// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNTextObservation] class.
var (
	_VNTextObservationClass     VNTextObservationClass
	_VNTextObservationClassOnce sync.Once
)

func getVNTextObservationClass() VNTextObservationClass {
	_VNTextObservationClassOnce.Do(func() {
		_VNTextObservationClass = VNTextObservationClass{class: objc.GetClass("VNTextObservation")}
	})
	return _VNTextObservationClass
}

// GetVNTextObservationClass returns the class object for VNTextObservation.
func GetVNTextObservationClass() VNTextObservationClass {
	return getVNTextObservationClass()
}

type VNTextObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNTextObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTextObservationClass) Alloc() VNTextObservation {
	rv := objc.Send[VNTextObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Information about regions of text that an image-analysis request detects.
//
// # Overview
//
// This type of observation results from a [VNDetectTextRectanglesRequest]. It
// expresses the location of each detected character by its bounding box.
//
// # Finding Individual Characters
//
//   - [VNTextObservation.CharacterBoxes]: An array of detected individual character bounding boxes.
//
// See: https://developer.apple.com/documentation/Vision/VNTextObservation
type VNTextObservation struct {
	VNRectangleObservation
}

// VNTextObservationFromID constructs a [VNTextObservation] from an objc.ID.
//
// Information about regions of text that an image-analysis request detects.
func VNTextObservationFromID(id objc.ID) VNTextObservation {
	return VNTextObservation{VNRectangleObservation: VNRectangleObservationFromID(id)}
}

// NOTE: VNTextObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTextObservation] class.
//
// # Finding Individual Characters
//
//   - [IVNTextObservation.CharacterBoxes]: An array of detected individual character bounding boxes.
//
// See: https://developer.apple.com/documentation/Vision/VNTextObservation
type IVNTextObservation interface {
	IVNRectangleObservation

	// Topic: Finding Individual Characters

	// An array of detected individual character bounding boxes.
	CharacterBoxes() []VNRectangleObservation
}

// Init initializes the instance.
func (t VNTextObservation) Init() VNTextObservation {
	rv := objc.Send[VNTextObservation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTextObservation) Autorelease() VNTextObservation {
	rv := objc.Send[VNTextObservation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTextObservation creates a new VNTextObservation instance.
func NewVNTextObservation() VNTextObservation {
	class := getVNTextObservationClass()
	rv := objc.Send[VNTextObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a rectangle observation from its corner points.
//
// requestRevision: The rectangle detector revision number. A higher revision indicates more
// recent iterations of the framework.
//
// topLeft: The upper-left corner point.
//
// bottomLeft: The lower-left corner point.
//
// bottomRight: The lower-right corner point.
//
// topRight: The upper-right corner point.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func NewTextObservationRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) VNTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNTextObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return VNTextObservationFromID(rv)
}

// Creates a rectangle observation from its corner points.
//
// requestRevision: The rectangle detector revision number. A higher revision indicates more
// recent iterations of the framework.
//
// topLeft: The upper-left corner point.
//
// topRight: The upper-right corner point.
//
// bottomRight: The lower-right corner point.
//
// bottomLeft: The lower-left corner point.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func NewTextObservationRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) VNTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNTextObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return VNTextObservationFromID(rv)
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewTextObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNTextObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNTextObservationFromID(rv)
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
func NewTextObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNTextObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNTextObservationFromID(rv)
}

// An array of detected individual character bounding boxes.
//
// # Discussion
//
// If the associated [VNDetectTextRectanglesRequest] request indicates
// interest in character boxes by setting the option `reportCharacterBoxes` to
// true, this property is non-`nil`. If no characters are found, it remains
// empty.
//
// See: https://developer.apple.com/documentation/Vision/VNTextObservation/characterBoxes
func (t VNTextObservation) CharacterBoxes() []VNRectangleObservation {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("characterBoxes"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNRectangleObservation {
		return VNRectangleObservationFromID(id)
	})
}
