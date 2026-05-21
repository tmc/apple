// Code generated from Apple documentation. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NWAdvertiseDescriptor is a description used to advertise the Bonjour service that a listener provides.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_t
type NWAdvertiseDescriptor = objectivec.Object

// NWAdvertiseDescriptorFromID constructs a [NWAdvertiseDescriptor] from an objc.ID.
func NWAdvertiseDescriptorFromID(id objc.ID) NWAdvertiseDescriptor {
	return NWAdvertiseDescriptor{ID: id}
}

// NWBrowseDescriptor is a service description used to discover Bonjour services.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_t
type NWBrowseDescriptor = objectivec.Object

// NWBrowseDescriptorFromID constructs a [NWBrowseDescriptor] from an objc.ID.
func NWBrowseDescriptorFromID(id objc.ID) NWBrowseDescriptor {
	return NWBrowseDescriptor{ID: id}
}

// NWBrowseResultEnumerateInterface is a handler that enumerates the interfaces associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interface_t
type NWBrowseResultEnumerateInterface = func(objectivec.Object) bool

// NWBrowseResult is a discovered service and metadata about the service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_t
type NWBrowseResult = objectivec.Object

// NWBrowseResultFromID constructs a [NWBrowseResult] from an objc.ID.
func NWBrowseResultFromID(id objc.ID) NWBrowseResult {
	return NWBrowseResult{ID: id}
}

// NWBrowserBrowseResultsChangedHandler is a handler that delivers updates about discovered services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_browse_results_changed_handler_t
type NWBrowserBrowseResultsChangedHandler = func(objectivec.Object, objectivec.Object, bool)

// NWBrowserStateChangedHandler is a handler that delivers browser state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_state_changed_handler_t
type NWBrowserStateChangedHandler = func(NWBrowserState, NWError)

// NWBrowser is an object you use to browse for available network services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_t
type NWBrowser = objectivec.Object

// NWBrowserFromID constructs a [NWBrowser] from an objc.ID.
func NWBrowserFromID(id objc.ID) NWBrowser {
	return NWBrowser{ID: id}
}

// NWConnectionBooleanEventHandler is a handler that receives Boolean state updates from a connection, such as viability and better path state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_boolean_event_handler_t
type NWConnectionBooleanEventHandler = func(bool)

// See: https://developer.apple.com/documentation/Network/nw_connection_group_new_connection_handler_t
type NWConnectionGroupNewConnectionHandler = func(objectivec.Object)

// NWConnectionGroupReceiveHandler is a handler that receives inbound messages from members of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_receive_handler_t
type NWConnectionGroupReceiveHandler = func(objectivec.Object, objectivec.Object, bool)

// NWConnectionGroupSendCompletion is a completion to notify you when data has been processed and sent.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_send_completion_t
type NWConnectionGroupSendCompletion = func(NWError)

// NWConnectionGroupStateChangedHandler is a handler that receives connection group state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_state_changed_handler_t
type NWConnectionGroupStateChangedHandler = func(NWConnectionGroupState, NWError)

// NWConnectionGroup is an object you use to communicate with a group of endpoints, such as an IP multicast group on a local network.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_t
type NWConnectionGroup = objectivec.Object

// NWConnectionGroupFromID constructs a [NWConnectionGroup] from an objc.ID.
func NWConnectionGroupFromID(id objc.ID) NWConnectionGroup {
	return NWConnectionGroup{ID: id}
}

// NWConnectionPathEventHandler is a handler that delivers network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_path_event_handler_t
type NWConnectionPathEventHandler = func(objectivec.Object)

// NWConnectionReceiveCompletion is a completion handler that indicates when content has been received by the connection, or that an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive_completion_t
type NWConnectionReceiveCompletion = func(objectivec.Object, objectivec.Object, bool, NWError)

// NWConnectionSendCompletion is a completion handler that indicates when the connection has finished processing sent content.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_send_completion_t
type NWConnectionSendCompletion = func(NWError)

// NWConnectionStateChangedHandler is a handler that delivers connection state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_state_changed_handler_t
type NWConnectionStateChangedHandler = func(NWConnectionState, NWError)

// NWConnection is a bidirectional data connection between a local endpoint and a remote endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_t
type NWConnection = objectivec.Object

// NWConnectionFromID constructs a [NWConnection] from an objc.ID.
func NWConnectionFromID(id objc.ID) NWConnection {
	return NWConnection{ID: id}
}

