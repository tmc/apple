// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSManagedObject] class.
var (
	_NSManagedObjectClass     NSManagedObjectClass
	_NSManagedObjectClassOnce sync.Once
)

func getNSManagedObjectClass() NSManagedObjectClass {
	_NSManagedObjectClassOnce.Do(func() {
		_NSManagedObjectClass = NSManagedObjectClass{class: objc.GetClass("NSManagedObject")}
	})
	return _NSManagedObjectClass
}

// GetNSManagedObjectClass returns the class object for NSManagedObject.
func GetNSManagedObjectClass() NSManagedObjectClass {
	return getNSManagedObjectClass()
}

type NSManagedObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSManagedObjectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSManagedObjectClass) Alloc() NSManagedObject {
	rv := objc.Send[NSManagedObject](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The base class that all Core Data model objects inherit from.
//
// # Overview
//
// A managed object has an associated entity description
// ([NSEntityDescription]) that provides metadata about the object, including
// the name of the entity that the object represents and the names of its
// attributes and relationships. A managed object also has an associated
// managed object context that tracks changes to the object graph.
//
// You can’t use instances of direct subclasses of [NSObject], or any other
// class that doesn’t inherit from [NSManagedObject], with a managed object
// context. You may create custom subclasses of [NSManagedObject], although
// this isn’t always necessary. If you don’t need custom logic, you can
// create a complete object graph with [NSManagedObject] instances.
//
// If you instantiate a managed object directly, you must call the designated
// initializer [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext].
//
// # Data Storage
//
// In some respects, an [NSManagedObject] acts like a dictionary—it’s a
// generic container object that provides efficient storage for the properties
// defined by its associated [NSEntityDescription] instance. [NSManagedObject]
// supports a range of common types for attribute values, including string,
// date, and number (see [NSAttributeDescription] for full details).
// Therefore, typically you don’t need to define instance variables in
// subclasses. Sometimes, however, you want to use types that aren’t
// supported directly, such as colors and C structures. For example, in a
// graphics application you might want to define a Rectangle entity that has
// color and bounds attributes that are an instance of [NSColor] and an
// [NSRect] struct, respectively. For some types you can use a transformable
// attribute, for others this may require you to create a subclass of
// [NSManagedObject].
//
// # Faulting
//
// Managed objects typically represent data held in a persistent store. In
// some situations a managed object may be a fault—an object whose property
// values haven’t yet been loaded from the external data store. When you
// access persistent property values, the fault “fires” and the data is
// retrieved from the store automatically. This can be a comparatively
// expensive process (potentially requiring a round trip to the persistent
// store), and you may wish to avoid unnecessarily firing a fault. See
// [Faulting and Uniquing] for more details on faults.
//
// You can safely invoke the following methods and properties on a fault
// without causing it to fire: [isEqual(_:)], [hash], [superclass], [class],
// [self()], [isProxy()], [isKind(of:)], [isMember(of:)], [conforms(to:)],
// [responds(to:)], [description], [NSManagedObject.ManagedObjectContext],
// [NSManagedObject.Entity], [NSManagedObject.ObjectID],
// [NSManagedObject.Inserted], [NSManagedObject.Updated],
// [NSManagedObject.Deleted], [NSManagedObject.FaultingState], and
// [NSManagedObject.Fault]. Because `isEqual` and `hash` don’t cause a fault
// to fire, managed objects can typically be placed in collections without
// firing a fault. Note, however, that invoking key-value coding methods on
// the collection object might in turn result in an invocation of “ on a
// managed object, which would fire the fault.
//
// Although the `description` property doesn’t cause a fault to fire, if you
// implement a custom `description` that accesses the object’s persistent
// properties, this does cause a fault to fire. You are strongly discouraged
// from overriding `description` in this way.
//
// # Subclassing Notes
//
// In combination with the entity description in the managed object model,
// [NSManagedObject] provides a rich set of default behaviors including
// support for arbitrary properties and value validation. If you decide to
// subclass [NSManagedObject] to implement custom features, make sure you
// don’t disrupt Core Data’s behavior.
//
// # Methods and Properties You Must Not Override
//
// [NSManagedObject] itself customizes many features of [NSObject] so that
// managed objects can be properly integrated into the Core Data
// infrastructure. Core Data relies on the [NSManagedObject] implementation of
// the following methods and properties, which you therefore absolutely must
// not override: [NSManagedObject.PrimitiveValueForKey],
// [NSManagedObject.SetPrimitiveValueForKey], [isEqual(_:)], [hash],
// [superclass], [class], [self()], [isProxy()], [isKind(of:)],
// [isMember(of:)], [conforms(to:)], [responds(to:)],
// [NSManagedObject.ManagedObjectContext], [NSManagedObject.Entity],
// [NSManagedObject.ObjectID], [NSManagedObject.Inserted],
// [NSManagedObject.Updated], [NSManagedObject.Deleted], and
// [NSManagedObject.Fault], [alloc], [allocWithZone:], [new],
// [instancesRespond(to:)], [instanceMethod(for:)], [method(for:)],
// [methodSignatureForSelector:], [instanceMethodSignatureForSelector:], or
// [isSubclass(of:)].
//
// # Methods and Properties You Shouldn’t Override
//
// As with any class, you are strongly discouraged from overriding the
// key-value observing methods such as [willChangeValue(forKey:)] and
// [didChangeValue(forKey:withSetMutation:using:)]. Avoid overriding
// `description`—if this method fires a fault during a debugging operation,
// the results may be unpredictable. Also avoid overriding
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext], or
// `dealloc`. Changing values in the
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext] method
// won’t be noticed by the context, and if you aren’t careful, those
// changes may not be saved. Perform most initialization customization in one
// of the `awake…` methods. If you do override
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext], make sure
// you adhere to the requirements set out in the method description. See
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext].
//
// Don’t override `dealloc` because [NSManagedObject.DidTurnIntoFault] is
// usually a better time to clear values—a managed object may not be
// reclaimed for some time after it has been turned into a fault. Core Data
// doesn’t guarantee that `dealloc` will be called in all scenarios (such as
// when the application quits). Therefore, don’t include required side
// effects (like saving or changes to the file system, user preferences, and
// so on) in these methods.
//
// In summary, for
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext] and
// `dealloc`, Core Data reserves exclusive control over the life cycle of the
// managed object (that is, raw memory management). This is so that the
// framework can provide features such as uniquing and by consequence,
// relationship maintenance, as well as much better performance than would be
// possible otherwise.
//
// # Additional Override Considerations
//
// The following methods are intended to be fine grained and aren’t suitable
// for large-scale operations. Don’t fetch or save in these methods. In
// particular, they shouldn’t have side effects on the managed object
// context.
//
// - [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext] -
// [NSManagedObject.DidTurnIntoFault] - [NSManagedObject.WillTurnIntoFault] -
// `dealloc`
//
// In addition, if you plan to override `awakeFromInsert`, `awakeFromFetch`,
// and validation methods, first invoke `super.Method()`, the superclass’s
// implementation. Don’t modify relationships in
// [NSManagedObject.AwakeFromFetch]—see the method description for details.
//
// # Custom Accessor Methods
//
// Typically, you don’t need to write custom accessor methods for properties
// that are defined in the entity of a managed object’s corresponding
// managed object model. If you need to do so, follow the implementation
// patterns described in Managed Object Accessor Methods in [Core Data
// Programming Guide].
//
// Core Data automatically generates accessor methods (and primitive accessor
// methods) for you. For attributes and to-one relationships, Core Data
// generates the standard get and set accessor methods; for to-many
// relationships, Core Data generates the indexed accessor methods as
// described in [Achieving Basic Key-Value Coding Compliance] in [Key-Value
// Coding Programming Guide]. You do however need to declare the accessor
// methods or use Objective-C properties to suppress compiler warnings. For a
// full discussion, see Managed Object Accessor Methods in [Core Data
// Programming Guide].
//
// # Custom Instance Variables
//
// By default, [NSManagedObject] stores its properties in an internal
// structure as objects, and in general Core Data is more efficient working
// with storage under its own control rather than by using custom instance
// variables.
//
// [NSManagedObject] provides support for a range of common types for
// attribute values, including string, date, and number (see
// [NSAttributeDescription] for full details). If you want to use types that
// aren’t supported directly, like colors and C structures, you can either
// use transformable attributes or create a subclass of [NSManagedObject].
//
// Sometimes it’s convenient to represent variables as scalars—in drawing
// applications, for example, where variables represent dimensions and x and y
// coordinates and are frequently used in calculations. To represent
// attributes as scalars, you declare instance variables as you do in any
// other class. You also need to implement suitable accessor methods as
// described in Managed Object Accessor Methods.
//
// If you define custom instance variables for example to store derived
// attributes or other transient properties, clean up these variables in
// [NSManagedObject.DidTurnIntoFault] rather than `dealloc`.
//
// # Validation Methods
//
// [NSManagedObject] provides consistent hooks for validating property and
// inter-property values. You typically shouldn’t override
// [NSManagedObject.ValidateValueForKeyError]. Instead implement methods of
// the form `validate:`, as defined by the NSKeyValueCoding protocol. If you
// want to validate inter-property values, you can override
// [NSManagedObject.ValidateForUpdate] and/or related validation methods.
//
// Don’t call “ within custom property validation methods—if you do, you
// create an infinite loop when “ is invoked at runtime. If you do implement
// custom validation methods, don’t call them directly. Instead, call “
// with the appropriate key. This ensures that any constraints defined in the
// managed object model are applied.
//
// If you implement custom inter-property validation methods like
// [NSManagedObject.ValidateForUpdate], call the superclass’s implementation
// first. This ensures that individual property validation methods are also
// invoked. If there are multiple validation failures in one operation,
// collect them in an array and add the array—using the key
// [NSDetailedErrorsKey]—to the userInfo dictionary in the [NSError] object
// you return. For an example, see Managed Object Validation.
//
// # Creating a Managed Object
//
//   - [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext]: Initializes a managed object from an entity description and inserts it into the specified managed object context.
//   - [NSManagedObject.InitWithContext]: Initializes a managed object subclass and inserts it into the specified managed object context.
//
// # Getting a Managed Object’s Identity
//
//   - [NSManagedObject.Entity]: The entity description of the managed object.
//   - [NSManagedObject.ObjectID]: The object ID of the managed object.
//
// # Getting State Information
//
//   - [NSManagedObject.ManagedObjectContext]: The managed object context with which the managed object is registered.
//   - [NSManagedObject.HasChanges]: A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes.
//   - [NSManagedObject.IsInserted]: A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//   - [NSManagedObject.IsUpdated]: A Boolean value that indicates whether the managed object has unsaved changes.
//   - [NSManagedObject.IsDeleted]: A Boolean value that indicates whether the managed object will be deleted during the next save.
//   - [NSManagedObject.IsFault]: A Boolean value that indicates whether the managed object is a fault.
//   - [NSManagedObject.FaultingState]: The faulting state of the managed object.
//   - [NSManagedObject.HasFaultForRelationshipNamed]: Returns a Boolean value that indicates whether the relationship for a given key is a fault.
//   - [NSManagedObject.HasPersistentChangedValues]: A Boolean value that indicates whether the managed object has persistent changes.
//
// # Managing Change Events
//
//   - [NSManagedObject.AwakeFromFetch]: Provides an opportunity to add code into the life cycle of the managed object when fufilling it from a fault.
//   - [NSManagedObject.AwakeFromInsert]: Provides an opportunity to add code into the life cycle of the managed object when initially creating it.
//   - [NSManagedObject.AwakeFromSnapshotEvents]: Provides an opportunity to add code into the life cycle of the managed object when fulfilling it from a snapshot.
//   - [NSManagedObject.ChangedValues]: Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
//   - [NSManagedObject.ChangedValuesForCurrentEvent]: Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
//   - [NSManagedObject.CommittedValuesForKeys]: Returns a dictionary of the most recent fetched or saved values of the managed object for the properties of the specified keys.
//   - [NSManagedObject.PrepareForDeletion]: Provides an opportunity to add code into the life cycle of the managed object before deleting it.
//   - [NSManagedObject.WillSave]: Provides an opportunity to add code into the life cycle of the managed object before saving it.
//   - [NSManagedObject.DidSave]: Provides an opportunity to add code into the life cycle of the managed object after the managed object’s context completes a save operation.
//   - [NSManagedObject.WillTurnIntoFault]: Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault.
//   - [NSManagedObject.DidTurnIntoFault]: Provides an opportunity to add code into the life cycle of the managed object after converting it to a fault.
//
// # Supporting Key-Value Coding
//
//   - [NSManagedObject.PrimitiveValueForKey]: Returns the value for the specified property from the managed object’s private internal storage .
//   - [NSManagedObject.SetPrimitiveValueForKey]: Sets the value of a given property in the managed object’s private internal storage.
//   - [NSManagedObject.ObjectIDsForRelationshipNamed]: Returns the object IDs for all of the managed objects that are in the named relationship.
//
// # Managing Data Validation
//
//   - [NSManagedObject.ValidateForDelete]: Determines whether the managed object can be deleted in its current state.
//   - [NSManagedObject.ValidateForInsert]: Determines whether the managed object can be inserted in its current state.
//   - [NSManagedObject.ValidateForUpdate]: Determines whether the managed object’s current state is valid.
//
// # Supporting Key-Value Observing
//
//   - [NSManagedObject.DidAccessValueForKey]: Provides support for key-value observing access notification.
//   - [NSManagedObject.ObservationInfo]: Returns the observation info of the managed object.
//   - [NSManagedObject.SetObservationInfo]: Sets the observation info of the managed object.
//   - [NSManagedObject.WillAccessValueForKey]: Provides support for key-value observing access notification.
//   - [NSManagedObject.DidChangeValueForKeyWithSetMutationUsingObjects]: Provides an opportunity to respond when a change was made to a specified to-many relationship.
//   - [NSManagedObject.WillChangeValueForKeyWithSetMutationUsingObjects]: Provides an opportunity to respond when a change is about to be made to a specified to-many relationship.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject
//
// [Achieving Basic Key-Value Coding Compliance]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/AccessorConventions.html#//apple_ref/doc/uid/20002174
// [Core Data Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreData/index.html#//apple_ref/doc/uid/TP40001075
// [Faulting and Uniquing]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreData/FaultingandUniquing.html#//apple_ref/doc/uid/TP40001075-CH18
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
// [allocWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/allocWithZone:
// [alloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/alloc
// [class]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/class
// [conforms(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/conforms(to:)
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
// [didChangeValue(forKey:withSetMutation:using:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChangeValue(forKey:withSetMutation:using:)
// [hash]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
// [instanceMethod(for:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethod(for:)
// [instanceMethodSignatureForSelector:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethodSignatureForSelector:
// [instancesRespond(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instancesRespond(to:)
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [isKind(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
// [isMember(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isMember(of:)
// [isProxy()]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isProxy()
// [isSubclass(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSubclass(of:)
// [method(for:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/method(for:)
// [methodSignatureForSelector:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/methodSignatureForSelector:
// [new]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/new
// [responds(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/responds(to:)
// [self()]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/self()
// [superclass]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/superclass
// [willChangeValue(forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/willChangeValue(forKey:)
type NSManagedObject struct {
	objectivec.Object
}

// NSManagedObjectFromID constructs a [NSManagedObject] from an objc.ID.
//
// The base class that all Core Data model objects inherit from.
func NSManagedObjectFromID(id objc.ID) NSManagedObject {
	return NSManagedObject{objectivec.Object{ID: id}}
}

// NOTE: NSManagedObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSManagedObject] class.
//
// # Creating a Managed Object
//
//   - [INSManagedObject.InitWithEntityInsertIntoManagedObjectContext]: Initializes a managed object from an entity description and inserts it into the specified managed object context.
//   - [INSManagedObject.InitWithContext]: Initializes a managed object subclass and inserts it into the specified managed object context.
//
// # Getting a Managed Object’s Identity
//
//   - [INSManagedObject.Entity]: The entity description of the managed object.
//   - [INSManagedObject.ObjectID]: The object ID of the managed object.
//
// # Getting State Information
//
//   - [INSManagedObject.ManagedObjectContext]: The managed object context with which the managed object is registered.
//   - [INSManagedObject.HasChanges]: A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes.
//   - [INSManagedObject.IsInserted]: A Boolean value that indicates whether the managed object has been inserted in a managed object context.
//   - [INSManagedObject.IsUpdated]: A Boolean value that indicates whether the managed object has unsaved changes.
//   - [INSManagedObject.IsDeleted]: A Boolean value that indicates whether the managed object will be deleted during the next save.
//   - [INSManagedObject.IsFault]: A Boolean value that indicates whether the managed object is a fault.
//   - [INSManagedObject.FaultingState]: The faulting state of the managed object.
//   - [INSManagedObject.HasFaultForRelationshipNamed]: Returns a Boolean value that indicates whether the relationship for a given key is a fault.
//   - [INSManagedObject.HasPersistentChangedValues]: A Boolean value that indicates whether the managed object has persistent changes.
//
// # Managing Change Events
//
//   - [INSManagedObject.AwakeFromFetch]: Provides an opportunity to add code into the life cycle of the managed object when fufilling it from a fault.
//   - [INSManagedObject.AwakeFromInsert]: Provides an opportunity to add code into the life cycle of the managed object when initially creating it.
//   - [INSManagedObject.AwakeFromSnapshotEvents]: Provides an opportunity to add code into the life cycle of the managed object when fulfilling it from a snapshot.
//   - [INSManagedObject.ChangedValues]: Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
//   - [INSManagedObject.ChangedValuesForCurrentEvent]: Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
//   - [INSManagedObject.CommittedValuesForKeys]: Returns a dictionary of the most recent fetched or saved values of the managed object for the properties of the specified keys.
//   - [INSManagedObject.PrepareForDeletion]: Provides an opportunity to add code into the life cycle of the managed object before deleting it.
//   - [INSManagedObject.WillSave]: Provides an opportunity to add code into the life cycle of the managed object before saving it.
//   - [INSManagedObject.DidSave]: Provides an opportunity to add code into the life cycle of the managed object after the managed object’s context completes a save operation.
//   - [INSManagedObject.WillTurnIntoFault]: Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault.
//   - [INSManagedObject.DidTurnIntoFault]: Provides an opportunity to add code into the life cycle of the managed object after converting it to a fault.
//
// # Supporting Key-Value Coding
//
//   - [INSManagedObject.PrimitiveValueForKey]: Returns the value for the specified property from the managed object’s private internal storage .
//   - [INSManagedObject.SetPrimitiveValueForKey]: Sets the value of a given property in the managed object’s private internal storage.
//   - [INSManagedObject.ObjectIDsForRelationshipNamed]: Returns the object IDs for all of the managed objects that are in the named relationship.
//
// # Managing Data Validation
//
//   - [INSManagedObject.ValidateForDelete]: Determines whether the managed object can be deleted in its current state.
//   - [INSManagedObject.ValidateForInsert]: Determines whether the managed object can be inserted in its current state.
//   - [INSManagedObject.ValidateForUpdate]: Determines whether the managed object’s current state is valid.
//
// # Supporting Key-Value Observing
//
//   - [INSManagedObject.DidAccessValueForKey]: Provides support for key-value observing access notification.
//   - [INSManagedObject.ObservationInfo]: Returns the observation info of the managed object.
//   - [INSManagedObject.SetObservationInfo]: Sets the observation info of the managed object.
//   - [INSManagedObject.WillAccessValueForKey]: Provides support for key-value observing access notification.
//   - [INSManagedObject.DidChangeValueForKeyWithSetMutationUsingObjects]: Provides an opportunity to respond when a change was made to a specified to-many relationship.
//   - [INSManagedObject.WillChangeValueForKeyWithSetMutationUsingObjects]: Provides an opportunity to respond when a change is about to be made to a specified to-many relationship.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject
type INSManagedObject interface {
	objectivec.IObject
	NSFetchRequestResult

	// Topic: Creating a Managed Object

	// Initializes a managed object from an entity description and inserts it into the specified managed object context.
	InitWithEntityInsertIntoManagedObjectContext(entity INSEntityDescription, context INSManagedObjectContext) NSManagedObject
	// Initializes a managed object subclass and inserts it into the specified managed object context.
	InitWithContext(moc INSManagedObjectContext) NSManagedObject

	// Topic: Getting a Managed Object’s Identity

	// The entity description of the managed object.
	Entity() INSEntityDescription
	// The object ID of the managed object.
	ObjectID() INSManagedObjectID

	// Topic: Getting State Information

	// The managed object context with which the managed object is registered.
	ManagedObjectContext() INSManagedObjectContext
	// A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes.
	HasChanges() bool
	// A Boolean value that indicates whether the managed object has been inserted in a managed object context.
	IsInserted() bool
	// A Boolean value that indicates whether the managed object has unsaved changes.
	IsUpdated() bool
	// A Boolean value that indicates whether the managed object will be deleted during the next save.
	IsDeleted() bool
	// A Boolean value that indicates whether the managed object is a fault.
	IsFault() bool
	// The faulting state of the managed object.
	FaultingState() uint
	// Returns a Boolean value that indicates whether the relationship for a given key is a fault.
	HasFaultForRelationshipNamed(key string) bool
	// A Boolean value that indicates whether the managed object has persistent changes.
	HasPersistentChangedValues() bool

	// Topic: Managing Change Events

	// Provides an opportunity to add code into the life cycle of the managed object when fufilling it from a fault.
	AwakeFromFetch()
	// Provides an opportunity to add code into the life cycle of the managed object when initially creating it.
	AwakeFromInsert()
	// Provides an opportunity to add code into the life cycle of the managed object when fulfilling it from a snapshot.
	AwakeFromSnapshotEvents(flags NSSnapshotEventType)
	// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
	ChangedValues() foundation.INSDictionary
	// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object.
	ChangedValuesForCurrentEvent() foundation.INSDictionary
	// Returns a dictionary of the most recent fetched or saved values of the managed object for the properties of the specified keys.
	CommittedValuesForKeys(keys []string) foundation.INSDictionary
	// Provides an opportunity to add code into the life cycle of the managed object before deleting it.
	PrepareForDeletion()
	// Provides an opportunity to add code into the life cycle of the managed object before saving it.
	WillSave()
	// Provides an opportunity to add code into the life cycle of the managed object after the managed object’s context completes a save operation.
	DidSave()
	// Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault.
	WillTurnIntoFault()
	// Provides an opportunity to add code into the life cycle of the managed object after converting it to a fault.
	DidTurnIntoFault()

	// Topic: Supporting Key-Value Coding

	// Returns the value for the specified property from the managed object’s private internal storage .
	PrimitiveValueForKey(key string) objectivec.IObject
	// Sets the value of a given property in the managed object’s private internal storage.
	SetPrimitiveValueForKey(value objectivec.IObject, key string)
	// Returns the object IDs for all of the managed objects that are in the named relationship.
	ObjectIDsForRelationshipNamed(key string) []NSManagedObjectID

	// Topic: Managing Data Validation

	// Determines whether the managed object can be deleted in its current state.
	ValidateForDelete() (bool, error)
	// Determines whether the managed object can be inserted in its current state.
	ValidateForInsert() (bool, error)
	// Determines whether the managed object’s current state is valid.
	ValidateForUpdate() (bool, error)

	// Topic: Supporting Key-Value Observing

	// Provides support for key-value observing access notification.
	DidAccessValueForKey(key string)
	// Returns the observation info of the managed object.
	ObservationInfo() unsafe.Pointer
	// Sets the observation info of the managed object.
	SetObservationInfo(inObservationInfo unsafe.Pointer)
	// Provides support for key-value observing access notification.
	WillAccessValueForKey(key string)
	// Provides an opportunity to respond when a change was made to a specified to-many relationship.
	DidChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind uint, inObjects foundation.INSSet)
	// Provides an opportunity to respond when a change is about to be made to a specified to-many relationship.
	WillChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind uint, inObjects foundation.INSSet)
}

