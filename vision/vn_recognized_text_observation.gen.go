// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNRecognizedTextObservation] class.
var (
	_VNRecognizedTextObservationClass     VNRecognizedTextObservationClass
	_VNRecognizedTextObservationClassOnce sync.Once
)

func getVNRecognizedTextObservationClass() VNRecognizedTextObservationClass {
	_VNRecognizedTextObservationClassOnce.Do(func() {
		_VNRecognizedTextObservationClass = VNRecognizedTextObservationClass{class: objc.GetClass("VNRecognizedTextObservation")}
	})
	return _VNRecognizedTextObservationClass
}

// GetVNRecognizedTextObservationClass returns the class object for VNRecognizedTextObservation.
func GetVNRecognizedTextObservationClass() VNRecognizedTextObservationClass {
	return getVNRecognizedTextObservationClass()
}

type VNRecognizedTextObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedTextObservationClass) Alloc() VNRecognizedTextObservation {
	rv := objc.Send[VNRecognizedTextObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A request that detects and recognizes regions of text in an image.
//
// # Overview
// 
// This type of observation results from a [VNRecognizeTextRequest]. It
// contains information about both the location and content of text and glyphs
// that Vision recognized in the input image.
//
// # Obtaining Recognized Text
//
//   - [VNRecognizedTextObservation.TopCandidates]: Requests the  top candidates for a recognized text string.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation
type VNRecognizedTextObservation struct {
	VNRectangleObservation
}

// VNRecognizedTextObservationFromID constructs a [VNRecognizedTextObservation] from an objc.ID.
//
// A request that detects and recognizes regions of text in an image.
func VNRecognizedTextObservationFromID(id objc.ID) VNRecognizedTextObservation {
	return VNRecognizedTextObservation{VNRectangleObservation: VNRectangleObservationFromID(id)}
}
// NOTE: VNRecognizedTextObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNRecognizedTextObservation] class.
//
// # Obtaining Recognized Text
//
//   - [IVNRecognizedTextObservation.TopCandidates]: Requests the  top candidates for a recognized text string.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation
type IVNRecognizedTextObservation interface {
	IVNRectangleObservation

	// Topic: Obtaining Recognized Text

	// Requests the  top candidates for a recognized text string.
	TopCandidates(maxCandidateCount uint) []VNRecognizedText

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (r VNRecognizedTextObservation) Init() VNRecognizedTextObservation {
	rv := objc.Send[VNRecognizedTextObservation](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedTextObservation) Autorelease() VNRecognizedTextObservation {
	rv := objc.Send[VNRecognizedTextObservation](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedTextObservation creates a new VNRecognizedTextObservation instance.
func NewVNRecognizedTextObservation() VNRecognizedTextObservation {
	class := getVNRecognizedTextObservationClass()
	rv := objc.Send[VNRecognizedTextObservation](objc.ID(class.class), objc.Sel("new"))
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
func NewRecognizedTextObservationRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) VNRecognizedTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedTextObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return VNRecognizedTextObservationFromID(rv)
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
func NewRecognizedTextObservationRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) VNRecognizedTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedTextObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return VNRecognizedTextObservationFromID(rv)
}


// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewRecognizedTextObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNRecognizedTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedTextObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNRecognizedTextObservationFromID(rv)
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
func NewRecognizedTextObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNRecognizedTextObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedTextObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNRecognizedTextObservationFromID(rv)
}







// Requests the top candidates for a recognized text string.
//
// maxCandidateCount: The maximum number of candidates to return. This can’t exceed 10.
//
// # Return Value
// 
// An array of the top candidates, sorted by decreasing confidence score.
//
// # Discussion
// 
// This function returns no more than candidates, but it may return fewer than
// candidates.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation/topCandidates(_:)
func (r VNRecognizedTextObservation) TopCandidates(maxCandidateCount uint) []VNRecognizedText {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("topCandidates:"), maxCandidateCount)
	return objc.ConvertSlice(rv, func(id objc.ID) VNRecognizedText {
		return VNRecognizedTextFromID(id)
	})
}
func (r VNRecognizedTextObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}




































