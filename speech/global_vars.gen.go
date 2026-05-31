// Code generated from Apple documentation. DO NOT EDIT.

package speech

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/Speech/SFSpeechErrorDomain
	SFSpeechErrorDomain foundation.NSErrorDomain
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SFSpeechErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SFSpeechErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}
