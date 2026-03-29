// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextTab] class.
var (
	_NSTextTabClass     NSTextTabClass
	_NSTextTabClassOnce sync.Once
)

func getNSTextTabClass() NSTextTabClass {
	_NSTextTabClassOnce.Do(func() {
		_NSTextTabClass = NSTextTabClass{class: objc.GetClass("NSTextTab")}
	})
	return _NSTextTabClass
}

// GetNSTextTabClass returns the class object for NSTextTab.
func GetNSTextTabClass() NSTextTabClass {
	return getNSTextTabClass()
}

type NSTextTabClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextTabClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextTabClass) Alloc() NSTextTab {
	rv := objc.Send[NSTextTab](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A tab in a paragraph.
//
// # Overview
// 
// A text tab represents a tab in an [NSParagraphStyle] object, storing an
// alignment type and location. [NSTextTab] objects are most frequently used
// with the TextKit system and with [NSRulerView] and [NSRulerMarker] objects.
// 
// The text system supports four alignment types: left, center, right, and
// decimal (based on the decimal separator character of the locale in effect).
// These alignment types are absolute, not based on the line sweep direction
// of text. For example, tabbed text is always positioned to the left of a
// right-aligned tab, whether the line sweep direction is left to right or
// right to left. A tab’s location, on the other hand, is relative to the
// back margin. A tab set at 1.5”, for example, is at 1.5” from the right
// in right to left text.
//
// # Creating a text tab
//
//   - [NSTextTab.InitWithTextAlignmentLocationOptions]: Initializes a text tab with the specified text alignment, location, and options.
//
// # Getting tab stop information
//
//   - [NSTextTab.Location]: The text tab’s ruler location relative to the back margin.
//
// # Getting text tab information
//
//   - [NSTextTab.Alignment]: The text alignment of the text tab.
//   - [NSTextTab.Options]: The dictionary of attributes for the text tab.
//
// # Deprecated
//
//   - [NSTextTab.InitWithTypeLocation]: Initializes a newly allocated text tab with the specified alignment and location.
//   - [NSTextTab.TabStopType]: The text tab’s type of tab stop.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab
type NSTextTab struct {
	objectivec.Object
}

// NSTextTabFromID constructs a [NSTextTab] from an objc.ID.
//
// A tab in a paragraph.
func NSTextTabFromID(id objc.ID) NSTextTab {
	return NSTextTab{objectivec.Object{ID: id}}
}
// NOTE: NSTextTab adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextTab] class.
//
// # Creating a text tab
//
//   - [INSTextTab.InitWithTextAlignmentLocationOptions]: Initializes a text tab with the specified text alignment, location, and options.
//
// # Getting tab stop information
//
//   - [INSTextTab.Location]: The text tab’s ruler location relative to the back margin.
//
// # Getting text tab information
//
//   - [INSTextTab.Alignment]: The text alignment of the text tab.
//   - [INSTextTab.Options]: The dictionary of attributes for the text tab.
//
// # Deprecated
//
//   - [INSTextTab.InitWithTypeLocation]: Initializes a newly allocated text tab with the specified alignment and location.
//   - [INSTextTab.TabStopType]: The text tab’s type of tab stop.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab
type INSTextTab interface {
	objectivec.IObject

	// Topic: Creating a text tab

	// Initializes a text tab with the specified text alignment, location, and options.
	InitWithTextAlignmentLocationOptions(alignment NSTextAlignment, loc float64, options foundation.INSDictionary) NSTextTab

	// Topic: Getting tab stop information

	// The text tab’s ruler location relative to the back margin.
	Location() float64

	// Topic: Getting text tab information

	// The text alignment of the text tab.
	Alignment() NSTextAlignment
	// The dictionary of attributes for the text tab.
	Options() foundation.INSDictionary

	// Topic: Deprecated

	// Initializes a newly allocated text tab with the specified alignment and location.
	InitWithTypeLocation(type_ NSTextTabType, loc float64) NSTextTab
	// The text tab’s type of tab stop.
	TabStopType() NSTextTabType

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextTab) Init() NSTextTab {
	rv := objc.Send[NSTextTab](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextTab) Autorelease() NSTextTab {
	rv := objc.Send[NSTextTab](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextTab creates a new NSTextTab instance.
func NewNSTextTab() NSTextTab {
	class := getNSTextTabClass()
	rv := objc.Send[NSTextTab](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a text tab with the specified text alignment, location, and
// options.
//
// alignment: The alignment of the text.
//
// loc: The position of the text tab on the ruler, relative to the back margin.
//
// options: Options to apply to the text tab.
//
// # Return Value
// 
// An initialized text tab.
//
// # Discussion
// 
// The text alignment is used to determine the position of text inside the tab
// column. See [NSParagraphStyle.TextTabType] for a mapping between alignments
// and tab stop types
//
// [NSParagraphStyle.TextTabType]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/init(textAlignment:location:options:)
func NewTextTabWithTextAlignmentLocationOptions(alignment NSTextAlignment, loc float64, options foundation.INSDictionary) NSTextTab {
	instance := getNSTextTabClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextAlignment:location:options:"), alignment, loc, options)
	return NSTextTabFromID(rv)
}

// Initializes a newly allocated text tab with the specified alignment and
// location.
//
// # Discussion
// 
// The location is relative to the back margin, based on the line sweep
// direction of the paragraph. The value in the `type` parameter can be any of
// the values described in [NSParagraphStyle.TextTabType].
//
// [NSParagraphStyle.TextTabType]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/init(type:location:)
func NewTextTabWithTypeLocation(type_ NSTextTabType, loc float64) NSTextTab {
	instance := getNSTextTabClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:location:"), type_, loc)
	return NSTextTabFromID(rv)
}

// Initializes a text tab with the specified text alignment, location, and
// options.
//
// alignment: The alignment of the text.
//
// loc: The position of the text tab on the ruler, relative to the back margin.
//
// options: Options to apply to the text tab.
//
// # Return Value
// 
// An initialized text tab.
//
// # Discussion
// 
// The text alignment is used to determine the position of text inside the tab
// column. See [NSParagraphStyle.TextTabType] for a mapping between alignments
// and tab stop types
//
// [NSParagraphStyle.TextTabType]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/init(textAlignment:location:options:)
func (t NSTextTab) InitWithTextAlignmentLocationOptions(alignment NSTextAlignment, loc float64, options foundation.INSDictionary) NSTextTab {
	rv := objc.Send[NSTextTab](t.ID, objc.Sel("initWithTextAlignment:location:options:"), alignment, loc, options)
	return rv
}
// Initializes a newly allocated text tab with the specified alignment and
// location.
//
// # Discussion
// 
// The location is relative to the back margin, based on the line sweep
// direction of the paragraph. The value in the `type` parameter can be any of
// the values described in [NSParagraphStyle.TextTabType].
//
// [NSParagraphStyle.TextTabType]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/init(type:location:)
func (t NSTextTab) InitWithTypeLocation(type_ NSTextTabType, loc float64) NSTextTab {
	rv := objc.Send[NSTextTab](t.ID, objc.Sel("initWithType:location:"), type_, loc)
	return rv
}
func (t NSTextTab) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the column terminators for the specified locale.
//
// aLocale: The locale to use when determining the terminators. Specify `nil` to use
// the system’s current locale. You can get the user’s locale using the
// [current] method of [NSLocale].
// //
// [NSLocale]: https://developer.apple.com/documentation/Foundation/NSLocale
// [current]: https://developer.apple.com/documentation/Foundation/NSLocale/current
//
// # Return Value
// 
// The characters for the column terminators.
//
// # Discussion
// 
// The returned value can be used as the value for [columnTerminators] to make
// a decimal tab stop.
//
// [columnTerminators]: https://developer.apple.com/documentation/AppKit/NSTextTab/OptionKey/columnTerminators
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/columnTerminators(for:)
func (_NSTextTabClass NSTextTabClass) ColumnTerminatorsForLocale(aLocale foundation.NSLocale) foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSTextTabClass.class), objc.Sel("columnTerminatorsForLocale:"), aLocale)
	return foundation.NSCharacterSetFromID(rv)
}

// The text tab’s ruler location relative to the back margin.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/location
func (t NSTextTab) Location() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("location"))
	return rv
}
// The text alignment of the text tab.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/alignment
func (t NSTextTab) Alignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](t.ID, objc.Sel("alignment"))
	return NSTextAlignment(rv)
}
// The dictionary of attributes for the text tab.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/options
func (t NSTextTab) Options() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("options"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The text tab’s type of tab stop.
//
// # Discussion
// 
// Possible values are listed in [NSParagraphStyle.TextTabType].
//
// [NSParagraphStyle.TextTabType]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/tabStopType
func (t NSTextTab) TabStopType() NSTextTabType {
	rv := objc.Send[NSTextTabType](t.ID, objc.Sel("tabStopType"))
	return NSTextTabType(rv)
}

