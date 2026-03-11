// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDecimalNumber] class.
var (
	_NSDecimalNumberClass     NSDecimalNumberClass
	_NSDecimalNumberClassOnce sync.Once
)

func getNSDecimalNumberClass() NSDecimalNumberClass {
	_NSDecimalNumberClassOnce.Do(func() {
		_NSDecimalNumberClass = NSDecimalNumberClass{class: objc.GetClass("NSDecimalNumber")}
	})
	return _NSDecimalNumberClass
}

// GetNSDecimalNumberClass returns the class object for NSDecimalNumber.
func GetNSDecimalNumberClass() NSDecimalNumberClass {
	return getNSDecimalNumberClass()
}

type NSDecimalNumberClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDecimalNumberClass) Alloc() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object for representing and performing arithmetic on base-10 numbers.
//
// # Overview
// 
// In Swift, this object bridges to [Decimal]; use [NSDecimalNumber] when you
// need reference semantics or other Foundation-specific behavior.
// 
// [NSDecimalNumber], an immutable subclass of [NSNumber], provides an
// object-oriented wrapper for doing base-10 arithmetic. An instance can
// represent any number that can be expressed as `mantissa x 10^exponent`
// where mantissa is a decimal integer up to 38 digits long, and exponent is
// an integer from –128 through 127.
//
// [Decimal]: https://developer.apple.com/documentation/Foundation/Decimal
//
// # Initializing a Decimal Number
//
//   - [NSDecimalNumber.InitWithDecimal]: Initializes a decimal number to represent a given decimal.
//   - [NSDecimalNumber.InitWithMantissaExponentIsNegative]: Initializes a decimal number using the given mantissa, exponent, and sign.
//   - [NSDecimalNumber.InitWithString]: Initializes a decimal number so that its value is equivalent to that in a given numeric string.
//   - [NSDecimalNumber.InitWithStringLocale]: Initializes a decimal number so that its value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// # Performing Arithmetic
//
//   - [NSDecimalNumber.DecimalNumberByAdding]: Adds this number to another given number.
//   - [NSDecimalNumber.DecimalNumberBySubtracting]: Subtracts another given number from this one.
//   - [NSDecimalNumber.DecimalNumberByMultiplyingBy]: Multiplies the number by another given number.
//   - [NSDecimalNumber.DecimalNumberByDividingBy]: Divides the number by another given number.
//   - [NSDecimalNumber.DecimalNumberByRaisingToPower]: Raises the number to a given power.
//   - [NSDecimalNumber.DecimalNumberByMultiplyingByPowerOf10]: Multiplies the number by 10 raised to the given power.
//   - [NSDecimalNumber.DecimalNumberByAddingWithBehavior]: Adds this number to another given number using the specified behavior.
//   - [NSDecimalNumber.DecimalNumberBySubtractingWithBehavior]: Subtracts this a given number from this one using the specified behavior.
//   - [NSDecimalNumber.DecimalNumberByMultiplyingByWithBehavior]: Multiplies this number by another given number using the specified behavior.
//   - [NSDecimalNumber.DecimalNumberByDividingByWithBehavior]: Divides this number by another given number using the specified behavior.
//   - [NSDecimalNumber.DecimalNumberByRaisingToPowerWithBehavior]: Raises the number to a given power using the specified behavior.
//   - [NSDecimalNumber.DecimalNumberByMultiplyingByPowerOf10WithBehavior]: Multiplies the number by 10 raised to the given power using the specified behavior.
//
// # Rounding Off
//
//   - [NSDecimalNumber.DecimalNumberByRoundingAccordingToBehavior]: Returns a rounded version of the decimal number using the specified rounding behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber
type NSDecimalNumber struct {
	NSNumber
}

