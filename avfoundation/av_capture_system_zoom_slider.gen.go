// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVCaptureSystemZoomSlider] class.
var (
	_AVCaptureSystemZoomSliderClass     AVCaptureSystemZoomSliderClass
	_AVCaptureSystemZoomSliderClassOnce sync.Once
)

func getAVCaptureSystemZoomSliderClass() AVCaptureSystemZoomSliderClass {
	_AVCaptureSystemZoomSliderClassOnce.Do(func() {
		_AVCaptureSystemZoomSliderClass = AVCaptureSystemZoomSliderClass{class: objc.GetClass("AVCaptureSystemZoomSlider")}
	})
	return _AVCaptureSystemZoomSliderClass
}

// GetAVCaptureSystemZoomSliderClass returns the class object for AVCaptureSystemZoomSlider.
func GetAVCaptureSystemZoomSliderClass() AVCaptureSystemZoomSliderClass {
	return getAVCaptureSystemZoomSliderClass()
}

type AVCaptureSystemZoomSliderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureSystemZoomSliderClass) Alloc() AVCaptureSystemZoomSlider {
	rv := objc.Send[AVCaptureSystemZoomSlider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A control that adjusts the video zoom factor of a capture device within the
// system-recommended range.
//
// # Overview
// 
// The system sets the slider’s range to the value of the
// [systemRecommendedVideoZoomRange] property of the device’s active format.
// If a device’s [AVCaptureSystemZoomSlider.ActiveFormat] value changes, the slider updates its range
// to the new format’s recommendation.
// 
// To use this control, add it to the capture session by calling the
// session’s [AddControl] method.
//
// [systemRecommendedVideoZoomRange]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/systemRecommendedVideoZoomRange
//
// # Creating a zoom slider
//
//   - [AVCaptureSystemZoomSlider.InitWithDevice]: Creates a slider to control the video zoom factor of a capture device.
//   - [AVCaptureSystemZoomSlider.InitWithDeviceAction]: Creates a slider to control the zoom level of the specified capture device with an action to respond to zoom changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider
type AVCaptureSystemZoomSlider struct {
	AVCaptureControl
}

// AVCaptureSystemZoomSliderFromID constructs a [AVCaptureSystemZoomSlider] from an objc.ID.
//
// A control that adjusts the video zoom factor of a capture device within the
// system-recommended range.
func AVCaptureSystemZoomSliderFromID(id objc.ID) AVCaptureSystemZoomSlider {
	return AVCaptureSystemZoomSlider{AVCaptureControl: AVCaptureControlFromID(id)}
}
// NOTE: AVCaptureSystemZoomSlider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureSystemZoomSlider] class.
//
// # Creating a zoom slider
//
//   - [IAVCaptureSystemZoomSlider.InitWithDevice]: Creates a slider to control the video zoom factor of a capture device.
//   - [IAVCaptureSystemZoomSlider.InitWithDeviceAction]: Creates a slider to control the zoom level of the specified capture device with an action to respond to zoom changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider
type IAVCaptureSystemZoomSlider interface {
	IAVCaptureControl

	// Topic: Creating a zoom slider

	// Creates a slider to control the video zoom factor of a capture device.
	InitWithDevice(device IAVCaptureDevice) AVCaptureSystemZoomSlider
	// Creates a slider to control the zoom level of the specified capture device with an action to respond to zoom changes.
	InitWithDeviceAction(device IAVCaptureDevice, action Float64Handler) AVCaptureSystemZoomSlider

	// The capture format in use by the device.
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	// The system’s recommended zoom range for this device format.
	SystemRecommendedVideoZoomRange() float64
	SetSystemRecommendedVideoZoomRange(value float64)
}

// Init initializes the instance.
func (c AVCaptureSystemZoomSlider) Init() AVCaptureSystemZoomSlider {
	rv := objc.Send[AVCaptureSystemZoomSlider](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureSystemZoomSlider) Autorelease() AVCaptureSystemZoomSlider {
	rv := objc.Send[AVCaptureSystemZoomSlider](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSystemZoomSlider creates a new AVCaptureSystemZoomSlider instance.
func NewAVCaptureSystemZoomSlider() AVCaptureSystemZoomSlider {
	class := getAVCaptureSystemZoomSliderClass()
	rv := objc.Send[AVCaptureSystemZoomSlider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a slider to control the video zoom factor of a capture device.
//
// device: The capture device to control.
//
// # Discussion
// 
// You can only create a zoom slider with a device that support’s setting
// its [VideoZoomFactor] property value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider/init(device:)
func NewCaptureSystemZoomSliderWithDevice(device IAVCaptureDevice) AVCaptureSystemZoomSlider {
	instance := getAVCaptureSystemZoomSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return AVCaptureSystemZoomSliderFromID(rv)
}

// Creates a slider to control the video zoom factor of a capture device.
//
// device: The capture device to control.
//
// # Discussion
// 
// You can only create a zoom slider with a device that support’s setting
// its [VideoZoomFactor] property value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider/init(device:)
func (c AVCaptureSystemZoomSlider) InitWithDevice(device IAVCaptureDevice) AVCaptureSystemZoomSlider {
	rv := objc.Send[AVCaptureSystemZoomSlider](c.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
// Creates a slider to control the zoom level of the specified capture device
// with an action to respond to zoom changes.
//
// device: The capture device to control.
//
// action: An action the system calls on the main actor to respond to changes to the
// device’s [VideoZoomFactor] property.
//
// # Discussion
// 
// The system calls the specified action only when the zoom slider changes the
// device’s [VideoZoomFactor] property value. If your app needs to react to
// other sources of video zoom factor changes like
// [RampToVideoZoomFactorWithRate], use key-value observation instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider/init(device:action:)
func (c AVCaptureSystemZoomSlider) InitWithDeviceAction(device IAVCaptureDevice, action Float64Handler) AVCaptureSystemZoomSlider {
_block1, _cleanup1 := NewFloat64Block(action)
	defer _cleanup1()
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithDevice:action:"), device, _block1)
	return AVCaptureSystemZoomSliderFromID(rv)
}

// The capture format in use by the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c AVCaptureSystemZoomSlider) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureSystemZoomSlider) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveFormat:"), value)
}
// The system’s recommended zoom range for this device format.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedvideozoomrange
func (c AVCaptureSystemZoomSlider) SystemRecommendedVideoZoomRange() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("systemRecommendedVideoZoomRange"))
	return rv
}
func (c AVCaptureSystemZoomSlider) SetSystemRecommendedVideoZoomRange(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setSystemRecommendedVideoZoomRange:"), value)
}

// InitWithDeviceActionSync is a synchronous wrapper around [AVCaptureSystemZoomSlider.InitWithDeviceAction].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureSystemZoomSlider) InitWithDeviceActionSync(ctx context.Context, device IAVCaptureDevice) (float64, error) {
	done := make(chan float64, 1)
	c.InitWithDeviceAction(device, func(val float64) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0.0, ctx.Err()
	}
}

