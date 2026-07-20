// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OS_os_workgroup_parallel] class.
var (
	_OS_os_workgroup_parallelClass     OS_os_workgroup_parallelClass
	_OS_os_workgroup_parallelClassOnce sync.Once
)

func getOS_os_workgroup_parallelClass() OS_os_workgroup_parallelClass {
	_OS_os_workgroup_parallelClassOnce.Do(func() {
		_OS_os_workgroup_parallelClass = OS_os_workgroup_parallelClass{class: objc.GetClass("OS_os_workgroup_parallel")}
	})
	return _OS_os_workgroup_parallelClass
}

// GetOS_os_workgroup_parallelClass returns the class object for OS_os_workgroup_parallel.
func GetOS_os_workgroup_parallelClass() OS_os_workgroup_parallelClass {
	return getOS_os_workgroup_parallelClass()
}

type OS_os_workgroup_parallelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OS_os_workgroup_parallelClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OS_os_workgroup_parallelClass) Alloc() OS_os_workgroup_parallel {
	rv := objc.Send[OS_os_workgroup_parallel](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.class
type OS_os_workgroup_parallel struct {
	OS_os_workgroup
}

// OS_os_workgroup_parallelFromID constructs a [OS_os_workgroup_parallel] from an objc.ID.
func OS_os_workgroup_parallelFromID(id objc.ID) OS_os_workgroup_parallel {
	return OS_os_workgroup_parallel{OS_os_workgroup: OS_os_workgroupFromID(id)}
}

// NOTE: OS_os_workgroup_parallel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OS_os_workgroup_parallel] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.class
type IOS_os_workgroup_parallel interface {
	IOS_os_workgroup
	OS_os_workgroup_parallelProtocol
}

// Init initializes the instance.
func (o OS_os_workgroup_parallel) Init() OS_os_workgroup_parallel {
	rv := objc.Send[OS_os_workgroup_parallel](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OS_os_workgroup_parallel) Autorelease() OS_os_workgroup_parallel {
	rv := objc.Send[OS_os_workgroup_parallel](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_os_workgroup_parallel creates a new OS_os_workgroup_parallel instance.
func NewOS_os_workgroup_parallel() OS_os_workgroup_parallel {
	class := getOS_os_workgroup_parallelClass()
	rv := objc.Send[OS_os_workgroup_parallel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Protocol methods for OS_os_workgroup_parallel
