// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLayoutManager] class.
var (
	_NSLayoutManagerClass     NSLayoutManagerClass
	_NSLayoutManagerClassOnce sync.Once
)

func getNSLayoutManagerClass() NSLayoutManagerClass {
	_NSLayoutManagerClassOnce.Do(func() {
		_NSLayoutManagerClass = NSLayoutManagerClass{class: objc.GetClass("NSLayoutManager")}
	})
	return _NSLayoutManagerClass
}

// GetNSLayoutManagerClass returns the class object for NSLayoutManager.
func GetNSLayoutManagerClass() NSLayoutManagerClass {
	return getNSLayoutManagerClass()
}

type NSLayoutManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLayoutManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutManagerClass) Alloc() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that coordinates the layout and display of text characters.
//
// # Overview
//
// [NSLayoutManager] maps Unicode character codes to glyphs, sets the glyphs
// in a series of [NSTextContainer] objects, and displays them in a series of
// [NSTextView] objects. In addition to its core function of laying out text,
// a layout manager object coordinates its text view objects, provides
// services to those text views to support [NSRulerView] instances for editing
// paragraph styles, and handles the layout and display of text attributes not
// inherent in glyphs (such as underline or strikethrough). You can create a
// subclass of [NSLayoutManager] to handle additional text attributes, whether
// inherent or not.
//
// # Text Antialiasing
//
// [NSLayoutManager] provides the threshold for text antialiasing. It looks at
// the [AppleAntiAliasingThreshold] default value. If the font size is smaller
// than or equal to this threshold size, the text is rendered aliased by
// [NSLayoutManager]. In macOS, you can change the threshold value from the
// Appearance pane of System Preferences.
//
// # Thread Safety of NSLayoutManager
//
// Generally speaking, a specific layout manager (and associated objects)
// should not be used in more than one block, operation, or thread at a time.
// Most layout managers are used on the main thread, since it is the main
// thread on which their text views are displayed, and since background layout
// occurs on the main thread.
//
// If you want to use a layout manager on a background thread, first make sure
// that text views associated with that layout manager (if any) are not
// displayed while the layout manager is being used on the background thread,
// and, second, turn off background layout for that layout manager while it is
// being used on the background thread. The most effective way to ensure that
// no text view is displayed, without knowing deep implementation, is just not
// to connect a text view to the layout manager.
//
// # Noncontiguous Layout
//
// Noncontiguous layout is an optional layout manager behavior. Previously,
// both glyph generation and layout were always performed, in order, from the
// beginning to the end of the document. When noncontiguous layout is turned
// on, however, the layout manager gains the option of performing glyph
// generation or layout for one portion of the document without having done so
// for previous sections. This can provide significant performance
// improvements for large documents.
//
// Noncontiguous layout is not turned on automatically because direct clients
// of [NSLayoutManager] typically have relied on the previous behavior—for
// example, by forcing layout for a specific glyph range, and then assuming
// that previous glyphs would therefore be laid out. Clients who use
// [NSLayoutManager] only indirectly—for example, those who use [NSTextView]
// without directly calling the underlying layout manager—can usually turn
// on noncontiguous layout without difficulty. Clients using [NSLayoutManager]
// directly need to examine their usage before turning on noncontiguous
// layout.
//
// Enable noncontiguous layout using the [NSLayoutManager.AllowsNonContiguousLayout] property.
// In addition, see the other methods in [NSLayoutManager], many of which
// enable you to ensure that glyph generation and layout are performed for
// specified portions of the text. The behavior of a number of other layout
// manager methods is affected by the state of noncontiguous layout, as noted
// in the discussion sections of those method descriptions.
//
// # Creating a layout manager
//
//   - [NSLayoutManager.InitWithCoder]: Creates a layout manager from data in an unarchiver.
//
// # Managing the layout process
//
//   - [NSLayoutManager.Delegate]: The layout manager’s delegate.
//   - [NSLayoutManager.SetDelegate]
//
// # Accessing the text storage
//
//   - [NSLayoutManager.TextStorage]: The text storage object that contains the content to lay out.
//   - [NSLayoutManager.SetTextStorage]
//   - [NSLayoutManager.ReplaceTextStorage]: Replaces the layout manager’s current text storage object with the specified object.
//
// # Configuring the global layout manager options
//
//   - [NSLayoutManager.AllowsNonContiguousLayout]: A Boolean value that indicates whether the layout manager allows noncontiguous layout.
//   - [NSLayoutManager.SetAllowsNonContiguousLayout]
//   - [NSLayoutManager.HasNonContiguousLayout]: A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
//   - [NSLayoutManager.ShowsInvisibleCharacters]: A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//   - [NSLayoutManager.SetShowsInvisibleCharacters]
//   - [NSLayoutManager.ShowsControlCharacters]: A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//   - [NSLayoutManager.SetShowsControlCharacters]
//   - [NSLayoutManager.UsesFontLeading]: A Boolean value that indicates whether the layout manager uses the leading of the font.
//   - [NSLayoutManager.SetUsesFontLeading]
//   - [NSLayoutManager.BackgroundLayoutEnabled]: A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
//   - [NSLayoutManager.SetBackgroundLayoutEnabled]
//   - [NSLayoutManager.LimitsLayoutForSuspiciousContents]: A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//   - [NSLayoutManager.SetLimitsLayoutForSuspiciousContents]
//   - [NSLayoutManager.UsesDefaultHyphenation]: A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//   - [NSLayoutManager.SetUsesDefaultHyphenation]
//
// # Managing the text containers
//
//   - [NSLayoutManager.TextContainers]: The current text containers of the layout manager.
//   - [NSLayoutManager.AddTextContainer]: Appends the specified text container to the series of text containers where the layout manager arranges text.
//   - [NSLayoutManager.InsertTextContainerAtIndex]: Inserts a text container at the specified index in the list of text containers.
//   - [NSLayoutManager.RemoveTextContainerAtIndex]: Removes the text container at the specified index and invalidates the layout as necessary.
//   - [NSLayoutManager.SetTextContainerForGlyphRange]: Associates a text container with the specified range of glyphs.
//   - [NSLayoutManager.TextContainerChangedGeometry]: Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
//   - [NSLayoutManager.TextContainerChangedTextView]: Updates the information necessary to manage text view objects for the specified text container.
//   - [NSLayoutManager.TextContainerForGlyphAtIndexEffectiveRange]: Returns the text container that manages the layout for the specified glyph, causing layout to happen as necessary.
//   - [NSLayoutManager.TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the text container that manages the layout for the specified glyph.
//   - [NSLayoutManager.UsedRectForTextContainer]: Returns the bounding rectangle for the glyphs in the specified text container.
//
// # Invalidating glyphs and layout
//
//   - [NSLayoutManager.InvalidateDisplayForCharacterRange]: Invalidates display for the specified character range.
//   - [NSLayoutManager.InvalidateDisplayForGlyphRange]: Invalidates a range of glyphs, requiring new layout information, and updates the appropriate regions of any text views that display those glyphs.
//   - [NSLayoutManager.InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange]: Invalidates and adjusts the glyphs in the specified character range.
//   - [NSLayoutManager.InvalidateLayoutForCharacterRangeActualCharacterRange]: Invalidates the layout information for the glyphs that map to the specified character range.
//   - [NSLayoutManager.ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange]: Notifies the layout manager when an edit action changes the contents of its text storage object.
//
// # Causing glyph generation and layout
//
//   - [NSLayoutManager.EnsureGlyphsForCharacterRange]: Forces the layout manager to generate glyphs for the specified character range if it hasn’t already.
//   - [NSLayoutManager.EnsureGlyphsForGlyphRange]: Forces the layout manager to generate glyphs for the specified glyph range if it hasn’t already.
//   - [NSLayoutManager.EnsureLayoutForBoundingRectInTextContainer]: Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
//   - [NSLayoutManager.EnsureLayoutForCharacterRange]: Forces the layout manager to perform layout for the specified character range if it hasn’t already.
//   - [NSLayoutManager.EnsureLayoutForGlyphRange]: Forces the layout manager to perform layout for the specified glyph range if it hasn’t already.
//   - [NSLayoutManager.EnsureLayoutForTextContainer]: Forces the layout manager to perform layout for the specified text container if it hasn’t already.
//   - [NSLayoutManager.GlyphGenerator]: The glyph generator that the layout manager uses.
//   - [NSLayoutManager.SetGlyphGenerator]
//
// # Accessing glyphs
//
//   - [NSLayoutManager.GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels]: Fills a passed-in buffer with a sequence of glyphs.
//   - [NSLayoutManager.CGGlyphAtIndex]: Returns the glyph at the specified index.
//   - [NSLayoutManager.CGGlyphAtIndexIsValidIndex]: Returns the glyph at the specified index along with information about whether the glyph index is valid.
//   - [NSLayoutManager.SetGlyphsPropertiesCharacterIndexesFontForGlyphRange]: Stores the initial glyphs and glyph properties for a character range.
//   - [NSLayoutManager.CharacterIndexForGlyphAtIndex]: Returns the index in the text storage for the first character of the specified glyph.
//   - [NSLayoutManager.GlyphIndexForCharacterAtIndex]: Returns the index of the first glyph of the character at the specified index.
//   - [NSLayoutManager.IsValidGlyphIndex]: Indicates whether the specified index refers to a valid glyph.
//   - [NSLayoutManager.NumberOfGlyphs]: The number of glyphs in the layout manager.
//   - [NSLayoutManager.PropertyForGlyphAtIndex]: Returns the glyph property of the glyph at the specified index.
//
// # Setting layout information
//
//   - [NSLayoutManager.SetAttachmentSizeForGlyphRange]: Sets the size to use when drawing a glyph that represents an attachment.
//   - [NSLayoutManager.SetDrawsOutsideLineFragmentForGlyphAtIndex]: Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
//   - [NSLayoutManager.SetExtraLineFragmentRectUsedRectTextContainer]: Sets the bounds and container for the extra line fragment.
//   - [NSLayoutManager.SetLineFragmentRectForGlyphRangeUsedRect]: Associates the line fragment bounds for the specified range of glyphs.
//   - [NSLayoutManager.SetLocationForStartOfGlyphRange]: Sets the location for the first glyph in the specified range.
//   - [NSLayoutManager.SetNotShownAttributeForGlyphAtIndex]: Sets the visibility of the glyph at the specified index.
//
// # Getting layout information
//
//   - [NSLayoutManager.AttachmentSizeForGlyphAtIndex]: Returns the size of the attachment glyph at the specified index.
//   - [NSLayoutManager.DrawsOutsideLineFragmentForGlyphAtIndex]: Indicates whether the glyph draws outside its line fragment rectangle.
//   - [NSLayoutManager.ExtraLineFragmentRect]: The rectangle for the extra line fragment at the end of a document.
//   - [NSLayoutManager.ExtraLineFragmentTextContainer]: The text container for the extra line fragment rectangle.
//   - [NSLayoutManager.ExtraLineFragmentUsedRect]: The rectangle that encloses the insertion point in the extra line fragment rectangle.
//   - [NSLayoutManager.FirstUnlaidCharacterIndex]: Returns the index for the first character in the layout manager that isn’t in the layout.
//   - [NSLayoutManager.FirstUnlaidGlyphIndex]: Returns the index for the first glyph in the layout manager that isn’t in the layout.
//   - [NSLayoutManager.GetFirstUnlaidCharacterIndexGlyphIndex]: Returns the indexes for the first character and glyph that have invalid layout information.
//   - [NSLayoutManager.LineFragmentRectForGlyphAtIndexEffectiveRange]: Returns the rectangle for the line fragment where the glyph lies and (optionally), by reference, the entire range of glyphs in that fragment.
//   - [NSLayoutManager.LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the line fragment rectangle that contains the glyph at the specified glyph index.
//   - [NSLayoutManager.LineFragmentUsedRectForGlyphAtIndexEffectiveRange]: Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//   - [NSLayoutManager.LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//   - [NSLayoutManager.LocationForGlyphAtIndex]: Returns the location for the specified glyph within its line fragment.
//   - [NSLayoutManager.NotShownAttributeForGlyphAtIndex]: Indicates whether the glyph at the specified index has a visible representation.
//   - [NSLayoutManager.TruncatedGlyphRangeInLineFragmentForGlyphAtIndex]: Returns the range of truncated glyphs for a line fragment that contains the specified index.
//
// # Performing advanced layout queries
//
//   - [NSLayoutManager.BoundingRectForGlyphRangeInTextContainer]: Returns the bounding rectangle for the specified glyphs in a container.
//   - [NSLayoutManager.CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints]: Returns the index of the character that lies beneath the specified point using the specified container’s coordinate system.
//   - [NSLayoutManager.CharacterRangeForGlyphRangeActualGlyphRange]: Returns the range of characters that correspond to the glyphs in the specified glyph range.
//   - [NSLayoutManager.EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock]: Enumerates enclosing rectangles for the specified glyph range in a text container.
//   - [NSLayoutManager.EnumerateLineFragmentsForGlyphRangeUsingBlock]: Enumerates line fragments intersecting with the specified glyph range.
//   - [NSLayoutManager.FractionOfDistanceThroughGlyphForPointInTextContainer]: Returns the fraction of the distance between the glyph at the specified point and the next glyph.
//   - [NSLayoutManager.GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes]: Returns insertion points in bulk for a specified line fragment.
//   - [NSLayoutManager.GlyphIndexForPointInTextContainer]: Returns the index of the glyph at the specified location in a text container.
//   - [NSLayoutManager.GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph]: Returns the index of the glyph at the specified point using the container’s coordinate system.
//   - [NSLayoutManager.GlyphRangeForBoundingRectInTextContainer]: Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//   - [NSLayoutManager.GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer]: Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//   - [NSLayoutManager.GlyphRangeForTextContainer]: Returns the range of glyphs lying within the specified text container.
//   - [NSLayoutManager.GlyphRangeForCharacterRangeActualCharacterRange]: Returns the range of glyphs that the specified range of characters generates.
//   - [NSLayoutManager.RangeOfNominallySpacedGlyphsContainingIndex]: Returns the range of displayable glyphs that surround the glyph at the specified index.
//
// # Drawing
//
//   - [NSLayoutManager.DrawBackgroundForGlyphRangeAtPoint]: Draws background marks for the specified glyphs, which must lie completely within a single text container.
//   - [NSLayoutManager.DrawGlyphsForGlyphRangeAtPoint]: Draws the specified glyphs, which must lie completely within a single text container.
//   - [NSLayoutManager.DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Draws a strikethrough for the specified glyphs.
//   - [NSLayoutManager.DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Draws underlining for the glyphs in a specified range.
//   - [NSLayoutManager.FillBackgroundRectArrayCountForCharacterRangeColor]: Fills background rectangles with a color.
//   - [NSLayoutManager.ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext]: Renders the glyphs at the specified positions, using the specified attributes.
//   - [NSLayoutManager.StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Calculates and draws strikethrough for the specified glyphs.
//   - [NSLayoutManager.UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Calculates subranges to underline for the specified glyphs and draws the underlining as appropriate.
//
// # Handling layout for text blocks
//
//   - [NSLayoutManager.SetLayoutRectForTextBlockGlyphRange]: Sets the layout rectangle that encloses the specified text block and glyph range.
//   - [NSLayoutManager.LayoutRectForTextBlockGlyphRange]: Returns the rectangle for the layout of the specified text block and glyph range.
//   - [NSLayoutManager.SetBoundsRectForTextBlockGlyphRange]: Sets the bounding rectangle that encloses the specified text block and glyph range.
//   - [NSLayoutManager.BoundsRectForTextBlockGlyphRange]: Returns the bounding rectangle that encloses the specified text block and glyph range.
//   - [NSLayoutManager.LayoutRectForTextBlockAtIndexEffectiveRange]: Returns the rectangle for the layout of the specified text block and glyph.
//   - [NSLayoutManager.BoundsRectForTextBlockAtIndexEffectiveRange]: Returns the bounding rectangle for the specified text block and glyph.
//
// # Managing attachments
//
//   - [NSLayoutManager.DefaultAttachmentScaling]: The default amount of scaling to apply when an attachment image is too large to fit in a text container.
//   - [NSLayoutManager.SetDefaultAttachmentScaling]
//   - [NSLayoutManager.ShowAttachmentCellInRectCharacterIndex]: Draws an attachment cell.
//
// # Handling Rulers
//
//   - [NSLayoutManager.RulerAccessoryViewForTextViewParagraphStyleRulerEnabled]: Returns the accessory view that the text system uses for its ruler.
//   - [NSLayoutManager.RulerMarkersForTextViewParagraphStyleRuler]: Returns an array of text ruler objects for the current selection.
//
// # Managing the responder chain
//
//   - [NSLayoutManager.LayoutManagerOwnsFirstResponderInWindow]: Indicates whether the first responder in the specified window is a text view for the layout manager.
//   - [NSLayoutManager.FirstTextView]: The first text view in the layout manager’s series of text views.
//   - [NSLayoutManager.TextViewForBeginningOfSelection]: The text view that contains the first glyph in the selection.
//
// # Managing the typesetter
//
//   - [NSLayoutManager.Typesetter]: The current typesetter.
//   - [NSLayoutManager.SetTypesetter]
//   - [NSLayoutManager.TypesetterBehavior]: The default typesetter behavior.
//   - [NSLayoutManager.SetTypesetterBehavior]
//   - [NSLayoutManager.DefaultLineHeightForFont]: Returns the default line height for a line of text that uses a specified font.
//   - [NSLayoutManager.DefaultBaselineOffsetForFont]: Returns the default baseline offset that the layout manager’s typesetter uses for the specified font.
//
// # Managing temporary attribute support
//
//   - [NSLayoutManager.AddTemporaryAttributesForCharacterRange]: Appends one or more temporary attributes to the attributes dictionary of the specified character range.
//   - [NSLayoutManager.AddTemporaryAttributeValueForCharacterRange]: Adds a temporary attribute to the characters in the specified range.
//   - [NSLayoutManager.SetTemporaryAttributesForCharacterRange]: Sets one or more temporary attributes for the specified character range.
//   - [NSLayoutManager.RemoveTemporaryAttributeForCharacterRange]: Removes a temporary attribute from the list of attributes for the specified character range.
//   - [NSLayoutManager.TemporaryAttributeAtCharacterIndexEffectiveRange]: Returns the value for the temporary attribute of a character, and the range it applies to.
//   - [NSLayoutManager.TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange]: Returns the value for the temporary attribute of a character, and the maximum range it applies to.
//   - [NSLayoutManager.TemporaryAttributesAtCharacterIndexEffectiveRange]: Returns the dictionary of temporary attributes for the specified character range.
//   - [NSLayoutManager.TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange]: Returns the temporary attributes for a character, and the maximum range they apply to.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager
type NSLayoutManager struct {
	objectivec.Object
}

