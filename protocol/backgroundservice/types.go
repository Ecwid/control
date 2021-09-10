package backgroundservice

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/serviceworker"
)

/*
	The Background Service that will be associated with the commands/events.
Every Background Service operates independently, but they share the same
API.
*/
type ServiceName string

/*
	A key-value pair for additional event information to pass along.
*/
type EventMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

/*

 */
type BackgroundServiceEvent struct {
	Timestamp                   common.TimeSinceEpoch        `json:"timestamp"`
	Origin                      string                       `json:"origin"`
	ServiceWorkerRegistrationId serviceworker.RegistrationID `json:"serviceWorkerRegistrationId"`
	Service                     ServiceName                  `json:"service"`
	EventName                   string                       `json:"eventName"`
	InstanceId                  string                       `json:"instanceId"`
	EventMetadata               []*EventMetadata             `json:"eventMetadata"`
}

type StartObservingArgs struct {
	Service ServiceName `json:"service"`
}

type StopObservingArgs struct {
	Service ServiceName `json:"service"`
}

type SetRecordingArgs struct {
	ShouldRecord bool        `json:"shouldRecord"`
	Service      ServiceName `json:"service"`
}

type ClearEventsArgs struct {
	Service ServiceName `json:"service"`
}
