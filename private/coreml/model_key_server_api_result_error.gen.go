// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPIResultError] class.
var (
	_ModelKeyServerAPIResultErrorClass     ModelKeyServerAPIResultErrorClass
	_ModelKeyServerAPIResultErrorClassOnce sync.Once
)

func getModelKeyServerAPIResultErrorClass() ModelKeyServerAPIResultErrorClass {
	_ModelKeyServerAPIResultErrorClassOnce.Do(func() {
		_ModelKeyServerAPIResultErrorClass = ModelKeyServerAPIResultErrorClass{class: objc.GetClass("ModelKeyServerAPIResultError")}
	})
	return _ModelKeyServerAPIResultErrorClass
}

// GetModelKeyServerAPIResultErrorClass returns the class object for ModelKeyServerAPIResultError.
func GetModelKeyServerAPIResultErrorClass() ModelKeyServerAPIResultErrorClass {
	return getModelKeyServerAPIResultErrorClass()
}

type ModelKeyServerAPIResultErrorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPIResultErrorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPIResultErrorClass) Alloc() ModelKeyServerAPIResultError {
	rv := objc.Send[ModelKeyServerAPIResultError](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ModelKeyServerAPIResultError.CopyTo]
//   - [ModelKeyServerAPIResultError.DictionaryRepresentation]
//   - [ModelKeyServerAPIResultError.HasMessage]
//   - [ModelKeyServerAPIResultError.MergeFrom]
//   - [ModelKeyServerAPIResultError.Message]
//   - [ModelKeyServerAPIResultError.SetMessage]
//   - [ModelKeyServerAPIResultError.ReadFrom]
//   - [ModelKeyServerAPIResultError.WriteTo]
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError
type ModelKeyServerAPIResultError struct {
	objectivec.Object
}

// ModelKeyServerAPIResultErrorFromID constructs a [ModelKeyServerAPIResultError] from an objc.ID.
func ModelKeyServerAPIResultErrorFromID(id objc.ID) ModelKeyServerAPIResultError {
	return ModelKeyServerAPIResultError{objectivec.Object{ID: id}}
}
// NOTE: ModelKeyServerAPIResultError struct embeds objectivec.Object (parent type unavailable) but
// IModelKeyServerAPIResultError embeds the parent interface; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPIResultError] class.
//
// # Methods
//
//   - [IModelKeyServerAPIResultError.CopyTo]
//   - [IModelKeyServerAPIResultError.DictionaryRepresentation]
//   - [IModelKeyServerAPIResultError.HasMessage]
//   - [IModelKeyServerAPIResultError.MergeFrom]
//   - [IModelKeyServerAPIResultError.Message]
//   - [IModelKeyServerAPIResultError.SetMessage]
//   - [IModelKeyServerAPIResultError.ReadFrom]
//   - [IModelKeyServerAPIResultError.WriteTo]
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError
type IModelKeyServerAPIResultError interface {
	IPBCodable

	// Topic: Methods

	CopyTo(to objectivec.IObject)
	DictionaryRepresentation() objectivec.IObject
	HasMessage() bool
	MergeFrom(from objectivec.IObject)
	Message() string
	SetMessage(value string)
	ReadFrom(from objectivec.IObject) bool
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPIResultError) Init() ModelKeyServerAPIResultError {
	rv := objc.Send[ModelKeyServerAPIResultError](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPIResultError) Autorelease() ModelKeyServerAPIResultError {
	rv := objc.Send[ModelKeyServerAPIResultError](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPIResultError creates a new ModelKeyServerAPIResultError instance.
func NewModelKeyServerAPIResultError() ModelKeyServerAPIResultError {
	class := getModelKeyServerAPIResultErrorClass()
	rv := objc.Send[ModelKeyServerAPIResultError](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/copyTo:
func (m ModelKeyServerAPIResultError) CopyTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/dictionaryRepresentation
func (m ModelKeyServerAPIResultError) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/mergeFrom:
func (m ModelKeyServerAPIResultError) MergeFrom(from objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/readFrom:
func (m ModelKeyServerAPIResultError) ReadFrom(from objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/writeTo:
func (m ModelKeyServerAPIResultError) WriteTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/hasMessage
func (m ModelKeyServerAPIResultError) HasMessage() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasMessage"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPIResultError/message
func (m ModelKeyServerAPIResultError) Message() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("message"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIResultError) SetMessage(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setMessage:"), objc.String(value))
}

