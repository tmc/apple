// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFont] class.
var (
	_NSFontClass     NSFontClass
	_NSFontClassOnce sync.Once
)

func getNSFontClass() NSFontClass {
	_NSFontClassOnce.Do(func() {
		_NSFontClass = NSFontClass{class: objc.GetClass("NSFont")}
	})
	return _NSFontClass
}

// GetNSFontClass returns the class object for NSFont.
func GetNSFontClass() NSFontClass {
	return getNSFontClass()
}

type NSFontClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontClass) Alloc() NSFont {
	rv := objc.Send[NSFont](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The representation of a font in an app.
//
// # Overview
// 
// [NSFont] objects represent fonts to an app, providing access to
// characteristics of the font and assistance in laying out glyphs relative to
// one another. Font objects are also used to establish the current font for
// drawing text directly into a graphics context, using the [NSFont.Set] method.
// 
// You don’t create [NSFont] objects using the `alloc` and `init` methods.
// Instead, you use either [NSFont.FontWithDescriptorSize] or [NSFont.FontWithNameSize] to
// look up an available font and alter its size or matrix to your needs. These
// methods check for an existing font object with the specified
// characteristics, returning it if there is one. Otherwise, they look up the
// font data requested and create the appropriate object. [NSFont] also
// defines a number of methods for getting standard system fonts, such as
// [NSFont.SystemFontOfSize], [NSFont.UserFontOfSize], and [NSFont.MessageFontOfSize]. To request
// the default size for these standard fonts, pass a negative number or `0` as
// the font size. See [macOS Human Interface Guidelines] for more information
// about system fonts.
//
// [macOS Human Interface Guidelines]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/OSXHIGuidelines/index.html#//apple_ref/doc/uid/20000957
//
// # Using a Font to Draw
//
//   - [NSFont.Set]: Sets this font as the font for the current graphics context.
//   - [NSFont.SetInContext]: Sets this font as the font for the specified graphics context.
//
// # Getting Font Metrics and Information
//
//   - [NSFont.PointSize]: The point size of the font.
//   - [NSFont.CoveredCharacterSet]: The character set containing all of the nominal characters that the font can render.
//   - [NSFont.FontDescriptor]: The font descriptor object for the font.
//   - [NSFont.FixedPitch]: A Boolean value indicating whether all glyphs in the font have the same advancement.
//   - [NSFont.MostCompatibleStringEncoding]: The string encoding that works best with the font.
//
// # Getting Information About Glyphs
//
//   - [NSFont.NumberOfGlyphs]: The number of glyphs in the font.
//   - [NSFont.NSControlGlyph]: The reserved code for a control glyph.
//   - [NSFont.SetNSControlGlyph]
//   - [NSFont.NSNullGlyph]: The reserved code for a null glyph.
//   - [NSFont.SetNSNullGlyph]
//
// # Getting Font Names
//
//   - [NSFont.DisplayName]: The name of the font, including family and face names, to use when displaying the font information to the user.
//   - [NSFont.FamilyName]: The family name of the font—for example, “Times” or “Helvetica.”
//   - [NSFont.FontName]: The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// # Vertical Fonts
//
//   - [NSFont.Vertical]: A Boolean value indicating whether the font is a vertical font.
//   - [NSFont.VerticalFont]: A vertical version of the font.
//
// # Instance Properties
//
//   - [NSFont.PrinterFont]: The scalable PostScript font corresponding to current font.
//   - [NSFont.RenderingMode]: The rendering mode of the font.
//   - [NSFont.ScreenFont]: The bitmapped screen font for the current font.
//
// # Instance Methods
//
//   - [NSFont.AdvancementForGlyph]: Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
//   - [NSFont.BoundingRectForGlyph]: Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
//   - [NSFont.GetAdvancementsForGlyphsCount]: Returns an array of the advancements for the specified glyphs rendered by the receiver.
//   - [NSFont.GetAdvancementsForPackedGlyphsLength]: Returns an array of the advancements for the specified packed glyphs and rendered by the receiver.
//   - [NSFont.GetBoundingRectsForGlyphsCount]: Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
//   - [NSFont.GlyphWithName]: Returns the named encoded glyph, or –1 if the receiver contains no such glyph.
//   - [NSFont.ScreenFontWithRenderingMode]: Returns a bitmapped screen font, when sent to a font object representing a scalable PostScript font, with the specified rendering mode, matching the receiver in typeface and matrix (or size), or `nil` if such a font can’t be found.
//   - [NSFont.FontWithSize]
//
// See: https://developer.apple.com/documentation/AppKit/NSFont
type NSFont struct {
	objectivec.Object
}

// NSFontFromID constructs a [NSFont] from an objc.ID.
//
// The representation of a font in an app.
func NSFontFromID(id objc.ID) NSFont {
	return NSFont{objectivec.Object{ID: id}}
}
// NOTE: NSFont adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFont] class.
//
// # Using a Font to Draw
//
//   - [INSFont.Set]: Sets this font as the font for the current graphics context.
//   - [INSFont.SetInContext]: Sets this font as the font for the specified graphics context.
//
// # Getting Font Metrics and Information
//
//   - [INSFont.PointSize]: The point size of the font.
//   - [INSFont.CoveredCharacterSet]: The character set containing all of the nominal characters that the font can render.
//   - [INSFont.FontDescriptor]: The font descriptor object for the font.
//   - [INSFont.FixedPitch]: A Boolean value indicating whether all glyphs in the font have the same advancement.
//   - [INSFont.MostCompatibleStringEncoding]: The string encoding that works best with the font.
//
// # Getting Information About Glyphs
//
//   - [INSFont.NumberOfGlyphs]: The number of glyphs in the font.
//   - [INSFont.NSControlGlyph]: The reserved code for a control glyph.
//   - [INSFont.SetNSControlGlyph]
//   - [INSFont.NSNullGlyph]: The reserved code for a null glyph.
//   - [INSFont.SetNSNullGlyph]
//
// # Getting Font Names
//
//   - [INSFont.DisplayName]: The name of the font, including family and face names, to use when displaying the font information to the user.
//   - [INSFont.FamilyName]: The family name of the font—for example, “Times” or “Helvetica.”
//   - [INSFont.FontName]: The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// # Vertical Fonts
//
//   - [INSFont.Vertical]: A Boolean value indicating whether the font is a vertical font.
//   - [INSFont.VerticalFont]: A vertical version of the font.
//
// # Instance Properties
//
//   - [INSFont.PrinterFont]: The scalable PostScript font corresponding to current font.
//   - [INSFont.RenderingMode]: The rendering mode of the font.
//   - [INSFont.ScreenFont]: The bitmapped screen font for the current font.
//
// # Instance Methods
//
//   - [INSFont.AdvancementForGlyph]: Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
//   - [INSFont.BoundingRectForGlyph]: Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
//   - [INSFont.GetAdvancementsForGlyphsCount]: Returns an array of the advancements for the specified glyphs rendered by the receiver.
//   - [INSFont.GetAdvancementsForPackedGlyphsLength]: Returns an array of the advancements for the specified packed glyphs and rendered by the receiver.
//   - [INSFont.GetBoundingRectsForGlyphsCount]: Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
//   - [INSFont.GlyphWithName]: Returns the named encoded glyph, or –1 if the receiver contains no such glyph.
//   - [INSFont.ScreenFontWithRenderingMode]: Returns a bitmapped screen font, when sent to a font object representing a scalable PostScript font, with the specified rendering mode, matching the receiver in typeface and matrix (or size), or `nil` if such a font can’t be found.
//   - [INSFont.FontWithSize]
//
// See: https://developer.apple.com/documentation/AppKit/NSFont
type INSFont interface {
	objectivec.IObject

	// Topic: Using a Font to Draw

	// Sets this font as the font for the current graphics context.
	Set()
	// Sets this font as the font for the specified graphics context.
	SetInContext(graphicsContext INSGraphicsContext)

	// Topic: Getting Font Metrics and Information

	// The point size of the font.
	PointSize() float64
	// The character set containing all of the nominal characters that the font can render.
	CoveredCharacterSet() foundation.NSCharacterSet
	// The font descriptor object for the font.
	FontDescriptor() INSFontDescriptor
	// A Boolean value indicating whether all glyphs in the font have the same advancement.
	FixedPitch() bool
	// The string encoding that works best with the font.
	MostCompatibleStringEncoding() uint

	// Topic: Getting Information About Glyphs

	// The number of glyphs in the font.
	NumberOfGlyphs() uint
	// The reserved code for a control glyph.
	NSControlGlyph() int
	SetNSControlGlyph(value int)
	// The reserved code for a null glyph.
	NSNullGlyph() int
	SetNSNullGlyph(value int)

	// Topic: Getting Font Names

	// The name of the font, including family and face names, to use when displaying the font information to the user.
	DisplayName() string
	// The family name of the font—for example, “Times” or “Helvetica.”
	FamilyName() string
	// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
	FontName() string

	// Topic: Vertical Fonts

	// A Boolean value indicating whether the font is a vertical font.
	Vertical() bool
	// A vertical version of the font.
	VerticalFont() NSFont

	// Topic: Instance Properties

	// The scalable PostScript font corresponding to current font.
	PrinterFont() NSFont
	// The rendering mode of the font.
	RenderingMode() NSFontRenderingMode
	// The bitmapped screen font for the current font.
	ScreenFont() NSFont

	// Topic: Instance Methods

	// Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
	AdvancementForGlyph(glyph NSGlyph) corefoundation.CGSize
	// Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
	BoundingRectForGlyph(glyph NSGlyph) corefoundation.CGRect
	// Returns an array of the advancements for the specified glyphs rendered by the receiver.
	GetAdvancementsForGlyphsCount(advancements foundation.NSSize, glyphs []NSGlyph, glyphCount uint)
	// Returns an array of the advancements for the specified packed glyphs and rendered by the receiver.
	GetAdvancementsForPackedGlyphsLength(advancements foundation.NSSize, packedGlyphs unsafe.Pointer, length uint)
	// Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
	GetBoundingRectsForGlyphsCount(bounds foundation.NSRect, glyphs []NSGlyph, glyphCount uint)
	// Returns the named encoded glyph, or –1 if the receiver contains no such glyph.
	GlyphWithName(name string) NSGlyph
	// Returns a bitmapped screen font, when sent to a font object representing a scalable PostScript font, with the specified rendering mode, matching the receiver in typeface and matrix (or size), or `nil` if such a font can’t be found.
	ScreenFontWithRenderingMode(renderingMode NSFontRenderingMode) NSFont
	FontWithSize(fontSize float64) NSFont

	// The top y-coordinate, offset from the baseline, of the font’s longest ascender.
	Ascender() float64
	SetAscender(value float64)
	// The font’s bounding rectangle, scaled to the font’s size.
	BoundingRectForFont() corefoundation.CGRect
	SetBoundingRectForFont(value corefoundation.CGRect)
	// The cap height of the font.
	CapHeight() float64
	SetCapHeight(value float64)
	// The bottom y-coordinate, offset from the baseline, of the font’s longest descender.
	Descender() float64
	SetDescender(value float64)
	// The number of degrees that the font is slanted counterclockwise from the vertical.
	ItalicAngle() float64
	SetItalicAngle(value float64)
	// The leading value of the font.
	Leading() float64
	SetLeading(value float64)
	// The transformation matrix associated with the font.
	Matrix() float64
	SetMatrix(value float64)
	// The maximum advance of any of the font’s glyphs.
	MaximumAdvancement() corefoundation.CGSize
	SetMaximumAdvancement(value corefoundation.CGSize)
	// The current transformation matrix of the font.
	TextTransform() objectivec.IObject
	SetTextTransform(value objectivec.IObject)
	// The baseline offset to use when drawing underlines with the font.
	UnderlinePosition() float64
	SetUnderlinePosition(value float64)
	// The thickness to use when drawing underlines with the font.
	UnderlineThickness() float64
	SetUnderlineThickness(value float64)
	// The x-height of the font.
	XHeight() float64
	SetXHeight(value float64)
	// Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
	AdvancementForCGGlyph(glyph coregraphics.CGFontIndex) corefoundation.CGSize
	// Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
	BoundingRectForCGGlyph(glyph coregraphics.CGFontIndex) corefoundation.CGRect
	// Returns an array of the advancements for the specified glyphs rendered by the receiver.
	GetAdvancementsForCGGlyphsCount(advancements foundation.NSSize, glyphs []coregraphics.CGFontIndex, glyphCount uint)
	// Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
	GetBoundingRectsForCGGlyphsCount(bounds foundation.NSRect, glyphs []coregraphics.CGFontIndex, glyphCount uint)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFont) Init() NSFont {
	rv := objc.Send[NSFont](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFont) Autorelease() NSFont {
	rv := objc.Send[NSFont](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFont creates a new NSFont instance.
func NewNSFont() NSFont {
	class := getNSFontClass()
	rv := objc.Send[NSFont](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a font object for the specified font descriptor and font size.
//
// fontDescriptor: A font descriptor object.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object for the specified descriptor and size.
//
// # Discussion
// 
// In most cases, you can simply use [FontWithNameSize] to create standard
// scaled fonts.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:size:)
func NewFontWithDescriptorSize(fontDescriptor INSFontDescriptor, fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(getNSFontClass().class), objc.Sel("fontWithDescriptor:size:"), fontDescriptor, fontSize)
	return NSFontFromID(rv)
}

// Returns a font object for the specified font descriptor and text transform.
//
// fontDescriptor: The font descriptor object describing the font to return.
//
// textTransform: An affine transformation applied to the font.
//
// # Return Value
// 
// A font object for the specified name and transform.
//
// # Discussion
// 
// In most cases, you can simply use [FontWithNameSize] to create standard
// scaled fonts. If `textTransform` is non-nil, it has precedence over
// [NSFontMatrixAttribute] in `fontDescriptor`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:textTransform:)
func NewFontWithDescriptorTextTransform(fontDescriptor INSFontDescriptor, textTransform foundation.NSAffineTransform) NSFont {
	rv := objc.Send[objc.ID](objc.ID(getNSFontClass().class), objc.Sel("fontWithDescriptor:textTransform:"), fontDescriptor, textTransform)
	return NSFontFromID(rv)
}

// Returns a font object for the specified font name and matrix.
//
// fontName: The fully specified family-face name of the font.
//
// fontMatrix: A transformation matrix applied to the font.
//
// # Return Value
// 
// A font object for the specified name and transformation matrix.
//
// # Discussion
// 
// The `fontName` is a fully specified family-face name, such as
// Helvetica-BoldOblique or Times-Roman (not a name as shown in the Font
// Panel). The `fontMatrix` is a standard 6-element transformation matrix as
// used in the PostScript language, specifically with the `makefont` operator.
// In most cases, you can simply use [FontWithNameSize] to create standard
// scaled fonts.
// 
// You can use the defined value [NSFontIdentityMatrix] for [1 0 0 1 0 0].
// Fonts created with a matrix other than [NSFontIdentityMatrix] don’t
// automatically flip themselves in flipped views.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/init(name:matrix:)
func NewFontWithNameMatrix(fontName string, fontMatrix unsafe.Pointer) NSFont {
	rv := objc.Send[objc.ID](objc.ID(getNSFontClass().class), objc.Sel("fontWithName:matrix:"), objc.String(fontName), fontMatrix)
	return NSFontFromID(rv)
}

// Creates a font object for the specified font name and font size.
//
// fontName: The fully specified family-face name of the font.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object for the specified name and size.
//
// # Discussion
// 
// The value of the `fontName` parameter is a fully specified family-face
// name, preferably the PostScript name, such as Helvetica-BoldOblique or
// Times-Roman. (The Font Book app displays PostScript names of fonts in the
// Font Info panel.)
// 
// Specifying `fontSize` is equivalent to using a font matrix of [`fontSize` 0
// 0 `fontSize` 0 0] with [FontWithDescriptorSize]. If you use a `fontSize` of
// 0.0, this method uses the default User Font size.
// 
// Fonts created with this method automatically flip themselves in flipped
// views. This method is the preferred means for creating fonts.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/init(name:size:)
func NewFontWithNameSize(fontName string, fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(getNSFontClass().class), objc.Sel("fontWithName:size:"), objc.String(fontName), fontSize)
	return NSFontFromID(rv)
}

// Sets this font as the font for the current graphics context.
//
// # Discussion
// 
// This method sets the font for the graphics system but does not affect the
// higher-level settings of the Cocoa text system, which are controlled by
// text attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/set()
func (f NSFont) Set() {
	objc.Send[objc.ID](f.ID, objc.Sel("set"))
}
// Sets this font as the font for the specified graphics context.
//
// graphicsContext: The graphics context for which the font is set.
//
// # Discussion
// 
// This method sets the font for the graphics system but does not affect the
// higher-level settings of the Cocoa text system, which are controlled by
// text attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/set(in:)
func (f NSFont) SetInContext(graphicsContext INSGraphicsContext) {
	objc.Send[objc.ID](f.ID, objc.Sel("setInContext:"), graphicsContext)
}
// Returns the nominal spacing for the given glyph—the distance the current
// point moves after showing the glyph—accounting for the receiver’s size.
//
// glyph: The glyph whose advancement is returned.
//
// # Return Value
// 
// The advancement spacing in points.
//
// # Discussion
// 
// This spacing is given according to the glyph’s movement direction, which
// is either strictly horizontal or strictly vertical.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/advancement(forGlyph:)
func (f NSFont) AdvancementForGlyph(glyph NSGlyph) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](f.ID, objc.Sel("advancementForGlyph:"), glyph)
	return corefoundation.CGSize(rv)
}
// Returns the bounding rectangle for the specified glyph, scaled to the
// receiver’s size.
//
// # Discussion
// 
// Japanese fonts encoded with the scheme “EUC12-NJE-CFEncoding” do not
// have individual metrics or bounding boxes available for the glyphs above
// 127. For those glyphs, this method returns the bounding rectangle for the
// font instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/boundingRect(forGlyph:)
func (f NSFont) BoundingRectForGlyph(glyph NSGlyph) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("boundingRectForGlyph:"), glyph)
	return corefoundation.CGRect(rv)
}
// Returns an array of the advancements for the specified glyphs rendered by
// the receiver.
//
// # Discussion
// 
// Returns in `advancements` an array of the advancements for the glyphs
// specified by `glyphs` and rendered by the receiver. The `glyphCount` must
// specify the count of glyphs passed in `glyphs`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forGlyphs:count:)
func (f NSFont) GetAdvancementsForGlyphsCount(advancements foundation.NSSize, glyphs []NSGlyph, glyphCount uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("getAdvancements:forGlyphs:count:"), advancements, objc.CArray(glyphs), glyphCount)
}
// Returns an array of the advancements for the specified packed glyphs and
// rendered by the receiver.
//
// # Discussion
// 
// Returns in `advancements` an array of the advancements for the packed
// glyphs specified by `packedGlyphs` and rendered by the receiver. The
// `glyphCount` must specify the count of glyphs passed in `packedGlyphs`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forPackedGlyphs:length:)
func (f NSFont) GetAdvancementsForPackedGlyphsLength(advancements foundation.NSSize, packedGlyphs unsafe.Pointer, length uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("getAdvancements:forPackedGlyphs:length:"), advancements, packedGlyphs, length)
}
// Returns an array of the bounding rectangles for the specified glyphs
// rendered by the receiver.
//
// # Discussion
// 
// Returns in `bounds` an array of the bounding rectangles for the glyphs
// specified by `glyphs` and rendered by the receiver. The `glyphCount` must
// specify the count of glyphs passed in `glyphs`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/getBoundingRects(_:forGlyphs:count:)
func (f NSFont) GetBoundingRectsForGlyphsCount(bounds foundation.NSRect, glyphs []NSGlyph, glyphCount uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("getBoundingRects:forGlyphs:count:"), bounds, objc.CArray(glyphs), glyphCount)
}
// Returns the named encoded glyph, or –1 if the receiver contains no such
// glyph.
//
// name: The name of the glyph.
//
// # Return Value
// 
// The named encoded glyph.
//
// # Discussion
// 
// Returns –1 if the glyph named `glyphName` isn’t encoded.
// 
// Glyph names in fonts do not always accurately identify the glyph. The
// layout manager, an instance of [NSLayoutManager], finds the correspondence
// between characters and glyphs. See [Text Layout Programming Guide] and
// [NSLayoutManager] for more information.
//
// [Text Layout Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextLayout/TextLayout.html#//apple_ref/doc/uid/10000158i
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/glyph(withName:)
func (f NSFont) GlyphWithName(name string) NSGlyph {
	rv := objc.Send[NSGlyph](f.ID, objc.Sel("glyphWithName:"), objc.String(name))
	return NSGlyph(rv)
}
// Returns a bitmapped screen font, when sent to a font object representing a
// scalable PostScript font, with the specified rendering mode, matching the
// receiver in typeface and matrix (or size), or `nil` if such a font can’t
// be found.
//
// # Discussion
// 
// For valid rendering modes, see [NSFontRenderingMode].
// 
// Screen fonts are for direct use with the window server only. Never use them
// with Application Kit objects, such as in `` methods. Internally, the
// Application Kit automatically uses the corresponding screen font for a font
// object as long as the view is not rotated or scaled.
//
// [NSFontRenderingMode]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/screenFont(with:)
func (f NSFont) ScreenFontWithRenderingMode(renderingMode NSFontRenderingMode) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("screenFontWithRenderingMode:"), renderingMode)
	return NSFontFromID(rv)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/withSize(_:)
