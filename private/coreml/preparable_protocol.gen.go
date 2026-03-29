// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLPreparable protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLPreparable
type MLPreparable interface {
	objectivec.IObject

	// PrepareWithConcurrencyHintError protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLPreparable/prepareWithConcurrencyHint:error:
	PrepareWithConcurrencyHintError(hint int64) (bool, error)
}

// MLPreparableObject wraps an existing Objective-C object that conforms to the MLPreparable protocol.
type MLPreparableObject struct {
	objectivec.Object
}
func (o MLPreparableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLPreparableObjectFromID constructs a [MLPreparableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLPreparableObjectFromID(id objc.ID) MLPreparableObject {
	return MLPreparableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPreparable/prepareWithConcurrencyHint:error:
func (o MLPreparableObject) PrepareWithConcurrencyHintError(hint int64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("prepareWithConcurrencyHint:error:"), hint)
	if err != nil {
		return false, err
	}
	return rv, nil
	}

