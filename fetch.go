package control

import (
	"encoding/json"

	"github.com/ecwid/control/protocol/fetch"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/transport/observe"
)

type Fetch struct {
	RequestID fetch.RequestId
	Request   *network.Request
}

func (s Session) FetchEnable(urlPattern string, resType network.ResourceType, stage fetch.RequestStage) error {
	return s.transport.Call(s.Ctx, "", "Fetch.enable", fetch.EnableArgs{
		Patterns: []*fetch.RequestPattern{
			{
				UrlPattern:   urlPattern,
				ResourceType: resType,
				RequestStage: stage,
			},
		},
	}, nil)
}

func (s Session) FetchDisable() error {
	return s.transport.Call(s.Ctx, "", "Fetch.disable", nil, nil)
}

func (s *Session) NewFetchEventCondition(predicate func(f *Fetch) bool) *Promise {
	return s.NewEventCondition("Fetch.requestPaused", func(value observe.Value) (bool, error) {
		if value.Method == "" {
			return false, nil
		}
		var event = new(fetch.RequestPaused)
		if err := json.Unmarshal(value.Params, event); err != nil {
			return false, err
		}
		fetch := &Fetch{
			RequestID: event.RequestId,
			Request:   event.Request,
		}
		return predicate(fetch), nil
	})
}

func (s Session) FetchContinue(f *Fetch) error {
	return s.transport.Call(s.Ctx, "", "Fetch.continueRequest", fetch.ContinueRequestArgs{
		RequestId: f.RequestID,
		Url:       f.Request.Url,
		Method:    f.Request.Method,
		PostData:  []byte(f.Request.PostData),
	}, nil)
}
