// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CBMutableCharacteristic] class.
var (
	_CBMutableCharacteristicClass     CBMutableCharacteristicClass
	_CBMutableCharacteristicClassOnce sync.Once
)

func getCBMutableCharacteristicClass() CBMutableCharacteristicClass {
	_CBMutableCharacteristicClassOnce.Do(func() {
		_CBMutableCharacteristicClass = CBMutableCharacteristicClass{class: objc.GetClass("CBMutableCharacteristic")}
	})
	return _CBMutableCharacteristicClass
}

// GetCBMutableCharacteristicClass returns the class object for CBMutableCharacteristic.
func GetCBMutableCharacteristicClass() CBMutableCharacteristicClass {
	return getCBMutableCharacteristicClass()
}

type CBMutableCharacteristicClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBMutableCharacteristicClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBMutableCharacteristicClass) Alloc() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A characteristic of a local peripheral’s service.
//
// # Overview
//
// [CBMutableCharacteristic] objects represent the characteristics of a local
// peripheral’s service. This class adds write access to many of the
// properties in the [CBCharacteristic] class, which it inherits from.
//
// You use this class to create a characteristic and to set its properties and
// permissions as desired. After you create and add a characteristic to a
// local service, you can publish it (and the service) to the peripheral’s
// local database with the [CBPeripheralManager.AddService] method of the
// [CBPeripheralManager] class. After you publish a characteristic, Core
// Bluetooth caches the characteristic and you can’t make changes to it.
//
// # Creating a Mutable Characteristic
//
//   - [CBMutableCharacteristic.InitWithTypePropertiesValuePermissions]: Creates a mutable characteristic with specified permissions, properties, and value.
//
// # Managing a Mutable Characteristic
//
//   - [CBMutableCharacteristic.Permissions]: The permissions of the characteristic value.
//   - [CBMutableCharacteristic.SetPermissions]
//   - [CBMutableCharacteristic.SubscribedCentrals]: A list of centrals that are currently subscribed to the characteristic’s value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic
type CBMutableCharacteristic struct {
	CBCharacteristic
}

// CBMutableCharacteristicFromID constructs a [CBMutableCharacteristic] from an objc.ID.
//
// A characteristic of a local peripheral’s service.
func CBMutableCharacteristicFromID(id objc.ID) CBMutableCharacteristic {
	return CBMutableCharacteristic{CBCharacteristic: CBCharacteristicFromID(id)}
}

// NOTE: CBMutableCharacteristic adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBMutableCharacteristic] class.
//
// # Creating a Mutable Characteristic
//
//   - [ICBMutableCharacteristic.InitWithTypePropertiesValuePermissions]: Creates a mutable characteristic with specified permissions, properties, and value.
//
// # Managing a Mutable Characteristic
//
//   - [ICBMutableCharacteristic.Permissions]: The permissions of the characteristic value.
//   - [ICBMutableCharacteristic.SetPermissions]
//   - [ICBMutableCharacteristic.SubscribedCentrals]: A list of centrals that are currently subscribed to the characteristic’s value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic
type ICBMutableCharacteristic interface {
	ICBCharacteristic

	// Topic: Creating a Mutable Characteristic

	// Creates a mutable characteristic with specified permissions, properties, and value.
	InitWithTypePropertiesValuePermissions(UUID ICBUUID, properties CBCharacteristicProperties, value foundation.NSData, permissions CBAttributePermissions) CBMutableCharacteristic

	// Topic: Managing a Mutable Characteristic

	// The permissions of the characteristic value.
	Permissions() CBAttributePermissions
	SetPermissions(value CBAttributePermissions)
	// A list of centrals that are currently subscribed to the characteristic’s value.
	SubscribedCentrals() []CBCentral
}

