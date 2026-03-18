// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreVideo/CVAttachmentMode
type CVAttachmentMode int

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
	KCVPixelBufferLock_ReadOnly CVPixelBufferLockFlags = 1
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
type CVSMPTETimeFlags int

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
type CVSMPTETimeType int

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
type CVTimeFlags int

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
type CVTimeStampFlags int

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
	KCVPixelFormatType_128RGBAFloat KCVPixelFormatType = 0
	KCVPixelFormatType_14Bayer_BGGR KCVPixelFormatType = 0
	KCVPixelFormatType_14Bayer_GBRG KCVPixelFormatType = 0
	KCVPixelFormatType_14Bayer_GRBG KCVPixelFormatType = 0
	KCVPixelFormatType_14Bayer_RGGB KCVPixelFormatType = 0
	// KCVPixelFormatType_16BE555: 16-bit BE RGB 555.
	KCVPixelFormatType_16BE555 KCVPixelFormatType = 0
	// KCVPixelFormatType_16BE565: 16-bit BE RGB 565.
	KCVPixelFormatType_16BE565 KCVPixelFormatType = 0
	// KCVPixelFormatType_16Gray: 16-bit Grayscale, 16-bit big-endian samples, black is zero.
	KCVPixelFormatType_16Gray KCVPixelFormatType = 0
	// KCVPixelFormatType_16LE555: 16-bit LE RGB 555.
	KCVPixelFormatType_16LE555 KCVPixelFormatType = 0
	// KCVPixelFormatType_16LE5551: 16-bit LE RGB 5551.
	KCVPixelFormatType_16LE5551 KCVPixelFormatType = 0
	// KCVPixelFormatType_16LE565: 16-bit LE RGB 565.
	KCVPixelFormatType_16LE565 KCVPixelFormatType = 0
	KCVPixelFormatType_16VersatileBayer KCVPixelFormatType = 0
	// KCVPixelFormatType_1IndexedGray_WhiteIsZero: 1 bit indexed gray, white is zero.
	KCVPixelFormatType_1IndexedGray_WhiteIsZero KCVPixelFormatType = 0
	// KCVPixelFormatType_1Monochrome: 1 bit indexed.
	KCVPixelFormatType_1Monochrome KCVPixelFormatType = 0
	// KCVPixelFormatType_24BGR: 24-bit BGR.
	KCVPixelFormatType_24BGR KCVPixelFormatType = 0
	// KCVPixelFormatType_24RGB: 24-bit RGB.
	KCVPixelFormatType_24RGB KCVPixelFormatType = 0
	// KCVPixelFormatType_2Indexed: 2-bit indexed.
	KCVPixelFormatType_2Indexed KCVPixelFormatType = 0
	// KCVPixelFormatType_2IndexedGray_WhiteIsZero: 2-bit indexed gray, white is zero.
	KCVPixelFormatType_2IndexedGray_WhiteIsZero KCVPixelFormatType = 0
	// KCVPixelFormatType_30RGB: 30-bit RGB, 10-bit big-endian samples, 2 unused padding bits (at least significant end).
	KCVPixelFormatType_30RGB KCVPixelFormatType = 0
	KCVPixelFormatType_30RGBLEPackedWideGamut KCVPixelFormatType = 0
	KCVPixelFormatType_30RGBLE_8A_BiPlanar KCVPixelFormatType = 0
	KCVPixelFormatType_30RGB_r210 KCVPixelFormatType = 0
	// KCVPixelFormatType_32ABGR: 32-bit ABGR.
	KCVPixelFormatType_32ABGR KCVPixelFormatType = 0
	// KCVPixelFormatType_32ARGB: 32-bit ARGB.
	KCVPixelFormatType_32ARGB KCVPixelFormatType = 0
	// KCVPixelFormatType_32AlphaGray: 32-bit AlphaGray, 16-bit big-endian samples, black is zero.
	KCVPixelFormatType_32AlphaGray KCVPixelFormatType = 0
	// KCVPixelFormatType_32BGRA: 32-bit BGRA.
	KCVPixelFormatType_32BGRA KCVPixelFormatType = 0
	// KCVPixelFormatType_32RGBA: 32-bit RGBA.
	KCVPixelFormatType_32RGBA KCVPixelFormatType = 0
	KCVPixelFormatType_40ARGBLEWideGamut KCVPixelFormatType = 0
	KCVPixelFormatType_40ARGBLEWideGamutPremultiplied KCVPixelFormatType = 0
	KCVPixelFormatType_420YpCbCr10BiPlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_420YpCbCr10BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_420YpCbCr8BiPlanarFullRange: Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]).
	KCVPixelFormatType_420YpCbCr8BiPlanarFullRange KCVPixelFormatType = 0
	// KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange: Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]).
	KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_420YpCbCr8Planar: Planar Component Y’CbCr 8-bit 4:2:0.
	KCVPixelFormatType_420YpCbCr8Planar KCVPixelFormatType = 0
	// KCVPixelFormatType_420YpCbCr8PlanarFullRange: Planar Component Y’CbCr 8-bit 4:2:0, full range.
	KCVPixelFormatType_420YpCbCr8PlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_420YpCbCr8VideoRange_8A_TriPlanar KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr10: Component Y’CbCr 10-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr10 KCVPixelFormatType = 0
	KCVPixelFormatType_422YpCbCr10BiPlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_422YpCbCr10BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr16: Component Y’CbCr 10,12,14,16-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr16 KCVPixelFormatType = 0
	KCVPixelFormatType_422YpCbCr16BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr8: Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
	KCVPixelFormatType_422YpCbCr8 KCVPixelFormatType = 0
	KCVPixelFormatType_422YpCbCr8BiPlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_422YpCbCr8BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr8FullRange: Component Y’CbCr 8-bit 4:2:2, full range, ordered Y’0 Cb Y’1 Cr.
	KCVPixelFormatType_422YpCbCr8FullRange KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr8_yuvs: Component Y’CbCr 8-bit 4:2:2, ordered Y’0 Cb Y’1 Cr.
	KCVPixelFormatType_422YpCbCr8_yuvs KCVPixelFormatType = 0
	// KCVPixelFormatType_422YpCbCr_4A_8BiPlanar: First plane: Video-range Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1; second plane: alpha 8-bit 0-255.
	KCVPixelFormatType_422YpCbCr_4A_8BiPlanar KCVPixelFormatType = 0
	// KCVPixelFormatType_4444AYpCbCr16: Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
	KCVPixelFormatType_4444AYpCbCr16 KCVPixelFormatType = 0
	// KCVPixelFormatType_4444AYpCbCr8: Component Y’CbCrA 8-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr.
	KCVPixelFormatType_4444AYpCbCr8 KCVPixelFormatType = 0
	KCVPixelFormatType_4444AYpCbCrFloat KCVPixelFormatType = 0
	// KCVPixelFormatType_4444YpCbCrA8: Component Y’CbCrA 8-bit 4:4:4:4, ordered Cb Y’ Cr A.
	KCVPixelFormatType_4444YpCbCrA8 KCVPixelFormatType = 0
	// KCVPixelFormatType_4444YpCbCrA8R: Component Y’CbCrA 8-bit 4:4:4:4, rendering format.
	KCVPixelFormatType_4444YpCbCrA8R KCVPixelFormatType = 0
	// KCVPixelFormatType_444YpCbCr10: Component Y’CbCr 10-bit 4:4:4.
	KCVPixelFormatType_444YpCbCr10 KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr10BiPlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr10BiPlanarVideoRange KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr16BiPlanarVideoRange KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr16VideoRange_16A_TriPlanar KCVPixelFormatType = 0
	// KCVPixelFormatType_444YpCbCr8: Component Y’CbCr 8-bit 4:4:4.
	KCVPixelFormatType_444YpCbCr8 KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr8BiPlanarFullRange KCVPixelFormatType = 0
	KCVPixelFormatType_444YpCbCr8BiPlanarVideoRange KCVPixelFormatType = 0
	// KCVPixelFormatType_48RGB: 48-bit RGB, 16-bit big-endian samples.
	KCVPixelFormatType_48RGB KCVPixelFormatType = 0
	// KCVPixelFormatType_4Indexed: 4-bit indexed.
	KCVPixelFormatType_4Indexed KCVPixelFormatType = 0
	// KCVPixelFormatType_4IndexedGray_WhiteIsZero: 4-bit indexed gray, white is zero.
	KCVPixelFormatType_4IndexedGray_WhiteIsZero KCVPixelFormatType = 0
	// KCVPixelFormatType_64ARGB: 64-bit ARGB, 16-bit big-endian samples.
	KCVPixelFormatType_64ARGB KCVPixelFormatType = 0
	// KCVPixelFormatType_64RGBAHalf: 64-bit RGBA IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_64RGBAHalf KCVPixelFormatType = 0
	KCVPixelFormatType_64RGBALE KCVPixelFormatType = 0
	KCVPixelFormatType_64RGBA_DownscaledProResRAW KCVPixelFormatType = 0
	// KCVPixelFormatType_8Indexed: 8-bit indexed.
	KCVPixelFormatType_8Indexed KCVPixelFormatType = 0
	// KCVPixelFormatType_8IndexedGray_WhiteIsZero: 8-bit indexed gray, white is zero.
	KCVPixelFormatType_8IndexedGray_WhiteIsZero KCVPixelFormatType = 0
	KCVPixelFormatType_96VersatileBayerPacked12 KCVPixelFormatType = 0
	KCVPixelFormatType_ARGB2101010LEPacked KCVPixelFormatType = 0
	KCVPixelFormatType_DepthFloat16 KCVPixelFormatType = 0
	KCVPixelFormatType_DepthFloat32 KCVPixelFormatType = 0
	KCVPixelFormatType_DisparityFloat16 KCVPixelFormatType = 0
	KCVPixelFormatType_DisparityFloat32 KCVPixelFormatType = 0
	KCVPixelFormatType_OneComponent10 KCVPixelFormatType = 0
	KCVPixelFormatType_OneComponent12 KCVPixelFormatType = 0
	KCVPixelFormatType_OneComponent16 KCVPixelFormatType = 0
	// KCVPixelFormatType_OneComponent16Half: 6-bit one component IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_OneComponent16Half KCVPixelFormatType = 0
	// KCVPixelFormatType_OneComponent32Float: 32-bit one component IEEE float, 32-bit little-endian samples.
	KCVPixelFormatType_OneComponent32Float KCVPixelFormatType = 0
	// KCVPixelFormatType_OneComponent8: 8-bit one component, black is zero.
	KCVPixelFormatType_OneComponent8 KCVPixelFormatType = 0
	KCVPixelFormatType_TwoComponent16 KCVPixelFormatType = 0
	// KCVPixelFormatType_TwoComponent16Half: 16-bit two component IEEE half-precision float, 16-bit little-endian samples.
	KCVPixelFormatType_TwoComponent16Half KCVPixelFormatType = 0
	// KCVPixelFormatType_TwoComponent32Float: 32-bit two component IEEE float, 32-bit little-endian samples.
	KCVPixelFormatType_TwoComponent32Float KCVPixelFormatType = 0
	// KCVPixelFormatType_TwoComponent8: 8-bit two component, black is zero.
	KCVPixelFormatType_TwoComponent8 KCVPixelFormatType = 0
)

