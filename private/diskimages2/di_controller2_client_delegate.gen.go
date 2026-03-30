// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIController2ClientDelegate] class.
var (
	_DIController2ClientDelegateClass     DIController2ClientDelegateClass
	_DIController2ClientDelegateClassOnce sync.Once
)

func getDIController2ClientDelegateClass() DIController2ClientDelegateClass {
	_DIController2ClientDelegateClassOnce.Do(func() {
		_DIController2ClientDelegateClass = DIController2ClientDelegateClass{class: objc.GetClass("DIController2ClientDelegate")}
	})
	return _DIController2ClientDelegateClass
}

// GetDIController2ClientDelegateClass returns the class object for DIController2ClientDelegate.
func GetDIController2ClientDelegateClass() DIController2ClientDelegateClass {
	return getDIController2ClientDelegateClass()
}

type DIController2ClientDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIController2ClientDelegateClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIController2ClientDelegateClass) Alloc() DIController2ClientDelegate {
	rv := objc.Send[DIController2ClientDelegate](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [DIController2ClientDelegate.DeviceHandle]
//   - [DIController2ClientDelegate.SetDeviceHandle]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientDelegate
type DIController2ClientDelegate struct {
	objectivec.Object
}

// DIController2ClientDelegateFromID constructs a [DIController2ClientDelegate] from an objc.ID.
func DIController2ClientDelegateFromID(id objc.ID) DIController2ClientDelegate {
	return DIController2ClientDelegate{objectivec.Object{ID: id}}
}

// Ensure DIController2ClientDelegate implements IDIController2ClientDelegate.
var _ IDIController2ClientDelegate = DIController2ClientDelegate{}

// An interface definition for the [DIController2ClientDelegate] class.
//
// # Methods
//
//   - [IDIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [IDIController2ClientDelegate.DeviceHandle]
//   - [IDIController2ClientDelegate.SetDeviceHandle]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientDelegate
type IDIController2ClientDelegate interface {
	objectivec.IObject

	// Topic: Methods

	AttachCompletedWithHandleReply(handle objectivec.IObject, reply VoidHandler)
	DeviceHandle() IDIDeviceHandle
	SetDeviceHandle(value IDIDeviceHandle)
}

// Init initializes the instance.
func (d DIController2ClientDelegate) Init() DIController2ClientDelegate {
	rv := objc.Send[DIController2ClientDelegate](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIController2ClientDelegate) Autorelease() DIController2ClientDelegate {
	rv := objc.Send[DIController2ClientDelegate](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIController2ClientDelegate creates a new DIController2ClientDelegate instance.
func NewDIController2ClientDelegate() DIController2ClientDelegate {
	class := getDIController2ClientDelegateClass()
	rv := objc.Send[DIController2ClientDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientDelegate/attachCompletedWithHandle:reply:
func (d DIController2ClientDelegate) AttachCompletedWithHandleReply(handle objectivec.IObject, reply VoidHandler) {
	_block1, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](d.ID, objc.Sel("attachCompletedWithHandle:reply:"), handle, _block1)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientDelegate/deviceHandle
func (d DIController2ClientDelegate) DeviceHandle() IDIDeviceHandle {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("deviceHandle"))
	return DIDeviceHandleFromID(objc.ID(rv))
}
func (d DIController2ClientDelegate) SetDeviceHandle(value IDIDeviceHandle) {
	objc.Send[struct{}](d.ID, objc.Sel("setDeviceHandle:"), value)
}

// AttachCompletedWithHandleReplySync is a synchronous wrapper around [DIController2ClientDelegate.AttachCompletedWithHandleReply].
// It blocks until the completion handler fires or the context is cancelled.
func (d DIController2ClientDelegate) AttachCompletedWithHandleReplySync(ctx context.Context, handle objectivec.IObject) error {
	done := make(chan struct{}, 1)
	d.AttachCompletedWithHandleReply(handle, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
