// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextLineFragment] class.
var (
	_NSTextLineFragmentClass     NSTextLineFragmentClass
	_NSTextLineFragmentClassOnce sync.Once
)

func getNSTextLineFragmentClass() NSTextLineFragmentClass {
	_NSTextLineFragmentClassOnce.Do(func() {
		_NSTextLineFragmentClass = NSTextLineFragmentClass{class: objc.GetClass("NSTextLineFragment")}
	})
	return _NSTextLineFragmentClass
}

// GetNSTextLineFragmentClass returns the class object for NSTextLineFragment.
func GetNSTextLineFragmentClass() NSTextLineFragmentClass {
	return getNSTextLineFragmentClass()
}

type NSTextLineFragmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextLineFragmentClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextLineFragmentClass) Alloc() NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a line fragment as a single textual layout and
// rendering unit inside a text layout fragment.
//
// # Creating line fragments
//
//   - [NSTextLineFragment.InitWithAttributedStringRange]: Creates a new line fragment from the attributed string for the range of characters you specify.
//   - [NSTextLineFragment.InitWithCoder]: Creates a new line fragment with from data in an unarchiver.
//   - [NSTextLineFragment.InitWithStringAttributesRange]: Creates a new line fragment using the string, attributes, and range you provide.
//
// # Line fragment characteristics
//
//   - [NSTextLineFragment.AttributedString]: The source attributed string.
//   - [NSTextLineFragment.CharacterRange]: The string range for the source attributed string that corresponds to this line fragment.
//   - [NSTextLineFragment.GlyphOrigin]: Rendering origin for the left-most glyph in the line fragment coordinate system.
//   - [NSTextLineFragment.TypographicBounds]: The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.
//
// # Finding specific text
//
//   - [NSTextLineFragment.CharacterIndexForPoint]: Returns character index for a point inside the line fragment coordinate system.
//   - [NSTextLineFragment.FractionOfDistanceThroughGlyphForPoint]: Returns character index for a point inside the line fragment coordinate system.
//   - [NSTextLineFragment.LocationForCharacterAtIndex]: Returns the location of the character at the specified index.
//
// # Drawing
//
//   - [NSTextLineFragment.DrawAtPointInContext]: Renders the line fragment contents at the rendering origin.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment
type NSTextLineFragment struct {
	objectivec.Object
}

// NSTextLineFragmentFromID constructs a [NSTextLineFragment] from an objc.ID.
//
// A class that represents a line fragment as a single textual layout and
// rendering unit inside a text layout fragment.
func NSTextLineFragmentFromID(id objc.ID) NSTextLineFragment {
	return NSTextLineFragment{objectivec.Object{ID: id}}
}
// NOTE: NSTextLineFragment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextLineFragment] class.
//
// # Creating line fragments
//
//   - [INSTextLineFragment.InitWithAttributedStringRange]: Creates a new line fragment from the attributed string for the range of characters you specify.
//   - [INSTextLineFragment.InitWithCoder]: Creates a new line fragment with from data in an unarchiver.
//   - [INSTextLineFragment.InitWithStringAttributesRange]: Creates a new line fragment using the string, attributes, and range you provide.
//
// # Line fragment characteristics
//
//   - [INSTextLineFragment.AttributedString]: The source attributed string.
//   - [INSTextLineFragment.CharacterRange]: The string range for the source attributed string that corresponds to this line fragment.
//   - [INSTextLineFragment.GlyphOrigin]: Rendering origin for the left-most glyph in the line fragment coordinate system.
//   - [INSTextLineFragment.TypographicBounds]: The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.
//
// # Finding specific text
//
//   - [INSTextLineFragment.CharacterIndexForPoint]: Returns character index for a point inside the line fragment coordinate system.
//   - [INSTextLineFragment.FractionOfDistanceThroughGlyphForPoint]: Returns character index for a point inside the line fragment coordinate system.
//   - [INSTextLineFragment.LocationForCharacterAtIndex]: Returns the location of the character at the specified index.
//
// # Drawing
//
//   - [INSTextLineFragment.DrawAtPointInContext]: Renders the line fragment contents at the rendering origin.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment
type INSTextLineFragment interface {
	objectivec.IObject

	// Topic: Creating line fragments

	// Creates a new line fragment from the attributed string for the range of characters you specify.
	InitWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSTextLineFragment
	// Creates a new line fragment with from data in an unarchiver.
	InitWithCoder(aDecoder foundation.INSCoder) NSTextLineFragment
	// Creates a new line fragment using the string, attributes, and range you provide.
	InitWithStringAttributesRange(string_ string, attributes foundation.INSDictionary, range_ foundation.NSRange) NSTextLineFragment

	// Topic: Line fragment characteristics

	// The source attributed string.
	AttributedString() foundation.NSAttributedString
	// The string range for the source attributed string that corresponds to this line fragment.
	CharacterRange() foundation.NSRange
	// Rendering origin for the left-most glyph in the line fragment coordinate system.
	GlyphOrigin() corefoundation.CGPoint
	// The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.
	TypographicBounds() corefoundation.CGRect

	// Topic: Finding specific text

	// Returns character index for a point inside the line fragment coordinate system.
	CharacterIndexForPoint(point corefoundation.CGPoint) int
	// Returns character index for a point inside the line fragment coordinate system.
	FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64
	// Returns the location of the character at the specified index.
	LocationForCharacterAtIndex(index int) corefoundation.CGPoint

	// Topic: Drawing

	// Renders the line fragment contents at the rendering origin.
	DrawAtPointInContext(point corefoundation.CGPoint, context coregraphics.CGContextRef)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextLineFragment) Init() NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextLineFragment) Autorelease() NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextLineFragment creates a new NSTextLineFragment instance.
