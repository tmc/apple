// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBCentralManager] class.
var (
	_CBCentralManagerClass     CBCentralManagerClass
	_CBCentralManagerClassOnce sync.Once
)

func getCBCentralManagerClass() CBCentralManagerClass {
	_CBCentralManagerClassOnce.Do(func() {
		_CBCentralManagerClass = CBCentralManagerClass{class: objc.GetClass("CBCentralManager")}
	})
	return _CBCentralManagerClass
}

// GetCBCentralManagerClass returns the class object for CBCentralManager.
func GetCBCentralManagerClass() CBCentralManagerClass {
	return getCBCentralManagerClass()
}

type CBCentralManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBCentralManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBCentralManagerClass) Alloc() CBCentralManager {
	rv := objc.Send[CBCentralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that scans for, discovers, connects to, and manages peripherals.
//
// # Overview
//
// [CBCentralManager] objects manage discovered or connected remote peripheral
// devices (represented by [CBPeripheral] objects), including scanning for,
// discovering, and connecting to advertising peripherals.
//
// Before calling the [CBCentralManager] methods, set the state of the central
// manager object to powered on, as indicated by the
// [CBCentralManagerState.poweredOn] constant. This state indicates that the
// central device (your iPhone or iPad, for instance) supports Bluetooth low
// energy and that Bluetooth is on and available for use.
//
// # Initializing a Central Manager
//
//   - [CBCentralManager.InitWithDelegateQueue]: Initializes the central manager with a specified delegate and dispatch queue.
//   - [CBCentralManager.InitWithDelegateQueueOptions]: Initializes the central manager with specified delegate, dispatch queue, and initialization options.
//
// # Establishing or Canceling Connections with Peripherals
//
//   - [CBCentralManager.ConnectPeripheralOptions]: Establishes a local connection to a peripheral.
//   - [CBCentralManager.CancelPeripheralConnection]: Cancels an active or pending local connection to a peripheral.
//
// # Retrieving Lists of Peripherals
//
//   - [CBCentralManager.RetrieveConnectedPeripheralsWithServices]: Returns a list of the peripherals connected to the system whose services match a given set of criteria.
//   - [CBCentralManager.RetrievePeripheralsWithIdentifiers]: Returns a list of known peripherals by their identifiers.
//
// # Scanning or Stopping Scans of Peripherals
//
//   - [CBCentralManager.ScanForPeripheralsWithServicesOptions]: Scans for peripherals that are advertising services.
//   - [CBCentralManager.StopScan]: Asks the central manager to stop scanning for peripherals.
//   - [CBCentralManager.IsScanning]: A Boolean value that indicates whether the central is currently scanning.
//
// # Monitoring Properties
//
//   - [CBCentralManager.Delegate]: The delegate object that you want to receive central manager events.
//   - [CBCentralManager.SetDelegate]
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager
//
// [CBCentralManagerState.poweredOn]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/poweredOn
type CBCentralManager struct {
	CBManager
}

// CBCentralManagerFromID constructs a [CBCentralManager] from an objc.ID.
//
// An object that scans for, discovers, connects to, and manages peripherals.
func CBCentralManagerFromID(id objc.ID) CBCentralManager {
	return CBCentralManager{CBManager: CBManagerFromID(id)}
}

// NOTE: CBCentralManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBCentralManager] class.
//
// # Initializing a Central Manager
//
//   - [ICBCentralManager.InitWithDelegateQueue]: Initializes the central manager with a specified delegate and dispatch queue.
//   - [ICBCentralManager.InitWithDelegateQueueOptions]: Initializes the central manager with specified delegate, dispatch queue, and initialization options.
//
// # Establishing or Canceling Connections with Peripherals
//
//   - [ICBCentralManager.ConnectPeripheralOptions]: Establishes a local connection to a peripheral.
//   - [ICBCentralManager.CancelPeripheralConnection]: Cancels an active or pending local connection to a peripheral.
//
// # Retrieving Lists of Peripherals
//
//   - [ICBCentralManager.RetrieveConnectedPeripheralsWithServices]: Returns a list of the peripherals connected to the system whose services match a given set of criteria.
//   - [ICBCentralManager.RetrievePeripheralsWithIdentifiers]: Returns a list of known peripherals by their identifiers.
//
// # Scanning or Stopping Scans of Peripherals
//
//   - [ICBCentralManager.ScanForPeripheralsWithServicesOptions]: Scans for peripherals that are advertising services.
//   - [ICBCentralManager.StopScan]: Asks the central manager to stop scanning for peripherals.
//   - [ICBCentralManager.IsScanning]: A Boolean value that indicates whether the central is currently scanning.
//
// # Monitoring Properties
//
//   - [ICBCentralManager.Delegate]: The delegate object that you want to receive central manager events.
//   - [ICBCentralManager.SetDelegate]
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager
type ICBCentralManager interface {
	ICBManager

	// Topic: Initializing a Central Manager

	// Initializes the central manager with a specified delegate and dispatch queue.
	InitWithDelegateQueue(delegate CBCentralManagerDelegate, queue dispatch.Queue) CBCentralManager
	// Initializes the central manager with specified delegate, dispatch queue, and initialization options.
	InitWithDelegateQueueOptions(delegate CBCentralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBCentralManager

	// Topic: Establishing or Canceling Connections with Peripherals

	// Establishes a local connection to a peripheral.
	ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.INSDictionary)
	// Cancels an active or pending local connection to a peripheral.
	CancelPeripheralConnection(peripheral ICBPeripheral)

	// Topic: Retrieving Lists of Peripherals

	// Returns a list of the peripherals connected to the system whose services match a given set of criteria.
	RetrieveConnectedPeripheralsWithServices(serviceUUIDs []CBUUID) []CBPeripheral
	// Returns a list of known peripherals by their identifiers.
	RetrievePeripheralsWithIdentifiers(identifiers []foundation.NSUUID) []CBPeripheral

	// Topic: Scanning or Stopping Scans of Peripherals

	// Scans for peripherals that are advertising services.
	ScanForPeripheralsWithServicesOptions(serviceUUIDs []CBUUID, options foundation.INSDictionary)
	// Asks the central manager to stop scanning for peripherals.
	StopScan()
	// A Boolean value that indicates whether the central is currently scanning.
	IsScanning() bool

	// Topic: Monitoring Properties

	// The delegate object that you want to receive central manager events.
	Delegate() CBCentralManagerDelegate
	SetDelegate(value CBCentralManagerDelegate)
}

// Init initializes the instance.
func (c CBCentralManager) Init() CBCentralManager {
	rv := objc.Send[CBCentralManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBCentralManager) Autorelease() CBCentralManager {
	rv := objc.Send[CBCentralManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCentralManager creates a new CBCentralManager instance.
func NewCBCentralManager() CBCentralManager {
	class := getCBCentralManagerClass()
	rv := objc.Send[CBCentralManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the central manager with a specified delegate and dispatch
// queue.
//
// delegate: The delegate that receives central events.
//
// queue: The dispatch queue used to dispatch the central role events. If the value
// is `nil`, the central manager dispatches central role events using the main
// queue.
//
// # Return Value
//
// Returns a newly initialized central manager.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func NewCBCentralManagerWithDelegateQueue(delegate CBCentralManagerDelegate, queue dispatch.Queue) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, uintptr(queue.Handle()))
	return CBCentralManagerFromID(rv)
}

// Initializes the central manager with specified delegate, dispatch queue,
// and initialization options.
//
// delegate: The delegate that receives the central events.
//
// queue: The dispatch queue used to dispatch the central role events. If the value
// is `nil`, the central manager dispatches central role events using the main
// queue.
//
// options: An optional dictionary that contains initialization options for a central
// manager. For available options, see [Central Manager Initialization
// Options].
//
// # Return Value
//
// Returns a newly initialized central manager.
//
// # Discussion
//
// This method is the designated initializer for the [CBCentralManager] class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:options:)
//
// [Central Manager Initialization Options]: https://developer.apple.com/documentation/CoreBluetooth/central-manager-initialization-options
func NewCBCentralManagerWithDelegateQueueOptions(delegate CBCentralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, uintptr(queue.Handle()), options)
	return CBCentralManagerFromID(rv)
}

// Initializes the central manager with a specified delegate and dispatch
// queue.
//
// delegate: The delegate that receives central events.
//
// queue: The dispatch queue used to dispatch the central role events. If the value
// is `nil`, the central manager dispatches central role events using the main
// queue.
//
// # Return Value
//
// Returns a newly initialized central manager.
//
// # Discussion
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (c CBCentralManager) InitWithDelegateQueue(delegate CBCentralManagerDelegate, queue dispatch.Queue) CBCentralManager {
	rv := objc.Send[CBCentralManager](c.ID, objc.Sel("initWithDelegate:queue:"), delegate, uintptr(queue.Handle()))
	return rv
}

// Initializes the central manager with specified delegate, dispatch queue,
// and initialization options.
//
// delegate: The delegate that receives the central events.
//
// queue: The dispatch queue used to dispatch the central role events. If the value
// is `nil`, the central manager dispatches central role events using the main
// queue.
//
// options: An optional dictionary that contains initialization options for a central
// manager. For available options, see [Central Manager Initialization
// Options].
//
// # Return Value
//
// Returns a newly initialized central manager.
//
// # Discussion
//
// This method is the designated initializer for the [CBCentralManager] class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:options:)
//
// [Central Manager Initialization Options]: https://developer.apple.com/documentation/CoreBluetooth/central-manager-initialization-options
func (c CBCentralManager) InitWithDelegateQueueOptions(delegate CBCentralManagerDelegate, queue dispatch.Queue, options foundation.INSDictionary) CBCentralManager {
	rv := objc.Send[CBCentralManager](c.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, uintptr(queue.Handle()), options)
	return rv
}

// Establishes a local connection to a peripheral.
//
// peripheral: The peripheral to which the central is attempting to connect.
//
// options: A dictionary to customize the behavior of the connection. For available
// options, see [Peripheral Connection Options].
//
// # Discussion
//
// After successfully establishing a local connection to a peripheral, the
// central manager object calls the [CentralManagerDidConnectPeripheral]
// method of its delegate object. If the connection attempt fails, the central
// manager object calls the [CentralManagerDidFailToConnectPeripheralError]
// method of its delegate object instead. Attempts to connect to a peripheral
// don’t time out. To explicitly cancel a pending connection to a
// peripheral, call the [CBCentralManager.CancelPeripheralConnection] method.
// Deallocating `peripheral` also implicitly calls
// [CBCentralManager.CancelPeripheralConnection].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/connect(_:options:)
//
// [Peripheral Connection Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-connection-options
func (c CBCentralManager) ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.INSDictionary) {
	objc.Send[objc.ID](c.ID, objc.Sel("connectPeripheral:options:"), peripheral, options)
}

// Cancels an active or pending local connection to a peripheral.
//
// peripheral: The peripheral to which the central manager is either trying to connect or
// has already connected.
//
// # Discussion
//
// This method is nonblocking, and any [CBPeripheral] class commands that are
// still pending to `peripheral` may not complete. Because other apps may
// still have a connection to the peripheral, canceling a local connection
// doesn’t guarantee that the underlying physical link is immediately
// disconnected. From the app’s perspective, however, the peripheral is
// effectively disconnected, and the central manager object calls the
// [CentralManagerDidDisconnectPeripheralError] method of its delegate object.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/cancelPeripheralConnection(_:)
func (c CBCentralManager) CancelPeripheralConnection(peripheral ICBPeripheral) {
	objc.Send[objc.ID](c.ID, objc.Sel("cancelPeripheralConnection:"), peripheral)
}

// Returns a list of the peripherals connected to the system whose services
// match a given set of criteria.
//
// serviceUUIDs: A list of service UUIDs, represented by [CBUUID] objects.
//
// # Return Value
//
// A list of the peripherals that are currently connected to the system and
// that contain any of the services specified in the `serviceUUID` parameter.
//
// # Discussion
//
// The list of connected peripherals can include those that other apps have
// connected. You need to connect these peripherals locally using the
// [CBCentralManager.ConnectPeripheralOptions] method before using them.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrieveConnectedPeripherals(withServices:)
func (c CBCentralManager) RetrieveConnectedPeripheralsWithServices(serviceUUIDs []CBUUID) []CBPeripheral {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("retrieveConnectedPeripheralsWithServices:"), objectivec.IObjectSliceToNSArray(serviceUUIDs))
	return objc.ConvertSlice(rv, func(id objc.ID) CBPeripheral {
		return CBPeripheralFromID(id)
	})
}

// Returns a list of known peripherals by their identifiers.
//
// identifiers: A list of peripheral identifiers (represented by [NSUUID] objects) from
// which [CBPeripheral] objects can be retrieved.
//
// # Return Value
//
// A list of peripherals that the central manager is able to match to the
// provided identifiers.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrievePeripherals(withIdentifiers:)
//
// [NSUUID]: https://developer.apple.com/documentation/Foundation/NSUUID
func (c CBCentralManager) RetrievePeripheralsWithIdentifiers(identifiers []foundation.NSUUID) []CBPeripheral {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("retrievePeripheralsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(identifiers))
	return objc.ConvertSlice(rv, func(id objc.ID) CBPeripheral {
		return CBPeripheralFromID(id)
	})
}

// Scans for peripherals that are advertising services.
//
// serviceUUIDs: An array of [CBUUID] objects that the app is interested in. Each [CBUUID]
// object represents the UUID of a service that a peripheral advertises.
//
// options: A dictionary of options for customizing the scan. For available options,
// see [Peripheral Scanning Options].
//
// # Discussion
//
// You can provide an array of [CBUUID] objects — representing service UUIDs
// — in the `serviceUUIDs` parameter. When you do, the central manager
// returns only peripherals that advertise the services you specify. If the
// `serviceUUIDs` parameter is `nil`, this method returns all discovered
// peripherals, regardless of their supported services.
//
// If the central manager is actively scanning with one set of parameters and
// it receives another set to scan, the new parameters override the previous
// set. When the central manager discovers a peripheral, it calls the
// [CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI] method of its
// delegate object.
//
// Your app can scan for Bluetooth devices in the background by specifying the
// `bluetooth-central` background mode. To do this, your app must explicitly
// scan for one or more services by specifying them in the `serviceUUIDs`
// parameter. The [CBCentralManager] scan option has no effect while scanning
// in the background.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/scanForPeripherals(withServices:options:)
//
// [Peripheral Scanning Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-scanning-options
func (c CBCentralManager) ScanForPeripheralsWithServicesOptions(serviceUUIDs []CBUUID, options foundation.INSDictionary) {
	objc.Send[objc.ID](c.ID, objc.Sel("scanForPeripheralsWithServices:options:"), objectivec.IObjectSliceToNSArray(serviceUUIDs), options)
}

// Asks the central manager to stop scanning for peripherals.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/stopScan()
func (c CBCentralManager) StopScan() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopScan"))
}

// A Boolean value that indicates whether the central is currently scanning.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/isScanning
func (c CBCentralManager) IsScanning() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isScanning"))
	return rv
}

// The delegate object that you want to receive central manager events.
//
// # Discussion
//
// For information about how to implement your central manager delegate, see
// [CBCentralManagerDelegate].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/delegate
func (c CBCentralManager) Delegate() CBCentralManagerDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return CBCentralManagerDelegateObjectFromID(rv)
}
func (c CBCentralManager) SetDelegate(value CBCentralManagerDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}
