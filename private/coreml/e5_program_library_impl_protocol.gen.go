// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5ProgramLibraryImpl protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryImpl
type MLE5ProgramLibraryImpl interface {
	objectivec.IObject

	// CreateProgramLibraryHandleWithRespecializationError protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryImpl/createProgramLibraryHandleWithRespecialization:error:
	CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rt_program_libraryRef, error)
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

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryImpl/createProgramLibraryHandleWithRespecialization:error:
func (o MLE5ProgramLibraryImplObject) CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rt_program_libraryRef, error) {
	rv, err := objc.SendWithError[E5rt_program_libraryRef](o.ID, objc.Sel("createProgramLibraryHandleWithRespecialization:error:"), respecialization)
	if err != nil {
		return 0, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryImpl/modelDisplayName
func (o MLE5ProgramLibraryImplObject) ModelDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelDisplayName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryImpl/serializedMILText
func (o MLE5ProgramLibraryImplObject) SerializedMILText() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("serializedMILText"))
	return objectivec.Object{ID: rv}
}
