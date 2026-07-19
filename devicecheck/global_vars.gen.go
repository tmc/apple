// Code generated from Apple documentation. DO NOT EDIT.

package devicecheck

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// DCErrorDomain is the error domain for errors associated with DeviceCheck APIs.
	//
	// See: https://developer.apple.com/documentation/DeviceCheck/DCErrorDomain
	DCErrorDomain foundation.NSErrorDomain
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "DCErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DCErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}
