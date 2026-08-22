// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSManagedObjectContext] class.
var (
	_NSManagedObjectContextClass     NSManagedObjectContextClass
	_NSManagedObjectContextClassOnce sync.Once
)

func getNSManagedObjectContextClass() NSManagedObjectContextClass {
	_NSManagedObjectContextClassOnce.Do(func() {
		_NSManagedObjectContextClass = NSManagedObjectContextClass{class: objc.GetClass("NSManagedObjectContext")}
	})
	return _NSManagedObjectContextClass
}

// GetNSManagedObjectContextClass returns the class object for NSManagedObjectContext.
func GetNSManagedObjectContextClass() NSManagedObjectContextClass {
	return getNSManagedObjectContextClass()
}

type NSManagedObjectContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSManagedObjectContextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSManagedObjectContextClass) Alloc() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object space to manipulate and track changes to managed objects.
//
// # Overview
//
// A context consists of a group of related model objects that represent an
// internally consistent view of one or more persistent stores. Changes to
// managed objects remain in memory in the associated context until Core Data
// saves that context to one or more persistent stores. A single managed
// object instance exists in one and only one context, but multiple copies of
// an object can exist in different contexts. Therefore, an object is unique
// to a particular context.
//
// # Life cycle management
//
// The context is a powerful object with a central role in the life cycle of
// managed objects, with responsibilities from life cycle management
// (including faulting) to validation, inverse relationship handling, and
// undo/redo. Through a context you can retrieve or “fetch” objects from a
// persistent store, make changes to those objects, and then either discard
// the changes or—again through the context—commit them back to the
// persistent store. The context is responsible for watching for changes in
// its objects and maintains an undo manager so you can have finer-grained
// control over undo and redo. You can insert new objects and delete ones you
// have fetched, and commit these modifications to the persistent store.
//
// All objects fetched from an external store are registered in a context
// together with a global identifier (an instance of [NSManagedObjectID])
// that’s used to uniquely identify each object to the external store.
//
// # Parent store
//
// Managed object contexts have a parent store from which they retrieve data
// representing managed objects and through which they commit changes to
// managed objects.
//
// Prior to OS X v10.7 and iOS v5.0, the parent store is always a persistent
// store coordinator. In macOS 10.7 and later and iOS v5.0 and later, the
// parent store may be another managed object context. Ultimately the root of
// a context’s ancestry must be a persistent store coordinator. The
// coordinator provides the managed object model and dispatches requests to
// the various persistent stores containing the data.
//
// If a context’s parent store is another managed object context, fetch and
// save operations are mediated by the parent context instead of a
// coordinator. This pattern has a number of usage scenarios, including:
//
// - Performing background operations on a second thread or queue. - Managing
// discardable edits, such as in an inspector window or view.
//
// As the first scenario implies, a parent context can service requests from
// children on different threads. You cannot, therefore, use parent contexts
// created with the thread confinement type (see [NSManagedObjectContext]).
//
// When you save changes in a context, the changes are only committed “one
// store up.” If you save a child context, changes are pushed to its parent.
// Changes are not saved to the persistent store until the root context is
// saved. (A root managed object context is one whose parent context is
// `nil`.) In addition, a parent does not pull changes from children before it
// saves. You must save a child context if you want ultimately to commit the
// changes.
//
// # Notifications
//
// A context posts notifications at various points—see
// [NSManagedObjectContextObjectsDidChange] for example. Typically, you should
// register to receive these notifications only from known contexts:
//
// Several system frameworks use Core Data internally. If you register to
// receive these notifications from all contexts (by passing `nil` as the
// object parameter to a method such as
// [addObserver(_:selector:name:object:)]), then you may receive unexpected
// notifications that are difficult to handle.
//
// # Concurrency
//
// Core Data uses thread (or serialized queue) confinement to protect managed
// objects and managed object contexts (see [Core Data Programming Guide]). A
// consequence of this is that a context assumes the default owner is the
// thread or queue that creates it. Don’t, therefore, initialize a context
// on one thread then pass it to another. Instead, pass a reference to a
// persistent store coordinator and have the receiving thread or queue create
// a new context using that. If you use [Operation], you must create the
// context in [main()] (for a serial queue) or [start()] (for a concurrent
// queue).
//
// When you create a context you specify the concurrency type with which
// you’ll use it. When you create a managed object context, you have two
// options for its thread (queue) association:
//
// - Private: The context creates and manages a private queue. - Main: The
// context associates with the main queue and is dependent on the
// application’s event loop; otherwise, it’s similar to a private context.
// Use this type for contexts that update view controllers and other user
// interface elements.
//
// You use contexts using the queue-based concurrency types in conjunction
// with [NSManagedObjectContext.PerformBlock] and
// [NSManagedObjectContext.PerformBlockAndWait]. You group “standard”
// messages to send to the context within a block to pass to one of these
// methods. There are two exceptions:
//
// - Setter methods on queue-based managed object contexts are thread-safe.
// You can invoke these methods directly on any thread. - If your code
// executes on the main thread, you can invoke methods on the main queue style
// contexts directly instead of using the block based API.
//
// [NSManagedObjectContext.PerformBlock] and
// [NSManagedObjectContext.PerformBlockAndWait] ensure the block operations
// execute on the correct queue for the context. The
// [NSManagedObjectContext.PerformBlock] method returns immediately and the
// context executes the block methods on its own thread. With the
// [NSManagedObjectContext.PerformBlockAndWait] method, the context still
// executes the block methods on its own thread, but the method doesn’t
// return until the block completes.
//
// It’s important to appreciate that blocks execute as a distinct body of
// work. As soon as your block ends, anyone else can enqueue another block,
// undo changes, reset the context, and so on. Thus blocks may be quite large,
// and typically end by invoking [NSManagedObjectContext.Save].
//
// You can also perform other operations, such as:
//
// # Subclassing notes
//
// You are strongly discouraged from subclassing [NSManagedObjectContext]. The
// change tracking and undo management mechanisms are highly optimized and
// hence intricate and delicate. Interposing your own additional logic that
// might impact [NSManagedObjectContext.ProcessPendingChanges] can have
// unforeseen consequences. In situations such as store migration, Core Data
// will create instances of [NSManagedObjectContext] for its own use. Under
// these circumstances, you cannot rely on any features of your custom
// subclass. Any [NSManagedObject] subclass must always be fully compatible
// with [NSManagedObjectContext] (that is, it cannot rely on features of a
// subclass of [NSManagedObjectContext]).
//
// # Creating a context
//
//   - [NSManagedObjectContext.InitWithConcurrencyType]: Creates a context that uses the specified concurrency type.
//
// # Configuring a context
//
//   - [NSManagedObjectContext.PersistentStoreCoordinator]: The persistent store coordinator of the context.
//   - [NSManagedObjectContext.SetPersistentStoreCoordinator]
//   - [NSManagedObjectContext.ParentContext]: The parent of the context.
//   - [NSManagedObjectContext.SetParentContext]
//   - [NSManagedObjectContext.Name]: The developer-provided name of the context.
//   - [NSManagedObjectContext.SetName]
//   - [NSManagedObjectContext.UserInfo]: The user information for the context.
//
// # Registering and fetching objects
//
//   - [NSManagedObjectContext.CountForFetchRequestError]: Returns the number of objects the specified request fetches when it executes.
//   - [NSManagedObjectContext.ObjectRegisteredForID]: Returns an object that exists in the context.
//   - [NSManagedObjectContext.ObjectWithID]: Returns either an existing object from the context or a fault that represents that object.
//   - [NSManagedObjectContext.ExistingObjectWithIDError]: Returns an existing object from either the context or the persistent store.
//   - [NSManagedObjectContext.RegisteredObjects]: The set of registered managed objects in the context.
//   - [NSManagedObjectContext.ExecuteRequestError]: Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result.
//   - [NSManagedObjectContext.RefreshAllObjects]: Refreshes all of the registered managed objects in the context.
//   - [NSManagedObjectContext.RetainsRegisteredObjects]: A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//   - [NSManagedObjectContext.SetRetainsRegisteredObjects]
//
// # Handling managed objects
//
//   - [NSManagedObjectContext.ShouldDeleteInaccessibleFaults]: A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//   - [NSManagedObjectContext.SetShouldDeleteInaccessibleFaults]
//   - [NSManagedObjectContext.InsertedObjects]: The set of objects that have been inserted into the context but not yet saved in a persistent store.
//   - [NSManagedObjectContext.UpdatedObjects]: The set of objects registered with the context that have uncommitted changes.
//   - [NSManagedObjectContext.DeletedObjects]: The set of objects that will be removed from their persistent store during the next save operation.
//   - [NSManagedObjectContext.ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty]: Creates a log of the inaccessible fault.
//   - [NSManagedObjectContext.InsertObject]: Registers an object to be inserted in the context’s persistent store the next time changes are saved.
//   - [NSManagedObjectContext.DeleteObject]: Specifies an object that should be removed from its persistent store when changes are committed.
//   - [NSManagedObjectContext.AssignObjectToPersistentStore]: Specifies the store in which a newly inserted object will be saved.
//   - [NSManagedObjectContext.ObtainPermanentIDsForObjectsError]: Converts to permanent IDs the object IDs of the objects in a given array.
//   - [NSManagedObjectContext.DetectConflictsForObject]: Marks an object for conflict detection.
//   - [NSManagedObjectContext.RefreshObjectMergeChanges]: Updates the persistent properties of a managed object to use the latest values from the persistent store.
//   - [NSManagedObjectContext.ProcessPendingChanges]: Forces the context to process changes to the object graph.
//
// # Managing concurrency
//
//   - [NSManagedObjectContext.AutomaticallyMergesChangesFromParent]: A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//   - [NSManagedObjectContext.SetAutomaticallyMergesChangesFromParent]
//   - [NSManagedObjectContext.ConcurrencyType]: The concurrency type for the context.
//   - [NSManagedObjectContext.MergePolicy]: The merge policy of the context.
//   - [NSManagedObjectContext.SetMergePolicy]
//   - [NSManagedObjectContext.QueryGenerationToken]: Returns the token associated with the query generation currently in use by this context.
//   - [NSManagedObjectContext.TransactionAuthor]: The author for the context that is used as an identifier in persistent history transactions.
//   - [NSManagedObjectContext.SetTransactionAuthor]
//   - [NSManagedObjectContext.MergeChangesFromContextDidSaveNotification]: Merges the changes specified in a given notification.
//   - [NSManagedObjectContext.SetQueryGenerationFromTokenError]: Sets the query generation this context should use.
//
// # Managing unsaved and uncommitted changes
//
//   - [NSManagedObjectContext.Save]: Attempts to commit unsaved changes to registered objects to the context’s parent store.
//   - [NSManagedObjectContext.HasChanges]: A Boolean value that indicates whether the context has uncommitted changes.
//
// # Undoing changes
//
//   - [NSManagedObjectContext.UndoManager]: The object that provides undo support for the context.
//   - [NSManagedObjectContext.SetUndoManager]
//   - [NSManagedObjectContext.Undo]: Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
//   - [NSManagedObjectContext.Redo]: Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph.
//   - [NSManagedObjectContext.Reset]: Returns the context to its base state.
//   - [NSManagedObjectContext.Rollback]: Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values.
//
// # Handling delete propagation
//
//   - [NSManagedObjectContext.PropagatesDeletesAtEndOfEvent]: A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//   - [NSManagedObjectContext.SetPropagatesDeletesAtEndOfEvent]
//
// # Managing the staleness interval
//
//   - [NSManagedObjectContext.StalenessInterval]: The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//   - [NSManagedObjectContext.SetStalenessInterval]
//
// # Performing block operations
//
//   - [NSManagedObjectContext.PerformBlock]: Asynchronously performs the specified closure on the context’s queue.
//   - [NSManagedObjectContext.PerformBlockAndWait]: Synchronously performs the specified closure on the context’s queue.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext
//
// [Core Data Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreData/index.html#//apple_ref/doc/uid/TP40001075
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
// [Operation]: https://developer.apple.com/documentation/Foundation/Operation
// [addObserver(_:selector:name:object:)]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
// [main()]: https://developer.apple.com/documentation/Foundation/Operation/main()
// [start()]: https://developer.apple.com/documentation/Foundation/Operation/start()
type NSManagedObjectContext struct {
	objectivec.Object
}

