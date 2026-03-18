// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSData] class.
var (
	_NSDataClass     NSDataClass
	_NSDataClassOnce sync.Once
)

func getNSDataClass() NSDataClass {
	_NSDataClassOnce.Do(func() {
		_NSDataClass = NSDataClass{class: objc.GetClass("NSData")}
	})
	return _NSDataClass
}

// GetNSDataClass returns the class object for NSData.
func GetNSDataClass() NSDataClass {
	return getNSDataClass()
}

type NSDataClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDataClass) Alloc() NSData {
	rv := objc.Send[NSData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A static byte buffer in memory.
//
// # Overview
// 
// In Swift, the buffer bridges to [Data]; use [NSData] when you need
// reference semantics or other Foundation-specific behavior.
// 
// [NSData] and its mutable subclass [NSMutableData] provide data objects, or
// object-oriented wrappers for byte buffers. Data objects let simple
// allocated buffers (that is, data with no embedded pointers) take on the
// behavior of Foundation objects.
// 
// The size of the data is subject to a theoretical limit of about 8 exabytes
// (1 EB = 10¹⁸ bytes; in practice, the limit should not be a factor).
// 
// [NSData] is with its Core Foundation counterpart, [CFData]. See [Toll-Free
// Bridging] for more information on toll-free bridging.
// 
// # Writing Data Atomically
// 
// [NSData] provides methods for atomically saving their contents to a file,
// which guarantee that the data is either saved in its entirety, or it fails
// completely. An atomic write first writes the data to a temporary file and
// then, only if this write succeeds, moves the temporary file to its final
// location.
// 
// Although atomic write operations minimize the risk of data loss due to
// corrupt or partially written files, they may not be appropriate when
// writing to a temporary directory, the user’s home directory or other
// publicly accessible directories. When you work with a publicly accessible
// file, treat that file as an untrusted and potentially dangerous resource.
// An attacker may compromise or corrupt these files. The attacker can also
// replace the files with hard or symbolic links, causing your write
// operations to overwrite or corrupt other system resources.
// 
// Avoid using the [NSData.WriteToURLAtomically] method (and the related methods)
// when working inside a publicly accessible directory. Instead, use
// [NSFileHandle] with an existing file descriptor to securely write the file.
// 
// For more information, see [Securing File Operations] in [Secure Coding
// Guide].
//
// [CFData]: https://developer.apple.com/documentation/CoreFoundation/CFData
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
// [Secure Coding Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Introduction.html#//apple_ref/doc/uid/TP40002415
// [Securing File Operations]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Articles/RaceConditions.html#//apple_ref/doc/uid/TP40002585-SW9
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating Data
//
//   - [NSData.InitWithBytesLength]: Initializes a data object filled with a given number of bytes copied from a given buffer.
//   - [NSData.InitWithBytesNoCopyLength]: Initializes a data object filled with a given number of bytes of data from a given buffer.
//   - [NSData.InitWithBytesNoCopyLengthDeallocator]: Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block.
//   - [NSData.InitWithBytesNoCopyLengthFreeWhenDone]: Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
//   - [NSData.InitWithData]: Initializes a data object with the contents of another data object.
//
// # Reading Data from a File
//
//   - [NSData.InitWithContentsOfFile]: Initializes a data object with the content of the file at a given path.
//   - [NSData.InitWithContentsOfFileOptionsError]: Initializes a data object with the content of the file at a given path.
//
// # Writing Data to a File
//
//   - [NSData.WriteToFileAtomically]: Writes the data object’s bytes to the file specified by a given path.
//   - [NSData.WriteToFileOptionsError]: Writes the data object’s bytes to the file specified by a given path.
//   - [NSData.WriteToURLAtomically]: Writes the data object’s bytes to the location specified by a given URL.
//   - [NSData.WriteToURLOptionsError]: Writes the data object’s bytes to the location specified by a given URL.
//
// # Encoding and Decoding Base64 Representations
//
//   - [NSData.Base64EncodedDataWithOptions]: Creates a Base64, UTF-8 encoded data object from the string using the given options.
//   - [NSData.Base64EncodedStringWithOptions]: Creates a Base64 encoded string from the string using the given options.
//
// # Accessing Underlying Bytes
//
//   - [NSData.Bytes]: A pointer to the data object’s contents.
//   - [NSData.EnumerateByteRangesUsingBlock]: Enumerates each range of bytes in the data object using a block.
//   - [NSData.GetBytesLength]: Copies a number of bytes from the start of the data object into a given buffer.
//   - [NSData.GetBytesRange]: Copies a range of bytes from the data object into a given buffer.
//
// # Finding Data
//
//   - [NSData.SubdataWithRange]: Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range.
//   - [NSData.RangeOfDataOptionsRange]: Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options.
//
// # Testing Data
//
//   - [NSData.IsEqualToData]: Returns a Boolean value indicating whether this data object is the same as another.
//   - [NSData.Length]: The number of bytes contained by the data object.
//
// # Describing Data
//
//   - [NSData.Description]: A string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// # Compressing and Decompressing Data
//
//   - [NSData.CompressedDataUsingAlgorithmError]: Returns a new data object by compressing the data object’s bytes.
//   - [NSData.DecompressedDataUsingAlgorithmError]: Returns a new data object by decompressing data object’s bytes.
//   - [NSData.NSCompressionErrorMaximum]: The end of the range of error codes reserved for compression errors.
//   - [NSData.SetNSCompressionErrorMaximum]
//   - [NSData.NSCompressionErrorMinimum]: The start of the range of error codes reserved for compression errors.
//   - [NSData.SetNSCompressionErrorMinimum]
//   - [NSData.NSCompressionFailedError]: An error code value that indicates a failure to compress data using the provided algorithm.
//   - [NSData.SetNSCompressionFailedError]
//   - [NSData.NSDecompressionFailedError]: An error code value that indicates a failure to decompress data using the provided algorithm.
//   - [NSData.SetNSDecompressionFailedError]
//
// # Initializers
//
//   - [NSData.InitWithBase64EncodedStringOptions]
//   - [NSData.InitWithBase64EncodedDataOptions]
//   - [NSData.InitWithContentsOfURL]
//   - [NSData.InitWithContentsOfURLOptionsError]
//
// See: https://developer.apple.com/documentation/Foundation/NSData
type NSData struct {
	objectivec.Object
}

// NSDataFromID constructs a [NSData] from an objc.ID.
//
// A static byte buffer in memory.
func NSDataFromID(id objc.ID) NSData {
	return NSData{objectivec.Object{ID: id}}
}
// NOTE: NSData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSData] class.
//
// # Creating Data
//
//   - [INSData.InitWithBytesLength]: Initializes a data object filled with a given number of bytes copied from a given buffer.
//   - [INSData.InitWithBytesNoCopyLength]: Initializes a data object filled with a given number of bytes of data from a given buffer.
//   - [INSData.InitWithBytesNoCopyLengthDeallocator]: Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block.
//   - [INSData.InitWithBytesNoCopyLengthFreeWhenDone]: Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
//   - [INSData.InitWithData]: Initializes a data object with the contents of another data object.
//
// # Reading Data from a File
//
//   - [INSData.InitWithContentsOfFile]: Initializes a data object with the content of the file at a given path.
//   - [INSData.InitWithContentsOfFileOptionsError]: Initializes a data object with the content of the file at a given path.
//
// # Writing Data to a File
//
//   - [INSData.WriteToFileAtomically]: Writes the data object’s bytes to the file specified by a given path.
//   - [INSData.WriteToFileOptionsError]: Writes the data object’s bytes to the file specified by a given path.
//   - [INSData.WriteToURLAtomically]: Writes the data object’s bytes to the location specified by a given URL.
//   - [INSData.WriteToURLOptionsError]: Writes the data object’s bytes to the location specified by a given URL.
//
// # Encoding and Decoding Base64 Representations
//
//   - [INSData.Base64EncodedDataWithOptions]: Creates a Base64, UTF-8 encoded data object from the string using the given options.
//   - [INSData.Base64EncodedStringWithOptions]: Creates a Base64 encoded string from the string using the given options.
//
// # Accessing Underlying Bytes
//
//   - [INSData.Bytes]: A pointer to the data object’s contents.
//   - [INSData.EnumerateByteRangesUsingBlock]: Enumerates each range of bytes in the data object using a block.
//   - [INSData.GetBytesLength]: Copies a number of bytes from the start of the data object into a given buffer.
//   - [INSData.GetBytesRange]: Copies a range of bytes from the data object into a given buffer.
//
// # Finding Data
//
//   - [INSData.SubdataWithRange]: Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range.
//   - [INSData.RangeOfDataOptionsRange]: Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options.
//
// # Testing Data
//
//   - [INSData.IsEqualToData]: Returns a Boolean value indicating whether this data object is the same as another.
//   - [INSData.Length]: The number of bytes contained by the data object.
//
// # Describing Data
//
//   - [INSData.Description]: A string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// # Compressing and Decompressing Data
//
//   - [INSData.CompressedDataUsingAlgorithmError]: Returns a new data object by compressing the data object’s bytes.
//   - [INSData.DecompressedDataUsingAlgorithmError]: Returns a new data object by decompressing data object’s bytes.
//   - [INSData.NSCompressionErrorMaximum]: The end of the range of error codes reserved for compression errors.
//   - [INSData.SetNSCompressionErrorMaximum]
//   - [INSData.NSCompressionErrorMinimum]: The start of the range of error codes reserved for compression errors.
//   - [INSData.SetNSCompressionErrorMinimum]
//   - [INSData.NSCompressionFailedError]: An error code value that indicates a failure to compress data using the provided algorithm.
//   - [INSData.SetNSCompressionFailedError]
//   - [INSData.NSDecompressionFailedError]: An error code value that indicates a failure to decompress data using the provided algorithm.
//   - [INSData.SetNSDecompressionFailedError]
//
// # Initializers
//
//   - [INSData.InitWithBase64EncodedStringOptions]
//   - [INSData.InitWithBase64EncodedDataOptions]
//   - [INSData.InitWithContentsOfURL]
//   - [INSData.InitWithContentsOfURLOptionsError]
//
// See: https://developer.apple.com/documentation/Foundation/NSData
type INSData interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Creating Data

	// Initializes a data object filled with a given number of bytes copied from a given buffer.
	InitWithBytesLength(bytes []byte) NSData
	// Initializes a data object filled with a given number of bytes of data from a given buffer.
	InitWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSData
	// Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block.
	InitWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) NSData
	// Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
	InitWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSData
	// Initializes a data object with the contents of another data object.
	InitWithData(data INSData) NSData

	// Topic: Reading Data from a File

	// Initializes a data object with the content of the file at a given path.
	InitWithContentsOfFile(path string) NSData
	// Initializes a data object with the content of the file at a given path.
	InitWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSData, error)

	// Topic: Writing Data to a File

	// Writes the data object’s bytes to the file specified by a given path.
	WriteToFileAtomically(path string, useAuxiliaryFile bool) bool
	// Writes the data object’s bytes to the file specified by a given path.
	WriteToFileOptionsError(path string, writeOptionsMask NSDataWritingOptions) (bool, error)
	// Writes the data object’s bytes to the location specified by a given URL.
	WriteToURLAtomically(url INSURL, atomically bool) bool
	// Writes the data object’s bytes to the location specified by a given URL.
	WriteToURLOptionsError(url INSURL, writeOptionsMask NSDataWritingOptions) (bool, error)

	// Topic: Encoding and Decoding Base64 Representations

	// Creates a Base64, UTF-8 encoded data object from the string using the given options.
	Base64EncodedDataWithOptions(options NSDataBase64EncodingOptions) INSData
	// Creates a Base64 encoded string from the string using the given options.
	Base64EncodedStringWithOptions(options NSDataBase64EncodingOptions) string

	// Topic: Accessing Underlying Bytes

	// A pointer to the data object’s contents.
	Bytes() unsafe.Pointer
	// Enumerates each range of bytes in the data object using a block.
	EnumerateByteRangesUsingBlock(block unsafe.Pointer)
	// Copies a number of bytes from the start of the data object into a given buffer.
	GetBytesLength(buffer unsafe.Pointer, length uint)
	// Copies a range of bytes from the data object into a given buffer.
	GetBytesRange(buffer unsafe.Pointer, range_ NSRange)

	// Topic: Finding Data

	// Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range.
	SubdataWithRange(range_ NSRange) INSData
	// Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options.
	RangeOfDataOptionsRange(dataToFind INSData, mask NSDataSearchOptions, searchRange NSRange) NSRange

	// Topic: Testing Data

	// Returns a Boolean value indicating whether this data object is the same as another.
	IsEqualToData(other INSData) bool
	// The number of bytes contained by the data object.
	Length() uint

	// Topic: Describing Data

	// A string that contains a hexadecimal representation of the data object’s contents in a property list format.
	Description() string

	// Topic: Compressing and Decompressing Data

	// Returns a new data object by compressing the data object’s bytes.
	CompressedDataUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (INSData, error)
	// Returns a new data object by decompressing data object’s bytes.
	DecompressedDataUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (INSData, error)
	// The end of the range of error codes reserved for compression errors.
	NSCompressionErrorMaximum() int
	SetNSCompressionErrorMaximum(value int)
	// The start of the range of error codes reserved for compression errors.
	NSCompressionErrorMinimum() int
	SetNSCompressionErrorMinimum(value int)
	// An error code value that indicates a failure to compress data using the provided algorithm.
	NSCompressionFailedError() int
	SetNSCompressionFailedError(value int)
	// An error code value that indicates a failure to decompress data using the provided algorithm.
	NSDecompressionFailedError() int
	SetNSDecompressionFailedError(value int)

	// Topic: Initializers

	InitWithBase64EncodedStringOptions(base64String string, options NSDataBase64DecodingOptions) NSData
	InitWithBase64EncodedDataOptions(base64Data INSData, options NSDataBase64DecodingOptions) NSData
	InitWithContentsOfURL(url INSURL) NSData
	InitWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSData, error)
}

