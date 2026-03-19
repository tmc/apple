// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MTLTextureBinding protocol.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding
type MTLTextureBinding interface {
	objectivec.IObject
	MTLBinding

	// ArrayLength protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/arrayLength
	ArrayLength() uint

	// DepthTexture protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/isDepthTexture
	IsDepthTexture() bool

	// TextureDataType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/textureDataType
	TextureDataType() MTLDataType

	// TextureType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/textureType
	TextureType() MTLTextureType
}

// MTLTextureBindingObject wraps an existing Objective-C object that conforms to the MTLTextureBinding protocol.
type MTLTextureBindingObject struct {
	objectivec.Object
}
func (o MTLTextureBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTextureBindingObjectFromID constructs a [MTLTextureBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTextureBindingObjectFromID(id objc.ID) MTLTextureBindingObject {
	return MTLTextureBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/arrayLength
func (o MTLTextureBindingObject) ArrayLength() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("arrayLength"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/isDepthTexture
func (o MTLTextureBindingObject) IsDepthTexture() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isDepthTexture"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/textureDataType
func (o MTLTextureBindingObject) TextureDataType() MTLDataType {
	
	rv := objc.Send[MTLDataType](o.ID, objc.Sel("textureDataType"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTextureBinding/textureType
func (o MTLTextureBindingObject) TextureType() MTLTextureType {
	
	rv := objc.Send[MTLTextureType](o.ID, objc.Sel("textureType"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
func (o MTLTextureBindingObject) Access() MTLBindingAccess {
	
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
func (o MTLTextureBindingObject) Index() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
func (o MTLTextureBindingObject) IsArgument() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
func (o MTLTextureBindingObject) IsUsed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
func (o MTLTextureBindingObject) Name() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
func (o MTLTextureBindingObject) Type() MTLBindingType {
	
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
	}

