// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSDatePicker] class.
var (
	_NSDatePickerClass     NSDatePickerClass
	_NSDatePickerClassOnce sync.Once
)

func getNSDatePickerClass() NSDatePickerClass {
	_NSDatePickerClassOnce.Do(func() {
		_NSDatePickerClass = NSDatePickerClass{class: objc.GetClass("NSDatePicker")}
	})
	return _NSDatePickerClass
}

// GetNSDatePickerClass returns the class object for NSDatePicker.
func GetNSDatePickerClass() NSDatePickerClass {
	return getNSDatePickerClass()
}

type NSDatePickerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDatePickerClass) Alloc() NSDatePicker {
	rv := objc.Send[NSDatePicker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A display of a calendar date with controls for editing the date value.
//
// # Overview
// 
// [NSDatePicker] uses an [NSDatePickerCell] to implement much of the
// control’s functionality. [NSDatePicker] provides cover methods for most
// of [NSDatePickerCell] methods, which invoke the corresponding cell method.
//
// # Configuring Date Pickers
//
//   - [NSDatePicker.Bezeled]: A Boolean value that indicates whether the date picker draws a bezeled border.
//   - [NSDatePicker.SetBezeled]
//   - [NSDatePicker.Bordered]: A Boolean value that indicates whether the date picker has a plain border.
//   - [NSDatePicker.SetBordered]
//   - [NSDatePicker.BackgroundColor]: The date picker’s background color.
//   - [NSDatePicker.SetBackgroundColor]
//   - [NSDatePicker.DrawsBackground]: A Boolean value that indicates whether the date picker draws the background.
//   - [NSDatePicker.SetDrawsBackground]
//   - [NSDatePicker.TextColor]: The date picker’s text color.
//   - [NSDatePicker.SetTextColor]
//   - [NSDatePicker.DatePickerStyle]: The date picker’s style.
//   - [NSDatePicker.SetDatePickerStyle]
//   - [NSDatePicker.PresentsCalendarOverlay]: A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
//   - [NSDatePicker.SetPresentsCalendarOverlay]
//   - [NSDatePicker.Delegate]: A delegate for the date picker’s cell
//   - [NSDatePicker.SetDelegate]
//   - [NSDatePicker.DatePickerElements]: A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//   - [NSDatePicker.SetDatePickerElements]
//
// # Controlling Date Picker Range and Mode
//
//   - [NSDatePicker.Calendar]: The calendar used by the date picker.
//   - [NSDatePicker.SetCalendar]
//   - [NSDatePicker.Locale]: The date picker’s locale.
//   - [NSDatePicker.SetLocale]
//   - [NSDatePicker.DatePickerMode]: The date picker’s mode.
//   - [NSDatePicker.SetDatePickerMode]
//   - [NSDatePicker.TimeZone]: The time zone for the date picker.
//   - [NSDatePicker.SetTimeZone]
//
// # Accessing Object Values
//
//   - [NSDatePicker.DateValue]: The date selected by the date picker.
//   - [NSDatePicker.SetDateValue]
//   - [NSDatePicker.TimeInterval]: The time interval selected by the date picker.
//   - [NSDatePicker.SetTimeInterval]
//
// # Constraining the Displayable/Selectable Range
//
//   - [NSDatePicker.MinDate]: The date picker’s minimum date value.
//   - [NSDatePicker.SetMinDate]
//   - [NSDatePicker.MaxDate]: The date picker’s maximum date value.
//   - [NSDatePicker.SetMaxDate]
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker
type NSDatePicker struct {
	NSControl
}

// NSDatePickerFromID constructs a [NSDatePicker] from an objc.ID.
//
// A display of a calendar date with controls for editing the date value.
func NSDatePickerFromID(id objc.ID) NSDatePicker {
	return NSDatePicker{NSControl: NSControlFromID(id)}
}
// NOTE: NSDatePicker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDatePicker] class.
//
// # Configuring Date Pickers
//
//   - [INSDatePicker.Bezeled]: A Boolean value that indicates whether the date picker draws a bezeled border.
//   - [INSDatePicker.SetBezeled]
//   - [INSDatePicker.Bordered]: A Boolean value that indicates whether the date picker has a plain border.
//   - [INSDatePicker.SetBordered]
//   - [INSDatePicker.BackgroundColor]: The date picker’s background color.
//   - [INSDatePicker.SetBackgroundColor]
//   - [INSDatePicker.DrawsBackground]: A Boolean value that indicates whether the date picker draws the background.
//   - [INSDatePicker.SetDrawsBackground]
//   - [INSDatePicker.TextColor]: The date picker’s text color.
//   - [INSDatePicker.SetTextColor]
//   - [INSDatePicker.DatePickerStyle]: The date picker’s style.
//   - [INSDatePicker.SetDatePickerStyle]
//   - [INSDatePicker.PresentsCalendarOverlay]: A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
//   - [INSDatePicker.SetPresentsCalendarOverlay]
//   - [INSDatePicker.Delegate]: A delegate for the date picker’s cell
//   - [INSDatePicker.SetDelegate]
//   - [INSDatePicker.DatePickerElements]: A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//   - [INSDatePicker.SetDatePickerElements]
//
// # Controlling Date Picker Range and Mode
//
//   - [INSDatePicker.Calendar]: The calendar used by the date picker.
//   - [INSDatePicker.SetCalendar]
//   - [INSDatePicker.Locale]: The date picker’s locale.
//   - [INSDatePicker.SetLocale]
//   - [INSDatePicker.DatePickerMode]: The date picker’s mode.
//   - [INSDatePicker.SetDatePickerMode]
//   - [INSDatePicker.TimeZone]: The time zone for the date picker.
//   - [INSDatePicker.SetTimeZone]
//
// # Accessing Object Values
//
//   - [INSDatePicker.DateValue]: The date selected by the date picker.
//   - [INSDatePicker.SetDateValue]
//   - [INSDatePicker.TimeInterval]: The time interval selected by the date picker.
//   - [INSDatePicker.SetTimeInterval]
//
// # Constraining the Displayable/Selectable Range
//
//   - [INSDatePicker.MinDate]: The date picker’s minimum date value.
//   - [INSDatePicker.SetMinDate]
//   - [INSDatePicker.MaxDate]: The date picker’s maximum date value.
//   - [INSDatePicker.SetMaxDate]
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker
type INSDatePicker interface {
	INSControl

	// Topic: Configuring Date Pickers

	// A Boolean value that indicates whether the date picker draws a bezeled border.
	Bezeled() bool
	SetBezeled(value bool)
	// A Boolean value that indicates whether the date picker has a plain border.
	Bordered() bool
	SetBordered(value bool)
	// The date picker’s background color.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean value that indicates whether the date picker draws the background.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// The date picker’s text color.
	TextColor() INSColor
	SetTextColor(value INSColor)
	// The date picker’s style.
	DatePickerStyle() NSDatePickerStyle
	SetDatePickerStyle(value NSDatePickerStyle)
	// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	// A delegate for the date picker’s cell
	Delegate() NSDatePickerCellDelegate
	SetDelegate(value NSDatePickerCellDelegate)
	// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
	DatePickerElements() NSDatePickerElementFlags
	SetDatePickerElements(value NSDatePickerElementFlags)

	// Topic: Controlling Date Picker Range and Mode

	// The calendar used by the date picker.
	Calendar() foundation.NSCalendar
	SetCalendar(value foundation.NSCalendar)
	// The date picker’s locale.
	Locale() foundation.NSLocale
	SetLocale(value foundation.NSLocale)
	// The date picker’s mode.
	DatePickerMode() NSDatePickerMode
	SetDatePickerMode(value NSDatePickerMode)
	// The time zone for the date picker.
	TimeZone() foundation.NSTimeZone
	SetTimeZone(value foundation.NSTimeZone)

	// Topic: Accessing Object Values

	// The date selected by the date picker.
	DateValue() foundation.INSDate
	SetDateValue(value foundation.INSDate)
	// The time interval selected by the date picker.
	TimeInterval() float64
	SetTimeInterval(value float64)

	// Topic: Constraining the Displayable/Selectable Range

	// The date picker’s minimum date value.
	MinDate() foundation.INSDate
	SetMinDate(value foundation.INSDate)
	// The date picker’s maximum date value.
	MaxDate() foundation.INSDate
	SetMaxDate(value foundation.INSDate)
}

// Init initializes the instance.
func (d NSDatePicker) Init() NSDatePicker {
	rv := objc.Send[NSDatePicker](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDatePicker) Autorelease() NSDatePicker {
	rv := objc.Send[NSDatePicker](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDatePicker creates a new NSDatePicker instance.
func NewNSDatePicker() NSDatePicker {
	class := getNSDatePickerClass()
	rv := objc.Send[NSDatePicker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewDatePickerWithCoder(coder foundation.INSCoder) NSDatePicker {
	instance := getNSDatePickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDatePickerFromID(rv)
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
func NewDatePickerWithFrame(frameRect corefoundation.CGRect) NSDatePicker {
	instance := getNSDatePickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSDatePickerFromID(rv)
}

// A Boolean value that indicates whether the date picker draws a bezeled
// border.
//
// # Discussion
// 
// This property contains [true] if the date picker has a bezeled boarder;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBezeled
func (d NSDatePicker) Bezeled() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isBezeled"))
	return rv
}
func (d NSDatePicker) SetBezeled(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setBezeled:"), value)
}

// A Boolean value that indicates whether the date picker has a plain border.
//
// # Discussion
// 
// This property contains [true] if the date picker has a plain boarder;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBordered
func (d NSDatePicker) Bordered() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isBordered"))
	return rv
}
func (d NSDatePicker) SetBordered(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setBordered:"), value)
}

// The date picker’s background color.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/backgroundColor
func (d NSDatePicker) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (d NSDatePicker) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](d.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value that indicates whether the date picker draws the
// background.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/drawsBackground
func (d NSDatePicker) DrawsBackground() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("drawsBackground"))
	return rv
}
func (d NSDatePicker) SetDrawsBackground(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDrawsBackground:"), value)
}

// The date picker’s text color.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/textColor
func (d NSDatePicker) TextColor() INSColor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (d NSDatePicker) SetTextColor(value INSColor) {
	objc.Send[struct{}](d.ID, objc.Sel("setTextColor:"), value)
}

// The date picker’s style.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerStyle
func (d NSDatePicker) DatePickerStyle() NSDatePickerStyle {
	rv := objc.Send[NSDatePickerStyle](d.ID, objc.Sel("datePickerStyle"))
	return NSDatePickerStyle(rv)
}
func (d NSDatePicker) SetDatePickerStyle(value NSDatePickerStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerStyle:"), value)
}

// A Boolean value that indicates whether to present a graphical calendar
// overlay when editing a calendar element within a text-field style date
// picker.
//
// # Discussion
// 
// The default value is [NO]. The overlay only appears for text-style date
// pickers when you select a calendar element. The overlay doesn’t appear
// when there are no calendar events or the value of this property is [YES].
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/presentsCalendarOverlay
func (d NSDatePicker) PresentsCalendarOverlay() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("presentsCalendarOverlay"))
	return rv
}
func (d NSDatePicker) SetPresentsCalendarOverlay(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setPresentsCalendarOverlay:"), value)
}

// A delegate for the date picker’s cell
//
// # Discussion
// 
// The date picker’s [NSDatePickerCell] instance handles all delegate
// methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/delegate
func (d NSDatePicker) Delegate() NSDatePickerCellDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return NSDatePickerCellDelegateObjectFromID(rv)
}
func (d NSDatePicker) SetDelegate(value NSDatePickerCellDelegate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

// A bitmask that indicates which visual elements of the date picker are
// currently shown, and which won’t be usable because they are hidden.
//
// # Discussion
// 
// See “Constants” in [NSDatePickerCell] for a description of the possible
// values.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerElements
func (d NSDatePicker) DatePickerElements() NSDatePickerElementFlags {
	rv := objc.Send[NSDatePickerElementFlags](d.ID, objc.Sel("datePickerElements"))
	return NSDatePickerElementFlags(rv)
}
func (d NSDatePicker) SetDatePickerElements(value NSDatePickerElementFlags) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerElements:"), value)
}

// The calendar used by the date picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/calendar
func (d NSDatePicker) Calendar() foundation.NSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return foundation.NSCalendarFromID(objc.ID(rv))
}
func (d NSDatePicker) SetCalendar(value foundation.NSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}

