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

type KAudioAggregateDeviceClassI uint32

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

type KAudioAggregateDeviceProperty uint32

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

type KAudioAggregateDevicePropertyMasterSub uint32

const (
	// Deprecated: use kAudioAggregateDevicePropertyMainSubDevice.
	KAudioAggregateDevicePropertyMasterSubDevice KAudioAggregateDevicePropertyMasterSub = 'a'<<24 | 'm'<<16 | 's'<<8 | 't' // 'amst'
)

func (e KAudioAggregateDevicePropertyMasterSub) String() string {
	switch e {
	case KAudioAggregateDevicePropertyMasterSubDevice:
		return "KAudioAggregateDevicePropertyMasterSubDevice"
	default:
		return fmt.Sprintf("KAudioAggregateDevicePropertyMasterSub(%d)", e)
	}
}

type KAudioAggregateDriftCompensation uint32

const (
	KAudioAggregateDriftCompensationHighQuality   KAudioAggregateDriftCompensation = 0x60
	KAudioAggregateDriftCompensationLowQuality    KAudioAggregateDriftCompensation = 0x20
	KAudioAggregateDriftCompensationMaxQuality    KAudioAggregateDriftCompensation = 0x7f
	KAudioAggregateDriftCompensationMediumQuality KAudioAggregateDriftCompensation = 0x40
	KAudioAggregateDriftCompensationMinQuality    KAudioAggregateDriftCompensation = 0
)

func (e KAudioAggregateDriftCompensation) String() string {
	switch e {
	case KAudioAggregateDriftCompensationHighQuality:
		return "KAudioAggregateDriftCompensationHighQuality"
	case KAudioAggregateDriftCompensationLowQuality:
		return "KAudioAggregateDriftCompensationLowQuality"
	case KAudioAggregateDriftCompensationMaxQuality:
		return "KAudioAggregateDriftCompensationMaxQuality"
	case KAudioAggregateDriftCompensationMediumQuality:
		return "KAudioAggregateDriftCompensationMediumQuality"
	case KAudioAggregateDriftCompensationMinQuality:
		return "KAudioAggregateDriftCompensationMinQuality"
	default:
		return fmt.Sprintf("KAudioAggregateDriftCompensation(%d)", e)
	}
}

type KAudioBooleanControlClassID uint32

const (
	KAudioBooleanControlClassIDValue KAudioBooleanControlClassID = 't'<<24 | 'o'<<16 | 'g'<<8 | 'l' // 'togl'
	KAudioClipLightControlClassID    KAudioBooleanControlClassID = 'c'<<24 | 'l'<<16 | 'i'<<8 | 'p' // 'clip'
	KAudioJackControlClassID         KAudioBooleanControlClassID = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
	KAudioLFEMuteControlClassID      KAudioBooleanControlClassID = 's'<<24 | 'u'<<16 | 'b'<<8 | 'm' // 'subm'
	KAudioListenbackControlClassID   KAudioBooleanControlClassID = 'l'<<24 | 's'<<16 | 'n'<<8 | 'b' // 'lsnb'
	KAudioMuteControlClassID         KAudioBooleanControlClassID = 'm'<<24 | 'u'<<16 | 't'<<8 | 'e' // 'mute'
	KAudioPhantomPowerControlClassID KAudioBooleanControlClassID = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'n' // 'phan'
	KAudioPhaseInvertControlClassID  KAudioBooleanControlClassID = 'p'<<24 | 'h'<<16 | 's'<<8 | 'i' // 'phsi'
	KAudioSoloControlClassID         KAudioBooleanControlClassID = 's'<<24 | 'o'<<16 | 'l'<<8 | 'o' // 'solo'
	KAudioTalkbackControlClassID     KAudioBooleanControlClassID = 't'<<24 | 'a'<<16 | 'l'<<8 | 'b' // 'talb'
)

func (e KAudioBooleanControlClassID) String() string {
	switch e {
	case KAudioBooleanControlClassIDValue:
		return "KAudioBooleanControlClassIDValue"
	case KAudioClipLightControlClassID:
		return "KAudioClipLightControlClassID"
	case KAudioJackControlClassID:
		return "KAudioJackControlClassID"
	case KAudioLFEMuteControlClassID:
		return "KAudioLFEMuteControlClassID"
	case KAudioListenbackControlClassID:
		return "KAudioListenbackControlClassID"
	case KAudioMuteControlClassID:
		return "KAudioMuteControlClassID"
	case KAudioPhantomPowerControlClassID:
		return "KAudioPhantomPowerControlClassID"
	case KAudioPhaseInvertControlClassID:
		return "KAudioPhaseInvertControlClassID"
	case KAudioSoloControlClassID:
		return "KAudioSoloControlClassID"
	case KAudioTalkbackControlClassID:
		return "KAudioTalkbackControlClassID"
	default:
		return fmt.Sprintf("KAudioBooleanControlClassID(%d)", e)
	}
}

type KAudioBooleanControlProperty uint32

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

type KAudioBootChimeVolumeControlClassI uint32

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

type KAudioBoxClassI uint32

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

type KAudioBoxProperty uint32

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

type KAudioClockDeviceClassI uint32

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

type KAudioClockDeviceProperty uint32

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

type KAudioClockSourceControlPropertyItem uint32

const (
	KAudioClockSourceControlPropertyItemKind KAudioClockSourceControlPropertyItem = 'c'<<24 | 'l'<<16 | 'k'<<8 | 'k' // 'clkk'
)

func (e KAudioClockSourceControlPropertyItem) String() string {
	switch e {
	case KAudioClockSourceControlPropertyItemKind:
		return "KAudioClockSourceControlPropertyItemKind"
	default:
		return fmt.Sprintf("KAudioClockSourceControlPropertyItem(%d)", e)
	}
}

type KAudioClockSourceItemKind uint32

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

type KAudioControlClassI uint32

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

type KAudioControlPropertyScope uint32

const (
	KAudioControlPropertyElement    KAudioControlPropertyScope = 'c'<<24 | 'e'<<16 | 'l'<<8 | 'm' // 'celm'
	KAudioControlPropertyScopeValue KAudioControlPropertyScope = 'c'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'cscp'
)

func (e KAudioControlPropertyScope) String() string {
	switch e {
	case KAudioControlPropertyElement:
		return "KAudioControlPropertyElement"
	case KAudioControlPropertyScopeValue:
		return "KAudioControlPropertyScopeValue"
	default:
		return fmt.Sprintf("KAudioControlPropertyScope(%d)", e)
	}
}

const KAudioControlPropertyVariant uint32 = 'c'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'cvar'

type KAudioDeviceClassI uint32

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

type KAudioDevicePropertyConfigurationApplication uint32