// NSDecimalNumberFromID constructs a [NSDecimalNumber] from an objc.ID.
//
// An object for representing and performing arithmetic on base-10 numbers.
func NSDecimalNumberFromID(id objc.ID) NSDecimalNumber {
	return NSDecimalNumber{NSNumber: NSNumberFromID(id)}
}
// NOTE: NSDecimalNumber adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDecimalNumber] class.
//
// # Initializing a Decimal Number
//
//   - [INSDecimalNumber.InitWithDecimal]: Initializes a decimal number to represent a given decimal.
//   - [INSDecimalNumber.InitWithMantissaExponentIsNegative]: Initializes a decimal number using the given mantissa, exponent, and sign.
//   - [INSDecimalNumber.InitWithString]: Initializes a decimal number so that its value is equivalent to that in a given numeric string.
//   - [INSDecimalNumber.InitWithStringLocale]: Initializes a decimal number so that its value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// # Performing Arithmetic
//
//   - [INSDecimalNumber.DecimalNumberByAdding]: Adds this number to another given number.
//   - [INSDecimalNumber.DecimalNumberBySubtracting]: Subtracts another given number from this one.
//   - [INSDecimalNumber.DecimalNumberByMultiplyingBy]: Multiplies the number by another given number.
//   - [INSDecimalNumber.DecimalNumberByDividingBy]: Divides the number by another given number.
//   - [INSDecimalNumber.DecimalNumberByRaisingToPower]: Raises the number to a given power.
//   - [INSDecimalNumber.DecimalNumberByMultiplyingByPowerOf10]: Multiplies the number by 10 raised to the given power.
//   - [INSDecimalNumber.DecimalNumberByAddingWithBehavior]: Adds this number to another given number using the specified behavior.
//   - [INSDecimalNumber.DecimalNumberBySubtractingWithBehavior]: Subtracts this a given number from this one using the specified behavior.
//   - [INSDecimalNumber.DecimalNumberByMultiplyingByWithBehavior]: Multiplies this number by another given number using the specified behavior.
//   - [INSDecimalNumber.DecimalNumberByDividingByWithBehavior]: Divides this number by another given number using the specified behavior.
//   - [INSDecimalNumber.DecimalNumberByRaisingToPowerWithBehavior]: Raises the number to a given power using the specified behavior.
//   - [INSDecimalNumber.DecimalNumberByMultiplyingByPowerOf10WithBehavior]: Multiplies the number by 10 raised to the given power using the specified behavior.
//
// # Rounding Off
//
//   - [INSDecimalNumber.DecimalNumberByRoundingAccordingToBehavior]: Returns a rounded version of the decimal number using the specified rounding behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber
type INSDecimalNumber interface {
	INSNumber

	// Topic: Initializing a Decimal Number

	// Initializes a decimal number to represent a given decimal.
	InitWithDecimal(dcm NSDecimal) NSDecimalNumber
	// Initializes a decimal number using the given mantissa, exponent, and sign.
	InitWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) NSDecimalNumber
	// Initializes a decimal number so that its value is equivalent to that in a given numeric string.
	InitWithString(numberValue string) NSDecimalNumber
	// Initializes a decimal number so that its value is equivalent to that in a given numeric string, interpreted using a given locale.
	InitWithStringLocale(numberValue string, locale objectivec.IObject) NSDecimalNumber

	// Topic: Performing Arithmetic

	// Adds this number to another given number.
	DecimalNumberByAdding(decimalNumber INSDecimalNumber) INSDecimalNumber
	// Subtracts another given number from this one.
	DecimalNumberBySubtracting(decimalNumber INSDecimalNumber) INSDecimalNumber
	// Multiplies the number by another given number.
	DecimalNumberByMultiplyingBy(decimalNumber INSDecimalNumber) INSDecimalNumber
	// Divides the number by another given number.
	DecimalNumberByDividingBy(decimalNumber INSDecimalNumber) INSDecimalNumber
	// Raises the number to a given power.
	DecimalNumberByRaisingToPower(power uint) INSDecimalNumber
	// Multiplies the number by 10 raised to the given power.
	DecimalNumberByMultiplyingByPowerOf10(power int16) INSDecimalNumber
	// Adds this number to another given number using the specified behavior.
	DecimalNumberByAddingWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber
	// Subtracts this a given number from this one using the specified behavior.
	DecimalNumberBySubtractingWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber
	// Multiplies this number by another given number using the specified behavior.
	DecimalNumberByMultiplyingByWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber
	// Divides this number by another given number using the specified behavior.
	DecimalNumberByDividingByWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber
	// Raises the number to a given power using the specified behavior.
	DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior NSDecimalNumberBehaviors) INSDecimalNumber
	// Multiplies the number by 10 raised to the given power using the specified behavior.
	DecimalNumberByMultiplyingByPowerOf10WithBehavior(power int16, behavior NSDecimalNumberBehaviors) INSDecimalNumber

	// Topic: Rounding Off

	// Returns a rounded version of the decimal number using the specified rounding behavior.
	DecimalNumberByRoundingAccordingToBehavior(behavior NSDecimalNumberBehaviors) INSDecimalNumber
}





