// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGlyphGenerator] class.
var (
	_NSGlyphGeneratorClass     NSGlyphGeneratorClass
	_NSGlyphGeneratorClassOnce sync.Once
)

func getNSGlyphGeneratorClass() NSGlyphGeneratorClass {
	_NSGlyphGeneratorClassOnce.Do(func() {
		_NSGlyphGeneratorClass = NSGlyphGeneratorClass{class: objc.GetClass("NSGlyphGenerator")}
	})
	return _NSGlyphGeneratorClass
}

// GetNSGlyphGeneratorClass returns the class object for NSGlyphGenerator.
func GetNSGlyphGeneratorClass() NSGlyphGeneratorClass {
	return getNSGlyphGeneratorClass()
}

type NSGlyphGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGlyphGeneratorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGlyphGeneratorClass) Alloc() NSGlyphGenerator {
	rv := objc.Send[NSGlyphGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that performs the initial, nominal glyph generation phase in the
// layout process.
//
// # Overview
//
// The nominal glyph generation pass essentially generates one glyph per
// character; the typesetter may later make substitutions in the glyph stream,
// for example, changing an acute accent glyph followed by an “e” glyph
// into a single acute-accented “é” glyph.
//
// [NSGlyphGenerator] communicates via the [NSGlyphStorage] protocol. An
// example of a class that conforms to the protocol is [NSLayoutManager].
//
// # Generating glyphs
//
//   - [NSGlyphGenerator.GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex]: Generates glyphs for the specified glyph storage object ([NSLayoutManager] by default).
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator
type NSGlyphGenerator struct {
	objectivec.Object
}

// NSGlyphGeneratorFromID constructs a [NSGlyphGenerator] from an objc.ID.
//
// An object that performs the initial, nominal glyph generation phase in the
// layout process.
func NSGlyphGeneratorFromID(id objc.ID) NSGlyphGenerator {
	return NSGlyphGenerator{objectivec.Object{ID: id}}
}

// NOTE: NSGlyphGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGlyphGenerator] class.
//
// # Generating glyphs
//
//   - [INSGlyphGenerator.GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex]: Generates glyphs for the specified glyph storage object ([NSLayoutManager] by default).
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator
type INSGlyphGenerator interface {
	objectivec.IObject

	// Topic: Generating glyphs

	// Generates glyphs for the specified glyph storage object ([NSLayoutManager] by default).
	GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage NSGlyphStorage, nChars uint, glyphIndex unsafe.Pointer, charIndex unsafe.Pointer)
}

// Init initializes the instance.
func (g NSGlyphGenerator) Init() NSGlyphGenerator {
	rv := objc.Send[NSGlyphGenerator](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGlyphGenerator) Autorelease() NSGlyphGenerator {
	rv := objc.Send[NSGlyphGenerator](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGlyphGenerator creates a new NSGlyphGenerator instance.
func NewNSGlyphGenerator() NSGlyphGenerator {
	class := getNSGlyphGeneratorClass()
	rv := objc.Send[NSGlyphGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates glyphs for the specified glyph storage object ([NSLayoutManager]
// by default).
//
// # Discussion
//
// Generates glyphs for the glyph storage object specified by `glyphStorage`,
// beginning with the character at `charIndex` and continuing for `nChars`
// characters. The `glyphIndex` specifies the index of the first glyph
// generated.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator/generateGlyphs(for:desiredNumberOfCharacters:glyphIndex:characterIndex:)
func (g NSGlyphGenerator) GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage NSGlyphStorage, nChars uint, glyphIndex unsafe.Pointer, charIndex unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), glyphStorage, nChars, glyphIndex, charIndex)
}

// Returns a shared instance of [NSGlyphGenerator].
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator/shared
func (_NSGlyphGeneratorClass NSGlyphGeneratorClass) SharedGlyphGenerator() NSGlyphGenerator {
	rv := objc.Send[objc.ID](objc.ID(_NSGlyphGeneratorClass.class), objc.Sel("sharedGlyphGenerator"))
	return NSGlyphGeneratorFromID(objc.ID(rv))
}
