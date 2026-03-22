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

// The class instance for the [NSButton] class.
var (
	_NSButtonClass     NSButtonClass
	_NSButtonClassOnce sync.Once
)

func getNSButtonClass() NSButtonClass {
	_NSButtonClassOnce.Do(func() {
		_NSButtonClass = NSButtonClass{class: objc.GetClass("NSButton")}
	})
	return _NSButtonClass
}

// GetNSButtonClass returns the class object for NSButton.
func GetNSButtonClass() NSButtonClass {
	return getNSButtonClass()
}

type NSButtonClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSButtonClass) Alloc() NSButton {
	rv := objc.Send[NSButton](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A control that defines an area on the screen that a user clicks to trigger
// an action.
//
// # Overview
// 
// Buttons are a standard control for initiating actions within your app. You
// can configure buttons with many different visual styles, but the behavior
// is the same. When a user clicks it, a button calls the action method of its
// associated target object. (If you configure a button as continuous, it
// calls its action method at timed intervals until the user releases the
// mouse button or the cursor leaves the button boundaries). You use the
// action method to perform your app-specific tasks.
// 
// There are multiple types of buttons, each with a different user interface
// and behavior. The [NSButtonCell] class defines the button types, and
// calling the [NSButton.SetButtonType] method configures them.
// 
// If you configure it as an accelerator button (type [NSAcceleratorButton] or
// [NSMultiLevelAcceleratorButton]), you can set a button to send action
// messages when changes in pressure occur when the user clicks the button.
// 
// Buttons can either have two states (on and off) or three states (on, off,
// and mixed). You enable a three-state button by calling the
// [NSButton.AllowsMixedState] method. On and off (also referred to as alternate and
// normal) states indicate that the user clicked or didn’t click the button.
// Mixed is typically used for checkboxes or radio buttons, which allow for an
// additional intermediate state. For example, suppose the state of a checkbox
// denotes whether a text field contains bold text. If all text in the text
// field is bold, then the checkbox is on. If none of the text is bold, then
// the checkbox is off. If some of the text is bold, then the checkbox is
// mixed.
// 
// For most types of buttons, the value of the button matches its state—the
// value is `1` for on, `0` for off, or `-1` for mixed. For pressure-sensitive
// buttons, the value of the button indicates pressure level instead.
// 
// [NSButton] and [NSMatrix] both provide a control view, which displays an
// [NSButtonCell] object. However, while a matrix requires you to access the
// button cell objects directly, most button class methods act as “covers”
// for identically declared button cell methods. In other words, the
// implementation of the button method invokes the corresponding button cell
// method for you, allowing you to be unconcerned with the existence of the
// button cell. The only button cell methods that don’t have covers relate
// to the font you use to display the key equivalent and to specific methods
// for highlighting or showing the state of the button.
//
// [NSAcceleratorButton]: https://developer.apple.com/documentation/AppKit/NSAcceleratorButton
// [NSMultiLevelAcceleratorButton]: https://developer.apple.com/documentation/AppKit/NSMultiLevelAcceleratorButton
//
// # Configuring buttons
//
//   - [NSButton.SetButtonType]: Sets the button’s type, which affects its user interface and behavior when clicked.
//   - [NSButton.GetPeriodicDelayInterval]: Returns by reference the delay and interval periods for a continuous button.
//   - [NSButton.SetPeriodicDelayInterval]: Sets the message delay and interval periods for a continuous button.
//   - [NSButton.ContentTintColor]: A tint color to use for the template image and text content.
//   - [NSButton.SetContentTintColor]
//   - [NSButton.HasDestructiveAction]: A Boolean value that defines whether a button’s action has a destructive effect.
//   - [NSButton.SetHasDestructiveAction]
//   - [NSButton.AlternateTitle]: The title that the button displays when the button is in an on state.
//   - [NSButton.SetAlternateTitle]
//   - [NSButton.AttributedTitle]: The title that the button displays in an off state, as an attributed string.
//   - [NSButton.SetAttributedTitle]
//   - [NSButton.AttributedAlternateTitle]: The title that the button displays as an attributed string when the button is in an on state.
//   - [NSButton.SetAttributedAlternateTitle]
//   - [NSButton.Title]: The title displayed on the button when it’s in an off state.
//   - [NSButton.SetTitle]
//   - [NSButton.SymbolConfiguration]: The combination of point size, weight, and scale to use when sizing and displaying symbol images.
//   - [NSButton.SetSymbolConfiguration]
//   - [NSButton.Sound]: The sound that plays when the user clicks the button.
//   - [NSButton.SetSound]
//   - [NSButton.SpringLoaded]: A Boolean value that indicates whether spring loading is enabled for the button.
//   - [NSButton.SetSpringLoaded]
//   - [NSButton.MaxAcceleratorLevel]: An integer value indicating the maximum pressure level for a button of type [NSMultiLevelAcceleratorButton](<doc://com.apple.appkit/documentation/AppKit/NSMultiLevelAcceleratorButton>).
//   - [NSButton.SetMaxAcceleratorLevel]
//   - [NSButton.TintProminence]: The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
//   - [NSButton.SetTintProminence]
//   - [NSButton.BorderShape]
//   - [NSButton.SetBorderShape]
//
// # Configuring button images
//
//   - [NSButton.Image]: The image that appears on the button when it’s in an off state, or `nil` if there is no such image.
//   - [NSButton.SetImage]
//   - [NSButton.AlternateImage]: An alternate image that appears on the button when the button is in an on state.
//   - [NSButton.SetAlternateImage]
//   - [NSButton.ImagePosition]: The position of the button’s image relative to its title.
//   - [NSButton.SetImagePosition]
//   - [NSButton.Bordered]: A Boolean value that determines whether the button has a border.
//   - [NSButton.SetBordered]
//   - [NSButton.Transparent]: A Boolean value that indicates whether the button is transparent.
//   - [NSButton.SetTransparent]
//   - [NSButton.BezelStyle]: The appearance of the button’s border.
//   - [NSButton.SetBezelStyle]
//   - [NSButton.BezelColor]: The color of the button’s bezel, in appearances that support it.
//   - [NSButton.SetBezelColor]
//   - [NSButton.ShowsBorderOnlyWhileMouseInside]: A Boolean value that determines whether the button displays its border only when the pointer is over it.
//   - [NSButton.SetShowsBorderOnlyWhileMouseInside]
//   - [NSButton.ImageHugsTitle]: A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//   - [NSButton.SetImageHugsTitle]
//   - [NSButton.ImageScaling]: The scaling mode applied to make the cell’s image fit the frame of the image view.
//   - [NSButton.SetImageScaling]
//
// # Managing button state
//
//   - [NSButton.AllowsMixedState]: A Boolean value that indicates whether the button allows a mixed state.
//   - [NSButton.SetAllowsMixedState]
//   - [NSButton.State]: The button’s state.
//   - [NSButton.SetState]
//   - [NSButton.SetNextState]: Sets the button to its next state.
//   - [NSButton.Highlight]: Highlights (or unhighlights) the button.
//
// # Accessing key equivalents
//
//   - [NSButton.KeyEquivalent]: The key-equivalent character of the button.
//   - [NSButton.SetKeyEquivalent]
//   - [NSButton.KeyEquivalentModifierMask]: The mask specifying the modifier keys for the button’s key equivalent.
//   - [NSButton.SetKeyEquivalentModifierMask]
//
// See: https://developer.apple.com/documentation/AppKit/NSButton
type NSButton struct {
	NSControl
}

// NSButtonFromID constructs a [NSButton] from an objc.ID.
//
// A control that defines an area on the screen that a user clicks to trigger
// an action.
func NSButtonFromID(id objc.ID) NSButton {
	return NSButton{NSControl: NSControlFromID(id)}
}
// NOTE: NSButton adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSButton] class.
//
// # Configuring buttons
//
//   - [INSButton.SetButtonType]: Sets the button’s type, which affects its user interface and behavior when clicked.
//   - [INSButton.GetPeriodicDelayInterval]: Returns by reference the delay and interval periods for a continuous button.
//   - [INSButton.SetPeriodicDelayInterval]: Sets the message delay and interval periods for a continuous button.
//   - [INSButton.ContentTintColor]: A tint color to use for the template image and text content.
//   - [INSButton.SetContentTintColor]
//   - [INSButton.HasDestructiveAction]: A Boolean value that defines whether a button’s action has a destructive effect.
//   - [INSButton.SetHasDestructiveAction]
//   - [INSButton.AlternateTitle]: The title that the button displays when the button is in an on state.
//   - [INSButton.SetAlternateTitle]
//   - [INSButton.AttributedTitle]: The title that the button displays in an off state, as an attributed string.
//   - [INSButton.SetAttributedTitle]
//   - [INSButton.AttributedAlternateTitle]: The title that the button displays as an attributed string when the button is in an on state.
//   - [INSButton.SetAttributedAlternateTitle]
//   - [INSButton.Title]: The title displayed on the button when it’s in an off state.
//   - [INSButton.SetTitle]
//   - [INSButton.SymbolConfiguration]: The combination of point size, weight, and scale to use when sizing and displaying symbol images.
//   - [INSButton.SetSymbolConfiguration]
//   - [INSButton.Sound]: The sound that plays when the user clicks the button.
//   - [INSButton.SetSound]
//   - [INSButton.SpringLoaded]: A Boolean value that indicates whether spring loading is enabled for the button.
//   - [INSButton.SetSpringLoaded]
//   - [INSButton.MaxAcceleratorLevel]: An integer value indicating the maximum pressure level for a button of type [NSMultiLevelAcceleratorButton](<doc://com.apple.appkit/documentation/AppKit/NSMultiLevelAcceleratorButton>).
//   - [INSButton.SetMaxAcceleratorLevel]
//   - [INSButton.TintProminence]: The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
//   - [INSButton.SetTintProminence]
//   - [INSButton.BorderShape]
//   - [INSButton.SetBorderShape]
//
// # Configuring button images
//
//   - [INSButton.Image]: The image that appears on the button when it’s in an off state, or `nil` if there is no such image.
//   - [INSButton.SetImage]
//   - [INSButton.AlternateImage]: An alternate image that appears on the button when the button is in an on state.
//   - [INSButton.SetAlternateImage]
//   - [INSButton.ImagePosition]: The position of the button’s image relative to its title.
//   - [INSButton.SetImagePosition]
//   - [INSButton.Bordered]: A Boolean value that determines whether the button has a border.
//   - [INSButton.SetBordered]
//   - [INSButton.Transparent]: A Boolean value that indicates whether the button is transparent.
//   - [INSButton.SetTransparent]
//   - [INSButton.BezelStyle]: The appearance of the button’s border.
//   - [INSButton.SetBezelStyle]
//   - [INSButton.BezelColor]: The color of the button’s bezel, in appearances that support it.
//   - [INSButton.SetBezelColor]
//   - [INSButton.ShowsBorderOnlyWhileMouseInside]: A Boolean value that determines whether the button displays its border only when the pointer is over it.
//   - [INSButton.SetShowsBorderOnlyWhileMouseInside]
//   - [INSButton.ImageHugsTitle]: A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//   - [INSButton.SetImageHugsTitle]
//   - [INSButton.ImageScaling]: The scaling mode applied to make the cell’s image fit the frame of the image view.
//   - [INSButton.SetImageScaling]
//
// # Managing button state
//
//   - [INSButton.AllowsMixedState]: A Boolean value that indicates whether the button allows a mixed state.
//   - [INSButton.SetAllowsMixedState]
//   - [INSButton.State]: The button’s state.
//   - [INSButton.SetState]
//   - [INSButton.SetNextState]: Sets the button to its next state.
//   - [INSButton.Highlight]: Highlights (or unhighlights) the button.
//
// # Accessing key equivalents
//
//   - [INSButton.KeyEquivalent]: The key-equivalent character of the button.
//   - [INSButton.SetKeyEquivalent]
//   - [INSButton.KeyEquivalentModifierMask]: The mask specifying the modifier keys for the button’s key equivalent.
//   - [INSButton.SetKeyEquivalentModifierMask]
//
// See: https://developer.apple.com/documentation/AppKit/NSButton
type INSButton interface {
	INSControl
	NSAccessibilityButton
	NSUserInterfaceCompression
	NSUserInterfaceValidations

	// Topic: Configuring buttons

	// Sets the button’s type, which affects its user interface and behavior when clicked.
	SetButtonType(type_ NSButtonType)
	// Returns by reference the delay and interval periods for a continuous button.
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	// Sets the message delay and interval periods for a continuous button.
	SetPeriodicDelayInterval(delay float32, interval float32)
	// A tint color to use for the template image and text content.
	ContentTintColor() INSColor
	SetContentTintColor(value INSColor)
	// A Boolean value that defines whether a button’s action has a destructive effect.
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	// The title that the button displays when the button is in an on state.
	AlternateTitle() string
	SetAlternateTitle(value string)
	// The title that the button displays in an off state, as an attributed string.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)
	// The title that the button displays as an attributed string when the button is in an on state.
	AttributedAlternateTitle() foundation.NSAttributedString
	SetAttributedAlternateTitle(value foundation.NSAttributedString)
	// The title displayed on the button when it’s in an off state.
	Title() string
	SetTitle(value string)
	// The combination of point size, weight, and scale to use when sizing and displaying symbol images.
	SymbolConfiguration() INSImageSymbolConfiguration
	SetSymbolConfiguration(value INSImageSymbolConfiguration)
	// The sound that plays when the user clicks the button.
	Sound() INSSound
	SetSound(value INSSound)
	// A Boolean value that indicates whether spring loading is enabled for the button.
	SpringLoaded() bool
	SetSpringLoaded(value bool)
	// An integer value indicating the maximum pressure level for a button of type [NSMultiLevelAcceleratorButton](<doc://com.apple.appkit/documentation/AppKit/NSMultiLevelAcceleratorButton>).
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	// The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
	TintProminence() NSTintProminence
	SetTintProminence(value NSTintProminence)
	BorderShape() NSControlBorderShape
	SetBorderShape(value NSControlBorderShape)

	// Topic: Configuring button images

	// The image that appears on the button when it’s in an off state, or `nil` if there is no such image.
	Image() INSImage
	SetImage(value INSImage)
	// An alternate image that appears on the button when the button is in an on state.
	AlternateImage() INSImage
	SetAlternateImage(value INSImage)
	// The position of the button’s image relative to its title.
	ImagePosition() NSCellImagePosition
	SetImagePosition(value NSCellImagePosition)
	// A Boolean value that determines whether the button has a border.
	Bordered() bool
	SetBordered(value bool)
	// A Boolean value that indicates whether the button is transparent.
	Transparent() bool
	SetTransparent(value bool)
	// The appearance of the button’s border.
	BezelStyle() NSBezelStyle
	SetBezelStyle(value NSBezelStyle)
	// The color of the button’s bezel, in appearances that support it.
	BezelColor() INSColor
	SetBezelColor(value INSColor)
	// A Boolean value that determines whether the button displays its border only when the pointer is over it.
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	// A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	// The scaling mode applied to make the cell’s image fit the frame of the image view.
	ImageScaling() NSImageScaling
	SetImageScaling(value NSImageScaling)

	// Topic: Managing button state

	// A Boolean value that indicates whether the button allows a mixed state.
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	// The button’s state.
	State() NSControlStateValue
	SetState(value NSControlStateValue)
	// Sets the button to its next state.
	SetNextState()
	// Highlights (or unhighlights) the button.
	Highlight(flag bool)

	// Topic: Accessing key equivalents

	// The key-equivalent character of the button.
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	// The mask specifying the modifier keys for the button’s key equivalent.
	KeyEquivalentModifierMask() NSEventModifierFlags
	SetKeyEquivalentModifierMask(value NSEventModifierFlags)
}

// Init initializes the instance.
func (b NSButton) Init() NSButton {
	rv := objc.Send[NSButton](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSButton) Autorelease() NSButton {
	rv := objc.Send[NSButton](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSButton creates a new NSButton instance.
func NewNSButton() NSButton {
	class := getNSButtonClass()
	rv := objc.Send[NSButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a standard checkbox with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewButtonCheckboxWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSButton {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return NSButtonFromID(rv)
}

// Creates a standard radio button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func NewButtonRadioButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSButton {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonClass().class), objc.Sel("radioButtonWithTitle:target:action:"), objc.String(title), target, action)
	return NSButtonFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewButtonWithCoder(coder foundation.INSCoder) NSButton {
	instance := getNSButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSButtonFromID(rv)
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
func NewButtonWithFrame(frameRect corefoundation.CGRect) NSButton {
	instance := getNSButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSButtonFromID(rv)
}

// Creates a standard push button with the image you specify.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// # Discussion
// 
// Set the image’s [AccessibilityDescription] property to ensure
// accessibility for this control.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func NewButtonWithImageTargetAction(image INSImage, target objectivec.IObject, action objc.SEL) NSButton {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonClass().class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return NSButtonFromID(rv)
}

// Creates a standard push button with a title and image.
//
// title: The localized title string to display on the button.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewButtonWithTitleImageTargetAction(title string, image INSImage, target objectivec.IObject, action objc.SEL) NSButton {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return NSButtonFromID(rv)
}

// Creates a standard push button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the control.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSButton {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return NSButtonFromID(rv)
}

// Sets the button’s type, which affects its user interface and behavior
// when clicked.
//
// type: A constant specifying the type of the button. The available button types
// are listed under [NSButton.ButtonType] in the [NSButtonCell] class.
// //
// [NSButton.ButtonType]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType
//
// # Discussion
// 
// This method causes the button to update to reflect the new type before the
// method finishes executing.
// 
// The types available are for the most common button types, which are also
// accessible in Interface Builder. You can configure different behavior with
// the [NSButtonCell] methods [HighlightsBy] and [ShowsStateBy].
// 
// Note that there is no `-buttonType` method. The set method sets various
// button properties that together establish the behavior of the type.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/setButtonType(_:)
func (b NSButton) SetButtonType(type_ NSButtonType) {
	objc.Send[objc.ID](b.ID, objc.Sel("setButtonType:"), type_)
}
// Returns by reference the delay and interval periods for a continuous
// button.
//
// delay: On return, the amount of time (in seconds) the button will pause before
// starting to periodically send action messages to the target object. The
// default delay is taken from a user’s default (60 seconds maximum). If the
// user hasn’t specified a default value, `delay` defaults to 0.4 seconds,
//
// interval: On return, the amount of time (in seconds) the button will pause between
// sending each action message. The default interval is taken from a user’s
// default (60 seconds maximum). If the user hasn’t specified a default
// value, `interval` defaults to 0.075 seconds.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/getPeriodicDelay(_:interval:)
func (b NSButton) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](b.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}
// Sets the message delay and interval periods for a continuous button.
//
// delay: The amount of time (in seconds) that a continuous button will pause before
// starting to periodically send action messages to the target object. The
// maximum allowed value is 60.0 seconds; if a larger value is supplied, it is
// ignored, and 60.0 seconds is used.
//
// interval: The amount of time (in seconds) the continuous button will pause between
// sending each action message. The maximum value is 60.0 seconds; if a larger
// value is supplied, it is ignored, and 60.0 seconds is used.
//
// # Discussion
// 
// The delay and interval values are used if the button is configured (by a
// [Continuous] message) to continuously send the action message to the target
// object while tracking the mouse.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b NSButton) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}
// Sets the priority compression options for this button.
//
// prioritizedOptions: An array of interface compression options.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b NSButton) CompressWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](b.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
}
// Returns the minimum size of the button by using the compression options.
//
// prioritizedOptions: An array of interface compression options.
//
// # Return Value
// 
// The size of the compressed button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/minimumSize(withPrioritizedCompressionOptions:)
func (b NSButton) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](b.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
	return corefoundation.CGSize(rv)
}
// Sets the button to its next state.
//
// # Discussion
// 
// If the button has three states, it cycles through them in this order: on,
// off, mixed, on, and so forth. If the button has two states, it toggles
// between them.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/setNextState()
func (b NSButton) SetNextState() {
	objc.Send[objc.ID](b.ID, objc.Sel("setNextState"))
}
// Highlights (or unhighlights) the button.
//
// flag: [true] to highlight the button; [false] to unhighlight the button. If the
// current state of the button matches `flag`, no action is taken.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Highlighting makes the button appear recessed, displays its alternate title
// or image, or causes the button to appear illuminated.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/highlight(_:)
func (b NSButton) Highlight(flag bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("highlight:"), flag)
}
// Simulates clicking the button.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton/accessibilityPerformPress()
func (b NSButton) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}
// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (b NSButton) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// A tint color to use for the template image and text content.
//
// # Discussion
// 
// The `contentTintColor` is only applicable to borderless buttons. This color
// is used in combination with other theme-appropriate effects.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b NSButton) ContentTintColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("contentTintColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSButton) SetContentTintColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setContentTintColor:"), value)
}
// A Boolean value that defines whether a button’s action has a destructive
// effect.
//
// # Discussion
// 
// The default value of `hasDestructiveAction` is `false`. Setting this to
// `true` allows the system to guard a destructive-action button against
// accidental presses, and can give the button a special appearance in certain
// contexts to caution against unintentional use.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/hasDestructiveAction
func (b NSButton) HasDestructiveAction() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("hasDestructiveAction"))
	return rv
}
func (b NSButton) SetHasDestructiveAction(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setHasDestructiveAction:"), value)
}
// The title that the button displays when the button is in an on state.
//
// # Discussion
// 
// This property contains the string that appears on the button when the
// button is in an on state, or the empty string if the button doesn’t
// display an alternate title. By default, a button’s alternate title is
// Button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/alternateTitle
func (b NSButton) AlternateTitle() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("alternateTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSButton) SetAlternateTitle(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setAlternateTitle:"), objc.String(value))
}
// The title that the button displays in an off state, as an attributed
// string.
//
// # Discussion
// 
// This property contains an string of class [NSAttributedString], which
// appears on the button when the button is in an off state. If the button
// doesn’t display a title, then this property contains an empty attributed
// string.
// 
// A button’s title is always displayed if the button doesn’t use its
// alternate contents for highlighting or displaying the on state. By default,
// a button’s title is Button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b NSButton) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (b NSButton) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](b.ID, objc.Sel("setAttributedTitle:"), value)
}
// The title that the button displays as an attributed string when the button
// is in an on state.
//
// # Discussion
// 
// This property contains the string that appears on the button when it’s in
// an on state, as an [NSAttributedString], or the empty string if the button
// doesn’t display an alternate title. By default, a button’s alternate
// title is Button.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/attributedAlternateTitle
func (b NSButton) AttributedAlternateTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attributedAlternateTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (b NSButton) SetAttributedAlternateTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](b.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}
// The title displayed on the button when it’s in an off state.
//
// # Discussion
// 
// This property contains the title displayed on the button when it’s in an
// off state or the empty string if the button doesn’t display a title. This
// title is always displayed if the button doesn’t use its alternate
// contents for highlighting or displaying the on state. By default, a
// button’s title is Button. Setting the value of this property redraws the
// button’s contents, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/title
func (b NSButton) Title() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSButton) SetTitle(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The combination of point size, weight, and scale to use when sizing and
// displaying symbol images.
//
// # Discussion
// 
// If you don’t provide a symbol configuration, the button uses its [Font]
// property to size and display the symbol image.
// 
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/symbolConfiguration
func (b NSButton) SymbolConfiguration() INSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("symbolConfiguration"))
	return NSImageSymbolConfigurationFromID(objc.ID(rv))
}
func (b NSButton) SetSymbolConfiguration(value INSImageSymbolConfiguration) {
	objc.Send[struct{}](b.ID, objc.Sel("setSymbolConfiguration:"), value)
}
// The sound that plays when the user clicks the button.
//
// # Discussion
// 
// The sound represented by this property is played during a mouse event, such
// as [NSLeftMouseDown].
//
// [NSLeftMouseDown]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDown
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b NSButton) Sound() INSSound {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sound"))
	return NSSoundFromID(objc.ID(rv))
}
func (b NSButton) SetSound(value INSSound) {
	objc.Send[struct{}](b.ID, objc.Sel("setSound:"), value)
}
// A Boolean value that indicates whether spring loading is enabled for the
// button.
//
// # Discussion
// 
// The value of this property is [true] if spring loading is enabled for the
// button, and [false] if it is not. The default is [false].
// 
// On pressure-sensitive systems, such as systems with the Force Touch
// trackpad, spring loading is a feature that allows a user to activate a
// button by dragging selected items over it and force clicking—pressing
// harder—without dropping the selected items. The user can then continue
// dragging the items, possibly to perform additional actions.
// 
// A practical example of this feature can be found in the Calendar app. A
// selected calendar event can be dragged over the Calendars button in the
// toolbar. Force clicking on the button displays the calendar list without
// releasing the selected event. The event can then be dropped onto a calendar
// in the list, which assigns it to that calendar.
// 
// If spring loading is enabled on a button and a user drags items over it,
// the button highlights to indicate that it responds to force clicking. If
// the user presses harder, additional highlighting occurs to indicate that
// the button was fully activated.
// 
// On systems that don’t support pressure sensitivity, simply hovering over
// the button for a short period of time is sufficient to activate the button.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/isSpringLoaded
func (b NSButton) SpringLoaded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isSpringLoaded"))
	return rv
}
func (b NSButton) SetSpringLoaded(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setSpringLoaded:"), value)
}
// An integer value indicating the maximum pressure level for a button of type
// [NSMultiLevelAcceleratorButton].
//
// [NSMultiLevelAcceleratorButton]: https://developer.apple.com/documentation/AppKit/NSMultiLevelAcceleratorButton
//
// # Discussion
// 
// A multilevel accelerator button is a variation of a standard accelerator
// button that allows for a configurable number of stepped pressure levels in
// a system that supports pressure-sensitivity, such as the Force Touch
// trackpad. As each level is reached, the user receives light tactile
// feedback, and an action is sent.
// 
// You configure the number of pressure levels for a multilevel accelerator
// button by adjusting the value of [MaxAcceleratorLevel]. For other types of
// buttons, this property value defaults to `1`. For multilevel accelerator
// buttons, this property value defaults to `2`, and may be set to a value
// between `1` and `5`.
// 
// The behavior of a multilevel accelerator button is reliant on a system that
// supports pressure sensitivity. On a system that doesn’t support pressure
// sensitivity, a multilevel accelerator button always has a value of `1` when
// the user clicks it.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/maxAcceleratorLevel
func (b NSButton) MaxAcceleratorLevel() int {
	rv := objc.Send[int](b.ID, objc.Sel("maxAcceleratorLevel"))
	return rv
}
func (b NSButton) SetMaxAcceleratorLevel(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setMaxAcceleratorLevel:"), value)
}
// The tint prominence of the button. Use tint prominence to gently suggest a
// hierarchy when multiple buttons perform similar actions. A button with
// primary tint prominence suggests the most preferred option, while secondary
// prominence indicates a reasonable alternative. See [NSTintProminence] for a
// list of possible values.
//
// [NSTintProminence]: https://developer.apple.com/documentation/AppKit/NSTintProminence
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/tintProminence
func (b NSButton) TintProminence() NSTintProminence {
	rv := objc.Send[NSTintProminence](b.ID, objc.Sel("tintProminence"))
	return NSTintProminence(rv)
}
func (b NSButton) SetTintProminence(value NSTintProminence) {
	objc.Send[struct{}](b.ID, objc.Sel("setTintProminence:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b NSButton) BorderShape() NSControlBorderShape {
	rv := objc.Send[NSControlBorderShape](b.ID, objc.Sel("borderShape"))
	return NSControlBorderShape(rv)
}
func (b NSButton) SetBorderShape(value NSControlBorderShape) {
	objc.Send[struct{}](b.ID, objc.Sel("setBorderShape:"), value)
}
// The image that appears on the button when it’s in an off state, or `nil`
// if there is no such image.
//
// # Discussion
// 
// The image contained in this property is always displayed on a button that
// doesn’t change its contents when highlighting or showing an on state.
// Buttons don’t display images by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/image
func (b NSButton) Image() INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (b NSButton) SetImage(value INSImage) {
	objc.Send[struct{}](b.ID, objc.Sel("setImage:"), value)
}
// An alternate image that appears on the button when the button is in an on
// state.
//
// # Discussion
// 
// The value of this property is `nil` if there is no alternate image for the
// button. Note that some button types don’t display an alternate image, and
// buttons don’t display images by default. If you use this property to set
// an image, the button redraws its contents if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/alternateImage
func (b NSButton) AlternateImage() INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("alternateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (b NSButton) SetAlternateImage(value INSImage) {
	objc.Send[struct{}](b.ID, objc.Sel("setAlternateImage:"), value)
}
// The position of the button’s image relative to its title.
//
// # Discussion
// 
// If the title is above, below, or overlapping the image, or if there is no
// image, the text is horizontally centered within the button. See
// [NSControl.ImagePosition] in [NSCell] for the list of possible image
// positions.
//
// [NSControl.ImagePosition]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/imagePosition
func (b NSButton) ImagePosition() NSCellImagePosition {
	rv := objc.Send[NSCellImagePosition](b.ID, objc.Sel("imagePosition"))
	return NSCellImagePosition(rv)
}
func (b NSButton) SetImagePosition(value NSCellImagePosition) {
	objc.Send[struct{}](b.ID, objc.Sel("setImagePosition:"), value)
}
// A Boolean value that determines whether the button has a border.
//
// # Discussion
// 
// The value of this property is [true] if the button has a border, [false]
// otherwise. A button’s border isn’t the single line of most other
// controls’ borders—instead, it’s a raised bezel. By default, buttons
// are bordered. If the bordered state of a button changes, it gets redrawn.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/isBordered
func (b NSButton) Bordered() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isBordered"))
	return rv
}
func (b NSButton) SetBordered(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setBordered:"), value)
}
// A Boolean value that indicates whether the button is transparent.
//
// # Discussion
// 
// The value of this property is [true] if the button is transparent, [false]
// otherwise. A transparent button never draws itself, but it receives mouse
// events, sends its action, and tracks the mouse properly. A transparent
// button can be useful for sensitizing an area on the screen so that an
// action gets sent to a target when the area receives a mouse click. Setting
// this property causes the button to redraw if necessary.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/isTransparent
func (b NSButton) Transparent() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isTransparent"))
	return rv
}
func (b NSButton) SetTransparent(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setTransparent:"), value)
}
// The appearance of the button’s border.
//
// # Return Value
// 
// The bezel style of the button. See [NSButton.BezelStyle] in [NSButtonCell]
// for the list of possible values.
// 
// # Discussion
// 
// Note that if the button is not bordered, the bezel style is ignored.
//
// [NSButton.BezelStyle]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/bezelStyle-swift.property
func (b NSButton) BezelStyle() NSBezelStyle {
	rv := objc.Send[NSBezelStyle](b.ID, objc.Sel("bezelStyle"))
	return NSBezelStyle(rv)
}
func (b NSButton) SetBezelStyle(value NSBezelStyle) {
	objc.Send[struct{}](b.ID, objc.Sel("setBezelStyle:"), value)
}
// The color of the button’s bezel, in appearances that support it.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b NSButton) BezelColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bezelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSButton) SetBezelColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setBezelColor:"), value)
}
// A Boolean value that determines whether the button displays its border only
// when the pointer is over it.
//
// # Discussion
// 
// The value of this property is [true] if the button’s border is displayed
// only when the pointer is over the button and the button is active. The
// value is [false] if the border is displayed all the time, regardless of the
// position of the pointer. By default, this method returns [false].
// 
// If [Bordered] is [false], the border is never displayed, regardless of the
// value of [ShowsBorderOnlyWhileMouseInside].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/showsBorderOnlyWhileMouseInside
func (b NSButton) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}
func (b NSButton) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}
// A Boolean value that determines how the button’s image and title are
// positioned together within the button bezel.
//
// # Discussion
// 
// When the value of this property is [false] (the default value), the
// button’s image is positioned according to the [ImagePosition] property at
// the edge of the button bezel, and the title is positioned within the
// remaining space.
// 
// When this property is [true], the button’s image is positioned directly
// adjacent to the title based on the [ImagePosition] property, and the image
// and title are positioned within the button bezel as a single unit.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/imageHugsTitle
func (b NSButton) ImageHugsTitle() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("imageHugsTitle"))
	return rv
}
func (b NSButton) SetImageHugsTitle(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setImageHugsTitle:"), value)
}
// The scaling mode applied to make the cell’s image fit the frame of the
// image view.
//
// # Discussion
// 
// The default value of this property is
// [NSImageScaling.scaleProportionallyDown].
//
// [NSImageScaling.scaleProportionallyDown]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/imageScaling
func (b NSButton) ImageScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](b.ID, objc.Sel("imageScaling"))
	return NSImageScaling(rv)
}
func (b NSButton) SetImageScaling(value NSImageScaling) {
	objc.Send[struct{}](b.ID, objc.Sel("setImageScaling:"), value)
}
// The compression options active for this button.
//
// # Discussion
// 
// Only compression options that have been applied and are actively being
// respected are returned. For more information about managing button sizes
// when space is restriced, see [NSUserInterfaceCompressionOptions].
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/activeCompressionOptions
func (b NSButton) ActiveCompressionOptions() INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("activeCompressionOptions"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the button allows a mixed state.
//
// # Discussion
// 
// The value of this property is [true] if the button has three states (on,
// off, and mixed), or [false] if the button has two states (on and off). The
// default value is [false]. On and off states (also referred to as alternate
// and normal) indicate that the button is either clicked or not clicked.
// Mixed state is typically used for checkboxes or radio buttons. For example,
// suppose the state of a checkbox is used to denote whether a text field
// contains bold text. If all of the text in the text field is bold, then the
// checkbox appears checked (on). If none of the text is bold, then the
// checkbox appears unchecked (off). If some of the text is bold, then the
// checkbox contains a dash (mixed).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/allowsMixedState
func (b NSButton) AllowsMixedState() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsMixedState"))
	return rv
}
func (b NSButton) SetAllowsMixedState(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsMixedState:"), value)
}
// The button’s state.
//
// # Discussion
// 
// The value of this property represents the button’s state. A button can
// have two or three states. If it has two, this value is either on
// ([NSOnState]) or off ([NSOffState]). If it has three, this value is on,
// off, or mixed ([NSMixedState]). A three-state button can be enabled by
// calling the [AllowsMixedState] method. On and off states (also referred to
// as alternate and normal) indicate that the button is either clicked or not
// clicked. Mixed state is typically used for checkboxes or radio buttons,
// which allow for an additional intermediate state. For example, suppose the
// state of a checkbox is used to denote whether a text field contains bold
// text. If all of the text in the text field is bold, then the checkbox
// appears checked (on). If none of the text is bold, then the checkbox
// appears unchecked (off). If some of the text is bold, then the checkbox
// contains a dash (mixed).
// 
// Note that if the button has only two states and you set the value of
// [State] to mixed, the button’s state changes to on. Setting this property
// redraws the button, if necessary.
// 
// Although using the enumerated constants is preferred, you can also set
// [State] to an integer value. If the button has two states, `0` is treated
// as [NSOffState], and a nonzero value is treated as [NSOnState]. If the
// button has three states, `0` is treated as [NSOffState]; a negative value,
// as [NSMixedState]; and a positive value, as [NSOnState].
// 
// To check whether the button uses the mixed state, use the
// [AllowsMixedState] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/state
func (b NSButton) State() NSControlStateValue {
	rv := objc.Send[NSControlStateValue](b.ID, objc.Sel("state"))
	return NSControlStateValue(rv)
}
func (b NSButton) SetState(value NSControlStateValue) {
	objc.Send[struct{}](b.ID, objc.Sel("setState:"), value)
}
// The key-equivalent character of the button.
//
// # Discussion
// 
// This property contains the button’s key equivalent, or the empty string
// if no equivalent has been defined. Buttons don’t have a default key
// equivalent.
// 
// If you set a key equivalent instead of an image, the button’s interior is
// redrawn. However, the key equivalent isn’t displayed if the image
// position is set to [NSNoImage], [NSImageOnly], or [NSImageOverlaps]; that
// is, the button must display both its title and its “image” (which is
// the key equivalent in this case), and they must not overlap.
// 
// To display a key equivalent on a button, set the image and alternate image
// to `nil`, set the key equivalent, and then set the image position.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalent
func (b NSButton) KeyEquivalent() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("keyEquivalent"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSButton) SetKeyEquivalent(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}
// The mask specifying the modifier keys for the button’s key equivalent.
//
// # Discussion
// 
// This property contains the mask specifying the modifier keys that are
// applied to the button’s key equivalent. Mask bits are defined in Modifier
// Flags. The only mask bits relevant in button key-equivalent modifier masks
// are [NSControlKeyMask], [NSAlternateKeyMask], and [NSCommandKeyMask].
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalentModifierMask
func (b NSButton) KeyEquivalentModifierMask() NSEventModifierFlags {
	rv := objc.Send[NSEventModifierFlags](b.ID, objc.Sel("keyEquivalentModifierMask"))
	return NSEventModifierFlags(rv)
}
func (b NSButton) SetKeyEquivalentModifierMask(value NSEventModifierFlags) {
	objc.Send[struct{}](b.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

			// Protocol methods for NSAccessibilityButton
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSButton) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}
// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSButton) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}
// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSButton) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSButton) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSUserInterfaceCompression
			

			// Protocol methods for NSUserInterfaceValidations
			

