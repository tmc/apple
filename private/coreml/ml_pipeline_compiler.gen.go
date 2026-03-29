// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPipelineCompiler] class.
var (
	_MLPipelineCompilerClass     MLPipelineCompilerClass
	_MLPipelineCompilerClassOnce sync.Once
)

func getMLPipelineCompilerClass() MLPipelineCompilerClass {
	_MLPipelineCompilerClassOnce.Do(func() {
		_MLPipelineCompilerClass = MLPipelineCompilerClass{class: objc.GetClass("MLPipelineCompiler")}
	})
	return _MLPipelineCompilerClass
}

// GetMLPipelineCompilerClass returns the class object for MLPipelineCompiler.
func GetMLPipelineCompilerClass() MLPipelineCompilerClass {
	return getMLPipelineCompilerClass()
}

type MLPipelineCompilerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineCompilerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineCompilerClass) Alloc() MLPipelineCompiler {
	rv := objc.Send[MLPipelineCompiler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler
type MLPipelineCompiler struct {
	objectivec.Object
}

// MLPipelineCompilerFromID constructs a [MLPipelineCompiler] from an objc.ID.
func MLPipelineCompilerFromID(id objc.ID) MLPipelineCompiler {
	return MLPipelineCompiler{objectivec.Object{ID: id}}
}
// Ensure MLPipelineCompiler implements IMLPipelineCompiler.
var _ IMLPipelineCompiler = MLPipelineCompiler{}

// An interface definition for the [MLPipelineCompiler] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler
type IMLPipelineCompiler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p MLPipelineCompiler) Init() MLPipelineCompiler {
	rv := objc.Send[MLPipelineCompiler](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPipelineCompiler) Autorelease() MLPipelineCompiler {
	rv := objc.Send[MLPipelineCompiler](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineCompiler creates a new MLPipelineCompiler instance.
func NewMLPipelineCompiler() MLPipelineCompiler {
	class := getMLPipelineCompilerClass()
	rv := objc.Send[MLPipelineCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/_archiveCustomModelNames:to:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) _archiveCustomModelNamesTo(names unsafe.Pointer, to unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("_archiveCustomModelNames:to:"), names, to)
}

// ArchiveCustomModelNamesTo is an exported wrapper for the private method _archiveCustomModelNamesTo.
func (_MLPipelineCompilerClass MLPipelineCompilerClass) ArchiveCustomModelNamesTo(names unsafe.Pointer, to unsafe.Pointer) {
	_MLPipelineCompilerClass._archiveCustomModelNamesTo(names, to)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/_archivePipelineModelDetailsFrom:toArchive:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) _archivePipelineModelDetailsFromToArchiveError(from unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("_archivePipelineModelDetailsFrom:toArchive:error:"), from, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_archivePipelineModelDetailsFrom:toArchive:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ArchivePipelineModelDetailsFromToArchiveError is an exported wrapper for the private method _archivePipelineModelDetailsFromToArchiveError.
func (_MLPipelineCompilerClass MLPipelineCompilerClass) ArchivePipelineModelDetailsFromToArchiveError(from unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	return _MLPipelineCompilerClass._archivePipelineModelDetailsFromToArchiveError(from, archive)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/_archivePipelineUpdateParameterForModels:to:updatable:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) _archivePipelineUpdateParameterForModelsToUpdatable(models unsafe.Pointer, to unsafe.Pointer, updatable bool) {
	objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("_archivePipelineUpdateParameterForModels:to:updatable:"), models, to, updatable)
}

// ArchivePipelineUpdateParameterForModelsToUpdatable is an exported wrapper for the private method _archivePipelineUpdateParameterForModelsToUpdatable.
func (_MLPipelineCompilerClass MLPipelineCompilerClass) ArchivePipelineUpdateParameterForModelsToUpdatable(models unsafe.Pointer, to unsafe.Pointer, updatable bool) {
	_MLPipelineCompilerClass._archivePipelineUpdateParameterForModelsToUpdatable(models, to, updatable)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/_compileSpecification:blobMapping:toArchive:options:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) _compileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("_compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/_compileWithModelsInPipeline:blobMapping:toArchive:options:updatable:osSizeTracking:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) _compileWithModelsInPipelineBlobMappingToArchiveOptionsUpdatableOsSizeTrackingError(pipeline unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject, updatable bool, tracking objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("_compileWithModelsInPipeline:blobMapping:toArchive:options:updatable:osSizeTracking:error:"), pipeline, mapping, archive, options, updatable, tracking, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// CompileWithModelsInPipelineBlobMappingToArchiveOptionsUpdatableOsSizeTrackingError is an exported wrapper for the private method _compileWithModelsInPipelineBlobMappingToArchiveOptionsUpdatableOsSizeTrackingError.
func (_MLPipelineCompilerClass MLPipelineCompilerClass) CompileWithModelsInPipelineBlobMappingToArchiveOptionsUpdatableOsSizeTrackingError(pipeline unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject, updatable bool, tracking objectivec.IObject) (objectivec.IObject, error) {
	return _MLPipelineCompilerClass._compileWithModelsInPipelineBlobMappingToArchiveOptionsUpdatableOsSizeTrackingError(pipeline, mapping, archive, options, updatable, tracking)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/compileSpecification:blobMapping:toArchive:options:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) CompileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/compileSpecification:toArchive:options:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineCompiler/compiledVersionForSpecification:options:error:
func (_MLPipelineCompilerClass MLPipelineCompilerClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineCompilerClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