// NSLayoutManagerFromID constructs a [NSLayoutManager] from an objc.ID.
//
// An object that coordinates the layout and display of text characters.
func NSLayoutManagerFromID(id objc.ID) NSLayoutManager {
	return NSLayoutManager{objectivec.Object{ID: id}}
}

// NOTE: NSLayoutManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLayoutManager] class.
//
// # Creating a layout manager
//
//   - [INSLayoutManager.InitWithCoder]: Creates a layout manager from data in an unarchiver.
//
// # Managing the layout process
//
//   - [INSLayoutManager.Delegate]: The layout manager’s delegate.
//   - [INSLayoutManager.SetDelegate]
//
// # Accessing the text storage
//
//   - [INSLayoutManager.TextStorage]: The text storage object that contains the content to lay out.
//   - [INSLayoutManager.SetTextStorage]
//   - [INSLayoutManager.ReplaceTextStorage]: Replaces the layout manager’s current text storage object with the specified object.
//
// # Configuring the global layout manager options
//
//   - [INSLayoutManager.AllowsNonContiguousLayout]: A Boolean value that indicates whether the layout manager allows noncontiguous layout.
//   - [INSLayoutManager.SetAllowsNonContiguousLayout]
//   - [INSLayoutManager.HasNonContiguousLayout]: A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
//   - [INSLayoutManager.ShowsInvisibleCharacters]: A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//   - [INSLayoutManager.SetShowsInvisibleCharacters]
//   - [INSLayoutManager.ShowsControlCharacters]: A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//   - [INSLayoutManager.SetShowsControlCharacters]
//   - [INSLayoutManager.UsesFontLeading]: A Boolean value that indicates whether the layout manager uses the leading of the font.
//   - [INSLayoutManager.SetUsesFontLeading]
//   - [INSLayoutManager.BackgroundLayoutEnabled]: A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
//   - [INSLayoutManager.SetBackgroundLayoutEnabled]
//   - [INSLayoutManager.LimitsLayoutForSuspiciousContents]: A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//   - [INSLayoutManager.SetLimitsLayoutForSuspiciousContents]
//   - [INSLayoutManager.UsesDefaultHyphenation]: A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//   - [INSLayoutManager.SetUsesDefaultHyphenation]
//
// # Managing the text containers
//
//   - [INSLayoutManager.TextContainers]: The current text containers of the layout manager.
//   - [INSLayoutManager.AddTextContainer]: Appends the specified text container to the series of text containers where the layout manager arranges text.
//   - [INSLayoutManager.InsertTextContainerAtIndex]: Inserts a text container at the specified index in the list of text containers.
//   - [INSLayoutManager.RemoveTextContainerAtIndex]: Removes the text container at the specified index and invalidates the layout as necessary.
//   - [INSLayoutManager.SetTextContainerForGlyphRange]: Associates a text container with the specified range of glyphs.
//   - [INSLayoutManager.TextContainerChangedGeometry]: Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
//   - [INSLayoutManager.TextContainerChangedTextView]: Updates the information necessary to manage text view objects for the specified text container.
//   - [INSLayoutManager.TextContainerForGlyphAtIndexEffectiveRange]: Returns the text container that manages the layout for the specified glyph, causing layout to happen as necessary.
//   - [INSLayoutManager.TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the text container that manages the layout for the specified glyph.
//   - [INSLayoutManager.UsedRectForTextContainer]: Returns the bounding rectangle for the glyphs in the specified text container.
//
// # Invalidating glyphs and layout
//
//   - [INSLayoutManager.InvalidateDisplayForCharacterRange]: Invalidates display for the specified character range.
//   - [INSLayoutManager.InvalidateDisplayForGlyphRange]: Invalidates a range of glyphs, requiring new layout information, and updates the appropriate regions of any text views that display those glyphs.
//   - [INSLayoutManager.InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange]: Invalidates and adjusts the glyphs in the specified character range.
//   - [INSLayoutManager.InvalidateLayoutForCharacterRangeActualCharacterRange]: Invalidates the layout information for the glyphs that map to the specified character range.
//   - [INSLayoutManager.ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange]: Notifies the layout manager when an edit action changes the contents of its text storage object.
//
// # Causing glyph generation and layout
//
//   - [INSLayoutManager.EnsureGlyphsForCharacterRange]: Forces the layout manager to generate glyphs for the specified character range if it hasn’t already.
//   - [INSLayoutManager.EnsureGlyphsForGlyphRange]: Forces the layout manager to generate glyphs for the specified glyph range if it hasn’t already.
//   - [INSLayoutManager.EnsureLayoutForBoundingRectInTextContainer]: Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
//   - [INSLayoutManager.EnsureLayoutForCharacterRange]: Forces the layout manager to perform layout for the specified character range if it hasn’t already.
//   - [INSLayoutManager.EnsureLayoutForGlyphRange]: Forces the layout manager to perform layout for the specified glyph range if it hasn’t already.
//   - [INSLayoutManager.EnsureLayoutForTextContainer]: Forces the layout manager to perform layout for the specified text container if it hasn’t already.
//   - [INSLayoutManager.GlyphGenerator]: The glyph generator that the layout manager uses.
//   - [INSLayoutManager.SetGlyphGenerator]
//
// # Accessing glyphs
//
//   - [INSLayoutManager.GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels]: Fills a passed-in buffer with a sequence of glyphs.
//   - [INSLayoutManager.CGGlyphAtIndex]: Returns the glyph at the specified index.
//   - [INSLayoutManager.CGGlyphAtIndexIsValidIndex]: Returns the glyph at the specified index along with information about whether the glyph index is valid.
//   - [INSLayoutManager.SetGlyphsPropertiesCharacterIndexesFontForGlyphRange]: Stores the initial glyphs and glyph properties for a character range.
//   - [INSLayoutManager.CharacterIndexForGlyphAtIndex]: Returns the index in the text storage for the first character of the specified glyph.
//   - [INSLayoutManager.GlyphIndexForCharacterAtIndex]: Returns the index of the first glyph of the character at the specified index.
//   - [INSLayoutManager.IsValidGlyphIndex]: Indicates whether the specified index refers to a valid glyph.
//   - [INSLayoutManager.NumberOfGlyphs]: The number of glyphs in the layout manager.
//   - [INSLayoutManager.PropertyForGlyphAtIndex]: Returns the glyph property of the glyph at the specified index.
//
// # Setting layout information
//
//   - [INSLayoutManager.SetAttachmentSizeForGlyphRange]: Sets the size to use when drawing a glyph that represents an attachment.
//   - [INSLayoutManager.SetDrawsOutsideLineFragmentForGlyphAtIndex]: Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
//   - [INSLayoutManager.SetExtraLineFragmentRectUsedRectTextContainer]: Sets the bounds and container for the extra line fragment.
//   - [INSLayoutManager.SetLineFragmentRectForGlyphRangeUsedRect]: Associates the line fragment bounds for the specified range of glyphs.
//   - [INSLayoutManager.SetLocationForStartOfGlyphRange]: Sets the location for the first glyph in the specified range.
//   - [INSLayoutManager.SetNotShownAttributeForGlyphAtIndex]: Sets the visibility of the glyph at the specified index.
//
// # Getting layout information
//
//   - [INSLayoutManager.AttachmentSizeForGlyphAtIndex]: Returns the size of the attachment glyph at the specified index.
//   - [INSLayoutManager.DrawsOutsideLineFragmentForGlyphAtIndex]: Indicates whether the glyph draws outside its line fragment rectangle.
//   - [INSLayoutManager.ExtraLineFragmentRect]: The rectangle for the extra line fragment at the end of a document.
//   - [INSLayoutManager.ExtraLineFragmentTextContainer]: The text container for the extra line fragment rectangle.
//   - [INSLayoutManager.ExtraLineFragmentUsedRect]: The rectangle that encloses the insertion point in the extra line fragment rectangle.
//   - [INSLayoutManager.FirstUnlaidCharacterIndex]: Returns the index for the first character in the layout manager that isn’t in the layout.
//   - [INSLayoutManager.FirstUnlaidGlyphIndex]: Returns the index for the first glyph in the layout manager that isn’t in the layout.
//   - [INSLayoutManager.GetFirstUnlaidCharacterIndexGlyphIndex]: Returns the indexes for the first character and glyph that have invalid layout information.
//   - [INSLayoutManager.LineFragmentRectForGlyphAtIndexEffectiveRange]: Returns the rectangle for the line fragment where the glyph lies and (optionally), by reference, the entire range of glyphs in that fragment.
//   - [INSLayoutManager.LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the line fragment rectangle that contains the glyph at the specified glyph index.
//   - [INSLayoutManager.LineFragmentUsedRectForGlyphAtIndexEffectiveRange]: Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//   - [INSLayoutManager.LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout]: Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//   - [INSLayoutManager.LocationForGlyphAtIndex]: Returns the location for the specified glyph within its line fragment.
//   - [INSLayoutManager.NotShownAttributeForGlyphAtIndex]: Indicates whether the glyph at the specified index has a visible representation.
//   - [INSLayoutManager.TruncatedGlyphRangeInLineFragmentForGlyphAtIndex]: Returns the range of truncated glyphs for a line fragment that contains the specified index.
//
// # Performing advanced layout queries
//
//   - [INSLayoutManager.BoundingRectForGlyphRangeInTextContainer]: Returns the bounding rectangle for the specified glyphs in a container.
//   - [INSLayoutManager.CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints]: Returns the index of the character that lies beneath the specified point using the specified container’s coordinate system.
//   - [INSLayoutManager.CharacterRangeForGlyphRangeActualGlyphRange]: Returns the range of characters that correspond to the glyphs in the specified glyph range.
//   - [INSLayoutManager.EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock]: Enumerates enclosing rectangles for the specified glyph range in a text container.
//   - [INSLayoutManager.EnumerateLineFragmentsForGlyphRangeUsingBlock]: Enumerates line fragments intersecting with the specified glyph range.
//   - [INSLayoutManager.FractionOfDistanceThroughGlyphForPointInTextContainer]: Returns the fraction of the distance between the glyph at the specified point and the next glyph.
//   - [INSLayoutManager.GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes]: Returns insertion points in bulk for a specified line fragment.
//   - [INSLayoutManager.GlyphIndexForPointInTextContainer]: Returns the index of the glyph at the specified location in a text container.
//   - [INSLayoutManager.GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph]: Returns the index of the glyph at the specified point using the container’s coordinate system.
//   - [INSLayoutManager.GlyphRangeForBoundingRectInTextContainer]: Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//   - [INSLayoutManager.GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer]: Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//   - [INSLayoutManager.GlyphRangeForTextContainer]: Returns the range of glyphs lying within the specified text container.
//   - [INSLayoutManager.GlyphRangeForCharacterRangeActualCharacterRange]: Returns the range of glyphs that the specified range of characters generates.
//   - [INSLayoutManager.RangeOfNominallySpacedGlyphsContainingIndex]: Returns the range of displayable glyphs that surround the glyph at the specified index.
//
// # Drawing
//
//   - [INSLayoutManager.DrawBackgroundForGlyphRangeAtPoint]: Draws background marks for the specified glyphs, which must lie completely within a single text container.
//   - [INSLayoutManager.DrawGlyphsForGlyphRangeAtPoint]: Draws the specified glyphs, which must lie completely within a single text container.
//   - [INSLayoutManager.DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Draws a strikethrough for the specified glyphs.
//   - [INSLayoutManager.DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Draws underlining for the glyphs in a specified range.
//   - [INSLayoutManager.FillBackgroundRectArrayCountForCharacterRangeColor]: Fills background rectangles with a color.
//   - [INSLayoutManager.ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext]: Renders the glyphs at the specified positions, using the specified attributes.
//   - [INSLayoutManager.StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Calculates and draws strikethrough for the specified glyphs.
//   - [INSLayoutManager.UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin]: Calculates subranges to underline for the specified glyphs and draws the underlining as appropriate.
//
// # Handling layout for text blocks
//
//   - [INSLayoutManager.SetLayoutRectForTextBlockGlyphRange]: Sets the layout rectangle that encloses the specified text block and glyph range.
//   - [INSLayoutManager.LayoutRectForTextBlockGlyphRange]: Returns the rectangle for the layout of the specified text block and glyph range.
//   - [INSLayoutManager.SetBoundsRectForTextBlockGlyphRange]: Sets the bounding rectangle that encloses the specified text block and glyph range.
//   - [INSLayoutManager.BoundsRectForTextBlockGlyphRange]: Returns the bounding rectangle that encloses the specified text block and glyph range.
//   - [INSLayoutManager.LayoutRectForTextBlockAtIndexEffectiveRange]: Returns the rectangle for the layout of the specified text block and glyph.
//   - [INSLayoutManager.BoundsRectForTextBlockAtIndexEffectiveRange]: Returns the bounding rectangle for the specified text block and glyph.
//
// # Managing attachments
//
//   - [INSLayoutManager.DefaultAttachmentScaling]: The default amount of scaling to apply when an attachment image is too large to fit in a text container.
//   - [INSLayoutManager.SetDefaultAttachmentScaling]
//   - [INSLayoutManager.ShowAttachmentCellInRectCharacterIndex]: Draws an attachment cell.
//
// # Handling Rulers
//
//   - [INSLayoutManager.RulerAccessoryViewForTextViewParagraphStyleRulerEnabled]: Returns the accessory view that the text system uses for its ruler.
//   - [INSLayoutManager.RulerMarkersForTextViewParagraphStyleRuler]: Returns an array of text ruler objects for the current selection.
//
// # Managing the responder chain
//
//   - [INSLayoutManager.LayoutManagerOwnsFirstResponderInWindow]: Indicates whether the first responder in the specified window is a text view for the layout manager.
//   - [INSLayoutManager.FirstTextView]: The first text view in the layout manager’s series of text views.
//   - [INSLayoutManager.TextViewForBeginningOfSelection]: The text view that contains the first glyph in the selection.
//
// # Managing the typesetter
//
//   - [INSLayoutManager.Typesetter]: The current typesetter.
//   - [INSLayoutManager.SetTypesetter]
//   - [INSLayoutManager.TypesetterBehavior]: The default typesetter behavior.
//   - [INSLayoutManager.SetTypesetterBehavior]
//   - [INSLayoutManager.DefaultLineHeightForFont]: Returns the default line height for a line of text that uses a specified font.
//   - [INSLayoutManager.DefaultBaselineOffsetForFont]: Returns the default baseline offset that the layout manager’s typesetter uses for the specified font.
//
// # Managing temporary attribute support
//
//   - [INSLayoutManager.AddTemporaryAttributesForCharacterRange]: Appends one or more temporary attributes to the attributes dictionary of the specified character range.
//   - [INSLayoutManager.AddTemporaryAttributeValueForCharacterRange]: Adds a temporary attribute to the characters in the specified range.
//   - [INSLayoutManager.SetTemporaryAttributesForCharacterRange]: Sets one or more temporary attributes for the specified character range.
//   - [INSLayoutManager.RemoveTemporaryAttributeForCharacterRange]: Removes a temporary attribute from the list of attributes for the specified character range.
//   - [INSLayoutManager.TemporaryAttributeAtCharacterIndexEffectiveRange]: Returns the value for the temporary attribute of a character, and the range it applies to.
//   - [INSLayoutManager.TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange]: Returns the value for the temporary attribute of a character, and the maximum range it applies to.
//   - [INSLayoutManager.TemporaryAttributesAtCharacterIndexEffectiveRange]: Returns the dictionary of temporary attributes for the specified character range.
//   - [INSLayoutManager.TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange]: Returns the temporary attributes for a character, and the maximum range they apply to.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager
type INSLayoutManager interface {
	objectivec.IObject

	// Topic: Creating a layout manager

	// Creates a layout manager from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSLayoutManager

	// Topic: Managing the layout process

	// The layout manager’s delegate.
	Delegate() NSLayoutManagerDelegate
	SetDelegate(value NSLayoutManagerDelegate)

	// Topic: Accessing the text storage

	// The text storage object that contains the content to lay out.
	TextStorage() NSTextStorage
	SetTextStorage(value NSTextStorage)
	// Replaces the layout manager’s current text storage object with the specified object.
	ReplaceTextStorage(newTextStorage NSTextStorage)

	// Topic: Configuring the global layout manager options

	// A Boolean value that indicates whether the layout manager allows noncontiguous layout.
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	// A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
	HasNonContiguousLayout() bool
	// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	// A Boolean value that indicates whether the layout manager uses the leading of the font.
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)

	// Topic: Managing the text containers

	// The current text containers of the layout manager.
	TextContainers() []NSTextContainer
	// Appends the specified text container to the series of text containers where the layout manager arranges text.
	AddTextContainer(container INSTextContainer)
	// Inserts a text container at the specified index in the list of text containers.
	InsertTextContainerAtIndex(container INSTextContainer, index uint)
	// Removes the text container at the specified index and invalidates the layout as necessary.
	RemoveTextContainerAtIndex(index uint)
	// Associates a text container with the specified range of glyphs.
	SetTextContainerForGlyphRange(container INSTextContainer, glyphRange foundation.NSRange)
	// Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
	TextContainerChangedGeometry(container INSTextContainer)
	// Updates the information necessary to manage text view objects for the specified text container.
	TextContainerChangedTextView(container INSTextContainer)
	// Returns the text container that manages the layout for the specified glyph, causing layout to happen as necessary.
	TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) INSTextContainer
	// Returns the text container that manages the layout for the specified glyph.
	TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) INSTextContainer
	// Returns the bounding rectangle for the glyphs in the specified text container.
	UsedRectForTextContainer(container INSTextContainer) corefoundation.CGRect

	// Topic: Invalidating glyphs and layout

	// Invalidates display for the specified character range.
	InvalidateDisplayForCharacterRange(charRange foundation.NSRange)
	// Invalidates a range of glyphs, requiring new layout information, and updates the appropriate regions of any text views that display those glyphs.
	InvalidateDisplayForGlyphRange(glyphRange foundation.NSRange)
	// Invalidates and adjusts the glyphs in the specified character range.
	InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.NSRange, delta int, actualCharRange foundation.NSRange)
	// Invalidates the layout information for the glyphs that map to the specified character range.
	InvalidateLayoutForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange)
	// Notifies the layout manager when an edit action changes the contents of its text storage object.
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage NSTextStorage, editMask NSTextStorageEditActions, newCharRange foundation.NSRange, delta int, invalidatedCharRange foundation.NSRange)

	// Topic: Causing glyph generation and layout

	// Forces the layout manager to generate glyphs for the specified character range if it hasn’t already.
	EnsureGlyphsForCharacterRange(charRange foundation.NSRange)
	// Forces the layout manager to generate glyphs for the specified glyph range if it hasn’t already.
	EnsureGlyphsForGlyphRange(glyphRange foundation.NSRange)
	// Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
	EnsureLayoutForBoundingRectInTextContainer(bounds corefoundation.CGRect, container INSTextContainer)
	// Forces the layout manager to perform layout for the specified character range if it hasn’t already.
	EnsureLayoutForCharacterRange(charRange foundation.NSRange)
	// Forces the layout manager to perform layout for the specified glyph range if it hasn’t already.
	EnsureLayoutForGlyphRange(glyphRange foundation.NSRange)
	// Forces the layout manager to perform layout for the specified text container if it hasn’t already.
	EnsureLayoutForTextContainer(container INSTextContainer)
	// The glyph generator that the layout manager uses.
	GlyphGenerator() INSGlyphGenerator
	SetGlyphGenerator(value INSGlyphGenerator)

	// Topic: Accessing glyphs

	// Fills a passed-in buffer with a sequence of glyphs.
	GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange foundation.NSRange, glyphBuffer *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexBuffer unsafe.Pointer, bidiLevelBuffer unsafe.Pointer) uint
	// Returns the glyph at the specified index.
	CGGlyphAtIndex(glyphIndex uint) coregraphics.CGFontIndex
	// Returns the glyph at the specified index along with information about whether the glyph index is valid.
	CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) coregraphics.CGFontIndex
	// Stores the initial glyphs and glyph properties for a character range.
	SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexes unsafe.Pointer, aFont NSFont, glyphRange foundation.NSRange)
	// Returns the index in the text storage for the first character of the specified glyph.
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	// Returns the index of the first glyph of the character at the specified index.
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	// Indicates whether the specified index refers to a valid glyph.
	IsValidGlyphIndex(glyphIndex uint) bool
	// The number of glyphs in the layout manager.
	NumberOfGlyphs() uint
	// Returns the glyph property of the glyph at the specified index.
	PropertyForGlyphAtIndex(glyphIndex uint) NSGlyphProperty

	// Topic: Setting layout information

	// Sets the size to use when drawing a glyph that represents an attachment.
	SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.NSRange)
	// Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
	SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint)
	// Sets the bounds and container for the extra line fragment.
	SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect corefoundation.CGRect, usedRect corefoundation.CGRect, container INSTextContainer)
	// Associates the line fragment bounds for the specified range of glyphs.
	SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect corefoundation.CGRect, glyphRange foundation.NSRange, usedRect corefoundation.CGRect)
	// Sets the location for the first glyph in the specified range.
	SetLocationForStartOfGlyphRange(location corefoundation.CGPoint, glyphRange foundation.NSRange)
	// Sets the visibility of the glyph at the specified index.
	SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint)

	// Topic: Getting layout information

	// Returns the size of the attachment glyph at the specified index.
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) corefoundation.CGSize
	// Indicates whether the glyph draws outside its line fragment rectangle.
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	// The rectangle for the extra line fragment at the end of a document.
	ExtraLineFragmentRect() corefoundation.CGRect
	// The text container for the extra line fragment rectangle.
	ExtraLineFragmentTextContainer() INSTextContainer
	// The rectangle that encloses the insertion point in the extra line fragment rectangle.
	ExtraLineFragmentUsedRect() corefoundation.CGRect
	// Returns the index for the first character in the layout manager that isn’t in the layout.
	FirstUnlaidCharacterIndex() uint
	// Returns the index for the first glyph in the layout manager that isn’t in the layout.
	FirstUnlaidGlyphIndex() uint
	// Returns the indexes for the first character and glyph that have invalid layout information.
	GetFirstUnlaidCharacterIndexGlyphIndex(charIndex unsafe.Pointer, glyphIndex unsafe.Pointer)
	// Returns the rectangle for the line fragment where the glyph lies and (optionally), by reference, the entire range of glyphs in that fragment.
	LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect
	// Returns the line fragment rectangle that contains the glyph at the specified glyph index.
	LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) corefoundation.CGRect
	// Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
	LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect
	// Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
	LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) corefoundation.CGRect
	// Returns the location for the specified glyph within its line fragment.
	LocationForGlyphAtIndex(glyphIndex uint) corefoundation.CGPoint
	// Indicates whether the glyph at the specified index has a visible representation.
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	// Returns the range of truncated glyphs for a line fragment that contains the specified index.
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.NSRange

	// Topic: Performing advanced layout queries

	// Returns the bounding rectangle for the specified glyphs in a container.
	BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.NSRange, container INSTextContainer) corefoundation.CGRect
	// Returns the index of the character that lies beneath the specified point using the specified container’s coordinate system.
	CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point corefoundation.CGPoint, container INSTextContainer, partialFraction unsafe.Pointer) uint
	// Returns the range of characters that correspond to the glyphs in the specified glyph range.
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.NSRange, actualGlyphRange foundation.NSRange) foundation.NSRange
	// Enumerates enclosing rectangles for the specified glyph range in a text container.
	EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.NSRange, selectedRange foundation.NSRange, textContainer INSTextContainer, block func(corefoundation.CGRect, *bool))
	// Enumerates line fragments intersecting with the specified glyph range.
	EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.NSRange, block TextContainerHandler)
	// Returns the fraction of the distance between the glyph at the specified point and the next glyph.
	FractionOfDistanceThroughGlyphForPointInTextContainer(point corefoundation.CGPoint, container INSTextContainer) float64
	// Returns insertion points in bulk for a specified line fragment.
	GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions unsafe.Pointer, charIndexes unsafe.Pointer) uint
	// Returns the index of the glyph at the specified location in a text container.
	GlyphIndexForPointInTextContainer(point corefoundation.CGPoint, container INSTextContainer) uint
	// Returns the index of the glyph at the specified point using the container’s coordinate system.
	GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point corefoundation.CGPoint, container INSTextContainer, partialFraction unsafe.Pointer) uint
	// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
	GlyphRangeForBoundingRectInTextContainer(bounds corefoundation.CGRect, container INSTextContainer) foundation.NSRange
	// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
	GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds corefoundation.CGRect, container INSTextContainer) foundation.NSRange
	// Returns the range of glyphs lying within the specified text container.
	GlyphRangeForTextContainer(container INSTextContainer) foundation.NSRange
	// Returns the range of glyphs that the specified range of characters generates.
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange) foundation.NSRange
	// Returns the range of displayable glyphs that surround the glyph at the specified index.
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.NSRange

	// Topic: Drawing

	// Draws background marks for the specified glyphs, which must lie completely within a single text container.
	DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.NSRange, origin corefoundation.CGPoint)
	// Draws the specified glyphs, which must lie completely within a single text container.
	DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.NSRange, origin corefoundation.CGPoint)
	// Draws a strikethrough for the specified glyphs.
	DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, strikethroughVal NSUnderlineStyle, baselineOffset float64, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint)
	// Draws underlining for the glyphs in a specified range.
	DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, underlineVal NSUnderlineStyle, baselineOffset float64, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint)
	// Fills background rectangles with a color.
	FillBackgroundRectArrayCountForCharacterRangeColor(rectArray []corefoundation.CGRect, rectCount uint, charRange foundation.NSRange, color INSColor)
	// Renders the glyphs at the specified positions, using the specified attributes.
	ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.CGFontIndex, positions []corefoundation.CGPoint, glyphCount int, font NSFont, textMatrix corefoundation.CGAffineTransform, attributes foundation.INSDictionary, CGContext coregraphics.CGContextRef)
	// Calculates and draws strikethrough for the specified glyphs.
	StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, strikethroughVal NSUnderlineStyle, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint)
	// Calculates subranges to underline for the specified glyphs and draws the underlining as appropriate.
	UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, underlineVal NSUnderlineStyle, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint)

	// Topic: Handling layout for text blocks

	// Sets the layout rectangle that encloses the specified text block and glyph range.
	SetLayoutRectForTextBlockGlyphRange(rect corefoundation.CGRect, block INSTextBlock, glyphRange foundation.NSRange)
	// Returns the rectangle for the layout of the specified text block and glyph range.
	LayoutRectForTextBlockGlyphRange(block INSTextBlock, glyphRange foundation.NSRange) corefoundation.CGRect
	// Sets the bounding rectangle that encloses the specified text block and glyph range.
	SetBoundsRectForTextBlockGlyphRange(rect corefoundation.CGRect, block INSTextBlock, glyphRange foundation.NSRange)
	// Returns the bounding rectangle that encloses the specified text block and glyph range.
	BoundsRectForTextBlockGlyphRange(block INSTextBlock, glyphRange foundation.NSRange) corefoundation.CGRect
	// Returns the rectangle for the layout of the specified text block and glyph.
	LayoutRectForTextBlockAtIndexEffectiveRange(block INSTextBlock, glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect
	// Returns the bounding rectangle for the specified text block and glyph.
	BoundsRectForTextBlockAtIndexEffectiveRange(block INSTextBlock, glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect

	// Topic: Managing attachments

	// The default amount of scaling to apply when an attachment image is too large to fit in a text container.
	DefaultAttachmentScaling() NSImageScaling
	SetDefaultAttachmentScaling(value NSImageScaling)
	// Draws an attachment cell.
	ShowAttachmentCellInRectCharacterIndex(cell INSCell, rect corefoundation.CGRect, attachmentIndex uint)

	// Topic: Handling Rulers

	// Returns the accessory view that the text system uses for its ruler.
	RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view INSTextView, style INSParagraphStyle, ruler INSRulerView, isEnabled bool) INSView
	// Returns an array of text ruler objects for the current selection.
	RulerMarkersForTextViewParagraphStyleRuler(view INSTextView, style INSParagraphStyle, ruler INSRulerView) []NSRulerMarker

	// Topic: Managing the responder chain

	// Indicates whether the first responder in the specified window is a text view for the layout manager.
	LayoutManagerOwnsFirstResponderInWindow(window INSWindow) bool
	// The first text view in the layout manager’s series of text views.
	FirstTextView() INSTextView
	// The text view that contains the first glyph in the selection.
	TextViewForBeginningOfSelection() INSTextView

	// Topic: Managing the typesetter

	// The current typesetter.
	Typesetter() INSTypesetter
	SetTypesetter(value INSTypesetter)
	// The default typesetter behavior.
	TypesetterBehavior() NSTypesetterBehavior
	SetTypesetterBehavior(value NSTypesetterBehavior)
	// Returns the default line height for a line of text that uses a specified font.
	DefaultLineHeightForFont(theFont NSFont) float64
	// Returns the default baseline offset that the layout manager’s typesetter uses for the specified font.
	DefaultBaselineOffsetForFont(theFont NSFont) float64

	// Topic: Managing temporary attribute support

	// Appends one or more temporary attributes to the attributes dictionary of the specified character range.
	AddTemporaryAttributesForCharacterRange(attrs foundation.INSDictionary, charRange foundation.NSRange)
	// Adds a temporary attribute to the characters in the specified range.
	AddTemporaryAttributeValueForCharacterRange(attrName foundation.NSString, value objectivec.IObject, charRange foundation.NSRange)
	// Sets one or more temporary attributes for the specified character range.
	SetTemporaryAttributesForCharacterRange(attrs foundation.INSDictionary, charRange foundation.NSRange)
	// Removes a temporary attribute from the list of attributes for the specified character range.
	RemoveTemporaryAttributeForCharacterRange(attrName foundation.NSString, charRange foundation.NSRange)
	// Returns the value for the temporary attribute of a character, and the range it applies to.
	TemporaryAttributeAtCharacterIndexEffectiveRange(attrName foundation.NSString, location uint, range_ foundation.NSRange) objectivec.IObject
	// Returns the value for the temporary attribute of a character, and the maximum range it applies to.
	TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.NSString, location uint, range_ foundation.NSRange, rangeLimit foundation.NSRange) objectivec.IObject
	// Returns the dictionary of temporary attributes for the specified character range.
	TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange foundation.NSRange) foundation.INSDictionary
	// Returns the temporary attributes for a character, and the maximum range they apply to.
	TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ foundation.NSRange, rangeLimit foundation.NSRange) foundation.INSDictionary

	// Returns the glyph at the specified index.
	GlyphAtIndex(glyphIndex uint) NSGlyph
	// Returns the glyph at a specified index, and optionally returns a flag indicating whether the requested index is valid.
	GlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) NSGlyph
	// Returns an array of rectangles and, by reference, the number of such rectangles, that define the region in the given container enclosing the given character range.
	RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount(charRange foundation.NSRange, selCharRange foundation.NSRange, container INSTextContainer, rectCount unsafe.Pointer) foundation.NSRect
	// Returns an array of rectangles and, by reference, the number of such rectangles, that define the region in the given container enclosing the given glyph range.
	RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount(glyphRange foundation.NSRange, selGlyphRange foundation.NSRange, container INSTextContainer, rectCount unsafe.Pointer) foundation.NSRect
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (l NSLayoutManager) Init() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutManager) Autorelease() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutManager creates a new NSLayoutManager instance.
func NewNSLayoutManager() NSLayoutManager {
	class := getNSLayoutManagerClass()
	rv := objc.Send[NSLayoutManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a layout manager from data in an unarchiver.
//
// coder: A coder that implements [NSCoder].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/init(coder:)
//
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
func NewLayoutManagerWithCoder(coder foundation.INSCoder) NSLayoutManager {
	instance := getNSLayoutManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSLayoutManagerFromID(rv)
}

// Creates a layout manager from data in an unarchiver.
//
// coder: A coder that implements [NSCoder].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/init(coder:)
//
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
func (l NSLayoutManager) InitWithCoder(coder foundation.INSCoder) NSLayoutManager {
	rv := objc.Send[NSLayoutManager](l.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Replaces the layout manager’s current text storage object with the
// specified object.
//
// newTextStorage: The text storage object to set.
//
// # Discussion
//
// Use this method to update the text storage uniformly for a group of related
// layout manager objects. Unlike changing the value in the textStorage
// property, this method replaces the text storage for all [NSLayoutManager]
// objects that share the current layout manager’s [NSTextStorage] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/replaceTextStorage(_:)
func (l NSLayoutManager) ReplaceTextStorage(newTextStorage NSTextStorage) {
	objc.Send[objc.ID](l.ID, objc.Sel("replaceTextStorage:"), newTextStorage)
}

// Appends the specified text container to the series of text containers where
// the layout manager arranges text.
//
// container: The text container to append.
//
// # Discussion
//
// Invalidates glyphs and layout as needed, but doesn’t perform glyph
// generation or layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTextContainer(_:)
func (l NSLayoutManager) AddTextContainer(container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("addTextContainer:"), container)
}

// Inserts a text container at the specified index in the list of text
// containers.
//
// container: The text container to insert.
//
// index: The index in the series of text containers at which to insert
// `aTextContainer`.
//
// # Discussion
//
// This method invalidates layout for all subsequent [NSTextContainer]
// objects, and invalidates glyph information as needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/insertTextContainer(_:at:)
func (l NSLayoutManager) InsertTextContainerAtIndex(container INSTextContainer, index uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("insertTextContainer:atIndex:"), container, index)
}

// Removes the text container at the specified index and invalidates the
// layout as necessary.
//
// index: The index of the text container to remove.
//
// # Discussion
//
// This method invalidates glyph information as needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTextContainer(at:)
func (l NSLayoutManager) RemoveTextContainerAtIndex(index uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("removeTextContainerAtIndex:"), index)
}

// Associates a text container with the specified range of glyphs.
//
// container: The text container to set.
//
// glyphRange: The range of glyphs to lay out.
//
// # Discussion
//
// The layout within the container is specified with the
// [SetLineFragmentRectForGlyphRangeUsedRect] and
// [SetLocationForStartOfGlyphRange] methods.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setTextContainer(_:forGlyphRange:)
func (l NSLayoutManager) SetTextContainerForGlyphRange(container INSTextContainer, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setTextContainer:forGlyphRange:"), container, glyphRange)
}

// Invalidates the layout information, and possibly glyphs, for the specified
// text container and all subsequent text container objects.
//
// container: The text container whose layout is invalidated.
//
// # Discussion
//
// This method is invoked automatically by other components of the text
// system; you should rarely need to invoke it directly. Subclasses of
// [NSTextContainer], however, must invoke this method any time their size of
// shape changes (a text container that dynamically adjusts its shape to wrap
// text around placed graphics, for example, must do so when a graphic is
// added, moved, or removed).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainerChangedGeometry(_:)
func (l NSLayoutManager) TextContainerChangedGeometry(container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("textContainerChangedGeometry:"), container)
}

