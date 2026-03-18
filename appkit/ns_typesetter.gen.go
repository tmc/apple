// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTypesetter] class.
var (
	_NSTypesetterClass     NSTypesetterClass
	_NSTypesetterClassOnce sync.Once
)

func getNSTypesetterClass() NSTypesetterClass {
	_NSTypesetterClassOnce.Do(func() {
		_NSTypesetterClass = NSTypesetterClass{class: objc.GetClass("NSTypesetter")}
	})
	return _NSTypesetterClass
}

// GetNSTypesetterClass returns the class object for NSTypesetter.
func GetNSTypesetterClass() NSTypesetterClass {
	return getNSTypesetterClass()
}

type NSTypesetterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTypesetterClass) Alloc() NSTypesetter {
	rv := objc.Send[NSTypesetter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that performs various type layout tasks.
//
// # Overview
// 
// [NSLayoutManager] uses concrete subclasses of [NSTypesetter] to perform
// line layout, which includes word wrapping, hyphenation, and line breaking
// in either vertical or horizontal rectangles. By default, the text system
// uses the concrete subclass [NSATSTypesetter].
// 
// # Subclassing Notes
// 
// [NSTypesetter] provides concrete subclasses with default implementation
// interfacing with the Cocoa text system. By subclassing [NSTypesetter], an
// application can override the [NSTypesetter.LayoutParagraphAtPoint] method to integrate a
// custom typesetting engine into the Cocoa text system. On the other hand, an
// application can subclass [NSATSTypesetter] and override the glyph storage
// interface to integrate the concrete subclass into its own custom layout
// system.
// 
// [NSTypesetter] methods belong to three categories: glyph storage interface
// methods, layout phase interface methods, and core typesetter methods. The
// glyph storage interface methods map to [NSLayoutManager] methods. The
// typesetter itself calls these methods, and their default implementations
// call the Cocoa layout manager. An [NSTypesetter] subclass can override
// these methods to call its own glyph storage facility, in which case it
// should override all of them. (This doesn’t preclude the overridden method
// calling its superclass implementation if appropriate).
// 
// The layout phase interface provides control points similar to delegate
// methods; if implemented, the system invokes these methods to notify an
// [NSTypesetter] subclass of events in the layout process so it can intervene
// as needed.
// 
// The remainder of the [NSTypesetter] methods are primitive, core typesetter
// methods. The core typesetter methods correlate with typesetting state
// attributes; the layout manager calls these methods to store its values
// before starting the layout process. If you subclass [NSTypesetter] and
// override the glyph storage interface methods, you can call the core methods
// to control the typesetter directly.
// 
// # Glyph Storage Interface
// 
// Override these methods to use [NSTypesetter]’s built-in concrete
// subclass, [NSATSTypesetter], with a custom glyph storage and layout system
// other than the Cocoa layout manager and text container mechanism.
// 
// - [NSTypesetter.CharacterRangeForGlyphRangeActualGlyphRange] -
// [NSTypesetter.GlyphRangeForCharacterRangeActualCharacterRange] -
// [getGlyphs(in:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:)]
// -
// [NSTypesetter.GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter]
// - [NSTypesetter.SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset] -
// [substituteGlyphs(in:withGlyphs:)] -
// [insertGlyph(_:atGlyphIndex:characterIndex:)] - [deleteGlyphs(in:)] -
// [NSTypesetter.SetNotShownAttributeForGlyphRange] -
// [NSTypesetter.SetDrawsOutsideLineFragmentForGlyphRange] -
// [NSTypesetter.SetLocationWithAdvancementsForStartOfGlyphRange] -
// [NSTypesetter.SetAttachmentSizeForGlyphRange] - [NSTypesetter.SetBidiLevelsForGlyphRange]
// 
// # Layout Phase Interface
// 
// Override these methods to customize the text layout process, including
// modifying line fragments, controlling line breaking and hyphenation, and
// controlling the behavior of tabs and other control glyphs.
// 
// - [NSTypesetter.WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset] -
// [NSTypesetter.ShouldBreakLineByWordBeforeCharacterAtIndex] -
// [NSTypesetter.ShouldBreakLineByHyphenatingBeforeCharacterAtIndex] -
// [NSTypesetter.HyphenationFactorForGlyphAtIndex] - [NSTypesetter.HyphenCharacterForGlyphAtIndex] -
// [NSTypesetter.BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex]
//
// [deleteGlyphs(in:)]: https://developer.apple.com/documentation/AppKit/NSTypesetter/deleteGlyphs(in:)
// [getGlyphs(in:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:)]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getGlyphs(in:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:)
// [insertGlyph(_:atGlyphIndex:characterIndex:)]: https://developer.apple.com/documentation/AppKit/NSTypesetter/insertGlyph(_:atGlyphIndex:characterIndex:)
// [substituteGlyphs(in:withGlyphs:)]: https://developer.apple.com/documentation/AppKit/NSTypesetter/substituteGlyphs(in:withGlyphs:)
//
// # Getting information about glyphs
//
//   - [NSTypesetter.BaselineOffsetInLayoutManagerGlyphIndex]: Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
//
// # Accessing the layout manager
//
//   - [NSTypesetter.LayoutManager]: Returns the layout manager for the text being typeset.
//   - [NSTypesetter.UsesFontLeading]: Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//   - [NSTypesetter.SetUsesFontLeading]
//   - [NSTypesetter.TypesetterBehavior]: Returns the current typesetter behavior.
//   - [NSTypesetter.SetTypesetterBehavior]
//   - [NSTypesetter.HyphenationFactor]: Returns the current hyphenation factor.
//   - [NSTypesetter.SetHyphenationFactor]
//
// # Managing text containers
//
//   - [NSTypesetter.CurrentTextContainer]: Returns the text container for the text being typeset.
//   - [NSTypesetter.TextContainers]: Returns an array containing the text containers belonging to the current layout manager.
//   - [NSTypesetter.LineFragmentPadding]: Returns the current line fragment padding, in points.
//   - [NSTypesetter.SetLineFragmentPadding]
//
// # Performing font substitution
//
//   - [NSTypesetter.SubstituteFontForFont]: Returns a screen font suitable for use in place of a given font.
//
// # Getting the location of text tabs
//
//   - [NSTypesetter.TextTabForGlyphLocationWritingDirectionMaxLocation]: Returns the text tab next closest to a given glyph location within the given parameters.
//
// # Bidirectional text processing
//
//   - [NSTypesetter.BidiProcessingEnabled]: Returns whether bidirectional text processing is enabled.
//   - [NSTypesetter.SetBidiProcessingEnabled]
//
// # Accessing paragraph typesetting information
//
//   - [NSTypesetter.CurrentParagraphStyle]: Returns the paragraph style object for the text being typeset.
//   - [NSTypesetter.AttributedString]: Returns the text backing store, usually an instance of [NSTextStorage].
//   - [NSTypesetter.SetAttributedString]
//   - [NSTypesetter.SetParagraphGlyphRangeSeparatorGlyphRange]: Sets the current glyph range being processed.
//   - [NSTypesetter.ParagraphGlyphRange]: Returns the glyph range currently being processed.
//   - [NSTypesetter.ParagraphSeparatorGlyphRange]: Returns the current paragraph separator range.
//   - [NSTypesetter.ParagraphCharacterRange]: Returns the character range currently being processed.
//   - [NSTypesetter.ParagraphSeparatorCharacterRange]: Returns the current paragraph separator character range.
//   - [NSTypesetter.AttributesForExtraLineFragment]: Returns the attributes used to lay out the extra line fragment.
//
// # Getting spacing information
//
//   - [NSTypesetter.LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect]: Returns the line spacing in effect following the specified glyph.
//   - [NSTypesetter.ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect]: Returns the paragraph spacing that is in effect after the specified glyph.
//   - [NSTypesetter.ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect]: Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph.
//
// # Laying out a paragraph
//
//   - [NSTypesetter.LayoutParagraphAtPoint]: Lays out glyphs in the current glyph range until the next paragraph separator is reached.
//   - [NSTypesetter.BeginParagraph]: Sets up layout parameters at the beginning of a paragraph.
//   - [NSTypesetter.EndParagraph]: Sets up layout parameters at the end of a paragraph.
//   - [NSTypesetter.BeginLineWithGlyphAtIndex]: Sets up layout parameters at the beginning of a line during typesetting.
//   - [NSTypesetter.EndLineWithGlyphRange]: Sets up layout parameters at the end of a line during typesetting.
//
// # Laying out characters
//
//   - [NSTypesetter.LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments]: Lays out characters in the given character range for the specified layout manager.
//
// # Laying out glyphs
//
//   - [NSTypesetter.LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex]: Lays out glyphs in the specified layout manager starting at a specified glyph.
//   - [NSTypesetter.BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex]: Returns the bounding rectangle for the specified control glyph with the specified parameters.
//   - [NSTypesetter.GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin]: Calculates the line fragment rectangle and line fragment used rectangle for blank lines.
//   - [NSTypesetter.GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter]: Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
//   - [NSTypesetter.HyphenCharacterForGlyphAtIndex]: Returns the hyphen character to be inserted after the specified glyph.
//   - [NSTypesetter.HyphenationFactorForGlyphAtIndex]: Returns the hyphenation factor in effect at a specified location.
//   - [NSTypesetter.ShouldBreakLineByHyphenatingBeforeCharacterAtIndex]: Returns whether the line being laid out should be broken by hyphenating at the specified character.
//   - [NSTypesetter.ShouldBreakLineByWordBeforeCharacterAtIndex]: Returns whether the line being laid out should be broken by a word break at the specified character.
//   - [NSTypesetter.WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset]: Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
//   - [NSTypesetter.SetHardInvalidationForGlyphRange]: Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
//
// # Interfacing with Glyph Storage
//
//   - [NSTypesetter.CharacterRangeForGlyphRangeActualGlyphRange]: Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs.
//   - [NSTypesetter.GlyphRangeForCharacterRangeActualCharacterRange]: Returns the range for the glyphs mapped to the characters of the text store in the specified range.
//   - [NSTypesetter.SetAttachmentSizeForGlyphRange]: Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at.
//   - [NSTypesetter.SetBidiLevelsForGlyphRange]: Sets the direction of the specified glyphs for bidirectional text.
//   - [NSTypesetter.SetDrawsOutsideLineFragmentForGlyphRange]: Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
//   - [NSTypesetter.SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset]: Sets the line fragment rectangle where the specified glyphs are laid out.
//   - [NSTypesetter.SetLocationWithAdvancementsForStartOfGlyphRange]: Sets the location where the specified glyphs are laid out.
//   - [NSTypesetter.SetNotShownAttributeForGlyphRange]: Sets whether the specified glyphs are not shown.
//
// # Deprecated
//
//   - [NSTypesetter.ActionForControlCharacterAtIndex]: Returns the action associated with a control character.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter
type NSTypesetter struct {
	objectivec.Object
}

// NSTypesetterFromID constructs a [NSTypesetter] from an objc.ID.
//
// An abstract class that performs various type layout tasks.
func NSTypesetterFromID(id objc.ID) NSTypesetter {
	return NSTypesetter{objectivec.Object{ID: id}}
}
// NOTE: NSTypesetter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTypesetter] class.
//
// # Getting information about glyphs
//
//   - [INSTypesetter.BaselineOffsetInLayoutManagerGlyphIndex]: Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
//
// # Accessing the layout manager
//
//   - [INSTypesetter.LayoutManager]: Returns the layout manager for the text being typeset.
//   - [INSTypesetter.UsesFontLeading]: Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//   - [INSTypesetter.SetUsesFontLeading]
//   - [INSTypesetter.TypesetterBehavior]: Returns the current typesetter behavior.
//   - [INSTypesetter.SetTypesetterBehavior]
//   - [INSTypesetter.HyphenationFactor]: Returns the current hyphenation factor.
//   - [INSTypesetter.SetHyphenationFactor]
//
// # Managing text containers
//
//   - [INSTypesetter.CurrentTextContainer]: Returns the text container for the text being typeset.
//   - [INSTypesetter.TextContainers]: Returns an array containing the text containers belonging to the current layout manager.
//   - [INSTypesetter.LineFragmentPadding]: Returns the current line fragment padding, in points.
//   - [INSTypesetter.SetLineFragmentPadding]
//
// # Performing font substitution
//
//   - [INSTypesetter.SubstituteFontForFont]: Returns a screen font suitable for use in place of a given font.
//
// # Getting the location of text tabs
//
//   - [INSTypesetter.TextTabForGlyphLocationWritingDirectionMaxLocation]: Returns the text tab next closest to a given glyph location within the given parameters.
//
// # Bidirectional text processing
//
//   - [INSTypesetter.BidiProcessingEnabled]: Returns whether bidirectional text processing is enabled.
//   - [INSTypesetter.SetBidiProcessingEnabled]
//
// # Accessing paragraph typesetting information
//
//   - [INSTypesetter.CurrentParagraphStyle]: Returns the paragraph style object for the text being typeset.
//   - [INSTypesetter.AttributedString]: Returns the text backing store, usually an instance of [NSTextStorage].
//   - [INSTypesetter.SetAttributedString]
//   - [INSTypesetter.SetParagraphGlyphRangeSeparatorGlyphRange]: Sets the current glyph range being processed.
//   - [INSTypesetter.ParagraphGlyphRange]: Returns the glyph range currently being processed.
//   - [INSTypesetter.ParagraphSeparatorGlyphRange]: Returns the current paragraph separator range.
//   - [INSTypesetter.ParagraphCharacterRange]: Returns the character range currently being processed.
//   - [INSTypesetter.ParagraphSeparatorCharacterRange]: Returns the current paragraph separator character range.
//   - [INSTypesetter.AttributesForExtraLineFragment]: Returns the attributes used to lay out the extra line fragment.
//
// # Getting spacing information
//
//   - [INSTypesetter.LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect]: Returns the line spacing in effect following the specified glyph.
//   - [INSTypesetter.ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect]: Returns the paragraph spacing that is in effect after the specified glyph.
//   - [INSTypesetter.ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect]: Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph.
//
// # Laying out a paragraph
//
//   - [INSTypesetter.LayoutParagraphAtPoint]: Lays out glyphs in the current glyph range until the next paragraph separator is reached.
//   - [INSTypesetter.BeginParagraph]: Sets up layout parameters at the beginning of a paragraph.
//   - [INSTypesetter.EndParagraph]: Sets up layout parameters at the end of a paragraph.
//   - [INSTypesetter.BeginLineWithGlyphAtIndex]: Sets up layout parameters at the beginning of a line during typesetting.
//   - [INSTypesetter.EndLineWithGlyphRange]: Sets up layout parameters at the end of a line during typesetting.
//
// # Laying out characters
//
//   - [INSTypesetter.LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments]: Lays out characters in the given character range for the specified layout manager.
//
// # Laying out glyphs
//
//   - [INSTypesetter.LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex]: Lays out glyphs in the specified layout manager starting at a specified glyph.
//   - [INSTypesetter.BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex]: Returns the bounding rectangle for the specified control glyph with the specified parameters.
//   - [INSTypesetter.GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin]: Calculates the line fragment rectangle and line fragment used rectangle for blank lines.
//   - [INSTypesetter.GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter]: Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
//   - [INSTypesetter.HyphenCharacterForGlyphAtIndex]: Returns the hyphen character to be inserted after the specified glyph.
//   - [INSTypesetter.HyphenationFactorForGlyphAtIndex]: Returns the hyphenation factor in effect at a specified location.
//   - [INSTypesetter.ShouldBreakLineByHyphenatingBeforeCharacterAtIndex]: Returns whether the line being laid out should be broken by hyphenating at the specified character.
//   - [INSTypesetter.ShouldBreakLineByWordBeforeCharacterAtIndex]: Returns whether the line being laid out should be broken by a word break at the specified character.
//   - [INSTypesetter.WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset]: Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
//   - [INSTypesetter.SetHardInvalidationForGlyphRange]: Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
//
// # Interfacing with Glyph Storage
//
//   - [INSTypesetter.CharacterRangeForGlyphRangeActualGlyphRange]: Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs.
//   - [INSTypesetter.GlyphRangeForCharacterRangeActualCharacterRange]: Returns the range for the glyphs mapped to the characters of the text store in the specified range.
//   - [INSTypesetter.SetAttachmentSizeForGlyphRange]: Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at.
//   - [INSTypesetter.SetBidiLevelsForGlyphRange]: Sets the direction of the specified glyphs for bidirectional text.
//   - [INSTypesetter.SetDrawsOutsideLineFragmentForGlyphRange]: Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
//   - [INSTypesetter.SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset]: Sets the line fragment rectangle where the specified glyphs are laid out.
//   - [INSTypesetter.SetLocationWithAdvancementsForStartOfGlyphRange]: Sets the location where the specified glyphs are laid out.
//   - [INSTypesetter.SetNotShownAttributeForGlyphRange]: Sets whether the specified glyphs are not shown.
//
// # Deprecated
//
//   - [INSTypesetter.ActionForControlCharacterAtIndex]: Returns the action associated with a control character.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter
type INSTypesetter interface {
	objectivec.IObject

	// Topic: Getting information about glyphs

	// Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
	BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr INSLayoutManager, glyphIndex uint) float64

	// Topic: Accessing the layout manager

	// Returns the layout manager for the text being typeset.
	LayoutManager() INSLayoutManager
	// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	// Returns the current typesetter behavior.
	TypesetterBehavior() NSTypesetterBehavior
	SetTypesetterBehavior(value NSTypesetterBehavior)
	// Returns the current hyphenation factor.
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)

	// Topic: Managing text containers

	// Returns the text container for the text being typeset.
	CurrentTextContainer() INSTextContainer
	// Returns an array containing the text containers belonging to the current layout manager.
	TextContainers() []NSTextContainer
	// Returns the current line fragment padding, in points.
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)

	// Topic: Performing font substitution

	// Returns a screen font suitable for use in place of a given font.
	SubstituteFontForFont(originalFont NSFont) NSFont

	// Topic: Getting the location of text tabs

	// Returns the text tab next closest to a given glyph location within the given parameters.
	TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction NSWritingDirection, maxLocation float64) INSTextTab

	// Topic: Bidirectional text processing

	// Returns whether bidirectional text processing is enabled.
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)

	// Topic: Accessing paragraph typesetting information

	// Returns the paragraph style object for the text being typeset.
	CurrentParagraphStyle() INSParagraphStyle
	// Returns the text backing store, usually an instance of [NSTextStorage].
	AttributedString() foundation.NSAttributedString
	SetAttributedString(value foundation.NSAttributedString)
	// Sets the current glyph range being processed.
	SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.NSRange, paragraphSeparatorRange foundation.NSRange)
	// Returns the glyph range currently being processed.
	ParagraphGlyphRange() foundation.NSRange
	// Returns the current paragraph separator range.
	ParagraphSeparatorGlyphRange() foundation.NSRange
	// Returns the character range currently being processed.
	ParagraphCharacterRange() foundation.NSRange
	// Returns the current paragraph separator character range.
	ParagraphSeparatorCharacterRange() foundation.NSRange
	// Returns the attributes used to lay out the extra line fragment.
	AttributesForExtraLineFragment() foundation.INSDictionary

	// Topic: Getting spacing information

	// Returns the line spacing in effect following the specified glyph.
	LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64
	// Returns the paragraph spacing that is in effect after the specified glyph.
	ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64
	// Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph.
	ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64

	// Topic: Laying out a paragraph

	// Lays out glyphs in the current glyph range until the next paragraph separator is reached.
	LayoutParagraphAtPoint(lineFragmentOrigin foundation.NSPoint) uint
	// Sets up layout parameters at the beginning of a paragraph.
	BeginParagraph()
	// Sets up layout parameters at the end of a paragraph.
	EndParagraph()
	// Sets up layout parameters at the beginning of a line during typesetting.
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	// Sets up layout parameters at the end of a line during typesetting.
	EndLineWithGlyphRange(lineGlyphRange foundation.NSRange)

	// Topic: Laying out characters

	// Lays out characters in the given character range for the specified layout manager.
	LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.NSRange, layoutManager INSLayoutManager, maxNumLines uint) foundation.NSRange

	// Topic: Laying out glyphs

	// Lays out glyphs in the specified layout manager starting at a specified glyph.
	LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager INSLayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph unsafe.Pointer)
	// Returns the bounding rectangle for the specified control glyph with the specified parameters.
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer INSTextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect
	// Calculates the line fragment rectangle and line fragment used rectangle for blank lines.
	GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect foundation.NSRect, lineFragmentUsedRect foundation.NSRect, paragraphSeparatorGlyphRange foundation.NSRange, lineOrigin corefoundation.CGPoint)
	// Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
	GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect foundation.NSRect, lineFragmentUsedRect foundation.NSRect, remainingRect foundation.NSRect, startingGlyphIndex uint, proposedRect corefoundation.CGRect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64)
	// Returns the hyphen character to be inserted after the specified glyph.
	HyphenCharacterForGlyphAtIndex(glyphIndex uint) uint32
	// Returns the hyphenation factor in effect at a specified location.
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32
	// Returns whether the line being laid out should be broken by hyphenating at the specified character.
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	// Returns whether the line being laid out should be broken by a word break at the specified character.
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	// Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect foundation.NSRect, glyphRange foundation.NSRange, usedRect foundation.NSRect, baselineOffset unsafe.Pointer)
	// Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
	SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.NSRange)

	// Topic: Interfacing with Glyph Storage

	// Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs.
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.NSRange, actualGlyphRange foundation.NSRange) foundation.NSRange
	// Returns the range for the glyphs mapped to the characters of the text store in the specified range.
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange) foundation.NSRange
	// Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at.
	SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.NSRange)
	// Sets the direction of the specified glyphs for bidirectional text.
	SetBidiLevelsForGlyphRange(levels unsafe.Pointer, glyphRange foundation.NSRange)
	// Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
	SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.NSRange)
	// Sets the line fragment rectangle where the specified glyphs are laid out.
	SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect corefoundation.CGRect, glyphRange foundation.NSRange, usedRect corefoundation.CGRect, baselineOffset float64)
	// Sets the location where the specified glyphs are laid out.
	SetLocationWithAdvancementsForStartOfGlyphRange(location corefoundation.CGPoint, advancements unsafe.Pointer, glyphRange foundation.NSRange)
	// Sets whether the specified glyphs are not shown.
	SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.NSRange)

	// Topic: Deprecated

	// Returns the action associated with a control character.
	ActionForControlCharacterAtIndex(charIndex uint) NSTypesetterControlCharacterAction
}

