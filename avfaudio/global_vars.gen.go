// Code generated from Apple documentation. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// AVAudioApplicationInputMuteStateChangeNotification is a notification the system posts when the app’s audio input mute state changes.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/inputMuteStateChangeNotification
	AVAudioApplicationInputMuteStateChangeNotification foundation.NSNotificationName
	// AVAudioUnitComponentManagerRegistrationsChangedNotification is a notification the component manager generates when it updates its list of components.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/registrationsChangedNotification
	AVAudioUnitComponentManagerRegistrationsChangedNotification foundation.NSNotificationName
	// AVSpeechSynthesisAvailableVoicesDidChangeNotification is a notification that indicates a change in available voices for speech synthesis.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/availableVoicesDidChangeNotification
	AVSpeechSynthesisAvailableVoicesDidChangeNotification foundation.NSNotificationName
)

var (
	// AVAudioApplicationMuteStateKey is a user information key to determine the app’s audio mute state.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/muteStateKey
	AVAudioApplicationMuteStateKey string
	// AVAudioBitRateStrategy_Constant is a value that represents a constant bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBitRateStrategy_Constant
	AVAudioBitRateStrategy_Constant string
	// AVAudioBitRateStrategy_LongTermAverage is a value that represents an average bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBitRateStrategy_LongTermAverage
	AVAudioBitRateStrategy_LongTermAverage string
	// AVAudioBitRateStrategy_Variable is a value that represents a variable bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBitRateStrategy_Variable
	AVAudioBitRateStrategy_Variable string
	// AVAudioBitRateStrategy_VariableConstrained is a value that represents a constrained variable bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBitRateStrategy_VariableConstrained
	AVAudioBitRateStrategy_VariableConstrained string
	// AVAudioEngineConfigurationChangeNotification is a notification the framework posts when the audio engine configuration changes.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineConfigurationChangeNotification
	AVAudioEngineConfigurationChangeNotification string
	// AVAudioFileTypeKey is a string that indicates the audio file type.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFileTypeKey
	AVAudioFileTypeKey string
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mutestatekey
	AVAudioSessionMuteStateKey foundation.NSString
	// AVAudioUnitComponentTagsDidChangeNotification is a notification that indicates when component tags change.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentTagsDidChangeNotification
	AVAudioUnitComponentTagsDidChangeNotification string
	// AVAudioUnitManufacturerNameApple is the audio unit manufacturer is Apple.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitManufacturerNameApple
	AVAudioUnitManufacturerNameApple string
	// AVAudioUnitTypeEffect is an audio unit type that represents an effect.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeEffect
	AVAudioUnitTypeEffect string
	// AVAudioUnitTypeFormatConverter is an audio unit type that represents a format converter.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeFormatConverter
	AVAudioUnitTypeFormatConverter string
	// AVAudioUnitTypeGenerator is an audio unit type that represents a generator.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeGenerator
	AVAudioUnitTypeGenerator string
	// AVAudioUnitTypeMIDIProcessor is an audio unit type that represents a MIDI processor.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeMIDIProcessor
	AVAudioUnitTypeMIDIProcessor string
	// AVAudioUnitTypeMixer is an audio unit type that represents a mixer.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeMixer
	AVAudioUnitTypeMixer string
	// AVAudioUnitTypeMusicDevice is an audio unit type that represents a music device.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeMusicDevice
	AVAudioUnitTypeMusicDevice string
	// AVAudioUnitTypeMusicEffect is an audio unit type that represents a music effect.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeMusicEffect
	AVAudioUnitTypeMusicEffect string
	// AVAudioUnitTypeOfflineEffect is an audio unit type that represents an offline effect.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeOfflineEffect
	AVAudioUnitTypeOfflineEffect string
	// AVAudioUnitTypeOutput is an audio unit type that represents an output.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypeOutput
	AVAudioUnitTypeOutput string
	// AVAudioUnitTypePanner is an audio unit type that represents a panner.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTypePanner
	AVAudioUnitTypePanner string
	// See: https://developer.apple.com/documentation/AVFAudio/AVChannelLayoutKey
	AVChannelLayoutKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderASPFrequencyKey
	AVEncoderASPFrequencyKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderAudioQualityForVBRKey
	AVEncoderAudioQualityForVBRKey string
	// AVEncoderAudioQualityKey is a constant that represents an integer from the audio quality enumeration.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderAudioQualityKey
	AVEncoderAudioQualityKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitDepthHintKey
	AVEncoderBitDepthHintKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitRateKey
	AVEncoderBitRateKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitRatePerChannelKey
	AVEncoderBitRatePerChannelKey string
	// AVEncoderBitRateStrategyKey is a constant that represents the bit rate strategy for the encoder to use.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitRateStrategyKey
	AVEncoderBitRateStrategyKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderContentSourceKey
	AVEncoderContentSourceKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVEncoderDynamicRangeControlConfigurationKey
	AVEncoderDynamicRangeControlConfigurationKey string
	// AVFormatIDKey is an integer value that represents the format of the audio data.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVFormatIDKey
	AVFormatIDKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVLinearPCMBitDepthKey
	AVLinearPCMBitDepthKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVLinearPCMIsBigEndianKey
	AVLinearPCMIsBigEndianKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVLinearPCMIsFloatKey
	AVLinearPCMIsFloatKey string
	// See: https://developer.apple.com/documentation/AVFAudio/AVLinearPCMIsNonInterleaved
	AVLinearPCMIsNonInterleaved string
	// AVNumberOfChannelsKey is an integer value that represents the number of channels.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
	AVNumberOfChannelsKey string
	// AVSampleRateConverterAlgorithmKey is a string value that represents the sample rate converter algorithm to use.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAlgorithmKey
	AVSampleRateConverterAlgorithmKey string
	// AVSampleRateConverterAlgorithm_Mastering is the mastering encoder bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAlgorithm_Mastering
	AVSampleRateConverterAlgorithm_Mastering string
	// AVSampleRateConverterAlgorithm_MinimumPhase is the minimum phase encoder bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAlgorithm_MinimumPhase
	AVSampleRateConverterAlgorithm_MinimumPhase string
	// AVSampleRateConverterAlgorithm_Normal is the usual encoder bit rate strategy.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAlgorithm_Normal
	AVSampleRateConverterAlgorithm_Normal string
	// AVSampleRateConverterAudioQualityKey is an integer value that represents the audio quality for conversion.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAudioQualityKey
	AVSampleRateConverterAudioQualityKey string
	// AVSampleRateKey is a floating point value that represents the sample rate, in hertz.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateKey
	AVSampleRateKey string
	// AVSpeechSynthesisIPANotationAttribute is a string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisIPANotationAttribute
	AVSpeechSynthesisIPANotationAttribute string
	// AVSpeechSynthesisVoiceIdentifierAlex is the voice that the system identifies as Alex.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceIdentifierAlex
	AVSpeechSynthesisVoiceIdentifierAlex string
)

