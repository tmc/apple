// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that provides updates on the use of a peripheral’s services.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate
type CBPeripheralDelegate interface {
	objectivec.IObject
}

// CBPeripheralDelegateObject wraps an existing Objective-C object that conforms to the CBPeripheralDelegate protocol.
type CBPeripheralDelegateObject struct {
	objectivec.Object
}

func (o CBPeripheralDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CBPeripheralDelegateObjectFromID constructs a [CBPeripheralDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CBPeripheralDelegateObjectFromID(id objc.ID) CBPeripheralDelegateObject {
	return CBPeripheralDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that peripheral service discovery succeeded.
//
// peripheral: The peripheral to which the services belong.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.DiscoverServices] method. If the peripheral successfully
// discovers services, you can access them through the peripheral’s
// [CBPeripheral.Services] property. If successful, the `error` parameter is
// `nil`. If unsuccessful, the `error` parameter returns the cause of the
// failure.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didDiscoverServices:)
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
func (o CBPeripheralDelegateObject) PeripheralDidDiscoverServices(peripheral ICBPeripheral, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didDiscoverServices:"), peripheral, error_)
}

// Tells the delegate that discovering included services within the indicated
// service completed.
//
// peripheral: The peripheral providing this information.
//
// service: The [CBService] object containing the included service.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.DiscoverIncludedServicesForService] method. If the peripheral
// successfully discovers services, you can access them through the
// service’s [CBService.IncludedServices] property. If successful, the
// `error` parameter is `nil`. If unsuccessful, the `error` parameter returns
// the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didDiscoverIncludedServicesFor:error:)
func (o CBPeripheralDelegateObject) PeripheralDidDiscoverIncludedServicesForServiceError(peripheral ICBPeripheral, service ICBService, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didDiscoverIncludedServicesForService:error:"), peripheral, service, error_)
}

// Tells the delegate that the peripheral found characteristics for a service.
//
// peripheral: The peripheral providing this information.
//
// service: The service to which the characteristics belong.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.DiscoverCharacteristicsForService] method. If the peripheral
// successfully discovers the characteristics of the specified service, you
// can access them through the service’s [CBService.Characteristics]
// property. If successful, the `error` parameter is `nil`. If unsuccessful,
// the `error` parameter returns the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didDiscoverCharacteristicsFor:error:)
func (o CBPeripheralDelegateObject) PeripheralDidDiscoverCharacteristicsForServiceError(peripheral ICBPeripheral, service ICBService, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didDiscoverCharacteristicsForService:error:"), peripheral, service, error_)
}

// Tells the delegate that the peripheral found descriptors for a
// characteristic.
//
// peripheral: The peripheral providing this information.
//
// characteristic: The characteristic to which the characteristic descriptors belong.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.DiscoverDescriptorsForCharacteristic] method. If the
// peripheral successfully discovers the descriptors of the specified
// characteristic, you can access them through the characteristic’s
// [CBCharacteristic.Descriptors] property. If successful, the `error`
// parameter is `nil`. If unsuccessful, the `error` parameter returns the
// cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didDiscoverDescriptorsFor:error:)
func (o CBPeripheralDelegateObject) PeripheralDidDiscoverDescriptorsForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didDiscoverDescriptorsForCharacteristic:error:"), peripheral, characteristic, error_)
}

// Tells the delegate that retrieving the specified characteristic’s value
// succeeded, or that the characteristic’s value changed.
//
// peripheral: The peripheral providing this information.
//
// characteristic: The characteristic containing the value.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.ReadValueForDescriptor] method. A peripheral also invokes
// this method to notify your app of a change to the value of the
// characteristic for which the app previously enabled notifications by
// calling [CBPeripheral.SetNotifyValueForCharacteristic]. If successful, the
// `error` parameter is `nil`. If unsuccessful, the `error` parameter returns
// the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didUpdateValueFor:error:)-1xyna
func (o CBPeripheralDelegateObject) PeripheralDidUpdateValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didUpdateValueForCharacteristic:error:"), peripheral, characteristic, error_)
}

// Tells the delegate that retrieving a specified characteristic
// descriptor’s value succeeded.
//
// peripheral: The peripheral providing this information.
//
// descriptor: The characteristic descriptor containing the value.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.ReadValueForDescriptor] method. If successful, the `error`
// parameter is `nil`. If unsuccessful, the `error` parameter returns the
// cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didUpdateValueFor:error:)-1t3wm
func (o CBPeripheralDelegateObject) PeripheralDidUpdateValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didUpdateValueForDescriptor:error:"), peripheral, descriptor, error_)
}

