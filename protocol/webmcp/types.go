package webmcp

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

/*
Tool annotations
*/
type Annotation struct {
	ReadOnly         bool `json:"readOnly,omitempty"`
	UntrustedContent bool `json:"untrustedContent,omitempty"`
	Autosubmit       bool `json:"autosubmit,omitempty"`
}

/*
Represents the status of a tool invocation.
*/
type InvocationStatus string

/*
Definition of a tool that can be invoked.
*/
type Tool struct {
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	InputSchema   any                 `json:"inputSchema,omitempty"`
	Annotations   *Annotation         `json:"annotations,omitempty"`
	FrameId       common.FrameId      `json:"frameId"`
	BackendNodeId dom.BackendNodeId   `json:"backendNodeId,omitempty"`
	StackTrace    *runtime.StackTrace `json:"stackTrace,omitempty"`
}

/*
Definition of a tool that was removed.
*/
type RemovedTool struct {
	Name    string         `json:"name"`
	FrameId common.FrameId `json:"frameId"`
}

type InvokeToolArgs struct {
	FrameId  common.FrameId `json:"frameId"`
	ToolName string         `json:"toolName"`
	Input    any            `json:"input"`
}

type InvokeToolVal struct {
	InvocationId string `json:"invocationId"`
}

type CancelInvocationArgs struct {
	InvocationId string `json:"invocationId"`
}
