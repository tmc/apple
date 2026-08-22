// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSRelationshipDescription] class.
var (
	_NSRelationshipDescriptionClass     NSRelationshipDescriptionClass
	_NSRelationshipDescriptionClassOnce sync.Once
)

func getNSRelationshipDescriptionClass() NSRelationshipDescriptionClass {
	_NSRelationshipDescriptionClassOnce.Do(func() {
		_NSRelationshipDescriptionClass = NSRelationshipDescriptionClass{class: objc.GetClass("NSRelationshipDescription")}
	})
	return _NSRelationshipDescriptionClass
}

// GetNSRelationshipDescriptionClass returns the class object for NSRelationshipDescription.
func GetNSRelationshipDescriptionClass() NSRelationshipDescriptionClass {
	return getNSRelationshipDescriptionClass()
}

type NSRelationshipDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSRelationshipDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRelationshipDescriptionClass) Alloc() NSRelationshipDescription {
	rv := objc.Send[NSRelationshipDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of a relationship between two entities.
//
// # Overview
//
// [NSRelationshipDescription] provides additional attributes that are
// specific to modeling a relationship between two entities. For the common
// attributes of all property types, see [NSPropertyDescription].
//
// For example, use this class to define a relationship’s cardinality —
// the number of managed objects the relationship can reference.
//
// - For a to-one relationship, set [NSRelationshipDescription.MaxCount] to
// `1`. - For a to-many relationship, set [NSRelationshipDescription.MaxCount]
// to a number greater than `1` to impose an upper limit; otherwise, use `0`
// to allow an unlimited number of referenced objects.
//
// At runtime, you can modify a relationship description until you associate
// its owning managed object model with a persistent store coordinator. If you
// attempt to modify the model after you associate it, Core Data throws an
// exception. To modify a model that’s in use, create and modify a copy and
// then discard any objects that belong to the original model.
//
// # Configuring the Destination
//
//   - [NSRelationshipDescription.InverseRelationship]: The relationship that represents the inverse of the current relationship.
//   - [NSRelationshipDescription.SetInverseRelationship]
//   - [NSRelationshipDescription.DestinationEntity]: The type of object the relationship contains.
//   - [NSRelationshipDescription.SetDestinationEntity]
//   - [NSRelationshipDescription.IsOrdered]: A Boolean value that determines whether the relationship preserves the order of the referenced managed objects.
//   - [NSRelationshipDescription.SetOrdered]
//
// # Configuring Cardinality
//
//   - [NSRelationshipDescription.IsToMany]: Returns a Boolean value that indicates whether the relationship can contain many managed objects.
//   - [NSRelationshipDescription.MinCount]: The minimum number of managed objects the relationship can reference.
//   - [NSRelationshipDescription.SetMinCount]
//   - [NSRelationshipDescription.MaxCount]: The maximum number of managed objects the relationship can reference.
//   - [NSRelationshipDescription.SetMaxCount]
//
// # Configuring Delete Behavior
//
//   - [NSRelationshipDescription.DeleteRule]: The rule to apply when you delete the relationship’s owning managed object.
//   - [NSRelationshipDescription.SetDeleteRule]
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription
type NSRelationshipDescription struct {
	NSPropertyDescription
}

// NSRelationshipDescriptionFromID constructs a [NSRelationshipDescription] from an objc.ID.
//
// A description of a relationship between two entities.
func NSRelationshipDescriptionFromID(id objc.ID) NSRelationshipDescription {
	return NSRelationshipDescription{NSPropertyDescription: NSPropertyDescriptionFromID(id)}
}

// NOTE: NSRelationshipDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRelationshipDescription] class.
//
// # Configuring the Destination
//
//   - [INSRelationshipDescription.InverseRelationship]: The relationship that represents the inverse of the current relationship.
//   - [INSRelationshipDescription.SetInverseRelationship]
//   - [INSRelationshipDescription.DestinationEntity]: The type of object the relationship contains.
//   - [INSRelationshipDescription.SetDestinationEntity]
//   - [INSRelationshipDescription.IsOrdered]: A Boolean value that determines whether the relationship preserves the order of the referenced managed objects.
//   - [INSRelationshipDescription.SetOrdered]
//
// # Configuring Cardinality
//
//   - [INSRelationshipDescription.IsToMany]: Returns a Boolean value that indicates whether the relationship can contain many managed objects.
//   - [INSRelationshipDescription.MinCount]: The minimum number of managed objects the relationship can reference.
//   - [INSRelationshipDescription.SetMinCount]
//   - [INSRelationshipDescription.MaxCount]: The maximum number of managed objects the relationship can reference.
//   - [INSRelationshipDescription.SetMaxCount]
//
// # Configuring Delete Behavior
//
//   - [INSRelationshipDescription.DeleteRule]: The rule to apply when you delete the relationship’s owning managed object.
//   - [INSRelationshipDescription.SetDeleteRule]
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription
type INSRelationshipDescription interface {
	INSPropertyDescription

	// Topic: Configuring the Destination

	// The relationship that represents the inverse of the current relationship.
	InverseRelationship() INSRelationshipDescription
	SetInverseRelationship(value INSRelationshipDescription)
	// The type of object the relationship contains.
	DestinationEntity() INSEntityDescription
	SetDestinationEntity(value INSEntityDescription)
	// A Boolean value that determines whether the relationship preserves the order of the referenced managed objects.
	IsOrdered() bool
	SetOrdered(value bool)

	// Topic: Configuring Cardinality

	// Returns a Boolean value that indicates whether the relationship can contain many managed objects.
	IsToMany() bool
	// The minimum number of managed objects the relationship can reference.
	MinCount() uint
	SetMinCount(value uint)
	// The maximum number of managed objects the relationship can reference.
	MaxCount() uint
	SetMaxCount(value uint)

	// Topic: Configuring Delete Behavior

	// The rule to apply when you delete the relationship’s owning managed object.
	DeleteRule() NSDeleteRule
	SetDeleteRule(value NSDeleteRule)
}

// Init initializes the instance.
func (r NSRelationshipDescription) Init() NSRelationshipDescription {
	rv := objc.Send[NSRelationshipDescription](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRelationshipDescription) Autorelease() NSRelationshipDescription {
	rv := objc.Send[NSRelationshipDescription](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRelationshipDescription creates a new NSRelationshipDescription instance.
func NewNSRelationshipDescription() NSRelationshipDescription {
	class := getNSRelationshipDescriptionClass()
	rv := objc.Send[NSRelationshipDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The relationship that represents the inverse of the current relationship.
//
// # Discussion
//
// The inverse relationship is the description of the current relationship
// from the destination entity’s perspective. For example, the inverse of a
// department’s relationship to an employee (a to-many relationship) is the
// employees’ relationship to the department (a to-one relationship).
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/inverseRelationship
func (r NSRelationshipDescription) InverseRelationship() INSRelationshipDescription {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("inverseRelationship"))
	return NSRelationshipDescriptionFromID(objc.ID(rv))
}
func (r NSRelationshipDescription) SetInverseRelationship(value INSRelationshipDescription) {
	objc.Send[struct{}](r.ID, objc.Sel("setInverseRelationship:"), value)
}

// The type of object the relationship contains.
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/destinationEntity
func (r NSRelationshipDescription) DestinationEntity() INSEntityDescription {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationEntity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}
func (r NSRelationshipDescription) SetDestinationEntity(value INSEntityDescription) {
	objc.Send[struct{}](r.ID, objc.Sel("setDestinationEntity:"), value)
}

// A Boolean value that determines whether the relationship preserves the
// order of the referenced managed objects.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/isOrdered
func (r NSRelationshipDescription) IsOrdered() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isOrdered"))
	return rv
}
func (r NSRelationshipDescription) SetOrdered(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setOrdered:"), value)
}

// Returns a Boolean value that indicates whether the relationship can contain
// many managed objects.
//
// # Discussion
//
// If [NSRelationshipDescription.MaxCount] is equal to `1`, implying a to-one
// relationship, this property returns false; otherwise, it returns true.
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/isToMany
func (r NSRelationshipDescription) IsToMany() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isToMany"))
	return rv
}

// The minimum number of managed objects the relationship can reference.
//
// # Discussion
//
// If you declare a relationship attribute as optional when defining your
// entities, the framework only enforces [NSRelationshipDescription.MinCount]
// and [NSRelationshipDescription.MaxCount] when that attribute is not `nil`.
//
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/minCount
func (r NSRelationshipDescription) MinCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("minCount"))
	return rv
}
func (r NSRelationshipDescription) SetMinCount(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMinCount:"), value)
}

// The maximum number of managed objects the relationship can reference.
//
// # Discussion
//
// If you declare a relationship attribute as optional when defining your
// entities, the framework only enforces [NSRelationshipDescription.MinCount]
// and [NSRelationshipDescription.MaxCount] when that attribute is not `nil`.
//
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/maxCount
func (r NSRelationshipDescription) MaxCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("maxCount"))
	return rv
}
func (r NSRelationshipDescription) SetMaxCount(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxCount:"), value)
}

// The rule to apply when you delete the relationship’s owning managed
// object.
//
// # Discussion
//
// The default value is [NSDeleteRule.nullifyDeleteRule]. For possible values,
// see [NSDeleteRule].
//
// See: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription/deleteRule
//
// [NSDeleteRule.nullifyDeleteRule]: https://developer.apple.com/documentation/CoreData/NSDeleteRule/nullifyDeleteRule
// [NSDeleteRule]: https://developer.apple.com/documentation/CoreData/NSDeleteRule
func (r NSRelationshipDescription) DeleteRule() NSDeleteRule {
	rv := objc.Send[NSDeleteRule](r.ID, objc.Sel("deleteRule"))
	return NSDeleteRule(rv)
}
func (r NSRelationshipDescription) SetDeleteRule(value NSDeleteRule) {
	objc.Send[struct{}](r.ID, objc.Sel("setDeleteRule:"), value)
}
