// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that provides updates for local peripheral state and interactions with remote central devices.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate
type CBPeripheralManagerDelegate interface {
	objectivec.IObject

	// Tells the delegate the peripheral manager’s state updated.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManagerDidUpdateState(_:)
	PeripheralManagerDidUpdateState(peripheral ICBPeripheralManager)
}

// CBPeripheralManagerDelegateObject wraps an existing Objective-C object that conforms to the CBPeripheralManagerDelegate protocol.
type CBPeripheralManagerDelegateObject struct {
	objectivec.Object
}

func (o CBPeripheralManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CBPeripheralManagerDelegateObjectFromID constructs a [CBPeripheralManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CBPeripheralManagerDelegateObjectFromID(id objc.ID) CBPeripheralManagerDelegateObject {
	return CBPeripheralManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the peripheral manager’s state updated.
//
// peripheral: The peripheral manager whose state has changed.
//
// # Discussion
//
// You implement this required method to ensure that Bluetooth low energy is
// available to use on the local peripheral device.
//
// Issue commands to the peripheral manager only when the peripheral manager
// is in the powered-on state, as indicated by the
// [CBPeripheralManagerState.poweredOn] constant. A state with a value lower
// than [CBPeripheralManagerState.poweredOn] implies that advertising has
// stopped and that any connected centrals have been disconnected. If the
// state moves below [CBPeripheralManagerState.poweredOff], advertising has
// stopped you must explicitly restart it. In addition, the powered off state
// clears the local database; in this case you must explicitly re-add all
// services. For a complete list and discussion of the possible values
// representing the state of the peripheral manager, see the
// [CBPeripheralManagerState] enumeration in [CBPeripheralManager].
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManagerDidUpdateState(_:)
//
// [CBPeripheralManagerState.poweredOff]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/poweredOff
// [CBPeripheralManagerState.poweredOn]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/poweredOn
// [CBPeripheralManagerState]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidUpdateState(peripheral ICBPeripheralManager) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManagerDidUpdateState:"), peripheral)
}

// Tells the delegate the system is about to restore the peripheral manager.
//
// peripheral: The peripheral manager undergoing state restoration.
//
// dict: A dictionary that contains information about the peripheral manager, which
// the system preserved when the app terminated. For the available keys to
// this dictionary, see [Peripheral Manager State Restoration Options].
//
// # Discussion
//
// For apps that opt in to the state preservation and restoration feature,
// Core Bluetooth invokes this method when relaunching your app into the
// background to complete some Bluetooth-related task. Use this method to
// synchronize the state of your app with the state of the Bluetooth system.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:willRestoreState:)
//
// [Peripheral Manager State Restoration Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-manager-state-restoration-options
func (o CBPeripheralManagerDelegateObject) PeripheralManagerWillRestoreState(peripheral ICBPeripheralManager, dict foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:willRestoreState:"), peripheral, dict)
}

// Tells the delegate the peripheral manager published a service to the local
// GATT database.
//
// peripheral: The peripheral manager adding the service.
//
// service: The service added to the local GATT database.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheralManager.AddService] method to publish a service to the local
// peripheral’s GATT database. If the service published successfully to the
// local database, the `error` parameter is `nil`. If unsuccessful, the
// `error` parameter provides the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didAdd:error:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidAddServiceError(peripheral ICBPeripheralManager, service ICBService, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didAddService:error:"), peripheral, service, error_)
}

// Tells the delegate the peripheral manager started advertising the local
// peripheral device’s data.
//
// peripheral: The peripheral manager that is starting advertising.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth calls this method when your app calls the
// [CBPeripheralManager.StartAdvertising] method to advertise the local
// peripheral device’s data. If successful, the `error` parameter is `nil`.
// If a problem prevents advertising the data, the `error` parameter returns
// the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManagerDidStartAdvertising(_:error:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidStartAdvertisingError(peripheral ICBPeripheralManager, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManagerDidStartAdvertising:error:"), peripheral, error_)
}

