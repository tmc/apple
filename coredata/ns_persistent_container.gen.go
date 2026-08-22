// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentContainer] class.
var (
	_NSPersistentContainerClass     NSPersistentContainerClass
	_NSPersistentContainerClassOnce sync.Once
)

func getNSPersistentContainerClass() NSPersistentContainerClass {
	_NSPersistentContainerClassOnce.Do(func() {
		_NSPersistentContainerClass = NSPersistentContainerClass{class: objc.GetClass("NSPersistentContainer")}
	})
	return _NSPersistentContainerClass
}

// GetNSPersistentContainerClass returns the class object for NSPersistentContainer.
func GetNSPersistentContainerClass() NSPersistentContainerClass {
	return getNSPersistentContainerClass()
}

type NSPersistentContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentContainerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentContainerClass) Alloc() NSPersistentContainer {
	rv := objc.Send[NSPersistentContainer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container that encapsulates the Core Data stack in your app.
//
// # Overview
//
// NSPersistentContainer simplifies the creation and management of the Core
// Data stack by handling the creation of the managed object model
// ([NSManagedObjectModel]), persistent store coordinator
// ([NSPersistentStoreCoordinator]), and the managed object context
// ([NSManagedObjectContext]).
//
// # Creating a Container
//
//   - [NSPersistentContainer.InitWithName]: Creates a container with the specified name.
//   - [NSPersistentContainer.InitWithNameManagedObjectModel]: Create a container with the specified name and managed object model.
//
// # Getting the Container’s Configuration
//
//   - [NSPersistentContainer.ManagedObjectModel]: The container’s managed object model.
//   - [NSPersistentContainer.Name]: The container’s name.
//   - [NSPersistentContainer.PersistentStoreCoordinator]: The container’s persistent store coordinator.
//
// # Managing Persistent Stores
//
//   - [NSPersistentContainer.PersistentStoreDescriptions]: The descriptions of the container’s persistent stores.
//   - [NSPersistentContainer.SetPersistentStoreDescriptions]
//   - [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler]: Loads the persistent stores.
//
// # Acquiring Contexts
//
//   - [NSPersistentContainer.NewBackgroundContext]: Returns a new managed object context that executes on a private queue.
//   - [NSPersistentContainer.ViewContext]: The main queue’s managed object context.
//
// # Performing Background Tasks
//
//   - [NSPersistentContainer.PerformBackgroundTask]: Executes a closure on a private queue using an ephemeral managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer
type NSPersistentContainer struct {
	objectivec.Object
}

// NSPersistentContainerFromID constructs a [NSPersistentContainer] from an objc.ID.
//
// A container that encapsulates the Core Data stack in your app.
func NSPersistentContainerFromID(id objc.ID) NSPersistentContainer {
	return NSPersistentContainer{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentContainer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentContainer] class.
//
// # Creating a Container
//
//   - [INSPersistentContainer.InitWithName]: Creates a container with the specified name.
//   - [INSPersistentContainer.InitWithNameManagedObjectModel]: Create a container with the specified name and managed object model.
//
// # Getting the Container’s Configuration
//
//   - [INSPersistentContainer.ManagedObjectModel]: The container’s managed object model.
//   - [INSPersistentContainer.Name]: The container’s name.
//   - [INSPersistentContainer.PersistentStoreCoordinator]: The container’s persistent store coordinator.
//
// # Managing Persistent Stores
//
//   - [INSPersistentContainer.PersistentStoreDescriptions]: The descriptions of the container’s persistent stores.
//   - [INSPersistentContainer.SetPersistentStoreDescriptions]
//   - [INSPersistentContainer.LoadPersistentStoresWithCompletionHandler]: Loads the persistent stores.
//
// # Acquiring Contexts
//
//   - [INSPersistentContainer.NewBackgroundContext]: Returns a new managed object context that executes on a private queue.
//   - [INSPersistentContainer.ViewContext]: The main queue’s managed object context.
//
// # Performing Background Tasks
//
//   - [INSPersistentContainer.PerformBackgroundTask]: Executes a closure on a private queue using an ephemeral managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer
type INSPersistentContainer interface {
	objectivec.IObject

	// Topic: Creating a Container

	// Creates a container with the specified name.
	InitWithName(name string) NSPersistentContainer
	// Create a container with the specified name and managed object model.
	InitWithNameManagedObjectModel(name string, model INSManagedObjectModel) NSPersistentContainer

	// Topic: Getting the Container’s Configuration

	// The container’s managed object model.
	ManagedObjectModel() INSManagedObjectModel
	// The container’s name.
	Name() string
	// The container’s persistent store coordinator.
	PersistentStoreCoordinator() INSPersistentStoreCoordinator

	// Topic: Managing Persistent Stores

	// The descriptions of the container’s persistent stores.
	PersistentStoreDescriptions() []NSPersistentStoreDescription
	SetPersistentStoreDescriptions(value []NSPersistentStoreDescription)
	// Loads the persistent stores.
	LoadPersistentStoresWithCompletionHandler(block PersistentStoreDescriptionErrorHandler)

	// Topic: Acquiring Contexts

	// Returns a new managed object context that executes on a private queue.
	NewBackgroundContext() INSManagedObjectContext
	// The main queue’s managed object context.
	ViewContext() INSManagedObjectContext

	// Topic: Performing Background Tasks

	// Executes a closure on a private queue using an ephemeral managed object context.
	PerformBackgroundTask(block ManagedObjectContextHandler)
}

// Init initializes the instance.
func (p NSPersistentContainer) Init() NSPersistentContainer {
	rv := objc.Send[NSPersistentContainer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentContainer) Autorelease() NSPersistentContainer {
	rv := objc.Send[NSPersistentContainer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentContainer creates a new NSPersistentContainer instance.
func NewNSPersistentContainer() NSPersistentContainer {
	class := getNSPersistentContainerClass()
	rv := objc.Send[NSPersistentContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a container with the specified name.
//
// name: The name of the [NSPersistentContainer] object.
//
// # Return Value
//
// A persistent container initialized with the given name.
//
// # Discussion
//
// By default, the provided name value is used to name the persistent store
// and is used to look up the name of the [NSManagedObjectModel] object to be
// used with the [NSPersistentContainer] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:)
func NewPersistentContainerWithName(name string) NSPersistentContainer {
	instance := getNSPersistentContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return NSPersistentContainerFromID(rv)
}

// Create a container with the specified name and managed object model.
//
// name: The name used by the persistent container.
//
// model: The managed object model to be used by the persistent container.
//
// # Return Value
//
// A persistent container initialized with the given name and model.
//
// # Discussion
//
// By default, the provided name value of the container is used as the name of
// the persisent store associated with the container. Passing in the
// [NSManagedObjectModel] object overrides the lookup of the model by the
// provided name value.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:managedObjectModel:)
func NewPersistentContainerWithNameManagedObjectModel(name string, model INSManagedObjectModel) NSPersistentContainer {
	instance := getNSPersistentContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:managedObjectModel:"), objc.String(name), model)
	return NSPersistentContainerFromID(rv)
}

// Creates a container with the specified name.
//
// name: The name of the [NSPersistentContainer] object.
//
// # Return Value
//
// A persistent container initialized with the given name.
//
// # Discussion
//
// By default, the provided name value is used to name the persistent store
// and is used to look up the name of the [NSManagedObjectModel] object to be
// used with the [NSPersistentContainer] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:)
func (p NSPersistentContainer) InitWithName(name string) NSPersistentContainer {
	rv := objc.Send[NSPersistentContainer](p.ID, objc.Sel("initWithName:"), objc.String(name))
	return rv
}

// Create a container with the specified name and managed object model.
//
// name: The name used by the persistent container.
//
// model: The managed object model to be used by the persistent container.
//
// # Return Value
//
// A persistent container initialized with the given name and model.
//
// # Discussion
//
// By default, the provided name value of the container is used as the name of
// the persisent store associated with the container. Passing in the
// [NSManagedObjectModel] object overrides the lookup of the model by the
// provided name value.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:managedObjectModel:)
func (p NSPersistentContainer) InitWithNameManagedObjectModel(name string, model INSManagedObjectModel) NSPersistentContainer {
	rv := objc.Send[NSPersistentContainer](p.ID, objc.Sel("initWithName:managedObjectModel:"), objc.String(name), model)
	return rv
}

// Loads the persistent stores.
//
// block: Once the loading of the persistent stores has completed, this block will be
// executed on the calling thread.
//
// # Discussion
//
// Once the persistent container has been initialized, you need to execute
// [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler] to
// instruct the container to load the persistent stores and complete the
// creation of the Core Data stack.
//
// Once the completion handler has fired, the stack is fully initialized and
// is ready for use. The completion handler will be called once for each
// persistent store that is created.
//
// If there is an error in the loading of the persistent stores, the [NSError]
// value will be populated.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/loadPersistentStores(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (p NSPersistentContainer) LoadPersistentStoresWithCompletionHandler(block PersistentStoreDescriptionErrorHandler) {
	_block0, _ := NewPersistentStoreDescriptionErrorBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("loadPersistentStoresWithCompletionHandler:"), _block0)
}

// Returns a new managed object context that executes on a private queue.
//
// # Return Value
//
// A newly created private managed object context.
//
// # Discussion
//
// Invoking this method causes the persistent container to create and return a
// new [NSManagedObjectContext] with the
// [NSManagedObjectContext.ConcurrencyType] set to
// [NSManagedObjectContextConcurrencyType.privateQueueConcurrencyType]. This
// new context will be associated with the [NSPersistentStoreCoordinator]
// directly and is set to consume [NSManagedObjectContextDidSave] broadcasts
// automatically.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/newBackgroundContext()
//
// [NSManagedObjectContextConcurrencyType.privateQueueConcurrencyType]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType/privateQueueConcurrencyType
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
func (p NSPersistentContainer) NewBackgroundContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("newBackgroundContext"))
	return NSManagedObjectContextFromID(rv)
}

// Executes a closure on a private queue using an ephemeral managed object
// context.
//
// block: A closure that is executed by the persistent container against a newly
// created private context. The private context is passed into the block as
// part of the execution of the block.
//
// # Discussion
//
// Each time this method is invoked, the persistent container creates a new
// [NSManagedObjectContext] with the [NSManagedObjectContext.ConcurrencyType]
// set to [NSManagedObjectContextConcurrencyType.privateQueueConcurrencyType].
// The persistent container then executes the passed in block against that
// newly created context on the context’s private queue.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/performBackgroundTask(_:)-39sch
//
// [NSManagedObjectContextConcurrencyType.privateQueueConcurrencyType]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType/privateQueueConcurrencyType
func (p NSPersistentContainer) PerformBackgroundTask(block ManagedObjectContextHandler) {
	_block0, _ := NewManagedObjectContextBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("performBackgroundTask:"), _block0)
}

// The container’s managed object model.
//
// # Discussion
//
// This property contains a reference to the [NSManagedObjectModel] object
// associated with this persistent container.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/managedObjectModel
func (p NSPersistentContainer) ManagedObjectModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The container’s name.
//
// # Discussion
//
// This property is passed in as part of the initialization of the persistent
// container. This name is used to locate the [NSManagedObjectModel] (if the
// [NSManagedObjectModel] object is not passed in as part of the
// initialization) and is used to name the persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/name
func (p NSPersistentContainer) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The container’s persistent store coordinator.
//
// # Discussion
//
// When the persistent container is initialized, it creates a persistent store
// coordinator as part of that initialization. That persistent store
// coordinator is referenced in this property.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentStoreCoordinator
func (p NSPersistentContainer) PersistentStoreCoordinator() INSPersistentStoreCoordinator {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("persistentStoreCoordinator"))
	return NSPersistentStoreCoordinatorFromID(objc.ID(rv))
}

// The descriptions of the container’s persistent stores.
//
// # Discussion
//
// If you want to override the type (or types) of persistent store(s) used by
// the persistent container, you can set this property with an array of
// [NSPersistentStoreDescription] objects.
//
// If you will be configuring custom persistent store descriptions, you must
// set this property before calling
// [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentStoreDescriptions
func (p NSPersistentContainer) PersistentStoreDescriptions() []NSPersistentStoreDescription {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("persistentStoreDescriptions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPersistentStoreDescription {
		return NSPersistentStoreDescriptionFromID(id)
	})
}
func (p NSPersistentContainer) SetPersistentStoreDescriptions(value []NSPersistentStoreDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setPersistentStoreDescriptions:"), objectivec.IObjectSliceToNSArray(value))
}

// The main queue’s managed object context.
//
// # Discussion
//
// This property contains a reference to the [NSManagedObjectContext] that is
// created and owned by the persistent container which is associated with the
// main queue of the application. This context is created automatically as
// part of the initialization of the persistent container.
//
// This context is associated directly with the [NSPersistentStoreCoordinator]
// and is non-generational by default.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/viewContext
func (p NSPersistentContainer) ViewContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("viewContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// The location of the directory that contains the persistent stores.
//
// See: https://developer.apple.com/documentation/coredata/nspersistentcontainer/defaultdirectoryurl-swift.type.property
func (_NSPersistentContainerClass NSPersistentContainerClass) DefaultDirectoryURL() foundation.NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentContainerClass.class), objc.Sel("defaultDirectoryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (_NSPersistentContainerClass NSPersistentContainerClass) SetDefaultDirectoryURL(value foundation.NSURL) {
	objc.Send[struct{}](objc.ID(_NSPersistentContainerClass.class), objc.Sel("setDefaultDirectoryURL:"), value)
}

// LoadPersistentStores is a synchronous wrapper around [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPersistentContainer) LoadPersistentStores(ctx context.Context) (*NSPersistentStoreDescription, error) {
	type result struct {
		val *NSPersistentStoreDescription
		err error
	}
	done := make(chan result, 1)
	p.LoadPersistentStoresWithCompletionHandler(func(val *NSPersistentStoreDescription, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PerformBackgroundTaskSync is a synchronous wrapper around [NSPersistentContainer.PerformBackgroundTask].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPersistentContainer) PerformBackgroundTaskSync(ctx context.Context) (*NSManagedObjectContext, error) {
	done := make(chan *NSManagedObjectContext, 1)
	p.PerformBackgroundTask(func(val *NSManagedObjectContext) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
