// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEFISignatureList] class.
var (
	_VZEFISignatureListClass     VZEFISignatureListClass
	_VZEFISignatureListClassOnce sync.Once
)

func getVZEFISignatureListClass() VZEFISignatureListClass {
	_VZEFISignatureListClassOnce.Do(func() {
		_VZEFISignatureListClass = VZEFISignatureListClass{class: objc.GetClass("VZEFISignatureList")}
	})
	return _VZEFISignatureListClass
}

// GetVZEFISignatureListClass returns the class object for VZEFISignatureList.
func GetVZEFISignatureListClass() VZEFISignatureListClass {
	return getVZEFISignatureListClass()
}

type VZEFISignatureListClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFISignatureListClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFISignatureListClass) Alloc() VZEFISignatureList {
	rv := objc.SendIfResponds[VZEFISignatureList](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFISignatureList.Signatures]
//   - [VZEFISignatureList.InitWithContentsOfURLError]
//   - [VZEFISignatureList.InitWithSignatures]
type VZEFISignatureList struct {
	objectivec.Object
}

// VZEFISignatureListFromID constructs a [VZEFISignatureList] from an objc.ID.
func VZEFISignatureListFromID(id objc.ID) VZEFISignatureList {
	return VZEFISignatureList{objectivec.Object{ID: id}}
}

// Ensure VZEFISignatureList implements IVZEFISignatureList.
var _ IVZEFISignatureList = VZEFISignatureList{}

// An interface definition for the [VZEFISignatureList] class.
//
// # Methods
//
//   - [IVZEFISignatureList.Signatures]
//   - [IVZEFISignatureList.InitWithContentsOfURLError]
//   - [IVZEFISignatureList.InitWithSignatures]
type IVZEFISignatureList interface {
	objectivec.IObject

	// Topic: Methods

	Signatures() foundation.INSArray
	InitWithContentsOfURLError(url foundation.NSURL) (VZEFISignatureList, error)
	InitWithSignatures(signatures objectivec.IObject) VZEFISignatureList
}

// Init initializes the instance.
func (v VZEFISignatureList) Init() VZEFISignatureList {
	rv := objc.SendIfResponds[VZEFISignatureList](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFISignatureList) Autorelease() VZEFISignatureList {
	rv := objc.SendIfResponds[VZEFISignatureList](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFISignatureList creates a new VZEFISignatureList instance.
func NewVZEFISignatureList() VZEFISignatureList {
	class := getVZEFISignatureListClass()
	rv := objc.SendIfResponds[VZEFISignatureList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZEFISignatureListWithContentsOfURLError(url foundation.NSURL) (VZEFISignatureList, error) {
	var errorPtr objc.ID
	instance := getVZEFISignatureListClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZEFISignatureList{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZEFISignatureList{}, objc.ErrInitFailed
	}
	return VZEFISignatureListFromID(rv), nil
}

func NewVZEFISignatureListWithSignatures(signatures objectivec.IObject) VZEFISignatureList {
	instance := getVZEFISignatureListClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSignatures:"), signatures)
	return VZEFISignatureListFromID(rv)
}

func (v VZEFISignatureList) InitWithContentsOfURLError(url foundation.NSURL) (VZEFISignatureList, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZEFISignatureList{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZEFISignatureListFromID(rv), nil

}
func (v VZEFISignatureList) InitWithSignatures(signatures objectivec.IObject) VZEFISignatureList {
	rv := objc.SendIfResponds[VZEFISignatureList](v.ID, objc.Sel("initWithSignatures:"), signatures)
	return rv
}

func (_VZEFISignatureListClass VZEFISignatureListClass) SignatureListFromSignatures(signatures unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZEFISignatureListClass.class), objc.Sel("signatureListFromSignatures:"), signatures)
	return objectivec.Object{ID: rv}
}

func (v VZEFISignatureList) Signatures() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("signatures"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
