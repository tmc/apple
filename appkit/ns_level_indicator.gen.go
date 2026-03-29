// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSLevelIndicator] class.
var (
	_NSLevelIndicatorClass     NSLevelIndicatorClass
	_NSLevelIndicatorClassOnce sync.Once
)

func getNSLevelIndicatorClass() NSLevelIndicatorClass {
	_NSLevelIndicatorClassOnce.Do(func() {
		_NSLevelIndicatorClass = NSLevelIndicatorClass{class: objc.GetClass("NSLevelIndicator")}
	})
	return _NSLevelIndicatorClass
}

// GetNSLevelIndicatorClass returns the class object for NSLevelIndicator.
func GetNSLevelIndicatorClass() NSLevelIndicatorClass {
	return getNSLevelIndicatorClass()
}

type NSLevelIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLevelIndicatorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLevelIndicatorClass) Alloc() NSLevelIndicator {
	rv := objc.Send[NSLevelIndicator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A visual representation of a level or quantity, using discrete values.
//
// # Overview
// 
// A level indicator is similar to an [NSSlider] object, but provides a more
// customized visual feedback to the user. Unlike sliders, level indicators do
// not have a “knob” indicating the current setting, and they do not allow
// the user to adjust the current setting. You set the value of the level
// indicator programmatically. The supported indicator styles include:
// 
// - A capacity style level indicator. The continuous mode for this style is
// often used to indicate conditions such as how much data is on hard disk.
// The discrete mode is similar to audio level indicators in audio playback
// applications. You can specify both a warning value and a critical value
// that provides additional visual feedback to the user. - A ranking style
// level indicator. This is similar to the star ranking displays provided in
// iTunes and iPhoto. You can also specify your own ranking image. - A
// relevancy style level indicator. This style is used to display the
// relevancy of a search result, for example in Mail.
// 
// [NSLevelIndicator] uses an [NSLevelIndicatorCell] to implement much of the
// control’s functionality. [NSLevelIndicator] provides cover methods for
// most of the [NSLevelIndicatorCell] methods, which call the corresponding
// cell method.
//
// # Configuring the Range of Values
//
//   - [NSLevelIndicator.MinValue]: The receiver’s minimum value.
//   - [NSLevelIndicator.SetMinValue]
//   - [NSLevelIndicator.MaxValue]: The receiver’s maximum value.
//   - [NSLevelIndicator.SetMaxValue]
//   - [NSLevelIndicator.WarningValue]: The receiver’s warning value.
//   - [NSLevelIndicator.SetWarningValue]
//   - [NSLevelIndicator.CriticalValue]: The receiver’s critical value.
//   - [NSLevelIndicator.SetCriticalValue]
//
// # Managing Tick Marks and Style
//
//   - [NSLevelIndicator.TickMarkPosition]: Determines how the receiver’s tick marks are aligned with it.
//   - [NSLevelIndicator.SetTickMarkPosition]
//   - [NSLevelIndicator.NumberOfTickMarks]: The number of tick marks associated with the receiver.
//   - [NSLevelIndicator.SetNumberOfTickMarks]
//   - [NSLevelIndicator.NumberOfMajorTickMarks]: The number of major tick marks associated with the receiver.
//   - [NSLevelIndicator.SetNumberOfMajorTickMarks]
//   - [NSLevelIndicator.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0).
//   - [NSLevelIndicator.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0).
//   - [NSLevelIndicator.LevelIndicatorStyle]: The appearance of the indicator.
//   - [NSLevelIndicator.SetLevelIndicatorStyle]
//
// # Configuring the Drawing Attributes
//
//   - [NSLevelIndicator.RatingImage]
//   - [NSLevelIndicator.SetRatingImage]
//   - [NSLevelIndicator.DrawsTieredCapacityLevels]
//   - [NSLevelIndicator.SetDrawsTieredCapacityLevels]
//   - [NSLevelIndicator.FillColor]
//   - [NSLevelIndicator.SetFillColor]
//   - [NSLevelIndicator.WarningFillColor]
//   - [NSLevelIndicator.SetWarningFillColor]
//   - [NSLevelIndicator.CriticalFillColor]
//   - [NSLevelIndicator.SetCriticalFillColor]
//
// # Managing Placeholder Information
//
//   - [NSLevelIndicator.RatingPlaceholderImage]
//   - [NSLevelIndicator.SetRatingPlaceholderImage]
//   - [NSLevelIndicator.PlaceholderVisibility]
//   - [NSLevelIndicator.SetPlaceholderVisibility]
//
// # Controlling the Edit Behavior
//
//   - [NSLevelIndicator.Editable]
//   - [NSLevelIndicator.SetEditable]
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator
type NSLevelIndicator struct {
	NSControl
}

// NSLevelIndicatorFromID constructs a [NSLevelIndicator] from an objc.ID.
//
// A visual representation of a level or quantity, using discrete values.
func NSLevelIndicatorFromID(id objc.ID) NSLevelIndicator {
	return NSLevelIndicator{NSControl: NSControlFromID(id)}
}
// NOTE: NSLevelIndicator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLevelIndicator] class.
//
// # Configuring the Range of Values
//
//   - [INSLevelIndicator.MinValue]: The receiver’s minimum value.
//   - [INSLevelIndicator.SetMinValue]
//   - [INSLevelIndicator.MaxValue]: The receiver’s maximum value.
//   - [INSLevelIndicator.SetMaxValue]
//   - [INSLevelIndicator.WarningValue]: The receiver’s warning value.
//   - [INSLevelIndicator.SetWarningValue]
//   - [INSLevelIndicator.CriticalValue]: The receiver’s critical value.
//   - [INSLevelIndicator.SetCriticalValue]
//
// # Managing Tick Marks and Style
//
//   - [INSLevelIndicator.TickMarkPosition]: Determines how the receiver’s tick marks are aligned with it.
//   - [INSLevelIndicator.SetTickMarkPosition]
//   - [INSLevelIndicator.NumberOfTickMarks]: The number of tick marks associated with the receiver.
//   - [INSLevelIndicator.SetNumberOfTickMarks]
//   - [INSLevelIndicator.NumberOfMajorTickMarks]: The number of major tick marks associated with the receiver.
//   - [INSLevelIndicator.SetNumberOfMajorTickMarks]
//   - [INSLevelIndicator.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0).
//   - [INSLevelIndicator.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0).
//   - [INSLevelIndicator.LevelIndicatorStyle]: The appearance of the indicator.
//   - [INSLevelIndicator.SetLevelIndicatorStyle]
//
// # Configuring the Drawing Attributes
//
//   - [INSLevelIndicator.RatingImage]
//   - [INSLevelIndicator.SetRatingImage]
//   - [INSLevelIndicator.DrawsTieredCapacityLevels]
//   - [INSLevelIndicator.SetDrawsTieredCapacityLevels]
//   - [INSLevelIndicator.FillColor]
//   - [INSLevelIndicator.SetFillColor]
//   - [INSLevelIndicator.WarningFillColor]
//   - [INSLevelIndicator.SetWarningFillColor]
//   - [INSLevelIndicator.CriticalFillColor]
//   - [INSLevelIndicator.SetCriticalFillColor]
//
// # Managing Placeholder Information
//
//   - [INSLevelIndicator.RatingPlaceholderImage]
//   - [INSLevelIndicator.SetRatingPlaceholderImage]
//   - [INSLevelIndicator.PlaceholderVisibility]
//   - [INSLevelIndicator.SetPlaceholderVisibility]
//
// # Controlling the Edit Behavior
//
//   - [INSLevelIndicator.Editable]
//   - [INSLevelIndicator.SetEditable]
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator
type INSLevelIndicator interface {
	INSControl

	// Topic: Configuring the Range of Values

	// The receiver’s minimum value.
	MinValue() float64
	SetMinValue(value float64)
	// The receiver’s maximum value.
	MaxValue() float64
	SetMaxValue(value float64)
	// The receiver’s warning value.
	WarningValue() float64
	SetWarningValue(value float64)
	// The receiver’s critical value.
	CriticalValue() float64
	SetCriticalValue(value float64)

	// Topic: Managing Tick Marks and Style

	// Determines how the receiver’s tick marks are aligned with it.
	TickMarkPosition() NSTickMarkPosition
	SetTickMarkPosition(value NSTickMarkPosition)
	// The number of tick marks associated with the receiver.
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	// The number of major tick marks associated with the receiver.
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	// Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0).
	TickMarkValueAtIndex(index int) float64
	// Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0).
	RectOfTickMarkAtIndex(index int) corefoundation.CGRect
	// The appearance of the indicator.
	LevelIndicatorStyle() NSLevelIndicatorStyle
	SetLevelIndicatorStyle(value NSLevelIndicatorStyle)

	// Topic: Configuring the Drawing Attributes

	RatingImage() INSImage
	SetRatingImage(value INSImage)
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	FillColor() INSColor
	SetFillColor(value INSColor)
	WarningFillColor() INSColor
	SetWarningFillColor(value INSColor)
	CriticalFillColor() INSColor
	SetCriticalFillColor(value INSColor)

	// Topic: Managing Placeholder Information

	RatingPlaceholderImage() INSImage
	SetRatingPlaceholderImage(value INSImage)
	PlaceholderVisibility() NSLevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value NSLevelIndicatorPlaceholderVisibility)

	// Topic: Controlling the Edit Behavior

	Editable() bool
	SetEditable(value bool)
}

// Init initializes the instance.
func (l NSLevelIndicator) Init() NSLevelIndicator {
	rv := objc.Send[NSLevelIndicator](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLevelIndicator) Autorelease() NSLevelIndicator {
	rv := objc.Send[NSLevelIndicator](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLevelIndicator creates a new NSLevelIndicator instance.
func NewNSLevelIndicator() NSLevelIndicator {
	class := getNSLevelIndicatorClass()
	rv := objc.Send[NSLevelIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewLevelIndicatorWithCoder(coder foundation.INSCoder) NSLevelIndicator {
	instance := getNSLevelIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSLevelIndicatorFromID(rv)
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
func NewLevelIndicatorWithFrame(frameRect corefoundation.CGRect) NSLevelIndicator {
	instance := getNSLevelIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSLevelIndicatorFromID(rv)
}

// Returns the receiver’s value represented by the tick mark at the
// specified index (the minimum-value tick mark has an index of 0).
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkValue(at:)
func (l NSLevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}
// Returns the bounding rectangle of the tick mark identified by the specified
// index (the minimum-value tick mark is at index 0).
//
// # Discussion
// 
// If no tick mark is associated with `index`, the method raises an
// [NSRangeException].
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/rectOfTickMark(at:)
func (l NSLevelIndicator) RectOfTickMarkAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return corefoundation.CGRect(rv)
}

// The receiver’s minimum value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/minValue
func (l NSLevelIndicator) MinValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("minValue"))
	return rv
}
func (l NSLevelIndicator) SetMinValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setMinValue:"), value)
}
// The receiver’s maximum value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/maxValue
func (l NSLevelIndicator) MaxValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("maxValue"))
	return rv
}
func (l NSLevelIndicator) SetMaxValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setMaxValue:"), value)
}
// The receiver’s warning value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningValue
func (l NSLevelIndicator) WarningValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("warningValue"))
	return rv
}
func (l NSLevelIndicator) SetWarningValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setWarningValue:"), value)
}
// The receiver’s critical value.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalValue
func (l NSLevelIndicator) CriticalValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("criticalValue"))
	return rv
}
func (l NSLevelIndicator) SetCriticalValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setCriticalValue:"), value)
}
// Determines how the receiver’s tick marks are aligned with it.
//
// # Discussion
// 
// The default alignments are [NSTickMarkBelow] and [NSTickMarkLeft]. This
// property has no effect if no tick marks have been assigned (that is,
// [NumberOfTickMarks] returns 0).
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/tickMarkPosition
func (l NSLevelIndicator) TickMarkPosition() NSTickMarkPosition {
	rv := objc.Send[NSTickMarkPosition](l.ID, objc.Sel("tickMarkPosition"))
	return NSTickMarkPosition(rv)
}
func (l NSLevelIndicator) SetTickMarkPosition(value NSTickMarkPosition) {
	objc.Send[struct{}](l.ID, objc.Sel("setTickMarkPosition:"), value)
}
// The number of tick marks associated with the receiver.
//
// # Discussion
// 
// By default, this value is 0, and no tick marks appear. The number of tick
// marks assigned to a slider, along with the slider’s minimum and maximum
// values, determines the values associated with the tick marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfTickMarks
func (l NSLevelIndicator) NumberOfTickMarks() int {
	rv := objc.Send[int](l.ID, objc.Sel("numberOfTickMarks"))
	return rv
}
func (l NSLevelIndicator) SetNumberOfTickMarks(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberOfTickMarks:"), value)
}
// The number of major tick marks associated with the receiver.
//
// # Discussion
// 
// The number of major tick marks must be less than or equal to the number of
// tick marks returned by [NumberOfTickMarks]. For example, if the number of
// tick marks is 11 and you specify 3 major tick marks, the resulting level
// indicator will display 3 major tick marks alternating with 8 minor tick
// marks, as in the example shown in [NSLevelIndicator].
// 
// [media-1965752]
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/numberOfMajorTickMarks
func (l NSLevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}
func (l NSLevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}
// The appearance of the indicator.
//
// # Discussion
// 
// See [NSLevelIndicator.Style] for possible styles.
//
// [NSLevelIndicator.Style]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/levelIndicatorStyle
func (l NSLevelIndicator) LevelIndicatorStyle() NSLevelIndicatorStyle {
	rv := objc.Send[NSLevelIndicatorStyle](l.ID, objc.Sel("levelIndicatorStyle"))
	return NSLevelIndicatorStyle(rv)
}
func (l NSLevelIndicator) SetLevelIndicatorStyle(value NSLevelIndicatorStyle) {
	objc.Send[struct{}](l.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingImage
func (l NSLevelIndicator) RatingImage() INSImage {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("ratingImage"))
	return NSImageFromID(objc.ID(rv))
}
func (l NSLevelIndicator) SetRatingImage(value INSImage) {
	objc.Send[struct{}](l.ID, objc.Sel("setRatingImage:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/drawsTieredCapacityLevels
func (l NSLevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("drawsTieredCapacityLevels"))
	return rv
}
func (l NSLevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setDrawsTieredCapacityLevels:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/fillColor
func (l NSLevelIndicator) FillColor() INSColor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("fillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (l NSLevelIndicator) SetFillColor(value INSColor) {
	objc.Send[struct{}](l.ID, objc.Sel("setFillColor:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/warningFillColor
func (l NSLevelIndicator) WarningFillColor() INSColor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("warningFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (l NSLevelIndicator) SetWarningFillColor(value INSColor) {
	objc.Send[struct{}](l.ID, objc.Sel("setWarningFillColor:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/criticalFillColor
func (l NSLevelIndicator) CriticalFillColor() INSColor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("criticalFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (l NSLevelIndicator) SetCriticalFillColor(value INSColor) {
	objc.Send[struct{}](l.ID, objc.Sel("setCriticalFillColor:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/ratingPlaceholderImage
func (l NSLevelIndicator) RatingPlaceholderImage() INSImage {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("ratingPlaceholderImage"))
	return NSImageFromID(objc.ID(rv))
}
func (l NSLevelIndicator) SetRatingPlaceholderImage(value INSImage) {
	objc.Send[struct{}](l.ID, objc.Sel("setRatingPlaceholderImage:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/placeholderVisibility-swift.property
func (l NSLevelIndicator) PlaceholderVisibility() NSLevelIndicatorPlaceholderVisibility {
	rv := objc.Send[NSLevelIndicatorPlaceholderVisibility](l.ID, objc.Sel("placeholderVisibility"))
	return NSLevelIndicatorPlaceholderVisibility(rv)
}
func (l NSLevelIndicator) SetPlaceholderVisibility(value NSLevelIndicatorPlaceholderVisibility) {
	objc.Send[struct{}](l.ID, objc.Sel("setPlaceholderVisibility:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/isEditable
func (l NSLevelIndicator) Editable() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isEditable"))
	return rv
}
func (l NSLevelIndicator) SetEditable(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setEditable:"), value)
}

