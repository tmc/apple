// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSEntityDescription] class.
var (
	_NSEntityDescriptionClass     NSEntityDescriptionClass
	_NSEntityDescriptionClassOnce sync.Once
)

func getNSEntityDescriptionClass() NSEntityDescriptionClass {
	_NSEntityDescriptionClassOnce.Do(func() {
		_NSEntityDescriptionClass = NSEntityDescriptionClass{class: objc.GetClass("NSEntityDescription")}
	})
	return _NSEntityDescriptionClass
}

// GetNSEntityDescriptionClass returns the class object for NSEntityDescription.
func GetNSEntityDescriptionClass() NSEntityDescriptionClass {
	return getNSEntityDescriptionClass()
}

type NSEntityDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSEntityDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSEntityDescriptionClass) Alloc() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of a Core Data entity.
//
// # Overview
//
// Entities are to managed objects what [Class] is to `id`, or — to use a
// database analogy — what tables are to rows. An instance specifies the
// entity’s name, its attributes and relationships (as instances of
// [NSAttributeDescription] and [NSRelationshipDescription]) and the class
// that represents it. Instances of that class correspond to entries in the
// associated persistent store. As a minimum, an entity description requires:
//
// - A name.
// - The class name of the corresponding managed object.
//
// If you don’t specify a class name, the framework uses [NSManagedObject].
//
// You define entities in a managed object model (an instance of
// [NSManagedObjectModel]) using Xcode’s data modeling tool. Core Data uses
// [NSEntityDescription] to map entries in the persistent store to managed
// objects in your app. It’s unlikely you’ll interact with entity
// descriptions directly unless you’re specifically working with models.
// [NSEntityDescription] provides a user dictionary for you to store any
// related, app-specific information.
//
// # Editing entity descriptions
//
// Entity descriptions are editable until an object graph manager uses them,
// which allows you to create or modify descriptions dynamically. However,
// once you associate the description’s managed object model with a
// persistent store coordinator, you can no longer modify it. The framework
// enforces this rule at runtime; any attempt to mutate the model, or any of
// its child objects, after you associate it with a persistent store
// coordinator results in an exception. If you need to modify a model that’s
// in use, create a copy of that model, modify it, and then discard the stale
// model.
//
// If you want to create an entity hierarchy, consider the relevant API. You
// can only set an entity’s [NSEntityDescription.Subentities], not an
// entity’s super-entity. To set an entity’s super-entity, set an array of
// subentities on the super entity that includes the desired entity; the
// entity hierarchy is built top-down.
//
// # Using entity descriptions in dictionaries
//
// The `copy` method of [NSEntityDescription] returns an entity such that:
//
// Since [NSDictionary] copies its keys and requires that keys both conform to
// the [NSCopying] protocol and have a property that `copy` returns an object
// for where the source and the copy are equal, don’t use entities as keys
// in a dictionary. Instead, use either the entity’s name as the key or use
// an [NSMapTable] with retain callbacks.
//
// # Fast enumeration
//
// [NSEntityDescription] implements the [NSFastEnumeration] protocol. Use this
// to enumerate over an entity’s properties, as the following example
// illustrates.
//
// # Getting descriptive information
//
//   - [NSEntityDescription.Name]: The entity name of the receiver.
//   - [NSEntityDescription.SetName]
//   - [NSEntityDescription.ManagedObjectModel]: The managed object model with which the receiver is associated.
//   - [NSEntityDescription.ManagedObjectClassName]: The name of the class that represents the receiver’s entity.
//   - [NSEntityDescription.SetManagedObjectClassName]
//   - [NSEntityDescription.RenamingIdentifier]: The renaming identifier for the receiver.
//   - [NSEntityDescription.SetRenamingIdentifier]
//   - [NSEntityDescription.IsAbstract]: A Boolean value that indicates whether the receiver represents an abstract entity.
//   - [NSEntityDescription.SetAbstract]
//   - [NSEntityDescription.UserInfo]: The user info dictionary of the receiver.
//   - [NSEntityDescription.SetUserInfo]
//   - [NSEntityDescription.CoreSpotlightDisplayNameExpression]: The expression that computes the CoreSpotlight display name for instances of the entity.
//   - [NSEntityDescription.SetCoreSpotlightDisplayNameExpression]
//
// # Managing inheritance
//
//   - [NSEntityDescription.SubentitiesByName]: A dictionary containing the receiver’s sub-entities.
//   - [NSEntityDescription.Subentities]: An array containing the sub-entities of the receiver.
//   - [NSEntityDescription.SetSubentities]
//   - [NSEntityDescription.Superentity]: The super-entity of the receiver.
//   - [NSEntityDescription.IsKindOfEntity]: Returns a Boolean value that indicates whether the receiver is a sub-entity of another given entity.
//
// # Working with properties
//
//   - [NSEntityDescription.PropertiesByName]: A dictionary containing the properties of the receiver.
//   - [NSEntityDescription.Properties]: An array containing the properties of the receiver.
//   - [NSEntityDescription.SetProperties]
//   - [NSEntityDescription.AttributesByName]: The attributes of the receiver in a dictionary.
//   - [NSEntityDescription.RelationshipsByName]: The relationships of the receiver in a dictionary.
//   - [NSEntityDescription.RelationshipsWithDestinationEntity]: Returns an array containing the relationships of the receiver where the entity description of the relationship is a given entity.
//
// # Configuring indexes and constraints
//
//   - [NSEntityDescription.Indexes]: An array of fetch index descriptions for the entity.
//   - [NSEntityDescription.SetIndexes]
//   - [NSEntityDescription.UniquenessConstraints]: An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
//   - [NSEntityDescription.SetUniquenessConstraints]
//
// # Managing versioning
//
//   - [NSEntityDescription.VersionHash]: The version hash for the receiver.
//   - [NSEntityDescription.VersionHashModifier]: The version hash modifier for the receiver.
//   - [NSEntityDescription.SetVersionHashModifier]
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription
//
// [NSCopying]: https://developer.apple.com/documentation/Foundation/NSCopying
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSFastEnumeration]: https://developer.apple.com/documentation/Foundation/NSFastEnumeration
// [NSMapTable]: https://developer.apple.com/documentation/Foundation/NSMapTable
type NSEntityDescription struct {
	objectivec.Object
}

