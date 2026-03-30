// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of optional methods that you can use to respond to drawing failures and manage incremental loads.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageDelegate
type NSImageDelegate interface {
	objectivec.IObject
}

// NSImageDelegateObject wraps an existing Objective-C object that conforms to the NSImageDelegate protocol.
type NSImageDelegateObject struct {
	objectivec.Object
}

func (o NSImageDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSImageDelegateObjectFromID constructs a [NSImageDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSImageDelegateObjectFromID(id objc.ID) NSImageDelegateObject {
	return NSImageDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the image object is unable, for whatever reason, to
// lock focus on its image or draw in the specified rectangle.
//
// sender: The [NSImage] object that encountered the problem.
//
// rect: The rectangle that the image object was attempting to draw.
//
// # Return Value
//
// An [NSImage] to draw in place of the one in `sender`, or `nil` if the
// delegate wants to draw the image itself.
//
// # Discussion
//
// The delegate can do one of the following:
//
// - Return another [NSImage] object to draw in the sender’s place. - Draw
// the image itself and return `nil`. - Simply return `nil` to indicate that
// `sender` should give up on the attempt at drawing the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageDelegate/imageDidNotDraw(_:in:)
func (o NSImageDelegateObject) ImageDidNotDrawInRect(sender INSImage, rect corefoundation.CGRect) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("imageDidNotDraw:inRect:"), sender, rect)
	return NSImageFromID(rv)
}
