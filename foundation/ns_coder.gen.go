// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCoder] class.
var (
	_NSCoderClass     NSCoderClass
	_NSCoderClassOnce sync.Once
)

func getNSCoderClass() NSCoderClass {
	_NSCoderClassOnce.Do(func() {
		_NSCoderClass = NSCoderClass{class: objc.GetClass("NSCoder")}
	})
	return _NSCoderClass
}

// GetNSCoderClass returns the class object for NSCoder.
func GetNSCoderClass() NSCoderClass {
	return getNSCoderClass()
}

type NSCoderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCoderClass) Alloc() NSCoder {
	rv := objc.Send[NSCoder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class that serves as the basis for objects that enable
// archiving and distribution of other objects.
//
// # Overview
// 
// [NSCoder] declares the interface used by concrete subclasses to transfer
// objects and other values between memory and some other format. This
// capability provides the basis for archiving (storing objects and data on
// disk) and distribution (copying objects and data items between different
// processes or threads). The concrete subclasses provided by Foundation for
// these purposes are [NSArchiver], [NSUnarchiver], [NSKeyedArchiver],
// [NSKeyedUnarchiver], and [NSPortCoder]. Concrete subclasses of [NSCoder]
// are “coder classes”, and instances of these classes are “coder
// objects” (or simply “coders”). A coder that can only encode values is
// an “encoder”, and one that can only decode values is a “decoder”.
// 
// [NSCoder] operates on objects, scalars, C arrays, structures, strings, and
// on pointers to these types. It doesn’t handle types whose implementation
// varies across platforms, such as `union`, `void *`, function pointers, and
// long chains of pointers. A coder stores object type information along with
// the data, so an object decoded from a stream of bytes is normally of the
// same class as the object that was originally encoded into the stream. An
// object can change its class when encoded, however; this is described in
// [Archives and Serializations Programming Guide].
// 
// The AVFoundation framework adds methods to the [NSCoder] class to make it
// easier to create archives including Core Media time structures, and extract
// Core Media time structure from archives.
// 
// # Subclassing Notes
// 
// For details of how to create a subclass of [NSCoder], see [Subclassing
// NSCoder] in [Archives and Serializations Programming Guide].
//
// [Archives and Serializations Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Archiving.html#//apple_ref/doc/uid/10000047i
// [NSArchiver]: https://developer.apple.com/documentation/Foundation/NSArchiver
// [NSPortCoder]: https://developer.apple.com/documentation/Foundation/NSPortCoder
// [NSUnarchiver]: https://developer.apple.com/documentation/Foundation/NSUnarchiver
// [Subclassing NSCoder]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Articles/subclassing.html#//apple_ref/doc/uid/20000951
//
// # Inspecting a Coder
//
//   - [NSCoder.AllowsKeyedCoding]: A Boolean value that indicates whether the receiver supports keyed coding of objects.
//   - [NSCoder.ContainsValueForKey]: Returns a Boolean value that indicates whether an encoded value is available for a string.
//   - [NSCoder.DecodingFailurePolicy]: The action the coder should take when decoding fails.
//
// # Encoding General Data
//
//   - [NSCoder.EncodeArrayOfObjCTypeCountAt]: Encodes an array of the given Objective-C type, provided the number of items and a pointer.
//   - [NSCoder.EncodeBoolForKey]: Encodes a Boolean value and associates it with the string `key`.
//   - [NSCoder.EncodeBycopyObject]: An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded.
//   - [NSCoder.EncodeByrefObject]: An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded.
//   - [NSCoder.EncodeBytesLength]: Encodes a buffer of data of an unspecified type.
//   - [NSCoder.EncodeBytesLengthForKey]: Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
//   - [NSCoder.EncodeConditionalObject]: An encoding method for subclasses to override to conditionally encode an object, preserving common references to it.
//   - [NSCoder.EncodeConditionalObjectForKey]: An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
//   - [NSCoder.EncodeDataObject]: Encodes a given data object.
//   - [NSCoder.EncodeDoubleForKey]: Encodes a double-precision floating point value and associates it with the string key.
//   - [NSCoder.EncodeFloatForKey]: Encodes a floating point value and associates it with the string key.
//   - [NSCoder.EncodeIntForKey]: Encodes a C integer value and associates it with the string key.
//   - [NSCoder.EncodeIntegerForKey]: Encodes an integer value and associates it with the string key.
//   - [NSCoder.EncodeInt32ForKey]: Encodes a 32-bit integer value and associates it with the string key.
//   - [NSCoder.EncodeInt64ForKey]: Encodes a 64-bit integer value and associates it with the string key.
//   - [NSCoder.EncodeObject]: Encodes an object.
//   - [NSCoder.EncodeObjectForKey]: Encodes an object and associates it with the string key.
//   - [NSCoder.EncodePoint]: Encodes a point.
//   - [NSCoder.EncodePointForKey]: Encodes a point and associates it with the string key.
//   - [NSCoder.EncodePropertyList]: Encodes a property list.
//   - [NSCoder.EncodeRect]: Encodes a rectangle structure.
//   - [NSCoder.EncodeRectForKey]: Encodes a rectangle structure and associates it with the string key.
//   - [NSCoder.EncodeRootObject]: An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object.
//   - [NSCoder.EncodeSize]: Encodes a size structure.
//   - [NSCoder.EncodeSizeForKey]: Encodes a size structure and associates it with the given string key.
//   - [NSCoder.EncodeValueOfObjCTypeAt]: Encodes a value of the given type at the given address.
//
// # Encoding Core Media Time Structures
//
//   - [NSCoder.EncodeCMTimeForKey]: Encodes a given Core Media time structure and associates it with a specified key.
//   - [NSCoder.EncodeCMTimeRangeForKey]: Encodes a given Core Media time range structure and associates it with a specified key.
//   - [NSCoder.EncodeCMTimeMappingForKey]: Encodes a given Core Media time mapping structure and associates it with a specified key.
//
// # Secure Coding
//
//   - [NSCoder.RequiresSecureCoding]: Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//   - [NSCoder.AllowedClasses]: The set of coded classes allowed for secure coding.
//
// # Decoding General Data
//
//   - [NSCoder.DecodeArrayOfObjCTypeCountAt]: Decodes an array of `count` items, whose Objective-C type is given by `itemType`.
//   - [NSCoder.DecodeBoolForKey]: Decodes and returns a boolean value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu>) and associated with the string `key`.
//   - [NSCoder.DecodeBytesForKeyReturnedLength]: Decodes a buffer of data that was previously encoded with [encodeBytes(_:length:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)>) and associated with the string `key`.
//   - [NSCoder.DecodeBytesWithReturnedLength]: Decodes a buffer of data whose types are unspecified.
//   - [NSCoder.DecodeDataObject]: Decodes and returns an [NSData] object that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-1qd1e>). Subclasses must override this method.
//   - [NSCoder.DecodeDoubleForKey]: Decodes and returns a double value that was previously encoded with either [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
//   - [NSCoder.DecodeFloatForKey]: Decodes and returns a float value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
//   - [NSCoder.DecodeIntForKey]: Decodes and returns an int value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [NSCoder.DecodeIntegerForKey]: Decodes and returns an NSInteger value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [NSCoder.DecodeInt32ForKey]: Decodes and returns a 32-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [NSCoder.DecodeInt64ForKey]: Decodes and returns a 64-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [NSCoder.DecodeObject]: Decodes and returns an object that was previously encoded with any of the `encode…Object` methods.
//   - [NSCoder.DecodeObjectForKey]: Decodes and returns a previously-encoded object that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu>) or [encodeConditionalObject(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)>) and associated with the string `key`.
//   - [NSCoder.DecodePoint]: Decodes and returns an NSPoint structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-75jv4>).
//   - [NSCoder.DecodePointForKey]: Decodes and returns an NSPoint structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif>).
//   - [NSCoder.DecodePropertyList]: Decodes a property list that was previously encoded with [encodePropertyList(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodePropertyList(_:)>).
//   - [NSCoder.DecodeRect]: Decodes and returns an NSRect structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-3c1wz>).
//   - [NSCoder.DecodeRectForKey]: Decodes and returns an NSRect structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx>).
//   - [NSCoder.DecodeSize]: Decodes and returns an NSSize structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-82i7c>).
//   - [NSCoder.DecodeSizeForKey]: Decodes and returns an NSSize structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu>).
//   - [NSCoder.DecodeValueOfObjCTypeAtSize]: Decodes a single value of a known type from the specified data buffer.
//   - [NSCoder.DecodePropertyListForKey]: Returns a decoded property list for the specified key.
//
// # Decoding Core Media Time Structures
//
//   - [NSCoder.DecodeCMTimeForKey]: Returns the Core Media time structure associated with a given key.
//   - [NSCoder.DecodeCMTimeRangeForKey]: Returns the Core Media time range structure associated with a given key.
//   - [NSCoder.DecodeCMTimeMappingForKey]: Returns the Core Media time mapping structure associated with a given key.
//
// # Managing Decode Errors
//
//   - [NSCoder.FailWithError]: Signals to this coder that the decode operation has failed.
//   - [NSCoder.Error]: An error in the top-level encode.
//
// # Getting Version Information
//
//   - [NSCoder.SystemVersion]: The system version in effect for the archive.
//   - [NSCoder.VersionForClassName]: This method is present for historical reasons and is not used with keyed archivers.
//
// # Error Codes
//
//   - [NSCoder.NSCoderErrorMaximum]: The end of the range of error codes reserved for coder errors.
//   - [NSCoder.SetNSCoderErrorMaximum]
//   - [NSCoder.NSCoderErrorMinimum]: The start of the range of error codes reserved for coder errors.
//   - [NSCoder.SetNSCoderErrorMinimum]
//   - [NSCoder.NSCoderReadCorruptError]: Decoding failed due to corrupt data.
//   - [NSCoder.SetNSCoderReadCorruptError]
//   - [NSCoder.NSCoderValueNotFoundError]: The requested data wasn’t found.
//   - [NSCoder.SetNSCoderValueNotFoundError]
//   - [NSCoder.NSCoderInvalidValueError]: Data wasn’t valid to encode.
//   - [NSCoder.SetNSCoderInvalidValueError]
//
// # Instance Methods
//
//   - [NSCoder.DecodeBytesForKeyMinimumLength]: Decode bytes from the decoder for a given key. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
//   - [NSCoder.DecodeBytesWithMinimumLength]: Decode bytes from the decoder. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder
type NSCoder struct {
	objectivec.Object
}

// NSCoderFromID constructs a [NSCoder] from an objc.ID.
//
// An abstract class that serves as the basis for objects that enable
// archiving and distribution of other objects.
func NSCoderFromID(id objc.ID) NSCoder {
	return NSCoder{objectivec.Object{id}}
}
// NOTE: NSCoder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCoder] class.
//
// # Inspecting a Coder
//
//   - [INSCoder.AllowsKeyedCoding]: A Boolean value that indicates whether the receiver supports keyed coding of objects.
//   - [INSCoder.ContainsValueForKey]: Returns a Boolean value that indicates whether an encoded value is available for a string.
//   - [INSCoder.DecodingFailurePolicy]: The action the coder should take when decoding fails.
//
// # Encoding General Data
//
//   - [INSCoder.EncodeArrayOfObjCTypeCountAt]: Encodes an array of the given Objective-C type, provided the number of items and a pointer.
//   - [INSCoder.EncodeBoolForKey]: Encodes a Boolean value and associates it with the string `key`.
//   - [INSCoder.EncodeBycopyObject]: An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded.
//   - [INSCoder.EncodeByrefObject]: An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded.
//   - [INSCoder.EncodeBytesLength]: Encodes a buffer of data of an unspecified type.
//   - [INSCoder.EncodeBytesLengthForKey]: Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
//   - [INSCoder.EncodeConditionalObject]: An encoding method for subclasses to override to conditionally encode an object, preserving common references to it.
//   - [INSCoder.EncodeConditionalObjectForKey]: An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
//   - [INSCoder.EncodeDataObject]: Encodes a given data object.
//   - [INSCoder.EncodeDoubleForKey]: Encodes a double-precision floating point value and associates it with the string key.
//   - [INSCoder.EncodeFloatForKey]: Encodes a floating point value and associates it with the string key.
//   - [INSCoder.EncodeIntForKey]: Encodes a C integer value and associates it with the string key.
//   - [INSCoder.EncodeIntegerForKey]: Encodes an integer value and associates it with the string key.
//   - [INSCoder.EncodeInt32ForKey]: Encodes a 32-bit integer value and associates it with the string key.
//   - [INSCoder.EncodeInt64ForKey]: Encodes a 64-bit integer value and associates it with the string key.
//   - [INSCoder.EncodeObject]: Encodes an object.
//   - [INSCoder.EncodeObjectForKey]: Encodes an object and associates it with the string key.
//   - [INSCoder.EncodePoint]: Encodes a point.
//   - [INSCoder.EncodePointForKey]: Encodes a point and associates it with the string key.
//   - [INSCoder.EncodePropertyList]: Encodes a property list.
//   - [INSCoder.EncodeRect]: Encodes a rectangle structure.
//   - [INSCoder.EncodeRectForKey]: Encodes a rectangle structure and associates it with the string key.
//   - [INSCoder.EncodeRootObject]: An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object.
//   - [INSCoder.EncodeSize]: Encodes a size structure.
//   - [INSCoder.EncodeSizeForKey]: Encodes a size structure and associates it with the given string key.
//   - [INSCoder.EncodeValueOfObjCTypeAt]: Encodes a value of the given type at the given address.
//
// # Encoding Core Media Time Structures
//
//   - [INSCoder.EncodeCMTimeForKey]: Encodes a given Core Media time structure and associates it with a specified key.
//   - [INSCoder.EncodeCMTimeRangeForKey]: Encodes a given Core Media time range structure and associates it with a specified key.
//   - [INSCoder.EncodeCMTimeMappingForKey]: Encodes a given Core Media time mapping structure and associates it with a specified key.
//
// # Secure Coding
//
//   - [INSCoder.RequiresSecureCoding]: Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//   - [INSCoder.AllowedClasses]: The set of coded classes allowed for secure coding.
//
// # Decoding General Data
//
//   - [INSCoder.DecodeArrayOfObjCTypeCountAt]: Decodes an array of `count` items, whose Objective-C type is given by `itemType`.
//   - [INSCoder.DecodeBoolForKey]: Decodes and returns a boolean value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu>) and associated with the string `key`.
//   - [INSCoder.DecodeBytesForKeyReturnedLength]: Decodes a buffer of data that was previously encoded with [encodeBytes(_:length:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)>) and associated with the string `key`.
//   - [INSCoder.DecodeBytesWithReturnedLength]: Decodes a buffer of data whose types are unspecified.
//   - [INSCoder.DecodeDataObject]: Decodes and returns an [NSData] object that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-1qd1e>). Subclasses must override this method.
//   - [INSCoder.DecodeDoubleForKey]: Decodes and returns a double value that was previously encoded with either [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
//   - [INSCoder.DecodeFloatForKey]: Decodes and returns a float value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
//   - [INSCoder.DecodeIntForKey]: Decodes and returns an int value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [INSCoder.DecodeIntegerForKey]: Decodes and returns an NSInteger value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [INSCoder.DecodeInt32ForKey]: Decodes and returns a 32-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [INSCoder.DecodeInt64ForKey]: Decodes and returns a 64-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
//   - [INSCoder.DecodeObject]: Decodes and returns an object that was previously encoded with any of the `encode…Object` methods.
//   - [INSCoder.DecodeObjectForKey]: Decodes and returns a previously-encoded object that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu>) or [encodeConditionalObject(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)>) and associated with the string `key`.
//   - [INSCoder.DecodePoint]: Decodes and returns an NSPoint structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-75jv4>).
//   - [INSCoder.DecodePointForKey]: Decodes and returns an NSPoint structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif>).
//   - [INSCoder.DecodePropertyList]: Decodes a property list that was previously encoded with [encodePropertyList(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodePropertyList(_:)>).
//   - [INSCoder.DecodeRect]: Decodes and returns an NSRect structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-3c1wz>).
//   - [INSCoder.DecodeRectForKey]: Decodes and returns an NSRect structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx>).
//   - [INSCoder.DecodeSize]: Decodes and returns an NSSize structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-82i7c>).
//   - [INSCoder.DecodeSizeForKey]: Decodes and returns an NSSize structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu>).
//   - [INSCoder.DecodeValueOfObjCTypeAtSize]: Decodes a single value of a known type from the specified data buffer.
//   - [INSCoder.DecodePropertyListForKey]: Returns a decoded property list for the specified key.
//
// # Decoding Core Media Time Structures
//
//   - [INSCoder.DecodeCMTimeForKey]: Returns the Core Media time structure associated with a given key.
//   - [INSCoder.DecodeCMTimeRangeForKey]: Returns the Core Media time range structure associated with a given key.
//   - [INSCoder.DecodeCMTimeMappingForKey]: Returns the Core Media time mapping structure associated with a given key.
//
// # Managing Decode Errors
//
//   - [INSCoder.FailWithError]: Signals to this coder that the decode operation has failed.
//   - [INSCoder.Error]: An error in the top-level encode.
//
// # Getting Version Information
//
//   - [INSCoder.SystemVersion]: The system version in effect for the archive.
//   - [INSCoder.VersionForClassName]: This method is present for historical reasons and is not used with keyed archivers.
//
// # Error Codes
//
//   - [INSCoder.NSCoderErrorMaximum]: The end of the range of error codes reserved for coder errors.
//   - [INSCoder.SetNSCoderErrorMaximum]
//   - [INSCoder.NSCoderErrorMinimum]: The start of the range of error codes reserved for coder errors.
//   - [INSCoder.SetNSCoderErrorMinimum]
//   - [INSCoder.NSCoderReadCorruptError]: Decoding failed due to corrupt data.
//   - [INSCoder.SetNSCoderReadCorruptError]
//   - [INSCoder.NSCoderValueNotFoundError]: The requested data wasn’t found.
//   - [INSCoder.SetNSCoderValueNotFoundError]
//   - [INSCoder.NSCoderInvalidValueError]: Data wasn’t valid to encode.
//   - [INSCoder.SetNSCoderInvalidValueError]
//
// # Instance Methods
//
//   - [INSCoder.DecodeBytesForKeyMinimumLength]: Decode bytes from the decoder for a given key. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
//   - [INSCoder.DecodeBytesWithMinimumLength]: Decode bytes from the decoder. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder
type INSCoder interface {
	objectivec.IObject

	// Topic: Inspecting a Coder

	// A Boolean value that indicates whether the receiver supports keyed coding of objects.
	AllowsKeyedCoding() bool
	// Returns a Boolean value that indicates whether an encoded value is available for a string.
	ContainsValueForKey(key string) bool
	// The action the coder should take when decoding fails.
	DecodingFailurePolicy() NSDecodingFailurePolicy

	// Topic: Encoding General Data

	// Encodes an array of the given Objective-C type, provided the number of items and a pointer.
	EncodeArrayOfObjCTypeCountAt(type_ []byte, count uint, array unsafe.Pointer)
	// Encodes a Boolean value and associates it with the string `key`.
	EncodeBoolForKey(value bool, key string)
	// An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded.
	EncodeBycopyObject(anObject objectivec.IObject)
	// An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded.
	EncodeByrefObject(anObject objectivec.IObject)
	// Encodes a buffer of data of an unspecified type.
	EncodeBytesLength(byteaddr unsafe.Pointer, length uint)
	// Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
	EncodeBytesLengthForKey(bytes uint8, length uint, key string)
	// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it.
	EncodeConditionalObject(object objectivec.IObject)
	// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
	EncodeConditionalObjectForKey(object objectivec.IObject, key string)
	// Encodes a given data object.
	EncodeDataObject(data INSData)
	// Encodes a double-precision floating point value and associates it with the string key.
	EncodeDoubleForKey(value float64, key string)
	// Encodes a floating point value and associates it with the string key.
	EncodeFloatForKey(value float32, key string)
	// Encodes a C integer value and associates it with the string key.
	EncodeIntForKey(value int, key string)
	// Encodes an integer value and associates it with the string key.
	EncodeIntegerForKey(value int, key string)
	// Encodes a 32-bit integer value and associates it with the string key.
	EncodeInt32ForKey(value int32, key string)
	// Encodes a 64-bit integer value and associates it with the string key.
	EncodeInt64ForKey(value int64, key string)
	// Encodes an object.
	EncodeObject(object objectivec.IObject)
	// Encodes an object and associates it with the string key.
	EncodeObjectForKey(object objectivec.IObject, key string)
	// Encodes a point.
	EncodePoint(point corefoundation.CGPoint)
	// Encodes a point and associates it with the string key.
	EncodePointForKey(point corefoundation.CGPoint, key string)
	// Encodes a property list.
	EncodePropertyList(aPropertyList objectivec.IObject)
	// Encodes a rectangle structure.
	EncodeRect(rect corefoundation.CGRect)
	// Encodes a rectangle structure and associates it with the string key.
	EncodeRectForKey(rect corefoundation.CGRect, key string)
	// An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object.
	EncodeRootObject(rootObject objectivec.IObject)
	// Encodes a size structure.
	EncodeSize(size corefoundation.CGSize)
	// Encodes a size structure and associates it with the given string key.
	EncodeSizeForKey(size corefoundation.CGSize, key string)
	// Encodes a value of the given type at the given address.
	EncodeValueOfObjCTypeAt(type_ []byte, addr unsafe.Pointer)

	// Topic: Encoding Core Media Time Structures

	// Encodes a given Core Media time structure and associates it with a specified key.
	EncodeCMTimeForKey(time objectivec.IObject, key string)
	// Encodes a given Core Media time range structure and associates it with a specified key.
	EncodeCMTimeRangeForKey(timeRange objectivec.IObject, key string)
	// Encodes a given Core Media time mapping structure and associates it with a specified key.
	EncodeCMTimeMappingForKey(timeMapping objectivec.IObject, key string)

	// Topic: Secure Coding

	// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
	RequiresSecureCoding() bool
	// The set of coded classes allowed for secure coding.
	AllowedClasses() INSSet

	// Topic: Decoding General Data

	// Decodes an array of `count` items, whose Objective-C type is given by `itemType`.
	DecodeArrayOfObjCTypeCountAt(itemType []byte, count uint, array unsafe.Pointer)
	// Decodes and returns a boolean value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu>) and associated with the string `key`.
	DecodeBoolForKey(key string) bool
	// Decodes a buffer of data that was previously encoded with [encodeBytes(_:length:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)>) and associated with the string `key`.
	DecodeBytesForKeyReturnedLength(key string, lengthp uint) uint8
	// Decodes a buffer of data whose types are unspecified.
	DecodeBytesWithReturnedLength(lengthp uint)
	// Decodes and returns an [NSData] object that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-1qd1e>). Subclasses must override this method.
	DecodeDataObject() INSData
	// Decodes and returns a double value that was previously encoded with either [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
	DecodeDoubleForKey(key string) float64
	// Decodes and returns a float value that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez>) or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu>) and associated with the string `key`.
	DecodeFloatForKey(key string) float32
	// Decodes and returns an int value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
	DecodeIntForKey(key string) int
	// Decodes and returns an NSInteger value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
	DecodeIntegerForKey(key string) int
	// Decodes and returns a 32-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
	DecodeInt32ForKey(key string) int32
	// Decodes and returns a 64-bit integer value that was previously encoded with [encodeCInt(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz>), [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z>), or [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg>) and associated with the string `key`.
	DecodeInt64ForKey(key string) int64
	// Decodes and returns an object that was previously encoded with any of the `encode…Object` methods.
	DecodeObject() objectivec.IObject
	// Decodes and returns a previously-encoded object that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu>) or [encodeConditionalObject(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)>) and associated with the string `key`.
	DecodeObjectForKey(key string) objectivec.IObject
	// Decodes and returns an NSPoint structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-75jv4>).
	DecodePoint() NSPoint
	// Decodes and returns an NSPoint structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif>).
	DecodePointForKey(key string) NSPoint
	// Decodes a property list that was previously encoded with [encodePropertyList(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encodePropertyList(_:)>).
	DecodePropertyList() objectivec.IObject
	// Decodes and returns an NSRect structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-3c1wz>).
	DecodeRect() NSRect
	// Decodes and returns an NSRect structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx>).
	DecodeRectForKey(key string) NSRect
	// Decodes and returns an NSSize structure that was previously encoded with [encode(_:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:)-82i7c>).
	DecodeSize() NSSize
	// Decodes and returns an NSSize structure that was previously encoded with [encode(_:forKey:)](<doc://com.apple.foundation/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu>).
	DecodeSizeForKey(key string) NSSize
	// Decodes a single value of a known type from the specified data buffer.
	DecodeValueOfObjCTypeAtSize(type_ []byte, data unsafe.Pointer, size uint)
	// Returns a decoded property list for the specified key.
	DecodePropertyListForKey(key string) objectivec.IObject

	// Topic: Decoding Core Media Time Structures

	// Returns the Core Media time structure associated with a given key.
	DecodeCMTimeForKey(key string) objectivec.IObject
	// Returns the Core Media time range structure associated with a given key.
	DecodeCMTimeRangeForKey(key string) objectivec.IObject
	// Returns the Core Media time mapping structure associated with a given key.
	DecodeCMTimeMappingForKey(key string) objectivec.IObject

	// Topic: Managing Decode Errors

	// Signals to this coder that the decode operation has failed.
	FailWithError(error_ INSError)
	// An error in the top-level encode.
	Error() INSError

	// Topic: Getting Version Information

	// The system version in effect for the archive.
	SystemVersion() uint32
	// This method is present for historical reasons and is not used with keyed archivers.
	VersionForClassName(className string) int

	// Topic: Error Codes

	// The end of the range of error codes reserved for coder errors.
	NSCoderErrorMaximum() int
	SetNSCoderErrorMaximum(value int)
	// The start of the range of error codes reserved for coder errors.
	NSCoderErrorMinimum() int
	SetNSCoderErrorMinimum(value int)
	// Decoding failed due to corrupt data.
	NSCoderReadCorruptError() int
	SetNSCoderReadCorruptError(value int)
	// The requested data wasn’t found.
	NSCoderValueNotFoundError() int
	SetNSCoderValueNotFoundError(value int)
	// Data wasn’t valid to encode.
	NSCoderInvalidValueError() int
	SetNSCoderInvalidValueError(value int)

	// Topic: Instance Methods

	// Decode bytes from the decoder for a given key. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
	DecodeBytesForKeyMinimumLength(key string, length uint) uint8
	// Decode bytes from the decoder. The length of the bytes must be greater than or equal to the `length` parameter. If the result exists, but is of insufficient length, then the decoder uses `failWithError` to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using [NSDecodingFailurePolicy].
	DecodeBytesWithMinimumLength(length uint)

	// Decodes the \c NSArray object for the given  \c key, which should be an \c NSArray, containing the given non-collection class (no nested arrays or arrays of dictionaries, etc) from the coder.
	DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key string) INSArray
	// Decodes the \c NSArray object for the given \c key, which should be an \c NSArray, containing the given non-collection classes (no nested arrays or arrays of dictionaries, etc) from the coder.
	DecodeArrayOfObjectsOfClassesForKey(classes INSSet, key string) INSArray
	// Decodes the \c NSDictionary object for the given \c key, which should be an \c NSDictionary<keyCls,objectCls> , with keys of type given in \c keyCls and objects of the given non-collection class \c objectCls (no nested dictionaries or other dictionaries contained in the dictionary, etc) from the coder.
	DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key string) INSDictionary
	// Decodes the \c NSDictionary object for the given \c key, which should be an \c NSDictionary, with keys of the types given in \c keyClasses and objects of the given non-collection classes in \c objectClasses (no nested dictionaries or other dictionaries contained in the dictionary, etc) from the given coder.
	DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses INSSet, objectClasses INSSet, key string) INSDictionary
	// Decodes an object for the key, restricted to the specified class.
	DecodeObjectOfClassForKey(aClass objc.Class, key string) objectivec.IObject
	// Decodes an object for the key, restricted to the specified classes.
	DecodeObjectOfClassesForKey(classes INSSet, key string) objectivec.IObject
	// Decodes the previously-encoded object associated by a key, populating an error if decoding fails.
	DecodeTopLevelObjectForKeyError(key string) (objectivec.IObject, error)
	// Decode an object as an expected type, failing if the archived type does not match.
	DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key string) (objectivec.IObject, error)
	// Decodes a series of potentially different Objective-C types.
	DecodeValuesOfObjCTypes(types []byte)
	// Encodes a series of values of potentially differing Objective-C types.
	EncodeValuesOfObjCTypes(types []byte)
	// This method is present for historical reasons and has no effect.
	ObjectZone() NSZone
}