// Tells the delegate that the peripheral successfully set a value for the
// characteristic.
//
// peripheral: The peripheral providing this information.
//
// characteristic: The characteristic containing the value.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method only when your app calls the
// [CBPeripheral.WriteValueForCharacteristicType] method with the
// [CBCharacteristicWriteType.withResponse] constant specified as the write
// type. If successful, the `error` parameter is `nil`. If unsuccessful, the
// `error` parameter returns the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didWriteValueFor:error:)-4f5ea
//
// [CBCharacteristicWriteType.withResponse]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType/withResponse
func (o CBPeripheralDelegateObject) PeripheralDidWriteValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didWriteValueForCharacteristic:error:"), peripheral, characteristic, error_)
}

// Tells the delegate that the peripheral successfully set a value for the
// descriptor.
//
// peripheral: The peripheral providing this information.
//
// descriptor: The characteristic descriptor containing the value.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.WriteValueForDescriptor] method. If successful, the `error`
// parameter is `nil`. If unsuccessful, the `error` parameter returns the
// cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didWriteValueFor:error:)-1ybl3
func (o CBPeripheralDelegateObject) PeripheralDidWriteValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didWriteValueForDescriptor:error:"), peripheral, descriptor, error_)
}

// Tells the delegate that a peripheral is again ready to send characteristic
// updates.
//
// peripheral: The peripheral providing this update.
//
// # Discussion
//
// The peripheral calls this delegate method after a failed call to
// [CBPeripheral.WriteValueForCharacteristicType], once `peripheral` is ready
// to send characteristic value updates.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheralIsReady(toSendWriteWithoutResponse:)
func (o CBPeripheralDelegateObject) PeripheralIsReadyToSendWriteWithoutResponse(peripheral ICBPeripheral) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralIsReadyToSendWriteWithoutResponse:"), peripheral)
}

// Tells the delegate that the peripheral received a request to start or stop
// providing notifications for a specified characteristic’s value.
//
// peripheral: The peripheral providing this information.
//
// characteristic: The characteristic for which to configure value notifications.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.SetNotifyValueForCharacteristic] method. If successful, the
// `error` parameter is `nil`. If unsuccessful, the `error` parameter returns
// the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didUpdateNotificationStateFor:error:)
func (o CBPeripheralDelegateObject) PeripheralDidUpdateNotificationStateForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didUpdateNotificationStateForCharacteristic:error:"), peripheral, characteristic, error_)
}

// Tells the delegate that retrieving the value of the peripheral’s current
// Received Signal Strength Indicator (RSSI) succeeded.
//
// peripheral: The peripheral providing this information.
//
// RSSI: The RSSI, in decibels, of the peripheral.
//
// error: The reason the call failed, or `nil` if no error occurred.
//
// # Discussion
//
// Core Bluetooth invokes this method when your app calls the
// [CBPeripheral.ReadRSSI] method, while the peripheral is connected to the
// central manager. If successful, the `error` parameter is `nil` and the
// parameter [RSSI] reports the peripheral’s signal strength, in decibels.
// If unsuccessful, the `error` parameter returns the cause of the failure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didReadRSSI:error:)
func (o CBPeripheralDelegateObject) PeripheralDidReadRSSIError(peripheral ICBPeripheral, RSSI foundation.NSNumber, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didReadRSSI:error:"), peripheral, RSSI, error_)
}

// Tells the delegate that a peripheral’s name changed.
//
// peripheral: The peripheral providing this information.
//
// # Discussion
//
// Core Bluetooth invokes this method whenever the peripheral’s Generic
// Access Profile (GAP) device name changes. Since a peripheral device can
// change its GAP device name, you can implement this method if your app needs
// to display the current name of the peripheral device.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheralDidUpdateName(_:)
func (o CBPeripheralDelegateObject) PeripheralDidUpdateName(peripheral ICBPeripheral) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheralDidUpdateName:"), peripheral)
}

// Tells the delegate that a peripheral’s services changed.
//
// peripheral: The peripheral providing this information.
//
// invalidatedServices: A list of services invalidated by this change.
//
// # Discussion
//
// Core Bluetooth invokes this method whenever one or more services of a
// peripheral change. A peripheral’s services have changed if:
//
// - The peripheral removes a service from its database. - The peripheral adds
// a new service to its database. - The peripheral adds back a
// previously-removed service, but at a different location in the database.
//
// The `invalidatedServices` parameter includes any changed services that you
// previously discovered; you can no longer use these services. You can use
// the [CBPeripheral.DiscoverServices] method to discover any new services
// that the peripheral added to its database. Use this same method to find out
// whether any of the invalidated services that you were using (and want to
// continue using) now have a different location in the peripheral’s
// database.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didModifyServices:)
func (o CBPeripheralDelegateObject) PeripheralDidModifyServices(peripheral ICBPeripheral, invalidatedServices []CBService) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didModifyServices:"), peripheral, objectivec.IObjectSliceToNSArray(invalidatedServices))
}

