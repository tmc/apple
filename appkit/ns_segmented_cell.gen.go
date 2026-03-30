// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSegmentedCell] class.
var (
	_NSSegmentedCellClass     NSSegmentedCellClass
	_NSSegmentedCellClassOnce sync.Once
)

func getNSSegmentedCellClass() NSSegmentedCellClass {
	_NSSegmentedCellClassOnce.Do(func() {
		_NSSegmentedCellClass = NSSegmentedCellClass{class: objc.GetClass("NSSegmentedCell")}
	})
	return _NSSegmentedCellClass
}

// GetNSSegmentedCellClass returns the class object for NSSegmentedCell.
func GetNSSegmentedCellClass() NSSegmentedCellClass {
	return getNSSegmentedCellClass()
}

type NSSegmentedCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSegmentedCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSegmentedCellClass) Alloc() NSSegmentedCell {
	rv := objc.Send[NSSegmentedCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An [NSSegmentedCell] object implements the appearance and behavior of a
// horizontal button divided into multiple segments. This class is used in
// conjunction with the [NSSegmentedControl] class to implement a segmented
// control.
//
// # Overview
//
// Use the methods of [NSSegmentedCell] to customize the attributes of a
// segmented control. To customize the appearance of individual segments, you
// can also subclass and override the [NSSegmentedCell.DrawSegmentInFrameWithView] method.
//
// # Specifying the Number of Segments
//
//   - [NSSegmentedCell.SegmentCount]: The number of segments in the segmented control.
//   - [NSSegmentedCell.SetSegmentCount]
//
// # Specifying the Selected Segment
//
//   - [NSSegmentedCell.SetSelectedForSegment]: Sets the selection state of the specified segment.
//   - [NSSegmentedCell.SelectSegmentWithTag]: Selects the segment with the specified tag.
//   - [NSSegmentedCell.MakeNextSegmentKey]: Selects the next segment.
//   - [NSSegmentedCell.MakePreviousSegmentKey]: Selects the previous segment.
//   - [NSSegmentedCell.SelectedSegment]: The index of the selected segment of the control, or `-1` if no segment is selected.
//   - [NSSegmentedCell.SetSelectedSegment]
//   - [NSSegmentedCell.IsSelectedForSegment]: Returns a Boolean value indicating whether the specified segment is selected,
//
// # Specifying the Tracking Mode
//
//   - [NSSegmentedCell.TrackingMode]: The tracking mode used for the segments of the control.
//   - [NSSegmentedCell.SetTrackingMode]
//
// # Configuring Individual Segments
//
//   - [NSSegmentedCell.SetLabelForSegment]: Sets the label for the specified segment.
//   - [NSSegmentedCell.LabelForSegment]: Returns the label of the specified segment.
//   - [NSSegmentedCell.SetImageForSegment]: Sets the image for the specified segment.
//   - [NSSegmentedCell.ImageForSegment]: Returns the image associated with the specified segment.
//   - [NSSegmentedCell.SetImageScalingForSegment]: Sets the image scaling mode for the specified segment.
//   - [NSSegmentedCell.ImageScalingForSegment]: Returns the image scaling mode associated with the specified segment.
//   - [NSSegmentedCell.SetWidthForSegment]: Sets the width of the specified segment.
//   - [NSSegmentedCell.WidthForSegment]: Returns the width of the specified segment.
//   - [NSSegmentedCell.SetEnabledForSegment]: Sets the enabled state of the specified segment
//   - [NSSegmentedCell.IsEnabledForSegment]: Returns a Boolean value indicating whether the specified segment is enabled.
//   - [NSSegmentedCell.SetMenuForSegment]: Sets the menu for the specified segment.
//   - [NSSegmentedCell.MenuForSegment]: Returns the menu for the specified segment.
//   - [NSSegmentedCell.SetToolTipForSegment]: Sets the tooltip for the specified segment.
//   - [NSSegmentedCell.ToolTipForSegment]: Returns the tooltip of the specified segment.
//   - [NSSegmentedCell.SetTagForSegment]: Sets the tag for the specified segment.
//   - [NSSegmentedCell.TagForSegment]: Returns the tag of the specified segment.
//
// # Drawing Custom Content
//
//   - [NSSegmentedCell.DrawSegmentInFrameWithView]: Draws the image and label of the segment in the specified view.
//
// # Specifying Segment Visual Styles
//
//   - [NSSegmentedCell.InteriorBackgroundStyleForSegment]: Returns the interior background style for the specified segment.
//   - [NSSegmentedCell.SegmentStyle]: The visual style used to display the segmented control.
//   - [NSSegmentedCell.SetSegmentStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell
type NSSegmentedCell struct {
	NSActionCell
}

// NSSegmentedCellFromID constructs a [NSSegmentedCell] from an objc.ID.
//
// An [NSSegmentedCell] object implements the appearance and behavior of a
// horizontal button divided into multiple segments. This class is used in
// conjunction with the [NSSegmentedControl] class to implement a segmented
// control.
func NSSegmentedCellFromID(id objc.ID) NSSegmentedCell {
	return NSSegmentedCell{NSActionCell: NSActionCellFromID(id)}
}

// NOTE: NSSegmentedCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSegmentedCell] class.
//
// # Specifying the Number of Segments
//
//   - [INSSegmentedCell.SegmentCount]: The number of segments in the segmented control.
//   - [INSSegmentedCell.SetSegmentCount]
//
// # Specifying the Selected Segment
//
//   - [INSSegmentedCell.SetSelectedForSegment]: Sets the selection state of the specified segment.
//   - [INSSegmentedCell.SelectSegmentWithTag]: Selects the segment with the specified tag.
//   - [INSSegmentedCell.MakeNextSegmentKey]: Selects the next segment.
//   - [INSSegmentedCell.MakePreviousSegmentKey]: Selects the previous segment.
//   - [INSSegmentedCell.SelectedSegment]: The index of the selected segment of the control, or `-1` if no segment is selected.
//   - [INSSegmentedCell.SetSelectedSegment]
//   - [INSSegmentedCell.IsSelectedForSegment]: Returns a Boolean value indicating whether the specified segment is selected,
//
// # Specifying the Tracking Mode
//
//   - [INSSegmentedCell.TrackingMode]: The tracking mode used for the segments of the control.
//   - [INSSegmentedCell.SetTrackingMode]
//
// # Configuring Individual Segments
//
//   - [INSSegmentedCell.SetLabelForSegment]: Sets the label for the specified segment.
//   - [INSSegmentedCell.LabelForSegment]: Returns the label of the specified segment.
//   - [INSSegmentedCell.SetImageForSegment]: Sets the image for the specified segment.
//   - [INSSegmentedCell.ImageForSegment]: Returns the image associated with the specified segment.
//   - [INSSegmentedCell.SetImageScalingForSegment]: Sets the image scaling mode for the specified segment.
//   - [INSSegmentedCell.ImageScalingForSegment]: Returns the image scaling mode associated with the specified segment.
//   - [INSSegmentedCell.SetWidthForSegment]: Sets the width of the specified segment.
//   - [INSSegmentedCell.WidthForSegment]: Returns the width of the specified segment.
//   - [INSSegmentedCell.SetEnabledForSegment]: Sets the enabled state of the specified segment
//   - [INSSegmentedCell.IsEnabledForSegment]: Returns a Boolean value indicating whether the specified segment is enabled.
//   - [INSSegmentedCell.SetMenuForSegment]: Sets the menu for the specified segment.
//   - [INSSegmentedCell.MenuForSegment]: Returns the menu for the specified segment.
//   - [INSSegmentedCell.SetToolTipForSegment]: Sets the tooltip for the specified segment.
//   - [INSSegmentedCell.ToolTipForSegment]: Returns the tooltip of the specified segment.
//   - [INSSegmentedCell.SetTagForSegment]: Sets the tag for the specified segment.
//   - [INSSegmentedCell.TagForSegment]: Returns the tag of the specified segment.
//
// # Drawing Custom Content
//
//   - [INSSegmentedCell.DrawSegmentInFrameWithView]: Draws the image and label of the segment in the specified view.
//
// # Specifying Segment Visual Styles
//
//   - [INSSegmentedCell.InteriorBackgroundStyleForSegment]: Returns the interior background style for the specified segment.
//   - [INSSegmentedCell.SegmentStyle]: The visual style used to display the segmented control.
//   - [INSSegmentedCell.SetSegmentStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell
type INSSegmentedCell interface {
	INSActionCell

	// Topic: Specifying the Number of Segments

	// The number of segments in the segmented control.
	SegmentCount() int
	SetSegmentCount(value int)

	// Topic: Specifying the Selected Segment

	// Sets the selection state of the specified segment.
	SetSelectedForSegment(selected bool, segment int)
	// Selects the segment with the specified tag.
	SelectSegmentWithTag(tag int) bool
	// Selects the next segment.
	MakeNextSegmentKey()
	// Selects the previous segment.
	MakePreviousSegmentKey()
	// The index of the selected segment of the control, or `-1` if no segment is selected.
	SelectedSegment() int
	SetSelectedSegment(value int)
	// Returns a Boolean value indicating whether the specified segment is selected,
	IsSelectedForSegment(segment int) bool

	// Topic: Specifying the Tracking Mode

	// The tracking mode used for the segments of the control.
	TrackingMode() NSSegmentSwitchTracking
	SetTrackingMode(value NSSegmentSwitchTracking)

	// Topic: Configuring Individual Segments

	// Sets the label for the specified segment.
	SetLabelForSegment(label string, segment int)
	// Returns the label of the specified segment.
	LabelForSegment(segment int) string
	// Sets the image for the specified segment.
	SetImageForSegment(image INSImage, segment int)
	// Returns the image associated with the specified segment.
	ImageForSegment(segment int) INSImage
	// Sets the image scaling mode for the specified segment.
	SetImageScalingForSegment(scaling NSImageScaling, segment int)
	// Returns the image scaling mode associated with the specified segment.
	ImageScalingForSegment(segment int) NSImageScaling
	// Sets the width of the specified segment.
	SetWidthForSegment(width float64, segment int)
	// Returns the width of the specified segment.
	WidthForSegment(segment int) float64
	// Sets the enabled state of the specified segment
	SetEnabledForSegment(enabled bool, segment int)
	// Returns a Boolean value indicating whether the specified segment is enabled.
	IsEnabledForSegment(segment int) bool
	// Sets the menu for the specified segment.
	SetMenuForSegment(menu INSMenu, segment int)
	// Returns the menu for the specified segment.
	MenuForSegment(segment int) INSMenu
	// Sets the tooltip for the specified segment.
	SetToolTipForSegment(toolTip string, segment int)
	// Returns the tooltip of the specified segment.
	ToolTipForSegment(segment int) string
	// Sets the tag for the specified segment.
	SetTagForSegment(tag int, segment int)
	// Returns the tag of the specified segment.
	TagForSegment(segment int) int

	// Topic: Drawing Custom Content

	// Draws the image and label of the segment in the specified view.
	DrawSegmentInFrameWithView(segment int, frame corefoundation.CGRect, controlView INSView)

	// Topic: Specifying Segment Visual Styles

	// Returns the interior background style for the specified segment.
	InteriorBackgroundStyleForSegment(segment int) NSBackgroundStyle
	// The visual style used to display the segmented control.
	SegmentStyle() NSSegmentStyle
	SetSegmentStyle(value NSSegmentStyle)
}

// Init initializes the instance.
func (s NSSegmentedCell) Init() NSSegmentedCell {
	rv := objc.Send[NSSegmentedCell](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSegmentedCell) Autorelease() NSSegmentedCell {
	rv := objc.Send[NSSegmentedCell](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSegmentedCell creates a new NSSegmentedCell instance.
func NewNSSegmentedCell() NSSegmentedCell {
	class := getNSSegmentedCellClass()
	rv := objc.Send[NSSegmentedCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSCell] object initialized with the specified image and set to
// have the cell’s default menu.
//
// image: The image to use for the cell. If this parameter is `nil`, no image is set.
//
// # Return Value
//
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
//
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewSegmentedCellImageCell(image INSImage) NSSegmentedCell {
	instance := getNSSegmentedCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSSegmentedCellFromID(rv)
}

// Returns an NSCell object initialized with the specified string and set to
// have the cell’s default menu.
//
// string: The initial string to use for the cell.
//
// # Return Value
//
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
//
// If no field editor (a shared [NSText] object) has been created for all
// [NSCell] objects, one is created.
//
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewSegmentedCellTextCell(string_ string) NSSegmentedCell {
	instance := getNSSegmentedCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSSegmentedCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewSegmentedCellWithCoder(coder foundation.INSCoder) NSSegmentedCell {
	instance := getNSSegmentedCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSegmentedCellFromID(rv)
}

// Sets the selection state of the specified segment.
//
// selected: true if you want to select the segment; otherwise, false.
//
// segment: The index of the segment whose selection state you want to set. This method
// raises an exception ([rangeException]) if the index is out of bounds.
//
// # Discussion
//
// If the control allows only a single selection, this method deselects any
// other selected segments.
//
// If the [TrackingMode] property of the segmented cell is set to
// [NSSegmentSwitchTrackingMomentary], then attempting to set the selected
// state of the segment will have no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setSelected(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetSelectedForSegment(selected bool, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSelected:forSegment:"), selected, segment)
}

// Selects the segment with the specified tag.
//
// tag: The tag associated with the desired segment.
//
// # Return Value
//
// true if the segment was selected successfully; otherwise, false.
//
// # Discussion
//
// Typically, you use Interface Builder to specify the tag for each segment.
// You may also set this value programmatically using the [SetTagForSegment]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/selectSegment(withTag:)
func (s NSSegmentedCell) SelectSegmentWithTag(tag int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}

// Selects the next segment.
//
// # Discussion
//
// The next segment is the one to the right of the currently selected segment.
// For the last segment, the selection wraps back to the beginning of the
// control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/makeNextSegmentKey()
func (s NSSegmentedCell) MakeNextSegmentKey() {
	objc.Send[objc.ID](s.ID, objc.Sel("makeNextSegmentKey"))
}

// Selects the previous segment.
//
// # Discussion
//
// The previous segment is the one to the left of the currently selected
// segment. For the first segment, the selection wraps around to the last
// segment of the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/makePreviousSegmentKey()
func (s NSSegmentedCell) MakePreviousSegmentKey() {
	objc.Send[objc.ID](s.ID, objc.Sel("makePreviousSegmentKey"))
}

// Returns a Boolean value indicating whether the specified segment is
// selected,
//
// segment: The index of the segment whose selection state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
//
// # Return Value
//
// true if the segment is selected; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/isSelected(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) IsSelectedForSegment(segment int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}

// Sets the label for the specified segment.
//
// label: The label you want to display in the segment. If the width of the string is
// greater than the width of the segment, the string’s text is truncated
// during drawing.
//
// segment: The index of the segment whose label you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setLabel(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetLabelForSegment(label string, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setLabel:forSegment:"), objc.String(label), segment)
}

// Returns the label of the specified segment.
//
// segment: The index of the segment whose label you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/label(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) LabelForSegment(segment int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("labelForSegment:"), segment)
	return foundation.NSStringFromID(rv).String()
}

// Sets the image for the specified segment.
//
// image: The image to apply to the segment or `nil` if you want to clear the
// existing image. Images are not scaled to fit inside a segment. If the image
// is larger than the available space, it is clipped.
//
// segment: The index of the segment whose image you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setImage(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetImageForSegment(image INSImage, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setImage:forSegment:"), image, segment)
}

// Returns the image associated with the specified segment.
//
// segment: The index of the segment whose image you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/image(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) ImageForSegment(segment int) INSImage {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("imageForSegment:"), segment)
	return NSImageFromID(rv)
}

// Sets the image scaling mode for the specified segment.
//
// scaling: The scaling mode to assign to the specified segment. For the possible
// values, see [NSSegmentedControl.Style].
//
// segment: The index of the segment whose image scaling mode you want to set. This
// method raises an exception ([rangeException]) if the index is out of
// bounds.
//
// # Discussion
//
// The image scaling mode for a segment affects how the image inside the
// corresponding cell is positioned and resized when the cell itself grows or
// shrinks. The image scaling mode does not itself cause the cell to change
// size in any way. If a cell does not contain an image, the scaling mode has
// no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setImageScaling(_:forSegment:)
//
// [NSSegmentedControl.Style]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetImageScalingForSegment(scaling NSImageScaling, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}

// Returns the image scaling mode associated with the specified segment.
//
// segment: The index of the segment whose image scaling mode you want to get. This
// method raises an exception ([rangeException]) if the index is out of
// bounds.
//
// # Return Value
//
// The scaling mode in use for the specified segment. For the possible values,
// see [NSSegmentedControl.Style]. If no value has been explicitly set,
// [NSImageScaleProportionallyDown] is returned.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/imageScaling(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [NSSegmentedControl.Style]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style
func (s NSSegmentedCell) ImageScalingForSegment(segment int) NSImageScaling {
	rv := objc.Send[NSImageScaling](s.ID, objc.Sel("imageScalingForSegment:"), segment)
	return NSImageScaling(rv)
}

// Sets the width of the specified segment.
//
// width: The width of the segment, measured in points. Specify the value `0` if you
// want the segment to be sized to fit the available space automatically.
//
// segment: The index of the segment whose width you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setWidth(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetWidthForSegment(width float64, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setWidth:forSegment:"), width, segment)
}

// Returns the width of the specified segment.
//
// segment: The index of the segment whose width you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/width(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) WidthForSegment(segment int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("widthForSegment:"), segment)
	return rv
}

