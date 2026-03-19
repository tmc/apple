// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextLayoutManager] class.
var (
	_NSTextLayoutManagerClass     NSTextLayoutManagerClass
	_NSTextLayoutManagerClassOnce sync.Once
)

func getNSTextLayoutManagerClass() NSTextLayoutManagerClass {
	_NSTextLayoutManagerClassOnce.Do(func() {
		_NSTextLayoutManagerClass = NSTextLayoutManagerClass{class: objc.GetClass("NSTextLayoutManager")}
	})
	return _NSTextLayoutManagerClass
}

// GetNSTextLayoutManagerClass returns the class object for NSTextLayoutManager.
func GetNSTextLayoutManagerClass() NSTextLayoutManagerClass {
	return getNSTextLayoutManagerClass()
}

type NSTextLayoutManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextLayoutManagerClass) Alloc() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The primary class that you use to manage text layout and presentation for
// custom text displays.
//
// # Overview
// 
// [NSTextLayoutManager] is the centerpiece of the TextKit object network that
// maintains the layout geometry through an array of [NSTextContainer]
// objects. It lays out results using [NSTextLayoutFragment] and
// [NSTextElement] objects vended from a [NSTextContentManager] that
// participates in the content layout process.
//
// # Creating a layout manager
//
//   - [NSTextLayoutManager.InitWithCoder]: Creates a new text layout manager with the coder you provide.
//
// # Configuring global layout manager options
//
//   - [NSTextLayoutManager.LayoutQueue]: The queue that the framework dispatches layout operations on.
//   - [NSTextLayoutManager.SetLayoutQueue]
//   - [NSTextLayoutManager.RenderingAttributesValidator]: A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//   - [NSTextLayoutManager.SetRenderingAttributesValidator]
//   - [NSTextLayoutManager.UsesFontLeading]: A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//   - [NSTextLayoutManager.SetUsesFontLeading]
//   - [NSTextLayoutManager.UsesHyphenation]: A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//   - [NSTextLayoutManager.SetUsesHyphenation]
//   - [NSTextLayoutManager.LimitsLayoutForSuspiciousContents]: A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//   - [NSTextLayoutManager.SetLimitsLayoutForSuspiciousContents]
//
// # Managing the layout process
//
//   - [NSTextLayoutManager.Delegate]: The delegate for the text layout manager object.
//   - [NSTextLayoutManager.SetDelegate]
//
// # Accessing the text storage
//
//   - [NSTextLayoutManager.TextContentManager]: Returns the text content manager associated with this text layout manager.
//   - [NSTextLayoutManager.TextContainer]: The text container object that provides geometric information for the layout destination.
//   - [NSTextLayoutManager.SetTextContainer]
//   - [NSTextLayoutManager.TextSelectionNavigation]: Returns a text selection manager configured to have the text layout manager as its data source.
//   - [NSTextLayoutManager.SetTextSelectionNavigation]
//   - [NSTextLayoutManager.TextSelections]: An array of text selections associated by the text layout manager.
//   - [NSTextLayoutManager.SetTextSelections]
//   - [NSTextLayoutManager.UsageBoundsForTextContainer]: Returns the usage bounds for the text container.
//   - [NSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]: Enumerates text segments of a specific type and in the text range you provide.
//   - [NSTextLayoutManager.ReplaceTextContentManager]: Replaces the current text content manager with a new one you provide.
//   - [NSTextLayoutManager.ReplaceContentsInRangeWithAttributedString]: Replaces content at the location you specify with an attributed string you provide.
//   - [NSTextLayoutManager.ReplaceContentsInRangeWithTextElements]: Replaces content at the location you specify with the text elements string you provide.
//
// # Adjusting rendering
//
//   - [NSTextLayoutManager.AddRenderingAttributeValueForTextRange]: Sets the rendering attribute for the value and range you specify.
//   - [NSTextLayoutManager.EnumerateRenderingAttributesFromLocationReverseUsingBlock]: Enumerates the rendering attributes from a location you specify.
//   - [NSTextLayoutManager.RenderingAttributesForLinkAtLocation]: Returns a dictionary of rendering attributes for rendering a link.
//   - [NSTextLayoutManager.InvalidateRenderingAttributesForTextRange]: Invalidates the rendering attributes of the specified text range.
//   - [NSTextLayoutManager.RemoveRenderingAttributeForTextRange]: Removes the rendering attribute from the specified text range.
//   - [NSTextLayoutManager.SetRenderingAttributesForTextRange]: Sets the rendering attributes for the range you specify.
//
// # Causing layout generation
//
//   - [NSTextLayoutManager.TextViewportLayoutController]: The text viewport layout controller associated with the layout manager’s text container.
//   - [NSTextLayoutManager.InvalidateLayoutForRange]: Invalidates the layout information for specified text range.
//   - [NSTextLayoutManager.TextLayoutFragmentForLocation]: Returns the text layout fragment from the document at the specified location.
//   - [NSTextLayoutManager.TextLayoutFragmentForPosition]: Returns the text layout fragment at the position you specify in the text container.
//   - [NSTextLayoutManager.EnsureLayoutForBounds]: Performs the layout for filling the bounds you specify inside the last text container.
//   - [NSTextLayoutManager.EnsureLayoutForRange]: Performs the layout for specified text range.
//   - [NSTextLayoutManager.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock]: Enumerates the text layout fragments starting at the specified location.
//
// # Instance Properties
//
//   - [NSTextLayoutManager.ResolvesNaturalAlignmentWithBaseWritingDirection]: Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
//   - [NSTextLayoutManager.SetResolvesNaturalAlignmentWithBaseWritingDirection]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager
type NSTextLayoutManager struct {
	objectivec.Object
}

