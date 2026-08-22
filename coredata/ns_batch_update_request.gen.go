// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSBatchUpdateRequest] class.
var (
	_NSBatchUpdateRequestClass     NSBatchUpdateRequestClass
	_NSBatchUpdateRequestClassOnce sync.Once
)

func getNSBatchUpdateRequestClass() NSBatchUpdateRequestClass {
	_NSBatchUpdateRequestClassOnce.Do(func() {
		_NSBatchUpdateRequestClass = NSBatchUpdateRequestClass{class: objc.GetClass("NSBatchUpdateRequest")}
	})
	return _NSBatchUpdateRequestClass
}

// GetNSBatchUpdateRequestClass returns the class object for NSBatchUpdateRequest.
func GetNSBatchUpdateRequestClass() NSBatchUpdateRequestClass {
	return getNSBatchUpdateRequestClass()
}

type NSBatchUpdateRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchUpdateRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchUpdateRequestClass) Alloc() NSBatchUpdateRequest {
	rv := objc.Send[NSBatchUpdateRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A request to Core Data to do a batch update of data in a persistent store
// without loading any data into memory.
//
// # Creating a Request
//
//   - [NSBatchUpdateRequest.InitWithEntity]: Creates a batch-update request for a managed entity.
//   - [NSBatchUpdateRequest.InitWithEntityName]: Creates a batch-update request for a named managed entity.
//
// # Configuring a Request
//
//   - [NSBatchUpdateRequest.Entity]: The managed entity to update data for.
//   - [NSBatchUpdateRequest.EntityName]: The name of the managed entity to update data for.
//   - [NSBatchUpdateRequest.IncludesSubentities]: A Boolean value that indicates whether to update subentities.
//   - [NSBatchUpdateRequest.SetIncludesSubentities]
//   - [NSBatchUpdateRequest.Predicate]: A predicate that identifies the objects to update.
//   - [NSBatchUpdateRequest.SetPredicate]
//   - [NSBatchUpdateRequest.PropertiesToUpdate]: A dictionary of property description pairs that describe the updates.
//   - [NSBatchUpdateRequest.SetPropertiesToUpdate]
//   - [NSBatchUpdateRequest.ResultType]: The type of result that Core Data returns from the request.
//   - [NSBatchUpdateRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest
type NSBatchUpdateRequest struct {
	NSPersistentStoreRequest
}

// NSBatchUpdateRequestFromID constructs a [NSBatchUpdateRequest] from an objc.ID.
//
// A request to Core Data to do a batch update of data in a persistent store
// without loading any data into memory.
func NSBatchUpdateRequestFromID(id objc.ID) NSBatchUpdateRequest {
	return NSBatchUpdateRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSBatchUpdateRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchUpdateRequest] class.
//
// # Creating a Request
//
//   - [INSBatchUpdateRequest.InitWithEntity]: Creates a batch-update request for a managed entity.
//   - [INSBatchUpdateRequest.InitWithEntityName]: Creates a batch-update request for a named managed entity.
//
// # Configuring a Request
//
//   - [INSBatchUpdateRequest.Entity]: The managed entity to update data for.
//   - [INSBatchUpdateRequest.EntityName]: The name of the managed entity to update data for.
//   - [INSBatchUpdateRequest.IncludesSubentities]: A Boolean value that indicates whether to update subentities.
//   - [INSBatchUpdateRequest.SetIncludesSubentities]
//   - [INSBatchUpdateRequest.Predicate]: A predicate that identifies the objects to update.
//   - [INSBatchUpdateRequest.SetPredicate]
//   - [INSBatchUpdateRequest.PropertiesToUpdate]: A dictionary of property description pairs that describe the updates.
//   - [INSBatchUpdateRequest.SetPropertiesToUpdate]
//   - [INSBatchUpdateRequest.ResultType]: The type of result that Core Data returns from the request.
//   - [INSBatchUpdateRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest
type INSBatchUpdateRequest interface {
	INSPersistentStoreRequest

	// Topic: Creating a Request

	// Creates a batch-update request for a managed entity.
	InitWithEntity(entity INSEntityDescription) NSBatchUpdateRequest
	// Creates a batch-update request for a named managed entity.
	InitWithEntityName(entityName string) NSBatchUpdateRequest

	// Topic: Configuring a Request

	// The managed entity to update data for.
	Entity() INSEntityDescription
	// The name of the managed entity to update data for.
	EntityName() string
	// A Boolean value that indicates whether to update subentities.
	IncludesSubentities() bool
	SetIncludesSubentities(value bool)
	// A predicate that identifies the objects to update.
	Predicate() foundation.NSPredicate
	SetPredicate(value foundation.NSPredicate)
	// A dictionary of property description pairs that describe the updates.
	PropertiesToUpdate() foundation.INSDictionary
	SetPropertiesToUpdate(value foundation.INSDictionary)
	// The type of result that Core Data returns from the request.
	ResultType() NSBatchUpdateRequestResultType
	SetResultType(value NSBatchUpdateRequestResultType)
}

// Init initializes the instance.
func (b NSBatchUpdateRequest) Init() NSBatchUpdateRequest {
	rv := objc.Send[NSBatchUpdateRequest](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchUpdateRequest) Autorelease() NSBatchUpdateRequest {
	rv := objc.Send[NSBatchUpdateRequest](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchUpdateRequest creates a new NSBatchUpdateRequest instance.
func NewNSBatchUpdateRequest() NSBatchUpdateRequest {
	class := getNSBatchUpdateRequestClass()
	rv := objc.Send[NSBatchUpdateRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a batch-update request for a managed entity.
//
// entity: The managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entity:)
func NewBatchUpdateRequestWithEntity(entity INSEntityDescription) NSBatchUpdateRequest {
	instance := getNSBatchUpdateRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntity:"), entity)
	return NSBatchUpdateRequestFromID(rv)
}

// Creates a batch-update request for a named managed entity.
//
// entityName: The name of the managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entityName:)
func NewBatchUpdateRequestWithEntityName(entityName string) NSBatchUpdateRequest {
	instance := getNSBatchUpdateRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	return NSBatchUpdateRequestFromID(rv)
}

// Creates a batch-update request for a managed entity.
//
// entity: The managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entity:)
func (b NSBatchUpdateRequest) InitWithEntity(entity INSEntityDescription) NSBatchUpdateRequest {
	rv := objc.Send[NSBatchUpdateRequest](b.ID, objc.Sel("initWithEntity:"), entity)
	return rv
}

// Creates a batch-update request for a named managed entity.
//
// entityName: The name of the managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entityName:)
func (b NSBatchUpdateRequest) InitWithEntityName(entityName string) NSBatchUpdateRequest {
	rv := objc.Send[NSBatchUpdateRequest](b.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	return rv
}

// The managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/entity
func (b NSBatchUpdateRequest) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// The name of the managed entity to update data for.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/entityName
func (b NSBatchUpdateRequest) EntityName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("entityName"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether to update subentities.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/includesSubentities
func (b NSBatchUpdateRequest) IncludesSubentities() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("includesSubentities"))
	return rv
}
func (b NSBatchUpdateRequest) SetIncludesSubentities(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setIncludesSubentities:"), value)
}

// A predicate that identifies the objects to update.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/predicate
func (b NSBatchUpdateRequest) Predicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("predicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (b NSBatchUpdateRequest) SetPredicate(value foundation.NSPredicate) {
	objc.Send[struct{}](b.ID, objc.Sel("setPredicate:"), value)
}

// A dictionary of property description pairs that describe the updates.
//
// # Discussion
//
// The dictionary keys are either [NSPropertyDescription] objects or strings
// that identify the property name.
//
// The dictionary values are either a constant value or an [NSExpression] that
// evaluates to a scalar value.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/propertiesToUpdate
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
func (b NSBatchUpdateRequest) PropertiesToUpdate() foundation.INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("propertiesToUpdate"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b NSBatchUpdateRequest) SetPropertiesToUpdate(value foundation.INSDictionary) {
	objc.Send[struct{}](b.ID, objc.Sel("setPropertiesToUpdate:"), value)
}

// The type of result that Core Data returns from the request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/resultType
func (b NSBatchUpdateRequest) ResultType() NSBatchUpdateRequestResultType {
	rv := objc.Send[NSBatchUpdateRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchUpdateRequestResultType(rv)
}
func (b NSBatchUpdateRequest) SetResultType(value NSBatchUpdateRequestResultType) {
	objc.Send[struct{}](b.ID, objc.Sel("setResultType:"), value)
}
