// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OS_os_workgroup_interval] class.
var (
	_OS_os_workgroup_intervalClass     OS_os_workgroup_intervalClass
	_OS_os_workgroup_intervalClassOnce sync.Once
)

func getOS_os_workgroup_intervalClass() OS_os_workgroup_intervalClass {
	_OS_os_workgroup_intervalClassOnce.Do(func() {
		_OS_os_workgroup_intervalClass = OS_os_workgroup_intervalClass{class: objc.GetClass("OS_os_workgroup_interval")}
	})
	return _OS_os_workgroup_intervalClass
}

// GetOS_os_workgroup_intervalClass returns the class object for OS_os_workgroup_interval.
func GetOS_os_workgroup_intervalClass() OS_os_workgroup_intervalClass {
	return getOS_os_workgroup_intervalClass()
}

type OS_os_workgroup_intervalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OS_os_workgroup_intervalClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OS_os_workgroup_intervalClass) Alloc() OS_os_workgroup_interval {
	rv := objc.Send[OS_os_workgroup_interval](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.class
type OS_os_workgroup_interval struct {
	OS_os_workgroup
}

// OS_os_workgroup_intervalFromID constructs a [OS_os_workgroup_interval] from an objc.ID.
func OS_os_workgroup_intervalFromID(id objc.ID) OS_os_workgroup_interval {
	return OS_os_workgroup_interval{OS_os_workgroup: OS_os_workgroupFromID(id)}
}

// NOTE: OS_os_workgroup_interval adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OS_os_workgroup_interval] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.class
type IOS_os_workgroup_interval interface {
	IOS_os_workgroup
	OS_os_workgroup_intervalProtocol
}

// Init initializes the instance.
func (o OS_os_workgroup_interval) Init() OS_os_workgroup_interval {
	rv := objc.Send[OS_os_workgroup_interval](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OS_os_workgroup_interval) Autorelease() OS_os_workgroup_interval {
	rv := objc.Send[OS_os_workgroup_interval](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_os_workgroup_interval creates a new OS_os_workgroup_interval instance.
func NewOS_os_workgroup_interval() OS_os_workgroup_interval {
	class := getOS_os_workgroup_intervalClass()
	rv := objc.Send[OS_os_workgroup_interval](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Protocol methods for OS_os_workgroup_interval
