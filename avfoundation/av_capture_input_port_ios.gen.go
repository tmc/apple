// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The device type of the source camera that provides data to the port.
//
// # Discussion
//
// All ports contained in an input’s [Ports] property have the same source
// device type, which each equal to the [DeviceType] property of the input’s
// device.
//
// When working with virtual devices such as the [builtInDualCamera] in an
// [AVCaptureMultiCamSession], it’s possible to stream media from the
// virtual device’s constituent device streams by discovering and connecting
// hidden ports. In the case of the [builtInDualCamera], its constituent
// devices are the wide-angle and telephoto cameras.
//
// By calling [PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition]:, you
// may discover ports originating from one or more of the virtual device’s
// constituent devices and then make connections using those ports.
// Constituent device ports are never present in their owning virtual device
// input’s ports array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/sourceDeviceType
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
func (c AVCaptureInputPort) SourceDeviceType() AVCaptureDeviceType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sourceDeviceType"))
	return AVCaptureDeviceType(foundation.NSStringFromID(rv).String())
}

// The position of the source device providing input through this port.
//
// # Discussion
//
// All ports contained in an [AVCaptureInput] object’s [Ports] array have
// the same [SourceDevicePosition] value.
//
// When working with a microphone input in an [AVCaptureMultiCamSession],
// it’s possible to record multiple microphone directions simultaneously.
// For example, you can record audio from the front microphone input to pair
// with video from the front camera, and record audio from the back microphone
// input to pair with video from the back camera.
//
// By calling the input’s
// [PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition] method, you may
// discover additional hidden ports originating from the source audio device.
// These ports represent individual microphones positioned to pick up audio
// from one particular direction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/sourceDevicePosition
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
func (c AVCaptureInputPort) SourceDevicePosition() AVCaptureDevicePosition {
	rv := objc.Send[AVCaptureDevicePosition](c.ID, objc.Sel("sourceDevicePosition"))
	return AVCaptureDevicePosition(rv)
}