// Updates the information necessary to manage text view objects for the
// specified text container.
//
// container: The text container whose text view has changed.
//
// # Discussion
//
// This method is called by a text container, whenever its text view changes,
// to keep notifications synchronized. You should rarely need to invoke it
// directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainerChangedTextView(_:)
func (l NSLayoutManager) TextContainerChangedTextView(container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("textContainerChangedTextView:"), container)
}

// Returns the text container that manages the layout for the specified glyph,
// causing layout to happen as necessary.
//
// glyphIndex: Index of a glyph in the returned container.
//
// effectiveGlyphRange: If not [NULL], on output, points to the whole range of glyphs that are in
// the returned container.
//
// # Return Value
//
// The text container in which the glyph at `glyphIndex` is laid out.
//
// # Discussion
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment. If noncontiguous layout is not
// enabled and `effectiveGlyphRange` is not [NULL], this method additionally
// causes glyph generation and layout for the entire text container containing
// the specified glyph.
//
// Overriding this method is not recommended. Any changes to the returned
// glyph range should be done at the typesetter level.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainer(forGlyphAt:effectiveRange:)
func (l NSLayoutManager) TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) INSTextContainer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("textContainerForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return NSTextContainerFromID(rv)
}

// Returns the text container that manages the layout for the specified glyph.
//
// glyphIndex: Index of a glyph in the returned container.
//
// effectiveGlyphRange: If not [NULL], on output, points to the whole range of glyphs that are in
// the returned container.
//
// flag: If true, glyph generation and layout are not performed, so this option
// should not be used unless layout is known to be complete for the range in
// question, or unless noncontiguous layout is enabled; if false, both are
// performed as needed.
//
// # Return Value
//
// The text container in which the glyph at `glyphIndex` is laid out.
//
// # Discussion
//
// This method is primarily for use from within [NSTypesetter], after layout
// is complete for the range in question, but before the layout manager’s
// call to [NSTypesetter] has returned. In that case glyph and layout holes
// have not yet been recalculated, so the layout manager does not yet know
// that layout is complete for that range, and this variant must be used.
//
// Overriding this method is not recommended. Any changes to the returned
// glyph range should be done at the typesetter level.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainer(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l NSLayoutManager) TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) INSTextContainer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("textContainerForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return NSTextContainerFromID(rv)
}

