// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBatchInsertRequest] class.
var (
	_NSBatchInsertRequestClass     NSBatchInsertRequestClass
	_NSBatchInsertRequestClassOnce sync.Once
)

func getNSBatchInsertRequestClass() NSBatchInsertRequestClass {
	_NSBatchInsertRequestClassOnce.Do(func() {
		_NSBatchInsertRequestClass = NSBatchInsertRequestClass{class: objc.GetClass("NSBatchInsertRequest")}
	})
	return _NSBatchInsertRequestClass
}

// GetNSBatchInsertRequestClass returns the class object for NSBatchInsertRequest.
func GetNSBatchInsertRequestClass() NSBatchInsertRequestClass {
	return getNSBatchInsertRequestClass()
}

type NSBatchInsertRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchInsertRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchInsertRequestClass) Alloc() NSBatchInsertRequest {
	rv := objc.Send[NSBatchInsertRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A request to insert a batch of data in a persistent store.
//
// # Creating a Request
//
//   - [NSBatchInsertRequest.InitWithEntityManagedObjectHandler]: Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
//   - [NSBatchInsertRequest.InitWithEntityNameManagedObjectHandler]: Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
//   - [NSBatchInsertRequest.InitWithEntityObjects]: Creates a batch-insertion request for a managed entity, and provides an array of data dictionaries for insertion.
//   - [NSBatchInsertRequest.InitWithEntityNameObjects]: Creates a batch-insertion request for a named managed entity, and provides an array of data dictionaries for insertion.
//
// # Configuring a Request
//
//   - [NSBatchInsertRequest.DictionaryHandler]: A closure that provides a dictionary for your app to insert data into.
//   - [NSBatchInsertRequest.SetDictionaryHandler]
//   - [NSBatchInsertRequest.Entity]: The managed entity to insert data into.
//   - [NSBatchInsertRequest.EntityName]: The name of the managed entity to insert data into.
//   - [NSBatchInsertRequest.ManagedObjectHandler]: A closure that provides a managed object for your app to insert data into.
//   - [NSBatchInsertRequest.SetManagedObjectHandler]
//   - [NSBatchInsertRequest.ObjectsToInsert]: An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values.
//   - [NSBatchInsertRequest.SetObjectsToInsert]
//   - [NSBatchInsertRequest.ResultType]: The type of result that Core Data returns from this request.
//   - [NSBatchInsertRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest
type NSBatchInsertRequest struct {
	NSPersistentStoreRequest
}

// NSBatchInsertRequestFromID constructs a [NSBatchInsertRequest] from an objc.ID.
//
// A request to insert a batch of data in a persistent store.
func NSBatchInsertRequestFromID(id objc.ID) NSBatchInsertRequest {
	return NSBatchInsertRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSBatchInsertRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchInsertRequest] class.
//
// # Creating a Request
//
//   - [INSBatchInsertRequest.InitWithEntityManagedObjectHandler]: Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
//   - [INSBatchInsertRequest.InitWithEntityNameManagedObjectHandler]: Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
//   - [INSBatchInsertRequest.InitWithEntityObjects]: Creates a batch-insertion request for a managed entity, and provides an array of data dictionaries for insertion.
//   - [INSBatchInsertRequest.InitWithEntityNameObjects]: Creates a batch-insertion request for a named managed entity, and provides an array of data dictionaries for insertion.
//
// # Configuring a Request
//
//   - [INSBatchInsertRequest.DictionaryHandler]: A closure that provides a dictionary for your app to insert data into.
//   - [INSBatchInsertRequest.SetDictionaryHandler]
//   - [INSBatchInsertRequest.Entity]: The managed entity to insert data into.
//   - [INSBatchInsertRequest.EntityName]: The name of the managed entity to insert data into.
//   - [INSBatchInsertRequest.ManagedObjectHandler]: A closure that provides a managed object for your app to insert data into.
//   - [INSBatchInsertRequest.SetManagedObjectHandler]
//   - [INSBatchInsertRequest.ObjectsToInsert]: An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values.
//   - [INSBatchInsertRequest.SetObjectsToInsert]
//   - [INSBatchInsertRequest.ResultType]: The type of result that Core Data returns from this request.
//   - [INSBatchInsertRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest
type INSBatchInsertRequest interface {
	INSPersistentStoreRequest

	// Topic: Creating a Request

	// Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
	InitWithEntityManagedObjectHandler(entity INSEntityDescription, handler BoolManagedObjectHandler) NSBatchInsertRequest
	// Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
	InitWithEntityNameManagedObjectHandler(entityName string, handler BoolManagedObjectHandler) NSBatchInsertRequest
	// Creates a batch-insertion request for a managed entity, and provides an array of data dictionaries for insertion.
	InitWithEntityObjects(entity INSEntityDescription, dictionaries foundation.INSDictionary) NSBatchInsertRequest
	// Creates a batch-insertion request for a named managed entity, and provides an array of data dictionaries for insertion.
	InitWithEntityNameObjects(entityName string, dictionaries foundation.INSDictionary) NSBatchInsertRequest

	// Topic: Configuring a Request

	// A closure that provides a dictionary for your app to insert data into.
	DictionaryHandler() objectivec.IObject
	SetDictionaryHandler(value objc.ID)
	// The managed entity to insert data into.
	Entity() INSEntityDescription
	// The name of the managed entity to insert data into.
	EntityName() string
	// A closure that provides a managed object for your app to insert data into.
	ManagedObjectHandler() BoolManagedObjectHandler
	SetManagedObjectHandler(value BoolManagedObjectHandler)
	// An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values.
	ObjectsToInsert() foundation.INSDictionary
	SetObjectsToInsert(value foundation.INSDictionary)
	// The type of result that Core Data returns from this request.
	ResultType() NSBatchInsertRequestResultType
	SetResultType(value NSBatchInsertRequestResultType)
}

// Init initializes the instance.
func (b NSBatchInsertRequest) Init() NSBatchInsertRequest {
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchInsertRequest) Autorelease() NSBatchInsertRequest {
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchInsertRequest creates a new NSBatchInsertRequest instance.
func NewNSBatchInsertRequest() NSBatchInsertRequest {
	class := getNSBatchInsertRequestClass()
	rv := objc.Send[NSBatchInsertRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a batch-insertion request for a named managed entity, and provides
// an array of data dictionaries for insertion.
//
// entityName: The name of the managed entity to insert data into.
//
// dictionaries: An array of dictionaries that represents objects to insert. Each dictionary
// contains an attribute name key and a value.
//
// # Return Value
//
// A batch-insertion request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:objects:)
func NewBatchInsertRequestWithEntityNameObjects(entityName string, dictionaries foundation.INSDictionary) NSBatchInsertRequest {
	instance := getNSBatchInsertRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntityName:objects:"), objc.String(entityName), dictionaries)
	return NSBatchInsertRequestFromID(rv)
}

// Creates a batch-insertion request for a managed entity, and provides an
// array of data dictionaries for insertion.
//
// entity: The managed entity to insert data into.
//
// dictionaries: An array of dictionaries that represents objects to insert. Each dictionary
// contains an attribute name key and a value.
//
// # Return Value
//
// A batch-insertion request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entity:objects:)
func NewBatchInsertRequestWithEntityObjects(entity INSEntityDescription, dictionaries foundation.INSDictionary) NSBatchInsertRequest {
	instance := getNSBatchInsertRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntity:objects:"), entity, dictionaries)
	return NSBatchInsertRequestFromID(rv)
}

// Creates a batch-insertion request for a managed entity, and specifies a
// closure that inserts data into the entity.
//
// entity: A managed entity to insert data into.
//
// handler: A closure that inserts data into the managed entity.
//
// # Return Value
//
// A batch-insertion request.
//
// # Discussion
//
// Core Data repeatedly calls the provided closure until it returns `true`,
// then finishes the request and saves the data.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entity:managedObjectHandler:)
func (b NSBatchInsertRequest) InitWithEntityManagedObjectHandler(entity INSEntityDescription, handler BoolManagedObjectHandler) NSBatchInsertRequest {
	_block1, _ := NewBoolManagedObjectBlock(handler)
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("initWithEntity:managedObjectHandler:"), entity, _block1)
	return rv
}

// Creates a batch-insertion request for a named managed entity, and specifies
// a closure that inserts data into the entity.
//
// entityName: The name of the managed entity that defines the object to create.
//
// handler: A closure that inserts data into the managed entity.
//
// # Return Value
//
// A batch-insertion request.
//
// # Discussion
//
// Core Data repeatedly calls the provided closure until it returns `true`,
// then finishes the request and saves the data.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:managedObjectHandler:)
func (b NSBatchInsertRequest) InitWithEntityNameManagedObjectHandler(entityName string, handler BoolManagedObjectHandler) NSBatchInsertRequest {
	_block1, _ := NewBoolManagedObjectBlock(handler)
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("initWithEntityName:managedObjectHandler:"), objc.String(entityName), _block1)
	return rv
}

// Creates a batch-insertion request for a managed entity, and provides an
// array of data dictionaries for insertion.
//
// entity: The managed entity to insert data into.
//
// dictionaries: An array of dictionaries that represents objects to insert. Each dictionary
// contains an attribute name key and a value.
//
// # Return Value
//
// A batch-insertion request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entity:objects:)
func (b NSBatchInsertRequest) InitWithEntityObjects(entity INSEntityDescription, dictionaries foundation.INSDictionary) NSBatchInsertRequest {
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("initWithEntity:objects:"), entity, dictionaries)
	return rv
}

// Creates a batch-insertion request for a named managed entity, and provides
// an array of data dictionaries for insertion.
//
// entityName: The name of the managed entity to insert data into.
//
// dictionaries: An array of dictionaries that represents objects to insert. Each dictionary
// contains an attribute name key and a value.
//
// # Return Value
//
// A batch-insertion request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:objects:)
func (b NSBatchInsertRequest) InitWithEntityNameObjects(entityName string, dictionaries foundation.INSDictionary) NSBatchInsertRequest {
	rv := objc.Send[NSBatchInsertRequest](b.ID, objc.Sel("initWithEntityName:objects:"), objc.String(entityName), dictionaries)
	return rv
}

// A closure that provides a dictionary for your app to insert data into.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/dictionaryHandler
func (b NSBatchInsertRequest) DictionaryHandler() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("dictionaryHandler"))
	return objectivec.Object{ID: rv}
}
func (b NSBatchInsertRequest) SetDictionaryHandler(value objc.ID) {
	objc.Send[struct{}](b.ID, objc.Sel("setDictionaryHandler:"), value)
}

