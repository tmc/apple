// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CIFaceFeature] class.
var (
	_CIFaceFeatureClass     CIFaceFeatureClass
	_CIFaceFeatureClassOnce sync.Once
)

func getCIFaceFeatureClass() CIFaceFeatureClass {
	_CIFaceFeatureClassOnce.Do(func() {
		_CIFaceFeatureClass = CIFaceFeatureClass{class: objc.GetClass("CIFaceFeature")}
	})
	return _CIFaceFeatureClass
}

// GetCIFaceFeatureClass returns the class object for CIFaceFeature.
func GetCIFaceFeatureClass() CIFaceFeatureClass {
	return getCIFaceFeatureClass()
}

type CIFaceFeatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIFaceFeatureClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIFaceFeatureClass) Alloc() CIFaceFeature {
	rv := objc.Send[CIFaceFeature](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about a face detected in a still or video image.
//
// # Overview
//
// The properties of a [CIFaceFeature] object provide information about the
// face’s eyes and mouth. A face object in a video can also have properties
// that track its location over time, tracking ID and frame count.
//
// # Locating Faces
//
//   - [CIFaceFeature.HasFaceAngle]: A Boolean value that indicates whether information about face rotation is available.
//   - [CIFaceFeature.FaceAngle]: The rotation of the face.
//
// # Identifying Facial Features
//
//   - [CIFaceFeature.HasLeftEyePosition]: A Boolean value that indicates whether the detector found the face’s left eye.
//   - [CIFaceFeature.HasRightEyePosition]: A Boolean value that indicates whether the detector found the face’s right eye.
//   - [CIFaceFeature.HasMouthPosition]: A Boolean value that indicates whether the detector found the face’s mouth.
//   - [CIFaceFeature.LeftEyePosition]: The image coordinate of the center of the left eye.
//   - [CIFaceFeature.RightEyePosition]: The image coordinate of the center of the right eye.
//   - [CIFaceFeature.MouthPosition]: The image coordinate of the center of the mouth.
//   - [CIFaceFeature.HasSmile]: A Boolean value that indicates whether a smile is detected in the face.
//   - [CIFaceFeature.LeftEyeClosed]: A Boolean value that indicates whether a closed left eye is detected in the face.
//   - [CIFaceFeature.RightEyeClosed]: A Boolean value that indicates whether a closed right eye is detected in the face.
//
// # Tracking Distinct Faces in Video
//
//   - [CIFaceFeature.HasTrackingID]: A Boolean value that indicates whether the face object has a tracking ID.
//   - [CIFaceFeature.TrackingID]: The tracking identifier of the face object.
//   - [CIFaceFeature.HasTrackingFrameCount]: A Boolean value that indicates the face object has a tracking frame count.
//   - [CIFaceFeature.TrackingFrameCount]: The tracking frame count of the face.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature
type CIFaceFeature struct {
	CIFeature
}

// CIFaceFeatureFromID constructs a [CIFaceFeature] from an objc.ID.
//
// Information about a face detected in a still or video image.
func CIFaceFeatureFromID(id objc.ID) CIFaceFeature {
	return CIFaceFeature{CIFeature: CIFeatureFromID(id)}
}

// NOTE: CIFaceFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIFaceFeature] class.
//
// # Locating Faces
//
//   - [ICIFaceFeature.HasFaceAngle]: A Boolean value that indicates whether information about face rotation is available.
//   - [ICIFaceFeature.FaceAngle]: The rotation of the face.
//
// # Identifying Facial Features
//
//   - [ICIFaceFeature.HasLeftEyePosition]: A Boolean value that indicates whether the detector found the face’s left eye.
//   - [ICIFaceFeature.HasRightEyePosition]: A Boolean value that indicates whether the detector found the face’s right eye.
//   - [ICIFaceFeature.HasMouthPosition]: A Boolean value that indicates whether the detector found the face’s mouth.
//   - [ICIFaceFeature.LeftEyePosition]: The image coordinate of the center of the left eye.
//   - [ICIFaceFeature.RightEyePosition]: The image coordinate of the center of the right eye.
//   - [ICIFaceFeature.MouthPosition]: The image coordinate of the center of the mouth.
//   - [ICIFaceFeature.HasSmile]: A Boolean value that indicates whether a smile is detected in the face.
//   - [ICIFaceFeature.LeftEyeClosed]: A Boolean value that indicates whether a closed left eye is detected in the face.
//   - [ICIFaceFeature.RightEyeClosed]: A Boolean value that indicates whether a closed right eye is detected in the face.
//
// # Tracking Distinct Faces in Video
//
//   - [ICIFaceFeature.HasTrackingID]: A Boolean value that indicates whether the face object has a tracking ID.
//   - [ICIFaceFeature.TrackingID]: The tracking identifier of the face object.
//   - [ICIFaceFeature.HasTrackingFrameCount]: A Boolean value that indicates the face object has a tracking frame count.
//   - [ICIFaceFeature.TrackingFrameCount]: The tracking frame count of the face.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature
type ICIFaceFeature interface {
	ICIFeature

	// Topic: Locating Faces

	// A Boolean value that indicates whether information about face rotation is available.
	HasFaceAngle() bool
	// The rotation of the face.
	FaceAngle() float32

	// Topic: Identifying Facial Features

	// A Boolean value that indicates whether the detector found the face’s left eye.
	HasLeftEyePosition() bool
	// A Boolean value that indicates whether the detector found the face’s right eye.
	HasRightEyePosition() bool
	// A Boolean value that indicates whether the detector found the face’s mouth.
	HasMouthPosition() bool
	// The image coordinate of the center of the left eye.
	LeftEyePosition() corefoundation.CGPoint
	// The image coordinate of the center of the right eye.
	RightEyePosition() corefoundation.CGPoint
	// The image coordinate of the center of the mouth.
	MouthPosition() corefoundation.CGPoint
	// A Boolean value that indicates whether a smile is detected in the face.
	HasSmile() bool
	// A Boolean value that indicates whether a closed left eye is detected in the face.
	LeftEyeClosed() bool
	// A Boolean value that indicates whether a closed right eye is detected in the face.
	RightEyeClosed() bool

	// Topic: Tracking Distinct Faces in Video

	// A Boolean value that indicates whether the face object has a tracking ID.
	HasTrackingID() bool
	// The tracking identifier of the face object.
	TrackingID() int
	// A Boolean value that indicates the face object has a tracking frame count.
	HasTrackingFrameCount() bool
	// The tracking frame count of the face.
	TrackingFrameCount() int
}

// Init initializes the instance.
func (f CIFaceFeature) Init() CIFaceFeature {
	rv := objc.Send[CIFaceFeature](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CIFaceFeature) Autorelease() CIFaceFeature {
	rv := objc.Send[CIFaceFeature](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIFaceFeature creates a new CIFaceFeature instance.
func NewCIFaceFeature() CIFaceFeature {
	class := getCIFaceFeatureClass()
	rv := objc.Send[CIFaceFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether information about face rotation is
// available.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-swift.property
func (f CIFaceFeature) HasFaceAngle() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasFaceAngle"))
	return rv
}

// The rotation of the face.
//
// # Discussion
//
// Rotation is measured counterclockwise in degrees, with zero indicating that
// a line drawn between the eyes is horizontal relative to the image
// orientation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-swift.property
func (f CIFaceFeature) FaceAngle() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("faceAngle"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s left
// eye.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-swift.property
func (f CIFaceFeature) HasLeftEyePosition() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasLeftEyePosition"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s
// right eye.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-swift.property
func (f CIFaceFeature) HasRightEyePosition() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasRightEyePosition"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s
// mouth.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-swift.property
func (f CIFaceFeature) HasMouthPosition() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasMouthPosition"))
	return rv
}

// The image coordinate of the center of the left eye.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-swift.property
func (f CIFaceFeature) LeftEyePosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f.ID, objc.Sel("leftEyePosition"))
	return corefoundation.CGPoint(rv)
}

// The image coordinate of the center of the right eye.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-swift.property
func (f CIFaceFeature) RightEyePosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f.ID, objc.Sel("rightEyePosition"))
	return corefoundation.CGPoint(rv)
}

// The image coordinate of the center of the mouth.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-swift.property
func (f CIFaceFeature) MouthPosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f.ID, objc.Sel("mouthPosition"))
	return corefoundation.CGPoint(rv)
}