// Init initializes the instance.
func (c NSCoder) Init() NSCoder {
	rv := objc.Send[NSCoder](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCoder) Autorelease() NSCoder {
	rv := objc.Send[NSCoder](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCoder creates a new NSCoder instance.
func NewNSCoder() NSCoder {
	class := getNSCoderClass()
	rv := objc.Send[NSCoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a Boolean value that indicates whether an encoded value is
// available for a string.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
// 
// The string is passed as `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/containsValue(forKey:)
func (c NSCoder) ContainsValueForKey(key string) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsValueForKey:"), objc.String(key))
	return rv
}

// Encodes an array of the given Objective-C type, provided the number of
// items and a pointer.
//
// # Discussion
// 
// The values are encoded from the buffer beginning at `address`. `itemType`
// must contain exactly one type code. [NSCoder]’s implementation invokes
// [EncodeValueOfObjCTypeAt] to encode the entire array of items. Subclasses
// that implement the [EncodeValueOfObjCTypeAt] method do not need to override
// this method.
// 
// This method must be matched by a subsequent [DecodeArrayOfObjCTypeCountAt]
// message.
// 
// For information on creating an Objective-C type code suitable for
// `itemType`, see [Type Encodings].
// 
// # Special Considerations
// 
// You should not use this method to encode C arrays of Objective-C objects.
// See [DecodeArrayOfObjCTypeCountAt] for more details.
//
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeArray(ofObjCType:count:at:)
func (c NSCoder) EncodeArrayOfObjCTypeCountAt(type_ []byte, count uint, array unsafe.Pointer) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeArrayOfObjCType:count:at:"), objc.CArray(type_), count, array)
}

// Encodes a Boolean value and associates it with the string `key`.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu
func (c NSCoder) EncodeBoolForKey(value bool, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBool:forKey:"), value, objc.String(key))
}

// An encoding method for subclasses to override such that it creates a copy,
// rather than a proxy, when decoded.
//
// # Discussion
// 
// [NSCoder]’s implementation simply invokes [EncodeObject].
// 
// This method must be matched by a corresponding [DecodeObject] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBycopyObject(_:)
func (c NSCoder) EncodeBycopyObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBycopyObject:"), anObject)
}

// An encoding method for subclasses to override such that it creates a proxy,
// rather than a copy, when decoded.
//
// # Discussion
// 
// [NSCoder]’s implementation simply invokes [EncodeObject].
// 
// This method must be matched by a corresponding [DecodeObject] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeByrefObject(_:)
func (c NSCoder) EncodeByrefObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeByrefObject:"), anObject)
}

// Encodes a buffer of data of an unspecified type.
//
// # Discussion
// 
// The buffer to be encoded begins at `address`, and its length in bytes is
// given by `numBytes`.
// 
// This method must be matched by a corresponding
// [DecodeBytesWithReturnedLength] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:)
func (c NSCoder) EncodeBytesLength(byteaddr unsafe.Pointer, length uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBytes:length:"), byteaddr, length)
}

// Encodes a buffer of data, given its length and a pointer, and associates it
// with a string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)
func (c NSCoder) EncodeBytesLengthForKey(bytes uint8, length uint, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBytes:length:forKey:"), bytes, length, objc.String(key))
}

// An encoding method for subclasses to override to conditionally encode an
// object, preserving common references to it.
//
// # Discussion
// 
// In the overriding method, `object` should be encoded only if it’s
// unconditionally encoded elsewhere (with any other `encode...Object:`
// method).
// 
// This method must be matched by a subsequent [DecodeObject] message. Upon
// decoding, if `object` was never encoded unconditionally, `decodeObject`
// returns `nil` in place of `object`. However, if `object` was encoded
// unconditionally, all references to `object` must be resolved.
// 
// [NSCoder]’s implementation simply invokes [EncodeObject].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:)
func (c NSCoder) EncodeConditionalObject(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeConditionalObject:"), object)
}

// An encoding method for subclasses to override to conditionally encode an
// object, preserving common references to it, only if it has been
// unconditionally encoded.
//
// # Discussion
// 
// Subclasses must override this method if they support keyed coding.
// 
// The encoded object is decoded with the [DecodeObjectForKey] method. If
// `objv` was never encoded unconditionally, [DecodeObjectForKey] returns
// `nil` in place of `objv`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)
func (c NSCoder) EncodeConditionalObjectForKey(object objectivec.IObject, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeConditionalObject:forKey:"), object, objc.String(key))
}

// Encodes a given data object.
//
// # Discussion
// 
// Subclasses must override this method.
// 
// This method must be matched by a subsequent [DecodeDataObject] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-1qd1e
func (c NSCoder) EncodeDataObject(data INSData) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeDataObject:"), data)
}

