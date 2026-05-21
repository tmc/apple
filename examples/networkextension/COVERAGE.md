# NetworkExtension example coverage

The examples under this directory are local runtime checks. They avoid
preference writes, system-extension install, provider activation, keychain
identity material, and traffic capture because those require entitlements or
host policy outside a normal `go test` run.

## Covered

- `networkextension-tcp-auth-demo`: registers an Objective-C object that
  conforms to `NWTCPConnectionAuthenticationDelegate` and invokes the
  trust/identity optional selectors directly.
- `networkextension-provider-callback-demo`: registers a local Objective-C
  object with the lifecycle selectors used by `NEPacketTunnelProvider` and
  invokes start, stop, message, sleep, and wake selectors directly with
  generated completion blocks.
- `networkextension-flow-callback-demo`: registers a local Objective-C object
  with selectors shaped like app proxy flow, tunnel provider session, VPN
  disconnect, URL filter, and connection identity callbacks. It invokes the
  generated error, data, data/error, datagram-array/error, packet-array,
  packet/protocol-array, URL verdict, and identity plus certificate-chain block
  constructors directly, and calls generated app proxy flow, packet tunnel
  flow, and tunnel provider session wrapper methods against the local object.
- `networkextension-settings-demo`: constructs packet tunnel, IPv4, IPv6, DNS,
  proxy, filter rule, filter provider configuration, IPSec, and IKEv2 settings
  objects and checks setter/getter round trips. It also covers local IPv4 and
  IPv6 automatic-addressing class factories, including the IPv6 link-local
  factory, base tunnel DNS/proxy settings, filter provider identity/password
  references, vendor
  configuration, packet-provider bundle identifiers, tunnel overhead, proxy
  PAC/static settings, DNS identity references, and VPN credential references
  using non-secret in-memory `NSData`.
- `networkextension-policy-config-demo`: constructs per-app VPN app rules,
  on-demand connection rules, tunnel provider manager/protocol policy,
  transparent proxy and ethernet tunnel settings, filter verdict factories,
  packet and endpoint value objects, default packet metadata, and relay policy
  objects. It checks only local object construction and property round trips,
  including local URL filter verdict callbacks, tunnel provider configuration
  dictionaries, relay header fields, raw public keys, identity data, and
  excluded domain/FQDN policy lists.
- `networkextension-manager-config-demo`: configures local manager objects for
  VPN and DNS settings configurations. It checks in-memory manager properties
  and on-demand disconnect/ignore rules, invokes generated manager array
  completion blocks locally, and avoids preference load/save/remove calls.
- `networkextension-value-object-demo`: constructs TLS parameters, IKEv2
  security association and PPK objects, DNS proxy provider protocol values,
  additional host/network and Network.framework endpoint network-rule
  constructors, and default path/TCP/UDP state and endpoint objects. It also
  covers package-level global strings, TLS cipher-suite sets, IKEv2
  security-association algorithm fields, DNS proxy provider configuration
  dictionaries, and generated data-array completion blocks locally. It checks
  only local object construction, block invocation, and property round trips.
- `networkextension-filter-dns-provider-demo`: registers a local Objective-C
  object with selectors shaped like DNS proxy and content-filter providers,
  then calls the generated `NEDNSProxyProvider`, `NEFilterProvider`, and
  `NEFilterDataProvider` methods against it. It checks lifecycle completion
  blocks, DNS proxy flow decisions, filter reports/configuration, filter data
  verdict returns, settings application, flow resume/update calls, and
  `NEFilterPacketProvider` packet handler/delay/allow selectors. It also
  constructs local default `NEFilterFlow`, `NEFilterSocketFlow`,
  `NEFilterReport`, and `NEFilterPacketContext` objects and checks stable
  zero-value accessors.
- `networkextension-provider-wrapper-demo`: registers local Objective-C
  objects with selectors shaped like packet tunnel and app proxy providers,
  then calls the generated `NEPacketTunnelProvider` and `NEAppProxyProvider`
  wrapper methods against them. It checks lifecycle completion blocks,
  synchronous context wrappers with immediate local completions, cancellation
  methods, packet tunnel accessors, app-proxy flow decisions, and class
  availability for the Ethernet tunnel and transparent proxy provider
  subclasses.

## Deliberately Not Covered By Local Examples

- Manager preference operations such as load, save, and remove on
  `NEVPNManager`, `NETunnelProviderManager`, `NEFilterManager`,
  `NEDNSSettingsManager`, `NEDNSProxyManager`, `NEAppProxyProviderManager`,
  `NERelayManager`, and `NETransparentProxyManager`.
- DNS proxy manager property coverage. `NEDNSProxyManager.sharedManager`
  returned nil in a normal local process during example validation, so the
  local examples do not claim DNS proxy manager round-trip coverage.
- Filter manager property coverage. `NEFilterManager.sharedManager` returned
  nil in a normal local process during example validation, so the local
  examples cover filter provider configuration objects but not the singleton
  filter manager.
- Previously generated filter report, filter flow, and filter provider
  pseudo-properties that do not respond as instance selectors at runtime,
  including `NEFilterReport.shouldReport`,
  `NEFilterFlow.NEFilterFlowBytesMax`, and
  `NEFilterProvider.NEFilterErrorDomain`. These are no longer emitted by the
  current generator; local examples may use package-level global constants such
  as `networkextension.NEFilterErrorDomain`.
- Swift-only pseudo-properties that do not respond as Objective-C instance
  selectors, including the previously emitted
  `NEFilterPacketProvider.Handler` and `NEAppProxyFlow.Interface` accessors.
  These are no longer emitted by the current generator. The packet-provider
  smoke covers the real `packetHandler`, `setPacketHandler:`,
  `delayCurrentPacket:`, and `allowPacket:` selectors instead.
- `NEVPNIKEv2PPKConfiguration.allowPostQuantumKeyExchangeFallback`, which did
  not respond as an instance selector during example validation. The examples
  still cover the corresponding `NEVPNProtocolIKEv2` property.
- Live provider runtime APIs that require system-extension packaging or app
  extension activation, including packet flows, proxy flows, DNS proxy
  provider callbacks on real DNS traffic, filter data decisions on real
  flows, active relay traffic, and active `NEVPNConnection` sessions.
- Direct construction of `NEEthernetTunnelProvider` and
  `NETransparentProxyProvider`. The classes are available at runtime, but
  `new` returned nil in a normal local process during example validation; the
  examples cover inherited packet/app-proxy wrapper shapes through local
  callback objects instead.
- Real keychain-backed identity/password material. The PPK value-object demo
  and settings demo use non-secret in-memory `NSData` to exercise generated
  property shapes; they do not load or store keychain items.