// Init initializes the instance.
func (c CBMutableCharacteristic) Init() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBMutableCharacteristic) Autorelease() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableCharacteristic creates a new CBMutableCharacteristic instance.
func NewCBMutableCharacteristic() CBMutableCharacteristic {
	class := getCBMutableCharacteristicClass()
	rv := objc.Send[CBMutableCharacteristic](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a mutable characteristic with specified permissions, properties,
// and value.
//
// UUID: A 128-bit UUID that identifies the characteristic.
//
// properties: The properties of the characteristic.
//
// value: The characteristic value to cache. If `nil`, the value is dynamic and the
// peripheral manager fetches it on demand.
//
// permissions: The permissions of the characteristic value.
//
// # Return Value
//
// A newly initialized mutable characteristic.
//
// # Discussion
//
// If you specify a value for the characteristic, the characteristic caches
// the value and sets its properties and permissions to
// [CBCharacteristicPropertyRead] and [CBAttributePermissionsReadable],
// respectively. Therefore, if you need the value of a characteristic to be
// writeable, or if you expect the value to change during the lifetime of the
// published service to which the characteristic belongs, you must specify the
// value as `nil`. This ensures that the characteristic treats the value
// dynamically. With a dynamic value, the peripheral manager requests the
// value whenever the peripheral manager receives a read or write request from
// a central. The peripheral does this by calling the
// [PeripheralManagerDidReceiveReadRequest] and
// [PeripheralManagerDidReceiveWriteRequests] methods of its delegate object,
// respectively.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/init(type:properties:value:permissions:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func NewCBMutableCharacteristicWithTypePropertiesValuePermissions(UUID ICBUUID, properties CBCharacteristicProperties, value foundation.NSData, permissions CBAttributePermissions) CBMutableCharacteristic {
	instance := getCBMutableCharacteristicClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:properties:value:permissions:"), UUID, properties, value, permissions)
	return CBMutableCharacteristicFromID(rv)
}

// Creates a mutable characteristic with specified permissions, properties,
// and value.
//
// UUID: A 128-bit UUID that identifies the characteristic.
//
// properties: The properties of the characteristic.
//
// value: The characteristic value to cache. If `nil`, the value is dynamic and the
// peripheral manager fetches it on demand.
//
// permissions: The permissions of the characteristic value.
//
// # Return Value
//
// A newly initialized mutable characteristic.
//
// # Discussion
//
// If you specify a value for the characteristic, the characteristic caches
// the value and sets its properties and permissions to
// [CBCharacteristicPropertyRead] and [CBAttributePermissionsReadable],
// respectively. Therefore, if you need the value of a characteristic to be
// writeable, or if you expect the value to change during the lifetime of the
// published service to which the characteristic belongs, you must specify the
// value as `nil`. This ensures that the characteristic treats the value
// dynamically. With a dynamic value, the peripheral manager requests the
// value whenever the peripheral manager receives a read or write request from
// a central. The peripheral does this by calling the
// [PeripheralManagerDidReceiveReadRequest] and
// [PeripheralManagerDidReceiveWriteRequests] methods of its delegate object,
// respectively.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/init(type:properties:value:permissions:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBMutableCharacteristic) InitWithTypePropertiesValuePermissions(UUID ICBUUID, properties CBCharacteristicProperties, value foundation.NSData, permissions CBAttributePermissions) CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](c.ID, objc.Sel("initWithType:properties:value:permissions:"), UUID, properties, value, permissions)
	return rv
}

// The permissions of the characteristic value.
//
// # Discussion
//
// Characteristic permissions represent the read, write, and encryption
// permissions for a characteristic’s value. For a complete list and
// discussion of the available characteristic permissions, see
// [CBAttributePermissions].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/permissions
//
// [CBAttributePermissions]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions
func (c CBMutableCharacteristic) Permissions() CBAttributePermissions {
	rv := objc.Send[CBAttributePermissions](c.ID, objc.Sel("permissions"))
	return CBAttributePermissions(rv)
}
func (c CBMutableCharacteristic) SetPermissions(value CBAttributePermissions) {
	objc.Send[struct{}](c.ID, objc.Sel("setPermissions:"), value)
}

// A list of centrals that are currently subscribed to the characteristic’s
// value.
//
// # Discussion
//
// The value of this property is an array of [CBCentral] objects that
// currently subscribe to the characteristic’s value. The array is empty if
// the characteristic isn’t configured to support notifications or
// indications. Even if the characteristic’s configuration supports
// notifications or indications, the array is empty if centrals aren’t
// subscribing to the characteristic’s value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/subscribedCentrals
func (c CBMutableCharacteristic) SubscribedCentrals() []CBCentral {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("subscribedCentrals"))
	return objc.ConvertSlice(rv, func(id objc.ID) CBCentral {
		return CBCentralFromID(id)
	})
}