const (
	KAudioDevicePropertyAvailableNominalSampleRates    KAudioDevicePropertyConfigurationApplication = 'n'<<24 | 's'<<16 | 'r'<<8 | '#' // 'nsr#'
	KAudioDevicePropertyClockDomain                    KAudioDevicePropertyConfigurationApplication = 'c'<<24 | 'l'<<16 | 'k'<<8 | 'd' // 'clkd'
	KAudioDevicePropertyConfigurationApplicationValue  KAudioDevicePropertyConfigurationApplication = 'c'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'capp'
	KAudioDevicePropertyDeviceCanBeDefaultDevice       KAudioDevicePropertyConfigurationApplication = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KAudioDevicePropertyDeviceCanBeDefaultSystemDevice KAudioDevicePropertyConfigurationApplication = 's'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'sflt'
	KAudioDevicePropertyDeviceIsAlive                  KAudioDevicePropertyConfigurationApplication = 'l'<<24 | 'i'<<16 | 'v'<<8 | 'n' // 'livn'
	KAudioDevicePropertyDeviceIsRunning                KAudioDevicePropertyConfigurationApplication = 'g'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'goin'
	KAudioDevicePropertyDeviceUID                      KAudioDevicePropertyConfigurationApplication = 'u'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'uid '
	KAudioDevicePropertyIcon                           KAudioDevicePropertyConfigurationApplication = 'i'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'icon'
	KAudioDevicePropertyIsHidden                       KAudioDevicePropertyConfigurationApplication = 'h'<<24 | 'i'<<16 | 'd'<<8 | 'n' // 'hidn'
	KAudioDevicePropertyLatency                        KAudioDevicePropertyConfigurationApplication = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KAudioDevicePropertyModelUID                       KAudioDevicePropertyConfigurationApplication = 'm'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'muid'
	KAudioDevicePropertyNominalSampleRate              KAudioDevicePropertyConfigurationApplication = 'n'<<24 | 's'<<16 | 'r'<<8 | 't' // 'nsrt'
	KAudioDevicePropertyPreferredChannelLayout         KAudioDevicePropertyConfigurationApplication = 's'<<24 | 'r'<<16 | 'n'<<8 | 'd' // 'srnd'
	KAudioDevicePropertyPreferredChannelsForStereo     KAudioDevicePropertyConfigurationApplication = 'd'<<24 | 'c'<<16 | 'h'<<8 | '2' // 'dch2'
	KAudioDevicePropertyRelatedDevices                 KAudioDevicePropertyConfigurationApplication = 'a'<<24 | 'k'<<16 | 'i'<<8 | 'n' // 'akin'
	KAudioDevicePropertySafetyOffset                   KAudioDevicePropertyConfigurationApplication = 's'<<24 | 'a'<<16 | 'f'<<8 | 't' // 'saft'
	KAudioDevicePropertyStreams                        KAudioDevicePropertyConfigurationApplication = 's'<<24 | 't'<<16 | 'm'<<8 | '#' // 'stm#'
	KAudioDevicePropertyTransportType                  KAudioDevicePropertyConfigurationApplication = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KAudioObjectPropertyControlList                    KAudioDevicePropertyConfigurationApplication = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
)

func (e KAudioDevicePropertyConfigurationApplication) String() string {
	switch e {
	case KAudioDevicePropertyAvailableNominalSampleRates:
		return "KAudioDevicePropertyAvailableNominalSampleRates"
	case KAudioDevicePropertyClockDomain:
		return "KAudioDevicePropertyClockDomain"
	case KAudioDevicePropertyConfigurationApplicationValue:
		return "KAudioDevicePropertyConfigurationApplicationValue"
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
	case KAudioObjectPropertyControlList:
		return "KAudioObjectPropertyControlList"
	default:
		return fmt.Sprintf("KAudioDevicePropertyConfigurationApplication(%d)", e)
	}
}

type KAudioDevicePropertyDeviceName uint32

const (
	KAudioDevicePropertyBufferSize                       KAudioDevicePropertyDeviceName = 'b'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'bsiz'
	KAudioDevicePropertyBufferSizeRange                  KAudioDevicePropertyDeviceName = 'b'<<24 | 's'<<16 | 'z'<<8 | '#' // 'bsz#'
	KAudioDevicePropertyChannelCategoryName              KAudioDevicePropertyDeviceName = 'c'<<24 | 'c'<<16 | 'n'<<8 | 'm' // 'ccnm'
	KAudioDevicePropertyChannelCategoryNameCFString      KAudioDevicePropertyDeviceName = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KAudioDevicePropertyChannelName                      KAudioDevicePropertyDeviceName = 'c'<<24 | 'h'<<16 | 'n'<<8 | 'm' // 'chnm'
	KAudioDevicePropertyChannelNameCFString              KAudioDevicePropertyDeviceName = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KAudioDevicePropertyChannelNominalLineLevelNameForID KAudioDevicePropertyDeviceName = 'c'<<24 | 'n'<<16 | 'l'<<8 | 'v' // 'cnlv'
	KAudioDevicePropertyChannelNumberName                KAudioDevicePropertyDeviceName = 'c'<<24 | 'n'<<16 | 'n'<<8 | 'm' // 'cnnm'
	KAudioDevicePropertyChannelNumberNameCFString        KAudioDevicePropertyDeviceName = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KAudioDevicePropertyClockSourceNameForID             KAudioDevicePropertyDeviceName = 'c'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'cscn'
	KAudioDevicePropertyDataSourceNameForID              KAudioDevicePropertyDeviceName = 's'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'sscn'
	KAudioDevicePropertyDeviceManufacturer               KAudioDevicePropertyDeviceName = 'm'<<24 | 'a'<<16 | 'k'<<8 | 'r' // 'makr'
	KAudioDevicePropertyDeviceManufacturerCFString       KAudioDevicePropertyDeviceName = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KAudioDevicePropertyDeviceNameValue                  KAudioDevicePropertyDeviceName = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	KAudioDevicePropertyDeviceNameCFString               KAudioDevicePropertyDeviceName = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KAudioDevicePropertyHighPassFilterSettingNameForID   KAudioDevicePropertyDeviceName = 'c'<<24 | 'h'<<16 | 'i'<<8 | 'p' // 'chip'
	KAudioDevicePropertyPlayThruDestinationNameForID     KAudioDevicePropertyDeviceName = 'm'<<24 | 'd'<<16 | 'd'<<8 | 'n' // 'mddn'
	KAudioDevicePropertyRegisterBufferList               KAudioDevicePropertyDeviceName = 'r'<<24 | 'b'<<16 | 'u'<<8 | 'f' // 'rbuf'
	KAudioDevicePropertyStreamFormat                     KAudioDevicePropertyDeviceName = 's'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'sfmt'
	KAudioDevicePropertyStreamFormatMatch                KAudioDevicePropertyDeviceName = 's'<<24 | 'f'<<16 | 'm'<<8 | 'm' // 'sfmm'
	KAudioDevicePropertyStreamFormatSupported            KAudioDevicePropertyDeviceName = 's'<<24 | 'f'<<16 | 'm'<<8 | '?' // 'sfm?'
	KAudioDevicePropertyStreamFormats                    KAudioDevicePropertyDeviceName = 's'<<24 | 'f'<<16 | 'm'<<8 | '#' // 'sfm#'
	KAudioDevicePropertySupportsMixing                   KAudioDevicePropertyDeviceName = 'm'<<24 | 'i'<<16 | 'x'<<8 | '?' // 'mix?'
)

func (e KAudioDevicePropertyDeviceName) String() string {
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
	case KAudioDevicePropertyChannelNominalLineLevelNameForID:
		return "KAudioDevicePropertyChannelNominalLineLevelNameForID"
	case KAudioDevicePropertyChannelNumberName:
		return "KAudioDevicePropertyChannelNumberName"
	case KAudioDevicePropertyChannelNumberNameCFString:
		return "KAudioDevicePropertyChannelNumberNameCFString"
	case KAudioDevicePropertyClockSourceNameForID:
		return "KAudioDevicePropertyClockSourceNameForID"
	case KAudioDevicePropertyDataSourceNameForID:
		return "KAudioDevicePropertyDataSourceNameForID"
	case KAudioDevicePropertyDeviceManufacturer:
		return "KAudioDevicePropertyDeviceManufacturer"
	case KAudioDevicePropertyDeviceManufacturerCFString:
		return "KAudioDevicePropertyDeviceManufacturerCFString"
	case KAudioDevicePropertyDeviceNameValue:
		return "KAudioDevicePropertyDeviceNameValue"
	case KAudioDevicePropertyDeviceNameCFString:
		return "KAudioDevicePropertyDeviceNameCFString"
	case KAudioDevicePropertyHighPassFilterSettingNameForID:
		return "KAudioDevicePropertyHighPassFilterSettingNameForID"
	case KAudioDevicePropertyPlayThruDestinationNameForID:
		return "KAudioDevicePropertyPlayThruDestinationNameForID"
	case KAudioDevicePropertyRegisterBufferList:
		return "KAudioDevicePropertyRegisterBufferList"
	case KAudioDevicePropertyStreamFormat:
		return "KAudioDevicePropertyStreamFormat"
	case KAudioDevicePropertyStreamFormatMatch:
		return "KAudioDevicePropertyStreamFormatMatch"
	case KAudioDevicePropertyStreamFormatSupported:
		return "KAudioDevicePropertyStreamFormatSupported"
	case KAudioDevicePropertyStreamFormats:
		return "KAudioDevicePropertyStreamFormats"
	case KAudioDevicePropertySupportsMixing:
		return "KAudioDevicePropertySupportsMixing"
	default:
		return fmt.Sprintf("KAudioDevicePropertyDeviceName(%d)", e)
	}
}

