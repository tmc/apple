// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTreeNode] class.
var (
	_NSTreeNodeClass     NSTreeNodeClass
	_NSTreeNodeClassOnce sync.Once
)

func getNSTreeNodeClass() NSTreeNodeClass {
	_NSTreeNodeClassOnce.Do(func() {
		_NSTreeNodeClass = NSTreeNodeClass{class: objc.GetClass("NSTreeNode")}
	})
	return _NSTreeNodeClass
}

// GetNSTreeNodeClass returns the class object for NSTreeNode.
func GetNSTreeNodeClass() NSTreeNodeClass {
	return getNSTreeNodeClass()
}

type NSTreeNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTreeNodeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTreeNodeClass) Alloc() NSTreeNode {
	rv := objc.Send[NSTreeNode](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A node in a tree of nodes.
//
// # Overview
// 
// [NSTreeNode] simplifies the creation and management of trees of objects.
// Each tree node represents a model object. A tree node with `nil` as its
// parent node is considered the root of the tree.
//
// # Creating tree nodes
//
//   - [NSTreeNode.InitWithRepresentedObject]: Initializes a newly allocated tree node that represents the specified object.
//
// # Getting information about a node
//
//   - [NSTreeNode.RepresentedObject]: The object the tree node represents.
//   - [NSTreeNode.IndexPath]: The position of the receiver relative to its root parent.
//   - [NSTreeNode.Leaf]: A Boolean that indicates whether the receiver is a leaf node.
//   - [NSTreeNode.ChildNodes]: An array containing receiver’s child nodes.
//   - [NSTreeNode.MutableChildNodes]: A mutable array that provides read-write access to the receiver’s child nodes.
//   - [NSTreeNode.DescendantNodeAtIndexPath]: Returns the receiver’s descendant at the specified index path.
//   - [NSTreeNode.ParentNode]: The receiver’s parent node.
//
// # Sorting the subtree
//
//   - [NSTreeNode.SortWithSortDescriptorsRecursively]: Sorts the receiver’s subtree using the values of the represented objects with the specified sort descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode
type NSTreeNode struct {
	objectivec.Object
}

// NSTreeNodeFromID constructs a [NSTreeNode] from an objc.ID.
//
// A node in a tree of nodes.
func NSTreeNodeFromID(id objc.ID) NSTreeNode {
	return NSTreeNode{objectivec.Object{ID: id}}
}
// NOTE: NSTreeNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTreeNode] class.
//
// # Creating tree nodes
//
//   - [INSTreeNode.InitWithRepresentedObject]: Initializes a newly allocated tree node that represents the specified object.
//
// # Getting information about a node
//
//   - [INSTreeNode.RepresentedObject]: The object the tree node represents.
//   - [INSTreeNode.IndexPath]: The position of the receiver relative to its root parent.
//   - [INSTreeNode.Leaf]: A Boolean that indicates whether the receiver is a leaf node.
//   - [INSTreeNode.ChildNodes]: An array containing receiver’s child nodes.
//   - [INSTreeNode.MutableChildNodes]: A mutable array that provides read-write access to the receiver’s child nodes.
//   - [INSTreeNode.DescendantNodeAtIndexPath]: Returns the receiver’s descendant at the specified index path.
//   - [INSTreeNode.ParentNode]: The receiver’s parent node.
//
// # Sorting the subtree
//
//   - [INSTreeNode.SortWithSortDescriptorsRecursively]: Sorts the receiver’s subtree using the values of the represented objects with the specified sort descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode
type INSTreeNode interface {
	objectivec.IObject

	// Topic: Creating tree nodes

	// Initializes a newly allocated tree node that represents the specified object.
	InitWithRepresentedObject(modelObject objectivec.IObject) NSTreeNode

	// Topic: Getting information about a node

	// The object the tree node represents.
	RepresentedObject() objectivec.IObject
	// The position of the receiver relative to its root parent.
	IndexPath() objc.ID
	// A Boolean that indicates whether the receiver is a leaf node.
	Leaf() bool
	// An array containing receiver’s child nodes.
	ChildNodes() []NSTreeNode
	// A mutable array that provides read-write access to the receiver’s child nodes.
	MutableChildNodes() foundation.INSArray
	// Returns the receiver’s descendant at the specified index path.
	DescendantNodeAtIndexPath(indexPath foundation.INSIndexPath) INSTreeNode
	// The receiver’s parent node.
	ParentNode() INSTreeNode

	// Topic: Sorting the subtree

	// Sorts the receiver’s subtree using the values of the represented objects with the specified sort descriptors.
	SortWithSortDescriptorsRecursively(sortDescriptors []foundation.NSSortDescriptor, recursively bool)
}

// Init initializes the instance.
func (t NSTreeNode) Init() NSTreeNode {
	rv := objc.Send[NSTreeNode](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTreeNode) Autorelease() NSTreeNode {
	rv := objc.Send[NSTreeNode](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTreeNode creates a new NSTreeNode instance.
func NewNSTreeNode() NSTreeNode {
	class := getNSTreeNodeClass()
	rv := objc.Send[NSTreeNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated tree node that represents the specified
// object.
//
// modelObject: The object the tree node represents.
//
// # Return Value
// 
// An initialized tree node that represents `modelObject`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/init(representedObject:)
func NewTreeNodeWithRepresentedObject(modelObject objectivec.IObject) NSTreeNode {
	instance := getNSTreeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRepresentedObject:"), modelObject)
	return NSTreeNodeFromID(rv)
}

// Initializes a newly allocated tree node that represents the specified
// object.
//
// modelObject: The object the tree node represents.
//
// # Return Value
// 
// An initialized tree node that represents `modelObject`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/init(representedObject:)
func (t NSTreeNode) InitWithRepresentedObject(modelObject objectivec.IObject) NSTreeNode {
	rv := objc.Send[NSTreeNode](t.ID, objc.Sel("initWithRepresentedObject:"), modelObject)
	return rv
}
// Returns the receiver’s descendant at the specified index path.
//
// indexPath: An index path specifying a descendant of the receiver.
//
// # Return Value
// 
// A tree node, or `nil` if the node does not exist.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/descendant(at:)
func (t NSTreeNode) DescendantNodeAtIndexPath(indexPath foundation.INSIndexPath) INSTreeNode {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("descendantNodeAtIndexPath:"), indexPath)
	return NSTreeNodeFromID(rv)
}
// Sorts the receiver’s subtree using the values of the represented objects
// with the specified sort descriptors.
//
// sortDescriptors: Array of sort descriptors specifying how to sort the represented objects.
//
// recursively: A Boolean that specifies whether the child nodes should be sorted
// recursively.
//
// # Discussion
// 
// All the represented objects in the child nodes must be key-value coding
// compliant for the keys specified in the sort descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/sort(with:recursively:)
func (t NSTreeNode) SortWithSortDescriptorsRecursively(sortDescriptors []foundation.NSSortDescriptor, recursively bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("sortWithSortDescriptors:recursively:"), objectivec.IObjectSliceToNSArray(sortDescriptors), recursively)
}

// The object the tree node represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/representedObject
func (t NSTreeNode) RepresentedObject() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("representedObject"))
	return objectivec.Object{ID: rv}
}
// The position of the receiver relative to its root parent.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/indexPath
func (t NSTreeNode) IndexPath() objc.ID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("indexPath"))
	return rv
}
// A Boolean that indicates whether the receiver is a leaf node.
//
// # Discussion
// 
// [true] if the receiver is a leaf node (has no child nodes), otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/isLeaf
func (t NSTreeNode) Leaf() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isLeaf"))
	return rv
}
// An array containing receiver’s child nodes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/children
func (t NSTreeNode) ChildNodes() []NSTreeNode {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("childNodes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTreeNode {
		return NSTreeNodeFromID(id)
	})
}
// A mutable array that provides read-write access to the receiver’s child
// nodes.
//
// # Discussion
// 
// Nodes that are inserted into this array have their parent nodes set to the
// receiver. Nodes that are removed from this array automatically have their
// parent node set to `nil`. The array that is returned is observable using
// key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/mutableChildren
func (t NSTreeNode) MutableChildNodes() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("mutableChildNodes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// The receiver’s parent node.
//
// See: https://developer.apple.com/documentation/AppKit/NSTreeNode/parent
func (t NSTreeNode) ParentNode() INSTreeNode {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("parentNode"))
	return NSTreeNodeFromID(objc.ID(rv))
}

