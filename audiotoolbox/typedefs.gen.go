// Code generated from Apple documentation. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/coremidi"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
)

// AUAudioChannelCount is a number of audio channels.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioChannelCount
type AUAudioChannelCount = uint32

// AUAudioFrameCount is a number of audio sample frames.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioFrameCount
type AUAudioFrameCount = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioObjectID
type AUAudioObjectID = uint32

// AUAudioUnitStatus is a result code returned from an audio unit’s render function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitStatus
type AUAudioUnitStatus = int32

// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerBlock
type AUEventListenerBlock = func(unsafe.Pointer, *AudioUnitEvent, uint64, float32)

// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerProc
type AUEventListenerProc = func(unsafe.Pointer, unsafe.Pointer, uintptr, uint64, float32)

// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRef
type AUEventListenerRef uintptr

// AUEventSampleTime is expresses time as a sample count.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventSampleTime
type AUEventSampleTime = int64

// AUGraph is an opaque type representing an audio processing graph.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraph
type AUGraph = uintptr

// AUHostMusicalContextBlock is a block through which hosts provide musical tempo, time signature, and beat position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUHostMusicalContextBlock
type AUHostMusicalContextBlock = func([]float64, []float64, *int, []float64, *int, []float64) bool

// AUHostTransportStateBlock is a block through which hosts provide information about their transport state.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUHostTransportStateBlock
type AUHostTransportStateBlock = func(*AUHostTransportStateFlags, []float64, []float64, []float64) bool

// AUImplementorDisplayNameWithLengthCallback is a block called to obtain a parameter node’s display name, possibly truncated to a desired length.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUImplementorDisplayNameWithLengthCallback
type AUImplementorDisplayNameWithLengthCallback = func(AUParameterNode, int) string

// AUImplementorStringFromValueCallback is a block called to convert a parameter value to a string representation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUImplementorStringFromValueCallback
type AUImplementorStringFromValueCallback = func(AUParameter, *float32) string

// AUImplementorValueFromStringCallback is a block called to convert a string to a parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUImplementorValueFromStringCallback
type AUImplementorValueFromStringCallback = func(AUParameter, string) float32

// AUImplementorValueObserver is a block called to notify the audio unit implementation of changes to a parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUImplementorValueObserver
type AUImplementorValueObserver = func(AUParameter, float32)

// AUImplementorValueProvider is a block called to fetch a parameter’s current value from the audio unit implementation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUImplementorValueProvider
type AUImplementorValueProvider = func(AUParameter) float32

// AUInputHandler is a block to notify the host of an I/O unit that an input is available.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUInputHandler
type AUInputHandler = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int)

// AUInputSamplesInOutputCallback is called by the system when an audio unit has provided a buffer of output samples.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallback
type AUInputSamplesInOutputCallback = func(inRefCon unsafe.Pointer, inOutputTimeStamp uintptr, inInputSample float64, inNumberInputSamples float64)

// AUInternalRenderBlock is a block to render the audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUInternalRenderBlock
type AUInternalRenderBlock = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int, *coreaudiotypes.AudioBufferList, unsafe.Pointer, func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, unsafe.Pointer, int, *coreaudiotypes.AudioBufferList) int32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AUMIDICIProfileChangedBlock
type AUMIDICIProfileChangedBlock = func(uint32, uint32, *coremidi.MIDICIProfile, bool)

// See: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEventListBlock
type AUMIDIEventListBlock = func(int, uint32, *coremidi.MIDIEventList) int32

// AUMIDIOutputCallback is when called by a host application, gets MIDI data from an audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallback
type AUMIDIOutputCallback = func(userData unsafe.Pointer, timeStamp uintptr, midiOutNum uint32, pktlist uintptr) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputEventBlock
type AUMIDIOutputEventBlock = func(int, uint32, int, *uint8) int32

// AUNode is a member of an audio processing graph, associated with an audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUNode
type AUNode = int32

// See: https://developer.apple.com/documentation/AudioToolbox/AUNodeConnection
type AUNodeConnection = AudioUnitNodeConnection

