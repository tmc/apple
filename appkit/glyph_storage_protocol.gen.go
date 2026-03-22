// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that a glyph storage object must implement to interact properly with [NSGlyphGenerator](<doc://com.apple.appkit/documentation/AppKit/NSGlyphGenerator>).
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage
type NSGlyphStorage interface {
	objectivec.IObject

	// Returns the text storage object from which the [NSGlyphGenerator] object procures characters for glyph generation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/attributedString()
	AttributedString() foundation.NSAttributedString

	// Returns the current layout options.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/layoutOptions()
	LayoutOptions() uint

	// Inserts the given glyphs into the glyph cache and maps them to the specified characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/insertGlyphs(_:length:forStartingGlyphAt:characterIndex:)
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs NSGlyph, length uint, glyphIndex uint, charIndex uint)

	// Sets a custom attribute value for a given glyph.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/setIntAttribute(_:value:forGlyphAt:)
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
}

// NSGlyphStorageObject wraps an existing Objective-C object that conforms to the NSGlyphStorage protocol.
type NSGlyphStorageObject struct {
	objectivec.Object
}
func (o NSGlyphStorageObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSGlyphStorageObjectFromID constructs a [NSGlyphStorageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSGlyphStorageObjectFromID(id objc.ID) NSGlyphStorageObject {
	return NSGlyphStorageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the text storage object from which the [NSGlyphGenerator] object
// procures characters for glyph generation.
//
// # Return Value
// 
// The receiver’s text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/attributedString()
func (o NSGlyphStorageObject) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(rv)
	}
// Returns the current layout options.
//
// # Return Value
// 
// The layout options as a bit mask, as defined in Constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/layoutOptions()
func (o NSGlyphStorageObject) LayoutOptions() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("layoutOptions"))
	return rv
	}
// Inserts the given glyphs into the glyph cache and maps them to the
// specified characters.
//
// glyphs: The glyphs to insert.
//
// length: Number of glyphs to insert.
//
// glyphIndex: Location in the glyph cache to begin inserting glyphs.
//
// charIndex: Index of first character to be mapped.
//
// # Discussion
// 
// This is a bulk insert method for the glyph cache.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/insertGlyphs(_:length:forStartingGlyphAt:characterIndex:)
func (o NSGlyphStorageObject) InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs NSGlyph, length uint, glyphIndex uint, charIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
	}
// Sets a custom attribute value for a given glyph.
//
// attributeTag: The custom attribute.
//
// val: The new attribute value.
//
// glyphIndex: Index of the glyph whose attribute is set.
//
// # Discussion
// 
// Custom attributes are glyph attributes such as [NSGlyphInscription] or
// attributes defined by subclasses. Subclasses that define their own custom
// attributes must override this method and provide their own storage for the
// attribute values. Nonnegative tags are reserved; you can define your own
// attributes with negative tags and set values using this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/setIntAttribute(_:value:forGlyphAt:)
func (o NSGlyphStorageObject) SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
	}

