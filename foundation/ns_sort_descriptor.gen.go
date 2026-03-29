// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSortDescriptor] class.
var (
	_NSSortDescriptorClass     NSSortDescriptorClass
	_NSSortDescriptorClassOnce sync.Once
)

func getNSSortDescriptorClass() NSSortDescriptorClass {
	_NSSortDescriptorClassOnce.Do(func() {
		_NSSortDescriptorClass = NSSortDescriptorClass{class: objc.GetClass("NSSortDescriptor")}
	})
	return _NSSortDescriptorClass
}

// GetNSSortDescriptorClass returns the class object for NSSortDescriptor.
func GetNSSortDescriptorClass() NSSortDescriptorClass {
	return getNSSortDescriptorClass()
}

type NSSortDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSortDescriptorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSortDescriptorClass) Alloc() NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An immutable description of how to order a collection of objects according
// to a property common to all the objects.
//
// # Overview
// 
// You construct instances of [NSSortDescriptor] by specifying the key path of
// the property to compare and the order of the sort (ascending or
// descending). Optionally, you can also specify a selector to use to perform
// the comparison, which allows you to specify other comparison selectors,
// such as [LocalizedStandardCompare] and [LocalizedCaseInsensitiveCompare].
// Sorting raises an exception if the objects don’t respond to the sort
// descriptor’s comparison selector.
// 
// You can use sort descriptors for the following:
// 
// - Sorting an array (an instance of [NSArray] or [NSMutableArray] — see
// [SortedArrayUsingDescriptors] and [SortUsingDescriptors]) - Comparing two
// objects directly (see [NSSortDescriptor.CompareObjectToObject]) - Specifying the order of
// objects that return from a Core Data fetch request (see [NSSortDescriptor.SortDescriptors])
//
// # Creating a Sort Descriptor
//
//   - [NSSortDescriptor.InitWithKeyAscending]: Creates a sort descriptor with a specified string key path and sort order.
//   - [NSSortDescriptor.InitWithKeyAscendingSelector]: Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
//   - [NSSortDescriptor.InitWithKeyAscendingComparator]: Creates a sort descriptor with a specified string key path and ordering, and a comparator block.
//
// # Getting Information About a Sort Descriptor
//
//   - [NSSortDescriptor.Ascending]: A Boolean value that indicates whether the receiver specifies sorting in ascending order.
//   - [NSSortDescriptor.Key]: The key that specifies the property to compare during sorting.
//   - [NSSortDescriptor.KeyPath]: The key path that specifies the property to compare during sorting.
//   - [NSSortDescriptor.SetKeyPath]
//   - [NSSortDescriptor.Selector]: The selector for comparing objects.
//   - [NSSortDescriptor.Comparator]: The comparator for the sort descriptor.
//
// # Using Sort Descriptors
//
//   - [NSSortDescriptor.CompareObjectToObject]: Returns a comparison result value that indicates the sort order of two objects.
//   - [NSSortDescriptor.ReversedSortDescriptor]: Returns a sort descriptor that reverses the sort order.
//   - [NSSortDescriptor.AllowEvaluation]: Forces a securely decoded sort descriptor to allow evaluation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor
type NSSortDescriptor struct {
	objectivec.Object
}

