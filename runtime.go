package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

var ErrExecutionContextDestroyed = errors.New("execution context destroyed")

type RuntimeExceptionError struct {
	value *runtime.ExceptionDetails
}

type JSMapEntry struct {
	Key   any
	Value any
}

type JSRegExp struct {
	Pattern string
	Flags   string
}

func (e RuntimeExceptionError) Error() string {
	if e.value.Exception.Description != "" {
		return e.value.Exception.Description
	}
	b, _ := json.Marshal(e.value)
	return string(b)
}

type nodeType float64

const (
	nodeTypeElement               nodeType = 1  // An Element node like <p> or <div>
	nodeTypeAttribute             nodeType = 2  // An Attribute of an Element
	nodeTypeText                  nodeType = 3  // The actual Text inside an Element or Attr
	nodeTypeCDataSection          nodeType = 4  // A CDATASection
	nodeTypeProcessingInstruction nodeType = 7  // A ProcessingInstruction of an XML document
	nodeTypeComment               nodeType = 8  // A Comment node
	nodeTypeDocument              nodeType = 9  // A Document node
	nodeTypeDocumentType          nodeType = 10 // A DocumentType node
	nodeTypeFragment              nodeType = 11 // A DocumentFragment node
)

type ObjectHandle interface {
	GetRemoteObjectID() runtime.RemoteObjectId
}

type objectHandleValue runtime.RemoteObjectId

func (o objectHandleValue) GetRemoteObjectID() runtime.RemoteObjectId {
	return runtime.RemoteObjectId(o)
}

func objectHandle(id runtime.RemoteObjectId) ObjectHandle {
	return objectHandleValue(id)
}

func getNodeType(deepSerializedValue any) (nodeType, bool) {
	obj, ok := deepSerializedValue.(map[string]any)
	if !ok {
		return 0, false
	}
	raw, ok := obj["nodeType"].(float64)
	if !ok {
		return 0, false
	}
	return nodeType(raw), true
}

func deepUnserializePair(pair map[string]any) any {
	t, ok := pair["type"].(string)
	if !ok || t == "" {
		return pair["value"]
	}
	decoded, err := decodeDeepKind(t, pair["value"], "")
	if err != nil {
		return pair["value"]
	}
	return decoded
}

func deepUnserializeValue(value any) any {
	switch v := value.(type) {
	case map[string]any:
		return deepUnserializePair(v)
	default:
		return value
	}
}

func deepUnserializeArray(value any) any {
	if value == nil {
		return value
	}
	items, ok := value.([]any)
	if !ok {
		return value
	}
	arr := make([]any, 0, len(items))
	for _, item := range items {
		arr = append(arr, deepUnserializeValue(item))
	}
	return arr
}

func deepUnserializeObject(value any) any {
	if value == nil {
		return value
	}
	entries, ok := value.([]any)
	if !ok {
		return value
	}
	obj := make(map[string]any, len(entries))
	for _, entry := range entries {
		val, ok := entry.([]any)
		if !ok || len(val) < 2 {
			continue
		}
		key, ok := val[0].(string)
		if !ok || key == "" {
			continue
		}
		pair, ok := val[1].(map[string]any)
		if !ok {
			obj[key] = val[1]
			continue
		}
		obj[key] = deepUnserializePair(pair)
	}
	return obj
}

func deepUnserializeMap(value any) any {
	entries, ok := value.([]any)
	if !ok {
		return value
	}
	out := make([]JSMapEntry, 0, len(entries))
	for _, entry := range entries {
		pair, ok := entry.([]any)
		if !ok || len(pair) < 2 {
			continue
		}
		out = append(out, JSMapEntry{
			Key:   deepUnserializeValue(pair[0]),
			Value: deepUnserializeValue(pair[1]),
		})
	}
	return out
}

func deepUnserializeSet(value any) any {
	return deepUnserializeArray(value)
}

func deepUnserializeBigInt(value any) any {
	s, ok := value.(string)
	if !ok || s == "" {
		return value
	}
	s = strings.TrimSuffix(s, "n")
	bi := new(big.Int)
	if _, ok := bi.SetString(s, 10); ok {
		return bi
	}
	return value
}

