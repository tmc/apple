// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods that delegates of layout manager objects implement.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate
type NSLayoutManagerDelegate interface {
	objectivec.IObject
}

// NSLayoutManagerDelegateObject wraps an existing Objective-C object that conforms to the NSLayoutManagerDelegate protocol.
type NSLayoutManagerDelegateObject struct {
	objectivec.Object
}

func (o NSLayoutManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSLayoutManagerDelegateObjectFromID constructs a [NSLayoutManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSLayoutManagerDelegateObjectFromID(id objc.ID) NSLayoutManagerDelegateObject {
	return NSLayoutManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate when the specified layout manager invalidates layout
// information (not glyph information).
//
// sender: The layout manager that invalidated layout.
//
// # Discussion
//
// This method is invoked only when layout was complete and then became
// invalidated for some reason. Delegates can use this information to show an
// indicator of background layout or to enable a button that forces immediate
// layout of text.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManagerDidInvalidateLayout(_:)
func (o NSLayoutManagerDelegateObject) LayoutManagerDidInvalidateLayout(sender INSLayoutManager) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutManagerDidInvalidateLayout:"), sender)
}

// Enables customization of the initial glyph generation process.
//
// layoutManager: The layout manager doing the layout.
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
// # Return Value
//
// The actual glyph range stored in this method. By returning `0`, it can
// indicate for the layout manager to do the default processing.
//
// # Discussion
//
// This message is sent whenever the layout manager is about to store the
// initial glyph information via
// [SetGlyphsPropertiesCharacterIndexesFontForGlyphRange]. To customize the
// initial glyph generation process, this method can invoke
// [SetGlyphsPropertiesCharacterIndexesFontForGlyphRange] with modified glyph
// information.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:)
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager INSLayoutManager, glyphs *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexes unsafe.Pointer, aFont NSFont, glyphRange foundation.NSRange) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"), layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
	return rv
}

// Returns the control character action for the control character at the
// specified character index.
//
// layoutManager: The layout manager doing the layout.
//
// action: The proposed control character action for the character at the given index.
// Possible values are enumerated by [NSLayoutManager.ControlCharacterAction].
//
// charIndex: The index of the control character for which the action is proposed.
//
// # Return Value
//
// The control character action for the control character at the given index.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldUse:forControlCharacterAt:)
//
// [NSLayoutManager.ControlCharacterAction]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager INSLayoutManager, action NSControlCharacterAction, charIndex uint) NSControlCharacterAction {
	rv := objc.Send[NSControlCharacterAction](o.ID, objc.Sel("layoutManager:shouldUseAction:forControlCharacterAtIndex:"), layoutManager, action, charIndex)
	return rv
}

// Informs the delegate when the layout manager finishes laying out text in
// the specified text container.
//
// layoutManager: The layout manager doing the layout.
//
// textContainer: The text container in which layout is complete. If `nil`, if there aren’t
// enough containers to hold all the text; the delegate can use this
// information as a cue to add another text container.
//
// layoutFinishedFlag: If true, `aLayoutManager` is finished laying out its text—this also means
// that `aTextContainer` is the final text container used by the layout
// manager. Delegates can use this information to show an indicator or
// background or to enable or disable a button that forces immediate layout of
// text.
//
// # Discussion
//
// This message is sent whenever a text container has been filled. This method
// can be useful for paginating.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:didCompleteLayoutFor:atEnd:)
func (o NSLayoutManagerDelegateObject) LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager INSLayoutManager, textContainer INSTextContainer, layoutFinishedFlag bool) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutManager:didCompleteLayoutForTextContainer:atEnd:"), layoutManager, textContainer, layoutFinishedFlag)
}

