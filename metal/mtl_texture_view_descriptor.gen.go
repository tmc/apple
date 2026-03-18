// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTextureViewDescriptor] class.
var (
	_MTLTextureViewDescriptorClass     MTLTextureViewDescriptorClass
	_MTLTextureViewDescriptorClassOnce sync.Once
)

func getMTLTextureViewDescriptorClass() MTLTextureViewDescriptorClass {
	_MTLTextureViewDescriptorClassOnce.Do(func() {
		_MTLTextureViewDescriptorClass = MTLTextureViewDescriptorClass{class: objc.GetClass("MTLTextureViewDescriptor")}
	})
	return _MTLTextureViewDescriptorClass
}

// GetMTLTextureViewDescriptorClass returns the class object for MTLTextureViewDescriptor.
func GetMTLTextureViewDescriptorClass() MTLTextureViewDescriptorClass {
	return getMTLTextureViewDescriptorClass()
}

type MTLTextureViewDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTextureViewDescriptorClass) Alloc() MTLTextureViewDescriptor {
	rv := objc.Send[MTLTextureViewDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [MTLTextureViewDescriptor.PixelFormat]
//   - [MTLTextureViewDescriptor.SetPixelFormat]
//   - [MTLTextureViewDescriptor.Swizzle]
//   - [MTLTextureViewDescriptor.SetSwizzle]
//   - [MTLTextureViewDescriptor.TextureType]
//   - [MTLTextureViewDescriptor.SetTextureType]
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor
type MTLTextureViewDescriptor struct {
	objectivec.Object
}

// MTLTextureViewDescriptorFromID constructs a [MTLTextureViewDescriptor] from an objc.ID.
func MTLTextureViewDescriptorFromID(id objc.ID) MTLTextureViewDescriptor {
	return MTLTextureViewDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLTextureViewDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTextureViewDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLTextureViewDescriptor.PixelFormat]
//   - [IMTLTextureViewDescriptor.SetPixelFormat]
//   - [IMTLTextureViewDescriptor.Swizzle]
//   - [IMTLTextureViewDescriptor.SetSwizzle]
//   - [IMTLTextureViewDescriptor.TextureType]
//   - [IMTLTextureViewDescriptor.SetTextureType]
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor
type IMTLTextureViewDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	PixelFormat() MTLPixelFormat
	SetPixelFormat(value MTLPixelFormat)
	Swizzle() MTLTextureSwizzleChannels
	SetSwizzle(value MTLTextureSwizzleChannels)
	TextureType() MTLTextureType
	SetTextureType(value MTLTextureType)

	LevelRange() foundation.NSRange
	SetLevelRange(value foundation.NSRange)
	SliceRange() foundation.NSRange
	SetSliceRange(value foundation.NSRange)
}

// Init initializes the instance.
func (t MTLTextureViewDescriptor) Init() MTLTextureViewDescriptor {
	rv := objc.Send[MTLTextureViewDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTextureViewDescriptor) Autorelease() MTLTextureViewDescriptor {
	rv := objc.Send[MTLTextureViewDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTextureViewDescriptor creates a new MTLTextureViewDescriptor instance.
func NewMTLTextureViewDescriptor() MTLTextureViewDescriptor {
	class := getMTLTextureViewDescriptorClass()
	rv := objc.Send[MTLTextureViewDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// # Discussion
// 
// A desired pixel format of a texture view.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/pixelFormat
func (t MTLTextureViewDescriptor) PixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](t.ID, objc.Sel("pixelFormat"))
	return MTLPixelFormat(rv)
}
func (t MTLTextureViewDescriptor) SetPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setPixelFormat:"), value)
}

//
// # Discussion
// 
// A desired swizzle format of a texture view.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/swizzle
func (t MTLTextureViewDescriptor) Swizzle() MTLTextureSwizzleChannels {
	rv := objc.Send[MTLTextureSwizzleChannels](t.ID, objc.Sel("swizzle"))
	return MTLTextureSwizzleChannels(rv)
}
func (t MTLTextureViewDescriptor) SetSwizzle(value MTLTextureSwizzleChannels) {
	objc.Send[struct{}](t.ID, objc.Sel("setSwizzle:"), value)
}

//
// # Discussion
// 
// A desired texture view of a texture view.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t MTLTextureViewDescriptor) TextureType() MTLTextureType {
	rv := objc.Send[MTLTextureType](t.ID, objc.Sel("textureType"))
	return MTLTextureType(rv)
}
func (t MTLTextureViewDescriptor) SetTextureType(value MTLTextureType) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextureType:"), value)
}

//
// # Discussion
// 
// A desired range of mip levels of a texture view.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/levelRange-7e7f3
func (t MTLTextureViewDescriptor) LevelRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("levelRange"))
	return foundation.NSRange(rv)
}
func (t MTLTextureViewDescriptor) SetLevelRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setLevelRange:"), value)
}

//
// # Discussion
// 
// A desired range of slices of a texture view.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/sliceRange-3cs9b
func (t MTLTextureViewDescriptor) SliceRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("sliceRange"))
	return foundation.NSRange(rv)
}
func (t MTLTextureViewDescriptor) SetSliceRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setSliceRange:"), value)
}