// Init initializes the instance.
func (t NSTypesetter) Init() NSTypesetter {
	rv := objc.Send[NSTypesetter](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTypesetter) Autorelease() NSTypesetter {
	rv := objc.Send[NSTypesetter](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTypesetter creates a new NSTypesetter instance.
func NewNSTypesetter() NSTypesetter {
	class := getNSTypesetterClass()
	rv := objc.Send[NSTypesetter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the distance from the bottom of the line fragment rectangle in
// which the glyph resides to the glyph baseline.
//
// layoutMgr: The layout manager used for the drawing.
//
// glyphIndex: The index of the glyph in question.
//
// # Return Value
// 
// The distance from the bottom of the line fragment rectangle in which the
// glyph resides to the glyph baseline.
//
// # Discussion
// 
// The text system uses this value to calculate the vertical position of
// underlines.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/baselineOffset(in:glyphIndex:)
func (t NSTypesetter) BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr INSLayoutManager, glyphIndex uint) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("baselineOffsetInLayoutManager:glyphIndex:"), layoutMgr, glyphIndex)
	return rv
}

// Returns a screen font suitable for use in place of a given font.
//
// originalFont: The original font.
//
// # Return Value
// 
// A screen font suitable for use in place of `originalFont`. This method
// returns `originalFont` if a screen font can’t be used or isn’t
// available.
//
// # Discussion
// 
// A screen font can only be substituted if the receiver is set to use screen
// fonts and if no text view associated with the receiver is scaled or
// rotated.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/substituteFont(for:)
func (t NSTypesetter) SubstituteFontForFont(originalFont NSFont) NSFont {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("substituteFontForFont:"), originalFont)
	return NSFontFromID(rv)
}

// Returns the text tab next closest to a given glyph location within the
// given parameters.
//
// glyphLocation: The location at which to start searching.
//
// direction: The direction in which to search.
//
// maxLocation: The maximum location for the search.
//
// # Return Value
// 
// The text tab next closest to `glyphLocation`, indexing in `direction` but
// not beyond `maxLocation`.
//
// # Discussion
// 
// The typesetter calls this method whenever it finds a tab character. To
// determine the width to advance the next glyph, the typesetter examines the
// [NSParagraphStyle] object’s tab array and the default tab interval.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/textTab(forGlyphLocation:writingDirection:maxLocation:)
func (t NSTypesetter) TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction NSWritingDirection, maxLocation float64) INSTextTab {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return NSTextTabFromID(rv)
}

