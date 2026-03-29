// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAsset] class.
var (
	_MLModelAssetClass     MLModelAssetClass
	_MLModelAssetClassOnce sync.Once
)

func getMLModelAssetClass() MLModelAssetClass {
	_MLModelAssetClassOnce.Do(func() {
		_MLModelAssetClass = MLModelAssetClass{class: objc.GetClass("MLModelAsset")}
	})
	return _MLModelAssetClass
}

// GetMLModelAssetClass returns the class object for MLModelAsset.
func GetMLModelAssetClass() MLModelAssetClass {
	return getMLModelAssetClass()
}

type MLModelAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetClass) Alloc() MLModelAsset {
	rv := objc.Send[MLModelAsset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAsset.ArchiveData]
//   - [MLModelAsset.SetArchiveData]
//   - [MLModelAsset.Classifier]
//   - [MLModelAsset.ClassifierWithError]
//   - [MLModelAsset.CompiledModelURL]
//   - [MLModelAsset.CompiledURL]
//   - [MLModelAsset.DescriptionVendor]
//   - [MLModelAsset.LastConfiguration]
//   - [MLModelAsset.SetLastConfiguration]
//   - [MLModelAsset.Load]
//   - [MLModelAsset.Model]
//   - [MLModelAsset.ModelStructureWithCompletionHandler]
//   - [MLModelAsset.ModelVendor]
//   - [MLModelAsset.ModelWithConfigurationError]
//   - [MLModelAsset.ModelWithError]
//   - [MLModelAsset.Regressor]
//   - [MLModelAsset.RegressorWithError]
//   - [MLModelAsset.ResourceFactory]
//   - [MLModelAsset.StorageType]
//   - [MLModelAsset.StructureVendor]
//   - [MLModelAsset.InitWithArchiveData]
//   - [MLModelAsset.InitWithResourceFactoryConfiguration]
//   - [MLModelAsset.InitWithURLConfigurationError]
//   - [MLModelAsset.InitWithURLError]
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset
type MLModelAsset struct {
	objectivec.Object
}

// MLModelAssetFromID constructs a [MLModelAsset] from an objc.ID.
func MLModelAssetFromID(id objc.ID) MLModelAsset {
	return MLModelAsset{objectivec.Object{ID: id}}
}
// Ensure MLModelAsset implements IMLModelAsset.
var _ IMLModelAsset = MLModelAsset{}

// An interface definition for the [MLModelAsset] class.
//
// # Methods
//
//   - [IMLModelAsset.ArchiveData]
//   - [IMLModelAsset.SetArchiveData]
//   - [IMLModelAsset.Classifier]
//   - [IMLModelAsset.ClassifierWithError]
//   - [IMLModelAsset.CompiledModelURL]
//   - [IMLModelAsset.CompiledURL]
//   - [IMLModelAsset.DescriptionVendor]
//   - [IMLModelAsset.LastConfiguration]
//   - [IMLModelAsset.SetLastConfiguration]
//   - [IMLModelAsset.Load]
//   - [IMLModelAsset.Model]
//   - [IMLModelAsset.ModelStructureWithCompletionHandler]
//   - [IMLModelAsset.ModelVendor]
//   - [IMLModelAsset.ModelWithConfigurationError]
//   - [IMLModelAsset.ModelWithError]
//   - [IMLModelAsset.Regressor]
//   - [IMLModelAsset.RegressorWithError]
//   - [IMLModelAsset.ResourceFactory]
//   - [IMLModelAsset.StorageType]
//   - [IMLModelAsset.StructureVendor]
//   - [IMLModelAsset.InitWithArchiveData]
//   - [IMLModelAsset.InitWithResourceFactoryConfiguration]
//   - [IMLModelAsset.InitWithURLConfigurationError]
//   - [IMLModelAsset.InitWithURLError]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset
type IMLModelAsset interface {
	objectivec.IObject

	// Topic: Methods

	ArchiveData() foundation.INSDictionary
	SetArchiveData(value foundation.INSDictionary)
	Classifier() objectivec.IObject
	ClassifierWithError() (objectivec.IObject, error)
	CompiledModelURL() foundation.INSURL
	CompiledURL() foundation.INSURL
	DescriptionVendor() IMLModelAssetDescriptionVendor
	LastConfiguration() IMLModelConfiguration
	SetLastConfiguration(value IMLModelConfiguration)
	Load(load []objectivec.IObject) bool
	Model() IMLModel
	ModelStructureWithCompletionHandler(handler ErrorHandler)
	ModelVendor() IMLModelAssetModelVendor
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	ModelWithError() (objectivec.IObject, error)
	Regressor() objectivec.IObject
	RegressorWithError() (objectivec.IObject, error)
	ResourceFactory() IMLModelAssetResourceFactory
	StorageType() int64
	StructureVendor() IMLModelAssetModelStructureVendor
	InitWithArchiveData(data objectivec.IObject) MLModelAsset
	InitWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset
	InitWithURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLModelAsset, error)
	InitWithURLError(url foundation.INSURL) (MLModelAsset, error)
}