type KAudioDevicePropertyJackIsConnected uint32

const (
	KAudioDevicePropertyChannelNominalLineLevel                  KAudioDevicePropertyJackIsConnected = 'n'<<24 | 'l'<<16 | 'v'<<8 | 'l' // 'nlvl'
	KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString KAudioDevicePropertyJackIsConnected = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'l' // 'lcnl'
	KAudioDevicePropertyChannelNominalLineLevels                 KAudioDevicePropertyJackIsConnected = 'n'<<24 | 'l'<<16 | 'v'<<8 | '#' // 'nlv#'
	KAudioDevicePropertyClipLight                                KAudioDevicePropertyJackIsConnected = 'c'<<24 | 'l'<<16 | 'i'<<8 | 'p' // 'clip'
	KAudioDevicePropertyClockSource                              KAudioDevicePropertyJackIsConnected = 'c'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'csrc'
	KAudioDevicePropertyClockSourceKindForID                     KAudioDevicePropertyJackIsConnected = 'c'<<24 | 's'<<16 | 'c'<<8 | 'k' // 'csck'
	KAudioDevicePropertyClockSourceNameForIDCFString             KAudioDevicePropertyJackIsConnected = 'l'<<24 | 'c'<<16 | 's'<<8 | 'n' // 'lcsn'
	KAudioDevicePropertyClockSources                             KAudioDevicePropertyJackIsConnected = 'c'<<24 | 's'<<16 | 'c'<<8 | '#' // 'csc#'
	KAudioDevicePropertyDataSource                               KAudioDevicePropertyJackIsConnected = 's'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'ssrc'
	KAudioDevicePropertyDataSourceKindForID                      KAudioDevicePropertyJackIsConnected = 's'<<24 | 's'<<16 | 'c'<<8 | 'k' // 'ssck'
	KAudioDevicePropertyDataSourceNameForIDCFString              KAudioDevicePropertyJackIsConnected = 'l'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'lscn'
	KAudioDevicePropertyDataSources                              KAudioDevicePropertyJackIsConnected = 's'<<24 | 's'<<16 | 'c'<<8 | '#' // 'ssc#'
	KAudioDevicePropertyHighPassFilterSetting                    KAudioDevicePropertyJackIsConnected = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'f' // 'hipf'
	KAudioDevicePropertyHighPassFilterSettingNameForIDCFString   KAudioDevicePropertyJackIsConnected = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'l' // 'hipl'
	KAudioDevicePropertyHighPassFilterSettings                   KAudioDevicePropertyJackIsConnected = 'h'<<24 | 'i'<<16 | 'p'<<8 | '#' // 'hip#'
	KAudioDevicePropertyJackIsConnectedValue                     KAudioDevicePropertyJackIsConnected = 'j'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'jack'
	KAudioDevicePropertyListenback                               KAudioDevicePropertyJackIsConnected = 'l'<<24 | 's'<<16 | 'n'<<8 | 'b' // 'lsnb'
	KAudioDevicePropertyMute                                     KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'u'<<16 | 't'<<8 | 'e' // 'mute'
	KAudioDevicePropertyPhantomPower                             KAudioDevicePropertyJackIsConnected = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'n' // 'phan'
	KAudioDevicePropertyPhaseInvert                              KAudioDevicePropertyJackIsConnected = 'p'<<24 | 'h'<<16 | 's'<<8 | 'i' // 'phsi'
	KAudioDevicePropertyPlayThru                                 KAudioDevicePropertyJackIsConnected = 't'<<24 | 'h'<<16 | 'r'<<8 | 'u' // 'thru'
	KAudioDevicePropertyPlayThruDestination                      KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'd'<<16 | 'd'<<8 | 's' // 'mdds'
	KAudioDevicePropertyPlayThruDestinationNameForIDCFString     KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'd'<<16 | 'd'<<8 | 'c' // 'mddc'
	KAudioDevicePropertyPlayThruDestinations                     KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'd'<<16 | 'd'<<8 | '#' // 'mdd#'
	KAudioDevicePropertyPlayThruSolo                             KAudioDevicePropertyJackIsConnected = 't'<<24 | 'h'<<16 | 'r'<<8 | 's' // 'thrs'
	KAudioDevicePropertyPlayThruStereoPan                        KAudioDevicePropertyJackIsConnected = 'm'<<24 | 's'<<16 | 'p'<<8 | 'n' // 'mspn'
	KAudioDevicePropertyPlayThruStereoPanChannels                KAudioDevicePropertyJackIsConnected = 'm'<<24 | 's'<<16 | 'p'<<8 | '#' // 'msp#'
	KAudioDevicePropertyPlayThruVolumeDecibels                   KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'v'<<16 | 'd'<<8 | 'b' // 'mvdb'
	KAudioDevicePropertyPlayThruVolumeDecibelsToScalar           KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'v'<<16 | '2'<<8 | 's' // 'mv2s'
	KAudioDevicePropertyPlayThruVolumeRangeDecibels              KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'v'<<16 | 'd'<<8 | '#' // 'mvd#'
	KAudioDevicePropertyPlayThruVolumeScalar                     KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'v'<<16 | 's'<<8 | 'c' // 'mvsc'
	KAudioDevicePropertyPlayThruVolumeScalarToDecibels           KAudioDevicePropertyJackIsConnected = 'm'<<24 | 'v'<<16 | '2'<<8 | 'd' // 'mv2d'
	KAudioDevicePropertySolo                                     KAudioDevicePropertyJackIsConnected = 's'<<24 | 'o'<<16 | 'l'<<8 | 'o' // 'solo'
	KAudioDevicePropertyStereoPan                                KAudioDevicePropertyJackIsConnected = 's'<<24 | 'p'<<16 | 'a'<<8 | 'n' // 'span'
	KAudioDevicePropertyStereoPanChannels                        KAudioDevicePropertyJackIsConnected = 's'<<24 | 'p'<<16 | 'n'<<8 | '#' // 'spn#'
	KAudioDevicePropertySubMute                                  KAudioDevicePropertyJackIsConnected = 's'<<24 | 'm'<<16 | 'u'<<8 | 't' // 'smut'
	KAudioDevicePropertySubVolumeDecibels                        KAudioDevicePropertyJackIsConnected = 's'<<24 | 'v'<<16 | 'l'<<8 | 'd' // 'svld'
	KAudioDevicePropertySubVolumeDecibelsToScalar                KAudioDevicePropertyJackIsConnected = 's'<<24 | 'd'<<16 | '2'<<8 | 'v' // 'sd2v'
	KAudioDevicePropertySubVolumeRangeDecibels                   KAudioDevicePropertyJackIsConnected = 's'<<24 | 'v'<<16 | 'd'<<8 | '#' // 'svd#'
	KAudioDevicePropertySubVolumeScalar                          KAudioDevicePropertyJackIsConnected = 's'<<24 | 'v'<<16 | 'l'<<8 | 'm' // 'svlm'
	KAudioDevicePropertySubVolumeScalarToDecibels                KAudioDevicePropertyJackIsConnected = 's'<<24 | 'v'<<16 | '2'<<8 | 'd' // 'sv2d'
	KAudioDevicePropertyTalkback                                 KAudioDevicePropertyJackIsConnected = 't'<<24 | 'a'<<16 | 'l'<<8 | 'b' // 'talb'
	KAudioDevicePropertyVoiceActivityDetectionEnable             KAudioDevicePropertyJackIsConnected = 'v'<<24 | 'A'<<16 | 'd'<<8 | '+' // 'vAd+'
	KAudioDevicePropertyVoiceActivityDetectionState              KAudioDevicePropertyJackIsConnected = 'v'<<24 | 'A'<<16 | 'd'<<8 | 'S' // 'vAdS'
	KAudioDevicePropertyVolumeDecibels                           KAudioDevicePropertyJackIsConnected = 'v'<<24 | 'o'<<16 | 'l'<<8 | 'd' // 'vold'
	KAudioDevicePropertyVolumeDecibelsToScalar                   KAudioDevicePropertyJackIsConnected = 'd'<<24 | 'b'<<16 | '2'<<8 | 'v' // 'db2v'
	KAudioDevicePropertyVolumeRangeDecibels                      KAudioDevicePropertyJackIsConnected = 'v'<<24 | 'd'<<16 | 'b'<<8 | '#' // 'vdb#'
	KAudioDevicePropertyVolumeScalar                             KAudioDevicePropertyJackIsConnected = 'v'<<24 | 'o'<<16 | 'l'<<8 | 'm' // 'volm'
	KAudioDevicePropertyVolumeScalarToDecibels                   KAudioDevicePropertyJackIsConnected = 'v'<<24 | '2'<<16 | 'd'<<8 | 'b' // 'v2db'
	KAudioDevicePropertyWantsControlsRestored                    KAudioDevicePropertyJackIsConnected = 'r'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'resc'
	KAudioDevicePropertyWantsStreamFormatsRestored               KAudioDevicePropertyJackIsConnected = 'r'<<24 | 'e'<<16 | 's'<<8 | 'f' // 'resf'
)

