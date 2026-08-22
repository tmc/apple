// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFetchRequest] class.
var (
	_NSFetchRequestClass     NSFetchRequestClass
	_NSFetchRequestClassOnce sync.Once
)

func getNSFetchRequestClass() NSFetchRequestClass {
	_NSFetchRequestClassOnce.Do(func() {
		_NSFetchRequestClass = NSFetchRequestClass{class: objc.GetClass("NSFetchRequest")}
	})
	return _NSFetchRequestClass
}

// GetNSFetchRequestClass returns the class object for NSFetchRequest.
func GetNSFetchRequestClass() NSFetchRequestClass {
	return getNSFetchRequestClass()
}

type NSFetchRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchRequestClass) Alloc() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of search criteria used to retrieve data from a persistent
// store.
//
// # Overview
//
// An instance of [NSFetchRequest] collects the criteria needed to select and
// optionally to sort a group of [NSManagedObject] managed objects held in an
// [NSPersistentStore] persistent store. A fetch request contains an
// [NSEntityDescription] or an entity name that specifies which entity to
// search. It frequently also contains:
//
// - An [NSPredicate] predicate that specifies which properties to filter by
// and the constraints on selection, such as, `“last name begins with a
// ‘J’”`. If you don’t specify a predicate, then the system fetches
// all instances of the entity that you specified, subject to other
// constraints. For more information, see [fetch(_:)]. - An array of
// [NSSortDescriptor] sort descriptors that specify how to order the returned
// objects, such as ascending by last name and then by first name.
//
// You can also specify other aspects of a fetch request:
//
// [NSFetchRequest.FetchLimit]: The maximum number of objects that a request
// returns [NSFetchRequest.FetchOffset]: The number of objects to skip
// [NSFetchRequest.AffectedStores]: Which data stores the request accesses
// [NSFetchRequest.ResultType]: Whether the fetch returns managed objects,
// object IDs, dictionaries, or a count
// [NSFetchRequest.IncludesPropertyValues] and: Whether objects are fully
// populated with their properties [NSFetchRequest.ReturnsObjectsAsFaults]:
// Whether the objects are faults [NSFetchRequest.IncludesSubentities]:
// Whether the fetch includes subentities of the fetched entity
// [NSFetchRequest.PropertiesToFetch]: Which properties to fetch
// [NSFetchRequest.IncludesPendingChanges]: Whether to include unsaved changes
//
// Use [NSFetchRequest.Execute] to perform the fetch directly on the managed
// object context that’s associated with the current queue. Or use one of
// the [NSManagedObjectContext] methods such as
// [NSManagedObjectContext.PerformBlock] to execute the fetch.
//
// In [SwiftUI], you can use a [FetchRequest] property wrapper to execute the
// fetch and assign the results to a property. First, create the request:
//
// Then use a [FetchRequest] property wrapper with the request to declare a
// property that receives the objects that the fetch returns:
//
// You often predefine fetch requests in an [NSManagedObjectModel] managed
// object model to provide an API to retrieve a stored fetch request by name.
// Stored fetch requests can include placeholders for variable substitution,
// and serve as templates for later completion. Fetch request templates allow
// you to predefine queries with variables to substitute at runtime.
//
// # Managing the Fetch Request’s Entity
//
//   - [NSFetchRequest.InitWithEntityName]: Initializes a fetch request configured with a given entity name.
//   - [NSFetchRequest.EntityName]: The name of the entity the request is configured to fetch.
//   - [NSFetchRequest.Entity]: The entity specified for the fetch request.
//   - [NSFetchRequest.SetEntity]
//   - [NSFetchRequest.IncludesSubentities]: A Boolean value that indicates whether the fetch request includes subentities in the results.
//   - [NSFetchRequest.SetIncludesSubentities]
//
// # Specifying Fetch Constraints
//
//   - [NSFetchRequest.Predicate]: The predicate of the fetch request.
//   - [NSFetchRequest.SetPredicate]
//   - [NSFetchRequest.FetchLimit]: The fetch limit of the fetch request.
//   - [NSFetchRequest.SetFetchLimit]
//   - [NSFetchRequest.FetchOffset]: The fetch offset of the fetch request.
//   - [NSFetchRequest.SetFetchOffset]
//   - [NSFetchRequest.FetchBatchSize]: The batch size of the objects specified in the fetch request.
//   - [NSFetchRequest.SetFetchBatchSize]
//
// # Sorting the Results
//
//   - [NSFetchRequest.SortDescriptors]: The sort descriptors of the fetch request.
//   - [NSFetchRequest.SetSortDescriptors]
//
// # Prefetching Related Objects
//
//   - [NSFetchRequest.RelationshipKeyPathsForPrefetching]: The relationship key paths to prefetch along with the entity for the request.
//   - [NSFetchRequest.SetRelationshipKeyPathsForPrefetching]
//
// # Managing How Results Are Returned
//
//   - [NSFetchRequest.ResultType]: The result type of the fetch request.
//   - [NSFetchRequest.SetResultType]
//   - [NSFetchRequest.IncludesPendingChanges]: A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context.
//   - [NSFetchRequest.SetIncludesPendingChanges]
//   - [NSFetchRequest.PropertiesToFetch]: A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
//   - [NSFetchRequest.SetPropertiesToFetch]
//   - [NSFetchRequest.ReturnsDistinctResults]: A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by [propertiesToFetch](<https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch>).
//   - [NSFetchRequest.SetReturnsDistinctResults]
//   - [NSFetchRequest.IncludesPropertyValues]: A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store.
//   - [NSFetchRequest.SetIncludesPropertyValues]
//   - [NSFetchRequest.ShouldRefreshRefetchedObjects]: A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store.
//   - [NSFetchRequest.SetShouldRefreshRefetchedObjects]
//   - [NSFetchRequest.ReturnsObjectsAsFaults]: A Boolean value that indicates whether the objects resulting from a fetch request are faults.
//   - [NSFetchRequest.SetReturnsObjectsAsFaults]
//
// # Grouping and Filtering Dictionary Results
//
//   - [NSFetchRequest.PropertiesToGroupBy]: An array of objects that indicates how data should be grouped before a select statement is run in a SQL database.
//   - [NSFetchRequest.SetPropertiesToGroupBy]
//   - [NSFetchRequest.HavingPredicate]: The predicate used to filter rows being returned by a query containing a GROUP BY directive.
//   - [NSFetchRequest.SetHavingPredicate]
//
// # Executing a Fetch Request Directly
//
//   - [NSFetchRequest.Execute]: Executes the fetch request against the managed object context that is associated with the current queue.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest
//
// [FetchRequest]: https://developer.apple.com/documentation/SwiftUI/FetchRequest
// [NSManagedObject]: https://developer.apple.com/library/archive/releasenotes/Cocoa/CoreDataReleaseNotes/index.html#//apple_ref/doc/uid/TP40006503-SW6
// [NSPredicate]: https://developer.apple.com/documentation/Foundation/NSPredicate
// [NSSortDescriptor]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor
// [SwiftUI]: https://developer.apple.com/documentation/SwiftUI
// [fetch(_:)]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/fetch(_:)-38ys1
type NSFetchRequest struct {
	NSPersistentStoreRequest
}

