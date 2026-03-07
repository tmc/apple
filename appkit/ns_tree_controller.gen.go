// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTreeController] class.
var (
	_NSTreeControllerClass     NSTreeControllerClass
	_NSTreeControllerClassOnce sync.Once
)

func getNSTreeControllerClass() NSTreeControllerClass {
	_NSTreeControllerClassOnce.Do(func() {
		_NSTreeControllerClass = NSTreeControllerClass{class: objc.GetClass("NSTreeController")}
	})
	return _NSTreeControllerClass
}

// GetNSTreeControllerClass returns the class object for NSTreeController.
func GetNSTreeControllerClass() NSTreeControllerClass {
	return getNSTreeControllerClass()
}

type NSTreeControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTreeControllerClass) Alloc() NSTreeController {
	rv := objc.Send[NSTreeController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A bindings-compatible controller that manages a tree of objects.
//
// # Overview
// 
// The [NSTreeController] class provides selection and sort management. Its
// primary purpose is to act as the controller when binding [NSOutlineView]
// and [NSBrowser] instances to a hierarchical collection of objects. The root
// content object of the tree can be a single object, or an array of objects.
// 
// An [NSTreeController] object requires that you describe how the tree of
// objects is traversed by specifying the key-path for child objects specified
// by [NSTreeController.ChildrenKeyPath]. All child objects for the tree must be key-value
// coding compliant for the same child key path. If necessary, you should add
// properties to your model classes that map the child key name to the
// appropriate class-specific property name.
// 
// Child objects can implement a count method (specified to the tree
// controller using [NSTreeController.CountKeyPath]) that, if provided, returns the number of
// child objects available. Your model objects are expected to update the
// value of the count key path in a key-value observing compliant method.
// Optionally, you can also provide a leaf key path using [NSTreeController.LeafKeyPath] that
// specifies a key in your model object that returns [true] if the object is a
// leaf node, and [false] if it is not. Changes to the leaf node value of the
// child object should be made in a key-value observing compliant manner.
// Providing the leaf node key path can improve performance, because it
// prevents the [NSTreeController] from having to examine the child object to
// determine if it is a leaf node.
// 
// For more information about using NSTreeController in your app, see
// [Navigating Hierarchical Data Using Outline and Split Views].
//
// [Navigating Hierarchical Data Using Outline and Split Views]: https://developer.apple.com/documentation/AppKit/navigating-hierarchical-data-using-outline-and-split-views
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Managing Sort Descriptors
//
//   - [NSTreeController.SortDescriptors]: An array containing the sort descriptors used to arrange the tree controller’s content.
//   - [NSTreeController.SetSortDescriptors]
//
// # Arranging Objects
//
//   - [NSTreeController.ArrangedObjects]: The tree controller’s sorted content objects.
//   - [NSTreeController.RearrangeObjects]: Use this method to trigger reordering of the tree controller’s content.
//
// # Getting the current selection
//
//   - [NSTreeController.SelectionIndexPath]: The index path of the first selected object.
//   - [NSTreeController.SelectionIndexPaths]: An array containing the index paths of the currently selected objects.
//   - [NSTreeController.SelectedNodes]: An array containing the tree controller’s selected tree nodes.
//
// # Managing Selections
//
//   - [NSTreeController.SelectsInsertedObjects]: A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//   - [NSTreeController.SetSelectsInsertedObjects]
//   - [NSTreeController.AddSelectionIndexPaths]: Adds the objects at the specified `indexPaths` in the tree controller’s content to the current selection.
//   - [NSTreeController.RemoveSelectionIndexPaths]: Removes the objects at the specified index paths from the tree controller’s current selection.
//   - [NSTreeController.AvoidsEmptySelection]: A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//   - [NSTreeController.SetAvoidsEmptySelection]
//   - [NSTreeController.PreservesSelection]: A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//   - [NSTreeController.SetPreservesSelection]
//   - [NSTreeController.AlwaysUsesMultipleValuesMarker]: A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//   - [NSTreeController.SetAlwaysUsesMultipleValuesMarker]
//
// # Adding, inserting and removing objects
//
//   - [NSTreeController.AddChild]: Adds a child object to the currently selected item.
//   - [NSTreeController.CanAddChild]: A Boolean value that indicates if a child object can be added to the tree controller’s content.
//   - [NSTreeController.CanInsert]: A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//   - [NSTreeController.CanInsertChild]: A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//   - [NSTreeController.Insert]: Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content.
//   - [NSTreeController.InsertChild]: Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content as a child of the current selection.
//   - [NSTreeController.InsertObjectAtArrangedObjectIndexPath]: Inserts `object` into the tree controller’s arranged objects array at the location specified by `indexPath`, and adds it to the tree controller’s content.
//   - [NSTreeController.InsertObjectsAtArrangedObjectIndexPaths]: Inserts `objects` into the tree controller’s arranged objects array at the locations specified in `indexPaths`, and adds them to the tree controller’s content.
//   - [NSTreeController.RemoveObjectAtArrangedObjectIndexPath]: Removes the object at the specified `indexPath` in the tree controller’s arranged objects from the tree controller’s content.
//   - [NSTreeController.RemoveObjectsAtArrangedObjectIndexPaths]: Removes the objects at the specified `indexPaths` in the tree controller’s arranged objects from the tree controller’s content.
//   - [NSTreeController.MoveNodeToIndexPath]: Moves the specified tree node to the new index path.
//   - [NSTreeController.MoveNodesToIndexPath]: Moves the specified tree nodes to the new index path.
//
// # Specifying model attributes
//
//   - [NSTreeController.ChildrenKeyPath]: The key path used to find the children in the tree controller’s objects.
//   - [NSTreeController.SetChildrenKeyPath]
//   - [NSTreeController.ChildrenKeyPathForNode]: Returns the key path used to find the children in the specified tree node.
//   - [NSTreeController.CountKeyPath]: The key path used to find the number of children for a node.
//   - [NSTreeController.SetCountKeyPath]
//   - [NSTreeController.CountKeyPathForNode]: Returns the key path that provides the number of children for a specified node.
//   - [NSTreeController.LeafKeyPath]: The key path used by the tree controller to determine if a node is a leaf key.
//   - [NSTreeController.SetLeafKeyPath]
//   - [NSTreeController.LeafKeyPathForNode]: Returns the key path that specifies whether the node is a leaf node.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController
type NSTreeController struct {
	NSObjectController
}

// NSTreeControllerFromID constructs a [NSTreeController] from an objc.ID.
//
// A bindings-compatible controller that manages a tree of objects.
func NSTreeControllerFromID(id objc.ID) NSTreeController {
	return NSTreeController{NSObjectController: NSObjectControllerFromID(id)}
}
// NOTE: NSTreeController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTreeController] class.
//
// # Managing Sort Descriptors
//
//   - [INSTreeController.SortDescriptors]: An array containing the sort descriptors used to arrange the tree controller’s content.
//   - [INSTreeController.SetSortDescriptors]
//
// # Arranging Objects
//
//   - [INSTreeController.ArrangedObjects]: The tree controller’s sorted content objects.
//   - [INSTreeController.RearrangeObjects]: Use this method to trigger reordering of the tree controller’s content.
//
// # Getting the current selection
//
//   - [INSTreeController.SelectionIndexPath]: The index path of the first selected object.
//   - [INSTreeController.SelectionIndexPaths]: An array containing the index paths of the currently selected objects.
//   - [INSTreeController.SelectedNodes]: An array containing the tree controller’s selected tree nodes.
//
// # Managing Selections
//
//   - [INSTreeController.SelectsInsertedObjects]: A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//   - [INSTreeController.SetSelectsInsertedObjects]
//   - [INSTreeController.AddSelectionIndexPaths]: Adds the objects at the specified `indexPaths` in the tree controller’s content to the current selection.
//   - [INSTreeController.RemoveSelectionIndexPaths]: Removes the objects at the specified index paths from the tree controller’s current selection.
//   - [INSTreeController.AvoidsEmptySelection]: A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//   - [INSTreeController.SetAvoidsEmptySelection]
//   - [INSTreeController.PreservesSelection]: A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//   - [INSTreeController.SetPreservesSelection]
//   - [INSTreeController.AlwaysUsesMultipleValuesMarker]: A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//   - [INSTreeController.SetAlwaysUsesMultipleValuesMarker]
//
// # Adding, inserting and removing objects
//
//   - [INSTreeController.AddChild]: Adds a child object to the currently selected item.
//   - [INSTreeController.CanAddChild]: A Boolean value that indicates if a child object can be added to the tree controller’s content.
//   - [INSTreeController.CanInsert]: A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//   - [INSTreeController.CanInsertChild]: A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//   - [INSTreeController.Insert]: Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content.
//   - [INSTreeController.InsertChild]: Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content as a child of the current selection.
//   - [INSTreeController.InsertObjectAtArrangedObjectIndexPath]: Inserts `object` into the tree controller’s arranged objects array at the location specified by `indexPath`, and adds it to the tree controller’s content.
//   - [INSTreeController.InsertObjectsAtArrangedObjectIndexPaths]: Inserts `objects` into the tree controller’s arranged objects array at the locations specified in `indexPaths`, and adds them to the tree controller’s content.
//   - [INSTreeController.RemoveObjectAtArrangedObjectIndexPath]: Removes the object at the specified `indexPath` in the tree controller’s arranged objects from the tree controller’s content.
//   - [INSTreeController.RemoveObjectsAtArrangedObjectIndexPaths]: Removes the objects at the specified `indexPaths` in the tree controller’s arranged objects from the tree controller’s content.
//   - [INSTreeController.MoveNodeToIndexPath]: Moves the specified tree node to the new index path.
//   - [INSTreeController.MoveNodesToIndexPath]: Moves the specified tree nodes to the new index path.
//
// # Specifying model attributes
//
//   - [INSTreeController.ChildrenKeyPath]: The key path used to find the children in the tree controller’s objects.
//   - [INSTreeController.SetChildrenKeyPath]
//   - [INSTreeController.ChildrenKeyPathForNode]: Returns the key path used to find the children in the specified tree node.
//   - [INSTreeController.CountKeyPath]: The key path used to find the number of children for a node.
//   - [INSTreeController.SetCountKeyPath]
//   - [INSTreeController.CountKeyPathForNode]: Returns the key path that provides the number of children for a specified node.
//   - [INSTreeController.LeafKeyPath]: The key path used by the tree controller to determine if a node is a leaf key.
//   - [INSTreeController.SetLeafKeyPath]
//   - [INSTreeController.LeafKeyPathForNode]: Returns the key path that specifies whether the node is a leaf node.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController
type INSTreeController interface {
	INSObjectController

	// Topic: Managing Sort Descriptors

	// An array containing the sort descriptors used to arrange the tree controller’s content.
	SortDescriptors() []foundation.NSSortDescriptor
	SetSortDescriptors(value []foundation.NSSortDescriptor)

	// Topic: Arranging Objects

	// The tree controller’s sorted content objects.
	ArrangedObjects() INSTreeNode
	// Use this method to trigger reordering of the tree controller’s content.
	RearrangeObjects()

	// Topic: Getting the current selection

	// The index path of the first selected object.
	SelectionIndexPath() objc.ID
	// An array containing the index paths of the currently selected objects.
	SelectionIndexPaths() []objc.ID
	// An array containing the tree controller’s selected tree nodes.
	SelectedNodes() []NSTreeNode

	// Topic: Managing Selections

	// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	// Adds the objects at the specified `indexPaths` in the tree controller’s content to the current selection.
	AddSelectionIndexPaths(indexPaths []objc.ID) bool
	// Removes the objects at the specified index paths from the tree controller’s current selection.
	RemoveSelectionIndexPaths(indexPaths []objc.ID) bool
	// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)

	// Topic: Adding, inserting and removing objects

	// Adds a child object to the currently selected item.
	AddChild(sender objectivec.IObject)
	// A Boolean value that indicates if a child object can be added to the tree controller’s content.
	CanAddChild() bool
	// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
	CanInsert() bool
	// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
	CanInsertChild() bool
	// Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content.
	Insert(sender objectivec.IObject)
	// Creates a new object of the class specified by `objectClass` and inserts it into the tree controller’s content as a child of the current selection.
	InsertChild(sender objectivec.IObject)
	// Inserts `object` into the tree controller’s arranged objects array at the location specified by `indexPath`, and adds it to the tree controller’s content.
	InsertObjectAtArrangedObjectIndexPath(object objectivec.IObject, indexPath objectivec.IObject)
	// Inserts `objects` into the tree controller’s arranged objects array at the locations specified in `indexPaths`, and adds them to the tree controller’s content.
	InsertObjectsAtArrangedObjectIndexPaths(objects foundation.INSArray, indexPaths []objc.ID)
	// Removes the object at the specified `indexPath` in the tree controller’s arranged objects from the tree controller’s content.
	RemoveObjectAtArrangedObjectIndexPath(indexPath objectivec.IObject)
	// Removes the objects at the specified `indexPaths` in the tree controller’s arranged objects from the tree controller’s content.
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []objc.ID)
	// Moves the specified tree node to the new index path.
	MoveNodeToIndexPath(node INSTreeNode, indexPath objectivec.IObject)
	// Moves the specified tree nodes to the new index path.
	MoveNodesToIndexPath(nodes []NSTreeNode, startingIndexPath objectivec.IObject)

	// Topic: Specifying model attributes

	// The key path used to find the children in the tree controller’s objects.
	ChildrenKeyPath() string
	SetChildrenKeyPath(value string)
	// Returns the key path used to find the children in the specified tree node.
	ChildrenKeyPathForNode(node INSTreeNode) string
	// The key path used to find the number of children for a node.
	CountKeyPath() string
	SetCountKeyPath(value string)
	// Returns the key path that provides the number of children for a specified node.
	CountKeyPathForNode(node INSTreeNode) string
	// The key path used by the tree controller to determine if a node is a leaf key.
	LeafKeyPath() string
	SetLeafKeyPath(value string)
	// Returns the key path that specifies whether the node is a leaf node.
	LeafKeyPathForNode(node INSTreeNode) string

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTreeController) Init() NSTreeController {
	rv := objc.Send[NSTreeController](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTreeController) Autorelease() NSTreeController {
	rv := objc.Send[NSTreeController](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTreeController creates a new NSTreeController instance.
func NewNSTreeController() NSTreeController {
	class := getNSTreeControllerClass()
	rv := objc.Send[NSTreeController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSObjectController/init(coder:)
func NewTreeControllerWithCoder(coder foundation.INSCoder) NSTreeController {
	instance := getNSTreeControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTreeControllerFromID(rv)
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
func NewTreeControllerWithContent(content objectivec.IObject) NSTreeController {
	instance := getNSTreeControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), content)
	return NSTreeControllerFromID(rv)
}







// Use this method to trigger reordering of the tree controller’s content.
//
// # Discussion
// 
// Subclasses should invoke this method if any parameter that affects the
// arranged objects changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/rearrangeObjects()
func (t NSTreeController) RearrangeObjects() {
	objc.Send[objc.ID](t.ID, objc.Sel("rearrangeObjects"))
}

// Sets the tree controller’s current selection.
//
// indexPath: The proposed new selection.
//
// # Return Value
// 
// Return [true] if the selection has changed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a `commitEditing` message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/setSelectionIndexPath(_:)
func (t NSTreeController) SetSelectionIndexPath(indexPath objectivec.IObject) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("setSelectionIndexPath:"), indexPath)
	return rv
}

// Sets the tree controller’s current selection to the specified index
// paths.
//
// indexPaths: An array of [NSIndexPath] objects specifying the selected objects.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Return Value
// 
// Return [true] if the selection has changed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a `commitEditing` message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/setSelectionIndexPaths(_:)
func (t NSTreeController) SetSelectionIndexPaths(indexPaths []objc.ID) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("setSelectionIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return rv
}

// Adds the objects at the specified `indexPaths` in the tree controller’s
// content to the current selection.
//
// # Discussion
// 
// Attempting to change the selection may cause a `commitEditing` message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/addSelectionIndexPaths(_:)
func (t NSTreeController) AddSelectionIndexPaths(indexPaths []objc.ID) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("addSelectionIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return rv
}

// Removes the objects at the specified index paths from the tree
// controller’s current selection.
//
// # Return Value
// 
// [true] if the selection was changed.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Attempting to change the selection may cause a `commitEditing` message
// which fails, thus denying the selection change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/removeSelectionIndexPaths(_:)
func (t NSTreeController) RemoveSelectionIndexPaths(indexPaths []objc.ID) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("removeSelectionIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return rv
}

// Adds a child object to the currently selected item.
//
// # Discussion
// 
// The `sender` is typically the object that invoked this method.
// 
// If the receiver is in object mode `newObject` is called and the returned
// object is added as a child. If the receiver is in entity mode a new object
// is created that is appropriate for the relationship as specified by the
// entity, and `newObject` is not used.
// 
// # Special Considerations
// 
// The result of this method is deferred until the next iteration of the run
// loop so that the error presentation mechanism can provide feedback as a
// sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/addChild(_:)
func (t NSTreeController) AddChild(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addChild:"), sender)
}

// Creates a new object of the class specified by `objectClass` and inserts it
// into the tree controller’s content.
//
// # Discussion
// 
// The `sender` is typically the object that invoked this method.
// 
// If the receiver is in object mode `newObject` is called and the returned
// object is inserted into the collection. If the receiver is in entity mode a
// new object is created that is appropriate as specified by the entity, and
// `newObject` is not used.
// 
// # Special Considerations
// 
// The result of this method is deferred until the next iteration of the run
// loop so that the error presentation mechanism can provide feedback as a
// sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:)
func (t NSTreeController) Insert(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("insert:"), sender)
}

// Creates a new object of the class specified by `objectClass` and inserts it
// into the tree controller’s content as a child of the current selection.
//
// # Discussion
// 
// The `sender` is typically the object that invoked this method.
// 
// If the receiver is in object mode `newObject` is called and the returned
// object is inserted as a child. If the receiver is in entity mode a new
// object is created that is appropriate for the relationship as specified by
// the entity, and `newObject` is not used.
// 
// # Special Considerations
// 
// The result of this method is deferred until the next iteration of the run
// loop so that the error presentation mechanism can provide feedback as a
// sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/insertChild(_:)
func (t NSTreeController) InsertChild(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertChild:"), sender)
}

// Inserts `object` into the tree controller’s arranged objects array at the
// location specified by `indexPath`, and adds it to the tree controller’s
// content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPath:)
func (t NSTreeController) InsertObjectAtArrangedObjectIndexPath(object objectivec.IObject, indexPath objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertObject:atArrangedObjectIndexPath:"), object, indexPath)
}

// Inserts `objects` into the tree controller’s arranged objects array at
// the locations specified in `indexPaths`, and adds them to the tree
// controller’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPaths:)
func (t NSTreeController) InsertObjectsAtArrangedObjectIndexPaths(objects foundation.INSArray, indexPaths []objc.ID) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertObjects:atArrangedObjectIndexPaths:"), objects, objectivec.IDSliceToNSArray(indexPaths))
}

// Removes the object at the specified `indexPath` in the tree controller’s
// arranged objects from the tree controller’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObject(atArrangedObjectIndexPath:)
func (t NSTreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeObjectAtArrangedObjectIndexPath:"), indexPath)
}

// Removes the objects at the specified `indexPaths` in the tree
// controller’s arranged objects from the tree controller’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObjects(atArrangedObjectIndexPaths:)
func (t NSTreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []objc.ID) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeObjectsAtArrangedObjectIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
}

