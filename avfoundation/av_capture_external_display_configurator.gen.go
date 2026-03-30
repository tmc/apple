// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVCaptureExternalDisplayConfigurator] class.
var (
	_AVCaptureExternalDisplayConfiguratorClass     AVCaptureExternalDisplayConfiguratorClass
	_AVCaptureExternalDisplayConfiguratorClassOnce sync.Once
)

func getAVCaptureExternalDisplayConfiguratorClass() AVCaptureExternalDisplayConfiguratorClass {
	_AVCaptureExternalDisplayConfiguratorClassOnce.Do(func() {
		_AVCaptureExternalDisplayConfiguratorClass = AVCaptureExternalDisplayConfiguratorClass{class: objc.GetClass("AVCaptureExternalDisplayConfigurator")}
	})
	return _AVCaptureExternalDisplayConfiguratorClass
}

// GetAVCaptureExternalDisplayConfiguratorClass returns the class object for AVCaptureExternalDisplayConfigurator.
func GetAVCaptureExternalDisplayConfiguratorClass() AVCaptureExternalDisplayConfiguratorClass {
	return getAVCaptureExternalDisplayConfiguratorClass()
}

type AVCaptureExternalDisplayConfiguratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureExternalDisplayConfiguratorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureExternalDisplayConfiguratorClass) Alloc() AVCaptureExternalDisplayConfigurator {
	rv := objc.Send[AVCaptureExternalDisplayConfigurator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A configurator class allowing you to configure properties of an external
// display to match the camera’s active video format.
//
// # Overview
//
// An [AVCaptureExternalDisplayConfigurator] allows you to configure a
// connected external display to output a clean feed using a [CALayer]. Using
// the configurator, you can opt into automatic adjustment of the external
// display’s color space and / or frame rate to match your device’s
// capture configuration. These adjustments are only applied to the external
// display, not to the device.
//
// # Creating an external display configurator
//
//   - [AVCaptureExternalDisplayConfigurator.InitWithDevicePreviewLayerConfiguration]: An external display configurator instance that attempts to synchronize the preview layer configuration with the device capture configuration.
//
// # Inspecting the configurator
//
//   - [AVCaptureExternalDisplayConfigurator.ActiveExternalDisplayFrameRate]: The currently configured frame rate on the external display that’s displaying the preview layer.
//   - [AVCaptureExternalDisplayConfigurator.Device]: The device for which the coordinator configures the preview layer.
//   - [AVCaptureExternalDisplayConfigurator.Active]: This property tells you whether the configurator is actively configuring the external display.
//   - [AVCaptureExternalDisplayConfigurator.PreviewLayer]: The layer for which the configurator adjusts display properties to match the device’s state.
//
// # Stopping configuration
//
//   - [AVCaptureExternalDisplayConfigurator.Stop]: Forces the external display configurator to asynchronously stop configuring the external display.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator
type AVCaptureExternalDisplayConfigurator struct {
	objectivec.Object
}

// AVCaptureExternalDisplayConfiguratorFromID constructs a [AVCaptureExternalDisplayConfigurator] from an objc.ID.
//
// A configurator class allowing you to configure properties of an external
// display to match the camera’s active video format.
func AVCaptureExternalDisplayConfiguratorFromID(id objc.ID) AVCaptureExternalDisplayConfigurator {
	return AVCaptureExternalDisplayConfigurator{objectivec.Object{ID: id}}
}

// NOTE: AVCaptureExternalDisplayConfigurator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureExternalDisplayConfigurator] class.
//
// # Creating an external display configurator
//
//   - [IAVCaptureExternalDisplayConfigurator.InitWithDevicePreviewLayerConfiguration]: An external display configurator instance that attempts to synchronize the preview layer configuration with the device capture configuration.
//
// # Inspecting the configurator
//
//   - [IAVCaptureExternalDisplayConfigurator.ActiveExternalDisplayFrameRate]: The currently configured frame rate on the external display that’s displaying the preview layer.
//   - [IAVCaptureExternalDisplayConfigurator.Device]: The device for which the coordinator configures the preview layer.
//   - [IAVCaptureExternalDisplayConfigurator.Active]: This property tells you whether the configurator is actively configuring the external display.
//   - [IAVCaptureExternalDisplayConfigurator.PreviewLayer]: The layer for which the configurator adjusts display properties to match the device’s state.
//
// # Stopping configuration
//
//   - [IAVCaptureExternalDisplayConfigurator.Stop]: Forces the external display configurator to asynchronously stop configuring the external display.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator
type IAVCaptureExternalDisplayConfigurator interface {
	objectivec.IObject

	// Topic: Creating an external display configurator

	// An external display configurator instance that attempts to synchronize the preview layer configuration with the device capture configuration.
	InitWithDevicePreviewLayerConfiguration(device IAVCaptureDevice, previewLayer quartzcore.CALayer, configuration IAVCaptureExternalDisplayConfiguration) AVCaptureExternalDisplayConfigurator

	// Topic: Inspecting the configurator

	// The currently configured frame rate on the external display that’s displaying the preview layer.
	ActiveExternalDisplayFrameRate() float64
	// The device for which the coordinator configures the preview layer.
	Device() IAVCaptureDevice
	// This property tells you whether the configurator is actively configuring the external display.
	Active() bool
	// The layer for which the configurator adjusts display properties to match the device’s state.
	PreviewLayer() quartzcore.CALayer

	// Topic: Stopping configuration

	// Forces the external display configurator to asynchronously stop configuring the external display.
	Stop()

	// The capture format in use by the device.
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
}

// Init initializes the instance.
func (c AVCaptureExternalDisplayConfigurator) Init() AVCaptureExternalDisplayConfigurator {
	rv := objc.Send[AVCaptureExternalDisplayConfigurator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureExternalDisplayConfigurator) Autorelease() AVCaptureExternalDisplayConfigurator {
	rv := objc.Send[AVCaptureExternalDisplayConfigurator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureExternalDisplayConfigurator creates a new AVCaptureExternalDisplayConfigurator instance.
func NewAVCaptureExternalDisplayConfigurator() AVCaptureExternalDisplayConfigurator {
	class := getAVCaptureExternalDisplayConfiguratorClass()
	rv := objc.Send[AVCaptureExternalDisplayConfigurator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An external display configurator instance that attempts to synchronize the
// preview layer configuration with the device capture configuration.
//
// device: The device for which to monitor the configuration.
//
// previewLayer: The layer that is being used on an external display for displaying the
// camera preview.
//
// configuration: A configuration specifying which aspects of the camera’s active format to
// monitor and configure on the external display.
//
// # Return Value
//
// An [AVCaptureExternalDisplayConfigurator] instance.
//
// # Discussion
//
// An [AVCaptureExternalDisplayConfigurator] is only applicable to external
// displays. It determines which properties to configure on the external
// display based on your provided configuration (see
// [AVCaptureExternalDisplayConfiguration]). The configurator observes changes
// to your camera’‘s configuration, and when changes are observed, it
// modifies the external display’s properties to match.
//
// If multiple configurators are linked to the same external display ,the last
// one created becomes the active configurator for the external display (see
// [Active]).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/init(device:previewLayer:configuration:)
func NewCaptureExternalDisplayConfiguratorWithDevicePreviewLayerConfiguration(device IAVCaptureDevice, previewLayer quartzcore.CALayer, configuration IAVCaptureExternalDisplayConfiguration) AVCaptureExternalDisplayConfigurator {
	instance := getAVCaptureExternalDisplayConfiguratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:previewLayer:configuration:"), device, previewLayer, configuration)
	return AVCaptureExternalDisplayConfiguratorFromID(rv)
}

// An external display configurator instance that attempts to synchronize the
// preview layer configuration with the device capture configuration.
//
// device: The device for which to monitor the configuration.
//
// previewLayer: The layer that is being used on an external display for displaying the
// camera preview.
//
// configuration: A configuration specifying which aspects of the camera’s active format to
// monitor and configure on the external display.
//
// # Return Value
//
// An [AVCaptureExternalDisplayConfigurator] instance.
//
// # Discussion
//
// An [AVCaptureExternalDisplayConfigurator] is only applicable to external
// displays. It determines which properties to configure on the external
// display based on your provided configuration (see
// [AVCaptureExternalDisplayConfiguration]). The configurator observes changes
// to your camera’‘s configuration, and when changes are observed, it
// modifies the external display’s properties to match.
//
// If multiple configurators are linked to the same external display ,the last
// one created becomes the active configurator for the external display (see
// [Active]).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/init(device:previewLayer:configuration:)
func (c AVCaptureExternalDisplayConfigurator) InitWithDevicePreviewLayerConfiguration(device IAVCaptureDevice, previewLayer quartzcore.CALayer, configuration IAVCaptureExternalDisplayConfiguration) AVCaptureExternalDisplayConfigurator {
	rv := objc.Send[AVCaptureExternalDisplayConfigurator](c.ID, objc.Sel("initWithDevice:previewLayer:configuration:"), device, previewLayer, configuration)
	return rv
}

// Forces the external display configurator to asynchronously stop configuring
// the external display.
//
// # Discussion
//
// Call [Stop] to force the [AVCaptureExternalDisplayConfigurator] to
// asynchronously stop configuring the external display. Once stopped, the
// [Active] property changes to `false` and the
// [ActiveExternalDisplayFrameRate] becomes 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/stop()
func (c AVCaptureExternalDisplayConfigurator) Stop() {
	objc.Send[objc.ID](c.ID, objc.Sel("stop"))
}

// The currently configured frame rate on the external display that’s
// displaying the preview layer.
//
// # Discussion
//
// Observe this property to determine if the configured frame rate matches the
// max frame rate ([ActiveVideoMinFrameDuration]) of the device. When the
// [Active] property becomes `false`, this property changes to 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/activeExternalDisplayFrameRate
func (c AVCaptureExternalDisplayConfigurator) ActiveExternalDisplayFrameRate() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("activeExternalDisplayFrameRate"))
	return rv
}

// The device for which the coordinator configures the preview layer.
//
// # Discussion
//
// The value of this property is the [AVCaptureDevice] instance you provided
// when instantiating the configurator. [AVCaptureExternalDisplayConfigurator]
// holds a weak reference to the device. If the device is released, this
// property returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/device
func (c AVCaptureExternalDisplayConfigurator) Device() IAVCaptureDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("device"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}

// This property tells you whether the configurator is actively configuring
// the external display.
//
// # Discussion
//
// When this property returns `true`, the external display is successfully
// configured to match the device. If it returns`false`, the configurator is
// not making any configuration changes to the external display. If another
// [AVCaptureExternalDisplayConfigurator] instance takes over the
// configuration of the external display, this property returns `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isActive
func (c AVCaptureExternalDisplayConfigurator) Active() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isActive"))
	return rv
}

// The layer for which the configurator adjusts display properties to match
// the device’s state.
//
// # Discussion
//
// The value of this property is the [CALayer] instance that you provided when
// instantiating the configurator. You may specify either an
// [AVCaptureVideoPreviewLayer] or another [CALayer] instance that displays a
// camera’s video preview. [AVCaptureExternalDisplayConfigurator]holds a
// weak reference to the layer. If the layer is released, this property
// returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/previewLayer
func (c AVCaptureExternalDisplayConfigurator) PreviewLayer() quartzcore.CALayer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previewLayer"))
	return quartzcore.CALayerFromID(objc.ID(rv))
}

// The capture format in use by the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c AVCaptureExternalDisplayConfigurator) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureExternalDisplayConfigurator) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveFormat:"), value)
}

// Whether the external display supports matching frame rate to a capture
// device.
//
// # Discussion
//
// If `true`, you may instantiate a configurator with a configuration
// specifying [ShouldMatchFrameRate] set to `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isMatchingFrameRateSupported
func (_AVCaptureExternalDisplayConfiguratorClass AVCaptureExternalDisplayConfiguratorClass) ShouldMatchFrameRateSupported() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureExternalDisplayConfiguratorClass.class), objc.Sel("isMatchingFrameRateSupported"))
	return rv
}

// Whether the external display supports configuration to your preferred
// resolution.
//
// # Discussion
//
// If `true`, you may instantiate a configurator with a configuration
// specifying [PreferredResolution] set to `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isPreferredResolutionSupported
func (_AVCaptureExternalDisplayConfiguratorClass AVCaptureExternalDisplayConfiguratorClass) SupportsPreferredResolution() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureExternalDisplayConfiguratorClass.class), objc.Sel("isPreferredResolutionSupported"))
	return rv
}

// Whether the external display supports bypassing color space conversion.
//
// # Discussion
//
// If `true`, you may instantiate a configurator with a configuration
// specifying [BypassColorSpaceConversion] set to `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isBypassingColorSpaceConversionSupported
func (_AVCaptureExternalDisplayConfiguratorClass AVCaptureExternalDisplayConfiguratorClass) SupportsBypassingColorSpaceConversion() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureExternalDisplayConfiguratorClass.class), objc.Sel("isBypassingColorSpaceConversionSupported"))
	return rv
}
