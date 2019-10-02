package devtool

import "fmt"

// ExecutionContextDescription https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ExecutionContextDescription
type ExecutionContextDescription struct {
	ID      int64                  `json:"id"`
	Origin  string                 `json:"origin"`
	Name    string                 `json:"name"`
	AuxData map[string]interface{} `json:"auxData"`
}

// ExecutionContextCreated https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-executionContextCreated
type ExecutionContextCreated struct {
	Context *ExecutionContextDescription `json:"context"`
}

// ExecutionContextDestroyed https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-executionContextDestroyed
type ExecutionContextDestroyed struct {
	ExecutionContextID int64 `json:"executionContextId"`
}

// RemoteObject https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-RemoteObject
type RemoteObject struct {
	Type                string      `json:"type"`
	Subtype             string      `json:"subtype"`
	ClassName           string      `json:"className"`
	Value               interface{} `json:"value"`
	UnserializableValue string      `json:"unserializableValue"`
	Description         string      `json:"description"`
	ObjectID            string      `json:"objectId"`
}

// ExceptionDetails https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ExceptionDetails
type ExceptionDetails struct {
	ExceptionID        int64         `json:"exceptionId"`
	Text               string        `json:"text"`
	LineNumber         int64         `json:"lineNumber"`
	ColumnNumber       int64         `json:"columnNumber"`
	ScriptID           string        `json:"scriptId"`
	URL                string        `json:"url"`
	Exception          *RemoteObject `json:"exception"`
	ExecutionContextID int64         `json:"executionContextId"`
}

// PropertyDescriptor https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-PropertyDescriptor
type PropertyDescriptor struct {
	Name         string        `json:"name"`
	Value        *RemoteObject `json:"value"`
	Writable     bool          `json:"writable"`
	Get          *RemoteObject `json:"get"`
	Set          *RemoteObject `json:"set"`
	Configurable bool          `json:"configurable"`
	Enumerable   bool          `json:"enumerable"`
	WasThrown    bool          `json:"wasThrown"`
	IsOwn        bool          `json:"isOwn"`
	Symbol       *RemoteObject `json:"symbol"`
}

// CallArgument https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-CallArgument
type CallArgument struct {
	Value               interface{} `json:"value,omitempty"`
	UnserializableValue string      `json:"unserializableValue,omitempty"`
	ObjectID            string      `json:"objectId,omitempty"`
}

// EvaluatesExpression https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-evaluate
type EvaluatesExpression struct {
	Expression            string `json:"expression"`
	ObjectGroup           string `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI bool   `json:"includeCommandLineAPI,omitempty"`
	Silent                bool   `json:"silent,omitempty"`
	ContextID             int64  `json:"contextId,omitempty"`
	ReturnByValue         bool   `json:"returnByValue,omitempty"`
	GeneratePreview       bool   `json:"generatePreview,omitempty"`
	UserGesture           bool   `json:"userGesture,omitempty"`
	AwaitPromise          bool   `json:"awaitPromise,omitempty"`
	ThrowOnSideEffect     bool   `json:"throwOnSideEffect,omitempty"`
	Timeout               int64  `json:"timeout,omitempty"`
}

// EvaluatesResult https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-evaluate
type EvaluatesResult struct {
	Result           *RemoteObject     `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

func (e *ExceptionDetails) Error() string {
	return fmt.Sprintf("%+v", e.Exception)
}

// StackTrace https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-StackTrace
type StackTrace struct {
	Description string        `json:"description"`
	CallFrames  []*CallFrame  `json:"callFrames"`
	Parent      *StackTrace   `json:"parent"`
	ParentID    *StackTraceID `json:"parentId"`
}

// CallFrame https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-CallFrame
type CallFrame struct {
	FunctionName string `json:"functionName"`
	ScriptID     string `json:"scriptId"`
	URL          string `json:"url"`
	LineNumber   int64  `json:"lineNumber"`
	ColumnNumber int64  `json:"columnNumber"`
}

// StackTraceID https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-StackTraceId
type StackTraceID struct {
	ID         string `json:"id"`
	DebuggerID string `json:"debuggerId"`
}

// PropertiesResult https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-getProperties
type PropertiesResult struct {
	Result           []*PropertyDescriptor `json:"result"`
	ExceptionDetails *ExceptionDetails     `json:"exceptionDetails"`
}

// ConsoleAPICalled https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-consoleAPICalled
type ConsoleAPICalled struct {
	Type               string          `json:"type"`
	Args               []*RemoteObject `json:"args"`
	ExecutionContextID int64           `json:"executionContextId"`
	Timestamp          float64         `json:"timestamp"`
	StackTrace         *StackTrace     `json:"stackTrace"`
	Context            string          `json:"context"`
}

// Bool RemoteObject as bool value
func (r *RemoteObject) Bool() bool {
	return r.Type == "boolean" && r.Value.(bool)
}
