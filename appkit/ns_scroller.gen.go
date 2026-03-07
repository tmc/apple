// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScroller] class.
var (
	_NSScrollerClass     NSScrollerClass
	_NSScrollerClassOnce sync.Once
)

func getNSScrollerClass() NSScrollerClass {
	_NSScrollerClassOnce.Do(func() {
		_NSScrollerClass = NSScrollerClass{class: objc.GetClass("NSScroller")}
	})
	return _NSScrollerClass
}

// GetNSScrollerClass returns the class object for NSScroller.
func GetNSScrollerClass() NSScrollerClass {
	return getNSScrollerClass()
}

type NSScrollerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrollerClass) Alloc() NSScroller {
	rv := objc.Send[NSScroller](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that controls scrolling of a document view within a scroll view
// or other type of container view.
//
// # Overview
// 
// A scroller displays a slot containing a knob that the user can drag
// directly to the desired location. The knob indicates both the position
// within the document view and—by varying in size within the slot—the
// amount visible relative to the size of the document view.
// 
// Typically, you don’t need to program with scrollers; instead, you
// configure them with an [NSScrollView] object in a [Nib file].
// 
// Don’t use an scroller when a slider would be more appropriate. An
// [NSSlider] object represents a range of values for something in the
// application and lets the user choose a setting. A scroller represents the
// relative position of the visible portion of a view and lets the user choose
// which portion to view.
//
// [Nib file]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/NibFile.html#//apple_ref/doc/uid/TP40008195-CH34
//
// # Setting the Knob Position
//
//   - [NSScroller.KnobProportion]: The proportion of the knob slot that the knob should fill.
//   - [NSScroller.SetKnobProportion]
//
// # Calculating Layout
//
//   - [NSScroller.RectForPart]: Returns the rectangle occupied by `aPart`, which for this method is interpreted literally rather than as an indicator of scrolling direction.
//   - [NSScroller.TestPart]: Returns the part that would be hit by a mouse-down event at `aPoint` (expressed in the window’s coordinate system).
//   - [NSScroller.CheckSpaceForParts]: Checks to see if there is enough room in the receiver to display the knob and buttons.
//   - [NSScroller.UsableParts]: A value that indicates which parts of the receiver are displayed and usable.
//
// # Drawing Scroller Parts
//
//   - [NSScroller.DrawKnobSlotInRectHighlight]: Draws the portion of the scroller’s track, possibly including the line increment and decrement arrow buttons, that falls in the given rectangle.
//   - [NSScroller.DrawKnob]: Draws the knob.
//
// # Event Handling
//
//   - [NSScroller.HitPart]: A part code indicating the manner in which the scrolling should be performed.
//   - [NSScroller.TrackKnob]: Tracks the knob and sends action messages to the receiver’s target.
//
// # Managing Presentation Style
//
//   - [NSScroller.ScrollerStyle]: The scroller style for this scroller.
//   - [NSScroller.SetScrollerStyle]
//   - [NSScroller.KnobStyle]: The scroller’s knob style.
//   - [NSScroller.SetKnobStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller
type NSScroller struct {
	NSControl
}

// NSScrollerFromID constructs a [NSScroller] from an objc.ID.
//
// An object that controls scrolling of a document view within a scroll view
// or other type of container view.
func NSScrollerFromID(id objc.ID) NSScroller {
	return NSScroller{NSControl: NSControlFromID(id)}
}
// NOTE: NSScroller adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSScroller] class.
//
// # Setting the Knob Position
//
//   - [INSScroller.KnobProportion]: The proportion of the knob slot that the knob should fill.
//   - [INSScroller.SetKnobProportion]
//
// # Calculating Layout
//
//   - [INSScroller.RectForPart]: Returns the rectangle occupied by `aPart`, which for this method is interpreted literally rather than as an indicator of scrolling direction.
//   - [INSScroller.TestPart]: Returns the part that would be hit by a mouse-down event at `aPoint` (expressed in the window’s coordinate system).
//   - [INSScroller.CheckSpaceForParts]: Checks to see if there is enough room in the receiver to display the knob and buttons.
//   - [INSScroller.UsableParts]: A value that indicates which parts of the receiver are displayed and usable.
//
// # Drawing Scroller Parts
//
//   - [INSScroller.DrawKnobSlotInRectHighlight]: Draws the portion of the scroller’s track, possibly including the line increment and decrement arrow buttons, that falls in the given rectangle.
//   - [INSScroller.DrawKnob]: Draws the knob.
//
// # Event Handling
//
//   - [INSScroller.HitPart]: A part code indicating the manner in which the scrolling should be performed.
//   - [INSScroller.TrackKnob]: Tracks the knob and sends action messages to the receiver’s target.
//
// # Managing Presentation Style
//
//   - [INSScroller.ScrollerStyle]: The scroller style for this scroller.
//   - [INSScroller.SetScrollerStyle]
//   - [INSScroller.KnobStyle]: The scroller’s knob style.
//   - [INSScroller.SetKnobStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller
type INSScroller interface {
	INSControl

	// Topic: Setting the Knob Position

	// The proportion of the knob slot that the knob should fill.
	KnobProportion() float64
	SetKnobProportion(value float64)

	// Topic: Calculating Layout

	// Returns the rectangle occupied by `aPart`, which for this method is interpreted literally rather than as an indicator of scrolling direction.
	RectForPart(partCode NSScrollerPart) corefoundation.CGRect
	// Returns the part that would be hit by a mouse-down event at `aPoint` (expressed in the window’s coordinate system).
	TestPart(point corefoundation.CGPoint) NSScrollerPart
	// Checks to see if there is enough room in the receiver to display the knob and buttons.
	CheckSpaceForParts()
	// A value that indicates which parts of the receiver are displayed and usable.
	UsableParts() NSUsableScrollerParts

	// Topic: Drawing Scroller Parts

	// Draws the portion of the scroller’s track, possibly including the line increment and decrement arrow buttons, that falls in the given rectangle.
	DrawKnobSlotInRectHighlight(slotRect corefoundation.CGRect, flag bool)
	// Draws the knob.
	DrawKnob()

	// Topic: Event Handling

	// A part code indicating the manner in which the scrolling should be performed.
	HitPart() NSScrollerPart
	// Tracks the knob and sends action messages to the receiver’s target.
	TrackKnob(event INSEvent)

	// Topic: Managing Presentation Style

	// The scroller style for this scroller.
	ScrollerStyle() NSScrollerStyle
	SetScrollerStyle(value NSScrollerStyle)
	// The scroller’s knob style.
	KnobStyle() NSScrollerKnobStyle
	SetKnobStyle(value NSScrollerKnobStyle)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSScroller) Init() NSScroller {
	rv := objc.Send[NSScroller](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScroller) Autorelease() NSScroller {
	rv := objc.Send[NSScroller](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScroller creates a new NSScroller instance.
func NewNSScroller() NSScroller {
	class := getNSScrollerClass()
	rv := objc.Send[NSScroller](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewScrollerWithCoder(coder foundation.INSCoder) NSScroller {
	instance := getNSScrollerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrollerFromID(rv)
}


// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewScrollerWithFrame(frameRect corefoundation.CGRect) NSScroller {
	instance := getNSScrollerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrollerFromID(rv)
}







// Returns the rectangle occupied by `aPart`, which for this method is
// interpreted literally rather than as an indicator of scrolling direction.
//
// # Discussion
// 
// See [NSScroller.Part] for a list of possible values for `aPart`.
// 
// Note the interpretations of [NSScrollerDecrementPage] and
// [NSScrollerIncrementPage]. The actual part of an NSScroller that causes
// page-by-page scrolling varies, so as a convenience these part codes refer
// to useful parts different from the scroll buttons.
// 
// Returns [NSZeroRect] if the part requested isn’t present on the receiver.
//
// [NSScroller.Part]: https://developer.apple.com/documentation/AppKit/NSScroller/Part
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/rect(for:)
func (s NSScroller) RectForPart(partCode NSScrollerPart) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rectForPart:"), partCode)
	return corefoundation.CGRect(rv)
}

// Returns the part that would be hit by a mouse-down event at `aPoint`
// (expressed in the window’s coordinate system).
//
// # Discussion
// 
// See [NSScroller.Part] for a list of possible return values. In macOS 10.7
// and later, this method no longer returns [NSScroller.Part.incrementLine] or
// [NSScroller.Part.decrementLine].
// 
// Note the interpretations of [NSScrollerDecrementPage] and
// [NSScrollerIncrementPage]. The actual part of a scroller that causes
// page-by-page scrolling varies, so as a convenience these part codes refer
// to useful parts different from the scroll buttons.
//
// [NSScroller.Part.decrementLine]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/decrementLine
// [NSScroller.Part.incrementLine]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/incrementLine
// [NSScroller.Part]: https://developer.apple.com/documentation/AppKit/NSScroller/Part
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/testPart(_:)
func (s NSScroller) TestPart(point corefoundation.CGPoint) NSScrollerPart {
	rv := objc.Send[NSScrollerPart](s.ID, objc.Sel("testPart:"), point)
	return NSScrollerPart(rv)
}

// Checks to see if there is enough room in the receiver to display the knob
// and buttons.
//
// # Discussion
// 
// The [UsableParts] property contains the state calculated by this method.
// You should never need to invoke this method; it’s invoked automatically
// whenever the scroller’s size changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/checkSpaceForParts()
func (s NSScroller) CheckSpaceForParts() {
	objc.Send[objc.ID](s.ID, objc.Sel("checkSpaceForParts"))
}

// Draws the portion of the scroller’s track, possibly including the line
// increment and decrement arrow buttons, that falls in the given rectangle.
//
// slotRect: The rectangle in which to draw the knob slot.
//
// flag: If `flag` is [true], any scroll arrow button that falls within `slotRect`
// is drawn highlighted; otherwise it’s drawn normally.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Only one arrow button will be shown highlighted at a time, so you can
// expect this method to sometimes be invoked with a `slotRect` that
// encompasses only one arrow.
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/drawKnobSlot(in:highlight:)
func (s NSScroller) DrawKnobSlotInRectHighlight(slotRect corefoundation.CGRect, flag bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawKnobSlotInRect:highlight:"), slotRect, flag)
}

// Draws the knob.
//
// # Discussion
// 
// You should never need to invoke this method directly, but may wish to
// override it to customize the appearance of the knob.
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/drawKnob()
func (s NSScroller) DrawKnob() {
	objc.Send[objc.ID](s.ID, objc.Sel("drawKnob"))
}

// Tracks the knob and sends action messages to the receiver’s target.
//
// # Discussion
// 
// This method is invoked automatically when the receiver receives `theEvent`
// mouse-down event in the knob; you should not invoke it directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/trackKnob(with:)
func (s NSScroller) TrackKnob(event INSEvent) {
	objc.Send[objc.ID](s.ID, objc.Sel("trackKnob:"), event)
}
func (s NSScroller) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Returns the width for scrollers of the receiving class for a given control
// size and scroller style.
//
// controlSize: A control size.
//
// scrollerStyle: A scroller style.
//
// # Return Value
// 
// The width for scrollers of the receiving class for `controlSize` and
// `scrollerStyle`.
//
// # Discussion
// 
// You should use this method in preference to [scrollerWidthForControlSize:],
// which assumes a scroller style of [NSScroller.Style.legacy], and
// [scrollerWidth] which in addition assumes a control size of
// [NSRegularControlSize].
//
// [NSRegularControlSize]: https://developer.apple.com/documentation/AppKit/NSRegularControlSize
// [NSScroller.Style.legacy]: https://developer.apple.com/documentation/AppKit/NSScroller/Style/legacy
// [scrollerWidthForControlSize:]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidthForControlSize:
// [scrollerWidth]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidth
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidth(for:scrollerStyle:)
func (_NSScrollerClass NSScrollerClass) ScrollerWidthForControlSizeScrollerStyle(controlSize NSControlSize, scrollerStyle NSScrollerStyle) float64 {
	rv := objc.Send[float64](objc.ID(_NSScrollerClass.class), objc.Sel("scrollerWidthForControlSize:scrollerStyle:"), controlSize, scrollerStyle)
	return rv
}








// A value that indicates which parts of the receiver are displayed and
// usable.
//
// # Discussion
// 
// See [NSScroller.UsableParts] for a list of possible values.
//
// [NSScroller.UsableParts]: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/usableParts-swift.property
func (s NSScroller) UsableParts() NSUsableScrollerParts {
	rv := objc.Send[NSUsableScrollerParts](s.ID, objc.Sel("usableParts"))
	return NSUsableScrollerParts(rv)
}



// A part code indicating the manner in which the scrolling should be
// performed.
//
// # Discussion
// 
// This method is typically invoked by an [NSScrollView] object to determine
// how to scroll its document view when it receives an action message from the
// scroller.
// 
// See [NSScroller.Part] for a list of part codes. In macOS 10.7 and later,
// this method no longer returns [NSScroller.Part.incrementLine] or
// [NSScroller.Part.decrementLine].
//
// [NSScroller.Part.decrementLine]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/decrementLine
// [NSScroller.Part.incrementLine]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/incrementLine
// [NSScroller.Part]: https://developer.apple.com/documentation/AppKit/NSScroller/Part
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/hitPart
func (s NSScroller) HitPart() NSScrollerPart {
	rv := objc.Send[NSScrollerPart](s.ID, objc.Sel("hitPart"))
	return NSScrollerPart(rv)
}



// The scroller style for this scroller.
//
// # Discussion
// 
// For a scroller that’s managed by an [NSScrollView] object, the setter is
// automatically invoked by the scroll view with the appropriate setting,
// according to the user’s Appearance preference settings and possibly what
// pointing device(s) are present (see [PreferredScrollerStyle]).
// 
// For a list of valid scroller styles, see [NSScroller.Style].
//
// [NSScroller.Style]: https://developer.apple.com/documentation/AppKit/NSScroller/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerStyle
func (s NSScroller) ScrollerStyle() NSScrollerStyle {
	rv := objc.Send[NSScrollerStyle](s.ID, objc.Sel("scrollerStyle"))
	return NSScrollerStyle(rv)
}
func (s NSScroller) SetScrollerStyle(value NSScrollerStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollerStyle:"), value)
}



// The scroller’s knob style.
//
// # Discussion
// 
// The value of this property does not affect legacy scrollers.
// [NSScroller.KnobStyle.default] is appropriate for a wide range of content,
// but in some cases choosing an alternative knob style may enhance visibility
// of the scroller knob atop some kinds of content.
// 
// For a list of possible values, see [NSScroller.KnobStyle].
//
// [NSScroller.KnobStyle.default]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum/default
// [NSScroller.KnobStyle]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/knobStyle-swift.property
func (s NSScroller) KnobStyle() NSScrollerKnobStyle {
	rv := objc.Send[NSScrollerKnobStyle](s.ID, objc.Sel("knobStyle"))
	return NSScrollerKnobStyle(rv)
}
func (s NSScroller) SetKnobStyle(value NSScrollerKnobStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setKnobStyle:"), value)
}



// The proportion of the knob slot that the knob should fill.
//
// # Discussion
// 
// This property contains a floating-point value from 0.0 (minimal size) to
// 1.0 (fills the slot).
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/knobProportion
func (s NSScroller) KnobProportion() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("knobProportion"))
	return rv
}
func (s NSScroller) SetKnobProportion(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setKnobProportion:"), value)
}







// Returns the style of scrollers that applications should use wherever
// possible.
//
// # Return Value
// 
// The style of scrollers that applications should use wherever possible.
// 
// # Discussion
// 
// The preferred scroller style is determined by the Appearance preference
// panel’s “Show scroll bars” setting for the current user, and—when
// the user’s preference is set to “Automatically based on input
// device”—by the set of built-in and connected pointing devices and the
// user’s scroll capability preference settings for them. The preferred
// scroller style may therefore change over time, and applications should be
// prepared to adapt their user interfaces to the new scroller style if
// needed.
// 
// In most cases, updating to a new scroller style is automatic: When the
// preferred scroller style changes, AppKit notifies all [NSScrollView]
// instances, setting the [ScrollerStyle] property of each with the new style,
// which causes each scroll view to automatically re-tile (update its layout)
// to adapt to the new scroller style. Some [NSScrollView] instances may
// refuse the new scroller style setting if they cannot accommodate it for
// compatibility reasons (the presence of accessory views or legacy scroller
// subclasses prevent use of overlay scrollers), but most instances will
// switch to the specified new preferred scroller style.
// 
// If you need to be notified of changes to the preferred scroller style, you
// can register to receive [preferredScrollerStyleDidChangeNotification]
// notifications.
//
// [preferredScrollerStyleDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyleDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyle
func (_NSScrollerClass NSScrollerClass) PreferredScrollerStyle() NSScrollerStyle {
	rv := objc.Send[NSScrollerStyle](objc.ID(_NSScrollerClass.class), objc.Sel("preferredScrollerStyle"))
	return NSScrollerStyle(rv)
}



// Posted if the preferred scroller style changes.
//
// See: https://developer.apple.com/documentation/appkit/nsscroller/preferredscrollerstyledidchangenotification
func (_NSScrollerClass NSScrollerClass) PreferredScrollerStyleDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSScrollerClass.class), objc.Sel("NSPreferredScrollerStyleDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppKit/NSScroller/isCompatibleWithOverlayScrollers
func (_NSScrollerClass NSScrollerClass) CompatibleWithOverlayScrollers() bool {
	rv := objc.Send[bool](objc.ID(_NSScrollerClass.class), objc.Sel("isCompatibleWithOverlayScrollers"))
	return rv
}































