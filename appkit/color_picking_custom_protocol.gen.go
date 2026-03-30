// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that provides a way to add color pickers—custom user interfaces for color selection—to an app’s color panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom
type NSColorPickingCustom interface {
	objectivec.IObject
	NSColorPickingDefault

	// Adjusts the receiver to make the specified color the currently selected color.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/setColor(_:)
	SetColor(newColor INSColor)

	// Returns the receiver’s current mode (or submode, if applicable).
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/currentMode()
	CurrentMode() NSColorPanelMode

	// Returns a Boolean value indicating whether or not the receiver supports the specified picking mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/supportsMode(_:)
	SupportsMode(mode NSColorPanelMode) bool

	// Returns the view containing the receiver’s user interface.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/provideNewView(_:)
	ProvideNewView(initialRequest bool) INSView
}

// NSColorPickingCustomObject wraps an existing Objective-C object that conforms to the NSColorPickingCustom protocol.
type NSColorPickingCustomObject struct {
	objectivec.Object
}

func (o NSColorPickingCustomObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSColorPickingCustomObjectFromID constructs a [NSColorPickingCustomObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSColorPickingCustomObjectFromID(id objc.ID) NSColorPickingCustomObject {
	return NSColorPickingCustomObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Adjusts the receiver to make the specified color the currently selected
// color.
//
// newColor: The color to set as the currently selected color.
//
// # Discussion
//
// This method is invoked on the current color picker each time
// [NSColorPanel]‘s [Color] method is invoked. If `color` is actually
// different from the color picker’s color (as it would be if, for example,
// the user dragged a color into [NSColorPanel]‘s color well), this method
// could be used to update the color picker’s color to reflect the change.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/setColor(_:)
func (o NSColorPickingCustomObject) SetColor(newColor INSColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), newColor)
}

// Returns the receiver’s current mode (or submode, if applicable).
//
// # Return Value
//
// The current color picker mode. The returned value should be unique to your
// color picker. See this protocol description’s list of the unique values
// for the standard color pickers used by the Application Kit.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/currentMode()
func (o NSColorPickingCustomObject) CurrentMode() NSColorPanelMode {
	rv := objc.Send[NSColorPanelMode](o.ID, objc.Sel("currentMode"))
	return rv
}

// Returns a Boolean value indicating whether or not the receiver supports the
// specified picking mode.
//
// mode: The color picking mode.
//
// # Return Value
//
// true if the color picker supports the specified color picking mode;
// otherwise false.
//
// # Discussion
//
// This method is invoked when the [NSColorPanel] is first initialized: It is
// used to attempt to restore the user’s previously selected mode. It is
// also invoked by [NSColorPanel]‘s [Mode] method to find the color picker
// that supports a particular mode. See this protocol description’s list of
// the unique mode values for the standard color pickers used by the
// Application Kit.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/supportsMode(_:)
func (o NSColorPickingCustomObject) SupportsMode(mode NSColorPanelMode) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsMode:"), mode)
	return rv
}

// Returns the view containing the receiver’s user interface.
//
// initialRequest: true only when this method is first invoked for your color picker. If
// `initialRequest` is true, the method should perform any initialization
// required (such as lazily loading a nib file, initializing the view, or
// performing any other custom initialization required for your picker).
//
// # Return Value
//
// The view containing the color picker’s user interface. The [NSView]
// returned by this method should be set to automatically resize both its
// width and height.
//
// # Discussion
//
// This message is sent to the color picker whenever the color panel attempts
// to display it. This may be when the panel is first presented, when the user
// switches pickers, or when the picker is switched through an API.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingCustom/provideNewView(_:)
func (o NSColorPickingCustomObject) ProvideNewView(initialRequest bool) INSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("provideNewView:"), initialRequest)
	return NSViewFromID(rv)
}

// Specifies the receiver’s mode.
//
// mode: The color picker mode. The available modes are described in [Choosing the
// Color Pickers in a Color Panel].
//
// # Discussion
//
// This method is invoked by the [NSColorPanel] method [Mode] method to ensure
// the color picker reflects the current mode. For example, invoke this method
// during color picker initialization to ensure that all color pickers are
// restored to the mode the user left them in the last time an [NSColorPanel]
// was used.
//
// Most color pickers have only one mode and thus don’t need to do any work
// in this method. An example of a color picker that uses this method is the
// slider picker, which can choose from one of several submodes depending on
// the value of `mode`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/setMode(_:)
//
// [Choosing the Color Pickers in a Color Panel]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DrawColor/Tasks/ChoosingColorPickers.html#//apple_ref/doc/uid/20000792
func (o NSColorPickingCustomObject) SetMode(mode NSColorPanelMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setMode:"), mode)
}

