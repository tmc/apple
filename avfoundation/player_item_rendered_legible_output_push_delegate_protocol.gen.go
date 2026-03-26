// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A delegate that handles the rendered pixel buffers produced by a rendered legible output object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutputPushDelegate
type AVPlayerItemRenderedLegibleOutputPushDelegate interface {
	objectivec.IObject
	AVPlayerItemOutputPushDelegate
}

// AVPlayerItemRenderedLegibleOutputPushDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemRenderedLegibleOutputPushDelegate protocol.
type AVPlayerItemRenderedLegibleOutputPushDelegateObject struct {
	objectivec.Object
}
func (o AVPlayerItemRenderedLegibleOutputPushDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemRenderedLegibleOutputPushDelegateObjectFromID constructs a [AVPlayerItemRenderedLegibleOutputPushDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemRenderedLegibleOutputPushDelegateObjectFromID(id objc.ID) AVPlayerItemRenderedLegibleOutputPushDelegateObject {
	return AVPlayerItemRenderedLegibleOutputPushDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that new rendered caption images are available.
//
// output: The rendered legible output object.
//
// captionImages: An array of [AVRenderedCaptionImage] objects. A caption object consists of
// a [CVPixelBuffer] and its associated position, in pixels, relative to the
// video frame.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
//
// itemTime: The item time at which to present the caption images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutputPushDelegate/renderedLegibleOutput(_:didOutputRenderedCaptionImages:forItemTime:)
func (o AVPlayerItemRenderedLegibleOutputPushDelegateObject) RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(output IAVPlayerItemRenderedLegibleOutput, captionImages []AVRenderedCaptionImage, itemTime objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("renderedLegibleOutput:didOutputRenderedCaptionImages:forItemTime:"), output, objectivec.IObjectSliceToNSArray(captionImages), itemTime)
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
func (o AVPlayerItemRenderedLegibleOutputPushDelegateObject) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasFlushed:"), output)
	}

