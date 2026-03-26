// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataFaceObject] class.
var (
	_AVMetadataFaceObjectClass     AVMetadataFaceObjectClass
	_AVMetadataFaceObjectClassOnce sync.Once
)

func getAVMetadataFaceObjectClass() AVMetadataFaceObjectClass {
	_AVMetadataFaceObjectClassOnce.Do(func() {
		_AVMetadataFaceObjectClass = AVMetadataFaceObjectClass{class: objc.GetClass("AVMetadataFaceObject")}
	})
	return _AVMetadataFaceObjectClass
}

// GetAVMetadataFaceObjectClass returns the class object for AVMetadataFaceObject.
func GetAVMetadataFaceObjectClass() AVMetadataFaceObjectClass {
	return getAVMetadataFaceObjectClass()
}

type AVMetadataFaceObjectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataFaceObjectClass) Alloc() AVMetadataFaceObject {
	rv := objc.Send[AVMetadataFaceObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Face information detected by a metadata capture output.
//
// # Overview
// 
// The [AVMetadataFaceObject] class is a concrete subclass of
// [AVMetadataObject] that defines the features of a single detected face. You
// can retrieve instances of this class from the output of an
// [AVCaptureMetadataOutput] object on devices that support face detection.
//
// # Getting the face identifier
//
//   - [AVMetadataFaceObject.FaceID]: The unique ID for this face metadata object.
//
// # Accessing the face detection data
//
//   - [AVMetadataFaceObject.HasRollAngle]: A Boolean value indicating whether there is a valid roll angle associated with the face.
//   - [AVMetadataFaceObject.RollAngle]: The roll angle of the face specified in degrees.
//   - [AVMetadataFaceObject.HasYawAngle]: A Boolean value indicating whether there is a valid yaw angle associated with the face.
//   - [AVMetadataFaceObject.YawAngle]: The yaw angle of the face specified in degrees.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject
type AVMetadataFaceObject struct {
	AVMetadataObject
}

// AVMetadataFaceObjectFromID constructs a [AVMetadataFaceObject] from an objc.ID.
//
// Face information detected by a metadata capture output.
func AVMetadataFaceObjectFromID(id objc.ID) AVMetadataFaceObject {
	return AVMetadataFaceObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}
// NOTE: AVMetadataFaceObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataFaceObject] class.
//
// # Getting the face identifier
//
//   - [IAVMetadataFaceObject.FaceID]: The unique ID for this face metadata object.
//
// # Accessing the face detection data
//
//   - [IAVMetadataFaceObject.HasRollAngle]: A Boolean value indicating whether there is a valid roll angle associated with the face.
//   - [IAVMetadataFaceObject.RollAngle]: The roll angle of the face specified in degrees.
//   - [IAVMetadataFaceObject.HasYawAngle]: A Boolean value indicating whether there is a valid yaw angle associated with the face.
//   - [IAVMetadataFaceObject.YawAngle]: The yaw angle of the face specified in degrees.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject
type IAVMetadataFaceObject interface {
	IAVMetadataObject

	// Topic: Getting the face identifier

	// The unique ID for this face metadata object.
	FaceID() int

	// Topic: Accessing the face detection data

	// A Boolean value indicating whether there is a valid roll angle associated with the face.
	HasRollAngle() bool
	// The roll angle of the face specified in degrees.
	RollAngle() float64
	// A Boolean value indicating whether there is a valid yaw angle associated with the face.
	HasYawAngle() bool
	// The yaw angle of the face specified in degrees.
	YawAngle() float64
}

// Init initializes the instance.
func (m AVMetadataFaceObject) Init() AVMetadataFaceObject {
	rv := objc.Send[AVMetadataFaceObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataFaceObject) Autorelease() AVMetadataFaceObject {
	rv := objc.Send[AVMetadataFaceObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataFaceObject creates a new AVMetadataFaceObject instance.
func NewAVMetadataFaceObject() AVMetadataFaceObject {
	class := getAVMetadataFaceObjectClass()
	rv := objc.Send[AVMetadataFaceObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The unique ID for this face metadata object.
//
// # Discussion
// 
// Each time a face enters the picture, it is assigned a new unique
// identifier, which you can use to reference the face in your code. Face IDs
// are not reused, and the same face leaving and entering the picture again is
// assigned a new identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/faceID
func (m AVMetadataFaceObject) FaceID() int {
	rv := objc.Send[int](m.ID, objc.Sel("faceID"))
	return rv
}
// A Boolean value indicating whether there is a valid roll angle associated
// with the face.
//
// # Discussion
// 
// If the value of this property is [false], the value in the [RollAngle]
// property is invalid and must not be accessed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasRollAngle
func (m AVMetadataFaceObject) HasRollAngle() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasRollAngle"))
	return rv
}
// The roll angle of the face specified in degrees.
//
// # Discussion
// 
// The roll angle represents the side-to-side tilt of the face relative to the
// metadata’s bounding rectangle. A value of `0.0` yields a face that is
// level relative to the picture, whereas a value of `90` yields a face that
// is perpendicular relative to the picture.
// 
// You must check the value of the [HasRollAngle] property before accessing
// this property. If the value in the [HasRollAngle] property is [false],
// reading the value in this property raises an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/rollAngle
func (m AVMetadataFaceObject) RollAngle() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("rollAngle"))
	return rv
}
// A Boolean value indicating whether there is a valid yaw angle associated
// with the face.
//
// # Discussion
// 
// If the value of this property is [false], the value in the [YawAngle]
// property is invalid and must not be accessed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasYawAngle
func (m AVMetadataFaceObject) HasYawAngle() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasYawAngle"))
	return rv
}
// The yaw angle of the face specified in degrees.
//
// # Discussion
// 
// The yaw angle represents the rotation of the face around the vertical axis.
// A value of `0.0` yields a face that is looking directly at the camera,
// whereas a yaw angle of `90` degrees yields a face whose eye line is
// perpendicular to that of the camera.
// 
// You must check the value of the [HasYawAngle] property before accessing
// this property. If the value in the [HasYawAngle] property is [false],
// reading the value in this property raises an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/yawAngle
func (m AVMetadataFaceObject) YawAngle() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("yawAngle"))
	return rv
}

