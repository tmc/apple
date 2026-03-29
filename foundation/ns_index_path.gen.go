// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSIndexPath] class.
var (
	_NSIndexPathClass     NSIndexPathClass
	_NSIndexPathClassOnce sync.Once
)

func getNSIndexPathClass() NSIndexPathClass {
	_NSIndexPathClassOnce.Do(func() {
		_NSIndexPathClass = NSIndexPathClass{class: objc.GetClass("NSIndexPath")}
	})
	return _NSIndexPathClass
}

// GetNSIndexPathClass returns the class object for NSIndexPath.
func GetNSIndexPathClass() NSIndexPathClass {
	return getNSIndexPathClass()
}

type NSIndexPathClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSIndexPathClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSIndexPathClass) Alloc() NSIndexPath {
	rv := objc.Send[NSIndexPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A list of indexes that together represent the path to a specific location
// in a tree of nested arrays.
//
// # Overview
// 
// In Swift, this object bridges to [IndexPath]; use [NSIndexPath] when you
// need reference semantics or other Foundation-specific behavior.
// 
// Each index in an index path represents the index into an array of children
// from one node in the tree to another, deeper, node. For example, the index
// path `1.4.3.2` specifies the path shown in [Figure 1].
// 
// [media-1965825]
//
// [IndexPath]: https://developer.apple.com/documentation/Foundation/IndexPath
//
// # Creating and Initializing Index Paths
//
//   - [NSIndexPath.InitWithIndex]: Initializes an index path with a single node.
//   - [NSIndexPath.InitWithIndexesLength]: Initializes an index path with the given nodes and length.
//
// # Using Special Node Names
//
//   - [NSIndexPath.Section]: An index number identifying a section in a table view or collection view.
//   - [NSIndexPath.Item]: An index number identifying an item in a section of a collection view.
//
// # Counting Nodes
//
//   - [NSIndexPath.Length]: The number of nodes in the index path.
//
// # Adding and Removing Nodes
//
//   - [NSIndexPath.IndexPathByAddingIndex]: Returns an index path containing the nodes in the receiving index path plus another given index.
//   - [NSIndexPath.IndexPathByRemovingLastIndex]: Returns an index path with the nodes in the receiving index path, excluding the last one.
//
// # Comparing Index Paths
//
//   - [NSIndexPath.Compare]: Indicates the depth-first traversal order of the receiving index path and another index path.
//
// # Working with Indexes
//
//   - [NSIndexPath.IndexAtPosition]: Provides the value at a particular node in the index path.
//   - [NSIndexPath.GetIndexesRange]: Copies the indexes stored in the index path from the positions specified by the position range into the specified indexes.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath
type NSIndexPath struct {
	objectivec.Object
}

// NSIndexPathFromID constructs a [NSIndexPath] from an objc.ID.
//
// A list of indexes that together represent the path to a specific location
// in a tree of nested arrays.
func NSIndexPathFromID(id objc.ID) NSIndexPath {
	return NSIndexPath{objectivec.Object{ID: id}}
}
// NOTE: NSIndexPath adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSIndexPath] class.
//
// # Creating and Initializing Index Paths
//
//   - [INSIndexPath.InitWithIndex]: Initializes an index path with a single node.
//   - [INSIndexPath.InitWithIndexesLength]: Initializes an index path with the given nodes and length.
//
// # Using Special Node Names
//
//   - [INSIndexPath.Section]: An index number identifying a section in a table view or collection view.
//   - [INSIndexPath.Item]: An index number identifying an item in a section of a collection view.
//
// # Counting Nodes
//
//   - [INSIndexPath.Length]: The number of nodes in the index path.
//
// # Adding and Removing Nodes
//
//   - [INSIndexPath.IndexPathByAddingIndex]: Returns an index path containing the nodes in the receiving index path plus another given index.
//   - [INSIndexPath.IndexPathByRemovingLastIndex]: Returns an index path with the nodes in the receiving index path, excluding the last one.
//
// # Comparing Index Paths
//
//   - [INSIndexPath.Compare]: Indicates the depth-first traversal order of the receiving index path and another index path.
//
// # Working with Indexes
//
//   - [INSIndexPath.IndexAtPosition]: Provides the value at a particular node in the index path.
//   - [INSIndexPath.GetIndexesRange]: Copies the indexes stored in the index path from the positions specified by the position range into the specified indexes.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath
type INSIndexPath interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating and Initializing Index Paths

	// Initializes an index path with a single node.
	InitWithIndex(index uint) NSIndexPath
	// Initializes an index path with the given nodes and length.
	InitWithIndexesLength(indexes uint, length uint) NSIndexPath

	// Topic: Using Special Node Names

	// An index number identifying a section in a table view or collection view.
	Section() int
	// An index number identifying an item in a section of a collection view.
	Item() int

	// Topic: Counting Nodes

	// The number of nodes in the index path.
	Length() uint

	// Topic: Adding and Removing Nodes

	// Returns an index path containing the nodes in the receiving index path plus another given index.
	IndexPathByAddingIndex(index uint) objc.ID
	// Returns an index path with the nodes in the receiving index path, excluding the last one.
	IndexPathByRemovingLastIndex() objc.ID

	// Topic: Comparing Index Paths

	// Indicates the depth-first traversal order of the receiving index path and another index path.
	Compare(otherObject objectivec.IObject) ComparisonResult

	// Topic: Working with Indexes

	// Provides the value at a particular node in the index path.
	IndexAtPosition(position uint) uint
	// Copies the indexes stored in the index path from the positions specified by the position range into the specified indexes.
	GetIndexesRange(indexes unsafe.Pointer, positionRange NSRange)
}

// Init initializes the instance.
func (i NSIndexPath) Init() NSIndexPath {
	rv := objc.Send[NSIndexPath](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSIndexPath) Autorelease() NSIndexPath {
	rv := objc.Send[NSIndexPath](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSIndexPath creates a new NSIndexPath instance.
func NewNSIndexPath() NSIndexPath {
	class := getNSIndexPathClass()
	rv := objc.Send[NSIndexPath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an index path with the indexes of a specific item and section
// in a collection view.
//
// item: An index number identifying an item in a [UICollectionView] object in a
// section identified by the `section` parameter.
// //
// [UICollectionView]: https://developer.apple.com/documentation/UIKit/UICollectionView
//
// section: An index number identifying a section in a [UICollectionView] object.
// //
// [UICollectionView]: https://developer.apple.com/documentation/UIKit/UICollectionView
//
// # Return Value
// 
// An [NSIndexPath] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forItem:inSection:)
func NewIndexPathForItemInSection(item int, section int) NSIndexPath {
	rv := objc.Send[objc.ID](objc.ID(getNSIndexPathClass().class), objc.Sel("indexPathForItem:inSection:"), item, section)
	return NSIndexPathFromID(rv)
}

// Initializes an index path with the indexes of a specific row and section in
// a table view.
//
// row: An index number identifying a row in a [UITableView] object in a section
// identified by `section`.
// //
// [UITableView]: https://developer.apple.com/documentation/UIKit/UITableView
//
// section: An index number identifying a section in a [UITableView] object.
// //
// [UITableView]: https://developer.apple.com/documentation/UIKit/UITableView
//
// # Return Value
// 
// An [NSIndexPath] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
func NewIndexPathForRowInSection(row int, section int) NSIndexPath {
	rv := objc.Send[objc.ID](objc.ID(getNSIndexPathClass().class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return NSIndexPathFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewIndexPathWithCoder(coder INSCoder) NSIndexPath {
	instance := getNSIndexPathClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSIndexPathFromID(rv)
}

// Initializes an index path with a single node.
//
// index: Index of the item in node 0 to point to.
//
// # Return Value
// 
// Initialized [NSIndexPath] object representing a one-node index path with
// `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(index:)
func NewIndexPathWithIndex(index uint) NSIndexPath {
	instance := getNSIndexPathClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndex:"), index)
	return NSIndexPathFromID(rv)
}

// Initializes an index path with the given nodes and length.
//
// indexes: Array of indexes to make up the index path.
//
// length: Number of nodes to include in the index path.
//
// # Return Value
// 
// Initialized [NSIndexPath] object with `indexes` up to `length`.
//
// # Discussion
// 
// This method is a designated initializer of [NSIndexPath].
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(indexes:length:)
func NewIndexPathWithIndexesLength(indexes uint, length uint) NSIndexPath {
	instance := getNSIndexPathClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndexes:length:"), indexes, length)
	return NSIndexPathFromID(rv)
}

// Initializes an index path with a single node.
//
// index: Index of the item in node 0 to point to.
//
// # Return Value
// 
// Initialized [NSIndexPath] object representing a one-node index path with
// `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(index:)
func (i NSIndexPath) InitWithIndex(index uint) NSIndexPath {
	rv := objc.Send[NSIndexPath](i.ID, objc.Sel("initWithIndex:"), index)
	return rv
}
// Initializes an index path with the given nodes and length.
//
// indexes: Array of indexes to make up the index path.
//
// length: Number of nodes to include in the index path.
//
// # Return Value
// 
// Initialized [NSIndexPath] object with `indexes` up to `length`.
//
// # Discussion
// 
// This method is a designated initializer of [NSIndexPath].
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(indexes:length:)
func (i NSIndexPath) InitWithIndexesLength(indexes uint, length uint) NSIndexPath {
	rv := objc.Send[NSIndexPath](i.ID, objc.Sel("initWithIndexes:length:"), indexes, length)
	return rv
}
// Returns an index path containing the nodes in the receiving index path plus
// another given index.
//
// index: Index to append to the index path’s indexes.
//
// # Return Value
// 
// A new index path containing the receiving index path’s indexes and
// `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/adding(_:)
func (i NSIndexPath) IndexPathByAddingIndex(index uint) objc.ID {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indexPathByAddingIndex:"), index)
	return rv
}
// Returns an index path with the nodes in the receiving index path, excluding
// the last one.
//
// # Return Value
// 
// A new index path with the receiving index path’s indexes, excluding the
// last one.
//
// # Discussion
// 
// Returns an empty [NSIndexPath] instance if the receiving index path’s
// length is 1 or less.
// 
// # Special Considerations
// 
// In OS X v10.4 this method returns `nil` when the length of the receiving
// index path is 1 or less. On iOS and macOS 10.5 and later this method never
// returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/removingLastIndex()
func (i NSIndexPath) IndexPathByRemovingLastIndex() objc.ID {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indexPathByRemovingLastIndex"))
	return rv
}
// Indicates the depth-first traversal order of the receiving index path and
// another index path.
//
// otherObject: Index path to compare.
// 
// This value must not be `nil`. If the value is `nil`, the behavior is
// undefined.
//
// # Return Value
// 
// The depth-first traversal ordering of the receiving index path and
// `indexPath`.
//
// # Discussion
// 
// - [OrderedAscending]: The receiving index path comes before `indexPath`. -
// [OrderedDescending]: The receiving index path comes after `indexPath`. -
// [OrderedSame]: The receiving index path and `indexPath` are the same index
// path.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/compare(_:)
func (i NSIndexPath) Compare(otherObject objectivec.IObject) ComparisonResult {
	rv := objc.Send[ComparisonResult](i.ID, objc.Sel("compare:"), otherObject)
	return ComparisonResult(rv)
}
// Provides the value at a particular node in the index path.
//
// position: Index value of the desired node. Node numbering starts at zero.
//
// # Return Value
// 
// The index value at `node` or [NSNotFound] if the node is outside the range
// of the index path.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/index(atPosition:)
func (i NSIndexPath) IndexAtPosition(position uint) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexAtPosition:"), position)
	return rv
}
// Copies the indexes stored in the index path from the positions specified by
// the position range into the specified indexes.
//
// indexes: Pointer to a C array of at least as many [NSUInteger] objects as specified
// by the length of `positionRange`. On return, the array holds the index
// path’s indexes.
// //
// [NSUInteger]: https://developer.apple.com/documentation/ObjectiveC/NSUInteger
//
// positionRange: A range of valid positions within the index path. If the location plus the
// length of `positionRange` is greater than the length of the index path,
// this method raises an [rangeException].
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Discussion
// 
// You must allocate the memory for the C array.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/getIndexes(_:range:)
func (i NSIndexPath) GetIndexesRange(indexes unsafe.Pointer, positionRange NSRange) {
	objc.Send[objc.ID](i.ID, objc.Sel("getIndexes:range:"), indexes, positionRange)
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (i NSIndexPath) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (i NSIndexPath) InitWithCoder(coder INSCoder) NSIndexPath {
	rv := objc.Send[NSIndexPath](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates a one-node index path.
//
// index: Index of the item in node 0 to point to.
//
// # Return Value
// 
// One-node index path with `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/indexPathWithIndex:
func (_NSIndexPathClass NSIndexPathClass) IndexPathWithIndex(index uint) NSIndexPath {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexPathClass.class), objc.Sel("indexPathWithIndex:"), index)
	return NSIndexPathFromID(rv)
}
// Creates an index path with one or more nodes.
//
// indexes: Array of indexes to make up the index path.
//
// length: Number of nodes to include in the index path.
//
// # Return Value
// 
// Index path with `indexes` up to `length`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/indexPathWithIndexes:length:
func (_NSIndexPathClass NSIndexPathClass) IndexPathWithIndexesLength(indexes uint, length uint) NSIndexPath {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexPathClass.class), objc.Sel("indexPathWithIndexes:length:"), indexes, length)
	return NSIndexPathFromID(rv)
}

// An index number identifying a section in a table view or collection view.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/section
func (i NSIndexPath) Section() int {
	rv := objc.Send[int](i.ID, objc.Sel("section"))
	return rv
}
// An index number identifying an item in a section of a collection view.
//
// # Discussion
// 
// The section the item is in is identified by the value of [Section].
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/item
func (i NSIndexPath) Item() int {
	rv := objc.Send[int](i.ID, objc.Sel("item"))
	return rv
}
// The number of nodes in the index path.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/length
func (i NSIndexPath) Length() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("length"))
	return rv
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

