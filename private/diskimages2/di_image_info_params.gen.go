// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [DIImageInfoParams] class.
var (
	_DIImageInfoParamsClass     DIImageInfoParamsClass
	_DIImageInfoParamsClassOnce sync.Once
)

func getDIImageInfoParamsClass() DIImageInfoParamsClass {
	_DIImageInfoParamsClassOnce.Do(func() {
		_DIImageInfoParamsClass = DIImageInfoParamsClass{class: objc.GetClass("DIImageInfoParams")}
	})
	return _DIImageInfoParamsClass
}

// GetDIImageInfoParamsClass returns the class object for DIImageInfoParams.
func GetDIImageInfoParamsClass() DIImageInfoParamsClass {
	return getDIImageInfoParamsClass()
}

type DIImageInfoParamsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIImageInfoParamsClass) Alloc() DIImageInfoParams {
	rv := objc.Send[DIImageInfoParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIImageInfoParams.ImageInfo]
//   - [DIImageInfoParams.SetImageInfo]
//   - [DIImageInfoParams.ExtraInfo]
//   - [DIImageInfoParams.SetExtraInfo]
//   - [DIImageInfoParams.EncryptionInfoOnly]
//   - [DIImageInfoParams.SetEncryptionInfoOnly]
//   - [DIImageInfoParams.OpenEncryption]
//   - [DIImageInfoParams.SetOpenEncryption]
//   - [DIImageInfoParams.RetrieveWithError]
//   - [DIImageInfoParams.InitWithURLError]
//   - [DIImageInfoParams.InitWithExistingParamsError]
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams
type DIImageInfoParams struct {
	DIBaseParams
}

// DIImageInfoParamsFromID constructs a [DIImageInfoParams] from an objc.ID.
func DIImageInfoParamsFromID(id objc.ID) DIImageInfoParams {
	return DIImageInfoParams{DIBaseParams: DIBaseParamsFromID(id)}
}
// Ensure DIImageInfoParams implements IDIImageInfoParams.
var _ IDIImageInfoParams = DIImageInfoParams{}

// An interface definition for the [DIImageInfoParams] class.
//
// # Methods
//
//   - [IDIImageInfoParams.ImageInfo]
//   - [IDIImageInfoParams.SetImageInfo]
//   - [IDIImageInfoParams.ExtraInfo]
//   - [IDIImageInfoParams.SetExtraInfo]
//   - [IDIImageInfoParams.EncryptionInfoOnly]
//   - [IDIImageInfoParams.SetEncryptionInfoOnly]
//   - [IDIImageInfoParams.OpenEncryption]
//   - [IDIImageInfoParams.SetOpenEncryption]
//   - [IDIImageInfoParams.RetrieveWithError]
//   - [IDIImageInfoParams.InitWithURLError]
//   - [IDIImageInfoParams.InitWithExistingParamsError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams
type IDIImageInfoParams interface {
	IDIBaseParams

	// Topic: Methods

	ImageInfo() foundation.INSDictionary
	SetImageInfo(value foundation.INSDictionary)
	ExtraInfo() bool
	SetExtraInfo(value bool)
	EncryptionInfoOnly() bool
	SetEncryptionInfoOnly(value bool)
	OpenEncryption() bool
	SetOpenEncryption(value bool)
	RetrieveWithError() (bool, error)
	InitWithURLError(url foundation.INSURL) (DIImageInfoParams, error)
	InitWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error)
}

// Init initializes the instance.
func (d DIImageInfoParams) Init() DIImageInfoParams {
	rv := objc.Send[DIImageInfoParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIImageInfoParams) Autorelease() DIImageInfoParams {
	rv := objc.Send[DIImageInfoParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIImageInfoParams creates a new DIImageInfoParams instance.
func NewDIImageInfoParams() DIImageInfoParams {
	class := getDIImageInfoParamsClass()
	rv := objc.Send[DIImageInfoParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/initWithExistingParams:error:
func NewDIImageInfoParamsWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	instance := getDIImageInfoParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParamsFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return DIImageInfoParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/initWithURL:error:
func NewDIImageInfoParamsWithURLError(url foundation.INSURL) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	instance := getDIImageInfoParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParamsFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return DIImageInfoParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/retrieveWithError:
func (d DIImageInfoParams) RetrieveWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("retrieveWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("retrieveWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/initWithURL:error:
func (d DIImageInfoParams) InitWithURLError(url foundation.INSURL) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIImageInfoParamsFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/initWithExistingParams:error:
func (d DIImageInfoParams) InitWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIImageInfoParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/supportsSecureCoding
func (_DIImageInfoParamsClass DIImageInfoParamsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIImageInfoParamsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/isDiskImageWithURL:
func (_DIImageInfoParamsClass DIImageInfoParamsClass) IsDiskImageWithURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](objc.ID(_DIImageInfoParamsClass.class), objc.Sel("isDiskImageWithURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/encryptionInfoOnly
func (d DIImageInfoParams) EncryptionInfoOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("encryptionInfoOnly"))
	return rv
}
func (d DIImageInfoParams) SetEncryptionInfoOnly(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setEncryptionInfoOnly:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/extraInfo
func (d DIImageInfoParams) ExtraInfo() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("extraInfo"))
	return rv
}
func (d DIImageInfoParams) SetExtraInfo(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setExtraInfo:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/imageInfo
func (d DIImageInfoParams) ImageInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("imageInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DIImageInfoParams) SetImageInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setImageInfo:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIImageInfoParams/openEncryption
func (d DIImageInfoParams) OpenEncryption() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("openEncryption"))
	return rv
}
func (d DIImageInfoParams) SetOpenEncryption(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setOpenEncryption:"), value)
}

