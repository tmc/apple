// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIDeviceHandle] class.
var (
	_DIDeviceHandleClass     DIDeviceHandleClass
	_DIDeviceHandleClassOnce sync.Once
)

func getDIDeviceHandleClass() DIDeviceHandleClass {
	_DIDeviceHandleClassOnce.Do(func() {
		_DIDeviceHandleClass = DIDeviceHandleClass{class: objc.GetClass("DIDeviceHandle")}
	})
	return _DIDeviceHandleClass
}

// GetDIDeviceHandleClass returns the class object for DIDeviceHandle.
func GetDIDeviceHandleClass() DIDeviceHandleClass {
	return getDIDeviceHandleClass()
}

type DIDeviceHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDeviceHandleClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDeviceHandleClass) Alloc() DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIDeviceHandle.BSDName]
//   - [DIDeviceHandle.SetBSDName]
//   - [DIDeviceHandle.RegEntryID]
//   - [DIDeviceHandle.HandleRefCount]
//   - [DIDeviceHandle.SetHandleRefCount]
//   - [DIDeviceHandle.WaitForDeviceWithError]
//   - [DIDeviceHandle.WaitForQuietWithServiceError]
//   - [DIDeviceHandle.AddToRefCountWithError]
//   - [DIDeviceHandle.UpdateBSDNameWithBlockDeviceError]
//   - [DIDeviceHandle.XpcEndpoint]
//   - [DIDeviceHandle.SetXpcEndpoint]
//   - [DIDeviceHandle.Client2IOhandler]
//   - [DIDeviceHandle.SetClient2IOhandler]
//   - [DIDeviceHandle.EncodeWithCoder]
//   - [DIDeviceHandle.InitWithRegEntryID]
//   - [DIDeviceHandle.InitWithRegEntryIDXpcEndpoint]
//   - [DIDeviceHandle.InitWithCoder]
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle
type DIDeviceHandle struct {
	objectivec.Object
}

// DIDeviceHandleFromID constructs a [DIDeviceHandle] from an objc.ID.
func DIDeviceHandleFromID(id objc.ID) DIDeviceHandle {
	return DIDeviceHandle{objectivec.Object{ID: id}}
}
// Ensure DIDeviceHandle implements IDIDeviceHandle.
var _ IDIDeviceHandle = DIDeviceHandle{}

// An interface definition for the [DIDeviceHandle] class.
//
// # Methods
//
//   - [IDIDeviceHandle.BSDName]
//   - [IDIDeviceHandle.SetBSDName]
//   - [IDIDeviceHandle.RegEntryID]
//   - [IDIDeviceHandle.HandleRefCount]
//   - [IDIDeviceHandle.SetHandleRefCount]
//   - [IDIDeviceHandle.WaitForDeviceWithError]
//   - [IDIDeviceHandle.WaitForQuietWithServiceError]
//   - [IDIDeviceHandle.AddToRefCountWithError]
//   - [IDIDeviceHandle.UpdateBSDNameWithBlockDeviceError]
//   - [IDIDeviceHandle.XpcEndpoint]
//   - [IDIDeviceHandle.SetXpcEndpoint]
//   - [IDIDeviceHandle.Client2IOhandler]
//   - [IDIDeviceHandle.SetClient2IOhandler]
//   - [IDIDeviceHandle.EncodeWithCoder]
//   - [IDIDeviceHandle.InitWithRegEntryID]
//   - [IDIDeviceHandle.InitWithRegEntryIDXpcEndpoint]
//   - [IDIDeviceHandle.InitWithCoder]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle
type IDIDeviceHandle interface {
	objectivec.IObject

	// Topic: Methods

	BSDName() string
	SetBSDName(value string)
	RegEntryID() uint64
	HandleRefCount() bool
	SetHandleRefCount(value bool)
	WaitForDeviceWithError() (bool, error)
	WaitForQuietWithServiceError(service uint32) (bool, error)
	AddToRefCountWithError() (bool, error)
	UpdateBSDNameWithBlockDeviceError(blockDevice string) (bool, error)
	XpcEndpoint() foundation.NSXPCListenerEndpoint
	SetXpcEndpoint(value foundation.NSXPCListenerEndpoint)
	Client2IOhandler() IDIClient2IODaemonXPCHandler
	SetClient2IOhandler(value IDIClient2IODaemonXPCHandler)
	EncodeWithCoder(coder foundation.INSCoder)
	InitWithRegEntryID(regEntryID uint64) DIDeviceHandle
	InitWithRegEntryIDXpcEndpoint(regEntryID uint64, xpcEndpoint foundation.NSXPCListenerEndpoint) DIDeviceHandle
	InitWithCoder(coder foundation.INSCoder) DIDeviceHandle
}

