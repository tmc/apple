// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

var mLFeatureValueImageOptionCropAndScale MLFeatureValueImageOption

var mLFeatureValueImageOptionCropRect MLFeatureValueImageOption

var MLModelAuthorKey MLModelMetadataKey

var MLModelCollectionDidChangeNotification foundation.NSNotification

var MLModelCreatorDefinedKey MLModelMetadataKey

var MLModelDescriptionKey MLModelMetadataKey

var MLModelErrorDomain string

var MLModelLicenseKey MLModelMetadataKey

var MLModelVersionStringKey MLModelMetadataKey

var mLMultiArrayDataTypeFloat MLMultiArrayDataType

var mLMultiArrayDataTypeFloat64 MLMultiArrayDataType

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "MLFeatureValueImageOptionCropAndScale"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mLFeatureValueImageOptionCropAndScale = MLFeatureValueImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLFeatureValueImageOptionCropRect"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mLFeatureValueImageOptionCropRect = MLFeatureValueImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelAuthorKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelAuthorKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelCollectionDidChangeNotification"); err == nil && ptr != 0 {
		MLModelCollectionDidChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelCreatorDefinedKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelCreatorDefinedKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelDescriptionKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelLicenseKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelLicenseKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelVersionStringKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelVersionStringKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLMultiArrayDataTypeFloat"); err == nil && ptr != 0 {
		mLMultiArrayDataTypeFloat = *(*MLMultiArrayDataType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLMultiArrayDataTypeFloat64"); err == nil && ptr != 0 {
		mLMultiArrayDataTypeFloat64 = *(*MLMultiArrayDataType)(unsafe.Pointer(ptr))
	}

}

type MLFeatureValueImageOptionValues struct{}

// MLFeatureValueImageOptions provides typed accessors for [MLFeatureValueImageOption] constants.
var MLFeatureValueImageOptions MLFeatureValueImageOptionValues

// CropAndScale returns The option you use to crop and scale an image when creating an image feature value.
func (MLFeatureValueImageOptionValues) CropAndScale() MLFeatureValueImageOption { return mLFeatureValueImageOptionCropAndScale }

// CropRect returns The option you use to crop an image when creating an image feature value.
func (MLFeatureValueImageOptionValues) CropRect() MLFeatureValueImageOption { return mLFeatureValueImageOptionCropRect }


type MLMultiArrayDataTypeValues struct{}

// MLMultiArrayDataTypes provides typed accessors for [MLMultiArrayDataType] constants.
var MLMultiArrayDataTypes MLMultiArrayDataTypeValues

// Float returns Designates the multiarray’s elements as floats.
func (MLMultiArrayDataTypeValues) Float() MLMultiArrayDataType { return mLMultiArrayDataTypeFloat }

// Float64 returns Designates the multiarray’s elements as 64-bit floats.
func (MLMultiArrayDataTypeValues) Float64() MLMultiArrayDataType { return mLMultiArrayDataTypeFloat64 }