// NSSortDescriptorFromID constructs a [NSSortDescriptor] from an objc.ID.
//
// An immutable description of how to order a collection of objects according
// to a property common to all the objects.
func NSSortDescriptorFromID(id objc.ID) NSSortDescriptor {
	return NSSortDescriptor{objectivec.Object{ID: id}}
}
// NOTE: NSSortDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSortDescriptor] class.
//
// # Creating a Sort Descriptor
//
//   - [INSSortDescriptor.InitWithKeyAscending]: Creates a sort descriptor with a specified string key path and sort order.
//   - [INSSortDescriptor.InitWithKeyAscendingSelector]: Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
//   - [INSSortDescriptor.InitWithKeyAscendingComparator]: Creates a sort descriptor with a specified string key path and ordering, and a comparator block.
//
// # Getting Information About a Sort Descriptor
//
//   - [INSSortDescriptor.Ascending]: A Boolean value that indicates whether the receiver specifies sorting in ascending order.
//   - [INSSortDescriptor.Key]: The key that specifies the property to compare during sorting.
//   - [INSSortDescriptor.KeyPath]: The key path that specifies the property to compare during sorting.
//   - [INSSortDescriptor.SetKeyPath]
//   - [INSSortDescriptor.Selector]: The selector for comparing objects.
//   - [INSSortDescriptor.Comparator]: The comparator for the sort descriptor.
//
// # Using Sort Descriptors
//
//   - [INSSortDescriptor.CompareObjectToObject]: Returns a comparison result value that indicates the sort order of two objects.
//   - [INSSortDescriptor.ReversedSortDescriptor]: Returns a sort descriptor that reverses the sort order.
//   - [INSSortDescriptor.AllowEvaluation]: Forces a securely decoded sort descriptor to allow evaluation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor
type INSSortDescriptor interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a Sort Descriptor

	// Creates a sort descriptor with a specified string key path and sort order.
	InitWithKeyAscending(key string, ascending bool) NSSortDescriptor
	// Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
	InitWithKeyAscendingSelector(key string, ascending bool, selector objc.SEL) NSSortDescriptor
	// Creates a sort descriptor with a specified string key path and ordering, and a comparator block.
	InitWithKeyAscendingComparator(key string, ascending bool, cmptr NSComparator) NSSortDescriptor

	// Topic: Getting Information About a Sort Descriptor

	// A Boolean value that indicates whether the receiver specifies sorting in ascending order.
	Ascending() bool
	// The key that specifies the property to compare during sorting.
	Key() string
	// The key path that specifies the property to compare during sorting.
	KeyPath() objectivec.IObject
	SetKeyPath(value objectivec.IObject)
	// The selector for comparing objects.
	Selector() objc.SEL
	// The comparator for the sort descriptor.
	Comparator() NSComparator

	// Topic: Using Sort Descriptors

	// Returns a comparison result value that indicates the sort order of two objects.
	CompareObjectToObject(object1 objectivec.IObject, object2 objectivec.IObject) ComparisonResult
	// Returns a sort descriptor that reverses the sort order.
	ReversedSortDescriptor() objectivec.IObject
	// Forces a securely decoded sort descriptor to allow evaluation.
	AllowEvaluation()

	// The sort descriptors of the fetch request.
	SortDescriptors() INSSortDescriptor
	SetSortDescriptors(value INSSortDescriptor)
}

// Init initializes the instance.
func (s NSSortDescriptor) Init() NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSortDescriptor) Autorelease() NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSortDescriptor creates a new NSSortDescriptor instance.
func NewNSSortDescriptor() NSSortDescriptor {
	class := getNSSortDescriptorClass()
	rv := objc.Send[NSSortDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a sort descriptor by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(coder:)
func NewSortDescriptorWithCoder(coder INSCoder) NSSortDescriptor {
	instance := getNSSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSortDescriptorFromID(rv)
}

// Creates a sort descriptor with a specified string key path and sort order.
//
// key: The key path for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key path and sort
// order, and the default comparison selector ([compare]).
//
// [compare]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/ci_gslang_ext.html#//apple_ref/doc/uid/TP40004397-CH206-BCIECEBI
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:)
func NewSortDescriptorWithKeyAscending(key string, ascending bool) NSSortDescriptor {
	instance := getNSSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:ascending:"), objc.String(key), ascending)
	return NSSortDescriptorFromID(rv)
}

// Creates a sort descriptor with a specified string key path and ordering,
// and a comparator block.
//
// key: The property key for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// cmptr: A comparator block.
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key, ordering, and
// comparator.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:comparator:)
func NewSortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr NSComparator) NSSortDescriptor {
	instance := getNSSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:ascending:comparator:"), objc.String(key), ascending, cmptr)
	return NSSortDescriptorFromID(rv)
}

