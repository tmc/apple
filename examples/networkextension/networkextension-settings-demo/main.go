// Command networkextension-settings-demo exercises local NetworkExtension
// settings objects without installing a Network Extension.
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/networkextension"
)

func runSmoke() error {
	if err := checkPacketTunnelSettings(); err != nil {
		return err
	}
	if err := checkDNSSettings(); err != nil {
		return err
	}
	if err := checkFilterSettings(); err != nil {
		return err
	}
	if err := checkVPNSettings(); err != nil {
		return err
	}
	return nil
}

func checkPacketTunnelSettings() error {
	settings := networkextension.NewPacketTunnelNetworkSettingsWithTunnelRemoteAddress("198.51.100.1")
	if settings.GetID() == 0 {
		return fmt.Errorf("create packet tunnel settings")
	}
	if got := settings.TunnelRemoteAddress(); got != "198.51.100.1" {
		return fmt.Errorf("tunnel remote address = %q, want %q", got, "198.51.100.1")
	}

	ipv4 := networkextension.NewIPv4SettingsWithAddressesSubnetMasks(
		[]string{"10.0.0.2"},
		[]string{"255.255.255.0"},
	)
	ipv4.SetRouter("10.0.0.1")
	included4 := networkextension.NewIPv4RouteWithDestinationAddressSubnetMask("0.0.0.0", "0.0.0.0")
	included4.SetGatewayAddress("10.0.0.1")
	excluded4 := networkextension.NewIPv4RouteWithDestinationAddressSubnetMask("192.0.2.0", "255.255.255.0")
	ipv4.SetIncludedRoutes([]networkextension.NEIPv4Route{included4})
	ipv4.SetExcludedRoutes([]networkextension.NEIPv4Route{excluded4})
	settings.SetIPv4Settings(ipv4)
	auto4 := networkextension.GetNEIPv4SettingsClass().SettingsWithAutomaticAddressing()
	if auto4.GetID() == 0 {
		return fmt.Errorf("create automatic IPv4 settings")
	}

	ipv6Prefix := foundation.NewNumberWithInt(64)
	ipv6 := networkextension.NewIPv6SettingsWithAddressesNetworkPrefixLengths(
		[]string{"2001:db8::2"},
		[]foundation.NSNumber{ipv6Prefix},
	)
	included6 := networkextension.NewIPv6RouteWithDestinationAddressNetworkPrefixLength("::", foundation.NewNumberWithInt(0))
	included6.SetGatewayAddress("2001:db8::1")
	excluded6 := networkextension.NewIPv6RouteWithDestinationAddressNetworkPrefixLength("2001:db8:1::", foundation.NewNumberWithInt(64))
	ipv6.SetIncludedRoutes([]networkextension.NEIPv6Route{included6})
	ipv6.SetExcludedRoutes([]networkextension.NEIPv6Route{excluded6})
	settings.SetIPv6Settings(ipv6)
	auto6 := networkextension.GetNEIPv6SettingsClass().SettingsWithAutomaticAddressing()
	if auto6.GetID() == 0 {
		return fmt.Errorf("create automatic IPv6 settings")
	}
	linkLocal6 := networkextension.GetNEIPv6SettingsClass().SettingsWithLinkLocalAddressing()
	if linkLocal6.GetID() == 0 {
		return fmt.Errorf("create link-local IPv6 settings")
	}

	dns := networkextension.NewDNSSettingsWithServers([]string{"1.1.1.1", "2606:4700:4700::1111"})
	dns.SetSearchDomains([]string{"corp.example"})
	dns.SetDomainName("corp.example")
	dns.SetMatchDomains([]string{"corp.example"})
	dns.SetMatchDomainsNoSearch(true)
	dns.SetAllowFailover(true)
	settings.SetDNSSettings(dns)

	proxy := networkextension.NewNEProxySettings()
	httpServer := networkextension.NewProxyServerWithAddressPort("127.0.0.1", 8080)
	httpsServer := networkextension.NewProxyServerWithAddressPort("127.0.0.1", 8443)
	httpServer.SetAuthenticationRequired(true)
	httpServer.SetUsername("proxy-user")
	httpServer.SetPassword("proxy-password")
	httpsServer.SetAuthenticationRequired(true)
	httpsServer.SetUsername("secure-proxy-user")
	httpsServer.SetPassword("secure-proxy-password")
	proxy.SetAutoProxyConfigurationEnabled(true)
	proxy.SetProxyAutoConfigurationURL(foundation.NewURLWithString("https://proxy.example/proxy.pac"))
	proxy.SetProxyAutoConfigurationJavaScript("function FindProxyForURL(url, host) { return 'DIRECT'; }")
	proxy.SetHTTPEnabled(true)
	proxy.SetHTTPServer(httpServer)
	proxy.SetHTTPSEnabled(true)
	proxy.SetHTTPSServer(httpsServer)
	proxy.SetExcludeSimpleHostnames(true)
	proxy.SetExceptionList([]string{"localhost", "*.local"})
	proxy.SetMatchDomains([]string{"corp.example"})
	settings.SetProxySettings(proxy)

	baseTunnel := networkextension.NewTunnelNetworkSettingsWithTunnelRemoteAddress("203.0.113.10")
	baseTunnel.SetDNSSettings(dns)
	baseTunnel.SetProxySettings(proxy)
	if got := baseTunnel.TunnelRemoteAddress(); got != "203.0.113.10" {
		return fmt.Errorf("base tunnel remote address = %q", got)
	}
	if got := baseTunnel.DNSSettings().Servers(); !reflect.DeepEqual(got, []string{"1.1.1.1", "2606:4700:4700::1111"}) {
		return fmt.Errorf("base tunnel DNS servers = %v", got)
	}
	if got := baseTunnel.ProxySettings().HTTPServer().Port(); got != 8080 {
		return fmt.Errorf("base tunnel proxy port = %d, want 8080", got)
	}

	settings.SetMTU(foundation.NewNumberWithInt(1280))
	settings.SetTunnelOverheadBytes(foundation.NewNumberWithInt(80))

	if got := settings.IPv4Settings().Addresses(); !reflect.DeepEqual(got, []string{"10.0.0.2"}) {
		return fmt.Errorf("IPv4 addresses = %v, want [10.0.0.2]", got)
	}
	if got := settings.IPv4Settings().IncludedRoutes()[0].GatewayAddress(); got != "10.0.0.1" {
		return fmt.Errorf("IPv4 included gateway = %q, want %q", got, "10.0.0.1")
	}
	if got := settings.IPv6Settings().NetworkPrefixLengths()[0].IntValue(); got != 64 {
		return fmt.Errorf("IPv6 prefix length = %d, want 64", got)
	}
	if got := settings.DNSSettings().Servers(); !reflect.DeepEqual(got, []string{"1.1.1.1", "2606:4700:4700::1111"}) {
		return fmt.Errorf("DNS servers = %v", got)
	}
	if got := settings.DNSSettings().AllowFailover(); !got {
		return fmt.Errorf("DNS allow failover = false")
	}
	gotProxy := settings.ProxySettings()
	if !gotProxy.AutoProxyConfigurationEnabled() ||
		gotProxy.ProxyAutoConfigurationURL().AbsoluteString() != "https://proxy.example/proxy.pac" ||
		gotProxy.ProxyAutoConfigurationJavaScript() == "" {
		return fmt.Errorf("proxy auto-configuration did not round trip")
	}
	if !gotProxy.HTTPEnabled() || !gotProxy.HTTPSEnabled() {
		return fmt.Errorf("proxy HTTP enable flags did not round trip")
	}
	if !gotProxy.ExcludeSimpleHostnames() ||
		!reflect.DeepEqual(gotProxy.ExceptionList(), []string{"localhost", "*.local"}) ||
		!reflect.DeepEqual(gotProxy.MatchDomains(), []string{"corp.example"}) {
		return fmt.Errorf("proxy bypass lists did not round trip")
	}
	if got := gotProxy.HTTPServer().Address(); got != "127.0.0.1" {
		return fmt.Errorf("HTTP proxy address = %q, want 127.0.0.1", got)
	}
	if got := gotProxy.HTTPServer().Port(); got != 8080 {
		return fmt.Errorf("HTTP proxy port = %d, want 8080", got)
	}
	if got := gotProxy.HTTPServer().Username(); got != "proxy-user" {
		return fmt.Errorf("HTTP proxy username = %q, want proxy-user", got)
	}
	if !gotProxy.HTTPServer().AuthenticationRequired() || gotProxy.HTTPServer().Password() != "proxy-password" {
		return fmt.Errorf("HTTP proxy authentication did not round trip")
	}
	if got := gotProxy.HTTPSServer().Port(); got != 8443 {
		return fmt.Errorf("HTTPS proxy port = %d, want 8443", got)
	}
	if !gotProxy.HTTPSServer().AuthenticationRequired() || gotProxy.HTTPSServer().Username() != "secure-proxy-user" {
		return fmt.Errorf("HTTPS proxy authentication did not round trip")
	}
	if got := settings.MTU().IntValue(); got != 1280 {
		return fmt.Errorf("MTU = %d, want 1280", got)
	}
	if got := settings.TunnelOverheadBytes().IntValue(); got != 80 {
		return fmt.Errorf("tunnel overhead = %d, want 80", got)
	}
	return nil
}