// Tells the delegate that a remote central device subscribed to a
// characteristic’s value.
//
// peripheral: The peripheral manager connected to the remote central.
//
// central: The remote central device that subscribed to the characteristic’s value.
//
// characteristic: The characteristic subscribed to.
//
// # Discussion
//
// Core Bluetooth invokes this method when a remote central device subscribes
// to the value of one of the local peripheral’s characteristics, by
// enabling notifications or indications on the characteristic’s value. When
// called, start sending the subscribed central updates as the
// characteristic’s value changes. To send updated characteristic values to
// subscribed centrals, use the
// [CBPeripheralManager.UpdateValueForCharacteristicOnSubscribedCentrals]
// method of the [CBPeripheralManager] class.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:central:didSubscribeTo:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerCentralDidSubscribeToCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:central:didSubscribeToCharacteristic:"), peripheral, central, characteristic)
}

// Tells the delegate that a remote central device unsubscribed from a
// characteristic’s value.
//
// peripheral: The peripheral manager connected to the remote central.
//
// central: The remote central device that subscribed to the characteristic’s value.
//
// characteristic: The characteristic unsubscribed from.
//
// # Discussion
//
// Core Bluetooth calls this method when a remote central device unsubscribes
// from the value of one of the local peripheral’s characteristics, by
// disabling notifications or indications on the characteristic’s value.
// When called, stop sending the subscribed central updates of updates to the
// characteristic’s value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:central:didUnsubscribeFrom:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerCentralDidUnsubscribeFromCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:central:didUnsubscribeFromCharacteristic:"), peripheral, central, characteristic)
}

// Tells the delegate that a local peripheral device is ready to send
// characteristic value updates.
//
// peripheral: The peripheral manager that sends characteristic value updates.
//
// # Discussion
//
// When a call to the
// [CBPeripheralManager.UpdateValueForCharacteristicOnSubscribedCentrals]
// method fails because the underlying queue used to transmit the updated
// characteristic value is full, Core Bluetooth calls the
// [PeripheralManagerIsReadyToUpdateSubscribers] method when more space in the
// transmit queue becomes available. You can then implement this delegate
// method to resend the value.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManagerIsReady(toUpdateSubscribers:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerIsReadyToUpdateSubscribers(peripheral ICBPeripheralManager) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManagerIsReadyToUpdateSubscribers:"), peripheral)
}

// Tells the delegate that a local peripheral received an Attribute Protocol
// (ATT) read request for a characteristic with a dynamic value.
//
// peripheral: The peripheral manager that received the request.
//
// request: A [CBATTRequest] object that represents a request to read a
// characteristic’s value.
//
// # Discussion
//
// When you receive this callback, call the
// [CBPeripheralManager.RespondToRequestWithResult] method of the
// [CBPeripheralManager] class exactly once to respond to the read request.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didReceiveRead:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidReceiveReadRequest(peripheral ICBPeripheralManager, request ICBATTRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didReceiveReadRequest:"), peripheral, request)
}

// Tells the delegate that a local peripheral device received an Attribute
// Protocol (ATT) write request for a characteristic with a dynamic value.
//
// peripheral: The peripheral manager that received the request.
//
// requests: A list of one or more [CBATTRequest] objects, each representing a request
// to write the value of a characteristic.
//
// # Discussion
//
// In the same way that you respond to a read request, each time you receive
// this callback, call the [CBPeripheralManager.RespondToRequestWithResult]
// method of the [CBPeripheralManager] class exactly once. If the `requests`
// parameter contains multiple requests, treat them as you would a single
// request—if you can’t fulfill an individual request, you shouldn’t
// fulfill any of them. Instead, call the
// [CBPeripheralManager.RespondToRequestWithResult] method immediately, and
// provide a result that indicates the cause of the failure.
//
// When you respond to a write request, note that the first parameter of the
// [CBPeripheralManager.RespondToRequestWithResult] method expects a single
// [CBATTRequest] object, even though you received an array of them from the
// [PeripheralManagerDidReceiveWriteRequests] method. To respond properly,
// pass in the first request of the `requests` array.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didReceiveWrite:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidReceiveWriteRequests(peripheral ICBPeripheralManager, requests []CBATTRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didReceiveWriteRequests:"), peripheral, objectivec.IObjectSliceToNSArray(requests))
}