// Sets the current glyph range being processed.
//
// paragraphRange: The current glyph range being processed.
//
// paragraphSeparatorRange: The range of the paragraph separator character or characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setParagraphGlyphRange(_:separatorGlyphRange:)
func (t NSTypesetter) SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.NSRange, paragraphSeparatorRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}

// Returns the line spacing in effect following the specified glyph.
//
// glyphIndex: The index of the glyph in question.
//
// rect: The proposed line fragment rectangle.
//
// # Return Value
// 
// The line spacing in effect following the glyph at `glyphIndex`.
//
// # Discussion
// 
// The [NSATSTypesetter] calls this method to determine the number of points
// of space to include below the descenders in the used rectangle for the
// proposed line fragment rectangle `rect`.
// 
// Line spacing, also called leading, is an attribute of [NSParagraphStyle],
// which you can set on an [NSMutableParagraphStyle] object. A font typically
// includes a default minimum line spacing metric used if none is set in the
// paragraph style.
// 
// If the typesetter behavior specified in the layout manager is
// [NSTypesetterOriginalBehavior], the text system uses the original, private
// typesetter [NSSimpleHorizontalTypesetter], which adds the line spacing
// above the ascender. Similarly, [NSATSTypesetter] adds the line spacing
// above the ascender if the value is negative.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t NSTypesetter) LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Returns the paragraph spacing that is in effect after the specified glyph.
//
// glyphIndex: The index of the glyph in question.
//
// rect: The line fragment rectangle of the last line in the paragraph.
//
// # Return Value
// 
// The paragraph spacing—that is, the number of points of space added
// following a paragraph—that is in effect after the glyph specified by
// `glyphIndex`.
//
// # Discussion
// 
// The typesetter adds the number of points specified in the return value to
// the bottom of the line fragment rectangle specified by `rect` (but not to
// the used line fragment rectangle for that line). Paragraph spacing added
// after a paragraph correlates to the value returned by the
// [ParagraphSpacing] method of [NSParagraphStyle], which you can set using
// the [ParagraphSpacing] method of [NSMutableParagraphStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t NSTypesetter) ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Returns the number of points of space—added before a paragraph—that is
// in effect before the specified glyph.
//
// glyphIndex: The index of the glyph in question.
//
// rect: The line fragment rectangle of the first line in the paragraph.
//
// # Return Value
// 
// The number of points of space—added before a paragraph—that is in
// effect before the glyph specified by `glyphIndex`.
//
// # Discussion
// 
// The typesetter adds the number of points specified in the return value to
// the top of the line fragment rectangle specified by `rect` (but not to the
// used line fragment rectangle for that line). Paragraph spacing added before
// a paragraph correlates to the value returned by the
// [ParagraphSpacingBefore] method of [NSParagraphStyle], which you can set
// using the [ParagraphSpacingBefore] method of [NSMutableParagraphStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(beforeGlyphAt:withProposedLineFragmentRect:)
func (t NSTypesetter) ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Lays out glyphs in the current glyph range until the next paragraph
// separator is reached.
//
// lineFragmentOrigin: The upper-left corner of line fragment rectangle. On return,
// `lineFragmentOrigin` contains the next origin.
//
// # Return Value
// 
// The next glyph index; usually the index right after the paragraph
// separator, but it can be inside the paragraph range if, for example, the
// end of the text container is reached before the paragraph separator.
//
// # Discussion
// 
// Concrete subclasses must implement this method. A concrete implementation
// must invoke [BeginParagraph], [BeginLineWithGlyphAtIndex],
// [EndLineWithGlyphRange], and [EndParagraph].
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutParagraph(at:)
func (t NSTypesetter) LayoutParagraphAtPoint(lineFragmentOrigin foundation.NSPoint) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}

