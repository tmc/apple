// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NumberFormatter] class.
var (
	_NumberFormatterClass     NumberFormatterClass
	_NumberFormatterClassOnce sync.Once
)

func getNumberFormatterClass() NumberFormatterClass {
	_NumberFormatterClassOnce.Do(func() {
		_NumberFormatterClass = NumberFormatterClass{class: objc.GetClass("NSNumberFormatter")}
	})
	return _NumberFormatterClass
}

// GetNumberFormatterClass returns the class object for NSNumberFormatter.
func GetNumberFormatterClass() NumberFormatterClass {
	return getNumberFormatterClass()
}

type NumberFormatterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NumberFormatterClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NumberFormatterClass) Alloc() NumberFormatter {
	rv := objc.Send[NumberFormatter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that converts between numeric values and their textual
// representations.
//
// # Overview
//
// Instances of [NSNumberFormatter] format the textual representation of cells
// that contain [NSNumber] objects and convert textual representations of
// numeric values into [NSNumber] objects. The representation encompasses
// integers, floats, and doubles; floats and doubles can be formatted to a
// specified decimal position. [NSNumberFormatter] objects can also impose
// ranges on the numeric values cells can accept.
//
// # Significant Digits and Fraction Digits
//
// The [NSNumberFormatter] class provides flexible options for displaying
// non-zero fractional parts of numbers.
//
// If you set the [UsesSignificantDigits] property to true, you can configure
// [NSNumberFormatter] to display significant digits using the
// [MinimumSignificantDigits] and [MaximumSignificantDigits] properties. If
// [UsesSignificantDigits] is false, these properties are ignored. See
// Configuring Significant Digits.
//
// Otherwise, you can configure the minimum and maximum number of integer and
// fraction digits, or the numbers before and after the decimal separator,
// respectively, using the [MinimumIntegerDigits], [MaximumIntegerDigits],
// [MinimumFractionDigits], and [MaximumFractionDigits] properties. See
// Configuring Integer and Fraction Digits.
//
// # Thread Safety
//
// On iOS 7 and later [NSNumberFormatter] is thread-safe.
//
// In macOS 10.9 and later [NSNumberFormatter] is thread-safe so long as you
// are using the modern behavior in a 64-bit app.
//
// On earlier versions of the operating system, or when using the legacy
// formatter behavior or running in 32-bit in macOS, [NSNumberFormatter] is
// not thread-safe, and you therefore must not mutate a number formatter
// simultaneously from multiple threads.
//
// # Configuring Formatter Behavior and Style
//
//   - [NumberFormatter.FormatterBehavior]: The formatter behavior of the receiver.
//   - [NumberFormatter.SetFormatterBehavior]
//   - [NumberFormatter.NumberStyle]: The number style used by the receiver.
//   - [NumberFormatter.SetNumberStyle]
//   - [NumberFormatter.GeneratesDecimalNumbers]: Determines whether the receiver creates instances of [NSDecimalNumber](<doc://com.apple.foundation/documentation/Foundation/NSDecimalNumber>) when it converts strings to number objects.
//   - [NumberFormatter.SetGeneratesDecimalNumbers]
//
// # Converting Between Numbers and Strings
//
//   - [NumberFormatter.GetObjectValueForStringRangeError]: Returns by reference a cell-content object after creating it from a range of characters in a given string.
//   - [NumberFormatter.NumberFromString]: Returns an [NSNumber](<doc://com.apple.foundation/documentation/Foundation/NSNumber>) object created by parsing a given string.
//   - [NumberFormatter.StringFromNumber]: Returns a string containing the formatted value of the provided number object.
//
// # Managing Localization of Numbers
//
//   - [NumberFormatter.LocalizesFormat]: Determines whether the dollar sign character (`$`), decimal separator character (`.`), and thousand separator character (`,`) are converted to appropriately localized characters as specified by the user’s localization preference.
//   - [NumberFormatter.SetLocalizesFormat]
//   - [NumberFormatter.Locale]: The locale of the receiver.
//   - [NumberFormatter.SetLocale]
//
// # Configuring Rounding Behavior
//
//   - [NumberFormatter.RoundingBehavior]: The rounding behavior used by the receiver.
//   - [NumberFormatter.SetRoundingBehavior]
//   - [NumberFormatter.RoundingIncrement]: The rounding increment used by the receiver.
//   - [NumberFormatter.SetRoundingIncrement]
//   - [NumberFormatter.RoundingMode]: The rounding mode used by the receiver.
//   - [NumberFormatter.SetRoundingMode]
//
// # Configuring Integer and Fraction Digits
//
//   - [NumberFormatter.MinimumIntegerDigits]: The minimum number of digits before the decimal separator.
//   - [NumberFormatter.SetMinimumIntegerDigits]
//   - [NumberFormatter.MaximumIntegerDigits]: The maximum number of digits before the decimal separator.
//   - [NumberFormatter.SetMaximumIntegerDigits]
//   - [NumberFormatter.MinimumFractionDigits]: The minimum number of digits after the decimal separator.
//   - [NumberFormatter.SetMinimumFractionDigits]
//   - [NumberFormatter.MaximumFractionDigits]: The maximum number of digits after the decimal separator.
//   - [NumberFormatter.SetMaximumFractionDigits]
//
// # Configuring Significant Digits
//
//   - [NumberFormatter.UsesSignificantDigits]: A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//   - [NumberFormatter.SetUsesSignificantDigits]
//   - [NumberFormatter.MinimumSignificantDigits]: The minimum number of significant digits for the number formatter.
//   - [NumberFormatter.SetMinimumSignificantDigits]
//   - [NumberFormatter.MaximumSignificantDigits]: The maximum number of significant digits for the number formatter.
//   - [NumberFormatter.SetMaximumSignificantDigits]
//
// # Configuring Numeric Formats
//
//   - [NumberFormatter.Format]: The receiver’s format.
//   - [NumberFormatter.SetFormat]
//   - [NumberFormatter.FormattingContext]: The capitalization formatting context used when formatting a number.
//   - [NumberFormatter.SetFormattingContext]
//   - [NumberFormatter.FormatWidth]: The format width used by the receiver.
//   - [NumberFormatter.SetFormatWidth]
//   - [NumberFormatter.NegativeFormat]: The format the receiver uses to display negative values.
//   - [NumberFormatter.SetNegativeFormat]
//   - [NumberFormatter.PositiveFormat]: The format the receiver uses to display positive values.
//   - [NumberFormatter.SetPositiveFormat]
//   - [NumberFormatter.Multiplier]: The multiplier of the receiver.
//   - [NumberFormatter.SetMultiplier]
//
// # Configuring Numeric Symbols
//
//   - [NumberFormatter.PercentSymbol]: The string used to represent a percent symbol.
//   - [NumberFormatter.SetPercentSymbol]
//   - [NumberFormatter.PerMillSymbol]: The string used to represent a per-mill (per-thousand) symbol.
//   - [NumberFormatter.SetPerMillSymbol]
//   - [NumberFormatter.MinusSign]: The string used to represent a minus sign.
//   - [NumberFormatter.SetMinusSign]
//   - [NumberFormatter.PlusSign]: The string used to represent a plus sign.
//   - [NumberFormatter.SetPlusSign]
//   - [NumberFormatter.ExponentSymbol]: The string used to represent an exponent symbol.
//   - [NumberFormatter.SetExponentSymbol]
//   - [NumberFormatter.ZeroSymbol]: The string used to represent a zero value.
//   - [NumberFormatter.SetZeroSymbol]
//   - [NumberFormatter.NilSymbol]: The string used to represent a `nil` value.
//   - [NumberFormatter.SetNilSymbol]
//   - [NumberFormatter.NotANumberSymbol]: The string used to represent a NaN (“not a number”) value.
//   - [NumberFormatter.SetNotANumberSymbol]
//   - [NumberFormatter.NegativeInfinitySymbol]: The string used to represent a negative infinity symbol.
//   - [NumberFormatter.SetNegativeInfinitySymbol]
//   - [NumberFormatter.PositiveInfinitySymbol]: The string used to represent a positive infinity symbol.
//   - [NumberFormatter.SetPositiveInfinitySymbol]
//
// # Configuring the Format of Currency
//
//   - [NumberFormatter.CurrencySymbol]: The string used by the receiver as a local currency symbol.
//   - [NumberFormatter.SetCurrencySymbol]
//   - [NumberFormatter.CurrencyCode]: The receiver’s currency code.
//   - [NumberFormatter.SetCurrencyCode]
//   - [NumberFormatter.InternationalCurrencySymbol]: The international currency symbol used by the receiver.
//   - [NumberFormatter.SetInternationalCurrencySymbol]
//   - [NumberFormatter.CurrencyGroupingSeparator]: The currency grouping separator for the receiver.
//   - [NumberFormatter.SetCurrencyGroupingSeparator]
//
// # Configuring Numeric Prefixes and Suffixes
//
//   - [NumberFormatter.PositivePrefix]: The string the receiver uses as the prefix for positive values.
//   - [NumberFormatter.SetPositivePrefix]
//   - [NumberFormatter.PositiveSuffix]: The string the receiver uses as the suffix for positive values.
//   - [NumberFormatter.SetPositiveSuffix]
//   - [NumberFormatter.NegativePrefix]: The string the receiver uses as a prefix for negative values.
//   - [NumberFormatter.SetNegativePrefix]
//   - [NumberFormatter.NegativeSuffix]: The string the receiver uses as a suffix for negative values.
//   - [NumberFormatter.SetNegativeSuffix]
//
// # Configuring the Display of Numeric Values
//
//   - [NumberFormatter.TextAttributesForNegativeValues]: The text attributes to be used in displaying negative values.
//   - [NumberFormatter.SetTextAttributesForNegativeValues]
//   - [NumberFormatter.TextAttributesForPositiveValues]: The text attributes to be used in displaying positive values.
//   - [NumberFormatter.SetTextAttributesForPositiveValues]
//   - [NumberFormatter.AttributedStringForZero]: The attributed string that the receiver uses to display zero values.
//   - [NumberFormatter.SetAttributedStringForZero]
//   - [NumberFormatter.TextAttributesForZero]: The text attributes used to display a zero value.
//   - [NumberFormatter.SetTextAttributesForZero]
//   - [NumberFormatter.AttributedStringForNil]: The attributed string the receiver uses to display `nil` values.
//   - [NumberFormatter.SetAttributedStringForNil]
//   - [NumberFormatter.TextAttributesForNil]: The text attributes used to display the `nil` symbol.
//   - [NumberFormatter.SetTextAttributesForNil]
//   - [NumberFormatter.AttributedStringForNotANumber]: The attributed string the receiver uses to display “not a number” values.
//   - [NumberFormatter.SetAttributedStringForNotANumber]
//   - [NumberFormatter.TextAttributesForNotANumber]: The text attributes used to display the NaN (“not a number”) string.
//   - [NumberFormatter.SetTextAttributesForNotANumber]
//   - [NumberFormatter.TextAttributesForPositiveInfinity]: The text attributes used to display the positive infinity symbol.
//   - [NumberFormatter.SetTextAttributesForPositiveInfinity]
//   - [NumberFormatter.TextAttributesForNegativeInfinity]: The text attributes used to display the negative infinity symbol.
//   - [NumberFormatter.SetTextAttributesForNegativeInfinity]
//
// # Configuring Separators and Grouping Size
//
//   - [NumberFormatter.GroupingSeparator]: The string used by the receiver for a grouping separator.
//   - [NumberFormatter.SetGroupingSeparator]
//   - [NumberFormatter.UsesGroupingSeparator]: Determines whether the receiver displays the group separator.
//   - [NumberFormatter.SetUsesGroupingSeparator]
//   - [NumberFormatter.ThousandSeparator]: The character the receiver uses as a thousand separator.
//   - [NumberFormatter.SetThousandSeparator]
//   - [NumberFormatter.HasThousandSeparators]: Determines whether the receiver uses thousand separators.
//   - [NumberFormatter.SetHasThousandSeparators]
//   - [NumberFormatter.DecimalSeparator]: The character the receiver uses as a decimal separator.
//   - [NumberFormatter.SetDecimalSeparator]
//   - [NumberFormatter.AlwaysShowsDecimalSeparator]: Determines whether the receiver always shows the decimal separator, even for integer numbers.
//   - [NumberFormatter.SetAlwaysShowsDecimalSeparator]
//   - [NumberFormatter.CurrencyDecimalSeparator]: The string used by the receiver as a currency decimal separator.
//   - [NumberFormatter.SetCurrencyDecimalSeparator]
//   - [NumberFormatter.GroupingSize]: The grouping size of the receiver.
//   - [NumberFormatter.SetGroupingSize]
//   - [NumberFormatter.SecondaryGroupingSize]: The secondary grouping size of the receiver.
//   - [NumberFormatter.SetSecondaryGroupingSize]
//
// # Managing the Padding of Numbers
//
//   - [NumberFormatter.PaddingCharacter]: The string that the receiver uses to pad numbers in the formatted string representation.
//   - [NumberFormatter.SetPaddingCharacter]
//   - [NumberFormatter.PaddingPosition]: The padding position used by the receiver.
//   - [NumberFormatter.SetPaddingPosition]
//
// # Managing Input and Output Attributes
//
//   - [NumberFormatter.AllowsFloats]: Determines whether the receiver allows as input floating-point values (that is, values that include the period character [`.`]).
//   - [NumberFormatter.SetAllowsFloats]
//   - [NumberFormatter.Minimum]: The lowest number allowed as input by the receiver.
//   - [NumberFormatter.SetMinimum]
//   - [NumberFormatter.Maximum]: The highest number allowed as input by the receiver.
//   - [NumberFormatter.SetMaximum]
//
// # Managing Leniency Behavior
//
//   - [NumberFormatter.Lenient]: Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//   - [NumberFormatter.SetLenient]
//
// # Managing the Validation of Partial Numeric Strings
//
//   - [NumberFormatter.PartialStringValidationEnabled]: Determines whether partial string validation is enabled for the receiver.
//   - [NumberFormatter.SetPartialStringValidationEnabled]
//
// # Instance Properties
//
//   - [NumberFormatter.MinimumGroupingDigits]
//   - [NumberFormatter.SetMinimumGroupingDigits]
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter
type NumberFormatter struct {
	NSFormatter
}

// NumberFormatterFromID constructs a [NumberFormatter] from an objc.ID.
//
// A formatter that converts between numeric values and their textual
// representations.
func NumberFormatterFromID(id objc.ID) NumberFormatter {
	return NSNumberFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSNumberFormatterFromID is an alias for [NumberFormatterFromID] for cross-framework compatibility.
func NSNumberFormatterFromID(id objc.ID) NumberFormatter { return NumberFormatterFromID(id) }

// NOTE: NumberFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NumberFormatter] class.
//
// # Configuring Formatter Behavior and Style
//
//   - [INumberFormatter.FormatterBehavior]: The formatter behavior of the receiver.
//   - [INumberFormatter.SetFormatterBehavior]
//   - [INumberFormatter.NumberStyle]: The number style used by the receiver.
//   - [INumberFormatter.SetNumberStyle]
//   - [INumberFormatter.GeneratesDecimalNumbers]: Determines whether the receiver creates instances of [NSDecimalNumber](<doc://com.apple.foundation/documentation/Foundation/NSDecimalNumber>) when it converts strings to number objects.
//   - [INumberFormatter.SetGeneratesDecimalNumbers]
//
// # Converting Between Numbers and Strings
//
//   - [INumberFormatter.GetObjectValueForStringRangeError]: Returns by reference a cell-content object after creating it from a range of characters in a given string.
//   - [INumberFormatter.NumberFromString]: Returns an [NSNumber](<doc://com.apple.foundation/documentation/Foundation/NSNumber>) object created by parsing a given string.
//   - [INumberFormatter.StringFromNumber]: Returns a string containing the formatted value of the provided number object.
//
// # Managing Localization of Numbers
//
//   - [INumberFormatter.LocalizesFormat]: Determines whether the dollar sign character (`$`), decimal separator character (`.`), and thousand separator character (`,`) are converted to appropriately localized characters as specified by the user’s localization preference.
//   - [INumberFormatter.SetLocalizesFormat]
//   - [INumberFormatter.Locale]: The locale of the receiver.
//   - [INumberFormatter.SetLocale]
//
// # Configuring Rounding Behavior
//
//   - [INumberFormatter.RoundingBehavior]: The rounding behavior used by the receiver.
//   - [INumberFormatter.SetRoundingBehavior]
//   - [INumberFormatter.RoundingIncrement]: The rounding increment used by the receiver.
//   - [INumberFormatter.SetRoundingIncrement]
//   - [INumberFormatter.RoundingMode]: The rounding mode used by the receiver.
//   - [INumberFormatter.SetRoundingMode]
//
// # Configuring Integer and Fraction Digits
//
//   - [INumberFormatter.MinimumIntegerDigits]: The minimum number of digits before the decimal separator.
//   - [INumberFormatter.SetMinimumIntegerDigits]
//   - [INumberFormatter.MaximumIntegerDigits]: The maximum number of digits before the decimal separator.
//   - [INumberFormatter.SetMaximumIntegerDigits]
//   - [INumberFormatter.MinimumFractionDigits]: The minimum number of digits after the decimal separator.
//   - [INumberFormatter.SetMinimumFractionDigits]
//   - [INumberFormatter.MaximumFractionDigits]: The maximum number of digits after the decimal separator.
//   - [INumberFormatter.SetMaximumFractionDigits]
//
// # Configuring Significant Digits
//
//   - [INumberFormatter.UsesSignificantDigits]: A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//   - [INumberFormatter.SetUsesSignificantDigits]
//   - [INumberFormatter.MinimumSignificantDigits]: The minimum number of significant digits for the number formatter.
//   - [INumberFormatter.SetMinimumSignificantDigits]
//   - [INumberFormatter.MaximumSignificantDigits]: The maximum number of significant digits for the number formatter.
//   - [INumberFormatter.SetMaximumSignificantDigits]
//
// # Configuring Numeric Formats
//
//   - [INumberFormatter.Format]: The receiver’s format.
//   - [INumberFormatter.SetFormat]
//   - [INumberFormatter.FormattingContext]: The capitalization formatting context used when formatting a number.
//   - [INumberFormatter.SetFormattingContext]
//   - [INumberFormatter.FormatWidth]: The format width used by the receiver.
//   - [INumberFormatter.SetFormatWidth]
//   - [INumberFormatter.NegativeFormat]: The format the receiver uses to display negative values.
//   - [INumberFormatter.SetNegativeFormat]
//   - [INumberFormatter.PositiveFormat]: The format the receiver uses to display positive values.
//   - [INumberFormatter.SetPositiveFormat]
//   - [INumberFormatter.Multiplier]: The multiplier of the receiver.
//   - [INumberFormatter.SetMultiplier]
//
// # Configuring Numeric Symbols
//
//   - [INumberFormatter.PercentSymbol]: The string used to represent a percent symbol.
//   - [INumberFormatter.SetPercentSymbol]
//   - [INumberFormatter.PerMillSymbol]: The string used to represent a per-mill (per-thousand) symbol.
//   - [INumberFormatter.SetPerMillSymbol]
//   - [INumberFormatter.MinusSign]: The string used to represent a minus sign.
//   - [INumberFormatter.SetMinusSign]
//   - [INumberFormatter.PlusSign]: The string used to represent a plus sign.
//   - [INumberFormatter.SetPlusSign]
//   - [INumberFormatter.ExponentSymbol]: The string used to represent an exponent symbol.
//   - [INumberFormatter.SetExponentSymbol]
//   - [INumberFormatter.ZeroSymbol]: The string used to represent a zero value.
//   - [INumberFormatter.SetZeroSymbol]
//   - [INumberFormatter.NilSymbol]: The string used to represent a `nil` value.
//   - [INumberFormatter.SetNilSymbol]
//   - [INumberFormatter.NotANumberSymbol]: The string used to represent a NaN (“not a number”) value.
//   - [INumberFormatter.SetNotANumberSymbol]
//   - [INumberFormatter.NegativeInfinitySymbol]: The string used to represent a negative infinity symbol.
//   - [INumberFormatter.SetNegativeInfinitySymbol]
//   - [INumberFormatter.PositiveInfinitySymbol]: The string used to represent a positive infinity symbol.
//   - [INumberFormatter.SetPositiveInfinitySymbol]
//
// # Configuring the Format of Currency
//
//   - [INumberFormatter.CurrencySymbol]: The string used by the receiver as a local currency symbol.
//   - [INumberFormatter.SetCurrencySymbol]
//   - [INumberFormatter.CurrencyCode]: The receiver’s currency code.
//   - [INumberFormatter.SetCurrencyCode]
//   - [INumberFormatter.InternationalCurrencySymbol]: The international currency symbol used by the receiver.
//   - [INumberFormatter.SetInternationalCurrencySymbol]
//   - [INumberFormatter.CurrencyGroupingSeparator]: The currency grouping separator for the receiver.
//   - [INumberFormatter.SetCurrencyGroupingSeparator]
//
// # Configuring Numeric Prefixes and Suffixes
//
//   - [INumberFormatter.PositivePrefix]: The string the receiver uses as the prefix for positive values.
//   - [INumberFormatter.SetPositivePrefix]
//   - [INumberFormatter.PositiveSuffix]: The string the receiver uses as the suffix for positive values.
//   - [INumberFormatter.SetPositiveSuffix]
//   - [INumberFormatter.NegativePrefix]: The string the receiver uses as a prefix for negative values.
//   - [INumberFormatter.SetNegativePrefix]
//   - [INumberFormatter.NegativeSuffix]: The string the receiver uses as a suffix for negative values.
//   - [INumberFormatter.SetNegativeSuffix]
//
// # Configuring the Display of Numeric Values
//
//   - [INumberFormatter.TextAttributesForNegativeValues]: The text attributes to be used in displaying negative values.
//   - [INumberFormatter.SetTextAttributesForNegativeValues]
//   - [INumberFormatter.TextAttributesForPositiveValues]: The text attributes to be used in displaying positive values.
//   - [INumberFormatter.SetTextAttributesForPositiveValues]
//   - [INumberFormatter.AttributedStringForZero]: The attributed string that the receiver uses to display zero values.
//   - [INumberFormatter.SetAttributedStringForZero]
//   - [INumberFormatter.TextAttributesForZero]: The text attributes used to display a zero value.
//   - [INumberFormatter.SetTextAttributesForZero]
//   - [INumberFormatter.AttributedStringForNil]: The attributed string the receiver uses to display `nil` values.
//   - [INumberFormatter.SetAttributedStringForNil]
//   - [INumberFormatter.TextAttributesForNil]: The text attributes used to display the `nil` symbol.
//   - [INumberFormatter.SetTextAttributesForNil]
//   - [INumberFormatter.AttributedStringForNotANumber]: The attributed string the receiver uses to display “not a number” values.
//   - [INumberFormatter.SetAttributedStringForNotANumber]
//   - [INumberFormatter.TextAttributesForNotANumber]: The text attributes used to display the NaN (“not a number”) string.
//   - [INumberFormatter.SetTextAttributesForNotANumber]
//   - [INumberFormatter.TextAttributesForPositiveInfinity]: The text attributes used to display the positive infinity symbol.
//   - [INumberFormatter.SetTextAttributesForPositiveInfinity]
//   - [INumberFormatter.TextAttributesForNegativeInfinity]: The text attributes used to display the negative infinity symbol.
//   - [INumberFormatter.SetTextAttributesForNegativeInfinity]
//
// # Configuring Separators and Grouping Size
//
//   - [INumberFormatter.GroupingSeparator]: The string used by the receiver for a grouping separator.
//   - [INumberFormatter.SetGroupingSeparator]
//   - [INumberFormatter.UsesGroupingSeparator]: Determines whether the receiver displays the group separator.
//   - [INumberFormatter.SetUsesGroupingSeparator]
//   - [INumberFormatter.ThousandSeparator]: The character the receiver uses as a thousand separator.
//   - [INumberFormatter.SetThousandSeparator]
//   - [INumberFormatter.HasThousandSeparators]: Determines whether the receiver uses thousand separators.
//   - [INumberFormatter.SetHasThousandSeparators]
//   - [INumberFormatter.DecimalSeparator]: The character the receiver uses as a decimal separator.
//   - [INumberFormatter.SetDecimalSeparator]
//   - [INumberFormatter.AlwaysShowsDecimalSeparator]: Determines whether the receiver always shows the decimal separator, even for integer numbers.
//   - [INumberFormatter.SetAlwaysShowsDecimalSeparator]
//   - [INumberFormatter.CurrencyDecimalSeparator]: The string used by the receiver as a currency decimal separator.
//   - [INumberFormatter.SetCurrencyDecimalSeparator]
//   - [INumberFormatter.GroupingSize]: The grouping size of the receiver.
//   - [INumberFormatter.SetGroupingSize]
//   - [INumberFormatter.SecondaryGroupingSize]: The secondary grouping size of the receiver.
//   - [INumberFormatter.SetSecondaryGroupingSize]
//
// # Managing the Padding of Numbers
//
//   - [INumberFormatter.PaddingCharacter]: The string that the receiver uses to pad numbers in the formatted string representation.
//   - [INumberFormatter.SetPaddingCharacter]
//   - [INumberFormatter.PaddingPosition]: The padding position used by the receiver.
//   - [INumberFormatter.SetPaddingPosition]
//
// # Managing Input and Output Attributes
//
//   - [INumberFormatter.AllowsFloats]: Determines whether the receiver allows as input floating-point values (that is, values that include the period character [`.`]).
//   - [INumberFormatter.SetAllowsFloats]
//   - [INumberFormatter.Minimum]: The lowest number allowed as input by the receiver.
//   - [INumberFormatter.SetMinimum]
//   - [INumberFormatter.Maximum]: The highest number allowed as input by the receiver.
//   - [INumberFormatter.SetMaximum]
//
// # Managing Leniency Behavior
//
//   - [INumberFormatter.Lenient]: Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//   - [INumberFormatter.SetLenient]
//
// # Managing the Validation of Partial Numeric Strings
//
//   - [INumberFormatter.PartialStringValidationEnabled]: Determines whether partial string validation is enabled for the receiver.
//   - [INumberFormatter.SetPartialStringValidationEnabled]
//
// # Instance Properties
//
//   - [INumberFormatter.MinimumGroupingDigits]
//   - [INumberFormatter.SetMinimumGroupingDigits]
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter
type INumberFormatter interface {
	INSFormatter

	// Topic: Configuring Formatter Behavior and Style

	// The formatter behavior of the receiver.
	FormatterBehavior() NSNumberFormatterBehavior
	SetFormatterBehavior(value NSNumberFormatterBehavior)
	// The number style used by the receiver.
	NumberStyle() NSNumberFormatterStyle
	SetNumberStyle(value NSNumberFormatterStyle)
	// Determines whether the receiver creates instances of [NSDecimalNumber](<doc://com.apple.foundation/documentation/Foundation/NSDecimalNumber>) when it converts strings to number objects.
	GeneratesDecimalNumbers() bool
	SetGeneratesDecimalNumbers(value bool)

	// Topic: Converting Between Numbers and Strings

	// Returns by reference a cell-content object after creating it from a range of characters in a given string.
	GetObjectValueForStringRangeError(obj []objectivec.IObject, string_ string, rangep NSRange) (bool, error)
	// Returns an [NSNumber](<doc://com.apple.foundation/documentation/Foundation/NSNumber>) object created by parsing a given string.
	NumberFromString(string_ string) INSNumber
	// Returns a string containing the formatted value of the provided number object.
	StringFromNumber(number INSNumber) string

	// Topic: Managing Localization of Numbers

	// Determines whether the dollar sign character (`$`), decimal separator character (`.`), and thousand separator character (`,`) are converted to appropriately localized characters as specified by the user’s localization preference.
	LocalizesFormat() bool
	SetLocalizesFormat(value bool)
	// The locale of the receiver.
	Locale() INSLocale
	SetLocale(value INSLocale)

	// Topic: Configuring Rounding Behavior

	// The rounding behavior used by the receiver.
	RoundingBehavior() INSDecimalNumberHandler
	SetRoundingBehavior(value INSDecimalNumberHandler)
	// The rounding increment used by the receiver.
	RoundingIncrement() INSNumber
	SetRoundingIncrement(value INSNumber)
	// The rounding mode used by the receiver.
	RoundingMode() NSNumberFormatterRoundingMode
	SetRoundingMode(value NSNumberFormatterRoundingMode)

	// Topic: Configuring Integer and Fraction Digits

	// The minimum number of digits before the decimal separator.
	MinimumIntegerDigits() uint
	SetMinimumIntegerDigits(value uint)
	// The maximum number of digits before the decimal separator.
	MaximumIntegerDigits() uint
	SetMaximumIntegerDigits(value uint)
	// The minimum number of digits after the decimal separator.
	MinimumFractionDigits() uint
	SetMinimumFractionDigits(value uint)
	// The maximum number of digits after the decimal separator.
	MaximumFractionDigits() uint
	SetMaximumFractionDigits(value uint)

	// Topic: Configuring Significant Digits

	// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
	UsesSignificantDigits() bool
	SetUsesSignificantDigits(value bool)
	// The minimum number of significant digits for the number formatter.
	MinimumSignificantDigits() uint
	SetMinimumSignificantDigits(value uint)
	// The maximum number of significant digits for the number formatter.
	MaximumSignificantDigits() uint
	SetMaximumSignificantDigits(value uint)

	// Topic: Configuring Numeric Formats

	// The receiver’s format.
	Format() string
	SetFormat(value string)
	// The capitalization formatting context used when formatting a number.
	FormattingContext() NSFormattingContext
	SetFormattingContext(value NSFormattingContext)
	// The format width used by the receiver.
	FormatWidth() uint
	SetFormatWidth(value uint)
	// The format the receiver uses to display negative values.
	NegativeFormat() string
	SetNegativeFormat(value string)
	// The format the receiver uses to display positive values.
	PositiveFormat() string
	SetPositiveFormat(value string)
	// The multiplier of the receiver.
	Multiplier() INSNumber
	SetMultiplier(value INSNumber)

	// Topic: Configuring Numeric Symbols

	// The string used to represent a percent symbol.
	PercentSymbol() string
	SetPercentSymbol(value string)
	// The string used to represent a per-mill (per-thousand) symbol.
	PerMillSymbol() string
	SetPerMillSymbol(value string)
	// The string used to represent a minus sign.
	MinusSign() string
	SetMinusSign(value string)
	// The string used to represent a plus sign.
	PlusSign() string
	SetPlusSign(value string)
	// The string used to represent an exponent symbol.
	ExponentSymbol() string
	SetExponentSymbol(value string)
	// The string used to represent a zero value.
	ZeroSymbol() string
	SetZeroSymbol(value string)
	// The string used to represent a `nil` value.
	NilSymbol() string
	SetNilSymbol(value string)
	// The string used to represent a NaN (“not a number”) value.
	NotANumberSymbol() string
	SetNotANumberSymbol(value string)
	// The string used to represent a negative infinity symbol.
	NegativeInfinitySymbol() string
	SetNegativeInfinitySymbol(value string)
	// The string used to represent a positive infinity symbol.
	PositiveInfinitySymbol() string
	SetPositiveInfinitySymbol(value string)

	// Topic: Configuring the Format of Currency

	// The string used by the receiver as a local currency symbol.
	CurrencySymbol() string
	SetCurrencySymbol(value string)
	// The receiver’s currency code.
	CurrencyCode() string
	SetCurrencyCode(value string)
	// The international currency symbol used by the receiver.
	InternationalCurrencySymbol() string
	SetInternationalCurrencySymbol(value string)
	// The currency grouping separator for the receiver.
	CurrencyGroupingSeparator() string
	SetCurrencyGroupingSeparator(value string)

	// Topic: Configuring Numeric Prefixes and Suffixes

	// The string the receiver uses as the prefix for positive values.
	PositivePrefix() string
	SetPositivePrefix(value string)
	// The string the receiver uses as the suffix for positive values.
	PositiveSuffix() string
	SetPositiveSuffix(value string)
	// The string the receiver uses as a prefix for negative values.
	NegativePrefix() string
	SetNegativePrefix(value string)
	// The string the receiver uses as a suffix for negative values.
	NegativeSuffix() string
	SetNegativeSuffix(value string)

	// Topic: Configuring the Display of Numeric Values

	// The text attributes to be used in displaying negative values.
	TextAttributesForNegativeValues() INSDictionary
	SetTextAttributesForNegativeValues(value INSDictionary)
	// The text attributes to be used in displaying positive values.
	TextAttributesForPositiveValues() INSDictionary
	SetTextAttributesForPositiveValues(value INSDictionary)
	// The attributed string that the receiver uses to display zero values.
	AttributedStringForZero() INSAttributedString
	SetAttributedStringForZero(value INSAttributedString)
	// The text attributes used to display a zero value.
	TextAttributesForZero() INSDictionary
	SetTextAttributesForZero(value INSDictionary)
	// The attributed string the receiver uses to display `nil` values.
	AttributedStringForNil() INSAttributedString
	SetAttributedStringForNil(value INSAttributedString)
	// The text attributes used to display the `nil` symbol.
	TextAttributesForNil() INSDictionary
	SetTextAttributesForNil(value INSDictionary)
	// The attributed string the receiver uses to display “not a number” values.
	AttributedStringForNotANumber() INSAttributedString
	SetAttributedStringForNotANumber(value INSAttributedString)
	// The text attributes used to display the NaN (“not a number”) string.
	TextAttributesForNotANumber() INSDictionary
	SetTextAttributesForNotANumber(value INSDictionary)
	// The text attributes used to display the positive infinity symbol.
	TextAttributesForPositiveInfinity() INSDictionary
	SetTextAttributesForPositiveInfinity(value INSDictionary)
	// The text attributes used to display the negative infinity symbol.
	TextAttributesForNegativeInfinity() INSDictionary
	SetTextAttributesForNegativeInfinity(value INSDictionary)

	// Topic: Configuring Separators and Grouping Size

	// The string used by the receiver for a grouping separator.
	GroupingSeparator() string
	SetGroupingSeparator(value string)
	// Determines whether the receiver displays the group separator.
	UsesGroupingSeparator() bool
	SetUsesGroupingSeparator(value bool)
	// The character the receiver uses as a thousand separator.
	ThousandSeparator() string
	SetThousandSeparator(value string)
	// Determines whether the receiver uses thousand separators.
	HasThousandSeparators() bool
	SetHasThousandSeparators(value bool)
	// The character the receiver uses as a decimal separator.
	DecimalSeparator() string
	SetDecimalSeparator(value string)
	// Determines whether the receiver always shows the decimal separator, even for integer numbers.
	AlwaysShowsDecimalSeparator() bool
	SetAlwaysShowsDecimalSeparator(value bool)
	// The string used by the receiver as a currency decimal separator.
	CurrencyDecimalSeparator() string
	SetCurrencyDecimalSeparator(value string)
	// The grouping size of the receiver.
	GroupingSize() uint
	SetGroupingSize(value uint)
	// The secondary grouping size of the receiver.
	SecondaryGroupingSize() uint
	SetSecondaryGroupingSize(value uint)

	// Topic: Managing the Padding of Numbers

	// The string that the receiver uses to pad numbers in the formatted string representation.
	PaddingCharacter() string
	SetPaddingCharacter(value string)
	// The padding position used by the receiver.
	PaddingPosition() NSNumberFormatterPadPosition
	SetPaddingPosition(value NSNumberFormatterPadPosition)

	// Topic: Managing Input and Output Attributes

	// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [`.`]).
	AllowsFloats() bool
	SetAllowsFloats(value bool)
	// The lowest number allowed as input by the receiver.
	Minimum() INSNumber
	SetMinimum(value INSNumber)
	// The highest number allowed as input by the receiver.
	Maximum() INSNumber
	SetMaximum(value INSNumber)

	// Topic: Managing Leniency Behavior

	// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
	Lenient() bool
	SetLenient(value bool)

	// Topic: Managing the Validation of Partial Numeric Strings

	// Determines whether partial string validation is enabled for the receiver.
	PartialStringValidationEnabled() bool
	SetPartialStringValidationEnabled(value bool)

	// Topic: Instance Properties

	MinimumGroupingDigits() int
	SetMinimumGroupingDigits(value int)
}

// Init initializes the instance.
func (n NumberFormatter) Init() NumberFormatter {
	rv := objc.Send[NumberFormatter](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NumberFormatter) Autorelease() NumberFormatter {
	rv := objc.Send[NumberFormatter](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumberFormatter creates a new NumberFormatter instance.
func NewNumberFormatter() NumberFormatter {
	class := getNumberFormatterClass()
	rv := objc.Send[NumberFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewNumberFormatterWithCoder(coder INSCoder) NumberFormatter {
	instance := getNumberFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NumberFormatterFromID(rv)
}

// Returns by reference a cell-content object after creating it from a range
// of characters in a given string.
//
// obj: On return, contains an instance of [NSDecimalNumber] or [NSNumber] based on
// the current value of the [GeneratesDecimalNumbers] property. Returns `nil`
// by reference if conversion failed.
//
// string: A string object with the range of characters specified in `rangep` that is
// used to create `anObject`.
//
// rangep: A range of characters in `aString`. On return, contains the actual range of
// characters used to create the object.
//
// # Discussion
//
// If a string contains any characters other than numerical digits or
// locale-appropriate group or decimal separators, parsing will fail.
//
// Any leading or trailing space separator characters in a string are ignored.
// For example, the strings “ 5”, “5 “, and “5” all produce the
// number `5`.
//
// If there is an error, this method calls `control(_:)` on the delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/getObjectValue(_:for:range:)
func (n NumberFormatter) GetObjectValueForStringRangeError(obj []objectivec.IObject, string_ string, rangep NSRange) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("getObjectValue:forString:range:error:"), objectivec.IObjectSliceToNSArray(obj), objc.String(string_), rangep, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getObjectValue:forString:range:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns an [NSNumber] object created by parsing a given string.
//
// string: An [NSString] object that is parsed to generate the returned number object.
//
// # Return Value
//
// An [NSNumber] object created by parsing `string` using the receiver’s
// format, or `nil` if no single number could be parsed.
//
// # Discussion
//
// If a string contains any characters other than numerical digits or
// locale-appropriate group or decimal separators, parsing will fail.
//
// Any leading or trailing space separator characters in a string are ignored.
// For example, the strings “ 5”, “5 “, and “5” all produce the
// number `5`.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/number(from:)
func (n NumberFormatter) NumberFromString(string_ string) INSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("numberFromString:"), objc.String(string_))
	return NSNumberFromID(rv)
}

// Returns a string containing the formatted value of the provided number
// object.
//
// number: An [NSNumber] object that is parsed to create the returned string object.
//
// # Return Value
//
// A string containing the formatted value of `number` using the receiver’s
// current settings.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/string(from:)
func (n NumberFormatter) StringFromNumber(number INSNumber) string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("stringFromNumber:"), number)
	return NSStringFromID(rv).String()
}

// Sets the default formatter behavior for new instances of
// [NSNumberFormatter] .
//
// behavior: An [NSNumberFormatterBehavior] constant that indicates the revision of the
// class providing the default behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/setDefaultFormatterBehavior(_:)
func (_NumberFormatterClass NumberFormatterClass) SetDefaultFormatterBehavior(behavior NSNumberFormatterBehavior) {
	objc.Send[objc.ID](objc.ID(_NumberFormatterClass.class), objc.Sel("setDefaultFormatterBehavior:"), behavior)
}

// Returns an [NSNumberFormatterBehavior] constant that indicates default
// formatter behavior for new instances of [NSNumberFormatter].
//
// # Return Value
//
// An [NSNumberFormatterBehavior] constant that indicates default formatter
// behavior for new instances of [NSNumberFormatter].
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/defaultFormatterBehavior()
func (_NumberFormatterClass NumberFormatterClass) DefaultFormatterBehavior() NSNumberFormatterBehavior {
	rv := objc.Send[NSNumberFormatterBehavior](objc.ID(_NumberFormatterClass.class), objc.Sel("defaultFormatterBehavior"))
	return NSNumberFormatterBehavior(rv)
}

// Returns a localized number string with the specified style.
//
// num: The number to localize
//
// nstyle: The localization style to use. See [NumberFormatter.Style] for the
// supported values.
//
// # Return Value
//
// An appropriately formatted [NSString].
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizedString(from:number:)
//
// [NumberFormatter.Style]: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style
func (_NumberFormatterClass NumberFormatterClass) LocalizedStringFromNumberNumberStyle(num INSNumber, nstyle NSNumberFormatterStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_NumberFormatterClass.class), objc.Sel("localizedStringFromNumber:numberStyle:"), num, nstyle)
	return NSStringFromID(rv).String()
}

// The formatter behavior of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n NumberFormatter) FormatterBehavior() NSNumberFormatterBehavior {
	rv := objc.Send[NSNumberFormatterBehavior](n.ID, objc.Sel("formatterBehavior"))
	return NSNumberFormatterBehavior(rv)
}
func (n NumberFormatter) SetFormatterBehavior(value NSNumberFormatterBehavior) {
	objc.Send[struct{}](n.ID, objc.Sel("setFormatterBehavior:"), value)
}

// The number style used by the receiver.
//
// # Discussion
//
// Styles are essentially predetermined sets of values for certain properties.
// Examples of number-formatter styles are those used for decimal values,
// percentage values, and currency.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n NumberFormatter) NumberStyle() NSNumberFormatterStyle {
	rv := objc.Send[NSNumberFormatterStyle](n.ID, objc.Sel("numberStyle"))
	return NSNumberFormatterStyle(rv)
}
func (n NumberFormatter) SetNumberStyle(value NSNumberFormatterStyle) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumberStyle:"), value)
}