// NSFetchRequestFromID constructs a [NSFetchRequest] from an objc.ID.
//
// A description of search criteria used to retrieve data from a persistent
// store.
func NSFetchRequestFromID(id objc.ID) NSFetchRequest {
	return NSFetchRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSFetchRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchRequest] class.
//
// # Managing the Fetch Request’s Entity
//
//   - [INSFetchRequest.InitWithEntityName]: Initializes a fetch request configured with a given entity name.
//   - [INSFetchRequest.EntityName]: The name of the entity the request is configured to fetch.
//   - [INSFetchRequest.Entity]: The entity specified for the fetch request.
//   - [INSFetchRequest.SetEntity]
//   - [INSFetchRequest.IncludesSubentities]: A Boolean value that indicates whether the fetch request includes subentities in the results.
//   - [INSFetchRequest.SetIncludesSubentities]
//
// # Specifying Fetch Constraints
//
//   - [INSFetchRequest.Predicate]: The predicate of the fetch request.
//   - [INSFetchRequest.SetPredicate]
//   - [INSFetchRequest.FetchLimit]: The fetch limit of the fetch request.
//   - [INSFetchRequest.SetFetchLimit]
//   - [INSFetchRequest.FetchOffset]: The fetch offset of the fetch request.
//   - [INSFetchRequest.SetFetchOffset]
//   - [INSFetchRequest.FetchBatchSize]: The batch size of the objects specified in the fetch request.
//   - [INSFetchRequest.SetFetchBatchSize]
//
// # Sorting the Results
//
//   - [INSFetchRequest.SortDescriptors]: The sort descriptors of the fetch request.
//   - [INSFetchRequest.SetSortDescriptors]
//
// # Prefetching Related Objects
//
//   - [INSFetchRequest.RelationshipKeyPathsForPrefetching]: The relationship key paths to prefetch along with the entity for the request.
//   - [INSFetchRequest.SetRelationshipKeyPathsForPrefetching]
//
// # Managing How Results Are Returned
//
//   - [INSFetchRequest.ResultType]: The result type of the fetch request.
//   - [INSFetchRequest.SetResultType]
//   - [INSFetchRequest.IncludesPendingChanges]: A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context.
//   - [INSFetchRequest.SetIncludesPendingChanges]
//   - [INSFetchRequest.PropertiesToFetch]: A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
//   - [INSFetchRequest.SetPropertiesToFetch]
//   - [INSFetchRequest.ReturnsDistinctResults]: A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by [propertiesToFetch](<https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch>).
//   - [INSFetchRequest.SetReturnsDistinctResults]
//   - [INSFetchRequest.IncludesPropertyValues]: A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store.
//   - [INSFetchRequest.SetIncludesPropertyValues]
//   - [INSFetchRequest.ShouldRefreshRefetchedObjects]: A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store.
//   - [INSFetchRequest.SetShouldRefreshRefetchedObjects]
//   - [INSFetchRequest.ReturnsObjectsAsFaults]: A Boolean value that indicates whether the objects resulting from a fetch request are faults.
//   - [INSFetchRequest.SetReturnsObjectsAsFaults]
//
// # Grouping and Filtering Dictionary Results
//
//   - [INSFetchRequest.PropertiesToGroupBy]: An array of objects that indicates how data should be grouped before a select statement is run in a SQL database.
//   - [INSFetchRequest.SetPropertiesToGroupBy]
//   - [INSFetchRequest.HavingPredicate]: The predicate used to filter rows being returned by a query containing a GROUP BY directive.
//   - [INSFetchRequest.SetHavingPredicate]
//
// # Executing a Fetch Request Directly
//
//   - [INSFetchRequest.Execute]: Executes the fetch request against the managed object context that is associated with the current queue.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest
type INSFetchRequest interface {
	INSPersistentStoreRequest

	// Topic: Managing the Fetch Request’s Entity

	// Initializes a fetch request configured with a given entity name.
	InitWithEntityName(entityName string) NSFetchRequest
	// The name of the entity the request is configured to fetch.
	EntityName() string
	// The entity specified for the fetch request.
	Entity() INSEntityDescription
	SetEntity(value INSEntityDescription)
	// A Boolean value that indicates whether the fetch request includes subentities in the results.
	IncludesSubentities() bool
	SetIncludesSubentities(value bool)

	// Topic: Specifying Fetch Constraints

	// The predicate of the fetch request.
	Predicate() foundation.NSPredicate
	SetPredicate(value foundation.NSPredicate)
	// The fetch limit of the fetch request.
	FetchLimit() uint
	SetFetchLimit(value uint)
	// The fetch offset of the fetch request.
	FetchOffset() uint
	SetFetchOffset(value uint)
	// The batch size of the objects specified in the fetch request.
	FetchBatchSize() uint
	SetFetchBatchSize(value uint)

	// Topic: Sorting the Results

	// The sort descriptors of the fetch request.
	SortDescriptors() []foundation.NSSortDescriptor
	SetSortDescriptors(value []foundation.NSSortDescriptor)

	// Topic: Prefetching Related Objects

	// The relationship key paths to prefetch along with the entity for the request.
	RelationshipKeyPathsForPrefetching() []string
	SetRelationshipKeyPathsForPrefetching(value []string)

	// Topic: Managing How Results Are Returned

	// The result type of the fetch request.
	ResultType() NSFetchRequestResultType
	SetResultType(value NSFetchRequestResultType)
	// A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context.
	IncludesPendingChanges() bool
	SetIncludesPendingChanges(value bool)
	// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
	PropertiesToFetch() foundation.INSArray
	SetPropertiesToFetch(value foundation.INSArray)
	// A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by [propertiesToFetch](<https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch>).
	ReturnsDistinctResults() bool
	SetReturnsDistinctResults(value bool)
	// A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store.
	IncludesPropertyValues() bool
	SetIncludesPropertyValues(value bool)
	// A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store.
	ShouldRefreshRefetchedObjects() bool
	SetShouldRefreshRefetchedObjects(value bool)
	// A Boolean value that indicates whether the objects resulting from a fetch request are faults.
	ReturnsObjectsAsFaults() bool
	SetReturnsObjectsAsFaults(value bool)

	// Topic: Grouping and Filtering Dictionary Results

	// An array of objects that indicates how data should be grouped before a select statement is run in a SQL database.
	PropertiesToGroupBy() foundation.INSArray
	SetPropertiesToGroupBy(value foundation.INSArray)
	// The predicate used to filter rows being returned by a query containing a GROUP BY directive.
	HavingPredicate() foundation.NSPredicate
	SetHavingPredicate(value foundation.NSPredicate)

	// Topic: Executing a Fetch Request Directly

	// Executes the fetch request against the managed object context that is associated with the current queue.
	Execute() ([]objectivec.IObject, error)
}

