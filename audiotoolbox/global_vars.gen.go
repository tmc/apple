// Code generated from Apple documentation. DO NOT EDIT.

package audiotoolbox

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/AudioToolbox/kAudioComponentInstanceInvalidationNotification
	KAudioComponentInstanceInvalidationNotification string
	// See: https://developer.apple.com/documentation/AudioToolbox/kAudioComponentRegistrationsChangedNotification
	KAudioComponentRegistrationsChangedNotification string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kAudioComponentInstanceInvalidationNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KAudioComponentInstanceInvalidationNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kAudioComponentRegistrationsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KAudioComponentRegistrationsChangedNotification = objc.GoString(cstr)
			}
		}
	}

}
