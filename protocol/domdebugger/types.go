package domdebugger

import (
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

/*
DOM breakpoint type.
*/
type DOMBreakpointType string

/*
CSP Violation type.
*/
type CSPViolationType string

/*
Object event listener.
*/
type EventListener struct {
	Type            string                `json:"type"`
	UseCapture      bool                  `json:"useCapture"`
	Passive         bool                  `json:"passive"`
	Once            bool                  `json:"once"`
	ScriptId        runtime.ScriptId      `json:"scriptId"`
	LineNumber      int                   `json:"lineNumber"`
	ColumnNumber    int                   `json:"columnNumber"`
	Handler         *runtime.RemoteObject `json:"handler,omitempty"`
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"`
	BackendNodeId   dom.BackendNodeId     `json:"backendNodeId,omitempty"`
}

type GetEventListenersArgs struct {
	ObjectId runtime.RemoteObjectId `json:"objectId"`
	Depth    int                    `json:"depth,omitempty"`
	Pierce   bool                   `json:"pierce,omitempty"`
}

type GetEventListenersVal struct {
	Listeners []*EventListener `json:"listeners"`
}

type RemoveDOMBreakpointArgs struct {
	NodeId dom.NodeId        `json:"nodeId"`
	Type   DOMBreakpointType `json:"type"`
}

type RemoveEventListenerBreakpointArgs struct {
	EventName  string `json:"eventName"`
	TargetName string `json:"targetName,omitempty"`
}

type RemoveInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"`
}

type RemoveXHRBreakpointArgs struct {
	Url string `json:"url"`
}

type SetBreakOnCSPViolationArgs struct {
	ViolationTypes []CSPViolationType `json:"violationTypes"`
}

type SetDOMBreakpointArgs struct {
	NodeId dom.NodeId        `json:"nodeId"`
	Type   DOMBreakpointType `json:"type"`
}

type SetEventListenerBreakpointArgs struct {
	EventName  string `json:"eventName"`
	TargetName string `json:"targetName,omitempty"`
}

type SetInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"`
}

type SetXHRBreakpointArgs struct {
	Url string `json:"url"`
}
