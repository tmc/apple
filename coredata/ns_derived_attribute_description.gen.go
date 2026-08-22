// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSDerivedAttributeDescription] class.
var (
	_NSDerivedAttributeDescriptionClass     NSDerivedAttributeDescriptionClass
	_NSDerivedAttributeDescriptionClassOnce sync.Once
)

func getNSDerivedAttributeDescriptionClass() NSDerivedAttributeDescriptionClass {
	_NSDerivedAttributeDescriptionClassOnce.Do(func() {
		_NSDerivedAttributeDescriptionClass = NSDerivedAttributeDescriptionClass{class: objc.GetClass("NSDerivedAttributeDescription")}
	})
	return _NSDerivedAttributeDescriptionClass
}

// GetNSDerivedAttributeDescriptionClass returns the class object for NSDerivedAttributeDescription.
func GetNSDerivedAttributeDescriptionClass() NSDerivedAttributeDescriptionClass {
	return getNSDerivedAttributeDescriptionClass()
}

type NSDerivedAttributeDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDerivedAttributeDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDerivedAttributeDescriptionClass) Alloc() NSDerivedAttributeDescription {
	rv := objc.Send[NSDerivedAttributeDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of an attribute that derives its value by performing a
// calculation on a related attribute.
//
// # Overview
//
// Use derived attributes to optimize fetch performance; for example:
//
// - Create a derived `searchName` attribute to reflect a `name` attribute
// with case and diacritics removed for more efficient comparison. - Create a
// derived `relationshipCount` attribute to reflect the number of objects in a
// relationship and avoid having to do a join.
//
// Derived attributes support the following expressions:
//
// [Table data omitted]
//
// # Specifying the Derivation Expression
//
//   - [NSDerivedAttributeDescription.DerivationExpression]: An expression for generating derived data.
//   - [NSDerivedAttributeDescription.SetDerivationExpression]
//
// See: https://developer.apple.com/documentation/CoreData/NSDerivedAttributeDescription
type NSDerivedAttributeDescription struct {
	NSAttributeDescription
}

// NSDerivedAttributeDescriptionFromID constructs a [NSDerivedAttributeDescription] from an objc.ID.
//
// A description of an attribute that derives its value by performing a
// calculation on a related attribute.
func NSDerivedAttributeDescriptionFromID(id objc.ID) NSDerivedAttributeDescription {
	return NSDerivedAttributeDescription{NSAttributeDescription: NSAttributeDescriptionFromID(id)}
}

// NOTE: NSDerivedAttributeDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDerivedAttributeDescription] class.
//
// # Specifying the Derivation Expression
//
//   - [INSDerivedAttributeDescription.DerivationExpression]: An expression for generating derived data.
//   - [INSDerivedAttributeDescription.SetDerivationExpression]
//
// See: https://developer.apple.com/documentation/CoreData/NSDerivedAttributeDescription
type INSDerivedAttributeDescription interface {
	INSAttributeDescription

	// Topic: Specifying the Derivation Expression

	// An expression for generating derived data.
	DerivationExpression() foundation.NSExpression
	SetDerivationExpression(value foundation.NSExpression)
}

// Init initializes the instance.
func (d NSDerivedAttributeDescription) Init() NSDerivedAttributeDescription {
	rv := objc.Send[NSDerivedAttributeDescription](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDerivedAttributeDescription) Autorelease() NSDerivedAttributeDescription {
	rv := objc.Send[NSDerivedAttributeDescription](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDerivedAttributeDescription creates a new NSDerivedAttributeDescription instance.
func NewNSDerivedAttributeDescription() NSDerivedAttributeDescription {
	class := getNSDerivedAttributeDescriptionClass()
	rv := objc.Send[NSDerivedAttributeDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An expression for generating derived data.
//
// # Discussion
//
// # When using derived attributes in an SQL store, this expression should be
//
// - a keypath expression (including @operation components)
//
// a function expression using one of the predefined functions defined in
// [NSExpression]
//
// Any keypaths used in the expression must be accessible from the entity on
// which the derived attribute is specified.
//
// If you try to add a store to a coordinator whose model contains derived
// attributes of a type not supported by the store, the add fails and throws
// an error.
//
// See: https://developer.apple.com/documentation/CoreData/NSDerivedAttributeDescription/derivationExpression
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
func (d NSDerivedAttributeDescription) DerivationExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivationExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}
func (d NSDerivedAttributeDescription) SetDerivationExpression(value foundation.NSExpression) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivationExpression:"), value)
}
