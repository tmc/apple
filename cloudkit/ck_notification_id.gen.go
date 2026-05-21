// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKNotificationID] class.
var (
	_CKNotificationIDClass     CKNotificationIDClass
	_CKNotificationIDClassOnce sync.Once
)

func getCKNotificationIDClass() CKNotificationIDClass {
	_CKNotificationIDClassOnce.Do(func() {
		_CKNotificationIDClass = CKNotificationIDClass{class: objc.GetClass("CKNotificationID")}
	})
	return _CKNotificationIDClass
}

// GetCKNotificationIDClass returns the class object for CKNotificationID.
func GetCKNotificationIDClass() CKNotificationIDClass {
	return getCKNotificationIDClass()
}

type CKNotificationIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKNotificationIDClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKNotificationIDClass) Alloc() CKNotificationID {
	rv := objc.Send[CKNotificationID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that uniquely identifies a push notification that a container
// sends.
//
// # Overview
//
// You don’t create notification IDs directly. The server creates them when
// it creates instances of [CKNotification] that correspond to the push
// notifications that CloudKit sends to your app. You can compare two IDs
// using the [isEqual(_:)] method to determine whether two notifications are
// the same. This class defines no methods or properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/ID
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
type CKNotificationID struct {
	objectivec.Object
}

// CKNotificationIDFromID constructs a [CKNotificationID] from an objc.ID.
//
// An object that uniquely identifies a push notification that a container
// sends.
func CKNotificationIDFromID(id objc.ID) CKNotificationID {
	return CKNotificationID{objectivec.Object{ID: id}}
}

// NOTE: CKNotificationID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKNotificationID] class.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/ID
type ICKNotificationID interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKNotificationID) Init() CKNotificationID {
	rv := objc.Send[CKNotificationID](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKNotificationID) Autorelease() CKNotificationID {
	rv := objc.Send[CKNotificationID](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotificationID creates a new CKNotificationID instance.
func NewCKNotificationID() CKNotificationID {
	class := getCKNotificationIDClass()
	rv := objc.Send[CKNotificationID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKNotificationID) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
