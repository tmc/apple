// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSSliderCell] class.
var (
	_NSSliderCellClass     NSSliderCellClass
	_NSSliderCellClassOnce sync.Once
)

func getNSSliderCellClass() NSSliderCellClass {
	_NSSliderCellClassOnce.Do(func() {
		_NSSliderCellClass = NSSliderCellClass{class: objc.GetClass("NSSliderCell")}
	})
	return _NSSliderCellClass
}

// GetNSSliderCellClass returns the class object for NSSliderCell.
func GetNSSliderCellClass() NSSliderCellClass {
	return getNSSliderCellClass()
}

type NSSliderCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSliderCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderCellClass) Alloc() NSSliderCell {
	rv := objc.Send[NSSliderCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The appearance and behavior of an [NSSlider] object.
//
// # Overview
// 
// You can customize an [NSSliderCell] to a certain degree, using its
// properties. If this doesn’t give you sufficient flexibility, you can
// create a subclass. In that subclass, you can override any of the following
// methods: [NSSliderCell.KnobRectFlipped], [NSSliderCell.DrawBarInsideFlipped], [NSSliderCell.DrawKnob], and
// [NSSliderCell.PrefersTrackingUntilMouseUp].
//
// # Managing Cell Behavior
//
//   - [NSSliderCell.AltIncrementValue]: The amount by which the slider changes its value when the user Option-drags the knob.
//   - [NSSliderCell.SetAltIncrementValue]
//   - [NSSliderCell.TrackRect]: The rectangle within which the cell tracks the pointer while the mouse button is down.
//
// # Managing the Slider Type
//
//   - [NSSliderCell.SliderType]: The slider type, either linear or circular.
//   - [NSSliderCell.SetSliderType]
//
// # Displaying the Cell
//
//   - [NSSliderCell.BarRectFlipped]: Returns the rectangle in which the bar is drawn.
//   - [NSSliderCell.DrawTickMarks]: Draws the slider’s tick marks.
//   - [NSSliderCell.KnobRectFlipped]: Returns the rectangle in which the slider knob is drawn.
//   - [NSSliderCell.DrawBarInsideFlipped]: Draws the slider’s bar—but not its bezel or knob—inside the specified rectangle.
//   - [NSSliderCell.DrawKnob]: Calculates the rectangle in which the knob should be drawn, then calls [drawKnob(_:)](<doc://com.apple.appkit/documentation/AppKit/NSSliderCell/drawKnob(_:)>) to actually draw the knob.
//   - [NSSliderCell.DrawKnobWithKnobRect]: Draws the slider knob in the given rectangle.
//
// # Managing Cell Appearance
//
//   - [NSSliderCell.KnobThickness]: The thickness of the slider knob, in pixels.
//   - [NSSliderCell.Vertical]: An integer indicating the orientation (vertical or horizontal) of the slider.
//   - [NSSliderCell.SetVertical]
//
// # Managing Value Limits
//
//   - [NSSliderCell.MaxValue]: The maximum value the slider can send to its target.
//   - [NSSliderCell.SetMaxValue]
//   - [NSSliderCell.MinValue]: The minimum value the slider can send to its target.
//   - [NSSliderCell.SetMinValue]
//
// # Managing Tick Marks
//
//   - [NSSliderCell.AllowsTickMarkValuesOnly]: A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
//   - [NSSliderCell.SetAllowsTickMarkValuesOnly]
//   - [NSSliderCell.ClosestTickMarkValueToValue]: Returns the value of the tick mark closest to the specified value.
//   - [NSSliderCell.IndexOfTickMarkAtPoint]: Returns the index of the tick mark closest to the location of the slider represented by the specified point.
//   - [NSSliderCell.NumberOfTickMarks]: The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
//   - [NSSliderCell.SetNumberOfTickMarks]
//   - [NSSliderCell.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark at the specified index.
//   - [NSSliderCell.TickMarkPosition]: The position of the tick marks relative to the receiver.
//   - [NSSliderCell.SetTickMarkPosition]
//   - [NSSliderCell.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell
type NSSliderCell struct {
	NSActionCell
}

// NSSliderCellFromID constructs a [NSSliderCell] from an objc.ID.
//
// The appearance and behavior of an [NSSlider] object.
func NSSliderCellFromID(id objc.ID) NSSliderCell {
	return NSSliderCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSSliderCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSliderCell] class.
//
// # Managing Cell Behavior
//
//   - [INSSliderCell.AltIncrementValue]: The amount by which the slider changes its value when the user Option-drags the knob.
//   - [INSSliderCell.SetAltIncrementValue]
//   - [INSSliderCell.TrackRect]: The rectangle within which the cell tracks the pointer while the mouse button is down.
//
// # Managing the Slider Type
//
//   - [INSSliderCell.SliderType]: The slider type, either linear or circular.
//   - [INSSliderCell.SetSliderType]
//
// # Displaying the Cell
//
//   - [INSSliderCell.BarRectFlipped]: Returns the rectangle in which the bar is drawn.
//   - [INSSliderCell.DrawTickMarks]: Draws the slider’s tick marks.
//   - [INSSliderCell.KnobRectFlipped]: Returns the rectangle in which the slider knob is drawn.
//   - [INSSliderCell.DrawBarInsideFlipped]: Draws the slider’s bar—but not its bezel or knob—inside the specified rectangle.
//   - [INSSliderCell.DrawKnob]: Calculates the rectangle in which the knob should be drawn, then calls [drawKnob(_:)](<doc://com.apple.appkit/documentation/AppKit/NSSliderCell/drawKnob(_:)>) to actually draw the knob.
//   - [INSSliderCell.DrawKnobWithKnobRect]: Draws the slider knob in the given rectangle.
//
// # Managing Cell Appearance
//
//   - [INSSliderCell.KnobThickness]: The thickness of the slider knob, in pixels.
//   - [INSSliderCell.Vertical]: An integer indicating the orientation (vertical or horizontal) of the slider.
//   - [INSSliderCell.SetVertical]
//
// # Managing Value Limits
//
//   - [INSSliderCell.MaxValue]: The maximum value the slider can send to its target.
//   - [INSSliderCell.SetMaxValue]
//   - [INSSliderCell.MinValue]: The minimum value the slider can send to its target.
//   - [INSSliderCell.SetMinValue]
//
// # Managing Tick Marks
//
//   - [INSSliderCell.AllowsTickMarkValuesOnly]: A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
//   - [INSSliderCell.SetAllowsTickMarkValuesOnly]
//   - [INSSliderCell.ClosestTickMarkValueToValue]: Returns the value of the tick mark closest to the specified value.
//   - [INSSliderCell.IndexOfTickMarkAtPoint]: Returns the index of the tick mark closest to the location of the slider represented by the specified point.
//   - [INSSliderCell.NumberOfTickMarks]: The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
//   - [INSSliderCell.SetNumberOfTickMarks]
//   - [INSSliderCell.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark at the specified index.
//   - [INSSliderCell.TickMarkPosition]: The position of the tick marks relative to the receiver.
//   - [INSSliderCell.SetTickMarkPosition]
//   - [INSSliderCell.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell
type INSSliderCell interface {
	INSActionCell

	// Topic: Managing Cell Behavior

	// The amount by which the slider changes its value when the user Option-drags the knob.
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	// The rectangle within which the cell tracks the pointer while the mouse button is down.
	TrackRect() corefoundation.CGRect

	// Topic: Managing the Slider Type

	// The slider type, either linear or circular.
	SliderType() NSSliderType
	SetSliderType(value NSSliderType)

	// Topic: Displaying the Cell

	// Returns the rectangle in which the bar is drawn.
	BarRectFlipped(flipped bool) corefoundation.CGRect
	// Draws the slider’s tick marks.
	DrawTickMarks()
	// Returns the rectangle in which the slider knob is drawn.
	KnobRectFlipped(flipped bool) corefoundation.CGRect
	// Draws the slider’s bar—but not its bezel or knob—inside the specified rectangle.
	DrawBarInsideFlipped(rect corefoundation.CGRect, flipped bool)
	// Calculates the rectangle in which the knob should be drawn, then calls [drawKnob(_:)](<doc://com.apple.appkit/documentation/AppKit/NSSliderCell/drawKnob(_:)>) to actually draw the knob.
	DrawKnob()
	// Draws the slider knob in the given rectangle.
	DrawKnobWithKnobRect(knobRect corefoundation.CGRect)

	// Topic: Managing Cell Appearance

	// The thickness of the slider knob, in pixels.
	KnobThickness() float64
	// An integer indicating the orientation (vertical or horizontal) of the slider.
	Vertical() bool
	SetVertical(value bool)

	// Topic: Managing Value Limits

	// The maximum value the slider can send to its target.
	MaxValue() float64
	SetMaxValue(value float64)
	// The minimum value the slider can send to its target.
	MinValue() float64
	SetMinValue(value float64)

	// Topic: Managing Tick Marks

	// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	// Returns the value of the tick mark closest to the specified value.
	ClosestTickMarkValueToValue(value float64) float64
	// Returns the index of the tick mark closest to the location of the slider represented by the specified point.
	IndexOfTickMarkAtPoint(point corefoundation.CGPoint) int
	// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	// Returns the bounding rectangle of the tick mark at the specified index.
	RectOfTickMarkAtIndex(index int) corefoundation.CGRect
	// The position of the tick marks relative to the receiver.
	TickMarkPosition() NSTickMarkPosition
	SetTickMarkPosition(value NSTickMarkPosition)
	// Returns the receiver’s value represented by the tick mark at the specified index.
	TickMarkValueAtIndex(index int) float64
}

// Init initializes the instance.
func (s NSSliderCell) Init() NSSliderCell {
	rv := objc.Send[NSSliderCell](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderCell) Autorelease() NSSliderCell {
	rv := objc.Send[NSSliderCell](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderCell creates a new NSSliderCell instance.
func NewNSSliderCell() NSSliderCell {
	class := getNSSliderCellClass()
	rv := objc.Send[NSSliderCell](objc.ID(class.class), objc.Sel("new"))
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
func NewSliderCellImageCell(image INSImage) NSSliderCell {
	instance := getNSSliderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSSliderCellFromID(rv)
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
func NewSliderCellTextCell(string_ string) NSSliderCell {
	instance := getNSSliderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSSliderCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewSliderCellWithCoder(coder foundation.INSCoder) NSSliderCell {
	instance := getNSSliderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSliderCellFromID(rv)
}

// Returns the rectangle in which the bar is drawn.
//
// flipped: [true] if the coordinate system of the associated [NSSlider] or [NSMatrix]
// is flipped; otherwise [false]. You can determine whether this is the case
// by sending the [NSView] message [isFlipped] message to the [NSMatrix] or
// [NSSlider].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [isFlipped]: https://developer.apple.com/documentation/AppKit/NSView/isFlipped
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The rectangle in which the bar is drawn, specified in the coordinate system
// of the [NSSlider] or [NSMatrix] with which the receiver is associated. The
// bar doesn’t include the slider’s bezel or knob.
//
// # Discussion
// 
// You can override this method if custom bar artwork requires specific
// dimensions.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/barRect(flipped:)
func (s NSSliderCell) BarRectFlipped(flipped bool) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("barRectFlipped:"), flipped)
	return corefoundation.CGRect(rv)
}
// Draws the slider’s tick marks.
//
// # Discussion
// 
// You should not call this method explicitly. It’s included so you can
// override it in a subclass and draw custom tick marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawTickMarks()
func (s NSSliderCell) DrawTickMarks() {
	objc.Send[objc.ID](s.ID, objc.Sel("drawTickMarks"))
}
// Returns the rectangle in which the slider knob is drawn.
//
// flipped: [true] if the coordinate system of the associated [NSSlider] or [NSMatrix]
// is flipped; otherwise [false]. You can determine whether this is the case
// by sending the [NSView] message [isFlipped] message to the [NSMatrix] or
// [NSSlider].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [isFlipped]: https://developer.apple.com/documentation/AppKit/NSView/isFlipped
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The rectangle in which the knob is drawn, specified in the coordinate
// system of the [NSSlider] or [NSMatrix] with which the receiver is
// associated.
//
// # Discussion
// 
// The knob rectangle depends on where in the slider the knob belongs—that
// is, it depends on the receiver’s minimum and maximum values and on the
// value the position of the knob will represent.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobRect(flipped:)
func (s NSSliderCell) KnobRectFlipped(flipped bool) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("knobRectFlipped:"), flipped)
	return corefoundation.CGRect(rv)
}
// Draws the slider’s bar—but not its bezel or knob—inside the specified
// rectangle.
//
// rect: The bounds of the slider’s bar, not of its interior rectangle.
//
// flipped: A Boolean value that indicates whether the cell’s control view—that is,
// the [NSSlider] or [NSMatrix] associated with the [NSSliderCell]—has a
// flipped coordinate system.
//
// # Discussion
// 
// You should not call this method explicitly. It’s included so you can
// override it in a subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawBar(inside:flipped:)
func (s NSSliderCell) DrawBarInsideFlipped(rect corefoundation.CGRect, flipped bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawBarInside:flipped:"), rect, flipped)
}
// Calculates the rectangle in which the knob should be drawn, then calls
// [DrawKnob] to actually draw the knob.
//
// # Discussion
// 
// Before this message is sent, a `lockFocus` method must be sent to the
// cell’s control view.
// 
// You might call this method if you override one of the display methods
// belonging to [NSControl] or [NSCell].
// 
// # Special Considerations
// 
// If you create a subclass of [NSSliderCell], don’t override this method.
// Override [DrawKnob] instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawKnob()
func (s NSSliderCell) DrawKnob() {
	objc.Send[objc.ID](s.ID, objc.Sel("drawKnob"))
}
// Draws the slider knob in the given rectangle.
//
// knobRect: The rectangle in which to draw the slider knob.
//
// # Discussion
// 
// Before this message is sent, a [lockFocus()] message must be sent to the
// cell’s control view.
// 
// You should not call this method explicitly. It’s included so you can
// override it in a subclass.
//
// [lockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/lockFocus()
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawKnob(_:)
func (s NSSliderCell) DrawKnobWithKnobRect(knobRect corefoundation.CGRect) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawKnob:"), knobRect)
}
// Returns the value of the tick mark closest to the specified value.
//
// value: The value for which to obtain the closest tick mark.
//
// # Return Value
// 
// The value of the closest tick mark.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/closestTickMarkValue(toValue:)
func (s NSSliderCell) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}
// Returns the index of the tick mark closest to the location of the slider
// represented by the specified point.
//
// point: The point representing the slider location.
//
// # Return Value
// 
// The index of the tick mark closest to the specified location.
//
// # Discussion
// 
// If `point` is not within the bounding rectangle (plus an extra pixel of
// space) of any tick mark, the method returns [NSNotFound]. This method calls
// [RectOfTickMarkAtIndex] for each tick mark on the slider until it finds a
// tick mark containing `point`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/indexOfTickMark(at:)
func (s NSSliderCell) IndexOfTickMarkAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](s.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}
// Returns the bounding rectangle of the tick mark at the specified index.
//
// index: The index of the tick mark for which to return the bounding rectangle. The
// minimum-value tick mark is at index 0.
//
// # Return Value
// 
// The bounding rectangle of the specified tick mark.
//
// # Discussion
// 
// If no tick mark is associated with `index`, the method raises
// [NSRangeException].
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/rectOfTickMark(at:)
func (s NSSliderCell) RectOfTickMarkAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return corefoundation.CGRect(rv)
}
// Returns the receiver’s value represented by the tick mark at the
// specified index.
//
// index: The index of the tick mark for which to retrieve the value. The
// minimum-value tick mark has an index of 0.
//
// # Return Value
// 
// The value represented by the specified tick mark.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/tickMarkValue(at:)
func (s NSSliderCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}

