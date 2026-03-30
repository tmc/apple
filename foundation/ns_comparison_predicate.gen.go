// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSComparisonPredicate] class.
var (
	_NSComparisonPredicateClass     NSComparisonPredicateClass
	_NSComparisonPredicateClassOnce sync.Once
)

func getNSComparisonPredicateClass() NSComparisonPredicateClass {
	_NSComparisonPredicateClassOnce.Do(func() {
		_NSComparisonPredicateClass = NSComparisonPredicateClass{class: objc.GetClass("NSComparisonPredicate")}
	})
	return _NSComparisonPredicateClass
}

// GetNSComparisonPredicateClass returns the class object for NSComparisonPredicate.
func GetNSComparisonPredicateClass() NSComparisonPredicateClass {
	return getNSComparisonPredicateClass()
}

type NSComparisonPredicateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSComparisonPredicateClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSComparisonPredicateClass) Alloc() NSComparisonPredicate {
	rv := objc.Send[NSComparisonPredicate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specialized predicate for comparing expressions.
//
// # Overview
//
// Use comparison predicates to compare the results of two expressions. You
// create a comparison predicate with an operator, a left expression, and a
// right expression, and use instances of the [NSExpression] class to
// represent those expressions. When you evaluate the predicate, it returns a
// [BOOL] value as the result of invoking the operator with the results of
// evaluating the expressions.
//
// # Creating Comparison Predicates
//
//   - [NSComparisonPredicate.InitWithLeftExpressionRightExpressionCustomSelector]: Creates a predicate that you form by combining specified left and right expressions using a specified selector.
//   - [NSComparisonPredicate.InitWithLeftExpressionRightExpressionModifierTypeOptions]: Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
//
// # Getting Information About a Comparison Predicate
//
//   - [NSComparisonPredicate.ComparisonPredicateModifier]: The comparison predicate modifier for the receiver.
//   - [NSComparisonPredicate.CustomSelector]: The selector for the receiver.
//   - [NSComparisonPredicate.RightExpression]: The right expression for the receiver.
//   - [NSComparisonPredicate.LeftExpression]: The left expression for the receiver.
//   - [NSComparisonPredicate.Options]: The options to use for the receiver.
//   - [NSComparisonPredicate.PredicateOperatorType]: The predicate type for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate
type NSComparisonPredicate struct {
	NSPredicate
}

// NSComparisonPredicateFromID constructs a [NSComparisonPredicate] from an objc.ID.
//
// A specialized predicate for comparing expressions.
func NSComparisonPredicateFromID(id objc.ID) NSComparisonPredicate {
	return NSComparisonPredicate{NSPredicate: NSPredicateFromID(id)}
}

// NOTE: NSComparisonPredicate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSComparisonPredicate] class.
//
// # Creating Comparison Predicates
//
//   - [INSComparisonPredicate.InitWithLeftExpressionRightExpressionCustomSelector]: Creates a predicate that you form by combining specified left and right expressions using a specified selector.
//   - [INSComparisonPredicate.InitWithLeftExpressionRightExpressionModifierTypeOptions]: Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
//
// # Getting Information About a Comparison Predicate
//
//   - [INSComparisonPredicate.ComparisonPredicateModifier]: The comparison predicate modifier for the receiver.
//   - [INSComparisonPredicate.CustomSelector]: The selector for the receiver.
//   - [INSComparisonPredicate.RightExpression]: The right expression for the receiver.
//   - [INSComparisonPredicate.LeftExpression]: The left expression for the receiver.
//   - [INSComparisonPredicate.Options]: The options to use for the receiver.
//   - [INSComparisonPredicate.PredicateOperatorType]: The predicate type for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate
type INSComparisonPredicate interface {
	INSPredicate

	// Topic: Creating Comparison Predicates

	// Creates a predicate that you form by combining specified left and right expressions using a specified selector.
	InitWithLeftExpressionRightExpressionCustomSelector(lhs INSExpression, rhs INSExpression, selector objc.SEL) NSComparisonPredicate
	// Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
	InitWithLeftExpressionRightExpressionModifierTypeOptions(lhs INSExpression, rhs INSExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) NSComparisonPredicate

	// Topic: Getting Information About a Comparison Predicate

	// The comparison predicate modifier for the receiver.
	ComparisonPredicateModifier() NSComparisonPredicateModifier
	// The selector for the receiver.
	CustomSelector() objc.SEL
	// The right expression for the receiver.
	RightExpression() INSExpression
	// The left expression for the receiver.
	LeftExpression() INSExpression
	// The options to use for the receiver.
	Options() NSComparisonPredicateOptions
	// The predicate type for the receiver.
	PredicateOperatorType() NSPredicateOperatorType
}

// Init initializes the instance.
func (c NSComparisonPredicate) Init() NSComparisonPredicate {
	rv := objc.Send[NSComparisonPredicate](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSComparisonPredicate) Autorelease() NSComparisonPredicate {
	rv := objc.Send[NSComparisonPredicate](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSComparisonPredicate creates a new NSComparisonPredicate instance.
func NewNSComparisonPredicate() NSComparisonPredicate {
	class := getNSComparisonPredicateClass()
	rv := objc.Send[NSComparisonPredicate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a predicate by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(coder:)
func NewComparisonPredicateWithCoder(coder INSCoder) NSComparisonPredicate {
	instance := getNSComparisonPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSComparisonPredicateFromID(rv)
}

// Creates a predicate that you form by combining specified left and right
// expressions using a specified selector.
//
// lhs: The left hand expression.
//
// rhs: The right hand expression.
//
// selector: The selector to use. The method defined by the selector must take a single
// argument and return a [BOOL] value.
//
// # Return Value
//
// The receiver, initialized by combining the left and right expressions using
// `selector`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:customSelector:)
func NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector(lhs INSExpression, rhs INSExpression, selector objc.SEL) NSComparisonPredicate {
	instance := getNSComparisonPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	return NSComparisonPredicateFromID(rv)
}

// Creates a predicate to a specified type that you form by combining
// specified left and right expressions using a specified modifier and
// options.
//
// lhs: The left hand expression.
//
// rhs: The right hand expression.
//
// modifier: The modifier to apply.
//
// type: The predicate operator type.
//
// options: The options to apply (see [NSComparisonPredicate.Options]). For no options,
// pass `0`.
//
// # Return Value
//
// The receiver, initialized to a predicate of type `type` formed by combining
// the left and right expressions using the `modifier` and `options`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:modifier:type:options:)
//
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
func NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs INSExpression, rhs INSExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) NSComparisonPredicate {
	instance := getNSComparisonPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	return NSComparisonPredicateFromID(rv)
}

// Creates a predicate that you form by combining specified left and right
// expressions using a specified selector.
//
// lhs: The left hand expression.
//
// rhs: The right hand expression.
//
// selector: The selector to use. The method defined by the selector must take a single
// argument and return a [BOOL] value.
//
// # Return Value
//
// The receiver, initialized by combining the left and right expressions using
// `selector`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:customSelector:)
func (c NSComparisonPredicate) InitWithLeftExpressionRightExpressionCustomSelector(lhs INSExpression, rhs INSExpression, selector objc.SEL) NSComparisonPredicate {
	rv := objc.Send[NSComparisonPredicate](c.ID, objc.Sel("initWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	return rv
}

// Creates a predicate to a specified type that you form by combining
// specified left and right expressions using a specified modifier and
// options.
//
// lhs: The left hand expression.
//
// rhs: The right hand expression.
//
// modifier: The modifier to apply.
//
// type: The predicate operator type.
//
// options: The options to apply (see [NSComparisonPredicate.Options]). For no options,
// pass `0`.
//
// # Return Value
//
// The receiver, initialized to a predicate of type `type` formed by combining
// the left and right expressions using the `modifier` and `options`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:modifier:type:options:)
//
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
func (c NSComparisonPredicate) InitWithLeftExpressionRightExpressionModifierTypeOptions(lhs INSExpression, rhs INSExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) NSComparisonPredicate {
	rv := objc.Send[NSComparisonPredicate](c.ID, objc.Sel("initWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	return rv
}

// Returns a new predicate formed by combining the left and right expressions
// using a given selector.
//
// lhs: The left hand side expression.
//
// rhs: The right hand side expression.
//
// selector: The selector to use for comparison. The method defined by the selector must
// take a single argument and return a [BOOL] value.
//
// # Return Value
//
// A new predicate formed by combining the left and right expressions using
// `selector`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:customSelector:
func (_NSComparisonPredicateClass NSComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionCustomSelector(lhs INSExpression, rhs INSExpression, selector objc.SEL) NSComparisonPredicate {
	rv := objc.Send[objc.ID](objc.ID(_NSComparisonPredicateClass.class), objc.Sel("predicateWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	return NSComparisonPredicateFromID(rv)
}

// Creates and returns a predicate of a given type formed by combining given
// left and right expressions using a given modifier and options.
//
// lhs: The left hand expression.
//
// rhs: The right hand expression.
//
// modifier: The modifier to apply.
//
// type: The predicate operator type.
//
// options: The options to apply (see [NSComparisonPredicate.Options]). For no options,
// pass `0`.
//
// # Return Value
//
// A new predicate of type `type` formed by combining the given left and right
// expressions using the `modifier` and `options`.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:modifier:type:options:
//
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
func (_NSComparisonPredicateClass NSComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs INSExpression, rhs INSExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) NSComparisonPredicate {
	rv := objc.Send[objc.ID](objc.ID(_NSComparisonPredicateClass.class), objc.Sel("predicateWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	return NSComparisonPredicateFromID(rv)
}

// The comparison predicate modifier for the receiver.
//
// # Discussion
//
// The default value is [NSDirectPredicateModifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/comparisonPredicateModifier
func (c NSComparisonPredicate) ComparisonPredicateModifier() NSComparisonPredicateModifier {
	rv := objc.Send[NSComparisonPredicateModifier](c.ID, objc.Sel("comparisonPredicateModifier"))
	return NSComparisonPredicateModifier(rv)
}

// The selector for the receiver.
//
// # Discussion
//
// [NULL] if there is none.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/customSelector
func (c NSComparisonPredicate) CustomSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("customSelector"))
	return rv
}

// The right expression for the receiver.
//
// # Discussion
//
// `nil` if there is none.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/rightExpression
func (c NSComparisonPredicate) RightExpression() INSExpression {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rightExpression"))
	return NSExpressionFromID(objc.ID(rv))
}

// The left expression for the receiver.
//
// # Discussion
//
// `nil` if there is none.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/leftExpression
func (c NSComparisonPredicate) LeftExpression() INSExpression {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("leftExpression"))
	return NSExpressionFromID(objc.ID(rv))
}

// The options to use for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/options-swift.property
func (c NSComparisonPredicate) Options() NSComparisonPredicateOptions {
	rv := objc.Send[NSComparisonPredicateOptions](c.ID, objc.Sel("options"))
	return NSComparisonPredicateOptions(rv)
}

// The predicate type for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateOperatorType
func (c NSComparisonPredicate) PredicateOperatorType() NSPredicateOperatorType {
	rv := objc.Send[NSPredicateOperatorType](c.ID, objc.Sel("predicateOperatorType"))
	return NSPredicateOperatorType(rv)
}