// Sets up layout parameters at the beginning of a paragraph.
//
// # Discussion
// 
// Concrete subclasses should invoke this method at the beginning of their
// [LayoutParagraphAtPoint] implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginParagraph()
func (t NSTypesetter) BeginParagraph() {
	objc.Send[objc.ID](t.ID, objc.Sel("beginParagraph"))
}

// Sets up layout parameters at the end of a paragraph.
//
// # Discussion
// 
// Concrete subclasses should invoke this method at the end of their
// [LayoutParagraphAtPoint] implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/endParagraph()
func (t NSTypesetter) EndParagraph() {
	objc.Send[objc.ID](t.ID, objc.Sel("endParagraph"))
}

// Sets up layout parameters at the beginning of a line during typesetting.
//
// glyphIndex: The index of the first glyph to be laid out in the line.
//
// # Discussion
// 
// Concrete subclass implementations of [LayoutParagraphAtPoint] should invoke
// this method at the beginning of each line.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginLine(withGlyphAt:)
func (t NSTypesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("beginLineWithGlyphAtIndex:"), glyphIndex)
}

// Sets up layout parameters at the end of a line during typesetting.
//
// lineGlyphRange: The range of glyphs laid out in the line.
//
// # Discussion
// 
// Concrete subclass implementations of [LayoutParagraphAtPoint] should invoke
// this method at the end of each line.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/endLine(withGlyphRange:)
func (t NSTypesetter) EndLineWithGlyphRange(lineGlyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("endLineWithGlyphRange:"), lineGlyphRange)
}

