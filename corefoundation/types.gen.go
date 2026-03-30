// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"unsafe"
)

// C struct types

// CFAllocatorContext - A structure that defines the context or operating environment for an allocator (CFAllocator) object. Every Core Foundation allocator object must have a context defined for it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorContext
type CFAllocatorContext struct {
	Version         int                                // An integer of type [CFIndex]. Assign the version number of the allocator. Currently the only valid value is 0.
	Info            unsafe.Pointer                     // An untyped pointer to program-defined data. Allocate memory for this data and assign a pointer to it. This data is often control information for the allocator. You may assign [NULL].
	Retain          CFAllocatorRetainCallBack          // A prototype for a function callback that retains the data pointed to by the `info` field. In implementing this function, retain the data you have defined for the allocator context in this field. (This might make sense only if the data is a Core Foundation object.) You may set this function pointer to [NULL].
	Release         CFAllocatorReleaseCallBack         // A prototype for a function callback that releases the data pointed to by the `info` field. In implementing this function, release (or free) the data you have defined for the allocator context. You may set this function pointer to [NULL], but doing so might result in memory leaks.
	CopyDescription CFAllocatorCopyDescriptionCallBack // A prototype for a function callback that provides a description of the data pointed to by the `info` field. In implementing this function, return a reference to a CFString object that describes your allocator, particularly some characteristics of your program-defined data. You may set this function pointer to [NULL], in which case Core Foundation will provide a rudimentary description.
	Allocate        CFAllocatorAllocateCallBack        // A prototype for a function callback that allocates memory of a requested size. In implementing this function, allocate a block of memory of at least `size` bytes and return a pointer to the start of the block. The `hint` argument is a bitfield that you should currently not use (that is, assign 0). The `size` parameter should always be greater than 0. If it is not, or if problems in allocation occur, return [NULL]. This function pointer may not be assigned [NULL].
	Reallocate      CFAllocatorReallocateCallBack      // A prototype for a function callback that reallocates memory of a requested size for an existing block of memory.
	Deallocate      CFAllocatorDeallocateCallBack      // A prototype for a function callback that deallocates a given block of memory. In implementing this function, make the block of memory pointed to by `ptr` available for subsequent reuse by the allocator but unavailable for continued use by the program. The `ptr` parameter cannot be [NULL] and if the `ptr` parameter is not a block of memory that has been previously allocated by the allocator, the results are undefined; abnormal program termination can occur. You can set this callback to [NULL], in which case the [CFAllocatorDeallocate(_:_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)>) function has no effect.
	PreferredSize   CFAllocatorPreferredSizeCallBack   // A prototype for a function callback that determines whether there is enough free memory to satisfy a request. In implementing this function, return the actual size the allocator is likely to allocate given a request for a block of memory of size `size`. The `hint` argument is a bitfield that you should currently not use.

}

// CFArrayCallBacks - Structure containing the callbacks of a CFArray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCallBacks
type CFArrayCallBacks struct {
	Version         int                            // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
	Retain          CFArrayRetainCallBack          // The callback used to retain each value as they are added to the collection. If [NULL], values are not retained. See [CFArrayRetainCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFArrayRetainCallBack>) for a description of this callback.
	Release         CFArrayReleaseCallBack         // The callback used to release values as they are removed from the collection. If [NULL], values are not released. See [CFArrayReleaseCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFArrayReleaseCallBack>) for a description of this callback.
	CopyDescription CFArrayCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If [NULL], the collection will create a simple description of each value. See [CFArrayCopyDescriptionCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFArrayCopyDescriptionCallBack>) for a description of this callback.
	Equal           CFArrayEqualCallBack           // The callback used to compare values in the array for equality for some operations. If [NULL], the collection will use pointer equality to compare values in the collection. See [CFArrayEqualCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFArrayEqualCallBack>) for a description of this callback.

}

// CFBagCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values of a CFBag object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCallBacks
type CFBagCallBacks struct {
	Version         int                          // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
	Retain          CFBagRetainCallBack          // The callback used to retain each value as they are added to the collection. If [NULL], values are not retained. See [CFBagRetainCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFBagRetainCallBack>) for a descriptions of this function’s parameters.
	Release         CFBagReleaseCallBack         // The callback used to release values as they are removed from the collection. If [NULL], values are not released. See [CFBagReleaseCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFBagReleaseCallBack>) for a description of this callback.
	CopyDescription CFBagCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If [NULL], the collection will create a simple description of each value. See [CFBagCopyDescriptionCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFBagCopyDescriptionCallBack>) for a description of this callback.
	Equal           CFBagEqualCallBack           // The callback used to compare values in the collection for equality for some operations. If [NULL], the collection will use pointer equality to compare values in the collection. See [CFBagEqualCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFBagEqualCallBack>) for a description of this callback.
	Hash            CFBagHashCallBack            // The callback used to compute a hash code for values in a collection. If [NULL], the collection computes a hash code by converting the pointer value to an integer. See [CFBagHashCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFBagHashCallBack>) for a description of this callback.

}

// CFBinaryHeapCallBacks - Structure containing the callbacks for values for a [CFBinaryHeap] object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCallBacks
type CFBinaryHeapCallBacks struct {
	Version         int                                                                     // The version number of the structure type being passed in as a parameter to the [CFBinaryHeap] creation functions. This structure is version `0`.
	Compare         func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CFComparisonResult // The callback used to compare values in the binary heap in some operations. This field cannot be [NULL].
	CopyDescription func(unsafe.Pointer) uintptr                                            // Callback function used to get a description of a value in a binary heap.
	Release         func(uintptr, unsafe.Pointer)                                           // Callback function used to release a value before it is removed from a binary heap.
	Retain          func(uintptr, unsafe.Pointer) unsafe.Pointer                            // Callback function used to retain a value being added to a binary heap.

}

// CFBinaryHeapCompareContext - Not used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCompareContext
type CFBinaryHeapCompareContext struct {
	Version         int
	Info            unsafe.Pointer
	CopyDescription func(unsafe.Pointer) uintptr
	Release         func(unsafe.Pointer)
	Retain          func(unsafe.Pointer) unsafe.Pointer
}

// CFDictionaryKeyCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the keys in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryKeyCallBacks
type CFDictionaryKeyCallBacks struct {
	Version         int                                 // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
	Retain          CFDictionaryRetainCallBack          // The callback used to retain each key as they are added to the collection. This callback returns the value to use as the key in the dictionary, which is usually the value parameter passed to this callback, but may be a different value if a different value should be used as the key. If [NULL], keys are not retained. See [CFDictionaryRetainCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryRetainCallBack>) for a descriptions of this function’s parameters.
	Release         CFDictionaryReleaseCallBack         // The callback used to release keys as they are removed from the dictionary. If [NULL], keys are not released. See [CFDictionaryReleaseCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryReleaseCallBack>) for a description of this callback.
	CopyDescription CFDictionaryCopyDescriptionCallBack // The callback used to create a descriptive string representation of each key in the dictionary. If [NULL], the collection will create a simple description of each key. See [CFDictionaryCopyDescriptionCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryCopyDescriptionCallBack>) for a description of this callback.
	Equal           CFDictionaryEqualCallBack           // The callback used to compare keys in the dictionary for equality. If [NULL], the collection will use pointer equality to compare keys in the collection. See [CFDictionaryEqualCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryEqualCallBack>) for a description of this callback.
	Hash            CFDictionaryHashCallBack            // The callback used to compute a hash code for keys as they are used to access, add, or remove values in the dictionary. If [NULL], the collection computes a hash code by converting the pointer value to an integer. See [CFDictionaryHashCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryHashCallBack>) for a description of this callback.

}

// CFDictionaryValueCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryValueCallBacks
type CFDictionaryValueCallBacks struct {
	Version         int                                 // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
	Retain          CFDictionaryRetainCallBack          // The callback used to retain each value as they are added to the collection. This callback returns the value to use as the value in the dictionary, which is usually the value parameter passed to this callback, but may be a different value if a different value should be used as the value. If [NULL], values are not retained. See [CFDictionaryRetainCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryRetainCallBack>) for a descriptions of this function’s parameters.
	Release         CFDictionaryReleaseCallBack         // The callback used to release values as they are removed from the dictionary. If [NULL], values are not released. See [CFDictionaryReleaseCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryReleaseCallBack>) for a description of this callback.
	CopyDescription CFDictionaryCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the dictionary. If [NULL], the collection will create a simple description of each value. See [CFDictionaryCopyDescriptionCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryCopyDescriptionCallBack>) for a description of this callback.
	Equal           CFDictionaryEqualCallBack           // The callback used to compare values in the dictionary for equality. If [NULL], the collection will use pointer equality to compare values in the collection. See [CFDictionaryEqualCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFDictionaryEqualCallBack>) for a description of this callback.

}

// CFFileDescriptorContext - Defines a structure for the context of a CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorContext
type CFFileDescriptorContext struct {
	Version         int // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
	Info            unsafe.Pointer
	CopyDescription func(unsafe.Pointer) uintptr        // The callback used to create a descriptive string representation of the CFFileDescriptor.
	Release         func(unsafe.Pointer)                // The release callback used by the CFFileDescriptor.
	Retain          func(unsafe.Pointer) unsafe.Pointer // The retain callback used by the CFFileDescriptor.

}

// CFGregorianDate - Structure used to represent a point in time using the Gregorian calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDate
type CFGregorianDate struct {
	Day    int8
	Hour   int8
	Minute int8
	Month  int8
	Second float64
	Year   int32
}

// CFGregorianUnits - Structure used to represent a time interval in Gregorian units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnits
type CFGregorianUnits struct {
	Days    int32
	Hours   int32
	Minutes int32
	Months  int32
	Seconds float64
	Years   int32
}