// Informs the delegate when the layout manager invalidates layout due to a
// change in the geometry of the specified text container.
//
// layoutManager: The layout manager invalidating layout.
//
// textContainer: The text container that changed geometry.
//
// oldSize: The size of the text container before it changed geometry.
//
// # Discussion
//
// The delegate can react to the geometry change and perform adjustments such
// as recreating an exclusion path.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:textContainer:didChangeGeometryFrom:)
func (o NSLayoutManagerDelegateObject) LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager INSLayoutManager, textContainer INSTextContainer, oldSize corefoundation.CGSize) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutManager:textContainer:didChangeGeometryFromSize:"), layoutManager, textContainer, oldSize)
}

// Asks the delegate whether to break the line at the specified character.
//
// layoutManager: The layout manager doing the layout.
//
// charIndex: Index of the character delimiting the hyphenation point search.
//
// # Return Value
//
// true if the current hyphenation point is acceptable; false if the layout
// manager should find the next hyphenation opportunity before `charIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldBreakLineByHyphenatingBeforeCharacterAt:)
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager INSLayoutManager, charIndex uint) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), layoutManager, charIndex)
	return rv
}

// Asks the delegate whether to break the line at the specified word.
//
// layoutManager: The layout manager doing the layout.
//
// charIndex: Index of the character delimiting the break point search.
//
// # Return Value
//
// true if the current line break point is acceptable; false if the layout
// manager should find the next break point opportunity before `charIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldBreakLineByWordBeforeCharacterAt:)
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager INSLayoutManager, charIndex uint) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"), layoutManager, charIndex)
	return rv
}

// Returns the amount of space to add to the end of a line.
//
// layoutManager: The layout manager doing the layout.
//
// glyphIndex: The index of the glyph at the end of the line.
//
// rect: The proposed line fragment rectangle for the current line.
//
// # Return Value
//
// The line spacing after the current line.
//
// # Discussion
//
// This message is sent while each line is laid out to enable the layout
// manager delegate to customize the shape of line.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:lineSpacingAfterGlyphAt:withProposedLineFragmentRect:)
func (o NSLayoutManagerDelegateObject) LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager INSLayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
	return rv
}

// Returns the amount of space to add at the end of a paragraph.
//
// layoutManager: The layout manager doing the layout.
//
// glyphIndex: The index of the glyph at the end of the line.
//
// rect: The proposed line fragment rectangle for the current line.
//
// # Return Value
//
// The paragraph spacing after the current line.
//
// # Discussion
//
// This message is sent while each line is laid out to enable the layout
// manager delegate to customize the shape of line.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:paragraphSpacingAfterGlyphAt:withProposedLineFragmentRect:)
func (o NSLayoutManagerDelegateObject) LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager INSLayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
	return rv
}

// Returns the amount of space to add at the beginning of a paragraph.
//
// layoutManager: The layout manager doing the layout.
//
// glyphIndex: The index of the glyph at the beginning of the line.
//
// rect: The proposed line fragment rectangle for the current line.
//
// # Return Value
//
// The paragraph spacing before the current line.
//
// # Discussion
//
// This message is sent while each line is laid out to enable the layout
// manager delegate to customize the shape of line.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:paragraphSpacingBeforeGlyphAt:withProposedLineFragmentRect:)
func (o NSLayoutManagerDelegateObject) LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager INSLayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
	return rv
}

// Returns the bounding rectangle for the specified control glyph with the
// specified parameters.
//
// layoutManager: The layout manager doing the layout.
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
// The bounding rectangle for the specified control glyph with the specified
// parameters.
//
// # Discussion
//
// Sent for resolving the glyph metrics for
// [NSControlCharacterWhitespaceAction] control character.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:boundingBoxForControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:)
//
// [NSControlCharacterWhitespaceAction]: https://developer.apple.com/documentation/UIKit/NSControlCharacterWhitespaceAction
func (o NSLayoutManagerDelegateObject) LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager INSLayoutManager, glyphIndex uint, textContainer INSTextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}

