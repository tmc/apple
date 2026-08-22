// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the information to provide about a track within a media asset.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader
type METrackReader interface {
	objectivec.IObject

	// Loads the track info object with the properties of the media asset track.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/loadTrackInfo(completionHandler:)
	LoadTrackInfoWithCompletionHandler(completionHandler METrackInfoErrorHandler)

	// Provides a new sample cursor that points to the sample at or near the specified presentation timestamp.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursor(atPresentationTimeStamp:completionHandler:)
	GenerateSampleCursorAtPresentationTimeStampCompletionHandler(presentationTimeStamp coremedia.CMTime, completionHandler MESampleCursorErrorHandler)

	// Provides a new sample cursor that points to the first sample in decode order.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursorAtFirstSampleInDecodeOrder(completionHandler:)
	GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler(completionHandler MESampleCursorErrorHandler)

	// Provides a new sample cursor that points to the last sample in decode order.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursorAtLastSampleInDecodeOrder(completionHandler:)
	GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler(completionHandler MESampleCursorErrorHandler)
}

// METrackReaderObject wraps an existing Objective-C object that conforms to the METrackReader protocol.
type METrackReaderObject struct {
	objectivec.Object
}

func (o METrackReaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// METrackReaderObjectFromID constructs a [METrackReaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func METrackReaderObjectFromID(id objc.ID) METrackReaderObject {
	return METrackReaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Loads the track info object with the properties of the media asset track.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// # Discussion
//
// If this method fails to create a track info object, it returns `nil` and
// the error contains information about the failure.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/loadTrackInfo(completionHandler:)
func (o METrackReaderObject) LoadTrackInfoWithCompletionHandler(completionHandler METrackInfoErrorHandler) {
	_block0, _ := NewMETrackInfoErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadTrackInfoWithCompletionHandler:"), _block0)
}

// Provides a new sample cursor that points to the sample at or near the
// specified presentation timestamp.
//
// presentationTimeStamp: The presentation time stamp.
//
// completionHandler: The completion block to execute when the generate operation finishes.
//
// # Discussion
//
// The new sample cursor points to the last sample with a presentation time
// stamp (PTS) less than or equal to `presentationTimeStamp`, or if there are
// no such samples, the first sample in PTS order.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursor(atPresentationTimeStamp:completionHandler:)
func (o METrackReaderObject) GenerateSampleCursorAtPresentationTimeStampCompletionHandler(presentationTimeStamp coremedia.CMTime, completionHandler MESampleCursorErrorHandler) {
	_block1, _ := NewMESampleCursorErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("generateSampleCursorAtPresentationTimeStamp:completionHandler:"), presentationTimeStamp, _block1)
}

// Provides a new sample cursor that points to the first sample in decode
// order.
//
// completionHandler: The completion block to execute when the generate operation finishes.
//
// # Discussion
//
// The new sample cursor points to the first sample in decode order,
// regardless of presentation time stamp (PTS).
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursorAtFirstSampleInDecodeOrder(completionHandler:)
func (o METrackReaderObject) GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler(completionHandler MESampleCursorErrorHandler) {
	_block0, _ := NewMESampleCursorErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("generateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler:"), _block0)
}

// Provides a new sample cursor that points to the last sample in decode
// order.
//
// completionHandler: The completion block to execute when the generate operation finishes.
//
// # Discussion
//
// The new sample cursor points to the last sample in decode order, regardless
// of presentation time stamp (PTS).
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/generateSampleCursorAtLastSampleInDecodeOrder(completionHandler:)
func (o METrackReaderObject) GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler(completionHandler MESampleCursorErrorHandler) {
	_block0, _ := NewMESampleCursorErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("generateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler:"), _block0)
}

// Loads the total size in bytes of all the samples in the track.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/loadTotalSampleDataLength(completionHandler:)
func (o METrackReaderObject) LoadTotalSampleDataLengthWithCompletionHandler(completionHandler int64_tErrorHandler) {
	_block0, _ := Newint64_tErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadTotalSampleDataLengthWithCompletionHandler:"), _block0)
}

// Loads the approximate data rate of the track in bytes per second.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// # Discussion
//
// If this method fails, it sets `estimatedDataRate` to `0.0` and the error
// contains information about the failure.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/loadEstimatedDataRate(completionHandler:)
func (o METrackReaderObject) LoadEstimatedDataRateWithCompletionHandler(completionHandler Float32ErrorHandler) {
	_block0, _ := NewFloat32ErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadEstimatedDataRateWithCompletionHandler:"), _block0)
}

// Loads the array of metadata items from the media asset track.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackReader/loadMetadata(completionHandler:)
func (o METrackReaderObject) LoadMetadataWithCompletionHandler(completionHandler AVMetadataItemArrayErrorHandler) {
	_block0, _ := NewAVMetadataItemArrayErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadMetadataWithCompletionHandler:"), _block0)
}
