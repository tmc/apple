// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods you can implement to provide additional metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutputPushDelegate
type AVPlayerItemMetadataOutputPushDelegate interface {
	objectivec.IObject
	AVPlayerItemOutputPushDelegate
}

// AVPlayerItemMetadataOutputPushDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemMetadataOutputPushDelegate protocol.
type AVPlayerItemMetadataOutputPushDelegateObject struct {
	objectivec.Object
}
func (o AVPlayerItemMetadataOutputPushDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemMetadataOutputPushDelegateObjectFromID constructs a [AVPlayerItemMetadataOutputPushDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemMetadataOutputPushDelegateObjectFromID(id objc.ID) AVPlayerItemMetadataOutputPushDelegateObject {
	return AVPlayerItemMetadataOutputPushDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate a new collection of metadata items is available.
//
// output: The [AVPlayerItemMetadataOutput] source.
//
// groups: An array of [AVTimedMetadataGroups] that contain metadata items with
// requested identifiers, according to the format descriptions associated with
// the underlying tracks.
//
// track: An instance of [AVPlayerItemTrack] that indicates the source of the
// metadata items in the group.
//
// # Discussion
// 
// Each group provided in a single invocation of this method will have timing
// that does not overlap with any other group in the array.
// 
// Note that for some timed metadata formats carried by HTTP live streaming,
// the `timeRange` of each group must be reported as [indefinite], because its
// duration will be unknown until the next metadata group in the stream
// arrives. In these cases, the groups parameter will always contain a single
// group.
// 
// Groups are typically packaged into arrays for delivery to your delegate
// according to the chunking or interleaving of the underlying metadata data.
// 
// Note that if the item carries multiple metadata tracks containing metadata
// with the same metadata identifiers, this method can be invoked for each one
// separately, each with reference to the associated [AVPlayerItemTrack].
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutputPushDelegate/metadataOutput(_:didOutputTimedMetadataGroups:from:)
func (o AVPlayerItemMetadataOutputPushDelegateObject) MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output IAVPlayerItemMetadataOutput, groups []AVTimedMetadataGroup, track IAVPlayerItemTrack) {
	objc.Send[struct{}](o.ID, objc.Sel("metadataOutput:didOutputTimedMetadataGroups:fromPlayerItemTrack:"), output, objectivec.IObjectSliceToNSArray(groups), track)
	}
// Tells the delegate that the output is starting a new sequence of media
// data.
//
// output: The [AVPlayerItemOutput] object.
//
// # Discussion
// 
// This method is invoked after any seeking and change in playback direction.
// If you are maintaining any queued future media data, you may want to
// discard those objects after receiving this message.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPushDelegate/outputSequenceWasFlushed(_:)
func (o AVPlayerItemMetadataOutputPushDelegateObject) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasFlushed:"), output)
	}