// Init initializes the instance.
func (f NSFetchRequest) Init() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchRequest) Autorelease() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchRequest creates a new NSFetchRequest instance.
func NewNSFetchRequest() NSFetchRequest {
	class := getNSFetchRequestClass()
	rv := objc.Send[NSFetchRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a fetch request configured with a given entity name.
//
// entityName: The name of the entity to fetch.
//
// # Return Value
//
// A fetch request configured to fetch using the entity named `entityName`.
//
// # Discussion
//
// This method provides a convenient way to configure the entity for a fetch
// request without having to retrieve an [NSEntityDescription] object. When
// the fetch is executed, the request uses the managed object context to find
// the entity with the given name. The model associated with the context’s
// persistent store coordinator must contain an entity named `entityName`.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/init(entityName:)
func NewFetchRequestWithEntityName(entityName string) NSFetchRequest {
	instance := getNSFetchRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	return NSFetchRequestFromID(rv)
}

// Initializes a fetch request configured with a given entity name.
//
// entityName: The name of the entity to fetch.
//
// # Return Value
//
// A fetch request configured to fetch using the entity named `entityName`.
//
// # Discussion
//
// This method provides a convenient way to configure the entity for a fetch
// request without having to retrieve an [NSEntityDescription] object. When
// the fetch is executed, the request uses the managed object context to find
// the entity with the given name. The model associated with the context’s
// persistent store coordinator must contain an entity named `entityName`.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/init(entityName:)
func (f NSFetchRequest) InitWithEntityName(entityName string) NSFetchRequest {
	rv := objc.Send[NSFetchRequest](f.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	return rv
}

// Executes the fetch request against the managed object context that is
// associated with the current queue.
//
// # Discussion
//
// Calling `execute` on an [NSFetchRequest] will cause the [NSFetchRequest] to
// run against the managed object context ([NSManagedObjectContext]) that is
// associated with the queue on which the `execute` is called.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/execute()
func (f NSFetchRequest) Execute() ([]objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("execute:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	}), nil

}

// The name of the entity the request is configured to fetch.
//
// # Discussion
//
// The entity name property is populated whenever the NSFetchRequest is
// created with [NSFetchRequest.InitWithEntityName] or
// [fetchRequestWithEntityName:].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/entityName
//
// [fetchRequestWithEntityName:]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchRequestWithEntityName:
func (f NSFetchRequest) EntityName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("entityName"))
	return foundation.NSStringFromID(rv).String()
}

// The entity specified for the fetch request.
//
// # Discussion
//
// When an [NSFetchRequest] instance is created with `init()`, it is expected
// that the [NSPropertyDescription.Entity] property will be set. If this
// property is not set, the fetch request fails upon execution.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/entity
func (f NSFetchRequest) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}
func (f NSFetchRequest) SetEntity(value INSEntityDescription) {
	objc.Send[struct{}](f.ID, objc.Sel("setEntity:"), value)
}

// A Boolean value that indicates whether the fetch request includes
// subentities in the results.
//
// # Discussion
//
// The value is true if the request will include all subentities of the entity
// for the request; otherwise it is false. The default is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesSubentities
func (f NSFetchRequest) IncludesSubentities() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("includesSubentities"))
	return rv
}
func (f NSFetchRequest) SetIncludesSubentities(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setIncludesSubentities:"), value)
}