// Sets the enabled state of the specified segment
//
// enabled: true to enable the segment; otherwise, false to disable it.
//
// segment: The index of the segment you want to enable or disable. This method raises
// an exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setEnabled(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetEnabledForSegment(enabled bool, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}

// Returns a Boolean value indicating whether the specified segment is
// enabled.
//
// segment: The index of the segment whose enabled state you want to get. This method
// raises an exception ([rangeException]) if the index is out of bounds.
//
// # Return Value
//
// true if the segment is enabled; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/isEnabled(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) IsEnabledForSegment(segment int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}

// Sets the menu for the specified segment.
//
// menu: The menu you want to add to the segment or `nil` to clear the current menu.
// This menu is displayed when the user clicks and holds the mouse button
// while the mouse is over the segment.
//
// segment: The index of the segment whose menu you want to set. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// # Discussion
//
// Adding a menu to a segment allows that segment to be used as a pop-up
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setMenu(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetMenuForSegment(menu INSMenu, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setMenu:forSegment:"), menu, segment)
}

// Returns the menu for the specified segment.
//
// segment: The index of the segment whose menu you want to get. This method raises an
// exception ([rangeException]) if the index is out of bounds.
//
// # Return Value
//
// The menu associated with the segment; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/menu(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) MenuForSegment(segment int) INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("menuForSegment:"), segment)
	return NSMenuFromID(rv)
}

