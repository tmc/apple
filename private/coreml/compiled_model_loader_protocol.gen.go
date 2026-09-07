// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLCompiledModelLoader protocol.
type MLCompiledModelLoader interface {
	objectivec.IObject

	// LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError protocol.
	LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive MLModelInputArchiverRef, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error)
}

// MLCompiledModelLoaderObject wraps an existing Objective-C object that conforms to the MLCompiledModelLoader protocol.
type MLCompiledModelLoaderObject struct {
	objectivec.Object
}

func (o MLCompiledModelLoaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLCompiledModelLoaderObjectFromID constructs a [MLCompiledModelLoaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLCompiledModelLoaderObjectFromID(id objc.ID) MLCompiledModelLoaderObject {
	return MLCompiledModelLoaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLCompiledModelLoaderObject) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive MLModelInputArchiverRef, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
