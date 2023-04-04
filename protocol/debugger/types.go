package debugger

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
Breakpoint identifier.
*/
type BreakpointId string

/*
Call frame identifier.
*/
type CallFrameId string

/*
Location in the source code.
*/
type Location struct {
	ScriptId     runtime.ScriptId `json:"scriptId"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber int              `json:"columnNumber,omitempty"`
}

/*
Location in the source code.
*/
type ScriptPosition struct {
	LineNumber   int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}

/*
Location range within one script.
*/
type LocationRange struct {
	ScriptId runtime.ScriptId `json:"scriptId"`
	Start    *ScriptPosition  `json:"start"`
	End      *ScriptPosition  `json:"end"`
}

/*
JavaScript call frame. Array of call frames form the call stack.
*/
type CallFrame struct {
	CallFrameId      CallFrameId           `json:"callFrameId"`
	FunctionName     string                `json:"functionName"`
	FunctionLocation *Location             `json:"functionLocation,omitempty"`
	Location         *Location             `json:"location"`
	ScopeChain       []*Scope              `json:"scopeChain"`
	This             *runtime.RemoteObject `json:"this"`
	ReturnValue      *runtime.RemoteObject `json:"returnValue,omitempty"`
	CanBeRestarted   bool                  `json:"canBeRestarted,omitempty"`
}

/*
Scope description.
*/
type Scope struct {
	Type          string                `json:"type"`
	Object        *runtime.RemoteObject `json:"object"`
	Name          string                `json:"name,omitempty"`
	StartLocation *Location             `json:"startLocation,omitempty"`
	EndLocation   *Location             `json:"endLocation,omitempty"`
}

/*
Search match for resource.
*/
type SearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`
	LineContent string  `json:"lineContent"`
}

/*
 */
type BreakLocation struct {
	ScriptId     runtime.ScriptId `json:"scriptId"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber int              `json:"columnNumber,omitempty"`
	Type         string           `json:"type,omitempty"`
}

/*
 */
type WasmDisassemblyChunk struct {
	Lines           []string `json:"lines"`
	BytecodeOffsets []int    `json:"bytecodeOffsets"`
}

/*
Enum of possible script languages.
*/
type ScriptLanguage string

/*
Debug symbols available for a wasm script.
*/
type DebugSymbols struct {
	Type        string `json:"type"`
	ExternalURL string `json:"externalURL,omitempty"`
}

type ContinueToLocationArgs struct {
	Location         *Location `json:"location"`
	TargetCallFrames string    `json:"targetCallFrames,omitempty"`
}

type EnableArgs struct {
	MaxScriptsCacheSize float64 `json:"maxScriptsCacheSize,omitempty"`
}

type EnableVal struct {
	DebuggerId runtime.UniqueDebuggerId `json:"debuggerId"`
}

type EvaluateOnCallFrameArgs struct {
	CallFrameId           CallFrameId       `json:"callFrameId"`
	Expression            string            `json:"expression"`
	ObjectGroup           string            `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI bool              `json:"includeCommandLineAPI,omitempty"`
	Silent                bool              `json:"silent,omitempty"`
	ReturnByValue         bool              `json:"returnByValue,omitempty"`
	GeneratePreview       bool              `json:"generatePreview,omitempty"`
	ThrowOnSideEffect     bool              `json:"throwOnSideEffect,omitempty"`
	Timeout               runtime.TimeDelta `json:"timeout,omitempty"`
}

