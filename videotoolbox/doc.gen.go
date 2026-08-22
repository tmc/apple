// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

// Package videotoolbox provides Go bindings for the VideoToolbox framework.
//
// Work directly with hardware-accelerated video encoding and decoding
// capabilities.
//
// VideoToolbox is a low-level framework that provides direct access to
// hardware encoders and decoders. It provides services for video compression
// and decompression, and for conversion between raster image formats stored
// in CoreVideo pixel buffers. These services are provided in the form of
// session objects (compression, decompression, and pixel transfer), which are
// vended as Core Foundation (CF) types. Apps that don’t need direct access
// to hardware encoders and decoders shouldn’t need to use VideoToolbox
// directly.
//
// # Frame Processing
//
//   - [Frame processing]: An interface for accessing a range of different video-processing features. ([VTFrameProcessor], [VTFrameProcessorConfiguration], [VTFrameProcessorFrame], [VTFrameProcessorParameters], [VTFrameRateConversionConfiguration])
//
// # Motion Estimation
//
//   - VTMotionEstimationSession ([VTMotionEstimationOutputHandler])
//
// # Compression
//
//   - [Encoding video for low-latency conferencing]: Configure a compression session to optimize encoding for video-conferencing apps.
//   - [Encoding video for live streaming]: Configure a compression session to encode video for live streaming.
//   - [Encoding video for offline transcoding]: Configure a compression session to transcode video in offline workflows.
//   - VTCompressionSession: An object that compresses video data. ([VTEncodeInfoFlags])
//   - VTDecompressionSession: An object that decompresses video data. ([VTDecompressionMultiImageCapableOutputHandler], [VTDecodeFrameFlags], [VTDecodeInfoFlags], [VTDecompressionOutputCallback], [VTDecompressionOutputCallbackRecord])
//
// # Media Extension
//
//   - [VTExtensionPropertiesKey]: A key in a Media Extension extension properties dictionary.
//   - [VTCopyVideoDecoderExtensionProperties]: Returns information about the Media Extension video decoder required to decode the specified format.
//
// # HDR Metadata
//
//   - VTHDRPerFrameMetadataGenerationSession: An object that generates per-frame HDR metadata. ([VTHDRPerFrameMetadataGenerationHDRFormatType])
//
// # Codec Support
//
//   - [VTIsHardwareDecodeSupported]: Returns a Boolean value that indicates whether the current system supports hardware decode for the specified codec.
//   - [VTRegisterProfessionalVideoWorkflowVideoEncoders]: Loads encoders appropriate for the client’s professional video workflows.
//   - [VTRegisterProfessionalVideoWorkflowVideoDecoders]: Loads decoders appropriate for the client’s professional video workflows.
//   - [VTRegisterSupplementalVideoDecoderIfAvailable]: Registers a video decoder for the specified codec type, if one exists on the current system.
//   - [VTCopySupportedPropertyDictionaryForEncoder]: Builds a list of supported properties and encoder ID for an encoder.
//   - [VTCopyVideoEncoderList]: Builds a list of available video encoders.
//   - [Video Encoder List Keys]: Dictionary key constants to use to retrieve video encoder information.
//
// # Utilities
//
//   - [VTCreateCGImageFromCVPixelBuffer]: Creates a Core Graphics bitmap image or image mask using the provided pixel buffer.
//
// # Data Types
//
//   - [VTInt32Point]: A structure that represents a 32-bit integer point value.
//   - [VTInt32Size]: A structure that represents a 32-bit integer size value.
//
// # Errors
//
//   - [Error Code Constants]: Constants for Video Toolbox operation error codes.
//
// # Variables
//
//   - [KVTCompressionPropertyKey_ConstantQualityFactor]//
//
// # Key Types
//
//   - [VTSuperResolutionScalerConfiguration] - Configuration that you use to set up the super-resolution processor.
//   - [VTFrameRateConversionConfiguration] - An object that enables the frame rate conversion on a frame processing session.
//   - [VTMotionBlurConfiguration] - A configuration object to enable motion blur on a frame processing session.
//   - [VTOpticalFlowConfiguration] - A configuration object that enables optical flow on a frame processing session.
//   - [VTLowLatencySuperResolutionScalerConfiguration] - An object you use to configure frame processor for low-latency super-resolution scaler processing.
//   - [VTLowLatencyFrameInterpolationConfiguration] - Configuration that you use to program Video Toolbox frame processor for low-latency frame interpolation.
//   - [VTTemporalNoiseFilterConfiguration] - A configuration object to initiate a frame processor and use temporal noise-filter processor.
//   - [VTMotionBlurParameters] - This object contains both input and output parameters necessary to run the motion blur processor on a frame.
//   - [VTFrameRateConversionParameters] - An object that contains the required input and output parameters to run a frame rate conversion processor on a frame.
//   - [VTSuperResolutionScalerParameters] - An object that contains both input and output parameters that the super-resolution processor needs to run on a frame.
//
// [Encoding video for live streaming]: https://developer.apple.com/documentation/videotoolbox/encoding-video-for-live-streaming
// [Encoding video for low-latency conferencing]: https://developer.apple.com/documentation/videotoolbox/encoding-video-for-low-latency-conferencing
// [Encoding video for offline transcoding]: https://developer.apple.com/documentation/videotoolbox/encoding-video-for-offline-transcoding
// [Error Code Constants]: https://developer.apple.com/documentation/videotoolbox/1490398-error-code-constants
// [Frame processing]: https://developer.apple.com/documentation/videotoolbox/frame-processing
// [Video Encoder List Keys]: https://developer.apple.com/documentation/videotoolbox/video-encoder-list-keys
package videotoolbox

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the VideoToolbox library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/VideoToolbox.framework/VideoToolbox",
	"/usr/lib/libVideoToolbox.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: VideoToolbox: failed to load framework from any known path\n")
	}
}