// Moves the specified tree node to the new index path.
//
// node: A tree node.
//
// indexPath: An index path specifying the new position in the tree controller’s
// content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-s5xp
func (t NSTreeController) MoveNodeToIndexPath(node INSTreeNode, indexPath objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("moveNode:toIndexPath:"), node, indexPath)
}

// Moves the specified tree nodes to the new index path.
//
// nodes: An array of tree nodes.
//
// startingIndexPath: An index path specifying the starting position to move the tree nodes to in
// the tree controller’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-moi9
func (t NSTreeController) MoveNodesToIndexPath(nodes []NSTreeNode, startingIndexPath objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("moveNodes:toIndexPath:"), objectivec.IObjectSliceToNSArray(nodes), startingIndexPath)
}

// Returns the key path used to find the children in the specified tree node.
//
// node: A tree node in the tree controller’s content.
//
// # Return Value
// 
// A string containing the key path in `node` that provides the child nodes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath(for:)
func (t NSTreeController) ChildrenKeyPathForNode(node INSTreeNode) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("childrenKeyPathForNode:"), node)
	return foundation.NSStringFromID(rv).String()
}

// Returns the key path that provides the number of children for a specified
// node.
//
// node: A tree node in the tree controller’s content.
//
// # Return Value
// 
// A string containing the key path in `node` that provides the number of
// children.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath(for:)
func (t NSTreeController) CountKeyPathForNode(node INSTreeNode) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("countKeyPathForNode:"), node)
	return foundation.NSStringFromID(rv).String()
}

