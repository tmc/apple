// Code generated from Apple documentation. DO NOT EDIT.

package network

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// KNWErrorDomainDNS is a DNS error encountered in resolving, browsing, or advertising.
	//
	// See: https://developer.apple.com/documentation/Network/kNWErrorDomainDNS
	KNWErrorDomainDNS string
	// KNWErrorDomainPOSIX is a POSIX error, which is used for most network protocol and routing errors.
	//
	// See: https://developer.apple.com/documentation/Network/kNWErrorDomainPOSIX
	KNWErrorDomainPOSIX string
	// KNWErrorDomainTLS is a TLS error reported by a TLS connection or listener.
	//
	// See: https://developer.apple.com/documentation/Network/kNWErrorDomainTLS
	KNWErrorDomainTLS string
	// See: https://developer.apple.com/documentation/Network/kNWErrorDomainWiFiAware
	KNWErrorDomainWiFiAware string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kNWErrorDomainDNS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KNWErrorDomainDNS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kNWErrorDomainPOSIX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KNWErrorDomainPOSIX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kNWErrorDomainTLS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KNWErrorDomainTLS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kNWErrorDomainWiFiAware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KNWErrorDomainWiFiAware = objc.GoString(cstr)
			}
		}
	}

}
