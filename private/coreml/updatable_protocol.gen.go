// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLUpdatable protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdatable
type MLUpdatable interface {
	objectivec.IObject

	// CancelUpdate protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/cancelUpdate
	CancelUpdate()

	// ResumeUpdate protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/resumeUpdate
	ResumeUpdate()
}

// MLUpdatableObject wraps an existing Objective-C object that conforms to the MLUpdatable protocol.
type MLUpdatableObject struct {
	objectivec.Object
}

func (o MLUpdatableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLUpdatableObjectFromID constructs a [MLUpdatableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLUpdatableObjectFromID(id objc.ID) MLUpdatableObject {
	return MLUpdatableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/cancelUpdate
func (o MLUpdatableObject) CancelUpdate() {
	objc.Send[struct{}](o.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (o MLUpdatableObject) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/resumeUpdate
func (o MLUpdatableObject) ResumeUpdate() {
	objc.Send[struct{}](o.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/resumeUpdateWithParameters:
func (o MLUpdatableObject) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/setUpdateProgressHandlers:dispatchQueue:
func (o MLUpdatableObject) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdatable/updateModelWithData:
func (o MLUpdatableObject) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateModelWithData:"), data)
}
