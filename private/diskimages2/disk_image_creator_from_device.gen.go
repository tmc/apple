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

// The class instance for the [DiskImageCreatorFromDevice] class.
var (
	_DiskImageCreatorFromDeviceClass     DiskImageCreatorFromDeviceClass
	_DiskImageCreatorFromDeviceClassOnce sync.Once
)

func getDiskImageCreatorFromDeviceClass() DiskImageCreatorFromDeviceClass {
	_DiskImageCreatorFromDeviceClassOnce.Do(func() {
		_DiskImageCreatorFromDeviceClass = DiskImageCreatorFromDeviceClass{class: objc.GetClass("DiskImageCreatorFromDevice")}
	})
	return _DiskImageCreatorFromDeviceClass
}

// GetDiskImageCreatorFromDeviceClass returns the class object for DiskImageCreatorFromDevice.
func GetDiskImageCreatorFromDeviceClass() DiskImageCreatorFromDeviceClass {
	return getDiskImageCreatorFromDeviceClass()
}

type DiskImageCreatorFromDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageCreatorFromDeviceClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageCreatorFromDeviceClass) Alloc() DiskImageCreatorFromDevice {
	rv := objc.SendIfResponds[DiskImageCreatorFromDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DiskImageCreatorFromDevice.CreateImageWithSrcDeviceError]
//   - [DiskImageCreatorFromDevice.InitWithURLError]
type DiskImageCreatorFromDevice struct {
	BaseDiskImageCreator
}

// DiskImageCreatorFromDeviceFromID constructs a [DiskImageCreatorFromDevice] from an objc.ID.
func DiskImageCreatorFromDeviceFromID(id objc.ID) DiskImageCreatorFromDevice {
	return DiskImageCreatorFromDevice{BaseDiskImageCreator: BaseDiskImageCreatorFromID(id)}
}

// Ensure DiskImageCreatorFromDevice implements IDiskImageCreatorFromDevice.
var _ IDiskImageCreatorFromDevice = DiskImageCreatorFromDevice{}

// An interface definition for the [DiskImageCreatorFromDevice] class.
//
// # Methods
//
//   - [IDiskImageCreatorFromDevice.CreateImageWithSrcDeviceError]
//   - [IDiskImageCreatorFromDevice.InitWithURLError]
type IDiskImageCreatorFromDevice interface {
	IBaseDiskImageCreator

	// Topic: Methods

	CreateImageWithSrcDeviceError(device objectivec.IObject) (bool, error)
	InitWithURLError(url foundation.NSURL) (DiskImageCreatorFromDevice, error)
}

// Init initializes the instance.
func (d DiskImageCreatorFromDevice) Init() DiskImageCreatorFromDevice {
	rv := objc.SendIfResponds[DiskImageCreatorFromDevice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageCreatorFromDevice) Autorelease() DiskImageCreatorFromDevice {
	rv := objc.SendIfResponds[DiskImageCreatorFromDevice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageCreatorFromDevice creates a new DiskImageCreatorFromDevice instance.
func NewDiskImageCreatorFromDevice() DiskImageCreatorFromDevice {
	class := getDiskImageCreatorFromDeviceClass()
	rv := objc.SendIfResponds[DiskImageCreatorFromDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageCreatorFromDeviceWithURLDefaultFormatError(url foundation.NSURL, format int64) (DiskImageCreatorFromDevice, error) {
	var errorPtr objc.ID
	instance := getDiskImageCreatorFromDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DiskImageCreatorFromDevice{}, objc.ErrInitFailed
	}
	return DiskImageCreatorFromDeviceFromID(rv), nil
}

func NewDiskImageCreatorFromDeviceWithURLError(url foundation.NSURL) (DiskImageCreatorFromDevice, error) {
	var errorPtr objc.ID
	instance := getDiskImageCreatorFromDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DiskImageCreatorFromDevice{}, objc.ErrInitFailed
	}
	return DiskImageCreatorFromDeviceFromID(rv), nil
}

func (d DiskImageCreatorFromDevice) CreateImageWithSrcDeviceError(device objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createImageWithSrcDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createImageWithSrcDevice:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DiskImageCreatorFromDevice) InitWithURLError(url foundation.NSURL) (DiskImageCreatorFromDevice, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageCreatorFromDeviceFromID(rv), nil

}