// The predicate of the fetch request.
//
// # Discussion
//
// The predicate instance constrains the selection of objects the
// [NSFetchRequest] instance is to fetch.
//
// If the predicate is empty—for example, if it is an [AND] predicate whose
// array of elements contains no predicates—the request has its predicate
// set to `nil`. For more about predicates, see [Predicate Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/predicate
//
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
func (f NSFetchRequest) Predicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("predicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (f NSFetchRequest) SetPredicate(value foundation.NSPredicate) {
	objc.Send[struct{}](f.ID, objc.Sel("setPredicate:"), value)
}

// The fetch limit of the fetch request.
//
// # Discussion
//
// The fetch limit specifies the maximum number of objects that a request
// should return when executed.
//
// If you set a fetch limit, the framework makes a best effort to improve
// efficiency, but does not guarantee it. For every object store except the
// SQL store, a fetch request executed with a fetch limit in effect simply
// performs an unlimited fetch and throws away the unasked for rows.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchLimit
func (f NSFetchRequest) FetchLimit() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("fetchLimit"))
	return rv
}
func (f NSFetchRequest) SetFetchLimit(value uint) {
	objc.Send[struct{}](f.ID, objc.Sel("setFetchLimit:"), value)
}

// The fetch offset of the fetch request.
//
// # Discussion
//
// The default value is `0`.
//
// This setting allows you to specify an offset at which rows will begin being
// returned. Effectively, the request skips the specified number of matching
// entries. For example, given a fetch that typically returns `a, b, c, d`,
// specifying an offset of 1 will return `b, c, d`, and an offset of 4 will
// return an empty array. Offsets are ignored in nested requests such as
// subqueries.
//
// This property can be used to restrict the working set of data. In
// combination with [NSFetchRequest.FetchLimit], you can create a subrange of
// an arbitrary result set.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchOffset
func (f NSFetchRequest) FetchOffset() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("fetchOffset"))
	return rv
}
func (f NSFetchRequest) SetFetchOffset(value uint) {
	objc.Send[struct{}](f.ID, objc.Sel("setFetchOffset:"), value)
}

