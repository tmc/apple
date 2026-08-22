// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code
type CBATTError int

const (
	// CBATTErrorAttributeNotFound: The attribute wasn’t found within the specified attribute handle range.
	CBATTErrorAttributeNotFound CBATTError = 0xa
	// CBATTErrorAttributeNotLong: The ATT read blob request can’t read or write the attribute.
	CBATTErrorAttributeNotLong CBATTError = 0xb
	// CBATTErrorInsufficientAuthentication: Reading or writing the attribute’s value failed for lack of authentication.
	CBATTErrorInsufficientAuthentication CBATTError = 0x5
	// CBATTErrorInsufficientAuthorization: Reading or writing the attribute’s value failed for lack of authorization.
	CBATTErrorInsufficientAuthorization CBATTError = 0x8
	// CBATTErrorInsufficientEncryption: Reading or writing the attribute’s value failed for lack of encryption.
	CBATTErrorInsufficientEncryption CBATTError = 0xf
	// CBATTErrorInsufficientEncryptionKeySize: The encryption key size used for encrypting this link is insufficient.
	CBATTErrorInsufficientEncryptionKeySize CBATTError = 0xc
	// CBATTErrorInsufficientResources: Resources are insufficient to complete the ATT request.
	CBATTErrorInsufficientResources CBATTError = 0x11
	// CBATTErrorInvalidAttributeValueLength: The length of the attribute’s value is invalid for the intended operation.
	CBATTErrorInvalidAttributeValueLength CBATTError = 0xd
	// CBATTErrorInvalidHandle: The attribute handle is invalid on this peripheral.
	CBATTErrorInvalidHandle CBATTError = 0x1
	// CBATTErrorInvalidOffset: The specified offset value was past the end of the attribute’s value.
	CBATTErrorInvalidOffset CBATTError = 0x7
	// CBATTErrorInvalidPdu: The attribute Protocol Data Unit (PDU) is invalid.
	CBATTErrorInvalidPdu CBATTError = 0x4
	// CBATTErrorPrepareQueueFull: The prepare queue is full, as a result of there being too many write requests in the queue.
	CBATTErrorPrepareQueueFull CBATTError = 0x9
	// CBATTErrorReadNotPermitted: The permissions prohibit reading the attribute’s value.
	CBATTErrorReadNotPermitted CBATTError = 0x2
	// CBATTErrorRequestNotSupported: The attribute server doesn’t support the request received from the client.
	CBATTErrorRequestNotSupported CBATTError = 0x6
	// CBATTErrorSuccess: The ATT command or request successfully completed.
	CBATTErrorSuccess CBATTError = 0
	// CBATTErrorUnlikelyError: The ATT request encountered an unlikely error and wasn’t completed.
	CBATTErrorUnlikelyError CBATTError = 0xe
	// CBATTErrorUnsupportedGroupType: The attribute type isn’t a supported grouping attribute as defined by a higher-layer specification.
	CBATTErrorUnsupportedGroupType CBATTError = 0x10
	// CBATTErrorWriteNotPermitted: The permissions prohibit writing the attribute’s value.
	CBATTErrorWriteNotPermitted CBATTError = 0x3
)