// Init initializes the instance.
func (d NSData) Init() NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSData) Autorelease() NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSData creates a new NSData instance.
func NewNSData() NSData {
	class := getNSDataClass()
	rv := objc.Send[NSData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-4t5yq
func NewDataWithBase64EncodedDataOptions(base64Data INSData, options NSDataBase64DecodingOptions) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return NSDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-3ksry
func NewDataWithBase64EncodedStringOptions(base64String string, options NSDataBase64DecodingOptions) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedString:options:"), objc.String(base64String), options)
	return NSDataFromID(rv)
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
func NewDataWithBase64Encoding(base64String string) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64Encoding:"), objc.String(base64String))
	return NSDataFromID(rv)
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
func NewDataWithBytesLength(bytes []byte) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
	return NSDataFromID(rv)
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
func NewDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	return NSDataFromID(rv)
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
func NewDataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	return NSDataFromID(rv)
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
func NewDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return NSDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDataWithCoder(coder INSCoder) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDataFromID(rv)
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
func NewDataWithContentsOfFile(path string) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSDataFromID(rv)
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
func NewDataWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:options:error:"), objc.String(path), readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:)
func NewDataWithContentsOfURL(url INSURL) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:options:)
func NewDataWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil
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
func NewDataWithData(data INSData) NSData {
	instance := getNSDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSDataFromID(rv)
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
func (d NSData) InitWithBytesLength(bytes []byte) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
	return rv
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
func (d NSData) InitWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	return rv
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
func (d NSData) InitWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	return rv
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
func (d NSData) InitWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
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
func (d NSData) InitWithData(data INSData) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithData:"), data)
	return rv
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
func (d NSData) InitWithContentsOfFile(path string) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return rv
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
func (d NSData) InitWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithContentsOfFile:options:error:"), objc.String(path), readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

// Writes the data object’s bytes to the file specified by a given path.
//
// path: The location to which to write the receiver’s bytes. If `path` contains a
// tilde (~) character, you must expand it with [StringByExpandingTildeInPath]
// before invoking this method.
//
// useAuxiliaryFile: If [true], the data is written to a backup file, and then—assuming no
// errors occur—the backup file is renamed to the name specified by `path`;
// otherwise, the data is written directly to `path`.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the operation succeeds, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method may not be appropriate when writing to publicly accessible
// files. To securely write data to a public location, use [NSFileHandle]
// instead. For more information, see [Securing File Operations] in [Secure
// Coding Guide].
//
// [Secure Coding Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Introduction.html#//apple_ref/doc/uid/TP40002415
// [Securing File Operations]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Articles/RaceConditions.html#//apple_ref/doc/uid/TP40002585-SW9
//
// See: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:atomically:)
func (d NSData) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("writeToFile:atomically:"), objc.String(path), useAuxiliaryFile)
	return rv
}