// Tells the delegate that the peripheral manager created a listener for
// incoming L2CAP channel connections.
//
// peripheral: The peripheral manager that published the channel.
//
// PSM: The Protocol/Service Multiplexer (PSM) of the published channel.
//
// error: The error that prevented publishing, or `nil` if no error occurred.
//
// # Discussion
//
// The peripheral manager calls this method after you call
// [CBPeripheralManager.PublishL2CAPChannelWithEncryption]. The [PSM]
// parameter contains the PSM assigned for the published channel.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didPublishL2CAPChannel:error:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidPublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didPublishL2CAPChannel:error:"), peripheral, PSM, error_)
}

// Tells the delegate that the peripheral manager removed a published service
// from the local system.
//
// peripheral: The peripheral manager that stopped publishing.
//
// PSM: The Protocol/Service Multiplexer (PSM) of the channel that was unpublished.
//
// error: The error that occurred, or `nil` if no error occurred.
//
// # Discussion
//
// The peripheral manager calls this method after you call
// [CBPeripheralManager.UnpublishL2CAPChannel].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didUnpublishL2CAPChannel:error:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidUnpublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didUnpublishL2CAPChannel:error:"), peripheral, PSM, error_)
}

// Tells the delegate that the peripheral manager opened an L2CAP channel.
//
// peripheral: The peripheral manager that opened the channel.
//
// channel: The channel opened by the manager.
//
// error: The error that occurred, or `nil` if no error occurred.
//
// # Discussion
//
// The peripheral manager calls this method after you call
// [CBPeripheralManager.PublishL2CAPChannelWithEncryption].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerDelegate/peripheralManager(_:didOpen:error:)
func (o CBPeripheralManagerDelegateObject) PeripheralManagerDidOpenL2CAPChannelError(peripheral ICBPeripheralManager, channel ICBL2CAPChannel, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralManager:didOpenL2CAPChannel:error:"), peripheral, channel, error_)
}

// CBPeripheralManagerDelegateConfig holds optional typed callbacks for [CBPeripheralManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbperipheralmanagerdelegate
type CBPeripheralManagerDelegateConfig struct {

	// Monitoring Changes to the Peripheral Manager’s State
	// PeripheralManagerDidUpdateState — Tells the delegate the peripheral manager’s state updated.
	PeripheralManagerDidUpdateState func(peripheral CBPeripheralManager)
	// PeripheralManagerWillRestoreState — Tells the delegate the system is about to restore the peripheral manager.
	PeripheralManagerWillRestoreState func(peripheral CBPeripheralManager, dict foundation.INSDictionary)

	// Advertising Peripheral Data
	// PeripheralManagerDidStartAdvertisingError — Tells the delegate the peripheral manager started advertising the local peripheral device’s data.
	PeripheralManagerDidStartAdvertisingError func(peripheral CBPeripheralManager, error_ foundation.NSError)

	// Using L2CAP Channels
	// PeripheralManagerDidPublishL2CAPChannelError — Tells the delegate that the peripheral manager created a listener for incoming L2CAP channel connections.
	PeripheralManagerDidPublishL2CAPChannelError func(peripheral CBPeripheralManager, PSM CBL2CAPPSM, error_ foundation.NSError)
	// PeripheralManagerDidUnpublishL2CAPChannelError — Tells the delegate that the peripheral manager removed a published service from the local system.
	PeripheralManagerDidUnpublishL2CAPChannelError func(peripheral CBPeripheralManager, PSM CBL2CAPPSM, error_ foundation.NSError)

	// Other Methods
	// PeripheralManagerDidAddServiceError — Tells the delegate the peripheral manager published a service to the local GATT database.
	PeripheralManagerDidAddServiceError func(peripheral CBPeripheralManager, service CBService, error_ foundation.NSError)
	// PeripheralManagerCentralDidSubscribeToCharacteristic — Tells the delegate that a remote central device subscribed to a characteristic’s value.
	PeripheralManagerCentralDidSubscribeToCharacteristic func(peripheral CBPeripheralManager, central CBCentral, characteristic CBCharacteristic)
	// PeripheralManagerCentralDidUnsubscribeFromCharacteristic — Tells the delegate that a remote central device unsubscribed from a characteristic’s value.
	PeripheralManagerCentralDidUnsubscribeFromCharacteristic func(peripheral CBPeripheralManager, central CBCentral, characteristic CBCharacteristic)
	// PeripheralManagerIsReadyToUpdateSubscribers — Tells the delegate that a local peripheral device is ready to send characteristic value updates.
	PeripheralManagerIsReadyToUpdateSubscribers func(peripheral CBPeripheralManager)
	// PeripheralManagerDidReceiveReadRequest — Tells the delegate that a local peripheral received an Attribute Protocol (ATT) read request for a characteristic with a dynamic value.
	PeripheralManagerDidReceiveReadRequest func(peripheral CBPeripheralManager, request CBATTRequest)
	// PeripheralManagerDidOpenL2CAPChannelError — Tells the delegate that the peripheral manager opened an L2CAP channel.
	PeripheralManagerDidOpenL2CAPChannelError func(peripheral CBPeripheralManager, channel CBL2CAPChannel, error_ foundation.NSError)
}

