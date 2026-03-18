// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSArrayController] class.
var (
	_NSArrayControllerClass     NSArrayControllerClass
	_NSArrayControllerClassOnce sync.Once
)

func getNSArrayControllerClass() NSArrayControllerClass {
	_NSArrayControllerClassOnce.Do(func() {
		_NSArrayControllerClass = NSArrayControllerClass{class: objc.GetClass("NSArrayController")}
	})
	return _NSArrayControllerClass
}

// GetNSArrayControllerClass returns the class object for NSArrayController.
func GetNSArrayControllerClass() NSArrayControllerClass {
	return getNSArrayControllerClass()
}

type NSArrayControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSArrayControllerClass) Alloc() NSArrayController {
	rv := objc.Send[NSArrayController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bindings-compatible controller that manages a collection of objects.
//
// # Overview
// 
// Typically the collection that an [NSArrayController] manages is an array,
// however, if the controller manages a relationship of a managed object (see
// [NSManagedObject]) the collection may be a set. [NSArrayController]
// provides selection management and sorting capabilities.
//
// [NSManagedObject]: https://developer.apple.com/documentation/CoreData/NSManagedObject
//
// # Managing Sort Descriptors
//
//   - [NSArrayController.SortDescriptors]: An array of sort descriptor objects, used by the receiver to arrange its content.
//   - [NSArrayController.SetSortDescriptors]
//
// # Arranging Objects
//
//   - [NSArrayController.ArrangeObjects]: Returns a given array, appropriately sorted and filtered.
//   - [NSArrayController.ArrangedObjects]: An array containing the receiver’s content objects arranged using [arrange(_:)](<doc://com.apple.appkit/documentation/AppKit/NSArrayController/arrange(_:)>).
//   - [NSArrayController.RearrangeObjects]: Triggers filtering of the receiver’s content.
//
// # Selection Attributes
//
//   - [NSArrayController.AvoidsEmptySelection]: A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//   - [NSArrayController.SetAvoidsEmptySelection]
//   - [NSArrayController.PreservesSelection]: A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//   - [NSArrayController.SetPreservesSelection]
//   - [NSArrayController.AlwaysUsesMultipleValuesMarker]: A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//   - [NSArrayController.SetAlwaysUsesMultipleValuesMarker]
//
// # Managing selections
//
//   - [NSArrayController.SelectionIndex]: The index of the first object in the receiver’s selection
//   - [NSArrayController.SelectsInsertedObjects]: A Boolean value that indicates whether the receiver automatically selects inserted objects
//   - [NSArrayController.SetSelectsInsertedObjects]
//   - [NSArrayController.SelectionIndexes]: An index set containing the indexes of the receiver’s currently selected objects in the content array
//   - [NSArrayController.AddSelectionIndexes]: Adds the objects at the specified indexes in the receiver’s content array to the current selection.
//   - [NSArrayController.RemoveSelectionIndexes]: Removes the object as the specified indexes from the receiver’s current selection.
//   - [NSArrayController.AddSelectedObjects]: Adds the specified objects from the receiver’s content array to the current selection.
//   - [NSArrayController.RemoveSelectedObjects]: Removes the specified objects from the receiver’s current selection.
//   - [NSArrayController.SelectNext]: Selects the next object, relative to the current selection, in the receiver’s arranged content.
//   - [NSArrayController.CanSelectNext]: A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//   - [NSArrayController.SelectPrevious]: Selects the previous object, relative to the current selection, in the receiver’s arranged content.
//   - [NSArrayController.CanSelectPrevious]: A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// # Inserting
//
//   - [NSArrayController.CanInsert]: Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//   - [NSArrayController.Insert]: Creates a new object and inserts it into the receiver’s content array.
//
// # Adding and Removing Objects
//
//   - [NSArrayController.AddObjects]: Adds `objects` to the receiver’s content collection.
//   - [NSArrayController.InsertObjectAtArrangedObjectIndex]: Inserts `object` into the receiver’s arranged objects array at the location specified by `index`, and adds it to the receiver’s content collection.
//   - [NSArrayController.InsertObjectsAtArrangedObjectIndexes]: Inserts `object`s into the receiver’s arranged objects array at the locations specified in `indexes`, and adds it to the receiver’s content collection.
//   - [NSArrayController.RemoveObjectAtArrangedObjectIndex]: Removes the object at the specified `index` in the receiver’s arranged objects from the receiver’s content array.
//   - [NSArrayController.RemoveObjectsAtArrangedObjectIndexes]: Removes the objects at the specified `indexes` in the receiver’s arranged objects from the content array.
//   - [NSArrayController.RemoveObjects]: Removes `objects` from the receiver’s content collection.
//
// # Filtering Content
//
//   - [NSArrayController.ClearsFilterPredicateOnInsertion]: A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//   - [NSArrayController.SetClearsFilterPredicateOnInsertion]
//   - [NSArrayController.FilterPredicate]: A predicate used by the receiver to filter the array controller contents
//   - [NSArrayController.SetFilterPredicate]
//
// # Automatic Content Rearranging
//
//   - [NSArrayController.AutomaticallyRearrangesObjects]: A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//   - [NSArrayController.SetAutomaticallyRearrangesObjects]
//   - [NSArrayController.AutomaticRearrangementKeyPaths]: An array of key paths that trigger automatic content sorting or filtering
//   - [NSArrayController.DidChangeArrangementCriteria]: Invoked when any criteria for arranging objects change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController
type NSArrayController struct {
	NSObjectController
}

// NSArrayControllerFromID constructs a [NSArrayController] from an objc.ID.
//
// A bindings-compatible controller that manages a collection of objects.
func NSArrayControllerFromID(id objc.ID) NSArrayController {
	return NSArrayController{NSObjectController: NSObjectControllerFromID(id)}
}
// NOTE: NSArrayController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSArrayController] class.
//
// # Managing Sort Descriptors
//
//   - [INSArrayController.SortDescriptors]: An array of sort descriptor objects, used by the receiver to arrange its content.
//   - [INSArrayController.SetSortDescriptors]
//
// # Arranging Objects
//
//   - [INSArrayController.ArrangeObjects]: Returns a given array, appropriately sorted and filtered.
//   - [INSArrayController.ArrangedObjects]: An array containing the receiver’s content objects arranged using [arrange(_:)](<doc://com.apple.appkit/documentation/AppKit/NSArrayController/arrange(_:)>).
//   - [INSArrayController.RearrangeObjects]: Triggers filtering of the receiver’s content.
//
// # Selection Attributes
//
//   - [INSArrayController.AvoidsEmptySelection]: A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//   - [INSArrayController.SetAvoidsEmptySelection]
//   - [INSArrayController.PreservesSelection]: A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//   - [INSArrayController.SetPreservesSelection]
//   - [INSArrayController.AlwaysUsesMultipleValuesMarker]: A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//   - [INSArrayController.SetAlwaysUsesMultipleValuesMarker]
//
// # Managing selections
//
//   - [INSArrayController.SelectionIndex]: The index of the first object in the receiver’s selection
//   - [INSArrayController.SelectsInsertedObjects]: A Boolean value that indicates whether the receiver automatically selects inserted objects
//   - [INSArrayController.SetSelectsInsertedObjects]
//   - [INSArrayController.SelectionIndexes]: An index set containing the indexes of the receiver’s currently selected objects in the content array
//   - [INSArrayController.AddSelectionIndexes]: Adds the objects at the specified indexes in the receiver’s content array to the current selection.
//   - [INSArrayController.RemoveSelectionIndexes]: Removes the object as the specified indexes from the receiver’s current selection.
//   - [INSArrayController.AddSelectedObjects]: Adds the specified objects from the receiver’s content array to the current selection.
//   - [INSArrayController.RemoveSelectedObjects]: Removes the specified objects from the receiver’s current selection.
//   - [INSArrayController.SelectNext]: Selects the next object, relative to the current selection, in the receiver’s arranged content.
//   - [INSArrayController.CanSelectNext]: A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//   - [INSArrayController.SelectPrevious]: Selects the previous object, relative to the current selection, in the receiver’s arranged content.
//   - [INSArrayController.CanSelectPrevious]: A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// # Inserting
//
//   - [INSArrayController.CanInsert]: Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//   - [INSArrayController.Insert]: Creates a new object and inserts it into the receiver’s content array.
//
// # Adding and Removing Objects
//
//   - [INSArrayController.AddObjects]: Adds `objects` to the receiver’s content collection.
//   - [INSArrayController.InsertObjectAtArrangedObjectIndex]: Inserts `object` into the receiver’s arranged objects array at the location specified by `index`, and adds it to the receiver’s content collection.
//   - [INSArrayController.InsertObjectsAtArrangedObjectIndexes]: Inserts `object`s into the receiver’s arranged objects array at the locations specified in `indexes`, and adds it to the receiver’s content collection.
//   - [INSArrayController.RemoveObjectAtArrangedObjectIndex]: Removes the object at the specified `index` in the receiver’s arranged objects from the receiver’s content array.
//   - [INSArrayController.RemoveObjectsAtArrangedObjectIndexes]: Removes the objects at the specified `indexes` in the receiver’s arranged objects from the content array.
//   - [INSArrayController.RemoveObjects]: Removes `objects` from the receiver’s content collection.
//
// # Filtering Content
//
//   - [INSArrayController.ClearsFilterPredicateOnInsertion]: A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//   - [INSArrayController.SetClearsFilterPredicateOnInsertion]
//   - [INSArrayController.FilterPredicate]: A predicate used by the receiver to filter the array controller contents
//   - [INSArrayController.SetFilterPredicate]
//
// # Automatic Content Rearranging
//
//   - [INSArrayController.AutomaticallyRearrangesObjects]: A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//   - [INSArrayController.SetAutomaticallyRearrangesObjects]
//   - [INSArrayController.AutomaticRearrangementKeyPaths]: An array of key paths that trigger automatic content sorting or filtering
//   - [INSArrayController.DidChangeArrangementCriteria]: Invoked when any criteria for arranging objects change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController
type INSArrayController interface {
	INSObjectController

	// Topic: Managing Sort Descriptors

	// An array of sort descriptor objects, used by the receiver to arrange its content.
	SortDescriptors() []foundation.NSSortDescriptor
	SetSortDescriptors(value []foundation.NSSortDescriptor)

	// Topic: Arranging Objects

	// Returns a given array, appropriately sorted and filtered.
	ArrangeObjects(objects foundation.INSArray) foundation.INSArray
	// An array containing the receiver’s content objects arranged using [arrange(_:)](<doc://com.apple.appkit/documentation/AppKit/NSArrayController/arrange(_:)>).
	ArrangedObjects() objectivec.IObject
	// Triggers filtering of the receiver’s content.
	RearrangeObjects()

	// Topic: Selection Attributes

	// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)

	// Topic: Managing selections

	// The index of the first object in the receiver’s selection
	SelectionIndex() uint
	// A Boolean value that indicates whether the receiver automatically selects inserted objects
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	// An index set containing the indexes of the receiver’s currently selected objects in the content array
	SelectionIndexes() foundation.NSIndexSet
	// Adds the objects at the specified indexes in the receiver’s content array to the current selection.
	AddSelectionIndexes(indexes foundation.NSIndexSet) bool
	// Removes the object as the specified indexes from the receiver’s current selection.
	RemoveSelectionIndexes(indexes foundation.NSIndexSet) bool
	// Adds the specified objects from the receiver’s content array to the current selection.
	AddSelectedObjects(objects foundation.INSArray) bool
	// Removes the specified objects from the receiver’s current selection.
	RemoveSelectedObjects(objects foundation.INSArray) bool
	// Selects the next object, relative to the current selection, in the receiver’s arranged content.
	SelectNext(sender objectivec.IObject)
	// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
	CanSelectNext() bool
	// Selects the previous object, relative to the current selection, in the receiver’s arranged content.
	SelectPrevious(sender objectivec.IObject)
	// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
	CanSelectPrevious() bool

	// Topic: Inserting

	// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
	CanInsert() bool
	// Creates a new object and inserts it into the receiver’s content array.
	Insert(sender objectivec.IObject)

	// Topic: Adding and Removing Objects

	// Adds `objects` to the receiver’s content collection.
	AddObjects(objects foundation.INSArray)
	// Inserts `object` into the receiver’s arranged objects array at the location specified by `index`, and adds it to the receiver’s content collection.
	InsertObjectAtArrangedObjectIndex(object objectivec.IObject, index uint)
	// Inserts `object`s into the receiver’s arranged objects array at the locations specified in `indexes`, and adds it to the receiver’s content collection.
	InsertObjectsAtArrangedObjectIndexes(objects foundation.INSArray, indexes foundation.NSIndexSet)
	// Removes the object at the specified `index` in the receiver’s arranged objects from the receiver’s content array.
	RemoveObjectAtArrangedObjectIndex(index uint)
	// Removes the objects at the specified `indexes` in the receiver’s arranged objects from the content array.
	RemoveObjectsAtArrangedObjectIndexes(indexes foundation.NSIndexSet)
	// Removes `objects` from the receiver’s content collection.
	RemoveObjects(objects foundation.INSArray)

	// Topic: Filtering Content

	// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
	ClearsFilterPredicateOnInsertion() bool
	SetClearsFilterPredicateOnInsertion(value bool)
	// A predicate used by the receiver to filter the array controller contents
	FilterPredicate() foundation.INSPredicate
	SetFilterPredicate(value foundation.INSPredicate)

	// Topic: Automatic Content Rearranging

	// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
	AutomaticallyRearrangesObjects() bool
	SetAutomaticallyRearrangesObjects(value bool)
	// An array of key paths that trigger automatic content sorting or filtering
	AutomaticRearrangementKeyPaths() []string
	// Invoked when any criteria for arranging objects change.
	DidChangeArrangementCriteria()
}

// Init initializes the instance.
func (a NSArrayController) Init() NSArrayController {
	rv := objc.Send[NSArrayController](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSArrayController) Autorelease() NSArrayController {
	rv := objc.Send[NSArrayController](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSArrayController creates a new NSArrayController instance.
func NewNSArrayController() NSArrayController {
	class := getNSArrayControllerClass()
	rv := objc.Send[NSArrayController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(coder:)
func NewArrayControllerWithCoder(coder foundation.INSCoder) NSArrayController {
	instance := getNSArrayControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSArrayControllerFromID(rv)
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
func NewArrayControllerWithContent(content objectivec.IObject) NSArrayController {
	instance := getNSArrayControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), content)
	return NSArrayControllerFromID(rv)
}

// Returns a given array, appropriately sorted and filtered.
//
// # Return Value
// 
// An array containing `objects` filtered using the receiver’s filter
// predicate (see [FilterPredicate]) and sorted according to the receiver’s
// [SortDescriptors].
//
// # Discussion
// 
// Subclasses should override this method to use a different sort mechanism,
// provide custom object arrangement, or (typically only prior to OS X version
// 10.4, which provides a filter predicate) filter the objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/arrange(_:)
func (a NSArrayController) ArrangeObjects(objects foundation.INSArray) foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("arrangeObjects:"), objects)
	return foundation.NSArrayFromID(rv)
}

// Triggers filtering of the receiver’s content.
//
// # Discussion
// 
// This method invokes [ArrangeObjects].
// 
// When you detect that filtering criteria change (such as when listening to
// the text sent by an [NSSearchField] instance), invoke this method on
// `self`.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/rearrangeObjects()
func (a NSArrayController) RearrangeObjects() {
	objc.Send[objc.ID](a.ID, objc.Sel("rearrangeObjects"))
}

// Sets the receiver’s selection to the given index, and returns a Boolean
// value that indicates whether the selection was changed.
//
// index: The index for the selection.
//
// # Return Value
// 
// [true] if the selection was changed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/setSelectionIndex(_:)
func (a NSArrayController) SetSelectionIndex(index uint) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setSelectionIndex:"), index)
	return rv
}

// Sets the receiver’s selection indexes and returns a Boolean value that
// indicates whether the selection changed.
//
// indexes: The set of selection indexes for the receiver.
//
// # Return Value
// 
// [true] if the selection was changed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
// 
// To select all the receiver’s objects, indexes should be an index set with
// indexes `[0..XCUIElementTypeCount -1]`. To deselect all indexes, pass an
// empty index set.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/setSelectionIndexes(_:)
func (a NSArrayController) SetSelectionIndexes(indexes foundation.NSIndexSet) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setSelectionIndexes:"), indexes)
	return rv
}

// Adds the objects at the specified indexes in the receiver’s content array
// to the current selection.
//
// # Return Value
// 
// [true] if the selection was changed.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/addSelectionIndexes(_:)
func (a NSArrayController) AddSelectionIndexes(indexes foundation.NSIndexSet) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("addSelectionIndexes:"), indexes)
	return rv
}

