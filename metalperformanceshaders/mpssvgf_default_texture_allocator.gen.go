// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSSVGFDefaultTextureAllocator] class.
var (
	_MPSSVGFDefaultTextureAllocatorClass     MPSSVGFDefaultTextureAllocatorClass
	_MPSSVGFDefaultTextureAllocatorClassOnce sync.Once
)

func getMPSSVGFDefaultTextureAllocatorClass() MPSSVGFDefaultTextureAllocatorClass {
	_MPSSVGFDefaultTextureAllocatorClassOnce.Do(func() {
		_MPSSVGFDefaultTextureAllocatorClass = MPSSVGFDefaultTextureAllocatorClass{class: objc.GetClass("MPSSVGFDefaultTextureAllocator")}
	})
	return _MPSSVGFDefaultTextureAllocatorClass
}

// GetMPSSVGFDefaultTextureAllocatorClass returns the class object for MPSSVGFDefaultTextureAllocator.
func GetMPSSVGFDefaultTextureAllocatorClass() MPSSVGFDefaultTextureAllocatorClass {
	return getMPSSVGFDefaultTextureAllocatorClass()
}

type MPSSVGFDefaultTextureAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSSVGFDefaultTextureAllocatorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSSVGFDefaultTextureAllocatorClass) Alloc() MPSSVGFDefaultTextureAllocator {
	rv := objc.Send[MPSSVGFDefaultTextureAllocator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSSVGFDefaultTextureAllocator.InitWithDevice]
//
// # Instance Properties
//
//   - [MPSSVGFDefaultTextureAllocator.AllocatedTextureCount]
//   - [MPSSVGFDefaultTextureAllocator.Device]
//
// # Instance Methods
//
//   - [MPSSVGFDefaultTextureAllocator.Reset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator
type MPSSVGFDefaultTextureAllocator struct {
	objectivec.Object
}

// MPSSVGFDefaultTextureAllocatorFromID constructs a [MPSSVGFDefaultTextureAllocator] from an objc.ID.
func MPSSVGFDefaultTextureAllocatorFromID(id objc.ID) MPSSVGFDefaultTextureAllocator {
	return MPSSVGFDefaultTextureAllocator{objectivec.Object{ID: id}}
}

// NOTE: MPSSVGFDefaultTextureAllocator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSSVGFDefaultTextureAllocator] class.
//
// # Initializers
//
//   - [IMPSSVGFDefaultTextureAllocator.InitWithDevice]
//
// # Instance Properties
//
//   - [IMPSSVGFDefaultTextureAllocator.AllocatedTextureCount]
//   - [IMPSSVGFDefaultTextureAllocator.Device]
//
// # Instance Methods
//
//   - [IMPSSVGFDefaultTextureAllocator.Reset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator
type IMPSSVGFDefaultTextureAllocator interface {
	objectivec.IObject
	MPSSVGFTextureAllocator

	// Topic: Initializers

	InitWithDevice(device metal.MTLDevice) MPSSVGFDefaultTextureAllocator

	// Topic: Instance Properties

	AllocatedTextureCount() uint
	Device() metal.MTLDevice

	// Topic: Instance Methods

	Reset()
}

// Init initializes the instance.
func (s MPSSVGFDefaultTextureAllocator) Init() MPSSVGFDefaultTextureAllocator {
	rv := objc.Send[MPSSVGFDefaultTextureAllocator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSSVGFDefaultTextureAllocator) Autorelease() MPSSVGFDefaultTextureAllocator {
	rv := objc.Send[MPSSVGFDefaultTextureAllocator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSSVGFDefaultTextureAllocator creates a new MPSSVGFDefaultTextureAllocator instance.
func NewMPSSVGFDefaultTextureAllocator() MPSSVGFDefaultTextureAllocator {
	class := getMPSSVGFDefaultTextureAllocatorClass()
	rv := objc.Send[MPSSVGFDefaultTextureAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/init(device:)
func NewSVGFDefaultTextureAllocatorWithDevice(device metal.MTLDevice) MPSSVGFDefaultTextureAllocator {
	instance := getMPSSVGFDefaultTextureAllocatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSSVGFDefaultTextureAllocatorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/init(device:)
func (s MPSSVGFDefaultTextureAllocator) InitWithDevice(device metal.MTLDevice) MPSSVGFDefaultTextureAllocator {
	rv := objc.Send[MPSSVGFDefaultTextureAllocator](s.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/reset()
func (s MPSSVGFDefaultTextureAllocator) Reset() {
	objc.Send[objc.ID](s.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/return(_:)
func (s MPSSVGFDefaultTextureAllocator) ReturnTexture(texture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("returnTexture:"), texture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/texture(with:width:height:)
func (s MPSSVGFDefaultTextureAllocator) TextureWithPixelFormatWidthHeight(pixelFormat metal.MTLPixelFormat, width uint, height uint) metal.MTLTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textureWithPixelFormat:width:height:"), pixelFormat, width, height)
	return metal.MTLTextureObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/allocatedTextureCount
func (s MPSSVGFDefaultTextureAllocator) AllocatedTextureCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("allocatedTextureCount"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/device
func (s MPSSVGFDefaultTextureAllocator) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// Protocol methods for MPSSVGFTextureAllocator
