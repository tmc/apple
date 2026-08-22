// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSEntityMapping] class.
var (
	_NSEntityMappingClass     NSEntityMappingClass
	_NSEntityMappingClassOnce sync.Once
)

func getNSEntityMappingClass() NSEntityMappingClass {
	_NSEntityMappingClassOnce.Do(func() {
		_NSEntityMappingClass = NSEntityMappingClass{class: objc.GetClass("NSEntityMapping")}
	})
	return _NSEntityMappingClass
}

// GetNSEntityMappingClass returns the class object for NSEntityMapping.
func GetNSEntityMappingClass() NSEntityMappingClass {
	return getNSEntityMappingClass()
}

type NSEntityMappingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSEntityMappingClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSEntityMappingClass) Alloc() NSEntityMapping {
	rv := objc.Send[NSEntityMapping](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mapping instance that specifies how to map an entity from a source to a
// destination managed object model.
//
// # Managing Source Information
//
//   - [NSEntityMapping.SourceEntityName]: The source entity name for the entity mapping.
//   - [NSEntityMapping.SetSourceEntityName]
//   - [NSEntityMapping.SourceEntityVersionHash]: The version hash of the source entity for the entity mapping.
//   - [NSEntityMapping.SetSourceEntityVersionHash]
//   - [NSEntityMapping.SourceExpression]: The source expression for the entity mapping.
//   - [NSEntityMapping.SetSourceExpression]
//
// # Managing Destination Information
//
//   - [NSEntityMapping.DestinationEntityName]: The destination entity name for the entity mapping.
//   - [NSEntityMapping.SetDestinationEntityName]
//   - [NSEntityMapping.DestinationEntityVersionHash]: The version hash for the destination entity for the entity mapping.
//   - [NSEntityMapping.SetDestinationEntityVersionHash]
//
// # Managing Mapping Information
//
//   - [NSEntityMapping.Name]: The name of the entity mapping.
//   - [NSEntityMapping.SetName]
//   - [NSEntityMapping.MappingType]: The mapping type for the entity mapping.
//   - [NSEntityMapping.SetMappingType]
//   - [NSEntityMapping.EntityMigrationPolicyClassName]: The class name of the migration policy for the entity mapping.
//   - [NSEntityMapping.SetEntityMigrationPolicyClassName]
//   - [NSEntityMapping.AttributeMappings]: The array of attribute mappings for the entity mapping.
//   - [NSEntityMapping.SetAttributeMappings]
//   - [NSEntityMapping.RelationshipMappings]: The array of relationship mappings for the entity mapping.
//   - [NSEntityMapping.SetRelationshipMappings]
//   - [NSEntityMapping.UserInfo]: The user info dictionary for the entity mapping.
//   - [NSEntityMapping.SetUserInfo]
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping
type NSEntityMapping struct {
	objectivec.Object
}

// NSEntityMappingFromID constructs a [NSEntityMapping] from an objc.ID.
//
// A mapping instance that specifies how to map an entity from a source to a
// destination managed object model.
func NSEntityMappingFromID(id objc.ID) NSEntityMapping {
	return NSEntityMapping{objectivec.Object{ID: id}}
}

// NOTE: NSEntityMapping adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSEntityMapping] class.
//
// # Managing Source Information
//
//   - [INSEntityMapping.SourceEntityName]: The source entity name for the entity mapping.
//   - [INSEntityMapping.SetSourceEntityName]
//   - [INSEntityMapping.SourceEntityVersionHash]: The version hash of the source entity for the entity mapping.
//   - [INSEntityMapping.SetSourceEntityVersionHash]
//   - [INSEntityMapping.SourceExpression]: The source expression for the entity mapping.
//   - [INSEntityMapping.SetSourceExpression]
//
// # Managing Destination Information
//
//   - [INSEntityMapping.DestinationEntityName]: The destination entity name for the entity mapping.
//   - [INSEntityMapping.SetDestinationEntityName]
//   - [INSEntityMapping.DestinationEntityVersionHash]: The version hash for the destination entity for the entity mapping.
//   - [INSEntityMapping.SetDestinationEntityVersionHash]
//
// # Managing Mapping Information
//
//   - [INSEntityMapping.Name]: The name of the entity mapping.
//   - [INSEntityMapping.SetName]
//   - [INSEntityMapping.MappingType]: The mapping type for the entity mapping.
//   - [INSEntityMapping.SetMappingType]
//   - [INSEntityMapping.EntityMigrationPolicyClassName]: The class name of the migration policy for the entity mapping.
//   - [INSEntityMapping.SetEntityMigrationPolicyClassName]
//   - [INSEntityMapping.AttributeMappings]: The array of attribute mappings for the entity mapping.
//   - [INSEntityMapping.SetAttributeMappings]
//   - [INSEntityMapping.RelationshipMappings]: The array of relationship mappings for the entity mapping.
//   - [INSEntityMapping.SetRelationshipMappings]
//   - [INSEntityMapping.UserInfo]: The user info dictionary for the entity mapping.
//   - [INSEntityMapping.SetUserInfo]
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping
type INSEntityMapping interface {
	objectivec.IObject

	// Topic: Managing Source Information

	// The source entity name for the entity mapping.
	SourceEntityName() string
	SetSourceEntityName(value string)
	// The version hash of the source entity for the entity mapping.
	SourceEntityVersionHash() foundation.NSData
	SetSourceEntityVersionHash(value foundation.NSData)
	// The source expression for the entity mapping.
	SourceExpression() foundation.NSExpression
	SetSourceExpression(value foundation.NSExpression)

	// Topic: Managing Destination Information

	// The destination entity name for the entity mapping.
	DestinationEntityName() string
	SetDestinationEntityName(value string)
	// The version hash for the destination entity for the entity mapping.
	DestinationEntityVersionHash() foundation.NSData
	SetDestinationEntityVersionHash(value foundation.NSData)

	// Topic: Managing Mapping Information

	// The name of the entity mapping.
	Name() string
	SetName(value string)
	// The mapping type for the entity mapping.
	MappingType() NSEntityMappingType
	SetMappingType(value NSEntityMappingType)
	// The class name of the migration policy for the entity mapping.
	EntityMigrationPolicyClassName() string
	SetEntityMigrationPolicyClassName(value string)
	// The array of attribute mappings for the entity mapping.
	AttributeMappings() []NSPropertyMapping
	SetAttributeMappings(value []NSPropertyMapping)
	// The array of relationship mappings for the entity mapping.
	RelationshipMappings() []NSPropertyMapping
	SetRelationshipMappings(value []NSPropertyMapping)
	// The user info dictionary for the entity mapping.
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)
}

// Init initializes the instance.
func (e NSEntityMapping) Init() NSEntityMapping {
	rv := objc.Send[NSEntityMapping](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSEntityMapping) Autorelease() NSEntityMapping {
	rv := objc.Send[NSEntityMapping](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSEntityMapping creates a new NSEntityMapping instance.
func NewNSEntityMapping() NSEntityMapping {
	class := getNSEntityMappingClass()
	rv := objc.Send[NSEntityMapping](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The source entity name for the entity mapping.
//
// # Discussion
//
// Mappings are not directly bound to entity descriptions; you can use the
// [NSMigrationManager.SourceEntityForEntityMapping] method on the migration
// manager to retrieve the entity description for this entity name.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityName
func (e NSEntityMapping) SourceEntityName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("sourceEntityName"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityMapping) SetSourceEntityName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setSourceEntityName:"), objc.String(value))
}

// The version hash of the source entity for the entity mapping.
//
// # Discussion
//
// The version hash is calculated by Core Data based on the property values of
// the entity (see [NSEntityDescription]’s [NSEntityDescription.VersionHash]
// method). The `sourceEntityVersionHash` must equal the version hash of the
// source entity represented by the mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceEntityVersionHash
func (e NSEntityMapping) SourceEntityVersionHash() foundation.NSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("sourceEntityVersionHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (e NSEntityMapping) SetSourceEntityVersionHash(value foundation.NSData) {
	objc.Send[struct{}](e.ID, objc.Sel("setSourceEntityVersionHash:"), value)
}

// The source expression for the entity mapping.
//
// # Discussion
//
// The source expression is used to obtain the collection of managed objects
// to process through the mapping. The expression can be a fetch request
// expression, or any other expression that evaluates to a collection.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/sourceExpression
func (e NSEntityMapping) SourceExpression() foundation.NSExpression {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("sourceExpression"))
	return foundation.NSExpressionFromID(objc.ID(rv))
}
func (e NSEntityMapping) SetSourceExpression(value foundation.NSExpression) {
	objc.Send[struct{}](e.ID, objc.Sel("setSourceExpression:"), value)
}

// The destination entity name for the entity mapping.
//
// # Discussion
//
// Mappings are not directly bound to entity descriptions. You can use the
// migration manager’s
// [NSMigrationManager.DestinationEntityForEntityMapping] method to retrieve
// the entity description for this entity name.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityName
func (e NSEntityMapping) DestinationEntityName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("destinationEntityName"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityMapping) SetDestinationEntityName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setDestinationEntityName:"), objc.String(value))
}

// The version hash for the destination entity for the entity mapping.
//
// # Discussion
//
// The version hash is calculated by Core Data based on the property values of
// the entity (see [NSEntityDescription]’s [NSEntityDescription.VersionHash]
// method). The `destinationEntityVersionHash` must equal the version hash of
// the destination entity represented by the mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/destinationEntityVersionHash
func (e NSEntityMapping) DestinationEntityVersionHash() foundation.NSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("destinationEntityVersionHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (e NSEntityMapping) SetDestinationEntityVersionHash(value foundation.NSData) {
	objc.Send[struct{}](e.ID, objc.Sel("setDestinationEntityVersionHash:"), value)
}

// The name of the entity mapping.
//
// # Discussion
//
// The name is used only as a means of distinguishing mappings in a model. If
// not specified, the value defaults to SOURCE->DESTINATION.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/name
func (e NSEntityMapping) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityMapping) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}

// The mapping type for the entity mapping.
//
// # Discussion
//
// If you specify a custom entity mapping type, you must specify a value for
// the migration policy class name as well (see
// [NSEntityMapping.EntityMigrationPolicyClassName]).
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/mappingType
func (e NSEntityMapping) MappingType() NSEntityMappingType {
	rv := objc.Send[NSEntityMappingType](e.ID, objc.Sel("mappingType"))
	return NSEntityMappingType(rv)
}
func (e NSEntityMapping) SetMappingType(value NSEntityMappingType) {
	objc.Send[struct{}](e.ID, objc.Sel("setMappingType:"), value)
}

// The class name of the migration policy for the entity mapping.
//
// # Discussion
//
// If not specified, the default migration class name is
// [NSEntityMigrationPolicy]. You can specify a subclass to provide custom
// behavior.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/entityMigrationPolicyClassName
func (e NSEntityMapping) EntityMigrationPolicyClassName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("entityMigrationPolicyClassName"))
	return foundation.NSStringFromID(rv).String()
}
func (e NSEntityMapping) SetEntityMigrationPolicyClassName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setEntityMigrationPolicyClassName:"), objc.String(value))
}

