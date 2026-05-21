// Command networkextension-policy-config-demo exercises local
// NetworkExtension policy/configuration objects without saving preferences.
package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objectivec"
)

func runSmoke() error {
	appRule, err := checkAppRules()
	if err != nil {
		return err
	}
	demandRule, err := checkOnDemandRules()
	if err != nil {
		return err
	}
	if err := checkTunnelProviderPolicy(appRule, demandRule); err != nil {
		return err
	}
	if err := checkNetworkPolicySettings(); err != nil {
		return err
	}
	if err := checkFilterVerdicts(); err != nil {
		return err
	}
	if err := checkURLFilter(); err != nil {
		return err
	}
	if err := checkPacketAndEndpointValues(); err != nil {
		return err
	}
	if err := checkRelayPolicy(demandRule); err != nil {
		return err
	}
	return nil
}

func checkAppRules() (networkextension.NEAppRule, error) {
	rule := networkextension.NewAppRuleWithSigningIdentifierDesignatedRequirement(
		"com.example.policy",
		"anchor apple generic",
	)
	if rule.GetID() == 0 {
		return networkextension.NEAppRule{}, fmt.Errorf("create app rule")
	}
	if got := rule.MatchSigningIdentifier(); got != "com.example.policy" {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule signing identifier = %q", got)
	}
	if got := rule.MatchDesignatedRequirement(); got != "anchor apple generic" {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule requirement = %q", got)
	}

	rule.SetMatchPath("/Applications/Example.app")
	rule.SetMatchDomains(foundation.NSArrayFromID(objectivec.StringSliceToNSArray([]string{"example.invalid"})))
	if got := rule.MatchPath(); got != "/Applications/Example.app" {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule match path = %q", got)
	}
	if got := rule.MatchDomains().Count(); got != 1 {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule match domains = %d, want 1", got)
	}

	tool := networkextension.NewAppRuleWithSigningIdentifier("com.example.policy.tool")
	rule.SetMatchTools([]networkextension.NEAppRule{tool})
	tools := rule.MatchTools()
	if len(tools) != 1 {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule match tools = %d, want 1", len(tools))
	}
	if got := tools[0].MatchSigningIdentifier(); got != "com.example.policy.tool" {
		return networkextension.NEAppRule{}, fmt.Errorf("app rule tool signing identifier = %q", got)
	}
	return rule, nil
}

func checkOnDemandRules() (networkextension.NEOnDemandRule, error) {
	domainRule := networkextension.NewEvaluateConnectionRuleWithMatchDomainsAndAction(
		[]string{"example.invalid", "corp.example"},
		networkextension.NEEvaluateConnectionRuleActionConnectIfNeeded,
	)
	if domainRule.GetID() == 0 {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("create evaluate connection rule")
	}
	domainRule.SetUseDNSServers([]string{"192.0.2.53"})
	domainRule.SetProbeURL(foundation.NewURLWithString("https://example.invalid/probe"))
	if got := domainRule.MatchDomains(); !reflect.DeepEqual(got, []string{"example.invalid", "corp.example"}) {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate rule match domains = %v", got)
	}
	if got := domainRule.UseDNSServers(); !reflect.DeepEqual(got, []string{"192.0.2.53"}) {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate rule DNS servers = %v", got)
	}
	if got := domainRule.ProbeURL().AbsoluteString(); got != "https://example.invalid/probe" {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate rule probe URL = %q", got)
	}
	if got := domainRule.Action(); got != networkextension.NEEvaluateConnectionRuleActionConnectIfNeeded {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate rule action = %v", got)
	}

	connect := networkextension.NewNEOnDemandRuleConnect()
	if connect.GetID() == 0 {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("create connect rule")
	}
	connect.SetInterfaceTypeMatch(networkextension.NEOnDemandRuleInterfaceTypeWiFi)
	connect.SetSSIDMatch([]string{"Example Wi-Fi"})
	connect.SetDNSSearchDomainMatch([]string{"corp.example"})
	connect.SetDNSServerAddressMatch([]string{"192.0.2.53"})
	connect.SetProbeURL(foundation.NewURLWithString("https://example.invalid/on-demand"))
	if got := connect.Action(); got != networkextension.NEOnDemandRuleActionConnect {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("connect action = %v", got)
	}
	if got := connect.InterfaceTypeMatch(); got != networkextension.NEOnDemandRuleInterfaceTypeWiFi {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("connect interface type = %v", got)
	}
	if got := connect.SSIDMatch(); !reflect.DeepEqual(got, []string{"Example Wi-Fi"}) {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("connect SSID match = %v", got)
	}

	evaluate := networkextension.NewNEOnDemandRuleEvaluateConnection()
	if evaluate.GetID() == 0 {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("create evaluate on-demand rule")
	}
	evaluate.SetConnectionRules([]networkextension.NEEvaluateConnectionRule{domainRule})
	if got := evaluate.Action(); got != networkextension.NEOnDemandRuleActionEvaluateConnection {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate on-demand action = %v", got)
	}
	rules := evaluate.ConnectionRules()
	if len(rules) != 1 {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("evaluate on-demand rules = %d, want 1", len(rules))
	}
	if got := rules[0].Action(); got != networkextension.NEEvaluateConnectionRuleActionConnectIfNeeded {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("nested evaluate rule action = %v", got)
	}
	return evaluate.NEOnDemandRule, nil
}

