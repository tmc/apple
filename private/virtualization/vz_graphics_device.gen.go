// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDevice] class.
var (
	_VZGraphicsDeviceClass     VZGraphicsDeviceClass
	_VZGraphicsDeviceClassOnce sync.Once
)

func getVZGraphicsDeviceClass() VZGraphicsDeviceClass {
	_VZGraphicsDeviceClassOnce.Do(func() {
		_VZGraphicsDeviceClass = VZGraphicsDeviceClass{class: objc.GetClass("VZGraphicsDevice")}
	})
	return _VZGraphicsDeviceClass
}

// GetVZGraphicsDeviceClass returns the class object for VZGraphicsDevice.
func GetVZGraphicsDeviceClass() VZGraphicsDeviceClass {
	return getVZGraphicsDeviceClass()
}

type VZGraphicsDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGraphicsDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDeviceClass) Alloc() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGraphicsDevice._attachDisplayCompletionHandler]
//   - [VZGraphicsDevice._detachDisplayCompletionHandler]
//   - [VZGraphicsDevice._displayPortCount]
//   - [VZGraphicsDevice._initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays]
//   - [VZGraphicsDevice._validateDisplayForHotPlugError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice
type VZGraphicsDevice struct {
	objectivec.Object
}

// VZGraphicsDeviceFromID constructs a [VZGraphicsDevice] from an objc.ID.
func VZGraphicsDeviceFromID(id objc.ID) VZGraphicsDevice {
	return VZGraphicsDevice{objectivec.Object{ID: id}}
}

// Ensure VZGraphicsDevice implements IVZGraphicsDevice.
var _ IVZGraphicsDevice = VZGraphicsDevice{}

// An interface definition for the [VZGraphicsDevice] class.
//
// # Methods
//
//   - [IVZGraphicsDevice._attachDisplayCompletionHandler]
//   - [IVZGraphicsDevice._detachDisplayCompletionHandler]
//   - [IVZGraphicsDevice._displayPortCount]
//   - [IVZGraphicsDevice._initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays]
//   - [IVZGraphicsDevice._validateDisplayForHotPlugError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice
type IVZGraphicsDevice interface {
	objectivec.IObject

	// Topic: Methods

	_attachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler)
	_detachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler)
	_displayPortCount() uint64
	_initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays(machine objectivec.IObject, index uint64, count uint64, displays objectivec.IObject) objectivec.IObject
	_validateDisplayForHotPlugError(plug objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (g VZGraphicsDevice) Init() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDevice) Autorelease() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDevice creates a new VZGraphicsDevice instance.
func NewVZGraphicsDevice() VZGraphicsDevice {
	class := getVZGraphicsDeviceClass()
	rv := objc.Send[VZGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/_attachDisplay:completionHandler:
func (g VZGraphicsDevice) _attachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](g.ID, objc.Sel("_attachDisplay:completionHandler:"), display, _block1)
}

// AttachDisplayCompletionHandler is an exported wrapper for the private method _attachDisplayCompletionHandler.
func (g VZGraphicsDevice) AttachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler) {
	g._attachDisplayCompletionHandler(display, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/_detachDisplay:completionHandler:
func (g VZGraphicsDevice) _detachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](g.ID, objc.Sel("_detachDisplay:completionHandler:"), display, _block1)
}

// DetachDisplayCompletionHandler is an exported wrapper for the private method _detachDisplayCompletionHandler.
func (g VZGraphicsDevice) DetachDisplayCompletionHandler(display objectivec.IObject, handler ErrorHandler) {
	g._detachDisplayCompletionHandler(display, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/_initWithVirtualMachine:graphicsDeviceIndex:displayPortCount:displays:
func (g VZGraphicsDevice) _initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays(machine objectivec.IObject, index uint64, count uint64, displays objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_initWithVirtualMachine:graphicsDeviceIndex:displayPortCount:displays:"), machine, index, count, displays)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays is an exported wrapper for the private method _initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays.
func (g VZGraphicsDevice) InitWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays(machine objectivec.IObject, index uint64, count uint64, displays objectivec.IObject) objectivec.IObject {
	return g._initWithVirtualMachineGraphicsDeviceIndexDisplayPortCountDisplays(machine, index, count, displays)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/_validateDisplayForHotPlug:error:
func (g VZGraphicsDevice) _validateDisplayForHotPlugError(plug objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("_validateDisplayForHotPlug:error:"), plug, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_validateDisplayForHotPlug:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ValidateDisplayForHotPlugError is an exported wrapper for the private method _validateDisplayForHotPlugError.
func (g VZGraphicsDevice) ValidateDisplayForHotPlugError(plug objectivec.IObject) (bool, error) {
	return g._validateDisplayForHotPlugError(plug)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/_displayPortCount
func (g VZGraphicsDevice) _displayPortCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("_displayPortCount"))
	return rv
}

// _attachDisplay is a synchronous wrapper around [VZGraphicsDevice._attachDisplayCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (g VZGraphicsDevice) _attachDisplay(ctx context.Context, display objectivec.IObject) error {
	done := make(chan error, 1)
	g._attachDisplayCompletionHandler(display, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _detachDisplay is a synchronous wrapper around [VZGraphicsDevice._detachDisplayCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (g VZGraphicsDevice) _detachDisplay(ctx context.Context, display objectivec.IObject) error {
	done := make(chan error, 1)
	g._detachDisplayCompletionHandler(display, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
