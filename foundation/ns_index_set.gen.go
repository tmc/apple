// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSIndexSet] class.
var (
	_NSIndexSetClass     NSIndexSetClass
	_NSIndexSetClassOnce sync.Once
)

func getNSIndexSetClass() NSIndexSetClass {
	_NSIndexSetClassOnce.Do(func() {
		_NSIndexSetClass = NSIndexSetClass{class: objc.GetClass("NSIndexSet")}
	})
	return _NSIndexSetClass
}

// GetNSIndexSetClass returns the class object for NSIndexSet.
func GetNSIndexSetClass() NSIndexSetClass {
	return getNSIndexSetClass()
}

type NSIndexSetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSIndexSetClass) Alloc() NSIndexSet {
	rv := objc.Send[NSIndexSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An immutable collection of unique integer values that represent indexes in
// another collection.
//
// # Overview
// 
// In Swift, this type bridges to [IndexSet]; use [NSIndexSet] when you need
// reference semantics or other Foundation-specific behavior.
// 
// The [NSIndexSet] class represents an immutable collection of unique
// unsigned integers, known as because of the way they are used. This
// collection is referred to as an . Indexes must be in the range `0 ..
// NSNotFound - 1`.
// 
// You use index sets in your code to store indexes into some other data
// structure. For example, given an [NSArray] object, you could use an index
// set to identify a subset of objects in that array.
// 
// You should not use index sets to store an arbitrary collection of integer
// values because index sets store indexes as sorted ranges. This makes them
// more efficient than storing a collection of individual integers. It also
// means that each index value can only appear once in the index set.
// 
// The designated initializers of the [NSIndexSet] class are: [NSIndexSet.InitWithIndex],
// [NSIndexSet.InitWithIndexesInRange], and [NSIndexSet.InitWithIndexSet].
// 
// You must not subclass the [NSIndexSet] class.
// 
// The mutable subclass of [NSIndexSet] is [NSMutableIndexSet].
//
// [IndexSet]: https://developer.apple.com/documentation/Foundation/IndexSet
//
// # Creating Index Sets
//
//   - [NSIndexSet.InitWithIndex]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index.
//   - [NSIndexSet.InitWithIndexesInRange]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index range.
//   - [NSIndexSet.InitWithIndexSet]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index set.
//
// # Querying Index Sets
//
//   - [NSIndexSet.ContainsIndex]: Indicates whether the index set contains a specific index.
//   - [NSIndexSet.ContainsIndexes]: Indicates whether the receiving index set contains a superset of the indexes in another index set.
//   - [NSIndexSet.ContainsIndexesInRange]: Indicates whether the index set contains the indexes represented by an index range.
//   - [NSIndexSet.IntersectsIndexesInRange]: Indicates whether the index set contains any of the indexes in a range.
//   - [NSIndexSet.Count]: The number of indexes in the index set.
//   - [NSIndexSet.CountOfIndexesInRange]: Returns the number of indexes in the index set that are members of a given range.
//   - [NSIndexSet.IndexPassingTest]: Returns the index of the first object that passes the predicate Block test.
//   - [NSIndexSet.IndexesPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test.
//   - [NSIndexSet.IndexWithOptionsPassingTest]: Returns the index of the first object that passes the predicate Block test using the specified enumeration options.
//   - [NSIndexSet.IndexesWithOptionsPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test using the specified enumeration options.
//   - [NSIndexSet.IndexInRangeOptionsPassingTest]: Returns the index of the first object in the specified range that passes the predicate Block test.
//   - [NSIndexSet.IndexesInRangeOptionsPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects in the specified range that pass the Block test.
//
// # Enumerating Index Set Content
//
//   - [NSIndexSet.EnumerateRangesInRangeOptionsUsingBlock]: Enumerates over the ranges in the range of objects using the block
//   - [NSIndexSet.EnumerateRangesUsingBlock]: Executes a given block using each object in the index set, in the specified ranges.
//   - [NSIndexSet.EnumerateRangesWithOptionsUsingBlock]: Executes a given block using each object in the index set, in the specified ranges.
//
// # Comparing Index Sets
//
//   - [NSIndexSet.IsEqualToIndexSet]: Indicates whether the indexes in the receiving index set are the same indexes contained in another index set.
//
// # Getting Indexes
//
//   - [NSIndexSet.FirstIndex]: The first index in the index set.
//   - [NSIndexSet.LastIndex]: The last index in the index set.
//   - [NSIndexSet.IndexLessThanIndex]: Returns either the closest index in the index set that is less than a specific index or the not-found indicator.
//   - [NSIndexSet.IndexLessThanOrEqualToIndex]: Returns either the closest index in the index set that is less than or equal to a specific index or the not-found indicator.
//   - [NSIndexSet.IndexGreaterThanOrEqualToIndex]: Returns either the closest index in the index set that is greater than or equal to a specific index or the not-found indicator.
//   - [NSIndexSet.IndexGreaterThanIndex]: Returns either the closest index in the index set that is greater than a specific index or the not-found indicator.
//   - [NSIndexSet.GetIndexesMaxCountInIndexRange]: The index set fills an index buffer with the indexes contained both in the index set and in an index range, returning the number of indexes copied.
//
// # Enumerating Indexes
//
//   - [NSIndexSet.EnumerateIndexesUsingBlock]: Executes a given Block using each object in the index set.
//   - [NSIndexSet.EnumerateIndexesWithOptionsUsingBlock]: Executes a given Block over the index set’s indexes, using the specified enumeration options.
//   - [NSIndexSet.EnumerateIndexesInRangeOptionsUsingBlock]: Executes a given Block using the indexes in the specified range, using the specified enumeration options.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet
type NSIndexSet struct {
	objectivec.Object
}

// NSIndexSetFromID constructs a [NSIndexSet] from an objc.ID.
//
// An immutable collection of unique integer values that represent indexes in
// another collection.
func NSIndexSetFromID(id objc.ID) NSIndexSet {
	return NSIndexSet{objectivec.Object{id}}
}
// NOTE: NSIndexSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSIndexSet] class.
//
// # Creating Index Sets
//
//   - [INSIndexSet.InitWithIndex]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index.
//   - [INSIndexSet.InitWithIndexesInRange]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index range.
//   - [INSIndexSet.InitWithIndexSet]: Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index set.
//
// # Querying Index Sets
//
//   - [INSIndexSet.ContainsIndex]: Indicates whether the index set contains a specific index.
//   - [INSIndexSet.ContainsIndexes]: Indicates whether the receiving index set contains a superset of the indexes in another index set.
//   - [INSIndexSet.ContainsIndexesInRange]: Indicates whether the index set contains the indexes represented by an index range.
//   - [INSIndexSet.IntersectsIndexesInRange]: Indicates whether the index set contains any of the indexes in a range.
//   - [INSIndexSet.Count]: The number of indexes in the index set.
//   - [INSIndexSet.CountOfIndexesInRange]: Returns the number of indexes in the index set that are members of a given range.
//   - [INSIndexSet.IndexPassingTest]: Returns the index of the first object that passes the predicate Block test.
//   - [INSIndexSet.IndexesPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test.
//   - [INSIndexSet.IndexWithOptionsPassingTest]: Returns the index of the first object that passes the predicate Block test using the specified enumeration options.
//   - [INSIndexSet.IndexesWithOptionsPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test using the specified enumeration options.
//   - [INSIndexSet.IndexInRangeOptionsPassingTest]: Returns the index of the first object in the specified range that passes the predicate Block test.
//   - [INSIndexSet.IndexesInRangeOptionsPassingTest]: Returns an [NSIndexSet] containing the receiving index set’s objects in the specified range that pass the Block test.
//
// # Enumerating Index Set Content
//
//   - [INSIndexSet.EnumerateRangesInRangeOptionsUsingBlock]: Enumerates over the ranges in the range of objects using the block
//   - [INSIndexSet.EnumerateRangesUsingBlock]: Executes a given block using each object in the index set, in the specified ranges.
//   - [INSIndexSet.EnumerateRangesWithOptionsUsingBlock]: Executes a given block using each object in the index set, in the specified ranges.
//
// # Comparing Index Sets
//
//   - [INSIndexSet.IsEqualToIndexSet]: Indicates whether the indexes in the receiving index set are the same indexes contained in another index set.
//
// # Getting Indexes
//
//   - [INSIndexSet.FirstIndex]: The first index in the index set.
//   - [INSIndexSet.LastIndex]: The last index in the index set.
//   - [INSIndexSet.IndexLessThanIndex]: Returns either the closest index in the index set that is less than a specific index or the not-found indicator.
//   - [INSIndexSet.IndexLessThanOrEqualToIndex]: Returns either the closest index in the index set that is less than or equal to a specific index or the not-found indicator.
//   - [INSIndexSet.IndexGreaterThanOrEqualToIndex]: Returns either the closest index in the index set that is greater than or equal to a specific index or the not-found indicator.
//   - [INSIndexSet.IndexGreaterThanIndex]: Returns either the closest index in the index set that is greater than a specific index or the not-found indicator.
//   - [INSIndexSet.GetIndexesMaxCountInIndexRange]: The index set fills an index buffer with the indexes contained both in the index set and in an index range, returning the number of indexes copied.
//
// # Enumerating Indexes
//
//   - [INSIndexSet.EnumerateIndexesUsingBlock]: Executes a given Block using each object in the index set.
//   - [INSIndexSet.EnumerateIndexesWithOptionsUsingBlock]: Executes a given Block over the index set’s indexes, using the specified enumeration options.
//   - [INSIndexSet.EnumerateIndexesInRangeOptionsUsingBlock]: Executes a given Block using the indexes in the specified range, using the specified enumeration options.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet
type INSIndexSet interface {
	objectivec.IObject
	NSCopying
	NSMutableCopying

	// Topic: Creating Index Sets

	// Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index.
	InitWithIndex(value uint) NSIndexSet
	// Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index range.
	InitWithIndexesInRange(range_ NSRange) NSIndexSet
	// Initializes an allocated [NSIndexSet](<doc://com.apple.foundation/documentation/Foundation/NSIndexSet>) object with an index set.
	InitWithIndexSet(indexSet INSIndexSet) NSIndexSet

	// Topic: Querying Index Sets

	// Indicates whether the index set contains a specific index.
	ContainsIndex(value uint) bool
	// Indicates whether the receiving index set contains a superset of the indexes in another index set.
	ContainsIndexes(indexSet INSIndexSet) bool
	// Indicates whether the index set contains the indexes represented by an index range.
	ContainsIndexesInRange(range_ NSRange) bool
	// Indicates whether the index set contains any of the indexes in a range.
	IntersectsIndexesInRange(range_ NSRange) bool
	// The number of indexes in the index set.
	Count() uint
	// Returns the number of indexes in the index set that are members of a given range.
	CountOfIndexesInRange(range_ NSRange) uint
	// Returns the index of the first object that passes the predicate Block test.
	IndexPassingTest(predicate bool) uint
	// Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test.
	IndexesPassingTest(predicate bool) INSIndexSet
	// Returns the index of the first object that passes the predicate Block test using the specified enumeration options.
	IndexWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) uint
	// Returns an [NSIndexSet] containing the receiving index set’s objects that pass the Block test using the specified enumeration options.
	IndexesWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) INSIndexSet
	// Returns the index of the first object in the specified range that passes the predicate Block test.
	IndexInRangeOptionsPassingTest(range_ NSRange, opts NSEnumerationOptions, predicate bool) uint
	// Returns an [NSIndexSet] containing the receiving index set’s objects in the specified range that pass the Block test.
	IndexesInRangeOptionsPassingTest(range_ NSRange, opts NSEnumerationOptions, predicate bool) INSIndexSet

	// Topic: Enumerating Index Set Content

	// Enumerates over the ranges in the range of objects using the block
	EnumerateRangesInRangeOptionsUsingBlock(range_ NSRange, opts NSEnumerationOptions, block bool)
	// Executes a given block using each object in the index set, in the specified ranges.
	EnumerateRangesUsingBlock(block bool)
	// Executes a given block using each object in the index set, in the specified ranges.
	EnumerateRangesWithOptionsUsingBlock(opts NSEnumerationOptions, block bool)

	// Topic: Comparing Index Sets

	// Indicates whether the indexes in the receiving index set are the same indexes contained in another index set.
	IsEqualToIndexSet(indexSet INSIndexSet) bool

	// Topic: Getting Indexes

	// The first index in the index set.
	FirstIndex() uint
	// The last index in the index set.
	LastIndex() uint
	// Returns either the closest index in the index set that is less than a specific index or the not-found indicator.
	IndexLessThanIndex(value uint) uint
	// Returns either the closest index in the index set that is less than or equal to a specific index or the not-found indicator.
	IndexLessThanOrEqualToIndex(value uint) uint
	// Returns either the closest index in the index set that is greater than or equal to a specific index or the not-found indicator.
	IndexGreaterThanOrEqualToIndex(value uint) uint
	// Returns either the closest index in the index set that is greater than a specific index or the not-found indicator.
	IndexGreaterThanIndex(value uint) uint
	// The index set fills an index buffer with the indexes contained both in the index set and in an index range, returning the number of indexes copied.
	GetIndexesMaxCountInIndexRange(indexBuffer uint, bufferSize uint, range_ NSRangePointer) uint

	// Topic: Enumerating Indexes

	// Executes a given Block using each object in the index set.
	EnumerateIndexesUsingBlock(block bool)
	// Executes a given Block over the index set’s indexes, using the specified enumeration options.
	EnumerateIndexesWithOptionsUsingBlock(opts NSEnumerationOptions, block bool)
	// Executes a given Block using the indexes in the specified range, using the specified enumeration options.
	EnumerateIndexesInRangeOptionsUsingBlock(range_ NSRange, opts NSEnumerationOptions, block bool)

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	InitWithCoder(coder INSCoder) NSIndexSet
}





