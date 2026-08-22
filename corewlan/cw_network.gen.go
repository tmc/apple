// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CWNetwork] class.
var (
	_CWNetworkClass     CWNetworkClass
	_CWNetworkClassOnce sync.Once
)

func getCWNetworkClass() CWNetworkClass {
	_CWNetworkClassOnce.Do(func() {
		_CWNetworkClass = CWNetworkClass{class: objc.GetClass("CWNetwork")}
	})
	return _CWNetworkClass
}

// GetCWNetworkClass returns the class object for CWNetwork.
func GetCWNetworkClass() CWNetworkClass {
	return getCWNetworkClass()
}

type CWNetworkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWNetworkClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWNetworkClass) Alloc() CWNetwork {
	rv := objc.Send[CWNetwork](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates an IEEE 802.11 network, providing read-only accessors to
// various properties of the network.
//
// # Getting supported security types
//
//   - [CWNetwork.SupportsSecurity]: Method for determining which security types a network supports.
//
// # Getting supported PHY modes
//
//   - [CWNetwork.SupportsPHYMode]: Method for determining which PHY modes a network supports.
//
// # Comparing wireless networks
//
//   - [CWNetwork.IsEqualToNetwork]: Method for determining CWNetwork object equality.
//
// # Instance Properties
//
//   - [CWNetwork.BeaconInterval]: The beacon interval (ms) for the network.
//   - [CWNetwork.Bssid]: The basic service set identifier (BSSID) for the network, returned as UTF-8 string.
//   - [CWNetwork.CountryCode]: The country code (ISO/IEC 3166-1:1997) for the network.
//   - [CWNetwork.Ibss]: The network is an IBSS network.
//   - [CWNetwork.InformationElementData]: Information element data included in beacon or probe response frames.
//   - [CWNetwork.NoiseMeasurement]: The aggregate noise measurement (dBm) for the network.
//   - [CWNetwork.RssiValue]: The aggregate received signal strength indication (RSSI) measurement (dBm) for the network.
//   - [CWNetwork.Ssid]: The service set identifier (SSID) for the network, encoded as a string.
//   - [CWNetwork.SsidData]: The service set identifier (SSID) for the network, returned as data.
//   - [CWNetwork.WlanChannel]: The channel for the network.
//
// # Initializers
//
//   - [CWNetwork.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork
type CWNetwork struct {
	objectivec.Object
}

// CWNetworkFromID constructs a [CWNetwork] from an objc.ID.
//
// Encapsulates an IEEE 802.11 network, providing read-only accessors to
// various properties of the network.
func CWNetworkFromID(id objc.ID) CWNetwork {
	return CWNetwork{objectivec.Object{ID: id}}
}

// NOTE: CWNetwork adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWNetwork] class.
//
// # Getting supported security types
//
//   - [ICWNetwork.SupportsSecurity]: Method for determining which security types a network supports.
//
// # Getting supported PHY modes
//
//   - [ICWNetwork.SupportsPHYMode]: Method for determining which PHY modes a network supports.
//
// # Comparing wireless networks
//
//   - [ICWNetwork.IsEqualToNetwork]: Method for determining CWNetwork object equality.
//
// # Instance Properties
//
//   - [ICWNetwork.BeaconInterval]: The beacon interval (ms) for the network.
//   - [ICWNetwork.Bssid]: The basic service set identifier (BSSID) for the network, returned as UTF-8 string.
//   - [ICWNetwork.CountryCode]: The country code (ISO/IEC 3166-1:1997) for the network.
//   - [ICWNetwork.Ibss]: The network is an IBSS network.
//   - [ICWNetwork.InformationElementData]: Information element data included in beacon or probe response frames.
//   - [ICWNetwork.NoiseMeasurement]: The aggregate noise measurement (dBm) for the network.
//   - [ICWNetwork.RssiValue]: The aggregate received signal strength indication (RSSI) measurement (dBm) for the network.
//   - [ICWNetwork.Ssid]: The service set identifier (SSID) for the network, encoded as a string.
//   - [ICWNetwork.SsidData]: The service set identifier (SSID) for the network, returned as data.
//   - [ICWNetwork.WlanChannel]: The channel for the network.
//
// # Initializers
//
//   - [ICWNetwork.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork
type ICWNetwork interface {
	objectivec.IObject

	// Topic: Getting supported security types

	// Method for determining which security types a network supports.
	SupportsSecurity(security CWSecurity) bool

	// Topic: Getting supported PHY modes

	// Method for determining which PHY modes a network supports.
	SupportsPHYMode(phyMode CWPHYMode) bool

	// Topic: Comparing wireless networks

	// Method for determining CWNetwork object equality.
	IsEqualToNetwork(network ICWNetwork) bool

	// Topic: Instance Properties

	// The beacon interval (ms) for the network.
	BeaconInterval() int
	// The basic service set identifier (BSSID) for the network, returned as UTF-8 string.
	Bssid() string
	// The country code (ISO/IEC 3166-1:1997) for the network.
	CountryCode() string
	// The network is an IBSS network.
	Ibss() bool
	// Information element data included in beacon or probe response frames.
	InformationElementData() foundation.NSData
	// The aggregate noise measurement (dBm) for the network.
	NoiseMeasurement() int
	// The aggregate received signal strength indication (RSSI) measurement (dBm) for the network.
	RssiValue() int
	// The service set identifier (SSID) for the network, encoded as a string.
	Ssid() string
	// The service set identifier (SSID) for the network, returned as data.
	SsidData() foundation.NSData
	// The channel for the network.
	WlanChannel() ICWChannel

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CWNetwork

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CWNetwork) Init() CWNetwork {
	rv := objc.Send[CWNetwork](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWNetwork) Autorelease() CWNetwork {
	rv := objc.Send[CWNetwork](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWNetwork creates a new CWNetwork instance.
func NewCWNetwork() CWNetwork {
	class := getCWNetworkClass()
	rv := objc.Send[CWNetwork](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/init(coder:)
func NewCWNetworkWithCoder(coder foundation.INSCoder) CWNetwork {
	instance := getCWNetworkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWNetworkFromID(rv)
}

// Method for determining which security types a network supports.
//
// security: The security type.
//
// # Return Value
//
// YES if the network supports the specified security type.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsSecurity(_:)
func (c CWNetwork) SupportsSecurity(security CWSecurity) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsSecurity:"), security)
	return rv
}

// Method for determining which PHY modes a network supports.
//
// # Return Value
//
// YES if the network supports the specified PHY mode.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsPHYMode(_:)
func (c CWNetwork) SupportsPHYMode(phyMode CWPHYMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsPHYMode:"), phyMode)
	return rv
}

// Method for determining CWNetwork object equality.
//
// network: The CWNetwork object for which to test equality.
//
// # Return Value
//
// YES if the objects are equal.
//
// # Discussion
//
// Two CWNetwork objects are considered equal if their corresponding ssidData,
// securityType, and networkType properties are equal.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/isEqual(to:)
func (c CWNetwork) IsEqualToNetwork(network ICWNetwork) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEqualToNetwork:"), network)
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/init(coder:)
func (c CWNetwork) InitWithCoder(coder foundation.INSCoder) CWNetwork {
	rv := objc.Send[CWNetwork](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CWNetwork) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The beacon interval (ms) for the network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/beaconInterval
func (c CWNetwork) BeaconInterval() int {
	rv := objc.Send[int](c.ID, objc.Sel("beaconInterval"))
	return rv
}

// The basic service set identifier (BSSID) for the network, returned as UTF-8
// string.
//
// # Discussion
//
// Returns a UTF-8 string formatted as .
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/bssid
func (c CWNetwork) Bssid() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("bssid"))
	return foundation.NSStringFromID(rv).String()
}

// The country code (ISO/IEC 3166-1:1997) for the network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/countryCode
func (c CWNetwork) CountryCode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("countryCode"))
	return foundation.NSStringFromID(rv).String()
}

// The network is an IBSS network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ibss
func (c CWNetwork) Ibss() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("ibss"))
	return rv
}

// Information element data included in beacon or probe response frames.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/informationElementData
func (c CWNetwork) InformationElementData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("informationElementData"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The aggregate noise measurement (dBm) for the network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/noiseMeasurement
func (c CWNetwork) NoiseMeasurement() int {
	rv := objc.Send[int](c.ID, objc.Sel("noiseMeasurement"))
	return rv
}

// The aggregate received signal strength indication (RSSI) measurement (dBm)
// for the network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/rssiValue
func (c CWNetwork) RssiValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("rssiValue"))
	return rv
}

// The service set identifier (SSID) for the network, encoded as a string.
//
// # Discussion
//
// If the SSID can not be encoded as a valid UTF-8 or WinLatin1 string, this
// method returns nil.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssid
func (c CWNetwork) Ssid() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssid"))
	return foundation.NSStringFromID(rv).String()
}

// The service set identifier (SSID) for the network, returned as data.
//
// # Discussion
//
// The SSID is defined as 1-32 octets.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssidData
func (c CWNetwork) SsidData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssidData"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The channel for the network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/wlanChannel
func (c CWNetwork) WlanChannel() ICWChannel {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("wlanChannel"))
	return CWChannelFromID(objc.ID(rv))
}
