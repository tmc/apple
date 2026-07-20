// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreVideo/CVAttachmentMode
type CVAttachmentMode uint32

const (
	// KCVAttachmentMode_ShouldNotPropagate: Indicates to not propagate the attachment.
	KCVAttachmentMode_ShouldNotPropagate CVAttachmentMode = 0
	// KCVAttachmentMode_ShouldPropagate: Indicates to copy the attachment.
	KCVAttachmentMode_ShouldPropagate CVAttachmentMode = 1
)

func (e CVAttachmentMode) String() string {
	switch e {
	case KCVAttachmentMode_ShouldNotPropagate:
		return "KCVAttachmentMode_ShouldNotPropagate"
	case KCVAttachmentMode_ShouldPropagate:
		return "KCVAttachmentMode_ShouldPropagate"
	default:
		return fmt.Sprintf("CVAttachmentMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockFlags
type CVPixelBufferLockFlags int

const (
	// KCVPixelBufferLock_ReadOnly: A read-only buffer.
	KCVPixelBufferLock_ReadOnly CVPixelBufferLockFlags = 0x1
)

func (e CVPixelBufferLockFlags) String() string {
	switch e {
	case KCVPixelBufferLock_ReadOnly:
		return "KCVPixelBufferLock_ReadOnly"
	default:
		return fmt.Sprintf("CVPixelBufferLockFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlushFlags
type CVPixelBufferPoolFlushFlags int

const (
	// KCVPixelBufferPoolFlushExcessBuffers: The value to pass to flush all unused buffers regardless of age.
	KCVPixelBufferPoolFlushExcessBuffers CVPixelBufferPoolFlushFlags = 1
)

func (e CVPixelBufferPoolFlushFlags) String() string {
	switch e {
	case KCVPixelBufferPoolFlushExcessBuffers:
		return "KCVPixelBufferPoolFlushExcessBuffers"
	default:
		return fmt.Sprintf("CVPixelBufferPoolFlushFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeFlags
type CVSMPTETimeFlags uint32

const (
	// KCVSMPTETimeRunning: Time is running.
	KCVSMPTETimeRunning CVSMPTETimeFlags = 2
	// KCVSMPTETimeValid: The full time is valid.
	KCVSMPTETimeValid CVSMPTETimeFlags = 1
)

func (e CVSMPTETimeFlags) String() string {
	switch e {
	case KCVSMPTETimeRunning:
		return "KCVSMPTETimeRunning"
	case KCVSMPTETimeValid:
		return "KCVSMPTETimeValid"
	default:
		return fmt.Sprintf("CVSMPTETimeFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType
type CVSMPTETimeType uint32

const (
	// KCVSMPTETimeType24: 24 frames per second (standard film).
	KCVSMPTETimeType24 CVSMPTETimeType = 0
	// KCVSMPTETimeType25: 25 frames per second (standard PAL).
	KCVSMPTETimeType25 CVSMPTETimeType = 1
	// KCVSMPTETimeType2997: 29.97 frames per second (standard NTSC).
	KCVSMPTETimeType2997 CVSMPTETimeType = 4
	// KCVSMPTETimeType2997Drop: 29.97 drop frame.
	KCVSMPTETimeType2997Drop CVSMPTETimeType = 5
	// KCVSMPTETimeType30: 30 frames per second.
	KCVSMPTETimeType30 CVSMPTETimeType = 3
	// KCVSMPTETimeType30Drop: 30 drop frame.
	KCVSMPTETimeType30Drop CVSMPTETimeType = 2
	// KCVSMPTETimeType5994: 59.94 frames per second.
	KCVSMPTETimeType5994 CVSMPTETimeType = 7
	// KCVSMPTETimeType60: 60 frames per second.
	KCVSMPTETimeType60 CVSMPTETimeType = 6
)

func (e CVSMPTETimeType) String() string {
	switch e {
	case KCVSMPTETimeType24:
		return "KCVSMPTETimeType24"
	case KCVSMPTETimeType25:
		return "KCVSMPTETimeType25"
	case KCVSMPTETimeType2997:
		return "KCVSMPTETimeType2997"
	case KCVSMPTETimeType2997Drop:
		return "KCVSMPTETimeType2997Drop"
	case KCVSMPTETimeType30:
		return "KCVSMPTETimeType30"
	case KCVSMPTETimeType30Drop:
		return "KCVSMPTETimeType30Drop"
	case KCVSMPTETimeType5994:
		return "KCVSMPTETimeType5994"
	case KCVSMPTETimeType60:
		return "KCVSMPTETimeType60"
	default:
		return fmt.Sprintf("CVSMPTETimeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVTimeFlags
type CVTimeFlags int32

const (
	// KCVTimeIsIndefinite: The time value is unknown.
	KCVTimeIsIndefinite CVTimeFlags = 1
)

func (e CVTimeFlags) String() string {
	switch e {
	case KCVTimeIsIndefinite:
		return "KCVTimeIsIndefinite"
	default:
		return fmt.Sprintf("CVTimeFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags
type CVTimeStampFlags uint64

const (
	// KCVTimeStampBottomField: The timestamp represents the bottom lines of an interlaced image.
	KCVTimeStampBottomField CVTimeStampFlags = 131072
	// KCVTimeStampHostTimeValid: The value in the host time field is valid.
	KCVTimeStampHostTimeValid CVTimeStampFlags = 2
	// KCVTimeStampIsInterlaced: A convenience constant indicating that the timestamp is for an interlaced image.
	KCVTimeStampIsInterlaced CVTimeStampFlags = 0
	// KCVTimeStampRateScalarValid: The value in the rate scalar field is valid.
	KCVTimeStampRateScalarValid CVTimeStampFlags = 16
	// KCVTimeStampSMPTETimeValid: The value in the SMPTE time field is valid.
	KCVTimeStampSMPTETimeValid CVTimeStampFlags = 4
	// KCVTimeStampTopField: The timestamp represents the top lines of an interlaced image.
	KCVTimeStampTopField CVTimeStampFlags = 65536
	// KCVTimeStampVideoHostTimeValid: A convenience constant indicating that both the video time and host time fields are valid.
	KCVTimeStampVideoHostTimeValid CVTimeStampFlags = 0
	// KCVTimeStampVideoRefreshPeriodValid: The value in the video refresh period field is valid.
	KCVTimeStampVideoRefreshPeriodValid CVTimeStampFlags = 8
	// KCVTimeStampVideoTimeValid: The value in the video time field is valid.
	KCVTimeStampVideoTimeValid CVTimeStampFlags = 1
)

func (e CVTimeStampFlags) String() string {
	switch e {
	case KCVTimeStampBottomField:
		return "KCVTimeStampBottomField"
	case KCVTimeStampHostTimeValid:
		return "KCVTimeStampHostTimeValid"
	case KCVTimeStampIsInterlaced:
		return "KCVTimeStampIsInterlaced"
	case KCVTimeStampRateScalarValid:
		return "KCVTimeStampRateScalarValid"
	case KCVTimeStampSMPTETimeValid:
		return "KCVTimeStampSMPTETimeValid"
	case KCVTimeStampTopField:
		return "KCVTimeStampTopField"
	case KCVTimeStampVideoRefreshPeriodValid:
		return "KCVTimeStampVideoRefreshPeriodValid"
	case KCVTimeStampVideoTimeValid:
		return "KCVTimeStampVideoTimeValid"
	default:
		return fmt.Sprintf("CVTimeStampFlags(%d)", e)
	}
}

type KCVPixelFormatType uint

const (
	// KCVPixelFormatType_128RGBAFloat: 128-bit RGBA IEEE float, 32-bit little-endian samples.
	KCVPixelFormatType_128RGBAFloat KCVPixelFormatType = 'R'<<24 | 'G'<<16 | 'f'<<8 | 'A' // 'RGfA'
	KCVPixelFormatType_14Bayer_BGGR KCVPixelFormatType = 'b'<<24 | 'g'<<16 | 'g'<<8 | '4' // 'bgg4'
	KCVPixelFormatType_14Bayer_GBRG KCVPixelFormatType = 'g'<<24 | 'b'<<16 | 'r'<<8 | '4' // 'gbr4'
	KCVPixelFormatType_14Bayer_GRBG KCVPixelFormatType = 'g'<<24 | 'r'<<16 | 'b'<<8 | '4' // 'grb4'
	KCVPixelFormatType_14Bayer_RGGB KCVPixelFormatType = 'r'<<24 | 'g'<<16 | 'g'<<8 | '4' // 'rgg4'
	// KCVPixelFormatType_16BE555: 16-bit BE RGB 555.
	KCVPixelFormatType_16BE555 KCVPixelFormatType = 0x10
	// KCVPixelFormatType_16BE565: 16-bit BE RGB 565.
	KCVPixelFormatType_16BE565 KCVPixelFormatType = 'B'<<24 | '5'<<16 | '6'<<8 | '5' // 'B565'
	// KCVPixelFormatType_16Gray: 16-bit Grayscale, 16-bit big-endian samples, black is zero.
	KCVPixelFormatType_16Gray KCVPixelFormatType = 'b'<<24 | '1'<<16 | '6'<<8 | 'g' // 'b16g'
	// KCVPixelFormatType_16LE555: 16-bit LE RGB 555.
	KCVPixelFormatType_16LE555 KCVPixelFormatType = 'L'<<24 | '5'<<16 | '5'<<8 | '5' // 'L555'
	// KCVPixelFormatType_16LE5551: 16-bit LE RGB 5551.
	KCVPixelFormatType_16LE5551 KCVPixelFormatType = '5'<<24 | '5'<<16 | '5'<<8 | '1' // '5551'
	// KCVPixelFormatType_16LE565: 16-bit LE RGB 565.
	KCVPixelFormatType_16LE565          KCVPixelFormatType = 'L'<<24 | '5'<<16 | '6'<<8 | '5' // 'L565'
	KCVPixelFormatType_16VersatileBayer KCVPixelFormatType = 'b'<<24 | 'p'<<16 | '1'<<8 | '6' // 'bp16'
	// KCVPixelFormatType_1IndexedGray_WhiteIsZero: 1 bit indexed gray, white is zero.
	KCVPixelFormatType_1IndexedGray_WhiteIsZero KCVPixelFormatType = 0x21
	// KCVPixelFormatType_1Monochrome: 1 bit indexed.
	KCVPixelFormatType_1Monochrome KCVPixelFormatType = 0x1
	// KCVPixelFormatType_24BGR: 24-bit BGR.
	KCVPixelFormatType_24BGR KCVPixelFormatType = '2'<<24 | '4'<<16 | 'B'<<8 | 'G' // '24BG'
	// KCVPixelFormatType_24RGB: 24-bit RGB.
	KCVPixelFormatType_24RGB KCVPixelFormatType = 0x18
	// KCVPixelFormatType_2Indexed: 2-bit indexed.
	KCVPixelFormatType_2Indexed KCVPixelFormatType = 0x2
	// KCVPixelFormatType_2IndexedGray_WhiteIsZero: 2-bit indexed gray, white is zero.
	KCVPixelFormatType_2IndexedGray_WhiteIsZero KCVPixelFormatType = 0x22
	// KCVPixelFormatType_30RGB: 30-bit RGB, 10-bit big-endian samples, 2 unused padding bits (at least significant end).
	KCVPixelFormatType_30RGB                  KCVPixelFormatType = 'R'<<24 | '1'<<16 | '0'<<8 | 'k' // 'R10k'
	KCVPixelFormatType_30RGBLEPackedWideGamut KCVPixelFormatType = 'w'<<24 | '3'<<16 | '0'<<8 | 'r' // 'w30r'
	KCVPixelFormatType_30RGBLE_8A_BiPlanar    KCVPixelFormatType = 'b'<<24 | '3'<<16 | 'a'<<8 | '8' // 'b3a8'
	KCVPixelFormatType_30RGB_r210             KCVPixelFormatType = 'r'<<24 | '2'<<16 | '1'<<8 | '0' // 'r210'
	// KCVPixelFormatType_32ABGR: 32-bit ABGR.
	KCVPixelFormatType_32ABGR KCVPixelFormatType = 'A'<<24 | 'B'<<16 | 'G'<<8 | 'R' // 'ABGR'
	// KCVPixelFormatType_32ARGB: 32-bit ARGB.
	KCVPixelFormatType_32ARGB KCVPixelFormatType = 0x20
	// KCVPixelFormatType_32AlphaGray: 32-bit AlphaGray, 16-bit big-endian samples, black is zero.
	KCVPixelFormatType_32AlphaGray KCVPixelFormatType = 'b'<<24 | '3'<<16 | '2'<<8 | 'a' // 'b32a'
	// KCVPixelFormatType_32BGRA: 32-bit BGRA.
	KCVPixelFormatType_32BGRA KCVPixelFormatType = 'B'<<24 | 'G'<<16 | 'R'<<8 | 'A' // 'BGRA'
	// KCVPixelFormatType_32RGBA: 32-bit RGBA.
	KCVPixelFormatType_32RGBA                         KCVPixelFormatType = 'R'<<24 | 'G'<<16 | 'B'<<8 | 'A' // 'RGBA'
	KCVPixelFormatType_40ARGBLEWideGamut              KCVPixelFormatType = 'w'<<24 | '4'<<16 | '0'<<8 | 'a' // 'w40a'
	KCVPixelFormatType_40ARGBLEWideGamutPremultiplied KCVPixelFormatType = 'w'<<24 | '4'<<16 | '0'<<8 | 'm' // 'w40m'
	KCVPixelFormatType_420YpCbCr10BiPlanarFullRange   KCVPixelFormatType = 'x'<<24 | 'f'<<16 | '2'<<8 | '0' // 'xf20'
	KCVPixelFormatType_420YpCbCr10BiPlanarVideoRange  KCVPixelFormatType = 'x'<<24 | '4'<<16 | '2'<<8 | '0' // 'x420'
	// KCVPixelFormatType_420YpCbCr8BiPlanarFullRange: Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]).
	KCVPixelFormatType_420YpCbCr8BiPlanarFullRange KCVPixelFormatType = '4'<<24 | '2'<<16 | '0'<<8 | 'f' // '420f'
	// KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange: Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]).
	KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange KCVPixelFormatType = '4'<<24 | '2'<<16 | '0'<<8 | 'v' // '420v'
	// KCVPixelFormatType_420YpCbCr8Planar: Planar Component Y’CbCr 8-bit 4:2:0.
	KCVPixelFormatType_420YpCbCr8Planar KCVPixelFormatType = 'y'<<24 | '4'<<16 | '2'<<8 | '0' // 'y420'
	// KCVPixelFormatType_420YpCbCr8PlanarFullRange: Planar Component Y’CbCr 8-bit 4:2:0, full range.
	KCVPixelFormatType_420YpCbCr8PlanarFullRange         KCVPixelFormatType = 'f'<<24 | '4'<<16 | '2'<<8 | '0' // 'f420'
	KCVPixelFormatType_420YpCbCr8VideoRange_8A_TriPlanar KCVPixelFormatType = 'v'<<24 | '0'<<16 | 'a'<<8 | '8' // 'v0a8'
	// KCVPixelFormatType_422YpCbCr10: Component Y’CbCr 10-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr10                   KCVPixelFormatType = 'v'<<24 | '2'<<16 | '1'<<8 | '0' // 'v210'
	KCVPixelFormatType_422YpCbCr10BiPlanarFullRange  KCVPixelFormatType = 'x'<<24 | 'f'<<16 | '2'<<8 | '2' // 'xf22'
	KCVPixelFormatType_422YpCbCr10BiPlanarVideoRange KCVPixelFormatType = 'x'<<24 | '4'<<16 | '2'<<8 | '2' // 'x422'
	// KCVPixelFormatType_422YpCbCr16: Component Y’CbCr 10,12,14,16-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr16                   KCVPixelFormatType = 'v'<<24 | '2'<<16 | '1'<<8 | '6' // 'v216'
	KCVPixelFormatType_422YpCbCr16BiPlanarVideoRange KCVPixelFormatType = 's'<<24 | 'v'<<16 | '2'<<8 | '2' // 'sv22'
	// KCVPixelFormatType_422YpCbCr8: Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
	KCVPixelFormatType_422YpCbCr8                   KCVPixelFormatType = '2'<<24 | 'v'<<16 | 'u'<<8 | 'y' // '2vuy'
	KCVPixelFormatType_422YpCbCr8BiPlanarFullRange  KCVPixelFormatType = '4'<<24 | '2'<<16 | '2'<<8 | 'f' // '422f'
	KCVPixelFormatType_422YpCbCr8BiPlanarVideoRange KCVPixelFormatType = '4'<<24 | '2'<<16 | '2'<<8 | 'v' // '422v'
	// KCVPixelFormatType_422YpCbCr8FullRange: Component Y’CbCr 8-bit 4:2:2, full range, ordered Y’0 Cb Y’1 Cr.
	KCVPixelFormatType_422YpCbCr8FullRange KCVPixelFormatType = 'y'<<24 | 'u'<<16 | 'v'<<8 | 'f' // 'yuvf'
	// KCVPixelFormatType_422YpCbCr8_yuvs: Component Y’CbCr 8-bit 4:2:2, ordered Y’0 Cb Y’1 Cr.
	KCVPixelFormatType_422YpCbCr8_yuvs KCVPixelFormatType = 'y'<<24 | 'u'<<16 | 'v'<<8 | 's' // 'yuvs'
	// KCVPixelFormatType_422YpCbCr_4A_8BiPlanar: First plane: Video-range Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1; second plane: alpha 8-bit 0-255.
	KCVPixelFormatType_422YpCbCr_4A_8BiPlanar KCVPixelFormatType = 'a'<<24 | '2'<<16 | 'v'<<8 | 'y' // 'a2vy'
	// KCVPixelFormatType_4444AYpCbCr16: Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
	KCVPixelFormatType_4444AYpCbCr16 KCVPixelFormatType = 'y'<<24 | '4'<<16 | '1'<<8 | '6' // 'y416'
	// KCVPixelFormatType_4444AYpCbCr8: Component Y’CbCrA 8-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr.
	KCVPixelFormatType_4444AYpCbCr8     KCVPixelFormatType = 'y'<<24 | '4'<<16 | '0'<<8 | '8' // 'y408'
	KCVPixelFormatType_4444AYpCbCrFloat KCVPixelFormatType = 'r'<<24 | '4'<<16 | 'f'<<8 | 'l' // 'r4fl'
	// KCVPixelFormatType_4444YpCbCrA8: Component Y’CbCrA 8-bit 4:4:4:4, ordered Cb Y’ Cr A.
	KCVPixelFormatType_4444YpCbCrA8 KCVPixelFormatType = 'v'<<24 | '4'<<16 | '0'<<8 | '8' // 'v408'
	// KCVPixelFormatType_4444YpCbCrA8R: Component Y’CbCrA 8-bit 4:4:4:4, rendering format.
	KCVPixelFormatType_4444YpCbCrA8R KCVPixelFormatType = 'r'<<24 | '4'<<16 | '0'<<8 | '8' // 'r408'
	// KCVPixelFormatType_444YpCbCr10: Component Y’CbCr 10-bit 4:4:4.
	KCVPixelFormatType_444YpCbCr10                         KCVPixelFormatType = 'v'<<24 | '4'<<16 | '1'<<8 | '0' // 'v410'
	KCVPixelFormatType_444YpCbCr10BiPlanarFullRange        KCVPixelFormatType = 'x'<<24 | 'f'<<16 | '4'<<8 | '4' // 'xf44'
	KCVPixelFormatType_444YpCbCr10BiPlanarVideoRange       KCVPixelFormatType = 'x'<<24 | '4'<<16 | '4'<<8 | '4' // 'x444'
	KCVPixelFormatType_444YpCbCr16BiPlanarVideoRange       KCVPixelFormatType = 's'<<24 | 'v'<<16 | '4'<<8 | '4' // 'sv44'
	KCVPixelFormatType_444YpCbCr16VideoRange_16A_TriPlanar KCVPixelFormatType = 's'<<24 | '4'<<16 | 'a'<<8 | 's' // 's4as'
	// KCVPixelFormatType_444YpCbCr8: Component Y’CbCr 8-bit 4:4:4.
	KCVPixelFormatType_444YpCbCr8                   KCVPixelFormatType = 'v'<<24 | '3'<<16 | '0'<<8 | '8' // 'v308'
	KCVPixelFormatType_444YpCbCr8BiPlanarFullRange  KCVPixelFormatType = '4'<<24 | '4'<<16 | '4'<<8 | 'f' // '444f'
	KCVPixelFormatType_444YpCbCr8BiPlanarVideoRange KCVPixelFormatType = '4'<<24 | '4'<<16 | '4'<<8 | 'v' // '444v'
	// KCVPixelFormatType_48RGB: 48-bit RGB, 16-bit big-endian samples.
	KCVPixelFormatType_48RGB KCVPixelFormatType = 'b'<<24 | '4'<<16 | '8'<<8 | 'r' // 'b48r'
	// KCVPixelFormatType_4Indexed: 4-bit indexed.
	KCVPixelFormatType_4Indexed KCVPixelFormatType = 0x4
	// KCVPixelFormatType_4IndexedGray_WhiteIsZero: 4-bit indexed gray, white is zero.
	KCVPixelFormatType_4IndexedGray_WhiteIsZero KCVPixelFormatType = 0x24
	// KCVPixelFormatType_64ARGB: 64-bit ARGB, 16-bit big-endian samples.
	KCVPixelFormatType_64ARGB KCVPixelFormatType = 'b'<<24 | '6'<<16 | '4'<<8 | 'a' // 'b64a'
	// KCVPixelFormatType_64RGBAHalf: 64-bit RGBA IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_64RGBAHalf                 KCVPixelFormatType = 'R'<<24 | 'G'<<16 | 'h'<<8 | 'A' // 'RGhA'
	KCVPixelFormatType_64RGBALE                   KCVPixelFormatType = 'l'<<24 | '6'<<16 | '4'<<8 | 'r' // 'l64r'
	KCVPixelFormatType_64RGBA_DownscaledProResRAW KCVPixelFormatType = 'b'<<24 | 'p'<<16 | '6'<<8 | '4' // 'bp64'
	// KCVPixelFormatType_8Indexed: 8-bit indexed.
	KCVPixelFormatType_8Indexed KCVPixelFormatType = 0x8
	// KCVPixelFormatType_8IndexedGray_WhiteIsZero: 8-bit indexed gray, white is zero.
	KCVPixelFormatType_8IndexedGray_WhiteIsZero KCVPixelFormatType = 0x28
	KCVPixelFormatType_96VersatileBayerPacked12 KCVPixelFormatType = 'b'<<24 | 't'<<16 | 'p'<<8 | '2' // 'btp2'
	KCVPixelFormatType_ARGB2101010LEPacked      KCVPixelFormatType = 'l'<<24 | '1'<<16 | '0'<<8 | 'r' // 'l10r'
	KCVPixelFormatType_DepthFloat16             KCVPixelFormatType = 'h'<<24 | 'd'<<16 | 'e'<<8 | 'p' // 'hdep'
	KCVPixelFormatType_DepthFloat32             KCVPixelFormatType = 'f'<<24 | 'd'<<16 | 'e'<<8 | 'p' // 'fdep'
	KCVPixelFormatType_DisparityFloat16         KCVPixelFormatType = 'h'<<24 | 'd'<<16 | 'i'<<8 | 's' // 'hdis'
	KCVPixelFormatType_DisparityFloat32         KCVPixelFormatType = 'f'<<24 | 'd'<<16 | 'i'<<8 | 's' // 'fdis'
	KCVPixelFormatType_OneComponent10           KCVPixelFormatType = 'L'<<24 | '0'<<16 | '1'<<8 | '0' // 'L010'
	KCVPixelFormatType_OneComponent12           KCVPixelFormatType = 'L'<<24 | '0'<<16 | '1'<<8 | '2' // 'L012'
	KCVPixelFormatType_OneComponent16           KCVPixelFormatType = 'L'<<24 | '0'<<16 | '1'<<8 | '6' // 'L016'
	// KCVPixelFormatType_OneComponent16Half: 6-bit one component IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_OneComponent16Half KCVPixelFormatType = 'L'<<24 | '0'<<16 | '0'<<8 | 'h' // 'L00h'
	// KCVPixelFormatType_OneComponent32Float: 32-bit one component IEEE float, 32-bit little-endian samples.
	KCVPixelFormatType_OneComponent32Float KCVPixelFormatType = 'L'<<24 | '0'<<16 | '0'<<8 | 'f' // 'L00f'
	// KCVPixelFormatType_OneComponent8: 8-bit one component, black is zero.
	KCVPixelFormatType_OneComponent8  KCVPixelFormatType = 'L'<<24 | '0'<<16 | '0'<<8 | '8' // 'L008'
	KCVPixelFormatType_TwoComponent16 KCVPixelFormatType = '2'<<24 | 'C'<<16 | '1'<<8 | '6' // '2C16'
	// KCVPixelFormatType_TwoComponent16Half: 16-bit two component IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_TwoComponent16Half KCVPixelFormatType = '2'<<24 | 'C'<<16 | '0'<<8 | 'h' // '2C0h'
	// KCVPixelFormatType_TwoComponent32Float: 32-bit two component IEEE float, 32-bit little-endian samples.
	KCVPixelFormatType_TwoComponent32Float KCVPixelFormatType = '2'<<24 | 'C'<<16 | '0'<<8 | 'f' // '2C0f'
	// KCVPixelFormatType_TwoComponent8: 8-bit two component, black is zero.
	KCVPixelFormatType_TwoComponent8 KCVPixelFormatType = '2'<<24 | 'C'<<16 | '0'<<8 | '8' // '2C08'
)

func (e KCVPixelFormatType) String() string {
	switch e {
	case KCVPixelFormatType_128RGBAFloat:
		return "KCVPixelFormatType_128RGBAFloat"
	case KCVPixelFormatType_14Bayer_BGGR:
		return "KCVPixelFormatType_14Bayer_BGGR"
	case KCVPixelFormatType_14Bayer_GBRG:
		return "KCVPixelFormatType_14Bayer_GBRG"
	case KCVPixelFormatType_14Bayer_GRBG:
		return "KCVPixelFormatType_14Bayer_GRBG"
	case KCVPixelFormatType_14Bayer_RGGB:
		return "KCVPixelFormatType_14Bayer_RGGB"
	case KCVPixelFormatType_16BE555:
		return "KCVPixelFormatType_16BE555"
	case KCVPixelFormatType_16BE565:
		return "KCVPixelFormatType_16BE565"
	case KCVPixelFormatType_16Gray:
		return "KCVPixelFormatType_16Gray"
	case KCVPixelFormatType_16LE555:
		return "KCVPixelFormatType_16LE555"
	case KCVPixelFormatType_16LE5551:
		return "KCVPixelFormatType_16LE5551"
	case KCVPixelFormatType_16LE565:
		return "KCVPixelFormatType_16LE565"
	case KCVPixelFormatType_16VersatileBayer:
		return "KCVPixelFormatType_16VersatileBayer"
	case KCVPixelFormatType_1IndexedGray_WhiteIsZero:
		return "KCVPixelFormatType_1IndexedGray_WhiteIsZero"
	case KCVPixelFormatType_1Monochrome:
		return "KCVPixelFormatType_1Monochrome"
	case KCVPixelFormatType_24BGR:
		return "KCVPixelFormatType_24BGR"
	case KCVPixelFormatType_24RGB:
		return "KCVPixelFormatType_24RGB"
	case KCVPixelFormatType_2Indexed:
		return "KCVPixelFormatType_2Indexed"
	case KCVPixelFormatType_2IndexedGray_WhiteIsZero:
		return "KCVPixelFormatType_2IndexedGray_WhiteIsZero"
	case KCVPixelFormatType_30RGB:
		return "KCVPixelFormatType_30RGB"
	case KCVPixelFormatType_30RGBLEPackedWideGamut:
		return "KCVPixelFormatType_30RGBLEPackedWideGamut"
	case KCVPixelFormatType_30RGBLE_8A_BiPlanar:
		return "KCVPixelFormatType_30RGBLE_8A_BiPlanar"
	case KCVPixelFormatType_30RGB_r210:
		return "KCVPixelFormatType_30RGB_r210"
	case KCVPixelFormatType_32ABGR:
		return "KCVPixelFormatType_32ABGR"
	case KCVPixelFormatType_32ARGB:
		return "KCVPixelFormatType_32ARGB"
	case KCVPixelFormatType_32AlphaGray:
		return "KCVPixelFormatType_32AlphaGray"
	case KCVPixelFormatType_32BGRA:
		return "KCVPixelFormatType_32BGRA"
	case KCVPixelFormatType_32RGBA:
		return "KCVPixelFormatType_32RGBA"
	case KCVPixelFormatType_40ARGBLEWideGamut:
		return "KCVPixelFormatType_40ARGBLEWideGamut"
	case KCVPixelFormatType_40ARGBLEWideGamutPremultiplied:
		return "KCVPixelFormatType_40ARGBLEWideGamutPremultiplied"
	case KCVPixelFormatType_420YpCbCr10BiPlanarFullRange:
		return "KCVPixelFormatType_420YpCbCr10BiPlanarFullRange"
	case KCVPixelFormatType_420YpCbCr10BiPlanarVideoRange:
		return "KCVPixelFormatType_420YpCbCr10BiPlanarVideoRange"
	case KCVPixelFormatType_420YpCbCr8BiPlanarFullRange:
		return "KCVPixelFormatType_420YpCbCr8BiPlanarFullRange"
	case KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange:
		return "KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange"
	case KCVPixelFormatType_420YpCbCr8Planar:
		return "KCVPixelFormatType_420YpCbCr8Planar"
	case KCVPixelFormatType_420YpCbCr8PlanarFullRange:
		return "KCVPixelFormatType_420YpCbCr8PlanarFullRange"
	case KCVPixelFormatType_420YpCbCr8VideoRange_8A_TriPlanar:
		return "KCVPixelFormatType_420YpCbCr8VideoRange_8A_TriPlanar"
	case KCVPixelFormatType_422YpCbCr10:
		return "KCVPixelFormatType_422YpCbCr10"
	case KCVPixelFormatType_422YpCbCr10BiPlanarFullRange:
		return "KCVPixelFormatType_422YpCbCr10BiPlanarFullRange"
	case KCVPixelFormatType_422YpCbCr10BiPlanarVideoRange:
		return "KCVPixelFormatType_422YpCbCr10BiPlanarVideoRange"
	case KCVPixelFormatType_422YpCbCr16:
		return "KCVPixelFormatType_422YpCbCr16"
	case KCVPixelFormatType_422YpCbCr16BiPlanarVideoRange:
		return "KCVPixelFormatType_422YpCbCr16BiPlanarVideoRange"
	case KCVPixelFormatType_422YpCbCr8:
		return "KCVPixelFormatType_422YpCbCr8"
	case KCVPixelFormatType_422YpCbCr8BiPlanarFullRange:
		return "KCVPixelFormatType_422YpCbCr8BiPlanarFullRange"
	case KCVPixelFormatType_422YpCbCr8BiPlanarVideoRange:
		return "KCVPixelFormatType_422YpCbCr8BiPlanarVideoRange"
	case KCVPixelFormatType_422YpCbCr8FullRange:
		return "KCVPixelFormatType_422YpCbCr8FullRange"
	case KCVPixelFormatType_422YpCbCr8_yuvs:
		return "KCVPixelFormatType_422YpCbCr8_yuvs"
	case KCVPixelFormatType_422YpCbCr_4A_8BiPlanar:
		return "KCVPixelFormatType_422YpCbCr_4A_8BiPlanar"
	case KCVPixelFormatType_4444AYpCbCr16:
		return "KCVPixelFormatType_4444AYpCbCr16"
	case KCVPixelFormatType_4444AYpCbCr8:
		return "KCVPixelFormatType_4444AYpCbCr8"
	case KCVPixelFormatType_4444AYpCbCrFloat:
		return "KCVPixelFormatType_4444AYpCbCrFloat"
	case KCVPixelFormatType_4444YpCbCrA8:
		return "KCVPixelFormatType_4444YpCbCrA8"
	case KCVPixelFormatType_4444YpCbCrA8R:
		return "KCVPixelFormatType_4444YpCbCrA8R"
	case KCVPixelFormatType_444YpCbCr10:
		return "KCVPixelFormatType_444YpCbCr10"
	case KCVPixelFormatType_444YpCbCr10BiPlanarFullRange:
		return "KCVPixelFormatType_444YpCbCr10BiPlanarFullRange"
	case KCVPixelFormatType_444YpCbCr10BiPlanarVideoRange:
		return "KCVPixelFormatType_444YpCbCr10BiPlanarVideoRange"
	case KCVPixelFormatType_444YpCbCr16BiPlanarVideoRange:
		return "KCVPixelFormatType_444YpCbCr16BiPlanarVideoRange"
	case KCVPixelFormatType_444YpCbCr16VideoRange_16A_TriPlanar:
		return "KCVPixelFormatType_444YpCbCr16VideoRange_16A_TriPlanar"
	case KCVPixelFormatType_444YpCbCr8:
		return "KCVPixelFormatType_444YpCbCr8"
	case KCVPixelFormatType_444YpCbCr8BiPlanarFullRange:
		return "KCVPixelFormatType_444YpCbCr8BiPlanarFullRange"
	case KCVPixelFormatType_444YpCbCr8BiPlanarVideoRange:
		return "KCVPixelFormatType_444YpCbCr8BiPlanarVideoRange"
	case KCVPixelFormatType_48RGB:
		return "KCVPixelFormatType_48RGB"
	case KCVPixelFormatType_4Indexed:
		return "KCVPixelFormatType_4Indexed"
	case KCVPixelFormatType_4IndexedGray_WhiteIsZero:
		return "KCVPixelFormatType_4IndexedGray_WhiteIsZero"
	case KCVPixelFormatType_64ARGB:
		return "KCVPixelFormatType_64ARGB"
	case KCVPixelFormatType_64RGBAHalf:
		return "KCVPixelFormatType_64RGBAHalf"
	case KCVPixelFormatType_64RGBALE:
		return "KCVPixelFormatType_64RGBALE"
	case KCVPixelFormatType_64RGBA_DownscaledProResRAW:
		return "KCVPixelFormatType_64RGBA_DownscaledProResRAW"
	case KCVPixelFormatType_8Indexed:
		return "KCVPixelFormatType_8Indexed"
	case KCVPixelFormatType_8IndexedGray_WhiteIsZero:
		return "KCVPixelFormatType_8IndexedGray_WhiteIsZero"
	case KCVPixelFormatType_96VersatileBayerPacked12:
		return "KCVPixelFormatType_96VersatileBayerPacked12"
	case KCVPixelFormatType_ARGB2101010LEPacked:
		return "KCVPixelFormatType_ARGB2101010LEPacked"
	case KCVPixelFormatType_DepthFloat16:
		return "KCVPixelFormatType_DepthFloat16"
	case KCVPixelFormatType_DepthFloat32:
		return "KCVPixelFormatType_DepthFloat32"
	case KCVPixelFormatType_DisparityFloat16:
		return "KCVPixelFormatType_DisparityFloat16"
	case KCVPixelFormatType_DisparityFloat32:
		return "KCVPixelFormatType_DisparityFloat32"
	case KCVPixelFormatType_OneComponent10:
		return "KCVPixelFormatType_OneComponent10"
	case KCVPixelFormatType_OneComponent12:
		return "KCVPixelFormatType_OneComponent12"
	case KCVPixelFormatType_OneComponent16:
		return "KCVPixelFormatType_OneComponent16"
	case KCVPixelFormatType_OneComponent16Half:
		return "KCVPixelFormatType_OneComponent16Half"
	case KCVPixelFormatType_OneComponent32Float:
		return "KCVPixelFormatType_OneComponent32Float"
	case KCVPixelFormatType_OneComponent8:
		return "KCVPixelFormatType_OneComponent8"
	case KCVPixelFormatType_TwoComponent16:
		return "KCVPixelFormatType_TwoComponent16"
	case KCVPixelFormatType_TwoComponent16Half:
		return "KCVPixelFormatType_TwoComponent16Half"
	case KCVPixelFormatType_TwoComponent32Float:
		return "KCVPixelFormatType_TwoComponent32Float"
	case KCVPixelFormatType_TwoComponent8:
		return "KCVPixelFormatType_TwoComponent8"
	default:
		return fmt.Sprintf("KCVPixelFormatType(%d)", e)
	}
}

type KCVReturn int

const (
	// KCVReturnAllocationFailed: Memory allocation for a buffer or buffer pool failed.
	KCVReturnAllocationFailed KCVReturn = -6662
	// KCVReturnDisplayLinkAlreadyRunning: The specified display link is already running.
	KCVReturnDisplayLinkAlreadyRunning KCVReturn = -6671
	// KCVReturnDisplayLinkCallbacksNotSet: No callback registered for the specified display link.
	KCVReturnDisplayLinkCallbacksNotSet KCVReturn = -6673
	// KCVReturnDisplayLinkNotRunning: The specified display link is not running.
	KCVReturnDisplayLinkNotRunning KCVReturn = -6672
	// KCVReturnError: An otherwise undefined error occurred.
	KCVReturnError KCVReturn = -6660
	// KCVReturnFirst: Placeholder to mark the beginning of Core Video result codes (not returned by any functions).
	KCVReturnFirst KCVReturn = -6660
	// KCVReturnInvalidArgument: Invalid function parameter.
	KCVReturnInvalidArgument KCVReturn = -6661
	// KCVReturnInvalidDisplay: The display specified when creating a display link is invalid.
	KCVReturnInvalidDisplay KCVReturn = -6670
	// KCVReturnInvalidPixelBufferAttributes: A buffer cannot be created with the specified attributes.
	KCVReturnInvalidPixelBufferAttributes KCVReturn = -6682
	// KCVReturnInvalidPixelFormat: The buffer does not support the specified pixel format.
	KCVReturnInvalidPixelFormat KCVReturn = -6680
	// KCVReturnInvalidPoolAttributes: A buffer pool cannot be created with the specified attributes.
	KCVReturnInvalidPoolAttributes KCVReturn = -6691
	// KCVReturnInvalidSize: The buffer cannot support the requested buffer size (usually too big).
	KCVReturnInvalidSize KCVReturn = -6681
	// KCVReturnLast: Placeholder to mark the end of Core Video result codes (not returned by any functions).
	KCVReturnLast KCVReturn = -6699
	// KCVReturnPixelBufferNotMetalCompatible: The pixel buffer is not compatible with Metal due to an unsupported buffer size, pixel format, or attribute.
	KCVReturnPixelBufferNotMetalCompatible KCVReturn = -6684
	// KCVReturnPixelBufferNotOpenGLCompatible: The pixel buffer is not compatible with OpenGL due to an unsupported buffer size, pixel format, or attribute.
	KCVReturnPixelBufferNotOpenGLCompatible KCVReturn = -6683
	// KCVReturnPoolAllocationFailed: Allocation for a buffer pool failed, most likely due to a lack of resources.
	KCVReturnPoolAllocationFailed KCVReturn = -6690
	// KCVReturnRetry: A scan hasn’t completely traversed the [CVBufferPool] due to a concurrent operation.
	KCVReturnRetry KCVReturn = -6692
	// KCVReturnSuccess: Indicates the operation completed successfully.
	KCVReturnSuccess     KCVReturn = 0
	KCVReturnUnsupported KCVReturn = -6663
	// KCVReturnWouldExceedAllocationThreshold: Allocation for a pixel buffer failed because the threshold value set for the kCVPixelBufferPoolAllocationThresholdKey key in the CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:) function would be surpassed.
	KCVReturnWouldExceedAllocationThreshold KCVReturn = -6689
)

func (e KCVReturn) String() string {
	switch e {
	case KCVReturnAllocationFailed:
		return "KCVReturnAllocationFailed"
	case KCVReturnDisplayLinkAlreadyRunning:
		return "KCVReturnDisplayLinkAlreadyRunning"
	case KCVReturnDisplayLinkCallbacksNotSet:
		return "KCVReturnDisplayLinkCallbacksNotSet"
	case KCVReturnDisplayLinkNotRunning:
		return "KCVReturnDisplayLinkNotRunning"
	case KCVReturnError:
		return "KCVReturnError"
	case KCVReturnInvalidArgument:
		return "KCVReturnInvalidArgument"
	case KCVReturnInvalidDisplay:
		return "KCVReturnInvalidDisplay"
	case KCVReturnInvalidPixelBufferAttributes:
		return "KCVReturnInvalidPixelBufferAttributes"
	case KCVReturnInvalidPixelFormat:
		return "KCVReturnInvalidPixelFormat"
	case KCVReturnInvalidPoolAttributes:
		return "KCVReturnInvalidPoolAttributes"
	case KCVReturnInvalidSize:
		return "KCVReturnInvalidSize"
	case KCVReturnLast:
		return "KCVReturnLast"
	case KCVReturnPixelBufferNotMetalCompatible:
		return "KCVReturnPixelBufferNotMetalCompatible"
	case KCVReturnPixelBufferNotOpenGLCompatible:
		return "KCVReturnPixelBufferNotOpenGLCompatible"
	case KCVReturnPoolAllocationFailed:
		return "KCVReturnPoolAllocationFailed"
	case KCVReturnRetry:
		return "KCVReturnRetry"
	case KCVReturnSuccess:
		return "KCVReturnSuccess"
	case KCVReturnUnsupported:
		return "KCVReturnUnsupported"
	case KCVReturnWouldExceedAllocationThreshold:
		return "KCVReturnWouldExceedAllocationThreshold"
	default:
		return fmt.Sprintf("KCVReturn(%d)", e)
	}
}

type KcvpixelformattypeLossless uint

const (
	KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut              KcvpixelformattypeLossless = '&'<<24 | 'w'<<16 | '3'<<8 | 'r' // '&w3r'
	KCVPixelFormatType_Lossless_30RGBLE_8A_BiPlanar                 KcvpixelformattypeLossless = '&'<<24 | 'b'<<16 | '3'<<8 | '8' // '&b38'
	KCVPixelFormatType_Lossless_32BGRA                              KcvpixelformattypeLossless = '&'<<24 | 'B'<<16 | 'G'<<8 | 'A' // '&BGA'
	KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarFullRange  KcvpixelformattypeLossless = '&'<<24 | 'x'<<16 | 'f'<<8 | '0' // '&xf0'
	KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossless = '&'<<24 | 'x'<<16 | 'v'<<8 | '0' // '&xv0'
	KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarFullRange         KcvpixelformattypeLossless = '&'<<24 | '8'<<16 | 'f'<<8 | '0' // '&8f0'
	KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarVideoRange        KcvpixelformattypeLossless = '&'<<24 | '8'<<16 | 'v'<<8 | '0' // '&8v0'
	KCVPixelFormatType_Lossless_422YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossless = '&'<<24 | 'x'<<16 | 'v'<<8 | '2' // '&xv2'
	KCVPixelFormatType_Lossless_64RGBAHalf                          KcvpixelformattypeLossless = '&'<<24 | 'R'<<16 | 'h'<<8 | 'A' // '&RhA'
)

func (e KcvpixelformattypeLossless) String() string {
	switch e {
	case KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut:
		return "KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut"
	case KCVPixelFormatType_Lossless_30RGBLE_8A_BiPlanar:
		return "KCVPixelFormatType_Lossless_30RGBLE_8A_BiPlanar"
	case KCVPixelFormatType_Lossless_32BGRA:
		return "KCVPixelFormatType_Lossless_32BGRA"
	case KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarFullRange:
		return "KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarFullRange"
	case KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarVideoRange:
		return "KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarVideoRange"
	case KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarFullRange:
		return "KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarFullRange"
	case KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarVideoRange:
		return "KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarVideoRange"
	case KCVPixelFormatType_Lossless_422YpCbCr10PackedBiPlanarVideoRange:
		return "KCVPixelFormatType_Lossless_422YpCbCr10PackedBiPlanarVideoRange"
	case KCVPixelFormatType_Lossless_64RGBAHalf:
		return "KCVPixelFormatType_Lossless_64RGBAHalf"
	default:
		return fmt.Sprintf("KcvpixelformattypeLossless(%d)", e)
	}
}

type KcvpixelformattypeLossy uint

const (
	KCVPixelFormatType_Lossy_32BGRA                              KcvpixelformattypeLossy = '-'<<24 | 'B'<<16 | 'G'<<8 | 'A' // '-BGA'
	KCVPixelFormatType_Lossy_420YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossy = '-'<<24 | 'x'<<16 | 'v'<<8 | '0' // '-xv0'
	KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarFullRange         KcvpixelformattypeLossy = '-'<<24 | '8'<<16 | 'f'<<8 | '0' // '-8f0'
	KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarVideoRange        KcvpixelformattypeLossy = '-'<<24 | '8'<<16 | 'v'<<8 | '0' // '-8v0'
	KCVPixelFormatType_Lossy_422YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossy = '-'<<24 | 'x'<<16 | 'v'<<8 | '2' // '-xv2'
)

func (e KcvpixelformattypeLossy) String() string {
	switch e {
	case KCVPixelFormatType_Lossy_32BGRA:
		return "KCVPixelFormatType_Lossy_32BGRA"
	case KCVPixelFormatType_Lossy_420YpCbCr10PackedBiPlanarVideoRange:
		return "KCVPixelFormatType_Lossy_420YpCbCr10PackedBiPlanarVideoRange"
	case KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarFullRange:
		return "KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarFullRange"
	case KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarVideoRange:
		return "KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarVideoRange"
	case KCVPixelFormatType_Lossy_422YpCbCr10PackedBiPlanarVideoRange:
		return "KCVPixelFormatType_Lossy_422YpCbCr10PackedBiPlanarVideoRange"
	default:
		return fmt.Sprintf("KcvpixelformattypeLossy(%d)", e)
	}
}

type KcvversatilebayerBayerpattern uint

const (
	KCVVersatileBayer_BayerPattern_BGGR KcvversatilebayerBayerpattern = 3
	KCVVersatileBayer_BayerPattern_GBRG KcvversatilebayerBayerpattern = 2
	KCVVersatileBayer_BayerPattern_GRBG KcvversatilebayerBayerpattern = 1
	KCVVersatileBayer_BayerPattern_RGGB KcvversatilebayerBayerpattern = 0
)

func (e KcvversatilebayerBayerpattern) String() string {
	switch e {
	case KCVVersatileBayer_BayerPattern_BGGR:
		return "KCVVersatileBayer_BayerPattern_BGGR"
	case KCVVersatileBayer_BayerPattern_GBRG:
		return "KCVVersatileBayer_BayerPattern_GBRG"
	case KCVVersatileBayer_BayerPattern_GRBG:
		return "KCVVersatileBayer_BayerPattern_GRBG"
	case KCVVersatileBayer_BayerPattern_RGGB:
		return "KCVVersatileBayer_BayerPattern_RGGB"
	default:
		return fmt.Sprintf("KcvversatilebayerBayerpattern(%d)", e)
	}
}
