// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"fmt"
)

type NwBrowseResultChange int

const (
	// Nw_browse_result_change_identical: The compared services are identical.
	Nw_browse_result_change_identical NwBrowseResultChange = 1
	// Nw_browse_result_change_interface_added: The service was discovered over a new interface.
	Nw_browse_result_change_interface_added NwBrowseResultChange = 8
	// Nw_browse_result_change_interface_removed: The service was no longer discovered over a certain interface.
	Nw_browse_result_change_interface_removed NwBrowseResultChange = 16
	// Nw_browse_result_change_invalid: The comparison was invallid.
	Nw_browse_result_change_invalid NwBrowseResultChange = 0
	// Nw_browse_result_change_result_added: A new service was discovered.
	Nw_browse_result_change_result_added NwBrowseResultChange = 2
	// Nw_browse_result_change_result_removed: A previously discovered service was removed.
	Nw_browse_result_change_result_removed NwBrowseResultChange = 4
	// Nw_browse_result_change_txt_record_changed: The service’s associated TXT record changed.
	Nw_browse_result_change_txt_record_changed NwBrowseResultChange = 32
)

func (e NwBrowseResultChange) String() string {
	switch e {
	case Nw_browse_result_change_identical:
		return "Nw_browse_result_change_identical"
	case Nw_browse_result_change_interface_added:
		return "Nw_browse_result_change_interface_added"
	case Nw_browse_result_change_interface_removed:
		return "Nw_browse_result_change_interface_removed"
	case Nw_browse_result_change_invalid:
		return "Nw_browse_result_change_invalid"
	case Nw_browse_result_change_result_added:
		return "Nw_browse_result_change_result_added"
	case Nw_browse_result_change_result_removed:
		return "Nw_browse_result_change_result_removed"
	case Nw_browse_result_change_txt_record_changed:
		return "Nw_browse_result_change_txt_record_changed"
	default:
		return fmt.Sprintf("NwBrowseResultChange(%d)", e)
	}
}

type NwBrowserState uint

const (
	// Nw_browser_state_cancelled: The browser has been canceled.
	Nw_browser_state_cancelled NwBrowserState = 3
	// Nw_browser_state_failed: The browser has encountered a fatal error.
	Nw_browser_state_failed NwBrowserState = 2
	// Nw_browser_state_invalid: The browser is not valid.
	Nw_browser_state_invalid NwBrowserState = 0
	// Nw_browser_state_ready: The browser is registered for discovering services.
	Nw_browser_state_ready   NwBrowserState = 1
	Nw_browser_state_waiting NwBrowserState = 4
)

func (e NwBrowserState) String() string {
	switch e {
	case Nw_browser_state_cancelled:
		return "Nw_browser_state_cancelled"
	case Nw_browser_state_failed:
		return "Nw_browser_state_failed"
	case Nw_browser_state_invalid:
		return "Nw_browser_state_invalid"
	case Nw_browser_state_ready:
		return "Nw_browser_state_ready"
	case Nw_browser_state_waiting:
		return "Nw_browser_state_waiting"
	default:
		return fmt.Sprintf("NwBrowserState(%d)", e)
	}
}

type NwConnectionGroupState uint

const (
	// Nw_connection_group_state_cancelled: The connection group has been canceled.
	Nw_connection_group_state_cancelled NwConnectionGroupState = 4
	// Nw_connection_group_state_failed: The connection group encountered a fatal error.
	Nw_connection_group_state_failed NwConnectionGroupState = 3
	// Nw_connection_group_state_invalid: The connection group is not valid.
	Nw_connection_group_state_invalid NwConnectionGroupState = 0
	// Nw_connection_group_state_ready: The connection group is joined, and ready to send and receive data.
	Nw_connection_group_state_ready NwConnectionGroupState = 2
	// Nw_connection_group_state_waiting: The connection group is waiting for a network path change.
	Nw_connection_group_state_waiting NwConnectionGroupState = 1
)

func (e NwConnectionGroupState) String() string {
	switch e {
	case Nw_connection_group_state_cancelled:
		return "Nw_connection_group_state_cancelled"
	case Nw_connection_group_state_failed:
		return "Nw_connection_group_state_failed"
	case Nw_connection_group_state_invalid:
		return "Nw_connection_group_state_invalid"
	case Nw_connection_group_state_ready:
		return "Nw_connection_group_state_ready"
	case Nw_connection_group_state_waiting:
		return "Nw_connection_group_state_waiting"
	default:
		return fmt.Sprintf("NwConnectionGroupState(%d)", e)
	}
}

type NwConnectionState uint

const (
	// Nw_connection_state_cancelled: The connection has been canceled.
	Nw_connection_state_cancelled NwConnectionState = 5
	// Nw_connection_state_failed: The connection has disconnected or encountered an error.
	Nw_connection_state_failed NwConnectionState = 4
	// Nw_connection_state_invalid: The connection is not valid.
	Nw_connection_state_invalid NwConnectionState = 0
	// Nw_connection_state_preparing: The connection in the process of being established.
	Nw_connection_state_preparing NwConnectionState = 2
	// Nw_connection_state_ready: The connection is established, and ready to send and receive data.
	Nw_connection_state_ready NwConnectionState = 3
	// Nw_connection_state_waiting: The connection is waiting for a network path change.
	Nw_connection_state_waiting NwConnectionState = 1
)

