// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacBatteryPowerSourceDevice] class.
var (
	_VZMacBatteryPowerSourceDeviceClass     VZMacBatteryPowerSourceDeviceClass
	_VZMacBatteryPowerSourceDeviceClassOnce sync.Once
)

func getVZMacBatteryPowerSourceDeviceClass() VZMacBatteryPowerSourceDeviceClass {
	_VZMacBatteryPowerSourceDeviceClassOnce.Do(func() {
		_VZMacBatteryPowerSourceDeviceClass = VZMacBatteryPowerSourceDeviceClass{class: objc.GetClass("_VZMacBatteryPowerSourceDevice")}
	})
	return _VZMacBatteryPowerSourceDeviceClass
}

// GetVZMacBatteryPowerSourceDeviceClass returns the class object for _VZMacBatteryPowerSourceDevice.
func GetVZMacBatteryPowerSourceDeviceClass() VZMacBatteryPowerSourceDeviceClass {
	return getVZMacBatteryPowerSourceDeviceClass()
}

type VZMacBatteryPowerSourceDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacBatteryPowerSourceDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacBatteryPowerSourceDeviceClass) Alloc() VZMacBatteryPowerSourceDevice {
	rv := objc.Send[VZMacBatteryPowerSourceDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacBatteryPowerSourceDevice.BatterySourceDidUpdateCharge]
//   - [VZMacBatteryPowerSourceDevice.BatterySourceDidUpdateConnectivity]
//   - [VZMacBatteryPowerSourceDevice.Source]
//   - [VZMacBatteryPowerSourceDevice.SetSource]
//   - [VZMacBatteryPowerSourceDevice.DebugDescription]
//   - [VZMacBatteryPowerSourceDevice.Description]
//   - [VZMacBatteryPowerSourceDevice.Hash]
//   - [VZMacBatteryPowerSourceDevice.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice
type VZMacBatteryPowerSourceDevice struct {
	VZPowerSourceDevice
}

// VZMacBatteryPowerSourceDeviceFromID constructs a [VZMacBatteryPowerSourceDevice] from an objc.ID.
func VZMacBatteryPowerSourceDeviceFromID(id objc.ID) VZMacBatteryPowerSourceDevice {
	return VZMacBatteryPowerSourceDevice{VZPowerSourceDevice: VZPowerSourceDeviceFromID(id)}
}

// Ensure VZMacBatteryPowerSourceDevice implements IVZMacBatteryPowerSourceDevice.
var _ IVZMacBatteryPowerSourceDevice = VZMacBatteryPowerSourceDevice{}

// An interface definition for the [VZMacBatteryPowerSourceDevice] class.
//
// # Methods
//
//   - [IVZMacBatteryPowerSourceDevice.BatterySourceDidUpdateCharge]
//   - [IVZMacBatteryPowerSourceDevice.BatterySourceDidUpdateConnectivity]
//   - [IVZMacBatteryPowerSourceDevice.Source]
//   - [IVZMacBatteryPowerSourceDevice.SetSource]
//   - [IVZMacBatteryPowerSourceDevice.DebugDescription]
//   - [IVZMacBatteryPowerSourceDevice.Description]
//   - [IVZMacBatteryPowerSourceDevice.Hash]
//   - [IVZMacBatteryPowerSourceDevice.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice
type IVZMacBatteryPowerSourceDevice interface {
	IVZPowerSourceDevice

	// Topic: Methods

	BatterySourceDidUpdateCharge(source objectivec.IObject, charge float64)
	BatterySourceDidUpdateConnectivity(source objectivec.IObject, connectivity int64)
	Source() *VZMacBatterySource
	SetSource(value *VZMacBatterySource)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZMacBatteryPowerSourceDevice) Init() VZMacBatteryPowerSourceDevice {
	rv := objc.Send[VZMacBatteryPowerSourceDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacBatteryPowerSourceDevice) Autorelease() VZMacBatteryPowerSourceDevice {
	rv := objc.Send[VZMacBatteryPowerSourceDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacBatteryPowerSourceDevice creates a new VZMacBatteryPowerSourceDevice instance.
func NewVZMacBatteryPowerSourceDevice() VZMacBatteryPowerSourceDevice {
	class := getVZMacBatteryPowerSourceDeviceClass()
	rv := objc.Send[VZMacBatteryPowerSourceDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/batterySource:didUpdateCharge:
func (v VZMacBatteryPowerSourceDevice) BatterySourceDidUpdateCharge(source objectivec.IObject, charge float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("batterySource:didUpdateCharge:"), source, charge)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/batterySource:didUpdateConnectivity:
func (v VZMacBatteryPowerSourceDevice) BatterySourceDidUpdateConnectivity(source objectivec.IObject, connectivity int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("batterySource:didUpdateConnectivity:"), source, connectivity)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/debugDescription
func (v VZMacBatteryPowerSourceDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/description
func (v VZMacBatteryPowerSourceDevice) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/hash
func (v VZMacBatteryPowerSourceDevice) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/source
func (v VZMacBatteryPowerSourceDevice) Source() *VZMacBatterySource {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("source"))
	if rv == 0 {
		return nil
	}
	val := VZMacBatterySourceFromID(objc.ID(rv))
	return &val
}
func (v VZMacBatteryPowerSourceDevice) SetSource(value *VZMacBatterySource) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setSource:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setSource:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDevice/superclass
func (v VZMacBatteryPowerSourceDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
