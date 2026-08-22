// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the requirements for a RAW processor.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor
type MERAWProcessor interface {
	objectivec.IObject

	// Requests the processor to process a video frame.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/processFrame(fromImageBuffer:completionHandler:)
	ProcessFrameFromImageBufferCompletionHandler(inputFrame corevideo.CVImageBufferRef, completionHandler CVPixelBufferRefErrorHandler)

	// Requests the processor use the provided Metal device for processing.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/metalDeviceRegistryID
	MetalDeviceRegistryID() uint64
	SetMetalDeviceRegistryID(value uint64)

	// Returns the color-related Core Video image buffer keys and values that become attachments to the output pixel buffers.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/outputColorAttachments
	OutputColorAttachments() foundation.INSDictionary

	// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/processingParameters
	ProcessingParameters() []MERAWProcessingParameter

	// Indicates the readiness of the processor to accept more sample buffers.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/isReadyForMoreMediaData
	IsReadyForMoreMediaData() bool

	// # Discussion  The metadata returned is a pre-formatted NSData that represents a fully-formed sidecar file, and should be compatible with the MediaExtension FormatReader.  The metadata will contain the initial processing parameters from the sidecar file, along with any adjustments made on the RAW processor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/metadataForSidecarFile
	MetadataForSidecarFile() foundation.NSData
}

// MERAWProcessorObject wraps an existing Objective-C object that conforms to the MERAWProcessor protocol.
type MERAWProcessorObject struct {
	objectivec.Object
}

func (o MERAWProcessorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MERAWProcessorObjectFromID constructs a [MERAWProcessorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MERAWProcessorObjectFromID(id objc.ID) MERAWProcessorObject {
	return MERAWProcessorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests the processor to process a video frame.
//
// inputFrame: A CVPixelBuffer that contains a video frame to process.
//
// completionHandler: The handler is invoked when a frame processes and is ready to be sent back
// to the caller. This block does not need to be called in the order in which
// frames were submitted.
//
// # Discussion
//
// The completionHandler block must be called for every
// [ProcessFrameFromImageBufferCompletionHandler] call when processing is
// complete. The completion handler block should return either a processed
// pixel buffer or an error.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/processFrame(fromImageBuffer:completionHandler:)
func (o MERAWProcessorObject) ProcessFrameFromImageBufferCompletionHandler(inputFrame corevideo.CVImageBufferRef, completionHandler CVPixelBufferRefErrorHandler) {
	_block1, _ := NewCVPixelBufferRefErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("processFrameFromImageBuffer:completionHandler:"), inputFrame, _block1)
}

// Requests the processor use the provided Metal device for processing.
//
// # Discussion
//
// This optional property requests that [MERAWProcessor] use [MTLDevice]
// corresponding to this ID for any Metal-based processing. This is optional
// and doesn’t need to be implemented if the processor does not use Metal.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/metalDeviceRegistryID
func (o MERAWProcessorObject) MetalDeviceRegistryID() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("metalDeviceRegistryID"))
	return uint64(rv)
}

func (o MERAWProcessorObject) SetMetalDeviceRegistryID(value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setMetalDeviceRegistryID:"), value)
}

// Returns the color-related Core Video image buffer keys and values that
// become attachments to the output pixel buffers.
//
// # Discussion
//
// This is an optional property. Only color-related keys from [Image Buffer
// Attachment Keys] are permitted in the returned dictionary.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/outputColorAttachments
//
// [Image Buffer Attachment Keys]: https://developer.apple.com/documentation/CoreVideo/image-buffer-attachment-keys
func (o MERAWProcessorObject) OutputColorAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputColorAttachments"))
	return foundation.NSDictionaryFromID(rv)
}

// Provides a list of processing parameters that can be changed by the client
// of Video Toolbox session to influence processing behavior.
//
// # Discussion
//
// This property value is an array of [MERAWProcessingParameter] objects, each
// describing the parameter and providing an interface where the processing
// parameter value may be modified.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/processingParameters
func (o MERAWProcessorObject) ProcessingParameters() []MERAWProcessingParameter {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("processingParameters"))
	result := make([]MERAWProcessingParameter, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = MERAWProcessingParameterFromID(id)
	}
	return result
}

// Indicates the readiness of the processor to accept more sample buffers.
//
// # Discussion
//
// An [MERAWProcessor] operates asynchronously and often has a fixed capacity
// for buffers in flight in the processor. This property allows the processor
// to signal to [Video Toolbox] that its internal buffers are full and cannot
// accept more samples. The processor must use
// [MERAWProcessorReadyForMoreMediaDataDidChangeNotification] to notify Video
// Toolbox when this property changes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/isReadyForMoreMediaData
//
// [MERAWProcessorReadyForMoreMediaDataDidChangeNotification]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorReadyForMoreMediaDataDidChangeNotification
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
func (o MERAWProcessorObject) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReadyForMoreMediaData"))
	return bool(rv)
}

// # Discussion
//
// The metadata returned is a pre-formatted NSData that represents a
// fully-formed sidecar file, and should be compatible with the MediaExtension
// FormatReader.
//
// The metadata will contain the initial processing parameters from the
// sidecar file, along with any adjustments made on the RAW processor.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessor/metadataForSidecarFile
func (o MERAWProcessorObject) MetadataForSidecarFile() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metadataForSidecarFile"))
	return foundation.NSDataFromID(rv)
}