// NSManagedObjectContextFromID constructs a [NSManagedObjectContext] from an objc.ID.
//
// An object space to manipulate and track changes to managed objects.
func NSManagedObjectContextFromID(id objc.ID) NSManagedObjectContext {
	return NSManagedObjectContext{objectivec.Object{ID: id}}
}

// NOTE: NSManagedObjectContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSManagedObjectContext] class.
//
// # Creating a context
//
//   - [INSManagedObjectContext.InitWithConcurrencyType]: Creates a context that uses the specified concurrency type.
//
// # Configuring a context
//
//   - [INSManagedObjectContext.PersistentStoreCoordinator]: The persistent store coordinator of the context.
//   - [INSManagedObjectContext.SetPersistentStoreCoordinator]
//   - [INSManagedObjectContext.ParentContext]: The parent of the context.
//   - [INSManagedObjectContext.SetParentContext]
//   - [INSManagedObjectContext.Name]: The developer-provided name of the context.
//   - [INSManagedObjectContext.SetName]
//   - [INSManagedObjectContext.UserInfo]: The user information for the context.
//
// # Registering and fetching objects
//
//   - [INSManagedObjectContext.CountForFetchRequestError]: Returns the number of objects the specified request fetches when it executes.
//   - [INSManagedObjectContext.ObjectRegisteredForID]: Returns an object that exists in the context.
//   - [INSManagedObjectContext.ObjectWithID]: Returns either an existing object from the context or a fault that represents that object.
//   - [INSManagedObjectContext.ExistingObjectWithIDError]: Returns an existing object from either the context or the persistent store.
//   - [INSManagedObjectContext.RegisteredObjects]: The set of registered managed objects in the context.
//   - [INSManagedObjectContext.ExecuteRequestError]: Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result.
//   - [INSManagedObjectContext.RefreshAllObjects]: Refreshes all of the registered managed objects in the context.
//   - [INSManagedObjectContext.RetainsRegisteredObjects]: A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//   - [INSManagedObjectContext.SetRetainsRegisteredObjects]
//
// # Handling managed objects
//
//   - [INSManagedObjectContext.ShouldDeleteInaccessibleFaults]: A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//   - [INSManagedObjectContext.SetShouldDeleteInaccessibleFaults]
//   - [INSManagedObjectContext.InsertedObjects]: The set of objects that have been inserted into the context but not yet saved in a persistent store.
//   - [INSManagedObjectContext.UpdatedObjects]: The set of objects registered with the context that have uncommitted changes.
//   - [INSManagedObjectContext.DeletedObjects]: The set of objects that will be removed from their persistent store during the next save operation.
//   - [INSManagedObjectContext.ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty]: Creates a log of the inaccessible fault.
//   - [INSManagedObjectContext.InsertObject]: Registers an object to be inserted in the context’s persistent store the next time changes are saved.
//   - [INSManagedObjectContext.DeleteObject]: Specifies an object that should be removed from its persistent store when changes are committed.
//   - [INSManagedObjectContext.AssignObjectToPersistentStore]: Specifies the store in which a newly inserted object will be saved.
//   - [INSManagedObjectContext.ObtainPermanentIDsForObjectsError]: Converts to permanent IDs the object IDs of the objects in a given array.
//   - [INSManagedObjectContext.DetectConflictsForObject]: Marks an object for conflict detection.
//   - [INSManagedObjectContext.RefreshObjectMergeChanges]: Updates the persistent properties of a managed object to use the latest values from the persistent store.
//   - [INSManagedObjectContext.ProcessPendingChanges]: Forces the context to process changes to the object graph.
//
// # Managing concurrency
//
//   - [INSManagedObjectContext.AutomaticallyMergesChangesFromParent]: A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//   - [INSManagedObjectContext.SetAutomaticallyMergesChangesFromParent]
//   - [INSManagedObjectContext.ConcurrencyType]: The concurrency type for the context.
//   - [INSManagedObjectContext.MergePolicy]: The merge policy of the context.
//   - [INSManagedObjectContext.SetMergePolicy]
//   - [INSManagedObjectContext.QueryGenerationToken]: Returns the token associated with the query generation currently in use by this context.
//   - [INSManagedObjectContext.TransactionAuthor]: The author for the context that is used as an identifier in persistent history transactions.
//   - [INSManagedObjectContext.SetTransactionAuthor]
//   - [INSManagedObjectContext.MergeChangesFromContextDidSaveNotification]: Merges the changes specified in a given notification.
//   - [INSManagedObjectContext.SetQueryGenerationFromTokenError]: Sets the query generation this context should use.
//
// # Managing unsaved and uncommitted changes
//
//   - [INSManagedObjectContext.Save]: Attempts to commit unsaved changes to registered objects to the context’s parent store.
//   - [INSManagedObjectContext.HasChanges]: A Boolean value that indicates whether the context has uncommitted changes.
//
// # Undoing changes
//
//   - [INSManagedObjectContext.UndoManager]: The object that provides undo support for the context.
//   - [INSManagedObjectContext.SetUndoManager]
//   - [INSManagedObjectContext.Undo]: Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
//   - [INSManagedObjectContext.Redo]: Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph.
//   - [INSManagedObjectContext.Reset]: Returns the context to its base state.
//   - [INSManagedObjectContext.Rollback]: Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values.
//
// # Handling delete propagation
//
//   - [INSManagedObjectContext.PropagatesDeletesAtEndOfEvent]: A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//   - [INSManagedObjectContext.SetPropagatesDeletesAtEndOfEvent]
//
// # Managing the staleness interval
//
//   - [INSManagedObjectContext.StalenessInterval]: The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//   - [INSManagedObjectContext.SetStalenessInterval]
//
// # Performing block operations
//
//   - [INSManagedObjectContext.PerformBlock]: Asynchronously performs the specified closure on the context’s queue.
//   - [INSManagedObjectContext.PerformBlockAndWait]: Synchronously performs the specified closure on the context’s queue.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext
type INSManagedObjectContext interface {
	objectivec.IObject

	// Topic: Creating a context

	// Creates a context that uses the specified concurrency type.
	InitWithConcurrencyType(ct NSManagedObjectContextConcurrencyType) NSManagedObjectContext

	// Topic: Configuring a context

	// The persistent store coordinator of the context.
	PersistentStoreCoordinator() INSPersistentStoreCoordinator
	SetPersistentStoreCoordinator(value INSPersistentStoreCoordinator)
	// The parent of the context.
	ParentContext() INSManagedObjectContext
	SetParentContext(value INSManagedObjectContext)
	// The developer-provided name of the context.
	Name() string
	SetName(value string)
	// The user information for the context.
	UserInfo() foundation.INSDictionary

	// Topic: Registering and fetching objects

	// Returns the number of objects the specified request fetches when it executes.
	CountForFetchRequestError(request INSFetchRequest) (uint, error)
	// Returns an object that exists in the context.
	ObjectRegisteredForID(objectID INSManagedObjectID) INSManagedObject
	// Returns either an existing object from the context or a fault that represents that object.
	ObjectWithID(objectID INSManagedObjectID) INSManagedObject
	// Returns an existing object from either the context or the persistent store.
	ExistingObjectWithIDError(objectID INSManagedObjectID) (INSManagedObject, error)
	// The set of registered managed objects in the context.
	RegisteredObjects() foundation.INSSet
	// Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result.
	ExecuteRequestError(request INSPersistentStoreRequest) (INSPersistentStoreResult, error)
	// Refreshes all of the registered managed objects in the context.
	RefreshAllObjects()
	// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
	RetainsRegisteredObjects() bool
	SetRetainsRegisteredObjects(value bool)

	// Topic: Handling managed objects

	// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
	ShouldDeleteInaccessibleFaults() bool
	SetShouldDeleteInaccessibleFaults(value bool)
	// The set of objects that have been inserted into the context but not yet saved in a persistent store.
	InsertedObjects() foundation.INSSet
	// The set of objects registered with the context that have uncommitted changes.
	UpdatedObjects() foundation.INSSet
	// The set of objects that will be removed from their persistent store during the next save operation.
	DeletedObjects() foundation.INSSet
	// Creates a log of the inaccessible fault.
	ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault INSManagedObject, oid INSManagedObjectID, property INSPropertyDescription) bool
	// Registers an object to be inserted in the context’s persistent store the next time changes are saved.
	InsertObject(object INSManagedObject)
	// Specifies an object that should be removed from its persistent store when changes are committed.
	DeleteObject(object INSManagedObject)
	// Specifies the store in which a newly inserted object will be saved.
	AssignObjectToPersistentStore(object objectivec.IObject, store INSPersistentStore)
	// Converts to permanent IDs the object IDs of the objects in a given array.
	ObtainPermanentIDsForObjectsError(objects []NSManagedObject) (bool, error)
	// Marks an object for conflict detection.
	DetectConflictsForObject(object INSManagedObject)
	// Updates the persistent properties of a managed object to use the latest values from the persistent store.
	RefreshObjectMergeChanges(object INSManagedObject, flag bool)
	// Forces the context to process changes to the object graph.
	ProcessPendingChanges()

	// Topic: Managing concurrency

	// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
	AutomaticallyMergesChangesFromParent() bool
	SetAutomaticallyMergesChangesFromParent(value bool)
	// The concurrency type for the context.
	ConcurrencyType() NSManagedObjectContextConcurrencyType
	// The merge policy of the context.
	MergePolicy() objectivec.IObject
	SetMergePolicy(value objectivec.IObject)
	// Returns the token associated with the query generation currently in use by this context.
	QueryGenerationToken() INSQueryGenerationToken
	// The author for the context that is used as an identifier in persistent history transactions.
	TransactionAuthor() string
	SetTransactionAuthor(value string)
	// Merges the changes specified in a given notification.
	MergeChangesFromContextDidSaveNotification(notification foundation.NSNotification)
	// Sets the query generation this context should use.
	SetQueryGenerationFromTokenError(generation INSQueryGenerationToken) (bool, error)

	// Topic: Managing unsaved and uncommitted changes

	// Attempts to commit unsaved changes to registered objects to the context’s parent store.
	Save() (bool, error)
	// A Boolean value that indicates whether the context has uncommitted changes.
	HasChanges() bool

	// Topic: Undoing changes

	// The object that provides undo support for the context.
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	// Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
	Undo()
	// Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph.
	Redo()
	// Returns the context to its base state.
	Reset()
	// Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values.
	Rollback()

	// Topic: Handling delete propagation

	// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
	PropagatesDeletesAtEndOfEvent() bool
	SetPropagatesDeletesAtEndOfEvent(value bool)

	// Topic: Managing the staleness interval

	// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
	StalenessInterval() foundation.NSTimeInterval
	SetStalenessInterval(value foundation.NSTimeInterval)

	// Topic: Performing block operations

	// Asynchronously performs the specified closure on the context’s queue.
	PerformBlock(block VoidHandler)
	// Synchronously performs the specified closure on the context’s queue.
	PerformBlockAndWait(block VoidHandler)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m NSManagedObjectContext) Init() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSManagedObjectContext) Autorelease() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSManagedObjectContext creates a new NSManagedObjectContext instance.
