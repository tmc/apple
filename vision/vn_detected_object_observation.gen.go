// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VNDetectedObjectObservation] class.
var (
	_VNDetectedObjectObservationClass     VNDetectedObjectObservationClass
	_VNDetectedObjectObservationClassOnce sync.Once
)

func getVNDetectedObjectObservationClass() VNDetectedObjectObservationClass {
	_VNDetectedObjectObservationClassOnce.Do(func() {
		_VNDetectedObjectObservationClass = VNDetectedObjectObservationClass{class: objc.GetClass("VNDetectedObjectObservation")}
	})
	return _VNDetectedObjectObservationClass
}

// GetVNDetectedObjectObservationClass returns the class object for VNDetectedObjectObservation.
func GetVNDetectedObjectObservationClass() VNDetectedObjectObservationClass {
	return getVNDetectedObjectObservationClass()
}

type VNDetectedObjectObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectedObjectObservationClass) Alloc() VNDetectedObjectObservation {
	rv := objc.Send[VNDetectedObjectObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that provides the position and extent of an image feature
// that an image- analysis request detects.
//
// # Overview
// 
// This class is the observation type that [VNTrackObjectRequest] generates.
// It represents an object that the Vision request detects and tracks.
//
// # Locating a Detected Object
//
//   - [VNDetectedObjectObservation.BoundingBox]: The bounding box of the object that the request detects.
//
// # Accessing an Image Mask
//
//   - [VNDetectedObjectObservation.GlobalSegmentationMask]: A resulting pixel buffer from a request to generate a segmentation mask for an image.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation
type VNDetectedObjectObservation struct {
	VNObservation
}

// VNDetectedObjectObservationFromID constructs a [VNDetectedObjectObservation] from an objc.ID.
//
// An observation that provides the position and extent of an image feature
// that an image- analysis request detects.
func VNDetectedObjectObservationFromID(id objc.ID) VNDetectedObjectObservation {
	return VNDetectedObjectObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNDetectedObjectObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectedObjectObservation] class.
//
// # Locating a Detected Object
//
//   - [IVNDetectedObjectObservation.BoundingBox]: The bounding box of the object that the request detects.
//
// # Accessing an Image Mask
//
//   - [IVNDetectedObjectObservation.GlobalSegmentationMask]: A resulting pixel buffer from a request to generate a segmentation mask for an image.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation
type IVNDetectedObjectObservation interface {
	IVNObservation

	// Topic: Locating a Detected Object

	// The bounding box of the object that the request detects.
	BoundingBox() corefoundation.CGRect

	// Topic: Accessing an Image Mask

	// A resulting pixel buffer from a request to generate a segmentation mask for an image.
	GlobalSegmentationMask() IVNPixelBufferObservation
}

// Init initializes the instance.
func (d VNDetectedObjectObservation) Init() VNDetectedObjectObservation {
	rv := objc.Send[VNDetectedObjectObservation](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectedObjectObservation) Autorelease() VNDetectedObjectObservation {
	rv := objc.Send[VNDetectedObjectObservation](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectedObjectObservation creates a new VNDetectedObjectObservation instance.
func NewVNDetectedObjectObservation() VNDetectedObjectObservation {
	class := getVNDetectedObjectObservationClass()
	rv := objc.Send[VNDetectedObjectObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewDetectedObjectObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNDetectedObjectObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNDetectedObjectObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNDetectedObjectObservationFromID(rv)
}

// Creates an observation with a revision number and bounding box.
//
// requestRevision: The revision of the request to use.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func NewDetectedObjectObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNDetectedObjectObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNDetectedObjectObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNDetectedObjectObservationFromID(rv)
}

// The bounding box of the object that the request detects.
//
// # Discussion
// 
// The system normalizes the coordinates to the dimensions of the processed
// image, with the origin at the lower-left corner of the image.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/boundingBox
func (d VNDetectedObjectObservation) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d.ID, objc.Sel("boundingBox"))
	return corefoundation.CGRect(rv)
}
// A resulting pixel buffer from a request to generate a segmentation mask for
// an image.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/globalSegmentationMask
func (d VNDetectedObjectObservation) GlobalSegmentationMask() IVNPixelBufferObservation {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("globalSegmentationMask"))
	return VNPixelBufferObservationFromID(objc.ID(rv))
}