func (e NwConnectionState) String() string {
	switch e {
	case Nw_connection_state_cancelled:
		return "Nw_connection_state_cancelled"
	case Nw_connection_state_failed:
		return "Nw_connection_state_failed"
	case Nw_connection_state_invalid:
		return "Nw_connection_state_invalid"
	case Nw_connection_state_preparing:
		return "Nw_connection_state_preparing"
	case Nw_connection_state_ready:
		return "Nw_connection_state_ready"
	case Nw_connection_state_waiting:
		return "Nw_connection_state_waiting"
	default:
		return fmt.Sprintf("NwConnectionState(%d)", e)
	}
}

type NwDataTransferReportState uint

const (
	// Nw_data_transfer_report_state_collected: The data transfer report has completed, and data can be examined.
	Nw_data_transfer_report_state_collected NwDataTransferReportState = 2
	// Nw_data_transfer_report_state_collecting: The data transfer report has been started but is still collecting data.
	Nw_data_transfer_report_state_collecting NwDataTransferReportState = 1
)

func (e NwDataTransferReportState) String() string {
	switch e {
	case Nw_data_transfer_report_state_collected:
		return "Nw_data_transfer_report_state_collected"
	case Nw_data_transfer_report_state_collecting:
		return "Nw_data_transfer_report_state_collecting"
	default:
		return fmt.Sprintf("NwDataTransferReportState(%d)", e)
	}
}

type NwEndpointType uint

const (
	// Nw_endpoint_type_address: An endpoint represented as an IP address and port.
	Nw_endpoint_type_address NwEndpointType = 1
	// Nw_endpoint_type_bonjour_service: An endpoint represented as a Bonjour service.
	Nw_endpoint_type_bonjour_service NwEndpointType = 3
	// Nw_endpoint_type_host: An endpoint represented as a hostname and port.
	Nw_endpoint_type_host NwEndpointType = 2
	// Nw_endpoint_type_invalid: An undefined endpoint type.
	Nw_endpoint_type_invalid NwEndpointType = 0
	// Nw_endpoint_type_url: An endpoint represented as a URL, with host and port values inferred from the URL.
	Nw_endpoint_type_url NwEndpointType = 4
)

func (e NwEndpointType) String() string {
	switch e {
	case Nw_endpoint_type_address:
		return "Nw_endpoint_type_address"
	case Nw_endpoint_type_bonjour_service:
		return "Nw_endpoint_type_bonjour_service"
	case Nw_endpoint_type_host:
		return "Nw_endpoint_type_host"
	case Nw_endpoint_type_invalid:
		return "Nw_endpoint_type_invalid"
	case Nw_endpoint_type_url:
		return "Nw_endpoint_type_url"
	default:
		return fmt.Sprintf("NwEndpointType(%d)", e)
	}
}

type NwErrorDomain int

const (
	// Nw_error_domain_dns: A DNS error encountered in resolving, browsing, or advertising.
	Nw_error_domain_dns NwErrorDomain = 2
	// Nw_error_domain_invalid: The error is invalid.
	Nw_error_domain_invalid NwErrorDomain = 0
	// Nw_error_domain_posix: A POSIX error, which is used for most network protocol and routing errors.
	Nw_error_domain_posix NwErrorDomain = 1
	// Nw_error_domain_tls: A TLS error reported by a TLS connection or listener.
	Nw_error_domain_tls        NwErrorDomain = 3
	Nw_error_domain_wifi_aware NwErrorDomain = 4
)

func (e NwErrorDomain) String() string {
	switch e {
	case Nw_error_domain_dns:
		return "Nw_error_domain_dns"
	case Nw_error_domain_invalid:
		return "Nw_error_domain_invalid"
	case Nw_error_domain_posix:
		return "Nw_error_domain_posix"
	case Nw_error_domain_tls:
		return "Nw_error_domain_tls"
	case Nw_error_domain_wifi_aware:
		return "Nw_error_domain_wifi_aware"
	default:
		return fmt.Sprintf("NwErrorDomain(%d)", e)
	}
}

type NwEthernetChannelState uint

const (
	// Nw_ethernet_channel_state_cancelled: The channel has been canceled.
	Nw_ethernet_channel_state_cancelled NwEthernetChannelState = 5
	// Nw_ethernet_channel_state_failed: The channel has encountered a fatal error.
	Nw_ethernet_channel_state_failed NwEthernetChannelState = 4
	// Nw_ethernet_channel_state_invalid: The channel is not valid.
	Nw_ethernet_channel_state_invalid NwEthernetChannelState = 0
	// Nw_ethernet_channel_state_preparing: The channel is registering with the interface.
	Nw_ethernet_channel_state_preparing NwEthernetChannelState = 2
	// Nw_ethernet_channel_state_ready: The channel is able to send and receive Ethernet frames.
	Nw_ethernet_channel_state_ready NwEthernetChannelState = 3
	// Nw_ethernet_channel_state_waiting: The channel is waiting for its interface to become available.
	Nw_ethernet_channel_state_waiting NwEthernetChannelState = 1
)