// Removes the object as the specified indexes from the receiver’s current
// selection.
//
// # Return Value
// 
// [true] if the selection was changed.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/removeSelectionIndexes(_:)
func (a NSArrayController) RemoveSelectionIndexes(indexes foundation.NSIndexSet) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeSelectionIndexes:"), indexes)
	return rv
}

// Adds the specified objects from the receiver’s content array to the
// current selection.
//
// # Return Value
// 
// [true] if the selection was changed.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/addSelectedObjects(_:)
func (a NSArrayController) AddSelectedObjects(objects foundation.INSArray) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("addSelectedObjects:"), objects)
	return rv
}

// Removes the specified objects from the receiver’s current selection.
//
// # Return Value
// 
// [true] if the selection was changed.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a [CommitEditing] message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/removeSelectedObjects(_:)
func (a NSArrayController) RemoveSelectedObjects(objects foundation.INSArray) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeSelectedObjects:"), objects)
	return rv
}

// Selects the next object, relative to the current selection, in the
// receiver’s arranged content.
//
// # Discussion
// 
// The `sender` is typically the object that invoked this method.
// 
// # Special Considerations
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism (see
// [Error Responders and Error Recovery]) can provide feedback as a sheet.
//
// [Error Responders and Error Recovery]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ErrorHandlingCocoa/ErrorRespondRecover/ErrorRespondRecover.html#//apple_ref/doc/uid/TP40001806-CH203
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/selectNext(_:)
func (a NSArrayController) SelectNext(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("selectNext:"), sender)
}

