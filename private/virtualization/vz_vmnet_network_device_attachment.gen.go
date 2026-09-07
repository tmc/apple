// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVmnetNetworkDeviceAttachment] class.
var (
	_VZVmnetNetworkDeviceAttachmentClass     VZVmnetNetworkDeviceAttachmentClass
	_VZVmnetNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	_VZVmnetNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZVmnetNetworkDeviceAttachmentClass = VZVmnetNetworkDeviceAttachmentClass{class: objc.GetClass("VZVmnetNetworkDeviceAttachment")}
	})
	return _VZVmnetNetworkDeviceAttachmentClass
}

// GetVZVmnetNetworkDeviceAttachmentClass returns the class object for VZVmnetNetworkDeviceAttachment.
func GetVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	return getVZVmnetNetworkDeviceAttachmentClass()
}

type VZVmnetNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVmnetNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVmnetNetworkDeviceAttachmentClass) Alloc() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVmnetNetworkDeviceAttachment._clientAuditToken]
//   - [VZVmnetNetworkDeviceAttachment._setClientAuditToken]
type VZVmnetNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZVmnetNetworkDeviceAttachmentFromID constructs a [VZVmnetNetworkDeviceAttachment] from an objc.ID.
func VZVmnetNetworkDeviceAttachmentFromID(id objc.ID) VZVmnetNetworkDeviceAttachment {
	return VZVmnetNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZVmnetNetworkDeviceAttachment implements IVZVmnetNetworkDeviceAttachment.
var _ IVZVmnetNetworkDeviceAttachment = VZVmnetNetworkDeviceAttachment{}

// An interface definition for the [VZVmnetNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZVmnetNetworkDeviceAttachment._clientAuditToken]
//   - [IVZVmnetNetworkDeviceAttachment._setClientAuditToken]
type IVZVmnetNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Methods

	_clientAuditToken() unsafe.Pointer
	_setClientAuditToken(token unsafe.Pointer)
}

// Init initializes the instance.
func (v VZVmnetNetworkDeviceAttachment) Init() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVmnetNetworkDeviceAttachment) Autorelease() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVmnetNetworkDeviceAttachment creates a new VZVmnetNetworkDeviceAttachment instance.
func NewVZVmnetNetworkDeviceAttachment() VZVmnetNetworkDeviceAttachment {
	class := getVZVmnetNetworkDeviceAttachmentClass()
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVmnetNetworkDeviceAttachment) _clientAuditToken() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_clientAuditToken"))
	return rv
}

// ClientAuditToken is an exported wrapper for the private method _clientAuditToken.
func (v VZVmnetNetworkDeviceAttachment) ClientAuditToken() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_clientAuditToken")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_clientAuditToken"}
		return nil, err
	}
	return v._clientAuditToken(), nil
}

// CanClientAuditToken reports whether the receiver responds to the private selector _clientAuditToken.
func (v VZVmnetNetworkDeviceAttachment) CanClientAuditToken() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_clientAuditToken"))
}
func (v VZVmnetNetworkDeviceAttachment) _setClientAuditToken(token unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setClientAuditToken:"), token)
}

// SetClientAuditToken is an exported wrapper for the private method _setClientAuditToken.
func (v VZVmnetNetworkDeviceAttachment) SetClientAuditToken(token unsafe.Pointer) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setClientAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setClientAuditToken:"}
		return err
	}
	v._setClientAuditToken(token)
	return nil
}

// CanSetClientAuditToken reports whether the receiver responds to the private selector _setClientAuditToken:.
func (v VZVmnetNetworkDeviceAttachment) CanSetClientAuditToken() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setClientAuditToken:"))
}
