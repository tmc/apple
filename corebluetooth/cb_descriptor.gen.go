// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBDescriptor] class.
var (
	_CBDescriptorClass     CBDescriptorClass
	_CBDescriptorClassOnce sync.Once
)

func getCBDescriptorClass() CBDescriptorClass {
	_CBDescriptorClassOnce.Do(func() {
		_CBDescriptorClass = CBDescriptorClass{class: objc.GetClass("CBDescriptor")}
	})
	return _CBDescriptorClass
}

// GetCBDescriptorClass returns the class object for CBDescriptor.
func GetCBDescriptorClass() CBDescriptorClass {
	return getCBDescriptorClass()
}

type CBDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBDescriptorClass) Alloc() CBDescriptor {
	rv := objc.Send[CBDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides further information about a remote peripheral’s
// characteristic.
//
// # Overview
//
// [CBDescriptor] and its subclass [CBMutableDescriptor] represent a
// descriptor of a peripheral’s characteristic. In partcular, [CBDescriptor]
// objects represent the descriptors of a remote peripheral’s
// characteristic. Descriptors provide further information about a
// characteristic’s value. For example, they may describe the value in
// human-readable form and describe how to format the value for presentation
// purposes. Characteristic descriptors also indicate whether a
// characteristic’s value indicates or notifies a client (a central) when
// the value of the characteristic changes.
//
// [CBUUID] details six predefined descriptors and their corresponding value
// types. [CBDescriptor] lists the predefined descriptors and the [CBUUID]
// constants that represent them.
//
// [Table data omitted]
//
// # Identifying a Descriptor
//
//   - [CBDescriptor.Characteristic]: The characteristic to which this descriptor belongs.
//
// # Accessing Descriptor Data
//
//   - [CBDescriptor.Value]: The value of the descriptor.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor
type CBDescriptor struct {
	CBAttribute
}

// CBDescriptorFromID constructs a [CBDescriptor] from an objc.ID.
//
// An object that provides further information about a remote peripheral’s
// characteristic.
func CBDescriptorFromID(id objc.ID) CBDescriptor {
	return CBDescriptor{CBAttribute: CBAttributeFromID(id)}
}

// NOTE: CBDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBDescriptor] class.
//
// # Identifying a Descriptor
//
//   - [ICBDescriptor.Characteristic]: The characteristic to which this descriptor belongs.
//
// # Accessing Descriptor Data
//
//   - [ICBDescriptor.Value]: The value of the descriptor.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor
type ICBDescriptor interface {
	ICBAttribute

	// Topic: Identifying a Descriptor

	// The characteristic to which this descriptor belongs.
	Characteristic() ICBCharacteristic

	// Topic: Accessing Descriptor Data

	// The value of the descriptor.
	Value() objectivec.IObject
}

// Init initializes the instance.
func (c CBDescriptor) Init() CBDescriptor {
	rv := objc.Send[CBDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBDescriptor) Autorelease() CBDescriptor {
	rv := objc.Send[CBDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBDescriptor creates a new CBDescriptor instance.
func NewCBDescriptor() CBDescriptor {
	class := getCBDescriptorClass()
	rv := objc.Send[CBDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The characteristic to which this descriptor belongs.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor/characteristic
func (c CBDescriptor) Characteristic() ICBCharacteristic {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("characteristic"))
	return CBCharacteristicFromID(objc.ID(rv))
}

// The value of the descriptor.
//
// # Discussion
//
// The documentation for [CBUUID] details the value types for the various
// descriptor types.
//
// You can read the value of a descriptor by calling the
// [CBPeripheral.ReadValueForDescriptor] method of the [CBPeripheral] class.
// You can write the value of a descriptor by calling the
// [CBPeripheral.WriteValueForDescriptor] method of the [CBPeripheral] class.
// You can’t, however, use the [CBPeripheral.WriteValueForDescriptor] method
// to write the value of a client configuration descriptor
// ([CBUUIDClientCharacteristicConfigurationString]). Instead, you use the
// [CBPeripheral.SetNotifyValueForCharacteristic] method of the [CBPeripheral]
// class to configure client indications or notifications of a
// characteristic’s value on a server.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor/value
//
// [CBUUIDClientCharacteristicConfigurationString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDClientCharacteristicConfigurationString
func (c CBDescriptor) Value() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("value"))
	return objectivec.Object{ID: rv}
}
