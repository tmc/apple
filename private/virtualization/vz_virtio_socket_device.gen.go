// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSocketDevice] class.
var (
	_VZVirtioSocketDeviceClass     VZVirtioSocketDeviceClass
	_VZVirtioSocketDeviceClassOnce sync.Once
)

func getVZVirtioSocketDeviceClass() VZVirtioSocketDeviceClass {
	_VZVirtioSocketDeviceClassOnce.Do(func() {
		_VZVirtioSocketDeviceClass = VZVirtioSocketDeviceClass{class: objc.GetClass("VZVirtioSocketDevice")}
	})
	return _VZVirtioSocketDeviceClass
}

// GetVZVirtioSocketDeviceClass returns the class object for VZVirtioSocketDevice.
func GetVZVirtioSocketDeviceClass() VZVirtioSocketDeviceClass {
	return getVZVirtioSocketDeviceClass()
}

type VZVirtioSocketDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSocketDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketDeviceClass) Alloc() VZVirtioSocketDevice {
	rv := objc.SendIfResponds[VZVirtioSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioSocketDevice._configurationOptions]
//   - [VZVirtioSocketDevice._setDelegate]
//   - [VZVirtioSocketDevice._virtualMachineClientAuditToken]
type VZVirtioSocketDevice struct {
	VZSocketDevice
}

// VZVirtioSocketDeviceFromID constructs a [VZVirtioSocketDevice] from an objc.ID.
func VZVirtioSocketDeviceFromID(id objc.ID) VZVirtioSocketDevice {
	return VZVirtioSocketDevice{VZSocketDevice: VZSocketDeviceFromID(id)}
}

// Ensure VZVirtioSocketDevice implements IVZVirtioSocketDevice.
var _ IVZVirtioSocketDevice = VZVirtioSocketDevice{}

// An interface definition for the [VZVirtioSocketDevice] class.
//
// # Methods
//
//   - [IVZVirtioSocketDevice._configurationOptions]
//   - [IVZVirtioSocketDevice._setDelegate]
//   - [IVZVirtioSocketDevice._virtualMachineClientAuditToken]
type IVZVirtioSocketDevice interface {
	IVZSocketDevice

	// Topic: Methods

	_configurationOptions() objectivec.IObject
	_setDelegate(delegate objectivec.IObject)
	_virtualMachineClientAuditToken() unsafe.Pointer
}

// Init initializes the instance.
func (v VZVirtioSocketDevice) Init() VZVirtioSocketDevice {
	rv := objc.SendIfResponds[VZVirtioSocketDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketDevice) Autorelease() VZVirtioSocketDevice {
	rv := objc.SendIfResponds[VZVirtioSocketDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDevice creates a new VZVirtioSocketDevice instance.
func NewVZVirtioSocketDevice() VZVirtioSocketDevice {
	class := getVZVirtioSocketDeviceClass()
	rv := objc.SendIfResponds[VZVirtioSocketDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioSocketDevice) _configurationOptions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_configurationOptions"))
	return objectivec.Object{ID: rv}
}

// ConfigurationOptions is an exported wrapper for the private method _configurationOptions.
func (v VZVirtioSocketDevice) ConfigurationOptions() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_configurationOptions")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_configurationOptions"}
		return nil, err
	}
	return v._configurationOptions(), nil
}

// CanConfigurationOptions reports whether the receiver responds to the private selector _configurationOptions.
func (v VZVirtioSocketDevice) CanConfigurationOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_configurationOptions"))
}
func (v VZVirtioSocketDevice) _setDelegate(delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setDelegate:"), delegate)
}

// SetDelegate is an exported wrapper for the private method _setDelegate.
func (v VZVirtioSocketDevice) SetDelegate(delegate objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDelegate:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDelegate:"}
		return err
	}
	v._setDelegate(delegate)
	return nil
}

// CanSetDelegate reports whether the receiver responds to the private selector _setDelegate:.
func (v VZVirtioSocketDevice) CanSetDelegate() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDelegate:"))
}
func (v VZVirtioSocketDevice) _virtualMachineClientAuditToken() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_virtualMachineClientAuditToken"))
	return rv
}

// VirtualMachineClientAuditToken is an exported wrapper for the private method _virtualMachineClientAuditToken.
func (v VZVirtioSocketDevice) VirtualMachineClientAuditToken() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_virtualMachineClientAuditToken")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_virtualMachineClientAuditToken"}
		return nil, err
	}
	return v._virtualMachineClientAuditToken(), nil
}

// CanVirtualMachineClientAuditToken reports whether the receiver responds to the private selector _virtualMachineClientAuditToken.
func (v VZVirtioSocketDevice) CanVirtualMachineClientAuditToken() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_virtualMachineClientAuditToken"))
}
