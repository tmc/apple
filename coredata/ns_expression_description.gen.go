// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSExpressionDescription] class.
var (
	_NSExpressionDescriptionClass     NSExpressionDescriptionClass
	_NSExpressionDescriptionClassOnce sync.Once
)

func getNSExpressionDescriptionClass() NSExpressionDescriptionClass {
	_NSExpressionDescriptionClassOnce.Do(func() {
		_NSExpressionDescriptionClass = NSExpressionDescriptionClass{class: objc.GetClass("NSExpressionDescription")}
	})
	return _NSExpressionDescriptionClass
}

// GetNSExpressionDescriptionClass returns the class object for NSExpressionDescription.
func GetNSExpressionDescriptionClass() NSExpressionDescriptionClass {
	return getNSExpressionDescriptionClass()
}

type NSExpressionDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExpressionDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExpressionDescriptionClass) Alloc() NSExpressionDescription {
	rv := objc.Send[NSExpressionDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an expression to include with a fetch request.
//
// # Overview
//
// An expression description describes a value that a fetch request returns,
// which doesn’t appear as an attribute or relationship on an entity. For
// example, expressions can aggregate data, or transform an attribute’s
// value. You add expression descriptions to a fetch request using the
// [NSFetchRequest.PropertiesToFetch] method.
//
// # Configuring the Expression Description
//
//   - [NSExpressionDescription.Expression]: The expression to evaluate.
//   - [NSExpressionDescription.SetExpression]
//   - [NSExpressionDescription.ExpressionResultType]: The attribute type of the expression’s result.
//   - [NSExpressionDescription.SetExpressionResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSExpressionDescription
type NSExpressionDescription struct {
	NSPropertyDescription
}

// NSExpressionDescriptionFromID constructs a [NSExpressionDescription] from an objc.ID.
//
// An object that describes an expression to include with a fetch request.
func NSExpressionDescriptionFromID(id objc.ID) NSExpressionDescription {
	return NSExpressionDescription{NSPropertyDescription: NSPropertyDescriptionFromID(id)}
}

// NOTE: NSExpressionDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExpressionDescription] class.
//
// # Configuring the Expression Description
//
//   - [INSExpressionDescription.Expression]: The expression to evaluate.
//   - [INSExpressionDescription.SetExpression]
//   - [INSExpressionDescription.ExpressionResultType]: The attribute type of the expression’s result.
//   - [INSExpressionDescription.SetExpressionResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSExpressionDescription
type INSExpressionDescription interface {
	INSPropertyDescription

	// Topic: Configuring the Expression Description

	// The expression to evaluate.
	Expression() foundation.NSExpression
	SetExpression(value foundation.NSExpression)
	// The attribute type of the expression’s result.
	ExpressionResultType() NSAttributeType
	SetExpressionResultType(value NSAttributeType)
}

// Init initializes the instance.
func (e NSExpressionDescription) Init() NSExpressionDescription {
	rv := objc.Send[NSExpressionDescription](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExpressionDescription) Autorelease() NSExpressionDescription {
	rv := objc.Send[NSExpressionDescription](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExpressionDescription creates a new NSExpressionDescription instance.
func NewNSExpressionDescription() NSExpressionDescription {
	class := getNSExpressionDescriptionClass()
	rv := objc.Send[NSExpressionDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The expression to evaluate.
//
// See: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expression
func (e NSExpressionDescription) Expression() foundation.NSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("expression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}
func (e NSExpressionDescription) SetExpression(value foundation.NSExpression) {
	objc.Send[struct{}](e.ID, objc.Sel("setExpression:"), value)
}

// The attribute type of the expression’s result.
//
// See: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expressionResultType
func (e NSExpressionDescription) ExpressionResultType() NSAttributeType {
	rv := objc.Send[NSAttributeType](e.ID, objc.Sel("expressionResultType"))
	return NSAttributeType(rv)
}
func (e NSExpressionDescription) SetExpressionResultType(value NSAttributeType) {
	objc.Send[struct{}](e.ID, objc.Sel("setExpressionResultType:"), value)
}