// NSTextLayoutManagerFromID constructs a [NSTextLayoutManager] from an objc.ID.
//
// The primary class that you use to manage text layout and presentation for
// custom text displays.
func NSTextLayoutManagerFromID(id objc.ID) NSTextLayoutManager {
	return NSTextLayoutManager{objectivec.Object{ID: id}}
}
// NOTE: NSTextLayoutManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextLayoutManager] class.
//
// # Creating a layout manager
//
//   - [INSTextLayoutManager.InitWithCoder]: Creates a new text layout manager with the coder you provide.
//
// # Configuring global layout manager options
//
//   - [INSTextLayoutManager.LayoutQueue]: The queue that the framework dispatches layout operations on.
//   - [INSTextLayoutManager.SetLayoutQueue]
//   - [INSTextLayoutManager.RenderingAttributesValidator]: A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//   - [INSTextLayoutManager.SetRenderingAttributesValidator]
//   - [INSTextLayoutManager.UsesFontLeading]: A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//   - [INSTextLayoutManager.SetUsesFontLeading]
//   - [INSTextLayoutManager.UsesHyphenation]: A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//   - [INSTextLayoutManager.SetUsesHyphenation]
//   - [INSTextLayoutManager.LimitsLayoutForSuspiciousContents]: A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//   - [INSTextLayoutManager.SetLimitsLayoutForSuspiciousContents]
//
// # Managing the layout process
//
//   - [INSTextLayoutManager.Delegate]: The delegate for the text layout manager object.
//   - [INSTextLayoutManager.SetDelegate]
//
// # Accessing the text storage
//
//   - [INSTextLayoutManager.TextContentManager]: Returns the text content manager associated with this text layout manager.
//   - [INSTextLayoutManager.TextContainer]: The text container object that provides geometric information for the layout destination.
//   - [INSTextLayoutManager.SetTextContainer]
//   - [INSTextLayoutManager.TextSelectionNavigation]: Returns a text selection manager configured to have the text layout manager as its data source.
//   - [INSTextLayoutManager.SetTextSelectionNavigation]
//   - [INSTextLayoutManager.TextSelections]: An array of text selections associated by the text layout manager.
//   - [INSTextLayoutManager.SetTextSelections]
//   - [INSTextLayoutManager.UsageBoundsForTextContainer]: Returns the usage bounds for the text container.
//   - [INSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]: Enumerates text segments of a specific type and in the text range you provide.
//   - [INSTextLayoutManager.ReplaceTextContentManager]: Replaces the current text content manager with a new one you provide.
//   - [INSTextLayoutManager.ReplaceContentsInRangeWithAttributedString]: Replaces content at the location you specify with an attributed string you provide.
//   - [INSTextLayoutManager.ReplaceContentsInRangeWithTextElements]: Replaces content at the location you specify with the text elements string you provide.
//
// # Adjusting rendering
//
//   - [INSTextLayoutManager.AddRenderingAttributeValueForTextRange]: Sets the rendering attribute for the value and range you specify.
//   - [INSTextLayoutManager.EnumerateRenderingAttributesFromLocationReverseUsingBlock]: Enumerates the rendering attributes from a location you specify.
//   - [INSTextLayoutManager.RenderingAttributesForLinkAtLocation]: Returns a dictionary of rendering attributes for rendering a link.
//   - [INSTextLayoutManager.InvalidateRenderingAttributesForTextRange]: Invalidates the rendering attributes of the specified text range.
//   - [INSTextLayoutManager.RemoveRenderingAttributeForTextRange]: Removes the rendering attribute from the specified text range.
//   - [INSTextLayoutManager.SetRenderingAttributesForTextRange]: Sets the rendering attributes for the range you specify.
//
// # Causing layout generation
//
//   - [INSTextLayoutManager.TextViewportLayoutController]: The text viewport layout controller associated with the layout manager’s text container.
//   - [INSTextLayoutManager.InvalidateLayoutForRange]: Invalidates the layout information for specified text range.
//   - [INSTextLayoutManager.TextLayoutFragmentForLocation]: Returns the text layout fragment from the document at the specified location.
//   - [INSTextLayoutManager.TextLayoutFragmentForPosition]: Returns the text layout fragment at the position you specify in the text container.
//   - [INSTextLayoutManager.EnsureLayoutForBounds]: Performs the layout for filling the bounds you specify inside the last text container.
//   - [INSTextLayoutManager.EnsureLayoutForRange]: Performs the layout for specified text range.
//   - [INSTextLayoutManager.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock]: Enumerates the text layout fragments starting at the specified location.
//
// # Instance Properties
//
//   - [INSTextLayoutManager.ResolvesNaturalAlignmentWithBaseWritingDirection]: Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
//   - [INSTextLayoutManager.SetResolvesNaturalAlignmentWithBaseWritingDirection]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager
type INSTextLayoutManager interface {
	objectivec.IObject
	NSTextSelectionDataSource

	// Topic: Creating a layout manager

	// Creates a new text layout manager with the coder you provide.
	InitWithCoder(coder foundation.INSCoder) NSTextLayoutManager

	// Topic: Configuring global layout manager options

	// The queue that the framework dispatches layout operations on.
	LayoutQueue() foundation.NSOperationQueue
	SetLayoutQueue(value foundation.NSOperationQueue)
	// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
	RenderingAttributesValidator() TextLayoutManagerTextLayoutFragmentHandler
	SetRenderingAttributesValidator(value TextLayoutManagerTextLayoutFragmentHandler)
	// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
	UsesHyphenation() bool
	SetUsesHyphenation(value bool)
	// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)

	// Topic: Managing the layout process

	// The delegate for the text layout manager object.
	Delegate() NSTextLayoutManagerDelegate
	SetDelegate(value NSTextLayoutManagerDelegate)

	// Topic: Accessing the text storage

	// Returns the text content manager associated with this text layout manager.
	TextContentManager() INSTextContentManager
	// The text container object that provides geometric information for the layout destination.
	TextContainer() INSTextContainer
	SetTextContainer(value INSTextContainer)
	// Returns a text selection manager configured to have the text layout manager as its data source.
	TextSelectionNavigation() INSTextSelectionNavigation
	SetTextSelectionNavigation(value INSTextSelectionNavigation)
	// An array of text selections associated by the text layout manager.
	TextSelections() []NSTextSelection
	SetTextSelections(value []NSTextSelection)
	// Returns the usage bounds for the text container.
	UsageBoundsForTextContainer() corefoundation.CGRect
	// Enumerates text segments of a specific type and in the text range you provide.
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange INSTextRange, type_ NSTextLayoutManagerSegmentType, options NSTextLayoutManagerSegmentOptions, block CGRectTextContainerHandler)
	// Replaces the current text content manager with a new one you provide.
	ReplaceTextContentManager(textContentManager INSTextContentManager)
	// Replaces content at the location you specify with an attributed string you provide.
	ReplaceContentsInRangeWithAttributedString(range_ INSTextRange, attributedString foundation.NSAttributedString)
	// Replaces content at the location you specify with the text elements string you provide.
	ReplaceContentsInRangeWithTextElements(range_ INSTextRange, textElements []NSTextElement)

	// Topic: Adjusting rendering

	// Sets the rendering attribute for the value and range you specify.
	AddRenderingAttributeValueForTextRange(renderingAttribute foundation.NSString, value objectivec.IObject, textRange INSTextRange)
	// Enumerates the rendering attributes from a location you specify.
	EnumerateRenderingAttributesFromLocationReverseUsingBlock(location NSTextLocation, reverse bool, block VoidHandler)
	// Returns a dictionary of rendering attributes for rendering a link.
	RenderingAttributesForLinkAtLocation(link objectivec.IObject, location NSTextLocation) foundation.INSDictionary
	// Invalidates the rendering attributes of the specified text range.
	InvalidateRenderingAttributesForTextRange(textRange INSTextRange)
	// Removes the rendering attribute from the specified text range.
	RemoveRenderingAttributeForTextRange(renderingAttribute foundation.NSString, textRange INSTextRange)
	// Sets the rendering attributes for the range you specify.
	SetRenderingAttributesForTextRange(renderingAttributes foundation.INSDictionary, textRange INSTextRange)

	// Topic: Causing layout generation

	// The text viewport layout controller associated with the layout manager’s text container.
	TextViewportLayoutController() INSTextViewportLayoutController
	// Invalidates the layout information for specified text range.
	InvalidateLayoutForRange(range_ INSTextRange)
	// Returns the text layout fragment from the document at the specified location.
	TextLayoutFragmentForLocation(location NSTextLocation) INSTextLayoutFragment
	// Returns the text layout fragment at the position you specify in the text container.
	TextLayoutFragmentForPosition(position corefoundation.CGPoint) INSTextLayoutFragment
	// Performs the layout for filling the bounds you specify inside the last text container.
	EnsureLayoutForBounds(bounds corefoundation.CGRect)
	// Performs the layout for specified text range.
	EnsureLayoutForRange(range_ INSTextRange)
	// Enumerates the text layout fragments starting at the specified location.
	EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location NSTextLocation, options NSTextLayoutFragmentEnumerationOptions, block TextLayoutFragmentHandler) NSTextLocation

	// Topic: Instance Properties

	// Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextLayoutManager) Init() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextLayoutManager) Autorelease() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextLayoutManager creates a new NSTextLayoutManager instance.
