// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTextInsertionIndicator] class.
var (
	_NSTextInsertionIndicatorClass     NSTextInsertionIndicatorClass
	_NSTextInsertionIndicatorClassOnce sync.Once
)

func getNSTextInsertionIndicatorClass() NSTextInsertionIndicatorClass {
	_NSTextInsertionIndicatorClassOnce.Do(func() {
		_NSTextInsertionIndicatorClass = NSTextInsertionIndicatorClass{class: objc.GetClass("NSTextInsertionIndicator")}
	})
	return _NSTextInsertionIndicatorClass
}

// GetNSTextInsertionIndicatorClass returns the class object for NSTextInsertionIndicator.
func GetNSTextInsertionIndicatorClass() NSTextInsertionIndicatorClass {
	return getNSTextInsertionIndicatorClass()
}

type NSTextInsertionIndicatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextInsertionIndicatorClass) Alloc() NSTextInsertionIndicator {
	rv := objc.Send[NSTextInsertionIndicator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that represents the insertion indicator in text.
//
// # Overview
// 
// [NSTextView] and [NSTextField] both use [NSTextInsertionIndicator] to
// display the insertion indicator. You can use this indicator if you have
// your own text engine or need to display an indicator elsewhere.
// 
// To use the indicator, instantiate an [NSTextInsertionIndicator], then add
// the view to your view hierarchy. Set the indicator view’s frame to where
// you want to display a text insertion indicator. The indicator has the same
// height as the indicator view’s frame, and centers horizontally within the
// indicator view’s frame.
// 
// The [NSTextInsertionIndicator.DisplayMode] specifies whether the indicator
// hides, remains visible, or blinks (automatic).
// 
// When set to [NSTextInsertionIndicator.DisplayMode.automatic], the indicator
// stops blinking when you set the frame. The indicator starts blinking when
// the frame doesn’t change for a period of time. When the user dictates,
// the indicator displays a trailing glow when it is moved.
// 
// Set the [NSTextInsertionIndicator.DisplayMode] to
// [NSTextInsertionIndicator.DisplayMode.automatic] when your custom view
// becomes the first responder. When your custom view resigns first responder,
// set the [NSTextInsertionIndicator.DisplayMode] to [NSTextInsertionIndicator.DisplayMode.hidden] to
// indicate that key events aren’t sent to your view.
// 
// By default the indicator’s color is [textInsertionPointColor]. You can
// set a different color.
//
// [NSTextInsertionIndicator.DisplayMode.automatic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/automatic
// [NSTextInsertionIndicator.DisplayMode.hidden]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/hidden
// [NSTextInsertionIndicator.DisplayMode]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum
// [textInsertionPointColor]: https://developer.apple.com/documentation/AppKit/NSColor/textInsertionPointColor
//
// # Configuring indicators
//
//   - [NSTextInsertionIndicator.Color]: The color of this indicator.
//   - [NSTextInsertionIndicator.SetColor]
//   - [NSTextInsertionIndicator.EffectsViewInserter]: An optional closure the system calls during dictation.
//   - [NSTextInsertionIndicator.SetEffectsViewInserter]
//
// # Setting the display mode
//
//   - [NSTextInsertionIndicator.DisplayMode]: A value that describes the display mode of an indicator.
//   - [NSTextInsertionIndicator.SetDisplayMode]
//   - [NSTextInsertionIndicator.AutomaticModeOptions]: Options that affect the automatic display mode.
//   - [NSTextInsertionIndicator.SetAutomaticModeOptions]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator
type NSTextInsertionIndicator struct {
	NSView
}

// NSTextInsertionIndicatorFromID constructs a [NSTextInsertionIndicator] from an objc.ID.
//
// A view that represents the insertion indicator in text.
func NSTextInsertionIndicatorFromID(id objc.ID) NSTextInsertionIndicator {
	return NSTextInsertionIndicator{NSView: NSViewFromID(id)}
}
// NOTE: NSTextInsertionIndicator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextInsertionIndicator] class.
//
// # Configuring indicators
//
//   - [INSTextInsertionIndicator.Color]: The color of this indicator.
//   - [INSTextInsertionIndicator.SetColor]
//   - [INSTextInsertionIndicator.EffectsViewInserter]: An optional closure the system calls during dictation.
//   - [INSTextInsertionIndicator.SetEffectsViewInserter]
//
// # Setting the display mode
//
//   - [INSTextInsertionIndicator.DisplayMode]: A value that describes the display mode of an indicator.
//   - [INSTextInsertionIndicator.SetDisplayMode]
//   - [INSTextInsertionIndicator.AutomaticModeOptions]: Options that affect the automatic display mode.
//   - [INSTextInsertionIndicator.SetAutomaticModeOptions]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator
type INSTextInsertionIndicator interface {
	INSView

	// Topic: Configuring indicators

	// The color of this indicator.
	Color() INSColor
	SetColor(value INSColor)
	// An optional closure the system calls during dictation.
	EffectsViewInserter() ViewHandler
	SetEffectsViewInserter(value ViewHandler)

	// Topic: Setting the display mode

	// A value that describes the display mode of an indicator.
	DisplayMode() NSTextInsertionIndicatorDisplayMode
	SetDisplayMode(value NSTextInsertionIndicatorDisplayMode)
	// Options that affect the automatic display mode.
	AutomaticModeOptions() NSTextInsertionIndicatorAutomaticModeOptions
	SetAutomaticModeOptions(value NSTextInsertionIndicatorAutomaticModeOptions)
}

// Init initializes the instance.
func (t NSTextInsertionIndicator) Init() NSTextInsertionIndicator {
	rv := objc.Send[NSTextInsertionIndicator](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextInsertionIndicator) Autorelease() NSTextInsertionIndicator {
	rv := objc.Send[NSTextInsertionIndicator](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextInsertionIndicator creates a new NSTextInsertionIndicator instance.
func NewNSTextInsertionIndicator() NSTextInsertionIndicator {
	class := getNSTextInsertionIndicatorClass()
	rv := objc.Send[NSTextInsertionIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewTextInsertionIndicatorWithCoder(coder foundation.INSCoder) NSTextInsertionIndicator {
	instance := getNSTextInsertionIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextInsertionIndicatorFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewTextInsertionIndicatorWithFrame(frameRect corefoundation.CGRect) NSTextInsertionIndicator {
	instance := getNSTextInsertionIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTextInsertionIndicatorFromID(rv)
}

// The color of this indicator.
//
// # Discussion
// 
// If set to [nil], returns [textInsertionPointColor]. Defaults to
// [textInsertionPointColor].
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
// [textInsertionPointColor]: https://developer.apple.com/documentation/AppKit/NSColor/textInsertionPointColor
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/color
func (t NSTextInsertionIndicator) Color() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("color"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextInsertionIndicator) SetColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setColor:"), value)
}
// An optional closure the system calls during dictation.
//
// # Discussion
// 
// Use this property to add the view that displays the trailing glow to the
// view hierarchy. The system calls the closure when it needs to display the
// glow effect view.
// 
// During dictation the indicator displays a glow effect above the text view
// and below the insertion indicator. It’s the closure’s responsibility to
// add the glow effect view to the view hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/effectsViewInserter
func (t NSTextInsertionIndicator) EffectsViewInserter() ViewHandler {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("effectsViewInserter"))
	_ = rv
	return nil
}
func (t NSTextInsertionIndicator) SetEffectsViewInserter(value ViewHandler) {
	block, cleanup := NewViewBlock(value)
	defer cleanup()
	objc.Send[struct{}](t.ID, objc.Sel("setEffectsViewInserter:"), block)
}
// A value that describes the display mode of an indicator.
//
// # Discussion
// 
// Set the [DisplayMode] to [NSTextInsertionIndicator.DisplayMode.automatic]
// when your custom view becomes the first responder. When your custom view
// resigns first responder, set the [DisplayMode] to
// [NSTextInsertionIndicator.DisplayMode.hidden] to indicate that key events
// aren’t sent to your view.
//
// [NSTextInsertionIndicator.DisplayMode.automatic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/automatic
// [NSTextInsertionIndicator.DisplayMode.hidden]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/hidden
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t NSTextInsertionIndicator) DisplayMode() NSTextInsertionIndicatorDisplayMode {
	rv := objc.Send[NSTextInsertionIndicatorDisplayMode](t.ID, objc.Sel("displayMode"))
	return NSTextInsertionIndicatorDisplayMode(rv)
}
func (t NSTextInsertionIndicator) SetDisplayMode(value NSTextInsertionIndicatorDisplayMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setDisplayMode:"), value)
}
// Options that affect the automatic display mode.
//
// # Discussion
// 
// Defaults to [TextInsertionIndicatorAutomaticModeOptionsShowEffectsView].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/automaticModeOptions-swift.property
func (t NSTextInsertionIndicator) AutomaticModeOptions() NSTextInsertionIndicatorAutomaticModeOptions {
	rv := objc.Send[NSTextInsertionIndicatorAutomaticModeOptions](t.ID, objc.Sel("automaticModeOptions"))
	return NSTextInsertionIndicatorAutomaticModeOptions(rv)
}
func (t NSTextInsertionIndicator) SetAutomaticModeOptions(value NSTextInsertionIndicatorAutomaticModeOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticModeOptions:"), value)
}

// See: https://developer.apple.com/documentation/appkit/nscolor/textinsertionpointcolor
func (_NSTextInsertionIndicatorClass NSTextInsertionIndicatorClass) TextInsertionPointColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSTextInsertionIndicatorClass.class), objc.Sel("textInsertionPointColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSTextInsertionIndicatorClass NSTextInsertionIndicatorClass) SetTextInsertionPointColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSTextInsertionIndicatorClass.class), objc.Sel("setTextInsertionPointColor:"), value)
}