// NWContentContext is a representation of a message to send or receive, containing protocol metadata and send properties.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_t
type NWContentContext = objectivec.Object

// NWContentContextFromID constructs a [NWContentContext] from an objc.ID.
func NWContentContextFromID(id objc.ID) NWContentContext {
	return NWContentContext{ID: id}
}

// NWDataTransferReportCollectBlock is a block that is delivered when a data transfer report is fully collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect_block_t
type NWDataTransferReportCollectBlock = func(objectivec.Object)

// NWDataTransferReport is a report that provides metrics about data being sent and received on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_t
type NWDataTransferReport = objectivec.Object

// NWDataTransferReportFromID constructs a [NWDataTransferReport] from an objc.ID.
func NWDataTransferReportFromID(id objc.ID) NWDataTransferReport {
	return NWDataTransferReport{ID: id}
}

// NWEndpoint is a local or remote endpoint in a network connection.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_t
type NWEndpoint = objectivec.Object

// NWEndpointFromID constructs a [NWEndpoint] from an objc.ID.
func NWEndpointFromID(id objc.ID) NWEndpoint {
	return NWEndpoint{ID: id}
}

// NWError is the errors returned by the Network framework.
//
// See: https://developer.apple.com/documentation/Network/nw_error_t

type NWError struct {
	objectivec.Object
}

// IsZero reports whether the network error is nil.
func (o NWError) IsZero() bool {
	return o.ID == 0
}

// Error returns the Objective-C description for the network error.
func (o NWError) Error() string {
	if o.ID == 0 {
		return ""
	}
	return o.Description()
}

// Code returns the numeric error code from Network.framework.
func (o NWError) Code() int {
	if o.ID == 0 {
		return 0
	}
	return NWErrorGetErrorCode(o)
}

// Domain returns the structured Network.framework error domain.
func (o NWError) Domain() NWErrorDomain {
	if o.ID == 0 {
		return *new(NWErrorDomain)
	}
	return NWErrorGetErrorDomain(o)
}

// DomainString returns the string form of the Network.framework error domain.
func (o NWError) DomainString() string {
	if o.ID == 0 {
		return ""
	}
	return NWErrorGetErrorDomain(o).String()
}

// CopyCFError returns a retained Core Foundation copy of the network error.
func (o NWError) CopyCFError() corefoundation.CFErrorRef {
	if o.ID == 0 {
		return *new(corefoundation.CFErrorRef)
	}
	return NWErrorCopyCfError(o)
}

// NWEstablishmentReportAccessBlock is a block that delivers a connection’s establishment report when it’s in the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_access_block_t
type NWEstablishmentReportAccessBlock = func(objectivec.Object)

// NWEstablishmentReport is a report that provides metrics about how a connection was established.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_t
type NWEstablishmentReport = objectivec.Object

// NWEstablishmentReportFromID constructs a [NWEstablishmentReport] from an objc.ID.
func NWEstablishmentReportFromID(id objc.ID) NWEstablishmentReport {
	return NWEstablishmentReport{ID: id}
}

// NWEthernetAddress is a 48-bit Ethernet address.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_address_t
type NWEthernetAddress = kernel.Pointer

// NWEthernetChannelReceiveHandler is a handler that delivers inbound Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_receive_handler_t
type NWEthernetChannelReceiveHandler = func(objectivec.Object, uint32, *uint8, *uint8)

// NWEthernetChannelSendCompletion is a handler that indicates when an Ethernet frame has been sent, or if an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send_completion_t
type NWEthernetChannelSendCompletion = func(NWError)

// NWEthernetChannelStateChangedHandler is a handler that delivers Ethernet channel state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_state_changed_handler_t
type NWEthernetChannelStateChangedHandler = func(NWEthernetChannelState, NWError)

// NWEthernetChannel is an object you use to send and receive custom Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_t
type NWEthernetChannel = objectivec.Object

// NWEthernetChannelFromID constructs a [NWEthernetChannel] from an objc.ID.
func NWEthernetChannelFromID(id objc.ID) NWEthernetChannel {
	return NWEthernetChannel{ID: id}
}

// NWFramerBlock is a block to be invoked asynchronously on your framer protocol’s scheduling context.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_block_t
type NWFramerBlock = func()

