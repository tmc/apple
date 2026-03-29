// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSComboButton] class.
var (
	_NSComboButtonClass     NSComboButtonClass
	_NSComboButtonClassOnce sync.Once
)

func getNSComboButtonClass() NSComboButtonClass {
	_NSComboButtonClassOnce.Do(func() {
		_NSComboButtonClass = NSComboButtonClass{class: objc.GetClass("NSComboButton")}
	})
	return _NSComboButtonClass
}

// GetNSComboButtonClass returns the class object for NSComboButton.
func GetNSComboButtonClass() NSComboButtonClass {
	return getNSComboButtonClass()
}

type NSComboButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSComboButtonClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSComboButtonClass) Alloc() NSComboButton {
	rv := objc.Send[NSComboButton](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A button with a pull-down menu and a default action.
//
// # Overview
// 
// An [NSComboButton] object is a button that displays a title string, image,
// and an optional control for displaying a menu. Use this control in places
// where you want to offer a button with a default action and one or more
// alternative actions. Clicking the title or image executes the default
// action you provide, and clicking the menu control displays a menu for
// selecting a different action. If you configure the button to hide the menu
// control, a long-press gesture displays the menu.
// 
// After you create a combo button programmatically or in Interface Builder,
// choose the button [NSComboButton.Style] you want and add a title or image for your
// content. A combo button has a default action, which you specify at creation
// time. You can also change that action later using the inherited [NSComboButton.Target]
// and [NSComboButton.Action] properties. To specify one or more alternative actions,
// configure a menu with those actions and assign it to the button’s [Menu]
// property.
// 
// This control doesn’t use an [NSCell] object for its underlying
// implementation. It also doesn’t support the addition of a contextual
// menu.
//
// # Configuring the Button Appearance
//
//   - [NSComboButton.Style]: The appearance setting that determines how the button presents its menu .
//   - [NSComboButton.SetStyle]
//   - [NSComboButton.Title]: The localized string that the button displays.
//   - [NSComboButton.SetTitle]
//   - [NSComboButton.Image]: The image that the button displays.
//   - [NSComboButton.SetImage]
//   - [NSComboButton.ImageScaling]: The scaling behavior to apply to the button’s image.
//   - [NSComboButton.SetImageScaling]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton
type NSComboButton struct {
	NSControl
}

// NSComboButtonFromID constructs a [NSComboButton] from an objc.ID.
//
// A button with a pull-down menu and a default action.
func NSComboButtonFromID(id objc.ID) NSComboButton {
	return NSComboButton{NSControl: NSControlFromID(id)}
}
// NOTE: NSComboButton adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSComboButton] class.
//
// # Configuring the Button Appearance
//
//   - [INSComboButton.Style]: The appearance setting that determines how the button presents its menu .
//   - [INSComboButton.SetStyle]
//   - [INSComboButton.Title]: The localized string that the button displays.
//   - [INSComboButton.SetTitle]
//   - [INSComboButton.Image]: The image that the button displays.
//   - [INSComboButton.SetImage]
//   - [INSComboButton.ImageScaling]: The scaling behavior to apply to the button’s image.
//   - [INSComboButton.SetImageScaling]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton
type INSComboButton interface {
	INSControl

	// Topic: Configuring the Button Appearance

	// The appearance setting that determines how the button presents its menu .
	Style() NSComboButtonStyle
	SetStyle(value NSComboButtonStyle)
	// The localized string that the button displays.
	Title() string
	SetTitle(value string)
	// The image that the button displays.
	Image() INSImage
	SetImage(value INSImage)
	// The scaling behavior to apply to the button’s image.
	ImageScaling() NSImageScaling
	SetImageScaling(value NSImageScaling)
}

// Init initializes the instance.
func (c NSComboButton) Init() NSComboButton {
	rv := objc.Send[NSComboButton](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSComboButton) Autorelease() NSComboButton {
	rv := objc.Send[NSComboButton](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSComboButton creates a new NSComboButton instance.
func NewNSComboButton() NSComboButton {
	class := getNSComboButtonClass()
	rv := objc.Send[NSComboButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewComboButtonWithCoder(coder foundation.INSCoder) NSComboButton {
	instance := getNSComboButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSComboButtonFromID(rv)
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
func NewComboButtonWithFrame(frameRect corefoundation.CGRect) NSComboButton {
	instance := getNSComboButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSComboButtonFromID(rv)
}

// Creates a combo button that displays an image.
//
// image: The image to display in the button.
//
// menu: The menu to display when someone chooses an alternate action.
//
// target: The object that receives the default action message when someone clicks the
// button.
//
// action: The action message to send to the `target` object.
//
// # Return Value
// 
// A combo button configured with only the specified image.
//
// # Discussion
// 
// This method sets the [Title] property to an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/init(image:menu:target:action:)
func NewComboButtonWithImageMenuTargetAction(image INSImage, menu INSMenu, target objectivec.IObject, action objc.SEL) NSComboButton {
	rv := objc.Send[objc.ID](objc.ID(getNSComboButtonClass().class), objc.Sel("comboButtonWithImage:menu:target:action:"), image, menu, target, action)
	return NSComboButtonFromID(rv)
}

// Creates a combo button that displays both a title and image.
//
// title: The localized string to display in the button. Use the inherited
// [Alignment] property to set the text alignment for the string.
//
// image: The image to display in the button.
//
// menu: The menu to display when someone chooses an alternate action.
//
// target: The object that receives the default action message when someone clicks the
// button.
//
// action: The action message to send to the `target` object.
//
// # Return Value
// 
// A combo button configured with both a title and image.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/init(title:image:menu:target:action:)
func NewComboButtonWithTitleImageMenuTargetAction(title string, image INSImage, menu INSMenu, target objectivec.IObject, action objc.SEL) NSComboButton {
	rv := objc.Send[objc.ID](objc.ID(getNSComboButtonClass().class), objc.Sel("comboButtonWithTitle:image:menu:target:action:"), objc.String(title), image, menu, target, action)
	return NSComboButtonFromID(rv)
}

// Creates a combo button that displays a title.
//
// title: The localized string to display in the button. Use the inherited
// [Alignment] property to set the text alignment for the string.
//
// menu: The menu to display when someone chooses an alternate action.
//
// target: The object that receives the default action message when someone clicks the
// button.
//
// action: The action message to send to the `target` object.
//
// # Return Value
// 
// A combo button configured with only a title string.
//
// # Discussion
// 
// This method sets the [Image] property to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/init(title:menu:target:action:)
func NewComboButtonWithTitleMenuTargetAction(title string, menu INSMenu, target objectivec.IObject, action objc.SEL) NSComboButton {
	rv := objc.Send[objc.ID](objc.ID(getNSComboButtonClass().class), objc.Sel("comboButtonWithTitle:menu:target:action:"), objc.String(title), menu, target, action)
	return NSComboButtonFromID(rv)
}

// The appearance setting that determines how the button presents its menu .
//
// # Discussion
// 
// The default value of this property is [NSComboButton.Style.split].
//
// [NSComboButton.Style.split]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum/split
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/style-swift.property
func (c NSComboButton) Style() NSComboButtonStyle {
	rv := objc.Send[NSComboButtonStyle](c.ID, objc.Sel("style"))
	return NSComboButtonStyle(rv)
}
func (c NSComboButton) SetStyle(value NSComboButtonStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setStyle:"), value)
}
// The localized string that the button displays.
//
// # Discussion
// 
// The method you use to create the combo button sets the initial value of
// this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/title
func (c NSComboButton) Title() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (c NSComboButton) SetTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The image that the button displays.
//
// # Discussion
// 
// The combo button scales the image to fit within its bounds. Use the
// [ImageScaling] property to specify the scaling behavior to use with your
// image.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/image
func (c NSComboButton) Image() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (c NSComboButton) SetImage(value INSImage) {
	objc.Send[struct{}](c.ID, objc.Sel("setImage:"), value)
}
// The scaling behavior to apply to the button’s image.
//
// # Discussion
// 
// The default value of this property is
// [NSImageScaling.scaleProportionallyDown].
//
// [NSImageScaling.scaleProportionallyDown]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
//
// See: https://developer.apple.com/documentation/AppKit/NSComboButton/imageScaling
func (c NSComboButton) ImageScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](c.ID, objc.Sel("imageScaling"))
	return NSImageScaling(rv)
}
func (c NSComboButton) SetImageScaling(value NSImageScaling) {
	objc.Send[struct{}](c.ID, objc.Sel("setImageScaling:"), value)
}

