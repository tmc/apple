// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that provides updates for the discovery and management of peripheral devices.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate
type CBCentralManagerDelegate interface {
	objectivec.IObject

	// Tells the delegate the central manager’s state updated.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManagerDidUpdateState(_:)
	CentralManagerDidUpdateState(central ICBCentralManager)
}

// CBCentralManagerDelegateObject wraps an existing Objective-C object that conforms to the CBCentralManagerDelegate protocol.
type CBCentralManagerDelegateObject struct {
	objectivec.Object
}

func (o CBCentralManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CBCentralManagerDelegateObjectFromID constructs a [CBCentralManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CBCentralManagerDelegateObjectFromID(id objc.ID) CBCentralManagerDelegateObject {
	return CBCentralManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the central manager’s state updated.
//
// central: The central manager whose state has changed.
//
// # Discussion
//
// You implement this required method to ensure that the central device
// supports Bluetooth low energy and that it’s available to use. You should
// issue commands to the central manager only when the central manager’s
// [CBManager.State] indicates it’s powered on. A state with a value lower
// than [CBManagerState.poweredOn] implies that scanning has stopped, which in
// turn disconnects any previously-connected peripherals. If the state moves
// below [CBManagerState.poweredOff], all [CBPeripheral] objects obtained from
// this central manager become invalid; you must retrieve or discover these
// peripherals again. For a complete list of possible states, see
// [CBManagerState].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManagerDidUpdateState(_:)
//
// [CBManagerState.poweredOff]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/poweredOff
// [CBManagerState.poweredOn]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/poweredOn
// [CBManagerState]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState
func (o CBCentralManagerDelegateObject) CentralManagerDidUpdateState(central ICBCentralManager) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManagerDidUpdateState:"), central)
}

// Tells the delegate that the central manager connected to a peripheral.
//
// central: The central manager that provides this information.
//
// peripheral: The now-connected peripheral.
//
// # Discussion
//
// The manager invokes this method when a call to
// [CBCentralManager.ConnectPeripheralOptions] succeeds. You typically
// implement this method to set the peripheral’s delegate and discover its
// services.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didConnect:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (o CBCentralManagerDelegateObject) CentralManagerDidConnectPeripheral(central ICBCentralManager, peripheral ICBPeripheral) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didConnectPeripheral:"), central, peripheral)
}

// Tells the delegate that the central manager disconnected from a peripheral.
//
// central: The central manager that provides this information.
//
// peripheral: The now-disconnected peripheral.
//
// error: The cause of the failure, or `nil` if no error occurred.
//
// # Discussion
//
// The manager invokes this method when disconnecting a peripheral previously
// connected with the [CBCentralManager.ConnectPeripheralOptions] method. The
// error parameter contains the reason for the disconnection, unless the
// disconnect resulted from a call to
// [CBCentralManager.CancelPeripheralConnection]. After this method executes,
// the peripheral device’s [CBPeripheralDelegate] object receives no further
// method calls.
//
// All services, characteristics, and characteristic descriptors a peripheral
// become invalidated after it disconnects.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didDisconnectPeripheral:error:)
func (o CBCentralManagerDelegateObject) CentralManagerDidDisconnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didDisconnectPeripheral:error:"), central, peripheral, error_)
}

// Tells the delegate the central manager failed to create a connection with a
// peripheral.
//
// central: The central manager that provides this information.
//
// peripheral: The peripheral that failed to connect.
//
// error: The cause of the failure, or `nil` if no error occurred.
//
// # Discussion
//
// The manager invokes this method when a connection initiated with the
// [CBCentralManager.ConnectPeripheralOptions] method fails to complete.
// Because connection attempts don’t time out, a failed connection usually
// indicates a transient issue, in which case you may attempt connecting to
// the peripheral again.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didFailToConnect:error:)
func (o CBCentralManagerDelegateObject) CentralManagerDidFailToConnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didFailToConnectPeripheral:error:"), central, peripheral, error_)
}

