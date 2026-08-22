// Code generated from Apple documentation. DO NOT EDIT.

package videotoolbox

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/kernel"
)

// VTCompressionOutputCallback is a callback for the system to invoke when it’s finished compressing a frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionOutputCallback
type VTCompressionOutputCallback = func(outputCallbackRefCon unsafe.Pointer, sourceFrameRefCon unsafe.Pointer, status int32, infoFlags uint, sampleBuffer uintptr)

// VTCompressionOutputHandler is a callback for the system to invoke when it’s finished compressing a frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionOutputHandler
type VTCompressionOutputHandler = func(status int32, infoFlags uint, sampleBuffer unsafe.Pointer)

// VTCompressionSessionRef is a reference to a VideoToolbox compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSession
type VTCompressionSessionRef uintptr

// VTDecompressionMultiImageCapableOutputHandler is a type alias for callback that the system invokes when it finishes decompressing a frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionMultiImageCapableOutputHandler
type VTDecompressionMultiImageCapableOutputHandler = func(status int32, infoFlags uint, imageBuffer corevideo.CVImageBufferRef, taggedBufferGroup *uintptr, presentationTimeStamp coremedia.CMTime, presentationDuration coremedia.CMTime)

// VTDecompressionOutputCallback is the prototype for the callback invoked when frame decompression is complete.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionOutputCallback
type VTDecompressionOutputCallback = func(decompressionOutputRefCon unsafe.Pointer, sourceFrameRefCon unsafe.Pointer, status int32, infoFlags uint, imageBuffer uintptr, presentationTimeStamp coremedia.CMTime, presentationDuration coremedia.CMTime)

// VTDecompressionOutputHandler is the prototype for the block invoked when frame decompression is complete.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionOutputHandler
type VTDecompressionOutputHandler = func(status int32, infoFlags uint, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, presentationDuration coremedia.CMTime)

// VTDecompressionOutputMultiImageCallback is a callback that the system invokes when multi-image frame decompression completes.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionOutputMultiImageCallback
type VTDecompressionOutputMultiImageCallback = func(decompressionOutputMultiImageRefCon unsafe.Pointer, sourceFrameRefCon unsafe.Pointer, status int32, infoFlags uint, taggedBufferGroup uintptr, presentationTimeStamp coremedia.CMTime, presentationDuration coremedia.CMTime)

// VTDecompressionSessionRef is a reference to a decompression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSession
type VTDecompressionSessionRef uintptr

// VTExtensionPropertiesKey is a key in a Media Extension extension properties dictionary.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey
type VTExtensionPropertiesKey = corefoundation.CFStringRef

// VTFrameSiloRef is an object that stores a large number of sample buffers, as produced by a multipass compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSilo
type VTFrameSiloRef uintptr

// VTHDRPerFrameMetadataGenerationHDRFormatType is the HDR format type.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTHDRPerFrameMetadataGenerationHDRFormatType
type VTHDRPerFrameMetadataGenerationHDRFormatType = corefoundation.CFStringRef

// VTHDRPerFrameMetadataGenerationSessionRef is a mechanism for generating HDR Per Frame Metadata and attaching that metadata to a pixel buffer and the backing IOSurface.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTHDRPerFrameMetadataGenerationSessionRef
type VTHDRPerFrameMetadataGenerationSessionRef uintptr

// VTMotionEstimationOutputHandler is a block invoked by motion-estimation session when frame processing is complete.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationOutputHandler
type VTMotionEstimationOutputHandler = func(status int32, infoFlags uint, additionalInfo corefoundation.CFDictionaryRef, motionVectorPixelBuffer corevideo.CVImageBufferRef)

// VTMotionEstimationSessionRef is a reference to a Video Toolbox motion-estimation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionRef
type VTMotionEstimationSessionRef uintptr

// VTMultiPassStorageRef is an object for storing information for each frame of a multipass compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMultiPassStorage
type VTMultiPassStorageRef uintptr

// VTPixelRotationSessionRef is a reference to a pixel rotation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelRotationSession
type VTPixelRotationSessionRef uintptr

// VTPixelTransferSessionRef is a reference to a VideoToolbox pixel transfer session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelTransferSession
type VTPixelTransferSessionRef = kernel.Pointer

// VTRAWProcessingOutputHandler is a block the system calls when frame processing is complete.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingOutputHandler
type VTRAWProcessingOutputHandler = func(int32, corevideo.CVImageBufferRef)

// VTRAWProcessingParameterChangeHandler is a function the system calls when processing parameters change.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingParameterChangeHandler
type VTRAWProcessingParameterChangeHandler = func(newParameters corefoundation.CFArrayRef)

// VTRAWProcessingSessionRef is an object that processes frames in camera native formats such as RAW or Bayer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSession
type VTRAWProcessingSessionRef uintptr

// VTSessionRef is a reference to a VideoToolbox compression session, decompression session or pixel transfer session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSession
type VTSessionRef = corefoundation.CFTypeRef
