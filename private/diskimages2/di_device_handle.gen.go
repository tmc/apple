// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[DIDeviceHandle](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

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
	rv := objc.SendIfResponds[DIDeviceHandle](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDeviceHandle) Autorelease() DIDeviceHandle {
	rv := objc.SendIfResponds[DIDeviceHandle](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDeviceHandle creates a new DIDeviceHandle instance.
func NewDIDeviceHandle() DIDeviceHandle {
	class := getDIDeviceHandleClass()
	rv := objc.SendIfResponds[DIDeviceHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIDeviceHandleWithCoder(coder objectivec.IObject) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIDeviceHandleFromID(rv)
}

func NewDIDeviceHandleWithRegEntryID(regEntryID uint64) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRegEntryID:"), regEntryID)
	return DIDeviceHandleFromID(rv)
}

func NewDIDeviceHandleWithRegEntryIDXpcEndpoint(regEntryID uint64, xpcEndpoint foundation.NSXPCListenerEndpoint) DIDeviceHandle {
	instance := getDIDeviceHandleClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRegEntryID:xpcEndpoint:"), regEntryID, xpcEndpoint)
	return DIDeviceHandleFromID(rv)
}

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
func (d DIDeviceHandle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (d DIDeviceHandle) InitWithRegEntryID(regEntryID uint64) DIDeviceHandle {
	rv := objc.SendIfResponds[DIDeviceHandle](d.ID, objc.Sel("initWithRegEntryID:"), regEntryID)
	return rv
}
func (d DIDeviceHandle) InitWithRegEntryIDXpcEndpoint(regEntryID uint64, xpcEndpoint foundation.NSXPCListenerEndpoint) DIDeviceHandle {
	rv := objc.SendIfResponds[DIDeviceHandle](d.ID, objc.Sel("initWithRegEntryID:xpcEndpoint:"), regEntryID, xpcEndpoint)
	return rv
}
func (d DIDeviceHandle) InitWithCoder(coder foundation.INSCoder) DIDeviceHandle {
	rv := objc.SendIfResponds[DIDeviceHandle](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_DIDeviceHandleClass DIDeviceHandleClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DIDeviceHandleClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (d DIDeviceHandle) BSDName() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("BSDName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIDeviceHandle) SetBSDName(value string) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setBSDName:"), objc.String(value))
}
func (d DIDeviceHandle) HandleRefCount() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("handleRefCount"))
	return rv
}
func (d DIDeviceHandle) SetHandleRefCount(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setHandleRefCount:"), value)
}
func (d DIDeviceHandle) RegEntryID() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("regEntryID"))
	return rv
}
func (d DIDeviceHandle) XpcEndpoint() foundation.NSXPCListenerEndpoint {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("xpcEndpoint"))
	return foundation.NSXPCListenerEndpointFromID(objc.ID(rv))
}
func (d DIDeviceHandle) SetXpcEndpoint(value foundation.NSXPCListenerEndpoint) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setXpcEndpoint:"), value)
}
func (d DIDeviceHandle) Client2IOhandler() IDIClient2IODaemonXPCHandler {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("client2IOhandler"))
	return DIClient2IODaemonXPCHandlerFromID(objc.ID(rv))
}
func (d DIDeviceHandle) SetClient2IOhandler(value IDIClient2IODaemonXPCHandler) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setClient2IOhandler:"), value)
}