func (f NSFont) FontWithSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontWithSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the nominal spacing for the given glyph—the distance the current
// point moves after showing the glyph—accounting for the receiver’s size.
//
// glyph: The glyph whose advancement is returned.
//
// # Return Value
// 
// The advancement spacing in points.
//
// # Discussion
// 
// The spacing is given according to the glyph’s movement direction, which
// is either strictly horizontal or strictly vertical.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/advancement(forCGGlyph:)
func (f NSFont) AdvancementForCGGlyph(glyph coregraphics.CGFontIndex) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](f.ID, objc.Sel("advancementForCGGlyph:"), glyph)
	return corefoundation.CGSize(rv)
}
// Returns the bounding rectangle for the specified glyph, scaled to the
// receiver’s size.
//
// # Discussion
// 
// Japanese fonts encoded with the scheme “EUC12-NJE-CFEncoding” do not
// have individual metrics or bounding boxes available for the glyphs above
// 127. For those glyphs, this method returns the bounding rectangle for the
// font instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/boundingRect(forCGGlyph:)
func (f NSFont) BoundingRectForCGGlyph(glyph coregraphics.CGFontIndex) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("boundingRectForCGGlyph:"), glyph)
	return corefoundation.CGRect(rv)
}
// Returns an array of the advancements for the specified glyphs rendered by
// the receiver.
//
// # Discussion
// 
// Returns in `advancements` an array of the advancements for the glyphs
// specified by `glyphs` and rendered by the receiver. The `glyphCount` value
// must specify the count of glyphs passed in `glyphs`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forCGGlyphs:count:)
func (f NSFont) GetAdvancementsForCGGlyphsCount(advancements foundation.NSSize, glyphs []coregraphics.CGFontIndex, glyphCount uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("getAdvancements:forCGGlyphs:count:"), advancements, objc.CArray(glyphs), glyphCount)
}
// Returns an array of the bounding rectangles for the specified glyphs
// rendered by the receiver.
//
// # Discussion
// 
// Returns in `bounds` an array of the bounding rectangles for the glyphs
// specified by `glyphs` and rendered by the receiver. The `glyphCount` value
// must specify the count of glyphs passed in `glyphs`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/getBoundingRects(_:forCGGlyphs:count:)
func (f NSFont) GetBoundingRectsForCGGlyphsCount(bounds foundation.NSRect, glyphs []coregraphics.CGFontIndex, glyphCount uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("getBoundingRects:forCGGlyphs:count:"), bounds, objc.CArray(glyphs), glyphCount)
}
func (f NSFont) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the font used by default for documents and other text under the
// user’s control (that is, text whose font the user can normally change),
// in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the user font at the default size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/userFont(ofSize:)
func (_NSFontClass NSFontClass) UserFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("userFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used by default for documents and other text under the
// user’s control (that is, text whose font the user can normally change),
// when that font should be fixed-pitch, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the fixed-pitch font at the default
// size.
// 
// The system does not guarantee that all the glyphs in a fixed-pitch font are
// the same width. For example, certain Japanese fonts are dual-pitch, and
// other fonts may have nonspacing marks that can affect the display of other
// glyphs.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/userFixedPitchFont(ofSize:)
func (_NSFontClass NSFontClass) UserFixedPitchFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("userFixedPitchFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font associated with the text style.
//
// style: The text style for which to return a font. See [NSFontTextStyle] for
// available values.
//
// options: A dictionary you use to further configure the returned font. See
// [NSFontTextStyleOptionKey] for a list of valid keys. Pass an empty
// dictionary to use the default configuration.
//
// # Return Value
// 
// The font associated with the text style.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/preferredFont(forTextStyle:options:)
func (_NSFontClass NSFontClass) PreferredFontForTextStyleOptions(style NSFontTextStyle, options foundation.INSDictionary) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("preferredFontForTextStyle:options:"), objc.String(string(style)), options)
	return NSFontFromID(rv)
}
// Returns the standard system font with the specified size.
//
// fontSize: The desired font size specified in points. If you specify `0.0` or a
// negative number for this parameter, the method returns the system font at
// the default size.
//
// # Return Value
// 
// A font object containing the system font at the specified size.
//
// # Discussion
// 
// Use the returned font for standard interface items, including button
// labels, menu items, and so on that use the default font appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:)
func (_NSFontClass NSFontClass) SystemFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("systemFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the standard system font with the specified size and weight.
//
// fontSize: The desired font size specified in points. If you specify `0.0` or a
// negative number for this parameter, the method returns the system font at
// the default size.
//
// weight: The desired weight of font lines, specified as one of the constants in
// [NSFontWeight].
//
// # Return Value
// 
// A font object containing the system font at the specified size and weight.
//
// # Discussion
// 
// Use the returned font for standard interface items, including button
// labels, menu items, and so on that require a specific font style
// information.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:)
func (_NSFontClass NSFontClass) SystemFontOfSizeWeight(fontSize float64, weight NSFontWeight) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("systemFontOfSize:weight:"), fontSize, weight)
	return NSFontFromID(rv)
}
// Returns the standard system font in boldface type with the specified size.
//
// fontSize: The desired font size specified in points. If you specify `0.0` or a
// negative number for this parameter, the method returns the system font at
// the default size.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the boldface system font at the
// default size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/boldSystemFont(ofSize:)
func (_NSFontClass NSFontClass) BoldSystemFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("boldSystemFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns a monospace version of the system font with the specified size and
// weight.
//
// fontSize: The desired font size specified in points. If you specify `0.0` or a
// negative number for this parameter, the method returns the system font at
// the default size.
//
// weight: The desired weight of font lines, specified as one of the constants in
// [NSFontWeight].
//
// # Return Value
// 
// A font object containing a monospace version of the system font at the
// specified size and weight.
//
// # Discussion
// 
// Use the returned font for interface items that require monospaced glyphs.
// The returned font includes monospaced glyphs for the Latin characters and
// the symbols commonly found in source code. Glyphs for other symbols are
// usually wider or narrower than the monospaced characters. To ensure the
// font uses fixed spacing for all characters, apply the [fixedAdvance]
// attribute to the any strings you render.
//
// [fixedAdvance]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/fixedAdvance
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/monospacedSystemFont(ofSize:weight:)
func (_NSFontClass NSFontClass) MonospacedSystemFontOfSizeWeight(fontSize float64, weight NSFontWeight) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("monospacedSystemFontOfSize:weight:"), fontSize, weight)
	return NSFontFromID(rv)
}
// Returns a version of the standard system font that contains monospaced
// digit glyphs.
//
// fontSize: The desired font size specified in points. If you specify `0.0` or a
// negative number for this parameter, the method returns the system font at
// the default size.
//
// weight: The desired weight of font lines, specified as one of the constants in
// [NSFontWeight].
//
// # Return Value
// 
// A font object containing the system font with monospace digits at the
// specified size and weight.
//
// # Discussion
// 
// The font returned by this method has monospaced digit glyphs. Glyphs for
// other characters and symbols may be wider or narrower than the monospaced
// characters. To ensure the font uses fixed spacing for all characters, apply
// the [fixedAdvance] attribute to the any strings you render.
//
// [fixedAdvance]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/fixedAdvance
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/monospacedDigitSystemFont(ofSize:weight:)
func (_NSFontClass NSFontClass) MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight NSFontWeight) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("monospacedDigitSystemFontOfSize:weight:"), fontSize, weight)
	return NSFontFromID(rv)
}
// Returns the font used for standard interface labels in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size. If `fontSize` is 0 or negative,
// returns the label font with the default size.
//
// # Discussion
// 
// The label font (Lucida Grande Regular 10 point) is used for the labels on
// toolbar buttons and to label tick marks on full-size sliders. See The macOS
// Environment in [macOS Human Interface Guidelines] for more information
// about system fonts.
//
// [macOS Human Interface Guidelines]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/OSXHIGuidelines/index.html#//apple_ref/doc/uid/20000957
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/labelFont(ofSize:)
func (_NSFontClass NSFontClass) LabelFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("labelFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for standard interface items, such as button labels,
// menu items, and so on, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns this font at the default size. This
// method is equivalent to [SystemFontOfSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/messageFont(ofSize:)
func (_NSFontClass NSFontClass) MessageFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("messageFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for menu bar items, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the menu bar font with the default
// size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/menuBarFont(ofSize:)
func (_NSFontClass NSFontClass) MenuBarFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("menuBarFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for menu items, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the menu items font with the
// default size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/menuFont(ofSize:)
func (_NSFontClass NSFontClass) MenuFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("menuFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for the content of controls in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// For example, in a table, the user’s input uses the control content font,
// and the table’s header uses another font. If `fontSize` is 0 or negative,
// returns the control content font at the default size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/controlContentFont(ofSize:)
func (_NSFontClass NSFontClass) ControlContentFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("controlContentFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for window title bars, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the title bar font at the default
// size. This method is equivalent to [BoldSystemFontOfSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/titleBarFont(ofSize:)
func (_NSFontClass NSFontClass) TitleBarFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("titleBarFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for palette window title bars, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the palette title font at the
// default size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/paletteFont(ofSize:)
func (_NSFontClass NSFontClass) PaletteFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("paletteFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font used for tool tips labels, in the specified size.
//
// fontSize: The size in points to which the font is scaled.
//
// # Return Value
// 
// A font object of the specified size.
//
// # Discussion
// 
// If `fontSize` is 0 or negative, returns the tool tips font at the default
// size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/toolTipsFont(ofSize:)
func (_NSFontClass NSFontClass) ToolTipsFontOfSize(fontSize float64) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("toolTipsFontOfSize:"), fontSize)
	return NSFontFromID(rv)
}
// Returns the font size used for the specified control size.
//
// controlSize: The control size constant.
//
// # Return Value
// 
// The font size in points for the specified control size.
//
// # Discussion
// 
// If `controlSize` does not correspond to a valid [NSControlSize], returns
// the size of the standard system font.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/systemFontSize(for:)
func (_NSFontClass NSFontClass) SystemFontSizeForControlSize(controlSize NSControlSize) float64 {
	rv := objc.Send[float64](objc.ID(_NSFontClass.class), objc.Sel("systemFontSizeForControlSize:"), controlSize)
	return rv
}
// Sets the font used by default for documents and other text under the
// user’s control to the specified font.
//
// # Discussion
// 
// Specifying `aFont` as `nil` causes the default to be removed from the
// application domain.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/setUser(_:)
func (_NSFontClass NSFontClass) SetUserFont(font NSFont) {
	objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("setUserFont:"), font)
}
// Sets the font used by default for documents and other text under the
// user’s control, when that font should be fixed-pitch, to the specified
// font.
//
// # Discussion
// 
// Specifying `aFont` as `nil` causes the default to be removed from the
// application domain.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/setUserFixedPitch(_:)
func (_NSFontClass NSFontClass) SetUserFixedPitchFont(font NSFont) {
	objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("setUserFixedPitchFont:"), font)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:width:)
func (_NSFontClass NSFontClass) SystemFontOfSizeWeightWidth(fontSize float64, weight NSFontWeight, width NSFontWidth) NSFont {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("systemFontOfSize:weight:width:"), fontSize, weight, width)
	return NSFontFromID(rv)
}

// The point size of the font.
//
// # Discussion
// 
// If the font has a nonstandard matrix, the point size is the effective
// vertical point size.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/pointSize
func (f NSFont) PointSize() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("pointSize"))
	return rv
}
// The character set containing all of the nominal characters that the font
// can render.
//
// # Discussion
// 
// The nominal character set is all of the entries in the font’s `cmap`
// table.
// 
// The number of glyphs supported by a given font is often larger than the
// number of characters contained in the character set returned by this
// method. This is because characters and glyphs have a many-to-many mapping,
// rather than a strict one-to-one correspondence. In some cases a character
// may be represented by multiple glyphs, such as an “é” which may be an
// “e” glyph combined with an acute accent glyph “´”. In other cases,
// a single glyph may represent multiple characters, as in the case of a
// ligature, or joined letter. See [Typographical Concepts] in [Cocoa Text
// Architecture Guide] for more information.
//
// [Cocoa Text Architecture Guide]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Introduction/Introduction.html#//apple_ref/doc/uid/TP40009459
// [Typographical Concepts]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html#//apple_ref/doc/uid/TP40009459-CH6
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/coveredCharacterSet
func (f NSFont) CoveredCharacterSet() foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("coveredCharacterSet"))
	return foundation.NSCharacterSetFromID(objc.ID(rv))
}
// The font descriptor object for the font.
//
// # Discussion
// 
// The font descriptor contains a mutable dictionary of optional attributes
// for creating an [NSFont] object. For more information about font
// descriptors, see [NSFontDescriptor].
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/fontDescriptor
func (f NSFont) FontDescriptor() INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptor"))
	return NSFontDescriptorFromID(objc.ID(rv))
}
// A Boolean value indicating whether all glyphs in the font have the same
// advancement.
//
// # Discussion
// 
// The value of this property is [true] when all glyphs have the same
// advancement or [false] when they do not. Some Japanese fonts encoded with
// the scheme “EUC12-NJE-CFEncoding” return that they have the same
// advancement, but actually encode glyphs with one of two advancements, for
// historical compatibility. You may need to handle such fonts specially for
// some applications.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/isFixedPitch
func (f NSFont) FixedPitch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isFixedPitch"))
	return rv
}
// The string encoding that works best with the font.
//
// # Discussion
// 
// The string encoding in this property is the encoding with the fewest
// unmatched characters and glyphs in the font. If this value is
// [NSASCIIStringEncoding], the font could not determine the correct encoding;
// you should assume the font can render only ASCII characters. The font uses
// heuristically well-known font encodings to determine the value of this
// property, so for nonstandard encodings the property may not contain the
// optimal string encoding.
// 
// You can use the [data(using:)] or [data(using:allowLossyConversion:)]
// methods of [NSString] to convert strings to this encoding.
//
// [NSASCIIStringEncoding]: https://developer.apple.com/documentation/Foundation/NSASCIIStringEncoding
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [data(using:)]: https://developer.apple.com/documentation/Foundation/NSString/data(using:)
// [data(using:allowLossyConversion:)]: https://developer.apple.com/documentation/Foundation/NSString/data(using:allowLossyConversion:)
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/mostCompatibleStringEncoding
func (f NSFont) MostCompatibleStringEncoding() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("mostCompatibleStringEncoding"))
	return rv
}
// The number of glyphs in the font.
//
// # Discussion
// 
// Glyphs are numbered starting at 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/numberOfGlyphs
func (f NSFont) NumberOfGlyphs() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("numberOfGlyphs"))
	return rv
}
// The reserved code for a control glyph.
//
// See: https://developer.apple.com/documentation/appkit/nscontrolglyph
func (f NSFont) NSControlGlyph() int {
	rv := objc.Send[int](f.ID, objc.Sel("NSControlGlyph"))
	return rv
}
func (f NSFont) SetNSControlGlyph(value int) {
	objc.Send[struct{}](f.ID, objc.Sel("setNSControlGlyph:"), value)
}
// The reserved code for a null glyph.
//
// See: https://developer.apple.com/documentation/appkit/nsnullglyph
func (f NSFont) NSNullGlyph() int {
	rv := objc.Send[int](f.ID, objc.Sel("NSNullGlyph"))
	return rv
}
func (f NSFont) SetNSNullGlyph(value int) {
	objc.Send[struct{}](f.ID, objc.Sel("setNSNullGlyph:"), value)
}
// The name of the font, including family and face names, to use when
// displaying the font information to the user.
//
// # Discussion
// 
// The font’s display name is typically localized for the user’s language.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/displayName
func (f NSFont) DisplayName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
// The family name of the font—for example, “Times” or “Helvetica.”
//
// # Discussion
// 
// This name is the one that [NSFontManager] uses and may differ slightly from
// the AFM name.
// 
// The value in this property is intended for an application’s internal
// usage and not for display. To get a name that you can display to the user,
// use the [DisplayName] property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/familyName
func (f NSFont) FamilyName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("familyName"))
	return foundation.NSStringFromID(rv).String()
}
// The full name of the font, as used in PostScript language code—for
// example, “Times-Roman” or “Helvetica-Oblique.”
//
// # Discussion
// 
// The value in this property is intended for an application’s internal
// usage and not for display. To get a name that you can display to the user,
// use the [DisplayName] property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/fontName
func (f NSFont) FontName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontName"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value indicating whether the font is a vertical font.
//
// # Discussion
// 
// The value in this property is [true] for a vertical font or [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/isVertical
func (f NSFont) Vertical() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isVertical"))
	return rv
}
// A vertical version of the font.
//
// # Discussion
// 
// The value in this property is a vertical version of the font, if such a
// configuration is supported. If a vertical configuration is not supported,
// the value in the property is `self`.
// 
// A vertical font applies appropriate rotation to the text matrix in
// [SetInContext], returns vertical metrics, and enables the vertical glyph
// substitution feature by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/vertical-6ym79
func (f NSFont) VerticalFont() NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("verticalFont"))
	return NSFontFromID(objc.ID(rv))
}
// The scalable PostScript font corresponding to current font.
//
// # Discussion
// 
// For a font that already represents a scalable PostScript font, the value in
// this property is `self`. For a bitmapped screen font, the value is the
// corresponding scalable PostScript font.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/printer
func (f NSFont) PrinterFont() NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("printerFont"))
	return NSFontFromID(objc.ID(rv))
}
// The rendering mode of the font.
//
// # Discussion
// 
// For a list of valid rendering modes, see [NSFontRenderingMode].
//
// [NSFontRenderingMode]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/renderingMode
func (f NSFont) RenderingMode() NSFontRenderingMode {
	rv := objc.Send[NSFontRenderingMode](f.ID, objc.Sel("renderingMode"))
	return NSFontRenderingMode(rv)
}
// The bitmapped screen font for the current font.
//
// # Discussion
// 
// For a font object that represents a scalable PostScript font, this property
// contains a bitmapped screen font matching the receiver in typeface and
// matrix (or size), or `nil` if such a font cannot be found. For a bitmapped
// screen font, the value in this property is `nil`.
// 
// Screen fonts are for direct use with the window server only. Never use them
// with Application Kit objects, such as in `` methods. Internally, the
// Application Kit automatically uses the corresponding screen font for a font
// object as long as the view is not rotated or scaled.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/screen
func (f NSFont) ScreenFont() NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("screenFont"))
	return NSFontFromID(objc.ID(rv))
}
// The top y-coordinate, offset from the baseline, of the font’s longest
// ascender.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/ascender
func (f NSFont) Ascender() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("ascender"))
	return rv
}
func (f NSFont) SetAscender(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setAscender:"), value)
}
// The font’s bounding rectangle, scaled to the font’s size.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/boundingrectforfont
func (f NSFont) BoundingRectForFont() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("boundingRectForFont"))
	return corefoundation.CGRect(rv)
}
func (f NSFont) SetBoundingRectForFont(value corefoundation.CGRect) {
	objc.Send[struct{}](f.ID, objc.Sel("setBoundingRectForFont:"), value)
}
// The cap height of the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/capheight
func (f NSFont) CapHeight() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("capHeight"))
	return rv
}
func (f NSFont) SetCapHeight(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setCapHeight:"), value)
}
// The bottom y-coordinate, offset from the baseline, of the font’s longest
// descender.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/descender
func (f NSFont) Descender() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("descender"))
	return rv
}
func (f NSFont) SetDescender(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setDescender:"), value)
}
// The number of degrees that the font is slanted counterclockwise from the
// vertical.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/italicangle
func (f NSFont) ItalicAngle() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("italicAngle"))
	return rv
}
func (f NSFont) SetItalicAngle(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setItalicAngle:"), value)
}
// The leading value of the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/leading
func (f NSFont) Leading() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("leading"))
	return rv
}
func (f NSFont) SetLeading(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setLeading:"), value)
}
// The transformation matrix associated with the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/matrix
func (f NSFont) Matrix() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("matrix"))
	return rv
}
func (f NSFont) SetMatrix(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setMatrix:"), value)
}
// The maximum advance of any of the font’s glyphs.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/maximumadvancement
func (f NSFont) MaximumAdvancement() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](f.ID, objc.Sel("maximumAdvancement"))
	return corefoundation.CGSize(rv)
}
func (f NSFont) SetMaximumAdvancement(value corefoundation.CGSize) {
	objc.Send[struct{}](f.ID, objc.Sel("setMaximumAdvancement:"), value)
}
// The current transformation matrix of the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/texttransform
func (f NSFont) TextTransform() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("textTransform"))
	return objectivec.Object{ID: rv}
}
func (f NSFont) SetTextTransform(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setTextTransform:"), value)
}
// The baseline offset to use when drawing underlines with the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/underlineposition
func (f NSFont) UnderlinePosition() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("underlinePosition"))
	return rv
}
func (f NSFont) SetUnderlinePosition(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setUnderlinePosition:"), value)
}
// The thickness to use when drawing underlines with the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/underlinethickness
func (f NSFont) UnderlineThickness() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("underlineThickness"))
	return rv
}
func (f NSFont) SetUnderlineThickness(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setUnderlineThickness:"), value)
}
// The x-height of the font.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/xheight
func (f NSFont) XHeight() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("xHeight"))
	return rv
}
func (f NSFont) SetXHeight(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setXHeight:"), value)
}