// Customizes the line fragment geometry before committing it to the layout
// cache.
//
// layoutManager: The layout manager doing the work.
//
// lineFragmentRect: The proposed rectangle that contains the glyphs. You may modify this
// rectangle as needed.
//
// lineFragmentUsedRect: The portion of `lineFragmentRect` that actually contains glyphs or other
// rendered marks, including the text container’s line fragment padding.
// This rectangle must be equal to `lineFragmentRect` or wholly contained by
// it. You may modify this rectangle as needed.
//
// baselineOffset: The vertical distance (in pixels) from the line fragment origin to the
// baseline on which the glyphs align.
//
// textContainer: The text container for the line fragments.
//
// glyphRange: The range of glyphs being laid out.
//
// # Return Value
//
// true if you modified the layout information and want your modifications to
// be used or false if the original layout information should be used.
//
// # Discussion
//
// Use this method to modify the line fragment rectangles associated with the
// text container. It is your responsibility to ensure that the modified
// rectangles remain valid and still lie within the text container.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldSetLineFragmentRect:lineFragmentUsedRect:baselineOffset:in:forGlyphRange:)
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(layoutManager INSLayoutManager, lineFragmentRect *corefoundation.CGRect, lineFragmentUsedRect *corefoundation.CGRect, baselineOffset unsafe.Pointer, textContainer INSTextContainer, glyphRange foundation.NSRange) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldSetLineFragmentRect:lineFragmentUsedRect:baselineOffset:inTextContainer:forGlyphRange:"), layoutManager, lineFragmentRect, lineFragmentUsedRect, baselineOffset, textContainer, glyphRange)
	return rv
}

// Asks the delegate whether to use temporary attributes when drawing the
// text.
//
// layoutManager: The layout manager sending the message.
//
// attrs: The temporary attributes currently in effect for the given character range.
//
// toScreen: true if the layout manager is drawing to the screen; otherwise, false.
//
// charIndex: Index of the first character in the range being drawn.
//
// effectiveCharRange: On input and output, the effective range to which the temporary attributes
// apply.
//
// # Return Value
//
// The temporary attributes for the layout manager to use, or `nil` if no
// temporary attributes are to be used.
//
// # Discussion
//
// The default behavior, if this method is not implemented, is to use
// temporary attributes only when drawing to the screen, so an implementation
// to match that behavior would return `attrs` if `toScreen` is true and `nil`
// otherwise, without changing `effectiveCharRange`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutManagerDelegate/layoutManager(_:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:)
func (o NSLayoutManagerDelegateObject) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager INSLayoutManager, attrs foundation.INSDictionary, toScreen bool, charIndex uint, effectiveCharRange foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
	return foundation.NSDictionaryFromID(rv)
}

// NSLayoutManagerDelegateConfig holds optional typed callbacks for [NSLayoutManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nslayoutmanagerdelegate
type NSLayoutManagerDelegateConfig struct {

	// Invalidating glyphs and layout
	// LayoutManagerDidInvalidateLayout — Informs the delegate when the specified layout manager invalidates layout information (not glyph information).
	LayoutManagerDidInvalidateLayout func(sender NSLayoutManager)
	// LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange — Enables customization of the initial glyph generation process.
	LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange func(layoutManager NSLayoutManager, glyphs *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexes *uint, aFont NSFont, glyphRange foundation.NSRange) uint

	// Managing temporary attribute support
	// LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange — Asks the delegate whether to use temporary attributes when drawing the text.
	LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange func(layoutManager NSLayoutManager, attrs foundation.INSDictionary, toScreen bool, charIndex uint, effectiveCharRange foundation.NSRange) foundation.INSDictionary

	// Other Methods
	// LayoutManagerShouldUseActionForControlCharacterAtIndex — Returns the control character action for the control character at the specified character index.
	LayoutManagerShouldUseActionForControlCharacterAtIndex func(layoutManager NSLayoutManager, action NSControlCharacterAction, charIndex uint) NSControlCharacterAction
	// LayoutManagerDidCompleteLayoutForTextContainerAtEnd — Informs the delegate when the layout manager finishes laying out text in the specified text container.
	LayoutManagerDidCompleteLayoutForTextContainerAtEnd func(layoutManager NSLayoutManager, textContainer NSTextContainer, layoutFinishedFlag bool)
	// LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex — Asks the delegate whether to break the line at the specified character.
	LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex func(layoutManager NSLayoutManager, charIndex uint) bool
	// LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex — Asks the delegate whether to break the line at the specified word.
	LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex func(layoutManager NSLayoutManager, charIndex uint) bool
	// LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange — Customizes the line fragment geometry before committing it to the layout cache.
	LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange func(layoutManager NSLayoutManager, lineFragmentRect *corefoundation.CGRect, lineFragmentUsedRect *corefoundation.CGRect, baselineOffset *float64, textContainer NSTextContainer, glyphRange foundation.NSRange) bool
}

