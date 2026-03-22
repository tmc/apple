// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSPredicateValidating protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicateValidating
type NSPredicateValidating interface {
	objectivec.IObject
}

// NSPredicateValidatingObject wraps an existing Objective-C object that conforms to the NSPredicateValidating protocol.
type NSPredicateValidatingObject struct {
	objectivec.Object
}
func (o NSPredicateValidatingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPredicateValidatingObjectFromID constructs a [NSPredicateValidatingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPredicateValidatingObjectFromID(id objc.ID) NSPredicateValidatingObject {
	return NSPredicateValidatingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Foundation/NSPredicateValidating/visit(_:)-491gq
func (o NSPredicateValidatingObject) VisitPredicateError(predicate INSPredicate) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("visitPredicate:error:"), predicate)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicateValidating/visit(_:)-9r82q
func (o NSPredicateValidatingObject) VisitOperatorTypeError(operatorType NSPredicateOperatorType) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("visitOperatorType:error:"), operatorType)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicateValidating/visit(_:)-9s9ho
func (o NSPredicateValidatingObject) VisitExpressionError(expression INSExpression) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("visitExpression:error:"), expression)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/Foundation/NSPredicateValidating/visitExpressionKeyPath(_:scope:key:)
func (o NSPredicateValidatingObject) VisitExpressionKeyPathScopeKeyError(expression INSExpression, scope string, key string) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("visitExpressionKeyPath:scope:key:error:"), expression, objc.String(scope), objc.String(key))
	if err != nil {
		return false, err
	}
	return rv, nil
	}

