package debugger

import (
	"github.com/ecwid/control/protocol"
)

/*
Continues execution until specific location is reached.
*/
func ContinueToLocation(c protocol.Caller, args ContinueToLocationArgs) error {
	return c.Call("Debugger.continueToLocation", args, nil)
}

/*
Disables debugger for given page.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Debugger.disable", nil, nil)
}

/*
	Enables debugger for the given page. Clients should not assume that the debugging has been

enabled until the result for this command is received.
*/
func Enable(c protocol.Caller, args EnableArgs) (*EnableVal, error) {
	var val = &EnableVal{}
	return val, c.Call("Debugger.enable", args, val)
}

/*
Evaluates expression on a given call frame.
*/
func EvaluateOnCallFrame(c protocol.Caller, args EvaluateOnCallFrameArgs) (*EvaluateOnCallFrameVal, error) {
	var val = &EvaluateOnCallFrameVal{}
	return val, c.Call("Debugger.evaluateOnCallFrame", args, val)
}

/*
	Returns possible locations for breakpoint. scriptId in start and end range locations should be

the same.
*/
func GetPossibleBreakpoints(c protocol.Caller, args GetPossibleBreakpointsArgs) (*GetPossibleBreakpointsVal, error) {
	var val = &GetPossibleBreakpointsVal{}
	return val, c.Call("Debugger.getPossibleBreakpoints", args, val)
}

/*
Returns source for the script with given id.
*/
func GetScriptSource(c protocol.Caller, args GetScriptSourceArgs) (*GetScriptSourceVal, error) {
	var val = &GetScriptSourceVal{}
	return val, c.Call("Debugger.getScriptSource", args, val)
}

/*
 */
func DisassembleWasmModule(c protocol.Caller, args DisassembleWasmModuleArgs) (*DisassembleWasmModuleVal, error) {
	var val = &DisassembleWasmModuleVal{}
	return val, c.Call("Debugger.disassembleWasmModule", args, val)
}

/*
	Disassemble the next chunk of lines for the module corresponding to the

stream. If disassembly is complete, this API will invalidate the streamId
and return an empty chunk. Any subsequent calls for the now invalid stream
will return errors.
*/
func NextWasmDisassemblyChunk(c protocol.Caller, args NextWasmDisassemblyChunkArgs) (*NextWasmDisassemblyChunkVal, error) {
	var val = &NextWasmDisassemblyChunkVal{}
	return val, c.Call("Debugger.nextWasmDisassemblyChunk", args, val)
}

/*
Returns stack trace with given `stackTraceId`.
*/
func GetStackTrace(c protocol.Caller, args GetStackTraceArgs) (*GetStackTraceVal, error) {
	var val = &GetStackTraceVal{}
	return val, c.Call("Debugger.getStackTrace", args, val)
}

/*
Stops on the next JavaScript statement.
*/
func Pause(c protocol.Caller) error {
	return c.Call("Debugger.pause", nil, nil)
}

/*
Removes JavaScript breakpoint.
*/
func RemoveBreakpoint(c protocol.Caller, args RemoveBreakpointArgs) error {
	return c.Call("Debugger.removeBreakpoint", args, nil)
}

/*
	Restarts particular call frame from the beginning. The old, deprecated

behavior of `restartFrame` is to stay paused and allow further CDP commands
after a restart was scheduled. This can cause problems with restarting, so
we now continue execution immediatly after it has been scheduled until we
reach the beginning of the restarted frame.

To stay back-wards compatible, `restartFrame` now expects a `mode`
parameter to be present. If the `mode` parameter is missing, `restartFrame`
errors out.

The various return values are deprecated and `callFrames` is always empty.
Use the call frames from the `Debugger#paused` events instead, that fires
once V8 pauses at the beginning of the restarted function.
*/
func RestartFrame(c protocol.Caller, args RestartFrameArgs) error {
	return c.Call("Debugger.restartFrame", args, nil)
}

/*
Resumes JavaScript execution.
*/
func Resume(c protocol.Caller, args ResumeArgs) error {
	return c.Call("Debugger.resume", args, nil)
}

/*
Searches for given string in script content.
*/
func SearchInContent(c protocol.Caller, args SearchInContentArgs) (*SearchInContentVal, error) {
	var val = &SearchInContentVal{}
	return val, c.Call("Debugger.searchInContent", args, val)
}

/*
Enables or disables async call stacks tracking.
*/
func SetAsyncCallStackDepth(c protocol.Caller, args SetAsyncCallStackDepthArgs) error {
	return c.Call("Debugger.setAsyncCallStackDepth", args, nil)
}

