// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineStateSerialization] class.
var (
	_CKSyncEngineStateSerializationClass     CKSyncEngineStateSerializationClass
	_CKSyncEngineStateSerializationClassOnce sync.Once
)

func getCKSyncEngineStateSerializationClass() CKSyncEngineStateSerializationClass {
	_CKSyncEngineStateSerializationClassOnce.Do(func() {
		_CKSyncEngineStateSerializationClass = CKSyncEngineStateSerializationClass{class: objc.GetClass("CKSyncEngineStateSerialization")}
	})
	return _CKSyncEngineStateSerializationClass
}

// GetCKSyncEngineStateSerializationClass returns the class object for CKSyncEngineStateSerialization.
func GetCKSyncEngineStateSerializationClass() CKSyncEngineStateSerializationClass {
	return getCKSyncEngineStateSerializationClass()
}

type CKSyncEngineStateSerializationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineStateSerializationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineStateSerializationClass) Alloc() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An opaque object that contains the serialized representation of a sync
// engine’s current state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateSerialization
type CKSyncEngineStateSerialization struct {
	objectivec.Object
}

// CKSyncEngineStateSerializationFromID constructs a [CKSyncEngineStateSerialization] from an objc.ID.
//
// An opaque object that contains the serialized representation of a sync
// engine’s current state.
func CKSyncEngineStateSerializationFromID(id objc.ID) CKSyncEngineStateSerialization {
	return CKSyncEngineStateSerialization{objectivec.Object{ID: id}}
}

// NOTE: CKSyncEngineStateSerialization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSyncEngineStateSerialization] class.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateSerialization
type ICKSyncEngineStateSerialization interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKSyncEngineStateSerialization) Init() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineStateSerialization) Autorelease() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateSerialization creates a new CKSyncEngineStateSerialization instance.
func NewCKSyncEngineStateSerialization() CKSyncEngineStateSerialization {
	class := getCKSyncEngineStateSerializationClass()
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKSyncEngineStateSerialization) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
