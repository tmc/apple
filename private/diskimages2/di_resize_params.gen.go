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

// The class instance for the [DIResizeParams] class.
var (
	_DIResizeParamsClass     DIResizeParamsClass
	_DIResizeParamsClassOnce sync.Once
)

func getDIResizeParamsClass() DIResizeParamsClass {
	_DIResizeParamsClassOnce.Do(func() {
		_DIResizeParamsClass = DIResizeParamsClass{class: objc.GetClass("DIResizeParams")}
	})
	return _DIResizeParamsClass
}

// GetDIResizeParamsClass returns the class object for DIResizeParams.
func GetDIResizeParamsClass() DIResizeParamsClass {
	return getDIResizeParamsClass()
}

type DIResizeParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIResizeParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIResizeParamsClass) Alloc() DIResizeParams {
	rv := objc.Send[DIResizeParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIResizeParams.Size]
//   - [DIResizeParams.SetSize]
//   - [DIResizeParams.ResizeWithError]
//   - [DIResizeParams.InitWithURLSizeError]
//   - [DIResizeParams.InitWithExistingParamsSizeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams
type DIResizeParams struct {
	DIBaseParams
}

// DIResizeParamsFromID constructs a [DIResizeParams] from an objc.ID.
func DIResizeParamsFromID(id objc.ID) DIResizeParams {
	return DIResizeParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIResizeParams implements IDIResizeParams.
var _ IDIResizeParams = DIResizeParams{}

// An interface definition for the [DIResizeParams] class.
//
// # Methods
//
//   - [IDIResizeParams.Size]
//   - [IDIResizeParams.SetSize]
//   - [IDIResizeParams.ResizeWithError]
//   - [IDIResizeParams.InitWithURLSizeError]
//   - [IDIResizeParams.InitWithExistingParamsSizeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams
type IDIResizeParams interface {
	IDIBaseParams

	// Topic: Methods

	Size() uint64
	SetSize(value uint64)
	ResizeWithError() (bool, error)
	InitWithURLSizeError(url foundation.INSURL, size uint64) (DIResizeParams, error)
	InitWithExistingParamsSizeError(params IDIResizeParams, size uint64) (DIResizeParams, error)
}

// Init initializes the instance.
func (d DIResizeParams) Init() DIResizeParams {
	rv := objc.Send[DIResizeParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIResizeParams) Autorelease() DIResizeParams {
	rv := objc.Send[DIResizeParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIResizeParams creates a new DIResizeParams instance.
func NewDIResizeParams() DIResizeParams {
	class := getDIResizeParamsClass()
	rv := objc.Send[DIResizeParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/initWithCoder:
func NewDIResizeParamsWithCoder(coder objectivec.IObject) DIResizeParams {
	instance := getDIResizeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIResizeParamsFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/initWithExistingParams:size:error:
func NewDIResizeParamsWithExistingParamsSizeError(params IDIResizeParams, size uint64) (DIResizeParams, error) {
	var errorPtr objc.ID
	instance := getDIResizeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExistingParams:size:error:"), params, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIResizeParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIResizeParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithURL:error:
func NewDIResizeParamsWithURLError(url foundation.INSURL) (DIResizeParams, error) {
	var errorPtr objc.ID
	instance := getDIResizeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIResizeParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIResizeParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/initWithURL:size:error:
func NewDIResizeParamsWithURLSizeError(url foundation.INSURL, size uint64) (DIResizeParams, error) {
	var errorPtr objc.ID
	instance := getDIResizeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:size:error:"), url, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIResizeParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIResizeParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/resizeWithError:
func (d DIResizeParams) ResizeWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/initWithURL:size:error:
func (d DIResizeParams) InitWithURLSizeError(url foundation.INSURL, size uint64) (DIResizeParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:size:error:"), url, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIResizeParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIResizeParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/initWithExistingParams:size:error:
func (d DIResizeParams) InitWithExistingParamsSizeError(params IDIResizeParams, size uint64) (DIResizeParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithExistingParams:size:error:"), params, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIResizeParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIResizeParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIResizeParams/size
func (d DIResizeParams) Size() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("size"))
	return rv
}
func (d DIResizeParams) SetSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setSize:"), value)
}
