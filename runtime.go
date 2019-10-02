package witness

import "github.com/ecwid/witness/pkg/devtool"

func (session *Session) releaseObject(objectID string) error {
	_, err := session.blockingSend("Runtime.releaseObject", Map{"objectId": objectID})
	return err
}

func evaluateResult(msg bytes) (*devtool.RemoteObject, error) {
	result := new(devtool.EvaluatesResult)
	if err := msg.Unmarshal(result); err != nil {
		return nil, err
	}
	if result.ExceptionDetails != nil {
		return nil, result.ExceptionDetails
	}
	return result.Result, nil
}

// Evaluate Evaluates expression on global object.
func (session *Session) evaluate(expression string, contextID int64, async bool) (*devtool.RemoteObject, error) {
	exp := &devtool.EvaluatesExpression{
		Expression:    expression,
		ContextID:     contextID,
		AwaitPromise:  !async,
		ReturnByValue: false,
	}
	msg, err := session.blockingSend("Runtime.evaluate", exp)
	if err != nil {
		return nil, err
	}
	return evaluateResult(msg)
}

func (session *Session) getProperties(objectID string) ([]*devtool.PropertyDescriptor, error) {
	msg, err := session.blockingSend("Runtime.getProperties", Map{
		"objectId":               objectID,
		"ownProperties":          true,
		"accessorPropertiesOnly": false,
	})
	if err != nil {
		return nil, err
	}
	result := new(devtool.PropertiesResult)
	if err = msg.Unmarshal(result); err != nil {
		return nil, err
	}
	if result.ExceptionDetails != nil {
		return nil, result.ExceptionDetails
	}
	return result.Result, nil
}

func (session *Session) callFunctionOn(objectID string, functionDeclaration string, arg ...interface{}) (*devtool.RemoteObject, error) {
	args := make([]devtool.CallArgument, len(arg))
	for i, a := range arg {
		args[i] = devtool.CallArgument{Value: a}
	}
	msg, err := session.blockingSend("Runtime.callFunctionOn", Map{
		"functionDeclaration": functionDeclaration,
		"objectId":            objectID,
		"arguments":           args,
		"awaitPromise":        true,
		"returnByValue":       false,
	})
	if err != nil {
		return nil, err
	}
	return evaluateResult(msg)
}
