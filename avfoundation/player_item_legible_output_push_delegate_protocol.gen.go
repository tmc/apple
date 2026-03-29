// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Methods you can implement to provide alternative attributed-string output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutputPushDelegate
type AVPlayerItemLegibleOutputPushDelegate interface {
	objectivec.IObject
	AVPlayerItemOutputPushDelegate
}

// AVPlayerItemLegibleOutputPushDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemLegibleOutputPushDelegate protocol.
type AVPlayerItemLegibleOutputPushDelegateObject struct {
	objectivec.Object
}
func (o AVPlayerItemLegibleOutputPushDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemLegibleOutputPushDelegateObjectFromID constructs a [AVPlayerItemLegibleOutputPushDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemLegibleOutputPushDelegateObjectFromID(id objc.ID) AVPlayerItemLegibleOutputPushDelegateObject {
	return AVPlayerItemLegibleOutputPushDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate to process the delivery of new textual samples.
//
// output: The [AVPlayerItemLegibleOutput] source instance.
//
// strings: An array of [NSAttributedString] objects, each containing both the run of
// text and the descriptive markup.
// //
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// nativeSamples: An array of [CMSampleBuffer] objects, for media subtypes included in the
// array passed to the `output` object’s
// [InitWithMediaSubtypesForNativeRepresentation] method.
// //
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// itemTime: The item time at which the strings should be presented.
//
// # Discussion
// 
// For each media subtype in the array passed in to the `output` object’s
// [InitWithMediaSubtypesForNativeRepresentation] method, the delegate
// receives sample buffers carrying data in its native format via the
// `nativeSamples` parameter if there is media data of that subtype in the
// media resource.
// 
// For all other media subtypes present in the media resource, the delegate
// receives attributed strings in a common format via the `strings` parameter.
// See [CMTextMarkup] for the string attributes keys and values that are used
// in the attributed strings.
//
// [CMTextMarkup]: https://developer.apple.com/documentation/CoreMedia/cmtextmarkup
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutputPushDelegate/legibleOutput(_:didOutputAttributedStrings:nativeSampleBuffers:forItemTime:)
func (o AVPlayerItemLegibleOutputPushDelegateObject) LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output IAVPlayerItemLegibleOutput, strings []foundation.NSAttributedString, nativeSamples foundation.INSArray, itemTime coremedia.CMTime) {
	objc.Send[struct{}](o.ID, objc.Sel("legibleOutput:didOutputAttributedStrings:nativeSampleBuffers:forItemTime:"), output, objectivec.IObjectSliceToNSArray(strings), nativeSamples, itemTime)
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
func (o AVPlayerItemLegibleOutputPushDelegateObject) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasFlushed:"), output)
	}

