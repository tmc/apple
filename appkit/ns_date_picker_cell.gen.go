// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSDatePickerCell] class.
var (
	_NSDatePickerCellClass     NSDatePickerCellClass
	_NSDatePickerCellClassOnce sync.Once
)

func getNSDatePickerCellClass() NSDatePickerCellClass {
	_NSDatePickerCellClassOnce.Do(func() {
		_NSDatePickerCellClass = NSDatePickerCellClass{class: objc.GetClass("NSDatePickerCell")}
	})
	return _NSDatePickerCellClass
}

// GetNSDatePickerCellClass returns the class object for NSDatePickerCell.
func GetNSDatePickerCellClass() NSDatePickerCellClass {
	return getNSDatePickerCellClass()
}

type NSDatePickerCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDatePickerCellClass) Alloc() NSDatePickerCell {
	rv := objc.Send[NSDatePickerCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that controls the behavior of a date picker, or of a single date
// picker cell in a matrix.
//
// # Configuring Appearance
//
//   - [NSDatePickerCell.BackgroundColor]: The cell’s background color.
//   - [NSDatePickerCell.SetBackgroundColor]
//   - [NSDatePickerCell.DrawsBackground]: A Boolean value indicating whether the cell draws its background.
//   - [NSDatePickerCell.SetDrawsBackground]
//   - [NSDatePickerCell.TextColor]: The cell’s text color.
//   - [NSDatePickerCell.SetTextColor]
//   - [NSDatePickerCell.DatePickerStyle]: The date picker style to use.
//   - [NSDatePickerCell.SetDatePickerStyle]
//   - [NSDatePickerCell.DatePickerElements]: A bitmask that indicates which visual elements are shown by the date picker.
//   - [NSDatePickerCell.SetDatePickerElements]
//
// # Range Mode
//
//   - [NSDatePickerCell.DatePickerMode]: The mode in use by the date picker.
//   - [NSDatePickerCell.SetDatePickerMode]
//
// # Object Values
//
//   - [NSDatePickerCell.DateValue]: The date currently specified in the picker.
//   - [NSDatePickerCell.SetDateValue]
//   - [NSDatePickerCell.TimeInterval]: The time interval that represents the date range.
//   - [NSDatePickerCell.SetTimeInterval]
//   - [NSDatePickerCell.Calendar]: The calendar used by the date picker.
//   - [NSDatePickerCell.SetCalendar]
//   - [NSDatePickerCell.Locale]: The locale used to display dates.
//   - [NSDatePickerCell.SetLocale]
//   - [NSDatePickerCell.TimeZone]: The time zone used to display time-related values.
//   - [NSDatePickerCell.SetTimeZone]
//
// # Date Range Constraints
//
//   - [NSDatePickerCell.MinDate]: The minimum date that the picker allows as input.
//   - [NSDatePickerCell.SetMinDate]
//   - [NSDatePickerCell.MaxDate]: The maximum date that the picker allows as input.
//   - [NSDatePickerCell.SetMaxDate]
//
// # Getting and Setting the Delegate
//
//   - [NSDatePickerCell.Delegate]: The delegate associated with the date picker.
//   - [NSDatePickerCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell
type NSDatePickerCell struct {
	NSActionCell
}

// NSDatePickerCellFromID constructs a [NSDatePickerCell] from an objc.ID.
//
// An object that controls the behavior of a date picker, or of a single date
// picker cell in a matrix.
func NSDatePickerCellFromID(id objc.ID) NSDatePickerCell {
	return NSDatePickerCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSDatePickerCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDatePickerCell] class.
//
// # Configuring Appearance
//
//   - [INSDatePickerCell.BackgroundColor]: The cell’s background color.
//   - [INSDatePickerCell.SetBackgroundColor]
//   - [INSDatePickerCell.DrawsBackground]: A Boolean value indicating whether the cell draws its background.
//   - [INSDatePickerCell.SetDrawsBackground]
//   - [INSDatePickerCell.TextColor]: The cell’s text color.
//   - [INSDatePickerCell.SetTextColor]
//   - [INSDatePickerCell.DatePickerStyle]: The date picker style to use.
//   - [INSDatePickerCell.SetDatePickerStyle]
//   - [INSDatePickerCell.DatePickerElements]: A bitmask that indicates which visual elements are shown by the date picker.
//   - [INSDatePickerCell.SetDatePickerElements]
//
// # Range Mode
//
//   - [INSDatePickerCell.DatePickerMode]: The mode in use by the date picker.
//   - [INSDatePickerCell.SetDatePickerMode]
//
// # Object Values
//
//   - [INSDatePickerCell.DateValue]: The date currently specified in the picker.
//   - [INSDatePickerCell.SetDateValue]
//   - [INSDatePickerCell.TimeInterval]: The time interval that represents the date range.
//   - [INSDatePickerCell.SetTimeInterval]
//   - [INSDatePickerCell.Calendar]: The calendar used by the date picker.
//   - [INSDatePickerCell.SetCalendar]
//   - [INSDatePickerCell.Locale]: The locale used to display dates.
//   - [INSDatePickerCell.SetLocale]
//   - [INSDatePickerCell.TimeZone]: The time zone used to display time-related values.
//   - [INSDatePickerCell.SetTimeZone]
//
// # Date Range Constraints
//
//   - [INSDatePickerCell.MinDate]: The minimum date that the picker allows as input.
//   - [INSDatePickerCell.SetMinDate]
//   - [INSDatePickerCell.MaxDate]: The maximum date that the picker allows as input.
//   - [INSDatePickerCell.SetMaxDate]
//
// # Getting and Setting the Delegate
//
//   - [INSDatePickerCell.Delegate]: The delegate associated with the date picker.
//   - [INSDatePickerCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell
type INSDatePickerCell interface {
	INSActionCell

	// Topic: Configuring Appearance

	// The cell’s background color.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean value indicating whether the cell draws its background.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// The cell’s text color.
	TextColor() INSColor
	SetTextColor(value INSColor)
	// The date picker style to use.
	DatePickerStyle() NSDatePickerStyle
	SetDatePickerStyle(value NSDatePickerStyle)
	// A bitmask that indicates which visual elements are shown by the date picker.
	DatePickerElements() NSDatePickerElementFlags
	SetDatePickerElements(value NSDatePickerElementFlags)

	// Topic: Range Mode

	// The mode in use by the date picker.
	DatePickerMode() NSDatePickerMode
	SetDatePickerMode(value NSDatePickerMode)

	// Topic: Object Values

	// The date currently specified in the picker.
	DateValue() foundation.INSDate
	SetDateValue(value foundation.INSDate)
	// The time interval that represents the date range.
	TimeInterval() float64
	SetTimeInterval(value float64)
	// The calendar used by the date picker.
	Calendar() foundation.NSCalendar
	SetCalendar(value foundation.NSCalendar)
	// The locale used to display dates.
	Locale() foundation.NSLocale
	SetLocale(value foundation.NSLocale)
	// The time zone used to display time-related values.
	TimeZone() foundation.NSTimeZone
	SetTimeZone(value foundation.NSTimeZone)

	// Topic: Date Range Constraints

	// The minimum date that the picker allows as input.
	MinDate() foundation.INSDate
	SetMinDate(value foundation.INSDate)
	// The maximum date that the picker allows as input.
	MaxDate() foundation.INSDate
	SetMaxDate(value foundation.INSDate)

	// Topic: Getting and Setting the Delegate

	// The delegate associated with the date picker.
	Delegate() NSDatePickerCellDelegate
	SetDelegate(value NSDatePickerCellDelegate)
}

// Init initializes the instance.
func (d NSDatePickerCell) Init() NSDatePickerCell {
	rv := objc.Send[NSDatePickerCell](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDatePickerCell) Autorelease() NSDatePickerCell {
	rv := objc.Send[NSDatePickerCell](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDatePickerCell creates a new NSDatePickerCell instance.
func NewNSDatePickerCell() NSDatePickerCell {
	class := getNSDatePickerCellClass()
	rv := objc.Send[NSDatePickerCell](objc.ID(class.class), objc.Sel("new"))
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
func NewDatePickerCellImageCell(image INSImage) NSDatePickerCell {
	instance := getNSDatePickerCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSDatePickerCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(textCell:)
func NewDatePickerCellTextCell(string_ string) NSDatePickerCell {
	instance := getNSDatePickerCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSDatePickerCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(coder:)
func NewDatePickerCellWithCoder(coder foundation.INSCoder) NSDatePickerCell {
	instance := getNSDatePickerCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDatePickerCellFromID(rv)
}

// The cell’s background color.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/backgroundColor
func (d NSDatePickerCell) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](d.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value indicating whether the cell draws its background.
//
// # Discussion
// 
// When the value of this property is [true], the cell draws its background.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/drawsBackground
func (d NSDatePickerCell) DrawsBackground() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("drawsBackground"))
	return rv
}
func (d NSDatePickerCell) SetDrawsBackground(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDrawsBackground:"), value)
}

// The cell’s text color.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/textColor
func (d NSDatePickerCell) TextColor() INSColor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetTextColor(value INSColor) {
	objc.Send[struct{}](d.ID, objc.Sel("setTextColor:"), value)
}

// The date picker style to use.
//
// # Discussion
// 
// For a list of possible values, see [NSDatePicker.Style].
//
// [NSDatePicker.Style]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerStyle
func (d NSDatePickerCell) DatePickerStyle() NSDatePickerStyle {
	rv := objc.Send[NSDatePickerStyle](d.ID, objc.Sel("datePickerStyle"))
	return NSDatePickerStyle(rv)
}
func (d NSDatePickerCell) SetDatePickerStyle(value NSDatePickerStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerStyle:"), value)
}

// A bitmask that indicates which visual elements are shown by the date
// picker.
//
// # Discussion
// 
// Elements not included in the bitmask are hidden from view. For a list of
// possible values, see [NSDatePicker.ElementFlags].
//
// [NSDatePicker.ElementFlags]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerElements
func (d NSDatePickerCell) DatePickerElements() NSDatePickerElementFlags {
	rv := objc.Send[NSDatePickerElementFlags](d.ID, objc.Sel("datePickerElements"))
	return NSDatePickerElementFlags(rv)
}
func (d NSDatePickerCell) SetDatePickerElements(value NSDatePickerElementFlags) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerElements:"), value)
}

// The mode in use by the date picker.
//
// # Discussion
// 
// For a list of possible values, see [NSDatePicker.Mode].
//
// [NSDatePicker.Mode]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerMode
func (d NSDatePickerCell) DatePickerMode() NSDatePickerMode {
	rv := objc.Send[NSDatePickerMode](d.ID, objc.Sel("datePickerMode"))
	return NSDatePickerMode(rv)
}
func (d NSDatePickerCell) SetDatePickerMode(value NSDatePickerMode) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerMode:"), value)
}

// The date currently specified in the picker.
//
// # Discussion
// 
// Assign a date to this property to set the starting value for the picker.
// For range-based dates, use the [TimeInterval] property to set the extent of
// the time range.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/dateValue
func (d NSDatePickerCell) DateValue() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateValue"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetDateValue(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateValue:"), value)
}

// The time interval that represents the date range.
//
// # Discussion
// 
// The date range begins at the date in the [DateValue] property. The value in
// this property applies only when the date picker is in the [NSRangeDateMode]
// mode.
//
// [NSRangeDateMode]: https://developer.apple.com/documentation/AppKit/NSRangeDateMode
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeInterval
func (d NSDatePickerCell) TimeInterval() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("timeInterval"))
	return rv
}
func (d NSDatePickerCell) SetTimeInterval(value float64) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeInterval:"), value)
}

