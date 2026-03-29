// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPredicate] class.
var (
	_NSPredicateClass     NSPredicateClass
	_NSPredicateClassOnce sync.Once
)

func getNSPredicateClass() NSPredicateClass {
	_NSPredicateClassOnce.Do(func() {
		_NSPredicateClass = NSPredicateClass{class: objc.GetClass("NSPredicate")}
	})
	return _NSPredicateClass
}

// GetNSPredicateClass returns the class object for NSPredicate.
func GetNSPredicateClass() NSPredicateClass {
	return getNSPredicateClass()
}

type NSPredicateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPredicateClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPredicateClass) Alloc() NSPredicate {
	rv := objc.Send[NSPredicate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A definition of logical conditions for constraining a search for a fetch or
// for in-memory filtering.
//
// # Overview
// 
// Predicates represent logical conditions, which you can use to filter
// collections of objects. Although it’s common to create predicates
// directly from instances of [NSComparisonPredicate], [NSCompoundPredicate],
// and [NSExpression], you often create predicates from a format string that
// the class methods parse on [NSPredicate]. Examples of predicate format
// strings include:
// 
// - Simple comparisons, such as `grade == "7"` or `firstName like "Juan"` -
// Case- and diacritic-insensitive lookups, such as `name contains[cd]
// "stein"` - Logical operations, such as `(firstName like "Mei") OR (lastName
// like "Chen")` - Temporal range constraints, such as `date between
// {$YESTERDAY, $TOMORROW}` - Relational conditions, such as `group.Name()
// like "work*"` - Aggregate operations, such as
// `@sum.ItemsXCUIElementTypePrice() < 1000`
// 
// For a complete syntax reference, refer to the [Predicate Programming
// Guide].
// 
// You can also create predicates that include variables using the
// [NSPredicate.EvaluateWithObjectSubstitutionVariables] method so that you can predefine
// the predicate before substituting concrete values at runtime.
//
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
//
// # Creating a Predicate
//
//   - [NSPredicate.PredicateWithSubstitutionVariables]: Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary.
//
// # Evaluating a Predicate
//
//   - [NSPredicate.EvaluateWithObject]: Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies.
//   - [NSPredicate.EvaluateWithObjectSubstitutionVariables]: Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary.
//   - [NSPredicate.AllowEvaluation]: Forces a securely decoded predicate to allow evaluation.
//
// # Getting a String Representation
//
//   - [NSPredicate.PredicateFormat]: The predicate’s format string.
//
// # Instance Methods
//
//   - [NSPredicate.AllowEvaluationWithValidatorError]
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate
type NSPredicate struct {
	objectivec.Object
}

// NSPredicateFromID constructs a [NSPredicate] from an objc.ID.
//
// A definition of logical conditions for constraining a search for a fetch or
// for in-memory filtering.
func NSPredicateFromID(id objc.ID) NSPredicate {
	return NSPredicate{objectivec.Object{ID: id}}
}
// NOTE: NSPredicate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPredicate] class.
//
// # Creating a Predicate
//
//   - [INSPredicate.PredicateWithSubstitutionVariables]: Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary.
//
// # Evaluating a Predicate
//
//   - [INSPredicate.EvaluateWithObject]: Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies.
//   - [INSPredicate.EvaluateWithObjectSubstitutionVariables]: Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary.
//   - [INSPredicate.AllowEvaluation]: Forces a securely decoded predicate to allow evaluation.
//
// # Getting a String Representation
//
//   - [INSPredicate.PredicateFormat]: The predicate’s format string.
//
// # Instance Methods
//
//   - [INSPredicate.AllowEvaluationWithValidatorError]
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate
type INSPredicate interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a Predicate

	// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary.
	PredicateWithSubstitutionVariables(variables INSDictionary) INSPredicate

	// Topic: Evaluating a Predicate

	// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies.
	EvaluateWithObject(object objectivec.IObject) bool
	// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary.
	EvaluateWithObjectSubstitutionVariables(object objectivec.IObject, bindings INSDictionary) bool
	// Forces a securely decoded predicate to allow evaluation.
	AllowEvaluation()

	// Topic: Getting a String Representation

	// The predicate’s format string.
	PredicateFormat() string

	// Topic: Instance Methods

	AllowEvaluationWithValidatorError(validator NSPredicateValidating) (bool, error)
}

// Init initializes the instance.
func (p NSPredicate) Init() NSPredicate {
	rv := objc.Send[NSPredicate](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPredicate) Autorelease() NSPredicate {
	rv := objc.Send[NSPredicate](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPredicate creates a new NSPredicate instance.
func NewNSPredicate() NSPredicate {
	class := getNSPredicateClass()
	rv := objc.Send[NSPredicate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a predicate with a metadata query string.
//
// queryString: A metadata query string.
//
// # Discussion
// 
// For details of the format of the query string, see [File Metadata Query
// Expression Syntax].
//
// [File Metadata Query Expression Syntax]: https://developer.apple.com/library/archive/documentation/Carbon/Conceptual/SpotlightQuery/Concepts/QueryFormat.html#//apple_ref/doc/uid/TP40001849
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/init(fromMetadataQueryString:)
func NewPredicateFromMetadataQueryString(queryString string) NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSPredicateClass().class), objc.Sel("predicateFromMetadataQueryString:"), objc.String(queryString))
	return NSPredicateFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPredicateWithCoder(coder INSCoder) NSPredicate {
	instance := getNSPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPredicateFromID(rv)
}

// Creates a predicate by substituting the values in a specified array into a
// format string and parsing the result.
//
// predicateFormat: The format string for the new predicate.
//
// arguments: The arguments to substitute into `predicateFormat`. Values are substituted
// in the order they appear in the array.
//
// # Return Value
// 
// A new predicate by substituting the values in `arguments` into
// `predicateFormat`, and parsing the result.
//
// # Discussion
// 
// For details of the format of the format string and of limitations on
// variable substitution, see [Predicate Format String Syntax].
//
// [Predicate Format String Syntax]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pSyntax.html#//apple_ref/doc/uid/TP40001795
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:argumentArray:)
func NewPredicateWithFormatArgumentArray(predicateFormat string, arguments INSArray) NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSPredicateClass().class), objc.Sel("predicateWithFormat:argumentArray:"), objc.String(predicateFormat), arguments)
	return NSPredicateFromID(rv)
}

