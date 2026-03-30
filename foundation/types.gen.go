// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
)

// C struct types

// NSAffineTransformStruct - A structure that defines the three-by-three matrix that performs an affine transform between two coordinate systems.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransformStruct
type NSAffineTransformStruct struct {
	M11 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	M12 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	M21 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	M22 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	TX  float64 // An element of the transform matrix that contributes translation.
	TY  float64 // An element of the transform matrix that contributes translation.

}

// NSDecimal - A structure representing a base-10 number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Decimal
type NSDecimal struct {
	Sign      unsafe.Pointer // The sign of the decimal.
	Exponent  int            // The exponent of the decimal.
	_mantissa unsafe.Pointer
}

// NSEdgeInsets - A description of the distance between the edges of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEdgeInsets
type NSEdgeInsets struct {
	Top    float64 // The distance from the top of the source rectangle to the top of the result rectangle.
	Left   float64 // The distance from the left side of the source rectangle to the left side of the result rectangle.
	Bottom float64 // The distance from the bottom of the source rectangle to the bottom of the result rectangle.
	Right  float64 // The distance from the right side of the source rectangle to the right side of the result rectangle.

}

// NSFastEnumerationState - This defines the structure used as contextual information in the [NSFastEnumeration](<doc://com.apple.foundation/documentation/Foundation/NSFastEnumeration>) protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFastEnumerationState
type NSFastEnumerationState struct {
	State        uint           // Arbitrary state information used by the iterator. Typically this is set to `0` at the beginning of the iteration.
	MutationsPtr *uint          // Arbitrary state information used to detect whether the collection has been mutated.
	Extra        uint           // A C array that you can use to hold returned values.
	ItemsPtr     unsafe.Pointer // A C array of objects.

}

// NSHashEnumerator - Allows successive elements of a hash table to be returned each time this structure is passed to [NSNextHashEnumeratorItem(_:)](<doc://com.apple.foundation/documentation/Foundation/NSNextHashEnumeratorItem(_:)>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashEnumerator
type NSHashEnumerator struct {
}

// NSHashTableCallBacks - Defines a structure that contains the function pointers used to configure behavior of [NSHashTable] with respect to elements within a hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTableCallBacks
type NSHashTableCallBacks struct {
	Describe func(*NSHashTable, unsafe.Pointer) NSString             // Points to the function that produces an autoreleased NSString * describing the given element. If [NULL], then the hash table produces a generic string description.
	Hash     func(*NSHashTable, unsafe.Pointer) uint                 // Points to the function that must produce hash code for elements of the hash table. If [NULL], the pointer value is used as the hash code. Second parameter is the element for which hash code should be produced.
	IsEqual  func(*NSHashTable, unsafe.Pointer, unsafe.Pointer) bool // Points to the function that compares second and third parameters. If [NULL], then == is used for comparison.
	Release  func(*NSHashTable, unsafe.Pointer)                      // Points to the function that decrements a reference count for the given element, and if the reference count becomes 0, frees the given element. If [NULL], then nothing is done for reference counting or releasing.
	Retain   func(*NSHashTable, unsafe.Pointer)                      // Points to the function that increments a reference count for the given element. If [NULL], then nothing is done for reference counting.

}

// NSMapEnumerator - Allows successive elements of a map table to be returned each time this structure is passed to [NSNextMapEnumeratorPair(_:_:_:)](<doc://com.apple.foundation/documentation/Foundation/NSNextMapEnumeratorPair(_:_:_:)>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapEnumerator
type NSMapEnumerator struct {
}

// NSMapTableKeyCallBacks - The function pointers used to configure behavior of [NSMapTable] with respect to key elements within a map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTableKeyCallBacks
type NSMapTableKeyCallBacks struct {
	NotAKeyMarker unsafe.Pointer                                         // No key put in map table can be this value. An exception is raised if attempt is made to use this value as a key
	Describe      func(*NSMapTable, unsafe.Pointer) NSString             // Points to the function which produces an autoreleased NSString * describing the given element. If [NULL], then the map table produces a generic string description.
	Hash          func(*NSMapTable, unsafe.Pointer) uint                 // Points to the function which must produce hash code for key elements of the map table. If [NULL], the pointer value is used as the hash code. Second parameter is the element for which hash code should be produced.
	IsEqual       func(*NSMapTable, unsafe.Pointer, unsafe.Pointer) bool // Points to the function which compares second and third parameters. If [NULL], then == is used for comparison.
	Release       func(*NSMapTable, unsafe.Pointer)                      // Points to the function which decrements a reference count for the given element, and if the reference count becomes zero, frees the given element. If [NULL], then nothing is done for reference counting or releasing.
	Retain        func(*NSMapTable, unsafe.Pointer)                      // Points to the function which increments a reference count for the given element. If [NULL], then nothing is done for reference counting.

}

// NSMapTableValueCallBacks - The function pointers used to configure behavior of [NSMapTable] with respect to value elements within a map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTableValueCallBacks
type NSMapTableValueCallBacks struct {
	Describe func(*NSMapTable, unsafe.Pointer) NSString // Points to the function that produces an autoreleased NSString * describing the given element. If [NULL], then the map table produces a generic string description.
	Release  func(*NSMapTable, unsafe.Pointer)          // Points to the function that decrements a reference count for the given element, and if the reference count becomes zero, frees the given element. If [NULL], then nothing is done for reference counting or releasing.
	Retain   func(*NSMapTable, unsafe.Pointer)          // Points to the function that increments a reference count for the given element. If [NULL], then nothing is done for reference counting.

}

// NSOperatingSystemVersion - A structure that contains version information about the currently executing operating system, including major, minor, and patch version numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperatingSystemVersion
type NSOperatingSystemVersion struct {
	MajorVersion int // The major release number, such as 10 in version 10.9.3.
	MinorVersion int // The minor release number, such as 9 in version 10.9.3.
	PatchVersion int // The update release number, such as 3 in version 10.9.3.

}

// NSRange - A structure used to describe a portion of a series, such as characters in a string or objects in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
type NSRange struct {
	Location uint
	Length   uint
}

// NSSwappedDouble - Opaque structure containing endian-independent `double` value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSwappedDouble
type NSSwappedDouble struct {
	V uint64
}

// NSSwappedFloat - Opaque type containing an endian-independent `float` value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSwappedFloat
type NSSwappedFloat struct {
	V uint
}
