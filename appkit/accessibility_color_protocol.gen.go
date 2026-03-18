// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a color.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityColor
type NSAccessibilityColor interface {
	objectivec.IObject

	// Returns a localized description of the color for use in accessibility attributes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityColor/accessibilityName
	AccessibilityName() string
}

// NSAccessibilityColorObject wraps an existing Objective-C object that conforms to the NSAccessibilityColor protocol.
type NSAccessibilityColorObject struct {
	objectivec.Object
}
func (o NSAccessibilityColorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityColorObjectFromID constructs a [NSAccessibilityColorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityColorObjectFromID(id objc.ID) NSAccessibilityColorObject {
	return NSAccessibilityColorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a localized description of the color for use in accessibility
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityColor/accessibilityName

func (o NSAccessibilityColorObject) AccessibilityName() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityName"))
	return foundation.NSStringFromID(rv).String()
	}

