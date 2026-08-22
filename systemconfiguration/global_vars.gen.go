// Code generated from Apple documentation. DO NOT EDIT.

package systemconfiguration

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// KCFErrorDomainSystemConfiguration is a string constant identifying a Core Foundation error domain. See [CFError] for further information on error domains.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kCFErrorDomainSystemConfiguration
	KCFErrorDomainSystemConfiguration string
	// KSCBondStatusDeviceAggregationStatus is device is aggregating. See [Ethernet Bond Aggregation Status] for a list of possible values.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCBondStatusDeviceAggregationStatus
	KSCBondStatusDeviceAggregationStatus string
	// KSCBondStatusDeviceCollecting is can be `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCBondStatusDeviceCollecting
	KSCBondStatusDeviceCollecting string
	// KSCBondStatusDeviceDistributing is can be `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCBondStatusDeviceDistributing
	KSCBondStatusDeviceDistributing string
	// KSCCompAnyRegex is a regular expression pattern that matches any component.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompAnyRegex-swift.var
	KSCCompAnyRegex string
	// KSCCompGlobal is the Component key [Global].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompGlobal-swift.var
	KSCCompGlobal string
	// KSCCompHostNames is the Component key [HostNames].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompHostNames-swift.var
	KSCCompHostNames string
	// KSCCompInterface is the Component key [Interface].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompInterface-swift.var
	KSCCompInterface string
	// KSCCompNetwork is the Component key [Network].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompNetwork-swift.var
	KSCCompNetwork string
	// KSCCompService is the Component key [Service].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompService-swift.var
	KSCCompService string
	// KSCCompSystem is the Component key [System].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompSystem-swift.var
	KSCCompSystem string
	// KSCCompUsers is the Component key Users.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCCompUsers-swift.var
	KSCCompUsers string
	// KSCDynamicStoreDomainFile is the `File:` prefix.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreDomainFile-swift.var
	KSCDynamicStoreDomainFile string
	// KSCDynamicStoreDomainPlugin is the `Plugin:` prefix.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreDomainPlugin-swift.var
	KSCDynamicStoreDomainPlugin string
	// KSCDynamicStoreDomainPrefs is the `Prefs:` prefix.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreDomainPrefs-swift.var
	KSCDynamicStoreDomainPrefs string
	// KSCDynamicStoreDomainSetup is the `Setup:` prefix.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreDomainSetup-swift.var
	KSCDynamicStoreDomainSetup string
	// KSCDynamicStoreDomainState is the `State:` prefix.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreDomainState-swift.var
	KSCDynamicStoreDomainState string
	// KSCDynamicStorePropNetInterfaces is the dynamic store key [Interfaces], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropNetInterfaces-swift.var
	KSCDynamicStorePropNetInterfaces string
	// KSCDynamicStorePropNetPrimaryInterface is the dynamic store key [PrimaryInterface], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropNetPrimaryInterface-swift.var
	KSCDynamicStorePropNetPrimaryInterface string
	// KSCDynamicStorePropNetPrimaryService is the dynamic store key [PrimaryService], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropNetPrimaryService-swift.var
	KSCDynamicStorePropNetPrimaryService string
	// KSCDynamicStorePropNetServiceIDs is the dynamic store key [ServiceIDs], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropNetServiceIDs-swift.var
	KSCDynamicStorePropNetServiceIDs string
	// KSCDynamicStorePropSetupCurrentSet is the dynamic store key [CurrentSet], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropSetupCurrentSet-swift.var
	KSCDynamicStorePropSetupCurrentSet string
	// KSCDynamicStorePropSetupLastUpdated is the dynamic store key [LastUpdated].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStorePropSetupLastUpdated-swift.var
	KSCDynamicStorePropSetupLastUpdated string
	// KSCDynamicStoreUseSessionKeys is all keys added or set into the dynamic store should be per-session keys.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCDynamicStoreUseSessionKeys
	KSCDynamicStoreUseSessionKeys string
	// KSCEntNet6to4 is the network entity key for the `6to4` dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNet6to4-swift.var
	KSCEntNet6to4 string
	// KSCEntNetAirPort is the network entity key for the [AirPort] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetAirPort-swift.var
	KSCEntNetAirPort string
	// KSCEntNetDHCP is the network entity key for the [DHCP] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetDHCP-swift.var
	KSCEntNetDHCP string
	// KSCEntNetDNS is the network entity key for the [DNS] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetDNS-swift.var
	KSCEntNetDNS string
	// KSCEntNetEthernet is the network entity key for the [Ethernet] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetEthernet-swift.var
	KSCEntNetEthernet string
	// KSCEntNetFireWire is the network entity key for the [FireWire] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetFireWire-swift.var
	KSCEntNetFireWire string
	// KSCEntNetIPSec is the network entity key for the [IPSec] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetIPSec-swift.var
	KSCEntNetIPSec string
	// KSCEntNetIPv4 is the network entity key for the [IPv4] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetIPv4-swift.var
	KSCEntNetIPv4 string
	// KSCEntNetIPv6 is the network entity key for the [IPv6] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetIPv6-swift.var
	KSCEntNetIPv6 string
	// KSCEntNetInterface is the network entity key for the [Interface] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetInterface-swift.var
	KSCEntNetInterface string
	// KSCEntNetL2TP is the network entity key for the [L2TP] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetL2TP-swift.var
	KSCEntNetL2TP string
	// KSCEntNetLink is the network entity key for the [Link] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetLink-swift.var
	KSCEntNetLink string
	// KSCEntNetModem is the network entity key for the [Modem] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetModem-swift.var
	KSCEntNetModem string
	// KSCEntNetPPP is the network entity key for the [PPP] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetPPP-swift.var
	KSCEntNetPPP string
	// KSCEntNetPPPSerial is the network entity key for the [PPPSerial] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetPPPSerial-swift.var
	KSCEntNetPPPSerial string
	// KSCEntNetPPPoE is the network entity key for the [PPPoE] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetPPPoE-swift.var
	KSCEntNetPPPoE string
	// KSCEntNetProxies is the network entity key for the [Proxies] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetProxies-swift.var
	KSCEntNetProxies string
	// KSCEntNetSMB is the network entity key for the [SMB] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntNetSMB-swift.var
	KSCEntNetSMB string
	// KSCEntUsersConsoleUser is the CompUsers key [ConsoleUser].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCEntUsersConsoleUser-swift.var
	KSCEntUsersConsoleUser string
	// KSCNetworkInterfaceType6to4 is the 6to4 interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceType6to4
	KSCNetworkInterfaceType6to4 string
	// KSCNetworkInterfaceTypeBluetooth is the Bluetooth interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeBluetooth
	KSCNetworkInterfaceTypeBluetooth string
	// KSCNetworkInterfaceTypeBond is the Ethernet bond interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeBond
	KSCNetworkInterfaceTypeBond string
	// KSCNetworkInterfaceTypeEthernet is the Ethernet interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeEthernet
	KSCNetworkInterfaceTypeEthernet string
	// KSCNetworkInterfaceTypeFireWire is the FireWire interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeFireWire
	KSCNetworkInterfaceTypeFireWire string
	// KSCNetworkInterfaceTypeIEEE80211 is the IEEE 802.11 interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeIEEE80211
	KSCNetworkInterfaceTypeIEEE80211 string
	// KSCNetworkInterfaceTypeIPSec is the IPSec interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeIPSec
	KSCNetworkInterfaceTypeIPSec string
	// KSCNetworkInterfaceTypeIPv4 is the IPv4 interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeIPv4
	KSCNetworkInterfaceTypeIPv4 string
	// KSCNetworkInterfaceTypeL2TP is the L2TP interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeL2TP
	KSCNetworkInterfaceTypeL2TP string
	// KSCNetworkInterfaceTypeModem is the modem interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeModem
	KSCNetworkInterfaceTypeModem string
	// KSCNetworkInterfaceTypePPP is the PPP interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypePPP
	KSCNetworkInterfaceTypePPP string
	// KSCNetworkInterfaceTypeSerial is the serial interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeSerial
	KSCNetworkInterfaceTypeSerial string
	// KSCNetworkInterfaceTypeVLAN is the VLAN interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeVLAN
	KSCNetworkInterfaceTypeVLAN string
	// KSCNetworkInterfaceTypeWWAN is the WWAN interface.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceTypeWWAN
	KSCNetworkInterfaceTypeWWAN string
	// KSCNetworkProtocolTypeDNS is the DNS protocol.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkProtocolTypeDNS
	KSCNetworkProtocolTypeDNS string
	// KSCNetworkProtocolTypeIPv4 is the IPv4 protocol.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkProtocolTypeIPv4
	KSCNetworkProtocolTypeIPv4 string
	// KSCNetworkProtocolTypeIPv6 is the IPv6 protocol.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkProtocolTypeIPv6
	KSCNetworkProtocolTypeIPv6 string
	// KSCNetworkProtocolTypeProxies is protocol proxies.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkProtocolTypeProxies
	KSCNetworkProtocolTypeProxies string
	// KSCNetworkProtocolTypeSMB is the SMB procotol.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkProtocolTypeSMB
	KSCNetworkProtocolTypeSMB string
	// KSCPrefCurrentSet is the preference key [CurrentSet], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPrefCurrentSet-swift.var
	KSCPrefCurrentSet string
	// KSCPrefNetworkServices is the preference key for the [NetworkServices] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPrefNetworkServices-swift.var
	KSCPrefNetworkServices string
	// KSCPrefSets is the preference key for the [Sets] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPrefSets-swift.var
	KSCPrefSets string
	// KSCPrefSystem is the preference key for the [System] dictionary.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPrefSystem-swift.var
	KSCPrefSystem string
	// KSCPropInterfaceName is the generic key [InterfaceName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropInterfaceName-swift.var
	KSCPropInterfaceName string
	// KSCPropMACAddress is the generic key [MACAddress], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropMACAddress-swift.var
	KSCPropMACAddress string
	// KSCPropNet6to4Relay is the 6to4 key [Relay], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNet6to4Relay-swift.var
	KSCPropNet6to4Relay string
	// KSCPropNetDNSDomainName is the DNS key [DomainName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSDomainName-swift.var
	KSCPropNetDNSDomainName string
	// KSCPropNetDNSOptions is the DNS key [Options], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSOptions-swift.var
	KSCPropNetDNSOptions string
	// KSCPropNetDNSSearchDomains is the DNS key [SearchDomains], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSSearchDomains-swift.var
	KSCPropNetDNSSearchDomains string
	// KSCPropNetDNSSearchOrder is the DNS key [SearchOrder], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSSearchOrder-swift.var
	KSCPropNetDNSSearchOrder string
	// KSCPropNetDNSServerAddresses is the DNS key [ServerAddresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSServerAddresses-swift.var
	KSCPropNetDNSServerAddresses string
	// KSCPropNetDNSServerPort is the DNS key [ServerPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSServerPort-swift.var
	KSCPropNetDNSServerPort string
	// KSCPropNetDNSServerTimeout is the DNS key [ServerTimeout], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSServerTimeout-swift.var
	KSCPropNetDNSServerTimeout string
	// KSCPropNetDNSSortList is the DNS key [SortList], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSSortList-swift.var
	KSCPropNetDNSSortList string
	// KSCPropNetDNSSupplementalMatchDomains is the DNS key [SupplementalMatchDomains], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSSupplementalMatchDomains-swift.var
	KSCPropNetDNSSupplementalMatchDomains string
	// KSCPropNetDNSSupplementalMatchOrders is the DNS key [SupplementalMatchOrders], whose value is of type [CFArray], containing elements of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetDNSSupplementalMatchOrders-swift.var
	KSCPropNetDNSSupplementalMatchOrders string
	// KSCPropNetEthernetMTU is the Ethernet key [MTU], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetEthernetMTU-swift.var
	KSCPropNetEthernetMTU string
	// KSCPropNetEthernetMediaOptions is the Ethernet key [MediaOptions], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetEthernetMediaOptions-swift.var
	KSCPropNetEthernetMediaOptions string
	// KSCPropNetEthernetMediaSubType is the Ethernet key [MediaSubType], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetEthernetMediaSubType-swift.var
	KSCPropNetEthernetMediaSubType string
	// KSCPropNetIPSecAuthenticationMethod is the IPSec key [AuthenticationMethod], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecAuthenticationMethod-swift.var
	KSCPropNetIPSecAuthenticationMethod string
	// KSCPropNetIPSecConnectTime is the IPSec key ConnectTime.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecConnectTime-swift.var
	KSCPropNetIPSecConnectTime string
	// KSCPropNetIPSecLocalCertificate is the IPSec key [LocalCertificate], whose value is of type [CFData].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecLocalCertificate-swift.var
	KSCPropNetIPSecLocalCertificate string
	// KSCPropNetIPSecLocalIdentifier is the IPSec key [LocalIdentifier], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecLocalIdentifier-swift.var
	KSCPropNetIPSecLocalIdentifier string
	// KSCPropNetIPSecLocalIdentifierType is the IPSec key [LocalIdentifierType], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecLocalIdentifierType-swift.var
	KSCPropNetIPSecLocalIdentifierType string
	// KSCPropNetIPSecRemoteAddress is the IPSec key RemoteAddress.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecRemoteAddress-swift.var
	KSCPropNetIPSecRemoteAddress string
	// KSCPropNetIPSecSharedSecret is the IPSec key [SharedSecret], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecSharedSecret-swift.var
	KSCPropNetIPSecSharedSecret string
	// KSCPropNetIPSecSharedSecretEncryption is the IPSec key [SharedSecretEncryption], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecSharedSecretEncryption-swift.var
	KSCPropNetIPSecSharedSecretEncryption string
	// KSCPropNetIPSecStatus is the IPSec key Status.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecStatus-swift.var
	KSCPropNetIPSecStatus string
	// KSCPropNetIPSecXAuthEnabled is the IPSec key XAuthEnabled.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecXAuthEnabled-swift.var
	KSCPropNetIPSecXAuthEnabled string
	// KSCPropNetIPSecXAuthName is the IPSec key XAuthName.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecXAuthName-swift.var
	KSCPropNetIPSecXAuthName string
	// KSCPropNetIPSecXAuthPassword is the IPSec key XAuthPassword.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecXAuthPassword-swift.var
	KSCPropNetIPSecXAuthPassword string
	// KSCPropNetIPSecXAuthPasswordEncryption is the IPSec key XAuthPasswordEncryption.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPSecXAuthPasswordEncryption-swift.var
	KSCPropNetIPSecXAuthPasswordEncryption string
	// KSCPropNetIPv4Addresses is the IPv4 key [Addresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4Addresses-swift.var
	KSCPropNetIPv4Addresses string
	// KSCPropNetIPv4BroadcastAddresses is the IPv4 key [BroadcastAddresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4BroadcastAddresses-swift.var
	KSCPropNetIPv4BroadcastAddresses string
	// KSCPropNetIPv4ConfigMethod is the IPv4 key [ConfigMethod], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4ConfigMethod-swift.var
	KSCPropNetIPv4ConfigMethod string
	// KSCPropNetIPv4DHCPClientID is the IPv4 key [DHCPClientID], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4DHCPClientID-swift.var
	KSCPropNetIPv4DHCPClientID string
	// KSCPropNetIPv4DestAddresses is the IPv4 key [DestAddresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4DestAddresses-swift.var
	KSCPropNetIPv4DestAddresses string
	// KSCPropNetIPv4Router is the IPv4 key [Router], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4Router-swift.var
	KSCPropNetIPv4Router string
	// KSCPropNetIPv4SubnetMasks is the IPv4 key [SubnetMasks], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv4SubnetMasks-swift.var
	KSCPropNetIPv4SubnetMasks string
	// KSCPropNetIPv6Addresses is the IPv6 key [Addresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6Addresses-swift.var
	KSCPropNetIPv6Addresses string
	// KSCPropNetIPv6ConfigMethod is the IPv6 key [ConfigMethod], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6ConfigMethod-swift.var
	KSCPropNetIPv6ConfigMethod string
	// KSCPropNetIPv6DestAddresses is the IPv6 key [DestAddresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6DestAddresses-swift.var
	KSCPropNetIPv6DestAddresses string
	// KSCPropNetIPv6Flags is the IPv6 key [Flags], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6Flags-swift.var
	KSCPropNetIPv6Flags string
	// KSCPropNetIPv6PrefixLength is the IPv6 key [PrefixLength], whose value is of type [CFArray], containing elements of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6PrefixLength-swift.var
	KSCPropNetIPv6PrefixLength string
	// KSCPropNetIPv6Router is the IPv6 key [Router], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetIPv6Router-swift.var
	KSCPropNetIPv6Router string
	// KSCPropNetInterfaceDeviceName is the Interface key [DeviceName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetInterfaceDeviceName-swift.var
	KSCPropNetInterfaceDeviceName string
	// KSCPropNetInterfaceHardware is the Interface key [Hardware], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetInterfaceHardware-swift.var
	KSCPropNetInterfaceHardware string
	// KSCPropNetInterfaceSubType is the Interface key [SubType], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetInterfaceSubType-swift.var
	KSCPropNetInterfaceSubType string
	// KSCPropNetInterfaceType is the Interface key [Type], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetInterfaceType-swift.var
	KSCPropNetInterfaceType string
	// KSCPropNetInterfaces is the Network key [Interfaces], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetInterfaces-swift.var
	KSCPropNetInterfaces string
	// KSCPropNetL2TPIPSecSharedSecret is the L2TP key [IPSecSharedSecret], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetL2TPIPSecSharedSecret-swift.var
	KSCPropNetL2TPIPSecSharedSecret string
	// KSCPropNetL2TPIPSecSharedSecretEncryption is the L2TP key [IPSecSharedSecretEncryption], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetL2TPIPSecSharedSecretEncryption-swift.var
	KSCPropNetL2TPIPSecSharedSecretEncryption string
	// KSCPropNetL2TPTransport is the L2TP key [Transport], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetL2TPTransport-swift.var
	KSCPropNetL2TPTransport string
	// KSCPropNetLinkActive is the Link key [Active], whose value is of type [CFBoolean].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetLinkActive-swift.var
	KSCPropNetLinkActive string
	// KSCPropNetLinkDetaching is the Link key [Detaching], whose value is of type [CFBoolean].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetLinkDetaching-swift.var
	KSCPropNetLinkDetaching string
	// KSCPropNetLocalHostName is the Network key [LocalHostName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetLocalHostName-swift.var
	KSCPropNetLocalHostName string
	// KSCPropNetModemAccessPointName is the Modem key [AccessPointName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemAccessPointName-swift.var
	KSCPropNetModemAccessPointName string
	// KSCPropNetModemConnectSpeed is the Modem key [ConnectSpeed], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemConnectSpeed-swift.var
	KSCPropNetModemConnectSpeed string
	// KSCPropNetModemConnectionPersonality is the Modem key [ConnectionPersonality], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemConnectionPersonality-swift.var
	KSCPropNetModemConnectionPersonality string
	// KSCPropNetModemConnectionScript is the Modem key [ConnectionScript], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemConnectionScript-swift.var
	KSCPropNetModemConnectionScript string
	// KSCPropNetModemDataCompression is the Modem key [DataCompression], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemDataCompression-swift.var
	KSCPropNetModemDataCompression string
	// KSCPropNetModemDeviceContextID is the Modem key [DeviceContextID], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemDeviceContextID-swift.var
	KSCPropNetModemDeviceContextID string
	// KSCPropNetModemDeviceModel is the Modem key [DeviceModel], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemDeviceModel-swift.var
	KSCPropNetModemDeviceModel string
	// KSCPropNetModemDeviceVendor is the Modem key [DeviceVendor], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemDeviceVendor-swift.var
	KSCPropNetModemDeviceVendor string
	// KSCPropNetModemDialMode is the Modem key [DialMode], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemDialMode-swift.var
	KSCPropNetModemDialMode string
	// KSCPropNetModemErrorCorrection is the Modem key [ErrorCorrection], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemErrorCorrection-swift.var
	KSCPropNetModemErrorCorrection string
	// KSCPropNetModemHoldCallWaitingAudibleAlert is the Modem key [HoldCallWaitingAudibleAlert], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemHoldCallWaitingAudibleAlert-swift.var
	KSCPropNetModemHoldCallWaitingAudibleAlert string
	// KSCPropNetModemHoldDisconnectOnAnswer is the Modem key [HoldDisconnectOnAnswer], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemHoldDisconnectOnAnswer-swift.var
	KSCPropNetModemHoldDisconnectOnAnswer string
	// KSCPropNetModemHoldEnabled is the Modem key [HoldEnabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemHoldEnabled-swift.var
	KSCPropNetModemHoldEnabled string
	// KSCPropNetModemHoldReminder is the Modem key [HoldReminder], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemHoldReminder-swift.var
	KSCPropNetModemHoldReminder string
	// KSCPropNetModemHoldReminderTime is the Modem key [HoldReminderTime], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemHoldReminderTime-swift.var
	KSCPropNetModemHoldReminderTime string
	// KSCPropNetModemNote is the Modem key [Note], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemNote-swift.var
	KSCPropNetModemNote string
	// KSCPropNetModemPulseDial is the Modem key [PulseDial], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemPulseDial-swift.var
	KSCPropNetModemPulseDial string
	// KSCPropNetModemSpeaker is the Modem key [Speaker], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemSpeaker-swift.var
	KSCPropNetModemSpeaker string
	// KSCPropNetModemSpeed is the Modem key [Speed], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetModemSpeed-swift.var
	KSCPropNetModemSpeed string
	// KSCPropNetOverridePrimary is the Network key [OverridePrimary], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetOverridePrimary-swift.var
	KSCPropNetOverridePrimary string
	// KSCPropNetPPPACSPEnabled is the PPP key [ACSPEnabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPACSPEnabled-swift.var
	KSCPropNetPPPACSPEnabled string
	// KSCPropNetPPPAuthName is the PPP key [AuthName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPAuthName-swift.var
	KSCPropNetPPPAuthName string
	// KSCPropNetPPPAuthPassword is the PPP key [AuthPassword], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPAuthPassword-swift.var
	KSCPropNetPPPAuthPassword string
	// KSCPropNetPPPAuthPasswordEncryption is the PPP key [AuthPasswordEncryption], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPAuthPasswordEncryption-swift.var
	KSCPropNetPPPAuthPasswordEncryption string
	// KSCPropNetPPPAuthPrompt is the PPP key [AuthPrompt], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPAuthPrompt-swift.var
	KSCPropNetPPPAuthPrompt string
	// KSCPropNetPPPAuthProtocol is the PPP key [AuthProtocol], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPAuthProtocol-swift.var
	KSCPropNetPPPAuthProtocol string
	// KSCPropNetPPPCCPEnabled is the PPP key [CCPEnabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCCPEnabled-swift.var
	KSCPropNetPPPCCPEnabled string
	// KSCPropNetPPPCCPMPPE128Enabled is the PPP key [CCPMPPE128Enabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCCPMPPE128Enabled-swift.var
	KSCPropNetPPPCCPMPPE128Enabled string
	// KSCPropNetPPPCCPMPPE40Enabled is the PPP key [CCPMPPE40Enabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCCPMPPE40Enabled-swift.var
	KSCPropNetPPPCCPMPPE40Enabled string
	// KSCPropNetPPPCommAlternateRemoteAddress is the PPP key [CommAlternateRemoteAddress], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommAlternateRemoteAddress-swift.var
	KSCPropNetPPPCommAlternateRemoteAddress string
	// KSCPropNetPPPCommConnectDelay is the PPP key [CommConnectDelay], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommConnectDelay-swift.var
	KSCPropNetPPPCommConnectDelay string
	// KSCPropNetPPPCommDisplayTerminalWindow is the PPP key [CommDisplayTerminalWindow], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommDisplayTerminalWindow-swift.var
	KSCPropNetPPPCommDisplayTerminalWindow string
	// KSCPropNetPPPCommRedialCount is the PPP key [CommRedialCount], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommRedialCount-swift.var
	KSCPropNetPPPCommRedialCount string
	// KSCPropNetPPPCommRedialEnabled is the PPP key [CommRedialEnabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommRedialEnabled-swift.var
	KSCPropNetPPPCommRedialEnabled string
	// KSCPropNetPPPCommRedialInterval is the PPP key [CommRedialInterval], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommRedialInterval-swift.var
	KSCPropNetPPPCommRedialInterval string
	// KSCPropNetPPPCommRemoteAddress is the PPP key [CommRemoteAddress], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommRemoteAddress-swift.var
	KSCPropNetPPPCommRemoteAddress string
	// KSCPropNetPPPCommTerminalScript is the PPP key [CommTerminalScript], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommTerminalScript-swift.var
	KSCPropNetPPPCommTerminalScript string
	// KSCPropNetPPPCommUseTerminalScript is the PPP key [CommUseTerminalScript], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPCommUseTerminalScript-swift.var
	KSCPropNetPPPCommUseTerminalScript string
	// KSCPropNetPPPConnectTime is the PPP key [ConnectTime], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPConnectTime-swift.var
	KSCPropNetPPPConnectTime string
	// KSCPropNetPPPDeviceLastCause is the PPP key [DeviceLastCause], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDeviceLastCause-swift.var
	KSCPropNetPPPDeviceLastCause string
	// KSCPropNetPPPDialOnDemand is the PPP key [DialOnDemand], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDialOnDemand-swift.var
	KSCPropNetPPPDialOnDemand string
	// KSCPropNetPPPDisconnectOnFastUserSwitch is the PPP key [DisconnectOnFastUserSwitch], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectOnFastUserSwitch-swift.var
	KSCPropNetPPPDisconnectOnFastUserSwitch string
	// KSCPropNetPPPDisconnectOnIdle is the PPP key [DisconnectOnIdle], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectOnIdle-swift.var
	KSCPropNetPPPDisconnectOnIdle string
	// KSCPropNetPPPDisconnectOnIdleTimer is the PPP key [DisconnectOnIdleTimer], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectOnIdleTimer-swift.var
	KSCPropNetPPPDisconnectOnIdleTimer string
	// KSCPropNetPPPDisconnectOnLogout is the PPP key [DisconnectOnLogout], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectOnLogout-swift.var
	KSCPropNetPPPDisconnectOnLogout string
	// KSCPropNetPPPDisconnectOnSleep is the PPP key [DisconnectOnSleep], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectOnSleep-swift.var
	KSCPropNetPPPDisconnectOnSleep string
	// KSCPropNetPPPDisconnectTime is the PPP key [DisconnectTime], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPDisconnectTime-swift.var
	KSCPropNetPPPDisconnectTime string
	// KSCPropNetPPPIPCPCompressionVJ is the PPP key [IPCPCompressionVJ], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPIPCPCompressionVJ-swift.var
	KSCPropNetPPPIPCPCompressionVJ string
	// KSCPropNetPPPIPCPUsePeerDNS is the PPP key [IPCPUsePeerDNS], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPIPCPUsePeerDNS-swift.var
	KSCPropNetPPPIPCPUsePeerDNS string
	// KSCPropNetPPPIdleReminder is the PPP key [IdleReminder], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPIdleReminder-swift.var
	KSCPropNetPPPIdleReminder string
	// KSCPropNetPPPIdleReminderTimer is the PPP key [IdleReminderTimer], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPIdleReminderTimer-swift.var
	KSCPropNetPPPIdleReminderTimer string
	// KSCPropNetPPPLCPCompressionACField is the PPP key [LCPCompressionACField], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPCompressionACField-swift.var
	KSCPropNetPPPLCPCompressionACField string
	// KSCPropNetPPPLCPCompressionPField is the PPP key [LCPCompressionPField], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPCompressionPField-swift.var
	KSCPropNetPPPLCPCompressionPField string
	// KSCPropNetPPPLCPEchoEnabled is the PPP key [LCPEchoEnabled], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPEchoEnabled-swift.var
	KSCPropNetPPPLCPEchoEnabled string
	// KSCPropNetPPPLCPEchoFailure is the PPP key [LCPEchoFailure], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPEchoFailure-swift.var
	KSCPropNetPPPLCPEchoFailure string
	// KSCPropNetPPPLCPEchoInterval is the PPP key [LCPEchoInterval], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPEchoInterval-swift.var
	KSCPropNetPPPLCPEchoInterval string
	// KSCPropNetPPPLCPMRU is the PPP key [LCPMRU], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPMRU-swift.var
	KSCPropNetPPPLCPMRU string
	// KSCPropNetPPPLCPMTU is the PPP key [LCPMTU], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPMTU-swift.var
	KSCPropNetPPPLCPMTU string
	// KSCPropNetPPPLCPReceiveACCM is the PPP key [LCPReceiveACCM], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPReceiveACCM-swift.var
	KSCPropNetPPPLCPReceiveACCM string
	// KSCPropNetPPPLCPTransmitACCM is the PPP key [LCPTransmitACCM], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLCPTransmitACCM-swift.var
	KSCPropNetPPPLCPTransmitACCM string
	// KSCPropNetPPPLastCause is the PPP key [LastCause], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLastCause-swift.var
	KSCPropNetPPPLastCause string
	// KSCPropNetPPPLogfile is the PPP key [Logfile], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPLogfile-swift.var
	KSCPropNetPPPLogfile string
	// KSCPropNetPPPOverridePrimary is the Network key [PPPOverridePrimary], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPOverridePrimary-swift.var
	KSCPropNetPPPOverridePrimary string
	// KSCPropNetPPPRetryConnectTime is the PPP key [RetryConnectTime], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPRetryConnectTime-swift.var
	KSCPropNetPPPRetryConnectTime string
	// KSCPropNetPPPSessionTimer is the PPP key [SessionTimer], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPSessionTimer-swift.var
	KSCPropNetPPPSessionTimer string
	// KSCPropNetPPPStatus is the PPP key [Status], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPStatus-swift.var
	KSCPropNetPPPStatus string
	// KSCPropNetPPPUseSessionTimer is the PPP key [UseSessionTimer], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPUseSessionTimer-swift.var
	KSCPropNetPPPUseSessionTimer string
	// KSCPropNetPPPVerboseLogging is the PPP key [VerboseLogging], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetPPPVerboseLogging-swift.var
	KSCPropNetPPPVerboseLogging string
	// KSCPropNetProxiesExceptionsList is the Proxies key [ExceptionsList], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesExceptionsList-swift.var
	KSCPropNetProxiesExceptionsList string
	// KSCPropNetProxiesExcludeSimpleHostnames is the Proxies key [ExcludeSimpleHostnames], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesExcludeSimpleHostnames-swift.var
	KSCPropNetProxiesExcludeSimpleHostnames string
	// KSCPropNetProxiesFTPEnable is the Proxies key [FTPEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesFTPEnable-swift.var
	KSCPropNetProxiesFTPEnable string
	// KSCPropNetProxiesFTPPassive is the Proxies key [FTPPassive], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesFTPPassive-swift.var
	KSCPropNetProxiesFTPPassive string
	// KSCPropNetProxiesFTPPort is the Proxies key [FTPPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesFTPPort-swift.var
	KSCPropNetProxiesFTPPort string
	// KSCPropNetProxiesFTPProxy is the Proxies key [FTPProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesFTPProxy-swift.var
	KSCPropNetProxiesFTPProxy string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesFTPUser-swift.var
	KSCPropNetProxiesFTPUser string
	// KSCPropNetProxiesGopherEnable is the Proxies key [GopherEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesGopherEnable-swift.var
	KSCPropNetProxiesGopherEnable string
	// KSCPropNetProxiesGopherPort is the Proxies key [GopherPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesGopherPort-swift.var
	KSCPropNetProxiesGopherPort string
	// KSCPropNetProxiesGopherProxy is the Proxies key [GopherProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesGopherProxy-swift.var
	KSCPropNetProxiesGopherProxy string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesGopherUser-swift.var
	KSCPropNetProxiesGopherUser string
	// KSCPropNetProxiesHTTPEnable is the Proxies key [HTTPEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPEnable-swift.var
	KSCPropNetProxiesHTTPEnable string
	// KSCPropNetProxiesHTTPPort is the Proxies key [HTTPPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPPort-swift.var
	KSCPropNetProxiesHTTPPort string
	// KSCPropNetProxiesHTTPProxy is the Proxies key [HTTPProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPProxy-swift.var
	KSCPropNetProxiesHTTPProxy string
	// KSCPropNetProxiesHTTPSEnable is the Proxies key [HTTPSEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPSEnable-swift.var
	KSCPropNetProxiesHTTPSEnable string
	// KSCPropNetProxiesHTTPSPort is the Proxies key [HTTPSPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPSPort-swift.var
	KSCPropNetProxiesHTTPSPort string
	// KSCPropNetProxiesHTTPSProxy is the Proxies key [HTTPSProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPSProxy-swift.var
	KSCPropNetProxiesHTTPSProxy string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPSUser-swift.var
	KSCPropNetProxiesHTTPSUser string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesHTTPUser-swift.var
	KSCPropNetProxiesHTTPUser string
	// KSCPropNetProxiesProxyAutoConfigEnable is the Proxies key [ProxyAutoConfigEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesProxyAutoConfigEnable-swift.var
	KSCPropNetProxiesProxyAutoConfigEnable string
	// KSCPropNetProxiesProxyAutoConfigJavaScript is the Proxies key ProxyAutoConfigJavaScript.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesProxyAutoConfigJavaScript-swift.var
	KSCPropNetProxiesProxyAutoConfigJavaScript string
	// KSCPropNetProxiesProxyAutoConfigURLString is the Proxies key [ProxyAutoConfigURLString], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesProxyAutoConfigURLString-swift.var
	KSCPropNetProxiesProxyAutoConfigURLString string
	// KSCPropNetProxiesProxyAutoDiscoveryEnable is the Proxies key [ProxyAutoDiscoveryEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesProxyAutoDiscoveryEnable-swift.var
	KSCPropNetProxiesProxyAutoDiscoveryEnable string
	// KSCPropNetProxiesRTSPEnable is the Proxies key [RTSPEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesRTSPEnable-swift.var
	KSCPropNetProxiesRTSPEnable string
	// KSCPropNetProxiesRTSPPort is the Proxies key [RTSPPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesRTSPPort-swift.var
	KSCPropNetProxiesRTSPPort string
	// KSCPropNetProxiesRTSPProxy is the Proxies key [RTSPProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesRTSPProxy-swift.var
	KSCPropNetProxiesRTSPProxy string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesRTSPUser-swift.var
	KSCPropNetProxiesRTSPUser string
	// KSCPropNetProxiesSOCKSEnable is the Proxies key [SOCKSEnable], whose value is of type [CFNumber] and is equal to `0` or `1`.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesSOCKSEnable-swift.var
	KSCPropNetProxiesSOCKSEnable string
	// KSCPropNetProxiesSOCKSPort is the Proxies key [SOCKSPort], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesSOCKSPort-swift.var
	KSCPropNetProxiesSOCKSPort string
	// KSCPropNetProxiesSOCKSProxy is the Proxies key [SOCKSProxy], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesSOCKSProxy-swift.var
	KSCPropNetProxiesSOCKSProxy string
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetProxiesSOCKSUser-swift.var
	KSCPropNetProxiesSOCKSUser string
	// KSCPropNetSMBNetBIOSName is the SMB key [NetBIOSName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetSMBNetBIOSName-swift.var
	KSCPropNetSMBNetBIOSName string
	// KSCPropNetSMBNetBIOSNodeType is the SMB key [NetBIOSNodeType], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetSMBNetBIOSNodeType-swift.var
	KSCPropNetSMBNetBIOSNodeType string
	// KSCPropNetSMBWINSAddresses is the SMB key [WINSAddresses], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetSMBWINSAddresses-swift.var
	KSCPropNetSMBWINSAddresses string
	// KSCPropNetSMBWorkgroup is the SMB key [Workgroup], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetSMBWorkgroup-swift.var
	KSCPropNetSMBWorkgroup string
	// KSCPropNetServiceOrder is the Network key [ServiceOrder], whose value is of type [CFArray], containing elements of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropNetServiceOrder-swift.var
	KSCPropNetServiceOrder string
	// KSCPropSystemComputerName is the CompSystem key [ComputerName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropSystemComputerName-swift.var
	KSCPropSystemComputerName string
	// KSCPropSystemComputerNameEncoding is the CompSystem key [ComputerNameEncoding], whose value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropSystemComputerNameEncoding-swift.var
	KSCPropSystemComputerNameEncoding string
	// KSCPropUserDefinedName is the generic key [UserDefinedName], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropUserDefinedName-swift.var
	KSCPropUserDefinedName string
	// KSCPropVersion is the generic key [Version], whose value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCPropVersion-swift.var
	KSCPropVersion string
	// KSCResvInactive is the reserved key __INACTIVE__.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCResvInactive-swift.var
	KSCResvInactive string
	// KSCResvLink is the reserved key __LINK__.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCResvLink-swift.var
	KSCResvLink string
	// KSCValNetIPSecAuthenticationMethodCertificate is the constant value Certificate.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecAuthenticationMethodCertificate-swift.var
	KSCValNetIPSecAuthenticationMethodCertificate string
	// KSCValNetIPSecAuthenticationMethodHybrid is the constant value Hybrid.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecAuthenticationMethodHybrid-swift.var
	KSCValNetIPSecAuthenticationMethodHybrid string
	// KSCValNetIPSecAuthenticationMethodSharedSecret is the constant value SharedSecret.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecAuthenticationMethodSharedSecret-swift.var
	KSCValNetIPSecAuthenticationMethodSharedSecret string
	// KSCValNetIPSecLocalIdentifierTypeKeyID is the constant value KeyID.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecLocalIdentifierTypeKeyID-swift.var
	KSCValNetIPSecLocalIdentifierTypeKeyID string
	// KSCValNetIPSecSharedSecretEncryptionKeychain is the constant value Keychain.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecSharedSecretEncryptionKeychain-swift.var
	KSCValNetIPSecSharedSecretEncryptionKeychain string
	// KSCValNetIPSecXAuthPasswordEncryptionKeychain is the constant value Keychain.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecXAuthPasswordEncryptionKeychain-swift.var
	KSCValNetIPSecXAuthPasswordEncryptionKeychain string
	// KSCValNetIPSecXAuthPasswordEncryptionPrompt is the constant value Prompt.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPSecXAuthPasswordEncryptionPrompt-swift.var
	KSCValNetIPSecXAuthPasswordEncryptionPrompt string
	// KSCValNetIPv4ConfigMethodAutomatic is the constant value Automatic.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodAutomatic-swift.var
	KSCValNetIPv4ConfigMethodAutomatic string
	// KSCValNetIPv4ConfigMethodBOOTP is the constant value BOOTP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodBOOTP-swift.var
	KSCValNetIPv4ConfigMethodBOOTP string
	// KSCValNetIPv4ConfigMethodDHCP is the constant value DHCP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodDHCP-swift.var
	KSCValNetIPv4ConfigMethodDHCP string
	// KSCValNetIPv4ConfigMethodINFORM is the constant value INFORM.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodINFORM-swift.var
	KSCValNetIPv4ConfigMethodINFORM string
	// KSCValNetIPv4ConfigMethodLinkLocal is the constant value LinkLocal.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodLinkLocal-swift.var
	KSCValNetIPv4ConfigMethodLinkLocal string
	// KSCValNetIPv4ConfigMethodManual is the constant value Manual.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodManual-swift.var
	KSCValNetIPv4ConfigMethodManual string
	// KSCValNetIPv4ConfigMethodPPP is the constant value PPP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv4ConfigMethodPPP-swift.var
	KSCValNetIPv4ConfigMethodPPP string
	// KSCValNetIPv6ConfigMethod6to4 is the constant value 6to4.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv6ConfigMethod6to4-swift.var
	KSCValNetIPv6ConfigMethod6to4 string
	// KSCValNetIPv6ConfigMethodAutomatic is the constant value Automatic.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv6ConfigMethodAutomatic-swift.var
	KSCValNetIPv6ConfigMethodAutomatic string
	// KSCValNetIPv6ConfigMethodLinkLocal is the constant value LinkLocal.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv6ConfigMethodLinkLocal-swift.var
	KSCValNetIPv6ConfigMethodLinkLocal string
	// KSCValNetIPv6ConfigMethodManual is the constant value Manual.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv6ConfigMethodManual-swift.var
	KSCValNetIPv6ConfigMethodManual string
	// KSCValNetIPv6ConfigMethodRouterAdvertisement is the constant value RouterAdvertisement.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetIPv6ConfigMethodRouterAdvertisement-swift.var
	KSCValNetIPv6ConfigMethodRouterAdvertisement string
	// KSCValNetInterfaceSubTypeL2TP is the constant value L2TP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceSubTypeL2TP-swift.var
	KSCValNetInterfaceSubTypeL2TP string
	// KSCValNetInterfaceSubTypePPPSerial is the constant value PPPSerial.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceSubTypePPPSerial-swift.var
	KSCValNetInterfaceSubTypePPPSerial string
	// KSCValNetInterfaceSubTypePPPoE is the constant value PPPoE.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceSubTypePPPoE-swift.var
	KSCValNetInterfaceSubTypePPPoE string
	// KSCValNetInterfaceType6to4 is the constant value 6to4.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceType6to4-swift.var
	KSCValNetInterfaceType6to4 string
	// KSCValNetInterfaceTypeEthernet is the constant value Ethernet.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceTypeEthernet-swift.var
	KSCValNetInterfaceTypeEthernet string
	// KSCValNetInterfaceTypeFireWire is the constant value FireWire.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceTypeFireWire-swift.var
	KSCValNetInterfaceTypeFireWire string
	// KSCValNetInterfaceTypeIPSec is the constant value IPSec.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceTypeIPSec-swift.var
	KSCValNetInterfaceTypeIPSec string
	// KSCValNetInterfaceTypePPP is the constant value PPP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetInterfaceTypePPP-swift.var
	KSCValNetInterfaceTypePPP string
	// KSCValNetL2TPIPSecSharedSecretEncryptionKeychain is the constant value Keychain.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetL2TPIPSecSharedSecretEncryptionKeychain-swift.var
	KSCValNetL2TPIPSecSharedSecretEncryptionKeychain string
	// KSCValNetL2TPTransportIP is the constant value IP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetL2TPTransportIP-swift.var
	KSCValNetL2TPTransportIP string
	// KSCValNetL2TPTransportIPSec is the constant value IPSec.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetL2TPTransportIPSec-swift.var
	KSCValNetL2TPTransportIPSec string
	// KSCValNetModemDialModeIgnoreDialTone is the constant value IgnoreDialTone.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetModemDialModeIgnoreDialTone-swift.var
	KSCValNetModemDialModeIgnoreDialTone string
	// KSCValNetModemDialModeManual is the constant value Manual.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetModemDialModeManual-swift.var
	KSCValNetModemDialModeManual string
	// KSCValNetModemDialModeWaitForDialTone is the constant value WaitForDialTone.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetModemDialModeWaitForDialTone-swift.var
	KSCValNetModemDialModeWaitForDialTone string
	// KSCValNetPPPAuthPasswordEncryptionKeychain is the constant value Keychain.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthPasswordEncryptionKeychain-swift.var
	KSCValNetPPPAuthPasswordEncryptionKeychain string
	// KSCValNetPPPAuthPasswordEncryptionToken is the constant value Token.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthPasswordEncryptionToken-swift.var
	KSCValNetPPPAuthPasswordEncryptionToken string
	// KSCValNetPPPAuthPromptAfter is the constant value After.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthPromptAfter-swift.var
	KSCValNetPPPAuthPromptAfter string
	// KSCValNetPPPAuthPromptBefore is the constant value Before.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthPromptBefore-swift.var
	KSCValNetPPPAuthPromptBefore string
	// KSCValNetPPPAuthProtocolCHAP is the constant value CHAP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthProtocolCHAP-swift.var
	KSCValNetPPPAuthProtocolCHAP string
	// KSCValNetPPPAuthProtocolEAP is the constant value EAP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthProtocolEAP-swift.var
	KSCValNetPPPAuthProtocolEAP string
	// KSCValNetPPPAuthProtocolMSCHAP1 is the constant value MSCHAP1.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthProtocolMSCHAP1-swift.var
	KSCValNetPPPAuthProtocolMSCHAP1 string
	// KSCValNetPPPAuthProtocolMSCHAP2 is the constant value MSCHAP2.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthProtocolMSCHAP2-swift.var
	KSCValNetPPPAuthProtocolMSCHAP2 string
	// KSCValNetPPPAuthProtocolPAP is the constant value PAP.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetPPPAuthProtocolPAP-swift.var
	KSCValNetPPPAuthProtocolPAP string
	// KSCValNetSMBNetBIOSNodeTypeBroadcast is the constant value Broadcast.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetSMBNetBIOSNodeTypeBroadcast-swift.var
	KSCValNetSMBNetBIOSNodeTypeBroadcast string
	// KSCValNetSMBNetBIOSNodeTypeHybrid is the constant value Hybrid.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetSMBNetBIOSNodeTypeHybrid-swift.var
	KSCValNetSMBNetBIOSNodeTypeHybrid string
	// KSCValNetSMBNetBIOSNodeTypeMixed is the constant value Mixed.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetSMBNetBIOSNodeTypeMixed-swift.var
	KSCValNetSMBNetBIOSNodeTypeMixed string
	// KSCValNetSMBNetBIOSNodeTypePeer is the constant value Peer.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCValNetSMBNetBIOSNodeTypePeer-swift.var
	KSCValNetSMBNetBIOSNodeTypePeer string
)

