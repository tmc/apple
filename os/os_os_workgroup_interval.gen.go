// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OSOSWorkgroupInterval] class.
var (
	_OSOSWorkgroupIntervalClass     OSOSWorkgroupIntervalClass
	_OSOSWorkgroupIntervalClassOnce sync.Once
)

func getOSOSWorkgroupIntervalClass() OSOSWorkgroupIntervalClass {
	_OSOSWorkgroupIntervalClassOnce.Do(func() {
		_OSOSWorkgroupIntervalClass = OSOSWorkgroupIntervalClass{class: objc.GetClass("OS_os_workgroup_interval")}
	})
	return _OSOSWorkgroupIntervalClass
}

// GetOSOSWorkgroupIntervalClass returns the class object for OS_os_workgroup_interval.
func GetOSOSWorkgroupIntervalClass() OSOSWorkgroupIntervalClass {
	return getOSOSWorkgroupIntervalClass()
}

type OSOSWorkgroupIntervalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSOSWorkgroupIntervalClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSOSWorkgroupIntervalClass) Alloc() OSOSWorkgroupInterval {
	rv := objc.Send[OSOSWorkgroupInterval](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.class
type OSOSWorkgroupInterval struct {
	OSOSWorkgroup
}

// OSOSWorkgroupIntervalFromID constructs a [OSOSWorkgroupInterval] from an objc.ID.
func OSOSWorkgroupIntervalFromID(id objc.ID) OSOSWorkgroupInterval {
	return OSOSWorkgroupInterval{OSOSWorkgroup: OSOSWorkgroupFromID(id)}
}

// OS_os_workgroup_intervalFromID is an alias for [OSOSWorkgroupIntervalFromID] for cross-framework compatibility.
func OS_os_workgroup_intervalFromID(id objc.ID) OSOSWorkgroupInterval {
	return OSOSWorkgroupIntervalFromID(id)
}

// NOTE: OSOSWorkgroupInterval adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSOSWorkgroupInterval] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.class
type IOSOSWorkgroupInterval interface {
	IOSOSWorkgroup
	OS_os_workgroup_interval
}

// Init initializes the instance.
func (o OSOSWorkgroupInterval) Init() OSOSWorkgroupInterval {
	rv := objc.Send[OSOSWorkgroupInterval](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSOSWorkgroupInterval) Autorelease() OSOSWorkgroupInterval {
	rv := objc.Send[OSOSWorkgroupInterval](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSOSWorkgroupInterval creates a new OSOSWorkgroupInterval instance.
func NewOSOSWorkgroupInterval() OSOSWorkgroupInterval {
	class := getOSOSWorkgroupIntervalClass()
	rv := objc.Send[OSOSWorkgroupInterval](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Protocol methods for OS_os_workgroup_interval