// Delivers the result of an attempt to open an L2CAP channel.
//
// # Discussion
//
// This method delivers the result of a previous call to
// [CBPeripheral.OpenL2CAPChannel].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didOpen:error:)
func (o CBPeripheralDelegateObject) PeripheralDidOpenL2CAPChannelError(peripheral ICBPeripheral, channel ICBL2CAPChannel, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didOpenL2CAPChannel:error:"), peripheral, channel, error_)
}

// peripheral: The peripheral providing this update.
//
// error: If an error occurred, the cause of the failure.
//
// # Discussion
//
// This method is called when a channel sounding session completes.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didCompleteChannelSoundingSession:)
func (o CBPeripheralDelegateObject) PeripheralDidCompleteChannelSoundingSession(peripheral ICBPeripheral, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didCompleteChannelSoundingSession:"), peripheral, error_)
}

// peripheral: The peripheral providing this update.
//
// results: An object containing the results of a channel sounding procedure.
//
// error: If an error occurred, the cause of the failure.
//
// # Discussion
//
// This method returns the results of a channel sounding procedure.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralDelegate/peripheral(_:didReceive:error:)
func (o CBPeripheralDelegateObject) PeripheralDidReceiveChannelSoundingProcedureResultsError(peripheral ICBPeripheral, results *uintptr, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("peripheral:didReceiveChannelSoundingProcedureResults:error:"), peripheral, results, error_)
}

// CBPeripheralDelegateConfig holds optional typed callbacks for [CBPeripheralDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbperipheraldelegate
type CBPeripheralDelegateConfig struct {

	// Discovering Services
	// PeripheralDidDiscoverServices — Tells the delegate that peripheral service discovery succeeded.
	PeripheralDidDiscoverServices func(peripheral CBPeripheral, error_ foundation.NSError)

	// Retrieving a Peripheral’s RSSI Data
	// PeripheralDidReadRSSIError — Tells the delegate that retrieving the value of the peripheral’s current Received Signal Strength Indicator (RSSI) succeeded.
	PeripheralDidReadRSSIError func(peripheral CBPeripheral, RSSI foundation.NSNumber, error_ foundation.NSError)

	// Monitoring Changes to a Peripheral’s Name or Services
	// PeripheralDidUpdateName — Tells the delegate that a peripheral’s name changed.
	PeripheralDidUpdateName func(peripheral CBPeripheral)

	// Other Methods
	// PeripheralDidDiscoverIncludedServicesForServiceError — Tells the delegate that discovering included services within the indicated service completed.
	PeripheralDidDiscoverIncludedServicesForServiceError func(peripheral CBPeripheral, service CBService, error_ foundation.NSError)
	// PeripheralDidDiscoverCharacteristicsForServiceError — Tells the delegate that the peripheral found characteristics for a service.
	PeripheralDidDiscoverCharacteristicsForServiceError func(peripheral CBPeripheral, service CBService, error_ foundation.NSError)
	// PeripheralDidDiscoverDescriptorsForCharacteristicError — Tells the delegate that the peripheral found descriptors for a characteristic.
	PeripheralDidDiscoverDescriptorsForCharacteristicError func(peripheral CBPeripheral, characteristic CBCharacteristic, error_ foundation.NSError)
	// PeripheralDidUpdateValueForCharacteristicError — Tells the delegate that retrieving the specified characteristic’s value succeeded, or that the characteristic’s value changed.
	PeripheralDidUpdateValueForCharacteristicError func(peripheral CBPeripheral, characteristic CBCharacteristic, error_ foundation.NSError)
	// PeripheralDidUpdateValueForDescriptorError — Tells the delegate that retrieving a specified characteristic descriptor’s value succeeded.
	PeripheralDidUpdateValueForDescriptorError func(peripheral CBPeripheral, descriptor CBDescriptor, error_ foundation.NSError)
	// PeripheralDidWriteValueForCharacteristicError — Tells the delegate that the peripheral successfully set a value for the characteristic.
	PeripheralDidWriteValueForCharacteristicError func(peripheral CBPeripheral, characteristic CBCharacteristic, error_ foundation.NSError)
	// PeripheralDidWriteValueForDescriptorError — Tells the delegate that the peripheral successfully set a value for the descriptor.
	PeripheralDidWriteValueForDescriptorError func(peripheral CBPeripheral, descriptor CBDescriptor, error_ foundation.NSError)
	// PeripheralIsReadyToSendWriteWithoutResponse — Tells the delegate that a peripheral is again ready to send characteristic updates.
	PeripheralIsReadyToSendWriteWithoutResponse func(peripheral CBPeripheral)
	// PeripheralDidUpdateNotificationStateForCharacteristicError — Tells the delegate that the peripheral received a request to start or stop providing notifications for a specified characteristic’s value.
	PeripheralDidUpdateNotificationStateForCharacteristicError func(peripheral CBPeripheral, characteristic CBCharacteristic, error_ foundation.NSError)
	// PeripheralDidOpenL2CAPChannelError — Delivers the result of an attempt to open an L2CAP channel.
	PeripheralDidOpenL2CAPChannelError func(peripheral CBPeripheral, channel CBL2CAPChannel, error_ foundation.NSError)
}