func NewNSTextLayoutManager() NSTextLayoutManager {
	class := getNSTextLayoutManagerClass()
	rv := objc.Send[NSTextLayoutManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new text layout manager with the coder you provide.
//
// coder: A coder that implements [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/init(coder:)
func NewTextLayoutManagerWithCoder(coder foundation.INSCoder) NSTextLayoutManager {
	instance := getNSTextLayoutManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextLayoutManagerFromID(rv)
}

// Creates a new text layout manager with the coder you provide.
//
// coder: A coder that implements [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/init(coder:)
func (t NSTextLayoutManager) InitWithCoder(coder foundation.INSCoder) NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Enumerates text segments of a specific type and in the text range you
// provide.
//
// textRange: The range as an [NSTextRange].
//
// type: One of the available [NSTextLayoutManager.SegmentType] values.
// //
// [NSTextLayoutManager.SegmentType]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType
//
// options: One or more of the [NSTextLayoutManager.SegmentOptions] options.
// //
// [NSTextLayoutManager.SegmentOptions]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions
//
// block: A closure you provide to determine if the enumeration finishes early.
//
// # Discussion
// 
// A text segment is a logically and visually contiguous portion of the text
// content inside a line fragment that you specify with a single text range.
// The framework enumerates the segments visually from left to right.
// Returning `false` breaks out of the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateTextSegments(in:type:options:using:)
func (t NSTextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange INSTextRange, type_ NSTextLayoutManagerSegmentType, options NSTextLayoutManagerSegmentOptions, block CGRectTextContainerHandler) {
_block3, _cleanup3 := NewCGRectTextContainerBlock(block)
	defer _cleanup3()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateTextSegmentsInRange:type:options:usingBlock:"), textRange, type_, options, _block3)
}
// Replaces the current text content manager with a new one you provide.
//
// textContentManager: The new text context manager.
//
// # Discussion
// 
// Use this method to replace the current [NSTextContentManager] with a new
// one, leaving all related objects intact. This method makes sure the
// [NSTextLayoutManager] doesn’t get deallocated while migrating to the new
// manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replace(_:)
func (t NSTextLayoutManager) ReplaceTextContentManager(textContentManager INSTextContentManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceTextContentManager:"), textContentManager)
}
// Replaces content at the location you specify with an attributed string you
// provide.
//
// range: The range of the content to replace.
//
// attributedString: An attribued string to replace the content at `range`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-2elb
func (t NSTextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ INSTextRange, attributedString foundation.NSAttributedString) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceContentsInRange:withAttributedString:"), range_, attributedString)
}
// Replaces content at the location you specify with the text elements string
// you provide.
//
// range: The range of the content to replace.
//
// textElements: An array of text elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-80j0b
func (t NSTextLayoutManager) ReplaceContentsInRangeWithTextElements(range_ INSTextRange, textElements []NSTextElement) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, objectivec.IObjectSliceToNSArray(textElements))
}
// Sets the rendering attribute for the value and range you specify.
//
// renderingAttribute: The [NSAttributedString.Key] that represents the attribute.
// //
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
//
// value: The value for the attribute.
//
// textRange: The range over which to apply the attribute.
//
// # Discussion
// 
// Passing `nil` overrides the specified attribute by removing it from the
// final attributes the framework passes to the layout and rendering engine.
// This is a convenience method for [SetRenderingAttributesForTextRange].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/addRenderingAttribute(_:value:for:)
func (t NSTextLayoutManager) AddRenderingAttributeValueForTextRange(renderingAttribute foundation.NSString, value objectivec.IObject, textRange INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("addRenderingAttribute:value:forTextRange:"), renderingAttribute, value, textRange)
}
// Enumerates the rendering attributes from a location you specify.
//
// location: The location at which to start the enumeration.
//
// reverse: Whether to start the enumeration from the end of the range.
//
// block: A closure you provide to determine if the enumeration finishes early.
//
// # Discussion
// 
// This method only enumerates ranges with text that specify rendering
// attributes. Returning `false` from `block` breaks out of the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateRenderingAttributes(from:reverse:using:)
func (t NSTextLayoutManager) EnumerateRenderingAttributesFromLocationReverseUsingBlock(location NSTextLocation, reverse bool, block VoidHandler) {
_block2, _cleanup2 := NewVoidBlock(block)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateRenderingAttributesFromLocation:reverse:usingBlock:"), location, reverse, _block2)
}
// Returns a dictionary of rendering attributes for rendering a link.
//
// link: The link.
//
// location: The location of the link in the text.
//
// # Return Value
// 
// A dictionary of rendering attributes.
//
// # Discussion
// 
// As with other rendering attributes, specifying [NSNull] removes the
// attribute from the final attributes the framework uses for rendering. It
// has priority over the general rendering attributes.
//
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/renderingAttributes(forLink:at:)
func (t NSTextLayoutManager) RenderingAttributesForLinkAtLocation(link objectivec.IObject, location NSTextLocation) foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("renderingAttributesForLink:atLocation:"), link, location)
	return foundation.NSDictionaryFromID(rv)
}
// Invalidates the rendering attributes of the specified text range.
//
// textRange: The range of the text to invalidate.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/invalidateRenderingAttributes(for:)
func (t NSTextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateRenderingAttributesForTextRange:"), textRange)
}
// Removes the rendering attribute from the specified text range.
//
// renderingAttribute: The [NSAttributedString.Key] attribute to remove
// //
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
//
// textRange: The range over which to remove the rendering attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/removeRenderingAttribute(_:for:)
func (t NSTextLayoutManager) RemoveRenderingAttributeForTextRange(renderingAttribute foundation.NSString, textRange INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeRenderingAttribute:forTextRange:"), renderingAttribute, textRange)
}
// Sets the rendering attributes for the range you specify.
//
// renderingAttributes: A dictionary of rendering attributes.
//
// textRange: The text range over which to apply `renderingAttributes`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/setRenderingAttributes(_:for:)
func (t NSTextLayoutManager) SetRenderingAttributesForTextRange(renderingAttributes foundation.INSDictionary, textRange INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setRenderingAttributes:forTextRange:"), renderingAttributes, textRange)
}
// Invalidates the layout information for specified text range.
//
// range: The range of the layout to invalidate.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/invalidateLayout(for:)
func (t NSTextLayoutManager) InvalidateLayoutForRange(range_ INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateLayoutForRange:"), range_)
}
// Returns the text layout fragment from the document at the specified
// location.
//
// location: The starting location.
//
// # Return Value
// 
// An [NSTextLayoutFragment].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textLayoutFragment(for:)-68dez
func (t NSTextLayoutManager) TextLayoutFragmentForLocation(location NSTextLocation) INSTextLayoutFragment {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutFragmentForLocation:"), location)
	return NSTextLayoutFragmentFromID(rv)
}
// Returns the text layout fragment at the position you specify in the text
// container.
//
// position: A [CGPoint] that describes the position in the coordinate system for the
// text container.
// //
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// # Return Value
// 
// An [NSTextLayoutFragment].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textLayoutFragment(for:)-4dhrx
func (t NSTextLayoutManager) TextLayoutFragmentForPosition(position corefoundation.CGPoint) INSTextLayoutFragment {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutFragmentForPosition:"), position)
	return NSTextLayoutFragmentFromID(rv)
}
// Performs the layout for filling the bounds you specify inside the last text
// container.
//
// bounds: A [CGRect] that describes the layout bounds.
// //
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/ensureLayout(for:)-6ptsj
func (t NSTextLayoutManager) EnsureLayoutForBounds(bounds corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("ensureLayoutForBounds:"), bounds)
}
// Performs the layout for specified text range.
//
// range: The [NSTextRange] that describes the content to lay out.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/ensureLayout(for:)-3duae
func (t NSTextLayoutManager) EnsureLayoutForRange(range_ INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("ensureLayoutForRange:"), range_)
}
// Enumerates the text layout fragments starting at the specified location.
//
// location: The location where youstart the enumeration.
//
// options: One or more of the available [NSTextLayoutFragment.EnumerationOptions].
// //
// [NSTextLayoutFragment.EnumerationOptions]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions
//
// block: A closure you provide that determines if the enumeration finishes early.
//
// # Return Value
// 
// An [NSTextLocation], or `nil`.
//
// # Discussion
// 
// If `textLocation` is `nil`, the method starts at
// `self.TextContentManagerXCUIElementTypeDocumentRangeXCUIElementTypeLocation()`.The
// method uses `self.DocumentRangeXCUIElementTypeEndLocation()` for reverse
// enumeration. When enumerating backward, it starts with the fragment
// preceding the one containing `textLocation`. If the method enumerates at
// least one fragment, it returns the edge of the enumerated range.
// 
// The enumerated range might not match the range of the last element
// returned; it enumerates the elements in the sequence, but it can skip a
// range. For example, it can limit the maximum number of text elements the
// method enumerates for a single invocation or hide some elements from the
// layout.
// 
// Returning `false` from `block` breaks out of the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateTextLayoutFragments(from:options:using:)
func (t NSTextLayoutManager) EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location NSTextLocation, options NSTextLayoutFragmentEnumerationOptions, block TextLayoutFragmentHandler) NSTextLocation {
_block2, _cleanup2 := NewTextLayoutFragmentBlock(block)
	defer _cleanup2()
	rv := objc.Send[objc.ID](t.ID, objc.Sel("enumerateTextLayoutFragmentsFromLocation:options:usingBlock:"), location, options, _block2)
	return NSTextLocationObjectFromID(rv)
}
// Returns the base writing direction at the location you specify.
//
// location: The location where you want to examine the text’s writing direction.
//
// # Return Value
// 
// The [NSWritingDirection].
//
// [NSWritingDirection]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/baseWritingDirection(at:)
func (t NSTextLayoutManager) BaseWritingDirectionAtLocation(location NSTextLocation) NSTextSelectionNavigationWritingDirection {
	rv := objc.Send[NSTextSelectionNavigationWritingDirection](t.ID, objc.Sel("baseWritingDirectionAtLocation:"), location)
	return NSTextSelectionNavigationWritingDirection(rv)
}
// Enumerates all the insertion point caret offsets from left to right in
// visual order.
//
// location: The [NSTextLocation] to start from.
//
// block: The closure to invoke once for each logical caret edge in the line
// fragment, in left-to-right visual order. End the enumeration early by
// returning `false`.
//
// # Discussion
// 
// The `caretOffset` is in the coordinate system for the text container. When
// `leadingEdge` is `true`, it indicates that `caretOffset` is at the logical
// edge preceding the character. For left-to-right characters, the caret is on
// the left, and on the right for right-to-left characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/enumerateCaretOffsetsInLineFragment(at:using:)
func (t NSTextLayoutManager) EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location NSTextLocation, block NSTextLocationHandler) {
_block1, _cleanup1 := NewNSTextLocationBlock(block)
	defer _cleanup1()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateCaretOffsetsInLineFragmentAtLocation:usingBlock:"), location, _block1)
}
// Enumerates all the container boundaries starting from the location you
// specify.
//
// location: The location where the enumeration starts.
//
// reverse: A Boolean value that indicates the enumeration starts at the end of the
// container.
//
// block: A closure to invoke to evaluate the container boundaries; end the
// enumeration early by returning `false`.
//
// # Discussion
// 
// This is an optional method you implement to enumerate the text up to the
// container or page boundary when the text selection data provider supports
// this layout functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/enumerateContainerBoundaries(from:reverse:using:)
func (t NSTextLayoutManager) EnumerateContainerBoundariesFromLocationReverseUsingBlock(location NSTextLocation, reverse bool, block NSTextLocationHandler) {
_block2, _cleanup2 := NewNSTextLocationBlock(block)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateContainerBoundariesFromLocation:reverse:usingBlock:"), location, reverse, _block2)
}
// Returns the range of the line fragment that contains the point you specify.
//
// point: The starting point that contains the line fragment, in the coordinate
// system of `location`.
//
// location: The location of the line fragment.
//
// # Return Value
// 
// An [NSTextRange] that describes the location of the line fragment, or nil
// if the range isn’t found.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/lineFragmentRange(for:inContainerAt:)
func (t NSTextLayoutManager) LineFragmentRangeForPointInContainerAtLocation(point corefoundation.CGPoint, location NSTextLocation) INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("lineFragmentRangeForPoint:inContainerAtLocation:"), point, location)
	return NSTextRangeFromID(rv)
}
// Returns a new location using the location and offset you specify.
//
// location: The starting location in the selection.
//
// offset: An offset that describes the extent of the new location.
//
// # Return Value
// 
// A new `NSTextLocation, or nil` when the inputs don’t produce any legal
// location, such as when the input is an out of bounds index.
//
// # Discussion
// 
// The offset value can be positive or negative indicating the logical
// direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/location(_:offsetBy:)
func (t NSTextLayoutManager) LocationFromLocationWithOffset(location NSTextLocation, offset int) NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("locationFromLocation:withOffset:"), location, offset)
	return NSTextLocationObjectFromID(rv)
}
// Returns the offset between the two locations you specify.
//
// from: The starting location.
//
// to: The ending location.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/offset(from:to:)
func (t NSTextLayoutManager) OffsetFromLocationToLocation(from NSTextLocation, to NSTextLocation) int {
	rv := objc.Send[int](t.ID, objc.Sel("offsetFromLocation:toLocation:"), from, to)
	return rv
}
// Returns the layout orientation at the location you specify.
//
// location: The location where you want to examine the text’s layout orientation.
//
// # Return Value
// 
// Returns an [NSTextSelectionNavigation.LayoutOrientation] that describes the
// orientation of the layout.
//
// [NSTextSelectionNavigation.LayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/textLayoutOrientation(at:)
func (t NSTextLayoutManager) TextLayoutOrientationAtLocation(location NSTextLocation) NSTextSelectionNavigationLayoutOrientation {
	rv := objc.Send[NSTextSelectionNavigationLayoutOrientation](t.ID, objc.Sel("textLayoutOrientationAtLocation:"), location)
	return NSTextSelectionNavigationLayoutOrientation(rv)
}
// Returns a text range that corresponds to selection granularity of the
// enclosing location.
//
// selectionGranularity: One of the possible [NSTextSelection.Granularity] options.
// //
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
//
// location: A location that encloses the text range of interest.
//
// # Return Value
// 
// Returns the text range of the section, or `nil` when
// `documentRange.IsEmpty()` `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/textRange(for:enclosing:)
func (t NSTextLayoutManager) TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity NSTextSelectionGranularity, location NSTextLocation) INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textRangeForSelectionGranularity:enclosingLocation:"), selectionGranularity, location)
	return NSTextRangeFromID(rv)
}
func (t NSTextLayoutManager) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The queue that the framework dispatches layout operations on.
//
// # Discussion
// 
// If non-`nil`, it performs layout in the specified queue until
// `estimatedUsageBounds` becomes `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/layoutQueue
func (t NSTextLayoutManager) LayoutQueue() foundation.NSOperationQueue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("layoutQueue"))
	return foundation.NSOperationQueueFromID(objc.ID(rv))
}
func (t NSTextLayoutManager) SetLayoutQueue(value foundation.NSOperationQueue) {
	objc.Send[struct{}](t.ID, objc.Sel("setLayoutQueue:"), value)
}
// A callback block that the framework invokes whenever the text layout
// manager needs to validate the rendering attributes for the range.
//
// # Discussion
// 
// The validator uses [SetRenderingAttributesForTextRange] to fill the
// rendering attributes appropriate for the range inside `textLayoutFragment`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/renderingAttributesValidator
func (t NSTextLayoutManager) RenderingAttributesValidator() TextLayoutManagerTextLayoutFragmentHandler {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("renderingAttributesValidator"))
	_ = rv
	return nil
}
func (t NSTextLayoutManager) SetRenderingAttributesValidator(value TextLayoutManagerTextLayoutFragmentHandler) {
	block, cleanup := NewTextLayoutManagerTextLayoutFragmentBlock(value)
	defer cleanup()
	objc.Send[struct{}](t.ID, objc.Sel("setRenderingAttributesValidator:"), block)
}
// A Boolean value that controls whether the framework uses the leading
// information specified by the font when laying out text.
//
// # Discussion
// 
// If set to `true`, uses the leading as specified by the font. However, this
// isn’t appropriate for most UI text. Defaults to `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesFontLeading
func (t NSTextLayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFontLeading"))
	return rv
}
func (t NSTextLayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFontLeading:"), value)
}
// A Boolean values that controls whether the text layout manager attempts to
// hyphenate when wrapping lines.
//
// # Discussion
// 
// Defaults to `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesHyphenation
func (t NSTextLayoutManager) UsesHyphenation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesHyphenation"))
	return rv
}
func (t NSTextLayoutManager) SetUsesHyphenation(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesHyphenation:"), value)
}
// A Boolean value that controls internal security analysis for malicious
// inputs and activates defensive behaviors.
//
// # Discussion
// 
// By enabling this functionality, it’s possible that certain text — such
// as a very long paragraph — might result in an unexpected layout. Defaults
// to `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/limitsLayoutForSuspiciousContents
func (t NSTextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}
func (t NSTextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}
// The delegate for the text layout manager object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/delegate
func (t NSTextLayoutManager) Delegate() NSTextLayoutManagerDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextLayoutManagerDelegateObjectFromID(rv)
}
func (t NSTextLayoutManager) SetDelegate(value NSTextLayoutManagerDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// Returns the text content manager associated with this text layout manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContentManager
func (t NSTextLayoutManager) TextContentManager() INSTextContentManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textContentManager"))
	return NSTextContentManagerFromID(objc.ID(rv))
}
// The text container object that provides geometric information for the
// layout destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t NSTextLayoutManager) TextContainer() INSTextContainer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textContainer"))
	return NSTextContainerFromID(objc.ID(rv))
}
func (t NSTextLayoutManager) SetTextContainer(value INSTextContainer) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextContainer:"), value)
}
// Returns a text selection manager configured to have the text layout manager
// as its data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t NSTextLayoutManager) TextSelectionNavigation() INSTextSelectionNavigation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelectionNavigation"))
	return NSTextSelectionNavigationFromID(objc.ID(rv))
}
func (t NSTextLayoutManager) SetTextSelectionNavigation(value INSTextSelectionNavigation) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextSelectionNavigation:"), value)
}
// An array of text selections associated by the text layout manager.
//
// # Discussion
// 
// Each [NSTextSelection] represents an insertion point. The selection state
// that the framework shares among all viewports connected to the text layout
// manager through text containers.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelections
func (t NSTextLayoutManager) TextSelections() []NSTextSelection {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textSelections"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextSelection {
		return NSTextSelectionFromID(id)
	})
}
func (t NSTextLayoutManager) SetTextSelections(value []NSTextSelection) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextSelections:"), objectivec.IObjectSliceToNSArray(value))
}
// Returns the usage bounds for the text container.
//
// # Discussion
// 
// Views can observe this property to trigger a resize operation. For example,
// [NSView] calls [needsUpdateConstraints] when the usage bounds changes. This
// property is KVO-compliant.
//
// [needsUpdateConstraints]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usageBoundsForTextContainer
func (t NSTextLayoutManager) UsageBoundsForTextContainer() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("usageBoundsForTextContainer"))
	return corefoundation.CGRect(rv)
}
// The text viewport layout controller associated with the layout manager’s
// text container.
//
// # Discussion
// 
// The value of this property is `nil` if the text layout manager’s text
// container is `nil`. Set [TextContainer] before accessing this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textViewportLayoutController
func (t NSTextLayoutManager) TextViewportLayoutController() INSTextViewportLayoutController {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textViewportLayoutController"))
	return NSTextViewportLayoutControllerFromID(objc.ID(rv))
}
// Specifies the behavior for resolving `NSTextAlignment.Natural()` to the
// visual alignment.
//
// # Discussion
// 
// When set to `true`, the resolved visual alignment is determined by the
// resolved base writing direction; otherwise, it is using the user’s
// preferred language. The default value is `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/resolvesNaturalAlignmentWithBaseWritingDirection
func (t NSTextLayoutManager) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}
func (t NSTextLayoutManager) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
}
// Returns the starting and ending locations for the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionDataSource/documentRange
func (t NSTextLayoutManager) DocumentRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("documentRange"))
	return NSTextRangeFromID(objc.ID(rv))
}