// Selects the previous object, relative to the current selection, in the
// receiver’s arranged content.
//
// # Discussion
// 
// The `sender` is typically the object that invoked this method.
// 
// # Special Considerations
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism (see
// [Error Responders and Error Recovery]) can provide feedback as a sheet.
//
// [Error Responders and Error Recovery]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ErrorHandlingCocoa/ErrorRespondRecover/ErrorRespondRecover.html#//apple_ref/doc/uid/TP40001806-CH203
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/selectPrevious(_:)
func (a NSArrayController) SelectPrevious(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("selectPrevious:"), sender)
}

// Creates a new object and inserts it into the receiver’s content array.
//
// sender: Typically the object that invoked this method.
//
// # Discussion
// 
// If an entity name is specified (see [EntityName]), this method creates an
// instance of the of the class specified by the entity, otherwise this method
// creates an instance of the class specified by [ObjectClass].
// 
// # Special Considerations
// 
// Beginning with OS X v10.4 the result of this method is deferred until the
// next iteration of the runloop so that the error presentation mechanism (see
// [Error Responders and Error Recovery]) can provide feedback as a sheet.
//
// [Error Responders and Error Recovery]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ErrorHandlingCocoa/ErrorRespondRecover/ErrorRespondRecover.html#//apple_ref/doc/uid/TP40001806-CH203
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(_:)
func (a NSArrayController) Insert(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("insert:"), sender)
}