func checkTunnelProviderPolicy(appRule networkextension.NEAppRule, demandRule networkextension.NEOnDemandRule) error {
	manager := networkextension.GetNETunnelProviderManagerClass().ForPerAppVPN()
	if manager.GetID() == 0 {
		manager = networkextension.NewNETunnelProviderManager()
	}
	if manager.GetID() == 0 {
		return fmt.Errorf("create tunnel provider manager")
	}

	protocol := networkextension.NewNETunnelProviderProtocol()
	if protocol.GetID() == 0 {
		return fmt.Errorf("create tunnel provider protocol")
	}
	protocol.SetProviderBundleIdentifier("com.example.networkextension.provider")
	protocol.SetProviderConfiguration(foundation.GetNSDictionaryClass().Dictionary())
	protocol.SetServerAddress("vpn.example")
	protocol.SetUsername("example-user")
	protocol.SetDisconnectOnSleep(true)
	protocol.SetIncludeAllNetworks(true)
	protocol.SetExcludeLocalNetworks(true)
	if got := protocol.ProviderBundleIdentifier(); got != "com.example.networkextension.provider" {
		return fmt.Errorf("provider bundle identifier = %q", got)
	}
	if got := protocol.ProviderConfiguration().Count(); got != 0 {
		return fmt.Errorf("provider configuration count = %d, want 0", got)
	}
	if !protocol.IncludeAllNetworks() || !protocol.ExcludeLocalNetworks() {
		return fmt.Errorf("protocol route booleans did not round trip")
	}

	manager.SetLocalizedDescription("Example Tunnel Policy")
	manager.SetEnabled(true)
	manager.SetProtocolConfiguration(protocol)
	manager.SetOnDemandEnabled(true)
	manager.SetOnDemandRules([]networkextension.NEOnDemandRule{demandRule})
	manager.SetAppRules([]networkextension.NEAppRule{appRule})
	manager.SetAssociatedDomains([]string{"applinks:example.invalid"})
	manager.SetExcludedDomains([]string{"bypass.example"})
	manager.SetCalendarDomains([]string{"cal.example"})
	manager.SetContactsDomains([]string{"contacts.example"})
	manager.SetMailDomains([]string{"mail.example"})
	manager.SetSafariDomains([]string{"safari.example"})

	if got := manager.LocalizedDescription(); got != "Example Tunnel Policy" {
		return fmt.Errorf("manager description = %q", got)
	}
	if !manager.IsEnabled() || !manager.IsOnDemandEnabled() {
		return fmt.Errorf("manager booleans did not round trip")
	}
	if got := manager.ProtocolConfiguration().ServerAddress(); got != "vpn.example" {
		return fmt.Errorf("manager protocol server = %q", got)
	}
	if got := manager.AppRules(); len(got) != 1 {
		return fmt.Errorf("manager app rules = %d, want 1", len(got))
	}
	if got := manager.OnDemandRules(); len(got) != 1 {
		return fmt.Errorf("manager on-demand rules = %d, want 1", len(got))
	}
	if got := manager.AssociatedDomains(); !reflect.DeepEqual(got, []string{"applinks:example.invalid"}) {
		return fmt.Errorf("manager associated domains = %v", got)
	}
	if got := manager.SafariDomains(); !reflect.DeepEqual(got, []string{"safari.example"}) {
		return fmt.Errorf("manager Safari domains = %v", got)
	}
	return nil
}

