// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZBridgedNetworkDeviceAttachment] class.
var (
	_VZBridgedNetworkDeviceAttachmentClass     VZBridgedNetworkDeviceAttachmentClass
	_VZBridgedNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	_VZBridgedNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZBridgedNetworkDeviceAttachmentClass = VZBridgedNetworkDeviceAttachmentClass{class: objc.GetClass("VZBridgedNetworkDeviceAttachment")}
	})
	return _VZBridgedNetworkDeviceAttachmentClass
}

// GetVZBridgedNetworkDeviceAttachmentClass returns the class object for VZBridgedNetworkDeviceAttachment.
func GetVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	return getVZBridgedNetworkDeviceAttachmentClass()
}

type VZBridgedNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBridgedNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBridgedNetworkDeviceAttachmentClass) Alloc() VZBridgedNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZBridgedNetworkDeviceAttachment._macNatEnabled]
//   - [VZBridgedNetworkDeviceAttachment.Set_macNatEnabled]
//   - [VZBridgedNetworkDeviceAttachment._setMacNatEnabled]
type VZBridgedNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZBridgedNetworkDeviceAttachmentFromID constructs a [VZBridgedNetworkDeviceAttachment] from an objc.ID.
func VZBridgedNetworkDeviceAttachmentFromID(id objc.ID) VZBridgedNetworkDeviceAttachment {
	return VZBridgedNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZBridgedNetworkDeviceAttachment implements IVZBridgedNetworkDeviceAttachment.
var _ IVZBridgedNetworkDeviceAttachment = VZBridgedNetworkDeviceAttachment{}

// An interface definition for the [VZBridgedNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZBridgedNetworkDeviceAttachment._macNatEnabled]
//   - [IVZBridgedNetworkDeviceAttachment.Set_macNatEnabled]
//   - [IVZBridgedNetworkDeviceAttachment._setMacNatEnabled]
type IVZBridgedNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Methods

	_macNatEnabled() bool
	Set_macNatEnabled(value bool)
	_setMacNatEnabled(enabled bool)
}

// Init initializes the instance.
func (v VZBridgedNetworkDeviceAttachment) Init() VZBridgedNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZBridgedNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBridgedNetworkDeviceAttachment) Autorelease() VZBridgedNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZBridgedNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkDeviceAttachment creates a new VZBridgedNetworkDeviceAttachment instance.
func NewVZBridgedNetworkDeviceAttachment() VZBridgedNetworkDeviceAttachment {
	class := getVZBridgedNetworkDeviceAttachmentClass()
	rv := objc.SendIfResponds[VZBridgedNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZBridgedNetworkDeviceAttachment) _setMacNatEnabled(enabled bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setMacNatEnabled:"), enabled)
}

// SetMacNatEnabled is an exported wrapper for the private method _setMacNatEnabled.
func (v VZBridgedNetworkDeviceAttachment) SetMacNatEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setMacNatEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMacNatEnabled:"}
		return err
	}
	v._setMacNatEnabled(enabled)
	return nil
}

// CanSetMacNatEnabled reports whether the receiver responds to the private selector _setMacNatEnabled:.
func (v VZBridgedNetworkDeviceAttachment) CanSetMacNatEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setMacNatEnabled:"))
}

func (v VZBridgedNetworkDeviceAttachment) _macNatEnabled() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_macNatEnabled"))
	return rv
}

// CanMacNatEnabled reports whether the receiver responds to the private selector _macNatEnabled.
func (v VZBridgedNetworkDeviceAttachment) CanMacNatEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_macNatEnabled"))
}

// MacNatEnabled is an exported wrapper for the private property _macNatEnabled.
func (v VZBridgedNetworkDeviceAttachment) MacNatEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_macNatEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_macNatEnabled"}
	}
	return v._macNatEnabled(), nil
}
func (v VZBridgedNetworkDeviceAttachment) Set_macNatEnabled(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_macNatEnabled:"), value)
}
