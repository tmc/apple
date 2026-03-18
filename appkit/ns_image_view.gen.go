// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSImageView] class.
var (
	_NSImageViewClass     NSImageViewClass
	_NSImageViewClassOnce sync.Once
)

func getNSImageViewClass() NSImageViewClass {
	_NSImageViewClassOnce.Do(func() {
		_NSImageViewClass = NSImageViewClass{class: objc.GetClass("NSImageView")}
	})
	return _NSImageViewClass
}

// GetNSImageViewClass returns the class object for NSImageView.
func GetNSImageViewClass() NSImageViewClass {
	return getNSImageViewClass()
}

type NSImageViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSImageViewClass) Alloc() NSImageView {
	rv := objc.Send[NSImageView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A display of image data in a frame.
//
// # Overview
// 
// Image views can be static or editable. A static image view only displays
// the image that you specify. An editable image view object lets the user
// change the displayed image. You can also configure an image view to allow
// copying, pasting, deleting, and dragging of the image.
//
// # Specifying the image
//
//   - [NSImageView.SymbolConfiguration]
//   - [NSImageView.SetSymbolConfiguration]
//   - [NSImageView.Image]: The image displayed by the image view.
//   - [NSImageView.SetImage]
//
// # Specifying the visual characteristics
//
//   - [NSImageView.ImageFrameStyle]: The style of frame that appears around the image.
//   - [NSImageView.SetImageFrameStyle]
//   - [NSImageView.ImageAlignment]: The alignment of the cell’s image inside the image view.
//   - [NSImageView.SetImageAlignment]
//   - [NSImageView.ImageScaling]: The scaling mode applied to make the cell’s image fit the frame of the image view.
//   - [NSImageView.SetImageScaling]
//   - [NSImageView.Animates]: A Boolean value indicating whether the image view automatically plays animated images.
//   - [NSImageView.SetAnimates]
//   - [NSImageView.ContentTintColor]
//   - [NSImageView.SetContentTintColor]
//
// # Specifying the dynamic range
//
//   - [NSImageView.ImageDynamicRange]: The resolved dynamic range of the fully resolved image content.
//   - [NSImageView.PreferredImageDynamicRange]: The preferred dynamic range when displaying an image in the receiving image view.
//   - [NSImageView.SetPreferredImageDynamicRange]
//
// # Responding to user events
//
//   - [NSImageView.Editable]: A Boolean value indicating whether the user can drag a new image into the image view.
//   - [NSImageView.SetEditable]
//   - [NSImageView.AllowsCutCopyPaste]: A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//   - [NSImageView.SetAllowsCutCopyPaste]
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView
type NSImageView struct {
	NSControl
}

// NSImageViewFromID constructs a [NSImageView] from an objc.ID.
//
// A display of image data in a frame.
func NSImageViewFromID(id objc.ID) NSImageView {
	return NSImageView{NSControl: NSControlFromID(id)}
}
// NOTE: NSImageView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSImageView] class.
//
// # Specifying the image
//
//   - [INSImageView.SymbolConfiguration]
//   - [INSImageView.SetSymbolConfiguration]
//   - [INSImageView.Image]: The image displayed by the image view.
//   - [INSImageView.SetImage]
//
// # Specifying the visual characteristics
//
//   - [INSImageView.ImageFrameStyle]: The style of frame that appears around the image.
//   - [INSImageView.SetImageFrameStyle]
//   - [INSImageView.ImageAlignment]: The alignment of the cell’s image inside the image view.
//   - [INSImageView.SetImageAlignment]
//   - [INSImageView.ImageScaling]: The scaling mode applied to make the cell’s image fit the frame of the image view.
//   - [INSImageView.SetImageScaling]
//   - [INSImageView.Animates]: A Boolean value indicating whether the image view automatically plays animated images.
//   - [INSImageView.SetAnimates]
//   - [INSImageView.ContentTintColor]
//   - [INSImageView.SetContentTintColor]
//
// # Specifying the dynamic range
//
//   - [INSImageView.ImageDynamicRange]: The resolved dynamic range of the fully resolved image content.
//   - [INSImageView.PreferredImageDynamicRange]: The preferred dynamic range when displaying an image in the receiving image view.
//   - [INSImageView.SetPreferredImageDynamicRange]
//
// # Responding to user events
//
//   - [INSImageView.Editable]: A Boolean value indicating whether the user can drag a new image into the image view.
//   - [INSImageView.SetEditable]
//   - [INSImageView.AllowsCutCopyPaste]: A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//   - [INSImageView.SetAllowsCutCopyPaste]
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView
type INSImageView interface {
	INSControl
	NSAccessibilityImage
	NSMenuItemValidation

	// Topic: Specifying the image

	SymbolConfiguration() INSImageSymbolConfiguration
	SetSymbolConfiguration(value INSImageSymbolConfiguration)
	// The image displayed by the image view.
	Image() INSImage
	SetImage(value INSImage)

	// Topic: Specifying the visual characteristics

	// The style of frame that appears around the image.
	ImageFrameStyle() NSImageFrameStyle
	SetImageFrameStyle(value NSImageFrameStyle)
	// The alignment of the cell’s image inside the image view.
	ImageAlignment() NSImageAlignment
	SetImageAlignment(value NSImageAlignment)
	// The scaling mode applied to make the cell’s image fit the frame of the image view.
	ImageScaling() NSImageScaling
	SetImageScaling(value NSImageScaling)
	// A Boolean value indicating whether the image view automatically plays animated images.
	Animates() bool
	SetAnimates(value bool)
	ContentTintColor() INSColor
	SetContentTintColor(value INSColor)

	// Topic: Specifying the dynamic range

	// The resolved dynamic range of the fully resolved image content.
	ImageDynamicRange() NSImageDynamicRange
	// The preferred dynamic range when displaying an image in the receiving image view.
	PreferredImageDynamicRange() NSImageDynamicRange
	SetPreferredImageDynamicRange(value NSImageDynamicRange)

	// Topic: Responding to user events

	// A Boolean value indicating whether the user can drag a new image into the image view.
	Editable() bool
	SetEditable(value bool)
	// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
}

// Init initializes the instance.
func (i NSImageView) Init() NSImageView {
	rv := objc.Send[NSImageView](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSImageView) Autorelease() NSImageView {
	rv := objc.Send[NSImageView](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSImageView creates a new NSImageView instance.
func NewNSImageView() NSImageView {
	class := getNSImageViewClass()
	rv := objc.Send[NSImageView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewImageViewWithCoder(coder foundation.INSCoder) NSImageView {
	instance := getNSImageViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSImageViewFromID(rv)
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
func NewImageViewWithFrame(frameRect corefoundation.CGRect) NSImageView {
	instance := getNSImageViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSImageViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/init(image:)
func NewImageViewWithImage(image INSImage) NSImageView {
	rv := objc.Send[objc.ID](objc.ID(getNSImageViewClass().class), objc.Sel("imageViewWithImage:"), image)
	return NSImageViewFromID(rv)
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
// 
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (i NSImageView) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSImageView/symbolConfiguration
func (i NSImageView) SymbolConfiguration() INSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("symbolConfiguration"))
	return NSImageSymbolConfigurationFromID(objc.ID(rv))
}
func (i NSImageView) SetSymbolConfiguration(value INSImageSymbolConfiguration) {
	objc.Send[struct{}](i.ID, objc.Sel("setSymbolConfiguration:"), value)
}

// The image displayed by the image view.
//
// # Discussion
// 
// Use this property to change the image being displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/image
func (i NSImageView) Image() INSImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (i NSImageView) SetImage(value INSImage) {
	objc.Send[struct{}](i.ID, objc.Sel("setImage:"), value)
}

// The style of frame that appears around the image.
//
// # Discussion
// 
// The default value of this property is [NSImageView.FrameStyle.none].
//
// [NSImageView.FrameStyle.none]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/none
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/imageFrameStyle
func (i NSImageView) ImageFrameStyle() NSImageFrameStyle {
	rv := objc.Send[NSImageFrameStyle](i.ID, objc.Sel("imageFrameStyle"))
	return NSImageFrameStyle(rv)
}
func (i NSImageView) SetImageFrameStyle(value NSImageFrameStyle) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageFrameStyle:"), value)
}

// The alignment of the cell’s image inside the image view.
//
// # Discussion
// 
// The default value of this property is [NSImageAlignment.alignCenter].
//
// [NSImageAlignment.alignCenter]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignCenter
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/imageAlignment
func (i NSImageView) ImageAlignment() NSImageAlignment {
	rv := objc.Send[NSImageAlignment](i.ID, objc.Sel("imageAlignment"))
	return NSImageAlignment(rv)
}
func (i NSImageView) SetImageAlignment(value NSImageAlignment) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageAlignment:"), value)
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
// See: https://developer.apple.com/documentation/AppKit/NSImageView/imageScaling
func (i NSImageView) ImageScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](i.ID, objc.Sel("imageScaling"))
	return NSImageScaling(rv)
}
func (i NSImageView) SetImageScaling(value NSImageScaling) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageScaling:"), value)
}

// A Boolean value indicating whether the image view automatically plays
// animated images.
//
// # Discussion
// 
// When the value of this property is [true], the image view plays animated
// images automatically using the timing and looping characteristics stored in
// the image data. The default value of this property is [true].
// 
// Decoding an animated GIF file is the only way to create an animated
// [NSImage] object. If the image view has been assigned an image that was
// created from animated GIF data, setting this property to [true] enables
// automatic playback of the animation. If this property is set to [false],
// only the first frame of an animated image is displayed.
// 
// Loading an animated GIF file using an [NSImage] object produces an image
// that uses an [NSBitmapImageRep] object. The [currentFrame],
// [currentFrameDuration], and [frameCount] properties of the bitmap image
// representation determine the timing and looping characteristics of the
// animation. For more information, see [NSBitmapImageRep].
//
// [currentFrameDuration]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/currentFrameDuration
// [currentFrame]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/currentFrame
// [false]: https://developer.apple.com/documentation/Swift/false
// [frameCount]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey/frameCount
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/animates
func (i NSImageView) Animates() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("animates"))
	return rv
}
func (i NSImageView) SetAnimates(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setAnimates:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSImageView/contentTintColor
func (i NSImageView) ContentTintColor() INSColor {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("contentTintColor"))
	return NSColorFromID(objc.ID(rv))
}
func (i NSImageView) SetContentTintColor(value INSColor) {
	objc.Send[struct{}](i.ID, objc.Sel("setContentTintColor:"), value)
}

// The resolved dynamic range of the fully resolved image content.
//
// # Discussion
// 
// This property returns [NSImage.DynamicRange.unspecified] if the image view
// can’t resolve the image content or the image view hasn’t displayed.
//
// [NSImage.DynamicRange.unspecified]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/unspecified
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/imageDynamicRange
func (i NSImageView) ImageDynamicRange() NSImageDynamicRange {
	rv := objc.Send[NSImageDynamicRange](i.ID, objc.Sel("imageDynamicRange"))
	return NSImageDynamicRange(rv)
}

// The preferred dynamic range when displaying an image in the receiving image
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/preferredImageDynamicRange
func (i NSImageView) PreferredImageDynamicRange() NSImageDynamicRange {
	rv := objc.Send[NSImageDynamicRange](i.ID, objc.Sel("preferredImageDynamicRange"))
	return NSImageDynamicRange(rv)
}
func (i NSImageView) SetPreferredImageDynamicRange(value NSImageDynamicRange) {
	objc.Send[struct{}](i.ID, objc.Sel("setPreferredImageDynamicRange:"), value)
}

// A Boolean value indicating whether the user can drag a new image into the
// image view.
//
// # Discussion
// 
// When the value of this property is [true], the user can set the displayed
// image by dragging an image onto the image view. The default value of this
// property is [false], which causes the image view to display only the
// programmatically set image.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/isEditable
func (i NSImageView) Editable() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isEditable"))
	return rv
}
func (i NSImageView) SetEditable(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setEditable:"), value)
}

// A Boolean value indicating whether the image view lets the user cut, copy,
// and paste the image contents.
//
// # Discussion
// 
// When the value of this property is YES, the user can cut, copy, or paste
// the image in the image view.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/allowsCutCopyPaste
func (i NSImageView) AllowsCutCopyPaste() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("allowsCutCopyPaste"))
	return rv
}
func (i NSImageView) SetAllowsCutCopyPaste(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setAllowsCutCopyPaste:"), value)
}