var ()

var (
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/availableinputschangenotification
	AVAudioSessionAvailableInputsChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/interruptionnotification
	AVAudioSessionInterruptionNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mediaserviceswerelostnotification
	AVAudioSessionMediaServicesWereLostNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mediaserviceswereresetnotification
	AVAudioSessionMediaServicesWereResetNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/microphoneinjectioncapabilitieschangenotification
	AVAudioSessionMicrophoneInjectionCapabilitiesChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputmutestatechangenotification
	AVAudioSessionOutputMuteStateChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/renderingcapabilitieschangenotification
	AVAudioSessionRenderingCapabilitiesChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/renderingmodechangenotification
	AVAudioSessionRenderingModeChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/routechangenotification
	AVAudioSessionRouteChangeNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/silencesecondaryaudiohintnotification
	AVAudioSessionSilenceSecondaryAudioHintNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/spatialplaybackcapabilitieschangednotification
	AVAudioSessionSpatialPlaybackCapabilitiesChangedNotification foundation.NSNotification
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/userintenttounmuteoutputnotification
	AVAudioSessionUserIntentToUnmuteOutputNotification foundation.NSNotification
)

var (
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/ambient
	AVAudioSessionCategoryAmbient objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/audioprocessing
	AVAudioSessionCategoryAudioProcessing objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/multiroute
	AVAudioSessionCategoryMultiRoute objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/playandrecord
	AVAudioSessionCategoryPlayAndRecord objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/playback
	AVAudioSessionCategoryPlayback objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/record
	AVAudioSessionCategoryRecord objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/soloambient
	AVAudioSessionCategorySoloAmbient objc.ID
	// AVAudioSessionLocationLower is a value that indicates that the data source is located near the bottom end of the device.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Location/lower
	AVAudioSessionLocationLower objc.ID
	// AVAudioSessionLocationUpper is a value that indicates that the data source is located near the top end of the device.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Location/upper
	AVAudioSessionLocationUpper objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/default
	AVAudioSessionModeDefault objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/dualroute
	AVAudioSessionModeDualRoute objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/gamechat
	AVAudioSessionModeGameChat objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/measurement
	AVAudioSessionModeMeasurement objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/movieplayback
	AVAudioSessionModeMoviePlayback objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/shortformvideo
	AVAudioSessionModeShortFormVideo objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/spokenaudio
	AVAudioSessionModeSpokenAudio objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/videochat
	AVAudioSessionModeVideoChat objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/videorecording
	AVAudioSessionModeVideoRecording objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/voicechat
	AVAudioSessionModeVoiceChat objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/mode-swift.struct/voiceprompt
	AVAudioSessionModeVoicePrompt objc.ID
	// AVAudioSessionOrientationBack is a data source that points outward from the back of the device, away from the user.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/back
	AVAudioSessionOrientationBack objc.ID
	// AVAudioSessionOrientationBottom is a data source that points downward.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/bottom
	AVAudioSessionOrientationBottom objc.ID
	// AVAudioSessionOrientationFront is a data source that points outward from the front of the device, toward the user.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/front
	AVAudioSessionOrientationFront objc.ID
	// AVAudioSessionOrientationLeft is a data source that points outward to the left of the device, away from the user.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/left
	AVAudioSessionOrientationLeft objc.ID
	// AVAudioSessionOrientationRight is a data source that points outward to the right of the device, away from the user.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/right
	AVAudioSessionOrientationRight objc.ID
	// AVAudioSessionOrientationTop is a data source that points upward.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Orientation/top
	AVAudioSessionOrientationTop objc.ID
	// AVAudioSessionPolarPatternCardioid is a data source that’s most sensitive to sound from the direction of the data source and is nearly insensitive to sound from the opposite direction.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PolarPattern/cardioid
	AVAudioSessionPolarPatternCardioid objc.ID
	// AVAudioSessionPolarPatternOmnidirectional is a data source that’s equally sensitive to sound from any direction.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PolarPattern/omnidirectional
	AVAudioSessionPolarPatternOmnidirectional objc.ID
	// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/polarpattern/stereo
	AVAudioSessionPolarPatternStereo objc.ID
	// AVAudioSessionPolarPatternSubcardioid is a data source that’s most sensitive to sound from the direction of the data source and is less sensitive to sound from the opposite direction.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PolarPattern/subcardioid
	AVAudioSessionPolarPatternSubcardioid objc.ID
)