func (e NwEthernetChannelState) String() string {
	switch e {
	case Nw_ethernet_channel_state_cancelled:
		return "Nw_ethernet_channel_state_cancelled"
	case Nw_ethernet_channel_state_failed:
		return "Nw_ethernet_channel_state_failed"
	case Nw_ethernet_channel_state_invalid:
		return "Nw_ethernet_channel_state_invalid"
	case Nw_ethernet_channel_state_preparing:
		return "Nw_ethernet_channel_state_preparing"
	case Nw_ethernet_channel_state_ready:
		return "Nw_ethernet_channel_state_ready"
	case Nw_ethernet_channel_state_waiting:
		return "Nw_ethernet_channel_state_waiting"
	default:
		return fmt.Sprintf("NwEthernetChannelState(%d)", e)
	}
}

type NwFramerStartResult int

const (
	// Nw_framer_start_result_ready: The protocol is immediately ready to send and receive data.
	Nw_framer_start_result_ready NwFramerStartResult = 1
	// Nw_framer_start_result_will_mark_ready: The protocol will perform a handshake, preventing the overall connection from becoming ready until nw_framer_mark_ready(_:) is called.
	Nw_framer_start_result_will_mark_ready NwFramerStartResult = 2
)

func (e NwFramerStartResult) String() string {
	switch e {
	case Nw_framer_start_result_ready:
		return "Nw_framer_start_result_ready"
	case Nw_framer_start_result_will_mark_ready:
		return "Nw_framer_start_result_will_mark_ready"
	default:
		return fmt.Sprintf("NwFramerStartResult(%d)", e)
	}
}

type NwInterfaceRadioType uint

const (
	Nw_interface_radio_type_cell_cdma       NwInterfaceRadioType = 135
	Nw_interface_radio_type_cell_endc_mmw   NwInterfaceRadioType = 130
	Nw_interface_radio_type_cell_endc_sub6  NwInterfaceRadioType = 129
	Nw_interface_radio_type_cell_evdo       NwInterfaceRadioType = 136
	Nw_interface_radio_type_cell_gsm        NwInterfaceRadioType = 134
	Nw_interface_radio_type_cell_lte        NwInterfaceRadioType = 128
	Nw_interface_radio_type_cell_nr_sa_mmw  NwInterfaceRadioType = 132
	Nw_interface_radio_type_cell_nr_sa_sub6 NwInterfaceRadioType = 131
	Nw_interface_radio_type_cell_wcdma      NwInterfaceRadioType = 133
	Nw_interface_radio_type_unknown         NwInterfaceRadioType = 0
	Nw_interface_radio_type_wifi_a          NwInterfaceRadioType = 2
	Nw_interface_radio_type_wifi_ac         NwInterfaceRadioType = 5
	Nw_interface_radio_type_wifi_ax         NwInterfaceRadioType = 6
	Nw_interface_radio_type_wifi_b          NwInterfaceRadioType = 1
	Nw_interface_radio_type_wifi_g          NwInterfaceRadioType = 3
	Nw_interface_radio_type_wifi_n          NwInterfaceRadioType = 4
)

func (e NwInterfaceRadioType) String() string {
	switch e {
	case Nw_interface_radio_type_cell_cdma:
		return "Nw_interface_radio_type_cell_cdma"
	case Nw_interface_radio_type_cell_endc_mmw:
		return "Nw_interface_radio_type_cell_endc_mmw"
	case Nw_interface_radio_type_cell_endc_sub6:
		return "Nw_interface_radio_type_cell_endc_sub6"
	case Nw_interface_radio_type_cell_evdo:
		return "Nw_interface_radio_type_cell_evdo"
	case Nw_interface_radio_type_cell_gsm:
		return "Nw_interface_radio_type_cell_gsm"
	case Nw_interface_radio_type_cell_lte:
		return "Nw_interface_radio_type_cell_lte"
	case Nw_interface_radio_type_cell_nr_sa_mmw:
		return "Nw_interface_radio_type_cell_nr_sa_mmw"
	case Nw_interface_radio_type_cell_nr_sa_sub6:
		return "Nw_interface_radio_type_cell_nr_sa_sub6"
	case Nw_interface_radio_type_cell_wcdma:
		return "Nw_interface_radio_type_cell_wcdma"
	case Nw_interface_radio_type_unknown:
		return "Nw_interface_radio_type_unknown"
	case Nw_interface_radio_type_wifi_a:
		return "Nw_interface_radio_type_wifi_a"
	case Nw_interface_radio_type_wifi_ac:
		return "Nw_interface_radio_type_wifi_ac"
	case Nw_interface_radio_type_wifi_ax:
		return "Nw_interface_radio_type_wifi_ax"
	case Nw_interface_radio_type_wifi_b:
		return "Nw_interface_radio_type_wifi_b"
	case Nw_interface_radio_type_wifi_g:
		return "Nw_interface_radio_type_wifi_g"
	case Nw_interface_radio_type_wifi_n:
		return "Nw_interface_radio_type_wifi_n"
	default:
		return fmt.Sprintf("NwInterfaceRadioType(%d)", e)
	}
}

type NwInterfaceType uint

