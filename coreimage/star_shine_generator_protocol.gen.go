// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a star-shine generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator
type CIStarShineGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/center
	Center() corefoundation.CGPoint

	// The color to use for the outer shell of the circular star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/color
	Color() ICIColor

	// The angle of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossAngle
	CrossAngle() float32

	// The opacity of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossOpacity
	CrossOpacity() float32

	// The size of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossScale
	CrossScale() float32

	// The width of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossWidth
	CrossWidth() float32

	// The length of the cross spikes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/epsilon
	Epsilon() float32

	// The radius of the star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/radius
	Radius() float32

	// The x and y position to use as the center of the star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/center
	SetCenter(value corefoundation.CGPoint)

	// The color to use for the outer shell of the circular star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/color
	SetColor(value ICIColor)

	// The angle of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossAngle
	SetCrossAngle(value float32)

	// The opacity of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossOpacity
	SetCrossOpacity(value float32)

	// The size of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossScale
	SetCrossScale(value float32)

	// The width of the cross pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossWidth
	SetCrossWidth(value float32)

	// The length of the cross spikes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/epsilon
	SetEpsilon(value float32)

	// The radius of the star.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/radius
	SetRadius(value float32)
}

// CIStarShineGeneratorObject wraps an existing Objective-C object that conforms to the CIStarShineGenerator protocol.
type CIStarShineGeneratorObject struct {
	objectivec.Object
}
func (o CIStarShineGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIStarShineGeneratorObjectFromID constructs a [CIStarShineGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIStarShineGeneratorObjectFromID(id objc.ID) CIStarShineGeneratorObject {
	return CIStarShineGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the star.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/center
func (o CIStarShineGeneratorObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The color to use for the outer shell of the circular star.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/color
func (o CIStarShineGeneratorObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The angle of the cross pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossAngle
func (o CIStarShineGeneratorObject) CrossAngle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("crossAngle"))
	return rv
	}
// The opacity of the cross pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossOpacity
func (o CIStarShineGeneratorObject) CrossOpacity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("crossOpacity"))
	return rv
	}
// The size of the cross pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossScale
func (o CIStarShineGeneratorObject) CrossScale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("crossScale"))
	return rv
	}
// The width of the cross pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/crossWidth
func (o CIStarShineGeneratorObject) CrossWidth() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("crossWidth"))
	return rv
	}
// The length of the cross spikes.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/epsilon
func (o CIStarShineGeneratorObject) Epsilon() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
	}
// The radius of the star.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStarShineGenerator/radius
func (o CIStarShineGeneratorObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIStarShineGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIStarShineGeneratorObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIStarShineGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIStarShineGeneratorObject) SetCrossAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCrossAngle:"), value)
}

func (o CIStarShineGeneratorObject) SetCrossOpacity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCrossOpacity:"), value)
}

func (o CIStarShineGeneratorObject) SetCrossScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCrossScale:"), value)
}

func (o CIStarShineGeneratorObject) SetCrossWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCrossWidth:"), value)
}

func (o CIStarShineGeneratorObject) SetEpsilon(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setEpsilon:"), value)
}

func (o CIStarShineGeneratorObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

