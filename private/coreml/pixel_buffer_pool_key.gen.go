// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PixelBufferPoolKey] class.
var (
	_PixelBufferPoolKeyClass     PixelBufferPoolKeyClass
	_PixelBufferPoolKeyClassOnce sync.Once
)

func getPixelBufferPoolKeyClass() PixelBufferPoolKeyClass {
	_PixelBufferPoolKeyClassOnce.Do(func() {
		_PixelBufferPoolKeyClass = PixelBufferPoolKeyClass{class: objc.GetClass("PixelBufferPoolKey")}
	})
	return _PixelBufferPoolKeyClass
}

// GetPixelBufferPoolKeyClass returns the class object for PixelBufferPoolKey.
func GetPixelBufferPoolKeyClass() PixelBufferPoolKeyClass {
	return getPixelBufferPoolKeyClass()
}

type PixelBufferPoolKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PixelBufferPoolKeyClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PixelBufferPoolKeyClass) Alloc() PixelBufferPoolKey {
	rv := objc.Send[PixelBufferPoolKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [PixelBufferPoolKey.FrameSize]
//   - [PixelBufferPoolKey.PixelFormatType]
//   - [PixelBufferPoolKey.InitWithSizePixelFormatType]
// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey
type PixelBufferPoolKey struct {
	objectivec.Object
}

// PixelBufferPoolKeyFromID constructs a [PixelBufferPoolKey] from an objc.ID.
func PixelBufferPoolKeyFromID(id objc.ID) PixelBufferPoolKey {
	return PixelBufferPoolKey{objectivec.Object{ID: id}}
}
// Ensure PixelBufferPoolKey implements IPixelBufferPoolKey.
var _ IPixelBufferPoolKey = PixelBufferPoolKey{}

// An interface definition for the [PixelBufferPoolKey] class.
//
// # Methods
//
//   - [IPixelBufferPoolKey.FrameSize]
//   - [IPixelBufferPoolKey.PixelFormatType]
//   - [IPixelBufferPoolKey.InitWithSizePixelFormatType]
//
// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey
type IPixelBufferPoolKey interface {
	objectivec.IObject

	// Topic: Methods

	FrameSize() corefoundation.CGSize
	PixelFormatType() uint32
	InitWithSizePixelFormatType(size corefoundation.CGSize, type_ uint32) PixelBufferPoolKey
}

// Init initializes the instance.
func (p PixelBufferPoolKey) Init() PixelBufferPoolKey {
	rv := objc.Send[PixelBufferPoolKey](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PixelBufferPoolKey) Autorelease() PixelBufferPoolKey {
	rv := objc.Send[PixelBufferPoolKey](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPixelBufferPoolKey creates a new PixelBufferPoolKey instance.
func NewPixelBufferPoolKey() PixelBufferPoolKey {
	class := getPixelBufferPoolKeyClass()
	rv := objc.Send[PixelBufferPoolKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey/initWithSize:pixelFormatType:
func NewPixelBufferPoolKeyWithSizePixelFormatType(size corefoundation.CGSize, type_ uint32) PixelBufferPoolKey {
	instance := getPixelBufferPoolKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:pixelFormatType:"), size, type_)
	return PixelBufferPoolKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey/initWithSize:pixelFormatType:
func (p PixelBufferPoolKey) InitWithSizePixelFormatType(size corefoundation.CGSize, type_ uint32) PixelBufferPoolKey {
	rv := objc.Send[PixelBufferPoolKey](p.ID, objc.Sel("initWithSize:pixelFormatType:"), size, type_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey/frameSize
func (p PixelBufferPoolKey) FrameSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("frameSize"))
	return corefoundation.CGSize(rv)
}
// See: https://developer.apple.com/documentation/CoreML/PixelBufferPoolKey/pixelFormatType
func (p PixelBufferPoolKey) PixelFormatType() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("pixelFormatType"))
	return rv
}

