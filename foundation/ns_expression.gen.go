// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSExpression] class.
var (
	_NSExpressionClass     NSExpressionClass
	_NSExpressionClassOnce sync.Once
)

func getNSExpressionClass() NSExpressionClass {
	_NSExpressionClassOnce.Do(func() {
		_NSExpressionClass = NSExpressionClass{class: objc.GetClass("NSExpression")}
	})
	return _NSExpressionClass
}

// GetNSExpressionClass returns the class object for NSExpression.
func GetNSExpressionClass() NSExpressionClass {
	return getNSExpressionClass()
}

type NSExpressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExpressionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExpressionClass) Alloc() NSExpression {
	rv := objc.Send[NSExpression](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An expression for use in a comparison predicate.
//
// # Overview
// 
// Comparison operations in an [NSPredicate] derive from two expressions as
// instances of the [NSExpression] class. You create expressions for constant
// values, key paths, and so on.
// 
// Generally, anywhere in the [NSExpression] class hierarchy where there’s a
// composite API and subtypes that may only reasonably respond to a subset of
// that API, invoking a method that doesn’t make sense for that subtype
// throws an exception.
// 
// # Aggregate Expressions
// 
// [AggregateExpressionType] allows you to create predicates containing
// expressions that evaluate to collections that contain further expressions.
// The collection may be an [NSArray], [NSSet], or [NSDictionary] object.
// 
// Core Data doesn’t support aggregate expressions.
// 
// # Subquery Expressions
// 
// The [SubqueryExpressionType] creates a subexpression that returns a subset
// of a collection of objects. This allows you to create sophisticated queries
// across relationships, such as a search for multiple correlated values on
// the destination object of a relationship.
// 
// # Set Expressions
// 
// The set expressions ([UnionSetExpressionType],
// [IntersectSetExpressionType], and [MinusSetExpressionType]) combine results
// in a manner similar to the [NSSet] methods.
// 
// Both sides of these expressions must evaluate to a collection; the left
// side must evaluate to an [NSSet] object, and the right side can be any
// other collection type.
// 
// Core Data doesn’t support set expressions.
// 
// # Function Expressions
// 
// In macOS 10.4, [NSExpression] only supports a predefined set of functions:
// `sum`, `count`, `min`, `max`, and `average`. You access these predefined
// functions in the predicate syntax using custom keywords (for example,
// `MAX(1, 5, 10)`).
// 
// In macOS 10.5 and later, function expressions also support arbitrary method
// invocations. To implement this extended functionality, use the syntax
// `FUNCTION(receiver, selectorName, arguments, ...),` as in the following
// example:
// 
// All methods must take one or more `id` arguments and return an `id` value,
// although you can use the [CAST] expression to convert datatypes with lossy
// string representations (for example, `CAST(####, "NSDate")`). macOS 10.5
// extends the [CAST] expression to provide support for casting to classes for
// use in creating receivers for function expressions.
// 
// Although Core Data supports evaluation of the predefined functions, it
// doesn’t support the evaluation of custom predicate functions in the
// persistent stores (during a fetch).
//
// # Creating an Expression
//
//   - [NSExpression.InitWithExpressionType]: Creates the expression with the specified expression type.
//
// # Getting Information About an Expression
//
//   - [NSExpression.Arguments]: The arguments for the expression.
//   - [NSExpression.Collection]: The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//   - [NSExpression.ConstantValue]: The constant value of the expression.
//   - [NSExpression.ExpressionType]: The expression type for the expression.
//   - [NSExpression.Function]: The function for the expression.
//   - [NSExpression.KeyPath]: The key path for the expression.
//   - [NSExpression.Operand]: The operand for the expression.
//   - [NSExpression.Predicate]: The predicate of a subquery expression.
//   - [NSExpression.LeftExpression]: The left expression of an aggregate expression.
//   - [NSExpression.RightExpression]: The right expression of an aggregate expression.
//   - [NSExpression.Variable]: The variable for the expression.
//
// # Evaluating an Expression
//
//   - [NSExpression.ExpressionValueWithObjectContext]: Evaluates an expression using a specified object and context.
//   - [NSExpression.AllowEvaluation]: Forces a securely decoded expression to allow evaluation.
//   - [NSExpression.FalseExpression]: An expression to evalutate if a conditional expression’s predicate evaluates to false.
//   - [NSExpression.TrueExpression]: An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// # Accessing the Expression Block
//
//   - [NSExpression.ExpressionBlock]: The block that executes to evaluate the expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression
type NSExpression struct {
	objectivec.Object
}

// NSExpressionFromID constructs a [NSExpression] from an objc.ID.
//
// An expression for use in a comparison predicate.
func NSExpressionFromID(id objc.ID) NSExpression {
	return NSExpression{objectivec.Object{ID: id}}
}
// NOTE: NSExpression adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExpression] class.
//
// # Creating an Expression
//
//   - [INSExpression.InitWithExpressionType]: Creates the expression with the specified expression type.
//
// # Getting Information About an Expression
//
//   - [INSExpression.Arguments]: The arguments for the expression.
//   - [INSExpression.Collection]: The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//   - [INSExpression.ConstantValue]: The constant value of the expression.
//   - [INSExpression.ExpressionType]: The expression type for the expression.
//   - [INSExpression.Function]: The function for the expression.
//   - [INSExpression.KeyPath]: The key path for the expression.
//   - [INSExpression.Operand]: The operand for the expression.
//   - [INSExpression.Predicate]: The predicate of a subquery expression.
//   - [INSExpression.LeftExpression]: The left expression of an aggregate expression.
//   - [INSExpression.RightExpression]: The right expression of an aggregate expression.
//   - [INSExpression.Variable]: The variable for the expression.
//
// # Evaluating an Expression
//
//   - [INSExpression.ExpressionValueWithObjectContext]: Evaluates an expression using a specified object and context.
//   - [INSExpression.AllowEvaluation]: Forces a securely decoded expression to allow evaluation.
//   - [INSExpression.FalseExpression]: An expression to evalutate if a conditional expression’s predicate evaluates to false.
//   - [INSExpression.TrueExpression]: An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// # Accessing the Expression Block
//
//   - [INSExpression.ExpressionBlock]: The block that executes to evaluate the expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression
type INSExpression interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating an Expression

	// Creates the expression with the specified expression type.
	InitWithExpressionType(type_ NSExpressionType) NSExpression

	// Topic: Getting Information About an Expression

	// The arguments for the expression.
	Arguments() []NSExpression
	// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
	Collection() objectivec.IObject
	// The constant value of the expression.
	ConstantValue() objectivec.IObject
	// The expression type for the expression.
	ExpressionType() NSExpressionType
	// The function for the expression.
	Function() string
	// The key path for the expression.
	KeyPath() string
	// The operand for the expression.
	Operand() INSExpression
	// The predicate of a subquery expression.
	Predicate() INSPredicate
	// The left expression of an aggregate expression.
	LeftExpression() INSExpression
	// The right expression of an aggregate expression.
	RightExpression() INSExpression
	// The variable for the expression.
	Variable() string

	// Topic: Evaluating an Expression

	// Evaluates an expression using a specified object and context.
	ExpressionValueWithObjectContext(object objectivec.IObject, context INSDictionary) objectivec.IObject
	// Forces a securely decoded expression to allow evaluation.
	AllowEvaluation()
	// An expression to evalutate if a conditional expression’s predicate evaluates to false.
	FalseExpression() INSExpression
	// An expression to evalutate if a conditional expression’s predicate evaluates to true.
	TrueExpression() INSExpression

	// Topic: Accessing the Expression Block

	// The block that executes to evaluate the expression.
	ExpressionBlock() ArrayHandler
}

// Init initializes the instance.
func (e NSExpression) Init() NSExpression {
	rv := objc.Send[NSExpression](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExpression) Autorelease() NSExpression {
	rv := objc.Send[NSExpression](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExpression creates a new NSExpression instance.
func NewNSExpression() NSExpression {
	class := getNSExpressionClass()
	rv := objc.Send[NSExpression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an aggregate expression for a specified collection.
//
// subexpressions: A collection object (an instance of [NSArray], [NSSet], or [NSDictionary])
// that contains further expressions.
//
// # Return Value
// 
// A new expression that contains the expressions in `collection`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forAggregate:)
func NewExpressionForAggregate(subexpressions []NSExpression) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForAggregate:"), objectivec.IObjectSliceToNSArray(subexpressions))
	return NSExpressionFromID(rv)
}

// Creates an expression that returns a result, depending on the value of
// predicate.
//
// predicate: The predicate for determining whether the element belongs in the result
// collection.
//
// trueExpression: The expression for evaluation when the predicate evaluates to `true`.
//
// falseExpression: The expression for evaluation when the predicate evaluates to `false`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConditional:trueExpression:falseExpression:)
func NewExpressionForConditionalTrueExpressionFalseExpression(predicate INSPredicate, trueExpression INSExpression, falseExpression INSExpression) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForConditional:trueExpression:falseExpression:"), predicate, trueExpression, falseExpression)
	return NSExpressionFromID(rv)
}

// Creates an expression that represents a specified constant value.
//
// obj: The constant value the new expression is to represent.
//
// # Return Value
// 
// A new expression that represents the constant value, `obj`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConstantValue:)
func NewExpressionForConstantValue(obj objectivec.IObject) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForConstantValue:"), obj)
	return NSExpressionFromID(rv)
}

// Creates an expression that invokes one of the predefined functions.
//
// name: The name of the function to invoke.
//
// parameters: An array containing [NSExpression] objects that will be used as parameters
// during the invocation of selector.
// 
// For a selector taking no parameters, the array should be empty. For a
// selector taking one or more parameters, the array should contain one
// [NSExpression] object which will evaluate to an instance of the appropriate
// type for each parameter.
// 
// If there is a mismatch between the number of parameters expected and the
// number you provide during evaluation, an exception may be raised or missing
// parameters may simply be replaced by `nil` (which occurs depends on how
// many parameters are provided, and whether you have over- or underflow).
//
// # Return Value
// 
// A new expression that invokes the function `name` using the parameters in
// `parameters`.
//
// # Discussion
// 
// The `name` parameter can be one of the following predefined functions.
// 
// [Table data omitted]
// 
// This method raises an exception immediately if the selector is invalid; it
// raises an exception at runtime if the parameters are incorrect.
// 
// The `parameters` argument is a collection containing an expression which
// evaluates to a collection, as illustrated in the following examples:
// 
// # Special Considerations
// 
// This method throws an exception immediately if the selector is unknown; it
// throws at runtime if the parameters are incorrect.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:arguments:)
func NewExpressionForFunctionArguments(name string, parameters INSArray) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForFunction:arguments:"), objc.String(name), parameters)
	return NSExpressionFromID(rv)
}

// Creates an expression that returns the result of invoking a selector with a
// specified name using specified arguments.
//
// target: An [NSExpression] object which will evaluate an object on which the
// selector identified by `name` may be invoked.
//
// name: The name of the method to be invoked.
//
// parameters: An array containing [NSExpression] objects which can be evaluated to
// provide parameters for the method specified by `name`.
//
// # Return Value
// 
// An expression which will return the result of invoking the selector named
// `name` on the result of evaluating the target expression with the
// parameters specified by evaluating the elements of `parameters`.
//
// # Discussion
// 
// See the description of [ExpressionForFunctionArguments] for examples of how
// to construct the parameter array.
// 
// # Special Considerations
// 
// This method throws an exception immediately if the selector is unknown; it
// throws at runtime if the parameters are incorrect.
// 
// This expression effectively allows your application to invoke any method on
// any object it can navigate to at runtime. You must consider the security
// implications of this type of evaluation.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:selectorName:arguments:)
func NewExpressionForFunctionSelectorNameArguments(target INSExpression, name string, parameters INSArray) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForFunction:selectorName:arguments:"), target, objc.String(name), parameters)
	return NSExpressionFromID(rv)
}

// Creates an expression object that represents the intersection of a
// specified set and collection.
//
// left: An expression that evaluates to an [NSSet] object.
//
// right: An expression that evaluates to a collection object (an instance of
// [NSArray], [NSSet], or [NSDictionary]).
//
// # Return Value
// 
// A new [NSExpression] object that represents the intersection of `left` and
// `right`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forIntersectSet:with:)
func NewExpressionForIntersectSetWith(left INSExpression, right INSExpression) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForIntersectSet:with:"), left, right)
	return NSExpressionFromID(rv)
}

