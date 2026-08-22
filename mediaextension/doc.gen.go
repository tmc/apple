// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

// Package mediaextension provides Go bindings for the MediaExtension framework.
//
// This framework provides a means for developers to create format readers,
// video decoders, and RAW processors for media that the system doesn’t
// natively support.
//
// MediaExtension format readers encapsulate media assets that the system
// doesn’t natively support so that the system can recognize them.
// MediaExtension video decoders decode video formats that the system
// doesn’t natively support. MediaExtension RAW processors work together
// with video decoders to allow direct control over the RAW decoding process.
// Developers need to build format readers, video decoders, and RAW processors
// as [ExtensionKit] bundles and embed them in a host app. Once a user
// installs and runs the host app, the embedded extensions become available to
// any app on the user’s system that opts in to using them.
//
// # Format readers
//
//   - [MEFormatReader]: A protocol that defines the requirements for a format reader, which represents a single media asset. ([MEFormatReaderParseAdditionalFragmentsStatus])
//   - [MEFormatReaderExtension]: A protocol that defines a factory to create a new format reader with a byte source.
//   - [MEFormatReaderInstantiationOptions]: An object that contains options to pass to a format reader extension.
//   - [MEFileInfo]: An object that contains file properties from the media asset.
//   - [Format reader property list dictionaries]: Include property list dictionaries to describe a format reader and register the formats it supports.
//   - [Format reader entitlement]: Include an entitlement to indicate your extension is a MediaExtension format reader.
//
// # Track readers
//
//   - [METrackReader]: A protocol that defines the information to provide about a track within a media asset.
//   - [METrackInfo]: An object that includes track properties parsed from the media asset.
//
// # Sample cursors
//
//   - [MESampleCursor]: A protocol that defines the information to provide about samples within a track of a media asset, and enables stepping through samples in the track in decode or presentation order.
//   - [MESampleLocation]: An object that provides information about the sample location with the media.
//   - [MESampleCursorChunk]: An object that provides information about the chunk of media at the location of a sample.
//   - [MEEstimatedSampleLocation]: An object that provides information about the estimated sample location with the media.
//   - [MEHEVCDependencyInfo]: An object that provides information about the HEVC dependency attributes of a sample.
//
// # Byte sources
//
//   - [MEByteSource]: Provides read access to the data in a media asset file.
//
// # Video decoders
//
//   - [MEVideoDecoder]: A protocol that defines the requirements for a video decoder. ([MEDecodeFrameStatus])
//   - [MEVideoDecoderExtension]: A protocol that defines a factory to create new video decoders for a codec type that the extension implements.
//   - [MEDecodeFrameOptions]: An object that guides the video decoder operation on a per-frame basis.
//   - [MEVideoDecoderPixelBufferManager]: Describes pixel buffer requirements and creates new pixel buffers.
//   - [Video decoder property list dictionary]: Include a property list dictionary to describe a video decoder.
//   - [Video decoder entitlement]: Include an entitlement to indicate your extension is a MediaExtension video decoder.
//
// # RAW processors
//
//   - [MERAWProcessor]: A protocol that defines the requirements for a RAW processor.
//   - [MERAWProcessorExtension]: A protocol that defines a factory to create RAW processors for a codec type that the extension implements.
//   - [MERAWProcessorPixelBufferManager]: Describes pixel buffer requirements and creates new pixel buffers.
//   - [MERAWProcessingParameter]: An object for the RAW processor to describe each processing parameter the processor exposes.
//   - [RAW processor property list dictionary]: Include a property list dictionary to describe a RAW processor.
//   - [RAW processor entitlement]: Include an entitlement to indicate your extension is a MediaExtension RAW processor.
//
// # Errors
//
//   - [MediaExtensionErrorDomain]: The domain of the error.
//   - [MEError]: An enumeration that models media extension error codes.//
//
// # Key Types
//
//   - [MEHEVCDependencyInfo] - An object that provides information about the HEVC dependency attributes of a sample.
//   - [METrackInfo] - An object that includes track properties parsed from the media asset.
//   - [MEByteSource] - Provides read access to the data in a media asset file.
//   - [MESampleCursorChunk] - An object that provides information about the chunk of media at the location of a sample.
//   - [MEEstimatedSampleLocation] - An object that provides information about the estimated sample location with the media.
//   - [MERAWProcessingFloatParameter] - An object that describes a floating-point parameter of a RAW processor.
//   - [MERAWProcessingIntegerParameter] - An object that describes an integer parameter of a RAW processor.
//   - [MERAWProcessingParameter] - An object for the RAW processor to describe each processing parameter the processor exposes.
//   - [MEFileInfo] - An object that contains file properties from the media asset.
//   - [MERAWProcessingListParameter] - An object that describes a list parameter of a RAW processor.
//
// [Format reader entitlement]: https://developer.apple.com/documentation/mediaextension/format-reader-entitlement
// [Format reader property list dictionaries]: https://developer.apple.com/documentation/mediaextension/format-reader-property-list-dictionaries
// [RAW processor entitlement]: https://developer.apple.com/documentation/mediaextension/raw-processor-entitlement
// [RAW processor property list dictionary]: https://developer.apple.com/documentation/mediaextension/raw-processor-property-list-dictionary
// [Video decoder entitlement]: https://developer.apple.com/documentation/mediaextension/video-decoder-entitlement
// [Video decoder property list dictionary]: https://developer.apple.com/documentation/mediaextension/video-decoder-property-list-dictionary
package mediaextension

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the MediaExtension library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/MediaExtension.framework/MediaExtension",
	"/usr/lib/libMediaExtension.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: MediaExtension: failed to load framework from any known path\n")
	}
}
