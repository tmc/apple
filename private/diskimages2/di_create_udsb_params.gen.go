// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DICreateUDSBParams] class.
var (
	_DICreateUDSBParamsClass     DICreateUDSBParamsClass
	_DICreateUDSBParamsClassOnce sync.Once
)

func getDICreateUDSBParamsClass() DICreateUDSBParamsClass {
	_DICreateUDSBParamsClassOnce.Do(func() {
		_DICreateUDSBParamsClass = DICreateUDSBParamsClass{class: objc.GetClass("DICreateUDSBParams")}
	})
	return _DICreateUDSBParamsClass
}

// GetDICreateUDSBParamsClass returns the class object for DICreateUDSBParams.
func GetDICreateUDSBParamsClass() DICreateUDSBParamsClass {
	return getDICreateUDSBParamsClass()
}

type DICreateUDSBParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DICreateUDSBParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICreateUDSBParamsClass) Alloc() DICreateUDSBParams {
	rv := objc.Send[DICreateUDSBParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DICreateUDSBParams.SparseBundleBandSize]
//   - [DICreateUDSBParams.SetSparseBundleBandSize]
// See: https://developer.apple.com/documentation/DiskImages2/DICreateUDSBParams
type DICreateUDSBParams struct {
	DICreateParams
}

// DICreateUDSBParamsFromID constructs a [DICreateUDSBParams] from an objc.ID.
func DICreateUDSBParamsFromID(id objc.ID) DICreateUDSBParams {
	return DICreateUDSBParams{DICreateParams: DICreateParamsFromID(id)}
}
// Ensure DICreateUDSBParams implements IDICreateUDSBParams.
var _ IDICreateUDSBParams = DICreateUDSBParams{}

// An interface definition for the [DICreateUDSBParams] class.
//
// # Methods
//
//   - [IDICreateUDSBParams.SparseBundleBandSize]
//   - [IDICreateUDSBParams.SetSparseBundleBandSize]
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateUDSBParams
type IDICreateUDSBParams interface {
	IDICreateParams

	// Topic: Methods

	SparseBundleBandSize() uint64
	SetSparseBundleBandSize(value uint64)
}

// Init initializes the instance.
func (d DICreateUDSBParams) Init() DICreateUDSBParams {
	rv := objc.Send[DICreateUDSBParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICreateUDSBParams) Autorelease() DICreateUDSBParams {
	rv := objc.Send[DICreateUDSBParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICreateUDSBParams creates a new DICreateUDSBParams instance.
func NewDICreateUDSBParams() DICreateUDSBParams {
	class := getDICreateUDSBParamsClass()
	rv := objc.Send[DICreateUDSBParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithCoder:
func NewDICreateUDSBParamsWithCoder(coder objectivec.IObject) DICreateUDSBParams {
	instance := getDICreateUDSBParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DICreateUDSBParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:error:
func NewDICreateUDSBParamsWithURLError(url foundation.INSURL) (DICreateUDSBParams, error) {
	var errorPtr objc.ID
	instance := getDICreateUDSBParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateUDSBParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateUDSBParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateUDSBParams/initWithURL:numBlocks:error:
func NewDICreateUDSBParamsWithURLNumBlocksError(url foundation.INSURL, blocks uint64) (DICreateUDSBParams, error) {
	var errorPtr objc.ID
	instance := getDICreateUDSBParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:numBlocks:error:"), url, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateUDSBParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateUDSBParamsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DICreateUDSBParams/sparseBundleBandSize
func (d DICreateUDSBParams) SparseBundleBandSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("sparseBundleBandSize"))
	return rv
}
func (d DICreateUDSBParams) SetSparseBundleBandSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setSparseBundleBandSize:"), value)
}