func checkDNSSettings() error {
	doh := networkextension.NewDNSOverHTTPSSettingsWithServers([]string{"1.1.1.1"})
	doh.SetServerURL(foundation.NewURLWithString("https://dns.example/dns-query"))
	doh.SetIdentityReference(foundation.NewDataWithBytesLength([]byte("doh-identity")))
	if got := doh.DnsProtocol(); got != networkextension.NEDNSProtocolHTTPS {
		return fmt.Errorf("DoH protocol = %v, want %v", got, networkextension.NEDNSProtocolHTTPS)
	}
	if got := doh.ServerURL().AbsoluteString(); got != "https://dns.example/dns-query" {
		return fmt.Errorf("DoH server URL = %q", got)
	}
	if got := doh.IdentityReference().Length(); got != 12 {
		return fmt.Errorf("DoH identity reference length = %d, want 12", got)
	}

	dot := networkextension.NewDNSOverTLSSettingsWithServers([]string{"9.9.9.9"})
	dot.SetServerName("dns.quad9.net")
	dot.SetAllowFailover(true)
	dot.SetIdentityReference(foundation.NewDataWithBytesLength([]byte("dot-identity")))
	if got := dot.DnsProtocol(); got != networkextension.NEDNSProtocolTLS {
		return fmt.Errorf("DoT protocol = %v, want %v", got, networkextension.NEDNSProtocolTLS)
	}
	if got := dot.ServerName(); got != "dns.quad9.net" {
		return fmt.Errorf("DoT server name = %q", got)
	}
	if !dot.AllowFailover() {
		return fmt.Errorf("DoT allow failover = false, want true")
	}
	if got := dot.IdentityReference().Length(); got != 12 {
		return fmt.Errorf("DoT identity reference length = %d, want 12", got)
	}
	return nil
}