// Determines whether the receiver creates instances of [NSDecimalNumber] when
// it converts strings to number objects.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/generatesDecimalNumbers
func (n NumberFormatter) GeneratesDecimalNumbers() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("generatesDecimalNumbers"))
	return rv
}
func (n NumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setGeneratesDecimalNumbers:"), value)
}

// Determines whether the dollar sign character (`$`), decimal separator
// character (`.`), and thousand separator character (`,`) are converted to
// appropriately localized characters as specified by the user’s
// localization preference.
//
// # Discussion
//
// While the currency-symbol part of this feature may be useful in certain
// types of applications, it’s probably more likely that you would tie a
// particular application to a particular currency (that is, that you would
// “hard-code” the currency symbol and separators instead of having them
// dynamically change based on the user’s configuration). The reason for
// this, of course, is that [NSNumberFormatter] doesn’t perform currency
// conversions, it just formats numeric data. You wouldn’t want one user
// interpreting the value `"56324"` as US currency and another user who’s
// accessing the same data interpreting it as Japanese currency, simply based
// on each user’s localization preferences.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n NumberFormatter) LocalizesFormat() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("localizesFormat"))
	return rv
}
func (n NumberFormatter) SetLocalizesFormat(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setLocalizesFormat:"), value)
}

// The locale of the receiver.
//
// # Discussion
//
// The locale determines the default values for many formatter attributes,
// such as ISO region and language codes, currency code, calendar, system of
// measurement, and decimal separator.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n NumberFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (n NumberFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](n.ID, objc.Sel("setLocale:"), value)
}

