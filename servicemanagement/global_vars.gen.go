// Code generated from Apple documentation. DO NOT EDIT.

package servicemanagement

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/ServiceManagement/SMAppServiceErrorDomain
	SMAppServiceErrorDomain string
)

var (
	// KSMDomainSystemLaunchd is the system-level launch domain.
	//
	// See: https://developer.apple.com/documentation/ServiceManagement/kSMDomainSystemLaunchd
	KSMDomainSystemLaunchd string
	// KSMDomainUserLaunchd is the user-level launch domain.
	//
	// See: https://developer.apple.com/documentation/ServiceManagement/kSMDomainUserLaunchd
	KSMDomainUserLaunchd string
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SMAppServiceErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SMAppServiceErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSMDomainSystemLaunchd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSMDomainSystemLaunchd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSMDomainUserLaunchd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSMDomainUserLaunchd = objc.GoString(cstr)
			}
		}
	}

}

