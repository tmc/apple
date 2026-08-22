// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSFetchRequestExpression] class.
var (
	_NSFetchRequestExpressionClass     NSFetchRequestExpressionClass
	_NSFetchRequestExpressionClassOnce sync.Once
)

func getNSFetchRequestExpressionClass() NSFetchRequestExpressionClass {
	_NSFetchRequestExpressionClassOnce.Do(func() {
		_NSFetchRequestExpressionClass = NSFetchRequestExpressionClass{class: objc.GetClass("NSFetchRequestExpression")}
	})
	return _NSFetchRequestExpressionClass
}

// GetNSFetchRequestExpressionClass returns the class object for NSFetchRequestExpression.
func GetNSFetchRequestExpressionClass() NSFetchRequestExpressionClass {
	return getNSFetchRequestExpressionClass()
}

type NSFetchRequestExpressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchRequestExpressionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchRequestExpressionClass) Alloc() NSFetchRequestExpression {
	rv := objc.Send[NSFetchRequestExpression](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An expression that evaluates the result of a fetch request on a managed
// object context.
//
// # Overview
//
// [NSFetchRequestExpression] inherits from [NSExpression], which provides
// most of the basic behavior. The first argument must be an expression which
// evaluates to an [NSFetchRequest] object, and the second must be an
// expression which evaluates to an [NSManagedObjectContext] object. If you
// simply want the count for the request, the `countOnly` argument should be
// true.
//
// # Examining a Fetch Request Expression
//
//   - [NSFetchRequestExpression.RequestExpression]: The expression for the receiver’s fetch request.
//   - [NSFetchRequestExpression.ContextExpression]: The expression for the receiver’s managed object context.
//   - [NSFetchRequestExpression.IsCountOnlyRequest]: Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
type NSFetchRequestExpression struct {
	foundation.NSExpression
}

// NSFetchRequestExpressionFromID constructs a [NSFetchRequestExpression] from an objc.ID.
//
// An expression that evaluates the result of a fetch request on a managed
// object context.
func NSFetchRequestExpressionFromID(id objc.ID) NSFetchRequestExpression {
	return NSFetchRequestExpression{NSExpression: foundation.NSExpressionFromID(id)}
}

// NOTE: NSFetchRequestExpression adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchRequestExpression] class.
//
// # Examining a Fetch Request Expression
//
//   - [INSFetchRequestExpression.RequestExpression]: The expression for the receiver’s fetch request.
//   - [INSFetchRequestExpression.ContextExpression]: The expression for the receiver’s managed object context.
//   - [INSFetchRequestExpression.IsCountOnlyRequest]: Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression
type INSFetchRequestExpression interface {
	foundation.INSExpression

	// Topic: Examining a Fetch Request Expression

	// The expression for the receiver’s fetch request.
	RequestExpression() foundation.NSExpression
	// The expression for the receiver’s managed object context.
	ContextExpression() foundation.NSExpression
	// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
	IsCountOnlyRequest() bool
}

// Init initializes the instance.
func (f NSFetchRequestExpression) Init() NSFetchRequestExpression {
	rv := objc.Send[NSFetchRequestExpression](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchRequestExpression) Autorelease() NSFetchRequestExpression {
	rv := objc.Send[NSFetchRequestExpression](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchRequestExpression creates a new NSFetchRequestExpression instance.
func NewNSFetchRequestExpression() NSFetchRequestExpression {
	class := getNSFetchRequestExpressionClass()
	rv := objc.Send[NSFetchRequestExpression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an expression which will evaluate to the result of executing a
// fetch request on a context.
//
// fetch: An expression that evaluates to an instance of [NSFetchRequest].
//
// context: An expression that evaluates to an instance of [NSManagedObjectContext].
//
// countFlag: If true, when the new expression is evaluated the managed object context
// (from `context`) will perform
// [NSManagedObjectContext.CountForFetchRequestError] with the fetch request
// (from `fetch`). If false, when the new expression is evaluated the managed
// object context will perform [fetch(_:)] with the fetch request.
//
// # Return Value
//
// An expression which will evaluate to the result of executing a fetch
// request (from `fetch`) on a managed object context (from `context`).
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/expression(forFetch:context:countOnly:)
//
// [fetch(_:)]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/fetch(_:)-38ys1
func (_NSFetchRequestExpressionClass NSFetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch foundation.NSExpression, context foundation.NSExpression, countFlag bool) foundation.NSExpression {
	rv := objc.Send[objc.ID](objc.ID(_NSFetchRequestExpressionClass.class), objc.Sel("expressionForFetch:context:countOnly:"), fetch, context, countFlag)
	return foundation.NSExpressionFromID(rv)
}

// The expression for the receiver’s fetch request.
//
// # Discussion
//
// The expression must evaluate to an [NSFetchRequest] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/requestExpression
func (f NSFetchRequestExpression) RequestExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("requestExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}

// The expression for the receiver’s managed object context.
//
// # Discussion
//
// The expression must evaluate to an [NSManagedObjectContext] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/contextExpression
func (f NSFetchRequestExpression) ContextExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("contextExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}

// Returns a Boolean value that indicates whether the receiver represents a
// count-only fetch request.
//
// # Discussion
//
// true if the receiver represents a count-only fetch request, otherwise
// false. If this method returns false, the managed object context (from the
// [NSFetchRequestExpression.ContextExpression]) will perform [fetch(_:)]:
// with the [NSFetchRequestExpression.RequestExpression]; if this method
// returns true, the managed object context will perform
// [NSManagedObjectContext.CountForFetchRequestError] with the
// [NSFetchRequestExpression.RequestExpression].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/isCountOnlyRequest
//
// [fetch(_:)]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/fetch(_:)-38ys1
func (f NSFetchRequestExpression) IsCountOnlyRequest() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isCountOnlyRequest"))
	return rv
}
