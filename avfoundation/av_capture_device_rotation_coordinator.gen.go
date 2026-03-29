// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVCaptureDeviceRotationCoordinator] class.
var (
	_AVCaptureDeviceRotationCoordinatorClass     AVCaptureDeviceRotationCoordinatorClass
	_AVCaptureDeviceRotationCoordinatorClassOnce sync.Once
)

func getAVCaptureDeviceRotationCoordinatorClass() AVCaptureDeviceRotationCoordinatorClass {
	_AVCaptureDeviceRotationCoordinatorClassOnce.Do(func() {
		_AVCaptureDeviceRotationCoordinatorClass = AVCaptureDeviceRotationCoordinatorClass{class: objc.GetClass("AVCaptureDeviceRotationCoordinator")}
	})
	return _AVCaptureDeviceRotationCoordinatorClass
}

// GetAVCaptureDeviceRotationCoordinatorClass returns the class object for AVCaptureDeviceRotationCoordinator.
func GetAVCaptureDeviceRotationCoordinatorClass() AVCaptureDeviceRotationCoordinatorClass {
	return getAVCaptureDeviceRotationCoordinatorClass()
}

type AVCaptureDeviceRotationCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureDeviceRotationCoordinatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceRotationCoordinatorClass) Alloc() AVCaptureDeviceRotationCoordinator {
	rv := objc.Send[AVCaptureDeviceRotationCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that monitors the physical orientation of a capture device and
// provides adjustment angles to keep images level, relative to gravity.
//
// # Overview
// 
// Correctly rotate the photos and movies your app captures, and optionally, a
// live camera preview, by applying a coordinator’s
// [AVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelCapture] and
// [AVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelPreview] properties, respectively. Each
// rotation coordinator instance updates its properties so that your app can
// observe them and immediately apply them to the relevant components.
//
// # Creating a rotation coordinator
//
//   - [AVCaptureDeviceRotationCoordinator.InitWithDevicePreviewLayer]: Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
//
// # Compensating for a device’s rotation
//
//   - [AVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelCapture]: An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
//   - [AVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelPreview]: An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
//
// # Inspecting a coordinator’s configuration
//
//   - [AVCaptureDeviceRotationCoordinator.Device]: The capture device the coordinator monitors to track its physical rotation.
//   - [AVCaptureDeviceRotationCoordinator.PreviewLayer]: The layer that displays a camera preview the coordinator calculates a video rotation angle for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator
type AVCaptureDeviceRotationCoordinator struct {
	objectivec.Object
}

// AVCaptureDeviceRotationCoordinatorFromID constructs a [AVCaptureDeviceRotationCoordinator] from an objc.ID.
//
// A class that monitors the physical orientation of a capture device and
// provides adjustment angles to keep images level, relative to gravity.
func AVCaptureDeviceRotationCoordinatorFromID(id objc.ID) AVCaptureDeviceRotationCoordinator {
	return AVCaptureDeviceRotationCoordinator{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeviceRotationCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeviceRotationCoordinator] class.
//
// # Creating a rotation coordinator
//
//   - [IAVCaptureDeviceRotationCoordinator.InitWithDevicePreviewLayer]: Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
//
// # Compensating for a device’s rotation
//
//   - [IAVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelCapture]: An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
//   - [IAVCaptureDeviceRotationCoordinator.VideoRotationAngleForHorizonLevelPreview]: An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
//
// # Inspecting a coordinator’s configuration
//
//   - [IAVCaptureDeviceRotationCoordinator.Device]: The capture device the coordinator monitors to track its physical rotation.
//   - [IAVCaptureDeviceRotationCoordinator.PreviewLayer]: The layer that displays a camera preview the coordinator calculates a video rotation angle for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator
type IAVCaptureDeviceRotationCoordinator interface {
	objectivec.IObject

	// Topic: Creating a rotation coordinator

	// Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
	InitWithDevicePreviewLayer(device IAVCaptureDevice, previewLayer quartzcore.CALayer) AVCaptureDeviceRotationCoordinator

	// Topic: Compensating for a device’s rotation

	// An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
	VideoRotationAngleForHorizonLevelCapture() float64
	// An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
	VideoRotationAngleForHorizonLevelPreview() float64

	// Topic: Inspecting a coordinator’s configuration

	// The capture device the coordinator monitors to track its physical rotation.
	Device() IAVCaptureDevice
	// The layer that displays a camera preview the coordinator calculates a video rotation angle for.
	PreviewLayer() quartzcore.CALayer
}

// Init initializes the instance.
func (c AVCaptureDeviceRotationCoordinator) Init() AVCaptureDeviceRotationCoordinator {
	rv := objc.Send[AVCaptureDeviceRotationCoordinator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeviceRotationCoordinator) Autorelease() AVCaptureDeviceRotationCoordinator {
	rv := objc.Send[AVCaptureDeviceRotationCoordinator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceRotationCoordinator creates a new AVCaptureDeviceRotationCoordinator instance.
func NewAVCaptureDeviceRotationCoordinator() AVCaptureDeviceRotationCoordinator {
	class := getAVCaptureDeviceRotationCoordinatorClass()
	rv := objc.Send[AVCaptureDeviceRotationCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a coordinator that provides separate compensation angles for
// content your app takes with a capture device, and for your app’s camera
// preview.
//
// device: A capture device the new coordinator monitors to track its physical
// rotation to calculate its [VideoRotationAngleForHorizonLevelCapture]
// property.
//
// previewLayer: A layer that displays a camera preview the new coordinator monitors to
// calculate its [VideoRotationAngleForHorizonLevelPreview] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/init(device:previewLayer:)
func NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer(device IAVCaptureDevice, previewLayer quartzcore.CALayer) AVCaptureDeviceRotationCoordinator {
	instance := getAVCaptureDeviceRotationCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:previewLayer:"), device, previewLayer)
	return AVCaptureDeviceRotationCoordinatorFromID(rv)
}

// Creates a coordinator that provides separate compensation angles for
// content your app takes with a capture device, and for your app’s camera
// preview.
//
// device: A capture device the new coordinator monitors to track its physical
// rotation to calculate its [VideoRotationAngleForHorizonLevelCapture]
// property.
//
// previewLayer: A layer that displays a camera preview the new coordinator monitors to
// calculate its [VideoRotationAngleForHorizonLevelPreview] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/init(device:previewLayer:)
func (c AVCaptureDeviceRotationCoordinator) InitWithDevicePreviewLayer(device IAVCaptureDevice, previewLayer quartzcore.CALayer) AVCaptureDeviceRotationCoordinator {
	rv := objc.Send[AVCaptureDeviceRotationCoordinator](c.ID, objc.Sel("initWithDevice:previewLayer:"), device, previewLayer)
	return rv
}

// An angle the coordinator provides your app to apply to photos or videos it
// captures with the device so that they’re level relative to gravity.
//
// # Discussion
// 
// Your app can get immediate rotation angle updates from the rotation
// coordinator with key-value observation (KVO) of this property. The system
// calls key-value observation code on the main queue so that it has the same
// behavior as the [VideoRotationAngleForHorizonLevelPreview] property.
// 
// Apps typically apply the property’s value to an [AVCaptureConnection]
// instance’s [VideoRotationAngle] property, such as saving photos and
// videos with the correction rotation with [AVCapturePhotoOutput] and
// [AVCaptureMovieFileOutput], respectively.
// 
// Alternatively, if your app uses an [AVCaptureVideoDataOutput] instance with
// an [AVAssetWriter], such as for recording custom videos with effects,
// don’t rotate the video with [AVCaptureConnection]. Instead, set the
// rotation with an [AVAssetWriterInput] instance’s [Transform] property,
// which alters the output file’s metadata. With this approach,
// video-playing apps apply the rotation during playback, which uses less
// energy than rotating each frame with the capture connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/videoRotationAngleForHorizonLevelCapture
func (c AVCaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelCapture() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoRotationAngleForHorizonLevelCapture"))
	return rv
}
// An angle the coordinator provides your app to apply to the preview layer so
// that it’s level relative to gravity.
//
// # Discussion
// 
// Your app can get immediate rotation angle updates from the rotation
// coordinator with key-value observation (KVO) for this property. You can
// immediately update your app’s UI from its key-value observation code
// because the rotation coordinator notifies your app on the main queue.
// 
// Apps typically apply the property’s value to an [AVCaptureConnection]
// instance’s [VideoRotationAngle] property, such as displaying a camera
// preview with the correction for an [AVCaptureVideoPreviewLayer] instance.
// 
// Alternatively, if your app uses an [AVCaptureVideoDataOutput] instance to
// display a custom camera preview, such as with effects, don’t rotate the
// video with [AVCaptureConnection]. Instead, set the rotation in your
// [CALayer] instance’s [transform] property, such as with an
// [AVSampleBufferDisplayLayer] instance. This approach uses less energy than
// rotating each frame with the capture connection.
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [transform]: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/videoRotationAngleForHorizonLevelPreview
func (c AVCaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelPreview() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoRotationAngleForHorizonLevelPreview"))
	return rv
}
// The capture device the coordinator monitors to track its physical rotation.
//
// # Discussion
// 
// The coordinator updates its [VideoRotationAngleForHorizonLevelCapture]
// property by monitoring the device’s physical rotation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/device
func (c AVCaptureDeviceRotationCoordinator) Device() IAVCaptureDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("device"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}
// The layer that displays a camera preview the coordinator calculates a video
// rotation angle for.
//
// # Discussion
// 
// The coordinator updates its [VideoRotationAngleForHorizonLevelPreview]
// property by monitoring the layer and the physical rotation of [Device].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/previewLayer
func (c AVCaptureDeviceRotationCoordinator) PreviewLayer() quartzcore.CALayer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previewLayer"))
	return quartzcore.CALayerFromID(objc.ID(rv))
}