const (
	// Nw_interface_type_cellular: The network interface type used for communication over cellular networks.
	Nw_interface_type_cellular NwInterfaceType = 2
	// Nw_interface_type_loopback: The network interface type used for communication over local loopback networks.
	Nw_interface_type_loopback NwInterfaceType = 4
	// Nw_interface_type_other: The network interface type used for communication over virtual networks or networks of unknown types.
	Nw_interface_type_other NwInterfaceType = 0
	// Nw_interface_type_wifi: The network interface type used for communication over Wi-Fi networks.
	Nw_interface_type_wifi NwInterfaceType = 1
	// Nw_interface_type_wired: The network interface type used for communication over wired Ethernet networks.
	Nw_interface_type_wired NwInterfaceType = 3
)

func (e NwInterfaceType) String() string {
	switch e {
	case Nw_interface_type_cellular:
		return "Nw_interface_type_cellular"
	case Nw_interface_type_loopback:
		return "Nw_interface_type_loopback"
	case Nw_interface_type_other:
		return "Nw_interface_type_other"
	case Nw_interface_type_wifi:
		return "Nw_interface_type_wifi"
	case Nw_interface_type_wired:
		return "Nw_interface_type_wired"
	default:
		return fmt.Sprintf("NwInterfaceType(%d)", e)
	}
}

type NwIpEcnFlag uint

const (
	// Nw_ip_ecn_flag_ce: Congestion Experienced.
	Nw_ip_ecn_flag_ce NwIpEcnFlag = 3
	// Nw_ip_ecn_flag_ect_0: ECN Capable Transport (flag 0).
	Nw_ip_ecn_flag_ect_0 NwIpEcnFlag = 2
	// Nw_ip_ecn_flag_ect_1: ECN Capable Transport (flag 1).
	Nw_ip_ecn_flag_ect_1 NwIpEcnFlag = 1
	// Nw_ip_ecn_flag_non_ect: Non-ECN Capable Transport.
	Nw_ip_ecn_flag_non_ect NwIpEcnFlag = 0
)

func (e NwIpEcnFlag) String() string {
	switch e {
	case Nw_ip_ecn_flag_ce:
		return "Nw_ip_ecn_flag_ce"
	case Nw_ip_ecn_flag_ect_0:
		return "Nw_ip_ecn_flag_ect_0"
	case Nw_ip_ecn_flag_ect_1:
		return "Nw_ip_ecn_flag_ect_1"
	case Nw_ip_ecn_flag_non_ect:
		return "Nw_ip_ecn_flag_non_ect"
	default:
		return fmt.Sprintf("NwIpEcnFlag(%d)", e)
	}
}

type NwIpLocalAddressPreference uint

const (
	// Nw_ip_local_address_preference_default: Allow the system to decide which kind of local address to prefer for a connection or listener.
	Nw_ip_local_address_preference_default NwIpLocalAddressPreference = 0
	// Nw_ip_local_address_preference_stable: Prefer using stable local addresses.
	Nw_ip_local_address_preference_stable NwIpLocalAddressPreference = 2
	// Nw_ip_local_address_preference_temporary: Prefer using temporary local addresses.
	Nw_ip_local_address_preference_temporary NwIpLocalAddressPreference = 1
)

func (e NwIpLocalAddressPreference) String() string {
	switch e {
	case Nw_ip_local_address_preference_default:
		return "Nw_ip_local_address_preference_default"
	case Nw_ip_local_address_preference_stable:
		return "Nw_ip_local_address_preference_stable"
	case Nw_ip_local_address_preference_temporary:
		return "Nw_ip_local_address_preference_temporary"
	default:
		return fmt.Sprintf("NwIpLocalAddressPreference(%d)", e)
	}
}

type NwIpVersion uint

const (
	// Nw_ip_version_4: Require IP version 4.
	Nw_ip_version_4 NwIpVersion = 4
	// Nw_ip_version_6: Require IP version 6.
	Nw_ip_version_6 NwIpVersion = 6
	// Nw_ip_version_any: Allow any IP version.
	Nw_ip_version_any NwIpVersion = 0
)

func (e NwIpVersion) String() string {
	switch e {
	case Nw_ip_version_4:
		return "Nw_ip_version_4"
	case Nw_ip_version_6:
		return "Nw_ip_version_6"
	case Nw_ip_version_any:
		return "Nw_ip_version_any"
	default:
		return fmt.Sprintf("NwIpVersion(%d)", e)
	}
}

type NwLinkQuality uint

const (
	Nw_link_quality_good     NwLinkQuality = 30
	Nw_link_quality_minimal  NwLinkQuality = 10
	Nw_link_quality_moderate NwLinkQuality = 20
	Nw_link_quality_unknown  NwLinkQuality = 0
)

func (e NwLinkQuality) String() string {
	switch e {
	case Nw_link_quality_good:
		return "Nw_link_quality_good"
	case Nw_link_quality_minimal:
		return "Nw_link_quality_minimal"
	case Nw_link_quality_moderate:
		return "Nw_link_quality_moderate"
	case Nw_link_quality_unknown:
		return "Nw_link_quality_unknown"
	default:
		return fmt.Sprintf("NwLinkQuality(%d)", e)
	}
}

type NwListenerState uint