var (
	// KSCNetworkInterfaceIPv4 is a network interface that can used for layering other interfaces (for example, 6to4, PPTP, or L2TP) over an existing IPv4 network.
	//
	// See: https://developer.apple.com/documentation/SystemConfiguration/kSCNetworkInterfaceIPv4
	KSCNetworkInterfaceIPv4 SCNetworkInterfaceRef
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainSystemConfiguration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorDomainSystemConfiguration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCBondStatusDeviceAggregationStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCBondStatusDeviceAggregationStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCBondStatusDeviceCollecting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCBondStatusDeviceCollecting = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCBondStatusDeviceDistributing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCBondStatusDeviceDistributing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompAnyRegex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompAnyRegex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompGlobal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompGlobal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompHostNames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompHostNames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompInterface"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompInterface = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompNetwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompNetwork = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompService"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompService = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompSystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompSystem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCCompUsers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCCompUsers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreDomainFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreDomainFile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreDomainPlugin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreDomainPlugin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreDomainPrefs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreDomainPrefs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreDomainSetup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreDomainSetup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreDomainState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreDomainState = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropNetInterfaces"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropNetInterfaces = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropNetPrimaryInterface"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropNetPrimaryInterface = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropNetPrimaryService"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropNetPrimaryService = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropNetServiceIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropNetServiceIDs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropSetupCurrentSet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropSetupCurrentSet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStorePropSetupLastUpdated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStorePropSetupLastUpdated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCDynamicStoreUseSessionKeys"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCDynamicStoreUseSessionKeys = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNet6to4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNet6to4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetAirPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetAirPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetDHCP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetDHCP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetDNS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetDNS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetEthernet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetEthernet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetFireWire"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetFireWire = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetIPSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetIPSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetIPv4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetIPv4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetIPv6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetIPv6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetInterface"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetInterface = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetL2TP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetL2TP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetModem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetModem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetPPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetPPP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetPPPSerial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetPPPSerial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetPPPoE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetPPPoE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetProxies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetProxies = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntNetSMB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntNetSMB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCEntUsersConsoleUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCEntUsersConsoleUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceIPv4"); err == nil && ptr != 0 {
		KSCNetworkInterfaceIPv4 = objc.ValueAt[SCNetworkInterfaceRef](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceType6to4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceType6to4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeBluetooth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeBluetooth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeBond"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeBond = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeEthernet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeEthernet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeFireWire"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeFireWire = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeIEEE80211"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeIEEE80211 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeIPSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeIPSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeIPv4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeIPv4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeL2TP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeL2TP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeModem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeModem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypePPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypePPP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeSerial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeSerial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeVLAN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeVLAN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkInterfaceTypeWWAN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkInterfaceTypeWWAN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkProtocolTypeDNS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkProtocolTypeDNS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkProtocolTypeIPv4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkProtocolTypeIPv4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkProtocolTypeIPv6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkProtocolTypeIPv6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkProtocolTypeProxies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkProtocolTypeProxies = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCNetworkProtocolTypeSMB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCNetworkProtocolTypeSMB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPrefCurrentSet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPrefCurrentSet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPrefNetworkServices"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPrefNetworkServices = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPrefSets"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPrefSets = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPrefSystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPrefSystem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropInterfaceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropInterfaceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropMACAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropMACAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNet6to4Relay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNet6to4Relay = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSDomainName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSDomainName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSOptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSSearchDomains"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSSearchDomains = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSSearchOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSSearchOrder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSServerAddresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSServerAddresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSServerPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSServerPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSServerTimeout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSServerTimeout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSSortList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSSortList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSSupplementalMatchDomains"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSSupplementalMatchDomains = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetDNSSupplementalMatchOrders"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetDNSSupplementalMatchOrders = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetEthernetMTU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetEthernetMTU = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetEthernetMediaOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetEthernetMediaOptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetEthernetMediaSubType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetEthernetMediaSubType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecAuthenticationMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecAuthenticationMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecConnectTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecConnectTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecLocalCertificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecLocalCertificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecLocalIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecLocalIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecLocalIdentifierType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecLocalIdentifierType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecRemoteAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecRemoteAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecSharedSecret"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecSharedSecret = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecSharedSecretEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecSharedSecretEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecXAuthEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecXAuthEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecXAuthName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecXAuthName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecXAuthPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecXAuthPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPSecXAuthPasswordEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPSecXAuthPasswordEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4Addresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4Addresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4BroadcastAddresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4BroadcastAddresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4ConfigMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4ConfigMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4DHCPClientID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4DHCPClientID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4DestAddresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4DestAddresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4Router"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4Router = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv4SubnetMasks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv4SubnetMasks = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6Addresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6Addresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6ConfigMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6ConfigMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6DestAddresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6DestAddresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6Flags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6Flags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6PrefixLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6PrefixLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetIPv6Router"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetIPv6Router = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetInterfaceDeviceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetInterfaceDeviceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetInterfaceHardware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetInterfaceHardware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetInterfaceSubType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetInterfaceSubType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetInterfaceType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetInterfaceType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetInterfaces"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetInterfaces = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetL2TPIPSecSharedSecret"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetL2TPIPSecSharedSecret = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetL2TPIPSecSharedSecretEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetL2TPIPSecSharedSecretEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetL2TPTransport"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetL2TPTransport = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetLinkActive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetLinkActive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetLinkDetaching"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetLinkDetaching = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetLocalHostName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetLocalHostName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemAccessPointName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemAccessPointName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemConnectSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemConnectSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemConnectionPersonality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemConnectionPersonality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemConnectionScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemConnectionScript = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemDataCompression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemDataCompression = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemDeviceContextID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemDeviceContextID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemDeviceModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemDeviceModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemDeviceVendor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemDeviceVendor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemDialMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemDialMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemErrorCorrection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemErrorCorrection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemHoldCallWaitingAudibleAlert"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemHoldCallWaitingAudibleAlert = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemHoldDisconnectOnAnswer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemHoldDisconnectOnAnswer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemHoldEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemHoldEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemHoldReminder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemHoldReminder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemHoldReminderTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemHoldReminderTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemNote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemNote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemPulseDial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemPulseDial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemSpeaker"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemSpeaker = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetModemSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetModemSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetOverridePrimary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetOverridePrimary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPACSPEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPACSPEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPAuthName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPAuthName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPAuthPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPAuthPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPAuthPasswordEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPAuthPasswordEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPAuthPrompt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPAuthPrompt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPAuthProtocol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPAuthProtocol = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCCPEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCCPEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCCPMPPE128Enabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCCPMPPE128Enabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCCPMPPE40Enabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCCPMPPE40Enabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommAlternateRemoteAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommAlternateRemoteAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommConnectDelay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommConnectDelay = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommDisplayTerminalWindow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommDisplayTerminalWindow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommRedialCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommRedialCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommRedialEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommRedialEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommRedialInterval"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommRedialInterval = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommRemoteAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommRemoteAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommTerminalScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommTerminalScript = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPCommUseTerminalScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPCommUseTerminalScript = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPConnectTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPConnectTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDeviceLastCause"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDeviceLastCause = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDialOnDemand"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDialOnDemand = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectOnFastUserSwitch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectOnFastUserSwitch = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectOnIdle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectOnIdle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectOnIdleTimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectOnIdleTimer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectOnLogout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectOnLogout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectOnSleep"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectOnSleep = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPDisconnectTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPDisconnectTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPIPCPCompressionVJ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPIPCPCompressionVJ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPIPCPUsePeerDNS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPIPCPUsePeerDNS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPIdleReminder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPIdleReminder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPIdleReminderTimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPIdleReminderTimer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPCompressionACField"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPCompressionACField = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPCompressionPField"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPCompressionPField = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPEchoEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPEchoEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPEchoFailure"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPEchoFailure = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPEchoInterval"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPEchoInterval = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPMRU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPMRU = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPMTU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPMTU = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPReceiveACCM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPReceiveACCM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLCPTransmitACCM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLCPTransmitACCM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLastCause"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLastCause = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPLogfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPLogfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPOverridePrimary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPOverridePrimary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPRetryConnectTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPRetryConnectTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPSessionTimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPSessionTimer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPUseSessionTimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPUseSessionTimer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetPPPVerboseLogging"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetPPPVerboseLogging = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesExceptionsList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesExceptionsList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesExcludeSimpleHostnames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesExcludeSimpleHostnames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesFTPEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesFTPEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesFTPPassive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesFTPPassive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesFTPPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesFTPPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesFTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesFTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesFTPUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesFTPUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesGopherEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesGopherEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesGopherPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesGopherPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesGopherProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesGopherProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesGopherUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesGopherUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPSEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPSEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPSPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPSPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPSUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPSUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesHTTPUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesHTTPUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesProxyAutoConfigEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesProxyAutoConfigEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesProxyAutoConfigJavaScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesProxyAutoConfigJavaScript = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesProxyAutoConfigURLString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesProxyAutoConfigURLString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesProxyAutoDiscoveryEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesProxyAutoDiscoveryEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesRTSPEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesRTSPEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesRTSPPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesRTSPPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesRTSPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesRTSPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesRTSPUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesRTSPUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesSOCKSEnable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesSOCKSEnable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesSOCKSPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesSOCKSPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesSOCKSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesSOCKSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetProxiesSOCKSUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetProxiesSOCKSUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetSMBNetBIOSName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetSMBNetBIOSName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetSMBNetBIOSNodeType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetSMBNetBIOSNodeType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetSMBWINSAddresses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetSMBWINSAddresses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetSMBWorkgroup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetSMBWorkgroup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropNetServiceOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropNetServiceOrder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropSystemComputerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropSystemComputerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropSystemComputerNameEncoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropSystemComputerNameEncoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropUserDefinedName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropUserDefinedName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCPropVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCPropVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCResvInactive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCResvInactive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCResvLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCResvLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecAuthenticationMethodCertificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecAuthenticationMethodCertificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecAuthenticationMethodHybrid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecAuthenticationMethodHybrid = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecAuthenticationMethodSharedSecret"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecAuthenticationMethodSharedSecret = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecLocalIdentifierTypeKeyID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecLocalIdentifierTypeKeyID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecSharedSecretEncryptionKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecSharedSecretEncryptionKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecXAuthPasswordEncryptionKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecXAuthPasswordEncryptionKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPSecXAuthPasswordEncryptionPrompt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPSecXAuthPasswordEncryptionPrompt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodAutomatic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodAutomatic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodBOOTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodBOOTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodDHCP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodDHCP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodINFORM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodINFORM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodLinkLocal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodLinkLocal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodManual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodManual = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv4ConfigMethodPPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv4ConfigMethodPPP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv6ConfigMethod6to4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv6ConfigMethod6to4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv6ConfigMethodAutomatic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv6ConfigMethodAutomatic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv6ConfigMethodLinkLocal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv6ConfigMethodLinkLocal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv6ConfigMethodManual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv6ConfigMethodManual = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetIPv6ConfigMethodRouterAdvertisement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetIPv6ConfigMethodRouterAdvertisement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceSubTypeL2TP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceSubTypeL2TP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceSubTypePPPSerial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceSubTypePPPSerial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceSubTypePPPoE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceSubTypePPPoE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceType6to4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceType6to4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceTypeEthernet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceTypeEthernet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceTypeFireWire"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceTypeFireWire = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceTypeIPSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceTypeIPSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetInterfaceTypePPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetInterfaceTypePPP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetL2TPIPSecSharedSecretEncryptionKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetL2TPIPSecSharedSecretEncryptionKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetL2TPTransportIP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetL2TPTransportIP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetL2TPTransportIPSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetL2TPTransportIPSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetModemDialModeIgnoreDialTone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetModemDialModeIgnoreDialTone = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetModemDialModeManual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetModemDialModeManual = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetModemDialModeWaitForDialTone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetModemDialModeWaitForDialTone = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthPasswordEncryptionKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthPasswordEncryptionKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthPasswordEncryptionToken"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthPasswordEncryptionToken = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthPromptAfter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthPromptAfter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthPromptBefore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthPromptBefore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthProtocolCHAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthProtocolCHAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthProtocolEAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthProtocolEAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthProtocolMSCHAP1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthProtocolMSCHAP1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthProtocolMSCHAP2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthProtocolMSCHAP2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetPPPAuthProtocolPAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetPPPAuthProtocolPAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetSMBNetBIOSNodeTypeBroadcast"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetSMBNetBIOSNodeTypeBroadcast = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetSMBNetBIOSNodeTypeHybrid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetSMBNetBIOSNodeTypeHybrid = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetSMBNetBIOSNodeTypeMixed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetSMBNetBIOSNodeTypeMixed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSCValNetSMBNetBIOSNodeTypePeer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSCValNetSMBNetBIOSNodeTypePeer = objc.GoString(cstr)
			}
		}
	}

}
