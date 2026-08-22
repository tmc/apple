// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve
type AU3DMixerAttenuationCurve uint32

const (
	// K3DMixerAttenuationCurve_Exponential: An exponential attenuation curve.
	K3DMixerAttenuationCurve_Exponential AU3DMixerAttenuationCurve = 1
	// K3DMixerAttenuationCurve_Inverse: An inverse attenuation curve.
	K3DMixerAttenuationCurve_Inverse AU3DMixerAttenuationCurve = 2
	// K3DMixerAttenuationCurve_Linear: A linear attenuation curve.
	K3DMixerAttenuationCurve_Linear AU3DMixerAttenuationCurve = 3
	// K3DMixerAttenuationCurve_Power: An equal-power-based attenuation curve.
	K3DMixerAttenuationCurve_Power AU3DMixerAttenuationCurve = 0
)

func (e AU3DMixerAttenuationCurve) String() string {
	switch e {
	case K3DMixerAttenuationCurve_Exponential:
		return "K3DMixerAttenuationCurve_Exponential"
	case K3DMixerAttenuationCurve_Inverse:
		return "K3DMixerAttenuationCurve_Inverse"
	case K3DMixerAttenuationCurve_Linear:
		return "K3DMixerAttenuationCurve_Linear"
	case K3DMixerAttenuationCurve_Power:
		return "K3DMixerAttenuationCurve_Power"
	default:
		return fmt.Sprintf("AU3DMixerAttenuationCurve(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerRenderingFlags
type AU3DMixerRenderingFlags uint32

const (
	K3DMixerRenderingFlags_ConstantReverbBlend       AU3DMixerRenderingFlags = 64
	K3DMixerRenderingFlags_DistanceAttenuation       AU3DMixerRenderingFlags = 4
	K3DMixerRenderingFlags_DistanceDiffusion         AU3DMixerRenderingFlags = 16
	K3DMixerRenderingFlags_DistanceFilter            AU3DMixerRenderingFlags = 8
	K3DMixerRenderingFlags_DopplerShift              AU3DMixerRenderingFlags = 2
	K3DMixerRenderingFlags_InterAuralDelay           AU3DMixerRenderingFlags = 1
	K3DMixerRenderingFlags_LinearDistanceAttenuation AU3DMixerRenderingFlags = 32
)

func (e AU3DMixerRenderingFlags) String() string {
	switch e {
	case K3DMixerRenderingFlags_ConstantReverbBlend:
		return "K3DMixerRenderingFlags_ConstantReverbBlend"
	case K3DMixerRenderingFlags_DistanceAttenuation:
		return "K3DMixerRenderingFlags_DistanceAttenuation"
	case K3DMixerRenderingFlags_DistanceDiffusion:
		return "K3DMixerRenderingFlags_DistanceDiffusion"
	case K3DMixerRenderingFlags_DistanceFilter:
		return "K3DMixerRenderingFlags_DistanceFilter"
	case K3DMixerRenderingFlags_DopplerShift:
		return "K3DMixerRenderingFlags_DopplerShift"
	case K3DMixerRenderingFlags_InterAuralDelay:
		return "K3DMixerRenderingFlags_InterAuralDelay"
	case K3DMixerRenderingFlags_LinearDistanceAttenuation:
		return "K3DMixerRenderingFlags_LinearDistanceAttenuation"
	default:
		return fmt.Sprintf("AU3DMixerRenderingFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle
type AUAudioMixRenderingStyle uint32

const (
	KAudioMixRenderingStyle_Cinematic               AUAudioMixRenderingStyle = 0
	KAudioMixRenderingStyle_CinematicBackgroundStem AUAudioMixRenderingStyle = 3
	KAudioMixRenderingStyle_CinematicForegroundStem AUAudioMixRenderingStyle = 4
	KAudioMixRenderingStyle_InFrame                 AUAudioMixRenderingStyle = 2
	KAudioMixRenderingStyle_InFrameBackgroundStem   AUAudioMixRenderingStyle = 9
	KAudioMixRenderingStyle_InFrameForegroundStem   AUAudioMixRenderingStyle = 6
	KAudioMixRenderingStyle_Standard                AUAudioMixRenderingStyle = 7
	KAudioMixRenderingStyle_Studio                  AUAudioMixRenderingStyle = 1
	KAudioMixRenderingStyle_StudioBackgroundStem    AUAudioMixRenderingStyle = 8
	KAudioMixRenderingStyle_StudioForegroundStem    AUAudioMixRenderingStyle = 5
)

func (e AUAudioMixRenderingStyle) String() string {
	switch e {
	case KAudioMixRenderingStyle_Cinematic:
		return "KAudioMixRenderingStyle_Cinematic"
	case KAudioMixRenderingStyle_CinematicBackgroundStem:
		return "KAudioMixRenderingStyle_CinematicBackgroundStem"
	case KAudioMixRenderingStyle_CinematicForegroundStem:
		return "KAudioMixRenderingStyle_CinematicForegroundStem"
	case KAudioMixRenderingStyle_InFrame:
		return "KAudioMixRenderingStyle_InFrame"
	case KAudioMixRenderingStyle_InFrameBackgroundStem:
		return "KAudioMixRenderingStyle_InFrameBackgroundStem"
	case KAudioMixRenderingStyle_InFrameForegroundStem:
		return "KAudioMixRenderingStyle_InFrameForegroundStem"
	case KAudioMixRenderingStyle_Standard:
		return "KAudioMixRenderingStyle_Standard"
	case KAudioMixRenderingStyle_Studio:
		return "KAudioMixRenderingStyle_Studio"
	case KAudioMixRenderingStyle_StudioBackgroundStem:
		return "KAudioMixRenderingStyle_StudioBackgroundStem"
	case KAudioMixRenderingStyle_StudioForegroundStem:
		return "KAudioMixRenderingStyle_StudioForegroundStem"
	default:
		return fmt.Sprintf("AUAudioMixRenderingStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusType
type AUAudioUnitBusType int

const (
	// AUAudioUnitBusTypeInput: An input bus.
	AUAudioUnitBusTypeInput AUAudioUnitBusType = 1
	// AUAudioUnitBusTypeOutput: An output bus.
	AUAudioUnitBusTypeOutput AUAudioUnitBusType = 2
)

func (e AUAudioUnitBusType) String() string {
	switch e {
	case AUAudioUnitBusTypeInput:
		return "AUAudioUnitBusTypeInput"
	case AUAudioUnitBusTypeOutput:
		return "AUAudioUnitBusTypeOutput"
	default:
		return fmt.Sprintf("AUAudioUnitBusType(%d)", e)
	}
}

type AUEventSampleTimeImmediateConstants int64

const (
	AUEventSampleTimeImmediate AUEventSampleTimeImmediateConstants = -4294967296
)

func (e AUEventSampleTimeImmediateConstants) String() string {
	switch e {
	case AUEventSampleTimeImmediate:
		return "AUEventSampleTimeImmediate"
	default:
		return fmt.Sprintf("AUEventSampleTimeImmediateConstants(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUHostTransportStateFlags
type AUHostTransportStateFlags uint

const (
	// AUHostTransportStateChanged: Indicates such state changes as start, stop, or seeking to another position in the timeline.
	AUHostTransportStateChanged AUHostTransportStateFlags = 1
	// AUHostTransportStateCycling: Indicates that the host is cycling or looping.
	AUHostTransportStateCycling AUHostTransportStateFlags = 8
	// AUHostTransportStateMoving: Indicates that the audio transport is moving.
	AUHostTransportStateMoving AUHostTransportStateFlags = 2
	// AUHostTransportStateRecording: Indicates that the host is recording, or is prepared to record.
	AUHostTransportStateRecording AUHostTransportStateFlags = 4
)

func (e AUHostTransportStateFlags) String() string {
	switch e {
	case AUHostTransportStateChanged:
		return "AUHostTransportStateChanged"
	case AUHostTransportStateCycling:
		return "AUHostTransportStateCycling"
	case AUHostTransportStateMoving:
		return "AUHostTransportStateMoving"
	case AUHostTransportStateRecording:
		return "AUHostTransportStateRecording"
	default:
		return fmt.Sprintf("AUHostTransportStateFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEventType
type AUParameterAutomationEventType uint32

const (
	AUParameterAutomationEventTypeRelease AUParameterAutomationEventType = 2
	AUParameterAutomationEventTypeTouch   AUParameterAutomationEventType = 1
	AUParameterAutomationEventTypeValue   AUParameterAutomationEventType = 0
)

func (e AUParameterAutomationEventType) String() string {
	switch e {
	case AUParameterAutomationEventTypeRelease:
		return "AUParameterAutomationEventTypeRelease"
	case AUParameterAutomationEventTypeTouch:
		return "AUParameterAutomationEventTypeTouch"
	case AUParameterAutomationEventTypeValue:
		return "AUParameterAutomationEventTypeValue"
	default:
		return fmt.Sprintf("AUParameterAutomationEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterEventType
type AUParameterEventType uint32

const (
	// KParameterEvent_Immediate: An immediate change from the parameter’s previous value to a new value.
	KParameterEvent_Immediate AUParameterEventType = 1
	// KParameterEvent_Ramped: A gradual change from the parameter’s previous value to a new value, applied linearly over a specified period of time
	KParameterEvent_Ramped AUParameterEventType = 2
)

func (e AUParameterEventType) String() string {
	switch e {
	case KParameterEvent_Immediate:
		return "KParameterEvent_Immediate"
	case KParameterEvent_Ramped:
		return "KParameterEvent_Ramped"
	default:
		return fmt.Sprintf("AUParameterEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMappingFlags
type AUParameterMIDIMappingFlags uint32

const (
	KAUParameterMIDIMapping_AnyChannelFlag AUParameterMIDIMappingFlags = 1
	KAUParameterMIDIMapping_AnyNoteFlag    AUParameterMIDIMappingFlags = 2
	KAUParameterMIDIMapping_Bipolar        AUParameterMIDIMappingFlags = 16
	KAUParameterMIDIMapping_Bipolar_On     AUParameterMIDIMappingFlags = 32
	KAUParameterMIDIMapping_SubRange       AUParameterMIDIMappingFlags = 4
	KAUParameterMIDIMapping_Toggle         AUParameterMIDIMappingFlags = 8
)

func (e AUParameterMIDIMappingFlags) String() string {
	switch e {
	case KAUParameterMIDIMapping_AnyChannelFlag:
		return "KAUParameterMIDIMapping_AnyChannelFlag"
	case KAUParameterMIDIMapping_AnyNoteFlag:
		return "KAUParameterMIDIMapping_AnyNoteFlag"
	case KAUParameterMIDIMapping_Bipolar:
		return "KAUParameterMIDIMapping_Bipolar"
	case KAUParameterMIDIMapping_Bipolar_On:
		return "KAUParameterMIDIMapping_Bipolar_On"
	case KAUParameterMIDIMapping_SubRange:
		return "KAUParameterMIDIMapping_SubRange"
	case KAUParameterMIDIMapping_Toggle:
		return "KAUParameterMIDIMapping_Toggle"
	default:
		return fmt.Sprintf("AUParameterMIDIMappingFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AURenderEventType
type AURenderEventType uint8

const (
	// AURenderEventMIDI: A MIDI event.
	AURenderEventMIDI          AURenderEventType = 8
	AURenderEventMIDIEventList AURenderEventType = 10
	// AURenderEventMIDISysEx: A system-exclusive MIDI event.
	AURenderEventMIDISysEx AURenderEventType = 9
	// AURenderEventParameter: A parameter event.
	AURenderEventParameter AURenderEventType = 1
	// AURenderEventParameterRamp: A ramped parameter event.
	AURenderEventParameterRamp AURenderEventType = 2
)

func (e AURenderEventType) String() string {
	switch e {
	case AURenderEventMIDI:
		return "AURenderEventMIDI"
	case AURenderEventMIDIEventList:
		return "AURenderEventMIDIEventList"
	case AURenderEventMIDISysEx:
		return "AURenderEventMIDISysEx"
	case AURenderEventParameter:
		return "AURenderEventParameter"
	case AURenderEventParameterRamp:
		return "AURenderEventParameterRamp"
	default:
		return fmt.Sprintf("AURenderEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType
type AUReverbRoomType uint32

const (
	KReverbRoomType_Cathedral     AUReverbRoomType = 8
	KReverbRoomType_LargeChamber  AUReverbRoomType = 7
	KReverbRoomType_LargeHall     AUReverbRoomType = 4
	KReverbRoomType_LargeHall2    AUReverbRoomType = 12
	KReverbRoomType_LargeRoom     AUReverbRoomType = 2
	KReverbRoomType_LargeRoom2    AUReverbRoomType = 9
	KReverbRoomType_MediumChamber AUReverbRoomType = 6
	KReverbRoomType_MediumHall    AUReverbRoomType = 3
	KReverbRoomType_MediumHall2   AUReverbRoomType = 10
	KReverbRoomType_MediumHall3   AUReverbRoomType = 11
	KReverbRoomType_MediumRoom    AUReverbRoomType = 1
	KReverbRoomType_Plate         AUReverbRoomType = 5
	KReverbRoomType_SmallRoom     AUReverbRoomType = 0
)

func (e AUReverbRoomType) String() string {
	switch e {
	case KReverbRoomType_Cathedral:
		return "KReverbRoomType_Cathedral"
	case KReverbRoomType_LargeChamber:
		return "KReverbRoomType_LargeChamber"
	case KReverbRoomType_LargeHall:
		return "KReverbRoomType_LargeHall"
	case KReverbRoomType_LargeHall2:
		return "KReverbRoomType_LargeHall2"
	case KReverbRoomType_LargeRoom:
		return "KReverbRoomType_LargeRoom"
	case KReverbRoomType_LargeRoom2:
		return "KReverbRoomType_LargeRoom2"
	case KReverbRoomType_MediumChamber:
		return "KReverbRoomType_MediumChamber"
	case KReverbRoomType_MediumHall:
		return "KReverbRoomType_MediumHall"
	case KReverbRoomType_MediumHall2:
		return "KReverbRoomType_MediumHall2"
	case KReverbRoomType_MediumHall3:
		return "KReverbRoomType_MediumHall3"
	case KReverbRoomType_MediumRoom:
		return "KReverbRoomType_MediumRoom"
	case KReverbRoomType_Plate:
		return "KReverbRoomType_Plate"
	case KReverbRoomType_SmallRoom:
		return "KReverbRoomType_SmallRoom"
	default:
		return fmt.Sprintf("AUReverbRoomType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUScheduledAudioSliceFlags
type AUScheduledAudioSliceFlags uint32

const (
	KScheduledAudioSliceFlag_BeganToRender     AUScheduledAudioSliceFlags = 0x2
	KScheduledAudioSliceFlag_BeganToRenderLate AUScheduledAudioSliceFlags = 0x4
	KScheduledAudioSliceFlag_Complete          AUScheduledAudioSliceFlags = 0x1
	KScheduledAudioSliceFlag_Interrupt         AUScheduledAudioSliceFlags = 0x10
	KScheduledAudioSliceFlag_InterruptAtLoop   AUScheduledAudioSliceFlags = 0x20
	KScheduledAudioSliceFlag_Loop              AUScheduledAudioSliceFlags = 0x8
)

func (e AUScheduledAudioSliceFlags) String() string {
	switch e {
	case KScheduledAudioSliceFlag_BeganToRender:
		return "KScheduledAudioSliceFlag_BeganToRender"
	case KScheduledAudioSliceFlag_BeganToRenderLate:
		return "KScheduledAudioSliceFlag_BeganToRenderLate"
	case KScheduledAudioSliceFlag_Complete:
		return "KScheduledAudioSliceFlag_Complete"
	case KScheduledAudioSliceFlag_Interrupt:
		return "KScheduledAudioSliceFlag_Interrupt"
	case KScheduledAudioSliceFlag_InterruptAtLoop:
		return "KScheduledAudioSliceFlag_InterruptAtLoop"
	case KScheduledAudioSliceFlag_Loop:
		return "KScheduledAudioSliceFlag_Loop"
	default:
		return fmt.Sprintf("AUScheduledAudioSliceFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve
type AUSpatialMixerAttenuationCurve uint32

const (
	KSpatialMixerAttenuationCurve_Exponential AUSpatialMixerAttenuationCurve = 1
	KSpatialMixerAttenuationCurve_Inverse     AUSpatialMixerAttenuationCurve = 2
	KSpatialMixerAttenuationCurve_Linear      AUSpatialMixerAttenuationCurve = 3
	KSpatialMixerAttenuationCurve_Power       AUSpatialMixerAttenuationCurve = 0
)

func (e AUSpatialMixerAttenuationCurve) String() string {
	switch e {
	case KSpatialMixerAttenuationCurve_Exponential:
		return "KSpatialMixerAttenuationCurve_Exponential"
	case KSpatialMixerAttenuationCurve_Inverse:
		return "KSpatialMixerAttenuationCurve_Inverse"
	case KSpatialMixerAttenuationCurve_Linear:
		return "KSpatialMixerAttenuationCurve_Linear"
	case KSpatialMixerAttenuationCurve_Power:
		return "KSpatialMixerAttenuationCurve_Power"
	default:
		return fmt.Sprintf("AUSpatialMixerAttenuationCurve(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerOutputType
type AUSpatialMixerOutputType uint32

const (
	KSpatialMixerOutputType_BuiltInSpeakers AUSpatialMixerOutputType = 2
	KSpatialMixerOutputType_Headphones      AUSpatialMixerOutputType = 1
)

func (e AUSpatialMixerOutputType) String() string {
	switch e {
	case KSpatialMixerOutputType_BuiltInSpeakers:
		return "KSpatialMixerOutputType_BuiltInSpeakers"
	case KSpatialMixerOutputType_Headphones:
		return "KSpatialMixerOutputType_Headphones"
	default:
		return fmt.Sprintf("AUSpatialMixerOutputType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPersonalizedHRTFMode
type AUSpatialMixerPersonalizedHRTFMode uint32

const (
	KSpatialMixerPersonalizedHRTFMode_Auto AUSpatialMixerPersonalizedHRTFMode = 2
	KSpatialMixerPersonalizedHRTFMode_Off  AUSpatialMixerPersonalizedHRTFMode = 0
	KSpatialMixerPersonalizedHRTFMode_On   AUSpatialMixerPersonalizedHRTFMode = 1
)

func (e AUSpatialMixerPersonalizedHRTFMode) String() string {
	switch e {
	case KSpatialMixerPersonalizedHRTFMode_Auto:
		return "KSpatialMixerPersonalizedHRTFMode_Auto"
	case KSpatialMixerPersonalizedHRTFMode_Off:
		return "KSpatialMixerPersonalizedHRTFMode_Off"
	case KSpatialMixerPersonalizedHRTFMode_On:
		return "KSpatialMixerPersonalizedHRTFMode_On"
	default:
		return fmt.Sprintf("AUSpatialMixerPersonalizedHRTFMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPointSourceInHeadMode
type AUSpatialMixerPointSourceInHeadMode uint32

const (
	KSpatialMixerPointSourceInHeadMode_Bypass AUSpatialMixerPointSourceInHeadMode = 1
	KSpatialMixerPointSourceInHeadMode_Mono   AUSpatialMixerPointSourceInHeadMode = 0
)

func (e AUSpatialMixerPointSourceInHeadMode) String() string {
	switch e {
	case KSpatialMixerPointSourceInHeadMode_Bypass:
		return "KSpatialMixerPointSourceInHeadMode_Bypass"
	case KSpatialMixerPointSourceInHeadMode_Mono:
		return "KSpatialMixerPointSourceInHeadMode_Mono"
	default:
		return fmt.Sprintf("AUSpatialMixerPointSourceInHeadMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerRenderingFlags
type AUSpatialMixerRenderingFlags uint32

const (
	KSpatialMixerRenderingFlags_DistanceAttenuation AUSpatialMixerRenderingFlags = 4
	KSpatialMixerRenderingFlags_InterAuralDelay     AUSpatialMixerRenderingFlags = 1
)

func (e AUSpatialMixerRenderingFlags) String() string {
	switch e {
	case KSpatialMixerRenderingFlags_DistanceAttenuation:
		return "KSpatialMixerRenderingFlags_DistanceAttenuation"
	case KSpatialMixerRenderingFlags_InterAuralDelay:
		return "KSpatialMixerRenderingFlags_InterAuralDelay"
	default:
		return fmt.Sprintf("AUSpatialMixerRenderingFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode
type AUSpatialMixerSourceMode uint32

const (
	KSpatialMixerSourceMode_AmbienceBed      AUSpatialMixerSourceMode = 3
	KSpatialMixerSourceMode_Bypass           AUSpatialMixerSourceMode = 1
	KSpatialMixerSourceMode_PointSource      AUSpatialMixerSourceMode = 2
	KSpatialMixerSourceMode_SpatializeIfMono AUSpatialMixerSourceMode = 0
)

func (e AUSpatialMixerSourceMode) String() string {
	switch e {
	case KSpatialMixerSourceMode_AmbienceBed:
		return "KSpatialMixerSourceMode_AmbienceBed"
	case KSpatialMixerSourceMode_Bypass:
		return "KSpatialMixerSourceMode_Bypass"
	case KSpatialMixerSourceMode_PointSource:
		return "KSpatialMixerSourceMode_PointSource"
	case KSpatialMixerSourceMode_SpatializeIfMono:
		return "KSpatialMixerSourceMode_SpatializeIfMono"
	default:
		return fmt.Sprintf("AUSpatialMixerSourceMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUSpatializationAlgorithm
type AUSpatializationAlgorithm uint32

const (
	KSpatializationAlgorithm_EqualPowerPanning  AUSpatializationAlgorithm = 0
	KSpatializationAlgorithm_HRTF               AUSpatializationAlgorithm = 2
	KSpatializationAlgorithm_HRTFHQ             AUSpatializationAlgorithm = 6
	KSpatializationAlgorithm_SoundField         AUSpatializationAlgorithm = 3
	KSpatializationAlgorithm_SphericalHead      AUSpatializationAlgorithm = 1
	KSpatializationAlgorithm_StereoPassThrough  AUSpatializationAlgorithm = 5
	KSpatializationAlgorithm_UseOutputType      AUSpatializationAlgorithm = 7
	KSpatializationAlgorithm_VectorBasedPanning AUSpatializationAlgorithm = 4
)

func (e AUSpatializationAlgorithm) String() string {
	switch e {
	case KSpatializationAlgorithm_EqualPowerPanning:
		return "KSpatializationAlgorithm_EqualPowerPanning"
	case KSpatializationAlgorithm_HRTF:
		return "KSpatializationAlgorithm_HRTF"
	case KSpatializationAlgorithm_HRTFHQ:
		return "KSpatializationAlgorithm_HRTFHQ"
	case KSpatializationAlgorithm_SoundField:
		return "KSpatializationAlgorithm_SoundField"
	case KSpatializationAlgorithm_SphericalHead:
		return "KSpatializationAlgorithm_SphericalHead"
	case KSpatializationAlgorithm_StereoPassThrough:
		return "KSpatializationAlgorithm_StereoPassThrough"
	case KSpatializationAlgorithm_UseOutputType:
		return "KSpatializationAlgorithm_UseOutputType"
	case KSpatializationAlgorithm_VectorBasedPanning:
		return "KSpatializationAlgorithm_VectorBasedPanning"
	default:
		return fmt.Sprintf("AUSpatializationAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingLevel
type AUVoiceIOOtherAudioDuckingLevel uint32

const (
	// KAUVoiceIOOtherAudioDuckingLevelDefault: The default ducking level of other non-voice audio in a typical voice chat.
	KAUVoiceIOOtherAudioDuckingLevelDefault AUVoiceIOOtherAudioDuckingLevel = 0
	// KAUVoiceIOOtherAudioDuckingLevelMax: The maximum ducking level of other non-voice audio.
	KAUVoiceIOOtherAudioDuckingLevelMax AUVoiceIOOtherAudioDuckingLevel = 30
	// KAUVoiceIOOtherAudioDuckingLevelMid: A medium ducking level of other non-voice audio.
	KAUVoiceIOOtherAudioDuckingLevelMid AUVoiceIOOtherAudioDuckingLevel = 20
	// KAUVoiceIOOtherAudioDuckingLevelMin: The minimum ducking level of other non-voice audio.
	KAUVoiceIOOtherAudioDuckingLevelMin AUVoiceIOOtherAudioDuckingLevel = 10
)

func (e AUVoiceIOOtherAudioDuckingLevel) String() string {
	switch e {
	case KAUVoiceIOOtherAudioDuckingLevelDefault:
		return "KAUVoiceIOOtherAudioDuckingLevelDefault"
	case KAUVoiceIOOtherAudioDuckingLevelMax:
		return "KAUVoiceIOOtherAudioDuckingLevelMax"
	case KAUVoiceIOOtherAudioDuckingLevelMid:
		return "KAUVoiceIOOtherAudioDuckingLevelMid"
	case KAUVoiceIOOtherAudioDuckingLevelMin:
		return "KAUVoiceIOOtherAudioDuckingLevelMin"
	default:
		return fmt.Sprintf("AUVoiceIOOtherAudioDuckingLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOSpeechActivityEvent
type AUVoiceIOSpeechActivityEvent uint32

const (
	// KAUVoiceIOSpeechActivityHasEnded: A state that indicates speech ended.
	KAUVoiceIOSpeechActivityHasEnded AUVoiceIOSpeechActivityEvent = 1
	// KAUVoiceIOSpeechActivityHasStarted: A state that indicates speech started.
	KAUVoiceIOSpeechActivityHasStarted AUVoiceIOSpeechActivityEvent = 0
)

func (e AUVoiceIOSpeechActivityEvent) String() string {
	switch e {
	case KAUVoiceIOSpeechActivityHasEnded:
		return "KAUVoiceIOSpeechActivityHasEnded"
	case KAUVoiceIOSpeechActivityHasStarted:
		return "KAUVoiceIOSpeechActivityHasStarted"
	default:
		return fmt.Sprintf("AUVoiceIOSpeechActivityEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFadeType
type AudioBalanceFadeType uint32

const (
	// KAudioBalanceFadeType_EqualPower: Overall loudness remains constant, but gain can exceed 1.0.
	KAudioBalanceFadeType_EqualPower AudioBalanceFadeType = 1
	// KAudioBalanceFadeType_MaxUnityGain: Ensures that the overall gain value never exceeds 1.0 by fading one channel as the other channel’s level rises.
	KAudioBalanceFadeType_MaxUnityGain AudioBalanceFadeType = 0
)

func (e AudioBalanceFadeType) String() string {
	switch e {
	case KAudioBalanceFadeType_EqualPower:
		return "KAudioBalanceFadeType_EqualPower"
	case KAudioBalanceFadeType_MaxUnityGain:
		return "KAudioBalanceFadeType_MaxUnityGain"
	default:
		return fmt.Sprintf("AudioBalanceFadeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslationFlags
type AudioBytePacketTranslationFlags uint32

const (
	// KBytePacketTranslationFlag_IsEstimate: If set, the result value is an estimate.
	KBytePacketTranslationFlag_IsEstimate AudioBytePacketTranslationFlags = 1
)

func (e AudioBytePacketTranslationFlags) String() string {
	switch e {
	case KBytePacketTranslationFlag_IsEstimate:
		return "KBytePacketTranslationFlag_IsEstimate"
	default:
		return fmt.Sprintf("AudioBytePacketTranslationFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags
type AudioComponentFlags uint32

const (
	KAudioComponentFlag_CanLoadInProcess           AudioComponentFlags = 0x10
	KAudioComponentFlag_IsV3AudioUnit              AudioComponentFlags = 4
	KAudioComponentFlag_RequiresAsyncInstantiation AudioComponentFlags = 8
	KAudioComponentFlag_SandboxSafe                AudioComponentFlags = 2
	KAudioComponentFlag_Unsearchable               AudioComponentFlags = 1
)

func (e AudioComponentFlags) String() string {
	switch e {
	case KAudioComponentFlag_CanLoadInProcess:
		return "KAudioComponentFlag_CanLoadInProcess"
	case KAudioComponentFlag_IsV3AudioUnit:
		return "KAudioComponentFlag_IsV3AudioUnit"
	case KAudioComponentFlag_RequiresAsyncInstantiation:
		return "KAudioComponentFlag_RequiresAsyncInstantiation"
	case KAudioComponentFlag_SandboxSafe:
		return "KAudioComponentFlag_SandboxSafe"
	case KAudioComponentFlag_Unsearchable:
		return "KAudioComponentFlag_Unsearchable"
	default:
		return fmt.Sprintf("AudioComponentFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiationOptions
type AudioComponentInstantiationOptions uint32

const (
	KAudioComponentInstantiation_LoadInProcess    AudioComponentInstantiationOptions = 2
	KAudioComponentInstantiation_LoadOutOfProcess AudioComponentInstantiationOptions = 1
	KAudioComponentInstantiation_LoadedRemotely   AudioComponentInstantiationOptions = 2147483648
)

func (e AudioComponentInstantiationOptions) String() string {
	switch e {
	case KAudioComponentInstantiation_LoadInProcess:
		return "KAudioComponentInstantiation_LoadInProcess"
	case KAudioComponentInstantiation_LoadOutOfProcess:
		return "KAudioComponentInstantiation_LoadOutOfProcess"
	case KAudioComponentInstantiation_LoadedRemotely:
		return "KAudioComponentInstantiation_LoadedRemotely"
	default:
		return fmt.Sprintf("AudioComponentInstantiationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult
type AudioComponentValidationResult uint32

const (
	KAudioComponentValidationResult_Failed                 AudioComponentValidationResult = 2
	KAudioComponentValidationResult_Passed                 AudioComponentValidationResult = 1
	KAudioComponentValidationResult_TimedOut               AudioComponentValidationResult = 3
	KAudioComponentValidationResult_UnauthorizedError_Init AudioComponentValidationResult = 5
	KAudioComponentValidationResult_UnauthorizedError_Open AudioComponentValidationResult = 4
	KAudioComponentValidationResult_Unknown                AudioComponentValidationResult = 0
)

func (e AudioComponentValidationResult) String() string {
	switch e {
	case KAudioComponentValidationResult_Failed:
		return "KAudioComponentValidationResult_Failed"
	case KAudioComponentValidationResult_Passed:
		return "KAudioComponentValidationResult_Passed"
	case KAudioComponentValidationResult_TimedOut:
		return "KAudioComponentValidationResult_TimedOut"
	case KAudioComponentValidationResult_UnauthorizedError_Init:
		return "KAudioComponentValidationResult_UnauthorizedError_Init"
	case KAudioComponentValidationResult_UnauthorizedError_Open:
		return "KAudioComponentValidationResult_UnauthorizedError_Open"
	case KAudioComponentValidationResult_Unknown:
		return "KAudioComponentValidationResult_Unknown"
	default:
		return fmt.Sprintf("AudioComponentValidationResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterOptions
type AudioConverterOptions uint32

const (
	KAudioConverterOption_Unbuffered AudioConverterOptions = 65536
)

func (e AudioConverterOptions) String() string {
	switch e {
	case KAudioConverterOption_Unbuffered:
		return "KAudioConverterOption_Unbuffered"
	default:
		return fmt.Sprintf("AudioConverterOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileFlags
type AudioFileFlags uint32

const (
	// KAudioFileFlags_DontPageAlignAudioData: Typically, the audio data in a file is page aligned.
	KAudioFileFlags_DontPageAlignAudioData AudioFileFlags = 2
	// KAudioFileFlags_EraseFile: # Discussion
	KAudioFileFlags_EraseFile AudioFileFlags = 1
)

func (e AudioFileFlags) String() string {
	switch e {
	case KAudioFileFlags_DontPageAlignAudioData:
		return "KAudioFileFlags_DontPageAlignAudioData"
	case KAudioFileFlags_EraseFile:
		return "KAudioFileFlags_EraseFile"
	default:
		return fmt.Sprintf("AudioFileFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFilePermissions
type AudioFilePermissions int8

const (
	// KAudioFileReadPermission: File is read-only.
	KAudioFileReadPermission AudioFilePermissions = 0x1
	// KAudioFileReadWritePermission: File has read-write permission.
	KAudioFileReadWritePermission AudioFilePermissions = 0x3
	// KAudioFileWritePermission: File is write-only.
	KAudioFileWritePermission AudioFilePermissions = 0x2
)

func (e AudioFilePermissions) String() string {
	switch e {
	case KAudioFileReadPermission:
		return "KAudioFileReadPermission"
	case KAudioFileReadWritePermission:
		return "KAudioFileReadWritePermission"
	case KAudioFileWritePermission:
		return "KAudioFileWritePermission"
	default:
		return fmt.Sprintf("AudioFilePermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionFlags
type AudioFileRegionFlags uint32

const (
	// KAudioFileRegionFlag_LoopEnable: If set, the region is looped.
	KAudioFileRegionFlag_LoopEnable AudioFileRegionFlags = 1
	// KAudioFileRegionFlag_PlayBackward: If set, the region is played backward.
	KAudioFileRegionFlag_PlayBackward AudioFileRegionFlags = 4
	// KAudioFileRegionFlag_PlayForward: If set, the region is played forward.
	KAudioFileRegionFlag_PlayForward AudioFileRegionFlags = 2
)

func (e AudioFileRegionFlags) String() string {
	switch e {
	case KAudioFileRegionFlag_LoopEnable:
		return "KAudioFileRegionFlag_LoopEnable"
	case KAudioFileRegionFlag_PlayBackward:
		return "KAudioFileRegionFlag_PlayBackward"
	case KAudioFileRegionFlag_PlayForward:
		return "KAudioFileRegionFlag_PlayForward"
	default:
		return fmt.Sprintf("AudioFileRegionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseFlags
type AudioFileStreamParseFlags uint32

const (
	// KAudioFileStreamParseFlag_Discontinuity: Pass this flag to the AudioFileStreamParseBytes(_:_:_:_:) function to signal a discontinuity in the audio data.
	KAudioFileStreamParseFlag_Discontinuity AudioFileStreamParseFlags = 1
)

func (e AudioFileStreamParseFlags) String() string {
	switch e {
	case KAudioFileStreamParseFlag_Discontinuity:
		return "KAudioFileStreamParseFlag_Discontinuity"
	default:
		return fmt.Sprintf("AudioFileStreamParseFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyFlags
type AudioFileStreamPropertyFlags uint32

const (
	// KAudioFileStreamPropertyFlag_CacheProperty: A property listener sets this flag to instruct the parser to cache the property value so that it remains available after the callback returns.
	KAudioFileStreamPropertyFlag_CacheProperty AudioFileStreamPropertyFlags = 2
	// KAudioFileStreamPropertyFlag_PropertyIsCached: This flag is set when the callback AudioFileStream_PropertyListenerProc is invoked in the case that the value of the property has been cached and can be obtained later.
	KAudioFileStreamPropertyFlag_PropertyIsCached AudioFileStreamPropertyFlags = 1
)

func (e AudioFileStreamPropertyFlags) String() string {
	switch e {
	case KAudioFileStreamPropertyFlag_CacheProperty:
		return "KAudioFileStreamPropertyFlag_CacheProperty"
	case KAudioFileStreamPropertyFlag_PropertyIsCached:
		return "KAudioFileStreamPropertyFlag_PropertyIsCached"
	default:
		return fmt.Sprintf("AudioFileStreamPropertyFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeekFlags
type AudioFileStreamSeekFlags uint32

const (
	// KAudioFileStreamSeekFlag_OffsetIsEstimated: This flag is returned by the AudioFileStreamSeek(_:_:_:_:) function if the byte offset is only an estimate.
	KAudioFileStreamSeekFlag_OffsetIsEstimated AudioFileStreamSeekFlags = 1
)

func (e AudioFileStreamSeekFlags) String() string {
	switch e {
	case KAudioFileStreamSeekFlag_OffsetIsEstimated:
		return "KAudioFileStreamSeekFlag_OffsetIsEstimated"
	default:
		return fmt.Sprintf("AudioFileStreamSeekFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioPanningMode
type AudioPanningMode uint32

const (
	// KPanningMode_SoundField: The SoundField panning algorithm.
	KPanningMode_SoundField AudioPanningMode = 3
	// KPanningMode_VectorBasedPanning: A vector-based panning algorithm.
	KPanningMode_VectorBasedPanning AudioPanningMode = 4
)

func (e AudioPanningMode) String() string {
	switch e {
	case KPanningMode_SoundField:
		return "KPanningMode_SoundField"
	case KPanningMode_VectorBasedPanning:
		return "KPanningMode_VectorBasedPanning"
	default:
		return fmt.Sprintf("AudioPanningMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapFlags
type AudioQueueProcessingTapFlags uint32

const (
	KAudioQueueProcessingTap_EndOfStream   AudioQueueProcessingTapFlags = 512
	KAudioQueueProcessingTap_PostEffects   AudioQueueProcessingTapFlags = 2
	KAudioQueueProcessingTap_PreEffects    AudioQueueProcessingTapFlags = 1
	KAudioQueueProcessingTap_Siphon        AudioQueueProcessingTapFlags = 4
	KAudioQueueProcessingTap_StartOfStream AudioQueueProcessingTapFlags = 256
)

func (e AudioQueueProcessingTapFlags) String() string {
	switch e {
	case KAudioQueueProcessingTap_EndOfStream:
		return "KAudioQueueProcessingTap_EndOfStream"
	case KAudioQueueProcessingTap_PostEffects:
		return "KAudioQueueProcessingTap_PostEffects"
	case KAudioQueueProcessingTap_PreEffects:
		return "KAudioQueueProcessingTap_PreEffects"
	case KAudioQueueProcessingTap_Siphon:
		return "KAudioQueueProcessingTap_Siphon"
	case KAudioQueueProcessingTap_StartOfStream:
		return "KAudioQueueProcessingTap_StartOfStream"
	default:
		return fmt.Sprintf("AudioQueueProcessingTapFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags
type AudioSettingsFlags uint32

const (
	KAudioSettingsFlags_ExpertParameter        AudioSettingsFlags = 1
	KAudioSettingsFlags_InvisibleParameter     AudioSettingsFlags = 2
	KAudioSettingsFlags_MetaParameter          AudioSettingsFlags = 4
	KAudioSettingsFlags_UserInterfaceParameter AudioSettingsFlags = 8
)

func (e AudioSettingsFlags) String() string {
	switch e {
	case KAudioSettingsFlags_ExpertParameter:
		return "KAudioSettingsFlags_ExpertParameter"
	case KAudioSettingsFlags_InvisibleParameter:
		return "KAudioSettingsFlags_InvisibleParameter"
	case KAudioSettingsFlags_MetaParameter:
		return "KAudioSettingsFlags_MetaParameter"
	case KAudioSettingsFlags_UserInterfaceParameter:
		return "KAudioSettingsFlags_UserInterfaceParameter"
	default:
		return fmt.Sprintf("AudioSettingsFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType
type AudioUnitEventType uint32

const (
	KAudioUnitEvent_BeginParameterChangeGesture AudioUnitEventType = 1
	KAudioUnitEvent_EndParameterChangeGesture   AudioUnitEventType = 2
	KAudioUnitEvent_ParameterValueChange        AudioUnitEventType = 0
	KAudioUnitEvent_PropertyChange              AudioUnitEventType = 3
)

func (e AudioUnitEventType) String() string {
	switch e {
	case KAudioUnitEvent_BeginParameterChangeGesture:
		return "KAudioUnitEvent_BeginParameterChangeGesture"
	case KAudioUnitEvent_EndParameterChangeGesture:
		return "KAudioUnitEvent_EndParameterChangeGesture"
	case KAudioUnitEvent_ParameterValueChange:
		return "KAudioUnitEvent_ParameterValueChange"
	case KAudioUnitEvent_PropertyChange:
		return "KAudioUnitEvent_PropertyChange"
	default:
		return fmt.Sprintf("AudioUnitEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterOptions
type AudioUnitParameterOptions uint32

const (
	// KAudioUnitParameterFlag_CFNameRelease: If an audio unit can generate parameter names dynamically, it should set this flag.
	KAudioUnitParameterFlag_CFNameRelease      AudioUnitParameterOptions = 16
	KAudioUnitParameterFlag_CanRamp            AudioUnitParameterOptions = 33554432
	KAudioUnitParameterFlag_DisplayCubeRoot    AudioUnitParameterOptions = 262144
	KAudioUnitParameterFlag_DisplayCubed       AudioUnitParameterOptions = 196608
	KAudioUnitParameterFlag_DisplayExponential AudioUnitParameterOptions = 327680
	KAudioUnitParameterFlag_DisplayLogarithmic AudioUnitParameterOptions = 4194304
	KAudioUnitParameterFlag_DisplayMask        AudioUnitParameterOptions = 4653056
	KAudioUnitParameterFlag_DisplaySquareRoot  AudioUnitParameterOptions = 65536
	KAudioUnitParameterFlag_DisplaySquared     AudioUnitParameterOptions = 131072
	KAudioUnitParameterFlag_ExpertMode         AudioUnitParameterOptions = 67108864
	KAudioUnitParameterFlag_HasCFNameString    AudioUnitParameterOptions = 134217728
	KAudioUnitParameterFlag_HasClump           AudioUnitParameterOptions = 1048576
	KAudioUnitParameterFlag_IsElementMeta      AudioUnitParameterOptions = 536870912
	KAudioUnitParameterFlag_IsGlobalMeta       AudioUnitParameterOptions = 268435456
	KAudioUnitParameterFlag_IsHighResolution   AudioUnitParameterOptions = 8388608
	KAudioUnitParameterFlag_IsReadable         AudioUnitParameterOptions = 1073741824
	KAudioUnitParameterFlag_IsWritable         AudioUnitParameterOptions = 2147483648
	KAudioUnitParameterFlag_MeterReadOnly      AudioUnitParameterOptions = 32768
	KAudioUnitParameterFlag_NonRealTime        AudioUnitParameterOptions = 16777216
	KAudioUnitParameterFlag_OmitFromPresets    AudioUnitParameterOptions = 8192
	// KAudioUnitParameterFlag_PlotHistory: If set, getting the `kAudioUnitProperty_ParameterHistoryInfo` property fills out the AudioUnitParameterHistoryInfo struct containing the recommended update rate and history duration.
	KAudioUnitParameterFlag_PlotHistory       AudioUnitParameterOptions = 16384
	KAudioUnitParameterFlag_ValuesHaveStrings AudioUnitParameterOptions = 2097152
)

func (e AudioUnitParameterOptions) String() string {
	switch e {
	case KAudioUnitParameterFlag_CFNameRelease:
		return "KAudioUnitParameterFlag_CFNameRelease"
	case KAudioUnitParameterFlag_CanRamp:
		return "KAudioUnitParameterFlag_CanRamp"
	case KAudioUnitParameterFlag_DisplayCubeRoot:
		return "KAudioUnitParameterFlag_DisplayCubeRoot"
	case KAudioUnitParameterFlag_DisplayCubed:
		return "KAudioUnitParameterFlag_DisplayCubed"
	case KAudioUnitParameterFlag_DisplayExponential:
		return "KAudioUnitParameterFlag_DisplayExponential"
	case KAudioUnitParameterFlag_DisplayLogarithmic:
		return "KAudioUnitParameterFlag_DisplayLogarithmic"
	case KAudioUnitParameterFlag_DisplayMask:
		return "KAudioUnitParameterFlag_DisplayMask"
	case KAudioUnitParameterFlag_DisplaySquareRoot:
		return "KAudioUnitParameterFlag_DisplaySquareRoot"
	case KAudioUnitParameterFlag_DisplaySquared:
		return "KAudioUnitParameterFlag_DisplaySquared"
	case KAudioUnitParameterFlag_ExpertMode:
		return "KAudioUnitParameterFlag_ExpertMode"
	case KAudioUnitParameterFlag_HasCFNameString:
		return "KAudioUnitParameterFlag_HasCFNameString"
	case KAudioUnitParameterFlag_HasClump:
		return "KAudioUnitParameterFlag_HasClump"
	case KAudioUnitParameterFlag_IsElementMeta:
		return "KAudioUnitParameterFlag_IsElementMeta"
	case KAudioUnitParameterFlag_IsGlobalMeta:
		return "KAudioUnitParameterFlag_IsGlobalMeta"
	case KAudioUnitParameterFlag_IsHighResolution:
		return "KAudioUnitParameterFlag_IsHighResolution"
	case KAudioUnitParameterFlag_IsReadable:
		return "KAudioUnitParameterFlag_IsReadable"
	case KAudioUnitParameterFlag_IsWritable:
		return "KAudioUnitParameterFlag_IsWritable"
	case KAudioUnitParameterFlag_MeterReadOnly:
		return "KAudioUnitParameterFlag_MeterReadOnly"
	case KAudioUnitParameterFlag_NonRealTime:
		return "KAudioUnitParameterFlag_NonRealTime"
	case KAudioUnitParameterFlag_OmitFromPresets:
		return "KAudioUnitParameterFlag_OmitFromPresets"
	case KAudioUnitParameterFlag_PlotHistory:
		return "KAudioUnitParameterFlag_PlotHistory"
	case KAudioUnitParameterFlag_ValuesHaveStrings:
		return "KAudioUnitParameterFlag_ValuesHaveStrings"
	default:
		return fmt.Sprintf("AudioUnitParameterOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit
type AudioUnitParameterUnit uint32

const (
	// KAudioUnitParameterUnit_AbsoluteCents: An absolute unit of measure for the musical pitch of a note.
	KAudioUnitParameterUnit_AbsoluteCents AudioUnitParameterUnit = 20
	// KAudioUnitParameterUnit_BPM: A whole-number unit of measure for musical tempo, representing beats per minute.
	KAudioUnitParameterUnit_BPM AudioUnitParameterUnit = 22
	// KAudioUnitParameterUnit_Beats: A time unit of measure in musical beats.
	KAudioUnitParameterUnit_Beats AudioUnitParameterUnit = 23
	// KAudioUnitParameterUnit_Boolean: A Boolean-like unit of measure.
	KAudioUnitParameterUnit_Boolean AudioUnitParameterUnit = 2
	// KAudioUnitParameterUnit_Cents: A logarithmic unit of measure for a musical interval between two notes.
	KAudioUnitParameterUnit_Cents AudioUnitParameterUnit = 9
	// KAudioUnitParameterUnit_CustomUnit: A custom unit of measure.
	KAudioUnitParameterUnit_CustomUnit AudioUnitParameterUnit = 26
	// KAudioUnitParameterUnit_Decibels: A logarithmic unit of measure representing the ratio between two audio levels.
	KAudioUnitParameterUnit_Decibels AudioUnitParameterUnit = 13
	// KAudioUnitParameterUnit_Degrees: An angular degree unit of measure.
	KAudioUnitParameterUnit_Degrees AudioUnitParameterUnit = 15
	// KAudioUnitParameterUnit_EqualPowerCrossfade: An audio power unit of measure.
	KAudioUnitParameterUnit_EqualPowerCrossfade AudioUnitParameterUnit = 16
	// KAudioUnitParameterUnit_Generic: A generic unit of measure.
	KAudioUnitParameterUnit_Generic AudioUnitParameterUnit = 0
	// KAudioUnitParameterUnit_Hertz: A hertz unit of measure.
	KAudioUnitParameterUnit_Hertz AudioUnitParameterUnit = 8
	// KAudioUnitParameterUnit_Indexed: An indexed unit of measure.
	KAudioUnitParameterUnit_Indexed AudioUnitParameterUnit = 1
	// KAudioUnitParameterUnit_LinearGain: A linear unit of measure representing the difference between two audio levels.
	KAudioUnitParameterUnit_LinearGain      AudioUnitParameterUnit = 14
	KAudioUnitParameterUnit_MIDI2Controller AudioUnitParameterUnit = 27
	// KAudioUnitParameterUnit_MIDIController: A whole-number unit of measure corresponding to standard MIDI control numbers.
	KAudioUnitParameterUnit_MIDIController AudioUnitParameterUnit = 12
	// KAudioUnitParameterUnit_MIDINoteNumber: A whole-number unit of measure corresponding to audio frequency.
	KAudioUnitParameterUnit_MIDINoteNumber AudioUnitParameterUnit = 11
	// KAudioUnitParameterUnit_Meters: A distance unit of measure, corresponding to meters.
	KAudioUnitParameterUnit_Meters AudioUnitParameterUnit = 19
	// KAudioUnitParameterUnit_Milliseconds: A time unit of measure representing milliseconds.
	KAudioUnitParameterUnit_Milliseconds AudioUnitParameterUnit = 24
	// KAudioUnitParameterUnit_MixerFaderCurve1: An audio power unit of measure.
	KAudioUnitParameterUnit_MixerFaderCurve1 AudioUnitParameterUnit = 17
	// KAudioUnitParameterUnit_Octaves: A relative unit of measure for the musical interval between two notes.
	KAudioUnitParameterUnit_Octaves AudioUnitParameterUnit = 21
	// KAudioUnitParameterUnit_Pan: An audio position unit of measure.
	KAudioUnitParameterUnit_Pan AudioUnitParameterUnit = 18
	// KAudioUnitParameterUnit_Percent: A percentage unit of measure.
	KAudioUnitParameterUnit_Percent AudioUnitParameterUnit = 3
	// KAudioUnitParameterUnit_Phase: An angular degree unit of measure.
	KAudioUnitParameterUnit_Phase AudioUnitParameterUnit = 6
	// KAudioUnitParameterUnit_Rate: A multiplication factor unit of measure.
	KAudioUnitParameterUnit_Rate AudioUnitParameterUnit = 7
	// KAudioUnitParameterUnit_Ratio: A unitless ratio unit of measure.
	KAudioUnitParameterUnit_Ratio AudioUnitParameterUnit = 25
	// KAudioUnitParameterUnit_RelativeSemiTones: A relative unit of measure for a musical interval between two notes.
	KAudioUnitParameterUnit_RelativeSemiTones AudioUnitParameterUnit = 10
	// KAudioUnitParameterUnit_SampleFrames: A sample-frame-count unit of measure.
	KAudioUnitParameterUnit_SampleFrames AudioUnitParameterUnit = 5
	// KAudioUnitParameterUnit_Seconds: A whole-seconds unit of measure, indicating either absolute or relative time.
	KAudioUnitParameterUnit_Seconds AudioUnitParameterUnit = 4
)

func (e AudioUnitParameterUnit) String() string {
	switch e {
	case KAudioUnitParameterUnit_AbsoluteCents:
		return "KAudioUnitParameterUnit_AbsoluteCents"
	case KAudioUnitParameterUnit_BPM:
		return "KAudioUnitParameterUnit_BPM"
	case KAudioUnitParameterUnit_Beats:
		return "KAudioUnitParameterUnit_Beats"
	case KAudioUnitParameterUnit_Boolean:
		return "KAudioUnitParameterUnit_Boolean"
	case KAudioUnitParameterUnit_Cents:
		return "KAudioUnitParameterUnit_Cents"
	case KAudioUnitParameterUnit_CustomUnit:
		return "KAudioUnitParameterUnit_CustomUnit"
	case KAudioUnitParameterUnit_Decibels:
		return "KAudioUnitParameterUnit_Decibels"
	case KAudioUnitParameterUnit_Degrees:
		return "KAudioUnitParameterUnit_Degrees"
	case KAudioUnitParameterUnit_EqualPowerCrossfade:
		return "KAudioUnitParameterUnit_EqualPowerCrossfade"
	case KAudioUnitParameterUnit_Generic:
		return "KAudioUnitParameterUnit_Generic"
	case KAudioUnitParameterUnit_Hertz:
		return "KAudioUnitParameterUnit_Hertz"
	case KAudioUnitParameterUnit_Indexed:
		return "KAudioUnitParameterUnit_Indexed"
	case KAudioUnitParameterUnit_LinearGain:
		return "KAudioUnitParameterUnit_LinearGain"
	case KAudioUnitParameterUnit_MIDI2Controller:
		return "KAudioUnitParameterUnit_MIDI2Controller"
	case KAudioUnitParameterUnit_MIDIController:
		return "KAudioUnitParameterUnit_MIDIController"
	case KAudioUnitParameterUnit_MIDINoteNumber:
		return "KAudioUnitParameterUnit_MIDINoteNumber"
	case KAudioUnitParameterUnit_Meters:
		return "KAudioUnitParameterUnit_Meters"
	case KAudioUnitParameterUnit_Milliseconds:
		return "KAudioUnitParameterUnit_Milliseconds"
	case KAudioUnitParameterUnit_MixerFaderCurve1:
		return "KAudioUnitParameterUnit_MixerFaderCurve1"
	case KAudioUnitParameterUnit_Octaves:
		return "KAudioUnitParameterUnit_Octaves"
	case KAudioUnitParameterUnit_Pan:
		return "KAudioUnitParameterUnit_Pan"
	case KAudioUnitParameterUnit_Percent:
		return "KAudioUnitParameterUnit_Percent"
	case KAudioUnitParameterUnit_Phase:
		return "KAudioUnitParameterUnit_Phase"
	case KAudioUnitParameterUnit_Rate:
		return "KAudioUnitParameterUnit_Rate"
	case KAudioUnitParameterUnit_Ratio:
		return "KAudioUnitParameterUnit_Ratio"
	case KAudioUnitParameterUnit_RelativeSemiTones:
		return "KAudioUnitParameterUnit_RelativeSemiTones"
	case KAudioUnitParameterUnit_SampleFrames:
		return "KAudioUnitParameterUnit_SampleFrames"
	case KAudioUnitParameterUnit_Seconds:
		return "KAudioUnitParameterUnit_Seconds"
	default:
		return fmt.Sprintf("AudioUnitParameterUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEvent
type AudioUnitRemoteControlEvent uint32

const (
	KAudioUnitRemoteControlEvent_Rewind          AudioUnitRemoteControlEvent = 3
	KAudioUnitRemoteControlEvent_TogglePlayPause AudioUnitRemoteControlEvent = 1
	KAudioUnitRemoteControlEvent_ToggleRecord    AudioUnitRemoteControlEvent = 2
)

func (e AudioUnitRemoteControlEvent) String() string {
	switch e {
	case KAudioUnitRemoteControlEvent_Rewind:
		return "KAudioUnitRemoteControlEvent_Rewind"
	case KAudioUnitRemoteControlEvent_TogglePlayPause:
		return "KAudioUnitRemoteControlEvent_TogglePlayPause"
	case KAudioUnitRemoteControlEvent_ToggleRecord:
		return "KAudioUnitRemoteControlEvent_ToggleRecord"
	default:
		return fmt.Sprintf("AudioUnitRemoteControlEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderActionFlags
type AudioUnitRenderActionFlags uint32

const (
	// KAudioOfflineUnitRenderAction_Complete: This flag is set when an offline unit has completed either its preflight or performed render operation.
	KAudioOfflineUnitRenderAction_Complete AudioUnitRenderActionFlags = 128
	// KAudioOfflineUnitRenderAction_Preflight: This is used with offline audio units (of type `'auol'`).
	KAudioOfflineUnitRenderAction_Preflight AudioUnitRenderActionFlags = 32
	// KAudioOfflineUnitRenderAction_Render: Once an offline unit has been successfully preflighted, it is then put into its render mode.
	KAudioOfflineUnitRenderAction_Render AudioUnitRenderActionFlags = 64
	// KAudioUnitRenderAction_DoNotCheckRenderArgs: If this flag is set, then checks that are done on the arguments provided to render are not performed.
	KAudioUnitRenderAction_DoNotCheckRenderArgs AudioUnitRenderActionFlags = 512
	// KAudioUnitRenderAction_OutputIsSilence: This flag can be set in a render input callback (or in the audio unit’s render operation itself) and is used to indicate that the render buffer contains only silence.
	KAudioUnitRenderAction_OutputIsSilence AudioUnitRenderActionFlags = 16
	// KAudioUnitRenderAction_PostRender: Called on a render notification Proc - which is called either before or after the render operation of the audio unit.
	KAudioUnitRenderAction_PostRender AudioUnitRenderActionFlags = 8
	// KAudioUnitRenderAction_PostRenderError: If this flag is set on the post-render call an error was returned by the audio unit’s render operation.
	KAudioUnitRenderAction_PostRenderError AudioUnitRenderActionFlags = 256
	// KAudioUnitRenderAction_PreRender: Called on a render notification Proc - which is called either before or after the render operation of the audio unit.
	KAudioUnitRenderAction_PreRender AudioUnitRenderActionFlags = 4
)

func (e AudioUnitRenderActionFlags) String() string {
	switch e {
	case KAudioOfflineUnitRenderAction_Complete:
		return "KAudioOfflineUnitRenderAction_Complete"
	case KAudioOfflineUnitRenderAction_Preflight:
		return "KAudioOfflineUnitRenderAction_Preflight"
	case KAudioOfflineUnitRenderAction_Render:
		return "KAudioOfflineUnitRenderAction_Render"
	case KAudioUnitRenderAction_DoNotCheckRenderArgs:
		return "KAudioUnitRenderAction_DoNotCheckRenderArgs"
	case KAudioUnitRenderAction_OutputIsSilence:
		return "KAudioUnitRenderAction_OutputIsSilence"
	case KAudioUnitRenderAction_PostRender:
		return "KAudioUnitRenderAction_PostRender"
	case KAudioUnitRenderAction_PostRenderError:
		return "KAudioUnitRenderAction_PostRenderError"
	case KAudioUnitRenderAction_PreRender:
		return "KAudioUnitRenderAction_PreRender"
	default:
		return fmt.Sprintf("AudioUnitRenderActionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage
type CAClockMessage uint32

const (
	KCAClockMessage_Armed            CAClockMessage = 'a'<<24 | 'r'<<16 | 'm'<<8 | 'd' // 'armd'
	KCAClockMessage_Disarmed         CAClockMessage = 'd'<<24 | 'a'<<16 | 'r'<<8 | 'm' // 'darm'
	KCAClockMessage_PropertyChanged  CAClockMessage = 'p'<<24 | 'c'<<16 | 'h'<<8 | 'g' // 'pchg'
	KCAClockMessage_StartTimeSet     CAClockMessage = 's'<<24 | 't'<<16 | 'i'<<8 | 'm' // 'stim'
	KCAClockMessage_Started          CAClockMessage = 's'<<24 | 't'<<16 | 'r'<<8 | 't' // 'strt'
	KCAClockMessage_Stopped          CAClockMessage = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KCAClockMessage_WrongSMPTEFormat CAClockMessage = '?'<<24 | 's'<<16 | 'm'<<8 | 'p' // '?smp'
)

func (e CAClockMessage) String() string {
	switch e {
	case KCAClockMessage_Armed:
		return "KCAClockMessage_Armed"
	case KCAClockMessage_Disarmed:
		return "KCAClockMessage_Disarmed"
	case KCAClockMessage_PropertyChanged:
		return "KCAClockMessage_PropertyChanged"
	case KCAClockMessage_StartTimeSet:
		return "KCAClockMessage_StartTimeSet"
	case KCAClockMessage_Started:
		return "KCAClockMessage_Started"
	case KCAClockMessage_Stopped:
		return "KCAClockMessage_Stopped"
	case KCAClockMessage_WrongSMPTEFormat:
		return "KCAClockMessage_WrongSMPTEFormat"
	default:
		return fmt.Sprintf("CAClockMessage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID
type CAClockPropertyID uint32

const (
	KCAClockProperty_InternalTimebase      CAClockPropertyID = 'i'<<24 | 'n'<<16 | 't'<<8 | 'b' // 'intb'
	KCAClockProperty_MIDIClockDestinations CAClockPropertyID = 'm'<<24 | 'b'<<16 | 'c'<<8 | 'd' // 'mbcd'
	KCAClockProperty_MTCDestinations       CAClockPropertyID = 'm'<<24 | 't'<<16 | 'c'<<8 | 'd' // 'mtcd'
	KCAClockProperty_MTCFreewheelTime      CAClockPropertyID = 'm'<<24 | 't'<<16 | 'f'<<8 | 'w' // 'mtfw'
	KCAClockProperty_MeterTrack            CAClockPropertyID = 'm'<<24 | 'e'<<16 | 't'<<8 | 'r' // 'metr'
	KCAClockProperty_Name                  CAClockPropertyID = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	KCAClockProperty_SMPTEFormat           CAClockPropertyID = 's'<<24 | 'm'<<16 | 'p'<<8 | 'f' // 'smpf'
	KCAClockProperty_SMPTEOffset           CAClockPropertyID = 's'<<24 | 'm'<<16 | 'p'<<8 | 'o' // 'smpo'
	KCAClockProperty_SendMIDISPP           CAClockPropertyID = 'm'<<24 | 's'<<16 | 'p'<<8 | 'p' // 'mspp'
	KCAClockProperty_SyncMode              CAClockPropertyID = 's'<<24 | 'y'<<16 | 'n'<<8 | 'm' // 'synm'
	KCAClockProperty_SyncSource            CAClockPropertyID = 's'<<24 | 'y'<<16 | 'n'<<8 | 's' // 'syns'
	KCAClockProperty_TempoMap              CAClockPropertyID = 't'<<24 | 'm'<<16 | 'p'<<8 | 'o' // 'tmpo'
	KCAClockProperty_TimebaseSource        CAClockPropertyID = 'i'<<24 | 't'<<16 | 'b'<<8 | 's' // 'itbs'
)

func (e CAClockPropertyID) String() string {
	switch e {
	case KCAClockProperty_InternalTimebase:
		return "KCAClockProperty_InternalTimebase"
	case KCAClockProperty_MIDIClockDestinations:
		return "KCAClockProperty_MIDIClockDestinations"
	case KCAClockProperty_MTCDestinations:
		return "KCAClockProperty_MTCDestinations"
	case KCAClockProperty_MTCFreewheelTime:
		return "KCAClockProperty_MTCFreewheelTime"
	case KCAClockProperty_MeterTrack:
		return "KCAClockProperty_MeterTrack"
	case KCAClockProperty_Name:
		return "KCAClockProperty_Name"
	case KCAClockProperty_SMPTEFormat:
		return "KCAClockProperty_SMPTEFormat"
	case KCAClockProperty_SMPTEOffset:
		return "KCAClockProperty_SMPTEOffset"
	case KCAClockProperty_SendMIDISPP:
		return "KCAClockProperty_SendMIDISPP"
	case KCAClockProperty_SyncMode:
		return "KCAClockProperty_SyncMode"
	case KCAClockProperty_SyncSource:
		return "KCAClockProperty_SyncSource"
	case KCAClockProperty_TempoMap:
		return "KCAClockProperty_TempoMap"
	case KCAClockProperty_TimebaseSource:
		return "KCAClockProperty_TimebaseSource"
	default:
		return fmt.Sprintf("CAClockPropertyID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSyncMode
type CAClockSyncMode uint32

const (
	KCAClockSyncMode_Internal           CAClockSyncMode = 'i'<<24 | 'n'<<16 | 't'<<8 | 'r' // 'intr'
	KCAClockSyncMode_MIDIClockTransport CAClockSyncMode = 'm'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'mclk'
	KCAClockSyncMode_MTCTransport       CAClockSyncMode = 'm'<<24 | 'm'<<16 | 't'<<8 | 'c' // 'mmtc'
)

func (e CAClockSyncMode) String() string {
	switch e {
	case KCAClockSyncMode_Internal:
		return "KCAClockSyncMode_Internal"
	case KCAClockSyncMode_MIDIClockTransport:
		return "KCAClockSyncMode_MIDIClockTransport"
	case KCAClockSyncMode_MTCTransport:
		return "KCAClockSyncMode_MTCTransport"
	default:
		return fmt.Sprintf("CAClockSyncMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat
type CAClockTimeFormat uint32

const (
	KCAClockTimeFormat_AbsoluteSeconds CAClockTimeFormat = 'a'<<24 | 's'<<16 | 'e'<<8 | 'c' // 'asec'
	KCAClockTimeFormat_Beats           CAClockTimeFormat = 'b'<<24 | 'e'<<16 | 'a'<<8 | 't' // 'beat'
	KCAClockTimeFormat_HostTime        CAClockTimeFormat = 'h'<<24 | 'o'<<16 | 's'<<8 | 't' // 'host'
	KCAClockTimeFormat_SMPTESeconds    CAClockTimeFormat = 's'<<24 | 'm'<<16 | 'p'<<8 | 's' // 'smps'
	KCAClockTimeFormat_SMPTETime       CAClockTimeFormat = 's'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'smpt'
	KCAClockTimeFormat_Samples         CAClockTimeFormat = 's'<<24 | 'a'<<16 | 'm'<<8 | 'p' // 'samp'
	KCAClockTimeFormat_Seconds         CAClockTimeFormat = 's'<<24 | 'e'<<16 | 'c'<<8 | 's' // 'secs'
)

func (e CAClockTimeFormat) String() string {
	switch e {
	case KCAClockTimeFormat_AbsoluteSeconds:
		return "KCAClockTimeFormat_AbsoluteSeconds"
	case KCAClockTimeFormat_Beats:
		return "KCAClockTimeFormat_Beats"
	case KCAClockTimeFormat_HostTime:
		return "KCAClockTimeFormat_HostTime"
	case KCAClockTimeFormat_SMPTESeconds:
		return "KCAClockTimeFormat_SMPTESeconds"
	case KCAClockTimeFormat_SMPTETime:
		return "KCAClockTimeFormat_SMPTETime"
	case KCAClockTimeFormat_Samples:
		return "KCAClockTimeFormat_Samples"
	case KCAClockTimeFormat_Seconds:
		return "KCAClockTimeFormat_Seconds"
	default:
		return fmt.Sprintf("CAClockTimeFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAClockTimebase
type CAClockTimebase uint32

const (
	KCAClockTimebase_AudioDevice     CAClockTimebase = 'a'<<24 | 'u'<<16 | 'd'<<8 | 'i' // 'audi'
	KCAClockTimebase_AudioOutputUnit CAClockTimebase = 'a'<<24 | 'u'<<16 | 'o'<<8 | 'u' // 'auou'
	KCAClockTimebase_HostTime        CAClockTimebase = 'h'<<24 | 'o'<<16 | 's'<<8 | 't' // 'host'
)

func (e CAClockTimebase) String() string {
	switch e {
	case KCAClockTimebase_AudioDevice:
		return "KCAClockTimebase_AudioDevice"
	case KCAClockTimebase_AudioOutputUnit:
		return "KCAClockTimebase_AudioOutputUnit"
	case KCAClockTimebase_HostTime:
		return "KCAClockTimebase_HostTime"
	default:
		return fmt.Sprintf("CAClockTimebase(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAFFormatFlags
type CAFFormatFlags uint32

const (
	KCAFLinearPCMFormatFlagIsFloat        CAFFormatFlags = 1
	KCAFLinearPCMFormatFlagIsLittleEndian CAFFormatFlags = 2
)

func (e CAFFormatFlags) String() string {
	switch e {
	case KCAFLinearPCMFormatFlagIsFloat:
		return "KCAFLinearPCMFormatFlagIsFloat"
	case KCAFLinearPCMFormatFlagIsLittleEndian:
		return "KCAFLinearPCMFormatFlagIsLittleEndian"
	default:
		return fmt.Sprintf("CAFFormatFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/CAFRegionFlags
type CAFRegionFlags uint32

const (
	KCAFRegionFlag_LoopEnable   CAFRegionFlags = 1
	KCAFRegionFlag_PlayBackward CAFRegionFlags = 4
	KCAFRegionFlag_PlayForward  CAFRegionFlags = 2
)

func (e CAFRegionFlags) String() string {
	switch e {
	case KCAFRegionFlag_LoopEnable:
		return "KCAFRegionFlag_LoopEnable"
	case KCAFRegionFlag_PlayBackward:
		return "KCAFRegionFlag_PlayBackward"
	case KCAFRegionFlag_PlayForward:
		return "KCAFRegionFlag_PlayForward"
	default:
		return fmt.Sprintf("CAFRegionFlags(%d)", e)
	}
}

type K3DMixerParam uint32

const (
	K3DMixerParam_Azimuth                          K3DMixerParam = 0
	K3DMixerParam_BusEnable                        K3DMixerParam = 20
	K3DMixerParam_Distance                         K3DMixerParam = 2
	K3DMixerParam_DryWetReverbBlend                K3DMixerParam = 23
	K3DMixerParam_Elevation                        K3DMixerParam = 1
	K3DMixerParam_Gain                             K3DMixerParam = 3
	K3DMixerParam_GlobalReverbGainInDecibels       K3DMixerParam = 24
	K3DMixerParam_MaxGainInDecibels                K3DMixerParam = 22
	K3DMixerParam_MinGainInDecibels                K3DMixerParam = 21
	K3DMixerParam_ObstructionAttenuationInDecibels K3DMixerParam = 26
	K3DMixerParam_OcclusionAttenuationInDecibels   K3DMixerParam = 25
	K3DMixerParam_PlaybackRate                     K3DMixerParam = 4
	K3DMixerParam_PostAveragePower                 K3DMixerParam = 3000
	K3DMixerParam_PostPeakHoldLevel                K3DMixerParam = 4000
	K3DMixerParam_PreAveragePower                  K3DMixerParam = 1000
	K3DMixerParam_PrePeakHoldLevel                 K3DMixerParam = 2000
	// Deprecated: use K3DMixerParam_GlobalReverbGainInDecibels.
	K3DMixerParam_GlobalReverbGain K3DMixerParam = 6
	// Deprecated: use K3DMixerParam_MaxGainInDecibels.
	K3DMixerParam_MaxGain K3DMixerParam = 10
	// Deprecated: use K3DMixerParam_MinGainInDecibels.
	K3DMixerParam_MinGain K3DMixerParam = 9
	// Deprecated: use K3DMixerParam_ObstructionAttenuationInDecibels.
	K3DMixerParam_ObstructionAttenuation K3DMixerParam = 8
	// Deprecated: use K3DMixerParam_OcclusionAttenuationInDecibels.
	K3DMixerParam_OcclusionAttenuation K3DMixerParam = 7
	// Deprecated: use K3DMixerParam_DryWetReverbBlend.
	K3DMixerParam_ReverbBlend K3DMixerParam = 5
)

func (e K3DMixerParam) String() string {
	switch e {
	case K3DMixerParam_Azimuth:
		return "K3DMixerParam_Azimuth"
	case K3DMixerParam_BusEnable:
		return "K3DMixerParam_BusEnable"
	case K3DMixerParam_Distance:
		return "K3DMixerParam_Distance"
	case K3DMixerParam_DryWetReverbBlend:
		return "K3DMixerParam_DryWetReverbBlend"
	case K3DMixerParam_Elevation:
		return "K3DMixerParam_Elevation"
	case K3DMixerParam_Gain:
		return "K3DMixerParam_Gain"
	case K3DMixerParam_GlobalReverbGainInDecibels:
		return "K3DMixerParam_GlobalReverbGainInDecibels"
	case K3DMixerParam_MaxGainInDecibels:
		return "K3DMixerParam_MaxGainInDecibels"
	case K3DMixerParam_MinGainInDecibels:
		return "K3DMixerParam_MinGainInDecibels"
	case K3DMixerParam_ObstructionAttenuationInDecibels:
		return "K3DMixerParam_ObstructionAttenuationInDecibels"
	case K3DMixerParam_OcclusionAttenuationInDecibels:
		return "K3DMixerParam_OcclusionAttenuationInDecibels"
	case K3DMixerParam_PlaybackRate:
		return "K3DMixerParam_PlaybackRate"
	case K3DMixerParam_PostAveragePower:
		return "K3DMixerParam_PostAveragePower"
	case K3DMixerParam_PostPeakHoldLevel:
		return "K3DMixerParam_PostPeakHoldLevel"
	case K3DMixerParam_PreAveragePower:
		return "K3DMixerParam_PreAveragePower"
	case K3DMixerParam_PrePeakHoldLevel:
		return "K3DMixerParam_PrePeakHoldLevel"
	case K3DMixerParam_GlobalReverbGain:
		return "K3DMixerParam_GlobalReverbGain"
	case K3DMixerParam_MaxGain:
		return "K3DMixerParam_MaxGain"
	case K3DMixerParam_MinGain:
		return "K3DMixerParam_MinGain"
	case K3DMixerParam_ObstructionAttenuation:
		return "K3DMixerParam_ObstructionAttenuation"
	case K3DMixerParam_OcclusionAttenuation:
		return "K3DMixerParam_OcclusionAttenuation"
	case K3DMixerParam_ReverbBlend:
		return "K3DMixerParam_ReverbBlend"
	default:
		return fmt.Sprintf("K3DMixerParam(%d)", e)
	}
}

type KAUAudioMixParameter uint32

const (
	KAUAudioMixParameter_RemixAmount KAUAudioMixParameter = 1
	KAUAudioMixParameter_Style       KAUAudioMixParameter = 0
)

func (e KAUAudioMixParameter) String() string {
	switch e {
	case KAUAudioMixParameter_RemixAmount:
		return "KAUAudioMixParameter_RemixAmount"
	case KAUAudioMixParameter_Style:
		return "KAUAudioMixParameter_Style"
	default:
		return fmt.Sprintf("KAUAudioMixParameter(%d)", e)
	}
}

type KAUAudioMixProperty uint32

const (
	KAUAudioMixProperty_EnableSpatialization    KAUAudioMixProperty = 5001
	KAUAudioMixProperty_SpatialAudioMixMetadata KAUAudioMixProperty = 5000
)

func (e KAUAudioMixProperty) String() string {
	switch e {
	case KAUAudioMixProperty_EnableSpatialization:
		return "KAUAudioMixProperty_EnableSpatialization"
	case KAUAudioMixProperty_SpatialAudioMixMetadata:
		return "KAUAudioMixProperty_SpatialAudioMixMetadata"
	default:
		return fmt.Sprintf("KAUAudioMixProperty(%d)", e)
	}
}

type KAUGraphErr int32

const (
	// KAUGraphErr_CannotDoInCurrentContext: To avoid spinning or waiting in the render thread (a bad idea!), many of the calls to AUGraph can return: `kAUGraphErr_CannotDoInCurrentContext`.
	KAUGraphErr_CannotDoInCurrentContext KAUGraphErr = -10863
	KAUGraphErr_InvalidAudioUnit         KAUGraphErr = -10864
	// KAUGraphErr_InvalidConnection: The attempted connection between two nodes cannot be made.
	KAUGraphErr_InvalidConnection KAUGraphErr = -10861
	// KAUGraphErr_NodeNotFound: The specified node cannot be found.
	KAUGraphErr_NodeNotFound KAUGraphErr = -10860
	// KAUGraphErr_OutputNodeErr: Audio processing graphs can only contain one output unit.
	KAUGraphErr_OutputNodeErr KAUGraphErr = -10862
)

func (e KAUGraphErr) String() string {
	switch e {
	case KAUGraphErr_CannotDoInCurrentContext:
		return "KAUGraphErr_CannotDoInCurrentContext"
	case KAUGraphErr_InvalidAudioUnit:
		return "KAUGraphErr_InvalidAudioUnit"
	case KAUGraphErr_InvalidConnection:
		return "KAUGraphErr_InvalidConnection"
	case KAUGraphErr_NodeNotFound:
		return "KAUGraphErr_NodeNotFound"
	case KAUGraphErr_OutputNodeErr:
		return "KAUGraphErr_OutputNodeErr"
	default:
		return fmt.Sprintf("KAUGraphErr(%d)", e)
	}
}

type KAUGroupParameterID uint32

const (
	KAUGroupParameterID_AllNotesOff          KAUGroupParameterID = 123
	KAUGroupParameterID_AllSoundOff          KAUGroupParameterID = 120
	KAUGroupParameterID_ChannelPressure      KAUGroupParameterID = 0xd0
	KAUGroupParameterID_DataEntry            KAUGroupParameterID = 6
	KAUGroupParameterID_DataEntry_LSB        KAUGroupParameterID = 38
	KAUGroupParameterID_Expression           KAUGroupParameterID = 11
	KAUGroupParameterID_Expression_LSB       KAUGroupParameterID = 43
	KAUGroupParameterID_Foot                 KAUGroupParameterID = 4
	KAUGroupParameterID_Foot_LSB             KAUGroupParameterID = 36
	KAUGroupParameterID_KeyPressure          KAUGroupParameterID = 0xa0
	KAUGroupParameterID_KeyPressure_FirstKey KAUGroupParameterID = 256
	KAUGroupParameterID_KeyPressure_LastKey  KAUGroupParameterID = 383
	KAUGroupParameterID_ModWheel             KAUGroupParameterID = 1
	KAUGroupParameterID_ModWheel_LSB         KAUGroupParameterID = 33
	KAUGroupParameterID_Pan                  KAUGroupParameterID = 10
	KAUGroupParameterID_Pan_LSB              KAUGroupParameterID = 42
	KAUGroupParameterID_PitchBend            KAUGroupParameterID = 0xe0
	KAUGroupParameterID_ResetAllControllers  KAUGroupParameterID = 121
	KAUGroupParameterID_Sostenuto            KAUGroupParameterID = 66
	KAUGroupParameterID_Sustain              KAUGroupParameterID = 64
	KAUGroupParameterID_Volume               KAUGroupParameterID = 7
	KAUGroupParameterID_Volume_LSB           KAUGroupParameterID = 39
)

func (e KAUGroupParameterID) String() string {
	switch e {
	case KAUGroupParameterID_AllNotesOff:
		return "KAUGroupParameterID_AllNotesOff"
	case KAUGroupParameterID_AllSoundOff:
		return "KAUGroupParameterID_AllSoundOff"
	case KAUGroupParameterID_ChannelPressure:
		return "KAUGroupParameterID_ChannelPressure"
	case KAUGroupParameterID_DataEntry:
		return "KAUGroupParameterID_DataEntry"
	case KAUGroupParameterID_DataEntry_LSB:
		return "KAUGroupParameterID_DataEntry_LSB"
	case KAUGroupParameterID_Expression:
		return "KAUGroupParameterID_Expression"
	case KAUGroupParameterID_Expression_LSB:
		return "KAUGroupParameterID_Expression_LSB"
	case KAUGroupParameterID_Foot:
		return "KAUGroupParameterID_Foot"
	case KAUGroupParameterID_Foot_LSB:
		return "KAUGroupParameterID_Foot_LSB"
	case KAUGroupParameterID_KeyPressure:
		return "KAUGroupParameterID_KeyPressure"
	case KAUGroupParameterID_KeyPressure_FirstKey:
		return "KAUGroupParameterID_KeyPressure_FirstKey"
	case KAUGroupParameterID_KeyPressure_LastKey:
		return "KAUGroupParameterID_KeyPressure_LastKey"
	case KAUGroupParameterID_ModWheel:
		return "KAUGroupParameterID_ModWheel"
	case KAUGroupParameterID_ModWheel_LSB:
		return "KAUGroupParameterID_ModWheel_LSB"
	case KAUGroupParameterID_Pan:
		return "KAUGroupParameterID_Pan"
	case KAUGroupParameterID_Pan_LSB:
		return "KAUGroupParameterID_Pan_LSB"
	case KAUGroupParameterID_PitchBend:
		return "KAUGroupParameterID_PitchBend"
	case KAUGroupParameterID_ResetAllControllers:
		return "KAUGroupParameterID_ResetAllControllers"
	case KAUGroupParameterID_Sostenuto:
		return "KAUGroupParameterID_Sostenuto"
	case KAUGroupParameterID_Sustain:
		return "KAUGroupParameterID_Sustain"
	case KAUGroupParameterID_Volume:
		return "KAUGroupParameterID_Volume"
	case KAUGroupParameterID_Volume_LSB:
		return "KAUGroupParameterID_Volume_LSB"
	default:
		return fmt.Sprintf("KAUGroupParameterID(%d)", e)
	}
}

type KAULowShelfParam uint32

const (
	KAULowShelfParam_CutoffFrequency KAULowShelfParam = 0
	KAULowShelfParam_Gain            KAULowShelfParam = 1
)

func (e KAULowShelfParam) String() string {
	switch e {
	case KAULowShelfParam_CutoffFrequency:
		return "KAULowShelfParam_CutoffFrequency"
	case KAULowShelfParam_Gain:
		return "KAULowShelfParam_Gain"
	default:
		return fmt.Sprintf("KAULowShelfParam(%d)", e)
	}
}

type KAUNBandEQParam uint32

const (
	KAUNBandEQParam_Bandwidth  KAUNBandEQParam = 5000
	KAUNBandEQParam_BypassBand KAUNBandEQParam = 1000
	KAUNBandEQParam_FilterType KAUNBandEQParam = 2000
	KAUNBandEQParam_Frequency  KAUNBandEQParam = 3000
	KAUNBandEQParam_Gain       KAUNBandEQParam = 4000
	KAUNBandEQParam_GlobalGain KAUNBandEQParam = 0
)

func (e KAUNBandEQParam) String() string {
	switch e {
	case KAUNBandEQParam_Bandwidth:
		return "KAUNBandEQParam_Bandwidth"
	case KAUNBandEQParam_BypassBand:
		return "KAUNBandEQParam_BypassBand"
	case KAUNBandEQParam_FilterType:
		return "KAUNBandEQParam_FilterType"
	case KAUNBandEQParam_Frequency:
		return "KAUNBandEQParam_Frequency"
	case KAUNBandEQParam_Gain:
		return "KAUNBandEQParam_Gain"
	case KAUNBandEQParam_GlobalGain:
		return "KAUNBandEQParam_GlobalGain"
	default:
		return fmt.Sprintf("KAUNBandEQParam(%d)", e)
	}
}

type KAUNBandEQProperty uint32

const (
	KAUNBandEQProperty_BiquadCoefficients KAUNBandEQProperty = 2203
	KAUNBandEQProperty_MaxNumberOfBands   KAUNBandEQProperty = 2201
	KAUNBandEQProperty_NumberOfBands      KAUNBandEQProperty = 2200
)

func (e KAUNBandEQProperty) String() string {
	switch e {
	case KAUNBandEQProperty_BiquadCoefficients:
		return "KAUNBandEQProperty_BiquadCoefficients"
	case KAUNBandEQProperty_MaxNumberOfBands:
		return "KAUNBandEQProperty_MaxNumberOfBands"
	case KAUNBandEQProperty_NumberOfBands:
		return "KAUNBandEQProperty_NumberOfBands"
	default:
		return fmt.Sprintf("KAUNBandEQProperty(%d)", e)
	}
}

type KAUNetReceiveParam uint32

const (
	KAUNetReceiveParam_NumParameters KAUNetReceiveParam = 1
	KAUNetReceiveParam_Status        KAUNetReceiveParam = 0
)

func (e KAUNetReceiveParam) String() string {
	switch e {
	case KAUNetReceiveParam_NumParameters:
		return "KAUNetReceiveParam_NumParameters"
	case KAUNetReceiveParam_Status:
		return "KAUNetReceiveParam_Status"
	default:
		return fmt.Sprintf("KAUNetReceiveParam(%d)", e)
	}
}

type KAUNetReceiveProperty uint32

const (
	KAUNetReceiveProperty_Hostname KAUNetReceiveProperty = 3511
	KAUNetReceiveProperty_Password KAUNetReceiveProperty = 3512
)

func (e KAUNetReceiveProperty) String() string {
	switch e {
	case KAUNetReceiveProperty_Hostname:
		return "KAUNetReceiveProperty_Hostname"
	case KAUNetReceiveProperty_Password:
		return "KAUNetReceiveProperty_Password"
	default:
		return fmt.Sprintf("KAUNetReceiveProperty(%d)", e)
	}
}

type KAUNetSend uint32

const (
	KAUNetSendNumPresetFormats             KAUNetSend = 18
	KAUNetSendPresetFormat_AAC_128kbpspc   KAUNetSend = 7
	KAUNetSendPresetFormat_AAC_32kbpspc    KAUNetSend = 13
	KAUNetSendPresetFormat_AAC_40kbpspc    KAUNetSend = 12
	KAUNetSendPresetFormat_AAC_48kbpspc    KAUNetSend = 11
	KAUNetSendPresetFormat_AAC_64kbpspc    KAUNetSend = 10
	KAUNetSendPresetFormat_AAC_80kbpspc    KAUNetSend = 9
	KAUNetSendPresetFormat_AAC_96kbpspc    KAUNetSend = 8
	KAUNetSendPresetFormat_AAC_LD_32kbpspc KAUNetSend = 17
	KAUNetSendPresetFormat_AAC_LD_40kbpspc KAUNetSend = 16
	KAUNetSendPresetFormat_AAC_LD_48kbpspc KAUNetSend = 15
	KAUNetSendPresetFormat_AAC_LD_64kbpspc KAUNetSend = 14
	KAUNetSendPresetFormat_IMA4            KAUNetSend = 6
	KAUNetSendPresetFormat_Lossless16      KAUNetSend = 4
	KAUNetSendPresetFormat_Lossless24      KAUNetSend = 3
	KAUNetSendPresetFormat_PCMFloat32      KAUNetSend = 0
	KAUNetSendPresetFormat_PCMInt16        KAUNetSend = 2
	KAUNetSendPresetFormat_PCMInt24        KAUNetSend = 1
	KAUNetSendPresetFormat_ULaw            KAUNetSend = 5
)

func (e KAUNetSend) String() string {
	switch e {
	case KAUNetSendNumPresetFormats:
		return "KAUNetSendNumPresetFormats"
	case KAUNetSendPresetFormat_AAC_128kbpspc:
		return "KAUNetSendPresetFormat_AAC_128kbpspc"
	case KAUNetSendPresetFormat_AAC_32kbpspc:
		return "KAUNetSendPresetFormat_AAC_32kbpspc"
	case KAUNetSendPresetFormat_AAC_40kbpspc:
		return "KAUNetSendPresetFormat_AAC_40kbpspc"
	case KAUNetSendPresetFormat_AAC_48kbpspc:
		return "KAUNetSendPresetFormat_AAC_48kbpspc"
	case KAUNetSendPresetFormat_AAC_64kbpspc:
		return "KAUNetSendPresetFormat_AAC_64kbpspc"
	case KAUNetSendPresetFormat_AAC_80kbpspc:
		return "KAUNetSendPresetFormat_AAC_80kbpspc"
	case KAUNetSendPresetFormat_AAC_96kbpspc:
		return "KAUNetSendPresetFormat_AAC_96kbpspc"
	case KAUNetSendPresetFormat_AAC_LD_32kbpspc:
		return "KAUNetSendPresetFormat_AAC_LD_32kbpspc"
	case KAUNetSendPresetFormat_AAC_LD_40kbpspc:
		return "KAUNetSendPresetFormat_AAC_LD_40kbpspc"
	case KAUNetSendPresetFormat_AAC_LD_48kbpspc:
		return "KAUNetSendPresetFormat_AAC_LD_48kbpspc"
	case KAUNetSendPresetFormat_AAC_LD_64kbpspc:
		return "KAUNetSendPresetFormat_AAC_LD_64kbpspc"
	case KAUNetSendPresetFormat_IMA4:
		return "KAUNetSendPresetFormat_IMA4"
	case KAUNetSendPresetFormat_Lossless16:
		return "KAUNetSendPresetFormat_Lossless16"
	case KAUNetSendPresetFormat_Lossless24:
		return "KAUNetSendPresetFormat_Lossless24"
	case KAUNetSendPresetFormat_PCMFloat32:
		return "KAUNetSendPresetFormat_PCMFloat32"
	case KAUNetSendPresetFormat_PCMInt16:
		return "KAUNetSendPresetFormat_PCMInt16"
	case KAUNetSendPresetFormat_PCMInt24:
		return "KAUNetSendPresetFormat_PCMInt24"
	case KAUNetSendPresetFormat_ULaw:
		return "KAUNetSendPresetFormat_ULaw"
	default:
		return fmt.Sprintf("KAUNetSend(%d)", e)
	}
}

type KAUNetSendParam uint32

const (
	KAUNetSendParam_NumParameters KAUNetSendParam = 1
	KAUNetSendParam_Status        KAUNetSendParam = 0
)

func (e KAUNetSendParam) String() string {
	switch e {
	case KAUNetSendParam_NumParameters:
		return "KAUNetSendParam_NumParameters"
	case KAUNetSendParam_Status:
		return "KAUNetSendParam_Status"
	default:
		return fmt.Sprintf("KAUNetSendParam(%d)", e)
	}
}

type KAUNetSendProperty uint32

const (
	KAUNetSendProperty_Disconnect              KAUNetSendProperty = 3517
	KAUNetSendProperty_Password                KAUNetSendProperty = 3518
	KAUNetSendProperty_PortNum                 KAUNetSendProperty = 3513
	KAUNetSendProperty_ServiceName             KAUNetSendProperty = 3516
	KAUNetSendProperty_TransmissionFormat      KAUNetSendProperty = 3514
	KAUNetSendProperty_TransmissionFormatIndex KAUNetSendProperty = 3515
)

func (e KAUNetSendProperty) String() string {
	switch e {
	case KAUNetSendProperty_Disconnect:
		return "KAUNetSendProperty_Disconnect"
	case KAUNetSendProperty_Password:
		return "KAUNetSendProperty_Password"
	case KAUNetSendProperty_PortNum:
		return "KAUNetSendProperty_PortNum"
	case KAUNetSendProperty_ServiceName:
		return "KAUNetSendProperty_ServiceName"
	case KAUNetSendProperty_TransmissionFormat:
		return "KAUNetSendProperty_TransmissionFormat"
	case KAUNetSendProperty_TransmissionFormatIndex:
		return "KAUNetSendProperty_TransmissionFormatIndex"
	default:
		return fmt.Sprintf("KAUNetSendProperty(%d)", e)
	}
}

type KAUNetStatus uint32

const (
	KAUNetStatus_Connected    KAUNetStatus = 1
	KAUNetStatus_Connecting   KAUNetStatus = 4
	KAUNetStatus_Listening    KAUNetStatus = 5
	KAUNetStatus_NotConnected KAUNetStatus = 0
	KAUNetStatus_Overflow     KAUNetStatus = 2
	KAUNetStatus_Underflow    KAUNetStatus = 3
)

func (e KAUNetStatus) String() string {
	switch e {
	case KAUNetStatus_Connected:
		return "KAUNetStatus_Connected"
	case KAUNetStatus_Connecting:
		return "KAUNetStatus_Connecting"
	case KAUNetStatus_Listening:
		return "KAUNetStatus_Listening"
	case KAUNetStatus_NotConnected:
		return "KAUNetStatus_NotConnected"
	case KAUNetStatus_Overflow:
		return "KAUNetStatus_Overflow"
	case KAUNetStatus_Underflow:
		return "KAUNetStatus_Underflow"
	default:
		return fmt.Sprintf("KAUNetStatus(%d)", e)
	}
}

type KAUNodeInteraction uint32

const (
	// KAUNodeInteraction_Connection: connections between 2 audio units,
	KAUNodeInteraction_Connection KAUNodeInteraction = 1
	// KAUNodeInteraction_InputCallback: input callbacks being registered to a single audio unit’s input bus.
	KAUNodeInteraction_InputCallback KAUNodeInteraction = 2
)

func (e KAUNodeInteraction) String() string {
	switch e {
	case KAUNodeInteraction_Connection:
		return "KAUNodeInteraction_Connection"
	case KAUNodeInteraction_InputCallback:
		return "KAUNodeInteraction_InputCallback"
	default:
		return fmt.Sprintf("KAUNodeInteraction(%d)", e)
	}
}

type KAUSamplerParam uint32

const (
	KAUSamplerParam_CoarseTuning KAUSamplerParam = 901
	KAUSamplerParam_FineTuning   KAUSamplerParam = 902
	KAUSamplerParam_Gain         KAUSamplerParam = 900
	KAUSamplerParam_Pan          KAUSamplerParam = 903
)

func (e KAUSamplerParam) String() string {
	switch e {
	case KAUSamplerParam_CoarseTuning:
		return "KAUSamplerParam_CoarseTuning"
	case KAUSamplerParam_FineTuning:
		return "KAUSamplerParam_FineTuning"
	case KAUSamplerParam_Gain:
		return "KAUSamplerParam_Gain"
	case KAUSamplerParam_Pan:
		return "KAUSamplerParam_Pan"
	default:
		return fmt.Sprintf("KAUSamplerParam(%d)", e)
	}
}

type KAUSamplerProperty uint32

const (
	KAUSamplerProperty_BankAndPreset      KAUSamplerProperty = 4100
	KAUSamplerProperty_LoadPresetFromBank KAUSamplerProperty = 4100
)

func (e KAUSamplerProperty) String() string {
	switch e {
	case KAUSamplerProperty_BankAndPreset:
		return "KAUSamplerProperty_BankAndPreset"
	default:
		return fmt.Sprintf("KAUSamplerProperty(%d)", e)
	}
}

type KAUSoundIsolationParam uint32

const (
	KAUSoundIsolationParam_SoundToIsolate   KAUSoundIsolationParam = 1
	KAUSoundIsolationParam_WetDryMixPercent KAUSoundIsolationParam = 0
)

func (e KAUSoundIsolationParam) String() string {
	switch e {
	case KAUSoundIsolationParam_SoundToIsolate:
		return "KAUSoundIsolationParam_SoundToIsolate"
	case KAUSoundIsolationParam_WetDryMixPercent:
		return "KAUSoundIsolationParam_WetDryMixPercent"
	default:
		return fmt.Sprintf("KAUSoundIsolationParam(%d)", e)
	}
}

type KAUSoundIsolationSoundType int

const (
	KAUSoundIsolationSoundType_HighQualityVoice KAUSoundIsolationSoundType = 0
	KAUSoundIsolationSoundType_Voice            KAUSoundIsolationSoundType = 1
)

func (e KAUSoundIsolationSoundType) String() string {
	switch e {
	case KAUSoundIsolationSoundType_HighQualityVoice:
		return "KAUSoundIsolationSoundType_HighQualityVoice"
	case KAUSoundIsolationSoundType_Voice:
		return "KAUSoundIsolationSoundType_Voice"
	default:
		return fmt.Sprintf("KAUSoundIsolationSoundType(%d)", e)
	}
}

type KAUVoiceIOProperty uint32

const (
	// KAUVoiceIOProperty_BypassVoiceProcessing: A property that bypasses all processing for microphone uplink done by the voice processing unit.
	KAUVoiceIOProperty_BypassVoiceProcessing KAUVoiceIOProperty = 2100
	// KAUVoiceIOProperty_MuteOutput: A property to mute the output of the processed microphone uplink.
	KAUVoiceIOProperty_MuteOutput KAUVoiceIOProperty = 2104
	// KAUVoiceIOProperty_VoiceProcessingEnableAGC: A property to enable automatic gain control on the processed microphone uplink.
	KAUVoiceIOProperty_VoiceProcessingEnableAGC KAUVoiceIOProperty = 2101
)

func (e KAUVoiceIOProperty) String() string {
	switch e {
	case KAUVoiceIOProperty_BypassVoiceProcessing:
		return "KAUVoiceIOProperty_BypassVoiceProcessing"
	case KAUVoiceIOProperty_MuteOutput:
		return "KAUVoiceIOProperty_MuteOutput"
	case KAUVoiceIOProperty_VoiceProcessingEnableAGC:
		return "KAUVoiceIOProperty_VoiceProcessingEnableAGC"
	default:
		return fmt.Sprintf("KAUVoiceIOProperty(%d)", e)
	}
}

type KAudioCodecBitRateControlMode uint32

const (
	KAudioCodecBitRateControlMode_Constant            KAudioCodecBitRateControlMode = 0
	KAudioCodecBitRateControlMode_LongTermAverage     KAudioCodecBitRateControlMode = 1
	KAudioCodecBitRateControlMode_Variable            KAudioCodecBitRateControlMode = 3
	KAudioCodecBitRateControlMode_VariableConstrained KAudioCodecBitRateControlMode = 2
)

func (e KAudioCodecBitRateControlMode) String() string {
	switch e {
	case KAudioCodecBitRateControlMode_Constant:
		return "KAudioCodecBitRateControlMode_Constant"
	case KAudioCodecBitRateControlMode_LongTermAverage:
		return "KAudioCodecBitRateControlMode_LongTermAverage"
	case KAudioCodecBitRateControlMode_Variable:
		return "KAudioCodecBitRateControlMode_Variable"
	case KAudioCodecBitRateControlMode_VariableConstrained:
		return "KAudioCodecBitRateControlMode_VariableConstrained"
	default:
		return fmt.Sprintf("KAudioCodecBitRateControlMode(%d)", e)
	}
}

type KAudioCodecBitRateFormat uint32

const (
	KAudioCodecBitRateFormat_ABR KAudioCodecBitRateFormat = 1
	KAudioCodecBitRateFormat_CBR KAudioCodecBitRateFormat = 0
	KAudioCodecBitRateFormat_VBR KAudioCodecBitRateFormat = 2
)

func (e KAudioCodecBitRateFormat) String() string {
	switch e {
	case KAudioCodecBitRateFormat_ABR:
		return "KAudioCodecBitRateFormat_ABR"
	case KAudioCodecBitRateFormat_CBR:
		return "KAudioCodecBitRateFormat_CBR"
	case KAudioCodecBitRateFormat_VBR:
		return "KAudioCodecBitRateFormat_VBR"
	default:
		return fmt.Sprintf("KAudioCodecBitRateFormat(%d)", e)
	}
}

type KAudioCodecContentSource int32

const (
	KAudioCodecContentSource_AV_Spatial_Live               KAudioCodecContentSource = 41
	KAudioCodecContentSource_AV_Spatial_Offline            KAudioCodecContentSource = 39
	KAudioCodecContentSource_AV_Traditional_Live           KAudioCodecContentSource = 40
	KAudioCodecContentSource_AV_Traditional_Offline        KAudioCodecContentSource = 38
	KAudioCodecContentSource_AppleAV_Spatial_Live          KAudioCodecContentSource = 9
	KAudioCodecContentSource_AppleAV_Spatial_Offline       KAudioCodecContentSource = 7
	KAudioCodecContentSource_AppleAV_Traditional_Live      KAudioCodecContentSource = 8
	KAudioCodecContentSource_AppleAV_Traditional_Offline   KAudioCodecContentSource = 6
	KAudioCodecContentSource_AppleCapture_Spatial          KAudioCodecContentSource = 2
	KAudioCodecContentSource_AppleCapture_Spatial_Enhanced KAudioCodecContentSource = 3
	KAudioCodecContentSource_AppleCapture_Traditional      KAudioCodecContentSource = 1
	KAudioCodecContentSource_AppleMusic_Spatial            KAudioCodecContentSource = 5
	KAudioCodecContentSource_AppleMusic_Traditional        KAudioCodecContentSource = 4
	KAudioCodecContentSource_ApplePassthrough              KAudioCodecContentSource = 10
	KAudioCodecContentSource_Capture_Spatial               KAudioCodecContentSource = 34
	KAudioCodecContentSource_Capture_Spatial_Enhanced      KAudioCodecContentSource = 35
	KAudioCodecContentSource_Capture_Traditional           KAudioCodecContentSource = 33
	KAudioCodecContentSource_Music_Spatial                 KAudioCodecContentSource = 37
	KAudioCodecContentSource_Music_Traditional             KAudioCodecContentSource = 36
	KAudioCodecContentSource_Passthrough                   KAudioCodecContentSource = 42
	KAudioCodecContentSource_Reserved                      KAudioCodecContentSource = 0
	KAudioCodecContentSource_Unspecified                   KAudioCodecContentSource = -1
)

func (e KAudioCodecContentSource) String() string {
	switch e {
	case KAudioCodecContentSource_AV_Spatial_Live:
		return "KAudioCodecContentSource_AV_Spatial_Live"
	case KAudioCodecContentSource_AV_Spatial_Offline:
		return "KAudioCodecContentSource_AV_Spatial_Offline"
	case KAudioCodecContentSource_AV_Traditional_Live:
		return "KAudioCodecContentSource_AV_Traditional_Live"
	case KAudioCodecContentSource_AV_Traditional_Offline:
		return "KAudioCodecContentSource_AV_Traditional_Offline"
	case KAudioCodecContentSource_AppleAV_Spatial_Live:
		return "KAudioCodecContentSource_AppleAV_Spatial_Live"
	case KAudioCodecContentSource_AppleAV_Spatial_Offline:
		return "KAudioCodecContentSource_AppleAV_Spatial_Offline"
	case KAudioCodecContentSource_AppleAV_Traditional_Live:
		return "KAudioCodecContentSource_AppleAV_Traditional_Live"
	case KAudioCodecContentSource_AppleAV_Traditional_Offline:
		return "KAudioCodecContentSource_AppleAV_Traditional_Offline"
	case KAudioCodecContentSource_AppleCapture_Spatial:
		return "KAudioCodecContentSource_AppleCapture_Spatial"
	case KAudioCodecContentSource_AppleCapture_Spatial_Enhanced:
		return "KAudioCodecContentSource_AppleCapture_Spatial_Enhanced"
	case KAudioCodecContentSource_AppleCapture_Traditional:
		return "KAudioCodecContentSource_AppleCapture_Traditional"
	case KAudioCodecContentSource_AppleMusic_Spatial:
		return "KAudioCodecContentSource_AppleMusic_Spatial"
	case KAudioCodecContentSource_AppleMusic_Traditional:
		return "KAudioCodecContentSource_AppleMusic_Traditional"
	case KAudioCodecContentSource_ApplePassthrough:
		return "KAudioCodecContentSource_ApplePassthrough"
	case KAudioCodecContentSource_Capture_Spatial:
		return "KAudioCodecContentSource_Capture_Spatial"
	case KAudioCodecContentSource_Capture_Spatial_Enhanced:
		return "KAudioCodecContentSource_Capture_Spatial_Enhanced"
	case KAudioCodecContentSource_Capture_Traditional:
		return "KAudioCodecContentSource_Capture_Traditional"
	case KAudioCodecContentSource_Music_Spatial:
		return "KAudioCodecContentSource_Music_Spatial"
	case KAudioCodecContentSource_Music_Traditional:
		return "KAudioCodecContentSource_Music_Traditional"
	case KAudioCodecContentSource_Passthrough:
		return "KAudioCodecContentSource_Passthrough"
	case KAudioCodecContentSource_Reserved:
		return "KAudioCodecContentSource_Reserved"
	case KAudioCodecContentSource_Unspecified:
		return "KAudioCodecContentSource_Unspecified"
	default:
		return fmt.Sprintf("KAudioCodecContentSource(%d)", e)
	}
}

type KAudioCodecDelayMode uint32

const (
	KAudioCodecDelayMode_Compatibility KAudioCodecDelayMode = 0
	KAudioCodecDelayMode_Minimum       KAudioCodecDelayMode = 1
	KAudioCodecDelayMode_Optimal       KAudioCodecDelayMode = 2
)

func (e KAudioCodecDelayMode) String() string {
	switch e {
	case KAudioCodecDelayMode_Compatibility:
		return "KAudioCodecDelayMode_Compatibility"
	case KAudioCodecDelayMode_Minimum:
		return "KAudioCodecDelayMode_Minimum"
	case KAudioCodecDelayMode_Optimal:
		return "KAudioCodecDelayMode_Optimal"
	default:
		return fmt.Sprintf("KAudioCodecDelayMode(%d)", e)
	}
}

type KAudioCodecDynamicRangeControlConfiguration uint32

const (
	KAudioCodecDynamicRangeControlConfiguration_Capture KAudioCodecDynamicRangeControlConfiguration = 4
	KAudioCodecDynamicRangeControlConfiguration_Movie   KAudioCodecDynamicRangeControlConfiguration = 3
	KAudioCodecDynamicRangeControlConfiguration_Music   KAudioCodecDynamicRangeControlConfiguration = 1
	KAudioCodecDynamicRangeControlConfiguration_None    KAudioCodecDynamicRangeControlConfiguration = 0
	KAudioCodecDynamicRangeControlConfiguration_Speech  KAudioCodecDynamicRangeControlConfiguration = 2
)

func (e KAudioCodecDynamicRangeControlConfiguration) String() string {
	switch e {
	case KAudioCodecDynamicRangeControlConfiguration_Capture:
		return "KAudioCodecDynamicRangeControlConfiguration_Capture"
	case KAudioCodecDynamicRangeControlConfiguration_Movie:
		return "KAudioCodecDynamicRangeControlConfiguration_Movie"
	case KAudioCodecDynamicRangeControlConfiguration_Music:
		return "KAudioCodecDynamicRangeControlConfiguration_Music"
	case KAudioCodecDynamicRangeControlConfiguration_None:
		return "KAudioCodecDynamicRangeControlConfiguration_None"
	case KAudioCodecDynamicRangeControlConfiguration_Speech:
		return "KAudioCodecDynamicRangeControlConfiguration_Speech"
	default:
		return fmt.Sprintf("KAudioCodecDynamicRangeControlConfiguration(%d)", e)
	}
}

type KAudioCodecGetPropertyInfoSelect uint32

const (
	KAudioCodecAppendInputBufferListSelect   KAudioCodecGetPropertyInfoSelect = 0x9
	KAudioCodecAppendInputDataSelect         KAudioCodecGetPropertyInfoSelect = 0x6
	KAudioCodecGetPropertyInfoSelectValue    KAudioCodecGetPropertyInfoSelect = 0x1
	KAudioCodecGetPropertySelect             KAudioCodecGetPropertyInfoSelect = 0x2
	KAudioCodecInitializeSelect              KAudioCodecGetPropertyInfoSelect = 0x4
	KAudioCodecProduceOutputBufferListSelect KAudioCodecGetPropertyInfoSelect = 0xa
	KAudioCodecProduceOutputDataSelect       KAudioCodecGetPropertyInfoSelect = 0x7
	KAudioCodecResetSelect                   KAudioCodecGetPropertyInfoSelect = 0x8
	KAudioCodecSetPropertySelect             KAudioCodecGetPropertyInfoSelect = 0x3
	KAudioCodecUninitializeSelect            KAudioCodecGetPropertyInfoSelect = 0x5
)

func (e KAudioCodecGetPropertyInfoSelect) String() string {
	switch e {
	case KAudioCodecAppendInputBufferListSelect:
		return "KAudioCodecAppendInputBufferListSelect"
	case KAudioCodecAppendInputDataSelect:
		return "KAudioCodecAppendInputDataSelect"
	case KAudioCodecGetPropertyInfoSelectValue:
		return "KAudioCodecGetPropertyInfoSelectValue"
	case KAudioCodecGetPropertySelect:
		return "KAudioCodecGetPropertySelect"
	case KAudioCodecInitializeSelect:
		return "KAudioCodecInitializeSelect"
	case KAudioCodecProduceOutputBufferListSelect:
		return "KAudioCodecProduceOutputBufferListSelect"
	case KAudioCodecProduceOutputDataSelect:
		return "KAudioCodecProduceOutputDataSelect"
	case KAudioCodecResetSelect:
		return "KAudioCodecResetSelect"
	case KAudioCodecSetPropertySelect:
		return "KAudioCodecSetPropertySelect"
	case KAudioCodecUninitializeSelect:
		return "KAudioCodecUninitializeSelect"
	default:
		return fmt.Sprintf("KAudioCodecGetPropertyInfoSelect(%d)", e)
	}
}

type KAudioCodecNoError int32

const (
	KAudioCodecBadDataError              KAudioCodecNoError = 'b'<<24 | 'a'<<16 | 'd'<<8 | 'a' // 'bada'
	KAudioCodecBadPropertySizeError      KAudioCodecNoError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KAudioCodecIllegalOperationError     KAudioCodecNoError = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	KAudioCodecNoErrorValue              KAudioCodecNoError = 0
	KAudioCodecNotEnoughBufferSpaceError KAudioCodecNoError = '!'<<24 | 'b'<<16 | 'u'<<8 | 'f' // '!buf'
	KAudioCodecStateError                KAudioCodecNoError = '!'<<24 | 's'<<16 | 't'<<8 | 't' // '!stt'
	KAudioCodecUnknownPropertyError      KAudioCodecNoError = 'w'<<24 | 'h'<<16 | 'o'<<8 | '?' // 'who?'
	KAudioCodecUnspecifiedError          KAudioCodecNoError = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	KAudioCodecUnsupportedFormatError    KAudioCodecNoError = '!'<<24 | 'd'<<16 | 'a'<<8 | 't' // '!dat'
)

func (e KAudioCodecNoError) String() string {
	switch e {
	case KAudioCodecBadDataError:
		return "KAudioCodecBadDataError"
	case KAudioCodecBadPropertySizeError:
		return "KAudioCodecBadPropertySizeError"
	case KAudioCodecIllegalOperationError:
		return "KAudioCodecIllegalOperationError"
	case KAudioCodecNoErrorValue:
		return "KAudioCodecNoErrorValue"
	case KAudioCodecNotEnoughBufferSpaceError:
		return "KAudioCodecNotEnoughBufferSpaceError"
	case KAudioCodecStateError:
		return "KAudioCodecStateError"
	case KAudioCodecUnknownPropertyError:
		return "KAudioCodecUnknownPropertyError"
	case KAudioCodecUnspecifiedError:
		return "KAudioCodecUnspecifiedError"
	case KAudioCodecUnsupportedFormatError:
		return "KAudioCodecUnsupportedFormatError"
	default:
		return fmt.Sprintf("KAudioCodecNoError(%d)", e)
	}
}

type KAudioCodecOutputPrecedence uint32

const (
	KAudioCodecOutputPrecedenceBitRate    KAudioCodecOutputPrecedence = 1
	KAudioCodecOutputPrecedenceNone       KAudioCodecOutputPrecedence = 0
	KAudioCodecOutputPrecedenceSampleRate KAudioCodecOutputPrecedence = 2
)

func (e KAudioCodecOutputPrecedence) String() string {
	switch e {
	case KAudioCodecOutputPrecedenceBitRate:
		return "KAudioCodecOutputPrecedenceBitRate"
	case KAudioCodecOutputPrecedenceNone:
		return "KAudioCodecOutputPrecedenceNone"
	case KAudioCodecOutputPrecedenceSampleRate:
		return "KAudioCodecOutputPrecedenceSampleRate"
	default:
		return fmt.Sprintf("KAudioCodecOutputPrecedence(%d)", e)
	}
}

type KAudioCodecPrimeMethod uint32

const (
	KAudioCodecPrimeMethod_None   KAudioCodecPrimeMethod = 2
	KAudioCodecPrimeMethod_Normal KAudioCodecPrimeMethod = 1
	KAudioCodecPrimeMethod_Pre    KAudioCodecPrimeMethod = 0
)

func (e KAudioCodecPrimeMethod) String() string {
	switch e {
	case KAudioCodecPrimeMethod_None:
		return "KAudioCodecPrimeMethod_None"
	case KAudioCodecPrimeMethod_Normal:
		return "KAudioCodecPrimeMethod_Normal"
	case KAudioCodecPrimeMethod_Pre:
		return "KAudioCodecPrimeMethod_Pre"
	default:
		return fmt.Sprintf("KAudioCodecPrimeMethod(%d)", e)
	}
}

type KAudioCodecProduceOutputPacket uint32

const (
	KAudioCodecProduceOutputPacketAtEOF              KAudioCodecProduceOutputPacket = 5
	KAudioCodecProduceOutputPacketFailure            KAudioCodecProduceOutputPacket = 1
	KAudioCodecProduceOutputPacketNeedsMoreInputData KAudioCodecProduceOutputPacket = 4
	KAudioCodecProduceOutputPacketSuccess            KAudioCodecProduceOutputPacket = 2
	KAudioCodecProduceOutputPacketSuccessConcealed   KAudioCodecProduceOutputPacket = 6
	KAudioCodecProduceOutputPacketSuccessHasMore     KAudioCodecProduceOutputPacket = 3
)

func (e KAudioCodecProduceOutputPacket) String() string {
	switch e {
	case KAudioCodecProduceOutputPacketAtEOF:
		return "KAudioCodecProduceOutputPacketAtEOF"
	case KAudioCodecProduceOutputPacketFailure:
		return "KAudioCodecProduceOutputPacketFailure"
	case KAudioCodecProduceOutputPacketNeedsMoreInputData:
		return "KAudioCodecProduceOutputPacketNeedsMoreInputData"
	case KAudioCodecProduceOutputPacketSuccess:
		return "KAudioCodecProduceOutputPacketSuccess"
	case KAudioCodecProduceOutputPacketSuccessConcealed:
		return "KAudioCodecProduceOutputPacketSuccessConcealed"
	case KAudioCodecProduceOutputPacketSuccessHasMore:
		return "KAudioCodecProduceOutputPacketSuccessHasMore"
	default:
		return fmt.Sprintf("KAudioCodecProduceOutputPacket(%d)", e)
	}
}

type KAudioCodecPropertyInputBufferSize uint32

const (
	KAudioCodecPropertyASPFrequency                     KAudioCodecPropertyInputBufferSize = 'a'<<24 | 's'<<16 | 'p'<<8 | 'f' // 'aspf'
	KAudioCodecPropertyAdjustCompressionProfile         KAudioCodecPropertyInputBufferSize = '^'<<24 | 'p'<<16 | 'r'<<8 | 'o' // '^pro'
	KAudioCodecPropertyAdjustLocalQuality               KAudioCodecPropertyInputBufferSize = '^'<<24 | 'q'<<16 | 'a'<<8 | 'l' // '^qal'
	KAudioCodecPropertyAdjustTargetLevel                KAudioCodecPropertyInputBufferSize = '^'<<24 | 'p'<<16 | 't'<<8 | 'l' // '^ptl'
	KAudioCodecPropertyAdjustTargetLevelConstant        KAudioCodecPropertyInputBufferSize = '^'<<24 | 't'<<16 | 'l'<<8 | 'c' // '^tlc'
	KAudioCodecPropertyApplicableBitRateRange           KAudioCodecPropertyInputBufferSize = 'b'<<24 | 'r'<<16 | 't'<<8 | 'a' // 'brta'
	KAudioCodecPropertyApplicableInputSampleRates       KAudioCodecPropertyInputBufferSize = 'i'<<24 | 's'<<16 | 'r'<<8 | 'a' // 'isra'
	KAudioCodecPropertyApplicableOutputSampleRates      KAudioCodecPropertyInputBufferSize = 'o'<<24 | 's'<<16 | 'r'<<8 | 'a' // 'osra'
	KAudioCodecPropertyBitRateControlMode               KAudioCodecPropertyInputBufferSize = 'a'<<24 | 'c'<<16 | 'b'<<8 | 'f' // 'acbf'
	KAudioCodecPropertyBitRateForVBR                    KAudioCodecPropertyInputBufferSize = 'v'<<24 | 'b'<<16 | 'r'<<8 | 'b' // 'vbrb'
	KAudioCodecPropertyContentSource                    KAudioCodecPropertyInputBufferSize = 'c'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'csrc'
	KAudioCodecPropertyCurrentInputChannelLayout        KAudioCodecPropertyInputBufferSize = 'i'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'icl '
	KAudioCodecPropertyCurrentInputFormat               KAudioCodecPropertyInputBufferSize = 'i'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'ifmt'
	KAudioCodecPropertyCurrentInputSampleRate           KAudioCodecPropertyInputBufferSize = 'c'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'cisr'
	KAudioCodecPropertyCurrentOutputChannelLayout       KAudioCodecPropertyInputBufferSize = 'o'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'ocl '
	KAudioCodecPropertyCurrentOutputFormat              KAudioCodecPropertyInputBufferSize = 'o'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'ofmt'
	KAudioCodecPropertyCurrentOutputSampleRate          KAudioCodecPropertyInputBufferSize = 'c'<<24 | 'o'<<16 | 's'<<8 | 'r' // 'cosr'
	KAudioCodecPropertyCurrentTargetBitRate             KAudioCodecPropertyInputBufferSize = 'b'<<24 | 'r'<<16 | 'a'<<8 | 't' // 'brat'
	KAudioCodecPropertyDelayMode                        KAudioCodecPropertyInputBufferSize = 'd'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'dmod'
	KAudioCodecPropertyDynamicRangeControlConfiguration KAudioCodecPropertyInputBufferSize = 'c'<<24 | 'd'<<16 | 'r'<<8 | 'c' // 'cdrc'
	KAudioCodecPropertyDynamicRangeControlMode          KAudioCodecPropertyInputBufferSize = 'm'<<24 | 'd'<<16 | 'r'<<8 | 'c' // 'mdrc'
	KAudioCodecPropertyEmploysDependentPackets          KAudioCodecPropertyInputBufferSize = 'd'<<24 | 'p'<<16 | 'k'<<8 | '?' // 'dpk?'
	KAudioCodecPropertyFormatList                       KAudioCodecPropertyInputBufferSize = 'a'<<24 | 'c'<<16 | 'f'<<8 | 'l' // 'acfl'
	KAudioCodecPropertyHasVariablePacketByteSizes       KAudioCodecPropertyInputBufferSize = 'v'<<24 | 'p'<<16 | 'k'<<8 | '?' // 'vpk?'
	KAudioCodecPropertyInputBufferSizeValue             KAudioCodecPropertyInputBufferSize = 't'<<24 | 'b'<<16 | 'u'<<8 | 'f' // 'tbuf'
	KAudioCodecPropertyIsInitialized                    KAudioCodecPropertyInputBufferSize = 'i'<<24 | 'n'<<16 | 'i'<<8 | 't' // 'init'
	KAudioCodecPropertyMagicCookie                      KAudioCodecPropertyInputBufferSize = 'k'<<24 | 'u'<<16 | 'k'<<8 | 'i' // 'kuki'
	KAudioCodecPropertyMaximumPacketByteSize            KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'a'<<16 | 'k'<<8 | 'b' // 'pakb'
	KAudioCodecPropertyPacketFrameSize                  KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'a'<<16 | 'k'<<8 | 'f' // 'pakf'
	KAudioCodecPropertyPacketSizeLimitForVBR            KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'a'<<16 | 'k'<<8 | 'l' // 'pakl'
	KAudioCodecPropertyPaddedZeros                      KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'a'<<16 | 'd'<<8 | '0' // 'pad0'
	KAudioCodecPropertyPrimeInfo                        KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'r'<<16 | 'i'<<8 | 'm' // 'prim'
	KAudioCodecPropertyPrimeMethod                      KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'r'<<16 | 'm'<<8 | 'm' // 'prmm'
	KAudioCodecPropertyProgramTargetLevel               KAudioCodecPropertyInputBufferSize = 'p'<<24 | 'p'<<16 | 't'<<8 | 'l' // 'pptl'
	KAudioCodecPropertyProgramTargetLevelConstant       KAudioCodecPropertyInputBufferSize = 'p'<<24 | 't'<<16 | 'l'<<8 | 'c' // 'ptlc'
	KAudioCodecPropertyQualitySetting                   KAudioCodecPropertyInputBufferSize = 's'<<24 | 'r'<<16 | 'c'<<8 | 'q' // 'srcq'
	KAudioCodecPropertyRecommendedBitRateRange          KAudioCodecPropertyInputBufferSize = 'b'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'brtr'
	KAudioCodecPropertySettings                         KAudioCodecPropertyInputBufferSize = 'a'<<24 | 'c'<<16 | 's'<<8 | ' ' // 'acs '
	KAudioCodecPropertySoundQualityForVBR               KAudioCodecPropertyInputBufferSize = 'v'<<24 | 'b'<<16 | 'r'<<8 | 'q' // 'vbrq'
	KAudioCodecPropertyUsedInputBufferSize              KAudioCodecPropertyInputBufferSize = 'u'<<24 | 'b'<<16 | 'u'<<8 | 'f' // 'ubuf'
)

func (e KAudioCodecPropertyInputBufferSize) String() string {
	switch e {
	case KAudioCodecPropertyASPFrequency:
		return "KAudioCodecPropertyASPFrequency"
	case KAudioCodecPropertyAdjustCompressionProfile:
		return "KAudioCodecPropertyAdjustCompressionProfile"
	case KAudioCodecPropertyAdjustLocalQuality:
		return "KAudioCodecPropertyAdjustLocalQuality"
	case KAudioCodecPropertyAdjustTargetLevel:
		return "KAudioCodecPropertyAdjustTargetLevel"
	case KAudioCodecPropertyAdjustTargetLevelConstant:
		return "KAudioCodecPropertyAdjustTargetLevelConstant"
	case KAudioCodecPropertyApplicableBitRateRange:
		return "KAudioCodecPropertyApplicableBitRateRange"
	case KAudioCodecPropertyApplicableInputSampleRates:
		return "KAudioCodecPropertyApplicableInputSampleRates"
	case KAudioCodecPropertyApplicableOutputSampleRates:
		return "KAudioCodecPropertyApplicableOutputSampleRates"
	case KAudioCodecPropertyBitRateControlMode:
		return "KAudioCodecPropertyBitRateControlMode"
	case KAudioCodecPropertyBitRateForVBR:
		return "KAudioCodecPropertyBitRateForVBR"
	case KAudioCodecPropertyContentSource:
		return "KAudioCodecPropertyContentSource"
	case KAudioCodecPropertyCurrentInputChannelLayout:
		return "KAudioCodecPropertyCurrentInputChannelLayout"
	case KAudioCodecPropertyCurrentInputFormat:
		return "KAudioCodecPropertyCurrentInputFormat"
	case KAudioCodecPropertyCurrentInputSampleRate:
		return "KAudioCodecPropertyCurrentInputSampleRate"
	case KAudioCodecPropertyCurrentOutputChannelLayout:
		return "KAudioCodecPropertyCurrentOutputChannelLayout"
	case KAudioCodecPropertyCurrentOutputFormat:
		return "KAudioCodecPropertyCurrentOutputFormat"
	case KAudioCodecPropertyCurrentOutputSampleRate:
		return "KAudioCodecPropertyCurrentOutputSampleRate"
	case KAudioCodecPropertyCurrentTargetBitRate:
		return "KAudioCodecPropertyCurrentTargetBitRate"
	case KAudioCodecPropertyDelayMode:
		return "KAudioCodecPropertyDelayMode"
	case KAudioCodecPropertyDynamicRangeControlConfiguration:
		return "KAudioCodecPropertyDynamicRangeControlConfiguration"
	case KAudioCodecPropertyDynamicRangeControlMode:
		return "KAudioCodecPropertyDynamicRangeControlMode"
	case KAudioCodecPropertyEmploysDependentPackets:
		return "KAudioCodecPropertyEmploysDependentPackets"
	case KAudioCodecPropertyFormatList:
		return "KAudioCodecPropertyFormatList"
	case KAudioCodecPropertyHasVariablePacketByteSizes:
		return "KAudioCodecPropertyHasVariablePacketByteSizes"
	case KAudioCodecPropertyInputBufferSizeValue:
		return "KAudioCodecPropertyInputBufferSizeValue"
	case KAudioCodecPropertyIsInitialized:
		return "KAudioCodecPropertyIsInitialized"
	case KAudioCodecPropertyMagicCookie:
		return "KAudioCodecPropertyMagicCookie"
	case KAudioCodecPropertyMaximumPacketByteSize:
		return "KAudioCodecPropertyMaximumPacketByteSize"
	case KAudioCodecPropertyPacketFrameSize:
		return "KAudioCodecPropertyPacketFrameSize"
	case KAudioCodecPropertyPacketSizeLimitForVBR:
		return "KAudioCodecPropertyPacketSizeLimitForVBR"
	case KAudioCodecPropertyPaddedZeros:
		return "KAudioCodecPropertyPaddedZeros"
	case KAudioCodecPropertyPrimeInfo:
		return "KAudioCodecPropertyPrimeInfo"
	case KAudioCodecPropertyPrimeMethod:
		return "KAudioCodecPropertyPrimeMethod"
	case KAudioCodecPropertyProgramTargetLevel:
		return "KAudioCodecPropertyProgramTargetLevel"
	case KAudioCodecPropertyProgramTargetLevelConstant:
		return "KAudioCodecPropertyProgramTargetLevelConstant"
	case KAudioCodecPropertyQualitySetting:
		return "KAudioCodecPropertyQualitySetting"
	case KAudioCodecPropertyRecommendedBitRateRange:
		return "KAudioCodecPropertyRecommendedBitRateRange"
	case KAudioCodecPropertySettings:
		return "KAudioCodecPropertySettings"
	case KAudioCodecPropertySoundQualityForVBR:
		return "KAudioCodecPropertySoundQualityForVBR"
	case KAudioCodecPropertyUsedInputBufferSize:
		return "KAudioCodecPropertyUsedInputBufferSize"
	default:
		return fmt.Sprintf("KAudioCodecPropertyInputBufferSize(%d)", e)
	}
}

const KAudioCodecPropertyMinimumDelayMode uint32 = 'm'<<24 | 'd'<<16 | 'e'<<8 | 'l' // 'mdel'

type KAudioCodecPropertyNameCFString uint32

const (
	KAudioCodecPropertyFormatCFString       KAudioCodecPropertyNameCFString = 'l'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'lfor'
	KAudioCodecPropertyManufacturerCFString KAudioCodecPropertyNameCFString = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KAudioCodecPropertyNameCFStringValue    KAudioCodecPropertyNameCFString = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
)

func (e KAudioCodecPropertyNameCFString) String() string {
	switch e {
	case KAudioCodecPropertyFormatCFString:
		return "KAudioCodecPropertyFormatCFString"
	case KAudioCodecPropertyManufacturerCFString:
		return "KAudioCodecPropertyManufacturerCFString"
	case KAudioCodecPropertyNameCFStringValue:
		return "KAudioCodecPropertyNameCFStringValue"
	default:
		return fmt.Sprintf("KAudioCodecPropertyNameCFString(%d)", e)
	}
}

type KAudioCodecPropertyRequiresPacketDescription uint32

const (
	KAudioCodecBitRateFormatValue                     KAudioCodecPropertyRequiresPacketDescription = 'a'<<24 | 'c'<<16 | 'b'<<8 | 'f' // 'acbf'
	KAudioCodecDoesSampleRateConversion               KAudioCodecPropertyRequiresPacketDescription = 'l'<<24 | 'm'<<16 | 'r'<<8 | 'c' // 'lmrc'
	KAudioCodecExtendFrequencies                      KAudioCodecPropertyRequiresPacketDescription = 'a'<<24 | 'c'<<16 | 'e'<<8 | 'f' // 'acef'
	KAudioCodecInputFormatsForOutputFormat            KAudioCodecPropertyRequiresPacketDescription = 'i'<<24 | 'f'<<16 | '4'<<8 | 'o' // 'if4o'
	KAudioCodecOutputFormatsForInputFormat            KAudioCodecPropertyRequiresPacketDescription = 'o'<<24 | 'f'<<16 | '4'<<8 | 'i' // 'of4i'
	KAudioCodecOutputPrecedenceValue                  KAudioCodecPropertyRequiresPacketDescription = 'o'<<24 | 'p'<<16 | 'p'<<8 | 'r' // 'oppr'
	KAudioCodecPropertyAvailableBitRates              KAudioCodecPropertyRequiresPacketDescription = 'b'<<24 | 'r'<<16 | 't'<<8 | '#' // 'brt#'
	KAudioCodecPropertyAvailableInputChannelLayouts   KAudioCodecPropertyRequiresPacketDescription = 'a'<<24 | 'i'<<16 | 'c'<<8 | 'l' // 'aicl'
	KAudioCodecPropertyAvailableOutputChannelLayouts  KAudioCodecPropertyRequiresPacketDescription = 'a'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'aocl'
	KAudioCodecPropertyInputChannelLayout             KAudioCodecPropertyRequiresPacketDescription = 'i'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'icl '
	KAudioCodecPropertyOutputChannelLayout            KAudioCodecPropertyRequiresPacketDescription = 'o'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'ocl '
	KAudioCodecPropertyRequiresPacketDescriptionValue KAudioCodecPropertyRequiresPacketDescription = 'p'<<24 | 'a'<<16 | 'k'<<8 | 'd' // 'pakd'
	KAudioCodecPropertyZeroFramesPadded               KAudioCodecPropertyRequiresPacketDescription = 'p'<<24 | 'a'<<16 | 'd'<<8 | '0' // 'pad0'
	KAudioCodecUseRecommendedSampleRate               KAudioCodecPropertyRequiresPacketDescription = 'u'<<24 | 'r'<<16 | 's'<<8 | 'r' // 'ursr'
)

func (e KAudioCodecPropertyRequiresPacketDescription) String() string {
	switch e {
	case KAudioCodecBitRateFormatValue:
		return "KAudioCodecBitRateFormatValue"
	case KAudioCodecDoesSampleRateConversion:
		return "KAudioCodecDoesSampleRateConversion"
	case KAudioCodecExtendFrequencies:
		return "KAudioCodecExtendFrequencies"
	case KAudioCodecInputFormatsForOutputFormat:
		return "KAudioCodecInputFormatsForOutputFormat"
	case KAudioCodecOutputFormatsForInputFormat:
		return "KAudioCodecOutputFormatsForInputFormat"
	case KAudioCodecOutputPrecedenceValue:
		return "KAudioCodecOutputPrecedenceValue"
	case KAudioCodecPropertyAvailableBitRates:
		return "KAudioCodecPropertyAvailableBitRates"
	case KAudioCodecPropertyAvailableInputChannelLayouts:
		return "KAudioCodecPropertyAvailableInputChannelLayouts"
	case KAudioCodecPropertyAvailableOutputChannelLayouts:
		return "KAudioCodecPropertyAvailableOutputChannelLayouts"
	case KAudioCodecPropertyInputChannelLayout:
		return "KAudioCodecPropertyInputChannelLayout"
	case KAudioCodecPropertyOutputChannelLayout:
		return "KAudioCodecPropertyOutputChannelLayout"
	case KAudioCodecPropertyRequiresPacketDescriptionValue:
		return "KAudioCodecPropertyRequiresPacketDescriptionValue"
	case KAudioCodecPropertyZeroFramesPadded:
		return "KAudioCodecPropertyZeroFramesPadded"
	case KAudioCodecUseRecommendedSampleRate:
		return "KAudioCodecUseRecommendedSampleRate"
	default:
		return fmt.Sprintf("KAudioCodecPropertyRequiresPacketDescription(%d)", e)
	}
}

type KAudioCodecPropertySupportedInputFormats uint32

const (
	KAudioCodecPropertyAvailableBitRateRange            KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'b'<<16 | 'r'<<8 | 't' // 'abrt'
	KAudioCodecPropertyAvailableInputChannelLayoutTags  KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'i'<<16 | 'c'<<8 | 'l' // 'aicl'
	KAudioCodecPropertyAvailableInputSampleRates        KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'aisr'
	KAudioCodecPropertyAvailableNumberChannels          KAudioCodecPropertySupportedInputFormats = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'c' // 'cmnc'
	KAudioCodecPropertyAvailableOutputChannelLayoutTags KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'aocl'
	KAudioCodecPropertyAvailableOutputSampleRates       KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'o'<<16 | 's'<<8 | 'r' // 'aosr'
	KAudioCodecPropertyDoesSampleRateConversion         KAudioCodecPropertySupportedInputFormats = 'l'<<24 | 'm'<<16 | 'r'<<8 | 'c' // 'lmrc'
	KAudioCodecPropertyFormatInfo                       KAudioCodecPropertySupportedInputFormats = 'a'<<24 | 'c'<<16 | 'f'<<8 | 'i' // 'acfi'
	KAudioCodecPropertyInputFormatsForOutputFormat      KAudioCodecPropertySupportedInputFormats = 'i'<<24 | 'f'<<16 | '4'<<8 | 'o' // 'if4o'
	KAudioCodecPropertyMinimumNumberInputPackets        KAudioCodecPropertySupportedInputFormats = 'm'<<24 | 'n'<<16 | 'i'<<8 | 'p' // 'mnip'
	KAudioCodecPropertyMinimumNumberOutputPackets       KAudioCodecPropertySupportedInputFormats = 'm'<<24 | 'n'<<16 | 'o'<<8 | 'p' // 'mnop'
	KAudioCodecPropertyOutputFormatsForInputFormat      KAudioCodecPropertySupportedInputFormats = 'o'<<24 | 'f'<<16 | '4'<<8 | 'i' // 'of4i'
	KAudioCodecPropertySupportedInputFormatsValue       KAudioCodecPropertySupportedInputFormats = 'i'<<24 | 'f'<<16 | 'm'<<8 | '#' // 'ifm#'
	KAudioCodecPropertySupportedOutputFormats           KAudioCodecPropertySupportedInputFormats = 'o'<<24 | 'f'<<16 | 'm'<<8 | '#' // 'ofm#'
)

func (e KAudioCodecPropertySupportedInputFormats) String() string {
	switch e {
	case KAudioCodecPropertyAvailableBitRateRange:
		return "KAudioCodecPropertyAvailableBitRateRange"
	case KAudioCodecPropertyAvailableInputChannelLayoutTags:
		return "KAudioCodecPropertyAvailableInputChannelLayoutTags"
	case KAudioCodecPropertyAvailableInputSampleRates:
		return "KAudioCodecPropertyAvailableInputSampleRates"
	case KAudioCodecPropertyAvailableNumberChannels:
		return "KAudioCodecPropertyAvailableNumberChannels"
	case KAudioCodecPropertyAvailableOutputChannelLayoutTags:
		return "KAudioCodecPropertyAvailableOutputChannelLayoutTags"
	case KAudioCodecPropertyAvailableOutputSampleRates:
		return "KAudioCodecPropertyAvailableOutputSampleRates"
	case KAudioCodecPropertyDoesSampleRateConversion:
		return "KAudioCodecPropertyDoesSampleRateConversion"
	case KAudioCodecPropertyFormatInfo:
		return "KAudioCodecPropertyFormatInfo"
	case KAudioCodecPropertyInputFormatsForOutputFormat:
		return "KAudioCodecPropertyInputFormatsForOutputFormat"
	case KAudioCodecPropertyMinimumNumberInputPackets:
		return "KAudioCodecPropertyMinimumNumberInputPackets"
	case KAudioCodecPropertyMinimumNumberOutputPackets:
		return "KAudioCodecPropertyMinimumNumberOutputPackets"
	case KAudioCodecPropertyOutputFormatsForInputFormat:
		return "KAudioCodecPropertyOutputFormatsForInputFormat"
	case KAudioCodecPropertySupportedInputFormatsValue:
		return "KAudioCodecPropertySupportedInputFormatsValue"
	case KAudioCodecPropertySupportedOutputFormats:
		return "KAudioCodecPropertySupportedOutputFormats"
	default:
		return fmt.Sprintf("KAudioCodecPropertySupportedInputFormats(%d)", e)
	}
}

type KAudioCodecQuality uint32

const (
	KAudioCodecQuality_High   KAudioCodecQuality = 0x60
	KAudioCodecQuality_Low    KAudioCodecQuality = 0x20
	KAudioCodecQuality_Max    KAudioCodecQuality = 0x7f
	KAudioCodecQuality_Medium KAudioCodecQuality = 0x40
	KAudioCodecQuality_Min    KAudioCodecQuality = 0
)

func (e KAudioCodecQuality) String() string {
	switch e {
	case KAudioCodecQuality_High:
		return "KAudioCodecQuality_High"
	case KAudioCodecQuality_Low:
		return "KAudioCodecQuality_Low"
	case KAudioCodecQuality_Max:
		return "KAudioCodecQuality_Max"
	case KAudioCodecQuality_Medium:
		return "KAudioCodecQuality_Medium"
	case KAudioCodecQuality_Min:
		return "KAudioCodecQuality_Min"
	default:
		return fmt.Sprintf("KAudioCodecQuality(%d)", e)
	}
}

type KAudioComponentErr int32

const (
	KAudioComponentErr_DuplicateDescription   KAudioComponentErr = -66752
	KAudioComponentErr_InitializationTimedOut KAudioComponentErr = -66747
	KAudioComponentErr_InvalidFormat          KAudioComponentErr = -66746
	KAudioComponentErr_NotPermitted           KAudioComponentErr = -66748
	KAudioComponentErr_TooManyInstances       KAudioComponentErr = -66750
	KAudioComponentErr_UnsupportedType        KAudioComponentErr = -66751
)

func (e KAudioComponentErr) String() string {
	switch e {
	case KAudioComponentErr_DuplicateDescription:
		return "KAudioComponentErr_DuplicateDescription"
	case KAudioComponentErr_InitializationTimedOut:
		return "KAudioComponentErr_InitializationTimedOut"
	case KAudioComponentErr_InvalidFormat:
		return "KAudioComponentErr_InvalidFormat"
	case KAudioComponentErr_NotPermitted:
		return "KAudioComponentErr_NotPermitted"
	case KAudioComponentErr_TooManyInstances:
		return "KAudioComponentErr_TooManyInstances"
	case KAudioComponentErr_UnsupportedType:
		return "KAudioComponentErr_UnsupportedType"
	default:
		return fmt.Sprintf("KAudioComponentErr(%d)", e)
	}
}

type KAudioConverterProperty uint32

const (
	KAudioConverterPropertyDitherBitDepth KAudioConverterProperty = 'd'<<24 | 'b'<<16 | 'i'<<8 | 't' // 'dbit'
	KAudioConverterPropertyDithering      KAudioConverterProperty = 'd'<<24 | 'i'<<16 | 't'<<8 | 'h' // 'dith'
)

func (e KAudioConverterProperty) String() string {
	switch e {
	case KAudioConverterPropertyDitherBitDepth:
		return "KAudioConverterPropertyDitherBitDepth"
	case KAudioConverterPropertyDithering:
		return "KAudioConverterPropertyDithering"
	default:
		return fmt.Sprintf("KAudioConverterProperty(%d)", e)
	}
}

type KAudioConverterPropertyMaximumInputBufferSize uint32

const (
	// KAudioConverterPropertyMaximumInputBufferSizeValue: Deprecated.
	KAudioConverterPropertyMaximumInputBufferSizeValue KAudioConverterPropertyMaximumInputBufferSize = 'x'<<24 | 'i'<<16 | 'b'<<8 | 's' // 'xibs'
	// KAudioConverterSampleRateConverterAlgorithm: A value that indicates the sample rate conversion algorithm.
	KAudioConverterSampleRateConverterAlgorithm KAudioConverterPropertyMaximumInputBufferSize = 's'<<24 | 'r'<<16 | 'c'<<8 | 'i' // 'srci'
)

func (e KAudioConverterPropertyMaximumInputBufferSize) String() string {
	switch e {
	case KAudioConverterPropertyMaximumInputBufferSizeValue:
		return "KAudioConverterPropertyMaximumInputBufferSizeValue"
	case KAudioConverterSampleRateConverterAlgorithm:
		return "KAudioConverterSampleRateConverterAlgorithm"
	default:
		return fmt.Sprintf("KAudioConverterPropertyMaximumInputBufferSize(%d)", e)
	}
}

type KAudioConverterPropertyMinimumInputBufferSize uint32

const (
	// KAudioConverterApplicableEncodeBitRates: An array of [AudioValueRange] structures that describes applicable bit rates based on current settings.
	KAudioConverterApplicableEncodeBitRates KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'e'<<16 | 'b'<<8 | 'r' // 'aebr'
	// KAudioConverterApplicableEncodeSampleRates: An array of [AudioValueRange] structures that describes applicable sample rates based on current settings.
	KAudioConverterApplicableEncodeSampleRates KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'e'<<16 | 's'<<8 | 'r' // 'aesr'
	// KAudioConverterAvailableEncodeBitRates: An array of [AudioValueRange] structures that describes the available bit rates based on the input format.
	KAudioConverterAvailableEncodeBitRates KAudioConverterPropertyMinimumInputBufferSize = 'v'<<24 | 'e'<<16 | 'b'<<8 | 'r' // 'vebr'
	// KAudioConverterAvailableEncodeChannelLayoutTags: An array of [AudioChannelLayoutTag] values for the format and number of channels specified in the encoder’s input format.
	KAudioConverterAvailableEncodeChannelLayoutTags KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'e'<<16 | 'c'<<8 | 'l' // 'aecl'
	// KAudioConverterAvailableEncodeSampleRates: An array of [AudioValueRange] structures that describes the available sample rates based on the input format.
	KAudioConverterAvailableEncodeSampleRates KAudioConverterPropertyMinimumInputBufferSize = 'v'<<24 | 'e'<<16 | 's'<<8 | 'r' // 'vesr'
	// KAudioConverterChannelMap: An array of [SInt32] values that specify an input-to-output channel mapping.
	KAudioConverterChannelMap KAudioConverterPropertyMinimumInputBufferSize = 'c'<<24 | 'h'<<16 | 'm'<<8 | 'p' // 'chmp'
	// KAudioConverterCodecQuality: The rendering quality of a codec.
	KAudioConverterCodecQuality KAudioConverterPropertyMinimumInputBufferSize = 'c'<<24 | 'd'<<16 | 'q'<<8 | 'u' // 'cdqu'
	// KAudioConverterCompressionMagicCookie: A `void*` value that points to memory set up by the caller.
	KAudioConverterCompressionMagicCookie KAudioConverterPropertyMinimumInputBufferSize = 'c'<<24 | 'm'<<16 | 'g'<<8 | 'c' // 'cmgc'
	// KAudioConverterCurrentInputStreamDescription: The current, completely specified input [AudioStreamBasicDescription] structure.
	KAudioConverterCurrentInputStreamDescription KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'c'<<16 | 'i'<<8 | 'd' // 'acid'
	// KAudioConverterCurrentOutputStreamDescription: The current, completely specified output [AudioStreamBasicDescription] structure.
	KAudioConverterCurrentOutputStreamDescription KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'c'<<16 | 'o'<<8 | 'd' // 'acod'
	// KAudioConverterDecompressionMagicCookie: A `void*` value that points to memory set up by the caller.
	KAudioConverterDecompressionMagicCookie KAudioConverterPropertyMinimumInputBufferSize = 'd'<<24 | 'm'<<16 | 'g'<<8 | 'c' // 'dmgc'
	// KAudioConverterEncodeAdjustableSampleRate: A [Float64] value that specifies an output sample rate.
	KAudioConverterEncodeAdjustableSampleRate KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'j'<<16 | 's'<<8 | 'r' // 'ajsr'
	// KAudioConverterEncodeBitRate: A [UInt32] value containing the number of bits per second to aim for when encoding data.
	KAudioConverterEncodeBitRate KAudioConverterPropertyMinimumInputBufferSize = 'b'<<24 | 'r'<<16 | 'a'<<8 | 't' // 'brat'
	// KAudioConverterInputChannelLayout: An [AudioChannelLayout] structure that specifies an audio converter’s input channel layout.
	KAudioConverterInputChannelLayout KAudioConverterPropertyMinimumInputBufferSize = 'i'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'icl '
	// KAudioConverterOutputChannelLayout: An [AudioChannelLayout] structure that specifies an audio converter’s output channel layout.
	KAudioConverterOutputChannelLayout KAudioConverterPropertyMinimumInputBufferSize = 'o'<<24 | 'c'<<16 | 'l'<<8 | ' ' // 'ocl '
	// KAudioConverterPrimeInfo: An AudioConverterPrimeInfo structure.
	KAudioConverterPrimeInfo KAudioConverterPropertyMinimumInputBufferSize = 'p'<<24 | 'r'<<16 | 'i'<<8 | 'm' // 'prim'
	// KAudioConverterPrimeMethod: The priming method, usually for sample-rate conversion.
	KAudioConverterPrimeMethod KAudioConverterPropertyMinimumInputBufferSize = 'p'<<24 | 'r'<<16 | 'm'<<8 | 'm' // 'prmm'
	// KAudioConverterPropertyBitDepthHint: A [UInt32] value that designates the source bit depth to preserve.
	KAudioConverterPropertyBitDepthHint KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'c'<<16 | 'b'<<8 | 'd' // 'acbd'
	// KAudioConverterPropertyCalculateInputBufferSize: A [UInt32] value that on input holds a size, in bytes, that is desired for the output data.
	KAudioConverterPropertyCalculateInputBufferSize KAudioConverterPropertyMinimumInputBufferSize = 'c'<<24 | 'i'<<16 | 'b'<<8 | 's' // 'cibs'
	// KAudioConverterPropertyCalculateOutputBufferSize: A [UInt32] value that on input holds a size, in bytes, that is desired for the input data.
	KAudioConverterPropertyCalculateOutputBufferSize KAudioConverterPropertyMinimumInputBufferSize = 'c'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'cobs'
	KAudioConverterPropertyChannelMixMap             KAudioConverterPropertyMinimumInputBufferSize = 'm'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'mmap'
	// KAudioConverterPropertyFormatList: An array of [AudioFormatListItem] structures that describes the set of data formats produced by the encoder end of an audio converter.
	KAudioConverterPropertyFormatList KAudioConverterPropertyMinimumInputBufferSize = 'f'<<24 | 'l'<<16 | 's'<<8 | 't' // 'flst'
	// KAudioConverterPropertyInputCodecParameters: The value of this property varies from format to format and is considered private to the format.
	KAudioConverterPropertyInputCodecParameters KAudioConverterPropertyMinimumInputBufferSize = 'i'<<24 | 'c'<<16 | 'd'<<8 | 'p' // 'icdp'
	// KAudioConverterPropertyMaximumInputPacketSize: A [UInt32] value that indicates the size, in bytes, of the largest single packet of data in the input format.
	KAudioConverterPropertyMaximumInputPacketSize KAudioConverterPropertyMinimumInputBufferSize = 'x'<<24 | 'i'<<16 | 'p'<<8 | 's' // 'xips'
	// KAudioConverterPropertyMaximumOutputPacketSize: A [UInt32] value that indicates the size, in bytes, of the largest single packet of data in the output format.
	KAudioConverterPropertyMaximumOutputPacketSize KAudioConverterPropertyMinimumInputBufferSize = 'x'<<24 | 'o'<<16 | 'p'<<8 | 's' // 'xops'
	// KAudioConverterPropertyMinimumInputBufferSizeValue: A [UInt32] value that indicates the size, in bytes, of the smallest buffer of input data that can be supplied via the audio converter input callback or as the input to the AudioConverterConvertBuffer(_:_:_:_:_:) function.
	KAudioConverterPropertyMinimumInputBufferSizeValue KAudioConverterPropertyMinimumInputBufferSize = 'm'<<24 | 'i'<<16 | 'b'<<8 | 's' // 'mibs'
	// KAudioConverterPropertyMinimumOutputBufferSize: A [UInt32] value that indicates the size, in bytes, of the smallest buffer of output data that can be supplied to AudioConverterFillComplexBuffer or as the output to AudioConverterConvertBuffer
	KAudioConverterPropertyMinimumOutputBufferSize KAudioConverterPropertyMinimumInputBufferSize = 'm'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'mobs'
	// KAudioConverterPropertyOutputCodecParameters: The value of this property varies from format to format and is considered private to the format.
	KAudioConverterPropertyOutputCodecParameters KAudioConverterPropertyMinimumInputBufferSize = 'o'<<24 | 'c'<<16 | 'd'<<8 | 'p' // 'ocdp'
	KAudioConverterPropertyPerformDownmix        KAudioConverterPropertyMinimumInputBufferSize = 'd'<<24 | 'm'<<16 | 'i'<<8 | 'x' // 'dmix'
	// KAudioConverterPropertySettings: An array (of type [CFArray]) of property settings for converters.
	KAudioConverterPropertySettings KAudioConverterPropertyMinimumInputBufferSize = 'a'<<24 | 'c'<<16 | 'p'<<8 | 's' // 'acps'
	// KAudioConverterSampleRateConverterComplexityValue: The sample rate conversion algorithm.
	KAudioConverterSampleRateConverterComplexityValue KAudioConverterPropertyMinimumInputBufferSize = 's'<<24 | 'r'<<16 | 'c'<<8 | 'a' // 'srca'
	// KAudioConverterSampleRateConverterInitialPhase: A [Float64] value equal to `0.0`.
	KAudioConverterSampleRateConverterInitialPhase KAudioConverterPropertyMinimumInputBufferSize = 's'<<24 | 'r'<<16 | 'c'<<8 | 'p' // 'srcp'
	// KAudioConverterSampleRateConverterQuality: The rendering quality of the sample rate converter.
	KAudioConverterSampleRateConverterQuality KAudioConverterPropertyMinimumInputBufferSize = 's'<<24 | 'r'<<16 | 'c'<<8 | 'q' // 'srcq'
)

func (e KAudioConverterPropertyMinimumInputBufferSize) String() string {
	switch e {
	case KAudioConverterApplicableEncodeBitRates:
		return "KAudioConverterApplicableEncodeBitRates"
	case KAudioConverterApplicableEncodeSampleRates:
		return "KAudioConverterApplicableEncodeSampleRates"
	case KAudioConverterAvailableEncodeBitRates:
		return "KAudioConverterAvailableEncodeBitRates"
	case KAudioConverterAvailableEncodeChannelLayoutTags:
		return "KAudioConverterAvailableEncodeChannelLayoutTags"
	case KAudioConverterAvailableEncodeSampleRates:
		return "KAudioConverterAvailableEncodeSampleRates"
	case KAudioConverterChannelMap:
		return "KAudioConverterChannelMap"
	case KAudioConverterCodecQuality:
		return "KAudioConverterCodecQuality"
	case KAudioConverterCompressionMagicCookie:
		return "KAudioConverterCompressionMagicCookie"
	case KAudioConverterCurrentInputStreamDescription:
		return "KAudioConverterCurrentInputStreamDescription"
	case KAudioConverterCurrentOutputStreamDescription:
		return "KAudioConverterCurrentOutputStreamDescription"
	case KAudioConverterDecompressionMagicCookie:
		return "KAudioConverterDecompressionMagicCookie"
	case KAudioConverterEncodeAdjustableSampleRate:
		return "KAudioConverterEncodeAdjustableSampleRate"
	case KAudioConverterEncodeBitRate:
		return "KAudioConverterEncodeBitRate"
	case KAudioConverterInputChannelLayout:
		return "KAudioConverterInputChannelLayout"
	case KAudioConverterOutputChannelLayout:
		return "KAudioConverterOutputChannelLayout"
	case KAudioConverterPrimeInfo:
		return "KAudioConverterPrimeInfo"
	case KAudioConverterPrimeMethod:
		return "KAudioConverterPrimeMethod"
	case KAudioConverterPropertyBitDepthHint:
		return "KAudioConverterPropertyBitDepthHint"
	case KAudioConverterPropertyCalculateInputBufferSize:
		return "KAudioConverterPropertyCalculateInputBufferSize"
	case KAudioConverterPropertyCalculateOutputBufferSize:
		return "KAudioConverterPropertyCalculateOutputBufferSize"
	case KAudioConverterPropertyChannelMixMap:
		return "KAudioConverterPropertyChannelMixMap"
	case KAudioConverterPropertyFormatList:
		return "KAudioConverterPropertyFormatList"
	case KAudioConverterPropertyInputCodecParameters:
		return "KAudioConverterPropertyInputCodecParameters"
	case KAudioConverterPropertyMaximumInputPacketSize:
		return "KAudioConverterPropertyMaximumInputPacketSize"
	case KAudioConverterPropertyMaximumOutputPacketSize:
		return "KAudioConverterPropertyMaximumOutputPacketSize"
	case KAudioConverterPropertyMinimumInputBufferSizeValue:
		return "KAudioConverterPropertyMinimumInputBufferSizeValue"
	case KAudioConverterPropertyMinimumOutputBufferSize:
		return "KAudioConverterPropertyMinimumOutputBufferSize"
	case KAudioConverterPropertyOutputCodecParameters:
		return "KAudioConverterPropertyOutputCodecParameters"
	case KAudioConverterPropertyPerformDownmix:
		return "KAudioConverterPropertyPerformDownmix"
	case KAudioConverterPropertySettings:
		return "KAudioConverterPropertySettings"
	case KAudioConverterSampleRateConverterComplexityValue:
		return "KAudioConverterSampleRateConverterComplexityValue"
	case KAudioConverterSampleRateConverterInitialPhase:
		return "KAudioConverterSampleRateConverterInitialPhase"
	case KAudioConverterSampleRateConverterQuality:
		return "KAudioConverterSampleRateConverterQuality"
	default:
		return fmt.Sprintf("KAudioConverterPropertyMinimumInputBufferSize(%d)", e)
	}
}

type KAudioConverterQuality uint32

const (
	// KAudioConverterQuality_High: Specifies high sample rate conversion quality.
	KAudioConverterQuality_High KAudioConverterQuality = 0x60
	// KAudioConverterQuality_Low: Specifies low sample rate conversion quality.
	KAudioConverterQuality_Low KAudioConverterQuality = 0x20
	// KAudioConverterQuality_Max: Specifies maximum sample-rate conversion quality.
	KAudioConverterQuality_Max KAudioConverterQuality = 0x7f
	// KAudioConverterQuality_Medium: Specifies medium sample rate conversion quality.
	KAudioConverterQuality_Medium KAudioConverterQuality = 0x40
	// KAudioConverterQuality_Min: Specifies minimum sample rate conversion quality.
	KAudioConverterQuality_Min KAudioConverterQuality = 0
)

func (e KAudioConverterQuality) String() string {
	switch e {
	case KAudioConverterQuality_High:
		return "KAudioConverterQuality_High"
	case KAudioConverterQuality_Low:
		return "KAudioConverterQuality_Low"
	case KAudioConverterQuality_Max:
		return "KAudioConverterQuality_Max"
	case KAudioConverterQuality_Medium:
		return "KAudioConverterQuality_Medium"
	case KAudioConverterQuality_Min:
		return "KAudioConverterQuality_Min"
	default:
		return fmt.Sprintf("KAudioConverterQuality(%d)", e)
	}
}

type KAudioConverterSampleRateConverterComplexity uint32

const (
	// KAudioConverterSampleRateConverterComplexity_Linear: Specifies linear interpolation for sample rate conversion.
	KAudioConverterSampleRateConverterComplexity_Linear KAudioConverterSampleRateConverterComplexity = 'l'<<24 | 'i'<<16 | 'n'<<8 | 'e' // 'line'
	// KAudioConverterSampleRateConverterComplexity_Mastering: Specifies a mastering-quality sample rate conversion algorithm.
	KAudioConverterSampleRateConverterComplexity_Mastering    KAudioConverterSampleRateConverterComplexity = 'b'<<24 | 'a'<<16 | 't'<<8 | 's' // 'bats'
	KAudioConverterSampleRateConverterComplexity_MinimumPhase KAudioConverterSampleRateConverterComplexity = 'm'<<24 | 'i'<<16 | 'n'<<8 | 'p' // 'minp'
	// KAudioConverterSampleRateConverterComplexity_Normal: Specifies the normal-complexity sample rate conversion algorithm.
	KAudioConverterSampleRateConverterComplexity_Normal KAudioConverterSampleRateConverterComplexity = 'n'<<24 | 'o'<<16 | 'r'<<8 | 'm' // 'norm'
)

func (e KAudioConverterSampleRateConverterComplexity) String() string {
	switch e {
	case KAudioConverterSampleRateConverterComplexity_Linear:
		return "KAudioConverterSampleRateConverterComplexity_Linear"
	case KAudioConverterSampleRateConverterComplexity_Mastering:
		return "KAudioConverterSampleRateConverterComplexity_Mastering"
	case KAudioConverterSampleRateConverterComplexity_MinimumPhase:
		return "KAudioConverterSampleRateConverterComplexity_MinimumPhase"
	case KAudioConverterSampleRateConverterComplexity_Normal:
		return "KAudioConverterSampleRateConverterComplexity_Normal"
	default:
		return fmt.Sprintf("KAudioConverterSampleRateConverterComplexity(%d)", e)
	}
}

type KAudioDecoderComponentType uint32

const (
	// KAudioDecoderComponentTypeValue: A codec that translates data in some other format into linear PCM.
	KAudioDecoderComponentTypeValue KAudioDecoderComponentType = 'a'<<24 | 'd'<<16 | 'e'<<8 | 'c' // 'adec'
	// KAudioEncoderComponentType: A codec that translates linear PCM data into some other format
	KAudioEncoderComponentType    KAudioDecoderComponentType = 'a'<<24 | 'e'<<16 | 'n'<<8 | 'c' // 'aenc'
	KAudioUnityCodecComponentType KAudioDecoderComponentType = 'a'<<24 | 'c'<<16 | 'd'<<8 | 'c' // 'acdc'
)

func (e KAudioDecoderComponentType) String() string {
	switch e {
	case KAudioDecoderComponentTypeValue:
		return "KAudioDecoderComponentTypeValue"
	case KAudioEncoderComponentType:
		return "KAudioEncoderComponentType"
	case KAudioUnityCodecComponentType:
		return "KAudioUnityCodecComponentType"
	default:
		return fmt.Sprintf("KAudioDecoderComponentType(%d)", e)
	}
}

type KAudioFileAIFFType uint32

const (
	// KAudioFile3GP2Type: A 3GPP2 file, suitable for video content on CDMA mobile phones.
	KAudioFile3GP2Type KAudioFileAIFFType = '3'<<24 | 'g'<<16 | 'p'<<8 | '2' // '3gp2'
	// KAudioFile3GPType: A 3GPP file, suitable for video content on GSM mobile phones.
	KAudioFile3GPType KAudioFileAIFFType = '3'<<24 | 'g'<<16 | 'p'<<8 | 'p' // '3gpp'
	// KAudioFileAAC_ADTSType: An Advanced Audio Coding (AAC) Audio Data Transport Stream (ADTS) file.
	KAudioFileAAC_ADTSType KAudioFileAIFFType = 'a'<<24 | 'd'<<16 | 't'<<8 | 's' // 'adts'
	// KAudioFileAC3Type: An AC-3 file.
	KAudioFileAC3Type KAudioFileAIFFType = 'a'<<24 | 'c'<<16 | '-'<<8 | '3' // 'ac-3'
	// KAudioFileAIFCType: An Audio Interchange File Format Compressed (AIFF-C) file.
	KAudioFileAIFCType KAudioFileAIFFType = 'A'<<24 | 'I'<<16 | 'F'<<8 | 'C' // 'AIFC'
	// KAudioFileAIFFTypeValue: An Audio Interchange File Format (AIFF) file.
	KAudioFileAIFFTypeValue KAudioFileAIFFType = 'A'<<24 | 'I'<<16 | 'F'<<8 | 'F' // 'AIFF'
	// KAudioFileAMRType: An AMR (Adaptive Multi-Rate) file suitable for compressed speech.
	KAudioFileAMRType  KAudioFileAIFFType = 'a'<<24 | 'm'<<16 | 'r'<<8 | 'f' // 'amrf'
	KAudioFileBW64Type KAudioFileAIFFType = 'B'<<24 | 'W'<<16 | '6'<<8 | '4' // 'BW64'
	// KAudioFileCAFType: A Core Audio File Format file.
	KAudioFileCAFType        KAudioFileAIFFType = 'c'<<24 | 'a'<<16 | 'f'<<8 | 'f' // 'caff'
	KAudioFileFLACType       KAudioFileAIFFType = 'f'<<24 | 'l'<<16 | 'a'<<8 | 'c' // 'flac'
	KAudioFileLATMInLOASType KAudioFileAIFFType = 'l'<<24 | 'o'<<16 | 'a'<<8 | 's' // 'loas'
	// KAudioFileM4AType: An M4A file.
	KAudioFileM4AType KAudioFileAIFFType = 'm'<<24 | '4'<<16 | 'a'<<8 | 'f' // 'm4af'
	KAudioFileM4BType KAudioFileAIFFType = 'm'<<24 | '4'<<16 | 'b'<<8 | 'f' // 'm4bf'
	// KAudioFileMP1Type: An MPEG Audio Layer 1 (`XCUIElementTypeMp1`) file.
	KAudioFileMP1Type KAudioFileAIFFType = 'M'<<24 | 'P'<<16 | 'G'<<8 | '1' // 'MPG1'
	// KAudioFileMP2Type: An MPEG Audio Layer 2 (`XCUIElementTypeMp2`) file.
	KAudioFileMP2Type KAudioFileAIFFType = 'M'<<24 | 'P'<<16 | 'G'<<8 | '2' // 'MPG2'
	// KAudioFileMP3Type: An MPEG Audio Layer 3 (`XCUIElementTypeMp3`) file.
	KAudioFileMP3Type KAudioFileAIFFType = 'M'<<24 | 'P'<<16 | 'G'<<8 | '3' // 'MPG3'
	// KAudioFileMPEG4Type: An MPEG 4 file.
	KAudioFileMPEG4Type KAudioFileAIFFType = 'm'<<24 | 'p'<<16 | '4'<<8 | 'f' // 'mp4f'
	// KAudioFileNextType: A NeXT or Sun Microsystems file.
	KAudioFileNextType KAudioFileAIFFType = 'N'<<24 | 'e'<<16 | 'X'<<8 | 'T' // 'NeXT'
	KAudioFileRF64Type KAudioFileAIFFType = 'R'<<24 | 'F'<<16 | '6'<<8 | '4' // 'RF64'
	// KAudioFileSoundDesigner2Type: A Sound Designer II file.
	KAudioFileSoundDesigner2Type KAudioFileAIFFType = 'S'<<24 | 'd'<<16 | '2'<<8 | 'f' // 'Sd2f'
	// KAudioFileWAVEType: A Microsoft WAVE file.
	KAudioFileWAVEType   KAudioFileAIFFType = 'W'<<24 | 'A'<<16 | 'V'<<8 | 'E' // 'WAVE'
	KAudioFileWave64Type KAudioFileAIFFType = 'W'<<24 | '6'<<16 | '4'<<8 | 'f' // 'W64f'
)

func (e KAudioFileAIFFType) String() string {
	switch e {
	case KAudioFile3GP2Type:
		return "KAudioFile3GP2Type"
	case KAudioFile3GPType:
		return "KAudioFile3GPType"
	case KAudioFileAAC_ADTSType:
		return "KAudioFileAAC_ADTSType"
	case KAudioFileAC3Type:
		return "KAudioFileAC3Type"
	case KAudioFileAIFCType:
		return "KAudioFileAIFCType"
	case KAudioFileAIFFTypeValue:
		return "KAudioFileAIFFTypeValue"
	case KAudioFileAMRType:
		return "KAudioFileAMRType"
	case KAudioFileBW64Type:
		return "KAudioFileBW64Type"
	case KAudioFileCAFType:
		return "KAudioFileCAFType"
	case KAudioFileFLACType:
		return "KAudioFileFLACType"
	case KAudioFileLATMInLOASType:
		return "KAudioFileLATMInLOASType"
	case KAudioFileM4AType:
		return "KAudioFileM4AType"
	case KAudioFileM4BType:
		return "KAudioFileM4BType"
	case KAudioFileMP1Type:
		return "KAudioFileMP1Type"
	case KAudioFileMP2Type:
		return "KAudioFileMP2Type"
	case KAudioFileMP3Type:
		return "KAudioFileMP3Type"
	case KAudioFileMPEG4Type:
		return "KAudioFileMPEG4Type"
	case KAudioFileNextType:
		return "KAudioFileNextType"
	case KAudioFileRF64Type:
		return "KAudioFileRF64Type"
	case KAudioFileSoundDesigner2Type:
		return "KAudioFileSoundDesigner2Type"
	case KAudioFileWAVEType:
		return "KAudioFileWAVEType"
	case KAudioFileWave64Type:
		return "KAudioFileWave64Type"
	default:
		return fmt.Sprintf("KAudioFileAIFFType(%d)", e)
	}
}

type KAudioFileComponent uint32

const (
	KAudioFileComponent_AvailableFormatIDs                   KAudioFileComponent = 'f'<<24 | 'm'<<16 | 'i'<<8 | 'd' // 'fmid'
	KAudioFileComponent_AvailableStreamDescriptionsForFormat KAudioFileComponent = 's'<<24 | 'd'<<16 | 'i'<<8 | 'd' // 'sdid'
	KAudioFileComponent_CanRead                              KAudioFileComponent = 'c'<<24 | 'n'<<16 | 'r'<<8 | 'd' // 'cnrd'
	KAudioFileComponent_CanWrite                             KAudioFileComponent = 'c'<<24 | 'n'<<16 | 'w'<<8 | 'r' // 'cnwr'
	KAudioFileComponent_ExtensionsForType                    KAudioFileComponent = 'f'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'fext'
	KAudioFileComponent_FastDispatchTable                    KAudioFileComponent = 'f'<<24 | 'd'<<16 | 'f'<<8 | 't' // 'fdft'
	KAudioFileComponent_FileTypeName                         KAudioFileComponent = 'f'<<24 | 't'<<16 | 'n'<<8 | 'm' // 'ftnm'
	KAudioFileComponent_HFSTypeCodesForType                  KAudioFileComponent = 'f'<<24 | 'h'<<16 | 'f'<<8 | 's' // 'fhfs'
	KAudioFileComponent_MIMETypesForType                     KAudioFileComponent = 'f'<<24 | 'm'<<16 | 'i'<<8 | 'm' // 'fmim'
	KAudioFileComponent_UTIsForType                          KAudioFileComponent = 'f'<<24 | 'u'<<16 | 't'<<8 | 'i' // 'futi'
)

func (e KAudioFileComponent) String() string {
	switch e {
	case KAudioFileComponent_AvailableFormatIDs:
		return "KAudioFileComponent_AvailableFormatIDs"
	case KAudioFileComponent_AvailableStreamDescriptionsForFormat:
		return "KAudioFileComponent_AvailableStreamDescriptionsForFormat"
	case KAudioFileComponent_CanRead:
		return "KAudioFileComponent_CanRead"
	case KAudioFileComponent_CanWrite:
		return "KAudioFileComponent_CanWrite"
	case KAudioFileComponent_ExtensionsForType:
		return "KAudioFileComponent_ExtensionsForType"
	case KAudioFileComponent_FastDispatchTable:
		return "KAudioFileComponent_FastDispatchTable"
	case KAudioFileComponent_FileTypeName:
		return "KAudioFileComponent_FileTypeName"
	case KAudioFileComponent_HFSTypeCodesForType:
		return "KAudioFileComponent_HFSTypeCodesForType"
	case KAudioFileComponent_MIMETypesForType:
		return "KAudioFileComponent_MIMETypesForType"
	case KAudioFileComponent_UTIsForType:
		return "KAudioFileComponent_UTIsForType"
	default:
		return fmt.Sprintf("KAudioFileComponent(%d)", e)
	}
}

type KAudioFileCreateSelect uint32

const (
	KAudioFileCloseSelect                   KAudioFileCreateSelect = 0x6
	KAudioFileCountUserDataSelect           KAudioFileCreateSelect = 0x14
	KAudioFileCreateSelectValue             KAudioFileCreateSelect = 0x1
	KAudioFileCreateURLSelect               KAudioFileCreateSelect = 0x19
	KAudioFileDataIsThisFormatSelect        KAudioFileCreateSelect = 0x11
	KAudioFileExtensionIsThisFormatSelect   KAudioFileCreateSelect = 0xf
	KAudioFileFileDataIsThisFormatSelect    KAudioFileCreateSelect = 0x1b
	KAudioFileFileIsThisFormatSelect        KAudioFileCreateSelect = 0x10
	KAudioFileGetGlobalInfoSelect           KAudioFileCreateSelect = 0x13
	KAudioFileGetGlobalInfoSizeSelect       KAudioFileCreateSelect = 0x12
	KAudioFileGetPropertyInfoSelect         KAudioFileCreateSelect = 0xc
	KAudioFileGetPropertySelect             KAudioFileCreateSelect = 0xd
	KAudioFileGetUserDataAtOffsetSelect     KAudioFileCreateSelect = 0x1e
	KAudioFileGetUserDataSelect             KAudioFileCreateSelect = 0x16
	KAudioFileGetUserDataSize64Select       KAudioFileCreateSelect = 0x1d
	KAudioFileGetUserDataSizeSelect         KAudioFileCreateSelect = 0x15
	KAudioFileInitializeSelect              KAudioFileCreateSelect = 0x3
	KAudioFileInitializeWithCallbacksSelect KAudioFileCreateSelect = 0x5
	KAudioFileOpenSelect                    KAudioFileCreateSelect = 0x2
	KAudioFileOpenURLSelect                 KAudioFileCreateSelect = 0x1a
	KAudioFileOpenWithCallbacksSelect       KAudioFileCreateSelect = 0x4
	KAudioFileOptimizeSelect                KAudioFileCreateSelect = 0x7
	KAudioFileReadBytesSelect               KAudioFileCreateSelect = 0x8
	KAudioFileReadPacketDataSelect          KAudioFileCreateSelect = 0x1c
	KAudioFileReadPacketsSelect             KAudioFileCreateSelect = 0xa
	KAudioFileRemoveUserDataSelect          KAudioFileCreateSelect = 0x18
	KAudioFileSetPropertySelect             KAudioFileCreateSelect = 0xe
	KAudioFileSetUserDataSelect             KAudioFileCreateSelect = 0x17
	KAudioFileWriteBytesSelect              KAudioFileCreateSelect = 0x9
	KAudioFileWritePacketsSelect            KAudioFileCreateSelect = 0xb
)

func (e KAudioFileCreateSelect) String() string {
	switch e {
	case KAudioFileCloseSelect:
		return "KAudioFileCloseSelect"
	case KAudioFileCountUserDataSelect:
		return "KAudioFileCountUserDataSelect"
	case KAudioFileCreateSelectValue:
		return "KAudioFileCreateSelectValue"
	case KAudioFileCreateURLSelect:
		return "KAudioFileCreateURLSelect"
	case KAudioFileDataIsThisFormatSelect:
		return "KAudioFileDataIsThisFormatSelect"
	case KAudioFileExtensionIsThisFormatSelect:
		return "KAudioFileExtensionIsThisFormatSelect"
	case KAudioFileFileDataIsThisFormatSelect:
		return "KAudioFileFileDataIsThisFormatSelect"
	case KAudioFileFileIsThisFormatSelect:
		return "KAudioFileFileIsThisFormatSelect"
	case KAudioFileGetGlobalInfoSelect:
		return "KAudioFileGetGlobalInfoSelect"
	case KAudioFileGetGlobalInfoSizeSelect:
		return "KAudioFileGetGlobalInfoSizeSelect"
	case KAudioFileGetPropertyInfoSelect:
		return "KAudioFileGetPropertyInfoSelect"
	case KAudioFileGetPropertySelect:
		return "KAudioFileGetPropertySelect"
	case KAudioFileGetUserDataAtOffsetSelect:
		return "KAudioFileGetUserDataAtOffsetSelect"
	case KAudioFileGetUserDataSelect:
		return "KAudioFileGetUserDataSelect"
	case KAudioFileGetUserDataSize64Select:
		return "KAudioFileGetUserDataSize64Select"
	case KAudioFileGetUserDataSizeSelect:
		return "KAudioFileGetUserDataSizeSelect"
	case KAudioFileInitializeSelect:
		return "KAudioFileInitializeSelect"
	case KAudioFileInitializeWithCallbacksSelect:
		return "KAudioFileInitializeWithCallbacksSelect"
	case KAudioFileOpenSelect:
		return "KAudioFileOpenSelect"
	case KAudioFileOpenURLSelect:
		return "KAudioFileOpenURLSelect"
	case KAudioFileOpenWithCallbacksSelect:
		return "KAudioFileOpenWithCallbacksSelect"
	case KAudioFileOptimizeSelect:
		return "KAudioFileOptimizeSelect"
	case KAudioFileReadBytesSelect:
		return "KAudioFileReadBytesSelect"
	case KAudioFileReadPacketDataSelect:
		return "KAudioFileReadPacketDataSelect"
	case KAudioFileReadPacketsSelect:
		return "KAudioFileReadPacketsSelect"
	case KAudioFileRemoveUserDataSelect:
		return "KAudioFileRemoveUserDataSelect"
	case KAudioFileSetPropertySelect:
		return "KAudioFileSetPropertySelect"
	case KAudioFileSetUserDataSelect:
		return "KAudioFileSetUserDataSelect"
	case KAudioFileWriteBytesSelect:
		return "KAudioFileWriteBytesSelect"
	case KAudioFileWritePacketsSelect:
		return "KAudioFileWritePacketsSelect"
	default:
		return fmt.Sprintf("KAudioFileCreateSelect(%d)", e)
	}
}

type KAudioFileGlobalInfo uint32

const (
	// KAudioFileGlobalInfo_AllExtensions: A [CFArray] of [CFStrings] containing all recognized file extensions.
	KAudioFileGlobalInfo_AllExtensions KAudioFileGlobalInfo = 'a'<<24 | 'l'<<16 | 'x'<<8 | 't' // 'alxt'
	// KAudioFileGlobalInfo_AllHFSTypeCodes: An array of HFS type codes containing all recognized HFS type codes.
	KAudioFileGlobalInfo_AllHFSTypeCodes KAudioFileGlobalInfo = 'a'<<24 | 'h'<<16 | 'f'<<8 | 's' // 'ahfs'
	// KAudioFileGlobalInfo_AllMIMETypes: A [CFArray] of CF strings of all MIME types are recognized by Audio File Services.
	KAudioFileGlobalInfo_AllMIMETypes KAudioFileGlobalInfo = 'a'<<24 | 'm'<<16 | 'i'<<8 | 'm' // 'amim'
	// KAudioFileGlobalInfo_AllUTIs: A [CFArray] of [CFString] of all UTIs (Universal Type Identifiers) recognized by Audio File Services.
	KAudioFileGlobalInfo_AllUTIs KAudioFileGlobalInfo = 'a'<<24 | 'u'<<16 | 't'<<8 | 'i' // 'auti'
	// KAudioFileGlobalInfo_AvailableFormatIDs: An array of format IDs for formats that can be read.
	KAudioFileGlobalInfo_AvailableFormatIDs KAudioFileGlobalInfo = 'f'<<24 | 'm'<<16 | 'i'<<8 | 'd' // 'fmid'
	// KAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat: An array of audio stream basic description structures, which contain all the formats for a particular file type and format ID.
	KAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat KAudioFileGlobalInfo = 's'<<24 | 'd'<<16 | 'i'<<8 | 'd' // 'sdid'
	// KAudioFileGlobalInfo_ExtensionsForType: A [CFArray] of CF strings containing the recognized file extensions for a specified type.
	KAudioFileGlobalInfo_ExtensionsForType KAudioFileGlobalInfo = 'f'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'fext'
	// KAudioFileGlobalInfo_FileTypeName: The name for the file type.
	KAudioFileGlobalInfo_FileTypeName KAudioFileGlobalInfo = 'f'<<24 | 't'<<16 | 'n'<<8 | 'm' // 'ftnm'
	// KAudioFileGlobalInfo_HFSTypeCodesForType: An array of HFS type codes corresponding to a specified file type.
	KAudioFileGlobalInfo_HFSTypeCodesForType KAudioFileGlobalInfo = 'f'<<24 | 'h'<<16 | 'f'<<8 | 's' // 'fhfs'
	// KAudioFileGlobalInfo_MIMETypesForType: A [CFArray] of [CFString] of all MIME types recognized by a specified file type.
	KAudioFileGlobalInfo_MIMETypesForType KAudioFileGlobalInfo = 'f'<<24 | 'm'<<16 | 'i'<<8 | 'm' // 'fmim'
	// KAudioFileGlobalInfo_ReadableTypes: An array of [UInt32] values containing the file types (such as AIFF, WAVE, and so forth) that can be opened for reading.
	KAudioFileGlobalInfo_ReadableTypes KAudioFileGlobalInfo = 'a'<<24 | 'f'<<16 | 'r'<<8 | 'f' // 'afrf'
	// KAudioFileGlobalInfo_TypesForExtension: An array of all audio file type IDs that support a specified filename extension.
	KAudioFileGlobalInfo_TypesForExtension KAudioFileGlobalInfo = 't'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'text'
	// KAudioFileGlobalInfo_TypesForHFSTypeCode: An array of all audio file type IDs that support a specified [HFSTypeCode].
	KAudioFileGlobalInfo_TypesForHFSTypeCode KAudioFileGlobalInfo = 't'<<24 | 'h'<<16 | 'f'<<8 | 's' // 'thfs'
	// KAudioFileGlobalInfo_TypesForMIMEType: An array of all audio file type IDs that support a specified MIME type.
	KAudioFileGlobalInfo_TypesForMIMEType KAudioFileGlobalInfo = 't'<<24 | 'm'<<16 | 'i'<<8 | 'm' // 'tmim'
	// KAudioFileGlobalInfo_TypesForUTI: An array of all audio file type IDs that support a specified UTI.
	KAudioFileGlobalInfo_TypesForUTI KAudioFileGlobalInfo = 't'<<24 | 'u'<<16 | 't'<<8 | 'i' // 'tuti'
	// KAudioFileGlobalInfo_UTIsForType: A [CFArray] of [CFString] of all Universal Type Identifiers recognized by a specified file type.
	KAudioFileGlobalInfo_UTIsForType KAudioFileGlobalInfo = 'f'<<24 | 'u'<<16 | 't'<<8 | 'i' // 'futi'
	// KAudioFileGlobalInfo_WritableTypes: An array of [UInt32] values containing the file types (such as AIFF, WAVE, and so forth) that can be opened for writing.
	KAudioFileGlobalInfo_WritableTypes KAudioFileGlobalInfo = 'a'<<24 | 'f'<<16 | 'w'<<8 | 'f' // 'afwf'
)

func (e KAudioFileGlobalInfo) String() string {
	switch e {
	case KAudioFileGlobalInfo_AllExtensions:
		return "KAudioFileGlobalInfo_AllExtensions"
	case KAudioFileGlobalInfo_AllHFSTypeCodes:
		return "KAudioFileGlobalInfo_AllHFSTypeCodes"
	case KAudioFileGlobalInfo_AllMIMETypes:
		return "KAudioFileGlobalInfo_AllMIMETypes"
	case KAudioFileGlobalInfo_AllUTIs:
		return "KAudioFileGlobalInfo_AllUTIs"
	case KAudioFileGlobalInfo_AvailableFormatIDs:
		return "KAudioFileGlobalInfo_AvailableFormatIDs"
	case KAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat:
		return "KAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat"
	case KAudioFileGlobalInfo_ExtensionsForType:
		return "KAudioFileGlobalInfo_ExtensionsForType"
	case KAudioFileGlobalInfo_FileTypeName:
		return "KAudioFileGlobalInfo_FileTypeName"
	case KAudioFileGlobalInfo_HFSTypeCodesForType:
		return "KAudioFileGlobalInfo_HFSTypeCodesForType"
	case KAudioFileGlobalInfo_MIMETypesForType:
		return "KAudioFileGlobalInfo_MIMETypesForType"
	case KAudioFileGlobalInfo_ReadableTypes:
		return "KAudioFileGlobalInfo_ReadableTypes"
	case KAudioFileGlobalInfo_TypesForExtension:
		return "KAudioFileGlobalInfo_TypesForExtension"
	case KAudioFileGlobalInfo_TypesForHFSTypeCode:
		return "KAudioFileGlobalInfo_TypesForHFSTypeCode"
	case KAudioFileGlobalInfo_TypesForMIMEType:
		return "KAudioFileGlobalInfo_TypesForMIMEType"
	case KAudioFileGlobalInfo_TypesForUTI:
		return "KAudioFileGlobalInfo_TypesForUTI"
	case KAudioFileGlobalInfo_UTIsForType:
		return "KAudioFileGlobalInfo_UTIsForType"
	case KAudioFileGlobalInfo_WritableTypes:
		return "KAudioFileGlobalInfo_WritableTypes"
	default:
		return fmt.Sprintf("KAudioFileGlobalInfo(%d)", e)
	}
}

type KAudioFileLoopDirection uint32

const (
	// KAudioFileLoopDirection_Backward: Play the segment backward.
	KAudioFileLoopDirection_Backward KAudioFileLoopDirection = 3
	// KAudioFileLoopDirection_Forward: Play the segment forward.
	KAudioFileLoopDirection_Forward KAudioFileLoopDirection = 1
	// KAudioFileLoopDirection_ForwardAndBackward: Play the segment forward and backward.
	KAudioFileLoopDirection_ForwardAndBackward KAudioFileLoopDirection = 2
	// KAudioFileLoopDirection_NoLooping: The segment is not looped.
	KAudioFileLoopDirection_NoLooping KAudioFileLoopDirection = 0
)

func (e KAudioFileLoopDirection) String() string {
	switch e {
	case KAudioFileLoopDirection_Backward:
		return "KAudioFileLoopDirection_Backward"
	case KAudioFileLoopDirection_Forward:
		return "KAudioFileLoopDirection_Forward"
	case KAudioFileLoopDirection_ForwardAndBackward:
		return "KAudioFileLoopDirection_ForwardAndBackward"
	case KAudioFileLoopDirection_NoLooping:
		return "KAudioFileLoopDirection_NoLooping"
	default:
		return fmt.Sprintf("KAudioFileLoopDirection(%d)", e)
	}
}

type KAudioFileProperty uint32

const (
	// KAudioFilePropertyAlbumArtwork: An object containing the image data for the album artwork associated with an audio file.
	KAudioFilePropertyAlbumArtwork KAudioFileProperty = 'a'<<24 | 'a'<<16 | 'r'<<8 | 't' // 'aart'
	// KAudioFilePropertyAudioDataByteCount: Indicates the number of bytes of audio data in the designated file.
	KAudioFilePropertyAudioDataByteCount KAudioFileProperty = 'b'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'bcnt'
	// KAudioFilePropertyAudioDataPacketCount: Indicates the number of packets of audio data in the designated file.
	KAudioFilePropertyAudioDataPacketCount KAudioFileProperty = 'p'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'pcnt'
	KAudioFilePropertyAudioTrackCount      KAudioFileProperty = 'a'<<24 | 't'<<16 | 'c'<<8 | 't' // 'atct'
	// KAudioFilePropertyBitRate: The actual bit rate (number of audio data bits in the file divided by the duration of the file) for some file types, and the nominal bit rate (which bit rate the encoder was set to) for others.
	KAudioFilePropertyBitRate KAudioFileProperty = 'b'<<24 | 'r'<<16 | 'a'<<8 | 't' // 'brat'
	// KAudioFilePropertyByteToPacket: Passes an audio byte packet translation structure with the `mByte` field filled out and returns the `mPacket` and `mByteOffsetInPacket` fields.
	KAudioFilePropertyByteToPacket KAudioFileProperty = 'b'<<24 | 'y'<<16 | 'p'<<8 | 'k' // 'bypk'
	// KAudioFilePropertyChannelLayout: An audio channel layout structure.
	KAudioFilePropertyChannelLayout KAudioFileProperty = 'c'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'cmap'
	// KAudioFilePropertyChunkIDs: An array of  four-character codes for each kind of chunk in the file.
	KAudioFilePropertyChunkIDs KAudioFileProperty = 'c'<<24 | 'h'<<16 | 'i'<<8 | 'd' // 'chid'
	// KAudioFilePropertyDataFormat: An audio stream basic description containing the format of the audio data.
	KAudioFilePropertyDataFormat KAudioFileProperty = 'd'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'dfmt'
	// KAudioFilePropertyDataFormatName: This constant is deprecated in macOS 10.5 and later.
	KAudioFilePropertyDataFormatName KAudioFileProperty = 'f'<<24 | 'n'<<16 | 'm'<<8 | 'e' // 'fnme'
	// KAudioFilePropertyDataOffset: Indicates the byte offset in the file of the designated audio data.
	KAudioFilePropertyDataOffset KAudioFileProperty = 'd'<<24 | 'o'<<16 | 'f'<<8 | 'f' // 'doff'
	// KAudioFilePropertyDeferSizeUpdates: The default value (`0`) always updates header.
	KAudioFilePropertyDeferSizeUpdates KAudioFileProperty = 'd'<<24 | 's'<<16 | 'z'<<8 | 'u' // 'dszu'
	// KAudioFilePropertyEstimatedDuration: An estimated duration in seconds.
	KAudioFilePropertyEstimatedDuration KAudioFileProperty = 'e'<<24 | 'd'<<16 | 'u'<<8 | 'r' // 'edur'
	// KAudioFilePropertyFileFormat: The format of the audio data file.
	KAudioFilePropertyFileFormat KAudioFileProperty = 'f'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'ffmt'
	// KAudioFilePropertyFormatList: To support formats such as AAC SBR in which an encoded data stream can be decoded to multiple destination formats, this property’s value is an array of audio format list item values (declared in `AudioFormat.H()`) of those formats.
	KAudioFilePropertyFormatList KAudioFileProperty = 'f'<<24 | 'l'<<16 | 's'<<8 | 't' // 'flst'
	// KAudioFilePropertyFrameToPacket: Passes an audio frame packet translation structure with the `mFrame` field filled out and returns the `mPacket` and `mFrameOffsetInPacket` fields.
	KAudioFilePropertyFrameToPacket KAudioFileProperty = 'f'<<24 | 'r'<<16 | 'p'<<8 | 'k' // 'frpk'
	// KAudioFilePropertyID3Tag: A `void*` value pointing to memory set up by your application to contain a fully formatted ID3 tag.
	KAudioFilePropertyID3Tag       KAudioFileProperty = 'i'<<24 | 'd'<<16 | '3'<<8 | 't' // 'id3t'
	KAudioFilePropertyID3TagOffset KAudioFileProperty = 'i'<<24 | 'd'<<16 | '3'<<8 | 'o' // 'id3o'
	// KAudioFilePropertyInfoDictionary: A CF Dictionary with information about the data in the file.
	KAudioFilePropertyInfoDictionary KAudioFileProperty = 'i'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'info'
	// KAudioFilePropertyIsOptimized: Indicates whether a designated audio file has been optimized, that is, ready to start having sound data written to it.
	KAudioFilePropertyIsOptimized KAudioFileProperty = 'o'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'optm'
	// KAudioFilePropertyMagicCookieData: A pointer to memory set up by the caller.
	KAudioFilePropertyMagicCookieData KAudioFileProperty = 'm'<<24 | 'g'<<16 | 'i'<<8 | 'c' // 'mgic'
	// KAudioFilePropertyMarkerList: A list of audio file markers defined in the file.
	KAudioFilePropertyMarkerList KAudioFileProperty = 'm'<<24 | 'k'<<16 | 'l'<<8 | 's' // 'mkls'
	// KAudioFilePropertyMaximumPacketSize: Indicates the maximum size of a packet for the data in the designated file.
	KAudioFilePropertyMaximumPacketSize              KAudioFileProperty = 'p'<<24 | 's'<<16 | 'z'<<8 | 'e' // 'psze'
	KAudioFilePropertyNextIndependentPacket          KAudioFileProperty = 'n'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'nind'
	KAudioFilePropertyPacketRangeByteCountUpperBound KAudioFileProperty = 'p'<<24 | 'r'<<16 | 'u'<<8 | 'b' // 'prub'
	// KAudioFilePropertyPacketSizeUpperBound: The theoretical maximum packet size in the file.
	KAudioFilePropertyPacketSizeUpperBound KAudioFileProperty = 'p'<<24 | 'k'<<16 | 'u'<<8 | 'b' // 'pkub'
	// KAudioFilePropertyPacketTableInfo: Gets or sets an audio file packet table information structure for its supporting file types.
	KAudioFilePropertyPacketTableInfo KAudioFileProperty = 'p'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'pnfo'
	// KAudioFilePropertyPacketToByte: Passes an audio byte packet translation structure with the `mPacket` field filled out and returns the `mByte` field.
	KAudioFilePropertyPacketToByte           KAudioFileProperty = 'p'<<24 | 'k'<<16 | 'b'<<8 | 'y' // 'pkby'
	KAudioFilePropertyPacketToDependencyInfo KAudioFileProperty = 'p'<<24 | 'k'<<16 | 'd'<<8 | 'p' // 'pkdp'
	// KAudioFilePropertyPacketToFrame: Passes an audio frame packet translation structure with the `mPacket` field filled out and returns the `mFrame` field.
	KAudioFilePropertyPacketToFrame             KAudioFileProperty = 'p'<<24 | 'k'<<16 | 'f'<<8 | 'r' // 'pkfr'
	KAudioFilePropertyPacketToRollDistance      KAudioFileProperty = 'p'<<24 | 'k'<<16 | 'r'<<8 | 'l' // 'pkrl'
	KAudioFilePropertyPreviousIndependentPacket KAudioFileProperty = 'p'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'pind'
	// KAudioFilePropertyRegionList: The list of audio file region values defined in the file.
	KAudioFilePropertyRegionList KAudioFileProperty = 'r'<<24 | 'g'<<16 | 'l'<<8 | 's' // 'rgls'
	// KAudioFilePropertyReserveDuration: The duration in seconds of the data expected to be written.
	KAudioFilePropertyReserveDuration       KAudioFileProperty = 'r'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'rsrv'
	KAudioFilePropertyRestrictsRandomAccess KAudioFileProperty = 'r'<<24 | 'r'<<16 | 'a'<<8 | 'p' // 'rrap'
	// KAudioFilePropertySourceBitDepth: For compressed data, this property’s value is the bit depth of the source, uncompressed audio stream as an [SInt32] value, if known.
	KAudioFilePropertySourceBitDepth KAudioFileProperty = 's'<<24 | 'b'<<16 | 't'<<8 | 'd' // 'sbtd'
	KAudioFilePropertyUseAudioTrack  KAudioFileProperty = 'u'<<24 | 'a'<<16 | 't'<<8 | 'k' // 'uatk'
)

func (e KAudioFileProperty) String() string {
	switch e {
	case KAudioFilePropertyAlbumArtwork:
		return "KAudioFilePropertyAlbumArtwork"
	case KAudioFilePropertyAudioDataByteCount:
		return "KAudioFilePropertyAudioDataByteCount"
	case KAudioFilePropertyAudioDataPacketCount:
		return "KAudioFilePropertyAudioDataPacketCount"
	case KAudioFilePropertyAudioTrackCount:
		return "KAudioFilePropertyAudioTrackCount"
	case KAudioFilePropertyBitRate:
		return "KAudioFilePropertyBitRate"
	case KAudioFilePropertyByteToPacket:
		return "KAudioFilePropertyByteToPacket"
	case KAudioFilePropertyChannelLayout:
		return "KAudioFilePropertyChannelLayout"
	case KAudioFilePropertyChunkIDs:
		return "KAudioFilePropertyChunkIDs"
	case KAudioFilePropertyDataFormat:
		return "KAudioFilePropertyDataFormat"
	case KAudioFilePropertyDataFormatName:
		return "KAudioFilePropertyDataFormatName"
	case KAudioFilePropertyDataOffset:
		return "KAudioFilePropertyDataOffset"
	case KAudioFilePropertyDeferSizeUpdates:
		return "KAudioFilePropertyDeferSizeUpdates"
	case KAudioFilePropertyEstimatedDuration:
		return "KAudioFilePropertyEstimatedDuration"
	case KAudioFilePropertyFileFormat:
		return "KAudioFilePropertyFileFormat"
	case KAudioFilePropertyFormatList:
		return "KAudioFilePropertyFormatList"
	case KAudioFilePropertyFrameToPacket:
		return "KAudioFilePropertyFrameToPacket"
	case KAudioFilePropertyID3Tag:
		return "KAudioFilePropertyID3Tag"
	case KAudioFilePropertyID3TagOffset:
		return "KAudioFilePropertyID3TagOffset"
	case KAudioFilePropertyInfoDictionary:
		return "KAudioFilePropertyInfoDictionary"
	case KAudioFilePropertyIsOptimized:
		return "KAudioFilePropertyIsOptimized"
	case KAudioFilePropertyMagicCookieData:
		return "KAudioFilePropertyMagicCookieData"
	case KAudioFilePropertyMarkerList:
		return "KAudioFilePropertyMarkerList"
	case KAudioFilePropertyMaximumPacketSize:
		return "KAudioFilePropertyMaximumPacketSize"
	case KAudioFilePropertyNextIndependentPacket:
		return "KAudioFilePropertyNextIndependentPacket"
	case KAudioFilePropertyPacketRangeByteCountUpperBound:
		return "KAudioFilePropertyPacketRangeByteCountUpperBound"
	case KAudioFilePropertyPacketSizeUpperBound:
		return "KAudioFilePropertyPacketSizeUpperBound"
	case KAudioFilePropertyPacketTableInfo:
		return "KAudioFilePropertyPacketTableInfo"
	case KAudioFilePropertyPacketToByte:
		return "KAudioFilePropertyPacketToByte"
	case KAudioFilePropertyPacketToDependencyInfo:
		return "KAudioFilePropertyPacketToDependencyInfo"
	case KAudioFilePropertyPacketToFrame:
		return "KAudioFilePropertyPacketToFrame"
	case KAudioFilePropertyPacketToRollDistance:
		return "KAudioFilePropertyPacketToRollDistance"
	case KAudioFilePropertyPreviousIndependentPacket:
		return "KAudioFilePropertyPreviousIndependentPacket"
	case KAudioFilePropertyRegionList:
		return "KAudioFilePropertyRegionList"
	case KAudioFilePropertyReserveDuration:
		return "KAudioFilePropertyReserveDuration"
	case KAudioFilePropertyRestrictsRandomAccess:
		return "KAudioFilePropertyRestrictsRandomAccess"
	case KAudioFilePropertySourceBitDepth:
		return "KAudioFilePropertySourceBitDepth"
	case KAudioFilePropertyUseAudioTrack:
		return "KAudioFilePropertyUseAudioTrack"
	default:
		return fmt.Sprintf("KAudioFileProperty(%d)", e)
	}
}

type KAudioFileStreamError int32

const (
	// KAudioFileStreamError_BadPropertySize: The size of the buffer you provided for property data was not correct.
	KAudioFileStreamError_BadPropertySize KAudioFileStreamError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	// KAudioFileStreamError_DataUnavailable: The amount of data provided to the parser was insufficient to produce any result.
	KAudioFileStreamError_DataUnavailable KAudioFileStreamError = 'm'<<24 | 'o'<<16 | 'r'<<8 | 'e' // 'more'
	// KAudioFileStreamError_DiscontinuityCantRecover: A discontinuity has occurred in the audio data, and Audio File Stream Services cannot recover.
	KAudioFileStreamError_DiscontinuityCantRecover KAudioFileStreamError = 'd'<<24 | 's'<<16 | 'c'<<8 | '!' // 'dsc!'
	// KAudioFileStreamError_IllegalOperation: An illegal operation was attempted.
	KAudioFileStreamError_IllegalOperation KAudioFileStreamError = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	// KAudioFileStreamError_InvalidFile: The file is malformed, not a valid instance of an audio file of its type, or not recognized as an audio file.
	KAudioFileStreamError_InvalidFile KAudioFileStreamError = 'd'<<24 | 't'<<16 | 'a'<<8 | '?' // 'dta?'
	// KAudioFileStreamError_InvalidPacketOffset: A packet offset was less than `0`, or past the end of the file, or a corrupt packet size was read when building the packet table.
	KAudioFileStreamError_InvalidPacketOffset KAudioFileStreamError = 'p'<<24 | 'c'<<16 | 'k'<<8 | '?' // 'pck?'
	// KAudioFileStreamError_NotOptimized: It is not possible to produce output packets because the streamed audio file’s packet table or other defining information is not present or appears after the audio data.
	KAudioFileStreamError_NotOptimized KAudioFileStreamError = 'o'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'optm'
	// KAudioFileStreamError_UnspecifiedError: An unspecified error has occurred.
	KAudioFileStreamError_UnspecifiedError KAudioFileStreamError = 'w'<<24 | 'h'<<16 | 't'<<8 | '?' // 'wht?'
	// KAudioFileStreamError_UnsupportedDataFormat: The data format is not supported by the specified file type.
	KAudioFileStreamError_UnsupportedDataFormat KAudioFileStreamError = 'f'<<24 | 'm'<<16 | 't'<<8 | '?' // 'fmt?'
	// KAudioFileStreamError_UnsupportedFileType: The specified file type is not supported.
	KAudioFileStreamError_UnsupportedFileType KAudioFileStreamError = 't'<<24 | 'y'<<16 | 'p'<<8 | '?' // 'typ?'
	// KAudioFileStreamError_UnsupportedProperty: The property is not supported.
	KAudioFileStreamError_UnsupportedProperty KAudioFileStreamError = 'p'<<24 | 't'<<16 | 'y'<<8 | '?' // 'pty?'
	// KAudioFileStreamError_ValueUnknown: The property value is not present in this file before the audio data.
	KAudioFileStreamError_ValueUnknown KAudioFileStreamError = 'u'<<24 | 'n'<<16 | 'k'<<8 | '?' // 'unk?'
)

func (e KAudioFileStreamError) String() string {
	switch e {
	case KAudioFileStreamError_BadPropertySize:
		return "KAudioFileStreamError_BadPropertySize"
	case KAudioFileStreamError_DataUnavailable:
		return "KAudioFileStreamError_DataUnavailable"
	case KAudioFileStreamError_DiscontinuityCantRecover:
		return "KAudioFileStreamError_DiscontinuityCantRecover"
	case KAudioFileStreamError_IllegalOperation:
		return "KAudioFileStreamError_IllegalOperation"
	case KAudioFileStreamError_InvalidFile:
		return "KAudioFileStreamError_InvalidFile"
	case KAudioFileStreamError_InvalidPacketOffset:
		return "KAudioFileStreamError_InvalidPacketOffset"
	case KAudioFileStreamError_NotOptimized:
		return "KAudioFileStreamError_NotOptimized"
	case KAudioFileStreamError_UnspecifiedError:
		return "KAudioFileStreamError_UnspecifiedError"
	case KAudioFileStreamError_UnsupportedDataFormat:
		return "KAudioFileStreamError_UnsupportedDataFormat"
	case KAudioFileStreamError_UnsupportedFileType:
		return "KAudioFileStreamError_UnsupportedFileType"
	case KAudioFileStreamError_UnsupportedProperty:
		return "KAudioFileStreamError_UnsupportedProperty"
	case KAudioFileStreamError_ValueUnknown:
		return "KAudioFileStreamError_ValueUnknown"
	default:
		return fmt.Sprintf("KAudioFileStreamError(%d)", e)
	}
}

type KAudioFileStreamProperty uint32

const (
	// KAudioFileStreamProperty_AudioDataByteCount: A [UInt64] value indicating the number of bytes of audio data in the streamed file.
	KAudioFileStreamProperty_AudioDataByteCount KAudioFileStreamProperty = 'b'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'bcnt'
	// KAudioFileStreamProperty_AudioDataPacketCount: A [UInt64] value indicating the number of packets of audio data in the streamed file.
	KAudioFileStreamProperty_AudioDataPacketCount KAudioFileStreamProperty = 'p'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'pcnt'
	// KAudioFileStreamProperty_AverageBytesPerPacket: A [Float64] value indicating the average bytes per packet.
	KAudioFileStreamProperty_AverageBytesPerPacket KAudioFileStreamProperty = 'a'<<24 | 'b'<<16 | 'p'<<8 | 'p' // 'abpp'
	// KAudioFileStreamProperty_BitRate: A [UInt32] value indicating the bit rate of a stream in bits per second.
	KAudioFileStreamProperty_BitRate KAudioFileStreamProperty = 'b'<<24 | 'r'<<16 | 'a'<<8 | 't' // 'brat'
	// KAudioFileStreamProperty_ByteToPacket: Obtains the packet number corresponding to a byte number.
	KAudioFileStreamProperty_ByteToPacket KAudioFileStreamProperty = 'b'<<24 | 'y'<<16 | 'p'<<8 | 'k' // 'bypk'
	// KAudioFileStreamProperty_ChannelLayout: An [AudioChannelLayout] structure.
	KAudioFileStreamProperty_ChannelLayout KAudioFileStreamProperty = 'c'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'cmap'
	// KAudioFileStreamProperty_DataFormat: An [AudioStreamBasicDescription] structure describing the format of the audio data in the stream.
	KAudioFileStreamProperty_DataFormat KAudioFileStreamProperty = 'd'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'dfmt'
	// KAudioFileStreamProperty_DataOffset: An [SInt64] value indicating the byte offset in the streamed file at which the audio data starts.
	KAudioFileStreamProperty_DataOffset KAudioFileStreamProperty = 'd'<<24 | 'o'<<16 | 'f'<<8 | 'f' // 'doff'
	// KAudioFileStreamProperty_FileFormat: A four-character code that identifies the audio data format.
	KAudioFileStreamProperty_FileFormat KAudioFileStreamProperty = 'f'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'ffmt'
	// KAudioFileStreamProperty_FormatList: To support formats such as AAC with SBR where an encoded data stream can be decoded to multiple destination formats, this property returns an array of [AudioFormatListItem] structures (declared in `AudioFormat.H()`)—one for each of the destination formats.
	KAudioFileStreamProperty_FormatList KAudioFileStreamProperty = 'f'<<24 | 'l'<<16 | 's'<<8 | 't' // 'flst'
	// KAudioFileStreamProperty_FrameToPacket: Obtains the packet number corresponding to a frame number.
	KAudioFileStreamProperty_FrameToPacket  KAudioFileStreamProperty = 'f'<<24 | 'r'<<16 | 'p'<<8 | 'k' // 'frpk'
	KAudioFileStreamProperty_InfoDictionary KAudioFileStreamProperty = 'i'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'info'
	// KAudioFileStreamProperty_MagicCookieData: A pointer (`void *`) to a magic cookie.
	KAudioFileStreamProperty_MagicCookieData KAudioFileStreamProperty = 'm'<<24 | 'g'<<16 | 'i'<<8 | 'c' // 'mgic'
	// KAudioFileStreamProperty_MaximumPacketSize: A [UInt32] value indicating the maximum packet size of the data in the streamed file.
	KAudioFileStreamProperty_MaximumPacketSize     KAudioFileStreamProperty = 'p'<<24 | 's'<<16 | 'z'<<8 | 'e' // 'psze'
	KAudioFileStreamProperty_NextIndependentPacket KAudioFileStreamProperty = 'n'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'nind'
	// KAudioFileStreamProperty_PacketSizeUpperBound: A [UInt32] value indicating the theoretical maximum packet size in the streamed file.
	KAudioFileStreamProperty_PacketSizeUpperBound KAudioFileStreamProperty = 'p'<<24 | 'k'<<16 | 'u'<<8 | 'b' // 'pkub'
	// KAudioFileStreamProperty_PacketTableInfo: An [AudioFilePacketTableInfo] structure.
	KAudioFileStreamProperty_PacketTableInfo KAudioFileStreamProperty = 'p'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'pnfo'
	// KAudioFileStreamProperty_PacketToByte: Obtains the byte number corresponding to a packet number.
	KAudioFileStreamProperty_PacketToByte           KAudioFileStreamProperty = 'p'<<24 | 'k'<<16 | 'b'<<8 | 'y' // 'pkby'
	KAudioFileStreamProperty_PacketToDependencyInfo KAudioFileStreamProperty = 'p'<<24 | 'k'<<16 | 'd'<<8 | 'p' // 'pkdp'
	// KAudioFileStreamProperty_PacketToFrame: Obtains the frame number corresponding to a packet number.
	KAudioFileStreamProperty_PacketToFrame             KAudioFileStreamProperty = 'p'<<24 | 'k'<<16 | 'f'<<8 | 'r' // 'pkfr'
	KAudioFileStreamProperty_PacketToRollDistance      KAudioFileStreamProperty = 'p'<<24 | 'k'<<16 | 'r'<<8 | 'l' // 'pkrl'
	KAudioFileStreamProperty_PreviousIndependentPacket KAudioFileStreamProperty = 'p'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'pind'
	// KAudioFileStreamProperty_ReadyToProducePackets: # Discussion
	KAudioFileStreamProperty_ReadyToProducePackets KAudioFileStreamProperty = 'r'<<24 | 'e'<<16 | 'd'<<8 | 'y' // 'redy'
	KAudioFileStreamProperty_RestrictsRandomAccess KAudioFileStreamProperty = 'r'<<24 | 'r'<<16 | 'a'<<8 | 'p' // 'rrap'
)

func (e KAudioFileStreamProperty) String() string {
	switch e {
	case KAudioFileStreamProperty_AudioDataByteCount:
		return "KAudioFileStreamProperty_AudioDataByteCount"
	case KAudioFileStreamProperty_AudioDataPacketCount:
		return "KAudioFileStreamProperty_AudioDataPacketCount"
	case KAudioFileStreamProperty_AverageBytesPerPacket:
		return "KAudioFileStreamProperty_AverageBytesPerPacket"
	case KAudioFileStreamProperty_BitRate:
		return "KAudioFileStreamProperty_BitRate"
	case KAudioFileStreamProperty_ByteToPacket:
		return "KAudioFileStreamProperty_ByteToPacket"
	case KAudioFileStreamProperty_ChannelLayout:
		return "KAudioFileStreamProperty_ChannelLayout"
	case KAudioFileStreamProperty_DataFormat:
		return "KAudioFileStreamProperty_DataFormat"
	case KAudioFileStreamProperty_DataOffset:
		return "KAudioFileStreamProperty_DataOffset"
	case KAudioFileStreamProperty_FileFormat:
		return "KAudioFileStreamProperty_FileFormat"
	case KAudioFileStreamProperty_FormatList:
		return "KAudioFileStreamProperty_FormatList"
	case KAudioFileStreamProperty_FrameToPacket:
		return "KAudioFileStreamProperty_FrameToPacket"
	case KAudioFileStreamProperty_InfoDictionary:
		return "KAudioFileStreamProperty_InfoDictionary"
	case KAudioFileStreamProperty_MagicCookieData:
		return "KAudioFileStreamProperty_MagicCookieData"
	case KAudioFileStreamProperty_MaximumPacketSize:
		return "KAudioFileStreamProperty_MaximumPacketSize"
	case KAudioFileStreamProperty_NextIndependentPacket:
		return "KAudioFileStreamProperty_NextIndependentPacket"
	case KAudioFileStreamProperty_PacketSizeUpperBound:
		return "KAudioFileStreamProperty_PacketSizeUpperBound"
	case KAudioFileStreamProperty_PacketTableInfo:
		return "KAudioFileStreamProperty_PacketTableInfo"
	case KAudioFileStreamProperty_PacketToByte:
		return "KAudioFileStreamProperty_PacketToByte"
	case KAudioFileStreamProperty_PacketToDependencyInfo:
		return "KAudioFileStreamProperty_PacketToDependencyInfo"
	case KAudioFileStreamProperty_PacketToFrame:
		return "KAudioFileStreamProperty_PacketToFrame"
	case KAudioFileStreamProperty_PacketToRollDistance:
		return "KAudioFileStreamProperty_PacketToRollDistance"
	case KAudioFileStreamProperty_PreviousIndependentPacket:
		return "KAudioFileStreamProperty_PreviousIndependentPacket"
	case KAudioFileStreamProperty_ReadyToProducePackets:
		return "KAudioFileStreamProperty_ReadyToProducePackets"
	case KAudioFileStreamProperty_RestrictsRandomAccess:
		return "KAudioFileStreamProperty_RestrictsRandomAccess"
	default:
		return fmt.Sprintf("KAudioFileStreamProperty(%d)", e)
	}
}

type KAudioFileUnspecifiedError int32

const (
	// KAudioFileBadPropertySizeError: The size of the property data was not correct.
	KAudioFileBadPropertySizeError KAudioFileUnspecifiedError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	// KAudioFileDoesNotAllow64BitDataSizeError: The file offset was too large for the file type.
	KAudioFileDoesNotAllow64BitDataSizeError KAudioFileUnspecifiedError = 'o'<<24 | 'f'<<16 | 'f'<<8 | '?' // 'off?'
	// KAudioFileEndOfFileError: End of file.
	KAudioFileEndOfFileError KAudioFileUnspecifiedError = -39
	// KAudioFileFileNotFoundError: File not found.
	KAudioFileFileNotFoundError KAudioFileUnspecifiedError = -43
	// KAudioFileInvalidChunkError: Either the chunk does not exist in the file or it is not supported by the file.
	KAudioFileInvalidChunkError KAudioFileUnspecifiedError = 'c'<<24 | 'h'<<16 | 'k'<<8 | '?' // 'chk?'
	// KAudioFileInvalidFileError: The file is malformed, or otherwise not a valid instance of an audio file of its type.
	KAudioFileInvalidFileError             KAudioFileUnspecifiedError = 'd'<<24 | 't'<<16 | 'a'<<8 | '?' // 'dta?'
	KAudioFileInvalidPacketDependencyError KAudioFileUnspecifiedError = 'd'<<24 | 'e'<<16 | 'p'<<8 | '?' // 'dep?'
	// KAudioFileInvalidPacketOffsetError: A packet offset was past the end of the file, or not at the end of the file when a VBR format was written,  or a corrupt packet size was read when the packet table was built.
	KAudioFileInvalidPacketOffsetError KAudioFileUnspecifiedError = 'p'<<24 | 'c'<<16 | 'k'<<8 | '?' // 'pck?'
	// KAudioFileNotOpenError: The file is closed.
	KAudioFileNotOpenError KAudioFileUnspecifiedError = -38
	// KAudioFileNotOptimizedError: The chunks following the audio data chunk are preventing the extension of the audio data chunk.
	KAudioFileNotOptimizedError KAudioFileUnspecifiedError = 'o'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'optm'
	// KAudioFileOperationNotSupportedError: The operation cannot be performed.
	KAudioFileOperationNotSupportedError KAudioFileUnspecifiedError = 0x6f703f3f
	// KAudioFilePermissionsError: The operation violated the file permissions.
	KAudioFilePermissionsError KAudioFileUnspecifiedError = 'p'<<24 | 'r'<<16 | 'm'<<8 | '?' // 'prm?'
	// KAudioFilePositionError: Invalid file position.
	KAudioFilePositionError KAudioFileUnspecifiedError = -40
	// KAudioFileUnspecifiedErrorValue: An unspecified error has occurred.
	KAudioFileUnspecifiedErrorValue KAudioFileUnspecifiedError = 'w'<<24 | 'h'<<16 | 't'<<8 | '?' // 'wht?'
	// KAudioFileUnsupportedDataFormatError: The data format is not supported by this file type.
	KAudioFileUnsupportedDataFormatError KAudioFileUnspecifiedError = 'f'<<24 | 'm'<<16 | 't'<<8 | '?' // 'fmt?'
	// KAudioFileUnsupportedFileTypeError: The file type is not supported.
	KAudioFileUnsupportedFileTypeError KAudioFileUnspecifiedError = 't'<<24 | 'y'<<16 | 'p'<<8 | '?' // 'typ?'
	// KAudioFileUnsupportedPropertyError: The property is not supported.
	KAudioFileUnsupportedPropertyError KAudioFileUnspecifiedError = 'p'<<24 | 't'<<16 | 'y'<<8 | '?' // 'pty?'
)

func (e KAudioFileUnspecifiedError) String() string {
	switch e {
	case KAudioFileBadPropertySizeError:
		return "KAudioFileBadPropertySizeError"
	case KAudioFileDoesNotAllow64BitDataSizeError:
		return "KAudioFileDoesNotAllow64BitDataSizeError"
	case KAudioFileEndOfFileError:
		return "KAudioFileEndOfFileError"
	case KAudioFileFileNotFoundError:
		return "KAudioFileFileNotFoundError"
	case KAudioFileInvalidChunkError:
		return "KAudioFileInvalidChunkError"
	case KAudioFileInvalidFileError:
		return "KAudioFileInvalidFileError"
	case KAudioFileInvalidPacketDependencyError:
		return "KAudioFileInvalidPacketDependencyError"
	case KAudioFileInvalidPacketOffsetError:
		return "KAudioFileInvalidPacketOffsetError"
	case KAudioFileNotOpenError:
		return "KAudioFileNotOpenError"
	case KAudioFileNotOptimizedError:
		return "KAudioFileNotOptimizedError"
	case KAudioFileOperationNotSupportedError:
		return "KAudioFileOperationNotSupportedError"
	case KAudioFilePermissionsError:
		return "KAudioFilePermissionsError"
	case KAudioFilePositionError:
		return "KAudioFilePositionError"
	case KAudioFileUnspecifiedErrorValue:
		return "KAudioFileUnspecifiedErrorValue"
	case KAudioFileUnsupportedDataFormatError:
		return "KAudioFileUnsupportedDataFormatError"
	case KAudioFileUnsupportedFileTypeError:
		return "KAudioFileUnsupportedFileTypeError"
	case KAudioFileUnsupportedPropertyError:
		return "KAudioFileUnsupportedPropertyError"
	default:
		return fmt.Sprintf("KAudioFileUnspecifiedError(%d)", e)
	}
}

type KAudioFormat int32

const (
	KAudioFormatBadPropertySizeError  KAudioFormat = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KAudioFormatBadSpecifierSizeError KAudioFormat = '!'<<24 | 's'<<16 | 'p'<<8 | 'c' // '!spc'
	// KAudioFormatUnknownFormatError: The specified data format is not a known format.
	KAudioFormatUnknownFormatError KAudioFormat = '!'<<24 | 'f'<<16 | 'm'<<8 | 't' // '!fmt'
	// KAudioFormatUnspecifiedError: An unspecified error.
	KAudioFormatUnspecifiedError KAudioFormat = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	// KAudioFormatUnsupportedDataFormatError: The playback data format is unsupported (declared in `AudioFormat.H()`).
	KAudioFormatUnsupportedDataFormatError KAudioFormat = 'f'<<24 | 'm'<<16 | 't'<<8 | '?' // 'fmt?'
	// KAudioFormatUnsupportedPropertyError: The specified property is not supported.
	KAudioFormatUnsupportedPropertyError KAudioFormat = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
)

func (e KAudioFormat) String() string {
	switch e {
	case KAudioFormatBadPropertySizeError:
		return "KAudioFormatBadPropertySizeError"
	case KAudioFormatBadSpecifierSizeError:
		return "KAudioFormatBadSpecifierSizeError"
	case KAudioFormatUnknownFormatError:
		return "KAudioFormatUnknownFormatError"
	case KAudioFormatUnspecifiedError:
		return "KAudioFormatUnspecifiedError"
	case KAudioFormatUnsupportedDataFormatError:
		return "KAudioFormatUnsupportedDataFormatError"
	case KAudioFormatUnsupportedPropertyError:
		return "KAudioFormatUnsupportedPropertyError"
	default:
		return fmt.Sprintf("KAudioFormat(%d)", e)
	}
}

type KAudioFormatProperty uint32

const (
	// KAudioFormatProperty_ASBDFromESDS: An [AudioStreamBasicDescription] structure for a given elementary stream descriptor (ESDS).
	KAudioFormatProperty_ASBDFromESDS KAudioFormatProperty = 'e'<<24 | 's'<<16 | 's'<<8 | 'd' // 'essd'
	// KAudioFormatProperty_ASBDFromMPEGPacket: An [AudioStreamBasicDescription] structure for a given MPEG Packet.
	KAudioFormatProperty_ASBDFromMPEGPacket            KAudioFormatProperty = 'a'<<24 | 'd'<<16 | 'm'<<8 | 'p' // 'admp'
	KAudioFormatProperty_AreChannelLayoutsEquivalent   KAudioFormatProperty = 'c'<<24 | 'h'<<16 | 'e'<<8 | 'q' // 'cheq'
	KAudioFormatProperty_AvailableDecodeNumberChannels KAudioFormatProperty = 'a'<<24 | 'd'<<16 | 'n'<<8 | 'c' // 'adnc'
	// KAudioFormatProperty_AvailableEncodeBitRates: An array of [AudioValueRange] structures describing all available bit rates.
	KAudioFormatProperty_AvailableEncodeBitRates KAudioFormatProperty = 'a'<<24 | 'e'<<16 | 'b'<<8 | 'r' // 'aebr'
	// KAudioFormatProperty_AvailableEncodeChannelLayoutTags: An array of [AudioChannelLayoutTag] values for the format and number of channels specified.
	KAudioFormatProperty_AvailableEncodeChannelLayoutTags KAudioFormatProperty = 'a'<<24 | 'e'<<16 | 'c'<<8 | 'l' // 'aecl'
	// KAudioFormatProperty_AvailableEncodeNumberChannels: An array of [UInt32] values indicating the number of channels that can be encoded.
	KAudioFormatProperty_AvailableEncodeNumberChannels KAudioFormatProperty = 'a'<<24 | 'v'<<16 | 'n'<<8 | 'c' // 'avnc'
	// KAudioFormatProperty_AvailableEncodeSampleRates: An array of [AudioValueRange] structures.
	KAudioFormatProperty_AvailableEncodeSampleRates KAudioFormatProperty = 'a'<<24 | 'e'<<16 | 's'<<8 | 'r' // 'aesr'
	// KAudioFormatProperty_BalanceFade: An array of coefficients, each a [Float32] value, for applying left/right audio balance and front/back audio fade.
	KAudioFormatProperty_BalanceFade KAudioFormatProperty = 'b'<<24 | 'a'<<16 | 'l'<<8 | 'f' // 'balf'
	// KAudioFormatProperty_BitmapForLayoutTag: A bitmap for an [AudioChannelLayoutTag] value.
	KAudioFormatProperty_BitmapForLayoutTag KAudioFormatProperty = 'b'<<24 | 'm'<<16 | 't'<<8 | 'g' // 'bmtg'
	// KAudioFormatProperty_ChannelLayoutForBitmap: The channel descriptions for a standard channel layout,
	KAudioFormatProperty_ChannelLayoutForBitmap KAudioFormatProperty = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'b' // 'cmpb'
	// KAudioFormatProperty_ChannelLayoutForTag: The channel descriptions for a standard channel layout.
	KAudioFormatProperty_ChannelLayoutForTag KAudioFormatProperty = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'l' // 'cmpl'
	// KAudioFormatProperty_ChannelLayoutFromESDS: An [AudioChannelLayout] structure for a given elementary stream descriptor (ESDS).
	KAudioFormatProperty_ChannelLayoutFromESDS KAudioFormatProperty = 'e'<<24 | 's'<<16 | 'c'<<8 | 'l' // 'escl'
	KAudioFormatProperty_ChannelLayoutHash     KAudioFormatProperty = 'c'<<24 | 'h'<<16 | 'h'<<8 | 'a' // 'chha'
	// KAudioFormatProperty_ChannelLayoutName: The a name for a particular channel layout.
	KAudioFormatProperty_ChannelLayoutName KAudioFormatProperty = 'l'<<24 | 'o'<<16 | 'n'<<8 | 'm' // 'lonm'
	// KAudioFormatProperty_ChannelLayoutSimpleName: A simplified name for channel layout.
	KAudioFormatProperty_ChannelLayoutSimpleName KAudioFormatProperty = 'l'<<24 | 's'<<16 | 'n'<<8 | 'm' // 'lsnm'
	// KAudioFormatProperty_ChannelMap: An array of [SInt32] values for reordering input channels.
	KAudioFormatProperty_ChannelMap KAudioFormatProperty = 'c'<<24 | 'h'<<16 | 'm'<<8 | 'p' // 'chmp'
	// KAudioFormatProperty_ChannelName: The name for a particular channel.
	KAudioFormatProperty_ChannelName KAudioFormatProperty = 'c'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'cnam'
	// KAudioFormatProperty_ChannelShortName: An abbreviated name for a particular channel.
	KAudioFormatProperty_ChannelShortName KAudioFormatProperty = 'c'<<24 | 's'<<16 | 'n'<<8 | 'm' // 'csnm'
	// KAudioFormatProperty_DecodeFormatIDs: An array of [UInt32] values representing format identifiers for formats that are valid input formats for a converter.
	KAudioFormatProperty_DecodeFormatIDs KAudioFormatProperty = 'a'<<24 | 'c'<<16 | 'i'<<8 | 'f' // 'acif'
	// KAudioFormatProperty_Decoders: An array of [AudioClassDescription] structures for all installed decoders for the specified audio format.
	KAudioFormatProperty_Decoders KAudioFormatProperty = 'a'<<24 | 'v'<<16 | 'd'<<8 | 'e' // 'avde'
	// KAudioFormatProperty_EncodeFormatIDs: An array of [UInt32] values representing format identifiers for formats that are valid output formats for a converter.
	KAudioFormatProperty_EncodeFormatIDs KAudioFormatProperty = 'a'<<24 | 'c'<<16 | 'o'<<8 | 'f' // 'acof'
	// KAudioFormatProperty_Encoders: An array of [AudioClassDescription] structures for all installed encoders for the specified audio format.
	KAudioFormatProperty_Encoders KAudioFormatProperty = 'a'<<24 | 'v'<<16 | 'e'<<8 | 'n' // 'aven'
	// KAudioFormatProperty_FirstPlayableFormatFromList: The index of the first [AudioFormatListItem] that represents an audio format.
	KAudioFormatProperty_FirstPlayableFormatFromList   KAudioFormatProperty = 'f'<<24 | 'p'<<16 | 'f'<<8 | 'l' // 'fpfl'
	KAudioFormatProperty_FormatEmploysDependentPackets KAudioFormatProperty = 'f'<<24 | 'd'<<16 | 'e'<<8 | 'p' // 'fdep'
	// KAudioFormatProperty_FormatInfo: General information about a format.
	KAudioFormatProperty_FormatInfo        KAudioFormatProperty = 'f'<<24 | 'm'<<16 | 't'<<8 | 'i' // 'fmti'
	KAudioFormatProperty_FormatIsEncrypted KAudioFormatProperty = 'c'<<24 | 'r'<<16 | 'y'<<8 | 'p' // 'cryp'
	// KAudioFormatProperty_FormatIsExternallyFramed: Indicates whether or not a format requires external framing information.
	KAudioFormatProperty_FormatIsExternallyFramed KAudioFormatProperty = 'f'<<24 | 'e'<<16 | 'x'<<8 | 'f' // 'fexf'
	// KAudioFormatProperty_FormatIsVBR: Indicates whether or not a format has a variable number of bytes-per-packet.
	KAudioFormatProperty_FormatIsVBR KAudioFormatProperty = 'f'<<24 | 'v'<<16 | 'b'<<8 | 'r' // 'fvbr'
	// KAudioFormatProperty_FormatList: A list of structures describing the audio formats.
	KAudioFormatProperty_FormatList KAudioFormatProperty = 'f'<<24 | 'l'<<16 | 's'<<8 | 't' // 'flst'
	// KAudioFormatProperty_FormatName: A name for a given format.
	KAudioFormatProperty_FormatName KAudioFormatProperty = 'f'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'fnam'
	// KAudioFormatProperty_ID3TagSize: A [UInt32] value indicating the ID3 tag size.
	KAudioFormatProperty_ID3TagSize KAudioFormatProperty = 'i'<<24 | 'd'<<16 | '3'<<8 | 's' // 'id3s'
	// KAudioFormatProperty_ID3TagToDictionary: A [CFDictionary] object containing key/value pairs for the frames in the ID3 tag.
	KAudioFormatProperty_ID3TagToDictionary KAudioFormatProperty = 'i'<<24 | 'd'<<16 | '3'<<8 | 'd' // 'id3d'
	// KAudioFormatProperty_MatrixMixMap: A matrix of scaling coefficients for converting audio from one channel map to another in a standard way, if one is known.
	KAudioFormatProperty_MatrixMixMap KAudioFormatProperty = 'm'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'mmap'
	// KAudioFormatProperty_NumberOfChannelsForLayout: The number of valid channels.
	KAudioFormatProperty_NumberOfChannelsForLayout KAudioFormatProperty = 'n'<<24 | 'c'<<16 | 'h'<<8 | 'm' // 'nchm'
	// KAudioFormatProperty_OutputFormatList: A list of structures describing the audio formats.
	KAudioFormatProperty_OutputFormatList KAudioFormatProperty = 'o'<<24 | 'f'<<16 | 'l'<<8 | 's' // 'ofls'
	// KAudioFormatProperty_PanningMatrix: An array of [Float32] values, each representing the audio level of one channel.
	KAudioFormatProperty_PanningMatrix KAudioFormatProperty = 'p'<<24 | 'a'<<16 | 'n'<<8 | 'm' // 'panm'
	// KAudioFormatProperty_TagForChannelLayout: An [AudioChannelLayoutTag] value for a layout.
	KAudioFormatProperty_TagForChannelLayout KAudioFormatProperty = 'c'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'cmpt'
	// KAudioFormatProperty_TagsForNumberOfChannels: An array of AudioChannelLayoutTag values for the number of channels specified.
	KAudioFormatProperty_TagsForNumberOfChannels KAudioFormatProperty = 't'<<24 | 'a'<<16 | 'g'<<8 | 'c' // 'tagc'
	// KAudioFormatProperty_ValidateChannelLayout: The validity of an audio channel layout structure.
	KAudioFormatProperty_ValidateChannelLayout KAudioFormatProperty = 'v'<<24 | 'a'<<16 | 'c'<<8 | 'l' // 'vacl'
)

func (e KAudioFormatProperty) String() string {
	switch e {
	case KAudioFormatProperty_ASBDFromESDS:
		return "KAudioFormatProperty_ASBDFromESDS"
	case KAudioFormatProperty_ASBDFromMPEGPacket:
		return "KAudioFormatProperty_ASBDFromMPEGPacket"
	case KAudioFormatProperty_AreChannelLayoutsEquivalent:
		return "KAudioFormatProperty_AreChannelLayoutsEquivalent"
	case KAudioFormatProperty_AvailableDecodeNumberChannels:
		return "KAudioFormatProperty_AvailableDecodeNumberChannels"
	case KAudioFormatProperty_AvailableEncodeBitRates:
		return "KAudioFormatProperty_AvailableEncodeBitRates"
	case KAudioFormatProperty_AvailableEncodeChannelLayoutTags:
		return "KAudioFormatProperty_AvailableEncodeChannelLayoutTags"
	case KAudioFormatProperty_AvailableEncodeNumberChannels:
		return "KAudioFormatProperty_AvailableEncodeNumberChannels"
	case KAudioFormatProperty_AvailableEncodeSampleRates:
		return "KAudioFormatProperty_AvailableEncodeSampleRates"
	case KAudioFormatProperty_BalanceFade:
		return "KAudioFormatProperty_BalanceFade"
	case KAudioFormatProperty_BitmapForLayoutTag:
		return "KAudioFormatProperty_BitmapForLayoutTag"
	case KAudioFormatProperty_ChannelLayoutForBitmap:
		return "KAudioFormatProperty_ChannelLayoutForBitmap"
	case KAudioFormatProperty_ChannelLayoutForTag:
		return "KAudioFormatProperty_ChannelLayoutForTag"
	case KAudioFormatProperty_ChannelLayoutFromESDS:
		return "KAudioFormatProperty_ChannelLayoutFromESDS"
	case KAudioFormatProperty_ChannelLayoutHash:
		return "KAudioFormatProperty_ChannelLayoutHash"
	case KAudioFormatProperty_ChannelLayoutName:
		return "KAudioFormatProperty_ChannelLayoutName"
	case KAudioFormatProperty_ChannelLayoutSimpleName:
		return "KAudioFormatProperty_ChannelLayoutSimpleName"
	case KAudioFormatProperty_ChannelMap:
		return "KAudioFormatProperty_ChannelMap"
	case KAudioFormatProperty_ChannelName:
		return "KAudioFormatProperty_ChannelName"
	case KAudioFormatProperty_ChannelShortName:
		return "KAudioFormatProperty_ChannelShortName"
	case KAudioFormatProperty_DecodeFormatIDs:
		return "KAudioFormatProperty_DecodeFormatIDs"
	case KAudioFormatProperty_Decoders:
		return "KAudioFormatProperty_Decoders"
	case KAudioFormatProperty_EncodeFormatIDs:
		return "KAudioFormatProperty_EncodeFormatIDs"
	case KAudioFormatProperty_Encoders:
		return "KAudioFormatProperty_Encoders"
	case KAudioFormatProperty_FirstPlayableFormatFromList:
		return "KAudioFormatProperty_FirstPlayableFormatFromList"
	case KAudioFormatProperty_FormatEmploysDependentPackets:
		return "KAudioFormatProperty_FormatEmploysDependentPackets"
	case KAudioFormatProperty_FormatInfo:
		return "KAudioFormatProperty_FormatInfo"
	case KAudioFormatProperty_FormatIsEncrypted:
		return "KAudioFormatProperty_FormatIsEncrypted"
	case KAudioFormatProperty_FormatIsExternallyFramed:
		return "KAudioFormatProperty_FormatIsExternallyFramed"
	case KAudioFormatProperty_FormatIsVBR:
		return "KAudioFormatProperty_FormatIsVBR"
	case KAudioFormatProperty_FormatList:
		return "KAudioFormatProperty_FormatList"
	case KAudioFormatProperty_FormatName:
		return "KAudioFormatProperty_FormatName"
	case KAudioFormatProperty_ID3TagSize:
		return "KAudioFormatProperty_ID3TagSize"
	case KAudioFormatProperty_ID3TagToDictionary:
		return "KAudioFormatProperty_ID3TagToDictionary"
	case KAudioFormatProperty_MatrixMixMap:
		return "KAudioFormatProperty_MatrixMixMap"
	case KAudioFormatProperty_NumberOfChannelsForLayout:
		return "KAudioFormatProperty_NumberOfChannelsForLayout"
	case KAudioFormatProperty_OutputFormatList:
		return "KAudioFormatProperty_OutputFormatList"
	case KAudioFormatProperty_PanningMatrix:
		return "KAudioFormatProperty_PanningMatrix"
	case KAudioFormatProperty_TagForChannelLayout:
		return "KAudioFormatProperty_TagForChannelLayout"
	case KAudioFormatProperty_TagsForNumberOfChannels:
		return "KAudioFormatProperty_TagsForNumberOfChannels"
	case KAudioFormatProperty_ValidateChannelLayout:
		return "KAudioFormatProperty_ValidateChannelLayout"
	default:
		return fmt.Sprintf("KAudioFormatProperty(%d)", e)
	}
}

type KAudioHardwareService uint32

const (
	KAudioHardwareServiceDeviceProperty_VirtualMainBalance KAudioHardwareService = 'v'<<24 | 'm'<<16 | 'b'<<8 | 'c' // 'vmbc'
	KAudioHardwareServiceDeviceProperty_VirtualMainVolume  KAudioHardwareService = 'v'<<24 | 'm'<<16 | 'v'<<8 | 'c' // 'vmvc'
	// KAudioHardwareServiceProperty_ServiceRestarted: Used, with a HAL audio object property listener callback, as a flag that indicates a hardware service restart.
	KAudioHardwareServiceProperty_ServiceRestarted KAudioHardwareService = 's'<<24 | 'r'<<16 | 's'<<8 | 't' // 'srst'
	// Deprecated.
	KAudioHardwareServiceDeviceProperty_VirtualMasterBalance KAudioHardwareService = 'v'<<24 | 'm'<<16 | 'b'<<8 | 'c' // 'vmbc'
	// Deprecated.
	KAudioHardwareServiceDeviceProperty_VirtualMasterVolume KAudioHardwareService = 'v'<<24 | 'm'<<16 | 'v'<<8 | 'c' // 'vmvc'
)

func (e KAudioHardwareService) String() string {
	switch e {
	case KAudioHardwareServiceDeviceProperty_VirtualMainBalance:
		return "KAudioHardwareServiceDeviceProperty_VirtualMainBalance"
	case KAudioHardwareServiceDeviceProperty_VirtualMainVolume:
		return "KAudioHardwareServiceDeviceProperty_VirtualMainVolume"
	case KAudioHardwareServiceProperty_ServiceRestarted:
		return "KAudioHardwareServiceProperty_ServiceRestarted"
	default:
		return fmt.Sprintf("KAudioHardwareService(%d)", e)
	}
}

type KAudioOutputUnit uint32

const (
	// KAudioOutputUnitRange: The start of the numerical range for I/O audio unit function selectors.
	KAudioOutputUnitRange KAudioOutputUnit = 0x200
	// KAudioOutputUnitStartSelect: Used by the system to start an I/O audio unit when you call the AudioOutputUnitStart(_:) function.
	KAudioOutputUnitStartSelect KAudioOutputUnit = 0x201
	// KAudioOutputUnitStopSelect: Used by the system to stop an I/O audio unit when you call the AudioOutputUnitStop(_:) function.
	KAudioOutputUnitStopSelect KAudioOutputUnit = 0x202
)

func (e KAudioOutputUnit) String() string {
	switch e {
	case KAudioOutputUnitRange:
		return "KAudioOutputUnitRange"
	case KAudioOutputUnitStartSelect:
		return "KAudioOutputUnitStartSelect"
	case KAudioOutputUnitStopSelect:
		return "KAudioOutputUnitStopSelect"
	default:
		return fmt.Sprintf("KAudioOutputUnit(%d)", e)
	}
}

type KAudioQueue uint32

const (
	// KAudioQueueDeviceProperty_NumberChannels: Value is a read-only [UInt32] value representing the number of channels in the audio hardware device associated with an audio queue.
	KAudioQueueDeviceProperty_NumberChannels KAudioQueue = 'a'<<24 | 'q'<<16 | 'd'<<8 | 'c' // 'aqdc'
	// KAudioQueueDeviceProperty_SampleRate: Value is a read-only [Float64] value representing the sampling rate of the audio hardware device associated with an audio queue.
	KAudioQueueDeviceProperty_SampleRate KAudioQueue = 'a'<<24 | 'q'<<16 | 's'<<8 | 'r' // 'aqsr'
	// KAudioQueueProperty_ChannelLayout: Describes an audio queue channel layout.
	KAudioQueueProperty_ChannelLayout KAudioQueue = 'a'<<24 | 'q'<<16 | 'c'<<8 | 'l' // 'aqcl'
	// KAudioQueueProperty_ConverterError: Value is a  read-only [UInt32] value that indicates the most recent error (if any) encountered by the audio queue’s internal encoding/decoding process.
	KAudioQueueProperty_ConverterError KAudioQueue = 'q'<<24 | 'c'<<16 | 'v'<<8 | 'e' // 'qcve'
	// KAudioQueueProperty_CurrentDevice: The unique identifier (UID) of the audio hardware device.
	KAudioQueueProperty_CurrentDevice KAudioQueue = 'a'<<24 | 'q'<<16 | 'c'<<8 | 'd' // 'aqcd'
	// KAudioQueueProperty_CurrentLevelMeter: A read-only array of level meter status structures.
	KAudioQueueProperty_CurrentLevelMeter KAudioQueue = 'a'<<24 | 'q'<<16 | 'm'<<8 | 'v' // 'aqmv'
	// KAudioQueueProperty_CurrentLevelMeterDB: Value is a read-only array of AudioQueueLevelMeterState structures, one array element per audio channel.
	KAudioQueueProperty_CurrentLevelMeterDB KAudioQueue = 'a'<<24 | 'q'<<16 | 'm'<<8 | 'd' // 'aqmd'
	// KAudioQueueProperty_DecodeBufferSizeFrames: Value is a read/write [UInt32] value that is the size of the buffer into which a playback (output) audio queue decodes buffers.
	KAudioQueueProperty_DecodeBufferSizeFrames KAudioQueue = 'd'<<24 | 'c'<<16 | 'b'<<8 | 'f' // 'dcbf'
	// KAudioQueueProperty_EnableLevelMetering: Value is a read/write [UInt32] value that indicates whether audio level metering is enabled for an audio queue.
	KAudioQueueProperty_EnableLevelMetering       KAudioQueue = 'a'<<24 | 'q'<<16 | 'm'<<8 | 'e' // 'aqme'
	KAudioQueueProperty_EnableTimePitch           KAudioQueue = 'q'<<24 | '_'<<16 | 't'<<8 | 'p' // 'q_tp'
	KAudioQueueProperty_IntendedSpatialExperience KAudioQueue = 'i'<<24 | 's'<<16 | 'e'<<8 | 'o' // 'iseo'
	// KAudioQueueProperty_IsRunning: Value is a read-only [UInt32] value indicating whether or not the audio queue is running.
	KAudioQueueProperty_IsRunning KAudioQueue = 'a'<<24 | 'q'<<16 | 'r'<<8 | 'n' // 'aqrn'
	// KAudioQueueProperty_MagicCookie: Value is a read/write void pointer to a block of memory, which you set up, containing an audio format magic cookie.
	KAudioQueueProperty_MagicCookie KAudioQueue = 'a'<<24 | 'q'<<16 | 'm'<<8 | 'c' // 'aqmc'
	// KAudioQueueProperty_MaximumOutputPacketSize: Value is a read-only[UInt32] value that is the size, in bytes, of the largest single packet of data in the output format.
	KAudioQueueProperty_MaximumOutputPacketSize KAudioQueue = 'x'<<24 | 'o'<<16 | 'p'<<8 | 's' // 'xops'
	// KAudioQueueProperty_StreamDescription: An audio queue’s data format.
	KAudioQueueProperty_StreamDescription  KAudioQueue = 'a'<<24 | 'q'<<16 | 'f'<<8 | 't' // 'aqft'
	KAudioQueueProperty_TimePitchAlgorithm KAudioQueue = 'q'<<24 | 't'<<16 | 'p'<<8 | 'a' // 'qtpa'
	KAudioQueueProperty_TimePitchBypass    KAudioQueue = 'q'<<24 | 't'<<16 | 'p'<<8 | 'b' // 'qtpb'
)

func (e KAudioQueue) String() string {
	switch e {
	case KAudioQueueDeviceProperty_NumberChannels:
		return "KAudioQueueDeviceProperty_NumberChannels"
	case KAudioQueueDeviceProperty_SampleRate:
		return "KAudioQueueDeviceProperty_SampleRate"
	case KAudioQueueProperty_ChannelLayout:
		return "KAudioQueueProperty_ChannelLayout"
	case KAudioQueueProperty_ConverterError:
		return "KAudioQueueProperty_ConverterError"
	case KAudioQueueProperty_CurrentDevice:
		return "KAudioQueueProperty_CurrentDevice"
	case KAudioQueueProperty_CurrentLevelMeter:
		return "KAudioQueueProperty_CurrentLevelMeter"
	case KAudioQueueProperty_CurrentLevelMeterDB:
		return "KAudioQueueProperty_CurrentLevelMeterDB"
	case KAudioQueueProperty_DecodeBufferSizeFrames:
		return "KAudioQueueProperty_DecodeBufferSizeFrames"
	case KAudioQueueProperty_EnableLevelMetering:
		return "KAudioQueueProperty_EnableLevelMetering"
	case KAudioQueueProperty_EnableTimePitch:
		return "KAudioQueueProperty_EnableTimePitch"
	case KAudioQueueProperty_IntendedSpatialExperience:
		return "KAudioQueueProperty_IntendedSpatialExperience"
	case KAudioQueueProperty_IsRunning:
		return "KAudioQueueProperty_IsRunning"
	case KAudioQueueProperty_MagicCookie:
		return "KAudioQueueProperty_MagicCookie"
	case KAudioQueueProperty_MaximumOutputPacketSize:
		return "KAudioQueueProperty_MaximumOutputPacketSize"
	case KAudioQueueProperty_StreamDescription:
		return "KAudioQueueProperty_StreamDescription"
	case KAudioQueueProperty_TimePitchAlgorithm:
		return "KAudioQueueProperty_TimePitchAlgorithm"
	case KAudioQueueProperty_TimePitchBypass:
		return "KAudioQueueProperty_TimePitchBypass"
	default:
		return fmt.Sprintf("KAudioQueue(%d)", e)
	}
}

type KAudioQueueErr int32

const (
	// KAudioQueueErr_BufferEmpty: The audio queue buffer is empty (that is, the `mAudioDataByteSize` field = `0`).
	KAudioQueueErr_BufferEmpty         KAudioQueueErr = -66686
	KAudioQueueErr_BufferEnqueuedTwice KAudioQueueErr = -66666
	// KAudioQueueErr_BufferInQueue: The audio queue buffer cannot be disposed of when it is enqueued.
	KAudioQueueErr_BufferInQueue KAudioQueueErr = -66679
	// KAudioQueueErr_CannotStart: The audio queue has encountered a problem and cannot start.
	KAudioQueueErr_CannotStart    KAudioQueueErr = -66681
	KAudioQueueErr_CannotStartYet KAudioQueueErr = -66665
	// KAudioQueueErr_CodecNotFound: The requested codec was not found.
	KAudioQueueErr_CodecNotFound KAudioQueueErr = -66673
	// KAudioQueueErr_DisposalPending: The function cannot act on the audio queue because it is being asynchronously disposed of.
	KAudioQueueErr_DisposalPending KAudioQueueErr = -66685
	// KAudioQueueErr_EnqueueDuringReset: During a call to the AudioQueueReset(_:), AudioQueueStop(_:_:), or AudioQueueDispose(_:_:) functions, the system does not allow you to enqueue buffers.
	KAudioQueueErr_EnqueueDuringReset KAudioQueueErr = -66632
	// KAudioQueueErr_InvalidBuffer: The specified audio queue buffer does not belong to the specified audio queue.
	KAudioQueueErr_InvalidBuffer KAudioQueueErr = -66687
	// KAudioQueueErr_InvalidCodecAccess: The codec could not be accessed.
	KAudioQueueErr_InvalidCodecAccess KAudioQueueErr = -66672
	// KAudioQueueErr_InvalidDevice: The specified audio hardware device could not be located.
	KAudioQueueErr_InvalidDevice KAudioQueueErr = -66680
	// KAudioQueueErr_InvalidOfflineMode: The operation requires the audio queue to be in offline mode but it isn’t, or vice versa.
	KAudioQueueErr_InvalidOfflineMode KAudioQueueErr = -66626
	// KAudioQueueErr_InvalidParameter: The specified parameter ID is invalid.
	KAudioQueueErr_InvalidParameter KAudioQueueErr = -66682
	// KAudioQueueErr_InvalidProperty: The specified property ID is invalid.
	KAudioQueueErr_InvalidProperty KAudioQueueErr = -66684
	// KAudioQueueErr_InvalidPropertySize: The size of the specified property is invalid.
	KAudioQueueErr_InvalidPropertySize KAudioQueueErr = -66683
	// KAudioQueueErr_InvalidPropertyValue: The property value used is not valid.
	KAudioQueueErr_InvalidPropertyValue KAudioQueueErr = -66675
	// KAudioQueueErr_InvalidQueueType: The queue is an input queue but the function can only operate on an output queue, or vice versa.
	KAudioQueueErr_InvalidQueueType KAudioQueueErr = -66677
	// KAudioQueueErr_InvalidRunState: The queue is running but the function can only operate on the queue when it is stopped, or vice versa.
	KAudioQueueErr_InvalidRunState   KAudioQueueErr = -66678
	KAudioQueueErr_InvalidTapContext KAudioQueueErr = -66669
	KAudioQueueErr_InvalidTapType    KAudioQueueErr = -66667
	// KAudioQueueErr_Permissions: You do not have the required permissions to call the function.
	KAudioQueueErr_Permissions KAudioQueueErr = -66676
	// KAudioQueueErr_PrimeTimedOut: During a call to the AudioQueuePrime(_:_:_:) function, the audio queue’s audio converter failed to convert the requested number of sample frames.
	KAudioQueueErr_PrimeTimedOut KAudioQueueErr = -66674
	// KAudioQueueErr_QueueInvalidated: In iOS, the audio server has exited, causing the audio queue to become invalid.
	KAudioQueueErr_QueueInvalidated KAudioQueueErr = -66671
	// KAudioQueueErr_RecordUnderrun: During recording, data was lost because there was no enqueued buffer to store it in.
	KAudioQueueErr_RecordUnderrun KAudioQueueErr = -66668
	KAudioQueueErr_TooManyTaps    KAudioQueueErr = -66670
)

func (e KAudioQueueErr) String() string {
	switch e {
	case KAudioQueueErr_BufferEmpty:
		return "KAudioQueueErr_BufferEmpty"
	case KAudioQueueErr_BufferEnqueuedTwice:
		return "KAudioQueueErr_BufferEnqueuedTwice"
	case KAudioQueueErr_BufferInQueue:
		return "KAudioQueueErr_BufferInQueue"
	case KAudioQueueErr_CannotStart:
		return "KAudioQueueErr_CannotStart"
	case KAudioQueueErr_CannotStartYet:
		return "KAudioQueueErr_CannotStartYet"
	case KAudioQueueErr_CodecNotFound:
		return "KAudioQueueErr_CodecNotFound"
	case KAudioQueueErr_DisposalPending:
		return "KAudioQueueErr_DisposalPending"
	case KAudioQueueErr_EnqueueDuringReset:
		return "KAudioQueueErr_EnqueueDuringReset"
	case KAudioQueueErr_InvalidBuffer:
		return "KAudioQueueErr_InvalidBuffer"
	case KAudioQueueErr_InvalidCodecAccess:
		return "KAudioQueueErr_InvalidCodecAccess"
	case KAudioQueueErr_InvalidDevice:
		return "KAudioQueueErr_InvalidDevice"
	case KAudioQueueErr_InvalidOfflineMode:
		return "KAudioQueueErr_InvalidOfflineMode"
	case KAudioQueueErr_InvalidParameter:
		return "KAudioQueueErr_InvalidParameter"
	case KAudioQueueErr_InvalidProperty:
		return "KAudioQueueErr_InvalidProperty"
	case KAudioQueueErr_InvalidPropertySize:
		return "KAudioQueueErr_InvalidPropertySize"
	case KAudioQueueErr_InvalidPropertyValue:
		return "KAudioQueueErr_InvalidPropertyValue"
	case KAudioQueueErr_InvalidQueueType:
		return "KAudioQueueErr_InvalidQueueType"
	case KAudioQueueErr_InvalidRunState:
		return "KAudioQueueErr_InvalidRunState"
	case KAudioQueueErr_InvalidTapContext:
		return "KAudioQueueErr_InvalidTapContext"
	case KAudioQueueErr_InvalidTapType:
		return "KAudioQueueErr_InvalidTapType"
	case KAudioQueueErr_Permissions:
		return "KAudioQueueErr_Permissions"
	case KAudioQueueErr_PrimeTimedOut:
		return "KAudioQueueErr_PrimeTimedOut"
	case KAudioQueueErr_QueueInvalidated:
		return "KAudioQueueErr_QueueInvalidated"
	case KAudioQueueErr_RecordUnderrun:
		return "KAudioQueueErr_RecordUnderrun"
	case KAudioQueueErr_TooManyTaps:
		return "KAudioQueueErr_TooManyTaps"
	default:
		return fmt.Sprintf("KAudioQueueErr(%d)", e)
	}
}

type KAudioQueueParam uint32

const (
	// KAudioQueueParam_Pan: The stereo panning position of a source.
	KAudioQueueParam_Pan KAudioQueueParam = 13
	// KAudioQueueParam_Pitch: The number of cents to pitch-shift the audio queue’s playback, in the range `-2400` through `2400` cents (where 1200 cents corresponds to one musical octave.)
	KAudioQueueParam_Pitch KAudioQueueParam = 3
	// KAudioQueueParam_PlayRate: The playback rate for the audio queue, in the range `0.5` through `2.0`.
	KAudioQueueParam_PlayRate KAudioQueueParam = 2
	// KAudioQueueParam_Volume: The playback volume for the audio queue, ranging from `0.0` through `1.0` on a linear scale.
	KAudioQueueParam_Volume KAudioQueueParam = 1
	// KAudioQueueParam_VolumeRampTime: The number of seconds over which a volume change is ramped.
	KAudioQueueParam_VolumeRampTime KAudioQueueParam = 4
)

func (e KAudioQueueParam) String() string {
	switch e {
	case KAudioQueueParam_Pan:
		return "KAudioQueueParam_Pan"
	case KAudioQueueParam_Pitch:
		return "KAudioQueueParam_Pitch"
	case KAudioQueueParam_PlayRate:
		return "KAudioQueueParam_PlayRate"
	case KAudioQueueParam_Volume:
		return "KAudioQueueParam_Volume"
	case KAudioQueueParam_VolumeRampTime:
		return "KAudioQueueParam_VolumeRampTime"
	default:
		return fmt.Sprintf("KAudioQueueParam(%d)", e)
	}
}

type KAudioQueueProperty uint32

const (
	KAudioQueueProperty_ChannelAssignments KAudioQueueProperty = 'a'<<24 | 'q'<<16 | 'c'<<8 | 'a' // 'aqca'
)

func (e KAudioQueueProperty) String() string {
	switch e {
	case KAudioQueueProperty_ChannelAssignments:
		return "KAudioQueueProperty_ChannelAssignments"
	default:
		return fmt.Sprintf("KAudioQueueProperty(%d)", e)
	}
}

type KAudioQueueTimePitchAlgorithm uint32

const (
	KAudioQueueTimePitchAlgorithm_Spectral   KAudioQueueTimePitchAlgorithm = 's'<<24 | 'p'<<16 | 'e'<<8 | 'c' // 'spec'
	KAudioQueueTimePitchAlgorithm_TimeDomain KAudioQueueTimePitchAlgorithm = 't'<<24 | 'i'<<16 | 'd'<<8 | 'o' // 'tido'
	KAudioQueueTimePitchAlgorithm_Varispeed  KAudioQueueTimePitchAlgorithm = 'v'<<24 | 's'<<16 | 'p'<<8 | 'd' // 'vspd'
)

func (e KAudioQueueTimePitchAlgorithm) String() string {
	switch e {
	case KAudioQueueTimePitchAlgorithm_Spectral:
		return "KAudioQueueTimePitchAlgorithm_Spectral"
	case KAudioQueueTimePitchAlgorithm_TimeDomain:
		return "KAudioQueueTimePitchAlgorithm_TimeDomain"
	case KAudioQueueTimePitchAlgorithm_Varispeed:
		return "KAudioQueueTimePitchAlgorithm_Varispeed"
	default:
		return fmt.Sprintf("KAudioQueueTimePitchAlgorithm(%d)", e)
	}
}

type KAudioServices int32

const (
	// KAudioServicesBadPropertySizeError: The size of the property data was not correct.
	KAudioServicesBadPropertySizeError KAudioServices = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	// KAudioServicesBadSpecifierSizeError: The size of the specifier data was not correct.
	KAudioServicesBadSpecifierSizeError KAudioServices = '!'<<24 | 's'<<16 | 'p'<<8 | 'c' // '!spc'
	// KAudioServicesNoError: No error has occurred.
	KAudioServicesNoError KAudioServices = 0
	// KAudioServicesSystemSoundClientTimedOutError: System sound client message timed out.
	KAudioServicesSystemSoundClientTimedOutError          KAudioServices = -1501
	KAudioServicesSystemSoundExceededMaximumDurationError KAudioServices = -1502
	// KAudioServicesSystemSoundUnspecifiedError: An unspecified error has occurred.
	KAudioServicesSystemSoundUnspecifiedError KAudioServices = -1500
	// KAudioServicesUnsupportedPropertyError: The property is not supported.
	KAudioServicesUnsupportedPropertyError KAudioServices = 'p'<<24 | 't'<<16 | 'y'<<8 | '?' // 'pty?'
)

func (e KAudioServices) String() string {
	switch e {
	case KAudioServicesBadPropertySizeError:
		return "KAudioServicesBadPropertySizeError"
	case KAudioServicesBadSpecifierSizeError:
		return "KAudioServicesBadSpecifierSizeError"
	case KAudioServicesNoError:
		return "KAudioServicesNoError"
	case KAudioServicesSystemSoundClientTimedOutError:
		return "KAudioServicesSystemSoundClientTimedOutError"
	case KAudioServicesSystemSoundExceededMaximumDurationError:
		return "KAudioServicesSystemSoundExceededMaximumDurationError"
	case KAudioServicesSystemSoundUnspecifiedError:
		return "KAudioServicesSystemSoundUnspecifiedError"
	case KAudioServicesUnsupportedPropertyError:
		return "KAudioServicesUnsupportedPropertyError"
	default:
		return fmt.Sprintf("KAudioServices(%d)", e)
	}
}

type KAudioServicesProperty uint32

const (
	// KAudioServicesPropertyCompletePlaybackIfAppDies: A [UInt32] value, where `1` means that the audio file specified by a system sound passed in the `inSpecifier` parameter should finish playing even if the client application terminates.
	KAudioServicesPropertyCompletePlaybackIfAppDies KAudioServicesProperty = 'i'<<24 | 'f'<<16 | 'd'<<8 | 'i' // 'ifdi'
	// KAudioServicesPropertyIsUISound: A [UInt32] value, where `1` means that, for the audio file specified by a system sound passed in the `inSpecifier` parameter, the System Sound server respects the user setting in the Sound Effects preference and is silent when the user turns off sound effects.
	KAudioServicesPropertyIsUISound KAudioServicesProperty = 'i'<<24 | 's'<<16 | 'u'<<8 | 'i' // 'isui'
)

func (e KAudioServicesProperty) String() string {
	switch e {
	case KAudioServicesPropertyCompletePlaybackIfAppDies:
		return "KAudioServicesPropertyCompletePlaybackIfAppDies"
	case KAudioServicesPropertyIsUISound:
		return "KAudioServicesPropertyIsUISound"
	default:
		return fmt.Sprintf("KAudioServicesProperty(%d)", e)
	}
}

type KAudioSession uint32

const (
	// KAudioSessionBeginInterruption: Your app’s audio session has just been interrupted, such as by a phone call.
	KAudioSessionBeginInterruption KAudioSession = 1
	// KAudioSessionEndInterruption: The interruption to your app’s audio session has just ended.
	KAudioSessionEndInterruption KAudioSession = 0
)

func (e KAudioSession) String() string {
	switch e {
	case KAudioSessionBeginInterruption:
		return "KAudioSessionBeginInterruption"
	case KAudioSessionEndInterruption:
		return "KAudioSessionEndInterruption"
	default:
		return fmt.Sprintf("KAudioSession(%d)", e)
	}
}

type KAudioSessionMode uint32

const (
	// KAudioSessionMode_Default: The default mode; used unless you set a mode with the AudioSessionSetProperty(_:_:_:) function.
	KAudioSessionMode_Default  KAudioSessionMode = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KAudioSessionMode_GameChat KAudioSessionMode = 'g'<<24 | 'm'<<16 | 'c'<<8 | 't' // 'gmct'
	// KAudioSessionMode_Measurement: Specify this mode if your app is performing measurement of incoming audio.
	KAudioSessionMode_Measurement KAudioSessionMode = 'm'<<24 | 's'<<16 | 'm'<<8 | 't' // 'msmt'
	// KAudioSessionMode_VideoRecording: Specify this mode if your app is recording a movie.
	KAudioSessionMode_VideoRecording KAudioSessionMode = 'v'<<24 | 'r'<<16 | 'c'<<8 | 'd' // 'vrcd'
	// KAudioSessionMode_VoiceChat: Specify this mode if your app is performing two-way voice communication, such as using Voice over Internet Protocol (VoIP).
	KAudioSessionMode_VoiceChat KAudioSessionMode = 'v'<<24 | 'c'<<16 | 'c'<<8 | 't' // 'vcct'
)

func (e KAudioSessionMode) String() string {
	switch e {
	case KAudioSessionMode_Default:
		return "KAudioSessionMode_Default"
	case KAudioSessionMode_GameChat:
		return "KAudioSessionMode_GameChat"
	case KAudioSessionMode_Measurement:
		return "KAudioSessionMode_Measurement"
	case KAudioSessionMode_VideoRecording:
		return "KAudioSessionMode_VideoRecording"
	case KAudioSessionMode_VoiceChat:
		return "KAudioSessionMode_VoiceChat"
	default:
		return fmt.Sprintf("KAudioSessionMode(%d)", e)
	}
}

type KAudioSessionNoError uint32

const (
	// KAudioServicesNoHardwareError: The audio operation failed because the device has no audio input available.
	KAudioServicesNoHardwareError KAudioSessionNoError = 'n'<<24 | 'o'<<16 | 'h'<<8 | 'w' // 'nohw'
	// KAudioSessionAlreadyInitialized: The [AudioSessionInitialize] function was called more than once during the lifetime of your application.
	KAudioSessionAlreadyInitialized KAudioSessionNoError = 'i'<<24 | 'n'<<16 | 'i'<<8 | 't' // 'init'
	// KAudioSessionBadPropertySizeError: The size of the audio session property data was not correct.
	KAudioSessionBadPropertySizeError KAudioSessionNoError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	// KAudioSessionIncompatibleCategory: The specified audio session category cannot be used for the attempted audio operation.
	KAudioSessionIncompatibleCategory KAudioSessionNoError = '!'<<24 | 'c'<<16 | 'a'<<8 | 't' // '!cat'
	// KAudioSessionInitializationError: There was an error during audio session initialization.
	KAudioSessionInitializationError KAudioSessionNoError = 'i'<<24 | 'n'<<16 | 'i'<<8 | '?' // 'ini?'
	// KAudioSessionNoCategorySet: The audio operation failed because it requires the audio session to have an explicitly-set category, but none was set.
	KAudioSessionNoCategorySet KAudioSessionNoError = '?'<<24 | 'c'<<16 | 'a'<<8 | 't' // '?cat'
	// KAudioSessionNoErrorValue: No error has occurred.
	KAudioSessionNoErrorValue KAudioSessionNoError = 0
	// KAudioSessionNotActiveError: The audio operation failed because your application’s audio session was not active.
	KAudioSessionNotActiveError KAudioSessionNoError = '!'<<24 | 'a'<<16 | 'c'<<8 | 't' // '!act'
	// KAudioSessionNotInitialized: An Audio Session function was called without first initializing the session.
	KAudioSessionNotInitialized KAudioSessionNoError = '!'<<24 | 'i'<<16 | 'n'<<8 | 'i' // '!ini'
	// KAudioSessionUnspecifiedError: An unspecified audio session error has occurred.
	KAudioSessionUnspecifiedError KAudioSessionNoError = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	// KAudioSessionUnsupportedPropertyError: The audio session property is not supported.
	KAudioSessionUnsupportedPropertyError KAudioSessionNoError = 'p'<<24 | 't'<<16 | 'y'<<8 | '?' // 'pty?'
)

func (e KAudioSessionNoError) String() string {
	switch e {
	case KAudioServicesNoHardwareError:
		return "KAudioServicesNoHardwareError"
	case KAudioSessionAlreadyInitialized:
		return "KAudioSessionAlreadyInitialized"
	case KAudioSessionBadPropertySizeError:
		return "KAudioSessionBadPropertySizeError"
	case KAudioSessionIncompatibleCategory:
		return "KAudioSessionIncompatibleCategory"
	case KAudioSessionInitializationError:
		return "KAudioSessionInitializationError"
	case KAudioSessionNoCategorySet:
		return "KAudioSessionNoCategorySet"
	case KAudioSessionNoErrorValue:
		return "KAudioSessionNoErrorValue"
	case KAudioSessionNotActiveError:
		return "KAudioSessionNotActiveError"
	case KAudioSessionNotInitialized:
		return "KAudioSessionNotInitialized"
	case KAudioSessionUnspecifiedError:
		return "KAudioSessionUnspecifiedError"
	case KAudioSessionUnsupportedPropertyError:
		return "KAudioSessionUnsupportedPropertyError"
	default:
		return fmt.Sprintf("KAudioSessionNoError(%d)", e)
	}
}

type KAudioSessionOverrideAudioRoute uint32

const (
	// KAudioSessionOverrideAudioRoute_None: Specifies, for the kAudioSessionCategory_PlayAndRecord category, that output audio should go to the receiver.
	KAudioSessionOverrideAudioRoute_None KAudioSessionOverrideAudioRoute = 0
	// KAudioSessionOverrideAudioRoute_Speaker: Specifies, for the kAudioSessionCategory_PlayAndRecord category,  that output audio should go to the speaker, not the receiver.
	KAudioSessionOverrideAudioRoute_Speaker KAudioSessionOverrideAudioRoute = 's'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'spkr'
)

func (e KAudioSessionOverrideAudioRoute) String() string {
	switch e {
	case KAudioSessionOverrideAudioRoute_None:
		return "KAudioSessionOverrideAudioRoute_None"
	case KAudioSessionOverrideAudioRoute_Speaker:
		return "KAudioSessionOverrideAudioRoute_Speaker"
	default:
		return fmt.Sprintf("KAudioSessionOverrideAudioRoute(%d)", e)
	}
}

type KAudioSessionProperty uint32

const (
	// KAudioSessionProperty_AudioCategory: The category for the audio session.
	KAudioSessionProperty_AudioCategory KAudioSessionProperty = 'a'<<24 | 'c'<<16 | 'a'<<8 | 't' // 'acat'
	// KAudioSessionProperty_AudioInputAvailable: Indicates if audio input is available (a nonzero value) or not (a value of 0).
	KAudioSessionProperty_AudioInputAvailable KAudioSessionProperty = 'a'<<24 | 'i'<<16 | 'a'<<8 | 'v' // 'aiav'
	// KAudioSessionProperty_AudioRouteChange: The reason the audio route changed.
	KAudioSessionProperty_AudioRouteChange KAudioSessionProperty = 'r'<<24 | 'o'<<16 | 'c'<<8 | 'h' // 'roch'
	// KAudioSessionProperty_AudioRouteDescription: Information about an audio route.
	KAudioSessionProperty_AudioRouteDescription KAudioSessionProperty = 'c'<<24 | 'r'<<16 | 'a'<<8 | 'r' // 'crar'
	// KAudioSessionProperty_CurrentHardwareIOBufferDuration: Indicates the current hardware IO buffer duration, in seconds, as a read-only [Float32] value.
	KAudioSessionProperty_CurrentHardwareIOBufferDuration KAudioSessionProperty = 'c'<<24 | 'h'<<16 | 'b'<<8 | 'd' // 'chbd'
	// KAudioSessionProperty_CurrentHardwareInputLatency: Indicates the current hardware input latency, in seconds, as a read-only [Float32] value.
	KAudioSessionProperty_CurrentHardwareInputLatency KAudioSessionProperty = 'c'<<24 | 'i'<<16 | 'l'<<8 | 't' // 'cilt'
	// KAudioSessionProperty_CurrentHardwareInputNumberChannels: Indicates the current number of audio hardware input channels.
	KAudioSessionProperty_CurrentHardwareInputNumberChannels KAudioSessionProperty = 'c'<<24 | 'h'<<16 | 'i'<<8 | 'c' // 'chic'
	// KAudioSessionProperty_CurrentHardwareOutputLatency: Indicates the current hardware output latency, in seconds, as a read-only [Float32] value.
	KAudioSessionProperty_CurrentHardwareOutputLatency KAudioSessionProperty = 'c'<<24 | 'o'<<16 | 'l'<<8 | 't' // 'colt'
	// KAudioSessionProperty_CurrentHardwareOutputNumberChannels: Indicates the current number of audio hardware output channels.
	KAudioSessionProperty_CurrentHardwareOutputNumberChannels KAudioSessionProperty = 'c'<<24 | 'h'<<16 | 'o'<<8 | 'c' // 'choc'
	// KAudioSessionProperty_CurrentHardwareOutputVolume: Indicates the current audio output volume as [Float32] value between 0.0 and 1.0.
	KAudioSessionProperty_CurrentHardwareOutputVolume KAudioSessionProperty = 'c'<<24 | 'h'<<16 | 'o'<<8 | 'v' // 'chov'
	// KAudioSessionProperty_CurrentHardwareSampleRate: Indicates the current hardware sample rate.
	KAudioSessionProperty_CurrentHardwareSampleRate KAudioSessionProperty = 'c'<<24 | 'h'<<16 | 's'<<8 | 'r' // 'chsr'
	// KAudioSessionProperty_InputGainAvailable: A read-only [UInt32] value that indicates whether or not audio input gain adjustment is available, where a nonzero value means adjustment is available.
	KAudioSessionProperty_InputGainAvailable KAudioSessionProperty = 'i'<<24 | 'g'<<16 | 'a'<<8 | 'v' // 'igav'
	// KAudioSessionProperty_InputGainScalar: A read/write [Float32] value that indicates the audio input gain setting for the active input source.
	KAudioSessionProperty_InputGainScalar KAudioSessionProperty = 'i'<<24 | 'g'<<16 | 's'<<8 | 'c' // 'igsc'
	// KAudioSessionProperty_InputSource: The audio input source.
	KAudioSessionProperty_InputSource KAudioSessionProperty = 'i'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'isrc'
	// KAudioSessionProperty_InputSources: Details on the available audio input sources.
	KAudioSessionProperty_InputSources KAudioSessionProperty = 's'<<24 | 'r'<<16 | 'c'<<8 | 's' // 'srcs'
	// KAudioSessionProperty_InterruptionType: Indicates the type of an end-interruption event.
	KAudioSessionProperty_InterruptionType KAudioSessionProperty = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
	// KAudioSessionProperty_Mode: A read/write [UIInt32] value that specifies the audio session mode.
	KAudioSessionProperty_Mode KAudioSessionProperty = 'm'<<24 | 'o'<<16 | 'd'<<8 | 'e' // 'mode'
	// KAudioSessionProperty_OtherAudioIsPlaying: Indicates whether or not another app (typically, the iPod app) is currently playing audio.
	KAudioSessionProperty_OtherAudioIsPlaying KAudioSessionProperty = 'o'<<24 | 't'<<16 | 'h'<<8 | 'r' // 'othr'
	// KAudioSessionProperty_OtherMixableAudioShouldDuck: For audio session categories that allow audio mixing with other apps, specifies whether other audio should be reduced in level when your app produces sound.
	KAudioSessionProperty_OtherMixableAudioShouldDuck KAudioSessionProperty = 'd'<<24 | 'u'<<16 | 'c'<<8 | 'k' // 'duck'
	// KAudioSessionProperty_OutputDestination: The audio output destination.
	KAudioSessionProperty_OutputDestination KAudioSessionProperty = 'o'<<24 | 'd'<<16 | 's'<<8 | 't' // 'odst'
	// KAudioSessionProperty_OutputDestinations: Details on the available audio output destinations.
	KAudioSessionProperty_OutputDestinations KAudioSessionProperty = 'd'<<24 | 's'<<16 | 't'<<8 | 's' // 'dsts'
	// KAudioSessionProperty_OverrideAudioRoute: Specifies whether or not to override the audio session category’s typical audio route.
	KAudioSessionProperty_OverrideAudioRoute KAudioSessionProperty = 'o'<<24 | 'v'<<16 | 'r'<<8 | 'd' // 'ovrd'
	// KAudioSessionProperty_OverrideCategoryDefaultToSpeaker: Specifies whether or not to route audio to the speaker (instead of to the receiver) when no other audio route, such as a headset, is connected.
	KAudioSessionProperty_OverrideCategoryDefaultToSpeaker KAudioSessionProperty = 'c'<<24 | 's'<<16 | 'p'<<8 | 'k' // 'cspk'
	// KAudioSessionProperty_OverrideCategoryEnableBluetoothInput: Allows a paired Bluetooth device to appear as an available audio input route.
	KAudioSessionProperty_OverrideCategoryEnableBluetoothInput KAudioSessionProperty = 'c'<<24 | 'b'<<16 | 'l'<<8 | 'u' // 'cblu'
	// KAudioSessionProperty_OverrideCategoryMixWithOthers: Changes the mixing behavior of the kAudioSessionCategory_MediaPlayback and kAudioSessionCategory_PlayAndRecord audio session categories.
	KAudioSessionProperty_OverrideCategoryMixWithOthers KAudioSessionProperty = 'c'<<24 | 'm'<<16 | 'i'<<8 | 'x' // 'cmix'
	// KAudioSessionProperty_PreferredHardwareIOBufferDuration: Your preferred hardware I/O buffer duration in seconds.
	KAudioSessionProperty_PreferredHardwareIOBufferDuration KAudioSessionProperty = 'i'<<24 | 'o'<<16 | 'b'<<8 | 'd' // 'iobd'
	// KAudioSessionProperty_PreferredHardwareSampleRate: Your preferred hardware sample rate for the audio session.
	KAudioSessionProperty_PreferredHardwareSampleRate KAudioSessionProperty = 'h'<<24 | 'w'<<16 | 's'<<8 | 'r' // 'hwsr'
	// KAudioSessionProperty_ServerDied: Indicates if the audio server has died (indicated by a nonzero [UInt32] value) or is still running (a value of 0).
	KAudioSessionProperty_ServerDied KAudioSessionProperty = 'd'<<24 | 'i'<<16 | 'e'<<8 | 'd' // 'died'
)

func (e KAudioSessionProperty) String() string {
	switch e {
	case KAudioSessionProperty_AudioCategory:
		return "KAudioSessionProperty_AudioCategory"
	case KAudioSessionProperty_AudioInputAvailable:
		return "KAudioSessionProperty_AudioInputAvailable"
	case KAudioSessionProperty_AudioRouteChange:
		return "KAudioSessionProperty_AudioRouteChange"
	case KAudioSessionProperty_AudioRouteDescription:
		return "KAudioSessionProperty_AudioRouteDescription"
	case KAudioSessionProperty_CurrentHardwareIOBufferDuration:
		return "KAudioSessionProperty_CurrentHardwareIOBufferDuration"
	case KAudioSessionProperty_CurrentHardwareInputLatency:
		return "KAudioSessionProperty_CurrentHardwareInputLatency"
	case KAudioSessionProperty_CurrentHardwareInputNumberChannels:
		return "KAudioSessionProperty_CurrentHardwareInputNumberChannels"
	case KAudioSessionProperty_CurrentHardwareOutputLatency:
		return "KAudioSessionProperty_CurrentHardwareOutputLatency"
	case KAudioSessionProperty_CurrentHardwareOutputNumberChannels:
		return "KAudioSessionProperty_CurrentHardwareOutputNumberChannels"
	case KAudioSessionProperty_CurrentHardwareOutputVolume:
		return "KAudioSessionProperty_CurrentHardwareOutputVolume"
	case KAudioSessionProperty_CurrentHardwareSampleRate:
		return "KAudioSessionProperty_CurrentHardwareSampleRate"
	case KAudioSessionProperty_InputGainAvailable:
		return "KAudioSessionProperty_InputGainAvailable"
	case KAudioSessionProperty_InputGainScalar:
		return "KAudioSessionProperty_InputGainScalar"
	case KAudioSessionProperty_InputSource:
		return "KAudioSessionProperty_InputSource"
	case KAudioSessionProperty_InputSources:
		return "KAudioSessionProperty_InputSources"
	case KAudioSessionProperty_InterruptionType:
		return "KAudioSessionProperty_InterruptionType"
	case KAudioSessionProperty_Mode:
		return "KAudioSessionProperty_Mode"
	case KAudioSessionProperty_OtherAudioIsPlaying:
		return "KAudioSessionProperty_OtherAudioIsPlaying"
	case KAudioSessionProperty_OtherMixableAudioShouldDuck:
		return "KAudioSessionProperty_OtherMixableAudioShouldDuck"
	case KAudioSessionProperty_OutputDestination:
		return "KAudioSessionProperty_OutputDestination"
	case KAudioSessionProperty_OutputDestinations:
		return "KAudioSessionProperty_OutputDestinations"
	case KAudioSessionProperty_OverrideAudioRoute:
		return "KAudioSessionProperty_OverrideAudioRoute"
	case KAudioSessionProperty_OverrideCategoryDefaultToSpeaker:
		return "KAudioSessionProperty_OverrideCategoryDefaultToSpeaker"
	case KAudioSessionProperty_OverrideCategoryEnableBluetoothInput:
		return "KAudioSessionProperty_OverrideCategoryEnableBluetoothInput"
	case KAudioSessionProperty_OverrideCategoryMixWithOthers:
		return "KAudioSessionProperty_OverrideCategoryMixWithOthers"
	case KAudioSessionProperty_PreferredHardwareIOBufferDuration:
		return "KAudioSessionProperty_PreferredHardwareIOBufferDuration"
	case KAudioSessionProperty_PreferredHardwareSampleRate:
		return "KAudioSessionProperty_PreferredHardwareSampleRate"
	case KAudioSessionProperty_ServerDied:
		return "KAudioSessionProperty_ServerDied"
	default:
		return fmt.Sprintf("KAudioSessionProperty(%d)", e)
	}
}

type KAudioSessionRouteChangeReason uint32

const (
	// KAudioSessionRouteChangeReason_CategoryChange: The audio session category has changed.
	KAudioSessionRouteChangeReason_CategoryChange KAudioSessionRouteChangeReason = 3
	// KAudioSessionRouteChangeReason_NewDeviceAvailable: A new audio hardware device became available; for example, a headset was plugged in.
	KAudioSessionRouteChangeReason_NewDeviceAvailable KAudioSessionRouteChangeReason = 1
	// KAudioSessionRouteChangeReason_NoSuitableRouteForCategory: There is no audio hardware route for the audio session category.
	KAudioSessionRouteChangeReason_NoSuitableRouteForCategory KAudioSessionRouteChangeReason = 7
	// KAudioSessionRouteChangeReason_OldDeviceUnavailable: The previously-used audio hardware device is now unavailable; for example, a headset was unplugged.
	KAudioSessionRouteChangeReason_OldDeviceUnavailable KAudioSessionRouteChangeReason = 2
	// KAudioSessionRouteChangeReason_Override: The audio route has been overridden.
	KAudioSessionRouteChangeReason_Override                 KAudioSessionRouteChangeReason = 4
	KAudioSessionRouteChangeReason_RouteConfigurationChange KAudioSessionRouteChangeReason = 8
	// KAudioSessionRouteChangeReason_Unknown: The audio route changed but the reason is not known.
	KAudioSessionRouteChangeReason_Unknown KAudioSessionRouteChangeReason = 0
	// KAudioSessionRouteChangeReason_WakeFromSleep: The device woke from sleep.
	KAudioSessionRouteChangeReason_WakeFromSleep KAudioSessionRouteChangeReason = 6
)

func (e KAudioSessionRouteChangeReason) String() string {
	switch e {
	case KAudioSessionRouteChangeReason_CategoryChange:
		return "KAudioSessionRouteChangeReason_CategoryChange"
	case KAudioSessionRouteChangeReason_NewDeviceAvailable:
		return "KAudioSessionRouteChangeReason_NewDeviceAvailable"
	case KAudioSessionRouteChangeReason_NoSuitableRouteForCategory:
		return "KAudioSessionRouteChangeReason_NoSuitableRouteForCategory"
	case KAudioSessionRouteChangeReason_OldDeviceUnavailable:
		return "KAudioSessionRouteChangeReason_OldDeviceUnavailable"
	case KAudioSessionRouteChangeReason_Override:
		return "KAudioSessionRouteChangeReason_Override"
	case KAudioSessionRouteChangeReason_RouteConfigurationChange:
		return "KAudioSessionRouteChangeReason_RouteConfigurationChange"
	case KAudioSessionRouteChangeReason_Unknown:
		return "KAudioSessionRouteChangeReason_Unknown"
	case KAudioSessionRouteChangeReason_WakeFromSleep:
		return "KAudioSessionRouteChangeReason_WakeFromSleep"
	default:
		return fmt.Sprintf("KAudioSessionRouteChangeReason(%d)", e)
	}
}

type KAudioToolbox int32

const (
	KAudioToolboxErr_CannotDoInCurrentContext KAudioToolbox = -10863
	KAudioToolboxErr_EndOfTrack               KAudioToolbox = -10857
	KAudioToolboxErr_IllegalTrackDestination  KAudioToolbox = -10855
	KAudioToolboxErr_InvalidEventType         KAudioToolbox = -10853
	KAudioToolboxErr_InvalidPlayerState       KAudioToolbox = -10852
	KAudioToolboxErr_InvalidSequenceType      KAudioToolbox = -10846
	KAudioToolboxErr_NoSequence               KAudioToolbox = -10854
	KAudioToolboxErr_StartOfTrack             KAudioToolbox = -10856
	KAudioToolboxErr_TrackIndexError          KAudioToolbox = -10859
	KAudioToolboxErr_TrackNotFound            KAudioToolbox = -10858
	KAudioToolboxError_NoTrackDestination     KAudioToolbox = -66720
)

func (e KAudioToolbox) String() string {
	switch e {
	case KAudioToolboxErr_CannotDoInCurrentContext:
		return "KAudioToolboxErr_CannotDoInCurrentContext"
	case KAudioToolboxErr_EndOfTrack:
		return "KAudioToolboxErr_EndOfTrack"
	case KAudioToolboxErr_IllegalTrackDestination:
		return "KAudioToolboxErr_IllegalTrackDestination"
	case KAudioToolboxErr_InvalidEventType:
		return "KAudioToolboxErr_InvalidEventType"
	case KAudioToolboxErr_InvalidPlayerState:
		return "KAudioToolboxErr_InvalidPlayerState"
	case KAudioToolboxErr_InvalidSequenceType:
		return "KAudioToolboxErr_InvalidSequenceType"
	case KAudioToolboxErr_NoSequence:
		return "KAudioToolboxErr_NoSequence"
	case KAudioToolboxErr_StartOfTrack:
		return "KAudioToolboxErr_StartOfTrack"
	case KAudioToolboxErr_TrackIndexError:
		return "KAudioToolboxErr_TrackIndexError"
	case KAudioToolboxErr_TrackNotFound:
		return "KAudioToolboxErr_TrackNotFound"
	case KAudioToolboxError_NoTrackDestination:
		return "KAudioToolboxError_NoTrackDestination"
	default:
		return fmt.Sprintf("KAudioToolbox(%d)", e)
	}
}

type KAudioUnit uint32

const (
	// KAudioUnitAddPropertyListenerSelect: Used by the system to register a property listener callback function for an audio unit when you call the AudioUnitAddPropertyListener(_:_:_:_:) function.
	KAudioUnitAddPropertyListenerSelect KAudioUnit = 0xa
	KAudioUnitAddRenderNotifySelect     KAudioUnit = 0xf
	KAudioUnitComplexRenderSelect       KAudioUnit = 0x13
	// KAudioUnitGetParameterSelect: Used by the system to get the current value of an audio unit parameter when you call the AudioUnitGetParameter(_:_:_:_:_:) function.
	KAudioUnitGetParameterSelect KAudioUnit = 0x6
	// KAudioUnitGetPropertyInfoSelect: Used by the system to get property information from an audio unit when you call the AudioUnitGetPropertyInfo(_:_:_:_:_:_:) function.
	KAudioUnitGetPropertyInfoSelect KAudioUnit = 0x3
	KAudioUnitGetPropertySelect     KAudioUnit = 0x4
	KAudioUnitInitializeSelect      KAudioUnit = 0x1
	KAudioUnitProcessMultipleSelect KAudioUnit = 0x15
	KAudioUnitProcessSelect         KAudioUnit = 0x14
	KAudioUnitRange                 KAudioUnit = 0
	// KAudioUnitRemovePropertyListenerSelect: Used by the system to unregister a property listener callback function from an audio unit when you call the [AudioUnitRemovePropertyListener] function.
	KAudioUnitRemovePropertyListenerSelect             KAudioUnit = 0xb
	KAudioUnitRemovePropertyListenerWithUserDataSelect KAudioUnit = 0x12
	KAudioUnitRemoveRenderNotifySelect                 KAudioUnit = 0x10
	// KAudioUnitRenderSelect: Used by the system to invoke audio rendering by an audio unit when you call the AudioUnitRender(_:_:_:_:_:_:) function.
	KAudioUnitRenderSelect KAudioUnit = 0xe
	// KAudioUnitResetSelect: Used by the system to reset an audio unit when you call the AudioUnitReset(_:_:_:) function.
	KAudioUnitResetSelect KAudioUnit = 0x9
	// KAudioUnitScheduleParametersSelect: Used by the system to schedule an audio unit parameter value change when you call the AudioUnitScheduleParameters(_:_:_:) function.
	KAudioUnitScheduleParametersSelect KAudioUnit = 0x11
	// KAudioUnitSetParameterSelect: Used by the system to set the value of an audio unit parameter when you call the AudioUnitSetParameter(_:_:_:_:_:_:) function.
	KAudioUnitSetParameterSelect KAudioUnit = 0x7
	KAudioUnitSetPropertySelect  KAudioUnit = 0x5
	// KAudioUnitUninitializeSelect: Used by the system to uninitialize an audio unit when you call the AudioUnitUninitialize(_:) function.
	KAudioUnitUninitializeSelect KAudioUnit = 0x2
)

func (e KAudioUnit) String() string {
	switch e {
	case KAudioUnitAddPropertyListenerSelect:
		return "KAudioUnitAddPropertyListenerSelect"
	case KAudioUnitAddRenderNotifySelect:
		return "KAudioUnitAddRenderNotifySelect"
	case KAudioUnitComplexRenderSelect:
		return "KAudioUnitComplexRenderSelect"
	case KAudioUnitGetParameterSelect:
		return "KAudioUnitGetParameterSelect"
	case KAudioUnitGetPropertyInfoSelect:
		return "KAudioUnitGetPropertyInfoSelect"
	case KAudioUnitGetPropertySelect:
		return "KAudioUnitGetPropertySelect"
	case KAudioUnitInitializeSelect:
		return "KAudioUnitInitializeSelect"
	case KAudioUnitProcessMultipleSelect:
		return "KAudioUnitProcessMultipleSelect"
	case KAudioUnitProcessSelect:
		return "KAudioUnitProcessSelect"
	case KAudioUnitRange:
		return "KAudioUnitRange"
	case KAudioUnitRemovePropertyListenerSelect:
		return "KAudioUnitRemovePropertyListenerSelect"
	case KAudioUnitRemovePropertyListenerWithUserDataSelect:
		return "KAudioUnitRemovePropertyListenerWithUserDataSelect"
	case KAudioUnitRemoveRenderNotifySelect:
		return "KAudioUnitRemoveRenderNotifySelect"
	case KAudioUnitRenderSelect:
		return "KAudioUnitRenderSelect"
	case KAudioUnitResetSelect:
		return "KAudioUnitResetSelect"
	case KAudioUnitScheduleParametersSelect:
		return "KAudioUnitScheduleParametersSelect"
	case KAudioUnitSetParameterSelect:
		return "KAudioUnitSetParameterSelect"
	case KAudioUnitSetPropertySelect:
		return "KAudioUnitSetPropertySelect"
	case KAudioUnitUninitializeSelect:
		return "KAudioUnitUninitializeSelect"
	default:
		return fmt.Sprintf("KAudioUnit(%d)", e)
	}
}

type KAudioUnitCarbonView uint

const (
	KAudioUnitCarbonViewCreateSelect           KAudioUnitCarbonView = 0
	KAudioUnitCarbonViewRange                  KAudioUnitCarbonView = 0
	KAudioUnitCarbonViewSetEventListenerSelect KAudioUnitCarbonView = 0
)

func (e KAudioUnitCarbonView) String() string {
	switch e {
	case KAudioUnitCarbonViewCreateSelect:
		return "KAudioUnitCarbonViewCreateSelect"
	default:
		return fmt.Sprintf("KAudioUnitCarbonView(%d)", e)
	}
}

type KAudioUnitCarbonViewComponentType uint

const (
	KAUCarbonViewSubType_Generic           KAudioUnitCarbonViewComponentType = 0
	KAudioUnitCarbonViewComponentTypeValue KAudioUnitCarbonViewComponentType = 0
)

func (e KAudioUnitCarbonViewComponentType) String() string {
	switch e {
	case KAUCarbonViewSubType_Generic:
		return "KAUCarbonViewSubType_Generic"
	default:
		return fmt.Sprintf("KAudioUnitCarbonViewComponentType(%d)", e)
	}
}

type KAudioUnitErr int32

const (
	KAudioUnitErr_IllegalInstrument      KAudioUnitErr = -10873
	KAudioUnitErr_InstrumentTypeNotFound KAudioUnitErr = -10872
)

func (e KAudioUnitErr) String() string {
	switch e {
	case KAudioUnitErr_IllegalInstrument:
		return "KAudioUnitErr_IllegalInstrument"
	case KAudioUnitErr_InstrumentTypeNotFound:
		return "KAudioUnitErr_InstrumentTypeNotFound"
	default:
		return fmt.Sprintf("KAudioUnitErr(%d)", e)
	}
}

type KAudioUnitMigrateProperty uint32

const (
	KAudioUnitMigrateProperty_FromPlugin    KAudioUnitMigrateProperty = 4000
	KAudioUnitMigrateProperty_OldAutomation KAudioUnitMigrateProperty = 4001
)

func (e KAudioUnitMigrateProperty) String() string {
	switch e {
	case KAudioUnitMigrateProperty_FromPlugin:
		return "KAudioUnitMigrateProperty_FromPlugin"
	case KAudioUnitMigrateProperty_OldAutomation:
		return "KAudioUnitMigrateProperty_OldAutomation"
	default:
		return fmt.Sprintf("KAudioUnitMigrateProperty(%d)", e)
	}
}

type KAudioUnitOfflineProperty uint32

const (
	KAudioUnitOfflineProperty_InputSize             KAudioUnitOfflineProperty = 3020
	KAudioUnitOfflineProperty_OutputSize            KAudioUnitOfflineProperty = 3021
	KAudioUnitOfflineProperty_PreflightName         KAudioUnitOfflineProperty = 3024
	KAudioUnitOfflineProperty_PreflightRequirements KAudioUnitOfflineProperty = 3023
	KAudioUnitOfflineProperty_StartOffset           KAudioUnitOfflineProperty = 3022
)

func (e KAudioUnitOfflineProperty) String() string {
	switch e {
	case KAudioUnitOfflineProperty_InputSize:
		return "KAudioUnitOfflineProperty_InputSize"
	case KAudioUnitOfflineProperty_OutputSize:
		return "KAudioUnitOfflineProperty_OutputSize"
	case KAudioUnitOfflineProperty_PreflightName:
		return "KAudioUnitOfflineProperty_PreflightName"
	case KAudioUnitOfflineProperty_PreflightRequirements:
		return "KAudioUnitOfflineProperty_PreflightRequirements"
	case KAudioUnitOfflineProperty_StartOffset:
		return "KAudioUnitOfflineProperty_StartOffset"
	default:
		return fmt.Sprintf("KAudioUnitOfflineProperty(%d)", e)
	}
}

type KAudioUnitSRCAlgorithm uint32

const (
	KAudioUnitSRCAlgorithm_MediumQuality KAudioUnitSRCAlgorithm = 'c'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'csrc'
	KAudioUnitSRCAlgorithm_Polyphase     KAudioUnitSRCAlgorithm = 'p'<<24 | 'o'<<16 | 'l'<<8 | 'y' // 'poly'
)

func (e KAudioUnitSRCAlgorithm) String() string {
	switch e {
	case KAudioUnitSRCAlgorithm_MediumQuality:
		return "KAudioUnitSRCAlgorithm_MediumQuality"
	case KAudioUnitSRCAlgorithm_Polyphase:
		return "KAudioUnitSRCAlgorithm_Polyphase"
	default:
		return fmt.Sprintf("KAudioUnitSRCAlgorithm(%d)", e)
	}
}

type KAudioUnitSampleRateConverterComplexity uint32

const (
	// KAudioUnitSampleRateConverterComplexity_Linear: Basic sample rate conversion using linear interpolation.
	KAudioUnitSampleRateConverterComplexity_Linear KAudioUnitSampleRateConverterComplexity = 'l'<<24 | 'i'<<16 | 'n'<<8 | 'e' // 'line'
	// KAudioUnitSampleRateConverterComplexity_Mastering: Mastering quality sample rate conversion.
	KAudioUnitSampleRateConverterComplexity_Mastering KAudioUnitSampleRateConverterComplexity = 'b'<<24 | 'a'<<16 | 't'<<8 | 's' // 'bats'
	KAudioUnitSampleRateConverterComplexity_Normal    KAudioUnitSampleRateConverterComplexity = 'n'<<24 | 'o'<<16 | 'r'<<8 | 'm' // 'norm'
)

func (e KAudioUnitSampleRateConverterComplexity) String() string {
	switch e {
	case KAudioUnitSampleRateConverterComplexity_Linear:
		return "KAudioUnitSampleRateConverterComplexity_Linear"
	case KAudioUnitSampleRateConverterComplexity_Mastering:
		return "KAudioUnitSampleRateConverterComplexity_Mastering"
	case KAudioUnitSampleRateConverterComplexity_Normal:
		return "KAudioUnitSampleRateConverterComplexity_Normal"
	default:
		return fmt.Sprintf("KAudioUnitSampleRateConverterComplexity(%d)", e)
	}
}

type KAudioUnitScope uint32

const (
	KAudioUnitScope_Global KAudioUnitScope = 0
	// KAudioUnitScope_Group: In macOS, a context specific to the control scope of audio unit parameters.
	KAudioUnitScope_Group KAudioUnitScope = 3
	// KAudioUnitScope_Input: The context for audio data coming into an audio unit.
	KAudioUnitScope_Input     KAudioUnitScope = 1
	KAudioUnitScope_Layer     KAudioUnitScope = 6
	KAudioUnitScope_LayerItem KAudioUnitScope = 7
	// KAudioUnitScope_Note: In macOS, a scope for changes to an individual musical note.
	KAudioUnitScope_Note KAudioUnitScope = 5
	// KAudioUnitScope_Output: The context for audio data leaving an audio unit.
	KAudioUnitScope_Output KAudioUnitScope = 2
	KAudioUnitScope_Part   KAudioUnitScope = 4
)

func (e KAudioUnitScope) String() string {
	switch e {
	case KAudioUnitScope_Global:
		return "KAudioUnitScope_Global"
	case KAudioUnitScope_Group:
		return "KAudioUnitScope_Group"
	case KAudioUnitScope_Input:
		return "KAudioUnitScope_Input"
	case KAudioUnitScope_Layer:
		return "KAudioUnitScope_Layer"
	case KAudioUnitScope_LayerItem:
		return "KAudioUnitScope_LayerItem"
	case KAudioUnitScope_Note:
		return "KAudioUnitScope_Note"
	case KAudioUnitScope_Output:
		return "KAudioUnitScope_Output"
	case KAudioUnitScope_Part:
		return "KAudioUnitScope_Part"
	default:
		return fmt.Sprintf("KAudioUnitScope(%d)", e)
	}
}

type KAudioUnitType uint32

const (
	KAudioUnitType_Effect          KAudioUnitType = 'a'<<24 | 'u'<<16 | 'f'<<8 | 'x' // 'aufx'
	KAudioUnitType_FormatConverter KAudioUnitType = 'a'<<24 | 'u'<<16 | 'f'<<8 | 'c' // 'aufc'
	// KAudioUnitType_Generator: A generator unit provides audio output but has no audio input.
	KAudioUnitType_Generator     KAudioUnitType = 'a'<<24 | 'u'<<16 | 'g'<<8 | 'n' // 'augn'
	KAudioUnitType_MIDIProcessor KAudioUnitType = 'a'<<24 | 'u'<<16 | 'm'<<8 | 'i' // 'aumi'
	// KAudioUnitType_Mixer: A mixer unit takes a number of input channels and mixes them to provide one or more output channels.
	KAudioUnitType_Mixer KAudioUnitType = 'a'<<24 | 'u'<<16 | 'm'<<8 | 'x' // 'aumx'
	// KAudioUnitType_MusicDevice: An instrument unit can be used as a software musical instrument, such as a sampler or synthesizer.
	KAudioUnitType_MusicDevice KAudioUnitType = 'a'<<24 | 'u'<<16 | 'm'<<8 | 'u' // 'aumu'
	// KAudioUnitType_MusicEffect: An effect unit that can respond to MIDI control messages, typically through a mapping of  MIDI messages to parameters of the audio unit’s DSP algorithm.
	KAudioUnitType_MusicEffect KAudioUnitType = 'a'<<24 | 'u'<<16 | 'm'<<8 | 'f' // 'aumf'
	// KAudioUnitType_OfflineEffect: An offline effect unit provides digital signal processing of a sort that cannot proceed in realtime.
	KAudioUnitType_OfflineEffect KAudioUnitType = 'a'<<24 | 'u'<<16 | 'o'<<8 | 'l' // 'auol'
	// KAudioUnitType_Output: An output unit provides input, output, or both input and output simultaneously.
	KAudioUnitType_Output            KAudioUnitType = 'a'<<24 | 'u'<<16 | 'o'<<8 | 'u' // 'auou'
	KAudioUnitType_Panner            KAudioUnitType = 'a'<<24 | 'u'<<16 | 'p'<<8 | 'n' // 'aupn'
	KAudioUnitType_SpeechSynthesizer KAudioUnitType = 'a'<<24 | 'u'<<16 | 's'<<8 | 'p' // 'ausp'
)

func (e KAudioUnitType) String() string {
	switch e {
	case KAudioUnitType_Effect:
		return "KAudioUnitType_Effect"
	case KAudioUnitType_FormatConverter:
		return "KAudioUnitType_FormatConverter"
	case KAudioUnitType_Generator:
		return "KAudioUnitType_Generator"
	case KAudioUnitType_MIDIProcessor:
		return "KAudioUnitType_MIDIProcessor"
	case KAudioUnitType_Mixer:
		return "KAudioUnitType_Mixer"
	case KAudioUnitType_MusicDevice:
		return "KAudioUnitType_MusicDevice"
	case KAudioUnitType_MusicEffect:
		return "KAudioUnitType_MusicEffect"
	case KAudioUnitType_OfflineEffect:
		return "KAudioUnitType_OfflineEffect"
	case KAudioUnitType_Output:
		return "KAudioUnitType_Output"
	case KAudioUnitType_Panner:
		return "KAudioUnitType_Panner"
	case KAudioUnitType_SpeechSynthesizer:
		return "KAudioUnitType_SpeechSynthesizer"
	default:
		return fmt.Sprintf("KAudioUnitType(%d)", e)
	}
}

type KBandpassParam uint32

const (
	KBandpassParam_Bandwidth       KBandpassParam = 1
	KBandpassParam_CenterFrequency KBandpassParam = 0
)

func (e KBandpassParam) String() string {
	switch e {
	case KBandpassParam_Bandwidth:
		return "KBandpassParam_Bandwidth"
	case KBandpassParam_CenterFrequency:
		return "KBandpassParam_CenterFrequency"
	default:
		return fmt.Sprintf("KBandpassParam(%d)", e)
	}
}

type KCAClock int32

const (
	KCAClock_CannotSetTimeError         KCAClock = -66805
	KCAClock_InvalidPlayRateError       KCAClock = -66806
	KCAClock_InvalidPropertySizeError   KCAClock = -66815
	KCAClock_InvalidSMPTEFormatError    KCAClock = -66809
	KCAClock_InvalidSMPTEOffsetError    KCAClock = -66808
	KCAClock_InvalidSyncModeError       KCAClock = -66813
	KCAClock_InvalidSyncSourceError     KCAClock = -66812
	KCAClock_InvalidTimeFormatError     KCAClock = -66814
	KCAClock_InvalidTimebaseError       KCAClock = -66811
	KCAClock_InvalidTimebaseSourceError KCAClock = -66810
	KCAClock_InvalidUnitError           KCAClock = -66807
	KCAClock_UnknownPropertyError       KCAClock = -66816
)

func (e KCAClock) String() string {
	switch e {
	case KCAClock_CannotSetTimeError:
		return "KCAClock_CannotSetTimeError"
	case KCAClock_InvalidPlayRateError:
		return "KCAClock_InvalidPlayRateError"
	case KCAClock_InvalidPropertySizeError:
		return "KCAClock_InvalidPropertySizeError"
	case KCAClock_InvalidSMPTEFormatError:
		return "KCAClock_InvalidSMPTEFormatError"
	case KCAClock_InvalidSMPTEOffsetError:
		return "KCAClock_InvalidSMPTEOffsetError"
	case KCAClock_InvalidSyncModeError:
		return "KCAClock_InvalidSyncModeError"
	case KCAClock_InvalidSyncSourceError:
		return "KCAClock_InvalidSyncSourceError"
	case KCAClock_InvalidTimeFormatError:
		return "KCAClock_InvalidTimeFormatError"
	case KCAClock_InvalidTimebaseError:
		return "KCAClock_InvalidTimebaseError"
	case KCAClock_InvalidTimebaseSourceError:
		return "KCAClock_InvalidTimebaseSourceError"
	case KCAClock_InvalidUnitError:
		return "KCAClock_InvalidUnitError"
	case KCAClock_UnknownPropertyError:
		return "KCAClock_UnknownPropertyError"
	default:
		return fmt.Sprintf("KCAClock(%d)", e)
	}
}

type KCAF uint32

const (
	KCAF_AudioDataChunkID         KCAF = 'd'<<24 | 'a'<<16 | 't'<<8 | 'a' // 'data'
	KCAF_ChannelLayoutChunkID     KCAF = 'c'<<24 | 'h'<<16 | 'a'<<8 | 'n' // 'chan'
	KCAF_EditCommentsChunkID      KCAF = 'e'<<24 | 'd'<<16 | 'c'<<8 | 't' // 'edct'
	KCAF_FillerChunkID            KCAF = 'f'<<24 | 'r'<<16 | 'e'<<8 | 'e' // 'free'
	KCAF_FormatListID             KCAF = 'l'<<24 | 'd'<<16 | 's'<<8 | 'c' // 'ldsc'
	KCAF_InfoStringsChunkID       KCAF = 'i'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'info'
	KCAF_InstrumentChunkID        KCAF = 'i'<<24 | 'n'<<16 | 's'<<8 | 't' // 'inst'
	KCAF_MIDIChunkID              KCAF = 'm'<<24 | 'i'<<16 | 'd'<<8 | 'i' // 'midi'
	KCAF_MagicCookieID            KCAF = 'k'<<24 | 'u'<<16 | 'k'<<8 | 'i' // 'kuki'
	KCAF_MarkerChunkID            KCAF = 'm'<<24 | 'a'<<16 | 'r'<<8 | 'k' // 'mark'
	KCAF_OverviewChunkID          KCAF = 'o'<<24 | 'v'<<16 | 'v'<<8 | 'w' // 'ovvw'
	KCAF_PacketTableChunkID       KCAF = 'p'<<24 | 'a'<<16 | 'k'<<8 | 't' // 'pakt'
	KCAF_PeakChunkID              KCAF = 'p'<<24 | 'e'<<16 | 'a'<<8 | 'k' // 'peak'
	KCAF_RegionChunkID            KCAF = 'r'<<24 | 'e'<<16 | 'g'<<8 | 'n' // 'regn'
	KCAF_StreamDescriptionChunkID KCAF = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	KCAF_StringsChunkID           KCAF = 's'<<24 | 't'<<16 | 'r'<<8 | 'g' // 'strg'
	KCAF_UMIDChunkID              KCAF = 'u'<<24 | 'm'<<16 | 'i'<<8 | 'd' // 'umid'
	KCAF_UUIDChunkID              KCAF = 'u'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'uuid'
	KCAF_iXMLChunkID              KCAF = 'i'<<24 | 'X'<<16 | 'M'<<8 | 'L' // 'iXML'
)

func (e KCAF) String() string {
	switch e {
	case KCAF_AudioDataChunkID:
		return "KCAF_AudioDataChunkID"
	case KCAF_ChannelLayoutChunkID:
		return "KCAF_ChannelLayoutChunkID"
	case KCAF_EditCommentsChunkID:
		return "KCAF_EditCommentsChunkID"
	case KCAF_FillerChunkID:
		return "KCAF_FillerChunkID"
	case KCAF_FormatListID:
		return "KCAF_FormatListID"
	case KCAF_InfoStringsChunkID:
		return "KCAF_InfoStringsChunkID"
	case KCAF_InstrumentChunkID:
		return "KCAF_InstrumentChunkID"
	case KCAF_MIDIChunkID:
		return "KCAF_MIDIChunkID"
	case KCAF_MagicCookieID:
		return "KCAF_MagicCookieID"
	case KCAF_MarkerChunkID:
		return "KCAF_MarkerChunkID"
	case KCAF_OverviewChunkID:
		return "KCAF_OverviewChunkID"
	case KCAF_PacketTableChunkID:
		return "KCAF_PacketTableChunkID"
	case KCAF_PeakChunkID:
		return "KCAF_PeakChunkID"
	case KCAF_RegionChunkID:
		return "KCAF_RegionChunkID"
	case KCAF_StreamDescriptionChunkID:
		return "KCAF_StreamDescriptionChunkID"
	case KCAF_StringsChunkID:
		return "KCAF_StringsChunkID"
	case KCAF_UMIDChunkID:
		return "KCAF_UMIDChunkID"
	case KCAF_UUIDChunkID:
		return "KCAF_UUIDChunkID"
	case KCAF_iXMLChunkID:
		return "KCAF_iXMLChunkID"
	default:
		return fmt.Sprintf("KCAF(%d)", e)
	}
}

type KCAFMarkerType uint32

const (
	KCAFMarkerType_EditDestinationBegin KCAFMarkerType = 'd'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'dbeg'
	KCAFMarkerType_EditDestinationEnd   KCAFMarkerType = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'dend'
	KCAFMarkerType_EditSourceBegin      KCAFMarkerType = 'c'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'cbeg'
	KCAFMarkerType_EditSourceEnd        KCAFMarkerType = 'c'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'cend'
	KCAFMarkerType_Generic              KCAFMarkerType = 0
	KCAFMarkerType_Index                KCAFMarkerType = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'x' // 'indx'
	KCAFMarkerType_KeySignature         KCAFMarkerType = 'k'<<24 | 's'<<16 | 'i'<<8 | 'g' // 'ksig'
	KCAFMarkerType_ProgramEnd           KCAFMarkerType = 'p'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'pend'
	KCAFMarkerType_ProgramStart         KCAFMarkerType = 'p'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'pbeg'
	KCAFMarkerType_RegionEnd            KCAFMarkerType = 'r'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'rend'
	KCAFMarkerType_RegionStart          KCAFMarkerType = 'r'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'rbeg'
	KCAFMarkerType_RegionSyncPoint      KCAFMarkerType = 'r'<<24 | 's'<<16 | 'y'<<8 | 'c' // 'rsyc'
	KCAFMarkerType_ReleaseLoopEnd       KCAFMarkerType = 'r'<<24 | 'l'<<16 | 'e'<<8 | 'n' // 'rlen'
	KCAFMarkerType_ReleaseLoopStart     KCAFMarkerType = 'r'<<24 | 'l'<<16 | 'b'<<8 | 'g' // 'rlbg'
	KCAFMarkerType_SavedPlayPosition    KCAFMarkerType = 's'<<24 | 'p'<<16 | 'l'<<8 | 'y' // 'sply'
	KCAFMarkerType_SelectionEnd         KCAFMarkerType = 's'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'send'
	KCAFMarkerType_SelectionStart       KCAFMarkerType = 's'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'sbeg'
	KCAFMarkerType_SustainLoopEnd       KCAFMarkerType = 's'<<24 | 'l'<<16 | 'e'<<8 | 'n' // 'slen'
	KCAFMarkerType_SustainLoopStart     KCAFMarkerType = 's'<<24 | 'l'<<16 | 'b'<<8 | 'g' // 'slbg'
	KCAFMarkerType_Tempo                KCAFMarkerType = 't'<<24 | 'm'<<16 | 'p'<<8 | 'o' // 'tmpo'
	KCAFMarkerType_TimeSignature        KCAFMarkerType = 't'<<24 | 's'<<16 | 'i'<<8 | 'g' // 'tsig'
	KCAFMarkerType_TrackEnd             KCAFMarkerType = 't'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'tend'
	KCAFMarkerType_TrackStart           KCAFMarkerType = 't'<<24 | 'b'<<16 | 'e'<<8 | 'g' // 'tbeg'
)

func (e KCAFMarkerType) String() string {
	switch e {
	case KCAFMarkerType_EditDestinationBegin:
		return "KCAFMarkerType_EditDestinationBegin"
	case KCAFMarkerType_EditDestinationEnd:
		return "KCAFMarkerType_EditDestinationEnd"
	case KCAFMarkerType_EditSourceBegin:
		return "KCAFMarkerType_EditSourceBegin"
	case KCAFMarkerType_EditSourceEnd:
		return "KCAFMarkerType_EditSourceEnd"
	case KCAFMarkerType_Generic:
		return "KCAFMarkerType_Generic"
	case KCAFMarkerType_Index:
		return "KCAFMarkerType_Index"
	case KCAFMarkerType_KeySignature:
		return "KCAFMarkerType_KeySignature"
	case KCAFMarkerType_ProgramEnd:
		return "KCAFMarkerType_ProgramEnd"
	case KCAFMarkerType_ProgramStart:
		return "KCAFMarkerType_ProgramStart"
	case KCAFMarkerType_RegionEnd:
		return "KCAFMarkerType_RegionEnd"
	case KCAFMarkerType_RegionStart:
		return "KCAFMarkerType_RegionStart"
	case KCAFMarkerType_RegionSyncPoint:
		return "KCAFMarkerType_RegionSyncPoint"
	case KCAFMarkerType_ReleaseLoopEnd:
		return "KCAFMarkerType_ReleaseLoopEnd"
	case KCAFMarkerType_ReleaseLoopStart:
		return "KCAFMarkerType_ReleaseLoopStart"
	case KCAFMarkerType_SavedPlayPosition:
		return "KCAFMarkerType_SavedPlayPosition"
	case KCAFMarkerType_SelectionEnd:
		return "KCAFMarkerType_SelectionEnd"
	case KCAFMarkerType_SelectionStart:
		return "KCAFMarkerType_SelectionStart"
	case KCAFMarkerType_SustainLoopEnd:
		return "KCAFMarkerType_SustainLoopEnd"
	case KCAFMarkerType_SustainLoopStart:
		return "KCAFMarkerType_SustainLoopStart"
	case KCAFMarkerType_Tempo:
		return "KCAFMarkerType_Tempo"
	case KCAFMarkerType_TimeSignature:
		return "KCAFMarkerType_TimeSignature"
	case KCAFMarkerType_TrackEnd:
		return "KCAFMarkerType_TrackEnd"
	case KCAFMarkerType_TrackStart:
		return "KCAFMarkerType_TrackStart"
	default:
		return fmt.Sprintf("KCAFMarkerType(%d)", e)
	}
}

type KConverterPrimeMethod uint32

const (
	// KConverterPrimeMethod_None: Acts in “latency” mode.
	KConverterPrimeMethod_None KConverterPrimeMethod = 2
	// KConverterPrimeMethod_Normal: Prime with `trailing` frames only, for zero latency.
	KConverterPrimeMethod_Normal KConverterPrimeMethod = 1
	// KConverterPrimeMethod_Pre: Prime with `leading` + `trailing` input frames.
	KConverterPrimeMethod_Pre KConverterPrimeMethod = 0
)

func (e KConverterPrimeMethod) String() string {
	switch e {
	case KConverterPrimeMethod_None:
		return "KConverterPrimeMethod_None"
	case KConverterPrimeMethod_Normal:
		return "KConverterPrimeMethod_Normal"
	case KConverterPrimeMethod_Pre:
		return "KConverterPrimeMethod_Pre"
	default:
		return fmt.Sprintf("KConverterPrimeMethod(%d)", e)
	}
}

type KDelayParam uint32

const (
	KDelayParam_DelayTime    KDelayParam = 1
	KDelayParam_Feedback     KDelayParam = 2
	KDelayParam_LopassCutoff KDelayParam = 3
	KDelayParam_WetDryMix    KDelayParam = 0
)

func (e KDelayParam) String() string {
	switch e {
	case KDelayParam_DelayTime:
		return "KDelayParam_DelayTime"
	case KDelayParam_Feedback:
		return "KDelayParam_Feedback"
	case KDelayParam_LopassCutoff:
		return "KDelayParam_LopassCutoff"
	case KDelayParam_WetDryMix:
		return "KDelayParam_WetDryMix"
	default:
		return fmt.Sprintf("KDelayParam(%d)", e)
	}
}

type KDistortionParam uint32

const (
	KDistortionParam_CubicTerm      KDistortionParam = 8
	KDistortionParam_Decay          KDistortionParam = 1
	KDistortionParam_Decimation     KDistortionParam = 3
	KDistortionParam_DecimationMix  KDistortionParam = 5
	KDistortionParam_Delay          KDistortionParam = 0
	KDistortionParam_DelayMix       KDistortionParam = 2
	KDistortionParam_FinalMix       KDistortionParam = 15
	KDistortionParam_LinearTerm     KDistortionParam = 6
	KDistortionParam_PolynomialMix  KDistortionParam = 9
	KDistortionParam_RingModBalance KDistortionParam = 12
	KDistortionParam_RingModFreq1   KDistortionParam = 10
	KDistortionParam_RingModFreq2   KDistortionParam = 11
	KDistortionParam_RingModMix     KDistortionParam = 13
	KDistortionParam_Rounding       KDistortionParam = 4
	KDistortionParam_SoftClipGain   KDistortionParam = 14
	KDistortionParam_SquaredTerm    KDistortionParam = 7
)

func (e KDistortionParam) String() string {
	switch e {
	case KDistortionParam_CubicTerm:
		return "KDistortionParam_CubicTerm"
	case KDistortionParam_Decay:
		return "KDistortionParam_Decay"
	case KDistortionParam_Decimation:
		return "KDistortionParam_Decimation"
	case KDistortionParam_DecimationMix:
		return "KDistortionParam_DecimationMix"
	case KDistortionParam_Delay:
		return "KDistortionParam_Delay"
	case KDistortionParam_DelayMix:
		return "KDistortionParam_DelayMix"
	case KDistortionParam_FinalMix:
		return "KDistortionParam_FinalMix"
	case KDistortionParam_LinearTerm:
		return "KDistortionParam_LinearTerm"
	case KDistortionParam_PolynomialMix:
		return "KDistortionParam_PolynomialMix"
	case KDistortionParam_RingModBalance:
		return "KDistortionParam_RingModBalance"
	case KDistortionParam_RingModFreq1:
		return "KDistortionParam_RingModFreq1"
	case KDistortionParam_RingModFreq2:
		return "KDistortionParam_RingModFreq2"
	case KDistortionParam_RingModMix:
		return "KDistortionParam_RingModMix"
	case KDistortionParam_Rounding:
		return "KDistortionParam_Rounding"
	case KDistortionParam_SoftClipGain:
		return "KDistortionParam_SoftClipGain"
	case KDistortionParam_SquaredTerm:
		return "KDistortionParam_SquaredTerm"
	default:
		return fmt.Sprintf("KDistortionParam(%d)", e)
	}
}

type KDitherAlgorithm uint32

const (
	KDitherAlgorithm_NoiseShaping KDitherAlgorithm = 2
	KDitherAlgorithm_TPDF         KDitherAlgorithm = 1
)

func (e KDitherAlgorithm) String() string {
	switch e {
	case KDitherAlgorithm_NoiseShaping:
		return "KDitherAlgorithm_NoiseShaping"
	case KDitherAlgorithm_TPDF:
		return "KDitherAlgorithm_TPDF"
	default:
		return fmt.Sprintf("KDitherAlgorithm(%d)", e)
	}
}

type KDynamicRangeCompressionProfile uint32

const (
	KDynamicRangeCompressionProfile_GeneralCompression   KDynamicRangeCompressionProfile = 6
	KDynamicRangeCompressionProfile_LateNight            KDynamicRangeCompressionProfile = 1
	KDynamicRangeCompressionProfile_LimitedPlaybackRange KDynamicRangeCompressionProfile = 3
	KDynamicRangeCompressionProfile_NoisyEnvironment     KDynamicRangeCompressionProfile = 2
	KDynamicRangeCompressionProfile_None                 KDynamicRangeCompressionProfile = 0
)

func (e KDynamicRangeCompressionProfile) String() string {
	switch e {
	case KDynamicRangeCompressionProfile_GeneralCompression:
		return "KDynamicRangeCompressionProfile_GeneralCompression"
	case KDynamicRangeCompressionProfile_LateNight:
		return "KDynamicRangeCompressionProfile_LateNight"
	case KDynamicRangeCompressionProfile_LimitedPlaybackRange:
		return "KDynamicRangeCompressionProfile_LimitedPlaybackRange"
	case KDynamicRangeCompressionProfile_NoisyEnvironment:
		return "KDynamicRangeCompressionProfile_NoisyEnvironment"
	case KDynamicRangeCompressionProfile_None:
		return "KDynamicRangeCompressionProfile_None"
	default:
		return fmt.Sprintf("KDynamicRangeCompressionProfile(%d)", e)
	}
}

type KDynamicRangeControlMode uint32

const (
	KDynamicRangeControlMode_Heavy KDynamicRangeControlMode = 2
	KDynamicRangeControlMode_Light KDynamicRangeControlMode = 1
	KDynamicRangeControlMode_None  KDynamicRangeControlMode = 0
)

func (e KDynamicRangeControlMode) String() string {
	switch e {
	case KDynamicRangeControlMode_Heavy:
		return "KDynamicRangeControlMode_Heavy"
	case KDynamicRangeControlMode_Light:
		return "KDynamicRangeControlMode_Light"
	case KDynamicRangeControlMode_None:
		return "KDynamicRangeControlMode_None"
	default:
		return fmt.Sprintf("KDynamicRangeControlMode(%d)", e)
	}
}

type KDynamicsProcessorParam uint32

const (
	KDynamicsProcessorParam_AttackTime         KDynamicsProcessorParam = 4
	KDynamicsProcessorParam_CompressionAmount  KDynamicsProcessorParam = 1000
	KDynamicsProcessorParam_ExpansionRatio     KDynamicsProcessorParam = 2
	KDynamicsProcessorParam_ExpansionThreshold KDynamicsProcessorParam = 3
	KDynamicsProcessorParam_HeadRoom           KDynamicsProcessorParam = 1
	KDynamicsProcessorParam_InputAmplitude     KDynamicsProcessorParam = 2000
	KDynamicsProcessorParam_OutputAmplitude    KDynamicsProcessorParam = 3000
	KDynamicsProcessorParam_OverallGain        KDynamicsProcessorParam = 6
	KDynamicsProcessorParam_ReleaseTime        KDynamicsProcessorParam = 5
	KDynamicsProcessorParam_Threshold          KDynamicsProcessorParam = 0
	// Deprecated: use KDynamicsProcessorParam_OverallGain.
	KDynamicsProcessorParam_MasterGain KDynamicsProcessorParam = 6
)

func (e KDynamicsProcessorParam) String() string {
	switch e {
	case KDynamicsProcessorParam_AttackTime:
		return "KDynamicsProcessorParam_AttackTime"
	case KDynamicsProcessorParam_CompressionAmount:
		return "KDynamicsProcessorParam_CompressionAmount"
	case KDynamicsProcessorParam_ExpansionRatio:
		return "KDynamicsProcessorParam_ExpansionRatio"
	case KDynamicsProcessorParam_ExpansionThreshold:
		return "KDynamicsProcessorParam_ExpansionThreshold"
	case KDynamicsProcessorParam_HeadRoom:
		return "KDynamicsProcessorParam_HeadRoom"
	case KDynamicsProcessorParam_InputAmplitude:
		return "KDynamicsProcessorParam_InputAmplitude"
	case KDynamicsProcessorParam_OutputAmplitude:
		return "KDynamicsProcessorParam_OutputAmplitude"
	case KDynamicsProcessorParam_OverallGain:
		return "KDynamicsProcessorParam_OverallGain"
	case KDynamicsProcessorParam_ReleaseTime:
		return "KDynamicsProcessorParam_ReleaseTime"
	case KDynamicsProcessorParam_Threshold:
		return "KDynamicsProcessorParam_Threshold"
	default:
		return fmt.Sprintf("KDynamicsProcessorParam(%d)", e)
	}
}

type KExtAudioFileError int32

const (
	// KExtAudioFileError_AsyncWriteBufferOverflow: An asynchronous write operation could not be completed in time.
	KExtAudioFileError_AsyncWriteBufferOverflow KExtAudioFileError = -66570
	KExtAudioFileError_AsyncWriteTooLarge       KExtAudioFileError = -66569
	// KExtAudioFileError_InvalidChannelMap: The number of channels does not match the specified format.
	KExtAudioFileError_InvalidChannelMap     KExtAudioFileError = -66564
	KExtAudioFileError_InvalidDataFormat     KExtAudioFileError = -66566
	KExtAudioFileError_InvalidOperationOrder KExtAudioFileError = -66565
	KExtAudioFileError_InvalidProperty       KExtAudioFileError = -66561
	KExtAudioFileError_InvalidPropertySize   KExtAudioFileError = -66562
	// KExtAudioFileError_InvalidSeek: An attempt to write, or an offset, is out of bounds.
	KExtAudioFileError_InvalidSeek          KExtAudioFileError = -66568
	KExtAudioFileError_MaxPacketSizeUnknown KExtAudioFileError = -66567
	KExtAudioFileError_NonPCMClientFormat   KExtAudioFileError = -66563
)

func (e KExtAudioFileError) String() string {
	switch e {
	case KExtAudioFileError_AsyncWriteBufferOverflow:
		return "KExtAudioFileError_AsyncWriteBufferOverflow"
	case KExtAudioFileError_AsyncWriteTooLarge:
		return "KExtAudioFileError_AsyncWriteTooLarge"
	case KExtAudioFileError_InvalidChannelMap:
		return "KExtAudioFileError_InvalidChannelMap"
	case KExtAudioFileError_InvalidDataFormat:
		return "KExtAudioFileError_InvalidDataFormat"
	case KExtAudioFileError_InvalidOperationOrder:
		return "KExtAudioFileError_InvalidOperationOrder"
	case KExtAudioFileError_InvalidProperty:
		return "KExtAudioFileError_InvalidProperty"
	case KExtAudioFileError_InvalidPropertySize:
		return "KExtAudioFileError_InvalidPropertySize"
	case KExtAudioFileError_InvalidSeek:
		return "KExtAudioFileError_InvalidSeek"
	case KExtAudioFileError_MaxPacketSizeUnknown:
		return "KExtAudioFileError_MaxPacketSizeUnknown"
	case KExtAudioFileError_NonPCMClientFormat:
		return "KExtAudioFileError_NonPCMClientFormat"
	default:
		return fmt.Sprintf("KExtAudioFileError(%d)", e)
	}
}

type KExtAudioFilePacketTableInfoOverride int32

const (
	KExtAudioFilePacketTableInfoOverride_UseFileValue        KExtAudioFilePacketTableInfoOverride = -1
	KExtAudioFilePacketTableInfoOverride_UseFileValueIfValid KExtAudioFilePacketTableInfoOverride = -2
)

func (e KExtAudioFilePacketTableInfoOverride) String() string {
	switch e {
	case KExtAudioFilePacketTableInfoOverride_UseFileValue:
		return "KExtAudioFilePacketTableInfoOverride_UseFileValue"
	case KExtAudioFilePacketTableInfoOverride_UseFileValueIfValid:
		return "KExtAudioFilePacketTableInfoOverride_UseFileValueIfValid"
	default:
		return fmt.Sprintf("KExtAudioFilePacketTableInfoOverride(%d)", e)
	}
}

type KExtAudioFileProperty uint32

const (
	// KExtAudioFileProperty_AudioConverter: The audio converter object associated with the extended audio file object, if a converter is associated.
	KExtAudioFileProperty_AudioConverter KExtAudioFileProperty = 'a'<<24 | 'c'<<16 | 'n'<<8 | 'v' // 'acnv'
	// KExtAudioFileProperty_AudioFile: The audio file object associated with the extended audio file object.
	KExtAudioFileProperty_AudioFile KExtAudioFileProperty = 'a'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'afil'
	// KExtAudioFileProperty_ClientChannelLayout: The audio channel layout for your application.
	KExtAudioFileProperty_ClientChannelLayout KExtAudioFileProperty = 'c'<<24 | 'c'<<16 | 'l'<<8 | 'o' // 'cclo'
	// KExtAudioFileProperty_ClientDataFormat: The audio stream format for your application.
	KExtAudioFileProperty_ClientDataFormat KExtAudioFileProperty = 'c'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'cfmt'
	// KExtAudioFileProperty_ClientMaxPacketSize: Your application audio data format’s maximum packet size, in bytes.
	KExtAudioFileProperty_ClientMaxPacketSize KExtAudioFileProperty = 'c'<<24 | 'm'<<16 | 'p'<<8 | 's' // 'cmps'
	// KExtAudioFileProperty_CodecManufacturer: The manufacturer of the codec to be used by the extended audio file object.
	KExtAudioFileProperty_CodecManufacturer KExtAudioFileProperty = 'c'<<24 | 'm'<<16 | 'a'<<8 | 'n' // 'cman'
	// KExtAudioFileProperty_ConverterConfig: The configuration of the extended audio file object’s associated audio converter, as specified by the `kAudioConverterPropertySettings` property.
	KExtAudioFileProperty_ConverterConfig KExtAudioFileProperty = 'a'<<24 | 'c'<<16 | 'c'<<8 | 'f' // 'accf'
	// KExtAudioFileProperty_FileChannelLayout: A file’s channel layout.
	KExtAudioFileProperty_FileChannelLayout KExtAudioFileProperty = 'f'<<24 | 'c'<<16 | 'l'<<8 | 'o' // 'fclo'
	// KExtAudioFileProperty_FileDataFormat: A file’s data format.
	KExtAudioFileProperty_FileDataFormat KExtAudioFileProperty = 'f'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'ffmt'
	// KExtAudioFileProperty_FileLengthFrames: The associated audio file’s length in sample frames.
	KExtAudioFileProperty_FileLengthFrames KExtAudioFileProperty = '#'<<24 | 'f'<<16 | 'r'<<8 | 'm' // '#frm'
	// KExtAudioFileProperty_FileMaxPacketSize: The file data format’s maximum packet size, in bytes.
	KExtAudioFileProperty_FileMaxPacketSize KExtAudioFileProperty = 'f'<<24 | 'm'<<16 | 'p'<<8 | 's' // 'fmps'
	// KExtAudioFileProperty_IOBuffer: An audio data buffer.
	KExtAudioFileProperty_IOBuffer KExtAudioFileProperty = 'i'<<24 | 'o'<<16 | 'b'<<8 | 'f' // 'iobf'
	// KExtAudioFileProperty_IOBufferSizeBytes: The size of the buffer that the extended audio file object’s associated audio converter uses to read or write the associated audio file.
	KExtAudioFileProperty_IOBufferSizeBytes KExtAudioFileProperty = 'i'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'iobs'
	// KExtAudioFileProperty_PacketTable: This property can be used to override the priming and remainder information in an audio file, and also to retrieve the current priming and remainder frames information for an extended audio file object.
	KExtAudioFileProperty_PacketTable KExtAudioFileProperty = 'x'<<24 | 'p'<<16 | 't'<<8 | 'i' // 'xpti'
)

func (e KExtAudioFileProperty) String() string {
	switch e {
	case KExtAudioFileProperty_AudioConverter:
		return "KExtAudioFileProperty_AudioConverter"
	case KExtAudioFileProperty_AudioFile:
		return "KExtAudioFileProperty_AudioFile"
	case KExtAudioFileProperty_ClientChannelLayout:
		return "KExtAudioFileProperty_ClientChannelLayout"
	case KExtAudioFileProperty_ClientDataFormat:
		return "KExtAudioFileProperty_ClientDataFormat"
	case KExtAudioFileProperty_ClientMaxPacketSize:
		return "KExtAudioFileProperty_ClientMaxPacketSize"
	case KExtAudioFileProperty_CodecManufacturer:
		return "KExtAudioFileProperty_CodecManufacturer"
	case KExtAudioFileProperty_ConverterConfig:
		return "KExtAudioFileProperty_ConverterConfig"
	case KExtAudioFileProperty_FileChannelLayout:
		return "KExtAudioFileProperty_FileChannelLayout"
	case KExtAudioFileProperty_FileDataFormat:
		return "KExtAudioFileProperty_FileDataFormat"
	case KExtAudioFileProperty_FileLengthFrames:
		return "KExtAudioFileProperty_FileLengthFrames"
	case KExtAudioFileProperty_FileMaxPacketSize:
		return "KExtAudioFileProperty_FileMaxPacketSize"
	case KExtAudioFileProperty_IOBuffer:
		return "KExtAudioFileProperty_IOBuffer"
	case KExtAudioFileProperty_IOBufferSizeBytes:
		return "KExtAudioFileProperty_IOBufferSizeBytes"
	case KExtAudioFileProperty_PacketTable:
		return "KExtAudioFileProperty_PacketTable"
	default:
		return fmt.Sprintf("KExtAudioFileProperty(%d)", e)
	}
}

type KHighShelfParam uint32

const (
	KHighShelfParam_CutOffFrequency KHighShelfParam = 0
	KHighShelfParam_Gain            KHighShelfParam = 1
)

func (e KHighShelfParam) String() string {
	switch e {
	case KHighShelfParam_CutOffFrequency:
		return "KHighShelfParam_CutOffFrequency"
	case KHighShelfParam_Gain:
		return "KHighShelfParam_Gain"
	default:
		return fmt.Sprintf("KHighShelfParam(%d)", e)
	}
}

type KHint uint32

const (
	KHintAdvanced KHint = 1
	KHintBasic    KHint = 0
	KHintHidden   KHint = 2
)

func (e KHint) String() string {
	switch e {
	case KHintAdvanced:
		return "KHintAdvanced"
	case KHintBasic:
		return "KHintBasic"
	case KHintHidden:
		return "KHintHidden"
	default:
		return fmt.Sprintf("KHint(%d)", e)
	}
}

type KHipassParam uint32

const (
	KHipassParam_CutoffFrequency KHipassParam = 0
	KHipassParam_Resonance       KHipassParam = 1
)

func (e KHipassParam) String() string {
	switch e {
	case KHipassParam_CutoffFrequency:
		return "KHipassParam_CutoffFrequency"
	case KHipassParam_Resonance:
		return "KHipassParam_Resonance"
	default:
		return fmt.Sprintf("KHipassParam(%d)", e)
	}
}

type KInstrumentType uint32

const (
	KInstrumentType_AUPreset  KInstrumentType = 2
	KInstrumentType_Audiofile KInstrumentType = 3
	KInstrumentType_DLSPreset KInstrumentType = 1
	KInstrumentType_EXS24     KInstrumentType = 4
	KInstrumentType_SF2Preset KInstrumentType = 1
)

func (e KInstrumentType) String() string {
	switch e {
	case KInstrumentType_AUPreset:
		return "KInstrumentType_AUPreset"
	case KInstrumentType_Audiofile:
		return "KInstrumentType_Audiofile"
	case KInstrumentType_DLSPreset:
		return "KInstrumentType_DLSPreset"
	case KInstrumentType_EXS24:
		return "KInstrumentType_EXS24"
	default:
		return fmt.Sprintf("KInstrumentType(%d)", e)
	}
}

type KLimiterParam uint32

const (
	KLimiterParam_AttackTime KLimiterParam = 0
	KLimiterParam_DecayTime  KLimiterParam = 1
	KLimiterParam_PreGain    KLimiterParam = 2
)

func (e KLimiterParam) String() string {
	switch e {
	case KLimiterParam_AttackTime:
		return "KLimiterParam_AttackTime"
	case KLimiterParam_DecayTime:
		return "KLimiterParam_DecayTime"
	case KLimiterParam_PreGain:
		return "KLimiterParam_PreGain"
	default:
		return fmt.Sprintf("KLimiterParam(%d)", e)
	}
}

type KLowPassParam uint32

const (
	KLowPassParam_CutoffFrequency KLowPassParam = 0
	KLowPassParam_Resonance       KLowPassParam = 1
)

func (e KLowPassParam) String() string {
	switch e {
	case KLowPassParam_CutoffFrequency:
		return "KLowPassParam_CutoffFrequency"
	case KLowPassParam_Resonance:
		return "KLowPassParam_Resonance"
	default:
		return fmt.Sprintf("KLowPassParam(%d)", e)
	}
}

type KMatrixMixerParam uint32

const (
	KMatrixMixerParam_Enable                  KMatrixMixerParam = 1
	KMatrixMixerParam_PostAveragePower        KMatrixMixerParam = 3000
	KMatrixMixerParam_PostAveragePowerLinear  KMatrixMixerParam = 7000
	KMatrixMixerParam_PostPeakHoldLevel       KMatrixMixerParam = 4000
	KMatrixMixerParam_PostPeakHoldLevelLinear KMatrixMixerParam = 8000
	KMatrixMixerParam_PreAveragePower         KMatrixMixerParam = 1000
	KMatrixMixerParam_PreAveragePowerLinear   KMatrixMixerParam = 5000
	KMatrixMixerParam_PrePeakHoldLevel        KMatrixMixerParam = 2000
	KMatrixMixerParam_PrePeakHoldLevelLinear  KMatrixMixerParam = 6000
	KMatrixMixerParam_Volume                  KMatrixMixerParam = 0
)

func (e KMatrixMixerParam) String() string {
	switch e {
	case KMatrixMixerParam_Enable:
		return "KMatrixMixerParam_Enable"
	case KMatrixMixerParam_PostAveragePower:
		return "KMatrixMixerParam_PostAveragePower"
	case KMatrixMixerParam_PostAveragePowerLinear:
		return "KMatrixMixerParam_PostAveragePowerLinear"
	case KMatrixMixerParam_PostPeakHoldLevel:
		return "KMatrixMixerParam_PostPeakHoldLevel"
	case KMatrixMixerParam_PostPeakHoldLevelLinear:
		return "KMatrixMixerParam_PostPeakHoldLevelLinear"
	case KMatrixMixerParam_PreAveragePower:
		return "KMatrixMixerParam_PreAveragePower"
	case KMatrixMixerParam_PreAveragePowerLinear:
		return "KMatrixMixerParam_PreAveragePowerLinear"
	case KMatrixMixerParam_PrePeakHoldLevel:
		return "KMatrixMixerParam_PrePeakHoldLevel"
	case KMatrixMixerParam_PrePeakHoldLevelLinear:
		return "KMatrixMixerParam_PrePeakHoldLevelLinear"
	case KMatrixMixerParam_Volume:
		return "KMatrixMixerParam_Volume"
	default:
		return fmt.Sprintf("KMatrixMixerParam(%d)", e)
	}
}

type KMultiChannelMixerParam uint32

const (
	// KMultiChannelMixerParam_Enable: Enables a channel.
	KMultiChannelMixerParam_Enable KMultiChannelMixerParam = 1
	// KMultiChannelMixerParam_Pan: The panning value for the mixer.
	KMultiChannelMixerParam_Pan KMultiChannelMixerParam = 2
	// KMultiChannelMixerParam_PostAveragePower: The average power level of the channel, after the mixer, in decibels.
	KMultiChannelMixerParam_PostAveragePower KMultiChannelMixerParam = 3000
	// KMultiChannelMixerParam_PostPeakHoldLevel: The peak hold level of the channel, after the mixer, in decibels.
	KMultiChannelMixerParam_PostPeakHoldLevel KMultiChannelMixerParam = 4000
	// KMultiChannelMixerParam_PreAveragePower: The average power level of the channel, prior to the mixer, in decibels.
	KMultiChannelMixerParam_PreAveragePower KMultiChannelMixerParam = 1000
	// KMultiChannelMixerParam_PrePeakHoldLevel: The peak hold level of the channel, prior to the mixer, in decibels.
	KMultiChannelMixerParam_PrePeakHoldLevel KMultiChannelMixerParam = 2000
	// KMultiChannelMixerParam_Volume: The linear gain of the channel.
	KMultiChannelMixerParam_Volume KMultiChannelMixerParam = 0
)

func (e KMultiChannelMixerParam) String() string {
	switch e {
	case KMultiChannelMixerParam_Enable:
		return "KMultiChannelMixerParam_Enable"
	case KMultiChannelMixerParam_Pan:
		return "KMultiChannelMixerParam_Pan"
	case KMultiChannelMixerParam_PostAveragePower:
		return "KMultiChannelMixerParam_PostAveragePower"
	case KMultiChannelMixerParam_PostPeakHoldLevel:
		return "KMultiChannelMixerParam_PostPeakHoldLevel"
	case KMultiChannelMixerParam_PreAveragePower:
		return "KMultiChannelMixerParam_PreAveragePower"
	case KMultiChannelMixerParam_PrePeakHoldLevel:
		return "KMultiChannelMixerParam_PrePeakHoldLevel"
	case KMultiChannelMixerParam_Volume:
		return "KMultiChannelMixerParam_Volume"
	default:
		return fmt.Sprintf("KMultiChannelMixerParam(%d)", e)
	}
}

type KMultibandCompressorParam uint32

const (
	KMultibandCompressorParam_AttackTime         KMultibandCompressorParam = 13
	KMultibandCompressorParam_CompressionAmount1 KMultibandCompressorParam = 1000
	KMultibandCompressorParam_CompressionAmount2 KMultibandCompressorParam = 2000
	KMultibandCompressorParam_CompressionAmount3 KMultibandCompressorParam = 3000
	KMultibandCompressorParam_CompressionAmount4 KMultibandCompressorParam = 4000
	KMultibandCompressorParam_Crossover1         KMultibandCompressorParam = 2
	KMultibandCompressorParam_Crossover2         KMultibandCompressorParam = 3
	KMultibandCompressorParam_Crossover3         KMultibandCompressorParam = 4
	KMultibandCompressorParam_EQ1                KMultibandCompressorParam = 15
	KMultibandCompressorParam_EQ2                KMultibandCompressorParam = 16
	KMultibandCompressorParam_EQ3                KMultibandCompressorParam = 17
	KMultibandCompressorParam_EQ4                KMultibandCompressorParam = 18
	KMultibandCompressorParam_Headroom1          KMultibandCompressorParam = 9
	KMultibandCompressorParam_Headroom2          KMultibandCompressorParam = 10
	KMultibandCompressorParam_Headroom3          KMultibandCompressorParam = 11
	KMultibandCompressorParam_Headroom4          KMultibandCompressorParam = 12
	KMultibandCompressorParam_InputAmplitude1    KMultibandCompressorParam = 5000
	KMultibandCompressorParam_InputAmplitude2    KMultibandCompressorParam = 6000
	KMultibandCompressorParam_InputAmplitude3    KMultibandCompressorParam = 7000
	KMultibandCompressorParam_InputAmplitude4    KMultibandCompressorParam = 8000
	KMultibandCompressorParam_OutputAmplitude1   KMultibandCompressorParam = 9000
	KMultibandCompressorParam_OutputAmplitude2   KMultibandCompressorParam = 10000
	KMultibandCompressorParam_OutputAmplitude3   KMultibandCompressorParam = 11000
	KMultibandCompressorParam_OutputAmplitude4   KMultibandCompressorParam = 12000
	KMultibandCompressorParam_Postgain           KMultibandCompressorParam = 1
	KMultibandCompressorParam_Pregain            KMultibandCompressorParam = 0
	KMultibandCompressorParam_ReleaseTime        KMultibandCompressorParam = 14
	KMultibandCompressorParam_Threshold1         KMultibandCompressorParam = 5
	KMultibandCompressorParam_Threshold2         KMultibandCompressorParam = 6
	KMultibandCompressorParam_Threshold3         KMultibandCompressorParam = 7
	KMultibandCompressorParam_Threshold4         KMultibandCompressorParam = 8
)

func (e KMultibandCompressorParam) String() string {
	switch e {
	case KMultibandCompressorParam_AttackTime:
		return "KMultibandCompressorParam_AttackTime"
	case KMultibandCompressorParam_CompressionAmount1:
		return "KMultibandCompressorParam_CompressionAmount1"
	case KMultibandCompressorParam_CompressionAmount2:
		return "KMultibandCompressorParam_CompressionAmount2"
	case KMultibandCompressorParam_CompressionAmount3:
		return "KMultibandCompressorParam_CompressionAmount3"
	case KMultibandCompressorParam_CompressionAmount4:
		return "KMultibandCompressorParam_CompressionAmount4"
	case KMultibandCompressorParam_Crossover1:
		return "KMultibandCompressorParam_Crossover1"
	case KMultibandCompressorParam_Crossover2:
		return "KMultibandCompressorParam_Crossover2"
	case KMultibandCompressorParam_Crossover3:
		return "KMultibandCompressorParam_Crossover3"
	case KMultibandCompressorParam_EQ1:
		return "KMultibandCompressorParam_EQ1"
	case KMultibandCompressorParam_EQ2:
		return "KMultibandCompressorParam_EQ2"
	case KMultibandCompressorParam_EQ3:
		return "KMultibandCompressorParam_EQ3"
	case KMultibandCompressorParam_EQ4:
		return "KMultibandCompressorParam_EQ4"
	case KMultibandCompressorParam_Headroom1:
		return "KMultibandCompressorParam_Headroom1"
	case KMultibandCompressorParam_Headroom2:
		return "KMultibandCompressorParam_Headroom2"
	case KMultibandCompressorParam_Headroom3:
		return "KMultibandCompressorParam_Headroom3"
	case KMultibandCompressorParam_Headroom4:
		return "KMultibandCompressorParam_Headroom4"
	case KMultibandCompressorParam_InputAmplitude1:
		return "KMultibandCompressorParam_InputAmplitude1"
	case KMultibandCompressorParam_InputAmplitude2:
		return "KMultibandCompressorParam_InputAmplitude2"
	case KMultibandCompressorParam_InputAmplitude3:
		return "KMultibandCompressorParam_InputAmplitude3"
	case KMultibandCompressorParam_InputAmplitude4:
		return "KMultibandCompressorParam_InputAmplitude4"
	case KMultibandCompressorParam_OutputAmplitude1:
		return "KMultibandCompressorParam_OutputAmplitude1"
	case KMultibandCompressorParam_OutputAmplitude2:
		return "KMultibandCompressorParam_OutputAmplitude2"
	case KMultibandCompressorParam_OutputAmplitude3:
		return "KMultibandCompressorParam_OutputAmplitude3"
	case KMultibandCompressorParam_OutputAmplitude4:
		return "KMultibandCompressorParam_OutputAmplitude4"
	case KMultibandCompressorParam_Postgain:
		return "KMultibandCompressorParam_Postgain"
	case KMultibandCompressorParam_Pregain:
		return "KMultibandCompressorParam_Pregain"
	case KMultibandCompressorParam_ReleaseTime:
		return "KMultibandCompressorParam_ReleaseTime"
	case KMultibandCompressorParam_Threshold1:
		return "KMultibandCompressorParam_Threshold1"
	case KMultibandCompressorParam_Threshold2:
		return "KMultibandCompressorParam_Threshold2"
	case KMultibandCompressorParam_Threshold3:
		return "KMultibandCompressorParam_Threshold3"
	case KMultibandCompressorParam_Threshold4:
		return "KMultibandCompressorParam_Threshold4"
	default:
		return fmt.Sprintf("KMultibandCompressorParam(%d)", e)
	}
}

type KMultibandFilter uint32

const (
	KMultibandFilter_Bandwidth1     KMultibandFilter = 5
	KMultibandFilter_Bandwidth2     KMultibandFilter = 8
	KMultibandFilter_Bandwidth3     KMultibandFilter = 11
	KMultibandFilter_CenterFreq1    KMultibandFilter = 3
	KMultibandFilter_CenterFreq2    KMultibandFilter = 6
	KMultibandFilter_CenterFreq3    KMultibandFilter = 9
	KMultibandFilter_CenterGain1    KMultibandFilter = 4
	KMultibandFilter_CenterGain2    KMultibandFilter = 7
	KMultibandFilter_CenterGain3    KMultibandFilter = 10
	KMultibandFilter_HighFilterType KMultibandFilter = 12
	KMultibandFilter_HighFrequency  KMultibandFilter = 13
	KMultibandFilter_HighGain       KMultibandFilter = 14
	KMultibandFilter_LowFilterType  KMultibandFilter = 0
	KMultibandFilter_LowFrequency   KMultibandFilter = 1
	KMultibandFilter_LowGain        KMultibandFilter = 2
)

func (e KMultibandFilter) String() string {
	switch e {
	case KMultibandFilter_Bandwidth1:
		return "KMultibandFilter_Bandwidth1"
	case KMultibandFilter_Bandwidth2:
		return "KMultibandFilter_Bandwidth2"
	case KMultibandFilter_Bandwidth3:
		return "KMultibandFilter_Bandwidth3"
	case KMultibandFilter_CenterFreq1:
		return "KMultibandFilter_CenterFreq1"
	case KMultibandFilter_CenterFreq2:
		return "KMultibandFilter_CenterFreq2"
	case KMultibandFilter_CenterFreq3:
		return "KMultibandFilter_CenterFreq3"
	case KMultibandFilter_CenterGain1:
		return "KMultibandFilter_CenterGain1"
	case KMultibandFilter_CenterGain2:
		return "KMultibandFilter_CenterGain2"
	case KMultibandFilter_CenterGain3:
		return "KMultibandFilter_CenterGain3"
	case KMultibandFilter_HighFilterType:
		return "KMultibandFilter_HighFilterType"
	case KMultibandFilter_HighFrequency:
		return "KMultibandFilter_HighFrequency"
	case KMultibandFilter_HighGain:
		return "KMultibandFilter_HighGain"
	case KMultibandFilter_LowFilterType:
		return "KMultibandFilter_LowFilterType"
	case KMultibandFilter_LowFrequency:
		return "KMultibandFilter_LowFrequency"
	case KMultibandFilter_LowGain:
		return "KMultibandFilter_LowGain"
	default:
		return fmt.Sprintf("KMultibandFilter(%d)", e)
	}
}

type KMusicDevice uint32

const (
	KMusicDeviceMIDIEventListSelect     KMusicDevice = 0x107
	KMusicDeviceMIDIEventSelect         KMusicDevice = 0x101
	KMusicDevicePrepareInstrumentSelect KMusicDevice = 0x103
	KMusicDeviceRange                   KMusicDevice = 0x100
	KMusicDeviceReleaseInstrumentSelect KMusicDevice = 0x104
	KMusicDeviceStartNoteSelect         KMusicDevice = 0x105
	KMusicDeviceStopNoteSelect          KMusicDevice = 0x106
	KMusicDeviceSysExSelect             KMusicDevice = 0x102
)

func (e KMusicDevice) String() string {
	switch e {
	case KMusicDeviceMIDIEventListSelect:
		return "KMusicDeviceMIDIEventListSelect"
	case KMusicDeviceMIDIEventSelect:
		return "KMusicDeviceMIDIEventSelect"
	case KMusicDevicePrepareInstrumentSelect:
		return "KMusicDevicePrepareInstrumentSelect"
	case KMusicDeviceRange:
		return "KMusicDeviceRange"
	case KMusicDeviceReleaseInstrumentSelect:
		return "KMusicDeviceReleaseInstrumentSelect"
	case KMusicDeviceStartNoteSelect:
		return "KMusicDeviceStartNoteSelect"
	case KMusicDeviceStopNoteSelect:
		return "KMusicDeviceStopNoteSelect"
	case KMusicDeviceSysExSelect:
		return "KMusicDeviceSysExSelect"
	default:
		return fmt.Sprintf("KMusicDevice(%d)", e)
	}
}

type KMusicDeviceParam uint32

const (
	KMusicDeviceParam_ReverbVolume KMusicDeviceParam = 2
	KMusicDeviceParam_Tuning       KMusicDeviceParam = 0
	KMusicDeviceParam_Volume       KMusicDeviceParam = 1
)

func (e KMusicDeviceParam) String() string {
	switch e {
	case KMusicDeviceParam_ReverbVolume:
		return "KMusicDeviceParam_ReverbVolume"
	case KMusicDeviceParam_Tuning:
		return "KMusicDeviceParam_Tuning"
	case KMusicDeviceParam_Volume:
		return "KMusicDeviceParam_Volume"
	default:
		return fmt.Sprintf("KMusicDeviceParam(%d)", e)
	}
}

type KMusicDeviceSampleFrameMask uint32

const (
	KMusicDeviceSampleFrameMask_IsScheduled  KMusicDeviceSampleFrameMask = 0x1000000
	KMusicDeviceSampleFrameMask_SampleOffset KMusicDeviceSampleFrameMask = 0xffffff
)

func (e KMusicDeviceSampleFrameMask) String() string {
	switch e {
	case KMusicDeviceSampleFrameMask_IsScheduled:
		return "KMusicDeviceSampleFrameMask_IsScheduled"
	case KMusicDeviceSampleFrameMask_SampleOffset:
		return "KMusicDeviceSampleFrameMask_SampleOffset"
	default:
		return fmt.Sprintf("KMusicDeviceSampleFrameMask(%d)", e)
	}
}

type KMusicEventType uint32

const (
	// KMusicEventType_AUPreset: An event containing an audio unit user preset dictionary.
	KMusicEventType_AUPreset KMusicEventType = 10
	// KMusicEventType_ExtendedNote: A non-MIDI music event with variable number of parameters.
	KMusicEventType_ExtendedNote KMusicEventType = 1
	// KMusicEventType_ExtendedTempo: An event that signals a change in tempo, in beats-per-minute.
	KMusicEventType_ExtendedTempo KMusicEventType = 3
	// KMusicEventType_MIDIChannelMessage: A MIDI channel message, other than note control.
	KMusicEventType_MIDIChannelMessage KMusicEventType = 7
	// KMusicEventType_MIDINoteMessage: A MIDI note-on message, including duration.
	KMusicEventType_MIDINoteMessage KMusicEventType = 6
	// KMusicEventType_MIDIRawData: MIDI system-exclusive data.
	KMusicEventType_MIDIRawData KMusicEventType = 8
	// KMusicEventType_Meta: A standard MIDI file metaevent.
	KMusicEventType_Meta KMusicEventType = 5
	// KMusicEventType_NULL: A null music event.
	KMusicEventType_NULL KMusicEventType = 0
	// KMusicEventType_Parameter: An audio unit parameter event.
	KMusicEventType_Parameter KMusicEventType = 9
	// KMusicEventType_User: User-defined data.
	KMusicEventType_User KMusicEventType = 4
)

func (e KMusicEventType) String() string {
	switch e {
	case KMusicEventType_AUPreset:
		return "KMusicEventType_AUPreset"
	case KMusicEventType_ExtendedNote:
		return "KMusicEventType_ExtendedNote"
	case KMusicEventType_ExtendedTempo:
		return "KMusicEventType_ExtendedTempo"
	case KMusicEventType_MIDIChannelMessage:
		return "KMusicEventType_MIDIChannelMessage"
	case KMusicEventType_MIDINoteMessage:
		return "KMusicEventType_MIDINoteMessage"
	case KMusicEventType_MIDIRawData:
		return "KMusicEventType_MIDIRawData"
	case KMusicEventType_Meta:
		return "KMusicEventType_Meta"
	case KMusicEventType_NULL:
		return "KMusicEventType_NULL"
	case KMusicEventType_Parameter:
		return "KMusicEventType_Parameter"
	case KMusicEventType_User:
		return "KMusicEventType_User"
	default:
		return fmt.Sprintf("KMusicEventType(%d)", e)
	}
}

type KMusicNoteEvent uint32

const (
	KMusicNoteEvent_Unused             KMusicNoteEvent = 0xffffffff
	KMusicNoteEvent_UseGroupInstrument KMusicNoteEvent = 0xffffffff
)

func (e KMusicNoteEvent) String() string {
	switch e {
	case KMusicNoteEvent_Unused:
		return "KMusicNoteEvent_Unused"
	default:
		return fmt.Sprintf("KMusicNoteEvent(%d)", e)
	}
}

type KNewTimePitchParam uint32

const (
	KNewTimePitchParam_EnableSpectralCoherence     KNewTimePitchParam = 6
	KNewTimePitchParam_EnableTransientPreservation KNewTimePitchParam = 7
	KNewTimePitchParam_Pitch                       KNewTimePitchParam = 1
	KNewTimePitchParam_Rate                        KNewTimePitchParam = 0
	KNewTimePitchParam_Smoothness                  KNewTimePitchParam = 4
	// Deprecated: use KNewTimePitchParam_EnableSpectralCoherence.
	KNewTimePitchParam_EnablePeakLocking KNewTimePitchParam = 6
	// Deprecated: use KNewTimePitchParam_Smoothness.
	KNewTimePitchParam_Overlap KNewTimePitchParam = 4
)

func (e KNewTimePitchParam) String() string {
	switch e {
	case KNewTimePitchParam_EnableSpectralCoherence:
		return "KNewTimePitchParam_EnableSpectralCoherence"
	case KNewTimePitchParam_EnableTransientPreservation:
		return "KNewTimePitchParam_EnableTransientPreservation"
	case KNewTimePitchParam_Pitch:
		return "KNewTimePitchParam_Pitch"
	case KNewTimePitchParam_Rate:
		return "KNewTimePitchParam_Rate"
	case KNewTimePitchParam_Smoothness:
		return "KNewTimePitchParam_Smoothness"
	default:
		return fmt.Sprintf("KNewTimePitchParam(%d)", e)
	}
}

const KNumberOfResponseFrequencies uint32 = 1024

type KOfflinePreflight uint32

const (
	KOfflinePreflight_NotRequired KOfflinePreflight = 0
	KOfflinePreflight_Optional    KOfflinePreflight = 1
	KOfflinePreflight_Required    KOfflinePreflight = 2
)

func (e KOfflinePreflight) String() string {
	switch e {
	case KOfflinePreflight_NotRequired:
		return "KOfflinePreflight_NotRequired"
	case KOfflinePreflight_Optional:
		return "KOfflinePreflight_Optional"
	case KOfflinePreflight_Required:
		return "KOfflinePreflight_Required"
	default:
		return fmt.Sprintf("KOfflinePreflight(%d)", e)
	}
}

type KOtherPluginFormat uint32

const (
	KOtherPluginFormat_AU        KOtherPluginFormat = 3
	KOtherPluginFormat_Undefined KOtherPluginFormat = 0
	KOtherPluginFormat_kMAS      KOtherPluginFormat = 1
	KOtherPluginFormat_kVST      KOtherPluginFormat = 2
)

func (e KOtherPluginFormat) String() string {
	switch e {
	case KOtherPluginFormat_AU:
		return "KOtherPluginFormat_AU"
	case KOtherPluginFormat_Undefined:
		return "KOtherPluginFormat_Undefined"
	case KOtherPluginFormat_kMAS:
		return "KOtherPluginFormat_kMAS"
	case KOtherPluginFormat_kVST:
		return "KOtherPluginFormat_kVST"
	default:
		return fmt.Sprintf("KOtherPluginFormat(%d)", e)
	}
}

type KPannerParam uint32

const (
	KPannerParam_Azimuth     KPannerParam = 1
	KPannerParam_CoordScale  KPannerParam = 4
	KPannerParam_Distance    KPannerParam = 3
	KPannerParam_Elevation   KPannerParam = 2
	KPannerParam_Gain        KPannerParam = 0
	KPannerParam_RefDistance KPannerParam = 5
)

func (e KPannerParam) String() string {
	switch e {
	case KPannerParam_Azimuth:
		return "KPannerParam_Azimuth"
	case KPannerParam_CoordScale:
		return "KPannerParam_CoordScale"
	case KPannerParam_Distance:
		return "KPannerParam_Distance"
	case KPannerParam_Elevation:
		return "KPannerParam_Elevation"
	case KPannerParam_Gain:
		return "KPannerParam_Gain"
	case KPannerParam_RefDistance:
		return "KPannerParam_RefDistance"
	default:
		return fmt.Sprintf("KPannerParam(%d)", e)
	}
}

type KParametricEQParam uint32

const (
	KParametricEQParam_CenterFreq KParametricEQParam = 0
	KParametricEQParam_Gain       KParametricEQParam = 2
	KParametricEQParam_Q          KParametricEQParam = 1
)

func (e KParametricEQParam) String() string {
	switch e {
	case KParametricEQParam_CenterFreq:
		return "KParametricEQParam_CenterFreq"
	case KParametricEQParam_Gain:
		return "KParametricEQParam_Gain"
	case KParametricEQParam_Q:
		return "KParametricEQParam_Q"
	default:
		return fmt.Sprintf("KParametricEQParam(%d)", e)
	}
}

type KProgramTargetLevel uint32

const (
	KProgramTargetLevel_Minus20dB KProgramTargetLevel = 3
	KProgramTargetLevel_Minus23dB KProgramTargetLevel = 2
	KProgramTargetLevel_Minus31dB KProgramTargetLevel = 1
	KProgramTargetLevel_None      KProgramTargetLevel = 0
)

func (e KProgramTargetLevel) String() string {
	switch e {
	case KProgramTargetLevel_Minus20dB:
		return "KProgramTargetLevel_Minus20dB"
	case KProgramTargetLevel_Minus23dB:
		return "KProgramTargetLevel_Minus23dB"
	case KProgramTargetLevel_Minus31dB:
		return "KProgramTargetLevel_Minus31dB"
	case KProgramTargetLevel_None:
		return "KProgramTargetLevel_None"
	default:
		return fmt.Sprintf("KProgramTargetLevel(%d)", e)
	}
}

type KRandomParam uint32

const (
	KRandomParam_BoundA KRandomParam = 0
	KRandomParam_BoundB KRandomParam = 1
	KRandomParam_Curve  KRandomParam = 2
)

func (e KRandomParam) String() string {
	switch e {
	case KRandomParam_BoundA:
		return "KRandomParam_BoundA"
	case KRandomParam_BoundB:
		return "KRandomParam_BoundB"
	case KRandomParam_Curve:
		return "KRandomParam_Curve"
	default:
		return fmt.Sprintf("KRandomParam(%d)", e)
	}
}

type KRenderQuality uint32

const (
	KRenderQuality_High   KRenderQuality = 96
	KRenderQuality_Low    KRenderQuality = 32
	KRenderQuality_Max    KRenderQuality = 127
	KRenderQuality_Medium KRenderQuality = 64
	KRenderQuality_Min    KRenderQuality = 0
)

func (e KRenderQuality) String() string {
	switch e {
	case KRenderQuality_High:
		return "KRenderQuality_High"
	case KRenderQuality_Low:
		return "KRenderQuality_Low"
	case KRenderQuality_Max:
		return "KRenderQuality_Max"
	case KRenderQuality_Medium:
		return "KRenderQuality_Medium"
	case KRenderQuality_Min:
		return "KRenderQuality_Min"
	default:
		return fmt.Sprintf("KRenderQuality(%d)", e)
	}
}

type KReverb2Param uint32

const (
	KReverb2Param_DecayTimeAt0Hz       KReverb2Param = 4
	KReverb2Param_DecayTimeAtNyquist   KReverb2Param = 5
	KReverb2Param_DryWetMix            KReverb2Param = 0
	KReverb2Param_Gain                 KReverb2Param = 1
	KReverb2Param_LegacyMode           KReverb2Param = 0
	KReverb2Param_MaxDelayTime         KReverb2Param = 3
	KReverb2Param_MinDelayTime         KReverb2Param = 2
	KReverb2Param_RandomizeReflections KReverb2Param = 6
)

func (e KReverb2Param) String() string {
	switch e {
	case KReverb2Param_DecayTimeAt0Hz:
		return "KReverb2Param_DecayTimeAt0Hz"
	case KReverb2Param_DecayTimeAtNyquist:
		return "KReverb2Param_DecayTimeAtNyquist"
	case KReverb2Param_DryWetMix:
		return "KReverb2Param_DryWetMix"
	case KReverb2Param_Gain:
		return "KReverb2Param_Gain"
	case KReverb2Param_MaxDelayTime:
		return "KReverb2Param_MaxDelayTime"
	case KReverb2Param_MinDelayTime:
		return "KReverb2Param_MinDelayTime"
	case KReverb2Param_RandomizeReflections:
		return "KReverb2Param_RandomizeReflections"
	default:
		return fmt.Sprintf("KReverb2Param(%d)", e)
	}
}

type KReverbParam uint32

const (
	KReverbParam_DryWetMix       KReverbParam = 0
	KReverbParam_LargeBrightness KReverbParam = 10
	KReverbParam_LargeDelay      KReverbParam = 5
	KReverbParam_LargeDelayRange KReverbParam = 8
	KReverbParam_LargeDensity    KReverbParam = 7
	KReverbParam_LargeSize       KReverbParam = 3
	KReverbParam_ModulationDepth KReverbParam = 13
	KReverbParam_ModulationRate  KReverbParam = 12
	KReverbParam_PreDelay        KReverbParam = 4
	KReverbParam_SmallBrightness KReverbParam = 9
	KReverbParam_SmallDelayRange KReverbParam = 11
	KReverbParam_SmallDensity    KReverbParam = 6
	KReverbParam_SmallLargeMix   KReverbParam = 1
	KReverbParam_SmallSize       KReverbParam = 2
)

func (e KReverbParam) String() string {
	switch e {
	case KReverbParam_DryWetMix:
		return "KReverbParam_DryWetMix"
	case KReverbParam_LargeBrightness:
		return "KReverbParam_LargeBrightness"
	case KReverbParam_LargeDelay:
		return "KReverbParam_LargeDelay"
	case KReverbParam_LargeDelayRange:
		return "KReverbParam_LargeDelayRange"
	case KReverbParam_LargeDensity:
		return "KReverbParam_LargeDensity"
	case KReverbParam_LargeSize:
		return "KReverbParam_LargeSize"
	case KReverbParam_ModulationDepth:
		return "KReverbParam_ModulationDepth"
	case KReverbParam_ModulationRate:
		return "KReverbParam_ModulationRate"
	case KReverbParam_PreDelay:
		return "KReverbParam_PreDelay"
	case KReverbParam_SmallBrightness:
		return "KReverbParam_SmallBrightness"
	case KReverbParam_SmallDelayRange:
		return "KReverbParam_SmallDelayRange"
	case KReverbParam_SmallDensity:
		return "KReverbParam_SmallDensity"
	case KReverbParam_SmallLargeMix:
		return "KReverbParam_SmallLargeMix"
	case KReverbParam_SmallSize:
		return "KReverbParam_SmallSize"
	default:
		return fmt.Sprintf("KReverbParam(%d)", e)
	}
}

type KRogerBeepParam uint32

const (
	KRogerBeepParam_InGateThreshold      KRogerBeepParam = 0
	KRogerBeepParam_InGateThresholdTime  KRogerBeepParam = 1
	KRogerBeepParam_OutGateThreshold     KRogerBeepParam = 2
	KRogerBeepParam_OutGateThresholdTime KRogerBeepParam = 3
	KRogerBeepParam_RogerGain            KRogerBeepParam = 6
	KRogerBeepParam_RogerType            KRogerBeepParam = 5
	KRogerBeepParam_Sensitivity          KRogerBeepParam = 4
)

func (e KRogerBeepParam) String() string {
	switch e {
	case KRogerBeepParam_InGateThreshold:
		return "KRogerBeepParam_InGateThreshold"
	case KRogerBeepParam_InGateThresholdTime:
		return "KRogerBeepParam_InGateThresholdTime"
	case KRogerBeepParam_OutGateThreshold:
		return "KRogerBeepParam_OutGateThreshold"
	case KRogerBeepParam_OutGateThresholdTime:
		return "KRogerBeepParam_OutGateThresholdTime"
	case KRogerBeepParam_RogerGain:
		return "KRogerBeepParam_RogerGain"
	case KRogerBeepParam_RogerType:
		return "KRogerBeepParam_RogerType"
	case KRogerBeepParam_Sensitivity:
		return "KRogerBeepParam_Sensitivity"
	default:
		return fmt.Sprintf("KRogerBeepParam(%d)", e)
	}
}

type KRoundTripAACParam uint32

const (
	KRoundTripAACParam_BitRate                    KRoundTripAACParam = 1
	KRoundTripAACParam_CompressedFormatSampleRate KRoundTripAACParam = 3
	KRoundTripAACParam_EncodingStrategy           KRoundTripAACParam = 1
	KRoundTripAACParam_Format                     KRoundTripAACParam = 0
	KRoundTripAACParam_Quality                    KRoundTripAACParam = 2
	KRoundTripAACParam_RateOrQuality              KRoundTripAACParam = 2
)

func (e KRoundTripAACParam) String() string {
	switch e {
	case KRoundTripAACParam_BitRate:
		return "KRoundTripAACParam_BitRate"
	case KRoundTripAACParam_CompressedFormatSampleRate:
		return "KRoundTripAACParam_CompressedFormatSampleRate"
	case KRoundTripAACParam_Format:
		return "KRoundTripAACParam_Format"
	case KRoundTripAACParam_Quality:
		return "KRoundTripAACParam_Quality"
	default:
		return fmt.Sprintf("KRoundTripAACParam(%d)", e)
	}
}

type KSequenceTrackProperty uint32

const (
	// KSequenceTrackProperty_AutomatedParameters: Indicates whether or not a music track’s purpose is audio unit parameter automation.
	KSequenceTrackProperty_AutomatedParameters KSequenceTrackProperty = 4
	// KSequenceTrackProperty_LoopInfo: Looping information for a music track.
	KSequenceTrackProperty_LoopInfo KSequenceTrackProperty = 0
	// KSequenceTrackProperty_MuteStatus: The mute/unmute state of a music track.
	KSequenceTrackProperty_MuteStatus KSequenceTrackProperty = 2
	// KSequenceTrackProperty_OffsetTime: A music track’s start time in terms of beat number.
	KSequenceTrackProperty_OffsetTime KSequenceTrackProperty = 1
	// KSequenceTrackProperty_SoloStatus: The solo/unsolo state of a music track.
	KSequenceTrackProperty_SoloStatus KSequenceTrackProperty = 3
	// KSequenceTrackProperty_TimeResolution: The time resolution for a sequence of music events.
	KSequenceTrackProperty_TimeResolution KSequenceTrackProperty = 6
	// KSequenceTrackProperty_TrackLength: The time of the last music event in a music track, plus time required for note fade-outs and so on.
	KSequenceTrackProperty_TrackLength KSequenceTrackProperty = 5
)

func (e KSequenceTrackProperty) String() string {
	switch e {
	case KSequenceTrackProperty_AutomatedParameters:
		return "KSequenceTrackProperty_AutomatedParameters"
	case KSequenceTrackProperty_LoopInfo:
		return "KSequenceTrackProperty_LoopInfo"
	case KSequenceTrackProperty_MuteStatus:
		return "KSequenceTrackProperty_MuteStatus"
	case KSequenceTrackProperty_OffsetTime:
		return "KSequenceTrackProperty_OffsetTime"
	case KSequenceTrackProperty_SoloStatus:
		return "KSequenceTrackProperty_SoloStatus"
	case KSequenceTrackProperty_TimeResolution:
		return "KSequenceTrackProperty_TimeResolution"
	case KSequenceTrackProperty_TrackLength:
		return "KSequenceTrackProperty_TrackLength"
	default:
		return fmt.Sprintf("KSequenceTrackProperty(%d)", e)
	}
}

type KSpatialMixerParam uint32

const (
	KSpatialMixerParam_Azimuth                KSpatialMixerParam = 0
	KSpatialMixerParam_Distance               KSpatialMixerParam = 2
	KSpatialMixerParam_Elevation              KSpatialMixerParam = 1
	KSpatialMixerParam_Enable                 KSpatialMixerParam = 5
	KSpatialMixerParam_Gain                   KSpatialMixerParam = 3
	KSpatialMixerParam_GlobalReverbGain       KSpatialMixerParam = 9
	KSpatialMixerParam_HeadPitch              KSpatialMixerParam = 20
	KSpatialMixerParam_HeadRoll               KSpatialMixerParam = 21
	KSpatialMixerParam_HeadYaw                KSpatialMixerParam = 19
	KSpatialMixerParam_MaxGain                KSpatialMixerParam = 7
	KSpatialMixerParam_MinGain                KSpatialMixerParam = 6
	KSpatialMixerParam_ObstructionAttenuation KSpatialMixerParam = 11
	KSpatialMixerParam_OcclusionAttenuation   KSpatialMixerParam = 10
	KSpatialMixerParam_PlaybackRate           KSpatialMixerParam = 4
	KSpatialMixerParam_ReverbBlend            KSpatialMixerParam = 8
)

func (e KSpatialMixerParam) String() string {
	switch e {
	case KSpatialMixerParam_Azimuth:
		return "KSpatialMixerParam_Azimuth"
	case KSpatialMixerParam_Distance:
		return "KSpatialMixerParam_Distance"
	case KSpatialMixerParam_Elevation:
		return "KSpatialMixerParam_Elevation"
	case KSpatialMixerParam_Enable:
		return "KSpatialMixerParam_Enable"
	case KSpatialMixerParam_Gain:
		return "KSpatialMixerParam_Gain"
	case KSpatialMixerParam_GlobalReverbGain:
		return "KSpatialMixerParam_GlobalReverbGain"
	case KSpatialMixerParam_HeadPitch:
		return "KSpatialMixerParam_HeadPitch"
	case KSpatialMixerParam_HeadRoll:
		return "KSpatialMixerParam_HeadRoll"
	case KSpatialMixerParam_HeadYaw:
		return "KSpatialMixerParam_HeadYaw"
	case KSpatialMixerParam_MaxGain:
		return "KSpatialMixerParam_MaxGain"
	case KSpatialMixerParam_MinGain:
		return "KSpatialMixerParam_MinGain"
	case KSpatialMixerParam_ObstructionAttenuation:
		return "KSpatialMixerParam_ObstructionAttenuation"
	case KSpatialMixerParam_OcclusionAttenuation:
		return "KSpatialMixerParam_OcclusionAttenuation"
	case KSpatialMixerParam_PlaybackRate:
		return "KSpatialMixerParam_PlaybackRate"
	case KSpatialMixerParam_ReverbBlend:
		return "KSpatialMixerParam_ReverbBlend"
	default:
		return fmt.Sprintf("KSpatialMixerParam(%d)", e)
	}
}

type KSpeakerConfiguration uint32

const (
	KSpeakerConfiguration_5_0        KSpeakerConfiguration = 3
	KSpeakerConfiguration_5_1        KSpeakerConfiguration = 3
	KSpeakerConfiguration_HeadPhones KSpeakerConfiguration = 0
	KSpeakerConfiguration_Quad       KSpeakerConfiguration = 2
	KSpeakerConfiguration_Stereo     KSpeakerConfiguration = 1
)

func (e KSpeakerConfiguration) String() string {
	switch e {
	case KSpeakerConfiguration_5_0:
		return "KSpeakerConfiguration_5_0"
	case KSpeakerConfiguration_HeadPhones:
		return "KSpeakerConfiguration_HeadPhones"
	case KSpeakerConfiguration_Quad:
		return "KSpeakerConfiguration_Quad"
	case KSpeakerConfiguration_Stereo:
		return "KSpeakerConfiguration_Stereo"
	default:
		return fmt.Sprintf("KSpeakerConfiguration(%d)", e)
	}
}

type KStereoMixerParam uint32

const (
	KStereoMixerParam_Pan               KStereoMixerParam = 1
	KStereoMixerParam_PostAveragePower  KStereoMixerParam = 3000
	KStereoMixerParam_PostPeakHoldLevel KStereoMixerParam = 4000
	KStereoMixerParam_PreAveragePower   KStereoMixerParam = 1000
	KStereoMixerParam_PrePeakHoldLevel  KStereoMixerParam = 2000
	KStereoMixerParam_Volume            KStereoMixerParam = 0
)

func (e KStereoMixerParam) String() string {
	switch e {
	case KStereoMixerParam_Pan:
		return "KStereoMixerParam_Pan"
	case KStereoMixerParam_PostAveragePower:
		return "KStereoMixerParam_PostAveragePower"
	case KStereoMixerParam_PostPeakHoldLevel:
		return "KStereoMixerParam_PostPeakHoldLevel"
	case KStereoMixerParam_PreAveragePower:
		return "KStereoMixerParam_PreAveragePower"
	case KStereoMixerParam_PrePeakHoldLevel:
		return "KStereoMixerParam_PrePeakHoldLevel"
	case KStereoMixerParam_Volume:
		return "KStereoMixerParam_Volume"
	default:
		return fmt.Sprintf("KStereoMixerParam(%d)", e)
	}
}

type KTimePitchParam uint32

const (
	KTimePitchParam_EffectBlend KTimePitchParam = 2
	KTimePitchParam_Pitch       KTimePitchParam = 1
	KTimePitchParam_Rate        KTimePitchParam = 0
)

func (e KTimePitchParam) String() string {
	switch e {
	case KTimePitchParam_EffectBlend:
		return "KTimePitchParam_EffectBlend"
	case KTimePitchParam_Pitch:
		return "KTimePitchParam_Pitch"
	case KTimePitchParam_Rate:
		return "KTimePitchParam_Rate"
	default:
		return fmt.Sprintf("KTimePitchParam(%d)", e)
	}
}

type KaudioconvertererrFormatnotsupported int32

const (
	KAudioConverterErr_BadPropertySizeError      KaudioconvertererrFormatnotsupported = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KAudioConverterErr_FormatNotSupported        KaudioconvertererrFormatnotsupported = 'f'<<24 | 'm'<<16 | 't'<<8 | '?' // 'fmt?'
	KAudioConverterErr_InputSampleRateOutOfRange KaudioconvertererrFormatnotsupported = '!'<<24 | 'i'<<16 | 's'<<8 | 'r' // '!isr'
	KAudioConverterErr_InvalidInputSize          KaudioconvertererrFormatnotsupported = 'i'<<24 | 'n'<<16 | 's'<<8 | 'z' // 'insz'
	// KAudioConverterErr_InvalidOutputSize: The byte size is not an integer multiple of the frame size.
	KAudioConverterErr_InvalidOutputSize               KaudioconvertererrFormatnotsupported = 'o'<<24 | 't'<<16 | 's'<<8 | 'z' // 'otsz'
	KAudioConverterErr_OperationNotSupported           KaudioconvertererrFormatnotsupported = 0x6f703f3f
	KAudioConverterErr_OutputSampleRateOutOfRange      KaudioconvertererrFormatnotsupported = '!'<<24 | 'o'<<16 | 's'<<8 | 'r' // '!osr'
	KAudioConverterErr_PropertyNotSupported            KaudioconvertererrFormatnotsupported = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	KAudioConverterErr_RequiresPacketDescriptionsError KaudioconvertererrFormatnotsupported = '!'<<24 | 'p'<<16 | 'k'<<8 | 'd' // '!pkd'
	KAudioConverterErr_UnspecifiedError                KaudioconvertererrFormatnotsupported = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
)

func (e KaudioconvertererrFormatnotsupported) String() string {
	switch e {
	case KAudioConverterErr_BadPropertySizeError:
		return "KAudioConverterErr_BadPropertySizeError"
	case KAudioConverterErr_FormatNotSupported:
		return "KAudioConverterErr_FormatNotSupported"
	case KAudioConverterErr_InputSampleRateOutOfRange:
		return "KAudioConverterErr_InputSampleRateOutOfRange"
	case KAudioConverterErr_InvalidInputSize:
		return "KAudioConverterErr_InvalidInputSize"
	case KAudioConverterErr_InvalidOutputSize:
		return "KAudioConverterErr_InvalidOutputSize"
	case KAudioConverterErr_OperationNotSupported:
		return "KAudioConverterErr_OperationNotSupported"
	case KAudioConverterErr_OutputSampleRateOutOfRange:
		return "KAudioConverterErr_OutputSampleRateOutOfRange"
	case KAudioConverterErr_PropertyNotSupported:
		return "KAudioConverterErr_PropertyNotSupported"
	case KAudioConverterErr_RequiresPacketDescriptionsError:
		return "KAudioConverterErr_RequiresPacketDescriptionsError"
	case KAudioConverterErr_UnspecifiedError:
		return "KAudioConverterErr_UnspecifiedError"
	default:
		return fmt.Sprintf("KaudioconvertererrFormatnotsupported(%d)", e)
	}
}

type KaudiofilemarkertypeGeneric uint32

const (
	// KAudioFileMarkerType_Generic: A generic marker.
	KAudioFileMarkerType_Generic KaudiofilemarkertypeGeneric = 0
)

func (e KaudiofilemarkertypeGeneric) String() string {
	switch e {
	case KAudioFileMarkerType_Generic:
		return "KAudioFileMarkerType_Generic"
	default:
		return fmt.Sprintf("KaudiofilemarkertypeGeneric(%d)", e)
	}
}

type KaudiooutputunitpropertyCurrentdevice uint32

const (
	KAudioOutputUnitProperty_ChannelMap KaudiooutputunitpropertyCurrentdevice = 2002
	// KAudioOutputUnitProperty_CurrentDevice: A read/write audio device ID object, of type [AudioDeviceID], valid on the audio unit global scope.
	KAudioOutputUnitProperty_CurrentDevice KaudiooutputunitpropertyCurrentdevice = 2000
	// KAudioOutputUnitProperty_EnableIO: Specifies whether audio I/O is enabled for an I/O unit bus-scope combination.
	KAudioOutputUnitProperty_EnableIO                  KaudiooutputunitpropertyCurrentdevice = 2003
	KAudioOutputUnitProperty_HasIO                     KaudiooutputunitpropertyCurrentdevice = 2006
	KAudioOutputUnitProperty_IntendedSpatialExperience KaudiooutputunitpropertyCurrentdevice = 2016
	// KAudioOutputUnitProperty_IsRunning: Indicates whether an audio unit is running ([TRUE]) or not ([FALSE]).
	KAudioOutputUnitProperty_IsRunning KaudiooutputunitpropertyCurrentdevice = 2001
	// KAudioOutputUnitProperty_OSWorkgroup: The workgroup associated with the audio device underlying this Audio Unit.
	KAudioOutputUnitProperty_OSWorkgroup KaudiooutputunitpropertyCurrentdevice = 2015
	// KAudioOutputUnitProperty_SetInputCallback: A read/write [AURenderCallbackStruct] data structure valid on the audio unit global scope.
	KAudioOutputUnitProperty_SetInputCallback KaudiooutputunitpropertyCurrentdevice = 2005
	// KAudioOutputUnitProperty_StartTime: A write-only [AudioOutputUnitStartAtTimeParams] data structure valid on the audio unit global scope.
	KAudioOutputUnitProperty_StartTime KaudiooutputunitpropertyCurrentdevice = 2004
	// KAudioOutputUnitProperty_StartTimestampsAtZero: A read/write [UInt32] value valid on the audio unit global scope.
	KAudioOutputUnitProperty_StartTimestampsAtZero KaudiooutputunitpropertyCurrentdevice = 2007
)

func (e KaudiooutputunitpropertyCurrentdevice) String() string {
	switch e {
	case KAudioOutputUnitProperty_ChannelMap:
		return "KAudioOutputUnitProperty_ChannelMap"
	case KAudioOutputUnitProperty_CurrentDevice:
		return "KAudioOutputUnitProperty_CurrentDevice"
	case KAudioOutputUnitProperty_EnableIO:
		return "KAudioOutputUnitProperty_EnableIO"
	case KAudioOutputUnitProperty_HasIO:
		return "KAudioOutputUnitProperty_HasIO"
	case KAudioOutputUnitProperty_IntendedSpatialExperience:
		return "KAudioOutputUnitProperty_IntendedSpatialExperience"
	case KAudioOutputUnitProperty_IsRunning:
		return "KAudioOutputUnitProperty_IsRunning"
	case KAudioOutputUnitProperty_OSWorkgroup:
		return "KAudioOutputUnitProperty_OSWorkgroup"
	case KAudioOutputUnitProperty_SetInputCallback:
		return "KAudioOutputUnitProperty_SetInputCallback"
	case KAudioOutputUnitProperty_StartTime:
		return "KAudioOutputUnitProperty_StartTime"
	case KAudioOutputUnitProperty_StartTimestampsAtZero:
		return "KAudioOutputUnitProperty_StartTimestampsAtZero"
	default:
		return fmt.Sprintf("KaudiooutputunitpropertyCurrentdevice(%d)", e)
	}
}

type KaudiooutputunitpropertyMidicallbacks uint32

const (
	KAudioOutputUnitProperty_HostReceivesRemoteControlEvents KaudiooutputunitpropertyMidicallbacks = 2011
	KAudioOutputUnitProperty_HostTransportState              KaudiooutputunitpropertyMidicallbacks = 2013
	KAudioOutputUnitProperty_MIDICallbacks                   KaudiooutputunitpropertyMidicallbacks = 2010
	KAudioOutputUnitProperty_NodeComponentDescription        KaudiooutputunitpropertyMidicallbacks = 2014
	KAudioOutputUnitProperty_RemoteControlToHost             KaudiooutputunitpropertyMidicallbacks = 2012
)

func (e KaudiooutputunitpropertyMidicallbacks) String() string {
	switch e {
	case KAudioOutputUnitProperty_HostReceivesRemoteControlEvents:
		return "KAudioOutputUnitProperty_HostReceivesRemoteControlEvents"
	case KAudioOutputUnitProperty_HostTransportState:
		return "KAudioOutputUnitProperty_HostTransportState"
	case KAudioOutputUnitProperty_MIDICallbacks:
		return "KAudioOutputUnitProperty_MIDICallbacks"
	case KAudioOutputUnitProperty_NodeComponentDescription:
		return "KAudioOutputUnitProperty_NodeComponentDescription"
	case KAudioOutputUnitProperty_RemoteControlToHost:
		return "KAudioOutputUnitProperty_RemoteControlToHost"
	default:
		return fmt.Sprintf("KaudiooutputunitpropertyMidicallbacks(%d)", e)
	}
}

type KaudiosessioncategoryAmbientsound uint32

const (
	// KAudioSessionCategory_AmbientSound: For an app in which sound playback is nonprimary—that is, your app can be used successfully with the sound turned off.
	KAudioSessionCategory_AmbientSound KaudiosessioncategoryAmbientsound = 'a'<<24 | 'm'<<16 | 'b'<<8 | 'i' // 'ambi'
	// KAudioSessionCategory_AudioProcessing: For using an audio hardware codec or signal processor while not playing or recording audio.
	KAudioSessionCategory_AudioProcessing KaudiosessioncategoryAmbientsound = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'c' // 'proc'
	// KAudioSessionCategory_MediaPlayback: For playing recorded music or other sounds that are central to the successful use of your app.
	KAudioSessionCategory_MediaPlayback KaudiosessioncategoryAmbientsound = 'm'<<24 | 'e'<<16 | 'd'<<8 | 'i' // 'medi'
	// KAudioSessionCategory_PlayAndRecord: Allows recording (input) and playback (output) of audio, such as for a VOIP (voice over IP) app.
	KAudioSessionCategory_PlayAndRecord KaudiosessioncategoryAmbientsound = 'p'<<24 | 'l'<<16 | 'a'<<8 | 'r' // 'plar'
	// KAudioSessionCategory_RecordAudio: For recording audio; this category silences playback audio.
	KAudioSessionCategory_RecordAudio KaudiosessioncategoryAmbientsound = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'a' // 'reca'
	// KAudioSessionCategory_SoloAmbientSound: The default category, used unless you set a category with the AudioSessionSetProperty(_:_:_:) function.
	KAudioSessionCategory_SoloAmbientSound KaudiosessioncategoryAmbientsound = 's'<<24 | 'o'<<16 | 'l'<<8 | 'o' // 'solo'
)

func (e KaudiosessioncategoryAmbientsound) String() string {
	switch e {
	case KAudioSessionCategory_AmbientSound:
		return "KAudioSessionCategory_AmbientSound"
	case KAudioSessionCategory_AudioProcessing:
		return "KAudioSessionCategory_AudioProcessing"
	case KAudioSessionCategory_MediaPlayback:
		return "KAudioSessionCategory_MediaPlayback"
	case KAudioSessionCategory_PlayAndRecord:
		return "KAudioSessionCategory_PlayAndRecord"
	case KAudioSessionCategory_RecordAudio:
		return "KAudioSessionCategory_RecordAudio"
	case KAudioSessionCategory_SoloAmbientSound:
		return "KAudioSessionCategory_SoloAmbientSound"
	default:
		return fmt.Sprintf("KaudiosessioncategoryAmbientsound(%d)", e)
	}
}

type KaudiosessioncategoryUserinterfacesoundeffects uint32

const (
	// KAudioSessionCategory_LiveAudio: For live performance of music, such as for an app that simulates a piano.
	KAudioSessionCategory_LiveAudio KaudiosessioncategoryUserinterfacesoundeffects = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'e' // 'live'
	// KAudioSessionCategory_UserInterfaceSoundEffects: For sound effects such as touch feedback, explosions, and so on.
	KAudioSessionCategory_UserInterfaceSoundEffects KaudiosessioncategoryUserinterfacesoundeffects = 'u'<<24 | 'i'<<16 | 'f'<<8 | 'x' // 'uifx'
)

func (e KaudiosessioncategoryUserinterfacesoundeffects) String() string {
	switch e {
	case KAudioSessionCategory_LiveAudio:
		return "KAudioSessionCategory_LiveAudio"
	case KAudioSessionCategory_UserInterfaceSoundEffects:
		return "KAudioSessionCategory_UserInterfaceSoundEffects"
	default:
		return fmt.Sprintf("KaudiosessioncategoryUserinterfacesoundeffects(%d)", e)
	}
}

type KaudiosessioninterruptiontypeShould uint32

const (
	// KAudioSessionInterruptionType_ShouldNotResume: Indicates that the interruption that has just ended was one for which it is not appropriate to resume playback; for example, your app had been interrupted by iPod playback.
	KAudioSessionInterruptionType_ShouldNotResume KaudiosessioninterruptiontypeShould = '!'<<24 | 'r'<<16 | 's'<<8 | 'm' // '!rsm'
	// KAudioSessionInterruptionType_ShouldResume: Indicates that the interruption that has just ended was one for which it is appropriate to immediately resume playback; for example, an incoming phone call was rejected by the user.
	KAudioSessionInterruptionType_ShouldResume KaudiosessioninterruptiontypeShould = 'i'<<24 | 'r'<<16 | 's'<<8 | 'm' // 'irsm'
)

func (e KaudiosessioninterruptiontypeShould) String() string {
	switch e {
	case KAudioSessionInterruptionType_ShouldNotResume:
		return "KAudioSessionInterruptionType_ShouldNotResume"
	case KAudioSessionInterruptionType_ShouldResume:
		return "KAudioSessionInterruptionType_ShouldResume"
	default:
		return fmt.Sprintf("KaudiosessioninterruptiontypeShould(%d)", e)
	}
}

type KaudiosessionpropertyAudioroute uint32

const (
	KAudioSessionProperty_AudioRoute KaudiosessionpropertyAudioroute = 'r'<<24 | 'o'<<16 | 'u'<<8 | 't' // 'rout'
)

func (e KaudiosessionpropertyAudioroute) String() string {
	switch e {
	case KAudioSessionProperty_AudioRoute:
		return "KAudioSessionProperty_AudioRoute"
	default:
		return fmt.Sprintf("KaudiosessionpropertyAudioroute(%d)", e)
	}
}

type KaudiosessionsetactiveflagNotifyothersondeactivation uint32

const (
	// KAudioSessionSetActiveFlag_NotifyOthersOnDeactivation: Indicates that when your audio session deactivates, other audio sessions that had been interrupted by your session can return to their active state.
	KAudioSessionSetActiveFlag_NotifyOthersOnDeactivation KaudiosessionsetactiveflagNotifyothersondeactivation = 1
)

func (e KaudiosessionsetactiveflagNotifyothersondeactivation) String() string {
	switch e {
	case KAudioSessionSetActiveFlag_NotifyOthersOnDeactivation:
		return "KAudioSessionSetActiveFlag_NotifyOthersOnDeactivation"
	default:
		return fmt.Sprintf("KaudiosessionsetactiveflagNotifyothersondeactivation(%d)", e)
	}
}

type KaudiounitcarbonvieweventMouse uint

const (
	KAudioUnitCarbonViewEvent_MouseDownInControl KaudiounitcarbonvieweventMouse = 0
	KAudioUnitCarbonViewEvent_MouseUpInControl   KaudiounitcarbonvieweventMouse = 0
)

func (e KaudiounitcarbonvieweventMouse) String() string {
	switch e {
	case KAudioUnitCarbonViewEvent_MouseDownInControl:
		return "KAudioUnitCarbonViewEvent_MouseDownInControl"
	default:
		return fmt.Sprintf("KaudiounitcarbonvieweventMouse(%d)", e)
	}
}

type KaudiounitclumpidSystem uint32

const (
	// KAudioUnitClumpID_System: Reserved for system use.
	KAudioUnitClumpID_System KaudiounitclumpidSystem = 0
)

func (e KaudiounitclumpidSystem) String() string {
	switch e {
	case KAudioUnitClumpID_System:
		return "KAudioUnitClumpID_System"
	default:
		return fmt.Sprintf("KaudiounitclumpidSystem(%d)", e)
	}
}

type KaudiouniterrInvalidproperty int32

const (
	KAudioComponentErr_InstanceInvalidated     KaudiouniterrInvalidproperty = -66749
	KAudioComponentErr_InstanceTimedOut        KaudiouniterrInvalidproperty = -66754
	KAudioUnitErr_CannotDoInCurrentContext     KaudiouniterrInvalidproperty = -10863
	KAudioUnitErr_ComponentManagerNotSupported KaudiouniterrInvalidproperty = -66740
	KAudioUnitErr_ExtensionNotFound            KaudiouniterrInvalidproperty = -66744
	KAudioUnitErr_FailedInitialization         KaudiouniterrInvalidproperty = -10875
	KAudioUnitErr_FileNotSpecified             KaudiouniterrInvalidproperty = -10869
	KAudioUnitErr_FormatNotSupported           KaudiouniterrInvalidproperty = -10868
	KAudioUnitErr_Initialized                  KaudiouniterrInvalidproperty = -10849
	KAudioUnitErr_InvalidElement               KaudiouniterrInvalidproperty = -10877
	KAudioUnitErr_InvalidFile                  KaudiouniterrInvalidproperty = -10871
	KAudioUnitErr_InvalidFilePath              KaudiouniterrInvalidproperty = -66742
	KAudioUnitErr_InvalidOfflineRender         KaudiouniterrInvalidproperty = -10848
	KAudioUnitErr_InvalidParameter             KaudiouniterrInvalidproperty = -10878
	KAudioUnitErr_InvalidParameterValue        KaudiouniterrInvalidproperty = -66743
	KAudioUnitErr_InvalidProperty              KaudiouniterrInvalidproperty = -10879
	KAudioUnitErr_InvalidPropertyValue         KaudiouniterrInvalidproperty = -10851
	KAudioUnitErr_InvalidScope                 KaudiouniterrInvalidproperty = -10866
	KAudioUnitErr_MIDIOutputBufferFull         KaudiouniterrInvalidproperty = -66753
	KAudioUnitErr_MissingKey                   KaudiouniterrInvalidproperty = -66741
	KAudioUnitErr_MultipleVoiceProcessors      KaudiouniterrInvalidproperty = -66635
	KAudioUnitErr_NoConnection                 KaudiouniterrInvalidproperty = -10876
	KAudioUnitErr_PropertyNotInUse             KaudiouniterrInvalidproperty = -10850
	KAudioUnitErr_PropertyNotWritable          KaudiouniterrInvalidproperty = -10865
	KAudioUnitErr_RenderTimeout                KaudiouniterrInvalidproperty = -66745
	KAudioUnitErr_TooManyFramesToProcess       KaudiouniterrInvalidproperty = -10874
	KAudioUnitErr_Unauthorized                 KaudiouniterrInvalidproperty = -10847
	KAudioUnitErr_Uninitialized                KaudiouniterrInvalidproperty = -10867
	KAudioUnitErr_UnknownFileType              KaudiouniterrInvalidproperty = -10870
)

func (e KaudiouniterrInvalidproperty) String() string {
	switch e {
	case KAudioComponentErr_InstanceInvalidated:
		return "KAudioComponentErr_InstanceInvalidated"
	case KAudioComponentErr_InstanceTimedOut:
		return "KAudioComponentErr_InstanceTimedOut"
	case KAudioUnitErr_CannotDoInCurrentContext:
		return "KAudioUnitErr_CannotDoInCurrentContext"
	case KAudioUnitErr_ComponentManagerNotSupported:
		return "KAudioUnitErr_ComponentManagerNotSupported"
	case KAudioUnitErr_ExtensionNotFound:
		return "KAudioUnitErr_ExtensionNotFound"
	case KAudioUnitErr_FailedInitialization:
		return "KAudioUnitErr_FailedInitialization"
	case KAudioUnitErr_FileNotSpecified:
		return "KAudioUnitErr_FileNotSpecified"
	case KAudioUnitErr_FormatNotSupported:
		return "KAudioUnitErr_FormatNotSupported"
	case KAudioUnitErr_Initialized:
		return "KAudioUnitErr_Initialized"
	case KAudioUnitErr_InvalidElement:
		return "KAudioUnitErr_InvalidElement"
	case KAudioUnitErr_InvalidFile:
		return "KAudioUnitErr_InvalidFile"
	case KAudioUnitErr_InvalidFilePath:
		return "KAudioUnitErr_InvalidFilePath"
	case KAudioUnitErr_InvalidOfflineRender:
		return "KAudioUnitErr_InvalidOfflineRender"
	case KAudioUnitErr_InvalidParameter:
		return "KAudioUnitErr_InvalidParameter"
	case KAudioUnitErr_InvalidParameterValue:
		return "KAudioUnitErr_InvalidParameterValue"
	case KAudioUnitErr_InvalidProperty:
		return "KAudioUnitErr_InvalidProperty"
	case KAudioUnitErr_InvalidPropertyValue:
		return "KAudioUnitErr_InvalidPropertyValue"
	case KAudioUnitErr_InvalidScope:
		return "KAudioUnitErr_InvalidScope"
	case KAudioUnitErr_MIDIOutputBufferFull:
		return "KAudioUnitErr_MIDIOutputBufferFull"
	case KAudioUnitErr_MissingKey:
		return "KAudioUnitErr_MissingKey"
	case KAudioUnitErr_MultipleVoiceProcessors:
		return "KAudioUnitErr_MultipleVoiceProcessors"
	case KAudioUnitErr_NoConnection:
		return "KAudioUnitErr_NoConnection"
	case KAudioUnitErr_PropertyNotInUse:
		return "KAudioUnitErr_PropertyNotInUse"
	case KAudioUnitErr_PropertyNotWritable:
		return "KAudioUnitErr_PropertyNotWritable"
	case KAudioUnitErr_RenderTimeout:
		return "KAudioUnitErr_RenderTimeout"
	case KAudioUnitErr_TooManyFramesToProcess:
		return "KAudioUnitErr_TooManyFramesToProcess"
	case KAudioUnitErr_Unauthorized:
		return "KAudioUnitErr_Unauthorized"
	case KAudioUnitErr_Uninitialized:
		return "KAudioUnitErr_Uninitialized"
	case KAudioUnitErr_UnknownFileType:
		return "KAudioUnitErr_UnknownFileType"
	default:
		return fmt.Sprintf("KaudiouniterrInvalidproperty(%d)", e)
	}
}

type KaudiounitmanufacturerApple uint32

const (
	KAudioUnitManufacturer_Apple KaudiounitmanufacturerApple = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'appl'
)

func (e KaudiounitmanufacturerApple) String() string {
	switch e {
	case KAudioUnitManufacturer_Apple:
		return "KAudioUnitManufacturer_Apple"
	default:
		return fmt.Sprintf("KaudiounitmanufacturerApple(%d)", e)
	}
}

type KaudiounitparameterflagGlobal uint32

const (
	KAudioUnitParameterFlag_Global KaudiounitparameterflagGlobal = 1
	KAudioUnitParameterFlag_Group  KaudiounitparameterflagGlobal = 8
	KAudioUnitParameterFlag_Input  KaudiounitparameterflagGlobal = 2
	KAudioUnitParameterFlag_Output KaudiounitparameterflagGlobal = 4
)

func (e KaudiounitparameterflagGlobal) String() string {
	switch e {
	case KAudioUnitParameterFlag_Global:
		return "KAudioUnitParameterFlag_Global"
	case KAudioUnitParameterFlag_Group:
		return "KAudioUnitParameterFlag_Group"
	case KAudioUnitParameterFlag_Input:
		return "KAudioUnitParameterFlag_Input"
	case KAudioUnitParameterFlag_Output:
		return "KAudioUnitParameterFlag_Output"
	default:
		return fmt.Sprintf("KaudiounitparameterflagGlobal(%d)", e)
	}
}

type KaudiounitparameterflagHasname uint32

const (
	KAudioUnitParameterFlag_HasName KaudiounitparameterflagHasname = 2097152
)

func (e KaudiounitparameterflagHasname) String() string {
	switch e {
	case KAudioUnitParameterFlag_HasName:
		return "KAudioUnitParameterFlag_HasName"
	default:
		return fmt.Sprintf("KaudiounitparameterflagHasname(%d)", e)
	}
}

type KaudiounitparameternameFull int32

const (
	KAudioUnitParameterName_Full KaudiounitparameternameFull = -1
)

func (e KaudiounitparameternameFull) String() string {
	switch e {
	case KAudioUnitParameterName_Full:
		return "KAudioUnitParameterName_Full"
	default:
		return fmt.Sprintf("KaudiounitparameternameFull(%d)", e)
	}
}

type Kaudiounitproperty3dmixerdistanceparams uint32

const (
	KAudioUnitProperty_3DMixerAttenuationCurve Kaudiounitproperty3dmixerdistanceparams = 3013
	KAudioUnitProperty_3DMixerDistanceAtten    Kaudiounitproperty3dmixerdistanceparams = 3004
	KAudioUnitProperty_3DMixerDistanceParams   Kaudiounitproperty3dmixerdistanceparams = 3010
	KAudioUnitProperty_3DMixerRenderingFlags   Kaudiounitproperty3dmixerdistanceparams = 3003
	KAudioUnitProperty_DopplerShift            Kaudiounitproperty3dmixerdistanceparams = 3002
	KAudioUnitProperty_ReverbPreset            Kaudiounitproperty3dmixerdistanceparams = 3012
)

func (e Kaudiounitproperty3dmixerdistanceparams) String() string {
	switch e {
	case KAudioUnitProperty_3DMixerAttenuationCurve:
		return "KAudioUnitProperty_3DMixerAttenuationCurve"
	case KAudioUnitProperty_3DMixerDistanceAtten:
		return "KAudioUnitProperty_3DMixerDistanceAtten"
	case KAudioUnitProperty_3DMixerDistanceParams:
		return "KAudioUnitProperty_3DMixerDistanceParams"
	case KAudioUnitProperty_3DMixerRenderingFlags:
		return "KAudioUnitProperty_3DMixerRenderingFlags"
	case KAudioUnitProperty_DopplerShift:
		return "KAudioUnitProperty_DopplerShift"
	case KAudioUnitProperty_ReverbPreset:
		return "KAudioUnitProperty_ReverbPreset"
	default:
		return fmt.Sprintf("Kaudiounitproperty3dmixerdistanceparams(%d)", e)
	}
}

type KaudiounitpropertyAllparametermidimappings uint32

const (
	KAudioUnitProperty_AddParameterMIDIMapping    KaudiounitpropertyAllparametermidimappings = 42
	KAudioUnitProperty_AllParameterMIDIMappings   KaudiounitpropertyAllparametermidimappings = 41
	KAudioUnitProperty_HotMapParameterMIDIMapping KaudiounitpropertyAllparametermidimappings = 44
	KAudioUnitProperty_RemoveParameterMIDIMapping KaudiounitpropertyAllparametermidimappings = 43
)

func (e KaudiounitpropertyAllparametermidimappings) String() string {
	switch e {
	case KAudioUnitProperty_AddParameterMIDIMapping:
		return "KAudioUnitProperty_AddParameterMIDIMapping"
	case KAudioUnitProperty_AllParameterMIDIMappings:
		return "KAudioUnitProperty_AllParameterMIDIMappings"
	case KAudioUnitProperty_HotMapParameterMIDIMapping:
		return "KAudioUnitProperty_HotMapParameterMIDIMapping"
	case KAudioUnitProperty_RemoveParameterMIDIMapping:
		return "KAudioUnitProperty_RemoveParameterMIDIMapping"
	default:
		return fmt.Sprintf("KaudiounitpropertyAllparametermidimappings(%d)", e)
	}
}

type KaudiounitpropertyClassinfo uint32

const (
	KAudioUnitProperty_AUHostIdentifier KaudiounitpropertyClassinfo = 46
	// KAudioUnitProperty_AudioChannelLayout: A read/write [AudioChannelLayout] data structure valid on the audio unit input and output scopes.
	KAudioUnitProperty_AudioChannelLayout    KaudiounitpropertyClassinfo = 19
	KAudioUnitProperty_AudioUnitMIDIProtocol KaudiounitpropertyClassinfo = 64
	// KAudioUnitProperty_BypassEffect: A read/write [UInt32] value, representing a Boolean value, valid on the audio unit global scope.
	KAudioUnitProperty_BypassEffect KaudiounitpropertyClassinfo = 21
	// KAudioUnitProperty_CPULoad: A read-only [Float64] value valid on the audio unit global scope.
	KAudioUnitProperty_CPULoad KaudiounitpropertyClassinfo = 6
	// KAudioUnitProperty_ClassInfo: Describes the state of an audio unit.
	KAudioUnitProperty_ClassInfo KaudiounitpropertyClassinfo = 0
	// KAudioUnitProperty_ClassInfoFromDocument: A read/write CFDictionary object, valid on the audio unit global scope.
	KAudioUnitProperty_ClassInfoFromDocument KaudiounitpropertyClassinfo = 50
	// KAudioUnitProperty_CocoaUI: A read-only [AudioUnitCocoaViewInfo] data structure valid on the audio unit global scope.
	KAudioUnitProperty_CocoaUI             KaudiounitpropertyClassinfo = 31
	KAudioUnitProperty_ContextName         KaudiounitpropertyClassinfo = 25
	KAudioUnitProperty_DependentParameters KaudiounitpropertyClassinfo = 45
	// KAudioUnitProperty_ElementCount: A read/write [UInt32] value valid on any audio unit scope.
	KAudioUnitProperty_ElementCount KaudiounitpropertyClassinfo = 11
	// KAudioUnitProperty_ElementName: The name of the specified element.
	KAudioUnitProperty_ElementName KaudiounitpropertyClassinfo = 30
	// KAudioUnitProperty_FactoryPresets: So-called factory presets (as opposed to user-configured presets) are ones supplied with an audio unit by the manufacturer.
	KAudioUnitProperty_FactoryPresets KaudiounitpropertyClassinfo = 24
	// KAudioUnitProperty_FastDispatch: A read-only `void *` value valid on the audio unit global scope.
	KAudioUnitProperty_FastDispatch       KaudiounitpropertyClassinfo = 5
	KAudioUnitProperty_FrequencyResponse  KaudiounitpropertyClassinfo = 52
	KAudioUnitProperty_GetUIComponentList KaudiounitpropertyClassinfo = 18
	KAudioUnitProperty_HostCallbacks      KaudiounitpropertyClassinfo = 27
	KAudioUnitProperty_HostMIDIProtocol   KaudiounitpropertyClassinfo = 65
	KAudioUnitProperty_IconLocation       KaudiounitpropertyClassinfo = 39
	// KAudioUnitProperty_InPlaceProcessing: A read/write [UInt32] value, representing a Boolean value, valid on the audio unit global scope.
	KAudioUnitProperty_InPlaceProcessing KaudiounitpropertyClassinfo = 29
	// KAudioUnitProperty_InputSamplesInOutput: A read/write AUInputSamplesInOutputCallbackStruct struct, valid on the audio unit global scope.
	KAudioUnitProperty_InputSamplesInOutput KaudiounitpropertyClassinfo = 49
	// KAudioUnitProperty_LastRenderError: A read-only [OSStatus] value valid on the audio unit global scope.
	KAudioUnitProperty_LastRenderError      KaudiounitpropertyClassinfo = 22
	KAudioUnitProperty_LastRenderSampleTime KaudiounitpropertyClassinfo = 61
	// KAudioUnitProperty_Latency: A read-only [Float64] value valid on the audio unit global scope.
	KAudioUnitProperty_Latency                  KaudiounitpropertyClassinfo = 12
	KAudioUnitProperty_LoadedOutOfProcess       KaudiounitpropertyClassinfo = 62
	KAudioUnitProperty_MIDIOutputBufferSizeHint KaudiounitpropertyClassinfo = 66
	// KAudioUnitProperty_MIDIOutputCallback: A write-only AUMIDIOutputCallbackStruct struct, valid on the audio unit global scope.
	KAudioUnitProperty_MIDIOutputCallback KaudiounitpropertyClassinfo = 48
	// KAudioUnitProperty_MIDIOutputCallbackInfo: A read-only CFArray object valid on the audio unit global scope.
	KAudioUnitProperty_MIDIOutputCallbackInfo      KaudiounitpropertyClassinfo = 47
	KAudioUnitProperty_MIDIOutputEventListCallback KaudiounitpropertyClassinfo = 63
	// KAudioUnitProperty_MakeConnection: A write-only AudioUnitConnection data structure valid on the audio unit input scope.
	KAudioUnitProperty_MakeConnection KaudiounitpropertyClassinfo = 1
	// KAudioUnitProperty_MaximumFramesPerSlice: Specifies the maximum number of sample frames an audio unit is prepared to supply on one invocation of its AudioUnitRender(_:_:_:_:_:_:) function.
	KAudioUnitProperty_MaximumFramesPerSlice KaudiounitpropertyClassinfo = 14
	KAudioUnitProperty_NickName              KaudiounitpropertyClassinfo = 54
	KAudioUnitProperty_OfflineRender         KaudiounitpropertyClassinfo = 37
	// KAudioUnitProperty_ParameterClumpName: A read-only [AudioUnitParameterNameInfo] struct, valid on any audio unit scope.
	KAudioUnitProperty_ParameterClumpName KaudiounitpropertyClassinfo = 35
	// KAudioUnitProperty_ParameterHistoryInfo: For parameters that have the flag_PlotHistory flag set, getting this property fills out the AudioUnitParameterHistoryInfo struct containing the recommended update rate and history duration.
	KAudioUnitProperty_ParameterHistoryInfo KaudiounitpropertyClassinfo = 53
	// KAudioUnitProperty_ParameterIDName: A shortened version of an audio unit parameter name, suitable for compact display situations.
	KAudioUnitProperty_ParameterIDName KaudiounitpropertyClassinfo = 34
	KAudioUnitProperty_ParameterInfo   KaudiounitpropertyClassinfo = 4
	// KAudioUnitProperty_ParameterList: A list of read-only parameter ID values valid on any audio unit scope.
	KAudioUnitProperty_ParameterList KaudiounitpropertyClassinfo = 3
	// KAudioUnitProperty_ParameterStringFromValue: A read-only AudioUnitParameterStringFromValue struct, valid on any audio unit scope.
	KAudioUnitProperty_ParameterStringFromValue KaudiounitpropertyClassinfo = 33
	KAudioUnitProperty_ParameterValueFromString KaudiounitpropertyClassinfo = 38
	// KAudioUnitProperty_ParameterValueStrings: An array of names for a named, indexed audio unit parameter.
	KAudioUnitProperty_ParameterValueStrings KaudiounitpropertyClassinfo = 16
	KAudioUnitProperty_ParametersForOverview KaudiounitpropertyClassinfo = 57
	// KAudioUnitProperty_PresentPreset: The active factory preset for an audio unit.
	KAudioUnitProperty_PresentPreset       KaudiounitpropertyClassinfo = 36
	KAudioUnitProperty_PresentationLatency KaudiounitpropertyClassinfo = 40
	// KAudioUnitProperty_RenderContextObserver: The block that the system calls when the rendering context changes.
	KAudioUnitProperty_RenderContextObserver KaudiounitpropertyClassinfo = 60
	// KAudioUnitProperty_RenderQuality: A read/write [UInt32] value valid on the audio unit global scope.
	KAudioUnitProperty_RenderQuality         KaudiounitpropertyClassinfo = 26
	KAudioUnitProperty_RequestViewController KaudiounitpropertyClassinfo = 56
	KAudioUnitProperty_SampleRate            KaudiounitpropertyClassinfo = 2
	KAudioUnitProperty_SetExternalBuffer     KaudiounitpropertyClassinfo = 15
	KAudioUnitProperty_SetRenderCallback     KaudiounitpropertyClassinfo = 23
	// KAudioUnitProperty_ShouldAllocateBuffer: A read/write [UInt32] value valid on the audio unit input and output scopes, settable individually on each element.
	KAudioUnitProperty_ShouldAllocateBuffer KaudiounitpropertyClassinfo = 51
	KAudioUnitProperty_StreamFormat         KaudiounitpropertyClassinfo = 8
	// KAudioUnitProperty_SupportedChannelLayoutTags: A read-only array on [AudioChannelLayoutTag] structures, valid on the audio unit input and output scopes.
	KAudioUnitProperty_SupportedChannelLayoutTags KaudiounitpropertyClassinfo = 32
	// KAudioUnitProperty_SupportedNumChannels: A read-only array of channel information structures valid on the audio unit global scope.
	KAudioUnitProperty_SupportedNumChannels KaudiounitpropertyClassinfo = 13
	KAudioUnitProperty_SupportsMPE          KaudiounitpropertyClassinfo = 58
	// KAudioUnitProperty_TailTime: A read-only [Float64] value valid on the audio unit global scope.
	KAudioUnitProperty_TailTime KaudiounitpropertyClassinfo = 20
)

func (e KaudiounitpropertyClassinfo) String() string {
	switch e {
	case KAudioUnitProperty_AUHostIdentifier:
		return "KAudioUnitProperty_AUHostIdentifier"
	case KAudioUnitProperty_AudioChannelLayout:
		return "KAudioUnitProperty_AudioChannelLayout"
	case KAudioUnitProperty_AudioUnitMIDIProtocol:
		return "KAudioUnitProperty_AudioUnitMIDIProtocol"
	case KAudioUnitProperty_BypassEffect:
		return "KAudioUnitProperty_BypassEffect"
	case KAudioUnitProperty_CPULoad:
		return "KAudioUnitProperty_CPULoad"
	case KAudioUnitProperty_ClassInfo:
		return "KAudioUnitProperty_ClassInfo"
	case KAudioUnitProperty_ClassInfoFromDocument:
		return "KAudioUnitProperty_ClassInfoFromDocument"
	case KAudioUnitProperty_CocoaUI:
		return "KAudioUnitProperty_CocoaUI"
	case KAudioUnitProperty_ContextName:
		return "KAudioUnitProperty_ContextName"
	case KAudioUnitProperty_DependentParameters:
		return "KAudioUnitProperty_DependentParameters"
	case KAudioUnitProperty_ElementCount:
		return "KAudioUnitProperty_ElementCount"
	case KAudioUnitProperty_ElementName:
		return "KAudioUnitProperty_ElementName"
	case KAudioUnitProperty_FactoryPresets:
		return "KAudioUnitProperty_FactoryPresets"
	case KAudioUnitProperty_FastDispatch:
		return "KAudioUnitProperty_FastDispatch"
	case KAudioUnitProperty_FrequencyResponse:
		return "KAudioUnitProperty_FrequencyResponse"
	case KAudioUnitProperty_GetUIComponentList:
		return "KAudioUnitProperty_GetUIComponentList"
	case KAudioUnitProperty_HostCallbacks:
		return "KAudioUnitProperty_HostCallbacks"
	case KAudioUnitProperty_HostMIDIProtocol:
		return "KAudioUnitProperty_HostMIDIProtocol"
	case KAudioUnitProperty_IconLocation:
		return "KAudioUnitProperty_IconLocation"
	case KAudioUnitProperty_InPlaceProcessing:
		return "KAudioUnitProperty_InPlaceProcessing"
	case KAudioUnitProperty_InputSamplesInOutput:
		return "KAudioUnitProperty_InputSamplesInOutput"
	case KAudioUnitProperty_LastRenderError:
		return "KAudioUnitProperty_LastRenderError"
	case KAudioUnitProperty_LastRenderSampleTime:
		return "KAudioUnitProperty_LastRenderSampleTime"
	case KAudioUnitProperty_Latency:
		return "KAudioUnitProperty_Latency"
	case KAudioUnitProperty_LoadedOutOfProcess:
		return "KAudioUnitProperty_LoadedOutOfProcess"
	case KAudioUnitProperty_MIDIOutputBufferSizeHint:
		return "KAudioUnitProperty_MIDIOutputBufferSizeHint"
	case KAudioUnitProperty_MIDIOutputCallback:
		return "KAudioUnitProperty_MIDIOutputCallback"
	case KAudioUnitProperty_MIDIOutputCallbackInfo:
		return "KAudioUnitProperty_MIDIOutputCallbackInfo"
	case KAudioUnitProperty_MIDIOutputEventListCallback:
		return "KAudioUnitProperty_MIDIOutputEventListCallback"
	case KAudioUnitProperty_MakeConnection:
		return "KAudioUnitProperty_MakeConnection"
	case KAudioUnitProperty_MaximumFramesPerSlice:
		return "KAudioUnitProperty_MaximumFramesPerSlice"
	case KAudioUnitProperty_NickName:
		return "KAudioUnitProperty_NickName"
	case KAudioUnitProperty_OfflineRender:
		return "KAudioUnitProperty_OfflineRender"
	case KAudioUnitProperty_ParameterClumpName:
		return "KAudioUnitProperty_ParameterClumpName"
	case KAudioUnitProperty_ParameterHistoryInfo:
		return "KAudioUnitProperty_ParameterHistoryInfo"
	case KAudioUnitProperty_ParameterIDName:
		return "KAudioUnitProperty_ParameterIDName"
	case KAudioUnitProperty_ParameterInfo:
		return "KAudioUnitProperty_ParameterInfo"
	case KAudioUnitProperty_ParameterList:
		return "KAudioUnitProperty_ParameterList"
	case KAudioUnitProperty_ParameterStringFromValue:
		return "KAudioUnitProperty_ParameterStringFromValue"
	case KAudioUnitProperty_ParameterValueFromString:
		return "KAudioUnitProperty_ParameterValueFromString"
	case KAudioUnitProperty_ParameterValueStrings:
		return "KAudioUnitProperty_ParameterValueStrings"
	case KAudioUnitProperty_ParametersForOverview:
		return "KAudioUnitProperty_ParametersForOverview"
	case KAudioUnitProperty_PresentPreset:
		return "KAudioUnitProperty_PresentPreset"
	case KAudioUnitProperty_PresentationLatency:
		return "KAudioUnitProperty_PresentationLatency"
	case KAudioUnitProperty_RenderContextObserver:
		return "KAudioUnitProperty_RenderContextObserver"
	case KAudioUnitProperty_RenderQuality:
		return "KAudioUnitProperty_RenderQuality"
	case KAudioUnitProperty_RequestViewController:
		return "KAudioUnitProperty_RequestViewController"
	case KAudioUnitProperty_SampleRate:
		return "KAudioUnitProperty_SampleRate"
	case KAudioUnitProperty_SetExternalBuffer:
		return "KAudioUnitProperty_SetExternalBuffer"
	case KAudioUnitProperty_SetRenderCallback:
		return "KAudioUnitProperty_SetRenderCallback"
	case KAudioUnitProperty_ShouldAllocateBuffer:
		return "KAudioUnitProperty_ShouldAllocateBuffer"
	case KAudioUnitProperty_StreamFormat:
		return "KAudioUnitProperty_StreamFormat"
	case KAudioUnitProperty_SupportedChannelLayoutTags:
		return "KAudioUnitProperty_SupportedChannelLayoutTags"
	case KAudioUnitProperty_SupportedNumChannels:
		return "KAudioUnitProperty_SupportedNumChannels"
	case KAudioUnitProperty_SupportsMPE:
		return "KAudioUnitProperty_SupportsMPE"
	case KAudioUnitProperty_TailTime:
		return "KAudioUnitProperty_TailTime"
	default:
		return fmt.Sprintf("KaudiounitpropertyClassinfo(%d)", e)
	}
}

type KaudiounitpropertyDeferredrenderer uint32

const (
	KAudioUnitProperty_DeferredRendererExtraLatency KaudiounitpropertyDeferredrenderer = 3321
	KAudioUnitProperty_DeferredRendererPullSize     KaudiounitpropertyDeferredrenderer = 3320
	KAudioUnitProperty_DeferredRendererWaitFrames   KaudiounitpropertyDeferredrenderer = 3322
)

func (e KaudiounitpropertyDeferredrenderer) String() string {
	switch e {
	case KAudioUnitProperty_DeferredRendererExtraLatency:
		return "KAudioUnitProperty_DeferredRendererExtraLatency"
	case KAudioUnitProperty_DeferredRendererPullSize:
		return "KAudioUnitProperty_DeferredRendererPullSize"
	case KAudioUnitProperty_DeferredRendererWaitFrames:
		return "KAudioUnitProperty_DeferredRendererWaitFrames"
	default:
		return fmt.Sprintf("KaudiounitpropertyDeferredrenderer(%d)", e)
	}
}

type KaudiounitpropertyDistanceattenuationdata uint32

const (
	KAudioUnitProperty_DistanceAttenuationData KaudiounitpropertyDistanceattenuationdata = 3600
)

func (e KaudiounitpropertyDistanceattenuationdata) String() string {
	switch e {
	case KAudioUnitProperty_DistanceAttenuationData:
		return "KAudioUnitProperty_DistanceAttenuationData"
	default:
		return fmt.Sprintf("KaudiounitpropertyDistanceattenuationdata(%d)", e)
	}
}

type KaudiounitpropertyMeteringmode uint32

const (
	KAudioUnitProperty_InputAnchorTimeStamp KaudiounitpropertyMeteringmode = 3016
	// KAudioUnitProperty_MatrixDimensions: Indicates the total number of channels for input and output of a given matrix mixer.
	KAudioUnitProperty_MatrixDimensions KaudiounitpropertyMeteringmode = 3009
	// KAudioUnitProperty_MatrixLevels: Describes the internal state of a matrix mixer.
	KAudioUnitProperty_MatrixLevels KaudiounitpropertyMeteringmode = 3006
	// KAudioUnitProperty_MeterClipping: Indicates audio clipping that has occurred since this property was last accessed.
	KAudioUnitProperty_MeterClipping KaudiounitpropertyMeteringmode = 3011
	// KAudioUnitProperty_MeteringMode: Specifies whether metering is enabled or disabled for a particular scope-element combination.
	KAudioUnitProperty_MeteringMode KaudiounitpropertyMeteringmode = 3007
)

func (e KaudiounitpropertyMeteringmode) String() string {
	switch e {
	case KAudioUnitProperty_InputAnchorTimeStamp:
		return "KAudioUnitProperty_InputAnchorTimeStamp"
	case KAudioUnitProperty_MatrixDimensions:
		return "KAudioUnitProperty_MatrixDimensions"
	case KAudioUnitProperty_MatrixLevels:
		return "KAudioUnitProperty_MatrixLevels"
	case KAudioUnitProperty_MeterClipping:
		return "KAudioUnitProperty_MeterClipping"
	case KAudioUnitProperty_MeteringMode:
		return "KAudioUnitProperty_MeteringMode"
	default:
		return fmt.Sprintf("KaudiounitpropertyMeteringmode(%d)", e)
	}
}

type KaudiounitpropertyRemotecontroleventlistener uint32

const (
	KAudioUnitProperty_IsInterAppConnected        KaudiounitpropertyRemotecontroleventlistener = 101
	KAudioUnitProperty_PeerURL                    KaudiounitpropertyRemotecontroleventlistener = 102
	KAudioUnitProperty_RemoteControlEventListener KaudiounitpropertyRemotecontroleventlistener = 100
)

func (e KaudiounitpropertyRemotecontroleventlistener) String() string {
	switch e {
	case KAudioUnitProperty_IsInterAppConnected:
		return "KAudioUnitProperty_IsInterAppConnected"
	case KAudioUnitProperty_PeerURL:
		return "KAudioUnitProperty_PeerURL"
	case KAudioUnitProperty_RemoteControlEventListener:
		return "KAudioUnitProperty_RemoteControlEventListener"
	default:
		return fmt.Sprintf("KaudiounitpropertyRemotecontroleventlistener(%d)", e)
	}
}

type KaudiounitpropertyReverbroomtype uint32

const (
	KAudioUnitProperty_ReverbRoomType                              KaudiounitpropertyReverbroomtype = 10
	KAudioUnitProperty_SpatialMixerAnyInputIsUsingPersonalizedHRTF KaudiounitpropertyReverbroomtype = 3116
	KAudioUnitProperty_SpatialMixerAttenuationCurve                KaudiounitpropertyReverbroomtype = 3013
	KAudioUnitProperty_SpatialMixerDistanceParams                  KaudiounitpropertyReverbroomtype = 3010
	KAudioUnitProperty_SpatialMixerEnableHeadTracking              KaudiounitpropertyReverbroomtype = 3111
	KAudioUnitProperty_SpatialMixerOutputType                      KaudiounitpropertyReverbroomtype = 3100
	KAudioUnitProperty_SpatialMixerPersonalizedHRTFMode            KaudiounitpropertyReverbroomtype = 3113
	KAudioUnitProperty_SpatialMixerPointSourceInHeadMode           KaudiounitpropertyReverbroomtype = 3103
	KAudioUnitProperty_SpatialMixerRenderingFlags                  KaudiounitpropertyReverbroomtype = 3003
	KAudioUnitProperty_SpatialMixerSourceMode                      KaudiounitpropertyReverbroomtype = 3005
	KAudioUnitProperty_SpatializationAlgorithm                     KaudiounitpropertyReverbroomtype = 3000
	KAudioUnitProperty_UsesInternalReverb                          KaudiounitpropertyReverbroomtype = 1005
)

func (e KaudiounitpropertyReverbroomtype) String() string {
	switch e {
	case KAudioUnitProperty_ReverbRoomType:
		return "KAudioUnitProperty_ReverbRoomType"
	case KAudioUnitProperty_SpatialMixerAnyInputIsUsingPersonalizedHRTF:
		return "KAudioUnitProperty_SpatialMixerAnyInputIsUsingPersonalizedHRTF"
	case KAudioUnitProperty_SpatialMixerAttenuationCurve:
		return "KAudioUnitProperty_SpatialMixerAttenuationCurve"
	case KAudioUnitProperty_SpatialMixerDistanceParams:
		return "KAudioUnitProperty_SpatialMixerDistanceParams"
	case KAudioUnitProperty_SpatialMixerEnableHeadTracking:
		return "KAudioUnitProperty_SpatialMixerEnableHeadTracking"
	case KAudioUnitProperty_SpatialMixerOutputType:
		return "KAudioUnitProperty_SpatialMixerOutputType"
	case KAudioUnitProperty_SpatialMixerPersonalizedHRTFMode:
		return "KAudioUnitProperty_SpatialMixerPersonalizedHRTFMode"
	case KAudioUnitProperty_SpatialMixerPointSourceInHeadMode:
		return "KAudioUnitProperty_SpatialMixerPointSourceInHeadMode"
	case KAudioUnitProperty_SpatialMixerRenderingFlags:
		return "KAudioUnitProperty_SpatialMixerRenderingFlags"
	case KAudioUnitProperty_SpatialMixerSourceMode:
		return "KAudioUnitProperty_SpatialMixerSourceMode"
	case KAudioUnitProperty_SpatializationAlgorithm:
		return "KAudioUnitProperty_SpatializationAlgorithm"
	case KAudioUnitProperty_UsesInternalReverb:
		return "KAudioUnitProperty_UsesInternalReverb"
	default:
		return fmt.Sprintf("KaudiounitpropertyReverbroomtype(%d)", e)
	}
}

type KaudiounitpropertySamplerateconvertercomplexity uint32

const (
	KAudioUnitProperty_SampleRateConverterComplexity KaudiounitpropertySamplerateconvertercomplexity = 3014
)

func (e KaudiounitpropertySamplerateconvertercomplexity) String() string {
	switch e {
	case KAudioUnitProperty_SampleRateConverterComplexity:
		return "KAudioUnitProperty_SampleRateConverterComplexity"
	default:
		return fmt.Sprintf("KaudiounitpropertySamplerateconvertercomplexity(%d)", e)
	}
}

type KaudiounitpropertyScheduleaudioslice uint32

const (
	KAudioUnitProperty_CurrentPlayTime        KaudiounitpropertyScheduleaudioslice = 3302
	KAudioUnitProperty_ScheduleAudioSlice     KaudiounitpropertyScheduleaudioslice = 3300
	KAudioUnitProperty_ScheduleStartTimeStamp KaudiounitpropertyScheduleaudioslice = 3301
)

func (e KaudiounitpropertyScheduleaudioslice) String() string {
	switch e {
	case KAudioUnitProperty_CurrentPlayTime:
		return "KAudioUnitProperty_CurrentPlayTime"
	case KAudioUnitProperty_ScheduleAudioSlice:
		return "KAudioUnitProperty_ScheduleAudioSlice"
	case KAudioUnitProperty_ScheduleStartTimeStamp:
		return "KAudioUnitProperty_ScheduleStartTimeStamp"
	default:
		return fmt.Sprintf("KaudiounitpropertyScheduleaudioslice(%d)", e)
	}
}

type KaudiounitpropertyScheduledfile uint32

const (
	KAudioUnitProperty_ScheduledFileBufferSizeFrames KaudiounitpropertyScheduledfile = 3313
	KAudioUnitProperty_ScheduledFileIDs              KaudiounitpropertyScheduledfile = 3310
	KAudioUnitProperty_ScheduledFileNumberBuffers    KaudiounitpropertyScheduledfile = 3314
	KAudioUnitProperty_ScheduledFilePrime            KaudiounitpropertyScheduledfile = 3312
	KAudioUnitProperty_ScheduledFileRegion           KaudiounitpropertyScheduledfile = 3311
)

func (e KaudiounitpropertyScheduledfile) String() string {
	switch e {
	case KAudioUnitProperty_ScheduledFileBufferSizeFrames:
		return "KAudioUnitProperty_ScheduledFileBufferSizeFrames"
	case KAudioUnitProperty_ScheduledFileIDs:
		return "KAudioUnitProperty_ScheduledFileIDs"
	case KAudioUnitProperty_ScheduledFileNumberBuffers:
		return "KAudioUnitProperty_ScheduledFileNumberBuffers"
	case KAudioUnitProperty_ScheduledFilePrime:
		return "KAudioUnitProperty_ScheduledFilePrime"
	case KAudioUnitProperty_ScheduledFileRegion:
		return "KAudioUnitProperty_ScheduledFileRegion"
	default:
		return fmt.Sprintf("KaudiounitpropertyScheduledfile(%d)", e)
	}
}

type KaudiounitpropertySpeakerconfiguration uint32

const (
	KAudioUnitProperty_SpeakerConfiguration KaudiounitpropertySpeakerconfiguration = 3001
)

func (e KaudiounitpropertySpeakerconfiguration) String() string {
	switch e {
	case KAudioUnitProperty_SpeakerConfiguration:
		return "KAudioUnitProperty_SpeakerConfiguration"
	default:
		return fmt.Sprintf("KaudiounitpropertySpeakerconfiguration(%d)", e)
	}
}

type KaudiounitpropertySrcalgorithm uint32

const (
	KAudioOfflineUnitProperty_InputSize   KaudiounitpropertySrcalgorithm = 3020
	KAudioOfflineUnitProperty_OutputSize  KaudiounitpropertySrcalgorithm = 3021
	KAudioUnitProperty_BusCount           KaudiounitpropertySrcalgorithm = 11
	KAudioUnitProperty_CurrentPreset      KaudiounitpropertySrcalgorithm = 28
	KAudioUnitProperty_MIDIControlMapping KaudiounitpropertySrcalgorithm = 17
	KAudioUnitProperty_ParameterValueName KaudiounitpropertySrcalgorithm = 33
	KAudioUnitProperty_SRCAlgorithm       KaudiounitpropertySrcalgorithm = 9
)

func (e KaudiounitpropertySrcalgorithm) String() string {
	switch e {
	case KAudioOfflineUnitProperty_InputSize:
		return "KAudioOfflineUnitProperty_InputSize"
	case KAudioOfflineUnitProperty_OutputSize:
		return "KAudioOfflineUnitProperty_OutputSize"
	case KAudioUnitProperty_BusCount:
		return "KAudioUnitProperty_BusCount"
	case KAudioUnitProperty_CurrentPreset:
		return "KAudioUnitProperty_CurrentPreset"
	case KAudioUnitProperty_MIDIControlMapping:
		return "KAudioUnitProperty_MIDIControlMapping"
	case KAudioUnitProperty_ParameterValueName:
		return "KAudioUnitProperty_ParameterValueName"
	case KAudioUnitProperty_SRCAlgorithm:
		return "KAudioUnitProperty_SRCAlgorithm"
	default:
		return fmt.Sprintf("KaudiounitpropertySrcalgorithm(%d)", e)
	}
}

type KaudiounitsubtypeAuconverter uint32

const (
	KAudioUnitSubType_AUAudioMix       KaudiounitsubtypeAuconverter = 'a'<<24 | 'm'<<16 | 'i'<<8 | 'x' // 'amix'
	KAudioUnitSubType_AUConverter      KaudiounitsubtypeAuconverter = 'c'<<24 | 'o'<<16 | 'n'<<8 | 'v' // 'conv'
	KAudioUnitSubType_AUiPodTimeOther  KaudiounitsubtypeAuconverter = 'i'<<24 | 'p'<<16 | 't'<<8 | 'o' // 'ipto'
	KAudioUnitSubType_DeferredRenderer KaudiounitsubtypeAuconverter = 'd'<<24 | 'e'<<16 | 'f'<<8 | 'r' // 'defr'
	KAudioUnitSubType_Merger           KaudiounitsubtypeAuconverter = 'm'<<24 | 'e'<<16 | 'r'<<8 | 'g' // 'merg'
	KAudioUnitSubType_MultiSplitter    KaudiounitsubtypeAuconverter = 'm'<<24 | 's'<<16 | 'p'<<8 | 'l' // 'mspl'
	KAudioUnitSubType_NewTimePitch     KaudiounitsubtypeAuconverter = 'n'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'nutp'
	KAudioUnitSubType_RoundTripAAC     KaudiounitsubtypeAuconverter = 'r'<<24 | 'a'<<16 | 'a'<<8 | 'c' // 'raac'
	KAudioUnitSubType_Splitter         KaudiounitsubtypeAuconverter = 's'<<24 | 'p'<<16 | 'l'<<8 | 't' // 'splt'
	// KAudioUnitSubType_Varispeed: An audio unit that can control playback rate.
	KAudioUnitSubType_Varispeed KaudiounitsubtypeAuconverter = 'v'<<24 | 'a'<<16 | 'r'<<8 | 'i' // 'vari'
)

func (e KaudiounitsubtypeAuconverter) String() string {
	switch e {
	case KAudioUnitSubType_AUAudioMix:
		return "KAudioUnitSubType_AUAudioMix"
	case KAudioUnitSubType_AUConverter:
		return "KAudioUnitSubType_AUConverter"
	case KAudioUnitSubType_AUiPodTimeOther:
		return "KAudioUnitSubType_AUiPodTimeOther"
	case KAudioUnitSubType_DeferredRenderer:
		return "KAudioUnitSubType_DeferredRenderer"
	case KAudioUnitSubType_Merger:
		return "KAudioUnitSubType_Merger"
	case KAudioUnitSubType_MultiSplitter:
		return "KAudioUnitSubType_MultiSplitter"
	case KAudioUnitSubType_NewTimePitch:
		return "KAudioUnitSubType_NewTimePitch"
	case KAudioUnitSubType_RoundTripAAC:
		return "KAudioUnitSubType_RoundTripAAC"
	case KAudioUnitSubType_Splitter:
		return "KAudioUnitSubType_Splitter"
	case KAudioUnitSubType_Varispeed:
		return "KAudioUnitSubType_Varispeed"
	default:
		return fmt.Sprintf("KaudiounitsubtypeAuconverter(%d)", e)
	}
}

type KaudiounitsubtypeDlssynth uint32

const (
	// KAudioUnitSubType_DLSSynth: A multitimbral instrument unit that can use sample banks in either DLS or SoundFont formats.
	KAudioUnitSubType_DLSSynth  KaudiounitsubtypeDlssynth = 'd'<<24 | 'l'<<16 | 's'<<8 | ' ' // 'dls '
	KAudioUnitSubType_MIDISynth KaudiounitsubtypeDlssynth = 'm'<<24 | 's'<<16 | 'y'<<8 | 'n' // 'msyn'
	KAudioUnitSubType_Sampler   KaudiounitsubtypeDlssynth = 's'<<24 | 'a'<<16 | 'm'<<8 | 'p' // 'samp'
)

func (e KaudiounitsubtypeDlssynth) String() string {
	switch e {
	case KAudioUnitSubType_DLSSynth:
		return "KAudioUnitSubType_DLSSynth"
	case KAudioUnitSubType_MIDISynth:
		return "KAudioUnitSubType_MIDISynth"
	case KAudioUnitSubType_Sampler:
		return "KAudioUnitSubType_Sampler"
	default:
		return fmt.Sprintf("KaudiounitsubtypeDlssynth(%d)", e)
	}
}

type KaudiounitsubtypeGenericoutput uint32

const (
	KAudioUnitSubType_GenericOutput KaudiounitsubtypeGenericoutput = 'g'<<24 | 'e'<<16 | 'n'<<8 | 'r' // 'genr'
	// KAudioUnitSubType_VoiceProcessingIO: An audio unit that interfaces to the audio inputs and outputs of iOS devices and provides voice processing features.
	KAudioUnitSubType_VoiceProcessingIO KaudiounitsubtypeGenericoutput = 'v'<<24 | 'p'<<16 | 'i'<<8 | 'o' // 'vpio'
)

func (e KaudiounitsubtypeGenericoutput) String() string {
	switch e {
	case KAudioUnitSubType_GenericOutput:
		return "KAudioUnitSubType_GenericOutput"
	case KAudioUnitSubType_VoiceProcessingIO:
		return "KAudioUnitSubType_VoiceProcessingIO"
	default:
		return fmt.Sprintf("KaudiounitsubtypeGenericoutput(%d)", e)
	}
}

type KaudiounitsubtypeGraphiceq uint32

const (
	KAudioUnitSubType_AUFilter KaudiounitsubtypeGraphiceq = 'f'<<24 | 'i'<<16 | 'l'<<8 | 't' // 'filt'
	// KAudioUnitSubType_GraphicEQ: An audio unit that provides a 10- or 31-band graphic equalizer.
	KAudioUnitSubType_GraphicEQ    KaudiounitsubtypeGraphiceq = 'g'<<24 | 'r'<<16 | 'e'<<8 | 'q' // 'greq'
	KAudioUnitSubType_MatrixReverb KaudiounitsubtypeGraphiceq = 'm'<<24 | 'r'<<16 | 'e'<<8 | 'v' // 'mrev'
	// KAudioUnitSubType_MultiBandCompressor: An audio unit that provides four-bands of dynamic compression or expansion.
	KAudioUnitSubType_MultiBandCompressor KaudiounitsubtypeGraphiceq = 'm'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'mcmp'
	KAudioUnitSubType_NetSend             KaudiounitsubtypeGraphiceq = 'n'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'nsnd'
	// KAudioUnitSubType_Pitch: An audio unit for modifying the pitch of a signal.
	KAudioUnitSubType_Pitch     KaudiounitsubtypeGraphiceq = 't'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'tmpt'
	KAudioUnitSubType_RogerBeep KaudiounitsubtypeGraphiceq = 'r'<<24 | 'o'<<16 | 'g'<<8 | 'r' // 'rogr'
)

func (e KaudiounitsubtypeGraphiceq) String() string {
	switch e {
	case KAudioUnitSubType_AUFilter:
		return "KAudioUnitSubType_AUFilter"
	case KAudioUnitSubType_GraphicEQ:
		return "KAudioUnitSubType_GraphicEQ"
	case KAudioUnitSubType_MatrixReverb:
		return "KAudioUnitSubType_MatrixReverb"
	case KAudioUnitSubType_MultiBandCompressor:
		return "KAudioUnitSubType_MultiBandCompressor"
	case KAudioUnitSubType_NetSend:
		return "KAudioUnitSubType_NetSend"
	case KAudioUnitSubType_Pitch:
		return "KAudioUnitSubType_Pitch"
	case KAudioUnitSubType_RogerBeep:
		return "KAudioUnitSubType_RogerBeep"
	default:
		return fmt.Sprintf("KaudiounitsubtypeGraphiceq(%d)", e)
	}
}

type KaudiounitsubtypeHaloutput uint32

const (
	KAudioUnitSubType_DefaultOutput KaudiounitsubtypeHaloutput = 'd'<<24 | 'e'<<16 | 'f'<<8 | ' ' // 'def '
	KAudioUnitSubType_HALOutput     KaudiounitsubtypeHaloutput = 'a'<<24 | 'h'<<16 | 'a'<<8 | 'l' // 'ahal'
	KAudioUnitSubType_SystemOutput  KaudiounitsubtypeHaloutput = 's'<<24 | 'y'<<16 | 's'<<8 | ' ' // 'sys '
)

func (e KaudiounitsubtypeHaloutput) String() string {
	switch e {
	case KAudioUnitSubType_DefaultOutput:
		return "KAudioUnitSubType_DefaultOutput"
	case KAudioUnitSubType_HALOutput:
		return "KAudioUnitSubType_HALOutput"
	case KAudioUnitSubType_SystemOutput:
		return "KAudioUnitSubType_SystemOutput"
	default:
		return fmt.Sprintf("KaudiounitsubtypeHaloutput(%d)", e)
	}
}

type KaudiounitsubtypeMultichannelmixer uint32

const (
	KAudioUnitSubType_MatrixMixer KaudiounitsubtypeMultichannelmixer = 'm'<<24 | 'x'<<16 | 'm'<<8 | 'x' // 'mxmx'
	// KAudioUnitSubType_MultiChannelMixer: An audio unit that can have any number of input buses, with any number of channels on each input bus, and one output bus.
	KAudioUnitSubType_MultiChannelMixer KaudiounitsubtypeMultichannelmixer = 'm'<<24 | 'c'<<16 | 'm'<<8 | 'x' // 'mcmx'
	KAudioUnitSubType_SpatialMixer      KaudiounitsubtypeMultichannelmixer = '3'<<24 | 'd'<<16 | 'e'<<8 | 'm' // '3dem'
)

func (e KaudiounitsubtypeMultichannelmixer) String() string {
	switch e {
	case KAudioUnitSubType_MatrixMixer:
		return "KAudioUnitSubType_MatrixMixer"
	case KAudioUnitSubType_MultiChannelMixer:
		return "KAudioUnitSubType_MultiChannelMixer"
	case KAudioUnitSubType_SpatialMixer:
		return "KAudioUnitSubType_SpatialMixer"
	default:
		return fmt.Sprintf("KaudiounitsubtypeMultichannelmixer(%d)", e)
	}
}

type KaudiounitsubtypeNetreceive uint32

const (
	KAudioUnitSubType_AudioFilePlayer      KaudiounitsubtypeNetreceive = 'a'<<24 | 'f'<<16 | 'p'<<8 | 'l' // 'afpl'
	KAudioUnitSubType_NetReceive           KaudiounitsubtypeNetreceive = 'n'<<24 | 'r'<<16 | 'c'<<8 | 'v' // 'nrcv'
	KAudioUnitSubType_ScheduledSoundPlayer KaudiounitsubtypeNetreceive = 's'<<24 | 's'<<16 | 'p'<<8 | 'l' // 'sspl'
)

func (e KaudiounitsubtypeNetreceive) String() string {
	switch e {
	case KAudioUnitSubType_AudioFilePlayer:
		return "KAudioUnitSubType_AudioFilePlayer"
	case KAudioUnitSubType_NetReceive:
		return "KAudioUnitSubType_NetReceive"
	case KAudioUnitSubType_ScheduledSoundPlayer:
		return "KAudioUnitSubType_ScheduledSoundPlayer"
	default:
		return fmt.Sprintf("KaudiounitsubtypeNetreceive(%d)", e)
	}
}

type KaudiounitsubtypePeaklimiter uint32

const (
	KAudioUnitSubType_AUSoundIsolation KaudiounitsubtypePeaklimiter = 'v'<<24 | 'o'<<16 | 'i'<<8 | 's' // 'vois'
	KAudioUnitSubType_BandPassFilter   KaudiounitsubtypePeaklimiter = 'b'<<24 | 'p'<<16 | 'a'<<8 | 's' // 'bpas'
	// KAudioUnitSubType_Delay: An audio unit that introduces a time delay to a signal.
	KAudioUnitSubType_Delay             KaudiounitsubtypePeaklimiter = 'd'<<24 | 'e'<<16 | 'l'<<8 | 'y' // 'dely'
	KAudioUnitSubType_Distortion        KaudiounitsubtypePeaklimiter = 'd'<<24 | 'i'<<16 | 's'<<8 | 't' // 'dist'
	KAudioUnitSubType_DynamicsProcessor KaudiounitsubtypePeaklimiter = 'd'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'dcmp'
	KAudioUnitSubType_HighPassFilter    KaudiounitsubtypePeaklimiter = 'h'<<24 | 'p'<<16 | 'a'<<8 | 's' // 'hpas'
	// KAudioUnitSubType_HighShelfFilter: An audio unit suitable for implementing a treble control in an audio playback or recording system.
	KAudioUnitSubType_HighShelfFilter KaudiounitsubtypePeaklimiter = 'h'<<24 | 's'<<16 | 'h'<<8 | 'f' // 'hshf'
	// KAudioUnitSubType_LowPassFilter: An audio unit that passes frequencies below a specified cutoff frequency, and blocks frequencies above that cutoff frequency.
	KAudioUnitSubType_LowPassFilter KaudiounitsubtypePeaklimiter = 'l'<<24 | 'p'<<16 | 'a'<<8 | 's' // 'lpas'
	// KAudioUnitSubType_LowShelfFilter: An audio unit suitable for implementing a bass control in an audio playback or recording system.
	KAudioUnitSubType_LowShelfFilter KaudiounitsubtypePeaklimiter = 'l'<<24 | 's'<<16 | 'h'<<8 | 'f' // 'lshf'
	KAudioUnitSubType_NBandEQ        KaudiounitsubtypePeaklimiter = 'n'<<24 | 'b'<<16 | 'e'<<8 | 'q' // 'nbeq'
	// KAudioUnitSubType_ParametricEQ: An audio unit that provides a filter whose center frequency, boost/cut level, and Q can be adjusted.
	KAudioUnitSubType_ParametricEQ KaudiounitsubtypePeaklimiter = 'p'<<24 | 'm'<<16 | 'e'<<8 | 'q' // 'pmeq'
	// KAudioUnitSubType_PeakLimiter: An audio unit that enforces an upper dynamic limit on an audio signal.
	KAudioUnitSubType_PeakLimiter KaudiounitsubtypePeaklimiter = 'l'<<24 | 'm'<<16 | 't'<<8 | 'r' // 'lmtr'
	KAudioUnitSubType_Reverb2     KaudiounitsubtypePeaklimiter = 'r'<<24 | 'v'<<16 | 'b'<<8 | '2' // 'rvb2'
	// KAudioUnitSubType_SampleDelay: An audio unit that provides a time delay for a specified number of samples.
	KAudioUnitSubType_SampleDelay KaudiounitsubtypePeaklimiter = 's'<<24 | 'd'<<16 | 'l'<<8 | 'y' // 'sdly'
)

func (e KaudiounitsubtypePeaklimiter) String() string {
	switch e {
	case KAudioUnitSubType_AUSoundIsolation:
		return "KAudioUnitSubType_AUSoundIsolation"
	case KAudioUnitSubType_BandPassFilter:
		return "KAudioUnitSubType_BandPassFilter"
	case KAudioUnitSubType_Delay:
		return "KAudioUnitSubType_Delay"
	case KAudioUnitSubType_Distortion:
		return "KAudioUnitSubType_Distortion"
	case KAudioUnitSubType_DynamicsProcessor:
		return "KAudioUnitSubType_DynamicsProcessor"
	case KAudioUnitSubType_HighPassFilter:
		return "KAudioUnitSubType_HighPassFilter"
	case KAudioUnitSubType_HighShelfFilter:
		return "KAudioUnitSubType_HighShelfFilter"
	case KAudioUnitSubType_LowPassFilter:
		return "KAudioUnitSubType_LowPassFilter"
	case KAudioUnitSubType_LowShelfFilter:
		return "KAudioUnitSubType_LowShelfFilter"
	case KAudioUnitSubType_NBandEQ:
		return "KAudioUnitSubType_NBandEQ"
	case KAudioUnitSubType_ParametricEQ:
		return "KAudioUnitSubType_ParametricEQ"
	case KAudioUnitSubType_PeakLimiter:
		return "KAudioUnitSubType_PeakLimiter"
	case KAudioUnitSubType_Reverb2:
		return "KAudioUnitSubType_Reverb2"
	case KAudioUnitSubType_SampleDelay:
		return "KAudioUnitSubType_SampleDelay"
	default:
		return fmt.Sprintf("KaudiounitsubtypePeaklimiter(%d)", e)
	}
}

type KaudiounitsubtypeSampler uint32

const ()

type KaudiounitsubtypeScheduledsoundplayer uint32

const ()

type KaudiounitsubtypeSphericalheadpanner uint32

const (
	KAudioUnitSubType_HRTFPanner          KaudiounitsubtypeSphericalheadpanner = 'h'<<24 | 'r'<<16 | 't'<<8 | 'f' // 'hrtf'
	KAudioUnitSubType_SoundFieldPanner    KaudiounitsubtypeSphericalheadpanner = 'a'<<24 | 'm'<<16 | 'b'<<8 | 'i' // 'ambi'
	KAudioUnitSubType_SphericalHeadPanner KaudiounitsubtypeSphericalheadpanner = 's'<<24 | 'p'<<16 | 'h'<<8 | 'r' // 'sphr'
	KAudioUnitSubType_VectorPanner        KaudiounitsubtypeSphericalheadpanner = 'v'<<24 | 'b'<<16 | 'a'<<8 | 's' // 'vbas'
)

func (e KaudiounitsubtypeSphericalheadpanner) String() string {
	switch e {
	case KAudioUnitSubType_HRTFPanner:
		return "KAudioUnitSubType_HRTFPanner"
	case KAudioUnitSubType_SoundFieldPanner:
		return "KAudioUnitSubType_SoundFieldPanner"
	case KAudioUnitSubType_SphericalHeadPanner:
		return "KAudioUnitSubType_SphericalHeadPanner"
	case KAudioUnitSubType_VectorPanner:
		return "KAudioUnitSubType_VectorPanner"
	default:
		return fmt.Sprintf("KaudiounitsubtypeSphericalheadpanner(%d)", e)
	}
}

type KaudiounitsubtypeStereomixer uint32

const (
	// KAudioUnitSubType_StereoMixer: An audio unit that can have any number of input buses, each of which is mono or stereo, and one stereo output bus.
	KAudioUnitSubType_StereoMixer KaudiounitsubtypeStereomixer = 's'<<24 | 'm'<<16 | 'x'<<8 | 'r' // 'smxr'
	// Deprecated.
	KAudioUnitSubType_3DMixer KaudiounitsubtypeStereomixer = '3'<<24 | 'd'<<16 | 'm'<<8 | 'x' // '3dmx'
)

func (e KaudiounitsubtypeStereomixer) String() string {
	switch e {
	case KAudioUnitSubType_StereoMixer:
		return "KAudioUnitSubType_StereoMixer"
	case KAudioUnitSubType_3DMixer:
		return "KAudioUnitSubType_3DMixer"
	default:
		return fmt.Sprintf("KaudiounitsubtypeStereomixer(%d)", e)
	}
}

type KaudiounitsubtypeTimepitch uint32

const (
	// KAudioUnitSubType_TimePitch: An audio unit that can provide independent control of playback rate and pitch.
	KAudioUnitSubType_TimePitch KaudiounitsubtypeTimepitch = 't'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'tmpt'
)

func (e KaudiounitsubtypeTimepitch) String() string {
	switch e {
	case KAudioUnitSubType_TimePitch:
		return "KAudioUnitSubType_TimePitch"
	default:
		return fmt.Sprintf("KaudiounitsubtypeTimepitch(%d)", e)
	}
}

type KaudiounittypeRemote uint32

const (
	KAudioUnitType_RemoteEffect      KaudiounittypeRemote = 'a'<<24 | 'u'<<16 | 'r'<<8 | 'x' // 'aurx'
	KAudioUnitType_RemoteGenerator   KaudiounittypeRemote = 'a'<<24 | 'u'<<16 | 'r'<<8 | 'g' // 'aurg'
	KAudioUnitType_RemoteInstrument  KaudiounittypeRemote = 'a'<<24 | 'u'<<16 | 'r'<<8 | 'i' // 'auri'
	KAudioUnitType_RemoteMusicEffect KaudiounittypeRemote = 'a'<<24 | 'u'<<16 | 'r'<<8 | 'm' // 'aurm'
)

func (e KaudiounittypeRemote) String() string {
	switch e {
	case KAudioUnitType_RemoteEffect:
		return "KAudioUnitType_RemoteEffect"
	case KAudioUnitType_RemoteGenerator:
		return "KAudioUnitType_RemoteGenerator"
	case KAudioUnitType_RemoteInstrument:
		return "KAudioUnitType_RemoteInstrument"
	case KAudioUnitType_RemoteMusicEffect:
		return "KAudioUnitType_RemoteMusicEffect"
	default:
		return fmt.Sprintf("KaudiounittypeRemote(%d)", e)
	}
}

type KaumidisynthpropertyEnablepreload uint32

const (
	KAUMIDISynthProperty_EnablePreload KaumidisynthpropertyEnablepreload = 4119
)

func (e KaumidisynthpropertyEnablepreload) String() string {
	switch e {
	case KAUMIDISynthProperty_EnablePreload:
		return "KAUMIDISynthProperty_EnablePreload"
	default:
		return fmt.Sprintf("KaumidisynthpropertyEnablepreload(%d)", e)
	}
}

type KaunbandeqfiltertypeParametric int

const (
	KAUNBandEQFilterType_2ndOrderButterworthHighPass KaunbandeqfiltertypeParametric = 2
	KAUNBandEQFilterType_2ndOrderButterworthLowPass  KaunbandeqfiltertypeParametric = 1
	KAUNBandEQFilterType_BandPass                    KaunbandeqfiltertypeParametric = 5
	KAUNBandEQFilterType_BandStop                    KaunbandeqfiltertypeParametric = 6
	KAUNBandEQFilterType_HighShelf                   KaunbandeqfiltertypeParametric = 8
	KAUNBandEQFilterType_LowShelf                    KaunbandeqfiltertypeParametric = 7
	KAUNBandEQFilterType_Parametric                  KaunbandeqfiltertypeParametric = 0
	KAUNBandEQFilterType_ResonantHighPass            KaunbandeqfiltertypeParametric = 4
	KAUNBandEQFilterType_ResonantHighShelf           KaunbandeqfiltertypeParametric = 10
	KAUNBandEQFilterType_ResonantLowPass             KaunbandeqfiltertypeParametric = 3
	KAUNBandEQFilterType_ResonantLowShelf            KaunbandeqfiltertypeParametric = 9
	KNumAUNBandEQFilterTypes                         KaunbandeqfiltertypeParametric = 11
)

func (e KaunbandeqfiltertypeParametric) String() string {
	switch e {
	case KAUNBandEQFilterType_2ndOrderButterworthHighPass:
		return "KAUNBandEQFilterType_2ndOrderButterworthHighPass"
	case KAUNBandEQFilterType_2ndOrderButterworthLowPass:
		return "KAUNBandEQFilterType_2ndOrderButterworthLowPass"
	case KAUNBandEQFilterType_BandPass:
		return "KAUNBandEQFilterType_BandPass"
	case KAUNBandEQFilterType_BandStop:
		return "KAUNBandEQFilterType_BandStop"
	case KAUNBandEQFilterType_HighShelf:
		return "KAUNBandEQFilterType_HighShelf"
	case KAUNBandEQFilterType_LowShelf:
		return "KAUNBandEQFilterType_LowShelf"
	case KAUNBandEQFilterType_Parametric:
		return "KAUNBandEQFilterType_Parametric"
	case KAUNBandEQFilterType_ResonantHighPass:
		return "KAUNBandEQFilterType_ResonantHighPass"
	case KAUNBandEQFilterType_ResonantHighShelf:
		return "KAUNBandEQFilterType_ResonantHighShelf"
	case KAUNBandEQFilterType_ResonantLowPass:
		return "KAUNBandEQFilterType_ResonantLowPass"
	case KAUNBandEQFilterType_ResonantLowShelf:
		return "KAUNBandEQFilterType_ResonantLowShelf"
	case KNumAUNBandEQFilterTypes:
		return "KNumAUNBandEQFilterTypes"
	default:
		return fmt.Sprintf("KaunbandeqfiltertypeParametric(%d)", e)
	}
}

type KauparameterlistenerAnyparameter uint32

const (
	KAUParameterListener_AnyParameter KauparameterlistenerAnyparameter = 0xffffffff
)

func (e KauparameterlistenerAnyparameter) String() string {
	switch e {
	case KAUParameterListener_AnyParameter:
		return "KAUParameterListener_AnyParameter"
	default:
		return fmt.Sprintf("KauparameterlistenerAnyparameter(%d)", e)
	}
}

type KausamplerDefault uint32

const (
	KAUSampler_DefaultBankLSB           KausamplerDefault = 0
	KAUSampler_DefaultMelodicBankMSB    KausamplerDefault = 0x79
	KAUSampler_DefaultPercussionBankMSB KausamplerDefault = 0x78
)

func (e KausamplerDefault) String() string {
	switch e {
	case KAUSampler_DefaultBankLSB:
		return "KAUSampler_DefaultBankLSB"
	case KAUSampler_DefaultMelodicBankMSB:
		return "KAUSampler_DefaultMelodicBankMSB"
	case KAUSampler_DefaultPercussionBankMSB:
		return "KAUSampler_DefaultPercussionBankMSB"
	default:
		return fmt.Sprintf("KausamplerDefault(%d)", e)
	}
}

type KausamplerpropertyLoad uint32

const (
	KAUSamplerProperty_LoadAudioFiles KausamplerpropertyLoad = 4101
	KAUSamplerProperty_LoadInstrument KausamplerpropertyLoad = 4102
)

func (e KausamplerpropertyLoad) String() string {
	switch e {
	case KAUSamplerProperty_LoadAudioFiles:
		return "KAUSamplerProperty_LoadAudioFiles"
	case KAUSamplerProperty_LoadInstrument:
		return "KAUSamplerProperty_LoadInstrument"
	default:
		return fmt.Sprintf("KausamplerpropertyLoad(%d)", e)
	}
}

type KauvoiceioerrUnexpectednumberofinputchannels int32

const (
	// KAUVoiceIOErr_UnexpectedNumberOfInputChannels: An error that indicates that the audio unit encountered an unexpected number of input channels during initialization.
	KAUVoiceIOErr_UnexpectedNumberOfInputChannels KauvoiceioerrUnexpectednumberofinputchannels = -66784
)

func (e KauvoiceioerrUnexpectednumberofinputchannels) String() string {
	switch e {
	case KAUVoiceIOErr_UnexpectedNumberOfInputChannels:
		return "KAUVoiceIOErr_UnexpectedNumberOfInputChannels"
	default:
		return fmt.Sprintf("KauvoiceioerrUnexpectednumberofinputchannels(%d)", e)
	}
}

type KauvoiceiopropertyMutedspeechactivityeventlistener uint32

const (
	// KAUVoiceIOProperty_MutedSpeechActivityEventListener: A property to register a listener that the system calls when it detects speech while the user has the microphone muted.
	KAUVoiceIOProperty_MutedSpeechActivityEventListener KauvoiceiopropertyMutedspeechactivityeventlistener = 2106
)

func (e KauvoiceiopropertyMutedspeechactivityeventlistener) String() string {
	switch e {
	case KAUVoiceIOProperty_MutedSpeechActivityEventListener:
		return "KAUVoiceIOProperty_MutedSpeechActivityEventListener"
	default:
		return fmt.Sprintf("KauvoiceiopropertyMutedspeechactivityeventlistener(%d)", e)
	}
}

type KauvoiceiopropertyOtheraudioduckingconfiguration uint32

const (
	KAUVoiceIOProperty_OtherAudioDuckingConfiguration KauvoiceiopropertyOtheraudioduckingconfiguration = 2108
)

func (e KauvoiceiopropertyOtheraudioduckingconfiguration) String() string {
	switch e {
	case KAUVoiceIOProperty_OtherAudioDuckingConfiguration:
		return "KAUVoiceIOProperty_OtherAudioDuckingConfiguration"
	default:
		return fmt.Sprintf("KauvoiceiopropertyOtheraudioduckingconfiguration(%d)", e)
	}
}

type KauvoiceiopropertyVoiceprocessingquality uint32

const (
	KAUVoiceIOProperty_VoiceProcessingQuality KauvoiceiopropertyVoiceprocessingquality = 2103
)

func (e KauvoiceiopropertyVoiceprocessingquality) String() string {
	switch e {
	case KAUVoiceIOProperty_VoiceProcessingQuality:
		return "KAUVoiceIOProperty_VoiceProcessingQuality"
	default:
		return fmt.Sprintf("KauvoiceiopropertyVoiceprocessingquality(%d)", e)
	}
}

type KcafFile uint32

const (
	KCAF_FileType            KcafFile = 'c'<<24 | 'a'<<16 | 'f'<<8 | 'f' // 'caff'
	KCAF_FileVersion_Initial KcafFile = 1
)

func (e KcafFile) String() string {
	switch e {
	case KCAF_FileType:
		return "KCAF_FileType"
	case KCAF_FileVersion_Initial:
		return "KCAF_FileVersion_Initial"
	default:
		return fmt.Sprintf("KcafFile(%d)", e)
	}
}

type KcafSmpte uint32

const (
	KCAF_SMPTE_TimeType2398     KcafSmpte = 12
	KCAF_SMPTE_TimeType24       KcafSmpte = 1
	KCAF_SMPTE_TimeType25       KcafSmpte = 2
	KCAF_SMPTE_TimeType2997     KcafSmpte = 5
	KCAF_SMPTE_TimeType2997Drop KcafSmpte = 6
	KCAF_SMPTE_TimeType30       KcafSmpte = 4
	KCAF_SMPTE_TimeType30Drop   KcafSmpte = 3
	KCAF_SMPTE_TimeType50       KcafSmpte = 11
	KCAF_SMPTE_TimeType5994     KcafSmpte = 8
	KCAF_SMPTE_TimeType5994Drop KcafSmpte = 10
	KCAF_SMPTE_TimeType60       KcafSmpte = 7
	KCAF_SMPTE_TimeType60Drop   KcafSmpte = 9
	KCAF_SMPTE_TimeTypeNone     KcafSmpte = 0
)

func (e KcafSmpte) String() string {
	switch e {
	case KCAF_SMPTE_TimeType2398:
		return "KCAF_SMPTE_TimeType2398"
	case KCAF_SMPTE_TimeType24:
		return "KCAF_SMPTE_TimeType24"
	case KCAF_SMPTE_TimeType25:
		return "KCAF_SMPTE_TimeType25"
	case KCAF_SMPTE_TimeType2997:
		return "KCAF_SMPTE_TimeType2997"
	case KCAF_SMPTE_TimeType2997Drop:
		return "KCAF_SMPTE_TimeType2997Drop"
	case KCAF_SMPTE_TimeType30:
		return "KCAF_SMPTE_TimeType30"
	case KCAF_SMPTE_TimeType30Drop:
		return "KCAF_SMPTE_TimeType30Drop"
	case KCAF_SMPTE_TimeType50:
		return "KCAF_SMPTE_TimeType50"
	case KCAF_SMPTE_TimeType5994:
		return "KCAF_SMPTE_TimeType5994"
	case KCAF_SMPTE_TimeType5994Drop:
		return "KCAF_SMPTE_TimeType5994Drop"
	case KCAF_SMPTE_TimeType60:
		return "KCAF_SMPTE_TimeType60"
	case KCAF_SMPTE_TimeType60Drop:
		return "KCAF_SMPTE_TimeType60Drop"
	case KCAF_SMPTE_TimeTypeNone:
		return "KCAF_SMPTE_TimeTypeNone"
	default:
		return fmt.Sprintf("KcafSmpte(%d)", e)
	}
}

type KgraphiceqparamNumberofbands uint32

const (
	KGraphicEQParam_NumberOfBands KgraphiceqparamNumberofbands = 10000
)

func (e KgraphiceqparamNumberofbands) String() string {
	switch e {
	case KGraphicEQParam_NumberOfBands:
		return "KGraphicEQParam_NumberOfBands"
	default:
		return fmt.Sprintf("KgraphiceqparamNumberofbands(%d)", e)
	}
}

type KhaloutputparamVolume uint32

const (
	KHALOutputParam_Volume KhaloutputparamVolume = 14
)

func (e KhaloutputparamVolume) String() string {
	switch e {
	case KHALOutputParam_Volume:
		return "KHALOutputParam_Volume"
	default:
		return fmt.Sprintf("KhaloutputparamVolume(%d)", e)
	}
}

type KmusicdevicepropertyGroupoutputbus uint32

const (
	KAudioUnitProperty_PannerMode        KmusicdevicepropertyGroupoutputbus = 3008
	KMusicDeviceProperty_GroupOutputBus  KmusicdevicepropertyGroupoutputbus = 1002
	KMusicDeviceProperty_SoundBankFSSpec KmusicdevicepropertyGroupoutputbus = 1003
)

func (e KmusicdevicepropertyGroupoutputbus) String() string {
	switch e {
	case KAudioUnitProperty_PannerMode:
		return "KAudioUnitProperty_PannerMode"
	case KMusicDeviceProperty_GroupOutputBus:
		return "KMusicDeviceProperty_GroupOutputBus"
	case KMusicDeviceProperty_SoundBankFSSpec:
		return "KMusicDeviceProperty_SoundBankFSSpec"
	default:
		return fmt.Sprintf("KmusicdevicepropertyGroupoutputbus(%d)", e)
	}
}

type KmusicdevicepropertyInstrumentcount uint32

const (
	KMusicDeviceProperty_BankName        KmusicdevicepropertyInstrumentcount = 1007
	KMusicDeviceProperty_InstrumentCount KmusicdevicepropertyInstrumentcount = 1000
	KMusicDeviceProperty_SoundBankURL    KmusicdevicepropertyInstrumentcount = 1100
)

func (e KmusicdevicepropertyInstrumentcount) String() string {
	switch e {
	case KMusicDeviceProperty_BankName:
		return "KMusicDeviceProperty_BankName"
	case KMusicDeviceProperty_InstrumentCount:
		return "KMusicDeviceProperty_InstrumentCount"
	case KMusicDeviceProperty_SoundBankURL:
		return "KMusicDeviceProperty_SoundBankURL"
	default:
		return fmt.Sprintf("KmusicdevicepropertyInstrumentcount(%d)", e)
	}
}

type KmusicdevicepropertyInstrumentname uint32

const (
	KMusicDeviceProperty_InstrumentName   KmusicdevicepropertyInstrumentname = 1001
	KMusicDeviceProperty_InstrumentNumber KmusicdevicepropertyInstrumentname = 1004
)

func (e KmusicdevicepropertyInstrumentname) String() string {
	switch e {
	case KMusicDeviceProperty_InstrumentName:
		return "KMusicDeviceProperty_InstrumentName"
	case KMusicDeviceProperty_InstrumentNumber:
		return "KMusicDeviceProperty_InstrumentNumber"
	default:
		return fmt.Sprintf("KmusicdevicepropertyInstrumentname(%d)", e)
	}
}

type KmusicdevicepropertyMidixmlnames uint32

const (
	KMusicDeviceProperty_DualSchedulingMode    KmusicdevicepropertyMidixmlnames = 1013
	KMusicDeviceProperty_MIDIXMLNames          KmusicdevicepropertyMidixmlnames = 1006
	KMusicDeviceProperty_PartGroup             KmusicdevicepropertyMidixmlnames = 1010
	KMusicDeviceProperty_SupportsStartStopNote KmusicdevicepropertyMidixmlnames = 1014
)

func (e KmusicdevicepropertyMidixmlnames) String() string {
	switch e {
	case KMusicDeviceProperty_DualSchedulingMode:
		return "KMusicDeviceProperty_DualSchedulingMode"
	case KMusicDeviceProperty_MIDIXMLNames:
		return "KMusicDeviceProperty_MIDIXMLNames"
	case KMusicDeviceProperty_PartGroup:
		return "KMusicDeviceProperty_PartGroup"
	case KMusicDeviceProperty_SupportsStartStopNote:
		return "KMusicDeviceProperty_SupportsStartStopNote"
	default:
		return fmt.Sprintf("KmusicdevicepropertyMidixmlnames(%d)", e)
	}
}

type KmusicdevicepropertyUsesinternalreverb uint32

const (
	KMusicDeviceProperty_SoundBankData      KmusicdevicepropertyUsesinternalreverb = 1008
	KMusicDeviceProperty_SoundBankFSRef     KmusicdevicepropertyUsesinternalreverb = 1012
	KMusicDeviceProperty_StreamFromDisk     KmusicdevicepropertyUsesinternalreverb = 1011
	KMusicDeviceProperty_UsesInternalReverb KmusicdevicepropertyUsesinternalreverb = 1005
)

func (e KmusicdevicepropertyUsesinternalreverb) String() string {
	switch e {
	case KMusicDeviceProperty_SoundBankData:
		return "KMusicDeviceProperty_SoundBankData"
	case KMusicDeviceProperty_SoundBankFSRef:
		return "KMusicDeviceProperty_SoundBankFSRef"
	case KMusicDeviceProperty_StreamFromDisk:
		return "KMusicDeviceProperty_StreamFromDisk"
	case KMusicDeviceProperty_UsesInternalReverb:
		return "KMusicDeviceProperty_UsesInternalReverb"
	default:
		return fmt.Sprintf("KmusicdevicepropertyUsesinternalreverb(%d)", e)
	}
}

type KmusiceventtypeExtendedcontrol uint32

const (
	KMusicEventType_ExtendedControl KmusiceventtypeExtendedcontrol = 2
)

func (e KmusiceventtypeExtendedcontrol) String() string {
	switch e {
	case KMusicEventType_ExtendedControl:
		return "KMusicEventType_ExtendedControl"
	default:
		return fmt.Sprintf("KmusiceventtypeExtendedcontrol(%d)", e)
	}
}

type KreverbparamFilter uint32

const (
	KReverbParam_FilterBandwidth KreverbparamFilter = 15
	KReverbParam_FilterEnable    KreverbparamFilter = 18
	KReverbParam_FilterFrequency KreverbparamFilter = 14
	KReverbParam_FilterGain      KreverbparamFilter = 16
	KReverbParam_FilterType      KreverbparamFilter = 17
)

func (e KreverbparamFilter) String() string {
	switch e {
	case KReverbParam_FilterBandwidth:
		return "KReverbParam_FilterBandwidth"
	case KReverbParam_FilterEnable:
		return "KReverbParam_FilterEnable"
	case KReverbParam_FilterFrequency:
		return "KReverbParam_FilterFrequency"
	case KReverbParam_FilterGain:
		return "KReverbParam_FilterGain"
	case KReverbParam_FilterType:
		return "KReverbParam_FilterType"
	default:
		return fmt.Sprintf("KreverbparamFilter(%d)", e)
	}
}

type KsampledelayparamDelayframes uint32

const (
	KSampleDelayParam_DelayFrames KsampledelayparamDelayframes = 0
)

func (e KsampledelayparamDelayframes) String() string {
	switch e {
	case KSampleDelayParam_DelayFrames:
		return "KSampleDelayParam_DelayFrames"
	default:
		return fmt.Sprintf("KsampledelayparamDelayframes(%d)", e)
	}
}

type KsystemsoundidUserpreferredalert uint32

const (
	// KSystemSoundID_FlashScreen: On the desktop, use this constant with the AudioServicesPlayAlertSound(_:) function to display a flash of light on the screen.
	KSystemSoundID_FlashScreen KsystemsoundidUserpreferredalert = 0xffe
	// KSystemSoundID_UserPreferredAlert: On the desktop, use this constant with the AudioServicesPlayAlertSound(_:) function to play the alert specified in the Sound preference pane.
	KSystemSoundID_UserPreferredAlert KsystemsoundidUserpreferredalert = 0x1000
	// KUserPreferredAlert: A deprecated sound identifier.
	KUserPreferredAlert KsystemsoundidUserpreferredalert = 4096
)

func (e KsystemsoundidUserpreferredalert) String() string {
	switch e {
	case KSystemSoundID_FlashScreen:
		return "KSystemSoundID_FlashScreen"
	case KSystemSoundID_UserPreferredAlert:
		return "KSystemSoundID_UserPreferredAlert"
	default:
		return fmt.Sprintf("KsystemsoundidUserpreferredalert(%d)", e)
	}
}

type KsystemsoundidVibrate uint32

const (
	// KSystemSoundID_Vibrate: On the iPhone, use this constant with the AudioServicesPlayAlertSound(_:) function to invoke a brief vibration.
	KSystemSoundID_Vibrate KsystemsoundidVibrate = 0xfff
)

func (e KsystemsoundidVibrate) String() string {
	switch e {
	case KSystemSoundID_Vibrate:
		return "KSystemSoundID_Vibrate"
	default:
		return fmt.Sprintf("KsystemsoundidVibrate(%d)", e)
	}
}

type KvarispeedparamPlayback uint32

const (
	KVarispeedParam_PlaybackCents KvarispeedparamPlayback = 1
	KVarispeedParam_PlaybackRate  KvarispeedparamPlayback = 0
)

func (e KvarispeedparamPlayback) String() string {
	switch e {
	case KVarispeedParam_PlaybackCents:
		return "KVarispeedParam_PlaybackCents"
	case KVarispeedParam_PlaybackRate:
		return "KVarispeedParam_PlaybackRate"
	default:
		return fmt.Sprintf("KvarispeedparamPlayback(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileFlags
type MusicSequenceFileFlags uint32

const (
	KMusicSequenceFileFlags_Default MusicSequenceFileFlags = 0
	// KMusicSequenceFileFlags_EraseFile: Specifies that an existing file should be erased when creating a new file.
	KMusicSequenceFileFlags_EraseFile MusicSequenceFileFlags = 1
)

func (e MusicSequenceFileFlags) String() string {
	switch e {
	case KMusicSequenceFileFlags_Default:
		return "KMusicSequenceFileFlags_Default"
	case KMusicSequenceFileFlags_EraseFile:
		return "KMusicSequenceFileFlags_EraseFile"
	default:
		return fmt.Sprintf("MusicSequenceFileFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileTypeID
type MusicSequenceFileTypeID uint32

const (
	KMusicSequenceFile_AnyType MusicSequenceFileTypeID = 0
	// KMusicSequenceFile_MIDIType: A MIDI file type
	KMusicSequenceFile_MIDIType MusicSequenceFileTypeID = 'm'<<24 | 'i'<<16 | 'd'<<8 | 'i' // 'midi'
	// KMusicSequenceFile_iMelodyType: An iMelody file type.
	KMusicSequenceFile_iMelodyType MusicSequenceFileTypeID = 'i'<<24 | 'm'<<16 | 'e'<<8 | 'l' // 'imel'
)

func (e MusicSequenceFileTypeID) String() string {
	switch e {
	case KMusicSequenceFile_AnyType:
		return "KMusicSequenceFile_AnyType"
	case KMusicSequenceFile_MIDIType:
		return "KMusicSequenceFile_MIDIType"
	case KMusicSequenceFile_iMelodyType:
		return "KMusicSequenceFile_iMelodyType"
	default:
		return fmt.Sprintf("MusicSequenceFileTypeID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadFlags
type MusicSequenceLoadFlags uint32

const (
	// KMusicSequenceLoadSMF_ChannelsToTracks: If this flag is set the resultant Sequence will contain a tempo track, 1 track for each MIDI Channel that is found in the SMF, 1 track for SysEx or MetaEvents - and this will be the last track in the sequence after the LoadSMFWithFlags calls.
	KMusicSequenceLoadSMF_ChannelsToTracks MusicSequenceLoadFlags = 1
	KMusicSequenceLoadSMF_PreserveTracks   MusicSequenceLoadFlags = 0
)

func (e MusicSequenceLoadFlags) String() string {
	switch e {
	case KMusicSequenceLoadSMF_ChannelsToTracks:
		return "KMusicSequenceLoadSMF_ChannelsToTracks"
	case KMusicSequenceLoadSMF_PreserveTracks:
		return "KMusicSequenceLoadSMF_PreserveTracks"
	default:
		return fmt.Sprintf("MusicSequenceLoadFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceType
type MusicSequenceType uint32

const (
	// KMusicSequenceType_Beats: Used for a music sequence that corresponds to a normal MIDI file.
	KMusicSequenceType_Beats MusicSequenceType = 'b'<<24 | 'e'<<16 | 'a'<<8 | 't' // 'beat'
	// KMusicSequenceType_Samples: Used for audio samples; a music sequence of this type cannot be saved to a MIDI file.
	KMusicSequenceType_Samples MusicSequenceType = 's'<<24 | 'a'<<16 | 'm'<<8 | 'p' // 'samp'
	// KMusicSequenceType_Seconds: Used for a music sequence that corresponds to a MIDI file, but employs SMPTE timecode.
	KMusicSequenceType_Seconds MusicSequenceType = 's'<<24 | 'e'<<16 | 'c'<<8 | 's' // 'secs'
)

func (e MusicSequenceType) String() string {
	switch e {
	case KMusicSequenceType_Beats:
		return "KMusicSequenceType_Beats"
	case KMusicSequenceType_Samples:
		return "KMusicSequenceType_Samples"
	case KMusicSequenceType_Seconds:
		return "KMusicSequenceType_Seconds"
	default:
		return fmt.Sprintf("MusicSequenceType(%d)", e)
	}
}
