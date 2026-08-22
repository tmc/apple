// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPIFetchKeyResult] class.
var (
	_ModelKeyServerAPIFetchKeyResultClass     ModelKeyServerAPIFetchKeyResultClass
	_ModelKeyServerAPIFetchKeyResultClassOnce sync.Once
)

func getModelKeyServerAPIFetchKeyResultClass() ModelKeyServerAPIFetchKeyResultClass {
	_ModelKeyServerAPIFetchKeyResultClassOnce.Do(func() {
		_ModelKeyServerAPIFetchKeyResultClass = ModelKeyServerAPIFetchKeyResultClass{class: objc.GetClass("ModelKeyServerAPIFetchKeyResult")}
	})
	return _ModelKeyServerAPIFetchKeyResultClass
}

// GetModelKeyServerAPIFetchKeyResultClass returns the class object for ModelKeyServerAPIFetchKeyResult.
func GetModelKeyServerAPIFetchKeyResultClass() ModelKeyServerAPIFetchKeyResultClass {
	return getModelKeyServerAPIFetchKeyResultClass()
}

type ModelKeyServerAPIFetchKeyResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPIFetchKeyResultClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPIFetchKeyResultClass) Alloc() ModelKeyServerAPIFetchKeyResult {
	rv := objc.SendIfResponds[ModelKeyServerAPIFetchKeyResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ModelKeyServerAPIFetchKeyResult.StringAsKey]
//   - [ModelKeyServerAPIFetchKeyResult.ClearOneofValuesForKey]
//   - [ModelKeyServerAPIFetchKeyResult.CopyTo]
//   - [ModelKeyServerAPIFetchKeyResult.DictionaryRepresentation]
//   - [ModelKeyServerAPIFetchKeyResult.HasKey]
//   - [ModelKeyServerAPIFetchKeyResult.SetHasKey]
//   - [ModelKeyServerAPIFetchKeyResult.HasKeyId]
//   - [ModelKeyServerAPIFetchKeyResult.HasModelName]
//   - [ModelKeyServerAPIFetchKeyResult.HasRawKey]
//   - [ModelKeyServerAPIFetchKeyResult.HasSignedKey]
//   - [ModelKeyServerAPIFetchKeyResult.HasTeamId]
//   - [ModelKeyServerAPIFetchKeyResult.Key]
//   - [ModelKeyServerAPIFetchKeyResult.SetKey]
//   - [ModelKeyServerAPIFetchKeyResult.KeyAsString]
//   - [ModelKeyServerAPIFetchKeyResult.KeyId]
//   - [ModelKeyServerAPIFetchKeyResult.SetKeyId]
//   - [ModelKeyServerAPIFetchKeyResult.MergeFrom]
//   - [ModelKeyServerAPIFetchKeyResult.ModelName]
//   - [ModelKeyServerAPIFetchKeyResult.SetModelName]
//   - [ModelKeyServerAPIFetchKeyResult.RawKey]
//   - [ModelKeyServerAPIFetchKeyResult.SetRawKey]
//   - [ModelKeyServerAPIFetchKeyResult.ReadFrom]
//   - [ModelKeyServerAPIFetchKeyResult.SignedKey]
//   - [ModelKeyServerAPIFetchKeyResult.SetSignedKey]
//   - [ModelKeyServerAPIFetchKeyResult.TeamId]
//   - [ModelKeyServerAPIFetchKeyResult.SetTeamId]
//   - [ModelKeyServerAPIFetchKeyResult.WriteTo]
type ModelKeyServerAPIFetchKeyResult struct {
	objectivec.Object
}

// ModelKeyServerAPIFetchKeyResultFromID constructs a [ModelKeyServerAPIFetchKeyResult] from an objc.ID.
func ModelKeyServerAPIFetchKeyResultFromID(id objc.ID) ModelKeyServerAPIFetchKeyResult {
	return ModelKeyServerAPIFetchKeyResult{objectivec.Object{ID: id}}
}

// NOTE: ModelKeyServerAPIFetchKeyResult embeds objectivec.Object because the parent type is
// unavailable, but IModelKeyServerAPIFetchKeyResult embeds IPBCodable, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPIFetchKeyResult] class.
//
// # Methods
//
//   - [IModelKeyServerAPIFetchKeyResult.StringAsKey]
//   - [IModelKeyServerAPIFetchKeyResult.ClearOneofValuesForKey]
//   - [IModelKeyServerAPIFetchKeyResult.CopyTo]
//   - [IModelKeyServerAPIFetchKeyResult.DictionaryRepresentation]
//   - [IModelKeyServerAPIFetchKeyResult.HasKey]
//   - [IModelKeyServerAPIFetchKeyResult.SetHasKey]
//   - [IModelKeyServerAPIFetchKeyResult.HasKeyId]
//   - [IModelKeyServerAPIFetchKeyResult.HasModelName]
//   - [IModelKeyServerAPIFetchKeyResult.HasRawKey]
//   - [IModelKeyServerAPIFetchKeyResult.HasSignedKey]
//   - [IModelKeyServerAPIFetchKeyResult.HasTeamId]
//   - [IModelKeyServerAPIFetchKeyResult.Key]
//   - [IModelKeyServerAPIFetchKeyResult.SetKey]
//   - [IModelKeyServerAPIFetchKeyResult.KeyAsString]
//   - [IModelKeyServerAPIFetchKeyResult.KeyId]
//   - [IModelKeyServerAPIFetchKeyResult.SetKeyId]
//   - [IModelKeyServerAPIFetchKeyResult.MergeFrom]
//   - [IModelKeyServerAPIFetchKeyResult.ModelName]
//   - [IModelKeyServerAPIFetchKeyResult.SetModelName]
//   - [IModelKeyServerAPIFetchKeyResult.RawKey]
//   - [IModelKeyServerAPIFetchKeyResult.SetRawKey]
//   - [IModelKeyServerAPIFetchKeyResult.ReadFrom]
//   - [IModelKeyServerAPIFetchKeyResult.SignedKey]
//   - [IModelKeyServerAPIFetchKeyResult.SetSignedKey]
//   - [IModelKeyServerAPIFetchKeyResult.TeamId]
//   - [IModelKeyServerAPIFetchKeyResult.SetTeamId]
//   - [IModelKeyServerAPIFetchKeyResult.WriteTo]
type IModelKeyServerAPIFetchKeyResult interface {
	IPBCodable

	// Topic: Methods

	StringAsKey(key objectivec.IObject) int
	ClearOneofValuesForKey()
	CopyTo(to objectivec.IObject)
	DictionaryRepresentation() objectivec.IObject
	HasKey() bool
	SetHasKey(value bool)
	HasKeyId() bool
	HasModelName() bool
	HasRawKey() bool
	HasSignedKey() bool
	HasTeamId() bool
	Key() int
	SetKey(value int)
	KeyAsString(string_ int) objectivec.IObject
	KeyId() string
	SetKeyId(value string)
	MergeFrom(from objectivec.IObject)
	ModelName() string
	SetModelName(value string)
	RawKey() IModelKeyServerAPIRawKey
	SetRawKey(value IModelKeyServerAPIRawKey)
	ReadFrom(from objectivec.IObject) bool
	SignedKey() IModelKeyServerAPISignedKey
	SetSignedKey(value IModelKeyServerAPISignedKey)
	TeamId() string
	SetTeamId(value string)
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPIFetchKeyResult) Init() ModelKeyServerAPIFetchKeyResult {
	rv := objc.SendIfResponds[ModelKeyServerAPIFetchKeyResult](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPIFetchKeyResult) Autorelease() ModelKeyServerAPIFetchKeyResult {
	rv := objc.SendIfResponds[ModelKeyServerAPIFetchKeyResult](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPIFetchKeyResult creates a new ModelKeyServerAPIFetchKeyResult instance.
func NewModelKeyServerAPIFetchKeyResult() ModelKeyServerAPIFetchKeyResult {
	class := getModelKeyServerAPIFetchKeyResultClass()
	rv := objc.SendIfResponds[ModelKeyServerAPIFetchKeyResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m ModelKeyServerAPIFetchKeyResult) StringAsKey(key objectivec.IObject) int {
	rv := objc.SendIfResponds[int](m.ID, objc.Sel("StringAsKey:"), key)
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) ClearOneofValuesForKey() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("clearOneofValuesForKey"))
}
func (m ModelKeyServerAPIFetchKeyResult) CopyTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
func (m ModelKeyServerAPIFetchKeyResult) DictionaryRepresentation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (m ModelKeyServerAPIFetchKeyResult) KeyAsString(string_ int) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("keyAsString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m ModelKeyServerAPIFetchKeyResult) MergeFrom(from objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
func (m ModelKeyServerAPIFetchKeyResult) ReadFrom(from objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) WriteTo(to objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

func (m ModelKeyServerAPIFetchKeyResult) HasKey() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasKey"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) SetHasKey(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setHasKey:"), value)
}
func (m ModelKeyServerAPIFetchKeyResult) HasKeyId() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasKeyId"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) HasModelName() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasModelName"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) HasRawKey() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasRawKey"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) HasSignedKey() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasSignedKey"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) HasTeamId() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasTeamId"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) Key() int {
	rv := objc.SendIfResponds[int](m.ID, objc.Sel("key"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyResult) SetKey(value int) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setKey:"), value)
}
func (m ModelKeyServerAPIFetchKeyResult) KeyId() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("keyId"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIFetchKeyResult) SetKeyId(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setKeyId:"), objc.String(value))
}
func (m ModelKeyServerAPIFetchKeyResult) ModelName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIFetchKeyResult) SetModelName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelName:"), objc.String(value))
}
func (m ModelKeyServerAPIFetchKeyResult) RawKey() IModelKeyServerAPIRawKey {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("rawKey"))
	return ModelKeyServerAPIRawKeyFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIFetchKeyResult) SetRawKey(value IModelKeyServerAPIRawKey) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setRawKey:"), value)
}
func (m ModelKeyServerAPIFetchKeyResult) SignedKey() IModelKeyServerAPISignedKey {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("signedKey"))
	return ModelKeyServerAPISignedKeyFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIFetchKeyResult) SetSignedKey(value IModelKeyServerAPISignedKey) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSignedKey:"), value)
}
func (m ModelKeyServerAPIFetchKeyResult) TeamId() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("teamId"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIFetchKeyResult) SetTeamId(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTeamId:"), objc.String(value))
}