// NWFramerCleanupHandler is a handler that tells your protocol to clean up all allocations before being deallocated.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_cleanup_handler_t
type NWFramerCleanupHandler = func(objectivec.Object)

// NWFramerInputHandler is a handler that notifies your protocol that new inbound data is available to parse.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_input_handler_t
type NWFramerInputHandler = func(objectivec.Object) uint64

// NWFramerMessageDisposeValue is a handler that’s invoked when your custom value needs to be released due to a message being released or the value being replaced.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_dispose_value_t
type NWFramerMessageDisposeValue = func(kernel.Pointer)

// NWFramerMessage is a message for a custom protocol, in which you can store arbitrary key-value pairs.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_t
type NWFramerMessage = NWProtocolMetadata

// NWFramerOutputHandler is a handler that notifies your protocol about a new outbound message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_output_handler_t
type NWFramerOutputHandler = func(objectivec.Object, objectivec.Object, uint32, bool)

// NWFramerParseCompletion is a handler that examines a range of data being sent or received.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_completion_t
type NWFramerParseCompletion = func(*uint8, uint32, bool) uint64

// NWFramerStartHandler is a handler that represents the entry point into your custom protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_start_handler_t
type NWFramerStartHandler = func(objectivec.Object) NWFramerStartResult

// NWFramerStopHandler is a handler that requests that your protocol send any final messages to close the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_stop_handler_t
type NWFramerStopHandler = func(objectivec.Object) bool

// NWFramer is an object that represents a single instance of your custom protocol running in a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_t
type NWFramer = objectivec.Object

// NWFramerFromID constructs a [NWFramer] from an objc.ID.
func NWFramerFromID(id objc.ID) NWFramer {
	return NWFramer{ID: id}
}

// NWFramerWakeupHandler is a handler that delivers a scheduled wakeup event.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_wakeup_handler_t
type NWFramerWakeupHandler = func(objectivec.Object)

// NWGroupDescriptorEnumerateEndpointsBlock is a handler that lists all endpoints added to the group descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints_block_t
type NWGroupDescriptorEnumerateEndpointsBlock = func(objectivec.Object) bool

// NWGroupDescriptor is a type that defines a group of endpoints with which you can communicate, such as a multicast group.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_t
type NWGroupDescriptor = objectivec.Object

// NWGroupDescriptorFromID constructs a [NWGroupDescriptor] from an objc.ID.
func NWGroupDescriptorFromID(id objc.ID) NWGroupDescriptor {
	return NWGroupDescriptor{ID: id}
}

// NWInterface is an interface that a network connection uses to send and receive data.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_t
type NWInterface = objectivec.Object

// NWInterfaceFromID constructs a [NWInterface] from an objc.ID.
func NWInterfaceFromID(id objc.ID) NWInterface {
	return NWInterface{ID: id}
}

// NWListenerAdvertisedEndpointChangedHandler is a handler that indicates changes to the service endpoints being advertised as they are added and removed.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_advertised_endpoint_changed_handler_t
type NWListenerAdvertisedEndpointChangedHandler = func(objectivec.Object, bool)

// See: https://developer.apple.com/documentation/Network/nw_listener_new_connection_group_handler_t
type NWListenerNewConnectionGroupHandler = func(objectivec.Object)

// NWListenerNewConnectionHandler is a handler that delivers inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_new_connection_handler_t
type NWListenerNewConnectionHandler = func(objectivec.Object)

// NWListenerStateChangedHandler is a handler that delivers listener state updates with associated errors.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_state_changed_handler_t
type NWListenerStateChangedHandler = func(NWListenerState, NWError)

// NWListener is an object you use to listen for incoming network connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_t
type NWListener = objectivec.Object

// NWListenerFromID constructs a [NWListener] from an objc.ID.
func NWListenerFromID(id objc.ID) NWListener {
	return NWListener{ID: id}
}

// NWObject is the generic type for objects in the Network framework.
//
// See: https://developer.apple.com/documentation/Network/nw_object_t
type NWObject = objectivec.Object

// NWObjectFromID constructs a [NWObject] from an objc.ID.
func NWObjectFromID(id objc.ID) NWObject {
	return NWObject{ID: id}
}

// NWParametersConfigureProtocolBlock is a block to configure protocol options during the creation of a parameters object.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_configure_protocol_block_t
type NWParametersConfigureProtocolBlock = func(objectivec.Object)