// Returns the bounding rectangle for the glyphs in the specified text
// container.
//
// # Discussion
//
// Returns the text container’s currently used area, which determines the
// size that the view would need to be in order to display all the glyphs that
// are currently laid out in the container. This causes neither glyph
// generation nor layout.
//
// Used rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usedRect(for:)
func (l NSLayoutManager) UsedRectForTextContainer(container INSTextContainer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("usedRectForTextContainer:"), container)
	return corefoundation.CGRect(rv)
}

// Invalidates display for the specified character range.
//
// charRange: The character range for which display is invalidated.
//
// # Discussion
//
// Parts of the range that are not laid out are remembered and redisplayed
// later when the layout is available. Does not actually cause layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateDisplay(forCharacterRange:)
func (l NSLayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("invalidateDisplayForCharacterRange:"), charRange)
}

// Invalidates a range of glyphs, requiring new layout information, and
// updates the appropriate regions of any text views that display those
// glyphs.
//
// glyphRange: The range of glyphs to invalidate.
//
// # Discussion
//
// You should rarely need to invoke this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateDisplay(forGlyphRange:)
func (l NSLayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("invalidateDisplayForGlyphRange:"), glyphRange)
}

// Invalidates and adjusts the glyphs in the specified character range.
//
// charRange: The range of characters for which to invalidate glyphs.
//
// delta: The number of characters added or removed.
//
// actualCharRange: If not [NULL], on output, the actual range invalidated after any necessary
// expansion. This range can be larger than the range of characters given due
// to the effect of context on glyphs and layout.
//
// # Discussion
//
// This method invalidates the cached glyphs for the characters in the given
// character range, adjusts the character indices of all the subsequent glyphs
// by the change in length, and invalidates the new character range. This
// method invalidates only glyph information and performs no glyph generation
// or layout. Because invalidating glyphs also invalidates layout, after
// invoking this method you should also invoke
// [InvalidateLayoutForCharacterRangeActualCharacterRange], passing
// `charRange` as the first argument.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateGlyphs(forCharacterRange:changeInLength:actualCharacterRange:)
func (l NSLayoutManager) InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.NSRange, delta int, actualCharRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:"), charRange, delta, actualCharRange)
}

// Invalidates the layout information for the glyphs that map to the specified
// character range.
//
// charRange: The range of characters to invalidate.
//
// actualCharRange: If not [NULL], on output, the actual range invalidated after any necessary
// expansion.
//
// # Discussion
//
// This method has the same effect as
// [invalidateLayout(forCharacterRange:isSoft:actualCharacterRange:)] with
// `flag` set to false.
//
// This method only invalidates information; it performs no glyph generation
// or layout. You should rarely need to invoke this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateLayout(forCharacterRange:actualCharacterRange:)
//
// [invalidateLayout(forCharacterRange:isSoft:actualCharacterRange:)]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateLayout(forCharacterRange:isSoft:actualCharacterRange:)
func (l NSLayoutManager) InvalidateLayoutForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("invalidateLayoutForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
}

// Notifies the layout manager when an edit action changes the contents of its
// text storage object.
//
// textStorage: The text storage object processing edits.
//
// editMask: The types of edits done: [NSTextStorageEditedAttributes],
// [NSTextStorageEditedCharacters], or both.
//
// newCharRange: The range in the final string that was explicitly edited.
//
// delta: The length delta for the editing changes.
//
// invalidatedCharRange: The range of characters that changed as a result of attribute fixing. This
// invalidated range is either equal to `newCharRange` or larger.
//
// # Discussion
//
// The [ProcessEditing] method of [NSTextStorage] calls this method to notify
// the layout manager of an edit action. Layout managers must not change the
// contents of the text storage during the execution of this message.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/processEditing(for:edited:range:changeInLength:invalidatedRange:)
func (l NSLayoutManager) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage NSTextStorage, editMask NSTextStorageEditActions, newCharRange foundation.NSRange, delta int, invalidatedCharRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}

// Forces the layout manager to generate glyphs for the specified character
// range if it hasn’t already.
//
// charRange: The character range for which glyphs are generated.
//
// # Discussion
//
// The layout manager reserves the right to perform glyph generation for
// larger ranges. If noncontiguous layout is disabled, then the affected range
// is always effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureGlyphs(forCharacterRange:)
func (l NSLayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureGlyphsForCharacterRange:"), charRange)
}

// Forces the layout manager to generate glyphs for the specified glyph range
// if it hasn’t already.
//
// glyphRange: The glyph range for which glyphs are generated.
//
// # Discussion
//
// The layout manager reserves the right to perform glyph generation for
// larger ranges. If noncontiguous layout is disabled, then the affected range
// is always effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureGlyphs(forGlyphRange:)
func (l NSLayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureGlyphsForGlyphRange:"), glyphRange)
}

// Forces the layout manager to perform layout for the specified area in the
// specified text container if it hasn’t already.
//
// bounds: The area for which layout is performed.
//
// container: The text container containing the area for which layout is performed.
//
// # Discussion
//
// The layout manager reserves the right to perform layout for larger ranges.
// If noncontiguous layout is disabled, then the affected range is always
// effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forBoundingRect:in:)
func (l NSLayoutManager) EnsureLayoutForBoundingRectInTextContainer(bounds corefoundation.CGRect, container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureLayoutForBoundingRect:inTextContainer:"), bounds, container)
}

// Forces the layout manager to perform layout for the specified character
// range if it hasn’t already.
//
// charRange: The character range for which layout is performed.
//
// # Discussion
//
// The layout manager reserves the right to perform layout for larger ranges.
// If noncontiguous layout is disabled, then the affected range is always
// effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forCharacterRange:)
func (l NSLayoutManager) EnsureLayoutForCharacterRange(charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureLayoutForCharacterRange:"), charRange)
}

// Forces the layout manager to perform layout for the specified glyph range
// if it hasn’t already.
//
// glyphRange: The glyph range for which layout is performed.
//
// # Discussion
//
// The layout manager reserves the right to perform layout for larger ranges.
// If noncontiguous layout is disabled, then the affected range is always
// effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forGlyphRange:)
func (l NSLayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureLayoutForGlyphRange:"), glyphRange)
}

// Forces the layout manager to perform layout for the specified text
// container if it hasn’t already.
//
// container: The text container for which layout is performed.
//
// # Discussion
//
// The layout manager reserves the right to perform layout for larger ranges.
// If noncontiguous layout is disabled, then the affected range is always
// effectively extended to start at the beginning of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(for:)
func (l NSLayoutManager) EnsureLayoutForTextContainer(container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("ensureLayoutForTextContainer:"), container)
}

// Fills a passed-in buffer with a sequence of glyphs.
//
// glyphRange: The range of glyphs to fill in.
//
// glyphBuffer: On output, the sequence of glyphs in the given glyph range.
//
// props: If not [NULL], on output, the glyph properties corresponding to the
// filled-in glyphs.
//
// charIndexBuffer: If not [NULL], on output, the indexes of the original characters
// corresponding to the given glyph range. Note that a glyph at index 1 is not
// necessarily mapped to the character at index 1, since a glyph may be for a
// ligature or accent.
//
// bidiLevelBuffer: If not [NULL], on output, the direction of each glyph for bidirectional
// text. The values range from 0 to 61 as defined by Unicode Standard Annex
// #9. An even value means the glyph goes left-to-right, and an odd value
// means the glyph goes right-to-left.
//
// # Return Value
//
// The number of glyphs returned in `glyphBuffer`.
//
// # Discussion
//
// Each pointer passed in should either be [NULL] or else point to sufficient
// memory to hold `glyphRange.Length()` elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getGlyphs(in:glyphs:properties:characterIndexes:bidiLevels:)
func (l NSLayoutManager) GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange foundation.NSRange, glyphBuffer *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexBuffer unsafe.Pointer, bidiLevelBuffer unsafe.Pointer) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("getGlyphsInRange:glyphs:properties:characterIndexes:bidiLevels:"), glyphRange, glyphBuffer, props, charIndexBuffer, bidiLevelBuffer)
	return rv
}

// Returns the glyph at the specified index.
//
// glyphIndex: The index of the glyph that you want. If the index is out of range, this
// method raises an exception with the error [rangeException].
//
// # Return Value
//
// The glyph at the specified index.
//
// # Discussion
//
// Calling this method generates all of the glyphs (as needed) up to and
// including the glyph at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/cgGlyph(at:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (l NSLayoutManager) CGGlyphAtIndex(glyphIndex uint) coregraphics.CGFontIndex {
	rv := objc.Send[coregraphics.CGFontIndex](l.ID, objc.Sel("CGGlyphAtIndex:"), glyphIndex)
	return coregraphics.CGFontIndex(rv)
}

// Returns the glyph at the specified index along with information about
// whether the glyph index is valid.
//
// glyphIndex: The index of the glyph that you want.
//
// isValidIndex: An optional Boolean variable. On return, the variable is set to true if the
// glyph index is valid or false if it is not.
//
// # Return Value
//
// The glyph at the specified index or [kCGFontIndexInvalid] if the index is
// out of range.
//
// # Discussion
//
// If noncontiguous layout is disabled, calling this method generates all
// glyphs up to and including the one at `glyphIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/cgGlyph(at:isValidIndex:)
//
// [kCGFontIndexInvalid]: https://developer.apple.com/documentation/CoreGraphics/kCGFontIndexInvalid
func (l NSLayoutManager) CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) coregraphics.CGFontIndex {
	rv := objc.Send[coregraphics.CGFontIndex](l.ID, objc.Sel("CGGlyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return coregraphics.CGFontIndex(rv)
}

// Stores the initial glyphs and glyph properties for a character range.
//
// glyphs: A pointer to the layout manager’s glyph cache.
//
// props: A pointer to a buffer containing glyph properties for the glyphs in the
// cache.
//
// charIndexes: A pointer to the starting index for the characters in the text storage for
// which glyphs are generated.
//
// aFont: A font to override the font attributes in the text storage for the
// specified character range.
//
// glyphRange: The range of glyphs in the glyph cache to set.
//
// # Discussion
//
// This method is invoked by text system during the glyph generation process.
// The only place apps are allowed to call this method directly is from an
// implementation of the [NSLayoutManagerDelegate] protocol method
// [LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange].
//
// Each array has `glyphRange.Length()` items. The specified `charIndexes`
// must be contiguous (no skipped indexes), enabling multiple items to have a
// same character index (as when one character index generates multiple glyph
// IDs). Due to font substitution, `aFont` passed into this method might not
// match the font in the attributes dictionary. Calling this method for a
// character range that has previously calculated layout information
// invalidates the layout and display.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setGlyphs(_:properties:characterIndexes:font:forGlyphRange:)
func (l NSLayoutManager) SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexes unsafe.Pointer, aFont NSFont, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setGlyphs:properties:characterIndexes:font:forGlyphRange:"), glyphs, props, charIndexes, aFont, glyphRange)
}

// Returns the index in the text storage for the first character of the
// specified glyph.
//
// glyphIndex: The index of the glyph for which to return the associated character.
//
// # Return Value
//
// The index of the first character associated with the glyph at the specified
// index.
//
// # Discussion
//
// If noncontiguous layout is not enabled, this method causes generation of
// all glyphs up to and including `glyphIndex`. This method accepts an index
// beyond the last glyph, returning an index extrapolated from the last actual
// glyph index.
//
// In many cases it’s better to use the range-mapping methods,
// [CharacterRangeForGlyphRangeActualGlyphRange] and
// [GlyphRangeForCharacterRangeActualCharacterRange], which provide more
// comprehensive information.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterIndexForGlyph(at:)
func (l NSLayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("characterIndexForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the index of the first glyph of the character at the specified
// index.
//
// charIndex: The index of the character for which to return the associated glyph.
//
// # Return Value
//
// The index of the first glyph associated with the character at the specified
// index.
//
// # Discussion
//
// If noncontiguous layout is not enabled, this method causes generation of
// all glyphs up to and including those associated with the specified
// character. This method accepts an index beyond the last character,
// returning an index extrapolated from the last actual character index.
//
// In many cases it’s better to use the range-mapping methods,
// [CharacterRangeForGlyphRangeActualGlyphRange] and
// [GlyphRangeForCharacterRangeActualCharacterRange], which provide more
// comprehensive information.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndexForCharacter(at:)
func (l NSLayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("glyphIndexForCharacterAtIndex:"), charIndex)
	return rv
}

// Indicates whether the specified index refers to a valid glyph.
//
// glyphIndex: The index of a glyph in the receiver.
//
// # Return Value
//
// true if the specified `glyphIndex` refers to a valid glyph, otherwise
// false.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/isValidGlyphIndex(_:)
func (l NSLayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isValidGlyphIndex:"), glyphIndex)
	return rv
}

// Returns the glyph property of the glyph at the specified index.
//
// glyphIndex: The glyph whose glyph property is returned.
//
// # Return Value
//
// The glyph property associated with the specified glyph.
// [NSLayoutManager.GlyphProperty] lists the values that can be returned.
//
// # Discussion
//
// If noncontiguous layout is not enabled, this method causes generation of
// all glyphs up to and including the one at `glyphIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/propertyForGlyph(at:)
//
// [NSLayoutManager.GlyphProperty]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty
func (l NSLayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) NSGlyphProperty {
	rv := objc.Send[NSGlyphProperty](l.ID, objc.Sel("propertyForGlyphAtIndex:"), glyphIndex)
	return NSGlyphProperty(rv)
}

// Sets the size to use when drawing a glyph that represents an attachment.
//
// attachmentSize: The glyph size to set.
//
// glyphRange: The attachment glyph’s position in the glyph stream.
//
// # Discussion
//
// For a glyph corresponding to an attachment, this method should be called to
// set the size for the attachment cell to occupy. The glyph’s value should
// be [NSControlGlyph].
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setAttachmentSize(_:forGlyphRange:)
func (l NSLayoutManager) SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

// Indicates whether the specified glyph exceeds the bounds of the line
// fragment for its layout.
//
// flag: If true, sets the given glyph to draw outside its line fragment; if false,
// the glyph does not draw outside.
//
// glyphIndex: Index of the glyph to set.
//
// # Discussion
//
// This can happen when text is set at a fixed line height. For example, if
// the user specifies a fixed line height of 12 points and sets the font size
// to 24 points, the glyphs will exceed their layout rectangles. This
// information is important for determining whether additional lines need to
// be redrawn as a result of changes to any given line fragment.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setDrawsOutsideLineFragment(_:forGlyphAt:)
func (l NSLayoutManager) SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphAtIndex:"), flag, glyphIndex)
}

// Sets the bounds and container for the extra line fragment.
//
// fragmentRect: The rectangle to set.
//
// usedRect: Indicates where the insertion point is drawn.
//
// container: The text container where the rectangle is to be laid out.
//
// # Discussion
//
// The extra line fragment is used when the text backing ends with a hard line
// break or when the text backing is totally empty, to define the extra line
// which needs to be displayed at the end of the text. If the text backing is
// not empty and does not end with a hard line break, this should be set to
// [NSZeroRect] and `nil`.
//
// Line fragment rectangles and line fragment used rectangles are always in
// container coordinates.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setExtraLineFragmentRect(_:usedRect:textContainer:)
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
func (l NSLayoutManager) SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect corefoundation.CGRect, usedRect corefoundation.CGRect, container INSTextContainer) {
	objc.Send[objc.ID](l.ID, objc.Sel("setExtraLineFragmentRect:usedRect:textContainer:"), fragmentRect, usedRect, container)
}

