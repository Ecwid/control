package runtime

import (
	"github.com/ecwid/control/protocol"
)

/*
Add handler to promise with given promise object id.
*/
func AwaitPromise(c protocol.Caller, args AwaitPromiseArgs) (*AwaitPromiseVal, error) {
	var val = &AwaitPromiseVal{}
	return val, c.Call("Runtime.awaitPromise", args, val)
}

/*
	Calls function with given declaration on the given object. Object group of the result is

inherited from the target object.
*/
func CallFunctionOn(c protocol.Caller, args CallFunctionOnArgs) (*CallFunctionOnVal, error) {
	var val = &CallFunctionOnVal{}
	return val, c.Call("Runtime.callFunctionOn", args, val)
}

/*
Compiles expression.
*/
func CompileScript(c protocol.Caller, args CompileScriptArgs) (*CompileScriptVal, error) {
	var val = &CompileScriptVal{}
	return val, c.Call("Runtime.compileScript", args, val)
}

/*
Disables reporting of execution contexts creation.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Runtime.disable", nil, nil)
}

/*
Discards collected exceptions and console API calls.
*/
func DiscardConsoleEntries(c protocol.Caller) error {
	return c.Call("Runtime.discardConsoleEntries", nil, nil)
}

/*
	Enables reporting of execution contexts creation by means of `executionContextCreated` event.

When the reporting gets enabled the event will be sent immediately for each existing execution
context.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Runtime.enable", nil, nil)
}

/*
Evaluates expression on global object.
*/
func Evaluate(c protocol.Caller, args EvaluateArgs) (*EvaluateVal, error) {
	var val = &EvaluateVal{}
	return val, c.Call("Runtime.evaluate", args, val)
}

/*
Returns the isolate id.
*/
func GetIsolateId(c protocol.Caller) (*GetIsolateIdVal, error) {
	var val = &GetIsolateIdVal{}
	return val, c.Call("Runtime.getIsolateId", nil, val)
}

/*
	Returns the JavaScript heap usage.

It is the total usage of the corresponding isolate not scoped to a particular Runtime.
*/
func GetHeapUsage(c protocol.Caller) (*GetHeapUsageVal, error) {
	var val = &GetHeapUsageVal{}
	return val, c.Call("Runtime.getHeapUsage", nil, val)
}

/*
	Returns properties of a given object. Object group of the result is inherited from the target

object.
*/
func GetProperties(c protocol.Caller, args GetPropertiesArgs) (*GetPropertiesVal, error) {
	var val = &GetPropertiesVal{}
	return val, c.Call("Runtime.getProperties", args, val)
}

/*
Returns all let, const and class variables from global scope.
*/
func GlobalLexicalScopeNames(c protocol.Caller, args GlobalLexicalScopeNamesArgs) (*GlobalLexicalScopeNamesVal, error) {
	var val = &GlobalLexicalScopeNamesVal{}
	return val, c.Call("Runtime.globalLexicalScopeNames", args, val)
}

/*
 */
func QueryObjects(c protocol.Caller, args QueryObjectsArgs) (*QueryObjectsVal, error) {
	var val = &QueryObjectsVal{}
	return val, c.Call("Runtime.queryObjects", args, val)
}

/*
Releases remote object with given id.
*/
func ReleaseObject(c protocol.Caller, args ReleaseObjectArgs) error {
	return c.Call("Runtime.releaseObject", args, nil)
}

/*
Releases all remote objects that belong to a given group.
*/
func ReleaseObjectGroup(c protocol.Caller, args ReleaseObjectGroupArgs) error {
	return c.Call("Runtime.releaseObjectGroup", args, nil)
}

/*
Tells inspected instance to run if it was waiting for debugger to attach.
*/
func RunIfWaitingForDebugger(c protocol.Caller) error {
	return c.Call("Runtime.runIfWaitingForDebugger", nil, nil)
}

/*
Runs script with given id in a given context.
*/
func RunScript(c protocol.Caller, args RunScriptArgs) (*RunScriptVal, error) {
	var val = &RunScriptVal{}
	return val, c.Call("Runtime.runScript", args, val)
}

/*
Enables or disables async call stacks tracking.
*/
func SetAsyncCallStackDepth(c protocol.Caller, args SetAsyncCallStackDepthArgs) error {
	return c.Call("Runtime.setAsyncCallStackDepth", args, nil)
}

/*
 */
func SetCustomObjectFormatterEnabled(c protocol.Caller, args SetCustomObjectFormatterEnabledArgs) error {
	return c.Call("Runtime.setCustomObjectFormatterEnabled", args, nil)
}

/*
 */
func SetMaxCallStackSizeToCapture(c protocol.Caller, args SetMaxCallStackSizeToCaptureArgs) error {
	return c.Call("Runtime.setMaxCallStackSizeToCapture", args, nil)
}

/*
	Terminate current or next JavaScript execution.

Will cancel the termination when the outer-most script execution ends.
*/
func TerminateExecution(c protocol.Caller) error {
	return c.Call("Runtime.terminateExecution", nil, nil)
}

/*
	If executionContextId is empty, adds binding with the given name on the

global objects of all inspected contexts, including those created later,
bindings survive reloads.
Binding function takes exactly one argument, this argument should be string,
in case of any other input, function throws an exception.
Each binding function call produces Runtime.bindingCalled notification.
*/
func AddBinding(c protocol.Caller, args AddBindingArgs) error {
	return c.Call("Runtime.addBinding", args, nil)
}

/*
	This method does not remove binding function from global object but

unsubscribes current runtime agent from Runtime.bindingCalled notifications.
*/
func RemoveBinding(c protocol.Caller, args RemoveBindingArgs) error {
	return c.Call("Runtime.removeBinding", args, nil)
}

/*
	This method tries to lookup and populate exception details for a

JavaScript Error object.
Note that the stackTrace portion of the resulting exceptionDetails will
only be populated if the Runtime domain was enabled at the time when the
Error was thrown.
*/
func GetExceptionDetails(c protocol.Caller, args GetExceptionDetailsArgs) (*GetExceptionDetailsVal, error) {
	var val = &GetExceptionDetailsVal{}
	return val, c.Call("Runtime.getExceptionDetails", args, val)
}
