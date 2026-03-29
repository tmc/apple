// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNRecognizedText] class.
var (
	_VNRecognizedTextClass     VNRecognizedTextClass
	_VNRecognizedTextClassOnce sync.Once
)

func getVNRecognizedTextClass() VNRecognizedTextClass {
	_VNRecognizedTextClassOnce.Do(func() {
		_VNRecognizedTextClass = VNRecognizedTextClass{class: objc.GetClass("VNRecognizedText")}
	})
	return _VNRecognizedTextClass
}

// GetVNRecognizedTextClass returns the class object for VNRecognizedText.
func GetVNRecognizedTextClass() VNRecognizedTextClass {
	return getVNRecognizedTextClass()
}

type VNRecognizedTextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRecognizedTextClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedTextClass) Alloc() VNRecognizedText {
	rv := objc.Send[VNRecognizedText](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Text recognized in an image through a text recognition request.
//
// # Overview
// 
// A single [VNRecognizedTextObservation] can contain multiple recognized text
// objects—one for each candidate.
//
// # Determining Recognized Text
//
//   - [VNRecognizedText.String]: The top candidate for recognized text.
//   - [VNRecognizedText.Confidence]: A normalized confidence score for the text recognition result.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedText
type VNRecognizedText struct {
	objectivec.Object
}

// VNRecognizedTextFromID constructs a [VNRecognizedText] from an objc.ID.
//
// Text recognized in an image through a text recognition request.
func VNRecognizedTextFromID(id objc.ID) VNRecognizedText {
	return VNRecognizedText{objectivec.Object{ID: id}}
}
// NOTE: VNRecognizedText adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizedText] class.
//
// # Determining Recognized Text
//
//   - [IVNRecognizedText.String]: The top candidate for recognized text.
//   - [IVNRecognizedText.Confidence]: A normalized confidence score for the text recognition result.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedText
type IVNRecognizedText interface {
	objectivec.IObject
	VNRequestRevisionProviding

	// Topic: Determining Recognized Text

	// The top candidate for recognized text.
	String() string
	// A normalized confidence score for the text recognition result.
	Confidence() VNConfidence

	// Calculates the bounding box around the characters in the range of a string.
	BoundingBoxForRangeError(range_ foundation.NSRange) (IVNRectangleObservation, error)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r VNRecognizedText) Init() VNRecognizedText {
	rv := objc.Send[VNRecognizedText](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedText) Autorelease() VNRecognizedText {
	rv := objc.Send[VNRecognizedText](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedText creates a new VNRecognizedText instance.
func NewVNRecognizedText() VNRecognizedText {
	class := getVNRecognizedTextClass()
	rv := objc.Send[VNRecognizedText](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Calculates the bounding box around the characters in the range of a string.
//
// range: The range of the characters in the text string to draw a bounding box
// around.
//
// error: An error that contains information about failed computation of the bounding
// box, or `nil` if no error occurred.
//
// # Discussion
// 
// Bounding boxes aren’t always an exact fit around the characters. Use them
// to display in user interfaces to provide general guidance, but avoid using
// their contents for image processing.
// 
// The bounding box that returns from this method differs based on the value
// of [RecognitionLevel], as follows:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedText/boundingBoxForRange:error:
func (r VNRecognizedText) BoundingBoxForRangeError(range_ foundation.NSRange) (IVNRectangleObservation, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("boundingBoxForRange:error:"), range_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNRectangleObservation{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNRectangleObservationFromID(rv), nil

}
func (r VNRecognizedText) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The top candidate for recognized text.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedText/string
func (r VNRecognizedText) String() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
// A normalized confidence score for the text recognition result.
//
// # Discussion
// 
// The confidence level is a normalized value between `0.0` and `1.0`, where
// `1.0` represents the highest confidence.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedText/confidence
func (r VNRecognizedText) Confidence() VNConfidence {
	rv := objc.Send[VNConfidence](r.ID, objc.Sel("confidence"))
	return VNConfidence(rv)
}
// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (r VNRecognizedText) RequestRevision() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("requestRevision"))
	return rv
}

			// Protocol methods for VNRequestRevisionProviding
			

