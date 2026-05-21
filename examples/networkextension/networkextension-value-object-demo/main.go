// Command networkextension-value-object-demo exercises local
// NetworkExtension value and state objects.
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/network"
	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
	"github.com/tmc/apple/objectivec"
)

func runSmoke() error {
	if err := checkPackageGlobals(); err != nil {
		return err
	}
	if err := checkTLSParameters(); err != nil {
		return err
	}
	if err := checkIKEv2PostQuantumValues(); err != nil {
		return err
	}
	if err := checkDNSProxyProviderProtocol(); err != nil {
		return err
	}
	if err := checkNetworkRules(); err != nil {
		return err
	}
	if err := checkDefaultNetworkStateObjects(); err != nil {
		return err
	}
	if err := checkGeneratedArrayBlocks(); err != nil {
		return err
	}
	return nil
}

func checkPackageGlobals() error {
	for _, global := range []struct {
		name  string
		value string
	}{
		{"NEAppProxyErrorDomain", networkextension.NEAppProxyErrorDomain},
		{"NEDNSProxyConfigurationDidChangeNotification", networkextension.NEDNSProxyConfigurationDidChangeNotification},
		{"NEDNSProxyErrorDomain", networkextension.NEDNSProxyErrorDomain},
		{"NEDNSSettingsConfigurationDidChangeNotification", networkextension.NEDNSSettingsConfigurationDidChangeNotification},
		{"NEDNSSettingsErrorDomain", networkextension.NEDNSSettingsErrorDomain},
		{"NEFilterConfigurationDidChangeNotification", networkextension.NEFilterConfigurationDidChangeNotification},
		{"NEFilterErrorDomain", networkextension.NEFilterErrorDomain},
		{"NERelayClientErrorDomain", networkextension.NERelayClientErrorDomain},
		{"NERelayConfigurationDidChangeNotification", networkextension.NERelayConfigurationDidChangeNotification},
		{"NERelayErrorDomain", networkextension.NERelayErrorDomain},
		{"NETunnelProviderErrorDomain", networkextension.NETunnelProviderErrorDomain},
		{"NEVPNConfigurationChangeNotification", networkextension.NEVPNConfigurationChangeNotification},
		{"NEVPNConnectionErrorDomain", networkextension.NEVPNConnectionErrorDomain},
		{"NEVPNConnectionStartOptionPassword", networkextension.NEVPNConnectionStartOptionPassword},
		{"NEVPNConnectionStartOptionUsername", networkextension.NEVPNConnectionStartOptionUsername},
		{"NEVPNErrorDomain", networkextension.NEVPNErrorDomain},
		{"NEVPNStatusDidChangeNotification", networkextension.NEVPNStatusDidChangeNotification},
	} {
		if global.value == "" {
			return fmt.Errorf("%s is empty", global.name)
		}
	}
	return nil
}

func checkTLSParameters() error {
	params := networkextension.NewNWTLSParameters()
	if params.GetID() == 0 {
		return fmt.Errorf("create tls parameters")
	}
	session := foundation.NewDataWithBytesLength([]byte{0x01, 0x02, 0x03})
	cipherSuites := foundation.NewSetWithArray([]objectivec.IObject{
		foundation.NewNumberWithInt(0x1301),
	})
	params.SetTLSSessionID(session)
	params.SetSSLCipherSuites(cipherSuites)
	params.SetMinimumSSLProtocolVersion(0x0303)
	params.SetMaximumSSLProtocolVersion(0x0304)
	if got := params.TLSSessionID().Length(); got != 3 {
		return fmt.Errorf("tls session id length = %d, want 3", got)
	}
	if got := params.SSLCipherSuites().Count(); got != 1 {
		return fmt.Errorf("tls cipher suites = %d, want 1", got)
	}
	if got := params.MinimumSSLProtocolVersion(); got != 0x0303 {
		return fmt.Errorf("minimum tls version = %#x, want 0x0303", got)
	}
	if got := params.MaximumSSLProtocolVersion(); got != 0x0304 {
		return fmt.Errorf("maximum tls version = %#x, want 0x0304", got)
	}
	return nil
}

