// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether you can write the composition to the
// Saved Photos album.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isCompatibleWithSavedPhotosAlbum
func (c AVComposition) IsCompatibleWithSavedPhotosAlbum() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
		return rv
}
func (c AVComposition) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setCompatibleWithSavedPhotosAlbum:"), value)
}

