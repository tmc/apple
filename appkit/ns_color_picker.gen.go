// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorPicker] class.
var (
	_NSColorPickerClass     NSColorPickerClass
	_NSColorPickerClassOnce sync.Once
)

func getNSColorPickerClass() NSColorPickerClass {
	_NSColorPickerClassOnce.Do(func() {
		_NSColorPickerClass = NSColorPickerClass{class: objc.GetClass("NSColorPicker")}
	})
	return _NSColorPickerClass
}

// GetNSColorPickerClass returns the class object for NSColorPicker.
func GetNSColorPickerClass() NSColorPickerClass {
	return getNSColorPickerClass()
}

type NSColorPickerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorPickerClass) Alloc() NSColorPicker {
	rv := objc.Send[NSColorPicker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass that implements the default color picking protocol.
//
// # Overview
// 
// The [NSColorPickingDefault] and [NSColorPickingCustom] protocols define a
// way to add color pickers (custom user interfaces for color selection) to
// the color panel.
//
// # Initializing the Color Picker Object
//
//   - [NSColorPicker.InitWithPickerMaskColorPanel]: Initializes the color picker with the specified color panel and color picker mode mask.
//
// # Getting the Color Panel
//
//   - [NSColorPicker.ColorPanel]: The color panel instance that owns the color picker.
//
// # Adding Button Images
//
//   - [NSColorPicker.InsertNewButtonImageIn]: Sets the image used for the specified button cell.
//
// # Setting the Mode
//
//   - [NSColorPicker.SetMode]: Overriden to set the color picker’s mode.
//
// # Managing Color Lists
//
//   - [NSColorPicker.AttachColorList]: Overriden to attach a color list to a color picker.
//   - [NSColorPicker.DetachColorList]: Overriden to detach a color list from a color picker.
//
// # Responding to View Changes
//
//   - [NSColorPicker.ViewSizeChanged]: Overriden to respond to a size change.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker
type NSColorPicker struct {
	objectivec.Object
}

// NSColorPickerFromID constructs a [NSColorPicker] from an objc.ID.
//
// An abstract superclass that implements the default color picking protocol.
func NSColorPickerFromID(id objc.ID) NSColorPicker {
	return NSColorPicker{objectivec.Object{ID: id}}
}
// NOTE: NSColorPicker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorPicker] class.
//
// # Initializing the Color Picker Object
//
//   - [INSColorPicker.InitWithPickerMaskColorPanel]: Initializes the color picker with the specified color panel and color picker mode mask.
//
// # Getting the Color Panel
//
//   - [INSColorPicker.ColorPanel]: The color panel instance that owns the color picker.
//
// # Adding Button Images
//
//   - [INSColorPicker.InsertNewButtonImageIn]: Sets the image used for the specified button cell.
//
// # Setting the Mode
//
//   - [INSColorPicker.SetMode]: Overriden to set the color picker’s mode.
//
// # Managing Color Lists
//
//   - [INSColorPicker.AttachColorList]: Overriden to attach a color list to a color picker.
//   - [INSColorPicker.DetachColorList]: Overriden to detach a color list from a color picker.
//
// # Responding to View Changes
//
//   - [INSColorPicker.ViewSizeChanged]: Overriden to respond to a size change.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker
type INSColorPicker interface {
	objectivec.IObject

	// Topic: Initializing the Color Picker Object

	// Initializes the color picker with the specified color panel and color picker mode mask.
	InitWithPickerMaskColorPanel(mask uint, owningColorPanel INSColorPanel) NSColorPicker

	// Topic: Getting the Color Panel

	// The color panel instance that owns the color picker.
	ColorPanel() INSColorPanel

	// Topic: Adding Button Images

	// Sets the image used for the specified button cell.
	InsertNewButtonImageIn(newButtonImage INSImage, buttonCell INSButtonCell)

	// Topic: Setting the Mode

	// Overriden to set the color picker’s mode.
	SetMode(mode NSColorPanelMode)

	// Topic: Managing Color Lists

	// Overriden to attach a color list to a color picker.
	AttachColorList(colorList INSColorList)
	// Overriden to detach a color list from a color picker.
	DetachColorList(colorList INSColorList)

	// Topic: Responding to View Changes

	// Overriden to respond to a size change.
	ViewSizeChanged(sender objectivec.IObject)

	// Sent when the color panel’s opacity controls have been hidden or displayed.
	AlphaControlAddedOrRemoved(sender objectivec.IObject)
}

// Init initializes the instance.
func (c NSColorPicker) Init() NSColorPicker {
	rv := objc.Send[NSColorPicker](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorPicker) Autorelease() NSColorPicker {
	rv := objc.Send[NSColorPicker](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorPicker creates a new NSColorPicker instance.
func NewNSColorPicker() NSColorPicker {
	class := getNSColorPickerClass()
	rv := objc.Send[NSColorPicker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the color picker with the specified color panel and color
// picker mode mask.
//
// mask: The color picker mask.
//
// owningColorPanel: The [NSColorPanel] that owns the color picker. This value is cached so it
// can be accessed using the [ColorPanel] property.
//
// # Return Value
// 
// An initialized color picker object.
//
// # Discussion
// 
// Override this method to respond to the values in `mask` or do other custom
// initialization. If you override this method in a subclass, you should
// forward the message to `super` as part of the implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/init(pickerMask:colorPanel:)
func NewColorPickerWithPickerMaskColorPanel(mask uint, owningColorPanel INSColorPanel) NSColorPicker {
	instance := getNSColorPickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPickerMask:colorPanel:"), mask, owningColorPanel)
	return NSColorPickerFromID(rv)
}

// Initializes the color picker with the specified color panel and color
// picker mode mask.
//
// mask: The color picker mask.
//
// owningColorPanel: The [NSColorPanel] that owns the color picker. This value is cached so it
// can be accessed using the [ColorPanel] property.
//
// # Return Value
// 
// An initialized color picker object.
//
// # Discussion
// 
// Override this method to respond to the values in `mask` or do other custom
// initialization. If you override this method in a subclass, you should
// forward the message to `super` as part of the implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/init(pickerMask:colorPanel:)
func (c NSColorPicker) InitWithPickerMaskColorPanel(mask uint, owningColorPanel INSColorPanel) NSColorPicker {
	rv := objc.Send[NSColorPicker](c.ID, objc.Sel("initWithPickerMask:colorPanel:"), mask, owningColorPanel)
	return rv
}
// Sets the image used for the specified button cell.
//
// newButtonImage: The image used for the specified button cell.
//
// buttonCell: The button cell for which to set the image.
//
// # Discussion
// 
// Called by the color panel to insert a new image into the specified cell by
// invoking [NSButtonCell]’s setImage: method. Override this method to
// customize `newButtonImage` before insertion in `buttonCell`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/insertNewButtonImage(_:in:)
func (c NSColorPicker) InsertNewButtonImageIn(newButtonImage INSImage, buttonCell INSButtonCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertNewButtonImage:in:"), newButtonImage, buttonCell)
}
// Overriden to set the color picker’s mode.
//
// mode: A constant specifying the color picking mode. These constants are defined
// in `AppKit/NSColorPanel.H()`.
//
// # Discussion
// 
// In grayscale-alpha, red-green-blue, cyan-magenta-yellow-black, and
// hue-saturation-brightness modes, the user adjusts colors by manipulating
// sliders. In the custom palette mode, the user can load an [NSImage] file
// (TIFF or EPS) into the [NSColorPanel], then select colors from the image.
// In custom color list mode, the user can create and load lists of named
// colors. The two custom modes provide [NSPopUpList]s for loading and saving
// files. Finally, color wheel mode provides a simplified control for
// selecting colors.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/setMode(_:)
func (c NSColorPicker) SetMode(mode NSColorPanelMode) {
	objc.Send[objc.ID](c.ID, objc.Sel("setMode:"), mode)
}
// Overriden to attach a color list to a color picker.
//
// colorList: The color list to attach to the color picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/attachColorList(_:)
func (c NSColorPicker) AttachColorList(colorList INSColorList) {
	objc.Send[objc.ID](c.ID, objc.Sel("attachColorList:"), colorList)
}
// Overriden to detach a color list from a color picker.
//
// colorList: The color list to detach.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/detachColorList(_:)
func (c NSColorPicker) DetachColorList(colorList INSColorList) {
	objc.Send[objc.ID](c.ID, objc.Sel("detachColorList:"), colorList)
}
// Overriden to respond to a size change.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/viewSizeChanged(_:)
func (c NSColorPicker) ViewSizeChanged(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("viewSizeChanged:"), sender)
}
// Sent when the color panel’s opacity controls have been hidden or
// displayed.
//
// sender: The color panel sending the message.
//
// # Discussion
// 
// This method is invoked automatically when the opacity slider of the
// [NSColorPanel] is added or removed; you never invoke this method directly.
// 
// If the color picker has its own opacity controls, it should hide or display
// them, depending on whether the sender’s [ShowsAlpha] method returns
// [false] or [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/alphaControlAddedOrRemoved(_:)
func (c NSColorPicker) AlphaControlAddedOrRemoved(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("alphaControlAddedOrRemoved:"), sender)
}

// The color panel instance that owns the color picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/colorPanel
func (c NSColorPicker) ColorPanel() INSColorPanel {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorPanel"))
	return NSColorPanelFromID(objc.ID(rv))
}
// The button image used by the color picker.
//
// # Discussion
// 
// The image placed on the mode button the user uses to select this color
// picker. This is the same image the color panel uses as an argument when
// sending the [InsertNewButtonImageIn] message. Override this property’s
// getter method to provide a custom button image. The default implementation
// looks in the color picker’s bundle for a TIFF file named after the color
// picker’s class, with the extension “`XCUIElementTypeTiff`”.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/provideNewButtonImage
func (c NSColorPicker) ProvideNewButtonImage() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("provideNewButtonImage"))
	return NSImageFromID(objc.ID(rv))
}
// The tool tip that is shown when the mouse cursor is over the color
// picker’s button image.
//
// # Discussion
// 
// Override this property’s getter method to provide a custom tool tip. The
// default implementation returns the name of the receiver’s class. If you
// want the color picker to have no tool tip, return an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/buttonToolTip
func (c NSColorPicker) ButtonToolTip() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("buttonToolTip"))
	return foundation.NSStringFromID(rv).String()
}
// The minimum content size.
//
// # Discussion
// 
// The containing [NSColorPanel] object does not allow the color picker to be
// made smaller than this size.
// 
// Override this property’s getter method to return a minimum size for the
// color picker’s content area. The default implementation obtains the
// minimum content size from the view-autoresizing behavior specified for the
// color picker. You should not have to override this method if you properly
// set up the color picker’s auto-sizing attributes in Interface Builder.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPicker/minContentSize
func (c NSColorPicker) MinContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("minContentSize"))
	return corefoundation.CGSize(rv)
}

