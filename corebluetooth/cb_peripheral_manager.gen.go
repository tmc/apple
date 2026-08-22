// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBPeripheralManager] class.
var (
	_CBPeripheralManagerClass     CBPeripheralManagerClass
	_CBPeripheralManagerClassOnce sync.Once
)

func getCBPeripheralManagerClass() CBPeripheralManagerClass {
	_CBPeripheralManagerClassOnce.Do(func() {
		_CBPeripheralManagerClass = CBPeripheralManagerClass{class: objc.GetClass("CBPeripheralManager")}
	})
	return _CBPeripheralManagerClass
}

// GetCBPeripheralManagerClass returns the class object for CBPeripheralManager.
func GetCBPeripheralManagerClass() CBPeripheralManagerClass {
	return getCBPeripheralManagerClass()
}

type CBPeripheralManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBPeripheralManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBPeripheralManagerClass) Alloc() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages and advertises peripheral services exposed by this
// app.
//
// # Overview
//
// Core Bluetooth uses [CBPeripheralManager] objects to manage published
// services within the local peripheral’s Generic Attribute Profile (GATT)
// database and to advertise these services to central devices (represented by
// [CBCentral] objects). While a service is in the database, any connected
// central can see and connect to it. That said, if your app hasn’t
// specified the `bluetooth-peripheral` background mode, the contents of its
// services become disabled when it’s in the background or in a suspended
// state. In this scenario, any remote central trying to access the
// service’s characteristic value or characteristic descriptors receives an
// error.
//
// Before you call [CBPeripheralManager] methods, the peripheral manager
// object must be in the powered-on state, as indicated by the
// [CBPeripheralManagerState.poweredOn]. This state indicates that the device
// (your iPhone or iPad, for instance) supports Bluetooth low energy and that
// its Bluetooth is on and available for use.
//
// In watchOS, tvOS, and visionOS, you can’t advertise services using a
// [CBPeripheralManager] object because support for doing so is unavailable.
//
// # Initializing a Peripheral Manager
//
//   - [CBPeripheralManager.InitWithDelegateQueue]: Initializes the peripheral manager with a specified delegate and dispatch queue.
//   - [CBPeripheralManager.InitWithDelegateQueueOptions]: Initializes the peripheral manager with a specified delegate, dispatch queue, and initialization options.
//   - [CBPeripheralManager.Delegate]: The delegate object specified to receive peripheral events.
//   - [CBPeripheralManager.SetDelegate]
//
// # Adding and Removing Services
//
//   - [CBPeripheralManager.AddService]: Publishes a service and any of its associated characteristics and characteristic descriptors to the local GATT database.
//   - [CBPeripheralManager.RemoveService]: Removes a specified published service from the local GATT database.
//   - [CBPeripheralManager.RemoveAllServices]: Removes all published services from the local GATT database.
//
// # Managing Advertising
//
//   - [CBPeripheralManager.StartAdvertising]: Advertises peripheral manager data.
//   - [CBPeripheralManager.StopAdvertising]: Stops advertising peripheral manager data.
//   - [CBPeripheralManager.IsAdvertising]: A Boolean value that indicates whether the peripheral is advertising data.
//
// # Sending Updates of a Characteristic’s Value
//
//   - [CBPeripheralManager.UpdateValueForCharacteristicOnSubscribedCentrals]: Send an updated characteristic value to one or more subscribed centrals, using a notification or indication.
//
// # Responding to Read and Write Requests
//
//   - [CBPeripheralManager.RespondToRequestWithResult]: Responds to a read or write request from a connected central.
//
// # Setting Connection Latency
//
//   - [CBPeripheralManager.SetDesiredConnectionLatencyForCentral]: Sets the desired connection latency for an existing connection to a central device.
//
// # Using L2CAP Channels
//
//   - [CBPeripheralManager.PublishL2CAPChannelWithEncryption]: Creates a listener for incoming L2CAP channel connections.
//   - [CBPeripheralManager.UnpublishL2CAPChannel]: Removes a published service from the local system.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager
//
// [CBPeripheralManagerState.poweredOn]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/poweredOn
type CBPeripheralManager struct {
	CBManager
}

// CBPeripheralManagerFromID constructs a [CBPeripheralManager] from an objc.ID.
//
// An object that manages and advertises peripheral services exposed by this
// app.
func CBPeripheralManagerFromID(id objc.ID) CBPeripheralManager {
	return CBPeripheralManager{CBManager: CBManagerFromID(id)}
}

