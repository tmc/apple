// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIIOIterator] class.
var (
	_DIIOIteratorClass     DIIOIteratorClass
	_DIIOIteratorClassOnce sync.Once
)

func getDIIOIteratorClass() DIIOIteratorClass {
	_DIIOIteratorClassOnce.Do(func() {
		_DIIOIteratorClass = DIIOIteratorClass{class: objc.GetClass("DIIOIterator")}
	})
	return _DIIOIteratorClass
}

// GetDIIOIteratorClass returns the class object for DIIOIterator.
func GetDIIOIteratorClass() DIIOIteratorClass {
	return getDIIOIteratorClass()
}

type DIIOIteratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIIOIteratorClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIIOIteratorClass) Alloc() DIIOIterator {
	rv := objc.Send[DIIOIterator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIIOIterator.CopyNextObject]
//   - [DIIOIterator.StartedOver]
//   - [DIIOIterator.SetStartedOver]
//   - [DIIOIterator.InitWithIOIteratorRetain]
// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator
type DIIOIterator struct {
	DIIOObject
}

// DIIOIteratorFromID constructs a [DIIOIterator] from an objc.ID.
func DIIOIteratorFromID(id objc.ID) DIIOIterator {
	return DIIOIterator{DIIOObject: DIIOObjectFromID(id)}
}
// Ensure DIIOIterator implements IDIIOIterator.
var _ IDIIOIterator = DIIOIterator{}

// An interface definition for the [DIIOIterator] class.
//
// # Methods
//
//   - [IDIIOIterator.CopyNextObject]
//   - [IDIIOIterator.StartedOver]
//   - [IDIIOIterator.SetStartedOver]
//   - [IDIIOIterator.InitWithIOIteratorRetain]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator
type IDIIOIterator interface {
	IDIIOObject

	// Topic: Methods

	CopyNextObject() uint32
	StartedOver() bool
	SetStartedOver(value bool)
	InitWithIOIteratorRetain(iOIterator uint32, retain bool) DIIOIterator
}

// Init initializes the instance.
func (d DIIOIterator) Init() DIIOIterator {
	rv := objc.Send[DIIOIterator](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIIOIterator) Autorelease() DIIOIterator {
	rv := objc.Send[DIIOIterator](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIIOIterator creates a new DIIOIterator instance.
func NewDIIOIterator() DIIOIterator {
	class := getDIIOIteratorClass()
	rv := objc.Send[DIIOIterator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithClassName:error:
func NewDIIOIteratorWithClassNameError(name objectivec.IObject) (DIIOIterator, error) {
	var errorPtr objc.ID
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOIterator{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOIteratorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithDIIOObject:
func NewDIIOIteratorWithDIIOObject(dIIOObject objectivec.IObject) DIIOIterator {
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDIIOObject:"), dIIOObject)
	return DIIOIteratorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator/initWithIOIterator:retain:
func NewDIIOIteratorWithIOIteratorRetain(iOIterator uint32, retain bool) DIIOIterator {
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOIterator:retain:"), iOIterator, retain)
	return DIIOIteratorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:
func NewDIIOIteratorWithIOObject(iOObject uint32) DIIOIterator {
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:"), iOObject)
	return DIIOIteratorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIOObject:retain:
func NewDIIOIteratorWithIOObjectRetain(iOObject uint32, retain bool) DIIOIterator {
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:retain:"), iOObject, retain)
	return DIIOIteratorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithIteratorNext:
func NewDIIOIteratorWithIteratorNext(next objectivec.IObject) DIIOIterator {
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIteratorNext:"), next)
	return DIIOIteratorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOObject/initWithRegistryEntryID:error:
func NewDIIOIteratorWithRegistryEntryIDError(id uint64) (DIIOIterator, error) {
	var errorPtr objc.ID
	instance := getDIIOIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegistryEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOIterator{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOIteratorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator/copyNextObject
func (d DIIOIterator) CopyNextObject() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("copyNextObject"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator/initWithIOIterator:retain:
func (d DIIOIterator) InitWithIOIteratorRetain(iOIterator uint32, retain bool) DIIOIterator {
	rv := objc.Send[DIIOIterator](d.ID, objc.Sel("initWithIOIterator:retain:"), iOIterator, retain)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIIOIterator/startedOver
func (d DIIOIterator) StartedOver() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("startedOver"))
	return rv
}
func (d DIIOIterator) SetStartedOver(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setStartedOver:"), value)
}

