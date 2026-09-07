// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLProgram protocol.
type MLProgram interface {
	objectivec.IObject

	// SerializedMILText protocol.
	SerializedMILText() objectivec.IObject
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

func (o MLProgramObject) SerializedMILText() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("serializedMILText"))
	return objectivec.Object{ID: rv}
}
