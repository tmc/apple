// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLUpdatable protocol.
type MLUpdatable interface {
	objectivec.IObject

	// CancelUpdate protocol.
	CancelUpdate()

	// LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError protocol.
	LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error)

	// ResumeUpdate protocol.
	ResumeUpdate()

	// ResumeUpdateWithParameters protocol.
	ResumeUpdateWithParameters(parameters objectivec.IObject)

	// SetUpdateProgressHandlersDispatchQueue protocol.
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)

	// UpdateModelWithData protocol.
	UpdateModelWithData(data objectivec.IObject)
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

func (o MLUpdatableObject) CancelUpdate() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cancelUpdate"))
}
func (o MLUpdatableObject) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLUpdatableObject) ResumeUpdate() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("resumeUpdate"))
}
func (o MLUpdatableObject) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
func (o MLUpdatableObject) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (o MLUpdatableObject) UpdateModelWithData(data objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("updateModelWithData:"), data)
}
