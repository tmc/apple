// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the information to provide about samples within a track of a media asset, and enables stepping through samples in the track in decode or presentation order.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor
type MESampleCursor interface {
	objectivec.IObject
	foundation.NSCopying

	// Moves the cursor a given number of samples in decode order.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/stepInDecodeOrder(by:completionHandler:)
	StepInDecodeOrderByCountCompletionHandler(stepCount int64, completionHandler int64_tErrorHandler)

	// Moves the cursor a given number of samples in presentation order.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/stepInPresentationOrder(by:completionHandler:)
	StepInPresentationOrderByCountCompletionHandler(stepCount int64, completionHandler int64_tErrorHandler)

	// The presentation timestamp (PTS) of the sample at the current position of the cursor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/presentationTimeStamp
	PresentationTimeStamp() coremedia.CMTime

	// The decode timestamp (DTS) of the sample at the current position of the cursor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/decodeTimeStamp
	DecodeTimeStamp() coremedia.CMTime

	// The decode duration of the sample at the current position.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/currentSampleDuration
	CurrentSampleDuration() coremedia.CMTime

	// The format description for the sample at the current position of the cursor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/currentSampleFormatDescription
	CurrentSampleFormatDescription() coremedia.CMFormatDescriptionRef

	// Decoder synchronization information about the sample the cursor points to.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/syncInfo
	SyncInfo() avfoundation.AVSampleCursorSyncInfo

	// Dependency information about the sample the cursor points to.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/dependencyInfo
	DependencyInfo() avfoundation.AVSampleCursorDependencyInfo

	// Additional information that’s necessary to recover complete sample dependency information.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/hevcDependencyInfo
	HevcDependencyInfo() IMEHEVCDependencyInfo

	// The duration of the playable content starting from the cursor position.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/decodeTimeOfLastSampleReachableByForwardSteppingThatIsAlreadyLoadedByByteSource
	DecodeTimeOfLastSampleReachableByForwardSteppingThatIsAlreadyLoadedByByteSource() coremedia.CMTime
}

// MESampleCursorObject wraps an existing Objective-C object that conforms to the MESampleCursor protocol.
type MESampleCursorObject struct {
	foundation.NSCopyingObject
}

func (o MESampleCursorObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MESampleCursorObjectFromID constructs a [MESampleCursorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MESampleCursorObjectFromID(id objc.ID) MESampleCursorObject {
	return MESampleCursorObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// Moves the cursor a given number of samples in decode order.
//
// stepCount: The number of samples to move. If positive, the cursor steps forward. If
// negative, the cursor steps backward.
//
// completionHandler: The completion block to execute when the move operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/stepInDecodeOrder(by:completionHandler:)
func (o MESampleCursorObject) StepInDecodeOrderByCountCompletionHandler(stepCount int64, completionHandler int64_tErrorHandler) {
	_block1, _ := Newint64_tErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("stepInDecodeOrderByCount:completionHandler:"), stepCount, _block1)
}

// Moves the cursor a given number of samples in presentation order.
//
// stepCount: The number of samples to move. If positive, the cursor steps forward. If
// negative, the cursor steps backward.
//
// completionHandler: The completion block to execute when the move operation finishes.
//
// # Discussion
//
// If the request would advance the cursor past the last sample or before the
// first sample, the cursor points to that limiting sample and
// `actualStepCount` is equal to the number of samples the cursor moved. If
// decode order and presentation order are the same, in other words, the
// samples aren’t reordered, this method has the same effect as
// [stepInDecodeOrder(byCount:)].
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/stepInPresentationOrder(by:completionHandler:)
//
// [stepInDecodeOrder(byCount:)]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/stepInDecodeOrder(byCount:)
func (o MESampleCursorObject) StepInPresentationOrderByCountCompletionHandler(stepCount int64, completionHandler int64_tErrorHandler) {
	_block1, _ := Newint64_tErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("stepInPresentationOrderByCount:completionHandler:"), stepCount, _block1)
}

// Tests for an earlier boundary in sample reordering.
//
// cursor: A sample cursor to use to test the sample reordering boundary.
//
// # Return Value
//
// true if it’s possible that earlier samples in decode order can have a
// later presentation timestamp than that of the specified cursor; otherwise
// false.
//
// # Discussion
//
// This method tests for a boundary in the reordering from decode order to
// presentation order. This determines when it’s possible for any sample
// earlier in decode order than the current sample to have a later
// presentation time than the current sample of the specified cursor. You can
// use this test to limit backward scans, such as to start forward playback.
// For example, with the argument cursor fixed, step the cursor backward until
// it’s impossible for any earlier-in-decode-order samples to be
// later-in-presentation-order than the argument cursor sample.
//
// Don’t implement this method for formats where sample reordering doesn’t
// make sense for the track content, which also indicates that the samples
// aren’t reordered.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/samplesWithEarlierDTSsMayHaveLaterPTSs(than:)
func (o MESampleCursorObject) SamplesWithEarlierDTSsMayHaveLaterPTSsThanCursor(cursor MESampleCursor) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("samplesWithEarlierDTSsMayHaveLaterPTSsThanCursor:"), cursor)
	return rv
}

