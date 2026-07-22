package webmcp

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/runtime"
)

/*
Event fired when new tools are added.
*/
type ToolsAdded struct {
	Tools []*Tool `json:"tools"`
}

/*
Event fired when tools are removed.
*/
type ToolsRemoved struct {
	Tools []*RemovedTool `json:"tools"`
}

/*
Event fired when a tool invocation starts.
*/
type ToolInvoked struct {
	ToolName     string         `json:"toolName"`
	FrameId      common.FrameId `json:"frameId"`
	InvocationId string         `json:"invocationId"`
	Input        string         `json:"input"`
}

/*
Event fired when a tool invocation completes or fails.
*/
type ToolResponded struct {
	InvocationId string                `json:"invocationId"`
	Status       InvocationStatus      `json:"status"`
	Output       any                   `json:"output,omitempty"`
	ErrorText    string                `json:"errorText,omitempty"`
	Exception    *runtime.RemoteObject `json:"exception,omitempty"`
}