// Init initializes the instance.
func (d NSDecimalNumber) Init() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDecimalNumber) Autorelease() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDecimalNumber creates a new NSDecimalNumber instance.
func NewNSDecimalNumber() NSDecimalNumber {
	class := getNSDecimalNumberClass()
	rv := objc.Send[NSDecimalNumber](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(GCPoint2:)
// point is a [gamecontroller.GCPoint2].
func NewDecimalNumberValueWithGCPoint2(point objectivec.IObject) NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(getNSDecimalNumberClass().class), objc.Sel("valueWithGCPoint2:"), point)
	return NSDecimalNumberFromID(rv)
}


// Initializes a value object to contain the specified value, interpreted with
// the specified Objective-C type.
//
// value: A pointer to data to be stored in the new value object.
//
// type: The Objective-C type of `value`, as provided by the `@encode()` compiler
// directive. Do not hard-code this parameter as a C string.
//
// # Return Value
// 
// An initialized value object that contains `value`, which is interpreted as
// being of the Objective-C type `type`. The returned object might be
// different than the original receiver.
//
// # Discussion
// 
// See [Number and Value Programming Topics] for other considerations in
// creating a value object.
// 
// This is the designated initializer for the [NSValue] class.
//
// [Number and Value Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/NumbersandValues/NumbersandValues.html#//apple_ref/doc/uid/10000038i
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(bytes:objCType:)
func NewDecimalNumberWithBytesObjCType(value unsafe.Pointer, type_ []byte) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:objCType:"), value, unsafe.Pointer(unsafe.SliceData(type_)))
	return NSDecimalNumberFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(coder:)
func NewDecimalNumberWithCoder(coder INSCoder) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDecimalNumberFromID(rv)
}


// Initializes a decimal number to represent a given decimal.
//
// dcm: The value of the new object.
//
// # Return Value
// 
// An [NSDecimalNumber] object initialized to represent `dcm`.
//
// # Discussion
// 
// This method is the designated initializer for [NSDecimalNumber].
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(decimal:)
func NewDecimalNumberWithDecimal(dcm NSDecimal) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDecimal:"), dcm)
	return NSDecimalNumberFromID(rv)
}


// Initializes a decimal number using the given mantissa, exponent, and sign.
//
// mantissa: The mantissa for the new decimal number object.
//
// exponent: The exponent for the new decimal number object.
//
// flag: A Boolean value that specifies whether the sign of the number is negative.
//
// # Return Value
// 
// An [NSDecimalNumber] object initialized using the given mantissa, exponent,
// and sign.
//
// # Discussion
// 
// The arguments express a number in a type of scientific notation that
// requires the mantissa to be an integer. So, for example, if the number to
// be represented is 1.23, it is expressed as 123x10^–2—`mantissa` is 123;
// `exponent` is –2; and `isNegative`, which refers to the sign of the
// mantissa, is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(mantissa:exponent:isNegative:)
func NewDecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return NSDecimalNumberFromID(rv)
}


// Initializes a decimal number so that its value is equivalent to that in a
// given numeric string.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number. For a listing of acceptable and unacceptable
// strings, see [InitWithStringLocale].
//
// # Discussion
// 
// Don’t use this initializer if `numberValue` has a fractional part, since
// the lack of a locale makes handling the decimal separator ambiguous. The
// separator is a period in some locales (like in the United States) and a
// comma in others (such as France).
// 
// To parse a numeric string with a fractional part, use
// [InitWithStringLocale] instead. When working with numeric representations
// with a known format, pass a fixed locale to ensure consistent results
// independent of the user’s current device settings. For localized parsing
// that uses the user’s current device settings, pass [CurrentLocale].
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:)
func NewDecimalNumberWithString(numberValue string) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(numberValue))
	return NSDecimalNumberFromID(rv)
}


