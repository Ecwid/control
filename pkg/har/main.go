package har

import (
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/devtool"
)

// Rec har recorder
type Rec struct {
	run         bool
	unsubscribe func()
	Log         *HAR `json:"log"`
}

// Unsubscribe ...
func (r *Rec) Unsubscribe() {
	if r.run {
		r.unsubscribe()
		r.run = false
	}
}

// Serialize stop recording and serialize to bytes
func (r *Rec) Serialize() ([]byte, error) {
	r.Unsubscribe()
	return json.Marshal(r)
}

// New start har records
func New(messages witness.Message) *Rec {

	r := &Rec{
		Log: &HAR{
			Version: "1.2",
			Creator: &Creator{
				Name:    "ecwid-witness",
				Version: "0.1",
			},
			Pages:   make([]*Page, 0),
			Entries: make([]*Entry, 0),
		},
		run: true,
	}

	var events chan *witness.Event
	events, r.unsubscribe = messages.Listen(
		"Page.frameStartedLoading",
		"Page.navigatedWithinDocument",
		"Network.requestWillBeSent",
		"Network.responseReceived",
		"Network.dataReceived",
		"Network.loadingFinished",
		"Network.requestServedFromCache",
		"Network.loadingFailed",
		"Page.loadEventFired",
		"Page.domContentEventFired",
	)
	go process(r, events)

	return r
}

func (r *Rec) getOrAddPage(ID string) *Page {
	for _, p := range r.Log.Pages {
		if p.ID == ID {
			return p
		}
	}
	p := &Page{
		ID:              ID,
		StartedDateTime: time.Now(),
		PageTiming:      &PageTiming{},
	}
	r.Log.Pages = append(r.Log.Pages, p)
	return p
}

func (r *Rec) getCurPage() *Page {
	return r.Log.Pages[len(r.Log.Pages)-1]
}

func (r *Rec) entryByRequestID(id string) *Entry {
	for _, e := range r.Log.Entries {
		if e.Request.ID == id {
			return e
		}
	}
	return nil
}