func checkIKEv2PostQuantumValues() error {
	protocol := networkextension.NewNEVPNProtocolIKEv2()
	if protocol.GetID() == 0 {
		return fmt.Errorf("create ikev2 protocol")
	}
	protocol.SetAllowPostQuantumKeyExchangeFallback(true)
	if !protocol.AllowPostQuantumKeyExchangeFallback() {
		return fmt.Errorf("ikev2 post-quantum fallback = false")
	}

	securityAssociation := networkextension.NewNEVPNIKEv2SecurityAssociationParameters()
	if securityAssociation.GetID() == 0 {
		return fmt.Errorf("create ikev2 security association")
	}
	securityAssociation.SetEncryptionAlgorithm(networkextension.NEVPNIKEv2EncryptionAlgorithmAES256GCM)
	securityAssociation.SetIntegrityAlgorithm(networkextension.NEVPNIKEv2IntegrityAlgorithmSHA384)
	securityAssociation.SetDiffieHellmanGroup(networkextension.NEVPNIKEv2DiffieHellmanGroup20)
	securityAssociation.SetLifetimeMinutes(60)
	securityAssociation.SetPostQuantumKeyExchangeMethods([]foundation.NSNumber{
		foundation.NewNumberWithInt(int(networkextension.NEVPNIKEv2PostQuantumKeyExchangeMethod36)),
	})
	if got := securityAssociation.EncryptionAlgorithm(); got != networkextension.NEVPNIKEv2EncryptionAlgorithmAES256GCM {
		return fmt.Errorf("ikev2 encryption algorithm = %v", got)
	}
	if got := securityAssociation.IntegrityAlgorithm(); got != networkextension.NEVPNIKEv2IntegrityAlgorithmSHA384 {
		return fmt.Errorf("ikev2 integrity algorithm = %v", got)
	}
	if got := securityAssociation.DiffieHellmanGroup(); got != networkextension.NEVPNIKEv2DiffieHellmanGroup20 {
		return fmt.Errorf("ikev2 Diffie-Hellman group = %v", got)
	}
	if got := securityAssociation.LifetimeMinutes(); got != 60 {
		return fmt.Errorf("ikev2 lifetime minutes = %d, want 60", got)
	}
	methods := securityAssociation.PostQuantumKeyExchangeMethods()
	if len(methods) != 1 {
		return fmt.Errorf("post-quantum methods = %d, want 1", len(methods))
	}
	if got := methods[0].IntValue(); got != int(networkextension.NEVPNIKEv2PostQuantumKeyExchangeMethod36) {
		return fmt.Errorf("post-quantum method = %d", got)
	}

	keyData := foundation.NewDataWithBytesLength([]byte("key"))
	ppk := networkextension.NewVPNIKEv2PPKConfigurationWithIdentifierKeychainReference("ppk-id", keyData)
	if ppk.GetID() == 0 {
		return fmt.Errorf("create ikev2 ppk configuration")
	}
	ppk.SetIsMandatory(true)
	if got := ppk.Identifier(); got != "ppk-id" {
		return fmt.Errorf("ppk identifier = %q", got)
	}
	if got := ppk.KeychainReference().Length(); got != 3 {
		return fmt.Errorf("ppk keychain reference length = %d, want 3", got)
	}
	if !ppk.IsMandatory() {
		return fmt.Errorf("ppk mandatory = false")
	}

	protocol.SetPpkConfiguration(ppk)
	if got := protocol.PpkConfiguration().Identifier(); got != "ppk-id" {
		return fmt.Errorf("protocol ppk identifier = %q", got)
	}
	return nil
}

func checkDNSProxyProviderProtocol() error {
	protocol := networkextension.NewNEDNSProxyProviderProtocol()
	if protocol.GetID() == 0 {
		return fmt.Errorf("create dns proxy provider protocol")
	}
	protocol.SetProviderBundleIdentifier("com.example.networkextension.dnsproxy")
	protocol.SetProviderConfiguration(foundation.GetNSDictionaryClass().Dictionary())
	protocol.SetServerAddress("dnsproxy.example")
	protocol.SetUsername("dns-user")
	if got := protocol.ProviderBundleIdentifier(); got != "com.example.networkextension.dnsproxy" {
		return fmt.Errorf("dns proxy bundle id = %q", got)
	}
	if got := protocol.ProviderConfiguration().Count(); got != 0 {
		return fmt.Errorf("dns proxy provider configuration count = %d, want 0", got)
	}
	if got := protocol.ServerAddress(); got != "dnsproxy.example" {
		return fmt.Errorf("dns proxy server = %q", got)
	}
	if got := protocol.Username(); got != "dns-user" {
		return fmt.Errorf("dns proxy username = %q", got)
	}
	return nil
}

