// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSObject] class.
var (
	_OSObjectClass     OSObjectClass
	_OSObjectClassOnce sync.Once
)

func getOSObjectClass() OSObjectClass {
	_OSObjectClassOnce.Do(func() {
		_OSObjectClass = OSObjectClass{class: objc.GetClass("OS_object")}
	})
	return _OSObjectClass
}

// GetOSObjectClass returns the class object for OS_object.
func GetOSObjectClass() OSObjectClass {
	return getOSObjectClass()
}

type OSObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSObjectClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSObjectClass) Alloc() OSObject {
	rv := objc.Send[OSObject](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_object
type OSObject struct {
	objectivec.Object
}

// OSObjectFromID constructs a [OSObject] from an objc.ID.
func OSObjectFromID(id objc.ID) OSObject {
	return OSObject{objectivec.Object{ID: id}}
}

// OS_objectFromID is an alias for [OSObjectFromID] for cross-framework compatibility.
func OS_objectFromID(id objc.ID) OSObject { return OSObjectFromID(id) }

// Ensure OSObject implements IOSObject.
var _ IOSObject = OSObject{}

// An interface definition for the [OSObject] class.
//
// See: https://developer.apple.com/documentation/os/OS_object
type IOSObject interface {
	objectivec.IObject
}

// Init initializes the instance.
func (o OSObject) Init() OSObject {
	rv := objc.Send[OSObject](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSObject) Autorelease() OSObject {
	rv := objc.Send[OSObject](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSObject creates a new OSObject instance.
func NewOSObject() OSObject {
	class := getOSObjectClass()
	rv := objc.Send[OSObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}