// Associates the line fragment bounds for the specified range of glyphs.
//
// fragmentRect: The rectangle of the line fragment.
//
// glyphRange: The range of glyphs to be associated with `fragmentRect`.
//
// usedRect: The portion of `fragmentRect` that actually contains glyphs or other marks
// that are drawn (including the text container’s line fragment padding.
// Must be equal to or contained within `fragmentRect`.
//
// # Discussion
//
// The typesetter must specify the text container first with
// [SetTextContainerForGlyphRange], and it sets the exact positions of the
// glyphs afterwards with [SetLocationForStartOfGlyphRange].
//
// In the course of layout, all glyphs should end up being included in a range
// passed to this method, but only glyphs that start a new line fragment
// should be at the start of such ranges.
//
// Line fragment rectangles and line fragment used rectangles are always in
// container coordinates.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLineFragmentRect(_:forGlyphRange:usedRect:)
func (l NSLayoutManager) SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect corefoundation.CGRect, glyphRange foundation.NSRange, usedRect corefoundation.CGRect) {
	objc.Send[objc.ID](l.ID, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:"), fragmentRect, glyphRange, usedRect)
}

// Sets the location for the first glyph in the specified range.
//
// location: The location to which the first glyph is set, relative to the origin of the
// glyph’s line fragment origin.
//
// glyphRange: The glyphs whose location is set.
//
// # Discussion
//
// Setting the location for a glyph range implies that its first glyph is not
// nominally spaced with respect to the previous glyph. In the course of
// layout, all glyphs should end up being included in a range passed to this
// method, but only glyphs that start a new nominal range should be at the
// start of such ranges. The first glyph in a line fragment should always
// start a new nominal range. Glyph locations are given relative to their line
// fragment rectangle’s origin.
//
// Before setting the location for a glyph range, you must specify the text
// container with [SetTextContainerForGlyphRange] and the line fragment
// rectangle with [SetLineFragmentRectForGlyphRangeUsedRect].
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLocation(_:forStartOfGlyphRange:)
func (l NSLayoutManager) SetLocationForStartOfGlyphRange(location corefoundation.CGPoint, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setLocation:forStartOfGlyphRange:"), location, glyphRange)
}

// Sets the visibility of the glyph at the specified index.
//
// flag: If true, the glyph is not shown; if false, it is shown.
//
// glyphIndex: Index of the glyph whose attribute is set.
//
// # Discussion
//
// The typesetter decides which glyphs are not shown and sets this attribute
// in the layout manager to ensure that those glyphs are not displayed. For
// example, a tab or newline character doesn’t leave any marks; it just
// indicates where following glyphs are laid out.
//
// Raises an [NSRangeException] if `glyphIndex` is out of bounds.
//
// This method is used by the layout mechanism and should be invoked only
// during typesetting, in almost all cases only by the typesetter. For
// example, a custom typesetter might invoke it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setNotShownAttribute(_:forGlyphAt:)
func (l NSLayoutManager) SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("setNotShownAttribute:forGlyphAtIndex:"), flag, glyphIndex)
}

// Returns the size of the attachment glyph at the specified index.
//
// glyphIndex: The index of the attachment glyph.
//
// # Return Value
//
// The layout manager calls this method for glyphs representing attachments,
// and returns the size that the attachment cell occupies. Returns `{-1.0,
// -1.0}` if there is no attachment laid for the specified glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/attachmentSize(forGlyphAt:)
func (l NSLayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](l.ID, objc.Sel("attachmentSizeForGlyphAtIndex:"), glyphIndex)
	return corefoundation.CGSize(rv)
}

// Indicates whether the glyph draws outside its line fragment rectangle.
//
// glyphIndex: Index of the glyph.
//
// # Return Value
//
// true if the glyph at `glyphIndex` exceeds the bounds of the line fragment
// where it’s laid out, false otherwise.
//
// # Discussion
//
// Exceeding bounds can happen when text is set at a fixed line height. For
// example, if the user specifies a fixed line height of 12 points and sets
// the font size to 24 points, the glyphs will exceed their layout rectangles.
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment.
//
// Glyphs that draw outside their line fragment rectangles aren’t considered
// when calculating enclosing rectangles with the
// [RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount]
// and
// [RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount]
// methods. They are, however, considered by
// [BoundingRectForGlyphRangeInTextContainer].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawsOutsideLineFragment(forGlyphAt:)
func (l NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("drawsOutsideLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the index for the first character in the layout manager that
// isn’t in the layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstUnlaidCharacterIndex()
func (l NSLayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("firstUnlaidCharacterIndex"))
	return rv
}

// Returns the index for the first glyph in the layout manager that isn’t in
// the layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstUnlaidGlyphIndex()
func (l NSLayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("firstUnlaidGlyphIndex"))
	return rv
}

// Returns the indexes for the first character and glyph that have invalid
// layout information.
//
// charIndex: On return, if not [NULL], the index of the first character that has invalid
// layout information
//
// glyphIndex: On return, if not [NULL], the index of the first glyph that has invalid
// layout information.
//
// # Discussion
//
// Either parameter may be [NULL], in which case the receiver simply ignores
// it.
//
// As part of its implementation, this method calls
// [FirstUnlaidCharacterIndex] and [FirstUnlaidGlyphIndex]. To change this
// method’s behavior, override those two methods instead of this one.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getFirstUnlaidCharacterIndex(_:glyphIndex:)
func (l NSLayoutManager) GetFirstUnlaidCharacterIndexGlyphIndex(charIndex unsafe.Pointer, glyphIndex unsafe.Pointer) {
	objc.Send[objc.ID](l.ID, objc.Sel("getFirstUnlaidCharacterIndex:glyphIndex:"), charIndex, glyphIndex)
}

// Returns the rectangle for the line fragment where the glyph lies and
// (optionally), by reference, the entire range of glyphs in that fragment.
//
// glyphIndex: The glyph for which to return the line fragment rectangle.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the line fragment.
//
// # Return Value
//
// The line fragment in which the given glyph is laid out.
//
// # Discussion
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// for all of the text up to and including that line fragment.
//
// Line fragment rectangles are always in container coordinates.
//
// Overriding this method is not recommended. If the line fragment rectangle
// needs to be modified, that should be done at the typesetter level or by
// calling [SetLineFragmentRectForGlyphRangeUsedRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentRect(forGlyphAt:effectiveRange:)
func (l NSLayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("lineFragmentRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return corefoundation.CGRect(rv)
}

// Returns the line fragment rectangle that contains the glyph at the
// specified glyph index.
//
// glyphIndex: The glyph for which to return the line fragment rectangle.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the line fragment.
//
// flag: If true, glyph generation and layout are not performed, so this option
// should not be used unless layout is known to be complete for the range in
// question, or unless noncontiguous layout is enabled; if false, both are
// performed as needed.
//
// # Return Value
//
// The line fragment in which the given glyph is laid out.
//
// # Discussion
//
// This method is primarily for use from within [NSTypesetter], after layout
// is complete for the range in question, but before the layout manager’s
// call to [NSTypesetter] has returned. In that case glyph and layout holes
// have not yet been recalculated, so the layout manager does not yet know
// that layout is complete for that range, and this variant must be used.
//
// Overriding this method is not recommended. If the line fragment rectangle
// needs to be modified, that should be done at the typesetter level or by
// calling [SetLineFragmentRectForGlyphRangeUsedRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentRect(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l NSLayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("lineFragmentRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return corefoundation.CGRect(rv)
}

// Returns the usage rectangle for the line fragment and (optionally) returns
// the entire range of glyphs in that fragment.
//
// glyphIndex: The glyph for which to return the line fragment used rectangle.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the line fragment.
//
// # Return Value
//
// The used rectangle for the line fragment in which the given glyph is laid
// out.
//
// # Discussion
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment.
//
// Line fragment used rectangles are always in container coordinates.
//
// Overriding this method is not recommended. If the line fragment used
// rectangle needs to be modified, that should be done at the typesetter level
// or by calling [SetLineFragmentRectForGlyphRangeUsedRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentUsedRect(forGlyphAt:effectiveRange:)
func (l NSLayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return corefoundation.CGRect(rv)
}

// Returns the usage rectangle for the line fragment and (optionally) returns
// the entire range of glyphs in that fragment.
//
// glyphIndex: The glyph for which to return the line fragment used rectangle.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the line fragment.
//
// flag: If true, glyph generation and layout are not performed, so this option
// should not be used unless layout is known to be complete for the range in
// question, or unless noncontiguous layout is enabled; if false, both are
// performed as needed.
//
// # Return Value
//
// The used rectangle for the line fragment in which the given glyph is laid
// out.
//
// # Discussion
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment.
//
// Line fragment used rectangles are always in container coordinates.
//
// Overriding this method is not recommended. If the line fragment used
// rectangle needs to be modified, that should be done at the typesetter level
// or by calling [SetLineFragmentRectForGlyphRangeUsedRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentUsedRect(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l NSLayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange foundation.NSRange, flag bool) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return corefoundation.CGRect(rv)
}

// Returns the location for the specified glyph within its line fragment.
//
// glyphIndex: The glyph whose location is returned.
//
// # Return Value
//
// The location of the given glyph.
//
// # Discussion
//
// If the given glyph does not have an explicit location set for it (for
// example, if it is part of (but not first in) a sequence of nominally spaced
// characters), the location is calculated by glyph advancements from the
// location of the most recent preceding glyph with a location set.
//
// Glyph locations are relative to their line fragment rectangle’s origin.
// The line fragment rectangle in turn is defined in the coordinate system of
// the text container where it resides.
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/location(forGlyphAt:)
func (l NSLayoutManager) LocationForGlyphAtIndex(glyphIndex uint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l.ID, objc.Sel("locationForGlyphAtIndex:"), glyphIndex)
	return corefoundation.CGPoint(rv)
}

// Indicates whether the glyph at the specified index has a visible
// representation.
//
// glyphIndex: Index of the glyph.
//
// # Return Value
//
// true if the glyph at `glyphIndex` is not shown; otherwise false.
//
// # Discussion
//
// Some glyphs are not shown. For example, a tab, newline, or attachment glyph
// is not shown; it just affects the layout of following glyphs or locates the
// attachment graphic. Space characters, however, typically are shown as
// glyphs with a displacement, although they leave no visible marks.
//
// This method causes glyph generation and layout for the line fragment
// containing the specified glyph, or if noncontiguous layout is not enabled,
// up to and including that line fragment.
//
// Raises an [NSRangeException] if `glyphIndex` is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/notShownAttribute(forGlyphAt:)
func (l NSLayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("notShownAttributeForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the range of truncated glyphs for a line fragment that contains the
// specified index.
//
// glyphIndex: A glyph whose line fragment is tested.
//
// # Return Value
//
// The range of truncated glyphs for a line fragment containing the specified
// glyph, or, when there is no truncation for the line fragment, `{NSNotFound,
// 0}`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/truncatedGlyphRange(inLineFragmentForGlyphAt:)
func (l NSLayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("truncatedGlyphRangeInLineFragmentForGlyphAtIndex:"), glyphIndex)
	return foundation.NSRange(rv)
}

// Returns the bounding rectangle for the specified glyphs in a container.
//
// glyphRange: The range of glyphs for which to return the bounding rectangle.
//
// container: The text container in which the glyphs are laid out.
//
// # Return Value
//
// The bounding rectangle enclosing the given range of glyphs.
//
// # Discussion
//
// This method returns a single bounding rectangle (in container coordinates)
// enclosing all glyphs and other marks drawn in the given text container for
// the given glyph range, including glyphs that draw outside their line
// fragment rectangles and text attributes such as underlining.
//
// The range is intersected with the container’s range before computing the
// bounding rectangle. This method can be used to translate glyph ranges into
// display rectangles for invalidation and redrawing when a range of glyphs
// changes. Bounding rectangles are always in container coordinates.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundingRect(forGlyphRange:in:)
func (l NSLayoutManager) BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.NSRange, container INSTextContainer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("boundingRectForGlyphRange:inTextContainer:"), glyphRange, container)
	return corefoundation.CGRect(rv)
}

// Returns the index of the character that lies beneath the specified point
// using the specified container’s coordinate system.
//
// point: The point to test.
//
// container: The text container within which the point is tested.
//
// partialFraction: A fraction of the distance from the insertion point, logically before the
// given character to the next one.
//
// # Return Value
//
// The index of the character falling under `point`.
//
// # Discussion
//
// Analogous to
// [GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph], but
// expressed in character index terms. The method returns the index of the
// character falling under `point`, expressed in coordinate system of
// `container`; if no character is under the point, the nearest character is
// returned, where nearest is defined according to the requirements of
// selection by mouse. However, this is not simply equivalent to taking the
// result of the corresponding glyph index method and converting it to a
// character index, because in some cases a single glyph represents more than
// one selectable character, for example an fi ligature glyph. In that case,
// there is an insertion point within the glyph, and this method returns one
// character or the other, depending on whether the specified point lies to
// the left or the right of that insertion point.
//
// In general, this method returns only character indexes for which there is
// an insertion point. The `partialFraction` is a fraction of the distance
// from the insertion point, logically before the given character to the next
// one, which may be either to the right or to the left depending on
// directionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterIndex(for:in:fractionOfDistanceBetweenInsertionPoints:)
func (l NSLayoutManager) CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point corefoundation.CGPoint, container INSTextContainer, partialFraction unsafe.Pointer) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:"), point, container, partialFraction)
	return rv
}

// Returns the range of characters that correspond to the glyphs in the
// specified glyph range.
//
// glyphRange: The glyph range for which to return the character range.
//
// actualGlyphRange: If not [NULL], on output, points to the full range of glyphs generated by
// the character range returned. This range may be identical or slightly
// larger than the requested glyph range. For example, if the text storage
// contains the character “`Ö`” and the glyph cache contains the two
// atomic glyphs “[O]” and “`¨`”, and if `glyphRange` encloses only
// the first or second glyph, then `actualGlyphRange` is set to enclose both
// glyphs.
//
// # Return Value
//
// The range of characters that generated the glyphs in `glyphRange`.
//
// # Discussion
//
// If the length of `glyphRange` is `0`, the resulting character range is a
// zero-length range just after the character(s) corresponding to the
// preceding glyph, and `actualGlyphRange` is also zero-length. If
// `glyphRange` extends beyond the text length, the method truncates the
// result to the number of characters in the text.
//
// If noncontiguous layout is not enabled, this method forces the generation
// of glyphs for all characters up to and including the end of the returned
// range.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterRange(forGlyphRange:actualGlyphRange:)
func (l NSLayoutManager) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.NSRange, actualGlyphRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return foundation.NSRange(rv)
}

