// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

// Package audiotoolbox provides Go bindings for the AudioToolbox framework.
//
// Record or play audio, convert formats, parse audio streams, and configure
// your audio session.
//
// The AudioToolbox framework provides interfaces for recording, playback, and
// stream parsing. In iOS, the framework provides additional interfaces for
// managing audio sessions.
//
// # Essentials
//
//   - [Porting your audio code to Apple silicon]: Eliminate issues in your audio-specific code when running on Apple silicon Mac computers.
//
// # Audio Units
//
//   - [Generating spatial audio from a multichannel audio stream]: Convert 8-channel audio to 2-channel spatial audio by using a spatial mixer audio unit.
//   - [Audio Unit v3 Plug-Ins]: Deliver custom audio effects, instruments, and other audio behaviors using an Audio Unit v3 app extension. ([AUAudioUnit], [AUAudioUnitBus], [AUAudioUnitBusArray], [AUAudioUnitPreset], [AUAudioUnitV2Bridge])
//   - [Audio Components]: Find, load, and configure audio components, such as Audio Units and audio codecs. ([AudioComponent], [AudioComponentInstantiationOptions], [AudioComponentDescription], [AudioComponentInstance], [AudioComponentFlags])
//   - [Audio Unit Properties]: Obtain information about the built-in mixers, equalizers, filters, effects, and other Audio Unit app extensions. ([HostCallbackInfo], [AUSpatialMixerAttenuationCurve], [AUSpatialMixerRenderingFlags], [AUSpatialMixerOutputType], [AUSpatialMixerPointSourceInHeadMode])
//   - [Audio Unit Voice I/O]: Configure system voice processing and respond to speech events. ([AUVoiceIOMutedSpeechActivityEventListener], [AUVoiceIOSpeechActivityEvent], [AUVoiceIOOtherAudioDuckingConfiguration])
//
// # Playback and Recording
//
//   - [Audio Queue Services]: Connect to audio hardware and manage the recording or playback process. ([AudioQueueRef], [AudioQueueInputCallbackBlock], [AudioQueueOutputCallbackBlock], [AudioQueueTimelineRef], [AudioQueueInputCallback])
//   - [Audio Services]: Play short sounds or trigger a vibration effect on iOS devices with the appropriate hardware. ([SystemSoundID], [AudioServicesSystemSoundCompletionProc], [AudioServicesPropertyID])
//   - [Music Player]: Create and play a sequence of tracks, and manage aspects of playback in response to standard events. ([MusicPlayer], [MusicTimeStamp], [MusicEventIterator], [MusicEventType], [ExtendedNoteOnEvent])
//   - [Anchoring sound to a window or volume]: Provide unique app experiences by attaching sounds to windows and volumes in 3D space.
//
// # Audio Files and Formats
//
//   - [Audio Format Services]: Access information about audio formats and codecs. ([AudioBalanceFade], [AudioFormatInfo], [AudioFormatListItem], [AudioPanningInfo], [ExtendedAudioFormatInfo])
//   - [Audio File Services]: Read or write a variety of audio data to or from disk or a memory buffer. ([AudioFile_ReadProc], [AudioFile_WriteProc], [AudioFile_GetSizeProc], [AudioFile_SetSizeProc], [AudioBytePacketTranslationFlags])
//   - [Extended Audio File Services]: Read and write compressed files and linear PCM audio files using a simplified interface. ([ExtAudioFilePacketTableInfoOverride], [ExtAudioFileRef], [ExtAudioFilePropertyID])
//   - [Audio File Stream Services]: Parse streamed audio files as the data arrives on the user’s computer. ([AudioFileStream_PropertyListenerProc], [AudioFileStream_PacketsProc], [AudioFileStreamPropertyID], [AudioFileStreamID])
//   - [Audio File Components]: Get information about audio file formats, and about files containing audio data. ([AudioFileComponent], [AudioFileComponentPropertyID], [AudioFileComponentCreateURLProc], [AudioFileComponentOpenWithCallbacksProc], [AudioFileComponentOpenURLProc])
//   - [Core Audio File Format]: Parse the structure of Core Audio files. ([CAFAudioDescription], [CAFAudioFormatListItem], [CAFChunkHeader], [CAFDataChunk], [CAFFileHeader])
//
// # Utilities
//
//   - [Analyzing audio performance with Instruments]: Ensure a smooth and immersive audio experience in your apps using Audio System Trace.
//   - [Audio Converter Services]: Convert between linear PCM audio formats, and between linear PCM and compressed formats. ([AudioConverterComplexInputDataProc], [AudioConverterInputDataProc], [AudioConverterPrimeInfo], [AudioConverterRef], [AudioConverterPropertyID])
//   - [Audio Session Support]: Describe the properties that you associate with audio sessions and audio routes. ([AudioSessionInterruptionType])
//   - [Audio Toolbox Debugging]: Obtain the internal state of Core Audio objects during the development and debugging of your code.
//   - [Workgroup Management]: Coordinate the activity of custom real-time audio threads with those of the system and other processes.
//   - [Audio Codec]: Translate audio data from one format to another. ([AudioCodecMagicCookieInfo], [AudioCodecPrimeInfo], [AudioCodec], [AudioCodecAppendInputBufferListProc], [AudioCodecAppendInputDataProc])
//   - [Clock Utilities]: Manage time-related information associated with audio playback. ([CAClockRef], [CAClockListenerProc], [CAClockMessage], [CAClockTime], [CAClockTimeFormat])
//
// # Functions
//
//   - [AudioConverterFillComplexBufferRealtimeSafe]
//   - [AudioConverterFillComplexBufferWithPacketDependencies]
//   - [AudioFileWritePacketsWithDependencies]
//
// # Type Aliases
//
//   - [AudioConverterComplexInputDataProcRealtimeSafe]
//
// # Enumerations
//
//   - [AUAudioMixRenderingStyle]//
//
// # Key Types
//
//   - [AUAudioUnit] - A class that defines a host’s interface to an audio unit.
//   - [AUParameter] - An object that represents a single audio unit parameter.
//   - [AUAudioUnitBus] - A class that defines an input or output connection point on an audio unit.
//   - [AUParameterNode] - An object that represents a node in an audio unit’s parameter tree.
//   - [AUAudioUnitBusArray] - A class that defines a container for an audio unit’s input or output busses.
//   - [AUParameterTree] - An object that represents a top-level group node that contains all of an audio unit’s parameters.
//   - [AUAudioUnitPreset] - A class that describes an interface for custom parameter settings provided by the audio unit developer.
//   - [AUAudioUnitV2Bridge] - A class that wraps a version 2 audio unit as version 3 audio unit.
//   - [AUParameterGroup] - A parameter group object represents a group of related audio unit parameters.
//
// [Analyzing audio performance with Instruments]: https://developer.apple.com/documentation/audiotoolbox/analyzing-audio-performance-with-instruments
// [Anchoring sound to a window or volume]: https://developer.apple.com/documentation/audiotoolbox/spatializing-sound-from-a-uiscene
// [Audio Codec]: https://developer.apple.com/documentation/audiotoolbox/audio-codec
// [Audio Components]: https://developer.apple.com/documentation/audiotoolbox/audio-components
// [Audio Converter Services]: https://developer.apple.com/documentation/audiotoolbox/audio-converter-services
// [Audio File Components]: https://developer.apple.com/documentation/audiotoolbox/audio-file-components
// [Audio File Services]: https://developer.apple.com/documentation/audiotoolbox/audio-file-services
// [Audio File Stream Services]: https://developer.apple.com/documentation/audiotoolbox/audio-file-stream-services
// [Audio Format Services]: https://developer.apple.com/documentation/audiotoolbox/audio-format-services
// [Audio Queue Services]: https://developer.apple.com/documentation/audiotoolbox/audio-queue-services
// [Audio Services]: https://developer.apple.com/documentation/audiotoolbox/audio-services
// [Audio Session Support]: https://developer.apple.com/documentation/audiotoolbox/audio-session-support
// [Audio Toolbox Debugging]: https://developer.apple.com/documentation/audiotoolbox/audio-toolbox-debugging
// [Audio Unit Properties]: https://developer.apple.com/documentation/audiotoolbox/audio-unit-properties
// [Audio Unit Voice I/O]: https://developer.apple.com/documentation/audiotoolbox/audio-unit-voice-i-o
// [Audio Unit v3 Plug-Ins]: https://developer.apple.com/documentation/audiotoolbox/audio-unit-v3-plug-ins
// [Clock Utilities]: https://developer.apple.com/documentation/audiotoolbox/clock-utilities
// [Core Audio File Format]: https://developer.apple.com/documentation/audiotoolbox/core-audio-file-format
// [Extended Audio File Services]: https://developer.apple.com/documentation/audiotoolbox/extended-audio-file-services
// [Generating spatial audio from a multichannel audio stream]: https://developer.apple.com/documentation/audiotoolbox/generating-spatial-audio-from-a-multichannel-audio-stream
// [Music Player]: https://developer.apple.com/documentation/audiotoolbox/music-player
// [Porting your audio code to Apple silicon]: https://developer.apple.com/documentation/Apple-Silicon/porting-your-audio-code-to-apple-silicon
// [Workgroup Management]: https://developer.apple.com/documentation/audiotoolbox/workgroup-management
package audiotoolbox

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the AudioToolbox library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/AudioToolbox.framework/AudioToolbox",
	"/usr/lib/libAudioToolbox.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: AudioToolbox: failed to load framework from any known path\n")
	}
}
