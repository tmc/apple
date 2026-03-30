// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNFaceObservation] class.
var (
	_VNFaceObservationClass     VNFaceObservationClass
	_VNFaceObservationClassOnce sync.Once
)

func getVNFaceObservationClass() VNFaceObservationClass {
	_VNFaceObservationClassOnce.Do(func() {
		_VNFaceObservationClass = VNFaceObservationClass{class: objc.GetClass("VNFaceObservation")}
	})
	return _VNFaceObservationClass
}

// GetVNFaceObservationClass returns the class object for VNFaceObservation.
func GetVNFaceObservationClass() VNFaceObservationClass {
	return getVNFaceObservationClass()
}

type VNFaceObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNFaceObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFaceObservationClass) Alloc() VNFaceObservation {
	rv := objc.Send[VNFaceObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Face or facial-feature information that an image analysis request detects.
//
// # Overview
//
// This type of observation results from a [VNDetectFaceRectanglesRequest]. It
// contains information about facial landmarks and regions it finds in the
// image.
//
// # Identifying Landmarks
//
//   - [VNFaceObservation.Landmarks]: The facial features of the detected face.
//
// # Determining Facial Orientation
//
//   - [VNFaceObservation.Roll]: The roll angle of a face in radians.
//   - [VNFaceObservation.Yaw]: The yaw angle of a face in radians.
//   - [VNFaceObservation.Pitch]: The pitch angle of a face in radians.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation
type VNFaceObservation struct {
	VNDetectedObjectObservation
}

// VNFaceObservationFromID constructs a [VNFaceObservation] from an objc.ID.
//
// Face or facial-feature information that an image analysis request detects.
func VNFaceObservationFromID(id objc.ID) VNFaceObservation {
	return VNFaceObservation{VNDetectedObjectObservation: VNDetectedObjectObservationFromID(id)}
}

// NOTE: VNFaceObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNFaceObservation] class.
//
// # Identifying Landmarks
//
//   - [IVNFaceObservation.Landmarks]: The facial features of the detected face.
//
// # Determining Facial Orientation
//
//   - [IVNFaceObservation.Roll]: The roll angle of a face in radians.
//   - [IVNFaceObservation.Yaw]: The yaw angle of a face in radians.
//   - [IVNFaceObservation.Pitch]: The pitch angle of a face in radians.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation
type IVNFaceObservation interface {
	IVNDetectedObjectObservation

	// Topic: Identifying Landmarks

	// The facial features of the detected face.
	Landmarks() IVNFaceLandmarks2D

	// Topic: Determining Facial Orientation

	// The roll angle of a face in radians.
	Roll() foundation.NSNumber
	// The yaw angle of a face in radians.
	Yaw() foundation.NSNumber
	// The pitch angle of a face in radians.
	Pitch() foundation.NSNumber

	// A value that indicates the quality of the face capture.
	FaceCaptureQuality() foundation.NSNumber
	// The results of the face-capture quality request.
	Results() IVNFaceObservation
	SetResults(value IVNFaceObservation)
}

// Init initializes the instance.
func (f VNFaceObservation) Init() VNFaceObservation {
	rv := objc.Send[VNFaceObservation](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFaceObservation) Autorelease() VNFaceObservation {
	rv := objc.Send[VNFaceObservation](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFaceObservation creates a new VNFaceObservation instance.
func NewVNFaceObservation() VNFaceObservation {
	class := getVNFaceObservationClass()
	rv := objc.Send[VNFaceObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewFaceObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNFaceObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNFaceObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNFaceObservationFromID(rv)
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
func NewFaceObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNFaceObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNFaceObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNFaceObservationFromID(rv)
}

// Creates an observation that contains the roll, yaw, and pitch of the face.
//
// requestRevision: The revision of the request.
//
// boundingBox: The bounding rectangle of the detected face landmark.
//
// roll: The rotational angle of the face landmark around the z-axis.
//
// yaw: The rotational angle of the face landmark around the y-axis.
//
// pitch: The rotational angle of the face landmark around the x-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/init(requestRevision:boundingBox:roll:yaw:pitch:)
func NewFaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision uint, boundingBox corefoundation.CGRect, roll foundation.NSNumber, yaw foundation.NSNumber, pitch foundation.NSNumber) VNFaceObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNFaceObservationClass().class), objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:pitch:"), requestRevision, boundingBox, roll, yaw, pitch)
	return VNFaceObservationFromID(rv)
}

// The facial features of the detected face.
//
// # Discussion
//
// This value is `nil` for face observations produced by a
// [VNDetectFaceRectanglesRequest] analysis. Use the
// [VNDetectFaceLandmarksRequest] class to find facial features.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/landmarks
func (f VNFaceObservation) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("landmarks"))
	return VNFaceLandmarks2DFromID(objc.ID(rv))
}

// The roll angle of a face in radians.
//
// # Discussion
//
// This value indicates the rotational angle of the face around the z-axis.
//
// If the request doesn’t calculate the angle, the value is `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/roll
func (f VNFaceObservation) Roll() foundation.NSNumber {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("roll"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The yaw angle of a face in radians.
//
// # Discussion
//
// This value indicates the rotational angle of the face around the y-axis.
//
// If the request doesn’t calculate the angle, the value is `nil.`
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/yaw
func (f VNFaceObservation) Yaw() foundation.NSNumber {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("yaw"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The pitch angle of a face in radians.
//
// # Discussion
//
// This value indicates the rotational angle of the face around the x-axis.
//
// If the request doesn’t calculate the angle, the value is `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/pitch
func (f VNFaceObservation) Pitch() foundation.NSNumber {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("pitch"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A value that indicates the quality of the face capture.
//
// # Discussion
//
// The capture quality of the face allows you to compare the quality of the
// face in terms of its capture attributes: lighting, blur, and prime
// positioning. Use this value to compare the capture quality of a face
// against other captures of the same face in a specified set.
//
// The value of this property value ranges from `0.0` to `1.0`. Faces with
// quality closer to `1.0` are better lit, sharper, and more centrally
// positioned than faces with quality closer to `0.0`.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservation/faceCaptureQuality-2o4xv
func (f VNFaceObservation) FaceCaptureQuality() foundation.NSNumber {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("faceCaptureQuality"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The results of the face-capture quality request.
//
// See: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequest/results
func (f VNFaceObservation) Results() IVNFaceObservation {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("results"))
	return VNFaceObservationFromID(objc.ID(rv))
}
func (f VNFaceObservation) SetResults(value IVNFaceObservation) {
	objc.Send[struct{}](f.ID, objc.Sel("setResults:"), value)
}