// Tells the delegate that a connection event occurred which matches the
// registered options.
//
// # Discussion
//
// The manager calls this method when it observes a connection event that
// matches the options provided to
// [CBCentralManager.RegisterForConnectionEventsWithOptions].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:connectionEventDidOccur:for:)
func (o CBCentralManagerDelegateObject) CentralManagerConnectionEventDidOccurForPeripheral(central ICBCentralManager, event CBConnectionEvent, peripheral ICBPeripheral) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:connectionEventDidOccur:forPeripheral:"), central, event, peripheral)
}

// Tells the delegate the central manager discovered a peripheral while
// scanning for devices.
//
// central: The central manager that provides the update.
//
// peripheral: The discovered peripheral.
//
// advertisementData: A dictionary containing any advertisement data.
//
// RSSI: The current received signal strength indicator (RSSI) of the peripheral, in
// decibels.
//
// # Discussion
//
// You can access the advertisement data with the keys listed in
// [Advertisement Data Retrieval Keys]. You must retain a local copy of the
// peripheral if you want to perform commands on it. Use the RSSI data to
// determine the proximity of a discoverable peripheral device, and whether
// you want to connect to it automatically.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didDiscover:advertisementData:rssi:)
//
// [Advertisement Data Retrieval Keys]: https://developer.apple.com/documentation/CoreBluetooth/advertisement-data-retrieval-keys
func (o CBCentralManagerDelegateObject) CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI(central ICBCentralManager, peripheral ICBPeripheral, advertisementData foundation.INSDictionary, RSSI foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didDiscoverPeripheral:advertisementData:RSSI:"), central, peripheral, advertisementData, RSSI)
}

// Tells the delegate the system is about to restore the central manager, as
// part of relaunching the app into the background.
//
// central: The central manager that provides this information.
//
// dict: A dictionary that contains information about the central manager preserved
// by the system when it terminated the app. For the available keys to this
// dictionary, see [Central Manager State Restoration Options].
//
// # Discussion
//
// This method only applies to apps that opt in to the state preservation and
// restoration feature of Core Bluetooth. The system invokes this method when
// relaunching your app into the background to complete some Bluetooth-related
// task. Use this method to synchronize the state of your app with the state
// of the Bluetooth system.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:willRestoreState:)
//
// [Central Manager State Restoration Options]: https://developer.apple.com/documentation/CoreBluetooth/central-manager-state-restoration-options
func (o CBCentralManagerDelegateObject) CentralManagerWillRestoreState(central ICBCentralManager, dict foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:willRestoreState:"), central, dict)
}

// Tells the delegate the authorization status changed for a ANCS-requiring
// connected peripheral.
//
// central: The central manager providing this information.
//
// peripheral: The [CBPeripheral] that caused the event.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didUpdateANCSAuthorizationFor:)
func (o CBCentralManagerDelegateObject) CentralManagerDidUpdateANCSAuthorizationForPeripheral(central ICBCentralManager, peripheral ICBPeripheral) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didUpdateANCSAuthorizationForPeripheral:"), central, peripheral)
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerDelegate/centralManager(_:didDisconnectPeripheral:timestamp:isReconnecting:error:)
func (o CBCentralManagerDelegateObject) CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError(central ICBCentralManager, peripheral ICBPeripheral, timestamp corefoundation.CFAbsoluteTime, isReconnecting bool, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("centralManager:didDisconnectPeripheral:timestamp:isReconnecting:error:"), central, peripheral, timestamp, isReconnecting, error_)
}

