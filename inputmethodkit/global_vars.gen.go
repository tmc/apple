// Code generated from Apple documentation. DO NOT EDIT.

package inputmethodkit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidatesOpacityAttributeName
	IMKCandidatesOpacityAttributeName string
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidatesSendServerKeyEventFirst
	IMKCandidatesSendServerKeyEventFirst string
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKControllerClass
	IMKControllerClass string
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKDelegateClass
	IMKDelegateClass string
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKModeDictionary
	IMKModeDictionary string
	// See: https://developer.apple.com/documentation/InputMethodKit/kIMKCommandClientName
	KIMKCommandClientName string
	// See: https://developer.apple.com/documentation/InputMethodKit/kIMKCommandMenuItemName
	KIMKCommandMenuItemName string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IMKCandidatesOpacityAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IMKCandidatesOpacityAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IMKCandidatesSendServerKeyEventFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IMKCandidatesSendServerKeyEventFirst = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IMKControllerClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IMKControllerClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IMKDelegateClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IMKDelegateClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IMKModeDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IMKModeDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIMKCommandClientName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIMKCommandClientName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIMKCommandMenuItemName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIMKCommandMenuItemName = objc.GoString(cstr)
			}
		}
	}

}
