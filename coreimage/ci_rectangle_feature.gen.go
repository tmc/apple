// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIRectangleFeature] class.
var (
	_CIRectangleFeatureClass     CIRectangleFeatureClass
	_CIRectangleFeatureClassOnce sync.Once
)

func getCIRectangleFeatureClass() CIRectangleFeatureClass {
	_CIRectangleFeatureClassOnce.Do(func() {
		_CIRectangleFeatureClass = CIRectangleFeatureClass{class: objc.GetClass("CIRectangleFeature")}
	})
	return _CIRectangleFeatureClass
}

// GetCIRectangleFeatureClass returns the class object for CIRectangleFeature.
func GetCIRectangleFeatureClass() CIRectangleFeatureClass {
	return getCIRectangleFeatureClass()
}

type CIRectangleFeatureClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIRectangleFeatureClass) Alloc() CIRectangleFeature {
	rv := objc.Send[CIRectangleFeature](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about a rectangular region detected in a still or video image.
//
// # Overview
// 
// A detected rectangle feature is not necessarily rectangular in the plane of
// the image; rather, the feature identifies a shape that may be rectangular
// in space (for example a book on a desk) but which appears as a four-sided
// polygon in the image. The properties of a [CIRectangleFeature] object
// identify its four corners in image coordinates.
// 
// You can use rectangle feature detection together with the
// [CIPerspectiveCorrection] filter to transform the feature to a normal
// orientation.
// 
// To detect rectangles in an image or video, choose [CIRectangleFeature.CIDetectorTypeRectangle]
// when initializing a [CIDetector] object, and use the
// [CIDetectorAspectRatio] and [CIDetectorFocalLength] options to specify the
// approximate shape of rectangular features to search for. The detector
// returns at most one rectangle feature, the most prominent found in the
// image.
//
// [CIRectangleFeature.CIDetectorTypeRectangle]: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeRectangle
//
// # Identifying the Corners of a Detected Rectangle
//
//   - [CIRectangleFeature.BottomLeft]: The lower-left corner of the detected rectangle, in image coordinates.
//   - [CIRectangleFeature.BottomRight]: The lower-right corner of the detected rectangle, in image coordinates.
//   - [CIRectangleFeature.TopLeft]: The upper-left corner of the detected rectangle, in image coordinates.
//   - [CIRectangleFeature.TopRight]: The upper-right corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature
type CIRectangleFeature struct {
	CIFeature
}

// CIRectangleFeatureFromID constructs a [CIRectangleFeature] from an objc.ID.
//
// Information about a rectangular region detected in a still or video image.
func CIRectangleFeatureFromID(id objc.ID) CIRectangleFeature {
	return CIRectangleFeature{CIFeature: CIFeatureFromID(id)}
}
// NOTE: CIRectangleFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIRectangleFeature] class.
//
// # Identifying the Corners of a Detected Rectangle
//
//   - [ICIRectangleFeature.BottomLeft]: The lower-left corner of the detected rectangle, in image coordinates.
//   - [ICIRectangleFeature.BottomRight]: The lower-right corner of the detected rectangle, in image coordinates.
//   - [ICIRectangleFeature.TopLeft]: The upper-left corner of the detected rectangle, in image coordinates.
//   - [ICIRectangleFeature.TopRight]: The upper-right corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature
type ICIRectangleFeature interface {
	ICIFeature

	// Topic: Identifying the Corners of a Detected Rectangle

	// The lower-left corner of the detected rectangle, in image coordinates.
	BottomLeft() corefoundation.CGPoint
	// The lower-right corner of the detected rectangle, in image coordinates.
	BottomRight() corefoundation.CGPoint
	// The upper-left corner of the detected rectangle, in image coordinates.
	TopLeft() corefoundation.CGPoint
	// The upper-right corner of the detected rectangle, in image coordinates.
	TopRight() corefoundation.CGPoint

	// A detector that searches for rectangular areas in a still image or video, returning 
	CIDetectorTypeRectangle() string
}

// Init initializes the instance.
func (r CIRectangleFeature) Init() CIRectangleFeature {
	rv := objc.Send[CIRectangleFeature](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CIRectangleFeature) Autorelease() CIRectangleFeature {
	rv := objc.Send[CIRectangleFeature](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIRectangleFeature creates a new CIRectangleFeature instance.
func NewCIRectangleFeature() CIRectangleFeature {
	class := getCIRectangleFeatureClass()
	rv := objc.Send[CIRectangleFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The lower-left corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomLeft-swift.property
func (r CIRectangleFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}
// The lower-right corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomRight-swift.property
func (r CIRectangleFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}
// The upper-left corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topLeft-swift.property
func (r CIRectangleFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}
// The upper-right corner of the detected rectangle, in image coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topRight-swift.property
func (r CIRectangleFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}
// A detector that searches for rectangular areas in a still image or video,
// returning
//
// See: https://developer.apple.com/documentation/coreimage/cidetectortyperectangle
func (r CIRectangleFeature) CIDetectorTypeRectangle() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("CIDetectorTypeRectangle"))
	return foundation.NSStringFromID(rv).String()
}