// Enumerates enclosing rectangles for the specified glyph range in a text
// container.
//
// glyphRange: The glyph range for which to return enclosing rectangles.
//
// selectedRange: Selected glyphs within `glyphRange`, which can affect the size of the
// rectangles. If not interested in selection rectangles, pass `{NSNotFound,
// 0}` as the selected range.
//
// textContainer: The text container in which the glyphs are laid out.
//
// block: The block to apply to the glyph range. The block has two arguments:
//
// rect: The current enclosing rectangle. stop: A reference to a Boolean
// value. The block can set the value to true to stop further processing of
// the array. The stop argument is an out-only argument. You should only set
// this Boolean to true within the block.
//
// # Discussion
//
// These rectangles are always in container coordinates. They can be used to
// draw the text background or highlight for the given range of characters.
// The rectangles don’t necessarily enclose glyphs that draw outside their
// line fragment rectangles; use [BoundingRectForGlyphRangeInTextContainer] to
// determine the area that contains all drawing performed for a range of
// glyphs.
//
// If a selected range is given in the second argument, the rectangles
// returned are correct for drawing the selection. Selection rectangles are
// generally more complicated than enclosing rectangles, and supplying a
// selected range determines whether the method does this extra work. This
// method does the minimum amount of work required to answer the question.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/enumerateEnclosingRects(forGlyphRange:withinSelectedGlyphRange:in:using:)
func (l NSLayoutManager) EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.NSRange, selectedRange foundation.NSRange, textContainer INSTextContainer, block func(corefoundation.CGRect, *bool)) {
	_block3 := objc.NewBlock(func(_ objc.Block, arg0 corefoundation.CGRect, arg1 *bool) { block(arg0, arg1) })
	defer _block3.Release()
	objc.Send[objc.ID](l.ID, objc.Sel("enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:"), glyphRange, selectedRange, textContainer, objc.ID(_block3))
}

// Enumerates line fragments intersecting with the specified glyph range.
//
// glyphRange: The glyph range for which to return line fragment rectangles.
//
// block: The block to apply to the glyph range. The block has five arguments:
//
// rect: The current line fragment rectangle. usedRect: The portion of the
// line fragment rectangle that actually contains glyphs or other marks that
// are drawn (including the text container’s line fragment padding).
// textContainer: The text container in which the glyphs are laid out.
// glyphRange: The range of glyphs laid out in the current line fragment.
// stop: A reference to a Boolean value. The block can set the value to true
// to stop further processing of the array. The stop argument is an out-only
// argument. You should only set this Boolean to true within the block.
//
// # Discussion
//
// This method causes glyph generation and layout for the line fragment
// containing the glyphs in the specified range, or if noncontiguous layout is
// not enabled, for all of the text up to and including that line fragment.
//
// Line fragment rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/enumerateLineFragments(forGlyphRange:using:)
func (l NSLayoutManager) EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.NSRange, block TextContainerHandler) {
	_block1, _ := NewTextContainerBlock(block)
	objc.Send[objc.ID](l.ID, objc.Sel("enumerateLineFragmentsForGlyphRange:usingBlock:"), glyphRange, _block1)
}

// Returns the fraction of the distance between the glyph at the specified
// point and the next glyph.
//
// # Discussion
//
// This method is a primitive for
// [GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph]. You
// should always call the main method, not the primitives.
//
// Overriding should be done for the primitive methods. Existing subclasses
// that do not do this overriding will not have their implementations
// available to Java developers.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/fractionOfDistanceThroughGlyph(for:in:)
func (l NSLayoutManager) FractionOfDistanceThroughGlyphForPointInTextContainer(point corefoundation.CGPoint, container INSTextContainer) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:inTextContainer:"), point, container)
	return rv
}

// Returns insertion points in bulk for a specified line fragment.
//
// charIndex: The character index of one character within the line fragment.
//
// aFlag: If true, returns alternate, rather than primary, insertion points.
//
// dFlag: If true, returns insertion points in display, rather than logical, order.
//
// positions: On output, the positions of the insertion points, in the order specified.
//
// charIndexes: On output, the indexes of the characters corresponding to the returned
// insertion points.
//
// # Return Value
//
// The number of insertion points returned.
//
// # Discussion
//
// The method allows clients to obtain all insertion points for a line
// fragment in one call. Each pointer passed in should either be [NULL] or
// else point to sufficient memory to hold as many elements as there are
// insertion points in the line fragment (which cannot be more than the number
// of characters + 1). The returned positions indicate a transverse offset
// relative to the line fragment rectangle’s origin. Internal caching is
// used to ensure that repeated calls to this method for the same line
// fragment (possibly with differing values for other arguments) are not
// significantly more expensive than a single call.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getLineFragmentInsertionPoints(forCharacterAt:alternatePositions:inDisplayOrder:positions:characterIndexes:)
func (l NSLayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions unsafe.Pointer, charIndexes unsafe.Pointer) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:"), charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}

// Returns the index of the glyph at the specified location in a text
// container.
//
// # Discussion
//
// This method is a primitive for
// [GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph]. You
// should always call the main method, not the primitives.
//
// Overriding should be done for the primitive methods. Existing subclasses
// that do not do this overriding will not have their implementations
// available to Java developers.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndex(for:in:)
func (l NSLayoutManager) GlyphIndexForPointInTextContainer(point corefoundation.CGPoint, container INSTextContainer) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("glyphIndexForPoint:inTextContainer:"), point, container)
	return rv
}

// Returns the index of the glyph at the specified point using the
// container’s coordinate system.
//
// point: The point for which to return the glyph, in coordinates of `container`.
//
// container: The container in which the returned glyph is laid out.
//
// partialFraction: If not [NULL], on output, the fraction of the distance between the location
// of the glyph returned and the location of the next glyph.
//
// # Return Value
//
// The index of the glyph falling under the given point, expressed in the
// given container’s coordinate system.
//
// # Discussion
//
// If no glyph is under `point`, the nearest glyph is returned, where nearest
// is defined according to the requirements of selection by mouse. Clients who
// wish to determine whether the the point actually lies within the bounds of
// the glyph returned should follow this with a call to
// [BoundingRectForGlyphRangeInTextContainer] and test whether the point falls
// in the rectangle returned by that method. If `partialFraction` is non-NULL,
// it returns by reference the fraction of the distance between the location
// of the glyph returned and the location of the next glyph.
//
// For purposes such as dragging out a selection or placing the insertion
// point, a partial percentage less than or equal to 0.5 indicates that
// `point` should be considered as falling before the glyph index returned; a
// partial percentage greater than 0.5 indicates that it should be considered
// as falling after the glyph index returned. If the nearest glyph doesn’t
// lie under `point` at all (for example, if `point` is beyond the beginning
// or end of a line), this ratio is 0 or 1.
//
// If the glyph stream contains the glyphs “A” and “b”, with the width
// of “A” being 13 points, and the user clicks at a location 8 points into
// “A”, `partialFraction` is 8/13, or 0.615. In this case, the point given
// should be considered as falling between “A” and “b” for purposes
// such as dragging out a selection or placing the insertion point.
//
// Performs glyph generation and layout if needed.
//
// As part of its implementation, this method calls
// [FractionOfDistanceThroughGlyphForPointInTextContainer] and
// [GlyphIndexForPointInTextContainer]. To change this method’s behavior,
// override those two methods instead of this one.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndex(for:in:fractionOfDistanceThroughGlyph:)
func (l NSLayoutManager) GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point corefoundation.CGPoint, container INSTextContainer, partialFraction unsafe.Pointer) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("glyphIndexForPoint:inTextContainer:fractionOfDistanceThroughGlyph:"), point, container, partialFraction)
	return rv
}

// Returns the smallest contiguous range for glyphs lying wholly or partially
// within the specified rectangle of the text container.
//
// bounds: The bounding rectangle for which to return glyphs.
//
// container: The text container in which the glyphs are laid out.
//
// # Return Value
//
// The range of glyphs that would need to be displayed in order to draw all
// glyphs that fall (even partially) within the given bounding rectangle. The
// range returned can include glyphs that don’t fall inside or intersect
// `bounds`, although the first and last glyphs in the range always do. At
// most this method returns the glyph range for the whole container.
//
// # Discussion
//
// This method is used to determine which glyphs need to be displayed within a
// given rectangle.
//
// Performs glyph generation and layout if needed. Bounding rectangles are
// always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forBoundingRect:in:)
func (l NSLayoutManager) GlyphRangeForBoundingRectInTextContainer(bounds corefoundation.CGRect, container INSTextContainer) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("glyphRangeForBoundingRect:inTextContainer:"), bounds, container)
	return foundation.NSRange(rv)
}

// Returns the smallest contiguous range for glyphs lying wholly or partially
// within the specified rectangle of the text container.
//
// bounds: The bounding rectangle for which to return glyphs.
//
// container: The text container in which the glyphs are laid out.
//
// # Return Value
//
// The range of glyphs that would need to be displayed in order to draw all
// glyphs that fall (even partially) within the given bounding rectangle. The
// range returned can include glyphs that don’t fall inside or intersect
// `bounds`, although the first and last glyphs in the range always do. At
// most this method returns the glyph range for the whole container.
//
// # Discussion
//
// Unlike [GlyphRangeForBoundingRectInTextContainer], this variant of the
// method doesn’t perform glyph generation or layout. Its results, though
// faster, can be incorrect. This method is primarily for use by [NSTextView];
// you should rarely need to use it yourself.
//
// Bounding rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forBoundingRectWithoutAdditionalLayout:in:)
func (l NSLayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds corefoundation.CGRect, container INSTextContainer) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:"), bounds, container)
	return foundation.NSRange(rv)
}

// Returns the range of glyphs lying within the specified text container.
//
// # Discussion
//
// This is a less efficient method than the similar
// [TextContainerForGlyphAtIndexEffectiveRange].
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(for:)
func (l NSLayoutManager) GlyphRangeForTextContainer(container INSTextContainer) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("glyphRangeForTextContainer:"), container)
	return foundation.NSRange(rv)
}

// Returns the range of glyphs that the specified range of characters
// generates.
//
// charRange: The character range for which to return the generated glyph range.
//
// actualCharRange: If not [NULL], on output, points to the actual range of characters that
// fully define the glyph range returned. This range may be identical to or
// slightly larger than the requested character range. For example, if the
// text storage contains the characters “[O]” and `"¨"`, and the glyph
// store contains the single precomposed glyph “¨`Ö`”, and if
// `charRange` encloses only the first or second character, then
// `actualCharRange` is set to enclose both characters.
//
// # Return Value
//
// The range of glyphs generated by `charRange`.
//
// # Discussion
//
// If the length of `charRange` is `0`, the resulting glyph range is a
// zero-length range just after the glyph(s) corresponding to the preceding
// character, and `actualCharRange` will also be zero-length. If `charRange`
// extends beyond the text length, the method truncates the result to the
// number of glyphs in the text.
//
// If noncontiguous layout is not enabled, this method forces the generation
// of glyphs for all characters up to and including the end of the specified
// range.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forCharacterRange:actualCharacterRange:)
func (l NSLayoutManager) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.NSRange, actualCharRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return foundation.NSRange(rv)
}

// Returns the range of displayable glyphs that surround the glyph at the
// specified index.
//
// glyphIndex: Index of the glyph to test.
//
// # Return Value
//
// The range of nominally spaced glyphs.
//
// # Discussion
//
// This method returns the range for the glyphs around the given glyph that
// can be displayed using only their advancements from the font, without
// pairwise kerning or other adjustments to spacing. The range returned begins
// with the first glyph, counting back from `glyphIndex`, that has a location
// set, and it continues up to, but does not include, the next glyph that has
// a location set.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/range(ofNominallySpacedGlyphsContaining:)
func (l NSLayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](l.ID, objc.Sel("rangeOfNominallySpacedGlyphsContainingIndex:"), glyphIndex)
	return foundation.NSRange(rv)
}

// Draws background marks for the specified glyphs, which must lie completely
// within a single text container.
//
// glyphsToShow: The range of glyphs for which the background is drawn.
//
// origin: The position of the text container in the coordinate system of the
// currently focused view.
//
// # Discussion
//
// This method is called by [NSTextView] for drawing. You can override it to
// perform additional drawing, or to replace text drawing entirely, but not to
// change layout. You can call this method directly, but focus must already be
// locked on the destination view or image.
//
// Background marks are such things as selection highlighting, text background
// color, and any background for marked text, along with block decoration such
// as table backgrounds and borders.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawBackground(forGlyphRange:at:)
func (l NSLayoutManager) DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.NSRange, origin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("drawBackgroundForGlyphRange:atPoint:"), glyphsToShow, origin)
}

// Draws the specified glyphs, which must lie completely within a single text
// container.
//
// glyphsToShow: The range of glyphs that are drawn.
//
// origin: The position of the text container in the coordinate system of the
// currently focused view.
//
// # Discussion
//
// This method is called by [NSTextView] for drawing. You can override it to
// perform additional drawing, or to replace text drawing entirely, but not to
// change layout. You can call this method directly, but focus must already be
// locked on the destination view or image. This method expects the coordinate
// system of the view to be flipped.
//
// This method draws the actual glyphs, including attachments, as well as any
// underlines or strikethoughs.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawGlyphs(forGlyphRange:at:)
func (l NSLayoutManager) DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.NSRange, origin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("drawGlyphsForGlyphRange:atPoint:"), glyphsToShow, origin)
}

// Draws a strikethrough for the specified glyphs.
//
// glyphRange: The range of glyphs for which to draw a strikethrough. The range must
// belong to a single line fragment rectangle (as returned by
// [LineFragmentRectForGlyphAtIndexEffectiveRange]).
//
// strikethroughVal: The style of strikethrough to draw. This value is a mask derived from the
// value for [underlineStyle]—for example, `(NSUnderlinePatternDash |
// NSUnderlineStyleThick)`. Subclasses can define custom strikethrough styles.
//
// baselineOffset: Indicates how far above the text baseline the underline should be drawn.
//
// lineRect: The line fragment rectangle containing the glyphs to draw strikethrough
// for.
//
// lineGlyphRange: The range of all glyphs within `lineRect`.
//
// containerOrigin: The origin of the line fragment rectangle’s [NSTextContainer] in its
// [NSTextView].
//
// # Discussion
//
// This method is invoked automatically by
// [StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin];
// you should rarely need to invoke it directly. This method’s
// `strikethroughVal` parameter does not take account of any setting
// for[NSUnderlineByWordMask] because that’s taken care of by
// [UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawStrikethrough(forGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
//
// [underlineStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/underlineStyle
// [NSUnderlineByWordMask]: https://developer.apple.com/documentation/AppKit/NSUnderlineByWordMask
func (l NSLayoutManager) DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, strikethroughVal NSUnderlineStyle, baselineOffset float64, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

// Draws underlining for the glyphs in a specified range.
//
// glyphRange: A range of glyphs, which must belong to a single line fragment rectangle
// (as returned by [LineFragmentRectForGlyphAtIndexEffectiveRange]).
//
// underlineVal: The style of underlining to draw. This value is a mask derived from the
// value for [underlineStyle]—for example, `(NSUnderlinePatternDash |
// NSUnderlineStyleThick)`. Subclasses can define custom underlining styles.
//
// baselineOffset: Specifies the distance from the bottom of the bounding box of the specified
// glyphs in the specified range to their baseline.
//
// lineRect: The line fragment rectangle containing the glyphs to draw underlining for.
//
// lineGlyphRange: The range of all glyphs within `lineRect`.
//
// containerOrigin: The origin of the `lineRectNSTextContainer` in its [NSTextView].
//
// # Discussion
//
// This method is invoked automatically by
// [UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin];
// you should rarely need to invoke it directly. This method’s
// `underlineVal` parameter does not take account of any setting
// for[NSUnderlineByWordMask] because that’s taken care of by
// [UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawUnderline(forGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
//
// [underlineStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/underlineStyle
// [NSUnderlineByWordMask]: https://developer.apple.com/documentation/AppKit/NSUnderlineByWordMask
func (l NSLayoutManager) DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, underlineVal NSUnderlineStyle, baselineOffset float64, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

// Fills background rectangles with a color.
//
// rectArray: The array of rectangles to fill.
//
// rectCount: The number of rectangles in `rectArray`.
//
// charRange: The range of characters whose background rectangles are filled.
//
// color: The fill color.
//
// # Discussion
//
// This is the primitive method used by [DrawBackgroundForGlyphRangeAtPoint],
// providing a finer-grained override point for actually filling rectangles
// with a particular background color for a background color attribute, a
// selected or marked range highlight, a block decoration, or any other
// rectangle fill needed by that method. As with
// [showPackedGlyphs:length:glyphRange:atPoint:font:color:printingAdjustment:],
// the `charRange` and `color` parameters are passed in merely for
// informational purposes; the color is already set in the graphics state. If
// for any reason you modify it, you must restore it before returning from
// this method.
//
// This method operates in terms of character ranges, because the relevant
// attributes are expressed on characters, and they don’t always lie on
// glyph boundaries—for example, when one character of an “fi” ligature
// is highlighted.
//
// You should never call this method, but you might override it. The default
// implementation simply fills the rectangles in the specified array. The
// graphics operation used for this fill is controlled by a link check; for
// compatibility reasons, it is [NSCompositeCopy] for applications linked
// prior to OS X v10.6 and [NSCompositeSourceOver] for applications linked on
// macOS 10.6 or later. Applications can control the operation used, or modify
// the drawing, by overriding this method in an [NSLayoutManager] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/fillBackgroundRectArray(_:count:forCharacterRange:color:)
//
// [NSCompositeCopy]: https://developer.apple.com/documentation/AppKit/NSCompositeCopy
// [NSCompositeSourceOver]: https://developer.apple.com/documentation/AppKit/NSCompositeSourceOver
// [showPackedGlyphs:length:glyphRange:atPoint:font:color:printingAdjustment:]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showPackedGlyphs:length:glyphRange:atPoint:font:color:printingAdjustment:
func (l NSLayoutManager) FillBackgroundRectArrayCountForCharacterRangeColor(rectArray []corefoundation.CGRect, rectCount uint, charRange foundation.NSRange, color INSColor) {
	objc.Send[objc.ID](l.ID, objc.Sel("fillBackgroundRectArray:count:forCharacterRange:color:"), objc.CArray(rectArray), rectCount, charRange, color)
}

