package debugger

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
Fired when breakpoint is resolved to an actual script and location.
*/
type BreakpointResolved struct {
	BreakpointId BreakpointId `json:"breakpointId"`
	Location     *Location    `json:"location"`
}

/*
Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
*/
type Paused struct {
	CallFrames        []*CallFrame          `json:"callFrames"`
	Reason            string                `json:"reason"`
	Data              interface{}           `json:"data,omitempty"`
	HitBreakpoints    []string              `json:"hitBreakpoints,omitempty"`
	AsyncStackTrace   *runtime.StackTrace   `json:"asyncStackTrace,omitempty"`
	AsyncStackTraceId *runtime.StackTraceId `json:"asyncStackTraceId,omitempty"`
}

/*
Fired when the virtual machine resumed execution.
*/
type Resumed interface{}

/*
Fired when virtual machine fails to parse the script.
*/
type ScriptFailedToParse struct {
	ScriptId                runtime.ScriptId           `json:"scriptId"`
	Url                     string                     `json:"url"`
	StartLine               int                        `json:"startLine"`
	StartColumn             int                        `json:"startColumn"`
	EndLine                 int                        `json:"endLine"`
	EndColumn               int                        `json:"endColumn"`
	ExecutionContextId      runtime.ExecutionContextId `json:"executionContextId"`
	Hash                    string                     `json:"hash"`
	ExecutionContextAuxData interface{}                `json:"executionContextAuxData,omitempty"`
	SourceMapURL            string                     `json:"sourceMapURL,omitempty"`
	HasSourceURL            bool                       `json:"hasSourceURL,omitempty"`
	IsModule                bool                       `json:"isModule,omitempty"`
	Length                  int                        `json:"length,omitempty"`
	StackTrace              *runtime.StackTrace        `json:"stackTrace,omitempty"`
	CodeOffset              int                        `json:"codeOffset,omitempty"`
	ScriptLanguage          ScriptLanguage             `json:"scriptLanguage,omitempty"`
	EmbedderName            string                     `json:"embedderName,omitempty"`
}

/*
	Fired when virtual machine parses script. This event is also fired for all known and uncollected

scripts upon enabling debugger.
*/
type ScriptParsed struct {
	ScriptId                runtime.ScriptId           `json:"scriptId"`
	Url                     string                     `json:"url"`
	StartLine               int                        `json:"startLine"`
	StartColumn             int                        `json:"startColumn"`
	EndLine                 int                        `json:"endLine"`
	EndColumn               int                        `json:"endColumn"`
	ExecutionContextId      runtime.ExecutionContextId `json:"executionContextId"`
	Hash                    string                     `json:"hash"`
	ExecutionContextAuxData interface{}                `json:"executionContextAuxData,omitempty"`
	IsLiveEdit              bool                       `json:"isLiveEdit,omitempty"`
	SourceMapURL            string                     `json:"sourceMapURL,omitempty"`
	HasSourceURL            bool                       `json:"hasSourceURL,omitempty"`
	IsModule                bool                       `json:"isModule,omitempty"`
	Length                  int                        `json:"length,omitempty"`
	StackTrace              *runtime.StackTrace        `json:"stackTrace,omitempty"`
	CodeOffset              int                        `json:"codeOffset,omitempty"`
	ScriptLanguage          ScriptLanguage             `json:"scriptLanguage,omitempty"`
	DebugSymbols            *DebugSymbols              `json:"debugSymbols,omitempty"`
	EmbedderName            string                     `json:"embedderName,omitempty"`
}