func (e KAudioDevicePropertyJackIsConnected) String() string {
	switch e {
	case KAudioDevicePropertyChannelNominalLineLevel:
		return "KAudioDevicePropertyChannelNominalLineLevel"
	case KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString:
		return "KAudioDevicePropertyChannelNominalLineLevelNameForIDCFString"
	case KAudioDevicePropertyChannelNominalLineLevels:
		return "KAudioDevicePropertyChannelNominalLineLevels"
	case KAudioDevicePropertyClipLight:
		return "KAudioDevicePropertyClipLight"
	case KAudioDevicePropertyClockSource:
		return "KAudioDevicePropertyClockSource"
	case KAudioDevicePropertyClockSourceKindForID:
		return "KAudioDevicePropertyClockSourceKindForID"
	case KAudioDevicePropertyClockSourceNameForIDCFString:
		return "KAudioDevicePropertyClockSourceNameForIDCFString"
	case KAudioDevicePropertyClockSources:
		return "KAudioDevicePropertyClockSources"
	case KAudioDevicePropertyDataSource:
		return "KAudioDevicePropertyDataSource"
	case KAudioDevicePropertyDataSourceKindForID:
		return "KAudioDevicePropertyDataSourceKindForID"
	case KAudioDevicePropertyDataSourceNameForIDCFString:
		return "KAudioDevicePropertyDataSourceNameForIDCFString"
	case KAudioDevicePropertyDataSources:
		return "KAudioDevicePropertyDataSources"
	case KAudioDevicePropertyHighPassFilterSetting:
		return "KAudioDevicePropertyHighPassFilterSetting"
	case KAudioDevicePropertyHighPassFilterSettingNameForIDCFString:
		return "KAudioDevicePropertyHighPassFilterSettingNameForIDCFString"
	case KAudioDevicePropertyHighPassFilterSettings:
		return "KAudioDevicePropertyHighPassFilterSettings"
	case KAudioDevicePropertyJackIsConnectedValue:
		return "KAudioDevicePropertyJackIsConnectedValue"
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
	case KAudioDevicePropertyPlayThruVolumeRangeDecibels:
		return "KAudioDevicePropertyPlayThruVolumeRangeDecibels"
	case KAudioDevicePropertyPlayThruVolumeScalar:
		return "KAudioDevicePropertyPlayThruVolumeScalar"
	case KAudioDevicePropertyPlayThruVolumeScalarToDecibels:
		return "KAudioDevicePropertyPlayThruVolumeScalarToDecibels"
	case KAudioDevicePropertySolo:
		return "KAudioDevicePropertySolo"
	case KAudioDevicePropertyStereoPan:
		return "KAudioDevicePropertyStereoPan"
	case KAudioDevicePropertyStereoPanChannels:
		return "KAudioDevicePropertyStereoPanChannels"
	case KAudioDevicePropertySubMute:
		return "KAudioDevicePropertySubMute"
	case KAudioDevicePropertySubVolumeDecibels:
		return "KAudioDevicePropertySubVolumeDecibels"
	case KAudioDevicePropertySubVolumeDecibelsToScalar:
		return "KAudioDevicePropertySubVolumeDecibelsToScalar"
	case KAudioDevicePropertySubVolumeRangeDecibels:
		return "KAudioDevicePropertySubVolumeRangeDecibels"
	case KAudioDevicePropertySubVolumeScalar:
		return "KAudioDevicePropertySubVolumeScalar"
	case KAudioDevicePropertySubVolumeScalarToDecibels:
		return "KAudioDevicePropertySubVolumeScalarToDecibels"
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
		return fmt.Sprintf("KAudioDevicePropertyJackIsConnected(%d)", e)
	}
}

type KAudioDevicePropertyPlugIn uint32

const (
	KAudioDeviceProcessorOverload                KAudioDevicePropertyPlugIn = 'o'<<24 | 'v'<<16 | 'e'<<8 | 'r' // 'over'
	KAudioDevicePropertyActualSampleRate         KAudioDevicePropertyPlugIn = 'a'<<24 | 's'<<16 | 'r'<<8 | 't' // 'asrt'
	KAudioDevicePropertyBufferFrameSize          KAudioDevicePropertyPlugIn = 'f'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'fsiz'
	KAudioDevicePropertyBufferFrameSizeRange     KAudioDevicePropertyPlugIn = 'f'<<24 | 's'<<16 | 'z'<<8 | '#' // 'fsz#'
	KAudioDevicePropertyClockDevice              KAudioDevicePropertyPlugIn = 'a'<<24 | 'p'<<16 | 'c'<<8 | 'd' // 'apcd'
	KAudioDevicePropertyDeviceHasChanged         KAudioDevicePropertyPlugIn = 'd'<<24 | 'i'<<16 | 'f'<<8 | 'f' // 'diff'
	KAudioDevicePropertyDeviceIsRunningSomewhere KAudioDevicePropertyPlugIn = 'g'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'gone'
	KAudioDevicePropertyHogMode                  KAudioDevicePropertyPlugIn = 'o'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'oink'
	KAudioDevicePropertyIOCycleUsage             KAudioDevicePropertyPlugIn = 'n'<<24 | 'c'<<16 | 'y'<<8 | 'c' // 'ncyc'
	KAudioDevicePropertyIOProcStreamUsage        KAudioDevicePropertyPlugIn = 's'<<24 | 'u'<<16 | 's'<<8 | 'e' // 'suse'
	KAudioDevicePropertyIOStoppedAbnormally      KAudioDevicePropertyPlugIn = 's'<<24 | 't'<<16 | 'p'<<8 | 'd' // 'stpd'
	// KAudioDevicePropertyIOThreadOSWorkgroup: The device’s workgroup object, which you use to coordinate your threads with the threads of the device.
	KAudioDevicePropertyIOThreadOSWorkgroup          KAudioDevicePropertyPlugIn = 'o'<<24 | 's'<<16 | 'w'<<8 | 'g' // 'oswg'
	KAudioDevicePropertyPlugInValue                  KAudioDevicePropertyPlugIn = 'p'<<24 | 'l'<<16 | 'u'<<8 | 'g' // 'plug'
	KAudioDevicePropertyProcessMute                  KAudioDevicePropertyPlugIn = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'm' // 'appm'
	KAudioDevicePropertyStreamConfiguration          KAudioDevicePropertyPlugIn = 's'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'slay'
	KAudioDevicePropertyUsesVariableBufferFrameSizes KAudioDevicePropertyPlugIn = 'v'<<24 | 'f'<<16 | 's'<<8 | 'z' // 'vfsz'
)