// Returns the key path that specifies whether the node is a leaf node.
//
// node: A tree node in the tree controller’s content.
//
// # Return Value
// 
// A string containing the key path in `node` that specifies that the node is
// a leaf node.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath(for:)
func (t NSTreeController) LeafKeyPathForNode(node INSTreeNode) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("leafKeyPathForNode:"), node)
	return foundation.NSStringFromID(rv).String()
}
func (t NSTreeController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// An array containing the sort descriptors used to arrange the tree
// controller’s content.
//
// # Discussion
// 
// When the value of this property is an empty array, the tree controller has
// no sort descriptors configured, which means that the contents are arranged
// in their natural order. This property is observable using key-value
// observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/sortDescriptors
func (t NSTreeController) SortDescriptors() []foundation.NSSortDescriptor {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("sortDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSSortDescriptor {
		return foundation.NSSortDescriptorFromID(id)
	})
}
func (t NSTreeController) SetSortDescriptors(value []foundation.NSSortDescriptor) {
	objc.Send[struct{}](t.ID, objc.Sel("setSortDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}



// The tree controller’s sorted content objects.
//
// # Discussion
// 
// The value of this property represents a proxy root tree node containing the
// tree controller’s sorted content objects. The proxy object responds to
// [ChildNodes] and [DescendantNodeAtIndexPath] messages. This property is
// observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/arrangedObjects
func (t NSTreeController) ArrangedObjects() INSTreeNode {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("arrangedObjects"))
	return NSTreeNodeFromID(objc.ID(rv))
}



// The index path of the first selected object.
//
// # Discussion
// 
// The value of this property is `nil` if there is no selection. This property
// is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPath
func (t NSTreeController) SelectionIndexPath() objc.ID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectionIndexPath"))
	return rv
}



// An array containing the index paths of the currently selected objects.
//
// # Discussion
// 
// This property contains an array containing [NSIndexPath] objects for each
// of the selected objects in the tree controller’s content. This property
// is observable using key-value observing.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPaths
func (t NSTreeController) SelectionIndexPaths() []objc.ID {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("selectionIndexPaths"))
	return rv
}



