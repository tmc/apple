// Code generated from Apple documentation. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
)

// AVAudio3DVector is a structure that represents a vector in 3D space, in degrees.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DVector
type AVAudio3DVector = AVAudio3DPoint

// AVAudioChannelCount is the number of audio channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelCount
type AVAudioChannelCount = uint32

// AVAudioConverterInputBlock is a block to get input data for conversion, as necessary.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputBlock
type AVAudioConverterInputBlock = func(uint32, *AVAudioConverterInputStatus) AVAudioBuffer

// AVAudioEngineManualRenderingBlock is the type that represents a block that renders the engine when operating in manual rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingBlock
type AVAudioEngineManualRenderingBlock = func(uint32, unsafe.Pointer, *int) AVAudioEngineManualRenderingStatus

// AVAudioFrameCount is a number of audio sample frames.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFrameCount
type AVAudioFrameCount = uint32

// AVAudioFramePosition is a position in an audio file or stream.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFramePosition
type AVAudioFramePosition = int64

// AVAudioIONodeInputBlock is the type that represents a block to render operation calls to get input data when in manual rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONodeInputBlock
type AVAudioIONodeInputBlock = func(uint32) unsafe.Pointer

// AVAudioNodeBus is the index of a bus on an audio node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNodeBus
type AVAudioNodeBus = uint

// AVAudioNodeCompletionHandler is a general callback handler for an audio node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNodeCompletionHandler
type AVAudioNodeCompletionHandler = func()

// AVAudioNodeTapBlock is the block that receives copies of the output of an audio node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNodeTapBlock
type AVAudioNodeTapBlock = func(AVAudioPCMBuffer, AVAudioTime)

// AVAudioPacketCount is the number of packets of audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPacketCount
type AVAudioPacketCount = uint32

// AVAudioPlayerNodeCompletionHandler is the callback handler for buffer or file completion.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionHandler
type AVAudioPlayerNodeCompletionHandler = func(AVAudioPlayerNodeCompletionCallbackType)

// AVAudioSequencerInfoDictionaryKey is constants that defines metadata keys for a sequencer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/InfoDictionaryKey
type AVAudioSequencerInfoDictionaryKey = string

// AVAudioSequencerUserCallback is a callback the sequencer calls asynchronously during playback when it encounters a user event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencerUserCallback
type AVAudioSequencerUserCallback = func(AVMusicTrack, foundation.NSData, float64)

// AVAudioSinkNodeReceiverBlock is a block that receives audio data from an audio sink node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNodeReceiverBlock
type AVAudioSinkNodeReceiverBlock = func(unsafe.Pointer, uint32, unsafe.Pointer) int

// AVAudioSourceNodeRenderBlock is a block that supplies audio data to an audio source node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNodeRenderBlock
type AVAudioSourceNodeRenderBlock = func(*int8, unsafe.Pointer, uint32, unsafe.Pointer) int

// AVMIDIPlayerCompletionHandler is a callback the system invokes when MIDI playback completes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayerCompletionHandler
type AVMIDIPlayerCompletionHandler = func()

// AVMusicEventEnumerationBlock is a type you use to enumerate and remove music events, if necessary.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicEventEnumerationBlock
type AVMusicEventEnumerationBlock = func(AVMusicEvent, []float64, *int8)

// AVMusicTimeStamp is a fractional number of beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTimeStamp
type AVMusicTimeStamp = float64

// AVSpeechSynthesisProviderOutputBlock is a type that represents the method for sending marker information to the host.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderOutputBlock
type AVSpeechSynthesisProviderOutputBlock = func([]AVSpeechSynthesisMarker, AVSpeechSynthesisProviderRequest)

// AVSpeechSynthesizerBufferCallback is a type that defines a callback that receives a buffer of generated speech.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/BufferCallback
type AVSpeechSynthesizerBufferCallback = func(AVAudioBuffer)

// AVSpeechSynthesizerMarkerCallback is a type that defines a callback that receives speech markers.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/MarkerCallback
type AVSpeechSynthesizerMarkerCallback = func([]AVSpeechSynthesisMarker)
