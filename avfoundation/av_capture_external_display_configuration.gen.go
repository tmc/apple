// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureExternalDisplayConfiguration] class.
var (
	_AVCaptureExternalDisplayConfigurationClass     AVCaptureExternalDisplayConfigurationClass
	_AVCaptureExternalDisplayConfigurationClassOnce sync.Once
)

func getAVCaptureExternalDisplayConfigurationClass() AVCaptureExternalDisplayConfigurationClass {
	_AVCaptureExternalDisplayConfigurationClassOnce.Do(func() {
		_AVCaptureExternalDisplayConfigurationClass = AVCaptureExternalDisplayConfigurationClass{class: objc.GetClass("AVCaptureExternalDisplayConfiguration")}
	})
	return _AVCaptureExternalDisplayConfigurationClass
}

// GetAVCaptureExternalDisplayConfigurationClass returns the class object for AVCaptureExternalDisplayConfiguration.
func GetAVCaptureExternalDisplayConfigurationClass() AVCaptureExternalDisplayConfigurationClass {
	return getAVCaptureExternalDisplayConfigurationClass()
}

type AVCaptureExternalDisplayConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureExternalDisplayConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureExternalDisplayConfigurationClass) Alloc() AVCaptureExternalDisplayConfiguration {
	rv := objc.Send[AVCaptureExternalDisplayConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class you use to specify a configuration to your external display
// configurator.
//
// # Overview
//
// Using an [AVCaptureExternalDisplayConfiguration], you direct your
// [AVCaptureExternalDisplayConfigurator] how to configure an external display
// to match your device’s active video format.
//
// # Modifying the configuration
//
//   - [AVCaptureExternalDisplayConfiguration.BypassColorSpaceConversion]: A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//   - [AVCaptureExternalDisplayConfiguration.SetBypassColorSpaceConversion]
//   - [AVCaptureExternalDisplayConfiguration.PreferredResolution]: Your preferred external display resolution.
//   - [AVCaptureExternalDisplayConfiguration.SetPreferredResolution]
//   - [AVCaptureExternalDisplayConfiguration.ShouldMatchFrameRate]: A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//   - [AVCaptureExternalDisplayConfiguration.SetShouldMatchFrameRate]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration
type AVCaptureExternalDisplayConfiguration struct {
	objectivec.Object
}

// AVCaptureExternalDisplayConfigurationFromID constructs a [AVCaptureExternalDisplayConfiguration] from an objc.ID.
//
// A class you use to specify a configuration to your external display
// configurator.
func AVCaptureExternalDisplayConfigurationFromID(id objc.ID) AVCaptureExternalDisplayConfiguration {
	return AVCaptureExternalDisplayConfiguration{objectivec.Object{ID: id}}
}

// NOTE: AVCaptureExternalDisplayConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureExternalDisplayConfiguration] class.
//
// # Modifying the configuration
//
//   - [IAVCaptureExternalDisplayConfiguration.BypassColorSpaceConversion]: A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//   - [IAVCaptureExternalDisplayConfiguration.SetBypassColorSpaceConversion]
//   - [IAVCaptureExternalDisplayConfiguration.PreferredResolution]: Your preferred external display resolution.
//   - [IAVCaptureExternalDisplayConfiguration.SetPreferredResolution]
//   - [IAVCaptureExternalDisplayConfiguration.ShouldMatchFrameRate]: A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//   - [IAVCaptureExternalDisplayConfiguration.SetShouldMatchFrameRate]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration
type IAVCaptureExternalDisplayConfiguration interface {
	objectivec.IObject

	// Topic: Modifying the configuration

	// A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
	BypassColorSpaceConversion() bool
	SetBypassColorSpaceConversion(value bool)
	// Your preferred external display resolution.
	PreferredResolution() coremedia.CMVideoDimensions
	SetPreferredResolution(value coremedia.CMVideoDimensions)
	// A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
	ShouldMatchFrameRate() bool
	SetShouldMatchFrameRate(value bool)
}

// Init initializes the instance.
func (c AVCaptureExternalDisplayConfiguration) Init() AVCaptureExternalDisplayConfiguration {
	rv := objc.Send[AVCaptureExternalDisplayConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureExternalDisplayConfiguration) Autorelease() AVCaptureExternalDisplayConfiguration {
	rv := objc.Send[AVCaptureExternalDisplayConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureExternalDisplayConfiguration creates a new AVCaptureExternalDisplayConfiguration instance.
func NewAVCaptureExternalDisplayConfiguration() AVCaptureExternalDisplayConfiguration {
	class := getAVCaptureExternalDisplayConfigurationClass()
	rv := objc.Send[AVCaptureExternalDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A property indicating whether the color space of the configurator’s
// preview layer should be preserved on the output display by avoiding color
// space conversions.
//
// # Discussion
//
// Set [BypassColorSpaceConversion] to `true` if you would like the
// configurator’s [AVCaptureVideoPreviewLayer] color space preserved on the
// output display. This is accomplished by setting the working color space to
// match the color space of the external display. The color properties of the
// [CALayer] remain untouched. The default value is `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/bypassColorSpaceConversion
func (c AVCaptureExternalDisplayConfiguration) BypassColorSpaceConversion() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("bypassColorSpaceConversion"))
	return rv
}
func (c AVCaptureExternalDisplayConfiguration) SetBypassColorSpaceConversion(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setBypassColorSpaceConversion:"), value)
}

// Your preferred external display resolution.
//
// # Discussion
//
// Use [PreferredResolution] to set your desired resolution of the external
// display. When left at the default value of { 0, 0 }, the native resolution
// of the external display is used.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/preferredResolution
func (c AVCaptureExternalDisplayConfiguration) PreferredResolution() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("preferredResolution"))
	return coremedia.CMVideoDimensions(rv)
}
func (c AVCaptureExternalDisplayConfiguration) SetPreferredResolution(value coremedia.CMVideoDimensions) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredResolution:"), value)
}

// A property indicating whether the frame rate of the external display should
// be configured to match the camera’s frame rate.
//
// # Discussion
//
// If you want to configure your [AVCaptureVideoPreviewLayer] to match its
// source [ActiveVideoMinFrameDuration], set [ShouldMatchFrameRate] to `true`.
// The default value is `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/shouldMatchFrameRate
func (c AVCaptureExternalDisplayConfiguration) ShouldMatchFrameRate() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldMatchFrameRate"))
	return rv
}
func (c AVCaptureExternalDisplayConfiguration) SetShouldMatchFrameRate(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldMatchFrameRate:"), value)
}