// Initializes a decimal number so that its value is equivalent to that in a
// given numeric string, interpreted using a given locale.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number.
//
// locale: A dictionary that defines the locale (specifically the [decimalSeparator])
// to use to interpret the number in `numberValue`.
// //
// [decimalSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
//
// # Discussion
// 
// The locale parameter determines whether the `decimalSeparator` is a period
// (like in the United States) or a comma (like in France).
// 
// The following strings show examples of acceptable values for `numberValue`:
// 
// - `2500.6` (or `2500,6`, depending on locale) - `–2500.6` (or
// `–2500,6`) - `–2.5006e3` (or `–2,5006e3`) - `–2.5006E3` (or
// `–2,5006E3`)
// 
// The following strings are unacceptable:
// 
// - `2,500.6` - `2500 3/5` - `2.5006x10e3` - `two thousand five hundred and
// six tenths`
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:locale:)
func NewDecimalNumberWithStringLocale(numberValue string, locale objectivec.IObject) NSDecimalNumber {
	instance := getNSDecimalNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:locale:"), objc.String(numberValue), locale)
	return NSDecimalNumberFromID(rv)
}







// Initializes a decimal number to represent a given decimal.
//
// dcm: The value of the new object.
//
// # Return Value
// 
// An [NSDecimalNumber] object initialized to represent `dcm`.
//
// # Discussion
// 
// This method is the designated initializer for [NSDecimalNumber].
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(decimal:)
func (d NSDecimalNumber) InitWithDecimal(dcm NSDecimal) NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("initWithDecimal:"), dcm)
	return rv
}

// Initializes a decimal number using the given mantissa, exponent, and sign.
//
// mantissa: The mantissa for the new decimal number object.
//
// exponent: The exponent for the new decimal number object.
//
// flag: A Boolean value that specifies whether the sign of the number is negative.
//
// # Return Value
// 
// An [NSDecimalNumber] object initialized using the given mantissa, exponent,
// and sign.
//
// # Discussion
// 
// The arguments express a number in a type of scientific notation that
// requires the mantissa to be an integer. So, for example, if the number to
// be represented is 1.23, it is expressed as 123x10^–2—`mantissa` is 123;
// `exponent` is –2; and `isNegative`, which refers to the sign of the
// mantissa, is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(mantissa:exponent:isNegative:)
func (d NSDecimalNumber) InitWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

// Initializes a decimal number so that its value is equivalent to that in a
// given numeric string.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number. For a listing of acceptable and unacceptable
// strings, see [InitWithStringLocale].
//
// # Discussion
// 
// Don’t use this initializer if `numberValue` has a fractional part, since
// the lack of a locale makes handling the decimal separator ambiguous. The
// separator is a period in some locales (like in the United States) and a
// comma in others (such as France).
// 
// To parse a numeric string with a fractional part, use
// [InitWithStringLocale] instead. When working with numeric representations
// with a known format, pass a fixed locale to ensure consistent results
// independent of the user’s current device settings. For localized parsing
// that uses the user’s current device settings, pass [CurrentLocale].
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:)
func (d NSDecimalNumber) InitWithString(numberValue string) NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("initWithString:"), objc.String(numberValue))
	return rv
}

