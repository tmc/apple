// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ByteCountFormatter] class.
var (
	_ByteCountFormatterClass     ByteCountFormatterClass
	_ByteCountFormatterClassOnce sync.Once
)

func getByteCountFormatterClass() ByteCountFormatterClass {
	_ByteCountFormatterClassOnce.Do(func() {
		_ByteCountFormatterClass = ByteCountFormatterClass{class: objc.GetClass("NSByteCountFormatter")}
	})
	return _ByteCountFormatterClass
}

// GetByteCountFormatterClass returns the class object for NSByteCountFormatter.
func GetByteCountFormatterClass() ByteCountFormatterClass {
	return getByteCountFormatterClass()
}

type ByteCountFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that converts a byte count value into a localized description
// that is formatted with the appropriate byte modifier (KB, MB, GB and so
// on).
//
// # Overview
//
// # Creating Strings from Byte Count
//
//   - [ByteCountFormatter.StringFromByteCount]: Converts a byte count into a string without creating an [NSNumber] object.
//
// # Setting Formatting Styles
//
//   - [ByteCountFormatter.FormattingContext]: Specify the formatting context for the formatted string.
//   - [ByteCountFormatter.SetFormattingContext]
//   - [ByteCountFormatter.CountStyle]: Specify the number of bytes to be used for kilobytes.
//   - [ByteCountFormatter.SetCountStyle]
//   - [ByteCountFormatter.AllowsNonnumericFormatting]: Determines whether to allow more natural display of some values.
//   - [ByteCountFormatter.SetAllowsNonnumericFormatting]
//   - [ByteCountFormatter.IncludesActualByteCount]: Determines whether to include the number of bytes after the formatted string.
//   - [ByteCountFormatter.SetIncludesActualByteCount]
//   - [ByteCountFormatter.Adaptive]: Determines the display style of the size representation.
//   - [ByteCountFormatter.SetAdaptive]
//   - [ByteCountFormatter.AllowedUnits]: Specify the units that can be used in the output.
//   - [ByteCountFormatter.SetAllowedUnits]
//   - [ByteCountFormatter.IncludesCount]: Determines whether to include the count in the resulting formatted string.
//   - [ByteCountFormatter.SetIncludesCount]
//   - [ByteCountFormatter.IncludesUnit]: Determines whether to include the units in the resulting formatted string.
//   - [ByteCountFormatter.SetIncludesUnit]
//   - [ByteCountFormatter.ZeroPadsFractionDigits]: Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//   - [ByteCountFormatter.SetZeroPadsFractionDigits]
//
// # Instance Methods
//
//   - [ByteCountFormatter.StringFromMeasurement]
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter
type ByteCountFormatter struct {
	NSFormatter
}