// Writes the data object’s bytes to the file specified by a given path.
//
// path: The location to which to write the receiver’s bytes.
//
// writeOptionsMask: A mask that specifies options for writing the data. Constant components are
// described in [NSData.WritingOptions].
// //
// [NSData.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
//
// # Discussion
// 
// This method may not be appropriate when writing to publicly accessible
// files. To securely write data to a public location, use [NSFileHandle]
// instead. For more information, see [Securing File Operations] in [Secure
// Coding Guide].
//
// [Secure Coding Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Introduction.html#//apple_ref/doc/uid/TP40002415
// [Securing File Operations]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Articles/RaceConditions.html#//apple_ref/doc/uid/TP40002585-SW9
//
// See: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:options:)
func (d NSData) WriteToFileOptionsError(path string, writeOptionsMask NSDataWritingOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToFile:options:error:"), objc.String(path), writeOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToFile:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Writes the data object’s bytes to the location specified by a given URL.
//
// url: The location to which to write the receiver’s bytes. Only `//` URLs are
// supported.
//
// atomically: If [true], the data is written to a backup location, and then—assuming no
// errors occur—the backup location is renamed to the name specified by
// `aURL`; otherwise, the data is written directly to `aURL`. `atomically` is
// ignored if `aURL` is not of a type the supports atomic writes.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the operation succeeds, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Since at present only `//` URLs are supported, there is no difference
// between this method and [WriteToFileAtomically], except for the type of the
// first argument.
// 
// This method may not be appropriate when writing to publicly accessible
// files. To securely write data to a public location, use [NSFileHandle]
// instead. For more information, see [Securing File Operations] in [Secure
// Coding Guide].
//
// [Secure Coding Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Introduction.html#//apple_ref/doc/uid/TP40002415
// [Securing File Operations]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Articles/RaceConditions.html#//apple_ref/doc/uid/TP40002585-SW9
//
// See: https://developer.apple.com/documentation/Foundation/NSData/write(to:atomically:)
func (d NSData) WriteToURLAtomically(url INSURL, atomically bool) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}

