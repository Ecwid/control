package witness

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ecwid/witness/pkg/devtool"
	"github.com/ecwid/witness/pkg/log"
	"github.com/gorilla/websocket"
)

type rpcMessage struct {
	ID        int64       `json:"id"`
	Method    string      `json:"method"`
	Params    interface{} `json:"params,omitempty"`
	SessionID string      `json:"sessionId,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    bytes  `json:"data,omitempty"`
}

func (e rpcError) Error() string {
	return e.Message
}

type rpcEvent struct {
	Method string `json:"method"`
	Params bytes  `json:"params"`
}

// rpcResponse cdp message response
type rpcResponse struct {
	ID        int64     `json:"id"`
	Result    bytes     `json:"result"`
	SessionID string    `json:"sessionId,omitempty"`
	Error     *rpcError `json:"error,omitempty"`
}

type rpcRecv struct {
	rpcEvent
	rpcResponse
}

func (r rpcRecv) isEvent() bool {
	return r.ID == 0 && r.Method != ""
}

func (r rpcResponse) isError() bool {
	return r.Error != nil
}

type stats struct {
	messages int64
	events   int64
}

type timeouts struct {
	Navigation time.Duration
	Implicitly time.Duration
	Poll       time.Duration
	WSTimeout  time.Duration
}

var dto = &timeouts{
	Navigation: time.Second * 40,
	Implicitly: time.Second * 30,
	Poll:       time.Millisecond * 500,
	WSTimeout:  time.Minute * 1,
}

// CDP ...
type CDP struct {
	mx          sync.Mutex
	nextID      int64
	conn        *websocket.Conn
	chanSend    chan rpcMessage
	chanReceive map[int64]chan rpcResponse
	close       chan bool
	sessions    map[string]*Session
	context     context.Context
	stats       *stats
	Timeouts    *timeouts
}

// New create new client to interact with browser by CDP
func New(cntx context.Context, webSocketURL string) (*CDP, error) {
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		return nil, err
	}
	c := &CDP{
		conn:        conn,
		chanSend:    make(chan rpcMessage),            /* channel for message sending */
		chanReceive: make(map[int64]chan rpcResponse), /* channel for receive message response */
		sessions:    make(map[string]*Session),
		close:       make(chan bool),
		context:     cntx,
		stats:       new(stats),
		Timeouts:    dto,
	}
	go c.transmitter()
	go c.receiver()
	if _, err := c.get(c.sendOverProtocol("", "Target.setDiscoverTargets", Map{"discover": true})); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CDP) get(ch chan rpcResponse) (rpcResponse, error) {
	select {
	case message := <-ch:
		return message, nil
	case <-c.context.Done():
		return rpcResponse{}, c.context.Err()
	case <-time.After(c.Timeouts.WSTimeout):
		return rpcResponse{}, ErrDevtoolTimeout
	}
}

// DefaultPage ...
func (c *CDP) DefaultPage() (*Session, error) {
	tick := time.NewTicker(time.Millisecond * 250)
	implicitly := time.NewTimer(time.Second * 10)
	defer tick.Stop()
	defer implicitly.Stop()
	for {
		select {
		case <-implicitly.C:
			return nil, ErrNoPageTarget
		case <-tick.C:
			ts, err := c.getTargets()
			if err != nil {
				return nil, err
			}
			for _, t := range ts {
				if t.Type == "page" {
					return c.newSession(t.TargetID)
				}
			}
		}
	}
}

func (c *CDP) getTargets() ([]*devtool.TargetInfo, error) {
	recv := c.sendOverProtocol("", "Target.getTargets", Map{})
	v, err := c.get(recv)
	if err != nil {
		return nil, err
	}
	t := devtool.TargetInfos{}
	if err := v.Result.Unmarshal(&t); err != nil {
		return nil, err
	}
	return t.TargetInfos, nil
}

func (c *CDP) deleteSession(sessionID string) {
	c.mx.Lock()
	delete(c.sessions, sessionID)
	c.mx.Unlock()
}

func (c *CDP) addSession(session *Session) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.sessions[session.id] = session
}

// SendOverProtocol send a message through cdp protocol
func (c *CDP) sendOverProtocol(sessionID string, method string, params interface{}) chan rpcResponse {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.nextID++
	c.chanReceive[c.nextID] = make(chan rpcResponse, 1)
	c.chanSend <- rpcMessage{
		ID:        c.nextID,
		SessionID: sessionID,
		Method:    method,
		Params:    params,
	}
	return c.chanReceive[c.nextID]
}

// Close close browser and websocket connection
func (c *CDP) Close() {
	<-c.sendOverProtocol("", "Browser.close", Map{})
	close(c.close)
	log.Printf(log.LevelFatal, "messages sent: %d", c.stats.messages)
	log.Printf(log.LevelFatal, "events received: %d", c.stats.events)
}

func tostring(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func (c *CDP) transmitter() {
	for {
		select {
		case req := <-c.chanSend:
			log.Printf(log.LevelProtocolMessage, "\033[1;36msend -> %s\033[0m", tostring(req))
			if err := c.conn.WriteJSON(req); err != nil {
				log.Print(log.LevelFatal, err)
				break
			}
		case <-c.context.Done():
			log.Print(log.LevelFatal, c.context.Err())
			close(c.close)
		case <-c.close:
			return
		}
	}
}

func (c *CDP) incoming(message rpcRecv) {
	if message.isEvent() {
		log.Printf(log.LevelProtocolEvents, "\033[1;30mevent <- %s\033[0m", tostring(message.rpcEvent))
		c.stats.events++
		if message.SessionID != "" {
			c.sessions[message.SessionID].incomingEvent <- message.rpcEvent
		} else {
			for _, e := range c.sessions {
				e.incomingEvent <- message.rpcEvent
			}
		}
	} else {
		c.stats.messages++
		if message.isError() {
			log.Printf(log.LevelProtocolErrors, "\033[1;31mrecv <- %s\033[0m", tostring(message.rpcResponse.Error))
		} else {
			log.Printf(log.LevelProtocolMessage, "\033[1;34mrecv <- %s\033[0m", tostring(message.rpcResponse))
		}
		if recv, ok := c.chanReceive[message.ID]; ok {
			recv <- message.rpcResponse
			c.mx.Lock()
			delete(c.chanReceive, message.ID)
			c.mx.Unlock()
		}
	}
}

func (c *CDP) receiver() {
	defer c.conn.Close()
	var message rpcRecv
	var err error
	for {
		select {
		case <-c.close:
			return
		default:
			message = rpcRecv{}
			if err = c.conn.ReadJSON(&message); err != nil {
				log.Print(log.LevelFatal, err)
				return
			}
			c.incoming(message)
		}
	}
}