// Adds `objects` to the receiver’s content collection.
//
// # Discussion
// 
// If [SelectsInsertedObjects] is [true] (the default), the added objects are
// selected in the array controller.
// 
// It is important to note that inserting many objects with
// [SelectsInsertedObjects] on can cause a significant performance penalty. In
// this case it is more efficient to use the [Content] method to set the
// array, or to set [SelectsInsertedObjects] to [false] before adding the
// objects with [AddObjects].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/add(contentsOf:)
func (a NSArrayController) AddObjects(objects foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("addObjects:"), objects)
}

// Inserts `object` into the receiver’s arranged objects array at the
// location specified by `index`, and adds it to the receiver’s content
// collection.
//
// # Discussion
// 
// Subclasses can override this method to provide customized arranged objects
// support. An error is returned if the given index is outside of the
// [ArrangedObjects] range, or if the given object would not appear in the
// arrangedObjects. Set the [ClearsFilterPredicateOnInsertion] to [true] to
// allow insertion.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(_:atArrangedObjectIndex:)
func (a NSArrayController) InsertObjectAtArrangedObjectIndex(object objectivec.IObject, index uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("insertObject:atArrangedObjectIndex:"), object, index)
}

// Inserts `object`s into the receiver’s arranged objects array at the
// locations specified in `indexes`, and adds it to the receiver’s content
// collection.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(contentsOf:atArrangedObjectIndexes:)
func (a NSArrayController) InsertObjectsAtArrangedObjectIndexes(objects foundation.INSArray, indexes foundation.NSIndexSet) {
	objc.Send[objc.ID](a.ID, objc.Sel("insertObjects:atArrangedObjectIndexes:"), objects, indexes)
}

