// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CBCharacteristic] class.
var (
	_CBCharacteristicClass     CBCharacteristicClass
	_CBCharacteristicClassOnce sync.Once
)

func getCBCharacteristicClass() CBCharacteristicClass {
	_CBCharacteristicClassOnce.Do(func() {
		_CBCharacteristicClass = CBCharacteristicClass{class: objc.GetClass("CBCharacteristic")}
	})
	return _CBCharacteristicClass
}

// GetCBCharacteristicClass returns the class object for CBCharacteristic.
func GetCBCharacteristicClass() CBCharacteristicClass {
	return getCBCharacteristicClass()
}

type CBCharacteristicClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBCharacteristicClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBCharacteristicClass) Alloc() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A characteristic of a remote peripheral’s service.
//
// # Overview
//
// [CBCharacteristic] and its subclass [CBMutableCharacteristic] represent
// further information about a peripheral’s service. In particular,
// [CBCharacteristic] objects represent the characteristics of a remote
// peripheral’s service. A characteristic contains a single value and any
// number of descriptors describing that value. The properties of a
// characteristic determine how you can use a characteristic’s value, and
// how you access the descriptors.
//
// # Identifying a Characteristic
//
//   - [CBCharacteristic.Service]: The service to which this characteristic belongs.
//
// # Accessing Characteristic Data
//
//   - [CBCharacteristic.Value]: The value of the characteristic.
//   - [CBCharacteristic.Descriptors]: A list of the descriptors discovered in this characteristic.
//   - [CBCharacteristic.Properties]: The properties of the characteristic.
//   - [CBCharacteristic.IsNotifying]: A Boolean value that indicates whether the characteristic is currently notifying a subscribed central of its value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic
type CBCharacteristic struct {
	CBAttribute
}

// CBCharacteristicFromID constructs a [CBCharacteristic] from an objc.ID.
//
// A characteristic of a remote peripheral’s service.
func CBCharacteristicFromID(id objc.ID) CBCharacteristic {
	return CBCharacteristic{CBAttribute: CBAttributeFromID(id)}
}

// NOTE: CBCharacteristic adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBCharacteristic] class.
//
// # Identifying a Characteristic
//
//   - [ICBCharacteristic.Service]: The service to which this characteristic belongs.
//
// # Accessing Characteristic Data
//
//   - [ICBCharacteristic.Value]: The value of the characteristic.
//   - [ICBCharacteristic.Descriptors]: A list of the descriptors discovered in this characteristic.
//   - [ICBCharacteristic.Properties]: The properties of the characteristic.
//   - [ICBCharacteristic.IsNotifying]: A Boolean value that indicates whether the characteristic is currently notifying a subscribed central of its value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic
type ICBCharacteristic interface {
	ICBAttribute

	// Topic: Identifying a Characteristic

	// The service to which this characteristic belongs.
	Service() ICBService

	// Topic: Accessing Characteristic Data

	// The value of the characteristic.
	Value() foundation.NSData
	// A list of the descriptors discovered in this characteristic.
	Descriptors() []CBDescriptor
	// The properties of the characteristic.
	Properties() CBCharacteristicProperties
	// A Boolean value that indicates whether the characteristic is currently notifying a subscribed central of its value.
	IsNotifying() bool
}

// Init initializes the instance.
func (c CBCharacteristic) Init() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBCharacteristic) Autorelease() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCharacteristic creates a new CBCharacteristic instance.
func NewCBCharacteristic() CBCharacteristic {
	class := getCBCharacteristicClass()
	rv := objc.Send[CBCharacteristic](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The service to which this characteristic belongs.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/service
func (c CBCharacteristic) Service() ICBService {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("service"))
	return CBServiceFromID(objc.ID(rv))
}

// The value of the characteristic.
//
// # Discussion
//
// This property contains the value of the characteristic. For example, a
// temperature measurement characteristic of a health thermometer service may
// have a value that indicates a temperature in Celsius.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/value
func (c CBCharacteristic) Value() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("value"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A list of the descriptors discovered in this characteristic.
//
// # Discussion
//
// The value of this property is an array of [CBDescriptor] objects that
// represent a characteristic’s descriptors. Characteristic descriptors
// provide more information about a characteristic’s value. For example,
// they may describe the value in human-readable form and describe how to
// format the value for presentation purposes. For more information about
// characteristic descriptors, see [CBDescriptor].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/descriptors
func (c CBCharacteristic) Descriptors() []CBDescriptor {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("descriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) CBDescriptor {
		return CBDescriptorFromID(id)
	})
}

// The properties of the characteristic.
//
// # Discussion
//
// The properties of a characteristic determine the access to and use of the
// characteristic’s value and descriptors. For a list of the possible values
// representing the properties of a characteristic, see
// [CBCharacteristicProperties].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/properties
//
// [CBCharacteristicProperties]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicProperties
func (c CBCharacteristic) Properties() CBCharacteristicProperties {
	rv := objc.Send[CBCharacteristicProperties](c.ID, objc.Sel("properties"))
	return CBCharacteristicProperties(rv)
}

// A Boolean value that indicates whether the characteristic is currently
// notifying a subscribed central of its value.
//
// # Discussion
//
// This value is true if you enabled notifications or indications for the
// characteristic by successfully calling the
// [CBPeripheral.SetNotifyValueForCharacteristic] method of the [CBPeripheral]
// class. In this case, the peripheral updates its connected central that
// whenever the characteristic’s value changes.
//
// If the value of the property is false, notifications (or indications)
// aren’t enabled for the characteristic, and the peripheral doesn’t
// update its connected central when the characteristic’s value changes.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/isNotifying
func (c CBCharacteristic) IsNotifying() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isNotifying"))
	return rv
}
