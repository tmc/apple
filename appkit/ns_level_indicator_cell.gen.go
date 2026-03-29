// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSLevelIndicatorCell] class.
var (
	_NSLevelIndicatorCellClass     NSLevelIndicatorCellClass
	_NSLevelIndicatorCellClassOnce sync.Once
)

func getNSLevelIndicatorCellClass() NSLevelIndicatorCellClass {
	_NSLevelIndicatorCellClassOnce.Do(func() {
		_NSLevelIndicatorCellClass = NSLevelIndicatorCellClass{class: objc.GetClass("NSLevelIndicatorCell")}
	})
	return _NSLevelIndicatorCellClass
}

// GetNSLevelIndicatorCellClass returns the class object for NSLevelIndicatorCell.
func GetNSLevelIndicatorCellClass() NSLevelIndicatorCellClass {
	return getNSLevelIndicatorCellClass()
}

type NSLevelIndicatorCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLevelIndicatorCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLevelIndicatorCellClass) Alloc() NSLevelIndicatorCell {
	rv := objc.Send[NSLevelIndicatorCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// [NSLevelIndicatorCell] is a subclass of [NSActionCell] that provides
// several level indicator display styles including: capacity, ranking and
// relevancy. The capacity style provides both continuous and discrete modes.
//
// # Initializing NSLevelIndicatorCell Objects
//
//   - [NSLevelIndicatorCell.InitWithLevelIndicatorStyle]: Initializes the receiver with the style specified by `levelIndicatorStyle`.
//
// # Configuring the Range of Values
//
//   - [NSLevelIndicatorCell.MinValue]: The minimum value of the control.
//   - [NSLevelIndicatorCell.SetMinValue]
//   - [NSLevelIndicatorCell.MaxValue]: The maximum value of the control.
//   - [NSLevelIndicatorCell.SetMaxValue]
//   - [NSLevelIndicatorCell.LevelIndicatorStyle]: The style of the level indicator control.
//   - [NSLevelIndicatorCell.SetLevelIndicatorStyle]
//   - [NSLevelIndicatorCell.WarningValue]: The warning value of the level indicator control.
//   - [NSLevelIndicatorCell.SetWarningValue]
//   - [NSLevelIndicatorCell.CriticalValue]: The critical value of the level indicator control.
//   - [NSLevelIndicatorCell.SetCriticalValue]
//
// # Managing Tick Marks
//
//   - [NSLevelIndicatorCell.TickMarkPosition]: The placement of tick marks on the level indicator control.
//   - [NSLevelIndicatorCell.SetTickMarkPosition]
//   - [NSLevelIndicatorCell.NumberOfTickMarks]: The number of tick marks displayed by the control.
//   - [NSLevelIndicatorCell.SetNumberOfTickMarks]
//   - [NSLevelIndicatorCell.NumberOfMajorTickMarks]: The number of major tick marks displayed by the control.
//   - [NSLevelIndicatorCell.SetNumberOfMajorTickMarks]
//   - [NSLevelIndicatorCell.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at index (the minimum-value tick mark has an index of 0).
//   - [NSLevelIndicatorCell.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark identified by `index` (the minimum-value tick mark is at index 0).
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell
type NSLevelIndicatorCell struct {
	NSActionCell
}

// NSLevelIndicatorCellFromID constructs a [NSLevelIndicatorCell] from an objc.ID.
//
// [NSLevelIndicatorCell] is a subclass of [NSActionCell] that provides
// several level indicator display styles including: capacity, ranking and
// relevancy. The capacity style provides both continuous and discrete modes.
func NSLevelIndicatorCellFromID(id objc.ID) NSLevelIndicatorCell {
	return NSLevelIndicatorCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSLevelIndicatorCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLevelIndicatorCell] class.
//
// # Initializing NSLevelIndicatorCell Objects
//
//   - [INSLevelIndicatorCell.InitWithLevelIndicatorStyle]: Initializes the receiver with the style specified by `levelIndicatorStyle`.
//
// # Configuring the Range of Values
//
//   - [INSLevelIndicatorCell.MinValue]: The minimum value of the control.
//   - [INSLevelIndicatorCell.SetMinValue]
//   - [INSLevelIndicatorCell.MaxValue]: The maximum value of the control.
//   - [INSLevelIndicatorCell.SetMaxValue]
//   - [INSLevelIndicatorCell.LevelIndicatorStyle]: The style of the level indicator control.
//   - [INSLevelIndicatorCell.SetLevelIndicatorStyle]
//   - [INSLevelIndicatorCell.WarningValue]: The warning value of the level indicator control.
//   - [INSLevelIndicatorCell.SetWarningValue]
//   - [INSLevelIndicatorCell.CriticalValue]: The critical value of the level indicator control.
//   - [INSLevelIndicatorCell.SetCriticalValue]
//
// # Managing Tick Marks
//
//   - [INSLevelIndicatorCell.TickMarkPosition]: The placement of tick marks on the level indicator control.
//   - [INSLevelIndicatorCell.SetTickMarkPosition]
//   - [INSLevelIndicatorCell.NumberOfTickMarks]: The number of tick marks displayed by the control.
//   - [INSLevelIndicatorCell.SetNumberOfTickMarks]
//   - [INSLevelIndicatorCell.NumberOfMajorTickMarks]: The number of major tick marks displayed by the control.
//   - [INSLevelIndicatorCell.SetNumberOfMajorTickMarks]
//   - [INSLevelIndicatorCell.TickMarkValueAtIndex]: Returns the receiver’s value represented by the tick mark at index (the minimum-value tick mark has an index of 0).
//   - [INSLevelIndicatorCell.RectOfTickMarkAtIndex]: Returns the bounding rectangle of the tick mark identified by `index` (the minimum-value tick mark is at index 0).
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell
type INSLevelIndicatorCell interface {
	INSActionCell

	// Topic: Initializing NSLevelIndicatorCell Objects

	// Initializes the receiver with the style specified by `levelIndicatorStyle`.
	InitWithLevelIndicatorStyle(levelIndicatorStyle NSLevelIndicatorStyle) NSLevelIndicatorCell

	// Topic: Configuring the Range of Values

	// The minimum value of the control.
	MinValue() float64
	SetMinValue(value float64)
	// The maximum value of the control.
	MaxValue() float64
	SetMaxValue(value float64)
	// The style of the level indicator control.
	LevelIndicatorStyle() NSLevelIndicatorStyle
	SetLevelIndicatorStyle(value NSLevelIndicatorStyle)
	// The warning value of the level indicator control.
	WarningValue() float64
	SetWarningValue(value float64)
	// The critical value of the level indicator control.
	CriticalValue() float64
	SetCriticalValue(value float64)

	// Topic: Managing Tick Marks

	// The placement of tick marks on the level indicator control.
	TickMarkPosition() NSTickMarkPosition
	SetTickMarkPosition(value NSTickMarkPosition)
	// The number of tick marks displayed by the control.
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	// The number of major tick marks displayed by the control.
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	// Returns the receiver’s value represented by the tick mark at index (the minimum-value tick mark has an index of 0).
	TickMarkValueAtIndex(index int) float64
	// Returns the bounding rectangle of the tick mark identified by `index` (the minimum-value tick mark is at index 0).
	RectOfTickMarkAtIndex(index int) corefoundation.CGRect
}

// Init initializes the instance.
func (l NSLevelIndicatorCell) Init() NSLevelIndicatorCell {
	rv := objc.Send[NSLevelIndicatorCell](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLevelIndicatorCell) Autorelease() NSLevelIndicatorCell {
	rv := objc.Send[NSLevelIndicatorCell](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLevelIndicatorCell creates a new NSLevelIndicatorCell instance.
func NewNSLevelIndicatorCell() NSLevelIndicatorCell {
	class := getNSLevelIndicatorCellClass()
	rv := objc.Send[NSLevelIndicatorCell](objc.ID(class.class), objc.Sel("new"))
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
func NewLevelIndicatorCellImageCell(image INSImage) NSLevelIndicatorCell {
	instance := getNSLevelIndicatorCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSLevelIndicatorCellFromID(rv)
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
func NewLevelIndicatorCellTextCell(string_ string) NSLevelIndicatorCell {
	instance := getNSLevelIndicatorCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSLevelIndicatorCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewLevelIndicatorCellWithCoder(coder foundation.INSCoder) NSLevelIndicatorCell {
	instance := getNSLevelIndicatorCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSLevelIndicatorCellFromID(rv)
}

// Initializes the receiver with the style specified by `levelIndicatorStyle`.
//
// # Discussion
// 
// The default value and minimum value are `0.0`. The default maximum value is
// dependent on `levelIndicatorStyle`. For continuous styles, the default
// maximum value is `100.0`. For discrete styles, the default maximum value is
// `5.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/init(levelIndicatorStyle:)
func NewLevelIndicatorCellWithLevelIndicatorStyle(levelIndicatorStyle NSLevelIndicatorStyle) NSLevelIndicatorCell {
	instance := getNSLevelIndicatorCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLevelIndicatorStyle:"), levelIndicatorStyle)
	return NSLevelIndicatorCellFromID(rv)
}

// Initializes the receiver with the style specified by `levelIndicatorStyle`.
//
// # Discussion
// 
// The default value and minimum value are `0.0`. The default maximum value is
// dependent on `levelIndicatorStyle`. For continuous styles, the default
// maximum value is `100.0`. For discrete styles, the default maximum value is
// `5.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/init(levelIndicatorStyle:)
func (l NSLevelIndicatorCell) InitWithLevelIndicatorStyle(levelIndicatorStyle NSLevelIndicatorStyle) NSLevelIndicatorCell {
	rv := objc.Send[NSLevelIndicatorCell](l.ID, objc.Sel("initWithLevelIndicatorStyle:"), levelIndicatorStyle)
	return rv
}
// Returns the receiver’s value represented by the tick mark at index (the
// minimum-value tick mark has an index of 0).
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/tickMarkValue(at:)
func (l NSLevelIndicatorCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}
// Returns the bounding rectangle of the tick mark identified by `index` (the
// minimum-value tick mark is at index 0).
//
// # Discussion
// 
// If no tick mark is associated with `index`, the method raises a
// [NSRangeException].
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/rectOfTickMark(at:)
func (l NSLevelIndicatorCell) RectOfTickMarkAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return corefoundation.CGRect(rv)
}

// The minimum value of the control.
//
// # Discussion
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/minValue
func (l NSLevelIndicatorCell) MinValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("minValue"))
	return rv
}
func (l NSLevelIndicatorCell) SetMinValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setMinValue:"), value)
}
// The maximum value of the control.
//
// # Discussion
// 
// The maximum value is dependent on the style of the control. For continuous
// styles, the default value of this property is `100.0`. For discrete styles,
// the default maximum value is `5.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/maxValue
func (l NSLevelIndicatorCell) MaxValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("maxValue"))
	return rv
}
func (l NSLevelIndicatorCell) SetMaxValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setMaxValue:"), value)
}
// The style of the level indicator control.
//
// # Discussion
// 
// The style determines the default minimum and maximum values of the control.
// For a list of possible styles, see [NSLevelIndicator.Style].
//
// [NSLevelIndicator.Style]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/levelIndicatorStyle
func (l NSLevelIndicatorCell) LevelIndicatorStyle() NSLevelIndicatorStyle {
	rv := objc.Send[NSLevelIndicatorStyle](l.ID, objc.Sel("levelIndicatorStyle"))
	return NSLevelIndicatorStyle(rv)
}
func (l NSLevelIndicatorCell) SetLevelIndicatorStyle(value NSLevelIndicatorStyle) {
	objc.Send[struct{}](l.ID, objc.Sel("setLevelIndicatorStyle:"), value)
}
// The warning value of the level indicator control.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/warningValue
func (l NSLevelIndicatorCell) WarningValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("warningValue"))
	return rv
}
func (l NSLevelIndicatorCell) SetWarningValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setWarningValue:"), value)
}
// The critical value of the level indicator control.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/criticalValue
func (l NSLevelIndicatorCell) CriticalValue() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("criticalValue"))
	return rv
}
func (l NSLevelIndicatorCell) SetCriticalValue(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setCriticalValue:"), value)
}
// The placement of tick marks on the level indicator control.
//
// # Discussion
// 
// Use this property to set the position where the control draws tick marks.
// Regardless of the value in this property, tick marks are not drawn if the
// value in the [NumberOfTickMarks] property is `0`.
// 
// The default value of this property is [NSTickMarkBelow], which also
// corresponds to the value [NSTickMarkLeft] for vertically oriented level
// indicators.
//
// [NSTickMarkBelow]: https://developer.apple.com/documentation/AppKit/NSTickMarkBelow
// [NSTickMarkLeft]: https://developer.apple.com/documentation/AppKit/NSTickMarkLeft
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/tickMarkPosition
func (l NSLevelIndicatorCell) TickMarkPosition() NSTickMarkPosition {
	rv := objc.Send[NSTickMarkPosition](l.ID, objc.Sel("tickMarkPosition"))
	return NSTickMarkPosition(rv)
}
func (l NSLevelIndicatorCell) SetTickMarkPosition(value NSTickMarkPosition) {
	objc.Send[struct{}](l.ID, objc.Sel("setTickMarkPosition:"), value)
}
// The number of tick marks displayed by the control.
//
// # Discussion
// 
// The value in this property represents all of the tick marks displayed by
// the control, including those positioned at the minimum and maximum values.
// The default value of this property is `0`, which results in no tick marks
// being displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfTickMarks
func (l NSLevelIndicatorCell) NumberOfTickMarks() int {
	rv := objc.Send[int](l.ID, objc.Sel("numberOfTickMarks"))
	return rv
}
func (l NSLevelIndicatorCell) SetNumberOfTickMarks(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberOfTickMarks:"), value)
}
// The number of major tick marks displayed by the control.
//
// # Discussion
// 
// The value in this property must be less than or equal to the value in the
// [NumberOfTickMarks] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell/numberOfMajorTickMarks
func (l NSLevelIndicatorCell) NumberOfMajorTickMarks() int {
	rv := objc.Send[int](l.ID, objc.Sel("numberOfMajorTickMarks"))
	return rv
}
func (l NSLevelIndicatorCell) SetNumberOfMajorTickMarks(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberOfMajorTickMarks:"), value)
}

