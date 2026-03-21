// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIDetector] class.
var (
	_CIDetectorClass     CIDetectorClass
	_CIDetectorClassOnce sync.Once
)

func getCIDetectorClass() CIDetectorClass {
	_CIDetectorClassOnce.Do(func() {
		_CIDetectorClass = CIDetectorClass{class: objc.GetClass("CIDetector")}
	})
	return _CIDetectorClass
}

// GetCIDetectorClass returns the class object for CIDetector.
func GetCIDetectorClass() CIDetectorClass {
	return getCIDetectorClass()
}

type CIDetectorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIDetectorClass) Alloc() CIDetector {
	rv := objc.Send[CIDetector](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An image processor that identifies notable features, such as faces and
// barcodes, in a still image or video.
//
// # Overview
// 
// A [CIDetector] object uses image processing to search for and identify
// notable features (faces, rectangles, and barcodes) in a still image or
// video. Detected features are represented by [CIFeature] objects that
// provide more information about each feature.
// 
// This class can maintain many state variables that can impact performance.
// So for best performance, reuse [CIDetector] instances instead of creating
// new ones.
//
// # Using a Detector Object to Find Features
//
//   - [CIDetector.FeaturesInImage]: Searches for features in an image.
//   - [CIDetector.FeaturesInImageOptions]: Searches for features in an image based on the specified image orientation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDetector
type CIDetector struct {
	objectivec.Object
}

// CIDetectorFromID constructs a [CIDetector] from an objc.ID.
//
// An image processor that identifies notable features, such as faces and
// barcodes, in a still image or video.
func CIDetectorFromID(id objc.ID) CIDetector {
	return CIDetector{objectivec.Object{ID: id}}
}
// NOTE: CIDetector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIDetector] class.
//
// # Using a Detector Object to Find Features
//
//   - [ICIDetector.FeaturesInImage]: Searches for features in an image.
//   - [ICIDetector.FeaturesInImageOptions]: Searches for features in an image based on the specified image orientation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDetector
type ICIDetector interface {
	objectivec.IObject

	// Topic: Using a Detector Object to Find Features

	// Searches for features in an image.
	FeaturesInImage(image ICIImage) []CIFeature
	// Searches for features in an image based on the specified image orientation.
	FeaturesInImageOptions(image ICIImage, options foundation.INSDictionary) []CIFeature
}

// Init initializes the instance.
func (d CIDetector) Init() CIDetector {
	rv := objc.Send[CIDetector](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d CIDetector) Autorelease() CIDetector {
	rv := objc.Send[CIDetector](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIDetector creates a new CIDetector instance.
func NewCIDetector() CIDetector {
	class := getCIDetectorClass()
	rv := objc.Send[CIDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a configured detector.
//
// type: A string indicating the kind of detector you are interested in. See
// [Detector Types].
// //
// [Detector Types]: https://developer.apple.com/documentation/CoreImage/detector-types
//
// context: A Core Image context that the detector can use when analyzing an image.
//
// options: A dictionary containing details on how you want the detector to be
// configured. See [Detector Configuration Keys].
// //
// [Detector Configuration Keys]: https://developer.apple.com/documentation/CoreImage/detector-configuration-keys
//
// # Return Value
// 
// A configured detector.
//
// # Discussion
// 
// A [CIDetector] object can potentially create and hold a significant amount
// of resources. Where possible, reuse the same [CIDetector] instance. Also,
// when processing images with a detector object, your application performs
// better if the [CIContext] used to initialize the detector is the same
// context used to process the [ciImage] objects.
//
// [ciImage]: https://developer.apple.com/documentation/UIKit/UIImage/ciImage
//
// See: https://developer.apple.com/documentation/CoreImage/CIDetector/init(ofType:context:options:)
func NewDetectorOfTypeContextOptions(type_ string, context ICIContext, options foundation.INSDictionary) CIDetector {
	rv := objc.Send[objc.ID](objc.ID(getCIDetectorClass().class), objc.Sel("detectorOfType:context:options:"), objc.String(type_), context, options)
	return CIDetectorFromID(rv)
}

// Searches for features in an image.
//
// image: The image you want to examine.
//
// # Return Value
// 
// An array of [CIFeature] objects. Each object represents a feature detected
// in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:)
func (d CIDetector) FeaturesInImage(image ICIImage) []CIFeature {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("featuresInImage:"), image)
	return objc.ConvertSlice(rv, func(id objc.ID) CIFeature {
		return CIFeatureFromID(id)
	})
}
// Searches for features in an image based on the specified image orientation.
//
// image: The image you want to examine.
//
// options: A dictionary that specifies feature detection options. See [Feature
// Detection Keys] for allowed keys and their possible values.
// //
// [Feature Detection Keys]: https://developer.apple.com/documentation/CoreImage/feature-detection-keys
//
// # Return Value
// 
// An array of [CIFeature] objects. Each object represents a feature detected
// in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:options:)
func (d CIDetector) FeaturesInImageOptions(image ICIImage, options foundation.INSDictionary) []CIFeature {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("featuresInImage:options:"), image, options)
	return objc.ConvertSlice(rv, func(id objc.ID) CIFeature {
		return CIFeatureFromID(id)
	})
}

