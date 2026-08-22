// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The basic interface for a data axis in a chart.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor
type AXDataAxisDescriptor interface {
	objectivec.IObject
	foundation.NSCopying

	// The title of the axis.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
	Title() string
	SetTitle(value string)

	// An attributed version of the axis title.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)
}

// AXDataAxisDescriptorObject wraps an existing Objective-C object that conforms to the AXDataAxisDescriptor protocol.
type AXDataAxisDescriptorObject struct {
	foundation.NSCopyingObject
}

func (o AXDataAxisDescriptorObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// AXDataAxisDescriptorObjectFromID constructs a [AXDataAxisDescriptorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AXDataAxisDescriptorObjectFromID(id objc.ID) AXDataAxisDescriptorObject {
	return AXDataAxisDescriptorObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// The title of the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
func (o AXDataAxisDescriptorObject) Title() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

func (o AXDataAxisDescriptorObject) SetTitle(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setTitle:"), objc.String(value))
}

// An attributed version of the axis title.
//
// # Discussion
//
// If you set the value of this property, the system uses this value instead
// of [Title].
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
func (o AXDataAxisDescriptorObject) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(rv)
}

func (o AXDataAxisDescriptorObject) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAttributedTitle:"), value)
}