func NewNSManagedObjectContext() NSManagedObjectContext {
	class := getNSManagedObjectContextClass()
	rv := objc.Send[NSManagedObjectContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a context that uses the specified concurrency type.
//
// ct: The context’s concurrency type. For possible values, see
// [NSManagedObjectContextConcurrencyType].
//
// # Discussion
//
// For more information, see [NSManagedObjectContext].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/init(concurrencyType:)
//
// [NSManagedObjectContextConcurrencyType]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType
func NewManagedObjectContextWithConcurrencyType(ct NSManagedObjectContextConcurrencyType) NSManagedObjectContext {
	instance := getNSManagedObjectContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConcurrencyType:"), ct)
	return NSManagedObjectContextFromID(rv)
}

// Creates a context that uses the specified concurrency type.
//
// ct: The context’s concurrency type. For possible values, see
// [NSManagedObjectContextConcurrencyType].
//
// # Discussion
//
// For more information, see [NSManagedObjectContext].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/init(concurrencyType:)
//
// [NSManagedObjectContextConcurrencyType]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType
func (m NSManagedObjectContext) InitWithConcurrencyType(ct NSManagedObjectContextConcurrencyType) NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m.ID, objc.Sel("initWithConcurrencyType:"), ct)
	return rv
}

