// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleWordTaggerParameters] class.
var (
	_MLAppleWordTaggerParametersClass     MLAppleWordTaggerParametersClass
	_MLAppleWordTaggerParametersClassOnce sync.Once
)

func getMLAppleWordTaggerParametersClass() MLAppleWordTaggerParametersClass {
	_MLAppleWordTaggerParametersClassOnce.Do(func() {
		_MLAppleWordTaggerParametersClass = MLAppleWordTaggerParametersClass{class: objc.GetClass("MLAppleWordTaggerParameters")}
	})
	return _MLAppleWordTaggerParametersClass
}

// GetMLAppleWordTaggerParametersClass returns the class object for MLAppleWordTaggerParameters.
func GetMLAppleWordTaggerParametersClass() MLAppleWordTaggerParametersClass {
	return getMLAppleWordTaggerParametersClass()
}

type MLAppleWordTaggerParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleWordTaggerParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleWordTaggerParametersClass) Alloc() MLAppleWordTaggerParameters {
	rv := objc.Send[MLAppleWordTaggerParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleWordTaggerParameters.InputFeatureName]
//   - [MLAppleWordTaggerParameters.SetInputFeatureName]
//   - [MLAppleWordTaggerParameters.Language]
//   - [MLAppleWordTaggerParameters.SetLanguage]
//   - [MLAppleWordTaggerParameters.Metadata]
//   - [MLAppleWordTaggerParameters.SetMetadata]
//   - [MLAppleWordTaggerParameters.ModelParameterData]
//   - [MLAppleWordTaggerParameters.SetModelParameterData]
//   - [MLAppleWordTaggerParameters.Revision]
//   - [MLAppleWordTaggerParameters.SetRevision]
//   - [MLAppleWordTaggerParameters.TagNames]
//   - [MLAppleWordTaggerParameters.SetTagNames]
//   - [MLAppleWordTaggerParameters.TokenLengthsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.SetTokenLengthsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.TokenLocationsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.SetTokenLocationsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.TokenTagsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.SetTokenTagsOutputFeatureName]
//   - [MLAppleWordTaggerParameters.TokensOutputFeatureName]
//   - [MLAppleWordTaggerParameters.SetTokensOutputFeatureName]
//   - [MLAppleWordTaggerParameters.InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesError]
//   - [MLAppleWordTaggerParameters.InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesMetadataError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters
type MLAppleWordTaggerParameters struct {
	objectivec.Object
}

// MLAppleWordTaggerParametersFromID constructs a [MLAppleWordTaggerParameters] from an objc.ID.
func MLAppleWordTaggerParametersFromID(id objc.ID) MLAppleWordTaggerParameters {
	return MLAppleWordTaggerParameters{objectivec.Object{ID: id}}
}
// Ensure MLAppleWordTaggerParameters implements IMLAppleWordTaggerParameters.
var _ IMLAppleWordTaggerParameters = MLAppleWordTaggerParameters{}

// An interface definition for the [MLAppleWordTaggerParameters] class.
//
// # Methods
//
//   - [IMLAppleWordTaggerParameters.InputFeatureName]
//   - [IMLAppleWordTaggerParameters.SetInputFeatureName]
//   - [IMLAppleWordTaggerParameters.Language]
//   - [IMLAppleWordTaggerParameters.SetLanguage]
//   - [IMLAppleWordTaggerParameters.Metadata]
//   - [IMLAppleWordTaggerParameters.SetMetadata]
//   - [IMLAppleWordTaggerParameters.ModelParameterData]
//   - [IMLAppleWordTaggerParameters.SetModelParameterData]
//   - [IMLAppleWordTaggerParameters.Revision]
//   - [IMLAppleWordTaggerParameters.SetRevision]
//   - [IMLAppleWordTaggerParameters.TagNames]
//   - [IMLAppleWordTaggerParameters.SetTagNames]
//   - [IMLAppleWordTaggerParameters.TokenLengthsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.SetTokenLengthsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.TokenLocationsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.SetTokenLocationsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.TokenTagsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.SetTokenTagsOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.TokensOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.SetTokensOutputFeatureName]
//   - [IMLAppleWordTaggerParameters.InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesError]
//   - [IMLAppleWordTaggerParameters.InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesMetadataError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters
type IMLAppleWordTaggerParameters interface {
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
	Revision() uint64
	SetRevision(value uint64)
	TagNames() foundation.INSArray
	SetTagNames(value foundation.INSArray)
	TokenLengthsOutputFeatureName() string
	SetTokenLengthsOutputFeatureName(value string)
	TokenLocationsOutputFeatureName() string
	SetTokenLocationsOutputFeatureName(value string)
	TokenTagsOutputFeatureName() string
	SetTokenTagsOutputFeatureName(value string)
	TokensOutputFeatureName() string
	SetTokensOutputFeatureName(value string)
	InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleWordTaggerParameters, error)
	InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleWordTaggerParameters, error)
}

