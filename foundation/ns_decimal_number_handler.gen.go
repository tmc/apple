// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDecimalNumberHandler] class.
var (
	_NSDecimalNumberHandlerClass     NSDecimalNumberHandlerClass
	_NSDecimalNumberHandlerClassOnce sync.Once
)

func getNSDecimalNumberHandlerClass() NSDecimalNumberHandlerClass {
	_NSDecimalNumberHandlerClassOnce.Do(func() {
		_NSDecimalNumberHandlerClass = NSDecimalNumberHandlerClass{class: objc.GetClass("NSDecimalNumberHandler")}
	})
	return _NSDecimalNumberHandlerClass
}

// GetNSDecimalNumberHandlerClass returns the class object for NSDecimalNumberHandler.
func GetNSDecimalNumberHandlerClass() NSDecimalNumberHandlerClass {
	return getNSDecimalNumberHandlerClass()
}

type NSDecimalNumberHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDecimalNumberHandlerClass) Alloc() NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that adopts the decimal number behaviors protocol.
//
// # Overview
// 
// This class allows you to set the way an [NSDecimalNumber] object rounds off
// and handles errors, without having to create a custom class.
// 
// You can use an instance of this class as an argument to any of the
// [NSDecimalNumber] methods that end with `...Behavior:`. If you don’t
// think you need special behavior, you probably don’t need this class—it
// is likely that [NSDecimalNumber]’s default behavior will suit your needs.
// 
// For more information, see the [NSDecimalNumberBehaviors] protocol
// specification.
//
// # Initializing a decimal number handler
//
//   - [NSDecimalNumberHandler.InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero]: Returns an [NSDecimalNumberHandler] object initialized so it behaves as specified by the method’s arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler
type NSDecimalNumberHandler struct {
	objectivec.Object
}

// NSDecimalNumberHandlerFromID constructs a [NSDecimalNumberHandler] from an objc.ID.
//
// A class that adopts the decimal number behaviors protocol.
func NSDecimalNumberHandlerFromID(id objc.ID) NSDecimalNumberHandler {
	return NSDecimalNumberHandler{objectivec.Object{ID: id}}
}
// NOTE: NSDecimalNumberHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDecimalNumberHandler] class.
//
// # Initializing a decimal number handler
//
//   - [INSDecimalNumberHandler.InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero]: Returns an [NSDecimalNumberHandler] object initialized so it behaves as specified by the method’s arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler
type INSDecimalNumberHandler interface {
	objectivec.IObject
	NSCoding

	// Topic: Initializing a decimal number handler

	// Returns an [NSDecimalNumberHandler] object initialized so it behaves as specified by the method’s arguments.
	InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode NSRoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) NSDecimalNumberHandler

	// The rounding behavior used by the receiver.
	RoundingBehavior() INSDecimalNumberHandler
	SetRoundingBehavior(value INSDecimalNumberHandler)
	// The rounding increment used by the receiver.
	RoundingIncrement() INSNumber
	SetRoundingIncrement(value INSNumber)
	// Specifies what an [NSDecimalNumber] object will do when it encounters an error.
	ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.SEL, error_ NSCalculationError, leftOperand INSDecimalNumber, rightOperand INSDecimalNumber) INSDecimalNumber
	// Returns the number of digits allowed after the decimal separator.
	Scale() int16
}

// Init initializes the instance.
func (d NSDecimalNumberHandler) Init() NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDecimalNumberHandler) Autorelease() NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDecimalNumberHandler creates a new NSDecimalNumberHandler instance.
func NewNSDecimalNumberHandler() NSDecimalNumberHandler {
	class := getNSDecimalNumberHandlerClass()
	rv := objc.Send[NSDecimalNumberHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDecimalNumberHandlerWithCoder(coder INSCoder) NSDecimalNumberHandler {
	instance := getNSDecimalNumberHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDecimalNumberHandlerFromID(rv)
}

// Returns an [NSDecimalNumberHandler] object initialized so it behaves as
// specified by the method’s arguments.
//
// roundingMode: The rounding mode to use. There are four possible values: [NSRoundUp],
// [NSRoundDown], [NSRoundPlain], and [NSRoundBankers].
//
// scale: The number of digits a rounded value should have after its decimal point.
//
// exact: If [true], in the event of an exactness error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// overflow: If [true], in the event of an overflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// underflow: If [true], in the event of an underflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// divideByZero: If [true], in the event of a divide by zero error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized [NSDecimalNumberHandler] object initialized with customized
// behavior. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// See the [NSDecimalNumberBehaviors] protocol specification for a complete
// explanation of the possible behaviors.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler/init(roundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:)
func NewDecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode NSRoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) NSDecimalNumberHandler {
	instance := getNSDecimalNumberHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return NSDecimalNumberHandlerFromID(rv)
}

// Returns an [NSDecimalNumberHandler] object initialized so it behaves as
// specified by the method’s arguments.
//
// roundingMode: The rounding mode to use. There are four possible values: [NSRoundUp],
// [NSRoundDown], [NSRoundPlain], and [NSRoundBankers].
//
// scale: The number of digits a rounded value should have after its decimal point.
//
// exact: If [true], in the event of an exactness error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// overflow: If [true], in the event of an overflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// underflow: If [true], in the event of an underflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// divideByZero: If [true], in the event of a divide by zero error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized [NSDecimalNumberHandler] object initialized with customized
// behavior. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// See the [NSDecimalNumberBehaviors] protocol specification for a complete
// explanation of the possible behaviors.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler/init(roundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:)
func (d NSDecimalNumberHandler) InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode NSRoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](d.ID, objc.Sel("initWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSDecimalNumberHandler) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
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
// [Exception Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Exceptions/Exceptions.html#//apple_ref/doc/uid/10000012i
// [NSDecimalNumber.CalculationError]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberBehaviors/exceptionDuringOperation(_:error:leftOperand:rightOperand:)
func (d NSDecimalNumberHandler) ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.SEL, error_ NSCalculationError, leftOperand INSDecimalNumber, rightOperand INSDecimalNumber) INSDecimalNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("exceptionDuringOperation:error:leftOperand:rightOperand:"), operation, error_, leftOperand, rightOperand)
	return NSDecimalNumberFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (d NSDecimalNumberHandler) InitWithCoder(coder INSCoder) NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](d.ID, objc.Sel("initWithCoder:"), coder)
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
func (d NSDecimalNumberHandler) Scale() int16 {
	rv := objc.Send[int16](d.ID, objc.Sel("scale"))
	return rv
}

