// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBPeripheral] class.
var (
	_CBPeripheralClass     CBPeripheralClass
	_CBPeripheralClassOnce sync.Once
)

func getCBPeripheralClass() CBPeripheralClass {
	_CBPeripheralClassOnce.Do(func() {
		_CBPeripheralClass = CBPeripheralClass{class: objc.GetClass("CBPeripheral")}
	})
	return _CBPeripheralClass
}

// GetCBPeripheralClass returns the class object for CBPeripheral.
func GetCBPeripheralClass() CBPeripheralClass {
	return getCBPeripheralClass()
}

type CBPeripheralClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBPeripheralClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBPeripheralClass) Alloc() CBPeripheral {
	rv := objc.Send[CBPeripheral](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A remote peripheral device.
//
// # Overview
//
// The [CBPeripheral] class represents remote peripheral devices that your app
// discovers with a central manager (an instance of [CBCentralManager]).
// Peripherals use universally unique identifiers (UUIDs), represented by
// [NSUUID] objects, to identify themselves. Peripherals may contain one or
// more services or provide useful information about their connected signal
// strength.
//
// You use this class to discover, explore, and interact with the services
// available on a remote peripheral that supports Bluetooth low energy. A
// service encapsulates the way part of the device behaves. For example, one
// service of a heart rate monitor may be to expose heart rate data from a
// sensor. Services themselves contain of characteristics or included services
// (references to other services). Characteristics provide further details
// about a peripheral’s service. For example, the heart rate service may
// contain multiple characteristics. One characteristic could describe the
// intended body location of the device’s heart rate sensor, and another
// characteristic could transmit the heart rate measurement data. Finally,
// characteristics contain any number of descriptors that provide more
// information about the characteristic’s value, such as a human-readable
// description and a way to format the value.
//
// # Identifying a Peripheral
//
//   - [CBPeripheral.Name]: The name of the peripheral.
//   - [CBPeripheral.Delegate]: The delegate object specified to receive peripheral events.
//   - [CBPeripheral.SetDelegate]
//
// # Discovering Services
//
//   - [CBPeripheral.DiscoverServices]: Discovers the specified services of the peripheral.
//   - [CBPeripheral.DiscoverIncludedServicesForService]: Discovers the specified included services of a previously-discovered service.
//   - [CBPeripheral.Services]: A list of a peripheral’s discovered services.
//
// # Discovering Characteristics and Descriptors
//
//   - [CBPeripheral.DiscoverCharacteristicsForService]: Discovers the specified characteristics of a service.
//   - [CBPeripheral.DiscoverDescriptorsForCharacteristic]: Discovers the descriptors of a characteristic.
//
// # Reading Characteristic and Descriptor Values
//
//   - [CBPeripheral.ReadValueForCharacteristic]: Retrieves the value of a specified characteristic.
//   - [CBPeripheral.ReadValueForDescriptor]: Retrieves the value of a specified characteristic descriptor.
//
// # Writing Characteristic and Descriptor Values
//
//   - [CBPeripheral.WriteValueForCharacteristicType]: Writes the value of a characteristic.
//   - [CBPeripheral.WriteValueForDescriptor]: Writes the value of a characteristic descriptor.
//   - [CBPeripheral.MaximumWriteValueLengthForType]: The maximum amount of data, in bytes, you can send to a characteristic in a single write type.
//
// # Setting Notifications for a Characteristic’s Value
//
//   - [CBPeripheral.SetNotifyValueForCharacteristic]: Sets notifications or indications for the value of a specified characteristic.
//
// # Monitoring a Peripheral’s Connection State
//
//   - [CBPeripheral.State]: The connection state of the peripheral.
//   - [CBPeripheral.CanSendWriteWithoutResponse]: A Boolean value that indicates whether the remote device can send a write without a response.
//
// # Accessing a Peripheral’s Signal Strength
//
//   - [CBPeripheral.ReadRSSI]: Retrieves the current RSSI value for the peripheral while connected to the central manager.
//
// # Working with L2CAP Channels
//
//   - [CBPeripheral.OpenL2CAPChannel]: Attempts to open an L2CAP channel to the peripheral using the supplied Protocol/Service Multiplexer (PSM).
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral
//
// [NSUUID]: https://developer.apple.com/documentation/Foundation/NSUUID
type CBPeripheral struct {
	CBPeer
}

// CBPeripheralFromID constructs a [CBPeripheral] from an objc.ID.
//
// A remote peripheral device.
func CBPeripheralFromID(id objc.ID) CBPeripheral {
	return CBPeripheral{CBPeer: CBPeerFromID(id)}
}

// NOTE: CBPeripheral adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBPeripheral] class.
//
// # Identifying a Peripheral
//
//   - [ICBPeripheral.Name]: The name of the peripheral.
//   - [ICBPeripheral.Delegate]: The delegate object specified to receive peripheral events.
//   - [ICBPeripheral.SetDelegate]
//
// # Discovering Services
//
//   - [ICBPeripheral.DiscoverServices]: Discovers the specified services of the peripheral.
//   - [ICBPeripheral.DiscoverIncludedServicesForService]: Discovers the specified included services of a previously-discovered service.
//   - [ICBPeripheral.Services]: A list of a peripheral’s discovered services.
//
// # Discovering Characteristics and Descriptors
//
//   - [ICBPeripheral.DiscoverCharacteristicsForService]: Discovers the specified characteristics of a service.
//   - [ICBPeripheral.DiscoverDescriptorsForCharacteristic]: Discovers the descriptors of a characteristic.
//
// # Reading Characteristic and Descriptor Values
//
//   - [ICBPeripheral.ReadValueForCharacteristic]: Retrieves the value of a specified characteristic.
//   - [ICBPeripheral.ReadValueForDescriptor]: Retrieves the value of a specified characteristic descriptor.
//
// # Writing Characteristic and Descriptor Values
//
//   - [ICBPeripheral.WriteValueForCharacteristicType]: Writes the value of a characteristic.
//   - [ICBPeripheral.WriteValueForDescriptor]: Writes the value of a characteristic descriptor.
//   - [ICBPeripheral.MaximumWriteValueLengthForType]: The maximum amount of data, in bytes, you can send to a characteristic in a single write type.
//
// # Setting Notifications for a Characteristic’s Value
//
//   - [ICBPeripheral.SetNotifyValueForCharacteristic]: Sets notifications or indications for the value of a specified characteristic.
//
// # Monitoring a Peripheral’s Connection State
//
//   - [ICBPeripheral.State]: The connection state of the peripheral.
//   - [ICBPeripheral.CanSendWriteWithoutResponse]: A Boolean value that indicates whether the remote device can send a write without a response.
//
// # Accessing a Peripheral’s Signal Strength
//
//   - [ICBPeripheral.ReadRSSI]: Retrieves the current RSSI value for the peripheral while connected to the central manager.
//
// # Working with L2CAP Channels
//
//   - [ICBPeripheral.OpenL2CAPChannel]: Attempts to open an L2CAP channel to the peripheral using the supplied Protocol/Service Multiplexer (PSM).
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral
type ICBPeripheral interface {
	ICBPeer

	// Topic: Identifying a Peripheral

	// The name of the peripheral.
	Name() string
	// The delegate object specified to receive peripheral events.
	Delegate() CBPeripheralDelegate
	SetDelegate(value CBPeripheralDelegate)

	// Topic: Discovering Services

	// Discovers the specified services of the peripheral.
	DiscoverServices(serviceUUIDs []CBUUID)
	// Discovers the specified included services of a previously-discovered service.
	DiscoverIncludedServicesForService(includedServiceUUIDs []CBUUID, service ICBService)
	// A list of a peripheral’s discovered services.
	Services() []CBService

	// Topic: Discovering Characteristics and Descriptors

	// Discovers the specified characteristics of a service.
	DiscoverCharacteristicsForService(characteristicUUIDs []CBUUID, service ICBService)
	// Discovers the descriptors of a characteristic.
	DiscoverDescriptorsForCharacteristic(characteristic ICBCharacteristic)

	// Topic: Reading Characteristic and Descriptor Values

	// Retrieves the value of a specified characteristic.
	ReadValueForCharacteristic(characteristic ICBCharacteristic)
	// Retrieves the value of a specified characteristic descriptor.
	ReadValueForDescriptor(descriptor ICBDescriptor)

	// Topic: Writing Characteristic and Descriptor Values

	// Writes the value of a characteristic.
	WriteValueForCharacteristicType(data foundation.NSData, characteristic ICBCharacteristic, type_ CBCharacteristicWriteType)
	// Writes the value of a characteristic descriptor.
	WriteValueForDescriptor(data foundation.NSData, descriptor ICBDescriptor)
	// The maximum amount of data, in bytes, you can send to a characteristic in a single write type.
	MaximumWriteValueLengthForType(type_ CBCharacteristicWriteType) uint

	// Topic: Setting Notifications for a Characteristic’s Value

	// Sets notifications or indications for the value of a specified characteristic.
	SetNotifyValueForCharacteristic(enabled bool, characteristic ICBCharacteristic)

	// Topic: Monitoring a Peripheral’s Connection State

	// The connection state of the peripheral.
	State() CBPeripheralState
	// A Boolean value that indicates whether the remote device can send a write without a response.
	CanSendWriteWithoutResponse() bool

	// Topic: Accessing a Peripheral’s Signal Strength

	// Retrieves the current RSSI value for the peripheral while connected to the central manager.
	ReadRSSI()

	// Topic: Working with L2CAP Channels

	// Attempts to open an L2CAP channel to the peripheral using the supplied Protocol/Service Multiplexer (PSM).
	OpenL2CAPChannel(PSM CBL2CAPPSM)
}

// Init initializes the instance.
func (c CBPeripheral) Init() CBPeripheral {
	rv := objc.Send[CBPeripheral](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBPeripheral) Autorelease() CBPeripheral {
	rv := objc.Send[CBPeripheral](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeripheral creates a new CBPeripheral instance.
func NewCBPeripheral() CBPeripheral {
	class := getCBPeripheralClass()
	rv := objc.Send[CBPeripheral](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Discovers the specified services of the peripheral.
//
// serviceUUIDs: An array of [CBUUID] objects that you are interested in. Each [CBUUID]
// object represents a UUID that identifies the type of service you want to
// discover.
//
// # Discussion
//
// You can provide an array of [CBUUID] objects—representing service
// UUIDs—in the `serviceUUIDs` parameter. When you do, the peripheral
// returns only the services of the peripheral that match the provided UUIDs.
//
// When the peripheral discovers one or more services, it calls the
// [PeripheralDidDiscoverServices]: method of its delegate object. After a
// peripheral discovers services, you can access them through the
// peripheral’s [CBPeripheral.Services] property.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverServices(_:)
func (c CBPeripheral) DiscoverServices(serviceUUIDs []CBUUID) {
	objc.Send[objc.ID](c.ID, objc.Sel("discoverServices:"), objectivec.IObjectSliceToNSArray(serviceUUIDs))
}

// Discovers the specified included services of a previously-discovered
// service.
//
// includedServiceUUIDs: An array of [CBUUID] objects that you are interested in. Here, each
// [CBUUID] object represents a UUID that identifies the type of included
// service you want to discover.
//
// service: The previously-discovered service whose included services you want to
// discover.
//
// # Discussion
//
// You can provide an array of [CBUUID] objects—representing included
// service UUIDs—in the `includedServiceUUIDs` parameter. When you do, the
// peripheral returns only the services of the peripheral that match the
// provided UUIDs.
//
// When the peripheral discovers one or more included services of the
// specified service, it calls the
// [PeripheralDidDiscoverIncludedServicesForServiceError] method of its
// delegate object. After the service discovers its included services, you can
// access them through the service’s [CBService.IncludedServices] property.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverIncludedServices(_:for:)
func (c CBPeripheral) DiscoverIncludedServicesForService(includedServiceUUIDs []CBUUID, service ICBService) {
	objc.Send[objc.ID](c.ID, objc.Sel("discoverIncludedServices:forService:"), objectivec.IObjectSliceToNSArray(includedServiceUUIDs), service)
}

// Discovers the specified characteristics of a service.
//
// characteristicUUIDs: An array of [CBUUID] objects that you are interested in. Each [CBUUID]
// object represents a UUID that identifies the type of a characteristic you
// want to discover.
//
// service: The service whose characteristics you want to discover.
//
// # Discussion
//
// You can provide an array of [CBUUID] objects—representing characteristic
// UUIDs— in the `characteristicUUIDs` parameter. When you do, the
// peripheral returns only the characteristics of the service that match the
// provided UUIDs. If the `characteristicUUIDs` parameter is `nil`, this
// method returns all characteristics of the service.
//
// When the peripheral discovers one or more characteristics of the specified
// service, it calls the [PeripheralDidDiscoverCharacteristicsForServiceError]
// method of its delegate object. After the peripheral discovers the
// service’s characteristics, you can access them through the service’s
// [CBService.Characteristics] property.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverCharacteristics(_:for:)
func (c CBPeripheral) DiscoverCharacteristicsForService(characteristicUUIDs []CBUUID, service ICBService) {
	objc.Send[objc.ID](c.ID, objc.Sel("discoverCharacteristics:forService:"), objectivec.IObjectSliceToNSArray(characteristicUUIDs), service)
}

// Discovers the descriptors of a characteristic.
//
// characteristic: The characteristic whose descriptors you want to discover.
//
// # Discussion
//
// When the peripheral discovers one or more descriptors of the specified
// characteristic, it calls the
// [PeripheralDidDiscoverDescriptorsForCharacteristicError] method of its
// delegate object. After the peripheral discovers the descriptors of the
// characteristic, you can access them through the characteristic’s
// [CBCharacteristic.Descriptors] property.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverDescriptors(for:)
func (c CBPeripheral) DiscoverDescriptorsForCharacteristic(characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c.ID, objc.Sel("discoverDescriptorsForCharacteristic:"), characteristic)
}

// Retrieves the value of a specified characteristic.
//
// characteristic: The characteristic whose value you want to read.
//
// # Discussion
//
// When you call this method to read the value of a characteristic, the
// peripheral calls the [PeripheralDidUpdateValueForDescriptorError] method of
// its delegate object. If the peripheral successfully reads the value of the
// characteristic, you can access it through the characteristic’s
// [CBCharacteristic.Value] property.
//
// Not all characteristics have a readable value. You can determine whether a
// characteristic’s value is readable by accessing the relevant properties
// of the [CBCharacteristicProperties] enumeration.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readValue(for:)-6u2kr
//
// [CBCharacteristicProperties]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicProperties
func (c CBPeripheral) ReadValueForCharacteristic(characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c.ID, objc.Sel("readValueForCharacteristic:"), characteristic)
}

// Retrieves the value of a specified characteristic descriptor.
//
// descriptor: The characteristic descriptor whose value you want to read.
//
// # Discussion
//
// When you call this method to read the value of a characteristic descriptor,
// the peripheral calls the [PeripheralDidUpdateValueForDescriptorError]
// method of its delegate object. If the peripheral successfully retrieves the
// value of the characteristic descriptor, you can access it through the
// characteristic descriptor’s [CBDescriptor.Value] property.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readValue(for:)-91hhp
func (c CBPeripheral) ReadValueForDescriptor(descriptor ICBDescriptor) {
	objc.Send[objc.ID](c.ID, objc.Sel("readValueForDescriptor:"), descriptor)
}

// Writes the value of a characteristic.
//
// data: The value to write.
//
// characteristic: The characteristic containing the value to write.
//
// type: The type of write to execute. For a list of the possible types of writes to
// a characteristic’s value, see [CBCharacteristicWriteType].
//
// # Discussion
//
// When you call this method to write the value of a characteristic, the
// peripheral calls the [PeripheralDidWriteValueForDescriptorError] method of
// its delegate object only if you specified the write type as
// [CBCharacteristicWriteType.withResponse]. The response you receive through
// the [PeripheralDidWriteValueForDescriptorError] delegate method indicates
// whether the write was successful; if the write failed, it details the cause
// of the failure in an error.
//
// On the other hand, if you specify the write type as
// [CBCharacteristicWriteType.withoutResponse], Core Bluetooth attempts to
// write the value but doesn’t guarantee success. If the write doesn’t
// succeed in this case, you aren’t notified and you don’t receive an
// error indicating the cause of the failure.
//
// Use the [CBCharacteristicPropertyWrite] and
// [CBCharacteristicPropertyWriteWithoutResponse] members of the
// characteristic’s [CBCharacteristic.Properties] enumeration to determine
// which kinds of writes you can perform.
//
// This method copies the data passed into the `data` parameter, and you can
// dispose of it after the method returns.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/writeValue(_:for:type:)
//
// [CBCharacteristicWriteType]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType
// [CBCharacteristicWriteType.withResponse]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType/withResponse
// [CBCharacteristicWriteType.withoutResponse]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType/withoutResponse
func (c CBPeripheral) WriteValueForCharacteristicType(data foundation.NSData, characteristic ICBCharacteristic, type_ CBCharacteristicWriteType) {
	objc.Send[objc.ID](c.ID, objc.Sel("writeValue:forCharacteristic:type:"), data, characteristic, type_)
}

// Writes the value of a characteristic descriptor.
//
// data: The value to write.
//
// descriptor: The descriptor containing the value to write.
//
// # Discussion
//
// When you call this method to write the value of a characteristic
// descriptor, the peripheral calls the
// [PeripheralDidWriteValueForDescriptorError] method of its delegate object.
//
// This method copies the `data` passed into the data parameter, and you can
// dispose of it after the method returns.
//
// You can’t use this method to write the value of a client configuration
// descriptor (represented by the
// [CBUUIDClientCharacteristicConfigurationString] constant), which describes
// the configuration of notification or indications for a characteristic’s
// value. If you want to manage notifications or indications for a
// characteristic’s value, you must use the
// [CBPeripheral.SetNotifyValueForCharacteristic] method instead.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/writeValue(_:for:)
//
// [CBUUIDClientCharacteristicConfigurationString]: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDClientCharacteristicConfigurationString
func (c CBPeripheral) WriteValueForDescriptor(data foundation.NSData, descriptor ICBDescriptor) {
	objc.Send[objc.ID](c.ID, objc.Sel("writeValue:forDescriptor:"), data, descriptor)
}

// The maximum amount of data, in bytes, you can send to a characteristic in a
// single write type.
//
// type: The characteristic write type to inspect.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/maximumWriteValueLength(for:)
func (c CBPeripheral) MaximumWriteValueLengthForType(type_ CBCharacteristicWriteType) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maximumWriteValueLengthForType:"), type_)
	return rv
}

// Sets notifications or indications for the value of a specified
// characteristic.
//
// enabled: A Boolean value that indicates whether to receive notifications or
// indications whenever the characteristic’s value changes. true if you want
// to enable notifications or indications for the characteristic’s value.
// false if you don’t want to receive notifications or indications whenever
// the characteristic’s value changes.
//
// characteristic: The specified characteristic.
//
// # Discussion
//
// When you enable notifications for the characteristic’s value, the
// peripheral calls the
// [PeripheralDidUpdateNotificationStateForCharacteristicError] method of its
// delegate object to indicate if the action succeeded. If successful, the
// peripheral then calls the [PeripheralDidUpdateValueForDescriptorError]
// method of its delegate object whenever the characteristic value changes.
// Because the peripheral chooses when it sends an update, your app should
// prepare to handle them as long as notifications or indications remain
// enabled. If the specified characteristic’s configuration allows both
// notifications and indications, calling this method enables notifications
// only. You can disable notifications and indications for a
// characteristic’s value by calling this method with the `enabled`
// parameter set to false.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/setNotifyValue(_:for:)
func (c CBPeripheral) SetNotifyValueForCharacteristic(enabled bool, characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c.ID, objc.Sel("setNotifyValue:forCharacteristic:"), enabled, characteristic)
}

// Retrieves the current RSSI value for the peripheral while connected to the
// central manager.
//
// # Discussion
//
// On macOS, when you call this method to retrieve the Received Signal
// Strength Indicator (RSSI) of the peripheral while connected to the central
// manager, the peripheral calls the [peripheralDidUpdateRSSI(_:error:)]
// method of its delegate object. If retrieving the RSSI value of the
// peripheral succeeds, you can access it through the peripheral’s [rssi]
// property.
//
// On iOS and tvOS, when you call this method to retrieve the RSSI of the
// peripheral while connected to the central manager, the peripheral calls the
// [PeripheralDidReadRSSIError] method of its delegate object, which includes
// the RSSI value as a parameter.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readRSSI()
//
// [peripheralDidUpdateRSSI(_:error:)]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheralDidUpdateRSSI(_:error:)
// [rssi]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/rssi
func (c CBPeripheral) ReadRSSI() {
	objc.Send[objc.ID](c.ID, objc.Sel("readRSSI"))
}

// Attempts to open an L2CAP channel to the peripheral using the supplied
// Protocol/Service Multiplexer (PSM).
//
// PSM: The PSM of the channel to open.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/openL2CAPChannel(_:)
func (c CBPeripheral) OpenL2CAPChannel(PSM CBL2CAPPSM) {
	objc.Send[objc.ID](c.ID, objc.Sel("openL2CAPChannel:"), PSM)
}

// The name of the peripheral.
//
// # Discussion
//
// Use this property to retrieve a human-readable name of the peripheral. A
// peripheral may have two different name types: one that the device
// advertises and another that the device publishes in its database as its
// Bluetooth low energy Generic Access Profile (GAP) device name. If a
// peripheral has both types of names, this property returns its GAP device
// name.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/name
func (c CBPeripheral) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The delegate object specified to receive peripheral events.
//
// # Discussion
//
// For information about how to implement your peripheral delegate, see
// [CBPeripheralDelegate].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/delegate
func (c CBPeripheral) Delegate() CBPeripheralDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return CBPeripheralDelegateObjectFromID(rv)
}
func (c CBPeripheral) SetDelegate(value CBPeripheralDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}

// A list of a peripheral’s discovered services.
//
// # Discussion
//
// Returns an array of services (represented by [CBService] objects) that
// successful call to the [CBPeripheral.DiscoverServices] method discovered.
// If you haven’t yet called the [CBPeripheral.DiscoverServices] method to
// discover the services of the peripheral, or if there was an error in doing
// so, the value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/services
func (c CBPeripheral) Services() []CBService {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("services"))
	return objc.ConvertSlice(rv, func(id objc.ID) CBService {
		return CBServiceFromID(id)
	})
}

// The connection state of the peripheral.
//
// # Discussion
//
// This property represents the current connection state of the peripheral.
// For a list of the possible values, see [CBPeripheralState].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/state
//
// [CBPeripheralState]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralState
func (c CBPeripheral) State() CBPeripheralState {
	rv := objc.Send[CBPeripheralState](c.ID, objc.Sel("state"))
	return CBPeripheralState(rv)
}

// A Boolean value that indicates whether the remote device can send a write
// without a response.
//
// # Discussion
//
// If this value is false, flushing all current writes sets the value to true.
// This also results in a call to the delegate’s
// [PeripheralIsReadyToSendWriteWithoutResponse].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/canSendWriteWithoutResponse
func (c CBPeripheral) CanSendWriteWithoutResponse() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canSendWriteWithoutResponse"))
	return rv
}
