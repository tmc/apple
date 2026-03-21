// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a copy machine transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition
type CICopyMachineTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/angle
	Angle() float32

	// The color of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/color
	Color() ICIColor

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/extent
	Extent() corefoundation.CGRect

	// The opacity of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/opacity
	Opacity() float32

	// The width of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/width
	Width() float32

	// The angle of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/angle
	SetAngle(value float32)

	// The color of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/color
	SetColor(value ICIColor)

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/extent
	SetExtent(value corefoundation.CGRect)

	// The opacity of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/opacity
	SetOpacity(value float32)

	// The width of the copier light.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/width
	SetWidth(value float32)
}

// CICopyMachineTransitionObject wraps an existing Objective-C object that conforms to the CICopyMachineTransition protocol.
type CICopyMachineTransitionObject struct {
	objectivec.Object
}
func (o CICopyMachineTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICopyMachineTransitionObjectFromID constructs a [CICopyMachineTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICopyMachineTransitionObjectFromID(id objc.ID) CICopyMachineTransitionObject {
	return CICopyMachineTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the copier light.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/angle
func (o CICopyMachineTransitionObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The color of the copier light.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/color
func (o CICopyMachineTransitionObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/extent
func (o CICopyMachineTransitionObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The opacity of the copier light.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/opacity
func (o CICopyMachineTransitionObject) Opacity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("opacity"))
	return rv
	}
// The width of the copier light.
//
// See: https://developer.apple.com/documentation/CoreImage/CICopyMachineTransition/width
func (o CICopyMachineTransitionObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICopyMachineTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CICopyMachineTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CICopyMachineTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CICopyMachineTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CICopyMachineTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CICopyMachineTransitionObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CICopyMachineTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CICopyMachineTransitionObject) SetOpacity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpacity:"), value)
}

func (o CICopyMachineTransitionObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

func (o CICopyMachineTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CICopyMachineTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CICopyMachineTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}

