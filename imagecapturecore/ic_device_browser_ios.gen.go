// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.
//go:build ios
// +build ios

package imagecapturecore

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/requestContentsAuthorization(completion:)
func (d ICDeviceBrowser) RequestContentsAuthorizationWithCompletion(completion StringHandler) {
	_block0, _ := NewStringBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("requestContentsAuthorizationWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/requestControlAuthorization(completion:)
func (d ICDeviceBrowser) RequestControlAuthorizationWithCompletion(completion StringHandler) {
	_block0, _ := NewStringBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("requestControlAuthorizationWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/resetContentsAuthorization(completion:)
func (d ICDeviceBrowser) ResetContentsAuthorizationWithCompletion(completion StringHandler) {
	_block0, _ := NewStringBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("resetContentsAuthorizationWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/resetControlAuthorization(completion:)
func (d ICDeviceBrowser) ResetControlAuthorizationWithCompletion(completion StringHandler) {
	_block0, _ := NewStringBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("resetControlAuthorizationWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/contentsAuthorizationStatus
func (d ICDeviceBrowser) ContentsAuthorizationStatus() ICAuthorizationStatus {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("contentsAuthorizationStatus"))
	return ICAuthorizationStatus(foundation.NSStringFromID(rv).String())
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/controlAuthorizationStatus
func (d ICDeviceBrowser) ControlAuthorizationStatus() ICAuthorizationStatus {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("controlAuthorizationStatus"))
	return ICAuthorizationStatus(foundation.NSStringFromID(rv).String())
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/isSuspended
func (d ICDeviceBrowser) IsSuspended() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isSuspended"))
	return rv
}