// NOTE: CBPeripheralManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBPeripheralManager] class.
//
// # Initializing a Peripheral Manager
//
//   - [ICBPeripheralManager.InitWithDelegateQueue]: Initializes the peripheral manager with a specified delegate and dispatch queue.
//   - [ICBPeripheralManager.InitWithDelegateQueueOptions]: Initializes the peripheral manager with a specified delegate, dispatch queue, and initialization options.
//   - [ICBPeripheralManager.Delegate]: The delegate object specified to receive peripheral events.
//   - [ICBPeripheralManager.SetDelegate]
//
// # Adding and Removing Services
//
//   - [ICBPeripheralManager.AddService]: Publishes a service and any of its associated characteristics and characteristic descriptors to the local GATT database.
//   - [ICBPeripheralManager.RemoveService]: Removes a specified published service from the local GATT database.
//   - [ICBPeripheralManager.RemoveAllServices]: Removes all published services from the local GATT database.
//
// # Managing Advertising
//
//   - [ICBPeripheralManager.StartAdvertising]: Advertises peripheral manager data.
//   - [ICBPeripheralManager.StopAdvertising]: Stops advertising peripheral manager data.
//   - [ICBPeripheralManager.IsAdvertising]: A Boolean value that indicates whether the peripheral is advertising data.
//
// # Sending Updates of a Characteristic’s Value
//
//   - [ICBPeripheralManager.UpdateValueForCharacteristicOnSubscribedCentrals]: Send an updated characteristic value to one or more subscribed centrals, using a notification or indication.
//
// # Responding to Read and Write Requests
//
//   - [ICBPeripheralManager.RespondToRequestWithResult]: Responds to a read or write request from a connected central.
//
// # Setting Connection Latency
//
//   - [ICBPeripheralManager.SetDesiredConnectionLatencyForCentral]: Sets the desired connection latency for an existing connection to a central device.
//
// # Using L2CAP Channels
//
//   - [ICBPeripheralManager.PublishL2CAPChannelWithEncryption]: Creates a listener for incoming L2CAP channel connections.
//   - [ICBPeripheralManager.UnpublishL2CAPChannel]: Removes a published service from the local system.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager
type ICBPeripheralManager interface {
	ICBManager

	// Topic: Initializing a Peripheral Manager

	// Initializes the peripheral manager with a specified delegate and dispatch queue.
	InitWithDelegateQueue(delegate CBPeripheralManagerDelegate, queue dispatch.Queue) CBPeripheralManager
	// Initializes the peripheral manager with a specified delegate, dispatch queue, and initialization options.
	InitWithDelegateQueueOptions(delegate CBPeripheralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBPeripheralManager
	// The delegate object specified to receive peripheral events.
	Delegate() CBPeripheralManagerDelegate
	SetDelegate(value CBPeripheralManagerDelegate)

	// Topic: Adding and Removing Services

	// Publishes a service and any of its associated characteristics and characteristic descriptors to the local GATT database.
	AddService(service ICBMutableService)
	// Removes a specified published service from the local GATT database.
	RemoveService(service ICBMutableService)
	// Removes all published services from the local GATT database.
	RemoveAllServices()

	// Topic: Managing Advertising

	// Advertises peripheral manager data.
	StartAdvertising(advertisementData foundation.INSDictionary)
	// Stops advertising peripheral manager data.
	StopAdvertising()
	// A Boolean value that indicates whether the peripheral is advertising data.
	IsAdvertising() bool

	// Topic: Sending Updates of a Characteristic’s Value

	// Send an updated characteristic value to one or more subscribed centrals, using a notification or indication.
	UpdateValueForCharacteristicOnSubscribedCentrals(value foundation.NSData, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool

	// Topic: Responding to Read and Write Requests

	// Responds to a read or write request from a connected central.
	RespondToRequestWithResult(request ICBATTRequest, result CBATTError)

	// Topic: Setting Connection Latency

	// Sets the desired connection latency for an existing connection to a central device.
	SetDesiredConnectionLatencyForCentral(latency CBPeripheralManagerConnectionLatency, central ICBCentral)

	// Topic: Using L2CAP Channels

	// Creates a listener for incoming L2CAP channel connections.
	PublishL2CAPChannelWithEncryption(encryptionRequired bool)
	// Removes a published service from the local system.
	UnpublishL2CAPChannel(PSM CBL2CAPPSM)
}

// Init initializes the instance.
func (c CBPeripheralManager) Init() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBPeripheralManager) Autorelease() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeripheralManager creates a new CBPeripheralManager instance.
func NewCBPeripheralManager() CBPeripheralManager {
	class := getCBPeripheralManagerClass()
	rv := objc.Send[CBPeripheralManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the peripheral manager with a specified delegate and dispatch
// queue.
//
// delegate: The delegate to receive the peripheral role events.
//
// queue: The dispatch queue for dispatching the peripheral role events. If the value
// is `nil`, the peripheral manager dispatches peripheral role events using
// the main queue.
//
// # Return Value
//
// Returns a newly initialized peripheral manager.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func NewCBPeripheralManagerWithDelegateQueue(delegate CBPeripheralManagerDelegate, queue dispatch.Queue) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, uintptr(queue.Handle()))
	return CBPeripheralManagerFromID(rv)
}

// Initializes the peripheral manager with a specified delegate, dispatch
// queue, and initialization options.
//
// delegate: The delegate to receive the peripheral role events.
//
// queue: The dispatch queue for dispatching the peripheral role events. If the value
// is `nil`, the peripheral manager dispatches peripheral role events using
// the main queue.
//
// options: An optional dictionary containing initialization options for a peripheral
// manager. For available options, see [Peripheral Manager Initialization
// Options].
//
// # Return Value
//
// Returns a newly initialized peripheral manager.
//
// # Discussion
//
// This method is the designated initializer for the [CBPeripheralManager]
// class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:options:)
//
// [Peripheral Manager Initialization Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-manager-initialization-options
func NewCBPeripheralManagerWithDelegateQueueOptions(delegate CBPeripheralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, uintptr(queue.Handle()), options)
	return CBPeripheralManagerFromID(rv)
}

// Initializes the peripheral manager with a specified delegate and dispatch
// queue.
//
// delegate: The delegate to receive the peripheral role events.
//
// queue: The dispatch queue for dispatching the peripheral role events. If the value
// is `nil`, the peripheral manager dispatches peripheral role events using
// the main queue.
//
// # Return Value
//
// Returns a newly initialized peripheral manager.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBPeripheralManager) InitWithDelegateQueue(delegate CBPeripheralManagerDelegate, queue dispatch.Queue) CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c.ID, objc.Sel("initWithDelegate:queue:"), delegate, uintptr(queue.Handle()))
	return rv
}

