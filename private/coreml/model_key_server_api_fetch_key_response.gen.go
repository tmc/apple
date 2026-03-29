// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPIFetchKeyResponse] class.
var (
	_ModelKeyServerAPIFetchKeyResponseClass     ModelKeyServerAPIFetchKeyResponseClass
	_ModelKeyServerAPIFetchKeyResponseClassOnce sync.Once
)

func getModelKeyServerAPIFetchKeyResponseClass() ModelKeyServerAPIFetchKeyResponseClass {
	_ModelKeyServerAPIFetchKeyResponseClassOnce.Do(func() {
		_ModelKeyServerAPIFetchKeyResponseClass = ModelKeyServerAPIFetchKeyResponseClass{class: objc.GetClass("ModelKeyServerAPIFetchKeyResponse")}
	})
	return _ModelKeyServerAPIFetchKeyResponseClass
}

// GetModelKeyServerAPIFetchKeyResponseClass returns the class object for ModelKeyServerAPIFetchKeyResponse.
func GetModelKeyServerAPIFetchKeyResponseClass() ModelKeyServerAPIFetchKeyResponseClass {
	return getModelKeyServerAPIFetchKeyResponseClass()
}

type ModelKeyServerAPIFetchKeyResponseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPIFetchKeyResponseClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPIFetchKeyResponseClass) Alloc() ModelKeyServerAPIFetchKeyResponse {
	rv := objc.Send[ModelKeyServerAPIFetchKeyResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ModelKeyServerAPIFetchKeyResponse.StringAsResult]
//   - [ModelKeyServerAPIFetchKeyResponse.ClearOneofValuesForResult]
//   - [ModelKeyServerAPIFetchKeyResponse.CopyTo]
//   - [ModelKeyServerAPIFetchKeyResponse.DictionaryRepresentation]
//   - [ModelKeyServerAPIFetchKeyResponse.Error]
//   - [ModelKeyServerAPIFetchKeyResponse.SetError]
//   - [ModelKeyServerAPIFetchKeyResponse.HasError]
//   - [ModelKeyServerAPIFetchKeyResponse.HasResult]
//   - [ModelKeyServerAPIFetchKeyResponse.SetHasResult]
//   - [ModelKeyServerAPIFetchKeyResponse.HasSuccess]
//   - [ModelKeyServerAPIFetchKeyResponse.MergeFrom]
//   - [ModelKeyServerAPIFetchKeyResponse.ReadFrom]
//   - [ModelKeyServerAPIFetchKeyResponse.Result]
//   - [ModelKeyServerAPIFetchKeyResponse.SetResult]
//   - [ModelKeyServerAPIFetchKeyResponse.ResultAsString]
//   - [ModelKeyServerAPIFetchKeyResponse.Success]
//   - [ModelKeyServerAPIFetchKeyResponse.SetSuccess]
//   - [ModelKeyServerAPIFetchKeyResponse.WriteTo]
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse
type ModelKeyServerAPIFetchKeyResponse struct {
	objectivec.Object
}

// ModelKeyServerAPIFetchKeyResponseFromID constructs a [ModelKeyServerAPIFetchKeyResponse] from an objc.ID.
func ModelKeyServerAPIFetchKeyResponseFromID(id objc.ID) ModelKeyServerAPIFetchKeyResponse {
	return ModelKeyServerAPIFetchKeyResponse{objectivec.Object{ID: id}}
}
// NOTE: ModelKeyServerAPIFetchKeyResponse struct embeds objectivec.Object (parent type unavailable) but
// IModelKeyServerAPIFetchKeyResponse embeds the parent interface; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPIFetchKeyResponse] class.
//
// # Methods
//
//   - [IModelKeyServerAPIFetchKeyResponse.StringAsResult]
//   - [IModelKeyServerAPIFetchKeyResponse.ClearOneofValuesForResult]
//   - [IModelKeyServerAPIFetchKeyResponse.CopyTo]
//   - [IModelKeyServerAPIFetchKeyResponse.DictionaryRepresentation]
//   - [IModelKeyServerAPIFetchKeyResponse.Error]
//   - [IModelKeyServerAPIFetchKeyResponse.SetError]
//   - [IModelKeyServerAPIFetchKeyResponse.HasError]
//   - [IModelKeyServerAPIFetchKeyResponse.HasResult]
//   - [IModelKeyServerAPIFetchKeyResponse.SetHasResult]
//   - [IModelKeyServerAPIFetchKeyResponse.HasSuccess]
//   - [IModelKeyServerAPIFetchKeyResponse.MergeFrom]
//   - [IModelKeyServerAPIFetchKeyResponse.ReadFrom]
//   - [IModelKeyServerAPIFetchKeyResponse.Result]
//   - [IModelKeyServerAPIFetchKeyResponse.SetResult]
//   - [IModelKeyServerAPIFetchKeyResponse.ResultAsString]
//   - [IModelKeyServerAPIFetchKeyResponse.Success]
//   - [IModelKeyServerAPIFetchKeyResponse.SetSuccess]
//   - [IModelKeyServerAPIFetchKeyResponse.WriteTo]
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse
type IModelKeyServerAPIFetchKeyResponse interface {
	IPBCodable

	// Topic: Methods

	StringAsResult(result objectivec.IObject) int
	ClearOneofValuesForResult()
	CopyTo(to objectivec.IObject)
	DictionaryRepresentation() objectivec.IObject
	Error() IModelKeyServerAPIResultError
	SetError(value IModelKeyServerAPIResultError)
	HasError() bool
	HasResult() bool
	SetHasResult(value bool)
	HasSuccess() bool
	MergeFrom(from objectivec.IObject)
	ReadFrom(from objectivec.IObject) bool
	Result() int
	SetResult(value int)
	ResultAsString(string_ int) objectivec.IObject
	Success() IModelKeyServerAPIFetchKeyResult
	SetSuccess(value IModelKeyServerAPIFetchKeyResult)
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPIFetchKeyResponse) Init() ModelKeyServerAPIFetchKeyResponse {
	rv := objc.Send[ModelKeyServerAPIFetchKeyResponse](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPIFetchKeyResponse) Autorelease() ModelKeyServerAPIFetchKeyResponse {
	rv := objc.Send[ModelKeyServerAPIFetchKeyResponse](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPIFetchKeyResponse creates a new ModelKeyServerAPIFetchKeyResponse instance.
func NewModelKeyServerAPIFetchKeyResponse() ModelKeyServerAPIFetchKeyResponse {
	class := getModelKeyServerAPIFetchKeyResponseClass()
	rv := objc.Send[ModelKeyServerAPIFetchKeyResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/StringAsResult:
func (m ModelKeyServerAPIFetchKeyResponse) StringAsResult(result objectivec.IObject) int {
	rv := objc.Send[int](m.ID, objc.Sel("StringAsResult:"), result)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/clearOneofValuesForResult
func (m ModelKeyServerAPIFetchKeyResponse) ClearOneofValuesForResult() {
	objc.Send[objc.ID](m.ID, objc.Sel("clearOneofValuesForResult"))
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/copyTo:
func (m ModelKeyServerAPIFetchKeyResponse) CopyTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/dictionaryRepresentation
func (m ModelKeyServerAPIFetchKeyResponse) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/mergeFrom:
func (m ModelKeyServerAPIFetchKeyResponse) MergeFrom(from objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/readFrom:
func (m ModelKeyServerAPIFetchKeyResponse) ReadFrom(from objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/resultAsString:
func (m ModelKeyServerAPIFetchKeyResponse) ResultAsString(string_ int) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resultAsString:"), string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/writeTo:
func (m ModelKeyServerAPIFetchKeyResponse) WriteTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/error
func (m ModelKeyServerAPIFetchKeyResponse) Error() IModelKeyServerAPIResultError {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("error"))
	return ModelKeyServerAPIResultErrorFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIFetchKeyResponse) SetError(value IModelKeyServerAPIResultError) {
	objc.Send[struct{}](m.ID, objc.Sel("setError:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/hasError
func (m ModelKeyServerAPIFetchKeyResponse) HasError() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasError"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/hasResult
func (m ModelKeyServerAPIFetchKeyResponse) HasResult() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasResult"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResponse) SetHasResult(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasResult:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/hasSuccess
func (m ModelKeyServerAPIFetchKeyResponse) HasSuccess() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasSuccess"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/result
func (m ModelKeyServerAPIFetchKeyResponse) Result() int {
	rv := objc.Send[int](m.ID, objc.Sel("result"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResponse) SetResult(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setResult:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIFetchKeyResponse/success
func (m ModelKeyServerAPIFetchKeyResponse) Success() IModelKeyServerAPIFetchKeyResult {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("success"))
	return ModelKeyServerAPIFetchKeyResultFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIFetchKeyResponse) SetSuccess(value IModelKeyServerAPIFetchKeyResult) {
	objc.Send[struct{}](m.ID, objc.Sel("setSuccess:"), value)
}

