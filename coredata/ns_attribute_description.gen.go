// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAttributeDescription] class.
var (
	_NSAttributeDescriptionClass     NSAttributeDescriptionClass
	_NSAttributeDescriptionClassOnce sync.Once
)

func getNSAttributeDescriptionClass() NSAttributeDescriptionClass {
	_NSAttributeDescriptionClassOnce.Do(func() {
		_NSAttributeDescriptionClass = NSAttributeDescriptionClass{class: objc.GetClass("NSAttributeDescription")}
	})
	return _NSAttributeDescriptionClass
}

// GetNSAttributeDescriptionClass returns the class object for NSAttributeDescription.
func GetNSAttributeDescriptionClass() NSAttributeDescriptionClass {
	return getNSAttributeDescriptionClass()
}

type NSAttributeDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAttributeDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAttributeDescriptionClass) Alloc() NSAttributeDescription {
	rv := objc.Send[NSAttributeDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of a single attribute belonging to an entity.
//
// # Overview
//
// [NSAttributeDescription] inherits from [NSPropertyDescription], which
// provides most of the basic behavior. Instances of [NSAttributeDescription]
// are used to describe attributes, as distinct from relationships. The class
// adds the ability to specify the attribute type, and to specify a default
// value. In a managed object model, you must specify the type of all
// attributes—you can only use the undefined attribute type
// ([NSUndefinedAttributeType]) for transient attributes.
//
// # Editing Attribute Descriptions
//
// Attribute descriptions are editable until they are used by an object graph
// manager. This allows you to create or modify them dynamically. However,
// once a description is used (when the managed object model to which it
// belongs is associated with a persistent store coordinator), it must not
// (indeed cannot) be changed. This is enforced at runtime: any attempt to
// mutate a model or any of its sub-objects after the model is associated with
// a persistent store coordinator causes an exception to be thrown. If you
// need to modify a model that is in use, create a copy, modify the copy, and
// then discard the objects with the old model.
//
// # Managing the type
//
//   - [NSAttributeDescription.AttributeValueClassName]: The class name that represents the attribute’s value.
//   - [NSAttributeDescription.SetAttributeValueClassName]
//   - [NSAttributeDescription.AttributeType]: The attribute’s type.
//   - [NSAttributeDescription.SetAttributeType]
//
// # Configuring the behavior
//
//   - [NSAttributeDescription.AllowsCloudEncryption]: A Boolean value that determines whether to encrypt the attribute’s value.
//   - [NSAttributeDescription.SetAllowsCloudEncryption]
//   - [NSAttributeDescription.AllowsExternalBinaryDataStorage]: A Boolean value that indicates whether the attribute allows external binary storage.
//   - [NSAttributeDescription.SetAllowsExternalBinaryDataStorage]
//   - [NSAttributeDescription.DefaultValue]: The default value of the attribute.
//   - [NSAttributeDescription.SetDefaultValue]
//   - [NSAttributeDescription.PreservesValueInHistoryOnDeletion]: A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
//   - [NSAttributeDescription.SetPreservesValueInHistoryOnDeletion]
//   - [NSAttributeDescription.ValueTransformerName]: The name of the transformer to use for the attribute value.
//   - [NSAttributeDescription.SetValueTransformerName]
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription
type NSAttributeDescription struct {
	NSPropertyDescription
}

// NSAttributeDescriptionFromID constructs a [NSAttributeDescription] from an objc.ID.
//
// A description of a single attribute belonging to an entity.
func NSAttributeDescriptionFromID(id objc.ID) NSAttributeDescription {
	return NSAttributeDescription{NSPropertyDescription: NSPropertyDescriptionFromID(id)}
}

// NOTE: NSAttributeDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAttributeDescription] class.
//
// # Managing the type
//
//   - [INSAttributeDescription.AttributeValueClassName]: The class name that represents the attribute’s value.
//   - [INSAttributeDescription.SetAttributeValueClassName]
//   - [INSAttributeDescription.AttributeType]: The attribute’s type.
//   - [INSAttributeDescription.SetAttributeType]
//
// # Configuring the behavior
//
//   - [INSAttributeDescription.AllowsCloudEncryption]: A Boolean value that determines whether to encrypt the attribute’s value.
//   - [INSAttributeDescription.SetAllowsCloudEncryption]
//   - [INSAttributeDescription.AllowsExternalBinaryDataStorage]: A Boolean value that indicates whether the attribute allows external binary storage.
//   - [INSAttributeDescription.SetAllowsExternalBinaryDataStorage]
//   - [INSAttributeDescription.DefaultValue]: The default value of the attribute.
//   - [INSAttributeDescription.SetDefaultValue]
//   - [INSAttributeDescription.PreservesValueInHistoryOnDeletion]: A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
//   - [INSAttributeDescription.SetPreservesValueInHistoryOnDeletion]
//   - [INSAttributeDescription.ValueTransformerName]: The name of the transformer to use for the attribute value.
//   - [INSAttributeDescription.SetValueTransformerName]
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription
type INSAttributeDescription interface {
	INSPropertyDescription

	// Topic: Managing the type

	// The class name that represents the attribute’s value.
	AttributeValueClassName() string
	SetAttributeValueClassName(value string)
	// The attribute’s type.
	AttributeType() NSAttributeType
	SetAttributeType(value NSAttributeType)

	// Topic: Configuring the behavior

	// A Boolean value that determines whether to encrypt the attribute’s value.
	AllowsCloudEncryption() bool
	SetAllowsCloudEncryption(value bool)
	// A Boolean value that indicates whether the attribute allows external binary storage.
	AllowsExternalBinaryDataStorage() bool
	SetAllowsExternalBinaryDataStorage(value bool)
	// The default value of the attribute.
	DefaultValue() objectivec.IObject
	SetDefaultValue(value objectivec.IObject)
	// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion.
	PreservesValueInHistoryOnDeletion() bool
	SetPreservesValueInHistoryOnDeletion(value bool)
	// The name of the transformer to use for the attribute value.
	ValueTransformerName() string
	SetValueTransformerName(value string)
}

// Init initializes the instance.
func (a NSAttributeDescription) Init() NSAttributeDescription {
	rv := objc.Send[NSAttributeDescription](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAttributeDescription) Autorelease() NSAttributeDescription {
	rv := objc.Send[NSAttributeDescription](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAttributeDescription creates a new NSAttributeDescription instance.
func NewNSAttributeDescription() NSAttributeDescription {
	class := getNSAttributeDescriptionClass()
	rv := objc.Send[NSAttributeDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The class name that represents the attribute’s value.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeValueClassName
func (a NSAttributeDescription) AttributeValueClassName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributeValueClassName"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAttributeDescription) SetAttributeValueClassName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributeValueClassName:"), objc.String(value))
}

// The attribute’s type.
//
// # Discussion
//
// Don’t change an attribute’s type after you add its containing managed
// object model to a persistent store coordinator; otherwise, Core Data throws
// an exception.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/attributeType-swift.property
func (a NSAttributeDescription) AttributeType() NSAttributeType {
	rv := objc.Send[NSAttributeType](a.ID, objc.Sel("attributeType"))
	return NSAttributeType(rv)
}
func (a NSAttributeDescription) SetAttributeType(value NSAttributeType) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributeType:"), value)
}

// A Boolean value that determines whether to encrypt the attribute’s value.
//
// # Discussion
//
// Set this property to true to store the attribute’s value in an encrypted
// form in iCloud. Only use this property with new attributes. Core Data
// doesn’t support encrypting attributes that already exist in your CloudKit
// schema, or attributes that represent relationships between entities.
//
// You can also set this property using the Allow Cloud Encryption attribute
// in the Attributes inspector of the Core Data model editor.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsCloudEncryption
func (a NSAttributeDescription) AllowsCloudEncryption() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsCloudEncryption"))
	return rv
}
func (a NSAttributeDescription) SetAllowsCloudEncryption(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowsCloudEncryption:"), value)
}

// A Boolean value that indicates whether the attribute allows external binary
// storage.
//
// # Discussion
//
// true if the attribute allows external binary storage, otherwise false. If
// this value is true, the corresponding attribute may be stored in a file
// external to the persistent store itself.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/allowsExternalBinaryDataStorage
func (a NSAttributeDescription) AllowsExternalBinaryDataStorage() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsExternalBinaryDataStorage"))
	return rv
}
func (a NSAttributeDescription) SetAllowsExternalBinaryDataStorage(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowsExternalBinaryDataStorage:"), value)
}

// The default value of the attribute.
//
// # Discussion
//
// Default values are retained by a managed object model, not copied. This
// means that attribute values do not have to implement the [NSCopying]
// protocol, however it also means that you should not modify any objects
// after they have been set as default values.
//
// # Special Considerations
//
// Setting the default value raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/defaultValue
func (a NSAttributeDescription) DefaultValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("defaultValue"))
	return objectivec.Object{ID: rv}
}
func (a NSAttributeDescription) SetDefaultValue(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setDefaultValue:"), value)
}

// A Boolean value that indicates whether the attribute records its value in
// the persistent history transaction for a managed object’s deletion.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/preservesValueInHistoryOnDeletion
func (a NSAttributeDescription) PreservesValueInHistoryOnDeletion() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("preservesValueInHistoryOnDeletion"))
	return rv
}
func (a NSAttributeDescription) SetPreservesValueInHistoryOnDeletion(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreservesValueInHistoryOnDeletion:"), value)
}

// The name of the transformer to use for the attribute value.
//
// # Discussion
//
// The attribute must be of type [NSTransformedAttributeType].
//
// The transformer must output an [NSData] object from [transformedValue(_:)]
// and must allow reverse transformations.
//
// If this value is `nil`, Core Data uses a default a transformer that uses
// [NSCoding] to archive and unarchive the attribute value.
//
// See: https://developer.apple.com/documentation/CoreData/NSAttributeDescription/valueTransformerName
//
// [NSCoding]: https://developer.apple.com/documentation/Foundation/NSCoding
// [transformedValue(_:)]: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValue(_:)
func (a NSAttributeDescription) ValueTransformerName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("valueTransformerName"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAttributeDescription) SetValueTransformerName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setValueTransformerName:"), objc.String(value))
}
