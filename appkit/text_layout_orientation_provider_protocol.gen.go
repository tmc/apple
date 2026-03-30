// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that define the orientation of text for an object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutOrientationProvider
type NSTextLayoutOrientationProvider interface {
	objectivec.IObject

	// The default layout orientation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutOrientationProvider/layoutOrientation
	LayoutOrientation() NSTextLayoutOrientation
}

// NSTextLayoutOrientationProviderObject wraps an existing Objective-C object that conforms to the NSTextLayoutOrientationProvider protocol.
type NSTextLayoutOrientationProviderObject struct {
	objectivec.Object
}

func (o NSTextLayoutOrientationProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextLayoutOrientationProviderObjectFromID constructs a [NSTextLayoutOrientationProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextLayoutOrientationProviderObjectFromID(id objc.ID) NSTextLayoutOrientationProviderObject {
	return NSTextLayoutOrientationProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The default layout orientation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutOrientationProvider/layoutOrientation
func (o NSTextLayoutOrientationProviderObject) LayoutOrientation() NSTextLayoutOrientation {
	rv := objc.Send[NSTextLayoutOrientation](o.ID, objc.Sel("layoutOrientation"))
	return rv
}