// Initializes the peripheral manager with a specified delegate, dispatch
// queue, and initialization options.
//
// delegate: The delegate to receive the peripheral role events.
//
// queue: The dispatch queue for dispatching the peripheral role events. If the value
// is `nil`, the peripheral manager dispatches peripheral role events using
// the main queue.
//
// options: An optional dictionary containing initialization options for a peripheral
// manager. For available options, see [Peripheral Manager Initialization
// Options].
//
// # Return Value
//
// Returns a newly initialized peripheral manager.
//
// # Discussion
//
// This method is the designated initializer for the [CBPeripheralManager]
// class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:options:)
//
// [Peripheral Manager Initialization Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-manager-initialization-options
func (c CBPeripheralManager) InitWithDelegateQueueOptions(delegate CBPeripheralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, uintptr(queue.Handle()), options)
	return rv
}

// Publishes a service and any of its associated characteristics and
// characteristic descriptors to the local GATT database.
//
// service: The service you want to publish.
//
// # Discussion
//
// When you add a service to the database, the peripheral manager calls the
// [PeripheralManagerDidAddServiceError] method of its delegate object. If the
// service contains any included services, you must first publish them.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/add(_:)
func (c CBPeripheralManager) AddService(service ICBMutableService) {
	objc.Send[objc.ID](c.ID, objc.Sel("addService:"), service)
}

// Removes a specified published service from the local GATT database.
//
// service: The service you want to remove.
//
// # Discussion
//
// Because apps on the local peripheral device share the GATT database, more
// than one instance of a service may exist in the database. As a result, this
// method removes only the instance of the service that your app added to the
// database (using the [CBPeripheralManager.AddService] method). If any other
// services contains this service, you must first remove them.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/remove(_:)
func (c CBPeripheralManager) RemoveService(service ICBMutableService) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeService:"), service)
}