func (e KAudioDevicePropertyPlugIn) String() string {
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
	case KAudioDevicePropertyPlugInValue:
		return "KAudioDevicePropertyPlugInValue"
	case KAudioDevicePropertyProcessMute:
		return "KAudioDevicePropertyProcessMute"
	case KAudioDevicePropertyStreamConfiguration:
		return "KAudioDevicePropertyStreamConfiguration"
	case KAudioDevicePropertyUsesVariableBufferFrameSizes:
		return "KAudioDevicePropertyUsesVariableBufferFrameSizes"
	default:
		return fmt.Sprintf("KAudioDevicePropertyPlugIn(%d)", e)
	}
}

type KAudioDevicePropertyScope uint32

const (
	KAudioDevicePropertyScopeInput       KAudioDevicePropertyScope = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KAudioDevicePropertyScopeOutput      KAudioDevicePropertyScope = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KAudioDevicePropertyScopePlayThrough KAudioDevicePropertyScope = 'p'<<24 | 't'<<16 | 'r'<<8 | 'u' // 'ptru'
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

type KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction uint32

const (
	KAudioDevicePropertyDriverShouldOwniSub                            KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction = 'i'<<24 | 's'<<16 | 'u'<<8 | 'b' // 'isub'
	KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction = 'm'<<24 | 'v'<<16 | 't'<<8 | 'f' // 'mvtf'
	KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction      KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction = 's'<<24 | 'v'<<16 | 't'<<8 | 'f' // 'svtf'
	KAudioDevicePropertyVolumeDecibelsToScalarTransferFunctionValue    KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction = 'v'<<24 | 'c'<<16 | 't'<<8 | 'f' // 'vctf'
)

func (e KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction) String() string {
	switch e {
	case KAudioDevicePropertyDriverShouldOwniSub:
		return "KAudioDevicePropertyDriverShouldOwniSub"
	case KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction:
		return "KAudioDevicePropertyPlayThruVolumeDecibelsToScalarTransferFunction"
	case KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction:
		return "KAudioDevicePropertySubVolumeDecibelsToScalarTransferFunction"
	case KAudioDevicePropertyVolumeDecibelsToScalarTransferFunctionValue:
		return "KAudioDevicePropertyVolumeDecibelsToScalarTransferFunctionValue"
	default:
		return fmt.Sprintf("KAudioDevicePropertyVolumeDecibelsToScalarTransferFunction(%d)", e)
	}
}

type KAudioDevicePropertyZeroTimeStampPeriod uint

const (
	KAudioDevicePropertyClockAlgorithm           KAudioDevicePropertyZeroTimeStampPeriod = 0
	KAudioDevicePropertyClockIsStable            KAudioDevicePropertyZeroTimeStampPeriod = 0
	KAudioDevicePropertyZeroTimeStampPeriodValue KAudioDevicePropertyZeroTimeStampPeriod = 0
)

func (e KAudioDevicePropertyZeroTimeStampPeriod) String() string {
	switch e {
	case KAudioDevicePropertyClockAlgorithm:
		return "KAudioDevicePropertyClockAlgorithm"
	default:
		return fmt.Sprintf("KAudioDevicePropertyZeroTimeStampPeriod(%d)", e)
	}
}

type KAudioDeviceStartTime uint32

const (
	KAudioDeviceStartTimeDontConsultDeviceFlag KAudioDeviceStartTime = 2
	KAudioDeviceStartTimeDontConsultHALFlag    KAudioDeviceStartTime = 4
	KAudioDeviceStartTimeIsInputFlag           KAudioDeviceStartTime = 1
)

func (e KAudioDeviceStartTime) String() string {
	switch e {
	case KAudioDeviceStartTimeDontConsultDeviceFlag:
		return "KAudioDeviceStartTimeDontConsultDeviceFlag"
	case KAudioDeviceStartTimeDontConsultHALFlag:
		return "KAudioDeviceStartTimeDontConsultHALFlag"
	case KAudioDeviceStartTimeIsInputFlag:
		return "KAudioDeviceStartTimeIsInputFlag"
	default:
		return fmt.Sprintf("KAudioDeviceStartTime(%d)", e)
	}
}

type KAudioDeviceTransportType uint32

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
	KAudioDeviceTransportTypeContinuityCapture KAudioDeviceTransportType = 'c'<<24 | 'c'<<16 | 'a'<<8 | 'p' // 'ccap'
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
	case KAudioDeviceTransportTypeContinuityCapture:
		return "KAudioDeviceTransportTypeContinuityCapture"
	default:
		return fmt.Sprintf("KAudioDeviceTransportType(%d)", e)
	}
}

type KAudioDeviceTransportTypeAuto uint32

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

const KAudioDeviceUnknown uint32 = 0

type KAudioEndPointClassI uint32

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

type KAudioEndPointDeviceClassI uint32

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

type KAudioEndPointDeviceProperty uint32

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

type KAudioHardwareNoError int32

const (
	KAudioDevicePermissionsError            KAudioHardwareNoError = '!'<<24 | 'h'<<16 | 'o'<<8 | 'g' // '!hog'
	KAudioDeviceUnsupportedFormatError      KAudioHardwareNoError = '!'<<24 | 'd'<<16 | 'a'<<8 | 't' // '!dat'
	KAudioHardwareBadDeviceError            KAudioHardwareNoError = '!'<<24 | 'd'<<16 | 'e'<<8 | 'v' // '!dev'
	KAudioHardwareBadObjectError            KAudioHardwareNoError = '!'<<24 | 'o'<<16 | 'b'<<8 | 'j' // '!obj'
	KAudioHardwareBadPropertySizeError      KAudioHardwareNoError = '!'<<24 | 's'<<16 | 'i'<<8 | 'z' // '!siz'
	KAudioHardwareBadStreamError            KAudioHardwareNoError = '!'<<24 | 's'<<16 | 't'<<8 | 'r' // '!str'
	KAudioHardwareIllegalOperationError     KAudioHardwareNoError = 'n'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'nope'
	KAudioHardwareNoErrorValue              KAudioHardwareNoError = 0
	KAudioHardwareNotReadyError             KAudioHardwareNoError = 'n'<<24 | 'r'<<16 | 'd'<<8 | 'y' // 'nrdy'
	KAudioHardwareNotRunningError           KAudioHardwareNoError = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KAudioHardwareUnknownPropertyError      KAudioHardwareNoError = 'w'<<24 | 'h'<<16 | 'o'<<8 | '?' // 'who?'
	KAudioHardwareUnspecifiedError          KAudioHardwareNoError = 'w'<<24 | 'h'<<16 | 'a'<<8 | 't' // 'what'
	KAudioHardwareUnsupportedOperationError KAudioHardwareNoError = 'u'<<24 | 'n'<<16 | 'o'<<8 | 'p' // 'unop'
)

func (e KAudioHardwareNoError) String() string {
	switch e {
	case KAudioDevicePermissionsError:
		return "KAudioDevicePermissionsError"
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
	case KAudioHardwareNoErrorValue:
		return "KAudioHardwareNoErrorValue"
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
	default:
		return fmt.Sprintf("KAudioHardwareNoError(%d)", e)
	}
}

type KAudioHardwarePropertyBootChimeVolume uint32

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

type KAudioHardwarePropertyDevices uint32

