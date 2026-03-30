// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDevice] class.
var (
	_VZStorageDeviceClass     VZStorageDeviceClass
	_VZStorageDeviceClassOnce sync.Once
)

func getVZStorageDeviceClass() VZStorageDeviceClass {
	_VZStorageDeviceClassOnce.Do(func() {
		_VZStorageDeviceClass = VZStorageDeviceClass{class: objc.GetClass("VZStorageDevice")}
	})
	return _VZStorageDeviceClass
}

// GetVZStorageDeviceClass returns the class object for VZStorageDevice.
func GetVZStorageDeviceClass() VZStorageDeviceClass {
	return getVZStorageDeviceClass()
}

type VZStorageDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZStorageDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceClass) Alloc() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZStorageDevice._attachment]
//   - [VZStorageDevice._initWithAttachment]
//   - [VZStorageDevice._initWithVirtualMachineAttachment]
//   - [VZStorageDevice._initWithVirtualMachineStorageDeviceIndexAttachment]
//   - [VZStorageDevice._setAttachmentCompletionHandler]
//   - [VZStorageDevice._setVirtualMachine]
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type VZStorageDevice struct {
	objectivec.Object
}

// VZStorageDeviceFromID constructs a [VZStorageDevice] from an objc.ID.
func VZStorageDeviceFromID(id objc.ID) VZStorageDevice {
	return VZStorageDevice{objectivec.Object{ID: id}}
}

// Ensure VZStorageDevice implements IVZStorageDevice.
var _ IVZStorageDevice = VZStorageDevice{}

// An interface definition for the [VZStorageDevice] class.
//
// # Methods
//
//   - [IVZStorageDevice._attachment]
//   - [IVZStorageDevice._initWithAttachment]
//   - [IVZStorageDevice._initWithVirtualMachineAttachment]
//   - [IVZStorageDevice._initWithVirtualMachineStorageDeviceIndexAttachment]
//   - [IVZStorageDevice._setAttachmentCompletionHandler]
//   - [IVZStorageDevice._setVirtualMachine]
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type IVZStorageDevice interface {
	objectivec.IObject

	// Topic: Methods

	_attachment() objectivec.IObject
	_initWithAttachment(attachment objectivec.IObject) objectivec.IObject
	_initWithVirtualMachineAttachment(machine objectivec.IObject, attachment objectivec.IObject) objectivec.IObject
	_initWithVirtualMachineStorageDeviceIndexAttachment(machine objectivec.IObject, index uint64, attachment objectivec.IObject) objectivec.IObject
	_setAttachmentCompletionHandler(attachment objectivec.IObject, handler ErrorHandler)
	_setVirtualMachine(machine objectivec.IObject)
}

// Init initializes the instance.
func (s VZStorageDevice) Init() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDevice) Autorelease() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDevice creates a new VZStorageDevice instance.
func NewVZStorageDevice() VZStorageDevice {
	class := getVZStorageDeviceClass()
	rv := objc.Send[VZStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_attachment
func (s VZStorageDevice) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}

// Attachment is an exported wrapper for the private method _attachment.
func (s VZStorageDevice) Attachment() objectivec.IObject {
	return s._attachment()
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithAttachment:
func (s VZStorageDevice) _initWithAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

// InitWithAttachment is an exported wrapper for the private method _initWithAttachment.
func (s VZStorageDevice) InitWithAttachment(attachment objectivec.IObject) objectivec.IObject {
	return s._initWithAttachment(attachment)
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithVirtualMachine:attachment:
func (s VZStorageDevice) _initWithVirtualMachineAttachment(machine objectivec.IObject, attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithVirtualMachine:attachment:"), machine, attachment)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineAttachment is an exported wrapper for the private method _initWithVirtualMachineAttachment.
func (s VZStorageDevice) InitWithVirtualMachineAttachment(machine objectivec.IObject, attachment objectivec.IObject) objectivec.IObject {
	return s._initWithVirtualMachineAttachment(machine, attachment)
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithVirtualMachine:storageDeviceIndex:attachment:
func (s VZStorageDevice) _initWithVirtualMachineStorageDeviceIndexAttachment(machine objectivec.IObject, index uint64, attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithVirtualMachine:storageDeviceIndex:attachment:"), machine, index, attachment)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineStorageDeviceIndexAttachment is an exported wrapper for the private method _initWithVirtualMachineStorageDeviceIndexAttachment.
func (s VZStorageDevice) InitWithVirtualMachineStorageDeviceIndexAttachment(machine objectivec.IObject, index uint64, attachment objectivec.IObject) objectivec.IObject {
	return s._initWithVirtualMachineStorageDeviceIndexAttachment(machine, index, attachment)
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_setAttachment:completionHandler:
func (s VZStorageDevice) _setAttachmentCompletionHandler(attachment objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("_setAttachment:completionHandler:"), attachment, _block1)
}

// SetAttachmentCompletionHandler is an exported wrapper for the private method _setAttachmentCompletionHandler.
func (s VZStorageDevice) SetAttachmentCompletionHandler(attachment objectivec.IObject, handler ErrorHandler) {
	s._setAttachmentCompletionHandler(attachment, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_setVirtualMachine:
func (s VZStorageDevice) _setVirtualMachine(machine objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_setVirtualMachine:"), machine)
}

// SetVirtualMachine is an exported wrapper for the private method _setVirtualMachine.
func (s VZStorageDevice) SetVirtualMachine(machine objectivec.IObject) {
	s._setVirtualMachine(machine)
}

// _setAttachment is a synchronous wrapper around [VZStorageDevice._setAttachmentCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s VZStorageDevice) _setAttachment(ctx context.Context, attachment objectivec.IObject) error {
	done := make(chan error, 1)
	s._setAttachmentCompletionHandler(attachment, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