// Lays out characters in the given character range for the specified layout
// manager.
//
// characterRange: The range of the characters to lay out.
//
// layoutManager: The layout manager that does the drawing.
//
// maxNumLines: The maximum number of line fragments to lay out. Specify [NSUIntegerMax]
// for unlimited number of line fragments.
//
// # Return Value
// 
// The method returns the actual character range that the receiving
// [NSTypesetter] processed.
//
// # Discussion
// 
// The layout process can be interrupted when the number of line fragments
// exceeds `maxNumLines`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutCharacters(in:for:maximumNumberOfLineFragments:)
func (t NSTypesetter) LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.NSRange, layoutManager INSLayoutManager, maxNumLines uint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:"), characterRange, layoutManager, maxNumLines)
	return foundation.NSRange(rv)
}

// Lays out glyphs in the specified layout manager starting at a specified
// glyph.
//
// layoutManager: The layout manager in which to lay out glyphs.
//
// startGlyphIndex: The index of the starting glyph.
//
// maxNumLines: The maximum number of lines to generate. Fewer lines may be laid out if the
// glyph storage runs out of glyphs.
//
// nextGlyph: On return, set to the index of the next glyph that needs to be laid out.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutGlyphs(in:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:)
func (t NSTypesetter) LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager INSLayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph unsafe.Pointer) {
	objc.Send[objc.ID](t.ID, objc.Sel("layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:"), layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}

// Returns the bounding rectangle for the specified control glyph with the
// specified parameters.
//
// glyphIndex: The index of the control glyph in question.
//
// textContainer: The text container to use to calculate the position.
//
// proposedRect: The proposed line fragment rectangle.
//
// glyphPosition: The position of the glyph in `textContainer`.
//
// charIndex: The character index in `textContainer`.
//
// # Return Value
// 
// The bounding rectangle of the control glyph at `glyphIndex`, at the given
// `glyphPosition` and character index `charIndex`, in `textContainer`.
//
// # Discussion
// 
// The typesetter calls this method when it encounters a control glyph. The
// default behavior is to return zero width for control glyphs. A subclass can
// override this method to do something different, such as implement a way to
// display control characters.
// 
// [NSGlyphGenerator] can choose whether or not to map control characters to
// [ControlGlyph]. Tab characters, for example, do not use this facility.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/boundingBox(forControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:)
func (t NSTypesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer INSTextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return corefoundation.CGRect(rv)
}

// Calculates the line fragment rectangle and line fragment used rectangle for
// blank lines.
//
// lineFragmentRect: On return, the calculated line fragment rectangle.
//
// lineFragmentUsedRect: On return, the used rectangle (the portion of the line fragment rectangle
// that actually contains marks).
//
// paragraphSeparatorGlyphRange: The range of glyphs under consideration. A `paragraphSeparatorGlyphRange`
// with length 0 indicates an extra line fragment (which occurs if the last
// character in the paragraph is a line separator).
//
// lineOrigin: The origin point of the line fragment rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:)
func (t NSTypesetter) GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect foundation.NSRect, lineFragmentUsedRect foundation.NSRect, paragraphSeparatorGlyphRange foundation.NSRange, lineOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}