const (
	// Nw_listener_state_cancelled: The listener has been canceled.
	Nw_listener_state_cancelled NwListenerState = 4
	// Nw_listener_state_failed: The listener has encountered a fatal error.
	Nw_listener_state_failed NwListenerState = 3
	// Nw_listener_state_invalid: The listener is not valid.
	Nw_listener_state_invalid NwListenerState = 0
	// Nw_listener_state_ready: The listener is running and able to receive incoming connections.
	Nw_listener_state_ready NwListenerState = 2
	// Nw_listener_state_waiting: The listener is waiting for a network to become available.
	Nw_listener_state_waiting NwListenerState = 1
)

func (e NwListenerState) String() string {
	switch e {
	case Nw_listener_state_cancelled:
		return "Nw_listener_state_cancelled"
	case Nw_listener_state_failed:
		return "Nw_listener_state_failed"
	case Nw_listener_state_invalid:
		return "Nw_listener_state_invalid"
	case Nw_listener_state_ready:
		return "Nw_listener_state_ready"
	case Nw_listener_state_waiting:
		return "Nw_listener_state_waiting"
	default:
		return fmt.Sprintf("NwListenerState(%d)", e)
	}
}

type NwMultipathService uint

const (
	// Nw_multipath_service_aggregate: Enable multipath to maximize bandwidth across multiple interfaces.
	Nw_multipath_service_aggregate NwMultipathService = 3
	// Nw_multipath_service_disabled: Disable multipath.
	Nw_multipath_service_disabled NwMultipathService = 0
	// Nw_multipath_service_handover: Enable multipath, but only use other interfaces when the primary interface is lost.
	Nw_multipath_service_handover NwMultipathService = 1
	// Nw_multipath_service_interactive: Enable multipath to use other interfaces when the primary interface encounters loss or delay.
	Nw_multipath_service_interactive NwMultipathService = 2
)

func (e NwMultipathService) String() string {
	switch e {
	case Nw_multipath_service_aggregate:
		return "Nw_multipath_service_aggregate"
	case Nw_multipath_service_disabled:
		return "Nw_multipath_service_disabled"
	case Nw_multipath_service_handover:
		return "Nw_multipath_service_handover"
	case Nw_multipath_service_interactive:
		return "Nw_multipath_service_interactive"
	default:
		return fmt.Sprintf("NwMultipathService(%d)", e)
	}
}

type NwMultipathVersion int

const (
	Nw_multipath_version_0           NwMultipathVersion = 0
	Nw_multipath_version_1           NwMultipathVersion = 1
	Nw_multipath_version_unspecified NwMultipathVersion = -1
)

func (e NwMultipathVersion) String() string {
	switch e {
	case Nw_multipath_version_0:
		return "Nw_multipath_version_0"
	case Nw_multipath_version_1:
		return "Nw_multipath_version_1"
	case Nw_multipath_version_unspecified:
		return "Nw_multipath_version_unspecified"
	default:
		return fmt.Sprintf("NwMultipathVersion(%d)", e)
	}
}

type NwParametersExpiredDnsBehavior uint

const (
	// Nw_parameters_expired_dns_behavior_allow: Explicitly allow the use of expired DNS answers.
	Nw_parameters_expired_dns_behavior_allow NwParametersExpiredDnsBehavior = 1
	// Nw_parameters_expired_dns_behavior_default: Let the system determine whether or not to allow expired DNS answers.
	Nw_parameters_expired_dns_behavior_default    NwParametersExpiredDnsBehavior = 0
	Nw_parameters_expired_dns_behavior_persistent NwParametersExpiredDnsBehavior = 3
	// Nw_parameters_expired_dns_behavior_prohibit: Explicitly prohibit the use of expired DNS answers.
	Nw_parameters_expired_dns_behavior_prohibit NwParametersExpiredDnsBehavior = 2
)

func (e NwParametersExpiredDnsBehavior) String() string {
	switch e {
	case Nw_parameters_expired_dns_behavior_allow:
		return "Nw_parameters_expired_dns_behavior_allow"
	case Nw_parameters_expired_dns_behavior_default:
		return "Nw_parameters_expired_dns_behavior_default"
	case Nw_parameters_expired_dns_behavior_persistent:
		return "Nw_parameters_expired_dns_behavior_persistent"
	case Nw_parameters_expired_dns_behavior_prohibit:
		return "Nw_parameters_expired_dns_behavior_prohibit"
	default:
		return fmt.Sprintf("NwParametersExpiredDnsBehavior(%d)", e)
	}
}

type NwPathStatus int

const (
	// Nw_path_status_invalid: The path is not valid.
	Nw_path_status_invalid NwPathStatus = 0
	// Nw_path_status_satisfiable: The path is not currently available, but establishing a new connection may activate the path.
	Nw_path_status_satisfiable NwPathStatus = 3
	// Nw_path_status_satisfied: The path is available to establish connections and send data.
	Nw_path_status_satisfied NwPathStatus = 1
	// Nw_path_status_unsatisfied: The path is not available for use.
	Nw_path_status_unsatisfied NwPathStatus = 2
)

func (e NwPathStatus) String() string {
	switch e {
	case Nw_path_status_invalid:
		return "Nw_path_status_invalid"
	case Nw_path_status_satisfiable:
		return "Nw_path_status_satisfiable"
	case Nw_path_status_satisfied:
		return "Nw_path_status_satisfied"
	case Nw_path_status_unsatisfied:
		return "Nw_path_status_unsatisfied"
	default:
		return fmt.Sprintf("NwPathStatus(%d)", e)
	}
}

