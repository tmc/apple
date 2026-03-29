// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModelSpecificationSaver protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelSpecificationSaver
type MLModelSpecificationSaver interface {
	objectivec.IObject

	// SaveModelToSpecification protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelSpecificationSaver/saveModelToSpecification:
	SaveModelToSpecification(specification []objectivec.IObject) unsafe.Pointer
}

// MLModelSpecificationSaverObject wraps an existing Objective-C object that conforms to the MLModelSpecificationSaver protocol.
type MLModelSpecificationSaverObject struct {
	objectivec.Object
}
func (o MLModelSpecificationSaverObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLModelSpecificationSaverObjectFromID constructs a [MLModelSpecificationSaverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLModelSpecificationSaverObjectFromID(id objc.ID) MLModelSpecificationSaverObject {
	return MLModelSpecificationSaverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelSpecificationSaver/saveModelToSpecification:
func (o MLModelSpecificationSaverObject) SaveModelToSpecification(specification []objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("saveModelToSpecification:"), objectivec.IObjectSliceToNSArray(specification))
	return rv
	}