// Removes the object at the specified `index` in the receiver’s arranged
// objects from the receiver’s content array.
//
// # Discussion
// 
// See [RemoveObject] for a discussion of the semantics of removing objects
// when using Core Data.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(atArrangedObjectIndex:)
func (a NSArrayController) RemoveObjectAtArrangedObjectIndex(index uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObjectAtArrangedObjectIndex:"), index)
}

// Removes the objects at the specified `indexes` in the receiver’s arranged
// objects from the content array.
//
// # Discussion
// 
// See [RemoveObject] for a discussion of the semantics of removing objects
// when using Core Data.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(atArrangedObjectIndexes:)
func (a NSArrayController) RemoveObjectsAtArrangedObjectIndexes(indexes foundation.NSIndexSet) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObjectsAtArrangedObjectIndexes:"), indexes)
}

// Removes `objects` from the receiver’s content collection.
//
// # Discussion
// 
// See [RemoveObject] for a discussion of the semantics of removing objects
// when using Core Data.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(contentsOf:)
func (a NSArrayController) RemoveObjects(objects foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObjects:"), objects)
}

// Invoked when any criteria for arranging objects change.
//
// # Discussion
// 
// This method is invoked by the controller itself when any criteria for
// arranging objects change (sort descriptors or filter predicates) to reset
// the key paths for automatic rearranging.
// 
// # Special Considerations
// 
// If you implement a subclass of [NSArrayController] and override
// [RearrangeObjects] to use additional arrangement criteria, you should
// invoke this method if those criteria change.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/didChangeArrangementCriteria()
func (a NSArrayController) DidChangeArrangementCriteria() {
	objc.Send[objc.ID](a.ID, objc.Sel("didChangeArrangementCriteria"))
}