// Init initializes the instance.
func (i NSIndexSet) Init() NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSIndexSet) Autorelease() NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSIndexSet creates a new NSIndexSet instance.
func NewNSIndexSet() NSIndexSet {
	class := getNSIndexSetClass()
	rv := objc.Send[NSIndexSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewIndexSetWithCoder(coder INSCoder) NSIndexSet {
	instance := getNSIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSIndexSetFromID(rv)
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
func NewIndexSetWithIndex(value uint) NSIndexSet {
	instance := getNSIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndex:"), value)
	return NSIndexSetFromID(rv)
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
func NewIndexSetWithIndexSet(indexSet INSIndexSet) NSIndexSet {
	instance := getNSIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndexSet:"), indexSet)
	return NSIndexSetFromID(rv)
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
func NewIndexSetWithIndexesInRange(range_ NSRange) NSIndexSet {
	instance := getNSIndexSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIndexesInRange:"), range_)
	return NSIndexSetFromID(rv)
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
func (i NSIndexSet) InitWithIndex(value uint) NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("initWithIndex:"), value)
	return rv
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
func (i NSIndexSet) InitWithIndexesInRange(range_ NSRange) NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("initWithIndexesInRange:"), range_)
	return rv
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
func (i NSIndexSet) InitWithIndexSet(indexSet INSIndexSet) NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("initWithIndexSet:"), indexSet)
	return rv
}

// Indicates whether the index set contains a specific index.
//
// value: Index being inquired about.
//
// # Return Value
// 
// [true] when the index set contains `index`, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(_:)-bb19
func (i NSIndexSet) ContainsIndex(value uint) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("containsIndex:"), value)
	return rv
}

