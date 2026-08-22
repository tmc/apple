// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPropertyDescription] class.
var (
	_NSPropertyDescriptionClass     NSPropertyDescriptionClass
	_NSPropertyDescriptionClassOnce sync.Once
)

func getNSPropertyDescriptionClass() NSPropertyDescriptionClass {
	_NSPropertyDescriptionClassOnce.Do(func() {
		_NSPropertyDescriptionClass = NSPropertyDescriptionClass{class: objc.GetClass("NSPropertyDescription")}
	})
	return _NSPropertyDescriptionClass
}

// GetNSPropertyDescriptionClass returns the class object for NSPropertyDescription.
func GetNSPropertyDescriptionClass() NSPropertyDescriptionClass {
	return getNSPropertyDescriptionClass()
}

type NSPropertyDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPropertyDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPropertyDescriptionClass) Alloc() NSPropertyDescription {
	rv := objc.Send[NSPropertyDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of a single property belonging to an entity.
//
// # Overview
//
// A property describes a single value within an object managed by the Core
// Data Framework. There are different types of property, each represented by
// a subclass which encapsulates the specific property behavior—see
// [NSAttributeDescription], [NSRelationshipDescription], and
// [NSFetchedPropertyDescription].
//
// Note that a property name cannot be the same as any no-parameter method
// name of [NSObject] or [NSManagedObject]. For example, you cannot give a
// property the name “description”. There are hundreds of methods on
// [NSObject] which may conflict with property names—and this list can grow
// without warning from frameworks or other libraries. You should avoid very
// general words (like “font”, and “color”) and words or phrases which
// overlap with Cocoa paradigms (such as “isEditing” and
// “objectSpecifier”).
//
// Properties—relationships as well as attributes—may be transient. A
// managed object context knows about transient properties and tracks changes
// made to them. Transient properties are ignored by the persistent store, and
// not just during saves: you cannot fetch using a predicate based on
// transients (although you can use transient properties to filter in memory
// yourself).
//
// # Editing Property Descriptions
//
// Property descriptions are editable until they are used by an object graph
// manager (such as a persistent store coordinator). This allows you to create
// or modify them dynamically. However, once a description is used (when the
// managed object model to which it belongs is associated with a persistent
// store coordinator), it must not (indeed cannot) be changed. This is
// enforced at runtime: any attempt to mutate a model or any of its
// sub-objects after the model is associated with a persistent store
// coordinator causes an exception to be thrown. If you need to modify a model
// that is in use, create a copy, modify the copy, and then discard the
// objects with the old model.
//
// # Accessing Features of a Property
//
//   - [NSPropertyDescription.Entity]: The entity description of the receiver.
//   - [NSPropertyDescription.IsOptional]: A Boolean value that indicates whether the receiver is optional.
//   - [NSPropertyDescription.SetOptional]
//   - [NSPropertyDescription.IsTransient]: A Boolean value that indicates whether the receiver is transient.
//   - [NSPropertyDescription.SetTransient]
//   - [NSPropertyDescription.Name]: The name of the receiver.
//   - [NSPropertyDescription.SetName]
//   - [NSPropertyDescription.UserInfo]: The user info dictionary of the receiver.
//   - [NSPropertyDescription.SetUserInfo]
//
// # Supporting Validation
//
//   - [NSPropertyDescription.ValidationPredicates]: The validation predicates of the receiver.
//   - [NSPropertyDescription.ValidationWarnings]: The error strings associated with the receiver’s validation predicates.
//   - [NSPropertyDescription.SetValidationPredicatesWithValidationWarnings]: Sets the validation predicates and warnings of the receiver.
//
// # Supporting Versioning
//
//   - [NSPropertyDescription.VersionHash]: The version hash for the receiver.
//   - [NSPropertyDescription.VersionHashModifier]: The version hash modifier for the receiver.
//   - [NSPropertyDescription.SetVersionHashModifier]
//   - [NSPropertyDescription.RenamingIdentifier]: The renaming identifier for the receiver.
//   - [NSPropertyDescription.SetRenamingIdentifier]
//
// # Specifying Spotlight Support
//
//   - [NSPropertyDescription.IsIndexedBySpotlight]: A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//   - [NSPropertyDescription.SetIndexedBySpotlight]
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription
type NSPropertyDescription struct {
	objectivec.Object
}

// NSPropertyDescriptionFromID constructs a [NSPropertyDescription] from an objc.ID.
//
// A description of a single property belonging to an entity.
func NSPropertyDescriptionFromID(id objc.ID) NSPropertyDescription {
	return NSPropertyDescription{objectivec.Object{ID: id}}
}

// NOTE: NSPropertyDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPropertyDescription] class.
//
// # Accessing Features of a Property
//
//   - [INSPropertyDescription.Entity]: The entity description of the receiver.
//   - [INSPropertyDescription.IsOptional]: A Boolean value that indicates whether the receiver is optional.
//   - [INSPropertyDescription.SetOptional]
//   - [INSPropertyDescription.IsTransient]: A Boolean value that indicates whether the receiver is transient.
//   - [INSPropertyDescription.SetTransient]
//   - [INSPropertyDescription.Name]: The name of the receiver.
//   - [INSPropertyDescription.SetName]
//   - [INSPropertyDescription.UserInfo]: The user info dictionary of the receiver.
//   - [INSPropertyDescription.SetUserInfo]
//
// # Supporting Validation
//
//   - [INSPropertyDescription.ValidationPredicates]: The validation predicates of the receiver.
//   - [INSPropertyDescription.ValidationWarnings]: The error strings associated with the receiver’s validation predicates.
//   - [INSPropertyDescription.SetValidationPredicatesWithValidationWarnings]: Sets the validation predicates and warnings of the receiver.
//
// # Supporting Versioning
//
//   - [INSPropertyDescription.VersionHash]: The version hash for the receiver.
//   - [INSPropertyDescription.VersionHashModifier]: The version hash modifier for the receiver.
//   - [INSPropertyDescription.SetVersionHashModifier]
//   - [INSPropertyDescription.RenamingIdentifier]: The renaming identifier for the receiver.
//   - [INSPropertyDescription.SetRenamingIdentifier]
//
// # Specifying Spotlight Support
//
//   - [INSPropertyDescription.IsIndexedBySpotlight]: A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
//   - [INSPropertyDescription.SetIndexedBySpotlight]
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription
type INSPropertyDescription interface {
	objectivec.IObject

	// Topic: Accessing Features of a Property

	// The entity description of the receiver.
	Entity() INSEntityDescription
	// A Boolean value that indicates whether the receiver is optional.
	IsOptional() bool
	SetOptional(value bool)
	// A Boolean value that indicates whether the receiver is transient.
	IsTransient() bool
	SetTransient(value bool)
	// The name of the receiver.
	Name() string
	SetName(value string)
	// The user info dictionary of the receiver.
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)

	// Topic: Supporting Validation

	// The validation predicates of the receiver.
	ValidationPredicates() []foundation.NSPredicate
	// The error strings associated with the receiver’s validation predicates.
	ValidationWarnings() foundation.INSArray
	// Sets the validation predicates and warnings of the receiver.
	SetValidationPredicatesWithValidationWarnings(validationPredicates []foundation.NSPredicate, validationWarnings []string)

	// Topic: Supporting Versioning

	// The version hash for the receiver.
	VersionHash() foundation.NSData
	// The version hash modifier for the receiver.
	VersionHashModifier() string
	SetVersionHashModifier(value string)
	// The renaming identifier for the receiver.
	RenamingIdentifier() string
	SetRenamingIdentifier(value string)

	// Topic: Specifying Spotlight Support

	// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index.
	IsIndexedBySpotlight() bool
	SetIndexedBySpotlight(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPropertyDescription) Init() NSPropertyDescription {
	rv := objc.Send[NSPropertyDescription](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPropertyDescription) Autorelease() NSPropertyDescription {
	rv := objc.Send[NSPropertyDescription](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPropertyDescription creates a new NSPropertyDescription instance.
func NewNSPropertyDescription() NSPropertyDescription {
	class := getNSPropertyDescriptionClass()
	rv := objc.Send[NSPropertyDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the validation predicates and warnings of the receiver.
//
// validationPredicates: An array containing the validation predicates for the receiver.
//
// validationWarnings: An array containing the validation warnings for the receiver.
//
// # Discussion
//
// The `validationPredicates` and `validationWarnings` arrays should contain
// the same number of elements, and corresponding elements should appear at
// the same index in each array.
//
// Instead of implementing individual validation methods, you can use this
// method to provide a list of predicates that are evaluated against the
// managed objects and a list of corresponding error messages (which can be
// localized).
//
// # Special Considerations
//
// This method raises an exception if the receiver’s model has been used by
// an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/setValidationPredicates(_:withValidationWarnings:)
func (p NSPropertyDescription) SetValidationPredicatesWithValidationWarnings(validationPredicates []foundation.NSPredicate, validationWarnings []string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setValidationPredicates:withValidationWarnings:"), objectivec.IObjectSliceToNSArray(validationPredicates), objectivec.StringSliceToNSArray(validationWarnings))
}
func (p NSPropertyDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The entity description of the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/entity
func (p NSPropertyDescription) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the receiver is optional.
//
// # Discussion
//
// true if the receiver is optional, otherwise false. The optionality flag
// specifies whether a property’s value can be `nil` before an object can be
// saved to a persistent store.
//
// # Special Considerations
//
// Setting this property raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isOptional
func (p NSPropertyDescription) IsOptional() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isOptional"))
	return rv
}
func (p NSPropertyDescription) SetOptional(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setOptional:"), value)
}

// A Boolean value that indicates whether the receiver is transient.
//
// # Discussion
//
// true if the receiver is transient, otherwise false. The transient flag
// specifies whether or not a property’s value is ignored when an object is
// saved to a persistent store. Transient properties are not saved to the
// persistent store, but are still managed for undo, redo, validation, and so
// on.
//
// # Special Considerations
//
// Setting this property raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isTransient
func (p NSPropertyDescription) IsTransient() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isTransient"))
	return rv
}
func (p NSPropertyDescription) SetTransient(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setTransient:"), value)
}

// The name of the receiver.
//
// # Discussion
//
// A property name cannot be the same as any no-parameter method name of
// [NSObject] or [NSManagedObject]. Since there are hundreds of methods on
// [NSObject] which may conflict with property names, you should avoid very
// general words (like “font”, and “color”) and words or phrases that
// overlap with Cocoa paradigms (such as “isEditing” and
// “objectSpecifier”).
//
// # Special Considerations
//
// Setting the name raises an exception if the receiver’s model has been
// used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p NSPropertyDescription) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPropertyDescription) SetName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setName:"), objc.String(value))
}

// The user info dictionary of the receiver.
//
// # Discussion
//
// Setting the user info raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/userInfo
func (p NSPropertyDescription) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p NSPropertyDescription) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setUserInfo:"), value)
}

// The validation predicates of the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/validationPredicates
func (p NSPropertyDescription) ValidationPredicates() []foundation.NSPredicate {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("validationPredicates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSPredicate {
		return foundation.NSPredicateFromID(id)
	})
}

// The error strings associated with the receiver’s validation predicates.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/validationWarnings
func (p NSPropertyDescription) ValidationWarnings() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("validationWarnings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// The version hash for the receiver.
//
// # Discussion
//
// The version hash is used to uniquely identify a property based on its
// configuration. The version hash uses only values which affect the
// persistence of data and the user-defined
// [NSPropertyDescription.VersionHashModifier] value. (The values which affect
// persistence are the name of the property, and the flags for `isOptional`,
// `isTransient`, and `isReadOnly`.) This value is stored as part of the
// version information in the metadata for stores, as well as a definition of
// a property involved in an [NSPropertyMapping] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/versionHash
func (p NSPropertyDescription) VersionHash() foundation.NSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("versionHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The version hash modifier for the receiver.
//
// # Discussion
//
// This value is included in the version hash for the property. You use it to
// mark or denote a property as being a different “version” than another
// even if all of the values which affect persistence are equal. (Such a
// difference is important in cases where the attributes of a property are
// unchanged but the format or content of its data are changed.)
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/versionHashModifier
func (p NSPropertyDescription) VersionHashModifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("versionHashModifier"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPropertyDescription) SetVersionHashModifier(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setVersionHashModifier:"), objc.String(value))
}

// The renaming identifier for the receiver.
//
// # Discussion
//
// This is used to resolve naming conflicts between models. When creating an
// entity mapping between entities in two managed object models, a source
// entity property and a destination entity property that share the same
// identifier indicate that a property mapping should be configured to migrate
// from the source to the destination. If unset, the identifier will return
// the property’s name.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/renamingIdentifier
func (p NSPropertyDescription) RenamingIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("renamingIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPropertyDescription) SetRenamingIdentifier(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setRenamingIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether Core Data adds the property’s
// value to the Core Spotlight index.
//
// # Discussion
//
// You can also set this property using the Index in Spotlight attribute in
// the Attributes inspector of the Core Data model editor.
//
// See: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/isIndexedBySpotlight
func (p NSPropertyDescription) IsIndexedBySpotlight() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isIndexedBySpotlight"))
	return rv
}
func (p NSPropertyDescription) SetIndexedBySpotlight(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexedBySpotlight:"), value)
}
