// Code generated from Apple documentation. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

var VZErrorDomain foundation.NSErrorDomain

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "VZErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VZErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}