// Renders the glyphs at the specified positions, using the specified
// attributes.
//
// glyphs: The glyphs to draw, which may include embedded [NULL] bytes.
//
// positions: The positions at which to draw the glyphs in the user space coordinate
// system.
//
// glyphCount: The number of glyphs to draw.
//
// font: The font to apply to the graphics state. This value can be different from
// the [font] value in the `attributes` argument because of various font
// substitutions that the system automatically executes.
//
// textMatrix: The affine transform mapping the text space coordinate system to the user
// space coordinate system. The `tx` and `ty` components of `textMatrix` are
// ignored because Quartz overrides them with the glyph positions.
//
// attributes: A dictionary of glyph attributes. For a list of possible keys and values,
// see [Glyph Attributes].
//
// CGContext: A graphics context object already configured with the information in the
// `font`, `textMatrix`, and `attributes` parameters
//
// # Discussion
//
// The layout manager calls this primitive method when it is time to lay out a
// set of glyphs in the specified graphics context.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showCGGlyphs(_:positions:count:font:textMatrix:attributes:in:)
//
// [font]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/font
// [Glyph Attributes]: https://developer.apple.com/documentation/AppKit/glyph-attributes
func (l NSLayoutManager) ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.CGFontIndex, positions []corefoundation.CGPoint, glyphCount int, font NSFont, textMatrix corefoundation.CGAffineTransform, attributes foundation.INSDictionary, CGContext coregraphics.CGContextRef) {
	objc.Send[objc.ID](l.ID, objc.Sel("showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:"), objc.CArray(glyphs), objc.CArray(positions), glyphCount, font, textMatrix, attributes, CGContext)
}

// Calculates and draws strikethrough for the specified glyphs.
//
// glyphRange: The range of glyphs for which to draw a strikethrough. The range must
// belong to a single line fragment rectangle (as returned by
// [LineFragmentRectForGlyphAtIndexEffectiveRange]).
//
// strikethroughVal: The style of underlining to draw. This value is a mask derived from the
// value for [underlineStyle]—for example, `(NSUnderlinePatternDash |
// NSUnderlineStyleThick | NSUnderlineByWordMask)`. Subclasses can define
// custom underlining styles.
//
// lineRect: The line fragment rectangle containing the glyphs to draw strikethrough
// for.
//
// lineGlyphRange: The range of all glyphs within `lineRect`.
//
// containerOrigin: The origin of the line fragment rectangle’s [NSTextContainer] in its
// [NSTextView].
//
// # Discussion
//
// This method determines which glyphs actually need to have a strikethrough
// drawn based on `strikethroughVal`. After determining which glyphs to draw
// strikethrough on, this method invokes
// [DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]
// for each contiguous range of glyphs that requires it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/strikethroughGlyphRange(_:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
//
// [underlineStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/underlineStyle
func (l NSLayoutManager) StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, strikethroughVal NSUnderlineStyle, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}

// Calculates subranges to underline for the specified glyphs and draws the
// underlining as appropriate.
//
// glyphRange: A range of glyphs, which must belong to a single line fragment rectangle
// (as returned by [LineFragmentRectForGlyphAtIndexEffectiveRange]).
//
// underlineVal: The style of underlining to draw. This value is a mask derived from the
// value for [underlineStyle]—for example, `(NSUnderlinePatternDash |
// NSUnderlineStyleThick | NSUnderlineByWordMask)`. Subclasses can define
// custom underlining styles.
//
// lineRect: The line fragment rectangle containing the glyphs to draw underlining for.
//
// lineGlyphRange: The range of all glyphs within that line fragment rectangle.
//
// containerOrigin: The origin of the line fragment rectangle’s [NSTextContainer] in its
// [NSTextView].
//
// # Discussion
//
// This method determines which glyphs actually need to be underlined based on
// `underlineVal`. With [NSUnderlineStyleSingle], for example, leading and
// trailing whitespace isn’t underlined, but whitespace between visible
// glyphs is. A potential word-underline style would omit underlining on any
// whitespace. After determining which glyphs to draw underlining on, this
// method invokes
// [DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin]
// for each contiguous range of glyphs that requires it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/underlineGlyphRange(_:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
//
// [underlineStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/underlineStyle
func (l NSLayoutManager) UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.NSRange, underlineVal NSUnderlineStyle, lineRect corefoundation.CGRect, lineGlyphRange foundation.NSRange, containerOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}

// Sets the layout rectangle that encloses the specified text block and glyph
// range.
//
// rect: The layout rectangle to set.
//
// block: The text block whose layout rectangle is set.
//
// glyphRange: The range of glyphs in the text block.
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLayoutRect(_:for:glyphRange:)
func (l NSLayoutManager) SetLayoutRectForTextBlockGlyphRange(rect corefoundation.CGRect, block INSTextBlock, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setLayoutRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}

// Returns the rectangle for the layout of the specified text block and glyph
// range.
//
// # Return Value
//
// The layout rectangle, or [NSZeroRect] if no rectangle has been set for the
// specified block since the last invalidation.
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutRect(for:glyphRange:)
func (l NSLayoutManager) LayoutRectForTextBlockGlyphRange(block INSTextBlock, glyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("layoutRectForTextBlock:glyphRange:"), block, glyphRange)
	return corefoundation.CGRect(rv)
}

// Sets the bounding rectangle that encloses the specified text block and
// glyph range.
//
// rect: The bounding rectangle to set.
//
// block: The text block whose bounding rectangle is set.
//
// glyphRange: The range of glyphs in the text block.
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setBoundsRect(_:for:glyphRange:)
func (l NSLayoutManager) SetBoundsRectForTextBlockGlyphRange(rect corefoundation.CGRect, block INSTextBlock, glyphRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setBoundsRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}

// Returns the bounding rectangle that encloses the specified text block and
// glyph range.
//
// block: The text block whose bounds rectangle is returned.
//
// glyphRange: The range of glyphs in the text block.
//
// # Return Value
//
// The bounding rectangle, or [NSZeroRect] if no rectangle has been set for
// the specified block since the last invalidation
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundsRect(for:glyphRange:)
func (l NSLayoutManager) BoundsRectForTextBlockGlyphRange(block INSTextBlock, glyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("boundsRectForTextBlock:glyphRange:"), block, glyphRange)
	return corefoundation.CGRect(rv)
}

// Returns the rectangle for the layout of the specified text block and glyph.
//
// block: The text block whose layout rectangle is returned.
//
// glyphIndex: Index of the glyph.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the text block.
//
// # Return Value
//
// The layout rectangle of the text block, or [NSZeroRect] if no rectangle has
// been set for the specified block since the last invalidation.
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutRect(for:at:effectiveRange:)
func (l NSLayoutManager) LayoutRectForTextBlockAtIndexEffectiveRange(block INSTextBlock, glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("layoutRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return corefoundation.CGRect(rv)
}

// Returns the bounding rectangle for the specified text block and glyph.
//
// block: The text block whose bounding rectangle is returned.
//
// glyphIndex: Index of the glyph.
//
// effectiveGlyphRange: If not [NULL], on output, the range for all glyphs in the text block.
//
// # Return Value
//
// The bounding rectangle of the text block, or [NSZeroRect] if no rectangle
// has been set for the specified block since the last invalidation.
//
// # Discussion
//
// This method causes glyph generation but not layout. Block layout rectangles
// and bounds rectangles are always in container coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundsRect(for:at:effectiveRange:)
func (l NSLayoutManager) BoundsRectForTextBlockAtIndexEffectiveRange(block INSTextBlock, glyphIndex uint, effectiveGlyphRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("boundsRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return corefoundation.CGRect(rv)
}

// Draws an attachment cell.
//
// cell: The attachment cell to draw.
//
// rect: The rectangle within which to draw `cell`.
//
// attachmentIndex: The location of the attachment cell.
//
// # Discussion
//
// The `attachmentIndex` parameter is provided for cells that alter their
// appearance based on their location.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showAttachmentCell(_:in:characterIndex:)
func (l NSLayoutManager) ShowAttachmentCellInRectCharacterIndex(cell INSCell, rect corefoundation.CGRect, attachmentIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("showAttachmentCell:inRect:characterIndex:"), cell, rect, attachmentIndex)
}

// Returns the accessory view that the text system uses for its ruler.
//
// view: The text view using the layout manager.
//
// style: Sets the state of the controls in the accessory view; must not be `nil`.
//
// ruler: The ruler view whose accessory view is returned.
//
// isEnabled: If true, the accessory view is enabled and accepts mouse and keyboard
// events; if false it’s disabled.
//
// # Return Value
//
// The accessory view containing tab wells, text alignment buttons, and so on.
//
// # Discussion
//
// If you have turned off automatic ruler updating through the use of
// [UsesRuler] so that you can do more complex things, but you still want to
// display the appropriate accessory view, you can use this method.
//
// This method is invoked automatically by the [NSTextView] object using the
// layout manager. You should rarely need to invoke it, but you can override
// it to customize ruler support. If you do use this method directly, note
// that it neither installs the ruler accessory view nor sets the markers for
// the [NSRulerView] object. You must install the accessory view into the
// ruler using the [NSRulerView] method [AccessoryView]. To set the markers,
// use [RulerMarkersForTextViewParagraphStyleRuler] to get the markers needed,
// and then send [Markers] to the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rulerAccessoryView(for:paragraphStyle:ruler:enabled:)
func (l NSLayoutManager) RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view INSTextView, style INSParagraphStyle, ruler INSRulerView, isEnabled bool) INSView {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:"), view, style, ruler, isEnabled)
	return NSViewFromID(rv)
}

// Returns an array of text ruler objects for the current selection.
//
// view: The text view using the layout manager.
//
// style: Sets the state of the controls in the accessory view; must not be `nil`.
//
// ruler: The ruler view whose ruler markers are returned.
//
// # Return Value
//
// An array of [NSRulerMarker] objects representing such things as left and
// right margins, first-line indent, and tab stops.
//
// # Discussion
//
// If you have turned off automatic ruler updating through the use of
// [UsesRuler] so that you can do more complex things, but you still want to
// display the appropriate accessory view, you can use this method.
//
// This method is invoked automatically by the [NSTextView] object using the
// layout manager. You should rarely need to invoke it, but you can override
// it to add new kinds of markers or otherwise customize ruler support.
//
// You can set the returned ruler markers with the [NSRulerView] method
// [Markers].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rulerMarkers(for:paragraphStyle:ruler:)
func (l NSLayoutManager) RulerMarkersForTextViewParagraphStyleRuler(view INSTextView, style INSParagraphStyle, ruler INSRulerView) []NSRulerMarker {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("rulerMarkersForTextView:paragraphStyle:ruler:"), view, style, ruler)
	return objc.ConvertSlice(rv, func(id objc.ID) NSRulerMarker {
		return NSRulerMarkerFromID(id)
	})
}

// Indicates whether the first responder in the specified window is a text
// view for the layout manager.
//
// window: The window whose first responder is tested.
//
// # Return Value
//
// true if the first responder in `window` is a text view associated with the
// receiver; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutManagerOwnsFirstResponder(in:)
func (l NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow(window INSWindow) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("layoutManagerOwnsFirstResponderInWindow:"), window)
	return rv
}

// Returns the default line height for a line of text that uses a specified
// font.
//
// theFont: The font for which to determine the default line height.
//
// # Return Value
//
// The default line height for a line of text drawn using `theFont`.
//
// # Discussion
//
// The value returned may vary according to the layout manager’s typesetter
// behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultLineHeight(for:)
func (l NSLayoutManager) DefaultLineHeightForFont(theFont NSFont) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("defaultLineHeightForFont:"), theFont)
	return rv
}

// Returns the default baseline offset that the layout manager’s typesetter
// uses for the specified font.
//
// theFont: The font for which to return the default baseline offset.
//
// # Return Value
//
// The default baseline offset for a line of text drawn using `theFont`.
//
// # Discussion
//
// The value returned may vary according to the layout manager’s typesetter
// behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultBaselineOffset(for:)
func (l NSLayoutManager) DefaultBaselineOffsetForFont(theFont NSFont) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("defaultBaselineOffsetForFont:"), theFont)
	return rv
}

// Appends one or more temporary attributes to the attributes dictionary of
// the specified character range.
//
// attrs: Attributes dictionary containing the temporary attributes to add.
//
// charRange: The range of characters to which the specified attributes apply.
//
// # Discussion
//
// Temporary attributes are used only for onscreen drawing and are not
// persistent in any way. [NSTextView] uses them to color misspelled words
// when continuous spell checking is enabled. Currently the only temporary
// attributes recognized are those that do not affect layout (colors,
// underlines, and so on).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttributes(_:forCharacterRange:)
func (l NSLayoutManager) AddTemporaryAttributesForCharacterRange(attrs foundation.INSDictionary, charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Adds a temporary attribute to the characters in the specified range.
//
// attrName: The name of a temporary attribute.
//
// value: The temporary attribute value associated with `attrName`.
//
// charRange: The range of characters to which the specified attribute-value pair
// applies.
//
// # Discussion
//
// Raises an [InvalidArgumentException] if `attrName` or `value` is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttribute(_:value:forCharacterRange:)
func (l NSLayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName foundation.NSString, value objectivec.IObject, charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}

// Sets one or more temporary attributes for the specified character range.
//
// attrs: Attributes dictionary containing the temporary attributes to set.
//
// charRange: The range of characters to which the specified attributes apply.
//
// # Discussion
//
// Temporary attributes are used only for onscreen drawing and are not
// persistent in any way. [NSTextView] uses them to color misspelled words
// when continuous spell checking is enabled. Currently the only temporary
// attributes recognized are those that do not affect layout (colors,
// underlines, and so on).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setTemporaryAttributes(_:forCharacterRange:)
func (l NSLayoutManager) SetTemporaryAttributesForCharacterRange(attrs foundation.INSDictionary, charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("setTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Removes a temporary attribute from the list of attributes for the specified
// character range.
//
// attrName: The name of a temporary attribute.
//
// charRange: The range of characters from which to remove the specified temporary
// attribute.
//
// # Discussion
//
// Temporary attributes are used only for onscreen drawing and are not
// persistent in any way. [NSTextView] uses them to color misspelled words
// when continuous spell checking is enabled. Currently the only temporary
// attributes recognized are those that do not affect layout (colors,
// underlines, and so on).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTemporaryAttribute(_:forCharacterRange:)
func (l NSLayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName foundation.NSString, charRange foundation.NSRange) {
	objc.Send[objc.ID](l.ID, objc.Sel("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

// Returns the value for the temporary attribute of a character, and the range
// it applies to.
//
// attrName: The name of a temporary attribute.
//
// location: The index for which to return attributes. This value must not exceed the
// bounds of the receiver.
//
// range: If non-[NULL]:
//
// - If the named attribute exists at `location`, on output, contains the
// range over which the named attribute’s value applies. - If the named
// attribute does not exist at `location`, on output, contains the range over
// which the attribute does not exist.
//
// The range isn’t necessarily the maximum range covered by `attrName`, and
// its extent is implementation-dependent. If you need the maximum range, use
// [TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange]. If you
// don’t need this value, pass [NULL].
//
// # Return Value
//
// The value for the temporary attribute named `attrName` of the character at
// index `location`, or `nil` if there is no such attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttribute(_:atCharacterIndex:effectiveRange:)
func (l NSLayoutManager) TemporaryAttributeAtCharacterIndexEffectiveRange(attrName foundation.NSString, location uint, range_ foundation.NSRange) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("temporaryAttribute:atCharacterIndex:effectiveRange:"), attrName, location, range_)
	return objectivec.Object{ID: rv}
}

// Returns the value for the temporary attribute of a character, and the
// maximum range it applies to.
//
// attrName: The name of a temporary attribute.
//
// location: The index for which to return attributes. This value must not exceed the
// bounds of the receiver.
//
// range: If non-[NULL]:
//
// - If the named attribute exists at `location`, on output, contains the
// maximum range over which the named attribute’s value applies, clipped to
// `rangeLimit`. - If the named attribute does not exist at `location`, on
// output, contains the maximum range over which the attribute does not exist.
//
// If you don’t need this value, pass [NULL].
//
// rangeLimit: The range over which to search for continuous presence of `attrName`. This
// value must not exceed the bounds of the receiver.
//
// # Return Value
//
// The value for the attribute named `attrName` of the character at
// `location`, or `nil` if there is no such attribute.
//
// # Discussion
//
// If you don’t need the longest effective range, it’s far more efficient
// to use the [TemporaryAttributeAtCharacterIndexEffectiveRange] method to
// retrieve the attribute value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttribute(_:atCharacterIndex:longestEffectiveRange:in:)
func (l NSLayoutManager) TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.NSString, location uint, range_ foundation.NSRange, rangeLimit foundation.NSRange) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return objectivec.Object{ID: rv}
}

// Returns the dictionary of temporary attributes for the specified character
// range.
//
// # Return Value
//
// The dictionary of temporary attributes for the character range specified in
// `effectiveCharRange` at character index `charIndex`.
//
// # Discussion
//
// Temporary attributes are used only for onscreen drawing and are not
// persistent in any way. [NSTextView] uses them to color misspelled words
// when continuous spell checking is enabled. Currently the only temporary
// attributes recognized are those that do not affect layout (colors,
// underlines, and so on).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttributes(atCharacterIndex:effectiveRange:)
func (l NSLayoutManager) TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("temporaryAttributesAtCharacterIndex:effectiveRange:"), charIndex, effectiveCharRange)
	return foundation.NSDictionaryFromID(rv)
}

