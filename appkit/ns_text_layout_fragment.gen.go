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

// The class instance for the [NSTextLayoutFragment] class.
var (
	_NSTextLayoutFragmentClass     NSTextLayoutFragmentClass
	_NSTextLayoutFragmentClassOnce sync.Once
)

func getNSTextLayoutFragmentClass() NSTextLayoutFragmentClass {
	_NSTextLayoutFragmentClassOnce.Do(func() {
		_NSTextLayoutFragmentClass = NSTextLayoutFragmentClass{class: objc.GetClass("NSTextLayoutFragment")}
	})
	return _NSTextLayoutFragmentClass
}

// GetNSTextLayoutFragmentClass returns the class object for NSTextLayoutFragment.
func GetNSTextLayoutFragmentClass() NSTextLayoutFragmentClass {
	return getNSTextLayoutFragmentClass()
}

type NSTextLayoutFragmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextLayoutFragmentClass) Alloc() NSTextLayoutFragment {
	rv := objc.Send[NSTextLayoutFragment](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A class that represents the layout fragment typically corresponding to a
// rendering surface, such as a layer or view subclass.
//
// # Creating a layout fragment
//
//   - [NSTextLayoutFragment.InitWithCoder]: Creates a new layout fragment with the coder you provide.
//   - [NSTextLayoutFragment.InitWithTextElementRange]: Create a new layout fragment using the provided text element and range.
//
// # Getting line fragments
//
//   - [NSTextLayoutFragment.TextLineFragments]: An array of text line fragments.
//   - [NSTextLayoutFragment.TextLineFragmentForTextLocationIsUpstreamAffinity]: Returns a text line fragment from a specific text location in the document.
//   - [NSTextLayoutFragment.TextLineFragmentForVerticalOffsetRequiresExactMatch]: Returns the text line fragment for the vertical offset you provide, or the closest text line fragment beyond the vertical offset.
//
// # Getting element information
//
//   - [NSTextLayoutFragment.State]: The layout information state.
//   - [NSTextLayoutFragment.RangeInElement]: The range inside the text element relative to the document origin.
//   - [NSTextLayoutFragment.TextElement]: The parent text element.
//
// # Accessing the layout manager
//
//   - [NSTextLayoutFragment.TextLayoutManager]: The layout manager for this text layout fragment.
//
// # Drawing the fragment and attachments
//
//   - [NSTextLayoutFragment.LayoutFragmentFrame]: The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.
//   - [NSTextLayoutFragment.RenderingSurfaceBounds]: The bounds defining the area required for rendering the contents.
//   - [NSTextLayoutFragment.DrawAtPointInContext]: Renders the visual representation of this element in the specified graphics context.
//   - [NSTextLayoutFragment.InvalidateLayout]: Invalidates any layout information associated with the text layout fragment.
//   - [NSTextLayoutFragment.TextAttachmentViewProviders]: The attachment view provider associated with the text layout fragment.
//   - [NSTextLayoutFragment.FrameForTextAttachmentAtLocation]: Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify.
//
// # Accessing the layout processing queue
//
//   - [NSTextLayoutFragment.LayoutQueue]: The queue on which the framework dispatches layout operations.
//   - [NSTextLayoutFragment.SetLayoutQueue]
//
// # Defining margins and padding
//
//   - [NSTextLayoutFragment.BottomMargin]: The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.
//   - [NSTextLayoutFragment.LeadingPadding]: The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.
//   - [NSTextLayoutFragment.TopMargin]: The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.
//   - [NSTextLayoutFragment.TrailingPadding]: The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment
type NSTextLayoutFragment struct {
	objectivec.Object
}

// NSTextLayoutFragmentFromID constructs a [NSTextLayoutFragment] from an objc.ID.
//
// A class that represents the layout fragment typically corresponding to a
// rendering surface, such as a layer or view subclass.
func NSTextLayoutFragmentFromID(id objc.ID) NSTextLayoutFragment {
	return NSTextLayoutFragment{objectivec.Object{id}}
}
// NOTE: NSTextLayoutFragment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextLayoutFragment] class.
//
// # Creating a layout fragment
//
//   - [INSTextLayoutFragment.InitWithCoder]: Creates a new layout fragment with the coder you provide.
//   - [INSTextLayoutFragment.InitWithTextElementRange]: Create a new layout fragment using the provided text element and range.
//
// # Getting line fragments
//
//   - [INSTextLayoutFragment.TextLineFragments]: An array of text line fragments.
//   - [INSTextLayoutFragment.TextLineFragmentForTextLocationIsUpstreamAffinity]: Returns a text line fragment from a specific text location in the document.
//   - [INSTextLayoutFragment.TextLineFragmentForVerticalOffsetRequiresExactMatch]: Returns the text line fragment for the vertical offset you provide, or the closest text line fragment beyond the vertical offset.
//
// # Getting element information
//
//   - [INSTextLayoutFragment.State]: The layout information state.
//   - [INSTextLayoutFragment.RangeInElement]: The range inside the text element relative to the document origin.
//   - [INSTextLayoutFragment.TextElement]: The parent text element.
//
// # Accessing the layout manager
//
//   - [INSTextLayoutFragment.TextLayoutManager]: The layout manager for this text layout fragment.
//
// # Drawing the fragment and attachments
//
//   - [INSTextLayoutFragment.LayoutFragmentFrame]: The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.
//   - [INSTextLayoutFragment.RenderingSurfaceBounds]: The bounds defining the area required for rendering the contents.
//   - [INSTextLayoutFragment.DrawAtPointInContext]: Renders the visual representation of this element in the specified graphics context.
//   - [INSTextLayoutFragment.InvalidateLayout]: Invalidates any layout information associated with the text layout fragment.
//   - [INSTextLayoutFragment.TextAttachmentViewProviders]: The attachment view provider associated with the text layout fragment.
//   - [INSTextLayoutFragment.FrameForTextAttachmentAtLocation]: Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify.
//
// # Accessing the layout processing queue
//
//   - [INSTextLayoutFragment.LayoutQueue]: The queue on which the framework dispatches layout operations.
//   - [INSTextLayoutFragment.SetLayoutQueue]
//
// # Defining margins and padding
//
//   - [INSTextLayoutFragment.BottomMargin]: The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.
//   - [INSTextLayoutFragment.LeadingPadding]: The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.
//   - [INSTextLayoutFragment.TopMargin]: The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.
//   - [INSTextLayoutFragment.TrailingPadding]: The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment
type INSTextLayoutFragment interface {
	objectivec.IObject

	// Topic: Creating a layout fragment

	// Creates a new layout fragment with the coder you provide.
	InitWithCoder(coder foundation.INSCoder) NSTextLayoutFragment
	// Create a new layout fragment using the provided text element and range.
	InitWithTextElementRange(textElement INSTextElement, rangeInElement INSTextRange) NSTextLayoutFragment

	// Topic: Getting line fragments

	// An array of text line fragments.
	TextLineFragments() []NSTextLineFragment
	// Returns a text line fragment from a specific text location in the document.
	TextLineFragmentForTextLocationIsUpstreamAffinity(textLocation NSTextLocation, isUpstreamAffinity bool) INSTextLineFragment
	// Returns the text line fragment for the vertical offset you provide, or the closest text line fragment beyond the vertical offset.
	TextLineFragmentForVerticalOffsetRequiresExactMatch(verticalOffset float64, requiresExactMatch bool) INSTextLineFragment

	// Topic: Getting element information

	// The layout information state.
	State() NSTextLayoutFragmentState
	// The range inside the text element relative to the document origin.
	RangeInElement() INSTextRange
	// The parent text element.
	TextElement() INSTextElement

	// Topic: Accessing the layout manager

	// The layout manager for this text layout fragment.
	TextLayoutManager() INSTextLayoutManager

	// Topic: Drawing the fragment and attachments

	// The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.
	LayoutFragmentFrame() corefoundation.CGRect
	// The bounds defining the area required for rendering the contents.
	RenderingSurfaceBounds() corefoundation.CGRect
	// Renders the visual representation of this element in the specified graphics context.
	DrawAtPointInContext(point corefoundation.CGPoint, context coregraphics.CGContextRef)
	// Invalidates any layout information associated with the text layout fragment.
	InvalidateLayout()
	// The attachment view provider associated with the text layout fragment.
	TextAttachmentViewProviders() []NSTextAttachmentViewProvider
	// Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify.
	FrameForTextAttachmentAtLocation(location NSTextLocation) corefoundation.CGRect

	// Topic: Accessing the layout processing queue

	// The queue on which the framework dispatches layout operations.
	LayoutQueue() foundation.NSOperationQueue
	SetLayoutQueue(value foundation.NSOperationQueue)

	// Topic: Defining margins and padding

	// The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.
	BottomMargin() float64
	// The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.
	LeadingPadding() float64
	// The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.
	TopMargin() float64
	// The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.
	TrailingPadding() float64

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTextLayoutFragment) Init() NSTextLayoutFragment {
	rv := objc.Send[NSTextLayoutFragment](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextLayoutFragment) Autorelease() NSTextLayoutFragment {
	rv := objc.Send[NSTextLayoutFragment](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextLayoutFragment creates a new NSTextLayoutFragment instance.
func NewNSTextLayoutFragment() NSTextLayoutFragment {
	class := getNSTextLayoutFragmentClass()
	rv := objc.Send[NSTextLayoutFragment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new layout fragment with the coder you provide.
//
// coder: A coder that conforms to [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(coder:)
func NewTextLayoutFragmentWithCoder(coder foundation.INSCoder) NSTextLayoutFragment {
	instance := getNSTextLayoutFragmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextLayoutFragmentFromID(rv)
}


// Create a new layout fragment using the provided text element and range.
//
// textElement: An [NSTextElement].
//
// rangeInElement: A range that defines the boundaries of the text for the new layout
// fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(textElement:range:)
func NewTextLayoutFragmentWithTextElementRange(textElement INSTextElement, rangeInElement INSTextRange) NSTextLayoutFragment {
	instance := getNSTextLayoutFragmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextElement:range:"), textElement, rangeInElement)
	return NSTextLayoutFragmentFromID(rv)
}







// Creates a new layout fragment with the coder you provide.
//
// coder: A coder that conforms to [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(coder:)
func (t NSTextLayoutFragment) InitWithCoder(coder foundation.INSCoder) NSTextLayoutFragment {
	rv := objc.Send[NSTextLayoutFragment](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Create a new layout fragment using the provided text element and range.
//
// textElement: An [NSTextElement].
//
// rangeInElement: A range that defines the boundaries of the text for the new layout
// fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(textElement:range:)
func (t NSTextLayoutFragment) InitWithTextElementRange(textElement INSTextElement, rangeInElement INSTextRange) NSTextLayoutFragment {
	rv := objc.Send[NSTextLayoutFragment](t.ID, objc.Sel("initWithTextElement:range:"), textElement, rangeInElement)
	return rv
}

// Returns a text line fragment from a specific text location in the document.
//
// textLocation: A text location that a text line fragment contains.
//
// isUpstreamAffinity: A Boolean value that indicates whether the text line fragment ends at the
// text location you provide.
//
// # Return Value
// 
// The text line fragment that contains or ends at the text location you
// provide, or `nil` if there isn’t a match.
//
// # Discussion
// 
// Set `isUpstreamAffinity` to [true] to find a text fragment by its element
// range end location, such as when you enumerate over line fragments in
// reverse order. Set `isUpstreamAffinity` to [false] to find a text fragment
// that contains `textLocation`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragment(for:isUpstreamAffinity:)
func (t NSTextLayoutFragment) TextLineFragmentForTextLocationIsUpstreamAffinity(textLocation NSTextLocation, isUpstreamAffinity bool) INSTextLineFragment {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLineFragmentForTextLocation:isUpstreamAffinity:"), textLocation, isUpstreamAffinity)
	return NSTextLineFragmentFromID(rv)
}

// Returns the text line fragment for the vertical offset you provide, or the
// closest text line fragment beyond the vertical offset.
//
// verticalOffset: A float value that indicates a vertical distance, expressed in points, from
// the layout fragment frame’s origin.
//
// requiresExactMatch: A Boolean value that indicates whether the method returns an exact match,
// or returns the closest match if there isn’t an exact match. The default
// value is [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A text line fragment, or `nil` if there isn’t a match.
//
// # Discussion
// 
// Set `requiresExactMatch` to [true] to find the text line fragment that
// contains the vertical offset, or set `requiresExactMatch` to [false] to
// find the closest text line fragment matching or beyond the vertical offset.
// Returns `nil` if there isn’t a match.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragment(forVerticalOffset:requiresExactMatch:)
func (t NSTextLayoutFragment) TextLineFragmentForVerticalOffsetRequiresExactMatch(verticalOffset float64, requiresExactMatch bool) INSTextLineFragment {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLineFragmentForVerticalOffset:requiresExactMatch:"), verticalOffset, requiresExactMatch)
	return NSTextLineFragmentFromID(rv)
}

// Renders the visual representation of this element in the specified graphics
// context.
//
// point: The origin as a [CGPoint].
// //
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// context: The rendering context.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/draw(at:in:)
func (t NSTextLayoutFragment) DrawAtPointInContext(point corefoundation.CGPoint, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawAtPoint:inContext:"), point, context)
}

// Invalidates any layout information associated with the text layout
// fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/invalidateLayout()
func (t NSTextLayoutFragment) InvalidateLayout() {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateLayout"))
}

// Returns the frame in the text layout fragment coordinate system for the
// attachment at the location you specify.
//
// location: The [NSTextLocation] that describes the location in the text layout
// fragment.
//
// # Return Value
// 
// The frame rectangle that describes the text layout fragment.
//
// # Discussion
// 
// Returns [CGRectZero] if `location` isn’t with any attachment or the state
// isn’t [NSTextLayoutFragment.State.layoutAvailable].
//
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
// [NSTextLayoutFragment.State.layoutAvailable]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/layoutAvailable
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/frameForTextAttachment(at:)
func (t NSTextLayoutFragment) FrameForTextAttachmentAtLocation(location NSTextLocation) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("frameForTextAttachmentAtLocation:"), location)
	return corefoundation.CGRect(rv)
}
func (t NSTextLayoutFragment) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// An array of text line fragments.
//
// # Discussion
// 
// Valid when [NSTextLayoutFragment.State.layoutAvailable]. This property is
// KVO-compliant.
//
// [NSTextLayoutFragment.State.layoutAvailable]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/layoutAvailable
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragments
func (t NSTextLayoutFragment) TextLineFragments() []NSTextLineFragment {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textLineFragments"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextLineFragment {
		return NSTextLineFragmentFromID(id)
	})
}



// The layout information state.
//
// # Discussion
// 
// This property is KVO-compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/state-swift.property
func (t NSTextLayoutFragment) State() NSTextLayoutFragmentState {
	rv := objc.Send[NSTextLayoutFragmentState](t.ID, objc.Sel("state"))
	return NSTextLayoutFragmentState(rv)
}



// The range inside the text element relative to the document origin.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/rangeInElement
func (t NSTextLayoutFragment) RangeInElement() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("rangeInElement"))
	return NSTextRangeFromID(objc.ID(rv))
}



// The parent text element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textElement
func (t NSTextLayoutFragment) TextElement() INSTextElement {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textElement"))
	return NSTextElementFromID(objc.ID(rv))
}



// The layout manager for this text layout fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLayoutManager
func (t NSTextLayoutFragment) TextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}



// The rectangle the framework uses for tiling the layout fragment inside the
// target layout coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/layoutFragmentFrame
func (t NSTextLayoutFragment) LayoutFragmentFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("layoutFragmentFrame"))
	return corefoundation.CGRect(rv)
}



