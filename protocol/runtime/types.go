package runtime

/*
Unique script identifier.
*/
type ScriptId string

/*
	Represents the value serialiazed by the WebDriver BiDi specification

https://w3c.github.io/webdriver-bidi.
*/
type DeepSerializedValue struct {
	Type                     string `json:"type"`
	Value                    any    `json:"value,omitempty"`
	ObjectId                 string `json:"objectId,omitempty"`
	WeakLocalObjectReference int    `json:"weakLocalObjectReference,omitempty"`
}

/*
Unique object identifier.
*/
type RemoteObjectId string

/*
	Primitive value which cannot be JSON-stringified. Includes values `-0`, `NaN`, `Infinity`,

`-Infinity`, and bigint literals.
*/
type UnserializableValue string

/*
Mirror object referencing original JavaScript object.
*/
type RemoteObject struct {
	Type                string               `json:"type"`
	Subtype             string               `json:"subtype,omitempty"`
	ClassName           string               `json:"className,omitempty"`
	Value               any                  `json:"value,omitempty"`
	UnserializableValue UnserializableValue  `json:"unserializableValue,omitempty"`
	Description         string               `json:"description,omitempty"`
	DeepSerializedValue *DeepSerializedValue `json:"deepSerializedValue,omitempty"`
	ObjectId            RemoteObjectId       `json:"objectId,omitempty"`
	Preview             *ObjectPreview       `json:"preview,omitempty"`
	CustomPreview       *CustomPreview       `json:"customPreview,omitempty"`
}

/*
 */
type CustomPreview struct {
	Header       string         `json:"header"`
	BodyGetterId RemoteObjectId `json:"bodyGetterId,omitempty"`
}

/*
Object containing abbreviated remote object value.
*/
type ObjectPreview struct {
	Type        string             `json:"type"`
	Subtype     string             `json:"subtype,omitempty"`
	Description string             `json:"description,omitempty"`
	Overflow    bool               `json:"overflow"`
	Properties  []*PropertyPreview `json:"properties"`
	Entries     []*EntryPreview    `json:"entries,omitempty"`
}

/*
 */
type PropertyPreview struct {
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Value        string         `json:"value,omitempty"`
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"`
	Subtype      string         `json:"subtype,omitempty"`
}

/*
 */
type EntryPreview struct {
	Key   *ObjectPreview `json:"key,omitempty"`
	Value *ObjectPreview `json:"value"`
}

/*
Object property descriptor.
*/
type PropertyDescriptor struct {
	Name         string        `json:"name"`
	Value        *RemoteObject `json:"value,omitempty"`
	Writable     bool          `json:"writable,omitempty"`
	Get          *RemoteObject `json:"get,omitempty"`
	Set          *RemoteObject `json:"set,omitempty"`
	Configurable bool          `json:"configurable"`
	Enumerable   bool          `json:"enumerable"`
	WasThrown    bool          `json:"wasThrown,omitempty"`
	IsOwn        bool          `json:"isOwn,omitempty"`
	Symbol       *RemoteObject `json:"symbol,omitempty"`
}

/*
Object internal property descriptor. This property isn't normally visible in JavaScript code.
*/
type InternalPropertyDescriptor struct {
	Name  string        `json:"name"`
	Value *RemoteObject `json:"value,omitempty"`
}

/*
Object private field descriptor.
*/
type PrivatePropertyDescriptor struct {
	Name  string        `json:"name"`
	Value *RemoteObject `json:"value,omitempty"`
	Get   *RemoteObject `json:"get,omitempty"`
	Set   *RemoteObject `json:"set,omitempty"`
}