// The managed entity to insert data into.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/entity
func (b NSBatchInsertRequest) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// The name of the managed entity to insert data into.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/entityName
func (b NSBatchInsertRequest) EntityName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("entityName"))
	return foundation.NSStringFromID(rv).String()
}

// A closure that provides a managed object for your app to insert data into.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/managedObjectHandler
func (b NSBatchInsertRequest) ManagedObjectHandler() BoolManagedObjectHandler {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("managedObjectHandler"))
	_ = rv
	return nil
}
func (b NSBatchInsertRequest) SetManagedObjectHandler(value BoolManagedObjectHandler) {
	block, cleanup := NewBoolManagedObjectBlock(value)
	defer cleanup()
	objc.Send[struct{}](b.ID, objc.Sel("setManagedObjectHandler:"), block)
}

// An array of dictionaries that represents the objects to insert with the
// keys as attribute names and their assigned values.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/objectsToInsert
func (b NSBatchInsertRequest) ObjectsToInsert() foundation.INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("objectsToInsert"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b NSBatchInsertRequest) SetObjectsToInsert(value foundation.INSDictionary) {
	objc.Send[struct{}](b.ID, objc.Sel("setObjectsToInsert:"), value)
}

// The type of result that Core Data returns from this request.
//
// # Discussion
//
// The default is [NSBatchInsertRequestResultType.statusOnly].
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/resultType
//
// [NSBatchInsertRequestResultType.statusOnly]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType/statusOnly
func (b NSBatchInsertRequest) ResultType() NSBatchInsertRequestResultType {
	rv := objc.Send[NSBatchInsertRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchInsertRequestResultType(rv)
}
func (b NSBatchInsertRequest) SetResultType(value NSBatchInsertRequestResultType) {
	objc.Send[struct{}](b.ID, objc.Sel("setResultType:"), value)
}
