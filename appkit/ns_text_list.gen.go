// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextList] class.
var (
	_NSTextListClass     NSTextListClass
	_NSTextListClassOnce sync.Once
)

func getNSTextListClass() NSTextListClass {
	_NSTextListClassOnce.Do(func() {
		_NSTextListClass = NSTextListClass{class: objc.GetClass("NSTextList")}
	})
	return _NSTextListClass
}

// GetNSTextListClass returns the class object for NSTextList.
func GetNSTextListClass() NSTextListClass {
	return getNSTextListClass()
}

type NSTextListClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextListClass) Alloc() NSTextList {
	rv := objc.Send[NSTextList](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A section of text that forms a single list.
//
// # Overview
// 
// The visible elements of the list, including list markers, appear in the
// text as they do for lists created by hand. The list object, however, allows
// the list to be recognized as such by the text system. This enables
// automatic creation of markers and spacing. Text lists are used in HTML
// import and export.
// 
// Text lists appear as attributes on paragraphs, as part of the paragraph
// style. An [NSParagraphStyle] may have an array of text lists, representing
// the nested lists containing the paragraph, in order from outermost to
// innermost. For example, if list1 contains four paragraphs, the middle two
// of which are also in the inner list2, then the text lists array for the
// first and fourth paragraphs is (list1), while the text lists array for the
// second and third paragraphs is (list1, list2).
// 
// The methods implementing this are [NSTextList.TextLists] on [NSParagraphStyle], and
// [NSTextList.TextLists] on [NSMutableParagraphStyle].
// 
// In addition, [NSAttributedString] has convenience methods for lists, such
// as [range(of:at:)], which determines the range covered by a list, and
// [itemNumber(in:at:)], which determines the ordinal position within a list
// of a particular item.
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [itemNumber(in:at:)]: https://developer.apple.com/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
// [range(of:at:)]: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
//
// # Creating a text list
//
//   - [NSTextList.InitWithCoder]: Initializes and returns a newly allocated text list item.
//   - [NSTextList.InitWithMarkerFormatOptions]: Returns an initialized text list.
//   - [NSTextList.InitWithMarkerFormatOptionsStartingItemNumber]: Returns a new text list with the format, options, and starting item number you provide.
//
// # Working with markers
//
//   - [NSTextList.MarkerFormat]: Returns the marker format string used by the receiver.
//   - [NSTextList.MarkerForItemNumber]: Returns the computed value for a specific ordinal position in the list.
//
// # Getting list options
//
//   - [NSTextList.Ordered]
//   - [NSTextList.ListOptions]: Returns the list options mask value of the receiver.
//
// # Managing item numbering
//
//   - [NSTextList.StartingItemNumber]: Sets the starting item number for the text list.
//   - [NSTextList.SetStartingItemNumber]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList
type NSTextList struct {
	objectivec.Object
}

// NSTextListFromID constructs a [NSTextList] from an objc.ID.
//
// A section of text that forms a single list.
func NSTextListFromID(id objc.ID) NSTextList {
	return NSTextList{objectivec.Object{ID: id}}
}
// NOTE: NSTextList adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextList] class.
//
// # Creating a text list
//
//   - [INSTextList.InitWithCoder]: Initializes and returns a newly allocated text list item.
//   - [INSTextList.InitWithMarkerFormatOptions]: Returns an initialized text list.
//   - [INSTextList.InitWithMarkerFormatOptionsStartingItemNumber]: Returns a new text list with the format, options, and starting item number you provide.
//
// # Working with markers
//
//   - [INSTextList.MarkerFormat]: Returns the marker format string used by the receiver.
//   - [INSTextList.MarkerForItemNumber]: Returns the computed value for a specific ordinal position in the list.
//
// # Getting list options
//
//   - [INSTextList.Ordered]
//   - [INSTextList.ListOptions]: Returns the list options mask value of the receiver.
//
// # Managing item numbering
//
//   - [INSTextList.StartingItemNumber]: Sets the starting item number for the text list.
//   - [INSTextList.SetStartingItemNumber]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList
type INSTextList interface {
	objectivec.IObject

	// Topic: Creating a text list

	// Initializes and returns a newly allocated text list item.
	InitWithCoder(coder foundation.INSCoder) NSTextList
	// Returns an initialized text list.
	InitWithMarkerFormatOptions(markerFormat NSTextListMarkerFormat, options uint) NSTextList
	// Returns a new text list with the format, options, and starting item number you provide.
	InitWithMarkerFormatOptionsStartingItemNumber(markerFormat NSTextListMarkerFormat, options NSTextListOptions, startingItemNumber int) NSTextList

	// Topic: Working with markers

	// Returns the marker format string used by the receiver.
	MarkerFormat() NSTextListMarkerFormat
	// Returns the computed value for a specific ordinal position in the list.
	MarkerForItemNumber(itemNumber int) string

	// Topic: Getting list options

	Ordered() bool
	// Returns the list options mask value of the receiver.
	ListOptions() NSTextListOptions

	// Topic: Managing item numbering

	// Sets the starting item number for the text list.
	StartingItemNumber() int
	SetStartingItemNumber(value int)

	// The text lists that contain the paragraph.
	TextLists() INSTextList
	SetTextLists(value INSTextList)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextList) Init() NSTextList {
	rv := objc.Send[NSTextList](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextList) Autorelease() NSTextList {
	rv := objc.Send[NSTextList](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextList creates a new NSTextList instance.
func NewNSTextList() NSTextList {
	class := getNSTextListClass()
	rv := objc.Send[NSTextList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a newly allocated text list item.
//
// coder: An instance of [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(coder:)
func NewTextListWithCoder(coder foundation.INSCoder) NSTextList {
	instance := getNSTextListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextListFromID(rv)
}

// Returns an initialized text list.
//
// markerFormat: The marker format for the text list.
//
// options: The marker options for the text list. Values for `mask` are listed in
// Constants.
//
// # Return Value
// 
// An initialized text list.
//
// # Discussion
// 
// The marker format is specified as a constant string, except for a numbering
// specifier, which takes the form `{`keyword`}`. The currently supported
// values for keyword include:
// 
// - `box` - `check` - `circle` - `diamond` - `disc` - `hyphen` - `square` -
// `lower-hexadecimal` - `upper-hexadecimal` - `octal` - `lower-alpha` or
// `lower-latin` - `upper-alpha` or `upper-latin` - `lower-roman` -
// `upper-roman` - `decimal`
// 
// Thus, for example, `@"({decimal})"` would specify the format for a list
// numbered (1), (2), (3), and so on, and `@"{upper-roman}"` would specify the
// format for a list numbered I, II, III, IV, and so on. (All of these
// keywords are included in the Cascading Style Sheets level 3 specification.)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(markerFormat:options:)
func NewTextListWithMarkerFormatOptions(markerFormat NSTextListMarkerFormat, options uint) NSTextList {
	instance := getNSTextListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkerFormat:options:"), objc.String(string(markerFormat)), options)
	return NSTextListFromID(rv)
}

// Returns a new text list with the format, options, and starting item number
// you provide.
//
// markerFormat: One of the possible [NSTextListMarkerFormat] formats.
//
// options: One or more of the possible [NSTextList.Options] options.
// //
// [NSTextList.Options]: https://developer.apple.com/documentation/AppKit/NSTextList/Options
//
// startingItemNumber: An integer that represents the stating item number.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(markerFormat:options:startingItemNumber:)
func NewTextListWithMarkerFormatOptionsStartingItemNumber(markerFormat NSTextListMarkerFormat, options NSTextListOptions, startingItemNumber int) NSTextList {
	instance := getNSTextListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkerFormat:options:startingItemNumber:"), objc.String(string(markerFormat)), options, startingItemNumber)
	return NSTextListFromID(rv)
}

// Initializes and returns a newly allocated text list item.
//
// coder: An instance of [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(coder:)
func (t NSTextList) InitWithCoder(coder foundation.INSCoder) NSTextList {
	rv := objc.Send[NSTextList](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns an initialized text list.
//
// markerFormat: The marker format for the text list.
//
// options: The marker options for the text list. Values for `mask` are listed in
// Constants.
//
// # Return Value
// 
// An initialized text list.
//
// # Discussion
// 
// The marker format is specified as a constant string, except for a numbering
// specifier, which takes the form `{`keyword`}`. The currently supported
// values for keyword include:
// 
// - `box` - `check` - `circle` - `diamond` - `disc` - `hyphen` - `square` -
// `lower-hexadecimal` - `upper-hexadecimal` - `octal` - `lower-alpha` or
// `lower-latin` - `upper-alpha` or `upper-latin` - `lower-roman` -
// `upper-roman` - `decimal`
// 
// Thus, for example, `@"({decimal})"` would specify the format for a list
// numbered (1), (2), (3), and so on, and `@"{upper-roman}"` would specify the
// format for a list numbered I, II, III, IV, and so on. (All of these
// keywords are included in the Cascading Style Sheets level 3 specification.)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(markerFormat:options:)
func (t NSTextList) InitWithMarkerFormatOptions(markerFormat NSTextListMarkerFormat, options uint) NSTextList {
	rv := objc.Send[NSTextList](t.ID, objc.Sel("initWithMarkerFormat:options:"), objc.String(string(markerFormat)), options)
	return rv
}

// Returns a new text list with the format, options, and starting item number
// you provide.
//
// markerFormat: One of the possible [NSTextListMarkerFormat] formats.
//
// options: One or more of the possible [NSTextList.Options] options.
// //
// [NSTextList.Options]: https://developer.apple.com/documentation/AppKit/NSTextList/Options
//
// startingItemNumber: An integer that represents the stating item number.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/init(markerFormat:options:startingItemNumber:)
func (t NSTextList) InitWithMarkerFormatOptionsStartingItemNumber(markerFormat NSTextListMarkerFormat, options NSTextListOptions, startingItemNumber int) NSTextList {
	rv := objc.Send[NSTextList](t.ID, objc.Sel("initWithMarkerFormat:options:startingItemNumber:"), objc.String(string(markerFormat)), options, startingItemNumber)
	return rv
}

// Returns the computed value for a specific ordinal position in the list.
//
// itemNumber: The ordinal position in the list whose computed marker value is desired.
//
// # Return Value
// 
// The computed maker value for `itemNum`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/marker(forItemNumber:)
func (t NSTextList) MarkerForItemNumber(itemNumber int) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("markerForItemNumber:"), itemNumber)
	return foundation.NSStringFromID(rv).String()
}
func (t NSTextList) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the marker format string used by the receiver.
//
// # Return Value
// 
// The marker format string used by the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/markerFormat-swift.property
func (t NSTextList) MarkerFormat() NSTextListMarkerFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("markerFormat"))
	return NSTextListMarkerFormat(foundation.NSStringFromID(rv).String())
}

// See: https://developer.apple.com/documentation/AppKit/NSTextList/isOrdered
func (t NSTextList) Ordered() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isOrdered"))
	return rv
}

// Returns the list options mask value of the receiver.
//
// # Return Value
// 
// The list options mask value of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/listOptions
func (t NSTextList) ListOptions() NSTextListOptions {
	rv := objc.Send[NSTextListOptions](t.ID, objc.Sel("listOptions"))
	return NSTextListOptions(rv)
}

// Sets the starting item number for the text list.
//
// # Discussion
// 
// The default value is `1`. This value will be used only for ordered lists,
// and ignored in other cases.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/startingItemNumber
func (t NSTextList) StartingItemNumber() int {
	rv := objc.Send[int](t.ID, objc.Sel("startingItemNumber"))
	return rv
}
func (t NSTextList) SetStartingItemNumber(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setStartingItemNumber:"), value)
}

// The text lists that contain the paragraph.
//
// See: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (t NSTextList) TextLists() INSTextList {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLists"))
	return NSTextListFromID(objc.ID(rv))
}
func (t NSTextList) SetTextLists(value INSTextList) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextLists:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextList/includesTextListMarkers
func (_NSTextListClass NSTextListClass) IncludesTextListMarkers() bool {
	rv := objc.Send[bool](objc.ID(_NSTextListClass.class), objc.Sel("includesTextListMarkers"))
	return rv
}

