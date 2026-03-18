// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSegmentedControl] class.
var (
	_NSSegmentedControlClass     NSSegmentedControlClass
	_NSSegmentedControlClassOnce sync.Once
)

func getNSSegmentedControlClass() NSSegmentedControlClass {
	_NSSegmentedControlClassOnce.Do(func() {
		_NSSegmentedControlClass = NSSegmentedControlClass{class: objc.GetClass("NSSegmentedControl")}
	})
	return _NSSegmentedControlClass
}

// GetNSSegmentedControlClass returns the class object for NSSegmentedControl.
func GetNSSegmentedControlClass() NSSegmentedControlClass {
	return getNSSegmentedControlClass()
}

type NSSegmentedControlClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSegmentedControlClass) Alloc() NSSegmentedControl {
	rv := objc.Send[NSSegmentedControl](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Display one or more buttons in a single horizontal group.
//
// # Overview
// 
// The [NSSegmentedControl] class uses an [NSSegmentedCell] class to implement
// much of the control’s functionality. Most methods in [NSSegmentedControl]
// are simply cover methods that call the corresponding method in
// [NSSegmentedCell]. The methods of [NSSegmentedCell] that do not have covers
// relate to accessing and setting values for tags and tooltips,
// programatically setting the key segment, and establishing the mode of the
// control.
// 
// The features of a segmented control include the following:
// 
// - A segment can have an image, text (label), menu, tooltip, and tag. - A
// segmented control can contain images or text, but not both. - Either the
// control or individual segments can be enabled or disabled. - Segmented
// controls have four tracking modes, described in
// [NSSegmentedControl.SwitchTracking]. You use these modes with the
// [NSSegmentedControl.TrackingMode] property. - Each segment can be either a fixed width or
// autosized to fit the contents. - If a segment has text and is marked as
// autosizing, then the text may be truncated so that the control completely
// fits. - If an image is too large to fit in a segment, it is clipped. - If
// Full Keyboard Access is enabled in System Preferences > Keyboard, the
// keyboard may be used to move between and select segments.
//
// [NSSegmentedControl.SwitchTracking]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
//
// # Specifying the segment behavior
//
//   - [NSSegmentedControl.TrackingMode]: The type of tracking behavior the control exhibits.
//   - [NSSegmentedControl.SetTrackingMode]
//   - [NSSegmentedControl.SegmentStyle]: The visual style used to display the control.
//   - [NSSegmentedControl.SetSegmentStyle]
//
// # Specifying number of segments
//
//   - [NSSegmentedControl.SegmentCount]: The number of segments in the control.
//   - [NSSegmentedControl.SetSegmentCount]
//
// # Configuring the segment text
//
//   - [NSSegmentedControl.LabelForSegment]: Returns the label of the specified segment.
//   - [NSSegmentedControl.SetLabelForSegment]: Sets the label for the specified segment.
//   - [NSSegmentedControl.SetAlignmentForSegment]
//   - [NSSegmentedControl.AlignmentForSegment]
//
// # Configuring a segment image
//
//   - [NSSegmentedControl.SetImageForSegment]: Sets the image for the specified segment.
//   - [NSSegmentedControl.ImageForSegment]: Returns the image associated with the specified segment.
//   - [NSSegmentedControl.SetImageScalingForSegment]: Sets the scaling mode used to display the specified segment’s image.
//   - [NSSegmentedControl.ImageScalingForSegment]: Returns the scaling mode used to display the specified segment’s image.
//
// # Configuring a segment menu
//
//   - [NSSegmentedControl.SetMenuForSegment]: Sets the menu for the specified segment.
//   - [NSSegmentedControl.MenuForSegment]: Returns the menu for the specified segment.
//   - [NSSegmentedControl.SetShowsMenuIndicatorForSegment]
//   - [NSSegmentedControl.ShowsMenuIndicatorForSegment]
//   - [NSSegmentedControl.SpringLoaded]: A Boolean value that indicates whether spring loading is enabled for the control.
//   - [NSSegmentedControl.SetSpringLoaded]
//
// # Managing the selected segment
//
//   - [NSSegmentedControl.SelectedSegment]: The index of the selected segment of the control, or `-1` if no segment is selected.
//   - [NSSegmentedControl.SetSelectedSegment]
//   - [NSSegmentedControl.IndexOfSelectedItem]
//   - [NSSegmentedControl.SelectSegmentWithTag]: Selects the segment with the specified tag.
//   - [NSSegmentedControl.SetSelectedForSegment]: Sets the selection state of the specified segment.
//   - [NSSegmentedControl.IsSelectedForSegment]: Returns a Boolean value indicating whether the specified segment is selected.
//   - [NSSegmentedControl.SelectedSegmentBezelColor]: The color of the selected segment’s bezel, in appearances that support it.
//   - [NSSegmentedControl.SetSelectedSegmentBezelColor]
//   - [NSSegmentedControl.DoubleValueForSelectedSegment]: When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// # Adjusting the segment spacing
//
//   - [NSSegmentedControl.SetWidthForSegment]: Sets the width of the specified segment.
//   - [NSSegmentedControl.WidthForSegment]: Returns the width of the specified segment.
//   - [NSSegmentedControl.SegmentDistribution]
//   - [NSSegmentedControl.SetSegmentDistribution]
//
// # Specifying the border shape
//
//   - [NSSegmentedControl.BorderShape]
//   - [NSSegmentedControl.SetBorderShape]
//
// # Enabling and disabling segments
//
//   - [NSSegmentedControl.SetEnabledForSegment]: Sets the enabled state of the specified segment
//   - [NSSegmentedControl.IsEnabledForSegment]: Returns a Boolean value indicating whether the specified segment is enabled.
//
// # Managing tags and tooltips
//
//   - [NSSegmentedControl.TagForSegment]
//   - [NSSegmentedControl.SetTagForSegment]
//   - [NSSegmentedControl.SetToolTipForSegment]
//   - [NSSegmentedControl.ToolTipForSegment]
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl
type NSSegmentedControl struct {
	NSControl
}

// NSSegmentedControlFromID constructs a [NSSegmentedControl] from an objc.ID.
//
// Display one or more buttons in a single horizontal group.
func NSSegmentedControlFromID(id objc.ID) NSSegmentedControl {
	return NSSegmentedControl{NSControl: NSControlFromID(id)}
}
// NOTE: NSSegmentedControl adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSegmentedControl] class.
//
// # Specifying the segment behavior
//
//   - [INSSegmentedControl.TrackingMode]: The type of tracking behavior the control exhibits.
//   - [INSSegmentedControl.SetTrackingMode]
//   - [INSSegmentedControl.SegmentStyle]: The visual style used to display the control.
//   - [INSSegmentedControl.SetSegmentStyle]
//
// # Specifying number of segments
//
//   - [INSSegmentedControl.SegmentCount]: The number of segments in the control.
//   - [INSSegmentedControl.SetSegmentCount]
//
// # Configuring the segment text
//
//   - [INSSegmentedControl.LabelForSegment]: Returns the label of the specified segment.
//   - [INSSegmentedControl.SetLabelForSegment]: Sets the label for the specified segment.
//   - [INSSegmentedControl.SetAlignmentForSegment]
//   - [INSSegmentedControl.AlignmentForSegment]
//
// # Configuring a segment image
//
//   - [INSSegmentedControl.SetImageForSegment]: Sets the image for the specified segment.
//   - [INSSegmentedControl.ImageForSegment]: Returns the image associated with the specified segment.
//   - [INSSegmentedControl.SetImageScalingForSegment]: Sets the scaling mode used to display the specified segment’s image.
//   - [INSSegmentedControl.ImageScalingForSegment]: Returns the scaling mode used to display the specified segment’s image.
//
// # Configuring a segment menu
//
//   - [INSSegmentedControl.SetMenuForSegment]: Sets the menu for the specified segment.
//   - [INSSegmentedControl.MenuForSegment]: Returns the menu for the specified segment.
//   - [INSSegmentedControl.SetShowsMenuIndicatorForSegment]
//   - [INSSegmentedControl.ShowsMenuIndicatorForSegment]
//   - [INSSegmentedControl.SpringLoaded]: A Boolean value that indicates whether spring loading is enabled for the control.
//   - [INSSegmentedControl.SetSpringLoaded]
//
// # Managing the selected segment
//
//   - [INSSegmentedControl.SelectedSegment]: The index of the selected segment of the control, or `-1` if no segment is selected.
//   - [INSSegmentedControl.SetSelectedSegment]
//   - [INSSegmentedControl.IndexOfSelectedItem]
//   - [INSSegmentedControl.SelectSegmentWithTag]: Selects the segment with the specified tag.
//   - [INSSegmentedControl.SetSelectedForSegment]: Sets the selection state of the specified segment.
//   - [INSSegmentedControl.IsSelectedForSegment]: Returns a Boolean value indicating whether the specified segment is selected.
//   - [INSSegmentedControl.SelectedSegmentBezelColor]: The color of the selected segment’s bezel, in appearances that support it.
//   - [INSSegmentedControl.SetSelectedSegmentBezelColor]
//   - [INSSegmentedControl.DoubleValueForSelectedSegment]: When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// # Adjusting the segment spacing
//
//   - [INSSegmentedControl.SetWidthForSegment]: Sets the width of the specified segment.
//   - [INSSegmentedControl.WidthForSegment]: Returns the width of the specified segment.
//   - [INSSegmentedControl.SegmentDistribution]
//   - [INSSegmentedControl.SetSegmentDistribution]
//
// # Specifying the border shape
//
//   - [INSSegmentedControl.BorderShape]
//   - [INSSegmentedControl.SetBorderShape]
//
// # Enabling and disabling segments
//
//   - [INSSegmentedControl.SetEnabledForSegment]: Sets the enabled state of the specified segment
//   - [INSSegmentedControl.IsEnabledForSegment]: Returns a Boolean value indicating whether the specified segment is enabled.
//
// # Managing tags and tooltips
//
//   - [INSSegmentedControl.TagForSegment]
//   - [INSSegmentedControl.SetTagForSegment]
//   - [INSSegmentedControl.SetToolTipForSegment]
//   - [INSSegmentedControl.ToolTipForSegment]
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl
type INSSegmentedControl interface {
	INSControl
	NSUserInterfaceCompression

	// Topic: Specifying the segment behavior

	// The type of tracking behavior the control exhibits.
	TrackingMode() NSSegmentSwitchTracking
	SetTrackingMode(value NSSegmentSwitchTracking)
	// The visual style used to display the control.
	SegmentStyle() NSSegmentStyle
	SetSegmentStyle(value NSSegmentStyle)

	// Topic: Specifying number of segments

	// The number of segments in the control.
	SegmentCount() int
	SetSegmentCount(value int)

	// Topic: Configuring the segment text

	// Returns the label of the specified segment.
	LabelForSegment(segment int) string
	// Sets the label for the specified segment.
	SetLabelForSegment(label string, segment int)
	SetAlignmentForSegment(alignment NSTextAlignment, segment int)
	AlignmentForSegment(segment int) NSTextAlignment

	// Topic: Configuring a segment image

	// Sets the image for the specified segment.
	SetImageForSegment(image INSImage, segment int)
	// Returns the image associated with the specified segment.
	ImageForSegment(segment int) INSImage
	// Sets the scaling mode used to display the specified segment’s image.
	SetImageScalingForSegment(scaling NSImageScaling, segment int)
	// Returns the scaling mode used to display the specified segment’s image.
	ImageScalingForSegment(segment int) NSImageScaling

	// Topic: Configuring a segment menu

	// Sets the menu for the specified segment.
	SetMenuForSegment(menu INSMenu, segment int)
	// Returns the menu for the specified segment.
	MenuForSegment(segment int) INSMenu
	SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int)
	ShowsMenuIndicatorForSegment(segment int) bool
	// A Boolean value that indicates whether spring loading is enabled for the control.
	SpringLoaded() bool
	SetSpringLoaded(value bool)

	// Topic: Managing the selected segment

	// The index of the selected segment of the control, or `-1` if no segment is selected.
	SelectedSegment() int
	SetSelectedSegment(value int)
	IndexOfSelectedItem() int
	// Selects the segment with the specified tag.
	SelectSegmentWithTag(tag int) bool
	// Sets the selection state of the specified segment.
	SetSelectedForSegment(selected bool, segment int)
	// Returns a Boolean value indicating whether the specified segment is selected.
	IsSelectedForSegment(segment int) bool
	// The color of the selected segment’s bezel, in appearances that support it.
	SelectedSegmentBezelColor() INSColor
	SetSelectedSegmentBezelColor(value INSColor)
	// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
	DoubleValueForSelectedSegment() float64

	// Topic: Adjusting the segment spacing

	// Sets the width of the specified segment.
	SetWidthForSegment(width float64, segment int)
	// Returns the width of the specified segment.
	WidthForSegment(segment int) float64
	SegmentDistribution() NSSegmentDistribution
	SetSegmentDistribution(value NSSegmentDistribution)

	// Topic: Specifying the border shape

	BorderShape() NSControlBorderShape
	SetBorderShape(value NSControlBorderShape)

	// Topic: Enabling and disabling segments

	// Sets the enabled state of the specified segment
	SetEnabledForSegment(enabled bool, segment int)
	// Returns a Boolean value indicating whether the specified segment is enabled.
	IsEnabledForSegment(segment int) bool

	// Topic: Managing tags and tooltips

	TagForSegment(segment int) int
	SetTagForSegment(tag int, segment int)
	SetToolTipForSegment(toolTip string, segment int)
	ToolTipForSegment(segment int) string
}

// Init initializes the instance.
func (s NSSegmentedControl) Init() NSSegmentedControl {
	rv := objc.Send[NSSegmentedControl](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSegmentedControl) Autorelease() NSSegmentedControl {
	rv := objc.Send[NSSegmentedControl](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSegmentedControl creates a new NSSegmentedControl instance.
func NewNSSegmentedControl() NSSegmentedControl {
	class := getNSSegmentedControlClass()
	rv := objc.Send[NSSegmentedControl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewSegmentedControlWithCoder(coder foundation.INSCoder) NSSegmentedControl {
	instance := getNSSegmentedControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSegmentedControlFromID(rv)
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
func NewSegmentedControlWithFrame(frameRect corefoundation.CGRect) NSSegmentedControl {
	instance := getNSSegmentedControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSegmentedControlFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(images:trackingMode:target:action:)
func NewSegmentedControlWithImagesTrackingModeTargetAction(images []NSImage, trackingMode NSSegmentSwitchTracking, target objectivec.IObject, action objc.SEL) NSSegmentedControl {
	rv := objc.Send[objc.ID](objc.ID(getNSSegmentedControlClass().class), objc.Sel("segmentedControlWithImages:trackingMode:target:action:"), objectivec.IObjectSliceToNSArray(images), trackingMode, target, action)
	return NSSegmentedControlFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(labels:trackingMode:target:action:)
func NewSegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode NSSegmentSwitchTracking, target objectivec.IObject, action objc.SEL) NSSegmentedControl {
	rv := objc.Send[objc.ID](objc.ID(getNSSegmentedControlClass().class), objc.Sel("segmentedControlWithLabels:trackingMode:target:action:"), objectivec.StringSliceToNSArray(labels), trackingMode, target, action)
	return NSSegmentedControlFromID(rv)
}

// Returns the label of the specified segment.
//
// segment: The index of the segment whose label you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// The label of the segment. The returned string contains the entire text of
// the label, even if that text is normally truncated during drawing.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/label(forSegment:)
func (s NSSegmentedControl) LabelForSegment(segment int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("labelForSegment:"), segment)
	return foundation.NSStringFromID(rv).String()
}

// Sets the label for the specified segment.
//
// label: The label you want to display in the segment. If the width of the string is
// greater than the width of the segment, the string’s text is truncated
// during drawing.
//
// segment: The index of the segment whose label you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setLabel(_:forSegment:)
func (s NSSegmentedControl) SetLabelForSegment(label string, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setLabel:forSegment:"), objc.String(label), segment)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setAlignment(_:forSegment:)
func (s NSSegmentedControl) SetAlignmentForSegment(alignment NSTextAlignment, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAlignment:forSegment:"), alignment, segment)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/alignment(forSegment:)
func (s NSSegmentedControl) AlignmentForSegment(segment int) NSTextAlignment {
	rv := objc.Send[NSTextAlignment](s.ID, objc.Sel("alignmentForSegment:"), segment)
	return NSTextAlignment(rv)
}

// Sets the image for the specified segment.
//
// image: The image to apply to the segment or `nil` if you want to clear the
// existing image. Images are not scaled to fit inside a segment. If the image
// is larger than the available space, it is clipped.
//
// segment: The index of the segment whose image you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setImage(_:forSegment:)
func (s NSSegmentedControl) SetImageForSegment(image INSImage, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setImage:forSegment:"), image, segment)
}

// Returns the image associated with the specified segment.
//
// segment: The index of the segment whose image you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// The image associated with the segment; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/image(forSegment:)
func (s NSSegmentedControl) ImageForSegment(segment int) INSImage {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("imageForSegment:"), segment)
	return NSImageFromID(rv)
}

// Sets the scaling mode used to display the specified segment’s image.
//
// scaling: One of the image scaling constants. For a list of possible values, see
// [NSImageScaling].
// //
// [NSImageScaling]: https://developer.apple.com/documentation/AppKit/NSImageScaling
//
// segment: The index of the segment whose enabled state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setImageScaling(_:forSegment:)
func (s NSSegmentedControl) SetImageScalingForSegment(scaling NSImageScaling, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}

// Returns the scaling mode used to display the specified segment’s image.
//
// segment: The index of the segment whose enabled state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// One of the image scaling constants. For a list of possible values, see
// [NSImageScaling]. The default value is
// [NSImageScaling.scaleProportionallyDown].
//
// [NSImageScaling.scaleProportionallyDown]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
// [NSImageScaling]: https://developer.apple.com/documentation/AppKit/NSImageScaling
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/imageScaling(forSegment:)
func (s NSSegmentedControl) ImageScalingForSegment(segment int) NSImageScaling {
	rv := objc.Send[NSImageScaling](s.ID, objc.Sel("imageScalingForSegment:"), segment)
	return NSImageScaling(rv)
}

// Sets the menu for the specified segment.
//
// menu: The menu you want to add to the segment or `nil` to clear the current menu.
//
// segment: The index of the segment whose menu you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Discussion
// 
// Adding a menu to a segment allows the segment to be used as a pop-up
// button. If the segment has a menu only, then the menu displays when the
// user clicks the segment. If the segment has both a menu and an action, then
// the action triggers when the user clicks the segment and the menu displays
// when the user clicks and holds the segment.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setMenu(_:forSegment:)
func (s NSSegmentedControl) SetMenuForSegment(menu INSMenu, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setMenu:forSegment:"), menu, segment)
}

// Returns the menu for the specified segment.
//
// segment: The index of the segment whose menu you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// The menu associated with the segment; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/menu(forSegment:)
func (s NSSegmentedControl) MenuForSegment(segment int) INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("menuForSegment:"), segment)
	return NSMenuFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setShowsMenuIndicator(_:forSegment:)
func (s NSSegmentedControl) SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setShowsMenuIndicator:forSegment:"), showsMenuIndicator, segment)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/showsMenuIndicator(forSegment:)
func (s NSSegmentedControl) ShowsMenuIndicatorForSegment(segment int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsMenuIndicatorForSegment:"), segment)
	return rv
}

// Selects the segment with the specified tag.
//
// tag: The tag associated with the desired segment. A tag is an integer value that
// can be assigned to a segment as a way of identifying it without knowing its
// position in the control.
//
// # Return Value
// 
// [true] if the segment was selected successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Typically, you use Interface Builder to specify the tag for each segment.
// You may also set this value programmatically using the `` method of
// [NSSegmentedCell].
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectSegment(withTag:)
func (s NSSegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}

// Sets the selection state of the specified segment.
//
// selected: [true] if you want to select the segment; otherwise, [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// segment: The index of the segment whose selection state you want to set. This method
// raises an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Discussion
// 
// If the control allows only a single selection, this method deselects any
// other selected segments.
// 
// If the [TrackingMode] property of the segment is set to
// [NSSegmentedControl.SwitchTracking.momentary], then attempting to set the
// selected state of the segment will have no effect.
//
// [NSSegmentedControl.SwitchTracking.momentary]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentary
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setSelected(_:forSegment:)
func (s NSSegmentedControl) SetSelectedForSegment(selected bool, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSelected:forSegment:"), selected, segment)
}

// Returns a Boolean value indicating whether the specified segment is
// selected.
//
// segment: The index of the segment whose selection state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// [true] if the segment is selected; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isSelected(forSegment:)
func (s NSSegmentedControl) IsSelectedForSegment(segment int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}

// Sets the width of the specified segment.
//
// width: The width of the segment, measured in points. Specify the value `0` if you
// want the segment to be sized to fit the available space automatically.
//
// segment: The index of the segment whose width you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setWidth(_:forSegment:)
func (s NSSegmentedControl) SetWidthForSegment(width float64, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setWidth:forSegment:"), width, segment)
}

// Returns the width of the specified segment.
//
// segment: The index of the segment whose width you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// The width of the segment, measured in points, or 0 if the segment is sized
// to fit the available space automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/width(forSegment:)
func (s NSSegmentedControl) WidthForSegment(segment int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("widthForSegment:"), segment)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/compress(withPrioritizedCompressionOptions:)
func (s NSSegmentedControl) CompressWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) {
	objc.Send[objc.ID](s.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/minimumSize(withPrioritizedCompressionOptions:)
func (s NSSegmentedControl) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
	return corefoundation.CGSize(rv)
}

// Sets the enabled state of the specified segment
//
// enabled: [true] to enable the segment; otherwise, [false] to disable it.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// segment: The index of the segment you want to enable or disable. This method raises
// an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setEnabled(_:forSegment:)
func (s NSSegmentedControl) SetEnabledForSegment(enabled bool, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}

// Returns a Boolean value indicating whether the specified segment is
// enabled.
//
// segment: The index of the segment whose enabled state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// [true] if the segment is enabled; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isEnabled(forSegment:)
func (s NSSegmentedControl) IsEnabledForSegment(segment int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/tag(forSegment:)
func (s NSSegmentedControl) TagForSegment(segment int) int {
	rv := objc.Send[int](s.ID, objc.Sel("tagForSegment:"), segment)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setTag(_:forSegment:)
func (s NSSegmentedControl) SetTagForSegment(tag int, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setTag:forSegment:"), tag, segment)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setToolTip(_:forSegment:)
func (s NSSegmentedControl) SetToolTipForSegment(toolTip string, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setToolTip:forSegment:"), objc.String(toolTip), segment)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/toolTip(forSegment:)
func (s NSSegmentedControl) ToolTipForSegment(segment int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("toolTipForSegment:"), segment)
	return foundation.NSStringFromID(rv).String()
}

// The type of tracking behavior the control exhibits.
//
// # Discussion
// 
// An [NSSegmentedControl.SwitchTracking] value specifies how the control
// responds when the user presses a keyboard key or clicks, force clicks
// (applies pressure in a pressure-sensitive system), releases pressure, and
// so on. For possible values see [NSSegmentedControl.SwitchTracking].
//
// [NSSegmentedControl.SwitchTracking]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/trackingMode
func (s NSSegmentedControl) TrackingMode() NSSegmentSwitchTracking {
	rv := objc.Send[NSSegmentSwitchTracking](s.ID, objc.Sel("trackingMode"))
	return NSSegmentSwitchTracking(rv)
}
func (s NSSegmentedControl) SetTrackingMode(value NSSegmentSwitchTracking) {
	objc.Send[struct{}](s.ID, objc.Sel("setTrackingMode:"), value)
}

// The visual style used to display the control.
//
// # Discussion
// 
// An [NSSegmentStyle] value that specifies the visual display used by the
// control. For possible values, see [NSSegmentedControl.Style].
//
// [NSSegmentedControl.Style]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentStyle
func (s NSSegmentedControl) SegmentStyle() NSSegmentStyle {
	rv := objc.Send[NSSegmentStyle](s.ID, objc.Sel("segmentStyle"))
	return NSSegmentStyle(rv)
}
func (s NSSegmentedControl) SetSegmentStyle(value NSSegmentStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setSegmentStyle:"), value)
}

// The number of segments in the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentCount
func (s NSSegmentedControl) SegmentCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("segmentCount"))
	return rv
}
func (s NSSegmentedControl) SetSegmentCount(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSegmentCount:"), value)
}

// A Boolean value that indicates whether spring loading is enabled for the
// control.
//
// # Discussion
// 
// The value of this property is [true] if spring loading is enabled for the
// control, and [false] if it is not. The default is [false].
// 
// On pressure-sensitive systems, such as systems with the Force Touch
// trackpad, spring loading is a feature that allows a user to activate a
// segment in a segmented control by dragging selected items over it and force
// clicking—pressing harder—without dropping the selected items. The user
// can then continue dragging the items, possibly to perform additional
// actions.
// 
// A practical example of this feature can be found in the Calendar app. A
// selected calendar event can be dragged over the day, week, month, or year
// segments in the toolbar. Force clicking on a segment switches the calendar
// view without releasing the selected calendar event. The calendar event can
// then be dropped at the desired location in the new calendar view.
// 
// When spring loading is enabled on a segmented control and a user drags
// something over a segment, the segment highlights to indicate that it
// responds to force clicking. In this situation, if the user presses harder,
// additional highlighting occurs to indicate that the segment was activated.
// 
// On systems that don’t support pressure sensitivity, simply hovering over
// the segment for a short period of time is sufficient to activate the
// segment.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isSpringLoaded
func (s NSSegmentedControl) SpringLoaded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSpringLoaded"))
	return rv
}
func (s NSSegmentedControl) SetSpringLoaded(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpringLoaded:"), value)
}

// The index of the selected segment of the control, or `-1` if no segment is
// selected.
//
// # Discussion
// 
// If the control allows multiple selections, this property contains the most
// recently selected segment. If the index is out of bounds, an exception
// ([rangeException]) is raised.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegment
func (s NSSegmentedControl) SelectedSegment() int {
	rv := objc.Send[int](s.ID, objc.Sel("selectedSegment"))
	return rv
}
func (s NSSegmentedControl) SetSelectedSegment(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectedSegment:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/indexOfSelectedItem
func (s NSSegmentedControl) IndexOfSelectedItem() int {
	rv := objc.Send[int](s.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The color of the selected segment’s bezel, in appearances that support
// it.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s NSSegmentedControl) SelectedSegmentBezelColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectedSegmentBezelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (s NSSegmentedControl) SetSelectedSegmentBezelColor(value INSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}

// When the tracking mode for the control is set to use a momentary
// accelerator, returns a value for the selected segment.
//
// # Return Value
// 
// The value of the selected segment interpreted as a double-precision
// floating-point number.
// 
// # Discussion
// 
// This method is intended for use with controls whose tracking mode is set to
// [NSSegmentedControl.SwitchTracking.momentaryAccelerator]. An assertion will
// occur if this method is called for other types of segmented controls.
//
// [NSSegmentedControl.SwitchTracking.momentaryAccelerator]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentaryAccelerator
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/doubleValueForSelectedSegment
func (s NSSegmentedControl) DoubleValueForSelectedSegment() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("doubleValueForSelectedSegment"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentDistribution
func (s NSSegmentedControl) SegmentDistribution() NSSegmentDistribution {
	rv := objc.Send[NSSegmentDistribution](s.ID, objc.Sel("segmentDistribution"))
	return NSSegmentDistribution(rv)
}
func (s NSSegmentedControl) SetSegmentDistribution(value NSSegmentDistribution) {
	objc.Send[struct{}](s.ID, objc.Sel("setSegmentDistribution:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/activeCompressionOptions
func (s NSSegmentedControl) ActiveCompressionOptions() INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("activeCompressionOptions"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/borderShape
func (s NSSegmentedControl) BorderShape() NSControlBorderShape {
	rv := objc.Send[NSControlBorderShape](s.ID, objc.Sel("borderShape"))
	return NSControlBorderShape(rv)
}
func (s NSSegmentedControl) SetBorderShape(value NSControlBorderShape) {
	objc.Send[struct{}](s.ID, objc.Sel("setBorderShape:"), value)
}

			// Protocol methods for NSUserInterfaceCompression
			