// The default preferred image dynamic range.
//
// # Discussion
// 
// This property defaults to [NSImage.DynamicRange.constrainedHigh] on macOS
// 14 and higher, and [NSImage.DynamicRange.standard] otherwise. Set this
// property to another [NSImage.DynamicRange] value to change the default for
// all subsequently created image views in your app.
//
// [NSImage.DynamicRange.constrainedHigh]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/constrainedHigh
// [NSImage.DynamicRange.standard]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/standard
// [NSImage.DynamicRange]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
//
// See: https://developer.apple.com/documentation/AppKit/NSImageView/defaultPreferredImageDynamicRange
func (_NSImageViewClass NSImageViewClass) DefaultPreferredImageDynamicRange() NSImageDynamicRange {
	rv := objc.Send[NSImageDynamicRange](objc.ID(_NSImageViewClass.class), objc.Sel("defaultPreferredImageDynamicRange"))
	return NSImageDynamicRange(rv)
}
func (_NSImageViewClass NSImageViewClass) SetDefaultPreferredImageDynamicRange(value NSImageDynamicRange) {
	objc.Send[struct{}](objc.ID(_NSImageViewClass.class), objc.Sel("setDefaultPreferredImageDynamicRange:"), value)
}

			// Protocol methods for NSAccessibilityImage
			

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

func (o NSImageView) AccessibilityFrame() corefoundation.CGRect {
	
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

func (o NSImageView) AccessibilityParent() objectivec.IObject {
	
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

func (o NSImageView) AccessibilityIdentifier() string {
	
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

func (o NSImageView) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSMenuItemValidation
			

