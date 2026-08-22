// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETOpaqueCopy] class.
var (
	_ETOpaqueCopyClass     ETOpaqueCopyClass
	_ETOpaqueCopyClassOnce sync.Once
)

func getETOpaqueCopyClass() ETOpaqueCopyClass {
	_ETOpaqueCopyClassOnce.Do(func() {
		_ETOpaqueCopyClass = ETOpaqueCopyClass{class: objc.GetClass("ETOpaqueCopy")}
	})
	return _ETOpaqueCopyClass
}

// GetETOpaqueCopyClass returns the class object for ETOpaqueCopy.
func GetETOpaqueCopyClass() ETOpaqueCopyClass {
	return getETOpaqueCopyClass()
}

type ETOpaqueCopyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETOpaqueCopyClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETOpaqueCopyClass) Alloc() ETOpaqueCopy {
	rv := objc.SendIfResponds[ETOpaqueCopy](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETOpaqueCopy.GetBlob]
//   - [ETOpaqueCopy.InitWithAbstractBlobContainer]
type ETOpaqueCopy struct {
	objectivec.Object
}

// ETOpaqueCopyFromID constructs a [ETOpaqueCopy] from an objc.ID.
func ETOpaqueCopyFromID(id objc.ID) ETOpaqueCopy {
	return ETOpaqueCopy{objectivec.Object{ID: id}}
}

// Ensure ETOpaqueCopy implements IETOpaqueCopy.
var _ IETOpaqueCopy = ETOpaqueCopy{}

// An interface definition for the [ETOpaqueCopy] class.
//
// # Methods
//
//   - [IETOpaqueCopy.GetBlob]
//   - [IETOpaqueCopy.InitWithAbstractBlobContainer]
type IETOpaqueCopy interface {
	objectivec.IObject

	// Topic: Methods

	GetBlob() unsafe.Pointer
	InitWithAbstractBlobContainer(container unsafe.Pointer) ETOpaqueCopy
}

// Init initializes the instance.
func (e ETOpaqueCopy) Init() ETOpaqueCopy {
	rv := objc.SendIfResponds[ETOpaqueCopy](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETOpaqueCopy) Autorelease() ETOpaqueCopy {
	rv := objc.SendIfResponds[ETOpaqueCopy](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETOpaqueCopy creates a new ETOpaqueCopy instance.
func NewETOpaqueCopy() ETOpaqueCopy {
	class := getETOpaqueCopyClass()
	rv := objc.SendIfResponds[ETOpaqueCopy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewETOpaqueCopyWithAbstractBlobContainer(container unsafe.Pointer) ETOpaqueCopy {
	instance := getETOpaqueCopyClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAbstractBlobContainer:"), container)
	return ETOpaqueCopyFromID(rv)
}

func (e ETOpaqueCopy) GetBlob() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("getBlob"))
	return rv
}
func (e ETOpaqueCopy) InitWithAbstractBlobContainer(container unsafe.Pointer) ETOpaqueCopy {
	rv := objc.SendIfResponds[ETOpaqueCopy](e.ID, objc.Sel("initWithAbstractBlobContainer:"), container)
	return rv
}