// The rounding behavior used by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n NumberFormatter) RoundingBehavior() INSDecimalNumberHandler {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("roundingBehavior"))
	return NSDecimalNumberHandlerFromID(objc.ID(rv))
}
func (n NumberFormatter) SetRoundingBehavior(value INSDecimalNumberHandler) {
	objc.Send[struct{}](n.ID, objc.Sel("setRoundingBehavior:"), value)
}

// The rounding increment used by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n NumberFormatter) RoundingIncrement() INSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("roundingIncrement"))
	return NSNumberFromID(objc.ID(rv))
}
func (n NumberFormatter) SetRoundingIncrement(value INSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setRoundingIncrement:"), value)
}

// The rounding mode used by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n NumberFormatter) RoundingMode() NSNumberFormatterRoundingMode {
	rv := objc.Send[NSNumberFormatterRoundingMode](n.ID, objc.Sel("roundingMode"))
	return NSNumberFormatterRoundingMode(rv)
}
func (n NumberFormatter) SetRoundingMode(value NSNumberFormatterRoundingMode) {
	objc.Send[struct{}](n.ID, objc.Sel("setRoundingMode:"), value)
}

// The minimum number of digits before the decimal separator.
//
// # Discussion
//
// By default, this property is set to `0`.
//
// The following code demonstrates the effect of setting
// [MinimumIntegerDigits] when formatting a number:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n NumberFormatter) MinimumIntegerDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("minimumIntegerDigits"))
	return rv
}
func (n NumberFormatter) SetMinimumIntegerDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumIntegerDigits:"), value)
}

