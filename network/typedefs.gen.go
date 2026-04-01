// Code generated from Apple documentation. DO NOT EDIT.

package network

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// Nw_advertise_descriptor_t is a description used to advertise the Bonjour service that a listener provides.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_t
type Nw_advertise_descriptor_t = objectivec.Object

// Nw_browse_descriptor_t is a service description used to discover Bonjour services.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_t
type Nw_browse_descriptor_t = objectivec.Object

// Nw_browse_result_change_t is flags describing ways in which discovered services can change between specific results.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_change_t
type Nw_browse_result_change_t = uint64

// Nw_browse_result_enumerate_interface_t is a handler that enumerates the interfaces associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interface_t
type Nw_browse_result_enumerate_interface_t = func(objectivec.Object) bool

// Nw_browse_result_t is a discovered service and metadata about the service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_t
type Nw_browse_result_t = objectivec.Object

// Nw_browser_browse_results_changed_handler_t is a handler that delivers updates about discovered services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_browse_results_changed_handler_t
type Nw_browser_browse_results_changed_handler_t = func(objectivec.Object, objectivec.Object, bool)

// Nw_browser_state_changed_handler_t is a handler that delivers browser state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_state_changed_handler_t
type Nw_browser_state_changed_handler_t = func(objectivec.IObject, objectivec.Object)

// Nw_browser_t is an object you use to browse for available network services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_t
type Nw_browser_t = objectivec.Object

// Nw_connection_boolean_event_handler_t is a handler that receives Boolean state updates from a connection, such as viability and better path state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_boolean_event_handler_t
type Nw_connection_boolean_event_handler_t = func(bool)

// See: https://developer.apple.com/documentation/Network/nw_connection_group_new_connection_handler_t
type Nw_connection_group_new_connection_handler_t = func(objectivec.Object)

// Nw_connection_group_receive_handler_t is a handler that receives inbound messages from members of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_receive_handler_t
type Nw_connection_group_receive_handler_t = func(objectivec.Object, objectivec.Object, bool)

// Nw_connection_group_send_completion_t is a completion to notify you when data has been processed and sent.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_send_completion_t
type Nw_connection_group_send_completion_t = func(objectivec.Object)

// Nw_connection_group_state_changed_handler_t is a handler that receives connection group state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_state_changed_handler_t
type Nw_connection_group_state_changed_handler_t = func(objectivec.IObject, objectivec.Object)

// Nw_connection_group_t is an object you use to communicate with a group of endpoints, such as an IP multicast group on a local network.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_t
type Nw_connection_group_t = objectivec.Object

// Nw_connection_path_event_handler_t is a handler that delivers network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_path_event_handler_t
type Nw_connection_path_event_handler_t = func(objectivec.Object)

// Nw_connection_receive_completion_t is a completion handler that indicates when content has been received by the connection, or that an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive_completion_t
type Nw_connection_receive_completion_t = func(objectivec.Object, objectivec.Object, bool, objectivec.Object)

// Nw_connection_send_completion_t is a completion handler that indicates when the connection has finished processing sent content.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_send_completion_t
type Nw_connection_send_completion_t = func(objectivec.Object)

// Nw_connection_state_changed_handler_t is a handler that delivers connection state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_state_changed_handler_t
type Nw_connection_state_changed_handler_t = func(objectivec.IObject, objectivec.Object)

// Nw_connection_t is a bidirectional data connection between a local endpoint and a remote endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_t
type Nw_connection_t = objectivec.Object

// Nw_content_context_t is a representation of a message to send or receive, containing protocol metadata and send properties.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_t
type Nw_content_context_t = objectivec.Object

// Nw_data_transfer_report_collect_block_t is a block that is delivered when a data transfer report is fully collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect_block_t
type Nw_data_transfer_report_collect_block_t = func(objectivec.Object)

// Nw_data_transfer_report_t is a report that provides metrics about data being sent and received on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_t
type Nw_data_transfer_report_t = objectivec.Object

// Nw_endpoint_t is a local or remote endpoint in a network connection.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_t
type Nw_endpoint_t = objectivec.Object

// Nw_error_t is the errors returned by the Network framework.
//
// See: https://developer.apple.com/documentation/Network/nw_error_t
type Nw_error_t = objectivec.Object

// Nw_establishment_report_access_block_t is a block that delivers a connection’s establishment report when it’s in the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_access_block_t
type Nw_establishment_report_access_block_t = func(objectivec.Object)

// Nw_establishment_report_t is a report that provides metrics about how a connection was established.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_t
type Nw_establishment_report_t = objectivec.Object

// Nw_ethernet_address_t is a 48-bit Ethernet address.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_address_t
type Nw_ethernet_address_t = unsafe.Pointer

// Nw_ethernet_channel_receive_handler_t is a handler that delivers inbound Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_receive_handler_t
type Nw_ethernet_channel_receive_handler_t = func(objectivec.Object, uint32, *uint8, *uint8)

// Nw_ethernet_channel_send_completion_t is a handler that indicates when an Ethernet frame has been sent, or if an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send_completion_t
type Nw_ethernet_channel_send_completion_t = func(objectivec.Object)