// Encodes a double-precision floating point value and associates it with the
// string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu
func (c NSCoder) EncodeDoubleForKey(value float64, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeDouble:forKey:"), value, objc.String(key))
}

// Encodes a floating point value and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez
func (c NSCoder) EncodeFloatForKey(value float32, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeFloat:forKey:"), value, objc.String(key))
}

// Encodes a C integer value and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)
func (c NSCoder) EncodeIntForKey(value int, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeInt:forKey:"), value, objc.String(key))
}

// Encodes an integer value and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz
func (c NSCoder) EncodeIntegerForKey(value int, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeInteger:forKey:"), value, objc.String(key))
}

// Encodes a 32-bit integer value and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z
func (c NSCoder) EncodeInt32ForKey(value int32, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeInt32:forKey:"), value, objc.String(key))
}

// Encodes a 64-bit integer value and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg
func (c NSCoder) EncodeInt64ForKey(value int64, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeInt64:forKey:"), value, objc.String(key))
}

// Encodes an object.
//
// # Discussion
// 
// [NSCoder]’s implementation simply invokes [EncodeValueOfObjCTypeAt] to
// encode `object`. Subclasses can override this method to encode a reference
// to `object` instead of `object` itself. For example, [NSArchiver] detects
// duplicate objects and encodes a reference to the original object rather
// than encode the same object twice.
// 
// This method must be matched by a subsequent [DecodeObject] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-9648d
func (c NSCoder) EncodeObject(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeObject:"), object)
}

