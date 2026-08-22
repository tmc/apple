// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreWLAN/CWChannelBand
type CWChannelBand int

const (
	// KCWChannelBand2GHz: 2.4GHz channel band.
	KCWChannelBand2GHz CWChannelBand = 1
	// KCWChannelBand5GHz: 5GHz channel band.
	KCWChannelBand5GHz CWChannelBand = 2
	// KCWChannelBandUnknown: Unknown channel band.
	KCWChannelBandUnknown CWChannelBand = 0
)

func (e CWChannelBand) String() string {
	switch e {
	case KCWChannelBand2GHz:
		return "KCWChannelBand2GHz"
	case KCWChannelBand5GHz:
		return "KCWChannelBand5GHz"
	case KCWChannelBandUnknown:
		return "KCWChannelBandUnknown"
	default:
		return fmt.Sprintf("CWChannelBand(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWChannelWidth
type CWChannelWidth int

const (
	// KCWChannelWidth160MHz: 160MHz channel width.
	KCWChannelWidth160MHz CWChannelWidth = 4
	// KCWChannelWidth20MHz: 20MHz channel width.
	KCWChannelWidth20MHz CWChannelWidth = 1
	// KCWChannelWidth40MHz: 40MHz channel width.
	KCWChannelWidth40MHz CWChannelWidth = 2
	// KCWChannelWidth80MHz: 80MHz channel width.
	KCWChannelWidth80MHz CWChannelWidth = 3
	// KCWChannelWidthUnknown: Unknown channel width.
	KCWChannelWidthUnknown CWChannelWidth = 0
)

func (e CWChannelWidth) String() string {
	switch e {
	case KCWChannelWidth160MHz:
		return "KCWChannelWidth160MHz"
	case KCWChannelWidth20MHz:
		return "KCWChannelWidth20MHz"
	case KCWChannelWidth40MHz:
		return "KCWChannelWidth40MHz"
	case KCWChannelWidth80MHz:
		return "KCWChannelWidth80MHz"
	case KCWChannelWidthUnknown:
		return "KCWChannelWidthUnknown"
	default:
		return fmt.Sprintf("CWChannelWidth(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWCipherKeyFlags
type CWCipherKeyFlags uint

const (
	// KCWCipherKeyFlagsMulticast: A flag that indicates to use the cipher key for multicast packets.
	KCWCipherKeyFlagsMulticast CWCipherKeyFlags = 4
	// KCWCipherKeyFlagsNone: Open System authentication.
	KCWCipherKeyFlagsNone CWCipherKeyFlags = 0
	// KCWCipherKeyFlagsRx: A flag that indicates to use the cipher key for packets received by the interface.
	KCWCipherKeyFlagsRx CWCipherKeyFlags = 16
	// KCWCipherKeyFlagsTx: A flag that indicates to use the cipher key for packets sent from the interface.
	KCWCipherKeyFlagsTx CWCipherKeyFlags = 8
	// KCWCipherKeyFlagsUnicast: A flag that indicates to use the cipher key for unicast packets.
	KCWCipherKeyFlagsUnicast CWCipherKeyFlags = 2
)

func (e CWCipherKeyFlags) String() string {
	switch e {
	case KCWCipherKeyFlagsMulticast:
		return "KCWCipherKeyFlagsMulticast"
	case KCWCipherKeyFlagsNone:
		return "KCWCipherKeyFlagsNone"
	case KCWCipherKeyFlagsRx:
		return "KCWCipherKeyFlagsRx"
	case KCWCipherKeyFlagsTx:
		return "KCWCipherKeyFlagsTx"
	case KCWCipherKeyFlagsUnicast:
		return "KCWCipherKeyFlagsUnicast"
	default:
		return fmt.Sprintf("CWCipherKeyFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWErr
type CWErr int

const ()

// See: https://developer.apple.com/documentation/CoreWLAN/CWEventType
type CWEventType int

const (
	// CWEventTypeBSSIDDidChange: Posts when the current BSSID of any Wi-Fi interface changes.
	CWEventTypeBSSIDDidChange CWEventType = 3
	CWEventTypeBtCoexStats    CWEventType = 9
	// CWEventTypeCountryCodeDidChange: Posts when the adopted country code of any Wi-Fi interface changes.
	CWEventTypeCountryCodeDidChange CWEventType = 4
	// CWEventTypeLinkDidChange: Posts when the link state for any Wi-Fi interface changes.
	CWEventTypeLinkDidChange CWEventType = 5
	// CWEventTypeLinkQualityDidChange: Posts when the RSSI or transmit rate for any Wi-Fi interface changes.
	CWEventTypeLinkQualityDidChange CWEventType = 6
	// CWEventTypeModeDidChange: Posts when the operating mode of any Wi-Fi interface changes.
	CWEventTypeModeDidChange CWEventType = 7
	// CWEventTypeNone: No specified event type.
	CWEventTypeNone CWEventType = 0
	// CWEventTypePowerDidChange: Posts when the power state of any Wi-Fi interface changes.
	CWEventTypePowerDidChange CWEventType = 1
	// CWEventTypeSSIDDidChange: Posts when the current SSID of any Wi-Fi interface changes.
	CWEventTypeSSIDDidChange CWEventType = 2
	// CWEventTypeScanCacheUpdated: Posts when the scan cache of any Wi-Fi interface is updated with new scan results.
	CWEventTypeScanCacheUpdated CWEventType = 8
	// CWEventTypeUnknown: Unknown event type.
	CWEventTypeUnknown CWEventType = 9223372036854775807
)

func (e CWEventType) String() string {
	switch e {
	case CWEventTypeBSSIDDidChange:
		return "CWEventTypeBSSIDDidChange"
	case CWEventTypeBtCoexStats:
		return "CWEventTypeBtCoexStats"
	case CWEventTypeCountryCodeDidChange:
		return "CWEventTypeCountryCodeDidChange"
	case CWEventTypeLinkDidChange:
		return "CWEventTypeLinkDidChange"
	case CWEventTypeLinkQualityDidChange:
		return "CWEventTypeLinkQualityDidChange"
	case CWEventTypeModeDidChange:
		return "CWEventTypeModeDidChange"
	case CWEventTypeNone:
		return "CWEventTypeNone"
	case CWEventTypePowerDidChange:
		return "CWEventTypePowerDidChange"
	case CWEventTypeSSIDDidChange:
		return "CWEventTypeSSIDDidChange"
	case CWEventTypeScanCacheUpdated:
		return "CWEventTypeScanCacheUpdated"
	case CWEventTypeUnknown:
		return "CWEventTypeUnknown"
	default:
		return fmt.Sprintf("CWEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWIBSSModeSecurity
type CWIBSSModeSecurity int

const (
	// KCWIBSSModeSecurityNone: Open System authentication.
	KCWIBSSModeSecurityNone CWIBSSModeSecurity = 0
	// KCWIBSSModeSecurityWEP104: WPA Personal authentication.
	KCWIBSSModeSecurityWEP104 CWIBSSModeSecurity = 2
	// KCWIBSSModeSecurityWEP40: WEP security.
	KCWIBSSModeSecurityWEP40 CWIBSSModeSecurity = 1
)

func (e CWIBSSModeSecurity) String() string {
	switch e {
	case KCWIBSSModeSecurityNone:
		return "KCWIBSSModeSecurityNone"
	case KCWIBSSModeSecurityWEP104:
		return "KCWIBSSModeSecurityWEP104"
	case KCWIBSSModeSecurityWEP40:
		return "KCWIBSSModeSecurityWEP40"
	default:
		return fmt.Sprintf("CWIBSSModeSecurity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWInterfaceMode
type CWInterfaceMode int

const (
	// KCWInterfaceModeHostAP: Interface is participating in an infrastructure network as an access point.
	KCWInterfaceModeHostAP CWInterfaceMode = 3
	// KCWInterfaceModeIBSS: Interface is participating in an IBSS network.
	KCWInterfaceModeIBSS CWInterfaceMode = 2
	// KCWInterfaceModeNone: Interface is not in any mode.
	KCWInterfaceModeNone CWInterfaceMode = 0
	// KCWInterfaceModeStation: Interface is participating in an infrastructure network as a non-AP station.
	KCWInterfaceModeStation CWInterfaceMode = 1
)

func (e CWInterfaceMode) String() string {
	switch e {
	case KCWInterfaceModeHostAP:
		return "KCWInterfaceModeHostAP"
	case KCWInterfaceModeIBSS:
		return "KCWInterfaceModeIBSS"
	case KCWInterfaceModeNone:
		return "KCWInterfaceModeNone"
	case KCWInterfaceModeStation:
		return "KCWInterfaceModeStation"
	default:
		return fmt.Sprintf("CWInterfaceMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDomain
type CWKeychainDomain int

const (
	// KCWKeychainDomainNone: No keychain domain specified.
	KCWKeychainDomainNone CWKeychainDomain = 0
	// KCWKeychainDomainSystem: The system keychain domain.
	KCWKeychainDomainSystem CWKeychainDomain = 2
	// KCWKeychainDomainUser: The user keychain domain.
	KCWKeychainDomainUser CWKeychainDomain = 1
)

func (e CWKeychainDomain) String() string {
	switch e {
	case KCWKeychainDomainNone:
		return "KCWKeychainDomainNone"
	case KCWKeychainDomainSystem:
		return "KCWKeychainDomainSystem"
	case KCWKeychainDomainUser:
		return "KCWKeychainDomainUser"
	default:
		return fmt.Sprintf("CWKeychainDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWPHYMode
type CWPHYMode int

const (
	// KCWPHYMode11a: IEEE 802.11a PHY.
	KCWPHYMode11a CWPHYMode = 1
	// KCWPHYMode11ac: IEEE 802.11ac PHY.
	KCWPHYMode11ac CWPHYMode = 5
	// KCWPHYMode11b: IEEE 802.11b PHY.
	KCWPHYMode11b CWPHYMode = 2
	// KCWPHYMode11g: IEEE 802.11g PHY.
	KCWPHYMode11g CWPHYMode = 3
	// KCWPHYMode11n: IEEE 802.11n PHY.
	KCWPHYMode11n CWPHYMode = 4
	// KCWPHYModeNone: No specified mode.
	KCWPHYModeNone CWPHYMode = 0
)

func (e CWPHYMode) String() string {
	switch e {
	case KCWPHYMode11a:
		return "KCWPHYMode11a"
	case KCWPHYMode11ac:
		return "KCWPHYMode11ac"
	case KCWPHYMode11b:
		return "KCWPHYMode11b"
	case KCWPHYMode11g:
		return "KCWPHYMode11g"
	case KCWPHYMode11n:
		return "KCWPHYMode11n"
	case KCWPHYModeNone:
		return "KCWPHYModeNone"
	default:
		return fmt.Sprintf("CWPHYMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWSecurity
type CWSecurity int

const (
	// KCWSecurityDynamicWEP: Dynamic WEP security.
	KCWSecurityDynamicWEP CWSecurity = 6
	// KCWSecurityEnterprise: Enterprise authentication.
	KCWSecurityEnterprise CWSecurity = 10
	// KCWSecurityNone: Open System authentication.
	KCWSecurityNone CWSecurity = 0
	// KCWSecurityPersonal: Personal authentication.
	KCWSecurityPersonal CWSecurity = 5
	// KCWSecurityUnknown: Unknown security type.
	KCWSecurityUnknown CWSecurity = 9223372036854775807
	// KCWSecurityWEP: WEP security.
	KCWSecurityWEP CWSecurity = 1
	// KCWSecurityWPA2Enterprise: WPA2 Enterprise authentication.
	KCWSecurityWPA2Enterprise CWSecurity = 9
	// KCWSecurityWPA2Personal: WPA2 Personal authentication.
	KCWSecurityWPA2Personal CWSecurity = 4
	// KCWSecurityWPA3Enterprise: WPA3 Enterprise authentication.
	KCWSecurityWPA3Enterprise CWSecurity = 12
	// KCWSecurityWPA3Personal: WPA3 Personal authentication.
	KCWSecurityWPA3Personal CWSecurity = 11
	// KCWSecurityWPA3Transition: WPA3 Transition (WPA3/WPA2 Personal) authentication.
	KCWSecurityWPA3Transition CWSecurity = 13
	// KCWSecurityWPAEnterprise: WPA Enterprise authentication.
	KCWSecurityWPAEnterprise CWSecurity = 7
	// KCWSecurityWPAEnterpriseMixed: WPA/WPA2 Enterprise authentication.
	KCWSecurityWPAEnterpriseMixed CWSecurity = 8
	// KCWSecurityWPAPersonal: WPA Personal authentication.
	KCWSecurityWPAPersonal CWSecurity = 2
	// KCWSecurityWPAPersonalMixed: WPA/WPA2 Personal authentication.
	KCWSecurityWPAPersonalMixed CWSecurity = 3
)

func (e CWSecurity) String() string {
	switch e {
	case KCWSecurityDynamicWEP:
		return "KCWSecurityDynamicWEP"
	case KCWSecurityEnterprise:
		return "KCWSecurityEnterprise"
	case KCWSecurityNone:
		return "KCWSecurityNone"
	case KCWSecurityPersonal:
		return "KCWSecurityPersonal"
	case KCWSecurityUnknown:
		return "KCWSecurityUnknown"
	case KCWSecurityWEP:
		return "KCWSecurityWEP"
	case KCWSecurityWPA2Enterprise:
		return "KCWSecurityWPA2Enterprise"
	case KCWSecurityWPA2Personal:
		return "KCWSecurityWPA2Personal"
	case KCWSecurityWPA3Enterprise:
		return "KCWSecurityWPA3Enterprise"
	case KCWSecurityWPA3Personal:
		return "KCWSecurityWPA3Personal"
	case KCWSecurityWPA3Transition:
		return "KCWSecurityWPA3Transition"
	case KCWSecurityWPAEnterprise:
		return "KCWSecurityWPAEnterprise"
	case KCWSecurityWPAEnterpriseMixed:
		return "KCWSecurityWPAEnterpriseMixed"
	case KCWSecurityWPAPersonal:
		return "KCWSecurityWPAPersonal"
	case KCWSecurityWPAPersonalMixed:
		return "KCWSecurityWPAPersonalMixed"
	default:
		return fmt.Sprintf("CWSecurity(%d)", e)
	}
}
