// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSecureModel] class.
var (
	_MLSecureModelClass     MLSecureModelClass
	_MLSecureModelClassOnce sync.Once
)

func getMLSecureModelClass() MLSecureModelClass {
	_MLSecureModelClassOnce.Do(func() {
		_MLSecureModelClass = MLSecureModelClass{class: objc.GetClass("MLSecureModel")}
	})
	return _MLSecureModelClass
}

// GetMLSecureModelClass returns the class object for MLSecureModel.
func GetMLSecureModelClass() MLSecureModelClass {
	return getMLSecureModelClass()
}

type MLSecureModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSecureModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSecureModelClass) Alloc() MLSecureModel {
	rv := objc.Send[MLSecureModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSecureModel.ConnectionToModelSecurityService]
//   - [MLSecureModel.SetConnectionToModelSecurityService]
//   - [MLSecureModel.ParameterValueForKeyError]
//   - [MLSecureModel.PredictionFromFeaturesError]
//   - [MLSecureModel.PredictionFromFeaturesOptionsError]
//   - [MLSecureModel.PredictionsFromBatchError]
//   - [MLSecureModel.PredictionsFromBatchOptionsError]
//   - [MLSecureModel.SecureModelProxy]
//   - [MLSecureModel.SetSecureModelProxy]
//   - [MLSecureModel.Metadata]
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel
type MLSecureModel struct {
	MLModel
}

// MLSecureModelFromID constructs a [MLSecureModel] from an objc.ID.
func MLSecureModelFromID(id objc.ID) MLSecureModel {
	return MLSecureModel{MLModel: MLModelFromID(id)}
}
// Ensure MLSecureModel implements IMLSecureModel.
var _ IMLSecureModel = MLSecureModel{}

// An interface definition for the [MLSecureModel] class.
//
// # Methods
//
//   - [IMLSecureModel.ConnectionToModelSecurityService]
//   - [IMLSecureModel.SetConnectionToModelSecurityService]
//   - [IMLSecureModel.ParameterValueForKeyError]
//   - [IMLSecureModel.PredictionFromFeaturesError]
//   - [IMLSecureModel.PredictionFromFeaturesOptionsError]
//   - [IMLSecureModel.PredictionsFromBatchError]
//   - [IMLSecureModel.PredictionsFromBatchOptionsError]
//   - [IMLSecureModel.SecureModelProxy]
//   - [IMLSecureModel.SetSecureModelProxy]
//   - [IMLSecureModel.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel
type IMLSecureModel interface {
	IMLModel

	// Topic: Methods

	ConnectionToModelSecurityService() foundation.NSXPCConnection
	SetConnectionToModelSecurityService(value foundation.NSXPCConnection)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	SecureModelProxy() objectivec.Object
	SetSecureModelProxy(value objectivec.Object)
	Metadata() IMLModelMetadata
}

// Init initializes the instance.
func (s MLSecureModel) Init() MLSecureModel {
	rv := objc.Send[MLSecureModel](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSecureModel) Autorelease() MLSecureModel {
	rv := objc.Send[MLSecureModel](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSecureModel creates a new MLSecureModel instance.
func NewMLSecureModel() MLSecureModel {
	class := getMLSecureModelClass()
	rv := objc.Send[MLSecureModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewSecureModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLSecureModel, error) {
	var errorPtr objc.ID
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSecureModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSecureModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewSecureModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLSecureModel, error) {
	var errorPtr objc.ID
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSecureModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSecureModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/initWithCoder:
func NewSecureModelWithCoder(coder objectivec.IObject) MLSecureModel {
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLSecureModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewSecureModelWithConfiguration(configuration objectivec.IObject) MLSecureModel {
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLSecureModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewSecureModelWithDescription(description objectivec.IObject) MLSecureModel {
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLSecureModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewSecureModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLSecureModel {
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLSecureModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewSecureModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLSecureModel {
	instance := getMLSecureModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLSecureModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/parameterValueForKey:error:
func (s MLSecureModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/predictionFromFeatures:error:
func (s MLSecureModel) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/predictionFromFeatures:options:error:
func (s MLSecureModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/predictionsFromBatch:error:
func (s MLSecureModel) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/predictionsFromBatch:options:error:
func (s MLSecureModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/modelWithContentsOfURL:configuration:decryptCredential:error:
func (_MLSecureModelClass MLSecureModelClass) ModelWithContentsOfURLConfigurationDecryptCredentialError(url foundation.INSURL, configuration objectivec.IObject, credential objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSecureModelClass.class), objc.Sel("modelWithContentsOfURL:configuration:decryptCredential:error:"), url, configuration, credential, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/modelWithContentsOfURL:decryptCredential:error:
func (_MLSecureModelClass MLSecureModelClass) ModelWithContentsOfURLDecryptCredentialError(url foundation.INSURL, credential objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSecureModelClass.class), objc.Sel("modelWithContentsOfURL:decryptCredential:error:"), url, credential, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/sandboxExtensionPathsForModelURL:
func (_MLSecureModelClass MLSecureModelClass) SandboxExtensionPathsForModelURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLSecureModelClass.class), objc.Sel("sandboxExtensionPathsForModelURL:"), url)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/sandboxExtensionTokenForModelURL:
func (_MLSecureModelClass MLSecureModelClass) SandboxExtensionTokenForModelURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLSecureModelClass.class), objc.Sel("sandboxExtensionTokenForModelURL:"), url)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/supportsSecureCoding
func (_MLSecureModelClass MLSecureModelClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLSecureModelClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/connectionToModelSecurityService
func (s MLSecureModel) ConnectionToModelSecurityService() foundation.NSXPCConnection {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("connectionToModelSecurityService"))
	return foundation.NSXPCConnectionFromID(objc.ID(rv))
}
func (s MLSecureModel) SetConnectionToModelSecurityService(value foundation.NSXPCConnection) {
	objc.Send[struct{}](s.ID, objc.Sel("setConnectionToModelSecurityService:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/metadata
func (s MLSecureModel) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSecureModel/secureModelProxy
func (s MLSecureModel) SecureModelProxy() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("secureModelProxy"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s MLSecureModel) SetSecureModelProxy(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setSecureModelProxy:"), value)
}

