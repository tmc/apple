// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
)

// The class instance for the [NSMutableData] class.
var (
	_NSMutableDataClass     NSMutableDataClass
	_NSMutableDataClassOnce sync.Once
)

func getNSMutableDataClass() NSMutableDataClass {
	_NSMutableDataClassOnce.Do(func() {
		_NSMutableDataClass = NSMutableDataClass{class: objc.GetClass("NSMutableData")}
	})
	return _NSMutableDataClass
}

// GetNSMutableDataClass returns the class object for NSMutableData.
func GetNSMutableDataClass() NSMutableDataClass {
	return getNSMutableDataClass()
}

type NSMutableDataClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableDataClass) Alloc() NSMutableData {
	rv := objc.Send[NSMutableData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object representing a dynamic byte buffer in memory.
//
// # Overview
// 
// In Swift, this object bridges to [Data]; use [NSMutableData] when you need
// reference semantics or other Foundation-specific behavior.
// 
// [NSMutableData] and its superclass [NSData] provide data objects, or
// object-oriented wrappers for byte buffers. Data objects let simple
// allocated buffers (that is, data with no embedded pointers) take on the
// behavior of Foundation objects. They are typically used for data storage
// and are also useful in Distributed Objects applications, where data
// contained in data objects can be copied or moved between applications.
// [NSData] creates static data objects, and [NSMutableData] creates dynamic
// data objects. You can easily convert one type of data object to the other
// with the initializer that takes an [NSData] object or an [NSMutableData]
// object as an argument.
// 
// The following [NSData] methods change when used on a mutable data object:
// 
// - [NSMutableData.InitWithBytesNoCopyLengthFreeWhenDone] -
// [NSMutableData.InitWithBytesNoCopyLengthDeallocator] - [NSMutableData.InitWithBytesNoCopyLength] -
// [DataWithBytesNoCopyLengthFreeWhenDone] - [DataWithBytesNoCopyLength]
// 
// When called, the bytes are immediately copied and then the buffer is freed.
// 
// [NSMutableData] is “toll-free bridged” with its Core Foundation
// counterpart, [CFData]. See [Toll-Free Bridging] for more information on
// toll-free bridging.
//
// [CFData]: https://developer.apple.com/documentation/CoreFoundation/CFData
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating Mutable Data
//
//   - [NSMutableData.InitWithCapacity]: Returns an initialized mutable data object capable of holding the specified number of bytes.
//   - [NSMutableData.InitWithLength]: Initializes and returns a mutable data object containing a given number of zeroed bytes.
//
// # Accessing Raw Bytes
//
//   - [NSMutableData.MutableBytes]: A pointer to the data contained by the mutable data object.
//
// # Adding Bytes
//
//   - [NSMutableData.AppendBytesLength]: Appends to the receiver a given number of bytes from a given buffer.
//   - [NSMutableData.AppendData]: Appends the content of another data object to the receiver.
//   - [NSMutableData.IncreaseLengthBy]: Increases the length of the receiver by a given number of bytes.
//
// # Modifying Bytes
//
//   - [NSMutableData.ReplaceBytesInRangeWithBytes]: Replaces with a given set of bytes a given range within the contents of the receiver.
//   - [NSMutableData.ReplaceBytesInRangeWithBytesLength]: Replaces with a given set of bytes a given range within the contents of the receiver.
//   - [NSMutableData.ResetBytesInRange]: Replaces with zeroes the contents of the receiver in a given range.
//   - [NSMutableData.SetData]: Replaces the entire contents of the receiver with the contents of another data object.
//
// # Compressing and Decompressing Data
//
//   - [NSMutableData.CompressUsingAlgorithmError]: Compresses the data object’s bytes using an algorithm that you specify.
//   - [NSMutableData.DecompressUsingAlgorithmError]: Decompresses the data object’s bytes.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData
type NSMutableData struct {
	NSData
}

// NSMutableDataFromID constructs a [NSMutableData] from an objc.ID.
//
// An object representing a dynamic byte buffer in memory.
func NSMutableDataFromID(id objc.ID) NSMutableData {
	return NSMutableData{NSData: NSDataFromID(id)}
}
// NOTE: NSMutableData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMutableData] class.
//
// # Creating Mutable Data
//
//   - [INSMutableData.InitWithCapacity]: Returns an initialized mutable data object capable of holding the specified number of bytes.
//   - [INSMutableData.InitWithLength]: Initializes and returns a mutable data object containing a given number of zeroed bytes.
//
// # Accessing Raw Bytes
//
//   - [INSMutableData.MutableBytes]: A pointer to the data contained by the mutable data object.
//
// # Adding Bytes
//
//   - [INSMutableData.AppendBytesLength]: Appends to the receiver a given number of bytes from a given buffer.
//   - [INSMutableData.AppendData]: Appends the content of another data object to the receiver.
//   - [INSMutableData.IncreaseLengthBy]: Increases the length of the receiver by a given number of bytes.
//
// # Modifying Bytes
//
//   - [INSMutableData.ReplaceBytesInRangeWithBytes]: Replaces with a given set of bytes a given range within the contents of the receiver.
//   - [INSMutableData.ReplaceBytesInRangeWithBytesLength]: Replaces with a given set of bytes a given range within the contents of the receiver.
//   - [INSMutableData.ResetBytesInRange]: Replaces with zeroes the contents of the receiver in a given range.
//   - [INSMutableData.SetData]: Replaces the entire contents of the receiver with the contents of another data object.
//
// # Compressing and Decompressing Data
//
//   - [INSMutableData.CompressUsingAlgorithmError]: Compresses the data object’s bytes using an algorithm that you specify.
//   - [INSMutableData.DecompressUsingAlgorithmError]: Decompresses the data object’s bytes.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData
type INSMutableData interface {
	INSData

	// Topic: Creating Mutable Data

	// Returns an initialized mutable data object capable of holding the specified number of bytes.
	InitWithCapacity(capacity uint) NSMutableData
	// Initializes and returns a mutable data object containing a given number of zeroed bytes.
	InitWithLength(length uint) NSMutableData

	// Topic: Accessing Raw Bytes

	// A pointer to the data contained by the mutable data object.
	MutableBytes() unsafe.Pointer

	// Topic: Adding Bytes

	// Appends to the receiver a given number of bytes from a given buffer.
	AppendBytesLength(bytes []byte)
	// Appends the content of another data object to the receiver.
	AppendData(other INSData)
	// Increases the length of the receiver by a given number of bytes.
	IncreaseLengthBy(extraLength uint)

	// Topic: Modifying Bytes

	// Replaces with a given set of bytes a given range within the contents of the receiver.
	ReplaceBytesInRangeWithBytes(range_ NSRange, bytes unsafe.Pointer)
	// Replaces with a given set of bytes a given range within the contents of the receiver.
	ReplaceBytesInRangeWithBytesLength(range_ NSRange, replacementBytes unsafe.Pointer, replacementLength uint)
	// Replaces with zeroes the contents of the receiver in a given range.
	ResetBytesInRange(range_ NSRange)
	// Replaces the entire contents of the receiver with the contents of another data object.
	SetData(data INSData)

	// Topic: Compressing and Decompressing Data

	// Compresses the data object’s bytes using an algorithm that you specify.
	CompressUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (bool, error)
	// Decompresses the data object’s bytes.
	DecompressUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (bool, error)
}





// Init initializes the instance.
func (m NSMutableData) Init() NSMutableData {
	rv := objc.Send[NSMutableData](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableData) Autorelease() NSMutableData {
	rv := objc.Send[NSMutableData](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableData creates a new NSMutableData instance.
func NewNSMutableData() NSMutableData {
	class := getNSMutableDataClass()
	rv := objc.Send[NSMutableData](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-4t5yq
func NewMutableDataWithBase64EncodedDataOptions(base64Data INSData, options NSDataBase64DecodingOptions) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return NSMutableDataFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-3ksry
func NewMutableDataWithBase64EncodedStringOptions(base64String string, options NSDataBase64DecodingOptions) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedString:options:"), objc.String(base64String), options)
	return NSMutableDataFromID(rv)
}


// Initializes a data object initialized with the given Base64 encoded string.
//
// base64String: A Base-64 encoded string.
//
// # Return Value
// 
// A data object built by Base-64 decoding the provided string. Returns `nil`
// if the data object could not be decoded.
//
// # Discussion
// 
// Although this method was only introduced publicly for iOS 7, it has existed
// since iOS 4; you can use it if your application needs to target an
// operating system prior to iOS 7. This method behaves like
// [init(base64EncodedString:options:)], but ignores all unknown characters.
//
// [init(base64EncodedString:options:)]: https://developer.apple.com/documentation/Foundation/NSData/init(base64EncodedString:options:)
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoding:)
func NewMutableDataWithBase64Encoding(base64String string) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64Encoding:"), objc.String(base64String))
	return NSMutableDataFromID(rv)
}


// Initializes a data object filled with a given number of bytes copied from a
// given buffer.
//
// # Discussion
// 
// A data object initialized by adding to it `length` bytes of data copied
// from the buffer `bytes`. The returned object might be different than the
// original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(bytes:length:)
func NewMutableDataWithBytesLength(bytes []byte) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
	return NSMutableDataFromID(rv)
}


// Initializes a data object filled with a given number of bytes of data from
// a given buffer.
//
// bytes: A buffer containing data for the new object. `bytes` must point to a memory
// block allocated with `malloc`.
//
// length: The number of bytes to hold from `bytes`. This value must not exceed the
// length of `bytes`.
//
// # Return Value
// 
// A data object initialized by adding to it `length` bytes of data from the
// buffer `bytes`. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// The returned object takes ownership of the `bytes` pointer and frees it on
// deallocation. Therefore, `bytes` must point to a memory block allocated
// with `malloc`.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:)
func NewMutableDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	return NSMutableDataFromID(rv)
}


// Initializes a data object filled with a given number of bytes of data from
// a given buffer, with a custom deallocator block.
//
// bytes: A buffer containing data for the new object.
//
// length: The number of bytes to hold from `bytes`. This value must not exceed the
// length of `bytes`.
//
// deallocator: A block to invoke when the resulting [NSData] object is deallocated.
//
// # Return Value
// 
// A data object initialized by adding to it `length` bytes of data from the
// buffer `bytes`. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// Use this method to define your own deallocation behavior for the data
// buffer you provide.
// 
// In order to avoid any inadvertent strong reference cycles, you should avoid
// capturing pointers to any objects that may in turn maintain strong
// references to the [NSData] object. This includes explicit references to
// `self`, and implicit references to `self` due to direct instance variable
// access. To make it easier to avoid these references, the `deallocator`
// block takes two parameters, a pointer to the `buffer`, and its length; you
// should always use these values instead of trying to use references from
// outside the block.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:deallocator:)
func NewMutableDataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator uint) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	return NSMutableDataFromID(rv)
}