// Init initializes the instance.
func (m MLModelAsset) Init() MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAsset) Autorelease() MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAsset creates a new MLModelAsset instance.
func NewMLModelAsset() MLModelAsset {
	class := getMLModelAssetClass()
	rv := objc.Send[MLModelAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithArchiveData:
func NewModelAssetWithArchiveData(data objectivec.IObject) MLModelAsset {
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArchiveData:"), data)
	return MLModelAssetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithResourceFactory:configuration:
func NewModelAssetWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset {
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResourceFactory:configuration:"), factory, configuration)
	return MLModelAssetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithURL:configuration:error:
func NewModelAssetWithURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLModelAsset, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithURL:error:
func NewModelAssetWithURLError(url foundation.INSURL) (MLModelAsset, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/classifierWithError:
func (m MLModelAsset) ClassifierWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classifierWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/load:
func (m MLModelAsset) Load(load []objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("load:"), objectivec.IObjectSliceToNSArray(load))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelStructureWithCompletionHandler:
func (m MLModelAsset) ModelStructureWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelWithConfiguration:error:
func (m MLModelAsset) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelWithError:
func (m MLModelAsset) ModelWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/regressorWithError:
func (m MLModelAsset) RegressorWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regressorWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/storageType
func (m MLModelAsset) StorageType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("storageType"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithArchiveData:
func (m MLModelAsset) InitWithArchiveData(data objectivec.IObject) MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("initWithArchiveData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithResourceFactory:configuration:
func (m MLModelAsset) InitWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("initWithResourceFactory:configuration:"), factory, configuration)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithURL:configuration:error:
func (m MLModelAsset) InitWithURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/initWithURL:error:
func (m MLModelAsset) InitWithURLError(url foundation.INSURL) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/_modelAssetWithSpecificationData:blobMapping:error:
func (_MLModelAssetClass MLModelAssetClass) _modelAssetWithSpecificationDataBlobMappingError(data objectivec.IObject, mapping objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("_modelAssetWithSpecificationData:blobMapping:error:"), data, mapping, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/fetchNetworkURLFromCompiledModelAtURL:error:
func (_MLModelAssetClass MLModelAssetClass) FetchNetworkURLFromCompiledModelAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("fetchNetworkURLFromCompiledModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/isANESupported
func (_MLModelAssetClass MLModelAssetClass) IsANESupported() bool {
	rv := objc.Send[bool](objc.ID(_MLModelAssetClass.class), objc.Sel("isANESupported"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetDataByLoadingBlobFileReferencesInModelSpecificationAtURL:blobMapping:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetDataByLoadingBlobFileReferencesInModelSpecificationAtURLBlobMappingError(url foundation.INSURL, mapping []objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetDataByLoadingBlobFileReferencesInModelSpecificationAtURL:blobMapping:error:"), url, objectivec.IObjectSliceToNSArray(mapping), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetDataByResolvingBlobFileReferencesIntoInMemoryValuesInModelSpecificationAtURL:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetDataByResolvingBlobFileReferencesIntoInMemoryValuesInModelSpecificationAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetDataByResolvingBlobFileReferencesIntoInMemoryValuesInModelSpecificationAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecification:compilerOptions:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationCompilerOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecification:compilerOptions:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecification:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationError(specification unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecificationData:blobMapping:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationDataBlobMappingError(data objectivec.IObject, mapping objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), data, mapping, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecificationData:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecificationURL:compilerOptions:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationURLCompilerOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationURL:compilerOptions:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithSpecificationURL:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithURL:configuration:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelAssetWithURL:error:
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/needsANECompilationForModelAtURL:result:error:
func (_MLModelAssetClass MLModelAssetClass) NeedsANECompilationForModelAtURLResultError(url foundation.NSURL) (bool, error) {
	var result bool
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelAssetClass.class), objc.Sel("needsANECompilationForModelAtURL:result:error:"), url, unsafe.Pointer(&result), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("needsANECompilationForModelAtURL:result:error: returned NO with nil NSError")
	}
	return result, nil
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/purgeANEBinaryForModelAtURL:error:
func (_MLModelAssetClass MLModelAssetClass) PurgeANEBinaryForModelAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelAssetClass.class), objc.Sel("purgeANEBinaryForModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("purgeANEBinaryForModelAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/purgeANEIRForModelAtURL:error:
func (_MLModelAssetClass MLModelAssetClass) PurgeANEIRForModelAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelAssetClass.class), objc.Sel("purgeANEIRForModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("purgeANEIRForModelAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/archiveData
func (m MLModelAsset) ArchiveData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("archiveData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelAsset) SetArchiveData(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setArchiveData:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/classifier
func (m MLModelAsset) Classifier() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/compiledModelURL
func (m MLModelAsset) CompiledModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/compiledURL
func (m MLModelAsset) CompiledURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/descriptionVendor
func (m MLModelAsset) DescriptionVendor() IMLModelAssetDescriptionVendor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("descriptionVendor"))
	return MLModelAssetDescriptionVendorFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/lastConfiguration
func (m MLModelAsset) LastConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lastConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModelAsset) SetLastConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setLastConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/model
func (m MLModelAsset) Model() IMLModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelVendor
func (m MLModelAsset) ModelVendor() IMLModelAssetModelVendor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelVendor"))
	return MLModelAssetModelVendorFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/regressor
func (m MLModelAsset) Regressor() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regressor"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/resourceFactory
func (m MLModelAsset) ResourceFactory() IMLModelAssetResourceFactory {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resourceFactory"))
	return MLModelAssetResourceFactoryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/structureVendor
func (m MLModelAsset) StructureVendor() IMLModelAssetModelStructureVendor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("structureVendor"))
	return MLModelAssetModelStructureVendorFromID(objc.ID(rv))
}

// ModelStructure is a synchronous wrapper around [MLModelAsset.ModelStructureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAsset) ModelStructure(ctx context.Context) error {
	done := make(chan error, 1)
	m.ModelStructureWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

