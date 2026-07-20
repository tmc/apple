// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OS_os_workgroup] class.
var (
	_OS_os_workgroupClass     OS_os_workgroupClass
	_OS_os_workgroupClassOnce sync.Once
)

func getOS_os_workgroupClass() OS_os_workgroupClass {
	_OS_os_workgroupClassOnce.Do(func() {
		_OS_os_workgroupClass = OS_os_workgroupClass{class: objc.GetClass("OS_os_workgroup")}
	})
	return _OS_os_workgroupClass
}

// GetOS_os_workgroupClass returns the class object for OS_os_workgroup.
func GetOS_os_workgroupClass() OS_os_workgroupClass {
	return getOS_os_workgroupClass()
}

type OS_os_workgroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OS_os_workgroupClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OS_os_workgroupClass) Alloc() OS_os_workgroup {
	rv := objc.Send[OS_os_workgroup](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup
type OS_os_workgroup struct {
	OS_object
}

// OS_os_workgroupFromID constructs a [OS_os_workgroup] from an objc.ID.
func OS_os_workgroupFromID(id objc.ID) OS_os_workgroup {
	return OS_os_workgroup{OS_object: OS_objectFromID(id)}
}

// Ensure OS_os_workgroup implements IOS_os_workgroup.
var _ IOS_os_workgroup = OS_os_workgroup{}

// An interface definition for the [OS_os_workgroup] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup
type IOS_os_workgroup interface {
	IOS_object
}

// Init initializes the instance.
func (o OS_os_workgroup) Init() OS_os_workgroup {
	rv := objc.Send[OS_os_workgroup](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OS_os_workgroup) Autorelease() OS_os_workgroup {
	rv := objc.Send[OS_os_workgroup](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_os_workgroup creates a new OS_os_workgroup instance.
func NewOS_os_workgroup() OS_os_workgroup {
	class := getOS_os_workgroupClass()
	rv := objc.Send[OS_os_workgroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}