// Writes the data object’s bytes to the location specified by a given URL.
//
// url: The location to which to write the receiver’s bytes.
//
// writeOptionsMask: A mask that specifies options for writing the data. Constant components are
// described in [NSData.WritingOptions].
// //
// [NSData.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
//
// # Discussion
// 
// Since at present only `//` URLs are supported, there is no difference
// between this method and [WriteToFileOptionsError], except for the type of
// the first argument.
// 
// This method may not be appropriate when writing to publicly accessible
// files. To securely write data to a public location, use [NSFileHandle]
// instead. For more information, see[Securing File Operations] in [Secure
// Coding Guide].
//
// [Secure Coding Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Introduction.html#//apple_ref/doc/uid/TP40002415
// [Securing File Operations]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/SecureCodingGuide/Articles/RaceConditions.html#//apple_ref/doc/uid/TP40002585-SW9
//
// See: https://developer.apple.com/documentation/Foundation/NSData/write(to:options:)
func (d NSData) WriteToURLOptionsError(url INSURL, writeOptionsMask NSDataWritingOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:options:error:"), url, writeOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a Base64, UTF-8 encoded data object from the string using the given
// options.
//
// options: A mask that specifies options for Base64 encoding the data. Possible values
// are given in [NSData.Base64EncodingOptions].
// //
// [NSData.Base64EncodingOptions]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
//
// # Return Value
// 
// A Base64, UTF-8 encoded data object.
//
// # Discussion
// 
// By default, no line endings are inserted.
// 
// If you specify one of the line length options
// ([DataBase64Encoding64CharacterLineLength] or
// [DataBase64Encoding76CharacterLineLength]) but don’t specify the kind of
// line ending to insert, the default line ending is Carriage Return + Line
// Feed.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedData(options:)
func (d NSData) Base64EncodedDataWithOptions(options NSDataBase64EncodingOptions) INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("base64EncodedDataWithOptions:"), options)
	return NSDataFromID(rv)
}