// Init initializes the instance.
func (m NSManagedObject) Init() NSManagedObject {
	rv := objc.Send[NSManagedObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSManagedObject) Autorelease() NSManagedObject {
	rv := objc.Send[NSManagedObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSManagedObject creates a new NSManagedObject instance.
func NewNSManagedObject() NSManagedObject {
	class := getNSManagedObjectClass()
	rv := objc.Send[NSManagedObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a managed object subclass and inserts it into the specified
// managed object context.
//
// # Return Value
//
// An initialized instance of the appropriate subclass.
//
// # Discussion
//
// This method is only legal to call on subclasses of [NSManagedObject] that
// represent a single entity in the model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(context:)
func NewManagedObjectWithContext(moc INSManagedObjectContext) NSManagedObject {
	instance := getNSManagedObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContext:"), moc)
	return NSManagedObjectFromID(rv)
}

// Initializes a managed object from an entity description and inserts it into
// the specified managed object context.
//
// entity: The entity of which to create an instance.
//
// The model associated with `context`’s persistent store coordinator must
// contain `entity`.
//
// context: The context into which the new instance is inserted.
//
// # Return Value
//
// An initialized instance of the appropriate class for `entity`.
//
// # Discussion
//
// [NSManagedObject] uses dynamic class generation to support the Objective-C
// 2 properties feature (see Declared Properties) by automatically creating a
// subclass of the class appropriate for `entity`. “ therefore returns an
// instance of the appropriate class for `entity`. The dynamically-generated
// subclass will be based on the class specified by the entity, so specifying
// a custom class in your model will supersede the class passed to `alloc`.
//
// If `context` is not `nil`, this method invokes `[context self]` (which
// causes [NSManagedObject.AwakeFromInsert] to be invoked).
//
// You are discouraged from overriding this method—you should instead
// override [NSManagedObject.AwakeFromInsert] and/or
// [NSManagedObject.AwakeFromFetch] (if there is logic common to these
// methods, it should be factored into a third method which is invoked from
// both). If you do perform custom initialization in this method, you may
// cause problems with undo and redo operations.
//
// In many applications, there is no need to subsequently assign a
// newly-created managed object to a particular store—see
// [NSManagedObjectContext.AssignObjectToPersistentStore]. If your application
// has multiple stores and you do need to assign an object to a specific
// store, you should not do so in a managed object’s initializer method.
// Such an assignment is controller- not model-level logic.
//
// # Special Considerations
//
// If you override
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext], you must
// ensure that you set `self` to the return value from invocation of
// `super`’s implementation, as shown in the following example:
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(entity:insertInto:)
func NewManagedObjectWithEntityInsertIntoManagedObjectContext(entity INSEntityDescription, context INSManagedObjectContext) NSManagedObject {
	instance := getNSManagedObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), entity, context)
	return NSManagedObjectFromID(rv)
}

// Initializes a managed object from an entity description and inserts it into
// the specified managed object context.
//
// entity: The entity of which to create an instance.
//
// The model associated with `context`’s persistent store coordinator must
// contain `entity`.
//
// context: The context into which the new instance is inserted.
//
// # Return Value
//
// An initialized instance of the appropriate class for `entity`.
//
// # Discussion
//
// [NSManagedObject] uses dynamic class generation to support the Objective-C
// 2 properties feature (see Declared Properties) by automatically creating a
// subclass of the class appropriate for `entity`. “ therefore returns an
// instance of the appropriate class for `entity`. The dynamically-generated
// subclass will be based on the class specified by the entity, so specifying
// a custom class in your model will supersede the class passed to `alloc`.
//
// If `context` is not `nil`, this method invokes `[context self]` (which
// causes [NSManagedObject.AwakeFromInsert] to be invoked).
//
// You are discouraged from overriding this method—you should instead
// override [NSManagedObject.AwakeFromInsert] and/or
// [NSManagedObject.AwakeFromFetch] (if there is logic common to these
// methods, it should be factored into a third method which is invoked from
// both). If you do perform custom initialization in this method, you may
// cause problems with undo and redo operations.
//
// In many applications, there is no need to subsequently assign a
// newly-created managed object to a particular store—see
// [NSManagedObjectContext.AssignObjectToPersistentStore]. If your application
// has multiple stores and you do need to assign an object to a specific
// store, you should not do so in a managed object’s initializer method.
// Such an assignment is controller- not model-level logic.
//
// # Special Considerations
//
// If you override
// [NSManagedObject.InitWithEntityInsertIntoManagedObjectContext], you must
// ensure that you set `self` to the return value from invocation of
// `super`’s implementation, as shown in the following example:
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(entity:insertInto:)
func (m NSManagedObject) InitWithEntityInsertIntoManagedObjectContext(entity INSEntityDescription, context INSManagedObjectContext) NSManagedObject {
	rv := objc.Send[NSManagedObject](m.ID, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), entity, context)
	return rv
}

// Initializes a managed object subclass and inserts it into the specified
// managed object context.
//
// # Return Value
//
// An initialized instance of the appropriate subclass.
//
// # Discussion
//
// This method is only legal to call on subclasses of [NSManagedObject] that
// represent a single entity in the model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(context:)
func (m NSManagedObject) InitWithContext(moc INSManagedObjectContext) NSManagedObject {
	rv := objc.Send[NSManagedObject](m.ID, objc.Sel("initWithContext:"), moc)
	return rv
}

// Returns a Boolean value that indicates whether the relationship for a given
// key is a fault.
//
// key: The name of one of the receiver’s relationships.
//
// # Return Value
//
// true if the relationship for `key` is a fault; otherwise, false.
//
// # Discussion
//
// If the specified relationship is a fault, calling this method does not
// result in the fault firing.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/hasFault(forRelationshipNamed:)
func (m NSManagedObject) HasFaultForRelationshipNamed(key string) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasFaultForRelationshipNamed:"), objc.String(key))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed
// object when fufilling it from a fault.
//
// # Discussion
//
// You typically use this method to compute derived values or to recreate
// transient relationships from the receiver’s persistent properties.
//
// The managed object context’s change processing is explicitly disabled
// around this method so that you can use public setters to establish
// transient values and other caches without dirtying the object or its
// context. Because of this, however, you should not modify relationships in
// this method as the inverse will not be set.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/awakeFromFetch()
func (m NSManagedObject) AwakeFromFetch() {
	objc.Send[objc.ID](m.ID, objc.Sel("awakeFromFetch"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object when initially creating it.
//
// # Discussion
//
// You typically use this method to initialize special default property
// values. This method is invoked only once in the object’s lifetime.
//
// If you want to set attribute values in an implementation of this method,
// you should typically use primitive accessor methods (either
// [NSManagedObject.SetPrimitiveValueForKey] or—better—the appropriate
// custom primitive accessors). This ensures that the new values are treated
// as baseline values rather than being recorded as undoable changes for the
// properties in question.
//
// # Special Considerations
//
// If you create a managed object then perform undo operations to bring the
// managed object context to a state prior to the object’s creation, then
// perform redo operations to bring the managed object context back to a state
// after the object’s creation, [NSManagedObject.AwakeFromInsert] is not
// invoked a second time.
//
// You are typically discouraged from performing fetches within an
// implementation of [NSManagedObject.AwakeFromInsert]. Although it is
// allowed, execution of the fetch request can trigger the sending of internal
// Core Data notifications which may have unwanted side-effects. For example,
// in macOS, an instance of [NSArrayController] may end up inserting a new
// object into its content array twice.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/awakeFromInsert()
//
// [NSArrayController]: https://developer.apple.com/documentation/AppKit/NSArrayController
func (m NSManagedObject) AwakeFromInsert() {
	objc.Send[objc.ID](m.ID, objc.Sel("awakeFromInsert"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object when fulfilling it from a snapshot.
//
// flags: A bit mask of [NSManagedObject.DidChangeValueForKey] constants to denote
// the event or events that led to the method being invoked.
//
// For possible values, see [NSSnapshotEventType].
//
// # Discussion
//
// You typically use this method to compute derived values or to recreate
// transient relationships from the receiver’s persistent properties.
//
// If you want to set attribute values and need to avoid emitting key-value
// observation change notifications, you should use primitive accessor methods
// (either [NSManagedObject.SetPrimitiveValueForKey] or—better—the
// appropriate custom primitive accessors). This ensures that the new values
// are treated as baseline values rather than being recorded as undoable
// changes for the properties in question.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/awake(fromSnapshotEvents:)
//
// [NSSnapshotEventType]: https://developer.apple.com/documentation/CoreData/NSSnapshotEventType
func (m NSManagedObject) AwakeFromSnapshotEvents(flags NSSnapshotEventType) {
	objc.Send[objc.ID](m.ID, objc.Sel("awakeFromSnapshotEvents:"), flags)
}

// Returns a dictionary containing the keys and new values of persistent
// properties with changes since the last fetching or saving of the managed
// object.
//
// # Return Value
//
// A dictionary with keys that are the names of persistent properties with
// changes since last fetching or saving the receiver, and with the new values
// for those properties.
//
// # Discussion
//
// This method only reports changes to properties that are persistent
// properties of the receiver, not changes to transient properties or custom
// instance variables.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/changedValues()
func (m NSManagedObject) ChangedValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("changedValues"))
	return foundation.NSDictionaryFromID(rv)
}

// Returns a dictionary containing the keys and new values of persistent
// properties with changes since the last fetching or saving of the managed
// object.
//
// # Return Value
//
// A dictionary with keys that are the names of persistent properties with
// changes since the last posting of [NSManagedObjectContextObjectsDidChange],
// and with the new values for those properties.
//
// # Discussion
//
// This method only reports changes to properties that are persistent
// properties of the receiver, not changes to transient properties or custom
// instance variables.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/changedValuesForCurrentEvent()
//
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
func (m NSManagedObject) ChangedValuesForCurrentEvent() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("changedValuesForCurrentEvent"))
	return foundation.NSDictionaryFromID(rv)
}

// Returns a dictionary of the most recent fetched or saved values of the
// managed object for the properties of the specified keys.
//
// keys: An array containing names of properties of the receiver, or `nil`.
//
// # Return Value
//
// A dictionary containing the last fetched or saved values of the receiver
// for the properties specified by `keys`.
//
// # Discussion
//
// `nil` values are represented by an instance of [NSNull].
//
// This method only reports values of properties that are defined as
// persistent properties of the receiver, not values of transient properties
// or of custom instance variables.
//
// You can invoke this method with the `keys` value of `nil` to retrieve
// committed values for all the receiver’s properties, as illustrated by the
// following example.
//
// It is more efficient to use `nil` than to pass an array of all the property
// keys.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/committedValues(forKeys:)
//
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
func (m NSManagedObject) CommittedValuesForKeys(keys []string) foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("committedValuesForKeys:"), objectivec.StringSliceToNSArray(keys))
	return foundation.NSDictionaryFromID(rv)
}

// Provides an opportunity to add code into the life cycle of the managed
// object before deleting it.
//
// # Discussion
//
// You can implement this method to perform any operations required before the
// object is deleted, such as custom propagation before relationships are torn
// down, or reconfiguration of objects using key-value observing.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/prepareForDeletion()
func (m NSManagedObject) PrepareForDeletion() {
	objc.Send[objc.ID](m.ID, objc.Sel("prepareForDeletion"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object before saving it.
//
// # Discussion
//
// This method can have “side effects” on persistent values. You can use
// it to, for example, compute persistent values from other transient or
// scratchpad values.
//
// If you want to update a persistent property value, you should typically
// test for equality of any new value with the existing value before making a
// change. If you change property values using standard accessor methods, Core
// Data will observe the resultant change notification and so invoke
// `willSave` again before saving the object’s managed object context. If
// you continue to modify a value in `willSave`, `willSave` will continue to
// be called until your program crashes.
//
// For example, if you set a last-modified timestamp, you should check whether
// either you previously set it in the same save operation, or that the
// existing timestamp is not less than a small delta from the current time.
// Typically it’s better to calculate the timestamp once for all the objects
// being saved (for example, in response to an
// [NSManagedObjectContextWillSaveNotification]).
//
// If you change property values using primitive accessors, you avoid the
// possibility of infinite recursion, but Core Data will not notice the change
// you make.
//
// The sense of “save” in the method name is that of a database commit
// statement and so applies to deletions as well as to updates to objects. For
// subclasses, this method is therefore an appropriate locus for code to be
// executed when an object deleted as well as “saved to disk.” You can
// find out if an object is marked for deletion with
// [NSManagedObject.Deleted].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/willSave()
func (m NSManagedObject) WillSave() {
	objc.Send[objc.ID](m.ID, objc.Sel("willSave"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object after the managed object’s context completes a save operation.
//
// # Discussion
//
// You can use this method to notify other objects after a save, and to
// compute transient values from persistent values.
//
// This method can have “side effects” on the persistent values, however
// any changes you make using standard accessor methods will by default dirty
// the managed object context and leave your context with unsaved changes.
// Moreover, if the object’s context has an undo manager, such changes will
// add an undo operation. For document-based applications, changes made in
// `didSave` will therefore come into the next undo grouping, which can lead
// to “empty” undo operations from the user’s perspective. You may want
// to disable undo registration to avoid this issue.
//
// The sense of “save” in the method name is that of a database commit
// statement and so applies to deletions as well as to updates to objects. For
// subclasses, this method is therefore an appropriate locus for code to be
// executed when an object deleted as well as “saved to disk.” You can
// find out if an object is marked for deletion with
// [NSManagedObject.Deleted].
//
// # Special Considerations
//
// You cannot attempt to resurrect a deleted object in `didSave`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/didSave()
func (m NSManagedObject) DidSave() {
	objc.Send[objc.ID](m.ID, objc.Sel("didSave"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object before converting it to a fault.
//
// # Discussion
//
// This method is the companion of the [NSManagedObject.DidTurnIntoFault]
// method. You can use it to (re)set state which requires access to property
// values (for example, observers across key paths). The default
// implementation does nothing.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/willTurnIntoFault()
func (m NSManagedObject) WillTurnIntoFault() {
	objc.Send[objc.ID](m.ID, objc.Sel("willTurnIntoFault"))
}

// Provides an opportunity to add code into the life cycle of the managed
// object after converting it to a fault.
//
// # Discussion
//
// You use this method to clear out custom data caches—transient values
// declared as entity properties are typically already cleared out by the time
// this method is invoked (see, for example,
// [NSManagedObjectContext.RefreshObjectMergeChanges]).
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/didTurnIntoFault()
func (m NSManagedObject) DidTurnIntoFault() {
	objc.Send[objc.ID](m.ID, objc.Sel("didTurnIntoFault"))
}

// Returns the value for the specified property from the managed object’s
// private internal storage .
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
//
// The value of the property specified by `key`. Returns `nil` if no value has
// been set.
//
// # Discussion
//
// This method does not invoke the access notification methods
// ([NSManagedObject.WillAccessValueForKey] and
// [NSManagedObject.DidAccessValueForKey]). This method is used primarily by
// subclasses that implement custom accessor methods that need direct access
// to the receiver’s private storage.
//
// # Special Considerations
//
// Subclasses should not override this method.
//
// The following points also apply:
//
// - Primitive accessor methods are only supported on modeled properties. If
// you invoke a primitive accessor on an unmodeled property, it will instead
// operate upon a random modeled property. (The debug libraries and frameworks
// (available from [Apple Developer Website]) have assertions to test for
// passing unmodeled keys to these methods.) - You are strongly encouraged to
// use the dynamically-generated accessors rather than using this method
// directly (for example, “ instead of `@"name"`). The dynamic accessors are
// much more efficient, and allow for compile-time checking.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/primitiveValue(forKey:)
//
// [Apple Developer Website]: http://developer.apple.com/
func (m NSManagedObject) PrimitiveValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("primitiveValueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Sets the value of a given property in the managed object’s private
// internal storage.
//
// value: The new value for the property specified by `key`.
//
// key: The name of one of the receiver’s properties.
//
// # Discussion
//
// Sets in the receiver’s private internal storage the value of the property
// specified by `key` to `value`. If `key` identifies a to-one relationship,
// relates the object specified by `value` to the receiver, unrelating the
// previously related object if there was one. Given a collection object and a
// key that identifies a to-many relationship, relates the objects contained
// in the collection to the receiver, unrelating previously related objects if
// there were any.
//
// This method does not invoke the change notification methods
// ([willChangeValue(forKey:)] and [didChangeValue(forKey:)]). It is typically
// used by subclasses that implement custom accessor methods that need direct
// access to the receiver’s private internal storage. It is also used by the
// Core Data framework to initialize the receiver with values from a
// persistent store or to restore a value from a snapshot.
//
// # Special Considerations
//
// You must not override this method.
//
// You should typically use this method only to modify attributes (usually
// transient), not relationships. If you try to set a to-many relationship to
// a new [NSMutableSet] object, it will (eventually) fail. In the unusual
// event that you need to modify a relationship using this method, you first
// get the existing set using “ (ensure the method does not return `nil`),
// create a mutable copy, and then modify the copy—as illustrated in the
// following example:
//
// If the relationship is bi-directional (that is, if an inverse relationship
// is specified) then you are also responsible for maintaining the inverse
// relationship (regardless of cardinality)—in contrast with Core Data’s
// normal behavior described in Using Managed Objects.
//
// The following points also apply:
//
// - Primitive accessor methods are only supported on modeled properties. If
// you invoke a primitive accessor on an unmodeled property, it will instead
// operate upon a random modeled property. (The debug libraries and frameworks
// from (available from the [Apple Developer Website]) have assertions to test
// for passing unmodeled keys to these methods.) - You are strongly encouraged
// to use the dynamically-generated accessors rather than using this method
// directly (for example, “ instead of `newName @"name"`). The dynamic
// accessors are much more efficient, and allow for compile-time checking.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/setPrimitiveValue(_:forKey:)
//
// [Apple Developer Website]: http://developer.apple.com/
// [didChangeValue(forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChangeValue(forKey:)
// [willChangeValue(forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/willChangeValue(forKey:)
func (m NSManagedObject) SetPrimitiveValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setPrimitiveValue:forKey:"), value, objc.String(key))
}

// Returns the object IDs for all of the managed objects that are in the named
// relationship.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/objectIDs(forRelationshipNamed:)
func (m NSManagedObject) ObjectIDsForRelationshipNamed(key string) []NSManagedObjectID {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("objectIDsForRelationshipNamed:"), objc.String(key))
	return objc.ConvertSlice(rv, func(id objc.ID) NSManagedObjectID {
		return NSManagedObjectIDFromID(id)
	})
}

// Determines whether the managed object can be deleted in its current state.
//
// # Discussion
//
// An object cannot be deleted if it has a relationship has a “deny”
// delete rule and that relationship has a destination object.
//
// [NSManagedObject]‘s implementation sends the receiver’s entity
// description a message which performs basic checking based on the presence
// or absence of values.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/validateForDelete()
func (m NSManagedObject) ValidateForDelete() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateForDelete:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateForDelete: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines whether the managed object can be inserted in its current state.
//
// # Discussion
//
// Subclasses should invoke super’s implementation before performing their
// own validation, and should combine any error returned by super’s
// implementation with their own (see Managed Object Validation).
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/validateForInsert()
func (m NSManagedObject) ValidateForInsert() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateForInsert:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateForInsert: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines whether the managed object’s current state is valid.
//
// # Discussion
//
// [NSManagedObject]‘s implementation iterates through all of the
// receiver’s properties validating each in turn. If this results in more
// than one error, the `userInfo` dictionary in the [NSError] returned in
// `error` contains a key [NSDetailedErrorsKey]; the corresponding value is an
// array containing the individual validation errors. If you pass [NULL] as
// the error, validation will abort after the first failure.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/validateForUpdate()
func (m NSManagedObject) ValidateForUpdate() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateForUpdate:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateForUpdate: returned NO with nil NSError")
	}
	return rv, nil

}

// Provides support for key-value observing access notification.
//
// key: The name of one of the receiver’s properties.
//
// # Discussion
//
// Together with [NSManagedObject.WillAccessValueForKey], this method is used
// to fire faults, to maintain inverse relationships, and so on. Each read
// access must be wrapped in this method pair (in the same way that each write
// access must be wrapped in the “/“ method pair). In the default
// implementation of [NSManagedObject] these methods are invoked for you
// automatically. If, say, you create a custom subclass that uses explicit
// instance variables, you must invoke them yourself, as in the following
// example.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/didAccessValue(forKey:)
func (m NSManagedObject) DidAccessValueForKey(key string) {
	objc.Send[objc.ID](m.ID, objc.Sel("didAccessValueForKey:"), objc.String(key))
}

// Returns the observation info of the managed object.
//
// # Return Value
//
// The observation info of the receiver.
//
// # Discussion
//
// For more about key-value observation, see [Key-Value Observing Programming
// Guide].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/observationInfo()
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (m NSManagedObject) ObservationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("observationInfo"))
	return rv
}

// Sets the observation info of the managed object.
//
// inObservationInfo: The new observation info for the receiver.
//
// # Discussion
//
// For more about observation information, see [Key-Value Observing
// Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/setObservationInfo(_:)
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (m NSManagedObject) SetObservationInfo(inObservationInfo unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObservationInfo:"), inObservationInfo)
}

// Provides support for key-value observing access notification.
//
// key: The name of one of the receiver’s properties.
//
// # Discussion
//
// See [NSManagedObject.DidAccessValueForKey] for more details. You can invoke
// this method with the `key` value of `nil` to ensure that a fault has been
// fired, as illustrated by the following example.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/willAccessValue(forKey:)
func (m NSManagedObject) WillAccessValueForKey(key string) {
	objc.Send[objc.ID](m.ID, objc.Sel("willAccessValueForKey:"), objc.String(key))
}

// Provides an opportunity to respond when a change was made to a specified
// to-many relationship.
//
// inKey: The name of a property that is a to-many relationship.
//
// inMutationKind: The type of change that was made.
//
// inObjects: The objects that were involved in the change (see
// [NSKeyValueSetMutationKind]).
//
// # Discussion
//
// For more details, see [Key-Value Observing Programming Guide].
//
// You must not override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/didChangeValue(forKey:withSetMutation:using:)
//
// [NSKeyValueSetMutationKind]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (m NSManagedObject) DidChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind uint, inObjects foundation.INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("didChangeValueForKey:withSetMutation:usingObjects:"), objc.String(inKey), inMutationKind, inObjects)
}

// Provides an opportunity to respond when a change is about to be made to a
// specified to-many relationship.
//
// inKey: The name of a property that is a to-many relationship
//
// inMutationKind: The type of change that will be made.
//
// inObjects: The objects that were involved in the change (see
// [NSKeyValueSetMutationKind]).
//
// # Discussion
//
// For more details, see [Key-Value Observing Programming Guide].
//
// You must not override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/willChangeValue(forKey:withSetMutation:using:)
//
// [NSKeyValueSetMutationKind]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (m NSManagedObject) WillChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind uint, inObjects foundation.INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("willChangeValueForKey:withSetMutation:usingObjects:"), objc.String(inKey), inMutationKind, inObjects)
}

// Returns the entity description that is associated with this subclass.
//
// # Discussion
//
// This method is only legal to call on subclasses of [NSManagedObject] that
// represent a single entity in the model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/entity()
func (_NSManagedObjectClass NSManagedObjectClass) Entity() NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSManagedObjectClass.class), objc.Sel("entity"))
	return NSEntityDescriptionFromID(rv)
}