// CBCentralManagerDelegateConfig holds optional typed callbacks for [CBCentralManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbcentralmanagerdelegate
type CBCentralManagerDelegateConfig struct {

	// Monitoring Connections with Peripherals
	// CentralManagerDidDisconnectPeripheralError — Tells the delegate that the central manager disconnected from a peripheral.
	CentralManagerDidDisconnectPeripheralError func(central CBCentralManager, peripheral CBPeripheral, error_ foundation.NSError)

	// Monitoring the Central Manager’s State
	// CentralManagerDidUpdateState — Tells the delegate the central manager’s state updated.
	CentralManagerDidUpdateState func(central CBCentralManager)
	// CentralManagerWillRestoreState — Tells the delegate the system is about to restore the central manager, as part of relaunching the app into the background.
	CentralManagerWillRestoreState func(central CBCentralManager, dict foundation.INSDictionary)

	// Other Methods
	// CentralManagerDidConnectPeripheral — Tells the delegate that the central manager connected to a peripheral.
	CentralManagerDidConnectPeripheral func(central CBCentralManager, peripheral CBPeripheral)
	// CentralManagerDidFailToConnectPeripheralError — Tells the delegate the central manager failed to create a connection with a peripheral.
	CentralManagerDidFailToConnectPeripheralError func(central CBCentralManager, peripheral CBPeripheral, error_ foundation.NSError)
	// CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI — Tells the delegate the central manager discovered a peripheral while scanning for devices.
	CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI func(central CBCentralManager, peripheral CBPeripheral, advertisementData foundation.INSDictionary, RSSI foundation.NSNumber)
}

// NewCBCentralManagerDelegate creates an Objective-C object implementing the [CBCentralManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CBCentralManagerDelegateObject] satisfies the [CBCentralManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbcentralmanagerdelegate
func NewCBCentralManagerDelegate(config CBCentralManagerDelegateConfig) CBCentralManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCBCentralManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CentralManagerDidUpdateState != nil {
		fn := config.CentralManagerDidUpdateState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManagerDidUpdateState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManagerDidUpdateState:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				fn(central)
				_delegateDone = true
			},
		})
	}

	if config.CentralManagerDidConnectPeripheral != nil {
		fn := config.CentralManagerDidConnectPeripheral
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManager:didConnectPeripheral:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID, peripheralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManager:didConnectPeripheral:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				peripheral := CBPeripheralFromID(peripheralID)
				fn(central, peripheral)
				_delegateDone = true
			},
		})
	}

	if config.CentralManagerDidDisconnectPeripheralError != nil {
		fn := config.CentralManagerDidDisconnectPeripheralError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManager:didDisconnectPeripheral:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID, peripheralID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManager:didDisconnectPeripheral:error:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				peripheral := CBPeripheralFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(central, peripheral, error_)
				_delegateDone = true
			},
		})
	}

	if config.CentralManagerDidFailToConnectPeripheralError != nil {
		fn := config.CentralManagerDidFailToConnectPeripheralError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManager:didFailToConnectPeripheral:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID, peripheralID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManager:didFailToConnectPeripheral:error:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				peripheral := CBPeripheralFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(central, peripheral, error_)
				_delegateDone = true
			},
		})
	}

	if config.CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI != nil {
		fn := config.CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManager:didDiscoverPeripheral:advertisementData:RSSI:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID, peripheralID objc.ID, advertisementDataID objc.ID, RSSIID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManager:didDiscoverPeripheral:advertisementData:RSSI:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				peripheral := CBPeripheralFromID(peripheralID)
				advertisementData := foundation.NSDictionaryFromID(advertisementDataID)
				RSSI := foundation.NSNumberFromID(RSSIID)
				fn(central, peripheral, advertisementData, RSSI)
				_delegateDone = true
			},
		})
	}

	if config.CentralManagerWillRestoreState != nil {
		fn := config.CentralManagerWillRestoreState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("centralManager:willRestoreState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centralID objc.ID, dictID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBCentralManagerDelegate", "centralManager:willRestoreState:")
					}
				}()
				central := CBCentralManagerFromID(centralID)
				dict := foundation.NSDictionaryFromID(dictID)
				fn(central, dict)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CBCentralManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCBCentralManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CBCentralManagerDelegateObjectFromID(instance)
}
