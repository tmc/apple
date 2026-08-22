// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

// Package iobluetooth provides Go bindings for the IOBluetooth framework.
//
// Gain user-space access to Bluetooth devices.
//
// The Bluetooth framework supports user-space access to Bluetooth devices,
// including both C and Objective-C APIs.
//
// # Classes
//
//   - [IOBluetoothDevice]: An instance of IOBluetoothDevice represents a single remote Bluetooth device.
//   - [IOBluetoothDeviceInquiry]: Object representing a device inquiry that finds Bluetooth devices in-range of the computer, and (optionally) retrieves name information for them.
//   - [IOBluetoothDevicePair]: An instance of IOBluetoothDevicePair represents a pairing attempt to a remote Bluetooth device.
//   - [IOBluetoothDeviceRef]: An object that represents a Bluetooth I/O device.
//   - [IOBluetoothHandsFree]: Hands free profile class.
//   - [IOBluetoothHandsFreeAudioGateway]: An object that sends data to a connected Bluetooth hands-free phone or headset and processes commands from it.
//   - [IOBluetoothHandsFreeDevice]: An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHostController]: This class is a representation of a Bluetooth Host Controller Interface that is present on the local computer (either plugged in externally or available internally).
//   - [IOBluetoothL2CAPChannel]: An instance of IOBluetoothL2CAPChannel represents a single open L2CAP channel.
//   - [IOBluetoothL2CAPChannelRef]
//   - [IOBluetoothOBEXSession]: An OBEX Session with a Bluetooth RFCOMM channel as the transport.
//   - [IOBluetoothObject]
//   - [IOBluetoothObjectRef]
//   - [IOBluetoothRFCOMMChannel]: An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
//   - [IOBluetoothRFCOMMChannelRef]
//   - [IOBluetoothSDPDataElement]: An instance of this class represents a single SDP data element as defined by the Bluetooth SDP spec.
//   - [IOBluetoothSDPDataElementRef]
//   - [IOBluetoothSDPServiceAttribute]: IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
//   - [IOBluetoothSDPServiceRecord]: An instance of this class represents a single SDP service record.
//   - [IOBluetoothSDPServiceRecordRef]
//   - [IOBluetoothSDPUUID]: An NSData subclass that represents a UUID as defined in the Bluetooth SDP spec.
//   - [IOBluetoothSDPUUIDRef]
//   - [IOBluetoothUserNotification]: Represents a registered notification.
//   - [IOBluetoothUserNotificationRef]
//   - [OBEXFileTransferServices]: Implements advanced OBEX operations in addition to simple PUT and GET.
//   - [OBEXSession]: Object representing an OBEX connection to a remote target. ([OBEXTransportEvent], [OBEXTransportEventType])
//
// # Protocols
//
//   - [IOBluetoothDeviceAsyncCallbacks]
//   - [IOBluetoothDeviceInquiryDelegate]: This category on NSObject describes the delegate methods for the IOBluetoothDeviceInquiry object.
//   - [IOBluetoothDevicePairDelegate]
//   - [IOBluetoothHandsFreeAudioGatewayDelegate]: A set of optional methods for receiving information about status changes for a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHandsFreeDelegate]
//   - [IOBluetoothHandsFreeDeviceDelegate]: A set of optional methods for receiving status change updates and information about a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothL2CAPChannelDelegate]
//   - [IOBluetoothRFCOMMChannelDelegate]
//
// # Key Types
//
//   - [IOBluetoothDevice] - An instance of IOBluetoothDevice represents a single remote Bluetooth device.
//   - [IOBluetoothOBEXSession] - An OBEX Session with a Bluetooth RFCOMM channel as the transport.
//   - [OBEXSession] - Object representing an OBEX connection to a remote target.
//   - [IOBluetoothHandsFreeDevice] - An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHandsFree] - Hands free profile class.
//   - [IOBluetoothL2CAPChannel] - An instance of IOBluetoothL2CAPChannel represents a single open L2CAP channel.
//   - [IOBluetoothRFCOMMChannel] - An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
//   - [OBEXFileTransferServices] - Implements advanced OBEX operations in addition to simple PUT and GET.
//   - [IOBluetoothSDPServiceRecord] - An instance of this class represents a single SDP service record.
//   - [IOBluetoothSDPDataElement] - An instance of this class represents a single SDP data element as defined by the Bluetooth SDP spec.
package iobluetooth

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the IOBluetooth library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/IOBluetooth.framework/IOBluetooth",
	"/usr/lib/libIOBluetooth.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: IOBluetooth: failed to load framework from any known path\n")
	}
}
