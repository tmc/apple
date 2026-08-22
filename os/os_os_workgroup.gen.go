// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OSOSWorkgroup] class.
var (
	_OSOSWorkgroupClass     OSOSWorkgroupClass
	_OSOSWorkgroupClassOnce sync.Once
)

func getOSOSWorkgroupClass() OSOSWorkgroupClass {
	_OSOSWorkgroupClassOnce.Do(func() {
		_OSOSWorkgroupClass = OSOSWorkgroupClass{class: objc.GetClass("OS_os_workgroup")}
	})
	return _OSOSWorkgroupClass
}

// GetOSOSWorkgroupClass returns the class object for OS_os_workgroup.
func GetOSOSWorkgroupClass() OSOSWorkgroupClass {
	return getOSOSWorkgroupClass()
}

type OSOSWorkgroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSOSWorkgroupClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSOSWorkgroupClass) Alloc() OSOSWorkgroup {
	rv := objc.Send[OSOSWorkgroup](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup
type OSOSWorkgroup struct {
	OSObject
}

// OSOSWorkgroupFromID constructs a [OSOSWorkgroup] from an objc.ID.
func OSOSWorkgroupFromID(id objc.ID) OSOSWorkgroup {
	return OSOSWorkgroup{OSObject: OSObjectFromID(id)}
}

// OS_os_workgroupFromID is an alias for [OSOSWorkgroupFromID] for cross-framework compatibility.
func OS_os_workgroupFromID(id objc.ID) OSOSWorkgroup { return OSOSWorkgroupFromID(id) }

// Ensure OSOSWorkgroup implements IOSOSWorkgroup.
var _ IOSOSWorkgroup = OSOSWorkgroup{}

// An interface definition for the [OSOSWorkgroup] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup
type IOSOSWorkgroup interface {
	IOSObject
}

// Init initializes the instance.
func (o OSOSWorkgroup) Init() OSOSWorkgroup {
	rv := objc.Send[OSOSWorkgroup](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSOSWorkgroup) Autorelease() OSOSWorkgroup {
	rv := objc.Send[OSOSWorkgroup](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSOSWorkgroup creates a new OSOSWorkgroup instance.
func NewOSOSWorkgroup() OSOSWorkgroup {
	class := getOSOSWorkgroupClass()
	rv := objc.Send[OSOSWorkgroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}
