// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBMutableDescriptor] class.
var (
	_CBMutableDescriptorClass     CBMutableDescriptorClass
	_CBMutableDescriptorClassOnce sync.Once
)

func getCBMutableDescriptorClass() CBMutableDescriptorClass {
	_CBMutableDescriptorClassOnce.Do(func() {
		_CBMutableDescriptorClass = CBMutableDescriptorClass{class: objc.GetClass("CBMutableDescriptor")}
	})
	return _CBMutableDescriptorClass
}

// GetCBMutableDescriptorClass returns the class object for CBMutableDescriptor.
func GetCBMutableDescriptorClass() CBMutableDescriptorClass {
	return getCBMutableDescriptorClass()
}

type CBMutableDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBMutableDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBMutableDescriptorClass) Alloc() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides additional information about a local peripheral’s
// characteristic.
//
// # Overview
//
// You use the [CBMutableDescriptor] class to create a local characteristic
// descriptor. After you create a descriptor and associate it with a local
// characteristic, you can publish it to the peripheral’s local database
// using the [CBPeripheralManager.AddService] method of the
// [CBPeripheralManager] class. This also publishes the characteristic and
// local service to which the descriptor belongs. After you publish a local
// descriptor, Core Bluetooth caches the descriptor and you can no longer make
// changes to it.
//
// [CBUUID] details predefined descriptor types and their corresponding value
// types. That said, only two of these are currently supported when creating
// local, mutable descriptors: the characteristic user description descriptor
// and the characteristic format descriptor. [CBUUID] declares these as the
// constants [CBUUIDCharacteristicUserDescriptionString] and
// [CBUUIDCharacteristicFormatString], respectively. The system automatically
// creates the extended properties descriptor and the client configuration
// descriptor, depending on the properties of the characteristic to which the
// descriptor belongs.
//
// # Creating a Mutable Descriptor
//
//   - [CBMutableDescriptor.InitWithTypeValue]: Creates a mutable descriptor with a specified value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor
//
// [CBUUIDCharacteristicFormatString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
// [CBUUIDCharacteristicUserDescriptionString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
type CBMutableDescriptor struct {
	CBDescriptor
}

// CBMutableDescriptorFromID constructs a [CBMutableDescriptor] from an objc.ID.
//
// An object that provides additional information about a local peripheral’s
// characteristic.
func CBMutableDescriptorFromID(id objc.ID) CBMutableDescriptor {
	return CBMutableDescriptor{CBDescriptor: CBDescriptorFromID(id)}
}

// NOTE: CBMutableDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBMutableDescriptor] class.
//
// # Creating a Mutable Descriptor
//
//   - [ICBMutableDescriptor.InitWithTypeValue]: Creates a mutable descriptor with a specified value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor
type ICBMutableDescriptor interface {
	ICBDescriptor

	// Topic: Creating a Mutable Descriptor

	// Creates a mutable descriptor with a specified value.
	InitWithTypeValue(UUID ICBUUID, value objectivec.IObject) CBMutableDescriptor
}

// Init initializes the instance.
func (c CBMutableDescriptor) Init() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBMutableDescriptor) Autorelease() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableDescriptor creates a new CBMutableDescriptor instance.
func NewCBMutableDescriptor() CBMutableDescriptor {
	class := getCBMutableDescriptorClass()
	rv := objc.Send[CBMutableDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a mutable descriptor with a specified value.
//
// UUID: A 128-bit UUID that identifies the characteristic. You must use only one of
// the two currently supported descriptor types:
// [CBUUIDCharacteristicUserDescriptionString] or
// [CBUUIDCharacteristicFormatString]. For more details about these descriptor
// types, see [CBUUID].
//
// value: The descriptor value to cache. You must provide a non-`nil` value. Once
// published, you can’t update the value dynamically.
//
// # Return Value
//
// A newly initialized mutable descriptor.
//
// # Discussion
//
// The value type of `value` depends on the type of descriptor:
//
// - The value type of [CBUUIDCharacteristicUserDescriptionString] is a string
// you use to provide a human-readable description of the characteristic’s
// value. - The value type of a [CBUUIDCharacteristicFormatString] is an
// [NSData] object that you use to specify how to format the
// characteristic’s value for presentation purposes.
//
// If you want to create a local characteristic format descriptor, the
// descriptor’s value must conform to the attribute value of the
// characteristic format descriptor as defined in the Bluetooth 4.0
// specification, Volume 3, Part G, Section 3.3.3.5.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor/init(type:value:)
//
// [CBUUIDCharacteristicFormatString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
// [CBUUIDCharacteristicUserDescriptionString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// [CBUUIDCharacteristicFormatString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
// [CBUUIDCharacteristicUserDescriptionString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
func NewCBMutableDescriptorWithTypeValue(UUID ICBUUID, value objectivec.IObject) CBMutableDescriptor {
	instance := getCBMutableDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:value:"), UUID, value)
	return CBMutableDescriptorFromID(rv)
}

// Creates a mutable descriptor with a specified value.
//
// UUID: A 128-bit UUID that identifies the characteristic. You must use only one of
// the two currently supported descriptor types:
// [CBUUIDCharacteristicUserDescriptionString] or
// [CBUUIDCharacteristicFormatString]. For more details about these descriptor
// types, see [CBUUID].
//
// value: The descriptor value to cache. You must provide a non-`nil` value. Once
// published, you can’t update the value dynamically.
//
// # Return Value
//
// A newly initialized mutable descriptor.
//
// # Discussion
//
// The value type of `value` depends on the type of descriptor:
//
// - The value type of [CBUUIDCharacteristicUserDescriptionString] is a string
// you use to provide a human-readable description of the characteristic’s
// value. - The value type of a [CBUUIDCharacteristicFormatString] is an
// [NSData] object that you use to specify how to format the
// characteristic’s value for presentation purposes.
//
// If you want to create a local characteristic format descriptor, the
// descriptor’s value must conform to the attribute value of the
// characteristic format descriptor as defined in the Bluetooth 4.0
// specification, Volume 3, Part G, Section 3.3.3.5.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor/init(type:value:)
//
// [CBUUIDCharacteristicFormatString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
// [CBUUIDCharacteristicUserDescriptionString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// [CBUUIDCharacteristicFormatString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
// [CBUUIDCharacteristicUserDescriptionString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
func (c CBMutableDescriptor) InitWithTypeValue(UUID ICBUUID, value objectivec.IObject) CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c.ID, objc.Sel("initWithType:value:"), UUID, value)
	return rv
}
