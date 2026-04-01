// Code generated from Apple documentation for Network. DO NOT EDIT.

// Package network provides Go bindings for the Network framework.
//
// Create network connections to send and receive data using transport and security protocols.
//
// Use this framework when you need direct access to protocols like TLS, TCP, and UDP for your custom application protocols. Continue to use [URLSession](<doc://com.apple.documentation/documentation/Foundation/URLSession>), which is built upon this framework, for loading HTTP- and URL-based resources. For in-depth advice on where to start with networking, see [TN3151: Choosing the right networking API](<doc://com.apple.documentation/documentation/Technotes/tn3151-choosing-the-right-networking-api>).
//
// # Essentials
//
//   - NWEndpoint: A local or remote endpoint in a network connection.
//   - NWParameters: An object that stores the protocols to use for connections, options for sending data, and network path constraints.
//
// # Connections and Listeners
//
//   - NWConnection: A bidirectional data connection between a local endpoint and a remote endpoint.
//   - NWListener: An object you use to listen for incoming network connections.
//   - NWBrowser: An object you use to browse for available network services.
//   - NWConnectionGroup: An object you use to communicate with a group of endpoints, such as an IP multicast group on a local network.
//   - NWEthernetChannel: An object you use to send and receive custom Ethernet frames.
//
// # Network Protocols
//
//   - Building a custom peer-to-peer protocol: Use networking frameworks to create a custom protocol for playing a game across iOS, iPadOS, watchOS, and tvOS devices.
//   - Connecting iPadOS and visionOS apps over the local network: Build an iPadOS companion app to control your visionOS app.
//   - NWProtocolTCP: A network protocol for connections that use the Transmission Control Protocol.
//   - NWProtocolTLS: A network protocol for connections that use Transport Layer Security.
//   - NWProtocolQUIC: A network protocol for connections that use the QUIC transport protocol.
//   - NWProtocolUDP: A network protocol for connections that use the User Datagram Protocol.
//   - NWProtocolIP: A network protocol for configuring the Internet Protocol on connections.
//   - NWProtocolWebSocket: A network protocol for connections that use WebSocket.
//   - NWProtocolFramer: A customizable network protocol for defining application message parsers.
//
// # Network Security and Privacy
//
//   - Security Options: Configure security options for TLS handshakes.
//   - Privacy Management: Configure parameters related to user privacy. ([Nw_parameters_attribution_t])
//   - Creating an Identity for Local Network TLS: Learn how to create and use a digital identity in your application for local network TLS.
//
// # Paths and Interfaces
//
//   - NWPath: An object that contains information about the properties of the network that a connection uses, or that are available to your app.
//   - NWPathMonitor: An observer that you use to monitor and react to network changes.
//   - NWInterface: An interface that a network connection uses to send and receive data.
//
// # Errors
//
//   - NWError: The errors returned by objects in the Network framework.
//
// # Network Debugging
//
//   - Choosing a Network Debugging Tool: Decide which tool works best for your network debugging problem.
//   - Debugging HTTP Server-Side Errors: Understand HTTP server-side errors and how to debug them.
//   - Debugging HTTPS Problems with CFNetwork Diagnostic Logging: Use CFNetwork diagnostic logging to investigate HTTP and HTTPS problems.
//   - Recording a Packet Trace: Learn how to record a low-level trace of network traffic.
//   - Taking Advantage of Third-Party Network Debugging Tools: Learn about the available third-party network debugging tools.
//   - Testing and Debugging L4S in Your App: Learn how to verify your app on an L4S-capable host and network to improve your app’s responsiveness.
//
// # C-Language Symbols
//
//   - C-Language Symbols ([OS_nw_advertise_descriptor], [OS_nw_browse_descriptor], [OS_nw_browse_result], [OS_nw_browser], [OS_nw_connection])
//
// # Classes
//
//   - NWMultiplexGroup
//   - NetworkBrowser: Discover advertised services and devices on the network.
//   - NetworkChannel: A base class supporting sending and recieving data through an arbitrary network channel.
//   - NetworkConnection: Connect to an endpoint on the network to send and receive data.
//   - NetworkListener: Listen for incoming network connections.
//
// # Protocols
//
//   - BrowserProvider: BrowserProviders can be used when creating NetworkBrowsers.
//   - Connectable: Describes types that can be used to make NetworkConnections.
//   - ConnectionStorage: Types that conform to ConnectionStorage can be used as additional storage within a connection.
//   - DatagramProtocol: Types that conform to DatagramProtocol send and receive messages with minimal or no metadata, usually constrained to a fixed maximum size.
//   - FramerProtocol: Framer protocols allow custom framing and serialization of messages on a connection.
//   - ListenerProvider: Extensible support for configuring advertise descriptors to define the service a listener should advertise.
//   - MessageProtocol: Types that conform to MessageProtocol send and receive messages.
//   - MultiplexProtocol: Types that conform to MultiplexProtocol are allowed to be the top protocol in a network protocol stack for multiplexing network connection objects.
//   - NWParametersProvider: Types that conform to the NWParametersProvider protocol can be used to generate an NWParameters.
//   - NetworkCoder
//   - NetworkDecoder: A type that conforms to the NetworkEncoder protocol can decode data to an Encodable object
//   - NetworkEncoder: A type that conforms to the NetworkEncoder protocol can encode a Encodable object to Data
//   - NetworkFixedWidthInteger
//   - NetworkMetadataProtocol: Types that conform to NetworkProtocolOptions can be used when configuring protocol stacks.
//   - NetworkProtocolOptions
//   - OneToOneProtocol: Types that conform to OneToOneProtocol are allowed to be the top protocol in a network protocol stack for non-multiplexed connections.
//   - StreamProtocol: Types that conform to the StreamProtocol protocol expose methods for sending and receiving byte streams.
//
// # Variables
//
//   - kNWErrorDomainWiFiAware
//   - nw_error_domain_wifi_aware
//   - nw_link_quality_good
//   - nw_link_quality_minimal
//   - nw_link_quality_moderate
//   - nw_link_quality_unknown
//
// # Functions
//
//   - nw_parameters_get_allow_ultra_constrained(_:)
//   - nw_parameters_set_allow_ultra_constrained(_:_:)
//   - nw_path_get_link_quality(_:)
//   - nw_path_is_ultra_constrained(_:)
//   - withNetworkConnection(to:using:_:)
//   - withNetworkConnection(to:using:_:)
//   - withNetworkConnection(to:using:_:)
//   - withNetworkConnection(to:using:_:)
//
// [Network Documentation]: https://developer.apple.com/documentation/Network
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