func deepUnserializeDate(value any) any {
	s, ok := value.(string)
	if !ok || s == "" {
		return value
	}
	if parsed, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return parsed
	}
	if parsed, err := time.Parse(time.RFC3339, s); err == nil {
		return parsed
	}
	return value
}

func deepUnserializeRegExp(value any) any {
	obj, ok := value.(map[string]any)
	if !ok {
		return value
	}
	pattern, _ := obj["pattern"].(string)
	flags, _ := obj["flags"].(string)
	return JSRegExp{Pattern: pattern, Flags: flags}
}

func decodeDeepKind(kind string, value any, objectID runtime.RemoteObjectId) (any, error) {
	switch kind {
	case "boolean", "string", "number":
		return value, nil
	case "undefined", "null":
		return nil, nil
	case "bigint":
		return deepUnserializeBigInt(value), nil
	case "regexp":
		return deepUnserializeRegExp(value), nil
	case "date":
		return deepUnserializeDate(value), nil
	case "symbol":
		return deepUnserializeValue(value), nil
	case "array":
		return deepUnserializeArray(value), nil
	case "object":
		return deepUnserializeObject(value), nil
	case "map":
		return deepUnserializeMap(value), nil
	case "set":
		return deepUnserializeSet(value), nil
	case "error":
		return deepUnserializeObject(value), nil
	case "typedarray", "arraybuffer":
		return deepUnserializeArray(value), nil
	case "promise", "function", "weakmap", "weakset", "proxy", "window", "generator":
		if objectID != "" {
			return objectHandle(objectID), nil
		}
		return deepUnserializeValue(value), nil
	default:
		return deepUnserializeValue(value), nil
	}
}

// implemented deep-serialized kinds
// undefined, null, string, number, boolean, bigint, regexp, date, symbol,
// array, object, function, map, set, weakmap, weakset, error, proxy,
// promise, typedarray, arraybuffer, node, window, generator, nodelist
func (f *Frame) unserialize(value *runtime.RemoteObject) (any, error) {
	if value == nil {
		return nil, errors.New("cannot unserialize nil remote object")
	}
	dsv := value.DeepSerializedValue
	if dsv == nil {
		return value.Value, nil
	}

	switch dsv.Type {
	case "node":
		nodeTypeValue, ok := getNodeType(dsv.Value)
		if !ok || value.ObjectId == "" {
			return deepUnserializeValue(dsv.Value), nil
		}
		switch nodeTypeValue {
		case nodeTypeElement, nodeTypeDocument:
			return &Node{
				object: objectHandle(value.ObjectId),
				frame:  f,
			}, nil
		default:
			return deepUnserializeValue(dsv.Value), nil
		}

	case "nodelist":
		if value.Description == "NodeList(0)" {
			return nil, nil
		}
		if value.ObjectId != "" {
			return f.requestNodeList(value.ObjectId)
		}
		return deepUnserializeValue(dsv.Value), nil

	default:
		return decodeDeepKind(dsv.Type, dsv.Value, value.ObjectId)
	}
}

func (f *Frame) requestNodeList(objectId runtime.RemoteObjectId) (NodeList, error) {
	descriptor, err := f.getProperties(objectHandle(objectId), true, false, false, false)
	if err != nil {
		return nil, err
	}

	type nodeListEntry struct {
		index int
		node  *Node
	}

	entries := make([]nodeListEntry, 0, len(descriptor.Result))
	for _, d := range descriptor.Result {
		if !d.Enumerable || d.Value == nil || d.Value.ObjectId == "" {
			continue
		}
		index, convErr := strconv.Atoi(d.Name)
		if convErr != nil {
			continue
		}
		entries = append(entries, nodeListEntry{
			index: index,
			node: &Node{
				object:            objectHandle(d.Value.ObjectId),
				requestedSelector: fmt.Sprintf("NodeList[%d]", index),
				frame:             f,
			},
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].index < entries[j].index
	})

	nodeList := make(NodeList, 0, len(entries))
	for _, entry := range entries {
		nodeList = append(nodeList, entry.node)
	}
	return nodeList, nil
}

