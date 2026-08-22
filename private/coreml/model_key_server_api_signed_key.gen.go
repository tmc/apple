// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[ModelKeyServerAPISignedKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
type ModelKeyServerAPISignedKey struct {
	objectivec.Object
}

// ModelKeyServerAPISignedKeyFromID constructs a [ModelKeyServerAPISignedKey] from an objc.ID.
func ModelKeyServerAPISignedKeyFromID(id objc.ID) ModelKeyServerAPISignedKey {
	return ModelKeyServerAPISignedKey{objectivec.Object{ID: id}}
}

// NOTE: ModelKeyServerAPISignedKey embeds objectivec.Object because the parent type is
// unavailable, but IModelKeyServerAPISignedKey embeds IPBCodable, which that fallback
// cannot satisfy; skip compile-time assertion.

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
type IModelKeyServerAPISignedKey interface {
	IPBCodable

	// Topic: Methods

	CopyTo(to objectivec.IObject)
	Data() foundation.NSData
	SetData(value foundation.NSData)
	DictionaryRepresentation() objectivec.IObject
	HasData() bool
	MergeFrom(from objectivec.IObject)
	ReadFrom(from objectivec.IObject) bool
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPISignedKey) Init() ModelKeyServerAPISignedKey {
	rv := objc.SendIfResponds[ModelKeyServerAPISignedKey](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPISignedKey) Autorelease() ModelKeyServerAPISignedKey {
	rv := objc.SendIfResponds[ModelKeyServerAPISignedKey](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPISignedKey creates a new ModelKeyServerAPISignedKey instance.
func NewModelKeyServerAPISignedKey() ModelKeyServerAPISignedKey {
	class := getModelKeyServerAPISignedKeyClass()
	rv := objc.SendIfResponds[ModelKeyServerAPISignedKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m ModelKeyServerAPISignedKey) CopyTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
func (m ModelKeyServerAPISignedKey) DictionaryRepresentation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (m ModelKeyServerAPISignedKey) MergeFrom(from objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
func (m ModelKeyServerAPISignedKey) ReadFrom(from objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
func (m ModelKeyServerAPISignedKey) WriteTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

func (m ModelKeyServerAPISignedKey) Data() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m ModelKeyServerAPISignedKey) SetData(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setData:"), value)
}
func (m ModelKeyServerAPISignedKey) HasData() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasData"))
	return rv
}