// Indicates whether the receiving index set contains a superset of the
// indexes in another index set.
//
// indexSet: Index set being inquired about.
//
// # Return Value
// 
// [true] when the receiving index set contains a superset of the indexes in
// `indexSet`, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(_:)-5j2kh
func (i NSIndexSet) ContainsIndexes(indexSet INSIndexSet) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("containsIndexes:"), indexSet)
	return rv
}

// Indicates whether the index set contains the indexes represented by an
// index range.
//
// range: The index range being inquired about.
//
// # Return Value
// 
// [true] when the index set contains the indexes in `indexRange`, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(in:)
func (i NSIndexSet) ContainsIndexesInRange(range_ NSRange) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("containsIndexesInRange:"), range_)
	return rv
}

// Indicates whether the index set contains any of the indexes in a range.
//
// range: Index range being inquired about.
//
// # Return Value
// 
// [true] when the index set contains one or more of the indexes in
// `indexRange`, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/intersects(in:)
func (i NSIndexSet) IntersectsIndexesInRange(range_ NSRange) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("intersectsIndexesInRange:"), range_)
	return rv
}

// Returns the number of indexes in the index set that are members of a given
// range.
//
// range: Index range being inquired about.
//
// # Return Value
// 
// Number of indexes in the index set that are members of `indexRange`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/countOfIndexes(in:)
func (i NSIndexSet) CountOfIndexesInRange(range_ NSRange) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("countOfIndexesInRange:"), range_)
	return rv
}

