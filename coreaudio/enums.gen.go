// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceClockAlgorithmSelector
type AudioDeviceClockAlgorithmSelector uint32

const (
	KAudioDeviceClockAlgorithm12PtMovingWindowAverage AudioDeviceClockAlgorithmSelector = 0
	KAudioDeviceClockAlgorithmRaw                     AudioDeviceClockAlgorithmSelector = 0
	KAudioDeviceClockAlgorithmSimpleIIR               AudioDeviceClockAlgorithmSelector = 0
)

func (e AudioDeviceClockAlgorithmSelector) String() string {
	switch e {
	case KAudioDeviceClockAlgorithm12PtMovingWindowAverage:
		return "KAudioDeviceClockAlgorithm12PtMovingWindowAverage"
	default:
		return fmt.Sprintf("AudioDeviceClockAlgorithmSelector(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePowerHint
type AudioHardwarePowerHint uint32

const (
	KAudioHardwarePowerHintFavorSavingPower AudioHardwarePowerHint = 1
	KAudioHardwarePowerHintNone             AudioHardwarePowerHint = 0
)

func (e AudioHardwarePowerHint) String() string {
	switch e {
	case KAudioHardwarePowerHintFavorSavingPower:
		return "KAudioHardwarePowerHintFavorSavingPower"
	case KAudioHardwarePowerHintNone:
		return "KAudioHardwarePowerHintNone"
	default:
		return fmt.Sprintf("AudioHardwarePowerHint(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction
type AudioLevelControlTransferFunction uint32

const (
	KAudioLevelControlTranferFunction10Over1 AudioLevelControlTransferFunction = 13
	KAudioLevelControlTranferFunction11Over1 AudioLevelControlTransferFunction = 14
	KAudioLevelControlTranferFunction12Over1 AudioLevelControlTransferFunction = 15
	KAudioLevelControlTranferFunction1Over2  AudioLevelControlTransferFunction = 2
	KAudioLevelControlTranferFunction1Over3  AudioLevelControlTransferFunction = 1
	KAudioLevelControlTranferFunction2Over1  AudioLevelControlTransferFunction = 5
	KAudioLevelControlTranferFunction3Over1  AudioLevelControlTransferFunction = 6
	KAudioLevelControlTranferFunction3Over2  AudioLevelControlTransferFunction = 4
	KAudioLevelControlTranferFunction3Over4  AudioLevelControlTransferFunction = 3
	KAudioLevelControlTranferFunction4Over1  AudioLevelControlTransferFunction = 7
	KAudioLevelControlTranferFunction5Over1  AudioLevelControlTransferFunction = 8
	KAudioLevelControlTranferFunction6Over1  AudioLevelControlTransferFunction = 9
	KAudioLevelControlTranferFunction7Over1  AudioLevelControlTransferFunction = 10
	KAudioLevelControlTranferFunction8Over1  AudioLevelControlTransferFunction = 11
	KAudioLevelControlTranferFunction9Over1  AudioLevelControlTransferFunction = 12
	KAudioLevelControlTranferFunctionLinear  AudioLevelControlTransferFunction = 0
)

func (e AudioLevelControlTransferFunction) String() string {
	switch e {
	case KAudioLevelControlTranferFunction10Over1:
		return "KAudioLevelControlTranferFunction10Over1"
	case KAudioLevelControlTranferFunction11Over1:
		return "KAudioLevelControlTranferFunction11Over1"
	case KAudioLevelControlTranferFunction12Over1:
		return "KAudioLevelControlTranferFunction12Over1"
	case KAudioLevelControlTranferFunction1Over2:
		return "KAudioLevelControlTranferFunction1Over2"
	case KAudioLevelControlTranferFunction1Over3:
		return "KAudioLevelControlTranferFunction1Over3"
	case KAudioLevelControlTranferFunction2Over1:
		return "KAudioLevelControlTranferFunction2Over1"
	case KAudioLevelControlTranferFunction3Over1:
		return "KAudioLevelControlTranferFunction3Over1"
	case KAudioLevelControlTranferFunction3Over2:
		return "KAudioLevelControlTranferFunction3Over2"
	case KAudioLevelControlTranferFunction3Over4:
		return "KAudioLevelControlTranferFunction3Over4"
	case KAudioLevelControlTranferFunction4Over1:
		return "KAudioLevelControlTranferFunction4Over1"
	case KAudioLevelControlTranferFunction5Over1:
		return "KAudioLevelControlTranferFunction5Over1"
	case KAudioLevelControlTranferFunction6Over1:
		return "KAudioLevelControlTranferFunction6Over1"
	case KAudioLevelControlTranferFunction7Over1:
		return "KAudioLevelControlTranferFunction7Over1"
	case KAudioLevelControlTranferFunction8Over1:
		return "KAudioLevelControlTranferFunction8Over1"
	case KAudioLevelControlTranferFunction9Over1:
		return "KAudioLevelControlTranferFunction9Over1"
	case KAudioLevelControlTranferFunctionLinear:
		return "KAudioLevelControlTranferFunctionLinear"
	default:
		return fmt.Sprintf("AudioLevelControlTransferFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation
type AudioServerPlugInIOOperation uint32

const (
	KAudioServerPlugInIOOperationConvertInput  AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationConvertMix    AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationCycle         AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationMixOutput     AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationProcessInput  AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationProcessMix    AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationProcessOutput AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationReadInput     AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationThread        AudioServerPlugInIOOperation = 0
	KAudioServerPlugInIOOperationWriteMix      AudioServerPlugInIOOperation = 0
)

func (e AudioServerPlugInIOOperation) String() string {
	switch e {
	case KAudioServerPlugInIOOperationConvertInput:
		return "KAudioServerPlugInIOOperationConvertInput"
	default:
		return fmt.Sprintf("AudioServerPlugInIOOperation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudio/CATapMuteBehavior
type CATapMuteBehavior int

const (
	CATapMuted           CATapMuteBehavior = 0
	CATapMutedWhenTapped CATapMuteBehavior = 0
	CATapUnmuted         CATapMuteBehavior = 0
)

func (e CATapMuteBehavior) String() string {
	switch e {
	case CATapMuted:
		return "CATapMuted"
	default:
		return fmt.Sprintf("CATapMuteBehavior(%d)", e)
	}
}

type KAudio uint

const (
	KAudioBooleanControlClassID                        KAudio = 't'<<24 | 'o'<<16 | 'g'<<8 | 'l' // 'togl'
	KAudioClipLightControlClassID                      KAudio = 'c'<<24 | 'l'<<16 | 'i'<<8 | 'p' // 'clip'
	KAudioClockSourceControlClassID                    KAudio = 'c'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'clck'
	KAudioDataDestinationControlClassID                KAudio = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KAudioDataSourceControlClassID                     KAudio = 'd'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'dsrc'
	KAudioDevicePermissionsError                       KAudio = '!'<<24 | 'h'<<16 | 'o'<<8 | 'g' // '!hog'
	KAudioDevicePropertyAvailableNominalSampleRates    KAudio = 'n'<<24 | 's'<<16 | 'r'<<8 | '#' // 'nsr#'
	KAudioDevicePropertyClockDomain                    KAudio = 'c'<<24 | 'l'<<16 | 'k'<<8 | 'd' // 'clkd'
	KAudioDevicePropertyConfigurationApplication       KAudio = 'c'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'capp'
	KAudioDevicePropertyDeviceCanBeDefaultDevice       KAudio = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KAudioDevicePropertyDeviceCanBeDefaultSystemDevice KAudio = 's'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'sflt'
	KAudioDevicePropertyDeviceIsAlive                  KAudio = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'n' // 'livn'
	KAudioDevicePropertyDeviceIsRunning                KAudio = 'g'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'goin'
	KAudioDevicePropertyDeviceUID                      KAudio = 'u'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'uid '
	KAudioDevicePropertyIcon                           KAudio = 'i'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'icon'
	KAudioDevicePropertyIsHidden                       KAudio = 'h'<<24 | 'i'<<16 | 'd'<<8 | 'n' // 'hidn'
	KAudioDevicePropertyLatency                        KAudio = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KAudioDevicePropertyModelUID                       KAudio = 'm'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'muid'
	KAudioDevicePropertyNominalSampleRate              KAudio = 'n'<<24 | 's'<<16 | 'r'<<8 | 't' // 'nsrt'
	KAudioDevicePropertyPreferredChannelLayout         KAudio = 's'<<24 | 'r'<<16 | 'n'<<8 | 'd' // 'srnd'
	KAudioDevicePropertyPreferredChannelsForStereo     KAudio = 'd'<<24 | 'c'<<16 | 'h'<<8 | '2' // 'dch2'
	KAudioDevicePropertyRelatedDevices                 KAudio = 'a'<<24 | 'k'<<16 | 'i'<<8 | 'n' // 'akin'
	KAudioDevicePropertySafetyOffset                   KAudio = 's'<<24 | 'a'<<16 | 'f'<<8 | 't' // 'saft'
	KAudioDevicePropertyStreams                        KAudio = 's'<<24 | 't'<<16 | 'm'<<8 | '#' // 'stm#'
	KAudioDevicePropertyTransportType                  KAudio = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KAudioDeviceUnsupportedFormatError                 KAudio = '!'<<24 | 'd'<<16 | 'a'<<8 | 't' // '!dat'
	KAudioHardwareBadDeviceError                       KAudio = '!'<<24 | 'd'<<16 | 'e'<<8 | 'v' // '!dev'
	KAudioHardwareBadObjectError                       KAudio = '!'<<24 | 'o'<<16 | 'b'<<8 | 'j' // '!obj'
	KAudioHardwareBadPropertySizeError                 KAudio = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KAudioHardwareBadStreamError                       KAudio = '!'<<24 | 's'<<16 | 't'<<8 | 'r' // '!str'
	KAudioHardwareIllegalOperationError                KAudio = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	KAudioHardwareNoError                              KAudio = 0
	KAudioHardwareNotReadyError                        KAudio = 'n'<<24 | 'r'<<16 | 'd'<<8 | 'y' // 'nrdy'
	KAudioHardwareNotRunningError                      KAudio = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KAudioHardwareUnknownPropertyError                 KAudio = 'w'<<24 | 'h'<<16 | 'o'<<8 | '?' // 'who?'
	KAudioHardwareUnspecifiedError                     KAudio = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	KAudioHardwareUnsupportedOperationError            KAudio = 'u'<<24 | 'n'<<16 | 'o'<<8 | 'p' // 'unop'
	KAudioHighPassFilterControlClassID                 KAudio = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'f' // 'hipf'
	KAudioJackControlClassID                           KAudio = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
	KAudioLFEMuteControlClassID                        KAudio = 's'<<24 | 'u'<<16 | 'b'<<8 | 'm' // 'subm'
	KAudioLFEVolumeControlClassID                      KAudio = 's'<<24 | 'u'<<16 | 'b'<<8 | 'v' // 'subv'
	KAudioLevelControlClassID                          KAudio = 'l'<<24 | 'e'<<16 | 'v'<<8 | 'l' // 'levl'
	KAudioLineLevelControlClassID                      KAudio = 'n'<<24 | 'l'<<16 | 'v'<<8 | 'l' // 'nlvl'
	KAudioListenbackControlClassID                     KAudio = 'l'<<24 | 's'<<16 | 'n'<<8 | 'b' // 'lsnb'
	KAudioMuteControlClassID                           KAudio = 'm'<<24 | 'u'<<16 | 't'<<8 | 'e' // 'mute'
	KAudioObjectPropertyControlList                    KAudio = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
	KAudioPhantomPowerControlClassID                   KAudio = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'n' // 'phan'
	KAudioPhaseInvertControlClassID                    KAudio = 'p'<<24 | 'h'<<16 | 's'<<8 | 'i' // 'phsi'
	KAudioSelectorControlClassID                       KAudio = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
	KAudioSoloControlClassID                           KAudio = 's'<<24 | 'o'<<16 | 'l'<<8 | 'o' // 'solo'
	KAudioTalkbackControlClassID                       KAudio = 't'<<24 | 'a'<<16 | 'l'<<8 | 'b' // 'talb'
	KAudioVolumeControlClassID                         KAudio = 'v'<<24 | 'l'<<16 | 'm'<<8 | 'e' // 'vlme'
)

func (e KAudio) String() string {
	switch e {
	case KAudioBooleanControlClassID:
		return "KAudioBooleanControlClassID"
	case KAudioClipLightControlClassID:
		return "KAudioClipLightControlClassID"
	case KAudioClockSourceControlClassID:
		return "KAudioClockSourceControlClassID"
	case KAudioDataDestinationControlClassID:
		return "KAudioDataDestinationControlClassID"
	case KAudioDataSourceControlClassID:
		return "KAudioDataSourceControlClassID"
	case KAudioDevicePermissionsError:
		return "KAudioDevicePermissionsError"
	case KAudioDevicePropertyAvailableNominalSampleRates:
		return "KAudioDevicePropertyAvailableNominalSampleRates"
	case KAudioDevicePropertyClockDomain:
		return "KAudioDevicePropertyClockDomain"
	case KAudioDevicePropertyConfigurationApplication:
		return "KAudioDevicePropertyConfigurationApplication"
	case KAudioDevicePropertyDeviceCanBeDefaultDevice:
		return "KAudioDevicePropertyDeviceCanBeDefaultDevice"
	case KAudioDevicePropertyDeviceCanBeDefaultSystemDevice:
		return "KAudioDevicePropertyDeviceCanBeDefaultSystemDevice"
	case KAudioDevicePropertyDeviceIsAlive:
		return "KAudioDevicePropertyDeviceIsAlive"
	case KAudioDevicePropertyDeviceIsRunning:
		return "KAudioDevicePropertyDeviceIsRunning"
	case KAudioDevicePropertyDeviceUID:
		return "KAudioDevicePropertyDeviceUID"
	case KAudioDevicePropertyIcon:
		return "KAudioDevicePropertyIcon"
	case KAudioDevicePropertyIsHidden:
		return "KAudioDevicePropertyIsHidden"
	case KAudioDevicePropertyLatency:
		return "KAudioDevicePropertyLatency"
	case KAudioDevicePropertyModelUID:
		return "KAudioDevicePropertyModelUID"
	case KAudioDevicePropertyNominalSampleRate:
		return "KAudioDevicePropertyNominalSampleRate"
	case KAudioDevicePropertyPreferredChannelLayout:
		return "KAudioDevicePropertyPreferredChannelLayout"
	case KAudioDevicePropertyPreferredChannelsForStereo:
		return "KAudioDevicePropertyPreferredChannelsForStereo"
	case KAudioDevicePropertyRelatedDevices:
		return "KAudioDevicePropertyRelatedDevices"
	case KAudioDevicePropertySafetyOffset:
		return "KAudioDevicePropertySafetyOffset"
	case KAudioDevicePropertyStreams:
		return "KAudioDevicePropertyStreams"
	case KAudioDevicePropertyTransportType:
		return "KAudioDevicePropertyTransportType"
	case KAudioDeviceUnsupportedFormatError:
		return "KAudioDeviceUnsupportedFormatError"
	case KAudioHardwareBadDeviceError:
		return "KAudioHardwareBadDeviceError"
	case KAudioHardwareBadObjectError:
		return "KAudioHardwareBadObjectError"
	case KAudioHardwareBadPropertySizeError:
		return "KAudioHardwareBadPropertySizeError"
	case KAudioHardwareBadStreamError:
		return "KAudioHardwareBadStreamError"
	case KAudioHardwareIllegalOperationError:
		return "KAudioHardwareIllegalOperationError"
	case KAudioHardwareNoError:
		return "KAudioHardwareNoError"
	case KAudioHardwareNotReadyError:
		return "KAudioHardwareNotReadyError"
	case KAudioHardwareNotRunningError:
		return "KAudioHardwareNotRunningError"
	case KAudioHardwareUnknownPropertyError:
		return "KAudioHardwareUnknownPropertyError"
	case KAudioHardwareUnspecifiedError:
		return "KAudioHardwareUnspecifiedError"
	case KAudioHardwareUnsupportedOperationError:
		return "KAudioHardwareUnsupportedOperationError"
	case KAudioHighPassFilterControlClassID:
		return "KAudioHighPassFilterControlClassID"
	case KAudioJackControlClassID:
		return "KAudioJackControlClassID"
	case KAudioLFEMuteControlClassID:
		return "KAudioLFEMuteControlClassID"
	case KAudioLFEVolumeControlClassID:
		return "KAudioLFEVolumeControlClassID"
	case KAudioLevelControlClassID:
		return "KAudioLevelControlClassID"
	case KAudioLineLevelControlClassID:
		return "KAudioLineLevelControlClassID"
	case KAudioListenbackControlClassID:
		return "KAudioListenbackControlClassID"
	case KAudioMuteControlClassID:
		return "KAudioMuteControlClassID"
	case KAudioObjectPropertyControlList:
		return "KAudioObjectPropertyControlList"
	case KAudioPhantomPowerControlClassID:
		return "KAudioPhantomPowerControlClassID"
	case KAudioPhaseInvertControlClassID:
		return "KAudioPhaseInvertControlClassID"
	case KAudioSelectorControlClassID:
		return "KAudioSelectorControlClassID"
	case KAudioSoloControlClassID:
		return "KAudioSoloControlClassID"
	case KAudioTalkbackControlClassID:
		return "KAudioTalkbackControlClassID"
	case KAudioVolumeControlClassID:
		return "KAudioVolumeControlClassID"
	default:
		return fmt.Sprintf("KAudio(%d)", e)
	}
}

type KAudioAggregateDeviceClassI uint

const (
	KAudioAggregateDeviceClassID KAudioAggregateDeviceClassI = 'a'<<24 | 'a'<<16 | 'g'<<8 | 'g' // 'aagg'
)

func (e KAudioAggregateDeviceClassI) String() string {
	switch e {
	case KAudioAggregateDeviceClassID:
		return "KAudioAggregateDeviceClassID"
	default:
		return fmt.Sprintf("KAudioAggregateDeviceClassI(%d)", e)
	}
}

type KAudioAggregateDeviceProperty uint

const (
	KAudioAggregateDevicePropertyActiveSubDeviceList KAudioAggregateDeviceProperty = 'a'<<24 | 'g'<<16 | 'r'<<8 | 'p' // 'agrp'
	KAudioAggregateDevicePropertyClockDevice         KAudioAggregateDeviceProperty = 'a'<<24 | 'p'<<16 | 'c'<<8 | 'd' // 'apcd'
	KAudioAggregateDevicePropertyComposition         KAudioAggregateDeviceProperty = 'a'<<24 | 'c'<<16 | 'o'<<8 | 'm' // 'acom'
	KAudioAggregateDevicePropertyFullSubDeviceList   KAudioAggregateDeviceProperty = 'g'<<24 | 'r'<<16 | 'u'<<8 | 'p' // 'grup'
	KAudioAggregateDevicePropertyMainSubDevice       KAudioAggregateDeviceProperty = 'a'<<24 | 'm'<<16 | 's'<<8 | 't' // 'amst'
	KAudioAggregateDevicePropertySubTapList          KAudioAggregateDeviceProperty = 'a'<<24 | 't'<<16 | 'a'<<8 | 'p' // 'atap'
	KAudioAggregateDevicePropertyTapList             KAudioAggregateDeviceProperty = 't'<<24 | 'a'<<16 | 'p'<<8 | '#' // 'tap#'
)

func (e KAudioAggregateDeviceProperty) String() string {
	switch e {
	case KAudioAggregateDevicePropertyActiveSubDeviceList:
		return "KAudioAggregateDevicePropertyActiveSubDeviceList"
	case KAudioAggregateDevicePropertyClockDevice:
		return "KAudioAggregateDevicePropertyClockDevice"
	case KAudioAggregateDevicePropertyComposition:
		return "KAudioAggregateDevicePropertyComposition"
	case KAudioAggregateDevicePropertyFullSubDeviceList:
		return "KAudioAggregateDevicePropertyFullSubDeviceList"
	case KAudioAggregateDevicePropertyMainSubDevice:
		return "KAudioAggregateDevicePropertyMainSubDevice"
	case KAudioAggregateDevicePropertySubTapList:
		return "KAudioAggregateDevicePropertySubTapList"
	case KAudioAggregateDevicePropertyTapList:
		return "KAudioAggregateDevicePropertyTapList"
	default:
		return fmt.Sprintf("KAudioAggregateDeviceProperty(%d)", e)
	}
}

type KAudioAggregateDevicePropertyMasterSub uint

const (
	// Deprecated.
	KAudioAggregateDevicePropertyMasterSubDevice KAudioAggregateDevicePropertyMasterSub = 0
)

func (e KAudioAggregateDevicePropertyMasterSub) String() string {
	switch e {
	case KAudioAggregateDevicePropertyMasterSubDevice:
		return "KAudioAggregateDevicePropertyMasterSubDevice"
	default:
		return fmt.Sprintf("KAudioAggregateDevicePropertyMasterSub(%d)", e)
	}
}

type KAudioAggregateDriftCompensation uint

const (
	KAudioAggregateDriftCompensationHighQuality   KAudioAggregateDriftCompensation = 0
	KAudioAggregateDriftCompensationLowQuality    KAudioAggregateDriftCompensation = 0
	KAudioAggregateDriftCompensationMaxQuality    KAudioAggregateDriftCompensation = 0
	KAudioAggregateDriftCompensationMediumQuality KAudioAggregateDriftCompensation = 0
	KAudioAggregateDriftCompensationMinQuality    KAudioAggregateDriftCompensation = 0
)

func (e KAudioAggregateDriftCompensation) String() string {
	switch e {
	case KAudioAggregateDriftCompensationHighQuality:
		return "KAudioAggregateDriftCompensationHighQuality"
	default:
		return fmt.Sprintf("KAudioAggregateDriftCompensation(%d)", e)
	}
}

type KAudioBooleanControlProperty uint

const (
	KAudioBooleanControlPropertyValue KAudioBooleanControlProperty = 'b'<<24 | 'c'<<16 | 'v'<<8 | 'l' // 'bcvl'
)

func (e KAudioBooleanControlProperty) String() string {
	switch e {
	case KAudioBooleanControlPropertyValue:
		return "KAudioBooleanControlPropertyValue"
	default:
		return fmt.Sprintf("KAudioBooleanControlProperty(%d)", e)
	}
}

type KAudioBootChimeVolumeControlClassI uint

const (
	KAudioBootChimeVolumeControlClassID KAudioBootChimeVolumeControlClassI = 'p'<<24 | 'r'<<16 | 'a'<<8 | 'm' // 'pram'
)

func (e KAudioBootChimeVolumeControlClassI) String() string {
	switch e {
	case KAudioBootChimeVolumeControlClassID:
		return "KAudioBootChimeVolumeControlClassID"
	default:
		return fmt.Sprintf("KAudioBootChimeVolumeControlClassI(%d)", e)
	}
}

type KAudioBoxClassI uint

const (
	KAudioBoxClassID KAudioBoxClassI = 'a'<<24 | 'b'<<16 | 'o'<<8 | 'x' // 'abox'
)

func (e KAudioBoxClassI) String() string {
	switch e {
	case KAudioBoxClassID:
		return "KAudioBoxClassID"
	default:
		return fmt.Sprintf("KAudioBoxClassI(%d)", e)
	}
}

type KAudioBoxProperty uint

const (
	KAudioBoxPropertyAcquired          KAudioBoxProperty = 'b'<<24 | 'x'<<16 | 'o'<<8 | 'n' // 'bxon'
	KAudioBoxPropertyAcquisitionFailed KAudioBoxProperty = 'b'<<24 | 'x'<<16 | 'o'<<8 | 'f' // 'bxof'
	KAudioBoxPropertyBoxUID            KAudioBoxProperty = 'b'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'buid'
	KAudioBoxPropertyClockDeviceList   KAudioBoxProperty = 'b'<<24 | 'c'<<16 | 'l'<<8 | '#' // 'bcl#'
	KAudioBoxPropertyDeviceList        KAudioBoxProperty = 'b'<<24 | 'd'<<16 | 'v'<<8 | '#' // 'bdv#'
	KAudioBoxPropertyHasAudio          KAudioBoxProperty = 'b'<<24 | 'h'<<16 | 'a'<<8 | 'u' // 'bhau'
	KAudioBoxPropertyHasMIDI           KAudioBoxProperty = 'b'<<24 | 'h'<<16 | 'm'<<8 | 'i' // 'bhmi'
	KAudioBoxPropertyHasVideo          KAudioBoxProperty = 'b'<<24 | 'h'<<16 | 'v'<<8 | 'i' // 'bhvi'
	KAudioBoxPropertyIsProtected       KAudioBoxProperty = 'b'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'bpro'
	KAudioBoxPropertyTransportType     KAudioBoxProperty = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
)

func (e KAudioBoxProperty) String() string {
	switch e {
	case KAudioBoxPropertyAcquired:
		return "KAudioBoxPropertyAcquired"
	case KAudioBoxPropertyAcquisitionFailed:
		return "KAudioBoxPropertyAcquisitionFailed"
	case KAudioBoxPropertyBoxUID:
		return "KAudioBoxPropertyBoxUID"
	case KAudioBoxPropertyClockDeviceList:
		return "KAudioBoxPropertyClockDeviceList"
	case KAudioBoxPropertyDeviceList:
		return "KAudioBoxPropertyDeviceList"
	case KAudioBoxPropertyHasAudio:
		return "KAudioBoxPropertyHasAudio"
	case KAudioBoxPropertyHasMIDI:
		return "KAudioBoxPropertyHasMIDI"
	case KAudioBoxPropertyHasVideo:
		return "KAudioBoxPropertyHasVideo"
	case KAudioBoxPropertyIsProtected:
		return "KAudioBoxPropertyIsProtected"
	case KAudioBoxPropertyTransportType:
		return "KAudioBoxPropertyTransportType"
	default:
		return fmt.Sprintf("KAudioBoxProperty(%d)", e)
	}
}

type KAudioClockDeviceClassI uint

const (
	KAudioClockDeviceClassID KAudioClockDeviceClassI = 'a'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'aclk'
)

func (e KAudioClockDeviceClassI) String() string {
	switch e {
	case KAudioClockDeviceClassID:
		return "KAudioClockDeviceClassID"
	default:
		return fmt.Sprintf("KAudioClockDeviceClassI(%d)", e)
	}
}

type KAudioClockDeviceProperty uint

const (
	KAudioClockDevicePropertyAvailableNominalSampleRates KAudioClockDeviceProperty = 'n'<<24 | 's'<<16 | 'r'<<8 | '#' // 'nsr#'
	KAudioClockDevicePropertyClockDomain                 KAudioClockDeviceProperty = 'c'<<24 | 'l'<<16 | 'k'<<8 | 'd' // 'clkd'
	KAudioClockDevicePropertyControlList                 KAudioClockDeviceProperty = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
	KAudioClockDevicePropertyDeviceIsAlive               KAudioClockDeviceProperty = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'n' // 'livn'
	KAudioClockDevicePropertyDeviceIsRunning             KAudioClockDeviceProperty = 'g'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'goin'
	KAudioClockDevicePropertyDeviceUID                   KAudioClockDeviceProperty = 'c'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'cuid'
	KAudioClockDevicePropertyLatency                     KAudioClockDeviceProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KAudioClockDevicePropertyNominalSampleRate           KAudioClockDeviceProperty = 'n'<<24 | 's'<<16 | 'r'<<8 | 't' // 'nsrt'
	KAudioClockDevicePropertyTransportType               KAudioClockDeviceProperty = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
)

func (e KAudioClockDeviceProperty) String() string {
	switch e {
	case KAudioClockDevicePropertyAvailableNominalSampleRates:
		return "KAudioClockDevicePropertyAvailableNominalSampleRates"
	case KAudioClockDevicePropertyClockDomain:
		return "KAudioClockDevicePropertyClockDomain"
	case KAudioClockDevicePropertyControlList:
		return "KAudioClockDevicePropertyControlList"
	case KAudioClockDevicePropertyDeviceIsAlive:
		return "KAudioClockDevicePropertyDeviceIsAlive"
	case KAudioClockDevicePropertyDeviceIsRunning:
		return "KAudioClockDevicePropertyDeviceIsRunning"
	case KAudioClockDevicePropertyDeviceUID:
		return "KAudioClockDevicePropertyDeviceUID"
	case KAudioClockDevicePropertyLatency:
		return "KAudioClockDevicePropertyLatency"
	case KAudioClockDevicePropertyNominalSampleRate:
		return "KAudioClockDevicePropertyNominalSampleRate"
	case KAudioClockDevicePropertyTransportType:
		return "KAudioClockDevicePropertyTransportType"
	default:
		return fmt.Sprintf("KAudioClockDeviceProperty(%d)", e)
	}
}

type KAudioClockSourceControlPropertyItem uint

const (
	KAudioClockSourceControlPropertyItemKind KAudioClockSourceControlPropertyItem = 0
)

func (e KAudioClockSourceControlPropertyItem) String() string {
	switch e {
	case KAudioClockSourceControlPropertyItemKind:
		return "KAudioClockSourceControlPropertyItemKind"
	default:
		return fmt.Sprintf("KAudioClockSourceControlPropertyItem(%d)", e)
	}
}

type KAudioClockSourceItemKind uint

const (
	KAudioClockSourceItemKindInternal KAudioClockSourceItemKind = 'i'<<24 | 'n'<<16 | 't'<<8 | ' ' // 'int '
)

func (e KAudioClockSourceItemKind) String() string {
	switch e {
	case KAudioClockSourceItemKindInternal:
		return "KAudioClockSourceItemKindInternal"
	default:
		return fmt.Sprintf("KAudioClockSourceItemKind(%d)", e)
	}
}

type KAudioControlClassI uint

const (
	KAudioControlClassID KAudioControlClassI = 'a'<<24 | 'c'<<16 | 't'<<8 | 'l' // 'actl'
)

func (e KAudioControlClassI) String() string {
	switch e {
	case KAudioControlClassID:
		return "KAudioControlClassID"
	default:
		return fmt.Sprintf("KAudioControlClassI(%d)", e)
	}
}

type KAudioControlProperty uint

const (
	KAudioControlPropertyElement KAudioControlProperty = 'c'<<24 | 'e'<<16 | 'l'<<8 | 'm' // 'celm'
	KAudioControlPropertyScope   KAudioControlProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'cscp'
	KAudioControlPropertyVariant KAudioControlProperty = 'c'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'cvar'
)

func (e KAudioControlProperty) String() string {
	switch e {
	case KAudioControlPropertyElement:
		return "KAudioControlPropertyElement"
	case KAudioControlPropertyScope:
		return "KAudioControlPropertyScope"
	case KAudioControlPropertyVariant:
		return "KAudioControlPropertyVariant"
	default:
		return fmt.Sprintf("KAudioControlProperty(%d)", e)
	}
}

type KAudioDevice uint

const (
	KAudioDeviceProcessorOverload                KAudioDevice = 'o'<<24 | 'v'<<16 | 'e'<<8 | 'r' // 'over'
	KAudioDevicePropertyActualSampleRate         KAudioDevice = 'a'<<24 | 's'<<16 | 'r'<<8 | 't' // 'asrt'
	KAudioDevicePropertyBufferFrameSize          KAudioDevice = 'f'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'fsiz'
	KAudioDevicePropertyBufferFrameSizeRange     KAudioDevice = 'f'<<24 | 's'<<16 | 'z'<<8 | '#' // 'fsz#'
	KAudioDevicePropertyClockDevice              KAudioDevice = 'a'<<24 | 'p'<<16 | 'c'<<8 | 'd' // 'apcd'
	KAudioDevicePropertyDeviceHasChanged         KAudioDevice = 'd'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'diff'
	KAudioDevicePropertyDeviceIsRunningSomewhere KAudioDevice = 'g'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'gone'
	KAudioDevicePropertyHogMode                  KAudioDevice = 'o'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'oink'
	KAudioDevicePropertyIOCycleUsage             KAudioDevice = 'n'<<24 | 'c'<<16 | 'y'<<8 | 'c' // 'ncyc'
	KAudioDevicePropertyIOProcStreamUsage        KAudioDevice = 's'<<24 | 'u'<<16 | 's'<<8 | 'e' // 'suse'
	KAudioDevicePropertyIOStoppedAbnormally      KAudioDevice = 's'<<24 | 't'<<16 | 'p'<<8 | 'd' // 'stpd'
	// KAudioDevicePropertyIOThreadOSWorkgroup: The device’s workgroup object, which you use to coordinate your threads with the threads of the device.
	KAudioDevicePropertyIOThreadOSWorkgroup          KAudioDevice = 'o'<<24 | 's'<<16 | 'w'<<8 | 'g' // 'oswg'
	KAudioDevicePropertyPlugIn                       KAudioDevice = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'g' // 'plug'
	KAudioDevicePropertyProcessMute                  KAudioDevice = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'm' // 'appm'
	KAudioDevicePropertyStreamConfiguration          KAudioDevice = 's'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'slay'
	KAudioDevicePropertyUsesVariableBufferFrameSizes KAudioDevice = 'v'<<24 | 'f'<<16 | 's'<<8 | 'z' // 'vfsz'
	KAudioDeviceUnknown                              KAudioDevice = 0
)

func (e KAudioDevice) String() string {
	switch e {
	case KAudioDeviceProcessorOverload:
		return "KAudioDeviceProcessorOverload"
	case KAudioDevicePropertyActualSampleRate:
		return "KAudioDevicePropertyActualSampleRate"
	case KAudioDevicePropertyBufferFrameSize:
		return "KAudioDevicePropertyBufferFrameSize"
	case KAudioDevicePropertyBufferFrameSizeRange:
		return "KAudioDevicePropertyBufferFrameSizeRange"
	case KAudioDevicePropertyClockDevice:
		return "KAudioDevicePropertyClockDevice"
	case KAudioDevicePropertyDeviceHasChanged:
		return "KAudioDevicePropertyDeviceHasChanged"
	case KAudioDevicePropertyDeviceIsRunningSomewhere:
		return "KAudioDevicePropertyDeviceIsRunningSomewhere"
	case KAudioDevicePropertyHogMode:
		return "KAudioDevicePropertyHogMode"
	case KAudioDevicePropertyIOCycleUsage:
		return "KAudioDevicePropertyIOCycleUsage"
	case KAudioDevicePropertyIOProcStreamUsage:
		return "KAudioDevicePropertyIOProcStreamUsage"
	case KAudioDevicePropertyIOStoppedAbnormally:
		return "KAudioDevicePropertyIOStoppedAbnormally"
	case KAudioDevicePropertyIOThreadOSWorkgroup:
		return "KAudioDevicePropertyIOThreadOSWorkgroup"
	case KAudioDevicePropertyPlugIn:
		return "KAudioDevicePropertyPlugIn"
	case KAudioDevicePropertyProcessMute:
		return "KAudioDevicePropertyProcessMute"
	case KAudioDevicePropertyStreamConfiguration:
		return "KAudioDevicePropertyStreamConfiguration"
	case KAudioDevicePropertyUsesVariableBufferFrameSizes:
		return "KAudioDevicePropertyUsesVariableBufferFrameSizes"
	case KAudioDeviceUnknown:
		return "KAudioDeviceUnknown"
	default:
		return fmt.Sprintf("KAudioDevice(%d)", e)
	}
}

type KAudioDeviceClassI uint

const (
	KAudioDeviceClassID KAudioDeviceClassI = 'a'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'adev'
)

func (e KAudioDeviceClassI) String() string {
	switch e {
	case KAudioDeviceClassID:
		return "KAudioDeviceClassID"
	default:
		return fmt.Sprintf("KAudioDeviceClassI(%d)", e)
	}
}

type KAudioDeviceProperty uint

const (
	KAudioDevicePropertyBufferSize                                     KAudioDeviceProperty = 'b'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'bsiz'
	KAudioDevicePropertyBufferSizeRange                                KAudioDeviceProperty = 'b'<<24 | 's'<<16 | 'z'<<8 | '#' // 'bsz#'
	KAudioDevicePropertyChannelCategoryName                            KAudioDeviceProperty = 'c'<<24 | 'c'<<16 | 'n'<<8 | 'm' // 'ccnm'
	KAudioDevicePropertyChannelCategoryNameCFString                    KAudioDeviceProperty = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KAudioDevicePropertyChannelName                                    KAudioDeviceProperty = 'c'<<24 | 'h'<<16 | 'n'<<8 | 'm' // 'chnm'
	KAudioDevicePropertyChannelNameCFString                            KAudioDeviceProperty = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KAudioDevicePropertyChannelNominalLineLevel                        KAudioDeviceProperty = 'n'<<24 | 'l'<<16 | 'v'<<8 | 'l' // 'nlvl'
	KAudioDevicePropertyChannelNominalLineLevelNameForID               KAudioDeviceProperty = 'c'<<24 | 'n'<<16 | 'l'<<8 | 'v' // 'cnlv'
	KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString       KAudioDeviceProperty = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'l' // 'lcnl'
	KAudioDevicePropertyChannelNominalLineLevels                       KAudioDeviceProperty = 'n'<<24 | 'l'<<16 | 'v'<<8 | '#' // 'nlv#'
	KAudioDevicePropertyChannelNumberName                              KAudioDeviceProperty = 'c'<<24 | 'n'<<16 | 'n'<<8 | 'm' // 'cnnm'
	KAudioDevicePropertyChannelNumberNameCFString                      KAudioDeviceProperty = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KAudioDevicePropertyClipLight                                      KAudioDeviceProperty = 'c'<<24 | 'l'<<16 | 'i'<<8 | 'p' // 'clip'
	KAudioDevicePropertyClockAlgorithm                                 KAudioDeviceProperty = 0
	KAudioDevicePropertyClockIsStable                                  KAudioDeviceProperty = 0
	KAudioDevicePropertyClockSource                                    KAudioDeviceProperty = 'c'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'csrc'
	KAudioDevicePropertyClockSourceKindForID                           KAudioDeviceProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | 'k' // 'csck'
	KAudioDevicePropertyClockSourceNameForID                           KAudioDeviceProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'cscn'
	KAudioDevicePropertyClockSourceNameForIDCFString                   KAudioDeviceProperty = 'l'<<24 | 'c'<<16 | 's'<<8 | 'n' // 'lcsn'
	KAudioDevicePropertyClockSources                                   KAudioDeviceProperty = 'c'<<24 | 's'<<16 | 'c'<<8 | '#' // 'csc#'
	KAudioDevicePropertyDataSource                                     KAudioDeviceProperty = 's'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'ssrc'
	KAudioDevicePropertyDataSourceKindForID                            KAudioDeviceProperty = 's'<<24 | 's'<<16 | 'c'<<8 | 'k' // 'ssck'
	KAudioDevicePropertyDataSourceNameForID                            KAudioDeviceProperty = 's'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'sscn'
	KAudioDevicePropertyDataSourceNameForIDCFString                    KAudioDeviceProperty = 'l'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'lscn'
	KAudioDevicePropertyDataSources                                    KAudioDeviceProperty = 's'<<24 | 's'<<16 | 'c'<<8 | '#' // 'ssc#'
	KAudioDevicePropertyDeviceManufacturer                             KAudioDeviceProperty = 'm'<<24 | 'a'<<16 | 'k'<<8 | 'r' // 'makr'
	KAudioDevicePropertyDeviceManufacturerCFString                     KAudioDeviceProperty = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KAudioDevicePropertyDeviceName                                     KAudioDeviceProperty = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	KAudioDevicePropertyDeviceNameCFString                             KAudioDeviceProperty = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KAudioDevicePropertyDriverShouldOwniSub                            KAudioDeviceProperty = 'i'<<24 | 's'<<16 | 'u'<<8 | 'b' // 'isub'
	KAudioDevicePropertyHighPassFilterSetting                          KAudioDeviceProperty = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'f' // 'hipf'
	KAudioDevicePropertyHighPassFilterSettingNameForID                 KAudioDeviceProperty = 'c'<<24 | 'h'<<16 | 'i'<<8 | 'p' // 'chip'
	KAudioDevicePropertyHighPassFilterSettingNameForIDCFString         KAudioDeviceProperty = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'l' // 'hipl'
	KAudioDevicePropertyHighPassFilterSettings                         KAudioDeviceProperty = 'h'<<24 | 'i'<<16 | 'p'<<8 | '#' // 'hip#'
	KAudioDevicePropertyJackIsConnected                                KAudioDeviceProperty = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
	KAudioDevicePropertyListenback                                     KAudioDeviceProperty = 'l'<<24 | 's'<<16 | 'n'<<8 | 'b' // 'lsnb'
	KAudioDevicePropertyMute                                           KAudioDeviceProperty = 'm'<<24 | 'u'<<16 | 't'<<8 | 'e' // 'mute'
	KAudioDevicePropertyPhantomPower                                   KAudioDeviceProperty = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'n' // 'phan'
	KAudioDevicePropertyPhaseInvert                                    KAudioDeviceProperty = 'p'<<24 | 'h'<<16 | 's'<<8 | 'i' // 'phsi'
	KAudioDevicePropertyPlayThru                                       KAudioDeviceProperty = 't'<<24 | 'h'<<16 | 'r'<<8 | 'u' // 'thru'
	KAudioDevicePropertyPlayThruDestination                            KAudioDeviceProperty = 'm'<<24 | 'd'<<16 | 'd'<<8 | 's' // 'mdds'
	KAudioDevicePropertyPlayThruDestinationNameForID                   KAudioDeviceProperty = 'm'<<24 | 'd'<<16 | 'd'<<8 | 'n' // 'mddn'
	KAudioDevicePropertyPlayThruDestinationNameForIDCFString           KAudioDeviceProperty = 'm'<<24 | 'd'<<16 | 'd'<<8 | 'c' // 'mddc'
	KAudioDevicePropertyPlayThruDestinations                           KAudioDeviceProperty = 'm'<<24 | 'd'<<16 | 'd'<<8 | '#' // 'mdd#'
	KAudioDevicePropertyPlayThruSolo                                   KAudioDeviceProperty = 't'<<24 | 'h'<<16 | 'r'<<8 | 's' // 'thrs'
	KAudioDevicePropertyPlayThruStereoPan                              KAudioDeviceProperty = 'm'<<24 | 's'<<16 | 'p'<<8 | 'n' // 'mspn'
	KAudioDevicePropertyPlayThruStereoPanChannels                      KAudioDeviceProperty = 'm'<<24 | 's'<<16 | 'p'<<8 | '#' // 'msp#'
	KAudioDevicePropertyPlayThruVolumeDecibels                         KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | 'd'<<8 | 'b' // 'mvdb'
	KAudioDevicePropertyPlayThruVolumeDecibelsToScalar                 KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | '2'<<8 | 's' // 'mv2s'
	KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | 't'<<8 | 'f' // 'mvtf'
	KAudioDevicePropertyPlayThruVolumeRangeDecibels                    KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | 'd'<<8 | '#' // 'mvd#'
	KAudioDevicePropertyPlayThruVolumeScalar                           KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | 's'<<8 | 'c' // 'mvsc'
	KAudioDevicePropertyPlayThruVolumeScalarToDecibels                 KAudioDeviceProperty = 'm'<<24 | 'v'<<16 | '2'<<8 | 'd' // 'mv2d'
	KAudioDevicePropertyRegisterBufferList                             KAudioDeviceProperty = 'r'<<24 | 'b'<<16 | 'u'<<8 | 'f' // 'rbuf'
	KAudioDevicePropertySolo                                           KAudioDeviceProperty = 's'<<24 | 'o'<<16 | 'l'<<8 | 'o' // 'solo'
	KAudioDevicePropertyStereoPan                                      KAudioDeviceProperty = 's'<<24 | 'p'<<16 | 'a'<<8 | 'n' // 'span'
	KAudioDevicePropertyStereoPanChannels                              KAudioDeviceProperty = 's'<<24 | 'p'<<16 | 'n'<<8 | '#' // 'spn#'
	KAudioDevicePropertyStreamFormat                                   KAudioDeviceProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'sfmt'
	KAudioDevicePropertyStreamFormatMatch                              KAudioDeviceProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | 'm' // 'sfmm'
	KAudioDevicePropertyStreamFormatSupported                          KAudioDeviceProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | '?' // 'sfm?'
	KAudioDevicePropertyStreamFormats                                  KAudioDeviceProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | '#' // 'sfm#'
	KAudioDevicePropertySubMute                                        KAudioDeviceProperty = 's'<<24 | 'm'<<16 | 'u'<<8 | 't' // 'smut'
	KAudioDevicePropertySubVolumeDecibels                              KAudioDeviceProperty = 's'<<24 | 'v'<<16 | 'l'<<8 | 'd' // 'svld'
	KAudioDevicePropertySubVolumeDecibelsToScalar                      KAudioDeviceProperty = 's'<<24 | 'd'<<16 | '2'<<8 | 'v' // 'sd2v'
	KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction      KAudioDeviceProperty = 's'<<24 | 'v'<<16 | 't'<<8 | 'f' // 'svtf'
	KAudioDevicePropertySubVolumeRangeDecibels                         KAudioDeviceProperty = 's'<<24 | 'v'<<16 | 'd'<<8 | '#' // 'svd#'
	KAudioDevicePropertySubVolumeScalar                                KAudioDeviceProperty = 's'<<24 | 'v'<<16 | 'l'<<8 | 'm' // 'svlm'
	KAudioDevicePropertySubVolumeScalarToDecibels                      KAudioDeviceProperty = 's'<<24 | 'v'<<16 | '2'<<8 | 'd' // 'sv2d'
	KAudioDevicePropertySupportsMixing                                 KAudioDeviceProperty = 'm'<<24 | 'i'<<16 | 'x'<<8 | '?' // 'mix?'
	KAudioDevicePropertyTalkback                                       KAudioDeviceProperty = 't'<<24 | 'a'<<16 | 'l'<<8 | 'b' // 'talb'
	KAudioDevicePropertyVoiceActivityDetectionEnable                   KAudioDeviceProperty = 'v'<<24 | 'A'<<16 | 'd'<<8 | '+' // 'vAd+'
	KAudioDevicePropertyVoiceActivityDetectionState                    KAudioDeviceProperty = 'v'<<24 | 'A'<<16 | 'd'<<8 | 'S' // 'vAdS'
	KAudioDevicePropertyVolumeDecibels                                 KAudioDeviceProperty = 'v'<<24 | 'o'<<16 | 'l'<<8 | 'd' // 'vold'
	KAudioDevicePropertyVolumeDecibelsToScalar                         KAudioDeviceProperty = 'd'<<24 | 'b'<<16 | '2'<<8 | 'v' // 'db2v'
	KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction         KAudioDeviceProperty = 'v'<<24 | 'c'<<16 | 't'<<8 | 'f' // 'vctf'
	KAudioDevicePropertyVolumeRangeDecibels                            KAudioDeviceProperty = 'v'<<24 | 'd'<<16 | 'b'<<8 | '#' // 'vdb#'
	KAudioDevicePropertyVolumeScalar                                   KAudioDeviceProperty = 'v'<<24 | 'o'<<16 | 'l'<<8 | 'm' // 'volm'
	KAudioDevicePropertyVolumeScalarToDecibels                         KAudioDeviceProperty = 'v'<<24 | '2'<<16 | 'd'<<8 | 'b' // 'v2db'
	KAudioDevicePropertyWantsControlsRestored                          KAudioDeviceProperty = 'r'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'resc'
	KAudioDevicePropertyWantsStreamFormatsRestored                     KAudioDeviceProperty = 'r'<<24 | 'e'<<16 | 's'<<8 | 'f' // 'resf'
	KAudioDevicePropertyZeroTimeStampPeriod                            KAudioDeviceProperty = 0
)

func (e KAudioDeviceProperty) String() string {
	switch e {
	case KAudioDevicePropertyBufferSize:
		return "KAudioDevicePropertyBufferSize"
	case KAudioDevicePropertyBufferSizeRange:
		return "KAudioDevicePropertyBufferSizeRange"
	case KAudioDevicePropertyChannelCategoryName:
		return "KAudioDevicePropertyChannelCategoryName"
	case KAudioDevicePropertyChannelCategoryNameCFString:
		return "KAudioDevicePropertyChannelCategoryNameCFString"
	case KAudioDevicePropertyChannelName:
		return "KAudioDevicePropertyChannelName"
	case KAudioDevicePropertyChannelNameCFString:
		return "KAudioDevicePropertyChannelNameCFString"
	case KAudioDevicePropertyChannelNominalLineLevel:
		return "KAudioDevicePropertyChannelNominalLineLevel"
	case KAudioDevicePropertyChannelNominalLineLevelNameForID:
		return "KAudioDevicePropertyChannelNominalLineLevelNameForID"
	case KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString:
		return "KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString"
	case KAudioDevicePropertyChannelNominalLineLevels:
		return "KAudioDevicePropertyChannelNominalLineLevels"
	case KAudioDevicePropertyChannelNumberName:
		return "KAudioDevicePropertyChannelNumberName"
	case KAudioDevicePropertyChannelNumberNameCFString:
		return "KAudioDevicePropertyChannelNumberNameCFString"
	case KAudioDevicePropertyClipLight:
		return "KAudioDevicePropertyClipLight"
	case KAudioDevicePropertyClockAlgorithm:
		return "KAudioDevicePropertyClockAlgorithm"
	case KAudioDevicePropertyClockSource:
		return "KAudioDevicePropertyClockSource"
	case KAudioDevicePropertyClockSourceKindForID:
		return "KAudioDevicePropertyClockSourceKindForID"
	case KAudioDevicePropertyClockSourceNameForID:
		return "KAudioDevicePropertyClockSourceNameForID"
	case KAudioDevicePropertyClockSourceNameForIDCFString:
		return "KAudioDevicePropertyClockSourceNameForIDCFString"
	case KAudioDevicePropertyClockSources:
		return "KAudioDevicePropertyClockSources"
	case KAudioDevicePropertyDataSource:
		return "KAudioDevicePropertyDataSource"
	case KAudioDevicePropertyDataSourceKindForID:
		return "KAudioDevicePropertyDataSourceKindForID"
	case KAudioDevicePropertyDataSourceNameForID:
		return "KAudioDevicePropertyDataSourceNameForID"
	case KAudioDevicePropertyDataSourceNameForIDCFString:
		return "KAudioDevicePropertyDataSourceNameForIDCFString"
	case KAudioDevicePropertyDataSources:
		return "KAudioDevicePropertyDataSources"
	case KAudioDevicePropertyDeviceManufacturer:
		return "KAudioDevicePropertyDeviceManufacturer"
	case KAudioDevicePropertyDeviceManufacturerCFString:
		return "KAudioDevicePropertyDeviceManufacturerCFString"
	case KAudioDevicePropertyDeviceName:
		return "KAudioDevicePropertyDeviceName"
	case KAudioDevicePropertyDeviceNameCFString:
		return "KAudioDevicePropertyDeviceNameCFString"
	case KAudioDevicePropertyDriverShouldOwniSub:
		return "KAudioDevicePropertyDriverShouldOwniSub"
	case KAudioDevicePropertyHighPassFilterSetting:
		return "KAudioDevicePropertyHighPassFilterSetting"
	case KAudioDevicePropertyHighPassFilterSettingNameForID:
		return "KAudioDevicePropertyHighPassFilterSettingNameForID"
	case KAudioDevicePropertyHighPassFilterSettingNameForIDCFString:
		return "KAudioDevicePropertyHighPassFilterSettingNameForIDCFString"
	case KAudioDevicePropertyHighPassFilterSettings:
		return "KAudioDevicePropertyHighPassFilterSettings"
	case KAudioDevicePropertyJackIsConnected:
		return "KAudioDevicePropertyJackIsConnected"
	case KAudioDevicePropertyListenback:
		return "KAudioDevicePropertyListenback"
	case KAudioDevicePropertyMute:
		return "KAudioDevicePropertyMute"
	case KAudioDevicePropertyPhantomPower:
		return "KAudioDevicePropertyPhantomPower"
	case KAudioDevicePropertyPhaseInvert:
		return "KAudioDevicePropertyPhaseInvert"
	case KAudioDevicePropertyPlayThru:
		return "KAudioDevicePropertyPlayThru"
	case KAudioDevicePropertyPlayThruDestination:
		return "KAudioDevicePropertyPlayThruDestination"
	case KAudioDevicePropertyPlayThruDestinationNameForID:
		return "KAudioDevicePropertyPlayThruDestinationNameForID"
	case KAudioDevicePropertyPlayThruDestinationNameForIDCFString:
		return "KAudioDevicePropertyPlayThruDestinationNameForIDCFString"
	case KAudioDevicePropertyPlayThruDestinations:
		return "KAudioDevicePropertyPlayThruDestinations"
	case KAudioDevicePropertyPlayThruSolo:
		return "KAudioDevicePropertyPlayThruSolo"
	case KAudioDevicePropertyPlayThruStereoPan:
		return "KAudioDevicePropertyPlayThruStereoPan"
	case KAudioDevicePropertyPlayThruStereoPanChannels:
		return "KAudioDevicePropertyPlayThruStereoPanChannels"
	case KAudioDevicePropertyPlayThruVolumeDecibels:
		return "KAudioDevicePropertyPlayThruVolumeDecibels"
	case KAudioDevicePropertyPlayThruVolumeDecibelsToScalar:
		return "KAudioDevicePropertyPlayThruVolumeDecibelsToScalar"
	case KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction:
		return "KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction"
	case KAudioDevicePropertyPlayThruVolumeRangeDecibels:
		return "KAudioDevicePropertyPlayThruVolumeRangeDecibels"
	case KAudioDevicePropertyPlayThruVolumeScalar:
		return "KAudioDevicePropertyPlayThruVolumeScalar"
	case KAudioDevicePropertyPlayThruVolumeScalarToDecibels:
		return "KAudioDevicePropertyPlayThruVolumeScalarToDecibels"
	case KAudioDevicePropertyRegisterBufferList:
		return "KAudioDevicePropertyRegisterBufferList"
	case KAudioDevicePropertySolo:
		return "KAudioDevicePropertySolo"
	case KAudioDevicePropertyStereoPan:
		return "KAudioDevicePropertyStereoPan"
	case KAudioDevicePropertyStereoPanChannels:
		return "KAudioDevicePropertyStereoPanChannels"
	case KAudioDevicePropertyStreamFormat:
		return "KAudioDevicePropertyStreamFormat"
	case KAudioDevicePropertyStreamFormatMatch:
		return "KAudioDevicePropertyStreamFormatMatch"
	case KAudioDevicePropertyStreamFormatSupported:
		return "KAudioDevicePropertyStreamFormatSupported"
	case KAudioDevicePropertyStreamFormats:
		return "KAudioDevicePropertyStreamFormats"
	case KAudioDevicePropertySubMute:
		return "KAudioDevicePropertySubMute"
	case KAudioDevicePropertySubVolumeDecibels:
		return "KAudioDevicePropertySubVolumeDecibels"
	case KAudioDevicePropertySubVolumeDecibelsToScalar:
		return "KAudioDevicePropertySubVolumeDecibelsToScalar"
	case KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction:
		return "KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction"
	case KAudioDevicePropertySubVolumeRangeDecibels:
		return "KAudioDevicePropertySubVolumeRangeDecibels"
	case KAudioDevicePropertySubVolumeScalar:
		return "KAudioDevicePropertySubVolumeScalar"
	case KAudioDevicePropertySubVolumeScalarToDecibels:
		return "KAudioDevicePropertySubVolumeScalarToDecibels"
	case KAudioDevicePropertySupportsMixing:
		return "KAudioDevicePropertySupportsMixing"
	case KAudioDevicePropertyTalkback:
		return "KAudioDevicePropertyTalkback"
	case KAudioDevicePropertyVoiceActivityDetectionEnable:
		return "KAudioDevicePropertyVoiceActivityDetectionEnable"
	case KAudioDevicePropertyVoiceActivityDetectionState:
		return "KAudioDevicePropertyVoiceActivityDetectionState"
	case KAudioDevicePropertyVolumeDecibels:
		return "KAudioDevicePropertyVolumeDecibels"
	case KAudioDevicePropertyVolumeDecibelsToScalar:
		return "KAudioDevicePropertyVolumeDecibelsToScalar"
	case KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction:
		return "KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction"
	case KAudioDevicePropertyVolumeRangeDecibels:
		return "KAudioDevicePropertyVolumeRangeDecibels"
	case KAudioDevicePropertyVolumeScalar:
		return "KAudioDevicePropertyVolumeScalar"
	case KAudioDevicePropertyVolumeScalarToDecibels:
		return "KAudioDevicePropertyVolumeScalarToDecibels"
	case KAudioDevicePropertyWantsControlsRestored:
		return "KAudioDevicePropertyWantsControlsRestored"
	case KAudioDevicePropertyWantsStreamFormatsRestored:
		return "KAudioDevicePropertyWantsStreamFormatsRestored"
	default:
		return fmt.Sprintf("KAudioDeviceProperty(%d)", e)
	}
}

type KAudioDevicePropertyScope uint

const (
	KAudioDevicePropertyScopeInput       KAudioDevicePropertyScope = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KAudioDevicePropertyScopeOutput      KAudioDevicePropertyScope = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KAudioDevicePropertyScopePlayThrough KAudioDevicePropertyScope = 0
)

func (e KAudioDevicePropertyScope) String() string {
	switch e {
	case KAudioDevicePropertyScopeInput:
		return "KAudioDevicePropertyScopeInput"
	case KAudioDevicePropertyScopeOutput:
		return "KAudioDevicePropertyScopeOutput"
	case KAudioDevicePropertyScopePlayThrough:
		return "KAudioDevicePropertyScopePlayThrough"
	default:
		return fmt.Sprintf("KAudioDevicePropertyScope(%d)", e)
	}
}

type KAudioDeviceStartTime uint

const (
	KAudioDeviceStartTimeDontConsultDeviceFlag KAudioDeviceStartTime = 0
	KAudioDeviceStartTimeDontConsultHALFlag    KAudioDeviceStartTime = 0
	KAudioDeviceStartTimeIsInputFlag           KAudioDeviceStartTime = 0
)

func (e KAudioDeviceStartTime) String() string {
	switch e {
	case KAudioDeviceStartTimeDontConsultDeviceFlag:
		return "KAudioDeviceStartTimeDontConsultDeviceFlag"
	default:
		return fmt.Sprintf("KAudioDeviceStartTime(%d)", e)
	}
}

type KAudioDeviceTransportType uint

const (
	KAudioDeviceTransportTypeAVB                       KAudioDeviceTransportType = 'e'<<24 | 'a'<<16 | 'v'<<8 | 'b' // 'eavb'
	KAudioDeviceTransportTypeAggregate                 KAudioDeviceTransportType = 'g'<<24 | 'r'<<16 | 'u'<<8 | 'p' // 'grup'
	KAudioDeviceTransportTypeAirPlay                   KAudioDeviceTransportType = 'a'<<24 | 'i'<<16 | 'r'<<8 | 'p' // 'airp'
	KAudioDeviceTransportTypeBluetooth                 KAudioDeviceTransportType = 'b'<<24 | 'l'<<16 | 'u'<<8 | 'e' // 'blue'
	KAudioDeviceTransportTypeBluetoothLE               KAudioDeviceTransportType = 'b'<<24 | 'l'<<16 | 'e'<<8 | 'a' // 'blea'
	KAudioDeviceTransportTypeBuiltIn                   KAudioDeviceTransportType = 'b'<<24 | 'l'<<16 | 't'<<8 | 'n' // 'bltn'
	KAudioDeviceTransportTypeContinuityCaptureWired    KAudioDeviceTransportType = 'c'<<24 | 'c'<<16 | 'w'<<8 | 'd' // 'ccwd'
	KAudioDeviceTransportTypeContinuityCaptureWireless KAudioDeviceTransportType = 'c'<<24 | 'c'<<16 | 'w'<<8 | 'l' // 'ccwl'
	KAudioDeviceTransportTypeDisplayPort               KAudioDeviceTransportType = 'd'<<24 | 'p'<<16 | 'r'<<8 | 't' // 'dprt'
	KAudioDeviceTransportTypeFireWire                  KAudioDeviceTransportType = '1'<<24 | '3'<<16 | '9'<<8 | '4' // '1394'
	KAudioDeviceTransportTypeHDMI                      KAudioDeviceTransportType = 'h'<<24 | 'd'<<16 | 'm'<<8 | 'i' // 'hdmi'
	KAudioDeviceTransportTypePCI                       KAudioDeviceTransportType = 'p'<<24 | 'c'<<16 | 'i'<<8 | ' ' // 'pci '
	KAudioDeviceTransportTypeThunderbolt               KAudioDeviceTransportType = 't'<<24 | 'h'<<16 | 'u'<<8 | 'n' // 'thun'
	KAudioDeviceTransportTypeUSB                       KAudioDeviceTransportType = 'u'<<24 | 's'<<16 | 'b'<<8 | ' ' // 'usb '
	KAudioDeviceTransportTypeUnknown                   KAudioDeviceTransportType = 0
	KAudioDeviceTransportTypeVirtual                   KAudioDeviceTransportType = 'v'<<24 | 'i'<<16 | 'r'<<8 | 't' // 'virt'
	// Deprecated.
	KAudioDeviceTransportTypeContinuityCapture KAudioDeviceTransportType = 0
)

func (e KAudioDeviceTransportType) String() string {
	switch e {
	case KAudioDeviceTransportTypeAVB:
		return "KAudioDeviceTransportTypeAVB"
	case KAudioDeviceTransportTypeAggregate:
		return "KAudioDeviceTransportTypeAggregate"
	case KAudioDeviceTransportTypeAirPlay:
		return "KAudioDeviceTransportTypeAirPlay"
	case KAudioDeviceTransportTypeBluetooth:
		return "KAudioDeviceTransportTypeBluetooth"
	case KAudioDeviceTransportTypeBluetoothLE:
		return "KAudioDeviceTransportTypeBluetoothLE"
	case KAudioDeviceTransportTypeBuiltIn:
		return "KAudioDeviceTransportTypeBuiltIn"
	case KAudioDeviceTransportTypeContinuityCaptureWired:
		return "KAudioDeviceTransportTypeContinuityCaptureWired"
	case KAudioDeviceTransportTypeContinuityCaptureWireless:
		return "KAudioDeviceTransportTypeContinuityCaptureWireless"
	case KAudioDeviceTransportTypeDisplayPort:
		return "KAudioDeviceTransportTypeDisplayPort"
	case KAudioDeviceTransportTypeFireWire:
		return "KAudioDeviceTransportTypeFireWire"
	case KAudioDeviceTransportTypeHDMI:
		return "KAudioDeviceTransportTypeHDMI"
	case KAudioDeviceTransportTypePCI:
		return "KAudioDeviceTransportTypePCI"
	case KAudioDeviceTransportTypeThunderbolt:
		return "KAudioDeviceTransportTypeThunderbolt"
	case KAudioDeviceTransportTypeUSB:
		return "KAudioDeviceTransportTypeUSB"
	case KAudioDeviceTransportTypeUnknown:
		return "KAudioDeviceTransportTypeUnknown"
	case KAudioDeviceTransportTypeVirtual:
		return "KAudioDeviceTransportTypeVirtual"
	default:
		return fmt.Sprintf("KAudioDeviceTransportType(%d)", e)
	}
}

type KAudioDeviceTransportTypeAuto uint

const (
	KAudioDeviceTransportTypeAutoAggregate KAudioDeviceTransportTypeAuto = 'f'<<24 | 'g'<<16 | 'r'<<8 | 'p' // 'fgrp'
)

func (e KAudioDeviceTransportTypeAuto) String() string {
	switch e {
	case KAudioDeviceTransportTypeAutoAggregate:
		return "KAudioDeviceTransportTypeAutoAggregate"
	default:
		return fmt.Sprintf("KAudioDeviceTransportTypeAuto(%d)", e)
	}
}

type KAudioEndPointClassI uint

const (
	KAudioEndPointClassID KAudioEndPointClassI = 'e'<<24 | 'n'<<16 | 'd'<<8 | 'p' // 'endp'
)

func (e KAudioEndPointClassI) String() string {
	switch e {
	case KAudioEndPointClassID:
		return "KAudioEndPointClassID"
	default:
		return fmt.Sprintf("KAudioEndPointClassI(%d)", e)
	}
}

type KAudioEndPointDeviceClassI uint

const (
	KAudioEndPointDeviceClassID KAudioEndPointDeviceClassI = 'e'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'edev'
)

func (e KAudioEndPointDeviceClassI) String() string {
	switch e {
	case KAudioEndPointDeviceClassID:
		return "KAudioEndPointDeviceClassID"
	default:
		return fmt.Sprintf("KAudioEndPointDeviceClassI(%d)", e)
	}
}

type KAudioEndPointDeviceProperty uint

const (
	KAudioEndPointDevicePropertyComposition  KAudioEndPointDeviceProperty = 'a'<<24 | 'c'<<16 | 'o'<<8 | 'm' // 'acom'
	KAudioEndPointDevicePropertyEndPointList KAudioEndPointDeviceProperty = 'a'<<24 | 'g'<<16 | 'r'<<8 | 'p' // 'agrp'
	KAudioEndPointDevicePropertyIsPrivate    KAudioEndPointDeviceProperty = 'p'<<24 | 'r'<<16 | 'i'<<8 | 'v' // 'priv'
)

func (e KAudioEndPointDeviceProperty) String() string {
	switch e {
	case KAudioEndPointDevicePropertyComposition:
		return "KAudioEndPointDevicePropertyComposition"
	case KAudioEndPointDevicePropertyEndPointList:
		return "KAudioEndPointDevicePropertyEndPointList"
	case KAudioEndPointDevicePropertyIsPrivate:
		return "KAudioEndPointDevicePropertyIsPrivate"
	default:
		return fmt.Sprintf("KAudioEndPointDeviceProperty(%d)", e)
	}
}

type KAudioHardwareProperty uint

const (
	KAudioHardwarePropertyBoxList                             KAudioHardwareProperty = 'b'<<24 | 'o'<<16 | 'x'<<8 | '#' // 'box#'
	KAudioHardwarePropertyClockDeviceList                     KAudioHardwareProperty = 'c'<<24 | 'l'<<16 | 'k'<<8 | '#' // 'clk#'
	KAudioHardwarePropertyDefaultInputDevice                  KAudioHardwareProperty = 'd'<<24 | 'I'<<16 | 'n'<<8 | ' ' // 'dIn '
	KAudioHardwarePropertyDefaultOutputDevice                 KAudioHardwareProperty = 'd'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'dOut'
	KAudioHardwarePropertyDefaultSystemOutputDevice           KAudioHardwareProperty = 's'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'sOut'
	KAudioHardwarePropertyDeviceForUID                        KAudioHardwareProperty = 'd'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'duid'
	KAudioHardwarePropertyDevices                             KAudioHardwareProperty = 'd'<<24 | 'e'<<16 | 'v'<<8 | '#' // 'dev#'
	KAudioHardwarePropertyHogModeIsAllowed                    KAudioHardwareProperty = 'h'<<24 | 'o'<<16 | 'g'<<8 | 'r' // 'hogr'
	KAudioHardwarePropertyIsInitingOrExiting                  KAudioHardwareProperty = 'i'<<24 | 'n'<<16 | 'o'<<8 | 't' // 'inot'
	KAudioHardwarePropertyMixStereoToMono                     KAudioHardwareProperty = 's'<<24 | 't'<<16 | 'm'<<8 | 'o' // 'stmo'
	KAudioHardwarePropertyPlugInForBundleID                   KAudioHardwareProperty = 'p'<<24 | 'i'<<16 | 'b'<<8 | 'i' // 'pibi'
	KAudioHardwarePropertyPlugInList                          KAudioHardwareProperty = 'p'<<24 | 'l'<<16 | 'g'<<8 | '#' // 'plg#'
	KAudioHardwarePropertyPowerHint                           KAudioHardwareProperty = 'p'<<24 | 'o'<<16 | 'w'<<8 | 'h' // 'powh'
	KAudioHardwarePropertyProcessInputMute                    KAudioHardwareProperty = 'p'<<24 | 'm'<<16 | 'i'<<8 | 'n' // 'pmin'
	KAudioHardwarePropertyProcessIsAudible                    KAudioHardwareProperty = 'p'<<24 | 'm'<<16 | 'u'<<8 | 't' // 'pmut'
	KAudioHardwarePropertyProcessIsMain                       KAudioHardwareProperty = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'main'
	KAudioHardwarePropertyProcessObjectList                   KAudioHardwareProperty = 'p'<<24 | 'r'<<16 | 's'<<8 | '#' // 'prs#'
	KAudioHardwarePropertyRunLoop                             KAudioHardwareProperty = 'r'<<24 | 'n'<<16 | 'l'<<8 | 'p' // 'rnlp'
	KAudioHardwarePropertyServiceRestarted                    KAudioHardwareProperty = 's'<<24 | 'r'<<16 | 's'<<8 | 't' // 'srst'
	KAudioHardwarePropertySleepingIsAllowed                   KAudioHardwareProperty = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KAudioHardwarePropertyTapList                             KAudioHardwareProperty = 't'<<24 | 'p'<<16 | 's'<<8 | '#' // 'tps#'
	KAudioHardwarePropertyTranslateBundleIDToPlugIn           KAudioHardwareProperty = 'b'<<24 | 'i'<<16 | 'd'<<8 | 'p' // 'bidp'
	KAudioHardwarePropertyTranslateBundleIDToTransportManager KAudioHardwareProperty = 't'<<24 | 'm'<<16 | 'b'<<8 | 'i' // 'tmbi'
	KAudioHardwarePropertyTranslatePIDToProcessObject         KAudioHardwareProperty = 'i'<<24 | 'd'<<16 | '2'<<8 | 'p' // 'id2p'
	KAudioHardwarePropertyTranslateUIDToBox                   KAudioHardwareProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'b' // 'uidb'
	KAudioHardwarePropertyTranslateUIDToClockDevice           KAudioHardwareProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'c' // 'uidc'
	KAudioHardwarePropertyTranslateUIDToDevice                KAudioHardwareProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'd' // 'uidd'
	KAudioHardwarePropertyTranslateUIDToTap                   KAudioHardwareProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 't' // 'uidt'
	KAudioHardwarePropertyTransportManagerList                KAudioHardwareProperty = 't'<<24 | 'm'<<16 | 'g'<<8 | '#' // 'tmg#'
	KAudioHardwarePropertyUnloadingIsAllowed                  KAudioHardwareProperty = 'u'<<24 | 'n'<<16 | 'l'<<8 | 'd' // 'unld'
	KAudioHardwarePropertyUserIDChanged                       KAudioHardwareProperty = 'e'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'euid'
	KAudioHardwarePropertyUserSessionIsActiveOrHeadless       KAudioHardwareProperty = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
	// Deprecated.
	KAudioHardwarePropertyProcessIsMaster KAudioHardwareProperty = 0
)

func (e KAudioHardwareProperty) String() string {
	switch e {
	case KAudioHardwarePropertyBoxList:
		return "KAudioHardwarePropertyBoxList"
	case KAudioHardwarePropertyClockDeviceList:
		return "KAudioHardwarePropertyClockDeviceList"
	case KAudioHardwarePropertyDefaultInputDevice:
		return "KAudioHardwarePropertyDefaultInputDevice"
	case KAudioHardwarePropertyDefaultOutputDevice:
		return "KAudioHardwarePropertyDefaultOutputDevice"
	case KAudioHardwarePropertyDefaultSystemOutputDevice:
		return "KAudioHardwarePropertyDefaultSystemOutputDevice"
	case KAudioHardwarePropertyDeviceForUID:
		return "KAudioHardwarePropertyDeviceForUID"
	case KAudioHardwarePropertyDevices:
		return "KAudioHardwarePropertyDevices"
	case KAudioHardwarePropertyHogModeIsAllowed:
		return "KAudioHardwarePropertyHogModeIsAllowed"
	case KAudioHardwarePropertyIsInitingOrExiting:
		return "KAudioHardwarePropertyIsInitingOrExiting"
	case KAudioHardwarePropertyMixStereoToMono:
		return "KAudioHardwarePropertyMixStereoToMono"
	case KAudioHardwarePropertyPlugInForBundleID:
		return "KAudioHardwarePropertyPlugInForBundleID"
	case KAudioHardwarePropertyPlugInList:
		return "KAudioHardwarePropertyPlugInList"
	case KAudioHardwarePropertyPowerHint:
		return "KAudioHardwarePropertyPowerHint"
	case KAudioHardwarePropertyProcessInputMute:
		return "KAudioHardwarePropertyProcessInputMute"
	case KAudioHardwarePropertyProcessIsAudible:
		return "KAudioHardwarePropertyProcessIsAudible"
	case KAudioHardwarePropertyProcessIsMain:
		return "KAudioHardwarePropertyProcessIsMain"
	case KAudioHardwarePropertyProcessObjectList:
		return "KAudioHardwarePropertyProcessObjectList"
	case KAudioHardwarePropertyRunLoop:
		return "KAudioHardwarePropertyRunLoop"
	case KAudioHardwarePropertyServiceRestarted:
		return "KAudioHardwarePropertyServiceRestarted"
	case KAudioHardwarePropertySleepingIsAllowed:
		return "KAudioHardwarePropertySleepingIsAllowed"
	case KAudioHardwarePropertyTapList:
		return "KAudioHardwarePropertyTapList"
	case KAudioHardwarePropertyTranslateBundleIDToPlugIn:
		return "KAudioHardwarePropertyTranslateBundleIDToPlugIn"
	case KAudioHardwarePropertyTranslateBundleIDToTransportManager:
		return "KAudioHardwarePropertyTranslateBundleIDToTransportManager"
	case KAudioHardwarePropertyTranslatePIDToProcessObject:
		return "KAudioHardwarePropertyTranslatePIDToProcessObject"
	case KAudioHardwarePropertyTranslateUIDToBox:
		return "KAudioHardwarePropertyTranslateUIDToBox"
	case KAudioHardwarePropertyTranslateUIDToClockDevice:
		return "KAudioHardwarePropertyTranslateUIDToClockDevice"
	case KAudioHardwarePropertyTranslateUIDToDevice:
		return "KAudioHardwarePropertyTranslateUIDToDevice"
	case KAudioHardwarePropertyTranslateUIDToTap:
		return "KAudioHardwarePropertyTranslateUIDToTap"
	case KAudioHardwarePropertyTransportManagerList:
		return "KAudioHardwarePropertyTransportManagerList"
	case KAudioHardwarePropertyUnloadingIsAllowed:
		return "KAudioHardwarePropertyUnloadingIsAllowed"
	case KAudioHardwarePropertyUserIDChanged:
		return "KAudioHardwarePropertyUserIDChanged"
	case KAudioHardwarePropertyUserSessionIsActiveOrHeadless:
		return "KAudioHardwarePropertyUserSessionIsActiveOrHeadless"
	case KAudioHardwarePropertyProcessIsMaster:
		return "KAudioHardwarePropertyProcessIsMaster"
	default:
		return fmt.Sprintf("KAudioHardwareProperty(%d)", e)
	}
}

type KAudioHardwarePropertyBootChimeVolume uint

const (
	KAudioHardwarePropertyBootChimeVolumeDecibels                         KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'b'<<16 | 'v'<<8 | 'd' // 'bbvd'
	KAudioHardwarePropertyBootChimeVolumeDecibelsToScalar                 KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'd'<<16 | '2'<<8 | 'v' // 'bd2v'
	KAudioHardwarePropertyBootChimeVolumeDecibelsToScalarTransferFunction KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'v'<<16 | 't'<<8 | 'f' // 'bvtf'
	KAudioHardwarePropertyBootChimeVolumeRangeDecibels                    KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'b'<<16 | 'd'<<8 | '#' // 'bbd#'
	KAudioHardwarePropertyBootChimeVolumeScalar                           KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'b'<<16 | 'v'<<8 | 's' // 'bbvs'
	KAudioHardwarePropertyBootChimeVolumeScalarToDecibels                 KAudioHardwarePropertyBootChimeVolume = 'b'<<24 | 'v'<<16 | '2'<<8 | 'd' // 'bv2d'
)

func (e KAudioHardwarePropertyBootChimeVolume) String() string {
	switch e {
	case KAudioHardwarePropertyBootChimeVolumeDecibels:
		return "KAudioHardwarePropertyBootChimeVolumeDecibels"
	case KAudioHardwarePropertyBootChimeVolumeDecibelsToScalar:
		return "KAudioHardwarePropertyBootChimeVolumeDecibelsToScalar"
	case KAudioHardwarePropertyBootChimeVolumeDecibelsToScalarTransferFunction:
		return "KAudioHardwarePropertyBootChimeVolumeDecibelsToScalarTransferFunction"
	case KAudioHardwarePropertyBootChimeVolumeRangeDecibels:
		return "KAudioHardwarePropertyBootChimeVolumeRangeDecibels"
	case KAudioHardwarePropertyBootChimeVolumeScalar:
		return "KAudioHardwarePropertyBootChimeVolumeScalar"
	case KAudioHardwarePropertyBootChimeVolumeScalarToDecibels:
		return "KAudioHardwarePropertyBootChimeVolumeScalarToDecibels"
	default:
		return fmt.Sprintf("KAudioHardwarePropertyBootChimeVolume(%d)", e)
	}
}

type KAudioISubOwnerControlClassI uint

const (
	KAudioISubOwnerControlClassID KAudioISubOwnerControlClassI = 'a'<<24 | 't'<<16 | 'c'<<8 | 'h' // 'atch'
)

func (e KAudioISubOwnerControlClassI) String() string {
	switch e {
	case KAudioISubOwnerControlClassID:
		return "KAudioISubOwnerControlClassID"
	default:
		return fmt.Sprintf("KAudioISubOwnerControlClassI(%d)", e)
	}
}

type KAudioLevelControlProperty uint

const (
	KAudioLevelControlPropertyConvertDecibelsToScalar KAudioLevelControlProperty = 'l'<<24 | 'c'<<16 | 'd'<<8 | 's' // 'lcds'
	KAudioLevelControlPropertyConvertScalarToDecibels KAudioLevelControlProperty = 'l'<<24 | 'c'<<16 | 's'<<8 | 'd' // 'lcsd'
	KAudioLevelControlPropertyDecibelRange            KAudioLevelControlProperty = 'l'<<24 | 'c'<<16 | 'd'<<8 | 'r' // 'lcdr'
	KAudioLevelControlPropertyDecibelValue            KAudioLevelControlProperty = 'l'<<24 | 'c'<<16 | 'd'<<8 | 'v' // 'lcdv'
	KAudioLevelControlPropertyScalarValue             KAudioLevelControlProperty = 'l'<<24 | 'c'<<16 | 's'<<8 | 'v' // 'lcsv'
)

func (e KAudioLevelControlProperty) String() string {
	switch e {
	case KAudioLevelControlPropertyConvertDecibelsToScalar:
		return "KAudioLevelControlPropertyConvertDecibelsToScalar"
	case KAudioLevelControlPropertyConvertScalarToDecibels:
		return "KAudioLevelControlPropertyConvertScalarToDecibels"
	case KAudioLevelControlPropertyDecibelRange:
		return "KAudioLevelControlPropertyDecibelRange"
	case KAudioLevelControlPropertyDecibelValue:
		return "KAudioLevelControlPropertyDecibelValue"
	case KAudioLevelControlPropertyScalarValue:
		return "KAudioLevelControlPropertyScalarValue"
	default:
		return fmt.Sprintf("KAudioLevelControlProperty(%d)", e)
	}
}

type KAudioLevelControlPropertyDecibelsToScalarTransfer uint

const (
	KAudioLevelControlPropertyDecibelsToScalarTransferFunction KAudioLevelControlPropertyDecibelsToScalarTransfer = 'l'<<24 | 'c'<<16 | 't'<<8 | 'f' // 'lctf'
)

func (e KAudioLevelControlPropertyDecibelsToScalarTransfer) String() string {
	switch e {
	case KAudioLevelControlPropertyDecibelsToScalarTransferFunction:
		return "KAudioLevelControlPropertyDecibelsToScalarTransferFunction"
	default:
		return fmt.Sprintf("KAudioLevelControlPropertyDecibelsToScalarTransfer(%d)", e)
	}
}

type KAudioObject uint

const (
	KAudioObjectUnknown KAudioObject = 0
)

func (e KAudioObject) String() string {
	switch e {
	case KAudioObjectUnknown:
		return "KAudioObjectUnknown"
	default:
		return fmt.Sprintf("KAudioObject(%d)", e)
	}
}

type KAudioObjectClassI uint

const (
	KAudioObjectClassIDValue KAudioObjectClassI = 'a'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'aobj'
)

func (e KAudioObjectClassI) String() string {
	switch e {
	case KAudioObjectClassIDValue:
		return "KAudioObjectClassIDValue"
	default:
		return fmt.Sprintf("KAudioObjectClassI(%d)", e)
	}
}

type KAudioObjectClassID uint

const (
	KAudioObjectClassIDWildcard KAudioObjectClassID = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KAudioObjectClassID) String() string {
	switch e {
	case KAudioObjectClassIDWildcard:
		return "KAudioObjectClassIDWildcard"
	default:
		return fmt.Sprintf("KAudioObjectClassID(%d)", e)
	}
}

type KAudioObjectPlugIn uint

const (
	KAudioObjectPlugInObject KAudioObjectPlugIn = 0
)

func (e KAudioObjectPlugIn) String() string {
	switch e {
	case KAudioObjectPlugInObject:
		return "KAudioObjectPlugInObject"
	default:
		return fmt.Sprintf("KAudioObjectPlugIn(%d)", e)
	}
}

type KAudioObjectProperty uint

const (
	KAudioObjectPropertyBaseClass           KAudioObjectProperty = 'b'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'bcls'
	KAudioObjectPropertyClass               KAudioObjectProperty = 'c'<<24 | 'l'<<16 | 'a'<<8 | 's' // 'clas'
	KAudioObjectPropertyCreator             KAudioObjectProperty = 'o'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'oplg'
	KAudioObjectPropertyElementCategoryName KAudioObjectProperty = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KAudioObjectPropertyElementMain         KAudioObjectProperty = 0
	KAudioObjectPropertyElementName         KAudioObjectProperty = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KAudioObjectPropertyElementNumberName   KAudioObjectProperty = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KAudioObjectPropertyFirmwareVersion     KAudioObjectProperty = 'f'<<24 | 'w'<<16 | 'v'<<8 | 'n' // 'fwvn'
	KAudioObjectPropertyIdentify            KAudioObjectProperty = 'i'<<24 | 'd'<<16 | 'e'<<8 | 'n' // 'iden'
	KAudioObjectPropertyListenerAdded       KAudioObjectProperty = 'l'<<24 | 'i'<<16 | 's'<<8 | 'a' // 'lisa'
	KAudioObjectPropertyListenerRemoved     KAudioObjectProperty = 'l'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'lisr'
	KAudioObjectPropertyManufacturer        KAudioObjectProperty = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KAudioObjectPropertyModelName           KAudioObjectProperty = 'l'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'lmod'
	KAudioObjectPropertyName                KAudioObjectProperty = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KAudioObjectPropertyOwnedObjects        KAudioObjectProperty = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KAudioObjectPropertyOwner               KAudioObjectProperty = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
	KAudioObjectPropertyScopeGlobal         KAudioObjectProperty = 'g'<<24 | 'l'<<16 | 'o'<<8 | 'b' // 'glob'
	KAudioObjectPropertyScopeInput          KAudioObjectProperty = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KAudioObjectPropertyScopeOutput         KAudioObjectProperty = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KAudioObjectPropertyScopePlayThrough    KAudioObjectProperty = 'p'<<24 | 't'<<16 | 'r'<<8 | 'u' // 'ptru'
	KAudioObjectPropertySerialNumber        KAudioObjectProperty = 's'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'snum'
	// Deprecated.
	KAudioObjectPropertyElementMaster KAudioObjectProperty = 0
)

func (e KAudioObjectProperty) String() string {
	switch e {
	case KAudioObjectPropertyBaseClass:
		return "KAudioObjectPropertyBaseClass"
	case KAudioObjectPropertyClass:
		return "KAudioObjectPropertyClass"
	case KAudioObjectPropertyCreator:
		return "KAudioObjectPropertyCreator"
	case KAudioObjectPropertyElementCategoryName:
		return "KAudioObjectPropertyElementCategoryName"
	case KAudioObjectPropertyElementMain:
		return "KAudioObjectPropertyElementMain"
	case KAudioObjectPropertyElementName:
		return "KAudioObjectPropertyElementName"
	case KAudioObjectPropertyElementNumberName:
		return "KAudioObjectPropertyElementNumberName"
	case KAudioObjectPropertyFirmwareVersion:
		return "KAudioObjectPropertyFirmwareVersion"
	case KAudioObjectPropertyIdentify:
		return "KAudioObjectPropertyIdentify"
	case KAudioObjectPropertyListenerAdded:
		return "KAudioObjectPropertyListenerAdded"
	case KAudioObjectPropertyListenerRemoved:
		return "KAudioObjectPropertyListenerRemoved"
	case KAudioObjectPropertyManufacturer:
		return "KAudioObjectPropertyManufacturer"
	case KAudioObjectPropertyModelName:
		return "KAudioObjectPropertyModelName"
	case KAudioObjectPropertyName:
		return "KAudioObjectPropertyName"
	case KAudioObjectPropertyOwnedObjects:
		return "KAudioObjectPropertyOwnedObjects"
	case KAudioObjectPropertyOwner:
		return "KAudioObjectPropertyOwner"
	case KAudioObjectPropertyScopeGlobal:
		return "KAudioObjectPropertyScopeGlobal"
	case KAudioObjectPropertyScopeInput:
		return "KAudioObjectPropertyScopeInput"
	case KAudioObjectPropertyScopeOutput:
		return "KAudioObjectPropertyScopeOutput"
	case KAudioObjectPropertyScopePlayThrough:
		return "KAudioObjectPropertyScopePlayThrough"
	case KAudioObjectPropertySerialNumber:
		return "KAudioObjectPropertySerialNumber"
	default:
		return fmt.Sprintf("KAudioObjectProperty(%d)", e)
	}
}

type KAudioObjectPropertyCustomPropertyInfo uint

const (
	KAudioObjectPropertyCustomPropertyInfoList KAudioObjectPropertyCustomPropertyInfo = 0
)

func (e KAudioObjectPropertyCustomPropertyInfo) String() string {
	switch e {
	case KAudioObjectPropertyCustomPropertyInfoList:
		return "KAudioObjectPropertyCustomPropertyInfoList"
	default:
		return fmt.Sprintf("KAudioObjectPropertyCustomPropertyInfo(%d)", e)
	}
}

type KAudioObjectPropertyElement uint

const (
	KAudioObjectPropertyElementWildcard KAudioObjectPropertyElement = 0
)

func (e KAudioObjectPropertyElement) String() string {
	switch e {
	case KAudioObjectPropertyElementWildcard:
		return "KAudioObjectPropertyElementWildcard"
	default:
		return fmt.Sprintf("KAudioObjectPropertyElement(%d)", e)
	}
}

type KAudioObjectPropertyScope uint

const (
	KAudioObjectPropertyScopeWildcard KAudioObjectPropertyScope = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KAudioObjectPropertyScope) String() string {
	switch e {
	case KAudioObjectPropertyScopeWildcard:
		return "KAudioObjectPropertyScopeWildcard"
	default:
		return fmt.Sprintf("KAudioObjectPropertyScope(%d)", e)
	}
}

type KAudioObjectPropertySelector uint

const (
	KAudioObjectPropertySelectorWildcard KAudioObjectPropertySelector = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KAudioObjectPropertySelector) String() string {
	switch e {
	case KAudioObjectPropertySelectorWildcard:
		return "KAudioObjectPropertySelectorWildcard"
	default:
		return fmt.Sprintf("KAudioObjectPropertySelector(%d)", e)
	}
}

type KAudioObjectSystem uint

const (
	KAudioObjectSystemObject KAudioObjectSystem = 0
)

func (e KAudioObjectSystem) String() string {
	switch e {
	case KAudioObjectSystemObject:
		return "KAudioObjectSystemObject"
	default:
		return fmt.Sprintf("KAudioObjectSystem(%d)", e)
	}
}

type KAudioPlugIn uint

const (
	KAudioPlugInCreateAggregateDevice  KAudioPlugIn = 'c'<<24 | 'a'<<16 | 'g'<<8 | 'g' // 'cagg'
	KAudioPlugInDestroyAggregateDevice KAudioPlugIn = 'd'<<24 | 'a'<<16 | 'g'<<8 | 'g' // 'dagg'
)

func (e KAudioPlugIn) String() string {
	switch e {
	case KAudioPlugInCreateAggregateDevice:
		return "KAudioPlugInCreateAggregateDevice"
	case KAudioPlugInDestroyAggregateDevice:
		return "KAudioPlugInDestroyAggregateDevice"
	default:
		return fmt.Sprintf("KAudioPlugIn(%d)", e)
	}
}

type KAudioPlugInClassI uint

const (
	KAudioPlugInClassID KAudioPlugInClassI = 'a'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'aplg'
)

func (e KAudioPlugInClassI) String() string {
	switch e {
	case KAudioPlugInClassID:
		return "KAudioPlugInClassID"
	default:
		return fmt.Sprintf("KAudioPlugInClassI(%d)", e)
	}
}

type KAudioPlugInProperty uint

const (
	KAudioPlugInPropertyBoxList                   KAudioPlugInProperty = 'b'<<24 | 'o'<<16 | 'x'<<8 | '#' // 'box#'
	KAudioPlugInPropertyBundleID                  KAudioPlugInProperty = 'p'<<24 | 'i'<<16 | 'i'<<8 | 'd' // 'piid'
	KAudioPlugInPropertyClockDeviceList           KAudioPlugInProperty = 'c'<<24 | 'l'<<16 | 'k'<<8 | '#' // 'clk#'
	KAudioPlugInPropertyDeviceList                KAudioPlugInProperty = 'd'<<24 | 'e'<<16 | 'v'<<8 | '#' // 'dev#'
	KAudioPlugInPropertyTranslateUIDToBox         KAudioPlugInProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'b' // 'uidb'
	KAudioPlugInPropertyTranslateUIDToClockDevice KAudioPlugInProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'c' // 'uidc'
	KAudioPlugInPropertyTranslateUIDToDevice      KAudioPlugInProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'd' // 'uidd'
)

func (e KAudioPlugInProperty) String() string {
	switch e {
	case KAudioPlugInPropertyBoxList:
		return "KAudioPlugInPropertyBoxList"
	case KAudioPlugInPropertyBundleID:
		return "KAudioPlugInPropertyBundleID"
	case KAudioPlugInPropertyClockDeviceList:
		return "KAudioPlugInPropertyClockDeviceList"
	case KAudioPlugInPropertyDeviceList:
		return "KAudioPlugInPropertyDeviceList"
	case KAudioPlugInPropertyTranslateUIDToBox:
		return "KAudioPlugInPropertyTranslateUIDToBox"
	case KAudioPlugInPropertyTranslateUIDToClockDevice:
		return "KAudioPlugInPropertyTranslateUIDToClockDevice"
	case KAudioPlugInPropertyTranslateUIDToDevice:
		return "KAudioPlugInPropertyTranslateUIDToDevice"
	default:
		return fmt.Sprintf("KAudioPlugInProperty(%d)", e)
	}
}

type KAudioPlugInPropertyResource uint

const (
	KAudioPlugInPropertyResourceBundle KAudioPlugInPropertyResource = 0
)

func (e KAudioPlugInPropertyResource) String() string {
	switch e {
	case KAudioPlugInPropertyResourceBundle:
		return "KAudioPlugInPropertyResourceBundle"
	default:
		return fmt.Sprintf("KAudioPlugInPropertyResource(%d)", e)
	}
}

type KAudioProcessClassI uint

const (
	KAudioProcessClassID KAudioProcessClassI = 'c'<<24 | 'l'<<16 | 'n'<<8 | 't' // 'clnt'
)

func (e KAudioProcessClassI) String() string {
	switch e {
	case KAudioProcessClassID:
		return "KAudioProcessClassID"
	default:
		return fmt.Sprintf("KAudioProcessClassI(%d)", e)
	}
}

type KAudioProcessProperty uint

const (
	KAudioProcessPropertyBundleID        KAudioProcessProperty = 'p'<<24 | 'b'<<16 | 'i'<<8 | 'd' // 'pbid'
	KAudioProcessPropertyDevices         KAudioProcessProperty = 'p'<<24 | 'd'<<16 | 'v'<<8 | '#' // 'pdv#'
	KAudioProcessPropertyIsRunning       KAudioProcessProperty = 'p'<<24 | 'i'<<16 | 'r'<<8 | '?' // 'pir?'
	KAudioProcessPropertyIsRunningInput  KAudioProcessProperty = 'p'<<24 | 'i'<<16 | 'r'<<8 | 'i' // 'piri'
	KAudioProcessPropertyIsRunningOutput KAudioProcessProperty = 'p'<<24 | 'i'<<16 | 'r'<<8 | 'o' // 'piro'
	KAudioProcessPropertyPID             KAudioProcessProperty = 'p'<<24 | 'p'<<16 | 'i'<<8 | 'd' // 'ppid'
)

func (e KAudioProcessProperty) String() string {
	switch e {
	case KAudioProcessPropertyBundleID:
		return "KAudioProcessPropertyBundleID"
	case KAudioProcessPropertyDevices:
		return "KAudioProcessPropertyDevices"
	case KAudioProcessPropertyIsRunning:
		return "KAudioProcessPropertyIsRunning"
	case KAudioProcessPropertyIsRunningInput:
		return "KAudioProcessPropertyIsRunningInput"
	case KAudioProcessPropertyIsRunningOutput:
		return "KAudioProcessPropertyIsRunningOutput"
	case KAudioProcessPropertyPID:
		return "KAudioProcessPropertyPID"
	default:
		return fmt.Sprintf("KAudioProcessProperty(%d)", e)
	}
}

type KAudioPropertyWildcard uint

const (
	KAudioPropertyWildcardChannel KAudioPropertyWildcard = 0
	KAudioPropertyWildcardSection KAudioPropertyWildcard = 0
)

func (e KAudioPropertyWildcard) String() string {
	switch e {
	case KAudioPropertyWildcardChannel:
		return "KAudioPropertyWildcardChannel"
	default:
		return fmt.Sprintf("KAudioPropertyWildcard(%d)", e)
	}
}

type KAudioPropertyWildcardPropertyI uint

const (
	KAudioPropertyWildcardPropertyID KAudioPropertyWildcardPropertyI = 0
)

func (e KAudioPropertyWildcardPropertyI) String() string {
	switch e {
	case KAudioPropertyWildcardPropertyID:
		return "KAudioPropertyWildcardPropertyID"
	default:
		return fmt.Sprintf("KAudioPropertyWildcardPropertyI(%d)", e)
	}
}

type KAudioSelectorControlItemKind uint

const (
	KAudioSelectorControlItemKindSpacer KAudioSelectorControlItemKind = 's'<<24 | 'p'<<16 | 'c'<<8 | 'r' // 'spcr'
)

func (e KAudioSelectorControlItemKind) String() string {
	switch e {
	case KAudioSelectorControlItemKindSpacer:
		return "KAudioSelectorControlItemKindSpacer"
	default:
		return fmt.Sprintf("KAudioSelectorControlItemKind(%d)", e)
	}
}

type KAudioSelectorControlProperty uint

const (
	KAudioSelectorControlPropertyAvailableItems KAudioSelectorControlProperty = 's'<<24 | 'c'<<16 | 'a'<<8 | 'i' // 'scai'
	KAudioSelectorControlPropertyCurrentItem    KAudioSelectorControlProperty = 's'<<24 | 'c'<<16 | 'c'<<8 | 'i' // 'scci'
	KAudioSelectorControlPropertyItemKind       KAudioSelectorControlProperty = 'c'<<24 | 'l'<<16 | 'k'<<8 | 'k' // 'clkk'
	KAudioSelectorControlPropertyItemName       KAudioSelectorControlProperty = 's'<<24 | 'c'<<16 | 'i'<<8 | 'n' // 'scin'
)

func (e KAudioSelectorControlProperty) String() string {
	switch e {
	case KAudioSelectorControlPropertyAvailableItems:
		return "KAudioSelectorControlPropertyAvailableItems"
	case KAudioSelectorControlPropertyCurrentItem:
		return "KAudioSelectorControlPropertyCurrentItem"
	case KAudioSelectorControlPropertyItemKind:
		return "KAudioSelectorControlPropertyItemKind"
	case KAudioSelectorControlPropertyItemName:
		return "KAudioSelectorControlPropertyItemName"
	default:
		return fmt.Sprintf("KAudioSelectorControlProperty(%d)", e)
	}
}

type KAudioServerPlugInCustomPropertyDataType uint

const (
	KAudioServerPlugInCustomPropertyDataTypeCFPropertyList KAudioServerPlugInCustomPropertyDataType = 0
	KAudioServerPlugInCustomPropertyDataTypeCFString       KAudioServerPlugInCustomPropertyDataType = 0
	KAudioServerPlugInCustomPropertyDataTypeNone           KAudioServerPlugInCustomPropertyDataType = 0
)

func (e KAudioServerPlugInCustomPropertyDataType) String() string {
	switch e {
	case KAudioServerPlugInCustomPropertyDataTypeCFPropertyList:
		return "KAudioServerPlugInCustomPropertyDataTypeCFPropertyList"
	default:
		return fmt.Sprintf("KAudioServerPlugInCustomPropertyDataType(%d)", e)
	}
}

type KAudioServerPlugInHostClientI uint

const (
	KAudioServerPlugInHostClientID KAudioServerPlugInHostClientI = 0
)

func (e KAudioServerPlugInHostClientI) String() string {
	switch e {
	case KAudioServerPlugInHostClientID:
		return "KAudioServerPlugInHostClientID"
	default:
		return fmt.Sprintf("KAudioServerPlugInHostClientI(%d)", e)
	}
}

type KAudioSliderControlClassI uint

const (
	KAudioSliderControlClassID KAudioSliderControlClassI = 's'<<24 | 'l'<<16 | 'd'<<8 | 'r' // 'sldr'
)

func (e KAudioSliderControlClassI) String() string {
	switch e {
	case KAudioSliderControlClassID:
		return "KAudioSliderControlClassID"
	default:
		return fmt.Sprintf("KAudioSliderControlClassI(%d)", e)
	}
}

type KAudioSliderControlProperty uint

const (
	KAudioSliderControlPropertyRange KAudioSliderControlProperty = 's'<<24 | 'd'<<16 | 'r'<<8 | 'r' // 'sdrr'
	KAudioSliderControlPropertyValue KAudioSliderControlProperty = 's'<<24 | 'd'<<16 | 'r'<<8 | 'v' // 'sdrv'
)

func (e KAudioSliderControlProperty) String() string {
	switch e {
	case KAudioSliderControlPropertyRange:
		return "KAudioSliderControlPropertyRange"
	case KAudioSliderControlPropertyValue:
		return "KAudioSliderControlPropertyValue"
	default:
		return fmt.Sprintf("KAudioSliderControlProperty(%d)", e)
	}
}

type KAudioStereoPanControlClassI uint

const (
	KAudioStereoPanControlClassID KAudioStereoPanControlClassI = 's'<<24 | 'p'<<16 | 'a'<<8 | 'n' // 'span'
)

func (e KAudioStereoPanControlClassI) String() string {
	switch e {
	case KAudioStereoPanControlClassID:
		return "KAudioStereoPanControlClassID"
	default:
		return fmt.Sprintf("KAudioStereoPanControlClassI(%d)", e)
	}
}

type KAudioStereoPanControlProperty uint

const (
	KAudioStereoPanControlPropertyPanningChannels KAudioStereoPanControlProperty = 's'<<24 | 'p'<<16 | 'c'<<8 | 'c' // 'spcc'
	KAudioStereoPanControlPropertyValue           KAudioStereoPanControlProperty = 's'<<24 | 'p'<<16 | 'c'<<8 | 'v' // 'spcv'
)

func (e KAudioStereoPanControlProperty) String() string {
	switch e {
	case KAudioStereoPanControlPropertyPanningChannels:
		return "KAudioStereoPanControlPropertyPanningChannels"
	case KAudioStereoPanControlPropertyValue:
		return "KAudioStereoPanControlPropertyValue"
	default:
		return fmt.Sprintf("KAudioStereoPanControlProperty(%d)", e)
	}
}

type KAudioStream uint

const (
	KAudioStreamUnknown KAudioStream = 0
)

func (e KAudioStream) String() string {
	switch e {
	case KAudioStreamUnknown:
		return "KAudioStreamUnknown"
	default:
		return fmt.Sprintf("KAudioStream(%d)", e)
	}
}

type KAudioStreamClassI uint

const (
	KAudioStreamClassID KAudioStreamClassI = 'a'<<24 | 's'<<16 | 't'<<8 | 'r' // 'astr'
)

func (e KAudioStreamClassI) String() string {
	switch e {
	case KAudioStreamClassID:
		return "KAudioStreamClassID"
	default:
		return fmt.Sprintf("KAudioStreamClassI(%d)", e)
	}
}

type KAudioStreamProperty uint

const (
	KAudioStreamPropertyAvailablePhysicalFormats KAudioStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | 'a' // 'pfta'
	KAudioStreamPropertyAvailableVirtualFormats  KAudioStreamProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | 'a' // 'sfma'
	KAudioStreamPropertyDirection                KAudioStreamProperty = 's'<<24 | 'd'<<16 | 'i'<<8 | 'r' // 'sdir'
	KAudioStreamPropertyIsActive                 KAudioStreamProperty = 's'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'sact'
	KAudioStreamPropertyLatency                  KAudioStreamProperty = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KAudioStreamPropertyOwningDevice             KAudioStreamProperty = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
	KAudioStreamPropertyPhysicalFormat           KAudioStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | ' ' // 'pft '
	KAudioStreamPropertyPhysicalFormatMatch      KAudioStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | 'm' // 'pftm'
	KAudioStreamPropertyPhysicalFormatSupported  KAudioStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | '?' // 'pft?'
	KAudioStreamPropertyPhysicalFormats          KAudioStreamProperty = 'p'<<24 | 'f'<<16 | 't'<<8 | '#' // 'pft#'
	KAudioStreamPropertyStartingChannel          KAudioStreamProperty = 's'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'schn'
	KAudioStreamPropertyTerminalType             KAudioStreamProperty = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
	KAudioStreamPropertyVirtualFormat            KAudioStreamProperty = 's'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'sfmt'
)

func (e KAudioStreamProperty) String() string {
	switch e {
	case KAudioStreamPropertyAvailablePhysicalFormats:
		return "KAudioStreamPropertyAvailablePhysicalFormats"
	case KAudioStreamPropertyAvailableVirtualFormats:
		return "KAudioStreamPropertyAvailableVirtualFormats"
	case KAudioStreamPropertyDirection:
		return "KAudioStreamPropertyDirection"
	case KAudioStreamPropertyIsActive:
		return "KAudioStreamPropertyIsActive"
	case KAudioStreamPropertyLatency:
		return "KAudioStreamPropertyLatency"
	case KAudioStreamPropertyOwningDevice:
		return "KAudioStreamPropertyOwningDevice"
	case KAudioStreamPropertyPhysicalFormat:
		return "KAudioStreamPropertyPhysicalFormat"
	case KAudioStreamPropertyPhysicalFormatMatch:
		return "KAudioStreamPropertyPhysicalFormatMatch"
	case KAudioStreamPropertyPhysicalFormatSupported:
		return "KAudioStreamPropertyPhysicalFormatSupported"
	case KAudioStreamPropertyPhysicalFormats:
		return "KAudioStreamPropertyPhysicalFormats"
	case KAudioStreamPropertyStartingChannel:
		return "KAudioStreamPropertyStartingChannel"
	case KAudioStreamPropertyTerminalType:
		return "KAudioStreamPropertyTerminalType"
	case KAudioStreamPropertyVirtualFormat:
		return "KAudioStreamPropertyVirtualFormat"
	default:
		return fmt.Sprintf("KAudioStreamProperty(%d)", e)
	}
}

type KAudioStreamTerminalType uint

const (
	KAudioStreamTerminalTypeDigitalAudioInterface KAudioStreamTerminalType = 's'<<24 | 'p'<<16 | 'd'<<8 | 'f' // 'spdf'
	KAudioStreamTerminalTypeDisplayPort           KAudioStreamTerminalType = 'd'<<24 | 'p'<<16 | 'r'<<8 | 't' // 'dprt'
	KAudioStreamTerminalTypeHDMI                  KAudioStreamTerminalType = 'h'<<24 | 'd'<<16 | 'm'<<8 | 'i' // 'hdmi'
	KAudioStreamTerminalTypeHeadphones            KAudioStreamTerminalType = 'h'<<24 | 'd'<<16 | 'p'<<8 | 'h' // 'hdph'
	KAudioStreamTerminalTypeHeadsetMicrophone     KAudioStreamTerminalType = 'h'<<24 | 'm'<<16 | 'i'<<8 | 'c' // 'hmic'
	KAudioStreamTerminalTypeLFESpeaker            KAudioStreamTerminalType = 'l'<<24 | 'f'<<16 | 'e'<<8 | 's' // 'lfes'
	KAudioStreamTerminalTypeLine                  KAudioStreamTerminalType = 'l'<<24 | 'i'<<16 | 'n'<<8 | 'e' // 'line'
	KAudioStreamTerminalTypeMicrophone            KAudioStreamTerminalType = 'm'<<24 | 'i'<<16 | 'c'<<8 | 'r' // 'micr'
	KAudioStreamTerminalTypeReceiverMicrophone    KAudioStreamTerminalType = 'r'<<24 | 'm'<<16 | 'i'<<8 | 'c' // 'rmic'
	KAudioStreamTerminalTypeReceiverSpeaker       KAudioStreamTerminalType = 'r'<<24 | 's'<<16 | 'p'<<8 | 'k' // 'rspk'
	KAudioStreamTerminalTypeSpeaker               KAudioStreamTerminalType = 's'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'spkr'
	KAudioStreamTerminalTypeTTY                   KAudioStreamTerminalType = 't'<<24 | 't'<<16 | 'y'<<8 | '_' // 'tty_'
	KAudioStreamTerminalTypeUnknown               KAudioStreamTerminalType = 0
)

func (e KAudioStreamTerminalType) String() string {
	switch e {
	case KAudioStreamTerminalTypeDigitalAudioInterface:
		return "KAudioStreamTerminalTypeDigitalAudioInterface"
	case KAudioStreamTerminalTypeDisplayPort:
		return "KAudioStreamTerminalTypeDisplayPort"
	case KAudioStreamTerminalTypeHDMI:
		return "KAudioStreamTerminalTypeHDMI"
	case KAudioStreamTerminalTypeHeadphones:
		return "KAudioStreamTerminalTypeHeadphones"
	case KAudioStreamTerminalTypeHeadsetMicrophone:
		return "KAudioStreamTerminalTypeHeadsetMicrophone"
	case KAudioStreamTerminalTypeLFESpeaker:
		return "KAudioStreamTerminalTypeLFESpeaker"
	case KAudioStreamTerminalTypeLine:
		return "KAudioStreamTerminalTypeLine"
	case KAudioStreamTerminalTypeMicrophone:
		return "KAudioStreamTerminalTypeMicrophone"
	case KAudioStreamTerminalTypeReceiverMicrophone:
		return "KAudioStreamTerminalTypeReceiverMicrophone"
	case KAudioStreamTerminalTypeReceiverSpeaker:
		return "KAudioStreamTerminalTypeReceiverSpeaker"
	case KAudioStreamTerminalTypeSpeaker:
		return "KAudioStreamTerminalTypeSpeaker"
	case KAudioStreamTerminalTypeTTY:
		return "KAudioStreamTerminalTypeTTY"
	case KAudioStreamTerminalTypeUnknown:
		return "KAudioStreamTerminalTypeUnknown"
	default:
		return fmt.Sprintf("KAudioStreamTerminalType(%d)", e)
	}
}

type KAudioSubDeviceClassI uint

const (
	KAudioSubDeviceClassID KAudioSubDeviceClassI = 'a'<<24 | 's'<<16 | 'u'<<8 | 'b' // 'asub'
)

func (e KAudioSubDeviceClassI) String() string {
	switch e {
	case KAudioSubDeviceClassID:
		return "KAudioSubDeviceClassID"
	default:
		return fmt.Sprintf("KAudioSubDeviceClassI(%d)", e)
	}
}

type KAudioSubDeviceDriftCompensation uint

const (
	// Deprecated.
	KAudioSubDeviceDriftCompensationHighQuality KAudioSubDeviceDriftCompensation = 0
	// Deprecated.
	KAudioSubDeviceDriftCompensationLowQuality KAudioSubDeviceDriftCompensation = 0
	// Deprecated.
	KAudioSubDeviceDriftCompensationMaxQuality KAudioSubDeviceDriftCompensation = 0
	// Deprecated.
	KAudioSubDeviceDriftCompensationMediumQuality KAudioSubDeviceDriftCompensation = 0
	// Deprecated.
	KAudioSubDeviceDriftCompensationMinQuality KAudioSubDeviceDriftCompensation = 0
)

func (e KAudioSubDeviceDriftCompensation) String() string {
	switch e {
	case KAudioSubDeviceDriftCompensationHighQuality:
		return "KAudioSubDeviceDriftCompensationHighQuality"
	default:
		return fmt.Sprintf("KAudioSubDeviceDriftCompensation(%d)", e)
	}
}

type KAudioSubDeviceProperty uint

const (
	KAudioSubDevicePropertyDriftCompensation        KAudioSubDeviceProperty = 'd'<<24 | 'r'<<16 | 'f'<<8 | 't' // 'drft'
	KAudioSubDevicePropertyDriftCompensationQuality KAudioSubDeviceProperty = 'd'<<24 | 'r'<<16 | 'f'<<8 | 'q' // 'drfq'
	KAudioSubDevicePropertyExtraLatency             KAudioSubDeviceProperty = 'x'<<24 | 'l'<<16 | 't'<<8 | 'c' // 'xltc'
)

func (e KAudioSubDeviceProperty) String() string {
	switch e {
	case KAudioSubDevicePropertyDriftCompensation:
		return "KAudioSubDevicePropertyDriftCompensation"
	case KAudioSubDevicePropertyDriftCompensationQuality:
		return "KAudioSubDevicePropertyDriftCompensationQuality"
	case KAudioSubDevicePropertyExtraLatency:
		return "KAudioSubDevicePropertyExtraLatency"
	default:
		return fmt.Sprintf("KAudioSubDeviceProperty(%d)", e)
	}
}

type KAudioSubTapClassI uint

const (
	KAudioSubTapClassID KAudioSubTapClassI = 's'<<24 | 't'<<16 | 'a'<<8 | 'p' // 'stap'
)

func (e KAudioSubTapClassI) String() string {
	switch e {
	case KAudioSubTapClassID:
		return "KAudioSubTapClassID"
	default:
		return fmt.Sprintf("KAudioSubTapClassI(%d)", e)
	}
}

type KAudioSubTapProperty uint

const (
	KAudioSubTapPropertyDriftCompensation        KAudioSubTapProperty = 'd'<<24 | 'r'<<16 | 'f'<<8 | 't' // 'drft'
	KAudioSubTapPropertyDriftCompensationQuality KAudioSubTapProperty = 'd'<<24 | 'r'<<16 | 'f'<<8 | 'q' // 'drfq'
	KAudioSubTapPropertyExtraLatency             KAudioSubTapProperty = 'x'<<24 | 'l'<<16 | 't'<<8 | 'c' // 'xltc'
)

func (e KAudioSubTapProperty) String() string {
	switch e {
	case KAudioSubTapPropertyDriftCompensation:
		return "KAudioSubTapPropertyDriftCompensation"
	case KAudioSubTapPropertyDriftCompensationQuality:
		return "KAudioSubTapPropertyDriftCompensationQuality"
	case KAudioSubTapPropertyExtraLatency:
		return "KAudioSubTapPropertyExtraLatency"
	default:
		return fmt.Sprintf("KAudioSubTapProperty(%d)", e)
	}
}

type KAudioSystemObjectClassI uint

const (
	KAudioSystemObjectClassID KAudioSystemObjectClassI = 'a'<<24 | 's'<<16 | 'y'<<8 | 's' // 'asys'
)

func (e KAudioSystemObjectClassI) String() string {
	switch e {
	case KAudioSystemObjectClassID:
		return "KAudioSystemObjectClassID"
	default:
		return fmt.Sprintf("KAudioSystemObjectClassI(%d)", e)
	}
}

type KAudioTapClassI uint

const (
	KAudioTapClassID KAudioTapClassI = 't'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'tcls'
)

func (e KAudioTapClassI) String() string {
	switch e {
	case KAudioTapClassID:
		return "KAudioTapClassID"
	default:
		return fmt.Sprintf("KAudioTapClassI(%d)", e)
	}
}

type KAudioTapProperty uint

const (
	KAudioTapPropertyDescription KAudioTapProperty = 't'<<24 | 'd'<<16 | 's'<<8 | 'c' // 'tdsc'
	KAudioTapPropertyFormat      KAudioTapProperty = 't'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'tfmt'
	KAudioTapPropertyUID         KAudioTapProperty = 't'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'tuid'
)

func (e KAudioTapProperty) String() string {
	switch e {
	case KAudioTapPropertyDescription:
		return "KAudioTapPropertyDescription"
	case KAudioTapPropertyFormat:
		return "KAudioTapPropertyFormat"
	case KAudioTapPropertyUID:
		return "KAudioTapPropertyUID"
	default:
		return fmt.Sprintf("KAudioTapProperty(%d)", e)
	}
}

type KAudioTransportManager uint

const (
	KAudioTransportManagerCreateEndPointDevice  KAudioTransportManager = 'c'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'cdev'
	KAudioTransportManagerDestroyEndPointDevice KAudioTransportManager = 'd'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'ddev'
)

func (e KAudioTransportManager) String() string {
	switch e {
	case KAudioTransportManagerCreateEndPointDevice:
		return "KAudioTransportManagerCreateEndPointDevice"
	case KAudioTransportManagerDestroyEndPointDevice:
		return "KAudioTransportManagerDestroyEndPointDevice"
	default:
		return fmt.Sprintf("KAudioTransportManager(%d)", e)
	}
}

type KAudioTransportManagerClassI uint

const (
	KAudioTransportManagerClassID KAudioTransportManagerClassI = 't'<<24 | 'r'<<16 | 'p'<<8 | 'm' // 'trpm'
)

func (e KAudioTransportManagerClassI) String() string {
	switch e {
	case KAudioTransportManagerClassID:
		return "KAudioTransportManagerClassID"
	default:
		return fmt.Sprintf("KAudioTransportManagerClassI(%d)", e)
	}
}

type KAudioTransportManagerProperty uint

const (
	KAudioTransportManagerPropertyEndPointList           KAudioTransportManagerProperty = 'e'<<24 | 'n'<<16 | 'd'<<8 | '#' // 'end#'
	KAudioTransportManagerPropertyTranslateUIDToEndPoint KAudioTransportManagerProperty = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'e' // 'uide'
	KAudioTransportManagerPropertyTransportType          KAudioTransportManagerProperty = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
)

func (e KAudioTransportManagerProperty) String() string {
	switch e {
	case KAudioTransportManagerPropertyEndPointList:
		return "KAudioTransportManagerPropertyEndPointList"
	case KAudioTransportManagerPropertyTranslateUIDToEndPoint:
		return "KAudioTransportManagerPropertyTranslateUIDToEndPoint"
	case KAudioTransportManagerPropertyTransportType:
		return "KAudioTransportManagerPropertyTransportType"
	default:
		return fmt.Sprintf("KAudioTransportManagerProperty(%d)", e)
	}
}
