// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCameraCalibrationData] class.
var (
	_AVCameraCalibrationDataClass     AVCameraCalibrationDataClass
	_AVCameraCalibrationDataClassOnce sync.Once
)

func getAVCameraCalibrationDataClass() AVCameraCalibrationDataClass {
	_AVCameraCalibrationDataClassOnce.Do(func() {
		_AVCameraCalibrationDataClass = AVCameraCalibrationDataClass{class: objc.GetClass("AVCameraCalibrationData")}
	})
	return _AVCameraCalibrationDataClass
}

// GetAVCameraCalibrationDataClass returns the class object for AVCameraCalibrationData.
func GetAVCameraCalibrationDataClass() AVCameraCalibrationDataClass {
	return getAVCameraCalibrationDataClass()
}

type AVCameraCalibrationDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCameraCalibrationDataClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCameraCalibrationDataClass) Alloc() AVCameraCalibrationData {
	rv := objc.Send[AVCameraCalibrationData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Information about the camera characteristics used to capture images and
// depth data.
//
// # Overview
//
// Information about the calibration of a camera—such as its pixel focal
// length, principal point, and lens distortion characteristics—helps to
// determine the geometric relationships between the camera device and the
// images it captures. You can use this information to accurately render
// visual effects into images produced by a camera or perform computer vision
// tasks such as correcting images for geometric distortions.
//
// # Mapping pixels to scene geometry
//
//   - [AVCameraCalibrationData.IntrinsicMatrix]: A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.
//   - [AVCameraCalibrationData.IntrinsicMatrixReferenceDimensions]: The image dimensions to which the camera’s intrinsic matrix values are relative.
//   - [AVCameraCalibrationData.ExtrinsicMatrix]: A matrix relating a camera’s position and orientation to a world or scene coordinate system.
//   - [AVCameraCalibrationData.PixelSize]: The size, in millimeters, of one image pixel.
//
// # Correcting for lens distortion
//
//   - [AVCameraCalibrationData.LensDistortionLookupTable]: A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.
//   - [AVCameraCalibrationData.InverseLensDistortionLookupTable]: A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.
//   - [AVCameraCalibrationData.LensDistortionCenter]: The offset of the distortion center of the camera lens from the top-left corner of the image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData
type AVCameraCalibrationData struct {
	objectivec.Object
}

// AVCameraCalibrationDataFromID constructs a [AVCameraCalibrationData] from an objc.ID.
//
// Information about the camera characteristics used to capture images and
// depth data.
func AVCameraCalibrationDataFromID(id objc.ID) AVCameraCalibrationData {
	return AVCameraCalibrationData{objectivec.Object{ID: id}}
}

// NOTE: AVCameraCalibrationData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCameraCalibrationData] class.
//
// # Mapping pixels to scene geometry
//
//   - [IAVCameraCalibrationData.IntrinsicMatrix]: A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.
//   - [IAVCameraCalibrationData.IntrinsicMatrixReferenceDimensions]: The image dimensions to which the camera’s intrinsic matrix values are relative.
//   - [IAVCameraCalibrationData.ExtrinsicMatrix]: A matrix relating a camera’s position and orientation to a world or scene coordinate system.
//   - [IAVCameraCalibrationData.PixelSize]: The size, in millimeters, of one image pixel.
//
// # Correcting for lens distortion
//
//   - [IAVCameraCalibrationData.LensDistortionLookupTable]: A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.
//   - [IAVCameraCalibrationData.InverseLensDistortionLookupTable]: A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.
//   - [IAVCameraCalibrationData.LensDistortionCenter]: The offset of the distortion center of the camera lens from the top-left corner of the image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData
type IAVCameraCalibrationData interface {
	objectivec.IObject

	// Topic: Mapping pixels to scene geometry

	// A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.
	IntrinsicMatrix() unsafe.Pointer
	// The image dimensions to which the camera’s intrinsic matrix values are relative.
	IntrinsicMatrixReferenceDimensions() corefoundation.CGSize
	// A matrix relating a camera’s position and orientation to a world or scene coordinate system.
	ExtrinsicMatrix() unsafe.Pointer
	// The size, in millimeters, of one image pixel.
	PixelSize() float32

	// Topic: Correcting for lens distortion

	// A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.
	LensDistortionLookupTable() foundation.INSData
	// A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.
	InverseLensDistortionLookupTable() foundation.INSData
	// The offset of the distortion center of the camera lens from the top-left corner of the image.
	LensDistortionCenter() corefoundation.CGPoint
}

// Init initializes the instance.
func (c AVCameraCalibrationData) Init() AVCameraCalibrationData {
	rv := objc.Send[AVCameraCalibrationData](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCameraCalibrationData) Autorelease() AVCameraCalibrationData {
	rv := objc.Send[AVCameraCalibrationData](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCameraCalibrationData creates a new AVCameraCalibrationData instance.
func NewAVCameraCalibrationData() AVCameraCalibrationData {
	class := getAVCameraCalibrationDataClass()
	rv := objc.Send[AVCameraCalibrationData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A matrix that relates a camera’s internal properties to an ideal
// pinhole-camera model.
//
// # Discussion
//
// The intrinsic matrix allows you to transform 3D coordinates to 2D
// coordinates on an image plane using the pinhole camera model. Equations
// like the following commonly represent the intrinsic matrix as `K:`
//
// [media-2902623]
//
// The equation expresses all values in pixels. The values `fx` and `fy` are
// the pixel focal length, and are identical for square pixels. The `ox` and
// `oy` values are the offsets of the principal point, from the top-left
// corner of the image frame. The principal point is relative to the top-left
// corner of the top-left pixel. Each pixel value represents a sampled value
// from the center of that pixel.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/intrinsicMatrix
func (c AVCameraCalibrationData) IntrinsicMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("intrinsicMatrix"))
	return rv
}

// The image dimensions to which the camera’s intrinsic matrix values are
// relative.
//
// # Discussion
//
// The [IntrinsicMatrix] property measures focal length and principal point
// offset in pixels, but those values are meaningful only in the context of an
// image of this size.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/intrinsicMatrixReferenceDimensions
func (c AVCameraCalibrationData) IntrinsicMatrixReferenceDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("intrinsicMatrixReferenceDimensions"))
	return corefoundation.CGSize(rv)
}

// A matrix relating a camera’s position and orientation to a world or scene
// coordinate system.
//
// # Discussion
//
// The extrinsic matrix consists of a unitless 3 x 3 rotation matrix ([R]) on
// the left and a 3 x 1 column vector translation (`t`) on the right. The
// translation vector’s units are millimeters.
//
// [media-2902624]
//
// The camera’s pose is expressed with respect to a reference camera
// (camera-to-world view). If the rotation matrix is an identity matrix, then
// this camera is the reference camera.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/extrinsicMatrix
func (c AVCameraCalibrationData) ExtrinsicMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("extrinsicMatrix"))
	return rv
}

// The size, in millimeters, of one image pixel.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/pixelSize
func (c AVCameraCalibrationData) PixelSize() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pixelSize"))
	return rv
}