// Sets the tooltip for the specified segment.
//
// toolTip: The text of the tooltip you want to display for the segment.
//
// segment: The index of the segment whose tooltip you want to set. This method raises
// an exception ([rangeException]) if the index is out of bounds.
//
// # Discussion
//
// Tooltips are currently not displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setToolTip(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetToolTipForSegment(toolTip string, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setToolTip:forSegment:"), objc.String(toolTip), segment)
}

// Returns the tooltip of the specified segment.
//
// segment: The index of the segment whose tooltip you want to get. This method raises
// an exception ([rangeException]) if the index is out of bounds.
//
// # Return Value
//
// The text of the tooltip.
//
// # Discussion
//
// Tooltips are currently not displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/toolTip(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) ToolTipForSegment(segment int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("toolTipForSegment:"), segment)
	return foundation.NSStringFromID(rv).String()
}

// Sets the tag for the specified segment.
//
// tag: The tag of the segment.
//
// segment: The index of the segment whose tool tag you want to set. This method raises
// an exception ([rangeException]) if the index is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setTag(_:forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SetTagForSegment(tag int, segment int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setTag:forSegment:"), tag, segment)
}

// Returns the tag of the specified segment.
//
// segment: The index of the segment whose tool tag you want to get. This method raises
// an exception ([rangeException]) if the index is out of bounds.
//
// # Return Value
//
// The tag of the segment.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/tag(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) TagForSegment(segment int) int {
	rv := objc.Send[int](s.ID, objc.Sel("tagForSegment:"), segment)
	return rv
}