// Returns the index of the first object that passes the predicate Block test.
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The index of the first object that passes the predicate test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(passingTest:)
func (i NSIndexSet) IndexPassingTest(predicate bool) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexPassingTest:"), predicate)
	return rv
}

// Returns an [NSIndexSet] containing the receiving index set’s objects that
// pass the Block test.
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSIndexSet] containing the indexes of the receiving index set that
// passed the predicate Block test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(passingTest:)
func (i NSIndexSet) IndexesPassingTest(predicate bool) INSIndexSet {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indexesPassingTest:"), predicate)
	return NSIndexSetFromID(rv)
}

// Returns the index of the first object that passes the predicate Block test
// using the specified enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The index of the first object that passes the predicate test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(options:passingTest:)
func (i NSIndexSet) IndexWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns an [NSIndexSet] containing the receiving index set’s objects that
// pass the Block test using the specified enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSIndexSet] containing the indexes of the receiving index set that
// passed the predicate Block test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(options:passingTest:)
func (i NSIndexSet) IndexesWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) INSIndexSet {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indexesWithOptions:passingTest:"), opts, predicate)
	return NSIndexSetFromID(rv)
}

// Returns the index of the first object in the specified range that passes
// the predicate Block test.
//
// range: The range of indexes to test.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The index of the first object that passes the predicate test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(in:options:passingTest:)
func (i NSIndexSet) IndexInRangeOptionsPassingTest(range_ NSRange, opts NSEnumerationOptions, predicate bool) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

