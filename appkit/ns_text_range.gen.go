// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextRange] class.
var (
	_NSTextRangeClass     NSTextRangeClass
	_NSTextRangeClassOnce sync.Once
)

func getNSTextRangeClass() NSTextRangeClass {
	_NSTextRangeClassOnce.Do(func() {
		_NSTextRangeClass = NSTextRangeClass{class: objc.GetClass("NSTextRange")}
	})
	return _NSTextRangeClass
}

// GetNSTextRangeClass returns the class object for NSTextRange.
func GetNSTextRangeClass() NSTextRangeClass {
	return getNSTextRangeClass()
}

type NSTextRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextRangeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextRangeClass) Alloc() NSTextRange {
	rv := objc.Send[NSTextRange](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a contiguous range between two locations inside
// document contents.
//
// # Overview
// 
// An [NSTextRange] consists of the starting and terminating locations. There
// the two basic properties: [NSTextRange.Location] and [NSTextRange.EndLocation], respectively. The
// terminating [NSTextRange.Location], [NSTextRange.EndLocation], is directly following the last
// location in the range. For example, a location contains a range if
// `(range.Location() <= location) && (location < range.EndLocation())` is
// `true`.
//
// # Creating a text range
//
//   - [NSTextRange.InitWithLocation]: Creates a new text range at the location you specify.
//   - [NSTextRange.InitWithLocationEndLocation]: Creates a new text range with the starting and ending locations you specify.
//
// # Characteristics of the text range
//
//   - [NSTextRange.Location]: The starting location of the text range.
//   - [NSTextRange.EndLocation]: The ending location of the text range.
//   - [NSTextRange.Empty]: Returns whether the text range is empty.
//
// # Comparing text ranges
//
//   - [NSTextRange.TextRangeByIntersectingWithTextRange]: Returns the range, if any, where two text ranges intersect.
//   - [NSTextRange.IntersectsWithTextRange]: Determines if two ranges intersect.
//   - [NSTextRange.IsEqualToTextRange]: Compares two text ranges.
//   - [NSTextRange.TextRangeByFormingUnionWithTextRange]: Returns a new text range by forming the union with the text range you provide.
//
// # Finding text within the text range
//
//   - [NSTextRange.ContainsLocation]: Determines if the text location you specify is in the current text range.
//   - [NSTextRange.ContainsRange]: Determines if the text range you specify is in the current text range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange
type NSTextRange struct {
	objectivec.Object
}

// NSTextRangeFromID constructs a [NSTextRange] from an objc.ID.
//
// A class that represents a contiguous range between two locations inside
// document contents.
func NSTextRangeFromID(id objc.ID) NSTextRange {
	return NSTextRange{objectivec.Object{ID: id}}
}
// NOTE: NSTextRange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextRange] class.
//
// # Creating a text range
//
//   - [INSTextRange.InitWithLocation]: Creates a new text range at the location you specify.
//   - [INSTextRange.InitWithLocationEndLocation]: Creates a new text range with the starting and ending locations you specify.
//
// # Characteristics of the text range
//
//   - [INSTextRange.Location]: The starting location of the text range.
//   - [INSTextRange.EndLocation]: The ending location of the text range.
//   - [INSTextRange.Empty]: Returns whether the text range is empty.
//
// # Comparing text ranges
//
//   - [INSTextRange.TextRangeByIntersectingWithTextRange]: Returns the range, if any, where two text ranges intersect.
//   - [INSTextRange.IntersectsWithTextRange]: Determines if two ranges intersect.
//   - [INSTextRange.IsEqualToTextRange]: Compares two text ranges.
//   - [INSTextRange.TextRangeByFormingUnionWithTextRange]: Returns a new text range by forming the union with the text range you provide.
//
// # Finding text within the text range
//
//   - [INSTextRange.ContainsLocation]: Determines if the text location you specify is in the current text range.
//   - [INSTextRange.ContainsRange]: Determines if the text range you specify is in the current text range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange
type INSTextRange interface {
	objectivec.IObject

	// Topic: Creating a text range

	// Creates a new text range at the location you specify.
	InitWithLocation(location NSTextLocation) NSTextRange
	// Creates a new text range with the starting and ending locations you specify.
	InitWithLocationEndLocation(location NSTextLocation, endLocation NSTextLocation) NSTextRange

	// Topic: Characteristics of the text range

	// The starting location of the text range.
	Location() NSTextLocation
	// The ending location of the text range.
	EndLocation() NSTextLocation
	// Returns whether the text range is empty.
	Empty() bool

	// Topic: Comparing text ranges

	// Returns the range, if any, where two text ranges intersect.
	TextRangeByIntersectingWithTextRange(textRange INSTextRange) INSTextRange
	// Determines if two ranges intersect.
	IntersectsWithTextRange(textRange INSTextRange) bool
	// Compares two text ranges.
	IsEqualToTextRange(textRange INSTextRange) bool
	// Returns a new text range by forming the union with the text range you provide.
	TextRangeByFormingUnionWithTextRange(textRange INSTextRange) INSTextRange

	// Topic: Finding text within the text range

	// Determines if the text location you specify is in the current text range.
	ContainsLocation(location NSTextLocation) bool
	// Determines if the text range you specify is in the current text range.
	ContainsRange(textRange INSTextRange) bool
}

// Init initializes the instance.
func (t NSTextRange) Init() NSTextRange {
	rv := objc.Send[NSTextRange](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextRange) Autorelease() NSTextRange {
	rv := objc.Send[NSTextRange](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextRange creates a new NSTextRange instance.
func NewNSTextRange() NSTextRange {
	class := getNSTextRangeClass()
	rv := objc.Send[NSTextRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new text range at the location you specify.
//
// location: An [NSTextLocation].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:)
func NewTextRangeWithLocation(location NSTextLocation) NSTextRange {
	instance := getNSTextRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:"), location)
	return NSTextRangeFromID(rv)
}

// Creates a new text range with the starting and ending locations you
// specify.
//
// location: The starting location.
//
// endLocation: The ending location.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:end:)
func NewTextRangeWithLocationEndLocation(location NSTextLocation, endLocation NSTextLocation) NSTextRange {
	instance := getNSTextRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:endLocation:"), location, endLocation)
	return NSTextRangeFromID(rv)
}

// Creates a new text range at the location you specify.
//
// location: An [NSTextLocation].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:)
func (t NSTextRange) InitWithLocation(location NSTextLocation) NSTextRange {
	rv := objc.Send[NSTextRange](t.ID, objc.Sel("initWithLocation:"), location)
	return rv
}
// Creates a new text range with the starting and ending locations you
// specify.
//
// location: The starting location.
//
// endLocation: The ending location.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/init(location:end:)
func (t NSTextRange) InitWithLocationEndLocation(location NSTextLocation, endLocation NSTextLocation) NSTextRange {
	rv := objc.Send[NSTextRange](t.ID, objc.Sel("initWithLocation:endLocation:"), location, endLocation)
	return rv
}
// Returns the range, if any, where two text ranges intersect.
//
// textRange: The range used to compare against the current range to evaluate for
// differences.
//
// # Return Value
// 
// An [NSRange] that represents the intersection of the ranges, or `nil` if
// they don’t intersect.
//
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/intersection(_:)
func (t NSTextRange) TextRangeByIntersectingWithTextRange(textRange INSTextRange) INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textRangeByIntersectingWithTextRange:"), textRange)
	return NSTextRangeFromID(rv)
}
// Determines if two ranges intersect.
//
// textRange: The range used to compare against the current range to evaluate for
// differences.
//
// # Return Value
// 
// Returns `true` if the ranges intersect.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/intersects(_:)
func (t NSTextRange) IntersectsWithTextRange(textRange INSTextRange) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("intersectsWithTextRange:"), textRange)
	return rv
}
// Compares two text ranges.
//
// textRange: The range used to compare against the current range to evaluate for
// differences.
//
// # Return Value
// 
// Returns `true` if the ranges are equal.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/isEqual(to:)
func (t NSTextRange) IsEqualToTextRange(textRange INSTextRange) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEqualToTextRange:"), textRange)
	return rv
}
// Returns a new text range by forming the union with the text range you
// provide.
//
// textRange: The range to use to create the union.
//
// # Return Value
// 
// An [NSTextRange] that represent the union of the two ranges.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/union(_:)
func (t NSTextRange) TextRangeByFormingUnionWithTextRange(textRange INSTextRange) INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textRangeByFormingUnionWithTextRange:"), textRange)
	return NSTextRangeFromID(rv)
}
// Determines if the text location you specify is in the current text range.
//
// location: An [NSTextLocation].
//
// # Return Value
// 
// Returns `true` if the location is in the range otherwise `false` .
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-7hvi0
func (t NSTextRange) ContainsLocation(location NSTextLocation) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("containsLocation:"), location)
	return rv
}
// Determines if the text range you specify is in the current text range.
//
// textRange: An [NSTextRange].
//
// # Return Value
// 
// Returns `true` if the range you provide is in the current range; otherwise
// `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/contains(_:)-5j4y2
func (t NSTextRange) ContainsRange(textRange INSTextRange) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("containsRange:"), textRange)
	return rv
}

// The starting location of the text range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/location
func (t NSTextRange) Location() NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("location"))
	return NSTextLocationObjectFromID(rv)
}
// The ending location of the text range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/endLocation
func (t NSTextRange) EndLocation() NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("endLocation"))
	return NSTextLocationObjectFromID(rv)
}
// Returns whether the text range is empty.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextRange/isEmpty
func (t NSTextRange) Empty() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEmpty"))
	return rv
}

