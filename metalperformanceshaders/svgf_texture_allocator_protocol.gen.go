// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSSVGFTextureAllocator protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator
type MPSSVGFTextureAllocator interface {
	objectivec.IObject

	// ReturnTexture protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator/return(_:)
	ReturnTexture(texture metal.MTLTexture)

	// TextureWithPixelFormatWidthHeight protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator/texture(with:width:height:)
	TextureWithPixelFormatWidthHeight(pixelFormat metal.MTLPixelFormat, width uint, height uint) metal.MTLTexture
}

// MPSSVGFTextureAllocatorObject wraps an existing Objective-C object that conforms to the MPSSVGFTextureAllocator protocol.
type MPSSVGFTextureAllocatorObject struct {
	objectivec.Object
}

func (o MPSSVGFTextureAllocatorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSSVGFTextureAllocatorObjectFromID constructs a [MPSSVGFTextureAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSSVGFTextureAllocatorObjectFromID(id objc.ID) MPSSVGFTextureAllocatorObject {
	return MPSSVGFTextureAllocatorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator/return(_:)
func (o MPSSVGFTextureAllocatorObject) ReturnTexture(texture metal.MTLTexture) {
	objc.Send[struct{}](o.ID, objc.Sel("returnTexture:"), texture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator/texture(with:width:height:)
func (o MPSSVGFTextureAllocatorObject) TextureWithPixelFormatWidthHeight(pixelFormat metal.MTLPixelFormat, width uint, height uint) metal.MTLTexture {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textureWithPixelFormat:width:height:"), pixelFormat, width, height)
	return metal.MTLTextureObjectFromID(rv)
}