// Returns an initialized fetch request with the entity this subclass
// represents.
//
// # Discussion
//
// This method is only legal to call on subclasses of [NSManagedObject] that
// represent a single entity in the model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/fetchRequest()
func (_NSManagedObjectClass NSManagedObjectClass) FetchRequest() NSFetchRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSManagedObjectClass.class), objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(rv)
}

// The entity description of the managed object.
//
// # Discussion
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/entity-swift.property
func (m NSManagedObject) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// The object ID of the managed object.
//
// # Discussion
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/objectID
func (m NSManagedObject) ObjectID() INSManagedObjectID {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectID"))
	return NSManagedObjectIDFromID(objc.ID(rv))
}

// The managed object context with which the managed object is registered.
//
// # Discussion
//
// May be `nil` if the receiver has been deleted from its context.
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/managedObjectContext
func (m NSManagedObject) ManagedObjectContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("managedObjectContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the managed object has been
// inserted, has been deleted, or has unsaved changes.
//
// # Discussion
//
// true if the receiver has been inserted, has been deleted, or has unsaved
// changes, otherwise false. The result is the equivalent of OR-ing the values
// of [NSManagedObject.Inserted], [NSManagedObject.Deleted], and
// [NSManagedObject.Updated].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/hasChanges
func (m NSManagedObject) HasChanges() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasChanges"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted
// in a managed object context.
//
// # Discussion
//
// true if the receiver has been inserted in a managed object context,
// otherwise false. If the receiver is a fault, accessing this property does
// not cause it to fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/isInserted
func (m NSManagedObject) IsInserted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isInserted"))
	return rv
}

