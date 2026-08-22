// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImages2] class.
var (
	_DiskImages2Class     DiskImages2Class
	_DiskImages2ClassOnce sync.Once
)

func getDiskImages2Class() DiskImages2Class {
	_DiskImages2ClassOnce.Do(func() {
		_DiskImages2Class = DiskImages2Class{class: objc.GetClass("DiskImages2")}
	})
	return _DiskImages2Class
}

// GetDiskImages2Class returns the class object for DiskImages2.
func GetDiskImages2Class() DiskImages2Class {
	return getDiskImages2Class()
}

type DiskImages2Class struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImages2Class) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImages2Class) Alloc() DiskImages2 {
	rv := objc.SendIfResponds[DiskImages2](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImages2 struct {
	objectivec.Object
}

// DiskImages2FromID constructs a [DiskImages2] from an objc.ID.
func DiskImages2FromID(id objc.ID) DiskImages2 {
	return DiskImages2{objectivec.Object{ID: id}}
}

// Ensure DiskImages2 implements IDiskImages2.
var _ IDiskImages2 = DiskImages2{}

// An interface definition for the [DiskImages2] class.
type IDiskImages2 interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DiskImages2) Init() DiskImages2 {
	rv := objc.SendIfResponds[DiskImages2](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImages2) Autorelease() DiskImages2 {
	rv := objc.SendIfResponds[DiskImages2](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImages2 creates a new DiskImages2 instance.
func NewDiskImages2() DiskImages2 {
	class := getDiskImages2Class()
	rv := objc.SendIfResponds[DiskImages2](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_DiskImages2Class DiskImages2Class) AttachWithParamsHandleError(params IDIAttachParams, handle IDIDeviceHandle) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("attachWithParams:handle:error:"), params, handle, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("attachWithParams:handle:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) ManagedAttachWithParamsHandleError(params IDIAttachParams, handle IDIDeviceHandle) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("managedAttachWithParams:handle:error:"), params, handle, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("managedAttachWithParams:handle:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) CreateBlankWithParamsError(params IDICreateASIFParams) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("createBlankWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createBlankWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) ResizeWithParamsError(params IDIResizeParams) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("resizeWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) RetrieveInfoWithParamsError(params IDIImageInfoParams) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("retrieveInfoWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("retrieveInfoWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) ConvertWithParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("convertWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("convertWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) VerifyWithParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("verifyWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) ChangePasswordWithParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("changePasswordWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("changePasswordWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) EmbedUserDataWithParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("embedUserDataWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("embedUserDataWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) RetrieveUserDataWithParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImages2Class.class), objc.Sel("retrieveUserDataWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("retrieveUserDataWithParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImages2Class DiskImages2Class) ImageURLFromDeviceError(device string) (foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImages2Class.class), objc.Sel("imageURLFromDevice:error:"), objc.String(device), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSURL{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSURLFromID(rv), nil

}
func (_DiskImages2Class DiskImages2Class) IsEncryptedImageWithURLError(url foundation.NSURL) (foundation.NSNumber, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImages2Class.class), objc.Sel("isEncryptedImageWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSNumber{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSNumberFromID(rv), nil

}
func (_DiskImages2Class DiskImages2Class) DebugLogsEnabled() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DiskImages2Class.class), objc.Sel("debugLogsEnabled"))
	return rv
}
func (_DiskImages2Class DiskImages2Class) ForwardLogs() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DiskImages2Class.class), objc.Sel("forwardLogs"))
	return rv
}
func (_DiskImages2Class DiskImages2Class) SetDebugLogsEnabled(enabled bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_DiskImages2Class.class), objc.Sel("setDebugLogsEnabled:"), enabled)
}
func (_DiskImages2Class DiskImages2Class) SetForwardLogs(forwardLogs bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_DiskImages2Class.class), objc.Sel("setForwardLogs:"), forwardLogs)
}
func (_DiskImages2Class DiskImages2Class) ConvertWithParamsCompletionBlock(params objectivec.IObject, block VoidHandler) objectivec.IObject {
	_block1, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[objc.ID](objc.ID(_DiskImages2Class.class), objc.Sel("convertWithParams:completionBlock:"), params, _block1)
	return objectivec.Object{ID: rv}
}

// ConvertWithParamsCompletionBlockSync is a synchronous wrapper around [DiskImages2.ConvertWithParamsCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (dc DiskImages2Class) ConvertWithParamsCompletionBlockSync(ctx context.Context, params objectivec.IObject) error {
	done := make(chan struct{}, 1)
	dc.ConvertWithParamsCompletionBlock(params, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