func (e KCVPixelFormatType) String() string {
	switch e {
	case KCVPixelFormatType_128RGBAFloat:
		return "KCVPixelFormatType_128RGBAFloat"
	default:
		return fmt.Sprintf("KCVPixelFormatType(%d)", e)
	}
}

type KCVReturn uint

const (
	// KCVReturnAllocationFailed: Memory allocation for a buffer or buffer pool failed.
	KCVReturnAllocationFailed KCVReturn = 0
	// KCVReturnDisplayLinkAlreadyRunning: The specified display link is already running.
	KCVReturnDisplayLinkAlreadyRunning KCVReturn = 0
	// KCVReturnDisplayLinkCallbacksNotSet: No callback registered for the specified display link.
	KCVReturnDisplayLinkCallbacksNotSet KCVReturn = 0
	// KCVReturnDisplayLinkNotRunning: The specified display link is not running.
	KCVReturnDisplayLinkNotRunning KCVReturn = 0
	// KCVReturnError: An otherwise undefined error occurred.
	KCVReturnError KCVReturn = 0
	// KCVReturnFirst: Placeholder to mark the beginning of Core Video result codes (not returned by any functions).
	KCVReturnFirst KCVReturn = 0
	// KCVReturnInvalidArgument: Invalid function parameter.
	KCVReturnInvalidArgument KCVReturn = 0
	// KCVReturnInvalidDisplay: The display specified when creating a display link is invalid.
	KCVReturnInvalidDisplay KCVReturn = 0
	// KCVReturnInvalidPixelBufferAttributes: A buffer cannot be created with the specified attributes.
	KCVReturnInvalidPixelBufferAttributes KCVReturn = 0
	// KCVReturnInvalidPixelFormat: The buffer does not support the specified pixel format.
	KCVReturnInvalidPixelFormat KCVReturn = 0
	// KCVReturnInvalidPoolAttributes: A buffer pool cannot be created with the specified attributes.
	KCVReturnInvalidPoolAttributes KCVReturn = 0
	// KCVReturnInvalidSize: The buffer cannot support the requested buffer size (usually too big).
	KCVReturnInvalidSize KCVReturn = 0
	// KCVReturnLast: Placeholder to mark the end of Core Video result codes (not returned by any functions).
	KCVReturnLast KCVReturn = 0
	// KCVReturnPixelBufferNotMetalCompatible: The pixel buffer is not compatible with Metal due to an unsupported buffer size, pixel format, or attribute.
	KCVReturnPixelBufferNotMetalCompatible KCVReturn = 0
	// KCVReturnPixelBufferNotOpenGLCompatible: The pixel buffer is not compatible with OpenGL due to an unsupported buffer size, pixel format, or attribute.
	KCVReturnPixelBufferNotOpenGLCompatible KCVReturn = 0
	// KCVReturnPoolAllocationFailed: Allocation for a buffer pool failed, most likely due to a lack of resources.
	KCVReturnPoolAllocationFailed KCVReturn = 0
	// KCVReturnRetry: A scan hasn’t completely traversed the [CVBufferPool] due to a concurrent operation.
	KCVReturnRetry KCVReturn = 0
	// KCVReturnSuccess: Indicates the operation completed successfully.
	KCVReturnSuccess KCVReturn = 0
	KCVReturnUnsupported KCVReturn = 0
	// KCVReturnWouldExceedAllocationThreshold: Allocation for a pixel buffer failed because the threshold value set for the kCVPixelBufferPoolAllocationThresholdKey key in the CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:) function would be surpassed.
	KCVReturnWouldExceedAllocationThreshold KCVReturn = 0
)