// Creates a predicate by substituting the values in an argument list into a
// format string and parsing the result.
//
// predicateFormat: The format string for the new predicate.
//
// argList: The arguments to substitute into `predicateFormat`. Values are substituted
// in the order they appear in the argument list.
//
// # Return Value
// 
// A new predicate by substituting the values in `argList` into
// `predicateFormat` and parsing the result.
//
// # Discussion
// 
// For details of the format of the format string and of limitations on
// variable substitution, see [Predicate Format String Syntax].
//
// [Predicate Format String Syntax]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pSyntax.html#//apple_ref/doc/uid/TP40001795
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:arguments:)
func NewPredicateWithFormatArguments(predicateFormat string, argList unsafe.Pointer) NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSPredicateClass().class), objc.Sel("predicateWithFormat:arguments:"), objc.String(predicateFormat), argList)
	return NSPredicateFromID(rv)
}

// Creates and returns a predicate that always evaluates to a specified
// Boolean value.
//
// value: The Boolean value to which the new predicate should evaluate.
//
// # Return Value
// 
// A predicate that always evaluates to `value`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/init(value:)
func NewPredicateWithValue(value bool) NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSPredicateClass().class), objc.Sel("predicateWithValue:"), value)
	return NSPredicateFromID(rv)
}

// Returns a copy of the predicate and substitutes the predicates variables
// with specified values from a specified substitution variables dictionary.
//
// variables: The substitution variables dictionary. The dictionary must contain
// key-value pairs for all variables in the receiver.
//
// # Return Value
// 
// A copy of the receiver with the predicate’s variables substituted by
// values specified in `variables`.
//
// # Discussion
// 
// The predicate itself is not modified by this method, so you can reuse it
// for any number of substitutions.
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/withSubstitutionVariables(_:)
func (p NSPredicate) PredicateWithSubstitutionVariables(variables INSDictionary) INSPredicate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return NSPredicateFromID(rv)
}
// Returns a Boolean value that indicates whether the specified object matches
// the conditions that the predicate specifies.
//
// object: The object against which to evaluate the predicate.
//
// # Return Value
// 
// [true] if `object` matches the conditions specified by the predicate,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:)
func (p NSPredicate) EvaluateWithObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("evaluateWithObject:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the specified object matches
// the conditions that the predicate specifies after substituting in the
// values from a specified variables dictionary.
//
// object: The object against which to evaluate the predicate.
//
// bindings: The substitution variables dictionary. The dictionary must contain
// key-value pairs for all variables in the predicate.
//
// # Return Value
// 
// [true] if `object` matches the conditions specified by the predicate after
// substituting in the values in `bindings` for any replacement tokens,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns the same result as the two step process of first
// invoking [PredicateWithSubstitutionVariables] on the predicate and then
// invoking [EvaluateWithObject] on the returned value. This method is
// optimized for situations which require repeatedly evaluating a predicate
// with substitution variables with different variable substitutions.
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:substitutionVariables:)
func (p NSPredicate) EvaluateWithObjectSubstitutionVariables(object objectivec.IObject, bindings INSDictionary) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("evaluateWithObject:substitutionVariables:"), object, bindings)
	return rv
}
// Forces a securely decoded predicate to allow evaluation.
//
// # Discussion
// 
// When securely decoding [NSPredicate] objects that are encoded using
// [NSSecureCoding], evaluation is disabled because it is potentially unsafe
// to evaluate predicates you get out of an archive.
// 
// Before you enable evaluation, you should validate key paths, selectors, and
// other details to ensure no erroneous or malicious code will be executed.
// Once you’ve verified the predicate, you can enable the receiver for
// evaluation by calling [AllowEvaluation].
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/allowEvaluation()
func (p NSPredicate) AllowEvaluation() {
	objc.Send[objc.ID](p.ID, objc.Sel("allowEvaluation"))
}
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/allowEvaluation(validator:)
func (p NSPredicate) AllowEvaluationWithValidatorError(validator NSPredicateValidating) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("allowEvaluationWithValidator:error:"), validator, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("allowEvaluationWithValidator:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (p NSPredicate) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (p NSPredicate) InitWithCoder(coder INSCoder) NSPredicate {
	rv := objc.Send[NSPredicate](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates and returns a new predicate formed by creating a new string with a
// specified format and parsing the result.
//
// predicateFormat: The format string for the new predicate.
//
// # Return Value
// 
// A new predicate formed by creating a new string with `format` and parsing
// the result.
//
// # Discussion
// 
// Pass a comma-separated list of trailing variadic arguments to substitute
// into `format`.
// 
// For details of the format of the format string and of limitations on
// variable substitution, see [Predicate Format String Syntax].
//
// [Predicate Format String Syntax]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pSyntax.html#//apple_ref/doc/uid/TP40001795
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/predicateWithFormat:
func (_NSPredicateClass NSPredicateClass) PredicateWithFormat(predicateFormat string) NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_NSPredicateClass.class), objc.Sel("predicateWithFormat:"), objc.String(predicateFormat))
	return NSPredicateFromID(rv)
}

// The predicate’s format string.
//
// # Discussion
// 
// The return value of this property is not guaranteed to be the same as the
// string used to create the predicate using [PredicateWithFormat] or similar
// methods.
// 
// You cannot use this method to create a persistent representation of a
// predicate suitable for recreating the original predicate. If you need a
// persistent representation of a predicate, create an archive instead, as
// described in [Object archiving] ([NSPredicate] adopts the [NSCoding]
// protocol).
//
// [Object archiving]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Archiving.html#//apple_ref/doc/uid/TP40008195-CH1
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicate/predicateFormat
func (p NSPredicate) PredicateFormat() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("predicateFormat"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

