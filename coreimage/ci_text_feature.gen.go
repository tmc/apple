// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CITextFeature] class.
var (
	_CITextFeatureClass     CITextFeatureClass
	_CITextFeatureClassOnce sync.Once
)

func getCITextFeatureClass() CITextFeatureClass {
	_CITextFeatureClassOnce.Do(func() {
		_CITextFeatureClass = CITextFeatureClass{class: objc.GetClass("CITextFeature")}
	})
	return _CITextFeatureClass
}

// GetCITextFeatureClass returns the class object for CITextFeature.
func GetCITextFeatureClass() CITextFeatureClass {
	return getCITextFeatureClass()
}

type CITextFeatureClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CITextFeatureClass) Alloc() CITextFeature {
	rv := objc.Send[CITextFeature](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about a text that was detected in a still or video image.
//
// # Overview
// 
// A detected text feature is not necessarily rectangular in the plane of the
// image; rather, the feature identifies a shape that may be rectangular in
// space (for example a text on a sign) but which appears as a four-sided
// polygon in the image. The properties of a [CITextFeature] object identify
// its four corners in image coordinates.
// 
// To detect text in an image or video, choose the [CITextFeature.CIDetectorTypeText] type
// when initializing a [CIDetector] object, and use the
// [CIDetectorImageOrientation] option to specify the desired orientation for
// finding upright text.
//
// [CITextFeature.CIDetectorTypeText]: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeText
//
// # Locating Features Within a Detected Region
//
//   - [CITextFeature.SubFeatures]: An array containing additional features detected within the feature.
//
// # Identifying the Corners of a Detected Text Region
//
//   - [CITextFeature.BottomLeft]: The image coordinate of the lower-left corner of the detected text.
//   - [CITextFeature.BottomRight]: The image coordinate of the lower-right corner of the detected text.
//   - [CITextFeature.TopLeft]: The image coordinate of the upper-left corner of the detected text.
//   - [CITextFeature.TopRight]: The image coordinate of the upper-right corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature
type CITextFeature struct {
	CIFeature
}

// CITextFeatureFromID constructs a [CITextFeature] from an objc.ID.
//
// Information about a text that was detected in a still or video image.
func CITextFeatureFromID(id objc.ID) CITextFeature {
	return CITextFeature{CIFeature: CIFeatureFromID(id)}
}
// NOTE: CITextFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CITextFeature] class.
//
// # Locating Features Within a Detected Region
//
//   - [ICITextFeature.SubFeatures]: An array containing additional features detected within the feature.
//
// # Identifying the Corners of a Detected Text Region
//
//   - [ICITextFeature.BottomLeft]: The image coordinate of the lower-left corner of the detected text.
//   - [ICITextFeature.BottomRight]: The image coordinate of the lower-right corner of the detected text.
//   - [ICITextFeature.TopLeft]: The image coordinate of the upper-left corner of the detected text.
//   - [ICITextFeature.TopRight]: The image coordinate of the upper-right corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature
type ICITextFeature interface {
	ICIFeature

	// Topic: Locating Features Within a Detected Region

	// An array containing additional features detected within the feature.
	SubFeatures() foundation.INSArray

	// Topic: Identifying the Corners of a Detected Text Region

	// The image coordinate of the lower-left corner of the detected text.
	BottomLeft() corefoundation.CGPoint
	// The image coordinate of the lower-right corner of the detected text.
	BottomRight() corefoundation.CGPoint
	// The image coordinate of the upper-left corner of the detected text.
	TopLeft() corefoundation.CGPoint
	// The image coordinate of the upper-right corner of the detected text.
	TopRight() corefoundation.CGPoint

	// A detector that searches for text in a still image or video, returning 
	CIDetectorTypeText() string
}

// Init initializes the instance.
func (t CITextFeature) Init() CITextFeature {
	rv := objc.Send[CITextFeature](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CITextFeature) Autorelease() CITextFeature {
	rv := objc.Send[CITextFeature](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCITextFeature creates a new CITextFeature instance.
func NewCITextFeature() CITextFeature {
	class := getCITextFeatureClass()
	rv := objc.Send[CITextFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array containing additional features detected within the feature.
//
// # Discussion
// 
// A text detector can identify both a major region that is likely to contain
// text as well as the areas within that region that likely to contain
// individual text features. Such features might be single characters, groups
// of closely-packed characters, or entire words.
// 
// To detect sub-features, `/CIDetector/` needs to be called with the
// [CIDetectorReturnSubFeatures] option set to true.
//
// [CIDetectorReturnSubFeatures]: https://developer.apple.com/documentation/CoreImage/CIDetectorReturnSubFeatures
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature/subFeatures
func (t CITextFeature) SubFeatures() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("subFeatures"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// The image coordinate of the lower-left corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomLeft
func (t CITextFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the lower-right corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomRight
func (t CITextFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the upper-left corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature/topLeft
func (t CITextFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the upper-right corner of the detected text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextFeature/topRight
func (t CITextFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}
// A detector that searches for text in a still image or video, returning
//
// See: https://developer.apple.com/documentation/coreimage/cidetectortypetext
func (t CITextFeature) CIDetectorTypeText() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("CIDetectorTypeText"))
	return foundation.NSStringFromID(rv).String()
}

