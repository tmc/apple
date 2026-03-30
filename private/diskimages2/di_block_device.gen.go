// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIBlockDevice] class.
var (
	_DIBlockDeviceClass     DIBlockDeviceClass
	_DIBlockDeviceClassOnce sync.Once
)

func getDIBlockDeviceClass() DIBlockDeviceClass {
	_DIBlockDeviceClassOnce.Do(func() {
		_DIBlockDeviceClass = DIBlockDeviceClass{class: objc.GetClass("DIBlockDevice")}
	})
	return _DIBlockDeviceClass
}

// GetDIBlockDeviceClass returns the class object for DIBlockDevice.
func GetDIBlockDeviceClass() DIBlockDeviceClass {
	return getDIBlockDeviceClass()
}

type DIBlockDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBlockDeviceClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBlockDeviceClass) Alloc() DIBlockDevice {
	rv := objc.Send[DIBlockDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIBlockDevice.CopyIOMediaWithError]
//   - [DIBlockDevice.CopyRootBlockDeviceWithError]
//   - [DIBlockDevice.DiskImageDevice]
//   - [DIBlockDevice.MediumType]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice
type DIBlockDevice struct {
	DIIOObject
}

// DIBlockDeviceFromID constructs a [DIBlockDevice] from an objc.ID.
func DIBlockDeviceFromID(id objc.ID) DIBlockDevice {
	return DIBlockDevice{DIIOObject: DIIOObjectFromID(id)}
}

// Ensure DIBlockDevice implements IDIBlockDevice.
var _ IDIBlockDevice = DIBlockDevice{}

// An interface definition for the [DIBlockDevice] class.
//
// # Methods
//
//   - [IDIBlockDevice.CopyIOMediaWithError]
//   - [IDIBlockDevice.CopyRootBlockDeviceWithError]
//   - [IDIBlockDevice.DiskImageDevice]
//   - [IDIBlockDevice.MediumType]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice
type IDIBlockDevice interface {
	IDIIOObject

	// Topic: Methods

	CopyIOMediaWithError() (objectivec.IObject, error)
	CopyRootBlockDeviceWithError() (objectivec.IObject, error)
	DiskImageDevice() bool
	MediumType() string
}

// Init initializes the instance.
func (d DIBlockDevice) Init() DIBlockDevice {
	rv := objc.Send[DIBlockDevice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBlockDevice) Autorelease() DIBlockDevice {
	rv := objc.Send[DIBlockDevice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBlockDevice creates a new DIBlockDevice instance.
func NewDIBlockDevice() DIBlockDevice {
	class := getDIBlockDeviceClass()
	rv := objc.Send[DIBlockDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithClassName:error:
func NewDIBlockDeviceWithClassNameError(name objectivec.IObject) (DIBlockDevice, error) {
	var errorPtr objc.ID
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIBlockDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIBlockDeviceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithDIIOObject:
func NewDIBlockDeviceWithDIIOObject(dIIOObject objectivec.IObject) DIBlockDevice {
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDIIOObject:"), dIIOObject)
	return DIBlockDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:
func NewDIBlockDeviceWithIOObject(iOObject uint32) DIBlockDevice {
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:"), iOObject)
	return DIBlockDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:retain:
func NewDIBlockDeviceWithIOObjectRetain(iOObject uint32, retain bool) DIBlockDevice {
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:retain:"), iOObject, retain)
	return DIBlockDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIteratorNext:
func NewDIBlockDeviceWithIteratorNext(next objectivec.IObject) DIBlockDevice {
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIteratorNext:"), next)
	return DIBlockDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithRegistryEntryID:error:
func NewDIBlockDeviceWithRegistryEntryIDError(id uint64) (DIBlockDevice, error) {
	var errorPtr objc.ID
	instance := getDIBlockDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegistryEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIBlockDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIBlockDeviceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice/copyIOMediaWithError:
func (d DIBlockDevice) CopyIOMediaWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyIOMediaWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice/copyRootBlockDeviceWithError:
func (d DIBlockDevice) CopyRootBlockDeviceWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyRootBlockDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice/copyUnmatchedDiskImageWithRegEntryID:error:
func (_DIBlockDeviceClass DIBlockDeviceClass) CopyUnmatchedDiskImageWithRegEntryIDError(id uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIBlockDeviceClass.class), objc.Sel("copyUnmatchedDiskImageWithRegEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice/diskImageDevice
func (d DIBlockDevice) DiskImageDevice() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("diskImageDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBlockDevice/mediumType
func (d DIBlockDevice) MediumType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mediumType"))
	return foundation.NSStringFromID(rv).String()
}
