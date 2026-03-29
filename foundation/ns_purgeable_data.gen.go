// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPurgeableData] class.
var (
	_NSPurgeableDataClass     NSPurgeableDataClass
	_NSPurgeableDataClassOnce sync.Once
)

func getNSPurgeableDataClass() NSPurgeableDataClass {
	_NSPurgeableDataClassOnce.Do(func() {
		_NSPurgeableDataClass = NSPurgeableDataClass{class: objc.GetClass("NSPurgeableData")}
	})
	return _NSPurgeableDataClass
}

// GetNSPurgeableDataClass returns the class object for NSPurgeableData.
func GetNSPurgeableDataClass() NSPurgeableDataClass {
	return getNSPurgeableDataClass()
}

type NSPurgeableDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPurgeableDataClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPurgeableDataClass) Alloc() NSPurgeableData {
	rv := objc.Send[NSPurgeableData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mutable data object containing bytes that can be discarded when they’re
// no longer needed.
//
// # Overview
// 
// [NSPurgeableData] objects inherit their creation methods from their
// superclass, [NSMutableData] while providing a default implementation of the
// [NSDiscardableContent] protocol.
// 
// All [NSPurgeableData] objects begin “accessed” to ensure that they are
// not instantly discarded. The [NSPurgeableData.BeginContentAccess] method marks the
// object’s bytes as “accessed,” thus protecting them from being
// discarded, and must be called before accessing the object, or else an
// exception will be raised. This method returns [true] if the bytes have not
// been discarded and if they have been successfully marked as “accessed”.
// Any method that directly or indirectly accesses these bytes or their length
// when they are not “accessed” will raise an exception. When you are done
// with the data, call [NSPurgeableData.EndContentAccess] to allow them to be discarded in
// order to quickly free up memory.
// 
// You may use these objects by themselves, and do not necessarily have to use
// them in conjunction with [NSCache] to get the purging behavior. The
// [NSCache] class incorporates a caching mechanism with some auto-removal
// policies to ensure that its memory footprint does not get too large.
// 
// [NSPurgeableData] objects should not be used as keys in hashing-based
// collections, because the value of the bytes pointer can change after every
// mutation of the data.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSPurgeableData
type NSPurgeableData struct {
	NSMutableData
}

// NSPurgeableDataFromID constructs a [NSPurgeableData] from an objc.ID.
//
// A mutable data object containing bytes that can be discarded when they’re
// no longer needed.
func NSPurgeableDataFromID(id objc.ID) NSPurgeableData {
	return NSPurgeableData{NSMutableData: NSMutableDataFromID(id)}
}
// NOTE: NSPurgeableData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPurgeableData] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPurgeableData
type INSPurgeableData interface {
	INSMutableData
	NSDiscardableContent
}

// Init initializes the instance.
func (p NSPurgeableData) Init() NSPurgeableData {
	rv := objc.Send[NSPurgeableData](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPurgeableData) Autorelease() NSPurgeableData {
	rv := objc.Send[NSPurgeableData](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPurgeableData creates a new NSPurgeableData instance.
func NewNSPurgeableData() NSPurgeableData {
	class := getNSPurgeableDataClass()
	rv := objc.Send[NSPurgeableData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-4t5yq
func NewPurgeableDataWithBase64EncodedDataOptions(base64Data INSData, options NSDataBase64DecodingOptions) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return NSPurgeableDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoded:options:)-3ksry
func NewPurgeableDataWithBase64EncodedStringOptions(base64String string, options NSDataBase64DecodingOptions) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64EncodedString:options:"), objc.String(base64String), options)
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithBase64Encoding(base64String string) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBase64Encoding:"), objc.String(base64String))
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithBytesLength(bytes []byte) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithCapacity(capacity uint) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	return NSPurgeableDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPurgeableDataWithCoder(coder INSCoder) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithContentsOfFile(path string) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithContentsOfFileOptionsError(path string, readOptionsMask NSDataReadingOptions) (NSPurgeableData, error) {
	var errorPtr objc.ID
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:options:error:"), objc.String(path), readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPurgeableData{}, NSErrorFrom(errorPtr)
	}
	return NSPurgeableDataFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:)
func NewPurgeableDataWithContentsOfURL(url INSURL) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSPurgeableDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOf:options:)
func NewPurgeableDataWithContentsOfURLOptionsError(url INSURL, readOptionsMask NSDataReadingOptions) (NSPurgeableData, error) {
	var errorPtr objc.ID
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPurgeableData{}, NSErrorFrom(errorPtr)
	}
	return NSPurgeableDataFromID(rv), nil
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
func NewPurgeableDataWithData(data INSData) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSPurgeableDataFromID(rv)
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
func NewPurgeableDataWithLength(length uint) NSPurgeableData {
	instance := getNSPurgeableDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLength:"), length)
	return NSPurgeableDataFromID(rv)
}

// Returns a Boolean value indicating whether the discardable contents are
// still available and have been successfully accessed.
//
// # Return Value
// 
// YES if the discardable contents are still available and have now been
// successfully accessed; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// Call this method if the object’s memory is needed or is about to be used.
// This method increments the counter variable, thus protecting the object’s
// memory from possibly being discarded. The implementing class may decide
// that this method will try to recreate the contents if they have been
// discarded and return YES if the re-creation was successful. Implementors of
// this protocol should raise exceptions if the [NSDiscardableContent] objects
// are used when the `beginContentAccess` method has not been called on them.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/beginContentAccess()
func (p NSPurgeableData) BeginContentAccess() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("beginContentAccess"))
	return rv
}
// Called to discard the contents of the receiver if the value of the accessed
// counter is 0.
//
// # Discussion
// 
// This method should only discard the contents of the object if the value of
// the accessed counter is 0. Otherwise, it should do nothing.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/discardContentIfPossible()
func (p NSPurgeableData) DiscardContentIfPossible() {
	objc.Send[objc.ID](p.ID, objc.Sel("discardContentIfPossible"))
}
// Called if the discardable contents are no longer being accessed.
//
// # Discussion
// 
// This method decrements the counter variable of the object, which will
// usually bring the value of the counter variable back down to 0, which
// allows the discardable contents of the object to be thrown away if
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/endContentAccess()
func (p NSPurgeableData) EndContentAccess() {
	objc.Send[objc.ID](p.ID, objc.Sel("endContentAccess"))
}
// Returns a Boolean value indicating whether the content has been discarded.
//
// # Return Value
// 
// [true] if the content has been discarded; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/isContentDiscarded()
func (p NSPurgeableData) IsContentDiscarded() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isContentDiscarded"))
	return rv
}

			// Protocol methods for NSDiscardableContent
			