// Removes all published services from the local GATT database.
//
// # Discussion
//
// Use this when you want to remove all services you’ve previously
// published, for example, if your app has a toggle button to expose GATT
// services.
//
// Because apps on the local peripheral device share the GATT database, this
// method removes only the services that you added using the
// [CBPeripheralManager.AddService] method. This call doesn’t remove any
// services published by other apps on the local peripheral device.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/removeAllServices()
func (c CBPeripheralManager) RemoveAllServices() {
	objc.Send[objc.ID](c.ID, objc.Sel("removeAllServices"))
}

// Advertises peripheral manager data.
//
// advertisementData: An optional dictionary containing the data you want to advertise. The
// peripheral manager only supports two keys:
// [CBAdvertisementDataLocalNameKey] and [CBAdvertisementDataServiceUUIDsKey].
//
// # Discussion
//
// When you start advertising peripheral data, the peripheral manager calls
// the [PeripheralManagerDidStartAdvertisingError] method of its delegate
// object.
//
// Core Bluetooth advertises data on a “best effort” basis, due to limited
// space and because there may be multiple apps advertising simultaneously.
// While in the foreground, your app can use up to 28 bytes of space in the
// initial advertisement data for any combination of the supported advertising
// data keys. If no this space remains, there’s an additional 10 bytes of
// space in the scan response, usable only for the local name (represented by
// the value of the [CBAdvertisementDataLocalNameKey] key). Note that these
// sizes don’t include the 2 bytes of header information required for each
// new data type.
//
// Any service UUIDs contained in the value of the
// [CBAdvertisementDataServiceUUIDsKey] key that don’t fit in the allotted
// space go to a special “overflow” area. These services are discoverable
// only by an iOS device explicitly scanning for them.
//
// While your app is in the background, the local name isn’t advertised and
// all service UUIDs are in the overflow area.
//
// For details about the format of advertising and response data, see the
// Bluetooth 4.0 specification, Volume 3, Part C, Section 11.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/startAdvertising(_:)
//
// [CBAdvertisementDataLocalNameKey]: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataLocalNameKey
// [CBAdvertisementDataServiceUUIDsKey]: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataServiceUUIDsKey
//
// [CBAdvertisementDataLocalNameKey]: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataLocalNameKey
// [CBAdvertisementDataServiceUUIDsKey]: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataServiceUUIDsKey
func (c CBPeripheralManager) StartAdvertising(advertisementData foundation.INSDictionary) {
	objc.Send[objc.ID](c.ID, objc.Sel("startAdvertising:"), advertisementData)
}

// Stops advertising peripheral manager data.
//
// # Discussion
//
// Call this method when you no longer want to advertise peripheral manager
// data.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/stopAdvertising()
func (c CBPeripheralManager) StopAdvertising() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopAdvertising"))
}

// Send an updated characteristic value to one or more subscribed centrals,
// using a notification or indication.
//
// value: The characteristic value you want to send via a notification or indication.
//
// characteristic: The characteristic whose value has changed.
//
// centrals: A list of centrals (represented by [CBCentral] objects) that have
// subscribed to receive updates of the characteristic’s value. If `nil`,
// the manager updates all subscribed centrals. The manager ignores any
// centrals that haven’t subscribed to the characteristic’s value.
//
// # Return Value
//
// This value is true if the update is successfully sent to the subscribed
// central or centrals. false if the update isn’t successfully sent because
// the underlying transmit queue is full.
//
// # Discussion
//
// You use this method to send updates of a characteristic’s value—through
// a notification or indication—to selected centrals that have subscribed to
// that characteristic’s value. If the method returns false because the
// underlying transmit queue is full, the peripheral manager calls the
// [PeripheralManagerIsReadyToUpdateSubscribers] method of its delegate object
// when more space in the transmit queue becomes available. After you receive
// this delegate method callback, you may resend the update.
//
// If the length of the `value` parameter exceeds the length of the
// [CBCentral.MaximumUpdateValueLength] property of a subscribed [CBCentral],
// the `value` parameter truncates accordingly.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/updateValue(_:for:onSubscribedCentrals:)
func (c CBPeripheralManager) UpdateValueForCharacteristicOnSubscribedCentrals(value foundation.NSData, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("updateValue:forCharacteristic:onSubscribedCentrals:"), value, characteristic, objectivec.IObjectSliceToNSArray(centrals))
	return rv
}