// NSEntityDescriptionFromID constructs a [NSEntityDescription] from an objc.ID.
//
// A description of a Core Data entity.
func NSEntityDescriptionFromID(id objc.ID) NSEntityDescription {
	return NSEntityDescription{objectivec.Object{ID: id}}
}

// NOTE: NSEntityDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSEntityDescription] class.
//
// # Getting descriptive information
//
//   - [INSEntityDescription.Name]: The entity name of the receiver.
//   - [INSEntityDescription.SetName]
//   - [INSEntityDescription.ManagedObjectModel]: The managed object model with which the receiver is associated.
//   - [INSEntityDescription.ManagedObjectClassName]: The name of the class that represents the receiver’s entity.
//   - [INSEntityDescription.SetManagedObjectClassName]
//   - [INSEntityDescription.RenamingIdentifier]: The renaming identifier for the receiver.
//   - [INSEntityDescription.SetRenamingIdentifier]
//   - [INSEntityDescription.IsAbstract]: A Boolean value that indicates whether the receiver represents an abstract entity.
//   - [INSEntityDescription.SetAbstract]
//   - [INSEntityDescription.UserInfo]: The user info dictionary of the receiver.
//   - [INSEntityDescription.SetUserInfo]
//   - [INSEntityDescription.CoreSpotlightDisplayNameExpression]: The expression that computes the CoreSpotlight display name for instances of the entity.
//   - [INSEntityDescription.SetCoreSpotlightDisplayNameExpression]
//
// # Managing inheritance
//
//   - [INSEntityDescription.SubentitiesByName]: A dictionary containing the receiver’s sub-entities.
//   - [INSEntityDescription.Subentities]: An array containing the sub-entities of the receiver.
//   - [INSEntityDescription.SetSubentities]
//   - [INSEntityDescription.Superentity]: The super-entity of the receiver.
//   - [INSEntityDescription.IsKindOfEntity]: Returns a Boolean value that indicates whether the receiver is a sub-entity of another given entity.
//
// # Working with properties
//
//   - [INSEntityDescription.PropertiesByName]: A dictionary containing the properties of the receiver.
//   - [INSEntityDescription.Properties]: An array containing the properties of the receiver.
//   - [INSEntityDescription.SetProperties]
//   - [INSEntityDescription.AttributesByName]: The attributes of the receiver in a dictionary.
//   - [INSEntityDescription.RelationshipsByName]: The relationships of the receiver in a dictionary.
//   - [INSEntityDescription.RelationshipsWithDestinationEntity]: Returns an array containing the relationships of the receiver where the entity description of the relationship is a given entity.
//
// # Configuring indexes and constraints
//
//   - [INSEntityDescription.Indexes]: An array of fetch index descriptions for the entity.
//   - [INSEntityDescription.SetIndexes]
//   - [INSEntityDescription.UniquenessConstraints]: An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
//   - [INSEntityDescription.SetUniquenessConstraints]
//
// # Managing versioning
//
//   - [INSEntityDescription.VersionHash]: The version hash for the receiver.
//   - [INSEntityDescription.VersionHashModifier]: The version hash modifier for the receiver.
//   - [INSEntityDescription.SetVersionHashModifier]
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription
type INSEntityDescription interface {
	objectivec.IObject

	// Topic: Getting descriptive information

	// The entity name of the receiver.
	Name() string
	SetName(value string)
	// The managed object model with which the receiver is associated.
	ManagedObjectModel() INSManagedObjectModel
	// The name of the class that represents the receiver’s entity.
	ManagedObjectClassName() string
	SetManagedObjectClassName(value string)
	// The renaming identifier for the receiver.
	RenamingIdentifier() string
	SetRenamingIdentifier(value string)
	// A Boolean value that indicates whether the receiver represents an abstract entity.
	IsAbstract() bool
	SetAbstract(value bool)
	// The user info dictionary of the receiver.
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)
	// The expression that computes the CoreSpotlight display name for instances of the entity.
	CoreSpotlightDisplayNameExpression() foundation.NSExpression
	SetCoreSpotlightDisplayNameExpression(value foundation.NSExpression)

	// Topic: Managing inheritance

	// A dictionary containing the receiver’s sub-entities.
	SubentitiesByName() foundation.INSDictionary
	// An array containing the sub-entities of the receiver.
	Subentities() []NSEntityDescription
	SetSubentities(value []NSEntityDescription)
	// The super-entity of the receiver.
	Superentity() INSEntityDescription
	// Returns a Boolean value that indicates whether the receiver is a sub-entity of another given entity.
	IsKindOfEntity(entity INSEntityDescription) bool

	// Topic: Working with properties

	// A dictionary containing the properties of the receiver.
	PropertiesByName() foundation.INSDictionary
	// An array containing the properties of the receiver.
	Properties() []NSPropertyDescription
	SetProperties(value []NSPropertyDescription)
	// The attributes of the receiver in a dictionary.
	AttributesByName() foundation.INSDictionary
	// The relationships of the receiver in a dictionary.
	RelationshipsByName() foundation.INSDictionary
	// Returns an array containing the relationships of the receiver where the entity description of the relationship is a given entity.
	RelationshipsWithDestinationEntity(entity INSEntityDescription) []NSRelationshipDescription

	// Topic: Configuring indexes and constraints

	// An array of fetch index descriptions for the entity.
	Indexes() []NSFetchIndexDescription
	SetIndexes(value []NSFetchIndexDescription)
	// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity.
	UniquenessConstraints() []foundation.NSArray
	SetUniquenessConstraints(value []foundation.NSArray)

	// Topic: Managing versioning

	// The version hash for the receiver.
	VersionHash() foundation.NSData
	// The version hash modifier for the receiver.
	VersionHashModifier() string
	SetVersionHashModifier(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (e NSEntityDescription) Init() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSEntityDescription) Autorelease() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSEntityDescription creates a new NSEntityDescription instance.
