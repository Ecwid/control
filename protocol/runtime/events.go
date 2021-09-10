package runtime

/*
	Notification is issued every time when binding is called.
*/
type BindingCalled struct {
	Name               string             `json:"name"`
	Payload            string             `json:"payload"`
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}

/*
	Issued when console API was called.
*/
type ConsoleAPICalled struct {
	Type               string             `json:"type"`
	Args               []*RemoteObject    `json:"args"`
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
	Timestamp          Timestamp          `json:"timestamp"`
	StackTrace         *StackTrace        `json:"stackTrace,omitempty"`
	Context            string             `json:"context,omitempty"`
}

/*
	Issued when unhandled exception was revoked.
*/
type ExceptionRevoked struct {
	Reason      string `json:"reason"`
	ExceptionId int    `json:"exceptionId"`
}

/*
	Issued when exception was thrown and unhandled.
*/
type ExceptionThrown struct {
	Timestamp        Timestamp         `json:"timestamp"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

/*
	Issued when new execution context is created.
*/
type ExecutionContextCreated struct {
	Context *ExecutionContextDescription `json:"context"`
}

/*
	Issued when execution context is destroyed.
*/
type ExecutionContextDestroyed struct {
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}

/*
	Issued when all executionContexts were cleared in browser
*/
type ExecutionContextsCleared interface{}

/*
	Issued when object should be inspected (for example, as a result of inspect() command line API
call).
*/
type InspectRequested struct {
	Object *RemoteObject `json:"object"`
	Hints  interface{}   `json:"hints"`
}
