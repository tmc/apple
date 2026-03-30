// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNRectangleObservation] class.
var (
	_VNRectangleObservationClass     VNRectangleObservationClass
	_VNRectangleObservationClassOnce sync.Once
)

func getVNRectangleObservationClass() VNRectangleObservationClass {
	_VNRectangleObservationClassOnce.Do(func() {
		_VNRectangleObservationClass = VNRectangleObservationClass{class: objc.GetClass("VNRectangleObservation")}
	})
	return _VNRectangleObservationClass
}

// GetVNRectangleObservationClass returns the class object for VNRectangleObservation.
func GetVNRectangleObservationClass() VNRectangleObservationClass {
	return getVNRectangleObservationClass()
}

type VNRectangleObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRectangleObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRectangleObservationClass) Alloc() VNRectangleObservation {
	rv := objc.Send[VNRectangleObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the four vertices of a detected rectangle.
//
// # Accessing the Coordinates
//
//   - [VNRectangleObservation.BottomLeft]: The coordinates of the lower-left corner of the observation bounding box.
//   - [VNRectangleObservation.BottomRight]: The coordinates of the lower-right corner of the observation bounding box.
//   - [VNRectangleObservation.TopLeft]: The coordinates of the upper-left corner of the observation bounding box.
//   - [VNRectangleObservation.TopRight]: The coordinates of the upper-right corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation
type VNRectangleObservation struct {
	VNDetectedObjectObservation
}

// VNRectangleObservationFromID constructs a [VNRectangleObservation] from an objc.ID.
//
// An object that represents the four vertices of a detected rectangle.
func VNRectangleObservationFromID(id objc.ID) VNRectangleObservation {
	return VNRectangleObservation{VNDetectedObjectObservation: VNDetectedObjectObservationFromID(id)}
}

// NOTE: VNRectangleObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRectangleObservation] class.
//
// # Accessing the Coordinates
//
//   - [IVNRectangleObservation.BottomLeft]: The coordinates of the lower-left corner of the observation bounding box.
//   - [IVNRectangleObservation.BottomRight]: The coordinates of the lower-right corner of the observation bounding box.
//   - [IVNRectangleObservation.TopLeft]: The coordinates of the upper-left corner of the observation bounding box.
//   - [IVNRectangleObservation.TopRight]: The coordinates of the upper-right corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation
type IVNRectangleObservation interface {
	IVNDetectedObjectObservation

	// Topic: Accessing the Coordinates

	// The coordinates of the lower-left corner of the observation bounding box.
	BottomLeft() corefoundation.CGPoint
	// The coordinates of the lower-right corner of the observation bounding box.
	BottomRight() corefoundation.CGPoint
	// The coordinates of the upper-left corner of the observation bounding box.
	TopLeft() corefoundation.CGPoint
	// The coordinates of the upper-right corner of the observation bounding box.
	TopRight() corefoundation.CGPoint

	// The results of a document segmentation request.
	Results() IVNRectangleObservation
	SetResults(value IVNRectangleObservation)
}

// Init initializes the instance.
func (r VNRectangleObservation) Init() VNRectangleObservation {
	rv := objc.Send[VNRectangleObservation](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRectangleObservation) Autorelease() VNRectangleObservation {
	rv := objc.Send[VNRectangleObservation](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRectangleObservation creates a new VNRectangleObservation instance.
func NewVNRectangleObservation() VNRectangleObservation {
	class := getVNRectangleObservationClass()
	rv := objc.Send[VNRectangleObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewRectangleObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNRectangleObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRectangleObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNRectangleObservationFromID(rv)
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
func NewRectangleObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNRectangleObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRectangleObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNRectangleObservationFromID(rv)
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
func NewRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) VNRectangleObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return VNRectangleObservationFromID(rv)
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
func NewRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) VNRectangleObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return VNRectangleObservationFromID(rv)
}

// The coordinates of the lower-left corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/bottomLeft
func (r VNRectangleObservation) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}

// The coordinates of the lower-right corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/bottomRight
func (r VNRectangleObservation) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}

// The coordinates of the upper-left corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/topLeft
func (r VNRectangleObservation) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}

// The coordinates of the upper-right corner of the observation bounding box.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/topRight
func (r VNRectangleObservation) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}

// The results of a document segmentation request.
//
// See: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (r VNRectangleObservation) Results() IVNRectangleObservation {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("results"))
	return VNRectangleObservationFromID(objc.ID(rv))
}
func (r VNRectangleObservation) SetResults(value IVNRectangleObservation) {
	objc.Send[struct{}](r.ID, objc.Sel("setResults:"), value)
}