// Draws the image and label of the segment in the specified view.
//
// segment: The index of the segment to draw. This method raises an exception
// ([rangeException]) if the index is out of bounds.
//
// frame: The rectangle in which to draw the segment’s image and label. This
// rectangle is specified in user space coordinates of the specified view.
//
// controlView: The view that contains the segment.
//
// # Discussion
//
// You can override this method to provide a custom appearance for segmented
// controls. You should not call this method directly. It is called for you
// automatically by the control when it needs to be redrawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/drawSegment(_:inFrame:with:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) DrawSegmentInFrameWithView(segment int, frame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawSegment:inFrame:withView:"), segment, frame, controlView)
}

// Returns the interior background style for the specified segment.
//
// segment: The index of the segment whose background style you want to get. This
// method raises an exception ([rangeException]) if the index is out of
// bounds..
//
// # Return Value
//
// The background style to use for specified segment. See
// [NSView.BackgroundStyle] for possible values.
//
// # Discussion
//
// The interior background style describes the surface drawn onto in
// [DrawInteriorWithFrameInView].
//
// This is both an override point and a useful method to call. In a custom
// segment cell with a custom bezel, you can override this method to describe
// the surface on a per-segment basis.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/interiorBackgroundStyle(forSegment:)
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [NSView.BackgroundStyle]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
func (s NSSegmentedCell) InteriorBackgroundStyleForSegment(segment int) NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](s.ID, objc.Sel("interiorBackgroundStyleForSegment:"), segment)
	return NSBackgroundStyle(rv)
}