// Initializes a decimal number so that its value is equivalent to that in a
// given numeric string, interpreted using a given locale.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number.
//
// locale: A dictionary that defines the locale (specifically the [decimalSeparator])
// to use to interpret the number in `numberValue`.
// //
// [decimalSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
//
// # Discussion
// 
// The locale parameter determines whether the `decimalSeparator` is a period
// (like in the United States) or a comma (like in France).
// 
// The following strings show examples of acceptable values for `numberValue`:
// 
// - `2500.6` (or `2500,6`, depending on locale) - `–2500.6` (or
// `–2500,6`) - `–2.5006e3` (or `–2,5006e3`) - `–2.5006E3` (or
// `–2,5006E3`)
// 
// The following strings are unacceptable:
// 
// - `2,500.6` - `2500 3/5` - `2.5006x10e3` - `two thousand five hundred and
// six tenths`
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:locale:)
func (d NSDecimalNumber) InitWithStringLocale(numberValue string, locale objectivec.IObject) NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d.ID, objc.Sel("initWithString:locale:"), objc.String(numberValue), locale)
	return rv
}

// Adds this number to another given number.
//
// decimalNumber: The number to add to the receiver.
//
// # Return Value
// 
// A new [NSDecimalNumber] object whose value is the sum of the receiver and
// `decimalNumber`.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:)
func (d NSDecimalNumber) DecimalNumberByAdding(decimalNumber INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByAdding:"), decimalNumber)
	return NSDecimalNumberFromID(rv)
}

// Subtracts another given number from this one.
//
// decimalNumber: The number to subtract from the receiver.
//
// # Return Value
// 
// A new [NSDecimalNumber] object whose value is `decimalNumber` subtracted
// from the receiver.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// when rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:)
func (d NSDecimalNumber) DecimalNumberBySubtracting(decimalNumber INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberBySubtracting:"), decimalNumber)
	return NSDecimalNumberFromID(rv)
}

// Multiplies the number by another given number.
//
// decimalNumber: The number by which to multiply the receiver.
//
// # Return Value
// 
// A new [NSDecimalNumber] object whose value is `decimalNumber` multiplied by
// the receiver.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// when rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:)
func (d NSDecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByMultiplyingBy:"), decimalNumber)
	return NSDecimalNumberFromID(rv)
}

// Divides the number by another given number.
//
// decimalNumber: The number by which to divide the receiver.
//
// # Return Value
// 
// A new [NSDecimalNumber] object whose value is the value of the receiver
// divided by `decimalNumber`.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:)
func (d NSDecimalNumber) DecimalNumberByDividingBy(decimalNumber INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByDividingBy:"), decimalNumber)
	return NSDecimalNumberFromID(rv)
}

// Raises the number to a given power.
//
// power: The power to which to raise the receiver.
//
// # Return Value
// 
// A new [NSDecimalNumber] object whose value is the value of the receiver
// raised to the power `power`.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// when rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:)
func (d NSDecimalNumber) DecimalNumberByRaisingToPower(power uint) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByRaisingToPower:"), power)
	return NSDecimalNumberFromID(rv)
}

// Multiplies the number by 10 raised to the given power.
//
// # Discussion
// 
// This method uses the default behavior when handling calculation errors and
// when rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:)
func (d NSDecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power int16) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:"), power)
	return NSDecimalNumberFromID(rv)
}

// Adds this number to another given number using the specified behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:withBehavior:)
func (d NSDecimalNumber) DecimalNumberByAddingWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByAdding:withBehavior:"), decimalNumber, behavior)
	return NSDecimalNumberFromID(rv)
}

// Subtracts this a given number from this one using the specified behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:withBehavior:)
func (d NSDecimalNumber) DecimalNumberBySubtractingWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberBySubtracting:withBehavior:"), decimalNumber, behavior)
	return NSDecimalNumberFromID(rv)
}

// Multiplies this number by another given number using the specified
// behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:withBehavior:)
func (d NSDecimalNumber) DecimalNumberByMultiplyingByWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByMultiplyingBy:withBehavior:"), decimalNumber, behavior)
	return NSDecimalNumberFromID(rv)
}

// Divides this number by another given number using the specified behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:withBehavior:)
func (d NSDecimalNumber) DecimalNumberByDividingByWithBehavior(decimalNumber INSDecimalNumber, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByDividingBy:withBehavior:"), decimalNumber, behavior)
	return NSDecimalNumberFromID(rv)
}

// Raises the number to a given power using the specified behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:withBehavior:)
func (d NSDecimalNumber) DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByRaisingToPower:withBehavior:"), power, behavior)
	return NSDecimalNumberFromID(rv)
}

