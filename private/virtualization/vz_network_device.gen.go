// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDevice] class.
var (
	_VZNetworkDeviceClass     VZNetworkDeviceClass
	_VZNetworkDeviceClassOnce sync.Once
)

func getVZNetworkDeviceClass() VZNetworkDeviceClass {
	_VZNetworkDeviceClassOnce.Do(func() {
		_VZNetworkDeviceClass = VZNetworkDeviceClass{class: objc.GetClass("VZNetworkDevice")}
	})
	return _VZNetworkDeviceClass
}

// GetVZNetworkDeviceClass returns the class object for VZNetworkDevice.
func GetVZNetworkDeviceClass() VZNetworkDeviceClass {
	return getVZNetworkDeviceClass()
}

type VZNetworkDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNetworkDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceClass) Alloc() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNetworkDevice._type]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type VZNetworkDevice struct {
	objectivec.Object
}

// VZNetworkDeviceFromID constructs a [VZNetworkDevice] from an objc.ID.
func VZNetworkDeviceFromID(id objc.ID) VZNetworkDevice {
	return VZNetworkDevice{objectivec.Object{ID: id}}
}

// Ensure VZNetworkDevice implements IVZNetworkDevice.
var _ IVZNetworkDevice = VZNetworkDevice{}

// An interface definition for the [VZNetworkDevice] class.
//
// # Methods
//
//   - [IVZNetworkDevice._type]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type IVZNetworkDevice interface {
	objectivec.IObject

	// Topic: Methods

	_type() int64
}

// Init initializes the instance.
func (v VZNetworkDevice) Init() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNetworkDevice) Autorelease() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDevice creates a new VZNetworkDevice instance.
func NewVZNetworkDevice() VZNetworkDevice {
	class := getVZNetworkDeviceClass()
	rv := objc.Send[VZNetworkDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice/_type
func (v VZNetworkDevice) _type() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_type"))
	return rv
}

// CanType reports whether the receiver responds to the private selector _type.
func (v VZNetworkDevice) CanType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_type"))
}

// Type is an exported wrapper for the private property _type.
func (v VZNetworkDevice) Type() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_type")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_type"}
	}
	return v._type(), nil
}