// AUParameterAddress is a numeric identifier for an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterAddress
type AUParameterAddress = uint64

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationObserver
type AUParameterAutomationObserver = func(int, *AUParameterAutomationEvent)

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerBlock
type AUParameterListenerBlock = func(unsafe.Pointer, *AudioUnitParameter, float32)

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerProc
type AUParameterListenerProc = func(unsafe.Pointer, unsafe.Pointer, uintptr, float32)

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerRef
type AUParameterListenerRef uintptr

// AUParameterObserver is a block called after the value of a parameter changes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterObserver
type AUParameterObserver = func(uint64, float32)

// AUParameterObserverToken is a token representing an installed parameter observer block.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterObserverToken
type AUParameterObserverToken = unsafe.Pointer

// AUParameterRecordingObserver is a block called to record parameter changes as automation events.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterRecordingObserver
type AUParameterRecordingObserver = func(int, *AURecordedParameterEvent)

// AURenderBlock is a block to render the audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AURenderBlock
type AURenderBlock = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int, *coreaudiotypes.AudioBufferList, func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, unsafe.Pointer, int, *coreaudiotypes.AudioBufferList) int32) int32

// AURenderCallback is called by the system when an audio unit requires input samples, or before and after a render operation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AURenderCallback
type AURenderCallback = func(inRefCon unsafe.Pointer, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp uintptr, inBusNumber uint32, inNumberFrames uint32, ioData uintptr) int32

// AURenderContextObserver is a custom block that tells the audio unit which thread context to use for the next render cycle.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AURenderContextObserver
type AURenderContextObserver = func(*AudioUnitRenderContext)

// AURenderObserver is a block called when an audio unit renders audio.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AURenderObserver
type AURenderObserver = func(uint, *coreaudiotypes.AudioTimeStamp, uint32, int)

// AURenderPullInputBlock is a block to supply audio input to a render block.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AURenderPullInputBlock
type AURenderPullInputBlock = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int, *coreaudiotypes.AudioBufferList) int32

// AUScheduleMIDIEventBlock is a block to schedule MIDI events.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUScheduleMIDIEventBlock
type AUScheduleMIDIEventBlock = func(int, uint32, int, *uint8)

// AUScheduleParameterBlock is a block to schedule parameter changes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUScheduleParameterBlock
type AUScheduleParameterBlock = func(int, uint32, uint64, float32)

// AUValue is a value of an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUValue
type AUValue = float32

// AUVoiceIOMutedSpeechActivityEventListener is a block that the system calls to indicate speech activity while the user has the microphone muted.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOMutedSpeechActivityEventListener
type AUVoiceIOMutedSpeechActivityEventListener = func(event AUVoiceIOSpeechActivityEvent)

