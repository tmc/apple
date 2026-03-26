// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
)

// A count of the media segments downloaded from the server to this client.
//
// # Discussion
// 
// The property corresponds to “sc-count”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfSegmentsDownloaded
func (p AVPlayerItemAccessLogEvent) NumberOfSegmentsDownloaded() int {
rv := objc.Send[int](p.ID, objc.Sel("numberOfSegmentsDownloaded"))
		return rv
}

