// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorWell] class.
var (
	_NSColorWellClass     NSColorWellClass
	_NSColorWellClassOnce sync.Once
)

func getNSColorWellClass() NSColorWellClass {
	_NSColorWellClassOnce.Do(func() {
		_NSColorWellClass = NSColorWellClass{class: objc.GetClass("NSColorWell")}
	})
	return _NSColorWellClass
}

// GetNSColorWellClass returns the class object for NSColorWell.
func GetNSColorWellClass() NSColorWellClass {
	return getNSColorWellClass()
}

type NSColorWellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorWellClass) Alloc() NSColorWell {
	rv := objc.Send[NSColorWell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A control that displays a color value and lets the user change that color
// value.
//
// # Overview
// 
// An [NSColorWell] object lets people select colors from your interface.
// Incorporate this type of control if your app supports custom color
// selection. For example, a drawing app might include a color well to let
// someone choose the color to use when drawing. A color well control displays
// the currently selected color, and interactions with the color well display
// interfaces for selecting new colors.
// 
// When you create a color well programmatically or in Interface Builder,
// specify the appearance and interaction style you want. The color well
// supports color selection using a color picker popover or the system
// [NSColorPanel] object. When someone selects a new color in one of these
// interfaces, the color well updates its selected color to match. You can
// also provide your own color selection process using a custom action and
// update the color yourself.
//
// # Managing the selected color
//
//   - [NSColorWell.Color]: The currently selected color for the color well.
//   - [NSColorWell.SetColor]
//   - [NSColorWell.TakeColorFrom]: Changes the currently selected color to the color of the specified object.
//   - [NSColorWell.SupportsAlpha]: A Boolean value that determines whether the color picker supports alpha values.
//   - [NSColorWell.SetSupportsAlpha]
//
// # Supporting high dynamic range (HDR) colors
//
//   - [NSColorWell.MaximumLinearExposure]: The maximum linear exposure a color in this color well can be set to. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked for this well may have a linear exposure applied to it.
//   - [NSColorWell.SetMaximumLinearExposure]
//
// # Configuring the appearance
//
//   - [NSColorWell.ColorWellStyle]: The appearance and interaction style to apply to the color well.
//   - [NSColorWell.SetColorWellStyle]
//   - [NSColorWell.Image]: The image to display on the button portion of a color well that adopts the expanded style.
//   - [NSColorWell.SetImage]
//   - [NSColorWell.Bordered]: A Boolean value that determines whether the color well has a border.
//   - [NSColorWell.SetBordered]
//
// # Activating and deactivating color wells
//
//   - [NSColorWell.Activate]: Activates the color well, displays the color panel, and synchronizes the two UI elements.
//   - [NSColorWell.Active]: A Boolean value that indicates whether the color well is currently active.
//   - [NSColorWell.Deactivate]: Deactivates the color well.
//
// # Drawing color wells
//
//   - [NSColorWell.DrawWellInside]: Draws the area inside the color well at the specified location without drawing borders.
//
// # Customizing the color selection behavior
//
//   - [NSColorWell.PulldownAction]: The action to perform when someone clicks in the color area of the color well.
//   - [NSColorWell.SetPulldownAction]
//   - [NSColorWell.PulldownTarget]: The target object that defines the action you want to perform when someone interacts with the color well.
//   - [NSColorWell.SetPulldownTarget]
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell
type NSColorWell struct {
	NSControl
}

// NSColorWellFromID constructs a [NSColorWell] from an objc.ID.
//
// A control that displays a color value and lets the user change that color
// value.
func NSColorWellFromID(id objc.ID) NSColorWell {
	return NSColorWell{NSControl: NSControlFromID(id)}
}
// NOTE: NSColorWell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSColorWell] class.
//
// # Managing the selected color
//
//   - [INSColorWell.Color]: The currently selected color for the color well.
//   - [INSColorWell.SetColor]
//   - [INSColorWell.TakeColorFrom]: Changes the currently selected color to the color of the specified object.
//   - [INSColorWell.SupportsAlpha]: A Boolean value that determines whether the color picker supports alpha values.
//   - [INSColorWell.SetSupportsAlpha]
//
// # Supporting high dynamic range (HDR) colors
//
//   - [INSColorWell.MaximumLinearExposure]: The maximum linear exposure a color in this color well can be set to. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked for this well may have a linear exposure applied to it.
//   - [INSColorWell.SetMaximumLinearExposure]
//
// # Configuring the appearance
//
//   - [INSColorWell.ColorWellStyle]: The appearance and interaction style to apply to the color well.
//   - [INSColorWell.SetColorWellStyle]
//   - [INSColorWell.Image]: The image to display on the button portion of a color well that adopts the expanded style.
//   - [INSColorWell.SetImage]
//   - [INSColorWell.Bordered]: A Boolean value that determines whether the color well has a border.
//   - [INSColorWell.SetBordered]
//
// # Activating and deactivating color wells
//
//   - [INSColorWell.Activate]: Activates the color well, displays the color panel, and synchronizes the two UI elements.
//   - [INSColorWell.Active]: A Boolean value that indicates whether the color well is currently active.
//   - [INSColorWell.Deactivate]: Deactivates the color well.
//
// # Drawing color wells
//
//   - [INSColorWell.DrawWellInside]: Draws the area inside the color well at the specified location without drawing borders.
//
// # Customizing the color selection behavior
//
//   - [INSColorWell.PulldownAction]: The action to perform when someone clicks in the color area of the color well.
//   - [INSColorWell.SetPulldownAction]
//   - [INSColorWell.PulldownTarget]: The target object that defines the action you want to perform when someone interacts with the color well.
//   - [INSColorWell.SetPulldownTarget]
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell
type INSColorWell interface {
	INSControl

	// Topic: Managing the selected color

	// The currently selected color for the color well.
	Color() INSColor
	SetColor(value INSColor)
	// Changes the currently selected color to the color of the specified object.
	TakeColorFrom(sender objectivec.IObject)
	// A Boolean value that determines whether the color picker supports alpha values.
	SupportsAlpha() bool
	SetSupportsAlpha(value bool)

	// Topic: Supporting high dynamic range (HDR) colors

	// The maximum linear exposure a color in this color well can be set to. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked for this well may have a linear exposure applied to it.
	MaximumLinearExposure() float64
	SetMaximumLinearExposure(value float64)

	// Topic: Configuring the appearance

	// The appearance and interaction style to apply to the color well.
	ColorWellStyle() NSColorWellStyle
	SetColorWellStyle(value NSColorWellStyle)
	// The image to display on the button portion of a color well that adopts the expanded style.
	Image() INSImage
	SetImage(value INSImage)
	// A Boolean value that determines whether the color well has a border.
	Bordered() bool
	SetBordered(value bool)

	// Topic: Activating and deactivating color wells

	// Activates the color well, displays the color panel, and synchronizes the two UI elements.
	Activate(exclusive bool)
	// A Boolean value that indicates whether the color well is currently active.
	Active() bool
	// Deactivates the color well.
	Deactivate()

	// Topic: Drawing color wells

	// Draws the area inside the color well at the specified location without drawing borders.
	DrawWellInside(insideRect corefoundation.CGRect)

	// Topic: Customizing the color selection behavior

	// The action to perform when someone clicks in the color area of the color well.
	PulldownAction() objc.SEL
	SetPulldownAction(value objc.SEL)
	// The target object that defines the action you want to perform when someone interacts with the color well.
	PulldownTarget() objectivec.IObject
	SetPulldownTarget(value objectivec.IObject)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSColorWell) Init() NSColorWell {
	rv := objc.Send[NSColorWell](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorWell) Autorelease() NSColorWell {
	rv := objc.Send[NSColorWell](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorWell creates a new NSColorWell instance.
func NewNSColorWell() NSColorWell {
	class := getNSColorWellClass()
	rv := objc.Send[NSColorWell](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewColorWellWithCoder(coder foundation.INSCoder) NSColorWell {
	instance := getNSColorWellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSColorWellFromID(rv)
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
func NewColorWellWithFrame(frameRect corefoundation.CGRect) NSColorWell {
	instance := getNSColorWellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSColorWellFromID(rv)
}


// Creates a color well that adopts the specified appearance style.
//
// style: The style to use to configure the color well control. For a list of
// possible values, see [NSColorWell.Style].
// //
// [NSColorWell.Style]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
//
// # Return Value
// 
// A color well configured with the specified style.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/init(style:)
func NewColorWellWithStyle(style NSColorWellStyle) NSColorWell {
	rv := objc.Send[objc.ID](objc.ID(getNSColorWellClass().class), objc.Sel("colorWellWithStyle:"), style)
	return NSColorWellFromID(rv)
}







// Changes the currently selected color to the color of the specified object.
//
// sender: The object from which to take the new color.
//
// # Discussion
// 
// This method attempts to access a property or accessor method named `color`.
// If the object doesn’t implement a `color` accessor, this method does
// nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/takeColorFrom(_:)
func (c NSColorWell) TakeColorFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeColorFrom:"), sender)
}

// Activates the color well, displays the color panel, and synchronizes the
// two UI elements.
//
// exclusive: [true] to deactivate any other color wells; [false] to keep them active. If
// a color panel is active with `exclusive` set to [true] and another is
// subsequently activated with `exclusive` set to [false], the exclusive
// setting of the first panel is ignored.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When you call this method, the color well displays the standard color panel
// and sets the panel’s current color to the value in the color well. When
// someone changes the color in the color panel, the color well updates its
// selected color to match. If the color well’s [Bordered] property is
// [true], the color well highlights that border while it’s active.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/activate(_:)
func (c NSColorWell) Activate(exclusive bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("activate:"), exclusive)
}

// Deactivates the color well.
//
// # Discussion
// 
// This method detaches the color well from the system color panel. Future
// selections in the color panel don’t update the color well’s current
// color.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/deactivate()
func (c NSColorWell) Deactivate() {
	objc.Send[objc.ID](c.ID, objc.Sel("deactivate"))
}

// Draws the area inside the color well at the specified location without
// drawing borders.
//
// insideRect: The rectangle specifying the area within which to draw.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/drawWell(inside:)
func (c NSColorWell) DrawWellInside(insideRect corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawWellInside:"), insideRect)
}
func (c NSColorWell) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The currently selected color for the color well.
//
// # Discussion
// 
// Use this property to get the currently selected color, or to set the
// current color programmatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/color
func (c NSColorWell) Color() INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("color"))
	return NSColorFromID(objc.ID(rv))
}
func (c NSColorWell) SetColor(value INSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setColor:"), value)
}



// A Boolean value that determines whether the color picker supports alpha
// values.
//
// # Discussion
// 
// If this property is [false], people can select only fully opaque colors
// from the color picker. A value of [false] also hides the alpha slider.
// Setting this property to [true] enables partial color opacity, and also
// makes the alpha slider visible.
// 
// If [IgnoresAlpha] is [true], this property always returns [false],
// disabling alpha globally.
// 
// By default this value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/supportsAlpha
func (c NSColorWell) SupportsAlpha() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsAlpha"))
	return rv
}
func (c NSColorWell) SetSupportsAlpha(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportsAlpha:"), value)
}



// The maximum linear exposure a color in this color well can be set to.
// Defaults to 1 and ignores any value less than 1. If set to a value >= 2,
// the color picked for this well may have a linear exposure applied to it.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/maximumLinearExposure
func (c NSColorWell) MaximumLinearExposure() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("maximumLinearExposure"))
	return rv
}
func (c NSColorWell) SetMaximumLinearExposure(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumLinearExposure:"), value)
}



// The appearance and interaction style to apply to the color well.
//
// # Discussion
// 
// The value of this property determines how the color well presents itself,
// and how interactions affect it. For details, see [NSColorWell.Style].
//
// [NSColorWell.Style]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/colorWellStyle
func (c NSColorWell) ColorWellStyle() NSColorWellStyle {
	rv := objc.Send[NSColorWellStyle](c.ID, objc.Sel("colorWellStyle"))
	return NSColorWellStyle(rv)
}
func (c NSColorWell) SetColorWellStyle(value NSColorWellStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setColorWellStyle:"), value)
}



// The image to display on the button portion of a color well that adopts the
// expanded style.
//
// # Discussion
// 
// The color well applies the image only when the [ColorWellStyle] property is
// set to [NSColorWell.Style.expanded].
//
// [NSColorWell.Style.expanded]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/image
func (c NSColorWell) Image() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (c NSColorWell) SetImage(value INSImage) {
	objc.Send[struct{}](c.ID, objc.Sel("setImage:"), value)
}



// A Boolean value that determines whether the color well has a border.
//
// # Discussion
// 
// If the value of this property is [true], the color well has a border; if
// it’s [false], the color well doesn’t have a border. The default value
// of this property is [true].
// 
// A borderless color well doesn’t display the Colors window when someone
// clicks it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/isBordered
func (c NSColorWell) Bordered() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBordered"))
	return rv
}
func (c NSColorWell) SetBordered(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setBordered:"), value)
}



