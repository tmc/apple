// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that declares three methods that control the discretionary aspects of working with decimal numbers.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors
type NSDecimalNumberBehaviors interface {
	objectivec.IObject

	// Returns the way that [NSDecimalNumber]’s `decimalNumberBy...` methods round their return values.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/roundingMode()
	RoundingMode() NSRoundingMode

	// Returns the number of digits allowed after the decimal separator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/scale()
	Scale() int16

	// Specifies what an [NSDecimalNumber] object will do when it encounters an error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/exceptionDuringOperation(_:error:leftOperand:rightOperand:)
	ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.SEL, error_ NSCalculationError, leftOperand INSDecimalNumber, rightOperand INSDecimalNumber) INSDecimalNumber
}

// NSDecimalNumberBehaviorsObject wraps an existing Objective-C object that conforms to the NSDecimalNumberBehaviors protocol.
type NSDecimalNumberBehaviorsObject struct {
	objectivec.Object
}

func (o NSDecimalNumberBehaviorsObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDecimalNumberBehaviorsObjectFromID constructs a [NSDecimalNumberBehaviorsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDecimalNumberBehaviorsObjectFromID(id objc.ID) NSDecimalNumberBehaviorsObject {
	return NSDecimalNumberBehaviorsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the way that [NSDecimalNumber]’s `decimalNumberBy...` methods
// round their return values.
//
// # Return Value
//
// Returns the current rounding mode. See [NSDecimalNumber.RoundingMode] for
// possible values.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/roundingMode()
//
// [NSDecimalNumber.RoundingMode]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode
func (o NSDecimalNumberBehaviorsObject) RoundingMode() NSRoundingMode {
	rv := objc.Send[NSRoundingMode](o.ID, objc.Sel("roundingMode"))
	return rv
}

// Returns the number of digits allowed after the decimal separator.
//
// # Return Value
//
// The number of digits allowed after the decimal separator.
//
// # Discussion
//
// This method limits the precision of the values returned by
// [NSDecimalNumber]’s `decimalNumberBy...` methods. If [Scale] returns a
// negative value, it affects the digits before the decimal separator as well.
// If [Scale] returns [NSDecimalNoScale], the number of digits is unlimited.
//
// Assuming that [RoundingMode] returns [NSRoundPlain], different values of
// [Scale] have the following effects on the number 123.456:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/scale()
func (o NSDecimalNumberBehaviorsObject) Scale() int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("scale"))
	return rv
}

// Specifies what an [NSDecimalNumber] object will do when it encounters an
// error.
//
// operation: The method that was being executed when the error occurred.
//
// error: The type of error that was generated.
//
// leftOperand: The left operand.
//
// rightOperand: The right operand.
//
// # Discussion
//
// There are four possible values for `error`, described in
// [NSDecimalNumber.CalculationError]. The first three have to do with limits
// on the ability of [NSDecimalNumber] to represent decimal numbers. An
// [NSDecimalNumber] object can represent any number that can be expressed as
// mantissa x 10^exponent, where mantissa is a decimal integer up to 38 digits
// long, and exponent is between –256 and 256. The fourth results from the
// caller trying to divide by `0`.
//
// In implementing [ExceptionDuringOperationErrorLeftOperandRightOperand], you
// can handle each of these errors in several ways:
//
// - Raise an exception. For an explanation of exceptions, see [Exception
// Programming Topics]. - Return `nil`. The calling method will return its
// value as though no error had occurred. If `error` is
// [NSCalculationLossOfPrecision], `operation` will return an imprecise
// value—that is, one constrained to 38 significant digits. If `error` is
// [NSCalculationUnderflow] or [NSCalculationOverflow], `operation` will
// return [NSDecimalNumber]‘s `notANumber`. You shouldn’t return `nil` if
// `error` is [NSDivideByZero]. - Correct the error and return a valid
// [NSDecimalNumber] object. The calling method will use this as its own
// return value.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/exceptionDuringOperation(_:error:leftOperand:rightOperand:)
//
// [Exception Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Exceptions/Exceptions.html#//apple_ref/doc/uid/10000012i
// [NSDecimalNumber.CalculationError]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError
func (o NSDecimalNumberBehaviorsObject) ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.SEL, error_ NSCalculationError, leftOperand INSDecimalNumber, rightOperand INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("exceptionDuringOperation:error:leftOperand:rightOperand:"), operation, error_, leftOperand, rightOperand)
	return NSDecimalNumberFromID(rv)
}