// A Boolean value that indicates whether a smile is detected in the face.
//
// # Discussion
//
// To detect smiles, `/CIDetector/` needs to be called with the
// [CIDetectorSmile] option set to true.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-swift.property
//
// [CIDetectorSmile]: https://developer.apple.com/documentation/CoreImage/CIDetectorSmile
func (f CIFaceFeature) HasSmile() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasSmile"))
	return rv
}

// A Boolean value that indicates whether a closed left eye is detected in the
// face.
//
// # Discussion
//
// To detect closed eyes, `/CIDetector/` needs to be called with the
// [CIDetectorEyeBlink] option set to true.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-swift.property
//
// [CIDetectorEyeBlink]: https://developer.apple.com/documentation/CoreImage/CIDetectorEyeBlink
func (f CIFaceFeature) LeftEyeClosed() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("leftEyeClosed"))
	return rv
}

// A Boolean value that indicates whether a closed right eye is detected in
// the face.
//
// # Discussion
//
// To detect closed eyes, `/CIDetector/` needs to be called with the
// [CIDetectorEyeBlink] option set to true.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-swift.property
//
// [CIDetectorEyeBlink]: https://developer.apple.com/documentation/CoreImage/CIDetectorEyeBlink
func (f CIFaceFeature) RightEyeClosed() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("rightEyeClosed"))
	return rv
}

// A Boolean value that indicates whether the face object has a tracking ID.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-swift.property
func (f CIFaceFeature) HasTrackingID() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasTrackingID"))
	return rv
}

// The tracking identifier of the face object.
//
// # Discussion
//
// Core Image provides a tracking identifier for faces it detects in a video
// stream, which you can use to identify when a CIFaceFeature objects detected
// in one video frame is the same face detected in a previous video frame.
//
// This identifier persists only as long as a face is in the frame and is not
// associated with a specific face. In other words, if a face moves out of the
// video frame and comes back into the frame later, another ID is assigned.
// (Core Image detects faces, but does not recognize specific faces.)
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-swift.property
func (f CIFaceFeature) TrackingID() int {
	rv := objc.Send[int](f.ID, objc.Sel("trackingID"))
	return rv
}

// A Boolean value that indicates the face object has a tracking frame count.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-swift.property
func (f CIFaceFeature) HasTrackingFrameCount() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasTrackingFrameCount"))
	return rv
}

// The tracking frame count of the face.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-swift.property
func (f CIFaceFeature) TrackingFrameCount() int {
	rv := objc.Send[int](f.ID, objc.Sel("trackingFrameCount"))
	return rv
}
