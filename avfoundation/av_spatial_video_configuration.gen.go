// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpatialVideoConfiguration] class.
var (
	_AVSpatialVideoConfigurationClass     AVSpatialVideoConfigurationClass
	_AVSpatialVideoConfigurationClassOnce sync.Once
)

func getAVSpatialVideoConfigurationClass() AVSpatialVideoConfigurationClass {
	_AVSpatialVideoConfigurationClassOnce.Do(func() {
		_AVSpatialVideoConfigurationClass = AVSpatialVideoConfigurationClass{class: objc.GetClass("AVSpatialVideoConfiguration")}
	})
	return _AVSpatialVideoConfigurationClass
}

// GetAVSpatialVideoConfigurationClass returns the class object for AVSpatialVideoConfiguration.
func GetAVSpatialVideoConfigurationClass() AVSpatialVideoConfigurationClass {
	return getAVSpatialVideoConfigurationClass()
}

type AVSpatialVideoConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpatialVideoConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpatialVideoConfigurationClass) Alloc() AVSpatialVideoConfiguration {
	rv := objc.Send[AVSpatialVideoConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An AVSpatialVideoConfiguration specifies spatial video properties.
//
// # Creating a configuration
//
//   - [AVSpatialVideoConfiguration.InitWithFormatDescription]: Initializes an AVSpatialVideoConfiguration with a format description.
//
// # Modifying the configuration
//
//   - [AVSpatialVideoConfiguration.CameraCalibrationDataLensCollection]: Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//   - [AVSpatialVideoConfiguration.SetCameraCalibrationDataLensCollection]
//   - [AVSpatialVideoConfiguration.CameraSystemBaseline]: Specifies the distance between centers of the lenses of the camera system that created the video.
//   - [AVSpatialVideoConfiguration.SetCameraSystemBaseline]
//   - [AVSpatialVideoConfiguration.DisparityAdjustment]: Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//   - [AVSpatialVideoConfiguration.SetDisparityAdjustment]
//   - [AVSpatialVideoConfiguration.HorizontalFieldOfView]: Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//   - [AVSpatialVideoConfiguration.SetHorizontalFieldOfView]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class
type AVSpatialVideoConfiguration struct {
	objectivec.Object
}

// AVSpatialVideoConfigurationFromID constructs a [AVSpatialVideoConfiguration] from an objc.ID.
//
// An AVSpatialVideoConfiguration specifies spatial video properties.
func AVSpatialVideoConfigurationFromID(id objc.ID) AVSpatialVideoConfiguration {
	return AVSpatialVideoConfiguration{objectivec.Object{ID: id}}
}

// Ensure AVSpatialVideoConfiguration implements IAVSpatialVideoConfiguration.
var _ IAVSpatialVideoConfiguration = AVSpatialVideoConfiguration{}

// An interface definition for the [AVSpatialVideoConfiguration] class.
//
// # Creating a configuration
//
//   - [IAVSpatialVideoConfiguration.InitWithFormatDescription]: Initializes an AVSpatialVideoConfiguration with a format description.
//
// # Modifying the configuration
//
//   - [IAVSpatialVideoConfiguration.CameraCalibrationDataLensCollection]: Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//   - [IAVSpatialVideoConfiguration.SetCameraCalibrationDataLensCollection]
//   - [IAVSpatialVideoConfiguration.CameraSystemBaseline]: Specifies the distance between centers of the lenses of the camera system that created the video.
//   - [IAVSpatialVideoConfiguration.SetCameraSystemBaseline]
//   - [IAVSpatialVideoConfiguration.DisparityAdjustment]: Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//   - [IAVSpatialVideoConfiguration.SetDisparityAdjustment]
//   - [IAVSpatialVideoConfiguration.HorizontalFieldOfView]: Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//   - [IAVSpatialVideoConfiguration.SetHorizontalFieldOfView]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class
type IAVSpatialVideoConfiguration interface {
	objectivec.IObject

	// Topic: Creating a configuration

	// Initializes an AVSpatialVideoConfiguration with a format description.
	InitWithFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVSpatialVideoConfiguration

	// Topic: Modifying the configuration

	// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
	CameraCalibrationDataLensCollection() foundation.INSDictionary
	SetCameraCalibrationDataLensCollection(value foundation.INSDictionary)
	// Specifies the distance between centers of the lenses of the camera system that created the video.
	CameraSystemBaseline() foundation.NSNumber
	SetCameraSystemBaseline(value foundation.NSNumber)
	// Specifies a relative shift of the left and right images, which changes the zero parallax plane.
	DisparityAdjustment() foundation.NSNumber
	SetDisparityAdjustment(value foundation.NSNumber)
	// Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
	HorizontalFieldOfView() foundation.NSNumber
	SetHorizontalFieldOfView(value foundation.NSNumber)
}

// Init initializes the instance.
func (s AVSpatialVideoConfiguration) Init() AVSpatialVideoConfiguration {
	rv := objc.Send[AVSpatialVideoConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpatialVideoConfiguration) Autorelease() AVSpatialVideoConfiguration {
	rv := objc.Send[AVSpatialVideoConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpatialVideoConfiguration creates a new AVSpatialVideoConfiguration instance.
func NewAVSpatialVideoConfiguration() AVSpatialVideoConfiguration {
	class := getAVSpatialVideoConfigurationClass()
	rv := objc.Send[AVSpatialVideoConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an AVSpatialVideoConfiguration with a format description.
//
// formatDescription: Format description to use to initialize the AVSpatialVideoConfiguration.
//
// # Return Value
//
// # An instance of AVSpatialVideoConfiguration
//
// # Discussion
//
// The format description is not stored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func NewSpatialVideoConfigurationWithFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVSpatialVideoConfiguration {
	instance := getAVSpatialVideoConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormatDescription:"), formatDescription)
	return AVSpatialVideoConfigurationFromID(rv)
}

// Initializes an AVSpatialVideoConfiguration with a format description.
//
// formatDescription: Format description to use to initialize the AVSpatialVideoConfiguration.
//
// # Return Value
//
// # An instance of AVSpatialVideoConfiguration
//
// # Discussion
//
// The format description is not stored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func (s AVSpatialVideoConfiguration) InitWithFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVSpatialVideoConfiguration {
	rv := objc.Send[AVSpatialVideoConfiguration](s.ID, objc.Sel("initWithFormatDescription:"), formatDescription)
	return rv
}

// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//
// # Discussion
//
// The property value is an array of dictionaries describing the camera
// calibration data for each lens. The camera calibration data includes
// intrinsics and extrinics with other parameters. This property is only
// applicable when the projection kind is
// kCMTagProjectionTypeParametricImmersive. Can be nil if the value is
// unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraCalibrationDataLensCollection
func (s AVSpatialVideoConfiguration) CameraCalibrationDataLensCollection() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("cameraCalibrationDataLensCollection"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s AVSpatialVideoConfiguration) SetCameraCalibrationDataLensCollection(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setCameraCalibrationDataLensCollection:"), value)
}

// Specifies the distance between centers of the lenses of the camera system
// that created the video.
//
// # Discussion
//
// The distance is in micrometers or thousandths of a millimeter. Can be nil
// if the value is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraSystemBaseline
func (s AVSpatialVideoConfiguration) CameraSystemBaseline() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("cameraSystemBaseline"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s AVSpatialVideoConfiguration) SetCameraSystemBaseline(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setCameraSystemBaseline:"), value)
}

// Specifies a relative shift of the left and right images, which changes the
// zero parallax plane.
//
// # Discussion
//
// The value is in normalized image space and measured over the range of
// -10000 to 10000 mapping to the uniform range [-1.0…1.0]. The interval of
// 0.0 to 1.0 or 0 to 10000 maps onto the stereo eye view image width. The
// negative interval 0.0 to -1.0 or 0 to -10000 similarly map onto the stereo
// eye view image width. Can be nil if the value is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/disparityAdjustment
func (s AVSpatialVideoConfiguration) DisparityAdjustment() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("disparityAdjustment"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s AVSpatialVideoConfiguration) SetDisparityAdjustment(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisparityAdjustment:"), value)
}

// Specifies horizontal field of view in thousandths of a degree. Can be nil
// if the value is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/horizontalFieldOfView
func (s AVSpatialVideoConfiguration) HorizontalFieldOfView() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("horizontalFieldOfView"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s AVSpatialVideoConfiguration) SetHorizontalFieldOfView(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalFieldOfView:"), value)
}
