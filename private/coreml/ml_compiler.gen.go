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

// The class instance for the [MLCompiler] class.
var (
	_MLCompilerClass     MLCompilerClass
	_MLCompilerClassOnce sync.Once
)

func getMLCompilerClass() MLCompilerClass {
	_MLCompilerClassOnce.Do(func() {
		_MLCompilerClass = MLCompilerClass{class: objc.GetClass("MLCompiler")}
	})
	return _MLCompilerClass
}

// GetMLCompilerClass returns the class object for MLCompiler.
func GetMLCompilerClass() MLCompilerClass {
	return getMLCompilerClass()
}

type MLCompilerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCompilerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCompilerClass) Alloc() MLCompiler {
	rv := objc.Send[MLCompiler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCompiler
type MLCompiler struct {
	objectivec.Object
}

// MLCompilerFromID constructs a [MLCompiler] from an objc.ID.
func MLCompilerFromID(id objc.ID) MLCompiler {
	return MLCompiler{objectivec.Object{ID: id}}
}
// Ensure MLCompiler implements IMLCompiler.
var _ IMLCompiler = MLCompiler{}

// An interface definition for the [MLCompiler] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler
type IMLCompiler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c MLCompiler) Init() MLCompiler {
	rv := objc.Send[MLCompiler](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCompiler) Autorelease() MLCompiler {
	rv := objc.Send[MLCompiler](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompiler creates a new MLCompiler instance.
func NewMLCompiler() MLCompiler {
	class := getMLCompilerClass()
	rv := objc.Send[MLCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_compileSpecification:blobMapping:toArchive:options:error:
func (_MLCompilerClass MLCompilerClass) _compileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("_compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_compileSpecificationAtURL:toURL:compiledModelName:overridingModelDescription:options:error:
func (_MLCompilerClass MLCompilerClass) _compileSpecificationAtURLToURLCompiledModelNameOverridingModelDescriptionOptionsError(url foundation.INSURL, url2 foundation.INSURL, name objectivec.IObject, description objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("_compileSpecificationAtURL:toURL:compiledModelName:overridingModelDescription:options:error:"), url, url2, name, description, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// CompileSpecificationAtURLToURLCompiledModelNameOverridingModelDescriptionOptionsError is an exported wrapper for the private method _compileSpecificationAtURLToURLCompiledModelNameOverridingModelDescriptionOptionsError.
func (_MLCompilerClass MLCompilerClass) CompileSpecificationAtURLToURLCompiledModelNameOverridingModelDescriptionOptionsError(url foundation.INSURL, url2 foundation.INSURL, name objectivec.IObject, description objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	return _MLCompilerClass._compileSpecificationAtURLToURLCompiledModelNameOverridingModelDescriptionOptionsError(url, url2, name, description, options)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_loadSpecificationAtURL:to:error:
func (_MLCompilerClass MLCompilerClass) _loadSpecificationAtURLToError(url foundation.INSURL, to unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("_loadSpecificationAtURL:to:error:"), url, to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_loadSpecificationAtURL:to:error: returned NO with nil NSError")
	}
	return rv, nil

}

// LoadSpecificationAtURLToError is an exported wrapper for the private method _loadSpecificationAtURLToError.
func (_MLCompilerClass MLCompilerClass) LoadSpecificationAtURLToError(url foundation.INSURL, to unsafe.Pointer) (bool, error) {
	return _MLCompilerClass._loadSpecificationAtURLToError(url, to)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_updateFeatures:withFeatures:
func (_MLCompilerClass MLCompilerClass) _updateFeaturesWithFeatures(features unsafe.Pointer, features2 objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("_updateFeatures:withFeatures:"), features, features2)
}

// UpdateFeaturesWithFeatures is an exported wrapper for the private method _updateFeaturesWithFeatures.
func (_MLCompilerClass MLCompilerClass) UpdateFeaturesWithFeatures(features unsafe.Pointer, features2 objectivec.IObject) {
	_MLCompilerClass._updateFeaturesWithFeatures(features, features2)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_updateMetadata:withMetadata:
func (_MLCompilerClass MLCompilerClass) _updateMetadataWithMetadata(metadata unsafe.Pointer, metadata2 objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("_updateMetadata:withMetadata:"), metadata, metadata2)
}

// UpdateMetadataWithMetadata is an exported wrapper for the private method _updateMetadataWithMetadata.
func (_MLCompilerClass MLCompilerClass) UpdateMetadataWithMetadata(metadata unsafe.Pointer, metadata2 objectivec.IObject) {
	_MLCompilerClass._updateMetadataWithMetadata(metadata, metadata2)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/_updateSpecification:withModelDescription:
func (_MLCompilerClass MLCompilerClass) _updateSpecificationWithModelDescription(specification unsafe.Pointer, description objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("_updateSpecification:withModelDescription:"), specification, description)
}

// UpdateSpecificationWithModelDescription is an exported wrapper for the private method _updateSpecificationWithModelDescription.
func (_MLCompilerClass MLCompilerClass) UpdateSpecificationWithModelDescription(specification unsafe.Pointer, description objectivec.IObject) {
	_MLCompilerClass._updateSpecificationWithModelDescription(specification, description)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/addMLProgramToCompiledModelAtURL:error:
func (_MLCompilerClass MLCompilerClass) AddMLProgramToCompiledModelAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("addMLProgramToCompiledModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/addMLProgramToCompiledModelAtURL:withCompilationMode:dryRun:oarchiveForModelCompilation:compilerEvent:error:
func (_MLCompilerClass MLCompilerClass) AddMLProgramToCompiledModelAtURLWithCompilationModeDryRunOarchiveForModelCompilationCompilerEventError(url foundation.INSURL, mode int, run bool, compilation unsafe.Pointer, event objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("addMLProgramToCompiledModelAtURL:withCompilationMode:dryRun:oarchiveForModelCompilation:compilerEvent:error:"), url, mode, run, compilation, event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/canAddMLProgramToCompiledModelAtURL:
func (_MLCompilerClass MLCompilerClass) CanAddMLProgramToCompiledModelAtURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("canAddMLProgramToCompiledModelAtURL:"), url)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileModelAtURL:toURL:options:error:
func (_MLCompilerClass MLCompilerClass) CompileModelAtURLToURLOptionsError(url foundation.INSURL, url2 foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileModelAtURL:toURL:options:error:"), url, url2, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileSpecification:blobMapping:compilerClass:toArchive:options:error:
func (_MLCompilerClass MLCompilerClass) CompileSpecificationBlobMappingCompilerClassToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, class objc.Class, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileSpecification:blobMapping:compilerClass:toArchive:options:error:"), specification, mapping, class, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileSpecification:blobMapping:toArchive:options:compilerEvent:error:
func (_MLCompilerClass MLCompilerClass) CompileSpecificationBlobMappingToArchiveOptionsCompilerEventError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject, event objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileSpecification:blobMapping:toArchive:options:compilerEvent:error:"), specification, mapping, archive, options, event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileSpecification:blobMapping:toArchive:options:error:
func (_MLCompilerClass MLCompilerClass) CompileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileSpecification:toArchive:options:error:
func (_MLCompilerClass MLCompilerClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compileSpecificationAtURL:toURL:options:error:
func (_MLCompilerClass MLCompilerClass) CompileSpecificationAtURLToURLOptionsError(url foundation.INSURL, url2 foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compileSpecificationAtURL:toURL:options:error:"), url, url2, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compiledVersionForModelAtURL:options:error:
func (_MLCompilerClass MLCompilerClass) CompiledVersionForModelAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compiledVersionForModelAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compiledVersionForSpecification:options:error:
func (_MLCompilerClass MLCompilerClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/compiledVersionForSpecificationAtURL:options:error:
func (_MLCompilerClass MLCompilerClass) CompiledVersionForSpecificationAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("compiledVersionForSpecificationAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/contentsOfDirectoryAtURL:error:
func (_MLCompilerClass MLCompilerClass) ContentsOfDirectoryAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("contentsOfDirectoryAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/encryptCompiledModelAtURL:options:error:
func (_MLCompilerClass MLCompilerClass) EncryptCompiledModelAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("encryptCompiledModelAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("encryptCompiledModelAtURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/encryptFileAtURL:options:error:
func (_MLCompilerClass MLCompilerClass) EncryptFileAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("encryptFileAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("encryptFileAtURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/fillCompilerEvent:withMetadataFromModelAt:error:
func (_MLCompilerClass MLCompilerClass) FillCompilerEventWithMetadataFromModelAtError(event objectivec.IObject, at objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("fillCompilerEvent:withMetadataFromModelAt:error:"), event, at, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("fillCompilerEvent:withMetadataFromModelAt:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/fingerprintSpecificationAtURL:toArchive:hash:error:
func (_MLCompilerClass MLCompilerClass) FingerprintSpecificationAtURLToArchiveHashError(url foundation.INSURL, archive unsafe.Pointer, hash objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("fingerprintSpecificationAtURL:toArchive:hash:error:"), url, archive, hash, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("fingerprintSpecificationAtURL:toArchive:hash:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/hashSpecificationAtURL:
func (_MLCompilerClass MLCompilerClass) HashSpecificationAtURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("hashSpecificationAtURL:"), url)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/storeEncryptionInfoInCompiledArchive:options:error:
func (_MLCompilerClass MLCompilerClass) StoreEncryptionInfoInCompiledArchiveOptionsError(archive unsafe.Pointer, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLCompilerClass.class), objc.Sel("storeEncryptionInfoInCompiledArchive:options:error:"), archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("storeEncryptionInfoInCompiledArchive:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLCompiler/versionInfo
func (_MLCompilerClass MLCompilerClass) VersionInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerClass.class), objc.Sel("versionInfo"))
	return objectivec.Object{ID: rv}
}