func NewNSEntityDescription() NSEntityDescription {
	class := getNSEntityDescriptionClass()
	rv := objc.Send[NSEntityDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver is a sub-entity
// of another given entity.
//
// entity: An entity.
//
// # Return Value
//
// true if the receiver is a sub-entity of `entity`, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/isKindOf(entity:)
func (e NSEntityDescription) IsKindOfEntity(entity INSEntityDescription) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isKindOfEntity:"), entity)
	return rv
}

// Returns an array containing the relationships of the receiver where the
// entity description of the relationship is a given entity.
//
// entity: An entity description.
//
// # Return Value
//
// An array containing the relationships of the receiver where the entity
// description of the relationship is `entity`. Elements in the array are
// instances of [NSRelationshipDescription].
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/relationships(forDestination:)
func (e NSEntityDescription) RelationshipsWithDestinationEntity(entity INSEntityDescription) []NSRelationshipDescription {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("relationshipsWithDestinationEntity:"), entity)
	return objc.ConvertSlice(rv, func(id objc.ID) NSRelationshipDescription {
		return NSRelationshipDescriptionFromID(id)
	})
}
func (e NSEntityDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates, configures, and returns an instance of the class for the entity
// with a given name.
//
// entityName: The name of an entity.
//
// context: The managed object context to use.
//
// # Return Value
//
// A new, autoreleased, fully configured instance of the class for the entity
// named `entityName`. The instance has its entity description set and is
// inserted it into `context`.
//
// # Discussion
//
// This method makes it easy for you to create instances of a given entity
// without worrying about the details of managed object creation. The method
// is conceptually similar to the following code example.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/insertNewObject(forEntityName:into:)
func (_NSEntityDescriptionClass NSEntityDescriptionClass) InsertNewObjectForEntityForNameInManagedObjectContext(entityName string, context INSManagedObjectContext) NSManagedObject {
	rv := objc.Send[objc.ID](objc.ID(_NSEntityDescriptionClass.class), objc.Sel("insertNewObjectForEntityForName:inManagedObjectContext:"), objc.String(entityName), context)
	return NSManagedObjectFromID(rv)
}

// Returns the entity with the specified name from the managed object model
// associated with the specified managed object context’s persistent store
// coordinator.
//
// entityName: The name of an entity.
//
// context: The managed object context to use. Must not be `nil`.
//
// # Return Value
//
// The entity with the specified name from the managed object model associated
// with `context`’s persistent store coordinator.
//
// # Discussion
//
// Raises [internalInconsistencyException] if `context` is `nil`.
//
// This method is functionally equivalent to the following code example.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/entity(forEntityName:in:)
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
func (_NSEntityDescriptionClass NSEntityDescriptionClass) EntityForNameInManagedObjectContext(entityName string, context INSManagedObjectContext) NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSEntityDescriptionClass.class), objc.Sel("entityForName:inManagedObjectContext:"), objc.String(entityName), context)
	return NSEntityDescriptionFromID(rv)
}

