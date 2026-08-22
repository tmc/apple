// Code generated from Apple documentation. DO NOT EDIT.

package cryptotokenkit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// TKErrorDomain is the domain for all CryptoTokenKit framework errors.
	//
	// See: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorDomain
	TKErrorDomain string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "TKErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TKErrorDomain = objc.GoString(cstr)
			}
		}
	}

}
