// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

// Package networkextension provides Go bindings for the NetworkExtension framework.
//
// Customize and extend core networking features.
//
// With the NetworkExtension framework, you can customize and extend the
// system’s core networking features. Specifically, you can:
//
// # Wi-Fi management
//
//   - [Wi-Fi configuration]: Add persistent Wi-Fi configurations, or temporarily move the device to a specific Wi-Fi network.
//   - [Configuring a Wi-Fi accessory to join a network]: Associate an iOS device with an accessory’s network to deliver network configuration information.
//   - [Hotspot helper]: Integrate your app with the iOS hotspot network subsystem.
//
// # Virtual private networks
//
//   - [Routing your VPN network traffic]: Configure your VPN to include and exclude some network traffic.
//   - [Personal VPN]: Create and manage a VPN configuration that uses one of the built-in VPN protocols (IPsec or IKEv2). ([NEVPNManager], [NEVPNProtocolIKEv2], [NEVPNProtocolIPSec], [NEVPNProtocol], [NEVPNConnection])
//   - [Packet tunnel provider]: Implement a VPN client for a packet-oriented, custom VPN protocol. ([NEPacketTunnelProvider], [NETunnelProvider], [NEProvider], [NEPacketTunnelNetworkSettings], [NETunnelNetworkSettings])
//   - [App proxy provider]: Implement a VPN client for a flow-oriented, custom VPN protocol. ([NEAppProxyProvider], [NETunnelProvider], [NEProvider], [NETunnelNetworkSettings], [NEAppProxyTCPFlow])
//
// # Network relays
//
//   - Relays: Create and manage a system-wide network relay configuration that uses built-in proxying for TCP and UDP traffic over HTTP/3 and HTTP/2. ([NERelayManager], [NERelay])
//
// # Content filters
//
//   - [Content filter providers]: Create an on-device network content filter. ([NEFilterDataProvider], [NEFilterPacketProvider], [NEFilterProvider], [NEFilterFlow], [NEFilterSocketFlow])
//
// # URL filters
//
//   - [URL filters]: Create a filter that analyzes full URLs, while preserving privacy. ([NEURLFilter])
//
// # DNS configurations
//
//   - [DNS settings]: Create and manage a system-wide DNS configuration that uses built-in encrypted DNS protocols. ([NEDNSSettingsManager], [NEDNSOverHTTPSSettings], [NEDNSOverTLSSettings])
//   - [DNS proxy provider]: Create an on-device DNS proxy using a custom protocol. ([NEDNSProxyProvider], [NEDNSSettings], [NEAppProxyTCPFlow], [NEAppProxyUDPFlow], [NEAppProxyFlow])
//
// # Local networking
//
//   - [Local push connectivity]: Provide functionality similar to Apple Push Notification Service when access to the wider internet is unavailable.
//
// # Variables
//
//   - [NEDNSSettingsConfigurationDidChangeNotification]
//   - [NERelayClientErrorDomain]
//   - [NERelayConfigurationDidChangeNotification]
//
// # Macros
//
//   - NEAPPPROXYFLOW_EXPORT
//   - NEAPPPUSH_EXPORT
//   - NECOPYNULLABLE
//   - NEDNSPROXY_EXPORT
//   - NEDNSSETTINGS_EXPORT
//   - NEFILTER_DATA_PROVIDER_EXTERN
//   - NEFILTER_EXPORT
//   - NEHSCFG_EXPORT
//   - NEHS_EXPORT
//   - NERELAY_EXPORT
//   - NETUNNELPROVIDER_EXPORT
//   - NEURLFILTER_EXPORT
//   - NEVPN_EXPORT
//
// # Type Aliases
//
//   - [NWEndpointArray]
//
// # Enumerations
//
//   - [NERelayManagerClientError]//
//
// # Key Types
//
//   - [NEVPNProtocolIKEv2] - Settings for an IKEv2 VPN configuration.
//   - [NERelayManager] - An object you use to create and manage a network relay configuration.
//   - [NEVPNProtocol] - Settings common to both IKEv2 and IPsec VPN configurations.
//   - [NENetworkRule] - A rule to match attributes of network traffic.
//   - [NETunnelProviderManager] - An object to create and manage the tunnel provider’s VPN configuration.
//   - [NEVPNManager] - An object to create and manage a Personal VPN configuration.
//   - [NEFilterDataProvider] - The principal class for a filter data provider extension.
//   - [NEFilterProviderConfiguration] - Configuration parameters for a content filter.
//   - [NEProxySettings] - [NEProxySettings] contains HTTP proxy settings.
//   - [NEAppProxyFlow] - An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// [App proxy provider]: https://developer.apple.com/documentation/networkextension/app-proxy-provider
// [Configuring a Wi-Fi accessory to join a network]: https://developer.apple.com/documentation/networkextension/configuring-a-wi-fi-accessory-to-join-a-network
// [Content filter providers]: https://developer.apple.com/documentation/networkextension/content-filter-providers
// [DNS proxy provider]: https://developer.apple.com/documentation/networkextension/dns-proxy-provider
// [DNS settings]: https://developer.apple.com/documentation/networkextension/dns-settings
// [Hotspot helper]: https://developer.apple.com/documentation/networkextension/hotspot-helper
// [Local push connectivity]: https://developer.apple.com/documentation/networkextension/local-push-connectivity
// [Packet tunnel provider]: https://developer.apple.com/documentation/networkextension/packet-tunnel-provider
// [Personal VPN]: https://developer.apple.com/documentation/networkextension/personal-vpn
// [Routing your VPN network traffic]: https://developer.apple.com/documentation/networkextension/routing-your-vpn-network-traffic
// [URL filters]: https://developer.apple.com/documentation/networkextension/url-filters
// [Wi-Fi configuration]: https://developer.apple.com/documentation/networkextension/wi-fi-configuration
package networkextension

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the NetworkExtension library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/NetworkExtension.framework/NetworkExtension",
	"/usr/lib/libNetworkExtension.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: NetworkExtension: failed to load framework from any known path\n")
}
