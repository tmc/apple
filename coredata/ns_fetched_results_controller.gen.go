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

// The class instance for the [NSFetchedResultsController] class.
var (
	_NSFetchedResultsControllerClass     NSFetchedResultsControllerClass
	_NSFetchedResultsControllerClassOnce sync.Once
)

func getNSFetchedResultsControllerClass() NSFetchedResultsControllerClass {
	_NSFetchedResultsControllerClassOnce.Do(func() {
		_NSFetchedResultsControllerClass = NSFetchedResultsControllerClass{class: objc.GetClass("NSFetchedResultsController")}
	})
	return _NSFetchedResultsControllerClass
}

// GetNSFetchedResultsControllerClass returns the class object for NSFetchedResultsController.
func GetNSFetchedResultsControllerClass() NSFetchedResultsControllerClass {
	return getNSFetchedResultsControllerClass()
}

type NSFetchedResultsControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchedResultsControllerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchedResultsControllerClass) Alloc() NSFetchedResultsController {
	rv := objc.Send[NSFetchedResultsController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A controller that you use to manage the results of a Core Data fetch
// request and to display data to the user.
//
// # Overview
//
// While you can use table views can in several ways, fetched results
// controllers primarily assist you with a primary list view. [UITableView]
// expects its data source to provide cells as an array of sections made up of
// rows. You configure a fetched results controller using a fetch request —
// an object that specifies what type of entity to fetch and how to sort the
// results. You can also add criteria for when to include a specific instance
// of the entity.
//
// The fetched results controller efficiently analyzes the result of the fetch
// request and computes all the information about sections in the result set.
// It also computes all the information for the index based on the result set.
//
// In addition, fetched results controllers:
//
// - Optionally monitor changes to objects in the associated managed object
// context, and report changes in the results set to its delegate (see
// [NSFetchedResultsController]). - Optionally cache the results of its
// computation to enable redisplaying the same data without repeating the work
// to fetch it. For more information, see [NSFetchedResultsController].
//
// A controller thus effectively has three modes of operation, determined by
// whether it has a delegate and whether you set the cache file name.
//
// - No tracking: The delegate is `nil`. The controller provides access to the
// data as it was when it fetched it. - Memory-only tracking: the delegate is
// non-`nil` and the file cache name is `nil`. The controller monitors objects
// in its result set and updates section and ordering information in response
// to relevant changes. - Full persistent tracking: the delegate and the file
// cache name are non-`nil`. The controller monitors objects in its result set
// and updates section and ordering information in response to relevant
// changes. The controller maintains a persistent cache of the results of its
// computation.
//
// # Using NSFetchedResultsController
//
// # Creating the fetched results controller
//
// You typically create an instance of [NSFetchedResultsController] as an
// instance variable of a table view controller. When you initialize the fetch
// results controller, you provide four parameters:
//
// - A fetch request. This must contain at least one sort descriptor to order
// the results. - A managed object context. The controller uses this context
// to execute the fetch request. - Optionally, a key path on result objects
// that returns the section name. The controller uses the key path to split
// the results into sections (passing `nil` indicates that the controller
// should generate a single section). - Optionally, the name of the cache file
// the controller should use (passing `nil` prevents caching). Using a cache
// can avoid the overhead of computing the section and index information.
//
// After creating an instance, you invoke
// [NSFetchedResultsController.PerformFetch] to actually execute the fetch:
//
// # The controller’s delegate
//
// If you set a delegate for a fetched results controller, the controller
// registers to receive change notifications from its managed object context.
// The controller processes any change in the context that affects the result
// set or section information and updates the results as necessary. The
// controller notifies the delegate when result objects change location or
// when changes occur in sections. For more information, see
// [NSFetchedResultsControllerDelegate]. You typically use these methods to
// update the display of the table view.
//
// # The cache
//
// Where possible, a controller uses a cache to avoid the need to repeat work
// performed in setting up any sections and ordering the contents. The system
// maintains the cache across launches of your application.
//
// When you initialize an instance of [NSFetchedResultsController], you
// typically specify a cache name. If you don’t specify a cache name, the
// controller doesn’t cache data. When you create a controller, it looks for
// an existing cache with the given name:
//
// - If the controller can’t find an appropriate cache, it calculates the
// required sections and the order of objects within sections. It then writes
// this information to disk. - If it finds a cache with the same name, the
// controller tests the cache to determine whether its contents are still
// valid. The controller compares the current entity name, entity version
// hash, sort descriptors, and section key-path with those stored in the
// cache, as well as the modification date of the cached information file and
// the persistent store file.
//
// If the cache is consistent with the current information, the controller
// reuses the previously-computed information.
//
// If the cache isn’t consistent with the current information, then the
// controller recomputes the required information and updates the cache.
//
// Any time the section and ordering information change, the controller
// updates cache.
//
// If you create multiple fetched results controllers with different
// configurations, such as different sort descriptors, give each configuration
// a different cache name.
//
// You can purge a cache using
// [NSFetchedResultsControllerClass.DeleteCacheWithName].
//
// # Implementing the table view datasource methods
//
// You ask the object to provide relevant information in your implementation
// of the table view data source methods:
//
// # Responding to changes
//
// I[NSFetchedResultsController] responds to changes at the model layer, and
// informs its delegate when result objects change location or when sections
// change.
//
// If you allow a user to reorder table rows, then your implementation of the
// delegate methods must take this into account; see
// [NSFetchedResultsControllerDelegate].
//
// The controller doesn’t show changes until after its managed object
// context receives a [NSManagedObjectContext.ProcessPendingChanges] message.
// Therefore, if you change the value of a managed object’s attribute so
// that its location in a fetched results controller’s results set changes,
// its index as reported by the controller won’t typically change until the
// end of the current event cycle, when the system calls
// [NSManagedObjectContext.ProcessPendingChanges]. For example, the following
// code would log `“same”`:
//
// # Modifying the fetch request
//
// You can’t change the fetch request to modify the results. Do the
// following if you want to change the fetch request:
//
// - Delete the cache if you’re using one, by calling
// [NSFetchedResultsControllerClass.DeleteCacheWithName]. - Change the fetch
// request. - Call [NSFetchedResultsController.PerformFetch].
//
// # Handling object invalidation
//
// When a managed object context notifies the fetched results controller of
// invalidated individual objects, the controller treats these as deleted
// objects and sends the proper delegate calls.
//
// Simultaneous invalidation of all the objects in a managed object context is
// possible, for example, as a result of calling
// [NSManagedObjectContext.Reset], or if you remove a store from the
// persistent store coordinator. When this happens,
// [NSFetchedResultsController] doesn’t invalidate all objects, nor does it
// send individual notifications for object deletions. Instead, you need to
// call [NSFetchedResultsController.PerformFetch] to reset the state of the
// controller then reload the data in the table view ([reloadData()]).
//
// # Subclassing notes
//
// You create a subclass of this class if you want to customize the creation
// of sections and index titles. You override
// [NSFetchedResultsController.SectionIndexTitleForSectionName] if you want
// the section index title to be something other than the capitalized first
// letter of the section name. You override
// [NSFetchedResultsController.SectionIndexTitles] if you want the index
// titles to be something other than the array created by calling
// [NSFetchedResultsController.SectionIndexTitleForSectionName] on all the
// known sections.
//
// # Initializing a Fetched Results Controller
//
//   - [NSFetchedResultsController.InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName]: Returns a fetch request controller initialized using the given arguments.
//   - [NSFetchedResultsController.PerformFetch]: Executes the controller’s fetch request.
//
// # Getting Configuration Information
//
//   - [NSFetchedResultsController.FetchRequest]: The fetch request used to do the fetching.
//   - [NSFetchedResultsController.ManagedObjectContext]: The managed object context used to fetch objects.
//   - [NSFetchedResultsController.SectionNameKeyPath]: The key path of the attribute that determines which section the fetched entity belongs to.
//   - [NSFetchedResultsController.CacheName]: The name of the file used to cache section information.
//   - [NSFetchedResultsController.Delegate]: The object that is notified when the fetched results changed.
//   - [NSFetchedResultsController.SetDelegate]
//
// # Accessing Results
//
//   - [NSFetchedResultsController.FetchedObjects]: The results of the fetch.
//   - [NSFetchedResultsController.ObjectAtIndexPath]: Returns the object at the given index path in the fetch results.
//   - [NSFetchedResultsController.IndexPathForObject]: Returns the index path of a given object.
//
// # Querying Section Information
//
//   - [NSFetchedResultsController.Sections]: The sections for the fetch results.
//   - [NSFetchedResultsController.SectionForSectionIndexTitleAtIndex]: Returns the section number for a given section title and index in the section index.
//
// # Configuring Section Information
//
//   - [NSFetchedResultsController.SectionIndexTitleForSectionName]: Returns the corresponding section index entry for a given section name.
//   - [NSFetchedResultsController.SectionIndexTitles]: The array of section index titles.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
//
// [UITableView]: https://developer.apple.com/documentation/UIKit/UITableView
// [reloadData()]: https://developer.apple.com/documentation/UIKit/UITableView/reloadData()
type NSFetchedResultsController struct {
	objectivec.Object
}

// NSFetchedResultsControllerFromID constructs a [NSFetchedResultsController] from an objc.ID.
//
// A controller that you use to manage the results of a Core Data fetch
// request and to display data to the user.
func NSFetchedResultsControllerFromID(id objc.ID) NSFetchedResultsController {
	return NSFetchedResultsController{objectivec.Object{ID: id}}
}

// NOTE: NSFetchedResultsController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchedResultsController] class.
//
// # Initializing a Fetched Results Controller
//
//   - [INSFetchedResultsController.InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName]: Returns a fetch request controller initialized using the given arguments.
//   - [INSFetchedResultsController.PerformFetch]: Executes the controller’s fetch request.
//
// # Getting Configuration Information
//
//   - [INSFetchedResultsController.FetchRequest]: The fetch request used to do the fetching.
//   - [INSFetchedResultsController.ManagedObjectContext]: The managed object context used to fetch objects.
//   - [INSFetchedResultsController.SectionNameKeyPath]: The key path of the attribute that determines which section the fetched entity belongs to.
//   - [INSFetchedResultsController.CacheName]: The name of the file used to cache section information.
//   - [INSFetchedResultsController.Delegate]: The object that is notified when the fetched results changed.
//   - [INSFetchedResultsController.SetDelegate]
//
// # Accessing Results
//
//   - [INSFetchedResultsController.FetchedObjects]: The results of the fetch.
//   - [INSFetchedResultsController.ObjectAtIndexPath]: Returns the object at the given index path in the fetch results.
//   - [INSFetchedResultsController.IndexPathForObject]: Returns the index path of a given object.
//
// # Querying Section Information
//
//   - [INSFetchedResultsController.Sections]: The sections for the fetch results.
//   - [INSFetchedResultsController.SectionForSectionIndexTitleAtIndex]: Returns the section number for a given section title and index in the section index.
//
// # Configuring Section Information
//
//   - [INSFetchedResultsController.SectionIndexTitleForSectionName]: Returns the corresponding section index entry for a given section name.
//   - [INSFetchedResultsController.SectionIndexTitles]: The array of section index titles.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
type INSFetchedResultsController interface {
	objectivec.IObject

	// Topic: Initializing a Fetched Results Controller

	// Returns a fetch request controller initialized using the given arguments.
	InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest INSFetchRequest, context INSManagedObjectContext, sectionNameKeyPath string, name string) NSFetchedResultsController
	// Executes the controller’s fetch request.
	PerformFetch() (bool, error)

	// Topic: Getting Configuration Information

	// The fetch request used to do the fetching.
	FetchRequest() INSFetchRequest
	// The managed object context used to fetch objects.
	ManagedObjectContext() INSManagedObjectContext
	// The key path of the attribute that determines which section the fetched entity belongs to.
	SectionNameKeyPath() string
	// The name of the file used to cache section information.
	CacheName() string
	// The object that is notified when the fetched results changed.
	Delegate() NSFetchedResultsControllerDelegate
	SetDelegate(value NSFetchedResultsControllerDelegate)

	// Topic: Accessing Results

	// The results of the fetch.
	FetchedObjects() []objectivec.IObject
	// Returns the object at the given index path in the fetch results.
	ObjectAtIndexPath(indexPath foundation.NSIndexPath) unsafe.Pointer
	// Returns the index path of a given object.
	IndexPathForObject(object unsafe.Pointer) foundation.NSIndexPath

	// Topic: Querying Section Information

	// The sections for the fetch results.
	Sections() []objectivec.IObject
	// Returns the section number for a given section title and index in the section index.
	SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int

	// Topic: Configuring Section Information

	// Returns the corresponding section index entry for a given section name.
	SectionIndexTitleForSectionName(sectionName string) string
	// The array of section index titles.
	SectionIndexTitles() []string
}

// Init initializes the instance.
func (f NSFetchedResultsController) Init() NSFetchedResultsController {
	rv := objc.Send[NSFetchedResultsController](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchedResultsController) Autorelease() NSFetchedResultsController {
	rv := objc.Send[NSFetchedResultsController](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchedResultsController creates a new NSFetchedResultsController instance.
func NewNSFetchedResultsController() NSFetchedResultsController {
	class := getNSFetchedResultsControllerClass()
	rv := objc.Send[NSFetchedResultsController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a fetch request controller initialized using the given arguments.
//
// fetchRequest: The fetch request used to get the objects.
//
// The fetch request must have at least one sort descriptor. If the controller
// generates sections, the first sort descriptor in the array is used to group
// the objects into sections; its key must either be the same as
// `sectionNameKeyPath` or the relative ordering using its key must match that
// using `sectionNameKeyPath`.
//
// context: The managed object against which `fetchRequest` is executed.
//
// sectionNameKeyPath: A key path on result objects that returns the section name. Pass `nil` to
// indicate that the controller should generate a single section.
//
// The section name is used to pre-compute the section information.
//
// If this key path is not the same as that specified by the first sort
// descriptor in `fetchRequest`, they must generate the same relative
// orderings. For example, the first sort descriptor in `fetchRequest` might
// specify the key for a persistent property; `sectionNameKeyPath` might
// specify a key for a transient property derived from the persistent
// property.
//
// name: The name of the cache file the receiver should use. Pass `nil` to prevent
// caching.
//
// Pre-computed section info is cached to a private directory under this name.
// If Core Data finds a cache stored with this name, it is checked to see if
// it matches the `fetchRequest`. If it does, the cache is loaded
// directly—this avoids the overhead of computing the section and index
// information. If the cached information doesn’t match the request, the
// cache is deleted and recomputed when the fetch happens.
//
// # Return Value
//
// The receiver initialized with the specified fetch request, context, key
// path, and cache name.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/init(fetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:)
func NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest INSFetchRequest, context INSManagedObjectContext, sectionNameKeyPath string, name string) NSFetchedResultsController {
	instance := getNSFetchedResultsControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:"), fetchRequest, context, objc.String(sectionNameKeyPath), objc.String(name))
	return NSFetchedResultsControllerFromID(rv)
}

// Returns a fetch request controller initialized using the given arguments.
//
// fetchRequest: The fetch request used to get the objects.
//
// The fetch request must have at least one sort descriptor. If the controller
// generates sections, the first sort descriptor in the array is used to group
// the objects into sections; its key must either be the same as
// `sectionNameKeyPath` or the relative ordering using its key must match that
// using `sectionNameKeyPath`.
//
// context: The managed object against which `fetchRequest` is executed.
//
// sectionNameKeyPath: A key path on result objects that returns the section name. Pass `nil` to
// indicate that the controller should generate a single section.
//
// The section name is used to pre-compute the section information.
//
// If this key path is not the same as that specified by the first sort
// descriptor in `fetchRequest`, they must generate the same relative
// orderings. For example, the first sort descriptor in `fetchRequest` might
// specify the key for a persistent property; `sectionNameKeyPath` might
// specify a key for a transient property derived from the persistent
// property.
//
// name: The name of the cache file the receiver should use. Pass `nil` to prevent
// caching.
//
// Pre-computed section info is cached to a private directory under this name.
// If Core Data finds a cache stored with this name, it is checked to see if
// it matches the `fetchRequest`. If it does, the cache is loaded
// directly—this avoids the overhead of computing the section and index
// information. If the cached information doesn’t match the request, the
// cache is deleted and recomputed when the fetch happens.
//
// # Return Value
//
// The receiver initialized with the specified fetch request, context, key
// path, and cache name.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/init(fetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:)
func (f NSFetchedResultsController) InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest INSFetchRequest, context INSManagedObjectContext, sectionNameKeyPath string, name string) NSFetchedResultsController {
	rv := objc.Send[NSFetchedResultsController](f.ID, objc.Sel("initWithFetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:"), fetchRequest, context, objc.String(sectionNameKeyPath), objc.String(name))
	return rv
}

// Executes the controller’s fetch request.
//
// # Discussion
//
// After you execute this method, access the controller’s fetched objects
// using the [NSFetchedResultsController.FetchedObjects] property.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/performFetch()
func (f NSFetchedResultsController) PerformFetch() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("performFetch:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performFetch: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the object at the given index path in the fetch results.
//
// indexPath: An index path in the fetch results.
//
// If `indexPath` does not describe a valid index path in the fetch results,
// an exception is raised.
//
// # Return Value
//
// The object at a given index path in the fetch results.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/object(at:)
func (f NSFetchedResultsController) ObjectAtIndexPath(indexPath foundation.NSIndexPath) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f.ID, objc.Sel("objectAtIndexPath:"), indexPath)
	return rv
}

// Returns the index path of a given object.
//
// object: An object in the receiver’s fetch results.
//
// # Return Value
//
// The index path of `object` in the receiver’s fetch results, or `nil` if
// `object` could not be found.
//
// # Discussion
//
// In versions of iOS before 3.2, this method raises an exception if `object`
// is not contained in the receiver’s fetch results.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/indexPath(forObject:)
func (f NSFetchedResultsController) IndexPathForObject(object unsafe.Pointer) foundation.NSIndexPath {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("indexPathForObject:"), object)
	return foundation.NSIndexPathFromID(rv)
}

// Returns the section number for a given section title and index in the
// section index.
//
// title: The title of a section
//
// sectionIndex: The index of a section.
//
// # Return Value
//
// The section number for the given section title and index in the section
// index
//
// # Discussion
//
// You would typically call this method when executing
// [UITableViewDataSource]’s [tableView(_:sectionForSectionIndexTitle:at:)]
// method.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/section(forSectionIndexTitle:at:)
//
// [tableView(_:sectionForSectionIndexTitle:at:)]: https://developer.apple.com/documentation/UIKit/UITableViewDataSource/tableView(_:sectionForSectionIndexTitle:at:)
func (f NSFetchedResultsController) SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int {
	rv := objc.Send[int](f.ID, objc.Sel("sectionForSectionIndexTitle:atIndex:"), objc.String(title), sectionIndex)
	return rv
}

// Returns the corresponding section index entry for a given section name.
//
// sectionName: The name of a section.
//
// # Return Value
//
// The section index entry corresponding to the section with name
// `sectionName`.
//
// # Discussion
//
// The default implementation returns the capitalized first letter of the
// section name.
//
// You should override this method if you need a different way to convert from
// a section name to its name in the section index.
//
// # Special Considerations
//
// You only need this method if you use a section index.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionIndexTitle(forSectionName:)
func (f NSFetchedResultsController) SectionIndexTitleForSectionName(sectionName string) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sectionIndexTitleForSectionName:"), objc.String(sectionName))
	return foundation.NSStringFromID(rv).String()
}

// Deletes the cached section information with the given name.
//
// name: The name of the cache file to delete.
//
// If `name` is `nil`, deletes all cache files.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/deleteCache(withName:)
func (_NSFetchedResultsControllerClass NSFetchedResultsControllerClass) DeleteCacheWithName(name string) {
	objc.Send[objc.ID](objc.ID(_NSFetchedResultsControllerClass.class), objc.Sel("deleteCacheWithName:"), objc.String(name))
}

// The fetch request used to do the fetching.
//
// # Discussion
//
// If you want to modify the fetch request, you must follow the steps
// described in [NSFetchedResultsController].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/fetchRequest
func (f NSFetchedResultsController) FetchRequest() INSFetchRequest {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}

// The managed object context used to fetch objects.
//
// # Discussion
//
// The controller registers to listen to change notifications on this context
// and properly update its result set and section information.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/managedObjectContext
func (f NSFetchedResultsController) ManagedObjectContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("managedObjectContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// The key path of the attribute that determines which section the fetched
// entity belongs to.
//
// # Discussion
//
// This property returns the value you specify for the `sectionNameKeyPath`
// parameter when you initialize the fetched results controller.
//
// If the controller generates sections, typically this property’s value
// matches the specified key path of the first sort descriptor in the
// controller’s fetch request. If the two key paths don’t match, then they
// must generate the same relative ordering. For example, the fetch
// request’s first sort descriptor might specify the key path of a
// persistent attribute, but [NSFetchedResultsController.SectionNameKeyPath]
// might specify the key path of a transient attribute that derives its value
// from the persistent one.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionNameKeyPath
func (f NSFetchedResultsController) SectionNameKeyPath() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sectionNameKeyPath"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the file used to cache section information.
//
// # Discussion
//
// The file itself is stored in a private directory; you can only access it by
// name using [NSFetchedResultsControllerClass.DeleteCacheWithName]
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/cacheName
func (f NSFetchedResultsController) CacheName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("cacheName"))
	return foundation.NSStringFromID(rv).String()
}

