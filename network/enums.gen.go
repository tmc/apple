// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"fmt"
)

type NWBrowseResultChange uint32

const (
	// NWBrowseResultChangeIdentical: The compared services are identical.
	NWBrowseResultChangeIdentical NWBrowseResultChange = 0x1
	// NWBrowseResultChangeInterfaceAdded: The service was discovered over a new interface.
	NWBrowseResultChangeInterfaceAdded NWBrowseResultChange = 0x8
	// NWBrowseResultChangeInterfaceRemoved: The service was no longer discovered over a certain interface.
	NWBrowseResultChangeInterfaceRemoved NWBrowseResultChange = 0x10
	// NWBrowseResultChangeInvalid: The comparison was invallid.
	NWBrowseResultChangeInvalid NWBrowseResultChange = 0
	// NWBrowseResultChangeResultAdded: A new service was discovered.
	NWBrowseResultChangeResultAdded NWBrowseResultChange = 0x2
	// NWBrowseResultChangeResultRemoved: A previously discovered service was removed.
	NWBrowseResultChangeResultRemoved NWBrowseResultChange = 0x4
	// NWBrowseResultChangeTXTRecordChanged: The service’s associated TXT record changed.
	NWBrowseResultChangeTXTRecordChanged NWBrowseResultChange = 0x20
)

