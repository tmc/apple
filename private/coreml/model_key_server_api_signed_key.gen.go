// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPISignedKey] class.
var (
	_ModelKeyServerAPISignedKeyClass     ModelKeyServerAPISignedKeyClass
	_ModelKeyServerAPISignedKeyClassOnce sync.Once
)

func getModelKeyServerAPISignedKeyClass() ModelKeyServerAPISignedKeyClass {
	_ModelKeyServerAPISignedKeyClassOnce.Do(func() {
		_ModelKeyServerAPISignedKeyClass = ModelKeyServerAPISignedKeyClass{class: objc.GetClass("ModelKeyServerAPISignedKey")}
	})
	return _ModelKeyServerAPISignedKeyClass
}

// GetModelKeyServerAPISignedKeyClass returns the class object for ModelKeyServerAPISignedKey.
func GetModelKeyServerAPISignedKeyClass() ModelKeyServerAPISignedKeyClass {
	return getModelKeyServerAPISignedKeyClass()
}

type ModelKeyServerAPISignedKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPISignedKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPISignedKeyClass) Alloc() ModelKeyServerAPISignedKey {
	rv := objc.Send[ModelKeyServerAPISignedKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ModelKeyServerAPISignedKey.CopyTo]
//   - [ModelKeyServerAPISignedKey.Data]
//   - [ModelKeyServerAPISignedKey.SetData]
//   - [ModelKeyServerAPISignedKey.DictionaryRepresentation]
//   - [ModelKeyServerAPISignedKey.HasData]
//   - [ModelKeyServerAPISignedKey.MergeFrom]
//   - [ModelKeyServerAPISignedKey.ReadFrom]
//   - [ModelKeyServerAPISignedKey.WriteTo]
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey
type ModelKeyServerAPISignedKey struct {
	objectivec.Object
}

// ModelKeyServerAPISignedKeyFromID constructs a [ModelKeyServerAPISignedKey] from an objc.ID.
func ModelKeyServerAPISignedKeyFromID(id objc.ID) ModelKeyServerAPISignedKey {
	return ModelKeyServerAPISignedKey{objectivec.Object{ID: id}}
}
// NOTE: ModelKeyServerAPISignedKey struct embeds objectivec.Object (parent type unavailable) but
// IModelKeyServerAPISignedKey embeds the parent interface; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPISignedKey] class.
//
// # Methods
//
//   - [IModelKeyServerAPISignedKey.CopyTo]
//   - [IModelKeyServerAPISignedKey.Data]
//   - [IModelKeyServerAPISignedKey.SetData]
//   - [IModelKeyServerAPISignedKey.DictionaryRepresentation]
//   - [IModelKeyServerAPISignedKey.HasData]
//   - [IModelKeyServerAPISignedKey.MergeFrom]
//   - [IModelKeyServerAPISignedKey.ReadFrom]
//   - [IModelKeyServerAPISignedKey.WriteTo]
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey
type IModelKeyServerAPISignedKey interface {
	IPBCodable

	// Topic: Methods

	CopyTo(to objectivec.IObject)
	Data() foundation.INSData
	SetData(value foundation.INSData)
	DictionaryRepresentation() objectivec.IObject
	HasData() bool
	MergeFrom(from objectivec.IObject)
	ReadFrom(from objectivec.IObject) bool
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPISignedKey) Init() ModelKeyServerAPISignedKey {
	rv := objc.Send[ModelKeyServerAPISignedKey](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPISignedKey) Autorelease() ModelKeyServerAPISignedKey {
	rv := objc.Send[ModelKeyServerAPISignedKey](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPISignedKey creates a new ModelKeyServerAPISignedKey instance.
func NewModelKeyServerAPISignedKey() ModelKeyServerAPISignedKey {
	class := getModelKeyServerAPISignedKeyClass()
	rv := objc.Send[ModelKeyServerAPISignedKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/copyTo:
func (m ModelKeyServerAPISignedKey) CopyTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/dictionaryRepresentation
func (m ModelKeyServerAPISignedKey) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/mergeFrom:
func (m ModelKeyServerAPISignedKey) MergeFrom(from objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/readFrom:
func (m ModelKeyServerAPISignedKey) ReadFrom(from objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/writeTo:
func (m ModelKeyServerAPISignedKey) WriteTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/data
func (m ModelKeyServerAPISignedKey) Data() foundation.INSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m ModelKeyServerAPISignedKey) SetData(value foundation.INSData) {
	objc.Send[struct{}](m.ID, objc.Sel("setData:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/ModelKeyServerAPISignedKey/hasData
func (m ModelKeyServerAPISignedKey) HasData() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasData"))
	return rv
}