// ByteCountFormatterFromID constructs a [ByteCountFormatter] from an objc.ID.
//
// A formatter that converts a byte count value into a localized description
// that is formatted with the appropriate byte modifier (KB, MB, GB and so
// on).
func ByteCountFormatterFromID(id objc.ID) ByteCountFormatter {
	return NSByteCountFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSByteCountFormatterFromID is an alias for [ByteCountFormatterFromID] for cross-framework compatibility.
func NSByteCountFormatterFromID(id objc.ID) ByteCountFormatter { return ByteCountFormatterFromID(id) }
// NOTE: ByteCountFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ByteCountFormatter] class.
//
// # Creating Strings from Byte Count
//
//   - [IByteCountFormatter.StringFromByteCount]: Converts a byte count into a string without creating an [NSNumber] object.
//
// # Setting Formatting Styles
//
//   - [IByteCountFormatter.FormattingContext]: Specify the formatting context for the formatted string.
//   - [IByteCountFormatter.SetFormattingContext]
//   - [IByteCountFormatter.CountStyle]: Specify the number of bytes to be used for kilobytes.
//   - [IByteCountFormatter.SetCountStyle]
//   - [IByteCountFormatter.AllowsNonnumericFormatting]: Determines whether to allow more natural display of some values.
//   - [IByteCountFormatter.SetAllowsNonnumericFormatting]
//   - [IByteCountFormatter.IncludesActualByteCount]: Determines whether to include the number of bytes after the formatted string.
//   - [IByteCountFormatter.SetIncludesActualByteCount]
//   - [IByteCountFormatter.Adaptive]: Determines the display style of the size representation.
//   - [IByteCountFormatter.SetAdaptive]
//   - [IByteCountFormatter.AllowedUnits]: Specify the units that can be used in the output.
//   - [IByteCountFormatter.SetAllowedUnits]
//   - [IByteCountFormatter.IncludesCount]: Determines whether to include the count in the resulting formatted string.
//   - [IByteCountFormatter.SetIncludesCount]
//   - [IByteCountFormatter.IncludesUnit]: Determines whether to include the units in the resulting formatted string.
//   - [IByteCountFormatter.SetIncludesUnit]
//   - [IByteCountFormatter.ZeroPadsFractionDigits]: Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//   - [IByteCountFormatter.SetZeroPadsFractionDigits]
//
// # Instance Methods
//
//   - [IByteCountFormatter.StringFromMeasurement]
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter
type IByteCountFormatter interface {
	INSFormatter

	// Topic: Creating Strings from Byte Count

	// Converts a byte count into a string without creating an [NSNumber] object.
	StringFromByteCount(byteCount int64) string

	// Topic: Setting Formatting Styles

	// Specify the formatting context for the formatted string.
	FormattingContext() NSFormattingContext
	SetFormattingContext(value NSFormattingContext)
	// Specify the number of bytes to be used for kilobytes.
	CountStyle() NSByteCountFormatterCountStyle
	SetCountStyle(value NSByteCountFormatterCountStyle)
	// Determines whether to allow more natural display of some values.
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	// Determines whether to include the number of bytes after the formatted string.
	IncludesActualByteCount() bool
	SetIncludesActualByteCount(value bool)
	// Determines the display style of the size representation.
	Adaptive() bool
	SetAdaptive(value bool)
	// Specify the units that can be used in the output.
	AllowedUnits() NSByteCountFormatterUnits
	SetAllowedUnits(value NSByteCountFormatterUnits)
	// Determines whether to include the count in the resulting formatted string.
	IncludesCount() bool
	SetIncludesCount(value bool)
	// Determines whether to include the units in the resulting formatted string.
	IncludesUnit() bool
	SetIncludesUnit(value bool)
	// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
	ZeroPadsFractionDigits() bool
	SetZeroPadsFractionDigits(value bool)

	// Topic: Instance Methods

	StringFromMeasurement(measurement INSMeasurement) string
}

// Init initializes the instance.
func (b ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b ByteCountFormatter) Autorelease() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewByteCountFormatter creates a new ByteCountFormatter instance.
func NewByteCountFormatter() ByteCountFormatter {
	class := getByteCountFormatterClass()
	rv := objc.Send[ByteCountFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewByteCountFormatterWithCoder(coder INSCoder) ByteCountFormatter {
	instance := getByteCountFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ByteCountFormatterFromID(rv)
}

// Converts a byte count into a string without creating an [NSNumber] object.
//
// byteCount: The byte count.
//
// # Return Value
// 
// A string containing the formatted `byteCount` value.
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:)
func (b ByteCountFormatter) StringFromByteCount(byteCount int64) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("stringFromByteCount:"), byteCount)
	return NSStringFromID(rv).String()
}
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:)
func (b ByteCountFormatter) StringFromMeasurement(measurement INSMeasurement) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("stringFromMeasurement:"), measurement)
	return NSStringFromID(rv).String()
}

// Converts a byte count into the specified string format without creating an
// [NSNumber] object.
//
// byteCount: The byte count.
//
// countStyle: The formatter style. See [ByteCountFormatter.CountStyle] for possible
// values.
// //
// [ByteCountFormatter.CountStyle]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum
//
// # Return Value
// 
// A string containing the formatted `byteCount` value.
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:countStyle:)
func (_ByteCountFormatterClass ByteCountFormatterClass) StringFromByteCountCountStyle(byteCount int64, countStyle NSByteCountFormatterCountStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_ByteCountFormatterClass.class), objc.Sel("stringFromByteCount:countStyle:"), byteCount, countStyle)
	return NSStringFromID(rv).String()
}
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:countStyle:)
func (_ByteCountFormatterClass ByteCountFormatterClass) StringFromMeasurementCountStyle(measurement INSMeasurement, countStyle NSByteCountFormatterCountStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_ByteCountFormatterClass.class), objc.Sel("stringFromMeasurement:countStyle:"), measurement, countStyle)
	return NSStringFromID(rv).String()
}

