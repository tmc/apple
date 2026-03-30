// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextSelection] class.
var (
	_NSTextSelectionClass     NSTextSelectionClass
	_NSTextSelectionClassOnce sync.Once
)

func getNSTextSelectionClass() NSTextSelectionClass {
	_NSTextSelectionClassOnce.Do(func() {
		_NSTextSelectionClass = NSTextSelectionClass{class: objc.GetClass("NSTextSelection")}
	})
	return _NSTextSelectionClass
}

// GetNSTextSelectionClass returns the class object for NSTextSelection.
func GetNSTextSelectionClass() NSTextSelectionClass {
	return getNSTextSelectionClass()
}

type NSTextSelectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextSelectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextSelectionClass) Alloc() NSTextSelection {
	rv := objc.Send[NSTextSelection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a single logical selection context that corresponds
// to an insertion point.
//
// # Creating a text selection
//
//   - [NSTextSelection.InitWithLocationAffinity]: Creates a new text selection with the location and selection affinity you provide.
//   - [NSTextSelection.InitWithRangeAffinityGranularity]: Creates a new text selection with the range, selection affinity, and granularity you provide.
//   - [NSTextSelection.InitWithRangesAffinityGranularity]: Creates a new text selection with the ranges, selection affinity, and granularity you provide.
//   - [NSTextSelection.InitWithCoder]: Creates a test selection from data in an unarchiver.
//
// # Characteristics of a selection
//
//   - [NSTextSelection.Affinity]: Returns the selection affinity of the text selection.
//   - [NSTextSelection.AnchorPositionOffset]: Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//   - [NSTextSelection.SetAnchorPositionOffset]
//   - [NSTextSelection.Granularity]: The granularity of the selection.
//   - [NSTextSelection.Logical]: A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//   - [NSTextSelection.SetLogical]
//   - [NSTextSelection.Transient]: A Boolean value that indicates transient text selection during drag handling.
//   - [NSTextSelection.SecondarySelectionLocation]: Specifies the secondary character location when user taps or clicks at a directional boundary.
//   - [NSTextSelection.SetSecondarySelectionLocation]
//   - [NSTextSelection.TextRanges]: Represents an array of noncontiguous logical ranges in the selection.
//   - [NSTextSelection.TypingAttributes]: The template attributes the framework uses for characters that replace the contents of this selection.
//   - [NSTextSelection.SetTypingAttributes]
//
// # Creating subselections
//
//   - [NSTextSelection.TextSelectionWithTextRanges]: Creates a subselection of the current text selection with the ranges you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection
type NSTextSelection struct {
	objectivec.Object
}

// NSTextSelectionFromID constructs a [NSTextSelection] from an objc.ID.
//
// A class that represents a single logical selection context that corresponds
// to an insertion point.
func NSTextSelectionFromID(id objc.ID) NSTextSelection {
	return NSTextSelection{objectivec.Object{ID: id}}
}

// NOTE: NSTextSelection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextSelection] class.
//
// # Creating a text selection
//
//   - [INSTextSelection.InitWithLocationAffinity]: Creates a new text selection with the location and selection affinity you provide.
//   - [INSTextSelection.InitWithRangeAffinityGranularity]: Creates a new text selection with the range, selection affinity, and granularity you provide.
//   - [INSTextSelection.InitWithRangesAffinityGranularity]: Creates a new text selection with the ranges, selection affinity, and granularity you provide.
//   - [INSTextSelection.InitWithCoder]: Creates a test selection from data in an unarchiver.
//
// # Characteristics of a selection
//
//   - [INSTextSelection.Affinity]: Returns the selection affinity of the text selection.
//   - [INSTextSelection.AnchorPositionOffset]: Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
//   - [INSTextSelection.SetAnchorPositionOffset]
//   - [INSTextSelection.Granularity]: The granularity of the selection.
//   - [INSTextSelection.Logical]: A Boolean value that indicates whether the framework interprets the selection as logical or visual.
//   - [INSTextSelection.SetLogical]
//   - [INSTextSelection.Transient]: A Boolean value that indicates transient text selection during drag handling.
//   - [INSTextSelection.SecondarySelectionLocation]: Specifies the secondary character location when user taps or clicks at a directional boundary.
//   - [INSTextSelection.SetSecondarySelectionLocation]
//   - [INSTextSelection.TextRanges]: Represents an array of noncontiguous logical ranges in the selection.
//   - [INSTextSelection.TypingAttributes]: The template attributes the framework uses for characters that replace the contents of this selection.
//   - [INSTextSelection.SetTypingAttributes]
//
// # Creating subselections
//
//   - [INSTextSelection.TextSelectionWithTextRanges]: Creates a subselection of the current text selection with the ranges you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection
type INSTextSelection interface {
	objectivec.IObject

	// Topic: Creating a text selection

	// Creates a new text selection with the location and selection affinity you provide.
	InitWithLocationAffinity(location NSTextLocation, affinity NSTextSelectionAffinity) NSTextSelection
	// Creates a new text selection with the range, selection affinity, and granularity you provide.
	InitWithRangeAffinityGranularity(range_ INSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection
	// Creates a new text selection with the ranges, selection affinity, and granularity you provide.
	InitWithRangesAffinityGranularity(textRanges []NSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection
	// Creates a test selection from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSTextSelection

	// Topic: Characteristics of a selection

	// Returns the selection affinity of the text selection.
	Affinity() NSTextSelectionAffinity
	// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location.
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	// The granularity of the selection.
	Granularity() NSTextSelectionGranularity
	// A Boolean value that indicates whether the framework interprets the selection as logical or visual.
	Logical() bool
	SetLogical(value bool)
	// A Boolean value that indicates transient text selection during drag handling.
	Transient() bool
	// Specifies the secondary character location when user taps or clicks at a directional boundary.
	SecondarySelectionLocation() NSTextLocation
	SetSecondarySelectionLocation(value NSTextLocation)
	// Represents an array of noncontiguous logical ranges in the selection.
	TextRanges() []NSTextRange
	// The template attributes the framework uses for characters that replace the contents of this selection.
	TypingAttributes() foundation.INSDictionary
	SetTypingAttributes(value foundation.INSDictionary)

	// Topic: Creating subselections

	// Creates a subselection of the current text selection with the ranges you specify.
	TextSelectionWithTextRanges(textRanges []NSTextRange) INSTextSelection

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextSelection) Init() NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextSelection) Autorelease() NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextSelection creates a new NSTextSelection instance.
func NewNSTextSelection() NSTextSelection {
	class := getNSTextSelectionClass()
	rv := objc.Send[NSTextSelection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a test selection from data in an unarchiver.
//
// coder: A coder that subclasses [NSCoder].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(coder:)
//
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
func NewTextSelectionWithCoder(coder foundation.INSCoder) NSTextSelection {
	instance := getNSTextSelectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextSelectionFromID(rv)
}

// Creates a new text selection with the location and selection affinity you
// provide.
//
// location: The text location
//
// affinity: One of the possible [NSTextSelection.Affinity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
func NewTextSelectionWithLocationAffinity(location NSTextLocation, affinity NSTextSelectionAffinity) NSTextSelection {
	instance := getNSTextSelectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:affinity:"), location, affinity)
	return NSTextSelectionFromID(rv)
}

// Creates a new text selection with the range, selection affinity, and
// granularity you provide.
//
// range: The range of the selection.
//
// affinity: One of the available [NSTextSelection.Affinity] options.
//
// granularity: One of the available [NSTextSelection.Granularity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(range:affinity:granularity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
func NewTextSelectionWithRangeAffinityGranularity(range_ INSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection {
	instance := getNSTextSelectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRange:affinity:granularity:"), range_, affinity, granularity)
	return NSTextSelectionFromID(rv)
}

// Creates a new text selection with the ranges, selection affinity, and
// granularity you provide.
//
// textRanges: An array of text ranges.
//
// affinity: One of the available [NSTextSelection.Affinity] options.
//
// granularity: One of the available [NSTextSelection.Granularity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:granularity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
func NewTextSelectionWithRangesAffinityGranularity(textRanges []NSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection {
	instance := getNSTextSelectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRanges:affinity:granularity:"), objectivec.IObjectSliceToNSArray(textRanges), affinity, granularity)
	return NSTextSelectionFromID(rv)
}

// Creates a new text selection with the location and selection affinity you
// provide.
//
// location: The text location
//
// affinity: One of the possible [NSTextSelection.Affinity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
func (t NSTextSelection) InitWithLocationAffinity(location NSTextLocation, affinity NSTextSelectionAffinity) NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("initWithLocation:affinity:"), location, affinity)
	return rv
}

// Creates a new text selection with the range, selection affinity, and
// granularity you provide.
//
// range: The range of the selection.
//
// affinity: One of the available [NSTextSelection.Affinity] options.
//
// granularity: One of the available [NSTextSelection.Granularity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(range:affinity:granularity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
func (t NSTextSelection) InitWithRangeAffinityGranularity(range_ INSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("initWithRange:affinity:granularity:"), range_, affinity, granularity)
	return rv
}

// Creates a new text selection with the ranges, selection affinity, and
// granularity you provide.
//
// textRanges: An array of text ranges.
//
// affinity: One of the available [NSTextSelection.Affinity] options.
//
// granularity: One of the available [NSTextSelection.Granularity] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(_:affinity:granularity:)
//
// [NSTextSelection.Affinity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
func (t NSTextSelection) InitWithRangesAffinityGranularity(textRanges []NSTextRange, affinity NSTextSelectionAffinity, granularity NSTextSelectionGranularity) NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("initWithRanges:affinity:granularity:"), objectivec.IObjectSliceToNSArray(textRanges), affinity, granularity)
	return rv
}

// Creates a test selection from data in an unarchiver.
//
// coder: A coder that subclasses [NSCoder].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/init(coder:)
//
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
func (t NSTextSelection) InitWithCoder(coder foundation.INSCoder) NSTextSelection {
	rv := objc.Send[NSTextSelection](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates a subselection of the current text selection with the ranges you
// specify.
//
// textRanges: An array of text ranges.
//
// # Return Value
//
// A new [NSTextSelection].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/textSelection(_:)
func (t NSTextSelection) TextSelectionWithTextRanges(textRanges []NSTextRange) INSTextSelection {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelectionWithTextRanges:"), objectivec.IObjectSliceToNSArray(textRanges))
	return NSTextSelectionFromID(rv)
}
func (t NSTextSelection) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the selection affinity of the text selection.
//
// # Discussion
//
// The affinity is [NSTextSelectionAffinityDownstream] by default. For a
// zero-length selection, it describes the visual location of the text cursor
// between the head of line containing the selection location (downstream) or
// tail of the previous line (upstream). For a selection with contents, it
// describes the logical direction of non-anchored edge of the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/affinity-swift.property
func (t NSTextSelection) Affinity() NSTextSelectionAffinity {
	rv := objc.Send[NSTextSelectionAffinity](t.ID, objc.Sel("affinity"))
	return NSTextSelectionAffinity(rv)
}

// Represents the anchor position offset from the beginning of a line fragment
// in the visual order for the initial tap or click location.
//
// # Discussion
//
// That starts from the left for a horizontal line fragment, and from the top
// for a vertical. Navigating between lines uses this point when the current
// line fragment associated with the selection is shorter than the next line
// visited. Defaults to `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/anchorPositionOffset
func (t NSTextSelection) AnchorPositionOffset() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("anchorPositionOffset"))
	return rv
}
func (t NSTextSelection) SetAnchorPositionOffset(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setAnchorPositionOffset:"), value)
}

// The granularity of the selection.
//
// # Discussion
//
// The default is [NSTextSelectionGranularityCharacter]. Selection extending
// operations modify the selection by `granularity`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/granularity-swift.property
func (t NSTextSelection) Granularity() NSTextSelectionGranularity {
	rv := objc.Send[NSTextSelectionGranularity](t.ID, objc.Sel("granularity"))
	return NSTextSelectionGranularity(rv)
}

// A Boolean value that indicates whether the framework interprets the
// selection as logical or visual.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/isLogical
func (t NSTextSelection) Logical() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isLogical"))
	return rv
}
func (t NSTextSelection) SetLogical(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setLogical:"), value)
}

// A Boolean value that indicates transient text selection during drag
// handling.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/isTransient
func (t NSTextSelection) Transient() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isTransient"))
	return rv
}

// Specifies the secondary character location when user taps or clicks at a
// directional boundary.
//
// # Discussion
//
// Setting a non-`nil` location has a side effect of making `isLogical` to
// `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/secondarySelectionLocation
func (t NSTextSelection) SecondarySelectionLocation() NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("secondarySelectionLocation"))
	return NSTextLocationObjectFromID(rv)
}
func (t NSTextSelection) SetSecondarySelectionLocation(value NSTextLocation) {
	objc.Send[struct{}](t.ID, objc.Sel("setSecondarySelectionLocation:"), value)
}

// Represents an array of noncontiguous logical ranges in the selection.
//
// # Discussion
//
// The array must be logically ordered. When editing, all ranges in a text
// selection constitute a single insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/textRanges
func (t NSTextSelection) TextRanges() []NSTextRange {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextRange {
		return NSTextRangeFromID(id)
	})
}

// The template attributes the framework uses for characters that replace the
// contents of this selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelection/typingAttributes
func (t NSTextSelection) TypingAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("typingAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextSelection) SetTypingAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setTypingAttributes:"), value)
}
