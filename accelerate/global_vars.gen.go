// Code generated from Apple documentation. DO NOT EDIT.

package accelerate

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	// KvImage_ARGBToYpCbCrMatrix_ITU_R_601_4 is rGB-to-Y’CbCr conversion matrix for ITU Recommendation BT.601-4.
	//
	// See: https://developer.apple.com/documentation/Accelerate/kvImage_ARGBToYpCbCrMatrix_ITU_R_601_4
	KvImage_ARGBToYpCbCrMatrix_ITU_R_601_4 VImage_ARGBToYpCbCrMatrix
	// KvImage_ARGBToYpCbCrMatrix_ITU_R_709_2 is rGB-to-Y’CbCr conversion matrix for ITU Recommendation BT.709-2.
	//
	// See: https://developer.apple.com/documentation/Accelerate/kvImage_ARGBToYpCbCrMatrix_ITU_R_709_2
	KvImage_ARGBToYpCbCrMatrix_ITU_R_709_2 VImage_ARGBToYpCbCrMatrix
)

var (
	// KvImage_YpCbCrToARGBMatrix_ITU_R_601_4 is y’CbCr-to-RGB conversion matrix for ITU Recommendation BT.601-4.
	//
	// See: https://developer.apple.com/documentation/Accelerate/kvImage_YpCbCrToARGBMatrix_ITU_R_601_4
	KvImage_YpCbCrToARGBMatrix_ITU_R_601_4 VImage_YpCbCrToARGBMatrix
	// KvImage_YpCbCrToARGBMatrix_ITU_R_709_2 is y’CbCr-to-RGB conversion matrix for ITU Recommendation BT.709-2.
	//
	// See: https://developer.apple.com/documentation/Accelerate/kvImage_YpCbCrToARGBMatrix_ITU_R_709_2
	KvImage_YpCbCrToARGBMatrix_ITU_R_709_2 VImage_YpCbCrToARGBMatrix
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kvImage_ARGBToYpCbCrMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		KvImage_ARGBToYpCbCrMatrix_ITU_R_601_4 = *(*VImage_ARGBToYpCbCrMatrix)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kvImage_ARGBToYpCbCrMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		KvImage_ARGBToYpCbCrMatrix_ITU_R_709_2 = *(*VImage_ARGBToYpCbCrMatrix)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kvImage_YpCbCrToARGBMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		KvImage_YpCbCrToARGBMatrix_ITU_R_601_4 = *(*VImage_YpCbCrToARGBMatrix)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kvImage_YpCbCrToARGBMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		KvImage_YpCbCrToARGBMatrix_ITU_R_709_2 = *(*VImage_YpCbCrToARGBMatrix)(unsafe.Pointer(ptr))
	}

}