// Nw_ethernet_channel_state_changed_handler_t is a handler that delivers Ethernet channel state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_state_changed_handler_t
type Nw_ethernet_channel_state_changed_handler_t = func(objectivec.IObject, objectivec.Object)

// Nw_ethernet_channel_t is an object you use to send and receive custom Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_t
type Nw_ethernet_channel_t = objectivec.Object

// Nw_framer_block_t is a block to be invoked asynchronously on your framer protocol’s scheduling context.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_block_t
type Nw_framer_block_t = func()

// Nw_framer_cleanup_handler_t is a handler that tells your protocol to clean up all allocations before being deallocated.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_cleanup_handler_t
type Nw_framer_cleanup_handler_t = func(objectivec.Object)

// Nw_framer_input_handler_t is a handler that notifies your protocol that new inbound data is available to parse.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_input_handler_t
type Nw_framer_input_handler_t = func(objectivec.Object) uint64

// Nw_framer_message_dispose_value_t is a handler that’s invoked when your custom value needs to be released due to a message being released or the value being replaced.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_dispose_value_t
type Nw_framer_message_dispose_value_t = func(unsafe.Pointer)

// Nw_framer_message_t is a message for a custom protocol, in which you can store arbitrary key-value pairs.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_t
type Nw_framer_message_t = Nw_protocol_metadata_t

// Nw_framer_output_handler_t is a handler that notifies your protocol about a new outbound message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_output_handler_t
type Nw_framer_output_handler_t = func(objectivec.Object, objectivec.Object, uint32, bool)

// Nw_framer_parse_completion_t is a handler that examines a range of data being sent or received.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_completion_t
type Nw_framer_parse_completion_t = func(*uint8, uint32, bool) uint64

// Nw_framer_start_handler_t is a handler that represents the entry point into your custom protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_start_handler_t
type Nw_framer_start_handler_t = func(objectivec.Object) objectivec.IObject

// Nw_framer_stop_handler_t is a handler that requests that your protocol send any final messages to close the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_stop_handler_t
type Nw_framer_stop_handler_t = func(objectivec.Object) bool

// Nw_framer_t is an object that represents a single instance of your custom protocol running in a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_t
type Nw_framer_t = objectivec.Object

// Nw_framer_wakeup_handler_t is a handler that delivers a scheduled wakeup event.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_wakeup_handler_t
type Nw_framer_wakeup_handler_t = func(objectivec.Object)

// Nw_group_descriptor_enumerate_endpoints_block_t is a handler that lists all endpoints added to the group descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints_block_t
type Nw_group_descriptor_enumerate_endpoints_block_t = func(objectivec.Object) bool

// Nw_group_descriptor_t is a type that defines a group of endpoints with which you can communicate, such as a multicast group.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_t
type Nw_group_descriptor_t = objectivec.Object

// Nw_interface_t is an interface that a network connection uses to send and receive data.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_t
type Nw_interface_t = objectivec.Object

// Nw_listener_advertised_endpoint_changed_handler_t is a handler that indicates changes to the service endpoints being advertised as they are added and removed.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_advertised_endpoint_changed_handler_t
type Nw_listener_advertised_endpoint_changed_handler_t = func(objectivec.Object, bool)

// See: https://developer.apple.com/documentation/Network/nw_listener_new_connection_group_handler_t
type Nw_listener_new_connection_group_handler_t = func(objectivec.Object)

// Nw_listener_new_connection_handler_t is a handler that delivers inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_new_connection_handler_t
type Nw_listener_new_connection_handler_t = func(objectivec.Object)

// Nw_listener_state_changed_handler_t is a handler that delivers listener state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_state_changed_handler_t
type Nw_listener_state_changed_handler_t = func(objectivec.IObject, objectivec.Object)

// Nw_listener_t is an object you use to listen for incoming network connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_t
type Nw_listener_t = objectivec.Object

// Nw_object_t is the generic type for objects in the Network framework.
//
// See: https://developer.apple.com/documentation/Network/nw_object_t
type Nw_object_t = objectivec.Object

// Nw_parameters_configure_protocol_block_t is a block to configure protocol options during the creation of a parameters object.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_configure_protocol_block_t
type Nw_parameters_configure_protocol_block_t = func(objectivec.Object)

// Nw_parameters_iterate_interface_types_block_t is a block that allows inspection of a list of interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_interface_types_block_t
type Nw_parameters_iterate_interface_types_block_t = func(objectivec.IObject) bool

// Nw_parameters_iterate_interfaces_block_t is a block that allows inspection of a list of interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_interfaces_block_t
type Nw_parameters_iterate_interfaces_block_t = func(objectivec.Object) bool

// Nw_parameters_t is an object that stores the protocols to use for connections, options for sending data, and network path constraints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_t
type Nw_parameters_t = objectivec.Object

// Nw_path_enumerate_gateways_block_t is a block that enumerates the gateways configured on the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways_block_t
type Nw_path_enumerate_gateways_block_t = func(objectivec.Object) bool

// Nw_path_enumerate_interfaces_block_t is a block that enumerates the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces_block_t
type Nw_path_enumerate_interfaces_block_t = func(objectivec.Object) bool

