// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLCustomLayerFactory protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerFactory
type MLCustomLayerFactory interface {
	objectivec.IObject
}

// MLCustomLayerFactoryObject wraps an existing Objective-C object that conforms to the MLCustomLayerFactory protocol.
type MLCustomLayerFactoryObject struct {
	objectivec.Object
}

func (o MLCustomLayerFactoryObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLCustomLayerFactoryObjectFromID constructs a [MLCustomLayerFactoryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLCustomLayerFactoryObjectFromID(id objc.ID) MLCustomLayerFactoryObject {
	return MLCustomLayerFactoryObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerFactory/createCustomLayer:withParameters:error:
func (o MLCustomLayerFactoryObject) CreateCustomLayerWithParametersError(layer objectivec.IObject, parameters objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("createCustomLayer:withParameters:error:"), layer, parameters)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