// Creates a sort descriptor with a specified string key path, ordering, and
// comparison selector.
//
// key: The key path for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// selector: The method to use when comparing the properties of objects, such as
// [LocalizedStandardCompare]. The selector must specify a method that you
// implement according to the value of the property that the key path
// identifies. Pass the selector a single parameter, the object to compare
// against, and it returns the appropriate [ComparisonResult] constant.
// //
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key path, sort order,
// and comparison selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:selector:)
func NewSortDescriptorWithKeyAscendingSelector(key string, ascending bool, selector objc.SEL) NSSortDescriptor {
	instance := getNSSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:ascending:selector:"), objc.String(key), ascending, selector)
	return NSSortDescriptorFromID(rv)
}

// Creates a sort descriptor with a specified string key path and sort order.
//
// key: The key path for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key path and sort
// order, and the default comparison selector ([compare]).
//
// [compare]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/ci_gslang_ext.html#//apple_ref/doc/uid/TP40004397-CH206-BCIECEBI
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:)
func (s NSSortDescriptor) InitWithKeyAscending(key string, ascending bool) NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("initWithKey:ascending:"), objc.String(key), ascending)
	return rv
}
// Creates a sort descriptor with a specified string key path, ordering, and
// comparison selector.
//
// key: The key path for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// selector: The method to use when comparing the properties of objects, such as
// [LocalizedStandardCompare]. The selector must specify a method that you
// implement according to the value of the property that the key path
// identifies. Pass the selector a single parameter, the object to compare
// against, and it returns the appropriate [ComparisonResult] constant.
// //
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key path, sort order,
// and comparison selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:selector:)
func (s NSSortDescriptor) InitWithKeyAscendingSelector(key string, ascending bool, selector objc.SEL) NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("initWithKey:ascending:selector:"), objc.String(key), ascending, selector)
	return rv
}
// Creates a sort descriptor with a specified string key path and ordering,
// and a comparator block.
//
// key: The property key for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// cmptr: A comparator block.
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key, ordering, and
// comparator.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:comparator:)
func (s NSSortDescriptor) InitWithKeyAscendingComparator(key string, ascending bool, cmptr NSComparator) NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("initWithKey:ascending:comparator:"), objc.String(key), ascending, cmptr)
	return rv
}
// Creates a sort descriptor by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(coder:)
func (s NSSortDescriptor) InitWithCoder(coder INSCoder) NSSortDescriptor {
	rv := objc.Send[NSSortDescriptor](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Returns a comparison result value that indicates the sort order of two
// objects.
//
// object1: The object to compare with `object2`. This object must have a property
// accessible using the key-path specified by [Key].
//
// object2: The object to compare with `object1`. This object must have a property
// accessible using the key-path specified by [Key].
//
// # Return Value
// 
// [OrderedAscending] if `object1` is less than `object2`, [OrderedDescending]
// if `object1` is greater than `object2`, or [OrderedSame] if `object1` is
// equal to `object2`.
//
// # Discussion
// 
// The ordering is determined by comparing the values specified by [Key] of
// `object1` and `object2` using the selector specified by [Selector].
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/compare(_:to:)
func (s NSSortDescriptor) CompareObjectToObject(object1 objectivec.IObject, object2 objectivec.IObject) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("compareObject:toObject:"), object1, object2)
	return ComparisonResult(rv)
}
// Forces a securely decoded sort descriptor to allow evaluation.
//
// # Discussion
// 
// When securely decoding [NSSortDescriptor] objects that are encoded using
// [NSSecureCoding], evaluation is disabled because it is potentially unsafe
// to evaluate descriptors you get out of an archive.
// 
// Before you enable evaluation, you should validate key paths, selectors, and
// related properties to ensure no erroneous or malicious code will be
// executed. Once you’ve preflighted the sort descriptor, you can enable the
// sort descriptor for evaluation by calling [AllowEvaluation].
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/allowEvaluation()
func (s NSSortDescriptor) AllowEvaluation() {
	objc.Send[objc.ID](s.ID, objc.Sel("allowEvaluation"))
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSSortDescriptor) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns a sort descriptor with the specified key path and
// ordering.
//
// key: The key path to use when performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order, otherwise
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A sort descriptor initialized with the specified key path and sort order,
// and the default comparison selector ([compare]).
//
// [compare]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/ci_gslang_ext.html#//apple_ref/doc/uid/TP40004397-CH206-BCIECEBI
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:
func (_NSSortDescriptorClass NSSortDescriptorClass) SortDescriptorWithKeyAscending(key string, ascending bool) NSSortDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSSortDescriptorClass.class), objc.Sel("sortDescriptorWithKey:ascending:"), objc.String(key), ascending)
	return NSSortDescriptorFromID(rv)
}
// Creates and returns a sort descriptor initialized with the specified key
// path and ordering, and a comparator block.
//
// key: The property for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// cmptr: A comparator block.
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key, ordering, and
// comparator.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:comparator:
func (_NSSortDescriptorClass NSSortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr NSComparator) NSSortDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSSortDescriptorClass.class), objc.Sel("sortDescriptorWithKey:ascending:comparator:"), objc.String(key), ascending, cmptr)
	return NSSortDescriptorFromID(rv)
}
// Creates a sort descriptor with the specified key path, ordering, and
// comparison selector.
//
// key: The key path for performing a comparison.
// 
// For information about key paths, see [Key-Value Coding Programming Guide].
// //
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
//
// ascending: [true] if the receiver specifies sorting in ascending order; otherwise,
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// selector: The method to use when comparing the properties of objects, for example
// [LocalizedStandardCompare]. The selector must specify a method implemented
// by the value of the property identified by the key path. The selector used
// for the comparison is passed a single parameter, the object to compare
// against, and it returns the appropriate [ComparisonResult] constant.
// //
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
//
// # Return Value
// 
// A sort descriptor that initializes with the specified key path, sort order,
// and comparison selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:selector:
func (_NSSortDescriptorClass NSSortDescriptorClass) SortDescriptorWithKeyAscendingSelector(key string, ascending bool, selector objc.SEL) NSSortDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSSortDescriptorClass.class), objc.Sel("sortDescriptorWithKey:ascending:selector:"), objc.String(key), ascending, selector)
	return NSSortDescriptorFromID(rv)
}

