// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// The class instance for the [CWInterface] class.
var (
	_CWInterfaceClass     CWInterfaceClass
	_CWInterfaceClassOnce sync.Once
)

func getCWInterfaceClass() CWInterfaceClass {
	_CWInterfaceClassOnce.Do(func() {
		_CWInterfaceClass = CWInterfaceClass{class: objc.GetClass("CWInterface")}
	})
	return _CWInterfaceClass
}

// GetCWInterfaceClass returns the class object for CWInterface.
func GetCWInterfaceClass() CWInterfaceClass {
	return getCWInterfaceClass()
}

type CWInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWInterfaceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWInterfaceClass) Alloc() CWInterface {
	rv := objc.Send[CWInterface](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates an IEEE 802.11 interface.
//
// # Overview
//
// Provides access to various WLAN interface parameters, and operations such
// as scanning for networks, association, and creating computer-to-computer
// (ad-hoc) networks.
//
// # Setting interface parameters
//
//   - [CWInterface.SetPairwiseMasterKeyError]: Sets the interface pairwise primary key (PMK).
//   - [CWInterface.SetPowerError]: Sets the interface power state.
//   - [CWInterface.SetWEPKeyFlagsIndexError]: Sets the interface WEP key.
//   - [CWInterface.SetWLANChannelError]: Sets the interface channel.
//
// # Scanning for networks
//
//   - [CWInterface.ScanForNetworksWithNameError]: Scans for networks.
//   - [CWInterface.ScanForNetworksWithSSIDError]: Scans for networks.
//
// # Disassociating from a network
//
//   - [CWInterface.Disassociate]: Disassociates from the current network.
//
// # Committing a configuration
//
//   - [CWInterface.CommitConfigurationAuthorizationError]: Commit a configuration for the given WLAN interface.
//
// # Associating to a network
//
//   - [CWInterface.AssociateToEnterpriseNetworkIdentityUsernamePasswordError]: Connects to the given enterprise network.
//   - [CWInterface.AssociateToNetworkPasswordError]: Associates to a given network using the given network passphrase.
//
// # Instance Properties
//
//   - [CWInterface.InterfaceName]: The BSD name of the interface.
//
// # Instance Methods
//
//   - [CWInterface.ActivePHYMode]: The current active PHY modes for the interface.
//   - [CWInterface.Bssid]: The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
//   - [CWInterface.CachedScanResults]: The networks currently in the scan cache for the WLAN interface.
//   - [CWInterface.Configuration]: The current configuration for the given WLAN interface.
//   - [CWInterface.CountryCode]: The current country code (ISO/IEC 3166-1:1997) for the interface.
//   - [CWInterface.HardwareAddress]: The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
//   - [CWInterface.InterfaceMode]: The current mode for the interface.
//   - [CWInterface.NoiseMeasurement]: The current aggregate noise measurement (dBm) for the interface.
//   - [CWInterface.PowerOn]: The interface power state is set to “ON”.
//   - [CWInterface.RssiValue]: The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
//   - [CWInterface.ScanForNetworksWithNameIncludeHiddenError]: Scans for networks with the name you specify, optionally including hidden networks.
//   - [CWInterface.ScanForNetworksWithSSIDIncludeHiddenError]: Scans for networks with the SSID you specify, optionally including hidden networks.
//   - [CWInterface.Security]: The current security mode for the interface.
//   - [CWInterface.ServiceActive]: The interface has its corresponding network service enabled.
//   - [CWInterface.Ssid]: The current service set identifier (SSID) for the interface, encoded as a string.
//   - [CWInterface.SsidData]: The current service set identifier (SSID) for the interface, returned as data.
//   - [CWInterface.SupportedWLANChannels]: An array of channels supported by the interface for the active country code.
//   - [CWInterface.TransmitPower]: The current transmit power (mW) for the interface.
//   - [CWInterface.TransmitRate]: The current transmit rate (Mbps) for the interface.
//   - [CWInterface.WlanChannel]: The current channel for the interface.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface
type CWInterface struct {
	objectivec.Object
}

// CWInterfaceFromID constructs a [CWInterface] from an objc.ID.
//
// Encapsulates an IEEE 802.11 interface.
func CWInterfaceFromID(id objc.ID) CWInterface {
	return CWInterface{objectivec.Object{ID: id}}
}

// NOTE: CWInterface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWInterface] class.
//
// # Setting interface parameters
//
//   - [ICWInterface.SetPairwiseMasterKeyError]: Sets the interface pairwise primary key (PMK).
//   - [ICWInterface.SetPowerError]: Sets the interface power state.
//   - [ICWInterface.SetWEPKeyFlagsIndexError]: Sets the interface WEP key.
//   - [ICWInterface.SetWLANChannelError]: Sets the interface channel.
//
// # Scanning for networks
//
//   - [ICWInterface.ScanForNetworksWithNameError]: Scans for networks.
//   - [ICWInterface.ScanForNetworksWithSSIDError]: Scans for networks.
//
// # Disassociating from a network
//
//   - [ICWInterface.Disassociate]: Disassociates from the current network.
//
// # Committing a configuration
//
//   - [ICWInterface.CommitConfigurationAuthorizationError]: Commit a configuration for the given WLAN interface.
//
// # Associating to a network
//
//   - [ICWInterface.AssociateToEnterpriseNetworkIdentityUsernamePasswordError]: Connects to the given enterprise network.
//   - [ICWInterface.AssociateToNetworkPasswordError]: Associates to a given network using the given network passphrase.
//
// # Instance Properties
//
//   - [ICWInterface.InterfaceName]: The BSD name of the interface.
//
// # Instance Methods
//
//   - [ICWInterface.ActivePHYMode]: The current active PHY modes for the interface.
//   - [ICWInterface.Bssid]: The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
//   - [ICWInterface.CachedScanResults]: The networks currently in the scan cache for the WLAN interface.
//   - [ICWInterface.Configuration]: The current configuration for the given WLAN interface.
//   - [ICWInterface.CountryCode]: The current country code (ISO/IEC 3166-1:1997) for the interface.
//   - [ICWInterface.HardwareAddress]: The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
//   - [ICWInterface.InterfaceMode]: The current mode for the interface.
//   - [ICWInterface.NoiseMeasurement]: The current aggregate noise measurement (dBm) for the interface.
//   - [ICWInterface.PowerOn]: The interface power state is set to “ON”.
//   - [ICWInterface.RssiValue]: The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
//   - [ICWInterface.ScanForNetworksWithNameIncludeHiddenError]: Scans for networks with the name you specify, optionally including hidden networks.
//   - [ICWInterface.ScanForNetworksWithSSIDIncludeHiddenError]: Scans for networks with the SSID you specify, optionally including hidden networks.
//   - [ICWInterface.Security]: The current security mode for the interface.
//   - [ICWInterface.ServiceActive]: The interface has its corresponding network service enabled.
//   - [ICWInterface.Ssid]: The current service set identifier (SSID) for the interface, encoded as a string.
//   - [ICWInterface.SsidData]: The current service set identifier (SSID) for the interface, returned as data.
//   - [ICWInterface.SupportedWLANChannels]: An array of channels supported by the interface for the active country code.
//   - [ICWInterface.TransmitPower]: The current transmit power (mW) for the interface.
//   - [ICWInterface.TransmitRate]: The current transmit rate (Mbps) for the interface.
//   - [ICWInterface.WlanChannel]: The current channel for the interface.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface
type ICWInterface interface {
	objectivec.IObject

	// Topic: Setting interface parameters

	// Sets the interface pairwise primary key (PMK).
	SetPairwiseMasterKeyError(key foundation.NSData) (bool, error)
	// Sets the interface power state.
	SetPowerError(power bool) (bool, error)
	// Sets the interface WEP key.
	SetWEPKeyFlagsIndexError(key foundation.NSData, flags CWCipherKeyFlags, index int) (bool, error)
	// Sets the interface channel.
	SetWLANChannelError(channel ICWChannel) (bool, error)

	// Topic: Scanning for networks

	// Scans for networks.
	ScanForNetworksWithNameError(networkName string) (foundation.INSSet, error)
	// Scans for networks.
	ScanForNetworksWithSSIDError(ssid foundation.NSData) (foundation.INSSet, error)

	// Topic: Disassociating from a network

	// Disassociates from the current network.
	Disassociate()

	// Topic: Committing a configuration

	// Commit a configuration for the given WLAN interface.
	CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization objectivec.IObject) (bool, error)

	// Topic: Associating to a network

	// Connects to the given enterprise network.
	AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity security.SecIdentityRef, username string, password string) (bool, error)
	// Associates to a given network using the given network passphrase.
	AssociateToNetworkPasswordError(network ICWNetwork, password string) (bool, error)

	// Topic: Instance Properties

	// The BSD name of the interface.
	InterfaceName() string

	// Topic: Instance Methods

	// The current active PHY modes for the interface.
	ActivePHYMode() CWPHYMode
	// The current basic service set identifier (BSSID) for the interface, returned as a UTF-8 string.
	Bssid() string
	// The networks currently in the scan cache for the WLAN interface.
	CachedScanResults() foundation.INSSet
	// The current configuration for the given WLAN interface.
	Configuration() ICWConfiguration
	// The current country code (ISO/IEC 3166-1:1997) for the interface.
	CountryCode() string
	// The hardware media access control (MAC) address for the interface, returned as a UTF-8 string.
	HardwareAddress() string
	// The current mode for the interface.
	InterfaceMode() CWInterfaceMode
	// The current aggregate noise measurement (dBm) for the interface.
	NoiseMeasurement() int
	// The interface power state is set to “ON”.
	PowerOn() bool
	// The current aggregate received signal strength indication (RSSI) measurement (dBm) for the interface.
	RssiValue() int
	// Scans for networks with the name you specify, optionally including hidden networks.
	ScanForNetworksWithNameIncludeHiddenError(networkName string, includeHidden bool) (foundation.INSSet, error)
	// Scans for networks with the SSID you specify, optionally including hidden networks.
	ScanForNetworksWithSSIDIncludeHiddenError(ssid foundation.NSData, includeHidden bool) (foundation.INSSet, error)
	// The current security mode for the interface.
	Security() CWSecurity
	// The interface has its corresponding network service enabled.
	ServiceActive() bool
	// The current service set identifier (SSID) for the interface, encoded as a string.
	Ssid() string
	// The current service set identifier (SSID) for the interface, returned as data.
	SsidData() foundation.NSData
	// An array of channels supported by the interface for the active country code.
	SupportedWLANChannels() foundation.INSSet
	// The current transmit power (mW) for the interface.
	TransmitPower() int
	// The current transmit rate (Mbps) for the interface.
	TransmitRate() float64
	// The current channel for the interface.
	WlanChannel() ICWChannel
}