// Returns an [NSIndexSet] containing the receiving index set’s objects in
// the specified range that pass the Block test.
//
// range: The range of indexes to test.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// predicate: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// 
// The Block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSIndexSet] containing the indexes of the receiving index set that
// passed the predicate Block test.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(in:options:passingTest:)
func (i NSIndexSet) IndexesInRangeOptionsPassingTest(range_ NSRange, opts NSEnumerationOptions, predicate bool) INSIndexSet {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indexesInRange:options:passingTest:"), range_, opts, predicate)
	return NSIndexSetFromID(rv)
}

// Enumerates over the ranges in the range of objects using the block
//
// range: The range of items to enumerate. If the range intersects a range of the
// receiver’s indexes, then that intersection will be passed to the block.
//
// opts: A bitmask that specifies the [NSEnumerationOptions] for the enumeration.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// block: The block to apply to elements in the index set.
// 
// The block takes two arguments:
// 
// range: The range of elements. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the array.
// The stop argument is an out-only argument. You should only ever set this
// Boolean to [true] within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the indexed set range to the last object in the range. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(in:options:using:)
func (i NSIndexSet) EnumerateRangesInRangeOptionsUsingBlock(range_ NSRange, opts NSEnumerationOptions, block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateRangesInRange:options:usingBlock:"), range_, opts, block)
}

