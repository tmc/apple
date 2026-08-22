// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothUserNotification] class.
var (
	_IOBluetoothUserNotificationClass     IOBluetoothUserNotificationClass
	_IOBluetoothUserNotificationClassOnce sync.Once
)

func getIOBluetoothUserNotificationClass() IOBluetoothUserNotificationClass {
	_IOBluetoothUserNotificationClassOnce.Do(func() {
		_IOBluetoothUserNotificationClass = IOBluetoothUserNotificationClass{class: objc.GetClass("IOBluetoothUserNotification")}
	})
	return _IOBluetoothUserNotificationClass
}

// GetIOBluetoothUserNotificationClass returns the class object for IOBluetoothUserNotification.
func GetIOBluetoothUserNotificationClass() IOBluetoothUserNotificationClass {
	return getIOBluetoothUserNotificationClass()
}

type IOBluetoothUserNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothUserNotificationClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothUserNotificationClass) Alloc() IOBluetoothUserNotification {
	rv := objc.Send[IOBluetoothUserNotification](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// Represents a registered notification.
//
// # Overview
//
// When registering for various notifications in the system, an
// IOBluetoothUserNotification object is returned. To unregister from the
// notification, call -unregister on the IOBluetoothUserNotification object.
// Once -unregister is called, the object will no longer be valid.
//
// # Instance Methods
//
//   - [IOBluetoothUserNotification.Unregister]: Called to unregister the target notification.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification
type IOBluetoothUserNotification struct {
	objectivec.Object
}

// IOBluetoothUserNotificationFromID constructs a [IOBluetoothUserNotification] from an objc.ID.
//
// Represents a registered notification.
func IOBluetoothUserNotificationFromID(id objc.ID) IOBluetoothUserNotification {
	return IOBluetoothUserNotification{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothUserNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothUserNotification] class.
//
// # Instance Methods
//
//   - [IIOBluetoothUserNotification.Unregister]: Called to unregister the target notification.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification
type IIOBluetoothUserNotification interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Called to unregister the target notification.
	Unregister()
}

// Init initializes the instance.
func (b IOBluetoothUserNotification) Init() IOBluetoothUserNotification {
	rv := objc.Send[IOBluetoothUserNotification](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothUserNotification) Autorelease() IOBluetoothUserNotification {
	rv := objc.Send[IOBluetoothUserNotification](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothUserNotification creates a new IOBluetoothUserNotification instance.
func NewIOBluetoothUserNotification() IOBluetoothUserNotification {
	class := getIOBluetoothUserNotificationClass()
	rv := objc.Send[IOBluetoothUserNotification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called to unregister the target notification.
//
// # Discussion
//
// Once this method has completed, the target IOBluetoothUserNotification will
// no longer be valid.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification/unregister()
func (b IOBluetoothUserNotification) Unregister() {
	objc.Send[objc.ID](b.ID, objc.Sel("unregister"))
}