func NewNSTextLineFragment() NSTextLineFragment {
	class := getNSTextLineFragmentClass()
	rv := objc.Send[NSTextLineFragment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new line fragment from the attributed string for the range of
// characters you specify.
//
// attributedString: The attributed string.
//
// range: An [NSRange] that specifies which characters to include.
// //
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(attributedString:range:)
func NewTextLineFragmentWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSTextLineFragment {
	instance := getNSTextLineFragmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:range:"), attributedString, range_)
	return NSTextLineFragmentFromID(rv)
}

// Creates a new line fragment with from data in an unarchiver.
//
// aDecoder: A decoder that conforms to the [NSCoder] protocol.
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(coder:)
func NewTextLineFragmentWithCoder(aDecoder foundation.INSCoder) NSTextLineFragment {
	instance := getNSTextLineFragmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return NSTextLineFragmentFromID(rv)
}

// Creates a new line fragment using the string, attributes, and range you
// provide.
//
// string: An attributed string.
//
// attributes: A dictionary of attributes.
//
// range: The range to use from `string`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(string:attributes:range:)
func NewTextLineFragmentWithStringAttributesRange(string_ string, attributes foundation.INSDictionary, range_ foundation.NSRange) NSTextLineFragment {
	instance := getNSTextLineFragmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:attributes:range:"), objc.String(string_), attributes, range_)
	return NSTextLineFragmentFromID(rv)
}

// Creates a new line fragment from the attributed string for the range of
// characters you specify.
//
// attributedString: The attributed string.
//
// range: An [NSRange] that specifies which characters to include.
// //
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(attributedString:range:)
func (t NSTextLineFragment) InitWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](t.ID, objc.Sel("initWithAttributedString:range:"), attributedString, range_)
	return rv
}
// Creates a new line fragment with from data in an unarchiver.
//
// aDecoder: A decoder that conforms to the [NSCoder] protocol.
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(coder:)
func (t NSTextLineFragment) InitWithCoder(aDecoder foundation.INSCoder) NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](t.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}
// Creates a new line fragment using the string, attributes, and range you
// provide.
//
// string: An attributed string.
//
// attributes: A dictionary of attributes.
//
// range: The range to use from `string`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/init(string:attributes:range:)
func (t NSTextLineFragment) InitWithStringAttributesRange(string_ string, attributes foundation.INSDictionary, range_ foundation.NSRange) NSTextLineFragment {
	rv := objc.Send[NSTextLineFragment](t.ID, objc.Sel("initWithString:attributes:range:"), objc.String(string_), attributes, range_)
	return rv
}
// Returns character index for a point inside the line fragment coordinate
// system.
//
// point: The distance is from the upstream edge.
//
// # Return Value
// 
// An integer that represents the character index at `point`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/characterIndex(for:)
func (t NSTextLineFragment) CharacterIndexForPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t.ID, objc.Sel("characterIndexForPoint:"), point)
	return rv
}
// Returns character index for a point inside the line fragment coordinate
// system.
//
// point: A [CGPoint] that represents the point inside the line fragment.
// //
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// # Return Value
// 
// The fraction of distance from the upstream edge.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/fractionOfDistanceThroughGlyph(for:)
func (t NSTextLineFragment) FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
}
// Returns the location of the character at the specified index.
//
// index: An integer that represents the position in the text.
//
// # Return Value
// 
// A [CGPoint] that’s on the upstream edge of the glyph. It’s in the
// coordinate system relative to the line fragment origin.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/locationForCharacter(at:)
func (t NSTextLineFragment) LocationForCharacterAtIndex(index int) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("locationForCharacterAtIndex:"), index)
	return corefoundation.CGPoint(rv)
}
// Renders the line fragment contents at the rendering origin.
//
// point: The origin as a [CGPoint].
//
// context: The drawing context.
//
// # Discussion
// 
// You can specify the origin as (`NSMinX(typographicBounds) +
// glyphOrigin.X(), NSMinY(typographicBounds) + glyphOrigin.Y())` relative to
// the line fragment group coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/draw(at:in:)
func (t NSTextLineFragment) DrawAtPointInContext(point corefoundation.CGPoint, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawAtPoint:inContext:"), point, context)
}
func (t NSTextLineFragment) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The source attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/attributedString
func (t NSTextLineFragment) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
// The string range for the source attributed string that corresponds to this
// line fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/characterRange
func (t NSTextLineFragment) CharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("characterRange"))
	return foundation.NSRange(rv)
}
// Rendering origin for the left-most glyph in the line fragment coordinate
// system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/glyphOrigin
func (t NSTextLineFragment) GlyphOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("glyphOrigin"))
	return corefoundation.CGPoint(rv)
}
// The typographic bounds that specifies the dimensions of the line fragment
// for laying out line fragments to each other.
//
// # Discussion
// 
// The origin value is offset from the beginning of the line fragment group
// belonging to the parent layout fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLineFragment/typographicBounds
func (t NSTextLineFragment) TypographicBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("typographicBounds"))
	return corefoundation.CGRect(rv)
}

