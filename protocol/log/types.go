package log

import (
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Log entry.
*/
type LogEntry struct {
	Source           string                  `json:"source"`
	Level            string                  `json:"level"`
	Text             string                  `json:"text"`
	Timestamp        runtime.Timestamp       `json:"timestamp"`
	Url              string                  `json:"url,omitempty"`
	LineNumber       int                     `json:"lineNumber,omitempty"`
	StackTrace       *runtime.StackTrace     `json:"stackTrace,omitempty"`
	NetworkRequestId network.RequestId       `json:"networkRequestId,omitempty"`
	WorkerId         string                  `json:"workerId,omitempty"`
	Args             []*runtime.RemoteObject `json:"args,omitempty"`
}

/*
	Violation configuration setting.
*/
type ViolationSetting struct {
	Name      string  `json:"name"`
	Threshold float64 `json:"threshold"`
}

type StartViolationsReportArgs struct {
	Config []*ViolationSetting `json:"config"`
}