// Creates a Base64 encoded string from the string using the given options.
//
// options: A mask that specifies options for Base-64 encoding the data. Possible
// values are given in [NSData.Base64EncodingOptions].
// //
// [NSData.Base64EncodingOptions]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
//
// # Return Value
// 
// A Base64 encoded string.
//
// # Discussion
// 
// By default, no line endings are inserted.
// 
// If you specify one of the line length options
// ([DataBase64Encoding64CharacterLineLength] or
// [DataBase64Encoding76CharacterLineLength]) but don’t specify the kind of
// line ending to insert, the default line ending is Carriage Return + Line
// Feed.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedString(options:)
func (d NSData) Base64EncodedStringWithOptions(options NSDataBase64EncodingOptions) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("base64EncodedStringWithOptions:"), options)
	return NSStringFromID(rv).String()
}

// Enumerates each range of bytes in the data object using a block.
//
// block: The block to apply to byte ranges in the array.
// 
// The block takes three arguments:
// 
// bytes: The bytes for the current range. This pointer is valid until the
// data object is deallocated. byteRange: The range of the current data bytes.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further processing of the data. The stop argument is an out-only
// argument. You should only ever set this Boolean to [true] within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The enumeration block is called once for each contiguous region of memory
// in the receiver (once total for a contiguous [NSData] object), until either
// all bytes have been enumerated, or the `stop` parameter is set to [true].
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSData/enumerateBytes(_:)
func (d NSData) EnumerateByteRangesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateByteRangesUsingBlock:"), block)
}

