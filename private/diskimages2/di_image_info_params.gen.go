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

// Class returns the underlying Objective-C class pointer.
func (dc DIImageInfoParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIImageInfoParamsClass) Alloc() DIImageInfoParams {
	rv := objc.SendIfResponds[DIImageInfoParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

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
//   - [DIImageInfoParams.InitWithExistingParamsError]
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
//   - [IDIImageInfoParams.InitWithExistingParamsError]
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
	InitWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error)
}

// Init initializes the instance.
func (d DIImageInfoParams) Init() DIImageInfoParams {
	rv := objc.SendIfResponds[DIImageInfoParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIImageInfoParams) Autorelease() DIImageInfoParams {
	rv := objc.SendIfResponds[DIImageInfoParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIImageInfoParams creates a new DIImageInfoParams instance.
func NewDIImageInfoParams() DIImageInfoParams {
	class := getDIImageInfoParamsClass()
	rv := objc.SendIfResponds[DIImageInfoParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIImageInfoParamsWithCoder(coder objectivec.IObject) DIImageInfoParams {
	instance := getDIImageInfoParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIImageInfoParamsFromID(rv)
}

func NewDIImageInfoParamsWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	instance := getDIImageInfoParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIImageInfoParams{}, objc.ErrInitFailed
	}
	return DIImageInfoParamsFromID(rv), nil
}

func NewDIImageInfoParamsWithURLError(url foundation.NSURL) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	instance := getDIImageInfoParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIImageInfoParams{}, objc.ErrInitFailed
	}
	return DIImageInfoParamsFromID(rv), nil
}

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
func (d DIImageInfoParams) InitWithExistingParamsError(params IDIImageInfoParams) (DIImageInfoParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIImageInfoParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIImageInfoParamsFromID(rv), nil

}

func (_DIImageInfoParamsClass DIImageInfoParamsClass) IsDiskImageWithURL(url foundation.NSURL) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DIImageInfoParamsClass.class), objc.Sel("isDiskImageWithURL:"), url)
	return rv
}

func (d DIImageInfoParams) EncryptionInfoOnly() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("encryptionInfoOnly"))
	return rv
}
func (d DIImageInfoParams) SetEncryptionInfoOnly(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setEncryptionInfoOnly:"), value)
}
func (d DIImageInfoParams) ExtraInfo() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("extraInfo"))
	return rv
}
func (d DIImageInfoParams) SetExtraInfo(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setExtraInfo:"), value)
}
func (d DIImageInfoParams) ImageInfo() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("imageInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DIImageInfoParams) SetImageInfo(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setImageInfo:"), value)
}
func (d DIImageInfoParams) OpenEncryption() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("openEncryption"))
	return rv
}
func (d DIImageInfoParams) SetOpenEncryption(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setOpenEncryption:"), value)
}
