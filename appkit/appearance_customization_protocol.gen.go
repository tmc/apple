// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods for getting and setting the appearance attributes of a view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization
type NSAppearanceCustomization interface {
	objectivec.IObject

	// The appearance of the receiver, in an [NSAppearance] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
	Appearance() INSAppearance

	// The appearance that will be used when the receiver is drawn onscreen, in an [NSAppearance] object. (read-only)
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
	EffectiveAppearance() INSAppearance

	// The appearance of the receiver, in an [NSAppearance] object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
	SetAppearance(value INSAppearance)
}

// NSAppearanceCustomizationObject wraps an existing Objective-C object that conforms to the NSAppearanceCustomization protocol.
type NSAppearanceCustomizationObject struct {
	objectivec.Object
}
func (o NSAppearanceCustomizationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAppearanceCustomizationObjectFromID constructs a [NSAppearanceCustomizationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAppearanceCustomizationObjectFromID(id objc.ID) NSAppearanceCustomizationObject {
	return NSAppearanceCustomizationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The appearance of the receiver, in an [NSAppearance] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (o NSAppearanceCustomizationObject) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(rv)
	}
// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (o NSAppearanceCustomizationObject) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(rv)
	}

func (o NSAppearanceCustomizationObject) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](o.ID, objc.Sel("setAppearance:"), value)
}

