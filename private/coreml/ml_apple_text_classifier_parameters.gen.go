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
	rv := objc.SendIfResponds[MLAppleTextClassifierParameters](objc.ID(mc.class), objc.Sel("alloc"))
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
	ModelParameterData() foundation.NSData
	SetModelParameterData(value foundation.NSData)
	OutputFeatureName() string
	SetOutputFeatureName(value string)
	Revision() uint64
	SetRevision(value uint64)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error)
	InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error)
}

// Init initializes the instance.
func (m MLAppleTextClassifierParameters) Init() MLAppleTextClassifierParameters {
	rv := objc.SendIfResponds[MLAppleTextClassifierParameters](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLAppleTextClassifierParameters) Autorelease() MLAppleTextClassifierParameters {
	rv := objc.SendIfResponds[MLAppleTextClassifierParameters](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleTextClassifierParameters creates a new MLAppleTextClassifierParameters instance.
func NewMLAppleTextClassifierParameters() MLAppleTextClassifierParameters {
	class := getMLAppleTextClassifierParametersClass()
	rv := objc.SendIfResponds[MLAppleTextClassifierParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAppleTextClassifierParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierParametersClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:"), data, language, name, name2, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleTextClassifierParameters{}, objc.ErrInitFailed
	}
	return MLAppleTextClassifierParametersFromID(rv), nil
}

func NewAppleTextClassifierParametersWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierParametersClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleTextClassifierParameters{}, objc.ErrInitFailed
	}
	return MLAppleTextClassifierParametersFromID(rv), nil
}

func (m MLAppleTextClassifierParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:error:"), data, language, name, name2, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil

}
func (m MLAppleTextClassifierParameters) InitWithDataLanguageInputFeatureNameOutputFeatureNameModelDataLabelNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleTextClassifierParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:language:inputFeatureName:outputFeatureName:modelData:labelNames:metadata:error:"), data, language, name, name2, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifierParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierParametersFromID(rv), nil

}

func (m MLAppleTextClassifierParameters) InputFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleTextClassifierParameters) SetInputFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}
func (m MLAppleTextClassifierParameters) LabelNames() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLAppleTextClassifierParameters) SetLabelNames(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLabelNames:"), value)
}
func (m MLAppleTextClassifierParameters) Language() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleTextClassifierParameters) SetLanguage(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLanguage:"), objc.String(value))
}
func (m MLAppleTextClassifierParameters) Metadata() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLAppleTextClassifierParameters) SetMetadata(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMetadata:"), value)
}
func (m MLAppleTextClassifierParameters) ModelParameterData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLAppleTextClassifierParameters) SetModelParameterData(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelParameterData:"), value)
}
func (m MLAppleTextClassifierParameters) OutputFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAppleTextClassifierParameters) SetOutputFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setOutputFeatureName:"), objc.String(value))
}
func (m MLAppleTextClassifierParameters) Revision() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("revision"))
	return rv
}
func (m MLAppleTextClassifierParameters) SetRevision(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setRevision:"), value)
}
