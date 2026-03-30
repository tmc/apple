// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleWordEmbeddingParameters] class.
var (
	_MLAppleWordEmbeddingParametersClass     MLAppleWordEmbeddingParametersClass
	_MLAppleWordEmbeddingParametersClassOnce sync.Once
)

func getMLAppleWordEmbeddingParametersClass() MLAppleWordEmbeddingParametersClass {
	_MLAppleWordEmbeddingParametersClassOnce.Do(func() {
		_MLAppleWordEmbeddingParametersClass = MLAppleWordEmbeddingParametersClass{class: objc.GetClass("MLAppleWordEmbeddingParameters")}
	})
	return _MLAppleWordEmbeddingParametersClass
}

// GetMLAppleWordEmbeddingParametersClass returns the class object for MLAppleWordEmbeddingParameters.
func GetMLAppleWordEmbeddingParametersClass() MLAppleWordEmbeddingParametersClass {
	return getMLAppleWordEmbeddingParametersClass()
}

type MLAppleWordEmbeddingParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleWordEmbeddingParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleWordEmbeddingParametersClass) Alloc() MLAppleWordEmbeddingParameters {
	rv := objc.Send[MLAppleWordEmbeddingParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleWordEmbeddingParameters.InputFeatureName]
//   - [MLAppleWordEmbeddingParameters.SetInputFeatureName]
//   - [MLAppleWordEmbeddingParameters.Language]
//   - [MLAppleWordEmbeddingParameters.SetLanguage]
//   - [MLAppleWordEmbeddingParameters.Metadata]
//   - [MLAppleWordEmbeddingParameters.SetMetadata]
//   - [MLAppleWordEmbeddingParameters.ModelParameterData]
//   - [MLAppleWordEmbeddingParameters.SetModelParameterData]
//   - [MLAppleWordEmbeddingParameters.OutputFeatureName]
//   - [MLAppleWordEmbeddingParameters.SetOutputFeatureName]
//   - [MLAppleWordEmbeddingParameters.Revision]
//   - [MLAppleWordEmbeddingParameters.SetRevision]
//   - [MLAppleWordEmbeddingParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError]
//   - [MLAppleWordEmbeddingParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters
type MLAppleWordEmbeddingParameters struct {
	objectivec.Object
}

// MLAppleWordEmbeddingParametersFromID constructs a [MLAppleWordEmbeddingParameters] from an objc.ID.
func MLAppleWordEmbeddingParametersFromID(id objc.ID) MLAppleWordEmbeddingParameters {
	return MLAppleWordEmbeddingParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleWordEmbeddingParameters implements IMLAppleWordEmbeddingParameters.
var _ IMLAppleWordEmbeddingParameters = MLAppleWordEmbeddingParameters{}

// An interface definition for the [MLAppleWordEmbeddingParameters] class.
//
// # Methods
//
//   - [IMLAppleWordEmbeddingParameters.InputFeatureName]
//   - [IMLAppleWordEmbeddingParameters.SetInputFeatureName]
//   - [IMLAppleWordEmbeddingParameters.Language]
//   - [IMLAppleWordEmbeddingParameters.SetLanguage]
//   - [IMLAppleWordEmbeddingParameters.Metadata]
//   - [IMLAppleWordEmbeddingParameters.SetMetadata]
//   - [IMLAppleWordEmbeddingParameters.ModelParameterData]
//   - [IMLAppleWordEmbeddingParameters.SetModelParameterData]
//   - [IMLAppleWordEmbeddingParameters.OutputFeatureName]
//   - [IMLAppleWordEmbeddingParameters.SetOutputFeatureName]
//   - [IMLAppleWordEmbeddingParameters.Revision]
//   - [IMLAppleWordEmbeddingParameters.SetRevision]
//   - [IMLAppleWordEmbeddingParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError]
//   - [IMLAppleWordEmbeddingParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters
type IMLAppleWordEmbeddingParameters interface {
	objectivec.IObject

	// Topic: Methods

	InputFeatureName() string
	SetInputFeatureName(value string)
	Language() string
	SetLanguage(value string)
	Metadata() foundation.INSDictionary
	SetMetadata(value foundation.INSDictionary)
	ModelParameterData() foundation.INSData
	SetModelParameterData(value foundation.INSData)
	OutputFeatureName() string
	SetOutputFeatureName(value string)
	Revision() uint64
	SetRevision(value uint64)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error)
}

// Init initializes the instance.
func (a MLAppleWordEmbeddingParameters) Init() MLAppleWordEmbeddingParameters {
	rv := objc.Send[MLAppleWordEmbeddingParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleWordEmbeddingParameters) Autorelease() MLAppleWordEmbeddingParameters {
	rv := objc.Send[MLAppleWordEmbeddingParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleWordEmbeddingParameters creates a new MLAppleWordEmbeddingParameters instance.
func NewMLAppleWordEmbeddingParameters() MLAppleWordEmbeddingParameters {
	class := getMLAppleWordEmbeddingParametersClass()
	rv := objc.Send[MLAppleWordEmbeddingParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:error:
func NewAppleWordEmbeddingParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:error:"), data, language, name, name2, data2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:
func NewAppleWordEmbeddingParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:"), data, language, name, name2, data2, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:error:
func (a MLAppleWordEmbeddingParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:error:"), data, language, name, name2, data2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:
func (a MLAppleWordEmbeddingParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:"), data, language, name, name2, data2, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/inputFeatureName
func (a MLAppleWordEmbeddingParameters) InputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordEmbeddingParameters) SetInputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/language
func (a MLAppleWordEmbeddingParameters) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordEmbeddingParameters) SetLanguage(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguage:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/metadata
func (a MLAppleWordEmbeddingParameters) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a MLAppleWordEmbeddingParameters) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/modelParameterData
func (a MLAppleWordEmbeddingParameters) ModelParameterData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a MLAppleWordEmbeddingParameters) SetModelParameterData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelParameterData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/outputFeatureName
func (a MLAppleWordEmbeddingParameters) OutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordEmbeddingParameters) SetOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbeddingParameters/revision
func (a MLAppleWordEmbeddingParameters) Revision() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("revision"))
	return rv
}
func (a MLAppleWordEmbeddingParameters) SetRevision(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRevision:"), value)
}
