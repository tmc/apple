// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
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
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/supportedMultiCamDeviceSets
func (c AVCaptureDeviceDiscoverySession) SupportedMultiCamDeviceSets() foundation.INSSet {
rv := objc.Send[objc.ID](c.ID, objc.Sel("supportedMultiCamDeviceSets"))
return foundation.NSSetFromID(rv)
}

