// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that objects implement to provide data for, and manage text selections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource
type NSTextSelectionDataSource interface {
	objectivec.IObject

	// Returns the starting and ending locations for the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/documentRange
	DocumentRange() INSTextRange

	// Enumerates all the insertion point caret offsets from left to right in visual order.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/enumerateCaretOffsetsInLineFragment(at:using:)
	EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location NSTextLocation, block NSTextLocationHandler)

	// Returns a new location using the location and offset you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/location(_:offsetBy:)
	LocationFromLocationWithOffset(location NSTextLocation, offset int) NSTextLocation

	// Returns the range of the line fragment that contains the point you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/lineFragmentRange(for:inContainerAt:)
	LineFragmentRangeForPointInContainerAtLocation(point corefoundation.CGPoint, location NSTextLocation) INSTextRange

	// Returns the offset between the two locations you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/offset(from:to:)
	OffsetFromLocationToLocation(from NSTextLocation, to NSTextLocation) int

	// Returns a text range that corresponds to selection granularity of the enclosing location.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/textRange(for:enclosing:)
	TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity NSTextSelectionGranularity, location NSTextLocation) INSTextRange

	// Returns the base writing direction at the location you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/baseWritingDirection(at:)
	BaseWritingDirectionAtLocation(location NSTextLocation) NSTextSelectionNavigationWritingDirection
}

// NSTextSelectionDataSourceObject wraps an existing Objective-C object that conforms to the NSTextSelectionDataSource protocol.
type NSTextSelectionDataSourceObject struct {
	objectivec.Object
}
func (o NSTextSelectionDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextSelectionDataSourceObjectFromID constructs a [NSTextSelectionDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextSelectionDataSourceObjectFromID(id objc.ID) NSTextSelectionDataSourceObject {
	return NSTextSelectionDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the starting and ending locations for the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/documentRange
func (o NSTextSelectionDataSourceObject) DocumentRange() INSTextRange {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("documentRange"))
	return NSTextRangeFromID(rv)
	}
// Enumerates all the insertion point caret offsets from left to right in
// visual order.
//
// location: The [NSTextLocation] to start from.
//
// block: The closure to invoke once for each logical caret edge in the line
// fragment, in left-to-right visual order. End the enumeration early by
// returning `false`.
//
// # Discussion
// 
// The `caretOffset` is in the coordinate system for the text container. When
// `leadingEdge` is `true`, it indicates that `caretOffset` is at the logical
// edge preceding the character. For left-to-right characters, the caret is on
// the left, and on the right for right-to-left characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/enumerateCaretOffsetsInLineFragment(at:using:)
func (o NSTextSelectionDataSourceObject) EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location NSTextLocation, block NSTextLocationHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("enumerateCaretOffsetsInLineFragmentAtLocation:usingBlock:"), location, block)
	}
// Returns a new location using the location and offset you specify.
//
// location: The starting location in the selection.
//
// offset: An offset that describes the extent of the new location.
//
// # Return Value
// 
// A new `NSTextLocation, or nil` when the inputs don’t produce any legal
// location, such as when the input is an out of bounds index.
//
// # Discussion
// 
// The offset value can be positive or negative indicating the logical
// direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/location(_:offsetBy:)
func (o NSTextSelectionDataSourceObject) LocationFromLocationWithOffset(location NSTextLocation, offset int) NSTextLocation {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("locationFromLocation:withOffset:"), location, offset)
	return NSTextLocationObjectFromID(rv)
	}
// Returns the range of the line fragment that contains the point you specify.
//
// point: The starting point that contains the line fragment, in the coordinate
// system of `location`.
//
// location: The location of the line fragment.
//
// # Return Value
// 
// An [NSTextRange] that describes the location of the line fragment, or nil
// if the range isn’t found.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/lineFragmentRange(for:inContainerAt:)
func (o NSTextSelectionDataSourceObject) LineFragmentRangeForPointInContainerAtLocation(point corefoundation.CGPoint, location NSTextLocation) INSTextRange {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lineFragmentRangeForPoint:inContainerAtLocation:"), point, location)
	return NSTextRangeFromID(rv)
	}
// Returns the offset between the two locations you specify.
//
// from: The starting location.
//
// to: The ending location.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/offset(from:to:)
func (o NSTextSelectionDataSourceObject) OffsetFromLocationToLocation(from NSTextLocation, to NSTextLocation) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("offsetFromLocation:toLocation:"), from, to)
	return rv
	}
// Returns a text range that corresponds to selection granularity of the
// enclosing location.
//
// selectionGranularity: One of the possible [NSTextSelection.Granularity] options.
// //
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
//
// location: A location that encloses the text range of interest.
//
// # Return Value
// 
// Returns the text range of the section, or `nil` when
// `documentRange.IsEmpty()` `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/textRange(for:enclosing:)
func (o NSTextSelectionDataSourceObject) TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity NSTextSelectionGranularity, location NSTextLocation) INSTextRange {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textRangeForSelectionGranularity:enclosingLocation:"), selectionGranularity, location)
	return NSTextRangeFromID(rv)
	}
// Returns the base writing direction at the location you specify.
//
// location: The location where you want to examine the text’s writing direction.
//
// # Return Value
// 
// The [NSWritingDirection].
//
// [NSWritingDirection]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/baseWritingDirection(at:)
func (o NSTextSelectionDataSourceObject) BaseWritingDirectionAtLocation(location NSTextLocation) NSTextSelectionNavigationWritingDirection {
	
	rv := objc.Send[NSTextSelectionNavigationWritingDirection](o.ID, objc.Sel("baseWritingDirectionAtLocation:"), location)
	return rv
	}
// Enumerates all the container boundaries starting from the location you
// specify.
//
// location: The location where the enumeration starts.
//
// reverse: A Boolean value that indicates the enumeration starts at the end of the
// container.
//
// block: A closure to invoke to evaluate the container boundaries; end the
// enumeration early by returning `false`.
//
// # Discussion
// 
// This is an optional method you implement to enumerate the text up to the
// container or page boundary when the text selection data provider supports
// this layout functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/enumerateContainerBoundaries(from:reverse:using:)
func (o NSTextSelectionDataSourceObject) EnumerateContainerBoundariesFromLocationReverseUsingBlock(location NSTextLocation, reverse bool, block NSTextLocationHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("enumerateContainerBoundariesFromLocation:reverse:usingBlock:"), location, reverse, block)
	}
// Returns the layout orientation at the location you specify.
//
// location: The location where you want to examine the text’s layout orientation.
//
// # Return Value
// 
// Returns an [NSTextSelectionNavigation.LayoutOrientation] that describes the
// orientation of the layout.
//
// [NSTextSelectionNavigation.LayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/textLayoutOrientation(at:)
func (o NSTextSelectionDataSourceObject) TextLayoutOrientationAtLocation(location NSTextLocation) NSTextSelectionNavigationLayoutOrientation {
	
	rv := objc.Send[NSTextSelectionNavigationLayoutOrientation](o.ID, objc.Sel("textLayoutOrientationAtLocation:"), location)
	return rv
	}