// Initializes a newly allocated data object by adding the given number of
// bytes from the given buffer.
//
// bytes: A buffer containing data for the new object. If `flag` is [true], `bytes`
// must point to a memory block allocated with `malloc`.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// length: The number of bytes to hold from `bytes`. This value must not exceed the
// length of `bytes`.
//
// b: If [true], the returned object takes ownership of the `bytes` pointer and
// frees it on deallocation.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:freeWhenDone:)
func NewMutableDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return NSMutableDataFromID(rv)
}


// Returns an initialized mutable data object capable of holding the specified
// number of bytes.
//
// capacity: The number of bytes the data object can initially contain.
//
// # Return Value
// 
// An initialized [NSMutableData] object capable of holding `capacity` bytes.
// The returned object has the same memory alignment guarantees as
// `malloc(_:)`.
//
// # Discussion
// 
// This method doesn’t necessarily allocate the requested memory right away.
// Mutable data objects allocate additional memory as needed, so `capacity`
// simply establishes the object’s initial capacity. When it does allocate
// the initial memory, though, it allocates the specified amount. This method
// sets the length of the data object to `0`.
// 
// If the capacity specified in `capacity` is greater than four memory pages
// in size, this method may round the amount of requested memory up to the
// nearest full page.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/init(capacity:)
func NewMutableDataWithCapacity(capacity uint) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	return NSMutableDataFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMutableDataWithCoder(coder INSCoder) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableDataFromID(rv)
}


