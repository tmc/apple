// Code generated from Apple documentation. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// FSDirectoryCookieInitial is the constant initial value for the directory-enumeration cookie.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSDirectoryCookie/initial
	FSDirectoryCookieInitial FSDirectoryCookie
)

var (
	// FSDirectoryVerifierInitial is the constant initial value for the directory-enumeration verifier.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSDirectoryVerifier/initial
	FSDirectoryVerifierInitial FSDirectoryVerifier
)

var (
	// FSKitErrorDomain is an error domain for FSKit errors.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSKitErrorDomain
	FSKitErrorDomain foundation.NSErrorDomain
)

var (
	// FSKitVersionNumber is project version number for FSKit.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSKitVersionNumber
	FSKitVersionNumber float64
)

var (
	// FSKitVersionString is project version string for FSKit.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSKitVersionString
	FSKitVersionString uint8
)

var (
	// See: https://developer.apple.com/documentation/FSKit/FSOperationID/unspecified
	FSOperationIDUnspecified FSOperationID
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSDirectoryCookieInitial"); err == nil && ptr != 0 {
		FSDirectoryCookieInitial = *(*FSDirectoryCookie)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSDirectoryVerifierInitial"); err == nil && ptr != 0 {
		FSDirectoryVerifierInitial = *(*FSDirectoryVerifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSKitErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FSKitErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSKitVersionNumber"); err == nil && ptr != 0 {
		FSKitVersionNumber = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSKitVersionString"); err == nil && ptr != 0 {
		FSKitVersionString = *(*uint8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "FSOperationIDUnspecified"); err == nil && ptr != 0 {
		FSOperationIDUnspecified = *(*FSOperationID)(unsafe.Pointer(ptr))
	}

}