// The maximum number of digits before the decimal separator.
//
// # Discussion
//
// By default, this property is set to `42`.
//
// The following code demonstrates the effect of setting
// [MaximumIntegerDigits] when formatting a number:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n NumberFormatter) MaximumIntegerDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("maximumIntegerDigits"))
	return rv
}
func (n NumberFormatter) SetMaximumIntegerDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumIntegerDigits:"), value)
}

// The minimum number of digits after the decimal separator.
//
// # Discussion
//
// By default, this property is set to `0`.
//
// The following code demonstrates the effect of setting
// [MinimumFractionDigits] when formatting a number:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n NumberFormatter) MinimumFractionDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("minimumFractionDigits"))
	return rv
}
func (n NumberFormatter) SetMinimumFractionDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumFractionDigits:"), value)
}

// The maximum number of digits after the decimal separator.
//
// # Discussion
//
// By default, this property is set to `0`.
//
// The following code demonstrates the effect of setting
// [MaximumFractionDigits] when formatting a number:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n NumberFormatter) MaximumFractionDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("maximumFractionDigits"))
	return rv
}
func (n NumberFormatter) SetMaximumFractionDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumFractionDigits:"), value)
}

// A Boolean value indicating whether the formatter uses minimum and maximum
// significant digits when formatting numbers.
//
// # Discussion
//
// The [NSNumberFormatter] class has two ways of determining how many digits
// to represent:_ _using integer and fraction digits and using significant
// digits.
//
// When this property is set to false, numbers are formatted according to
// whether you want them formatted as fractions or as integers. For more
// information, see Configuring Integer and Fraction Digits. This property is
// false by default.
//
// Set this property to true to format numbers according to the significant
// digits configuration specified by the [MinimumSignificantDigits] and
// [MaximumSignificantDigits] properties. By default, the minimum number of
// significant digits is 1, and the maximum number of significant digits is 6.
//
// The following code demonstrates the effect of configuring
// [UsesSignificantDigits] when formatting various numbers:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n NumberFormatter) UsesSignificantDigits() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesSignificantDigits"))
	return rv
}
func (n NumberFormatter) SetUsesSignificantDigits(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUsesSignificantDigits:"), value)
}

