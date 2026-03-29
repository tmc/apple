// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGlyphInfo] class.
var (
	_NSGlyphInfoClass     NSGlyphInfoClass
	_NSGlyphInfoClassOnce sync.Once
)

func getNSGlyphInfoClass() NSGlyphInfoClass {
	_NSGlyphInfoClassOnce.Do(func() {
		_NSGlyphInfoClass = NSGlyphInfoClass{class: objc.GetClass("NSGlyphInfo")}
	})
	return _NSGlyphInfoClass
}

// GetNSGlyphInfoClass returns the class object for NSGlyphInfo.
func GetNSGlyphInfoClass() NSGlyphInfoClass {
	return getNSGlyphInfoClass()
}

type NSGlyphInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGlyphInfoClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGlyphInfoClass) Alloc() NSGlyphInfo {
	rv := objc.Send[NSGlyphInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A glyph attribute in an attributed string.
//
// # Overview
// 
// Glyphs are the graphic representations of characters, stored in a font,
// that the text system draws on a display or printed page. Before text can be
// laid out, the layout manager (<[NSLayoutManager]) generates a stream of
// glyphs, using the character and font information specified by the
// attributed string and contained in the font file. [NSGlyphInfo] represents
// a glyph attribute value ([NSGlyphInfoAttributeName]) in an attributed
// string ([NSAttributedString]) and provides a means to override the standard
// glyph generation process and substitute a specified glyph over the
// attribute’s range.
// 
// Glyph attributes are integer values that the layout manager uses to denote
// special handling for particular glyphs during rendering. [NSGlyphInfo]
// enables you to override a font’s built-in mapping from a Unicode
// character code to a corresponding glyph ID. Overriding the mapping allows
// you to specify a variant glyph for a given character if the font contains
// multiple variations for that character or to specify a glyph that doesn’t
// have a standard mapping (such as some ligature glyphs).
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSGlyphInfoAttributeName]: https://developer.apple.com/documentation/AppKit/NSGlyphInfoAttributeName
//
// # Getting information about a glyph info object
//
//   - [NSGlyphInfo.BaseString]: The string containing the character represented by the glyph.
//   - [NSGlyphInfo.GlyphID]: The glyph identifier, specified as the index into the internal glyph table of the font.
//
// # Deprecated
//
//   - [NSGlyphInfo.CharacterIdentifier]: The receiver’s character identifier (CID).
//   - [NSGlyphInfo.CharacterCollection]: A value specifying the glyph–to–character identifier mapping of the receiver.
//   - [NSGlyphInfo.GlyphName]: The receiver’s glyph name.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo
type NSGlyphInfo struct {
	objectivec.Object
}

// NSGlyphInfoFromID constructs a [NSGlyphInfo] from an objc.ID.
//
// A glyph attribute in an attributed string.
func NSGlyphInfoFromID(id objc.ID) NSGlyphInfo {
	return NSGlyphInfo{objectivec.Object{ID: id}}
}
// NOTE: NSGlyphInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGlyphInfo] class.
//
// # Getting information about a glyph info object
//
//   - [INSGlyphInfo.BaseString]: The string containing the character represented by the glyph.
//   - [INSGlyphInfo.GlyphID]: The glyph identifier, specified as the index into the internal glyph table of the font.
//
// # Deprecated
//
//   - [INSGlyphInfo.CharacterIdentifier]: The receiver’s character identifier (CID).
//   - [INSGlyphInfo.CharacterCollection]: A value specifying the glyph–to–character identifier mapping of the receiver.
//   - [INSGlyphInfo.GlyphName]: The receiver’s glyph name.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo
type INSGlyphInfo interface {
	objectivec.IObject

	// Topic: Getting information about a glyph info object

	// The string containing the character represented by the glyph.
	BaseString() string
	// The glyph identifier, specified as the index into the internal glyph table of the font.
	GlyphID() coregraphics.CGFontIndex

	// Topic: Deprecated

	// The receiver’s character identifier (CID).
	CharacterIdentifier() uint
	// A value specifying the glyph–to–character identifier mapping of the receiver.
	CharacterCollection() NSCharacterCollection
	// The receiver’s glyph name.
	GlyphName() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g NSGlyphInfo) Init() NSGlyphInfo {
	rv := objc.Send[NSGlyphInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGlyphInfo) Autorelease() NSGlyphInfo {
	rv := objc.Send[NSGlyphInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGlyphInfo creates a new NSGlyphInfo instance.
func NewNSGlyphInfo() NSGlyphInfo {
	class := getNSGlyphInfoClass()
	rv := objc.Send[NSGlyphInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a glyph info object from the specified glyph identifier and font
// informaton.
//
// glyph: The requested [CGGlyph] object.
// //
// [CGGlyph]: https://developer.apple.com/documentation/CoreGraphics/CGGlyph
//
// font: The font containing the glyph.
//
// string: A string containing the character represented by the glyph.
//
// # Return Value
// 
// A glyph info object for the specified glyph or `nil` if the glyph
// information is not available.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(cgGlyph:for:baseString:)
func NewGlyphInfoWithCGGlyphForFontBaseString(glyph coregraphics.CGFontIndex, font NSFont, string_ string) NSGlyphInfo {
	rv := objc.Send[objc.ID](objc.ID(getNSGlyphInfoClass().class), objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return NSGlyphInfoFromID(rv)
}

// Instantiates and returns an [NSGlyphInfo] object using a character
// identifier and a character collection.
//
// cid: A character identifier.
//
// characterCollection: A string constant representing a character collection. Possible values are
// described in [NSCharacterCollection].
// //
// [NSCharacterCollection]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection
//
// string: The part of the attributed string the returned instance is intended to
// override.
//
// # Return Value
// 
// The created [NSGlyphInfo] object or `nil` if the object couldn’t be
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(characterIdentifier:collection:baseString:)
func NewGlyphInfoWithCharacterIdentifierCollectionBaseString(cid uint, characterCollection NSCharacterCollection, string_ string) NSGlyphInfo {
	rv := objc.Send[objc.ID](objc.ID(getNSGlyphInfoClass().class), objc.Sel("glyphInfoWithCharacterIdentifier:collection:baseString:"), cid, characterCollection, objc.String(string_))
	return NSGlyphInfoFromID(rv)
}

// Instantiates and returns a glyph information object using a glyph index and
// a specified font.
//
// glyph: The identifier of the glyph.
//
// font: The font object to be associated with the returned [NSGlyphInfo] object,
//
// string: The part of the attributed string the returned instance is intended to
// override.
//
// # Return Value
// 
// The created [NSGlyphInfo] object or `nil` if the object couldn’t be
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyph:forFont:baseString:)
func NewGlyphInfoWithGlyphForFontBaseString(glyph NSGlyph, font NSFont, string_ string) NSGlyphInfo {
	rv := objc.Send[objc.ID](objc.ID(getNSGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return NSGlyphInfoFromID(rv)
}

// Instantiates and returns a glyph information object using a glyph name and
// a specified font.
//
// glyphName: The name of the glyph.
//
// font: The font object to be associated with the returned [NSGlyphInfo] object,
//
// string: The part of the attributed string the returned instance is intended to
// override.
//
// # Return Value
// 
// The created [NSGlyphInfo] object or `nil` if the object couldn’t be
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyphName:forFont:baseString:)
func NewGlyphInfoWithGlyphNameForFontBaseString(glyphName string, font NSFont, string_ string) NSGlyphInfo {
	rv := objc.Send[objc.ID](objc.ID(getNSGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyphName:forFont:baseString:"), objc.String(glyphName), font, objc.String(string_))
	return NSGlyphInfoFromID(rv)
}

func (g NSGlyphInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The string containing the character represented by the glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/baseString
func (g NSGlyphInfo) BaseString() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("baseString"))
	return foundation.NSStringFromID(rv).String()
}
// The glyph identifier, specified as the index into the internal glyph table
// of the font.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphID
func (g NSGlyphInfo) GlyphID() coregraphics.CGFontIndex {
	rv := objc.Send[coregraphics.CGFontIndex](g.ID, objc.Sel("glyphID"))
	return coregraphics.CGFontIndex(rv)
}
// The receiver’s character identifier (CID).
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterIdentifier
func (g NSGlyphInfo) CharacterIdentifier() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("characterIdentifier"))
	return rv
}
// A value specifying the glyph–to–character identifier mapping of the
// receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterCollection
func (g NSGlyphInfo) CharacterCollection() NSCharacterCollection {
	rv := objc.Send[NSCharacterCollection](g.ID, objc.Sel("characterCollection"))
	return NSCharacterCollection(rv)
}
// The receiver’s glyph name.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphName
func (g NSGlyphInfo) GlyphName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("glyphName"))
	return foundation.NSStringFromID(rv).String()
}