// Returns the number of objects the specified request fetches when it
// executes.
//
// request: A fetch request that specifies the search criteria for the fetch.
//
// # Return Value
//
// The number of objects a given fetch request would have returned if it had
// been passed to [fetch(_:)], or [NSNotFound] if an error occurs.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/count(for:)-93zbm
//
// [fetch(_:)]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/fetch(_:)-38ys1
func (m NSManagedObjectContext) CountForFetchRequestError(request INSFetchRequest) (uint, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint](m.ID, objc.Sel("countForFetchRequest:error:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Returns an object that exists in the context.
//
// objectID: The identifier of the object to retrieve. For more information, see
// [NSManagedObjectID].
//
// # Return Value
//
// The identified object, if it’s known to the context; otherwise, `nil`.
//
// # Discussion
//
// This method provides a convenient way to retrieve an object from the
// context’s [NSManagedObjectContext.RegisteredObjects] property. A `nil`
// return value means the context doesn’t recognize the specified object;
// the object might still exist in the persistent store. If you need to query
// both the context and the store, use
// [NSManagedObjectContext.ExistingObjectWithIDError] instead.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObject(for:)
func (m NSManagedObjectContext) ObjectRegisteredForID(objectID INSManagedObjectID) INSManagedObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectRegisteredForID:"), objectID)
	return NSManagedObjectFromID(rv)
}

// Returns either an existing object from the context or a fault that
// represents that object.
//
// objectID: The identifier of the object to retrieve. For more information, see
// [NSManagedObjectID].
//
// # Return Value
//
// The identified object, if its known to the context; otherwise, a fault with
// its [NSManagedObject.ObjectID] property set to `objectID`.
//
// # Discussion
//
// If the context doesn’t recognize the specified object, this method
// returns a fault — a placeholder object that doesn’t load its properties
// until your code accesses them. The context then fetches the corresponding
// values from the persistent store and uses those values to turn the fault
// into a fully realized object.
//
// When this method returns a fault, Core Data makes no attempts to verify the
// existence of the underlying object in the persistent store. If the object
// doesn’t exist when the context tries to the fetch the object’s values,
// the framework throws an exception.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/object(with:)
func (m NSManagedObjectContext) ObjectWithID(objectID INSManagedObjectID) INSManagedObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectWithID:"), objectID)
	return NSManagedObjectFromID(rv)
}