// An array containing the tree controller’s selected tree nodes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/selectedNodes
func (t NSTreeController) SelectedNodes() []NSTreeNode {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("selectedNodes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTreeNode {
		return NSTreeNodeFromID(id)
	})
}



// A Boolean value that indicates whether the tree controller automatically
// selects objects as they are inserted.
//
// # Discussion
// 
// The default value of this property is [true]. This property is observable
// using key-value observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/selectsInsertedObjects
func (t NSTreeController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}
func (t NSTreeController) SetSelectsInsertedObjects(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}



// A Boolean value that indicates whether the tree controller requires the
// content array to attempt to maintain a selection at all times, avoiding an
// empty selection.
//
// # Discussion
// 
// When the value of this property is [true], the tree controller maintains a
// selection unless there are no objects in the content. The default value is
// [true]. This property is observable using key-value observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/avoidsEmptySelection
func (t NSTreeController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}
func (t NSTreeController) SetAvoidsEmptySelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}



// A Boolean value that indicates whether the tree controller will attempt to
// preserve the current selection when the content changes.
//
// # Discussion
// 
// When the value of this property is [true], the selection is preserved, if
// possible. The default value is [true]. This property is observable using
// key-value observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/preservesSelection
func (t NSTreeController) PreservesSelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("preservesSelection"))
	return rv
}
func (t NSTreeController) SetPreservesSelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreservesSelection:"), value)
}



