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

// The class instance for the [MLLoader] class.
var (
	_MLLoaderClass     MLLoaderClass
	_MLLoaderClassOnce sync.Once
)

func getMLLoaderClass() MLLoaderClass {
	_MLLoaderClassOnce.Do(func() {
		_MLLoaderClass = MLLoaderClass{class: objc.GetClass("MLLoader")}
	})
	return _MLLoaderClass
}

// GetMLLoaderClass returns the class object for MLLoader.
func GetMLLoaderClass() MLLoaderClass {
	return getMLLoaderClass()
}

type MLLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLoaderClass) Alloc() MLLoader {
	rv := objc.Send[MLLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLoader
type MLLoader struct {
	objectivec.Object
}

// MLLoaderFromID constructs a [MLLoader] from an objc.ID.
func MLLoaderFromID(id objc.ID) MLLoader {
	return MLLoader{objectivec.Object{ID: id}}
}
// Ensure MLLoader implements IMLLoader.
var _ IMLLoader = MLLoader{}

// An interface definition for the [MLLoader] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader
type IMLLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l MLLoader) Init() MLLoader {
	rv := objc.Send[MLLoader](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLoader) Autorelease() MLLoader {
	rv := objc.Send[MLLoader](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLoader creates a new MLLoader instance.
func NewMLLoader() MLLoader {
	class := getMLLoaderClass()
	rv := objc.Send[MLLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_conformConfiguration:usingModelArchive:
func (_MLLoaderClass MLLoaderClass) _conformConfigurationUsingModelArchive(configuration objectivec.IObject, archive unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_conformConfiguration:usingModelArchive:"), configuration, archive)
	return objectivec.Object{ID: rv}
}

// ConformConfigurationUsingModelArchive is an exported wrapper for the private method _conformConfigurationUsingModelArchive.
func (_MLLoaderClass MLLoaderClass) ConformConfigurationUsingModelArchive(configuration objectivec.IObject, archive unsafe.Pointer) objectivec.IObject {
	return _MLLoaderClass._conformConfigurationUsingModelArchive(configuration, archive)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_createDecryptSessionForModelAtURL:configuration:decryptSession:loaderEvent:error:
func (_MLLoaderClass MLLoaderClass) _createDecryptSessionForModelAtURLConfigurationDecryptSessionLoaderEventError(url foundation.INSURL, configuration objectivec.IObject, session []objectivec.IObject, event objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLLoaderClass.class), objc.Sel("_createDecryptSessionForModelAtURL:configuration:decryptSession:loaderEvent:error:"), url, configuration, objectivec.IObjectSliceToNSArray(session), event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_createDecryptSessionForModelAtURL:configuration:decryptSession:loaderEvent:error: returned NO with nil NSError")
	}
	return rv, nil

}

// CreateDecryptSessionForModelAtURLConfigurationDecryptSessionLoaderEventError is an exported wrapper for the private method _createDecryptSessionForModelAtURLConfigurationDecryptSessionLoaderEventError.
func (_MLLoaderClass MLLoaderClass) CreateDecryptSessionForModelAtURLConfigurationDecryptSessionLoaderEventError(url foundation.INSURL, configuration objectivec.IObject, session []objectivec.IObject, event objectivec.IObject) (bool, error) {
	return _MLLoaderClass._createDecryptSessionForModelAtURLConfigurationDecryptSessionLoaderEventError(url, configuration, session, event)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_findCodedObjectURLInModelArchive:
func (_MLLoaderClass MLLoaderClass) _findCodedObjectURLInModelArchive(archive unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_findCodedObjectURLInModelArchive:"), archive)
	return objectivec.Object{ID: rv}
}

// FindCodedObjectURLInModelArchive is an exported wrapper for the private method _findCodedObjectURLInModelArchive.
func (_MLLoaderClass MLLoaderClass) FindCodedObjectURLInModelArchive(archive unsafe.Pointer) objectivec.IObject {
	return _MLLoaderClass._findCodedObjectURLInModelArchive(archive)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadModelAssetDescriptionFromArchive:configuration:modelVersion:compilerVersion:loadingClasses:error:
func (_MLLoaderClass MLLoaderClass) _loadModelAssetDescriptionFromArchiveConfigurationModelVersionCompilerVersionLoadingClassesError(archive unsafe.Pointer, configuration objectivec.IObject, version objectivec.IObject, version2 objectivec.IObject, classes objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadModelAssetDescriptionFromArchive:configuration:modelVersion:compilerVersion:loadingClasses:error:"), archive, configuration, version, version2, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadModelAssetDescriptionFromArchiveConfigurationModelVersionCompilerVersionLoadingClassesError is an exported wrapper for the private method _loadModelAssetDescriptionFromArchiveConfigurationModelVersionCompilerVersionLoadingClassesError.
func (_MLLoaderClass MLLoaderClass) LoadModelAssetDescriptionFromArchiveConfigurationModelVersionCompilerVersionLoadingClassesError(archive unsafe.Pointer, configuration objectivec.IObject, version objectivec.IObject, version2 objectivec.IObject, classes objectivec.IObject) (objectivec.IObject, error) {
	return _MLLoaderClass._loadModelAssetDescriptionFromArchiveConfigurationModelVersionCompilerVersionLoadingClassesError(archive, configuration, version, version2, classes)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadModelFromArchive:configuration:loaderEvent:useUpdatableModelLoaders:error:
func (_MLLoaderClass MLLoaderClass) _loadModelFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive unsafe.Pointer, configuration objectivec.IObject, event objectivec.IObject, loaders bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadModelFromArchive:configuration:loaderEvent:useUpdatableModelLoaders:error:"), archive, configuration, event, loaders, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadModelFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError is an exported wrapper for the private method _loadModelFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError.
func (_MLLoaderClass MLLoaderClass) LoadModelFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive unsafe.Pointer, configuration objectivec.IObject, event objectivec.IObject, loaders bool) (objectivec.IObject, error) {
	return _MLLoaderClass._loadModelFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive, configuration, event, loaders)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadModelFromArchive:configuration:modelVersion:compilerVersion:loaderEvent:useUpdatableModelLoaders:loadingClasses:error:
func (_MLLoaderClass MLLoaderClass) _loadModelFromArchiveConfigurationModelVersionCompilerVersionLoaderEventUseUpdatableModelLoadersLoadingClassesError(archive unsafe.Pointer, configuration objectivec.IObject, version objectivec.IObject, version2 objectivec.IObject, event objectivec.IObject, loaders bool, classes objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadModelFromArchive:configuration:modelVersion:compilerVersion:loaderEvent:useUpdatableModelLoaders:loadingClasses:error:"), archive, configuration, version, version2, event, loaders, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadModelFromArchiveConfigurationModelVersionCompilerVersionLoaderEventUseUpdatableModelLoadersLoadingClassesError is an exported wrapper for the private method _loadModelFromArchiveConfigurationModelVersionCompilerVersionLoaderEventUseUpdatableModelLoadersLoadingClassesError.
func (_MLLoaderClass MLLoaderClass) LoadModelFromArchiveConfigurationModelVersionCompilerVersionLoaderEventUseUpdatableModelLoadersLoadingClassesError(archive unsafe.Pointer, configuration objectivec.IObject, version objectivec.IObject, version2 objectivec.IObject, event objectivec.IObject, loaders bool, classes objectivec.IObject) (objectivec.IObject, error) {
	return _MLLoaderClass._loadModelFromArchiveConfigurationModelVersionCompilerVersionLoaderEventUseUpdatableModelLoadersLoadingClassesError(archive, configuration, version, version2, event, loaders, classes)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadModelFromAssetAtURL:configuration:loaderEvent:error:
func (_MLLoaderClass MLLoaderClass) _loadModelFromAssetAtURLConfigurationLoaderEventError(url foundation.INSURL, configuration objectivec.IObject, event objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadModelFromAssetAtURL:configuration:loaderEvent:error:"), url, configuration, event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadModelFromAssetAtURLConfigurationLoaderEventError is an exported wrapper for the private method _loadModelFromAssetAtURLConfigurationLoaderEventError.
func (_MLLoaderClass MLLoaderClass) LoadModelFromAssetAtURLConfigurationLoaderEventError(url foundation.INSURL, configuration objectivec.IObject, event objectivec.IObject) (objectivec.IObject, error) {
	return _MLLoaderClass._loadModelFromAssetAtURLConfigurationLoaderEventError(url, configuration, event)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadModelWithClass:fromArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLLoaderClass MLLoaderClass) _loadModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class objc.Class, archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadModelWithClass:fromArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), class, archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError is an exported wrapper for the private method _loadModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError.
func (_MLLoaderClass MLLoaderClass) LoadModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class objc.Class, archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	return _MLLoaderClass._loadModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class, archive, info, info2, configuration)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadUpdatableModelWithClass:fromArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLLoaderClass MLLoaderClass) _loadUpdatableModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class objc.Class, archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadUpdatableModelWithClass:fromArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), class, archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadUpdatableModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError is an exported wrapper for the private method _loadUpdatableModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError.
func (_MLLoaderClass MLLoaderClass) LoadUpdatableModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class objc.Class, archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	return _MLLoaderClass._loadUpdatableModelWithClassFromArchiveModelVersionInfoCompilerVersionInfoConfigurationError(class, archive, info, info2, configuration)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_loadWithModelLoaderFromArchive:configuration:loaderEvent:useUpdatableModelLoaders:error:
func (_MLLoaderClass MLLoaderClass) _loadWithModelLoaderFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive unsafe.Pointer, configuration objectivec.IObject, event objectivec.IObject, loaders bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_loadWithModelLoaderFromArchive:configuration:loaderEvent:useUpdatableModelLoaders:error:"), archive, configuration, event, loaders, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// LoadWithModelLoaderFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError is an exported wrapper for the private method _loadWithModelLoaderFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError.
func (_MLLoaderClass MLLoaderClass) LoadWithModelLoaderFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive unsafe.Pointer, configuration objectivec.IObject, event objectivec.IObject, loaders bool) (objectivec.IObject, error) {
	return _MLLoaderClass._loadWithModelLoaderFromArchiveConfigurationLoaderEventUseUpdatableModelLoadersError(archive, configuration, event, loaders)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_populateLoaderAndPredictionEvent:model:configuration:loadTimeDuration:
func (_MLLoaderClass MLLoaderClass) _populateLoaderAndPredictionEventModelConfigurationLoadTimeDuration(event objectivec.IObject, model objectivec.IObject, configuration objectivec.IObject, duration uint64) {
	objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_populateLoaderAndPredictionEvent:model:configuration:loadTimeDuration:"), event, model, configuration, duration)
}

// PopulateLoaderAndPredictionEventModelConfigurationLoadTimeDuration is an exported wrapper for the private method _populateLoaderAndPredictionEventModelConfigurationLoadTimeDuration.
func (_MLLoaderClass MLLoaderClass) PopulateLoaderAndPredictionEventModelConfigurationLoadTimeDuration(event objectivec.IObject, model objectivec.IObject, configuration objectivec.IObject, duration uint64) {
	_MLLoaderClass._populateLoaderAndPredictionEventModelConfigurationLoadTimeDuration(event, model, configuration, duration)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/_unarchiveCodedModelObjectAtURL:error:
func (_MLLoaderClass MLLoaderClass) _unarchiveCodedModelObjectAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("_unarchiveCodedModelObjectAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// UnarchiveCodedModelObjectAtURLError is an exported wrapper for the private method _unarchiveCodedModelObjectAtURLError.
func (_MLLoaderClass MLLoaderClass) UnarchiveCodedModelObjectAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	return _MLLoaderClass._unarchiveCodedModelObjectAtURLError(url)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/checkAssetPath:error:
func (_MLLoaderClass MLLoaderClass) CheckAssetPathError(path objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLLoaderClass.class), objc.Sel("checkAssetPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkAssetPath:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelAssetDescriptionFromArchive:error:
func (_MLLoaderClass MLLoaderClass) LoadModelAssetDescriptionFromArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelAssetDescriptionFromArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelAssetDescriptionFromAssetAtURL:error:
func (_MLLoaderClass MLLoaderClass) LoadModelAssetDescriptionFromAssetAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelAssetDescriptionFromAssetAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelFromArchive:configuration:error:
func (_MLLoaderClass MLLoaderClass) LoadModelFromArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelFromArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelFromArchive:error:
func (_MLLoaderClass MLLoaderClass) LoadModelFromArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelFromArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelFromAssetAtURL:configuration:error:
func (_MLLoaderClass MLLoaderClass) LoadModelFromAssetAtURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelFromAssetAtURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadModelFromAssetAtURL:error:
func (_MLLoaderClass MLLoaderClass) LoadModelFromAssetAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadModelFromAssetAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadUpdatableModelFromArchive:configuration:error:
func (_MLLoaderClass MLLoaderClass) LoadUpdatableModelFromArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadUpdatableModelFromArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLLoader/loadUpdatableModelFromAssetAtURL:configuration:error:
func (_MLLoaderClass MLLoaderClass) LoadUpdatableModelFromAssetAtURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLoaderClass.class), objc.Sel("loadUpdatableModelFromAssetAtURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