// Init initializes the instance.
func (a MLAppleWordTaggerParameters) Init() MLAppleWordTaggerParameters {
	rv := objc.Send[MLAppleWordTaggerParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleWordTaggerParameters) Autorelease() MLAppleWordTaggerParameters {
	rv := objc.Send[MLAppleWordTaggerParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleWordTaggerParameters creates a new MLAppleWordTaggerParameters instance.
func NewMLAppleWordTaggerParameters() MLAppleWordTaggerParameters {
	class := getMLAppleWordTaggerParametersClass()
	rv := objc.Send[MLAppleWordTaggerParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:error:
func NewAppleWordTaggerParametersWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleWordTaggerParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordTaggerParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:error:"), data, language, name, name2, name3, name4, name5, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTaggerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:metadata:error:
func NewAppleWordTaggerParametersWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleWordTaggerParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordTaggerParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:metadata:error:"), data, language, name, name2, name3, name4, name5, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTaggerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:error:
func (a MLAppleWordTaggerParameters) InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject) (MLAppleWordTaggerParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:error:"), data, language, name, name2, name3, name4, name5, data2, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTaggerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerParametersFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:metadata:error:
func (a MLAppleWordTaggerParameters) InitWithDataLanguageInputFeatureNameTokensFeatureNameTokenTagsFeatureNameTokenLocationsFeatureNameTokenLengthsFeatureNameModelDataTagNamesMetadataError(data uint64, language objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, data2 objectivec.IObject, names objectivec.IObject, metadata objectivec.IObject) (MLAppleWordTaggerParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:language:inputFeatureName:tokensFeatureName:tokenTagsFeatureName:tokenLocationsFeatureName:tokenLengthsFeatureName:modelData:tagNames:metadata:error:"), data, language, name, name2, name3, name4, name5, data2, names, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTaggerParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/inputFeatureName
func (a MLAppleWordTaggerParameters) InputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetInputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setInputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/language
func (a MLAppleWordTaggerParameters) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetLanguage(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguage:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/metadata
func (a MLAppleWordTaggerParameters) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a MLAppleWordTaggerParameters) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/modelParameterData
func (a MLAppleWordTaggerParameters) ModelParameterData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelParameterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a MLAppleWordTaggerParameters) SetModelParameterData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelParameterData:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/revision
func (a MLAppleWordTaggerParameters) Revision() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("revision"))
	return rv
}
func (a MLAppleWordTaggerParameters) SetRevision(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRevision:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/tagNames
func (a MLAppleWordTaggerParameters) TagNames() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tagNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a MLAppleWordTaggerParameters) SetTagNames(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setTagNames:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/tokenLengthsOutputFeatureName
func (a MLAppleWordTaggerParameters) TokenLengthsOutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tokenLengthsOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetTokenLengthsOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setTokenLengthsOutputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/tokenLocationsOutputFeatureName
func (a MLAppleWordTaggerParameters) TokenLocationsOutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tokenLocationsOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetTokenLocationsOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setTokenLocationsOutputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/tokenTagsOutputFeatureName
func (a MLAppleWordTaggerParameters) TokenTagsOutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tokenTagsOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetTokenTagsOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setTokenTagsOutputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTaggerParameters/tokensOutputFeatureName
func (a MLAppleWordTaggerParameters) TokensOutputFeatureName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tokensOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (a MLAppleWordTaggerParameters) SetTokensOutputFeatureName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setTokensOutputFeatureName:"), objc.String(value))
}