// A Boolean value that indicates whether the tree controller always returns
// the multiple values marker when multiple objects are selected, even if the
// selected items have the same value.
//
// # Discussion
// 
// Setting this property to [true] can increase performance if your
// application doesn’t allow editing multiple values. The default is
// [false]. This property is observable using key-value observing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/alwaysUsesMultipleValuesMarker
func (t NSTreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}
func (t NSTreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}



// A Boolean value that indicates if a child object can be added to the tree
// controller’s content.
//
// # Discussion
// 
// The value of this property is [true] if a child object can be added to the
// tree controller’s content. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/canAddChild
func (t NSTreeController) CanAddChild() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canAddChild"))
	return rv
}



// A Boolean value that indicates if an object can be inserted into the tree
// controller’s content.
//
// # Discussion
// 
// The value of this property is [true] if an object can be inserted into the
// tree controller’s content. This property is observable using key-value
// observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsert
func (t NSTreeController) CanInsert() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canInsert"))
	return rv
}



// A Boolean value that indicates if a child object can be inserted into the
// tree controller’s content.
//
// # Discussion
// 
// The value of this property is [true] if a child object can be inserted into
// the tree controller’s content. This property is observable using
// key-value observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsertChild
func (t NSTreeController) CanInsertChild() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canInsertChild"))
	return rv
}



// The key path used to find the children in the tree controller’s objects.
//
// # Discussion
// 
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath
func (t NSTreeController) ChildrenKeyPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("childrenKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTreeController) SetChildrenKeyPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setChildrenKeyPath:"), objc.String(value))
}



// The key path used to find the number of children for a node.
//
// # Discussion
// 
// Specifying this key path (if the data is available in the model object) can
// increase performance, but disables insert and remove functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath
func (t NSTreeController) CountKeyPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("countKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTreeController) SetCountKeyPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCountKeyPath:"), objc.String(value))
}



// The key path used by the tree controller to determine if a node is a leaf
// key.
//
// # Discussion
// 
// Specifying a key path for this property is optional. If the tree controller
// is able to determine that a node is a leaf node, it can disable inserting
// or adding children to those nodes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath
func (t NSTreeController) LeafKeyPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("leafKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTreeController) SetLeafKeyPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLeafKeyPath:"), objc.String(value))
}




























