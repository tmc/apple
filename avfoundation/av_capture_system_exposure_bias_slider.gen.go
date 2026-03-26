// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVCaptureSystemExposureBiasSlider] class.
var (
	_AVCaptureSystemExposureBiasSliderClass     AVCaptureSystemExposureBiasSliderClass
	_AVCaptureSystemExposureBiasSliderClassOnce sync.Once
)

func getAVCaptureSystemExposureBiasSliderClass() AVCaptureSystemExposureBiasSliderClass {
	_AVCaptureSystemExposureBiasSliderClassOnce.Do(func() {
		_AVCaptureSystemExposureBiasSliderClass = AVCaptureSystemExposureBiasSliderClass{class: objc.GetClass("AVCaptureSystemExposureBiasSlider")}
	})
	return _AVCaptureSystemExposureBiasSliderClass
}

// GetAVCaptureSystemExposureBiasSliderClass returns the class object for AVCaptureSystemExposureBiasSlider.
func GetAVCaptureSystemExposureBiasSliderClass() AVCaptureSystemExposureBiasSliderClass {
	return getAVCaptureSystemExposureBiasSliderClass()
}

type AVCaptureSystemExposureBiasSliderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureSystemExposureBiasSliderClass) Alloc() AVCaptureSystemExposureBiasSlider {
	rv := objc.Send[AVCaptureSystemExposureBiasSlider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A control that adjusts the exposure bias of a capture device within the
// system-recommended range.
//
// # Overview
// 
// This control defines its range by querying the
// [systemRecommendedExposureBiasRange] property of the device’s active
// format. If a device’s [AVCaptureSystemExposureBiasSlider.ActiveFormat] value changes, the slider updates
// its range with the new format’s system-recommended value.
// 
// To use this control, add it to the capture session by calling the
// session’s [AddControl] method.
//
// [systemRecommendedExposureBiasRange]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/systemRecommendedExposureBiasRange
//
// # Creating an exposure bias slider
//
//   - [AVCaptureSystemExposureBiasSlider.InitWithDevice]: Creates a slider to control the exposure bias of the specified capture device.
//   - [AVCaptureSystemExposureBiasSlider.InitWithDeviceAction]: Creates a slider to control the exposure bias of the specified capture device with an action to respond to exposure bias changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider
type AVCaptureSystemExposureBiasSlider struct {
	AVCaptureControl
}

// AVCaptureSystemExposureBiasSliderFromID constructs a [AVCaptureSystemExposureBiasSlider] from an objc.ID.
//
// A control that adjusts the exposure bias of a capture device within the
// system-recommended range.
func AVCaptureSystemExposureBiasSliderFromID(id objc.ID) AVCaptureSystemExposureBiasSlider {
	return AVCaptureSystemExposureBiasSlider{AVCaptureControl: AVCaptureControlFromID(id)}
}
// NOTE: AVCaptureSystemExposureBiasSlider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureSystemExposureBiasSlider] class.
//
// # Creating an exposure bias slider
//
//   - [IAVCaptureSystemExposureBiasSlider.InitWithDevice]: Creates a slider to control the exposure bias of the specified capture device.
//   - [IAVCaptureSystemExposureBiasSlider.InitWithDeviceAction]: Creates a slider to control the exposure bias of the specified capture device with an action to respond to exposure bias changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider
type IAVCaptureSystemExposureBiasSlider interface {
	IAVCaptureControl

	// Topic: Creating an exposure bias slider

	// Creates a slider to control the exposure bias of the specified capture device.
	InitWithDevice(device IAVCaptureDevice) AVCaptureSystemExposureBiasSlider
	// Creates a slider to control the exposure bias of the specified capture device with an action to respond to exposure bias changes.
	InitWithDeviceAction(device IAVCaptureDevice, action Float32Handler) AVCaptureSystemExposureBiasSlider

	// The capture format in use by the device.
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	// The system’s recommended exposure bias range for this device format.
	SystemRecommendedExposureBiasRange() float32
	SetSystemRecommendedExposureBiasRange(value float32)
}

// Init initializes the instance.
func (c AVCaptureSystemExposureBiasSlider) Init() AVCaptureSystemExposureBiasSlider {
	rv := objc.Send[AVCaptureSystemExposureBiasSlider](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureSystemExposureBiasSlider) Autorelease() AVCaptureSystemExposureBiasSlider {
	rv := objc.Send[AVCaptureSystemExposureBiasSlider](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSystemExposureBiasSlider creates a new AVCaptureSystemExposureBiasSlider instance.
func NewAVCaptureSystemExposureBiasSlider() AVCaptureSystemExposureBiasSlider {
	class := getAVCaptureSystemExposureBiasSliderClass()
	rv := objc.Send[AVCaptureSystemExposureBiasSlider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a slider to control the exposure bias of the specified capture
// device.
//
// device: The capture device to control.
//
// # Discussion
// 
// You can only create an exposure bias slider with a device that support’s
// setting its [ExposureTargetBias] property value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:)
func NewCaptureSystemExposureBiasSliderWithDevice(device IAVCaptureDevice) AVCaptureSystemExposureBiasSlider {
	instance := getAVCaptureSystemExposureBiasSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return AVCaptureSystemExposureBiasSliderFromID(rv)
}

// Creates a slider to control the exposure bias of the specified capture
// device.
//
// device: The capture device to control.
//
// # Discussion
// 
// You can only create an exposure bias slider with a device that support’s
// setting its [ExposureTargetBias] property value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:)
func (c AVCaptureSystemExposureBiasSlider) InitWithDevice(device IAVCaptureDevice) AVCaptureSystemExposureBiasSlider {
	rv := objc.Send[AVCaptureSystemExposureBiasSlider](c.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
// Creates a slider to control the exposure bias of the specified capture
// device with an action to respond to exposure bias changes.
//
// device: The capture device to control.
//
// action: An action the system calls on the main actor to handle changes to the
// device’s [ExposureTargetBias] property.
//
// # Discussion
// 
// The system only calls the specified action when the exposure bias slider
// changes the device’s [VideoZoomFactor] property value. If you need to
// react to other sources of changes to the exposure target bias, use
// key-value observation instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:action:)
func (c AVCaptureSystemExposureBiasSlider) InitWithDeviceAction(device IAVCaptureDevice, action Float32Handler) AVCaptureSystemExposureBiasSlider {
_block1, _cleanup1 := NewFloat32Block(action)
	defer _cleanup1()
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithDevice:action:"), device, _block1)
	return AVCaptureSystemExposureBiasSliderFromID(rv)
}

// The capture format in use by the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c AVCaptureSystemExposureBiasSlider) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureSystemExposureBiasSlider) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveFormat:"), value)
}
// The system’s recommended exposure bias range for this device format.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c AVCaptureSystemExposureBiasSlider) SystemRecommendedExposureBiasRange() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return rv
}
func (c AVCaptureSystemExposureBiasSlider) SetSystemRecommendedExposureBiasRange(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setSystemRecommendedExposureBiasRange:"), value)
}

// InitWithDeviceActionSync is a synchronous wrapper around [AVCaptureSystemExposureBiasSlider.InitWithDeviceAction].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureSystemExposureBiasSlider) InitWithDeviceActionSync(ctx context.Context, device IAVCaptureDevice) (float32, error) {
	done := make(chan float32, 1)
	c.InitWithDeviceAction(device, func(val float32) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0.0, ctx.Err()
	}
}