// Init initializes the instance.
func (c CWInterface) Init() CWInterface {
	rv := objc.Send[CWInterface](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWInterface) Autorelease() CWInterface {
	rv := objc.Send[CWInterface](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWInterface creates a new CWInterface instance.
func NewCWInterface() CWInterface {
	class := getCWInterfaceClass()
	rv := objc.Send[CWInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the interface pairwise primary key (PMK).
//
// key: An NSData object containing the pairwise primary key (PMK).
//
// # Discussion
//
// key must be 32 octets. If key is nil, this method clears the PMK for the
// interface.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPairwiseMasterKey(_:)
func (c CWInterface) SetPairwiseMasterKeyError(key foundation.NSData) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("setPairwiseMasterKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPairwiseMasterKey:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the interface power state.
//
// power: A Boolean value corresponding to the power state. NO indicates the
// “OFF” state.
//
// # Discussion
//
// This operation may require an administrator password.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setPower(_:)
func (c CWInterface) SetPowerError(power bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("setPower:error:"), power, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPower:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the interface WEP key.
//
// key: An NSData object containing the WEP key.
//
// flags: An NSUInteger indicating which cipher key flags to use for the specified
// key.
//
// index: An NSUInteger indicating which default key index to use for the specified
// key.
//
// # Discussion
//
// key must be 5 octets for WEP-40 or 13 octets for WEP-104. If key is nil,
// this method clears the WEP key for the interface. index must correspond to
// default key index 1-4.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWEPKey(_:flags:index:)
func (c CWInterface) SetWEPKeyFlagsIndexError(key foundation.NSData, flags CWCipherKeyFlags, index int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("setWEPKey:flags:index:error:"), key, flags, index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWEPKey:flags:index:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the interface channel.
//
// channel: A CWChannel object corresponding to the channel.
//
// # Discussion
//
// The channel cannot be changed if the interface is associated to a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/setWLANChannel(_:)
func (c CWInterface) SetWLANChannelError(channel ICWChannel) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("setWLANChannel:error:"), channel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWLANChannel:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Scans for networks.
//
// networkName: The name (SSID) of the network for which to scan.
//
// # Return Value
//
// A set of CWNetwork objects.
//
// # Discussion
//
// If ssid parameter is present, a directed scan will be performed by the
// interface, otherwise a broadcast scan will be performed. This method will
// block for the duration of the scan.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:)
func (c CWInterface) ScanForNetworksWithNameError(networkName string) (foundation.INSSet, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scanForNetworksWithName:error:"), objc.String(networkName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSSetFromID(rv), nil

}

// Scans for networks.
//
// ssid: The SSID for which to scan.
//
// # Return Value
//
// A set of CWNetwork objects.
//
// # Discussion
//
// If ssid parameter is present, a directed scan will be performed by the
// interface, otherwise a broadcast scan will be performed. This method will
// block for the duration of the scan.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:)
func (c CWInterface) ScanForNetworksWithSSIDError(ssid foundation.NSData) (foundation.INSSet, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scanForNetworksWithSSID:error:"), ssid, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSSetFromID(rv), nil

}

// Disassociates from the current network.
//
// # Discussion
//
// This method has no effect if the interface is not associated to a network.
// This operation may require an administrator password.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/disassociate()
func (c CWInterface) Disassociate() {
	objc.Send[objc.ID](c.ID, objc.Sel("disassociate"))
}

// Commit a configuration for the given WLAN interface.
//
// configuration: The configuration to commit.
//
// authorization: An SFAuthorization object to use for authorizing the commit. This parameter
// is optional and can be passed as nil.
//
// authorization is a [*securityfoundation.SFAuthorization].
//
// # Discussion
//
// This method requires the caller have root privileges or obtain
// administrator privileges with the authorization parameter.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/commitConfiguration(_:authorization:)
func (c CWInterface) CommitConfigurationAuthorizationError(configuration ICWConfiguration, authorization objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("commitConfiguration:authorization:error:"), configuration, authorization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("commitConfiguration:authorization:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Connects to the given enterprise network.
//
// network: The network to which the interface will associate.
//
// identity: The identity to use for IEEE 802.1X authentication. Holds the corresponding
// client certificate.
//
// username: The username to use for IEEE 802.1X authentication.
//
// password: The password to use for IEEE 802.1X authentication.
//
// # Discussion
//
// This method will block for the duration of the association. This operation
// may require an administrator password.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(toEnterpriseNetwork:identity:username:password:)
func (c CWInterface) AssociateToEnterpriseNetworkIdentityUsernamePasswordError(network ICWNetwork, identity security.SecIdentityRef, username string, password string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("associateToEnterpriseNetwork:identity:username:password:error:"), network, identity, objc.String(username), objc.String(password), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("associateToEnterpriseNetwork:identity:username:password:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Associates to a given network using the given network passphrase.
//
// network: The network to which the interface will associate.
//
// password: The network passphrase or key. Required for association to WEP, WPA
// Personal, and WPA2 Personal networks.
//
// # Discussion
//
// This method will block for the duration of the association. This operation
// may require an administrator password.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/associate(to:password:)
func (c CWInterface) AssociateToNetworkPasswordError(network ICWNetwork, password string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("associateToNetwork:password:error:"), network, objc.String(password), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("associateToNetwork:password:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The current active PHY modes for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current active PHY mode. Returns
// kCWPHYModeNone in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/activePHYMode()
func (c CWInterface) ActivePHYMode() CWPHYMode {
	rv := objc.Send[CWPHYMode](c.ID, objc.Sel("activePHYMode"))
	return CWPHYMode(rv)
}

// The current basic service set identifier (BSSID) for the interface,
// returned as a UTF-8 string.
//
// # Discussion
//
// Dynamically queries the interface for the current BSSID. Returns a UTF-8
// string formatted as , or nil in the case of an error, or if the interface
// is not participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/bssid()
func (c CWInterface) Bssid() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("bssid"))
	return foundation.NSStringFromID(rv).String()
}

// The networks currently in the scan cache for the WLAN interface.
//
// # Discussion
//
// Returns nil in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/cachedScanResults()
func (c CWInterface) CachedScanResults() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cachedScanResults"))
	return foundation.NSSetFromID(rv)
}

// The current configuration for the given WLAN interface.
//
// # Discussion
//
// Returns nil in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/configuration()
func (c CWInterface) Configuration() ICWConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return CWConfigurationFromID(rv)
}

// The current country code (ISO/IEC 3166-1:1997) for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current country code. Returns nil
// in the case of an error, or if the interface is OFF.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/countryCode()
func (c CWInterface) CountryCode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("countryCode"))
	return foundation.NSStringFromID(rv).String()
}

// The hardware media access control (MAC) address for the interface, returned
// as a UTF-8 string.
//
// # Discussion
//
// The standard format for printing a MAC-48 address is used to represent the
// MAC address as a string. Returns nil in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/hardwareAddress()
func (c CWInterface) HardwareAddress() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("hardwareAddress"))
	return foundation.NSStringFromID(rv).String()
}

// The current mode for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current mode. Returns
// kCWInterfaceModeNone in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceMode()
func (c CWInterface) InterfaceMode() CWInterfaceMode {
	rv := objc.Send[CWInterfaceMode](c.ID, objc.Sel("interfaceMode"))
	return CWInterfaceMode(rv)
}

// The current aggregate noise measurement (dBm) for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current aggregate noise
// measurement. Returns 0 in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/noiseMeasurement()
func (c CWInterface) NoiseMeasurement() int {
	rv := objc.Send[int](c.ID, objc.Sel("noiseMeasurement"))
	return rv
}

// The interface power state is set to “ON”.
//
// # Discussion
//
// Dynamically queries the interface for the current power state. Returns NO
// in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/powerOn()
func (c CWInterface) PowerOn() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("powerOn"))
	return rv
}

// The current aggregate received signal strength indication (RSSI)
// measurement (dBm) for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current aggregate RSSI
// measurement. Returns 0 in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/rssiValue()
func (c CWInterface) RssiValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("rssiValue"))
	return rv
}

// Scans for networks with the name you specify, optionally including hidden
// networks.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withName:includeHidden:)
func (c CWInterface) ScanForNetworksWithNameIncludeHiddenError(networkName string, includeHidden bool) (foundation.INSSet, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scanForNetworksWithName:includeHidden:error:"), objc.String(networkName), includeHidden, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSSetFromID(rv), nil

}

// Scans for networks with the SSID you specify, optionally including hidden
// networks.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/scanForNetworks(withSSID:includeHidden:)
func (c CWInterface) ScanForNetworksWithSSIDIncludeHiddenError(ssid foundation.NSData, includeHidden bool) (foundation.INSSet, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scanForNetworksWithSSID:includeHidden:error:"), ssid, includeHidden, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSSetFromID(rv), nil

}

// The current security mode for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the security mode. Returns
// kCWSecurityUnknown in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/security()
func (c CWInterface) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c.ID, objc.Sel("security"))
	return CWSecurity(rv)
}