// Returns the temporary attributes for a character, and the maximum range
// they apply to.
//
// location: The index for which to return attributes. This value must not exceed the
// bounds of the receiver.
//
// range: If not [NULL], on output, contains the maximum range over which the
// attributes and values are the same as those at `location`, clipped to
// `rangeLimit`.
//
// rangeLimit: The range over which to search for continuous presence of the attributes at
// `location`. This value must not exceed the bounds of the receiver.
//
// # Return Value
//
// The attributes for the character at `location`.
//
// # Discussion
//
// If you don’t need the longest effective range, it’s far more efficient
// to use the [TemporaryAttributesAtCharacterIndexEffectiveRange] method to
// retrieve the attribute value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttributes(atCharacterIndex:longestEffectiveRange:in:)
func (l NSLayoutManager) TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ foundation.NSRange, rangeLimit foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("temporaryAttributesAtCharacterIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return foundation.NSDictionaryFromID(rv)
}

// Returns the text storage object from which the [NSGlyphGenerator] object
// procures characters for glyph generation.
//
// # Return Value
//
// The receiver’s text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/attributedString()
func (l NSLayoutManager) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(rv)
}

// Returns the glyph at the specified index.
//
// glyphIndex: The index of a glyph in the receiver. This value must not exceed the bounds
// of the receiver’s glyph array.
//
// # Return Value
//
// The glyph at `glyphIndex`.
//
// # Discussion
//
// Raises an [NSRangeException] if `glyphIndex` is out of bounds.
//
// Performs glyph generation if needed. To avoid an exception with
// [GlyphAtIndex] you must first check the glyph index against the number of
// glyphs, which requires generating all glyphs. Another method,
// [GlyphAtIndexIsValidIndex], generates glyphs only up to the one requested,
// so using it can be more efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyph(at:)
func (l NSLayoutManager) GlyphAtIndex(glyphIndex uint) NSGlyph {
	rv := objc.Send[NSGlyph](l.ID, objc.Sel("glyphAtIndex:"), glyphIndex)
	return NSGlyph(rv)
}

// Returns the glyph at a specified index, and optionally returns a flag
// indicating whether the requested index is valid.
//
// glyphIndex: The index of the glyph to be returned.
//
// isValidIndex: If not [NULL], on output, true if the requested index is in range;
// otherwise false.
//
// # Return Value
//
// The glyph at the requested index, or [NSNullGlyph] if the requested index
// is out of the range `{0,` “NSLayoutManager/numberOfGlyphs```}`.
//
// # Discussion
//
// If noncontiguous layout is not enabled, this method causes generation of
// all glyphs up to and including `glyphIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyph(at:isValidIndex:)
func (l NSLayoutManager) GlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) NSGlyph {
	rv := objc.Send[NSGlyph](l.ID, objc.Sel("glyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return NSGlyph(rv)
}

// Returns the current layout options.
//
// # Return Value
//
// The layout options as a bit mask, as defined in Constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyphStorage/layoutOptions()
func (l NSLayoutManager) LayoutOptions() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("layoutOptions"))
	return rv
}

// Returns an array of rectangles and, by reference, the number of such
// rectangles, that define the region in the given container enclosing the
// given character range.
//
// charRange: The character range for which to return rectangles.
//
// selCharRange: Selected characters within `charRange`, which can affect the size of the
// rectangles; it must be equal to or contain `charRange`. If the caller is
// interested in this more from an enclosing point of view rather than a
// selection point of view, pass `{NSNotFound, 0}` as the selected range.
//
// container: The text container in which the text is laid out.
//
// rectCount: The number of rectangles returned.
//
// # Return Value
//
// The array of rectangles enclosing the given range.
//
// # Discussion
//
// These rectangles can be used to draw the text background or highlight for
// the given range of characters. If a selected range is given in
// `selCharRange`, the rectangles returned are correct for drawing the
// selection. Selection rectangles are generally more complicated than
// enclosing rectangles and supplying a selected range is the clue this method
// uses to determine whether to go to the trouble of doing this special work.
//
// This method will do the minimum amount of work required to answer the
// question. The resulting array is owned by the layout manager and will be
// reused when this method,
// [RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount],
// or [BoundingRectForGlyphRangeInTextContainer] is called. One of these
// methods may be called indirectly. If you aren’t going to use the
// rectangles right away, you should copy them to another location. These
// rectangles are always in container coordinates.
//
// The number of rectangles returned isn’t necessarily the number of lines
// enclosing the specified range. Contiguous lines can share an enclosing
// rectangle, and lines broken into several fragments have a separate
// enclosing rectangle for each fragment.
//
// These rectangles don’t necessarily enclose glyphs that draw outside their
// line fragment rectangles; use [BoundingRectForGlyphRangeInTextContainer] to
// determine the area that contains all drawing performed for a range of
// glyphs.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rectArray(forCharacterRange:withinSelectedCharacterRange:in:rectCount:)
func (l NSLayoutManager) RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount(charRange foundation.NSRange, selCharRange foundation.NSRange, container INSTextContainer, rectCount unsafe.Pointer) foundation.NSRect {
	rv := objc.Send[foundation.NSRect](l.ID, objc.Sel("rectArrayForCharacterRange:withinSelectedCharacterRange:inTextContainer:rectCount:"), charRange, selCharRange, container, rectCount)
	return foundation.NSRect(rv)
}

// Returns an array of rectangles and, by reference, the number of such
// rectangles, that define the region in the given container enclosing the
// given glyph range.
//
// glyphRange: The glyph range for which to return rectangles.
//
// selGlyphRange: Selected glyphs within `glyphRange`, which can affect the size of the
// rectangles; it must be equal to or contain `glyphRange`. If the caller is
// interested in this more from an enclosing point of view rather than a
// selection point of view, pass `{NSNotFound, 0}` as the selected range.
//
// container: The text container in which the text is laid out.
//
// rectCount: The number of rectangles returned.
//
// # Return Value
//
// The array of rectangles enclosing the given range.
//
// # Discussion
//
// These rectangles can be used to draw the text background or highlight for
// the given range of characters. If a selected range is given in
// `selGlyphRange`, the rectangles returned are correct for drawing the
// selection. Selection rectangles are generally more complicated than
// enclosing rectangles and supplying a selected range is the clue this method
// uses to determine whether to go to the trouble of doing this special work.
//
// The number of rectangles returned isn’t necessarily the number of lines
// enclosing the specified range. Contiguous lines can share an enclosing
// rectangle, and lines broken into several fragments have a separate
// enclosing rectangle for each fragment.
//
// This method will do the minimum amount of work required to answer the
// question. The resulting array is owned by the layout manager and will be
// reused when this method,
// [RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount],
// or [BoundingRectForGlyphRangeInTextContainer] is called. One of these
// methods may be called indirectly. If you aren’t going to use the
// rectangles right away, you should copy them to another location. These
// rectangles are always in container coordinates.
//
// The purpose of this method is to calculate line rectangles for drawing the
// text background and highlighting. These rectangles don’t necessarily
// enclose glyphs that draw outside their line fragment rectangles; use
// [BoundingRectForGlyphRangeInTextContainer] to determine the area that
// contains all drawing performed for a range of glyphs.
//
// Performs glyph generation and layout if needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rectArray(forGlyphRange:withinSelectedGlyphRange:in:rectCount:)
func (l NSLayoutManager) RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount(glyphRange foundation.NSRange, selGlyphRange foundation.NSRange, container INSTextContainer, rectCount unsafe.Pointer) foundation.NSRect {
	rv := objc.Send[foundation.NSRect](l.ID, objc.Sel("rectArrayForGlyphRange:withinSelectedGlyphRange:inTextContainer:rectCount:"), glyphRange, selGlyphRange, container, rectCount)
	return foundation.NSRect(rv)
}
func (l NSLayoutManager) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
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
func (l NSLayoutManager) InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs NSGlyph, length uint, glyphIndex uint, charIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
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
func (l NSLayoutManager) SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}

// The layout manager’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/delegate
func (l NSLayoutManager) Delegate() NSLayoutManagerDelegate {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("delegate"))
	return NSLayoutManagerDelegateObjectFromID(rv)
}
func (l NSLayoutManager) SetDelegate(value NSLayoutManagerDelegate) {
	objc.Send[struct{}](l.ID, objc.Sel("setDelegate:"), value)
}

// The text storage object that contains the content to lay out.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textStorage
func (l NSLayoutManager) TextStorage() NSTextStorage {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("textStorage"))
	return NSTextStorageFromID(objc.ID(rv))
}
func (l NSLayoutManager) SetTextStorage(value NSTextStorage) {
	objc.Send[struct{}](l.ID, objc.Sel("setTextStorage:"), value)
}

// A Boolean value that indicates whether the layout manager allows
// noncontiguous layout.
//
// # Discussion
//
// For more information about noncontiguous layout, see [NSLayoutManager].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/allowsNonContiguousLayout
func (l NSLayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("allowsNonContiguousLayout"))
	return rv
}
func (l NSLayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setAllowsNonContiguousLayout:"), value)
}

// A Boolean value that indicates whether the layout manager currently has any
// areas of noncontiguous layout.
//
// # Discussion
//
// There may be times at which there is no noncontiguous layout, such as when
// layout is complete; this method enables the layout manager to report that
// to clients.
//
// For more information about noncontiguous layout, see [NSLayoutManager].
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/hasNonContiguousLayout
func (l NSLayoutManager) HasNonContiguousLayout() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("hasNonContiguousLayout"))
	return rv
}

// A Boolean value that indicates whether to substitute visible glyphs for
// whitespace and other typically invisible characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsInvisibleCharacters
func (l NSLayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("showsInvisibleCharacters"))
	return rv
}
func (l NSLayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setShowsInvisibleCharacters:"), value)
}

// A Boolean value that indicates whether the layout manager substitutes
// visible glyphs for control characters in the layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsControlCharacters
func (l NSLayoutManager) ShowsControlCharacters() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("showsControlCharacters"))
	return rv
}
func (l NSLayoutManager) SetShowsControlCharacters(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setShowsControlCharacters:"), value)
}

// A Boolean value that indicates whether the layout manager uses the leading
// of the font.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesFontLeading
func (l NSLayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("usesFontLeading"))
	return rv
}
func (l NSLayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setUsesFontLeading:"), value)
}

// A Boolean value that indicates whether the layout manager generates glyphs
// and lays them out when the app’s run loop is idle.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/backgroundLayoutEnabled
func (l NSLayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("backgroundLayoutEnabled"))
	return rv
}
func (l NSLayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setBackgroundLayoutEnabled:"), value)
}

// A Boolean value that indicates whether the layout manager avoids laying out
// unusually long or suspicious input.
//
// # Discussion
//
// The default value of this property is false, which causes the layout
// manager to lay out whatever text you give it. Changing the value to true
// causes the layout manager to generate invalid layout information when it
// detects potentially suspicious content.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/limitsLayoutForSuspiciousContents
func (l NSLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}
func (l NSLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}

// A Boolean value that indicates whether the layout manager uses the default
// hyphenation rules to wrap lines.
//
// # Discussion
//
// When the value of this property is true, the layout manager makes a
// best-effort attempt to hyphenate text when wrapping lines. You may override
// this hyphenation behavior on a per-paragraph basis using the
// [hyphenationFactor] property of [NSParagraphStyle] The default value of
// this property is false, which prevents the layout manager from hyphenating
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesDefaultHyphenation
//
// [hyphenationFactor]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/hyphenationFactor
func (l NSLayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}
func (l NSLayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}

// The current text containers of the layout manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainers
func (l NSLayoutManager) TextContainers() []NSTextContainer {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("textContainers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextContainer {
		return NSTextContainerFromID(id)
	})
}

// The glyph generator that the layout manager uses.
//
// # Discussion
//
// Setting the glyph generator invalidates all glyphs and layout in the layout
// manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphGenerator
func (l NSLayoutManager) GlyphGenerator() INSGlyphGenerator {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("glyphGenerator"))
	return NSGlyphGeneratorFromID(objc.ID(rv))
}
func (l NSLayoutManager) SetGlyphGenerator(value INSGlyphGenerator) {
	objc.Send[struct{}](l.ID, objc.Sel("setGlyphGenerator:"), value)
}

// The number of glyphs in the layout manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/numberOfGlyphs
func (l NSLayoutManager) NumberOfGlyphs() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("numberOfGlyphs"))
	return rv
}

// The rectangle for the extra line fragment at the end of a document.
//
// # Discussion
//
// The layout manager uses the extra line fragment when the last character in
// a document causes a line or paragraph break. This extra line fragment has
// no corresponding glyph.
//
// The rectangle is defined in the coordinate system of its [NSTextContainer].
// [NSZeroRect] if there is no such rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentRect
func (l NSLayoutManager) ExtraLineFragmentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("extraLineFragmentRect"))
	return corefoundation.CGRect(rv)
}

// The text container for the extra line fragment rectangle.
//
// # Discussion
//
// This rectangle is used to display the insertion point at the end of a text
// (either in an empty text or after a final paragraph separator).
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentTextContainer
func (l NSLayoutManager) ExtraLineFragmentTextContainer() INSTextContainer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("extraLineFragmentTextContainer"))
	return NSTextContainerFromID(objc.ID(rv))
}

// The rectangle that encloses the insertion point in the extra line fragment
// rectangle.
//
// # Discussion
//
// The rectangle is defined in the coordinate system of its [NSTextContainer].
// [NSZeroRect] if there is no extra line fragment rectangle.
//
// The extra line fragment used rectangle is twice as wide (or tall) as the
// text container’s line fragment padding, with the insertion point itself
// in the middle.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentUsedRect
func (l NSLayoutManager) ExtraLineFragmentUsedRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("extraLineFragmentUsedRect"))
	return corefoundation.CGRect(rv)
}

// The default amount of scaling to apply when an attachment image is too
// large to fit in a text container.
//
// # Discussion
//
// Attachment cells control their own size and drawing, so this setting is
// only advisory to them, but Application Kit–supplied attachment cells
// respect it.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultAttachmentScaling
func (l NSLayoutManager) DefaultAttachmentScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](l.ID, objc.Sel("defaultAttachmentScaling"))
	return NSImageScaling(rv)
}
func (l NSLayoutManager) SetDefaultAttachmentScaling(value NSImageScaling) {
	objc.Send[struct{}](l.ID, objc.Sel("setDefaultAttachmentScaling:"), value)
}

// The first text view in the layout manager’s series of text views.
//
// # Discussion
//
// This [NSTextView] object is the recipient of various [NSText] and
// [NSTextView] notifications.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstTextView
func (l NSLayoutManager) FirstTextView() INSTextView {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("firstTextView"))
	return NSTextViewFromID(objc.ID(rv))
}

// The text view that contains the first glyph in the selection.
//
// # Discussion
//
// This property does not cause layout if the beginning of the selected range
// is not yet laid out.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textViewForBeginningOfSelection
func (l NSLayoutManager) TextViewForBeginningOfSelection() INSTextView {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("textViewForBeginningOfSelection"))
	return NSTextViewFromID(objc.ID(rv))
}

// The current typesetter.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetter
func (l NSLayoutManager) Typesetter() INSTypesetter {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("typesetter"))
	return NSTypesetterFromID(objc.ID(rv))
}
func (l NSLayoutManager) SetTypesetter(value INSTypesetter) {
	objc.Send[struct{}](l.ID, objc.Sel("setTypesetter:"), value)
}

// The default typesetter behavior.
//
// # Discussion
//
// The typesetter behavior affects glyph spacing and line height.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l NSLayoutManager) TypesetterBehavior() NSTypesetterBehavior {
	rv := objc.Send[NSTypesetterBehavior](l.ID, objc.Sel("typesetterBehavior"))
	return NSTypesetterBehavior(rv)
}
func (l NSLayoutManager) SetTypesetterBehavior(value NSTypesetterBehavior) {
	objc.Send[struct{}](l.ID, objc.Sel("setTypesetterBehavior:"), value)
}