// Sets the image of a given button cell.
//
// newButtonImage: The image to set for the button cell.
//
// buttonCell: The [NSButtonCell] object that lets the user choose the picker from the
// color panel—the color picker’s representation in the [NSMatrix] of the
// [NSColorPanel].
//
// # Discussion
//
// This method should perform application-specific manipulation of the image
// before it’s inserted and displayed by the button cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/insertNewButtonImage(_:in:)
func (o NSColorPickingCustomObject) InsertNewButtonImageIn(newButtonImage INSImage, buttonCell INSButtonCell) {
	objc.Send[struct{}](o.ID, objc.Sel("insertNewButtonImage:in:"), newButtonImage, buttonCell)
}

// Provides the image of the button used to select the receiver in the color
// panel.
//
// # Return Value
//
// The image for the mode button the user uses to select this picker in the
// color panel; that is, the color picker’s representation in the [NSMatrix]
// of the [NSColorPanel].
//
// # Discussion
//
// This image is the same one the color panel uses as an argument when sending
// the [InsertNewButtonImageIn] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/provideNewButtonImage()
func (o NSColorPickingCustomObject) ProvideNewButtonImage() INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("provideNewButtonImage"))
	return NSImageFromID(rv)
}

// Indicates the receiver’s minimum content size.
//
// # Discussion
//
// The receiver does not allow a size smaller than `minContentSize`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/minContentSize()
func (o NSColorPickingCustomObject) MinContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("minContentSize"))
	return rv
}

// Provides the toolbar button help tag.
//
// # Return Value
//
// Help tag text.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/buttonToolTip()
func (o NSColorPickingCustomObject) ButtonToolTip() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttonToolTip"))
	return foundation.NSStringFromID(rv).String()
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
// them, depending on whether the sender’s [ShowsAlpha] method returns false
// or true.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/alphaControlAddedOrRemoved(_:)
func (o NSColorPickingCustomObject) AlphaControlAddedOrRemoved(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("alphaControlAddedOrRemoved:"), sender)
}

// Tells the recever when the color panel’s view size changes in a way that
// might affect the color picker.
//
// sender: The [NSColorPanel] that contains the color picker.
//
// # Discussion
//
// Use this method to perform special preparation when resizing the color
// picker’s view. Because this method is invoked only as appropriate, it’s
// better to implement this method than to override the method “ for the
// [NSView] in which the color picker’s user interface is contained.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/viewSizeChanged(_:)
func (o NSColorPickingCustomObject) ViewSizeChanged(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("viewSizeChanged:"), sender)
}

// Tells the receiver to attach the given color list, if it isn’t already
// displaying the list.
//
// colorList: The color list to display.
//
// # Discussion
//
// You never invoke this method; it’s invoked automatically by the
// [NSColorPanel] object when its [AttachColorList] method is invoked. Because
// the [NSColorPanel] list mode manages [NSColorList] objects, this method
// need only be implemented by a custom color picker that manages
// [NSColorList] objects itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/attachColorList(_:)
func (o NSColorPickingCustomObject) AttachColorList(colorList INSColorList) {
	objc.Send[struct{}](o.ID, objc.Sel("attachColorList:"), colorList)
}

// Tells the receiver to detach the given color list, unless the receiver
// isn’t displaying the list.
//
// colorList: The color list to detach.
//
// # Discussion
//
// You never invoke this method; it’s invoked automatically by the
// [NSColorPanel] object when its [DetachColorList] method is invoked. Because
// the [NSColorPanel] list mode manages [NSColorList] objects, this method
// need only be implemented by a custom color picker that manages
// [NSColorList] objects itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/detachColorList(_:)
func (o NSColorPickingCustomObject) DetachColorList(colorList INSColorList) {
	objc.Send[struct{}](o.ID, objc.Sel("detachColorList:"), colorList)
}