// Copies a number of bytes from the start of the data object into a given
// buffer.
//
// buffer: A buffer into which to copy data.
//
// length: The number of bytes from the start of the receiver’s data to copy to
// `buffer`.
//
// # Discussion
// 
// The number of bytes copied is the smaller of the `length` parameter and the
// [Length] of the data encapsulated in the object.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:length:)
func (d NSData) GetBytesLength(buffer unsafe.Pointer, length uint) {
	objc.Send[objc.ID](d.ID, objc.Sel("getBytes:length:"), buffer, length)
}

// Copies a range of bytes from the data object into a given buffer.
//
// buffer: A buffer into which to copy data.
//
// range: The range of bytes in the receiver’s data to copy to `buffer`. The range
// must lie within the range of bytes of the receiver’s data.
//
// # Discussion
// 
// If `range` isn’t within the receiver’s range of bytes, an
// [rangeException] is raised.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:range:)
func (d NSData) GetBytesRange(buffer unsafe.Pointer, range_ NSRange) {
	objc.Send[objc.ID](d.ID, objc.Sel("getBytes:range:"), buffer, range_)
}

// Returns a new data object containing the data object’s bytes that fall
// within the limits specified by a given range.
//
// range: The range in the receiver from which to get the data. If this range is not
// within the data object’s range of bytes, [rangeException] is raised.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// A data object containing the receiver’s bytes that fall within the limits
// specified by `range`.
//
// # Discussion
// 
// A sample using this method can be found in [Working With Binary Data].
//
// [Working With Binary Data]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/BinaryData/Tasks/WorkingBinaryData.html#//apple_ref/doc/uid/20000717
//
// See: https://developer.apple.com/documentation/Foundation/NSData/subdata(with:)
func (d NSData) SubdataWithRange(range_ NSRange) INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("subdataWithRange:"), range_)
	return NSDataFromID(rv)
}

// Finds and returns the range of the first occurrence of the given data,
// within the given range, subject to given options.
//
// dataToFind: The data for which to search.
//
// mask: A mask specifying search options. The [NSData.SearchOptions] options may be
// specified singly or by combining them with the C bitwise [OR] operator.
// //
// [NSData.SearchOptions]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
//
// searchRange: The range within the receiver in which to search for `dataToFind`. If this
// range is not within the data object’s range of bytes, [rangeException] is
// raised.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Return Value
// 
// An [NSRange] structure giving the location and length of `dataToFind`
// within `searchRange`, modulo the options in `mask`. The range returned is
// relative to the start of the searched data, not the passed-in search range.
// Returns `{``NSNotFound``, 0}` if `dataToFind` is not found or is empty.
//
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// See: https://developer.apple.com/documentation/Foundation/NSData/range(of:options:in:)
func (d NSData) RangeOfDataOptionsRange(dataToFind INSData, mask NSDataSearchOptions, searchRange NSRange) NSRange {
	rv := objc.Send[NSRange](d.ID, objc.Sel("rangeOfData:options:range:"), dataToFind, mask, searchRange)
	return NSRange(rv)
}

