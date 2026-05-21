// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZTemporaryRAMStorageDeviceAttachment] class.
var (
	_VZTemporaryRAMStorageDeviceAttachmentClass     VZTemporaryRAMStorageDeviceAttachmentClass
	_VZTemporaryRAMStorageDeviceAttachmentClassOnce sync.Once
)

func getVZTemporaryRAMStorageDeviceAttachmentClass() VZTemporaryRAMStorageDeviceAttachmentClass {
	_VZTemporaryRAMStorageDeviceAttachmentClassOnce.Do(func() {
		_VZTemporaryRAMStorageDeviceAttachmentClass = VZTemporaryRAMStorageDeviceAttachmentClass{class: objc.GetClass("_VZTemporaryRAMStorageDeviceAttachment")}
	})
	return _VZTemporaryRAMStorageDeviceAttachmentClass
}

// GetVZTemporaryRAMStorageDeviceAttachmentClass returns the class object for _VZTemporaryRAMStorageDeviceAttachment.
func GetVZTemporaryRAMStorageDeviceAttachmentClass() VZTemporaryRAMStorageDeviceAttachmentClass {
	return getVZTemporaryRAMStorageDeviceAttachmentClass()
}

type VZTemporaryRAMStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZTemporaryRAMStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZTemporaryRAMStorageDeviceAttachmentClass) Alloc() VZTemporaryRAMStorageDeviceAttachment {
	rv := objc.Send[VZTemporaryRAMStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZTemporaryRAMStorageDeviceAttachment.URL]
//   - [VZTemporaryRAMStorageDeviceAttachment._getAttachmentWithQueueCompletionHandler]
//   - [VZTemporaryRAMStorageDeviceAttachment.IsReadOnly]
//   - [VZTemporaryRAMStorageDeviceAttachment.InitWithURLReadOnlyError]
//   - [VZTemporaryRAMStorageDeviceAttachment.ReadOnly]
type VZTemporaryRAMStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZTemporaryRAMStorageDeviceAttachmentFromID constructs a [VZTemporaryRAMStorageDeviceAttachment] from an objc.ID.
func VZTemporaryRAMStorageDeviceAttachmentFromID(id objc.ID) VZTemporaryRAMStorageDeviceAttachment {
	return VZTemporaryRAMStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}

// Ensure VZTemporaryRAMStorageDeviceAttachment implements IVZTemporaryRAMStorageDeviceAttachment.
var _ IVZTemporaryRAMStorageDeviceAttachment = VZTemporaryRAMStorageDeviceAttachment{}

// An interface definition for the [VZTemporaryRAMStorageDeviceAttachment] class.
//
// # Methods
//
//   - [IVZTemporaryRAMStorageDeviceAttachment.URL]
//   - [IVZTemporaryRAMStorageDeviceAttachment._getAttachmentWithQueueCompletionHandler]
//   - [IVZTemporaryRAMStorageDeviceAttachment.IsReadOnly]
//   - [IVZTemporaryRAMStorageDeviceAttachment.InitWithURLReadOnlyError]
//   - [IVZTemporaryRAMStorageDeviceAttachment.ReadOnly]
type IVZTemporaryRAMStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Methods

	URL() foundation.NSURL
	_getAttachmentWithQueueCompletionHandler(queue dispatch.Queue, handler ErrorHandler)
	IsReadOnly() bool
	InitWithURLReadOnlyError(url foundation.NSURL, only bool) (VZTemporaryRAMStorageDeviceAttachment, error)
	ReadOnly() bool
}

// Init initializes the instance.
func (v VZTemporaryRAMStorageDeviceAttachment) Init() VZTemporaryRAMStorageDeviceAttachment {
	rv := objc.Send[VZTemporaryRAMStorageDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZTemporaryRAMStorageDeviceAttachment) Autorelease() VZTemporaryRAMStorageDeviceAttachment {
	rv := objc.Send[VZTemporaryRAMStorageDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZTemporaryRAMStorageDeviceAttachment creates a new VZTemporaryRAMStorageDeviceAttachment instance.
func NewVZTemporaryRAMStorageDeviceAttachment() VZTemporaryRAMStorageDeviceAttachment {
	class := getVZTemporaryRAMStorageDeviceAttachmentClass()
	rv := objc.Send[VZTemporaryRAMStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZTemporaryRAMStorageDeviceAttachmentWithURLReadOnlyError(url foundation.NSURL, only bool) (VZTemporaryRAMStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZTemporaryRAMStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:readOnly:error:"), url, only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZTemporaryRAMStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZTemporaryRAMStorageDeviceAttachmentFromID(rv), nil
}

func (v VZTemporaryRAMStorageDeviceAttachment) _getAttachmentWithQueueCompletionHandler(queue dispatch.Queue, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_getAttachmentWithQueue:completionHandler:"), uintptr(queue.Handle()), _block1)
}

// GetAttachmentWithQueueCompletionHandler is an exported wrapper for the private method _getAttachmentWithQueueCompletionHandler.
func (v VZTemporaryRAMStorageDeviceAttachment) GetAttachmentWithQueueCompletionHandler(queue dispatch.Queue, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_getAttachmentWithQueue:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_getAttachmentWithQueue:completionHandler:"}
		return err
	}
	v._getAttachmentWithQueueCompletionHandler(queue, handler)
	return nil
}

// CanGetAttachmentWithQueueCompletionHandler reports whether the receiver responds to the private selector _getAttachmentWithQueue:completionHandler:.
func (v VZTemporaryRAMStorageDeviceAttachment) CanGetAttachmentWithQueueCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_getAttachmentWithQueue:completionHandler:"))
}
func (v VZTemporaryRAMStorageDeviceAttachment) IsReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isReadOnly"))
	return rv
}
func (v VZTemporaryRAMStorageDeviceAttachment) InitWithURLReadOnlyError(url foundation.NSURL, only bool) (VZTemporaryRAMStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithURL:readOnly:error:"), url, only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZTemporaryRAMStorageDeviceAttachment), foundation.NSErrorFrom(errorPtr)
	}
	return VZTemporaryRAMStorageDeviceAttachmentFromID(rv), nil

}

func (v VZTemporaryRAMStorageDeviceAttachment) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (v VZTemporaryRAMStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("readOnly"))
	return rv
}

// _getAttachmentWithQueue is a synchronous wrapper around [VZTemporaryRAMStorageDeviceAttachment._getAttachmentWithQueueCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZTemporaryRAMStorageDeviceAttachment) _getAttachmentWithQueue(ctx context.Context, queue dispatch.Queue) error {
	done := make(chan error, 1)
	v._getAttachmentWithQueueCompletionHandler(queue, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