// Returns an existing object from either the context or the persistent store.
//
// objectID: The identifier of the object to retrieve. For more information, see
// [NSManagedObjectID].
//
// # Return Value
//
// The identified object from either the context or the persistent store.
//
// # Discussion
//
// If the context recognizes the specified object, the method returns that
// object. Otherwise, the context fetches and returns a fully realized object
// from the persistent store; unlike [NSManagedObjectContext.ObjectWithID],
// this method never returns a fault. If the object doesn’t exist in both
// the context and the persistent store, the method throws an error.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/existingObject(with:)
func (m NSManagedObjectContext) ExistingObjectWithIDError(objectID INSManagedObjectID) (INSManagedObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("existingObjectWithID:error:"), objectID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSManagedObject{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSManagedObjectFromID(rv), nil

}

// Passes a request to the persistent store without affecting the contents of
// the managed object context, and returns a persistent store result.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/execute(_:)
func (m NSManagedObjectContext) ExecuteRequestError(request INSPersistentStoreRequest) (INSPersistentStoreResult, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("executeRequest:error:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentStoreResult{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentStoreResultFromID(rv), nil

}

// Refreshes all of the registered managed objects in the context.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refreshAllObjects()
func (m NSManagedObjectContext) RefreshAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("refreshAllObjects"))
}

// Creates a log of the inaccessible fault.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldHandleInaccessibleFault(_:for:triggeredByProperty:)
func (m NSManagedObjectContext) ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault INSManagedObject, oid INSManagedObjectID, property INSPropertyDescription) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shouldHandleInaccessibleFault:forObjectID:triggeredByProperty:"), fault, oid, property)
	return rv
}

// Registers an object to be inserted in the context’s persistent store the
// next time changes are saved.
//
// object: A managed object.
//
// # Discussion
//
// The managed object (`object`) is registered in the receiver with a
// temporary global ID. It is assigned a permanent global ID when changes are
// committed. If the current transaction is rolled back (for example, if the
// receiver is sent a [NSManagedObjectContext.Rollback] message) before a save
// operation, the object is unregistered from the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insert(_:)
func (m NSManagedObjectContext) InsertObject(object INSManagedObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertObject:"), object)
}

// Specifies an object that should be removed from its persistent store when
// changes are committed.
//
// object: A managed object.
//
// # Discussion
//
// When changes are committed, `object` will be removed from the uniquing
// tables. If `object` has not yet been saved to a persistent store, it is
// simply removed from the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/delete(_:)
func (m NSManagedObjectContext) DeleteObject(object INSManagedObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("deleteObject:"), object)
}

// Specifies the store in which a newly inserted object will be saved.
//
// object: A managed object.
//
// store: A persistent store.
//
// # Discussion
//
// You can obtain a store from the persistent store coordinator, using for
// example [NSPersistentStoreCoordinator.PersistentStoreForURL].
//
// # Special Considerations
//
// It is only necessary to use this method if the receiver’s persistent
// store coordinator manages multiple writable stores that have `object`‘s
// entity in their configuration. Maintaining configurations in the managed
// object model can eliminate the need for invoking this method directly in
// many situations. If the receiver’s persistent store coordinator manages
// only a single writable store, or if only one store has `object`’s entity
// in its model, `object` will automatically be assigned to that store.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/assign(_:to:)
func (m NSManagedObjectContext) AssignObjectToPersistentStore(object objectivec.IObject, store INSPersistentStore) {
	objc.Send[objc.ID](m.ID, objc.Sel("assignObject:toPersistentStore:"), object, store)
}

// Converts to permanent IDs the object IDs of the objects in a given array.
//
// objects: An array of managed objects.
//
// # Discussion
//
// This method converts the object ID of each managed object in `objects` to a
// permanent ID. Although the object will have a permanent ID, it will still
// respond positively to [NSManagedObject.Inserted] until it is saved. Any
// object that already has a permanent ID is ignored.
//
// Any object not already assigned to a store is assigned based on the same
// rules Core Data uses for assignment during a save operation (first writable
// store supporting the entity, and appropriate for the instance and its
// related items).
//
// # Special Considerations
//
// This method results in a transaction with the underlying store which
// changes the file’s modification date.
//
// In macOS, this results an additional consideration if you invoke this
// method on the managed object context associated with an instance of
// [NSPersistentDocument]. Instances of [NSDocument] need to know that they
// are in sync with the underlying content. To avoid problems, after invoking
// this method you must therefore update the document’s modification date
// (using [fileModificationDate]).
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/obtainPermanentIDs(for:)
//
// [NSPersistentDocument]: https://developer.apple.com/documentation/AppKit/NSPersistentDocument
// [fileModificationDate]: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (m NSManagedObjectContext) ObtainPermanentIDsForObjectsError(objects []NSManagedObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("obtainPermanentIDsForObjects:error:"), objectivec.IObjectSliceToNSArray(objects), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("obtainPermanentIDsForObjects:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Marks an object for conflict detection.
//
// object: A managed object.
//
// # Discussion
//
// If on the next invocation of [NSManagedObjectContext.Save] `object` has
// been modified in its persistent store, the save fails. This allows
// optimistic locking for unchanged objects. Conflict detection is always
// performed on changed or deleted objects.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/detectConflicts(for:)
func (m NSManagedObjectContext) DetectConflictsForObject(object INSManagedObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("detectConflictsForObject:"), object)
}

