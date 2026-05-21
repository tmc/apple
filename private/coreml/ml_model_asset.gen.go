// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
type IMLModelAsset interface {
	objectivec.IObject

	// Topic: Methods

	ArchiveData() foundation.INSDictionary
	SetArchiveData(value foundation.INSDictionary)
	Classifier() unsafe.Pointer
	ClassifierWithError() (objectivec.IObject, error)
	CompiledModelURL() foundation.NSURL
	CompiledURL() foundation.NSURL
	DescriptionVendor() IMLModelAssetDescriptionVendor
	LastConfiguration() IMLModelConfiguration
	SetLastConfiguration(value IMLModelConfiguration)
	Load(load []objectivec.IObject) bool
	Model() IMLModel
	ModelStructureWithCompletionHandler(handler ErrorHandler)
	ModelVendor() IMLModelAssetModelVendor
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	ModelWithError() (objectivec.IObject, error)
	Regressor() unsafe.Pointer
	RegressorWithError() (objectivec.IObject, error)
	ResourceFactory() IMLModelAssetResourceFactory
	StorageType() int64
	StructureVendor() IMLModelAssetModelStructureVendor
	InitWithArchiveData(data objectivec.IObject) MLModelAsset
	InitWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset
	InitWithURLConfigurationError(url foundation.NSURL, configuration objectivec.IObject) (MLModelAsset, error)
	InitWithURLError(url foundation.NSURL) (MLModelAsset, error)
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

func NewModelAssetWithArchiveData(data objectivec.IObject) MLModelAsset {
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArchiveData:"), data)
	return MLModelAssetFromID(rv)
}

func NewModelAssetWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset {
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResourceFactory:configuration:"), factory, configuration)
	return MLModelAssetFromID(rv)
}

func NewModelAssetWithURLConfigurationError(url foundation.NSURL, configuration objectivec.IObject) (MLModelAsset, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}

func NewModelAssetWithURLError(url foundation.NSURL) (MLModelAsset, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}

func (m MLModelAsset) ClassifierWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classifierWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelAsset) Load(load []objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("load:"), objectivec.IObjectSliceToNSArray(load))
	return rv
}
func (m MLModelAsset) ModelStructureWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithCompletionHandler:"), _block0)
}
func (m MLModelAsset) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelAsset) ModelWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelAsset) RegressorWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regressorWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelAsset) StorageType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("storageType"))
	return rv
}
func (m MLModelAsset) InitWithArchiveData(data objectivec.IObject) MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("initWithArchiveData:"), data)
	return rv
}
func (m MLModelAsset) InitWithResourceFactoryConfiguration(factory objectivec.IObject, configuration objectivec.IObject) MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("initWithResourceFactory:configuration:"), factory, configuration)
	return rv
}
func (m MLModelAsset) InitWithURLConfigurationError(url foundation.NSURL, configuration objectivec.IObject) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil

}
func (m MLModelAsset) InitWithURLError(url foundation.NSURL) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAsset{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil

}

func (_MLModelAssetClass MLModelAssetClass) _modelAssetWithSpecificationDataBlobMappingError(data objectivec.IObject, mapping objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("_modelAssetWithSpecificationData:blobMapping:error:"), data, mapping, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) FetchNetworkURLFromCompiledModelAtURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("fetchNetworkURLFromCompiledModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) IsANESupported() bool {
	rv := objc.Send[bool](objc.ID(_MLModelAssetClass.class), objc.Sel("isANESupported"))
	return rv
}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetDataByLoadingBlobFileReferencesInModelSpecificationAtURLBlobMappingError(url foundation.NSURL, mapping []objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetDataByLoadingBlobFileReferencesInModelSpecificationAtURL:blobMapping:error:"), url, objectivec.IObjectSliceToNSArray(mapping), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetDataByResolvingBlobFileReferencesIntoInMemoryValuesInModelSpecificationAtURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetDataByResolvingBlobFileReferencesIntoInMemoryValuesInModelSpecificationAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationCompilerOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecification:compilerOptions:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationError(specification unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationDataBlobMappingError(data objectivec.IObject, mapping objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), data, mapping, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationURLCompilerOptionsError(url foundation.NSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationURL:compilerOptions:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithSpecificationURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithSpecificationURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithURLConfigurationError(url foundation.NSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelAssetClass MLModelAssetClass) ModelAssetWithURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetClass.class), objc.Sel("modelAssetWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (_MLModelAssetClass MLModelAssetClass) PurgeANEBinaryForModelAtURLError(url foundation.NSURL) (bool, error) {
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
func (_MLModelAssetClass MLModelAssetClass) PurgeANEIRForModelAtURLError(url foundation.NSURL) (bool, error) {
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

func (m MLModelAsset) ArchiveData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("archiveData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelAsset) SetArchiveData(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setArchiveData:"), value)
}
func (m MLModelAsset) Classifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("classifier"))
	return rv
}
func (m MLModelAsset) CompiledModelURL() foundation.NSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelAsset) CompiledURL() foundation.NSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelAsset) DescriptionVendor() IMLModelAssetDescriptionVendor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("descriptionVendor"))
	return MLModelAssetDescriptionVendorFromID(objc.ID(rv))
}
func (m MLModelAsset) LastConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lastConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModelAsset) SetLastConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setLastConfiguration:"), value)
}
func (m MLModelAsset) Model() IMLModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLModelAsset) ModelVendor() IMLModelAssetModelVendor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelVendor"))
	return MLModelAssetModelVendorFromID(objc.ID(rv))
}
func (m MLModelAsset) Regressor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("regressor"))
	return rv
}
func (m MLModelAsset) ResourceFactory() IMLModelAssetResourceFactory {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resourceFactory"))
	return MLModelAssetResourceFactoryFromID(objc.ID(rv))
}
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
