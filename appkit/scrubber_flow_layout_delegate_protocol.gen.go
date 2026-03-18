// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol that a scrubber delegate can adopt to provide the size of an item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayoutDelegate
type NSScrubberFlowLayoutDelegate interface {
	objectivec.IObject
	NSScrubberDelegate
}

// NSScrubberFlowLayoutDelegateObject wraps an existing Objective-C object that conforms to the NSScrubberFlowLayoutDelegate protocol.
type NSScrubberFlowLayoutDelegateObject struct {
	objectivec.Object
}
func (o NSScrubberFlowLayoutDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSScrubberFlowLayoutDelegateObjectFromID constructs a [NSScrubberFlowLayoutDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSScrubberFlowLayoutDelegateObjectFromID(id objc.ID) NSScrubberFlowLayoutDelegateObject {
	return NSScrubberFlowLayoutDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate for the size of each item in a scrubber whose items are
// arranged in a flow layout.
//
// scrubber: The scrubber object displaying the items arranged in a flow layout.
//
// layout: The layout object requesting the information.
//
// itemIndex: The index of the item in the scrubber.
//
// # Return Value
// 
// The width and height of the item at the specified index in the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayoutDelegate/scrubber(_:layout:sizeForItemAt:)

func (o NSScrubberFlowLayoutDelegateObject) ScrubberLayoutSizeForItemAtIndex(scrubber INSScrubber, layout INSScrubberFlowLayout, itemIndex int) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("scrubber:layout:sizeForItemAtIndex:"), scrubber, layout, itemIndex)
	return rv
	}

// Tells the delegate that the item at the specified index was selected.
//
// scrubber: The scrubber object that is notifying you of the selection change.
//
// selectedIndex: The index of the item that was selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didSelectItemAt:)

func (o NSScrubberFlowLayoutDelegateObject) ScrubberDidSelectItemAtIndex(scrubber INSScrubber, selectedIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didSelectItemAtIndex:"), scrubber, selectedIndex)
	}

// Tells the delegate that the item at the specified index was highlighted.
//
// scrubber: The scrubber object that is notifying you of the highlight change.
//
// highlightedIndex: The index of the item that is now highlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didHighlightItemAt:)

func (o NSScrubberFlowLayoutDelegateObject) ScrubberDidHighlightItemAtIndex(scrubber INSScrubber, highlightedIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didHighlightItemAtIndex:"), scrubber, highlightedIndex)
	}

// Tells the delegate that the range of items currently visible in the
// scrubber has changed.
//
// scrubber: The scrubber object that is notifying you of the change in the range of
// items that are currently visible.
//
// visibleRange: The range of items that are now visible in the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didChangeVisibleRange:)

func (o NSScrubberFlowLayoutDelegateObject) ScrubberDidChangeVisibleRange(scrubber INSScrubber, visibleRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didChangeVisibleRange:"), scrubber, visibleRange)
	}

// Tells the delegate that the user is panning or scrolling the scrubber.
//
// scrubber: The scrubber the user is interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didBeginInteracting(with:)

func (o NSScrubberFlowLayoutDelegateObject) DidBeginInteractingWithScrubber(scrubber INSScrubber) {
	
	objc.Send[struct{}](o.ID, objc.Sel("didBeginInteractingWithScrubber:"), scrubber)
	}

// Tells the delegate that a pan or scroll interaction with the scrubber has
// ended.
//
// scrubber: The scrubber the user was interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didFinishInteracting(with:)

func (o NSScrubberFlowLayoutDelegateObject) DidFinishInteractingWithScrubber(scrubber INSScrubber) {
	
	objc.Send[struct{}](o.ID, objc.Sel("didFinishInteractingWithScrubber:"), scrubber)
	}

// Tells the delegate that a user interaction with the scrubber has been
// canceled.
//
// scrubber: The scrubber the user is interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didCancelInteracting(with:)

func (o NSScrubberFlowLayoutDelegateObject) DidCancelInteractingWithScrubber(scrubber INSScrubber) {
	
	objc.Send[struct{}](o.ID, objc.Sel("didCancelInteractingWithScrubber:"), scrubber)
	}