// Executes a given block using each object in the index set, in the specified
// ranges.
//
// block: The block to apply to elements in the index set.
// 
// The block takes two arguments:
// 
// range: The range of objects of the elements in the index set. stop: A
// reference to a Boolean value. The block can set the value to [true] to stop
// further processing of the array. The stop argument is an out-only argument.
// You should only ever set this Boolean to [true] within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the Block parameter is `nil` this method will raise an exception.
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(_:)
func (i NSIndexSet) EnumerateRangesUsingBlock(block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateRangesUsingBlock:"), block)
}

// Executes a given block using each object in the index set, in the specified
// ranges.
//
// opts: A bitmask that specifies the [NSEnumerationOptions] for the enumeration
// (whether it should be performed concurrently and whether it should be
// performed in reverse order).
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// block: The block to apply to elements in the index set.
// 
// The block takes two arguments:
// 
// range: The range of objects of the elements in the index set. stop: A
// reference to a Boolean value. The block can set the value to [true] to stop
// further processing of the array. The stop argument is an out-only argument.
// You should only ever set this Boolean to [true] within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the indexed set range to the last object in the range. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(options:using:)
func (i NSIndexSet) EnumerateRangesWithOptionsUsingBlock(opts NSEnumerationOptions, block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateRangesWithOptions:usingBlock:"), opts, block)
}

// Indicates whether the indexes in the receiving index set are the same
// indexes contained in another index set.
//
// indexSet: Index set being inquired about.
//
// # Return Value
// 
// [true] when the indexes in the receiving index set are the same indexes
// `indexSet` contains, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/isEqual(to:)
func (i NSIndexSet) IsEqualToIndexSet(indexSet INSIndexSet) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isEqualToIndexSet:"), indexSet)
	return rv
}

// Returns either the closest index in the index set that is less than a
// specific index or the not-found indicator.
//
// value: Index being inquired about.
//
// # Return Value
// 
// Closest index in the index set less than `index`; NSNotFound when the index
// set contains no qualifying index.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexLessThanIndex(_:)
func (i NSIndexSet) IndexLessThanIndex(value uint) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexLessThanIndex:"), value)
	return rv
}

// Returns either the closest index in the index set that is less than or
// equal to a specific index or the not-found indicator.
//
// value: Index being inquired about.
//
// # Return Value
// 
// Closest index in the index set less than or equal to `index`; NSNotFound
// when the index set contains no qualifying index.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexLessThanOrEqual(to:)
func (i NSIndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexLessThanOrEqualToIndex:"), value)
	return rv
}

// Returns either the closest index in the index set that is greater than or
// equal to a specific index or the not-found indicator.
//
// value: Index being inquired about.
//
// # Return Value
// 
// Closest index in the index set greater than or equal to `index`; NSNotFound
// when the index set contains no qualifying index.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexGreaterThanOrEqual(to:)
func (i NSIndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexGreaterThanOrEqualToIndex:"), value)
	return rv
}

// Returns either the closest index in the index set that is greater than a
// specific index or the not-found indicator.
//
// value: Index being inquired about.
//
// # Return Value
// 
// Closest index in the index set greater than `index`; NSNotFound when the
// index set contains no qualifying index.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexGreaterThanIndex(_:)
func (i NSIndexSet) IndexGreaterThanIndex(value uint) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("indexGreaterThanIndex:"), value)
	return rv
}