// Initializes a data object with the content of the file at a given path.
//
// path: The absolute path of the file from which to read data.
//
// # Return Value
// 
// A data object initialized by reading into it the data from the file
// specified by `path`.
//
// # Discussion
// 
// This method is equivalent to [InitWithContentsOfFileOptionsError] with no
// options.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:)
func NewMutableDataWithContentsOfFile(path string) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSMutableDataFromID(rv)
}


// Initializes a data object with the content of the file at a given path.
//
// path: The absolute path of the file from which to read data.
//
// readOptionsMask: A mask that specifies options for reading the data. Constant components are
// described in [NSData.ReadingOptions].
// //
// [NSData.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
//
// # Return Value
// 
// A data object initialized by reading into it the data from the file
// specified by `path`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:options:)
func NewMutableDataWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSMutableData, error) {
	var errorPtr objc.ID
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:options:error:"), objc.String(path), readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableDataFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableDataFromID(rv), nil
}


//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:)
func NewMutableDataWithContentsOfURL(url INSURL) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSMutableDataFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:options:)
func NewMutableDataWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSMutableData, error) {
	var errorPtr objc.ID
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableDataFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableDataFromID(rv), nil
}


// Initializes a data object with the contents of another data object.
//
// data: A data object.
//
// # Return Value
// 
// A data object initialized with the contents `data`.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(data:)
func NewMutableDataWithData(data INSData) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSMutableDataFromID(rv)
}