// An array of sort descriptor objects, used by the receiver to arrange its
// content.
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/sortDescriptors
func (a NSArrayController) SortDescriptors() []foundation.NSSortDescriptor {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSSortDescriptor {
		return foundation.NSSortDescriptorFromID(id)
	})
}
func (a NSArrayController) SetSortDescriptors(value []foundation.NSSortDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setSortDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}

// An array containing the receiver’s content objects arranged using
// [ArrangeObjects].
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/arrangedObjects
func (a NSArrayController) ArrangedObjects() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("arrangedObjects"))
	return objectivec.Object{ID: rv}
}

// A Boolean value that indicates whether the receiver requires that the
// content array attempt to maintain a selection
//
// # Discussion
// 
// The default is [true]. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/avoidsEmptySelection
func (a NSArrayController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}
func (a NSArrayController) SetAvoidsEmptySelection(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}

// A Boolean value that indicates whether the receiver will attempt to
// preserve the current selection when the content changes
//
// # Discussion
// 
// The default is [true]. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/preservesSelection
func (a NSArrayController) PreservesSelection() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("preservesSelection"))
	return rv
}
func (a NSArrayController) SetPreservesSelection(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreservesSelection:"), value)
}

// A Boolean value that indicates whether the receiver always returns the
// multiple values marker when multiple objects are selected
//
// # Discussion
// 
// The default is [false]. Setting to [true] can increase performance if your
// application doesn’t allow editing multiple values. This property is
// observable using key-value observing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/alwaysUsesMultipleValuesMarker
func (a NSArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}
func (a NSArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}

