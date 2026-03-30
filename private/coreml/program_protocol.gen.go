// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLProgram protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLProgram
type MLProgram interface {
	objectivec.IObject
}

// MLProgramObject wraps an existing Objective-C object that conforms to the MLProgram protocol.
type MLProgramObject struct {
	objectivec.Object
}

func (o MLProgramObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLProgramObjectFromID constructs a [MLProgramObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLProgramObjectFromID(id objc.ID) MLProgramObject {
	return MLProgramObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgram/serializedMILText
func (o MLProgramObject) SerializedMILText() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("serializedMILText"))
	return objectivec.Object{ID: rv}
}