// Encodes an object and associates it with the string key.
//
// # Discussion
// 
// Subclasses must override this method to identify multiple encodings of
// `objv` and encode a reference to `objv` instead. For example,
// [NSKeyedArchiver] detects duplicate objects and encodes a reference to the
// original object rather than encode the same object twice.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu
func (c NSCoder) EncodeObjectForKey(object objectivec.IObject, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeObject:forKey:"), object, objc.String(key))
}

// Encodes a point.
//
// # Discussion
// 
// [NSCoder]’s implementation invokes [EncodeValueOfObjCTypeAt] to encode
// `point`.
// 
// This method must be matched by a subsequent [DecodePoint] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-75jv4
func (c NSCoder) EncodePoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodePoint:"), point)
}

// Encodes a point and associates it with the string key.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif
func (c NSCoder) EncodePointForKey(point corefoundation.CGPoint, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodePoint:forKey:"), point, objc.String(key))
}

// Encodes a property list.
//
// # Discussion
// 
// [NSCoder]’s implementation invokes [EncodeValueOfObjCTypeAt] to encode
// `aPropertyList`.
// 
// This method must be matched by a subsequent [DecodePropertyList] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodePropertyList(_:)
func (c NSCoder) EncodePropertyList(aPropertyList objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodePropertyList:"), aPropertyList)
}

