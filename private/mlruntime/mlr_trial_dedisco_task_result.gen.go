// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTrialDediscoTaskResult] class.
var (
	_MLRTrialDediscoTaskResultClass     MLRTrialDediscoTaskResultClass
	_MLRTrialDediscoTaskResultClassOnce sync.Once
)

func getMLRTrialDediscoTaskResultClass() MLRTrialDediscoTaskResultClass {
	_MLRTrialDediscoTaskResultClassOnce.Do(func() {
		_MLRTrialDediscoTaskResultClass = MLRTrialDediscoTaskResultClass{class: objc.GetClass("MLRTrialDediscoTaskResult")}
	})
	return _MLRTrialDediscoTaskResultClass
}

// GetMLRTrialDediscoTaskResultClass returns the class object for MLRTrialDediscoTaskResult.
func GetMLRTrialDediscoTaskResultClass() MLRTrialDediscoTaskResultClass {
	return getMLRTrialDediscoTaskResultClass()
}

type MLRTrialDediscoTaskResultClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTrialDediscoTaskResultClass) Alloc() MLRTrialDediscoTaskResult {
	rv := objc.Send[MLRTrialDediscoTaskResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTrialDediscoTaskResult.AdditionalKeyVariables]
//   - [MLRTrialDediscoTaskResult.NamespaceName]
//   - [MLRTrialDediscoTaskResult.RecipeFactorName]
//   - [MLRTrialDediscoTaskResult.RecordDataEncodingSchemaMetadataErrorOut]
//   - [MLRTrialDediscoTaskResult.SubmitWithTRIClientError]
//   - [MLRTrialDediscoTaskResult.VariablesFromTrialClient]
//   - [MLRTrialDediscoTaskResult.InitWithJSONResultIdentifier]
//   - [MLRTrialDediscoTaskResult.InitWithJSONResultNamespaceNameFactorNameAdditionalKeyVariables]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult
type MLRTrialDediscoTaskResult struct {
	MLRTrialTaskResult
}

// MLRTrialDediscoTaskResultFromID constructs a [MLRTrialDediscoTaskResult] from an objc.ID.
func MLRTrialDediscoTaskResultFromID(id objc.ID) MLRTrialDediscoTaskResult {
	return MLRTrialDediscoTaskResult{MLRTrialTaskResult: MLRTrialTaskResultFromID(id)}
}
// Ensure MLRTrialDediscoTaskResult implements IMLRTrialDediscoTaskResult.
var _ IMLRTrialDediscoTaskResult = MLRTrialDediscoTaskResult{}

// An interface definition for the [MLRTrialDediscoTaskResult] class.
//
// # Methods
//
//   - [IMLRTrialDediscoTaskResult.AdditionalKeyVariables]
//   - [IMLRTrialDediscoTaskResult.NamespaceName]
//   - [IMLRTrialDediscoTaskResult.RecipeFactorName]
//   - [IMLRTrialDediscoTaskResult.RecordDataEncodingSchemaMetadataErrorOut]
//   - [IMLRTrialDediscoTaskResult.SubmitWithTRIClientError]
//   - [IMLRTrialDediscoTaskResult.VariablesFromTrialClient]
//   - [IMLRTrialDediscoTaskResult.InitWithJSONResultIdentifier]
//   - [IMLRTrialDediscoTaskResult.InitWithJSONResultNamespaceNameFactorNameAdditionalKeyVariables]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult
type IMLRTrialDediscoTaskResult interface {
	IMLRTrialTaskResult

	// Topic: Methods

	AdditionalKeyVariables() foundation.INSDictionary
	NamespaceName() string
	RecipeFactorName() string
	RecordDataEncodingSchemaMetadataErrorOut(record objectivec.IObject, data objectivec.IObject, schema objectivec.IObject, metadata objectivec.IObject, out []objectivec.IObject) bool
	SubmitWithTRIClientError(tRIClient objectivec.IObject) (bool, error)
	VariablesFromTrialClient(client objectivec.IObject) objectivec.IObject
	InitWithJSONResultIdentifier(jSONResult objectivec.IObject, identifier objectivec.IObject) MLRTrialDediscoTaskResult
	InitWithJSONResultNamespaceNameFactorNameAdditionalKeyVariables(jSONResult objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, variables objectivec.IObject) MLRTrialDediscoTaskResult
}

// Init initializes the instance.
func (r MLRTrialDediscoTaskResult) Init() MLRTrialDediscoTaskResult {
	rv := objc.Send[MLRTrialDediscoTaskResult](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTrialDediscoTaskResult) Autorelease() MLRTrialDediscoTaskResult {
	rv := objc.Send[MLRTrialDediscoTaskResult](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTrialDediscoTaskResult creates a new MLRTrialDediscoTaskResult instance.
func NewMLRTrialDediscoTaskResult() MLRTrialDediscoTaskResult {
	class := getMLRTrialDediscoTaskResultClass()
	rv := objc.Send[MLRTrialDediscoTaskResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult/initWithJSONResult:
func NewRTrialDediscoTaskResultWithJSONResult(jSONResult objectivec.IObject) MLRTrialDediscoTaskResult {
	instance := getMLRTrialDediscoTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSONResult:"), jSONResult)
	return MLRTrialDediscoTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/initWithJSONResult:identifier:
func NewRTrialDediscoTaskResultWithJSONResultIdentifier(jSONResult objectivec.IObject, identifier objectivec.IObject) MLRTrialDediscoTaskResult {
	instance := getMLRTrialDediscoTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSONResult:identifier:"), jSONResult, identifier)
	return MLRTrialDediscoTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/initWithJSONResult:namespaceName:factorName:additionalKeyVariables:
func NewRTrialDediscoTaskResultWithJSONResultNamespaceNameFactorNameAdditionalKeyVariables(jSONResult objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, variables objectivec.IObject) MLRTrialDediscoTaskResult {
	instance := getMLRTrialDediscoTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSONResult:namespaceName:factorName:additionalKeyVariables:"), jSONResult, name, name2, variables)
	return MLRTrialDediscoTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/record:data:encodingSchema:metadata:errorOut:
func (r MLRTrialDediscoTaskResult) RecordDataEncodingSchemaMetadataErrorOut(record objectivec.IObject, data objectivec.IObject, schema objectivec.IObject, metadata objectivec.IObject, out []objectivec.IObject) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("record:data:encodingSchema:metadata:errorOut:"), record, data, schema, metadata, objectivec.IObjectSliceToNSArray(out))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/submitWithTRIClient:error:
func (r MLRTrialDediscoTaskResult) SubmitWithTRIClientError(tRIClient objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](r.ID, objc.Sel("submitWithTRIClient:error:"), tRIClient, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("submitWithTRIClient:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/variablesFromTrialClient:
func (r MLRTrialDediscoTaskResult) VariablesFromTrialClient(client objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("variablesFromTrialClient:"), client)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/initWithJSONResult:identifier:
func (r MLRTrialDediscoTaskResult) InitWithJSONResultIdentifier(jSONResult objectivec.IObject, identifier objectivec.IObject) MLRTrialDediscoTaskResult {
	rv := objc.Send[MLRTrialDediscoTaskResult](r.ID, objc.Sel("initWithJSONResult:identifier:"), jSONResult, identifier)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/initWithJSONResult:namespaceName:factorName:additionalKeyVariables:
func (r MLRTrialDediscoTaskResult) InitWithJSONResultNamespaceNameFactorNameAdditionalKeyVariables(jSONResult objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, variables objectivec.IObject) MLRTrialDediscoTaskResult {
	rv := objc.Send[MLRTrialDediscoTaskResult](r.ID, objc.Sel("initWithJSONResult:namespaceName:factorName:additionalKeyVariables:"), jSONResult, name, name2, variables)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/baseKeyFromFormat:variables:
func (_MLRTrialDediscoTaskResultClass MLRTrialDediscoTaskResultClass) BaseKeyFromFormatVariables(format objectivec.IObject, variables objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLRTrialDediscoTaskResultClass.class), objc.Sel("baseKeyFromFormat:variables:"), format, variables)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/additionalKeyVariables
func (r MLRTrialDediscoTaskResult) AdditionalKeyVariables() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("additionalKeyVariables"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/namespaceName
func (r MLRTrialDediscoTaskResult) NamespaceName() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("namespaceName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoTaskResult/recipeFactorName
func (r MLRTrialDediscoTaskResult) RecipeFactorName() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recipeFactorName"))
	return foundation.NSStringFromID(rv).String()
}