func checkNetworkPolicySettings() error {
	host := networkextension.NewNWHostEndpointWithHostnamePort("example.invalid", "443")
	included := networkextension.NewNetworkRuleWithDestinationHostProtocol(host, networkextension.NENetworkRuleProtocolTCP)
	excludedHost := networkextension.NewNWHostEndpointWithHostnamePort("bypass.example", "443")
	excluded := networkextension.NewNetworkRuleWithDestinationHostProtocol(excludedHost, networkextension.NENetworkRuleProtocolTCP)

	transparent := networkextension.NewTransparentProxyNetworkSettingsWithTunnelRemoteAddress("198.51.100.10")
	if transparent.GetID() == 0 {
		return fmt.Errorf("create transparent proxy settings")
	}
	transparent.SetIncludedNetworkRules([]networkextension.NENetworkRule{included})
	transparent.SetExcludedNetworkRules([]networkextension.NENetworkRule{excluded})
	if got := transparent.TunnelRemoteAddress(); got != "198.51.100.10" {
		return fmt.Errorf("transparent proxy remote address = %q", got)
	}
	if got := transparent.IncludedNetworkRules(); len(got) != 1 {
		return fmt.Errorf("transparent proxy included rules = %d, want 1", len(got))
	}
	if got := transparent.ExcludedNetworkRules(); len(got) != 1 {
		return fmt.Errorf("transparent proxy excluded rules = %d, want 1", len(got))
	}

	ethernet := networkextension.NewEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu(
		"198.51.100.20",
		"02:00:5e:00:53:01",
		1500,
	)
	if ethernet.GetID() == 0 {
		return fmt.Errorf("create ethernet tunnel settings")
	}
	if got := ethernet.TunnelRemoteAddress(); got != "198.51.100.20" {
		return fmt.Errorf("ethernet tunnel remote address = %q", got)
	}
	if got := ethernet.EthernetAddress(); got != "02:00:5e:00:53:01" {
		return fmt.Errorf("ethernet address = %q", got)
	}
	return nil
}

func checkFilterVerdicts() error {
	data := networkextension.NewFilterDataVerdictWithPassBytesPeekBytes(128, 64)
	if data.GetID() == 0 {
		return fmt.Errorf("create data verdict")
	}
	data.SetShouldReport(true)
	data.SetStatisticsReportFrequency(networkextension.NEFilterReportFrequencyMedium)
	if !data.ShouldReport() {
		return fmt.Errorf("data verdict shouldReport = false")
	}
	if got := data.StatisticsReportFrequency(); got != networkextension.NEFilterReportFrequencyMedium {
		return fmt.Errorf("data verdict report frequency = %v", got)
	}

	flowClass := networkextension.GetNEFilterNewFlowVerdictClass()
	flow := flowClass.FilterDataVerdictWithFilterInboundPeekInboundBytesFilterOutboundPeekOutboundBytes(true, 256, true, 512)
	if flow.GetID() == 0 {
		return fmt.Errorf("create new-flow verdict")
	}
	flow.SetShouldReport(true)
	flow.SetStatisticsReportFrequency(networkextension.NEFilterReportFrequencyHigh)
	if got := flow.StatisticsReportFrequency(); got != networkextension.NEFilterReportFrequencyHigh {
		return fmt.Errorf("new-flow verdict report frequency = %v", got)
	}
	if flowClass.AllowVerdict().GetID() == 0 || flowClass.DropVerdict().GetID() == 0 || flowClass.PauseVerdict().GetID() == 0 {
		return fmt.Errorf("new-flow verdict class factories returned nil")
	}

	dataClass := networkextension.GetNEFilterDataVerdictClass()
	if dataClass.AllowVerdict().GetID() == 0 || dataClass.DropVerdict().GetID() == 0 || dataClass.PauseVerdict().GetID() == 0 {
		return fmt.Errorf("data verdict class factories returned nil")
	}
	return nil
}