// The batch size of the objects specified in the fetch request.
//
// # Discussion
//
// The default value is `0`. A batch size of `0` is treated as infinite, which
// disables the batch fetching behavior.
//
// If you set a nonzero batch size, the collection of objects returned when an
// instance of [NSFetchRequest] is executed is broken into batches. When the
// fetch is executed, the entire request is evaluated and the identities of
// all matching objects recorded, but only data for objects up to the
// `batchSize` will be fetched from the persistent store at a time. The array
// returned from executing the request is a proxy object that transparently
// fetches subsequent batches on demand. (In database terms, this is an
// in-memory cursor.)
//
// You can use this feature to restrict the working set of data in your
// application. In combination with [NSFetchRequest.FetchLimit], you can
// create a subrange of an arbitrary result set.
//
// # Special Considerations
//
// For purposes of thread safety, when the fetch is executed, consider the
// array proxy returned as being owned by the managed object context the
// request is executed against. Treat the array proxy as if it were a managed
// object registered with that context.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchBatchSize
func (f NSFetchRequest) FetchBatchSize() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("fetchBatchSize"))
	return rv
}
func (f NSFetchRequest) SetFetchBatchSize(value uint) {
	objc.Send[struct{}](f.ID, objc.Sel("setFetchBatchSize:"), value)
}