// Returns the size of the standard system font.
//
// # Return Value
// 
// The standard system font size in points.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/systemFontSize
func (_NSFontClass NSFontClass) SystemFontSize() float64 {
	rv := objc.Send[float64](objc.ID(_NSFontClass.class), objc.Sel("systemFontSize"))
	return rv
}
// Returns the size of the standard small system font.
//
// # Return Value
// 
// The small system font size in points.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/smallSystemFontSize
func (_NSFontClass NSFontClass) SmallSystemFontSize() float64 {
	rv := objc.Send[float64](objc.ID(_NSFontClass.class), objc.Sel("smallSystemFontSize"))
	return rv
}
// Returns the size of the standard label font.
//
// # Return Value
// 
// The label font size in points.
// 
// # Discussion
// 
// The label font (Lucida Grande Regular 10 point) is used for the labels on
// toolbar buttons and to label tick marks on full-size sliders. See The macOS
// Environment in [macOS Human Interface Guidelines] for more information
// about system fonts.
//
// [macOS Human Interface Guidelines]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/OSXHIGuidelines/index.html#//apple_ref/doc/uid/20000957
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/labelFontSize
func (_NSFontClass NSFontClass) LabelFontSize() float64 {
	rv := objc.Send[float64](objc.ID(_NSFontClass.class), objc.Sel("labelFontSize"))
	return rv
}
// Posted after the threshold for antialiasing changes.
//
// See: https://developer.apple.com/documentation/appkit/nsfont/antialiasthresholdchangednotification
func (_NSFontClass NSFontClass) AntialiasThresholdChangedNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSFontClass.class), objc.Sel("NSAntialiasThresholdChangedNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

