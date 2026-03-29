// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/ClockType-swift.enum
type CMIOExtensionStreamClockType int

const (
	// CMIOExtensionStreamClockTypeCustom: Indicates that the stream’s clock is specific to the device hosting the stream.
	CMIOExtensionStreamClockTypeCustom CMIOExtensionStreamClockType = 2
	// CMIOExtensionStreamClockTypeHostTime: Indicates that the stream uses the host time clock.
	CMIOExtensionStreamClockTypeHostTime CMIOExtensionStreamClockType = 0
	// CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID: Indicates that the stream uses the clock of the linked Core Audio device.
	CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID CMIOExtensionStreamClockType = 1
)

func (e CMIOExtensionStreamClockType) String() string {
	switch e {
	case CMIOExtensionStreamClockTypeCustom:
		return "CMIOExtensionStreamClockTypeCustom"
	case CMIOExtensionStreamClockTypeHostTime:
		return "CMIOExtensionStreamClockTypeHostTime"
	case CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID:
		return "CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID"
	default:
		return fmt.Sprintf("CMIOExtensionStreamClockType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/Direction-swift.enum
type CMIOExtensionStreamDirection int

const (
	// CMIOExtensionStreamDirectionSink: A stream that consumes sample buffers for playback.
	CMIOExtensionStreamDirectionSink CMIOExtensionStreamDirection = 1
	// CMIOExtensionStreamDirectionSource: A stream that provides sample buffers for capture.
	CMIOExtensionStreamDirectionSource CMIOExtensionStreamDirection = 0
)

func (e CMIOExtensionStreamDirection) String() string {
	switch e {
	case CMIOExtensionStreamDirectionSink:
		return "CMIOExtensionStreamDirectionSink"
	case CMIOExtensionStreamDirectionSource:
		return "CMIOExtensionStreamDirectionSource"
	default:
		return fmt.Sprintf("CMIOExtensionStreamDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/DiscontinuityFlags
type CMIOExtensionStreamDiscontinuityFlags uint32

const (
	// CMIOExtensionStreamDiscontinuityFlagSampleDropped: A flag that indicates a discontinuity in the stream due to a dropped frame.
	CMIOExtensionStreamDiscontinuityFlagSampleDropped CMIOExtensionStreamDiscontinuityFlags = 64
	// CMIOExtensionStreamDiscontinuityFlagTime: A flag that indicates a time discontinuity in the stream.
	CMIOExtensionStreamDiscontinuityFlagTime CMIOExtensionStreamDiscontinuityFlags = 2
	// CMIOExtensionStreamDiscontinuityFlagUnknown: A flag that indicates a stream discontinuity due to an unknown reason.
	CMIOExtensionStreamDiscontinuityFlagUnknown CMIOExtensionStreamDiscontinuityFlags = 1
)

func (e CMIOExtensionStreamDiscontinuityFlags) String() string {
	switch e {
	case CMIOExtensionStreamDiscontinuityFlagSampleDropped:
		return "CMIOExtensionStreamDiscontinuityFlagSampleDropped"
	case CMIOExtensionStreamDiscontinuityFlagTime:
		return "CMIOExtensionStreamDiscontinuityFlagTime"
	case CMIOExtensionStreamDiscontinuityFlagUnknown:
		return "CMIOExtensionStreamDiscontinuityFlagUnknown"
	default:
		return fmt.Sprintf("CMIOExtensionStreamDiscontinuityFlags(%d)", e)
	}
}

type KCMIO uint

const (
	KCMIOBacklightCompensationControlClassID KCMIO = 'b'<<24 | 'k'<<16 | 'l'<<8 | 't' // 'bklt'
	KCMIOBlackLevelControlClassID KCMIO = 'b'<<24 | 'k'<<16 | 'l'<<8 | 'v' // 'bklv'
	KCMIOBooleanControlClassID KCMIO = 't'<<24 | 'o'<<16 | 'g'<<8 | 'l' // 'togl'
	KCMIOBrightnessControlClassID KCMIO = 'b'<<24 | 'r'<<16 | 'i'<<8 | 't' // 'brit'
	KCMIOContrastControlClassID KCMIO = 'c'<<24 | 't'<<16 | 's'<<8 | 't' // 'ctst'
	KCMIOControlClassID KCMIO = 'a'<<24 | 'c'<<16 | 't'<<8 | 'l' // 'actl'
	KCMIODevicePermissionsError KCMIO = '!'<<24 | 'h'<<16 | 'o'<<8 | 'g' // '!hog'
	KCMIODeviceUnsupportedFormatError KCMIO = '!'<<24 | 'd'<<16 | 'a'<<8 | 't' // '!dat'
	KCMIODirectionControlClassID KCMIO = 'd'<<24 | 'i'<<16 | 'r'<<8 | 'e' // 'dire'
	KCMIOExposureControlClassID KCMIO = 'x'<<24 | 'p'<<16 | 's'<<8 | 'r' // 'xpsr'
	KCMIOFeatureControlClassID KCMIO = 'f'<<24 | 't'<<16 | 'c'<<8 | 't' // 'ftct'
	KCMIOFocusControlClassID KCMIO = 'f'<<24 | 'c'<<16 | 'u'<<8 | 's' // 'fcus'
	KCMIOGainControlClassID KCMIO = 'g'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'gain'
	KCMIOGammaControlClassID KCMIO = 'g'<<24 | 'm'<<16 | 'm'<<8 | 'a' // 'gmma'
	KCMIOHardwareBadDeviceError KCMIO = '!'<<24 | 'd'<<16 | 'e'<<8 | 'v' // '!dev'
	KCMIOHardwareBadObjectError KCMIO = '!'<<24 | 'o'<<16 | 'b'<<8 | 'j' // '!obj'
	KCMIOHardwareBadPropertySizeError KCMIO = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KCMIOHardwareBadStreamError KCMIO = '!'<<24 | 's'<<16 | 't'<<8 | 'r' // '!str'
	KCMIOHardwareIllegalOperationError KCMIO = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	KCMIOHardwareNoError KCMIO = 0
	KCMIOHardwareNotRunningError KCMIO = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KCMIOHardwareNotStoppedError KCMIO = 'r'<<24 | 'u'<<16 | 'n'<<8 | ' ' // 'run '
	KCMIOHardwareSuspendedBySystemError KCMIO = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'y' // 'deny'
	KCMIOHardwareUnknownPropertyError KCMIO = 'w'<<24 | 'h'<<16 | 'o'<<8 | '?' // 'who?'
	KCMIOHardwareUnspecifiedError KCMIO = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	KCMIOHardwareUnsupportedOperationError KCMIO = 'u'<<24 | 'n'<<16 | 'o'<<8 | 'p' // 'unop'
	KCMIOHueControlClassID KCMIO = 'h'<<24 | 'u'<<16 | 'e'<<8 | ' ' // 'hue '
	KCMIOIrisControlClassID KCMIO = 'i'<<24 | 'r'<<16 | 'i'<<8 | 's' // 'iris'
	KCMIOJackControlClassID KCMIO = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
	KCMIONoiseReductionControlClassID KCMIO = 's'<<24 | '2'<<16 | 'n'<<8 | 'r' // 's2nr'
	KCMIOObjectSystemObject KCMIO = 0
	KCMIOOpticalFilterClassID KCMIO = 'o'<<24 | 'p'<<16 | 'f'<<8 | 't' // 'opft'
	KCMIOPanControlClassID KCMIO = 'p'<<24 | 'a'<<16 | 'n'<<8 | ' ' // 'pan '
	KCMIOPanTiltAbsoluteControlClassID KCMIO = 'p'<<24 | 't'<<16 | 'a'<<8 | 'b' // 'ptab'
	KCMIOPanTiltRelativeControlClassID KCMIO = 'p'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ptrl'
	KCMIOPowerLineFrequencyControlClassID KCMIO = 'p'<<24 | 'w'<<16 | 'f'<<8 | 'q' // 'pwfq'
	KCMIORollAbsoluteControlClassID KCMIO = 'r'<<24 | 'o'<<16 | 'l'<<8 | 'a' // 'rola'
	KCMIOSaturationControlClassID KCMIO = 's'<<24 | 'a'<<16 | 't'<<8 | 'u' // 'satu'
	KCMIOSelectorControlClassID KCMIO = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
	KCMIOSharpnessControlClassID KCMIO = 's'<<24 | 'h'<<16 | 'r'<<8 | 'p' // 'shrp'
	KCMIOShutterControlClassID KCMIO = 's'<<24 | 'h'<<16 | 't'<<8 | 'r' // 'shtr'
	KCMIOSystemObjectClassID KCMIO = 'a'<<24 | 's'<<16 | 'y'<<8 | 's' // 'asys'
	KCMIOTemperatureControlClassID KCMIO = 't'<<24 | 'e'<<16 | 'm'<<8 | 'p' // 'temp'
	KCMIOTiltControlClassID KCMIO = 't'<<24 | 'i'<<16 | 'l'<<8 | 't' // 'tilt'
	KCMIOWhiteBalanceControlClassID KCMIO = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'l' // 'whbl'
	KCMIOWhiteBalanceUControlClassID KCMIO = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'u' // 'whbu'
	KCMIOWhiteBalanceVControlClassID KCMIO = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'v' // 'whbv'
	KCMIOWhiteLevelControlClassID KCMIO = 'w'<<24 | 'h'<<16 | 'l'<<8 | 'v' // 'whlv'
	KCMIOZoomControlClassID KCMIO = 'z'<<24 | 'o'<<16 | 'o'<<8 | 'm' // 'zoom'
	KCMIOZoomRelativeControlClassID KCMIO = 'z'<<24 | 'o'<<16 | 'm'<<8 | 'r' // 'zomr'
)

func (e KCMIO) String() string {
	switch e {
	case KCMIOBacklightCompensationControlClassID:
		return "KCMIOBacklightCompensationControlClassID"
	case KCMIOBlackLevelControlClassID:
		return "KCMIOBlackLevelControlClassID"
	case KCMIOBooleanControlClassID:
		return "KCMIOBooleanControlClassID"
	case KCMIOBrightnessControlClassID:
		return "KCMIOBrightnessControlClassID"
	case KCMIOContrastControlClassID:
		return "KCMIOContrastControlClassID"
	case KCMIOControlClassID:
		return "KCMIOControlClassID"
	case KCMIODevicePermissionsError:
		return "KCMIODevicePermissionsError"
	case KCMIODeviceUnsupportedFormatError:
		return "KCMIODeviceUnsupportedFormatError"
	case KCMIODirectionControlClassID:
		return "KCMIODirectionControlClassID"
	case KCMIOExposureControlClassID:
		return "KCMIOExposureControlClassID"
	case KCMIOFeatureControlClassID:
		return "KCMIOFeatureControlClassID"
	case KCMIOFocusControlClassID:
		return "KCMIOFocusControlClassID"
	case KCMIOGainControlClassID:
		return "KCMIOGainControlClassID"
	case KCMIOGammaControlClassID:
		return "KCMIOGammaControlClassID"
	case KCMIOHardwareBadDeviceError:
		return "KCMIOHardwareBadDeviceError"
	case KCMIOHardwareBadObjectError:
		return "KCMIOHardwareBadObjectError"
	case KCMIOHardwareBadPropertySizeError:
		return "KCMIOHardwareBadPropertySizeError"
	case KCMIOHardwareBadStreamError:
		return "KCMIOHardwareBadStreamError"
	case KCMIOHardwareIllegalOperationError:
		return "KCMIOHardwareIllegalOperationError"
	case KCMIOHardwareNoError:
		return "KCMIOHardwareNoError"
	case KCMIOHardwareNotRunningError:
		return "KCMIOHardwareNotRunningError"
	case KCMIOHardwareNotStoppedError:
		return "KCMIOHardwareNotStoppedError"
	case KCMIOHardwareSuspendedBySystemError:
		return "KCMIOHardwareSuspendedBySystemError"
	case KCMIOHardwareUnknownPropertyError:
		return "KCMIOHardwareUnknownPropertyError"
	case KCMIOHardwareUnspecifiedError:
		return "KCMIOHardwareUnspecifiedError"
	case KCMIOHardwareUnsupportedOperationError:
		return "KCMIOHardwareUnsupportedOperationError"
	case KCMIOHueControlClassID:
		return "KCMIOHueControlClassID"
	case KCMIOIrisControlClassID:
		return "KCMIOIrisControlClassID"
	case KCMIOJackControlClassID:
		return "KCMIOJackControlClassID"
	case KCMIONoiseReductionControlClassID:
		return "KCMIONoiseReductionControlClassID"
	case KCMIOOpticalFilterClassID:
		return "KCMIOOpticalFilterClassID"
	case KCMIOPanControlClassID:
		return "KCMIOPanControlClassID"
	case KCMIOPanTiltAbsoluteControlClassID:
		return "KCMIOPanTiltAbsoluteControlClassID"
	case KCMIOPanTiltRelativeControlClassID:
		return "KCMIOPanTiltRelativeControlClassID"
	case KCMIOPowerLineFrequencyControlClassID:
		return "KCMIOPowerLineFrequencyControlClassID"
	case KCMIORollAbsoluteControlClassID:
		return "KCMIORollAbsoluteControlClassID"
	case KCMIOSaturationControlClassID:
		return "KCMIOSaturationControlClassID"
	case KCMIOSelectorControlClassID:
		return "KCMIOSelectorControlClassID"
	case KCMIOSharpnessControlClassID:
		return "KCMIOSharpnessControlClassID"
	case KCMIOShutterControlClassID:
		return "KCMIOShutterControlClassID"
	case KCMIOSystemObjectClassID:
		return "KCMIOSystemObjectClassID"
	case KCMIOTemperatureControlClassID:
		return "KCMIOTemperatureControlClassID"
	case KCMIOTiltControlClassID:
		return "KCMIOTiltControlClassID"
	case KCMIOWhiteBalanceControlClassID:
		return "KCMIOWhiteBalanceControlClassID"
	case KCMIOWhiteBalanceUControlClassID:
		return "KCMIOWhiteBalanceUControlClassID"
	case KCMIOWhiteBalanceVControlClassID:
		return "KCMIOWhiteBalanceVControlClassID"
	case KCMIOWhiteLevelControlClassID:
		return "KCMIOWhiteLevelControlClassID"
	case KCMIOZoomControlClassID:
		return "KCMIOZoomControlClassID"
	case KCMIOZoomRelativeControlClassID:
		return "KCMIOZoomRelativeControlClassID"
	default:
		return fmt.Sprintf("KCMIO(%d)", e)
	}
}

type KCMIOAVCDeviceType uint

const (
	KCMIOAVCDeviceType_DVCPro100_720p KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | 'p' // 'dvhp'
	KCMIOAVCDeviceType_DVCPro100_NTSC KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '1'<<8 | 'n' // 'dv1n'
	KCMIOAVCDeviceType_DVCPro100_PAL KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '1'<<8 | 'p' // 'dv1p'
	KCMIOAVCDeviceType_DVCPro50_NTSC KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '5'<<8 | 'n' // 'dv5n'
	KCMIOAVCDeviceType_DVCPro50_PAL KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '5'<<8 | 'p' // 'dv5p'
	KCMIOAVCDeviceType_DVCProHD_1080i50 KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | '5' // 'dvh5'
	KCMIOAVCDeviceType_DVCProHD_1080i60 KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | '6' // 'dvh6'
	KCMIOAVCDeviceType_DVCPro_NTSC KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'p'<<8 | 'n' // 'dvpn'
	KCMIOAVCDeviceType_DVCPro_PAL KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'p'<<8 | 'p' // 'dvpp'
	KCMIOAVCDeviceType_DV_NTSC KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'c'<<8 | ' ' // 'dvc '
	KCMIOAVCDeviceType_DV_PAL KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'c'<<8 | 'p' // 'dvcp'
	KCMIOAVCDeviceType_MPEG2 KCMIOAVCDeviceType = 'm'<<24 | 'p'<<16 | 'g'<<8 | '2' // 'mpg2'
	KCMIOAVCDeviceType_Unknown KCMIOAVCDeviceType = 'u'<<24 | 'n'<<16 | 'k'<<8 | 'n' // 'unkn'
)

func (e KCMIOAVCDeviceType) String() string {
	switch e {
	case KCMIOAVCDeviceType_DVCPro100_720p:
		return "KCMIOAVCDeviceType_DVCPro100_720p"
	case KCMIOAVCDeviceType_DVCPro100_NTSC:
		return "KCMIOAVCDeviceType_DVCPro100_NTSC"
	case KCMIOAVCDeviceType_DVCPro100_PAL:
		return "KCMIOAVCDeviceType_DVCPro100_PAL"
	case KCMIOAVCDeviceType_DVCPro50_NTSC:
		return "KCMIOAVCDeviceType_DVCPro50_NTSC"
	case KCMIOAVCDeviceType_DVCPro50_PAL:
		return "KCMIOAVCDeviceType_DVCPro50_PAL"
	case KCMIOAVCDeviceType_DVCProHD_1080i50:
		return "KCMIOAVCDeviceType_DVCProHD_1080i50"
	case KCMIOAVCDeviceType_DVCProHD_1080i60:
		return "KCMIOAVCDeviceType_DVCProHD_1080i60"
	case KCMIOAVCDeviceType_DVCPro_NTSC:
		return "KCMIOAVCDeviceType_DVCPro_NTSC"
	case KCMIOAVCDeviceType_DVCPro_PAL:
		return "KCMIOAVCDeviceType_DVCPro_PAL"
	case KCMIOAVCDeviceType_DV_NTSC:
		return "KCMIOAVCDeviceType_DV_NTSC"
	case KCMIOAVCDeviceType_DV_PAL:
		return "KCMIOAVCDeviceType_DV_PAL"
	case KCMIOAVCDeviceType_MPEG2:
		return "KCMIOAVCDeviceType_MPEG2"
	case KCMIOAVCDeviceType_Unknown:
		return "KCMIOAVCDeviceType_Unknown"
	default:
		return fmt.Sprintf("KCMIOAVCDeviceType(%d)", e)
	}
}

const KCMIOBooleanControlPropertyValue uint = 'b'<<24 | 'c'<<16 | 'v'<<8 | 'l' // 'bcvl'

type KCMIOControlProperty uint

const (
	KCMIOControlPropertyElement KCMIOControlProperty = 'c'<<24 | 'e'<<16 | 'l'<<8 | 'm' // 'celm'
	KCMIOControlPropertyScope KCMIOControlProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'cscp'
	KCMIOControlPropertyVariant KCMIOControlProperty = 'c'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'cvar'
)

func (e KCMIOControlProperty) String() string {
	switch e {
	case KCMIOControlPropertyElement:
		return "KCMIOControlPropertyElement"
	case KCMIOControlPropertyScope:
		return "KCMIOControlPropertyScope"
	case KCMIOControlPropertyVariant:
		return "KCMIOControlPropertyVariant"
	default:
		return fmt.Sprintf("KCMIOControlProperty(%d)", e)
	}
}

type KCMIOData uint

const (
	KCMIODataDestinationControlClassID KCMIOData = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KCMIODataSourceControlClassID KCMIOData = 'd'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'dsrc'
)

func (e KCMIOData) String() string {
	switch e {
	case KCMIODataDestinationControlClassID:
		return "KCMIODataDestinationControlClassID"
	case KCMIODataSourceControlClassID:
		return "KCMIODataSourceControlClassID"
	default:
		return fmt.Sprintf("KCMIOData(%d)", e)
	}
}

type KCMIODeckShuttle uint

const (
	KCMIODeckShuttlePause KCMIODeckShuttle = 0
	KCMIODeckShuttlePlay1x KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayFast KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayFaster KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayFastest KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayHighSpeed KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayNextFrame KCMIODeckShuttle = 0
	KCMIODeckShuttlePlayPreviousFrame KCMIODeckShuttle = 0
	KCMIODeckShuttlePlaySlow1 KCMIODeckShuttle = 0
	KCMIODeckShuttlePlaySlow2 KCMIODeckShuttle = 0
	KCMIODeckShuttlePlaySlow3 KCMIODeckShuttle = 0
	KCMIODeckShuttlePlaySlowest KCMIODeckShuttle = 0
	KCMIODeckShuttleReverse1x KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseFast KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseFaster KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseFastest KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseHighSpeed KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseSlow1 KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseSlow2 KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseSlow3 KCMIODeckShuttle = 0
	KCMIODeckShuttleReverseSlowest KCMIODeckShuttle = 0
)

func (e KCMIODeckShuttle) String() string {
	switch e {
	case KCMIODeckShuttlePause:
		return "KCMIODeckShuttlePause"
	default:
		return fmt.Sprintf("KCMIODeckShuttle(%d)", e)
	}
}

type KCMIODeckState uint

const (
	KCMIODeckStateFastForward KCMIODeckState = 0
	KCMIODeckStateFastRewind KCMIODeckState = 0
	KCMIODeckStatePause KCMIODeckState = 0
	KCMIODeckStatePlay KCMIODeckState = 0
	KCMIODeckStatePlayReverse KCMIODeckState = 0
	KCMIODeckStatePlaySlow KCMIODeckState = 0
	KCMIODeckStateReverseSlow KCMIODeckState = 0
	KCMIODeckStateStop KCMIODeckState = 0
)

func (e KCMIODeckState) String() string {
	switch e {
	case KCMIODeckStateFastForward:
		return "KCMIODeckStateFastForward"
	default:
		return fmt.Sprintf("KCMIODeckState(%d)", e)
	}
}

type KCMIODeckStatus int

const (
	KCMIODeckStatusBusy KCMIODeckStatus = 0
	KCMIODeckStatusLocal KCMIODeckStatus = 0
	KCMIODeckStatusNoDevice KCMIODeckStatus = 0
	KCMIODeckStatusNotThreaded KCMIODeckStatus = 0
	KCMIODeckStatusOpcode KCMIODeckStatus = 0
	KCMIODeckStatusSearchingForDevice KCMIODeckStatus = 0
	KCMIODeckStatusTapeInserted KCMIODeckStatus = 0
)

func (e KCMIODeckStatus) String() string {
	switch e {
	case KCMIODeckStatusBusy:
		return "KCMIODeckStatusBusy"
	default:
		return fmt.Sprintf("KCMIODeckStatus(%d)", e)
	}
}

type KCMIODevice uint

const (
	KCMIODeviceClassID KCMIODevice = 'a'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'adev'
	KCMIODevicePropertyScopeInput KCMIODevice = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KCMIODevicePropertyScopeOutput KCMIODevice = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KCMIODevicePropertyScopePlayThrough KCMIODevice = 'p'<<24 | 't'<<16 | 'r'<<8 | 'u' // 'ptru'
	KCMIODeviceUnknown KCMIODevice = 0
)

func (e KCMIODevice) String() string {
	switch e {
	case KCMIODeviceClassID:
		return "KCMIODeviceClassID"
	case KCMIODevicePropertyScopeInput:
		return "KCMIODevicePropertyScopeInput"
	case KCMIODevicePropertyScopeOutput:
		return "KCMIODevicePropertyScopeOutput"
	case KCMIODevicePropertyScopePlayThrough:
		return "KCMIODevicePropertyScopePlayThrough"
	case KCMIODeviceUnknown:
		return "KCMIODeviceUnknown"
	default:
		return fmt.Sprintf("KCMIODevice(%d)", e)
	}
}

type KCMIODeviceAVCSignal uint

const (
	KCMIODeviceAVCSignalMode8mmNTSC KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalMode8mmPAL KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeAudio KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro100_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro100_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro25_525_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro25_625_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro50_525_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVCPro50_625_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeDVHS KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHD1125_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHD1250_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHDV1_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHDV1_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHDV2_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHDV2_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHi8NTSC KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeHi8PAL KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG12Mbps_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG12Mbps_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG25Mbps_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG25Mbps_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG6Mbps_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMPEG6Mbps_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMicroMV12Mbps_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMicroMV12Mbps_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMicroMV6Mbps_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeMicroMV6Mbps_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSD525_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSD625_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSDL525_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSDL625_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSVHS525_60 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSVHS625_50 KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSMESECAM KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSMPAL KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSNPAL KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSNTSC KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSPAL KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeVHSSECAM KCMIODeviceAVCSignal = 0
)

func (e KCMIODeviceAVCSignal) String() string {
	switch e {
	case KCMIODeviceAVCSignalMode8mmNTSC:
		return "KCMIODeviceAVCSignalMode8mmNTSC"
	default:
		return fmt.Sprintf("KCMIODeviceAVCSignal(%d)", e)
	}
}

type KCMIODeviceProperty uint

const (
	KCMIODevicePropertyAVCDeviceSignalMode KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 's'<<8 | 'm' // 'pmsm'
	KCMIODevicePropertyAVCDeviceType KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'a'<<8 | 't' // 'pmat'
	KCMIODevicePropertyCanProcessAVCCommand KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'a'<<8 | 'c' // 'pmac'
	KCMIODevicePropertyCanProcessRS422Command KCMIODeviceProperty = 'r'<<24 | '4'<<16 | '2'<<8 | '2' // 'r422'
	KCMIODevicePropertyCanSwitchFrameRatesWithoutFrameDrops KCMIODeviceProperty = 'f'<<24 | 'r'<<16 | 'n'<<8 | 'd' // 'frnd'
	KCMIODevicePropertyClientSyncDiscontinuity KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'c'<<8 | 's' // 'pmcs'
	KCMIODevicePropertyDeviceCanBeDefaultDevice KCMIODeviceProperty = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KCMIODevicePropertyDeviceControl KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | 'h' // 'pmnh'
	KCMIODevicePropertyDeviceHasChanged KCMIODeviceProperty = 'd'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'diff'
	KCMIODevicePropertyDeviceHasStreamingError KCMIODeviceProperty = 's'<<24 | 'e'<<16 | 'r'<<8 | 'r' // 'serr'
	KCMIODevicePropertyDeviceIsAlive KCMIODeviceProperty = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'n' // 'livn'
	KCMIODevicePropertyDeviceIsRunning KCMIODeviceProperty = 'g'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'goin'
	KCMIODevicePropertyDeviceIsRunningSomewhere KCMIODeviceProperty = 'g'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'gone'
	KCMIODevicePropertyDeviceUID KCMIODeviceProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'uid '
	KCMIODevicePropertyExcludeNonDALAccess KCMIODeviceProperty = 'i'<<24 | 'x'<<16 | 'n'<<8 | 'a' // 'ixna'
	KCMIODevicePropertyHogMode KCMIODeviceProperty = 'o'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'oink'
	KCMIODevicePropertyIIDCCSRData KCMIODeviceProperty = 'c'<<24 | 's'<<16 | 'r'<<8 | 'd' // 'csrd'
	KCMIODevicePropertyIIDCInitialUnitSpace KCMIODeviceProperty = 'i'<<24 | 'u'<<16 | 'n'<<8 | 's' // 'iuns'
	KCMIODevicePropertyLatency KCMIODeviceProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KCMIODevicePropertyLinkedAndSyncedCoreAudioDeviceUID KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 's'<<8 | 'd' // 'plsd'
	KCMIODevicePropertyLinkedCoreAudioDeviceUID KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'd' // 'plud'
	KCMIODevicePropertyLocation KCMIODeviceProperty = 'd'<<24 | 'l'<<16 | 'o'<<8 | 'c' // 'dloc'
	KCMIODevicePropertyModelUID KCMIODeviceProperty = 'm'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'muid'
	KCMIODevicePropertyPlugIn KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'g' // 'plug'
	KCMIODevicePropertySMPTETimeCallback KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 's'<<8 | 'c' // 'pmsc'
	KCMIODevicePropertyStreamConfiguration KCMIODeviceProperty = 's'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'slay'
	KCMIODevicePropertyStreams KCMIODeviceProperty = 's'<<24 | 't'<<16 | 'm'<<8 | '#' // 'stm#'
	KCMIODevicePropertySuspendedByUser KCMIODeviceProperty = 's'<<24 | 'b'<<16 | 'y'<<8 | 'u' // 'sbyu'
	KCMIODevicePropertyTransportType KCMIODeviceProperty = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KCMIODevicePropertyVideoDigitizerComponents KCMIODeviceProperty = 'v'<<24 | 'd'<<16 | 'i'<<8 | 'g' // 'vdig'
	// Deprecated.
	KCMIODevicePropertyDeviceMaster KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | 'h' // 'pmnh'
)

func (e KCMIODeviceProperty) String() string {
	switch e {
	case KCMIODevicePropertyAVCDeviceSignalMode:
		return "KCMIODevicePropertyAVCDeviceSignalMode"
	case KCMIODevicePropertyAVCDeviceType:
		return "KCMIODevicePropertyAVCDeviceType"
	case KCMIODevicePropertyCanProcessAVCCommand:
		return "KCMIODevicePropertyCanProcessAVCCommand"
	case KCMIODevicePropertyCanProcessRS422Command:
		return "KCMIODevicePropertyCanProcessRS422Command"
	case KCMIODevicePropertyCanSwitchFrameRatesWithoutFrameDrops:
		return "KCMIODevicePropertyCanSwitchFrameRatesWithoutFrameDrops"
	case KCMIODevicePropertyClientSyncDiscontinuity:
		return "KCMIODevicePropertyClientSyncDiscontinuity"
	case KCMIODevicePropertyDeviceCanBeDefaultDevice:
		return "KCMIODevicePropertyDeviceCanBeDefaultDevice"
	case KCMIODevicePropertyDeviceControl:
		return "KCMIODevicePropertyDeviceControl"
	case KCMIODevicePropertyDeviceHasChanged:
		return "KCMIODevicePropertyDeviceHasChanged"
	case KCMIODevicePropertyDeviceHasStreamingError:
		return "KCMIODevicePropertyDeviceHasStreamingError"
	case KCMIODevicePropertyDeviceIsAlive:
		return "KCMIODevicePropertyDeviceIsAlive"
	case KCMIODevicePropertyDeviceIsRunning:
		return "KCMIODevicePropertyDeviceIsRunning"
	case KCMIODevicePropertyDeviceIsRunningSomewhere:
		return "KCMIODevicePropertyDeviceIsRunningSomewhere"
	case KCMIODevicePropertyDeviceUID:
		return "KCMIODevicePropertyDeviceUID"
	case KCMIODevicePropertyExcludeNonDALAccess:
		return "KCMIODevicePropertyExcludeNonDALAccess"
	case KCMIODevicePropertyHogMode:
		return "KCMIODevicePropertyHogMode"
	case KCMIODevicePropertyIIDCCSRData:
		return "KCMIODevicePropertyIIDCCSRData"
	case KCMIODevicePropertyIIDCInitialUnitSpace:
		return "KCMIODevicePropertyIIDCInitialUnitSpace"
	case KCMIODevicePropertyLatency:
		return "KCMIODevicePropertyLatency"
	case KCMIODevicePropertyLinkedAndSyncedCoreAudioDeviceUID:
		return "KCMIODevicePropertyLinkedAndSyncedCoreAudioDeviceUID"
	case KCMIODevicePropertyLinkedCoreAudioDeviceUID:
		return "KCMIODevicePropertyLinkedCoreAudioDeviceUID"
	case KCMIODevicePropertyLocation:
		return "KCMIODevicePropertyLocation"
	case KCMIODevicePropertyModelUID:
		return "KCMIODevicePropertyModelUID"
	case KCMIODevicePropertyPlugIn:
		return "KCMIODevicePropertyPlugIn"
	case KCMIODevicePropertySMPTETimeCallback:
		return "KCMIODevicePropertySMPTETimeCallback"
	case KCMIODevicePropertyStreamConfiguration:
		return "KCMIODevicePropertyStreamConfiguration"
	case KCMIODevicePropertyStreams:
		return "KCMIODevicePropertyStreams"
	case KCMIODevicePropertySuspendedByUser:
		return "KCMIODevicePropertySuspendedByUser"
	case KCMIODevicePropertyTransportType:
		return "KCMIODevicePropertyTransportType"
	case KCMIODevicePropertyVideoDigitizerComponents:
		return "KCMIODevicePropertyVideoDigitizerComponents"
	default:
		return fmt.Sprintf("KCMIODeviceProperty(%d)", e)
	}
}

const (
	KCMIODevicePropertyLocationBuiltInDisplay uint = 0
	KCMIODevicePropertyLocationExternalDevice uint = 0
	KCMIODevicePropertyLocationExternalDisplay uint = 0
	KCMIODevicePropertyLocationExternalWirelessDevice uint = 0
	KCMIODevicePropertyLocationUnknown uint = 0
)

type KCMIOExposureControlProperty uint

const (
	KCMIOExposureControlPropertyConvergenceSpeed KCMIOExposureControlProperty = 'e'<<24 | 'c'<<16 | 's'<<8 | 'p' // 'ecsp'
	KCMIOExposureControlPropertyIntegrationTime KCMIOExposureControlProperty = 'e'<<24 | 'i'<<16 | 'n'<<8 | 't' // 'eint'
	KCMIOExposureControlPropertyLockThreshold KCMIOExposureControlProperty = 'e'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'elck'
	KCMIOExposureControlPropertyMaximumGain KCMIOExposureControlProperty = 'e'<<24 | 'm'<<16 | 'a'<<8 | 'x' // 'emax'
	KCMIOExposureControlPropertyRegionOfInterest KCMIOExposureControlProperty = 'e'<<24 | 'r'<<16 | 'o'<<8 | 'i' // 'eroi'
	KCMIOExposureControlPropertyStability KCMIOExposureControlProperty = 'e'<<24 | 's'<<16 | 't'<<8 | 'y' // 'esty'
	KCMIOExposureControlPropertyStable KCMIOExposureControlProperty = 'e'<<24 | 's'<<16 | 't'<<8 | 'b' // 'estb'
	KCMIOExposureControlPropertyTarget KCMIOExposureControlProperty = 'e'<<24 | 't'<<16 | 'g'<<8 | 't' // 'etgt'
	KCMIOExposureControlPropertyUnlockThreshold KCMIOExposureControlProperty = 'e'<<24 | 'u'<<16 | 'l'<<8 | 'k' // 'eulk'
)

func (e KCMIOExposureControlProperty) String() string {
	switch e {
	case KCMIOExposureControlPropertyConvergenceSpeed:
		return "KCMIOExposureControlPropertyConvergenceSpeed"
	case KCMIOExposureControlPropertyIntegrationTime:
		return "KCMIOExposureControlPropertyIntegrationTime"
	case KCMIOExposureControlPropertyLockThreshold:
		return "KCMIOExposureControlPropertyLockThreshold"
	case KCMIOExposureControlPropertyMaximumGain:
		return "KCMIOExposureControlPropertyMaximumGain"
	case KCMIOExposureControlPropertyRegionOfInterest:
		return "KCMIOExposureControlPropertyRegionOfInterest"
	case KCMIOExposureControlPropertyStability:
		return "KCMIOExposureControlPropertyStability"
	case KCMIOExposureControlPropertyStable:
		return "KCMIOExposureControlPropertyStable"
	case KCMIOExposureControlPropertyTarget:
		return "KCMIOExposureControlPropertyTarget"
	case KCMIOExposureControlPropertyUnlockThreshold:
		return "KCMIOExposureControlPropertyUnlockThreshold"
	default:
		return fmt.Sprintf("KCMIOExposureControlProperty(%d)", e)
	}
}

type KCMIOFeatureControlProperty uint

const (
	KCMIOFeatureControlPropertyAbsoluteNative KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'a' // 'fcna'
	KCMIOFeatureControlPropertyAbsoluteRange KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'r' // 'fcar'
	KCMIOFeatureControlPropertyAbsoluteUnitName KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'u'<<8 | 'n' // 'fcun'
	KCMIOFeatureControlPropertyAbsoluteValue KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'v' // 'fcav'
	KCMIOFeatureControlPropertyAutomaticManual KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'm' // 'fcam'
	KCMIOFeatureControlPropertyConvertAbsoluteToNative KCMIOFeatureControlProperty = 'f'<<24 | 'a'<<16 | '2'<<8 | 'n' // 'fa2n'
	KCMIOFeatureControlPropertyConvertNativeToAbsolute KCMIOFeatureControlProperty = 'f'<<24 | 'n'<<16 | '2'<<8 | 'a' // 'fn2a'
	KCMIOFeatureControlPropertyNativeData KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'd' // 'fcnd'
	KCMIOFeatureControlPropertyNativeDataRange KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'd'<<8 | 'r' // 'fcdr'
	KCMIOFeatureControlPropertyNativeRange KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'fcnr'
	KCMIOFeatureControlPropertyNativeValue KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'v' // 'fcnv'
	KCMIOFeatureControlPropertyOnOff KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'o'<<8 | 'o' // 'fcoo'
	KCMIOFeatureControlPropertyTune KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 't'<<8 | 'n' // 'fctn'
)

func (e KCMIOFeatureControlProperty) String() string {
	switch e {
	case KCMIOFeatureControlPropertyAbsoluteNative:
		return "KCMIOFeatureControlPropertyAbsoluteNative"
	case KCMIOFeatureControlPropertyAbsoluteRange:
		return "KCMIOFeatureControlPropertyAbsoluteRange"
	case KCMIOFeatureControlPropertyAbsoluteUnitName:
		return "KCMIOFeatureControlPropertyAbsoluteUnitName"
	case KCMIOFeatureControlPropertyAbsoluteValue:
		return "KCMIOFeatureControlPropertyAbsoluteValue"
	case KCMIOFeatureControlPropertyAutomaticManual:
		return "KCMIOFeatureControlPropertyAutomaticManual"
	case KCMIOFeatureControlPropertyConvertAbsoluteToNative:
		return "KCMIOFeatureControlPropertyConvertAbsoluteToNative"
	case KCMIOFeatureControlPropertyConvertNativeToAbsolute:
		return "KCMIOFeatureControlPropertyConvertNativeToAbsolute"
	case KCMIOFeatureControlPropertyNativeData:
		return "KCMIOFeatureControlPropertyNativeData"
	case KCMIOFeatureControlPropertyNativeDataRange:
		return "KCMIOFeatureControlPropertyNativeDataRange"
	case KCMIOFeatureControlPropertyNativeRange:
		return "KCMIOFeatureControlPropertyNativeRange"
	case KCMIOFeatureControlPropertyNativeValue:
		return "KCMIOFeatureControlPropertyNativeValue"
	case KCMIOFeatureControlPropertyOnOff:
		return "KCMIOFeatureControlPropertyOnOff"
	case KCMIOFeatureControlPropertyTune:
		return "KCMIOFeatureControlPropertyTune"
	default:
		return fmt.Sprintf("KCMIOFeatureControlProperty(%d)", e)
	}
}

type KCMIOHardwareProperty uint

const (
	KCMIOHardwarePropertyAllowScreenCaptureDevices KCMIOHardwareProperty = 'y'<<24 | 'e'<<16 | 's'<<8 | ' ' // 'yes '
	KCMIOHardwarePropertyAllowWirelessScreenCaptureDevices KCMIOHardwareProperty = 'w'<<24 | 's'<<16 | 'c'<<8 | 'd' // 'wscd'
	KCMIOHardwarePropertyDefaultInputDevice KCMIOHardwareProperty = 'd'<<24 | 'I'<<16 | 'n'<<8 | ' ' // 'dIn '
	KCMIOHardwarePropertyDefaultOutputDevice KCMIOHardwareProperty = 'd'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'dOut'
	KCMIOHardwarePropertyDeviceForUID KCMIOHardwareProperty = 'd'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'duid'
	KCMIOHardwarePropertyDevices KCMIOHardwareProperty = 'd'<<24 | 'e'<<16 | 'v'<<8 | '#' // 'dev#'
	KCMIOHardwarePropertyIsInitingOrExiting KCMIOHardwareProperty = 'i'<<24 | 'n'<<16 | 'o'<<8 | 't' // 'inot'
	KCMIOHardwarePropertyPlugInForBundleID KCMIOHardwareProperty = 'p'<<24 | 'i'<<16 | 'b'<<8 | 'i' // 'pibi'
	KCMIOHardwarePropertyProcessIsMain KCMIOHardwareProperty = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'main'
	KCMIOHardwarePropertySleepingIsAllowed KCMIOHardwareProperty = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KCMIOHardwarePropertySuspendedBySystem KCMIOHardwareProperty = 's'<<24 | 'b'<<16 | 'y'<<8 | 's' // 'sbys'
	KCMIOHardwarePropertyUnloadingIsAllowed KCMIOHardwareProperty = 'u'<<24 | 'n'<<16 | 'l'<<8 | 'd' // 'unld'
	KCMIOHardwarePropertyUserSessionIsActiveOrHeadless KCMIOHardwareProperty = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
	// Deprecated.
	KCMIOHardwarePropertyProcessIsMaster KCMIOHardwareProperty = 'm'<<24 | 'a'<<16 | 's'<<8 | 't' // 'mast'
)

func (e KCMIOHardwareProperty) String() string {
	switch e {
	case KCMIOHardwarePropertyAllowScreenCaptureDevices:
		return "KCMIOHardwarePropertyAllowScreenCaptureDevices"
	case KCMIOHardwarePropertyAllowWirelessScreenCaptureDevices:
		return "KCMIOHardwarePropertyAllowWirelessScreenCaptureDevices"
	case KCMIOHardwarePropertyDefaultInputDevice:
		return "KCMIOHardwarePropertyDefaultInputDevice"
	case KCMIOHardwarePropertyDefaultOutputDevice:
		return "KCMIOHardwarePropertyDefaultOutputDevice"
	case KCMIOHardwarePropertyDeviceForUID:
		return "KCMIOHardwarePropertyDeviceForUID"
	case KCMIOHardwarePropertyDevices:
		return "KCMIOHardwarePropertyDevices"
	case KCMIOHardwarePropertyIsInitingOrExiting:
		return "KCMIOHardwarePropertyIsInitingOrExiting"
	case KCMIOHardwarePropertyPlugInForBundleID:
		return "KCMIOHardwarePropertyPlugInForBundleID"
	case KCMIOHardwarePropertyProcessIsMain:
		return "KCMIOHardwarePropertyProcessIsMain"
	case KCMIOHardwarePropertySleepingIsAllowed:
		return "KCMIOHardwarePropertySleepingIsAllowed"
	case KCMIOHardwarePropertySuspendedBySystem:
		return "KCMIOHardwarePropertySuspendedBySystem"
	case KCMIOHardwarePropertyUnloadingIsAllowed:
		return "KCMIOHardwarePropertyUnloadingIsAllowed"
	case KCMIOHardwarePropertyUserSessionIsActiveOrHeadless:
		return "KCMIOHardwarePropertyUserSessionIsActiveOrHeadless"
	case KCMIOHardwarePropertyProcessIsMaster:
		return "KCMIOHardwarePropertyProcessIsMaster"
	default:
		return fmt.Sprintf("KCMIOHardwareProperty(%d)", e)
	}
}

type KCMIOObject uint

const (
	KCMIOObjectClassID KCMIOObject = 'a'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'aobj'
	KCMIOObjectClassIDWildcard KCMIOObject = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KCMIOObjectPropertyElementMain KCMIOObject = 0
	KCMIOObjectPropertyScopeGlobal KCMIOObject = 'g'<<24 | 'l'<<16 | 'o'<<8 | 'b' // 'glob'
	KCMIOObjectUnknown KCMIOObject = 0
	// Deprecated.
	KCMIOObjectPropertyElementMaster KCMIOObject = 0
)

func (e KCMIOObject) String() string {
	switch e {
	case KCMIOObjectClassID:
		return "KCMIOObjectClassID"
	case KCMIOObjectClassIDWildcard:
		return "KCMIOObjectClassIDWildcard"
	case KCMIOObjectPropertyElementMain:
		return "KCMIOObjectPropertyElementMain"
	case KCMIOObjectPropertyScopeGlobal:
		return "KCMIOObjectPropertyScopeGlobal"
	default:
		return fmt.Sprintf("KCMIOObject(%d)", e)
	}
}

type KCMIOObjectProperty uint

const (
	KCMIOObjectPropertyClass KCMIOObjectProperty = 'c'<<24 | 'l'<<16 | 'a'<<8 | 's' // 'clas'
	KCMIOObjectPropertyCreator KCMIOObjectProperty = 'o'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'oplg'
	KCMIOObjectPropertyElementCategoryName KCMIOObjectProperty = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KCMIOObjectPropertyElementName KCMIOObjectProperty = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KCMIOObjectPropertyElementNumberName KCMIOObjectProperty = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KCMIOObjectPropertyElementWildcard KCMIOObjectProperty = 0
	KCMIOObjectPropertyListenerAdded KCMIOObjectProperty = 'l'<<24 | 'i'<<16 | 's'<<8 | 'a' // 'lisa'
	KCMIOObjectPropertyListenerRemoved KCMIOObjectProperty = 'l'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'lisr'
	KCMIOObjectPropertyManufacturer KCMIOObjectProperty = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KCMIOObjectPropertyName KCMIOObjectProperty = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KCMIOObjectPropertyOwnedObjects KCMIOObjectProperty = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KCMIOObjectPropertyOwner KCMIOObjectProperty = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
	KCMIOObjectPropertyScopeWildcard KCMIOObjectProperty = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KCMIOObjectPropertySelectorWildcard KCMIOObjectProperty = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KCMIOObjectProperty) String() string {
	switch e {
	case KCMIOObjectPropertyClass:
		return "KCMIOObjectPropertyClass"
	case KCMIOObjectPropertyCreator:
		return "KCMIOObjectPropertyCreator"
	case KCMIOObjectPropertyElementCategoryName:
		return "KCMIOObjectPropertyElementCategoryName"
	case KCMIOObjectPropertyElementName:
		return "KCMIOObjectPropertyElementName"
	case KCMIOObjectPropertyElementNumberName:
		return "KCMIOObjectPropertyElementNumberName"
	case KCMIOObjectPropertyElementWildcard:
		return "KCMIOObjectPropertyElementWildcard"
	case KCMIOObjectPropertyListenerAdded:
		return "KCMIOObjectPropertyListenerAdded"
	case KCMIOObjectPropertyListenerRemoved:
		return "KCMIOObjectPropertyListenerRemoved"
	case KCMIOObjectPropertyManufacturer:
		return "KCMIOObjectPropertyManufacturer"
	case KCMIOObjectPropertyName:
		return "KCMIOObjectPropertyName"
	case KCMIOObjectPropertyOwnedObjects:
		return "KCMIOObjectPropertyOwnedObjects"
	case KCMIOObjectPropertyOwner:
		return "KCMIOObjectPropertyOwner"
	case KCMIOObjectPropertyScopeWildcard:
		return "KCMIOObjectPropertyScopeWildcard"
	default:
		return fmt.Sprintf("KCMIOObjectProperty(%d)", e)
	}
}

const KCMIOPlugInClassID uint = 'a'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'aplg'

type KCMIOPlugInProperty uint

const (
	KCMIOPlugInPropertyBundleID KCMIOPlugInProperty = 'p'<<24 | 'i'<<16 | 'i'<<8 | 'd' // 'piid'
	KCMIOPlugInPropertyIsExtension KCMIOPlugInProperty = 'p'<<24 | 'i'<<16 | 'i'<<8 | 'e' // 'piie'
)

func (e KCMIOPlugInProperty) String() string {
	switch e {
	case KCMIOPlugInPropertyBundleID:
		return "KCMIOPlugInPropertyBundleID"
	case KCMIOPlugInPropertyIsExtension:
		return "KCMIOPlugInPropertyIsExtension"
	default:
		return fmt.Sprintf("KCMIOPlugInProperty(%d)", e)
	}
}

type KCMIOSampleBuffer uint

const (
	KCMIOSampleBufferDiscontinuityFlag_BufferOverrun KCMIOSampleBuffer = 0
	// KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_CodecSettingsChanged KCMIOSampleBuffer = 0
	// KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_DataWasDropped KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_DataWasFlushed KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_DiscontinuityInDTS KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_DurationWasExtended KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_MalformedData KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_NoDataMarker KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_PacketError KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_RelatedToDiscontinuity KCMIOSampleBuffer = 0
	// KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_StreamDiscontinuity KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_TimecodeDiscontinuity KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_TimingReferenceJumped KCMIOSampleBuffer = 0
	// KCMIOSampleBufferDiscontinuityFlag_TrickPlay: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_TrickPlay KCMIOSampleBuffer = 0
	KCMIOSampleBufferDiscontinuityFlag_UnknownDiscontinuity KCMIOSampleBuffer = 0
	KCMIOSampleBufferNoDiscontinuities KCMIOSampleBuffer = 0
)

func (e KCMIOSampleBuffer) String() string {
	switch e {
	case KCMIOSampleBufferDiscontinuityFlag_BufferOverrun:
		return "KCMIOSampleBufferDiscontinuityFlag_BufferOverrun"
	default:
		return fmt.Sprintf("KCMIOSampleBuffer(%d)", e)
	}
}

type KCMIOSampleBufferNoDataEvent uint

const (
	KCMIOSampleBufferNoDataEvent_DeviceDidNotSync KCMIOSampleBufferNoDataEvent = 0
	KCMIOSampleBufferNoDataEvent_DeviceInWrongMode KCMIOSampleBufferNoDataEvent = 0
	KCMIOSampleBufferNoDataEvent_NoMedia KCMIOSampleBufferNoDataEvent = 0
	KCMIOSampleBufferNoDataEvent_ProcessingError KCMIOSampleBufferNoDataEvent = 0
	KCMIOSampleBufferNoDataEvent_SleepWakeCycle KCMIOSampleBufferNoDataEvent = 0
	KCMIOSampleBufferNoDataEvent_Unknown KCMIOSampleBufferNoDataEvent = 0
)

func (e KCMIOSampleBufferNoDataEvent) String() string {
	switch e {
	case KCMIOSampleBufferNoDataEvent_DeviceDidNotSync:
		return "KCMIOSampleBufferNoDataEvent_DeviceDidNotSync"
	default:
		return fmt.Sprintf("KCMIOSampleBufferNoDataEvent(%d)", e)
	}
}

type KCMIOSelectorControlProperty uint

const (
	KCMIOSelectorControlPropertyAvailableItemNames KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'a'<<8 | 'n' // 'scan'
	KCMIOSelectorControlPropertyAvailableItems KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'a'<<8 | 'i' // 'scai'
	KCMIOSelectorControlPropertyCurrentItem KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'c'<<8 | 'i' // 'scci'
	KCMIOSelectorControlPropertyItemName KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'i'<<8 | 'n' // 'scin'
)

func (e KCMIOSelectorControlProperty) String() string {
	switch e {
	case KCMIOSelectorControlPropertyAvailableItemNames:
		return "KCMIOSelectorControlPropertyAvailableItemNames"
	case KCMIOSelectorControlPropertyAvailableItems:
		return "KCMIOSelectorControlPropertyAvailableItems"
	case KCMIOSelectorControlPropertyCurrentItem:
		return "KCMIOSelectorControlPropertyCurrentItem"
	case KCMIOSelectorControlPropertyItemName:
		return "KCMIOSelectorControlPropertyItemName"
	default:
		return fmt.Sprintf("KCMIOSelectorControlProperty(%d)", e)
	}
}

type KCMIOStream uint

const (
	KCMIOStreamClassID KCMIOStream = 'a'<<24 | 's'<<16 | 't'<<8 | 'r' // 'astr'
	KCMIOStreamUnknown KCMIOStream = 0
)

func (e KCMIOStream) String() string {
	switch e {
	case KCMIOStreamClassID:
		return "KCMIOStreamClassID"
	case KCMIOStreamUnknown:
		return "KCMIOStreamUnknown"
	default:
		return fmt.Sprintf("KCMIOStream(%d)", e)
	}
}

type KCMIOStreamProperty uint

const (
	KCMIOStreamPropertyCanProcessDeckCommand KCMIOStreamProperty = 'p'<<24 | 'd'<<16 | 'c'<<8 | 'd' // 'pdcd'
	KCMIOStreamPropertyClock KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'c'<<8 | 'l' // 'pmcl'
	KCMIOStreamPropertyDeck KCMIOStreamProperty = 'd'<<24 | 'e'<<16 | 'c'<<8 | 'k' // 'deck'
	KCMIOStreamPropertyDeckCueing KCMIOStreamProperty = 'c'<<24 | 'u'<<16 | 'e'<<8 | 'c' // 'cuec'
	KCMIOStreamPropertyDeckDropness KCMIOStreamProperty = 'd'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'drop'
	KCMIOStreamPropertyDeckFrameNumber KCMIOStreamProperty = 't'<<24 | 'c'<<16 | 'o'<<8 | 'd' // 'tcod'
	KCMIOStreamPropertyDeckLocal KCMIOStreamProperty = 'l'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'locl'
	KCMIOStreamPropertyDeckThreaded KCMIOStreamProperty = 't'<<24 | 'h'<<16 | 'r'<<8 | 'd' // 'thrd'
	KCMIOStreamPropertyDeviceSyncTimeoutInMSec KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '2' // 'pmn2'
	KCMIOStreamPropertyDirection KCMIOStreamProperty = 's'<<24 | 'd'<<16 | 'i'<<8 | 'r' // 'sdir'
	KCMIOStreamPropertyEndOfData KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'e'<<8 | 'd' // 'pmed'
	KCMIOStreamPropertyFirstOutputPresentationTimeStamp KCMIOStreamProperty = 'p'<<24 | 'o'<<16 | 'p'<<8 | 't' // 'popt'
	KCMIOStreamPropertyFormatDescription KCMIOStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | ' ' // 'pft '
	KCMIOStreamPropertyFormatDescriptions KCMIOStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | 'a' // 'pfta'
	KCMIOStreamPropertyFrameRate KCMIOStreamProperty = 'n'<<24 | 'f'<<16 | 'r'<<8 | 't' // 'nfrt'
	KCMIOStreamPropertyFrameRateRanges KCMIOStreamProperty = 'f'<<24 | 'r'<<16 | 'r'<<8 | 'g' // 'frrg'
	KCMIOStreamPropertyFrameRates KCMIOStreamProperty = 'n'<<24 | 'f'<<16 | 'r'<<8 | '#' // 'nfr#'
	KCMIOStreamPropertyInitialPresentationTimeStampForLinkedAndSyncedAudio KCMIOStreamProperty = 'i'<<24 | 'p'<<16 | 'l'<<8 | 's' // 'ipls'
	KCMIOStreamPropertyLatency KCMIOStreamProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KCMIOStreamPropertyMinimumFrameRate KCMIOStreamProperty = 'm'<<24 | 'f'<<16 | 'r'<<8 | 't' // 'mfrt'
	KCMIOStreamPropertyNoDataEventCount KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '3' // 'pmn3'
	KCMIOStreamPropertyNoDataTimeoutInMSec KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '1' // 'pmn1'
	KCMIOStreamPropertyOutputBufferQueueSize KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'q' // 'pmoq'
	KCMIOStreamPropertyOutputBufferRepeatCount KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'r' // 'pmor'
	KCMIOStreamPropertyOutputBufferUnderrunCount KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'u' // 'pmou'
	KCMIOStreamPropertyOutputBuffersNeededForThrottledPlayback KCMIOStreamProperty = 'm'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'miff'
	KCMIOStreamPropertyOutputBuffersRequiredForStartup KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 's' // 'pmos'
	KCMIOStreamPropertyPreferredFormatDescription KCMIOStreamProperty = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'd' // 'prfd'
	KCMIOStreamPropertyPreferredFrameRate KCMIOStreamProperty = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'r' // 'prfr'
	KCMIOStreamPropertyScheduledOutputNotificationProc KCMIOStreamProperty = 's'<<24 | 'o'<<16 | 'n'<<8 | 'p' // 'sonp'
	KCMIOStreamPropertyStartingChannel KCMIOStreamProperty = 's'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'schn'
	KCMIOStreamPropertyStillImage KCMIOStreamProperty = 's'<<24 | 't'<<16 | 'm'<<8 | 'g' // 'stmg'
	KCMIOStreamPropertyStillImageFormatDescriptions KCMIOStreamProperty = 's'<<24 | 't'<<16 | 'f'<<8 | 't' // 'stft'
	KCMIOStreamPropertyTerminalType KCMIOStreamProperty = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
)

func (e KCMIOStreamProperty) String() string {
	switch e {
	case KCMIOStreamPropertyCanProcessDeckCommand:
		return "KCMIOStreamPropertyCanProcessDeckCommand"
	case KCMIOStreamPropertyClock:
		return "KCMIOStreamPropertyClock"
	case KCMIOStreamPropertyDeck:
		return "KCMIOStreamPropertyDeck"
	case KCMIOStreamPropertyDeckCueing:
		return "KCMIOStreamPropertyDeckCueing"
	case KCMIOStreamPropertyDeckDropness:
		return "KCMIOStreamPropertyDeckDropness"
	case KCMIOStreamPropertyDeckFrameNumber:
		return "KCMIOStreamPropertyDeckFrameNumber"
	case KCMIOStreamPropertyDeckLocal:
		return "KCMIOStreamPropertyDeckLocal"
	case KCMIOStreamPropertyDeckThreaded:
		return "KCMIOStreamPropertyDeckThreaded"
	case KCMIOStreamPropertyDeviceSyncTimeoutInMSec:
		return "KCMIOStreamPropertyDeviceSyncTimeoutInMSec"
	case KCMIOStreamPropertyDirection:
		return "KCMIOStreamPropertyDirection"
	case KCMIOStreamPropertyEndOfData:
		return "KCMIOStreamPropertyEndOfData"
	case KCMIOStreamPropertyFirstOutputPresentationTimeStamp:
		return "KCMIOStreamPropertyFirstOutputPresentationTimeStamp"
	case KCMIOStreamPropertyFormatDescription:
		return "KCMIOStreamPropertyFormatDescription"
	case KCMIOStreamPropertyFormatDescriptions:
		return "KCMIOStreamPropertyFormatDescriptions"
	case KCMIOStreamPropertyFrameRate:
		return "KCMIOStreamPropertyFrameRate"
	case KCMIOStreamPropertyFrameRateRanges:
		return "KCMIOStreamPropertyFrameRateRanges"
	case KCMIOStreamPropertyFrameRates:
		return "KCMIOStreamPropertyFrameRates"
	case KCMIOStreamPropertyInitialPresentationTimeStampForLinkedAndSyncedAudio:
		return "KCMIOStreamPropertyInitialPresentationTimeStampForLinkedAndSyncedAudio"
	case KCMIOStreamPropertyLatency:
		return "KCMIOStreamPropertyLatency"
	case KCMIOStreamPropertyMinimumFrameRate:
		return "KCMIOStreamPropertyMinimumFrameRate"
	case KCMIOStreamPropertyNoDataEventCount:
		return "KCMIOStreamPropertyNoDataEventCount"
	case KCMIOStreamPropertyNoDataTimeoutInMSec:
		return "KCMIOStreamPropertyNoDataTimeoutInMSec"
	case KCMIOStreamPropertyOutputBufferQueueSize:
		return "KCMIOStreamPropertyOutputBufferQueueSize"
	case KCMIOStreamPropertyOutputBufferRepeatCount:
		return "KCMIOStreamPropertyOutputBufferRepeatCount"
	case KCMIOStreamPropertyOutputBufferUnderrunCount:
		return "KCMIOStreamPropertyOutputBufferUnderrunCount"
	case KCMIOStreamPropertyOutputBuffersNeededForThrottledPlayback:
		return "KCMIOStreamPropertyOutputBuffersNeededForThrottledPlayback"
	case KCMIOStreamPropertyOutputBuffersRequiredForStartup:
		return "KCMIOStreamPropertyOutputBuffersRequiredForStartup"
	case KCMIOStreamPropertyPreferredFormatDescription:
		return "KCMIOStreamPropertyPreferredFormatDescription"
	case KCMIOStreamPropertyPreferredFrameRate:
		return "KCMIOStreamPropertyPreferredFrameRate"
	case KCMIOStreamPropertyScheduledOutputNotificationProc:
		return "KCMIOStreamPropertyScheduledOutputNotificationProc"
	case KCMIOStreamPropertyStartingChannel:
		return "KCMIOStreamPropertyStartingChannel"
	case KCMIOStreamPropertyStillImage:
		return "KCMIOStreamPropertyStillImage"
	case KCMIOStreamPropertyStillImageFormatDescriptions:
		return "KCMIOStreamPropertyStillImageFormatDescriptions"
	case KCMIOStreamPropertyTerminalType:
		return "KCMIOStreamPropertyTerminalType"
	default:
		return fmt.Sprintf("KCMIOStreamProperty(%d)", e)
	}
}