// Calculates line fragment rectangle, line fragment used rectangle, and
// remaining rectangle for a line fragment.
//
// lineFragmentRect: On return, the calculated line fragment rectangle.
//
// lineFragmentUsedRect: On return, the used rectangle (the portion of the line fragment rectangle
// that actually contains marks).
//
// remainingRect: On return, the remaining rectangle of `proposedRect`.
//
// startingGlyphIndex: The glyph index where the line fragment starts.
//
// proposedRect: The proposed rectangle of the line fragment.
//
// lineSpacing: The line spacing.
//
// paragraphSpacingBefore: The spacing before the paragraph.
//
// paragraphSpacingAfter: The spacing after the paragraph.
//
// # Discussion
// 
// The height of the line fragment is determined using `lineSpacing`,
// `paragraphSpacingBefore`, and `paragraphSpacingAfter` as well as
// `proposedRect`. The width for `lineFragmentUsedRect` is set to the
// `lineFragmentRect` width. In the standard implementation, paragraph spacing
// is included in the line fragment rectangle but not the line fragment used
// rectangle; line spacing is included in both.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:remaining:forStartingGlyphAt:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:)
func (t NSTypesetter) GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect foundation.NSRect, lineFragmentUsedRect foundation.NSRect, remainingRect foundation.NSRect, startingGlyphIndex uint, proposedRect corefoundation.CGRect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	objc.Send[objc.ID](t.ID, objc.Sel("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}

// Returns the hyphen character to be inserted after the specified glyph.
//
// glyphIndex: The index of the glyph in question.
//
// # Return Value
// 
// The hyphen character to be inserted after the glyph at `glyphIndex`.
//
// # Discussion
// 
// The typesetter calls this method before hyphenating. A subclass can
// override this method to return a different hyphen glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenCharacter(forGlyphAt:)
func (t NSTypesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the hyphenation factor in effect at a specified location.
//
// glyphIndex: The index of the glyph position to examine.
//
// # Return Value
// 
// The hyphenation factor in effect at `glyphIndex`. The hyphenation factor is
// a value ranging from 0.0 to 1.0 that controls when hyphenation is
// attempted. By default, the value is 0.0, meaning hyphenation is off. A
// factor of 1.0 causes hyphenation to be attempted always.
//
// # Discussion
// 
// The typesetter calls this method with a proposed hyphenation point for a
// line break to find the hyphenation factor in effect at that time. A
// subclass can override this method to customize the text layout process.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor(forGlyphAt:)
func (t NSTypesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns whether the line being laid out should be broken by hyphenating at
// the specified character.
//
// charIndex: The index of the character just after the proposed hyphenation would occur.
//
// # Return Value
// 
// [true] if the line should be broken by hyphenating, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The typesetter calls this method, if implemented by a subclass, before
// breaking a line by hyphenating before the character at `charIndex`,
// enabling the subclass to control line breaking.
// 
// A subclass can override this method to customize the text layout process.
// If the method returns [false], the typesetter continues looking for a break
// point.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byHyphenatingBeforeCharacterAt:)
func (t NSTypesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}

// Returns whether the line being laid out should be broken by a word break at
// the specified character.
//
// charIndex: The index of the character just after the proposed word break would occur.
//
// # Return Value
// 
// [true] if the line should be broken by a word break, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The typesetter calls this method, if implemented by a subclass, before
// breaking a line by word wrapping before the character at `charIndex`,
// enabling the subclass to control line breaking.
// 
// A subclass can override this method to customize the text layout process.
// If the method returns [false], the typesetter continues looking for a break
// point.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byWordBeforeCharacterAt:)
func (t NSTypesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}

// Called by the typesetter just prior to storing the actual line fragment
// rectangle location in the layout manager.
//
// lineRect: The rectangle in which the glyphs in `glyphRange` are laid out.
//
// glyphRange: The range of the glyphs to lay out.
//
// usedRect: The portion of `lineRect`, in the NSTextContainer object’s coordinate
// system, that actually contains glyphs or other marks that are drawn
// (including the text container’s line fragment padding). The `usedRect`
// must be equal to or contained within `lineRect`.
//
// baselineOffset: The vertical distance in pixels from the line fragment origin to the
// baseline on which the glyphs align.
//
// # Discussion
// 
// Called by the typesetter just prior to calling
// [SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset] which stores the
// actual line fragment rectangle location in the layout manager.
// 
// A subclass can override this method to customize the text layout process.
// For example, it could change the shape of the line fragment rectangle. The
// subclass is responsible for ensuring that the modified rectangle remains
// valid (for example, that it lies within the text container).
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/willSetLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t NSTypesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect foundation.NSRect, glyphRange foundation.NSRange, usedRect foundation.NSRect, baselineOffset unsafe.Pointer) {
	objc.Send[objc.ID](t.ID, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}

// Sets whether to force the layout manager to invalidate the specified
// portion of the glyph cache when invalidating layout.
//
// flag: [true] if the layout manager should invalidate the specified portion of the
// glyph cache, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// glyphRange: The range of glyphs in the cache to be marked for hard invalidation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setHardInvalidation(_:forGlyphRange:)
func (t NSTypesetter) SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}

// Returns the range for the characters in the receiver’s text store that
// are mapped to the specified glyphs.
//
// glyphRange: The range of glyphs.
//
// actualGlyphRange: On return, the range of all glyphs mapped to the characters in the
// receiver’s text store. May be [NULL].
//
// # Return Value
// 
// The range for the characters in the receiver’s text store that are mapped
// to the glyphs in `glyphRange`.
//
// # Discussion
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/characterRange(forGlyphRange:actualGlyphRange:)
func (t NSTypesetter) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.NSRange, actualGlyphRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return foundation.NSRange(rv)
}

// Returns the range for the glyphs mapped to the characters of the text store
// in the specified range.
//
// charRange: The range of the characters whose glyph range is desired.
//
// actualCharRange: On return, all characters mapped to those glyphs; may be [NULL].
//
// # Return Value
// 
// The range for the glyphs mapped to the characters of the text store in
// `charRange`.
//
// # Discussion
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/glyphRange(forCharacterRange:actualCharacterRange:)
func (t NSTypesetter) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return foundation.NSRange(rv)
}

// Sets the size the specified glyphs (assumed to be attachments) will be
// asked to draw themselves at.
//
// attachmentSize: The size the glyphs in `glyphRange` (assumed to be attachments) will be
// asked to draw themselves at.
//
// glyphRange: The range of glyphs the attachment size applies to.
//
// # Discussion
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setAttachmentSize(_:forGlyphRange:)
func (t NSTypesetter) SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

// Sets the direction of the specified glyphs for bidirectional text.
//
// levels: Values in `levels` can range from 0 to 61 as defined by Unicode Standard
// Annex #9.
//
// glyphRange: The range of glyphs for which the bidirectional text levels are desired.
//
// # Discussion
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setBidiLevels(_:forGlyphRange:)
func (t NSTypesetter) SetBidiLevelsForGlyphRange(levels unsafe.Pointer, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setBidiLevels:forGlyphRange:"), levels, glyphRange)
}

// Sets whether the specified glyphs exceed the bounds of the line fragment in
// which they are laid out.
//
// flag: [true] if the glyphs in `glyphRange` exceed the bounds of the line fragment
// in which they are laid out, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// glyphRange: The range of the glyphs in question.
//
// # Discussion
// 
// This can happen when text is set at a fixed line height. For example, if
// the user specifies a fixed line height of 12 points and sets the font size
// to 24 points, the glyphs will exceed their layout rectangles.
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setDrawsOutsideLineFragment(_:forGlyphRange:)
func (t NSTypesetter) SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}