func process(rec *Rec, events <-chan *witness.Event) {
	for e := range events {
		var err error

		switch e.Method {

		case "Page.frameStartedLoading", "Page.navigatedWithinDocument":
			j := map[string]string{}
			if err := json.Unmarshal(e.Params, &j); err != nil {
				log.Print(err)
				continue
			}
			page := rec.getOrAddPage(j["frameId"])
			if e.Method == "Page.navigatedWithinDocument" {
				page.Title = j["url"]
			}

		case "Network.requestWillBeSent":

			willBeSent := new(devtool.RequestWillBeSent)
			if err = json.Unmarshal(e.Params, willBeSent); err != nil {
				log.Print(err)
				continue
			}

			page := rec.getOrAddPage(willBeSent.FrameID)

			request := &Request{
				ID:        willBeSent.RequestID,
				Method:    willBeSent.Request.Method,
				URL:       willBeSent.Request.URL,
				BodySize:  len(willBeSent.Request.PostData),
				Timestamp: willBeSent.Timestamp,
				Headers:   parseHeaders(willBeSent.Request.Headers),
			}

			if willBeSent.Request.HasPostData {
				mimeType := ""
				if ct, ok := willBeSent.Request.Headers["Content-Type"]; ok {
					mimeType = ct.(string)
				}
				request.PostData = &PostData{
					MimeType: mimeType,
					Text:     willBeSent.Request.PostData,
				}
			}

			if request.QueryString, err = parseURL(willBeSent.Request.URL); err != nil {
				log.Print(err)
				continue
			}

			entry := &Entry{
				StartedDateTime: epoch(willBeSent.WallTime), //epoch float64, eg 1440589909.59248
				Pageref:         page.ID,
				Request:         request,
				Response:        &Response{Content: &Content{}},
				Cache:           &Cache{},
				PageTimings:     &PageTimings{},
			}

			if willBeSent.RedirectResponse != nil {
				entry := rec.entryByRequestID(willBeSent.RequestID)
				if entry == nil {
					log.Print("no original request for redirect response " + string(e.Params))
					continue
				}
				entry.Request.ID = entry.Request.ID + "r"
				addResponse(entry, willBeSent.Timestamp, willBeSent.RedirectResponse)
				entry.Response.RedirectURL = willBeSent.Request.URL
				entry.PageTimings.Receive = 0.0
			}
			rec.Log.Entries = append(rec.Log.Entries, entry)

		case "Network.responseReceived":
			received := new(devtool.ResponseReceived)
			if err = json.Unmarshal(e.Params, received); err != nil {
				log.Print(err)
				continue
			}
			entry := rec.entryByRequestID(received.RequestID)
			if entry == nil {
				log.Print("no request for response " + string(e.Params))
				continue
			}
			addResponse(entry, received.Timestamp, received.Response)

		case "Network.dataReceived":
			received := new(devtool.DataReceived)
			if err = json.Unmarshal(e.Params, received); err != nil {
				log.Print(err)
				continue
			}
			entry := rec.entryByRequestID(received.RequestID)
			if entry == nil {
				log.Print("no request for dataReceived " + string(e.Params))
				continue
			}
			if entry.Response != nil && entry.Response.Content != nil {
				entry.Response.Content.Size += received.DataLength
			}

		case "Network.loadingFinished":
			finished := new(devtool.LoadingFinished)
			if err = json.Unmarshal(e.Params, finished); err != nil {
				log.Print(err)
				continue
			}
			entry := rec.entryByRequestID(finished.RequestID)
			if entry == nil {
				log.Print("no request for loadingFinished " + string(e.Params))
				continue
			}
			entry.Response.BodySize = int64(finished.EncodedDataLength) - int64(entry.Response.HeadersSize)
			entry.Response.Content.Compression = entry.Response.Content.Size - entry.Response.BodySize
			entry.Time = (finished.Timestamp - entry.Request.Timestamp) * 1000
			entry.PageTimings.Receive += int(finished.Timestamp * 1000)

		case "Network.requestServedFromCache":
			servedFromCache := new(devtool.ServedFromCache)
			if err = json.Unmarshal(e.Params, servedFromCache); err != nil {
				log.Print(err)
				continue
			}
			entry := rec.entryByRequestID(servedFromCache.RequestID)
			if entry == nil {
				log.Print("no request for requestServedFromCache")
				continue
			}
			entry.Cache.BeforeRequest = &CacheObject{
				LastAccess: "",
				ETag:       "",
				HitCount:   0,
			}

		case "Network.loadingFailed":
			loadingFailed := new(devtool.LoadingFailed)
			if err = json.Unmarshal(e.Params, loadingFailed); err != nil {
				log.Print(err)
				continue
			}
			entry := rec.entryByRequestID(loadingFailed.RequestID)
			if entry == nil {
				log.Print("no request for loadingFailed")
				continue
			}
			entry.Response = &Response{
				Status:     0,
				StatusText: loadingFailed.ErrorText,
				Timestamp:  loadingFailed.Timestamp,
			}

		case "Page.loadEventFired":
			if rec.getCurPage() == nil {
				continue
			}
			loadEvent := map[string]float64{}
			if err = json.Unmarshal(e.Params, &loadEvent); err != nil {
				log.Print(err)
				continue
			}
			page := rec.getCurPage()
			page.PageTiming.OnLoad = int64((loadEvent["timestamp"] - page.Timestamp) * 1000)

		case "Page.domContentEventFired":
			if rec.getCurPage() == nil {
				continue
			}
			domContentEvent := map[string]float64{}
			if err = json.Unmarshal(e.Params, &domContentEvent); err != nil {
				log.Print(err)
				continue
			}
			page := rec.getCurPage()
			page.PageTiming.OnContentLoad = int64((domContentEvent["timestamp"] - page.Timestamp) * 1000)
		}

	}
}

func addResponse(entry *Entry, timestamp float64, response *devtool.Response) {
	entry.Request.HTTPVersion = response.Protocol
	entry.Request.Headers = parseHeaders(response.RequestHeaders)
	entry.Request.SetHeadersSize()
	entry.Request.SetCookies()

	// cookies := response.parseCookie()
	// for _, c := range cookies {
	// 	resp.Cookies = append(resp.Cookies, &Cookie{
	// 		Name:    c.Name,
	// 		Value:   c.Value,
	// 		Path:    c.Path,
	// 		Domain:  c.Domain,
	// 		Expires: string(c.Expires),
	// 		Secure:  c.Secure,
	// 	})
	// }

	entry.Response = &Response{
		Status:      response.Status,
		StatusText:  response.StatusText,
		HTTPVersion: entry.Request.HTTPVersion,
		Headers:     parseHeaders(response.Headers),
		Timestamp:   timestamp,
	}
	entry.Response.SetHeadersSize()
	entry.ServerIPAddress = response.RemoteIPAddress
	entry.Response.Content = &Content{
		MimeType: response.MimeType,
	}

	if response.Timing == nil {
		return
	}

	entry.PageTimings = &PageTimings{
		Blocked: int(math.Max(0.0, response.Timing.DNSStart)),
		DNS:     int(math.Max(0.0, response.Timing.DNSEnd-response.Timing.DNSStart)),
		Connect: int(math.Max(0.0, response.Timing.ConnectEnd-response.Timing.ConnectStart)),
		Send:    int(math.Max(0.0, response.Timing.SendEnd-response.Timing.SendStart)),
		Wait:    int(math.Max(0.0, response.Timing.ReceiveHeadersEnd-response.Timing.SendEnd)),
		SSL:     int(math.Max(0.0, response.Timing.SSLEnd-response.Timing.SSLStart)),
		Receive: int(0.0 - (response.Timing.RequestTime*1000 + response.Timing.ReceiveHeadersEnd)),
	}
}
