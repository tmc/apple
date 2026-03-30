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

// The class instance for the [DIChpassParams] class.
var (
	_DIChpassParamsClass     DIChpassParamsClass
	_DIChpassParamsClassOnce sync.Once
)

func getDIChpassParamsClass() DIChpassParamsClass {
	_DIChpassParamsClassOnce.Do(func() {
		_DIChpassParamsClass = DIChpassParamsClass{class: objc.GetClass("DIChpassParams")}
	})
	return _DIChpassParamsClass
}

// GetDIChpassParamsClass returns the class object for DIChpassParams.
func GetDIChpassParamsClass() DIChpassParamsClass {
	return getDIChpassParamsClass()
}

type DIChpassParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIChpassParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIChpassParamsClass) Alloc() DIChpassParams {
	rv := objc.Send[DIChpassParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIChpassParams.ChangePassWithXpcHandlerError]
//   - [DIChpassParams.ChpassWithError]
//   - [DIChpassParams.PrepareImageWithXpcHandlerFileModeEncChpassError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams
type DIChpassParams struct {
	DIBaseParams
}

// DIChpassParamsFromID constructs a [DIChpassParams] from an objc.ID.
func DIChpassParamsFromID(id objc.ID) DIChpassParams {
	return DIChpassParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIChpassParams implements IDIChpassParams.
var _ IDIChpassParams = DIChpassParams{}

// An interface definition for the [DIChpassParams] class.
//
// # Methods
//
//   - [IDIChpassParams.ChangePassWithXpcHandlerError]
//   - [IDIChpassParams.ChpassWithError]
//   - [IDIChpassParams.PrepareImageWithXpcHandlerFileModeEncChpassError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams
type IDIChpassParams interface {
	IDIBaseParams

	// Topic: Methods

	ChangePassWithXpcHandlerError(handler objectivec.IObject) (bool, error)
	ChpassWithError() (bool, error)
	PrepareImageWithXpcHandlerFileModeEncChpassError(handler objectivec.IObject, mode int64, chpass objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (d DIChpassParams) Init() DIChpassParams {
	rv := objc.Send[DIChpassParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIChpassParams) Autorelease() DIChpassParams {
	rv := objc.Send[DIChpassParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIChpassParams creates a new DIChpassParams instance.
func NewDIChpassParams() DIChpassParams {
	class := getDIChpassParamsClass()
	rv := objc.Send[DIChpassParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithCoder:
func NewDIChpassParamsWithCoder(coder objectivec.IObject) DIChpassParams {
	instance := getDIChpassParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIChpassParamsFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams/initWithURL:error:
func NewDIChpassParamsWithURLError(url foundation.INSURL) (DIChpassParams, error) {
	var errorPtr objc.ID
	instance := getDIChpassParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIChpassParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIChpassParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams/changePassWithXpcHandler:error:
func (d DIChpassParams) ChangePassWithXpcHandlerError(handler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("changePassWithXpcHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("changePassWithXpcHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams/chpassWithError:
func (d DIChpassParams) ChpassWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("chpassWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("chpassWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIChpassParams/prepareImageWithXpcHandler:fileMode:encChpass:error:
func (d DIChpassParams) PrepareImageWithXpcHandlerFileModeEncChpassError(handler objectivec.IObject, mode int64, chpass objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("prepareImageWithXpcHandler:fileMode:encChpass:error:"), handler, mode, chpass, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareImageWithXpcHandler:fileMode:encChpass:error: returned NO with nil NSError")
	}
	return rv, nil

}