// Encodes a rectangle structure.
//
// # Discussion
// 
// [NSCoder]’s implementation invokes [EncodeValueOfObjCTypeAt] to encode
// `rect`.
// 
// This method must be matched by a subsequent [DecodeRect] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-3c1wz
func (c NSCoder) EncodeRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeRect:"), rect)
}

// Encodes a rectangle structure and associates it with the string key.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx
func (c NSCoder) EncodeRectForKey(rect corefoundation.CGRect, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeRect:forKey:"), rect, objc.String(key))
}

// An encoding method for subclasses to override to encode an interconnected
// group of objects, starting with the provided root object.
//
// # Discussion
// 
// [NSCoder]’s implementation simply invokes [EncodeObject].
// 
// This method must be matched by a subsequent [DecodeObject] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeRootObject(_:)
func (c NSCoder) EncodeRootObject(rootObject objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeRootObject:"), rootObject)
}

// Encodes a size structure.
//
// # Discussion
// 
// [NSCoder]’s implementation invokes [EncodeValueOfObjCTypeAt] to encode
// `size`.
// 
// This method must be matched by a subsequent [DecodeSize] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-82i7c
func (c NSCoder) EncodeSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeSize:"), size)
}

// Encodes a size structure and associates it with the given string key.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu
func (c NSCoder) EncodeSizeForKey(size corefoundation.CGSize, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeSize:forKey:"), size, objc.String(key))
}

// Encodes a value of the given type at the given address.
//
// # Discussion
// 
// Subclasses must override this method, and match it with a subsequent
// [DecodeValueOfObjCTypeAt] message.
// 
// When calling this method, `valueType` must contain exactly one type code.
// 
// For information on creating an Objective-C type code suitable for
// `valueType`, see [Type Encodings].
// 
// # Special Considerations
// 
// You should not use this method to encode Objective-C objects. See
// [DecodeArrayOfObjCTypeCountAt] for more details.
//
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValue(ofObjCType:at:)
func (c NSCoder) EncodeValueOfObjCTypeAt(type_ []byte, addr unsafe.Pointer) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeValueOfObjCType:at:"), unsafe.Pointer(unsafe.SliceData(type_)), addr)
}

// Encodes a given Core Media time structure and associates it with a
// specified key.
//
// time: A [CMTime] structure.
//
// key: The key with which to associate `time` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wbby
func (c NSCoder) EncodeCMTimeForKey(time objectivec.IObject, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeCMTime:forKey:"), time, objc.String(key))
}

// Encodes a given Core Media time range structure and associates it with a
// specified key.
//
// timeRange: A [CMTimeRange] structure.
//
// key: The key with which to associate `timeRange` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-46lo8
func (c NSCoder) EncodeCMTimeRangeForKey(timeRange objectivec.IObject, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeCMTimeRange:forKey:"), timeRange, objc.String(key))
}

// Encodes a given Core Media time mapping structure and associates it with a
// specified key.
//
// timeMapping: A [CMTimeMapping] structure.
//
// key: The key with which to associate `timeMapping` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-8tefb
func (c NSCoder) EncodeCMTimeMappingForKey(timeMapping objectivec.IObject, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeCMTimeMapping:forKey:"), timeMapping, objc.String(key))
}

// Decodes an array of `count` items, whose Objective-C type is given by
// `itemType`.
//
// # Discussion
// 
// The items are decoded into the buffer beginning at `address`, which must be
// large enough to contain them all. `itemType` must contain exactly one type
// code. [NSCoder]’s implementation invokes [DecodeValueOfObjCTypeAt] to
// decode the entire array of items.
// 
// This method matches an [EncodeArrayOfObjCTypeCountAt] message used during
// encoding.
// 
// For information on creating an Objective-C type code suitable for
// `itemType`, see [Type Encodings].
// 
// # Special Considerations
// 
// You should not use this method to decode C arrays of Objective-C objects.
// For historical reasons, returned objects will have an additional ownership
// reference which you can only relinquish using [CFRelease].
//
// [CFRelease]: https://developer.apple.com/documentation/CoreFoundation/CFRelease
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArray(ofObjCType:count:at:)
func (c NSCoder) DecodeArrayOfObjCTypeCountAt(itemType []byte, count uint, array unsafe.Pointer) {
	objc.Send[objc.ID](c.ID, objc.Sel("decodeArrayOfObjCType:count:at:"), objc.CArray(itemType), count, array)
}

