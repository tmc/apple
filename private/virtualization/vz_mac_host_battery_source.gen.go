// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacHostBatterySource] class.
var (
	_VZMacHostBatterySourceClass     VZMacHostBatterySourceClass
	_VZMacHostBatterySourceClassOnce sync.Once
)

func getVZMacHostBatterySourceClass() VZMacHostBatterySourceClass {
	_VZMacHostBatterySourceClassOnce.Do(func() {
		_VZMacHostBatterySourceClass = VZMacHostBatterySourceClass{class: objc.GetClass("_VZMacHostBatterySource")}
	})
	return _VZMacHostBatterySourceClass
}

// GetVZMacHostBatterySourceClass returns the class object for _VZMacHostBatterySource.
func GetVZMacHostBatterySourceClass() VZMacHostBatterySourceClass {
	return getVZMacHostBatterySourceClass()
}

type VZMacHostBatterySourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacHostBatterySourceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacHostBatterySourceClass) Alloc() VZMacHostBatterySource {
	rv := objc.Send[VZMacHostBatterySource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHostBatterySource
type VZMacHostBatterySource struct {
	VZMacBatterySource
}

// VZMacHostBatterySourceFromID constructs a [VZMacHostBatterySource] from an objc.ID.
func VZMacHostBatterySourceFromID(id objc.ID) VZMacHostBatterySource {
	return VZMacHostBatterySource{VZMacBatterySource: VZMacBatterySourceFromID(id)}
}

// Ensure VZMacHostBatterySource implements IVZMacHostBatterySource.
var _ IVZMacHostBatterySource = VZMacHostBatterySource{}

// An interface definition for the [VZMacHostBatterySource] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHostBatterySource
type IVZMacHostBatterySource interface {
	IVZMacBatterySource
}

// Init initializes the instance.
func (v VZMacHostBatterySource) Init() VZMacHostBatterySource {
	rv := objc.Send[VZMacHostBatterySource](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacHostBatterySource) Autorelease() VZMacHostBatterySource {
	rv := objc.Send[VZMacHostBatterySource](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHostBatterySource creates a new VZMacHostBatterySource instance.
func NewVZMacHostBatterySource() VZMacHostBatterySource {
	class := getVZMacHostBatterySourceClass()
	rv := objc.Send[VZMacHostBatterySource](objc.ID(class.class), objc.Sel("new"))
	return rv
}