// Creates an expression that invokes the value function with a specified key
// path.
//
// keyPath: The key path that the new expression should evaluate.
//
// # Return Value
// 
// A new expression that invokes [value(forKeyPath:)] with `keyPath`.
//
// [value(forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKeyPath:)
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forKeyPath:)-1aqf5
func NewExpressionForKeyPath(keyPath string) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForKeyPath:"), objc.String(keyPath))
	return NSExpressionFromID(rv)
}

// Creates an expression object that represents the subtraction of a specified
// collection from a specified set.
//
// left: An expression that evaluates to an [NSSet] object.
//
// right: An expression that evaluates to a collection object (an instance of
// [NSArray], [NSSet], or [NSDictionary]).
//
// # Return Value
// 
// A new [NSExpression] object that represents the subtraction of `right` from
// `left`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forMinusSet:with:)
func NewExpressionForMinusSetWith(left INSExpression, right INSExpression) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForMinusSet:with:"), left, right)
	return NSExpressionFromID(rv)
}

// Creates an expression that filters a collection by storing elements in the
// collection in a specified variable and keeping the elements that the
// qualifier returns as true.
//
// expression: A predicate expression that evaluates to a collection.
//
// variable: Used as a local variable, and will shadow any instances of variable in the
// bindings dictionary. The variable is removed or the old value replaced once
// evaluation completes.
//
// predicate: The predicate used to determine whether the element belongs in the result
// collection.
//
// # Return Value
// 
// An expression that filters a collection by storing elements in the
// collection in the variable variable and keeping the elements for which
// qualifier returns true
//
// # Discussion
// 
// This method creates a sub-expression, evaluation of which returns a subset
// of a collection of objects. It allows you to create sophisticated queries
// across relationships, such as a search for multiple correlated values on
// the destination object of a relationship.
// 
// For example, suppose you have an Apartment entity that has a to-many
// relationship to a Resident entity, and that you want to create a query for
// all apartments inhabited by a resident whose first name is “Jane” and
// whose last name is “Doe”. Using only API available for OS X v 10.4, you
// could try the predicate:
// 
// but this will always return false since `resident.Firstname()` and
// `resident.Lastname()` both return collections. You could also try:
// 
// but this is also flawed—it returns true if there are two residents, one
// of whom is John Doe and one of whom is Jane Smith. The only way to find the
// desired apartments is to do two passes: one through residents to find
// “Jane Doe”, and one through apartments to find the ones where our Jane
// Does reside.
// 
// Subquery expressions provide a way to encapsulate this type of
// qualification into a single query.
// 
// The string format for a subquery expression is:
// 
// where `expression` is a predicate expression that evaluates to a
// collection, `variableExpression` is an expression which will be used to
// contain each individual element of `collection`, and `predicate` is the
// predicate used to determine whether the element belongs in the result
// collection.
// 
// Using subqueries, the apartment query could be reformulated as
// 
// or
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forSubquery:usingIteratorVariable:predicate:)
func NewExpressionForSubqueryUsingIteratorVariablePredicate(expression INSExpression, variable string, predicate INSPredicate) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), expression, objc.String(variable), predicate)
	return NSExpressionFromID(rv)
}

