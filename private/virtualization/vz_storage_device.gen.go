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
func (v VZStorageDevice) Init() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZStorageDevice) Autorelease() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDevice creates a new VZStorageDevice instance.
func NewVZStorageDevice() VZStorageDevice {
	class := getVZStorageDeviceClass()
	rv := objc.Send[VZStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_attachment
func (v VZStorageDevice) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}

// Attachment is an exported wrapper for the private method _attachment.
func (v VZStorageDevice) Attachment() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_attachment")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_attachment"}
		return nil, err
	}
	return v._attachment(), nil
}

// CanAttachment reports whether the receiver responds to the private selector _attachment.
func (v VZStorageDevice) CanAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_attachment"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithAttachment:
func (v VZStorageDevice) _initWithAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

// InitWithAttachment is an exported wrapper for the private method _initWithAttachment.
func (v VZStorageDevice) InitWithAttachment(attachment objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithAttachment:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithAttachment:"}
		return nil, err
	}
	return v._initWithAttachment(attachment), nil
}

// CanInitWithAttachment reports whether the receiver responds to the private selector _initWithAttachment:.
func (v VZStorageDevice) CanInitWithAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithAttachment:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithVirtualMachine:attachment:
func (v VZStorageDevice) _initWithVirtualMachineAttachment(machine objectivec.IObject, attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithVirtualMachine:attachment:"), machine, attachment)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineAttachment is an exported wrapper for the private method _initWithVirtualMachineAttachment.
func (v VZStorageDevice) InitWithVirtualMachineAttachment(machine objectivec.IObject, attachment objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:attachment:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithVirtualMachine:attachment:"}
		return nil, err
	}
	return v._initWithVirtualMachineAttachment(machine, attachment), nil
}

// CanInitWithVirtualMachineAttachment reports whether the receiver responds to the private selector _initWithVirtualMachine:attachment:.
func (v VZStorageDevice) CanInitWithVirtualMachineAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:attachment:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_initWithVirtualMachine:storageDeviceIndex:attachment:
func (v VZStorageDevice) _initWithVirtualMachineStorageDeviceIndexAttachment(machine objectivec.IObject, index uint64, attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithVirtualMachine:storageDeviceIndex:attachment:"), machine, index, attachment)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineStorageDeviceIndexAttachment is an exported wrapper for the private method _initWithVirtualMachineStorageDeviceIndexAttachment.
func (v VZStorageDevice) InitWithVirtualMachineStorageDeviceIndexAttachment(machine objectivec.IObject, index uint64, attachment objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:storageDeviceIndex:attachment:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithVirtualMachine:storageDeviceIndex:attachment:"}
		return nil, err
	}
	return v._initWithVirtualMachineStorageDeviceIndexAttachment(machine, index, attachment), nil
}

// CanInitWithVirtualMachineStorageDeviceIndexAttachment reports whether the receiver responds to the private selector _initWithVirtualMachine:storageDeviceIndex:attachment:.
func (v VZStorageDevice) CanInitWithVirtualMachineStorageDeviceIndexAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:storageDeviceIndex:attachment:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_setAttachment:completionHandler:
func (v VZStorageDevice) _setAttachmentCompletionHandler(attachment objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_setAttachment:completionHandler:"), attachment, _block1)
}

// SetAttachmentCompletionHandler is an exported wrapper for the private method _setAttachmentCompletionHandler.
func (v VZStorageDevice) SetAttachmentCompletionHandler(attachment objectivec.IObject, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setAttachment:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setAttachment:completionHandler:"}
		return err
	}
	v._setAttachmentCompletionHandler(attachment, handler)
	return nil
}

// CanSetAttachmentCompletionHandler reports whether the receiver responds to the private selector _setAttachment:completionHandler:.
func (v VZStorageDevice) CanSetAttachmentCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setAttachment:completionHandler:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice/_setVirtualMachine:
func (v VZStorageDevice) _setVirtualMachine(machine objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setVirtualMachine:"), machine)
}

// SetVirtualMachine is an exported wrapper for the private method _setVirtualMachine.
func (v VZStorageDevice) SetVirtualMachine(machine objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setVirtualMachine:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setVirtualMachine:"}
		return err
	}
	v._setVirtualMachine(machine)
	return nil
}

// CanSetVirtualMachine reports whether the receiver responds to the private selector _setVirtualMachine:.
func (v VZStorageDevice) CanSetVirtualMachine() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setVirtualMachine:"))
}

// _setAttachment is a synchronous wrapper around [VZStorageDevice._setAttachmentCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZStorageDevice) _setAttachment(ctx context.Context, attachment objectivec.IObject) error {
	done := make(chan error, 1)
	v._setAttachmentCompletionHandler(attachment, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