// NewCBPeripheralDelegate creates an Objective-C object implementing the [CBPeripheralDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CBPeripheralDelegateObject] satisfies the [CBPeripheralDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corebluetooth/cbperipheraldelegate
func NewCBPeripheralDelegate(config CBPeripheralDelegateConfig) CBPeripheralDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCBPeripheralDelegate_%d", n)

	var methods []objc.MethodDef

	if config.PeripheralDidDiscoverServices != nil {
		fn := config.PeripheralDidDiscoverServices
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didDiscoverServices:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didDiscoverServices:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidDiscoverIncludedServicesForServiceError != nil {
		fn := config.PeripheralDidDiscoverIncludedServicesForServiceError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didDiscoverIncludedServicesForService:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, serviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didDiscoverIncludedServicesForService:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				service := CBServiceFromID(serviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, service, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidDiscoverCharacteristicsForServiceError != nil {
		fn := config.PeripheralDidDiscoverCharacteristicsForServiceError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didDiscoverCharacteristicsForService:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, serviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didDiscoverCharacteristicsForService:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				service := CBServiceFromID(serviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, service, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidDiscoverDescriptorsForCharacteristicError != nil {
		fn := config.PeripheralDidDiscoverDescriptorsForCharacteristicError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didDiscoverDescriptorsForCharacteristic:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, characteristicID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didDiscoverDescriptorsForCharacteristic:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, characteristic, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidUpdateValueForCharacteristicError != nil {
		fn := config.PeripheralDidUpdateValueForCharacteristicError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didUpdateValueForCharacteristic:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, characteristicID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didUpdateValueForCharacteristic:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, characteristic, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidUpdateValueForDescriptorError != nil {
		fn := config.PeripheralDidUpdateValueForDescriptorError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didUpdateValueForDescriptor:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, descriptorID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didUpdateValueForDescriptor:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				descriptor := CBDescriptorFromID(descriptorID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, descriptor, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidWriteValueForCharacteristicError != nil {
		fn := config.PeripheralDidWriteValueForCharacteristicError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didWriteValueForCharacteristic:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, characteristicID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didWriteValueForCharacteristic:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, characteristic, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidWriteValueForDescriptorError != nil {
		fn := config.PeripheralDidWriteValueForDescriptorError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didWriteValueForDescriptor:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, descriptorID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didWriteValueForDescriptor:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				descriptor := CBDescriptorFromID(descriptorID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, descriptor, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralIsReadyToSendWriteWithoutResponse != nil {
		fn := config.PeripheralIsReadyToSendWriteWithoutResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralIsReadyToSendWriteWithoutResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheralIsReadyToSendWriteWithoutResponse:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				fn(peripheral)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidUpdateNotificationStateForCharacteristicError != nil {
		fn := config.PeripheralDidUpdateNotificationStateForCharacteristicError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didUpdateNotificationStateForCharacteristic:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, characteristicID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didUpdateNotificationStateForCharacteristic:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				characteristic := CBCharacteristicFromID(characteristicID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, characteristic, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidReadRSSIError != nil {
		fn := config.PeripheralDidReadRSSIError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didReadRSSI:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, RSSIID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didReadRSSI:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				RSSI := foundation.NSNumberFromID(RSSIID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, RSSI, error_)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidUpdateName != nil {
		fn := config.PeripheralDidUpdateName
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheralDidUpdateName:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheralDidUpdateName:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				fn(peripheral)
				_delegateDone = true
			},
		})
	}

	if config.PeripheralDidOpenL2CAPChannelError != nil {
		fn := config.PeripheralDidOpenL2CAPChannelError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("peripheral:didOpenL2CAPChannel:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, peripheralID objc.ID, channelID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CBPeripheralDelegate", "peripheral:didOpenL2CAPChannel:error:")
					}
				}()
				peripheral := CBPeripheralFromID(peripheralID)
				channel := CBL2CAPChannelFromID(channelID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(peripheral, channel, error_)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CBPeripheralDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCBPeripheralDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CBPeripheralDelegateObjectFromID(instance)
}