// The number of segments in the segmented control.
//
// # Discussion
//
// This property contains the number of segments the segmented control should
// have. If this value is less than the number of segments currently in the
// control, segments are removed from the right of the control. Similarly, if
// the number is greater than the current number of segments, the new segments
// are added on the right. This value must be between `0` and `2049`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentCount
func (s NSSegmentedCell) SegmentCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("segmentCount"))
	return rv
}
func (s NSSegmentedCell) SetSegmentCount(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSegmentCount:"), value)
}

// The index of the selected segment of the control, or `-1` if no segment is
// selected.
//
// # Discussion
//
// This property contains the zero-based index of the segment. If the control
// allows multiple selections, the value of this property is the index of the
// most recently selected segment.
//
// If you specify an index that is out of bounds, an exception
// ([rangeException]) is raised.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/selectedSegment
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
func (s NSSegmentedCell) SelectedSegment() int {
	rv := objc.Send[int](s.ID, objc.Sel("selectedSegment"))
	return rv
}
func (s NSSegmentedCell) SetSelectedSegment(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectedSegment:"), value)
}

// The tracking mode used for the segments of the control.
//
// # Discussion
//
// Possible values for `trackingMode` are described in
// [NSSegmentedControl.SwitchTracking]. The default value is
// [NSSegmentSwitchTrackingSelectOne].
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/trackingMode
//
// [NSSegmentedControl.SwitchTracking]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
func (s NSSegmentedCell) TrackingMode() NSSegmentSwitchTracking {
	rv := objc.Send[NSSegmentSwitchTracking](s.ID, objc.Sel("trackingMode"))
	return NSSegmentSwitchTracking(rv)
}
func (s NSSegmentedCell) SetTrackingMode(value NSSegmentSwitchTracking) {
	objc.Send[struct{}](s.ID, objc.Sel("setTrackingMode:"), value)
}

// The visual style used to display the segmented control.
//
// # Discussion
//
// This property contains an [NSSegmentStyle] value that specifies the visual
// display used by the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentStyle
func (s NSSegmentedCell) SegmentStyle() NSSegmentStyle {
	rv := objc.Send[NSSegmentStyle](s.ID, objc.Sel("segmentStyle"))
	return NSSegmentStyle(rv)
}
func (s NSSegmentedCell) SetSegmentStyle(value NSSegmentStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setSegmentStyle:"), value)
}
