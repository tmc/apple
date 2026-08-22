// Code generated from Apple documentation. DO NOT EDIT.

package accessibility

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/Accessibility/AXAnimatedImagesEnabledDidChangeNotification
	AXAnimatedImagesEnabledDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/prefersActionSliderAlternativeDidChangeNotification
	AXPrefersActionSliderAlternativeDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/Accessibility/AXPrefersHorizontalTextLayoutDidChangeNotification
	AXPrefersHorizontalTextLayoutDidChangeNotification foundation.NSNotificationName
	// AXPrefersNonBlinkingTextInsertionIndicatorDidChangeNotification is a notification that posts when the system setting to prefer a nonblinking cursor in editable text fields changes.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/prefersNonBlinkingTextInsertionIndicatorDidChangeNotification
	AXPrefersNonBlinkingTextInsertionIndicatorDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/reduceHighlightingEffectsEnabledDidChangeNotification
	AXReduceHighlightingEffectsEnabledDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/showBordersEnabledStatusDidChangeNotification
	AXShowBordersEnabledStatusDidChangeNotification foundation.NSNotificationName
)

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXAnimatedImagesEnabledDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXAnimatedImagesEnabledDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXPrefersActionSliderAlternativeDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXPrefersActionSliderAlternativeDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXPrefersHorizontalTextLayoutDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXPrefersHorizontalTextLayoutDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXPrefersNonBlinkingTextInsertionIndicatorDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXPrefersNonBlinkingTextInsertionIndicatorDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXReduceHighlightingEffectsEnabledDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXReduceHighlightingEffectsEnabledDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXShowBordersEnabledStatusDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AXShowBordersEnabledStatusDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyAutomation"); err == nil && ptr != 0 {
		AXTechnologys.Automation = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyFullKeyboardAccess"); err == nil && ptr != 0 {
		AXTechnologys.FullKeyboardAccess = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyHoverText"); err == nil && ptr != 0 {
		AXTechnologys.HoverText = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologySpeakScreen"); err == nil && ptr != 0 {
		AXTechnologys.SpeakScreen = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologySwitchControl"); err == nil && ptr != 0 {
		AXTechnologys.SwitchControl = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyVoiceControl"); err == nil && ptr != 0 {
		AXTechnologys.VoiceControl = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyVoiceOver"); err == nil && ptr != 0 {
		AXTechnologys.VoiceOver = objc.ValueAt[AXTechnology](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AXTechnologyZoom"); err == nil && ptr != 0 {
		AXTechnologys.Zoom = objc.ValueAt[AXTechnology](ptr)
	}

}

// AXTechnologys provides typed accessors for [AXTechnology] constants.
var AXTechnologys struct {
	Automation         AXTechnology
	FullKeyboardAccess AXTechnology
	HoverText          AXTechnology
	SpeakScreen        AXTechnology
	SwitchControl      AXTechnology
	VoiceControl       AXTechnology
	VoiceOver          AXTechnology
	Zoom               AXTechnology
}
