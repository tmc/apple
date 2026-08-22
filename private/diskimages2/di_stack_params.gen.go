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

// The class instance for the [DIStackParams] class.
var (
	_DIStackParamsClass     DIStackParamsClass
	_DIStackParamsClassOnce sync.Once
)

func getDIStackParamsClass() DIStackParamsClass {
	_DIStackParamsClassOnce.Do(func() {
		_DIStackParamsClass = DIStackParamsClass{class: objc.GetClass("DIStackParams")}
	})
	return _DIStackParamsClass
}

// GetDIStackParamsClass returns the class object for DIStackParams.
func GetDIStackParamsClass() DIStackParamsClass {
	return getDIStackParamsClass()
}

type DIStackParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIStackParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIStackParamsClass) Alloc() DIStackParams {
	rv := objc.SendIfResponds[DIStackParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIStackParams.AppendWithURLIsCacheError]
//   - [DIStackParams.AppendWithURLIsCacheNumBlocksError]
type DIStackParams struct {
	DIBaseParams
}

// DIStackParamsFromID constructs a [DIStackParams] from an objc.ID.
func DIStackParamsFromID(id objc.ID) DIStackParams {
	return DIStackParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIStackParams implements IDIStackParams.
var _ IDIStackParams = DIStackParams{}

// An interface definition for the [DIStackParams] class.
//
// # Methods
//
//   - [IDIStackParams.AppendWithURLIsCacheError]
//   - [IDIStackParams.AppendWithURLIsCacheNumBlocksError]
type IDIStackParams interface {
	IDIBaseParams

	// Topic: Methods

	AppendWithURLIsCacheError(url foundation.NSURL, cache bool) (bool, error)
	AppendWithURLIsCacheNumBlocksError(url foundation.NSURL, cache bool, blocks uint64) (bool, error)
}

// Init initializes the instance.
func (d DIStackParams) Init() DIStackParams {
	rv := objc.SendIfResponds[DIStackParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIStackParams) Autorelease() DIStackParams {
	rv := objc.SendIfResponds[DIStackParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIStackParams creates a new DIStackParams instance.
func NewDIStackParams() DIStackParams {
	class := getDIStackParamsClass()
	rv := objc.SendIfResponds[DIStackParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIStackParamsWithCoder(coder objectivec.IObject) DIStackParams {
	instance := getDIStackParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIStackParamsFromID(rv)
}

func NewDIStackParamsWithURLError(url foundation.NSURL) (DIStackParams, error) {
	var errorPtr objc.ID
	instance := getDIStackParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStackParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIStackParams{}, objc.ErrInitFailed
	}
	return DIStackParamsFromID(rv), nil
}

func (d DIStackParams) AppendWithURLIsCacheError(url foundation.NSURL, cache bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendWithURL:isCache:error:"), url, cache, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendWithURL:isCache:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIStackParams) AppendWithURLIsCacheNumBlocksError(url foundation.NSURL, cache bool, blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendWithURL:isCache:numBlocks:error:"), url, cache, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendWithURL:isCache:numBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