// NewNSLayoutManagerDelegate creates an Objective-C object implementing the [NSLayoutManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSLayoutManagerDelegateObject] satisfies the [NSLayoutManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nslayoutmanagerdelegate
func NewNSLayoutManagerDelegate(config NSLayoutManagerDelegateConfig) NSLayoutManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSLayoutManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.LayoutManagerDidInvalidateLayout != nil {
		fn := config.LayoutManagerDidInvalidateLayout
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManagerDidInvalidateLayout:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSLayoutManagerFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange != nil {
		fn := config.LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, glyphs *coregraphics.CGFontIndex, props NSGlyphProperty, charIndexes *uint, aFontID objc.ID, glyphRange foundation.NSRange) uint {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				aFont := NSFontFromID(aFontID)
				return fn(layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
			},
		})
	}

	if config.LayoutManagerShouldUseActionForControlCharacterAtIndex != nil {
		fn := config.LayoutManagerShouldUseActionForControlCharacterAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldUseAction:forControlCharacterAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, action NSControlCharacterAction, charIndex uint) NSControlCharacterAction {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				return fn(layoutManager, action, charIndex)
			},
		})
	}

	if config.LayoutManagerDidCompleteLayoutForTextContainerAtEnd != nil {
		fn := config.LayoutManagerDidCompleteLayoutForTextContainerAtEnd
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:didCompleteLayoutForTextContainer:atEnd:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, textContainerID objc.ID, layoutFinishedFlag bool) {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				textContainer := NSTextContainerFromID(textContainerID)
				fn(layoutManager, textContainer, layoutFinishedFlag)
			},
		})
	}

	if config.LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex != nil {
		fn := config.LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, charIndex uint) bool {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				return fn(layoutManager, charIndex)
			},
		})
	}

	if config.LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex != nil {
		fn := config.LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, charIndex uint) bool {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				return fn(layoutManager, charIndex)
			},
		})
	}

	if config.LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange != nil {
		fn := config.LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldSetLineFragmentRect:lineFragmentUsedRect:baselineOffset:inTextContainer:forGlyphRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, lineFragmentRect *corefoundation.CGRect, lineFragmentUsedRect *corefoundation.CGRect, baselineOffset *float64, textContainerID objc.ID, glyphRange foundation.NSRange) bool {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				textContainer := NSTextContainerFromID(textContainerID)
				return fn(layoutManager, lineFragmentRect, lineFragmentUsedRect, baselineOffset, textContainer, glyphRange)
			},
		})
	}

	if config.LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange != nil {
		fn := config.LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layoutManagerID objc.ID, attrsID objc.ID, toScreen bool, charIndex uint, effectiveCharRange foundation.NSRange) objc.ID {
				layoutManager := NSLayoutManagerFromID(layoutManagerID)
				attrs := foundation.NSDictionaryFromID(attrsID)
				return fn(layoutManager, attrs, toScreen, charIndex, effectiveCharRange).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSLayoutManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSLayoutManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSLayoutManagerDelegateObjectFromID(instance)
}
