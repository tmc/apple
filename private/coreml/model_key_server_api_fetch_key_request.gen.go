// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ModelKeyServerAPIFetchKeyRequest] class.
var (
	_ModelKeyServerAPIFetchKeyRequestClass     ModelKeyServerAPIFetchKeyRequestClass
	_ModelKeyServerAPIFetchKeyRequestClassOnce sync.Once
)

func getModelKeyServerAPIFetchKeyRequestClass() ModelKeyServerAPIFetchKeyRequestClass {
	_ModelKeyServerAPIFetchKeyRequestClassOnce.Do(func() {
		_ModelKeyServerAPIFetchKeyRequestClass = ModelKeyServerAPIFetchKeyRequestClass{class: objc.GetClass("ModelKeyServerAPIFetchKeyRequest")}
	})
	return _ModelKeyServerAPIFetchKeyRequestClass
}

// GetModelKeyServerAPIFetchKeyRequestClass returns the class object for ModelKeyServerAPIFetchKeyRequest.
func GetModelKeyServerAPIFetchKeyRequestClass() ModelKeyServerAPIFetchKeyRequestClass {
	return getModelKeyServerAPIFetchKeyRequestClass()
}

type ModelKeyServerAPIFetchKeyRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ModelKeyServerAPIFetchKeyRequestClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ModelKeyServerAPIFetchKeyRequestClass) Alloc() ModelKeyServerAPIFetchKeyRequest {
	rv := objc.Send[ModelKeyServerAPIFetchKeyRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ModelKeyServerAPIFetchKeyRequest.CopyTo]
//   - [ModelKeyServerAPIFetchKeyRequest.DictionaryRepresentation]
//   - [ModelKeyServerAPIFetchKeyRequest.HasKeyId]
//   - [ModelKeyServerAPIFetchKeyRequest.HasRawRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.SetHasRawRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.HasSignedKeyRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.HasTeamId]
//   - [ModelKeyServerAPIFetchKeyRequest.KeyId]
//   - [ModelKeyServerAPIFetchKeyRequest.SetKeyId]
//   - [ModelKeyServerAPIFetchKeyRequest.MergeFrom]
//   - [ModelKeyServerAPIFetchKeyRequest.RawRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.SetRawRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.ReadFrom]
//   - [ModelKeyServerAPIFetchKeyRequest.SignedKeyRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.SetSignedKeyRequest]
//   - [ModelKeyServerAPIFetchKeyRequest.TeamId]
//   - [ModelKeyServerAPIFetchKeyRequest.SetTeamId]
//   - [ModelKeyServerAPIFetchKeyRequest.WriteTo]
type ModelKeyServerAPIFetchKeyRequest struct {
	objectivec.Object
}

// ModelKeyServerAPIFetchKeyRequestFromID constructs a [ModelKeyServerAPIFetchKeyRequest] from an objc.ID.
func ModelKeyServerAPIFetchKeyRequestFromID(id objc.ID) ModelKeyServerAPIFetchKeyRequest {
	return ModelKeyServerAPIFetchKeyRequest{objectivec.Object{ID: id}}
}

// NOTE: ModelKeyServerAPIFetchKeyRequest struct embeds objectivec.Object (parent type unavailable) but
// IModelKeyServerAPIFetchKeyRequest embeds the parent interface; skip compile-time assertion.

// An interface definition for the [ModelKeyServerAPIFetchKeyRequest] class.
//
// # Methods
//
//   - [IModelKeyServerAPIFetchKeyRequest.CopyTo]
//   - [IModelKeyServerAPIFetchKeyRequest.DictionaryRepresentation]
//   - [IModelKeyServerAPIFetchKeyRequest.HasKeyId]
//   - [IModelKeyServerAPIFetchKeyRequest.HasRawRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.SetHasRawRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.HasSignedKeyRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.HasTeamId]
//   - [IModelKeyServerAPIFetchKeyRequest.KeyId]
//   - [IModelKeyServerAPIFetchKeyRequest.SetKeyId]
//   - [IModelKeyServerAPIFetchKeyRequest.MergeFrom]
//   - [IModelKeyServerAPIFetchKeyRequest.RawRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.SetRawRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.ReadFrom]
//   - [IModelKeyServerAPIFetchKeyRequest.SignedKeyRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.SetSignedKeyRequest]
//   - [IModelKeyServerAPIFetchKeyRequest.TeamId]
//   - [IModelKeyServerAPIFetchKeyRequest.SetTeamId]
//   - [IModelKeyServerAPIFetchKeyRequest.WriteTo]
type IModelKeyServerAPIFetchKeyRequest interface {
	IPBRequest

	// Topic: Methods

	CopyTo(to objectivec.IObject)
	DictionaryRepresentation() objectivec.IObject
	HasKeyId() bool
	HasRawRequest() bool
	SetHasRawRequest(value bool)
	HasSignedKeyRequest() bool
	HasTeamId() bool
	KeyId() string
	SetKeyId(value string)
	MergeFrom(from objectivec.IObject)
	RawRequest() bool
	SetRawRequest(value bool)
	ReadFrom(from objectivec.IObject) bool
	SignedKeyRequest() foundation.NSData
	SetSignedKeyRequest(value foundation.NSData)
	TeamId() string
	SetTeamId(value string)
	WriteTo(to objectivec.IObject)
}

// Init initializes the instance.
func (m ModelKeyServerAPIFetchKeyRequest) Init() ModelKeyServerAPIFetchKeyRequest {
	rv := objc.Send[ModelKeyServerAPIFetchKeyRequest](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m ModelKeyServerAPIFetchKeyRequest) Autorelease() ModelKeyServerAPIFetchKeyRequest {
	rv := objc.Send[ModelKeyServerAPIFetchKeyRequest](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelKeyServerAPIFetchKeyRequest creates a new ModelKeyServerAPIFetchKeyRequest instance.
func NewModelKeyServerAPIFetchKeyRequest() ModelKeyServerAPIFetchKeyRequest {
	class := getModelKeyServerAPIFetchKeyRequestClass()
	rv := objc.Send[ModelKeyServerAPIFetchKeyRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m ModelKeyServerAPIFetchKeyRequest) CopyTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyTo:"), to)
}
func (m ModelKeyServerAPIFetchKeyRequest) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (m ModelKeyServerAPIFetchKeyRequest) MergeFrom(from objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("mergeFrom:"), from)
}
func (m ModelKeyServerAPIFetchKeyRequest) ReadFrom(from objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("readFrom:"), from)
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) WriteTo(to objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("writeTo:"), to)
}

func (m ModelKeyServerAPIFetchKeyRequest) HasKeyId() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasKeyId"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) HasRawRequest() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasRawRequest"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) SetHasRawRequest(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasRawRequest:"), value)
}
func (m ModelKeyServerAPIFetchKeyRequest) HasSignedKeyRequest() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasSignedKeyRequest"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) HasTeamId() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasTeamId"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) KeyId() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyId"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIFetchKeyRequest) SetKeyId(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setKeyId:"), objc.String(value))
}
func (m ModelKeyServerAPIFetchKeyRequest) RawRequest() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("rawRequest"))
	return rv
}
func (m ModelKeyServerAPIFetchKeyRequest) SetRawRequest(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRawRequest:"), value)
}
func (m ModelKeyServerAPIFetchKeyRequest) SignedKeyRequest() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("signedKeyRequest"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m ModelKeyServerAPIFetchKeyRequest) SetSignedKeyRequest(value foundation.NSData) {
	objc.Send[struct{}](m.ID, objc.Sel("setSignedKeyRequest:"), value)
}
func (m ModelKeyServerAPIFetchKeyRequest) TeamId() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("teamId"))
	return foundation.NSStringFromID(rv).String()
}
func (m ModelKeyServerAPIFetchKeyRequest) SetTeamId(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setTeamId:"), objc.String(value))
}