// The entity name of the receiver.
//
// # Discussion
//
// Setting the name raises an exception if the receiver’s model has been
// used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/name
func (e NSEntityDescription) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityDescription) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}

// The managed object model with which the receiver is associated.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/managedObjectModel
func (e NSEntityDescription) ManagedObjectModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("managedObjectModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The name of the class that represents the receiver’s entity.
//
// # Discussion
//
// The class specified by `name` must [NSManagedObject] or a subclass of
// [NSManagedObject].
//
// # Special Considerations
//
// Setting the class name raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/managedObjectClassName
func (e NSEntityDescription) ManagedObjectClassName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("managedObjectClassName"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityDescription) SetManagedObjectClassName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setManagedObjectClassName:"), objc.String(value))
}

// The renaming identifier for the receiver.
//
// # Discussion
//
// The renaming identifier is used to resolve naming conflicts between models.
// When creating a mapping model between two managed object models, a source
// entity and a destination entity that share the same identifier indicate
// that an entity mapping should be configured to migrate from the source to
// the destination.
//
// If you do not set this value, the identifier will return the entity’s
// name.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/renamingIdentifier
func (e NSEntityDescription) RenamingIdentifier() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("renamingIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityDescription) SetRenamingIdentifier(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setRenamingIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether the receiver represents an abstract
// entity.
//
// # Return Value
//
// true if the receiver represents an abstract entity, otherwise false.
//
// # Discussion
//
// true if the receiver represents an abstract entity, otherwise false. An
// abstract entity might be Shape, with concrete sub-entities such as
// Rectangle, Triangle, and Circle.
//
// # Special Considerations
//
// Setting whether an entity is abstract raises an exception if the
// receiver’s model has been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/isAbstract
func (e NSEntityDescription) IsAbstract() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isAbstract"))
	return rv
}
func (e NSEntityDescription) SetAbstract(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setAbstract:"), value)
}

// The user info dictionary of the receiver.
//
// # Discussion
//
// Setting the user info dictionary raises an exception if the receiver’s
// model has been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/userInfo
func (e NSEntityDescription) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e NSEntityDescription) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setUserInfo:"), value)
}

