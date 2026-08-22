// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPIRawKey] class.
var (
	_ModelKeyServerAPIRawKeyClass     ModelKeyServerAPIRawKeyClass
	_ModelKeyServerAPIRawKeyClassOnce sync.Once
)

func getModelKeyServerAPIRawKeyClass() ModelKeyServerAPIRawKeyClass {
	_ModelKeyServerAPIRawKeyClassOnce.Do(func() {
		_ModelKeyServerAPIRawKeyClass = ModelKeyServerAPIRawKeyClass{class: objc.GetClass("ModelKeyServerAPIRawKey")}
	})
	return _ModelKeyServerAPIRawKeyClass
}

// GetModelKeyServerAPIRawKeyClass returns the class object for ModelKeyServerAPIRawKey.
func GetModelKeyServerAPIRawKeyClass() ModelKeyServerAPIRawKeyClass {
	return getModelKeyServerAPIRawKeyClass()
}

type ModelKeyServerAPIRawKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPIRawKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPIRawKeyClass) Alloc() ModelKeyServerAPIRawKey {
	rv := objc.SendIfResponds[ModelKeyServerAPIRawKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ModelKeyServerAPIRawKey.CopyTo]
//   - [ModelKeyServerAPIRawKey.DictionaryRepresentation]
//   - [ModelKeyServerAPIRawKey.EncryptionIv]
//   - [ModelKeyServerAPIRawKey.SetEncryptionIv]
//   - [ModelKeyServerAPIRawKey.EncryptionKey]
//   - [ModelKeyServerAPIRawKey.SetEncryptionKey]
//   - [ModelKeyServerAPIRawKey.HasEncryptionIv]
//   - [ModelKeyServerAPIRawKey.HasEncryptionKey]
//   - [ModelKeyServerAPIRawKey.MergeFrom]
//   - [ModelKeyServerAPIRawKey.ReadFrom]
//   - [ModelKeyServerAPIRawKey.WriteTo]
type ModelKeyServerAPIRawKey struct {
	objectivec.Object
}

// ModelKeyServerAPIRawKeyFromID constructs a [ModelKeyServerAPIRawKey] from an objc.ID.
func ModelKeyServerAPIRawKeyFromID(id objc.ID) ModelKeyServerAPIRawKey {
	return ModelKeyServerAPIRawKey{objectivec.Object{ID: id}}
}

// NOTE: ModelKeyServerAPIRawKey embeds objectivec.Object because the parent type is
// unavailable, but IModelKeyServerAPIRawKey embeds IPBCodable, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPIRawKey] class.
//
// # Methods
//
//   - [IModelKeyServerAPIRawKey.CopyTo]
//   - [IModelKeyServerAPIRawKey.DictionaryRepresentation]
//   - [IModelKeyServerAPIRawKey.EncryptionIv]
//   - [IModelKeyServerAPIRawKey.SetEncryptionIv]
//   - [IModelKeyServerAPIRawKey.EncryptionKey]
//   - [IModelKeyServerAPIRawKey.SetEncryptionKey]
//   - [IModelKeyServerAPIRawKey.HasEncryptionIv]
//   - [IModelKeyServerAPIRawKey.HasEncryptionKey]
//   - [IModelKeyServerAPIRawKey.MergeFrom]
//   - [IModelKeyServerAPIRawKey.ReadFrom]
//   - [IModelKeyServerAPIRawKey.WriteTo]
type IModelKeyServerAPIRawKey interface {
	IPBCodable

	// Topic: Methods

	CopyTo(to objectivec.IObject)
	DictionaryRepresentation() objectivec.IObject
	EncryptionIv() foundation.NSData
	SetEncryptionIv(value foundation.NSData)
	EncryptionKey() foundation.NSData
	SetEncryptionKey(value foundation.NSData)
	HasEncryptionIv() bool
	HasEncryptionKey() bool
	MergeFrom(from objectivec.IObject)
	ReadFrom(from objectivec.IObject) bool
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPIRawKey) Init() ModelKeyServerAPIRawKey {
	rv := objc.SendIfResponds[ModelKeyServerAPIRawKey](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPIRawKey) Autorelease() ModelKeyServerAPIRawKey {
	rv := objc.SendIfResponds[ModelKeyServerAPIRawKey](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPIRawKey creates a new ModelKeyServerAPIRawKey instance.
func NewModelKeyServerAPIRawKey() ModelKeyServerAPIRawKey {
	class := getModelKeyServerAPIRawKeyClass()
	rv := objc.SendIfResponds[ModelKeyServerAPIRawKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m ModelKeyServerAPIRawKey) CopyTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
func (m ModelKeyServerAPIRawKey) DictionaryRepresentation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (m ModelKeyServerAPIRawKey) MergeFrom(from objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
func (m ModelKeyServerAPIRawKey) ReadFrom(from objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
func (m ModelKeyServerAPIRawKey) WriteTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

func (m ModelKeyServerAPIRawKey) EncryptionIv() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encryptionIv"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIRawKey) SetEncryptionIv(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEncryptionIv:"), value)
}
func (m ModelKeyServerAPIRawKey) EncryptionKey() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encryptionKey"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIRawKey) SetEncryptionKey(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEncryptionKey:"), value)
}
func (m ModelKeyServerAPIRawKey) HasEncryptionIv() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasEncryptionIv"))
	return rv
}
func (m ModelKeyServerAPIRawKey) HasEncryptionKey() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasEncryptionKey"))
	return rv
}