// The calendar used by the date picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/calendar
func (d NSDatePickerCell) Calendar() foundation.NSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return foundation.NSCalendarFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetCalendar(value foundation.NSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}

// The locale used to display dates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/locale
func (d NSDatePickerCell) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetLocale(value foundation.NSLocale) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocale:"), value)
}

// The time zone used to display time-related values.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeZone
func (d NSDatePickerCell) TimeZone() foundation.NSTimeZone {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timeZone"))
	return foundation.NSTimeZoneFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetTimeZone(value foundation.NSTimeZone) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeZone:"), value)
}

// The minimum date that the picker allows as input.
//
// # Discussion
// 
// Set this property to `nil` if you want to allow any value for the minimum
// date.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/minDate
func (d NSDatePickerCell) MinDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("minDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetMinDate(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinDate:"), value)
}

// The maximum date that the picker allows as input.
//
// # Discussion
// 
// Set this property to `nil` if you want to allow any value for the maximum
// date.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate
func (d NSDatePickerCell) MaxDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("maxDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePickerCell) SetMaxDate(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaxDate:"), value)
}

// The delegate associated with the date picker.
//
// # Discussion
// 
// The delegate must conform to [NSDatePickerCellDelegate].
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/delegate
func (d NSDatePickerCell) Delegate() NSDatePickerCellDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return NSDatePickerCellDelegateObjectFromID(rv)
}
func (d NSDatePickerCell) SetDelegate(value NSDatePickerCellDelegate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