// The object that is notified when the fetched results changed.
//
// # Discussion
//
// If you do not specify a delegate, the controller does not track changes to
// managed objects associated with its managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/delegate
func (f NSFetchedResultsController) Delegate() NSFetchedResultsControllerDelegate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("delegate"))
	return NSFetchedResultsControllerDelegateObjectFromID(rv)
}
func (f NSFetchedResultsController) SetDelegate(value NSFetchedResultsControllerDelegate) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelegate:"), value)
}

// The results of the fetch.
//
// # Discussion
//
// The value of the property is `nil` if
// [NSFetchedResultsController.PerformFetch] hasn’t been called.
//
// The results array only includes instances of the entity specified by the
// fetch request ([NSFetchedResultsController.FetchRequest]) and that match
// its predicate. (If the fetch request has no predicate, then the results
// array includes all instances of the entity specified by the fetch request.)
//
// The results array reflects the in-memory state of managed objects in the
// controller’s managed object context, not their state in the persistent
// store. The returned array does not, however, update as managed objects are
// inserted, modified, or deleted.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/fetchedObjects
func (f NSFetchedResultsController) FetchedObjects() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("fetchedObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The sections for the fetch results.
//
// # Discussion
//
// The objects in the sections array implement the
// [NSFetchedResultsSectionInfo] protocol.
//
// You typically use the sections array when implementing
// [UITableViewDataSource] methods, such as [numberOfSections(in:)] and
// [tableView(_:titleForHeaderInSection:)].
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sections
//
// [numberOfSections(in:)]: https://developer.apple.com/documentation/UIKit/UITableViewDataSource/numberOfSections(in:)
// [tableView(_:titleForHeaderInSection:)]: https://developer.apple.com/documentation/UIKit/UITableViewDataSource/tableView(_:titleForHeaderInSection:)
func (f NSFetchedResultsController) Sections() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("sections"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The array of section index titles.
//
// # Discussion
//
// The default implementation returns the array created by calling
// [NSFetchedResultsController.SectionIndexTitleForSectionName] on all the
// known sections. You should override this method if you want to return a
// different array for the section index.
//
// # Special Considerations
//
// You only need this method if you use a section index.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController/sectionIndexTitles
func (f NSFetchedResultsController) SectionIndexTitles() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("sectionIndexTitles"))
	return objc.ConvertSliceToStrings(rv)
}
