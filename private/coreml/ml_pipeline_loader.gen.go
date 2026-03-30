// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPipelineLoader] class.
var (
	_MLPipelineLoaderClass     MLPipelineLoaderClass
	_MLPipelineLoaderClassOnce sync.Once
)

func getMLPipelineLoaderClass() MLPipelineLoaderClass {
	_MLPipelineLoaderClassOnce.Do(func() {
		_MLPipelineLoaderClass = MLPipelineLoaderClass{class: objc.GetClass("MLPipelineLoader")}
	})
	return _MLPipelineLoaderClass
}

// GetMLPipelineLoaderClass returns the class object for MLPipelineLoader.
func GetMLPipelineLoaderClass() MLPipelineLoaderClass {
	return getMLPipelineLoaderClass()
}

type MLPipelineLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineLoaderClass) Alloc() MLPipelineLoader {
	rv := objc.Send[MLPipelineLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPipelineLoader.DebugDescription]
//   - [MLPipelineLoader.Description]
//   - [MLPipelineLoader.Hash]
//   - [MLPipelineLoader.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader
type MLPipelineLoader struct {
	objectivec.Object
}

// MLPipelineLoaderFromID constructs a [MLPipelineLoader] from an objc.ID.
func MLPipelineLoaderFromID(id objc.ID) MLPipelineLoader {
	return MLPipelineLoader{objectivec.Object{ID: id}}
}

// Ensure MLPipelineLoader implements IMLPipelineLoader.
var _ IMLPipelineLoader = MLPipelineLoader{}

// An interface definition for the [MLPipelineLoader] class.
//
// # Methods
//
//   - [IMLPipelineLoader.DebugDescription]
//   - [IMLPipelineLoader.Description]
//   - [IMLPipelineLoader.Hash]
//   - [IMLPipelineLoader.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader
type IMLPipelineLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (p MLPipelineLoader) Init() MLPipelineLoader {
	rv := objc.Send[MLPipelineLoader](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPipelineLoader) Autorelease() MLPipelineLoader {
	rv := objc.Send[MLPipelineLoader](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineLoader creates a new MLPipelineLoader instance.
func NewMLPipelineLoader() MLPipelineLoader {
	class := getMLPipelineLoaderClass()
	rv := objc.Send[MLPipelineLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/ensureFrameworkSupportsCompilerVersion:error:
func (_MLPipelineLoaderClass MLPipelineLoaderClass) EnsureFrameworkSupportsCompilerVersionError(version objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLPipelineLoaderClass.class), objc.Sel("ensureFrameworkSupportsCompilerVersion:error:"), version, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("ensureFrameworkSupportsCompilerVersion:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/loadModelAssetDescriptionFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLPipelineLoaderClass MLPipelineLoaderClass) LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineLoaderClass.class), objc.Sel("loadModelAssetDescriptionFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLPipelineLoaderClass MLPipelineLoaderClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineLoaderClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/debugDescription
func (p MLPipelineLoader) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/description
func (p MLPipelineLoader) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/hash
func (p MLPipelineLoader) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineLoader/superclass
func (p MLPipelineLoader) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}