// A Boolean value that indicates whether the color well is currently active.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/isActive
func (c NSColorWell) Active() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isActive"))
	return rv
}



// The action to perform when someone clicks in the color area of the color
// well.
//
// # Discussion
// 
// Specify a custom action method to replace the system popover and color
// picker. For a color well with the [NSColorWell.Style.minimal] or
// [NSColorWell.Style.expanded] style, clicks in the color area normally
// display a popover with the system color picker. If you specify a value for
// this property and the [PulldownTarget] property, clicks in the color area
// execute your custom action method instead.
//
// [NSColorWell.Style.expanded]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
// [NSColorWell.Style.minimal]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/minimal
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/pulldownAction
func (c NSColorWell) PulldownAction() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("pulldownAction"))
	return rv
}
func (c NSColorWell) SetPulldownAction(value objc.SEL) {
	objc.Send[struct{}](c.ID, objc.Sel("setPulldownAction:"), value)
}



// The target object that defines the action you want to perform when someone
// interacts with the color well.
//
// # Discussion
// 
// Specify a custom action method to replace the system popover and color
// picker. For a color well with the [NSColorWell.Style.minimal] or
// [NSColorWell.Style.expanded] style, clicks in the color area normally
// display a popover with the system color picker. If you specify a value for
// this property and the [PulldownAction] property, clicks in the color area
// execute your custom action method instead.
//
// [NSColorWell.Style.expanded]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
// [NSColorWell.Style.minimal]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/minimal
//
// See: https://developer.apple.com/documentation/AppKit/NSColorWell/pulldownTarget
func (c NSColorWell) PulldownTarget() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pulldownTarget"))
	return objectivec.Object{ID: rv}
}
func (c NSColorWell) SetPulldownTarget(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setPulldownTarget:"), value)
}



































