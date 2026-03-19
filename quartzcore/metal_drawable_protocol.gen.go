// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// A Metal drawable associated with a Core Animation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDrawable
type CAMetalDrawable interface {
	objectivec.IObject
	metal.MTLDrawable

	// A Metal texture object that contains the drawable’s contents.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDrawable/texture
	Texture() metal.MTLTexture

	// The layer that owns this drawable object.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDrawable/layer
	Layer() ICAMetalLayer
}

// CAMetalDrawableObject wraps an existing Objective-C object that conforms to the CAMetalDrawable protocol.
type CAMetalDrawableObject struct {
	metal.MTLDrawableObject
}
func (o CAMetalDrawableObject) BaseObject() objectivec.Object {
	return o.MTLDrawableObject.BaseObject()
}

// CAMetalDrawableObjectFromID constructs a [CAMetalDrawableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CAMetalDrawableObjectFromID(id objc.ID) CAMetalDrawableObject {
	return CAMetalDrawableObject{
		MTLDrawableObject: metal.MTLDrawableObjectFromID(id),
	}
}

// A Metal texture object that contains the drawable’s contents.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDrawable/texture
func (o CAMetalDrawableObject) Texture() metal.MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("texture"))
	return metal.MTLTextureObjectFromID(rv)
	}
// The layer that owns this drawable object.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDrawable/layer
func (o CAMetalDrawableObject) Layer() ICAMetalLayer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("layer"))
	return CAMetalLayerFromID(rv)
	}