// Sets the line fragment rectangle where the specified glyphs are laid out.
//
// fragmentRect: The line fragment rectangle where the glyphs in `glyphRange` are laid out.
//
// glyphRange: The range of the specified glyphs.
//
// usedRect: The portion of `fragmentRect`, in the NSTextContainer object’s coordinate
// system, that actually contains glyphs or other marks that are drawn
// (including the text container’s line fragment padding). The `usedRect`
// must be equal to or contained within `fragmentRect`.
//
// baselineOffset: The vertical distance in pixels from the line fragment origin to the
// baseline on which the glyphs align.
//
// # Discussion
// 
// The exact positions of the glyphs must be set after the line fragment
// rectangle with ``.
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t NSTypesetter) SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect corefoundation.CGRect, glyphRange foundation.NSRange, usedRect corefoundation.CGRect, baselineOffset float64) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), fragmentRect, glyphRange, usedRect, baselineOffset)
}

// Sets the location where the specified glyphs are laid out.
//
// location: The location where the glyphs in `glyphRange` are laid out. The
// x-coordinate of `location` is expressed relative to the line fragment
// rectangle origin, and the y-coordinate is expressed relative to the
// baseline previously specified by
// [SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset].
//
// advancements: The nominal glyph advance width specified in the font metric information.
//
// glyphRange: The range of glyphs whose layout location is being set. This series of
// glyphs can be displayed with a single PostScript `show` operation (a
// nominal range).
//
// # Discussion
// 
// Setting the location for a series of glyphs implies that the glyphs
// preceding it can’t be included in a single `show` operation.
// 
// Before setting the location for a glyph range, you must specify line
// fragment rectangle with
// [SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset].
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLocation(_:withAdvancements:forStartOfGlyphRange:)
func (t NSTypesetter) SetLocationWithAdvancementsForStartOfGlyphRange(location corefoundation.CGPoint, advancements unsafe.Pointer, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}

// Sets whether the specified glyphs are not shown.
//
// flag: [true] if the glyphs in `glyphRange` are not shown, [false] if they are
// shown.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// glyphRange: The range of glyphs in question.
//
// # Discussion
// 
// For example, a tab or newline character doesn’t leave any marks; it just
// indicates where following glyphs are laid out.
// 
// A subclass can override this method to interact with custom glyph storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/setNotShownAttribute(_:forGlyphRange:)
func (t NSTypesetter) SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}

// Returns the action associated with a control character.
//
// charIndex: The index of the control character.
//
// # Return Value
// 
// The action associated with the control character at `charIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/actionForControlCharacter(at:)
func (t NSTypesetter) ActionForControlCharacterAtIndex(charIndex uint) NSTypesetterControlCharacterAction {
	rv := objc.Send[NSTypesetterControlCharacterAction](t.ID, objc.Sel("actionForControlCharacterAtIndex:"), charIndex)
	return NSTypesetterControlCharacterAction(rv)
}

// Returns a shared instance of a reentrant typesetter that implements
// typesetting with the specified behavior.
//
// behavior: The desired behavior.
//
// # Return Value
// 
// A shared instance of a reentrant typesetter that implements typesetting
// with the specified behavior.
//
// # Discussion
// 
// Possible return values are described in the
// [NSLayoutManager.TypesetterBehavior] section for [NSLayoutManager].
//
// [NSLayoutManager.TypesetterBehavior]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter(for:)
func (_NSTypesetterClass NSTypesetterClass) SharedSystemTypesetterForBehavior(behavior NSTypesetterBehavior) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSTypesetterClass.class), objc.Sel("sharedSystemTypesetterForBehavior:"), behavior)
	return objectivec.Object{ID: rv}
}

// Returns the interglyph spacing in the specified range when sent to a
// printer.
//
// layoutMgr: The layout manager that will do the drawing.
//
// nominallySpacedGlyphsRange: The range of the glyphs whose spacing is desired.
//
// packedGlyphs: The glyphs as they are packed for sending to be drawn in `layoutMgr`.
//
// packedGlyphsCount: The number of glyphs in `packedGlyphs`.
//
// # Return Value
// 
// The interglyph spacing in the specified range when sent to a printer. If
// the font metrics of the font used for displaying text on the screen is
// different from the font metrics of the font used in printing, then this
// interglyph spacing may need to be adjusted slightly to match that used on
// the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/printingAdjustment(in:forNominallySpacedGlyphRange:packedGlyphs:count:)
func (_NSTypesetterClass NSTypesetterClass) PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr INSLayoutManager, nominallySpacedGlyphsRange foundation.NSRange, packedGlyphs unsafe.Pointer, packedGlyphsCount uint) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](objc.ID(_NSTypesetterClass.class), objc.Sel("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), layoutMgr, nominallySpacedGlyphsRange, objc.CArray(packedGlyphs), packedGlyphsCount)
	return corefoundation.CGSize(rv)
}

// Returns the layout manager for the text being typeset.
//
// # Return Value
// 
// The layout manager for the text being typeset. This value is valid only
// while the typesetter is performing layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutManager
func (t NSTypesetter) LayoutManager() INSLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("layoutManager"))
	return NSLayoutManagerFromID(objc.ID(rv))
}

