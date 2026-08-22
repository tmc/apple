// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFetchIndexElementDescription] class.
var (
	_NSFetchIndexElementDescriptionClass     NSFetchIndexElementDescriptionClass
	_NSFetchIndexElementDescriptionClassOnce sync.Once
)

func getNSFetchIndexElementDescriptionClass() NSFetchIndexElementDescriptionClass {
	_NSFetchIndexElementDescriptionClassOnce.Do(func() {
		_NSFetchIndexElementDescriptionClass = NSFetchIndexElementDescriptionClass{class: objc.GetClass("NSFetchIndexElementDescription")}
	})
	return _NSFetchIndexElementDescriptionClass
}

// GetNSFetchIndexElementDescriptionClass returns the class object for NSFetchIndexElementDescription.
func GetNSFetchIndexElementDescriptionClass() NSFetchIndexElementDescriptionClass {
	return getNSFetchIndexElementDescriptionClass()
}

type NSFetchIndexElementDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchIndexElementDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchIndexElementDescriptionClass) Alloc() NSFetchIndexElementDescription {
	rv := objc.Send[NSFetchIndexElementDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Description of an Index Element
//
// # Creating an Index Element Description
//
//   - [NSFetchIndexElementDescription.InitWithPropertyCollationType]: Creates an index element description using the specified property description and collation type.
//
// # Inspecting an Index Element Description
//
//   - [NSFetchIndexElementDescription.CollationType]: The type of collation that the index element uses, either binary or R-tree.
//   - [NSFetchIndexElementDescription.SetCollationType]
//   - [NSFetchIndexElementDescription.IndexDescription]
//   - [NSFetchIndexElementDescription.IsAscending]: A Boolean value that controls whether an index that supports direction is an ascending or descending index.
//   - [NSFetchIndexElementDescription.SetAscending]
//   - [NSFetchIndexElementDescription.Property]: A property description.
//   - [NSFetchIndexElementDescription.PropertyName]: The specified name in the property description.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription
type NSFetchIndexElementDescription struct {
	objectivec.Object
}

// NSFetchIndexElementDescriptionFromID constructs a [NSFetchIndexElementDescription] from an objc.ID.
//
// Description of an Index Element
func NSFetchIndexElementDescriptionFromID(id objc.ID) NSFetchIndexElementDescription {
	return NSFetchIndexElementDescription{objectivec.Object{ID: id}}
}

// NOTE: NSFetchIndexElementDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchIndexElementDescription] class.
//
// # Creating an Index Element Description
//
//   - [INSFetchIndexElementDescription.InitWithPropertyCollationType]: Creates an index element description using the specified property description and collation type.
//
// # Inspecting an Index Element Description
//
//   - [INSFetchIndexElementDescription.CollationType]: The type of collation that the index element uses, either binary or R-tree.
//   - [INSFetchIndexElementDescription.SetCollationType]
//   - [INSFetchIndexElementDescription.IndexDescription]
//   - [INSFetchIndexElementDescription.IsAscending]: A Boolean value that controls whether an index that supports direction is an ascending or descending index.
//   - [INSFetchIndexElementDescription.SetAscending]
//   - [INSFetchIndexElementDescription.Property]: A property description.
//   - [INSFetchIndexElementDescription.PropertyName]: The specified name in the property description.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription
type INSFetchIndexElementDescription interface {
	objectivec.IObject

	// Topic: Creating an Index Element Description

	// Creates an index element description using the specified property description and collation type.
	InitWithPropertyCollationType(property INSPropertyDescription, collationType NSFetchIndexElementType) NSFetchIndexElementDescription

	// Topic: Inspecting an Index Element Description

	// The type of collation that the index element uses, either binary or R-tree.
	CollationType() NSFetchIndexElementType
	SetCollationType(value NSFetchIndexElementType)
	IndexDescription() INSFetchIndexDescription
	// A Boolean value that controls whether an index that supports direction is an ascending or descending index.
	IsAscending() bool
	SetAscending(value bool)
	// A property description.
	Property() INSPropertyDescription
	// The specified name in the property description.
	PropertyName() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFetchIndexElementDescription) Init() NSFetchIndexElementDescription {
	rv := objc.Send[NSFetchIndexElementDescription](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchIndexElementDescription) Autorelease() NSFetchIndexElementDescription {
	rv := objc.Send[NSFetchIndexElementDescription](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchIndexElementDescription creates a new NSFetchIndexElementDescription instance.
func NewNSFetchIndexElementDescription() NSFetchIndexElementDescription {
	class := getNSFetchIndexElementDescriptionClass()
	rv := objc.Send[NSFetchIndexElementDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an index element description using the specified property
// description and collation type.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/init(property:collationType:)
func NewFetchIndexElementDescriptionWithPropertyCollationType(property INSPropertyDescription, collationType NSFetchIndexElementType) NSFetchIndexElementDescription {
	instance := getNSFetchIndexElementDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProperty:collationType:"), property, collationType)
	return NSFetchIndexElementDescriptionFromID(rv)
}

// Creates an index element description using the specified property
// description and collation type.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/init(property:collationType:)
func (f NSFetchIndexElementDescription) InitWithPropertyCollationType(property INSPropertyDescription, collationType NSFetchIndexElementType) NSFetchIndexElementDescription {
	rv := objc.Send[NSFetchIndexElementDescription](f.ID, objc.Sel("initWithProperty:collationType:"), property, collationType)
	return rv
}
func (f NSFetchIndexElementDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The type of collation that the index element uses, either binary or R-tree.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/collationType
func (f NSFetchIndexElementDescription) CollationType() NSFetchIndexElementType {
	rv := objc.Send[NSFetchIndexElementType](f.ID, objc.Sel("collationType"))
	return NSFetchIndexElementType(rv)
}
func (f NSFetchIndexElementDescription) SetCollationType(value NSFetchIndexElementType) {
	objc.Send[struct{}](f.ID, objc.Sel("setCollationType:"), value)
}

// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/indexDescription
func (f NSFetchIndexElementDescription) IndexDescription() INSFetchIndexDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("indexDescription"))
	return NSFetchIndexDescriptionFromID(objc.ID(rv))
}

// A Boolean value that controls whether an index that supports direction is
// an ascending or descending index.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/isAscending
func (f NSFetchIndexElementDescription) IsAscending() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isAscending"))
	return rv
}
func (f NSFetchIndexElementDescription) SetAscending(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setAscending:"), value)
}

// A property description.
//
// # Discussion
//
// This property may also be an [NSExpressionDescription] that expresses a
// function.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/property
func (f NSFetchIndexElementDescription) Property() INSPropertyDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("property"))
	return NSPropertyDescriptionFromID(objc.ID(rv))
}

// The specified name in the property description.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementDescription/propertyName
func (f NSFetchIndexElementDescription) PropertyName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("propertyName"))
	return foundation.NSStringFromID(rv).String()
}
