// Code generated from Apple documentation for Network. DO NOT EDIT.

// Package network provides Go bindings for the Network framework.
//
// Create network connections to send and receive data using transport and
// security protocols.
//
// Use this framework when you need direct access to protocols like TLS, TCP,
// and UDP for your custom application protocols. Continue to use
// [NSURLSession], which is built upon this framework, for loading HTTP- and
// URL-based resources. For in-depth advice on where to start with networking,
// see TN3151: Choosing the right networking API.
//
// # Essentials
//
//   - [NWEndpoint]: A local or remote endpoint in a network connection. ([NWEndpointType])
//   - [NWParameters]: An object that stores the protocols to use for connections, options for sending data, and network path constraints. ([NWParametersConfigureProtocolBlock], [NWProtocolStack], [NWProtocolDefinition], [NWProtocolOptions], [NWParametersIterateInterfaceTypesBlock])
//
// # Connections and Listeners
//
//   - [Implementing netcat with Network Framework]: Build a simple tool that establishes network connections and transfers data.
//   - [NWConnection]: A bidirectional data connection between a local endpoint and a remote endpoint. ([NWConnectionState], [NWConnectionStateChangedHandler], [NWConnectionSendCompletion], [NWContentContext], [NWConnectionReceiveCompletion])
//   - [NWListener]: An object you use to listen for incoming network connections. ([NWListenerNewConnectionHandler], [NWAdvertiseDescriptor], [NWListenerAdvertisedEndpointChangedHandler], [NWListenerStateChangedHandler], [NWListenerState])
//   - [NWBrowser]: An object you use to browse for available network services. ([NWBrowseDescriptor], [NWBrowserBrowseResultsChangedHandler], [NWBrowseResult], [NWBrowserStateChangedHandler], [NWBrowserState])
//   - [NWConnectionGroup]: An object you use to communicate with a group of endpoints, such as an IP multicast group on a local network. ([NWGroupDescriptor], [NWGroupDescriptorEnumerateEndpointsBlock], [NWConnectionGroupReceiveHandler], [NWConnectionGroupSendCompletion], [NWConnectionGroupStateChangedHandler])
//   - [NWEthernetChannel]: An object you use to send and receive custom Ethernet frames. ([NWEthernetChannelStateChangedHandler], [NWEthernetChannelState], [NWEthernetChannelSendCompletion], [NWEthernetChannelReceiveHandler], [NWEthernetAddress])
//
// # Network Protocols
//
//   - [TCP Options]: Configure options for connections that use the Transmission Control Protocol.
//   - [TLS Options]: Configure options for connections that use Transport Layer Security.
//   - [QUIC Options]: Configure options for connections that use the QUIC transport protocol.
//   - [UDP Options]: Configure options for connections that use the User Datagram Protocol.
//   - [IP Options]: Configure Internet Protocol options on connections. ([NWIPEcnFlag], [NWIPVersion], [NWIPLocalAddressPreference])
//   - [WebSocket Options]: Configure options for connections that use WebSocket. ([NWWsVersion], [NWWsOpcode], [NWWsCloseCode], [NWWsPongHandler], [NWWsClientRequestHandler])
//   - [Framer Protocol Options]: Create custom protocols to frame applications messages over a connection. ([NWFramerStartHandler], [NWFramer], [NWFramerStartResult], [NWFramerOutputHandler], [NWFramerParseCompletion])
//
// # Network Security and Privacy
//
//   - [Security Options]: Configure security options for TLS handshakes.
//   - [Privacy Management]: Configure parameters related to user privacy. ([NWParametersAttribution])
//   - [Creating an Identity for Local Network TLS]: Learn how to create and use a digital identity in your application for local network TLS.
//
// # Paths and Interfaces
//
//   - [NWPath]: An object that contains information about the properties of the network that a connection uses, or that are available to your app. ([NWPathStatus], [NWPathEnumerateInterfacesBlock], [NWPathEnumerateGatewaysBlock])
//   - [NWPathMonitor]: An observer that you use to monitor and react to network changes. ([NWPathMonitorUpdateHandler], [NWPathMonitorCancelHandler])
//   - [NWInterface]: An interface that a network connection uses to send and receive data. ([NWInterfaceType])
//
// # Memory Management
//
//   - [NWRelease]: Releases a reference count on a Network.framework object.
//   - [NWRetain]: Adds a reference count to a Network.framework object.
//   - [NWObject]: The generic type for objects in the Network framework.
//
// # Errors
//
//   - [NWError]: The errors returned by the Network framework. ([NWErrorDomain])
//
// # Network Debugging
//
//   - [Choosing a Network Debugging Tool]: Decide which tool works best for your network debugging problem.
//   - [Debugging HTTP Server-Side Errors]: Understand HTTP server-side errors and how to debug them.
//   - [Debugging HTTPS Problems with CFNetwork Diagnostic Logging]: Use CFNetwork diagnostic logging to investigate HTTP and HTTPS problems.
//   - [Recording a Packet Trace]: Learn how to record a low-level trace of network traffic.
//   - [Taking Advantage of Third-Party Network Debugging Tools]: Learn about the available third-party network debugging tools.
//   - [Testing and Debugging L4S in Your App]: Learn how to verify your app on an L4S-capable host and network to improve your app’s responsiveness.
//
// # Protocols
//
//   - [OSNwAdvertiseDescriptor]
//   - [OSNwBrowseDescriptor]
//   - [OSNwBrowseResult]
//   - [OSNwBrowser]
//   - [OSNwConnection]
//   - [OSNwConnectionGroup]
//   - [OSNwContentContext]
//   - [OSNwDataTransferReport]
//   - [OSNwEndpoint]
//   - [OSNwError]
//   - [OSNwEstablishmentReport]
//   - [OSNwEthernetChannel]
//   - [OSNwFramer]
//   - [OSNwGroupDescriptor]
//   - [OSNwInterface]
//   - [OSNwListener]
//   - [OSNwObject]
//   - [OSNwParameters]
//   - [OSNwPath]
//   - [OSNwPathMonitor]
//   - [OSNwPrivacyContext]
//   - [OSNwProtocolDefinition]
//   - [OSNwProtocolMetadata]
//   - [OSNwProtocolOptions]
//   - [OSNwProtocolStack]
//   - [OSNwProxyConfig]
//   - [OSNwRelayHop]
//   - [OSNwResolutionReport]
//   - [OSNwResolverConfig]
//   - [OSNwTXTRecord]
//   - [OSNwWsRequest]
//   - [OSNwWsResponse]
//
// # Variables
//
//   - [KNWErrorDomainWiFiAware]
//
// # Functions
//
//   - [NWParametersGetAllowUltraConstrained]
//   - [NWParametersSetAllowUltraConstrained]
//   - [NWPathGetLinkQuality]
//   - [NWPathIsUltraConstrained]
//
// # Macros
//
//   - NW_ASSUME_EXTERNALLY_RETAINED_BEGIN
//   - NW_ASSUME_EXTERNALLY_RETAINED_END
//   - NW_EXTERNALLY_RETAINED
//
// # Enumerations
//
//   - [NWLinkQuality]//
//
// [Choosing a Network Debugging Tool]: https://developer.apple.com/documentation/network/choosing-a-network-debugging-tool
// [Creating an Identity for Local Network TLS]: https://developer.apple.com/documentation/network/creating-an-identity-for-local-network-tls
// [Debugging HTTP Server-Side Errors]: https://developer.apple.com/documentation/network/debugging-http-server-side-errors
// [Debugging HTTPS Problems with CFNetwork Diagnostic Logging]: https://developer.apple.com/documentation/network/debugging-https-problems-with-cfnetwork-diagnostic-logging
// [Framer Protocol Options]: https://developer.apple.com/documentation/network/framer-protocol-options
// [IP Options]: https://developer.apple.com/documentation/network/ip-options
// [Implementing netcat with Network Framework]: https://developer.apple.com/documentation/network/implementing-netcat-with-network-framework
// [Privacy Management]: https://developer.apple.com/documentation/network/privacy-management
// [QUIC Options]: https://developer.apple.com/documentation/network/quic-options
// [Recording a Packet Trace]: https://developer.apple.com/documentation/network/recording-a-packet-trace
// [Security Options]: https://developer.apple.com/documentation/network/security-options
// [TCP Options]: https://developer.apple.com/documentation/network/tcp-options
// [TLS Options]: https://developer.apple.com/documentation/network/tls-options
// [Taking Advantage of Third-Party Network Debugging Tools]: https://developer.apple.com/documentation/network/taking-advantage-of-third-party-network-debugging-tools
// [Testing and Debugging L4S in Your App]: https://developer.apple.com/documentation/network/testing-and-debugging-l4s-in-your-app
// [UDP Options]: https://developer.apple.com/documentation/network/udp-options
// [WebSocket Options]: https://developer.apple.com/documentation/network/websocket-options
package network

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Network library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Network.framework/Network",
	"/usr/lib/libNetwork.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: Network: failed to load framework from any known path\n")
}