// A Boolean value that indicates whether the managed object has unsaved
// changes.
//
// # Discussion
//
// true if the receiver has unsaved changes, otherwise false. The receiver has
// unsaved changes if it has been updated since its managed object context was
// last saved.
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/isUpdated
func (m NSManagedObject) IsUpdated() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUpdated"))
	return rv
}

// A Boolean value that indicates whether the managed object will be deleted
// during the next save.
//
// # Discussion
//
// true if Core Data will ask the persistent store to delete the object during
// the next save operation, otherwise false. It may return false at other
// times, particularly after the object has been deleted. The immediacy with
// which it will stop returning true depends on where the object is in the
// process of being deleted.
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/isDeleted
func (m NSManagedObject) IsDeleted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isDeleted"))
	return rv
}

// A Boolean value that indicates whether the managed object is a fault.
//
// # Discussion
//
// true if the receiver is a fault, otherwise false. Knowing whether an object
// is a fault is useful in many situations when computations are optional. It
// can also be used to avoid growing the object graph unnecessarily (which may
// improve performance as it can avoid time-consuming fetches from data
// stores).
//
// If this property is false, then the receiver’s data must be in memory.
// However, if this property is true, it does not mean that the data is not in
// memory. The data may be in memory, or it may not, depending on many factors
// influencing caching.
//
// If the receiver is a fault, accessing this property does not cause it to
// fire.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/isFault
func (m NSManagedObject) IsFault() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isFault"))
	return rv
}