func checkURLFilter() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	done := make(chan networkextension.NEURLFilterVerdict, 1)
	urlFilter := networkextension.GetNEURLFilterClass()
	url := foundation.NewURLWithString("https://example.com/")
	urlFilter.VerdictForURLCompletionHandler(url, func(verdict networkextension.NEURLFilterVerdict) {
		done <- verdict
	})
	select {
	case <-done:
	case <-ctx.Done():
		return fmt.Errorf("URL filter completion: %w", ctx.Err())
	}

	_, err := urlFilter.VerdictForURL(ctx, url)
	if err != nil {
		return fmt.Errorf("URL filter verdict: %w", err)
	}
	return nil
}

func checkPacketAndEndpointValues() error {
	packetData := foundation.NewDataWithBytesLength([]byte{0x45, 0x00, 0x00, 0x14})
	packet := networkextension.NewPacketWithDataProtocolFamily(packetData, 2)
	if packet.GetID() == 0 {
		return fmt.Errorf("create packet")
	}
	if got := packet.ProtocolFamily(); got != 2 {
		return fmt.Errorf("packet protocol family = %d", got)
	}
	if got := packet.Data().Length(); got != 4 {
		return fmt.Errorf("packet data length = %d", got)
	}
	if got := packet.Direction(); got != networkextension.NETrafficDirectionAny {
		return fmt.Errorf("packet direction = %v, want any", got)
	}
	if networkextension.GetNEFlowMetaDataClass().Class() == 0 {
		return fmt.Errorf("flow metadata class unavailable")
	}
	if got := packet.Metadata().GetID(); got != 0 {
		return fmt.Errorf("local packet metadata id = %v, want nil", got)
	}

	bonjour := networkextension.NewNWBonjourServiceEndpointWithNameTypeDomain("Example", "_demo._tcp", "local.")
	if bonjour.GetID() == 0 {
		return fmt.Errorf("create Bonjour endpoint")
	}
	if got := bonjour.Name(); got != "Example" {
		return fmt.Errorf("Bonjour endpoint name = %q", got)
	}
	if got := bonjour.Type(); got != "_demo._tcp" {
		return fmt.Errorf("Bonjour endpoint type = %q", got)
	}
	if got := bonjour.Domain(); got != "local." {
		return fmt.Errorf("Bonjour endpoint domain = %q", got)
	}
	return nil
}