const (
	KAudioHardwarePropertyBoxList                             KAudioHardwarePropertyDevices = 'b'<<24 | 'o'<<16 | 'x'<<8 | '#' // 'box#'
	KAudioHardwarePropertyClockDeviceList                     KAudioHardwarePropertyDevices = 'c'<<24 | 'l'<<16 | 'k'<<8 | '#' // 'clk#'
	KAudioHardwarePropertyDefaultInputDevice                  KAudioHardwarePropertyDevices = 'd'<<24 | 'I'<<16 | 'n'<<8 | ' ' // 'dIn '
	KAudioHardwarePropertyDefaultOutputDevice                 KAudioHardwarePropertyDevices = 'd'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'dOut'
	KAudioHardwarePropertyDefaultSystemOutputDevice           KAudioHardwarePropertyDevices = 's'<<24 | 'O'<<16 | 'u'<<8 | 't' // 'sOut'
	KAudioHardwarePropertyDevicesValue                        KAudioHardwarePropertyDevices = 'd'<<24 | 'e'<<16 | 'v'<<8 | '#' // 'dev#'
	KAudioHardwarePropertyHogModeIsAllowed                    KAudioHardwarePropertyDevices = 'h'<<24 | 'o'<<16 | 'g'<<8 | 'r' // 'hogr'
	KAudioHardwarePropertyIsInitingOrExiting                  KAudioHardwarePropertyDevices = 'i'<<24 | 'n'<<16 | 'o'<<8 | 't' // 'inot'
	KAudioHardwarePropertyMixStereoToMono                     KAudioHardwarePropertyDevices = 's'<<24 | 't'<<16 | 'm'<<8 | 'o' // 'stmo'
	KAudioHardwarePropertyPlugInList                          KAudioHardwarePropertyDevices = 'p'<<24 | 'l'<<16 | 'g'<<8 | '#' // 'plg#'
	KAudioHardwarePropertyPowerHint                           KAudioHardwarePropertyDevices = 'p'<<24 | 'o'<<16 | 'w'<<8 | 'h' // 'powh'
	KAudioHardwarePropertyProcessInputMute                    KAudioHardwarePropertyDevices = 'p'<<24 | 'm'<<16 | 'i'<<8 | 'n' // 'pmin'
	KAudioHardwarePropertyProcessIsAudible                    KAudioHardwarePropertyDevices = 'p'<<24 | 'm'<<16 | 'u'<<8 | 't' // 'pmut'
	KAudioHardwarePropertyProcessIsMain                       KAudioHardwarePropertyDevices = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'n' // 'main'
	KAudioHardwarePropertyProcessObjectList                   KAudioHardwarePropertyDevices = 'p'<<24 | 'r'<<16 | 's'<<8 | '#' // 'prs#'
	KAudioHardwarePropertyServiceRestarted                    KAudioHardwarePropertyDevices = 's'<<24 | 'r'<<16 | 's'<<8 | 't' // 'srst'
	KAudioHardwarePropertySleepingIsAllowed                   KAudioHardwarePropertyDevices = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KAudioHardwarePropertyTapList                             KAudioHardwarePropertyDevices = 't'<<24 | 'p'<<16 | 's'<<8 | '#' // 'tps#'
	KAudioHardwarePropertyTranslateBundleIDToPlugIn           KAudioHardwarePropertyDevices = 'b'<<24 | 'i'<<16 | 'd'<<8 | 'p' // 'bidp'
	KAudioHardwarePropertyTranslateBundleIDToTransportManager KAudioHardwarePropertyDevices = 't'<<24 | 'm'<<16 | 'b'<<8 | 'i' // 'tmbi'
	KAudioHardwarePropertyTranslatePIDToProcessObject         KAudioHardwarePropertyDevices = 'i'<<24 | 'd'<<16 | '2'<<8 | 'p' // 'id2p'
	KAudioHardwarePropertyTranslateUIDToBox                   KAudioHardwarePropertyDevices = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'b' // 'uidb'
	KAudioHardwarePropertyTranslateUIDToClockDevice           KAudioHardwarePropertyDevices = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'c' // 'uidc'
	KAudioHardwarePropertyTranslateUIDToDevice                KAudioHardwarePropertyDevices = 'u'<<24 | 'i'<<16 | 'd'<<8 | 'd' // 'uidd'
	KAudioHardwarePropertyTranslateUIDToTap                   KAudioHardwarePropertyDevices = 'u'<<24 | 'i'<<16 | 'd'<<8 | 't' // 'uidt'
	KAudioHardwarePropertyTransportManagerList                KAudioHardwarePropertyDevices = 't'<<24 | 'm'<<16 | 'g'<<8 | '#' // 'tmg#'
	KAudioHardwarePropertyUnloadingIsAllowed                  KAudioHardwarePropertyDevices = 'u'<<24 | 'n'<<16 | 'l'<<8 | 'd' // 'unld'
	KAudioHardwarePropertyUserIDChanged                       KAudioHardwarePropertyDevices = 'e'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'euid'
	KAudioHardwarePropertyUserSessionIsActiveOrHeadless       KAudioHardwarePropertyDevices = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
)

func (e KAudioHardwarePropertyDevices) String() string {
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
	case KAudioHardwarePropertyDevicesValue:
		return "KAudioHardwarePropertyDevicesValue"
	case KAudioHardwarePropertyHogModeIsAllowed:
		return "KAudioHardwarePropertyHogModeIsAllowed"
	case KAudioHardwarePropertyIsInitingOrExiting:
		return "KAudioHardwarePropertyIsInitingOrExiting"
	case KAudioHardwarePropertyMixStereoToMono:
		return "KAudioHardwarePropertyMixStereoToMono"
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
	default:
		return fmt.Sprintf("KAudioHardwarePropertyDevices(%d)", e)
	}
}

type KAudioHardwarePropertyRunLoop uint32

const (
	KAudioHardwarePropertyDeviceForUID      KAudioHardwarePropertyRunLoop = 'd'<<24 | 'u'<<16 | 'i'<<8 | 'd' // 'duid'
	KAudioHardwarePropertyPlugInForBundleID KAudioHardwarePropertyRunLoop = 'p'<<24 | 'i'<<16 | 'b'<<8 | 'i' // 'pibi'
	KAudioHardwarePropertyRunLoopValue      KAudioHardwarePropertyRunLoop = 'r'<<24 | 'n'<<16 | 'l'<<8 | 'p' // 'rnlp'
	// Deprecated: use kAudioHardwarePropertyProcessIsMain.
	KAudioHardwarePropertyProcessIsMaster KAudioHardwarePropertyRunLoop = 'm'<<24 | 'a'<<16 | 's'<<8 | 't' // 'mast'
)

func (e KAudioHardwarePropertyRunLoop) String() string {
	switch e {
	case KAudioHardwarePropertyDeviceForUID:
		return "KAudioHardwarePropertyDeviceForUID"
	case KAudioHardwarePropertyPlugInForBundleID:
		return "KAudioHardwarePropertyPlugInForBundleID"
	case KAudioHardwarePropertyRunLoopValue:
		return "KAudioHardwarePropertyRunLoopValue"
	case KAudioHardwarePropertyProcessIsMaster:
		return "KAudioHardwarePropertyProcessIsMaster"
	default:
		return fmt.Sprintf("KAudioHardwarePropertyRunLoop(%d)", e)
	}
}

type KAudioISubOwnerControlClassI uint32

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

type KAudioLevelControlClassID uint32

const (
	KAudioLFEVolumeControlClassID  KAudioLevelControlClassID = 's'<<24 | 'u'<<16 | 'b'<<8 | 'v' // 'subv'
	KAudioLevelControlClassIDValue KAudioLevelControlClassID = 'l'<<24 | 'e'<<16 | 'v'<<8 | 'l' // 'levl'
	KAudioVolumeControlClassID     KAudioLevelControlClassID = 'v'<<24 | 'l'<<16 | 'm'<<8 | 'e' // 'vlme'
)

func (e KAudioLevelControlClassID) String() string {
	switch e {
	case KAudioLFEVolumeControlClassID:
		return "KAudioLFEVolumeControlClassID"
	case KAudioLevelControlClassIDValue:
		return "KAudioLevelControlClassIDValue"
	case KAudioVolumeControlClassID:
		return "KAudioVolumeControlClassID"
	default:
		return fmt.Sprintf("KAudioLevelControlClassID(%d)", e)
	}
}

type KAudioLevelControlProperty uint32

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