// Initializes and returns a mutable data object containing a given number of
// zeroed bytes.
//
// length: The number of bytes the object initially contains.
//
// # Return Value
// 
// An initialized [NSMutableData] object containing `length` zeroed bytes. The
// returned object has the same memory alignment guarantees as `malloc(_:)`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/init(length:)
func NewMutableDataWithLength(length uint) NSMutableData {
	instance := getNSMutableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLength:"), length)
	return NSMutableDataFromID(rv)
}







// Returns an initialized mutable data object capable of holding the specified
// number of bytes.
//
// capacity: The number of bytes the data object can initially contain.
//
// # Return Value
// 
// An initialized [NSMutableData] object capable of holding `capacity` bytes.
// The returned object has the same memory alignment guarantees as
// `malloc(_:)`.
//
// # Discussion
// 
// This method doesn’t necessarily allocate the requested memory right away.
// Mutable data objects allocate additional memory as needed, so `capacity`
// simply establishes the object’s initial capacity. When it does allocate
// the initial memory, though, it allocates the specified amount. This method
// sets the length of the data object to `0`.
// 
// If the capacity specified in `capacity` is greater than four memory pages
// in size, this method may round the amount of requested memory up to the
// nearest full page.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/init(capacity:)
func (m NSMutableData) InitWithCapacity(capacity uint) NSMutableData {
	rv := objc.Send[NSMutableData](m.ID, objc.Sel("initWithCapacity:"), capacity)
	return rv
}

// Initializes and returns a mutable data object containing a given number of
// zeroed bytes.
//
// length: The number of bytes the object initially contains.
//
// # Return Value
// 
// An initialized [NSMutableData] object containing `length` zeroed bytes. The
// returned object has the same memory alignment guarantees as `malloc(_:)`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/init(length:)
func (m NSMutableData) InitWithLength(length uint) NSMutableData {
	rv := objc.Send[NSMutableData](m.ID, objc.Sel("initWithLength:"), length)
	return rv
}

