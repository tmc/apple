// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Sets of capture devices that you can use simultaneously in a multi-camera
// session.
//
// # Discussion
//
// You may use multiple cameras as device inputs to an
// [AVCaptureMultiCamSession], as long as one of the supported multi-camera
// device sets includes the device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/supportedMultiCamDeviceSets
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
func (c AVCaptureDeviceDiscoverySession) SupportedMultiCamDeviceSets() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportedMultiCamDeviceSets"))
	return foundation.NSSetFromID(rv)
}
