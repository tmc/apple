// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSClassDescription] class.
var (
	_NSClassDescriptionClass     NSClassDescriptionClass
	_NSClassDescriptionClassOnce sync.Once
)

func getNSClassDescriptionClass() NSClassDescriptionClass {
	_NSClassDescriptionClassOnce.Do(func() {
		_NSClassDescriptionClass = NSClassDescriptionClass{class: objc.GetClass("NSClassDescription")}
	})
	return _NSClassDescriptionClass
}

// GetNSClassDescriptionClass returns the class object for NSClassDescription.
func GetNSClassDescriptionClass() NSClassDescriptionClass {
	return getNSClassDescriptionClass()
}

type NSClassDescriptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSClassDescriptionClass) Alloc() NSClassDescription {
	rv := objc.Send[NSClassDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that provides the interface for querying the
// relationships and properties of a class.
//
// # Overview
// 
// Concrete subclasses of [NSClassDescription] provide the available
// attributes of objects of a particular class and the relationships between
// that class and other classes. Defining these relationships between classes
// allows for more intelligent and flexible manipulation of objects with
// key-value coding.
// 
// It is important to note that there are no class descriptions by default. To
// use [NSClassDescription] objects in your code you have to implement them
// for your model classes. For all concrete subclasses, you must provide
// implementations for all instance methods of [NSClassDescription].
// ([NSClassDescription] provides only the implementation for the class
// methods that maintain the cache of registered class descriptions.) Once
// created, you must register a class description with the
// [NSClassDescription] method [NSClassDescription.RegisterClassDescriptionForClass].
// 
// You can use the [NSString] objects in the arrays returned by methods such
// as [NSClassDescription.AttributeKeys] and [NSClassDescription.ToManyRelationshipKeys] to access—using key-value
// coding—the properties of an instance of the class to which a class
// description object corresponds. For more about attributes and
// relationships, see Cocoa Fundamentals Guide. For more about key-value
// coding, see [Key-Value Coding Programming Guide].
// 
// [NSScriptClassDescription], which is used to map the relationships between
// scriptable classes, is the only concrete subclass of [NSClassDescription]
// provided as part of the Cocoa framework.
//
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// # Attribute keys
//
//   - [NSClassDescription.AttributeKeys]: Overridden by subclasses to return the names of attributes of instances of the described class.
//
// # Relationship keys
//
//   - [NSClassDescription.InverseForRelationshipKey]: Overridden by subclasses to return the name of the inverse relationship from a relationship specified by a given key.
//   - [NSClassDescription.ToManyRelationshipKeys]: Overridden by subclasses to return the keys for the to-many relationship properties of instances of the described class.
//   - [NSClassDescription.ToOneRelationshipKeys]: Overridden by subclasses to return the keys for the to-one relationship properties of instances of the described class.
//
// # Notifications
//
//   - [NSClassDescription.NSClassDescriptionNeededForClass]: Posted by 
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription
type NSClassDescription struct {
	objectivec.Object
}

// NSClassDescriptionFromID constructs a [NSClassDescription] from an objc.ID.
//
// An abstract class that provides the interface for querying the
// relationships and properties of a class.
func NSClassDescriptionFromID(id objc.ID) NSClassDescription {
	return NSClassDescription{objectivec.Object{ID: id}}
}
// NOTE: NSClassDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSClassDescription] class.
//
// # Attribute keys
//
//   - [INSClassDescription.AttributeKeys]: Overridden by subclasses to return the names of attributes of instances of the described class.
//
// # Relationship keys
//
//   - [INSClassDescription.InverseForRelationshipKey]: Overridden by subclasses to return the name of the inverse relationship from a relationship specified by a given key.
//   - [INSClassDescription.ToManyRelationshipKeys]: Overridden by subclasses to return the keys for the to-many relationship properties of instances of the described class.
//   - [INSClassDescription.ToOneRelationshipKeys]: Overridden by subclasses to return the keys for the to-one relationship properties of instances of the described class.
//
// # Notifications
//
//   - [INSClassDescription.NSClassDescriptionNeededForClass]: Posted by 
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription
type INSClassDescription interface {
	objectivec.IObject

	// Topic: Attribute keys

	// Overridden by subclasses to return the names of attributes of instances of the described class.
	AttributeKeys() []string

	// Topic: Relationship keys

	// Overridden by subclasses to return the name of the inverse relationship from a relationship specified by a given key.
	InverseForRelationshipKey(relationshipKey objectivec.IObject) objectivec.IObject
	// Overridden by subclasses to return the keys for the to-many relationship properties of instances of the described class.
	ToManyRelationshipKeys() []string
	// Overridden by subclasses to return the keys for the to-one relationship properties of instances of the described class.
	ToOneRelationshipKeys() []string

	// Topic: Notifications

	// Posted by 
	NSClassDescriptionNeededForClass() NSNotificationName
}

// Init initializes the instance.
func (c NSClassDescription) Init() NSClassDescription {
	rv := objc.Send[NSClassDescription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSClassDescription) Autorelease() NSClassDescription {
	rv := objc.Send[NSClassDescription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSClassDescription creates a new NSClassDescription instance.
func NewNSClassDescription() NSClassDescription {
	class := getNSClassDescriptionClass()
	rv := objc.Send[NSClassDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the class description for a given class.
//
// aClass: The class for which to return a class description. See note below for
// important details.
//
// # Return Value
// 
// The class description for `aClass`, or `nil` if a class description cannot
// be found.
//
// # Discussion
// 
// If a class description for `aClass` is not found, the method posts an
// NSClassDescriptionNeededForClassNotification on behalf of `aClass`,
// allowing an observer to register a class description. The method then
// checks for a class description again. Returns `nil` if a class description
// is still not found.
// 
// If you have an instance of the receiver’s class, you can use the
// [NSObject] instance method [classDescription] instead.
//
// [classDescription]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classDescription
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/init(for:)
func NewClassDescriptionForClass(aClass objc.Class) NSClassDescription {
	rv := objc.Send[objc.ID](objc.ID(getNSClassDescriptionClass().class), objc.Sel("classDescriptionForClass:"), aClass)
	return NSClassDescriptionFromID(rv)
}

// Overridden by subclasses to return the name of the inverse relationship
// from a relationship specified by a given key.
//
// # Return Value
// 
// The name of the inverse relationship from the relationship specified by
// `relationshipKey`.
//
// # Discussion
// 
// For a given key that defines the name of the relationship from the
// receiver’s class to another class, returns the name of the relationship
// from the other class to the receiver’s class. For example, suppose an
// Employee class has a relationship named `department` to a Department class,
// and that Department has a relationship named `employees` to Employee. The
// statement:
// 
// returns the string `employees`.
// 
// If you have an instance of the class the receiver describes, you can use
// the [NSObject] instance method [inverse(forRelationshipKey:)] instead.
//
// [inverse(forRelationshipKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inverse(forRelationshipKey:)
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/inverse(forRelationshipKey:)
func (c NSClassDescription) InverseForRelationshipKey(relationshipKey objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inverseForRelationshipKey:"), relationshipKey)
	return objectivec.Object{ID: rv}
}

// Removes all [NSClassDescription] objects from the cache.
//
// # Discussion
// 
// You should rarely need to invoke this method. Use it whenever a registered
// [NSClassDescription] object might be replaced by a different version, such
// as when you have loaded a new provider of [NSClassDescription] objects, or
// when you are about to remove a provider of [NSClassDescription] objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/invalidateClassDescriptionCache()
func (_NSClassDescriptionClass NSClassDescriptionClass) InvalidateClassDescriptionCache() {
	objc.Send[objc.ID](objc.ID(_NSClassDescriptionClass.class), objc.Sel("invalidateClassDescriptionCache"))
}

// Registers an [NSClassDescription] object for a given class in the
// [NSClassDescription] cache.
//
// description: The class description to register.
//
// aClass: The class for which to register `description`.
//
// # Discussion
// 
// You should rarely need to directly invoke this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/register(_:for:)
func (_NSClassDescriptionClass NSClassDescriptionClass) RegisterClassDescriptionForClass(description INSClassDescription, aClass objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSClassDescriptionClass.class), objc.Sel("registerClassDescription:forClass:"), description, aClass)
}

// Overridden by subclasses to return the names of attributes of instances of
// the described class.
//
// # Return Value
// 
// An array of [NSString] objects containing the names of attributes of
// instances of the described class.
// 
// # Discussion
// 
// For example, a class description that describes Movie objects could return
// the attribute keys `title`, `dateReleased`, and `rating`.
// 
// If you have an instance of the class the receiver describes, you can use
// the [NSObject] instance method [attributeKeys] instead.
//
// [attributeKeys]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attributeKeys
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/attributeKeys
func (c NSClassDescription) AttributeKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("attributeKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Overridden by subclasses to return the keys for the to-many relationship
// properties of instances of the described class.
//
// # Return Value
// 
// An array of [NSString] objects containing the names of the to-many
// relationship properties of instances of the described class.
// 
// # Discussion
// 
// To-many relationship properties are arrays of objects.
// 
// If you have an instance of the class the receiver describes, you can use
// the [NSObject] instance method [toManyRelationshipKeys] instead.
//
// [toManyRelationshipKeys]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/toManyRelationshipKeys
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/toManyRelationshipKeys
func (c NSClassDescription) ToManyRelationshipKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("toManyRelationshipKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Overridden by subclasses to return the keys for the to-one relationship
// properties of instances of the described class.
//
// # Return Value
// 
// An array of [NSString] objects containing the names of the to-one
// relationship properties of instances of the described class.
// 
// # Discussion
// 
// To-one relationship properties are single objects.
// 
// If you have an instance of the class the receiver describes, you can use
// the [NSObject] instance method [toOneRelationshipKeys] instead.
//
// [toOneRelationshipKeys]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/toOneRelationshipKeys
//
// See: https://developer.apple.com/documentation/Foundation/NSClassDescription/toOneRelationshipKeys
func (c NSClassDescription) ToOneRelationshipKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("toOneRelationshipKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Posted by
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsclassdescriptionneededforclass
func (c NSClassDescription) NSClassDescriptionNeededForClass() NSNotificationName {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("NSClassDescriptionNeededForClassNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

