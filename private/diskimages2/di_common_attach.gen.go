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

// The class instance for the [DICommonAttach] class.
var (
	_DICommonAttachClass     DICommonAttachClass
	_DICommonAttachClassOnce sync.Once
)

func getDICommonAttachClass() DICommonAttachClass {
	_DICommonAttachClassOnce.Do(func() {
		_DICommonAttachClass = DICommonAttachClass{class: objc.GetClass("DICommonAttach")}
	})
	return _DICommonAttachClass
}

// GetDICommonAttachClass returns the class object for DICommonAttach.
func GetDICommonAttachClass() DICommonAttachClass {
	return getDICommonAttachClass()
}

type DICommonAttachClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DICommonAttachClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICommonAttachClass) Alloc() DICommonAttach {
	rv := objc.Send[DICommonAttach](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach
type DICommonAttach struct {
	objectivec.Object
}

// DICommonAttachFromID constructs a [DICommonAttach] from an objc.ID.
func DICommonAttachFromID(id objc.ID) DICommonAttach {
	return DICommonAttach{objectivec.Object{ID: id}}
}
// Ensure DICommonAttach implements IDICommonAttach.
var _ IDICommonAttach = DICommonAttach{}

// An interface definition for the [DICommonAttach] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach
type IDICommonAttach interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DICommonAttach) Init() DICommonAttach {
	rv := objc.Send[DICommonAttach](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICommonAttach) Autorelease() DICommonAttach {
	rv := objc.Send[DICommonAttach](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICommonAttach creates a new DICommonAttach instance.
func NewDICommonAttach() DICommonAttach {
	class := getDICommonAttachClass()
	rv := objc.Send[DICommonAttach](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/diskImageAttach:readOnly:autoMount:BSDName:error:
func (_DICommonAttachClass DICommonAttachClass) DiskImageAttachReadOnlyAutoMountBSDNameError(url foundation.INSURL, readOnly bool, autoMount bool, bsdName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("diskImageAttach:readOnly:autoMount:BSDName:error:"), url, readOnly, autoMount, objc.String(bsdName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("diskImageAttach:readOnly:autoMount:BSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/diskImageAttach:BSDName:error:
func (_DICommonAttachClass DICommonAttachClass) DiskImageAttachBSDNameError(url foundation.INSURL, bsdName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("diskImageAttach:BSDName:error:"), url, objc.String(bsdName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("diskImageAttach:BSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/defaultDiskImageAttach:BSDName:error:
func (_DICommonAttachClass DICommonAttachClass) DefaultDiskImageAttachBSDNameError(url foundation.INSURL, bsdName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("defaultDiskImageAttach:BSDName:error:"), url, objc.String(bsdName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("defaultDiskImageAttach:BSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/DI1_attachWithDictionary:BSDName:error:
func (_DICommonAttachClass DICommonAttachClass) DI1_attachWithDictionaryBSDNameError(dictionary objectivec.IObject, sDName []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("DI1_attachWithDictionary:BSDName:error:"), dictionary, objectivec.IObjectSliceToNSArray(sDName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("DI1_attachWithDictionary:BSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/DI2_attachWithParams:BSDName:error:
func (_DICommonAttachClass DICommonAttachClass) DI2_attachWithParamsBSDNameError(params objectivec.IObject, sDName []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("DI2_attachWithParams:BSDName:error:"), params, objectivec.IObjectSliceToNSArray(sDName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("DI2_attachWithParams:BSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICommonAttach/failWithDI1errorCode:error:
func (_DICommonAttachClass DICommonAttachClass) FailWithDI1errorCodeError(code int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICommonAttachClass.class), objc.Sel("failWithDI1errorCode:error:"), code, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithDI1errorCode:error: returned NO with nil NSError")
	}
	return rv, nil

}