// The bounds defining the area required for rendering the contents.
//
// # Discussion
// 
// The coordinate system is vertically flipped from the `layoutFragmentFrame`
// origin ({`0`,`0`} is at the upper-left corner). The size should be larger
// than `layoutFragmentFrame.Size()`. The origin could be in the negative
// coordinate since the rendering could stretch out of `layoutFragmentFrame`.
// Only valid when `state` greater than
// [NSTextLayoutFragment.State.estimatedUsageBounds].
//
// [NSTextLayoutFragment.State.estimatedUsageBounds]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/estimatedUsageBounds
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/renderingSurfaceBounds
func (t NSTextLayoutFragment) RenderingSurfaceBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("renderingSurfaceBounds"))
	return corefoundation.CGRect(rv)
}



// The attachment view provider associated with the text layout fragment.
//
// # Discussion
// 
// The property contents are only valid with
// [NSTextLayoutFragment.State.layoutAvailable].
//
// [NSTextLayoutFragment.State.layoutAvailable]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/layoutAvailable
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textAttachmentViewProviders
func (t NSTextLayoutFragment) TextAttachmentViewProviders() []NSTextAttachmentViewProvider {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textAttachmentViewProviders"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextAttachmentViewProvider {
		return NSTextAttachmentViewProviderFromID(id)
	})
}