// Appends to the receiver a given number of bytes from a given buffer.
//
// bytes: A buffer containing data to append to the receiver’s content.
//
// length: The number of bytes from `bytes` to append.
//
// # Discussion
// 
// A sample using this method can be found in [Working With Mutable Binary
// Data].
//
// [Working With Mutable Binary Data]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/BinaryData/Tasks/WorkingMutableData.html#//apple_ref/doc/uid/20002150
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:length:)
func (m NSMutableData) AppendBytesLength(bytes []byte) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
}

// Appends the content of another data object to the receiver.
//
// other: The data object whose content is to be appended to the contents of the
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:)
func (m NSMutableData) AppendData(other INSData) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendData:"), other)
}

// Increases the length of the receiver by a given number of bytes.
//
// extraLength: The number of bytes by which to increase the receiver’s length.
//
// # Discussion
// 
// The additional bytes are all set to `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/increaseLength(by:)
func (m NSMutableData) IncreaseLengthBy(extraLength uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("increaseLengthBy:"), extraLength)
}

// Replaces with a given set of bytes a given range within the contents of the
// receiver.
//
// range: The range within the receiver’s contents to replace with `bytes`. The
// range must not exceed the bounds of the receiver.
//
// bytes: The data to insert into the receiver’s contents.
//
// # Discussion
// 
// If the location of `range` isn’t within the receiver’s range of bytes,
// an [NSRangeException] is raised. The receiver is resized to accommodate the
// new bytes, if necessary.
// 
// A sample using this method is given in [Working With Mutable Binary Data].
//
// [Working With Mutable Binary Data]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/BinaryData/Tasks/WorkingMutableData.html#//apple_ref/doc/uid/20002150
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:)
func (m NSMutableData) ReplaceBytesInRangeWithBytes(range_ NSRange, bytes unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceBytesInRange:withBytes:"), range_, bytes)
}

// Replaces with a given set of bytes a given range within the contents of the
// receiver.
//
// range: The range within the receiver’s contents to replace with `bytes`. The
// range must not exceed the bounds of the receiver.
//
// replacementBytes: The data to insert into the receiver’s contents.
//
// replacementLength: The number of bytes to take from `replacementBytes`.
//
// # Discussion
// 
// If the length of `range` is not equal to `replacementLength`, the receiver
// is resized to accommodate the new bytes. Any bytes past `range` in the
// receiver are shifted to accommodate the new bytes. You can therefore pass
// [NULL] for `replacementBytes` and `0` for `replacementLength` to delete
// bytes in the receiver in the range `range`. You can also replace a range
// (which might be zero-length) with more bytes than the length of the range,
// which has the effect of insertion (or “replace some and insert more”).
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:length:)
func (m NSMutableData) ReplaceBytesInRangeWithBytesLength(range_ NSRange, replacementBytes unsafe.Pointer, replacementLength uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceBytesInRange:withBytes:length:"), range_, replacementBytes, replacementLength)
}

// Replaces with zeroes the contents of the receiver in a given range.
//
// range: The range within the contents of the receiver to be replaced by zeros. The
// range must not exceed the bounds of the receiver.
//
// # Discussion
// 
// If the location of `range` isn’t within the receiver’s range of bytes,
// an [NSRangeException] is raised. The receiver is resized to accommodate the
// new bytes, if necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/resetBytes(in:)
func (m NSMutableData) ResetBytesInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("resetBytesInRange:"), range_)
}

// Replaces the entire contents of the receiver with the contents of another
// data object.
//
// data: The data object whose content replaces that of the receiver.
//
// # Discussion
// 
// As part of its implementation, this method calls
// [ReplaceBytesInRangeWithBytes].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/setData(_:)
func (m NSMutableData) SetData(data INSData) {
	objc.Send[objc.ID](m.ID, objc.Sel("setData:"), data)
}

