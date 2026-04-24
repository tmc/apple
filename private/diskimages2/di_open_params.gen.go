// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIOpenParams] class.
var (
	_DIOpenParamsClass     DIOpenParamsClass
	_DIOpenParamsClassOnce sync.Once
)

func getDIOpenParamsClass() DIOpenParamsClass {
	_DIOpenParamsClassOnce.Do(func() {
		_DIOpenParamsClass = DIOpenParamsClass{class: objc.GetClass("DIOpenParams")}
	})
	return _DIOpenParamsClass
}

// GetDIOpenParamsClass returns the class object for DIOpenParams.
func GetDIOpenParamsClass() DIOpenParamsClass {
	return getDIOpenParamsClass()
}

type DIOpenParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIOpenParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIOpenParamsClass) Alloc() DIOpenParams {
	rv := objc.Send[DIOpenParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIOpenParams.UIOOpenMode]
//   - [DIOpenParams.OpenWithError]
//   - [DIOpenParams.UnlockImageWithOpenParams]
//   - [DIOpenParams.InitWithURLOpenModeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams
type DIOpenParams struct {
	DIBaseParams
}

// DIOpenParamsFromID constructs a [DIOpenParams] from an objc.ID.
func DIOpenParamsFromID(id objc.ID) DIOpenParams {
	return DIOpenParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIOpenParams implements IDIOpenParams.
var _ IDIOpenParams = DIOpenParams{}

// An interface definition for the [DIOpenParams] class.
//
// # Methods
//
//   - [IDIOpenParams.UIOOpenMode]
//   - [IDIOpenParams.OpenWithError]
//   - [IDIOpenParams.UnlockImageWithOpenParams]
//   - [IDIOpenParams.InitWithURLOpenModeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams
type IDIOpenParams interface {
	IDIBaseParams

	// Topic: Methods

	UIOOpenMode() int
	OpenWithError() (objectivec.IObject, error)
	UnlockImageWithOpenParams(params unsafe.Pointer) objectivec.IObject
	InitWithURLOpenModeError(url foundation.INSURL, mode int64) (DIOpenParams, error)
}

// Init initializes the instance.
func (d DIOpenParams) Init() DIOpenParams {
	rv := objc.Send[DIOpenParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIOpenParams) Autorelease() DIOpenParams {
	rv := objc.Send[DIOpenParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIOpenParams creates a new DIOpenParams instance.
func NewDIOpenParams() DIOpenParams {
	class := getDIOpenParamsClass()
	rv := objc.Send[DIOpenParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithCoder:
func NewDIOpenParamsWithCoder(coder objectivec.IObject) DIOpenParams {
	instance := getDIOpenParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIOpenParamsFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithURL:error:
func NewDIOpenParamsWithURLError(url foundation.INSURL) (DIOpenParams, error) {
	var errorPtr objc.ID
	instance := getDIOpenParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIOpenParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIOpenParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams/initWithURL:openMode:error:
func NewDIOpenParamsWithURLOpenModeError(url foundation.INSURL, mode int64) (DIOpenParams, error) {
	var errorPtr objc.ID
	instance := getDIOpenParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:openMode:error:"), url, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIOpenParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIOpenParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams/openWithError:
func (d DIOpenParams) OpenWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("openWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams/unlockImageWithOpenParams:
func (d DIOpenParams) UnlockImageWithOpenParams(params unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("unlockImageWithOpenParams:"), params)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams/initWithURL:openMode:error:
func (d DIOpenParams) InitWithURLOpenModeError(url foundation.INSURL, mode int64) (DIOpenParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:openMode:error:"), url, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIOpenParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIOpenParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIOpenParams/UIOOpenMode
func (d DIOpenParams) UIOOpenMode() int {
	rv := objc.Send[int](d.ID, objc.Sel("UIOOpenMode"))
	return rv
}