// The amount by which the slider changes its value when the user Option-drags
// the knob.
//
// # Discussion
// 
// By default, the value of this property is -1.0, and the slider behaves no
// differently with the Option key down than with it up. If you set this
// property, the number should correspond to the range of values the slider
// can represent—for example, if the slider has a minimum value of 5 and a
// maximum value of 10, the value should be between 0 and 5.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/altIncrementValue
func (s NSSliderCell) AltIncrementValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("altIncrementValue"))
	return rv
}
func (s NSSliderCell) SetAltIncrementValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAltIncrementValue:"), value)
}
// The rectangle within which the cell tracks the pointer while the mouse
// button is down.
//
// # Discussion
// 
// The tracking rectangle includes the slider bar, but not the bezel.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/trackRect
func (s NSSliderCell) TrackRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("trackRect"))
	return corefoundation.CGRect(rv)
}
// The slider type, either linear or circular.
//
// # Discussion
// 
// The value of this property is a constant indicating the type of the slider.
// Possible values are described in [NSSlider.SliderType].
// 
// When the value of this property is [NSCircularSlider], then you get a
// fixed-size circular slider. The minimum value ([MinValue]) is at the top,
// and the value increases clockwise around the dial. The maximum selectable
// value is just below [MaxValue]; for example, if [MaxValue] is 360, you can
// set the dial up to 359.999.
// 
// You can use the [NumberOfTickMarks] property to display tick marks, and you
// can use the [AllowsTickMarkValuesOnly] property to specify that values are
// limited to those values represented by tick marks. You can set this control
// to regular or small sizes; the mini size is not supported.
//
// [NSSlider.SliderType]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s NSSliderCell) SliderType() NSSliderType {
	rv := objc.Send[NSSliderType](s.ID, objc.Sel("sliderType"))
	return NSSliderType(rv)
}
func (s NSSliderCell) SetSliderType(value NSSliderType) {
	objc.Send[struct{}](s.ID, objc.Sel("setSliderType:"), value)
}
// The thickness of the slider knob, in pixels.
//
// # Discussion
// 
// The thickness is defined to be the extent of the knob along the long
// dimension of the bar. In a vertical slider, a knob’s thickness is its
// height; in a horizontal slider, its thickness is its width.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobThickness
func (s NSSliderCell) KnobThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("knobThickness"))
	return rv
}
// An integer indicating the orientation (vertical or horizontal) of the
// slider.
//
// # Discussion
// 
// The value of this property is 1 if the slider is vertical, 0 if it’s
// horizontal, and –1 if the orientation can’t be determined (for example,
// if the slider hasn’t been displayed yet). A slider is defined as vertical
// if its height is greater than its width.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/isVertical
func (s NSSliderCell) Vertical() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isVertical"))
	return rv
}
func (s NSSliderCell) SetVertical(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setVertical:"), value)
}
// The maximum value the slider can send to its target.
//
// # Discussion
// 
// A horizontal slider sends its maximum value when the knob is at the right
// end of the slider; a vertical slider sends it when the knob is at the top.
// The maximum selectable value for a circular slider is just below
// [MaxValue]; for example, if [MaxValue] is 360, you can set the dial up to
// 359.999.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/maxValue
func (s NSSliderCell) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}
func (s NSSliderCell) SetMaxValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxValue:"), value)
}
// The minimum value the slider can send to its target.
//
// # Discussion
// 
// A vertical slider sends this value when its knob is at the bottom; a
// horizontal slider sends it when its knob is all the way to the left; a
// circular slider sends it when its knob is at the top.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/minValue
func (s NSSliderCell) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}
func (s NSSliderCell) SetMinValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinValue:"), value)
}
// A Boolean value indicating whether the receiver fixes its values to those
// values represented by its tick marks.
//
// # Discussion
// 
// The value of this property is [true] if the slider’s values are limited
// to those values represented by tick marks; otherwise, [false]. For example,
// if you specify [true] for a slider that has a minimum value of 0, a maximum
// value of 100, and five markers, the allowable values are 0, 25, 50, 75, and
// 100. When users move the slider’s knob, it jumps to the tick mark nearest
// the pointer when the mouse button is released. Setting this property has no
// effect if the slider has no tick marks.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/allowsTickMarkValuesOnly
func (s NSSliderCell) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}
func (s NSSliderCell) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}
// The number of tick marks associated with the slider, including the tick
// marks assigned to the minimum and maximum values.
//
// # Discussion
// 
// By default, this value is 0, and no tick marks appear. The number of tick
// marks assigned to a slider, along with the slider’s minimum and maximum
// values, determines the values associated with the tick marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/numberOfTickMarks
func (s NSSliderCell) NumberOfTickMarks() int {
	rv := objc.Send[int](s.ID, objc.Sel("numberOfTickMarks"))
	return rv
}
func (s NSSliderCell) SetNumberOfTickMarks(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setNumberOfTickMarks:"), value)
}
// The position of the tick marks relative to the receiver.
//
// # Discussion
// 
// The value of this property is a constant indicating the position of the
// tick marks. Possible values are described in [NSSlider.TickMarkPosition].
// The default alignments are [NSTickMarkBelow] and [NSTickMarkLeft]. Setting
// this property has no effect if no tick marks have been assigned (that is,
// if [NumberOfTickMarks] is 0).
//
// [NSSlider.TickMarkPosition]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderCell/tickMarkPosition
func (s NSSliderCell) TickMarkPosition() NSTickMarkPosition {
	rv := objc.Send[NSTickMarkPosition](s.ID, objc.Sel("tickMarkPosition"))
	return NSTickMarkPosition(rv)
}
func (s NSSliderCell) SetTickMarkPosition(value NSTickMarkPosition) {
	objc.Send[struct{}](s.ID, objc.Sel("setTickMarkPosition:"), value)
}

