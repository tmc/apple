// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPowerSourceDevice] class.
var (
	_VZPowerSourceDeviceClass     VZPowerSourceDeviceClass
	_VZPowerSourceDeviceClassOnce sync.Once
)

func getVZPowerSourceDeviceClass() VZPowerSourceDeviceClass {
	_VZPowerSourceDeviceClassOnce.Do(func() {
		_VZPowerSourceDeviceClass = VZPowerSourceDeviceClass{class: objc.GetClass("_VZPowerSourceDevice")}
	})
	return _VZPowerSourceDeviceClass
}

// GetVZPowerSourceDeviceClass returns the class object for _VZPowerSourceDevice.
func GetVZPowerSourceDeviceClass() VZPowerSourceDeviceClass {
	return getVZPowerSourceDeviceClass()
}

type VZPowerSourceDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPowerSourceDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPowerSourceDeviceClass) Alloc() VZPowerSourceDevice {
	rv := objc.Send[VZPowerSourceDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZPowerSourceDevice._init]
// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDevice
type VZPowerSourceDevice struct {
	objectivec.Object
}

// VZPowerSourceDeviceFromID constructs a [VZPowerSourceDevice] from an objc.ID.
func VZPowerSourceDeviceFromID(id objc.ID) VZPowerSourceDevice {
	return VZPowerSourceDevice{objectivec.Object{ID: id}}
}
// Ensure VZPowerSourceDevice implements IVZPowerSourceDevice.
var _ IVZPowerSourceDevice = VZPowerSourceDevice{}

// An interface definition for the [VZPowerSourceDevice] class.
//
// # Methods
//
//   - [IVZPowerSourceDevice._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDevice
type IVZPowerSourceDevice interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZPowerSourceDevice) Init() VZPowerSourceDevice {
	rv := objc.Send[VZPowerSourceDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPowerSourceDevice) Autorelease() VZPowerSourceDevice {
	rv := objc.Send[VZPowerSourceDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPowerSourceDevice creates a new VZPowerSourceDevice instance.
func NewVZPowerSourceDevice() VZPowerSourceDevice {
	class := getVZPowerSourceDeviceClass()
	rv := objc.Send[VZPowerSourceDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDevice/_init
func (v VZPowerSourceDevice) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