// Returns an [NSDecimalNumberHandler] object with customized behavior.
//
// roundingMode: The rounding mode to use. There are four possible values: [NSRoundUp],
// [NSRoundDown], [NSRoundPlain], and [NSRoundBankers].
//
// scale: The number of digits a rounded value should have after its decimal point.
//
// exact: If [true], in the event of an exactness error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// overflow: If [true], in the event of an overflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// underflow: If [true], in the event of an underflow error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// divideByZero: If [true], in the event of a divide by zero error the handler will raise an
// exception, otherwise it will ignore the error and return control to the
// calling method
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSDecimalNumberHandler] object with customized behavior.
//
// # Discussion
// 
// See the [NSDecimalNumberBehaviors] protocol specification for a complete
// explanation of the possible behaviors.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler/decimalNumberHandlerWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:
func (_NSDecimalNumberHandlerClass NSDecimalNumberHandlerClass) DecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode NSRoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) NSDecimalNumberHandler {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberHandlerClass.class), objc.Sel("decimalNumberHandlerWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return NSDecimalNumberHandlerFromID(rv)
}

// The rounding behavior used by the receiver.
//
// See: https://developer.apple.com/documentation/foundation/numberformatter/roundingbehavior
func (d NSDecimalNumberHandler) RoundingBehavior() INSDecimalNumberHandler {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("roundingBehavior"))
	return NSDecimalNumberHandlerFromID(objc.ID(rv))
}
func (d NSDecimalNumberHandler) SetRoundingBehavior(value INSDecimalNumberHandler) {
	objc.Send[struct{}](d.ID, objc.Sel("setRoundingBehavior:"), value)
}
// The rounding increment used by the receiver.
//
// See: https://developer.apple.com/documentation/foundation/numberformatter/roundingincrement
func (d NSDecimalNumberHandler) RoundingIncrement() INSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("roundingIncrement"))
	return NSNumberFromID(objc.ID(rv))
}
func (d NSDecimalNumberHandler) SetRoundingIncrement(value INSNumber) {
	objc.Send[struct{}](d.ID, objc.Sel("setRoundingIncrement:"), value)
}
		
		
// The rounding mode used by the receiver.
//
// See: https://developer.apple.com/documentation/foundation/numberformatter/roundingmode-swift.property
func (d NSDecimalNumberHandler) RoundingMode() NSNumberFormatterRoundingMode {
	rv := objc.Send[NSNumberFormatterRoundingMode](d.ID, objc.Sel("roundingMode"))
	return NSNumberFormatterRoundingMode(rv)
}
func (d NSDecimalNumberHandler) SetRoundingMode(value NSNumberFormatterRoundingMode) {
	objc.Send[struct{}](d.ID, objc.Sel("setRoundingMode:"), value)
}

// Returns the default instance of [NSDecimalNumberHandler].
//
// # Return Value
// 
// The default instance of [NSDecimalNumberHandler].
// 
// # Discussion
// 
// This default decimal number handler rounds to the closest possible return
// value. It assumes your need for precision does not exceed 38 significant
// digits, and it raises an exception when its [NSDecimalNumber] object tries
// to divide by `0` or when its [NSDecimalNumber] object produces a number too
// big or too small to be represented.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler/default
func (_NSDecimalNumberHandlerClass NSDecimalNumberHandlerClass) DefaultDecimalNumberHandler() NSDecimalNumberHandler {
	rv := objc.Send[objc.ID](objc.ID(_NSDecimalNumberHandlerClass.class), objc.Sel("defaultDecimalNumberHandler"))
	return NSDecimalNumberHandlerFromID(objc.ID(rv))
}
		

			// Protocol methods for NSCoding
			