// The minimum number of significant digits for the number formatter.
//
// # Discussion
//
// You must set the [UsesSignificantDigits] property to true in order for this
// property to affect formatting behavior. By default, the minimum number of
// significant digits is 1.
//
// The following code demonstrates the effect of setting
// [MinimumSignificantDigits] when formatting various numbers:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n NumberFormatter) MinimumSignificantDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("minimumSignificantDigits"))
	return rv
}
func (n NumberFormatter) SetMinimumSignificantDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumSignificantDigits:"), value)
}

// The maximum number of significant digits for the number formatter.
//
// # Discussion
//
// You must set the [UsesSignificantDigits] property to true in order for this
// property to affect formatting behavior. By default, the maximum number of
// significant digits is 6. Values less than 1 are ignored.
//
// The following code demonstrates the effect of setting
// [MaximumSignificantDigits] when formatting various numbers:
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n NumberFormatter) MaximumSignificantDigits() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("maximumSignificantDigits"))
	return rv
}
func (n NumberFormatter) SetMaximumSignificantDigits(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumSignificantDigits:"), value)
}

// The receiver’s format.
//
// # Discussion
//
// The format string uses the format patterns from [Unicode Technical Standard
// #35]. For more information, see [Data Formatting Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
// [Unicode Technical Standard #35]: http://www.unicode.org/reports/tr35/tr35-numbers.html#Number_Format_Patterns
func (n NumberFormatter) Format() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("format"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetFormat(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setFormat:"), objc.String(value))
}

// The capitalization formatting context used when formatting a number.
//
// # Discussion
//
// Defaults to NSFormattingContextUnknown.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n NumberFormatter) FormattingContext() NSFormattingContext {
	rv := objc.Send[NSFormattingContext](n.ID, objc.Sel("formattingContext"))
	return NSFormattingContext(rv)
}
func (n NumberFormatter) SetFormattingContext(value NSFormattingContext) {
	objc.Send[struct{}](n.ID, objc.Sel("setFormattingContext:"), value)
}