var (
	// AVExtendedNoteOnEventDefaultInstrument is a constant that represents the default instrument identifier.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/defaultInstrument
	AVExtendedNoteOnEventDefaultInstrument uint32
)

var (
	// AVSpeechUtteranceDefaultSpeechRate is the default rate the speech synthesizer uses when speaking an utterance.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceDefaultSpeechRate
	AVSpeechUtteranceDefaultSpeechRate float32
	// AVSpeechUtteranceMaximumSpeechRate is the maximum rate the speech synthesizer uses when speaking an utterance.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceMaximumSpeechRate
	AVSpeechUtteranceMaximumSpeechRate float32
	// AVSpeechUtteranceMinimumSpeechRate is the minimum rate the speech synthesizer uses when speaking an utterance.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceMinimumSpeechRate
	AVSpeechUtteranceMinimumSpeechRate float32
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioApplicationInputMuteStateChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioApplicationInputMuteStateChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioApplicationMuteStateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioApplicationMuteStateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioBitRateStrategy_Constant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioBitRateStrategy_Constant = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioBitRateStrategy_LongTermAverage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioBitRateStrategy_LongTermAverage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioBitRateStrategy_Variable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioBitRateStrategy_Variable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioBitRateStrategy_VariableConstrained"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioBitRateStrategy_VariableConstrained = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioEngineConfigurationChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioEngineConfigurationChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioFileTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioFileTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Album = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyApproximateDurationInSeconds"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.ApproximateDurationInSeconds = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Artist = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyChannelLayout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.ChannelLayout = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyComments"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Comments = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Composer = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Copyright = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyEncodingApplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.EncodingApplication = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Genre = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyISRC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.ISRC = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyKeySignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.KeySignature = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyLyricist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Lyricist = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyNominalBitRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.NominalBitRate = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyRecordedDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.RecordedDate = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeySourceBitDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.SourceBitDepth = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeySourceEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.SourceEncoder = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeySubTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.SubTitle = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyTempo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Tempo = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyTimeSignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.TimeSignature = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Title = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyTrackNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.TrackNumber = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSequencerInfoDictionaryKeyYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioSequencerInfoDictionaryKeys.Year = AVAudioSequencerInfoDictionaryKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionAvailableInputsChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionAvailableInputsChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryAmbient"); err == nil && ptr != 0 {
		AVAudioSessionCategoryAmbient = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryAudioProcessing"); err == nil && ptr != 0 {
		AVAudioSessionCategoryAudioProcessing = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryMultiRoute"); err == nil && ptr != 0 {
		AVAudioSessionCategoryMultiRoute = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryPlayAndRecord"); err == nil && ptr != 0 {
		AVAudioSessionCategoryPlayAndRecord = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryPlayback"); err == nil && ptr != 0 {
		AVAudioSessionCategoryPlayback = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategoryRecord"); err == nil && ptr != 0 {
		AVAudioSessionCategoryRecord = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionCategorySoloAmbient"); err == nil && ptr != 0 {
		AVAudioSessionCategorySoloAmbient = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionInterruptionNotification"); err == nil && ptr != 0 {
		AVAudioSessionInterruptionNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionLocationLower"); err == nil && ptr != 0 {
		AVAudioSessionLocationLower = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionLocationUpper"); err == nil && ptr != 0 {
		AVAudioSessionLocationUpper = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionMediaServicesWereLostNotification"); err == nil && ptr != 0 {
		AVAudioSessionMediaServicesWereLostNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionMediaServicesWereResetNotification"); err == nil && ptr != 0 {
		AVAudioSessionMediaServicesWereResetNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionMicrophoneInjectionCapabilitiesChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionMicrophoneInjectionCapabilitiesChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeDefault"); err == nil && ptr != 0 {
		AVAudioSessionModeDefault = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeDualRoute"); err == nil && ptr != 0 {
		AVAudioSessionModeDualRoute = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeGameChat"); err == nil && ptr != 0 {
		AVAudioSessionModeGameChat = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeMeasurement"); err == nil && ptr != 0 {
		AVAudioSessionModeMeasurement = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeMoviePlayback"); err == nil && ptr != 0 {
		AVAudioSessionModeMoviePlayback = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeShortFormVideo"); err == nil && ptr != 0 {
		AVAudioSessionModeShortFormVideo = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeSpokenAudio"); err == nil && ptr != 0 {
		AVAudioSessionModeSpokenAudio = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeVideoChat"); err == nil && ptr != 0 {
		AVAudioSessionModeVideoChat = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeVideoRecording"); err == nil && ptr != 0 {
		AVAudioSessionModeVideoRecording = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeVoiceChat"); err == nil && ptr != 0 {
		AVAudioSessionModeVoiceChat = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionModeVoicePrompt"); err == nil && ptr != 0 {
		AVAudioSessionModeVoicePrompt = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionMuteStateKey"); err == nil && ptr != 0 {
		AVAudioSessionMuteStateKey = *(*foundation.NSString)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationBack"); err == nil && ptr != 0 {
		AVAudioSessionOrientationBack = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationBottom"); err == nil && ptr != 0 {
		AVAudioSessionOrientationBottom = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationFront"); err == nil && ptr != 0 {
		AVAudioSessionOrientationFront = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationLeft"); err == nil && ptr != 0 {
		AVAudioSessionOrientationLeft = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationRight"); err == nil && ptr != 0 {
		AVAudioSessionOrientationRight = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOrientationTop"); err == nil && ptr != 0 {
		AVAudioSessionOrientationTop = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionOutputMuteStateChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionOutputMuteStateChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionPolarPatternCardioid"); err == nil && ptr != 0 {
		AVAudioSessionPolarPatternCardioid = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionPolarPatternOmnidirectional"); err == nil && ptr != 0 {
		AVAudioSessionPolarPatternOmnidirectional = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionPolarPatternStereo"); err == nil && ptr != 0 {
		AVAudioSessionPolarPatternStereo = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionPolarPatternSubcardioid"); err == nil && ptr != 0 {
		AVAudioSessionPolarPatternSubcardioid = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionRenderingCapabilitiesChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionRenderingCapabilitiesChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionRenderingModeChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionRenderingModeChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionRouteChangeNotification"); err == nil && ptr != 0 {
		AVAudioSessionRouteChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionSilenceSecondaryAudioHintNotification"); err == nil && ptr != 0 {
		AVAudioSessionSilenceSecondaryAudioHintNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionSpatialPlaybackCapabilitiesChangedNotification"); err == nil && ptr != 0 {
		AVAudioSessionSpatialPlaybackCapabilitiesChangedNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioSessionUserIntentToUnmuteOutputNotification"); err == nil && ptr != 0 {
		AVAudioSessionUserIntentToUnmuteOutputNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitComponentManagerRegistrationsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitComponentManagerRegistrationsChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitComponentTagsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitComponentTagsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitManufacturerNameApple"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitManufacturerNameApple = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeFormatConverter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeFormatConverter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeGenerator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeGenerator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeMIDIProcessor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeMIDIProcessor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeMixer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeMixer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeMusicDevice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeMusicDevice = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeMusicEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeMusicEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeOfflineEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeOfflineEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypeOutput"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypeOutput = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioUnitTypePanner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioUnitTypePanner = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVChannelLayoutKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVChannelLayoutKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderASPFrequencyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderASPFrequencyKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderAudioQualityForVBRKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderAudioQualityForVBRKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderAudioQualityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderAudioQualityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderBitDepthHintKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderBitDepthHintKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderBitRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderBitRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderBitRatePerChannelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderBitRatePerChannelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderBitRateStrategyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderBitRateStrategyKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderContentSourceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderContentSourceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVEncoderDynamicRangeControlConfigurationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVEncoderDynamicRangeControlConfigurationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVExtendedNoteOnEventDefaultInstrument"); err == nil && ptr != 0 {
		AVExtendedNoteOnEventDefaultInstrument = *(*uint32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFormatIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFormatIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLinearPCMBitDepthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLinearPCMBitDepthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLinearPCMIsBigEndianKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLinearPCMIsBigEndianKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLinearPCMIsFloatKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLinearPCMIsFloatKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLinearPCMIsNonInterleaved"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLinearPCMIsNonInterleaved = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVNumberOfChannelsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVNumberOfChannelsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateConverterAlgorithmKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateConverterAlgorithmKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateConverterAlgorithm_Mastering"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateConverterAlgorithm_Mastering = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateConverterAlgorithm_MinimumPhase"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateConverterAlgorithm_MinimumPhase = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateConverterAlgorithm_Normal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateConverterAlgorithm_Normal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateConverterAudioQualityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateConverterAudioQualityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechSynthesisAvailableVoicesDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSpeechSynthesisAvailableVoicesDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechSynthesisIPANotationAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSpeechSynthesisIPANotationAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechSynthesisVoiceIdentifierAlex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSpeechSynthesisVoiceIdentifierAlex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechUtteranceDefaultSpeechRate"); err == nil && ptr != 0 {
		AVSpeechUtteranceDefaultSpeechRate = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechUtteranceMaximumSpeechRate"); err == nil && ptr != 0 {
		AVSpeechUtteranceMaximumSpeechRate = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpeechUtteranceMinimumSpeechRate"); err == nil && ptr != 0 {
		AVSpeechUtteranceMinimumSpeechRate = *(*float32)(unsafe.Pointer(ptr))
	}

}

