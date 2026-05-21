// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

// Package coreaudiotypes provides Go bindings for the CoreAudioTypes framework.
//
// Use specialized data types to interact with audio streams, complex buffers,
// and audiovisual timestamps.
//
// The Core Audio Types framework declares common data types and constants
// that other Core Audio interfaces use. This framework also includes several
// convenience functions.
//
// # Buffers
//
//   - [AudioBuffer]: A structure that holds a buffer of audio data.
//   - [AudioBufferList]: A structure that stores a variable-length array of audio buffers.
//
// # Channels
//
//   - [AudioChannelDescription]: A structure that describes a channel of audio data. ([AudioChannelLabel], [AudioChannelFlags])
//   - [AudioChannelLayout]: A structure that specifies a channel layout in a file or in hardware. ([AudioChannelBitmap], [AudioChannelLayoutTag])
//
// # Codecs
//
//   - [AudioClassDescription]: A structure that describes an audio codec.
//
// # Audio Time
//
//   - [AudioTimeStamp]: A structure that represents a timestamp value.
//   - [AudioTimeStampFlags]: A structure that represents flags for a timestamp.
//
// # SMPTE Time
//
//   - [SMPTETime]: A structure that defines an SMPTE time value.
//   - [SMPTETimeFlags]: A structure that defines SMPTE time flags.
//   - [SMPTETimeType]: Constants that define SMPTE time types.
//
// # Values
//
//   - [AudioValueRange]: A structure that represents a continuous range of values.
//   - [AudioValueTranslation]: A structure that stores buffers to use in translation operations.
//
// # Streams
//
//   - [AudioStreamBasicDescription]: A format specification for an audio stream.
//   - [AudioStreamPacketDescription]: A value that describes a packet in a buffer of audio data.
//   - [AudioFormatFlags]: A type definition for audio format flags.
//   - [AudioFormatID]: A type definition for audio format identifiers.
//   - [KAudioStreamAnyRate]: A value that indicates that an audio stream can use any sample rate.
//
// # Common Types
//
//   - [AVAudioInteger]: An integer type for audio operations.
//   - [AVAudioUInteger]: An unsigned integer type for audio operations.
//   - [AudioSessionID]: A unique identifier of an audio session.
//   - kAudioUnitSampleFractionBits: The number of fractional bits in fixed-point samples.
//   - COREAUDIOTYPES_VERSION: A value that represents the Core Audio Types version.
//   - [AudioFormatListItem]
//
// # Errors
//
//   - kAudio_ParamError: An error in the parameter list of the function.
//   - kAudio_MemFullError: An error that indicates that the heap zone is full.
//   - kAudio_FileNotFoundError: An error that indicates the file wasn’t found.
//   - kAudio_UnimplementedError: An error that indicates the app called an unimplemented system function.
//
// # Macros
//
//   - CA_CANONICAL_DEPRECATED
//   - CA_REALTIME_API
//   - TestAudioFormatNativeEndian
//
// # Enumeration Cases
//
//   - kAudioChannelLayoutTag_Ogg_3_0
//   - kAudioChannelLayoutTag_Ogg_4_0
//   - kAudioChannelLayoutTag_Ogg_5_0
//   - kAudioChannelLayoutTag_Ogg_5_1
//   - kAudioChannelLayoutTag_Ogg_6_1
//   - kAudioChannelLayoutTag_Ogg_7_1
//   - kAudioFormatAPAC
//   - kAudio_BadFilePathError
//   - kAudio_FilePermissionError
//   - kAudio_NoError
//   - kAudio_TooManyFilesOpenError
//
// # Enumerations
//
//   - [AVAudioSessionErrorCode]: Codes that describe error conditions that may occur when performing audio session operations.
package coreaudiotypes

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreAudioTypes library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreAudioTypes.framework/CoreAudioTypes",
	"/usr/lib/libCoreAudioTypes.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: CoreAudioTypes: failed to load framework from any known path\n")
}
