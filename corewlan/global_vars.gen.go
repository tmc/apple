// Code generated from Apple documentation. DO NOT EDIT.

package corewlan

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/CoreWLAN/CWErrorDomain
	CWErrorDomain string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CWErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CWErrorDomain = objc.GoString(cstr)
			}
		}
	}

}