type KAudioLevelControlPropertyDecibelsToScalarTransfer uint32

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

type KAudioObject uint32

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

type KAudioObjectClassI uint32

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

type KAudioObjectClassID uint32

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

type KAudioObjectPropertyBaseClass uint32

const (
	KAudioObjectPropertyBaseClassValue      KAudioObjectPropertyBaseClass = 'b'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'bcls'
	KAudioObjectPropertyClass               KAudioObjectPropertyBaseClass = 'c'<<24 | 'l'<<16 | 'a'<<8 | 's' // 'clas'
	KAudioObjectPropertyElementCategoryName KAudioObjectPropertyBaseClass = 'l'<<24 | 'c'<<16 | 'c'<<8 | 'n' // 'lccn'
	KAudioObjectPropertyElementName         KAudioObjectPropertyBaseClass = 'l'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'lchn'
	KAudioObjectPropertyElementNumberName   KAudioObjectPropertyBaseClass = 'l'<<24 | 'c'<<16 | 'n'<<8 | 'n' // 'lcnn'
	KAudioObjectPropertyFirmwareVersion     KAudioObjectPropertyBaseClass = 'f'<<24 | 'w'<<16 | 'v'<<8 | 'n' // 'fwvn'
	KAudioObjectPropertyIdentify            KAudioObjectPropertyBaseClass = 'i'<<24 | 'd'<<16 | 'e'<<8 | 'n' // 'iden'
	KAudioObjectPropertyManufacturer        KAudioObjectPropertyBaseClass = 'l'<<24 | 'm'<<16 | 'a'<<8 | 'k' // 'lmak'
	KAudioObjectPropertyModelName           KAudioObjectPropertyBaseClass = 'l'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'lmod'
	KAudioObjectPropertyName                KAudioObjectPropertyBaseClass = 'l'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'lnam'
	KAudioObjectPropertyOwnedObjects        KAudioObjectPropertyBaseClass = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KAudioObjectPropertyOwner               KAudioObjectPropertyBaseClass = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
	KAudioObjectPropertySerialNumber        KAudioObjectPropertyBaseClass = 's'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'snum'
)

func (e KAudioObjectPropertyBaseClass) String() string {
	switch e {
	case KAudioObjectPropertyBaseClassValue:
		return "KAudioObjectPropertyBaseClassValue"
	case KAudioObjectPropertyClass:
		return "KAudioObjectPropertyClass"
	case KAudioObjectPropertyElementCategoryName:
		return "KAudioObjectPropertyElementCategoryName"
	case KAudioObjectPropertyElementName:
		return "KAudioObjectPropertyElementName"
	case KAudioObjectPropertyElementNumberName:
		return "KAudioObjectPropertyElementNumberName"
	case KAudioObjectPropertyFirmwareVersion:
		return "KAudioObjectPropertyFirmwareVersion"
	case KAudioObjectPropertyIdentify:
		return "KAudioObjectPropertyIdentify"
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
	case KAudioObjectPropertySerialNumber:
		return "KAudioObjectPropertySerialNumber"
	default:
		return fmt.Sprintf("KAudioObjectPropertyBaseClass(%d)", e)
	}
}

type KAudioObjectPropertyCreator uint32

const (
	KAudioObjectPropertyCreatorValue    KAudioObjectPropertyCreator = 'o'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'oplg'
	KAudioObjectPropertyListenerAdded   KAudioObjectPropertyCreator = 'l'<<24 | 'i'<<16 | 's'<<8 | 'a' // 'lisa'
	KAudioObjectPropertyListenerRemoved KAudioObjectPropertyCreator = 'l'<<24 | 'i'<<16 | 's'<<8 | 'r' // 'lisr'
)