// NewCBPeripheralManagerDelegate creates an Objective-C object implementing the [CBPeripheralManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CBPeripheralManagerDelegateObject] satisfies the [CBPeripheralManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbperipheralmanagerdelegate
func NewCBPeripheralManagerDelegate(config CBPeripheralManagerDelegateConfig) CBPeripheralManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCBPeripheralManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.PeripheralManagerDidUpdateState != nil {
		fn := config.PeripheralManagerDidUpdateState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManagerDidUpdateState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManagerDidUpdateState:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				fn(peripheral)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerWillRestoreState != nil {
		fn := config.PeripheralManagerWillRestoreState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:willRestoreState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, dictID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:willRestoreState:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				dict := foundation.NSDictionaryFromID(dictID)
				fn(peripheral, dict)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidAddServiceError != nil {
		fn := config.PeripheralManagerDidAddServiceError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:didAddService:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, serviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:didAddService:error:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				service := CBServiceFromID(serviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, service, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidStartAdvertisingError != nil {
		fn := config.PeripheralManagerDidStartAdvertisingError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManagerDidStartAdvertising:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManagerDidStartAdvertising:error:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerCentralDidSubscribeToCharacteristic != nil {
		fn := config.PeripheralManagerCentralDidSubscribeToCharacteristic
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:central:didSubscribeToCharacteristic:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, centralID objc.ID, characteristicID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:central:didSubscribeToCharacteristic:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				central := CBCentralFromID(centralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				fn(peripheral, central, characteristic)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerCentralDidUnsubscribeFromCharacteristic != nil {
		fn := config.PeripheralManagerCentralDidUnsubscribeFromCharacteristic
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:central:didUnsubscribeFromCharacteristic:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, centralID objc.ID, characteristicID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:central:didUnsubscribeFromCharacteristic:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				central := CBCentralFromID(centralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				fn(peripheral, central, characteristic)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerIsReadyToUpdateSubscribers != nil {
		fn := config.PeripheralManagerIsReadyToUpdateSubscribers
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManagerIsReadyToUpdateSubscribers:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManagerIsReadyToUpdateSubscribers:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				fn(peripheral)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidReceiveReadRequest != nil {
		fn := config.PeripheralManagerDidReceiveReadRequest
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:didReceiveReadRequest:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, requestID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:didReceiveReadRequest:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				request := CBATTRequestFromID(requestID)
				fn(peripheral, request)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidPublishL2CAPChannelError != nil {
		fn := config.PeripheralManagerDidPublishL2CAPChannelError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:didPublishL2CAPChannel:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, PSM CBL2CAPPSM, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:didPublishL2CAPChannel:error:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, PSM, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidUnpublishL2CAPChannelError != nil {
		fn := config.PeripheralManagerDidUnpublishL2CAPChannelError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:didUnpublishL2CAPChannel:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, PSM CBL2CAPPSM, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:didUnpublishL2CAPChannel:error:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, PSM, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralManagerDidOpenL2CAPChannelError != nil {
		fn := config.PeripheralManagerDidOpenL2CAPChannelError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralManager:didOpenL2CAPChannel:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, channelID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralManagerDelegate", "peripheralManager:didOpenL2CAPChannel:error:")
					}
				}()
				peripheral := CBPeripheralManagerFromID(peripheralID)
				channel := CBL2CAPChannelFromID(channelID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, channel, error_)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CBPeripheralManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCBPeripheralManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CBPeripheralManagerDelegateObjectFromID(instance)
}
