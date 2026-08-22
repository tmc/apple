// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCompositeAttributeDescription] class.
var (
	_NSCompositeAttributeDescriptionClass     NSCompositeAttributeDescriptionClass
	_NSCompositeAttributeDescriptionClassOnce sync.Once
)

func getNSCompositeAttributeDescriptionClass() NSCompositeAttributeDescriptionClass {
	_NSCompositeAttributeDescriptionClassOnce.Do(func() {
		_NSCompositeAttributeDescriptionClass = NSCompositeAttributeDescriptionClass{class: objc.GetClass("NSCompositeAttributeDescription")}
	})
	return _NSCompositeAttributeDescriptionClass
}

// GetNSCompositeAttributeDescriptionClass returns the class object for NSCompositeAttributeDescription.
func GetNSCompositeAttributeDescriptionClass() NSCompositeAttributeDescriptionClass {
	return getNSCompositeAttributeDescriptionClass()
}

type NSCompositeAttributeDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCompositeAttributeDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCompositeAttributeDescriptionClass) Alloc() NSCompositeAttributeDescription {
	rv := objc.Send[NSCompositeAttributeDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of an attribute that derives its value by composing other
// attributes.
//
// # Overview
//
// Composite attributes enable you to define and store complex data types, and
// then query, index, and apply constraints to those types. Model classes use
// dictionaries to represent those composites in-memory, where each dictionary
// contains keys corresponding to the names of the underlying attributes. You
// may use composite attributes anywhere you use standard attributes,
// including lightweight migrations and CloudKit, through
// [NSPersistentCloudKitContainer]. You can even nest composites inside other
// composites to create complex object hierarchies without additional model
// classes.
//
// In most scenarios, prefer to use Xcode’s model editor to add composite
// attributes to your entities and then regenerate your model classes.
// However, if you need to create composites dynamically at runtime, create an
// instance of this class and populate its
// [NSCompositeAttributeDescription.Elements] property with the necessary
// attribute descriptions.
//
// You can access a composite’s underlying attributes using namespaced key
// paths and property-like setters and getters, as the following example
// demonstrates:
//
// # Composing attributes
//
//   - [NSCompositeAttributeDescription.Elements]: The composed attribute descriptions.
//   - [NSCompositeAttributeDescription.SetElements]
//
// See: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription
type NSCompositeAttributeDescription struct {
	NSAttributeDescription
}

// NSCompositeAttributeDescriptionFromID constructs a [NSCompositeAttributeDescription] from an objc.ID.
//
// A description of an attribute that derives its value by composing other
// attributes.
func NSCompositeAttributeDescriptionFromID(id objc.ID) NSCompositeAttributeDescription {
	return NSCompositeAttributeDescription{NSAttributeDescription: NSAttributeDescriptionFromID(id)}
}

// NOTE: NSCompositeAttributeDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCompositeAttributeDescription] class.
//
// # Composing attributes
//
//   - [INSCompositeAttributeDescription.Elements]: The composed attribute descriptions.
//   - [INSCompositeAttributeDescription.SetElements]
//
// See: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription
type INSCompositeAttributeDescription interface {
	INSAttributeDescription

	// Topic: Composing attributes

	// The composed attribute descriptions.
	Elements() []NSAttributeDescription
	SetElements(value []NSAttributeDescription)
}

// Init initializes the instance.
func (c NSCompositeAttributeDescription) Init() NSCompositeAttributeDescription {
	rv := objc.Send[NSCompositeAttributeDescription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCompositeAttributeDescription) Autorelease() NSCompositeAttributeDescription {
	rv := objc.Send[NSCompositeAttributeDescription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCompositeAttributeDescription creates a new NSCompositeAttributeDescription instance.
func NewNSCompositeAttributeDescription() NSCompositeAttributeDescription {
	class := getNSCompositeAttributeDescriptionClass()
	rv := objc.Send[NSCompositeAttributeDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The composed attribute descriptions.
//
// See: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription/elements
func (c NSCompositeAttributeDescription) Elements() []NSAttributeDescription {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("elements"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAttributeDescription {
		return NSAttributeDescriptionFromID(id)
	})
}
func (c NSCompositeAttributeDescription) SetElements(value []NSAttributeDescription) {
	objc.Send[struct{}](c.ID, objc.Sel("setElements:"), objectivec.IObjectSliceToNSArray(value))
}
