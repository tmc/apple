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
	rv := objc.SendIfResponds[MLAppleWordEmbeddingParameters](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLAppleWordEmbeddingParameters interface {
	objectivec.IObject

	// Topic: Methods

	InputFeatureName() string
	SetInputFeatureName(value string)
	Language() string
	SetLanguage(value string)
	Metadata() foundation.INSDictionary
	SetMetadata(value foundation.INSDictionary)
	ModelParameterData() foundation.NSData
	SetModelParameterData(value foundation.NSData)
	OutputFeatureName() string
	SetOutputFeatureName(value string)
	Revision() uint64
	SetRevision(value uint64)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error)
}

// Init initializes the instance.
func (m MLAppleWordEmbeddingParameters) Init() MLAppleWordEmbeddingParameters {
	rv := objc.SendIfResponds[MLAppleWordEmbeddingParameters](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLAppleWordEmbeddingParameters) Autorelease() MLAppleWordEmbeddingParameters {
	rv := objc.SendIfResponds[MLAppleWordEmbeddingParameters](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleWordEmbeddingParameters creates a new MLAppleWordEmbeddingParameters instance.
func NewMLAppleWordEmbeddingParameters() MLAppleWordEmbeddingParameters {
	class := getMLAppleWordEmbeddingParametersClass()
	rv := objc.SendIfResponds[MLAppleWordEmbeddingParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAppleWordEmbeddingParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingParametersClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:error:"), data, language, name, name2, data2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleWordEmbeddingParameters{}, objc.ErrInitFailed
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil
}

func NewAppleWordEmbeddingParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingParametersClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:"), data, language, name, name2, data2, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleWordEmbeddingParameters{}, objc.ErrInitFailed
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil
}

func (m MLAppleWordEmbeddingParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:error:"), data, language, name, name2, data2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil

}
func (m MLAppleWordEmbeddingParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, metadata objectivec.IObject) (MLAppleWordEmbeddingParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:metadata:error:"), data, language, name, name2, data2, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbeddingParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingParametersFromID(rv), nil

}

func (m MLAppleWordEmbeddingParameters) InputFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleWordEmbeddingParameters) SetInputFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}
func (m MLAppleWordEmbeddingParameters) Language() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleWordEmbeddingParameters) SetLanguage(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLanguage:"), objc.String(value))
}
func (m MLAppleWordEmbeddingParameters) Metadata() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLAppleWordEmbeddingParameters) SetMetadata(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMetadata:"), value)
}
func (m MLAppleWordEmbeddingParameters) ModelParameterData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLAppleWordEmbeddingParameters) SetModelParameterData(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelParameterData:"), value)
}
func (m MLAppleWordEmbeddingParameters) OutputFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleWordEmbeddingParameters) SetOutputFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setOutputFeatureName:"), objc.String(value))
}
func (m MLAppleWordEmbeddingParameters) Revision() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("revision"))
	return rv
}
func (m MLAppleWordEmbeddingParameters) SetRevision(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setRevision:"), value)
}