// Returns a Boolean value indicating whether this data object is the same as
// another.
//
// other: The data object with which to compare the receiver.
//
// # Return Value
// 
// [true] if the contents of `otherData` are equal to the contents of the
// receiver, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two data objects are equal if they hold the same number of bytes, and if
// the bytes at the same position in the objects are the same.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/isEqual(to:)
func (d NSData) IsEqualToData(other INSData) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEqualToData:"), other)
	return rv
}

// Returns a new data object by compressing the data object’s bytes.
//
// algorithm: An algorithm used to compress the data. For a list of available algorithms,
// see [NSData.CompressionAlgorithm].
// //
// [NSData.CompressionAlgorithm]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
//
// # Return Value
// 
// An [NSData] instance that contains the compressed buffer data.
//
// # Discussion
// 
// Use this method to compress in-memory data when you want to reduce memory
// usage and can afford the time to compress and decompress it. If your data
// object is already in a compressed format, such as media formats like JPEG
// images or AAC audio, additional compression may provide minimal or no
// reduction in memory usage.
// 
// To restore this data, use [DecompressedDataUsingAlgorithmError], and
// specify the algorithm originally used to compress the data.
// 
// The following example shows how to compress the data from a string and
// prints the sizes of the data instances to illustrate the amount of
// compression:
//
// See: https://developer.apple.com/documentation/Foundation/NSData/compressed(using:)
func (d NSData) CompressedDataUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

// Returns a new data object by decompressing data object’s bytes.
//
// algorithm: An algorithm used to decompress the data. For a list of available
// algorithms, see [NSData.CompressionAlgorithm].
// //
// [NSData.CompressionAlgorithm]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
//
// # Return Value
// 
// An [NSData] instance that contains the decompressed buffer data.
//
// # Discussion
// 
// Use this method to inflate in-memory data when you need uncompressed bytes.
// Specify the same algorithm used to compress the data to successfully
// decompress it.
// 
// The following example shows how to create a new [NSData] instance from data
// compressed with the [DataCompressionAlgorithmZlib] algorithm:
//
// See: https://developer.apple.com/documentation/Foundation/NSData/decompressed(using:)
func (d NSData) DecompressedDataUsingAlgorithmError(algorithm NSDataCompressionAlgorithm) (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-3ksry
func (d NSData) InitWithBase64EncodedStringOptions(base64String string, options NSDataBase64DecodingOptions) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBase64EncodedString:options:"), objc.String(base64String), options)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-4t5yq
func (d NSData) InitWithBase64EncodedDataOptions(base64Data INSData, options NSDataBase64DecodingOptions) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:)
func (d NSData) InitWithContentsOfURL(url INSURL) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:options:)
func (d NSData) InitWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSData) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (d NSData) InitWithCoder(coder INSCoder) NSData {
	rv := objc.Send[NSData](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates an empty data object.
//
// # Discussion
// 
// This method is declared primarily for the use of mutable subclasses of
// [NSData].
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// See: https://developer.apple.com/documentation/Foundation/NSData/data
func (_NSDataClass NSDataClass) Data() NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("data"))
	return NSDataFromID(rv)
}

// Creates a data object containing a given number of bytes copied from a
// given buffer.
//
// bytes: A buffer containing data for the new object.
//
// length: The number of bytes to copy from `bytes`. This value must not exceed the
// length of `bytes`.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytes:length:
func (_NSDataClass NSDataClass) DataWithBytesLength(bytes []byte) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
	return NSDataFromID(rv)
}

// Creates a data object that holds a given number of bytes from a given
// buffer.
//
// bytes: A buffer containing data for the new object. `bytes` must point to a memory
// block allocated with `malloc`.
//
// length: The number of bytes to hold from `bytes`. This value must not exceed the
// length of `bytes`.
//
// # Discussion
// 
// The returned object takes ownership of the `bytes` pointer and frees it on
// deallocation. Therefore, `bytes` must point to a memory block allocated
// with `malloc`.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:
func (_NSDataClass NSDataClass) DataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithBytesNoCopy:length:"), bytes, length)
	return NSDataFromID(rv)
}

