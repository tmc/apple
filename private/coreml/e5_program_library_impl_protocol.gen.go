// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5ProgramLibraryImpl protocol.
type MLE5ProgramLibraryImpl interface {
	objectivec.IObject

	// CreateProgramLibraryHandleWithRespecializationError protocol.
	CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error)

	// ModelDisplayName protocol.
	ModelDisplayName() objectivec.IObject

	// SerializedMILText protocol.
	SerializedMILText() objectivec.IObject
}

// MLE5ProgramLibraryImplObject wraps an existing Objective-C object that conforms to the MLE5ProgramLibraryImpl protocol.
type MLE5ProgramLibraryImplObject struct {
	objectivec.Object
}

func (o MLE5ProgramLibraryImplObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLE5ProgramLibraryImplObjectFromID constructs a [MLE5ProgramLibraryImplObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLE5ProgramLibraryImplObjectFromID(id objc.ID) MLE5ProgramLibraryImplObject {
	return MLE5ProgramLibraryImplObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLE5ProgramLibraryImplObject) CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("createProgramLibraryHandleWithRespecialization:error:"), respecialization)
	if err != nil {
		return *new(E5rtProgramLibraryRef), err
	}
	return E5rtProgramLibraryRef(rv), nil
}
func (o MLE5ProgramLibraryImplObject) ModelDisplayName() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("modelDisplayName"))
	return objectivec.Object{ID: rv}
}
func (o MLE5ProgramLibraryImplObject) SerializedMILText() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("serializedMILText"))
	return objectivec.Object{ID: rv}
}
