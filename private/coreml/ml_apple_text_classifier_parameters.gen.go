// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleTextClassifierParameters] class.
var (
	_MLAppleTextClassifierParametersClass     MLAppleTextClassifierParametersClass
	_MLAppleTextClassifierParametersClassOnce sync.Once
)

func getMLAppleTextClassifierParametersClass() MLAppleTextClassifierParametersClass {
	_MLAppleTextClassifierParametersClassOnce.Do(func() {
		_MLAppleTextClassifierParametersClass = MLAppleTextClassifierParametersClass{class: objc.GetClass("MLAppleTextClassifierParameters")}
	})
	return _MLAppleTextClassifierParametersClass
}

// GetMLAppleTextClassifierParametersClass returns the class object for MLAppleTextClassifierParameters.
func GetMLAppleTextClassifierParametersClass() MLAppleTextClassifierParametersClass {
	return getMLAppleTextClassifierParametersClass()
}

type MLAppleTextClassifierParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleTextClassifierParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleTextClassifierParametersClass) Alloc() MLAppleTextClassifierParameters {
	rv := objc.Send[MLAppleTextClassifierParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleTextClassifierParameters.InputFeatureName]
//   - [MLAppleTextClassifierParameters.SetInputFeatureName]
//   - [MLAppleTextClassifierParameters.LabelNames]
//   - [MLAppleTextClassifierParameters.SetLabelNames]
//   - [MLAppleTextClassifierParameters.Language]
//   - [MLAppleTextClassifierParameters.SetLanguage]
//   - [MLAppleTextClassifierParameters.Metadata]
//   - [MLAppleTextClassifierParameters.SetMetadata]
//   - [MLAppleTextClassifierParameters.ModelParameterData]
//   - [MLAppleTextClassifierParameters.SetModelParameterData]
//   - [MLAppleTextClassifierParameters.OutputFeatureName]
//   - [MLAppleTextClassifierParameters.SetOutputFeatureName]
//   - [MLAppleTextClassifierParameters.Revision]
//   - [MLAppleTextClassifierParameters.SetRevision]
//   - [MLAppleTextClassifierParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError]
//   - [MLAppleTextClassifierParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters
type MLAppleTextClassifierParameters struct {
	objectivec.Object
}

// MLAppleTextClassifierParametersFromID constructs a [MLAppleTextClassifierParameters] from an objc.ID.
func MLAppleTextClassifierParametersFromID(id objc.ID) MLAppleTextClassifierParameters {
	return MLAppleTextClassifierParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleTextClassifierParameters implements IMLAppleTextClassifierParameters.
var _ IMLAppleTextClassifierParameters = MLAppleTextClassifierParameters{}

// An interface definition for the [MLAppleTextClassifierParameters] class.
//
// # Methods
//
//   - [IMLAppleTextClassifierParameters.InputFeatureName]
//   - [IMLAppleTextClassifierParameters.SetInputFeatureName]
//   - [IMLAppleTextClassifierParameters.LabelNames]
//   - [IMLAppleTextClassifierParameters.SetLabelNames]
//   - [IMLAppleTextClassifierParameters.Language]
//   - [IMLAppleTextClassifierParameters.SetLanguage]
//   - [IMLAppleTextClassifierParameters.Metadata]
//   - [IMLAppleTextClassifierParameters.SetMetadata]
//   - [IMLAppleTextClassifierParameters.ModelParameterData]
//   - [IMLAppleTextClassifierParameters.SetModelParameterData]
//   - [IMLAppleTextClassifierParameters.OutputFeatureName]
//   - [IMLAppleTextClassifierParameters.SetOutputFeatureName]
//   - [IMLAppleTextClassifierParameters.Revision]
//   - [IMLAppleTextClassifierParameters.SetRevision]
//   - [IMLAppleTextClassifierParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError]
//   - [IMLAppleTextClassifierParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters
type IMLAppleTextClassifierParameters interface {
	objectivec.IObject

	// Topic: Methods

	InputFeatureName() string
	SetInputFeatureName(value string)
	LabelNames() foundation.INSArray
	SetLabelNames(value foundation.INSArray)
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
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error)
}

// Init initializes the instance.
func (a MLAppleTextClassifierParameters) Init() MLAppleTextClassifierParameters {
	rv := objc.Send[MLAppleTextClassifierParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleTextClassifierParameters) Autorelease() MLAppleTextClassifierParameters {
	rv := objc.Send[MLAppleTextClassifierParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleTextClassifierParameters creates a new MLAppleTextClassifierParameters instance.
func NewMLAppleTextClassifierParameters() MLAppleTextClassifierParameters {
	class := getMLAppleTextClassifierParametersClass()
	rv := objc.Send[MLAppleTextClassifierParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:
func NewAppleTextClassifierParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:"), data, language, name, name2, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:
func NewAppleTextClassifierParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:
func (a MLAppleTextClassifierParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:"), data, language, name, name2, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:
func (a MLAppleTextClassifierParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/inputFeatureName
func (a MLAppleTextClassifierParameters) InputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleTextClassifierParameters) SetInputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/labelNames
func (a MLAppleTextClassifierParameters) LabelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a MLAppleTextClassifierParameters) SetLabelNames(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setLabelNames:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/language
func (a MLAppleTextClassifierParameters) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleTextClassifierParameters) SetLanguage(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguage:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/metadata
func (a MLAppleTextClassifierParameters) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a MLAppleTextClassifierParameters) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/modelParameterData
func (a MLAppleTextClassifierParameters) ModelParameterData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a MLAppleTextClassifierParameters) SetModelParameterData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelParameterData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/outputFeatureName
func (a MLAppleTextClassifierParameters) OutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleTextClassifierParameters) SetOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleTextClassifierParameters/revision
func (a MLAppleTextClassifierParameters) Revision() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("revision"))
	return rv
}
func (a MLAppleTextClassifierParameters) SetRevision(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRevision:"), value)
}
