// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var ()

var (
	// MLModelAuthorKey is key for the author of the model.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelMetadataKey/author
	MLModelAuthorKey MLModelMetadataKey
	// MLModelCreatorDefinedKey is key for the model creator’s custom metadata.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelMetadataKey/creatorDefinedKey
	MLModelCreatorDefinedKey MLModelMetadataKey
	// MLModelDescriptionKey is key for the overall description of the model.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelMetadataKey/description
	MLModelDescriptionKey MLModelMetadataKey
	// MLModelLicenseKey is key for the license of the model.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelMetadataKey/license
	MLModelLicenseKey MLModelMetadataKey
	// MLModelVersionStringKey is key for the version of the model.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelMetadataKey/versionString
	MLModelVersionStringKey MLModelMetadataKey
)

var (
	// See: https://developer.apple.com/documentation/coreml/mlmodelcollection/didchangenotification
	MLModelCollectionDidChangeNotification foundation.NSNotification
)

var (
	// MLModelErrorDomain is the domain for Core ML errors.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModelErrorDomain
	MLModelErrorDomain string
)

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