// Creates an expression object that represents the union of a specified set
// and collection.
//
// left: An expression that evaluates to an [NSSet] object.
//
// right: An expression that evaluates to a collection object (an instance of
// [NSArray], [NSSet], or [NSDictionary]).
//
// # Return Value
// 
// An new [NSExpression] object that represents the union of `left` and
// `right`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forUnionSet:with:)
func NewExpressionForUnionSetWith(left INSExpression, right INSExpression) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForUnionSet:with:"), left, right)
	return NSExpressionFromID(rv)
}

// Creates an expression that extracts a value from the variable bindings
// dictionary for a specified key.
//
// string: The key for the variable to extract from the variable bindings dictionary.
//
// # Return Value
// 
// A new expression that extracts from the variable bindings dictionary the
// value for the key `string`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(forVariable:)
func NewExpressionForVariable(string_ string) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionForVariable:"), objc.String(string_))
	return NSExpressionFromID(rv)
}

// Creates an expression by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(coder:)
func NewExpressionWithCoder(coder INSCoder) NSExpression {
	instance := getNSExpressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSExpressionFromID(rv)
}

// Creates the expression with the specified expression type.
//
// type: The type of the new expression, as defined by
// [NSExpression.ExpressionType].
// //
// [NSExpression.ExpressionType]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum
//
// # Return Value
// 
// An initialized [NSExpression] object of the type `type`.
//
// # Discussion
// 
// This method is the designated initializer for [NSExpression].
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(expressionType:)
func NewExpressionWithExpressionType(type_ NSExpressionType) NSExpression {
	instance := getNSExpressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressionType:"), type_)
	return NSExpressionFromID(rv)
}

