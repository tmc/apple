// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DICreateRAWParams] class.
var (
	_DICreateRAWParamsClass     DICreateRAWParamsClass
	_DICreateRAWParamsClassOnce sync.Once
)

func getDICreateRAWParamsClass() DICreateRAWParamsClass {
	_DICreateRAWParamsClassOnce.Do(func() {
		_DICreateRAWParamsClass = DICreateRAWParamsClass{class: objc.GetClass("DICreateRAWParams")}
	})
	return _DICreateRAWParamsClass
}

// GetDICreateRAWParamsClass returns the class object for DICreateRAWParams.
func GetDICreateRAWParamsClass() DICreateRAWParamsClass {
	return getDICreateRAWParamsClass()
}

type DICreateRAWParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DICreateRAWParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICreateRAWParamsClass) Alloc() DICreateRAWParams {
	rv := objc.Send[DICreateRAWParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DICreateRAWParams
type DICreateRAWParams struct {
	DICreateParams
}

// DICreateRAWParamsFromID constructs a [DICreateRAWParams] from an objc.ID.
func DICreateRAWParamsFromID(id objc.ID) DICreateRAWParams {
	return DICreateRAWParams{DICreateParams: DICreateParamsFromID(id)}
}
// Ensure DICreateRAWParams implements IDICreateRAWParams.
var _ IDICreateRAWParams = DICreateRAWParams{}

// An interface definition for the [DICreateRAWParams] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateRAWParams
type IDICreateRAWParams interface {
	IDICreateParams
}

// Init initializes the instance.
func (d DICreateRAWParams) Init() DICreateRAWParams {
	rv := objc.Send[DICreateRAWParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICreateRAWParams) Autorelease() DICreateRAWParams {
	rv := objc.Send[DICreateRAWParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICreateRAWParams creates a new DICreateRAWParams instance.
func NewDICreateRAWParams() DICreateRAWParams {
	class := getDICreateRAWParamsClass()
	rv := objc.Send[DICreateRAWParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithCoder:
func NewDICreateRAWParamsWithCoder(coder objectivec.IObject) DICreateRAWParams {
	instance := getDICreateRAWParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DICreateRAWParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:error:
func NewDICreateRAWParamsWithURLError(url foundation.INSURL) (DICreateRAWParams, error) {
	var errorPtr objc.ID
	instance := getDICreateRAWParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateRAWParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateRAWParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateRAWParams/initWithURL:numBlocks:error:
func NewDICreateRAWParamsWithURLNumBlocksError(url foundation.INSURL, blocks uint64) (DICreateRAWParams, error) {
	var errorPtr objc.ID
	instance := getDICreateRAWParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:numBlocks:error:"), url, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateRAWParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateRAWParamsFromID(rv), nil
}