// Updates the persistent properties of a managed object to use the latest
// values from the persistent store.
//
// object: A managed object.
//
// flag: A Boolean value.
//
// If `flag` is false, the context discards pending changes and the managed
// object becomes a fault. Upon next access, the context reloads the
// object’s values from the persistent store or last cached state.
//
// If `flag` is true, the context reloads the object’s property values from
// the store or the cache. Then the context applies local changes over the
// newly loaded values. Merging the local values into `object` always
// succeeds, and never results in a merge conflict.
//
// # Discussion
//
// If you call this method before the
// [NSManagedObjectContext.StalenessInterval] expires, the context reloads the
// data from the cache instead of fetching from the store. If `flag` is true,
// this method doesn’t affect any transient properties. If `flag` is false,
// the object disposes the value of transient properties.
//
// You typically use this method to ensure data freshness if multiple managed
// object contexts share a single persistent store. You can use this method to
// resolve an optimistic locking failure when attempting to save.
//
// Turning `object` into a fault by setting `flag` to false breaks strong
// references to related managed objects. You can use this method to release a
// portion of your object graph if you want to constrain memory usage.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refresh(_:mergeChanges:)
func (m NSManagedObjectContext) RefreshObjectMergeChanges(object INSManagedObject, flag bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("refreshObject:mergeChanges:"), object, flag)
}

// Forces the context to process changes to the object graph.
//
// # Discussion
//
// This method causes changes to registered managed objects to be recorded
// with the undo manager.
//
// In AppKit-based applications, this method is invoked automatically at least
// once during the event loop (at the end of the loop)—it may be called more
// often than that if the framework needs to coalesce your changes before
// doing something else. You can also invoke it manually to coalesce any
// pending unprocessed changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/processPendingChanges()
func (m NSManagedObjectContext) ProcessPendingChanges() {
	objc.Send[objc.ID](m.ID, objc.Sel("processPendingChanges"))
}

// Merges the changes specified in a given notification.
//
// notification: An instance of an [NSManagedObjectContextDidSave] notification posted by
// another context.
//
// # Discussion
//
// This method refreshes any objects which have been updated in the other
// context, faults in any newly-inserted objects, and invokes
// [NSManagedObjectContext.DeleteObject]: on those which have been deleted.
//
// You can pass a [NSManagedObjectContextDidSave] notification posted by a
// managed object context on another thread, however you must not use the
// managed objects in the user info dictionary directly. For more details, see
// Concurrency with Core Data.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromContextDidSave:)
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
func (m NSManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification foundation.NSNotification) {
	objc.Send[objc.ID](m.ID, objc.Sel("mergeChangesFromContextDidSaveNotification:"), notification)
}

// Sets the query generation this context should use.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/setQueryGenerationFrom(_:)
func (m NSManagedObjectContext) SetQueryGenerationFromTokenError(generation INSQueryGenerationToken) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setQueryGenerationFromToken:error:"), generation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setQueryGenerationFromToken:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Attempts to commit unsaved changes to registered objects to the context’s
// parent store.
//
// # Discussion
//
// If there were multiple errors (for example several edited objects had
// validation failures) the description of [NSError] returned indicates that
// there were multiple errors, and its userInfo dictionary contains the key
// [NSDetailedErrors]. The value associated with the [NSDetailedErrors] key is
// an array that contains the individual [NSError] objects.
//
// If a context’s parent store is a persistent store coordinator, then
// changes are committed to the external store. If a context’s parent store
// is another managed object context, then [NSManagedObjectContext.Save] only
// updates managed objects in that parent store. To commit changes to the
// external store, you must save changes in the chain of contexts up to and
// including the context whose parent is the persistent store coordinator.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()
func (m NSManagedObjectContext) Save() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("save:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("save: returned NO with nil NSError")
	}
	return rv, nil

}

// Sends an undo message to the context’s undo manager, asking it to reverse
// the latest uncommitted changes applied to objects in the object graph.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undo()
func (m NSManagedObjectContext) Undo() {
	objc.Send[objc.ID](m.ID, objc.Sel("undo"))
}

// Sends a redo message to the context’s undo manager, asking it to reverse
// the latest undo operation applied to objects in the object graph.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/redo()
func (m NSManagedObjectContext) Redo() {
	objc.Send[objc.ID](m.ID, objc.Sel("redo"))
}

// Returns the context to its base state.
//
// # Discussion
//
// All the receiver’s managed objects are “forgotten.” If you use this
// method, you should ensure that you also discard references to any managed
// objects fetched using the receiver, since they will be invalid afterwards.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/reset()
func (m NSManagedObjectContext) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Removes everything from the undo stack, discards all insertions and
// deletions, and restores updated objects to their last committed values.
//
// # Discussion
//
// This method does not refetch data from the persistent store or stores.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/rollback()
func (m NSManagedObjectContext) Rollback() {
	objc.Send[objc.ID](m.ID, objc.Sel("rollback"))
}

// Asynchronously performs the specified closure on the context’s queue.
//
// block: The closure to perform.
//
// # Discussion
//
// This method encapsulates an autorelease pool and a call to
// [NSManagedObjectContext.ProcessPendingChanges].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/perform(_:)
func (m NSManagedObjectContext) PerformBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](m.ID, objc.Sel("performBlock:"), _block0)
}

