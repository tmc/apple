// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSlider] class.
var (
	_NSSliderClass     NSSliderClass
	_NSSliderClassOnce sync.Once
)

func getNSSliderClass() NSSliderClass {
	_NSSliderClassOnce.Do(func() {
		_NSSliderClass = NSSliderClass{class: objc.GetClass("NSSlider")}
	})
	return _NSSliderClass
}

// GetNSSliderClass returns the class object for NSSlider.
func GetNSSliderClass() NSSliderClass {
	return getNSSliderClass()
}

type NSSliderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSliderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderClass) Alloc() NSSlider {
	rv := objc.Send[NSSlider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A display of a bar representing a continuous range of numerical values and
// a knob representing the currently selected value.
//
// # Overview
// 
// A slider is a UI element that displays a range of values in the app.
// Sliders can be vertical or horizontal bars or circular dials. An indicator,
// or knob, notes the current setting. The user can move the knob in the
// slider’s bar—or rotate the knob in a circular slider—to change the
// setting.
// 
// The [NSSlider] class uses the [NSSliderCell] class to implement its user
// interface.
//
// # Managing the slider’s appearance
//
//   - [NSSlider.SliderType]: The type of the slider, such as vertical or circular.
//   - [NSSlider.SetSliderType]
//   - [NSSlider.AltIncrementValue]: The amount by which the slider changes its value when the user Option-drags the slider knob.
//   - [NSSlider.SetAltIncrementValue]
//   - [NSSlider.KnobThickness]: The knob’s thickness, in pixels.
//   - [NSSlider.Vertical]: An integer indicating the orientation (horizontal or vertical) of the slider.
//   - [NSSlider.SetVertical]
//   - [NSSlider.TrackFillColor]: The color of the filled portion of the slider track, in appearances that support it.
//   - [NSSlider.SetTrackFillColor]
//   - [NSSlider.TintProminence]: The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
//   - [NSSlider.SetTintProminence]
//
// # Asking about the value limits
//
//   - [NSSlider.MaxValue]: The maximum value the slider can send to its target.
//   - [NSSlider.SetMaxValue]
//   - [NSSlider.MinValue]: The minimum value the slider can send to its target.
//   - [NSSlider.SetMinValue]
//
// # Managing tick marks
//
//   - [NSSlider.AllowsTickMarkValuesOnly]: A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//   - [NSSlider.SetAllowsTickMarkValuesOnly]
//   - [NSSlider.ClosestTickMarkValueToValue]: Returns the value of the tick mark closest to the specified value.
//   - [NSSlider.IndexOfTickMarkAtPoint]: Returns the index of the tick mark closest to the location of the slider represented by the given point.
//   - [NSSlider.NumberOfTickMarks]: The number of tick marks associated with the slider.
//   - [NSSlider.SetNumberOfTickMarks]
//   - [NSSlider.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark at the given index.
//   - [NSSlider.TickMarkPosition]: Determines where the slider’s tick marks are displayed.
//   - [NSSlider.SetTickMarkPosition]
//   - [NSSlider.TickMarkValueAtIndex]: Returns the slider’s value represented by the tick mark at the specified index.
//
// # Instance Properties
//
//   - [NSSlider.NeutralValue]: The value this slider will be filled from. This slider will be filled from its `neutralValue` to its current value. If `neutralValue` has not been explicitly set before, access to `neutralValue` will return `minValue`.
//   - [NSSlider.SetNeutralValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider
type NSSlider struct {
	NSControl
}

// NSSliderFromID constructs a [NSSlider] from an objc.ID.
//
// A display of a bar representing a continuous range of numerical values and
// a knob representing the currently selected value.
func NSSliderFromID(id objc.ID) NSSlider {
	return NSSlider{NSControl: NSControlFromID(id)}
}
// NOTE: NSSlider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSlider] class.
//
// # Managing the slider’s appearance
//
//   - [INSSlider.SliderType]: The type of the slider, such as vertical or circular.
//   - [INSSlider.SetSliderType]
//   - [INSSlider.AltIncrementValue]: The amount by which the slider changes its value when the user Option-drags the slider knob.
//   - [INSSlider.SetAltIncrementValue]
//   - [INSSlider.KnobThickness]: The knob’s thickness, in pixels.
//   - [INSSlider.Vertical]: An integer indicating the orientation (horizontal or vertical) of the slider.
//   - [INSSlider.SetVertical]
//   - [INSSlider.TrackFillColor]: The color of the filled portion of the slider track, in appearances that support it.
//   - [INSSlider.SetTrackFillColor]
//   - [INSSlider.TintProminence]: The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
//   - [INSSlider.SetTintProminence]
//
// # Asking about the value limits
//
//   - [INSSlider.MaxValue]: The maximum value the slider can send to its target.
//   - [INSSlider.SetMaxValue]
//   - [INSSlider.MinValue]: The minimum value the slider can send to its target.
//   - [INSSlider.SetMinValue]
//
// # Managing tick marks
//
//   - [INSSlider.AllowsTickMarkValuesOnly]: A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//   - [INSSlider.SetAllowsTickMarkValuesOnly]
//   - [INSSlider.ClosestTickMarkValueToValue]: Returns the value of the tick mark closest to the specified value.
//   - [INSSlider.IndexOfTickMarkAtPoint]: Returns the index of the tick mark closest to the location of the slider represented by the given point.
//   - [INSSlider.NumberOfTickMarks]: The number of tick marks associated with the slider.
//   - [INSSlider.SetNumberOfTickMarks]
//   - [INSSlider.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark at the given index.
//   - [INSSlider.TickMarkPosition]: Determines where the slider’s tick marks are displayed.
//   - [INSSlider.SetTickMarkPosition]
//   - [INSSlider.TickMarkValueAtIndex]: Returns the slider’s value represented by the tick mark at the specified index.
//
// # Instance Properties
//
//   - [INSSlider.NeutralValue]: The value this slider will be filled from. This slider will be filled from its `neutralValue` to its current value. If `neutralValue` has not been explicitly set before, access to `neutralValue` will return `minValue`.
//   - [INSSlider.SetNeutralValue]
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider
type INSSlider interface {
	INSControl
	NSAccessibilitySlider

	// Topic: Managing the slider’s appearance

	// The type of the slider, such as vertical or circular.
	SliderType() NSSliderType
	SetSliderType(value NSSliderType)
	// The amount by which the slider changes its value when the user Option-drags the slider knob.
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	// The knob’s thickness, in pixels.
	KnobThickness() float64
	// An integer indicating the orientation (horizontal or vertical) of the slider.
	Vertical() bool
	SetVertical(value bool)
	// The color of the filled portion of the slider track, in appearances that support it.
	TrackFillColor() INSColor
	SetTrackFillColor(value INSColor)
	// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See [NSTintProminence](<doc://com.apple.appkit/documentation/AppKit/NSTintProminence>) for a list of possible values.
	TintProminence() NSTintProminence
	SetTintProminence(value NSTintProminence)

	// Topic: Asking about the value limits

	// The maximum value the slider can send to its target.
	MaxValue() float64
	SetMaxValue(value float64)
	// The minimum value the slider can send to its target.
	MinValue() float64
	SetMinValue(value float64)

	// Topic: Managing tick marks

	// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	// Returns the value of the tick mark closest to the specified value.
	ClosestTickMarkValueToValue(value float64) float64
	// Returns the index of the tick mark closest to the location of the slider represented by the given point.
	IndexOfTickMarkAtPoint(point corefoundation.CGPoint) int
	// The number of tick marks associated with the slider.
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	// Returns the bounding rectangle of the tick mark at the given index.
	RectOfTickMarkAtIndex(index int) corefoundation.CGRect
	// Determines where the slider’s tick marks are displayed.
	TickMarkPosition() NSTickMarkPosition
	SetTickMarkPosition(value NSTickMarkPosition)
	// Returns the slider’s value represented by the tick mark at the specified index.
	TickMarkValueAtIndex(index int) float64

	// Topic: Instance Properties

	// The value this slider will be filled from. This slider will be filled from its `neutralValue` to its current value. If `neutralValue` has not been explicitly set before, access to `neutralValue` will return `minValue`.
	NeutralValue() float64
	SetNeutralValue(value float64)
}

// Init initializes the instance.
func (s NSSlider) Init() NSSlider {
	rv := objc.Send[NSSlider](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSlider) Autorelease() NSSlider {
	rv := objc.Send[NSSlider](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSlider creates a new NSSlider instance.
func NewNSSlider() NSSlider {
	class := getNSSliderClass()
	rv := objc.Send[NSSlider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewSliderWithCoder(coder foundation.INSCoder) NSSlider {
	instance := getNSSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSliderFromID(rv)
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
func NewSliderWithFrame(frameRect corefoundation.CGRect) NSSlider {
	instance := getNSSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSliderFromID(rv)
}

// Creates a continuous horizontal slider whose values range from `0.0` to
// `1.0`.
//
// target: The target object that receives action messages from the control.
//
// action: The action message sent by the control.
//
// # Return Value
// 
// An initialized slider control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/init(target:action:)
func NewSliderWithTargetAction(target objectivec.IObject, action objc.SEL) NSSlider {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderClass().class), objc.Sel("sliderWithTarget:action:"), target, action)
	return NSSliderFromID(rv)
}

// Creates a continuous horizontal slider that represents values over the
// specified range.
//
// value: The initial value displayed by the control.
//
// minValue: The minimum value that the control can represent.
//
// maxValue: The maximum value that the control can represent.
//
// target: The target object that receives action messages from the control.
//
// action: The action message sent by the control.
//
// # Return Value
// 
// An initialized slider control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/init(value:minValue:maxValue:target:action:)
func NewSliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objectivec.IObject, action objc.SEL) NSSlider {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderClass().class), objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return NSSliderFromID(rv)
}

// Returns the value of the tick mark closest to the specified value.
//
// value: The value for which to return the closest tick mark.
//
// # Return Value
// 
// The value of the tick mark closest to `aValue`.
//
// # Discussion
// 
// In its implementation of this method, the slider invokes the method of the
// same name of its [NSSliderCell] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/closestTickMarkValue(toValue:)
func (s NSSlider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}
// Returns the index of the tick mark closest to the location of the slider
// represented by the given point.
//
// point: The point representing the location for which to retrieve the tick mark.
//
// # Return Value
// 
// The index of the tick mark closest to the location specified by `point`. If
// `point` is not within the bounding rectangle (plus an extra pixel of space)
// of any tick mark, the method returns [NSNotFound].
//
// # Discussion
// 
// In its implementation of this method, the receiving [NSSlider] instance
// invokes the method of the same name of its [NSSliderCell] instance. This
// method invokes [RectOfTickMarkAtIndex] for each tick mark on the slider
// until it finds a tick mark containing the point.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/indexOfTickMark(at:)
func (s NSSlider) IndexOfTickMarkAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](s.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}
// Returns the bounding rectangle of the tick mark at the given index.
//
// index: The index of the tick mark for which to retrieve the bounds. The
// minimum-value tick mark is at index `0`.
//
// # Return Value
// 
// The bounding rectangle of the specified tick mark.
//
// # Discussion
// 
// If no tick mark is associated with `index`, the method raises
// [NSRangeException]. In its implementation of this method, the receiving
// [NSSlider] instance invokes the method of the same name of its
// [NSSliderCell] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/rectOfTickMark(at:)
func (s NSSlider) RectOfTickMarkAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return corefoundation.CGRect(rv)
}
// Returns the slider’s value represented by the tick mark at the specified
// index.
//
// index: The index of the tick mark for which to return the value. The minimum-value
// tick mark has an index of `0`.
//
// # Return Value
// 
// The value of the specified tick mark.
//
// # Discussion
// 
// In its implementation of this method, the receiving [NSSlider] instance
// invokes the method of the same name of its [NSSliderCell] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkValue(at:)
func (s NSSlider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}
// Decrements the slider’s value.
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
// # Discussion
// 
// This method must post an [valueChanged] notification after changing the
// slider’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformDecrement()
func (s NSSlider) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}
// Increments the slider’s value.
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
// # Discussion
// 
// This method must post an [valueChanged] notification after changing the
// slider’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformIncrement()
func (s NSSlider) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}
// Returns the slider’s value.
//
// # Return Value
// 
// The value for the slider.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityValue()
func (s NSSlider) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// The type of the slider, such as vertical or circular.
//
// # Discussion
// 
// See [NSSlider.SliderType] for possible values.
//
// [NSSlider.SliderType]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/sliderType-swift.property
func (s NSSlider) SliderType() NSSliderType {
	rv := objc.Send[NSSliderType](s.ID, objc.Sel("sliderType"))
	return NSSliderType(rv)
}
func (s NSSlider) SetSliderType(value NSSliderType) {
	objc.Send[struct{}](s.ID, objc.Sel("setSliderType:"), value)
}
// The amount by which the slider changes its value when the user Option-drags
// the slider knob.
//
// # Discussion
// 
// By default, the value of this property is `-1.0`, and the slider behaves no
// differently with the Option key down than with it up. The value of this
// property must fit the range of values the slider can represent—for
// example, if the slider has a minimum value of `5` and a maximum value of
// `10`, the value should be between `0` and `5`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/altIncrementValue
func (s NSSlider) AltIncrementValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("altIncrementValue"))
	return rv
}
func (s NSSlider) SetAltIncrementValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAltIncrementValue:"), value)
}
// The knob’s thickness, in pixels.
//
// # Discussion
// 
// The thickness is defined to be the extent of the knob along the long
// dimension of the bar. In a vertical slider, a knob’s thickness is its
// height; in a horizontal slider, a knob’s thickness is its width.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/knobThickness
func (s NSSlider) KnobThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("knobThickness"))
	return rv
}
// An integer indicating the orientation (horizontal or vertical) of the
// slider.
//
// # Discussion
// 
// The value of this property is `1` if the slider is vertical, `0` if it’s
// horizontal, and `-1` if the orientation is unknown (for example, if the
// slider hasn’t been displayed yet). A slider is defined as vertical if its
// height is greater than its width.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s NSSlider) Vertical() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isVertical"))
	return rv
}
func (s NSSlider) SetVertical(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setVertical:"), value)
}
// The color of the filled portion of the slider track, in appearances that
// support it.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s NSSlider) TrackFillColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("trackFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (s NSSlider) SetTrackFillColor(value INSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setTrackFillColor:"), value)
}
// The tint prominence of the slider. The automatic behavior for a regular
// slider tints its track fill, while a slider with tick marks is untinted.
// Setting the tint prominence will override this default behavior and choose
// an explicit track fill tint behavior. See [NSTintProminence] for a list of
// possible values.
//
// [NSTintProminence]: https://developer.apple.com/documentation/AppKit/NSTintProminence
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/tintProminence
func (s NSSlider) TintProminence() NSTintProminence {
	rv := objc.Send[NSTintProminence](s.ID, objc.Sel("tintProminence"))
	return NSTintProminence(rv)
}
func (s NSSlider) SetTintProminence(value NSTintProminence) {
	objc.Send[struct{}](s.ID, objc.Sel("setTintProminence:"), value)
}
// The maximum value the slider can send to its target.
//
// # Discussion
// 
// The slider’s maximum value. A horizontal slider sends the maximum value
// when the knob is all the way to the trailing end of the bar. A vertical
// slider sends the maximum value when the knob is at the top.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s NSSlider) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}
func (s NSSlider) SetMaxValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxValue:"), value)
}
// The minimum value the slider can send to its target.
//
// # Discussion
// 
// A horizontal slider sends the minimum value when the knob is all the way to
// the leading end of the bar. A vertical slider sends the minimum value when
// its knob is at the bottom.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/minValue
func (s NSSlider) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}
func (s NSSlider) SetMinValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinValue:"), value)
}
// A Boolean value that indicates whether the slider fixes its values to those
// values represented by its tick marks.
//
// # Discussion
// 
// This value is [true] if the slider fixes its values to the values
// represented by its tick marks; otherwise it’s [false]. For example, if a
// slider has a minimum value of `0`, a maximum value of `100`, and five
// markers, the allowable values are `0`, `25`, `50`, `75`, and `100`. When
// users move the slider’s knob, it jumps to the tick mark nearest the
// cursor when the mouse button is released. This method has no effect if the
// slider has no tick marks. In its implementation of this method, the
// receiving [NSSlider] instance invokes the method of the same name of its
// [NSSliderCell] instance.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/allowsTickMarkValuesOnly
func (s NSSlider) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}
func (s NSSlider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}
// The number of tick marks associated with the slider.
//
// # Discussion
// 
// This property includes tick marks assigned to the minimum and maximum
// values. In its implementation of this property, the receiving [NSSlider]
// instance invokes the method of the same name of its [NSSliderCell]
// instance. By default, this value is 0, and no tick marks appear. The number
// of tick marks assigned to a slider, along with the slider’s minimum and
// maximum values, determines the values associated with the tick marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/numberOfTickMarks
func (s NSSlider) NumberOfTickMarks() int {
	rv := objc.Send[int](s.ID, objc.Sel("numberOfTickMarks"))
	return rv
}
func (s NSSlider) SetNumberOfTickMarks(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setNumberOfTickMarks:"), value)
}
// Determines where the slider’s tick marks are displayed.
//
// # Discussion
// 
// For horizontal sliders, use [NSSlider.TickMarkPosition.below] and
// [NSSlider.TickMarkPosition.above]. For vertical sliders, use [leading], and
// [trailing]. The default positions are `below` for horizontal and `leading`
// for vertical. This property has no effect if [NumberOfTickMarks] is `0`, or
// if the slider is circular. In its implementation of this property, the
// receiving [NSSlider] instance invokes the method of the same name of its
// [NSSliderCell] instance.
//
// [NSSlider.TickMarkPosition.above]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/above
// [NSSlider.TickMarkPosition.below]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/below
// [leading]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/leading
// [trailing]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/trailing
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkPosition-swift.property
func (s NSSlider) TickMarkPosition() NSTickMarkPosition {
	rv := objc.Send[NSTickMarkPosition](s.ID, objc.Sel("tickMarkPosition"))
	return NSTickMarkPosition(rv)
}
func (s NSSlider) SetTickMarkPosition(value NSTickMarkPosition) {
	objc.Send[struct{}](s.ID, objc.Sel("setTickMarkPosition:"), value)
}
// The value this slider will be filled from. This slider will be filled from
// its `neutralValue` to its current value. If `neutralValue` has not been
// explicitly set before, access to `neutralValue` will return `minValue`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSlider/neutralValue
func (s NSSlider) NeutralValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("neutralValue"))
	return rv
}
func (s NSSlider) SetNeutralValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setNeutralValue:"), value)
}

			// Protocol methods for NSAccessibilitySlider
			
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
func (o NSSlider) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSSlider) AccessibilityParent() objectivec.IObject {
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
func (o NSSlider) AccessibilityIdentifier() string {
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
func (o NSSlider) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