// The index of the first object in the receiver’s selection
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/selectionIndex
func (a NSArrayController) SelectionIndex() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("selectionIndex"))
	return rv
}

// A Boolean value that indicates whether the receiver automatically selects
// inserted objects
//
// # Discussion
// 
// The default is [true]. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/selectsInsertedObjects
func (a NSArrayController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}
func (a NSArrayController) SetSelectsInsertedObjects(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}

// An index set containing the indexes of the receiver’s currently selected
// objects in the content array
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/selectionIndexes
func (a NSArrayController) SelectionIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("selectionIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// A Boolean value indicating whether the next object, relative to the current
// selection, in the receiver’s content array can be selected
//
// # Discussion
// 
// This property can be used by a binding to enable user interface items. This
// property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/canSelectNext
func (a NSArrayController) CanSelectNext() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canSelectNext"))
	return rv
}

// A Boolean value indicating whether the previous object, relative to the
// current selection, in the receiver’s content array can be selected
//
// # Discussion
// 
// This property can be used by a binding to enable user interface items. This
// property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/canSelectPrevious
func (a NSArrayController) CanSelectPrevious() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canSelectPrevious"))
	return rv
}

// Returns a Boolean value that indicates whether an object can be inserted
// into the receiver’s content collection.
//
// # Return Value
// 
// [true] if an object can be inserted into the receiver’s content
// collection, otherwise [false].
// 
// # Discussion
// 
// The result of this method can be used by a binding to enable user interface
// items.
// 
// This property is observable using key-value observing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/canInsert
func (a NSArrayController) CanInsert() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canInsert"))
	return rv
}

// A Boolean value that indicates whether the receiver automatically clears an
// existing filter predicate when new items are inserted or added to the
// content
//
// # Discussion
// 
// The default is [true]. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/clearsFilterPredicateOnInsertion
func (a NSArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}
func (a NSArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}

// A predicate used by the receiver to filter the array controller contents
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/filterPredicate
func (a NSArrayController) FilterPredicate() foundation.INSPredicate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("filterPredicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}
func (a NSArrayController) SetFilterPredicate(value foundation.INSPredicate) {
	objc.Send[struct{}](a.ID, objc.Sel("setFilterPredicate:"), value)
}

// A Boolean that indicates if the receiver automatically rearranges its
// content to correspond to the current sort descriptors and filter predicates
//
// # Discussion
// 
// The default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticallyRearrangesObjects
func (a NSArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}
func (a NSArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}

// An array of key paths that trigger automatic content sorting or filtering
//
// # Discussion
// 
// Subclasses can override this property to customize the default behavior of
// the sort descriptors and filtering predicates, for example, if additional
// arrangement criteria are used in a custom implementation of
// [ArrangedObjects].
//
// See: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticRearrangementKeyPaths
func (a NSArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return objc.ConvertSliceToStrings(rv)
}