// The array of attribute mappings for the entity mapping.
//
// # Discussion
//
// The order of mappings in the array specifies the order in which the
// mappings will be processed during a migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/attributeMappings
func (e NSEntityMapping) AttributeMappings() []NSPropertyMapping {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("attributeMappings"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPropertyMapping {
		return NSPropertyMappingFromID(id)
	})
}
func (e NSEntityMapping) SetAttributeMappings(value []NSPropertyMapping) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttributeMappings:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of relationship mappings for the entity mapping.
//
// # Discussion
//
// The order of mappings in the array specifies the order in which the
// mappings will be processed during a migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/relationshipMappings
func (e NSEntityMapping) RelationshipMappings() []NSPropertyMapping {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("relationshipMappings"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPropertyMapping {
		return NSPropertyMappingFromID(id)
	})
}
func (e NSEntityMapping) SetRelationshipMappings(value []NSPropertyMapping) {
	objc.Send[struct{}](e.ID, objc.Sel("setRelationshipMappings:"), objectivec.IObjectSliceToNSArray(value))
}

// The user info dictionary for the entity mapping.
//
// # Discussion
//
// You can use the info dictionary in any way that might be useful in your
// migration. You can set the contents of the dictionary directory or using
// the appropriate inspector in the Xcode mapping model editor.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMapping/userInfo
func (e NSEntityMapping) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e NSEntityMapping) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setUserInfo:"), value)
}
