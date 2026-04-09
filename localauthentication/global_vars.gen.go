// Code generated from Apple documentation. DO NOT EDIT.

package localauthentication

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// LAErrorDomain is the error domain that the framework uses when issuing errors.
	//
	// See: https://developer.apple.com/documentation/LocalAuthentication/LAErrorDomain
	LAErrorDomain string
)

var (
	// LATouchIDAuthenticationMaximumAllowableReuseDuration is the maximum allowable reuse duration.
	//
	// See: https://developer.apple.com/documentation/LocalAuthentication/LATouchIDAuthenticationMaximumAllowableReuseDuration
	LATouchIDAuthenticationMaximumAllowableReuseDuration float64
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "LAErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LAErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "LATouchIDAuthenticationMaximumAllowableReuseDuration"); err == nil && ptr != 0 {
		LATouchIDAuthenticationMaximumAllowableReuseDuration = *(*float64)(unsafe.Pointer(ptr))
	}

}
