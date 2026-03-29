// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DICreateASIFParams] class.
var (
	_DICreateASIFParamsClass     DICreateASIFParamsClass
	_DICreateASIFParamsClassOnce sync.Once
)

func getDICreateASIFParamsClass() DICreateASIFParamsClass {
	_DICreateASIFParamsClassOnce.Do(func() {
		_DICreateASIFParamsClass = DICreateASIFParamsClass{class: objc.GetClass("DICreateASIFParams")}
	})
	return _DICreateASIFParamsClass
}

// GetDICreateASIFParamsClass returns the class object for DICreateASIFParams.
func GetDICreateASIFParamsClass() DICreateASIFParamsClass {
	return getDICreateASIFParamsClass()
}

type DICreateASIFParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DICreateASIFParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICreateASIFParamsClass) Alloc() DICreateASIFParams {
	rv := objc.Send[DICreateASIFParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DICreateASIFParams
type DICreateASIFParams struct {
	DICreateParams
}

// DICreateASIFParamsFromID constructs a [DICreateASIFParams] from an objc.ID.
func DICreateASIFParamsFromID(id objc.ID) DICreateASIFParams {
	return DICreateASIFParams{DICreateParams: DICreateParamsFromID(id)}
}
// Ensure DICreateASIFParams implements IDICreateASIFParams.
var _ IDICreateASIFParams = DICreateASIFParams{}

// An interface definition for the [DICreateASIFParams] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateASIFParams
type IDICreateASIFParams interface {
	IDICreateParams
}

// Init initializes the instance.
func (d DICreateASIFParams) Init() DICreateASIFParams {
	rv := objc.Send[DICreateASIFParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICreateASIFParams) Autorelease() DICreateASIFParams {
	rv := objc.Send[DICreateASIFParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICreateASIFParams creates a new DICreateASIFParams instance.
func NewDICreateASIFParams() DICreateASIFParams {
	class := getDICreateASIFParamsClass()
	rv := objc.Send[DICreateASIFParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithCoder:
func NewDICreateASIFParamsWithCoder(coder objectivec.IObject) DICreateASIFParams {
	instance := getDICreateASIFParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DICreateASIFParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:error:
func NewDICreateASIFParamsWithURLError(url foundation.INSURL) (DICreateASIFParams, error) {
	var errorPtr objc.ID
	instance := getDICreateASIFParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateASIFParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateASIFParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateASIFParams/initWithURL:numBlocks:error:
func NewDICreateASIFParamsWithURLNumBlocksError(url foundation.INSURL, numBlocks uint64) (DICreateASIFParams, error) {
	var errorPtr objc.ID
	instance := getDICreateASIFParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:numBlocks:error:"), url, numBlocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateASIFParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateASIFParamsFromID(rv), nil
}