// Tests for a later boundary in sample reordering.
//
// cursor: A sample cursor to use to test the sample reordering boundary.
//
// # Return Value
//
// true if it’s possible that later samples in decode order can have an
// earlier presentation timestamp than that of the specified cursor; otherwise
// false.
//
// # Discussion
//
// This method tests for a boundary in the reordering from decode order to
// presentation order. This determines when it’s possible for any sample
// later in decode order than the current sample to have an earllier
// presentation time than the current sample of the specified cursor. You can
// use this test to limit backward scans, such as to start forward playback.
// For example, with the argument cursor fixed, step the cursor forward until
// it’s impossible for any later-in-decode-order samples to be
// earlier-in-presentation-order than the argument cursor sample.
//
// Don’t implement this method for formats where sample reordering doesn’t
// make sense for the track content, which also indicates that the samples
// aren’t reordered.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/samplesWithLaterDTSsMayHaveEarlierPTSs(than:)
func (o MESampleCursorObject) SamplesWithLaterDTSsMayHaveEarlierPTSsThanCursor(cursor MESampleCursor) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("samplesWithLaterDTSsMayHaveEarlierPTSsThanCursor:"), cursor)
	return rv
}

// Returns an estimate of the sample location indicated by the cursor.
//
// # Return Value
//
// An object that provides information about the estimated sample location.
//
// # Discussion
//
// Some formats may need to read some data on a per-sample basis to produce
// the exact sample location. For these formats, it’s more efficient to read
// a larger chunk of data that contains both the data to produce the exact
// sample location and the actual sample data.
//
// Pass the value this method returns to
// [RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError]
// to obtain the exact sample location.
//
// To indicate that refinement isn’t necessary, return a value for
// [MEEstimatedSampleLocation.RefinementDataLocation] that has a zero length.
// If [MEEstimatedSampleLocation.RefinementDataLocation] has a non-zero
// length, the range for the estimated sample location needs to fully cover
// the refined range that
// [RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError]
// returns, and the refinement data location.
//
// This method fails with the error [MEError.Code.locationNotAvailable] if the
// sample location indicated by the cursor isn’t contiguous or the method
// isn’t supported. In this case, use
// [LoadSampleBufferContainingSamplesToEndCursorCompletionHandler] to load the
// sample data.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/estimatedSampleLocation()
//
// [MEError.Code.locationNotAvailable]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/locationNotAvailable
func (o MESampleCursorObject) EstimatedSampleLocationReturningError() (IMEEstimatedSampleLocation, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("estimatedSampleLocationReturningError:"))
	if err != nil {
		return nil, err
	}
	return MEEstimatedSampleLocationFromID(rv), nil
}

