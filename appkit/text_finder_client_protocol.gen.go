// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods implemented by objects that support searching using the [NSTextFinder](<doc://com.apple.appkit/documentation/AppKit/NSTextFinder>) class and the in-window text find bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient
type NSTextFinderClient interface {
	objectivec.IObject

	// Allows the client to specify a single string for searching.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/string
	String() string

	// Returns whether the text is selectable.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/isSelectable
	IsSelectable() bool

	// Returns whether multiple items can be selected.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/allowsMultipleSelection
	AllowsMultipleSelection() bool

	// Returns the currently selected range.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/firstSelectedRange
	FirstSelectedRange() foundation.NSRange

	// Returns an array of selected ranges.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/selectedRanges
	SelectedRanges() []foundation.NSValue

	// Returns whether the text is editable.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/isEditable
	IsEditable() bool

	// An array of visible character ranges.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/visibleCharacterRanges
	VisibleCharacterRanges() []foundation.NSValue

	// Returns an array of selected ranges.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/selectedRanges
	SetSelectedRanges(value []foundation.NSValue)
}

// NSTextFinderClientObject wraps an existing Objective-C object that conforms to the NSTextFinderClient protocol.
type NSTextFinderClientObject struct {
	objectivec.Object
}
func (o NSTextFinderClientObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextFinderClientObjectFromID constructs a [NSTextFinderClientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextFinderClientObjectFromID(id objc.ID) NSTextFinderClientObject {
	return NSTextFinderClientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Allows the client to specify a single string for searching.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/string
func (o NSTextFinderClientObject) String() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns whether the text is selectable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/isSelectable
func (o NSTextFinderClientObject) IsSelectable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSelectable"))
	return rv
	}
// Returns whether multiple items can be selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/allowsMultipleSelection
func (o NSTextFinderClientObject) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("allowsMultipleSelection"))
	return rv
	}
// Returns the currently selected range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/firstSelectedRange
func (o NSTextFinderClientObject) FirstSelectedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("firstSelectedRange"))
	return rv
	}
// Returns an array of selected ranges.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/selectedRanges
func (o NSTextFinderClientObject) SelectedRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("selectedRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}
// Returns whether the text is editable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/isEditable
func (o NSTextFinderClientObject) IsEditable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isEditable"))
	return rv
	}
// An array of visible character ranges.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/visibleCharacterRanges
func (o NSTextFinderClientObject) VisibleCharacterRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("visibleCharacterRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}
// Returns the found string that is created by conceptually mapping its
// content to a single string, which is composed of a concatenation of all its
// substrings.
//
// characterIndex: The given character index the client should return.
//
// outRange: Returns, by reference, the “effective range” of that substring in the
// full conceptually concatenated string
//
// outFlag: Returns, by-reference, whether the substring ends with a “search
// boundary”, meaning that NSTextFinder should not attempt to find any
// matches that overlap this boundary.
//
// # Return Value
// 
// Returns the found string.
//
// # Discussion
// 
// See [NSTextFinder] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/string(at:effectiveRange:endsWithSearchBoundary:)
func (o NSTextFinderClientObject) StringAtIndexEffectiveRangeEndsWithSearchBoundary(characterIndex uint, outRange foundation.NSRange, outFlag unsafe.Pointer) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("stringAtIndex:effectiveRange:endsWithSearchBoundary:"), characterIndex, outRange, outFlag)
	return foundation.NSStringFromID(rv).String()
	}
// Returns the full length of the conceptually concatenated string return by
// the `` method.
//
// # Return Value
// 
// Returns the full length of the conceptually concatenated string in the
// second model, that is, the sum of the lengths of all of its substrings.
//
// # Discussion
// 
// See [NSTextFinder] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/stringLength()
func (o NSTextFinderClientObject) StringLength() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("stringLength"))
	return rv
	}
