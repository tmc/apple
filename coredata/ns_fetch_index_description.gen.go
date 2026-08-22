// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFetchIndexDescription] class.
var (
	_NSFetchIndexDescriptionClass     NSFetchIndexDescriptionClass
	_NSFetchIndexDescriptionClassOnce sync.Once
)

func getNSFetchIndexDescriptionClass() NSFetchIndexDescriptionClass {
	_NSFetchIndexDescriptionClassOnce.Do(func() {
		_NSFetchIndexDescriptionClass = NSFetchIndexDescriptionClass{class: objc.GetClass("NSFetchIndexDescription")}
	})
	return _NSFetchIndexDescriptionClass
}

// GetNSFetchIndexDescriptionClass returns the class object for NSFetchIndexDescription.
func GetNSFetchIndexDescriptionClass() NSFetchIndexDescriptionClass {
	return getNSFetchIndexDescriptionClass()
}

type NSFetchIndexDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchIndexDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchIndexDescriptionClass) Alloc() NSFetchIndexDescription {
	rv := objc.Send[NSFetchIndexDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The description of the index.
//
// # Creating an Index Description
//
//   - [NSFetchIndexDescription.InitWithNameElements]: Creates a fetch index description using the specified name and element descriptions.
//
// # Inspecting an Index Description
//
//   - [NSFetchIndexDescription.Elements]: An array of fetch index element descriptions.
//   - [NSFetchIndexDescription.SetElements]
//   - [NSFetchIndexDescription.Entity]: The entity description for the fetch index description.
//   - [NSFetchIndexDescription.Name]: The name of the fetch index description.
//   - [NSFetchIndexDescription.SetName]
//   - [NSFetchIndexDescription.PartialIndexPredicate]: A predicate that selects rows for indexing, if the index is a partial index.
//   - [NSFetchIndexDescription.SetPartialIndexPredicate]
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription
type NSFetchIndexDescription struct {
	objectivec.Object
}

// NSFetchIndexDescriptionFromID constructs a [NSFetchIndexDescription] from an objc.ID.
//
// The description of the index.
func NSFetchIndexDescriptionFromID(id objc.ID) NSFetchIndexDescription {
	return NSFetchIndexDescription{objectivec.Object{ID: id}}
}

// NOTE: NSFetchIndexDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchIndexDescription] class.
//
// # Creating an Index Description
//
//   - [INSFetchIndexDescription.InitWithNameElements]: Creates a fetch index description using the specified name and element descriptions.
//
// # Inspecting an Index Description
//
//   - [INSFetchIndexDescription.Elements]: An array of fetch index element descriptions.
//   - [INSFetchIndexDescription.SetElements]
//   - [INSFetchIndexDescription.Entity]: The entity description for the fetch index description.
//   - [INSFetchIndexDescription.Name]: The name of the fetch index description.
//   - [INSFetchIndexDescription.SetName]
//   - [INSFetchIndexDescription.PartialIndexPredicate]: A predicate that selects rows for indexing, if the index is a partial index.
//   - [INSFetchIndexDescription.SetPartialIndexPredicate]
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription
type INSFetchIndexDescription interface {
	objectivec.IObject

	// Topic: Creating an Index Description

	// Creates a fetch index description using the specified name and element descriptions.
	InitWithNameElements(name string, elements []NSFetchIndexElementDescription) NSFetchIndexDescription

	// Topic: Inspecting an Index Description

	// An array of fetch index element descriptions.
	Elements() []NSFetchIndexElementDescription
	SetElements(value []NSFetchIndexElementDescription)
	// The entity description for the fetch index description.
	Entity() INSEntityDescription
	// The name of the fetch index description.
	Name() string
	SetName(value string)
	// A predicate that selects rows for indexing, if the index is a partial index.
	PartialIndexPredicate() foundation.NSPredicate
	SetPartialIndexPredicate(value foundation.NSPredicate)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFetchIndexDescription) Init() NSFetchIndexDescription {
	rv := objc.Send[NSFetchIndexDescription](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchIndexDescription) Autorelease() NSFetchIndexDescription {
	rv := objc.Send[NSFetchIndexDescription](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchIndexDescription creates a new NSFetchIndexDescription instance.
func NewNSFetchIndexDescription() NSFetchIndexDescription {
	class := getNSFetchIndexDescriptionClass()
	rv := objc.Send[NSFetchIndexDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fetch index description using the specified name and element
// descriptions.
//
// name: The name of the fetch index description.
//
// elements: An array of fetch index element descriptions.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/init(name:elements:)
func NewFetchIndexDescriptionWithNameElements(name string, elements []NSFetchIndexElementDescription) NSFetchIndexDescription {
	instance := getNSFetchIndexDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:elements:"), objc.String(name), objectivec.IObjectSliceToNSArray(elements))
	return NSFetchIndexDescriptionFromID(rv)
}

// Creates a fetch index description using the specified name and element
// descriptions.
//
// name: The name of the fetch index description.
//
// elements: An array of fetch index element descriptions.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/init(name:elements:)
func (f NSFetchIndexDescription) InitWithNameElements(name string, elements []NSFetchIndexElementDescription) NSFetchIndexDescription {
	rv := objc.Send[NSFetchIndexDescription](f.ID, objc.Sel("initWithName:elements:"), objc.String(name), objectivec.IObjectSliceToNSArray(elements))
	return rv
}
func (f NSFetchIndexDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An array of fetch index element descriptions.
//
// # Discussion
//
// Setting this property to an invalid value throws an exception, such as when
// the new value includes both R-tree and non-R-tree elements.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/elements
func (f NSFetchIndexDescription) Elements() []NSFetchIndexElementDescription {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("elements"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFetchIndexElementDescription {
		return NSFetchIndexElementDescriptionFromID(id)
	})
}
func (f NSFetchIndexDescription) SetElements(value []NSFetchIndexElementDescription) {
	objc.Send[struct{}](f.ID, objc.Sel("setElements:"), objectivec.IObjectSliceToNSArray(value))
}

// The entity description for the fetch index description.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/entity
func (f NSFetchIndexDescription) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// The name of the fetch index description.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/name
func (f NSFetchIndexDescription) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (f NSFetchIndexDescription) SetName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setName:"), objc.String(value))
}

// A predicate that selects rows for indexing, if the index is a partial
// index.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexDescription/partialIndexPredicate
func (f NSFetchIndexDescription) PartialIndexPredicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("partialIndexPredicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (f NSFetchIndexDescription) SetPartialIndexPredicate(value foundation.NSPredicate) {
	objc.Send[struct{}](f.ID, objc.Sel("setPartialIndexPredicate:"), value)
}
