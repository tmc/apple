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

// The class instance for the [DIUserDataParams] class.
var (
	_DIUserDataParamsClass     DIUserDataParamsClass
	_DIUserDataParamsClassOnce sync.Once
)

func getDIUserDataParamsClass() DIUserDataParamsClass {
	_DIUserDataParamsClassOnce.Do(func() {
		_DIUserDataParamsClass = DIUserDataParamsClass{class: objc.GetClass("DIUserDataParams")}
	})
	return _DIUserDataParamsClass
}

// GetDIUserDataParamsClass returns the class object for DIUserDataParams.
func GetDIUserDataParamsClass() DIUserDataParamsClass {
	return getDIUserDataParamsClass()
}

type DIUserDataParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIUserDataParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIUserDataParamsClass) Alloc() DIUserDataParams {
	rv := objc.Send[DIUserDataParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIUserDataParams.EmbedWithError]
//   - [DIUserDataParams.RetrieveWithError]
//   - [DIUserDataParams.UserDict]
//   - [DIUserDataParams.SetUserDict]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams
type DIUserDataParams struct {
	DIBaseParams
}

// DIUserDataParamsFromID constructs a [DIUserDataParams] from an objc.ID.
func DIUserDataParamsFromID(id objc.ID) DIUserDataParams {
	return DIUserDataParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIUserDataParams implements IDIUserDataParams.
var _ IDIUserDataParams = DIUserDataParams{}

// An interface definition for the [DIUserDataParams] class.
//
// # Methods
//
//   - [IDIUserDataParams.EmbedWithError]
//   - [IDIUserDataParams.RetrieveWithError]
//   - [IDIUserDataParams.UserDict]
//   - [IDIUserDataParams.SetUserDict]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams
type IDIUserDataParams interface {
	IDIBaseParams

	// Topic: Methods

	EmbedWithError() (bool, error)
	RetrieveWithError() (bool, error)
	UserDict() foundation.INSDictionary
	SetUserDict(value foundation.INSDictionary)
}

// Init initializes the instance.
func (d DIUserDataParams) Init() DIUserDataParams {
	rv := objc.Send[DIUserDataParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIUserDataParams) Autorelease() DIUserDataParams {
	rv := objc.Send[DIUserDataParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIUserDataParams creates a new DIUserDataParams instance.
func NewDIUserDataParams() DIUserDataParams {
	class := getDIUserDataParamsClass()
	rv := objc.Send[DIUserDataParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams/initWithCoder:
func NewDIUserDataParamsWithCoder(coder objectivec.IObject) DIUserDataParams {
	instance := getDIUserDataParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIUserDataParamsFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams/initWithURL:error:
func NewDIUserDataParamsWithURLError(url foundation.INSURL) (DIUserDataParams, error) {
	var errorPtr objc.ID
	instance := getDIUserDataParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIUserDataParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIUserDataParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams/embedWithError:
func (d DIUserDataParams) EmbedWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("embedWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("embedWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams/retrieveWithError:
func (d DIUserDataParams) RetrieveWithError() (bool, error) {
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

// See: https://developer.apple.com/documentation/DiskImages2/DIUserDataParams/userDict
func (d DIUserDataParams) UserDict() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("userDict"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DIUserDataParams) SetUserDict(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setUserDict:"), value)
}