// NWParametersIterateInterfaceTypesBlock is a block that allows inspection of a list of interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_interface_types_block_t
type NWParametersIterateInterfaceTypesBlock = func(NWInterfaceType) bool

// NWParametersIterateInterfacesBlock is a block that allows inspection of a list of interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_interfaces_block_t
type NWParametersIterateInterfacesBlock = func(objectivec.Object) bool

// NWParameters is an object that stores the protocols to use for connections, options for sending data, and network path constraints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_t
type NWParameters = objectivec.Object

// NWParametersFromID constructs a [NWParameters] from an objc.ID.
func NWParametersFromID(id objc.ID) NWParameters {
	return NWParameters{ID: id}
}

// NWPathEnumerateGatewaysBlock is a block that enumerates the gateways configured on the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways_block_t
type NWPathEnumerateGatewaysBlock = func(objectivec.Object) bool

// NWPathEnumerateInterfacesBlock is a block that enumerates the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces_block_t
type NWPathEnumerateInterfacesBlock = func(objectivec.Object) bool

// NWPathMonitorCancelHandler is a handler that indicates when a monitor has been cancelled.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel_handler_t
type NWPathMonitorCancelHandler = func()

// NWPathMonitor is an observer that you use to monitor and react to network changes.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_t
type NWPathMonitor = objectivec.Object

// NWPathMonitorFromID constructs a [NWPathMonitor] from an objc.ID.
func NWPathMonitorFromID(id objc.ID) NWPathMonitor {
	return NWPathMonitor{ID: id}
}

// NWPathMonitorUpdateHandler is a handler that delivers network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_update_handler_t
type NWPathMonitorUpdateHandler = func(objectivec.Object)

// NWPath is an object that contains information about the properties of the network that a connection uses, or that are available to your app.
//
// See: https://developer.apple.com/documentation/Network/nw_path_t
type NWPath = objectivec.Object

// NWPathFromID constructs a [NWPath] from an objc.ID.
func NWPathFromID(id objc.ID) NWPath {
	return NWPath{ID: id}
}

// NWPrivacyContext is an object that defines the privacy requirements for a set of connections.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_t
type NWPrivacyContext = objectivec.Object

// NWPrivacyContextFromID constructs a [NWPrivacyContext] from an objc.ID.
func NWPrivacyContextFromID(id objc.ID) NWPrivacyContext {
	return NWPrivacyContext{ID: id}
}

// NWProtocolDefinition is the abstract superclass for identifying a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_definition_t
type NWProtocolDefinition = objectivec.Object

// NWProtocolDefinitionFromID constructs a [NWProtocolDefinition] from an objc.ID.
func NWProtocolDefinitionFromID(id objc.ID) NWProtocolDefinition {
	return NWProtocolDefinition{ID: id}
}

// NWProtocolMetadata is the abstract superclass for specifying metadata about a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_t
type NWProtocolMetadata = objectivec.Object

// NWProtocolMetadataFromID constructs a [NWProtocolMetadata] from an objc.ID.
func NWProtocolMetadataFromID(id objc.ID) NWProtocolMetadata {
	return NWProtocolMetadata{ID: id}
}

// NWProtocolOptions is the abstract superclass for configuring the options of a network protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_t
type NWProtocolOptions = objectivec.Object

// NWProtocolOptionsFromID constructs a [NWProtocolOptions] from an objc.ID.
func NWProtocolOptionsFromID(id objc.ID) NWProtocolOptions {
	return NWProtocolOptions{ID: id}
}

// NWProtocolStackIterateProtocolsBlock is a block that allows you to inspect or modify a single protocol’s options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_protocols_block_t
type NWProtocolStackIterateProtocolsBlock = func(objectivec.Object)

// NWProtocolStack is an ordered set of protocol options that define the protocols that connections and listeners use.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_t
type NWProtocolStack = objectivec.Object

// NWProtocolStackFromID constructs a [NWProtocolStack] from an objc.ID.
func NWProtocolStackFromID(id objc.ID) NWProtocolStack {
	return NWProtocolStack{ID: id}
}

// NWProxyConfig is a proxy configuration for Relays, Oblivious HTTP, HTTP CONNECT, or SOCKSv5.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_t
type NWProxyConfig = objectivec.Object

// NWProxyConfigFromID constructs a [NWProxyConfig] from an objc.ID.
func NWProxyConfigFromID(id objc.ID) NWProxyConfig {
	return NWProxyConfig{ID: id}
}

