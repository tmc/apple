// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIIOObject] class.
var (
	_DIIOObjectClass     DIIOObjectClass
	_DIIOObjectClassOnce sync.Once
)

func getDIIOObjectClass() DIIOObjectClass {
	_DIIOObjectClassOnce.Do(func() {
		_DIIOObjectClass = DIIOObjectClass{class: objc.GetClass("DIIOObject")}
	})
	return _DIIOObjectClass
}

// GetDIIOObjectClass returns the class object for DIIOObject.
func GetDIIOObjectClass() DIIOObjectClass {
	return getDIIOObjectClass()
}

type DIIOObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIIOObjectClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIIOObjectClass) Alloc() DIIOObject {
	rv := objc.Send[DIIOObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIIOObject.CopyParentWithError]
//   - [DIIOObject.CopyPropertyWithClassKey]
//   - [DIIOObject.IoClassName]
//   - [DIIOObject.IoObj]
//   - [DIIOObject.IoObjectWithClassNameIterateParentError]
//   - [DIIOObject.NewIteratorWithOptionsError]
//   - [DIIOObject.RegistryEntryIDWithError]
//   - [DIIOObject.InitWithClassNameError]
//   - [DIIOObject.InitWithDIIOObject]
//   - [DIIOObject.InitWithIOObject]
//   - [DIIOObject.InitWithIOObjectRetain]
//   - [DIIOObject.InitWithIteratorNext]
//   - [DIIOObject.InitWithRegistryEntryIDError]
type DIIOObject struct {
	objectivec.Object
}

// DIIOObjectFromID constructs a [DIIOObject] from an objc.ID.
func DIIOObjectFromID(id objc.ID) DIIOObject {
	return DIIOObject{objectivec.Object{ID: id}}
}

// Ensure DIIOObject implements IDIIOObject.
var _ IDIIOObject = DIIOObject{}

// An interface definition for the [DIIOObject] class.
//
// # Methods
//
//   - [IDIIOObject.CopyParentWithError]
//   - [IDIIOObject.CopyPropertyWithClassKey]
//   - [IDIIOObject.IoClassName]
//   - [IDIIOObject.IoObj]
//   - [IDIIOObject.IoObjectWithClassNameIterateParentError]
//   - [IDIIOObject.NewIteratorWithOptionsError]
//   - [IDIIOObject.RegistryEntryIDWithError]
//   - [IDIIOObject.InitWithClassNameError]
//   - [IDIIOObject.InitWithDIIOObject]
//   - [IDIIOObject.InitWithIOObject]
//   - [IDIIOObject.InitWithIOObjectRetain]
//   - [IDIIOObject.InitWithIteratorNext]
//   - [IDIIOObject.InitWithRegistryEntryIDError]
type IDIIOObject interface {
	objectivec.IObject

	// Topic: Methods

	CopyParentWithError() (objectivec.IObject, error)
	CopyPropertyWithClassKey(class objectivec.Class, key objectivec.IObject) objectivec.IObject
	IoClassName() string
	IoObj() uint32
	IoObjectWithClassNameIterateParentError(name objectivec.IObject, parent bool) (objectivec.IObject, error)
	NewIteratorWithOptionsError(options uint32) (objectivec.IObject, error)
	RegistryEntryIDWithError() (uint64, error)
	InitWithClassNameError(name objectivec.IObject) (DIIOObject, error)
	InitWithDIIOObject(dIIOObject objectivec.IObject) DIIOObject
	InitWithIOObject(iOObject uint32) DIIOObject
	InitWithIOObjectRetain(iOObject uint32, retain bool) DIIOObject
	InitWithIteratorNext(next objectivec.IObject) DIIOObject
	InitWithRegistryEntryIDError(id uint64) (DIIOObject, error)
}

// Init initializes the instance.
func (d DIIOObject) Init() DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIIOObject) Autorelease() DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIIOObject creates a new DIIOObject instance.
func NewDIIOObject() DIIOObject {
	class := getDIIOObjectClass()
	rv := objc.Send[DIIOObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIIOObjectWithClassNameError(name objectivec.IObject) (DIIOObject, error) {
	var errorPtr objc.ID
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOObject{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOObjectFromID(rv), nil
}

func NewDIIOObjectWithDIIOObject(dIIOObject objectivec.IObject) DIIOObject {
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDIIOObject:"), dIIOObject)
	return DIIOObjectFromID(rv)
}

func NewDIIOObjectWithIOObject(iOObject uint32) DIIOObject {
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:"), iOObject)
	return DIIOObjectFromID(rv)
}

func NewDIIOObjectWithIOObjectRetain(iOObject uint32, retain bool) DIIOObject {
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOObject:retain:"), iOObject, retain)
	return DIIOObjectFromID(rv)
}

func NewDIIOObjectWithIteratorNext(next objectivec.IObject) DIIOObject {
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIteratorNext:"), next)
	return DIIOObjectFromID(rv)
}

func NewDIIOObjectWithRegistryEntryIDError(id uint64) (DIIOObject, error) {
	var errorPtr objc.ID
	instance := getDIIOObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegistryEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOObject{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOObjectFromID(rv), nil
}

func (d DIIOObject) CopyParentWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyParentWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIIOObject) CopyPropertyWithClassKey(class objectivec.Class, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyPropertyWithClass:key:"), class, key)
	return objectivec.Object{ID: rv}
}
func (d DIIOObject) IoObjectWithClassNameIterateParentError(name objectivec.IObject, parent bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ioObjectWithClassName:iterateParent:error:"), name, parent, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIIOObject) NewIteratorWithOptionsError(options uint32) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("newIteratorWithOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIIOObject) RegistryEntryIDWithError() (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](d.ID, objc.Sel("registryEntryIDWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (d DIIOObject) InitWithClassNameError(name objectivec.IObject) (DIIOObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithClassName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOObject{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOObjectFromID(rv), nil

}
func (d DIIOObject) InitWithDIIOObject(dIIOObject objectivec.IObject) DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("initWithDIIOObject:"), dIIOObject)
	return rv
}
func (d DIIOObject) InitWithIOObject(iOObject uint32) DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("initWithIOObject:"), iOObject)
	return rv
}
func (d DIIOObject) InitWithIOObjectRetain(iOObject uint32, retain bool) DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("initWithIOObject:retain:"), iOObject, retain)
	return rv
}
func (d DIIOObject) InitWithIteratorNext(next objectivec.IObject) DIIOObject {
	rv := objc.Send[DIIOObject](d.ID, objc.Sel("initWithIteratorNext:"), next)
	return rv
}
func (d DIIOObject) InitWithRegistryEntryIDError(id uint64) (DIIOObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithRegistryEntryID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIIOObject{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIIOObjectFromID(rv), nil

}

func (_DIIOObjectClass DIIOObjectClass) CopyDiskImagesControllerWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIIOObjectClass.class), objc.Sel("copyDiskImagesControllerWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (d DIIOObject) IoClassName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ioClassName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIIOObject) IoObj() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("ioObj"))
	return rv
}