// The queue on which the framework dispatches layout operations.
//
// # Discussion
// 
// If non-`nil`, the queue the framework uses for layout operations.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/layoutQueue
func (t NSTextLayoutFragment) LayoutQueue() foundation.NSOperationQueue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("layoutQueue"))
	return foundation.NSOperationQueueFromID(objc.ID(rv))
}
func (t NSTextLayoutFragment) SetLayoutQueue(value foundation.NSOperationQueue) {
	objc.Send[struct{}](t.ID, objc.Sel("setLayoutQueue:"), value)
}



// The amount of space reserved during paragraph layout between the bottom of
// the last line in the paragraph and the bottom of the text layout fragment.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/bottomMargin
func (t NSTextLayoutFragment) BottomMargin() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("bottomMargin"))
	return rv
}



// The amount of margin space reserved during paragraph layout between the
// leading edge of the text layout fragment and the start of the lines in the
// paragraph.
//
// # Discussion
// 
// This is with respect to according to the primary writing direction of the
// paragraph and the start of the lines in the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/leadingPadding
func (t NSTextLayoutFragment) LeadingPadding() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("leadingPadding"))
	return rv
}



// The amount of space reserved during paragraph layout between the top of the
// text layout fragment and the top of the first line in the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/topMargin
func (t NSTextLayoutFragment) TopMargin() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("topMargin"))
	return rv
}



// The amount of margin space reserved during paragraph layout between the end
// of the lines in the paragraph and the trailing edge of the text layout
// fragment.
//
// # Discussion
// 
// This is with respect to the primary writing direction of the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/trailingPadding
func (t NSTextLayoutFragment) TrailingPadding() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("trailingPadding"))
	return rv
}

























