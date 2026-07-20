// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode
type AVAudioSessionErrorCode int

const (
	// AVAudioSessionErrorCodeBadParam: An error code that indicates an attempt to set a property to an illegal value.
	AVAudioSessionErrorCodeBadParam AVAudioSessionErrorCode = -50
	// AVAudioSessionErrorCodeCannotInterruptOthers: An error code that indictates an attempt to make a nonmixable audio session active while the app was in the background.
	AVAudioSessionErrorCodeCannotInterruptOthers AVAudioSessionErrorCode = '!'<<24 | 'i'<<16 | 'n'<<8 | 't' // '!int'
	// AVAudioSessionErrorCodeCannotStartPlaying: An error code that indicates an attempt to start audio playback when it wasn’t allowed.
	AVAudioSessionErrorCodeCannotStartPlaying AVAudioSessionErrorCode = '!'<<24 | 'p'<<16 | 'l'<<8 | 'a' // '!pla'
	// AVAudioSessionErrorCodeCannotStartRecording: An error code that indicates an attempt to start audio recording, but the operation failed.
	AVAudioSessionErrorCodeCannotStartRecording AVAudioSessionErrorCode = '!'<<24 | 'r'<<16 | 'e'<<8 | 'c' // '!rec'
	// AVAudioSessionErrorCodeExpiredSession: An error code that indicates that an operation failed because the system deallocated the associated session.
	AVAudioSessionErrorCodeExpiredSession AVAudioSessionErrorCode = '!'<<24 | 's'<<16 | 'e'<<8 | 's' // '!ses'
	// AVAudioSessionErrorCodeIncompatibleCategory: An error code that indicates an attempt to perform an operation that the current audio session category doesn’t support.
	AVAudioSessionErrorCodeIncompatibleCategory AVAudioSessionErrorCode = '!'<<24 | 'c'<<16 | 'a'<<8 | 't' // '!cat'
	// AVAudioSessionErrorCodeInsufficientPriority: An error code that indicates the app isn’t allowed to set the audio category because it’s in use by another app.
	AVAudioSessionErrorCodeInsufficientPriority AVAudioSessionErrorCode = '!'<<24 | 'p'<<16 | 'r'<<8 | 'i' // '!pri'
	// AVAudioSessionErrorCodeIsBusy: An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	AVAudioSessionErrorCodeIsBusy AVAudioSessionErrorCode = '!'<<24 | 'a'<<16 | 'c'<<8 | 't' // '!act'
	// AVAudioSessionErrorCodeMediaServicesFailed: An error code that indictates an attempt to use the audio session during or after a Media Services failure.
	AVAudioSessionErrorCodeMediaServicesFailed AVAudioSessionErrorCode = 'm'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'msrv'
	// AVAudioSessionErrorCodeMissingEntitlement: An error code that indicates an attempt to perform an operation for which the app doesn’t have the required entitlements.
	AVAudioSessionErrorCodeMissingEntitlement AVAudioSessionErrorCode = 'e'<<24 | 'n'<<16 | 't'<<8 | '?' // 'ent?'
	// AVAudioSessionErrorCodeNone: An error code that indicates the operation succeeded.
	AVAudioSessionErrorCodeNone AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeResourceNotAvailable: An error code that indicates that an operation failed because the device doesn’t have sufficient hardware resources to complete the action.
	AVAudioSessionErrorCodeResourceNotAvailable AVAudioSessionErrorCode = '!'<<24 | 'r'<<16 | 'e'<<8 | 's' // '!res'
	// AVAudioSessionErrorCodeSessionNotActive: An error code that indicates the operation failed because the session isn’t active.
	AVAudioSessionErrorCodeSessionNotActive AVAudioSessionErrorCode = 'i'<<24 | 'n'<<16 | 'a'<<8 | 'c' // 'inac'
	// AVAudioSessionErrorCodeSiriIsRecording: An error code that indicates an attempt to perform an operation that isn’t allowed while Siri is recording.
	AVAudioSessionErrorCodeSiriIsRecording AVAudioSessionErrorCode = 's'<<24 | 'i'<<16 | 'r'<<8 | 'i' // 'siri'
	// AVAudioSessionErrorCodeUnspecified: An error code that indicates an unspecified error occurred.
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

type AVAudioSessionErrorInsufficient int

const (
	// Deprecated: use AVAudioSession.ErrorCode.insufficientPriority.
	AVAudioSessionErrorInsufficientPriority AVAudioSessionErrorInsufficient = '!'<<24 | 'p'<<16 | 'r'<<8 | 'i' // '!pri'
)

func (e AVAudioSessionErrorInsufficient) String() string {
	switch e {
	case AVAudioSessionErrorInsufficientPriority:
		return "AVAudioSessionErrorInsufficientPriority"
	default:
		return fmt.Sprintf("AVAudioSessionErrorInsufficient(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap
type AudioChannelBitmap uint32

const (
	// KAudioChannelBit_Center: The center channel.
	KAudioChannelBit_Center AudioChannelBitmap = 4
	// KAudioChannelBit_CenterSurround: The center surround channel.
	KAudioChannelBit_CenterSurround AudioChannelBitmap = 256
	// KAudioChannelBit_CenterTopFront: The top-front center channel.
	KAudioChannelBit_CenterTopFront AudioChannelBitmap = 8192
	// KAudioChannelBit_CenterTopMiddle: The top-middle center channel.
	KAudioChannelBit_CenterTopMiddle AudioChannelBitmap = 2048
	// KAudioChannelBit_CenterTopRear: The top-right center channel.
	KAudioChannelBit_CenterTopRear AudioChannelBitmap = 33554432
	// KAudioChannelBit_LFEScreen: The Low Frequency Effects (LFE) screen channel.
	KAudioChannelBit_LFEScreen AudioChannelBitmap = 8
	// KAudioChannelBit_Left: The left channel.
	KAudioChannelBit_Left AudioChannelBitmap = 1
	// KAudioChannelBit_LeftCenter: The left center channel.
	KAudioChannelBit_LeftCenter AudioChannelBitmap = 64
	// KAudioChannelBit_LeftSurround: The left surround channel.
	KAudioChannelBit_LeftSurround AudioChannelBitmap = 16
	// KAudioChannelBit_LeftSurroundDirect: The left surround direct channel.
	KAudioChannelBit_LeftSurroundDirect AudioChannelBitmap = 512
	// KAudioChannelBit_LeftTopFront: The left-top front channel.
	KAudioChannelBit_LeftTopFront AudioChannelBitmap = 4096
	// KAudioChannelBit_LeftTopMiddle: The left-top middle channel.
	KAudioChannelBit_LeftTopMiddle AudioChannelBitmap = 2097152
	// KAudioChannelBit_LeftTopRear: The left-top rear channel.
	KAudioChannelBit_LeftTopRear AudioChannelBitmap = 16777216
	// KAudioChannelBit_Right: The right channel.
	KAudioChannelBit_Right AudioChannelBitmap = 2
	// KAudioChannelBit_RightCenter: The right center channel.
	KAudioChannelBit_RightCenter AudioChannelBitmap = 128
	// KAudioChannelBit_RightSurround: The rIght surround channel.
	KAudioChannelBit_RightSurround AudioChannelBitmap = 32
	// KAudioChannelBit_RightSurroundDirect: The right surround direct channel.
	KAudioChannelBit_RightSurroundDirect AudioChannelBitmap = 1024
	// KAudioChannelBit_RightTopFront: The top-front front channel.
	KAudioChannelBit_RightTopFront AudioChannelBitmap = 16384
	// KAudioChannelBit_RightTopMiddle: The top-middle right channel.
	KAudioChannelBit_RightTopMiddle AudioChannelBitmap = 8388608
	// KAudioChannelBit_RightTopRear: The top-rear right channel.
	KAudioChannelBit_RightTopRear AudioChannelBitmap = 67108864
	// KAudioChannelBit_TopBackCenter: The top-back center channel.
	KAudioChannelBit_TopBackCenter AudioChannelBitmap = 65536
	// KAudioChannelBit_TopBackLeft: The top-back left channel.
	KAudioChannelBit_TopBackLeft AudioChannelBitmap = 32768
	// KAudioChannelBit_TopBackRight: The top-back right channel.
	KAudioChannelBit_TopBackRight AudioChannelBitmap = 131072
	// KAudioChannelBit_TopCenterSurround: The top center surround channel.
	KAudioChannelBit_TopCenterSurround AudioChannelBitmap = 2048
	// KAudioChannelBit_VerticalHeightCenter: The vertical height center channel.
	KAudioChannelBit_VerticalHeightCenter AudioChannelBitmap = 8192
	// KAudioChannelBit_VerticalHeightLeft: The vertical height left channel.
	KAudioChannelBit_VerticalHeightLeft AudioChannelBitmap = 4096
	// KAudioChannelBit_VerticalHeightRight: The vertical height right channel.
	KAudioChannelBit_VerticalHeightRight AudioChannelBitmap = 16384
)

func (e AudioChannelBitmap) String() string {
	switch e {
	case KAudioChannelBit_Center:
		return "KAudioChannelBit_Center"
	case KAudioChannelBit_CenterSurround:
		return "KAudioChannelBit_CenterSurround"
	case KAudioChannelBit_CenterTopFront:
		return "KAudioChannelBit_CenterTopFront"
	case KAudioChannelBit_CenterTopMiddle:
		return "KAudioChannelBit_CenterTopMiddle"
	case KAudioChannelBit_CenterTopRear:
		return "KAudioChannelBit_CenterTopRear"
	case KAudioChannelBit_LFEScreen:
		return "KAudioChannelBit_LFEScreen"
	case KAudioChannelBit_Left:
		return "KAudioChannelBit_Left"
	case KAudioChannelBit_LeftCenter:
		return "KAudioChannelBit_LeftCenter"
	case KAudioChannelBit_LeftSurround:
		return "KAudioChannelBit_LeftSurround"
	case KAudioChannelBit_LeftSurroundDirect:
		return "KAudioChannelBit_LeftSurroundDirect"
	case KAudioChannelBit_LeftTopFront:
		return "KAudioChannelBit_LeftTopFront"
	case KAudioChannelBit_LeftTopMiddle:
		return "KAudioChannelBit_LeftTopMiddle"
	case KAudioChannelBit_LeftTopRear:
		return "KAudioChannelBit_LeftTopRear"
	case KAudioChannelBit_Right:
		return "KAudioChannelBit_Right"
	case KAudioChannelBit_RightCenter:
		return "KAudioChannelBit_RightCenter"
	case KAudioChannelBit_RightSurround:
		return "KAudioChannelBit_RightSurround"
	case KAudioChannelBit_RightSurroundDirect:
		return "KAudioChannelBit_RightSurroundDirect"
	case KAudioChannelBit_RightTopFront:
		return "KAudioChannelBit_RightTopFront"
	case KAudioChannelBit_RightTopMiddle:
		return "KAudioChannelBit_RightTopMiddle"
	case KAudioChannelBit_RightTopRear:
		return "KAudioChannelBit_RightTopRear"
	case KAudioChannelBit_TopBackCenter:
		return "KAudioChannelBit_TopBackCenter"
	case KAudioChannelBit_TopBackLeft:
		return "KAudioChannelBit_TopBackLeft"
	case KAudioChannelBit_TopBackRight:
		return "KAudioChannelBit_TopBackRight"
	default:
		return fmt.Sprintf("AudioChannelBitmap(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex
type AudioChannelCoordinateIndex uint32

const (
	// KAudioChannelCoordinates_Azimuth: For spherical coordinates, `0` is front center, positive is right, negative is left, and measurements are in degrees.
	KAudioChannelCoordinates_Azimuth AudioChannelCoordinateIndex = 0
	// KAudioChannelCoordinates_BackFront: For rectangular coordinates, negative is back and positive is front.
	KAudioChannelCoordinates_BackFront AudioChannelCoordinateIndex = 1
	// KAudioChannelCoordinates_Distance: For spherical coordinates, distance is radially from the center.
	KAudioChannelCoordinates_Distance AudioChannelCoordinateIndex = 2
	// KAudioChannelCoordinates_DownUp: For rectangular coordinates, negative is below ground level, `0` is ground level, and positive is above ground level.
	KAudioChannelCoordinates_DownUp AudioChannelCoordinateIndex = 2
	// KAudioChannelCoordinates_Elevation: For spherical coordinates, `+90` is zenith, `0` is horizontal, `-90` is nadir, and measurements are in degrees.
	KAudioChannelCoordinates_Elevation AudioChannelCoordinateIndex = 1
	// KAudioChannelCoordinates_LeftRight: For rectangular coordinates, negative is left and positive is right.
	KAudioChannelCoordinates_LeftRight AudioChannelCoordinateIndex = 0
)

func (e AudioChannelCoordinateIndex) String() string {
	switch e {
	case KAudioChannelCoordinates_Azimuth:
		return "KAudioChannelCoordinates_Azimuth"
	case KAudioChannelCoordinates_BackFront:
		return "KAudioChannelCoordinates_BackFront"
	case KAudioChannelCoordinates_Distance:
		return "KAudioChannelCoordinates_Distance"
	default:
		return fmt.Sprintf("AudioChannelCoordinateIndex(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags
type AudioChannelFlags uint32

const (
	// KAudioChannelFlags_AllOff: All flags are clear.
	KAudioChannelFlags_AllOff AudioChannelFlags = 0
	// KAudioChannelFlags_Meters: A flag that indicates that unit values are in meters.
	KAudioChannelFlags_Meters AudioChannelFlags = 4
	// KAudioChannelFlags_RectangularCoordinates: A flag that indicates the channel uses the speaker position’s cartesian coordinates.
	KAudioChannelFlags_RectangularCoordinates AudioChannelFlags = 1
	// KAudioChannelFlags_SphericalCoordinates: A flag that indicates the channel uses the speaker position’s spherical coordinates.
	KAudioChannelFlags_SphericalCoordinates AudioChannelFlags = 2
)

func (e AudioChannelFlags) String() string {
	switch e {
	case KAudioChannelFlags_AllOff:
		return "KAudioChannelFlags_AllOff"
	case KAudioChannelFlags_Meters:
		return "KAudioChannelFlags_Meters"
	case KAudioChannelFlags_RectangularCoordinates:
		return "KAudioChannelFlags_RectangularCoordinates"
	case KAudioChannelFlags_SphericalCoordinates:
		return "KAudioChannelFlags_SphericalCoordinates"
	default:
		return fmt.Sprintf("AudioChannelFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags
type AudioTimeStampFlags uint32

const (
	// KAudioTimeStampHostTimeValid: A flag that indicates that the host time is valid.
	KAudioTimeStampHostTimeValid AudioTimeStampFlags = 2
	// KAudioTimeStampNothingValid: A flag that indicates no fields are valid.
	KAudioTimeStampNothingValid AudioTimeStampFlags = 0
	// KAudioTimeStampRateScalarValid: A flag that indicates that the rate scalar is valid.
	KAudioTimeStampRateScalarValid AudioTimeStampFlags = 4
	// KAudioTimeStampSMPTETimeValid: A flag that indicates that the SMPTE time is valid.
	KAudioTimeStampSMPTETimeValid AudioTimeStampFlags = 16
	// KAudioTimeStampSampleHostTimeValid: A flag that indicates that the sample frame time and the host time are valid.
	KAudioTimeStampSampleHostTimeValid AudioTimeStampFlags = 0
	// KAudioTimeStampSampleTimeValid: A flag that indicates that the sample frame time is valid.
	KAudioTimeStampSampleTimeValid AudioTimeStampFlags = 1
	// KAudioTimeStampWordClockTimeValid: A flag that indicates that the word clock time is valid.
	KAudioTimeStampWordClockTimeValid AudioTimeStampFlags = 8
)

func (e AudioTimeStampFlags) String() string {
	switch e {
	case KAudioTimeStampHostTimeValid:
		return "KAudioTimeStampHostTimeValid"
	case KAudioTimeStampNothingValid:
		return "KAudioTimeStampNothingValid"
	case KAudioTimeStampRateScalarValid:
		return "KAudioTimeStampRateScalarValid"
	case KAudioTimeStampSMPTETimeValid:
		return "KAudioTimeStampSMPTETimeValid"
	case KAudioTimeStampSampleTimeValid:
		return "KAudioTimeStampSampleTimeValid"
	case KAudioTimeStampWordClockTimeValid:
		return "KAudioTimeStampWordClockTimeValid"
	default:
		return fmt.Sprintf("AudioTimeStampFlags(%d)", e)
	}
}

type KAudio int

const (
	// KAudioFormat60958AC3: A key that specifies the AC-3 codec, which provides data packaged for transport over an IEC 60958-compliant digital audio interface, and uses standard flags.
	KAudioFormat60958AC3 KAudio = 'c'<<24 | 'a'<<16 | 'c'<<8 | '3' // 'cac3'
	// KAudioFormatAC3: A key that specifies the AC-3 codec, and uses no flags.
	KAudioFormatAC3 KAudio = 'a'<<24 | 'c'<<16 | '-'<<8 | '3' // 'ac-3'
	// KAudioFormatAES3: A key that specifies the codec defined by the AES3-2003 standard, and uses no flags.
	KAudioFormatAES3 KAudio = 'a'<<24 | 'e'<<16 | 's'<<8 | '3' // 'aes3'
	// KAudioFormatALaw: A key that specifies the A-law 2:1 codec, and uses no flags.
	KAudioFormatALaw KAudio = 'a'<<24 | 'l'<<16 | 'a'<<8 | 'w' // 'alaw'
	// KAudioFormatAMR: A key that specifies the Adaptive Multi-Rate (AMR) narrow band speech codec, and uses no flags.
	KAudioFormatAMR KAudio = 's'<<24 | 'a'<<16 | 'm'<<8 | 'r' // 'samr'
	// KAudioFormatAMR_WB: A key that specifies the AMR Wideband speech codec, and uses no flags.
	KAudioFormatAMR_WB KAudio = 's'<<24 | 'a'<<16 | 'w'<<8 | 'b' // 'sawb'
	KAudioFormatAPAC   KAudio = 'a'<<24 | 'p'<<16 | 'a'<<8 | 'c' // 'apac'
	// KAudioFormatAppleIMA4: A key that specifies Apple’s implementation of the IMA 4:1 ADPCM codec, and uses no flags.
	KAudioFormatAppleIMA4 KAudio = 'i'<<24 | 'm'<<16 | 'a'<<8 | '4' // 'ima4'
	// KAudioFormatAppleLossless: A key that specifies the Apple Lossless codec, and uses flags to indicate the bit depth of the source material.
	KAudioFormatAppleLossless KAudio = 'a'<<24 | 'l'<<16 | 'a'<<8 | 'c' // 'alac'
	// KAudioFormatAudible: A key that specifies the codec for Audible audio books, and uses no flags.
	KAudioFormatAudible KAudio = 'A'<<24 | 'U'<<16 | 'D'<<8 | 'B' // 'AUDB'
	// KAudioFormatDVIIntelIMA: A key that specifies the codec defined by DVI/Intel IMA ADPCM - ACM code 17, and uses no flags.
	KAudioFormatDVIIntelIMA KAudio = 0x6d730011
	// KAudioFormatEnhancedAC3: A key that specifies the Enhanced AC-3 codec, and uses no flags.
	KAudioFormatEnhancedAC3 KAudio = 'e'<<24 | 'c'<<16 | '-'<<8 | '3' // 'ec-3'
	// KAudioFormatFLAC: A key that specifies the Free Lossless Audio Codec (FLAC), and uses flags to indicate the bit depth of the source material.
	KAudioFormatFLAC KAudio = 'f'<<24 | 'l'<<16 | 'a'<<8 | 'c' // 'flac'
	// KAudioFormatLinearPCM: A key that specifies the linear PCM codec, and uses the standard flags.
	KAudioFormatLinearPCM KAudio = 'l'<<24 | 'p'<<16 | 'c'<<8 | 'm' // 'lpcm'
	// KAudioFormatMACE3: A key that specifies the MACE 3:1 codec, and uses no flags.
	KAudioFormatMACE3 KAudio = 'M'<<24 | 'A'<<16 | 'C'<<8 | '3' // 'MAC3'
	// KAudioFormatMACE6: A key that specifies the MACE C:1 codec, and uses no flags.
	KAudioFormatMACE6 KAudio = 'M'<<24 | 'A'<<16 | 'C'<<8 | '6' // 'MAC6'
	// KAudioFormatMIDIStream: A key that specifies the MIDI stream codec, and uses no flags.
	KAudioFormatMIDIStream KAudio = 'm'<<24 | 'i'<<16 | 'd'<<8 | 'i' // 'midi'
	// KAudioFormatMPEG4AAC: A key that specifies the MPEG-4 AAC Low Complexity codec, and uses no flags.
	KAudioFormatMPEG4AAC KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | ' ' // 'aac '
	// KAudioFormatMPEG4AAC_ELD: A key that specifies the MPEG-4 Enhanced Low Delay AAC codec, and uses no flags.
	KAudioFormatMPEG4AAC_ELD KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'e' // 'aace'
	// KAudioFormatMPEG4AAC_ELD_SBR: A key that specifies the MPEG-4 Enhanced Low Delay AAC codec with a spectral band replication (SBR) extension layer, and uses no flags.
	KAudioFormatMPEG4AAC_ELD_SBR KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'f' // 'aacf'
	// KAudioFormatMPEG4AAC_ELD_V2: A key that specifies the MPEG-4 Enhanced Low Delay AAC version 2 codec, and uses no flags.
	KAudioFormatMPEG4AAC_ELD_V2 KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'g' // 'aacg'
	// KAudioFormatMPEG4AAC_HE: A key that specifies the MPEG-4 High-Efficiency AAC codec, and uses no flags.
	KAudioFormatMPEG4AAC_HE KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'h' // 'aach'
	// KAudioFormatMPEG4AAC_HE_V2: A key that specifies the MPEG-4 High-Efficiency AAC version 2 codec, and uses no flags.
	KAudioFormatMPEG4AAC_HE_V2 KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'p' // 'aacp'
	// KAudioFormatMPEG4AAC_LD: A key that specifies the MPEG-4 Low Delay AAC codec, and uses no flags.
	KAudioFormatMPEG4AAC_LD KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 'l' // 'aacl'
	// KAudioFormatMPEG4AAC_Spatial: A key that specifies the MPEG-4 Spatial Audio Coding codec, and uses no flags.
	KAudioFormatMPEG4AAC_Spatial KAudio = 'a'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'aacs'
	// KAudioFormatMPEG4CELP: A key that specifies the MPEG-4 CELP codec, and uses flags to indicate the specific kind of data.
	KAudioFormatMPEG4CELP KAudio = 'c'<<24 | 'e'<<16 | 'l'<<8 | 'p' // 'celp'
	// KAudioFormatMPEG4HVXC: A key that specifies the MPEG-4 HVXC codec, and uses no flags.
	KAudioFormatMPEG4HVXC KAudio = 'h'<<24 | 'v'<<16 | 'x'<<8 | 'c' // 'hvxc'
	// KAudioFormatMPEG4TwinVQ: A key that specifies the MPEG-4 TwinVQ codec, and uses no flags.
	KAudioFormatMPEG4TwinVQ KAudio = 't'<<24 | 'w'<<16 | 'v'<<8 | 'q' // 'twvq'
	// KAudioFormatMPEGD_USAC: A key that specifies the MPEG-D Unified Speech and Audio Coding codec, and uses no flags.
	KAudioFormatMPEGD_USAC KAudio = 'u'<<24 | 's'<<16 | 'a'<<8 | 'c' // 'usac'
	// KAudioFormatMPEGLayer1: A key that specifies the MPEG-1/2, Layer I audio codec, and uses no flags.
	KAudioFormatMPEGLayer1 KAudio = '.'<<24 | 'm'<<16 | 'p'<<8 | '1' // '.mp1'
	// KAudioFormatMPEGLayer2: A key that specifies the MPEG-1/2, Layer II audio codec, and uses no flags.
	KAudioFormatMPEGLayer2 KAudio = '.'<<24 | 'm'<<16 | 'p'<<8 | '2' // '.mp2'
	// KAudioFormatMPEGLayer3: A key that specifies the MPEG-1/2, Layer III audio codec, and uses no flags.
	KAudioFormatMPEGLayer3 KAudio = '.'<<24 | 'm'<<16 | 'p'<<8 | '3' // '.mp3'
	// KAudioFormatMicrosoftGSM: A key that specifies the Microsoft GSM 6.10 - ACM code 49 codec, and uses no flags.
	KAudioFormatMicrosoftGSM KAudio = 0x6d730031
	// KAudioFormatOpus: A key that specifies the Opus codec, and uses no flags.
	KAudioFormatOpus KAudio = 'o'<<24 | 'p'<<16 | 'u'<<8 | 's' // 'opus'
	// KAudioFormatParameterValueStream: A key that specifies the A side-chain of float 32 data that an audio unit provides for sending high-density parameter value control information, and uses no flags.
	KAudioFormatParameterValueStream KAudio = 'a'<<24 | 'p'<<16 | 'v'<<8 | 's' // 'apvs'
	// KAudioFormatQDesign: A key that specifies the QDesign music codec, and uses no flags.
	KAudioFormatQDesign KAudio = 'Q'<<24 | 'D'<<16 | 'M'<<8 | 'C' // 'QDMC'
	// KAudioFormatQDesign2: A key that specifies the QDesign 2 music codec, and uses no flags.
	KAudioFormatQDesign2 KAudio = 'Q'<<24 | 'D'<<16 | 'M'<<8 | '2' // 'QDM2'
	// KAudioFormatQUALCOMM: A key that specifies the Qualcomm PureVoice codec, and uses no flags.
	KAudioFormatQUALCOMM KAudio = 'Q'<<24 | 'c'<<16 | 'l'<<8 | 'p' // 'Qclp'
	// KAudioFormatTimeCode: A key that specifies the A stream of audio timestamp structures, and uses audio timestamp flags.
	KAudioFormatTimeCode KAudio = 't'<<24 | 'i'<<16 | 'm'<<8 | 'e' // 'time'
	// KAudioFormatULaw: A key that specifies the μ-Law 2:1 codec, and uses no flags.
	KAudioFormatULaw KAudio = 'u'<<24 | 'l'<<16 | 'a'<<8 | 'w' // 'ulaw'
	// KAudioFormatiLBC: A key that specifies the internet Low Bitrate Codec (iLBC) narrow band speech codec, and uses no flags.
	KAudioFormatiLBC        KAudio = 'i'<<24 | 'l'<<16 | 'b'<<8 | 'c' // 'ilbc'
	KAudio_BadFilePathError KAudio = '!'<<24 | 'p'<<16 | 't'<<8 | 'h' // '!pth'
	// KAudio_FileNotFoundError: An error that indicates the file wasn’t found.
	KAudio_FileNotFoundError   KAudio = -43
	KAudio_FilePermissionError KAudio = -54
	// KAudio_MemFullError: An error that indicates that the heap zone is full.
	KAudio_MemFullError KAudio = -108
	KAudio_NoError      KAudio = 0
	// KAudio_ParamError: An error in the parameter list of the function.
	KAudio_ParamError            KAudio = -50
	KAudio_TooManyFilesOpenError KAudio = -42
	// KAudio_UnimplementedError: An error that indicates the app called an unimplemented system function.
	KAudio_UnimplementedError KAudio = -4
)

func (e KAudio) String() string {
	switch e {
	case KAudioFormat60958AC3:
		return "KAudioFormat60958AC3"
	case KAudioFormatAC3:
		return "KAudioFormatAC3"
	case KAudioFormatAES3:
		return "KAudioFormatAES3"
	case KAudioFormatALaw:
		return "KAudioFormatALaw"
	case KAudioFormatAMR:
		return "KAudioFormatAMR"
	case KAudioFormatAMR_WB:
		return "KAudioFormatAMR_WB"
	case KAudioFormatAPAC:
		return "KAudioFormatAPAC"
	case KAudioFormatAppleIMA4:
		return "KAudioFormatAppleIMA4"
	case KAudioFormatAppleLossless:
		return "KAudioFormatAppleLossless"
	case KAudioFormatAudible:
		return "KAudioFormatAudible"
	case KAudioFormatDVIIntelIMA:
		return "KAudioFormatDVIIntelIMA"
	case KAudioFormatEnhancedAC3:
		return "KAudioFormatEnhancedAC3"
	case KAudioFormatFLAC:
		return "KAudioFormatFLAC"
	case KAudioFormatLinearPCM:
		return "KAudioFormatLinearPCM"
	case KAudioFormatMACE3:
		return "KAudioFormatMACE3"
	case KAudioFormatMACE6:
		return "KAudioFormatMACE6"
	case KAudioFormatMIDIStream:
		return "KAudioFormatMIDIStream"
	case KAudioFormatMPEG4AAC:
		return "KAudioFormatMPEG4AAC"
	case KAudioFormatMPEG4AAC_ELD:
		return "KAudioFormatMPEG4AAC_ELD"
	case KAudioFormatMPEG4AAC_ELD_SBR:
		return "KAudioFormatMPEG4AAC_ELD_SBR"
	case KAudioFormatMPEG4AAC_ELD_V2:
		return "KAudioFormatMPEG4AAC_ELD_V2"
	case KAudioFormatMPEG4AAC_HE:
		return "KAudioFormatMPEG4AAC_HE"
	case KAudioFormatMPEG4AAC_HE_V2:
		return "KAudioFormatMPEG4AAC_HE_V2"
	case KAudioFormatMPEG4AAC_LD:
		return "KAudioFormatMPEG4AAC_LD"
	case KAudioFormatMPEG4AAC_Spatial:
		return "KAudioFormatMPEG4AAC_Spatial"
	case KAudioFormatMPEG4CELP:
		return "KAudioFormatMPEG4CELP"
	case KAudioFormatMPEG4HVXC:
		return "KAudioFormatMPEG4HVXC"
	case KAudioFormatMPEG4TwinVQ:
		return "KAudioFormatMPEG4TwinVQ"
	case KAudioFormatMPEGD_USAC:
		return "KAudioFormatMPEGD_USAC"
	case KAudioFormatMPEGLayer1:
		return "KAudioFormatMPEGLayer1"
	case KAudioFormatMPEGLayer2:
		return "KAudioFormatMPEGLayer2"
	case KAudioFormatMPEGLayer3:
		return "KAudioFormatMPEGLayer3"
	case KAudioFormatMicrosoftGSM:
		return "KAudioFormatMicrosoftGSM"
	case KAudioFormatOpus:
		return "KAudioFormatOpus"
	case KAudioFormatParameterValueStream:
		return "KAudioFormatParameterValueStream"
	case KAudioFormatQDesign:
		return "KAudioFormatQDesign"
	case KAudioFormatQDesign2:
		return "KAudioFormatQDesign2"
	case KAudioFormatQUALCOMM:
		return "KAudioFormatQUALCOMM"
	case KAudioFormatTimeCode:
		return "KAudioFormatTimeCode"
	case KAudioFormatULaw:
		return "KAudioFormatULaw"
	case KAudioFormatiLBC:
		return "KAudioFormatiLBC"
	case KAudio_BadFilePathError:
		return "KAudio_BadFilePathError"
	case KAudio_FileNotFoundError:
		return "KAudio_FileNotFoundError"
	case KAudio_FilePermissionError:
		return "KAudio_FilePermissionError"
	case KAudio_MemFullError:
		return "KAudio_MemFullError"
	case KAudio_NoError:
		return "KAudio_NoError"
	case KAudio_ParamError:
		return "KAudio_ParamError"
	case KAudio_TooManyFilesOpenError:
		return "KAudio_TooManyFilesOpenError"
	case KAudio_UnimplementedError:
		return "KAudio_UnimplementedError"
	default:
		return fmt.Sprintf("KAudio(%d)", e)
	}
}

type KAudioChannelLabel uint

const (
	// KAudioChannelLabel_Ambisonic_W: First order Ambisonic channel W.
	KAudioChannelLabel_Ambisonic_W KAudioChannelLabel = 200
	// KAudioChannelLabel_Ambisonic_X: First order Ambisonic channel X.
	KAudioChannelLabel_Ambisonic_X KAudioChannelLabel = 201
	// KAudioChannelLabel_Ambisonic_Y: First order Ambisonic channel Y.
	KAudioChannelLabel_Ambisonic_Y KAudioChannelLabel = 202
	// KAudioChannelLabel_Ambisonic_Z: First order Ambisonic channel Z.
	KAudioChannelLabel_Ambisonic_Z   KAudioChannelLabel = 203
	KAudioChannelLabel_BeginReserved KAudioChannelLabel = 0xf0000000
	KAudioChannelLabel_BinauralLeft  KAudioChannelLabel = 208
	KAudioChannelLabel_BinauralRight KAudioChannelLabel = 209
	// KAudioChannelLabel_Center: Center channel.
	KAudioChannelLabel_Center       KAudioChannelLabel = 3
	KAudioChannelLabel_CenterBottom KAudioChannelLabel = 59
	// KAudioChannelLabel_CenterSurround: Center surround channel; or for WAVE (.wav) files, back center or rear surround.
	KAudioChannelLabel_CenterSurround KAudioChannelLabel = 9
	// KAudioChannelLabel_CenterSurroundDirect: Back center, non diffuse channel.
	KAudioChannelLabel_CenterSurroundDirect KAudioChannelLabel = 44
	KAudioChannelLabel_CenterTopFront       KAudioChannelLabel = 14
	KAudioChannelLabel_CenterTopMiddle      KAudioChannelLabel = 12
	KAudioChannelLabel_CenterTopRear        KAudioChannelLabel = 53
	// KAudioChannelLabel_ClickTrack: Click track channel.
	KAudioChannelLabel_ClickTrack       KAudioChannelLabel = 304
	KAudioChannelLabel_DialogCentricMix KAudioChannelLabel = 43
	// KAudioChannelLabel_Discrete: Generic discrete channel.
	KAudioChannelLabel_Discrete KAudioChannelLabel = 400
	// KAudioChannelLabel_Discrete_0: Discrete channel 0.
	KAudioChannelLabel_Discrete_0 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_1: Discrete channel 1.
	KAudioChannelLabel_Discrete_1 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_10: Discrete channel 10.
	KAudioChannelLabel_Discrete_10 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_11: Discrete channel 11.
	KAudioChannelLabel_Discrete_11 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_12: Discrete channel 12.
	KAudioChannelLabel_Discrete_12 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_13: Discrete channel 13.
	KAudioChannelLabel_Discrete_13 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_14: Discrete channel 14.
	KAudioChannelLabel_Discrete_14 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_15: Discrete channel 15.
	KAudioChannelLabel_Discrete_15 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_2: Discrete channel 2.
	KAudioChannelLabel_Discrete_2 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_3: Discrete channel 3.
	KAudioChannelLabel_Discrete_3 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_4: Discrete channel 4.
	KAudioChannelLabel_Discrete_4 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_5: Discrete channel 5.
	KAudioChannelLabel_Discrete_5 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_6: Discrete channel 6.
	KAudioChannelLabel_Discrete_6 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_65535: Discrete channel 65536.
	KAudioChannelLabel_Discrete_65535 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_7: Discrete channel 7.
	KAudioChannelLabel_Discrete_7 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_8: Discrete channel 8.
	KAudioChannelLabel_Discrete_8 KAudioChannelLabel = 65536
	// KAudioChannelLabel_Discrete_9: Discrete channel 9.
	KAudioChannelLabel_Discrete_9  KAudioChannelLabel = 65536
	KAudioChannelLabel_EndReserved KAudioChannelLabel = 0xfffffffe
	// KAudioChannelLabel_ForeignLanguage: Foreign language channel.
	KAudioChannelLabel_ForeignLanguage KAudioChannelLabel = 305
	KAudioChannelLabel_HOA_ACN         KAudioChannelLabel = 500
	KAudioChannelLabel_HOA_ACN_0       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_1       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_10      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_11      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_12      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_13      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_14      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_15      KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_2       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_3       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_4       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_5       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_6       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_65024   KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_7       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_8       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_ACN_9       KAudioChannelLabel = 131072
	KAudioChannelLabel_HOA_N3D         KAudioChannelLabel = 196608
	KAudioChannelLabel_HOA_SN3D        KAudioChannelLabel = 131072
	// KAudioChannelLabel_Haptic: A channel for haptic (touch) data.
	KAudioChannelLabel_Haptic KAudioChannelLabel = 45
	// KAudioChannelLabel_HeadphonesLeft: Left channel of stereo headphones.
	KAudioChannelLabel_HeadphonesLeft KAudioChannelLabel = 301
	// KAudioChannelLabel_HeadphonesRight: Right channel of stereo headphones.
	KAudioChannelLabel_HeadphonesRight KAudioChannelLabel = 302
	// KAudioChannelLabel_HearingImpaired: Channel carrying audio for people who are deaf or hard of hearing.
	KAudioChannelLabel_HearingImpaired KAudioChannelLabel = 40
	// KAudioChannelLabel_LFE2: Low Frequency Effects 2.
	KAudioChannelLabel_LFE2 KAudioChannelLabel = 37
	KAudioChannelLabel_LFE3 KAudioChannelLabel = 62
	// KAudioChannelLabel_LFEScreen: Low Frequency Effects Screen; a subwoofer located in front of the theater.
	KAudioChannelLabel_LFEScreen KAudioChannelLabel = 4
	// KAudioChannelLabel_Left: Left channel.
	KAudioChannelLabel_Left             KAudioChannelLabel = 1
	KAudioChannelLabel_LeftBackSurround KAudioChannelLabel = 63
	KAudioChannelLabel_LeftBottom       KAudioChannelLabel = 57
	// KAudioChannelLabel_LeftCenter: Left center channel.
	KAudioChannelLabel_LeftCenter       KAudioChannelLabel = 7
	KAudioChannelLabel_LeftEdgeOfScreen KAudioChannelLabel = 65
	KAudioChannelLabel_LeftSideSurround KAudioChannelLabel = 55
	// KAudioChannelLabel_LeftSurround: Left surround channel; or for WAVE (.wav) files, back left.
	KAudioChannelLabel_LeftSurround KAudioChannelLabel = 5
	// KAudioChannelLabel_LeftSurroundDirect: Left surround direct channel; or for WAVE (.wav) files, side left.
	KAudioChannelLabel_LeftSurroundDirect KAudioChannelLabel = 10
	KAudioChannelLabel_LeftTopFront       KAudioChannelLabel = 13
	KAudioChannelLabel_LeftTopMiddle      KAudioChannelLabel = 49
	KAudioChannelLabel_LeftTopRear        KAudioChannelLabel = 52
	KAudioChannelLabel_LeftTopSurround    KAudioChannelLabel = 60
	// KAudioChannelLabel_LeftTotal: The left channel of matrix encoded 4 channel audio.
	KAudioChannelLabel_LeftTotal KAudioChannelLabel = 38
	// KAudioChannelLabel_LeftWide: Left wide channel.
	KAudioChannelLabel_LeftWide KAudioChannelLabel = 35
	// KAudioChannelLabel_MS_Mid: Mid channel of a Mid/Side recording.
	KAudioChannelLabel_MS_Mid KAudioChannelLabel = 204
	// KAudioChannelLabel_MS_Side: Side channel of a Mid/Side recording.
	KAudioChannelLabel_MS_Side KAudioChannelLabel = 205
	// KAudioChannelLabel_Mono: Monaural channel.
	KAudioChannelLabel_Mono KAudioChannelLabel = 42
	// KAudioChannelLabel_Narration: Narration channel.
	KAudioChannelLabel_Narration KAudioChannelLabel = 41
	KAudioChannelLabel_Object    KAudioChannelLabel = 262144
	// KAudioChannelLabel_RearSurroundLeft: Rear surround left channel.
	KAudioChannelLabel_RearSurroundLeft KAudioChannelLabel = 33
	// KAudioChannelLabel_RearSurroundRight: Rear surround right channel.
	KAudioChannelLabel_RearSurroundRight KAudioChannelLabel = 34
	// KAudioChannelLabel_Right: Right channel.
	KAudioChannelLabel_Right             KAudioChannelLabel = 2
	KAudioChannelLabel_RightBackSurround KAudioChannelLabel = 64
	KAudioChannelLabel_RightBottom       KAudioChannelLabel = 58
	// KAudioChannelLabel_RightCenter: Right center channel.
	KAudioChannelLabel_RightCenter       KAudioChannelLabel = 8
	KAudioChannelLabel_RightEdgeOfScreen KAudioChannelLabel = 66
	KAudioChannelLabel_RightSideSurround KAudioChannelLabel = 56
	// KAudioChannelLabel_RightSurround: Right surround channel; or for WAVE (.wav) files, back right.
	KAudioChannelLabel_RightSurround KAudioChannelLabel = 6
	// KAudioChannelLabel_RightSurroundDirect: Right surround direct channel; or for WAVE (.wav) files, side right.
	KAudioChannelLabel_RightSurroundDirect KAudioChannelLabel = 11
	KAudioChannelLabel_RightTopFront       KAudioChannelLabel = 15
	KAudioChannelLabel_RightTopMiddle      KAudioChannelLabel = 51
	KAudioChannelLabel_RightTopRear        KAudioChannelLabel = 54
	KAudioChannelLabel_RightTopSurround    KAudioChannelLabel = 61
	// KAudioChannelLabel_RightTotal: The right channel of matrix encoded 4 channel audio.
	KAudioChannelLabel_RightTotal KAudioChannelLabel = 39
	// KAudioChannelLabel_RightWide: Right wide channel.
	KAudioChannelLabel_RightWide KAudioChannelLabel = 36
	// KAudioChannelLabel_TopBackCenter: Top back center channel.
	KAudioChannelLabel_TopBackCenter KAudioChannelLabel = 17
	// KAudioChannelLabel_TopBackLeft: Top back left channel.
	KAudioChannelLabel_TopBackLeft KAudioChannelLabel = 16
	// KAudioChannelLabel_TopBackRight: Top back right channel.
	KAudioChannelLabel_TopBackRight KAudioChannelLabel = 18
	// KAudioChannelLabel_TopCenterSurround: Top center surround-sound channel.
	KAudioChannelLabel_TopCenterSurround KAudioChannelLabel = 12
	// KAudioChannelLabel_Unknown: Unknown role or unspecified other use for channel.
	KAudioChannelLabel_Unknown KAudioChannelLabel = 0xffffffff
	// KAudioChannelLabel_Unused: The channel is present, but has no intended role or destination.
	KAudioChannelLabel_Unused KAudioChannelLabel = 0
	// KAudioChannelLabel_UseCoordinates: The channel is described solely by the `mCoordinates` field of the [AudioChannelDescription] structure.
	KAudioChannelLabel_UseCoordinates KAudioChannelLabel = 100
	// KAudioChannelLabel_VerticalHeightCenter: Vertical height center channel; or for WAVE (.wav) files, top front center.
	KAudioChannelLabel_VerticalHeightCenter KAudioChannelLabel = 14
	// KAudioChannelLabel_VerticalHeightLeft: Vertical height left channel; or for WAVE (.wav) files, top front left.
	KAudioChannelLabel_VerticalHeightLeft KAudioChannelLabel = 13
	// KAudioChannelLabel_VerticalHeightRight: Vertical height right channel; or for WAVE (.wav) files, top front right.
	KAudioChannelLabel_VerticalHeightRight KAudioChannelLabel = 15
	// KAudioChannelLabel_XY_X: X channel of an X-Y recording.
	KAudioChannelLabel_XY_X KAudioChannelLabel = 206
	// KAudioChannelLabel_XY_Y: Y channel of an X-Y recording.
	KAudioChannelLabel_XY_Y KAudioChannelLabel = 207
)

func (e KAudioChannelLabel) String() string {
	switch e {
	case KAudioChannelLabel_Ambisonic_W:
		return "KAudioChannelLabel_Ambisonic_W"
	case KAudioChannelLabel_Ambisonic_X:
		return "KAudioChannelLabel_Ambisonic_X"
	case KAudioChannelLabel_Ambisonic_Y:
		return "KAudioChannelLabel_Ambisonic_Y"
	case KAudioChannelLabel_Ambisonic_Z:
		return "KAudioChannelLabel_Ambisonic_Z"
	case KAudioChannelLabel_BeginReserved:
		return "KAudioChannelLabel_BeginReserved"
	case KAudioChannelLabel_BinauralLeft:
		return "KAudioChannelLabel_BinauralLeft"
	case KAudioChannelLabel_BinauralRight:
		return "KAudioChannelLabel_BinauralRight"
	case KAudioChannelLabel_Center:
		return "KAudioChannelLabel_Center"
	case KAudioChannelLabel_CenterBottom:
		return "KAudioChannelLabel_CenterBottom"
	case KAudioChannelLabel_CenterSurround:
		return "KAudioChannelLabel_CenterSurround"
	case KAudioChannelLabel_CenterSurroundDirect:
		return "KAudioChannelLabel_CenterSurroundDirect"
	case KAudioChannelLabel_CenterTopFront:
		return "KAudioChannelLabel_CenterTopFront"
	case KAudioChannelLabel_CenterTopMiddle:
		return "KAudioChannelLabel_CenterTopMiddle"
	case KAudioChannelLabel_CenterTopRear:
		return "KAudioChannelLabel_CenterTopRear"
	case KAudioChannelLabel_ClickTrack:
		return "KAudioChannelLabel_ClickTrack"
	case KAudioChannelLabel_DialogCentricMix:
		return "KAudioChannelLabel_DialogCentricMix"
	case KAudioChannelLabel_Discrete:
		return "KAudioChannelLabel_Discrete"
	case KAudioChannelLabel_Discrete_0:
		return "KAudioChannelLabel_Discrete_0"
	case KAudioChannelLabel_EndReserved:
		return "KAudioChannelLabel_EndReserved"
	case KAudioChannelLabel_ForeignLanguage:
		return "KAudioChannelLabel_ForeignLanguage"
	case KAudioChannelLabel_HOA_ACN:
		return "KAudioChannelLabel_HOA_ACN"
	case KAudioChannelLabel_HOA_ACN_0:
		return "KAudioChannelLabel_HOA_ACN_0"
	case KAudioChannelLabel_HOA_N3D:
		return "KAudioChannelLabel_HOA_N3D"
	case KAudioChannelLabel_Haptic:
		return "KAudioChannelLabel_Haptic"
	case KAudioChannelLabel_HeadphonesLeft:
		return "KAudioChannelLabel_HeadphonesLeft"
	case KAudioChannelLabel_HeadphonesRight:
		return "KAudioChannelLabel_HeadphonesRight"
	case KAudioChannelLabel_HearingImpaired:
		return "KAudioChannelLabel_HearingImpaired"
	case KAudioChannelLabel_LFE2:
		return "KAudioChannelLabel_LFE2"
	case KAudioChannelLabel_LFE3:
		return "KAudioChannelLabel_LFE3"
	case KAudioChannelLabel_LFEScreen:
		return "KAudioChannelLabel_LFEScreen"
	case KAudioChannelLabel_Left:
		return "KAudioChannelLabel_Left"
	case KAudioChannelLabel_LeftBackSurround:
		return "KAudioChannelLabel_LeftBackSurround"
	case KAudioChannelLabel_LeftBottom:
		return "KAudioChannelLabel_LeftBottom"
	case KAudioChannelLabel_LeftCenter:
		return "KAudioChannelLabel_LeftCenter"
	case KAudioChannelLabel_LeftEdgeOfScreen:
		return "KAudioChannelLabel_LeftEdgeOfScreen"
	case KAudioChannelLabel_LeftSideSurround:
		return "KAudioChannelLabel_LeftSideSurround"
	case KAudioChannelLabel_LeftSurround:
		return "KAudioChannelLabel_LeftSurround"
	case KAudioChannelLabel_LeftSurroundDirect:
		return "KAudioChannelLabel_LeftSurroundDirect"
	case KAudioChannelLabel_LeftTopFront:
		return "KAudioChannelLabel_LeftTopFront"
	case KAudioChannelLabel_LeftTopMiddle:
		return "KAudioChannelLabel_LeftTopMiddle"
	case KAudioChannelLabel_LeftTopRear:
		return "KAudioChannelLabel_LeftTopRear"
	case KAudioChannelLabel_LeftTopSurround:
		return "KAudioChannelLabel_LeftTopSurround"
	case KAudioChannelLabel_LeftTotal:
		return "KAudioChannelLabel_LeftTotal"
	case KAudioChannelLabel_LeftWide:
		return "KAudioChannelLabel_LeftWide"
	case KAudioChannelLabel_MS_Mid:
		return "KAudioChannelLabel_MS_Mid"
	case KAudioChannelLabel_MS_Side:
		return "KAudioChannelLabel_MS_Side"
	case KAudioChannelLabel_Mono:
		return "KAudioChannelLabel_Mono"
	case KAudioChannelLabel_Narration:
		return "KAudioChannelLabel_Narration"
	case KAudioChannelLabel_Object:
		return "KAudioChannelLabel_Object"
	case KAudioChannelLabel_RearSurroundLeft:
		return "KAudioChannelLabel_RearSurroundLeft"
	case KAudioChannelLabel_RearSurroundRight:
		return "KAudioChannelLabel_RearSurroundRight"
	case KAudioChannelLabel_Right:
		return "KAudioChannelLabel_Right"
	case KAudioChannelLabel_RightBackSurround:
		return "KAudioChannelLabel_RightBackSurround"
	case KAudioChannelLabel_RightBottom:
		return "KAudioChannelLabel_RightBottom"
	case KAudioChannelLabel_RightCenter:
		return "KAudioChannelLabel_RightCenter"
	case KAudioChannelLabel_RightEdgeOfScreen:
		return "KAudioChannelLabel_RightEdgeOfScreen"
	case KAudioChannelLabel_RightSideSurround:
		return "KAudioChannelLabel_RightSideSurround"
	case KAudioChannelLabel_RightSurround:
		return "KAudioChannelLabel_RightSurround"
	case KAudioChannelLabel_RightSurroundDirect:
		return "KAudioChannelLabel_RightSurroundDirect"
	case KAudioChannelLabel_RightTopFront:
		return "KAudioChannelLabel_RightTopFront"
	case KAudioChannelLabel_RightTopMiddle:
		return "KAudioChannelLabel_RightTopMiddle"
	case KAudioChannelLabel_RightTopRear:
		return "KAudioChannelLabel_RightTopRear"
	case KAudioChannelLabel_RightTopSurround:
		return "KAudioChannelLabel_RightTopSurround"
	case KAudioChannelLabel_RightTotal:
		return "KAudioChannelLabel_RightTotal"
	case KAudioChannelLabel_RightWide:
		return "KAudioChannelLabel_RightWide"
	case KAudioChannelLabel_TopBackCenter:
		return "KAudioChannelLabel_TopBackCenter"
	case KAudioChannelLabel_TopBackLeft:
		return "KAudioChannelLabel_TopBackLeft"
	case KAudioChannelLabel_TopBackRight:
		return "KAudioChannelLabel_TopBackRight"
	case KAudioChannelLabel_Unknown:
		return "KAudioChannelLabel_Unknown"
	case KAudioChannelLabel_Unused:
		return "KAudioChannelLabel_Unused"
	case KAudioChannelLabel_UseCoordinates:
		return "KAudioChannelLabel_UseCoordinates"
	case KAudioChannelLabel_XY_X:
		return "KAudioChannelLabel_XY_X"
	case KAudioChannelLabel_XY_Y:
		return "KAudioChannelLabel_XY_Y"
	default:
		return fmt.Sprintf("KAudioChannelLabel(%d)", e)
	}
}

type KAudioChannelLayoutTag uint

const (
	// KAudioChannelLayoutTag_AAC_3_0: An AAC 3-channel audio layout.
	KAudioChannelLayoutTag_AAC_3_0 KAudioChannelLayoutTag = 7471104
	// KAudioChannelLayoutTag_AAC_4_0: An AAC 4-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_4_0 KAudioChannelLayoutTag = 7602176
	// KAudioChannelLayoutTag_AAC_5_0: An AAC 5-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_5_0 KAudioChannelLayoutTag = 7864320
	// KAudioChannelLayoutTag_AAC_5_1: An AAC 5.1-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_5_1 KAudioChannelLayoutTag = 8126464
	// KAudioChannelLayoutTag_AAC_6_0: An AAC 6-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_6_0 KAudioChannelLayoutTag = 9240576
	// KAudioChannelLayoutTag_AAC_6_1: An AAC 6.1-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_6_1 KAudioChannelLayoutTag = 9306112
	// KAudioChannelLayoutTag_AAC_7_0: An AAC 7-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_7_0 KAudioChannelLayoutTag = 9371648
	// KAudioChannelLayoutTag_AAC_7_1: An AAC 7.1-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_7_1 KAudioChannelLayoutTag = 8323072
	// KAudioChannelLayoutTag_AAC_7_1_B: An AAC 7.1-channel, configuration B, surround-based layout.
	KAudioChannelLayoutTag_AAC_7_1_B KAudioChannelLayoutTag = 11993088
	// KAudioChannelLayoutTag_AAC_7_1_C: An AAC 7.1-channel, configuration C, surround-based layout.
	KAudioChannelLayoutTag_AAC_7_1_C KAudioChannelLayoutTag = 12058624
	// KAudioChannelLayoutTag_AAC_Octagonal: An AAC 8-channel surround-based layout.
	KAudioChannelLayoutTag_AAC_Octagonal KAudioChannelLayoutTag = 9437184
	// KAudioChannelLayoutTag_AAC_Quadraphonic: An AAC quadraphonic surround-based layout.
	KAudioChannelLayoutTag_AAC_Quadraphonic KAudioChannelLayoutTag = 7077888
	// KAudioChannelLayoutTag_AC3_1_0_1: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_1_0_1 KAudioChannelLayoutTag = 9764864
	// KAudioChannelLayoutTag_AC3_2_1_1: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_2_1_1 KAudioChannelLayoutTag = 10027008
	// KAudioChannelLayoutTag_AC3_3_0: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_3_0 KAudioChannelLayoutTag = 9830400
	// KAudioChannelLayoutTag_AC3_3_0_1: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_3_0_1 KAudioChannelLayoutTag = 9961472
	// KAudioChannelLayoutTag_AC3_3_1: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_3_1 KAudioChannelLayoutTag = 9895936
	// KAudioChannelLayoutTag_AC3_3_1_1: An AC-3 layout.
	KAudioChannelLayoutTag_AC3_3_1_1 KAudioChannelLayoutTag = 10092544
	// KAudioChannelLayoutTag_Ambisonic_B_Format: An ambisonic B-format audio layout.
	KAudioChannelLayoutTag_Ambisonic_B_Format KAudioChannelLayoutTag = 7012352
	KAudioChannelLayoutTag_Atmos_5_1_2        KAudioChannelLayoutTag = 12713984
	KAudioChannelLayoutTag_Atmos_5_1_4        KAudioChannelLayoutTag = 12779520
	KAudioChannelLayoutTag_Atmos_7_1_2        KAudioChannelLayoutTag = 12845056
	KAudioChannelLayoutTag_Atmos_7_1_4        KAudioChannelLayoutTag = 12582912
	KAudioChannelLayoutTag_Atmos_9_1_6        KAudioChannelLayoutTag = 12648448
	// KAudioChannelLayoutTag_AudioUnit_4: A quadraphonic symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_4 KAudioChannelLayoutTag = 7077888
	// KAudioChannelLayoutTag_AudioUnit_5: A pentagonal symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_5 KAudioChannelLayoutTag = 7143424
	// KAudioChannelLayoutTag_AudioUnit_5_0: A 5-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_5_0 KAudioChannelLayoutTag = 7733248
	// KAudioChannelLayoutTag_AudioUnit_5_1: A 5.1-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_5_1 KAudioChannelLayoutTag = 7929856
	// KAudioChannelLayoutTag_AudioUnit_6: A hexagonal symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_6 KAudioChannelLayoutTag = 7208960
	// KAudioChannelLayoutTag_AudioUnit_6_0: A 6-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_6_0 KAudioChannelLayoutTag = 9109504
	// KAudioChannelLayoutTag_AudioUnit_6_1: A 6.1-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_6_1 KAudioChannelLayoutTag = 8192000
	// KAudioChannelLayoutTag_AudioUnit_7_0: A 7-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_0 KAudioChannelLayoutTag = 9175040
	// KAudioChannelLayoutTag_AudioUnit_7_0_Front: An alternate 7-channel surround-based layout, for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_0_Front KAudioChannelLayoutTag = 9699328
	// KAudioChannelLayoutTag_AudioUnit_7_1: A 7.1-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_1 KAudioChannelLayoutTag = 8388608
	// KAudioChannelLayoutTag_AudioUnit_7_1_Front: A 7.1-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_1_Front KAudioChannelLayoutTag = 8257536
	// KAudioChannelLayoutTag_AudioUnit_8: An octagonal symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_8 KAudioChannelLayoutTag = 7274496
	// KAudioChannelLayoutTag_BeginReserved: The beginning value for a reserved range of layout tags.
	KAudioChannelLayoutTag_BeginReserved KAudioChannelLayoutTag = 0xf0000000
	// KAudioChannelLayoutTag_Binaural: A binaural stereo audio layout.
	KAudioChannelLayoutTag_Binaural KAudioChannelLayoutTag = 6946816
	KAudioChannelLayoutTag_CICP_1   KAudioChannelLayoutTag = 6553600
	KAudioChannelLayoutTag_CICP_10  KAudioChannelLayoutTag = 8650752
	KAudioChannelLayoutTag_CICP_11  KAudioChannelLayoutTag = 8192000
	KAudioChannelLayoutTag_CICP_12  KAudioChannelLayoutTag = 8388608
	KAudioChannelLayoutTag_CICP_13  KAudioChannelLayoutTag = 13369344
	KAudioChannelLayoutTag_CICP_14  KAudioChannelLayoutTag = 13434880
	KAudioChannelLayoutTag_CICP_15  KAudioChannelLayoutTag = 13500416
	KAudioChannelLayoutTag_CICP_16  KAudioChannelLayoutTag = 13565952
	KAudioChannelLayoutTag_CICP_17  KAudioChannelLayoutTag = 13631488
	KAudioChannelLayoutTag_CICP_18  KAudioChannelLayoutTag = 13697024
	KAudioChannelLayoutTag_CICP_19  KAudioChannelLayoutTag = 13762560
	KAudioChannelLayoutTag_CICP_2   KAudioChannelLayoutTag = 6619136
	KAudioChannelLayoutTag_CICP_20  KAudioChannelLayoutTag = 13828096
	KAudioChannelLayoutTag_CICP_3   KAudioChannelLayoutTag = 7405568
	KAudioChannelLayoutTag_CICP_4   KAudioChannelLayoutTag = 7536640
	KAudioChannelLayoutTag_CICP_5   KAudioChannelLayoutTag = 7667712
	KAudioChannelLayoutTag_CICP_6   KAudioChannelLayoutTag = 7929856
	KAudioChannelLayoutTag_CICP_7   KAudioChannelLayoutTag = 8323072
	KAudioChannelLayoutTag_CICP_9   KAudioChannelLayoutTag = 8585216
	// KAudioChannelLayoutTag_Cube: A cubic audio layout.
	KAudioChannelLayoutTag_Cube KAudioChannelLayoutTag = 7340032
	// KAudioChannelLayoutTag_DTS_3_1: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_3_1 KAudioChannelLayoutTag = 11010048
	// KAudioChannelLayoutTag_DTS_4_1: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_4_1 KAudioChannelLayoutTag = 11075584
	// KAudioChannelLayoutTag_DTS_6_0_A: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_0_A KAudioChannelLayoutTag = 11141120
	// KAudioChannelLayoutTag_DTS_6_0_B: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_0_B KAudioChannelLayoutTag = 11206656
	// KAudioChannelLayoutTag_DTS_6_0_C: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_0_C KAudioChannelLayoutTag = 11272192
	// KAudioChannelLayoutTag_DTS_6_1_A: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_1_A KAudioChannelLayoutTag = 11337728
	// KAudioChannelLayoutTag_DTS_6_1_B: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_1_B KAudioChannelLayoutTag = 11403264
	// KAudioChannelLayoutTag_DTS_6_1_C: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_1_C KAudioChannelLayoutTag = 11468800
	// KAudioChannelLayoutTag_DTS_6_1_D: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_6_1_D KAudioChannelLayoutTag = 11927552
	// KAudioChannelLayoutTag_DTS_7_0: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_7_0 KAudioChannelLayoutTag = 11534336
	// KAudioChannelLayoutTag_DTS_7_1: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_7_1 KAudioChannelLayoutTag = 11599872
	// KAudioChannelLayoutTag_DTS_8_0_A: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_8_0_A KAudioChannelLayoutTag = 11665408
	// KAudioChannelLayoutTag_DTS_8_0_B: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_8_0_B KAudioChannelLayoutTag = 11730944
	// KAudioChannelLayoutTag_DTS_8_1_A: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_8_1_A KAudioChannelLayoutTag = 11796480
	// KAudioChannelLayoutTag_DTS_8_1_B: A Blu-ray Disc audio layout, defined by DTS (Digital Theater Systems Ltd.).
	KAudioChannelLayoutTag_DTS_8_1_B KAudioChannelLayoutTag = 11862016
	// KAudioChannelLayoutTag_DVD_0: A DVD monaural audio layout.
	KAudioChannelLayoutTag_DVD_0 KAudioChannelLayoutTag = 6553600
	// KAudioChannelLayoutTag_DVD_1: A DVD stereo audio layout.
	KAudioChannelLayoutTag_DVD_1 KAudioChannelLayoutTag = 6619136
	// KAudioChannelLayoutTag_DVD_10: A DVD 3.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_10 KAudioChannelLayoutTag = 8912896
	// KAudioChannelLayoutTag_DVD_11: A DVD 4.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_11 KAudioChannelLayoutTag = 8978432
	// KAudioChannelLayoutTag_DVD_12: A DVD 5.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_12 KAudioChannelLayoutTag = 7929856
	// KAudioChannelLayoutTag_DVD_13: A DVD 4-channel audio layout.
	KAudioChannelLayoutTag_DVD_13 KAudioChannelLayoutTag = 7536640
	// KAudioChannelLayoutTag_DVD_14: A DVD 5-channel audio layout.
	KAudioChannelLayoutTag_DVD_14 KAudioChannelLayoutTag = 7667712
	// KAudioChannelLayoutTag_DVD_15: A DVD 3.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_15 KAudioChannelLayoutTag = 8912896
	// KAudioChannelLayoutTag_DVD_16: A DVD 4.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_16 KAudioChannelLayoutTag = 8978432
	// KAudioChannelLayoutTag_DVD_17: A DVD 5.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_17 KAudioChannelLayoutTag = 7929856
	// KAudioChannelLayoutTag_DVD_18: A DVD 4.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_18 KAudioChannelLayoutTag = 9043968
	// KAudioChannelLayoutTag_DVD_19: A DVD 5-channel audio layout.
	KAudioChannelLayoutTag_DVD_19 KAudioChannelLayoutTag = 7733248
	// KAudioChannelLayoutTag_DVD_2: A DVD 3-channel audio layout.
	KAudioChannelLayoutTag_DVD_2 KAudioChannelLayoutTag = 8585216
	// KAudioChannelLayoutTag_DVD_20: A DVD 5.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_20 KAudioChannelLayoutTag = 7995392
	// KAudioChannelLayoutTag_DVD_3: A DVD 4-channel audio layout.
	KAudioChannelLayoutTag_DVD_3 KAudioChannelLayoutTag = 8650752
	// KAudioChannelLayoutTag_DVD_4: A DVD 2.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_4 KAudioChannelLayoutTag = 8716288
	// KAudioChannelLayoutTag_DVD_5: A DVD 3.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_5 KAudioChannelLayoutTag = 8781824
	// KAudioChannelLayoutTag_DVD_6: A DVD 4.1-channel audio layout.
	KAudioChannelLayoutTag_DVD_6 KAudioChannelLayoutTag = 8847360
	// KAudioChannelLayoutTag_DVD_7: A DVD 3-channel audio layout.
	KAudioChannelLayoutTag_DVD_7 KAudioChannelLayoutTag = 7405568
	// KAudioChannelLayoutTag_DVD_8: A DVD 4-channel audio layout.
	KAudioChannelLayoutTag_DVD_8 KAudioChannelLayoutTag = 7536640
	// KAudioChannelLayoutTag_DVD_9: A DVD 5-channel audio layout.
	KAudioChannelLayoutTag_DVD_9 KAudioChannelLayoutTag = 7667712
	// KAudioChannelLayoutTag_DiscreteInOrder: A tag used to map input channels to output channels without changing the channel order.
	KAudioChannelLayoutTag_DiscreteInOrder KAudioChannelLayoutTag = 9633792
	// KAudioChannelLayoutTag_EAC3_6_1_A: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_6_1_A KAudioChannelLayoutTag = 10289152
	// KAudioChannelLayoutTag_EAC3_6_1_B: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_6_1_B KAudioChannelLayoutTag = 10354688
	// KAudioChannelLayoutTag_EAC3_6_1_C: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_6_1_C KAudioChannelLayoutTag = 10420224
	// KAudioChannelLayoutTag_EAC3_7_1_A: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_A KAudioChannelLayoutTag = 10485760
	// KAudioChannelLayoutTag_EAC3_7_1_B: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_B KAudioChannelLayoutTag = 10551296
	// KAudioChannelLayoutTag_EAC3_7_1_C: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_C KAudioChannelLayoutTag = 10616832
	// KAudioChannelLayoutTag_EAC3_7_1_D: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_D KAudioChannelLayoutTag = 10682368
	// KAudioChannelLayoutTag_EAC3_7_1_E: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_E KAudioChannelLayoutTag = 10747904
	// KAudioChannelLayoutTag_EAC3_7_1_F: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_F KAudioChannelLayoutTag = 10813440
	// KAudioChannelLayoutTag_EAC3_7_1_G: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_G KAudioChannelLayoutTag = 10878976
	// KAudioChannelLayoutTag_EAC3_7_1_H: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC3_7_1_H KAudioChannelLayoutTag = 10944512
	// KAudioChannelLayoutTag_EAC_6_0_A: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC_6_0_A KAudioChannelLayoutTag = 10158080
	// KAudioChannelLayoutTag_EAC_7_0_A: A Blu-ray Disc audio layout for Enhanced AC-3, also known as Dolby Digital Plus.
	KAudioChannelLayoutTag_EAC_7_0_A KAudioChannelLayoutTag = 10223616
	// KAudioChannelLayoutTag_Emagic_Default_7_1: An Emagic 7.1-channel default audio layout.
	KAudioChannelLayoutTag_Emagic_Default_7_1 KAudioChannelLayoutTag = 8454144
	// KAudioChannelLayoutTag_EndReserved: The ending value for a reserved range of layout tags.
	KAudioChannelLayoutTag_EndReserved KAudioChannelLayoutTag = 0xfffeffff
	// KAudioChannelLayoutTag_HOA_ACN_N3D: A Higher-order Ambisonics, full 3D normalization audio layout.
	KAudioChannelLayoutTag_HOA_ACN_N3D KAudioChannelLayoutTag = 12517376
	// KAudioChannelLayoutTag_HOA_ACN_SN3D: A Higher-order Ambisonics, Schmidt semi-normalization audio layout.
	KAudioChannelLayoutTag_HOA_ACN_SN3D KAudioChannelLayoutTag = 12451840
	// KAudioChannelLayoutTag_Hexagonal: A hexagonal audio layout.
	KAudioChannelLayoutTag_Hexagonal KAudioChannelLayoutTag = 7208960
	// KAudioChannelLayoutTag_ITU_1_0: An ITU 1-channel audio layout.
	KAudioChannelLayoutTag_ITU_1_0 KAudioChannelLayoutTag = 6553600
	// KAudioChannelLayoutTag_ITU_2_0: An ITU 2-channel audio layout.
	KAudioChannelLayoutTag_ITU_2_0 KAudioChannelLayoutTag = 6619136
	// KAudioChannelLayoutTag_ITU_2_1: An ITU 2.1-channel audio layout.
	KAudioChannelLayoutTag_ITU_2_1 KAudioChannelLayoutTag = 8585216
	// KAudioChannelLayoutTag_ITU_2_2: An ITU 2.2-channel audio layout.
	KAudioChannelLayoutTag_ITU_2_2 KAudioChannelLayoutTag = 8650752
	// KAudioChannelLayoutTag_ITU_3_0: An ITU 3-channel audio layout.
	KAudioChannelLayoutTag_ITU_3_0 KAudioChannelLayoutTag = 7405568
	// KAudioChannelLayoutTag_ITU_3_1: An ITU 3.1-channel audio layout.
	KAudioChannelLayoutTag_ITU_3_1 KAudioChannelLayoutTag = 7536640
	// KAudioChannelLayoutTag_ITU_3_2: An ITU 3.2-channel audio layout.
	KAudioChannelLayoutTag_ITU_3_2 KAudioChannelLayoutTag = 7667712
	// KAudioChannelLayoutTag_ITU_3_2_1: An ITU 3.2.1-channel audio layout.
	KAudioChannelLayoutTag_ITU_3_2_1 KAudioChannelLayoutTag = 7929856
	// KAudioChannelLayoutTag_ITU_3_4_1: An ITU 3.4.1-channel audio layout.
	KAudioChannelLayoutTag_ITU_3_4_1           KAudioChannelLayoutTag = 8388608
	KAudioChannelLayoutTag_Logic_4_0_A         KAudioChannelLayoutTag = 7536640
	KAudioChannelLayoutTag_Logic_4_0_B         KAudioChannelLayoutTag = 7602176
	KAudioChannelLayoutTag_Logic_4_0_C         KAudioChannelLayoutTag = 12910592
	KAudioChannelLayoutTag_Logic_5_0_A         KAudioChannelLayoutTag = 7667712
	KAudioChannelLayoutTag_Logic_5_0_B         KAudioChannelLayoutTag = 7733248
	KAudioChannelLayoutTag_Logic_5_0_C         KAudioChannelLayoutTag = 7798784
	KAudioChannelLayoutTag_Logic_5_0_D         KAudioChannelLayoutTag = 7864320
	KAudioChannelLayoutTag_Logic_5_1_A         KAudioChannelLayoutTag = 7929856
	KAudioChannelLayoutTag_Logic_5_1_B         KAudioChannelLayoutTag = 7995392
	KAudioChannelLayoutTag_Logic_5_1_C         KAudioChannelLayoutTag = 8060928
	KAudioChannelLayoutTag_Logic_5_1_D         KAudioChannelLayoutTag = 8126464
	KAudioChannelLayoutTag_Logic_6_0_A         KAudioChannelLayoutTag = 9240576
	KAudioChannelLayoutTag_Logic_6_0_B         KAudioChannelLayoutTag = 12976128
	KAudioChannelLayoutTag_Logic_6_0_C         KAudioChannelLayoutTag = 9109504
	KAudioChannelLayoutTag_Logic_6_1_A         KAudioChannelLayoutTag = 9306112
	KAudioChannelLayoutTag_Logic_6_1_B         KAudioChannelLayoutTag = 13041664
	KAudioChannelLayoutTag_Logic_6_1_C         KAudioChannelLayoutTag = 8192000
	KAudioChannelLayoutTag_Logic_6_1_D         KAudioChannelLayoutTag = 13107200
	KAudioChannelLayoutTag_Logic_7_1_A         KAudioChannelLayoutTag = 8388608
	KAudioChannelLayoutTag_Logic_7_1_B         KAudioChannelLayoutTag = 13172736
	KAudioChannelLayoutTag_Logic_7_1_C         KAudioChannelLayoutTag = 8388608
	KAudioChannelLayoutTag_Logic_7_1_SDDS_A    KAudioChannelLayoutTag = 8257536
	KAudioChannelLayoutTag_Logic_7_1_SDDS_B    KAudioChannelLayoutTag = 8323072
	KAudioChannelLayoutTag_Logic_7_1_SDDS_C    KAudioChannelLayoutTag = 8454144
	KAudioChannelLayoutTag_Logic_Atmos_5_1_2   KAudioChannelLayoutTag = 12713984
	KAudioChannelLayoutTag_Logic_Atmos_5_1_4   KAudioChannelLayoutTag = 12779520
	KAudioChannelLayoutTag_Logic_Atmos_7_1_2   KAudioChannelLayoutTag = 12845056
	KAudioChannelLayoutTag_Logic_Atmos_7_1_4_A KAudioChannelLayoutTag = 12582912
	KAudioChannelLayoutTag_Logic_Atmos_7_1_4_B KAudioChannelLayoutTag = 13238272
	KAudioChannelLayoutTag_Logic_Atmos_7_1_6   KAudioChannelLayoutTag = 13303808
	KAudioChannelLayoutTag_Logic_Mono          KAudioChannelLayoutTag = 6553600
	KAudioChannelLayoutTag_Logic_Quadraphonic  KAudioChannelLayoutTag = 7077888
	KAudioChannelLayoutTag_Logic_Stereo        KAudioChannelLayoutTag = 6619136
	// KAudioChannelLayoutTag_MPEG_1_0: An MPEG 1-channel audio layout.
	KAudioChannelLayoutTag_MPEG_1_0 KAudioChannelLayoutTag = 6553600
	// KAudioChannelLayoutTag_MPEG_2_0: An MPEG 2-channel audio layout.
	KAudioChannelLayoutTag_MPEG_2_0 KAudioChannelLayoutTag = 6619136
	// KAudioChannelLayoutTag_MPEG_3_0_A: An MPEG 3-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_3_0_A KAudioChannelLayoutTag = 7405568
	// KAudioChannelLayoutTag_MPEG_3_0_B: An MPEG 3-channel, configuration B, audio layout.
	KAudioChannelLayoutTag_MPEG_3_0_B KAudioChannelLayoutTag = 7471104
	// KAudioChannelLayoutTag_MPEG_4_0_A: An MPEG 4-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_4_0_A KAudioChannelLayoutTag = 7536640
	// KAudioChannelLayoutTag_MPEG_4_0_B: An MPEG 4-channel, configuration B, audio layout
	KAudioChannelLayoutTag_MPEG_4_0_B KAudioChannelLayoutTag = 7602176
	// KAudioChannelLayoutTag_MPEG_5_0_A: An MPEG 5-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_5_0_A KAudioChannelLayoutTag = 7667712
	// KAudioChannelLayoutTag_MPEG_5_0_B: An MPEG 5-channel, configuration B, audio layout.
	KAudioChannelLayoutTag_MPEG_5_0_B KAudioChannelLayoutTag = 7733248
	// KAudioChannelLayoutTag_MPEG_5_0_C: An MPEG 5-channel, configuration C, audio layout.
	KAudioChannelLayoutTag_MPEG_5_0_C KAudioChannelLayoutTag = 7798784
	// KAudioChannelLayoutTag_MPEG_5_0_D: An MPEG 5-channel, configuration D, audio layout.
	KAudioChannelLayoutTag_MPEG_5_0_D KAudioChannelLayoutTag = 7864320
	// KAudioChannelLayoutTag_MPEG_5_0_E: 5 channels, L R Rls Rrs C
	KAudioChannelLayoutTag_MPEG_5_0_E KAudioChannelLayoutTag = 14155776
	// KAudioChannelLayoutTag_MPEG_5_1_A: An MPEG 5.1-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_5_1_A KAudioChannelLayoutTag = 7929856
	// KAudioChannelLayoutTag_MPEG_5_1_B: An MPEG 5.1-channel, configuration B, audio layout.
	KAudioChannelLayoutTag_MPEG_5_1_B KAudioChannelLayoutTag = 7995392
	// KAudioChannelLayoutTag_MPEG_5_1_C: An MPEG 5.1-channel, configuration C, audio layout.
	KAudioChannelLayoutTag_MPEG_5_1_C KAudioChannelLayoutTag = 8060928
	// KAudioChannelLayoutTag_MPEG_5_1_D: An MPEG 5.1-channel, configuration D, audio layout.
	KAudioChannelLayoutTag_MPEG_5_1_D KAudioChannelLayoutTag = 8126464
	// KAudioChannelLayoutTag_MPEG_5_1_E: 6 channels, L R Rls Rrs C LFE
	KAudioChannelLayoutTag_MPEG_5_1_E KAudioChannelLayoutTag = 14221312
	// KAudioChannelLayoutTag_MPEG_6_1_A: An MPEG 6.1-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_6_1_A KAudioChannelLayoutTag = 8192000
	// KAudioChannelLayoutTag_MPEG_6_1_B: 7 channels, L R Ls Rs C Cs LFE
	KAudioChannelLayoutTag_MPEG_6_1_B KAudioChannelLayoutTag = 14286848
	// KAudioChannelLayoutTag_MPEG_7_1_A: An MPEG 7.1-channel, configuration A, audio layout.
	KAudioChannelLayoutTag_MPEG_7_1_A KAudioChannelLayoutTag = 8257536
	// KAudioChannelLayoutTag_MPEG_7_1_B: An MPEG 7.1-channel, configuration B, audio layout.
	KAudioChannelLayoutTag_MPEG_7_1_B KAudioChannelLayoutTag = 8323072
	// KAudioChannelLayoutTag_MPEG_7_1_C: An MPEG 7.1-channel, configuration C, audio layout.
	KAudioChannelLayoutTag_MPEG_7_1_C KAudioChannelLayoutTag = 8388608
	// KAudioChannelLayoutTag_MPEG_7_1_D: 8 channels, L R Rls Rrs Ls Rs C LFE
	KAudioChannelLayoutTag_MPEG_7_1_D KAudioChannelLayoutTag = 14352384
	// KAudioChannelLayoutTag_MatrixStereo: A matrix-encoded stereo stream.
	KAudioChannelLayoutTag_MatrixStereo KAudioChannelLayoutTag = 6750208
	// KAudioChannelLayoutTag_MidSide: A middle and side channel audio layout.
	KAudioChannelLayoutTag_MidSide KAudioChannelLayoutTag = 6815744
	// KAudioChannelLayoutTag_Mono: A standard monophonic stream.
	KAudioChannelLayoutTag_Mono KAudioChannelLayoutTag = 6553600
	// KAudioChannelLayoutTag_Octagonal: An octagonal audio layout.
	KAudioChannelLayoutTag_Octagonal KAudioChannelLayoutTag = 7274496
	KAudioChannelLayoutTag_Ogg_3_0   KAudioChannelLayoutTag = 9830400
	KAudioChannelLayoutTag_Ogg_4_0   KAudioChannelLayoutTag = 12124160
	KAudioChannelLayoutTag_Ogg_5_0   KAudioChannelLayoutTag = 13893632
	KAudioChannelLayoutTag_Ogg_5_1   KAudioChannelLayoutTag = 13959168
	KAudioChannelLayoutTag_Ogg_6_1   KAudioChannelLayoutTag = 14024704
	KAudioChannelLayoutTag_Ogg_7_1   KAudioChannelLayoutTag = 14090240
	// KAudioChannelLayoutTag_Pentagonal: A pentalgonal audio layout.
	KAudioChannelLayoutTag_Pentagonal KAudioChannelLayoutTag = 7143424
	// KAudioChannelLayoutTag_Quadraphonic: A quadraphonic audio layout.
	KAudioChannelLayoutTag_Quadraphonic KAudioChannelLayoutTag = 7077888
	// KAudioChannelLayoutTag_SMPTE_DTV: An SMPTE DTV audio layout.
	KAudioChannelLayoutTag_SMPTE_DTV KAudioChannelLayoutTag = 8519680
	// KAudioChannelLayoutTag_Stereo: A standard stereophonic stream.
	KAudioChannelLayoutTag_Stereo KAudioChannelLayoutTag = 6619136
	// KAudioChannelLayoutTag_StereoHeadphones: A standard stereo stream; headphone playback implied.
	KAudioChannelLayoutTag_StereoHeadphones KAudioChannelLayoutTag = 6684672
	// KAudioChannelLayoutTag_TMH_10_2_full: An extended TMH 10.2 multiple-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_TMH_10_2_full KAudioChannelLayoutTag = 9568256
	// KAudioChannelLayoutTag_TMH_10_2_std: A TMH 10.2 multiple-channel surround-based layout .
	KAudioChannelLayoutTag_TMH_10_2_std KAudioChannelLayoutTag = 9502720
	// KAudioChannelLayoutTag_Unknown: The channel layout is unknown.
	KAudioChannelLayoutTag_Unknown KAudioChannelLayoutTag = 0xffff0000
	// KAudioChannelLayoutTag_UseChannelBitmap: A bitmap that defines the layout mapping.
	KAudioChannelLayoutTag_UseChannelBitmap KAudioChannelLayoutTag = 65536
	// KAudioChannelLayoutTag_UseChannelDescriptions: An array of audio channel description structures that defines the layout mapping.
	KAudioChannelLayoutTag_UseChannelDescriptions KAudioChannelLayoutTag = 0
	KAudioChannelLayoutTag_WAVE_2_1               KAudioChannelLayoutTag = 8716288
	KAudioChannelLayoutTag_WAVE_3_0               KAudioChannelLayoutTag = 7405568
	KAudioChannelLayoutTag_WAVE_4_0_A             KAudioChannelLayoutTag = 8650752
	KAudioChannelLayoutTag_WAVE_4_0_B             KAudioChannelLayoutTag = 12124160
	KAudioChannelLayoutTag_WAVE_5_0_A             KAudioChannelLayoutTag = 7667712
	KAudioChannelLayoutTag_WAVE_5_0_B             KAudioChannelLayoutTag = 12189696
	KAudioChannelLayoutTag_WAVE_5_1_A             KAudioChannelLayoutTag = 7929856
	KAudioChannelLayoutTag_WAVE_5_1_B             KAudioChannelLayoutTag = 12255232
	KAudioChannelLayoutTag_WAVE_6_1               KAudioChannelLayoutTag = 12320768
	KAudioChannelLayoutTag_WAVE_7_1               KAudioChannelLayoutTag = 12386304
	// KAudioChannelLayoutTag_XY: A coincident, angled microphone pair.
	KAudioChannelLayoutTag_XY KAudioChannelLayoutTag = 6881280
)

func (e KAudioChannelLayoutTag) String() string {
	switch e {
	case KAudioChannelLayoutTag_AAC_3_0:
		return "KAudioChannelLayoutTag_AAC_3_0"
	case KAudioChannelLayoutTag_AAC_4_0:
		return "KAudioChannelLayoutTag_AAC_4_0"
	case KAudioChannelLayoutTag_AAC_5_0:
		return "KAudioChannelLayoutTag_AAC_5_0"
	case KAudioChannelLayoutTag_AAC_5_1:
		return "KAudioChannelLayoutTag_AAC_5_1"
	case KAudioChannelLayoutTag_AAC_6_0:
		return "KAudioChannelLayoutTag_AAC_6_0"
	case KAudioChannelLayoutTag_AAC_6_1:
		return "KAudioChannelLayoutTag_AAC_6_1"
	case KAudioChannelLayoutTag_AAC_7_0:
		return "KAudioChannelLayoutTag_AAC_7_0"
	case KAudioChannelLayoutTag_AAC_7_1:
		return "KAudioChannelLayoutTag_AAC_7_1"
	case KAudioChannelLayoutTag_AAC_7_1_B:
		return "KAudioChannelLayoutTag_AAC_7_1_B"
	case KAudioChannelLayoutTag_AAC_7_1_C:
		return "KAudioChannelLayoutTag_AAC_7_1_C"
	case KAudioChannelLayoutTag_AAC_Octagonal:
		return "KAudioChannelLayoutTag_AAC_Octagonal"
	case KAudioChannelLayoutTag_AAC_Quadraphonic:
		return "KAudioChannelLayoutTag_AAC_Quadraphonic"
	case KAudioChannelLayoutTag_AC3_1_0_1:
		return "KAudioChannelLayoutTag_AC3_1_0_1"
	case KAudioChannelLayoutTag_AC3_2_1_1:
		return "KAudioChannelLayoutTag_AC3_2_1_1"
	case KAudioChannelLayoutTag_AC3_3_0:
		return "KAudioChannelLayoutTag_AC3_3_0"
	case KAudioChannelLayoutTag_AC3_3_0_1:
		return "KAudioChannelLayoutTag_AC3_3_0_1"
	case KAudioChannelLayoutTag_AC3_3_1:
		return "KAudioChannelLayoutTag_AC3_3_1"
	case KAudioChannelLayoutTag_AC3_3_1_1:
		return "KAudioChannelLayoutTag_AC3_3_1_1"
	case KAudioChannelLayoutTag_Ambisonic_B_Format:
		return "KAudioChannelLayoutTag_Ambisonic_B_Format"
	case KAudioChannelLayoutTag_Atmos_5_1_2:
		return "KAudioChannelLayoutTag_Atmos_5_1_2"
	case KAudioChannelLayoutTag_Atmos_5_1_4:
		return "KAudioChannelLayoutTag_Atmos_5_1_4"
	case KAudioChannelLayoutTag_Atmos_7_1_2:
		return "KAudioChannelLayoutTag_Atmos_7_1_2"
	case KAudioChannelLayoutTag_Atmos_7_1_4:
		return "KAudioChannelLayoutTag_Atmos_7_1_4"
	case KAudioChannelLayoutTag_Atmos_9_1_6:
		return "KAudioChannelLayoutTag_Atmos_9_1_6"
	case KAudioChannelLayoutTag_AudioUnit_5:
		return "KAudioChannelLayoutTag_AudioUnit_5"
	case KAudioChannelLayoutTag_AudioUnit_5_0:
		return "KAudioChannelLayoutTag_AudioUnit_5_0"
	case KAudioChannelLayoutTag_AudioUnit_5_1:
		return "KAudioChannelLayoutTag_AudioUnit_5_1"
	case KAudioChannelLayoutTag_AudioUnit_6:
		return "KAudioChannelLayoutTag_AudioUnit_6"
	case KAudioChannelLayoutTag_AudioUnit_6_0:
		return "KAudioChannelLayoutTag_AudioUnit_6_0"
	case KAudioChannelLayoutTag_AudioUnit_6_1:
		return "KAudioChannelLayoutTag_AudioUnit_6_1"
	case KAudioChannelLayoutTag_AudioUnit_7_0:
		return "KAudioChannelLayoutTag_AudioUnit_7_0"
	case KAudioChannelLayoutTag_AudioUnit_7_0_Front:
		return "KAudioChannelLayoutTag_AudioUnit_7_0_Front"
	case KAudioChannelLayoutTag_AudioUnit_7_1:
		return "KAudioChannelLayoutTag_AudioUnit_7_1"
	case KAudioChannelLayoutTag_AudioUnit_7_1_Front:
		return "KAudioChannelLayoutTag_AudioUnit_7_1_Front"
	case KAudioChannelLayoutTag_AudioUnit_8:
		return "KAudioChannelLayoutTag_AudioUnit_8"
	case KAudioChannelLayoutTag_BeginReserved:
		return "KAudioChannelLayoutTag_BeginReserved"
	case KAudioChannelLayoutTag_Binaural:
		return "KAudioChannelLayoutTag_Binaural"
	case KAudioChannelLayoutTag_CICP_1:
		return "KAudioChannelLayoutTag_CICP_1"
	case KAudioChannelLayoutTag_CICP_10:
		return "KAudioChannelLayoutTag_CICP_10"
	case KAudioChannelLayoutTag_CICP_13:
		return "KAudioChannelLayoutTag_CICP_13"
	case KAudioChannelLayoutTag_CICP_14:
		return "KAudioChannelLayoutTag_CICP_14"
	case KAudioChannelLayoutTag_CICP_15:
		return "KAudioChannelLayoutTag_CICP_15"
	case KAudioChannelLayoutTag_CICP_16:
		return "KAudioChannelLayoutTag_CICP_16"
	case KAudioChannelLayoutTag_CICP_17:
		return "KAudioChannelLayoutTag_CICP_17"
	case KAudioChannelLayoutTag_CICP_18:
		return "KAudioChannelLayoutTag_CICP_18"
	case KAudioChannelLayoutTag_CICP_19:
		return "KAudioChannelLayoutTag_CICP_19"
	case KAudioChannelLayoutTag_CICP_2:
		return "KAudioChannelLayoutTag_CICP_2"
	case KAudioChannelLayoutTag_CICP_20:
		return "KAudioChannelLayoutTag_CICP_20"
	case KAudioChannelLayoutTag_CICP_3:
		return "KAudioChannelLayoutTag_CICP_3"
	case KAudioChannelLayoutTag_CICP_4:
		return "KAudioChannelLayoutTag_CICP_4"
	case KAudioChannelLayoutTag_CICP_5:
		return "KAudioChannelLayoutTag_CICP_5"
	case KAudioChannelLayoutTag_CICP_9:
		return "KAudioChannelLayoutTag_CICP_9"
	case KAudioChannelLayoutTag_Cube:
		return "KAudioChannelLayoutTag_Cube"
	case KAudioChannelLayoutTag_DTS_3_1:
		return "KAudioChannelLayoutTag_DTS_3_1"
	case KAudioChannelLayoutTag_DTS_4_1:
		return "KAudioChannelLayoutTag_DTS_4_1"
	case KAudioChannelLayoutTag_DTS_6_0_A:
		return "KAudioChannelLayoutTag_DTS_6_0_A"
	case KAudioChannelLayoutTag_DTS_6_0_B:
		return "KAudioChannelLayoutTag_DTS_6_0_B"
	case KAudioChannelLayoutTag_DTS_6_0_C:
		return "KAudioChannelLayoutTag_DTS_6_0_C"
	case KAudioChannelLayoutTag_DTS_6_1_A:
		return "KAudioChannelLayoutTag_DTS_6_1_A"
	case KAudioChannelLayoutTag_DTS_6_1_B:
		return "KAudioChannelLayoutTag_DTS_6_1_B"
	case KAudioChannelLayoutTag_DTS_6_1_C:
		return "KAudioChannelLayoutTag_DTS_6_1_C"
	case KAudioChannelLayoutTag_DTS_6_1_D:
		return "KAudioChannelLayoutTag_DTS_6_1_D"
	case KAudioChannelLayoutTag_DTS_7_0:
		return "KAudioChannelLayoutTag_DTS_7_0"
	case KAudioChannelLayoutTag_DTS_7_1:
		return "KAudioChannelLayoutTag_DTS_7_1"
	case KAudioChannelLayoutTag_DTS_8_0_A:
		return "KAudioChannelLayoutTag_DTS_8_0_A"
	case KAudioChannelLayoutTag_DTS_8_0_B:
		return "KAudioChannelLayoutTag_DTS_8_0_B"
	case KAudioChannelLayoutTag_DTS_8_1_A:
		return "KAudioChannelLayoutTag_DTS_8_1_A"
	case KAudioChannelLayoutTag_DTS_8_1_B:
		return "KAudioChannelLayoutTag_DTS_8_1_B"
	case KAudioChannelLayoutTag_DVD_10:
		return "KAudioChannelLayoutTag_DVD_10"
	case KAudioChannelLayoutTag_DVD_11:
		return "KAudioChannelLayoutTag_DVD_11"
	case KAudioChannelLayoutTag_DVD_18:
		return "KAudioChannelLayoutTag_DVD_18"
	case KAudioChannelLayoutTag_DVD_20:
		return "KAudioChannelLayoutTag_DVD_20"
	case KAudioChannelLayoutTag_DVD_4:
		return "KAudioChannelLayoutTag_DVD_4"
	case KAudioChannelLayoutTag_DVD_5:
		return "KAudioChannelLayoutTag_DVD_5"
	case KAudioChannelLayoutTag_DVD_6:
		return "KAudioChannelLayoutTag_DVD_6"
	case KAudioChannelLayoutTag_DiscreteInOrder:
		return "KAudioChannelLayoutTag_DiscreteInOrder"
	case KAudioChannelLayoutTag_EAC3_6_1_A:
		return "KAudioChannelLayoutTag_EAC3_6_1_A"
	case KAudioChannelLayoutTag_EAC3_6_1_B:
		return "KAudioChannelLayoutTag_EAC3_6_1_B"
	case KAudioChannelLayoutTag_EAC3_6_1_C:
		return "KAudioChannelLayoutTag_EAC3_6_1_C"
	case KAudioChannelLayoutTag_EAC3_7_1_A:
		return "KAudioChannelLayoutTag_EAC3_7_1_A"
	case KAudioChannelLayoutTag_EAC3_7_1_B:
		return "KAudioChannelLayoutTag_EAC3_7_1_B"
	case KAudioChannelLayoutTag_EAC3_7_1_C:
		return "KAudioChannelLayoutTag_EAC3_7_1_C"
	case KAudioChannelLayoutTag_EAC3_7_1_D:
		return "KAudioChannelLayoutTag_EAC3_7_1_D"
	case KAudioChannelLayoutTag_EAC3_7_1_E:
		return "KAudioChannelLayoutTag_EAC3_7_1_E"
	case KAudioChannelLayoutTag_EAC3_7_1_F:
		return "KAudioChannelLayoutTag_EAC3_7_1_F"
	case KAudioChannelLayoutTag_EAC3_7_1_G:
		return "KAudioChannelLayoutTag_EAC3_7_1_G"
	case KAudioChannelLayoutTag_EAC3_7_1_H:
		return "KAudioChannelLayoutTag_EAC3_7_1_H"
	case KAudioChannelLayoutTag_EAC_6_0_A:
		return "KAudioChannelLayoutTag_EAC_6_0_A"
	case KAudioChannelLayoutTag_EAC_7_0_A:
		return "KAudioChannelLayoutTag_EAC_7_0_A"
	case KAudioChannelLayoutTag_Emagic_Default_7_1:
		return "KAudioChannelLayoutTag_Emagic_Default_7_1"
	case KAudioChannelLayoutTag_EndReserved:
		return "KAudioChannelLayoutTag_EndReserved"
	case KAudioChannelLayoutTag_HOA_ACN_N3D:
		return "KAudioChannelLayoutTag_HOA_ACN_N3D"
	case KAudioChannelLayoutTag_HOA_ACN_SN3D:
		return "KAudioChannelLayoutTag_HOA_ACN_SN3D"
	case KAudioChannelLayoutTag_Logic_4_0_C:
		return "KAudioChannelLayoutTag_Logic_4_0_C"
	case KAudioChannelLayoutTag_Logic_5_0_C:
		return "KAudioChannelLayoutTag_Logic_5_0_C"
	case KAudioChannelLayoutTag_Logic_5_1_C:
		return "KAudioChannelLayoutTag_Logic_5_1_C"
	case KAudioChannelLayoutTag_Logic_6_0_B:
		return "KAudioChannelLayoutTag_Logic_6_0_B"
	case KAudioChannelLayoutTag_Logic_6_1_B:
		return "KAudioChannelLayoutTag_Logic_6_1_B"
	case KAudioChannelLayoutTag_Logic_6_1_D:
		return "KAudioChannelLayoutTag_Logic_6_1_D"
	case KAudioChannelLayoutTag_Logic_7_1_B:
		return "KAudioChannelLayoutTag_Logic_7_1_B"
	case KAudioChannelLayoutTag_Logic_Atmos_7_1_4_B:
		return "KAudioChannelLayoutTag_Logic_Atmos_7_1_4_B"
	case KAudioChannelLayoutTag_Logic_Atmos_7_1_6:
		return "KAudioChannelLayoutTag_Logic_Atmos_7_1_6"
	case KAudioChannelLayoutTag_MPEG_5_0_E:
		return "KAudioChannelLayoutTag_MPEG_5_0_E"
	case KAudioChannelLayoutTag_MPEG_5_1_E:
		return "KAudioChannelLayoutTag_MPEG_5_1_E"
	case KAudioChannelLayoutTag_MPEG_6_1_B:
		return "KAudioChannelLayoutTag_MPEG_6_1_B"
	case KAudioChannelLayoutTag_MPEG_7_1_D:
		return "KAudioChannelLayoutTag_MPEG_7_1_D"
	case KAudioChannelLayoutTag_MatrixStereo:
		return "KAudioChannelLayoutTag_MatrixStereo"
	case KAudioChannelLayoutTag_MidSide:
		return "KAudioChannelLayoutTag_MidSide"
	case KAudioChannelLayoutTag_Ogg_4_0:
		return "KAudioChannelLayoutTag_Ogg_4_0"
	case KAudioChannelLayoutTag_Ogg_5_0:
		return "KAudioChannelLayoutTag_Ogg_5_0"
	case KAudioChannelLayoutTag_Ogg_5_1:
		return "KAudioChannelLayoutTag_Ogg_5_1"
	case KAudioChannelLayoutTag_Ogg_6_1:
		return "KAudioChannelLayoutTag_Ogg_6_1"
	case KAudioChannelLayoutTag_Ogg_7_1:
		return "KAudioChannelLayoutTag_Ogg_7_1"
	case KAudioChannelLayoutTag_SMPTE_DTV:
		return "KAudioChannelLayoutTag_SMPTE_DTV"
	case KAudioChannelLayoutTag_StereoHeadphones:
		return "KAudioChannelLayoutTag_StereoHeadphones"
	case KAudioChannelLayoutTag_TMH_10_2_full:
		return "KAudioChannelLayoutTag_TMH_10_2_full"
	case KAudioChannelLayoutTag_TMH_10_2_std:
		return "KAudioChannelLayoutTag_TMH_10_2_std"
	case KAudioChannelLayoutTag_Unknown:
		return "KAudioChannelLayoutTag_Unknown"
	case KAudioChannelLayoutTag_UseChannelBitmap:
		return "KAudioChannelLayoutTag_UseChannelBitmap"
	case KAudioChannelLayoutTag_UseChannelDescriptions:
		return "KAudioChannelLayoutTag_UseChannelDescriptions"
	case KAudioChannelLayoutTag_WAVE_5_0_B:
		return "KAudioChannelLayoutTag_WAVE_5_0_B"
	case KAudioChannelLayoutTag_WAVE_5_1_B:
		return "KAudioChannelLayoutTag_WAVE_5_1_B"
	case KAudioChannelLayoutTag_WAVE_6_1:
		return "KAudioChannelLayoutTag_WAVE_6_1"
	case KAudioChannelLayoutTag_WAVE_7_1:
		return "KAudioChannelLayoutTag_WAVE_7_1"
	case KAudioChannelLayoutTag_XY:
		return "KAudioChannelLayoutTag_XY"
	default:
		return fmt.Sprintf("KAudioChannelLayoutTag(%d)", e)
	}
}

type KAudioFormatFlagIsFloat uint

const (
	// KAppleLosslessFormatFlag_16BitSourceData: A flag that indicates Apple Lossless data sourced from 16-bit native endian signed integer data.
	KAppleLosslessFormatFlag_16BitSourceData KAudioFormatFlagIsFloat = 1
	// KAppleLosslessFormatFlag_20BitSourceData: A flag that indicates Apple Lossless data sourced from 20-bit native endian signed integer data aligned high in 24 bits.
	KAppleLosslessFormatFlag_20BitSourceData KAudioFormatFlagIsFloat = 2
	// KAppleLosslessFormatFlag_24BitSourceData: A flag that indicates Apple Lossless data sourced from 24-bit native endian signed integer data.
	KAppleLosslessFormatFlag_24BitSourceData KAudioFormatFlagIsFloat = 3
	// KAppleLosslessFormatFlag_32BitSourceData: A flag that indicates Apple Lossless data sourced from 32-bit native endian signed integer data.
	KAppleLosslessFormatFlag_32BitSourceData KAudioFormatFlagIsFloat = 4
	// KAudioFormatFlagIsAlignedHigh: A flag that indicates whether placement of the sample bits is with the high or low bits of the channel.
	KAudioFormatFlagIsAlignedHigh KAudioFormatFlagIsFloat = 16
	// KAudioFormatFlagIsBigEndian: A flag that indicates whether the format is big or little endian.
	KAudioFormatFlagIsBigEndian KAudioFormatFlagIsFloat = 2
	// KAudioFormatFlagIsFloatValue: A flag that indicates whether the format is floating point or integer.
	KAudioFormatFlagIsFloatValue KAudioFormatFlagIsFloat = 1
	// KAudioFormatFlagIsNonInterleaved: A flag that indicates whether the samples for each channel or frame are continguously located, and whether the layout of the channels or frames is end-to-end.
	KAudioFormatFlagIsNonInterleaved KAudioFormatFlagIsFloat = 32
	// KAudioFormatFlagIsNonMixable: A flag that indicates the format is nonmixable.
	KAudioFormatFlagIsNonMixable KAudioFormatFlagIsFloat = 64
	// KAudioFormatFlagIsPacked: A flag that indicates whether placement of the sample bits occupy the entire available bits of the channel.
	KAudioFormatFlagIsPacked KAudioFormatFlagIsFloat = 8
	// KAudioFormatFlagIsSignedInteger: A flag that indicates whether the format is signed or unsigned integer.
	KAudioFormatFlagIsSignedInteger KAudioFormatFlagIsFloat = 4
	// KAudioFormatFlagsAreAllClear: A flag that indicates whether all the flags are clear.
	KAudioFormatFlagsAreAllClear KAudioFormatFlagIsFloat = 0x80000000
	// KLinearPCMFormatFlagIsAlignedHigh: A flag that indicates whether placement of the sample bits is with the high or low bits of the channel.
	KLinearPCMFormatFlagIsAlignedHigh KAudioFormatFlagIsFloat = 16
	// KLinearPCMFormatFlagIsBigEndian: A flag that indicates whether the format is big or little endian.
	KLinearPCMFormatFlagIsBigEndian KAudioFormatFlagIsFloat = 2
	// KLinearPCMFormatFlagIsFloat: A flag that indicates whether the format is floating point or integer.
	KLinearPCMFormatFlagIsFloat KAudioFormatFlagIsFloat = 1
	// KLinearPCMFormatFlagIsNonInterleaved: A flag that indicates whether the samples for each channel or frame are continguously located, and whether the layout of the channels or frames is end-to-end.
	KLinearPCMFormatFlagIsNonInterleaved KAudioFormatFlagIsFloat = 32
	// KLinearPCMFormatFlagIsNonMixable: A flag that indicates the format is nonmixable.
	KLinearPCMFormatFlagIsNonMixable KAudioFormatFlagIsFloat = 64
	// KLinearPCMFormatFlagIsPacked: A flag that indicates whether placement of the sample bits occupy the entire available bits of the channel.
	KLinearPCMFormatFlagIsPacked KAudioFormatFlagIsFloat = 8
	// KLinearPCMFormatFlagIsSignedInteger: A flag that indicates whether the format is signed or unsigned integer.
	KLinearPCMFormatFlagIsSignedInteger KAudioFormatFlagIsFloat = 4
	// KLinearPCMFormatFlagsAreAllClear: A flag that indicates whether all the flags are clear.
	KLinearPCMFormatFlagsAreAllClear KAudioFormatFlagIsFloat = 2147483648
	// KLinearPCMFormatFlagsSampleFractionMask: A flag that indicates the sample fraction mask.
	KLinearPCMFormatFlagsSampleFractionMask KAudioFormatFlagIsFloat = 0
	// KLinearPCMFormatFlagsSampleFractionShift: A flag that indicates the bit position of the PCM flag’s 6-bit bitfield.
	KLinearPCMFormatFlagsSampleFractionShift KAudioFormatFlagIsFloat = 7
)

func (e KAudioFormatFlagIsFloat) String() string {
	switch e {
	case KAppleLosslessFormatFlag_16BitSourceData:
		return "KAppleLosslessFormatFlag_16BitSourceData"
	case KAppleLosslessFormatFlag_20BitSourceData:
		return "KAppleLosslessFormatFlag_20BitSourceData"
	case KAppleLosslessFormatFlag_24BitSourceData:
		return "KAppleLosslessFormatFlag_24BitSourceData"
	case KAppleLosslessFormatFlag_32BitSourceData:
		return "KAppleLosslessFormatFlag_32BitSourceData"
	case KAudioFormatFlagIsAlignedHigh:
		return "KAudioFormatFlagIsAlignedHigh"
	case KAudioFormatFlagIsNonInterleaved:
		return "KAudioFormatFlagIsNonInterleaved"
	case KAudioFormatFlagIsNonMixable:
		return "KAudioFormatFlagIsNonMixable"
	case KAudioFormatFlagIsPacked:
		return "KAudioFormatFlagIsPacked"
	case KAudioFormatFlagsAreAllClear:
		return "KAudioFormatFlagsAreAllClear"
	case KLinearPCMFormatFlagsSampleFractionMask:
		return "KLinearPCMFormatFlagsSampleFractionMask"
	case KLinearPCMFormatFlagsSampleFractionShift:
		return "KLinearPCMFormatFlagsSampleFractionShift"
	default:
		return fmt.Sprintf("KAudioFormatFlagIsFloat(%d)", e)
	}
}

type KAudioFormatFlags uint

const (
	// KAudioFormatFlagsAudioUnitCanonical: The flags for the canonical audio unit and processing sample type.
	KAudioFormatFlagsAudioUnitCanonical KAudioFormatFlags = 1
	// KAudioFormatFlagsCanonical: The set of flags for the canonical input-output audio sample type.
	KAudioFormatFlagsCanonical KAudioFormatFlags = 1
	// KAudioFormatFlagsNativeEndian: A flag that specifies whether the format is big endian, depending on the endianness of the processor at build time.
	KAudioFormatFlagsNativeEndian KAudioFormatFlags = 0
	// KAudioFormatFlagsNativeFloatPacked: The flags for the canonical format of fully packed, native endian floating-point data.
	KAudioFormatFlagsNativeFloatPacked KAudioFormatFlags = 1
)

func (e KAudioFormatFlags) String() string {
	switch e {
	case KAudioFormatFlagsAudioUnitCanonical:
		return "KAudioFormatFlagsAudioUnitCanonical"
	case KAudioFormatFlagsNativeEndian:
		return "KAudioFormatFlagsNativeEndian"
	default:
		return fmt.Sprintf("KAudioFormatFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID
type MPEG4ObjectID int

const (
	// KMPEG4Object_AAC_LC: A constant that specifies lossless coding, which provides compression with no loss of quality.
	KMPEG4Object_AAC_LC MPEG4ObjectID = 2
	// KMPEG4Object_AAC_LTP: A constant that specifies long-term prediction, which reduces redundancy in a coded signal.
	KMPEG4Object_AAC_LTP MPEG4ObjectID = 4
	// KMPEG4Object_AAC_Main: A constant that specifies advanced audio coding, which is the basic MPEG-4 technology.
	KMPEG4Object_AAC_Main MPEG4ObjectID = 1
	// KMPEG4Object_AAC_SBR: A constant that specifies spectral band replication, which reconstructs high-frequency content from lower frequencies and side information.
	KMPEG4Object_AAC_SBR MPEG4ObjectID = 5
	// KMPEG4Object_AAC_SSR: A constant that specifies scalable sampling rate, which provides different sampling frequencies for different targets.
	KMPEG4Object_AAC_SSR MPEG4ObjectID = 3
	// KMPEG4Object_AAC_Scalable: A constant that specifies scalable lossless coding.
	KMPEG4Object_AAC_Scalable MPEG4ObjectID = 6
	// KMPEG4Object_CELP: A constant that specifies code-excited linear prediction, which is a narrow-band/wide-band speech codec.
	KMPEG4Object_CELP MPEG4ObjectID = 8
	// KMPEG4Object_HVXC: A constant that specifies harmonic vector excitation coding, which is a very-low bit-rate parametric speech codec.
	KMPEG4Object_HVXC MPEG4ObjectID = 9
	// KMPEG4Object_TwinVQ: A constant that specifies transform-domain weighted interleaved vector quantization.
	KMPEG4Object_TwinVQ MPEG4ObjectID = 7
)

func (e MPEG4ObjectID) String() string {
	switch e {
	case KMPEG4Object_AAC_LC:
		return "KMPEG4Object_AAC_LC"
	case KMPEG4Object_AAC_LTP:
		return "KMPEG4Object_AAC_LTP"
	case KMPEG4Object_AAC_Main:
		return "KMPEG4Object_AAC_Main"
	case KMPEG4Object_AAC_SBR:
		return "KMPEG4Object_AAC_SBR"
	case KMPEG4Object_AAC_SSR:
		return "KMPEG4Object_AAC_SSR"
	case KMPEG4Object_AAC_Scalable:
		return "KMPEG4Object_AAC_Scalable"
	case KMPEG4Object_CELP:
		return "KMPEG4Object_CELP"
	case KMPEG4Object_HVXC:
		return "KMPEG4Object_HVXC"
	case KMPEG4Object_TwinVQ:
		return "KMPEG4Object_TwinVQ"
	default:
		return fmt.Sprintf("MPEG4ObjectID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags
type SMPTETimeFlags uint32

const (
	// KSMPTETimeRunning: # Discussion
	KSMPTETimeRunning SMPTETimeFlags = 2
	KSMPTETimeUnknown SMPTETimeFlags = 0
	// KSMPTETimeValid: # Discussion
	KSMPTETimeValid SMPTETimeFlags = 1
)

func (e SMPTETimeFlags) String() string {
	switch e {
	case KSMPTETimeRunning:
		return "KSMPTETimeRunning"
	case KSMPTETimeUnknown:
		return "KSMPTETimeUnknown"
	case KSMPTETimeValid:
		return "KSMPTETimeValid"
	default:
		return fmt.Sprintf("SMPTETimeFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType
type SMPTETimeType uint32

const (
	// KSMPTETimeType2398: # Discussion
	KSMPTETimeType2398 SMPTETimeType = 11
	// KSMPTETimeType24: # Discussion
	KSMPTETimeType24 SMPTETimeType = 0
	// KSMPTETimeType25: # Discussion
	KSMPTETimeType25 SMPTETimeType = 1
	// KSMPTETimeType2997: # Discussion
	KSMPTETimeType2997 SMPTETimeType = 4
	// KSMPTETimeType2997Drop: # Discussion
	KSMPTETimeType2997Drop SMPTETimeType = 5
	// KSMPTETimeType30: # Discussion
	KSMPTETimeType30 SMPTETimeType = 3
	// KSMPTETimeType30Drop: # Discussion
	KSMPTETimeType30Drop SMPTETimeType = 2
	// KSMPTETimeType50: # Discussion
	KSMPTETimeType50 SMPTETimeType = 10
	// KSMPTETimeType5994: # Discussion
	KSMPTETimeType5994 SMPTETimeType = 7
	// KSMPTETimeType5994Drop: # Discussion
	KSMPTETimeType5994Drop SMPTETimeType = 9
	// KSMPTETimeType60: # Discussion
	KSMPTETimeType60 SMPTETimeType = 6
	// KSMPTETimeType60Drop: # Discussion
	KSMPTETimeType60Drop SMPTETimeType = 8
)

func (e SMPTETimeType) String() string {
	switch e {
	case KSMPTETimeType2398:
		return "KSMPTETimeType2398"
	case KSMPTETimeType24:
		return "KSMPTETimeType24"
	case KSMPTETimeType25:
		return "KSMPTETimeType25"
	case KSMPTETimeType2997:
		return "KSMPTETimeType2997"
	case KSMPTETimeType2997Drop:
		return "KSMPTETimeType2997Drop"
	case KSMPTETimeType30:
		return "KSMPTETimeType30"
	case KSMPTETimeType30Drop:
		return "KSMPTETimeType30Drop"
	case KSMPTETimeType50:
		return "KSMPTETimeType50"
	case KSMPTETimeType5994:
		return "KSMPTETimeType5994"
	case KSMPTETimeType5994Drop:
		return "KSMPTETimeType5994Drop"
	case KSMPTETimeType60:
		return "KSMPTETimeType60"
	case KSMPTETimeType60Drop:
		return "KSMPTETimeType60Drop"
	default:
		return fmt.Sprintf("SMPTETimeType(%d)", e)
	}
}
