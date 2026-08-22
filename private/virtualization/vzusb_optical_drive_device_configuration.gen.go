// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"

	"github.com/tmc/apple/dispatch"
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
	rv := objc.SendIfResponds[VZUSBOpticalDriveDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler]
//   - [VZUSBOpticalDriveDeviceConfiguration.InitWithAttachment]
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
//   - [IVZUSBOpticalDriveDeviceConfiguration.InitWithAttachment]
type IVZUSBOpticalDriveDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	// Topic: Methods

	_getStorageDeviceWithQueueSessionCompletionHandler(queue dispatch.Queue, session *DispatchGroupSession, handler ErrorHandler)
	InitWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration
}

// Init initializes the instance.
func (v VZUSBOpticalDriveDeviceConfiguration) Init() VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.SendIfResponds[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBOpticalDriveDeviceConfiguration) Autorelease() VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.SendIfResponds[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBOpticalDriveDeviceConfiguration creates a new VZUSBOpticalDriveDeviceConfiguration instance.
func NewVZUSBOpticalDriveDeviceConfiguration() VZUSBOpticalDriveDeviceConfiguration {
	class := getVZUSBOpticalDriveDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZUSBOpticalDriveDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZUSBOpticalDriveDeviceConfigurationWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration {
	instance := getVZUSBOpticalDriveDeviceConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	return VZUSBOpticalDriveDeviceConfigurationFromID(rv)
}

func (v VZUSBOpticalDriveDeviceConfiguration) _getStorageDeviceWithQueueSessionCompletionHandler(queue dispatch.Queue, session *DispatchGroupSession, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_getStorageDeviceWithQueue:session:completionHandler:"), uintptr(queue.Handle()), session, _block2)
}

// GetStorageDeviceWithQueueSessionCompletionHandler is an exported wrapper for the private method _getStorageDeviceWithQueueSessionCompletionHandler.
func (v VZUSBOpticalDriveDeviceConfiguration) GetStorageDeviceWithQueueSessionCompletionHandler(queue dispatch.Queue, session *DispatchGroupSession, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_getStorageDeviceWithQueue:session:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_getStorageDeviceWithQueue:session:completionHandler:"}
		return err
	}
	v._getStorageDeviceWithQueueSessionCompletionHandler(queue, session, handler)
	return nil
}

// CanGetStorageDeviceWithQueueSessionCompletionHandler reports whether the receiver responds to the private selector _getStorageDeviceWithQueue:session:completionHandler:.
func (v VZUSBOpticalDriveDeviceConfiguration) CanGetStorageDeviceWithQueueSessionCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_getStorageDeviceWithQueue:session:completionHandler:"))
}
func (v VZUSBOpticalDriveDeviceConfiguration) InitWithAttachment(attachment objectivec.IObject) VZUSBOpticalDriveDeviceConfiguration {
	rv := objc.SendIfResponds[VZUSBOpticalDriveDeviceConfiguration](v.ID, objc.Sel("initWithAttachment:"), attachment)
	return rv
}

// _getStorageDeviceWithQueueSession is a synchronous wrapper around [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZUSBOpticalDriveDeviceConfiguration) _getStorageDeviceWithQueueSession(ctx context.Context, queue dispatch.Queue, session *DispatchGroupSession) error {
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
