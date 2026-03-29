// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
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
	InitWithFormatDescription(formatDescription uintptr) AVSpatialVideoConfiguration

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

	// A video composition tool to use with Core Animation in offline rendering.
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	SetAnimationTool(value IAVVideoCompositionCoreAnimationTool)
	// The color primaries used for video composition.
	ColorPrimaries() string
	SetColorPrimaries(value string)
	// The transfer function used for video composition.
	ColorTransferFunction() string
	SetColorTransferFunction(value string)
	// The YCbCr matrix used for video composition.
	ColorYCbCrMatrix() string
	SetColorYCbCrMatrix(value string)
	// A custom compositor class to use.
	CustomVideoCompositorClass() AVVideoCompositing
	SetCustomVideoCompositorClass(value AVVideoCompositing)
	// A time interval for which the video composition should render composed video frames.
	FrameDuration() coremedia.CMTime
	SetFrameDuration(value coremedia.CMTime)
	// The scale at which the video composition should render.
	RenderScale() float32
	SetRenderScale(value float32)
	// The size at which the video composition should render.
	RenderSize() corefoundation.CGSize
	SetRenderSize(value corefoundation.CGSize)
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
// An instance of AVSpatialVideoConfiguration
//
// # Discussion
// 
// The format description is not stored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func NewSpatialVideoConfigurationWithFormatDescription(formatDescription uintptr) AVSpatialVideoConfiguration {
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
// An instance of AVSpatialVideoConfiguration
//
// # Discussion
// 
// The format description is not stored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func (s AVSpatialVideoConfiguration) InitWithFormatDescription(formatDescription uintptr) AVSpatialVideoConfiguration {
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
// A video composition tool to use with Core Animation in offline rendering.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (s AVSpatialVideoConfiguration) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animationTool"))
	return AVVideoCompositionCoreAnimationToolFromID(objc.ID(rv))
}
func (s AVSpatialVideoConfiguration) SetAnimationTool(value IAVVideoCompositionCoreAnimationTool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAnimationTool:"), value)
}
// The color primaries used for video composition.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (s AVSpatialVideoConfiguration) ColorPrimaries() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("colorPrimaries"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpatialVideoConfiguration) SetColorPrimaries(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorPrimaries:"), objc.String(value))
}
// The transfer function used for video composition.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (s AVSpatialVideoConfiguration) ColorTransferFunction() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("colorTransferFunction"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpatialVideoConfiguration) SetColorTransferFunction(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorTransferFunction:"), objc.String(value))
}
// The YCbCr matrix used for video composition.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (s AVSpatialVideoConfiguration) ColorYCbCrMatrix() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("colorYCbCrMatrix"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpatialVideoConfiguration) SetColorYCbCrMatrix(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorYCbCrMatrix:"), objc.String(value))
}
// A custom compositor class to use.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/customvideocompositorclass
func (s AVSpatialVideoConfiguration) CustomVideoCompositorClass() AVVideoCompositing {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("customVideoCompositorClass"))
	return AVVideoCompositingObjectFromID(rv)
}
func (s AVSpatialVideoConfiguration) SetCustomVideoCompositorClass(value AVVideoCompositing) {
	objc.Send[struct{}](s.ID, objc.Sel("setCustomVideoCompositorClass:"), value)
}
// A time interval for which the video composition should render composed
// video frames.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/frameduration
func (s AVSpatialVideoConfiguration) FrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("frameDuration"))
	return coremedia.CMTime(rv)
}
func (s AVSpatialVideoConfiguration) SetFrameDuration(value coremedia.CMTime) {
	objc.Send[struct{}](s.ID, objc.Sel("setFrameDuration:"), value)
}
// The scale at which the video composition should render.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/renderscale
func (s AVSpatialVideoConfiguration) RenderScale() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("renderScale"))
	return rv
}
func (s AVSpatialVideoConfiguration) SetRenderScale(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setRenderScale:"), value)
}
// The size at which the video composition should render.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (s AVSpatialVideoConfiguration) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("renderSize"))
	return corefoundation.CGSize(rv)
}
func (s AVSpatialVideoConfiguration) SetRenderSize(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setRenderSize:"), value)
}

