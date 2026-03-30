// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether you can write the composition to the
// Saved Photos album.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isCompatibleWithSavedPhotosAlbum
func (m AVMutableMovie) IsCompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}
func (m AVMutableMovie) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setCompatibleWithSavedPhotosAlbum:"), value)
}