// Synchronously performs the specified closure on the context’s queue.
//
// block: The closure to perform.
//
// # Discussion
//
// This method supports reentrancy — meaning it’s safe to call the method
// again, from within the closure, before the previous invocation completes.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/performAndWait(_:)-ypye
func (m NSManagedObjectContext) PerformBlockAndWait(block VoidHandler) {
	_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](m.ID, objc.Sel("performBlockAndWait:"), _block0)
}
func (m NSManagedObjectContext) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Handles changes from other processes or from a serialized state.
//
// # Discussion
//
// This method more efficiently merges changes into multiple contexts as well
// as nested contexts. The dictionary keys should be one or more from an
// [NSManagedObjectContextObjectsDidChange]: [NSInsertedObjectsKey],
// [NSUpdatedObjectsKey], [NSDeletedObjectsKey]. The values should be an
// [NSArray] of either [NSManagedObjectID] or [NSURL] objects conforming to
// valid results from [NSManagedObjectID.URIRepresentation].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromRemoteContextSave:into:)
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSDeletedObjectsKey]: https://developer.apple.com/documentation/CoreData/NSDeletedObjectsKey
// [NSInsertedObjectsKey]: https://developer.apple.com/documentation/CoreData/NSInsertedObjectsKey
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [NSUpdatedObjectsKey]: https://developer.apple.com/documentation/CoreData/NSUpdatedObjectsKey
func (_NSManagedObjectContextClass NSManagedObjectContextClass) MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData foundation.INSDictionary, contexts []NSManagedObjectContext) {
	objc.Send[objc.ID](objc.ID(_NSManagedObjectContextClass.class), objc.Sel("mergeChangesFromRemoteContextSave:intoContexts:"), changeNotificationData, objectivec.IObjectSliceToNSArray(contexts))
}

// The persistent store coordinator of the context.
//
// # Return Value
//
// The persistent store coordinator of the receiver.
//
// # Discussion
//
// The coordinator provides the managed object model and handles persistency.
// Note that multiple contexts can share a coordinator. May not be `nil`.
//
// Setting [NSManagedObjectContext.PersistentStoreCoordinator] to `nil` will
// raise an exception. If you want to “disconnect” a context from its
// persistent store coordinator, you should simply set all strong references
// to the context to `nil` and allow it to be deallocated normally.
//
// For more details, see [NSManagedObjectContext].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/persistentStoreCoordinator
func (m NSManagedObjectContext) PersistentStoreCoordinator() INSPersistentStoreCoordinator {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("persistentStoreCoordinator"))
	return NSPersistentStoreCoordinatorFromID(objc.ID(rv))
}
func (m NSManagedObjectContext) SetPersistentStoreCoordinator(value INSPersistentStoreCoordinator) {
	objc.Send[struct{}](m.ID, objc.Sel("setPersistentStoreCoordinator:"), value)
}

// The parent of the context.
//
// # Discussion
//
// `nil` indicates there is no parent context. For more details, see
// [NSManagedObjectContext].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/parent
func (m NSManagedObjectContext) ParentContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parentContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}
func (m NSManagedObjectContext) SetParentContext(value INSManagedObjectContext) {
	objc.Send[struct{}](m.ID, objc.Sel("setParentContext:"), value)
}

// The developer-provided name of the context.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/name
func (m NSManagedObjectContext) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSManagedObjectContext) SetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setName:"), objc.String(value))
}

// The user information for the context.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/userInfo
func (m NSManagedObjectContext) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The set of registered managed objects in the context.
//
// # Discussion
//
// A managed object context does not post key-value observing notifications
// when the return value of `registeredObjects` changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObjects
func (m NSManagedObjectContext) RegisteredObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("registeredObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the context keeps strong references
// to all registered managed objects.
//
// # Discussion
//
// If set to true, the receiver keeps strong references to all registered
// managed objects. If set to false, then the receiver keeps strong references
// to registered objects only when they are inserted, updated, deleted, or
// locked. The default is false.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/retainsRegisteredObjects
func (m NSManagedObjectContext) RetainsRegisteredObjects() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("retainsRegisteredObjects"))
	return rv
}
func (m NSManagedObjectContext) SetRetainsRegisteredObjects(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRetainsRegisteredObjects:"), value)
}

// A Boolean value that determines whether the context turns inaccessible
// faults into deleted objects.
//
// # Discussion
//
// Use this property to control how the context behaves when it encounters an
// inaccessible fault — an object with no underlying data in the persistent
// store. For example, you might fetch an object that has a to-many
// relationship, but then a background context deletes the related objects
// from the store before you traverse that relationship.
//
// When this property is set to true, the context returns a managed object
// with the following characteristics:
//
// - The object’s attributes, including scalars, nullable, and mandatory
// attributes are all set to `nil` or `0`. - The object’s
// [NSManagedObject.Deleted] property is set to true, which adds the object to
// the context’s [NSManagedObjectContext.DeletedObjects] set. - The object
// is exempt from validation rules, including optionality, because the object
// is nonexistent and the context discards it when you next call
// [NSManagedObjectContext.Save] or [NSManagedObjectContext.Reset].
//
// When the context returns an object with these characteristics, your app can
// continue running and process this object in the same way as any other
// deleted object.
//
// When this property is set to false, the context throws an exception.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults
func (m NSManagedObjectContext) ShouldDeleteInaccessibleFaults() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}
func (m NSManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}

// The set of objects that have been inserted into the context but not yet
// saved in a persistent store.
//
// # Discussion
//
// A managed object context does not post key-value observing notifications
// when the return value of `insertedObjects` changes—it does, however, post
// a [NSManagedObjectContextObjectsDidChange] notification when a change is
// made, and a [NSManagedObjectContextWillSave] and a
// [NSManagedObjectContextDidSave] notification before and after changes are
// committed respectively.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insertedObjects
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
// [NSManagedObjectContextWillSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextWillSave
func (m NSManagedObjectContext) InsertedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("insertedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The set of objects registered with the context that have uncommitted
// changes.
//
// # Discussion
//
// A managed object context does not post key-value observing notifications
// when the return value of `updatedObjects` changes. A context does, however,
// post a [NSManagedObjectContextObjectsDidChange] notification when a change
// is made, and a [NSManagedObjectContextWillSave] notification and a
// [NSManagedObjectContextDidSave] notification before and after changes are
// committed respectively.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/updatedObjects
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
// [NSManagedObjectContextWillSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextWillSave
func (m NSManagedObjectContext) UpdatedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updatedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The set of objects that will be removed from their persistent store during
// the next save operation.
//
// # Discussion
//
// The returned set does not necessarily include all the objects that have
// been deleted (using [NSManagedObjectContext.DeleteObject])—if an object
// has been inserted and deleted without an intervening save operation, it may
// not be included in the set.
//
// A managed object context does not post key-value observing notifications
// when the return value of `deletedObjects` changes. A context does, however,
// post a [NSManagedObjectContextObjectsDidChange] notification when a change
// is made, and a [NSManagedObjectContextWillSave] notification and a
// [NSManagedObjectContextDidSave] notification before and after changes are
// committed respectively (although again the set of deleted objects given for
// an [NSManagedObjectContextDidSave] does not include objects that were
// inserted and deleted without an intervening save operation—that is, they
// had never been saved to a persistent store).
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/deletedObjects
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
// [NSManagedObjectContextObjectsDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextObjectsDidChange
// [NSManagedObjectContextWillSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextWillSave
func (m NSManagedObjectContext) DeletedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deletedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the context automatically merges
// changes saved to its persistent store coordinator or parent context.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/automaticallyMergesChangesFromParent
func (m NSManagedObjectContext) AutomaticallyMergesChangesFromParent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}
func (m NSManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}

