// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface for providing data for a braille map.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMapRenderer
type AXBrailleMapRenderer interface {
	objectivec.IObject

	// A region of the UI that the system converts into a braille map and displays on the braille display.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMapRenderer/accessibilityBrailleMapRenderRegion
	AccessibilityBrailleMapRenderRegion() corefoundation.CGRect
	SetAccessibilityBrailleMapRenderRegion(value corefoundation.CGRect)
}

// AXBrailleMapRendererObject wraps an existing Objective-C object that conforms to the AXBrailleMapRenderer protocol.
type AXBrailleMapRendererObject struct {
	objectivec.Object
}

func (o AXBrailleMapRendererObject) BaseObject() objectivec.Object {
	return o.Object
}

// AXBrailleMapRendererObjectFromID constructs a [AXBrailleMapRendererObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AXBrailleMapRendererObjectFromID(id objc.ID) AXBrailleMapRendererObject {
	return AXBrailleMapRendererObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A region of the UI that the system converts into a braille map and displays
// on the braille display.
//
// # Discussion
//
// When the accessibility element that implements this property has focus, the
// system uses this value to update the braille display automatically.
//
// Set this value to specify a region of the accessibility element —
// relative to its bounds — to display on the braille display. VoiceOver
// snapshots the region of the accessibility element that you specify,
// converts the data to a braille map, and renders the braille map to the
// braille display.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMapRenderer/accessibilityBrailleMapRenderRegion
func (o AXBrailleMapRendererObject) AccessibilityBrailleMapRenderRegion() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityBrailleMapRenderRegion"))
	return corefoundation.CGRect(rv)
}

func (o AXBrailleMapRendererObject) SetAccessibilityBrailleMapRenderRegion(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityBrailleMapRenderRegion:"), value)
}