// The format width used by the receiver.
//
// # Discussion
//
// The format width is the number of characters of a formatted number within a
// string that is either left justified or right justified based on the value
// contained in [PaddingPosition].
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n NumberFormatter) FormatWidth() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("formatWidth"))
	return rv
}
func (n NumberFormatter) SetFormatWidth(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setFormatWidth:"), value)
}

// The format the receiver uses to display negative values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n NumberFormatter) NegativeFormat() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("negativeFormat"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNegativeFormat(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNegativeFormat:"), objc.String(value))
}

// The format the receiver uses to display positive values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n NumberFormatter) PositiveFormat() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("positiveFormat"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPositiveFormat(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPositiveFormat:"), objc.String(value))
}

// The multiplier of the receiver.
//
// # Discussion
//
// A multiplier is a factor used in conversions between numbers and strings
// (that is, numbers as stored and numbers as displayed). When the input value
// is a string, the multiplier is used to divide, and when the input value is
// a number, the multiplier is used to multiply. These operations allow the
// formatted values to be different from the values that a program manipulates
// internally.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n NumberFormatter) Multiplier() INSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("multiplier"))
	return NSNumberFromID(objc.ID(rv))
}
func (n NumberFormatter) SetMultiplier(value INSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setMultiplier:"), value)
}

