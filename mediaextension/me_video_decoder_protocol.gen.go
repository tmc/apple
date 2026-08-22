// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the requirements for a video decoder.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder
type MEVideoDecoder interface {
	objectivec.IObject

	// Requests the extension to decode a video frame.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/decodeFrame(from:options:completionHandler:)
	DecodeFrameFromSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.CMSampleBufferRef, options IMEDecodeFrameOptions, completionHandler CVImageBufferRefMEDecodeFrameStatusErrorHandler)

	// A Boolean that specifies whether the content has interframe dependencies, if the decoder knows.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/contentHasInterframeDependencies
	ContentHasInterframeDependencies() bool

	// The recommended number of threads for the decoder to use.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/recommendedThreadCount
	RecommendedThreadCount() int
	SetRecommendedThreadCount(value int)

	// The actual number of threads the decoder uses.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/actualThreadCount
	ActualThreadCount() int

	// Provides hints about quality tradeoffs between pixel formats.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/supportedPixelFormatsOrderedByQuality
	SupportedPixelFormatsOrderedByQuality() []foundation.NSNumber

	// A request to decode at a lower resolution than full-size.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/reducedResolution
	ReducedResolution() corefoundation.CGSize
	SetReducedResolution(value corefoundation.CGSize)

	// Provides a list of output pixel formats where the decoder supports reduced resolution decoding.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/pixelFormatsWithReducedResolutionDecodeSupport
	PixelFormatsWithReducedResolutionDecodeSupport() []foundation.NSNumber

	// Indicates whether the decoder produces RAW output which requires the use of a RAW processor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/producesRAWOutput
	ProducesRAWOutput() bool

	// A Boolean value that indicates the readiness of the decoder to accept more sample buffers.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/isReadyForMoreMediaData
	IsReadyForMoreMediaData() bool
}

// MEVideoDecoderObject wraps an existing Objective-C object that conforms to the MEVideoDecoder protocol.
type MEVideoDecoderObject struct {
	objectivec.Object
}

func (o MEVideoDecoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MEVideoDecoderObjectFromID constructs a [MEVideoDecoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MEVideoDecoderObjectFromID(id objc.ID) MEVideoDecoderObject {
	return MEVideoDecoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests the extension to decode a video frame.
//
// sampleBuffer: A sample buffer that contains one video frame.
//
// options: Specific decode options for the frame.
//
// completionHandler: The completion block to execute when the decode operation finishes.
//
// # Discussion
//
// This method calls the completion handler for every sample buffer frame when
// decoding completes, but not necessarily in display order. The completion
// handler receives a decoded pixel buffer, a decode status that indicates a
// frame dropped, or an error. Use [MEVideoDecoderPixelBufferManager] to
// allocate an image buffer. If an error occurs that’s external to
// [MediaExtensionErrorDomain], the [VTDecompressionSession] receives it as
// [kVTVideoDecoderUnknownErr].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/decodeFrame(from:options:completionHandler:)
//
// [MediaExtensionErrorDomain]: https://developer.apple.com/documentation/MediaExtension/MediaExtensionErrorDomain
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [kVTVideoDecoderUnknownErr]: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderUnknownErr
func (o MEVideoDecoderObject) DecodeFrameFromSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.CMSampleBufferRef, options IMEDecodeFrameOptions, completionHandler CVImageBufferRefMEDecodeFrameStatusErrorHandler) {
	_block2, _ := NewCVImageBufferRefMEDecodeFrameStatusErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("decodeFrameFromSampleBuffer:options:completionHandler:"), sampleBuffer, options, _block2)
}

// Asks the extension whether the decoder can decode frames with the format
// description that you specify.
//
// formatDescription: The new format description to evaluate.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/canAccept(_:)
func (o MEVideoDecoderObject) CanAcceptFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canAcceptFormatDescription:"), formatDescription)
	return rv
}

// A Boolean that specifies whether the content has interframe dependencies,
// if the decoder knows.
//
// # Discussion
//
// The system queries this property on the extension when [Video Toolbox]
// queries the [kVTDecompressionPropertyKey_ContentHasInterframeDependencies]
// on the hosting [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/contentHasInterframeDependencies
//
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
// [kVTDecompressionPropertyKey_ContentHasInterframeDependencies]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ContentHasInterframeDependencies
func (o MEVideoDecoderObject) ContentHasInterframeDependencies() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("contentHasInterframeDependencies"))
	return bool(rv)
}