// Produces an exact sample location based on the estimated sample location
// and refinement data that you specify.
//
// estimatedSampleLocation: The estimated sample location.
//
// refinementData: The refinement data.
//
// refinementDataLength: The length of the refinement data in bytes.
//
// refinedLocationOut: The starting file offset and size of the sample in bytes.
//
// # Discussion
//
// Use [EstimatedSampleLocationReturningError] to obtain the estimated sample
// location to pass to this method.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/refineSampleLocation(_:refinementData:refinementDataLength:refinedLocation:)
func (o MESampleCursorObject) RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError(estimatedSampleLocation avfoundation.AVSampleCursorStorageRange, refinementData *uint8, refinementDataLength uintptr, refinedLocationOut *avfoundation.AVSampleCursorStorageRange) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("refineSampleLocation:refinementData:refinementDataLength:refinedLocation:error:"), estimatedSampleLocation, unsafe.Pointer(refinementData), refinementDataLength, unsafe.Pointer(refinedLocationOut))
	if err != nil {
		return false, err
	}
	return rv, nil
}

// Returns information about the chunk that holds the sample indicated by the
// cursor.
//
// # Return Value
//
// A sample cursor chunk.
//
// # Discussion
//
// If the sample resides in a contiguous chunk of the file among similar
// samples, this method returns information about that chunk.
//
// It may not be practical to use this method with some media assets. In this
// case, or if the cursor doesn’t support this method, it returns
// [MEError.Code.locationNotAvailable], which indicates to use
// [LoadSampleBufferContainingSamplesToEndCursorCompletionHandler] to load the
// sample data instead.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/chunkDetails()
//
// [MEError.Code.locationNotAvailable]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/locationNotAvailable
func (o MESampleCursorObject) ChunkDetailsReturningError() (IMESampleCursorChunk, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("chunkDetailsReturningError:"))
	if err != nil {
		return nil, err
	}
	return MESampleCursorChunkFromID(rv), nil
}

// Returns the location and byte source of the sample indicated by the cursor.
//
// # Return Value
//
// A sample location.
//
// # Discussion
//
// Sample data needs to be contiguous. If the sample data isn’t contiguous
// or the cursor doesn’t support this method, it fails with the error
// [MEError.Code.locationNotAvailable]. In this case, use
// [LoadSampleBufferContainingSamplesToEndCursorCompletionHandler] to load the
// data.
//
// If it’s not possible to implement this method, implement
// [EstimatedSampleLocationReturningError] to get an estimated sample
// location, and
// [RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError]
// to analyze this data and provide precise location and size info.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/sampleLocation()
//
// [MEError.Code.locationNotAvailable]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/locationNotAvailable
func (o MESampleCursorObject) SampleLocationReturningError() (IMESampleLocation, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("sampleLocationReturningError:"))
	if err != nil {
		return nil, err
	}
	return MESampleLocationFromID(rv), nil
}

// Builds a sample buffer that contains the samples at the cursor that you
// specify.
//
// endSampleCursor: If not `nil`, this cursor indicates the last sample that the new sample
// buffer should contain.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// # Discussion
//
// Plugin format readers that don’t implement [SampleLocationReturningError]
// or that always load sample data to answer cursor queries need to implement
// this method. If a plug-in format reader implements
// [SampleLocationReturningError], implementing
// [LoadSampleBufferContainingSamplesToEndCursorCompletionHandler] is
// optional.
//
// If there’s no sample data between the sample cursor and
// `endSampleCursor`, the sample buffer is empty. If an error occurs, the
// sample buffer is `nil`.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/loadSampleBufferContainingSamples(to:completionHandler:)
func (o MESampleCursorObject) LoadSampleBufferContainingSamplesToEndCursorCompletionHandler(endSampleCursor MESampleCursor, completionHandler CMSampleBufferRefErrorHandler) {
	_block1, _ := NewCMSampleBufferRefErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadSampleBufferContainingSamplesToEndCursor:completionHandler:"), endSampleCursor, _block1)
}

