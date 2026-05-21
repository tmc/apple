// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ETTaskContext protocol.
type ETTaskContext interface {
	objectivec.IObject
}

// ETTaskContextObject wraps an existing Objective-C object that conforms to the ETTaskContext protocol.
type ETTaskContextObject struct {
	objectivec.Object
}

func (o ETTaskContextObject) BaseObject() objectivec.Object {
	return o.Object
}

// ETTaskContextObjectFromID constructs a [ETTaskContextObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ETTaskContextObjectFromID(id objc.ID) ETTaskContextObject {
	return ETTaskContextObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ETTaskContextObject) DoInferenceOnDataError(data objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("doInferenceOnData:error:"), data)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o ETTaskContextObject) GetTensorNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("getTensorNamed:"), named)
	return objectivec.Object{ID: rv}
}
func (o ETTaskContextObject) SaveNetworkInplaceError(network objectivec.IObject, inplace bool) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("saveNetwork:inplace:error:"), network, inplace)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (o ETTaskContextObject) SetTensorNamedWithValueError(named objectivec.IObject, value objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setTensorNamed:withValue:error:"), named, value)
	if err != nil {
		return false, err
	}
	return rv, nil
}
