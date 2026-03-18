// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

var MLModelAuthorKey MLModelMetadataKey

var MLModelCollectionDidChangeNotification foundation.NSNotification

var MLModelCreatorDefinedKey MLModelMetadataKey

var MLModelDescriptionKey MLModelMetadataKey

var MLModelErrorDomain string

var MLModelLicenseKey MLModelMetadataKey

var MLModelVersionStringKey MLModelMetadataKey

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLFeatureValueImageOptionCropAndScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLFeatureValueImageOptions.CropAndScale = MLFeatureValueImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLFeatureValueImageOptionCropRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLFeatureValueImageOptions.CropRect = MLFeatureValueImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelAuthorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
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
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelCreatorDefinedKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelDescriptionKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelLicenseKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelLicenseKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MLModelVersionStringKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MLModelVersionStringKey = MLModelMetadataKey(objc.GoString(cstr))
			}
		}
	}

}

// MLFeatureValueImageOptions provides typed accessors for [MLFeatureValueImageOption] constants.
var MLFeatureValueImageOptions struct {
	// CropAndScale: The option you use to crop and scale an image when creating an image feature value.
	CropAndScale MLFeatureValueImageOption
	// CropRect: The option you use to crop an image when creating an image feature value.
	CropRect MLFeatureValueImageOption
}

