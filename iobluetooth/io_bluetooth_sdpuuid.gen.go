// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothSDPUUID] class.
var (
	_IOBluetoothSDPUUIDClass     IOBluetoothSDPUUIDClass
	_IOBluetoothSDPUUIDClassOnce sync.Once
)

func getIOBluetoothSDPUUIDClass() IOBluetoothSDPUUIDClass {
	_IOBluetoothSDPUUIDClassOnce.Do(func() {
		_IOBluetoothSDPUUIDClass = IOBluetoothSDPUUIDClass{class: objc.GetClass("IOBluetoothSDPUUID")}
	})
	return _IOBluetoothSDPUUIDClass
}

// GetIOBluetoothSDPUUIDClass returns the class object for IOBluetoothSDPUUID.
func GetIOBluetoothSDPUUIDClass() IOBluetoothSDPUUIDClass {
	return getIOBluetoothSDPUUIDClass()
}

type IOBluetoothSDPUUIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothSDPUUIDClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothSDPUUIDClass) Alloc() IOBluetoothSDPUUID {
	rv := objc.Send[IOBluetoothSDPUUID](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An NSData subclass that represents a UUID as defined in the Bluetooth SDP
// spec.
//
// # Overview
//
// The IOBluetoothSDPUUID class can represent a UUID of any valid size (16, 32
// or 128 bits). It provides the ability to compare two UUIDs no matter what
// their size as well as the ability to promote the size of a UUID to a larger
// one.
//
// # Initializers
//
//   - [IOBluetoothSDPUUID.InitWithUUID16]: Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
//   - [IOBluetoothSDPUUID.InitWithUUID32]: Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// # Instance Methods
//
//   - [IOBluetoothSDPUUID.ClassForArchiver]
//   - [IOBluetoothSDPUUID.ClassForCoder]
//   - [IOBluetoothSDPUUID.ClassForPortCoder]
//   - [IOBluetoothSDPUUID.GetUUIDWithLength]: Returns an IOBluetoothSDPUUID object matching the target UUID, but with the given number of bytes.
//   - [IOBluetoothSDPUUID.IsEqualToUUID]: Compares the target IOBluetoothSDPUUID object with the given otherUUID object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID
type IOBluetoothSDPUUID struct {
	foundation.NSData
}

// IOBluetoothSDPUUIDFromID constructs a [IOBluetoothSDPUUID] from an objc.ID.
//
// An NSData subclass that represents a UUID as defined in the Bluetooth SDP
// spec.
func IOBluetoothSDPUUIDFromID(id objc.ID) IOBluetoothSDPUUID {
	return IOBluetoothSDPUUID{NSData: foundation.NSDataFromID(id)}
}

// NOTE: IOBluetoothSDPUUID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothSDPUUID] class.
//
// # Initializers
//
//   - [IIOBluetoothSDPUUID.InitWithUUID16]: Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
//   - [IIOBluetoothSDPUUID.InitWithUUID32]: Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// # Instance Methods
//
//   - [IIOBluetoothSDPUUID.ClassForArchiver]
//   - [IIOBluetoothSDPUUID.ClassForCoder]
//   - [IIOBluetoothSDPUUID.ClassForPortCoder]
//   - [IIOBluetoothSDPUUID.GetUUIDWithLength]: Returns an IOBluetoothSDPUUID object matching the target UUID, but with the given number of bytes.
//   - [IIOBluetoothSDPUUID.IsEqualToUUID]: Compares the target IOBluetoothSDPUUID object with the given otherUUID object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID
type IIOBluetoothSDPUUID interface {
	foundation.INSData

	// Topic: Initializers

	// Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
	InitWithUUID16(uuid16 BluetoothSDPUUID16) IOBluetoothSDPUUID
	// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
	InitWithUUID32(uuid32 BluetoothSDPUUID32) IOBluetoothSDPUUID

	// Topic: Instance Methods

	ClassForArchiver() objectivec.Class
	ClassForCoder() objectivec.Class
	ClassForPortCoder() objectivec.Class
	// Returns an IOBluetoothSDPUUID object matching the target UUID, but with the given number of bytes.
	GetUUIDWithLength(newLength uint32) IIOBluetoothSDPUUID
	// Compares the target IOBluetoothSDPUUID object with the given otherUUID object.
	IsEqualToUUID(otherUUID IIOBluetoothSDPUUID) bool
}

// Init initializes the instance.
func (b IOBluetoothSDPUUID) Init() IOBluetoothSDPUUID {
	rv := objc.Send[IOBluetoothSDPUUID](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothSDPUUID) Autorelease() IOBluetoothSDPUUID {
	rv := objc.Send[IOBluetoothSDPUUID](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothSDPUUID creates a new IOBluetoothSDPUUID instance.
func NewIOBluetoothSDPUUID() IOBluetoothSDPUUID {
	class := getIOBluetoothSDPUUIDClass()
	rv := objc.Send[IOBluetoothSDPUUID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new IOBluetoothSDPUUID object with the given bytes of the given
// length.
//
// bytes: An array of bytes representing the UUID.
//
// length: The length of the array of bytes.
//
// # Return Value
//
// Returns the new IOBluetoothSDPUUID object or nil on failure.
//
// # Discussion
//
// If the length is invalid for a UUID, nil is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(bytes:length:)
func NewBluetoothSDPUUIDUuidWithBytesLength(bytes []byte) IOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](objc.ID(getIOBluetoothSDPUUIDClass().class), objc.Sel("uuidWithBytes:length:"), objc.BytesPointer(bytes), uint32(len(bytes)))
	return IOBluetoothSDPUUIDFromID(rv)
}

// Creates a new IOBluetoothSDPUUID object from the given NSData.
//
// data: The NSData containing the UUID bytes.
//
// # Return Value
//
// Returns the new IOBluetoothSDPUUID object or nil on failure.
//
// # Discussion
//
// If the length of the NSData is invalid for a UUID, nil is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(data:)
func NewBluetoothSDPUUIDUuidWithData(data foundation.NSData) IOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](objc.ID(getIOBluetoothSDPUUIDClass().class), objc.Sel("uuidWithData:"), data)
	return IOBluetoothSDPUUIDFromID(rv)
}

// Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
//
// uuid16: A scalar representing a 16-bit UUID
//
// # Return Value
//
// Returns self.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid16:)
func NewBluetoothSDPUUIDWithUUID16(uuid16 BluetoothSDPUUID16) IOBluetoothSDPUUID {
	instance := getIOBluetoothSDPUUIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID16:"), uuid16)
	return IOBluetoothSDPUUIDFromID(rv)
}

// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// uuid32: A scalar representing a 32-bit UUID
//
// # Return Value
//
// Returns self.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid32:)
func NewBluetoothSDPUUIDWithUUID32(uuid32 BluetoothSDPUUID32) IOBluetoothSDPUUID {
	instance := getIOBluetoothSDPUUIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID32:"), uuid32)
	return IOBluetoothSDPUUIDFromID(rv)
}

// Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
//
// uuid16: A scalar representing a 16-bit UUID
//
// # Return Value
//
// Returns self.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid16:)
func (b IOBluetoothSDPUUID) InitWithUUID16(uuid16 BluetoothSDPUUID16) IOBluetoothSDPUUID {
	rv := objc.Send[IOBluetoothSDPUUID](b.ID, objc.Sel("initWithUUID16:"), uuid16)
	return rv
}

// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// uuid32: A scalar representing a 32-bit UUID
//
// # Return Value
//
// Returns self.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid32:)
func (b IOBluetoothSDPUUID) InitWithUUID32(uuid32 BluetoothSDPUUID32) IOBluetoothSDPUUID {
	rv := objc.Send[IOBluetoothSDPUUID](b.ID, objc.Sel("initWithUUID32:"), uuid32)
	return rv
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForArchiver()
func (b IOBluetoothSDPUUID) ClassForArchiver() objectivec.Class {
	rv := objc.Send[objectivec.Class](b.ID, objc.Sel("classForArchiver"))
	return objectivec.Class(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForCoder()
func (b IOBluetoothSDPUUID) ClassForCoder() objectivec.Class {
	rv := objc.Send[objectivec.Class](b.ID, objc.Sel("classForCoder"))
	return objectivec.Class(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForPortCoder()
func (b IOBluetoothSDPUUID) ClassForPortCoder() objectivec.Class {
	rv := objc.Send[objectivec.Class](b.ID, objc.Sel("classForPortCoder"))
	return objectivec.Class(rv)
}

// Returns an IOBluetoothSDPUUID object matching the target UUID, but with the
// given number of bytes.
//
// newLength: The desired length for the UUID.
//
// # Return Value
//
// Returns an IOBluetoothSDPUUID object with the same data as the target but
// with the given length if it is possible to do so. Otherwise, nil is
// returned.
//
// # Discussion
//
// If the target object is the same length as newLength, it returns self. If
// newLength is greater it creates a new IOBluetoothSDPUUID object with the
// correct value for the given length. If newLength is smaller, it will
// attempt to create a new IOBluetoothSDPUUID that is smaller if the data
// matches the Bluetooth UUID base. This downconversion is currently
// unimplemented.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/getWithLength(_:)
func (b IOBluetoothSDPUUID) GetUUIDWithLength(newLength uint32) IIOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getUUIDWithLength:"), newLength)
	return IOBluetoothSDPUUIDFromID(rv)
}

// Compares the target IOBluetoothSDPUUID object with the given otherUUID
// object.
//
// otherUUID: The UUID object to be compared with the target.
//
// # Return Value
//
// Returns true if the UUID values of each object are equal. This includes the
// case where the sizes are different but the data itself is the same when the
// Bluetooth UUID base is applied.
//
// # Discussion
//
// This method will compare the two UUID values independent of their length.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/isEqual(to:)
func (b IOBluetoothSDPUUID) IsEqualToUUID(otherUUID IIOBluetoothSDPUUID) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isEqualToUUID:"), otherUUID)
	return rv
}

// Creates a new 16-bit IOBluetoothSDPUUID with the given UUID16
//
// uuid16: A scalar representing a 16-bit UUID
//
// # Return Value
//
// Returns the new IOBluetoothSDPUUID object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/uuid16(_:)
func (_IOBluetoothSDPUUIDClass IOBluetoothSDPUUIDClass) Uuid16(uuid16 BluetoothSDPUUID16) IOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPUUIDClass.class), objc.Sel("uuid16:"), uuid16)
	return IOBluetoothSDPUUIDFromID(rv)
}

// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// uuid32: A scalar representing a 32-bit UUID
//
// # Return Value
//
// Returns the new IOBluetoothSDPUUID object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/uuid32(_:)
func (_IOBluetoothSDPUUIDClass IOBluetoothSDPUUIDClass) Uuid32(uuid32 BluetoothSDPUUID32) IOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPUUIDClass.class), objc.Sel("uuid32:"), uuid32)
	return IOBluetoothSDPUUIDFromID(rv)
}
