// Code generated from Apple documentation. DO NOT EDIT.

package virtualization

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

var (
	// VZErrorDomain is the error domain for the Virtualization framework.
	//
	// See: https://developer.apple.com/documentation/Virtualization/VZErrorDomain
	VZErrorDomain foundation.NSErrorDomain
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VZErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VZErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}

