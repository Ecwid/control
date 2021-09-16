package control

import (
	"github.com/ecwid/control/protocol/runtime"
)

type primitiveRemoteObject runtime.RemoteObject

func (p primitiveRemoteObject) String() (string, error) {
	const to = "string"
	if p.Type == to {
		return p.Value.(string), nil
	}
	return "", RemoteObjectCastError{
		object: p,
		cast:   to,
	}
}

// Bool RemoteObject as bool value
func (p primitiveRemoteObject) Bool() (bool, error) {
	const to = "boolean"
	if p.Type == to {
		return p.Value.(bool), nil
	}
	return false, RemoteObjectCastError{
		object: p,
		cast:   to,
	}
}

func (f Frame) getProperties(objectID runtime.RemoteObjectId, ownProperties, accessorPropertiesOnly bool) ([]*runtime.PropertyDescriptor, error) {
	val, err := runtime.GetProperties(f, runtime.GetPropertiesArgs{
		ObjectId:               objectID,
		OwnProperties:          ownProperties,
		AccessorPropertiesOnly: accessorPropertiesOnly,
	})
	if err != nil {
		return nil, err
	}
	if val.ExceptionDetails != nil {
		return nil, RuntimeError(*val.ExceptionDetails)
	}
	return val.Result, nil
}
