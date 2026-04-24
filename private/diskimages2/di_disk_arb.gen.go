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

// The class instance for the [DIDiskArb] class.
var (
	_DIDiskArbClass     DIDiskArbClass
	_DIDiskArbClassOnce sync.Once
)

func getDIDiskArbClass() DIDiskArbClass {
	_DIDiskArbClassOnce.Do(func() {
		_DIDiskArbClass = DIDiskArbClass{class: objc.GetClass("DIDiskArb")}
	})
	return _DIDiskArbClass
}

// GetDIDiskArbClass returns the class object for DIDiskArb.
func GetDIDiskArbClass() DIDiskArbClass {
	return getDIDiskArbClass()
}

type DIDiskArbClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDiskArbClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDiskArbClass) Alloc() DIDiskArb {
	rv := objc.Send[DIDiskArb](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIDiskArb.AddDisappearedCallbackWithMountPointShadowMountPointsDelegate]
//   - [DIDiskArb.CallbackReached]
//   - [DIDiskArb.SetCallbackReached]
//   - [DIDiskArb.CopyDescriptionWithBSDName]
//   - [DIDiskArb.DaSession]
//   - [DIDiskArb.SetDaSession]
//   - [DIDiskArb.Delegate]
//   - [DIDiskArb.SetDelegate]
//   - [DIDiskArb.EjectWithBSDNameError]
//   - [DIDiskArb.InputMountedOnURL]
//   - [DIDiskArb.SetInputMountedOnURL]
//   - [DIDiskArb.MountWithDeviceNameArgsFilesystemMountURLError]
//   - [DIDiskArb.OnDiskDisappearedWithDisk]
//   - [DIDiskArb.OperationError]
//   - [DIDiskArb.SetOperationError]
//   - [DIDiskArb.ShadowMountedOnURLs]
//   - [DIDiskArb.SetShadowMountedOnURLs]
//   - [DIDiskArb.Stop]
//   - [DIDiskArb.UnmountWithMountPointError]
//   - [DIDiskArb.WaitForDAIdleWithError]
//   - [DIDiskArb.InitWithError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb
type DIDiskArb struct {
	objectivec.Object
}

// DIDiskArbFromID constructs a [DIDiskArb] from an objc.ID.
func DIDiskArbFromID(id objc.ID) DIDiskArb {
	return DIDiskArb{objectivec.Object{ID: id}}
}

// Ensure DIDiskArb implements IDIDiskArb.
var _ IDIDiskArb = DIDiskArb{}

// An interface definition for the [DIDiskArb] class.
//
// # Methods
//
//   - [IDIDiskArb.AddDisappearedCallbackWithMountPointShadowMountPointsDelegate]
//   - [IDIDiskArb.CallbackReached]
//   - [IDIDiskArb.SetCallbackReached]
//   - [IDIDiskArb.CopyDescriptionWithBSDName]
//   - [IDIDiskArb.DaSession]
//   - [IDIDiskArb.SetDaSession]
//   - [IDIDiskArb.Delegate]
//   - [IDIDiskArb.SetDelegate]
//   - [IDIDiskArb.EjectWithBSDNameError]
//   - [IDIDiskArb.InputMountedOnURL]
//   - [IDIDiskArb.SetInputMountedOnURL]
//   - [IDIDiskArb.MountWithDeviceNameArgsFilesystemMountURLError]
//   - [IDIDiskArb.OnDiskDisappearedWithDisk]
//   - [IDIDiskArb.OperationError]
//   - [IDIDiskArb.SetOperationError]
//   - [IDIDiskArb.ShadowMountedOnURLs]
//   - [IDIDiskArb.SetShadowMountedOnURLs]
//   - [IDIDiskArb.Stop]
//   - [IDIDiskArb.UnmountWithMountPointError]
//   - [IDIDiskArb.WaitForDAIdleWithError]
//   - [IDIDiskArb.InitWithError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb
type IDIDiskArb interface {
	objectivec.IObject

	// Topic: Methods

	AddDisappearedCallbackWithMountPointShadowMountPointsDelegate(point objectivec.IObject, points objectivec.IObject, delegate objectivec.IObject)
	CallbackReached() bool
	SetCallbackReached(value bool)
	CopyDescriptionWithBSDName(bSDName objectivec.IObject) objectivec.IObject
	DaSession() DASessionRef
	SetDaSession(value DASessionRef)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	EjectWithBSDNameError(bSDName objectivec.IObject) (bool, error)
	InputMountedOnURL() foundation.INSURL
	SetInputMountedOnURL(value foundation.INSURL)
	MountWithDeviceNameArgsFilesystemMountURLError(name objectivec.IObject, args objectivec.IObject, filesystem objectivec.IObject, url foundation.INSURL) (bool, error)
	OnDiskDisappearedWithDisk(disk DADiskRef)
	OperationError() foundation.INSError
	SetOperationError(value foundation.INSError)
	ShadowMountedOnURLs() foundation.INSArray
	SetShadowMountedOnURLs(value foundation.INSArray)
	Stop()
	UnmountWithMountPointError(point objectivec.IObject) (bool, error)
	WaitForDAIdleWithError() (bool, error)
	InitWithError() (DIDiskArb, error)
}

// Init initializes the instance.
func (d DIDiskArb) Init() DIDiskArb {
	rv := objc.Send[DIDiskArb](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDiskArb) Autorelease() DIDiskArb {
	rv := objc.Send[DIDiskArb](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDiskArb creates a new DIDiskArb instance.
func NewDIDiskArb() DIDiskArb {
	class := getDIDiskArbClass()
	rv := objc.Send[DIDiskArb](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/initWithError:
func NewDIDiskArbWithError() (DIDiskArb, error) {
	var errorPtr objc.ID
	instance := getDIDiskArbClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDiskArb{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIDiskArbFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/addDisappearedCallbackWithMountPoint:shadowMountPoints:delegate:
func (d DIDiskArb) AddDisappearedCallbackWithMountPointShadowMountPointsDelegate(point objectivec.IObject, points objectivec.IObject, delegate objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("addDisappearedCallbackWithMountPoint:shadowMountPoints:delegate:"), point, points, delegate)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/copyDescriptionWithBSDName:
func (d DIDiskArb) CopyDescriptionWithBSDName(bSDName objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyDescriptionWithBSDName:"), bSDName)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/ejectWithBSDName:error:
func (d DIDiskArb) EjectWithBSDNameError(bSDName objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("ejectWithBSDName:error:"), bSDName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("ejectWithBSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/mountWithDeviceName:args:filesystem:mountURL:error:
func (d DIDiskArb) MountWithDeviceNameArgsFilesystemMountURLError(name objectivec.IObject, args objectivec.IObject, filesystem objectivec.IObject, url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("mountWithDeviceName:args:filesystem:mountURL:error:"), name, args, filesystem, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("mountWithDeviceName:args:filesystem:mountURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/onDiskDisappearedWithDisk:
func (d DIDiskArb) OnDiskDisappearedWithDisk(disk DADiskRef) {
	objc.Send[objc.ID](d.ID, objc.Sel("onDiskDisappearedWithDisk:"), disk)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/stop
func (d DIDiskArb) Stop() {
	objc.Send[objc.ID](d.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/unmountWithMountPoint:error:
func (d DIDiskArb) UnmountWithMountPointError(point objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unmountWithMountPoint:error:"), point, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unmountWithMountPoint:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/waitForDAIdleWithError:
func (d DIDiskArb) WaitForDAIdleWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("waitForDAIdleWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("waitForDAIdleWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/initWithError:
func (d DIDiskArb) InitWithError() (DIDiskArb, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDiskArb{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIDiskArbFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/diskArbWithError:
func (_DIDiskArbClass DIDiskArbClass) DiskArbWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIDiskArbClass.class), objc.Sel("diskArbWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/callbackReached
func (d DIDiskArb) CallbackReached() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("callbackReached"))
	return rv
}
func (d DIDiskArb) SetCallbackReached(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setCallbackReached:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/daSession
func (d DIDiskArb) DaSession() DASessionRef {
	rv := objc.Send[DASessionRef](d.ID, objc.Sel("daSession"))
	return DASessionRef(rv)
}
func (d DIDiskArb) SetDaSession(value DASessionRef) {
	objc.Send[struct{}](d.ID, objc.Sel("setDaSession:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/delegate
func (d DIDiskArb) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (d DIDiskArb) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/inputMountedOnURL
func (d DIDiskArb) InputMountedOnURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputMountedOnURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d DIDiskArb) SetInputMountedOnURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setInputMountedOnURL:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/operationError
func (d DIDiskArb) OperationError() foundation.INSError {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("operationError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (d DIDiskArb) SetOperationError(value foundation.INSError) {
	objc.Send[struct{}](d.ID, objc.Sel("setOperationError:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/shadowMountedOnURLs
func (d DIDiskArb) ShadowMountedOnURLs() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("shadowMountedOnURLs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIDiskArb) SetShadowMountedOnURLs(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setShadowMountedOnURLs:"), value)
}
