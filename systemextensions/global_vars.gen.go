// Code generated from Apple documentation. DO NOT EDIT.

package systemextensions

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// SystemExtensionUsageDescriptionKey is a message that tells the user why the app is trying to install a system extension bundle.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/NSSystemExtensionUsageDescriptionKey
	SystemExtensionUsageDescriptionKey string
	// OSBundleUsageDescriptionKey is a message that tells the user why the app is trying to install a driver extension bundle.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSBundleUsageDescriptionKey
	OSBundleUsageDescriptionKey string
	// See: https://developer.apple.com/documentation/SystemExtensions/OSRelatedKernelExtensionKey
	OSRelatedKernelExtensionKey string
)

var (
	// OSSystemExtensionErrorDomain is the error domain identifying system extension errors.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionErrorDomain
	OSSystemExtensionErrorDomain foundation.NSErrorDomain
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSystemExtensionUsageDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SystemExtensionUsageDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "OSBundleUsageDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OSBundleUsageDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "OSRelatedKernelExtensionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OSRelatedKernelExtensionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "OSSystemExtensionErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OSSystemExtensionErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}