// Specify the formatting context for the formatted string.
//
// # Discussion
// 
// The default value is [NSFormattingContextUnknown]. See [NSFormatter] for
// possible values.
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/formattingContext
func (b ByteCountFormatter) FormattingContext() NSFormattingContext {
	rv := objc.Send[NSFormattingContext](b.ID, objc.Sel("formattingContext"))
	return NSFormattingContext(rv)
}
func (b ByteCountFormatter) SetFormattingContext(value NSFormattingContext) {
	objc.Send[struct{}](b.ID, objc.Sel("setFormattingContext:"), value)
}
// Specify the number of bytes to be used for kilobytes.
//
// # Discussion
// 
// The default setting is [ByteCountFormatterCountStyleFile], which is the
// system specific value for file and storage sizes.
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/countStyle-swift.property
func (b ByteCountFormatter) CountStyle() NSByteCountFormatterCountStyle {
	rv := objc.Send[NSByteCountFormatterCountStyle](b.ID, objc.Sel("countStyle"))
	return NSByteCountFormatterCountStyle(rv)
}
func (b ByteCountFormatter) SetCountStyle(value NSByteCountFormatterCountStyle) {
	objc.Send[struct{}](b.ID, objc.Sel("setCountStyle:"), value)
}
// Determines whether to allow more natural display of some values.
//
// # Discussion
// 
// Displays a more natural display of some values, such as zero, where it may
// be displayed as `Zero KB`, ignoring all other flags or options (with the
// exception of [ByteCountFormatterUseBytes], which would generate `Zero
// bytes`).The result is appropriate for standalone output.
// 
// Special handling of certain values such as zero is especially important in
// some languages, so it’s highly recommended that this property be left in
// its default state.
// 
// Default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowsNonnumericFormatting
func (b ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsNonnumericFormatting"))
	return rv
}
func (b ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsNonnumericFormatting:"), value)
}
// Determines whether to include the number of bytes after the formatted
// string.
//
// # Discussion
// 
// Setting this value to [true] causes the byte count to be displayed
// parenthetically (localized as appropriate), for instance `723 KB (722,842
// bytes)`. This will happen only if needed, that is, the first part is
// already not showing the exact byte count.
// 
// If [IncludesUnit] or [IncludesCount] are [false], then this setting has no
// effect.
// 
// Default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesActualByteCount
func (b ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("includesActualByteCount"))
	return rv
}
func (b ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setIncludesActualByteCount:"), value)
}
// Determines the display style of the size representation.
//
// # Discussion
// 
// The “adaptive” algorithm is platform specific and uses a different
// number of fraction digits based on the magnitude (in OS X v10.8: 0 fraction
// digits for bytes and KB; 1 fraction digits for MB; 2 for GB and above).
// Otherwise the result always tries to show at least three significant
// digits, introducing fraction digits as necessary.
// 
// Default is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/isAdaptive
func (b ByteCountFormatter) Adaptive() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isAdaptive"))
	return rv
}
func (b ByteCountFormatter) SetAdaptive(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAdaptive:"), value)
}
// Specify the units that can be used in the output.
//
// # Discussion
// 
// If the value is [ByteCountFormatterUseDefault], the formatter uses
// platform-appropriate settings; otherwise will only the specified units are
// used.
// 
// [ByteCountFormatter.Units] values can be combined using the C [OR] operator
// to specify complex formatting strings. The [ByteCountFormatterUseDefault]
// or [ByteCountFormatterUseAll] constants can be used with the C [AND] or the
// C [NOT] operators to create custom formats as well.
// 
// This is the default value if [ByteCountFormatterUseDefault].
//
// [ByteCountFormatter.Units]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowedUnits
func (b ByteCountFormatter) AllowedUnits() NSByteCountFormatterUnits {
	rv := objc.Send[NSByteCountFormatterUnits](b.ID, objc.Sel("allowedUnits"))
	return NSByteCountFormatterUnits(rv)
}
func (b ByteCountFormatter) SetAllowedUnits(value NSByteCountFormatterUnits) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowedUnits:"), value)
}
// Determines whether to include the count in the resulting formatted string.
//
// # Discussion
// 
// If set to [true] and [IncludesUnit] is set to [false], no unit is
// displayed. For example, a value of 723 KB is formatted as `723`.
// 
// You can get the set this property to [true] and the [IncludesUnit] to
// [true] individually to get both parts, separately. Note that putting them
// together yourself via string concatenation may be incorrect for some
// locales.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesCount
func (b ByteCountFormatter) IncludesCount() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("includesCount"))
	return rv
}
func (b ByteCountFormatter) SetIncludesCount(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setIncludesCount:"), value)
}
// Determines whether to include the units in the resulting formatted string.
//
// # Discussion
// 
// If set to [true] and [IncludesCount] is set to [false], no count is
// displayed. For example, a value of 723 KB is formatted as [KB].
// 
// You can get the set this property to [true] and the [IncludesCount] to
// [true] individually to get both parts, separately. Note that putting them
// together yourself via string concatenation may be incorrect for some
// locales.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesUnit
func (b ByteCountFormatter) IncludesUnit() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("includesUnit"))
	return rv
}
func (b ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setIncludesUnit:"), value)
}
// Determines whether to zero pad fraction digits so a consistent number of
// characters is displayed in a representation.
//
// # Discussion
// 
// Displaying values using zero pad fraction digits causes a consistent number
// of fraction digits are displayed, causing updating displays to remain more
// stable. For instance, if the [Adaptive] algorithm is used, this option
// formats 1.19 and 1.2 GB as `1.19 GB` and `1.20 GB`, respectively, while
// without the option the latter would be displayed as `1.2 GB`.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/zeroPadsFractionDigits
func (b ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("zeroPadsFractionDigits"))
	return rv
}
func (b ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setZeroPadsFractionDigits:"), value)
}

