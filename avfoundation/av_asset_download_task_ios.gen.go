// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
)

// The local file URL to where the task downloads the asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask/destinationURL
func (a AVAssetDownloadTask) DestinationURL() foundation.INSURL {
rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationURL"))
return foundation.NSURLFromID(rv)
}