func checkFilterSettings() error {
	host := networkextension.NewNWHostEndpointWithHostnamePort("example.com", "443")
	if got := host.Hostname(); got != "example.com" {
		return fmt.Errorf("host endpoint hostname = %q", got)
	}
	if got := host.Port(); got != "443" {
		return fmt.Errorf("host endpoint port = %q", got)
	}
	rule := networkextension.NewNetworkRuleWithDestinationHostProtocol(host, networkextension.NENetworkRuleProtocolTCP)
	filterRule := networkextension.NewFilterRuleWithNetworkRuleAction(rule, networkextension.NEFilterActionFilterData)
	settings := networkextension.NewFilterSettingsWithRulesDefaultAction(
		[]networkextension.NEFilterRule{filterRule},
		networkextension.NEFilterActionAllow,
	)
	if got := filterRule.Action(); got != networkextension.NEFilterActionFilterData {
		return fmt.Errorf("filter rule action = %v", got)
	}
	if got := settings.DefaultAction(); got != networkextension.NEFilterActionAllow {
		return fmt.Errorf("filter default action = %v", got)
	}
	if got := settings.Rules(); len(got) != 1 {
		return fmt.Errorf("filter rules = %d, want 1", len(got))
	}
	if got := filterRule.NetworkRule().MatchProtocol(); got != networkextension.NENetworkRuleProtocolTCP {
		return fmt.Errorf("filter rule network protocol = %v", got)
	}

	config := networkextension.NewNEFilterProviderConfiguration()
	config.SetFilterSockets(true)
	config.SetFilterPackets(true)
	config.SetServerAddress("filter.example")
	config.SetUsername("filter-user")
	config.SetOrganization("Example")
	config.SetVendorConfiguration(foundation.GetNSDictionaryClass().Dictionary())
	config.SetPasswordReference(foundation.NewDataWithBytesLength([]byte("password-ref")))
	config.SetIdentityReference(foundation.NewDataWithBytesLength([]byte("identity-ref")))
	config.SetFilterDataProviderBundleIdentifier("com.example.filter.data")
	config.SetFilterPacketProviderBundleIdentifier("com.example.filter.packet")
	if !config.FilterSockets() || !config.FilterPackets() {
		return fmt.Errorf("filter packet/socket booleans did not round trip")
	}
	if got := config.ServerAddress(); got != "filter.example" {
		return fmt.Errorf("filter server address = %q", got)
	}
	if got := config.Username(); got != "filter-user" {
		return fmt.Errorf("filter username = %q", got)
	}
	if got := config.Organization(); got != "Example" {
		return fmt.Errorf("filter organization = %q", got)
	}
	if got := config.VendorConfiguration().Count(); got != 0 {
		return fmt.Errorf("filter vendor config count = %d, want 0", got)
	}
	if got := config.PasswordReference().Length(); got != 12 {
		return fmt.Errorf("filter password reference length = %d, want 12", got)
	}
	if got := config.IdentityReference().Length(); got != 12 {
		return fmt.Errorf("filter identity reference length = %d, want 12", got)
	}
	if got := config.FilterDataProviderBundleIdentifier(); got != "com.example.filter.data" {
		return fmt.Errorf("filter data provider bundle id = %q", got)
	}
	if got := config.FilterPacketProviderBundleIdentifier(); got != "com.example.filter.packet" {
		return fmt.Errorf("filter packet provider bundle id = %q", got)
	}
	return nil
}