func (f Frame) evaluate(expression string, awaitPromise bool) (any, error) {
	var uid = f.executionContextID()
	if uid == "" {
		return nil, ErrExecutionContextDestroyed
	}
	value, err := runtime.Evaluate(f, runtime.EvaluateArgs{
		Expression:            expression,
		IncludeCommandLineAPI: true,
		UniqueContextId:       uid,
		AwaitPromise:          awaitPromise,
		SerializationOptions: &runtime.SerializationOptions{
			Serialization: "deep",
		},
	})
	if err != nil {
		return nil, err
	}
	if hasException(value.ExceptionDetails) {
		return nil, RuntimeExceptionError{value: value.ExceptionDetails}
	}
	return f.unserialize(value.Result)
}

func (f Frame) AwaitPromise(promise ObjectHandle) (any, error) {
	value, err := runtime.AwaitPromise(f, runtime.AwaitPromiseArgs{
		PromiseObjectId: promise.GetRemoteObjectID(),
		ReturnByValue:   true,
		GeneratePreview: false,
	})
	if err != nil {
		return nil, err
	}
	if hasException(value.ExceptionDetails) {
		return nil, RuntimeExceptionError{value: value.ExceptionDetails}
	}
	return f.unserialize(value.Result)
}

func toCallArgument(arg any) *runtime.CallArgument {
	callArg := &runtime.CallArgument{}
	switch a := arg.(type) {
	case ObjectHandle:
		callArg.ObjectId = a.GetRemoteObjectID()
	case runtime.RemoteObjectId:
		callArg.ObjectId = a
	default:
		callArg.Value = a
	}
	return callArg
}

func (f Frame) CallFunctionOn(self ObjectHandle, function string, awaitPromise bool, args ...any) (any, error) {
	arguments := make([]*runtime.CallArgument, 0, len(args))
	for _, arg := range args {
		arguments = append(arguments, toCallArgument(arg))
	}
	if len(arguments) == 0 {
		arguments = nil
	}
	value, err := runtime.CallFunctionOn(f, runtime.CallFunctionOnArgs{
		FunctionDeclaration: function,
		ObjectId:            self.GetRemoteObjectID(),
		AwaitPromise:        awaitPromise,
		Arguments:           arguments,
		SerializationOptions: &runtime.SerializationOptions{
			Serialization: "deep",
		},
	})
	if err != nil {
		return nil, err
	}
	if hasException(value.ExceptionDetails) {
		return nil, RuntimeExceptionError{value: value.ExceptionDetails}
	}
	return f.unserialize(value.Result)
}

func (f Frame) getProperties(self ObjectHandle, ownProperties, accessorPropertiesOnly, generatePreview, nonIndexedPropertiesOnly bool) (*runtime.GetPropertiesVal, error) {
	value, err := runtime.GetProperties(f, runtime.GetPropertiesArgs{
		ObjectId:                 self.GetRemoteObjectID(),
		OwnProperties:            ownProperties,
		AccessorPropertiesOnly:   accessorPropertiesOnly,
		GeneratePreview:          generatePreview,
		NonIndexedPropertiesOnly: nonIndexedPropertiesOnly,
	})
	if err != nil {
		return nil, err
	}
	if hasException(value.ExceptionDetails) {
		return nil, RuntimeExceptionError{value: value.ExceptionDetails}
	}
	return value, nil
}

func (f Frame) describeNode(self ObjectHandle) (*dom.Node, error) {
	value, err := dom.DescribeNode(f, dom.DescribeNodeArgs{
		ObjectId: self.GetRemoteObjectID(),
	})
	if err != nil {
		return nil, err
	}
	return value.Node, nil
}

func hasException(value *runtime.ExceptionDetails) bool {
	if value == nil {
		return false
	}
	if value.Exception != nil {
		return true
	}
	return false
}
