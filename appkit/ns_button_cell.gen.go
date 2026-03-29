// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSButtonCell] class.
var (
	_NSButtonCellClass     NSButtonCellClass
	_NSButtonCellClassOnce sync.Once
)

func getNSButtonCellClass() NSButtonCellClass {
	_NSButtonCellClassOnce.Do(func() {
		_NSButtonCellClass = NSButtonCellClass{class: objc.GetClass("NSButtonCell")}
	})
	return _NSButtonCellClass
}

// GetNSButtonCellClass returns the class object for NSButtonCell.
func GetNSButtonCellClass() NSButtonCellClass {
	return getNSButtonCellClass()
}

type NSButtonCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSButtonCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSButtonCellClass) Alloc() NSButtonCell {
	rv := objc.Send[NSButtonCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the user interface of a button or other clickable
// region of a view.
//
// # Overview
// 
// Setting the integer, float, double, or object value of an [NSButtonCell]
// object results in a call to [NSButtonCell.State] with the value converted to integer. In
// the case of [NSButtonCell.ObjectValue], `nil` is equivalent to `0`, and a non-`nil`
// object that doesn’t respond to [NSButtonCell.IntValue] sets the state to `1`.
// Otherwise, the state is set to the object’s [NSButtonCell.IntValue]. Similarly, for
// most button types, querying the integer, float, double, or object value of
// an [NSButtonCell] returns the current state in the requested
// representation. In the case of [NSButtonCell.ObjectValue], this is an [NSNumber]
// containing [true] for on, [false] for off, and integer value `-1` for the
// mixed state. For accelerator buttons (type [NSAcceleratorButton] or
// [NSMultiLevelAcceleratorButton]) on systems that support pressure
// sensitivity, querying [NSButtonCell.DoubleValue] returns the amount of pressure applied
// while pressing the button.
// 
// The configuration of an [NSButtonCell] object controls how the button
// object appears and behaves, but it’s [NSButton] that sends a message when
// the control is clicked. For more information on the behavior of
// [NSButtonCell], see the [NSButton] and [NSMatrix] class specifications, and
// [Button Programming Topics].
// 
// # Exceptions
// 
// In its implementation of the [Compare] method (declared in [NSCell]),
// [NSButtonCell] raises an [NSBadComparisonException] if the `otherCell`
// argument is not of the [NSButtonCell] class.
// 
// # Fonts
// 
// Setting the [Font] property does nothing if the button has no title or
// alternate title. If the button cell has a key equivalent, its font is not
// changed, but the key equivalent’s font size is changed to match the new
// title font.
//
// [Button Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Button/Button.html#//apple_ref/doc/uid/10000019i
// [NSAcceleratorButton]: https://developer.apple.com/documentation/AppKit/NSAcceleratorButton
// [NSMultiLevelAcceleratorButton]: https://developer.apple.com/documentation/AppKit/NSMultiLevelAcceleratorButton
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Setting Titles
//
//   - [NSButtonCell.AlternateTitle]: The string displayed by the button when it’s in its alternate state.
//   - [NSButtonCell.SetAlternateTitle]
//   - [NSButtonCell.AttributedAlternateTitle]: The title displayed by the button when it’s in its alternate state, as an attributed string.
//   - [NSButtonCell.SetAttributedAlternateTitle]
//   - [NSButtonCell.AttributedTitle]: The title displayed by the button when it’s in its normal state as an attributed string.
//   - [NSButtonCell.SetAttributedTitle]
//
// # Managing Images
//
//   - [NSButtonCell.AlternateImage]: The image the button displays in its alternate state.
//   - [NSButtonCell.SetAlternateImage]
//   - [NSButtonCell.ImagePosition]: The position of the button’s image relative to its title.
//   - [NSButtonCell.SetImagePosition]
//   - [NSButtonCell.ImageScaling]: The scale factor for the button’s image.
//   - [NSButtonCell.SetImageScaling]
//
// # Managing the Repeat Interval
//
//   - [NSButtonCell.SetPeriodicDelayInterval]: Sets the message delay and interval for the button.
//
// # Managing the Key Equivalent
//
//   - [NSButtonCell.KeyEquivalentModifierMask]: The mask that identifies the modifier keys for the button’s key equivalent.
//   - [NSButtonCell.SetKeyEquivalentModifierMask]
//
// # Managing Graphics Attributes
//
//   - [NSButtonCell.BackgroundColor]: The background color of the button.
//   - [NSButtonCell.SetBackgroundColor]
//   - [NSButtonCell.BezelStyle]: The appearance of the button’s border, if it has one.
//   - [NSButtonCell.SetBezelStyle]
//   - [NSButtonCell.ImageDimsWhenDisabled]: A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//   - [NSButtonCell.SetImageDimsWhenDisabled]
//   - [NSButtonCell.Transparent]: A Boolean value that indicates if the button is transparent.
//   - [NSButtonCell.SetTransparent]
//   - [NSButtonCell.ShowsBorderOnlyWhileMouseInside]: A Boolean value that indicates if the button displays its border only when the pointer is over it.
//   - [NSButtonCell.SetShowsBorderOnlyWhileMouseInside]
//
// # Displaying the Cell
//
//   - [NSButtonCell.HighlightsBy]: A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//   - [NSButtonCell.SetHighlightsBy]
//   - [NSButtonCell.SetButtonType]: Sets how the button highlights while pressed and how it shows its state.
//   - [NSButtonCell.ShowsStateBy]: The flags that indicate how the button cell shows its alternate state.
//   - [NSButtonCell.SetShowsStateBy]
//
// # Managing the Sound
//
//   - [NSButtonCell.Sound]: The sound that’s played when the user presses the button (that is during a mouse-down event).
//   - [NSButtonCell.SetSound]
//
// # Handling Events and Action Messages
//
//   - [NSButtonCell.MouseEntered]: Draws the button’s border.
//   - [NSButtonCell.MouseExited]: Erases the button’s border.
//
// # Drawing the Button Content
//
//   - [NSButtonCell.DrawBezelWithFrameInView]: Draws the border of the button using the current bezel style.
//   - [NSButtonCell.DrawImageWithFrameInView]: Draws the image associated with the button’s current state.
//   - [NSButtonCell.DrawTitleWithFrameInView]: Draws the button’s title centered vertically in a specified rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell
type NSButtonCell struct {
	NSActionCell
}

// NSButtonCellFromID constructs a [NSButtonCell] from an objc.ID.
//
// An object that defines the user interface of a button or other clickable
// region of a view.
func NSButtonCellFromID(id objc.ID) NSButtonCell {
	return NSButtonCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSButtonCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSButtonCell] class.
//
// # Setting Titles
//
//   - [INSButtonCell.AlternateTitle]: The string displayed by the button when it’s in its alternate state.
//   - [INSButtonCell.SetAlternateTitle]
//   - [INSButtonCell.AttributedAlternateTitle]: The title displayed by the button when it’s in its alternate state, as an attributed string.
//   - [INSButtonCell.SetAttributedAlternateTitle]
//   - [INSButtonCell.AttributedTitle]: The title displayed by the button when it’s in its normal state as an attributed string.
//   - [INSButtonCell.SetAttributedTitle]
//
// # Managing Images
//
//   - [INSButtonCell.AlternateImage]: The image the button displays in its alternate state.
//   - [INSButtonCell.SetAlternateImage]
//   - [INSButtonCell.ImagePosition]: The position of the button’s image relative to its title.
//   - [INSButtonCell.SetImagePosition]
//   - [INSButtonCell.ImageScaling]: The scale factor for the button’s image.
//   - [INSButtonCell.SetImageScaling]
//
// # Managing the Repeat Interval
//
//   - [INSButtonCell.SetPeriodicDelayInterval]: Sets the message delay and interval for the button.
//
// # Managing the Key Equivalent
//
//   - [INSButtonCell.KeyEquivalentModifierMask]: The mask that identifies the modifier keys for the button’s key equivalent.
//   - [INSButtonCell.SetKeyEquivalentModifierMask]
//
// # Managing Graphics Attributes
//
//   - [INSButtonCell.BackgroundColor]: The background color of the button.
//   - [INSButtonCell.SetBackgroundColor]
//   - [INSButtonCell.BezelStyle]: The appearance of the button’s border, if it has one.
//   - [INSButtonCell.SetBezelStyle]
//   - [INSButtonCell.ImageDimsWhenDisabled]: A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//   - [INSButtonCell.SetImageDimsWhenDisabled]
//   - [INSButtonCell.Transparent]: A Boolean value that indicates if the button is transparent.
//   - [INSButtonCell.SetTransparent]
//   - [INSButtonCell.ShowsBorderOnlyWhileMouseInside]: A Boolean value that indicates if the button displays its border only when the pointer is over it.
//   - [INSButtonCell.SetShowsBorderOnlyWhileMouseInside]
//
// # Displaying the Cell
//
//   - [INSButtonCell.HighlightsBy]: A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//   - [INSButtonCell.SetHighlightsBy]
//   - [INSButtonCell.SetButtonType]: Sets how the button highlights while pressed and how it shows its state.
//   - [INSButtonCell.ShowsStateBy]: The flags that indicate how the button cell shows its alternate state.
//   - [INSButtonCell.SetShowsStateBy]
//
// # Managing the Sound
//
//   - [INSButtonCell.Sound]: The sound that’s played when the user presses the button (that is during a mouse-down event).
//   - [INSButtonCell.SetSound]
//
// # Handling Events and Action Messages
//
//   - [INSButtonCell.MouseEntered]: Draws the button’s border.
//   - [INSButtonCell.MouseExited]: Erases the button’s border.
//
// # Drawing the Button Content
//
//   - [INSButtonCell.DrawBezelWithFrameInView]: Draws the border of the button using the current bezel style.
//   - [INSButtonCell.DrawImageWithFrameInView]: Draws the image associated with the button’s current state.
//   - [INSButtonCell.DrawTitleWithFrameInView]: Draws the button’s title centered vertically in a specified rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell
type INSButtonCell interface {
	INSActionCell

	// Topic: Setting Titles

	// The string displayed by the button when it’s in its alternate state.
	AlternateTitle() string
	SetAlternateTitle(value string)
	// The title displayed by the button when it’s in its alternate state, as an attributed string.
	AttributedAlternateTitle() foundation.NSAttributedString
	SetAttributedAlternateTitle(value foundation.NSAttributedString)
	// The title displayed by the button when it’s in its normal state as an attributed string.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)

	// Topic: Managing Images

	// The image the button displays in its alternate state.
	AlternateImage() INSImage
	SetAlternateImage(value INSImage)
	// The position of the button’s image relative to its title.
	ImagePosition() NSCellImagePosition
	SetImagePosition(value NSCellImagePosition)
	// The scale factor for the button’s image.
	ImageScaling() NSImageScaling
	SetImageScaling(value NSImageScaling)

	// Topic: Managing the Repeat Interval

	// Sets the message delay and interval for the button.
	SetPeriodicDelayInterval(delay float32, interval float32)

	// Topic: Managing the Key Equivalent

	// The mask that identifies the modifier keys for the button’s key equivalent.
	KeyEquivalentModifierMask() NSEventModifierFlags
	SetKeyEquivalentModifierMask(value NSEventModifierFlags)

	// Topic: Managing Graphics Attributes

	// The background color of the button.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// The appearance of the button’s border, if it has one.
	BezelStyle() NSBezelStyle
	SetBezelStyle(value NSBezelStyle)
	// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	// A Boolean value that indicates if the button is transparent.
	Transparent() bool
	SetTransparent(value bool)
	// A Boolean value that indicates if the button displays its border only when the pointer is over it.
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)

	// Topic: Displaying the Cell

	// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
	HighlightsBy() NSCellStyleMask
	SetHighlightsBy(value NSCellStyleMask)
	// Sets how the button highlights while pressed and how it shows its state.
	SetButtonType(type_ NSButtonType)
	// The flags that indicate how the button cell shows its alternate state.
	ShowsStateBy() NSCellStyleMask
	SetShowsStateBy(value NSCellStyleMask)

	// Topic: Managing the Sound

	// The sound that’s played when the user presses the button (that is during a mouse-down event).
	Sound() INSSound
	SetSound(value INSSound)

	// Topic: Handling Events and Action Messages

	// Draws the button’s border.
	MouseEntered(event INSEvent)
	// Erases the button’s border.
	MouseExited(event INSEvent)

	// Topic: Drawing the Button Content

	// Draws the border of the button using the current bezel style.
	DrawBezelWithFrameInView(frame corefoundation.CGRect, controlView INSView)
	// Draws the image associated with the button’s current state.
	DrawImageWithFrameInView(image INSImage, frame corefoundation.CGRect, controlView INSView)
	// Draws the button’s title centered vertically in a specified rectangle.
	DrawTitleWithFrameInView(title foundation.NSAttributedString, frame corefoundation.CGRect, controlView INSView) corefoundation.CGRect
}

// Init initializes the instance.
func (b NSButtonCell) Init() NSButtonCell {
	rv := objc.Send[NSButtonCell](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSButtonCell) Autorelease() NSButtonCell {
	rv := objc.Send[NSButtonCell](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSButtonCell creates a new NSButtonCell instance.
func NewNSButtonCell() NSButtonCell {
	class := getNSButtonCellClass()
	rv := objc.Send[NSButtonCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(imageCell:)
func NewButtonCellImageCell(image INSImage) NSButtonCell {
	instance := getNSButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSButtonCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(textCell:)
func NewButtonCellTextCell(string_ string) NSButtonCell {
	instance := getNSButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSButtonCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(coder:)
func NewButtonCellWithCoder(coder foundation.INSCoder) NSButtonCell {
	instance := getNSButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSButtonCellFromID(rv)
}

// Sets the message delay and interval for the button.
//
// delay: The amount of time (in seconds) that a continuous button will pause before
// starting to periodically send action messages to the target object.
// 
// The maximum value is 60.0 seconds; if a larger value is supplied, it’s
// ignored, and 60.0 seconds is used.
//
// interval: The amount of time (in seconds) between each action message.
// 
// The maximum value is 60.0 seconds; if a larger value is supplied, it’s
// ignored, and 60.0 seconds is used.
//
// # Discussion
// 
// These values are used if the button is configured (by a [Continuous]
// message) to continuously send the action message to the target object while
// tracking the mouse.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/setPeriodicDelay(_:interval:)
func (b NSButtonCell) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}
// Sets how the button highlights while pressed and how it shows its state.
//
// type: A constant specifying the type of button. This can be one of the constants
// defined in [NSButton.ButtonType].
// //
// [NSButton.ButtonType]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType
//
// # Discussion
// 
// The [SetButtonType] method redisplays the button before returning.
// 
// The types available are for the most common button types, which are also
// accessible in Interface Builder; you can configure different behavior with
// the [HighlightsBy] and [ShowsStateBy] properties.
// 
// Note that there is no `-buttonType` method. The set method sets various
// button properties that together establish the behavior of the type.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/setButtonType(_:)
func (b NSButtonCell) SetButtonType(type_ NSButtonType) {
	objc.Send[objc.ID](b.ID, objc.Sel("setButtonType:"), type_)
}
// Draws the button’s border.
//
// event: The event object generated by the mouse movement.
//
// # Discussion
// 
// This method is called only when the pointer moves onto the button and the
// value of [ShowsBorderOnlyWhileMouseInside] is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseEntered(with:)
func (b NSButtonCell) MouseEntered(event INSEvent) {
	objc.Send[objc.ID](b.ID, objc.Sel("mouseEntered:"), event)
}
// Erases the button’s border.
//
// event: The event object generated by the mouse movement.
//
// # Discussion
// 
// This method is called only when the pointer moves off the button and the
// value of [ShowsBorderOnlyWhileMouseInside] is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseExited(with:)
func (b NSButtonCell) MouseExited(event INSEvent) {
	objc.Send[objc.ID](b.ID, objc.Sel("mouseExited:"), event)
}
// Draws the border of the button using the current bezel style.
//
// frame: The bounding rectangle of the button.
//
// controlView: The control being drawn.
//
// # Discussion
// 
// This method is called automatically when the button is redrawn; you should
// not call it directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawBezel(withFrame:in:)
func (b NSButtonCell) DrawBezelWithFrameInView(frame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](b.ID, objc.Sel("drawBezelWithFrame:inView:"), frame, controlView)
}
// Draws the image associated with the button’s current state.
//
// image: The image associated with the button’s current state.
//
// frame: The bounding rectangle of the button.
//
// controlView: The control being drawn.
//
// # Discussion
// 
// This method is called automatically when the button is redrawn; you should
// not call it directly.
// 
// You specify the primary and alternate images for the button using Interface
// Builder.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawImage(_:withFrame:in:)
func (b NSButtonCell) DrawImageWithFrameInView(image INSImage, frame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](b.ID, objc.Sel("drawImage:withFrame:inView:"), image, frame, controlView)
}
// Draws the button’s title centered vertically in a specified rectangle.
//
// title: The title of the button.
//
// frame: The rectangle in which to draw the title.
//
// controlView: The control being drawn.
//
// # Return Value
// 
// The bounding rectangle for the text of the title.
//
// # Discussion
// 
// This method is called automatically when the button is redrawn; you should
// not call it directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawTitle(_:withFrame:in:)
func (b NSButtonCell) DrawTitleWithFrameInView(title foundation.NSAttributedString, frame corefoundation.CGRect, controlView INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("drawTitle:withFrame:inView:"), title, frame, controlView)
	return corefoundation.CGRect(rv)
}

// The string displayed by the button when it’s in its alternate state.
//
// # Discussion
// 
// The value of this property is the string that appears on the button when
// it’s in its alternate state, or the empty string if the button doesn’t
// display an alternate title. Note that some button types don’t display an
// alternate title. By default, a button’s alternate title is “Button.”
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateTitle
func (b NSButtonCell) AlternateTitle() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("alternateTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSButtonCell) SetAlternateTitle(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setAlternateTitle:"), objc.String(value))
}
// The title displayed by the button when it’s in its alternate state, as an
// attributed string.
//
// # Discussion
// 
// The value of this property is the attributed string that appears on the
// button when it’s in its alternate state, or the empty string if the
// button doesn’t display an alternate title. Note that some button types
// don’t display an alternate title. By default, a button’s alternate
// title is “Button.”
// 
// Graphics attributes that are set on the cell (such as `backgroundColor`,
// `alignment`, `font`, and so on) are overridden when corresponding
// properties are set for the attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedAlternateTitle
func (b NSButtonCell) AttributedAlternateTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attributedAlternateTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (b NSButtonCell) SetAttributedAlternateTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](b.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}
// The title displayed by the button when it’s in its normal state as an
// attributed string.
//
// # Discussion
// 
// The value of this property is the attributes string that appears on the
// button when it’s in its normal state, or an empty attributed string if
// the button doesn’t display a title. A button’s title is always
// displayed if the button doesn’t use its alternate contents for
// highlighting or displaying the alternate state. By default, a button’s
// title is “Button.” Setting this property redraws the button if
// necessary.
// 
// Graphics attributes configured for the cell (such as `backgroundColor`,
// `alignment`, `font`, and so on) are overridden when corresponding
// properties are set for the attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedTitle
func (b NSButtonCell) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (b NSButtonCell) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](b.ID, objc.Sel("setAttributedTitle:"), value)
}
// The image the button displays in its alternate state.
//
// # Discussion
// 
// The value of this property is the image displayed by the button when it’s
// in its alternate state, or `nil` if there is no alternate image. Note that
// some button types don’t display an alternate image. Buttons don’t
// display images by default. Setting this property may redraw the contents of
// the button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateImage
func (b NSButtonCell) AlternateImage() INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("alternateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (b NSButtonCell) SetAlternateImage(value INSImage) {
	objc.Send[struct{}](b.ID, objc.Sel("setAlternateImage:"), value)
}
// The position of the button’s image relative to its title.
//
// # Discussion
// 
// The value of this property is one of the image positions described in the
// “Constants” section of [NSCell]. If the title is above, below, or
// overlapping the image, or if there is no image, the text is horizontally
// centered within the button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/imagePosition
func (b NSButtonCell) ImagePosition() NSCellImagePosition {
	rv := objc.Send[NSCellImagePosition](b.ID, objc.Sel("imagePosition"))
	return NSCellImagePosition(rv)
}
func (b NSButtonCell) SetImagePosition(value NSCellImagePosition) {
	objc.Send[struct{}](b.ID, objc.Sel("setImagePosition:"), value)
}
// The scale factor for the button’s image.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageScaling
func (b NSButtonCell) ImageScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](b.ID, objc.Sel("imageScaling"))
	return NSImageScaling(rv)
}
func (b NSButtonCell) SetImageScaling(value NSImageScaling) {
	objc.Send[struct{}](b.ID, objc.Sel("setImageScaling:"), value)
}
// The mask that identifies the modifier keys for the button’s key
// equivalent.
//
// # Discussion
// 
// The value of this property is a mask that indicates the modifier keys that
// are applied to the button’s key equivalent. Mask bits are defined in
// `NSEvent.H()`. The only mask bits that are relevant in button
// key-equivalent modifier masks are [NSControlKeyMask], [NSAlternateKeyMask],
// and [NSCommandKeyMask] bits.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b NSButtonCell) KeyEquivalentModifierMask() NSEventModifierFlags {
	rv := objc.Send[NSEventModifierFlags](b.ID, objc.Sel("keyEquivalentModifierMask"))
	return NSEventModifierFlags(rv)
}
func (b NSButtonCell) SetKeyEquivalentModifierMask(value NSEventModifierFlags) {
	objc.Send[struct{}](b.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}
// The background color of the button.
//
// # Discussion
// 
// The background color is used only when drawing borderless buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/backgroundColor
func (b NSButtonCell) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSButtonCell) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setBackgroundColor:"), value)
}
// The appearance of the button’s border, if it has one.
//
// # Discussion
// 
// The value of this property is a constant that specifies the bezel style
// used by the button. See [NSButton.BezelStyle] for a list of possible
// values. If a button is borderless, the value of this property is ignored.
// 
// A button uses shading to look like it’s sticking out or pushed in. You
// can set the shading with the [gradientType] property.
//
// [NSButton.BezelStyle]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
// [gradientType]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/bezelStyle
func (b NSButtonCell) BezelStyle() NSBezelStyle {
	rv := objc.Send[NSBezelStyle](b.ID, objc.Sel("bezelStyle"))
	return NSBezelStyle(rv)
}
func (b NSButtonCell) SetBezelStyle(value NSBezelStyle) {
	objc.Send[struct{}](b.ID, objc.Sel("setBezelStyle:"), value)
}
// A Boolean value that indicates if the button’s image and text appear
// “dim” when the button is disabled.
//
// # Discussion
// 
// When the value of this property is [true], the button’s image and text
// are dimmed when the button is disabled; when it is [false], the image and
// text are not dimmed in the disabled state. By default, all button types
// except [NSSwitchButton] and [NSRadioButton] dim when disabled. When buttons
// of type [NSSwitchButton] and [NSRadioButton] are disabled, only the
// associated text dims.
// 
// The default setting for this state is reasserted whenever you invoke
// [SetButtonType], so be sure to specify the button cell’s type before you
// set [ImageDimsWhenDisabled].
//
// [NSRadioButton]: https://developer.apple.com/documentation/AppKit/NSRadioButton
// [NSSwitchButton]: https://developer.apple.com/documentation/AppKit/NSSwitchButton
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageDimsWhenDisabled
func (b NSButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("imageDimsWhenDisabled"))
	return rv
}
func (b NSButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setImageDimsWhenDisabled:"), value)
}
// A Boolean value that indicates if the button is transparent.
//
// # Discussion
// 
// When the value of this property is [true], the button is transparent, when
// it is [false], the button is not transparent. The default value is [false].
// 
// Setting this property redraws the button if necessary. A transparent button
// tracks the mouse and sends its action, but doesn’t draw. A transparent
// button is useful for sensitizing an area on the screen so that an action
// gets sent to a target when the area receives a mouse click.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/isTransparent
func (b NSButtonCell) Transparent() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isTransparent"))
	return rv
}
func (b NSButtonCell) SetTransparent(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setTransparent:"), value)
}
// A Boolean value that indicates if the button displays its border only when
// the pointer is over it.
//
// # Discussion
// 
// When the value of this property is [true] if the button’s border is
// displayed only when the pointer is over the button and the button is
// active. When it is [false], the border continues to display when the
// pointer is outside of the button’s bounds.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsBorderOnlyWhileMouseInside
func (b NSButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}
func (b NSButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}
// A set of flags that indicate how the button highlights when it receives a
// mouse-down event (that is, when the button is pressed).
//
// # Discussion
// 
// The value of this property is the logical OR of one or more of flags
// described in the “Constants” section of [NSCell].
// 
// If both [NSChangeGrayCellMask] and [NSChangeBackgroundCellMask] are
// specified, both are recorded, but the resulting behavior depends on the
// button cell’s image. If the button has no image, or if the image has no
// alpha (transparency) data, [NSChangeGrayCellMask] is used; if the image has
// alpha data, [NSChangeBackgroundCellMask] is used. This arrangement allows
// the color swap of the background to show through the image’s transparent
// pixels.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/highlightsBy
func (b NSButtonCell) HighlightsBy() NSCellStyleMask {
	rv := objc.Send[NSCellStyleMask](b.ID, objc.Sel("highlightsBy"))
	return NSCellStyleMask(rv)
}
func (b NSButtonCell) SetHighlightsBy(value NSCellStyleMask) {
	objc.Send[struct{}](b.ID, objc.Sel("setHighlightsBy:"), value)
}
// The flags that indicate how the button cell shows its alternate state.
//
// # Discussion
// 
// The value of this property is the logical [OR] of one or more of the cell
// masks described in the “Constants” section of [NSCell].
// 
// If both [NSChangeGrayCellMask] and [NSChangeBackgroundCellMask] are
// specified, both are recorded, but the actual behavior depends on the button
// cell’s image. If the button has no image, or if the image has no alpha
// (transparency) data, [NSChangeGrayCellMask] is used. If the image exists
// and has alpha data, [NSChangeBackgroundCellMask] is used; this arrangement
// allows the color swap of the background to show through the image’s
// transparent pixels.
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsStateBy
func (b NSButtonCell) ShowsStateBy() NSCellStyleMask {
	rv := objc.Send[NSCellStyleMask](b.ID, objc.Sel("showsStateBy"))
	return NSCellStyleMask(rv)
}
func (b NSButtonCell) SetShowsStateBy(value NSCellStyleMask) {
	objc.Send[struct{}](b.ID, objc.Sel("setShowsStateBy:"), value)
}
// The sound that’s played when the user presses the button (that is during
// a mouse-down event).
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/sound
func (b NSButtonCell) Sound() INSSound {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sound"))
	return NSSoundFromID(objc.ID(rv))
}
func (b NSButtonCell) SetSound(value INSSound) {
	objc.Send[struct{}](b.ID, objc.Sel("setSound:"), value)
}