// The string used to represent a percent symbol.
//
// # Discussion
//
// By default, this property is set to the percent sign (%).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n NumberFormatter) PercentSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("percentSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPercentSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPercentSymbol:"), objc.String(value))
}

// The string used to represent a per-mill (per-thousand) symbol.
//
// # Discussion
//
// By default, this property is set to the per mille sign (‰).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n NumberFormatter) PerMillSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("perMillSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPerMillSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPerMillSymbol:"), objc.String(value))
}

// The string used to represent a minus sign.
//
// # Discussion
//
// By default, this property is set to the minus sign (-).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n NumberFormatter) MinusSign() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("minusSign"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetMinusSign(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinusSign:"), objc.String(value))
}

// The string used to represent a plus sign.
//
// # Discussion
//
// By default, this property is set to the plus sign (+).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n NumberFormatter) PlusSign() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("plusSign"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPlusSign(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPlusSign:"), objc.String(value))
}

// The string used to represent an exponent symbol.
//
// # Discussion
//
// By default, this property is set to the latin capital letter e (E).
//
// The exponent symbol is the “E” or “e” in the scientific notation of
// numbers, as in “1.0E+42”.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n NumberFormatter) ExponentSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("exponentSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetExponentSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setExponentSymbol:"), objc.String(value))
}

// The string used to represent a zero value.
//
// # Discussion
//
// If not specified, zero values are formatted normally.
//
// You might, for example, set this property to “` “-“ `” in a
// spreadsheet used for accounting.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n NumberFormatter) ZeroSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("zeroSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetZeroSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setZeroSymbol:"), objc.String(value))
}

// The string used to represent a `nil` value.
//
// # Discussion
//
// By default, this property is set to an empty string (””).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n NumberFormatter) NilSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("nilSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNilSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNilSymbol:"), objc.String(value))
}

// The string used to represent a NaN (“not a number”) value.
//
// # Discussion
//
// By default, this property is set to the string “[NaN]”.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n NumberFormatter) NotANumberSymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("notANumberSymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNotANumberSymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNotANumberSymbol:"), objc.String(value))
}

// The string used to represent a negative infinity symbol.
//
// # Discussion
//
// By default, this property is set to the string “`-∞`”.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n NumberFormatter) NegativeInfinitySymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("negativeInfinitySymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNegativeInfinitySymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNegativeInfinitySymbol:"), objc.String(value))
}

// The string used to represent a positive infinity symbol.
//
// # Discussion
//
// By default, this property is set to the string “`+∞`”.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n NumberFormatter) PositiveInfinitySymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("positiveInfinitySymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPositiveInfinitySymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPositiveInfinitySymbol:"), objc.String(value))
}

// The string used by the receiver as a local currency symbol.
//
// # Discussion
//
// A region typically has a local currency symbol and an international
// currency symbol. The local symbol is used within the region, while the
// international currency symbol is used in international contexts to specify
// that region’s currency unambiguously. The local currency symbol is often
// represented by a Unicode code point.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n NumberFormatter) CurrencySymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currencySymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetCurrencySymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCurrencySymbol:"), objc.String(value))
}

// The receiver’s currency code.
//
// # Discussion
//
// A currency code is a three-letter code that is, in most cases, composed of
// a region’s two-character Internet region code plus an extra character to
// denote the currency unit. For example, the currency code for the Australian
// dollar is “AUD”. Currency codes are based on the ISO 4217 standard.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n NumberFormatter) CurrencyCode() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currencyCode"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetCurrencyCode(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCurrencyCode:"), objc.String(value))
}

// The international currency symbol used by the receiver.
//
// # Discussion
//
// A region typically has a local currency symbol and an international
// currency symbol. The local symbol is used within the region, while the
// international currency symbol is used in international contexts to specify
// that region’s currency unambiguously. The international currency symbol
// is often represented by a Unicode code point.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n NumberFormatter) InternationalCurrencySymbol() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internationalCurrencySymbol"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetInternationalCurrencySymbol(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternationalCurrencySymbol:"), objc.String(value))
}

// The currency grouping separator for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n NumberFormatter) CurrencyGroupingSeparator() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currencyGroupingSeparator"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetCurrencyGroupingSeparator(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCurrencyGroupingSeparator:"), objc.String(value))
}

// The string the receiver uses as the prefix for positive values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/positivePrefix
func (n NumberFormatter) PositivePrefix() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("positivePrefix"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPositivePrefix(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPositivePrefix:"), objc.String(value))
}

// The string the receiver uses as the suffix for positive values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n NumberFormatter) PositiveSuffix() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("positiveSuffix"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPositiveSuffix(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPositiveSuffix:"), objc.String(value))
}

// The string the receiver uses as a prefix for negative values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n NumberFormatter) NegativePrefix() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("negativePrefix"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNegativePrefix(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNegativePrefix:"), objc.String(value))
}

// The string the receiver uses as a suffix for negative values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n NumberFormatter) NegativeSuffix() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("negativeSuffix"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetNegativeSuffix(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setNegativeSuffix:"), objc.String(value))
}

// The text attributes to be used in displaying negative values.
//
// # Discussion
//
// This property is a dictionary that contains the attributes used to display
// negative values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n NumberFormatter) TextAttributesForNegativeValues() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForNegativeValues"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForNegativeValues(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForNegativeValues:"), value)
}

// The text attributes to be used in displaying positive values.
//
// # Discussion
//
// This property is a dictionary that contains the attributes used to display
// positive values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n NumberFormatter) TextAttributesForPositiveValues() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForPositiveValues"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForPositiveValues(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForPositiveValues:"), value)
}

