// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Optional methods that delegates implement to respond to layout changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManagerDelegate
type NSTextLayoutManagerDelegate interface {
	objectivec.IObject
}



// NSTextLayoutManagerDelegateObject wraps an existing Objective-C object that conforms to the NSTextLayoutManagerDelegate protocol.
type NSTextLayoutManagerDelegateObject struct {
	objectivec.Object
}
func (o NSTextLayoutManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSTextLayoutManagerDelegateObjectFromID constructs a [NSTextLayoutManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextLayoutManagerDelegateObjectFromID(id objc.ID) NSTextLayoutManagerDelegateObject {
	return NSTextLayoutManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// The method the framework calls to return a dictionary of attributes for
// rendering a link attribute name.
//
// textLayoutManager: The [NSTextLayoutManager].
//
// link: The link.
//
// location: The [NSTextLocation] of the link.
//
// renderingAttributes: A dictionary of attributes whose keys are [NSAttributedString.Key] values.
// //
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
//
// # Return Value
// 
// A dictionary of attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManagerDelegate/textLayoutManager(_:renderingAttributesForLink:at:defaultAttributes:)

func (o NSTextLayoutManagerDelegateObject) TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager INSTextLayoutManager, link objectivec.IObject, location NSTextLocation, renderingAttributes foundation.INSDictionary) foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textLayoutManager:renderingAttributesForLink:atLocation:defaultAttributes:"), textLayoutManager, link, location, renderingAttributes)
	return foundation.NSDictionaryFromID(rv)
	}

// The method the framework calls to determine the soft line break point.
//
// textLayoutManager: The text layout manager.
//
// location: The location of the proposed line break.
//
// hyphenating: A Boolean value that indicates the current hyphenation mode.
//
// # Return Value
// 
// A Boolean value that indicates if the framework should break the line at
// the current location.
//
// # Discussion
// 
// When `hyphenating` is `false`, [NSTextLayoutManager] tries to find the next
// line break opportunity before location. When hyphenating is `true`, it’s
// an auto-hyphenation point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManagerDelegate/textLayoutManager(_:shouldBreakLineBefore:hyphenating:)

func (o NSTextLayoutManagerDelegateObject) TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager INSTextLayoutManager, location NSTextLocation, hyphenating bool) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textLayoutManager:shouldBreakLineBeforeLocation:hyphenating:"), textLayoutManager, location, hyphenating)
	return rv
	}

// The method the framework calls to give the delegate an opportunity to
// return a custom text layout fragment.
//
// textLayoutManager: The text layout manager.
//
// location: The [NSTextLocation] of the link in the text element.
//
// textElement: The [NSTextElement] that the method could return a custom
// [NSTextLayoutFragment] from.
//
// # Return Value
// 
// An [NSTextLayoutFragment].
//
// # Discussion
// 
// Use this to provide an [NSTextLayoutFragment] specialized for an
// [NSTextElement] subclass targeted for the rendering surface.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManagerDelegate/textLayoutManager(_:textLayoutFragmentFor:in:)

func (o NSTextLayoutManagerDelegateObject) TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager INSTextLayoutManager, location NSTextLocation, textElement INSTextElement) INSTextLayoutFragment {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textLayoutManager:textLayoutFragmentForLocation:inTextElement:"), textLayoutManager, location, textElement)
	return NSTextLayoutFragmentFromID(rv)
	}