func checkVPNSettings() error {
	proxy := networkextension.NewNEProxySettings()
	proxy.SetHTTPEnabled(true)
	proxy.SetHTTPServer(networkextension.NewProxyServerWithAddressPort("127.0.0.1", 8080))

	ipsec := networkextension.NewNEVPNProtocolIPSec()
	ipsec.SetServerAddress("vpn.example")
	ipsec.SetUsername("vpn-user")
	ipsec.SetDisconnectOnSleep(true)
	ipsec.SetPasswordReference(foundation.NewDataWithBytesLength([]byte("password-ref")))
	ipsec.SetIdentityReference(foundation.NewDataWithBytesLength([]byte("identity-ref")))
	ipsec.SetIdentityData(foundation.NewDataWithBytesLength([]byte("identity-data")))
	ipsec.SetIdentityDataPassword("identity-password")
	ipsec.SetProxySettings(proxy)
	ipsec.SetAuthenticationMethod(networkextension.NEVPNIKEAuthenticationMethodSharedSecret)
	ipsec.SetUseExtendedAuthentication(true)
	ipsec.SetLocalIdentifier("local.example")
	ipsec.SetRemoteIdentifier("remote.example")
	ipsec.SetSharedSecretReference(foundation.NewDataWithBytesLength([]byte("shared-secret")))
	ipsec.SetIncludeAllNetworks(true)
	ipsec.SetExcludeLocalNetworks(true)
	ipsec.SetExcludeAPNs(true)
	ipsec.SetExcludeCellularServices(true)
	ipsec.SetExcludeDeviceCommunication(true)
	ipsec.SetEnforceRoutes(true)
	if got := ipsec.ServerAddress(); got != "vpn.example" {
		return fmt.Errorf("IPSec server address = %q", got)
	}
	if got := ipsec.Username(); got != "vpn-user" {
		return fmt.Errorf("IPSec username = %q", got)
	}
	if !ipsec.DisconnectOnSleep() ||
		!ipsec.IncludeAllNetworks() ||
		!ipsec.ExcludeLocalNetworks() ||
		!ipsec.ExcludeAPNs() ||
		!ipsec.ExcludeCellularServices() ||
		!ipsec.ExcludeDeviceCommunication() ||
		!ipsec.EnforceRoutes() {
		return fmt.Errorf("IPSec route booleans did not round trip")
	}
	if got := ipsec.PasswordReference().Length(); got != 12 {
		return fmt.Errorf("IPSec password reference length = %d, want 12", got)
	}
	if got := ipsec.IdentityReference().Length(); got != 12 {
		return fmt.Errorf("IPSec identity reference length = %d, want 12", got)
	}
	if got := ipsec.IdentityData().Length(); got != 13 {
		return fmt.Errorf("IPSec identity data length = %d, want 13", got)
	}
	if got := ipsec.IdentityDataPassword(); got != "identity-password" {
		return fmt.Errorf("IPSec identity data password = %q", got)
	}
	if got := ipsec.ProxySettings().HTTPServer().Port(); got != 8080 {
		return fmt.Errorf("IPSec proxy port = %d, want 8080", got)
	}
	if got := ipsec.AuthenticationMethod(); got != networkextension.NEVPNIKEAuthenticationMethodSharedSecret {
		return fmt.Errorf("IPSec auth method = %v", got)
	}
	if !ipsec.UseExtendedAuthentication() ||
		ipsec.LocalIdentifier() != "local.example" ||
		ipsec.RemoteIdentifier() != "remote.example" ||
		ipsec.SharedSecretReference().Length() != 13 {
		return fmt.Errorf("IPSec identifiers or shared secret did not round trip")
	}

	ike := networkextension.NewNEVPNProtocolIKEv2()
	ike.SetServerAddress("ike.example")
	ike.SetRemoteIdentifier("ike.remote.example")
	ike.SetLocalIdentifier("ike.local.example")
	ike.SetServerCertificateIssuerCommonName("Example Issuer")
	ike.SetServerCertificateCommonName("vpn.example")
	ike.SetCertificateType(networkextension.NEVPNIKEv2CertificateTypeRSA)
	ike.SetDeadPeerDetectionRate(networkextension.NEVPNIKEv2DeadPeerDetectionRateMedium)
	ike.SetDisableMOBIKE(true)
	ike.SetDisableRedirect(true)
	ike.SetEnablePFS(true)
	ike.SetEnableRevocationCheck(true)
	ike.SetStrictRevocationCheck(true)
	ike.SetUseConfigurationAttributeInternalIPSubnet(true)
	ike.SetMinimumTLSVersion(networkextension.NEVPNIKEv2TLSVersion1_2)
	ike.SetMaximumTLSVersion(networkextension.NEVPNIKEv2TLSVersion1_2)
	ike.SetMtu(1280)

	ike.IKESecurityAssociationParameters().SetEncryptionAlgorithm(networkextension.NEVPNIKEv2EncryptionAlgorithmAES256)
	ike.IKESecurityAssociationParameters().SetIntegrityAlgorithm(networkextension.NEVPNIKEv2IntegrityAlgorithmSHA256)
	ike.IKESecurityAssociationParameters().SetDiffieHellmanGroup(networkextension.NEVPNIKEv2DiffieHellmanGroup14)
	ike.IKESecurityAssociationParameters().SetLifetimeMinutes(60)
	ike.ChildSecurityAssociationParameters().SetEncryptionAlgorithm(networkextension.NEVPNIKEv2EncryptionAlgorithmAES128GCM)
	ike.ChildSecurityAssociationParameters().SetDiffieHellmanGroup(networkextension.NEVPNIKEv2DiffieHellmanGroup19)

	if got := ike.ServerAddress(); got != "ike.example" {
		return fmt.Errorf("IKEv2 server address = %q", got)
	}
	if got := ike.CertificateType(); got != networkextension.NEVPNIKEv2CertificateTypeRSA {
		return fmt.Errorf("IKEv2 certificate type = %v", got)
	}
	if ike.ServerCertificateIssuerCommonName() != "Example Issuer" ||
		ike.ServerCertificateCommonName() != "vpn.example" ||
		ike.DeadPeerDetectionRate() != networkextension.NEVPNIKEv2DeadPeerDetectionRateMedium ||
		!ike.DisableMOBIKE() ||
		!ike.DisableRedirect() ||
		!ike.EnablePFS() ||
		!ike.EnableRevocationCheck() ||
		!ike.StrictRevocationCheck() ||
		!ike.UseConfigurationAttributeInternalIPSubnet() ||
		ike.MinimumTLSVersion() != networkextension.NEVPNIKEv2TLSVersion1_2 ||
		ike.MaximumTLSVersion() != networkextension.NEVPNIKEv2TLSVersion1_2 {
		return fmt.Errorf("IKEv2 booleans or certificate names did not round trip")
	}
	if got := ike.IKESecurityAssociationParameters().EncryptionAlgorithm(); got != networkextension.NEVPNIKEv2EncryptionAlgorithmAES256 {
		return fmt.Errorf("IKE SA encryption = %v", got)
	}
	if got := ike.Mtu(); got != 1280 {
		return fmt.Errorf("IKEv2 MTU = %d, want 1280", got)
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-settings-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-settings-demo: NetworkExtension settings smoke ok")
}