// Creates a data object that holds a given number of bytes from a given
// buffer.
//
// bytes: A buffer containing data for the new object. If `freeWhenDone` is [true],
// `bytes` must point to a memory block allocated with `malloc`.
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
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:freeWhenDone:
func (_NSDataClass NSDataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return NSDataFromID(rv)
}

// Creates a data object by reading every byte from the file at a given path.
//
// path: The absolute path of the file from which to read data.
//
// # Discussion
// 
// This method returns `nil` if the data object could not be created. If you
// need to know the reason for failure, use
// [DataWithContentsOfFileOptionsError].
// 
// This method is equivalent to calling [DataWithContentsOfFileOptionsError]
// and passing no options.
// 
// A sample using this method can be found in [Working With Binary Data].
//
// [Working With Binary Data]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/BinaryData/Tasks/WorkingBinaryData.html#//apple_ref/doc/uid/20000717
//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:
func (_NSDataClass NSDataClass) DataWithContentsOfFile(path string) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithContentsOfFile:"), objc.String(path))
	return NSDataFromID(rv)
}

// Creates a data object by reading every byte from the file at a given path.
//
// path: The absolute path of the file from which to read data.
//
// readOptionsMask: A mask that specifies options for reading the data. Constant components are
// described in [NSData.ReadingOptions].
// //
// [NSData.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
//
// errorPtr: If an error occurs, upon return contains an error object that describes the
// problem.
//
// # Discussion
// 
// This method returns `nil` if the data object could not be created. In this
// case, `errorPtr` will contain an [NSError] indicating the problem.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:options:error:
func (_NSDataClass NSDataClass) DataWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithContentsOfFile:options:error:"), objc.String(path), readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfURL:
func (_NSDataClass NSDataClass) DataWithContentsOfURL(url INSURL) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithContentsOfURL:"), url)
	return NSDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfURL:options:error:
func (_NSDataClass NSDataClass) DataWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithContentsOfURL:options:error:"), url, readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

// Creates a data object containing the contents of another data object.
//
// data: A data object.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/dataWithData:
func (_NSDataClass NSDataClass) DataWithData(data INSData) NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSDataClass.class), objc.Sel("dataWithData:"), data)
	return NSDataFromID(rv)
}

// A pointer to the data object’s contents.
//
// # Discussion
// 
// If the [Length] of the [NSData] object is 0, this property returns `nil`.
// 
// For an immutable data object, the returned pointer is valid until the data
// object is deallocated. For a mutable data object, the returned pointer is
// valid until the data object is deallocated or the data is mutated.
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// See: https://developer.apple.com/documentation/Foundation/NSData/bytes
func (d NSData) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("bytes"))
	return rv
}

// The number of bytes contained by the data object.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/length
func (d NSData) Length() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("length"))
	return rv
}

// A string that contains a hexadecimal representation of the data object’s
// contents in a property list format.
//
// See: https://developer.apple.com/documentation/Foundation/NSData/description
func (d NSData) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

// The end of the range of error codes reserved for compression errors.
//
// See: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (d NSData) NSCompressionErrorMaximum() int {
	rv := objc.Send[int](d.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}
func (d NSData) SetNSCompressionErrorMaximum(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}

// The start of the range of error codes reserved for compression errors.
//
// See: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (d NSData) NSCompressionErrorMinimum() int {
	rv := objc.Send[int](d.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}
func (d NSData) SetNSCompressionErrorMinimum(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}

// An error code value that indicates a failure to compress data using the
// provided algorithm.
//
// See: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (d NSData) NSCompressionFailedError() int {
	rv := objc.Send[int](d.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}
func (d NSData) SetNSCompressionFailedError(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSCompressionFailedError:"), value)
}

// An error code value that indicates a failure to decompress data using the
// provided algorithm.
//
// See: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (d NSData) NSDecompressionFailedError() int {
	rv := objc.Send[int](d.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}
func (d NSData) SetNSDecompressionFailedError(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSMutableCopying
			

			// Protocol methods for NSSecureCoding
			