// Returns whether the specified strings should be replaced.
//
// ranges: The ranges of the strings to replace.
//
// strings: The replacement strings.
//
// # Return Value
// 
// Returns [true] if the replacement should occur; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// See [NSTextFinder] for a complete description.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/shouldReplaceCharacters(inRanges:with:)
func (o NSTextFinderClientObject) ShouldReplaceCharactersInRangesWithStrings(ranges []foundation.NSValue, strings []string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldReplaceCharactersInRanges:withStrings:"), objectivec.IObjectSliceToNSArray(ranges), objectivec.StringSliceToNSArray(strings))
	return rv
	}
// Replaces the text in the specified range with the new string.
//
// range: The specified range of the text to replace.
//
// string: The replacement string.
//
// # Discussion
// 
// See [NSTextFinder] for a complete description.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/replaceCharacters(in:with:)
func (o NSTextFinderClientObject) ReplaceCharactersInRangeWithString(range_ foundation.NSRange, string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(string_))
	}
// Specifies whether text characters were replaced.
//
// # Discussion
// 
// See [NSTextFinder] for a complete description.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/didReplaceCharacters()
func (o NSTextFinderClientObject) DidReplaceCharacters() {
	objc.Send[struct{}](o.ID, objc.Sel("didReplaceCharacters"))
	}
// Returns the view the context is displayed in.
//
// index: The index of the view containing the located text.
//
// outRange: Returns, by reference, the entire range of the string displayed by the view
//
// # Return Value
// 
// Returns the view the contains the found text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/contentView(at:effectiveCharacterRange:)
func (o NSTextFinderClientObject) ContentViewAtIndexEffectiveCharacterRange(index uint, outRange foundation.NSRange) INSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentViewAtIndex:effectiveCharacterRange:"), index, outRange)
	return NSViewFromID(rv)
	}
// An array containing the located text in the content view’s coordinate
// system.
//
// range: The range of the located character string.
//
// # Return Value
// 
// An array containing the rectangles containing the located text in the
// content view object’s coordinate system and return that array. The
// rectangles are return wrapped as [NSValue] objects.
//
// # Discussion
// 
// The text finder uses this method to determine the location to display the
// find indicator.
// 
// The given range is guaranteed not to overlap multiple views.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/rects(forCharacterRange:)
func (o NSTextFinderClientObject) RectsForCharacterRange(range_ foundation.NSRange) []foundation.NSValue {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("rectsForCharacterRange:"), range_)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}
// Scrolls the specified range such that it is visible.
//
// range: The range to display.
//
// # Discussion
// 
// This method is used by all actions, but is not strictly required by any.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/scrollRangeToVisible(_:)
func (o NSTextFinderClientObject) ScrollRangeToVisible(range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollRangeToVisible:"), range_)
	}
// Draw the glyphs for the requested character range as they are drawn in the
// given content view.
//
// range: The character range.
//
// view: The content view.
//
// # Discussion
// 
// If the character range partially intersects a glyph range, then the full
// glyph is drawn to avoid additional layout.
// 
// The given range is guaranteed to be completely contained by the given view.
// When this method is called, a drawing context effectively identical to the
// one provided to the view’s [DrawRect] method is configured. This method
// is mainly used to draw find indicator contents, so implementations should
// check -the view property [isDrawingFindIndicator] to ensure that the text
// will be easily readable against the background of the find indicator when
// it returns [true]. If this method is not implemented, then the find
// indicator will be drawn using the content view’s [DrawRect] method
// instead.
//
// [isDrawingFindIndicator]: https://developer.apple.com/documentation/AppKit/NSView/isDrawingFindIndicator
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderClient/drawCharacters(in:forContentView:)
func (o NSTextFinderClientObject) DrawCharactersInRangeForContentView(range_ foundation.NSRange, view INSView) {
	objc.Send[struct{}](o.ID, objc.Sel("drawCharactersInRange:forContentView:"), range_, view)
	}

func (o NSTextFinderClientObject) SetSelectedRanges(value []foundation.NSValue) {
	objc.Send[struct{}](o.ID, objc.Sel("setSelectedRanges:"), objectivec.IObjectSliceToNSArray(value))
}