// A Boolean value that indicates whether the receiver specifies sorting in
// ascending order.
//
// # Discussion
// 
// [true] if the receiver specifies sorting in ascending order, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/ascending
func (s NSSortDescriptor) Ascending() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ascending"))
	return rv
}
// The key that specifies the property to compare during sorting.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/key
func (s NSSortDescriptor) Key() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("key"))
	return NSStringFromID(rv).String()
}
// The key path that specifies the property to compare during sorting.
//
// See: https://developer.apple.com/documentation/foundation/nssortdescriptor/keypath
func (s NSSortDescriptor) KeyPath() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("keyPath"))
	return objectivec.Object{ID: rv}
}
func (s NSSortDescriptor) SetKeyPath(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setKeyPath:"), value)
}
// The selector for comparing objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/selector
func (s NSSortDescriptor) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("selector"))
	return rv
}
// The comparator for the sort descriptor.
//
// # Discussion
// 
// Call this property only for sort descriptors initialized with
// [InitWithKeyAscendingComparator].
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/comparator
func (s NSSortDescriptor) Comparator() NSComparator {
	rv := objc.Send[NSComparator](s.ID, objc.Sel("comparator"))
	return NSComparator(rv)
}
// Returns a sort descriptor that reverses the sort order.
//
// See: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/reversedSortDescriptor
func (s NSSortDescriptor) ReversedSortDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("reversedSortDescriptor"))
	return objectivec.Object{ID: rv}
}
// The sort descriptors of the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (s NSSortDescriptor) SortDescriptors() INSSortDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sortDescriptors"))
	return NSSortDescriptorFromID(objc.ID(rv))
}
func (s NSSortDescriptor) SetSortDescriptors(value INSSortDescriptor) {
	objc.Send[struct{}](s.ID, objc.Sel("setSortDescriptors:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

