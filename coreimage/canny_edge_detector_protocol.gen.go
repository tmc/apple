// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CICannyEdgeDetector protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector
type CICannyEdgeDetector interface {
	objectivec.IObject
	CIFilterProtocol

	// GaussianSigma protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/gaussianSigma
	GaussianSigma() float32

	// HysteresisPasses protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/hysteresisPasses
	HysteresisPasses() int

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/inputImage
	InputImage() ICIImage

	// Perceptual protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/perceptual
	Perceptual() bool

	// ThresholdHigh protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdHigh
	ThresholdHigh() float32

	// ThresholdLow protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdLow
	ThresholdLow() float32

	// gaussianSigma protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/gaussianSigma
	SetGaussianSigma(value float32)

	// hysteresisPasses protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/hysteresisPasses
	SetHysteresisPasses(value int)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/inputImage
	SetInputImage(value ICIImage)

	// perceptual protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/perceptual
	SetPerceptual(value bool)

	// thresholdHigh protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdHigh
	SetThresholdHigh(value float32)

	// thresholdLow protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdLow
	SetThresholdLow(value float32)
}

// CICannyEdgeDetectorObject wraps an existing Objective-C object that conforms to the CICannyEdgeDetector protocol.
type CICannyEdgeDetectorObject struct {
	objectivec.Object
}

func (o CICannyEdgeDetectorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICannyEdgeDetectorObjectFromID constructs a [CICannyEdgeDetectorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICannyEdgeDetectorObjectFromID(id objc.ID) CICannyEdgeDetectorObject {
	return CICannyEdgeDetectorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/gaussianSigma
func (o CICannyEdgeDetectorObject) GaussianSigma() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gaussianSigma"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/hysteresisPasses
func (o CICannyEdgeDetectorObject) HysteresisPasses() int {
	rv := objc.Send[int](o.ID, objc.Sel("hysteresisPasses"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/inputImage
func (o CICannyEdgeDetectorObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/perceptual
func (o CICannyEdgeDetectorObject) Perceptual() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("perceptual"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdHigh
func (o CICannyEdgeDetectorObject) ThresholdHigh() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("thresholdHigh"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdLow
func (o CICannyEdgeDetectorObject) ThresholdLow() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("thresholdLow"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICannyEdgeDetectorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/gaussianSigma
func (o CICannyEdgeDetectorObject) SetGaussianSigma(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGaussianSigma:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/hysteresisPasses
func (o CICannyEdgeDetectorObject) SetHysteresisPasses(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setHysteresisPasses:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/inputImage
func (o CICannyEdgeDetectorObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/perceptual
func (o CICannyEdgeDetectorObject) SetPerceptual(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setPerceptual:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdHigh
func (o CICannyEdgeDetectorObject) SetThresholdHigh(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setThresholdHigh:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CICannyEdgeDetector/thresholdLow
func (o CICannyEdgeDetectorObject) SetThresholdLow(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setThresholdLow:"), value)
}