// Multiplies the number by 10 raised to the given power using the specified
// behavior.
//
// # Discussion
// 
// `behavior` specifies the handling of calculation errors and rounding.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:withBehavior:)
func (d NSDecimalNumber) DecimalNumberByMultiplyingByPowerOf10WithBehavior(power int16, behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:withBehavior:"), power, behavior)
	return NSDecimalNumberFromID(rv)
}

// Returns a rounded version of the decimal number using the specified
// rounding behavior.
//
// # Discussion
// 
// For a description of the different ways of rounding, see the [RoundingMode]
// method in the [NSDecimalNumberBehaviors] protocol specification.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/rounding(accordingToBehavior:)
func (d NSDecimalNumber) DecimalNumberByRoundingAccordingToBehavior(behavior NSDecimalNumberBehaviors) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decimalNumberByRoundingAccordingToBehavior:"), behavior)
	return NSDecimalNumberFromID(rv)
}





// Creates and returns a decimal number equivalent to a given decimal
// structure.
//
// dcm: An [NSDecimal] structure that specifies the value for the new decimal
// number object.
//
// # Return Value
// 
// An [NSDecimalNumber] object equivalent to `dcm`.
//
// # Discussion
// 
// You can initialize `dcm` programmatically or generate it using the
// [NSScanner] method, [scanDecimal(_:)]
//
// [scanDecimal(_:)]: https://developer.apple.com/documentation/Foundation/Scanner/scanDecimal(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithDecimal:
func (_NSDecimalNumberClass NSDecimalNumberClass) DecimalNumberWithDecimal(dcm NSDecimal) NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("decimalNumberWithDecimal:"), dcm)
	return NSDecimalNumberFromID(rv)
}

// Creates and returns a decimal number equivalent to the number specified by
// the arguments.
//
// mantissa: The mantissa for the new decimal number object.
//
// exponent: The exponent for the new decimal number object.
//
// flag: A Boolean value that specifies whether the sign of the number is negative.
//
// # Discussion
// 
// The arguments express a number in a kind of scientific notation that
// requires the mantissa to be an integer. So, for example, if the number to
// be represented is `–12.345`, it is expressed as
// `12345x10^–3`—`mantissa` is `12345`; `exponent` is `–3`; and `flag`
// is [true], as illustrated by the following example.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithMantissa:exponent:isNegative:
func (_NSDecimalNumberClass NSDecimalNumberClass) DecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return NSDecimalNumberFromID(rv)
}

// Creates a decimal number whose value is equivalent to that in a given
// numeric string.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number. For a listing of acceptable and unacceptable
// strings, see [DecimalNumberWithStringLocale].
//
// # Discussion
// 
// Don’t use this method if `numberValue` has a fractional part, because the
// lack of a locale makes handling the decimal separator ambiguous. The
// separator is a period in some locales (like in the United States) and a
// comma in others (such as France).
// 
// To parse a numeric string with a fractional part, use
// [DecimalNumberWithStringLocale] instead. When working with numeric
// representations with a known format, pass a fixed locale to ensure
// consistent results independent of the user’s current device settings. For
// localized parsing that uses the user’s current device settings, pass
// [CurrentLocale].
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:
func (_NSDecimalNumberClass NSDecimalNumberClass) DecimalNumberWithString(numberValue string) NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("decimalNumberWithString:"), objc.String(numberValue))
	return NSDecimalNumberFromID(rv)
}

