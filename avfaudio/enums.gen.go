// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode
type AVAudio3DMixingPointSourceInHeadMode int

const (
	// AVAudio3DMixingPointSourceInHeadModeBypass: The point source distributes into each output channel inside the head of the listener.
	AVAudio3DMixingPointSourceInHeadModeBypass AVAudio3DMixingPointSourceInHeadMode = 1
	// AVAudio3DMixingPointSourceInHeadModeMono: The point source remains a single mono source inside the head of the listener regardless of the channels it consists of.
	AVAudio3DMixingPointSourceInHeadModeMono AVAudio3DMixingPointSourceInHeadMode = 0
)

func (e AVAudio3DMixingPointSourceInHeadMode) String() string {
	switch e {
	case AVAudio3DMixingPointSourceInHeadModeBypass:
		return "AVAudio3DMixingPointSourceInHeadModeBypass"
	case AVAudio3DMixingPointSourceInHeadModeMono:
		return "AVAudio3DMixingPointSourceInHeadModeMono"
	default:
		return fmt.Sprintf("AVAudio3DMixingPointSourceInHeadMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm
type AVAudio3DMixingRenderingAlgorithm int

const (
	// AVAudio3DMixingRenderingAlgorithmAuto: Automatically selects the highest-quality rendering algorithm available for the current playback hardware.
	AVAudio3DMixingRenderingAlgorithmAuto AVAudio3DMixingRenderingAlgorithm = 7
	// AVAudio3DMixingRenderingAlgorithmEqualPowerPanning: An algorithm that pans the data of the mixer bus into a stereo field.
	AVAudio3DMixingRenderingAlgorithmEqualPowerPanning AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmHRTF: A high-quality algorithm that uses filtering to emulate 3D space in headphones.
	AVAudio3DMixingRenderingAlgorithmHRTF AVAudio3DMixingRenderingAlgorithm = 2
	// AVAudio3DMixingRenderingAlgorithmHRTFHQ: A higher-quality head-related transfer function rendering algorithm.
	AVAudio3DMixingRenderingAlgorithmHRTFHQ AVAudio3DMixingRenderingAlgorithm = 6
	// AVAudio3DMixingRenderingAlgorithmSoundField: An algorithm that renders to multichannel hardware.
	AVAudio3DMixingRenderingAlgorithmSoundField AVAudio3DMixingRenderingAlgorithm = 3
	// AVAudio3DMixingRenderingAlgorithmSphericalHead: An algorithm that emulates 3D space in headphones by simulating interaural time delays and other spatial cues.
	AVAudio3DMixingRenderingAlgorithmSphericalHead AVAudio3DMixingRenderingAlgorithm = 1
	// AVAudio3DMixingRenderingAlgorithmStereoPassThrough: An algorithm to use when the source data doesn’t need localization.
	AVAudio3DMixingRenderingAlgorithmStereoPassThrough AVAudio3DMixingRenderingAlgorithm = 5
)

func (e AVAudio3DMixingRenderingAlgorithm) String() string {
	switch e {
	case AVAudio3DMixingRenderingAlgorithmAuto:
		return "AVAudio3DMixingRenderingAlgorithmAuto"
	case AVAudio3DMixingRenderingAlgorithmEqualPowerPanning:
		return "AVAudio3DMixingRenderingAlgorithmEqualPowerPanning"
	case AVAudio3DMixingRenderingAlgorithmHRTF:
		return "AVAudio3DMixingRenderingAlgorithmHRTF"
	case AVAudio3DMixingRenderingAlgorithmHRTFHQ:
		return "AVAudio3DMixingRenderingAlgorithmHRTFHQ"
	case AVAudio3DMixingRenderingAlgorithmSoundField:
		return "AVAudio3DMixingRenderingAlgorithmSoundField"
	case AVAudio3DMixingRenderingAlgorithmSphericalHead:
		return "AVAudio3DMixingRenderingAlgorithmSphericalHead"
	case AVAudio3DMixingRenderingAlgorithmStereoPassThrough:
		return "AVAudio3DMixingRenderingAlgorithmStereoPassThrough"
	default:
		return fmt.Sprintf("AVAudio3DMixingRenderingAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode
type AVAudio3DMixingSourceMode int

const (
	// AVAudio3DMixingSourceModeAmbienceBed: The input channels spread around the listener as far-field sources that anchor to global space.
	AVAudio3DMixingSourceModeAmbienceBed AVAudio3DMixingSourceMode = 3
	// AVAudio3DMixingSourceModeBypass: A mode that does no spatial rendering.
	AVAudio3DMixingSourceModeBypass AVAudio3DMixingSourceMode = 1
	// AVAudio3DMixingSourceModePointSource: All channels of the bus render as a single source at the location of the source node.
	AVAudio3DMixingSourceModePointSource AVAudio3DMixingSourceMode = 2
	// AVAudio3DMixingSourceModeSpatializeIfMono: A mono input bus that renders as a point source at the location of the source node.
	AVAudio3DMixingSourceModeSpatializeIfMono AVAudio3DMixingSourceMode = 0
)

func (e AVAudio3DMixingSourceMode) String() string {
	switch e {
	case AVAudio3DMixingSourceModeAmbienceBed:
		return "AVAudio3DMixingSourceModeAmbienceBed"
	case AVAudio3DMixingSourceModeBypass:
		return "AVAudio3DMixingSourceModeBypass"
	case AVAudio3DMixingSourceModePointSource:
		return "AVAudio3DMixingSourceModePointSource"
	case AVAudio3DMixingSourceModeSpatializeIfMono:
		return "AVAudio3DMixingSourceModeSpatializeIfMono"
	default:
		return fmt.Sprintf("AVAudio3DMixingSourceMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum
type AVAudioApplicationMicrophoneInjectionPermission int

const (
	// AVAudioApplicationMicrophoneInjectionPermissionDenied: A person denies the app permission to add audio to calls.
	AVAudioApplicationMicrophoneInjectionPermissionDenied AVAudioApplicationMicrophoneInjectionPermission = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'y' // 'deny'
	// AVAudioApplicationMicrophoneInjectionPermissionGranted: A person grants the app permission to add audio to calls.
	AVAudioApplicationMicrophoneInjectionPermissionGranted AVAudioApplicationMicrophoneInjectionPermission = 'g'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'grnt'
	// AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled: A person disables this service for all apps.
	AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled AVAudioApplicationMicrophoneInjectionPermission = 's'<<24 | 'r'<<16 | 'd'<<8 | 's' // 'srds'
	// AVAudioApplicationMicrophoneInjectionPermissionUndetermined: The app hasn’t requested a person’s permission to add audio to calls.
	AVAudioApplicationMicrophoneInjectionPermissionUndetermined AVAudioApplicationMicrophoneInjectionPermission = 'u'<<24 | 'n'<<16 | 'd'<<8 | 't' // 'undt'
)

func (e AVAudioApplicationMicrophoneInjectionPermission) String() string {
	switch e {
	case AVAudioApplicationMicrophoneInjectionPermissionDenied:
		return "AVAudioApplicationMicrophoneInjectionPermissionDenied"
	case AVAudioApplicationMicrophoneInjectionPermissionGranted:
		return "AVAudioApplicationMicrophoneInjectionPermissionGranted"
	case AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled:
		return "AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled"
	case AVAudioApplicationMicrophoneInjectionPermissionUndetermined:
		return "AVAudioApplicationMicrophoneInjectionPermissionUndetermined"
	default:
		return fmt.Sprintf("AVAudioApplicationMicrophoneInjectionPermission(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum
type AVAudioApplicationRecordPermission int

const (
	// AVAudioApplicationRecordPermissionDenied: Indicates the user denies the app permission to record audio.
	AVAudioApplicationRecordPermissionDenied AVAudioApplicationRecordPermission = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'y' // 'deny'
	// AVAudioApplicationRecordPermissionGranted: Indicates the user grants the app permission to record audio.
	AVAudioApplicationRecordPermissionGranted AVAudioApplicationRecordPermission = 'g'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'grnt'
	// AVAudioApplicationRecordPermissionUndetermined: Indicates the app hasn’t requested recording permission.
	AVAudioApplicationRecordPermissionUndetermined AVAudioApplicationRecordPermission = 'u'<<24 | 'n'<<16 | 'd'<<8 | 't' // 'undt'
)

func (e AVAudioApplicationRecordPermission) String() string {
	switch e {
	case AVAudioApplicationRecordPermissionDenied:
		return "AVAudioApplicationRecordPermissionDenied"
	case AVAudioApplicationRecordPermissionGranted:
		return "AVAudioApplicationRecordPermissionGranted"
	case AVAudioApplicationRecordPermissionUndetermined:
		return "AVAudioApplicationRecordPermissionUndetermined"
	default:
		return fmt.Sprintf("AVAudioApplicationRecordPermission(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
type AVAudioCommonFormat uint

const (
	// AVAudioOtherFormat: A format other than one the enumeration specifies.
	AVAudioOtherFormat AVAudioCommonFormat = 0
	// AVAudioPCMFormatFloat32: A format that represents the standard format as native-endian floats.
	AVAudioPCMFormatFloat32 AVAudioCommonFormat = 1
	// AVAudioPCMFormatFloat64: A format that represents native-endian doubles.
	AVAudioPCMFormatFloat64 AVAudioCommonFormat = 2
	// AVAudioPCMFormatInt16: A format that represents signed 16-bit native-endian integers.
	AVAudioPCMFormatInt16 AVAudioCommonFormat = 3
	// AVAudioPCMFormatInt32: A format that represents signed 32-bit native-endian integers.
	AVAudioPCMFormatInt32 AVAudioCommonFormat = 4
)

func (e AVAudioCommonFormat) String() string {
	switch e {
	case AVAudioOtherFormat:
		return "AVAudioOtherFormat"
	case AVAudioPCMFormatFloat32:
		return "AVAudioPCMFormatFloat32"
	case AVAudioPCMFormatFloat64:
		return "AVAudioPCMFormatFloat64"
	case AVAudioPCMFormatInt16:
		return "AVAudioPCMFormatInt16"
	case AVAudioPCMFormatInt32:
		return "AVAudioPCMFormatInt32"
	default:
		return fmt.Sprintf("AVAudioCommonFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource
type AVAudioContentSource int

const (
	AVAudioContentSource_AV_Spatial_Live AVAudioContentSource = 41
	AVAudioContentSource_AV_Spatial_Offline AVAudioContentSource = 39
	AVAudioContentSource_AV_Traditional_Live AVAudioContentSource = 40
	AVAudioContentSource_AV_Traditional_Offline AVAudioContentSource = 38
	AVAudioContentSource_AppleAV_Spatial_Live AVAudioContentSource = 9
	AVAudioContentSource_AppleAV_Spatial_Offline AVAudioContentSource = 7
	AVAudioContentSource_AppleAV_Traditional_Live AVAudioContentSource = 8
	AVAudioContentSource_AppleAV_Traditional_Offline AVAudioContentSource = 6
	AVAudioContentSource_AppleCapture_Spatial AVAudioContentSource = 2
	AVAudioContentSource_AppleCapture_Spatial_Enhanced AVAudioContentSource = 3
	AVAudioContentSource_AppleCapture_Traditional AVAudioContentSource = 1
	AVAudioContentSource_AppleMusic_Spatial AVAudioContentSource = 5
	AVAudioContentSource_AppleMusic_Traditional AVAudioContentSource = 4
	AVAudioContentSource_ApplePassthrough AVAudioContentSource = 10
	AVAudioContentSource_Capture_Spatial AVAudioContentSource = 34
	AVAudioContentSource_Capture_Spatial_Enhanced AVAudioContentSource = 35
	AVAudioContentSource_Capture_Traditional AVAudioContentSource = 33
	AVAudioContentSource_Music_Spatial AVAudioContentSource = 37
	AVAudioContentSource_Music_Traditional AVAudioContentSource = 36
	AVAudioContentSource_Passthrough AVAudioContentSource = 42
	AVAudioContentSource_Reserved AVAudioContentSource = 0
	AVAudioContentSource_Unspecified AVAudioContentSource = -1
)

func (e AVAudioContentSource) String() string {
	switch e {
	case AVAudioContentSource_AV_Spatial_Live:
		return "AVAudioContentSource_AV_Spatial_Live"
	case AVAudioContentSource_AV_Spatial_Offline:
		return "AVAudioContentSource_AV_Spatial_Offline"
	case AVAudioContentSource_AV_Traditional_Live:
		return "AVAudioContentSource_AV_Traditional_Live"
	case AVAudioContentSource_AV_Traditional_Offline:
		return "AVAudioContentSource_AV_Traditional_Offline"
	case AVAudioContentSource_AppleAV_Spatial_Live:
		return "AVAudioContentSource_AppleAV_Spatial_Live"
	case AVAudioContentSource_AppleAV_Spatial_Offline:
		return "AVAudioContentSource_AppleAV_Spatial_Offline"
	case AVAudioContentSource_AppleAV_Traditional_Live:
		return "AVAudioContentSource_AppleAV_Traditional_Live"
	case AVAudioContentSource_AppleAV_Traditional_Offline:
		return "AVAudioContentSource_AppleAV_Traditional_Offline"
	case AVAudioContentSource_AppleCapture_Spatial:
		return "AVAudioContentSource_AppleCapture_Spatial"
	case AVAudioContentSource_AppleCapture_Spatial_Enhanced:
		return "AVAudioContentSource_AppleCapture_Spatial_Enhanced"
	case AVAudioContentSource_AppleCapture_Traditional:
		return "AVAudioContentSource_AppleCapture_Traditional"
	case AVAudioContentSource_AppleMusic_Spatial:
		return "AVAudioContentSource_AppleMusic_Spatial"
	case AVAudioContentSource_AppleMusic_Traditional:
		return "AVAudioContentSource_AppleMusic_Traditional"
	case AVAudioContentSource_ApplePassthrough:
		return "AVAudioContentSource_ApplePassthrough"
	case AVAudioContentSource_Capture_Spatial:
		return "AVAudioContentSource_Capture_Spatial"
	case AVAudioContentSource_Capture_Spatial_Enhanced:
		return "AVAudioContentSource_Capture_Spatial_Enhanced"
	case AVAudioContentSource_Capture_Traditional:
		return "AVAudioContentSource_Capture_Traditional"
	case AVAudioContentSource_Music_Spatial:
		return "AVAudioContentSource_Music_Spatial"
	case AVAudioContentSource_Music_Traditional:
		return "AVAudioContentSource_Music_Traditional"
	case AVAudioContentSource_Passthrough:
		return "AVAudioContentSource_Passthrough"
	case AVAudioContentSource_Reserved:
		return "AVAudioContentSource_Reserved"
	case AVAudioContentSource_Unspecified:
		return "AVAudioContentSource_Unspecified"
	default:
		return fmt.Sprintf("AVAudioContentSource(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus
type AVAudioConverterInputStatus int

const (
	// AVAudioConverterInputStatus_EndOfStream: A status that indicates you’re at the end of an audio stream.
	AVAudioConverterInputStatus_EndOfStream AVAudioConverterInputStatus = 2
	// AVAudioConverterInputStatus_HaveData: A status that indicates the normal case where you supply data to the converter.
	AVAudioConverterInputStatus_HaveData AVAudioConverterInputStatus = 0
	// AVAudioConverterInputStatus_NoDataNow: A status that indicates you’re out of data.
	AVAudioConverterInputStatus_NoDataNow AVAudioConverterInputStatus = 1
)

func (e AVAudioConverterInputStatus) String() string {
	switch e {
	case AVAudioConverterInputStatus_EndOfStream:
		return "AVAudioConverterInputStatus_EndOfStream"
	case AVAudioConverterInputStatus_HaveData:
		return "AVAudioConverterInputStatus_HaveData"
	case AVAudioConverterInputStatus_NoDataNow:
		return "AVAudioConverterInputStatus_NoDataNow"
	default:
		return fmt.Sprintf("AVAudioConverterInputStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus
type AVAudioConverterOutputStatus int

const (
	// AVAudioConverterOutputStatus_EndOfStream: A status that indicates the method reaches the end of the stream, and doesn’t return any data.
	AVAudioConverterOutputStatus_EndOfStream AVAudioConverterOutputStatus = 2
	// AVAudioConverterOutputStatus_Error: A status that indicates the method encounters an error.
	AVAudioConverterOutputStatus_Error AVAudioConverterOutputStatus = 3
	// AVAudioConverterOutputStatus_HaveData: A status that indicates that the method returns all of the requested data.
	AVAudioConverterOutputStatus_HaveData AVAudioConverterOutputStatus = 0
	// AVAudioConverterOutputStatus_InputRanDry: A status that indicates the method doesn’t have enough input available to satisfy the request.
	AVAudioConverterOutputStatus_InputRanDry AVAudioConverterOutputStatus = 1
)

func (e AVAudioConverterOutputStatus) String() string {
	switch e {
	case AVAudioConverterOutputStatus_EndOfStream:
		return "AVAudioConverterOutputStatus_EndOfStream"
	case AVAudioConverterOutputStatus_Error:
		return "AVAudioConverterOutputStatus_Error"
	case AVAudioConverterOutputStatus_HaveData:
		return "AVAudioConverterOutputStatus_HaveData"
	case AVAudioConverterOutputStatus_InputRanDry:
		return "AVAudioConverterOutputStatus_InputRanDry"
	default:
		return fmt.Sprintf("AVAudioConverterOutputStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeMethod
type AVAudioConverterPrimeMethod int

const (
	// AVAudioConverterPrimeMethod_None: An option to prime the converter assumes leading and trailing frames are silence.
	AVAudioConverterPrimeMethod_None AVAudioConverterPrimeMethod = 2
	// AVAudioConverterPrimeMethod_Normal: An option to prime with trailing (zero latency) frames where the converter assumes the leading frames are silence.
	AVAudioConverterPrimeMethod_Normal AVAudioConverterPrimeMethod = 1
	// AVAudioConverterPrimeMethod_Pre: An option to prime with leading and trailing input frames.
	AVAudioConverterPrimeMethod_Pre AVAudioConverterPrimeMethod = 0
)

func (e AVAudioConverterPrimeMethod) String() string {
	switch e {
	case AVAudioConverterPrimeMethod_None:
		return "AVAudioConverterPrimeMethod_None"
	case AVAudioConverterPrimeMethod_Normal:
		return "AVAudioConverterPrimeMethod_Normal"
	case AVAudioConverterPrimeMethod_Pre:
		return "AVAudioConverterPrimeMethod_Pre"
	default:
		return fmt.Sprintf("AVAudioConverterPrimeMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration
type AVAudioDynamicRangeControlConfiguration int

const (
	AVAudioDynamicRangeControlConfiguration_Capture AVAudioDynamicRangeControlConfiguration = 4
	AVAudioDynamicRangeControlConfiguration_Movie AVAudioDynamicRangeControlConfiguration = 3
	AVAudioDynamicRangeControlConfiguration_Music AVAudioDynamicRangeControlConfiguration = 1
	AVAudioDynamicRangeControlConfiguration_None AVAudioDynamicRangeControlConfiguration = 0
	AVAudioDynamicRangeControlConfiguration_Speech AVAudioDynamicRangeControlConfiguration = 2
)

func (e AVAudioDynamicRangeControlConfiguration) String() string {
	switch e {
	case AVAudioDynamicRangeControlConfiguration_Capture:
		return "AVAudioDynamicRangeControlConfiguration_Capture"
	case AVAudioDynamicRangeControlConfiguration_Movie:
		return "AVAudioDynamicRangeControlConfiguration_Movie"
	case AVAudioDynamicRangeControlConfiguration_Music:
		return "AVAudioDynamicRangeControlConfiguration_Music"
	case AVAudioDynamicRangeControlConfiguration_None:
		return "AVAudioDynamicRangeControlConfiguration_None"
	case AVAudioDynamicRangeControlConfiguration_Speech:
		return "AVAudioDynamicRangeControlConfiguration_Speech"
	default:
		return fmt.Sprintf("AVAudioDynamicRangeControlConfiguration(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingError
type AVAudioEngineManualRenderingError int

const (
	// AVAudioEngineManualRenderingErrorInitialized: An operation that the system can’t perform because the engine is still running.
	AVAudioEngineManualRenderingErrorInitialized AVAudioEngineManualRenderingError = -80801
	// AVAudioEngineManualRenderingErrorInvalidMode: An operation the system can’t perform because the engine isn’t in manual rendering mode or the right variant of it.
	AVAudioEngineManualRenderingErrorInvalidMode AVAudioEngineManualRenderingError = -80800
	// AVAudioEngineManualRenderingErrorNotRunning: An operation the system can’t perform because the engine isn’t running.
	AVAudioEngineManualRenderingErrorNotRunning AVAudioEngineManualRenderingError = -80802
)

func (e AVAudioEngineManualRenderingError) String() string {
	switch e {
	case AVAudioEngineManualRenderingErrorInitialized:
		return "AVAudioEngineManualRenderingErrorInitialized"
	case AVAudioEngineManualRenderingErrorInvalidMode:
		return "AVAudioEngineManualRenderingErrorInvalidMode"
	case AVAudioEngineManualRenderingErrorNotRunning:
		return "AVAudioEngineManualRenderingErrorNotRunning"
	default:
		return fmt.Sprintf("AVAudioEngineManualRenderingError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingMode
type AVAudioEngineManualRenderingMode int

const (
	// AVAudioEngineManualRenderingModeOffline: An engine that operates in an offline mode.
	AVAudioEngineManualRenderingModeOffline AVAudioEngineManualRenderingMode = 0
	// AVAudioEngineManualRenderingModeRealtime: An engine that operates under real-time constraints and doesn’t make blocking calls while rendering.
	AVAudioEngineManualRenderingModeRealtime AVAudioEngineManualRenderingMode = 1
)

func (e AVAudioEngineManualRenderingMode) String() string {
	switch e {
	case AVAudioEngineManualRenderingModeOffline:
		return "AVAudioEngineManualRenderingModeOffline"
	case AVAudioEngineManualRenderingModeRealtime:
		return "AVAudioEngineManualRenderingModeRealtime"
	default:
		return fmt.Sprintf("AVAudioEngineManualRenderingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus
type AVAudioEngineManualRenderingStatus int

const (
	// AVAudioEngineManualRenderingStatusCannotDoInCurrentContext: An operation that the system can’t perform under the current conditions.
	AVAudioEngineManualRenderingStatusCannotDoInCurrentContext AVAudioEngineManualRenderingStatus = 2
	// AVAudioEngineManualRenderingStatusError: A problem that occurs during rendering and results in no data returning.
	AVAudioEngineManualRenderingStatusError AVAudioEngineManualRenderingStatus = -1
	// AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode: A condition that occurs when the input node doesn’t return enough input data to satisfy the render request at the time of the request.
	AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode AVAudioEngineManualRenderingStatus = 1
	// AVAudioEngineManualRenderingStatusSuccess: A status that indicates the successful return of the requested data.
	AVAudioEngineManualRenderingStatusSuccess AVAudioEngineManualRenderingStatus = 0
)

func (e AVAudioEngineManualRenderingStatus) String() string {
	switch e {
	case AVAudioEngineManualRenderingStatusCannotDoInCurrentContext:
		return "AVAudioEngineManualRenderingStatusCannotDoInCurrentContext"
	case AVAudioEngineManualRenderingStatusError:
		return "AVAudioEngineManualRenderingStatusError"
	case AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode:
		return "AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode"
	case AVAudioEngineManualRenderingStatusSuccess:
		return "AVAudioEngineManualRenderingStatusSuccess"
	default:
		return fmt.Sprintf("AVAudioEngineManualRenderingStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationModel
type AVAudioEnvironmentDistanceAttenuationModel int

const (
	// AVAudioEnvironmentDistanceAttenuationModelExponential: An exponential model that describes the drop-off in gain as the source moves away from the listener.
	AVAudioEnvironmentDistanceAttenuationModelExponential AVAudioEnvironmentDistanceAttenuationModel = 1
	// AVAudioEnvironmentDistanceAttenuationModelInverse: An inverse model that describes the drop-off in gain as the source moves away from the listener.
	AVAudioEnvironmentDistanceAttenuationModelInverse AVAudioEnvironmentDistanceAttenuationModel = 2
	// AVAudioEnvironmentDistanceAttenuationModelLinear: A linear model that describes the drop-off in gain as the source moves away from the listener.
	AVAudioEnvironmentDistanceAttenuationModelLinear AVAudioEnvironmentDistanceAttenuationModel = 3
)

func (e AVAudioEnvironmentDistanceAttenuationModel) String() string {
	switch e {
	case AVAudioEnvironmentDistanceAttenuationModelExponential:
		return "AVAudioEnvironmentDistanceAttenuationModelExponential"
	case AVAudioEnvironmentDistanceAttenuationModelInverse:
		return "AVAudioEnvironmentDistanceAttenuationModelInverse"
	case AVAudioEnvironmentDistanceAttenuationModelLinear:
		return "AVAudioEnvironmentDistanceAttenuationModelLinear"
	default:
		return fmt.Sprintf("AVAudioEnvironmentDistanceAttenuationModel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType
type AVAudioEnvironmentOutputType int

const (
	// AVAudioEnvironmentOutputTypeAuto: Automatically detects the playback route and picks the correct output.
	AVAudioEnvironmentOutputTypeAuto AVAudioEnvironmentOutputType = 0
	// AVAudioEnvironmentOutputTypeBuiltInSpeakers: Renders the audio output for built-in speakers on the current hardware.
	AVAudioEnvironmentOutputTypeBuiltInSpeakers AVAudioEnvironmentOutputType = 2
	// AVAudioEnvironmentOutputTypeExternalSpeakers: Renders the audio output for external speakers according to the audio environment node’s output channel layout.
	AVAudioEnvironmentOutputTypeExternalSpeakers AVAudioEnvironmentOutputType = 3
	// AVAudioEnvironmentOutputTypeHeadphones: Renders the audio output for headphones.
	AVAudioEnvironmentOutputTypeHeadphones AVAudioEnvironmentOutputType = 1
)

func (e AVAudioEnvironmentOutputType) String() string {
	switch e {
	case AVAudioEnvironmentOutputTypeAuto:
		return "AVAudioEnvironmentOutputTypeAuto"
	case AVAudioEnvironmentOutputTypeBuiltInSpeakers:
		return "AVAudioEnvironmentOutputTypeBuiltInSpeakers"
	case AVAudioEnvironmentOutputTypeExternalSpeakers:
		return "AVAudioEnvironmentOutputTypeExternalSpeakers"
	case AVAudioEnvironmentOutputTypeHeadphones:
		return "AVAudioEnvironmentOutputTypeHeadphones"
	default:
		return fmt.Sprintf("AVAudioEnvironmentOutputType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeBufferOptions
type AVAudioPlayerNodeBufferOptions uint

const (
	// AVAudioPlayerNodeBufferInterrupts: An option that indicates the buffer interrupts any buffer in a playing state.
	AVAudioPlayerNodeBufferInterrupts AVAudioPlayerNodeBufferOptions = 2
	// AVAudioPlayerNodeBufferInterruptsAtLoop: An option that indicates the buffer interrupts any buffer in a playing state at its loop point.
	AVAudioPlayerNodeBufferInterruptsAtLoop AVAudioPlayerNodeBufferOptions = 4
	// AVAudioPlayerNodeBufferLoops: An option that indicates the buffer loops indefinitely.
	AVAudioPlayerNodeBufferLoops AVAudioPlayerNodeBufferOptions = 1
)

func (e AVAudioPlayerNodeBufferOptions) String() string {
	switch e {
	case AVAudioPlayerNodeBufferInterrupts:
		return "AVAudioPlayerNodeBufferInterrupts"
	case AVAudioPlayerNodeBufferInterruptsAtLoop:
		return "AVAudioPlayerNodeBufferInterruptsAtLoop"
	case AVAudioPlayerNodeBufferLoops:
		return "AVAudioPlayerNodeBufferLoops"
	default:
		return fmt.Sprintf("AVAudioPlayerNodeBufferOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType
type AVAudioPlayerNodeCompletionCallbackType int

const (
	// AVAudioPlayerNodeCompletionDataConsumed: A completion handler that indicates the player consumes the buffer or file data.
	AVAudioPlayerNodeCompletionDataConsumed AVAudioPlayerNodeCompletionCallbackType = 0
	// AVAudioPlayerNodeCompletionDataPlayedBack: A completion handler that indicates the player finishes the buffer or file data.
	AVAudioPlayerNodeCompletionDataPlayedBack AVAudioPlayerNodeCompletionCallbackType = 2
	// AVAudioPlayerNodeCompletionDataRendered: A completion handler that indicates the player renders the buffer or file data.
	AVAudioPlayerNodeCompletionDataRendered AVAudioPlayerNodeCompletionCallbackType = 1
)

func (e AVAudioPlayerNodeCompletionCallbackType) String() string {
	switch e {
	case AVAudioPlayerNodeCompletionDataConsumed:
		return "AVAudioPlayerNodeCompletionDataConsumed"
	case AVAudioPlayerNodeCompletionDataPlayedBack:
		return "AVAudioPlayerNodeCompletionDataPlayedBack"
	case AVAudioPlayerNodeCompletionDataRendered:
		return "AVAudioPlayerNodeCompletionDataRendered"
	default:
		return fmt.Sprintf("AVAudioPlayerNodeCompletionCallbackType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality
type AVAudioQuality int

const (
	// AVAudioQualityHigh: A value that represents a high audio quality for encoding and conversion.
	AVAudioQualityHigh AVAudioQuality = 96
	// AVAudioQualityLow: A value that represents a low audio quality for encoding and conversion.
	AVAudioQualityLow AVAudioQuality = 32
	// AVAudioQualityMax: A value that represents a maximum audio quality for encoding and conversion.
	AVAudioQualityMax AVAudioQuality = 127
	// AVAudioQualityMedium: A value that represents a medium audio quality for encoding and conversion.
	AVAudioQualityMedium AVAudioQuality = 64
	// AVAudioQualityMin: A value that represents a minimum audio quality for encoding and conversion.
	AVAudioQualityMin AVAudioQuality = 0
)

func (e AVAudioQuality) String() string {
	switch e {
	case AVAudioQualityHigh:
		return "AVAudioQualityHigh"
	case AVAudioQualityLow:
		return "AVAudioQualityLow"
	case AVAudioQualityMax:
		return "AVAudioQualityMax"
	case AVAudioQualityMedium:
		return "AVAudioQualityMedium"
	case AVAudioQualityMin:
		return "AVAudioQualityMin"
	default:
		return fmt.Sprintf("AVAudioQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category
type AVAudioRoutingArbitrationCategory int

const (
	// AVAudioRoutingArbitrationCategoryPlayAndRecord: The app plays and records audio.
	AVAudioRoutingArbitrationCategoryPlayAndRecord AVAudioRoutingArbitrationCategory = 1
	// AVAudioRoutingArbitrationCategoryPlayAndRecordVoice: The app uses Voice over IP (VoIP).
	AVAudioRoutingArbitrationCategoryPlayAndRecordVoice AVAudioRoutingArbitrationCategory = 2
	// AVAudioRoutingArbitrationCategoryPlayback: The app plays audio.
	AVAudioRoutingArbitrationCategoryPlayback AVAudioRoutingArbitrationCategory = 0
)

func (e AVAudioRoutingArbitrationCategory) String() string {
	switch e {
	case AVAudioRoutingArbitrationCategoryPlayAndRecord:
		return "AVAudioRoutingArbitrationCategoryPlayAndRecord"
	case AVAudioRoutingArbitrationCategoryPlayAndRecordVoice:
		return "AVAudioRoutingArbitrationCategoryPlayAndRecordVoice"
	case AVAudioRoutingArbitrationCategoryPlayback:
		return "AVAudioRoutingArbitrationCategoryPlayback"
	default:
		return fmt.Sprintf("AVAudioRoutingArbitrationCategory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionActivationOptions
type AVAudioSessionActivationOptions uint

const (
	// AVAudioSessionActivationOptionNone: A value that indicates the system should activate the audio session with no options.
	AVAudioSessionActivationOptionNone AVAudioSessionActivationOptions = 0
)

func (e AVAudioSessionActivationOptions) String() string {
	switch e {
	case AVAudioSessionActivationOptionNone:
		return "AVAudioSessionActivationOptionNone"
	default:
		return fmt.Sprintf("AVAudioSessionActivationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy
type AVAudioSessionAnchoringStrategy int

const (
	AVAudioSessionAnchoringStrategyAutomatic AVAudioSessionAnchoringStrategy = 0
	AVAudioSessionAnchoringStrategyFront AVAudioSessionAnchoringStrategy = 0
	AVAudioSessionAnchoringStrategyScene AVAudioSessionAnchoringStrategy = 0
)

func (e AVAudioSessionAnchoringStrategy) String() string {
	switch e {
	case AVAudioSessionAnchoringStrategyAutomatic:
		return "AVAudioSessionAnchoringStrategyAutomatic"
	default:
		return fmt.Sprintf("AVAudioSessionAnchoringStrategy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct
type AVAudioSessionCategoryOptions uint

const (
	// AVAudioSessionCategoryOptionAllowAirPlay: An option that determines whether you can stream audio from this session to AirPlay devices.
	AVAudioSessionCategoryOptionAllowAirPlay AVAudioSessionCategoryOptions = 64
	// AVAudioSessionCategoryOptionAllowBluetoothA2DP: An option that determines whether you can stream audio from this session to Bluetooth devices that support the Advanced Audio Distribution Profile (A2DP).
	AVAudioSessionCategoryOptionAllowBluetoothA2DP AVAudioSessionCategoryOptions = 32
	// AVAudioSessionCategoryOptionAllowBluetoothHFP: An option that makes Bluetooth Hands-Free Profile (HFP) devices available for audio input.
	AVAudioSessionCategoryOptionAllowBluetoothHFP AVAudioSessionCategoryOptions = 4
	// AVAudioSessionCategoryOptionBluetoothHighQualityRecording: An option that indicates to enable high-quality audio for input and output routes.
	AVAudioSessionCategoryOptionBluetoothHighQualityRecording AVAudioSessionCategoryOptions = 524288
	// AVAudioSessionCategoryOptionDefaultToSpeaker: An option that determines whether audio from the session defaults to the built-in speaker instead of the receiver.
	AVAudioSessionCategoryOptionDefaultToSpeaker AVAudioSessionCategoryOptions = 8
	// AVAudioSessionCategoryOptionDuckOthers: An option that reduces the volume of other audio sessions while audio from this session plays.
	AVAudioSessionCategoryOptionDuckOthers AVAudioSessionCategoryOptions = 2
	// AVAudioSessionCategoryOptionFarFieldInput: This option should be used if a session prefers to use FarFieldInput when available.
	AVAudioSessionCategoryOptionFarFieldInput AVAudioSessionCategoryOptions = 262144
	// AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers: An option that determines whether to pause spoken audio content from other sessions when your app plays its audio.
	AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers AVAudioSessionCategoryOptions = 17
	// AVAudioSessionCategoryOptionMixWithOthers: An option that indicates whether audio from this session mixes with audio from active sessions in other audio apps.
	AVAudioSessionCategoryOptionMixWithOthers AVAudioSessionCategoryOptions = 1
	// AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption: An option that indicates whether the system interrupts the audio session when it mutes the built-in microphone.
	AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption AVAudioSessionCategoryOptions = 128
	// Deprecated.
	AVAudioSessionCategoryOptionAllowBluetooth AVAudioSessionCategoryOptions = 4
)

func (e AVAudioSessionCategoryOptions) String() string {
	switch e {
	case AVAudioSessionCategoryOptionAllowAirPlay:
		return "AVAudioSessionCategoryOptionAllowAirPlay"
	case AVAudioSessionCategoryOptionAllowBluetoothA2DP:
		return "AVAudioSessionCategoryOptionAllowBluetoothA2DP"
	case AVAudioSessionCategoryOptionAllowBluetoothHFP:
		return "AVAudioSessionCategoryOptionAllowBluetoothHFP"
	case AVAudioSessionCategoryOptionBluetoothHighQualityRecording:
		return "AVAudioSessionCategoryOptionBluetoothHighQualityRecording"
	case AVAudioSessionCategoryOptionDefaultToSpeaker:
		return "AVAudioSessionCategoryOptionDefaultToSpeaker"
	case AVAudioSessionCategoryOptionDuckOthers:
		return "AVAudioSessionCategoryOptionDuckOthers"
	case AVAudioSessionCategoryOptionFarFieldInput:
		return "AVAudioSessionCategoryOptionFarFieldInput"
	case AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers:
		return "AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers"
	case AVAudioSessionCategoryOptionMixWithOthers:
		return "AVAudioSessionCategoryOptionMixWithOthers"
	case AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption:
		return "AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption"
	default:
		return fmt.Sprintf("AVAudioSessionCategoryOptions(%d)", e)
	}
}

type AVAudioSessionErrorCode int

const (
	AVAudioSessionErrorCodeBadParam AVAudioSessionErrorCode = -50
	AVAudioSessionErrorCodeCannotInterruptOthers AVAudioSessionErrorCode = '!'<<24 | 'i'<<16 | 'n'<<8 | 't' // '!int'
	AVAudioSessionErrorCodeCannotStartPlaying AVAudioSessionErrorCode = '!'<<24 | 'p'<<16 | 'l'<<8 | 'a' // '!pla'
	AVAudioSessionErrorCodeCannotStartRecording AVAudioSessionErrorCode = '!'<<24 | 'r'<<16 | 'e'<<8 | 'c' // '!rec'
	AVAudioSessionErrorCodeExpiredSession AVAudioSessionErrorCode = '!'<<24 | 's'<<16 | 'e'<<8 | 's' // '!ses'
	AVAudioSessionErrorCodeIncompatibleCategory AVAudioSessionErrorCode = '!'<<24 | 'c'<<16 | 'a'<<8 | 't' // '!cat'
	AVAudioSessionErrorCodeInsufficientPriority AVAudioSessionErrorCode = '!'<<24 | 'p'<<16 | 'r'<<8 | 'i' // '!pri'
	AVAudioSessionErrorCodeIsBusy AVAudioSessionErrorCode = '!'<<24 | 'a'<<16 | 'c'<<8 | 't' // '!act'
	AVAudioSessionErrorCodeMediaServicesFailed AVAudioSessionErrorCode = 'm'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'msrv'
	AVAudioSessionErrorCodeMissingEntitlement AVAudioSessionErrorCode = 'e'<<24 | 'n'<<16 | 't'<<8 | '?' // 'ent?'
	AVAudioSessionErrorCodeNone AVAudioSessionErrorCode = 0
	AVAudioSessionErrorCodeResourceNotAvailable AVAudioSessionErrorCode = '!'<<24 | 'r'<<16 | 'e'<<8 | 's' // '!res'
	AVAudioSessionErrorCodeSessionNotActive AVAudioSessionErrorCode = 'i'<<24 | 'n'<<16 | 'a'<<8 | 'c' // 'inac'
	AVAudioSessionErrorCodeSiriIsRecording AVAudioSessionErrorCode = 's'<<24 | 'i'<<16 | 'r'<<8 | 'i' // 'siri'
	AVAudioSessionErrorCodeUnspecified AVAudioSessionErrorCode = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
)

func (e AVAudioSessionErrorCode) String() string {
	switch e {
	case AVAudioSessionErrorCodeBadParam:
		return "AVAudioSessionErrorCodeBadParam"
	case AVAudioSessionErrorCodeCannotInterruptOthers:
		return "AVAudioSessionErrorCodeCannotInterruptOthers"
	case AVAudioSessionErrorCodeCannotStartPlaying:
		return "AVAudioSessionErrorCodeCannotStartPlaying"
	case AVAudioSessionErrorCodeCannotStartRecording:
		return "AVAudioSessionErrorCodeCannotStartRecording"
	case AVAudioSessionErrorCodeExpiredSession:
		return "AVAudioSessionErrorCodeExpiredSession"
	case AVAudioSessionErrorCodeIncompatibleCategory:
		return "AVAudioSessionErrorCodeIncompatibleCategory"
	case AVAudioSessionErrorCodeInsufficientPriority:
		return "AVAudioSessionErrorCodeInsufficientPriority"
	case AVAudioSessionErrorCodeIsBusy:
		return "AVAudioSessionErrorCodeIsBusy"
	case AVAudioSessionErrorCodeMediaServicesFailed:
		return "AVAudioSessionErrorCodeMediaServicesFailed"
	case AVAudioSessionErrorCodeMissingEntitlement:
		return "AVAudioSessionErrorCodeMissingEntitlement"
	case AVAudioSessionErrorCodeNone:
		return "AVAudioSessionErrorCodeNone"
	case AVAudioSessionErrorCodeResourceNotAvailable:
		return "AVAudioSessionErrorCodeResourceNotAvailable"
	case AVAudioSessionErrorCodeSessionNotActive:
		return "AVAudioSessionErrorCodeSessionNotActive"
	case AVAudioSessionErrorCodeSiriIsRecording:
		return "AVAudioSessionErrorCodeSiriIsRecording"
	case AVAudioSessionErrorCodeUnspecified:
		return "AVAudioSessionErrorCodeUnspecified"
	default:
		return fmt.Sprintf("AVAudioSessionErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType
type AVAudioSessionIOType uint

const (
	// AVAudioSessionIOTypeAggregated: An I/O type that indicates if audio input and output should be presented in the same realtime I/O callback.
	AVAudioSessionIOTypeAggregated AVAudioSessionIOType = 1
	// AVAudioSessionIOTypeNotSpecified: The default audio session I/O type.
	AVAudioSessionIOTypeNotSpecified AVAudioSessionIOType = 0
)

func (e AVAudioSessionIOType) String() string {
	switch e {
	case AVAudioSessionIOTypeAggregated:
		return "AVAudioSessionIOTypeAggregated"
	case AVAudioSessionIOTypeNotSpecified:
		return "AVAudioSessionIOTypeNotSpecified"
	default:
		return fmt.Sprintf("AVAudioSessionIOType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionOptions
type AVAudioSessionInterruptionOptions uint

const (
	// AVAudioSessionInterruptionOptionShouldResume: An option that indicates the interruption by another audio session has ended and the app can resume its audio session.
	AVAudioSessionInterruptionOptionShouldResume AVAudioSessionInterruptionOptions = 1
)

func (e AVAudioSessionInterruptionOptions) String() string {
	switch e {
	case AVAudioSessionInterruptionOptionShouldResume:
		return "AVAudioSessionInterruptionOptionShouldResume"
	default:
		return fmt.Sprintf("AVAudioSessionInterruptionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason
type AVAudioSessionInterruptionReason uint

const (
	// AVAudioSessionInterruptionReasonBuiltInMicMuted: The system interrupts the audio session when the device mutes the built-in microphone.
	AVAudioSessionInterruptionReasonBuiltInMicMuted AVAudioSessionInterruptionReason = 2
	// AVAudioSessionInterruptionReasonDefault: The system interrupts this audio session when it activates another.
	AVAudioSessionInterruptionReasonDefault AVAudioSessionInterruptionReason = 0
	AVAudioSessionInterruptionReasonDeviceUnauthenticated AVAudioSessionInterruptionReason = 0
	// AVAudioSessionInterruptionReasonRouteDisconnected: The system interrupts the audio session due to a disconnection of an audio route.
	AVAudioSessionInterruptionReasonRouteDisconnected AVAudioSessionInterruptionReason = 4
	// AVAudioSessionInterruptionReasonSceneWasBackgrounded: The system backgrounds the scene and interrupts the audio session.
	AVAudioSessionInterruptionReasonSceneWasBackgrounded AVAudioSessionInterruptionReason = 0
	// Deprecated.
	AVAudioSessionInterruptionReasonAppWasSuspended AVAudioSessionInterruptionReason = 1
)

func (e AVAudioSessionInterruptionReason) String() string {
	switch e {
	case AVAudioSessionInterruptionReasonBuiltInMicMuted:
		return "AVAudioSessionInterruptionReasonBuiltInMicMuted"
	case AVAudioSessionInterruptionReasonDefault:
		return "AVAudioSessionInterruptionReasonDefault"
	case AVAudioSessionInterruptionReasonRouteDisconnected:
		return "AVAudioSessionInterruptionReasonRouteDisconnected"
	case AVAudioSessionInterruptionReasonAppWasSuspended:
		return "AVAudioSessionInterruptionReasonAppWasSuspended"
	default:
		return fmt.Sprintf("AVAudioSessionInterruptionReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType
type AVAudioSessionInterruptionType uint

const (
	// AVAudioSessionInterruptionTypeBegan: A type that indicates that the operating system began interrupting the audio session.
	AVAudioSessionInterruptionTypeBegan AVAudioSessionInterruptionType = 1
	// AVAudioSessionInterruptionTypeEnded: A type that indicates that the operating system ended interrupting the audio session.
	AVAudioSessionInterruptionTypeEnded AVAudioSessionInterruptionType = 0
)

func (e AVAudioSessionInterruptionType) String() string {
	switch e {
	case AVAudioSessionInterruptionTypeBegan:
		return "AVAudioSessionInterruptionTypeBegan"
	case AVAudioSessionInterruptionTypeEnded:
		return "AVAudioSessionInterruptionTypeEnded"
	default:
		return fmt.Sprintf("AVAudioSessionInterruptionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode
type AVAudioSessionMicrophoneInjectionMode int

const (
	// AVAudioSessionMicrophoneInjectionModeNone: A mode that indicates not to use spoken audio injection.
	AVAudioSessionMicrophoneInjectionModeNone AVAudioSessionMicrophoneInjectionMode = 0
	// AVAudioSessionMicrophoneInjectionModeSpokenAudio: A mode that indicates to inject spoken audio, like synthesized speech, along with microphone audio.
	AVAudioSessionMicrophoneInjectionModeSpokenAudio AVAudioSessionMicrophoneInjectionMode = 1
)

func (e AVAudioSessionMicrophoneInjectionMode) String() string {
	switch e {
	case AVAudioSessionMicrophoneInjectionModeNone:
		return "AVAudioSessionMicrophoneInjectionModeNone"
	case AVAudioSessionMicrophoneInjectionModeSpokenAudio:
		return "AVAudioSessionMicrophoneInjectionModeSpokenAudio"
	default:
		return fmt.Sprintf("AVAudioSessionMicrophoneInjectionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PortOverride
type AVAudioSessionPortOverride uint

const (
	// AVAudioSessionPortOverrideNone: A value that indicates not to override the output audio port.
	AVAudioSessionPortOverrideNone AVAudioSessionPortOverride = 0
	// AVAudioSessionPortOverrideSpeaker: A value that indicates to override the current inputs and outputs, and route audio to the built-in speaker and microphone.
	AVAudioSessionPortOverrideSpeaker AVAudioSessionPortOverride = 's'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'spkr'
)

func (e AVAudioSessionPortOverride) String() string {
	switch e {
	case AVAudioSessionPortOverrideNone:
		return "AVAudioSessionPortOverrideNone"
	case AVAudioSessionPortOverrideSpeaker:
		return "AVAudioSessionPortOverrideSpeaker"
	default:
		return fmt.Sprintf("AVAudioSessionPortOverride(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum
type AVAudioSessionPromptStyle uint

const (
	// AVAudioSessionPromptStyleNone: Your app shouldn’t issue prompts at this time.
	AVAudioSessionPromptStyleNone AVAudioSessionPromptStyle = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'none'
	// AVAudioSessionPromptStyleNormal: Your app may use long, verbal prompts.
	AVAudioSessionPromptStyleNormal AVAudioSessionPromptStyle = 'n'<<24 | 'r'<<16 | 'm'<<8 | 'l' // 'nrml'
	// AVAudioSessionPromptStyleShort: Your app should issue short, nonverbal prompts.
	AVAudioSessionPromptStyleShort AVAudioSessionPromptStyle = 's'<<24 | 'h'<<16 | 'r'<<8 | 't' // 'shrt'
)

func (e AVAudioSessionPromptStyle) String() string {
	switch e {
	case AVAudioSessionPromptStyleNone:
		return "AVAudioSessionPromptStyleNone"
	case AVAudioSessionPromptStyleNormal:
		return "AVAudioSessionPromptStyleNormal"
	case AVAudioSessionPromptStyleShort:
		return "AVAudioSessionPromptStyleShort"
	default:
		return fmt.Sprintf("AVAudioSessionPromptStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RecordPermission-swift.enum
type AVAudioSessionRecordPermission uint

const (
	// Deprecated.
	AVAudioSessionRecordPermissionDenied AVAudioSessionRecordPermission = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'y' // 'deny'
	// Deprecated.
	AVAudioSessionRecordPermissionGranted AVAudioSessionRecordPermission = 'g'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'grnt'
	// Deprecated.
	AVAudioSessionRecordPermissionUndetermined AVAudioSessionRecordPermission = 'u'<<24 | 'n'<<16 | 'd'<<8 | 't' // 'undt'
)

func (e AVAudioSessionRecordPermission) String() string {
	switch e {
	case AVAudioSessionRecordPermissionDenied:
		return "AVAudioSessionRecordPermissionDenied"
	case AVAudioSessionRecordPermissionGranted:
		return "AVAudioSessionRecordPermissionGranted"
	case AVAudioSessionRecordPermissionUndetermined:
		return "AVAudioSessionRecordPermissionUndetermined"
	default:
		return fmt.Sprintf("AVAudioSessionRecordPermission(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum
type AVAudioSessionRenderingMode int

const (
	// AVAudioSessionRenderingModeDolbyAtmos: A mode that represents Dolby Atmos.
	AVAudioSessionRenderingModeDolbyAtmos AVAudioSessionRenderingMode = 5
	// AVAudioSessionRenderingModeDolbyAudio: A mode that represents Dolby audio.
	AVAudioSessionRenderingModeDolbyAudio AVAudioSessionRenderingMode = 4
	// AVAudioSessionRenderingModeMonoStereo: A mode that represents non multi-channel audio.
	AVAudioSessionRenderingModeMonoStereo AVAudioSessionRenderingMode = 1
	// AVAudioSessionRenderingModeNotApplicable: A mode that represents there’s no asset in a loading or playing state.
	AVAudioSessionRenderingModeNotApplicable AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeSpatialAudio: A mode that represents a fallback for when hardware capabilities don’t support Dolby.
	AVAudioSessionRenderingModeSpatialAudio AVAudioSessionRenderingMode = 3
	// AVAudioSessionRenderingModeSurround: A mode that represents general multi-channel audio.
	AVAudioSessionRenderingModeSurround AVAudioSessionRenderingMode = 2
)

func (e AVAudioSessionRenderingMode) String() string {
	switch e {
	case AVAudioSessionRenderingModeDolbyAtmos:
		return "AVAudioSessionRenderingModeDolbyAtmos"
	case AVAudioSessionRenderingModeDolbyAudio:
		return "AVAudioSessionRenderingModeDolbyAudio"
	case AVAudioSessionRenderingModeMonoStereo:
		return "AVAudioSessionRenderingModeMonoStereo"
	case AVAudioSessionRenderingModeNotApplicable:
		return "AVAudioSessionRenderingModeNotApplicable"
	case AVAudioSessionRenderingModeSpatialAudio:
		return "AVAudioSessionRenderingModeSpatialAudio"
	case AVAudioSessionRenderingModeSurround:
		return "AVAudioSessionRenderingModeSurround"
	default:
		return fmt.Sprintf("AVAudioSessionRenderingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason
type AVAudioSessionRouteChangeReason uint

const (
	// AVAudioSessionRouteChangeReasonCategoryChange: A value that indicates that the category of the session object changed.
	AVAudioSessionRouteChangeReasonCategoryChange AVAudioSessionRouteChangeReason = 3
	// AVAudioSessionRouteChangeReasonNewDeviceAvailable: A value that indicates a user action, such as plugging in a headset, has made a preferred audio route available.
	AVAudioSessionRouteChangeReasonNewDeviceAvailable AVAudioSessionRouteChangeReason = 1
	// AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory: A value that indicates that the route changed because no suitable route is now available for the specified category.
	AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory AVAudioSessionRouteChangeReason = 7
	// AVAudioSessionRouteChangeReasonOldDeviceUnavailable: A value that indicates that the previous audio output path is no longer available.
	AVAudioSessionRouteChangeReasonOldDeviceUnavailable AVAudioSessionRouteChangeReason = 2
	// AVAudioSessionRouteChangeReasonOverride: A value that indicates that the output route was overridden by the app.
	AVAudioSessionRouteChangeReasonOverride AVAudioSessionRouteChangeReason = 4
	// AVAudioSessionRouteChangeReasonRouteConfigurationChange: A value that indicates that the configuration for a set of I/O ports has changed.
	AVAudioSessionRouteChangeReasonRouteConfigurationChange AVAudioSessionRouteChangeReason = 8
	// AVAudioSessionRouteChangeReasonUnknown: A value that indicates the reason for the change is unknown.
	AVAudioSessionRouteChangeReasonUnknown AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonWakeFromSleep: A value that indicates that the route changed when the device woke up from sleep.
	AVAudioSessionRouteChangeReasonWakeFromSleep AVAudioSessionRouteChangeReason = 6
)

func (e AVAudioSessionRouteChangeReason) String() string {
	switch e {
	case AVAudioSessionRouteChangeReasonCategoryChange:
		return "AVAudioSessionRouteChangeReasonCategoryChange"
	case AVAudioSessionRouteChangeReasonNewDeviceAvailable:
		return "AVAudioSessionRouteChangeReasonNewDeviceAvailable"
	case AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory:
		return "AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory"
	case AVAudioSessionRouteChangeReasonOldDeviceUnavailable:
		return "AVAudioSessionRouteChangeReasonOldDeviceUnavailable"
	case AVAudioSessionRouteChangeReasonOverride:
		return "AVAudioSessionRouteChangeReasonOverride"
	case AVAudioSessionRouteChangeReasonRouteConfigurationChange:
		return "AVAudioSessionRouteChangeReasonRouteConfigurationChange"
	case AVAudioSessionRouteChangeReasonUnknown:
		return "AVAudioSessionRouteChangeReasonUnknown"
	case AVAudioSessionRouteChangeReasonWakeFromSleep:
		return "AVAudioSessionRouteChangeReasonWakeFromSleep"
	default:
		return fmt.Sprintf("AVAudioSessionRouteChangeReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSelection
type AVAudioSessionRouteSelection int

const (
)

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum
type AVAudioSessionRouteSharingPolicy uint

const (
	// AVAudioSessionRouteSharingPolicyDefault: A policy that follows standard rules for routing audio output.
	AVAudioSessionRouteSharingPolicyDefault AVAudioSessionRouteSharingPolicy = 0
	// AVAudioSessionRouteSharingPolicyIndependent: A policy in which the route picker UI directs videos to a wireless route.
	AVAudioSessionRouteSharingPolicyIndependent AVAudioSessionRouteSharingPolicy = 2
	// AVAudioSessionRouteSharingPolicyLongFormAudio: A policy that routes output to the shared long-form audio output.
	AVAudioSessionRouteSharingPolicyLongFormAudio AVAudioSessionRouteSharingPolicy = 1
	// AVAudioSessionRouteSharingPolicyLongFormVideo: A policy that routes output to the shared long-form video output.
	AVAudioSessionRouteSharingPolicyLongFormVideo AVAudioSessionRouteSharingPolicy = 3
	// Deprecated.
	AVAudioSessionRouteSharingPolicyLongForm AVAudioSessionRouteSharingPolicy = 1
)

func (e AVAudioSessionRouteSharingPolicy) String() string {
	switch e {
	case AVAudioSessionRouteSharingPolicyDefault:
		return "AVAudioSessionRouteSharingPolicyDefault"
	case AVAudioSessionRouteSharingPolicyIndependent:
		return "AVAudioSessionRouteSharingPolicyIndependent"
	case AVAudioSessionRouteSharingPolicyLongFormAudio:
		return "AVAudioSessionRouteSharingPolicyLongFormAudio"
	case AVAudioSessionRouteSharingPolicyLongFormVideo:
		return "AVAudioSessionRouteSharingPolicyLongFormVideo"
	default:
		return fmt.Sprintf("AVAudioSessionRouteSharingPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SetActiveOptions
type AVAudioSessionSetActiveOptions uint

const (
	// AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation: An option that indicates that the system should notify other apps that you’ve deactivated your app’s audio session.
	AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation AVAudioSessionSetActiveOptions = 1
)

func (e AVAudioSessionSetActiveOptions) String() string {
	switch e {
	case AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation:
		return "AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation"
	default:
		return fmt.Sprintf("AVAudioSessionSetActiveOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SilenceSecondaryAudioHintType
type AVAudioSessionSilenceSecondaryAudioHintType uint

const (
	// AVAudioSessionSilenceSecondaryAudioHintTypeBegin: A value that indicates that another application’s primary audio has started.
	AVAudioSessionSilenceSecondaryAudioHintTypeBegin AVAudioSessionSilenceSecondaryAudioHintType = 1
	// AVAudioSessionSilenceSecondaryAudioHintTypeEnd: A value that indicates that another application’s primary audio has stopped.
	AVAudioSessionSilenceSecondaryAudioHintTypeEnd AVAudioSessionSilenceSecondaryAudioHintType = 0
)

func (e AVAudioSessionSilenceSecondaryAudioHintType) String() string {
	switch e {
	case AVAudioSessionSilenceSecondaryAudioHintTypeBegin:
		return "AVAudioSessionSilenceSecondaryAudioHintTypeBegin"
	case AVAudioSessionSilenceSecondaryAudioHintTypeEnd:
		return "AVAudioSessionSilenceSecondaryAudioHintTypeEnd"
	default:
		return fmt.Sprintf("AVAudioSessionSilenceSecondaryAudioHintType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize
type AVAudioSessionSoundStageSize int

const (
	// AVAudioSessionSoundStageSizeAutomatic: The system sets the sound stage size.
	AVAudioSessionSoundStageSizeAutomatic AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeLarge: A large sound stage.
	AVAudioSessionSoundStageSizeLarge AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeMedium: A medium sound stage.
	AVAudioSessionSoundStageSizeMedium AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeSmall: A small sound stage.
	AVAudioSessionSoundStageSizeSmall AVAudioSessionSoundStageSize = 0
)

func (e AVAudioSessionSoundStageSize) String() string {
	switch e {
	case AVAudioSessionSoundStageSizeAutomatic:
		return "AVAudioSessionSoundStageSizeAutomatic"
	default:
		return fmt.Sprintf("AVAudioSessionSoundStageSize(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum
type AVAudioSessionSpatialExperience int

const (
	AVAudioSessionSpatialExperienceBypassed AVAudioSessionSpatialExperience = 0
	AVAudioSessionSpatialExperienceFixed AVAudioSessionSpatialExperience = 0
	AVAudioSessionSpatialExperienceHeadTracked AVAudioSessionSpatialExperience = 0
)

func (e AVAudioSessionSpatialExperience) String() string {
	switch e {
	case AVAudioSessionSpatialExperienceBypassed:
		return "AVAudioSessionSpatialExperienceBypassed"
	default:
		return fmt.Sprintf("AVAudioSessionSpatialExperience(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation
type AVAudioStereoOrientation int

const (
	// AVAudioStereoOrientationLandscapeLeft: Audio capture should be horizontally oriented, with the USB-C or Lightning connector on the left.
	AVAudioStereoOrientationLandscapeLeft AVAudioStereoOrientation = 4
	// AVAudioStereoOrientationLandscapeRight: Audio capture should be horizontally oriented, with the USB-C or Lightning connector on the right.
	AVAudioStereoOrientationLandscapeRight AVAudioStereoOrientation = 3
	// AVAudioStereoOrientationNone: The audio session isn’t configured for stereo recording.
	AVAudioStereoOrientationNone AVAudioStereoOrientation = 0
	// AVAudioStereoOrientationPortrait: Audio capture should be vertically oriented, with the USB-C or Lightning connector on the bottom.
	AVAudioStereoOrientationPortrait AVAudioStereoOrientation = 1
	// AVAudioStereoOrientationPortraitUpsideDown: Audio capture should be vertically oriented, with the USB-C or Lightning connector on the top.
	AVAudioStereoOrientationPortraitUpsideDown AVAudioStereoOrientation = 2
)

func (e AVAudioStereoOrientation) String() string {
	switch e {
	case AVAudioStereoOrientationLandscapeLeft:
		return "AVAudioStereoOrientationLandscapeLeft"
	case AVAudioStereoOrientationLandscapeRight:
		return "AVAudioStereoOrientationLandscapeRight"
	case AVAudioStereoOrientationNone:
		return "AVAudioStereoOrientationNone"
	case AVAudioStereoOrientationPortrait:
		return "AVAudioStereoOrientationPortrait"
	case AVAudioStereoOrientationPortraitUpsideDown:
		return "AVAudioStereoOrientationPortraitUpsideDown"
	default:
		return fmt.Sprintf("AVAudioStereoOrientation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset
type AVAudioUnitDistortionPreset int

const (
	// AVAudioUnitDistortionPresetDrumsBitBrush: A preset that represents a bit brush drums distortion.
	AVAudioUnitDistortionPresetDrumsBitBrush AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetDrumsBufferBeats: A preset that represents a buffer beat drums distortion.
	AVAudioUnitDistortionPresetDrumsBufferBeats AVAudioUnitDistortionPreset = 1
	// AVAudioUnitDistortionPresetDrumsLoFi: A preset that represents a low fidelity drums distortion.
	AVAudioUnitDistortionPresetDrumsLoFi AVAudioUnitDistortionPreset = 2
	// AVAudioUnitDistortionPresetMultiBrokenSpeaker: A preset that represents a broken speaker distortion.
	AVAudioUnitDistortionPresetMultiBrokenSpeaker AVAudioUnitDistortionPreset = 3
	// AVAudioUnitDistortionPresetMultiCellphoneConcert: A preset that represents a cellphone concert distortion.
	AVAudioUnitDistortionPresetMultiCellphoneConcert AVAudioUnitDistortionPreset = 4
	// AVAudioUnitDistortionPresetMultiDecimated1: A preset that represents a variant of the decimated distortion.
	AVAudioUnitDistortionPresetMultiDecimated1 AVAudioUnitDistortionPreset = 5
	// AVAudioUnitDistortionPresetMultiDecimated2: A preset that represents a variant of the decimated distortion.
	AVAudioUnitDistortionPresetMultiDecimated2 AVAudioUnitDistortionPreset = 6
	// AVAudioUnitDistortionPresetMultiDecimated3: A preset that represents a variant of the decimated distortion.
	AVAudioUnitDistortionPresetMultiDecimated3 AVAudioUnitDistortionPreset = 7
	// AVAudioUnitDistortionPresetMultiDecimated4: A preset that represents a variant of the decimated distortion.
	AVAudioUnitDistortionPresetMultiDecimated4 AVAudioUnitDistortionPreset = 8
	// AVAudioUnitDistortionPresetMultiDistortedCubed: A preset that represents a distorted cubed distortion.
	AVAudioUnitDistortionPresetMultiDistortedCubed AVAudioUnitDistortionPreset = 10
	// AVAudioUnitDistortionPresetMultiDistortedFunk: A preset that represents a distorted funk distortion.
	AVAudioUnitDistortionPresetMultiDistortedFunk AVAudioUnitDistortionPreset = 9
	// AVAudioUnitDistortionPresetMultiDistortedSquared: A preset that represents a distorted squared distortion.
	AVAudioUnitDistortionPresetMultiDistortedSquared AVAudioUnitDistortionPreset = 11
	// AVAudioUnitDistortionPresetMultiEcho1: A preset that represents a variant of an echo distortion.
	AVAudioUnitDistortionPresetMultiEcho1 AVAudioUnitDistortionPreset = 12
	// AVAudioUnitDistortionPresetMultiEcho2: A preset that represents a variant of an echo distortion.
	AVAudioUnitDistortionPresetMultiEcho2 AVAudioUnitDistortionPreset = 13
	// AVAudioUnitDistortionPresetMultiEchoTight1: A preset that represents a variant of a tight echo distortion.
	AVAudioUnitDistortionPresetMultiEchoTight1 AVAudioUnitDistortionPreset = 14
	// AVAudioUnitDistortionPresetMultiEchoTight2: A preset that represents a variant of a tight echo distortion.
	AVAudioUnitDistortionPresetMultiEchoTight2 AVAudioUnitDistortionPreset = 15
	// AVAudioUnitDistortionPresetMultiEverythingIsBroken: A preset that represents an everything-is-broken distortion.
	AVAudioUnitDistortionPresetMultiEverythingIsBroken AVAudioUnitDistortionPreset = 16
	// AVAudioUnitDistortionPresetSpeechAlienChatter: A preset that represents an alien chatter distortion.
	AVAudioUnitDistortionPresetSpeechAlienChatter AVAudioUnitDistortionPreset = 17
	// AVAudioUnitDistortionPresetSpeechCosmicInterference: A preset that represents a cosmic interference distortion.
	AVAudioUnitDistortionPresetSpeechCosmicInterference AVAudioUnitDistortionPreset = 18
	// AVAudioUnitDistortionPresetSpeechGoldenPi: A preset that represents a golden pi distortion.
	AVAudioUnitDistortionPresetSpeechGoldenPi AVAudioUnitDistortionPreset = 19
	// AVAudioUnitDistortionPresetSpeechRadioTower: A preset that represents a radio tower distortion.
	AVAudioUnitDistortionPresetSpeechRadioTower AVAudioUnitDistortionPreset = 20
	// AVAudioUnitDistortionPresetSpeechWaves: A preset that represents a speech wave distortion.
	AVAudioUnitDistortionPresetSpeechWaves AVAudioUnitDistortionPreset = 21
)

func (e AVAudioUnitDistortionPreset) String() string {
	switch e {
	case AVAudioUnitDistortionPresetDrumsBitBrush:
		return "AVAudioUnitDistortionPresetDrumsBitBrush"
	case AVAudioUnitDistortionPresetDrumsBufferBeats:
		return "AVAudioUnitDistortionPresetDrumsBufferBeats"
	case AVAudioUnitDistortionPresetDrumsLoFi:
		return "AVAudioUnitDistortionPresetDrumsLoFi"
	case AVAudioUnitDistortionPresetMultiBrokenSpeaker:
		return "AVAudioUnitDistortionPresetMultiBrokenSpeaker"
	case AVAudioUnitDistortionPresetMultiCellphoneConcert:
		return "AVAudioUnitDistortionPresetMultiCellphoneConcert"
	case AVAudioUnitDistortionPresetMultiDecimated1:
		return "AVAudioUnitDistortionPresetMultiDecimated1"
	case AVAudioUnitDistortionPresetMultiDecimated2:
		return "AVAudioUnitDistortionPresetMultiDecimated2"
	case AVAudioUnitDistortionPresetMultiDecimated3:
		return "AVAudioUnitDistortionPresetMultiDecimated3"
	case AVAudioUnitDistortionPresetMultiDecimated4:
		return "AVAudioUnitDistortionPresetMultiDecimated4"
	case AVAudioUnitDistortionPresetMultiDistortedCubed:
		return "AVAudioUnitDistortionPresetMultiDistortedCubed"
	case AVAudioUnitDistortionPresetMultiDistortedFunk:
		return "AVAudioUnitDistortionPresetMultiDistortedFunk"
	case AVAudioUnitDistortionPresetMultiDistortedSquared:
		return "AVAudioUnitDistortionPresetMultiDistortedSquared"
	case AVAudioUnitDistortionPresetMultiEcho1:
		return "AVAudioUnitDistortionPresetMultiEcho1"
	case AVAudioUnitDistortionPresetMultiEcho2:
		return "AVAudioUnitDistortionPresetMultiEcho2"
	case AVAudioUnitDistortionPresetMultiEchoTight1:
		return "AVAudioUnitDistortionPresetMultiEchoTight1"
	case AVAudioUnitDistortionPresetMultiEchoTight2:
		return "AVAudioUnitDistortionPresetMultiEchoTight2"
	case AVAudioUnitDistortionPresetMultiEverythingIsBroken:
		return "AVAudioUnitDistortionPresetMultiEverythingIsBroken"
	case AVAudioUnitDistortionPresetSpeechAlienChatter:
		return "AVAudioUnitDistortionPresetSpeechAlienChatter"
	case AVAudioUnitDistortionPresetSpeechCosmicInterference:
		return "AVAudioUnitDistortionPresetSpeechCosmicInterference"
	case AVAudioUnitDistortionPresetSpeechGoldenPi:
		return "AVAudioUnitDistortionPresetSpeechGoldenPi"
	case AVAudioUnitDistortionPresetSpeechRadioTower:
		return "AVAudioUnitDistortionPresetSpeechRadioTower"
	case AVAudioUnitDistortionPresetSpeechWaves:
		return "AVAudioUnitDistortionPresetSpeechWaves"
	default:
		return fmt.Sprintf("AVAudioUnitDistortionPreset(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType
type AVAudioUnitEQFilterType int

const (
	// AVAudioUnitEQFilterTypeBandPass: A type that represents a bandpass filter.
	AVAudioUnitEQFilterTypeBandPass AVAudioUnitEQFilterType = 5
	// AVAudioUnitEQFilterTypeBandStop: A type that represents a band-stop filter, also known as a notch filter.
	AVAudioUnitEQFilterTypeBandStop AVAudioUnitEQFilterType = 6
	// AVAudioUnitEQFilterTypeHighPass: A type that represents a simple Butterworth second-order high-pass filter.
	AVAudioUnitEQFilterTypeHighPass AVAudioUnitEQFilterType = 2
	// AVAudioUnitEQFilterTypeHighShelf: A type that represents a high-shelf filter.
	AVAudioUnitEQFilterTypeHighShelf AVAudioUnitEQFilterType = 8
	// AVAudioUnitEQFilterTypeLowPass: A type that represents a simple Butterworth second-order low-pass filter.
	AVAudioUnitEQFilterTypeLowPass AVAudioUnitEQFilterType = 1
	// AVAudioUnitEQFilterTypeLowShelf: A type that represents a low-shelf filter.
	AVAudioUnitEQFilterTypeLowShelf AVAudioUnitEQFilterType = 7
	// AVAudioUnitEQFilterTypeParametric: A type that represents a parametric filter that derives from a Butterworth analog prototype.
	AVAudioUnitEQFilterTypeParametric AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeResonantHighPass: A type that represents a high-pass filter with resonance support using the bandwidth parameter.
	AVAudioUnitEQFilterTypeResonantHighPass AVAudioUnitEQFilterType = 4
	// AVAudioUnitEQFilterTypeResonantHighShelf: A type that represents a high-shelf filter with resonance support using the bandwidth parameter.
	AVAudioUnitEQFilterTypeResonantHighShelf AVAudioUnitEQFilterType = 10
	// AVAudioUnitEQFilterTypeResonantLowPass: A type that represents a low-pass filter with resonance support using the bandwidth parameter.
	AVAudioUnitEQFilterTypeResonantLowPass AVAudioUnitEQFilterType = 3
	// AVAudioUnitEQFilterTypeResonantLowShelf: A type that represents a low-shelf filter with resonance support using the bandwidth parameter.
	AVAudioUnitEQFilterTypeResonantLowShelf AVAudioUnitEQFilterType = 9
)

func (e AVAudioUnitEQFilterType) String() string {
	switch e {
	case AVAudioUnitEQFilterTypeBandPass:
		return "AVAudioUnitEQFilterTypeBandPass"
	case AVAudioUnitEQFilterTypeBandStop:
		return "AVAudioUnitEQFilterTypeBandStop"
	case AVAudioUnitEQFilterTypeHighPass:
		return "AVAudioUnitEQFilterTypeHighPass"
	case AVAudioUnitEQFilterTypeHighShelf:
		return "AVAudioUnitEQFilterTypeHighShelf"
	case AVAudioUnitEQFilterTypeLowPass:
		return "AVAudioUnitEQFilterTypeLowPass"
	case AVAudioUnitEQFilterTypeLowShelf:
		return "AVAudioUnitEQFilterTypeLowShelf"
	case AVAudioUnitEQFilterTypeParametric:
		return "AVAudioUnitEQFilterTypeParametric"
	case AVAudioUnitEQFilterTypeResonantHighPass:
		return "AVAudioUnitEQFilterTypeResonantHighPass"
	case AVAudioUnitEQFilterTypeResonantHighShelf:
		return "AVAudioUnitEQFilterTypeResonantHighShelf"
	case AVAudioUnitEQFilterTypeResonantLowPass:
		return "AVAudioUnitEQFilterTypeResonantLowPass"
	case AVAudioUnitEQFilterTypeResonantLowShelf:
		return "AVAudioUnitEQFilterTypeResonantLowShelf"
	default:
		return fmt.Sprintf("AVAudioUnitEQFilterType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset
type AVAudioUnitReverbPreset int

const (
	// AVAudioUnitReverbPresetCathedral: A preset that represents a reverb with the acoustic characteristics of a cathedral environment.
	AVAudioUnitReverbPresetCathedral AVAudioUnitReverbPreset = 8
	// AVAudioUnitReverbPresetLargeChamber: A preset that represents a reverb with the acoustic characteristics of a large-sized chamber environment.
	AVAudioUnitReverbPresetLargeChamber AVAudioUnitReverbPreset = 7
	// AVAudioUnitReverbPresetLargeHall: A preset that represents a reverb with the acoustic characteristics of a large-sized hall environment.
	AVAudioUnitReverbPresetLargeHall AVAudioUnitReverbPreset = 4
	// AVAudioUnitReverbPresetLargeHall2: A preset that represents a reverb with the acoustic characteristics of an alternative large-sized hall environment.
	AVAudioUnitReverbPresetLargeHall2 AVAudioUnitReverbPreset = 12
	// AVAudioUnitReverbPresetLargeRoom: A preset that represents a reverb with the acoustic characteristics of a large-sized room environment.
	AVAudioUnitReverbPresetLargeRoom AVAudioUnitReverbPreset = 2
	// AVAudioUnitReverbPresetLargeRoom2: A preset that represents a reverb with the acoustic characteristics of an alternative large-sized room environment.
	AVAudioUnitReverbPresetLargeRoom2 AVAudioUnitReverbPreset = 9
	// AVAudioUnitReverbPresetMediumChamber: A preset that represents a reverb with the acoustic characteristics of a medium-sized chamber environment.
	AVAudioUnitReverbPresetMediumChamber AVAudioUnitReverbPreset = 6
	// AVAudioUnitReverbPresetMediumHall: A preset that represents a reverb with the acoustic characteristics of a medium-sized hall environment.
	AVAudioUnitReverbPresetMediumHall AVAudioUnitReverbPreset = 3
	// AVAudioUnitReverbPresetMediumHall2: A preset that represents a reverb with the acoustic characteristics of an alternative medium-sized hall environment.
	AVAudioUnitReverbPresetMediumHall2 AVAudioUnitReverbPreset = 10
	// AVAudioUnitReverbPresetMediumHall3: A preset that represents a reverb with the acoustic characteristics of an alternative medium-sized hall environment.
	AVAudioUnitReverbPresetMediumHall3 AVAudioUnitReverbPreset = 11
	// AVAudioUnitReverbPresetMediumRoom: A preset that represents a reverb with the acoustic characteristics of a medium-sized room environment.
	AVAudioUnitReverbPresetMediumRoom AVAudioUnitReverbPreset = 1
	// AVAudioUnitReverbPresetPlate: A preset that represents a reverb with the acoustic characteristics of a plate environment.
	AVAudioUnitReverbPresetPlate AVAudioUnitReverbPreset = 5
	// AVAudioUnitReverbPresetSmallRoom: A preset that represents a reverb with the acoustic characteristics of a small-sized room environment.
	AVAudioUnitReverbPresetSmallRoom AVAudioUnitReverbPreset = 0
)

func (e AVAudioUnitReverbPreset) String() string {
	switch e {
	case AVAudioUnitReverbPresetCathedral:
		return "AVAudioUnitReverbPresetCathedral"
	case AVAudioUnitReverbPresetLargeChamber:
		return "AVAudioUnitReverbPresetLargeChamber"
	case AVAudioUnitReverbPresetLargeHall:
		return "AVAudioUnitReverbPresetLargeHall"
	case AVAudioUnitReverbPresetLargeHall2:
		return "AVAudioUnitReverbPresetLargeHall2"
	case AVAudioUnitReverbPresetLargeRoom:
		return "AVAudioUnitReverbPresetLargeRoom"
	case AVAudioUnitReverbPresetLargeRoom2:
		return "AVAudioUnitReverbPresetLargeRoom2"
	case AVAudioUnitReverbPresetMediumChamber:
		return "AVAudioUnitReverbPresetMediumChamber"
	case AVAudioUnitReverbPresetMediumHall:
		return "AVAudioUnitReverbPresetMediumHall"
	case AVAudioUnitReverbPresetMediumHall2:
		return "AVAudioUnitReverbPresetMediumHall2"
	case AVAudioUnitReverbPresetMediumHall3:
		return "AVAudioUnitReverbPresetMediumHall3"
	case AVAudioUnitReverbPresetMediumRoom:
		return "AVAudioUnitReverbPresetMediumRoom"
	case AVAudioUnitReverbPresetPlate:
		return "AVAudioUnitReverbPresetPlate"
	case AVAudioUnitReverbPresetSmallRoom:
		return "AVAudioUnitReverbPresetSmallRoom"
	default:
		return fmt.Sprintf("AVAudioUnitReverbPreset(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level
type AVAudioVoiceProcessingOtherAudioDuckingLevel int

const (
	// AVAudioVoiceProcessingOtherAudioDuckingLevelDefault: The default ducking level for typical voice chat.
	AVAudioVoiceProcessingOtherAudioDuckingLevelDefault AVAudioVoiceProcessingOtherAudioDuckingLevel = 0
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMax: Applies maximum ducking to other audio.
	AVAudioVoiceProcessingOtherAudioDuckingLevelMax AVAudioVoiceProcessingOtherAudioDuckingLevel = 30
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMid: Applies medium ducking to other audio.
	AVAudioVoiceProcessingOtherAudioDuckingLevelMid AVAudioVoiceProcessingOtherAudioDuckingLevel = 20
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMin: Applies minimum ducking to other audio.
	AVAudioVoiceProcessingOtherAudioDuckingLevelMin AVAudioVoiceProcessingOtherAudioDuckingLevel = 10
)

func (e AVAudioVoiceProcessingOtherAudioDuckingLevel) String() string {
	switch e {
	case AVAudioVoiceProcessingOtherAudioDuckingLevelDefault:
		return "AVAudioVoiceProcessingOtherAudioDuckingLevelDefault"
	case AVAudioVoiceProcessingOtherAudioDuckingLevelMax:
		return "AVAudioVoiceProcessingOtherAudioDuckingLevelMax"
	case AVAudioVoiceProcessingOtherAudioDuckingLevelMid:
		return "AVAudioVoiceProcessingOtherAudioDuckingLevelMid"
	case AVAudioVoiceProcessingOtherAudioDuckingLevelMin:
		return "AVAudioVoiceProcessingOtherAudioDuckingLevelMin"
	default:
		return fmt.Sprintf("AVAudioVoiceProcessingOtherAudioDuckingLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingSpeechActivityEvent
type AVAudioVoiceProcessingSpeechActivityEvent int

const (
	// AVAudioVoiceProcessingSpeechActivityEnded: Indicates the end of speech activity.
	AVAudioVoiceProcessingSpeechActivityEnded AVAudioVoiceProcessingSpeechActivityEvent = 1
	// AVAudioVoiceProcessingSpeechActivityStarted: Indicates the start of speech activity.
	AVAudioVoiceProcessingSpeechActivityStarted AVAudioVoiceProcessingSpeechActivityEvent = 0
)

func (e AVAudioVoiceProcessingSpeechActivityEvent) String() string {
	switch e {
	case AVAudioVoiceProcessingSpeechActivityEnded:
		return "AVAudioVoiceProcessingSpeechActivityEnded"
	case AVAudioVoiceProcessingSpeechActivityStarted:
		return "AVAudioVoiceProcessingSpeechActivityStarted"
	default:
		return fmt.Sprintf("AVAudioVoiceProcessingSpeechActivityEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum
type AVMIDIControlChangeMessageType int

const (
	// AVMIDIControlChangeMessageTypeAllNotesOff: An event type for muting all sounding notes while maintaining the release time.
	AVMIDIControlChangeMessageTypeAllNotesOff AVMIDIControlChangeMessageType = 123
	// AVMIDIControlChangeMessageTypeAllSoundOff: An event type for muting all sounding notes.
	AVMIDIControlChangeMessageTypeAllSoundOff AVMIDIControlChangeMessageType = 120
	// AVMIDIControlChangeMessageTypeAttackTime: An event type for controlling the attack time.
	AVMIDIControlChangeMessageTypeAttackTime AVMIDIControlChangeMessageType = 73
	// AVMIDIControlChangeMessageTypeBalance: An event type for controlling the left and right channel balance.
	AVMIDIControlChangeMessageTypeBalance AVMIDIControlChangeMessageType = 8
	// AVMIDIControlChangeMessageTypeBankSelect: An event type for switching bank selection.
	AVMIDIControlChangeMessageTypeBankSelect AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeBreath: An event type for a breath controller.
	AVMIDIControlChangeMessageTypeBreath AVMIDIControlChangeMessageType = 2
	// AVMIDIControlChangeMessageTypeBrightness: An event type for controlling the brightness.
	AVMIDIControlChangeMessageTypeBrightness AVMIDIControlChangeMessageType = 74
	// AVMIDIControlChangeMessageTypeChorusLevel: An event type for controlling the chorus level.
	AVMIDIControlChangeMessageTypeChorusLevel AVMIDIControlChangeMessageType = 93
	// AVMIDIControlChangeMessageTypeDataEntry: An event type for controlling the data entry parameters.
	AVMIDIControlChangeMessageTypeDataEntry AVMIDIControlChangeMessageType = 6
	// AVMIDIControlChangeMessageTypeDecayTime: An event type for controlling the decay time.
	AVMIDIControlChangeMessageTypeDecayTime AVMIDIControlChangeMessageType = 75
	// AVMIDIControlChangeMessageTypeExpression: An event type that represents an expression controller.
	AVMIDIControlChangeMessageTypeExpression AVMIDIControlChangeMessageType = 11
	// AVMIDIControlChangeMessageTypeFilterResonance: An event type for a filter resonance.
	AVMIDIControlChangeMessageTypeFilterResonance AVMIDIControlChangeMessageType = 71
	// AVMIDIControlChangeMessageTypeFoot: An event type for sending continuous stream of values when using a foot controller.
	AVMIDIControlChangeMessageTypeFoot AVMIDIControlChangeMessageType = 4
	// AVMIDIControlChangeMessageTypeHold2Pedal: An event type for holding notes.
	AVMIDIControlChangeMessageTypeHold2Pedal AVMIDIControlChangeMessageType = 69
	// AVMIDIControlChangeMessageTypeLegatoPedal: An event type for switching the legato pedal on or off.
	AVMIDIControlChangeMessageTypeLegatoPedal AVMIDIControlChangeMessageType = 68
	// AVMIDIControlChangeMessageTypeModWheel: An event type for modulating a vibrato effect.
	AVMIDIControlChangeMessageTypeModWheel AVMIDIControlChangeMessageType = 1
	// AVMIDIControlChangeMessageTypeMonoModeOff: An event type for setting the device mode to polyphonic.
	AVMIDIControlChangeMessageTypeMonoModeOff AVMIDIControlChangeMessageType = 127
	// AVMIDIControlChangeMessageTypeMonoModeOn: An event type for setting the device mode to monophonic.
	AVMIDIControlChangeMessageTypeMonoModeOn AVMIDIControlChangeMessageType = 126
	// AVMIDIControlChangeMessageTypeOmniModeOff: An event type for setting omni off mode.
	AVMIDIControlChangeMessageTypeOmniModeOff AVMIDIControlChangeMessageType = 124
	// AVMIDIControlChangeMessageTypeOmniModeOn: An event type for setting omni on mode.
	AVMIDIControlChangeMessageTypeOmniModeOn AVMIDIControlChangeMessageType = 125
	// AVMIDIControlChangeMessageTypePan: An event type for controlling the left and right channel pan.
	AVMIDIControlChangeMessageTypePan AVMIDIControlChangeMessageType = 10
	// AVMIDIControlChangeMessageTypePortamento: An event type for switching portamento on or off.
	AVMIDIControlChangeMessageTypePortamento AVMIDIControlChangeMessageType = 65
	// AVMIDIControlChangeMessageTypePortamentoTime: An event type for controlling the portamento rate.
	AVMIDIControlChangeMessageTypePortamentoTime AVMIDIControlChangeMessageType = 5
	// AVMIDIControlChangeMessageTypeRPN_LSB: An event type that represents the registered parameter number LSB.
	AVMIDIControlChangeMessageTypeRPN_LSB AVMIDIControlChangeMessageType = 100
	// AVMIDIControlChangeMessageTypeRPN_MSB: An event type that represents the registered parameter number MSB.
	AVMIDIControlChangeMessageTypeRPN_MSB AVMIDIControlChangeMessageType = 101
	// AVMIDIControlChangeMessageTypeReleaseTime: An event type for controlling the release time.
	AVMIDIControlChangeMessageTypeReleaseTime AVMIDIControlChangeMessageType = 72
	// AVMIDIControlChangeMessageTypeResetAllControllers: An event type for resetting all controllers to their default state.
	AVMIDIControlChangeMessageTypeResetAllControllers AVMIDIControlChangeMessageType = 121
	// AVMIDIControlChangeMessageTypeReverbLevel: An event type for controlling the reverb level.
	AVMIDIControlChangeMessageTypeReverbLevel AVMIDIControlChangeMessageType = 91
	// AVMIDIControlChangeMessageTypeSoft: An event type for lowering the volume of the notes.
	AVMIDIControlChangeMessageTypeSoft AVMIDIControlChangeMessageType = 67
	// AVMIDIControlChangeMessageTypeSostenuto: An event type for switching sostenuto on or off.
	AVMIDIControlChangeMessageTypeSostenuto AVMIDIControlChangeMessageType = 66
	// AVMIDIControlChangeMessageTypeSustain: An event type for switching a damper pedal on or off.
	AVMIDIControlChangeMessageTypeSustain AVMIDIControlChangeMessageType = 64
	// AVMIDIControlChangeMessageTypeVibratoDelay: An event type for controlling the vibrato delay.
	AVMIDIControlChangeMessageTypeVibratoDelay AVMIDIControlChangeMessageType = 78
	// AVMIDIControlChangeMessageTypeVibratoDepth: An event type for controlling the vibrato depth.
	AVMIDIControlChangeMessageTypeVibratoDepth AVMIDIControlChangeMessageType = 77
	// AVMIDIControlChangeMessageTypeVibratoRate: An event type for controlling the vibrato rate.
	AVMIDIControlChangeMessageTypeVibratoRate AVMIDIControlChangeMessageType = 76
	// AVMIDIControlChangeMessageTypeVolume: An event type for controlling the channel volume.
	AVMIDIControlChangeMessageTypeVolume AVMIDIControlChangeMessageType = 7
)

func (e AVMIDIControlChangeMessageType) String() string {
	switch e {
	case AVMIDIControlChangeMessageTypeAllNotesOff:
		return "AVMIDIControlChangeMessageTypeAllNotesOff"
	case AVMIDIControlChangeMessageTypeAllSoundOff:
		return "AVMIDIControlChangeMessageTypeAllSoundOff"
	case AVMIDIControlChangeMessageTypeAttackTime:
		return "AVMIDIControlChangeMessageTypeAttackTime"
	case AVMIDIControlChangeMessageTypeBalance:
		return "AVMIDIControlChangeMessageTypeBalance"
	case AVMIDIControlChangeMessageTypeBankSelect:
		return "AVMIDIControlChangeMessageTypeBankSelect"
	case AVMIDIControlChangeMessageTypeBreath:
		return "AVMIDIControlChangeMessageTypeBreath"
	case AVMIDIControlChangeMessageTypeBrightness:
		return "AVMIDIControlChangeMessageTypeBrightness"
	case AVMIDIControlChangeMessageTypeChorusLevel:
		return "AVMIDIControlChangeMessageTypeChorusLevel"
	case AVMIDIControlChangeMessageTypeDataEntry:
		return "AVMIDIControlChangeMessageTypeDataEntry"
	case AVMIDIControlChangeMessageTypeDecayTime:
		return "AVMIDIControlChangeMessageTypeDecayTime"
	case AVMIDIControlChangeMessageTypeExpression:
		return "AVMIDIControlChangeMessageTypeExpression"
	case AVMIDIControlChangeMessageTypeFilterResonance:
		return "AVMIDIControlChangeMessageTypeFilterResonance"
	case AVMIDIControlChangeMessageTypeFoot:
		return "AVMIDIControlChangeMessageTypeFoot"
	case AVMIDIControlChangeMessageTypeHold2Pedal:
		return "AVMIDIControlChangeMessageTypeHold2Pedal"
	case AVMIDIControlChangeMessageTypeLegatoPedal:
		return "AVMIDIControlChangeMessageTypeLegatoPedal"
	case AVMIDIControlChangeMessageTypeModWheel:
		return "AVMIDIControlChangeMessageTypeModWheel"
	case AVMIDIControlChangeMessageTypeMonoModeOff:
		return "AVMIDIControlChangeMessageTypeMonoModeOff"
	case AVMIDIControlChangeMessageTypeMonoModeOn:
		return "AVMIDIControlChangeMessageTypeMonoModeOn"
	case AVMIDIControlChangeMessageTypeOmniModeOff:
		return "AVMIDIControlChangeMessageTypeOmniModeOff"
	case AVMIDIControlChangeMessageTypeOmniModeOn:
		return "AVMIDIControlChangeMessageTypeOmniModeOn"
	case AVMIDIControlChangeMessageTypePan:
		return "AVMIDIControlChangeMessageTypePan"
	case AVMIDIControlChangeMessageTypePortamento:
		return "AVMIDIControlChangeMessageTypePortamento"
	case AVMIDIControlChangeMessageTypePortamentoTime:
		return "AVMIDIControlChangeMessageTypePortamentoTime"
	case AVMIDIControlChangeMessageTypeRPN_LSB:
		return "AVMIDIControlChangeMessageTypeRPN_LSB"
	case AVMIDIControlChangeMessageTypeRPN_MSB:
		return "AVMIDIControlChangeMessageTypeRPN_MSB"
	case AVMIDIControlChangeMessageTypeReleaseTime:
		return "AVMIDIControlChangeMessageTypeReleaseTime"
	case AVMIDIControlChangeMessageTypeResetAllControllers:
		return "AVMIDIControlChangeMessageTypeResetAllControllers"
	case AVMIDIControlChangeMessageTypeReverbLevel:
		return "AVMIDIControlChangeMessageTypeReverbLevel"
	case AVMIDIControlChangeMessageTypeSoft:
		return "AVMIDIControlChangeMessageTypeSoft"
	case AVMIDIControlChangeMessageTypeSostenuto:
		return "AVMIDIControlChangeMessageTypeSostenuto"
	case AVMIDIControlChangeMessageTypeSustain:
		return "AVMIDIControlChangeMessageTypeSustain"
	case AVMIDIControlChangeMessageTypeVibratoDelay:
		return "AVMIDIControlChangeMessageTypeVibratoDelay"
	case AVMIDIControlChangeMessageTypeVibratoDepth:
		return "AVMIDIControlChangeMessageTypeVibratoDepth"
	case AVMIDIControlChangeMessageTypeVibratoRate:
		return "AVMIDIControlChangeMessageTypeVibratoRate"
	case AVMIDIControlChangeMessageTypeVolume:
		return "AVMIDIControlChangeMessageTypeVolume"
	default:
		return fmt.Sprintf("AVMIDIControlChangeMessageType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType
type AVMIDIMetaEventType int

const (
	// AVMIDIMetaEventTypeCopyright: An event type that represents a copyright.
	AVMIDIMetaEventTypeCopyright AVMIDIMetaEventType = 2
	// AVMIDIMetaEventTypeCuePoint: An event type that represents a cue point.
	AVMIDIMetaEventTypeCuePoint AVMIDIMetaEventType = 7
	// AVMIDIMetaEventTypeEndOfTrack: An event type that represents the end of the track.
	AVMIDIMetaEventTypeEndOfTrack AVMIDIMetaEventType = 47
	// AVMIDIMetaEventTypeInstrument: An event type that represents an instrument.
	AVMIDIMetaEventTypeInstrument AVMIDIMetaEventType = 4
	// AVMIDIMetaEventTypeKeySignature: An event type that represents a key signature.
	AVMIDIMetaEventTypeKeySignature AVMIDIMetaEventType = 89
	// AVMIDIMetaEventTypeLyric: An event type that represents a lyric.
	AVMIDIMetaEventTypeLyric AVMIDIMetaEventType = 5
	// AVMIDIMetaEventTypeMarker: An event type that represents a marker.
	AVMIDIMetaEventTypeMarker AVMIDIMetaEventType = 6
	// AVMIDIMetaEventTypeMidiChannel: An event type that represents a MIDI channel.
	AVMIDIMetaEventTypeMidiChannel AVMIDIMetaEventType = 32
	// AVMIDIMetaEventTypeMidiPort: An event type that represents a MIDI port.
	AVMIDIMetaEventTypeMidiPort AVMIDIMetaEventType = 33
	// AVMIDIMetaEventTypeProprietaryEvent: An event type that represents a proprietary event.
	AVMIDIMetaEventTypeProprietaryEvent AVMIDIMetaEventType = 127
	// AVMIDIMetaEventTypeSequenceNumber: An event type that represents a sequence number.
	AVMIDIMetaEventTypeSequenceNumber AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeSmpteOffset: An event type that represents a SMPTE time offset.
	AVMIDIMetaEventTypeSmpteOffset AVMIDIMetaEventType = 84
	// AVMIDIMetaEventTypeTempo: An event type that represents a tempo.
	AVMIDIMetaEventTypeTempo AVMIDIMetaEventType = 81
	// AVMIDIMetaEventTypeText: An event type that represents text.
	AVMIDIMetaEventTypeText AVMIDIMetaEventType = 1
	// AVMIDIMetaEventTypeTimeSignature: An event type that represents a time signature.
	AVMIDIMetaEventTypeTimeSignature AVMIDIMetaEventType = 88
	// AVMIDIMetaEventTypeTrackName: An event type that represents a track name.
	AVMIDIMetaEventTypeTrackName AVMIDIMetaEventType = 3
)

func (e AVMIDIMetaEventType) String() string {
	switch e {
	case AVMIDIMetaEventTypeCopyright:
		return "AVMIDIMetaEventTypeCopyright"
	case AVMIDIMetaEventTypeCuePoint:
		return "AVMIDIMetaEventTypeCuePoint"
	case AVMIDIMetaEventTypeEndOfTrack:
		return "AVMIDIMetaEventTypeEndOfTrack"
	case AVMIDIMetaEventTypeInstrument:
		return "AVMIDIMetaEventTypeInstrument"
	case AVMIDIMetaEventTypeKeySignature:
		return "AVMIDIMetaEventTypeKeySignature"
	case AVMIDIMetaEventTypeLyric:
		return "AVMIDIMetaEventTypeLyric"
	case AVMIDIMetaEventTypeMarker:
		return "AVMIDIMetaEventTypeMarker"
	case AVMIDIMetaEventTypeMidiChannel:
		return "AVMIDIMetaEventTypeMidiChannel"
	case AVMIDIMetaEventTypeMidiPort:
		return "AVMIDIMetaEventTypeMidiPort"
	case AVMIDIMetaEventTypeProprietaryEvent:
		return "AVMIDIMetaEventTypeProprietaryEvent"
	case AVMIDIMetaEventTypeSequenceNumber:
		return "AVMIDIMetaEventTypeSequenceNumber"
	case AVMIDIMetaEventTypeSmpteOffset:
		return "AVMIDIMetaEventTypeSmpteOffset"
	case AVMIDIMetaEventTypeTempo:
		return "AVMIDIMetaEventTypeTempo"
	case AVMIDIMetaEventTypeText:
		return "AVMIDIMetaEventTypeText"
	case AVMIDIMetaEventTypeTimeSignature:
		return "AVMIDIMetaEventTypeTimeSignature"
	case AVMIDIMetaEventTypeTrackName:
		return "AVMIDIMetaEventTypeTrackName"
	default:
		return fmt.Sprintf("AVMIDIMetaEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMusicSequenceLoadOptions
type AVMusicSequenceLoadOptions uint

const (
	// AVMusicSequenceLoadSMF_ChannelsToTracks: An option that represents data on different MIDI channels mapped to multiple tracks.
	AVMusicSequenceLoadSMF_ChannelsToTracks AVMusicSequenceLoadOptions = 1
	// AVMusicSequenceLoadSMF_PreserveTracks: An option that preserves the tracks as they are.
	AVMusicSequenceLoadSMF_PreserveTracks AVMusicSequenceLoadOptions = 0
)

func (e AVMusicSequenceLoadOptions) String() string {
	switch e {
	case AVMusicSequenceLoadSMF_ChannelsToTracks:
		return "AVMusicSequenceLoadSMF_ChannelsToTracks"
	case AVMusicSequenceLoadSMF_PreserveTracks:
		return "AVMusicSequenceLoadSMF_PreserveTracks"
	default:
		return fmt.Sprintf("AVMusicSequenceLoadOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackLoopCount
type AVMusicTrackLoopCount int

const (
	// AVMusicTrackLoopCountForever: A track that loops forever.
	AVMusicTrackLoopCountForever AVMusicTrackLoopCount = -1
)

func (e AVMusicTrackLoopCount) String() string {
	switch e {
	case AVMusicTrackLoopCountForever:
		return "AVMusicTrackLoopCountForever"
	default:
		return fmt.Sprintf("AVMusicTrackLoopCount(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechBoundary
type AVSpeechBoundary int

const (
	// AVSpeechBoundaryImmediate: Indicates to pause or stop speech immediately.
	AVSpeechBoundaryImmediate AVSpeechBoundary = 0
	// AVSpeechBoundaryWord: Indicates to pause or stop speech after the synthesizer finishes speaking the current word.
	AVSpeechBoundaryWord AVSpeechBoundary = 1
)

func (e AVSpeechBoundary) String() string {
	switch e {
	case AVSpeechBoundaryImmediate:
		return "AVSpeechBoundaryImmediate"
	case AVSpeechBoundaryWord:
		return "AVSpeechBoundaryWord"
	default:
		return fmt.Sprintf("AVSpeechBoundary(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum
type AVSpeechSynthesisMarkerMark int

const (
	// AVSpeechSynthesisMarkerMarkBookmark: A Speech Synthesis Markup Language (SSML) mark tag.
	AVSpeechSynthesisMarkerMarkBookmark AVSpeechSynthesisMarkerMark = 4
	// AVSpeechSynthesisMarkerMarkParagraph: A type of text that represents a paragraph.
	AVSpeechSynthesisMarkerMarkParagraph AVSpeechSynthesisMarkerMark = 3
	// AVSpeechSynthesisMarkerMarkPhoneme: A type of text that represents a phoneme.
	AVSpeechSynthesisMarkerMarkPhoneme AVSpeechSynthesisMarkerMark = 0
	// AVSpeechSynthesisMarkerMarkSentence: A type of text that represents a sentence.
	AVSpeechSynthesisMarkerMarkSentence AVSpeechSynthesisMarkerMark = 2
	// AVSpeechSynthesisMarkerMarkWord: A type of text that represents a word.
	AVSpeechSynthesisMarkerMarkWord AVSpeechSynthesisMarkerMark = 1
)

func (e AVSpeechSynthesisMarkerMark) String() string {
	switch e {
	case AVSpeechSynthesisMarkerMarkBookmark:
		return "AVSpeechSynthesisMarkerMarkBookmark"
	case AVSpeechSynthesisMarkerMarkParagraph:
		return "AVSpeechSynthesisMarkerMarkParagraph"
	case AVSpeechSynthesisMarkerMarkPhoneme:
		return "AVSpeechSynthesisMarkerMarkPhoneme"
	case AVSpeechSynthesisMarkerMarkSentence:
		return "AVSpeechSynthesisMarkerMarkSentence"
	case AVSpeechSynthesisMarkerMarkWord:
		return "AVSpeechSynthesisMarkerMarkWord"
	default:
		return fmt.Sprintf("AVSpeechSynthesisMarkerMark(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum
type AVSpeechSynthesisPersonalVoiceAuthorizationStatus int

const (
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized: The user granted your app’s request to use personal voices.
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 3
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied: The user denied your app’s request to use personal voices.
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 1
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined: The app hasn’t requested authorization to use personal voices.
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 0
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported: The device doesn’t support personal voices.
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 2
)

func (e AVSpeechSynthesisPersonalVoiceAuthorizationStatus) String() string {
	switch e {
	case AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized:
		return "AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized"
	case AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied:
		return "AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied"
	case AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined:
		return "AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined"
	case AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported:
		return "AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported"
	default:
		return fmt.Sprintf("AVSpeechSynthesisPersonalVoiceAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceGender
type AVSpeechSynthesisVoiceGender int

const (
	// AVSpeechSynthesisVoiceGenderFemale: The female voice option.
	AVSpeechSynthesisVoiceGenderFemale AVSpeechSynthesisVoiceGender = 2
	// AVSpeechSynthesisVoiceGenderMale: The male voice option.
	AVSpeechSynthesisVoiceGenderMale AVSpeechSynthesisVoiceGender = 1
	// AVSpeechSynthesisVoiceGenderUnspecified: The nonspecific gender option.
	AVSpeechSynthesisVoiceGenderUnspecified AVSpeechSynthesisVoiceGender = 0
)

func (e AVSpeechSynthesisVoiceGender) String() string {
	switch e {
	case AVSpeechSynthesisVoiceGenderFemale:
		return "AVSpeechSynthesisVoiceGenderFemale"
	case AVSpeechSynthesisVoiceGenderMale:
		return "AVSpeechSynthesisVoiceGenderMale"
	case AVSpeechSynthesisVoiceGenderUnspecified:
		return "AVSpeechSynthesisVoiceGenderUnspecified"
	default:
		return fmt.Sprintf("AVSpeechSynthesisVoiceGender(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceQuality
type AVSpeechSynthesisVoiceQuality int

const (
	// AVSpeechSynthesisVoiceQualityDefault: A basic quality voice that’s  available on the device by default.
	AVSpeechSynthesisVoiceQualityDefault AVSpeechSynthesisVoiceQuality = 1
	// AVSpeechSynthesisVoiceQualityEnhanced: An enhanced quality voice that you must download to use.
	AVSpeechSynthesisVoiceQualityEnhanced AVSpeechSynthesisVoiceQuality = 2
	// AVSpeechSynthesisVoiceQualityPremium: A premium quality voice that you must download to use.
	AVSpeechSynthesisVoiceQualityPremium AVSpeechSynthesisVoiceQuality = 3
)

func (e AVSpeechSynthesisVoiceQuality) String() string {
	switch e {
	case AVSpeechSynthesisVoiceQualityDefault:
		return "AVSpeechSynthesisVoiceQualityDefault"
	case AVSpeechSynthesisVoiceQualityEnhanced:
		return "AVSpeechSynthesisVoiceQualityEnhanced"
	case AVSpeechSynthesisVoiceQualityPremium:
		return "AVSpeechSynthesisVoiceQualityPremium"
	default:
		return fmt.Sprintf("AVSpeechSynthesisVoiceQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits
type AVSpeechSynthesisVoiceTraits uint

const (
	// AVSpeechSynthesisVoiceTraitIsNoveltyVoice: The trait that indicates a voice is a novelty voice.
	AVSpeechSynthesisVoiceTraitIsNoveltyVoice AVSpeechSynthesisVoiceTraits = 1
	// AVSpeechSynthesisVoiceTraitIsPersonalVoice: The trait that indicates a voice is a personal voice.
	AVSpeechSynthesisVoiceTraitIsPersonalVoice AVSpeechSynthesisVoiceTraits = 2
	// AVSpeechSynthesisVoiceTraitNone: The trait that indicates a voice is a regular voice.
	AVSpeechSynthesisVoiceTraitNone AVSpeechSynthesisVoiceTraits = 0
)

func (e AVSpeechSynthesisVoiceTraits) String() string {
	switch e {
	case AVSpeechSynthesisVoiceTraitIsNoveltyVoice:
		return "AVSpeechSynthesisVoiceTraitIsNoveltyVoice"
	case AVSpeechSynthesisVoiceTraitIsPersonalVoice:
		return "AVSpeechSynthesisVoiceTraitIsPersonalVoice"
	case AVSpeechSynthesisVoiceTraitNone:
		return "AVSpeechSynthesisVoiceTraitNone"
	default:
		return fmt.Sprintf("AVSpeechSynthesisVoiceTraits(%d)", e)
	}
}

type AvaudiosessioninterruptionflagsShouldresume uint

const (
	// Deprecated.
	AVAudioSessionInterruptionFlags_ShouldResume AvaudiosessioninterruptionflagsShouldresume = 1
)

func (e AvaudiosessioninterruptionflagsShouldresume) String() string {
	switch e {
	case AVAudioSessionInterruptionFlags_ShouldResume:
		return "AVAudioSessionInterruptionFlags_ShouldResume"
	default:
		return fmt.Sprintf("AvaudiosessioninterruptionflagsShouldresume(%d)", e)
	}
}

type AvaudiosessionsetactiveflagsNotifyothersondeactivation uint

const (
	// Deprecated.
	AVAudioSessionSetActiveFlags_NotifyOthersOnDeactivation AvaudiosessionsetactiveflagsNotifyothersondeactivation = 1
)

func (e AvaudiosessionsetactiveflagsNotifyothersondeactivation) String() string {
	switch e {
	case AVAudioSessionSetActiveFlags_NotifyOthersOnDeactivation:
		return "AVAudioSessionSetActiveFlags_NotifyOthersOnDeactivation"
	default:
		return fmt.Sprintf("AvaudiosessionsetactiveflagsNotifyothersondeactivation(%d)", e)
	}
}

