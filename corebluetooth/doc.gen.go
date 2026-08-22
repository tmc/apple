// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

// Package corebluetooth provides Go bindings for the CoreBluetooth framework.
//
// Communicate with Bluetooth low energy and BR/EDR (“Classic”) Devices.
//
// The Core Bluetooth framework provides the classes needed for your apps to
// communicate with Bluetooth-equipped low energy (LE) and Basic Rate /
// Enhanced Data Rate (BR/EDR) wireless technology.
//
// # Centrals
//
//   - [CBCentral]: A remote device connected to a local app, which is acting as a peripheral.
//   - [CBCentralManager]: An object that scans for, discovers, connects to, and manages peripherals. ([CBConnectionEvent], [CBConnectionEventMatchingOption])
//   - [CBCentralManagerDelegate]: A protocol that provides updates for the discovery and management of peripheral devices.
//
// # Peripherals
//
//   - [CBPeripheral]: A remote peripheral device. ([CBCharacteristicWriteType], [CBPeripheralState], [CBL2CAPChannel], [CBL2CAPPSM])
//   - [CBPeripheralDelegate]: A protocol that provides updates on the use of a peripheral’s services.
//   - [CBPeripheralManager]: An object that manages and advertises peripheral services exposed by this app. ([CBPeripheralManagerAuthorizationStatus], [CBPeripheralManagerState], [CBPeripheralManagerConnectionLatency])
//   - [CBPeripheralManagerDelegate]: A protocol that provides updates for local peripheral state and interactions with remote central devices.
//   - [CBAttribute]: A representation of common aspects of services offered by a peripheral.
//   - [CBAttributePermissions]: Values that represent the read, write, and encryption permissions for a characteristic’s value.
//
// # Channel Sounding
//
//   - [CBUUIDCharacteristicObservationScheduleString]
//
// # Services
//
//   - [CBService]: A collection of data and associated behaviors that accomplish a function or feature of a device.
//   - [CBMutableService]: A service with writeable property values.
//   - [CBCharacteristic]: A characteristic of a remote peripheral’s service. ([CBCharacteristicProperties])
//   - [CBMutableCharacteristic]: A characteristic of a local peripheral’s service. ([CBAttributePermissions])
//   - [CBDescriptor]: An object that provides further information about a remote peripheral’s characteristic.
//   - [CBMutableDescriptor]: An object that provides additional information about a local peripheral’s characteristic.
//
// # Supporting Types
//
//   - [CBManager]: The abstract base class that manages central and peripheral objects. ([CBManagerState], [CBManagerAuthorization])
//   - [CBATTRequest]: A request that uses the Attribute Protocol (ATT).
//   - [CBPeer]: An object that represents a remote device.
//   - [CBUUID]: A universally unique identifier, as defined by Bluetooth standards.
//
// # Bluetooth Classic Support
//
//   - [Using Core Bluetooth Classic]: Discover and communicate with a Bluetooth Classic device by using the Core Bluetooth APIs.
//
// # Errors
//
//   - [CBErrorDomain]: The domain for Core Bluetooth errors.
//   - [CBError]: The codes for errors that Core Bluetooth returns during Bluetooth transactions.
//   - [CBATTErrorDomain]: The domain for Core Bluetooth ATT errors.
//   - [CBATTError]: The possible errors returned by a GATT server (a remote peripheral) during Bluetooth low energy ATT transactions.
//
// # Enumerations
//
//   - [CBChannelSoundingSessionConfigurationRole]//
//
// # Key Types
//
//   - [CBPeripheral] - A remote peripheral device.
//   - [CBPeripheralManager] - An object that manages and advertises peripheral services exposed by this app.
//   - [CBCentralManager] - An object that scans for, discovers, connects to, and manages peripherals.
//   - [CBMutableCharacteristic] - A characteristic of a local peripheral’s service.
//   - [CBCharacteristic] - A characteristic of a remote peripheral’s service.
//   - [CBUUID] - A universally unique identifier, as defined by Bluetooth standards.
//   - [CBATTRequest] - A request that uses the Attribute Protocol (ATT).
//   - [CBL2CAPChannel] - A live L2CAP connection to a remote device.
//   - [CBService] - A collection of data and associated behaviors that accomplish a function or feature of a device.
//   - [CBMutableService] - A service with writeable property values.
//
// [Using Core Bluetooth Classic]: https://developer.apple.com/documentation/corebluetooth/using-core-bluetooth-classic
package corebluetooth

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreBluetooth library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreBluetooth.framework/CoreBluetooth",
	"/usr/lib/libCoreBluetooth.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreBluetooth: failed to load framework from any known path\n")
	}
}