type NwPathUnsatisfiedReason uint

const (
	Nw_path_unsatisfied_reason_cellular_denied      NwPathUnsatisfiedReason = 1
	Nw_path_unsatisfied_reason_local_network_denied NwPathUnsatisfiedReason = 3
	Nw_path_unsatisfied_reason_not_available        NwPathUnsatisfiedReason = 0
	Nw_path_unsatisfied_reason_vpn_inactive         NwPathUnsatisfiedReason = 4
	Nw_path_unsatisfied_reason_wifi_denied          NwPathUnsatisfiedReason = 2
)

func (e NwPathUnsatisfiedReason) String() string {
	switch e {
	case Nw_path_unsatisfied_reason_cellular_denied:
		return "Nw_path_unsatisfied_reason_cellular_denied"
	case Nw_path_unsatisfied_reason_local_network_denied:
		return "Nw_path_unsatisfied_reason_local_network_denied"
	case Nw_path_unsatisfied_reason_not_available:
		return "Nw_path_unsatisfied_reason_not_available"
	case Nw_path_unsatisfied_reason_vpn_inactive:
		return "Nw_path_unsatisfied_reason_vpn_inactive"
	case Nw_path_unsatisfied_reason_wifi_denied:
		return "Nw_path_unsatisfied_reason_wifi_denied"
	default:
		return fmt.Sprintf("NwPathUnsatisfiedReason(%d)", e)
	}
}

type NwQuicStreamType uint

const (
	Nw_quic_stream_type_bidirectional  NwQuicStreamType = 1
	Nw_quic_stream_type_datagram       NwQuicStreamType = 3
	Nw_quic_stream_type_unidirectional NwQuicStreamType = 2
	Nw_quic_stream_type_unknown        NwQuicStreamType = 0
)

func (e NwQuicStreamType) String() string {
	switch e {
	case Nw_quic_stream_type_bidirectional:
		return "Nw_quic_stream_type_bidirectional"
	case Nw_quic_stream_type_datagram:
		return "Nw_quic_stream_type_datagram"
	case Nw_quic_stream_type_unidirectional:
		return "Nw_quic_stream_type_unidirectional"
	case Nw_quic_stream_type_unknown:
		return "Nw_quic_stream_type_unknown"
	default:
		return fmt.Sprintf("NwQuicStreamType(%d)", e)
	}
}

type NwReportResolutionProtocol uint

const (
	// Nw_report_resolution_protocol_https: The connection used HTTPS for DNS resolution.
	Nw_report_resolution_protocol_https NwReportResolutionProtocol = 4
	// Nw_report_resolution_protocol_tcp: The connection used cleartext TCP for DNS resolution.
	Nw_report_resolution_protocol_tcp NwReportResolutionProtocol = 2
	// Nw_report_resolution_protocol_tls: The connection used TLS for DNS resolution.
	Nw_report_resolution_protocol_tls NwReportResolutionProtocol = 3
	// Nw_report_resolution_protocol_udp: The connection used cleartext UDP for DNS resolution.
	Nw_report_resolution_protocol_udp NwReportResolutionProtocol = 1
	// Nw_report_resolution_protocol_unknown: The DNS response protocol is unknown or not applicable.
	Nw_report_resolution_protocol_unknown NwReportResolutionProtocol = 0
)

func (e NwReportResolutionProtocol) String() string {
	switch e {
	case Nw_report_resolution_protocol_https:
		return "Nw_report_resolution_protocol_https"
	case Nw_report_resolution_protocol_tcp:
		return "Nw_report_resolution_protocol_tcp"
	case Nw_report_resolution_protocol_tls:
		return "Nw_report_resolution_protocol_tls"
	case Nw_report_resolution_protocol_udp:
		return "Nw_report_resolution_protocol_udp"
	case Nw_report_resolution_protocol_unknown:
		return "Nw_report_resolution_protocol_unknown"
	default:
		return fmt.Sprintf("NwReportResolutionProtocol(%d)", e)
	}
}

type NwReportResolutionSource uint

const (
	// Nw_report_resolution_source_cache: The DNS response was retrieved from a local cache.
	Nw_report_resolution_source_cache NwReportResolutionSource = 2
	// Nw_report_resolution_source_expired_cache: The DNS response had expired and was retrieved from a local cache.
	Nw_report_resolution_source_expired_cache NwReportResolutionSource = 3
	// Nw_report_resolution_source_query: The DNS response was received from the network.
	Nw_report_resolution_source_query NwReportResolutionSource = 1
)

func (e NwReportResolutionSource) String() string {
	switch e {
	case Nw_report_resolution_source_cache:
		return "Nw_report_resolution_source_cache"
	case Nw_report_resolution_source_expired_cache:
		return "Nw_report_resolution_source_expired_cache"
	case Nw_report_resolution_source_query:
		return "Nw_report_resolution_source_query"
	default:
		return fmt.Sprintf("NwReportResolutionSource(%d)", e)
	}
}

type NwServiceClass uint

