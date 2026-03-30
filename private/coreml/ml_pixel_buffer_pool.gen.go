// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPixelBufferPool] class.
var (
	_MLPixelBufferPoolClass     MLPixelBufferPoolClass
	_MLPixelBufferPoolClassOnce sync.Once
)

func getMLPixelBufferPoolClass() MLPixelBufferPoolClass {
	_MLPixelBufferPoolClassOnce.Do(func() {
		_MLPixelBufferPoolClass = MLPixelBufferPoolClass{class: objc.GetClass("MLPixelBufferPool")}
	})
	return _MLPixelBufferPoolClass
}

// GetMLPixelBufferPoolClass returns the class object for MLPixelBufferPool.
func GetMLPixelBufferPoolClass() MLPixelBufferPoolClass {
	return getMLPixelBufferPoolClass()
}

type MLPixelBufferPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPixelBufferPoolClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPixelBufferPoolClass) Alloc() MLPixelBufferPool {
	rv := objc.Send[MLPixelBufferPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPixelBufferPool._pixelBufferPoolWithSizePixelFormatTypeError]
//   - [MLPixelBufferPool.CreatePixelBufferWithSizePixelFormatTypeError]
//   - [MLPixelBufferPool.PixelBufferPoolCache]
//
// See: https://developer.apple.com/documentation/CoreML/MLPixelBufferPool
type MLPixelBufferPool struct {
	objectivec.Object
}

// MLPixelBufferPoolFromID constructs a [MLPixelBufferPool] from an objc.ID.
func MLPixelBufferPoolFromID(id objc.ID) MLPixelBufferPool {
	return MLPixelBufferPool{objectivec.Object{ID: id}}
}

// Ensure MLPixelBufferPool implements IMLPixelBufferPool.
var _ IMLPixelBufferPool = MLPixelBufferPool{}

// An interface definition for the [MLPixelBufferPool] class.
//
// # Methods
//
//   - [IMLPixelBufferPool._pixelBufferPoolWithSizePixelFormatTypeError]
//   - [IMLPixelBufferPool.CreatePixelBufferWithSizePixelFormatTypeError]
//   - [IMLPixelBufferPool.PixelBufferPoolCache]
//
// See: https://developer.apple.com/documentation/CoreML/MLPixelBufferPool
type IMLPixelBufferPool interface {
	objectivec.IObject

	// Topic: Methods

	_pixelBufferPoolWithSizePixelFormatTypeError(size corefoundation.CGSize, type_ uint32) (corevideo.CVImageBufferRef, error)
	CreatePixelBufferWithSizePixelFormatTypeError(size corefoundation.CGSize, type_ uint32) (corevideo.CVImageBufferRef, error)
	PixelBufferPoolCache() foundation.INSDictionary
}

// Init initializes the instance.
func (p MLPixelBufferPool) Init() MLPixelBufferPool {
	rv := objc.Send[MLPixelBufferPool](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPixelBufferPool) Autorelease() MLPixelBufferPool {
	rv := objc.Send[MLPixelBufferPool](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPixelBufferPool creates a new MLPixelBufferPool instance.
func NewMLPixelBufferPool() MLPixelBufferPool {
	class := getMLPixelBufferPoolClass()
	rv := objc.Send[MLPixelBufferPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPixelBufferPool/_pixelBufferPoolWithSize:pixelFormatType:error:
func (p MLPixelBufferPool) _pixelBufferPoolWithSizePixelFormatTypeError(size corefoundation.CGSize, type_ uint32) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](p.ID, objc.Sel("_pixelBufferPoolWithSize:pixelFormatType:error:"), size, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// PixelBufferPoolWithSizePixelFormatTypeError is an exported wrapper for the private method _pixelBufferPoolWithSizePixelFormatTypeError.
func (p MLPixelBufferPool) PixelBufferPoolWithSizePixelFormatTypeError(size corefoundation.CGSize, type_ uint32) (corevideo.CVImageBufferRef, error) {
	return p._pixelBufferPoolWithSizePixelFormatTypeError(size, type_)
}

// See: https://developer.apple.com/documentation/CoreML/MLPixelBufferPool/createPixelBufferWithSize:pixelFormatType:error:
func (p MLPixelBufferPool) CreatePixelBufferWithSizePixelFormatTypeError(size corefoundation.CGSize, type_ uint32) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](p.ID, objc.Sel("createPixelBufferWithSize:pixelFormatType:error:"), size, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPixelBufferPool/pixelBufferPoolCache
func (p MLPixelBufferPool) PixelBufferPoolCache() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pixelBufferPoolCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