func checkRelayPolicy(demandRule networkextension.NEOnDemandRule) error {
	if networkextension.GetNERelayClass().Class() == 0 {
		return fmt.Errorf("relay class unavailable")
	}
	if networkextension.GetNERelayManagerClass().Class() == 0 {
		return fmt.Errorf("relay manager class unavailable")
	}

	relay := networkextension.NewNERelay()
	if relay.GetID() == 0 {
		return fmt.Errorf("create relay")
	}
	relay.SetHTTP2RelayURL(foundation.NewURLWithString("https://relay.example/h2"))
	relay.SetHTTP3RelayURL(foundation.NewURLWithString("https://relay.example/h3"))
	relay.SetDnsOverHTTPSURL(foundation.NewURLWithString("https://dns.example/dns-query"))
	relay.SetRawPublicKeys([]foundation.NSData{foundation.NewDataWithBytesLength([]byte("raw-key"))})
	relay.SetAdditionalHTTPHeaderFields(foundation.GetNSDictionaryClass().Dictionary())
	relay.SetIdentityData(foundation.NewDataWithBytesLength([]byte("relay-identity")))
	relay.SetIdentityDataPassword("relay-password")
	relay.SetSyntheticDNSAnswerIPv4Prefix("192.0.2.0/24")
	relay.SetSyntheticDNSAnswerIPv6Prefix("2001:db8::/32")
	if got := relay.HTTP2RelayURL().AbsoluteString(); got != "https://relay.example/h2" {
		return fmt.Errorf("relay HTTP/2 URL = %q", got)
	}
	if got := relay.HTTP3RelayURL().AbsoluteString(); got != "https://relay.example/h3" {
		return fmt.Errorf("relay HTTP/3 URL = %q", got)
	}
	if got := relay.DnsOverHTTPSURL().AbsoluteString(); got != "https://dns.example/dns-query" {
		return fmt.Errorf("relay DoH URL = %q", got)
	}
	if got := relay.RawPublicKeys(); len(got) != 1 || got[0].Length() != 7 {
		return fmt.Errorf("relay raw public keys = %d", len(got))
	}
	if got := relay.AdditionalHTTPHeaderFields().Count(); got != 0 {
		return fmt.Errorf("relay header field count = %d, want 0", got)
	}
	if got := relay.IdentityData().Length(); got != 14 {
		return fmt.Errorf("relay identity data length = %d, want 14", got)
	}
	if got := relay.IdentityDataPassword(); got != "relay-password" {
		return fmt.Errorf("relay identity password = %q", got)
	}
	if got := relay.SyntheticDNSAnswerIPv4Prefix(); got != "192.0.2.0/24" {
		return fmt.Errorf("relay IPv4 prefix = %q", got)
	}
	if got := relay.SyntheticDNSAnswerIPv6Prefix(); got != "2001:db8::/32" {
		return fmt.Errorf("relay IPv6 prefix = %q", got)
	}

	manager := networkextension.NewNERelayManager()
	if manager.GetID() == 0 {
		return fmt.Errorf("create relay manager")
	}
	manager.SetLocalizedDescription("Example Relay Policy")
	manager.SetEnabled(true)
	manager.SetRelays([]networkextension.NERelay{relay})
	manager.SetMatchDomains([]string{"relay.example"})
	manager.SetExcludedDomains([]string{"bypass.example"})
	manager.SetMatchFQDNs([]string{"api.relay.example"})
	manager.SetExcludedFQDNs([]string{"direct.relay.example"})
	manager.SetAllowDNSFailover(true)
	manager.SetUIToggleEnabled(true)
	manager.SetOnDemandRules([]networkextension.NEOnDemandRule{demandRule})

	if got := manager.LocalizedDescription(); got != "Example Relay Policy" {
		return fmt.Errorf("relay manager description = %q", got)
	}
	if !manager.IsEnabled() || !manager.IsDNSFailoverAllowed() || !manager.IsUIToggleEnabled() {
		return fmt.Errorf("relay manager booleans did not round trip")
	}
	if got := manager.Relays(); len(got) != 1 {
		return fmt.Errorf("relay manager relays = %d, want 1", len(got))
	}
	if got := manager.MatchDomains(); !reflect.DeepEqual(got, []string{"relay.example"}) {
		return fmt.Errorf("relay manager match domains = %v", got)
	}
	if got := manager.MatchFQDNs(); !reflect.DeepEqual(got, []string{"api.relay.example"}) {
		return fmt.Errorf("relay manager match FQDNs = %v", got)
	}
	if got := manager.ExcludedDomains(); !reflect.DeepEqual(got, []string{"bypass.example"}) {
		return fmt.Errorf("relay manager excluded domains = %v", got)
	}
	if got := manager.ExcludedFQDNs(); !reflect.DeepEqual(got, []string{"direct.relay.example"}) {
		return fmt.Errorf("relay manager excluded FQDNs = %v", got)
	}
	if got := manager.OnDemandRules(); len(got) != 1 {
		return fmt.Errorf("relay manager on-demand rules = %d, want 1", len(got))
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-policy-config-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-policy-config-demo: NetworkExtension policy/config smoke ok")
}