// Decodes and returns a boolean value that was previously encoded with
// [EncodeBoolForKey] and associated with the string `key`.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBool(forKey:)
func (c NSCoder) DecodeBoolForKey(key string) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("decodeBoolForKey:"), objc.String(key))
	return rv
}

// Decodes a buffer of data that was previously encoded with
// [EncodeBytesLengthForKey] and associated with the string `key`.
//
// # Discussion
// 
// The buffer’s length is returned by reference in `lengthp`. The returned
// bytes are immutable. Subclasses must override this method if they perform
// keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:returnedLength:)
func (c NSCoder) DecodeBytesForKeyReturnedLength(key string, lengthp uint) uint8 {
	rv := objc.Send[uint8](c.ID, objc.Sel("decodeBytesForKey:returnedLength:"), objc.String(key), lengthp)
	return rv
}

// Decodes a buffer of data whose types are unspecified.
//
// # Discussion
// 
// [NSCoder]‘s implementation invokes [DecodeValueOfObjCTypeAt] to decode
// the data as a series of bytes, which this method then places into a buffer
// and returns. The buffer’s length is returned by reference in `numBytes`.
// If you need the bytes beyond the scope of the current `@autoreleasepool`
// block, you must copy them.
// 
// This method matches an [EncodeBytesLength] message used during encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withReturnedLength:)
func (c NSCoder) DecodeBytesWithReturnedLength(lengthp uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("decodeBytesWithReturnedLength:"), lengthp)
}

// Decodes and returns an [NSData] object that was previously encoded with
// [EncodeDataObject]. Subclasses must override this method.
//
// # Discussion
// 
// The implementation of your overriding method must match the implementation
// of your [EncodeDataObject] method. For example, a typical
// [EncodeDataObject] method encodes the number of bytes of data followed by
// the bytes themselves. Your override of this method must read the number of
// bytes, create an [NSData] object of the appropriate size, and decode the
// bytes into the new [NSData] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeData()
func (c NSCoder) DecodeDataObject() INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeDataObject"))
	return NSDataFromID(rv)
}