// The sort descriptors of the fetch request.
//
// # Discussion
//
// The sort descriptors specify how the objects returned when the
// [NSFetchRequest] is issued should be ordered—for example, by last name
// and then by first name. The sort descriptors are applied in the order in
// which they appear in the `sortDescriptors` array (serially in
// lowest-array-index-first order).
//
// A value of `nil` is treated as no sort descriptors.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (f NSFetchRequest) SortDescriptors() []foundation.NSSortDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("sortDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSSortDescriptor {
		return foundation.NSSortDescriptorFromID(id)
	})
}
func (f NSFetchRequest) SetSortDescriptors(value []foundation.NSSortDescriptor) {
	objc.Send[struct{}](f.ID, objc.Sel("setSortDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}

// The relationship key paths to prefetch along with the entity for the
// request.
//
// # Discussion
//
// This property is an array of relationship key-path strings in
// NSKeyValueCoding notation (as you typically use with [value(forKeyPath:)]).
// The default value is an empty array (no prefetching).
//
// Prefetching allows Core Data to obtain related objects in a single fetch
// (per entity), rather than incurring subsequent access to the store for each
// individual record as their faults are tripped. For example, given an
// Employee entity with a relationship to a Department entity, suppose you
// fetch all the employees, and then for each print out their name and the
// name of the department to which they belong. In this case, a fault might be
// fired for each individual Department object. This can represent a
// significant overhead. You can avoid this by prefetching the department
// relationship in the Employee fetch, as illustrated in Listing 1.
//
// (For more details, see Core Data Performance in Core Data Programming
// Guide)
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/relationshipKeyPathsForPrefetching
//
// [value(forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKeyPath:)
func (f NSFetchRequest) RelationshipKeyPathsForPrefetching() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("relationshipKeyPathsForPrefetching"))
	return objc.ConvertSliceToStrings(rv)
}
func (f NSFetchRequest) SetRelationshipKeyPathsForPrefetching(value []string) {
	objc.Send[struct{}](f.ID, objc.Sel("setRelationshipKeyPathsForPrefetching:"), objectivec.StringSliceToNSArray(value))
}

// The result type of the fetch request.
//
// # Discussion
//
// The default value is [NSManagedObjectResultType].
//
// If you set the value to [NSManagedObjectIDResultType], and do not include
// property values in the request, sort orderings are demoted to “best
// efforts” hints.
//
// [NSFetchRequest.IncludesPendingChanges] discusses with whether pending
// changes are taken into account when the `resultType` is set to
// `managedObjectResultType`.
//
// [NSFetchRequest.IncludesPropertyValues] discusses whether property values
// are included or not by default when the `resultType` is set to
// `managedObjectResultType`.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/resultType
func (f NSFetchRequest) ResultType() NSFetchRequestResultType {
	rv := objc.Send[NSFetchRequestResultType](f.ID, objc.Sel("resultType"))
	return NSFetchRequestResultType(rv)
}
func (f NSFetchRequest) SetResultType(value NSFetchRequestResultType) {
	objc.Send[struct{}](f.ID, objc.Sel("setResultType:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed, it
// matches against currently unsaved changes in the managed object context.
//
// # Discussion
//
// This value is true if when the fetch is executed, the fetch will match
// against currently unsaved changes in the managed object context; otherwise
// the value is false. The default value is true.
//
// If the value is false, the fetch request doesn’t check unsaved changes
// and only returns objects that matched the predicate in the persistent
// store.
//
// # Special Considerations
//
// A value of true is not supported in conjunction with the result type
// [NSDictionaryResultType], including calculation of aggregate results (such
// as `max` and `min`). For dictionaries, the array returned from the fetch
// reflects the current state in the persistent store, and does not take into
// account any pending changes, insertions, or deletions in the context.
//
// If you need to take pending changes into account for some simple
// aggregations like `max` and `min`, you can instead use a normal fetch
// request, sorted on the attribute you want, with a fetch limit of 1.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPendingChanges
func (f NSFetchRequest) IncludesPendingChanges() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("includesPendingChanges"))
	return rv
}
func (f NSFetchRequest) SetIncludesPendingChanges(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setIncludesPendingChanges:"), value)
}

// A collection of either property descriptions or string property names that
// specify which properties should be returned by the fetch.
//
// # Discussion
//
// Property descriptions can either be instances of [NSPropertyDescription] or
// [NSString]. The property descriptions may represent attributes, to-one
// relationships, or expressions. The name of an attribute or relationship
// description must match the name of a description on the fetch request’s
// entity.
//
// # Special Considerations
//
// You must set the entity for the fetch request before setting this value;
// otherwise, [NSFetchRequest] throws an [invalidArgumentException] exception.
//
// This property can be set with [NSManagedObjectResultType] and thereby
// implement a partial faulting (whereby only some of the properties are
// populated) of the returned objects, as well as the [NSDictionaryResultType]
// to define what properties are included in the resulting [NSDictionary].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToFetch
//
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (f NSFetchRequest) PropertiesToFetch() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("propertiesToFetch"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (f NSFetchRequest) SetPropertiesToFetch(value foundation.INSArray) {
	objc.Send[struct{}](f.ID, objc.Sel("setPropertiesToFetch:"), value)
}

// A Boolean value that indicates whether the fetch request returns only
// distinct values for the fields specified by
// [NSFetchRequest.PropertiesToFetch].
//
// # Discussion
//
// This value is used only if a value has been set for
// [NSFetchRequest.PropertiesToFetch].
//
// This value is true if when the fetch is executed, it returns only distinct
// values for the fields specified by [NSFetchRequest.PropertiesToFetch];
// otherwise, the value is false. The default value is false.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsDistinctResults
func (f NSFetchRequest) ReturnsDistinctResults() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("returnsDistinctResults"))
	return rv
}
func (f NSFetchRequest) SetReturnsDistinctResults(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setReturnsDistinctResults:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed,
// property data is obtained from the persistent store.
//
// # Discussion
//
// This value is true if when the fetch is executed, property data is obtained
// from the persistent store; otherwise it is false. The default value is
// true.
//
// You can set [NSFetchRequest.IncludesPropertyValues] to false to avoid
// creating objects to represent property values and thereby reduce memory
// overhead. You typically should only do so, however, if you are sure that
// you will not need the actual property data, or you already have the
// information in the row cache. Otherwise, you will incur multiple trips to
// the database.
//
// During a normal fetch ([NSFetchRequest.IncludesPropertyValues] is true),
// Core Data fetches the object ID and property data for the matching records,
// fills the row cache with the information, and returns managed objects as
// faults (see [NSFetchRequest.ReturnsObjectsAsFaults]). Although these faults
// are managed objects, all of their property data still resides in the row
// cache until the fault is fired. When the fault is fired, Core Data
// retrieves the data from the row cache—there is no need to go back to the
// database.
//
// If [NSFetchRequest.IncludesPropertyValues] is false, then Core Data fetches
// only the object ID information for the matching records—it does not
// populate the row cache. Core Data still returns managed objects because it
// only needs managed object IDs to create faults. However, if you
// subsequently fire the fault, Core Data looks in the (empty) row cache,
// doesn’t find any data, and then goes back to the store a second time for
// the data.
//
// If [NSFetchRequest.IncludesPropertyValues] is true and
// [NSFetchRequest.ResultType] is set to [NSManagedObjectIDResultType], the
// properties are fetched even though they are not being presented to the
// application and can result in a significant performance penalty.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/includesPropertyValues
func (f NSFetchRequest) IncludesPropertyValues() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("includesPropertyValues"))
	return rv
}
func (f NSFetchRequest) SetIncludesPropertyValues(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setIncludesPropertyValues:"), value)
}