// Compresses the data object’s bytes using an algorithm that you specify.
//
// algorithm: The algorithm to use to compress the data. For a list of available
// algorithms, see [NSData.CompressionAlgorithm].
// //
// [NSData.CompressionAlgorithm]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
//
// # Discussion
// 
// Use this method to compress in-memory data when you want to reduce memory
// usage and can afford the time to compress and decompress the data. If your
// data object is already in a compressed format, such as media formats like
// JPEG images or AAC audio, [CompressUsingAlgorithmError] may provide minimal
// or no benefit.
// 
// The following example shows how to compress data from a string and prints
// the sizes of the data instances to illustrate the amount of compression:
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/compress(using:)
func (m NSMutableData) CompressUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("compressUsingAlgorithm:error:"), algorithm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("compressUsingAlgorithm:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Decompresses the data object’s bytes.
//
// algorithm: The algorithm to use for decompressing the data. For a list of available
// algorithms, see [NSData.CompressionAlgorithm].
// //
// [NSData.CompressionAlgorithm]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
//
// # Discussion
// 
// Use this method to inflate in-memory data when you need uncompressed bytes.
// Specify the same algorithm used to compress the data to successfully
// decompress it.
// 
// The following example shows how to inflate an instance of [NSMutableData]
// compressed with the [DataCompressionAlgorithmZlib] algorithm:
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/decompress(using:)
func (m NSMutableData) DecompressUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("decompressUsingAlgorithm:error:"), algorithm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("decompressUsingAlgorithm:error: returned NO with nil NSError")
	}
	return rv, nil

}





// Creates and returns an mutable data object containing a given number of
// zeroed bytes.
//
// length: The number of bytes the new data object initially contains.
//
// # Return Value
// 
// A new [NSMutableData] object of `length` bytes, filled with zeros. The
// returned object has the same memory alignment guarantees as `malloc(_:)`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/dataWithLength:
func (_NSMutableDataClass NSMutableDataClass) DataWithLength(length uint) NSMutableData {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableDataClass.class), objc.Sel("dataWithLength:"), length)
	return NSMutableDataFromID(rv)
}

// Creates and returns a mutable data object capable of holding the specified
// number of bytes.
//
// aNumItems: The number of bytes the new data object can initially contain.
//
// # Return Value
// 
// A new [NSMutableData] object capable of holding `aNumItems` bytes.
// 
// The returned object has the same memory alignment guarantees as
// `malloc(_:)`.
//
// # Discussion
// 
// This method doesn’t necessarily allocate the requested memory right away.
// Mutable data objects allocate additional memory as needed, so `aNumItems`
// simply establishes the object’s initial capacity. When it does allocate
// the initial memory, though, it allocates the specified amount. This method
// sets the length of the data object to `0`.
// 
// If the capacity specified in `aNumItems` is greater than four memory pages
// in size, this method may round the amount of requested memory up to the
// nearest full page.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/dataWithCapacity:
func (_NSMutableDataClass NSMutableDataClass) DataWithCapacity(aNumItems uint) NSMutableData {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableDataClass.class), objc.Sel("dataWithCapacity:"), aNumItems)
	return NSMutableDataFromID(rv)
}








// A pointer to the data contained by the mutable data object.
//
// # Discussion
// 
// If the length of the receiver’s data is not zero, this property is
// guaranteed to contain a pointer to the object’s internal bytes. If the
// length of receiver’s data zero, this property may or may not contain
// [NULL] dependent upon many factors related to how the object was created
// (moreover, in this case the method result might change between different
// releases). The returned pointer is valid until the data object is
// deallocated.
// 
// A sample using this method can be found in [Working With Mutable Binary
// Data].
//
// [Working With Mutable Binary Data]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/BinaryData/Tasks/WorkingMutableData.html#//apple_ref/doc/uid/20002150
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableData/mutableBytes
func (m NSMutableData) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("mutableBytes"))
	return rv
}

