func checkNetworkRules() error {
	remote := networkextension.NewNWHostEndpointWithHostnamePort("203.0.113.0", "0")
	local := networkextension.NewNWHostEndpointWithHostnamePort("192.0.2.0", "0")

	destination := networkextension.NewNetworkRuleWithDestinationNetworkPrefixProtocol(
		remote,
		24,
		networkextension.NENetworkRuleProtocolTCP,
	)
	if destination.GetID() == 0 {
		return fmt.Errorf("create destination network rule")
	}
	if got := destination.MatchRemotePrefix(); got != 24 {
		return fmt.Errorf("destination remote prefix = %d, want 24", got)
	}
	if got := destination.MatchProtocol(); got != networkextension.NENetworkRuleProtocolTCP {
		return fmt.Errorf("destination protocol = %v", got)
	}
	if destination.MatchRemoteEndpoint().GetID() == 0 {
		return fmt.Errorf("destination remote endpoint is nil")
	}

	bidirectional := networkextension.NewNetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection(
		remote,
		24,
		local,
		24,
		networkextension.NENetworkRuleProtocolUDP,
		networkextension.NETrafficDirectionOutbound,
	)
	if bidirectional.GetID() == 0 {
		return fmt.Errorf("create remote/local network rule")
	}
	if got := bidirectional.MatchRemotePrefix(); got != 24 {
		return fmt.Errorf("bidirectional remote prefix = %d, want 24", got)
	}
	if got := bidirectional.MatchLocalPrefix(); got != 24 {
		return fmt.Errorf("bidirectional local prefix = %d, want 24", got)
	}
	if got := bidirectional.MatchProtocol(); got != networkextension.NENetworkRuleProtocolUDP {
		return fmt.Errorf("bidirectional protocol = %v", got)
	}
	if got := bidirectional.MatchDirection(); got != networkextension.NETrafficDirectionOutbound {
		return fmt.Errorf("bidirectional direction = %v", got)
	}

	remoteEndpoint := network.NWEndpointCreateHost("203.0.113.1", "443")
	localEndpoint := network.NWEndpointCreateHost("192.0.2.1", "0")
	endpointHost := networkextension.NewNetworkRuleWithDestinationHostEndpointProtocol(remoteEndpoint, networkextension.NENetworkRuleProtocolTCP)
	if endpointHost.GetID() == 0 {
		return fmt.Errorf("create endpoint host rule")
	}
	if endpointHost.MatchRemoteHostOrNetworkEndpoint().GetID() == 0 {
		return fmt.Errorf("endpoint host rule remote endpoint is nil")
	}
	endpointNetwork := networkextension.NewNetworkRuleWithDestinationNetworkEndpointPrefixProtocol(remoteEndpoint, 24, networkextension.NENetworkRuleProtocolUDP)
	if endpointNetwork.GetID() == 0 {
		return fmt.Errorf("create endpoint network rule")
	}
	if got := endpointNetwork.MatchRemotePrefix(); got != 24 {
		return fmt.Errorf("endpoint network remote prefix = %d, want 24", got)
	}
	endpointBidirectional := networkextension.NewNetworkRuleWithRemoteNetworkEndpointRemotePrefixLocalNetworkEndpointLocalPrefixProtocolDirection(
		remoteEndpoint,
		24,
		localEndpoint,
		24,
		networkextension.NENetworkRuleProtocolAny,
		networkextension.NETrafficDirectionInbound,
	)
	if endpointBidirectional.GetID() == 0 {
		return fmt.Errorf("create endpoint remote/local network rule")
	}
	if endpointBidirectional.MatchRemoteHostOrNetworkEndpoint().GetID() == 0 {
		return fmt.Errorf("endpoint bidirectional remote endpoint is nil")
	}
	if endpointBidirectional.MatchLocalNetworkEndpoint().GetID() == 0 {
		return fmt.Errorf("endpoint bidirectional local endpoint is nil")
	}
	if got := endpointBidirectional.MatchDirection(); got != networkextension.NETrafficDirectionInbound {
		return fmt.Errorf("endpoint bidirectional direction = %v", got)
	}
	return nil
}

