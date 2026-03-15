// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Accelerate/BLAS_THREADING
type BLAS_THREADING int

const (
	BLAS_THREADING_MAX_OPTIONS BLAS_THREADING = 2
	// BLAS_THREADING_MULTI_THREADED: A constant that specifies that the Accelerate framework decides whether BLAS and LAPACK execute on single or multiple threads.
	BLAS_THREADING_MULTI_THREADED BLAS_THREADING = 0
	// BLAS_THREADING_SINGLE_THREADED: A constant that specifies BLAS and LAPACK execute on a single thread only.
	BLAS_THREADING_SINGLE_THREADED BLAS_THREADING = 1
)


func (e BLAS_THREADING) String() string {
	switch e {
	case BLAS_THREADING_MAX_OPTIONS:
		return "BLAS_THREADING_MAX_OPTIONS"
	case BLAS_THREADING_MULTI_THREADED:
		return "BLAS_THREADING_MULTI_THREADED"
	case BLAS_THREADING_SINGLE_THREADED:
		return "BLAS_THREADING_SINGLE_THREADED"
	default:
		return fmt.Sprintf("BLAS_THREADING(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/CBLAS_DIAG
type CBLAS_DIAG int

const (
	CblasNonUnit CBLAS_DIAG = 131
	CblasUnit CBLAS_DIAG = 132
)


func (e CBLAS_DIAG) String() string {
	switch e {
	case CblasNonUnit:
		return "CblasNonUnit"
	case CblasUnit:
		return "CblasUnit"
	default:
		return fmt.Sprintf("CBLAS_DIAG(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/CBLAS_ORDER
type CBLAS_ORDER int

const (
	CblasColMajor CBLAS_ORDER = 102
	CblasRowMajor CBLAS_ORDER = 101
)


func (e CBLAS_ORDER) String() string {
	switch e {
	case CblasColMajor:
		return "CblasColMajor"
	case CblasRowMajor:
		return "CblasRowMajor"
	default:
		return fmt.Sprintf("CBLAS_ORDER(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/CBLAS_SIDE
type CBLAS_SIDE int

const (
	CblasLeft CBLAS_SIDE = 141
	CblasRight CBLAS_SIDE = 142
)


func (e CBLAS_SIDE) String() string {
	switch e {
	case CblasLeft:
		return "CblasLeft"
	case CblasRight:
		return "CblasRight"
	default:
		return fmt.Sprintf("CBLAS_SIDE(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/CBLAS_TRANSPOSE
type CBLAS_TRANSPOSE int

const (
	AtlasConj CBLAS_TRANSPOSE = 114
	CblasConjTrans CBLAS_TRANSPOSE = 113
	CblasNoTrans CBLAS_TRANSPOSE = 111
	CblasTrans CBLAS_TRANSPOSE = 112
)


func (e CBLAS_TRANSPOSE) String() string {
	switch e {
	case AtlasConj:
		return "AtlasConj"
	case CblasConjTrans:
		return "CblasConjTrans"
	case CblasNoTrans:
		return "CblasNoTrans"
	case CblasTrans:
		return "CblasTrans"
	default:
		return fmt.Sprintf("CBLAS_TRANSPOSE(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/CBLAS_UPLO
type CBLAS_UPLO int

const (
	CblasLower CBLAS_UPLO = 122
	CblasUpper CBLAS_UPLO = 121
)


func (e CBLAS_UPLO) String() string {
	switch e {
	case CblasLower:
		return "CblasLower"
	case CblasUpper:
		return "CblasUpper"
	default:
		return fmt.Sprintf("CBLAS_UPLO(%d)", e)
	}
}




type Fft uint

const (
	// FFT_FORWARD: Forward FFT.
	FFT_FORWARD Fft = 0
	// FFT_INVERSE: Inverse FFT.
	FFT_INVERSE Fft = 0
	FFT_RADIX2 Fft = 0
	FFT_RADIX3 Fft = 1
	FFT_RADIX5 Fft = 0
)


func (e Fft) String() string {
	switch e {
	case FFT_FORWARD:
		return "FFT_FORWARD"
	case FFT_RADIX3:
		return "FFT_RADIX3"
	default:
		return fmt.Sprintf("Fft(%d)", e)
	}
}




type KFFT uint

const (
	KFFTRadix2 KFFT = 0
	KFFTRadix3 KFFT = 1
	KFFTRadix5 KFFT = 2
)


func (e KFFT) String() string {
	switch e {
	case KFFTRadix2:
		return "KFFTRadix2"
	case KFFTRadix3:
		return "KFFTRadix3"
	case KFFTRadix5:
		return "KFFTRadix5"
	default:
		return fmt.Sprintf("KFFT(%d)", e)
	}
}




type KFFTDirection int

const (
	KFFTDirection_Forward KFFTDirection = 0
	KFFTDirection_Inverse KFFTDirection = -1
)


func (e KFFTDirection) String() string {
	switch e {
	case KFFTDirection_Forward:
		return "KFFTDirection_Forward"
	case KFFTDirection_Inverse:
		return "KFFTDirection_Inverse"
	default:
		return fmt.Sprintf("KFFTDirection(%d)", e)
	}
}




type KRotate0DegreesClockwise uint

const (
	// KRotate0DegreesClockwiseValue: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesClockwiseValue KRotate0DegreesClockwise = 0
	// KRotate0DegreesCounterClockwise: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate180DegreesClockwise: A constant that specifies rotation by 180° clockwise.
	KRotate180DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate180DegreesCounterClockwise: A constant that specifies rotation by 180° counterclockwise.
	KRotate180DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate270DegreesClockwise: A constant that specifies rotation by 270° clockwise.
	KRotate270DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate270DegreesCounterClockwise: A constant that specifies rotation by 270° counterclockwise.
	KRotate270DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate90DegreesClockwise: A constant that specifies rotation by 90° clockwise.
	KRotate90DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate90DegreesCounterClockwise: A constant that specifies rotation by 90° counterclockwise.
	KRotate90DegreesCounterClockwise KRotate0DegreesClockwise = 0
)


func (e KRotate0DegreesClockwise) String() string {
	switch e {
	case KRotate0DegreesClockwiseValue:
		return "KRotate0DegreesClockwiseValue"
	default:
		return fmt.Sprintf("KRotate0DegreesClockwise(%d)", e)
	}
}




type KvImage uint

const (
	// KvImageBackgroundColorFill: A flag that uses the background color for missing pixels.
	KvImageBackgroundColorFill KvImage = 0
	// KvImageBufferSizeMismatch: The function requires the source and destination buffers to have the same height and the same width, but they do not.
	KvImageBufferSizeMismatch KvImage = 0
	KvImageColorSyncIsAbsent KvImage = 0
	// KvImageCopyInPlace: A flag that copies the value of the edge pixel in the source to the destination.
	KvImageCopyInPlace KvImage = 0
	KvImageCoreVideoIsAbsent KvImage = 0
	// KvImageDoNotClamp: A flag that disables clamping in some conversions to floating-point formats.
	KvImageDoNotClamp KvImage = 0
	// KvImageDoNotTile: A flag that disables vImage internal tiling routines.
	KvImageDoNotTile KvImage = 0
	// KvImageEdgeExtend: A flag that extends the edges of the image infinitely.
	KvImageEdgeExtend KvImage = 0
	// KvImageGetTempBufferSize: A flag that returns the minimum temporary buffer size for the operation, given the parameters provided.
	KvImageGetTempBufferSize KvImage = 0
	// KvImageHDRContent: A flag that uses HDR-aware methods.
	KvImageHDRContent KvImage = 0
	// KvImageHighQualityResampling: A flag that uses a higher-quality, slower resampling filter for geometry operations.
	KvImageHighQualityResampling KvImage = 0
	// KvImageInternalError: A serious error occured inside vImage, which prevented vImage from continuing.
	KvImageInternalError KvImage = 0
	KvImageInvalidCVImageFormat KvImage = 0
	// KvImageInvalidEdgeStyle: The edge style specified is invalid.
	KvImageInvalidEdgeStyle KvImage = 0
	KvImageInvalidImageFormat KvImage = 0
	KvImageInvalidImageObject KvImage = 0
	// KvImageInvalidKernelSize: Either the kernel height, the kernel width, or both, are even.
	KvImageInvalidKernelSize KvImage = 0
	// KvImageInvalidOffset_X: The `srcOffsetToROI_X` parameter that specifies the left edge of the region of interest is greater than the width of the source image.
	KvImageInvalidOffset_X KvImage = 0
	// KvImageInvalidOffset_Y: The `srcOffsetToROI_Y` parameter that specifies the top edge of the region of interest is greater than the height of the source image.
	KvImageInvalidOffset_Y KvImage = 0
	// KvImageInvalidParameter: Invalid parameter.
	KvImageInvalidParameter KvImage = 0
	KvImageInvalidRowBytes KvImage = 0
	// KvImageLeaveAlphaUnchanged: A flag that restricts the operation to red, green, and blue channels only.
	KvImageLeaveAlphaUnchanged KvImage = 0
	// KvImageMemoryAllocationError: An attempt to allocate memory failed.
	KvImageMemoryAllocationError KvImage = 0
	// KvImageNoAllocate: A flag that prevents vImage from allocating additional storage.
	KvImageNoAllocate KvImage = 0
	// KvImageNoError: The vImage function completed without error.
	KvImageNoError KvImage = 0
	// KvImageNoFlags: A flag that sets the behavior to the default.
	KvImageNoFlags KvImage = 0
	// KvImageNullPointerArgument: A pointer parameter is [NULL] and it must not be.
	KvImageNullPointerArgument KvImage = 0
	KvImageOutOfPlaceOperationRequired KvImage = 0
	// KvImagePrintDiagnosticsToConsole: A flag that prints a debug message if the operation fails.
	KvImagePrintDiagnosticsToConsole KvImage = 0
	// KvImageRoiLargerThanInputBuffer: The region of interest, as specified by the `srcOffsetToROI_X` and `srcOffsetToROI_Y` parameters and the height and width of the destination buffer, extends beyond the bottom edge or right edge of the source buffer.
	KvImageRoiLargerThanInputBuffer KvImage = 0
	// KvImageTruncateKernel: A flag that uses only the part of the kernel that overlaps the image.
	KvImageTruncateKernel KvImage = 0
	// KvImageUnknownFlagsBit: The flag is not recognized.
	KvImageUnknownFlagsBit KvImage = 0
	// KvImageUnsupportedConversion: Some lower level conversion APIs only support conversion among a sparse matrix of image formats.
	KvImageUnsupportedConversion KvImage = 0
	// KvImageUseFP16Accumulator: A flag that specifies vImage uses faster but lower-precision internal arithmetic for floating-point 16-bit operations.
	KvImageUseFP16Accumulator KvImage = 0
)


func (e KvImage) String() string {
	switch e {
	case KvImageBackgroundColorFill:
		return "KvImageBackgroundColorFill"
	default:
		return fmt.Sprintf("KvImage(%d)", e)
	}
}




type KvImageBufferTypeCode uint

const (
	// KvImageBufferTypeCode_Alpha: The buffer contains the alpha channel.
	KvImageBufferTypeCode_Alpha KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CGFormat: The buffer contains data describable as a vImage Core Graphics image format as a single buffer.
	KvImageBufferTypeCode_CGFormat KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Black: If the image has a CMYK color model, the buffer contains the black channel.
	KvImageBufferTypeCode_CMYK_Black KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Cyan: If the image has a CMYK color model, the buffer contains the cyan channel.
	KvImageBufferTypeCode_CMYK_Cyan KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Magenta: If the image has a CMYK color model, the buffer contains the magenta channel.
	KvImageBufferTypeCode_CMYK_Magenta KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Yellow: If the image has a CMYK color model, the buffer contains the yellow channel.
	KvImageBufferTypeCode_CMYK_Yellow KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CVPixelBuffer_YCbCr: The buffer contains luminance and both chroma channels interleaved according to the vImageConstCVImageFormat image type.
	KvImageBufferTypeCode_CVPixelBuffer_YCbCr KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Cb: The buffer contains the blue chrominance channel.
	KvImageBufferTypeCode_Cb KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Chroma: The buffer contains both chrominance channels, interleaved.
	KvImageBufferTypeCode_Chroma KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Chunky: The buffer contains chunky data not describable as a vImage Core Graphics image format.
	KvImageBufferTypeCode_Chunky KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel1 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel10 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel11 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel12 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel13 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel14 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel15 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel16 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel2 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel3 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel4 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel5 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel6 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel7 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel8 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel9 KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Cr: The buffer contains the red chrominance channel.
	KvImageBufferTypeCode_Cr KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_EndOfList: End of list marker.
	KvImageBufferTypeCode_EndOfList KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Indexed: The buffer contains data in an indexed colorspace.
	KvImageBufferTypeCode_Indexed KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_A: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_A KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_B: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_B KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_L: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_L KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Luminance: The buffer contains only luminance data.
	KvImageBufferTypeCode_Luminance KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Monochrome: The buffer contains a single color channel.
	KvImageBufferTypeCode_Monochrome KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Blue: If the image has a RGB color model, the buffer contains the blue channel.
	KvImageBufferTypeCode_RGB_Blue KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Green: If the image has a RGB color model, the buffer contains the green channel.
	KvImageBufferTypeCode_RGB_Green KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Red: If the image has a RGB color model, the buffer contains the red channel.
	KvImageBufferTypeCode_RGB_Red KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_UniqueFormatCount KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_X: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_X KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_Y: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_Y KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_Z: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_Z KvImageBufferTypeCode = 0
)


func (e KvImageBufferTypeCode) String() string {
	switch e {
	case KvImageBufferTypeCode_Alpha:
		return "KvImageBufferTypeCode_Alpha"
	default:
		return fmt.Sprintf("KvImageBufferTypeCode(%d)", e)
	}
}




type KvImageCVImageFormat uint

const (
	// KvImageCVImageFormat_AlphaIsOneHint: A hint that indicates the alpha channel is opaque.
	KvImageCVImageFormat_AlphaIsOneHint KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ChromaSiting: An error code that indicates the chroma siting information is absent.
	KvImageCVImageFormat_ChromaSiting KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ColorSpace: An error code that indicates the image’s color space is missing.
	KvImageCVImageFormat_ColorSpace KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ConversionMatrix: An error code that indicates the required conversion matrix is absent.
	KvImageCVImageFormat_ConversionMatrix KvImageCVImageFormat = 0
	// KvImageCVImageFormat_NoError: An error code that indicates the conversion completed without error.
	KvImageCVImageFormat_NoError KvImageCVImageFormat = 0
	// KvImageCVImageFormat_VideoChannelDescription: An error code that indicates the range and clipping information is missing.
	KvImageCVImageFormat_VideoChannelDescription KvImageCVImageFormat = 0
)


func (e KvImageCVImageFormat) String() string {
	switch e {
	case KvImageCVImageFormat_AlphaIsOneHint:
		return "KvImageCVImageFormat_AlphaIsOneHint"
	default:
		return fmt.Sprintf("KvImageCVImageFormat(%d)", e)
	}
}




type KvImageConvert uint

const (
	// KvImageConvert_DitherAtkinson: A constant that indicates the conversion will add Atkinson dithering to the image.
	KvImageConvert_DitherAtkinson KvImageConvert = 4
	// KvImageConvert_DitherFloydSteinberg: A constant that indicates the conversion will add Floyd-Steinberg dithering to the image.
	KvImageConvert_DitherFloydSteinberg KvImageConvert = 3
	// KvImageConvert_DitherNone: A constant that indicates the conversion will not apply dithering.
	KvImageConvert_DitherNone KvImageConvert = 0
	// KvImageConvert_DitherOrdered: A constant that indicates the conversion will add randomized, pre-computed blue noise to the image.
	KvImageConvert_DitherOrdered KvImageConvert = 1
	// KvImageConvert_DitherOrderedReproducible: A constant that indicates the conversion will add reproducible, pre-computed blue noise to the image.
	KvImageConvert_DitherOrderedReproducible KvImageConvert = 2
	// KvImageConvert_OrderedGaussianBlue: A constant that indicates the conversion will distribute the noise according to a Gaussian distribution.
	KvImageConvert_OrderedGaussianBlue KvImageConvert = 0
	KvImageConvert_OrderedNoiseShapeMask KvImageConvert = 0
	// KvImageConvert_OrderedUniformBlue: A constant that indicates the conversion will distribute the noise uniformly.
	KvImageConvert_OrderedUniformBlue KvImageConvert = 268435456
)


func (e KvImageConvert) String() string {
	switch e {
	case KvImageConvert_DitherAtkinson:
		return "KvImageConvert_DitherAtkinson"
	case KvImageConvert_DitherFloydSteinberg:
		return "KvImageConvert_DitherFloydSteinberg"
	case KvImageConvert_DitherNone:
		return "KvImageConvert_DitherNone"
	case KvImageConvert_DitherOrdered:
		return "KvImageConvert_DitherOrdered"
	case KvImageConvert_DitherOrderedReproducible:
		return "KvImageConvert_DitherOrderedReproducible"
	case KvImageConvert_OrderedUniformBlue:
		return "KvImageConvert_OrderedUniformBlue"
	default:
		return fmt.Sprintf("KvImageConvert(%d)", e)
	}
}




type KvImageGamma uint

const (
	// KvImageGamma_11_over_5_half_precision: A half-precision calculation using a gamma value of 11/5 or 2.2.
	KvImageGamma_11_over_5_half_precision KvImageGamma = 0
	// KvImageGamma_11_over_9_half_precision: A half-precision calculation using a gamma value of 11/9 or (11/5)/(9/5).
	KvImageGamma_11_over_9_half_precision KvImageGamma = 0
	// KvImageGamma_5_over_11_half_precision: A half-precision calculation using a gamma value of 5/11 or 1/2.2.
	KvImageGamma_5_over_11_half_precision KvImageGamma = 0
	// KvImageGamma_5_over_9_half_precision: A half-precision calculation using a gamma value of 5/9 or 1/1.8.
	KvImageGamma_5_over_9_half_precision KvImageGamma = 0
	// KvImageGamma_9_over_11_half_precision: A half-precision calculation using a gamma value of 9/11 or (9/5)/(11/5).
	KvImageGamma_9_over_11_half_precision KvImageGamma = 0
	// KvImageGamma_9_over_5_half_precision: A half-precision calculation using a gamma value of 9/5 or 1.8.
	KvImageGamma_9_over_5_half_precision KvImageGamma = 0
	// KvImageGamma_BT709_forward_half_precision: The ITU-R BT.709 standard.
	KvImageGamma_BT709_forward_half_precision KvImageGamma = 0
	// KvImageGamma_BT709_reverse_half_precision: The ITU-R BT.709 standard reverse.
	KvImageGamma_BT709_reverse_half_precision KvImageGamma = 0
	// KvImageGamma_UseGammaValue: A user-defined gamma value with full-precision calculation.
	KvImageGamma_UseGammaValue KvImageGamma = 0
	// KvImageGamma_UseGammaValue_half_precision: A user-defined gamma value with half-precision calculation.
	KvImageGamma_UseGammaValue_half_precision KvImageGamma = 0
	// KvImageGamma_sRGB_forward_half_precision: A half-precision calculation using the sRGB standard gamma value of 2.2.
	KvImageGamma_sRGB_forward_half_precision KvImageGamma = 0
	// KvImageGamma_sRGB_reverse_half_precision: A half-precision calculation using the sRGB standard gamma value of 1/2.2.
	KvImageGamma_sRGB_reverse_half_precision KvImageGamma = 0
)


func (e KvImageGamma) String() string {
	switch e {
	case KvImageGamma_11_over_5_half_precision:
		return "KvImageGamma_11_over_5_half_precision"
	default:
		return fmt.Sprintf("KvImageGamma(%d)", e)
	}
}




type KvImageInterpolation uint

const (
	// KvImageInterpolationLinear: Linear interpoation
	KvImageInterpolationLinear KvImageInterpolation = 0
	// KvImageInterpolationNearest: Nearest neigborhood
	KvImageInterpolationNearest KvImageInterpolation = 0
)


func (e KvImageInterpolation) String() string {
	switch e {
	case KvImageInterpolationLinear:
		return "KvImageInterpolationLinear"
	default:
		return fmt.Sprintf("KvImageInterpolation(%d)", e)
	}
}




type KvImageMatrixType uint

const (
	KvImageMatrixType_ARGBToYpCbCrMatrix KvImageMatrixType = 0
	KvImageMatrixType_None KvImageMatrixType = 0
)


func (e KvImageMatrixType) String() string {
	switch e {
	case KvImageMatrixType_ARGBToYpCbCrMatrix:
		return "KvImageMatrixType_ARGBToYpCbCrMatrix"
	default:
		return fmt.Sprintf("KvImageMatrixType(%d)", e)
	}
}




type KvimagePNGFilterValue uint

const (
	// KvImage_PNG_FILTER_VALUE_AVG: A filter that predicts a pixel value from the average of the pixels to the left and above the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_AVG KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_NONE: No filtering.
	KvImage_PNG_FILTER_VALUE_NONE KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_PAETH: A filter that predicts a pixel value by applying a linear function to the pixels located to the left, above, and to the upper-left of the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_PAETH KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_SUB: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located to the left.
	KvImage_PNG_FILTER_VALUE_SUB KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_UP: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located above.
	KvImage_PNG_FILTER_VALUE_UP KvimagePNGFilterValue = 0
)


func (e KvimagePNGFilterValue) String() string {
	switch e {
	case KvImage_PNG_FILTER_VALUE_AVG:
		return "KvImage_PNG_FILTER_VALUE_AVG"
	default:
		return fmt.Sprintf("KvimagePNGFilterValue(%d)", e)
	}
}




type VdspHa uint

const (
	// VDSP_HALF_WINDOW: Specifies that the window should only contain the bottom half of the values (`0` to `(N+1)/2`).
	VDSP_HALF_WINDOW VdspHa = 1
	// VDSP_HANN_DENORM: Specifies a denormalized Hann window.
	VDSP_HANN_DENORM VdspHa = 0
	// VDSP_HANN_NORM: Specifies a normalized Hann window
	VDSP_HANN_NORM VdspHa = 2
)


func (e VdspHa) String() string {
	switch e {
	case VDSP_HALF_WINDOW:
		return "VDSP_HALF_WINDOW"
	case VDSP_HANN_DENORM:
		return "VDSP_HANN_DENORM"
	case VDSP_HANN_NORM:
		return "VDSP_HANN_NORM"
	default:
		return fmt.Sprintf("VdspHa(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type
type VDSP_DCT_Type int

const (
	// VDSP_DCT_II: A constant that specifies a type II discrete cosine transform.
	VDSP_DCT_II VDSP_DCT_Type = 2
	// VDSP_DCT_III: A constant that specifies a type III discrete cosine transform.
	VDSP_DCT_III VDSP_DCT_Type = 3
	// VDSP_DCT_IV: A constant that specifies a type IV discrete cosine transform.
	VDSP_DCT_IV VDSP_DCT_Type = 4
)


func (e VDSP_DCT_Type) String() string {
	switch e {
	case VDSP_DCT_II:
		return "VDSP_DCT_II"
	case VDSP_DCT_III:
		return "VDSP_DCT_III"
	case VDSP_DCT_IV:
		return "VDSP_DCT_IV"
	default:
		return fmt.Sprintf("VDSP_DCT_Type(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Direction
type VDSP_DFT_Direction int

const (
	// VDSP_DFT_FORWARD: A constant that specifies a forward transform.
	VDSP_DFT_FORWARD VDSP_DFT_Direction = 0
	// VDSP_DFT_INVERSE: A constant that specifies an inverse transform.
	VDSP_DFT_INVERSE VDSP_DFT_Direction = -1
)


func (e VDSP_DFT_Direction) String() string {
	switch e {
	case VDSP_DFT_FORWARD:
		return "VDSP_DFT_FORWARD"
	case VDSP_DFT_INVERSE:
		return "VDSP_DFT_INVERSE"
	default:
		return fmt.Sprintf("VDSP_DFT_Direction(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_RealtoComplex
type VDSP_DFT_RealtoComplex int

const (
	VDSP_DFT_Interleaved_ComplextoComplex VDSP_DFT_RealtoComplex = 0
	VDSP_DFT_Interleaved_RealtoComplex VDSP_DFT_RealtoComplex = 1
)


func (e VDSP_DFT_RealtoComplex) String() string {
	switch e {
	case VDSP_DFT_Interleaved_ComplextoComplex:
		return "VDSP_DFT_Interleaved_ComplextoComplex"
	case VDSP_DFT_Interleaved_RealtoComplex:
		return "VDSP_DFT_Interleaved_RealtoComplex"
	default:
		return fmt.Sprintf("VDSP_DFT_RealtoComplex(%d)", e)
	}
}





