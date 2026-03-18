// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol the text content manager and its concrete subclasses conform to, which defines the interface for interacting with custom content types of a text document.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider
type NSTextElementProvider interface {
	objectivec.IObject

	// Describes the starting and ending locations for the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/documentRange
	DocumentRange() INSTextRange

	// Enumerates text elements starting at the text location you provide.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/enumerateTextElements(from:options:using:)
	EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation NSTextLocation, options NSTextContentManagerEnumerationOptions, block TextElementHandler) NSTextLocation

	// Replaces the characters specified by range with the text elements you provide.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/replaceContents(in:with:)
	ReplaceContentsInRangeWithTextElements(range_ INSTextRange, textElements []NSTextElement)

	// Synchronizes changes to the backing store.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/synchronizeToBackingStore(_:)
	SynchronizeToBackingStore(completionHandler ErrorHandler)
}

// NSTextElementProviderObject wraps an existing Objective-C object that conforms to the NSTextElementProvider protocol.
type NSTextElementProviderObject struct {
	objectivec.Object
}
func (o NSTextElementProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextElementProviderObjectFromID constructs a [NSTextElementProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextElementProviderObjectFromID(id objc.ID) NSTextElementProviderObject {
	return NSTextElementProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Describes the starting and ending locations for the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/documentRange

func (o NSTextElementProviderObject) DocumentRange() INSTextRange {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("documentRange"))
	return NSTextRangeFromID(rv)
	}

// Enumerates text elements starting at the text location you provide.
//
// textLocation: The [NSTextLocation] at which to start the enumeration.
//
// options: One of the possible [NSTextElementProviderEnumerationOptions] directions.
//
// block: A block you use to evaluate whether to continue the enumeration or tell the
// method to stop. Return `false` to end the enumeration process.
//
// # Return Value
// 
// An [NSTextLocation].
//
// # Discussion
// 
// If `textLocation` is `nil`, the method uses `documentRange.Location()` for
// forward enumeration and `documentRange.EndLocation()` for reverse
// enumeration. When enumerating backward, the method starts with the element
// preceding the one containing `textLocation`. If enumerated at least one
// element, it returns the edge of the enumerated range.
// 
// The enumerated range might not match the range of the last element
// returned. It enumerates the elements in the sequence, but it can skip a
// range (it can limit the maximum number of text elements enumerated for a
// single invocation or hide some elements from the layout).
// 
// Returning [NO] or `false` from block breaks out of the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/enumerateTextElements(from:options:using:)

func (o NSTextElementProviderObject) EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation NSTextLocation, options NSTextContentManagerEnumerationOptions, block TextElementHandler) NSTextLocation {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("enumerateTextElementsFromLocation:options:usingBlock:"), textLocation, options, block)
	return NSTextLocationObjectFromID(rv)
	}

// Replaces the characters specified by range with the text elements you
// provide.
//
// range: An [NSTextRange].
//
// textElements: The elements to replace that characters at `range`.
//
// # Discussion
// 
// If the edges of `range` aren’t at existing element range boundaries, the
// method either splits the element if it allows the operation (for example,
// [NSTextParagraph]), or the adjusts the replacement range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/replaceContents(in:with:)

func (o NSTextElementProviderObject) ReplaceContentsInRangeWithTextElements(range_ INSTextRange, textElements []NSTextElement) {
	
	objc.Send[struct{}](o.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, objectivec.IObjectSliceToNSArray(textElements))
	}

// Synchronizes changes to the backing store.
//
// completionHandler: A completion handler to run upon successful completion, or to process an
// error upon failure.
//
// # Discussion
// 
// If `completionHandler` is `nil`, performs the operation synchronously. The
// `completionHandler` gets passed `error` if the synchronization fails. It
// should block (or fails if synchronous) when there’s an active
// transaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/synchronizeToBackingStore(_:)

func (o NSTextElementProviderObject) SynchronizeToBackingStore(completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("synchronizeToBackingStore:"), completionHandler)
	}

// Returns a new location from location with offset you provide.
//
// location: An [NSTextLocation] in the text element.
//
// offset: An offset of the number of characters to or from `location`.
//
// # Return Value
// 
// An new [NSTextLocation], or `nil` of the offset exceeds the bounds of the
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/location(_:offsetBy:)

func (o NSTextElementProviderObject) LocationFromLocationWithOffset(location NSTextLocation, offset int) NSTextLocation {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("locationFromLocation:withOffset:"), location, offset)
	return NSTextLocationObjectFromID(rv)
	}

// A method you implement if the location backing store requires manual
// adjustment after editing.
//
// textRange: An [NSTextRange] that the method adjusts.
//
// forEditingTextSelection: A Boolean value that indicates if `textRange` is for the text selection
// associated with the edit session.
//
// # Return Value
// 
// When `textRange` is intersecting or following the current edited range, the
// method returns the range adjusted for the modification in the editing
// session. Returns `nil`, when no adjustment necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/adjustedRange(from:forEditingTextSelection:)

func (o NSTextElementProviderObject) AdjustedRangeFromRangeForEditingTextSelection(textRange INSTextRange, forEditingTextSelection bool) INSTextRange {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"), textRange, forEditingTextSelection)
	return NSTextRangeFromID(rv)
	}

// Returns the offset between the two specified locations.
//
// from: A starting location.
//
// to: An ending location.
//
// # Return Value
// 
// An [Integer] that represents the offset between the starting and ending
// locations.
//
// # Discussion
// 
// The return value could be positive or negative. This method can return
// [NSNotFound] when the method can’t represent an offset as an integer
// value. This can occur, for example, if the locations aren’t in the same
// document).
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/offset(from:to:)

func (o NSTextElementProviderObject) OffsetFromLocationToLocation(from NSTextLocation, to NSTextLocation) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("offsetFromLocation:toLocation:"), from, to)
	return rv
	}