// The recommended number of threads for the decoder to use.
//
// # Discussion
//
// The system sets this property on the extension when [Video Toolbox] sets
// [kVTDecompressionPropertyKey_ThreadCount] on the hosting
// [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/recommendedThreadCount
//
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
// [kVTDecompressionPropertyKey_ThreadCount]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ThreadCount
func (o MEVideoDecoderObject) RecommendedThreadCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("recommendedThreadCount"))
	return int(rv)
}

func (o MEVideoDecoderObject) SetRecommendedThreadCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setRecommendedThreadCount:"), value)
}

// The actual number of threads the decoder uses.
//
// # Discussion
//
// The system queries this property on the extension when [Video Toolbox]
// queries the [kVTDecompressionPropertyKey_ThreadCount] on the hosting
// [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/actualThreadCount
//
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
// [kVTDecompressionPropertyKey_ThreadCount]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ThreadCount
func (o MEVideoDecoderObject) ActualThreadCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("actualThreadCount"))
	return int(rv)
}

// Provides hints about quality tradeoffs between pixel formats.
//
// # Discussion
//
// This array contains [CMPixelFormatType] values, ordered by quality from
// best to worst. The system queries this property on the extension when
// [Video Toolbox] queries the
// [kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality] on the
// hosting [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/supportedPixelFormatsOrderedByQuality
//
// [CMPixelFormatType]: https://developer.apple.com/documentation/CoreMedia/CMPixelFormatType
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
// [kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality
func (o MEVideoDecoderObject) SupportedPixelFormatsOrderedByQuality() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("supportedPixelFormatsOrderedByQuality"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

// A request to decode at a lower resolution than full-size.
//
// # Discussion
//
// This optional property conveys a request for reduced resolution for
// decoding. Decoders that only support a fixed set of resolutions should pick
// the smallest resolution greater than or equal to the requested width and
// height. If the output [CVPixelBuffer] is not in a format where reduced
// resolution decoding is supported, this setting should be disregarded. This
// property is set on the extension when a Video Toolbox client sets the
// [kVTDecompressionPropertyKey_ReducedResolutionDecode] property on the
// hosting [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/reducedResolution
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [kVTDecompressionPropertyKey_ReducedResolutionDecode]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ReducedResolutionDecode
func (o MEVideoDecoderObject) ReducedResolution() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("reducedResolution"))
	return corefoundation.CGSize(rv)
}

func (o MEVideoDecoderObject) SetReducedResolution(value corefoundation.CGSize) {
	objc.Send[struct{}](o.ID, objc.Sel("setReducedResolution:"), value)
}

// Provides a list of output pixel formats where the decoder supports reduced
// resolution decoding.
//
// # Discussion
//
// This array contains [CMPixelFormatType] values. The system queries this
// property on the extension when [Video Toolbox] queries the
// [kVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport] on
// the hosting [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/pixelFormatsWithReducedResolutionDecodeSupport
//
// [CMPixelFormatType]: https://developer.apple.com/documentation/CoreMedia/CMPixelFormatType
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
// [kVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport
func (o MEVideoDecoderObject) PixelFormatsWithReducedResolutionDecodeSupport() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("pixelFormatsWithReducedResolutionDecodeSupport"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

// Indicates whether the decoder produces RAW output which requires the use of
// a RAW processor.
//
// # Discussion
//
// The extension should implement this property returning YES if the decoder
// produces RAW output which requires the use of an [MERAWProcessor] for
// post-decode processing to produce renderable output.
//
// This optional property is queried on the extension when a Video Toolbox
// client queries the [kVTDecompressionPropertyKey_DecoderProducesRAWOutput]
// property on the hosting [VTDecompressionSession].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/producesRAWOutput
//
// [VTDecompressionSession]: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
// [kVTDecompressionPropertyKey_DecoderProducesRAWOutput]: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_DecoderProducesRAWOutput
func (o MEVideoDecoderObject) ProducesRAWOutput() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("producesRAWOutput"))
	return bool(rv)
}

// A Boolean value that indicates the readiness of the decoder to accept more
// sample buffers.
//
// # Discussion
//
// Video decoders which operate asynchronously often have a fixed capacity for
// buffers in flight in the decoder. This property allows the decoder to
// signal to [Video Toolbox] that its internal buffers are full and it can’t
// accept more samples. The decoder needs to use
// [MEVideoDecoderReadyForMoreMediaDataDidChangeNotification] to notify Video
// Toolbox when this property changes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoder/isReadyForMoreMediaData
//
// [MEVideoDecoderReadyForMoreMediaDataDidChangeNotification]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderReadyForMoreMediaDataDidChangeNotification
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
func (o MEVideoDecoderObject) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReadyForMoreMediaData"))
	return bool(rv)
}
