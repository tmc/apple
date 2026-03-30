// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacSyntheticBatterySource] class.
var (
	_VZMacSyntheticBatterySourceClass     VZMacSyntheticBatterySourceClass
	_VZMacSyntheticBatterySourceClassOnce sync.Once
)

func getVZMacSyntheticBatterySourceClass() VZMacSyntheticBatterySourceClass {
	_VZMacSyntheticBatterySourceClassOnce.Do(func() {
		_VZMacSyntheticBatterySourceClass = VZMacSyntheticBatterySourceClass{class: objc.GetClass("_VZMacSyntheticBatterySource")}
	})
	return _VZMacSyntheticBatterySourceClass
}

// GetVZMacSyntheticBatterySourceClass returns the class object for _VZMacSyntheticBatterySource.
func GetVZMacSyntheticBatterySourceClass() VZMacSyntheticBatterySourceClass {
	return getVZMacSyntheticBatterySourceClass()
}

type VZMacSyntheticBatterySourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacSyntheticBatterySourceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacSyntheticBatterySourceClass) Alloc() VZMacSyntheticBatterySource {
	rv := objc.Send[VZMacSyntheticBatterySource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacSyntheticBatterySource.Charge]
//   - [VZMacSyntheticBatterySource.SetCharge]
//   - [VZMacSyntheticBatterySource.Connectivity]
//   - [VZMacSyntheticBatterySource.SetConnectivity]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSyntheticBatterySource
type VZMacSyntheticBatterySource struct {
	VZMacBatterySource
}

// VZMacSyntheticBatterySourceFromID constructs a [VZMacSyntheticBatterySource] from an objc.ID.
func VZMacSyntheticBatterySourceFromID(id objc.ID) VZMacSyntheticBatterySource {
	return VZMacSyntheticBatterySource{VZMacBatterySource: VZMacBatterySourceFromID(id)}
}

// Ensure VZMacSyntheticBatterySource implements IVZMacSyntheticBatterySource.
var _ IVZMacSyntheticBatterySource = VZMacSyntheticBatterySource{}

// An interface definition for the [VZMacSyntheticBatterySource] class.
//
// # Methods
//
//   - [IVZMacSyntheticBatterySource.Charge]
//   - [IVZMacSyntheticBatterySource.SetCharge]
//   - [IVZMacSyntheticBatterySource.Connectivity]
//   - [IVZMacSyntheticBatterySource.SetConnectivity]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSyntheticBatterySource
type IVZMacSyntheticBatterySource interface {
	IVZMacBatterySource

	// Topic: Methods

	Charge() float64
	SetCharge(value float64)
	Connectivity() int64
	SetConnectivity(value int64)
}

// Init initializes the instance.
func (v VZMacSyntheticBatterySource) Init() VZMacSyntheticBatterySource {
	rv := objc.Send[VZMacSyntheticBatterySource](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacSyntheticBatterySource) Autorelease() VZMacSyntheticBatterySource {
	rv := objc.Send[VZMacSyntheticBatterySource](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacSyntheticBatterySource creates a new VZMacSyntheticBatterySource instance.
func NewVZMacSyntheticBatterySource() VZMacSyntheticBatterySource {
	class := getVZMacSyntheticBatterySourceClass()
	rv := objc.Send[VZMacSyntheticBatterySource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacSyntheticBatterySource/charge
func (v VZMacSyntheticBatterySource) Charge() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("charge"))
	return rv
}
func (v VZMacSyntheticBatterySource) SetCharge(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setCharge:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacSyntheticBatterySource/connectivity
func (v VZMacSyntheticBatterySource) Connectivity() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("connectivity"))
	return rv
}
func (v VZMacSyntheticBatterySource) SetConnectivity(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setConnectivity:"), value)
}
