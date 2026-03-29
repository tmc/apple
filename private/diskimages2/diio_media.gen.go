// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIIOMedia] class.
var (
	_DIIOMediaClass     DIIOMediaClass
	_DIIOMediaClassOnce sync.Once
)

func getDIIOMediaClass() DIIOMediaClass {
	_DIIOMediaClassOnce.Do(func() {
		_DIIOMediaClass = DIIOMediaClass{class: objc.GetClass("DIIOMedia")}
	})
	return _DIIOMediaClass
}

// GetDIIOMediaClass returns the class object for DIIOMedia.
func GetDIIOMediaClass() DIIOMediaClass {
	return getDIIOMediaClass()
}

type DIIOMediaClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIIOMediaClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIIOMediaClass) Alloc() DIIOMedia {
	rv := objc.Send[DIIOMedia](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIIOMedia.BSDName]
//   - [DIIOMedia.CopyBlockDeviceWithError]
//   - [DIIOMedia.InitWithDevNameError]
// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia
type DIIOMedia struct {
	DIIOObject
}

// DIIOMediaFromID constructs a [DIIOMedia] from an objc.ID.
func DIIOMediaFromID(id objc.ID) DIIOMedia {
	return DIIOMedia{DIIOObject: DIIOObjectFromID(id)}
}
// Ensure DIIOMedia implements IDIIOMedia.
var _ IDIIOMedia = DIIOMedia{}

// An interface definition for the [DIIOMedia] class.
//
// # Methods
//
//   - [IDIIOMedia.BSDName]
//   - [IDIIOMedia.CopyBlockDeviceWithError]
//   - [IDIIOMedia.InitWithDevNameError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia
type IDIIOMedia interface {
	IDIIOObject

	// Topic: Methods

	BSDName() string
	CopyBlockDeviceWithError() (objectivec.IObject, error)
	InitWithDevNameError(name objectivec.IObject) (DIIOMedia, error)
}

// Init initializes the instance.
func (d DIIOMedia) Init() DIIOMedia {
	rv := objc.Send[DIIOMedia](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIIOMedia) Autorelease() DIIOMedia {
	rv := objc.Send[DIIOMedia](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIIOMedia creates a new DIIOMedia instance.
func NewDIIOMedia() DIIOMedia {
	class := getDIIOMediaClass()
	rv := objc.Send[DIIOMedia](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithClassName:error:
func NewDIIOMediaWithClassNameError(name objectivec.IObject) (DIIOMedia, error) {
	var errorPtr objc.ID
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOMedia{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOMediaFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithDIIOObject:
func NewDIIOMediaWithDIIOObject(dIIOObject objectivec.IObject) DIIOMedia {
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDIIOObject:"), dIIOObject)
	return DIIOMediaFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia/initWithDevName:error:
func NewDIIOMediaWithDevNameError(name objectivec.IObject) (DIIOMedia, error) {
	var errorPtr objc.ID
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOMedia{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOMediaFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:
func NewDIIOMediaWithIOObject(iOObject uint32) DIIOMedia {
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:"), iOObject)
	return DIIOMediaFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:retain:
func NewDIIOMediaWithIOObjectRetain(iOObject uint32, retain bool) DIIOMedia {
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:retain:"), iOObject, retain)
	return DIIOMediaFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIteratorNext:
func NewDIIOMediaWithIteratorNext(next objectivec.IObject) DIIOMedia {
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIteratorNext:"), next)
	return DIIOMediaFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithRegistryEntryID:error:
func NewDIIOMediaWithRegistryEntryIDError(id uint64) (DIIOMedia, error) {
	var errorPtr objc.ID
	instance := getDIIOMediaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegistryEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOMedia{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOMediaFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia/copyBlockDeviceWithError:
func (d DIIOMedia) CopyBlockDeviceWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyBlockDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia/initWithDevName:error:
func (d DIIOMedia) InitWithDevNameError(name objectivec.IObject) (DIIOMedia, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithDevName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOMedia{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOMediaFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOMedia/BSDName
func (d DIIOMedia) BSDName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("BSDName"))
	return foundation.NSStringFromID(rv).String()
}