// The expression that computes the CoreSpotlight display name for instances
// of the entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/coreSpotlightDisplayNameExpression
func (e NSEntityDescription) CoreSpotlightDisplayNameExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("coreSpotlightDisplayNameExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}
func (e NSEntityDescription) SetCoreSpotlightDisplayNameExpression(value foundation.NSExpression) {
	objc.Send[struct{}](e.ID, objc.Sel("setCoreSpotlightDisplayNameExpression:"), value)
}

// A dictionary containing the receiver’s sub-entities.
//
// # Return Value
//
// The keys in the dictionary are the sub-entity names, the corresponding
// values are instances of [NSEntityDescription].
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/subentitiesByName
func (e NSEntityDescription) SubentitiesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("subentitiesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// An array containing the sub-entities of the receiver.
//
// # Discussion
//
// The sub-entities are instances of [NSEntityDescription].
//
// # Special Considerations
//
// Setting the sub-entities raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/subentities
func (e NSEntityDescription) Subentities() []NSEntityDescription {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("subentities"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSEntityDescription {
		return NSEntityDescriptionFromID(id)
	})
}
func (e NSEntityDescription) SetSubentities(value []NSEntityDescription) {
	objc.Send[struct{}](e.ID, objc.Sel("setSubentities:"), objectivec.IObjectSliceToNSArray(value))
}

// The super-entity of the receiver.
//
// # Discussion
//
// If the receiver has no super-entity, returns `nil`.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/superentity
func (e NSEntityDescription) Superentity() INSEntityDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("superentity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// A dictionary containing the properties of the receiver.
//
// # Discussion
//
// The keys in the dictionary are the property names and the values are
// instances of [NSAttributeDescription] and/or [NSRelationshipDescription].
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/propertiesByName
func (e NSEntityDescription) PropertiesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("propertiesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// An array containing the properties of the receiver.
//
// # Discussion
//
// The elements in the array are instances of [NSAttributeDescription],
// [NSRelationshipDescription], and/or [NSFetchedPropertyDescription].
//
// # Special Considerations
//
// Setting the properties raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/properties
func (e NSEntityDescription) Properties() []NSPropertyDescription {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("properties"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPropertyDescription {
		return NSPropertyDescriptionFromID(id)
	})
}
func (e NSEntityDescription) SetProperties(value []NSPropertyDescription) {
	objc.Send[struct{}](e.ID, objc.Sel("setProperties:"), objectivec.IObjectSliceToNSArray(value))
}

// The attributes of the receiver in a dictionary.
//
// # Discussion
//
// The keys in the dictionary are the attribute names and the values are
// instances of [NSAttributeDescription]. .
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/attributesByName
func (e NSEntityDescription) AttributesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("attributesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The relationships of the receiver in a dictionary.
//
// # Discussion
//
// The keys in the dictionary are the relationship names and the values are
// instances of [NSRelationshipDescription].
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/relationshipsByName
func (e NSEntityDescription) RelationshipsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("relationshipsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// An array of fetch index descriptions for the entity.
//
// # Discussion
//
// This value doesn’t form part of the entity’s version hash, and stores
// that don’t natively support indexing may ignore it.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/indexes
func (e NSEntityDescription) Indexes() []NSFetchIndexDescription {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("indexes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFetchIndexDescription {
		return NSFetchIndexDescriptionFromID(id)
	})
}
func (e NSEntityDescription) SetIndexes(value []NSFetchIndexDescription) {
	objc.Send[struct{}](e.ID, objc.Sel("setIndexes:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of arrays that contains one or more attributes with a value that
// must be unique over the instances of that entity.
//
// # Discussion
//
// Each inner array contains one or more [NSAttributeDescription] objects or
// strings that contain the names of attributes on the entity.
//
// This value forms part of the entity’s version hash. Stores that don’t
// support uniqueness constraints must refuse to initialize when receiving a
// model that contains such constraints.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/uniquenessConstraints
func (e NSEntityDescription) UniquenessConstraints() []foundation.NSArray {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("uniquenessConstraints"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSArray {
		return foundation.NSArrayFromID(id)
	})
}
func (e NSEntityDescription) SetUniquenessConstraints(value []foundation.NSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setUniquenessConstraints:"), objectivec.IObjectSliceToNSArray(value))
}

// The version hash for the receiver.
//
// # Discussion
//
// The version hash is used to uniquely identify an entity based on the
// collection and configuration of properties for the entity. The version hash
// uses only values which affect the persistence of data and the user-defined
// [NSEntityDescription.VersionHashModifier] value. (The values which affect
// persistence are: the name of the entity, the version hash of the
// superentity (if present), if the entity is abstract, and all of the version
// hashes for the properties.) This value is stored as part of the version
// information in the metadata for stores which use this entity, as well as a
// definition of an entity involved in an [NSEntityMapping] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/versionHash
func (e NSEntityDescription) VersionHash() foundation.NSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("versionHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The version hash modifier for the receiver.
//
// # Discussion
//
// This value is included in the version hash for the entity. You use it to
// mark or denote an entity as being a different “version” than another
// even if all of the values which affect persistence are equal. (Such a
// difference is important in cases where, for example, the structure of an
// entity is unchanged but the format or content of data has changed.)
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityDescription/versionHashModifier
func (e NSEntityDescription) VersionHashModifier() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("versionHashModifier"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityDescription) SetVersionHashModifier(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setVersionHashModifier:"), objc.String(value))
}