// AudioCodec is an instance of a Component Manager component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodec
type AudioCodec = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferListProc
type AudioCodecAppendInputBufferListProc = func(unsafe.Pointer, uintptr, *uint32, uintptr, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputDataProc
type AudioCodecAppendInputDataProc = func(unsafe.Pointer, unsafe.Pointer, *uint32, *uint32, uintptr) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfoProc
type AudioCodecGetPropertyInfoProc = func(unsafe.Pointer, uint32, *uint32, *byte) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyProc
type AudioCodecGetPropertyProc = func(unsafe.Pointer, uint32, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitializeProc
type AudioCodecInitializeProc = func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferListProc
type AudioCodecProduceOutputBufferListProc = func(unsafe.Pointer, uintptr, *uint32, uintptr, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPacketsProc
type AudioCodecProduceOutputPacketsProc = func(unsafe.Pointer, unsafe.Pointer, *uint32, *uint32, uintptr, *uint32) int32

// AudioCodecPropertyID is an integer identifying an audio codec property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPropertyID
type AudioCodecPropertyID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecResetProc
type AudioCodecResetProc = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetPropertyProc
type AudioCodecSetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitializeProc
type AudioCodecUninitializeProc = func(unsafe.Pointer) int32

// AudioComponent is an audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponent
type AudioComponent = uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFactoryFunction
type AudioComponentFactoryFunction = func(uintptr) uintptr

// AudioComponentInstance is a component instance, or object, is an audio unit or audio codec.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstance
type AudioComponentInstance = uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentMethod
type AudioComponentMethod = func(unsafe.Pointer, unsafe.Pointer) int32

// AudioConverterComplexInputDataProc is supplies input data to the [AudioConverterFillComplexBuffer(_:_:_:_:_:_:)] function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProc
type AudioConverterComplexInputDataProc = func(inAudioConverter AudioConverterRef, ioNumberDataPackets *uint32, ioData uintptr, outDataPacketDescription uintptr, inUserData unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProcRealtimeSafe
type AudioConverterComplexInputDataProcRealtimeSafe = func(AudioConverterRef, *uint32, uintptr, uintptr, unsafe.Pointer) int32

// AudioConverterInputDataProc is deprecated. Use [AudioConverterFillComplexBuffer(_:_:_:_:_:_:)] instead.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterInputDataProc
type AudioConverterInputDataProc = func(AudioConverterRef, *uint32, unsafe.Pointer, unsafe.Pointer) int32

// AudioConverterPropertyID is an audio converter property identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPropertyID
type AudioConverterPropertyID = uint32

// AudioConverterRef is a reference to an audio converter object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterRef
type AudioConverterRef uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponent
type AudioFileComponent = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCloseProc
type AudioFileComponentCloseProc = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserDataProc
type AudioFileComponentCountUserDataProc = func(unsafe.Pointer, uint32, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURLProc
type AudioFileComponentCreateURLProc = func(unsafe.Pointer, uintptr, uintptr, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormatProc
type AudioFileComponentExtensionIsThisFormatProc = func(unsafe.Pointer, uintptr, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormatProc
type AudioFileComponentFileDataIsThisFormatProc = func(unsafe.Pointer, uint32, unsafe.Pointer, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoProc
type AudioFileComponentGetGlobalInfoProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSizeProc
type AudioFileComponentGetGlobalInfoSizeProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfoProc
type AudioFileComponentGetPropertyInfoProc = func(unsafe.Pointer, uint32, *uint32, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyProc
type AudioFileComponentGetPropertyProc = func(unsafe.Pointer, uint32, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffsetProc
type AudioFileComponentGetUserDataAtOffsetProc = func(unsafe.Pointer, uint32, uint32, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataProc
type AudioFileComponentGetUserDataProc = func(unsafe.Pointer, uint32, uint32, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64Proc
type AudioFileComponentGetUserDataSize64Proc = func(unsafe.Pointer, uint32, uint32, *uint64) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSizeProc
type AudioFileComponentGetUserDataSizeProc = func(unsafe.Pointer, uint32, uint32, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacksProc
type AudioFileComponentInitializeWithCallbacksProc = func(unsafe.Pointer, unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc, uint32, uintptr, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURLProc
type AudioFileComponentOpenURLProc = func(unsafe.Pointer, uintptr, int8, int32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacksProc
type AudioFileComponentOpenWithCallbacksProc = func(unsafe.Pointer, unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimizeProc
type AudioFileComponentOptimizeProc = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentPropertyID
type AudioFileComponentPropertyID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytesProc
type AudioFileComponentReadBytesProc = func(unsafe.Pointer, uint8, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketDataProc
type AudioFileComponentReadPacketDataProc = func(unsafe.Pointer, uint8, *uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketsProc
type AudioFileComponentReadPacketsProc = func(unsafe.Pointer, uint8, *uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserDataProc
type AudioFileComponentRemoveUserDataProc = func(unsafe.Pointer, uint32, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetPropertyProc
type AudioFileComponentSetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserDataProc
type AudioFileComponentSetUserDataProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytesProc
type AudioFileComponentWriteBytesProc = func(unsafe.Pointer, uint8, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePacketsProc
type AudioFileComponentWritePacketsProc = func(unsafe.Pointer, uint8, uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// AudioFileID is an opaque data type that represents an audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileID
type AudioFileID = uintptr

// AudioFilePropertyID is an audio file property identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFilePropertyID
type AudioFilePropertyID = uint32

// AudioFileStreamID is defines an opaque data type that represents an audio file stream parser.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamID
type AudioFileStreamID = uintptr

// AudioFileStreamPropertyID is uniquely identifies an audio file stream property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyID
type AudioFileStreamPropertyID = uint32

// AudioFileStream_PacketsProc is invoked by an audio file stream parser when it finds audio data in the audio file stream.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStream_PacketsProc
type AudioFileStream_PacketsProc = func(inClientData unsafe.Pointer, inNumberBytes uint32, inNumberPackets uint32, inInputData unsafe.Pointer, inPacketDescriptions uintptr)

// AudioFileStream_PropertyListenerProc is invoked by an audio file stream parser when it finds a property value in the audio file stream.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStream_PropertyListenerProc
type AudioFileStream_PropertyListenerProc = func(inClientData unsafe.Pointer, inAudioFileStream AudioFileStreamID, inPropertyID uint32, ioFlags *AudioFileStreamPropertyFlags)

// AudioFileTypeID is operating system constants that indicate the type of file to be written or a hint about what type of file to expect from data provided.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeID
type AudioFileTypeID = uint32

// AudioFile_GetSizeProc is gets file data size.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFile_GetSizeProc
type AudioFile_GetSizeProc = func(inClientData unsafe.Pointer) int64

// AudioFile_ReadProc is reads audio data when used in conjunction with the [AudioFileOpenWithCallbacks(_:_:_:_:_:_:_:)] or [AudioFileInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)] functions.).
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFile_ReadProc
type AudioFile_ReadProc = func(inClientData unsafe.Pointer, inPosition int64, requestCount uint32, buffer unsafe.Pointer, actualCount *uint32) int32

// AudioFile_SetSizeProc is sets file data size.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SetSizeProc
type AudioFile_SetSizeProc = func(unsafe.Pointer, int64) int32

// AudioFile_WriteProc is a callback for writing file data when used in conjunction with the [AudioFileOpenWithCallbacks(_:_:_:_:_:_:_:)] or [AudioFileCreateWithURL(_:_:_:_:_:)] functions.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFile_WriteProc
type AudioFile_WriteProc = func(inClientData unsafe.Pointer, inPosition int64, requestCount uint32, buffer unsafe.Pointer, actualCount *uint32) int32

// AudioFormatPropertyID is a type for four-char codes for audio format property identifiers.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFormatPropertyID
type AudioFormatPropertyID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartProc
type AudioOutputUnitStartProc = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStopProc
type AudioOutputUnitStopProc = func(unsafe.Pointer) int32

// AudioQueueBufferRef is a pointer to an audio queue buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBufferRef
type AudioQueueBufferRef = *AudioQueueBuffer

// AudioQueueInputCallback is called by the system when a recording audio queue has finished filling an audio queue buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueInputCallback
type AudioQueueInputCallback = func(inUserData unsafe.Pointer, inAQ AudioQueueRef, inBuffer uintptr, inStartTime uintptr, inNumberPacketDescriptions uint32, inPacketDescs uintptr)

// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueInputCallbackBlock
type AudioQueueInputCallbackBlock = func(*uintptr, *AudioQueueBuffer, *coreaudiotypes.AudioTimeStamp, uint32, *coreaudiotypes.AudioStreamPacketDescription)

// AudioQueueOutputCallback is called by the system when an audio queue buffer is available for reuse.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOutputCallback
type AudioQueueOutputCallback = func(inUserData unsafe.Pointer, inAQ AudioQueueRef, inBuffer uintptr)

// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOutputCallbackBlock
type AudioQueueOutputCallbackBlock = func(*uintptr, *AudioQueueBuffer)

// AudioQueueParameterID is a [UInt32] value that uniquely identifies an audio queue parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterID
type AudioQueueParameterID = uint32

// AudioQueueParameterValue is a [Float32] value for an audio queue parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterValue
type AudioQueueParameterValue = float32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapCallback
type AudioQueueProcessingTapCallback = func(unsafe.Pointer, AudioQueueProcessingTapRef, uint32, uintptr, *AudioQueueProcessingTapFlags, *uint32, uintptr)

// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapRef
type AudioQueueProcessingTapRef uintptr

// AudioQueuePropertyID is identifiers for audio queue properties.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePropertyID
type AudioQueuePropertyID = uint32

// AudioQueuePropertyListenerProc is called by the system when a specified audio queue property changes value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePropertyListenerProc
type AudioQueuePropertyListenerProc = func(inUserData unsafe.Pointer, inAQ AudioQueueRef, inID uint32)

// AudioQueueRef is defines an opaque data type that represents an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueRef
type AudioQueueRef = kernel.Pointer

// AudioQueueTimelineRef is defines an opaque data type that represents an audio queue timeline object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueTimelineRef
type AudioQueueTimelineRef uintptr

// AudioServicesPropertyID is the data type for a system sound property identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPropertyID
type AudioServicesPropertyID = uint32

// AudioServicesSystemSoundCompletionProc is a function the system invokes when a system sound finishes playing.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesSystemSoundCompletionProc
type AudioServicesSystemSoundCompletionProc = func(ssID uint32, clientData unsafe.Pointer)

// AudioSessionInterruptionType is values that indicate the nature of the interruption that ended.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInterruptionType
type AudioSessionInterruptionType = uint32

// AudioUnit is the data type for a plug-in component that provides audio processing or audio data generation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnit
type AudioUnit = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListenerProc
type AudioUnitAddPropertyListenerProc = func(unsafe.Pointer, uint32, AudioUnitPropertyListenerProc, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotifyProc
type AudioUnitAddRenderNotifyProc = func(unsafe.Pointer, func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uint32, uintptr) int32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitCarbonView
type AudioUnitCarbonView = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitCarbonViewEventID
type AudioUnitCarbonViewEventID = int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitComplexRenderProc
type AudioUnitComplexRenderProc = func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uint32, *uint32, uintptr, uintptr, unsafe.Pointer, *uint32) int32

// AudioUnitElement is the data type for an audio unit element identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitElement
type AudioUnitElement = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameterProc
type AudioUnitGetParameterProc = func(unsafe.Pointer, uint32, uint32, uint32, []float32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfoProc
type AudioUnitGetPropertyInfoProc = func(unsafe.Pointer, uint32, uint32, uint32, *uint32, *byte) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyProc
type AudioUnitGetPropertyProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitializeProc
type AudioUnitInitializeProc = func(unsafe.Pointer) int32

// AudioUnitParameterID is the data type for an audio unit parameter identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterID
type AudioUnitParameterID = uint32

// AudioUnitParameterIDName is a type definition for a data type that defines the short version of the name for an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterIDName
type AudioUnitParameterIDName = AudioUnitParameterNameInfo

// AudioUnitParameterValue is the data type for an audio unit parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValue
type AudioUnitParameterValue = float32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessMultipleProc
type AudioUnitProcessMultipleProc = func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uint32, uintptr, uint32, uintptr) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessProc
type AudioUnitProcessProc = func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uintptr) int32

// AudioUnitPropertyID is the data type for audio unit property keys.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyID
type AudioUnitPropertyID = uint32

// AudioUnitPropertyListenerProc is called by the system when the value of a specified audio unit property has changed.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyListenerProc
type AudioUnitPropertyListenerProc = func(inRefCon unsafe.Pointer, inUnit AudioComponentInstance, inID uint32, inScope uint32, inElement uint32)

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEventListener
type AudioUnitRemoteControlEventListener = func(AudioUnitRemoteControlEvent)

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerProc
type AudioUnitRemovePropertyListenerProc = func(unsafe.Pointer, uint32, AudioUnitPropertyListenerProc) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserDataProc
type AudioUnitRemovePropertyListenerWithUserDataProc = func(unsafe.Pointer, uint32, AudioUnitPropertyListenerProc, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotifyProc
type AudioUnitRemoveRenderNotifyProc = func(unsafe.Pointer, func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uint32, uintptr) int32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderProc
type AudioUnitRenderProc = func(unsafe.Pointer, *AudioUnitRenderActionFlags, uintptr, uint32, uint32, uintptr) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitResetProc
type AudioUnitResetProc = func(unsafe.Pointer, uint32, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParametersProc
type AudioUnitScheduleParametersProc = func(unsafe.Pointer, uintptr, uint32) int32

// AudioUnitScope is the data type for audio unit scope identifiers.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScope
type AudioUnitScope = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameterProc
type AudioUnitSetParameterProc = func(unsafe.Pointer, uint32, uint32, uint32, float32, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetPropertyProc
type AudioUnitSetPropertyProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitializeProc
type AudioUnitUninitializeProc = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockBeats
type CAClockBeats = float64

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockListenerProc
type CAClockListenerProc = func(unsafe.Pointer, CAClockMessage, unsafe.Pointer)

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockRef
type CAClockRef uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTEFormat
type CAClockSMPTEFormat = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSamples
type CAClockSamples = float64

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSeconds
type CAClockSeconds = float64

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockTempo
type CAClockTempo = float64

// See: https://developer.apple.com/documentation/AudioToolbox/CallHostBlock
type CallHostBlock = func(foundation.INSDictionary) foundation.INSDictionary

// See: https://developer.apple.com/documentation/AudioToolbox/CountUserDataFDF
type CountUserDataFDF = func(unsafe.Pointer, uint32, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFilePacketTableInfoOverride
type ExtAudioFilePacketTableInfoOverride = int32

// ExtAudioFilePropertyID is an audio file object property identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFilePropertyID
type ExtAudioFilePropertyID = uint32

// ExtAudioFileRef is an opaque structure representing an extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileRef
type ExtAudioFileRef uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/GetPropertyFDF
type GetPropertyFDF = func(unsafe.Pointer, uint32, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/GetPropertyInfoFDF
type GetPropertyInfoFDF = func(unsafe.Pointer, uint32, *uint32, *uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/GetUserDataFDF
type GetUserDataFDF = func(unsafe.Pointer, uint32, uint32, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/GetUserDataSizeFDF
type GetUserDataSizeFDF = func(unsafe.Pointer, uint32, uint32, *uint32) int32

// HostCallback_GetBeatAndTempo is when called by the system, provides beat and tempo information to an audio unit from a host application.
//
// See: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetBeatAndTempo
type HostCallback_GetBeatAndTempo = func(inHostUserData unsafe.Pointer, outCurrentBeat []float64, outCurrentTempo []float64) int32

// HostCallback_GetMusicalTimeLocation is when called by the system, provides musical timing information to an audio unit from a host application.
//
// See: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetMusicalTimeLocation
type HostCallback_GetMusicalTimeLocation = func(inHostUserData unsafe.Pointer, outDeltaSampleOffsetToNextBeat *uint32, outTimeSig_Numerator []float32, outTimeSig_Denominator *uint32, outCurrentMeasureDownBeat []float64) int32

// HostCallback_GetTransportState is when called by the system, provides audio transport state and timeline information to an audio unit from a host application.
//
// See: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState
type HostCallback_GetTransportState = func(inHostUserData unsafe.Pointer, outIsPlaying *byte, outTransportStateChanged *byte, outCurrentSampleInTimeLine []float64, outIsCycling *byte, outCycleStartBeat []float64, outCycleEndBeat []float64) int32

// See: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState2
type HostCallback_GetTransportState2 = func(unsafe.Pointer, *byte, *byte, *byte, []float64, *byte, []float64, []float64) int32

// MIDIChannelNumber is mIDI Channel, 0~15 (channels 1 through 16, respectively).
//
// See: https://developer.apple.com/documentation/AudioToolbox/MIDIChannelNumber
type MIDIChannelNumber = uint8

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceComponent
type MusicDeviceComponent = unsafe.Pointer

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceGroupID
type MusicDeviceGroupID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceInstrumentID
type MusicDeviceInstrumentID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEventProc
type MusicDeviceMIDIEventProc = func(unsafe.Pointer, uint32, uint32, uint32, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNoteProc
type MusicDeviceStartNoteProc = func(unsafe.Pointer, uint32, uint32, *uint32, uint32, uintptr) int32

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNoteProc
type MusicDeviceStopNoteProc = func(unsafe.Pointer, uint32, uint32, uint32) int32

// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysExProc
type MusicDeviceSysExProc = func(unsafe.Pointer, *byte, uint32) int32

// MusicEventIterator is a music event iterator sequentially handles events on a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIterator
type MusicEventIterator = kernel.Pointer

// MusicEventType is mIDI and other music event types, used by music event iterator functions.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventType
type MusicEventType = uint32

// MusicPlayer is a music player plays a music sequence (of type [MusicSequence]).
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayer
type MusicPlayer = uintptr

// MusicSequence is a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequence
type MusicSequence = uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceUserCallback
type MusicSequenceUserCallback = func(unsafe.Pointer, MusicSequence, MusicTrack, float64, uintptr, float64, float64)

// MusicTimeStamp is a timestamp for use by a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTimeStamp
type MusicTimeStamp = float64

// MusicTrack is a music track consists of a series of music events, each timestamped using units of beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrack
type MusicTrack = uintptr

// See: https://developer.apple.com/documentation/AudioToolbox/NoteInstanceID
type NoteInstanceID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/ReadBytesFDF
type ReadBytesFDF = func(unsafe.Pointer, uint8, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/ReadPacketDataFDF
type ReadPacketDataFDF = func(unsafe.Pointer, uint8, *uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/ReadPacketsFDF
type ReadPacketsFDF = func(unsafe.Pointer, uint8, *uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegionCompletionProc
type ScheduledAudioFileRegionCompletionProc = func(unsafe.Pointer, uintptr, int32)

// See: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSliceCompletionProc
type ScheduledAudioSliceCompletionProc = func(unsafe.Pointer, uintptr)

// See: https://developer.apple.com/documentation/AudioToolbox/SetPropertyFDF
type SetPropertyFDF = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/SetUserDataFDF
type SetUserDataFDF = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32

// SystemSoundID is a system sound object, identified with a sound file you want to play.
//
// See: https://developer.apple.com/documentation/AudioToolbox/SystemSoundID
type SystemSoundID = uint32

// See: https://developer.apple.com/documentation/AudioToolbox/WriteBytesFDF
type WriteBytesFDF = func(unsafe.Pointer, uint8, int64, *uint32, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/AudioToolbox/WritePacketsFDF
type WritePacketsFDF = func(unsafe.Pointer, uint8, uint32, uintptr, int64, *uint32, unsafe.Pointer) int32

// AudioFileStreamPacketsProc is a Go-name alias for AudioFileStream_PacketsProc.
type AudioFileStreamPacketsProc = AudioFileStream_PacketsProc

// AudioFileStreamPropertyListenerProc is a Go-name alias for AudioFileStream_PropertyListenerProc.
type AudioFileStreamPropertyListenerProc = AudioFileStream_PropertyListenerProc

// AudioFileGetSizeProc is a Go-name alias for AudioFile_GetSizeProc.
type AudioFileGetSizeProc = AudioFile_GetSizeProc

// AudioFileReadProc is a Go-name alias for AudioFile_ReadProc.
type AudioFileReadProc = AudioFile_ReadProc

// AudioFileSetSizeProc is a Go-name alias for AudioFile_SetSizeProc.
type AudioFileSetSizeProc = AudioFile_SetSizeProc

// AudioFileWriteProc is a Go-name alias for AudioFile_WriteProc.
type AudioFileWriteProc = AudioFile_WriteProc

// HostCallbackGetBeatAndTempo is a Go-name alias for HostCallback_GetBeatAndTempo.
type HostCallbackGetBeatAndTempo = HostCallback_GetBeatAndTempo

// HostCallbackGetMusicalTimeLocation is a Go-name alias for HostCallback_GetMusicalTimeLocation.
type HostCallbackGetMusicalTimeLocation = HostCallback_GetMusicalTimeLocation

// HostCallbackGetTransportState is a Go-name alias for HostCallback_GetTransportState.
type HostCallbackGetTransportState = HostCallback_GetTransportState

// HostCallbackGetTransportState2 is a Go-name alias for HostCallback_GetTransportState2.
type HostCallbackGetTransportState2 = HostCallback_GetTransportState2