// Decodes and returns a double value that was previously encoded with either
// [EncodeFloatForKey] or [EncodeDoubleForKey] and associated with the string
// `key`.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDouble(forKey:)
func (c NSCoder) DecodeDoubleForKey(key string) float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("decodeDoubleForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a float value that was previously encoded with
// [EncodeFloatForKey] or [EncodeDoubleForKey] and associated with the string
// `key`.
//
// # Discussion
// 
// If the value was encoded as a double, the extra precision is lost. If the
// encoded real value does not fit into a float, the method raises an
// [NSRangeException]. Subclasses must override this method if they perform
// keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeFloat(forKey:)
func (c NSCoder) DecodeFloatForKey(key string) float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("decodeFloatForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an int value that was previously encoded with
// [EncodeIntForKey], [EncodeIntegerForKey], [EncodeInt32ForKey], or
// [EncodeInt64ForKey] and associated with the string `key`.
//
// # Discussion
// 
// If the encoded integer does not fit into the default integer size, the
// method raises an [NSRangeException]. Subclasses must override this method
// if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCInt(forKey:)
func (c NSCoder) DecodeIntForKey(key string) int {
	rv := objc.Send[int](c.ID, objc.Sel("decodeIntForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an NSInteger value that was previously encoded with
// [EncodeIntForKey], [EncodeIntegerForKey], [EncodeInt32ForKey], or
// [EncodeInt64ForKey] and associated with the string `key`.
//
// # Discussion
// 
// If the encoded integer does not fit into the NSInteger size, the method
// raises an [NSRangeException]. Subclasses must override this method if they
// perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInteger(forKey:)
func (c NSCoder) DecodeIntegerForKey(key string) int {
	rv := objc.Send[int](c.ID, objc.Sel("decodeIntegerForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a 32-bit integer value that was previously encoded with
// [EncodeIntForKey], [EncodeIntegerForKey], [EncodeInt32ForKey], or
// [EncodeInt64ForKey] and associated with the string `key`.
//
// # Discussion
// 
// If the encoded integer does not fit into a 32-bit integer, the method
// raises an [NSRangeException]. Subclasses must override this method if they
// perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt32(forKey:)
func (c NSCoder) DecodeInt32ForKey(key string) int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("decodeInt32ForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a 64-bit integer value that was previously encoded with
// [EncodeIntForKey], [EncodeIntegerForKey], [EncodeInt32ForKey], or
// [EncodeInt64ForKey] and associated with the string `key`.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt64(forKey:)
func (c NSCoder) DecodeInt64ForKey(key string) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("decodeInt64ForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an object that was previously encoded with any of the
// `encode…Object` methods.
//
// # Discussion
// 
// [NSCoder]’s implementation invokes [DecodeValueOfObjCTypeAt] to decode
// the object data.
// 
// Subclasses may need to override this method if they override any of the
// corresponding `encode…Object` methods. For example, if an object was
// encoded conditionally using the [EncodeConditionalObject] method, this
// method needs to check whether the object had actually been encoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject()
func (c NSCoder) DecodeObject() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeObject"))
	return objectivec.Object{ID: rv}
}

// Decodes and returns a previously-encoded object that was previously encoded
// with [EncodeObjectForKey] or [EncodeConditionalObjectForKey] and associated
// with the string `key`.
//
// # Discussion
// 
// Subclasses must override this method if they perform keyed coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject(forKey:)
func (c NSCoder) DecodeObjectForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeObjectForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Decodes and returns an NSPoint structure that was previously encoded with
// [EncodePoint].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint()
func (c NSCoder) DecodePoint() NSPoint {
	rv := objc.Send[NSPoint](c.ID, objc.Sel("decodePoint"))
	return NSPoint(rv)
}

// Decodes and returns an NSPoint structure that was previously encoded with
// [EncodePointForKey].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint(forKey:)
func (c NSCoder) DecodePointForKey(key string) NSPoint {
	rv := objc.Send[NSPoint](c.ID, objc.Sel("decodePointForKey:"), objc.String(key))
	return NSPoint(rv)
}

// Decodes a property list that was previously encoded with
// [EncodePropertyList].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList()
func (c NSCoder) DecodePropertyList() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodePropertyList"))
	return objectivec.Object{ID: rv}
}

// Decodes and returns an NSRect structure that was previously encoded with
// [EncodeRect].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect()
func (c NSCoder) DecodeRect() NSRect {
	rv := objc.Send[NSRect](c.ID, objc.Sel("decodeRect"))
	return NSRect(rv)
}

// Decodes and returns an NSRect structure that was previously encoded with
// [EncodeRectForKey].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect(forKey:)
func (c NSCoder) DecodeRectForKey(key string) NSRect {
	rv := objc.Send[NSRect](c.ID, objc.Sel("decodeRectForKey:"), objc.String(key))
	return NSRect(rv)
}

// Decodes and returns an NSSize structure that was previously encoded with
// [EncodeSize].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize()
func (c NSCoder) DecodeSize() NSSize {
	rv := objc.Send[NSSize](c.ID, objc.Sel("decodeSize"))
	return NSSize(rv)
}

// Decodes and returns an NSSize structure that was previously encoded with
// [EncodeSizeForKey].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize(forKey:)
func (c NSCoder) DecodeSizeForKey(key string) NSSize {
	rv := objc.Send[NSSize](c.ID, objc.Sel("decodeSizeForKey:"), objc.String(key))
	return NSSize(rv)
}

// Decodes a single value of a known type from the specified data buffer.
//
// type: The Objective-C type to decode.
//
// data: The data buffer to decode from.
//
// size: The size of the data buffer.
//
// # Discussion
// 
// The `type` parameter must contain exactly one type code, and the buffer
// specified by `data` must be large enough to hold the value corresponding to
// that type code. For information on creating an Objective-C type code
// suitable for `type`, see [NSMethodSignature].
// 
// Subclasses must override this method and provide an implementation to
// decode the value. In your overriding implementation, decode the value into
// the buffer beginning at data.
// 
// This method matches an [EncodeValueOfObjCTypeAt] message used during
// encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValue(ofObjCType:at:size:)
func (c NSCoder) DecodeValueOfObjCTypeAtSize(type_ []byte, data unsafe.Pointer, size uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("decodeValueOfObjCType:at:size:"), unsafe.Pointer(unsafe.SliceData(type_)), data, size)
}

// Returns a decoded property list for the specified key.
//
// key: The coder key.
//
// # Return Value
// 
// A decoded object containing a property list.
//
// # Discussion
// 
// This method calls [DecodeObjectOfClassesForKey] with a set allowing only
// property list types.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList(forKey:)
func (c NSCoder) DecodePropertyListForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodePropertyListForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns the Core Media time structure associated with a given key.
//
// key: The key for a [CMTime] structure encoded in the receiver.
//
// # Return Value
// 
// The [CMTime] structure associated with `key` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTime(forKey:)
func (c NSCoder) DecodeCMTimeForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeCMTimeForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns the Core Media time range structure associated with a given key.
//
// key: The key for a [CMTimeRange] structure encoded in the receiver.
//
// # Return Value
// 
// The [CMTimeRange] structure associated with `key` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeRange(forKey:)
func (c NSCoder) DecodeCMTimeRangeForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeCMTimeRangeForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns the Core Media time mapping structure associated with a given key.
//
// key: The key for a [CMTimeMapping] structure encoded in the receiver.
//
// # Return Value
// 
// The [CMTimeMapping] structure associated with `key` in the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeMapping(forKey:)
func (c NSCoder) DecodeCMTimeMappingForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeCMTimeMappingForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Signals to this coder that the decode operation has failed.
//
// error: An error that indicates why decoding failed.
//
// # Discussion
// 
// Typically, you call this method in your [InitWithCoder] implementation. You
// should set the error when you detect problems such as lack of secure
// coding, data corruption, or a domain validation failure.
// 
// This method is only meaningful to call for decodes.
// 
// The effect of calling this method depends on the value of
// [DecodingFailurePolicy], as follows:
// 
// - If the policy is [DecodingFailurePolicyRaiseException], calling this
// method throws an exception immediately. Swift code cannot catch this kind
// of exception. - If the policy is [DecodingFailurePolicySetErrorAndReturn],
// calling this method sets the error property once per call to one of the
// `decode` methods. Calling it repeatedly has no effect until the call stack
// unwinds to one of these methods’ entry points.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/failWithError(_:)
func (c NSCoder) FailWithError(error_ INSError) {
	objc.Send[objc.ID](c.ID, objc.Sel("failWithError:"), error_)
}

// This method is present for historical reasons and is not used with keyed
// archivers.
//
// # Return Value
// 
// The version in effect for the class named `className` or [NSNotFound] if no
// class named `className` exists.
//
// # Discussion
// 
// The version number does apply not to [NSKeyedArchiver]/[NSKeyedUnarchiver].
// A keyed archiver does not encode class version numbers.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/version(forClassName:)
func (c NSCoder) VersionForClassName(className string) int {
	rv := objc.Send[int](c.ID, objc.Sel("versionForClassName:"), objc.String(className))
	return rv
}

// Decode bytes from the decoder for a given key. The length of the bytes must
// be greater than or equal to the `length` parameter. If the result exists,
// but is of insufficient length, then the decoder uses `failWithError` to
// fail the entire decode operation. The result of that is configurable on a
// per-NSCoder basis using [NSDecodingFailurePolicy].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:minimumLength:)
func (c NSCoder) DecodeBytesForKeyMinimumLength(key string, length uint) uint8 {
	rv := objc.Send[uint8](c.ID, objc.Sel("decodeBytesForKey:minimumLength:"), objc.String(key), length)
	return rv
}

// Decode bytes from the decoder. The length of the bytes must be greater than
// or equal to the `length` parameter. If the result exists, but is of
// insufficient length, then the decoder uses `failWithError` to fail the
// entire decode operation. The result of that is configurable on a
// per-NSCoder basis using [NSDecodingFailurePolicy].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withMinimumLength:)
func (c NSCoder) DecodeBytesWithMinimumLength(length uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("decodeBytesWithMinimumLength:"), length)
}

// Decodes the \c NSArray object for the given \c key, which should be an \c
// NSArray, containing the given non-collection class (no nested arrays or
// arrays of dictionaries, etc) from the coder.
//
// # Discussion
// 
// Requires \c NSSecureCoding otherwise an exception is thrown and sets the \c
// decodingFailurePolicy to \c NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the object for \c key is not of the expected types, or
// cannot be decoded, and sets the \c error on the decoder.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClass:forKey:
func (c NSCoder) DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key string) INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeArrayOfObjectsOfClass:forKey:"), cls, objc.String(key))
	return NSArrayFromID(rv)
}

// Decodes the \c NSArray object for the given \c key, which should be an \c
// NSArray, containing the given non-collection classes (no nested arrays or
// arrays of dictionaries, etc) from the coder.
//
// # Discussion
// 
// Requires \c NSSecureCoding otherwise an exception is thrown and sets the \c
// decodingFailurePolicy to \c NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the object for \c key is not of the expected types, or
// cannot be decoded, and sets the \c error on the decoder.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClasses:forKey:
func (c NSCoder) DecodeArrayOfObjectsOfClassesForKey(classes INSSet, key string) INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeArrayOfObjectsOfClasses:forKey:"), classes, objc.String(key))
	return NSArrayFromID(rv)
}

// Decodes the \c NSDictionary object for the given \c key, which should be an
// \c NSDictionary , with keys of type given in \c keyCls and objects of the
// given non-collection class \c objectCls (no nested dictionaries or other
// dictionaries contained in the dictionary, etc) from the coder.
//
// # Discussion
// 
// Requires \c NSSecureCoding otherwise an exception is thrown and sets the \c
// decodingFailurePolicy to \c NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the object for \c key is not of the expected types, or
// cannot be decoded, and sets the \c error on the decoder.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:
func (c NSCoder) DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key string) INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:"), keyCls, objectCls, objc.String(key))
	return NSDictionaryFromID(rv)
}

// Decodes the \c NSDictionary object for the given \c key, which should be an
// \c NSDictionary, with keys of the types given in \c keyClasses and objects
// of the given non-collection classes in \c objectClasses (no nested
// dictionaries or other dictionaries contained in the dictionary, etc) from
// the given coder.
//
// # Discussion
// 
// Requires \c NSSecureCoding otherwise an exception is thrown and sets the \c
// decodingFailurePolicy to \c NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the object for \c key is not of the expected types, or
// cannot be decoded, and sets the \c error on the decoder.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:
func (c NSCoder) DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses INSSet, objectClasses INSSet, key string) INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:"), keyClasses, objectClasses, objc.String(key))
	return NSDictionaryFromID(rv)
}

// Decodes an object for the key, restricted to the specified class.
//
// aClass: The expect class type.
//
// key: The coder key.
//
// # Return Value
// 
// The decoded object.
//
// # Discussion
// 
// If the coder responds [true] to [RequiresSecureCoding], then an exception
// will be thrown if the class to be decoded does not implement
// [NSSecureCoding] or is not [isKind(of:)] of `aClass`.
// 
// If the coder responds [false] to [RequiresSecureCoding], then the class
// argument is ignored and no check of the class of the decoded object is
// performed, exactly as if [DecodeObjectForKey] had been called.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isKind(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c NSCoder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, objc.String(key))
	return objectivec.Object{ID: rv}
}

// Decodes an object for the key, restricted to the specified classes.
//
// classes: A set of the expected classes.
//
// key: The coder key.
//
// # Return Value
// 
// The decoded object.
//
// # Discussion
// 
// The class of the object may be any class in the `classes` set, or a
// subclass of any class in the set. Otherwise, the behavior is the same as
// [DecodeObjectOfClassForKey].
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClasses:forKey:
func (c NSCoder) DecodeObjectOfClassesForKey(classes INSSet, key string) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeObjectOfClasses:forKey:"), classes, objc.String(key))
	return objectivec.Object{ID: rv}
}

