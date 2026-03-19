// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSStepperCell] class.
var (
	_NSStepperCellClass     NSStepperCellClass
	_NSStepperCellClassOnce sync.Once
)

func getNSStepperCellClass() NSStepperCellClass {
	_NSStepperCellClassOnce.Do(func() {
		_NSStepperCellClass = NSStepperCellClass{class: objc.GetClass("NSStepperCell")}
	})
	return _NSStepperCellClass
}

// GetNSStepperCellClass returns the class object for NSStepperCell.
func GetNSStepperCellClass() NSStepperCellClass {
	return getNSStepperCellClass()
}

type NSStepperCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStepperCellClass) Alloc() NSStepperCell {
	rv := objc.Send[NSStepperCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An [NSStepperCell] object controls the appearance and behavior of an
// [NSStepper] object.
//
// # Specifying value range
//
//   - [NSStepperCell.MaxValue]: The maximum value for the receiver.
//   - [NSStepperCell.SetMaxValue]
//   - [NSStepperCell.MinValue]: The minimum value for the receiver.
//   - [NSStepperCell.SetMinValue]
//   - [NSStepperCell.Increment]: The amount by which the receiver will change per increment or decrement.
//   - [NSStepperCell.SetIncrement]
//
// # Specifying how stepper cell responds
//
//   - [NSStepperCell.Autorepeat]: A Boolean value indicating how the receiver responds to mouse events.
//   - [NSStepperCell.SetAutorepeat]
//   - [NSStepperCell.ValueWraps]: A Boolean value indicating whether the receiver wraps around the minimum and maximum values.
//   - [NSStepperCell.SetValueWraps]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell
type NSStepperCell struct {
	NSActionCell
}

// NSStepperCellFromID constructs a [NSStepperCell] from an objc.ID.
//
// An [NSStepperCell] object controls the appearance and behavior of an
// [NSStepper] object.
func NSStepperCellFromID(id objc.ID) NSStepperCell {
	return NSStepperCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSStepperCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStepperCell] class.
//
// # Specifying value range
//
//   - [INSStepperCell.MaxValue]: The maximum value for the receiver.
//   - [INSStepperCell.SetMaxValue]
//   - [INSStepperCell.MinValue]: The minimum value for the receiver.
//   - [INSStepperCell.SetMinValue]
//   - [INSStepperCell.Increment]: The amount by which the receiver will change per increment or decrement.
//   - [INSStepperCell.SetIncrement]
//
// # Specifying how stepper cell responds
//
//   - [INSStepperCell.Autorepeat]: A Boolean value indicating how the receiver responds to mouse events.
//   - [INSStepperCell.SetAutorepeat]
//   - [INSStepperCell.ValueWraps]: A Boolean value indicating whether the receiver wraps around the minimum and maximum values.
//   - [INSStepperCell.SetValueWraps]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell
type INSStepperCell interface {
	INSActionCell

	// Topic: Specifying value range

	// The maximum value for the receiver.
	MaxValue() float64
	SetMaxValue(value float64)
	// The minimum value for the receiver.
	MinValue() float64
	SetMinValue(value float64)
	// The amount by which the receiver will change per increment or decrement.
	Increment() float64
	SetIncrement(value float64)

	// Topic: Specifying how stepper cell responds

	// A Boolean value indicating how the receiver responds to mouse events.
	Autorepeat() bool
	SetAutorepeat(value bool)
	// A Boolean value indicating whether the receiver wraps around the minimum and maximum values.
	ValueWraps() bool
	SetValueWraps(value bool)
}

// Init initializes the instance.
func (s NSStepperCell) Init() NSStepperCell {
	rv := objc.Send[NSStepperCell](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStepperCell) Autorelease() NSStepperCell {
	rv := objc.Send[NSStepperCell](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStepperCell creates a new NSStepperCell instance.
func NewNSStepperCell() NSStepperCell {
	class := getNSStepperCellClass()
	rv := objc.Send[NSStepperCell](objc.ID(class.class), objc.Sel("new"))
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
func NewStepperCellImageCell(image INSImage) NSStepperCell {
	instance := getNSStepperCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSStepperCellFromID(rv)
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
func NewStepperCellTextCell(string_ string) NSStepperCell {
	instance := getNSStepperCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSStepperCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewStepperCellWithCoder(coder foundation.INSCoder) NSStepperCell {
	instance := getNSStepperCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStepperCellFromID(rv)
}

// The maximum value for the receiver.
//
// # Discussion
// 
// The default is 59.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell/maxValue
func (s NSStepperCell) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}
func (s NSStepperCell) SetMaxValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxValue:"), value)
}
// The minimum value for the receiver.
//
// # Discussion
// 
// The default is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell/minValue
func (s NSStepperCell) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}
func (s NSStepperCell) SetMinValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinValue:"), value)
}
// The amount by which the receiver will change per increment or decrement.
//
// # Discussion
// 
// The default is 1.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell/increment
func (s NSStepperCell) Increment() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("increment"))
	return rv
}
func (s NSStepperCell) SetIncrement(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncrement:"), value)
}
// A Boolean value indicating how the receiver responds to mouse events.
//
// # Discussion
// 
// If [true], the first mouse down will do one increment (decrement), and,
// after a delay of 0.5 seconds, will increment (decrement) at a rate of ten
// times per second. If [false], the receiver will do one increment
// (decrement) on a mouse up. The default is [true]
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell/autorepeat
func (s NSStepperCell) Autorepeat() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("autorepeat"))
	return rv
}
func (s NSStepperCell) SetAutorepeat(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutorepeat:"), value)
}
// A Boolean value indicating whether the receiver wraps around the minimum
// and maximum values.
//
// # Discussion
// 
// [true] if, when incrementing or decrementing, the value wraps around to the
// minimum or maximum. [false] if the value stays pinned at the minimum or
// maximum. The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperCell/valueWraps
func (s NSStepperCell) ValueWraps() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("valueWraps"))
	return rv
}
func (s NSStepperCell) SetValueWraps(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setValueWraps:"), value)
}