// Responds to a read or write request from a connected central.
//
// request: The read or write request received from the connected central. For more
// information about read and write requests, see [CBATTRequest].
//
// result: The result of attempting to fulfill the request. For a list of possible
// results, see [Deprecated Constants].
//
// # Discussion
//
// When the peripheral manager receives a request from a connected central to
// read or write a characteristic’s value, it calls the
// [PeripheralManagerDidReceiveReadRequest] or
// [PeripheralManagerDidReceiveWriteRequests] method of its delegate object.
// To respond to the corresponding read or write request, you call this method
// whenever you recevie one of these delegate method callbacks.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/respond(to:withResult:)
//
// [Deprecated Constants]: https://developer.apple.com/documentation/CoreBluetooth/deprecated-constants
func (c CBPeripheralManager) RespondToRequestWithResult(request ICBATTRequest, result CBATTError) {
	objc.Send[objc.ID](c.ID, objc.Sel("respondToRequest:withResult:"), request, result)
}

// Sets the desired connection latency for an existing connection to a central
// device.
//
// latency: The desired connection latency. For a list of the possible connection
// latency values that you may set for the peripheral manager, see
// [CBPeripheralManagerConnectionLatency].
//
// central: The central to which the peripheral manager is currently connected.
//
// # Discussion
//
// The latency of a peripheral-central connection controls how frequently the
// peripheral and the peripheral’s connected central can exchange messages.
// By setting a desired connection latency, you manage the relationship
// between the frequency of the data exchange and the resulting battery
// performance of the peripheral device. When you call this method to set the
// connection latency, note that connection latency changes aren’t
// guaranteed. As a result, the latency may vary. If you don’t explicitly
// set a latency, the central device uses the connection latency it chose when
// establishing the connection. Typically, you don’t need to change the
// connection latency.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/setDesiredConnectionLatency(_:for:)
//
// [CBPeripheralManagerConnectionLatency]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerConnectionLatency
func (c CBPeripheralManager) SetDesiredConnectionLatencyForCentral(latency CBPeripheralManagerConnectionLatency, central ICBCentral) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDesiredConnectionLatency:forCentral:"), latency, central)
}

// Creates a listener for incoming L2CAP channel connections.
//
// encryptionRequired: true if the service requires link encryption before a stream can be
// established. false if the service supports use over an unsecured link.
//
// # Discussion
//
// The system determines an unused Protocol and Service Multiplexer (PSM) at
// the time of publishing, and provides it to your app with
// [PeripheralManagerDidPublishL2CAPChannelError]. L2CAP channels aren’t
// discoverable by themselves, so it’s the app’s responsibility to handle
// PSM discovery on the client.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/publishL2CAPChannel(withEncryption:)
func (c CBPeripheralManager) PublishL2CAPChannelWithEncryption(encryptionRequired bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("publishL2CAPChannelWithEncryption:"), encryptionRequired)
}

// Removes a published service from the local system.
//
// PSM: The Protocol and Service Multiplexer (PSM) to remove from the system.
//
// # Discussion
//
// After you make this call, the peripheral manager accepts no new connections
// for this PSM, and closes any existing L2CAP channels using this PSM.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/unpublishL2CAPChannel(_:)
func (c CBPeripheralManager) UnpublishL2CAPChannel(PSM CBL2CAPPSM) {
	objc.Send[objc.ID](c.ID, objc.Sel("unpublishL2CAPChannel:"), PSM)
}

// The delegate object specified to receive peripheral events.
//
// # Discussion
//
// For information about how to implement your peripheral manager delegate,
// see [CBPeripheralManagerDelegate].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/delegate
func (c CBPeripheralManager) Delegate() CBPeripheralManagerDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return CBPeripheralManagerDelegateObjectFromID(rv)
}
func (c CBPeripheralManager) SetDelegate(value CBPeripheralManagerDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the peripheral is advertising data.
//
// # Discussion
//
// This value is true if the peripheral is advertising data as a result of
// successfully calling the [CBPeripheralManager.StartAdvertising] method. The
// value is false if the peripheral is no longer advertising its data.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/isAdvertising
func (c CBPeripheralManager) IsAdvertising() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAdvertising"))
	return rv
}