// A map of floating-point values describing radial distortions imparted by
// the camera lens, for use in rectifying camera images.
//
// # Discussion
//
// Images captured by a camera are geometrically warped by small imperfections
// in the lens. To project from the 2D image plane back into the 3D world, the
// images must be distortion corrected, or made rectilinear. Lens distortion
// is modeled using a one-dimensional lookup table of 32-bit float values
// evenly distributed along a radius from the center of the distortion to a
// corner, with each value representing a magnification of the radius. This
// model assumes symmetrical lens distortion.
//
// When dealing with [AVDepthData] objects, the disparity/depth map
// representations are geometrically distorted to align with images produced
// by the camera.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/lensDistortionLookupTable
func (c AVCameraCalibrationData) LensDistortionLookupTable() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lensDistortionLookupTable"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A map of floating-point values describing radial distortions for use in
// reapplying camera geometry to a rectified image.
//
// # Discussion
//
// If you’ve rectified an image by removing the distortions characterized by
// the [LensDistortionLookupTable] property, and now wish to go back to a
// geometrically distorted image (for example, to render visual effects into
// the camera image or perform computer vision tasks such as scene
// reconstruction), use this inverse lookup table.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/inverseLensDistortionLookupTable
func (c AVCameraCalibrationData) InverseLensDistortionLookupTable() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inverseLensDistortionLookupTable"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The offset of the distortion center of the camera lens from the top-left
// corner of the image.
//
// # Discussion
//
// Due to geometric distortions in the image, the center of the distortion may
// not be equal to the optical center (principal point) of the lens. When
// making an image rectilinear, use the distortion center rather than the
// optical center of the image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/lensDistortionCenter
func (c AVCameraCalibrationData) LensDistortionCenter() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("lensDistortionCenter"))
	return corefoundation.CGPoint(rv)
}