// The interface has its corresponding network service enabled.
//
// # Discussion
//
// Returns NO in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/serviceActive()
func (c CWInterface) ServiceActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("serviceActive"))
	return rv
}

// The current service set identifier (SSID) for the interface, encoded as a
// string.
//
// # Discussion
//
// Dynamically queries the interface for the current SSID. Returns nil in the
// case of an error, or if the interface is not participating in a network, or
// if the SSID can not be encoded as a valid UTF-8 or WinLatin1 string.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssid()
func (c CWInterface) Ssid() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssid"))
	return foundation.NSStringFromID(rv).String()
}

// The current service set identifier (SSID) for the interface, returned as
// data.
//
// # Discussion
//
// Dynamically queries the interface for the current SSID. The SSID is 1-32
// octets. Returns nil in the case of an error, or if the interface is not
// participating in a network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/ssidData()
func (c CWInterface) SsidData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssidData"))
	return foundation.NSDataFromID(rv)
}

// An array of channels supported by the interface for the active country
// code.
//
// # Discussion
//
// Dynamically queries the interface for the supported channels. Returns an
// array of CWChannel objects, or nil in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/supportedWLANChannels()
func (c CWInterface) SupportedWLANChannels() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportedWLANChannels"))
	return foundation.NSSetFromID(rv)
}

// The current transmit power (mW) for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current transmit power. Returns 0
// in the case of an error.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitPower()
func (c CWInterface) TransmitPower() int {
	rv := objc.Send[int](c.ID, objc.Sel("transmitPower"))
	return rv
}

// The current transmit rate (Mbps) for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current transmit rate. Returns 0
// in the case of an error, or if the interface is not participating in a
// network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/transmitRate()
func (c CWInterface) TransmitRate() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("transmitRate"))
	return rv
}

// The current channel for the interface.
//
// # Discussion
//
// Dynamically queries the interface for the current channel. Returns nil in
// the case of an error, or if the interface is not participating in a
// network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/wlanChannel()
func (c CWInterface) WlanChannel() ICWChannel {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("wlanChannel"))
	return CWChannelFromID(rv)
}

// The BSD name of the interface.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWInterface/interfaceName
func (c CWInterface) InterfaceName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("interfaceName"))
	return foundation.NSStringFromID(rv).String()
}
