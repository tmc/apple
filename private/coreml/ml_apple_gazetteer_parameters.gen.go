// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleGazetteerParameters] class.
var (
	_MLAppleGazetteerParametersClass     MLAppleGazetteerParametersClass
	_MLAppleGazetteerParametersClassOnce sync.Once
)

func getMLAppleGazetteerParametersClass() MLAppleGazetteerParametersClass {
	_MLAppleGazetteerParametersClassOnce.Do(func() {
		_MLAppleGazetteerParametersClass = MLAppleGazetteerParametersClass{class: objc.GetClass("MLAppleGazetteerParameters")}
	})
	return _MLAppleGazetteerParametersClass
}

// GetMLAppleGazetteerParametersClass returns the class object for MLAppleGazetteerParameters.
func GetMLAppleGazetteerParametersClass() MLAppleGazetteerParametersClass {
	return getMLAppleGazetteerParametersClass()
}

type MLAppleGazetteerParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleGazetteerParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleGazetteerParametersClass) Alloc() MLAppleGazetteerParameters {
	rv := objc.Send[MLAppleGazetteerParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleGazetteerParameters.InputFeatureName]
//   - [MLAppleGazetteerParameters.SetInputFeatureName]
//   - [MLAppleGazetteerParameters.LabelNames]
//   - [MLAppleGazetteerParameters.SetLabelNames]
//   - [MLAppleGazetteerParameters.Language]
//   - [MLAppleGazetteerParameters.SetLanguage]
//   - [MLAppleGazetteerParameters.Metadata]
//   - [MLAppleGazetteerParameters.SetMetadata]
//   - [MLAppleGazetteerParameters.ModelParameterData]
//   - [MLAppleGazetteerParameters.SetModelParameterData]
//   - [MLAppleGazetteerParameters.OutputFeatureName]
//   - [MLAppleGazetteerParameters.SetOutputFeatureName]
//   - [MLAppleGazetteerParameters.Revision]
//   - [MLAppleGazetteerParameters.SetRevision]
//   - [MLAppleGazetteerParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters
type MLAppleGazetteerParameters struct {
	objectivec.Object
}

// MLAppleGazetteerParametersFromID constructs a [MLAppleGazetteerParameters] from an objc.ID.
func MLAppleGazetteerParametersFromID(id objc.ID) MLAppleGazetteerParameters {
	return MLAppleGazetteerParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleGazetteerParameters implements IMLAppleGazetteerParameters.
var _ IMLAppleGazetteerParameters = MLAppleGazetteerParameters{}

// An interface definition for the [MLAppleGazetteerParameters] class.
//
// # Methods
//
//   - [IMLAppleGazetteerParameters.InputFeatureName]
//   - [IMLAppleGazetteerParameters.SetInputFeatureName]
//   - [IMLAppleGazetteerParameters.LabelNames]
//   - [IMLAppleGazetteerParameters.SetLabelNames]
//   - [IMLAppleGazetteerParameters.Language]
//   - [IMLAppleGazetteerParameters.SetLanguage]
//   - [IMLAppleGazetteerParameters.Metadata]
//   - [IMLAppleGazetteerParameters.SetMetadata]
//   - [IMLAppleGazetteerParameters.ModelParameterData]
//   - [IMLAppleGazetteerParameters.SetModelParameterData]
//   - [IMLAppleGazetteerParameters.OutputFeatureName]
//   - [IMLAppleGazetteerParameters.SetOutputFeatureName]
//   - [IMLAppleGazetteerParameters.Revision]
//   - [IMLAppleGazetteerParameters.SetRevision]
//   - [IMLAppleGazetteerParameters.InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters
type IMLAppleGazetteerParameters interface {
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
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleGazetteerParameters, error)
}

// Init initializes the instance.
func (a MLAppleGazetteerParameters) Init() MLAppleGazetteerParameters {
	rv := objc.Send[MLAppleGazetteerParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleGazetteerParameters) Autorelease() MLAppleGazetteerParameters {
	rv := objc.Send[MLAppleGazetteerParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleGazetteerParameters creates a new MLAppleGazetteerParameters instance.
func NewMLAppleGazetteerParameters() MLAppleGazetteerParameters {
	class := getMLAppleGazetteerParametersClass()
	rv := objc.Send[MLAppleGazetteerParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:
func NewAppleGazetteerParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleGazetteerParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleGazetteerParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:
func (a MLAppleGazetteerParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleGazetteerParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/inputFeatureName
func (a MLAppleGazetteerParameters) InputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleGazetteerParameters) SetInputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/labelNames
func (a MLAppleGazetteerParameters) LabelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a MLAppleGazetteerParameters) SetLabelNames(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setLabelNames:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/language
func (a MLAppleGazetteerParameters) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleGazetteerParameters) SetLanguage(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguage:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/metadata
func (a MLAppleGazetteerParameters) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a MLAppleGazetteerParameters) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/modelParameterData
func (a MLAppleGazetteerParameters) ModelParameterData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a MLAppleGazetteerParameters) SetModelParameterData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelParameterData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/outputFeatureName
func (a MLAppleGazetteerParameters) OutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleGazetteerParameters) SetOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteerParameters/revision
func (a MLAppleGazetteerParameters) Revision() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("revision"))
	return rv
}
func (a MLAppleGazetteerParameters) SetRevision(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRevision:"), value)
}