type EvaluateOnCallFrameVal struct {
	Result           *runtime.RemoteObject     `json:"result"`
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type GetPossibleBreakpointsArgs struct {
	Start              *Location `json:"start"`
	End                *Location `json:"end,omitempty"`
	RestrictToFunction bool      `json:"restrictToFunction,omitempty"`
}

type GetPossibleBreakpointsVal struct {
	Locations []*BreakLocation `json:"locations"`
}

type GetScriptSourceArgs struct {
	ScriptId runtime.ScriptId `json:"scriptId"`
}

type GetScriptSourceVal struct {
	ScriptSource string `json:"scriptSource"`
	Bytecode     []byte `json:"bytecode,omitempty"`
}

type DisassembleWasmModuleArgs struct {
	ScriptId runtime.ScriptId `json:"scriptId"`
}

type DisassembleWasmModuleVal struct {
	StreamId            string                `json:"streamId,omitempty"`
	TotalNumberOfLines  int                   `json:"totalNumberOfLines"`
	FunctionBodyOffsets []int                 `json:"functionBodyOffsets"`
	Chunk               *WasmDisassemblyChunk `json:"chunk"`
}

type NextWasmDisassemblyChunkArgs struct {
	StreamId string `json:"streamId"`
}

type NextWasmDisassemblyChunkVal struct {
	Chunk *WasmDisassemblyChunk `json:"chunk"`
}

type GetStackTraceArgs struct {
	StackTraceId *runtime.StackTraceId `json:"stackTraceId"`
}

type GetStackTraceVal struct {
	StackTrace *runtime.StackTrace `json:"stackTrace"`
}

type RemoveBreakpointArgs struct {
	BreakpointId BreakpointId `json:"breakpointId"`
}

type RestartFrameArgs struct {
	CallFrameId CallFrameId `json:"callFrameId"`
	Mode        string      `json:"mode,omitempty"`
}

type ResumeArgs struct {
	TerminateOnResume bool `json:"terminateOnResume,omitempty"`
}

type SearchInContentArgs struct {
	ScriptId      runtime.ScriptId `json:"scriptId"`
	Query         string           `json:"query"`
	CaseSensitive bool             `json:"caseSensitive,omitempty"`
	IsRegex       bool             `json:"isRegex,omitempty"`
}

type SearchInContentVal struct {
	Result []*SearchMatch `json:"result"`
}

type SetAsyncCallStackDepthArgs struct {
	MaxDepth int `json:"maxDepth"`
}

type SetBlackboxPatternsArgs struct {
	Patterns []string `json:"patterns"`
}

type SetBlackboxedRangesArgs struct {
	ScriptId  runtime.ScriptId  `json:"scriptId"`
	Positions []*ScriptPosition `json:"positions"`
}

type SetBreakpointArgs struct {
	Location  *Location `json:"location"`
	Condition string    `json:"condition,omitempty"`
}

type SetBreakpointVal struct {
	BreakpointId   BreakpointId `json:"breakpointId"`
	ActualLocation *Location    `json:"actualLocation"`
}

type SetInstrumentationBreakpointArgs struct {
	Instrumentation string `json:"instrumentation"`
}

type SetInstrumentationBreakpointVal struct {
	BreakpointId BreakpointId `json:"breakpointId"`
}

type SetBreakpointByUrlArgs struct {
	LineNumber   int    `json:"lineNumber"`
	Url          string `json:"url,omitempty"`
	UrlRegex     string `json:"urlRegex,omitempty"`
	ScriptHash   string `json:"scriptHash,omitempty"`
	ColumnNumber int    `json:"columnNumber,omitempty"`
	Condition    string `json:"condition,omitempty"`
}

type SetBreakpointByUrlVal struct {
	BreakpointId BreakpointId `json:"breakpointId"`
	Locations    []*Location  `json:"locations"`
}

type SetBreakpointOnFunctionCallArgs struct {
	ObjectId  runtime.RemoteObjectId `json:"objectId"`
	Condition string                 `json:"condition,omitempty"`
}

type SetBreakpointOnFunctionCallVal struct {
	BreakpointId BreakpointId `json:"breakpointId"`
}

type SetBreakpointsActiveArgs struct {
	Active bool `json:"active"`
}

type SetPauseOnExceptionsArgs struct {
	State string `json:"state"`
}

type SetReturnValueArgs struct {
	NewValue *runtime.CallArgument `json:"newValue"`
}

type SetScriptSourceArgs struct {
	ScriptId             runtime.ScriptId `json:"scriptId"`
	ScriptSource         string           `json:"scriptSource"`
	DryRun               bool             `json:"dryRun,omitempty"`
	AllowTopFrameEditing bool             `json:"allowTopFrameEditing,omitempty"`
}

type SetScriptSourceVal struct {
	Status           string                    `json:"status"`
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type SetSkipAllPausesArgs struct {
	Skip bool `json:"skip"`
}

type SetVariableValueArgs struct {
	ScopeNumber  int                   `json:"scopeNumber"`
	VariableName string                `json:"variableName"`
	NewValue     *runtime.CallArgument `json:"newValue"`
	CallFrameId  CallFrameId           `json:"callFrameId"`
}

type StepIntoArgs struct {
	BreakOnAsyncCall bool             `json:"breakOnAsyncCall,omitempty"`
	SkipList         []*LocationRange `json:"skipList,omitempty"`
}

type StepOverArgs struct {
	SkipList []*LocationRange `json:"skipList,omitempty"`
}
