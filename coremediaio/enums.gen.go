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
	// CMIOExtensionStreamDiscontinuityFlagNone: A flag that indicates there’s no discontinuity in the stream.
	CMIOExtensionStreamDiscontinuityFlagNone CMIOExtensionStreamDiscontinuityFlags = 0
	// CMIOExtensionStreamDiscontinuityFlagSampleDropped: A flag that indicates a discontinuity in the stream due to a dropped frame.
	CMIOExtensionStreamDiscontinuityFlagSampleDropped CMIOExtensionStreamDiscontinuityFlags = 64
	// CMIOExtensionStreamDiscontinuityFlagTime: A flag that indicates a time discontinuity in the stream.
	CMIOExtensionStreamDiscontinuityFlagTime CMIOExtensionStreamDiscontinuityFlags = 2
	// CMIOExtensionStreamDiscontinuityFlagUnknown: A flag that indicates a stream discontinuity due to an unknown reason.
	CMIOExtensionStreamDiscontinuityFlagUnknown CMIOExtensionStreamDiscontinuityFlags = 1
)

func (e CMIOExtensionStreamDiscontinuityFlags) String() string {
	switch e {
	case CMIOExtensionStreamDiscontinuityFlagNone:
		return "CMIOExtensionStreamDiscontinuityFlagNone"
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

type KCMIOAVCDeviceType uint32

const (
	KCMIOAVCDeviceType_DVCPro100_720p   KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | 'p' // 'dvhp'
	KCMIOAVCDeviceType_DVCPro100_NTSC   KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '1'<<8 | 'n' // 'dv1n'
	KCMIOAVCDeviceType_DVCPro100_PAL    KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '1'<<8 | 'p' // 'dv1p'
	KCMIOAVCDeviceType_DVCPro50_NTSC    KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '5'<<8 | 'n' // 'dv5n'
	KCMIOAVCDeviceType_DVCPro50_PAL     KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | '5'<<8 | 'p' // 'dv5p'
	KCMIOAVCDeviceType_DVCProHD_1080i50 KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | '5' // 'dvh5'
	KCMIOAVCDeviceType_DVCProHD_1080i60 KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'h'<<8 | '6' // 'dvh6'
	KCMIOAVCDeviceType_DVCPro_NTSC      KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'p'<<8 | 'n' // 'dvpn'
	KCMIOAVCDeviceType_DVCPro_PAL       KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'p'<<8 | 'p' // 'dvpp'
	KCMIOAVCDeviceType_DV_NTSC          KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'c'<<8 | ' ' // 'dvc '
	KCMIOAVCDeviceType_DV_PAL           KCMIOAVCDeviceType = 'd'<<24 | 'v'<<16 | 'c'<<8 | 'p' // 'dvcp'
	KCMIOAVCDeviceType_MPEG2            KCMIOAVCDeviceType = 'm'<<24 | 'p'<<16 | 'g'<<8 | '2' // 'mpg2'
	KCMIOAVCDeviceType_Unknown          KCMIOAVCDeviceType = 'u'<<24 | 'n'<<16 | 'k'<<8 | 'n' // 'unkn'
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

type KCMIOBlackLevelControlClassID uint32

const (
	KCMIOBacklightCompensationControlClassID KCMIOBlackLevelControlClassID = 'b'<<24 | 'k'<<16 | 'l'<<8 | 't' // 'bklt'
	KCMIOBlackLevelControlClassIDValue       KCMIOBlackLevelControlClassID = 'b'<<24 | 'k'<<16 | 'l'<<8 | 'v' // 'bklv'
	KCMIOBrightnessControlClassID            KCMIOBlackLevelControlClassID = 'b'<<24 | 'r'<<16 | 'i'<<8 | 't' // 'brit'
	KCMIOContrastControlClassID              KCMIOBlackLevelControlClassID = 'c'<<24 | 't'<<16 | 's'<<8 | 't' // 'ctst'
	KCMIOExposureControlClassID              KCMIOBlackLevelControlClassID = 'x'<<24 | 'p'<<16 | 's'<<8 | 'r' // 'xpsr'
	KCMIOFocusControlClassID                 KCMIOBlackLevelControlClassID = 'f'<<24 | 'c'<<16 | 'u'<<8 | 's' // 'fcus'
	KCMIOGainControlClassID                  KCMIOBlackLevelControlClassID = 'g'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'gain'
	KCMIOGammaControlClassID                 KCMIOBlackLevelControlClassID = 'g'<<24 | 'm'<<16 | 'm'<<8 | 'a' // 'gmma'
	KCMIOHueControlClassID                   KCMIOBlackLevelControlClassID = 'h'<<24 | 'u'<<16 | 'e'<<8 | ' ' // 'hue '
	KCMIOIrisControlClassID                  KCMIOBlackLevelControlClassID = 'i'<<24 | 'r'<<16 | 'i'<<8 | 's' // 'iris'
	KCMIONoiseReductionControlClassID        KCMIOBlackLevelControlClassID = 's'<<24 | '2'<<16 | 'n'<<8 | 'r' // 's2nr'
	KCMIOOpticalFilterClassID                KCMIOBlackLevelControlClassID = 'o'<<24 | 'p'<<16 | 'f'<<8 | 't' // 'opft'
	KCMIOPanControlClassID                   KCMIOBlackLevelControlClassID = 'p'<<24 | 'a'<<16 | 'n'<<8 | ' ' // 'pan '
	KCMIOPanTiltAbsoluteControlClassID       KCMIOBlackLevelControlClassID = 'p'<<24 | 't'<<16 | 'a'<<8 | 'b' // 'ptab'
	KCMIOPanTiltRelativeControlClassID       KCMIOBlackLevelControlClassID = 'p'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ptrl'
	KCMIOPowerLineFrequencyControlClassID    KCMIOBlackLevelControlClassID = 'p'<<24 | 'w'<<16 | 'f'<<8 | 'q' // 'pwfq'
	KCMIORollAbsoluteControlClassID          KCMIOBlackLevelControlClassID = 'r'<<24 | 'o'<<16 | 'l'<<8 | 'a' // 'rola'
	KCMIOSaturationControlClassID            KCMIOBlackLevelControlClassID = 's'<<24 | 'a'<<16 | 't'<<8 | 'u' // 'satu'
	KCMIOSharpnessControlClassID             KCMIOBlackLevelControlClassID = 's'<<24 | 'h'<<16 | 'r'<<8 | 'p' // 'shrp'
	KCMIOShutterControlClassID               KCMIOBlackLevelControlClassID = 's'<<24 | 'h'<<16 | 't'<<8 | 'r' // 'shtr'
	KCMIOTemperatureControlClassID           KCMIOBlackLevelControlClassID = 't'<<24 | 'e'<<16 | 'm'<<8 | 'p' // 'temp'
	KCMIOTiltControlClassID                  KCMIOBlackLevelControlClassID = 't'<<24 | 'i'<<16 | 'l'<<8 | 't' // 'tilt'
	KCMIOWhiteBalanceControlClassID          KCMIOBlackLevelControlClassID = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'l' // 'whbl'
	KCMIOWhiteBalanceUControlClassID         KCMIOBlackLevelControlClassID = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'u' // 'whbu'
	KCMIOWhiteBalanceVControlClassID         KCMIOBlackLevelControlClassID = 'w'<<24 | 'h'<<16 | 'b'<<8 | 'v' // 'whbv'
	KCMIOWhiteLevelControlClassID            KCMIOBlackLevelControlClassID = 'w'<<24 | 'h'<<16 | 'l'<<8 | 'v' // 'whlv'
	KCMIOZoomControlClassID                  KCMIOBlackLevelControlClassID = 'z'<<24 | 'o'<<16 | 'o'<<8 | 'm' // 'zoom'
	KCMIOZoomRelativeControlClassID          KCMIOBlackLevelControlClassID = 'z'<<24 | 'o'<<16 | 'm'<<8 | 'r' // 'zomr'
)

func (e KCMIOBlackLevelControlClassID) String() string {
	switch e {
	case KCMIOBacklightCompensationControlClassID:
		return "KCMIOBacklightCompensationControlClassID"
	case KCMIOBlackLevelControlClassIDValue:
		return "KCMIOBlackLevelControlClassIDValue"
	case KCMIOBrightnessControlClassID:
		return "KCMIOBrightnessControlClassID"
	case KCMIOContrastControlClassID:
		return "KCMIOContrastControlClassID"
	case KCMIOExposureControlClassID:
		return "KCMIOExposureControlClassID"
	case KCMIOFocusControlClassID:
		return "KCMIOFocusControlClassID"
	case KCMIOGainControlClassID:
		return "KCMIOGainControlClassID"
	case KCMIOGammaControlClassID:
		return "KCMIOGammaControlClassID"
	case KCMIOHueControlClassID:
		return "KCMIOHueControlClassID"
	case KCMIOIrisControlClassID:
		return "KCMIOIrisControlClassID"
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
	case KCMIOSharpnessControlClassID:
		return "KCMIOSharpnessControlClassID"
	case KCMIOShutterControlClassID:
		return "KCMIOShutterControlClassID"
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
		return fmt.Sprintf("KCMIOBlackLevelControlClassID(%d)", e)
	}
}

type KCMIOBooleanControlProperty uint32

const (
	KCMIOBooleanControlPropertyValue KCMIOBooleanControlProperty = 'b'<<24 | 'c'<<16 | 'v'<<8 | 'l' // 'bcvl'
)

func (e KCMIOBooleanControlProperty) String() string {
	switch e {
	case KCMIOBooleanControlPropertyValue:
		return "KCMIOBooleanControlPropertyValue"
	default:
		return fmt.Sprintf("KCMIOBooleanControlProperty(%d)", e)
	}
}

type KCMIOControlClassID uint32

const (
	KCMIOBooleanControlClassID  KCMIOControlClassID = 't'<<24 | 'o'<<16 | 'g'<<8 | 'l' // 'togl'
	KCMIOControlClassIDValue    KCMIOControlClassID = 'a'<<24 | 'c'<<16 | 't'<<8 | 'l' // 'actl'
	KCMIOFeatureControlClassID  KCMIOControlClassID = 'f'<<24 | 't'<<16 | 'c'<<8 | 't' // 'ftct'
	KCMIOSelectorControlClassID KCMIOControlClassID = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
)

func (e KCMIOControlClassID) String() string {
	switch e {
	case KCMIOBooleanControlClassID:
		return "KCMIOBooleanControlClassID"
	case KCMIOControlClassIDValue:
		return "KCMIOControlClassIDValue"
	case KCMIOFeatureControlClassID:
		return "KCMIOFeatureControlClassID"
	case KCMIOSelectorControlClassID:
		return "KCMIOSelectorControlClassID"
	default:
		return fmt.Sprintf("KCMIOControlClassID(%d)", e)
	}
}

type KCMIOControlProperty uint32

const (
	KCMIOControlPropertyElement KCMIOControlProperty = 'c'<<24 | 'e'<<16 | 'l'<<8 | 'm' // 'celm'
	KCMIOControlPropertyScope   KCMIOControlProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'cscp'
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

type KCMIOData uint32

const (
	KCMIODataDestinationControlClassID KCMIOData = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KCMIODataSourceControlClassID      KCMIOData = 'd'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'dsrc'
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

type KCMIODeckShuttle int32

const (
	KCMIODeckShuttlePause             KCMIODeckShuttle = 0
	KCMIODeckShuttlePlay1x            KCMIODeckShuttle = 6
	KCMIODeckShuttlePlayFast          KCMIODeckShuttle = 7
	KCMIODeckShuttlePlayFaster        KCMIODeckShuttle = 8
	KCMIODeckShuttlePlayFastest       KCMIODeckShuttle = 9
	KCMIODeckShuttlePlayHighSpeed     KCMIODeckShuttle = 10
	KCMIODeckShuttlePlayNextFrame     KCMIODeckShuttle = 1
	KCMIODeckShuttlePlayPreviousFrame KCMIODeckShuttle = -1
	KCMIODeckShuttlePlaySlow1         KCMIODeckShuttle = 3
	KCMIODeckShuttlePlaySlow2         KCMIODeckShuttle = 4
	KCMIODeckShuttlePlaySlow3         KCMIODeckShuttle = 5
	KCMIODeckShuttlePlaySlowest       KCMIODeckShuttle = 2
	KCMIODeckShuttleReverse1x         KCMIODeckShuttle = -6
	KCMIODeckShuttleReverseFast       KCMIODeckShuttle = -7
	KCMIODeckShuttleReverseFaster     KCMIODeckShuttle = -8
	KCMIODeckShuttleReverseFastest    KCMIODeckShuttle = -9
	KCMIODeckShuttleReverseHighSpeed  KCMIODeckShuttle = -10
	KCMIODeckShuttleReverseSlow1      KCMIODeckShuttle = -3
	KCMIODeckShuttleReverseSlow2      KCMIODeckShuttle = -4
	KCMIODeckShuttleReverseSlow3      KCMIODeckShuttle = -5
	KCMIODeckShuttleReverseSlowest    KCMIODeckShuttle = -2
)

func (e KCMIODeckShuttle) String() string {
	switch e {
	case KCMIODeckShuttlePause:
		return "KCMIODeckShuttlePause"
	case KCMIODeckShuttlePlay1x:
		return "KCMIODeckShuttlePlay1x"
	case KCMIODeckShuttlePlayFast:
		return "KCMIODeckShuttlePlayFast"
	case KCMIODeckShuttlePlayFaster:
		return "KCMIODeckShuttlePlayFaster"
	case KCMIODeckShuttlePlayFastest:
		return "KCMIODeckShuttlePlayFastest"
	case KCMIODeckShuttlePlayHighSpeed:
		return "KCMIODeckShuttlePlayHighSpeed"
	case KCMIODeckShuttlePlayNextFrame:
		return "KCMIODeckShuttlePlayNextFrame"
	case KCMIODeckShuttlePlayPreviousFrame:
		return "KCMIODeckShuttlePlayPreviousFrame"
	case KCMIODeckShuttlePlaySlow1:
		return "KCMIODeckShuttlePlaySlow1"
	case KCMIODeckShuttlePlaySlow2:
		return "KCMIODeckShuttlePlaySlow2"
	case KCMIODeckShuttlePlaySlow3:
		return "KCMIODeckShuttlePlaySlow3"
	case KCMIODeckShuttlePlaySlowest:
		return "KCMIODeckShuttlePlaySlowest"
	case KCMIODeckShuttleReverse1x:
		return "KCMIODeckShuttleReverse1x"
	case KCMIODeckShuttleReverseFast:
		return "KCMIODeckShuttleReverseFast"
	case KCMIODeckShuttleReverseFaster:
		return "KCMIODeckShuttleReverseFaster"
	case KCMIODeckShuttleReverseFastest:
		return "KCMIODeckShuttleReverseFastest"
	case KCMIODeckShuttleReverseHighSpeed:
		return "KCMIODeckShuttleReverseHighSpeed"
	case KCMIODeckShuttleReverseSlow1:
		return "KCMIODeckShuttleReverseSlow1"
	case KCMIODeckShuttleReverseSlow2:
		return "KCMIODeckShuttleReverseSlow2"
	case KCMIODeckShuttleReverseSlow3:
		return "KCMIODeckShuttleReverseSlow3"
	case KCMIODeckShuttleReverseSlowest:
		return "KCMIODeckShuttleReverseSlowest"
	default:
		return fmt.Sprintf("KCMIODeckShuttle(%d)", e)
	}
}

type KCMIODeckState uint32

const (
	KCMIODeckStateFastForward KCMIODeckState = 6
	KCMIODeckStateFastRewind  KCMIODeckState = 7
	KCMIODeckStatePause       KCMIODeckState = 2
	KCMIODeckStatePlay        KCMIODeckState = 1
	KCMIODeckStatePlayReverse KCMIODeckState = 5
	KCMIODeckStatePlaySlow    KCMIODeckState = 3
	KCMIODeckStateReverseSlow KCMIODeckState = 4
	KCMIODeckStateStop        KCMIODeckState = 0
)

func (e KCMIODeckState) String() string {
	switch e {
	case KCMIODeckStateFastForward:
		return "KCMIODeckStateFastForward"
	case KCMIODeckStateFastRewind:
		return "KCMIODeckStateFastRewind"
	case KCMIODeckStatePause:
		return "KCMIODeckStatePause"
	case KCMIODeckStatePlay:
		return "KCMIODeckStatePlay"
	case KCMIODeckStatePlayReverse:
		return "KCMIODeckStatePlayReverse"
	case KCMIODeckStatePlaySlow:
		return "KCMIODeckStatePlaySlow"
	case KCMIODeckStateReverseSlow:
		return "KCMIODeckStateReverseSlow"
	case KCMIODeckStateStop:
		return "KCMIODeckStateStop"
	default:
		return fmt.Sprintf("KCMIODeckState(%d)", e)
	}
}

type KCMIODeckStatus uint32

const (
	KCMIODeckStatusBusy               KCMIODeckStatus = 1
	KCMIODeckStatusLocal              KCMIODeckStatus = 2
	KCMIODeckStatusNoDevice           KCMIODeckStatus = 7
	KCMIODeckStatusNotThreaded        KCMIODeckStatus = 3
	KCMIODeckStatusOpcode             KCMIODeckStatus = 5
	KCMIODeckStatusSearchingForDevice KCMIODeckStatus = 6
	KCMIODeckStatusTapeInserted       KCMIODeckStatus = 4
)

func (e KCMIODeckStatus) String() string {
	switch e {
	case KCMIODeckStatusBusy:
		return "KCMIODeckStatusBusy"
	case KCMIODeckStatusLocal:
		return "KCMIODeckStatusLocal"
	case KCMIODeckStatusNoDevice:
		return "KCMIODeckStatusNoDevice"
	case KCMIODeckStatusNotThreaded:
		return "KCMIODeckStatusNotThreaded"
	case KCMIODeckStatusOpcode:
		return "KCMIODeckStatusOpcode"
	case KCMIODeckStatusSearchingForDevice:
		return "KCMIODeckStatusSearchingForDevice"
	case KCMIODeckStatusTapeInserted:
		return "KCMIODeckStatusTapeInserted"
	default:
		return fmt.Sprintf("KCMIODeckStatus(%d)", e)
	}
}

type KCMIODevice uint32

const (
	KCMIODeviceClassID                  KCMIODevice = 'a'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'adev'
	KCMIODevicePropertyScopeInput       KCMIODevice = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KCMIODevicePropertyScopeOutput      KCMIODevice = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KCMIODevicePropertyScopePlayThrough KCMIODevice = 'p'<<24 | 't'<<16 | 'r'<<8 | 'u' // 'ptru'
	KCMIODeviceUnknown                  KCMIODevice = 0
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

type KCMIODeviceAVCSignal uint32

const (
	KCMIODeviceAVCSignalMode8mmNTSC          KCMIODeviceAVCSignal = 0x6
	KCMIODeviceAVCSignalMode8mmPAL           KCMIODeviceAVCSignal = 0x86
	KCMIODeviceAVCSignalModeAudio            KCMIODeviceAVCSignal = 0x20
	KCMIODeviceAVCSignalModeDVCPro100_50     KCMIODeviceAVCSignal = 0xf0
	KCMIODeviceAVCSignalModeDVCPro100_60     KCMIODeviceAVCSignal = 0x70
	KCMIODeviceAVCSignalModeDVCPro25_525_60  KCMIODeviceAVCSignal = 0x78
	KCMIODeviceAVCSignalModeDVCPro25_625_50  KCMIODeviceAVCSignal = 0xf8
	KCMIODeviceAVCSignalModeDVCPro50_525_60  KCMIODeviceAVCSignal = 0x74
	KCMIODeviceAVCSignalModeDVCPro50_625_50  KCMIODeviceAVCSignal = 0xf4
	KCMIODeviceAVCSignalModeDVHS             KCMIODeviceAVCSignal = 0x1
	KCMIODeviceAVCSignalModeHD1125_60        KCMIODeviceAVCSignal = 0x8
	KCMIODeviceAVCSignalModeHD1250_50        KCMIODeviceAVCSignal = 0x88
	KCMIODeviceAVCSignalModeHDV1_50          KCMIODeviceAVCSignal = 0x90
	KCMIODeviceAVCSignalModeHDV1_60          KCMIODeviceAVCSignal = 0x10
	KCMIODeviceAVCSignalModeHDV2_50          KCMIODeviceAVCSignal = 0x9a
	KCMIODeviceAVCSignalModeHDV2_60          KCMIODeviceAVCSignal = 0x1a
	KCMIODeviceAVCSignalModeHi8NTSC          KCMIODeviceAVCSignal = 0xe
	KCMIODeviceAVCSignalModeHi8PAL           KCMIODeviceAVCSignal = 0x8e
	KCMIODeviceAVCSignalModeMPEG12Mbps_50    KCMIODeviceAVCSignal = 0x94
	KCMIODeviceAVCSignalModeMPEG12Mbps_60    KCMIODeviceAVCSignal = 0x14
	KCMIODeviceAVCSignalModeMPEG25Mbps_50    KCMIODeviceAVCSignal = 0x90
	KCMIODeviceAVCSignalModeMPEG25Mbps_60    KCMIODeviceAVCSignal = 0x10
	KCMIODeviceAVCSignalModeMPEG6Mbps_50     KCMIODeviceAVCSignal = 0x98
	KCMIODeviceAVCSignalModeMPEG6Mbps_60     KCMIODeviceAVCSignal = 0x18
	KCMIODeviceAVCSignalModeMicroMV12Mbps_50 KCMIODeviceAVCSignal = 0xa4
	KCMIODeviceAVCSignalModeMicroMV12Mbps_60 KCMIODeviceAVCSignal = 0x24
	KCMIODeviceAVCSignalModeMicroMV6Mbps_50  KCMIODeviceAVCSignal = 0xa8
	KCMIODeviceAVCSignalModeMicroMV6Mbps_60  KCMIODeviceAVCSignal = 0x28
	KCMIODeviceAVCSignalModeSD525_60         KCMIODeviceAVCSignal = 0
	KCMIODeviceAVCSignalModeSD625_50         KCMIODeviceAVCSignal = 0x80
	KCMIODeviceAVCSignalModeSDL525_60        KCMIODeviceAVCSignal = 0x4
	KCMIODeviceAVCSignalModeSDL625_50        KCMIODeviceAVCSignal = 0x84
	KCMIODeviceAVCSignalModeSVHS525_60       KCMIODeviceAVCSignal = 0xd
	KCMIODeviceAVCSignalModeSVHS625_50       KCMIODeviceAVCSignal = 0xed
	KCMIODeviceAVCSignalModeVHSMESECAM       KCMIODeviceAVCSignal = 0xd5
	KCMIODeviceAVCSignalModeVHSMPAL          KCMIODeviceAVCSignal = 0x25
	KCMIODeviceAVCSignalModeVHSNPAL          KCMIODeviceAVCSignal = 0xb5
	KCMIODeviceAVCSignalModeVHSNTSC          KCMIODeviceAVCSignal = 0x5
	KCMIODeviceAVCSignalModeVHSPAL           KCMIODeviceAVCSignal = 0xa5
	KCMIODeviceAVCSignalModeVHSSECAM         KCMIODeviceAVCSignal = 0xc5
)

func (e KCMIODeviceAVCSignal) String() string {
	switch e {
	case KCMIODeviceAVCSignalMode8mmNTSC:
		return "KCMIODeviceAVCSignalMode8mmNTSC"
	case KCMIODeviceAVCSignalMode8mmPAL:
		return "KCMIODeviceAVCSignalMode8mmPAL"
	case KCMIODeviceAVCSignalModeAudio:
		return "KCMIODeviceAVCSignalModeAudio"
	case KCMIODeviceAVCSignalModeDVCPro100_50:
		return "KCMIODeviceAVCSignalModeDVCPro100_50"
	case KCMIODeviceAVCSignalModeDVCPro100_60:
		return "KCMIODeviceAVCSignalModeDVCPro100_60"
	case KCMIODeviceAVCSignalModeDVCPro25_525_60:
		return "KCMIODeviceAVCSignalModeDVCPro25_525_60"
	case KCMIODeviceAVCSignalModeDVCPro25_625_50:
		return "KCMIODeviceAVCSignalModeDVCPro25_625_50"
	case KCMIODeviceAVCSignalModeDVCPro50_525_60:
		return "KCMIODeviceAVCSignalModeDVCPro50_525_60"
	case KCMIODeviceAVCSignalModeDVCPro50_625_50:
		return "KCMIODeviceAVCSignalModeDVCPro50_625_50"
	case KCMIODeviceAVCSignalModeDVHS:
		return "KCMIODeviceAVCSignalModeDVHS"
	case KCMIODeviceAVCSignalModeHD1125_60:
		return "KCMIODeviceAVCSignalModeHD1125_60"
	case KCMIODeviceAVCSignalModeHD1250_50:
		return "KCMIODeviceAVCSignalModeHD1250_50"
	case KCMIODeviceAVCSignalModeHDV1_50:
		return "KCMIODeviceAVCSignalModeHDV1_50"
	case KCMIODeviceAVCSignalModeHDV1_60:
		return "KCMIODeviceAVCSignalModeHDV1_60"
	case KCMIODeviceAVCSignalModeHDV2_50:
		return "KCMIODeviceAVCSignalModeHDV2_50"
	case KCMIODeviceAVCSignalModeHDV2_60:
		return "KCMIODeviceAVCSignalModeHDV2_60"
	case KCMIODeviceAVCSignalModeHi8NTSC:
		return "KCMIODeviceAVCSignalModeHi8NTSC"
	case KCMIODeviceAVCSignalModeHi8PAL:
		return "KCMIODeviceAVCSignalModeHi8PAL"
	case KCMIODeviceAVCSignalModeMPEG12Mbps_50:
		return "KCMIODeviceAVCSignalModeMPEG12Mbps_50"
	case KCMIODeviceAVCSignalModeMPEG12Mbps_60:
		return "KCMIODeviceAVCSignalModeMPEG12Mbps_60"
	case KCMIODeviceAVCSignalModeMPEG6Mbps_50:
		return "KCMIODeviceAVCSignalModeMPEG6Mbps_50"
	case KCMIODeviceAVCSignalModeMPEG6Mbps_60:
		return "KCMIODeviceAVCSignalModeMPEG6Mbps_60"
	case KCMIODeviceAVCSignalModeMicroMV12Mbps_50:
		return "KCMIODeviceAVCSignalModeMicroMV12Mbps_50"
	case KCMIODeviceAVCSignalModeMicroMV12Mbps_60:
		return "KCMIODeviceAVCSignalModeMicroMV12Mbps_60"
	case KCMIODeviceAVCSignalModeMicroMV6Mbps_50:
		return "KCMIODeviceAVCSignalModeMicroMV6Mbps_50"
	case KCMIODeviceAVCSignalModeMicroMV6Mbps_60:
		return "KCMIODeviceAVCSignalModeMicroMV6Mbps_60"
	case KCMIODeviceAVCSignalModeSD525_60:
		return "KCMIODeviceAVCSignalModeSD525_60"
	case KCMIODeviceAVCSignalModeSD625_50:
		return "KCMIODeviceAVCSignalModeSD625_50"
	case KCMIODeviceAVCSignalModeSDL525_60:
		return "KCMIODeviceAVCSignalModeSDL525_60"
	case KCMIODeviceAVCSignalModeSDL625_50:
		return "KCMIODeviceAVCSignalModeSDL625_50"
	case KCMIODeviceAVCSignalModeSVHS525_60:
		return "KCMIODeviceAVCSignalModeSVHS525_60"
	case KCMIODeviceAVCSignalModeSVHS625_50:
		return "KCMIODeviceAVCSignalModeSVHS625_50"
	case KCMIODeviceAVCSignalModeVHSMESECAM:
		return "KCMIODeviceAVCSignalModeVHSMESECAM"
	case KCMIODeviceAVCSignalModeVHSMPAL:
		return "KCMIODeviceAVCSignalModeVHSMPAL"
	case KCMIODeviceAVCSignalModeVHSNPAL:
		return "KCMIODeviceAVCSignalModeVHSNPAL"
	case KCMIODeviceAVCSignalModeVHSNTSC:
		return "KCMIODeviceAVCSignalModeVHSNTSC"
	case KCMIODeviceAVCSignalModeVHSPAL:
		return "KCMIODeviceAVCSignalModeVHSPAL"
	case KCMIODeviceAVCSignalModeVHSSECAM:
		return "KCMIODeviceAVCSignalModeVHSSECAM"
	default:
		return fmt.Sprintf("KCMIODeviceAVCSignal(%d)", e)
	}
}

type KCMIODeviceProperty uint32

const (
	KCMIODevicePropertyAVCDeviceSignalMode                  KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 's'<<8 | 'm' // 'pmsm'
	KCMIODevicePropertyAVCDeviceType                        KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'a'<<8 | 't' // 'pmat'
	KCMIODevicePropertyCanProcessAVCCommand                 KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'a'<<8 | 'c' // 'pmac'
	KCMIODevicePropertyCanProcessRS422Command               KCMIODeviceProperty = 'r'<<24 | '4'<<16 | '2'<<8 | '2' // 'r422'
	KCMIODevicePropertyCanSwitchFrameRatesWithoutFrameDrops KCMIODeviceProperty = 'f'<<24 | 'r'<<16 | 'n'<<8 | 'd' // 'frnd'
	KCMIODevicePropertyClientSyncDiscontinuity              KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'c'<<8 | 's' // 'pmcs'
	KCMIODevicePropertyDeviceCanBeDefaultDevice             KCMIODeviceProperty = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KCMIODevicePropertyDeviceControl                        KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | 'h' // 'pmnh'
	KCMIODevicePropertyDeviceHasChanged                     KCMIODeviceProperty = 'd'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'diff'
	KCMIODevicePropertyDeviceHasStreamingError              KCMIODeviceProperty = 's'<<24 | 'e'<<16 | 'r'<<8 | 'r' // 'serr'
	KCMIODevicePropertyDeviceIsAlive                        KCMIODeviceProperty = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'n' // 'livn'
	KCMIODevicePropertyDeviceIsRunning                      KCMIODeviceProperty = 'g'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'goin'
	KCMIODevicePropertyDeviceIsRunningSomewhere             KCMIODeviceProperty = 'g'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'gone'
	KCMIODevicePropertyDeviceUID                            KCMIODeviceProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'uid '
	KCMIODevicePropertyExcludeNonDALAccess                  KCMIODeviceProperty = 'i'<<24 | 'x'<<16 | 'n'<<8 | 'a' // 'ixna'
	KCMIODevicePropertyHogMode                              KCMIODeviceProperty = 'o'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'oink'
	KCMIODevicePropertyIIDCCSRData                          KCMIODeviceProperty = 'c'<<24 | 's'<<16 | 'r'<<8 | 'd' // 'csrd'
	KCMIODevicePropertyIIDCInitialUnitSpace                 KCMIODeviceProperty = 'i'<<24 | 'u'<<16 | 'n'<<8 | 's' // 'iuns'
	KCMIODevicePropertyLatency                              KCMIODeviceProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KCMIODevicePropertyLinkedAndSyncedCoreAudioDeviceUID    KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 's'<<8 | 'd' // 'plsd'
	KCMIODevicePropertyLinkedCoreAudioDeviceUID             KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'd' // 'plud'
	KCMIODevicePropertyLocationValue                        KCMIODeviceProperty = 'd'<<24 | 'l'<<16 | 'o'<<8 | 'c' // 'dloc'
	KCMIODevicePropertyModelUID                             KCMIODeviceProperty = 'm'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'muid'
	KCMIODevicePropertyPlugIn                               KCMIODeviceProperty = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'g' // 'plug'
	KCMIODevicePropertySMPTETimeCallback                    KCMIODeviceProperty = 'p'<<24 | 'm'<<16 | 's'<<8 | 'c' // 'pmsc'
	KCMIODevicePropertyStreamConfiguration                  KCMIODeviceProperty = 's'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'slay'
	KCMIODevicePropertyStreams                              KCMIODeviceProperty = 's'<<24 | 't'<<16 | 'm'<<8 | '#' // 'stm#'
	KCMIODevicePropertySuspendedByUser                      KCMIODeviceProperty = 's'<<24 | 'b'<<16 | 'y'<<8 | 'u' // 'sbyu'
	KCMIODevicePropertyTransportType                        KCMIODeviceProperty = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KCMIODevicePropertyVideoDigitizerComponents             KCMIODeviceProperty = 'v'<<24 | 'd'<<16 | 'i'<<8 | 'g' // 'vdig'
	// Deprecated: use KCMIODevicePropertyDeviceControl.
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
	case KCMIODevicePropertyLocationValue:
		return "KCMIODevicePropertyLocationValue"
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

type KCMIODevicePropertyLocation uint32

const (
	KCMIODevicePropertyLocationBuiltInDisplay         KCMIODevicePropertyLocation = 1
	KCMIODevicePropertyLocationExternalDevice         KCMIODevicePropertyLocation = 3
	KCMIODevicePropertyLocationExternalDisplay        KCMIODevicePropertyLocation = 2
	KCMIODevicePropertyLocationExternalWirelessDevice KCMIODevicePropertyLocation = 4
	KCMIODevicePropertyLocationUnknown                KCMIODevicePropertyLocation = 0
)

func (e KCMIODevicePropertyLocation) String() string {
	switch e {
	case KCMIODevicePropertyLocationBuiltInDisplay:
		return "KCMIODevicePropertyLocationBuiltInDisplay"
	case KCMIODevicePropertyLocationExternalDevice:
		return "KCMIODevicePropertyLocationExternalDevice"
	case KCMIODevicePropertyLocationExternalDisplay:
		return "KCMIODevicePropertyLocationExternalDisplay"
	case KCMIODevicePropertyLocationExternalWirelessDevice:
		return "KCMIODevicePropertyLocationExternalWirelessDevice"
	case KCMIODevicePropertyLocationUnknown:
		return "KCMIODevicePropertyLocationUnknown"
	default:
		return fmt.Sprintf("KCMIODevicePropertyLocation(%d)", e)
	}
}

type KCMIOExposureControlProperty uint32

const (
	KCMIOExposureControlPropertyConvergenceSpeed KCMIOExposureControlProperty = 'e'<<24 | 'c'<<16 | 's'<<8 | 'p' // 'ecsp'
	KCMIOExposureControlPropertyIntegrationTime  KCMIOExposureControlProperty = 'e'<<24 | 'i'<<16 | 'n'<<8 | 't' // 'eint'
	KCMIOExposureControlPropertyLockThreshold    KCMIOExposureControlProperty = 'e'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'elck'
	KCMIOExposureControlPropertyMaximumGain      KCMIOExposureControlProperty = 'e'<<24 | 'm'<<16 | 'a'<<8 | 'x' // 'emax'
	KCMIOExposureControlPropertyRegionOfInterest KCMIOExposureControlProperty = 'e'<<24 | 'r'<<16 | 'o'<<8 | 'i' // 'eroi'
	KCMIOExposureControlPropertyStability        KCMIOExposureControlProperty = 'e'<<24 | 's'<<16 | 't'<<8 | 'y' // 'esty'
	KCMIOExposureControlPropertyStable           KCMIOExposureControlProperty = 'e'<<24 | 's'<<16 | 't'<<8 | 'b' // 'estb'
	KCMIOExposureControlPropertyTarget           KCMIOExposureControlProperty = 'e'<<24 | 't'<<16 | 'g'<<8 | 't' // 'etgt'
	KCMIOExposureControlPropertyUnlockThreshold  KCMIOExposureControlProperty = 'e'<<24 | 'u'<<16 | 'l'<<8 | 'k' // 'eulk'
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

type KCMIOFeatureControlProperty uint32

const (
	KCMIOFeatureControlPropertyAbsoluteNative          KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'a' // 'fcna'
	KCMIOFeatureControlPropertyAbsoluteRange           KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'r' // 'fcar'
	KCMIOFeatureControlPropertyAbsoluteUnitName        KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'u'<<8 | 'n' // 'fcun'
	KCMIOFeatureControlPropertyAbsoluteValue           KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'v' // 'fcav'
	KCMIOFeatureControlPropertyAutomaticManual         KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'a'<<8 | 'm' // 'fcam'
	KCMIOFeatureControlPropertyConvertAbsoluteToNative KCMIOFeatureControlProperty = 'f'<<24 | 'a'<<16 | '2'<<8 | 'n' // 'fa2n'
	KCMIOFeatureControlPropertyConvertNativeToAbsolute KCMIOFeatureControlProperty = 'f'<<24 | 'n'<<16 | '2'<<8 | 'a' // 'fn2a'
	KCMIOFeatureControlPropertyNativeData              KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'd' // 'fcnd'
	KCMIOFeatureControlPropertyNativeDataRange         KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'd'<<8 | 'r' // 'fcdr'
	KCMIOFeatureControlPropertyNativeRange             KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'fcnr'
	KCMIOFeatureControlPropertyNativeValue             KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'n'<<8 | 'v' // 'fcnv'
	KCMIOFeatureControlPropertyOnOff                   KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 'o'<<8 | 'o' // 'fcoo'
	KCMIOFeatureControlPropertyTune                    KCMIOFeatureControlProperty = 'f'<<24 | 'c'<<16 | 't'<<8 | 'n' // 'fctn'
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

type KCMIOHardwareNoError uint32

const (
	KCMIODevicePermissionsError            KCMIOHardwareNoError = '!'<<24 | 'h'<<16 | 'o'<<8 | 'g' // '!hog'
	KCMIODeviceUnsupportedFormatError      KCMIOHardwareNoError = '!'<<24 | 'd'<<16 | 'a'<<8 | 't' // '!dat'
	KCMIOHardwareBadDeviceError            KCMIOHardwareNoError = '!'<<24 | 'd'<<16 | 'e'<<8 | 'v' // '!dev'
	KCMIOHardwareBadObjectError            KCMIOHardwareNoError = '!'<<24 | 'o'<<16 | 'b'<<8 | 'j' // '!obj'
	KCMIOHardwareBadPropertySizeError      KCMIOHardwareNoError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KCMIOHardwareBadStreamError            KCMIOHardwareNoError = '!'<<24 | 's'<<16 | 't'<<8 | 'r' // '!str'
	KCMIOHardwareIllegalOperationError     KCMIOHardwareNoError = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	KCMIOHardwareNoErrorValue              KCMIOHardwareNoError = 0
	KCMIOHardwareNotRunningError           KCMIOHardwareNoError = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KCMIOHardwareNotStoppedError           KCMIOHardwareNoError = 'r'<<24 | 'u'<<16 | 'n'<<8 | ' ' // 'run '
	KCMIOHardwareSuspendedBySystemError    KCMIOHardwareNoError = 'd'<<24 | 'e'<<16 | 'n'<<8 | 'y' // 'deny'
	KCMIOHardwareUnknownPropertyError      KCMIOHardwareNoError = 'w'<<24 | 'h'<<16 | 'o'<<8 | '?' // 'who?'
	KCMIOHardwareUnspecifiedError          KCMIOHardwareNoError = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	KCMIOHardwareUnsupportedOperationError KCMIOHardwareNoError = 'u'<<24 | 'n'<<16 | 'o'<<8 | 'p' // 'unop'
)

func (e KCMIOHardwareNoError) String() string {
	switch e {
	case KCMIODevicePermissionsError:
		return "KCMIODevicePermissionsError"
	case KCMIODeviceUnsupportedFormatError:
		return "KCMIODeviceUnsupportedFormatError"
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
	case KCMIOHardwareNoErrorValue:
		return "KCMIOHardwareNoErrorValue"
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
	default:
		return fmt.Sprintf("KCMIOHardwareNoError(%d)", e)
	}
}

type KCMIOHardwareProperty uint32

const (
	KCMIOHardwarePropertyAllowScreenCaptureDevices         KCMIOHardwareProperty = 'y'<<24 | 'e'<<16 | 's'<<8 | ' ' // 'yes '
	KCMIOHardwarePropertyAllowWirelessScreenCaptureDevices KCMIOHardwareProperty = 'w'<<24 | 's'<<16 | 'c'<<8 | 'd' // 'wscd'
	KCMIOHardwarePropertyDefaultInputDevice                KCMIOHardwareProperty = 'd'<<24 | 'I'<<16 | 'n'<<8 | ' ' // 'dIn '
	KCMIOHardwarePropertyDefaultOutputDevice               KCMIOHardwareProperty = 'd'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'dOut'
	KCMIOHardwarePropertyDeviceForUID                      KCMIOHardwareProperty = 'd'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'duid'
	KCMIOHardwarePropertyDevices                           KCMIOHardwareProperty = 'd'<<24 | 'e'<<16 | 'v'<<8 | '#' // 'dev#'
	KCMIOHardwarePropertyIsInitingOrExiting                KCMIOHardwareProperty = 'i'<<24 | 'n'<<16 | 'o'<<8 | 't' // 'inot'
	KCMIOHardwarePropertyPlugInForBundleID                 KCMIOHardwareProperty = 'p'<<24 | 'i'<<16 | 'b'<<8 | 'i' // 'pibi'
	KCMIOHardwarePropertyProcessIsMain                     KCMIOHardwareProperty = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'main'
	KCMIOHardwarePropertySleepingIsAllowed                 KCMIOHardwareProperty = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KCMIOHardwarePropertySuspendedBySystem                 KCMIOHardwareProperty = 's'<<24 | 'b'<<16 | 'y'<<8 | 's' // 'sbys'
	KCMIOHardwarePropertyUnloadingIsAllowed                KCMIOHardwareProperty = 'u'<<24 | 'n'<<16 | 'l'<<8 | 'd' // 'unld'
	KCMIOHardwarePropertyUserSessionIsActiveOrHeadless     KCMIOHardwareProperty = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
	// Deprecated: use KCMIOHardwarePropertyProcessIsMain.
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

type KCMIOJackControlClassID uint32

const (
	KCMIODirectionControlClassID KCMIOJackControlClassID = 'd'<<24 | 'i'<<16 | 'r'<<8 | 'e' // 'dire'
	KCMIOJackControlClassIDValue KCMIOJackControlClassID = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
)

func (e KCMIOJackControlClassID) String() string {
	switch e {
	case KCMIODirectionControlClassID:
		return "KCMIODirectionControlClassID"
	case KCMIOJackControlClassIDValue:
		return "KCMIOJackControlClassIDValue"
	default:
		return fmt.Sprintf("KCMIOJackControlClassID(%d)", e)
	}
}

type KCMIOObject uint32

const (
	KCMIOObjectClassID             KCMIOObject = 'a'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'aobj'
	KCMIOObjectClassIDWildcard     KCMIOObject = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KCMIOObjectPropertyElementMain KCMIOObject = 0
	KCMIOObjectPropertyScopeGlobal KCMIOObject = 'g'<<24 | 'l'<<16 | 'o'<<8 | 'b' // 'glob'
	KCMIOObjectUnknown             KCMIOObject = 0
	// Deprecated: use KCMIOObjectPropertyElementMain.
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

type KCMIOObjectPropertyClass uint32

const (
	KCMIOObjectPropertyClassValue          KCMIOObjectPropertyClass = 'c'<<24 | 'l'<<16 | 'a'<<8 | 's' // 'clas'
	KCMIOObjectPropertyCreator             KCMIOObjectPropertyClass = 'o'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'oplg'
	KCMIOObjectPropertyElementCategoryName KCMIOObjectPropertyClass = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KCMIOObjectPropertyElementName         KCMIOObjectPropertyClass = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KCMIOObjectPropertyElementNumberName   KCMIOObjectPropertyClass = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KCMIOObjectPropertyListenerAdded       KCMIOObjectPropertyClass = 'l'<<24 | 'i'<<16 | 's'<<8 | 'a' // 'lisa'
	KCMIOObjectPropertyListenerRemoved     KCMIOObjectPropertyClass = 'l'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'lisr'
	KCMIOObjectPropertyManufacturer        KCMIOObjectPropertyClass = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KCMIOObjectPropertyName                KCMIOObjectPropertyClass = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KCMIOObjectPropertyOwnedObjects        KCMIOObjectPropertyClass = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KCMIOObjectPropertyOwner               KCMIOObjectPropertyClass = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
)

func (e KCMIOObjectPropertyClass) String() string {
	switch e {
	case KCMIOObjectPropertyClassValue:
		return "KCMIOObjectPropertyClassValue"
	case KCMIOObjectPropertyCreator:
		return "KCMIOObjectPropertyCreator"
	case KCMIOObjectPropertyElementCategoryName:
		return "KCMIOObjectPropertyElementCategoryName"
	case KCMIOObjectPropertyElementName:
		return "KCMIOObjectPropertyElementName"
	case KCMIOObjectPropertyElementNumberName:
		return "KCMIOObjectPropertyElementNumberName"
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
	default:
		return fmt.Sprintf("KCMIOObjectPropertyClass(%d)", e)
	}
}

type KCMIOObjectPropertySelectorWildcard uint32

const (
	KCMIOObjectPropertyElementWildcard       KCMIOObjectPropertySelectorWildcard = 0xffffffff
	KCMIOObjectPropertyScopeWildcard         KCMIOObjectPropertySelectorWildcard = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KCMIOObjectPropertySelectorWildcardValue KCMIOObjectPropertySelectorWildcard = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KCMIOObjectPropertySelectorWildcard) String() string {
	switch e {
	case KCMIOObjectPropertyElementWildcard:
		return "KCMIOObjectPropertyElementWildcard"
	case KCMIOObjectPropertyScopeWildcard:
		return "KCMIOObjectPropertyScopeWildcard"
	default:
		return fmt.Sprintf("KCMIOObjectPropertySelectorWildcard(%d)", e)
	}
}

type KCMIOPlugInClassI uint32

const (
	KCMIOPlugInClassID KCMIOPlugInClassI = 'a'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'aplg'
)

func (e KCMIOPlugInClassI) String() string {
	switch e {
	case KCMIOPlugInClassID:
		return "KCMIOPlugInClassID"
	default:
		return fmt.Sprintf("KCMIOPlugInClassI(%d)", e)
	}
}

type KCMIOPlugInProperty uint32

const (
	KCMIOPlugInPropertyBundleID    KCMIOPlugInProperty = 'p'<<24 | 'i'<<16 | 'i'<<8 | 'd' // 'piid'
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

type KCMIOSampleBuffer uint32

const (
	KCMIOSampleBufferDiscontinuityFlag_BufferOverrun KCMIOSampleBuffer = 128
	// KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity KCMIOSampleBuffer = 1024
	KCMIOSampleBufferDiscontinuityFlag_CodecSettingsChanged    KCMIOSampleBuffer = 131072
	// KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged      KCMIOSampleBuffer = 8192
	KCMIOSampleBufferDiscontinuityFlag_DataWasDropped         KCMIOSampleBuffer = 64
	KCMIOSampleBufferDiscontinuityFlag_DataWasFlushed         KCMIOSampleBuffer = 32
	KCMIOSampleBufferDiscontinuityFlag_DiscontinuityInDTS     KCMIOSampleBuffer = 256
	KCMIOSampleBufferDiscontinuityFlag_DurationWasExtended    KCMIOSampleBuffer = 32768
	KCMIOSampleBufferDiscontinuityFlag_MalformedData          KCMIOSampleBuffer = 16
	KCMIOSampleBufferDiscontinuityFlag_NoDataMarker           KCMIOSampleBuffer = 4096
	KCMIOSampleBufferDiscontinuityFlag_PacketError            KCMIOSampleBuffer = 4
	KCMIOSampleBufferDiscontinuityFlag_RelatedToDiscontinuity KCMIOSampleBuffer = 512
	// KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle        KCMIOSampleBuffer = 65536
	KCMIOSampleBufferDiscontinuityFlag_StreamDiscontinuity   KCMIOSampleBuffer = 8
	KCMIOSampleBufferDiscontinuityFlag_TimecodeDiscontinuity KCMIOSampleBuffer = 2
	KCMIOSampleBufferDiscontinuityFlag_TimingReferenceJumped KCMIOSampleBuffer = 16384
	// KCMIOSampleBufferDiscontinuityFlag_TrickPlay: # Discussion
	KCMIOSampleBufferDiscontinuityFlag_TrickPlay            KCMIOSampleBuffer = 2048
	KCMIOSampleBufferDiscontinuityFlag_UnknownDiscontinuity KCMIOSampleBuffer = 1
	KCMIOSampleBufferNoDiscontinuities                      KCMIOSampleBuffer = 0
)

func (e KCMIOSampleBuffer) String() string {
	switch e {
	case KCMIOSampleBufferDiscontinuityFlag_BufferOverrun:
		return "KCMIOSampleBufferDiscontinuityFlag_BufferOverrun"
	case KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity:
		return "KCMIOSampleBufferDiscontinuityFlag_ClientSyncDiscontinuity"
	case KCMIOSampleBufferDiscontinuityFlag_CodecSettingsChanged:
		return "KCMIOSampleBufferDiscontinuityFlag_CodecSettingsChanged"
	case KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged:
		return "KCMIOSampleBufferDiscontinuityFlag_DataFormatChanged"
	case KCMIOSampleBufferDiscontinuityFlag_DataWasDropped:
		return "KCMIOSampleBufferDiscontinuityFlag_DataWasDropped"
	case KCMIOSampleBufferDiscontinuityFlag_DataWasFlushed:
		return "KCMIOSampleBufferDiscontinuityFlag_DataWasFlushed"
	case KCMIOSampleBufferDiscontinuityFlag_DiscontinuityInDTS:
		return "KCMIOSampleBufferDiscontinuityFlag_DiscontinuityInDTS"
	case KCMIOSampleBufferDiscontinuityFlag_DurationWasExtended:
		return "KCMIOSampleBufferDiscontinuityFlag_DurationWasExtended"
	case KCMIOSampleBufferDiscontinuityFlag_MalformedData:
		return "KCMIOSampleBufferDiscontinuityFlag_MalformedData"
	case KCMIOSampleBufferDiscontinuityFlag_NoDataMarker:
		return "KCMIOSampleBufferDiscontinuityFlag_NoDataMarker"
	case KCMIOSampleBufferDiscontinuityFlag_PacketError:
		return "KCMIOSampleBufferDiscontinuityFlag_PacketError"
	case KCMIOSampleBufferDiscontinuityFlag_RelatedToDiscontinuity:
		return "KCMIOSampleBufferDiscontinuityFlag_RelatedToDiscontinuity"
	case KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle:
		return "KCMIOSampleBufferDiscontinuityFlag_SleepWakeCycle"
	case KCMIOSampleBufferDiscontinuityFlag_StreamDiscontinuity:
		return "KCMIOSampleBufferDiscontinuityFlag_StreamDiscontinuity"
	case KCMIOSampleBufferDiscontinuityFlag_TimecodeDiscontinuity:
		return "KCMIOSampleBufferDiscontinuityFlag_TimecodeDiscontinuity"
	case KCMIOSampleBufferDiscontinuityFlag_TimingReferenceJumped:
		return "KCMIOSampleBufferDiscontinuityFlag_TimingReferenceJumped"
	case KCMIOSampleBufferDiscontinuityFlag_TrickPlay:
		return "KCMIOSampleBufferDiscontinuityFlag_TrickPlay"
	case KCMIOSampleBufferDiscontinuityFlag_UnknownDiscontinuity:
		return "KCMIOSampleBufferDiscontinuityFlag_UnknownDiscontinuity"
	case KCMIOSampleBufferNoDiscontinuities:
		return "KCMIOSampleBufferNoDiscontinuities"
	default:
		return fmt.Sprintf("KCMIOSampleBuffer(%d)", e)
	}
}

type KCMIOSampleBufferNoDataEvent uint32

const (
	KCMIOSampleBufferNoDataEvent_DeviceDidNotSync  KCMIOSampleBufferNoDataEvent = 2
	KCMIOSampleBufferNoDataEvent_DeviceInWrongMode KCMIOSampleBufferNoDataEvent = 3
	KCMIOSampleBufferNoDataEvent_NoMedia           KCMIOSampleBufferNoDataEvent = 1
	KCMIOSampleBufferNoDataEvent_ProcessingError   KCMIOSampleBufferNoDataEvent = 4
	KCMIOSampleBufferNoDataEvent_SleepWakeCycle    KCMIOSampleBufferNoDataEvent = 5
	KCMIOSampleBufferNoDataEvent_Unknown           KCMIOSampleBufferNoDataEvent = 0
)

func (e KCMIOSampleBufferNoDataEvent) String() string {
	switch e {
	case KCMIOSampleBufferNoDataEvent_DeviceDidNotSync:
		return "KCMIOSampleBufferNoDataEvent_DeviceDidNotSync"
	case KCMIOSampleBufferNoDataEvent_DeviceInWrongMode:
		return "KCMIOSampleBufferNoDataEvent_DeviceInWrongMode"
	case KCMIOSampleBufferNoDataEvent_NoMedia:
		return "KCMIOSampleBufferNoDataEvent_NoMedia"
	case KCMIOSampleBufferNoDataEvent_ProcessingError:
		return "KCMIOSampleBufferNoDataEvent_ProcessingError"
	case KCMIOSampleBufferNoDataEvent_SleepWakeCycle:
		return "KCMIOSampleBufferNoDataEvent_SleepWakeCycle"
	case KCMIOSampleBufferNoDataEvent_Unknown:
		return "KCMIOSampleBufferNoDataEvent_Unknown"
	default:
		return fmt.Sprintf("KCMIOSampleBufferNoDataEvent(%d)", e)
	}
}

type KCMIOSelectorControlProperty uint32

const (
	KCMIOSelectorControlPropertyAvailableItemNames KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'a'<<8 | 'n' // 'scan'
	KCMIOSelectorControlPropertyAvailableItems     KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'a'<<8 | 'i' // 'scai'
	KCMIOSelectorControlPropertyCurrentItem        KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'c'<<8 | 'i' // 'scci'
	KCMIOSelectorControlPropertyItemName           KCMIOSelectorControlProperty = 's'<<24 | 'c'<<16 | 'i'<<8 | 'n' // 'scin'
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

type KCMIOStream uint32

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

type KCMIOStreamProperty uint32

const (
	KCMIOStreamPropertyCanProcessDeckCommand                               KCMIOStreamProperty = 'p'<<24 | 'd'<<16 | 'c'<<8 | 'd' // 'pdcd'
	KCMIOStreamPropertyClock                                               KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'c'<<8 | 'l' // 'pmcl'
	KCMIOStreamPropertyDeck                                                KCMIOStreamProperty = 'd'<<24 | 'e'<<16 | 'c'<<8 | 'k' // 'deck'
	KCMIOStreamPropertyDeckCueing                                          KCMIOStreamProperty = 'c'<<24 | 'u'<<16 | 'e'<<8 | 'c' // 'cuec'
	KCMIOStreamPropertyDeckDropness                                        KCMIOStreamProperty = 'd'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'drop'
	KCMIOStreamPropertyDeckFrameNumber                                     KCMIOStreamProperty = 't'<<24 | 'c'<<16 | 'o'<<8 | 'd' // 'tcod'
	KCMIOStreamPropertyDeckLocal                                           KCMIOStreamProperty = 'l'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'locl'
	KCMIOStreamPropertyDeckThreaded                                        KCMIOStreamProperty = 't'<<24 | 'h'<<16 | 'r'<<8 | 'd' // 'thrd'
	KCMIOStreamPropertyDeviceSyncTimeoutInMSec                             KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '2' // 'pmn2'
	KCMIOStreamPropertyDirection                                           KCMIOStreamProperty = 's'<<24 | 'd'<<16 | 'i'<<8 | 'r' // 'sdir'
	KCMIOStreamPropertyEndOfData                                           KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'e'<<8 | 'd' // 'pmed'
	KCMIOStreamPropertyFirstOutputPresentationTimeStamp                    KCMIOStreamProperty = 'p'<<24 | 'o'<<16 | 'p'<<8 | 't' // 'popt'
	KCMIOStreamPropertyFormatDescription                                   KCMIOStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | ' ' // 'pft '
	KCMIOStreamPropertyFormatDescriptions                                  KCMIOStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | 'a' // 'pfta'
	KCMIOStreamPropertyFrameRate                                           KCMIOStreamProperty = 'n'<<24 | 'f'<<16 | 'r'<<8 | 't' // 'nfrt'
	KCMIOStreamPropertyFrameRateRanges                                     KCMIOStreamProperty = 'f'<<24 | 'r'<<16 | 'r'<<8 | 'g' // 'frrg'
	KCMIOStreamPropertyFrameRates                                          KCMIOStreamProperty = 'n'<<24 | 'f'<<16 | 'r'<<8 | '#' // 'nfr#'
	KCMIOStreamPropertyInitialPresentationTimeStampForLinkedAndSyncedAudio KCMIOStreamProperty = 'i'<<24 | 'p'<<16 | 'l'<<8 | 's' // 'ipls'
	KCMIOStreamPropertyLatency                                             KCMIOStreamProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KCMIOStreamPropertyMinimumFrameRate                                    KCMIOStreamProperty = 'm'<<24 | 'f'<<16 | 'r'<<8 | 't' // 'mfrt'
	KCMIOStreamPropertyNoDataEventCount                                    KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '3' // 'pmn3'
	KCMIOStreamPropertyNoDataTimeoutInMSec                                 KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'n'<<8 | '1' // 'pmn1'
	KCMIOStreamPropertyOutputBufferQueueSize                               KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'q' // 'pmoq'
	KCMIOStreamPropertyOutputBufferRepeatCount                             KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'r' // 'pmor'
	KCMIOStreamPropertyOutputBufferUnderrunCount                           KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'u' // 'pmou'
	KCMIOStreamPropertyOutputBuffersNeededForThrottledPlayback             KCMIOStreamProperty = 'm'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'miff'
	KCMIOStreamPropertyOutputBuffersRequiredForStartup                     KCMIOStreamProperty = 'p'<<24 | 'm'<<16 | 'o'<<8 | 's' // 'pmos'
	KCMIOStreamPropertyPreferredFormatDescription                          KCMIOStreamProperty = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'd' // 'prfd'
	KCMIOStreamPropertyPreferredFrameRate                                  KCMIOStreamProperty = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'r' // 'prfr'
	KCMIOStreamPropertyScheduledOutputNotificationProc                     KCMIOStreamProperty = 's'<<24 | 'o'<<16 | 'n'<<8 | 'p' // 'sonp'
	KCMIOStreamPropertyStartingChannel                                     KCMIOStreamProperty = 's'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'schn'
	KCMIOStreamPropertyStillImage                                          KCMIOStreamProperty = 's'<<24 | 't'<<16 | 'm'<<8 | 'g' // 'stmg'
	KCMIOStreamPropertyStillImageFormatDescriptions                        KCMIOStreamProperty = 's'<<24 | 't'<<16 | 'f'<<8 | 't' // 'stft'
	KCMIOStreamPropertyTerminalType                                        KCMIOStreamProperty = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
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

type KCMIOSystemObjectClassID uint32

const (
	KCMIOObjectSystemObject       KCMIOSystemObjectClassID = 1
	KCMIOSystemObjectClassIDValue KCMIOSystemObjectClassID = 'a'<<24 | 's'<<16 | 'y'<<8 | 's' // 'asys'
)

func (e KCMIOSystemObjectClassID) String() string {
	switch e {
	case KCMIOObjectSystemObject:
		return "KCMIOObjectSystemObject"
	case KCMIOSystemObjectClassIDValue:
		return "KCMIOSystemObjectClassIDValue"
	default:
		return fmt.Sprintf("KCMIOSystemObjectClassID(%d)", e)
	}
}
