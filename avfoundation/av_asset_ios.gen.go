// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/corefoundation"
)

// A Boolean value that indicates whether you can write the asset to the Saved
// Photos album.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/isCompatibleWithSavedPhotosAlbum
func (a AVAsset) CompatibleWithSavedPhotosAlbum() bool {
rv := objc.Send[bool](a.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
		return rv
}
// The encoded or authored size of the visual portion of the asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/naturalSize
func (a AVAsset) NaturalSize() corefoundation.CGSize {
rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("naturalSize"))
		return corefoundation.CGSize(rv)
}
// The asset’s display mode preference for optimal playback of its content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredDisplayCriteria
func (a AVAsset) PreferredDisplayCriteria() objc.ID {
rv := objc.Send[objc.ID](a.ID, objc.Sel("preferredDisplayCriteria"))
		return rv
}