// The concurrency type for the context.
//
// # Discussion
//
// For more details on concurrency type, see [NSManagedObjectContext].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/concurrencyType-swift.property
func (m NSManagedObjectContext) ConcurrencyType() NSManagedObjectContextConcurrencyType {
	rv := objc.Send[NSManagedObjectContextConcurrencyType](m.ID, objc.Sel("concurrencyType"))
	return NSManagedObjectContextConcurrencyType(rv)
}

// The merge policy of the context.
//
// # Discussion
//
// The default is [NSErrorMergePolicy]. For possible values, see
// [NSMergePolicy].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergePolicy
//
// [NSErrorMergePolicy]: https://developer.apple.com/documentation/CoreData/NSErrorMergePolicy
func (m NSManagedObjectContext) MergePolicy() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mergePolicy"))
	return objectivec.Object{ID: rv}
}
func (m NSManagedObjectContext) SetMergePolicy(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setMergePolicy:"), value)
}

// Returns the token associated with the query generation currently in use by
// this context.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/queryGenerationToken
func (m NSManagedObjectContext) QueryGenerationToken() INSQueryGenerationToken {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("queryGenerationToken"))
	return NSQueryGenerationTokenFromID(objc.ID(rv))
}

// The author for the context that is used as an identifier in persistent
// history transactions.
//
// # Discussion
//
// Set a managed object context’s [NSManagedObjectContext.TransactionAuthor]
// before saving it to differentiate among multiple call sites that modify the
// same context. Doing this records an [NSPersistentHistoryTransaction.Author]
// in subsequent transactions.
//
// Reset the context’s [NSManagedObjectContext.TransactionAuthor] to nil
// after the save to prevent misattribution of future transactions.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/transactionAuthor
func (m NSManagedObjectContext) TransactionAuthor() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("transactionAuthor"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSManagedObjectContext) SetTransactionAuthor(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}

// A Boolean value that indicates whether the context has uncommitted changes.
//
// # Discussion
//
// If you are observing this property using key-value observing (KVO) you
// should not touch the context or its objects within your implementation of
// [observeValue(forKeyPath:of:change:context:)] for this notification. (This
// is because of the intricacy of the locations of the KVO notifications—for
// example, the context may be in the middle of an undo operation, or
// repairing a merge conflict.) If you need to send messages to the context or
// change any of its managed objects as a result of a change to the value of
// `hasChanges`, you must do so after the call stack unwinds (typically using
// [perform(_:with:afterDelay:)] or a similar method).
//
// # Special Considerations
//
// In macOS 10.6 and later, this property is [Key-value observing] compliant.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/hasChanges
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
// [observeValue(forKeyPath:of:change:context:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/observeValue(forKeyPath:of:change:context:)
// [perform(_:with:afterDelay:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:)
func (m NSManagedObjectContext) HasChanges() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasChanges"))
	return rv
}

// The object that provides undo support for the context.
//
// # Discussion
//
// Enable undo support for a context by setting this property to an instance
// of [UndoManager]. This can be an undo manager that’s exclusive to the
// context, or an existing undo manager if you want to integrate the
// context’s undo operations with those of the rest of your app.
//
// If your context uses an undo manager, you can realize a performance benefit
// by temporarily setting this property to `nil` when performing expensive
// operations on that context, such as importing a large number of objects.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undoManager
//
// [UndoManager]: https://developer.apple.com/documentation/Foundation/UndoManager
func (m NSManagedObjectContext) UndoManager() foundation.UndoManager {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("undoManager"))
	return foundation.UndoManagerFromID(objc.ID(rv))
}
func (m NSManagedObjectContext) SetUndoManager(value foundation.UndoManager) {
	objc.Send[struct{}](m.ID, objc.Sel("setUndoManager:"), value)
}

// A Boolean value that indicates whether the context propagates deletes at
// the end of the event in which a change was made.
//
// # Discussion
//
// true if the receiver propagates deletes at the end of the event in which a
// change was made, false if it propagates deletes only during a save
// operation. The default is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/propagatesDeletesAtEndOfEvent
func (m NSManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}
func (m NSManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPropagatesDeletesAtEndOfEvent:"), value)
}

// The maximum length of time that may have elapsed since the store previously
// fetched data before fulfilling a fault issues a new fetch.
//
// # Discussion
//
// The staleness interval controls whether fulfilling a fault uses data
// previously fetched by the application, or issues a new fetch (see also
// [NSManagedObjectContext.RefreshObjectMergeChanges]). The staleness interval
// does not affect objects currently in use (that is, it is not used to
// automatically update property values from a persistent store after a
// certain period of time).
//
// The expiration value is applied on a per object basis. It is the relative
// time until cached data (snapshots) should be considered stale. For example,
// a value of 300.0 informs the context to utilize cached information for no
// more than 5 minutes after an object was originally fetched.
//
// Note that the staleness interval is a hint and may not be supported by all
// persistent store types. It is not used by XML and binary stores, because
// these stores maintain all current values in memory.
//
// The default is a negative value, which represents infinite staleness
// allowed. `0.0` represents “no staleness acceptable”.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval
func (m NSManagedObjectContext) StalenessInterval() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](m.ID, objc.Sel("stalenessInterval"))
	return foundation.NSTimeInterval(rv)
}
func (m NSManagedObjectContext) SetStalenessInterval(value foundation.NSTimeInterval) {
	objc.Send[struct{}](m.ID, objc.Sel("setStalenessInterval:"), value)
}

// PerformBlockSync is a synchronous wrapper around [NSManagedObjectContext.PerformBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (m NSManagedObjectContext) PerformBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.PerformBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PerformBlockAndWaitSync is a synchronous wrapper around [NSManagedObjectContext.PerformBlockAndWait].
// It blocks until the completion handler fires or the context is cancelled.
func (m NSManagedObjectContext) PerformBlockAndWaitSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.PerformBlockAndWait(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