func checkDefaultNetworkStateObjects() error {
	path := networkextension.NewNWPath()
	if path.GetID() == 0 {
		return fmt.Errorf("create path")
	}
	if got := path.Status(); got != networkextension.NWPathStatusInvalid {
		return fmt.Errorf("path status = %v", got)
	}
	if path.IsExpensive() || path.IsConstrained() {
		return fmt.Errorf("new path has expensive or constrained state")
	}

	tcp := networkextension.NewNWTCPConnection()
	if tcp.GetID() == 0 {
		return fmt.Errorf("create tcp connection")
	}
	if got := tcp.State(); got != networkextension.NWTCPConnectionStateInvalid {
		return fmt.Errorf("tcp connection state = %v", got)
	}
	if tcp.IsViable() || tcp.HasBetterPath() || tcp.Error().GetID() != 0 {
		return fmt.Errorf("new tcp connection has unexpected live state")
	}
	if tcp.Endpoint().GetID() != 0 ||
		tcp.LocalAddress().GetID() != 0 ||
		tcp.RemoteAddress().GetID() != 0 ||
		tcp.ConnectedPath().GetID() != 0 ||
		tcp.TxtRecord().GetID() != 0 {
		return fmt.Errorf("new tcp connection has endpoint state")
	}

	udp := networkextension.NewNWUDPSession()
	if udp.GetID() == 0 {
		return fmt.Errorf("create udp session")
	}
	if got := udp.State(); got != networkextension.NWUDPSessionStateInvalid {
		return fmt.Errorf("udp session state = %v", got)
	}
	if udp.IsViable() || udp.HasBetterPath() || udp.MaximumDatagramLength() != 0 {
		return fmt.Errorf("new udp session has unexpected live state")
	}
	if udp.Endpoint().GetID() != 0 ||
		udp.ResolvedEndpoint().GetID() != 0 ||
		udp.CurrentPath().GetID() != 0 {
		return fmt.Errorf("new udp session has endpoint state")
	}
	return nil
}

func checkGeneratedArrayBlocks() error {
	data := foundation.NewDataWithBytesLength([]byte("datagram"))
	dataArray := objectivec.IObjectSliceToNSArray([]foundation.NSData{data})
	emptyArray := objc.Send[objc.ID](objc.ID(objc.GetClass("NSArray")), objc.Sel("array"))
	if dataArray == 0 || emptyArray == 0 {
		return fmt.Errorf("create NSArray values")
	}

	invoker := objcbridge.NewBlockInvoker()
	if err := expectNSDataArrayError(invoker, dataArray); err != nil {
		return err
	}
	if err := expectNSDataNWEndpointArrays(invoker, dataArray, emptyArray); err != nil {
		return err
	}
	return nil
}

func expectNSDataArrayError(invoker *objcbridge.BlockInvoker, dataArray objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNSDataArrayErrorBlock(func(datagrams *[]foundation.NSData, err error) {
		if err != nil {
			panic(err)
		}
		if datagrams == nil || len(*datagrams) != 1 {
			panic("unexpected datagrams")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectError(block, dataArray, 0); err != nil {
		return fmt.Errorf("invoke NSData array error block: %w", err)
	}
	if !called {
		return fmt.Errorf("NSData array error block was not called")
	}
	return nil
}

func expectNSDataNWEndpointArrays(invoker *objcbridge.BlockInvoker, dataArray objc.ID, endpointArray objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNSDataArrayNWEndpointArrayErrorBlock(func(datagrams *[]foundation.NSData, endpoints *[]networkextension.NWEndpoint, err error) {
		if err != nil {
			panic(err)
		}
		if datagrams == nil || len(*datagrams) != 1 {
			panic("unexpected datagrams")
		}
		if endpoints == nil || len(*endpoints) != 0 {
			panic("unexpected endpoints")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectObjectError(block, dataArray, endpointArray, 0); err != nil {
		return fmt.Errorf("invoke NSData endpoint array error block: %w", err)
	}
	if !called {
		return fmt.Errorf("NSData endpoint array error block was not called")
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-value-object-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-value-object-demo: NetworkExtension value object smoke ok")
}