func (e CBATTError) String() string {
	switch e {
	case CBATTErrorAttributeNotFound:
		return "CBATTErrorAttributeNotFound"
	case CBATTErrorAttributeNotLong:
		return "CBATTErrorAttributeNotLong"
	case CBATTErrorInsufficientAuthentication:
		return "CBATTErrorInsufficientAuthentication"
	case CBATTErrorInsufficientAuthorization:
		return "CBATTErrorInsufficientAuthorization"
	case CBATTErrorInsufficientEncryption:
		return "CBATTErrorInsufficientEncryption"
	case CBATTErrorInsufficientEncryptionKeySize:
		return "CBATTErrorInsufficientEncryptionKeySize"
	case CBATTErrorInsufficientResources:
		return "CBATTErrorInsufficientResources"
	case CBATTErrorInvalidAttributeValueLength:
		return "CBATTErrorInvalidAttributeValueLength"
	case CBATTErrorInvalidHandle:
		return "CBATTErrorInvalidHandle"
	case CBATTErrorInvalidOffset:
		return "CBATTErrorInvalidOffset"
	case CBATTErrorInvalidPdu:
		return "CBATTErrorInvalidPdu"
	case CBATTErrorPrepareQueueFull:
		return "CBATTErrorPrepareQueueFull"
	case CBATTErrorReadNotPermitted:
		return "CBATTErrorReadNotPermitted"
	case CBATTErrorRequestNotSupported:
		return "CBATTErrorRequestNotSupported"
	case CBATTErrorSuccess:
		return "CBATTErrorSuccess"
	case CBATTErrorUnlikelyError:
		return "CBATTErrorUnlikelyError"
	case CBATTErrorUnsupportedGroupType:
		return "CBATTErrorUnsupportedGroupType"
	case CBATTErrorWriteNotPermitted:
		return "CBATTErrorWriteNotPermitted"
	default:
		return fmt.Sprintf("CBATTError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions
type CBAttributePermissions uint

const (
	// CBAttributePermissionsReadEncryptionRequired: A permission that indicates only trusted devices can read the attribute’s value.
	CBAttributePermissionsReadEncryptionRequired CBAttributePermissions = 0x4
	// CBAttributePermissionsReadable: A permission that indicates a peripheral can read the attribute’s value.
	CBAttributePermissionsReadable CBAttributePermissions = 0x1
	// CBAttributePermissionsWriteEncryptionRequired: A permission that indicates only trusted devices can write the attribute’s value.
	CBAttributePermissionsWriteEncryptionRequired CBAttributePermissions = 0x8
	// CBAttributePermissionsWriteable: A permission that indicates a peripheral can write the attribute’s value.
	CBAttributePermissionsWriteable CBAttributePermissions = 0x2
)

func (e CBAttributePermissions) String() string {
	switch e {
	case CBAttributePermissionsReadEncryptionRequired:
		return "CBAttributePermissionsReadEncryptionRequired"
	case CBAttributePermissionsReadable:
		return "CBAttributePermissionsReadable"
	case CBAttributePermissionsWriteEncryptionRequired:
		return "CBAttributePermissionsWriteEncryptionRequired"
	case CBAttributePermissionsWriteable:
		return "CBAttributePermissionsWriteable"
	default:
		return fmt.Sprintf("CBAttributePermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/Feature
type CBCentralManagerFeature uint

const (
	// CBCentralManagerFeatureExtendedScanAndConnect: The hardware supports extended scans and enhanced connection creation.
	CBCentralManagerFeatureExtendedScanAndConnect CBCentralManagerFeature = 1
)

func (e CBCentralManagerFeature) String() string {
	switch e {
	case CBCentralManagerFeatureExtendedScanAndConnect:
		return "CBCentralManagerFeatureExtendedScanAndConnect"
	default:
		return fmt.Sprintf("CBCentralManagerFeature(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState
type CBCentralManagerState int

const (
	// CBCentralManagerStatePoweredOff: A state that indicates Bluetooth is currently powered off.
	CBCentralManagerStatePoweredOff CBCentralManagerState = 4
	// CBCentralManagerStatePoweredOn: A state that indicates Bluetooth is currently powered on and available to use.
	CBCentralManagerStatePoweredOn CBCentralManagerState = 5
	// CBCentralManagerStateResetting: A state that indicates the connection with the system service was momentarily lost.
	CBCentralManagerStateResetting CBCentralManagerState = 1
	// CBCentralManagerStateUnauthorized: A state that indicates the application isn’t authorized to use the Bluetooth low energy role.
	CBCentralManagerStateUnauthorized CBCentralManagerState = 3
	// CBCentralManagerStateUnknown: The manager’s state is unknown.
	CBCentralManagerStateUnknown CBCentralManagerState = 0
	// CBCentralManagerStateUnsupported: A state that indicates this device doesn’t support the Bluetooth low energy central or client role.
	CBCentralManagerStateUnsupported CBCentralManagerState = 2
)

func (e CBCentralManagerState) String() string {
	switch e {
	case CBCentralManagerStatePoweredOff:
		return "CBCentralManagerStatePoweredOff"
	case CBCentralManagerStatePoweredOn:
		return "CBCentralManagerStatePoweredOn"
	case CBCentralManagerStateResetting:
		return "CBCentralManagerStateResetting"
	case CBCentralManagerStateUnauthorized:
		return "CBCentralManagerStateUnauthorized"
	case CBCentralManagerStateUnknown:
		return "CBCentralManagerStateUnknown"
	case CBCentralManagerStateUnsupported:
		return "CBCentralManagerStateUnsupported"
	default:
		return fmt.Sprintf("CBCentralManagerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBChannelSoundingSessionConfiguration/Role-swift.enum
type CBChannelSoundingSessionConfigurationRole int

const ()

// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicProperties
type CBCharacteristicProperties uint

const (
	// CBCharacteristicPropertyAuthenticatedSignedWrites: A property that indicates the perhipheral allows signed writes of the characteristic’s value, without a response to indicate the write succeeded.
	CBCharacteristicPropertyAuthenticatedSignedWrites CBCharacteristicProperties = 0x40
	// CBCharacteristicPropertyBroadcast: A property that indicates the characteristic can broadcast its value using a characteristic configuration descriptor.
	CBCharacteristicPropertyBroadcast CBCharacteristicProperties = 0x1
	// CBCharacteristicPropertyExtendedProperties: A property that indicates the characteristic defines additional properties in the extended properties descriptor.
	CBCharacteristicPropertyExtendedProperties CBCharacteristicProperties = 0x80
	// CBCharacteristicPropertyIndicate: A property that indicates the peripheral permits notifications of the characteristic’s value, with a response from the central to indicate receipt of the notification.
	CBCharacteristicPropertyIndicate CBCharacteristicProperties = 0x20
	// CBCharacteristicPropertyIndicateEncryptionRequired: A property that indicates only trusted devices can enable indications of the characteristic’s value.
	CBCharacteristicPropertyIndicateEncryptionRequired CBCharacteristicProperties = 0x200
	// CBCharacteristicPropertyNotify: A property that indicates the peripheral permits notifications of the characteristic’s value, without a response from the central to indicate receipt of the notification.
	CBCharacteristicPropertyNotify CBCharacteristicProperties = 0x10
	// CBCharacteristicPropertyNotifyEncryptionRequired: A property that indicates that only trusted devices can enable notifications of the characteristic’s value.
	CBCharacteristicPropertyNotifyEncryptionRequired CBCharacteristicProperties = 0x100
	// CBCharacteristicPropertyRead: A property that indicates a peripheral can read the characteristic’s value.
	CBCharacteristicPropertyRead CBCharacteristicProperties = 0x2
	// CBCharacteristicPropertyWrite: A property that indicates a peripheral can write the characteristic’s value, with a response to indicate that the write succeeded.
	CBCharacteristicPropertyWrite CBCharacteristicProperties = 0x8
	// CBCharacteristicPropertyWriteWithoutResponse: A property that indicates a peripheral can write the characteristic’s value, without a response to indicate that the write succeeded.
	CBCharacteristicPropertyWriteWithoutResponse CBCharacteristicProperties = 0x4
)

func (e CBCharacteristicProperties) String() string {
	switch e {
	case CBCharacteristicPropertyAuthenticatedSignedWrites:
		return "CBCharacteristicPropertyAuthenticatedSignedWrites"
	case CBCharacteristicPropertyBroadcast:
		return "CBCharacteristicPropertyBroadcast"
	case CBCharacteristicPropertyExtendedProperties:
		return "CBCharacteristicPropertyExtendedProperties"
	case CBCharacteristicPropertyIndicate:
		return "CBCharacteristicPropertyIndicate"
	case CBCharacteristicPropertyIndicateEncryptionRequired:
		return "CBCharacteristicPropertyIndicateEncryptionRequired"
	case CBCharacteristicPropertyNotify:
		return "CBCharacteristicPropertyNotify"
	case CBCharacteristicPropertyNotifyEncryptionRequired:
		return "CBCharacteristicPropertyNotifyEncryptionRequired"
	case CBCharacteristicPropertyRead:
		return "CBCharacteristicPropertyRead"
	case CBCharacteristicPropertyWrite:
		return "CBCharacteristicPropertyWrite"
	case CBCharacteristicPropertyWriteWithoutResponse:
		return "CBCharacteristicPropertyWriteWithoutResponse"
	default:
		return fmt.Sprintf("CBCharacteristicProperties(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType
type CBCharacteristicWriteType int

const (
	// CBCharacteristicWriteWithResponse: Write a characteristic value, with a response from the peripheral to indicate whether the write was successful.
	CBCharacteristicWriteWithResponse CBCharacteristicWriteType = 0
	// CBCharacteristicWriteWithoutResponse: Write a characteristic value, without any response from the peripheral to indicate whether the write was successful.
	CBCharacteristicWriteWithoutResponse CBCharacteristicWriteType = 1
)

func (e CBCharacteristicWriteType) String() string {
	switch e {
	case CBCharacteristicWriteWithResponse:
		return "CBCharacteristicWriteWithResponse"
	case CBCharacteristicWriteWithoutResponse:
		return "CBCharacteristicWriteWithoutResponse"
	default:
		return fmt.Sprintf("CBCharacteristicWriteType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectionEvent
type CBConnectionEvent int

const (
	// CBConnectionEventPeerConnected: The peer has connected to the local device.
	CBConnectionEventPeerConnected CBConnectionEvent = 1
	// CBConnectionEventPeerDisconnected: The peer has disconnected from the local device.
	CBConnectionEventPeerDisconnected CBConnectionEvent = 0
)

func (e CBConnectionEvent) String() string {
	switch e {
	case CBConnectionEventPeerConnected:
		return "CBConnectionEventPeerConnected"
	case CBConnectionEventPeerDisconnected:
		return "CBConnectionEventPeerDisconnected"
	default:
		return fmt.Sprintf("CBConnectionEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code
type CBError int

const (
	// CBErrorAlreadyAdvertising: The peripheral is already advertising.
	CBErrorAlreadyAdvertising CBError = 9
	// CBErrorConnectionFailed: The connection failed.
	CBErrorConnectionFailed CBError = 10
	// CBErrorConnectionLimitReached: The device already has the maximum number of connections.
	CBErrorConnectionLimitReached CBError = 11
	// CBErrorConnectionTimeout: The connection timed out.
	CBErrorConnectionTimeout  CBError = 6
	CBErrorEncryptionTimedOut CBError = 15
	// CBErrorInvalidHandle: The specified attribute handle is invalid.
	CBErrorInvalidHandle CBError = 2
	// CBErrorInvalidParameters: The specified parameters are invalid.
	CBErrorInvalidParameters                         CBError = 1
	CBErrorLeGattExceededBackgroundNotificationLimit CBError = 17
	CBErrorLeGattNearBackgroundNotificationLimit     CBError = 18
	// CBErrorNotConnected: The device isn’t currently connected.
	CBErrorNotConnected CBError = 3
	// CBErrorOperationCancelled: The error represents a canceled operation.
	CBErrorOperationCancelled CBError = 5
	// CBErrorOperationNotSupported: The operation isn’t supported.
	CBErrorOperationNotSupported CBError = 13
	// CBErrorOutOfSpace: The device has run out of space to complete the intended operation.
	CBErrorOutOfSpace                    CBError = 4
	CBErrorPeerRemovedPairingInformation CBError = 14
	// CBErrorPeripheralDisconnected: The peripheral disconnected.
	CBErrorPeripheralDisconnected CBError = 7
	CBErrorTooManyLEPairedDevices CBError = 16
	// CBErrorUUIDNotAllowed: The specified UUID isn’t permitted.
	CBErrorUUIDNotAllowed CBError = 8
	// CBErrorUnknown: An unknown error occurred.
	CBErrorUnknown       CBError = 0
	CBErrorUnknownDevice CBError = 12
	// CBErrorUnkownDevice: A misspelled version of the unknown device error code.
	CBErrorUnkownDevice CBError = 12
)

func (e CBError) String() string {
	switch e {
	case CBErrorAlreadyAdvertising:
		return "CBErrorAlreadyAdvertising"
	case CBErrorConnectionFailed:
		return "CBErrorConnectionFailed"
	case CBErrorConnectionLimitReached:
		return "CBErrorConnectionLimitReached"
	case CBErrorConnectionTimeout:
		return "CBErrorConnectionTimeout"
	case CBErrorEncryptionTimedOut:
		return "CBErrorEncryptionTimedOut"
	case CBErrorInvalidHandle:
		return "CBErrorInvalidHandle"
	case CBErrorInvalidParameters:
		return "CBErrorInvalidParameters"
	case CBErrorLeGattExceededBackgroundNotificationLimit:
		return "CBErrorLeGattExceededBackgroundNotificationLimit"
	case CBErrorLeGattNearBackgroundNotificationLimit:
		return "CBErrorLeGattNearBackgroundNotificationLimit"
	case CBErrorNotConnected:
		return "CBErrorNotConnected"
	case CBErrorOperationCancelled:
		return "CBErrorOperationCancelled"
	case CBErrorOperationNotSupported:
		return "CBErrorOperationNotSupported"
	case CBErrorOutOfSpace:
		return "CBErrorOutOfSpace"
	case CBErrorPeerRemovedPairingInformation:
		return "CBErrorPeerRemovedPairingInformation"
	case CBErrorPeripheralDisconnected:
		return "CBErrorPeripheralDisconnected"
	case CBErrorTooManyLEPairedDevices:
		return "CBErrorTooManyLEPairedDevices"
	case CBErrorUUIDNotAllowed:
		return "CBErrorUUIDNotAllowed"
	case CBErrorUnknown:
		return "CBErrorUnknown"
	case CBErrorUnknownDevice:
		return "CBErrorUnknownDevice"
	default:
		return fmt.Sprintf("CBError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBManagerAuthorization
type CBManagerAuthorization int

const (
	// CBManagerAuthorizationAllowedAlways: A state that indicates the user has authorized Bluetooth at any time.
	CBManagerAuthorizationAllowedAlways CBManagerAuthorization = 3
	// CBManagerAuthorizationDenied: A state that indicates the user explicitly denied Bluetooth access for this app.
	CBManagerAuthorizationDenied CBManagerAuthorization = 2
	// CBManagerAuthorizationNotDetermined: A state that indicates the user has yet to authorize Bluetooth for this app.
	CBManagerAuthorizationNotDetermined CBManagerAuthorization = 0
	// CBManagerAuthorizationRestricted: A state that indicates this app isn’t authorized to use Bluetooth.
	CBManagerAuthorizationRestricted CBManagerAuthorization = 1
)

func (e CBManagerAuthorization) String() string {
	switch e {
	case CBManagerAuthorizationAllowedAlways:
		return "CBManagerAuthorizationAllowedAlways"
	case CBManagerAuthorizationDenied:
		return "CBManagerAuthorizationDenied"
	case CBManagerAuthorizationNotDetermined:
		return "CBManagerAuthorizationNotDetermined"
	case CBManagerAuthorizationRestricted:
		return "CBManagerAuthorizationRestricted"
	default:
		return fmt.Sprintf("CBManagerAuthorization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState
type CBManagerState int

const (
	// CBManagerStatePoweredOff: A state that indicates Bluetooth is currently powered off.
	CBManagerStatePoweredOff CBManagerState = 4
	// CBManagerStatePoweredOn: A state that indicates Bluetooth is currently powered on and available to use.
	CBManagerStatePoweredOn CBManagerState = 5
	// CBManagerStateResetting: A state that indicates the connection with the system service was momentarily lost.
	CBManagerStateResetting CBManagerState = 1
	// CBManagerStateUnauthorized: A state that indicates the application isn’t authorized to use the Bluetooth low energy role.
	CBManagerStateUnauthorized CBManagerState = 3
	// CBManagerStateUnknown: The manager’s state is unknown.
	CBManagerStateUnknown CBManagerState = 0
	// CBManagerStateUnsupported: A state that indicates this device doesn’t support the Bluetooth low energy central or client role.
	CBManagerStateUnsupported CBManagerState = 2
)

func (e CBManagerState) String() string {
	switch e {
	case CBManagerStatePoweredOff:
		return "CBManagerStatePoweredOff"
	case CBManagerStatePoweredOn:
		return "CBManagerStatePoweredOn"
	case CBManagerStateResetting:
		return "CBManagerStateResetting"
	case CBManagerStateUnauthorized:
		return "CBManagerStateUnauthorized"
	case CBManagerStateUnknown:
		return "CBManagerStateUnknown"
	case CBManagerStateUnsupported:
		return "CBManagerStateUnsupported"
	default:
		return fmt.Sprintf("CBManagerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerAuthorizationStatus
type CBPeripheralManagerAuthorizationStatus int

const (
	// CBPeripheralManagerAuthorizationStatusAuthorized: An authorization status that indicates the user authorized this app to share data using Bluetooth while in the background.
	CBPeripheralManagerAuthorizationStatusAuthorized CBPeripheralManagerAuthorizationStatus = 3
	// CBPeripheralManagerAuthorizationStatusDenied: An authorization status that indicates the user explicitly denied this app from sharing data using Bluetooth while in the background.
	CBPeripheralManagerAuthorizationStatusDenied CBPeripheralManagerAuthorizationStatus = 2
	// CBPeripheralManagerAuthorizationStatusNotDetermined: An authorization status that indicates the user hasn’t indicated whether this app can share data using Bluetooth while in the background.
	CBPeripheralManagerAuthorizationStatusNotDetermined CBPeripheralManagerAuthorizationStatus = 0
	// CBPeripheralManagerAuthorizationStatusRestricted: An authorization status that indicates this app isn’t authorized to share data using Bluetooth while in the background.
	CBPeripheralManagerAuthorizationStatusRestricted CBPeripheralManagerAuthorizationStatus = 1
)

func (e CBPeripheralManagerAuthorizationStatus) String() string {
	switch e {
	case CBPeripheralManagerAuthorizationStatusAuthorized:
		return "CBPeripheralManagerAuthorizationStatusAuthorized"
	case CBPeripheralManagerAuthorizationStatusDenied:
		return "CBPeripheralManagerAuthorizationStatusDenied"
	case CBPeripheralManagerAuthorizationStatusNotDetermined:
		return "CBPeripheralManagerAuthorizationStatusNotDetermined"
	case CBPeripheralManagerAuthorizationStatusRestricted:
		return "CBPeripheralManagerAuthorizationStatusRestricted"
	default:
		return fmt.Sprintf("CBPeripheralManagerAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerConnectionLatency
type CBPeripheralManagerConnectionLatency int

const (
	// CBPeripheralManagerConnectionLatencyHigh: A latency setting that prioritizes extending battery life over rapid communication.
	CBPeripheralManagerConnectionLatencyHigh CBPeripheralManagerConnectionLatency = 2
	// CBPeripheralManagerConnectionLatencyLow: A latency setting indicating that prioritizes rapid communication over battery life.
	CBPeripheralManagerConnectionLatencyLow CBPeripheralManagerConnectionLatency = 0
	// CBPeripheralManagerConnectionLatencyMedium: A latency setting that balances communication frequency and battery life.
	CBPeripheralManagerConnectionLatencyMedium CBPeripheralManagerConnectionLatency = 1
)

func (e CBPeripheralManagerConnectionLatency) String() string {
	switch e {
	case CBPeripheralManagerConnectionLatencyHigh:
		return "CBPeripheralManagerConnectionLatencyHigh"
	case CBPeripheralManagerConnectionLatencyLow:
		return "CBPeripheralManagerConnectionLatencyLow"
	case CBPeripheralManagerConnectionLatencyMedium:
		return "CBPeripheralManagerConnectionLatencyMedium"
	default:
		return fmt.Sprintf("CBPeripheralManagerConnectionLatency(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState
type CBPeripheralManagerState int

const (
	// CBPeripheralManagerStatePoweredOff: A manager state that indicates Bluetooth is currently powered off.
	CBPeripheralManagerStatePoweredOff CBPeripheralManagerState = 4
	// CBPeripheralManagerStatePoweredOn: A manager state that indicates Bluetooth is currently powered on and is available to use.
	CBPeripheralManagerStatePoweredOn CBPeripheralManagerState = 5
	// CBPeripheralManagerStateResetting: A manager state that indicates the connection with the system service was momentarily lost.
	CBPeripheralManagerStateResetting CBPeripheralManagerState = 1
	// CBPeripheralManagerStateUnauthorized: A manager state that indicates the app isn’t authorized to use the Bluetooth low energy peripheral/server role.
	CBPeripheralManagerStateUnauthorized CBPeripheralManagerState = 3
	// CBPeripheralManagerStateUnknown: A manager state that indicates the current state of the peripheral manager is unknown.
	CBPeripheralManagerStateUnknown CBPeripheralManagerState = 0
	// CBPeripheralManagerStateUnsupported: A manager state that indicates the platform doesn’t support the Bluetooth low energy peripheral/server role.
	CBPeripheralManagerStateUnsupported CBPeripheralManagerState = 2
)

func (e CBPeripheralManagerState) String() string {
	switch e {
	case CBPeripheralManagerStatePoweredOff:
		return "CBPeripheralManagerStatePoweredOff"
	case CBPeripheralManagerStatePoweredOn:
		return "CBPeripheralManagerStatePoweredOn"
	case CBPeripheralManagerStateResetting:
		return "CBPeripheralManagerStateResetting"
	case CBPeripheralManagerStateUnauthorized:
		return "CBPeripheralManagerStateUnauthorized"
	case CBPeripheralManagerStateUnknown:
		return "CBPeripheralManagerStateUnknown"
	case CBPeripheralManagerStateUnsupported:
		return "CBPeripheralManagerStateUnsupported"
	default:
		return fmt.Sprintf("CBPeripheralManagerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralState
type CBPeripheralState int

const (
	// CBPeripheralStateConnected: The peripheral is connected to the central manager.
	CBPeripheralStateConnected CBPeripheralState = 2
	// CBPeripheralStateConnecting: The peripheral is in the process of connecting to the central manager.
	CBPeripheralStateConnecting CBPeripheralState = 1
	// CBPeripheralStateDisconnected: The peripheral isn’t connected to the central manager.
	CBPeripheralStateDisconnected CBPeripheralState = 0
	// CBPeripheralStateDisconnecting: The peripheral is disconnecting from the central manager.
	CBPeripheralStateDisconnecting CBPeripheralState = 3
)

func (e CBPeripheralState) String() string {
	switch e {
	case CBPeripheralStateConnected:
		return "CBPeripheralStateConnected"
	case CBPeripheralStateConnecting:
		return "CBPeripheralStateConnecting"
	case CBPeripheralStateDisconnected:
		return "CBPeripheralStateDisconnected"
	case CBPeripheralStateDisconnecting:
		return "CBPeripheralStateDisconnecting"
	default:
		return fmt.Sprintf("CBPeripheralState(%d)", e)
	}
}