// The index set fills an index buffer with the indexes contained both in the
// index set and in an index range, returning the number of indexes copied.
//
// indexBuffer: Index buffer to fill.
//
// bufferSize: Maximum size of `indexBuffer`.
//
// range: Index range to compare with indexes in the index set; `nil` represents all
// the indexes in the index set. Indexes in the index range and in the index
// set are copied to `indexBuffer`. On output, the range of indexes not copied
// to `indexBuffer`.
//
// # Return Value
// 
// Number of indexes placed in `indexBuffer`.
//
// # Discussion
// 
// You are responsible for allocating the memory required for `indexBuffer`
// and for releasing it later.
// 
// Suppose you have an index set with contiguous indexes from 1 to 100. If you
// use this method to request a range of `(1, 100)`—which represents the set
// of indexes 1 through 100—and specify a buffer size of `20`, this method
// returns 20 indexes—1 through 20—in `indexBuffer` and sets `indexRange`
// to `(21, 80)`—which represents the indexes 21 through 100.
// 
// Use this method to retrieve entries quickly and efficiently from an index
// set. You can call this method repeatedly to retrieve blocks of index values
// and then process them. When doing so, use the return value and `indexRange`
// to determine when you have finished processing the desired indexes. When
// the return value is less than `bufferSize`, you have reached the end of the
// range.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/getIndexes(_:maxCount:inIndexRange:)
func (i NSIndexSet) GetIndexesMaxCountInIndexRange(indexBuffer uint, bufferSize uint, range_ NSRangePointer) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("getIndexes:maxCount:inIndexRange:"), indexBuffer, bufferSize, range_)
	return rv
}

// Executes a given Block using each object in the index set.
//
// block: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(_:)
func (i NSIndexSet) EnumerateIndexesUsingBlock(block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateIndexesUsingBlock:"), block)
}

// Executes a given Block over the index set’s indexes, using the specified
// enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// block: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(options:using:)
func (i NSIndexSet) EnumerateIndexesWithOptionsUsingBlock(opts NSEnumerationOptions, block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}

// Executes a given Block using the indexes in the specified range, using the
// specified enumeration options.
//
// range: The range to enumerate.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order). See [NSEnumerationOptions] for the supported values.
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// block: The Block to apply to elements in the set.
// 
// The Block takes two arguments:
// 
// idx: The index of the object. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to YES within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(in:options:using:)
func (i NSIndexSet) EnumerateIndexesInRangeOptionsUsingBlock(range_ NSRange, opts NSEnumerationOptions, block bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (i NSIndexSet) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (i NSIndexSet) InitWithCoder(coder INSCoder) NSIndexSet {
	rv := objc.Send[NSIndexSet](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}





// Creates an empty index set.
//
// # Return Value
// 
// [NSIndexSet] object with no members.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSet
func (_NSIndexSetClass NSIndexSetClass) IndexSet() NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexSetClass.class), objc.Sel("indexSet"))
	return NSIndexSetFromID(rv)
}

// Creates an index set with an index.
//
// value: An index. Must be in the range `0 .. NSNotFound - 1`.
//
// # Return Value
// 
// [NSIndexSet] object containing `index`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSetWithIndex:
func (_NSIndexSetClass NSIndexSetClass) IndexSetWithIndex(value uint) NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexSetClass.class), objc.Sel("indexSetWithIndex:"), value)
	return NSIndexSetFromID(rv)
}

// Creates an index set with an index range.
//
// range: An index range. Must be in the range `0 .. NSNotFound - 1`.
//
// # Return Value
// 
// [NSIndexSet] object containing `indexRange`.
//
// # Discussion
// 
// The resulting index set has a [FirstIndex] equal to the `location` of
// `indexRange`, and a [Count] equal to the `length` of `indexRange`.
// Specifying a zero-length range results in an empty index set.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSetWithIndexesInRange:
func (_NSIndexSetClass NSIndexSetClass) IndexSetWithIndexesInRange(range_ NSRange) NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexSetClass.class), objc.Sel("indexSetWithIndexesInRange:"), range_)
	return NSIndexSetFromID(rv)
}








// The number of indexes in the index set.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/count
func (i NSIndexSet) Count() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("count"))
	return rv
}



// The first index in the index set.
//
// # Discussion
// 
// First index in the index set or NSNotFound when the index set is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/firstIndex
func (i NSIndexSet) FirstIndex() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("firstIndex"))
	return rv
}



// The last index in the index set.
//
// # Discussion
// 
// Last index in the index set or NSNotFound when the index set is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSet/lastIndex
func (i NSIndexSet) LastIndex() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("lastIndex"))
	return rv
}
















			// Protocol methods for NSCopying
			




			// Protocol methods for NSMutableCopying
			
















