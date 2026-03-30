// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZHIDEventMonitor] class.
var (
	_VZHIDEventMonitorClass     VZHIDEventMonitorClass
	_VZHIDEventMonitorClassOnce sync.Once
)

func getVZHIDEventMonitorClass() VZHIDEventMonitorClass {
	_VZHIDEventMonitorClassOnce.Do(func() {
		_VZHIDEventMonitorClass = VZHIDEventMonitorClass{class: objc.GetClass("_VZHIDEventMonitor")}
	})
	return _VZHIDEventMonitorClass
}

// GetVZHIDEventMonitorClass returns the class object for _VZHIDEventMonitor.
func GetVZHIDEventMonitorClass() VZHIDEventMonitorClass {
	return getVZHIDEventMonitorClass()
}

type VZHIDEventMonitorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHIDEventMonitorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHIDEventMonitorClass) Alloc() VZHIDEventMonitor {
	rv := objc.Send[VZHIDEventMonitor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventMonitor
type VZHIDEventMonitor struct {
	objectivec.Object
}

// VZHIDEventMonitorFromID constructs a [VZHIDEventMonitor] from an objc.ID.
func VZHIDEventMonitorFromID(id objc.ID) VZHIDEventMonitor {
	return VZHIDEventMonitor{objectivec.Object{ID: id}}
}

// Ensure VZHIDEventMonitor implements IVZHIDEventMonitor.
var _ IVZHIDEventMonitor = VZHIDEventMonitor{}

// An interface definition for the [VZHIDEventMonitor] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventMonitor
type IVZHIDEventMonitor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZHIDEventMonitor) Init() VZHIDEventMonitor {
	rv := objc.Send[VZHIDEventMonitor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHIDEventMonitor) Autorelease() VZHIDEventMonitor {
	rv := objc.Send[VZHIDEventMonitor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHIDEventMonitor creates a new VZHIDEventMonitor instance.
func NewVZHIDEventMonitor() VZHIDEventMonitor {
	class := getVZHIDEventMonitorClass()
	rv := objc.Send[VZHIDEventMonitor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
