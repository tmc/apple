// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a tone curve filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve
type CIToneCurve interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/inputImage
	InputImage() ICIImage

	// A vector containing the position of the first point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point0
	Point0() corefoundation.CGPoint

	// A vector containing the position of the second point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point1
	Point1() corefoundation.CGPoint

	// A vector containing the position of the third point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point2
	Point2() corefoundation.CGPoint

	// A vector containing the position of the fourth point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point3
	Point3() corefoundation.CGPoint

	// A vector containing the position of the fifth point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point4
	Point4() corefoundation.CGPoint

	// If true, then the color effect will be extrapolated if the input image contains RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/extrapolate
	Extrapolate() bool

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/inputImage
	SetInputImage(value ICIImage)

	// A vector containing the position of the first point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point0
	SetPoint0(value corefoundation.CGPoint)

	// A vector containing the position of the second point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point1
	SetPoint1(value corefoundation.CGPoint)

	// A vector containing the position of the third point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point2
	SetPoint2(value corefoundation.CGPoint)

	// A vector containing the position of the fourth point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point3
	SetPoint3(value corefoundation.CGPoint)

	// A vector containing the position of the fifth point of the tone curve.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point4
	SetPoint4(value corefoundation.CGPoint)

	// If true, then the color effect will be extrapolated if the input image contains RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/extrapolate
	SetExtrapolate(value bool)
}

// CIToneCurveObject wraps an existing Objective-C object that conforms to the CIToneCurve protocol.
type CIToneCurveObject struct {
	objectivec.Object
}
func (o CIToneCurveObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIToneCurveObjectFromID constructs a [CIToneCurveObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIToneCurveObjectFromID(id objc.ID) CIToneCurveObject {
	return CIToneCurveObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/inputImage
func (o CIToneCurveObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A vector containing the position of the first point of the tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point0
func (o CIToneCurveObject) Point0() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point0"))
	return rv
	}
// A vector containing the position of the second point of the tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point1
func (o CIToneCurveObject) Point1() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point1"))
	return rv
	}
// A vector containing the position of the third point of the tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point2
func (o CIToneCurveObject) Point2() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point2"))
	return rv
	}
// A vector containing the position of the fourth point of the tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point3
func (o CIToneCurveObject) Point3() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point3"))
	return rv
	}
// A vector containing the position of the fifth point of the tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/point4
func (o CIToneCurveObject) Point4() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point4"))
	return rv
	}
// If true, then the color effect will be extrapolated if the input image
// contains RGB component values outside the range 0.0 to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneCurve/extrapolate
func (o CIToneCurveObject) Extrapolate() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("extrapolate"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIToneCurveObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIToneCurveObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIToneCurveObject) SetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint0:"), value)
}

func (o CIToneCurveObject) SetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint1:"), value)
}

func (o CIToneCurveObject) SetPoint2(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint2:"), value)
}

func (o CIToneCurveObject) SetPoint3(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint3:"), value)
}

func (o CIToneCurveObject) SetPoint4(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint4:"), value)
}

func (o CIToneCurveObject) SetExtrapolate(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtrapolate:"), value)
}

