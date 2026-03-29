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

// The class instance for the [DIVerifyParams] class.
var (
	_DIVerifyParamsClass     DIVerifyParamsClass
	_DIVerifyParamsClassOnce sync.Once
)

func getDIVerifyParamsClass() DIVerifyParamsClass {
	_DIVerifyParamsClassOnce.Do(func() {
		_DIVerifyParamsClass = DIVerifyParamsClass{class: objc.GetClass("DIVerifyParams")}
	})
	return _DIVerifyParamsClass
}

// GetDIVerifyParamsClass returns the class object for DIVerifyParams.
func GetDIVerifyParamsClass() DIVerifyParamsClass {
	return getDIVerifyParamsClass()
}

type DIVerifyParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIVerifyParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIVerifyParamsClass) Alloc() DIVerifyParams {
	rv := objc.Send[DIVerifyParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIVerifyParams.ShouldValidateShadows]
//   - [DIVerifyParams.SetShouldValidateShadows]
//   - [DIVerifyParams.VerifyWithError]
//   - [DIVerifyParams.InitWithURLShadowURLsError]
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams
type DIVerifyParams struct {
	DIBaseParams
}

// DIVerifyParamsFromID constructs a [DIVerifyParams] from an objc.ID.
func DIVerifyParamsFromID(id objc.ID) DIVerifyParams {
	return DIVerifyParams{DIBaseParams: DIBaseParamsFromID(id)}
}
// Ensure DIVerifyParams implements IDIVerifyParams.
var _ IDIVerifyParams = DIVerifyParams{}

// An interface definition for the [DIVerifyParams] class.
//
// # Methods
//
//   - [IDIVerifyParams.ShouldValidateShadows]
//   - [IDIVerifyParams.SetShouldValidateShadows]
//   - [IDIVerifyParams.VerifyWithError]
//   - [IDIVerifyParams.InitWithURLShadowURLsError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams
type IDIVerifyParams interface {
	IDIBaseParams

	// Topic: Methods

	ShouldValidateShadows() bool
	SetShouldValidateShadows(value bool)
	VerifyWithError() (bool, error)
	InitWithURLShadowURLsError(url foundation.INSURL, uRLs objectivec.IObject) (DIVerifyParams, error)
}

// Init initializes the instance.
func (d DIVerifyParams) Init() DIVerifyParams {
	rv := objc.Send[DIVerifyParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIVerifyParams) Autorelease() DIVerifyParams {
	rv := objc.Send[DIVerifyParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIVerifyParams creates a new DIVerifyParams instance.
func NewDIVerifyParams() DIVerifyParams {
	class := getDIVerifyParamsClass()
	rv := objc.Send[DIVerifyParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithCoder:
func NewDIVerifyParamsWithCoder(coder objectivec.IObject) DIVerifyParams {
	instance := getDIVerifyParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIVerifyParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams/initWithURL:error:
func NewDIVerifyParamsWithURLError(url foundation.INSURL) (DIVerifyParams, error) {
	var errorPtr objc.ID
	instance := getDIVerifyParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIVerifyParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIVerifyParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams/initWithURL:shadowURLs:error:
func NewDIVerifyParamsWithURLShadowURLsError(url foundation.INSURL, uRLs objectivec.IObject) (DIVerifyParams, error) {
	var errorPtr objc.ID
	instance := getDIVerifyParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:shadowURLs:error:"), url, uRLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIVerifyParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIVerifyParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams/verifyWithError:
func (d DIVerifyParams) VerifyWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("verifyWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams/initWithURL:shadowURLs:error:
func (d DIVerifyParams) InitWithURLShadowURLsError(url foundation.INSURL, uRLs objectivec.IObject) (DIVerifyParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:shadowURLs:error:"), url, uRLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIVerifyParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIVerifyParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIVerifyParams/shouldValidateShadows
func (d DIVerifyParams) ShouldValidateShadows() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldValidateShadows"))
	return rv
}
func (d DIVerifyParams) SetShouldValidateShadows(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setShouldValidateShadows:"), value)
}

