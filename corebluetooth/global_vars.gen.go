// Code generated from Apple documentation. DO NOT EDIT.

package corebluetooth

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// CBATTErrorDomain is the domain for Core Bluetooth ATT errors.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBATTErrorDomain
	CBATTErrorDomain string
	// CBAdvertisementDataIsConnectable is a Boolean value that indicates whether the advertising event type is connectable.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataIsConnectable
	CBAdvertisementDataIsConnectable string
	// CBAdvertisementDataLocalNameKey is the local name of a peripheral.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataLocalNameKey
	CBAdvertisementDataLocalNameKey string
	// CBAdvertisementDataManufacturerDataKey is the manufacturer data of a peripheral.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataManufacturerDataKey
	CBAdvertisementDataManufacturerDataKey string
	// CBAdvertisementDataOverflowServiceUUIDsKey is an array of UUIDs found in the overflow area of the advertisement data.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataOverflowServiceUUIDsKey
	CBAdvertisementDataOverflowServiceUUIDsKey string
	// CBAdvertisementDataServiceDataKey is a dictionary that contains service-specific advertisement data.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataServiceDataKey
	CBAdvertisementDataServiceDataKey string
	// CBAdvertisementDataServiceUUIDsKey is an array of service UUIDs.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataServiceUUIDsKey
	CBAdvertisementDataServiceUUIDsKey string
	// CBAdvertisementDataSolicitedServiceUUIDsKey is an array of solicited service UUIDs.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataSolicitedServiceUUIDsKey
	CBAdvertisementDataSolicitedServiceUUIDsKey string
	// CBAdvertisementDataTxPowerLevelKey is the transmit power of a peripheral.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBAdvertisementDataTxPowerLevelKey
	CBAdvertisementDataTxPowerLevelKey string
	// CBCentralManagerOptionRestoreIdentifierKey is a string containing a unique identifier (UID) for the central manager to instantiate.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerOptionRestoreIdentifierKey
	CBCentralManagerOptionRestoreIdentifierKey string
	// CBCentralManagerOptionShowPowerAlertKey is a Boolean value that specifies whether the system warns the user if the app instantiates the central manager when Bluetooth service isn’t available.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerOptionShowPowerAlertKey
	CBCentralManagerOptionShowPowerAlertKey string
	// CBCentralManagerRestoredStatePeripheralsKey is an array of peripherals for use when restoring the state of a central manager.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerRestoredStatePeripheralsKey
	CBCentralManagerRestoredStatePeripheralsKey string
	// CBCentralManagerRestoredStateScanOptionsKey is a dictionary of peripheral scan options for use when restoring state.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerRestoredStateScanOptionsKey
	CBCentralManagerRestoredStateScanOptionsKey string
	// CBCentralManagerRestoredStateScanServicesKey is an array of service IDs for use when restoring state.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerRestoredStateScanServicesKey
	CBCentralManagerRestoredStateScanServicesKey string
	// CBCentralManagerScanOptionAllowDuplicatesKey is a Boolean value that specifies whether the scan should run without duplicate filtering.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerScanOptionAllowDuplicatesKey
	CBCentralManagerScanOptionAllowDuplicatesKey string
	// CBCentralManagerScanOptionSolicitedServiceUUIDsKey is an array of service UUIDs that you want to scan for.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerScanOptionSolicitedServiceUUIDsKey
	CBCentralManagerScanOptionSolicitedServiceUUIDsKey string
	// CBConnectPeripheralOptionEnableAutoReconnect is a Boolean value that specifies whether the system automatically reconnects with a peripheral.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectPeripheralOptionEnableAutoReconnect
	CBConnectPeripheralOptionEnableAutoReconnect string
	// CBConnectPeripheralOptionNotifyOnConnectionKey is a Boolean value that specifies whether the system should display an alert when connecting a peripheral in the background.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectPeripheralOptionNotifyOnConnectionKey
	CBConnectPeripheralOptionNotifyOnConnectionKey string
	// CBConnectPeripheralOptionNotifyOnDisconnectionKey is a Boolean value that specifies whether the system should display an alert when disconnecting a peripheral in the background.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectPeripheralOptionNotifyOnDisconnectionKey
	CBConnectPeripheralOptionNotifyOnDisconnectionKey string
	// CBConnectPeripheralOptionNotifyOnNotificationKey is a Boolean value that specifies whether the system should display an alert for any notification sent by a peripheral.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectPeripheralOptionNotifyOnNotificationKey
	CBConnectPeripheralOptionNotifyOnNotificationKey string
	// CBConnectPeripheralOptionStartDelayKey is an option that indicates a delay before the system makes a connection.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBConnectPeripheralOptionStartDelayKey
	CBConnectPeripheralOptionStartDelayKey string
	// CBErrorDomain is the domain for Core Bluetooth errors.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBErrorDomain
	CBErrorDomain string
	// CBPeripheralManagerOptionRestoreIdentifierKey is a string containing a unique identifier (UID) for the peripheral manager to instantiate.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerOptionRestoreIdentifierKey
	CBPeripheralManagerOptionRestoreIdentifierKey string
	// CBPeripheralManagerOptionShowPowerAlertKey is a Boolean value specifying whether the system should warn if Bluetooth is in the powered-off state when instantiating the peripheral manager.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerOptionShowPowerAlertKey
	CBPeripheralManagerOptionShowPowerAlertKey string
	// CBPeripheralManagerRestoredStateAdvertisementDataKey is a dictionary of restored advertising data.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerRestoredStateAdvertisementDataKey
	CBPeripheralManagerRestoredStateAdvertisementDataKey string
	// CBPeripheralManagerRestoredStateServicesKey is an array of restored peripheral services.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerRestoredStateServicesKey
	CBPeripheralManagerRestoredStateServicesKey string
	// CBUUIDCharacteristicAggregateFormatString is the UUID for the Aggregate Format descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicAggregateFormatString
	CBUUIDCharacteristicAggregateFormatString string
	// CBUUIDCharacteristicExtendedPropertiesString is the UUID for the Extended Properties descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicExtendedPropertiesString
	CBUUIDCharacteristicExtendedPropertiesString string
	// CBUUIDCharacteristicFormatString is the UUID for the Presentation Format descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicFormatString
	CBUUIDCharacteristicFormatString string
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicObservationScheduleString
	CBUUIDCharacteristicObservationScheduleString string
	// CBUUIDCharacteristicUserDescriptionString is the UUID for the User Description descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicUserDescriptionString
	CBUUIDCharacteristicUserDescriptionString string
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDCharacteristicValidRangeString
	CBUUIDCharacteristicValidRangeString string
	// CBUUIDClientCharacteristicConfigurationString is the UUID for the Client Configuration descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDClientCharacteristicConfigurationString
	CBUUIDClientCharacteristicConfigurationString string
	// CBUUIDL2CAPPSMCharacteristicString is the PSM of an L2CAP channel associated with the GATT service containing this characteristic.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDL2CAPPSMCharacteristicString
	CBUUIDL2CAPPSMCharacteristicString string
	// CBUUIDServerCharacteristicConfigurationString is the UUID for the Server Configuration descriptor, as a string.
	//
	// See: https://developer.apple.com/documentation/CoreBluetooth/CBUUIDServerCharacteristicConfigurationString
	CBUUIDServerCharacteristicConfigurationString string
)

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBATTErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBATTErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataIsConnectable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataIsConnectable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataLocalNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataLocalNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataManufacturerDataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataManufacturerDataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataOverflowServiceUUIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataOverflowServiceUUIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataServiceDataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataServiceDataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataServiceUUIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataServiceUUIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataSolicitedServiceUUIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataSolicitedServiceUUIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBAdvertisementDataTxPowerLevelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBAdvertisementDataTxPowerLevelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerOptionRestoreIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerOptionRestoreIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerOptionShowPowerAlertKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerOptionShowPowerAlertKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerRestoredStatePeripheralsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerRestoredStatePeripheralsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerRestoredStateScanOptionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerRestoredStateScanOptionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerRestoredStateScanServicesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerRestoredStateScanServicesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerScanOptionAllowDuplicatesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerScanOptionAllowDuplicatesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBCentralManagerScanOptionSolicitedServiceUUIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBCentralManagerScanOptionSolicitedServiceUUIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectPeripheralOptionEnableAutoReconnect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectPeripheralOptionEnableAutoReconnect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectPeripheralOptionNotifyOnConnectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectPeripheralOptionNotifyOnConnectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectPeripheralOptionNotifyOnDisconnectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectPeripheralOptionNotifyOnDisconnectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectPeripheralOptionNotifyOnNotificationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectPeripheralOptionNotifyOnNotificationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectPeripheralOptionStartDelayKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectPeripheralOptionStartDelayKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectionEventMatchingOptionPeripheralUUIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectionEventMatchingOptions.PeripheralUUIDs = CBConnectionEventMatchingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBConnectionEventMatchingOptionServiceUUIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBConnectionEventMatchingOptions.ServiceUUIDs = CBConnectionEventMatchingOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBPeripheralManagerOptionRestoreIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBPeripheralManagerOptionRestoreIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBPeripheralManagerOptionShowPowerAlertKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBPeripheralManagerOptionShowPowerAlertKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBPeripheralManagerRestoredStateAdvertisementDataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBPeripheralManagerRestoredStateAdvertisementDataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBPeripheralManagerRestoredStateServicesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBPeripheralManagerRestoredStateServicesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicAggregateFormatString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicAggregateFormatString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicExtendedPropertiesString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicExtendedPropertiesString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicFormatString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicFormatString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicObservationScheduleString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicObservationScheduleString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicUserDescriptionString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicUserDescriptionString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDCharacteristicValidRangeString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDCharacteristicValidRangeString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDClientCharacteristicConfigurationString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDClientCharacteristicConfigurationString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDL2CAPPSMCharacteristicString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDL2CAPPSMCharacteristicString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CBUUIDServerCharacteristicConfigurationString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CBUUIDServerCharacteristicConfigurationString = objc.GoString(cstr)
			}
		}
	}

}

// CBConnectionEventMatchingOptions provides typed accessors for [CBConnectionEventMatchingOption] constants.
var CBConnectionEventMatchingOptions struct {
	// PeripheralUUIDs: An array of UUID objects that represents peripherals to match.
	PeripheralUUIDs CBConnectionEventMatchingOption
	// ServiceUUIDs: An array that represents service identifiers to match.
	ServiceUUIDs CBConnectionEventMatchingOption
}