const (
	// Nw_service_class_background: Bulk traffic, or traffic that can be deprioritized behind foreground traffic.
	Nw_service_class_background NwServiceClass = 1
	// Nw_service_class_best_effort: Default priority traffic.
	Nw_service_class_best_effort NwServiceClass = 0
	// Nw_service_class_interactive_video: Interactive video traffic.
	Nw_service_class_interactive_video NwServiceClass = 2
	// Nw_service_class_interactive_voice: Interactive voice traffic.
	Nw_service_class_interactive_voice NwServiceClass = 3
	// Nw_service_class_responsive_data: Responsive user-data traffic.
	Nw_service_class_responsive_data NwServiceClass = 4
	// Nw_service_class_signaling: Signaling control traffic.
	Nw_service_class_signaling NwServiceClass = 5
)

func (e NwServiceClass) String() string {
	switch e {
	case Nw_service_class_background:
		return "Nw_service_class_background"
	case Nw_service_class_best_effort:
		return "Nw_service_class_best_effort"
	case Nw_service_class_interactive_video:
		return "Nw_service_class_interactive_video"
	case Nw_service_class_interactive_voice:
		return "Nw_service_class_interactive_voice"
	case Nw_service_class_responsive_data:
		return "Nw_service_class_responsive_data"
	case Nw_service_class_signaling:
		return "Nw_service_class_signaling"
	default:
		return fmt.Sprintf("NwServiceClass(%d)", e)
	}
}

type NwTxtRecordFindKey uint

const (
	// Nw_txt_record_find_key_empty_value: The key is present and has an empty associated value.
	Nw_txt_record_find_key_empty_value NwTxtRecordFindKey = 3
	// Nw_txt_record_find_key_invalid: The key is not valid.
	Nw_txt_record_find_key_invalid NwTxtRecordFindKey = 0
	// Nw_txt_record_find_key_no_value: The key is present but has no associated value.
	Nw_txt_record_find_key_no_value NwTxtRecordFindKey = 2
	// Nw_txt_record_find_key_non_empty_value: The key has an associated value.
	Nw_txt_record_find_key_non_empty_value NwTxtRecordFindKey = 4
	// Nw_txt_record_find_key_not_present: The key is not present in the dictionary.
	Nw_txt_record_find_key_not_present NwTxtRecordFindKey = 1
)

func (e NwTxtRecordFindKey) String() string {
	switch e {
	case Nw_txt_record_find_key_empty_value:
		return "Nw_txt_record_find_key_empty_value"
	case Nw_txt_record_find_key_invalid:
		return "Nw_txt_record_find_key_invalid"
	case Nw_txt_record_find_key_no_value:
		return "Nw_txt_record_find_key_no_value"
	case Nw_txt_record_find_key_non_empty_value:
		return "Nw_txt_record_find_key_non_empty_value"
	case Nw_txt_record_find_key_not_present:
		return "Nw_txt_record_find_key_not_present"
	default:
		return fmt.Sprintf("NwTxtRecordFindKey(%d)", e)
	}
}

type NwWsCloseCode uint

const (
	// Nw_ws_close_code_abnormal_closure: This value is reserved for local errors and indicates that no Close message was received.
	Nw_ws_close_code_abnormal_closure NwWsCloseCode = 1006
	// Nw_ws_close_code_going_away: An endpoint is no longer available, such as when a server is down.
	Nw_ws_close_code_going_away NwWsCloseCode = 1001
	// Nw_ws_close_code_internal_server_error: The server is terminating the connection because it encountered an unexpected condition that prevented it from fulfilling the request.
	Nw_ws_close_code_internal_server_error NwWsCloseCode = 1011
	// Nw_ws_close_code_invalid_frame_payload_data: An endpoint is terminating the connection because it received data within a message that was inconsistent with the message type.
	Nw_ws_close_code_invalid_frame_payload_data NwWsCloseCode = 1007
	// Nw_ws_close_code_mandatory_extension: The WebSocket client expected the server to negotiate one or more extensions that were not negotiated.
	Nw_ws_close_code_mandatory_extension NwWsCloseCode = 1010
	// Nw_ws_close_code_message_too_big: An endpoint is terminating the connection because it received a message that is too big for it to process.
	Nw_ws_close_code_message_too_big NwWsCloseCode = 1009
	// Nw_ws_close_code_no_status_received: This value is reserved for local errors and indicates that no Close code was received.
	Nw_ws_close_code_no_status_received NwWsCloseCode = 1005
	// Nw_ws_close_code_normal_closure: A normal closure occurred with no errors.
	Nw_ws_close_code_normal_closure NwWsCloseCode = 1000
	// Nw_ws_close_code_policy_violation: An endpoint is terminating the connection because it received a message that violates its policy.
	Nw_ws_close_code_policy_violation NwWsCloseCode = 1008
	// Nw_ws_close_code_protocol_error: An endpoint is terminating the connection due to a protocol error.
	Nw_ws_close_code_protocol_error NwWsCloseCode = 1002
	// Nw_ws_close_code_tls_handshake: This value is reserved for local errors and indicates that the TLS handshake failed.
	Nw_ws_close_code_tls_handshake NwWsCloseCode = 1015
	// Nw_ws_close_code_unsupported_data: An endpoint is terminating the connection because it received a type of data it cannot accept.
	Nw_ws_close_code_unsupported_data NwWsCloseCode = 1003
)