/*
	Represents function call argument. Either remote object id `objectId`, primitive `value`,

unserializable primitive value or neither of (for undefined) them should be specified.
*/
type CallArgument struct {
	Value               interface{}         `json:"value,omitempty"`
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`
	ObjectId            RemoteObjectId      `json:"objectId,omitempty"`
}

/*
Id of an execution context.
*/
type ExecutionContextId int

/*
Description of an isolated world.
*/
type ExecutionContextDescription struct {
	Id       ExecutionContextId `json:"id"`
	Origin   string             `json:"origin"`
	Name     string             `json:"name"`
	UniqueId string             `json:"uniqueId"`
	AuxData  interface{}        `json:"auxData,omitempty"`
}

/*
	Detailed information about exception (or error) that was thrown during script compilation or

execution.
*/
type ExceptionDetails struct {
	ExceptionId        int                `json:"exceptionId"`
	Text               string             `json:"text"`
	LineNumber         int                `json:"lineNumber"`
	ColumnNumber       int                `json:"columnNumber"`
	ScriptId           ScriptId           `json:"scriptId,omitempty"`
	Url                string             `json:"url,omitempty"`
	StackTrace         *StackTrace        `json:"stackTrace,omitempty"`
	Exception          *RemoteObject      `json:"exception,omitempty"`
	ExecutionContextId ExecutionContextId `json:"executionContextId,omitempty"`
	ExceptionMetaData  interface{}        `json:"exceptionMetaData,omitempty"`
}

/*
Number of milliseconds since epoch.
*/
type Timestamp float64

/*
Number of milliseconds.
*/
type TimeDelta float64

/*
Stack entry for runtime errors and assertions.
*/
type CallFrame struct {
	FunctionName string   `json:"functionName"`
	ScriptId     ScriptId `json:"scriptId"`
	Url          string   `json:"url"`
	LineNumber   int      `json:"lineNumber"`
	ColumnNumber int      `json:"columnNumber"`
}

/*
Call frames for assertions or error messages.
*/
type StackTrace struct {
	Description string        `json:"description,omitempty"`
	CallFrames  []*CallFrame  `json:"callFrames"`
	Parent      *StackTrace   `json:"parent,omitempty"`
	ParentId    *StackTraceId `json:"parentId,omitempty"`
}

/*
Unique identifier of current debugger.
*/
type UniqueDebuggerId string

/*
	If `debuggerId` is set stack trace comes from another debugger and can be resolved there. This

allows to track cross-debugger calls. See `Runtime.StackTrace` and `Debugger.paused` for usages.
*/
type StackTraceId struct {
	Id         string           `json:"id"`
	DebuggerId UniqueDebuggerId `json:"debuggerId,omitempty"`
}

type AwaitPromiseArgs struct {
	PromiseObjectId RemoteObjectId `json:"promiseObjectId"`
	ReturnByValue   bool           `json:"returnByValue,omitempty"`
	GeneratePreview bool           `json:"generatePreview,omitempty"`
}

type AwaitPromiseVal struct {
	Result           *RemoteObject     `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type CallFunctionOnArgs struct {
	FunctionDeclaration  string                `json:"functionDeclaration"`
	ObjectId             RemoteObjectId        `json:"objectId,omitempty"`
	Arguments            []*CallArgument       `json:"arguments,omitempty"`
	Silent               bool                  `json:"silent,omitempty"`
	ReturnByValue        bool                  `json:"returnByValue,omitempty"`
	GeneratePreview      bool                  `json:"generatePreview,omitempty"`
	UserGesture          bool                  `json:"userGesture,omitempty"`
	AwaitPromise         bool                  `json:"awaitPromise,omitempty"`
	ExecutionContextId   ExecutionContextId    `json:"executionContextId,omitempty"`
	ObjectGroup          string                `json:"objectGroup,omitempty"`
	ThrowOnSideEffect    bool                  `json:"throwOnSideEffect,omitempty"`
	UniqueContextId      string                `json:"uniqueContextId,omitempty"`
	SerializationOptions *SerializationOptions `json:"serializationOptions,omitempty"`
}

type CallFunctionOnVal struct {
	Result           *RemoteObject     `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type CompileScriptArgs struct {
	Expression         string             `json:"expression"`
	SourceURL          string             `json:"sourceURL"`
	PersistScript      bool               `json:"persistScript"`
	ExecutionContextId ExecutionContextId `json:"executionContextId,omitempty"`
}

type CompileScriptVal struct {
	ScriptId         ScriptId          `json:"scriptId,omitempty"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type SerializationOptions struct {
	Serialization string `json:"serialization,omitempty"`
	MaxDepth      int    `json:"maxDepth,omitempty"`
}

type EvaluateArgs struct {
	Expression                  string                `json:"expression"`
	ObjectGroup                 string                `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI       bool                  `json:"includeCommandLineAPI,omitempty"`
	Silent                      bool                  `json:"silent,omitempty"`
	ContextId                   ExecutionContextId    `json:"contextId,omitempty"`
	ReturnByValue               bool                  `json:"returnByValue,omitempty"`
	GeneratePreview             bool                  `json:"generatePreview,omitempty"`
	UserGesture                 bool                  `json:"userGesture,omitempty"`
	AwaitPromise                bool                  `json:"awaitPromise,omitempty"`
	ThrowOnSideEffect           bool                  `json:"throwOnSideEffect,omitempty"`
	Timeout                     TimeDelta             `json:"timeout,omitempty"`
	DisableBreaks               bool                  `json:"disableBreaks,omitempty"`
	ReplMode                    bool                  `json:"replMode,omitempty"`
	AllowUnsafeEvalBlockedByCSP bool                  `json:"allowUnsafeEvalBlockedByCSP,omitempty"`
	UniqueContextId             string                `json:"uniqueContextId,omitempty"`
	SerializationOptions        *SerializationOptions `json:"serializationOptions,omitempty"`
}

type EvaluateVal struct {
	Result           *RemoteObject     `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type GetIsolateIdVal struct {
	Id string `json:"id"`
}

type GetHeapUsageVal struct {
	UsedSize  float64 `json:"usedSize"`
	TotalSize float64 `json:"totalSize"`
}

type GetPropertiesArgs struct {
	ObjectId                 RemoteObjectId `json:"objectId"`
	OwnProperties            bool           `json:"ownProperties,omitempty"`
	AccessorPropertiesOnly   bool           `json:"accessorPropertiesOnly,omitempty"`
	GeneratePreview          bool           `json:"generatePreview,omitempty"`
	NonIndexedPropertiesOnly bool           `json:"nonIndexedPropertiesOnly,omitempty"`
}

type GetPropertiesVal struct {
	Result             []*PropertyDescriptor         `json:"result"`
	InternalProperties []*InternalPropertyDescriptor `json:"internalProperties,omitempty"`
	PrivateProperties  []*PrivatePropertyDescriptor  `json:"privateProperties,omitempty"`
	ExceptionDetails   *ExceptionDetails             `json:"exceptionDetails,omitempty"`
}

type GlobalLexicalScopeNamesArgs struct {
	ExecutionContextId ExecutionContextId `json:"executionContextId,omitempty"`
}

type GlobalLexicalScopeNamesVal struct {
	Names []string `json:"names"`
}

type QueryObjectsArgs struct {
	PrototypeObjectId RemoteObjectId `json:"prototypeObjectId"`
	ObjectGroup       string         `json:"objectGroup,omitempty"`
}

type QueryObjectsVal struct {
	Objects *RemoteObject `json:"objects"`
}

type ReleaseObjectArgs struct {
	ObjectId RemoteObjectId `json:"objectId"`
}

type ReleaseObjectGroupArgs struct {
	ObjectGroup string `json:"objectGroup"`
}

type RunScriptArgs struct {
	ScriptId              ScriptId           `json:"scriptId"`
	ExecutionContextId    ExecutionContextId `json:"executionContextId,omitempty"`
	ObjectGroup           string             `json:"objectGroup,omitempty"`
	Silent                bool               `json:"silent,omitempty"`
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI,omitempty"`
	ReturnByValue         bool               `json:"returnByValue,omitempty"`
	GeneratePreview       bool               `json:"generatePreview,omitempty"`
	AwaitPromise          bool               `json:"awaitPromise,omitempty"`
}

type RunScriptVal struct {
	Result           *RemoteObject     `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type SetAsyncCallStackDepthArgs struct {
	MaxDepth int `json:"maxDepth"`
}

type SetCustomObjectFormatterEnabledArgs struct {
	Enabled bool `json:"enabled"`
}

type SetMaxCallStackSizeToCaptureArgs struct {
	Size int `json:"size"`
}

type AddBindingArgs struct {
	Name                 string `json:"name"`
	ExecutionContextName string `json:"executionContextName,omitempty"`
}

type RemoveBindingArgs struct {
	Name string `json:"name"`
}

type GetExceptionDetailsArgs struct {
	ErrorObjectId RemoteObjectId `json:"errorObjectId"`
}

type GetExceptionDetailsVal struct {
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}