// Nw_path_monitor_cancel_handler_t is a handler that indicates when a monitor has been cancelled.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel_handler_t
type Nw_path_monitor_cancel_handler_t = func()

// Nw_path_monitor_t is an observer that you use to monitor and react to network changes.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_t
type Nw_path_monitor_t = objectivec.Object

// Nw_path_monitor_update_handler_t is a handler that delivers network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_update_handler_t
type Nw_path_monitor_update_handler_t = func(objectivec.Object)

// Nw_path_t is an object that contains information about the properties of the network that a connection uses, or that are available to your app.
//
// See: https://developer.apple.com/documentation/Network/nw_path_t
type Nw_path_t = objectivec.Object

// Nw_privacy_context_t is an object that defines the privacy requirements for a set of connections.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_t
type Nw_privacy_context_t = objectivec.Object

// Nw_protocol_definition_t is the abstract superclass for identifying a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_definition_t
type Nw_protocol_definition_t = objectivec.Object

// Nw_protocol_metadata_t is the abstract superclass for specifying metadata about a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_t
type Nw_protocol_metadata_t = objectivec.Object

// Nw_protocol_options_t is the abstract superclass for configuring the options of a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_t
type Nw_protocol_options_t = objectivec.Object

// Nw_protocol_stack_iterate_protocols_block_t is a block that allows you to inspect or modify a single protocol’s options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_protocols_block_t
type Nw_protocol_stack_iterate_protocols_block_t = func(objectivec.Object)

// Nw_protocol_stack_t is an ordered set of protocol options that define the protocols that connections and listeners use.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_t
type Nw_protocol_stack_t = objectivec.Object

// Nw_proxy_config_t is a proxy configuration for Relays, Oblivious HTTP, HTTP CONNECT, or SOCKSv5.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_t
type Nw_proxy_config_t = objectivec.Object

// See: https://developer.apple.com/documentation/Network/nw_proxy_domain_enumerator_t
type Nw_proxy_domain_enumerator_t = func(string)

// Nw_relay_hop_t is a single relay server you can chain together with other servers.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_t
type Nw_relay_hop_t = objectivec.Object

// Nw_report_protocol_enumerator_t is a block used to enumerate protocol handshakes performed during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_report_protocol_enumerator_t
type Nw_report_protocol_enumerator_t = func(objectivec.Object, uint64, uint64) bool

// Nw_report_resolution_enumerator_t is a block used to enumerate resolution steps performed during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_report_resolution_enumerator_t
type Nw_report_resolution_enumerator_t = func(objectivec.IObject, uint64, uint32, objectivec.Object, objectivec.Object) bool

// Nw_report_resolution_report_enumerator_t is iterates a list of resolution steps, as [nw_resolution_report_t] objects, performed during connection establishment, in order from first resolved to last resolved.
//
// See: https://developer.apple.com/documentation/Network/nw_report_resolution_report_enumerator_t
type Nw_report_resolution_report_enumerator_t = func(objectivec.Object) bool

// Nw_resolution_report_t is a description of a single DNS resolution step.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_t
type Nw_resolution_report_t = objectivec.Object

// Nw_resolver_config_t is a DNS server configuration that uses TLS or HTTPS.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_t
type Nw_resolver_config_t = objectivec.Object

// Nw_txt_record_access_bytes_t is a block that provides access to the raw bytes of a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes_t
type Nw_txt_record_access_bytes_t = func(*uint8, uint32) bool

// Nw_txt_record_access_key_t is a block that returns a value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_key_t
type Nw_txt_record_access_key_t = func(string, objectivec.IObject, *uint8, uint32) bool

// Nw_txt_record_applier_t is a block that iterates over values and keys in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_applier_t
type Nw_txt_record_applier_t = func(string, objectivec.IObject, *uint8, uint32) bool

// Nw_txt_record_t is a dictionary representing a TXT record in a DNS packet.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_t
type Nw_txt_record_t = objectivec.Object

// Nw_ws_additional_header_enumerator_t is a block that enumerates additional HTTP headers in a WebSocket client request.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_additional_header_enumerator_t
type Nw_ws_additional_header_enumerator_t = func(string, string) bool

// Nw_ws_client_request_handler_t is a handler that delivers inbound client handshake requests.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_client_request_handler_t
type Nw_ws_client_request_handler_t = func(objectivec.Object) objectivec.Object

// Nw_ws_pong_handler_t is a handler that indicates that a Pong message has been received for a previously sent Ping message, or that an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_pong_handler_t
type Nw_ws_pong_handler_t = func(objectivec.Object)

// Nw_ws_request_t is a WebSocket handshake request sent from a client to a server.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_t
type Nw_ws_request_t = objectivec.Object

// Nw_ws_response_t is a WebSocket handshake reponse sent from a server to a client.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_t
type Nw_ws_response_t = objectivec.Object

// Nw_ws_subprotocol_enumerator_t is a block that enumerates the supported subprotocols in a WebSocket client request.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_subprotocol_enumerator_t
type Nw_ws_subprotocol_enumerator_t = func(string) bool