// See: https://developer.apple.com/documentation/Network/nw_proxy_domain_enumerator_t
type NWProxyDomainEnumerator = func(string)

// NWRelayHop is a single relay server you can chain together with other servers.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_t
type NWRelayHop = objectivec.Object

// NWRelayHopFromID constructs a [NWRelayHop] from an objc.ID.
func NWRelayHopFromID(id objc.ID) NWRelayHop {
	return NWRelayHop{ID: id}
}

// NWReportProtocolEnumerator is a block used to enumerate protocol handshakes performed during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_report_protocol_enumerator_t
type NWReportProtocolEnumerator = func(objectivec.Object, uint64, uint64) bool

// NWReportResolutionEnumerator is a block used to enumerate resolution steps performed during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_report_resolution_enumerator_t
type NWReportResolutionEnumerator = func(NWReportResolutionSource, uint64, uint32, objectivec.Object, objectivec.Object) bool

// NWReportResolutionReportEnumerator is iterates a list of resolution steps, as [nw_resolution_report_t] objects, performed during connection establishment, in order from first resolved to last resolved.
//
// See: https://developer.apple.com/documentation/Network/nw_report_resolution_report_enumerator_t
type NWReportResolutionReportEnumerator = func(objectivec.Object) bool

// NWResolutionReport is a description of a single DNS resolution step.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_t
type NWResolutionReport = objectivec.Object

// NWResolutionReportFromID constructs a [NWResolutionReport] from an objc.ID.
func NWResolutionReportFromID(id objc.ID) NWResolutionReport {
	return NWResolutionReport{ID: id}
}

// NWResolverConfig is a DNS server configuration that uses TLS or HTTPS.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_t
type NWResolverConfig = objectivec.Object

// NWResolverConfigFromID constructs a [NWResolverConfig] from an objc.ID.
func NWResolverConfigFromID(id objc.ID) NWResolverConfig {
	return NWResolverConfig{ID: id}
}

// NWTXTRecordAccessBytes is a block that provides access to the raw bytes of a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes_t
type NWTXTRecordAccessBytes = func(*uint8, uint32) bool

// NWTXTRecordAccessKey is a block that returns a value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_key_t
type NWTXTRecordAccessKey = func(string, NWTXTRecordFindKey, *uint8, uint32) bool

// NWTXTRecordApplier is a block that iterates over values and keys in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_applier_t
type NWTXTRecordApplier = func(string, NWTXTRecordFindKey, *uint8, uint32) bool

// NWTXTRecord is a dictionary representing a TXT record in a DNS packet.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_t
type NWTXTRecord = objectivec.Object

// NWTXTRecordFromID constructs a [NWTXTRecord] from an objc.ID.
func NWTXTRecordFromID(id objc.ID) NWTXTRecord {
	return NWTXTRecord{ID: id}
}

// NWWsAdditionalHeaderEnumerator is a block that enumerates additional HTTP headers in a WebSocket client request.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_additional_header_enumerator_t
type NWWsAdditionalHeaderEnumerator = func(string, string) bool

// NWWsClientRequestHandler is a handler that delivers inbound client handshake requests.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_client_request_handler_t
type NWWsClientRequestHandler = func(objectivec.Object) objectivec.Object

// NWWsPongHandler is a handler that indicates that a Pong message has been received for a previously sent Ping message, or that an error was encountered.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_pong_handler_t
type NWWsPongHandler = func(NWError)

// NWWsRequest is a WebSocket handshake request sent from a client to a server.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_t
type NWWsRequest = objectivec.Object

// NWWsRequestFromID constructs a [NWWsRequest] from an objc.ID.
func NWWsRequestFromID(id objc.ID) NWWsRequest {
	return NWWsRequest{ID: id}
}

// NWWsResponse is a WebSocket handshake reponse sent from a server to a client.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_t
type NWWsResponse = objectivec.Object

// NWWsResponseFromID constructs a [NWWsResponse] from an objc.ID.
func NWWsResponseFromID(id objc.ID) NWWsResponse {
	return NWWsResponse{ID: id}
}

// NWWsSubprotocolEnumerator is a block that enumerates the supported subprotocols in a WebSocket client request.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_subprotocol_enumerator_t
type NWWsSubprotocolEnumerator = func(string) bool