func (e KAudioObjectPropertyCreator) String() string {
	switch e {
	case KAudioObjectPropertyCreatorValue:
		return "KAudioObjectPropertyCreatorValue"
	case KAudioObjectPropertyListenerAdded:
		return "KAudioObjectPropertyListenerAdded"
	case KAudioObjectPropertyListenerRemoved:
		return "KAudioObjectPropertyListenerRemoved"
	default:
		return fmt.Sprintf("KAudioObjectPropertyCreator(%d)", e)
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

type KAudioObjectPropertyElement uint32

const (
	KAudioObjectPropertyElementWildcard KAudioObjectPropertyElement = 0xffffffff
)

func (e KAudioObjectPropertyElement) String() string {
	switch e {
	case KAudioObjectPropertyElementWildcard:
		return "KAudioObjectPropertyElementWildcard"
	default:
		return fmt.Sprintf("KAudioObjectPropertyElement(%d)", e)
	}
}

type KAudioObjectPropertyScope uint32

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

type KAudioObjectPropertyScopeGlobal uint32

const (
	KAudioObjectPropertyElementMain      KAudioObjectPropertyScopeGlobal = 0
	KAudioObjectPropertyScopeGlobalValue KAudioObjectPropertyScopeGlobal = 'g'<<24 | 'l'<<16 | 'o'<<8 | 'b' // 'glob'
	KAudioObjectPropertyScopeInput       KAudioObjectPropertyScopeGlobal = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	KAudioObjectPropertyScopeOutput      KAudioObjectPropertyScopeGlobal = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	KAudioObjectPropertyScopePlayThrough KAudioObjectPropertyScopeGlobal = 'p'<<24 | 't'<<16 | 'r'<<8 | 'u' // 'ptru'
	// Deprecated: use KAudioObjectPropertyElementMain.
	KAudioObjectPropertyElementMaster KAudioObjectPropertyScopeGlobal = 0
)

func (e KAudioObjectPropertyScopeGlobal) String() string {
	switch e {
	case KAudioObjectPropertyElementMain:
		return "KAudioObjectPropertyElementMain"
	case KAudioObjectPropertyScopeGlobalValue:
		return "KAudioObjectPropertyScopeGlobalValue"
	case KAudioObjectPropertyScopeInput:
		return "KAudioObjectPropertyScopeInput"
	case KAudioObjectPropertyScopeOutput:
		return "KAudioObjectPropertyScopeOutput"
	case KAudioObjectPropertyScopePlayThrough:
		return "KAudioObjectPropertyScopePlayThrough"
	default:
		return fmt.Sprintf("KAudioObjectPropertyScopeGlobal(%d)", e)
	}
}

type KAudioObjectPropertySelector uint32

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

type KAudioObjectSystem int32

const (
	KAudioObjectSystemObject KAudioObjectSystem = 1
)

func (e KAudioObjectSystem) String() string {
	switch e {
	case KAudioObjectSystemObject:
		return "KAudioObjectSystemObject"
	default:
		return fmt.Sprintf("KAudioObjectSystem(%d)", e)
	}
}

type KAudioPlugIn uint32

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

type KAudioPlugInClassI uint32

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

type KAudioPlugInProperty uint32

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

type KAudioProcessClassI uint32

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

type KAudioProcessProperty uint32

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

const KAudioPropertyWildcardChannel uint32 = 4294967295

type KAudioPropertyWildcardPropertyI uint32

const (
	KAudioPropertyWildcardPropertyID KAudioPropertyWildcardPropertyI = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e KAudioPropertyWildcardPropertyI) String() string {
	switch e {
	case KAudioPropertyWildcardPropertyID:
		return "KAudioPropertyWildcardPropertyID"
	default:
		return fmt.Sprintf("KAudioPropertyWildcardPropertyI(%d)", e)
	}
}

const KAudioPropertyWildcardSection uint8 = 0xff

type KAudioSelectorControlClassID uint32

const (
	KAudioClockSourceControlClassID     KAudioSelectorControlClassID = 'c'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'clck'
	KAudioDataDestinationControlClassID KAudioSelectorControlClassID = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KAudioDataSourceControlClassID      KAudioSelectorControlClassID = 'd'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'dsrc'
	KAudioHighPassFilterControlClassID  KAudioSelectorControlClassID = 'h'<<24 | 'i'<<16 | 'p'<<8 | 'f' // 'hipf'
	KAudioLineLevelControlClassID       KAudioSelectorControlClassID = 'n'<<24 | 'l'<<16 | 'v'<<8 | 'l' // 'nlvl'
	KAudioSelectorControlClassIDValue   KAudioSelectorControlClassID = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
)

func (e KAudioSelectorControlClassID) String() string {
	switch e {
	case KAudioClockSourceControlClassID:
		return "KAudioClockSourceControlClassID"
	case KAudioDataDestinationControlClassID:
		return "KAudioDataDestinationControlClassID"
	case KAudioDataSourceControlClassID:
		return "KAudioDataSourceControlClassID"
	case KAudioHighPassFilterControlClassID:
		return "KAudioHighPassFilterControlClassID"
	case KAudioLineLevelControlClassID:
		return "KAudioLineLevelControlClassID"
	case KAudioSelectorControlClassIDValue:
		return "KAudioSelectorControlClassIDValue"
	default:
		return fmt.Sprintf("KAudioSelectorControlClassID(%d)", e)
	}
}

type KAudioSelectorControlItemKind uint32

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

type KAudioSelectorControlProperty uint32

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

type KAudioSliderControlClassI uint32

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

type KAudioSliderControlProperty uint32

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

type KAudioStereoPanControlClassI uint32

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

type KAudioStereoPanControlProperty uint32

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

type KAudioStream uint32

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

type KAudioStreamClassI uint32

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

type KAudioStreamPropertyIsActive uint32

const (
	KAudioStreamPropertyAvailablePhysicalFormats KAudioStreamPropertyIsActive = 'p'<<24 | 'f'<<16 | 't'<<8 | 'a' // 'pfta'
	KAudioStreamPropertyAvailableVirtualFormats  KAudioStreamPropertyIsActive = 's'<<24 | 'f'<<16 | 'm'<<8 | 'a' // 'sfma'
	KAudioStreamPropertyDirection                KAudioStreamPropertyIsActive = 's'<<24 | 'd'<<16 | 'i'<<8 | 'r' // 'sdir'
	KAudioStreamPropertyIsActiveValue            KAudioStreamPropertyIsActive = 's'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'sact'
	KAudioStreamPropertyLatency                  KAudioStreamPropertyIsActive = 'l'<<24 | 't'<<16 | 'n'<<8 | 'c' // 'ltnc'
	KAudioStreamPropertyPhysicalFormat           KAudioStreamPropertyIsActive = 'p'<<24 | 'f'<<16 | 't'<<8 | ' ' // 'pft '
	KAudioStreamPropertyStartingChannel          KAudioStreamPropertyIsActive = 's'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'schn'
	KAudioStreamPropertyTerminalType             KAudioStreamPropertyIsActive = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
	KAudioStreamPropertyVirtualFormat            KAudioStreamPropertyIsActive = 's'<<24 | 'f'<<16 | 'm'<<8 | 't' // 'sfmt'
)

func (e KAudioStreamPropertyIsActive) String() string {
	switch e {
	case KAudioStreamPropertyAvailablePhysicalFormats:
		return "KAudioStreamPropertyAvailablePhysicalFormats"
	case KAudioStreamPropertyAvailableVirtualFormats:
		return "KAudioStreamPropertyAvailableVirtualFormats"
	case KAudioStreamPropertyDirection:
		return "KAudioStreamPropertyDirection"
	case KAudioStreamPropertyIsActiveValue:
		return "KAudioStreamPropertyIsActiveValue"
	case KAudioStreamPropertyLatency:
		return "KAudioStreamPropertyLatency"
	case KAudioStreamPropertyPhysicalFormat:
		return "KAudioStreamPropertyPhysicalFormat"
	case KAudioStreamPropertyStartingChannel:
		return "KAudioStreamPropertyStartingChannel"
	case KAudioStreamPropertyTerminalType:
		return "KAudioStreamPropertyTerminalType"
	case KAudioStreamPropertyVirtualFormat:
		return "KAudioStreamPropertyVirtualFormat"
	default:
		return fmt.Sprintf("KAudioStreamPropertyIsActive(%d)", e)
	}
}

type KAudioStreamPropertyOwningDevice uint32

const (
	KAudioStreamPropertyOwningDeviceValue       KAudioStreamPropertyOwningDevice = 's'<<24 | 't'<<16 | 'd'<<8 | 'v' // 'stdv'
	KAudioStreamPropertyPhysicalFormatMatch     KAudioStreamPropertyOwningDevice = 'p'<<24 | 'f'<<16 | 't'<<8 | 'm' // 'pftm'
	KAudioStreamPropertyPhysicalFormatSupported KAudioStreamPropertyOwningDevice = 'p'<<24 | 'f'<<16 | 't'<<8 | '?' // 'pft?'
	KAudioStreamPropertyPhysicalFormats         KAudioStreamPropertyOwningDevice = 'p'<<24 | 'f'<<16 | 't'<<8 | '#' // 'pft#'
)

func (e KAudioStreamPropertyOwningDevice) String() string {
	switch e {
	case KAudioStreamPropertyOwningDeviceValue:
		return "KAudioStreamPropertyOwningDeviceValue"
	case KAudioStreamPropertyPhysicalFormatMatch:
		return "KAudioStreamPropertyPhysicalFormatMatch"
	case KAudioStreamPropertyPhysicalFormatSupported:
		return "KAudioStreamPropertyPhysicalFormatSupported"
	case KAudioStreamPropertyPhysicalFormats:
		return "KAudioStreamPropertyPhysicalFormats"
	default:
		return fmt.Sprintf("KAudioStreamPropertyOwningDevice(%d)", e)
	}
}

type KAudioStreamTerminalType uint32

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

type KAudioSubDeviceClassI uint32

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

type KAudioSubDeviceDriftCompensation uint32

const (
	// Deprecated: use kAudioAggregateDriftCompensationHighQuality.
	KAudioSubDeviceDriftCompensationHighQuality KAudioSubDeviceDriftCompensation = 0x60
	// Deprecated: use kAudioAggregateDriftCompensationLowQuality.
	KAudioSubDeviceDriftCompensationLowQuality KAudioSubDeviceDriftCompensation = 0x20
	// Deprecated: use kAudioAggregateDriftCompensationMaxQuality.
	KAudioSubDeviceDriftCompensationMaxQuality KAudioSubDeviceDriftCompensation = 0x7f
	// Deprecated: use kAudioAggregateDriftCompensationMediumQuality.
	KAudioSubDeviceDriftCompensationMediumQuality KAudioSubDeviceDriftCompensation = 0x40
	// Deprecated: use kAudioAggregateDriftCompensationMinQuality.
	KAudioSubDeviceDriftCompensationMinQuality KAudioSubDeviceDriftCompensation = 0
)

func (e KAudioSubDeviceDriftCompensation) String() string {
	switch e {
	case KAudioSubDeviceDriftCompensationHighQuality:
		return "KAudioSubDeviceDriftCompensationHighQuality"
	case KAudioSubDeviceDriftCompensationLowQuality:
		return "KAudioSubDeviceDriftCompensationLowQuality"
	case KAudioSubDeviceDriftCompensationMaxQuality:
		return "KAudioSubDeviceDriftCompensationMaxQuality"
	case KAudioSubDeviceDriftCompensationMediumQuality:
		return "KAudioSubDeviceDriftCompensationMediumQuality"
	case KAudioSubDeviceDriftCompensationMinQuality:
		return "KAudioSubDeviceDriftCompensationMinQuality"
	default:
		return fmt.Sprintf("KAudioSubDeviceDriftCompensation(%d)", e)
	}
}

type KAudioSubDeviceProperty uint32

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

type KAudioSubTapClassI uint32

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

type KAudioSubTapProperty uint32

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

type KAudioSystemObjectClassI uint32

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

type KAudioTapClassI uint32

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

type KAudioTapProperty uint32

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

type KAudioTransportManager uint32

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

type KAudioTransportManagerClassI uint32

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

type KAudioTransportManagerProperty uint32

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