// Returns whether the typesetter uses the leading (or line gap) value
// specified in the font metric information of the current font.
//
// # Return Value
// 
// [true] if it uses the information in the font metrics, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t NSTypesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFontLeading"))
	return rv
}
func (t NSTypesetter) SetUsesFontLeading(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFontLeading:"), value)
}

// Returns the current typesetter behavior.
//
// # Return Value
// 
// The current typesetter behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/typesetterBehavior
func (t NSTypesetter) TypesetterBehavior() NSTypesetterBehavior {
	rv := objc.Send[NSTypesetterBehavior](t.ID, objc.Sel("typesetterBehavior"))
	return NSTypesetterBehavior(rv)
}
func (t NSTypesetter) SetTypesetterBehavior(value NSTypesetterBehavior) {
	objc.Send[struct{}](t.ID, objc.Sel("setTypesetterBehavior:"), value)
}

// Returns the current hyphenation factor.
//
// # Return Value
// 
// The hyphenation factor, a value ranging from 0.0 to 1.0 that controls when
// hyphenation is attempted. By default, the value is 0.0, meaning hyphenation
// is off. A factor of 1.0 causes hyphenation to be attempted always.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t NSTypesetter) HyphenationFactor() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("hyphenationFactor"))
	return rv
}
func (t NSTypesetter) SetHyphenationFactor(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setHyphenationFactor:"), value)
}

// Returns the text container for the text being typeset.
//
// # Return Value
// 
// The text container for the text being typeset. This value is valid only
// while the typesetter is performing layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentTextContainer
func (t NSTypesetter) CurrentTextContainer() INSTextContainer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentTextContainer"))
	return NSTextContainerFromID(objc.ID(rv))
}

// Returns an array containing the text containers belonging to the current
// layout manager.
//
// # Return Value
// 
// An array containing the text containers belonging to the current layout
// manager. This value is valid only while the typesetter is performing
// layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/textContainers
func (t NSTypesetter) TextContainers() []NSTextContainer {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textContainers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextContainer {
		return NSTextContainerFromID(id)
	})
}

// Returns the current line fragment padding, in points.
//
// # Return Value
// 
// The current line fragment padding, in points; that is, the portion on each
// end of the line fragment rectangle left blank.
// 
// # Discussion
// 
// Text is inset within the line fragment rectangle by this amount.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t NSTypesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("lineFragmentPadding"))
	return rv
}
func (t NSTypesetter) SetLineFragmentPadding(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setLineFragmentPadding:"), value)
}

// Returns whether bidirectional text processing is enabled.
//
// # Return Value
// 
// [true] if bidirectional text processing is enabled, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/bidiProcessingEnabled
func (t NSTypesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}
func (t NSTypesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}

// Returns the paragraph style object for the text being typeset.
//
// # Return Value
// 
// The paragraph style object for the text being typeset. This value is valid
// only while the typesetter is performing layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentParagraphStyle
func (t NSTypesetter) CurrentParagraphStyle() INSParagraphStyle {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentParagraphStyle"))
	return NSParagraphStyleFromID(objc.ID(rv))
}

// Returns the text backing store, usually an instance of [NSTextStorage].
//
// # Return Value
// 
// The text backing store.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributedString
func (t NSTypesetter) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (t NSTypesetter) SetAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttributedString:"), value)
}

// Returns the glyph range currently being processed.
//
// # Return Value
// 
// The glyph range currently being processed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphGlyphRange
func (t NSTypesetter) ParagraphGlyphRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("paragraphGlyphRange"))
	return foundation.NSRange(rv)
}

// Returns the current paragraph separator range.
//
// # Return Value
// 
// The current paragraph separator range, which is the full range that
// contains the current glyph range and that extends from one paragraph
// separator character to the next.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorGlyphRange
func (t NSTypesetter) ParagraphSeparatorGlyphRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return foundation.NSRange(rv)
}

// Returns the character range currently being processed.
//
// # Return Value
// 
// The character range currently being processed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphCharacterRange
func (t NSTypesetter) ParagraphCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("paragraphCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the current paragraph separator character range.
//
// # Return Value
// 
// The current paragraph separator character range, which is the full range
// that contains the current character range and that extends from one
// paragraph separator character to the next.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorCharacterRange
func (t NSTypesetter) ParagraphSeparatorCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the attributes used to lay out the extra line fragment.
//
// # Return Value
// 
// A dictionary of attributes used to lay out the extra line fragment.
// 
// # Discussion
// 
// The default implementation tries to use the [NSTextView] method
// [TypingAttributes] if possible; otherwise, it uses the attributes for the
// last character.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributesForExtraLineFragment
func (t NSTypesetter) AttributesForExtraLineFragment() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributesForExtraLineFragment"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Returns a shared instance of a reentrant typesetter.
//
// # Return Value
// 
// The shared system typesetter. This typesetter is reentrant.
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter
func (_NSTypesetterClass NSTypesetterClass) SharedSystemTypesetter() NSTypesetter {
	rv := objc.Send[objc.ID](objc.ID(_NSTypesetterClass.class), objc.Sel("sharedSystemTypesetter"))
	return NSTypesetterFromID(objc.ID(rv))
}

// Returns the default typesetter behavior.
//
// # Return Value
// 
// The default typesetter behavior.
// 
// # Discussion
// 
// Possible return values are described in the
// [NSLayoutManager.TypesetterBehavior] section for [NSLayoutManager].
//
// [NSLayoutManager.TypesetterBehavior]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTypesetter/defaultTypesetterBehavior
func (_NSTypesetterClass NSTypesetterClass) DefaultTypesetterBehavior() NSTypesetterBehavior {
	rv := objc.Send[NSTypesetterBehavior](objc.ID(_NSTypesetterClass.class), objc.Sel("defaultTypesetterBehavior"))
	return NSTypesetterBehavior(rv)
}

