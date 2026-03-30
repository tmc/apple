// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBOpticalDriveDeviceConfiguration] class.
var (
	_VZUSBOpticalDriveDeviceConfigurationClass     VZUSBOpticalDriveDeviceConfigurationClass
	_VZUSBOpticalDriveDeviceConfigurationClassOnce sync.Once
)

func getVZUSBOpticalDriveDeviceConfigurationClass() VZUSBOpticalDriveDeviceConfigurationClass {
	_VZUSBOpticalDriveDeviceConfigurationClassOnce.Do(func() {
		_VZUSBOpticalDriveDeviceConfigurationClass = VZUSBOpticalDriveDeviceConfigurationClass{class: objc.GetClass("_VZUSBOpticalDriveDeviceConfiguration")}
	})
	return _VZUSBOpticalDriveDeviceConfigurationClass
}

// GetVZUSBOpticalDriveDeviceConfigurationClass returns the class object for _VZUSBOpticalDriveDeviceConfiguration.
func GetVZUSBOpticalDriveDeviceConfigurationClass() VZUSBOpticalDriveDeviceConfigurationClass {
	return getVZUSBOpticalDriveDeviceConfigurationClass()
}

type VZUSBOpticalDriveDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBOpticalDriveDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBOpticalDriveDeviceConfigurationClass) Alloc() VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.Send[VZUSBOpticalDriveDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler]
//   - [VZUSBOpticalDriveDeviceConfiguration.EncodeWithEncoder]
//   - [VZUSBOpticalDriveDeviceConfiguration.InitWithAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration
type VZUSBOpticalDriveDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZUSBOpticalDriveDeviceConfigurationFromID constructs a [VZUSBOpticalDriveDeviceConfiguration] from an objc.ID.
func VZUSBOpticalDriveDeviceConfigurationFromID(id objc.ID) VZUSBOpticalDriveDeviceConfiguration {
	return VZUSBOpticalDriveDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}

// Ensure VZUSBOpticalDriveDeviceConfiguration implements IVZUSBOpticalDriveDeviceConfiguration.
var _ IVZUSBOpticalDriveDeviceConfiguration = VZUSBOpticalDriveDeviceConfiguration{}

// An interface definition for the [VZUSBOpticalDriveDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler]
//   - [IVZUSBOpticalDriveDeviceConfiguration.EncodeWithEncoder]
//   - [IVZUSBOpticalDriveDeviceConfiguration.InitWithAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration
type IVZUSBOpticalDriveDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	// Topic: Methods

	_getStorageDeviceWithQueueSessionCompletionHandler(queue objectivec.IObject, session unsafe.Pointer, handler ErrorHandler)
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	InitWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration
}

// Init initializes the instance.
func (v VZUSBOpticalDriveDeviceConfiguration) Init() VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.Send[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBOpticalDriveDeviceConfiguration) Autorelease() VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.Send[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBOpticalDriveDeviceConfiguration creates a new VZUSBOpticalDriveDeviceConfiguration instance.
func NewVZUSBOpticalDriveDeviceConfiguration() VZUSBOpticalDriveDeviceConfiguration {
	class := getVZUSBOpticalDriveDeviceConfigurationClass()
	rv := objc.Send[VZUSBOpticalDriveDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration/initWithAttachment:
func NewVZUSBOpticalDriveDeviceConfigurationWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration {
	instance := getVZUSBOpticalDriveDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	return VZUSBOpticalDriveDeviceConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration/_getStorageDeviceWithQueue:session:completionHandler:
func (v VZUSBOpticalDriveDeviceConfiguration) _getStorageDeviceWithQueueSessionCompletionHandler(queue objectivec.IObject, session unsafe.Pointer, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_getStorageDeviceWithQueue:session:completionHandler:"), queue, session, _block2)
}

// GetStorageDeviceWithQueueSessionCompletionHandler is an exported wrapper for the private method _getStorageDeviceWithQueueSessionCompletionHandler.
func (v VZUSBOpticalDriveDeviceConfiguration) GetStorageDeviceWithQueueSessionCompletionHandler(queue objectivec.IObject, session unsafe.Pointer, handler ErrorHandler) {
	v._getStorageDeviceWithQueueSessionCompletionHandler(queue, session, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration/encodeWithEncoder:
func (v VZUSBOpticalDriveDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBOpticalDriveDeviceConfiguration/initWithAttachment:
func (v VZUSBOpticalDriveDeviceConfiguration) InitWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.Send[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("initWithAttachment:"), attachment)
	return rv
}

// _getStorageDeviceWithQueueSession is a synchronous wrapper around [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZUSBOpticalDriveDeviceConfiguration) _getStorageDeviceWithQueueSession(ctx context.Context, queue objectivec.IObject, session unsafe.Pointer) error {
	done := make(chan error, 1)
	v._getStorageDeviceWithQueueSessionCompletionHandler(queue, session, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
