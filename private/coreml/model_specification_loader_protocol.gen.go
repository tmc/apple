// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModelSpecificationLoader protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelSpecificationLoader
type MLModelSpecificationLoader interface {
	objectivec.IObject
}

// MLModelSpecificationLoaderObject wraps an existing Objective-C object that conforms to the MLModelSpecificationLoader protocol.
type MLModelSpecificationLoaderObject struct {
	objectivec.Object
}

func (o MLModelSpecificationLoaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLModelSpecificationLoaderObjectFromID constructs a [MLModelSpecificationLoaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLModelSpecificationLoaderObjectFromID(id objc.ID) MLModelSpecificationLoaderObject {
	return MLModelSpecificationLoaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelSpecificationLoader/loadModelFromSpecification:configuration:error:
func (o MLModelSpecificationLoaderObject) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