// Creates a decimal number whose value is equivalent to that in a given
// numeric string, interpreted using a given locale.
//
// numberValue: A numeric string.
// 
// Besides digits, `numberValue` can include an initial `+` or `–`; a single
// [E] or `e`, to indicate the exponent of a number in scientific notation;
// and a single decimal separator character to divide the fractional from the
// integral part of the number.
//
// locale: A dictionary that defines the locale (specifically the [decimalSeparator])
// to use to interpret the number in `numberValue`.
// //
// [decimalSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
//
// # Discussion
// 
// The `locale` parameter determines whether the [decimalSeparator] is a
// period (like in the United States) or a comma (like in France).
// 
// The following strings show examples of acceptable values for `numberValue`:
// 
// - `2500.6` (or `2500,6`, depending on locale) - `–2500.6` (or
// `–2500,6`) - `–2.5006e3` (or `–2,5006e3`) - `–2.5006E3` (or
// `–2,5006E3`)
// 
// The following strings are unacceptable:
// 
// - `2,500.6` - `2500 3/5` - `2.5006x10e3` - `two thousand five hundred and
// six tenths`
//
// [decimalSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:locale:
func (_NSDecimalNumberClass NSDecimalNumberClass) DecimalNumberWithStringLocale(numberValue string, locale objectivec.IObject) NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("decimalNumberWithString:locale:"), objc.String(numberValue), locale)
	return NSDecimalNumberFromID(rv)
}












// A decimal number equivalent to the number 1.0.
//
// # Return Value
// 
// An [NSDecimalNumber] object equivalent to the number 1.0.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/one
func (_NSDecimalNumberClass NSDecimalNumberClass) One() NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("one"))
	return NSDecimalNumberFromID(objc.ID(rv))
}



// A decimal number equivalent to the number 0.0.
//
// # Return Value
// 
// An [NSDecimalNumber] object equivalent to the number 0.0.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/zero
func (_NSDecimalNumberClass NSDecimalNumberClass) Zero() NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("zero"))
	return NSDecimalNumberFromID(objc.ID(rv))
}



// A decimal number that specifies no number.
//
// # Return Value
// 
// An [NSDecimalNumber] object that specifies no number.
// 
// # Discussion
// 
// Any arithmetic method receiving [NotANumber] as an argument returns
// [NotANumber].
// 
// This value can be a useful way of handling non-numeric data in an input
// file. This method can also be a useful response to calculation errors. For
// more information on calculation errors, see the
// [ExceptionDuringOperationErrorLeftOperandRightOperand] method description
// in the [NSDecimalNumberBehaviors] protocol specification.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/notANumber
func (_NSDecimalNumberClass NSDecimalNumberClass) NotANumber() NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("notANumber"))
	return NSDecimalNumberFromID(objc.ID(rv))
}



// The way arithmetic methods round off and handle error conditions.
//
// # Discussion
// 
// By default, the arithmetic methods use the [NSRoundPlain] behavior; that
// is, the methods round to the closest possible return value. The methods
// assume your need for precision does not exceed 38 significant digits and
// raise exceptions when they try to divide by 0 or produce a number too big
// or too small to be represented.
// 
// If this default behavior doesn’t suit your application, you should use
// methods that let you specify the behavior, like
// [DecimalNumberByAddingWithBehavior]. If you find yourself using a
// particular behavior consistently, you can specify a different default
// behavior with `setDefaultBehavior(_:)`.
// 
// The default behavior is maintained separately for each thread in your app.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (_NSDecimalNumberClass NSDecimalNumberClass) DefaultBehavior() NSDecimalNumberBehaviors {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("defaultBehavior"))
	return NSDecimalNumberBehaviorsObjectFromID(rv)
}
func (_NSDecimalNumberClass NSDecimalNumberClass) SetDefaultBehavior(value NSDecimalNumberBehaviors) {
	objc.Send[struct{}](objc.ID(_NSDecimalNumberClass.class), objc.Sel("setDefaultBehavior:"), value)
}



// Returns the largest possible value of a decimal number.
//
// # Return Value
// 
// The largest possible value of an [NSDecimalNumber] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/maximum
func (_NSDecimalNumberClass NSDecimalNumberClass) MaximumDecimalNumber() NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("maximumDecimalNumber"))
	return NSDecimalNumberFromID(objc.ID(rv))
}



// Returns the smallest possible value of a decimal number.
//
// # Return Value
// 
// The smallest possible value of an [NSDecimalNumber] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/minimum
func (_NSDecimalNumberClass NSDecimalNumberClass) MinimumDecimalNumber() NSDecimalNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberClass.class), objc.Sel("minimumDecimalNumber"))
	return NSDecimalNumberFromID(objc.ID(rv))
}


