// A Boolean value that indicates whether the property values of fetched
// objects will be updated with the current values in the persistent store.
//
// # Discussion
//
// This value is true if the property values of fetched objects will be
// updated with the current values in the persistent store; otherwise, it is
// false.
//
// By default when you fetch objects, they maintain their current property
// values, even if the values in the persistent store have changed. Invoking
// this method with the parameter true means that when the fetch is executed,
// the property values of fetched objects are updated with the current values
// in the persistent store. This is a more convenient way to ensure that
// managed object property values are consistent with the store than by using
// [NSManagedObjectContext.RefreshObjectMergeChanges]
// ([NSManagedObjetContext]) for multiple objects in turn.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/shouldRefreshRefetchedObjects
func (f NSFetchRequest) ShouldRefreshRefetchedObjects() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("shouldRefreshRefetchedObjects"))
	return rv
}
func (f NSFetchRequest) SetShouldRefreshRefetchedObjects(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setShouldRefreshRefetchedObjects:"), value)
}

// A Boolean value that indicates whether the objects resulting from a fetch
// request are faults.
//
// # Discussion
//
// This value is true if the objects resulting from a fetch using the
// [NSFetchRequest] are faults; otherwise, it is false. The default value is
// true. This setting is not used if the result type (see
// [NSFetchRequest.ResultType]) is [NSManagedObjectIDResultType], as object
// IDs do not have property values. You can set
// [NSFetchRequest.ReturnsObjectsAsFaults] to false to gain a performance
// benefit if you know you will need to access the property values from the
// returned objects.
//
// When you execute a fetch, by default
// [NSFetchRequest.ReturnsObjectsAsFaults] is true; Core Data fetches the
// object data for the matching records, fills the row cache with the
// information, and returns managed object as faults. These faults are managed
// objects, but all of their property data resides in the row cache until the
// fault is fired. When the fault is fired, Core Data retrieves the data from
// the row cache. Although the overhead for this operation is small, for large
// datasets it may not be trivial. If you need to access the property values
// from the returned objects (for example, if you iterate over all the objects
// to calculate the average value of a particular attribute), then it is more
// efficient to set [NSFetchRequest.ReturnsObjectsAsFaults] to false to avoid
// the additional overhead.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/returnsObjectsAsFaults
func (f NSFetchRequest) ReturnsObjectsAsFaults() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("returnsObjectsAsFaults"))
	return rv
}
func (f NSFetchRequest) SetReturnsObjectsAsFaults(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setReturnsObjectsAsFaults:"), value)
}

