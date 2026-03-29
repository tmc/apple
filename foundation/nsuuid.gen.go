// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUUID] class.
var (
	_NSUUIDClass     NSUUIDClass
	_NSUUIDClassOnce sync.Once
)

func getNSUUIDClass() NSUUIDClass {
	_NSUUIDClassOnce.Do(func() {
		_NSUUIDClass = NSUUIDClass{class: objc.GetClass("NSUUID")}
	})
	return _NSUUIDClass
}

// GetNSUUIDClass returns the class object for NSUUID.
func GetNSUUIDClass() NSUUIDClass {
	return getNSUUIDClass()
}

type NSUUIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSUUIDClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUUIDClass) Alloc() NSUUID {
	rv := objc.Send[NSUUID](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A universally unique value that can be used to identify types, interfaces,
// and other items.
//
// # Overview
// 
// In Swift, this object bridges to [UUID]; use [NSUUID] when you need
// reference semantics or other Foundation-specific behavior.
// 
// UUIDs (Universally Unique Identifiers), also known as GUIDs (Globally
// Unique Identifiers) or IIDs (Interface Identifiers), are 128-bit values.
// UUIDs created by [NSUUID] conform to RFC 4122 version 4 and are created
// with random bytes.
// 
// The standard format for UUIDs represented in ASCII is a string punctuated
// by hyphens, for example `68753A44-4D6F-1226-9C60-0050E4C00067`. The hex
// representation looks, as you might expect, like a list of numerical values
// preceded by 0x. For example, `0xD7`, `0x36`, `0x95`, `0x0A`, `0x4D`,
// `0x6E`, `0x12`, `0x26`, `0x80`, `0x3A`, `0x00`, `0x50`, `0xE4`, `0xC0`,
// `0x00`, `0x67`. Because a UUID is expressed simply as an array of bytes,
// there are no endianness considerations for different platforms.
// 
// The [NSUUID] class is toll-free bridged with CoreFoundation’s [CFUUID].
// Use UUID strings to convert between [CFUUIDRef] and [NSUUID], if needed.
// Two [NSUUID] objects are not guaranteed to be comparable by pointer value
// (as [CFUUID] is); use [isEqual(_:)] to compare two [NSUUID] instances.
//
// [CFUUID]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Creating UUIDs
//
//   - [NSUUID.InitWithUUIDString]: Initializes a new UUID with the formatted string.
//   - [NSUUID.InitWithUUIDBytes]: Initializes a new UUID with the given bytes.
//
// # Get UUID Values
//
//   - [NSUUID.GetUUIDBytes]: Returns the UUID as bytes.
//   - [NSUUID.UUIDString]: The UUID as a string.
//
// # Instance Methods
//
//   - [NSUUID.Compare]
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID
type NSUUID struct {
	objectivec.Object
}

// NSUUIDFromID constructs a [NSUUID] from an objc.ID.
//
// A universally unique value that can be used to identify types, interfaces,
// and other items.
func NSUUIDFromID(id objc.ID) NSUUID {
	return NSUUID{objectivec.Object{ID: id}}
}
// NOTE: NSUUID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSUUID] class.
//
// # Creating UUIDs
//
//   - [INSUUID.InitWithUUIDString]: Initializes a new UUID with the formatted string.
//   - [INSUUID.InitWithUUIDBytes]: Initializes a new UUID with the given bytes.
//
// # Get UUID Values
//
//   - [INSUUID.GetUUIDBytes]: Returns the UUID as bytes.
//   - [INSUUID.UUIDString]: The UUID as a string.
//
// # Instance Methods
//
//   - [INSUUID.Compare]
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID
type INSUUID interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating UUIDs

	// Initializes a new UUID with the formatted string.
	InitWithUUIDString(string_ string) NSUUID
	// Initializes a new UUID with the given bytes.
	InitWithUUIDBytes(bytes unsafe.Pointer) NSUUID

	// Topic: Get UUID Values

	// Returns the UUID as bytes.
	GetUUIDBytes(uuid unsafe.Pointer)
	// The UUID as a string.
	UUIDString() string

	// Topic: Instance Methods

	Compare(otherUUID INSUUID) ComparisonResult
}

// Init initializes the instance.
func (u NSUUID) Init() NSUUID {
	rv := objc.Send[NSUUID](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUUID) Autorelease() NSUUID {
	rv := objc.Send[NSUUID](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUUID creates a new NSUUID instance.
func NewNSUUID() NSUUID {
	class := getNSUUIDClass()
	rv := objc.Send[NSUUID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUUIDWithCoder(coder INSCoder) NSUUID {
	instance := getNSUUIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSUUIDFromID(rv)
}

// Initializes a new UUID with the given bytes.
//
// bytes: Raw UUID bytes to use to create the UUID.
//
// # Return Value
// 
// A new UUID object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidBytes:)
func NewUUIDWithUUIDBytes(bytes unsafe.Pointer) NSUUID {
	instance := getNSUUIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUIDBytes:"), bytes)
	return NSUUIDFromID(rv)
}

// Initializes a new UUID with the formatted string.
//
// string: The source string containing the UUID. The standard format for UUIDs
// represented in ASCII is a string punctuated by hyphens, for example
// `68753A44-4D6F-1226-9C60-0050E4C00067`.
//
// # Return Value
// 
// A new UUID object. Returns `nil` for invalid strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidString:)
func NewUUIDWithUUIDString(string_ string) NSUUID {
	instance := getNSUUIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUIDString:"), objc.String(string_))
	return NSUUIDFromID(rv)
}

// Initializes a new UUID with the formatted string.
//
// string: The source string containing the UUID. The standard format for UUIDs
// represented in ASCII is a string punctuated by hyphens, for example
// `68753A44-4D6F-1226-9C60-0050E4C00067`.
//
// # Return Value
// 
// A new UUID object. Returns `nil` for invalid strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidString:)
func (u NSUUID) InitWithUUIDString(string_ string) NSUUID {
	rv := objc.Send[NSUUID](u.ID, objc.Sel("initWithUUIDString:"), objc.String(string_))
	return rv
}
// Initializes a new UUID with the given bytes.
//
// bytes: Raw UUID bytes to use to create the UUID.
//
// # Return Value
// 
// A new UUID object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/init(uuidBytes:)
func (u NSUUID) InitWithUUIDBytes(bytes unsafe.Pointer) NSUUID {
	rv := objc.Send[NSUUID](u.ID, objc.Sel("initWithUUIDBytes:"), bytes)
	return rv
}
// Returns the UUID as bytes.
//
// uuid: The value of uuid represented as raw bytes.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/getBytes(_:)
func (u NSUUID) GetUUIDBytes(uuid unsafe.Pointer) {
	objc.Send[objc.ID](u.ID, objc.Sel("getUUIDBytes:"), uuid)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/compare(_:)
func (u NSUUID) Compare(otherUUID INSUUID) ComparisonResult {
	rv := objc.Send[ComparisonResult](u.ID, objc.Sel("compare:"), otherUUID)
	return ComparisonResult(rv)
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u NSUUID) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u NSUUID) InitWithCoder(coder INSCoder) NSUUID {
	rv := objc.Send[NSUUID](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Create and returns a new UUID with RFC 4122 version 4 random bytes.
//
// # Return Value
// 
// A new UUID object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/UUID
func (_NSUUIDClass NSUUIDClass) UUID() NSUUID {
	rv := objc.Send[objc.ID](objc.ID(_NSUUIDClass.class), objc.Sel("UUID"))
	return NSUUIDFromID(rv)
}

// The UUID as a string.
//
// # Discussion
// 
// A string containing a formatted UUID for example
// `E621E1F8-C36C-495A-93FC-0C247A3E6E5F`.
// 
// Use this property when you need a string representation of the [NSUUID]
// object—for example, to compare with a [CFUUID] object.
//
// [CFUUID]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
//
// See: https://developer.apple.com/documentation/Foundation/NSUUID/uuidString
func (u NSUUID) UUIDString() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("UUIDString"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