func (e NWBrowseResultChange) String() string {
	switch e {
	case NWBrowseResultChangeIdentical:
		return "NWBrowseResultChangeIdentical"
	case NWBrowseResultChangeInterfaceAdded:
		return "NWBrowseResultChangeInterfaceAdded"
	case NWBrowseResultChangeInterfaceRemoved:
		return "NWBrowseResultChangeInterfaceRemoved"
	case NWBrowseResultChangeInvalid:
		return "NWBrowseResultChangeInvalid"
	case NWBrowseResultChangeResultAdded:
		return "NWBrowseResultChangeResultAdded"
	case NWBrowseResultChangeResultRemoved:
		return "NWBrowseResultChangeResultRemoved"
	case NWBrowseResultChangeTXTRecordChanged:
		return "NWBrowseResultChangeTXTRecordChanged"
	default:
		return fmt.Sprintf("NWBrowseResultChange(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_browser_state_t
type NWBrowserState uint32

const (
	// NWBrowserStateCancelled: The browser has been canceled.
	NWBrowserStateCancelled NWBrowserState = 3
	// NWBrowserStateFailed: The browser has encountered a fatal error.
	NWBrowserStateFailed NWBrowserState = 2
	// NWBrowserStateInvalid: The browser is not valid.
	NWBrowserStateInvalid NWBrowserState = 0
	// NWBrowserStateReady: The browser is registered for discovering services.
	NWBrowserStateReady   NWBrowserState = 1
	NWBrowserStateWaiting NWBrowserState = 4
)

func (e NWBrowserState) String() string {
	switch e {
	case NWBrowserStateCancelled:
		return "NWBrowserStateCancelled"
	case NWBrowserStateFailed:
		return "NWBrowserStateFailed"
	case NWBrowserStateInvalid:
		return "NWBrowserStateInvalid"
	case NWBrowserStateReady:
		return "NWBrowserStateReady"
	case NWBrowserStateWaiting:
		return "NWBrowserStateWaiting"
	default:
		return fmt.Sprintf("NWBrowserState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_connection_group_state_t
type NWConnectionGroupState uint32

const (
	// NWConnectionGroupStateCancelled: The connection group has been canceled.
	NWConnectionGroupStateCancelled NWConnectionGroupState = 4
	// NWConnectionGroupStateFailed: The connection group encountered a fatal error.
	NWConnectionGroupStateFailed NWConnectionGroupState = 3
	// NWConnectionGroupStateInvalid: The connection group is not valid.
	NWConnectionGroupStateInvalid NWConnectionGroupState = 0
	// NWConnectionGroupStateReady: The connection group is joined, and ready to send and receive data.
	NWConnectionGroupStateReady NWConnectionGroupState = 2
	// NWConnectionGroupStateWaiting: The connection group is waiting for a network path change.
	NWConnectionGroupStateWaiting NWConnectionGroupState = 1
)

func (e NWConnectionGroupState) String() string {
	switch e {
	case NWConnectionGroupStateCancelled:
		return "NWConnectionGroupStateCancelled"
	case NWConnectionGroupStateFailed:
		return "NWConnectionGroupStateFailed"
	case NWConnectionGroupStateInvalid:
		return "NWConnectionGroupStateInvalid"
	case NWConnectionGroupStateReady:
		return "NWConnectionGroupStateReady"
	case NWConnectionGroupStateWaiting:
		return "NWConnectionGroupStateWaiting"
	default:
		return fmt.Sprintf("NWConnectionGroupState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_connection_state_t
type NWConnectionState uint32

const (
	// NWConnectionStateCancelled: The connection has been canceled.
	NWConnectionStateCancelled NWConnectionState = 5
	// NWConnectionStateFailed: The connection has disconnected or encountered an error.
	NWConnectionStateFailed NWConnectionState = 4
	// NWConnectionStateInvalid: The connection is not valid.
	NWConnectionStateInvalid NWConnectionState = 0
	// NWConnectionStatePreparing: The connection in the process of being established.
	NWConnectionStatePreparing NWConnectionState = 2
	// NWConnectionStateReady: The connection is established, and ready to send and receive data.
	NWConnectionStateReady NWConnectionState = 3
	// NWConnectionStateWaiting: The connection is waiting for a network path change.
	NWConnectionStateWaiting NWConnectionState = 1
)

func (e NWConnectionState) String() string {
	switch e {
	case NWConnectionStateCancelled:
		return "NWConnectionStateCancelled"
	case NWConnectionStateFailed:
		return "NWConnectionStateFailed"
	case NWConnectionStateInvalid:
		return "NWConnectionStateInvalid"
	case NWConnectionStatePreparing:
		return "NWConnectionStatePreparing"
	case NWConnectionStateReady:
		return "NWConnectionStateReady"
	case NWConnectionStateWaiting:
		return "NWConnectionStateWaiting"
	default:
		return fmt.Sprintf("NWConnectionState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_state_t
type NWDataTransferReportState uint32

const (
	// NWDataTransferReportStateCollected: The data transfer report has completed, and data can be examined.
	NWDataTransferReportStateCollected NWDataTransferReportState = 2
	// NWDataTransferReportStateCollecting: The data transfer report has been started but is still collecting data.
	NWDataTransferReportStateCollecting NWDataTransferReportState = 1
)

func (e NWDataTransferReportState) String() string {
	switch e {
	case NWDataTransferReportStateCollected:
		return "NWDataTransferReportStateCollected"
	case NWDataTransferReportStateCollecting:
		return "NWDataTransferReportStateCollecting"
	default:
		return fmt.Sprintf("NWDataTransferReportState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_endpoint_type_t
type NWEndpointType uint32

const (
	// NWEndpointTypeAddress: An endpoint represented as an IP address and port.
	NWEndpointTypeAddress NWEndpointType = 1
	// NWEndpointTypeBonjourService: An endpoint represented as a Bonjour service.
	NWEndpointTypeBonjourService NWEndpointType = 3
	// NWEndpointTypeHost: An endpoint represented as a hostname and port.
	NWEndpointTypeHost NWEndpointType = 2
	// NWEndpointTypeInvalid: An undefined endpoint type.
	NWEndpointTypeInvalid NWEndpointType = 0
	// NWEndpointTypeURL: An endpoint represented as a URL, with host and port values inferred from the URL.
	NWEndpointTypeURL NWEndpointType = 4
)

func (e NWEndpointType) String() string {
	switch e {
	case NWEndpointTypeAddress:
		return "NWEndpointTypeAddress"
	case NWEndpointTypeBonjourService:
		return "NWEndpointTypeBonjourService"
	case NWEndpointTypeHost:
		return "NWEndpointTypeHost"
	case NWEndpointTypeInvalid:
		return "NWEndpointTypeInvalid"
	case NWEndpointTypeURL:
		return "NWEndpointTypeURL"
	default:
		return fmt.Sprintf("NWEndpointType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_error_domain_t
type NWErrorDomain uint32

const (
	// NWErrorDomainDns: A DNS error encountered in resolving, browsing, or advertising.
	NWErrorDomainDns NWErrorDomain = 2
	// NWErrorDomainInvalid: The error is invalid.
	NWErrorDomainInvalid NWErrorDomain = 0
	// NWErrorDomainPosix: A POSIX error, which is used for most network protocol and routing errors.
	NWErrorDomainPosix NWErrorDomain = 1
	// NWErrorDomainTLS: A TLS error reported by a TLS connection or listener.
	NWErrorDomainTLS       NWErrorDomain = 3
	NWErrorDomainWifiAware NWErrorDomain = 4
)

func (e NWErrorDomain) String() string {
	switch e {
	case NWErrorDomainDns:
		return "NWErrorDomainDns"
	case NWErrorDomainInvalid:
		return "NWErrorDomainInvalid"
	case NWErrorDomainPosix:
		return "NWErrorDomainPosix"
	case NWErrorDomainTLS:
		return "NWErrorDomainTLS"
	case NWErrorDomainWifiAware:
		return "NWErrorDomainWifiAware"
	default:
		return fmt.Sprintf("NWErrorDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_state_t
type NWEthernetChannelState uint32

const (
	// NWEthernetChannelStateCancelled: The channel has been canceled.
	NWEthernetChannelStateCancelled NWEthernetChannelState = 5
	// NWEthernetChannelStateFailed: The channel has encountered a fatal error.
	NWEthernetChannelStateFailed NWEthernetChannelState = 4
	// NWEthernetChannelStateInvalid: The channel is not valid.
	NWEthernetChannelStateInvalid NWEthernetChannelState = 0
	// NWEthernetChannelStatePreparing: The channel is registering with the interface.
	NWEthernetChannelStatePreparing NWEthernetChannelState = 2
	// NWEthernetChannelStateReady: The channel is able to send and receive Ethernet frames.
	NWEthernetChannelStateReady NWEthernetChannelState = 3
	// NWEthernetChannelStateWaiting: The channel is waiting for its interface to become available.
	NWEthernetChannelStateWaiting NWEthernetChannelState = 1
)

func (e NWEthernetChannelState) String() string {
	switch e {
	case NWEthernetChannelStateCancelled:
		return "NWEthernetChannelStateCancelled"
	case NWEthernetChannelStateFailed:
		return "NWEthernetChannelStateFailed"
	case NWEthernetChannelStateInvalid:
		return "NWEthernetChannelStateInvalid"
	case NWEthernetChannelStatePreparing:
		return "NWEthernetChannelStatePreparing"
	case NWEthernetChannelStateReady:
		return "NWEthernetChannelStateReady"
	case NWEthernetChannelStateWaiting:
		return "NWEthernetChannelStateWaiting"
	default:
		return fmt.Sprintf("NWEthernetChannelState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_framer_start_result_t
type NWFramerStartResult uint32

const (
	// NWFramerStartResultReady: The protocol is immediately ready to send and receive data.
	NWFramerStartResultReady NWFramerStartResult = 1
	// NWFramerStartResultWillMarkReady: The protocol will perform a handshake, preventing the overall connection from becoming ready until nw_framer_mark_ready(_:) is called.
	NWFramerStartResultWillMarkReady NWFramerStartResult = 2
)

func (e NWFramerStartResult) String() string {
	switch e {
	case NWFramerStartResultReady:
		return "NWFramerStartResultReady"
	case NWFramerStartResultWillMarkReady:
		return "NWFramerStartResultWillMarkReady"
	default:
		return fmt.Sprintf("NWFramerStartResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_interface_radio_type_t
type NWInterfaceRadioType uint32

const (
	NWInterfaceRadioTypeCellCdma     NWInterfaceRadioType = 0x87
	NWInterfaceRadioTypeCellEndcMmw  NWInterfaceRadioType = 0x82
	NWInterfaceRadioTypeCellEndcSub6 NWInterfaceRadioType = 0x81
	NWInterfaceRadioTypeCellEvdo     NWInterfaceRadioType = 0x88
	NWInterfaceRadioTypeCellGsm      NWInterfaceRadioType = 0x86
	NWInterfaceRadioTypeCellLte      NWInterfaceRadioType = 0x80
	NWInterfaceRadioTypeCellNrSaMmw  NWInterfaceRadioType = 0x84
	NWInterfaceRadioTypeCellNrSaSub6 NWInterfaceRadioType = 0x83
	NWInterfaceRadioTypeCellWcdma    NWInterfaceRadioType = 0x85
	NWInterfaceRadioTypeUnknown      NWInterfaceRadioType = 0
	NWInterfaceRadioTypeWifiA        NWInterfaceRadioType = 2
	NWInterfaceRadioTypeWifiAc       NWInterfaceRadioType = 5
	NWInterfaceRadioTypeWifiAx       NWInterfaceRadioType = 6
	NWInterfaceRadioTypeWifiB        NWInterfaceRadioType = 1
	NWInterfaceRadioTypeWifiG        NWInterfaceRadioType = 3
	NWInterfaceRadioTypeWifiN        NWInterfaceRadioType = 4
)

func (e NWInterfaceRadioType) String() string {
	switch e {
	case NWInterfaceRadioTypeCellCdma:
		return "NWInterfaceRadioTypeCellCdma"
	case NWInterfaceRadioTypeCellEndcMmw:
		return "NWInterfaceRadioTypeCellEndcMmw"
	case NWInterfaceRadioTypeCellEndcSub6:
		return "NWInterfaceRadioTypeCellEndcSub6"
	case NWInterfaceRadioTypeCellEvdo:
		return "NWInterfaceRadioTypeCellEvdo"
	case NWInterfaceRadioTypeCellGsm:
		return "NWInterfaceRadioTypeCellGsm"
	case NWInterfaceRadioTypeCellLte:
		return "NWInterfaceRadioTypeCellLte"
	case NWInterfaceRadioTypeCellNrSaMmw:
		return "NWInterfaceRadioTypeCellNrSaMmw"
	case NWInterfaceRadioTypeCellNrSaSub6:
		return "NWInterfaceRadioTypeCellNrSaSub6"
	case NWInterfaceRadioTypeCellWcdma:
		return "NWInterfaceRadioTypeCellWcdma"
	case NWInterfaceRadioTypeUnknown:
		return "NWInterfaceRadioTypeUnknown"
	case NWInterfaceRadioTypeWifiA:
		return "NWInterfaceRadioTypeWifiA"
	case NWInterfaceRadioTypeWifiAc:
		return "NWInterfaceRadioTypeWifiAc"
	case NWInterfaceRadioTypeWifiAx:
		return "NWInterfaceRadioTypeWifiAx"
	case NWInterfaceRadioTypeWifiB:
		return "NWInterfaceRadioTypeWifiB"
	case NWInterfaceRadioTypeWifiG:
		return "NWInterfaceRadioTypeWifiG"
	case NWInterfaceRadioTypeWifiN:
		return "NWInterfaceRadioTypeWifiN"
	default:
		return fmt.Sprintf("NWInterfaceRadioType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_interface_type_t
type NWInterfaceType uint32

const (
	// NWInterfaceTypeCellular: The network interface type used for communication over cellular networks.
	NWInterfaceTypeCellular NWInterfaceType = 2
	// NWInterfaceTypeLoopback: The network interface type used for communication over local loopback networks.
	NWInterfaceTypeLoopback NWInterfaceType = 4
	// NWInterfaceTypeOther: The network interface type used for communication over virtual networks or networks of unknown types.
	NWInterfaceTypeOther NWInterfaceType = 0
	// NWInterfaceTypeWifi: The network interface type used for communication over Wi-Fi networks.
	NWInterfaceTypeWifi NWInterfaceType = 1
	// NWInterfaceTypeWired: The network interface type used for communication over wired Ethernet networks.
	NWInterfaceTypeWired NWInterfaceType = 3
)

func (e NWInterfaceType) String() string {
	switch e {
	case NWInterfaceTypeCellular:
		return "NWInterfaceTypeCellular"
	case NWInterfaceTypeLoopback:
		return "NWInterfaceTypeLoopback"
	case NWInterfaceTypeOther:
		return "NWInterfaceTypeOther"
	case NWInterfaceTypeWifi:
		return "NWInterfaceTypeWifi"
	case NWInterfaceTypeWired:
		return "NWInterfaceTypeWired"
	default:
		return fmt.Sprintf("NWInterfaceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ip_ecn_flag_t
type NWIPEcnFlag uint32

const (
	// NWIPEcnFlagCe: Congestion Experienced.
	NWIPEcnFlagCe NWIPEcnFlag = 3
	// NWIPEcnFlagEct0: ECN Capable Transport (flag 0).
	NWIPEcnFlagEct0 NWIPEcnFlag = 2
	// NWIPEcnFlagEct1: ECN Capable Transport (flag 1).
	NWIPEcnFlagEct1 NWIPEcnFlag = 1
	// NWIPEcnFlagNonEct: Non-ECN Capable Transport.
	NWIPEcnFlagNonEct NWIPEcnFlag = 0
)

func (e NWIPEcnFlag) String() string {
	switch e {
	case NWIPEcnFlagCe:
		return "NWIPEcnFlagCe"
	case NWIPEcnFlagEct0:
		return "NWIPEcnFlagEct0"
	case NWIPEcnFlagEct1:
		return "NWIPEcnFlagEct1"
	case NWIPEcnFlagNonEct:
		return "NWIPEcnFlagNonEct"
	default:
		return fmt.Sprintf("NWIPEcnFlag(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ip_local_address_preference_t
type NWIPLocalAddressPreference uint32

const (
	// NWIPLocalAddressPreferenceDefault: Allow the system to decide which kind of local address to prefer for a connection or listener.
	NWIPLocalAddressPreferenceDefault NWIPLocalAddressPreference = 0
	// NWIPLocalAddressPreferenceStable: Prefer using stable local addresses.
	NWIPLocalAddressPreferenceStable NWIPLocalAddressPreference = 2
	// NWIPLocalAddressPreferenceTemporary: Prefer using temporary local addresses.
	NWIPLocalAddressPreferenceTemporary NWIPLocalAddressPreference = 1
)

func (e NWIPLocalAddressPreference) String() string {
	switch e {
	case NWIPLocalAddressPreferenceDefault:
		return "NWIPLocalAddressPreferenceDefault"
	case NWIPLocalAddressPreferenceStable:
		return "NWIPLocalAddressPreferenceStable"
	case NWIPLocalAddressPreferenceTemporary:
		return "NWIPLocalAddressPreferenceTemporary"
	default:
		return fmt.Sprintf("NWIPLocalAddressPreference(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ip_version_t
type NWIPVersion uint32

const (
	// NWIPVersion4: Require IP version 4.
	NWIPVersion4 NWIPVersion = 4
	// NWIPVersion6: Require IP version 6.
	NWIPVersion6 NWIPVersion = 6
	// NWIPVersionAny: Allow any IP version.
	NWIPVersionAny NWIPVersion = 0
)

func (e NWIPVersion) String() string {
	switch e {
	case NWIPVersion4:
		return "NWIPVersion4"
	case NWIPVersion6:
		return "NWIPVersion6"
	case NWIPVersionAny:
		return "NWIPVersionAny"
	default:
		return fmt.Sprintf("NWIPVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_link_quality_t
type NWLinkQuality uint32

const (
	NWLinkQualityGood     NWLinkQuality = 30
	NWLinkQualityMinimal  NWLinkQuality = 10
	NWLinkQualityModerate NWLinkQuality = 20
	NWLinkQualityUnknown  NWLinkQuality = 0
)

func (e NWLinkQuality) String() string {
	switch e {
	case NWLinkQualityGood:
		return "NWLinkQualityGood"
	case NWLinkQualityMinimal:
		return "NWLinkQualityMinimal"
	case NWLinkQualityModerate:
		return "NWLinkQualityModerate"
	case NWLinkQualityUnknown:
		return "NWLinkQualityUnknown"
	default:
		return fmt.Sprintf("NWLinkQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_listener_state_t
type NWListenerState uint32

const (
	// NWListenerStateCancelled: The listener has been canceled.
	NWListenerStateCancelled NWListenerState = 4
	// NWListenerStateFailed: The listener has encountered a fatal error.
	NWListenerStateFailed NWListenerState = 3
	// NWListenerStateInvalid: The listener is not valid.
	NWListenerStateInvalid NWListenerState = 0
	// NWListenerStateReady: The listener is running and able to receive incoming connections.
	NWListenerStateReady NWListenerState = 2
	// NWListenerStateWaiting: The listener is waiting for a network to become available.
	NWListenerStateWaiting NWListenerState = 1
)

func (e NWListenerState) String() string {
	switch e {
	case NWListenerStateCancelled:
		return "NWListenerStateCancelled"
	case NWListenerStateFailed:
		return "NWListenerStateFailed"
	case NWListenerStateInvalid:
		return "NWListenerStateInvalid"
	case NWListenerStateReady:
		return "NWListenerStateReady"
	case NWListenerStateWaiting:
		return "NWListenerStateWaiting"
	default:
		return fmt.Sprintf("NWListenerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_multipath_service_t
type NWMultipathService uint32

const (
	// NWMultipathServiceAggregate: Enable multipath to maximize bandwidth across multiple interfaces.
	NWMultipathServiceAggregate NWMultipathService = 3
	// NWMultipathServiceDisabled: Disable multipath.
	NWMultipathServiceDisabled NWMultipathService = 0
	// NWMultipathServiceHandover: Enable multipath, but only use other interfaces when the primary interface is lost.
	NWMultipathServiceHandover NWMultipathService = 1
	// NWMultipathServiceInteractive: Enable multipath to use other interfaces when the primary interface encounters loss or delay.
	NWMultipathServiceInteractive NWMultipathService = 2
)

func (e NWMultipathService) String() string {
	switch e {
	case NWMultipathServiceAggregate:
		return "NWMultipathServiceAggregate"
	case NWMultipathServiceDisabled:
		return "NWMultipathServiceDisabled"
	case NWMultipathServiceHandover:
		return "NWMultipathServiceHandover"
	case NWMultipathServiceInteractive:
		return "NWMultipathServiceInteractive"
	default:
		return fmt.Sprintf("NWMultipathService(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_multipath_version_t
type NWMultipathVersion int32

const (
	NWMultipathVersion0           NWMultipathVersion = 0
	NWMultipathVersion1           NWMultipathVersion = 1
	NWMultipathVersionUnspecified NWMultipathVersion = -1
)

func (e NWMultipathVersion) String() string {
	switch e {
	case NWMultipathVersion0:
		return "NWMultipathVersion0"
	case NWMultipathVersion1:
		return "NWMultipathVersion1"
	case NWMultipathVersionUnspecified:
		return "NWMultipathVersionUnspecified"
	default:
		return fmt.Sprintf("NWMultipathVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_parameters_attribution_t
type NWParametersAttribution uint8

const (
	// NWParametersAttributionDeveloper: A developer-initiated network request.
	NWParametersAttributionDeveloper NWParametersAttribution = 1
	// NWParametersAttributionUser: The user explicitly directs the app to make a network request.
	NWParametersAttributionUser NWParametersAttribution = 2
)

func (e NWParametersAttribution) String() string {
	switch e {
	case NWParametersAttributionDeveloper:
		return "NWParametersAttributionDeveloper"
	case NWParametersAttributionUser:
		return "NWParametersAttributionUser"
	default:
		return fmt.Sprintf("NWParametersAttribution(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_parameters_expired_dns_behavior_t
type NWParametersExpiredDnsBehavior uint32

const (
	// NWParametersExpiredDnsBehaviorAllow: Explicitly allow the use of expired DNS answers.
	NWParametersExpiredDnsBehaviorAllow NWParametersExpiredDnsBehavior = 1
	// NWParametersExpiredDnsBehaviorDefault: Let the system determine whether or not to allow expired DNS answers.
	NWParametersExpiredDnsBehaviorDefault    NWParametersExpiredDnsBehavior = 0
	NWParametersExpiredDnsBehaviorPersistent NWParametersExpiredDnsBehavior = 3
	// NWParametersExpiredDnsBehaviorProhibit: Explicitly prohibit the use of expired DNS answers.
	NWParametersExpiredDnsBehaviorProhibit NWParametersExpiredDnsBehavior = 2
)

func (e NWParametersExpiredDnsBehavior) String() string {
	switch e {
	case NWParametersExpiredDnsBehaviorAllow:
		return "NWParametersExpiredDnsBehaviorAllow"
	case NWParametersExpiredDnsBehaviorDefault:
		return "NWParametersExpiredDnsBehaviorDefault"
	case NWParametersExpiredDnsBehaviorPersistent:
		return "NWParametersExpiredDnsBehaviorPersistent"
	case NWParametersExpiredDnsBehaviorProhibit:
		return "NWParametersExpiredDnsBehaviorProhibit"
	default:
		return fmt.Sprintf("NWParametersExpiredDnsBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_path_status_t
type NWPathStatus uint32

const (
	// NWPathStatusInvalid: The path is not valid.
	NWPathStatusInvalid NWPathStatus = 0
	// NWPathStatusSatisfiable: The path is not currently available, but establishing a new connection may activate the path.
	NWPathStatusSatisfiable NWPathStatus = 3
	// NWPathStatusSatisfied: The path is available to establish connections and send data.
	NWPathStatusSatisfied NWPathStatus = 1
	// NWPathStatusUnsatisfied: The path is not available for use.
	NWPathStatusUnsatisfied NWPathStatus = 2
)

func (e NWPathStatus) String() string {
	switch e {
	case NWPathStatusInvalid:
		return "NWPathStatusInvalid"
	case NWPathStatusSatisfiable:
		return "NWPathStatusSatisfiable"
	case NWPathStatusSatisfied:
		return "NWPathStatusSatisfied"
	case NWPathStatusUnsatisfied:
		return "NWPathStatusUnsatisfied"
	default:
		return fmt.Sprintf("NWPathStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_path_unsatisfied_reason_t
type NWPathUnsatisfiedReason uint32

const (
	NWPathUnsatisfiedReasonCellularDenied     NWPathUnsatisfiedReason = 1
	NWPathUnsatisfiedReasonLocalNetworkDenied NWPathUnsatisfiedReason = 3
	NWPathUnsatisfiedReasonNotAvailable       NWPathUnsatisfiedReason = 0
	NWPathUnsatisfiedReasonVpnInactive        NWPathUnsatisfiedReason = 4
	NWPathUnsatisfiedReasonWifiDenied         NWPathUnsatisfiedReason = 2
)

func (e NWPathUnsatisfiedReason) String() string {
	switch e {
	case NWPathUnsatisfiedReasonCellularDenied:
		return "NWPathUnsatisfiedReasonCellularDenied"
	case NWPathUnsatisfiedReasonLocalNetworkDenied:
		return "NWPathUnsatisfiedReasonLocalNetworkDenied"
	case NWPathUnsatisfiedReasonNotAvailable:
		return "NWPathUnsatisfiedReasonNotAvailable"
	case NWPathUnsatisfiedReasonVpnInactive:
		return "NWPathUnsatisfiedReasonVpnInactive"
	case NWPathUnsatisfiedReasonWifiDenied:
		return "NWPathUnsatisfiedReasonWifiDenied"
	default:
		return fmt.Sprintf("NWPathUnsatisfiedReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_quic_stream_type_t
type NWQuicStreamType uint32

const (
	NWQuicStreamTypeBidirectional  NWQuicStreamType = 1
	NWQuicStreamTypeDatagram       NWQuicStreamType = 3
	NWQuicStreamTypeUnidirectional NWQuicStreamType = 2
	NWQuicStreamTypeUnknown        NWQuicStreamType = 0
)

func (e NWQuicStreamType) String() string {
	switch e {
	case NWQuicStreamTypeBidirectional:
		return "NWQuicStreamTypeBidirectional"
	case NWQuicStreamTypeDatagram:
		return "NWQuicStreamTypeDatagram"
	case NWQuicStreamTypeUnidirectional:
		return "NWQuicStreamTypeUnidirectional"
	case NWQuicStreamTypeUnknown:
		return "NWQuicStreamTypeUnknown"
	default:
		return fmt.Sprintf("NWQuicStreamType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_report_resolution_protocol_t
type NWReportResolutionProtocol uint32

const (
	// NWReportResolutionProtocolHttps: The connection used HTTPS for DNS resolution.
	NWReportResolutionProtocolHttps NWReportResolutionProtocol = 4
	// NWReportResolutionProtocolTCP: The connection used cleartext TCP for DNS resolution.
	NWReportResolutionProtocolTCP NWReportResolutionProtocol = 2
	// NWReportResolutionProtocolTLS: The connection used TLS for DNS resolution.
	NWReportResolutionProtocolTLS NWReportResolutionProtocol = 3
	// NWReportResolutionProtocolUDP: The connection used cleartext UDP for DNS resolution.
	NWReportResolutionProtocolUDP NWReportResolutionProtocol = 1
	// NWReportResolutionProtocolUnknown: The DNS response protocol is unknown or not applicable.
	NWReportResolutionProtocolUnknown NWReportResolutionProtocol = 0
)

func (e NWReportResolutionProtocol) String() string {
	switch e {
	case NWReportResolutionProtocolHttps:
		return "NWReportResolutionProtocolHttps"
	case NWReportResolutionProtocolTCP:
		return "NWReportResolutionProtocolTCP"
	case NWReportResolutionProtocolTLS:
		return "NWReportResolutionProtocolTLS"
	case NWReportResolutionProtocolUDP:
		return "NWReportResolutionProtocolUDP"
	case NWReportResolutionProtocolUnknown:
		return "NWReportResolutionProtocolUnknown"
	default:
		return fmt.Sprintf("NWReportResolutionProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_report_resolution_source_t
type NWReportResolutionSource uint32

const (
	// NWReportResolutionSourceCache: The DNS response was retrieved from a local cache.
	NWReportResolutionSourceCache NWReportResolutionSource = 2
	// NWReportResolutionSourceExpiredCache: The DNS response had expired and was retrieved from a local cache.
	NWReportResolutionSourceExpiredCache NWReportResolutionSource = 3
	// NWReportResolutionSourceQuery: The DNS response was received from the network.
	NWReportResolutionSourceQuery NWReportResolutionSource = 1
)

func (e NWReportResolutionSource) String() string {
	switch e {
	case NWReportResolutionSourceCache:
		return "NWReportResolutionSourceCache"
	case NWReportResolutionSourceExpiredCache:
		return "NWReportResolutionSourceExpiredCache"
	case NWReportResolutionSourceQuery:
		return "NWReportResolutionSourceQuery"
	default:
		return fmt.Sprintf("NWReportResolutionSource(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_service_class_t
type NWServiceClass uint32

const (
	// NWServiceClassBackground: Bulk traffic, or traffic that can be deprioritized behind foreground traffic.
	NWServiceClassBackground NWServiceClass = 1
	// NWServiceClassBestEffort: Default priority traffic.
	NWServiceClassBestEffort NWServiceClass = 0
	// NWServiceClassInteractiveVideo: Interactive video traffic.
	NWServiceClassInteractiveVideo NWServiceClass = 2
	// NWServiceClassInteractiveVoice: Interactive voice traffic.
	NWServiceClassInteractiveVoice NWServiceClass = 3
	// NWServiceClassResponsiveData: Responsive user-data traffic.
	NWServiceClassResponsiveData NWServiceClass = 4
	// NWServiceClassSignaling: Signaling control traffic.
	NWServiceClassSignaling NWServiceClass = 5
)

func (e NWServiceClass) String() string {
	switch e {
	case NWServiceClassBackground:
		return "NWServiceClassBackground"
	case NWServiceClassBestEffort:
		return "NWServiceClassBestEffort"
	case NWServiceClassInteractiveVideo:
		return "NWServiceClassInteractiveVideo"
	case NWServiceClassInteractiveVoice:
		return "NWServiceClassInteractiveVoice"
	case NWServiceClassResponsiveData:
		return "NWServiceClassResponsiveData"
	case NWServiceClassSignaling:
		return "NWServiceClassSignaling"
	default:
		return fmt.Sprintf("NWServiceClass(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_txt_record_find_key_t
type NWTXTRecordFindKey uint32

const (
	// NWTXTRecordFindKeyEmptyValue: The key is present and has an empty associated value.
	NWTXTRecordFindKeyEmptyValue NWTXTRecordFindKey = 3
	// NWTXTRecordFindKeyInvalid: The key is not valid.
	NWTXTRecordFindKeyInvalid NWTXTRecordFindKey = 0
	// NWTXTRecordFindKeyNoValue: The key is present but has no associated value.
	NWTXTRecordFindKeyNoValue NWTXTRecordFindKey = 2
	// NWTXTRecordFindKeyNonEmptyValue: The key has an associated value.
	NWTXTRecordFindKeyNonEmptyValue NWTXTRecordFindKey = 4
	// NWTXTRecordFindKeyNotPresent: The key is not present in the dictionary.
	NWTXTRecordFindKeyNotPresent NWTXTRecordFindKey = 1
)

func (e NWTXTRecordFindKey) String() string {
	switch e {
	case NWTXTRecordFindKeyEmptyValue:
		return "NWTXTRecordFindKeyEmptyValue"
	case NWTXTRecordFindKeyInvalid:
		return "NWTXTRecordFindKeyInvalid"
	case NWTXTRecordFindKeyNoValue:
		return "NWTXTRecordFindKeyNoValue"
	case NWTXTRecordFindKeyNonEmptyValue:
		return "NWTXTRecordFindKeyNonEmptyValue"
	case NWTXTRecordFindKeyNotPresent:
		return "NWTXTRecordFindKeyNotPresent"
	default:
		return fmt.Sprintf("NWTXTRecordFindKey(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ws_close_code_t
type NWWsCloseCode uint32

const (
	// NWWsCloseCodeAbnormalClosure: This value is reserved for local errors and indicates that no Close message was received.
	NWWsCloseCodeAbnormalClosure NWWsCloseCode = 1006
	// NWWsCloseCodeGoingAway: An endpoint is no longer available, such as when a server is down.
	NWWsCloseCodeGoingAway NWWsCloseCode = 1001
	// NWWsCloseCodeInternalServerError: The server is terminating the connection because it encountered an unexpected condition that prevented it from fulfilling the request.
	NWWsCloseCodeInternalServerError NWWsCloseCode = 1011
	// NWWsCloseCodeInvalidFramePayloadData: An endpoint is terminating the connection because it received data within a message that was inconsistent with the message type.
	NWWsCloseCodeInvalidFramePayloadData NWWsCloseCode = 1007
	// NWWsCloseCodeMandatoryExtension: The WebSocket client expected the server to negotiate one or more extensions that were not negotiated.
	NWWsCloseCodeMandatoryExtension NWWsCloseCode = 1010
	// NWWsCloseCodeMessageTooBig: An endpoint is terminating the connection because it received a message that is too big for it to process.
	NWWsCloseCodeMessageTooBig NWWsCloseCode = 1009
	// NWWsCloseCodeNoStatusReceived: This value is reserved for local errors and indicates that no Close code was received.
	NWWsCloseCodeNoStatusReceived NWWsCloseCode = 1005
	// NWWsCloseCodeNormalClosure: A normal closure occurred with no errors.
	NWWsCloseCodeNormalClosure NWWsCloseCode = 1000
	// NWWsCloseCodePolicyViolation: An endpoint is terminating the connection because it received a message that violates its policy.
	NWWsCloseCodePolicyViolation NWWsCloseCode = 1008
	// NWWsCloseCodeProtocolError: An endpoint is terminating the connection due to a protocol error.
	NWWsCloseCodeProtocolError NWWsCloseCode = 1002
	// NWWsCloseCodeTLSHandshake: This value is reserved for local errors and indicates that the TLS handshake failed.
	NWWsCloseCodeTLSHandshake NWWsCloseCode = 1015
	// NWWsCloseCodeUnsupportedData: An endpoint is terminating the connection because it received a type of data it cannot accept.
	NWWsCloseCodeUnsupportedData NWWsCloseCode = 1003
)

func (e NWWsCloseCode) String() string {
	switch e {
	case NWWsCloseCodeAbnormalClosure:
		return "NWWsCloseCodeAbnormalClosure"
	case NWWsCloseCodeGoingAway:
		return "NWWsCloseCodeGoingAway"
	case NWWsCloseCodeInternalServerError:
		return "NWWsCloseCodeInternalServerError"
	case NWWsCloseCodeInvalidFramePayloadData:
		return "NWWsCloseCodeInvalidFramePayloadData"
	case NWWsCloseCodeMandatoryExtension:
		return "NWWsCloseCodeMandatoryExtension"
	case NWWsCloseCodeMessageTooBig:
		return "NWWsCloseCodeMessageTooBig"
	case NWWsCloseCodeNoStatusReceived:
		return "NWWsCloseCodeNoStatusReceived"
	case NWWsCloseCodeNormalClosure:
		return "NWWsCloseCodeNormalClosure"
	case NWWsCloseCodePolicyViolation:
		return "NWWsCloseCodePolicyViolation"
	case NWWsCloseCodeProtocolError:
		return "NWWsCloseCodeProtocolError"
	case NWWsCloseCodeTLSHandshake:
		return "NWWsCloseCodeTLSHandshake"
	case NWWsCloseCodeUnsupportedData:
		return "NWWsCloseCodeUnsupportedData"
	default:
		return fmt.Sprintf("NWWsCloseCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ws_opcode_t
type NWWsOpcode int32

const (
	// NWWsOpcodeBinary: A binary data message.
	NWWsOpcodeBinary NWWsOpcode = 0x2
	// NWWsOpcodeClose: A message indicating a close of the connection.
	NWWsOpcodeClose NWWsOpcode = 0x8
	// NWWsOpcodeCont: A continuation message.
	NWWsOpcodeCont NWWsOpcode = 0
	// NWWsOpcodeInvalid: The message is not valid.
	NWWsOpcodeInvalid NWWsOpcode = -1
	// NWWsOpcodePing: A Ping message, which requests a Pong from the peer.
	NWWsOpcodePing NWWsOpcode = 0x9
	// NWWsOpcodePong: A Pong message in response to a Ping from the peer.
	NWWsOpcodePong NWWsOpcode = 0xa
	// NWWsOpcodeText: A text data message.
	NWWsOpcodeText NWWsOpcode = 0x1
)

func (e NWWsOpcode) String() string {
	switch e {
	case NWWsOpcodeBinary:
		return "NWWsOpcodeBinary"
	case NWWsOpcodeClose:
		return "NWWsOpcodeClose"
	case NWWsOpcodeCont:
		return "NWWsOpcodeCont"
	case NWWsOpcodeInvalid:
		return "NWWsOpcodeInvalid"
	case NWWsOpcodePing:
		return "NWWsOpcodePing"
	case NWWsOpcodePong:
		return "NWWsOpcodePong"
	case NWWsOpcodeText:
		return "NWWsOpcodeText"
	default:
		return fmt.Sprintf("NWWsOpcode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ws_response_status_t
type NWWsResponseStatus uint32

const (
	// NWWsResponseStatusAccept: The client request is being accepted.
	NWWsResponseStatusAccept NWWsResponseStatus = 1
	// NWWsResponseStatusInvalid: An invalid response status.
	NWWsResponseStatusInvalid NWWsResponseStatus = 0
	// NWWsResponseStatusReject: The client request is being rejected.
	NWWsResponseStatusReject NWWsResponseStatus = 2
)

func (e NWWsResponseStatus) String() string {
	switch e {
	case NWWsResponseStatusAccept:
		return "NWWsResponseStatusAccept"
	case NWWsResponseStatusInvalid:
		return "NWWsResponseStatusInvalid"
	case NWWsResponseStatusReject:
		return "NWWsResponseStatusReject"
	default:
		return fmt.Sprintf("NWWsResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_ws_version_t
type NWWsVersion uint32

const (
	// NWWsVersion13: Version 13 of the WebSocket protocol.
	NWWsVersion13 NWWsVersion = 1
	// NWWsVersionInvalid: An invalid version.
	NWWsVersionInvalid NWWsVersion = 0
)

func (e NWWsVersion) String() string {
	switch e {
	case NWWsVersion13:
		return "NWWsVersion13"
	case NWWsVersionInvalid:
		return "NWWsVersionInvalid"
	default:
		return fmt.Sprintf("NWWsVersion(%d)", e)
	}
}

// Nw_browser_state_t is a C-name alias for NWBrowserState.
type Nw_browser_state_t = NWBrowserState

// Nw_connection_group_state_t is a C-name alias for NWConnectionGroupState.
type Nw_connection_group_state_t = NWConnectionGroupState

// Nw_connection_state_t is a C-name alias for NWConnectionState.
type Nw_connection_state_t = NWConnectionState

// Nw_data_transfer_report_state_t is a C-name alias for NWDataTransferReportState.
type Nw_data_transfer_report_state_t = NWDataTransferReportState

// Nw_endpoint_type_t is a C-name alias for NWEndpointType.
type Nw_endpoint_type_t = NWEndpointType

// Nw_error_domain_t is a C-name alias for NWErrorDomain.
type Nw_error_domain_t = NWErrorDomain

// Nw_ethernet_channel_state_t is a C-name alias for NWEthernetChannelState.
type Nw_ethernet_channel_state_t = NWEthernetChannelState

// Nw_framer_start_result_t is a C-name alias for NWFramerStartResult.
type Nw_framer_start_result_t = NWFramerStartResult

// Nw_interface_radio_type_t is a C-name alias for NWInterfaceRadioType.
type Nw_interface_radio_type_t = NWInterfaceRadioType

// Nw_interface_type_t is a C-name alias for NWInterfaceType.
type Nw_interface_type_t = NWInterfaceType

// Nw_ip_ecn_flag_t is a C-name alias for NWIPEcnFlag.
type Nw_ip_ecn_flag_t = NWIPEcnFlag

// Nw_ip_local_address_preference_t is a C-name alias for NWIPLocalAddressPreference.
type Nw_ip_local_address_preference_t = NWIPLocalAddressPreference

// Nw_ip_version_t is a C-name alias for NWIPVersion.
type Nw_ip_version_t = NWIPVersion

// Nw_link_quality_t is a C-name alias for NWLinkQuality.
type Nw_link_quality_t = NWLinkQuality

// Nw_listener_state_t is a C-name alias for NWListenerState.
type Nw_listener_state_t = NWListenerState

// Nw_multipath_service_t is a C-name alias for NWMultipathService.
type Nw_multipath_service_t = NWMultipathService

// Nw_multipath_version_t is a C-name alias for NWMultipathVersion.
type Nw_multipath_version_t = NWMultipathVersion

// Nw_parameters_attribution_t is a C-name alias for NWParametersAttribution.
type Nw_parameters_attribution_t = NWParametersAttribution

// Nw_parameters_expired_dns_behavior_t is a C-name alias for NWParametersExpiredDnsBehavior.
type Nw_parameters_expired_dns_behavior_t = NWParametersExpiredDnsBehavior

// Nw_path_status_t is a C-name alias for NWPathStatus.
type Nw_path_status_t = NWPathStatus

// Nw_path_unsatisfied_reason_t is a C-name alias for NWPathUnsatisfiedReason.
type Nw_path_unsatisfied_reason_t = NWPathUnsatisfiedReason

// Nw_quic_stream_type_t is a C-name alias for NWQuicStreamType.
type Nw_quic_stream_type_t = NWQuicStreamType

// Nw_report_resolution_protocol_t is a C-name alias for NWReportResolutionProtocol.
type Nw_report_resolution_protocol_t = NWReportResolutionProtocol

// Nw_report_resolution_source_t is a C-name alias for NWReportResolutionSource.
type Nw_report_resolution_source_t = NWReportResolutionSource

// Nw_service_class_t is a C-name alias for NWServiceClass.
type Nw_service_class_t = NWServiceClass

// Nw_txt_record_find_key_t is a C-name alias for NWTXTRecordFindKey.
type Nw_txt_record_find_key_t = NWTXTRecordFindKey

// Nw_ws_close_code_t is a C-name alias for NWWsCloseCode.
type Nw_ws_close_code_t = NWWsCloseCode

// Nw_ws_opcode_t is a C-name alias for NWWsOpcode.
type Nw_ws_opcode_t = NWWsOpcode

// Nw_ws_response_status_t is a C-name alias for NWWsResponseStatus.
type Nw_ws_response_status_t = NWWsResponseStatus

// Nw_ws_version_t is a C-name alias for NWWsVersion.
type Nw_ws_version_t = NWWsVersion