// The date picker’s locale.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/locale
func (d NSDatePicker) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}
func (d NSDatePicker) SetLocale(value foundation.NSLocale) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocale:"), value)
}

// The date picker’s mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerMode
func (d NSDatePicker) DatePickerMode() NSDatePickerMode {
	rv := objc.Send[NSDatePickerMode](d.ID, objc.Sel("datePickerMode"))
	return NSDatePickerMode(rv)
}
func (d NSDatePicker) SetDatePickerMode(value NSDatePickerMode) {
	objc.Send[struct{}](d.ID, objc.Sel("setDatePickerMode:"), value)
}

// The time zone for the date picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeZone
func (d NSDatePicker) TimeZone() foundation.NSTimeZone {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timeZone"))
	return foundation.NSTimeZoneFromID(objc.ID(rv))
}
func (d NSDatePicker) SetTimeZone(value foundation.NSTimeZone) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeZone:"), value)
}

// The date selected by the date picker.
//
// # Discussion
// 
// When selecting a date range, this property represents the time interval’s
// starting date.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/dateValue
func (d NSDatePicker) DateValue() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateValue"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePicker) SetDateValue(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateValue:"), value)
}

// The time interval selected by the date picker.
//
// # Discussion
// 
// The time interval that represents the receiver’s date range. The date
// range begins at the date returned by [DateValue]. This method returns 0
// when the receiver is not in the NSRangeDateMode mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeInterval
func (d NSDatePicker) TimeInterval() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("timeInterval"))
	return rv
}
func (d NSDatePicker) SetTimeInterval(value float64) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeInterval:"), value)
}

// The date picker’s minimum date value.
//
// # Discussion
// 
// This property represents the minimum value that the date picker allows as
// input. `nil` indicates no minimum date.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/minDate
func (d NSDatePicker) MinDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("minDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePicker) SetMinDate(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinDate:"), value)
}

// The date picker’s maximum date value.
//
// # Discussion
// 
// This property represents the maximum value that the date picker allows as
// input. `nil` indicates no maximum date.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePicker/maxDate
func (d NSDatePicker) MaxDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("maxDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDatePicker) SetMaxDate(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaxDate:"), value)
}