func (e KCVReturn) String() string {
	switch e {
	case KCVReturnAllocationFailed:
		return "KCVReturnAllocationFailed"
	default:
		return fmt.Sprintf("KCVReturn(%d)", e)
	}
}

type KcvpixelformattypeLossless uint

const (
	KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_30RGBLE_8A_BiPlanar KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_32BGRA KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarFullRange KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_420YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarFullRange KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_420YpCbCr8BiPlanarVideoRange KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_422YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossless = 0
	KCVPixelFormatType_Lossless_64RGBAHalf KcvpixelformattypeLossless = 0
)

func (e KcvpixelformattypeLossless) String() string {
	switch e {
	case KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut:
		return "KCVPixelFormatType_Lossless_30RGBLEPackedWideGamut"
	default:
		return fmt.Sprintf("KcvpixelformattypeLossless(%d)", e)
	}
}

type KcvpixelformattypeLossy uint

const (
	KCVPixelFormatType_Lossy_32BGRA KcvpixelformattypeLossy = 0
	KCVPixelFormatType_Lossy_420YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossy = 0
	KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarFullRange KcvpixelformattypeLossy = 0
	KCVPixelFormatType_Lossy_420YpCbCr8BiPlanarVideoRange KcvpixelformattypeLossy = 0
	KCVPixelFormatType_Lossy_422YpCbCr10PackedBiPlanarVideoRange KcvpixelformattypeLossy = 0
)

func (e KcvpixelformattypeLossy) String() string {
	switch e {
	case KCVPixelFormatType_Lossy_32BGRA:
		return "KCVPixelFormatType_Lossy_32BGRA"
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

