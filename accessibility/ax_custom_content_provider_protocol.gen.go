// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface for customizing the accessibility content.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContentProvider
type AXCustomContentProvider interface {
	objectivec.IObject

	// An array of custom objects for creating accessible content.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXCustomContentProvider/accessibilityCustomContent
	AccessibilityCustomContent() []AXCustomContent
	SetAccessibilityCustomContent(value []AXCustomContent)

	// accessibilityCustomContentBlock protocol.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXCustomContentProvider/accessibilityCustomContentBlock
	AccessibilityCustomContentBlock() AXCustomContentReturnBlock
	SetAccessibilityCustomContentBlock(value objc.ID)
}

// AXCustomContentProviderObject wraps an existing Objective-C object that conforms to the AXCustomContentProvider protocol.
type AXCustomContentProviderObject struct {
	objectivec.Object
}

func (o AXCustomContentProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// AXCustomContentProviderObjectFromID constructs a [AXCustomContentProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AXCustomContentProviderObjectFromID(id objc.ID) AXCustomContentProviderObject {
	return AXCustomContentProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An array of custom objects for creating accessible content.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContentProvider/accessibilityCustomContent
func (o AXCustomContentProviderObject) AccessibilityCustomContent() []AXCustomContent {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomContent"))
	result := make([]AXCustomContent, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = AXCustomContentFromID(id)
	}
	return result
}

func (o AXCustomContentProviderObject) SetAccessibilityCustomContent(value []AXCustomContent) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomContent:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/Accessibility/AXCustomContentProvider/accessibilityCustomContentBlock
func (o AXCustomContentProviderObject) AccessibilityCustomContentBlock() AXCustomContentReturnBlock {
	rv := objc.Send[AXCustomContentReturnBlock](o.ID, objc.Sel("accessibilityCustomContentBlock"))
	return AXCustomContentReturnBlock(rv)
}

func (o AXCustomContentProviderObject) SetAccessibilityCustomContentBlock(value objc.ID) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomContentBlock:"), value)
}
