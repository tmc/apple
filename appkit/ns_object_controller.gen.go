// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSObjectController] class.
var (
	_NSObjectControllerClass     NSObjectControllerClass
	_NSObjectControllerClassOnce sync.Once
)

func getNSObjectControllerClass() NSObjectControllerClass {
	_NSObjectControllerClassOnce.Do(func() {
		_NSObjectControllerClass = NSObjectControllerClass{class: objc.GetClass("NSObjectController")}
	})
	return _NSObjectControllerClass
}

// GetNSObjectControllerClass returns the class object for NSObjectController.
func GetNSObjectControllerClass() NSObjectControllerClass {
	return getNSObjectControllerClass()
}

type NSObjectControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSObjectControllerClass) Alloc() NSObjectController {
	rv := objc.Send[NSObjectController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A controller that can manage an object’s properties referenced by
// key-value paths.
//
// # Overview
// 
// [NSObjectController] is a Cocoa bindings–compatible controller class.
// Properties of the content object of instances of this class can be bound to
// user interface elements to access and modify their values.
// 
// By default, the content of an [NSObjectController] instance is an
// [NSMutableDictionary] object. This allows a single [NSObjectController]
// instance to be used to manage many different properties referenced by
// key-value paths. The default content object class can be changed by calling
// [NSObjectController.ObjectClass], which subclasses must override. Your application should use
// a custom data class that is key-value compliant whenever possible.
// 
// # Object Controllers, Entity Mode, and Lazy Fetching
// 
// [NSObjectController] and its subclasses, when in entity mode, can now fetch
// lazily. With lazy fetching enabled using the property [NSObjectController.UsesLazyFetching],
// the controller will try to fetch only a small amount of data from available
// persistent stores. This can provide a significant improvement in memory use
// when a large amount of content is stored on disk but just a subset of that
// data is required in memory.
// 
// When set to use lazy fetching, a controller will fetch objects in batches.
// You can change the default batch size for your application by setting a
// value for the the user default
// “`com.Apple().CocoaBindings.LazyFetchBatchSize`”. If you have table
// views bound to an array controller set to use lazy fetching, the size of
// the controller’s batch size will grow as the table views’ visible row
// count grows.
// 
// Add, Insert, and Remove operations on controllers that use lazy fetching
// behave similarly to the same operations on a regular controller. The
// difference is that it is faster to sort an array controller using lazy
// fetching if:
// 
// - All of the keys in the `sortDescriptors` array are modeled, non transient
// properties. - All of the selectors in the `sortDescriptors` array are `` or
// ``. - There are no changes in the controller’s managed object context
//
// [NSMutableDictionary]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary
//
// # Initializing an object controller
//
//   - [NSObjectController.InitWithContent]: Initializes and returns an [NSObjectController] object with the given content.
//
// # Managing content
//
//   - [NSObjectController.Content]: The receiver’s content object.
//   - [NSObjectController.SetContent]
//   - [NSObjectController.AutomaticallyPreparesContent]: A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//   - [NSObjectController.SetAutomaticallyPreparesContent]
//   - [NSObjectController.PrepareContent]: Typically overridden by subclasses that require additional control over the creation of new objects.
//
// # Setting the content class
//
//   - [NSObjectController.ObjectClass]: The object class to use when creating new objects.
//   - [NSObjectController.SetObjectClass]
//
// # Managing objects
//
//   - [NSObjectController.NewObject]: Creates and returns a new object of the appropriate class.
//   - [NSObjectController.AddObject]: Sets the receiver’s content object.
//   - [NSObjectController.RemoveObject]: Removes a given object from the receiver’s content.
//   - [NSObjectController.Add]: Creates a new object and sets it as the receiver’s content object.
//   - [NSObjectController.CanAdd]: A Boolean value that indicates whether an object can be added to the receiver using [add(_:)](<doc://com.apple.appkit/documentation/AppKit/NSObjectController/add(_:)>).
//   - [NSObjectController.Remove]: Removes the receiver’s content object.
//   - [NSObjectController.CanRemove]: A Boolean value that indicates whether an object can be removed from the receiver.
//
// # Managing editing
//
//   - [NSObjectController.Editable]: A Boolean that indicates whether the receiver allows adding and removing objects.
//   - [NSObjectController.SetEditable]
//
// # Core Data support
//
//   - [NSObjectController.EntityName]: The entity name used by the receiver to create new objects.
//   - [NSObjectController.SetEntityName]
//   - [NSObjectController.Fetch]: Causes the receiver to fetch the data objects specified by the entity name and fetch predicate.
//   - [NSObjectController.UsesLazyFetching]: A Boolean that indicates whether the receiver uses lazy fetching.
//   - [NSObjectController.SetUsesLazyFetching]
//   - [NSObjectController.DefaultFetchRequest]: Returns the default fetch request used by the receiver.
//   - [NSObjectController.FetchPredicate]: The receiver’s fetch predicate.
//   - [NSObjectController.SetFetchPredicate]
//   - [NSObjectController.ManagedObjectContext]: The receiver’s managed object context.
//   - [NSObjectController.SetManagedObjectContext]
//   - [NSObjectController.FetchWithRequestMergeError]: Subclasses should override this method to customize a fetch request, for example to specify fetch limits.
//
// # Obtaining selections
//
//   - [NSObjectController.SelectedObjects]: An array of all objects to be affected by editing.
//   - [NSObjectController.Selection]: A proxy object representing the receiver’s selection.
//
// # Validating user interface items
//
//   - [NSObjectController.ValidateUserInterfaceItem]: Returns whether the receiver can handle the action method for a user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController
type NSObjectController struct {
	NSController
}

// NSObjectControllerFromID constructs a [NSObjectController] from an objc.ID.
//
// A controller that can manage an object’s properties referenced by
// key-value paths.
func NSObjectControllerFromID(id objc.ID) NSObjectController {
	return NSObjectController{NSController: NSControllerFromID(id)}
}
// NOTE: NSObjectController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSObjectController] class.
//
// # Initializing an object controller
//
//   - [INSObjectController.InitWithContent]: Initializes and returns an [NSObjectController] object with the given content.
//
// # Managing content
//
//   - [INSObjectController.Content]: The receiver’s content object.
//   - [INSObjectController.SetContent]
//   - [INSObjectController.AutomaticallyPreparesContent]: A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//   - [INSObjectController.SetAutomaticallyPreparesContent]
//   - [INSObjectController.PrepareContent]: Typically overridden by subclasses that require additional control over the creation of new objects.
//
// # Setting the content class
//
//   - [INSObjectController.ObjectClass]: The object class to use when creating new objects.
//   - [INSObjectController.SetObjectClass]
//
// # Managing objects
//
//   - [INSObjectController.NewObject]: Creates and returns a new object of the appropriate class.
//   - [INSObjectController.AddObject]: Sets the receiver’s content object.
//   - [INSObjectController.RemoveObject]: Removes a given object from the receiver’s content.
//   - [INSObjectController.Add]: Creates a new object and sets it as the receiver’s content object.
//   - [INSObjectController.CanAdd]: A Boolean value that indicates whether an object can be added to the receiver using [add(_:)](<doc://com.apple.appkit/documentation/AppKit/NSObjectController/add(_:)>).
//   - [INSObjectController.Remove]: Removes the receiver’s content object.
//   - [INSObjectController.CanRemove]: A Boolean value that indicates whether an object can be removed from the receiver.
//
// # Managing editing
//
//   - [INSObjectController.Editable]: A Boolean that indicates whether the receiver allows adding and removing objects.
//   - [INSObjectController.SetEditable]
//
// # Core Data support
//
//   - [INSObjectController.EntityName]: The entity name used by the receiver to create new objects.
//   - [INSObjectController.SetEntityName]
//   - [INSObjectController.Fetch]: Causes the receiver to fetch the data objects specified by the entity name and fetch predicate.
//   - [INSObjectController.UsesLazyFetching]: A Boolean that indicates whether the receiver uses lazy fetching.
//   - [INSObjectController.SetUsesLazyFetching]
//   - [INSObjectController.DefaultFetchRequest]: Returns the default fetch request used by the receiver.
//   - [INSObjectController.FetchPredicate]: The receiver’s fetch predicate.
//   - [INSObjectController.SetFetchPredicate]
//   - [INSObjectController.ManagedObjectContext]: The receiver’s managed object context.
//   - [INSObjectController.SetManagedObjectContext]
//   - [INSObjectController.FetchWithRequestMergeError]: Subclasses should override this method to customize a fetch request, for example to specify fetch limits.
//
// # Obtaining selections
//
//   - [INSObjectController.SelectedObjects]: An array of all objects to be affected by editing.
//   - [INSObjectController.Selection]: A proxy object representing the receiver’s selection.
//
// # Validating user interface items
//
//   - [INSObjectController.ValidateUserInterfaceItem]: Returns whether the receiver can handle the action method for a user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController
type INSObjectController interface {
	INSController

	// Topic: Initializing an object controller

	// Initializes and returns an [NSObjectController] object with the given content.
	InitWithContent(content objectivec.IObject) NSObjectController

	// Topic: Managing content

	// The receiver’s content object.
	Content() objectivec.IObject
	SetContent(value objectivec.IObject)
	// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	// Typically overridden by subclasses that require additional control over the creation of new objects.
	PrepareContent()

	// Topic: Setting the content class

	// The object class to use when creating new objects.
	ObjectClass() objc.Class
	SetObjectClass(value objc.Class)

	// Topic: Managing objects

	// Creates and returns a new object of the appropriate class.
	NewObject() objectivec.IObject
	// Sets the receiver’s content object.
	AddObject(object objectivec.IObject)
	// Removes a given object from the receiver’s content.
	RemoveObject(object objectivec.IObject)
	// Creates a new object and sets it as the receiver’s content object.
	Add(sender objectivec.IObject)
	// A Boolean value that indicates whether an object can be added to the receiver using [add(_:)](<doc://com.apple.appkit/documentation/AppKit/NSObjectController/add(_:)>).
	CanAdd() bool
	// Removes the receiver’s content object.
	Remove(sender objectivec.IObject)
	// A Boolean value that indicates whether an object can be removed from the receiver.
	CanRemove() bool

	// Topic: Managing editing

	// A Boolean that indicates whether the receiver allows adding and removing objects.
	Editable() bool
	SetEditable(value bool)

	// Topic: Core Data support

	// The entity name used by the receiver to create new objects.
	EntityName() string
	SetEntityName(value string)
	// Causes the receiver to fetch the data objects specified by the entity name and fetch predicate.
	Fetch(sender objectivec.IObject)
	// A Boolean that indicates whether the receiver uses lazy fetching.
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
	// Returns the default fetch request used by the receiver.
	DefaultFetchRequest() objectivec.IObject
	// The receiver’s fetch predicate.
	FetchPredicate() foundation.INSPredicate
	SetFetchPredicate(value foundation.INSPredicate)
	// The receiver’s managed object context.
	ManagedObjectContext() objectivec.IObject
	SetManagedObjectContext(value objectivec.IObject)
	// Subclasses should override this method to customize a fetch request, for example to specify fetch limits.
	FetchWithRequestMergeError(fetchRequest objectivec.IObject, merge bool) (bool, error)

	// Topic: Obtaining selections

	// An array of all objects to be affected by editing.
	SelectedObjects() foundation.INSArray
	// A proxy object representing the receiver’s selection.
	Selection() objectivec.IObject

	// Topic: Validating user interface items

	// Returns whether the receiver can handle the action method for a user interface item.
	ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool
}

// Init initializes the instance.
func (o NSObjectController) Init() NSObjectController {
	rv := objc.Send[NSObjectController](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSObjectController) Autorelease() NSObjectController {
	rv := objc.Send[NSObjectController](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSObjectController creates a new NSObjectController instance.
func NewNSObjectController() NSObjectController {
	class := getNSObjectControllerClass()
	rv := objc.Send[NSObjectController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(coder:)
func NewObjectControllerWithCoder(coder foundation.INSCoder) NSObjectController {
	instance := getNSObjectControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSObjectControllerFromID(rv)
}

// Initializes and returns an [NSObjectController] object with the given
// content.
//
// content: The content for the receiver.
//
// # Return Value
// 
// The initialized object controller, with its content object set to
// `content`.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(content:)
func NewObjectControllerWithContent(content objectivec.IObject) NSObjectController {
	instance := getNSObjectControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), content)
	return NSObjectControllerFromID(rv)
}

// Initializes and returns an [NSObjectController] object with the given
// content.
//
// content: The content for the receiver.
//
// # Return Value
// 
// The initialized object controller, with its content object set to
// `content`.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(content:)
func (o NSObjectController) InitWithContent(content objectivec.IObject) NSObjectController {
	rv := objc.Send[NSObjectController](o.ID, objc.Sel("initWithContent:"), content)
	return rv
}
// Typically overridden by subclasses that require additional control over the
// creation of new objects.
//
// # Discussion
// 
// Subclasses that implement this method are responsible for creating the new
// content object and setting it as the receiver’s content object. This
// method is only called if [AutomaticallyPreparesContent] has been set to
// [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/prepareContent()
func (o NSObjectController) PrepareContent() {
	objc.Send[objc.ID](o.ID, objc.Sel("prepareContent"))
}
// Creates and returns a new object of the appropriate class.
//
// # Return Value
// 
// A new object of the appropriate class. The returned object is implicitly
// retained, the sender is responsible for releasing it (with either release
// or autorelease).
//
// # Discussion
// 
// If an entity name is set (see [EntityName]), the object created is an
// instance of the class specified for that entity (and the object is inserted
// into the receiver’s managed object context). Otherwise the object created
// is an instance of the class returned by [ObjectClass].
// 
// This method is called when adding and inserting objects if
// [AutomaticallyPreparesContent] is [true].
// 
// The default implementation assumes the class returned by [ObjectClass] has
// a standard `init` method without arguments. If the object class being
// controlled is [NSManagedObject] (or a subclass thereof) its designated
// initializer ([init(entity:insertInto:)]) is called instead, using the
// entity and managed object context specified for the receiver.
//
// [init(entity:insertInto:)]: https://developer.apple.com/documentation/CoreData/NSManagedObject/init(entity:insertInto:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/newObject()
func (o NSObjectController) NewObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newObject"))
	return objectivec.Object{ID: rv}
}
// Sets the receiver’s content object.
//
// object: The content object for the receiver.
//
// # Discussion
// 
// If the receiver’s content is bound to another (primary) object or
// controller through a relationship key, the relationship of the primary
// object is changed. In a tree-like structure, the object is added after the
// current selection at the same depth. If there is no selection, the object
// is appended to the child nodes of the tree’s arranged objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/addObject(_:)
func (o NSObjectController) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("addObject:"), object)
}
// Removes a given object from the receiver’s content.
//
// object: The object to remove from the receiver.
//
// # Discussion
// 
// If `object` is the receiver’s content object, the receiver’s content is
// set to `nil`. If the receiver’s content is bound to another (primary)
// object or controller through a relationship key, the relationship of the
// primary object is cleared.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/removeObject(_:)
func (o NSObjectController) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("removeObject:"), object)
}
// Creates a new object and sets it as the receiver’s content object.
//
// sender: Typically the object that invoked this method.
//
// # Discussion
// 
// Creates a new object of the appropriate entity (specified by [EntityName])
// or class (specified by [ObjectClass])—see [NewObject]—and sets it as
// the receiver’s content object using [AddObject].
// 
// # Special Considerations
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism can
// provide feedback as a sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/add(_:)
func (o NSObjectController) Add(sender objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("add:"), sender)
}
// Removes the receiver’s content object.
//
// sender: Typically the object that invoked this method.
//
// # Discussion
// 
// Removes the receiver’s content object using [RemoveObject].
// 
// # Special Considerations
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism can
// provide feedback as a sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/remove(_:)
func (o NSObjectController) Remove(sender objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("remove:"), sender)
}
// Causes the receiver to fetch the data objects specified by the entity name
// and fetch predicate.
//
// sender: Typically the object that invoked this method.
//
// # Discussion
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism can
// provide feedback as a sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/fetch(_:)
func (o NSObjectController) Fetch(sender objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("fetch:"), sender)
}
// Returns the default fetch request used by the receiver.
//
// # Return Value
// 
// The default NSFetchResult used by the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/defaultFetchRequest()
func (o NSObjectController) DefaultFetchRequest() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("defaultFetchRequest"))
	return objectivec.Object{ID: rv}
}
// Subclasses should override this method to customize a fetch request, for
// example to specify fetch limits.
//
// fetchRequest: The fetch request to use for the fetch. Pass `nil` to use the default fetch
// request.
//
// merge: If [true], the receiver merges the existing content with the fetch result,
// otherwise the receiver replaces the entire content with the fetch result.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// fetchRequest is a [coredata.NSFetchRequest].
//
// # Discussion
// 
// This method performs a number of actions that you cannot reproduce. To
// customize this method, you should therefore create your own fetch request
// and then invoke `super`’s implementation with the new fetch request.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/fetch(with:merge:)
func (o NSObjectController) FetchWithRequestMergeError(fetchRequest objectivec.IObject, merge bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("fetchWithRequest:merge:error:"), fetchRequest, merge, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("fetchWithRequest:merge:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Returns whether the receiver can handle the action method for a user
// interface item.
//
// item: The user interface item to validate. You can send `item` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if the receiver can handle the action method; [false] if it cannot.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/validateUserInterfaceItem(_:)
func (o NSObjectController) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// The receiver’s content object.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/content
func (o NSObjectController) Content() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("content"))
	return objectivec.Object{ID: rv}
}
func (o NSObjectController) SetContent(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setContent:"), value)
}
// A Boolean that shows whether the receiver automatically creates and inserts
// new content objects automatically when loading from a nib file.
//
// # Discussion
// 
// If `flag` is [true] and the receiver is not using a managed object context,
// [PrepareContent] is used to create the content object. If `flag` is [true]
// and a managed object context is set, the initial content is fetched from
// the managed object context using the current fetch predicate. The default
// is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/automaticallyPreparesContent
func (o NSObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("automaticallyPreparesContent"))
	return rv
}
func (o NSObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutomaticallyPreparesContent:"), value)
}
// The object class to use when creating new objects.
//
// # Discussion
// 
// [NSObjectController]’s default implementation assumes that instances of
// `objectClass` are initialized using a standard `init` method that takes no
// arguments.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/objectClass
func (o NSObjectController) ObjectClass() objc.Class {
	rv := objc.Send[objc.Class](o.ID, objc.Sel("objectClass"))
	return rv
}
func (o NSObjectController) SetObjectClass(value objc.Class) {
	objc.Send[struct{}](o.ID, objc.Sel("setObjectClass:"), value)
}
// A Boolean value that indicates whether an object can be added to the
// receiver using [Add].
//
// # Discussion
// 
// Bindings can use this method to control the enabling of user interface
// objects.
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/canAdd
func (o NSObjectController) CanAdd() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canAdd"))
	return rv
}
// A Boolean value that indicates whether an object can be removed from the
// receiver.
//
// # Discussion
// 
// Bindings can use this method to control the enabling of user interface
// objects.
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/canRemove
func (o NSObjectController) CanRemove() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canRemove"))
	return rv
}
// A Boolean that indicates whether the receiver allows adding and removing
// objects.
//
// # Discussion
// 
// The default is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/isEditable
func (o NSObjectController) Editable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isEditable"))
	return rv
}
func (o NSObjectController) SetEditable(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setEditable:"), value)
}
// The entity name used by the receiver to create new objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/entityName
func (o NSObjectController) EntityName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("entityName"))
	return foundation.NSStringFromID(rv).String()
}
func (o NSObjectController) SetEntityName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setEntityName:"), objc.String(value))
}
// A Boolean that indicates whether the receiver uses lazy fetching.
//
// # Discussion
// 
// When enabled the controller uses a number of techniques that typically make
// managing large data sets more efficient. As with all optimizations, you
// should use suitable performance analysis tools (such as Instruments) to
// determine the best solution.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/usesLazyFetching
func (o NSObjectController) UsesLazyFetching() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("usesLazyFetching"))
	return rv
}
func (o NSObjectController) SetUsesLazyFetching(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setUsesLazyFetching:"), value)
}
// The receiver’s fetch predicate.
//
// # Discussion
// 
// The receiver uses `predicate` when fetching its content, for example in
// [Fetch]. If you need to customize the fetching behavior further, you can
// override [FetchWithRequestMergeError].
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/fetchPredicate
func (o NSObjectController) FetchPredicate() foundation.INSPredicate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fetchPredicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (o NSObjectController) SetFetchPredicate(value foundation.INSPredicate) {
	objc.Send[struct{}](o.ID, objc.Sel("setFetchPredicate:"), value)
}
// The receiver’s managed object context.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/managedObjectContext
func (o NSObjectController) ManagedObjectContext() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("managedObjectContext"))
	return objectivec.Object{ID: rv}
}
func (o NSObjectController) SetManagedObjectContext(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setManagedObjectContext:"), value)
}
// An array of all objects to be affected by editing.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/selectedObjects
func (o NSObjectController) SelectedObjects() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("selectedObjects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// A proxy object representing the receiver’s selection.
//
// # Discussion
// 
// This object is fully key-value coding compliant, but note that it is a
// proxy and so does not provide the full range of functionality that might be
// available in the source object.
//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/selection
func (o NSObjectController) Selection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("selection"))
	return objectivec.Object{ID: rv}
}