// The attributed string that the receiver uses to display zero values.
//
// # Discussion
//
// By default zero values are displayed according to the format specified for
// positive values; for more discussion of this subject see [Data Formatting
// Guide].
//
// # Special Considerations
//
// This method is for use with formatters using
// `NSNumberFormatterBehavior10_0` behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
func (n NumberFormatter) AttributedStringForZero() INSAttributedString {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attributedStringForZero"))
	return NSAttributedStringFromID(objc.ID(rv))
}
func (n NumberFormatter) SetAttributedStringForZero(value INSAttributedString) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttributedStringForZero:"), value)
}

// The text attributes used to display a zero value.
//
// # Discussion
//
// This property is a dictionary that contains the text attributes used to
// display zero values.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n NumberFormatter) TextAttributesForZero() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForZero"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForZero(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForZero:"), value)
}

// The attributed string the receiver uses to display `nil` values.
//
// # Discussion
//
// By default `nil` values are displayed as an empty string.
//
// # Special Considerations
//
// This method is for use with formatters using
// `NSNumberFormatterBehavior10_0` behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n NumberFormatter) AttributedStringForNil() INSAttributedString {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attributedStringForNil"))
	return NSAttributedStringFromID(objc.ID(rv))
}
func (n NumberFormatter) SetAttributedStringForNil(value INSAttributedString) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttributedStringForNil:"), value)
}

// The text attributes used to display the `nil` symbol.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n NumberFormatter) TextAttributesForNil() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForNil"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForNil(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForNil:"), value)
}

// The attributed string the receiver uses to display “not a number”
// values.
//
// # Discussion
//
// By default “not a number” values are displayed as the string “NaN”.
//
// # Special Considerations
//
// This method is for use with formatters using
// `NSNumberFormatterBehavior10_0` behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n NumberFormatter) AttributedStringForNotANumber() INSAttributedString {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attributedStringForNotANumber"))
	return NSAttributedStringFromID(objc.ID(rv))
}
func (n NumberFormatter) SetAttributedStringForNotANumber(value INSAttributedString) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttributedStringForNotANumber:"), value)
}

// The text attributes used to display the NaN (“not a number”) string.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNotANumber
func (n NumberFormatter) TextAttributesForNotANumber() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForNotANumber"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForNotANumber(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForNotANumber:"), value)
}

// The text attributes used to display the positive infinity symbol.
//
// # Discussion
//
// This property is a dictionary that contains the text attributes used to
// display the positive infinity string.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n NumberFormatter) TextAttributesForPositiveInfinity() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForPositiveInfinity"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForPositiveInfinity(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForPositiveInfinity:"), value)
}

// The text attributes used to display the negative infinity symbol.
//
// # Discussion
//
// This property is a dictionary that contains the text attributes used to
// display the negative infinity string.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n NumberFormatter) TextAttributesForNegativeInfinity() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("textAttributesForNegativeInfinity"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (n NumberFormatter) SetTextAttributesForNegativeInfinity(value INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setTextAttributesForNegativeInfinity:"), value)
}

// The string used by the receiver for a grouping separator.
//
// # Discussion
//
// For example, the grouping separator used in the United States is the comma
// (“10,000”) whereas in France it is the space (“10 000”).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n NumberFormatter) GroupingSeparator() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("groupingSeparator"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetGroupingSeparator(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setGroupingSeparator:"), objc.String(value))
}

// Determines whether the receiver displays the group separator.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n NumberFormatter) UsesGroupingSeparator() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesGroupingSeparator"))
	return rv
}
func (n NumberFormatter) SetUsesGroupingSeparator(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUsesGroupingSeparator:"), value)
}

// The character the receiver uses as a thousand separator.
//
// # Discussion
//
// If you don’t have thousand separators enabled through any other means
// (such as [Format]), using this method enables them.
//
// # Special Considerations
//
// This method is for use with formatters using
// `NSNumberFormatterBehavior10_0` behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n NumberFormatter) ThousandSeparator() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("thousandSeparator"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetThousandSeparator(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setThousandSeparator:"), objc.String(value))
}

// Determines whether the receiver uses thousand separators.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n NumberFormatter) HasThousandSeparators() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasThousandSeparators"))
	return rv
}
func (n NumberFormatter) SetHasThousandSeparators(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasThousandSeparators:"), value)
}

// The character the receiver uses as a decimal separator.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n NumberFormatter) DecimalSeparator() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("decimalSeparator"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetDecimalSeparator(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setDecimalSeparator:"), objc.String(value))
}

// Determines whether the receiver always shows the decimal separator, even
// for integer numbers.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("alwaysShowsDecimalSeparator"))
	return rv
}
func (n NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAlwaysShowsDecimalSeparator:"), value)
}

// The string used by the receiver as a currency decimal separator.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n NumberFormatter) CurrencyDecimalSeparator() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currencyDecimalSeparator"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetCurrencyDecimalSeparator(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCurrencyDecimalSeparator:"), objc.String(value))
}

// The grouping size of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n NumberFormatter) GroupingSize() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("groupingSize"))
	return rv
}
func (n NumberFormatter) SetGroupingSize(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setGroupingSize:"), value)
}

// The secondary grouping size of the receiver.
//
// # Discussion
//
// Some locales allow the specification of another grouping size for larger
// numbers. For example, some locales may represent a number such as 61, 242,
// 378.46 (as in the United States) as 6,12,42,378.46. In this case, the
// secondary grouping size (covering the groups of digits furthest from the
// decimal point) is 2.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n NumberFormatter) SecondaryGroupingSize() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("secondaryGroupingSize"))
	return rv
}
func (n NumberFormatter) SetSecondaryGroupingSize(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setSecondaryGroupingSize:"), value)
}

// The string that the receiver uses to pad numbers in the formatted string
// representation.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n NumberFormatter) PaddingCharacter() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("paddingCharacter"))
	return NSStringFromID(rv).String()
}
func (n NumberFormatter) SetPaddingCharacter(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setPaddingCharacter:"), objc.String(value))
}

// The padding position used by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n NumberFormatter) PaddingPosition() NSNumberFormatterPadPosition {
	rv := objc.Send[NSNumberFormatterPadPosition](n.ID, objc.Sel("paddingPosition"))
	return NSNumberFormatterPadPosition(rv)
}
func (n NumberFormatter) SetPaddingPosition(value NSNumberFormatterPadPosition) {
	objc.Send[struct{}](n.ID, objc.Sel("setPaddingPosition:"), value)
}

// Determines whether the receiver allows as input floating-point values (that
// is, values that include the period character [`.`]).
//
// # Discussion
//
// By default, floating point values are allowed.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n NumberFormatter) AllowsFloats() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("allowsFloats"))
	return rv
}
func (n NumberFormatter) SetAllowsFloats(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAllowsFloats:"), value)
}

// The lowest number allowed as input by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n NumberFormatter) Minimum() INSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("minimum"))
	return NSNumberFromID(objc.ID(rv))
}
func (n NumberFormatter) SetMinimum(value INSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimum:"), value)
}

// The highest number allowed as input by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n NumberFormatter) Maximum() INSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("maximum"))
	return NSNumberFromID(objc.ID(rv))
}
func (n NumberFormatter) SetMaximum(value INSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximum:"), value)
}

// Determines whether the receiver will use heuristics to guess at the number
// which is intended by a string.
//
// # Discussion
//
// If the formatter is set to be lenient, as with any guessing it may get the
// result number wrong (that is, a number other than that which was intended).
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n NumberFormatter) Lenient() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isLenient"))
	return rv
}
func (n NumberFormatter) SetLenient(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setLenient:"), value)
}

// Determines whether partial string validation is enabled for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n NumberFormatter) PartialStringValidationEnabled() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isPartialStringValidationEnabled"))
	return rv
}
func (n NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setPartialStringValidationEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n NumberFormatter) MinimumGroupingDigits() int {
	rv := objc.Send[int](n.ID, objc.Sel("minimumGroupingDigits"))
	return rv
}
func (n NumberFormatter) SetMinimumGroupingDigits(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumGroupingDigits:"), value)
}