// Returns the default set of attributes for rendering a link.
//
// # Discussion
// 
// The base [NSTextLayoutManager] class returns with [UnderlineStyleSingle]
// for [underlineStyle] and the platform link color for [foregroundColor]. The
// platform color for macOS is `linkColor`. Other platforms uses `blueColor`.
//
// [foregroundColor]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/foregroundColor
// [underlineStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/underlineStyle
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/linkRenderingAttributes
func (_NSTextLayoutManagerClass NSTextLayoutManagerClass) LinkRenderingAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSTextLayoutManagerClass.class), objc.Sel("linkRenderingAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

			// Protocol methods for NSTextSelectionDataSource
			

// EnumerateRenderingAttributesFromLocationReverseUsingBlockSync is a synchronous wrapper around [NSTextLayoutManager.EnumerateRenderingAttributesFromLocationReverseUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextLayoutManager) EnumerateRenderingAttributesFromLocationReverseUsingBlockSync(ctx context.Context, location NSTextLocation, reverse bool) error {
	done := make(chan struct{}, 1)
	t.EnumerateRenderingAttributesFromLocationReverseUsingBlock(location, reverse, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlockSync is a synchronous wrapper around [NSTextLayoutManager.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextLayoutManager) EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlockSync(ctx context.Context, location NSTextLocation, options NSTextLayoutFragmentEnumerationOptions) (*NSTextLayoutFragment, error) {
	done := make(chan *NSTextLayoutFragment, 1)
	t.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location, options, func(val *NSTextLayoutFragment) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