// The faulting state of the managed object.
//
// # Return Value
//
// `0` if the object is fully initialized as a managed object and not
// transitioning to or from another state, otherwise some other value.
//
// # Discussion
//
// `0` if the object is fully initialized as a managed object and not
// transitioning to or from another state, otherwise some other value. This
// property allows you to determine if an object is in a transitional phase
// when receiving a key-value observing change notification.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/faultingState
func (m NSManagedObject) FaultingState() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("faultingState"))
	return rv
}

// A Boolean value that indicates whether the managed object has persistent
// changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/hasPersistentChangedValues
func (m NSManagedObject) HasPersistentChangedValues() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasPersistentChangedValues"))
	return rv
}

// A Boolean value that indicates whether to mark instances of the class as
// having changes when an unmodeled property changes.
//
// # Return Value
//
// false if instances of the class should be marked as having changes if an
// unmodeled property is changed, otherwise true.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObject/contextShouldIgnoreUnmodeledPropertyChanges
func (_NSManagedObjectClass NSManagedObjectClass) ContextShouldIgnoreUnmodeledPropertyChanges() bool {
	rv := objc.Send[bool](objc.ID(_NSManagedObjectClass.class), objc.Sel("contextShouldIgnoreUnmodeledPropertyChanges"))
	return rv
}

// Protocol methods for NSFetchRequestResult
