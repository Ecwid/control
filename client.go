package witness

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ecwid/witness/pkg/devtool"
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

// Event ...
type Event struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

// rpcResponse cdp message response
type rpcResponse struct {
	ID        int64     `json:"id"`
	Result    bytes     `json:"result"`
	SessionID string    `json:"sessionId,omitempty"`
	Error     *rpcError `json:"error,omitempty"`
}

type rpcRecv struct {
	Event
	rpcResponse
}

func (r rpcRecv) isEvent() bool {
	return r.ID == 0 && r.Method != ""
}

func (r rpcResponse) isError() bool {
	return r.Error != nil
}

type stats struct {
	Messages int
	Events   int
	Sent     int
	Recv     int
}

type timeouts struct {
	Navigation time.Duration
	Implicitly time.Duration
	Poll       time.Duration
	WSTimeout  time.Duration
	internal   time.Duration
}

var dto = &timeouts{
	Navigation: time.Second * 60,
	Implicitly: time.Second * 60,
	Poll:       time.Millisecond * 500,
	WSTimeout:  time.Minute * 1,
	internal:   time.Second * 10,
}

// CDP chrome devtool protocol client
type CDP struct {
	mx          sync.Mutex
	nextID      int64
	conn        *websocket.Conn
	chanSend    chan []byte
	chanReceive map[int64]chan *rpcResponse
	close       chan bool
	sessions    sync.Map
	context     context.Context
	Stats       *stats
	Timeouts    *timeouts
	Logging     *wlog
}

// New create new client to interact with browser by CDP
func New(cntx context.Context, webSocketURL string) (*CDP, error) {
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		return nil, err
	}
	c := &CDP{
		conn:        conn,
		chanSend:    make(chan []byte),                 /* channel for message sending */
		chanReceive: make(map[int64]chan *rpcResponse), /* channel for receive message response */
		close:       make(chan bool),
		context:     cntx,
		Stats:       new(stats),
		Timeouts:    dto,
		Logging:     new(wlog),
	}
	c.Logging.Level = LevelProtocolErrors
	go c.transmitter()
	go c.receiver()
	if _, err := c.get(c.sendOverProtocol("", "Target.setDiscoverTargets", Map{"discover": true})); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CDP) get(ch chan *rpcResponse) (*rpcResponse, error) {
	select {
	case message := <-ch:
		return message, nil
	case <-c.context.Done():
		return nil, c.context.Err()
	case <-time.After(c.Timeouts.WSTimeout):
		return nil, ErrDevtoolTimeout
	}
}

// DefaultSession attach to default welcome page that opened after chrome start
func (c *CDP) DefaultSession() (*Session, error) {
	tick := time.NewTicker(time.Millisecond * 250)
	implicitly := time.NewTimer(c.Timeouts.internal)
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
	c.sessions.Delete(sessionID)
}

func (c *CDP) addSession(session *CDPSession) {
	c.sessions.Store(session.id, session)
}

// SendOverProtocol send a message through cdp protocol
func (c *CDP) sendOverProtocol(sessionID string, method string, params interface{}) chan *rpcResponse {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.nextID++
	c.chanReceive[c.nextID] = make(chan *rpcResponse, 1)
	message := &rpcMessage{
		ID:        c.nextID,
		SessionID: sessionID,
		Method:    method,
		Params:    params,
	}
	buf, err := json.Marshal(message)
	if err != nil {
		c.Logging.Print(LevelFatal, err)
	}
	c.chanSend <- buf
	return c.chanReceive[c.nextID]
}

// Close close browser and websocket connection
func (c *CDP) Close() {
	select {
	case <-c.sendOverProtocol("", "Browser.close", Map{}):
	case <-time.After(c.Timeouts.internal):
	}
	close(c.close)
}

func (c *CDP) transmitter() {
	for {
		select {
		case req := <-c.chanSend:
			c.Stats.Sent += len(req)
			c.Logging.Printf(LevelProtocolMessage, "\033[1;36msend -> %s\033[0m", string(req))
			if err := c.conn.WriteMessage(websocket.TextMessage, req); err != nil {
				c.Logging.Print(LevelFatal, err)
				break
			}
		case <-c.context.Done():
			c.Logging.Print(LevelFatal, c.context.Err())
			close(c.close)
		case <-c.close:
			return
		}
	}
}

func (c *CDP) incoming(recv []byte) {
	message := new(rpcRecv)
	if err := json.Unmarshal(recv, message); err != nil {
		c.Logging.Print(LevelFatal, err)
		return
	}
	if message.isEvent() {
		c.Logging.Printf(LevelProtocolEvents, "\033[1;30mevent <- %s\033[0m", string(recv))
		c.Stats.Events++
		if message.SessionID != "" {
			if sess, ok := c.sessions.Load(message.SessionID); ok {
				sess.(*CDPSession).incoming <- &message.Event
			}
		} else {
			c.sessions.Range(func(_ interface{}, v interface{}) bool {
				v.(*CDPSession).incoming <- &message.Event
				return true
			})
		}
	} else {
		c.Stats.Messages++
		if message.isError() {
			c.Logging.Printf(LevelProtocolErrors, "\033[1;31mrecv <- %s\033[0m", string(recv))
		} else {
			c.Logging.Printf(LevelProtocolMessage, "\033[1;34mrecv <- %s\033[0m", string(recv))
		}
		if recv, ok := c.chanReceive[message.ID]; ok {
			recv <- &message.rpcResponse
			c.mx.Lock()
			delete(c.chanReceive, message.ID)
			c.mx.Unlock()
		}
	}
}

func (c *CDP) receiver() {
	defer c.conn.Close()
	var err error
	var recv []byte
	for {
		select {
		case <-c.close:
			return
		default:
			_, recv, err = c.conn.ReadMessage()
			if err != nil {
				c.Logging.Print(LevelFatal, err)
				return
			}
			c.Stats.Recv += len(recv)
			c.incoming(recv)
		}
	}
}
