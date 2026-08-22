// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBUUID] class.
var (
	_CBUUIDClass     CBUUIDClass
	_CBUUIDClassOnce sync.Once
)

func getCBUUIDClass() CBUUIDClass {
	_CBUUIDClassOnce.Do(func() {
		_CBUUIDClass = CBUUIDClass{class: objc.GetClass("CBUUID")}
	})
	return _CBUUIDClass
}

// GetCBUUIDClass returns the class object for CBUUID.
func GetCBUUIDClass() CBUUIDClass {
	return getCBUUIDClass()
}

type CBUUIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBUUIDClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBUUIDClass) Alloc() CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A universally unique identifier, as defined by Bluetooth standards.
//
// # Overview
//
// Instances of the [CBUUID] class represent the 128-bit universally unique
// identifiers (UUIDs) of attributes used in Bluetooth low energy
// communication, such as a peripheral’s services, characteristics, and
// descriptors. This class provides a number of factory methods for dealing
// with long UUIDs when developing your app. For example, instead of passing
// around the string representation of a 128-bit Bluetooth low energy
// attribute in your code, you can create a [CBUUID] object that represents
// it, and pass that around instead.
//
// The Bluetooth Special Interest Group (SIG) publishes a list of
// commonly-used UUIDs, many of which are 16- or 32-bits for convenience. The
// [CBUUID] class provides methods that automatically transform these
// predefined shorter UUIDs into their 128-bit equivalent UUIDs. When you
// create a [CBUUID] object from a predefined 16- or 32-bit UUID, Core
// Bluetooth pre-fills the rest of the 128-bit UUID with the Bluetooth base
// UUID, as defined in the Bluetooth 4.0 specification, Volume 3, Part F,
// Section 3.2.1.
//
// In addition to providing methods for creating [CBUUID] objects, this class
// defines constants that represent the UUIDs of the Bluetooth-defined
// characteristic descriptors, as defined in the Bluetooth 4.0 specification,
// Volume 3, Part G, Section 3.3.3.
//
// # Inspecting CBUUID Properties
//
//   - [CBUUID.Data]: The data of the UUID.
//   - [CBUUID.UUIDString]: The UUID represented as a string.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID
type CBUUID struct {
	objectivec.Object
}

// CBUUIDFromID constructs a [CBUUID] from an objc.ID.
//
// A universally unique identifier, as defined by Bluetooth standards.
func CBUUIDFromID(id objc.ID) CBUUID {
	return CBUUID{objectivec.Object{ID: id}}
}

// NOTE: CBUUID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBUUID] class.
//
// # Inspecting CBUUID Properties
//
//   - [ICBUUID.Data]: The data of the UUID.
//   - [ICBUUID.UUIDString]: The UUID represented as a string.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID
type ICBUUID interface {
	objectivec.IObject

	// Topic: Inspecting CBUUID Properties

	// The data of the UUID.
	Data() foundation.NSData
	// The UUID represented as a string.
	UUIDString() string
}

// Init initializes the instance.
func (c CBUUID) Init() CBUUID {
	rv := objc.Send[CBUUID](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBUUID) Autorelease() CBUUID {
	rv := objc.Send[CBUUID](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBUUID creates a new CBUUID instance.
func NewCBUUID() CBUUID {
	class := getCBUUIDClass()
	rv := objc.Send[CBUUID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID data
// container.
//
// theData: Data containing a 16-, 32-, or 128-bit UUID.
//
// # Return Value
//
// A new [CBUUID] object for the specified UUID data.
//
// # Discussion
//
// This method is useful when handling the UUID of a Bluetooth attribute in
// raw bytes.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(data:)
func NewCBUUIDWithData(theData foundation.NSData) CBUUID {
	rv := objc.Send[objc.ID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithData:"), theData)
	return CBUUIDFromID(rv)
}

// Creates a Core Bluetooth UUID object from a Foundation UUID object.
//
// theUUID: A UUID represented by an [NSUUID] object.
//
// # Return Value
//
// A new [CBUUID] object for the specified UUID.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(nsuuid:)
//
// [NSUUID]: https://developer.apple.com/documentation/Foundation/NSUUID
func NewCBUUIDWithNSUUID(theUUID foundation.NSUUID) CBUUID {
	rv := objc.Send[objc.ID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithNSUUID:"), theUUID)
	return CBUUIDFromID(rv)
}

// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID
// string.
//
// theString: A string containing a 16-, 32-, or 128-bit UUID.
//
// # Return Value
//
// A new [CBUUID] object for the specified UUID string.
//
// # Discussion
//
// Specify 128-bit UUIDs as a string of hexadecimal digits punctuated by
// hyphens, for example, 68753A44-4D6F-1226-9C60-0050E4C00067. Specify 16- or
// 32-bit UUIDs as a string of 4 or 8 hexadecimal digits, respectively. For an
// example of how to use this method, see [Services and Characteristics Are
// Identified by UUIDs] and [Create Your Own UUIDs for Custom Services and
// Characteristics].
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(string:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
// [Create Your Own UUIDs for Custom Services and Characteristics]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/PerformingCommonPeripheralRoleTasks/PerformingCommonPeripheralRoleTasks.html#//apple_ref/doc/uid/TP40013257-CH4-SW9
// [Services and Characteristics Are Identified by UUIDs]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/PerformingCommonPeripheralRoleTasks/PerformingCommonPeripheralRoleTasks.html#//apple_ref/doc/uid/TP40013257-CH4-SW8
func NewCBUUIDWithString(theString string) CBUUID {
	rv := objc.Send[objc.ID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithString:"), objc.String(theString))
	return CBUUIDFromID(rv)
}

// The data of the UUID.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/data
func (c CBUUID) Data() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The UUID represented as a string.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/uuidString
func (c CBUUID) UUIDString() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("UUIDString"))
	return foundation.NSStringFromID(rv).String()
}
