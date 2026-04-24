// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXNotificationCenter] class.
var (
	_CPXNotificationCenterClass     CPXNotificationCenterClass
	_CPXNotificationCenterClassOnce sync.Once
)

func getCPXNotificationCenterClass() CPXNotificationCenterClass {
	_CPXNotificationCenterClassOnce.Do(func() {
		_CPXNotificationCenterClass = CPXNotificationCenterClass{class: objc.GetClass("CPXNotificationCenter")}
	})
	return _CPXNotificationCenterClass
}

// GetCPXNotificationCenterClass returns the class object for CPXNotificationCenter.
func GetCPXNotificationCenterClass() CPXNotificationCenterClass {
	return getCPXNotificationCenterClass()
}

type CPXNotificationCenterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXNotificationCenterClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXNotificationCenterClass) Alloc() CPXNotificationCenter {
	rv := objc.Send[CPXNotificationCenter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXNotificationCenter.NotifyLaunchServicesOfLastestEventTypeFlags]
//   - [CPXNotificationCenter.PostLocalNotificationDataLength]
//   - [CPXNotificationCenter.InitWithSession]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter
type CPXNotificationCenter struct {
	objectivec.Object
}

// CPXNotificationCenterFromID constructs a [CPXNotificationCenter] from an objc.ID.
func CPXNotificationCenterFromID(id objc.ID) CPXNotificationCenter {
	return CPXNotificationCenter{objectivec.Object{ID: id}}
}

// Ensure CPXNotificationCenter implements ICPXNotificationCenter.
var _ ICPXNotificationCenter = CPXNotificationCenter{}

// An interface definition for the [CPXNotificationCenter] class.
//
// # Methods
//
//   - [ICPXNotificationCenter.NotifyLaunchServicesOfLastestEventTypeFlags]
//   - [ICPXNotificationCenter.PostLocalNotificationDataLength]
//   - [ICPXNotificationCenter.InitWithSession]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter
type ICPXNotificationCenter interface {
	objectivec.IObject

	// Topic: Methods

	NotifyLaunchServicesOfLastestEventTypeFlags(type_ uint32, flags uint32)
	PostLocalNotificationDataLength(notification uint32, data unsafe.Pointer, length uint64)
	InitWithSession(session unsafe.Pointer) CPXNotificationCenter
}

// Init initializes the instance.
func (c CPXNotificationCenter) Init() CPXNotificationCenter {
	rv := objc.Send[CPXNotificationCenter](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXNotificationCenter) Autorelease() CPXNotificationCenter {
	rv := objc.Send[CPXNotificationCenter](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXNotificationCenter creates a new CPXNotificationCenter instance.
func NewCPXNotificationCenter() CPXNotificationCenter {
	class := getCPXNotificationCenterClass()
	rv := objc.Send[CPXNotificationCenter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter/initWithSession:
func NewCPXNotificationCenterWithSession(session unsafe.Pointer) CPXNotificationCenter {
	instance := getCPXNotificationCenterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return CPXNotificationCenterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter/notifyLaunchServicesOfLastestEventType:flags:
func (c CPXNotificationCenter) NotifyLaunchServicesOfLastestEventTypeFlags(type_ uint32, flags uint32) {
	objc.Send[objc.ID](c.ID, objc.Sel("notifyLaunchServicesOfLastestEventType:flags:"), type_, flags)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter/postLocalNotification:data:length:
func (c CPXNotificationCenter) PostLocalNotificationDataLength(notification uint32, data unsafe.Pointer, length uint64) {
	objc.Send[objc.ID](c.ID, objc.Sel("postLocalNotification:data:length:"), notification, data, length)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXNotificationCenter/initWithSession:
func (c CPXNotificationCenter) InitWithSession(session unsafe.Pointer) CPXNotificationCenter {
	rv := objc.Send[CPXNotificationCenter](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}