// Init initializes the instance.
func (d DIDeviceHandle) Init() DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDeviceHandle) Autorelease() DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDeviceHandle creates a new DIDeviceHandle instance.
func NewDIDeviceHandle() DIDeviceHandle {
	class := getDIDeviceHandleClass()
	rv := objc.Send[DIDeviceHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithCoder:
func NewDIDeviceHandleWithCoder(coder objectivec.IObject) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIDeviceHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithRegEntryID:
func NewDIDeviceHandleWithRegEntryID(regEntryID uint64) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegEntryID:"), regEntryID)
	return DIDeviceHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithRegEntryID:xpcEndpoint:
func NewDIDeviceHandleWithRegEntryIDXpcEndpoint(regEntryID uint64, xpcEndpoint foundation.NSXPCListenerEndpoint) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegEntryID:xpcEndpoint:"), regEntryID, xpcEndpoint)
	return DIDeviceHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/waitForDeviceWithError:
func (d DIDeviceHandle) WaitForDeviceWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("waitForDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("waitForDeviceWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/waitForQuietWithService:error:
func (d DIDeviceHandle) WaitForQuietWithServiceError(service uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("waitForQuietWithService:error:"), service, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("waitForQuietWithService:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/addToRefCountWithError:
func (d DIDeviceHandle) AddToRefCountWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addToRefCountWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addToRefCountWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/updateBSDNameWithBlockDevice:error:
func (d DIDeviceHandle) UpdateBSDNameWithBlockDeviceError(blockDevice string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("updateBSDNameWithBlockDevice:error:"), objc.String(blockDevice), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateBSDNameWithBlockDevice:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/encodeWithCoder:
func (d DIDeviceHandle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithRegEntryID:
func (d DIDeviceHandle) InitWithRegEntryID(regEntryID uint64) DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](d.ID, objc.Sel("initWithRegEntryID:"), regEntryID)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithRegEntryID:xpcEndpoint:
func (d DIDeviceHandle) InitWithRegEntryIDXpcEndpoint(regEntryID uint64, xpcEndpoint foundation.NSXPCListenerEndpoint) DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](d.ID, objc.Sel("initWithRegEntryID:xpcEndpoint:"), regEntryID, xpcEndpoint)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/initWithCoder:
func (d DIDeviceHandle) InitWithCoder(coder foundation.INSCoder) DIDeviceHandle {
	rv := objc.Send[DIDeviceHandle](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/supportsSecureCoding
func (_DIDeviceHandleClass DIDeviceHandleClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIDeviceHandleClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/BSDName
func (d DIDeviceHandle) BSDName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("BSDName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIDeviceHandle) SetBSDName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setBSDName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/handleRefCount
func (d DIDeviceHandle) HandleRefCount() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("handleRefCount"))
	return rv
}
func (d DIDeviceHandle) SetHandleRefCount(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setHandleRefCount:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/regEntryID
func (d DIDeviceHandle) RegEntryID() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("regEntryID"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/xpcEndpoint
func (d DIDeviceHandle) XpcEndpoint() foundation.NSXPCListenerEndpoint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("xpcEndpoint"))
	return foundation.NSXPCListenerEndpointFromID(objc.ID(rv))
}
func (d DIDeviceHandle) SetXpcEndpoint(value foundation.NSXPCListenerEndpoint) {
	objc.Send[struct{}](d.ID, objc.Sel("setXpcEndpoint:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIDeviceHandle/client2IOhandler
func (d DIDeviceHandle) Client2IOhandler() IDIClient2IODaemonXPCHandler {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("client2IOhandler"))
	return DIClient2IODaemonXPCHandlerFromID(objc.ID(rv))
}
func (d DIDeviceHandle) SetClient2IOhandler(value IDIClient2IODaemonXPCHandler) {
	objc.Send[struct{}](d.ID, objc.Sel("setClient2IOhandler:"), value)
}