// CFMachPortContext - A structure that contains program-defined data and callbacks with which you can configure a CFMachPort object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortContext
type CFMachPortContext struct {
	Version         int                                 // Version number of the structure. Must be `0`.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the CFMachPort object at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFMessagePortContext - A structure that contains program-defined data and callbacks with which you can configure a CFMessagePort object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortContext
type CFMessagePortContext struct {
	Version         int                                 // Version number of the structure. Must be `0`.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the message port at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFRange - A structure representing a range of sequential items in a container, such as characters in a buffer or elements in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRange
type CFRange struct {
	Location int // An integer representing the starting location of the range. For type compatibility with the rest of the system, `LONG_MAX` is the maximum value you should use for location.
	Length   int // An integer representing the number of items in the range. For type compatibility with the rest of the system, `LONG_MAX` is the maximum value you should use for length.

}

// CFRunLoopObserverContext - A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopObserver object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverContext
type CFRunLoopObserverContext struct {
	Version         int                                 // Version number of the structure. Must be `0`.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the run loop observer at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFRunLoopSourceContext - A structure that contains program-defined data and callbacks with which you can configure a version 0 CFRunLoopSource’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceContext
type CFRunLoopSourceContext struct {
	Version         int            // Version number of the structure. Must be 0.
	Info            unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the CFRunLoopSource at creation time. This pointer is passed to all the callbacks defined in the context.
	Cancel          func(unsafe.Pointer, uintptr, uintptr)
	CopyDescription func(unsafe.Pointer) uintptr               // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Equal           func(unsafe.Pointer, unsafe.Pointer) uint8 // An equality test callback for your program-defined `info` pointer. Can be [NULL].
	Hash            func(unsafe.Pointer) uint                  // A hash calculation callback for your program-defined `info` pointer. Can be [NULL].
	Perform         func(unsafe.Pointer)                       // A perform callback for the run loop source. This callback is called when the source has fired.
	Release         func(unsafe.Pointer)                       // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer        // A retain callback for your program-defined `info` pointer. Can be [NULL].
	Schedule        func(unsafe.Pointer, uintptr, uintptr)     // A scheduling callback for the run loop source. This callback is called when the source is added to a run loop mode. Can be [NULL].

}

// CFRunLoopSourceContext1 - A structure that contains program-defined data and callbacks with which you can configure a version 1 CFRunLoopSource’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceContext1
type CFRunLoopSourceContext1 struct {
	Version         int                                                               // Version number of the structure. Must be 1.
	Info            unsafe.Pointer                                                    // An arbitrary pointer to program-defined data, which can be associated with the run loop source at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr                                      // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Equal           func(unsafe.Pointer, unsafe.Pointer) uint8                        // An equality test callback for your program-defined `info` pointer. Can be [NULL].
	GetPort         func(unsafe.Pointer) uint                                         // A callback to retrieve the native Mach port represented by the source. This callback is called when the source is either added to or removed from a run loop mode.
	Hash            func(unsafe.Pointer) uint                                         // A hash calculation callback for your program-defined `info` pointer. Can be [NULL].
	Perform         func(unsafe.Pointer, int, uintptr, unsafe.Pointer) unsafe.Pointer // A perform callback for the run loop source. This callback is called when the source has fired.
	Release         func(unsafe.Pointer)                                              // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer                               // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFRunLoopTimerContext - A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopTimer’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerContext
type CFRunLoopTimerContext struct {
	Version         int                                 // Version number of the structure. Must be 0.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the run loop timer at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFSetCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values of a CFSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCallBacks
type CFSetCallBacks struct {
	Version         int                          // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is `0`.
	Retain          CFSetRetainCallBack          // The callback used to retain each value as they are added to the collection. If [NULL], values are not retained. See [CFSetRetainCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFSetRetainCallBack>) for a descriptions of this function’s parameters.
	Release         CFSetReleaseCallBack         // The callback used to release values as they are removed from the collection. If [NULL], values are not released. See [CFSetReleaseCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFSetReleaseCallBack>) for a description of this callback.
	CopyDescription CFSetCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If [NULL], the collection will create a simple description of each value. See [CFSetCopyDescriptionCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFSetCopyDescriptionCallBack>) for a description of this callback.
	Equal           CFSetEqualCallBack           // The callback used to compare values in the collection for equality for some operations. If [NULL], the collection will use pointer equality to compare values in the collection. See [CFSetEqualCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFSetEqualCallBack>) for a description of this callback.
	Hash            CFSetHashCallBack            // The callback used to compute a hash code for values in a collection. If [NULL], the collection computes a hash code by converting the pointer value to an integer. See [CFSetHashCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFSetHashCallBack>) for a description of this callback.

}

// CFSocketContext - A structure that contains program-defined data and callbacks with which you can configure a CFSocket object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketContext
type CFSocketContext struct {
	Version         int                                 // Version number of the structure. Must be `0`.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the CFSocket object at creation time. This pointer is passed to all the callbacks defined in the context.
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFSocketSignature - A structure that fully specifies the communication protocol and connection address of a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSignature
type CFSocketSignature struct {
	ProtocolFamily int32     // The protocol family of the socket.
	SocketType     int32     // The socket type of the socket.
	Protocol       int32     // The protocol type of the socket.
	Address        CFDataRef // A CFData object holding the contents of a `struct sockaddr` appropriate for the given protocol family (`struct sockaddr_in` or `struct sockaddr_in6`, for example), identifying the address of the socket.

}

// CFStreamClientContext - A structure that contains program-defined data and callbacks with which you can configure a stream’s client behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamClientContext
type CFStreamClientContext struct {
	Version         int                                 // Version number of the structure. Must be `0`.
	Info            unsafe.Pointer                      // An arbitrary pointer to program-defined data, which can be associated with the client. This pointer is passed to the callbacks defined in the context and to the client callback function [CFReadStreamClientCallBack](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFReadStreamClientCallBack>).
	CopyDescription func(unsafe.Pointer) uintptr        // A copy description callback for your program-defined `info` pointer. Can be [NULL].
	Release         func(unsafe.Pointer)                // A release callback for your program-defined `info` pointer. Can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // A retain callback for your program-defined `info` pointer. Can be [NULL].

}

// CFStreamError - The structure returned by [CFReadStreamGetError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFReadStreamGetError(_:)>) and [CFWriteStreamGetError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFWriteStreamGetError(_:)>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamError
type CFStreamError struct {
	Domain int   // The error domain that should be used to interpret the error. See [CFStreamErrorDomain](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFStreamErrorDomain>) for possible values.
	Error  int32 // The error code.

}

// CFStringInlineBuffer - Defines the buffer and related fields used for in-line buffer access of characters in CFString objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringInlineBuffer
type CFStringInlineBuffer struct {
	TheString           CFStringRef
	RangeToBuffer       CFRange
	BufferedRangeStart  int
	BufferedRangeEnd    int
	Buffer              uint16
	DirectCStringBuffer *byte
	DirectUniCharBuffer *uint16
}

// CFSwappedFloat32 - Structure holding a 32-bit float value in a platform-independentbyte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSwappedFloat32
type CFSwappedFloat32 struct {
	V uint32 // A 32-bit float value stored with a platform-independentbyte order.

}

// CFSwappedFloat64 - Structure holding a 64-bit float value in a platform-independentbyte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSwappedFloat64
type CFSwappedFloat64 struct {
	V uint64 // A 64-bit float value stored with a platform-independentbyte order.

}

// CFTreeContext - Structure containing program-defined data and callbacks for a CFTree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeContext
type CFTreeContext struct {
	Version         int                           // The version number of the structure type being passed in as a parameter to a CFTree creation function. This structure is version `0`.
	Info            unsafe.Pointer                // A C pointer to a program-defined block of data, referred to as the information pointer.
	Retain          CFTreeRetainCallBack          // The callback used to retain the `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. This value may be [NULL].
	Release         CFTreeReleaseCallBack         // The callback used to release a previously retained `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. This value may be [NULL].
	CopyDescription CFTreeCopyDescriptionCallBack // The callback used to provide a description of the `info` field.

}

// CFUUIDBytes - A 128-bit struct that represents a UUID as raw bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDBytes
type CFUUIDBytes struct {
	Byte0  uint8 // The first byte.
	Byte1  uint8 // The second byte.
	Byte2  uint8 // The third byte.
	Byte3  uint8 // The fourth byte.
	Byte4  uint8 // The fifth byte.
	Byte5  uint8 // The sixth byte.
	Byte6  uint8 // The seventh byte.
	Byte7  uint8 // The eighth byte.
	Byte8  uint8 // The ninth byte.
	Byte9  uint8 // The tenth byte.
	Byte10 uint8 // The eleventh byte.
	Byte11 uint8 // The twelfth byte.
	Byte12 uint8 // The thirteenth byte.
	Byte13 uint8 // The fourteenth byte.
	Byte14 uint8 // The fifteenth byte.
	Byte15 uint8 // The sixteenth byte.

}

// CFXMLAttributeDeclarationInfo - Contains information about an element attribute definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLAttributeDeclarationInfo
type CFXMLAttributeDeclarationInfo struct {
	AttributeName CFStringRef // The name of the attribute.
	TypeString    CFStringRef // Describes the declaration of a single attribute.
	DefaultString CFStringRef // The attribute’s default value.

}

// CFXMLAttributeListDeclarationInfo - Contains a list of the attributes associated with an element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLAttributeListDeclarationInfo
type CFXMLAttributeListDeclarationInfo struct {
	NumberOfAttributes int                            // The number of attributes in the array.
	Attributes         *CFXMLAttributeDeclarationInfo // A C array of attributes.

}

// CFXMLDocumentInfo - Contains the source URL and text encoding information for the XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLDocumentInfo
type CFXMLDocumentInfo struct {
	SourceURL CFURLRef // The source URL of the XML document.
	Encoding  uint32   // The text encoding of the XML document.

}

// CFXMLDocumentTypeInfo - Contains the external ID of the DTD.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLDocumentTypeInfo
type CFXMLDocumentTypeInfo struct {
	ExternalID CFXMLExternalID // The external ID of the DTD.

}

// CFXMLElementInfo - Contains a list of element attributes packaged as CFDictionary key/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLElementInfo
type CFXMLElementInfo struct {
	Attributes     CFDictionaryRef // The dictionary of attribute values.
	AttributeOrder CFArrayRef      // An array specifying the order in which the attributes appeared in the XML document.
	IsEmpty        bool            // A flag indicating whether the element was expressed in closed form.
	_reserved      unsafe.Pointer
}

// CFXMLElementTypeDeclarationInfo - Contains a description of the element type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLElementTypeDeclarationInfo
type CFXMLElementTypeDeclarationInfo struct {
	ContentDescription CFStringRef // A textual description of the element type.

}

// CFXMLEntityInfo - Contains information describing an XML entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityInfo
type CFXMLEntityInfo struct {
	EntityType      CFXMLEntityTypeCode // The entity type code.
	ReplacementText CFStringRef         // [NULL] if `entityType` is external or unparsed, otherwise the text that the entity should be replaced with.
	EntityID        CFXMLExternalID     // `entityID.SystemID()` will be [NULL] if `entityType` is internal.
	NotationName    CFStringRef         // [NULL] if `entityType` is parsed.

}

// CFXMLEntityReferenceInfo - Contains information describing an XML entity reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityReferenceInfo
type CFXMLEntityReferenceInfo struct {
	EntityType CFXMLEntityTypeCode // The entity type code.

}

// CFXMLExternalID - Contains the system and public IDs for an external entity reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLExternalID
type CFXMLExternalID struct {
	SystemID CFURLRef    // The systemID URL.
	PublicID CFStringRef // The publicID string.

}

// CFXMLNotationInfo - Contains the external ID of the notation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNotationInfo
type CFXMLNotationInfo struct {
	ExternalID CFXMLExternalID // The external ID of the notation.

}

// CFXMLParserCallBacks - Contains version information and function pointers to callbacks needed when parsing XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCallBacks
type CFXMLParserCallBacks struct {
	Version               int                                      // Version number. Must be `0`.
	CreateXMLStructure    CFXMLParserCreateXMLStructureCallBack    // Called when an XML structure is created.
	AddChild              CFXMLParserAddChildCallBack              // Called when a child is added.
	EndXMLStructure       CFXMLParserEndXMLStructureCallBack       // Called when an XML structure has ended.
	ResolveExternalEntity CFXMLParserResolveExternalEntityCallBack // Called when an external entity needs to be resolved.
	HandleError           CFXMLParserHandleErrorCallBack           // Called when a parse error needs to be handled.

}

// CFXMLParserContext - Contains version information and function pointers to callbacks used when handling a program-defined context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserContext
type CFXMLParserContext struct {
	Version         int                                // Version number of this structure. Must be 0.
	Info            unsafe.Pointer                     // An arbitrary program-defined value passed to all the callbacks in this structure and in the [CFXMLParserCallBacks](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFXMLParserCallBacks>) structure.
	Retain          CFXMLParserRetainCallBack          // A retain callback for your program-defined context data. Optional.
	Release         CFXMLParserReleaseCallBack         // A release callback for your program-defined context data. Optional.
	CopyDescription CFXMLParserCopyDescriptionCallBack // A copy description callback for your program-defined context data. Optional.

}

// CFXMLProcessingInstructionInfo - Contains the text of the processing instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLProcessingInstructionInfo
type CFXMLProcessingInstructionInfo struct {
	DataString CFStringRef // The text of the processing instruction.

}

// CGAffineTransform
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
type CGAffineTransform struct {
	A  float64
	B  float64
	C  float64
	D  float64
	Tx float64
	Ty float64
}

// CGAffineTransformComponents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransformComponents
type CGAffineTransformComponents struct {
	Scale           CGSize
	HorizontalShear float64
	Rotation        float64
	Translation     CGVector
}

// CGPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
type CGPoint struct {
	X float64
	Y float64
}

// CGRect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect
type CGRect struct {
	Origin CGPoint
	Size   CGSize
}

// CGSize - A structure that contains width and height values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGSize
type CGSize struct {
	Width  float64 // A width value.
	Height float64 // A height value.

}

// CGVector - A structure that contains a two-dimensional vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGVector
type CGVector struct {
	Dx float64 // The x component of the vector.
	Dy float64 // The y component of the vector.

}

// IUnknownVTbl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/IUnknownVTbl
type IUnknownVTbl struct {
	AddRef         func(unsafe.Pointer) uint
	QueryInterface func(unsafe.Pointer, CFUUIDBytes, unsafe.Pointer) int
	Release        func(unsafe.Pointer) uint
}