// The presentation timestamp (PTS) of the sample at the current position of
// the cursor.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/presentationTimeStamp
func (o MESampleCursorObject) PresentationTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](o.ID, objc.Sel("presentationTimeStamp"))
	return coremedia.CMTime(rv)
}

// The decode timestamp (DTS) of the sample at the current position of the
// cursor.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/decodeTimeStamp
func (o MESampleCursorObject) DecodeTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](o.ID, objc.Sel("decodeTimeStamp"))
	return coremedia.CMTime(rv)
}

// The decode duration of the sample at the current position.
//
// # Discussion
//
// This value is [indefinite] if the system needs to advance the sample past
// its current position to determine the decode duration. This can occur with
// streaming formats such as MPEG-2 transport streams.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/currentSampleDuration
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
func (o MESampleCursorObject) CurrentSampleDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](o.ID, objc.Sel("currentSampleDuration"))
	return coremedia.CMTime(rv)
}

// The format description for the sample at the current position of the
// cursor.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/currentSampleFormatDescription
func (o MESampleCursorObject) CurrentSampleFormatDescription() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](o.ID, objc.Sel("currentSampleFormatDescription"))
	return coremedia.CMFormatDescriptionRef(rv)
}

// Decoder synchronization information about the sample the cursor points to.
//
// # Discussion
//
// This value includes any valid flags set. Don’t implement this property if
// this kind of synchronization information doesn’t make sense for the
// sequence of samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/syncInfo
func (o MESampleCursorObject) SyncInfo() avfoundation.AVSampleCursorSyncInfo {
	rv := objc.Send[avfoundation.AVSampleCursorSyncInfo](o.ID, objc.Sel("syncInfo"))
	return avfoundation.AVSampleCursorSyncInfo(rv)
}

// Dependency information about the sample the cursor points to.
//
// # Discussion
//
// This value includes any valid flags set. Don’t implement this property if
// this kind of dependency information doesn’t make sense for the sequence
// of samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/dependencyInfo
func (o MESampleCursorObject) DependencyInfo() avfoundation.AVSampleCursorDependencyInfo {
	rv := objc.Send[avfoundation.AVSampleCursorDependencyInfo](o.ID, objc.Sel("dependencyInfo"))
	return avfoundation.AVSampleCursorDependencyInfo(rv)
}

// Additional information that’s necessary to recover complete sample
// dependency information.
//
// # Discussion
//
// This is an optional property that provides additional sample dependency
// information that [SyncInfo] and [DependencyInfo] don’t provide. Examples
// of this are the NAL unit type of an HEVC sync sample or the number of
// samples necessary to refresh the decoder after a USAC independent frame.
// Don’t implement this property for formats where this information
// doesn’t make sense.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/hevcDependencyInfo
func (o MESampleCursorObject) HevcDependencyInfo() IMEHEVCDependencyInfo {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("hevcDependencyInfo"))
	return MEHEVCDependencyInfoFromID(rv)
}

// The duration of the playable content starting from the cursor position.
//
// # Discussion
//
// Indicates the time difference between the current cursor decode timestamp
// (DTS) and the last reachable sample DTS. This is necessary to play certain
// assets such as those with HTTP URLs, because it indicates what samples the
// byte source has already loaded.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursor/decodeTimeOfLastSampleReachableByForwardSteppingThatIsAlreadyLoadedByByteSource
func (o MESampleCursorObject) DecodeTimeOfLastSampleReachableByForwardSteppingThatIsAlreadyLoadedByByteSource() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](o.ID, objc.Sel("decodeTimeOfLastSampleReachableByForwardSteppingThatIsAlreadyLoadedByByteSource"))
	return coremedia.CMTime(rv)
}