// An array of objects that indicates how data should be grouped before a
// select statement is run in a SQL database.
//
// # Discussion
//
// An array of [NSPropertyDescription] or [NSExpressionDescription] objects or
// key-path strings that indicate how data should be grouped before a select
// statement is run in an SQL database.
//
// If you use this setting, you must set the [NSFetchRequest.ResultType] to
// [NSDictionaryResultType], and the SELECT values must be literals,
// aggregates, or columns specified in `propertiesToGroupBy`.
//
// Aggregates will operate on the groups specified in `propertiesToGroupBy`
//
// rather than the whole table. If you set `propertiesToGroupBy`, you can also
// set a predicate to filter rows that are returned by `propertiesToGroupBy`.
//
// See [NSFetchRequest.HavingPredicate].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/propertiesToGroupBy
func (f NSFetchRequest) PropertiesToGroupBy() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("propertiesToGroupBy"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (f NSFetchRequest) SetPropertiesToGroupBy(value foundation.INSArray) {
	objc.Send[struct{}](f.ID, objc.Sel("setPropertiesToGroupBy:"), value)
}

// The predicate used to filter rows being returned by a query containing a
// GROUP BY directive.
//
// # Discussion
//
// If a `havingPredicate` value is supplied, the predicate will be run after.
// Specifying a `havingPredicate` requires that
// [NSFetchRequest.PropertiesToGroupBy] also be specified.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/havingPredicate
func (f NSFetchRequest) HavingPredicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("havingPredicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (f NSFetchRequest) SetHavingPredicate(value foundation.NSPredicate) {
	objc.Send[struct{}](f.ID, objc.Sel("setHavingPredicate:"), value)
}
