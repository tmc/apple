// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a line overlay filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay
type CILineOverlay interface {
	objectivec.IObject
	CIFilterProtocol

	// The noise level of the image, used with camera data, that’s removed before tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrNoiseLevel
	NRNoiseLevel() float32

	// The amount of sharpening done when removing noise in the image before tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrSharpness
	NRSharpness() float32

	// The amount of antialiasing to use on the edges produced by this filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/contrast
	Contrast() float32

	// The accentuation factor of the Sobel gradient information when tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/edgeIntensity
	EdgeIntensity() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/inputImage
	InputImage() ICIImage

	// A value that determines edge visibility.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/threshold
	Threshold() float32

	// The noise level of the image, used with camera data, that’s removed before tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrNoiseLevel
	SetNRNoiseLevel(value float32)

	// The amount of sharpening done when removing noise in the image before tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrSharpness
	SetNRSharpness(value float32)

	// The amount of antialiasing to use on the edges produced by this filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/contrast
	SetContrast(value float32)

	// The accentuation factor of the Sobel gradient information when tracing the edges of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/edgeIntensity
	SetEdgeIntensity(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/inputImage
	SetInputImage(value ICIImage)

	// A value that determines edge visibility.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/threshold
	SetThreshold(value float32)
}

// CILineOverlayObject wraps an existing Objective-C object that conforms to the CILineOverlay protocol.
type CILineOverlayObject struct {
	objectivec.Object
}
func (o CILineOverlayObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILineOverlayObjectFromID constructs a [CILineOverlayObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILineOverlayObjectFromID(id objc.ID) CILineOverlayObject {
	return CILineOverlayObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The noise level of the image, used with camera data, that’s removed
// before tracing the edges of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrNoiseLevel
func (o CILineOverlayObject) NRNoiseLevel() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("NRNoiseLevel"))
	return rv
	}
// The amount of sharpening done when removing noise in the image before
// tracing the edges of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/nrSharpness
func (o CILineOverlayObject) NRSharpness() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("NRSharpness"))
	return rv
	}
// The amount of antialiasing to use on the edges produced by this filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/contrast
func (o CILineOverlayObject) Contrast() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("contrast"))
	return rv
	}
// The accentuation factor of the Sobel gradient information when tracing the
// edges of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/edgeIntensity
func (o CILineOverlayObject) EdgeIntensity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("edgeIntensity"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/inputImage
func (o CILineOverlayObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A value that determines edge visibility.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineOverlay/threshold
func (o CILineOverlayObject) Threshold() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("threshold"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILineOverlayObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILineOverlayObject) SetNRNoiseLevel(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setNRNoiseLevel:"), value)
}

func (o CILineOverlayObject) SetNRSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setNRSharpness:"), value)
}

func (o CILineOverlayObject) SetContrast(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setContrast:"), value)
}

func (o CILineOverlayObject) SetEdgeIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setEdgeIntensity:"), value)
}

func (o CILineOverlayObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CILineOverlayObject) SetThreshold(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreshold:"), value)
}

