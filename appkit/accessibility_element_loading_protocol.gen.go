// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to support loading.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementLoading
type NSAccessibilityElementLoading interface {
	objectivec.IObject

	// Loads the target accessibility element with the specified load token.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementLoading/accessibilityElement(withToken:)
	AccessibilityElementWithToken(token NSAccessibilityLoadingToken) NSAccessibilityElement
}

// NSAccessibilityElementLoadingObject wraps an existing Objective-C object that conforms to the NSAccessibilityElementLoading protocol.
type NSAccessibilityElementLoadingObject struct {
	objectivec.Object
}
func (o NSAccessibilityElementLoadingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityElementLoadingObjectFromID constructs a [NSAccessibilityElementLoadingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityElementLoadingObjectFromID(id objc.ID) NSAccessibilityElementLoadingObject {
	return NSAccessibilityElementLoadingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Loads the target accessibility element with the specified load token.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementLoading/accessibilityElement(withToken:)
func (o NSAccessibilityElementLoadingObject) AccessibilityElementWithToken(token NSAccessibilityLoadingToken) NSAccessibilityElement {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityElementWithToken:"), token)
	return NSAccessibilityElementFromID(rv)
	}
// Returns the range that specifies the area of interest in text-based
// accessibility elements with the specified load token.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementLoading/accessibilityRangeInTargetElement(withToken:)
func (o NSAccessibilityElementLoadingObject) AccessibilityRangeInTargetElementWithToken(token NSAccessibilityLoadingToken) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeInTargetElementWithToken:"), token)
	return rv
	}

