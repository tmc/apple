// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OSOSWorkgroupParallel] class.
var (
	_OSOSWorkgroupParallelClass     OSOSWorkgroupParallelClass
	_OSOSWorkgroupParallelClassOnce sync.Once
)

func getOSOSWorkgroupParallelClass() OSOSWorkgroupParallelClass {
	_OSOSWorkgroupParallelClassOnce.Do(func() {
		_OSOSWorkgroupParallelClass = OSOSWorkgroupParallelClass{class: objc.GetClass("OS_os_workgroup_parallel")}
	})
	return _OSOSWorkgroupParallelClass
}

// GetOSOSWorkgroupParallelClass returns the class object for OS_os_workgroup_parallel.
func GetOSOSWorkgroupParallelClass() OSOSWorkgroupParallelClass {
	return getOSOSWorkgroupParallelClass()
}

type OSOSWorkgroupParallelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSOSWorkgroupParallelClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSOSWorkgroupParallelClass) Alloc() OSOSWorkgroupParallel {
	rv := objc.Send[OSOSWorkgroupParallel](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.class
type OSOSWorkgroupParallel struct {
	OSOSWorkgroup
}

// OSOSWorkgroupParallelFromID constructs a [OSOSWorkgroupParallel] from an objc.ID.
func OSOSWorkgroupParallelFromID(id objc.ID) OSOSWorkgroupParallel {
	return OSOSWorkgroupParallel{OSOSWorkgroup: OSOSWorkgroupFromID(id)}
}

// OS_os_workgroup_parallelFromID is an alias for [OSOSWorkgroupParallelFromID] for cross-framework compatibility.
func OS_os_workgroup_parallelFromID(id objc.ID) OSOSWorkgroupParallel {
	return OSOSWorkgroupParallelFromID(id)
}

// NOTE: OSOSWorkgroupParallel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSOSWorkgroupParallel] class.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.class
type IOSOSWorkgroupParallel interface {
	IOSOSWorkgroup
	OS_os_workgroup_parallel
}

// Init initializes the instance.
func (o OSOSWorkgroupParallel) Init() OSOSWorkgroupParallel {
	rv := objc.Send[OSOSWorkgroupParallel](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSOSWorkgroupParallel) Autorelease() OSOSWorkgroupParallel {
	rv := objc.Send[OSOSWorkgroupParallel](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSOSWorkgroupParallel creates a new OSOSWorkgroupParallel instance.
func NewOSOSWorkgroupParallel() OSOSWorkgroupParallel {
	class := getOSOSWorkgroupParallelClass()
	rv := objc.Send[OSOSWorkgroupParallel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Protocol methods for OS_os_workgroup_parallel
