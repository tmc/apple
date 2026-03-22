// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that provides basic behavior for a color picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault
type NSColorPickingDefault interface {
	objectivec.IObject

	// Specifies the receiver’s mode.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/setMode(_:)
	SetMode(mode NSColorPanelMode)

	// Sets the image of a given button cell.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/insertNewButtonImage(_:in:)
	InsertNewButtonImageIn(newButtonImage INSImage, buttonCell INSButtonCell)

	// Provides the image of the button used to select the receiver in the color panel.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/provideNewButtonImage()
	ProvideNewButtonImage() INSImage

	// Indicates the receiver’s minimum content size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/minContentSize()
	MinContentSize() corefoundation.CGSize

	// Provides the toolbar button help tag.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/buttonToolTip()
	ButtonToolTip() string

	// Tells the receiver to attach the given color list, if it isn’t already displaying the list.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/attachColorList(_:)
	AttachColorList(colorList INSColorList)

	// Tells the receiver to detach the given color list, unless the receiver isn’t displaying the list.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/detachColorList(_:)
	DetachColorList(colorList INSColorList)
}

// NSColorPickingDefaultObject wraps an existing Objective-C object that conforms to the NSColorPickingDefault protocol.
type NSColorPickingDefaultObject struct {
	objectivec.Object
}
func (o NSColorPickingDefaultObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSColorPickingDefaultObjectFromID constructs a [NSColorPickingDefaultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSColorPickingDefaultObjectFromID(id objc.ID) NSColorPickingDefaultObject {
	return NSColorPickingDefaultObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Specifies the receiver’s mode.
//
// mode: The color picker mode. The available modes are described in [Choosing the
// Color Pickers in a Color Panel].
// //
// [Choosing the Color Pickers in a Color Panel]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DrawColor/Tasks/ChoosingColorPickers.html#//apple_ref/doc/uid/20000792
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
func (o NSColorPickingDefaultObject) SetMode(mode NSColorPanelMode) {
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
func (o NSColorPickingDefaultObject) InsertNewButtonImageIn(newButtonImage INSImage, buttonCell INSButtonCell) {
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
func (o NSColorPickingDefaultObject) ProvideNewButtonImage() INSImage {
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
func (o NSColorPickingDefaultObject) MinContentSize() corefoundation.CGSize {
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
func (o NSColorPickingDefaultObject) ButtonToolTip() string {
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
// them, depending on whether the sender’s [ShowsAlpha] method returns
// [false] or [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/alphaControlAddedOrRemoved(_:)
func (o NSColorPickingDefaultObject) AlphaControlAddedOrRemoved(sender objectivec.IObject) {
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
// better to implement this method than to override the method `` for the
// [NSView] in which the color picker’s user interface is contained.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickingDefault/viewSizeChanged(_:)
func (o NSColorPickingDefaultObject) ViewSizeChanged(sender objectivec.IObject) {
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
func (o NSColorPickingDefaultObject) AttachColorList(colorList INSColorList) {
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
func (o NSColorPickingDefaultObject) DetachColorList(colorList INSColorList) {
	objc.Send[struct{}](o.ID, objc.Sel("detachColorList:"), colorList)
	}