/*
	Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in

scripts with url matching one of the patterns. VM will try to leave blackboxed script by
performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
*/
func SetBlackboxPatterns(c protocol.Caller, args SetBlackboxPatternsArgs) error {
	return c.Call("Debugger.setBlackboxPatterns", args, nil)
}

/*
	Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted

scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
Positions array contains positions where blackbox state is changed. First interval isn't
blackboxed. Array should be sorted.
*/
func SetBlackboxedRanges(c protocol.Caller, args SetBlackboxedRangesArgs) error {
	return c.Call("Debugger.setBlackboxedRanges", args, nil)
}

/*
Sets JavaScript breakpoint at a given location.
*/
func SetBreakpoint(c protocol.Caller, args SetBreakpointArgs) (*SetBreakpointVal, error) {
	var val = &SetBreakpointVal{}
	return val, c.Call("Debugger.setBreakpoint", args, val)
}

/*
Sets instrumentation breakpoint.
*/
func SetInstrumentationBreakpoint(c protocol.Caller, args SetInstrumentationBreakpointArgs) (*SetInstrumentationBreakpointVal, error) {
	var val = &SetInstrumentationBreakpointVal{}
	return val, c.Call("Debugger.setInstrumentationBreakpoint", args, val)
}

/*
	Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this

command is issued, all existing parsed scripts will have breakpoints resolved and returned in
`locations` property. Further matching script parsing will result in subsequent
`breakpointResolved` events issued. This logical breakpoint will survive page reloads.
*/
func SetBreakpointByUrl(c protocol.Caller, args SetBreakpointByUrlArgs) (*SetBreakpointByUrlVal, error) {
	var val = &SetBreakpointByUrlVal{}
	return val, c.Call("Debugger.setBreakpointByUrl", args, val)
}

/*
	Sets JavaScript breakpoint before each call to the given function.

If another function was created from the same source as a given one,
calling it will also trigger the breakpoint.
*/
func SetBreakpointOnFunctionCall(c protocol.Caller, args SetBreakpointOnFunctionCallArgs) (*SetBreakpointOnFunctionCallVal, error) {
	var val = &SetBreakpointOnFunctionCallVal{}
	return val, c.Call("Debugger.setBreakpointOnFunctionCall", args, val)
}

/*
Activates / deactivates all breakpoints on the page.
*/
func SetBreakpointsActive(c protocol.Caller, args SetBreakpointsActiveArgs) error {
	return c.Call("Debugger.setBreakpointsActive", args, nil)
}

/*
	Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions,

or caught exceptions, no exceptions. Initial pause on exceptions state is `none`.
*/
func SetPauseOnExceptions(c protocol.Caller, args SetPauseOnExceptionsArgs) error {
	return c.Call("Debugger.setPauseOnExceptions", args, nil)
}

/*
Changes return value in top frame. Available only at return break position.
*/
func SetReturnValue(c protocol.Caller, args SetReturnValueArgs) error {
	return c.Call("Debugger.setReturnValue", args, nil)
}

/*
	Edits JavaScript source live.

In general, functions that are currently on the stack can not be edited with
a single exception: If the edited function is the top-most stack frame and
that is the only activation of that function on the stack. In this case
the live edit will be successful and a `Debugger.restartFrame` for the
top-most function is automatically triggered.
*/
func SetScriptSource(c protocol.Caller, args SetScriptSourceArgs) (*SetScriptSourceVal, error) {
	var val = &SetScriptSourceVal{}
	return val, c.Call("Debugger.setScriptSource", args, val)
}

/*
Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
*/
func SetSkipAllPauses(c protocol.Caller, args SetSkipAllPausesArgs) error {
	return c.Call("Debugger.setSkipAllPauses", args, nil)
}

/*
	Changes value of variable in a callframe. Object-based scopes are not supported and must be

mutated manually.
*/
func SetVariableValue(c protocol.Caller, args SetVariableValueArgs) error {
	return c.Call("Debugger.setVariableValue", args, nil)
}

/*
Steps into the function call.
*/
func StepInto(c protocol.Caller, args StepIntoArgs) error {
	return c.Call("Debugger.stepInto", args, nil)
}

/*
Steps out of the function call.
*/
func StepOut(c protocol.Caller) error {
	return c.Call("Debugger.stepOut", nil, nil)
}

/*
Steps over the statement.
*/
func StepOver(c protocol.Caller, args StepOverArgs) error {
	return c.Call("Debugger.stepOver", args, nil)
}