// Creates the expression with the specified expression format and array of
// arguments.
//
// expressionFormat: The expression format.
//
// arguments: An array of arguments to be used with the `expressionFormat` string.
//
// # Return Value
// 
// An initialized [NSExpression] object with the specified arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:argumentArray:)
func NewExpressionWithFormatArgumentArray(expressionFormat string, arguments INSArray) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionWithFormat:argumentArray:"), objc.String(expressionFormat), arguments)
	return NSExpressionFromID(rv)
}

// Creates the expression with the specified expression format and arguments
// list.
//
// expressionFormat: The expression format.
//
// argList: A list of arguments to be inserted into the `expressionFormat` string. The
// argument list is terminated by `nil`.
//
// # Return Value
// 
// An initialized [NSExpression] object with the specified arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:arguments:)
func NewExpressionWithFormatArguments(expressionFormat string, argList unsafe.Pointer) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(getNSExpressionClass().class), objc.Sel("expressionWithFormat:arguments:"), objc.String(expressionFormat), argList)
	return NSExpressionFromID(rv)
}

// Creates the expression with the specified expression type.
//
// type: The type of the new expression, as defined by
// [NSExpression.ExpressionType].
// //
// [NSExpression.ExpressionType]: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum
//
// # Return Value
// 
// An initialized [NSExpression] object of the type `type`.
//
// # Discussion
// 
// This method is the designated initializer for [NSExpression].
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(expressionType:)
func (e NSExpression) InitWithExpressionType(type_ NSExpressionType) NSExpression {
	rv := objc.Send[NSExpression](e.ID, objc.Sel("initWithExpressionType:"), type_)
	return rv
}
// Creates an expression by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(coder:)
func (e NSExpression) InitWithCoder(coder INSCoder) NSExpression {
	rv := objc.Send[NSExpression](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Evaluates an expression using a specified object and context.
//
// object: The object against which the expression is evaluated.
//
// context: A dictionary that the expression can use to store temporary state for one
// predicate evaluation. Can be `nil`.
// 
// Note that `context` is mutable, and that it can only be accessed during the
// evaluation of the expression. You must not attempt to retain it for use
// elsewhere.
//
// # Return Value
// 
// The evaluated object.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e NSExpression) ExpressionValueWithObjectContext(object objectivec.IObject, context INSDictionary) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("expressionValueWithObject:context:"), object, context)
	return objectivec.Object{ID: rv}
}
// Forces a securely decoded expression to allow evaluation.
//
// # Discussion
// 
// When securely decoding an [NSExpression] object encoded using
// [NSSecureCoding], evaluation is disabled because it is potentially unsafe
// to evaluate expressions you get out of an archive.
// 
// Before you enable evaluation, you should validate key paths, selectors, etc
// to ensure no erroneous or malicious code will be executed. Once you’ve
// preflighted the expression, you can enable the expression for evaluation by
// calling `allowEvaluation`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/allowEvaluation()
func (e NSExpression) AllowEvaluation() {
	objc.Send[objc.ID](e.ID, objc.Sel("allowEvaluation"))
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (e NSExpression) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates an expression that represents the object you’re evaluating.
//
// # Return Value
// 
// A new expression that represents the object being evaluated.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForEvaluatedObject()
func (_NSExpressionClass NSExpressionClass) ExpressionForEvaluatedObject() NSExpression {
	rv := objc.Send[objc.ID](objc.ID(_NSExpressionClass.class), objc.Sel("expressionForEvaluatedObject"))
	return NSExpressionFromID(rv)
}
// Creates an expression that represents any key for a Spotlight query.
//
// # Return Value
// 
// A new expression that represents any key for a Spotlight query.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForAnyKey()
func (_NSExpressionClass NSExpressionClass) ExpressionForAnyKey() NSExpression {
	rv := objc.Send[objc.ID](objc.ID(_NSExpressionClass.class), objc.Sel("expressionForAnyKey"))
	return NSExpressionFromID(rv)
}
// Creates an expression object that uses the block for evaluating objects.
//
// block: The Block is applied to the object to be evaluated.
// 
// The Block takes three arguments and returns a value:
// 
// evaluatedObject: The object to be evaluated. expressions: An array of
// predicate expressions that evaluates to a collection. context: A dictionary
// that the expression can use to store temporary state for one predicate
// evaluation.
// 
// Note that `context` is mutable, and that it can only be accessed during the
// evaluation of the expression. You must not attempt to retain it for use
// elsewhere. ]
// 
// The Block returns the `evaluatedObject`.
//
// arguments: An array containing [NSExpression] objects that will be used as parameters
// during the invocation of selector.
// 
// For a selector taking no parameters, the array should be empty. For a
// selector taking one or more parameters, the array should contain one
// [NSExpression] object which will evaluate to an instance of the appropriate
// type for each parameter.
// 
// If there is a mismatch between the number of parameters expected and the
// number you provide during evaluation, an exception may be raised or missing
// parameters may simply be replaced by `nil` (which occurs depends on how
// many parameters are provided, and whether you have over- or underflow).
// 
// See [ExpressionForFunctionArguments] for a complete list of arguments.
//
// # Return Value
// 
// An expression that filters a collection using the specified Block.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/init(block:arguments:)
func (_NSExpressionClass NSExpressionClass) ExpressionForBlockArguments(block ArrayHandler, arguments []NSExpression) NSExpression {
_block0, _ := NewArrayBlock(block)
	rv := objc.Send[objc.ID](objc.ID(_NSExpressionClass.class), objc.Sel("expressionForBlock:arguments:"), _block0, arguments)
	return NSExpressionFromID(rv)
}
// Creates the expression with the specified expression arguments.
//
// expressionFormat: The expression format.
//
// # Return Value
// 
// An initialized [NSExpression] object with the specified format.
//
// # Discussion
// 
// After `expressionFormat`, pass a comma-separated list of arguments to
// substitute into format as variadic arguments. The list is terminated by
// `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionWithFormat:
func (_NSExpressionClass NSExpressionClass) ExpressionWithFormat(expressionFormat string) NSExpression {
	rv := objc.Send[objc.ID](objc.ID(_NSExpressionClass.class), objc.Sel("expressionWithFormat:"), objc.String(expressionFormat))
	return NSExpressionFromID(rv)
}

// The arguments for the expression.
//
// # Discussion
// 
// An expression’s arguments is the array of expressions that will be passed
// as parameters during invocation of the selector on the operand of a
// function expression.
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/arguments
func (e NSExpression) Arguments() []NSExpression {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("arguments"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSExpression {
		return NSExpressionFromID(id)
	})
}
// The collection of expressions in an aggregate expression, or the collection
// element of a subquery expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/collection
func (e NSExpression) Collection() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("collection"))
	return objectivec.Object{ID: rv}
}
// The constant value of the expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/constantValue
func (e NSExpression) ConstantValue() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("constantValue"))
	return objectivec.Object{ID: rv}
}
// The expression type for the expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionType-swift.property
func (e NSExpression) ExpressionType() NSExpressionType {
	rv := objc.Send[NSExpressionType](e.ID, objc.Sel("expressionType"))
	return NSExpressionType(rv)
}
// The function for the expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/function
func (e NSExpression) Function() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("function"))
	return NSStringFromID(rv).String()
}
// The key path for the expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/keyPath
func (e NSExpression) KeyPath() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("keyPath"))
	return NSStringFromID(rv).String()
}
// The operand for the expression.
//
// # Discussion
// 
// The operand for an expression is the object on which the expression’s
// selector or block will be invoked. The object is the result of evaluating a
// key path or one of the defined functions. Accessing this property raises an
// exception if it is not applicable to the expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/operand
func (e NSExpression) Operand() INSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("operand"))
	return NSExpressionFromID(objc.ID(rv))
}
// The predicate of a subquery expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/predicate
func (e NSExpression) Predicate() INSPredicate {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("predicate"))
	return NSPredicateFromID(objc.ID(rv))
}
// The left expression of an aggregate expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/left
func (e NSExpression) LeftExpression() INSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("leftExpression"))
	return NSExpressionFromID(objc.ID(rv))
}
// The right expression of an aggregate expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/right
func (e NSExpression) RightExpression() INSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("rightExpression"))
	return NSExpressionFromID(objc.ID(rv))
}
// The variable for the expression.
//
// # Discussion
// 
// Accessing this property raises an exception if it is not applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/variable
func (e NSExpression) Variable() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variable"))
	return NSStringFromID(rv).String()
}
// An expression to evalutate if a conditional expression’s predicate
// evaluates to false.
//
// # Discussion
// 
// Accessing this property raises an exception if it isn’t applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/false
func (e NSExpression) FalseExpression() INSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("falseExpression"))
	return NSExpressionFromID(objc.ID(rv))
}
// An expression to evalutate if a conditional expression’s predicate
// evaluates to true.
//
// # Discussion
// 
// Accessing this property raises an exception if it isn’t applicable to the
// expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/true
func (e NSExpression) TrueExpression() INSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("trueExpression"))
	return NSExpressionFromID(objc.ID(rv))
}
// The block that executes to evaluate the expression.
//
// See: https://developer.apple.com/documentation/Foundation/NSExpression/expressionBlock
func (e NSExpression) ExpressionBlock() ArrayHandler {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("expressionBlock"))
	_ = rv
	return nil
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