// AVAudioSequencerInfoDictionaryKeys provides typed accessors for [AVAudioSequencerInfoDictionaryKey] constants.
var AVAudioSequencerInfoDictionaryKeys struct {
	// Album: A key that represents the album.
	Album AVAudioSequencerInfoDictionaryKey
	// ApproximateDurationInSeconds: A key that represents the approximate duration.
	ApproximateDurationInSeconds AVAudioSequencerInfoDictionaryKey
	// Artist: A key that represents the artist.
	Artist AVAudioSequencerInfoDictionaryKey
	// ChannelLayout: A key that represents the channel layout.
	ChannelLayout AVAudioSequencerInfoDictionaryKey
	// Comments: A key that represents the comments.
	Comments AVAudioSequencerInfoDictionaryKey
	// Composer: A key that represents the composer.
	Composer AVAudioSequencerInfoDictionaryKey
	// Copyright: A key that represents the copyright statement.
	Copyright AVAudioSequencerInfoDictionaryKey
	// EncodingApplication: A key that represents the encoding application.
	EncodingApplication AVAudioSequencerInfoDictionaryKey
	// Genre: A key that represents the genre.
	Genre AVAudioSequencerInfoDictionaryKey
	// ISRC: A key that represents the international standard recording code.
	ISRC AVAudioSequencerInfoDictionaryKey
	// KeySignature: A key that represents the key signature.
	KeySignature AVAudioSequencerInfoDictionaryKey
	// Lyricist: A key that represents the lyricist.
	Lyricist AVAudioSequencerInfoDictionaryKey
	// NominalBitRate: A key that represents the nominal bit rate.
	NominalBitRate AVAudioSequencerInfoDictionaryKey
	// RecordedDate: A key that represents the date of the recording.
	RecordedDate AVAudioSequencerInfoDictionaryKey
	// SourceBitDepth: A key that represents the bit depth of the source.
	SourceBitDepth AVAudioSequencerInfoDictionaryKey
	// SourceEncoder: A key that represents the encoder the source uses.
	SourceEncoder AVAudioSequencerInfoDictionaryKey
	// SubTitle: A key that represents the subtitle.
	SubTitle AVAudioSequencerInfoDictionaryKey
	// Tempo: A key that represents the tempo.
	Tempo AVAudioSequencerInfoDictionaryKey
	// TimeSignature: A key that represents the time signature.
	TimeSignature AVAudioSequencerInfoDictionaryKey
	// Title: A key that represents the title.
	Title AVAudioSequencerInfoDictionaryKey
	// TrackNumber: A key that represents the track number.
	TrackNumber AVAudioSequencerInfoDictionaryKey
	// Year: A key that represents the year.
	Year AVAudioSequencerInfoDictionaryKey
}
