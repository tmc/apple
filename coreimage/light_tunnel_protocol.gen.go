// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CILightTunnel protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel
type CILightTunnel interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/radius
	Radius() float32

	// Rotation protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/rotation
	Rotation() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/radius
	SetRadius(value float32)

	// rotation protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/rotation
	SetRotation(value float32)
}

// CILightTunnelObject wraps an existing Objective-C object that conforms to the CILightTunnel protocol.
type CILightTunnelObject struct {
	objectivec.Object
}
func (o CILightTunnelObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILightTunnelObjectFromID constructs a [CILightTunnelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILightTunnelObjectFromID(id objc.ID) CILightTunnelObject {
	return CILightTunnelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/center
func (o CILightTunnelObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/inputImage
func (o CILightTunnelObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/radius
func (o CILightTunnelObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CILightTunnel/rotation
func (o CILightTunnelObject) Rotation() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("rotation"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILightTunnelObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILightTunnelObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CILightTunnelObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CILightTunnelObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CILightTunnelObject) SetRotation(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRotation:"), value)
}

