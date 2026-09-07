// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModelAssetDescriptionLoader protocol.
type MLModelAssetDescriptionLoader interface {
	objectivec.IObject

	// LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError protocol.
	LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error)
}

// MLModelAssetDescriptionLoaderObject wraps an existing Objective-C object that conforms to the MLModelAssetDescriptionLoader protocol.
type MLModelAssetDescriptionLoaderObject struct {
	objectivec.Object
}

func (o MLModelAssetDescriptionLoaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLModelAssetDescriptionLoaderObjectFromID constructs a [MLModelAssetDescriptionLoaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLModelAssetDescriptionLoaderObjectFromID(id objc.ID) MLModelAssetDescriptionLoaderObject {
	return MLModelAssetDescriptionLoaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLModelAssetDescriptionLoaderObject) LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("loadModelAssetDescriptionFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
