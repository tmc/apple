// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAsynchronousVideoCompositionRequest] class.
var (
	_AVAsynchronousVideoCompositionRequestClass     AVAsynchronousVideoCompositionRequestClass
	_AVAsynchronousVideoCompositionRequestClassOnce sync.Once
)

func getAVAsynchronousVideoCompositionRequestClass() AVAsynchronousVideoCompositionRequestClass {
	_AVAsynchronousVideoCompositionRequestClassOnce.Do(func() {
		_AVAsynchronousVideoCompositionRequestClass = AVAsynchronousVideoCompositionRequestClass{class: objc.GetClass("AVAsynchronousVideoCompositionRequest")}
	})
	return _AVAsynchronousVideoCompositionRequestClass
}

// GetAVAsynchronousVideoCompositionRequestClass returns the class object for AVAsynchronousVideoCompositionRequest.
func GetAVAsynchronousVideoCompositionRequestClass() AVAsynchronousVideoCompositionRequestClass {
	return getAVAsynchronousVideoCompositionRequestClass()
}

type AVAsynchronousVideoCompositionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAsynchronousVideoCompositionRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAsynchronousVideoCompositionRequestClass) Alloc() AVAsynchronousVideoCompositionRequest {
	rv := objc.Send[AVAsynchronousVideoCompositionRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that contains information a video compositor needs to render an
// output pixel buffer.
//
// # Overview
// 
// The video compositor must adopt the [AVVideoCompositing] protocol.
//
// # Inspecting the request
//
//   - [AVAsynchronousVideoCompositionRequest.CompositionTime]: A time for which to compose the frame.
//   - [AVAsynchronousVideoCompositionRequest.RenderContext]: The rendering context of the video composition.
//   - [AVAsynchronousVideoCompositionRequest.VideoCompositionInstruction]: A video composition instruction that indicates how to compose the frame.
//
// # Accessing source data
//
//   - [AVAsynchronousVideoCompositionRequest.SourceTimedMetadataByTrackID]: Returns a source timed metadata group for the track that contains the specified identifier.
//   - [AVAsynchronousVideoCompositionRequest.SourceTrackIDs]: The identifiers of tracks that contain source video.
//
// # Finishing the request
//
//   - [AVAsynchronousVideoCompositionRequest.FinishWithError]: Finishes the request with an error.
//   - [AVAsynchronousVideoCompositionRequest.FinishCancelledRequest]: Cancels the request to compose a video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest
type AVAsynchronousVideoCompositionRequest struct {
	objectivec.Object
}

// AVAsynchronousVideoCompositionRequestFromID constructs a [AVAsynchronousVideoCompositionRequest] from an objc.ID.
//
// An object that contains information a video compositor needs to render an
// output pixel buffer.
func AVAsynchronousVideoCompositionRequestFromID(id objc.ID) AVAsynchronousVideoCompositionRequest {
	return AVAsynchronousVideoCompositionRequest{objectivec.Object{ID: id}}
}
// NOTE: AVAsynchronousVideoCompositionRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAsynchronousVideoCompositionRequest] class.
//
// # Inspecting the request
//
//   - [IAVAsynchronousVideoCompositionRequest.CompositionTime]: A time for which to compose the frame.
//   - [IAVAsynchronousVideoCompositionRequest.RenderContext]: The rendering context of the video composition.
//   - [IAVAsynchronousVideoCompositionRequest.VideoCompositionInstruction]: A video composition instruction that indicates how to compose the frame.
//
// # Accessing source data
//
//   - [IAVAsynchronousVideoCompositionRequest.SourceTimedMetadataByTrackID]: Returns a source timed metadata group for the track that contains the specified identifier.
//   - [IAVAsynchronousVideoCompositionRequest.SourceTrackIDs]: The identifiers of tracks that contain source video.
//
// # Finishing the request
//
//   - [IAVAsynchronousVideoCompositionRequest.FinishWithError]: Finishes the request with an error.
//   - [IAVAsynchronousVideoCompositionRequest.FinishCancelledRequest]: Cancels the request to compose a video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest
type IAVAsynchronousVideoCompositionRequest interface {
	objectivec.IObject

	// Topic: Inspecting the request

	// A time for which to compose the frame.
	CompositionTime() coremedia.CMTime
	// The rendering context of the video composition.
	RenderContext() IAVVideoCompositionRenderContext
	// A video composition instruction that indicates how to compose the frame.
	VideoCompositionInstruction() AVVideoCompositionInstruction

	// Topic: Accessing source data

	// Returns a source timed metadata group for the track that contains the specified identifier.
	SourceTimedMetadataByTrackID(trackID int32) IAVTimedMetadataGroup
	// The identifiers of tracks that contain source video.
	SourceTrackIDs() []foundation.NSNumber

	// Topic: Finishing the request

	// Finishes the request with an error.
	FinishWithError(error_ foundation.INSError)
	// Cancels the request to compose a video frame.
	FinishCancelledRequest()

	// The identifiers of tracks that contain source metadata.
	SourceSampleDataTrackIDs() []foundation.NSNumber
	// Associates the pixel buffer with the specified spatial configuration.
	AttachSpatialVideoConfigurationToPixelBuffer(spatialVideoConfiguration IAVSpatialVideoConfiguration, pixelBuffer corevideo.CVImageBufferRef)
	// The method that the custom compositor calls when composition succeeds.
	FinishWithComposedTaggedBufferGroup(taggedBufferGroup coremedia.CMTaggedBufferGroupRef)
	// Returns the source CMTaggedBufferGroupRef for the given track ID.
	SourceTaggedBufferGroupByTrackID(trackID int32) coremedia.CMTaggedBufferGroupRef
}

// Init initializes the instance.
func (a AVAsynchronousVideoCompositionRequest) Init() AVAsynchronousVideoCompositionRequest {
	rv := objc.Send[AVAsynchronousVideoCompositionRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAsynchronousVideoCompositionRequest) Autorelease() AVAsynchronousVideoCompositionRequest {
	rv := objc.Send[AVAsynchronousVideoCompositionRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAsynchronousVideoCompositionRequest creates a new AVAsynchronousVideoCompositionRequest instance.
func NewAVAsynchronousVideoCompositionRequest() AVAsynchronousVideoCompositionRequest {
	class := getAVAsynchronousVideoCompositionRequestClass()
	rv := objc.Send[AVAsynchronousVideoCompositionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a source timed metadata group for the track that contains the
// specified identifier.
//
// trackID: The identifier of the track that contains the timed metadata.
//
// # Return Value
// 
// A timed metadata group, or `nil` if not found.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTimedMetadata(byTrackID:)
func (a AVAsynchronousVideoCompositionRequest) SourceTimedMetadataByTrackID(trackID int32) IAVTimedMetadataGroup {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sourceTimedMetadataByTrackID:"), trackID)
	return AVTimedMetadataGroupFromID(rv)
}
// Finishes the request with an error.
//
// error: Returns the error encountered during the compositing.
//
// # Discussion
// 
// A custom compositor calls this method to indicate that the attempt to
// compose a frame failed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finish(with:)
func (a AVAsynchronousVideoCompositionRequest) FinishWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishWithError:"), error_)
}
// Cancels the request to compose a video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finishCancelledRequest()
func (a AVAsynchronousVideoCompositionRequest) FinishCancelledRequest() {
	objc.Send[objc.ID](a.ID, objc.Sel("finishCancelledRequest"))
}
// Associates the pixel buffer with the specified spatial configuration.
//
// spatialVideoConfiguration: The spatial configuration to associate with the pixel buffer.
//
// pixelBuffer: The pixel buffer to associate with the spatial configuration. NOTE: The
// spatial configuration must be one of the spatial configurations specified
// in the `AVVideoComposition/spatialConfigurations` property. An exception
// will be thrown otherwise. NOTE: All pixel buffers from the custom
// compositor must be associated with the same spatial configuration. An
// exception will be thrown otherwise. A spatial configuration with all nil
// values indicates the video is not spatial. A nil spatial configuration also
// indicates the video is not spatial. The value can be nil, which indicates
// the output will not be spatial, but a spatial configuration with all nil
// values must be in the `AVVideoComposition/spatialConfigurations` property
// or an exception will be thrown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/attachSpatialVideoConfiguration:toPixelBuffer:
func (a AVAsynchronousVideoCompositionRequest) AttachSpatialVideoConfigurationToPixelBuffer(spatialVideoConfiguration IAVSpatialVideoConfiguration, pixelBuffer corevideo.CVImageBufferRef) {
	objc.Send[objc.ID](a.ID, objc.Sel("attachSpatialVideoConfiguration:toPixelBuffer:"), spatialVideoConfiguration, pixelBuffer)
}
// The method that the custom compositor calls when composition succeeds.
//
// taggedBufferGroup: The tagged buffer group containing the composed tagged buffers. The tagged
// buffers must be compatible with the outputBufferDescription specified in
// the video composition. The outputBufferDescription must not be nil when
// calling this function. NOTE: If `AVVideoComposition/spatialConfigurations`
// is not empty, then `attach()` must be called with one of the spatial
// configurations. An exception will be thrown otherwise. Also, all pixel
// buffers must be associated with the same spatial configuration. An
// exception will be thrown otherwise.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finishWithComposedTaggedBufferGroup:
func (a AVAsynchronousVideoCompositionRequest) FinishWithComposedTaggedBufferGroup(taggedBufferGroup coremedia.CMTaggedBufferGroupRef) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishWithComposedTaggedBufferGroup:"), taggedBufferGroup)
}
// Returns the source CMTaggedBufferGroupRef for the given track ID.
//
// trackID: The track ID for the requested source tagged buffer group.
//
// # Discussion
// 
// Returns nil if the video track does not contain tagged buffers. Returns nil
// if the track does not contain video. This function should only be called
// when supportsSourceTaggedBuffers is YES.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTaggedBufferGroupByTrackID:
func (a AVAsynchronousVideoCompositionRequest) SourceTaggedBufferGroupByTrackID(trackID int32) coremedia.CMTaggedBufferGroupRef {
	rv := objc.Send[coremedia.CMTaggedBufferGroupRef](a.ID, objc.Sel("sourceTaggedBufferGroupByTrackID:"), trackID)
	return coremedia.CMTaggedBufferGroupRef(rv)
}

// A time for which to compose the frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/compositionTime
func (a AVAsynchronousVideoCompositionRequest) CompositionTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("compositionTime"))
	return coremedia.CMTime(rv)
}
// The rendering context of the video composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/renderContext
func (a AVAsynchronousVideoCompositionRequest) RenderContext() IAVVideoCompositionRenderContext {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renderContext"))
	return AVVideoCompositionRenderContextFromID(objc.ID(rv))
}
// A video composition instruction that indicates how to compose the frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/videoCompositionInstruction
func (a AVAsynchronousVideoCompositionRequest) VideoCompositionInstruction() AVVideoCompositionInstruction {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoCompositionInstruction"))
	return AVVideoCompositionInstructionFromID(objc.ID(rv))
}
// The identifiers of tracks that contain source video.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTrackIDs
func (a AVAsynchronousVideoCompositionRequest) SourceTrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sourceTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The identifiers of tracks that contain source metadata.
//
// # Discussion
// 
// The track identifiers are of type [kCMMediaType_Metadata].
//
// [kCMMediaType_Metadata]: https://developer.apple.com/documentation/CoreMedia/kCMMediaType_Metadata
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceSampleDataTrackIDs-9vxz5
func (a AVAsynchronousVideoCompositionRequest) SourceSampleDataTrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

