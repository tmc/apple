// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMutableIndexSet] class.
var (
	_NSMutableIndexSetClass     NSMutableIndexSetClass
	_NSMutableIndexSetClassOnce sync.Once
)

func getNSMutableIndexSetClass() NSMutableIndexSetClass {
	_NSMutableIndexSetClassOnce.Do(func() {
		_NSMutableIndexSetClass = NSMutableIndexSetClass{class: objc.GetClass("NSMutableIndexSet")}
	})
	return _NSMutableIndexSetClass
}

// GetNSMutableIndexSetClass returns the class object for NSMutableIndexSet.
func GetNSMutableIndexSetClass() NSMutableIndexSetClass {
	return getNSMutableIndexSetClass()
}

type NSMutableIndexSetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableIndexSetClass) Alloc() NSMutableIndexSet {
	rv := objc.Send[NSMutableIndexSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A mutable collection of unique integer values that represent indexes in
// another collection.
//
// # Overview
// 
// In Swift, this type bridges to [IndexSet]; use [NSMutableIndexSet] when you
// need reference semantics or other Foundation-specific behavior.
// 
// The [NSMutableIndexSet] class represents a mutable collection of unique
// unsigned integers, known as because of the way they are used. This
// collection is referred to as a . The inclusive range of valid indexes is
// `0...(NSNotFound - 1)`; trying to use indexes outside this range is
// invalid.
// 
// The values in a mutable index set are always sorted, so the order in which
// values are added is irrelevant.
// 
// Do not subclass the [NSMutableIndexSet] class.
//
// [IndexSet]: https://developer.apple.com/documentation/Foundation/IndexSet
//
// # Adding Indexes
//
//   - [NSMutableIndexSet.AddIndex]: Adds an index  to the receiver.
//   - [NSMutableIndexSet.AddIndexes]: Adds the indexes in an index set to the receiver.
//   - [NSMutableIndexSet.AddIndexesInRange]: Adds the indexes in an index range to the receiver.
//
// # Removing Indexes
//
//   - [NSMutableIndexSet.RemoveIndex]: Removes an index from the receiver.
//   - [NSMutableIndexSet.RemoveIndexes]: Removes the indexes in an index set from the receiver.
//   - [NSMutableIndexSet.RemoveAllIndexes]: Removes the receiver’s indexes.
//   - [NSMutableIndexSet.RemoveIndexesInRange]: Removes the indexes in an index range from the receiver.
//
// # Shifting Index Groups
//
//   - [NSMutableIndexSet.ShiftIndexesStartingAtIndexBy]: Shifts a group of indexes to the left or the right within the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet
type NSMutableIndexSet struct {
	NSIndexSet
}

// NSMutableIndexSetFromID constructs a [NSMutableIndexSet] from an objc.ID.
//
// A mutable collection of unique integer values that represent indexes in
// another collection.
func NSMutableIndexSetFromID(id objc.ID) NSMutableIndexSet {
	return NSMutableIndexSet{NSIndexSet: NSIndexSetFromID(id)}
}
// NOTE: NSMutableIndexSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMutableIndexSet] class.
//
// # Adding Indexes
//
//   - [INSMutableIndexSet.AddIndex]: Adds an index  to the receiver.
//   - [INSMutableIndexSet.AddIndexes]: Adds the indexes in an index set to the receiver.
//   - [INSMutableIndexSet.AddIndexesInRange]: Adds the indexes in an index range to the receiver.
//
// # Removing Indexes
//
//   - [INSMutableIndexSet.RemoveIndex]: Removes an index from the receiver.
//   - [INSMutableIndexSet.RemoveIndexes]: Removes the indexes in an index set from the receiver.
//   - [INSMutableIndexSet.RemoveAllIndexes]: Removes the receiver’s indexes.
//   - [INSMutableIndexSet.RemoveIndexesInRange]: Removes the indexes in an index range from the receiver.
//
// # Shifting Index Groups
//
//   - [INSMutableIndexSet.ShiftIndexesStartingAtIndexBy]: Shifts a group of indexes to the left or the right within the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet
type INSMutableIndexSet interface {
	INSIndexSet

	// Topic: Adding Indexes

	// Adds an index  to the receiver.
	AddIndex(value uint)
	// Adds the indexes in an index set to the receiver.
	AddIndexes(indexSet INSIndexSet)
	// Adds the indexes in an index range to the receiver.
	AddIndexesInRange(range_ NSRange)

	// Topic: Removing Indexes

	// Removes an index from the receiver.
	RemoveIndex(value uint)
	// Removes the indexes in an index set from the receiver.
	RemoveIndexes(indexSet INSIndexSet)
	// Removes the receiver’s indexes.
	RemoveAllIndexes()
	// Removes the indexes in an index range from the receiver.
	RemoveIndexesInRange(range_ NSRange)

	// Topic: Shifting Index Groups

	// Shifts a group of indexes to the left or the right within the receiver.
	ShiftIndexesStartingAtIndexBy(index uint, delta int)
}





// Init initializes the instance.
func (m NSMutableIndexSet) Init() NSMutableIndexSet {
	rv := objc.Send[NSMutableIndexSet](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableIndexSet) Autorelease() NSMutableIndexSet {
	rv := objc.Send[NSMutableIndexSet](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableIndexSet creates a new NSMutableIndexSet instance.
func NewNSMutableIndexSet() NSMutableIndexSet {
	class := getNSMutableIndexSetClass()
	rv := objc.Send[NSMutableIndexSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMutableIndexSetWithCoder(coder INSCoder) NSMutableIndexSet {
	instance := getNSMutableIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableIndexSetFromID(rv)
}


// Initializes an allocated [NSIndexSet] object with an index.
//
// value: An index. Must be in the range `0 .. NSNotFound - 1`.
//
// # Return Value
// 
// Initialized [NSIndexSet] object with `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(index:)
func NewMutableIndexSetWithIndex(value uint) NSMutableIndexSet {
	instance := getNSMutableIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndex:"), value)
	return NSMutableIndexSetFromID(rv)
}


// Initializes an allocated [NSIndexSet] object with an index set.
//
// indexSet: An index set.
//
// # Return Value
// 
// Initialized [NSIndexSet] object with `indexSet`.
//
// # Discussion
// 
// This method is a designated initializer for [NSIndexSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(indexSet:)
func NewMutableIndexSetWithIndexSet(indexSet INSIndexSet) NSMutableIndexSet {
	instance := getNSMutableIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndexSet:"), indexSet)
	return NSMutableIndexSetFromID(rv)
}


// Initializes an allocated [NSIndexSet] object with an index range.
//
// range: An index range. Must be in the range `0 .. NSNotFound - 1`..
//
// # Return Value
// 
// Initialized [NSIndexSet] object with `indexRange`.
//
// # Discussion
// 
// This method raises an [rangeException] when `indexRange` would add an index
// that exceeds the maximum allowed value for unsigned integers.
// 
// The resulting index set has a [FirstIndex] equal to the `location` of
// `indexRange`, and a [Count] equal to the `length` of `indexRange`.
// Specifying a zero-length range results in an empty index set.
// 
// This method is a designated initializer for [NSIndexSet].
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(indexesIn:)
func NewMutableIndexSetWithIndexesInRange(range_ NSRange) NSMutableIndexSet {
	instance := getNSMutableIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndexesInRange:"), range_)
	return NSMutableIndexSetFromID(rv)
}







// Adds an index to the receiver.
//
// value: Index to add. Must be in the range `0 .. NSNotFound - 1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(_:)-6dtkj
func (m NSMutableIndexSet) AddIndex(value uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("addIndex:"), value)
}

// Adds the indexes in an index set to the receiver.
//
// indexSet: Index set to add.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(_:)-6zmti
func (m NSMutableIndexSet) AddIndexes(indexSet INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("addIndexes:"), indexSet)
}

// Adds the indexes in an index range to the receiver.
//
// range: Index range to add. Must be in the range `0 .. NSNotFound - 1`.
//
// # Discussion
// 
// This method raises an [rangeException] when `range` would add an index that
// exceeds the maximum allowed value for unsigned integers.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(in:)
func (m NSMutableIndexSet) AddIndexesInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("addIndexesInRange:"), range_)
}

// Removes an index from the receiver.
//
// value: Index to remove.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-5li0r
func (m NSMutableIndexSet) RemoveIndex(value uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeIndex:"), value)
}

// Removes the indexes in an index set from the receiver.
//
// indexSet: Index set to remove.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-196u2
func (m NSMutableIndexSet) RemoveIndexes(indexSet INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeIndexes:"), indexSet)
}

// Removes the receiver’s indexes.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/removeAllIndexes()
func (m NSMutableIndexSet) RemoveAllIndexes() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllIndexes"))
}

// Removes the indexes in an index range from the receiver.
//
// range: Index range to remove.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(in:)
func (m NSMutableIndexSet) RemoveIndexesInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeIndexesInRange:"), range_)
}

// Shifts a group of indexes to the left or the right within the receiver.
//
// index: Head of the group of indexes to shift.
//
// delta: Amount and direction of the shift. Positive integers shift the indexes to
// the right. Negative integers shift the indexes to the left.
//
// # Discussion
// 
// The group of indexes shifted is made up by `index` and the indexes that
// follow it in the set.
// 
// A left shift deletes the indexes in a range the length of `delta` preceding
// `index` from the set.
// 
// A right shift inserts empty space in the range `(``index``,``delta``)` in
// the receiver.
// 
// The resulting indexes must all be in the range `0 .. NSNotFound - 1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/shiftIndexesStarting(at:by:)
func (m NSMutableIndexSet) ShiftIndexesStartingAtIndexBy(index uint, delta int) {
	objc.Send[objc.ID](m.ID, objc.Sel("shiftIndexesStartingAtIndex:by:"), index, delta)
}







