func (e NwWsCloseCode) String() string {
	switch e {
	case Nw_ws_close_code_abnormal_closure:
		return "Nw_ws_close_code_abnormal_closure"
	case Nw_ws_close_code_going_away:
		return "Nw_ws_close_code_going_away"
	case Nw_ws_close_code_internal_server_error:
		return "Nw_ws_close_code_internal_server_error"
	case Nw_ws_close_code_invalid_frame_payload_data:
		return "Nw_ws_close_code_invalid_frame_payload_data"
	case Nw_ws_close_code_mandatory_extension:
		return "Nw_ws_close_code_mandatory_extension"
	case Nw_ws_close_code_message_too_big:
		return "Nw_ws_close_code_message_too_big"
	case Nw_ws_close_code_no_status_received:
		return "Nw_ws_close_code_no_status_received"
	case Nw_ws_close_code_normal_closure:
		return "Nw_ws_close_code_normal_closure"
	case Nw_ws_close_code_policy_violation:
		return "Nw_ws_close_code_policy_violation"
	case Nw_ws_close_code_protocol_error:
		return "Nw_ws_close_code_protocol_error"
	case Nw_ws_close_code_tls_handshake:
		return "Nw_ws_close_code_tls_handshake"
	case Nw_ws_close_code_unsupported_data:
		return "Nw_ws_close_code_unsupported_data"
	default:
		return fmt.Sprintf("NwWsCloseCode(%d)", e)
	}
}

type NwWsOpcode int

const (
	// Nw_ws_opcode_binary: A binary data message.
	Nw_ws_opcode_binary NwWsOpcode = 2
	// Nw_ws_opcode_close: A message indicating a close of the connection.
	Nw_ws_opcode_close NwWsOpcode = 8
	// Nw_ws_opcode_cont: A continuation message.
	Nw_ws_opcode_cont NwWsOpcode = 0
	// Nw_ws_opcode_invalid: The message is not valid.
	Nw_ws_opcode_invalid NwWsOpcode = -1
	// Nw_ws_opcode_ping: A Ping message, which requests a Pong from the peer.
	Nw_ws_opcode_ping NwWsOpcode = 9
	// Nw_ws_opcode_pong: A Pong message in response to a Ping from the peer.
	Nw_ws_opcode_pong NwWsOpcode = 10
	// Nw_ws_opcode_text: A text data message.
	Nw_ws_opcode_text NwWsOpcode = 1
)

func (e NwWsOpcode) String() string {
	switch e {
	case Nw_ws_opcode_binary:
		return "Nw_ws_opcode_binary"
	case Nw_ws_opcode_close:
		return "Nw_ws_opcode_close"
	case Nw_ws_opcode_cont:
		return "Nw_ws_opcode_cont"
	case Nw_ws_opcode_invalid:
		return "Nw_ws_opcode_invalid"
	case Nw_ws_opcode_ping:
		return "Nw_ws_opcode_ping"
	case Nw_ws_opcode_pong:
		return "Nw_ws_opcode_pong"
	case Nw_ws_opcode_text:
		return "Nw_ws_opcode_text"
	default:
		return fmt.Sprintf("NwWsOpcode(%d)", e)
	}
}

type NwWsResponseStatus int

const (
	// Nw_ws_response_status_accept: The client request is being accepted.
	Nw_ws_response_status_accept NwWsResponseStatus = 1
	// Nw_ws_response_status_invalid: An invalid response status.
	Nw_ws_response_status_invalid NwWsResponseStatus = 0
	// Nw_ws_response_status_reject: The client request is being rejected.
	Nw_ws_response_status_reject NwWsResponseStatus = 2
)

func (e NwWsResponseStatus) String() string {
	switch e {
	case Nw_ws_response_status_accept:
		return "Nw_ws_response_status_accept"
	case Nw_ws_response_status_invalid:
		return "Nw_ws_response_status_invalid"
	case Nw_ws_response_status_reject:
		return "Nw_ws_response_status_reject"
	default:
		return fmt.Sprintf("NwWsResponseStatus(%d)", e)
	}
}

type NwWsVersion uint

const (
	// Nw_ws_version_13: Version 13 of the WebSocket protocol.
	Nw_ws_version_13 NwWsVersion = 1
	// Nw_ws_version_invalid: An invalid version.
	Nw_ws_version_invalid NwWsVersion = 0
)

func (e NwWsVersion) String() string {
	switch e {
	case Nw_ws_version_13:
		return "Nw_ws_version_13"
	case Nw_ws_version_invalid:
		return "Nw_ws_version_invalid"
	default:
		return fmt.Sprintf("NwWsVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Network/nw_parameters_attribution_t
type Nw_parameters_attribution_t uint8

const (
	// Nw_parameters_attribution_developer: A developer-initiated network request.
	Nw_parameters_attribution_developer Nw_parameters_attribution_t = 1
	// Nw_parameters_attribution_user: The user explicitly directs the app to make a network request.
	Nw_parameters_attribution_user Nw_parameters_attribution_t = 2
)

func (e Nw_parameters_attribution_t) String() string {
	switch e {
	case Nw_parameters_attribution_developer:
		return "Nw_parameters_attribution_developer"
	case Nw_parameters_attribution_user:
		return "Nw_parameters_attribution_user"
	default:
		return fmt.Sprintf("Nw_parameters_attribution_t(%d)", e)
	}
}