// Decodes the previously-encoded object associated by a key, populating an
// error if decoding fails.
//
// key: The key that identifies the object to decode.
//
// error: An [NSError] reference. On return, if this value is not `nil`, it
// represents an error encountered while decoding.
//
// # Return Value
// 
// The decoded object, or `nil` if decoding fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectForKey:error:
func (c NSCoder) DecodeTopLevelObjectForKeyError(key string) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeTopLevelObjectForKey:error:"), objc.String(key), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Decode an object as an expected type, failing if the archived type does not
// match.
//
// aClass: The expected class of the object being decoded.
//
// key: The archive key indicating the member to decode.
//
// error: On return, an [NSError] indicating why decoding failed, or `nil` if no
// error occurred.
//
// # Return Value
// 
// The decoded object, or `nil` if decoding fails.
//
// # Discussion
// 
// If the coder responds [true] to [RequiresSecureCoding], then the coder
// calls [FailWithError] in either the following cases:
// 
// - The class indicated by `cls` does not implement [NSSecureCoding]. - The
// unarchived class does not match `cls`, nor do any of its superclasses.
// 
// If the coder does not require secure coding, it ignores the `cls` parameter
// and does not check the decoded object.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectOfClass:forKey:error:
func (c NSCoder) DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key string) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeTopLevelObjectOfClass:forKey:error:"), aClass, objc.String(key), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Decodes a series of potentially different Objective-C types.
//
// # Discussion
// 
// `valueTypes` is a single C string containing any number of type codes. The
// variable arguments to this method consist of one or more pointer arguments,
// each of which specifies the buffer in which to place a single decoded
// value. For each type code in `valueTypes`, you must specify a corresponding
// pointer argument whose buffer is large enough to hold the decoded value.
// 
// This method matches an [EncodeValuesOfObjCTypes] message used during
// encoding.
// 
// [NSCoder]’s implementation invokes [DecodeValueOfObjCTypeAt] to decode
// individual types. Subclasses that implement the [DecodeValueOfObjCTypeAt]
// method do not need to override this method.
// 
// For information on creating Objective-C type codes suitable for
// `valueTypes`, see [Type Encodings].
// 
// # Special Considerations
// 
// You should not use this method to decode Objective-C objects. See
// [DecodeArrayOfObjCTypeCountAt] for more details.
//
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValuesOfObjCTypes:
func (c NSCoder) DecodeValuesOfObjCTypes(types []byte) {
	objc.Send[objc.ID](c.ID, objc.Sel("decodeValuesOfObjCTypes:"), unsafe.Pointer(unsafe.SliceData(types)))
}

// Encodes a series of values of potentially differing Objective-C types.
//
// # Discussion
// 
// `valueTypes` is a C string containing any number of type codes. The
// variable arguments to this method consist of one or more pointer arguments,
// each of which specifies a buffer containing the value to be encoded. For
// each type code in `valueTypes`, you must specify a corresponding pointer
// argument.
// 
// This method must be matched by a subsequent [DecodeValuesOfObjCTypes]
// message.
// 
// [NSCoder]’s implementation invokes [EncodeValueOfObjCTypeAt] to encode
// individual types. Subclasses that implement the [EncodeValueOfObjCTypeAt]
// method do not need to override this method. However, subclasses that
// provide a more efficient approach for encoding a series of values may
// override this method to implement that approach.
// 
// For information on creating Objective-C type codes suitable for
// `valueTypes`, see [Type Encodings].
// 
// # Special Considerations
// 
// You should not use this method to encode Objective-C objects. See
// [DecodeArrayOfObjCTypeCountAt] for more details.
//
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValuesOfObjCTypes:
func (c NSCoder) EncodeValuesOfObjCTypes(types []byte) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeValuesOfObjCTypes:"), unsafe.Pointer(unsafe.SliceData(types)))
}

// This method is present for historical reasons and has no effect.
//
// # Discussion
// 
// [NSCoder]’s implementation returns the default memory zone, as given by
// `NSDefaultMallocZone()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/objectZone
func (c NSCoder) ObjectZone() NSZone {
	rv := objc.Send[NSZone](c.ID, objc.Sel("objectZone"))
	return NSZone(rv)
}












// A Boolean value that indicates whether the receiver supports keyed coding
// of objects.
//
// # Discussion
// 
// [false] by default. Concrete subclasses that support keyed coding, such as
// [NSKeyedArchiver], need to override this property to return [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/allowsKeyedCoding
func (c NSCoder) AllowsKeyedCoding() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsKeyedCoding"))
	return rv
}



// The action the coder should take when decoding fails.
//
// # Discussion
// 
// A decode call can fail for the following reasons:
// 
// - The keyed archive data is corrupt or missing. - A type mismatch occurs,
// such as expecting a class by calling [decodeObject(of:forKey:)] but
// encountering a numeric type instead. This also occurs when
// [DecodeIntegerForKey] encounters a value encoded as floating-point, or vice
// versa. - A secure coding violation occurs. This happens when you attempt to
// decode an object that doesn’t conform to [NSSecureCoding]. This also
// happens when the encoded type doesn’t match any of the types passed to
// [decodeObject(of:forKey:)].
//
// [decodeObject(of:forKey:)]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject(of:forKey:)-roif
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c NSCoder) DecodingFailurePolicy() NSDecodingFailurePolicy {
	rv := objc.Send[NSDecodingFailurePolicy](c.ID, objc.Sel("decodingFailurePolicy"))
	return NSDecodingFailurePolicy(rv)
}



// Indicates whether the archiver requires all archived classes to resist
// object substitution attacks.
//
// # Discussion
// 
// [true] if this coder requires secure coding; [false] otherwise.
// 
// Secure coders check a set of allowed classes before decoding objects, and
// all objects must implement the [NSSecureCoding] protocol.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/requiresSecureCoding
func (c NSCoder) RequiresSecureCoding() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requiresSecureCoding"))
	return rv
}



// The set of coded classes allowed for secure coding.
//
// # Discussion
// 
// Secure coders check this set of allowed classes before decoding objects,
// and all objects must implement the [NSSecureCoding] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/allowedClasses
func (c NSCoder) AllowedClasses() INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("allowedClasses"))
	return NSSetFromID(objc.ID(rv))
}



// An error in the top-level encode.
//
// # Discussion
// 
// The meaning of this property depends on the setting of the
// [DecodingFailurePolicy] property. For
// [DecodingFailurePolicyRaiseException], this property is always `nil`. For
// [DecodingFailurePolicySetErrorAndReturn], a non-`nil` value represents the
// first error encountered while decoding the archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/error
func (c NSCoder) Error() INSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("error"))
	return NSErrorFromID(objc.ID(rv))
}



// The system version in effect for the archive.
//
// # Discussion
// 
// During encoding, the current version. During decoding, the version that was
// in effect when the data was encoded.
// 
// Subclasses that implement decoding must override this property to return
// the system version of the data being decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/systemVersion
func (c NSCoder) SystemVersion() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("systemVersion"))
	return rv
}



// The end of the range of error codes reserved for coder errors.
//
// See: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c NSCoder) NSCoderErrorMaximum() int {
	rv := objc.Send[int](c.ID, objc.Sel("NSCoderErrorMaximum"))
	return rv
}
func (c NSCoder) SetNSCoderErrorMaximum(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNSCoderErrorMaximum:"), value)
}



// The start of the range of error codes reserved for coder errors.
//
// See: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c NSCoder) NSCoderErrorMinimum() int {
	rv := objc.Send[int](c.ID, objc.Sel("NSCoderErrorMinimum"))
	return rv
}
func (c NSCoder) SetNSCoderErrorMinimum(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNSCoderErrorMinimum:"), value)
}



// Decoding failed due to corrupt data.
//
// See: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c NSCoder) NSCoderReadCorruptError() int {
	rv := objc.Send[int](c.ID, objc.Sel("NSCoderReadCorruptError"))
	return rv
}
func (c NSCoder) SetNSCoderReadCorruptError(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNSCoderReadCorruptError:"), value)
}



// The requested data wasn’t found.
//
// See: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c NSCoder) NSCoderValueNotFoundError() int {
	rv := objc.Send[int](c.ID, objc.Sel("NSCoderValueNotFoundError"))
	return rv
}
func (c NSCoder) SetNSCoderValueNotFoundError(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNSCoderValueNotFoundError:"), value)
}



// Data wasn’t valid to encode.
//
// See: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c NSCoder) NSCoderInvalidValueError() int {
	rv := objc.Send[int](c.ID, objc.Sel("NSCoderInvalidValueError"))
	return rv
}
func (c NSCoder) SetNSCoderInvalidValueError(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNSCoderInvalidValueError:"), value)
}






















