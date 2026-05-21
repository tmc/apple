// Command networkextension-manager-config-demo exercises local
// NetworkExtension manager objects without loading or saving preferences.
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

func runSmoke() error {
	demandRule, err := newDemandRule()
	if err != nil {
		return err
	}
	if err := checkVPNManager(demandRule); err != nil {
		return err
	}
	if err := checkDNSSettingsManager(demandRule); err != nil {
		return err
	}
	if err := checkIgnoreDemandRule(); err != nil {
		return err
	}
	if err := checkManagerArrayBlocks(); err != nil {
		return err
	}
	return nil
}

func newDemandRule() (networkextension.NEOnDemandRule, error) {
	rule := networkextension.NewNEOnDemandRuleDisconnect()
	if rule.GetID() == 0 {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("create on-demand disconnect rule")
	}
	rule.SetInterfaceTypeMatch(networkextension.NEOnDemandRuleInterfaceTypeEthernet)
	rule.SetDNSSearchDomainMatch([]string{"corp.example"})
	if got := rule.Action(); got != networkextension.NEOnDemandRuleActionDisconnect {
		return networkextension.NEOnDemandRule{}, fmt.Errorf("on-demand action = %v", got)
	}
	return rule.NEOnDemandRule, nil
}

func checkIgnoreDemandRule() error {
	rule := networkextension.NewNEOnDemandRuleIgnore()
	if rule.GetID() == 0 {
		return fmt.Errorf("create on-demand ignore rule")
	}
	rule.SetInterfaceTypeMatch(networkextension.NEOnDemandRuleInterfaceTypeAny)
	rule.SetDNSServerAddressMatch([]string{"192.0.2.54"})
	rule.SetSSIDMatch([]string{"Guest Wi-Fi"})
	if got := rule.Action(); got != networkextension.NEOnDemandRuleActionIgnore {
		return fmt.Errorf("ignore rule action = %v", got)
	}
	if got := rule.DNSServerAddressMatch(); !reflect.DeepEqual(got, []string{"192.0.2.54"}) {
		return fmt.Errorf("ignore rule DNS servers = %v", got)
	}
	if got := rule.SSIDMatch(); !reflect.DeepEqual(got, []string{"Guest Wi-Fi"}) {
		return fmt.Errorf("ignore rule SSID match = %v", got)
	}
	return nil
}

func checkVPNManager(demandRule networkextension.NEOnDemandRule) error {
	manager := networkextension.GetNEVPNManagerClass().SharedManager()
	if manager.GetID() == 0 {
		return fmt.Errorf("create vpn manager")
	}

	protocol := networkextension.NewNEVPNProtocolIKEv2()
	protocol.SetServerAddress("vpn-manager.example")
	protocol.SetRemoteIdentifier("remote.vpn-manager.example")
	protocol.SetLocalIdentifier("local.vpn-manager.example")
	protocol.SetUsername("vpn-user")
	protocol.SetDisconnectOnSleep(true)

	manager.SetLocalizedDescription("Example VPN Manager")
	manager.SetEnabled(true)
	manager.SetProtocolConfiguration(protocol)
	manager.SetOnDemandEnabled(true)
	manager.SetOnDemandRules([]networkextension.NEOnDemandRule{demandRule})

	if got := manager.LocalizedDescription(); got != "Example VPN Manager" {
		return fmt.Errorf("vpn manager description = %q", got)
	}
	if !manager.IsEnabled() || !manager.IsOnDemandEnabled() {
		return fmt.Errorf("vpn manager booleans did not round trip")
	}
	if got := manager.ProtocolConfiguration().ServerAddress(); got != "vpn-manager.example" {
		return fmt.Errorf("vpn manager protocol server = %q", got)
	}
	if got := manager.OnDemandRules(); len(got) != 1 {
		return fmt.Errorf("vpn manager on-demand rules = %d, want 1", len(got))
	}
	return nil
}

func checkDNSSettingsManager(demandRule networkextension.NEOnDemandRule) error {
	manager := networkextension.GetNEDNSSettingsManagerClass().SharedManager()
	if manager.GetID() == 0 {
		return fmt.Errorf("create dns settings manager")
	}

	settings := networkextension.NewDNSOverTLSSettingsWithServers([]string{"9.9.9.9", "149.112.112.112"})
	settings.SetServerName("dns.quad9.net")
	settings.SetMatchDomains([]string{"corp.example"})
	settings.SetAllowFailover(true)

	manager.SetLocalizedDescription("Example DNS Settings Manager")
	manager.SetDnsSettings(settings)
	manager.SetOnDemandRules([]networkextension.NEOnDemandRule{demandRule})

	if got := manager.LocalizedDescription(); got != "Example DNS Settings Manager" {
		return fmt.Errorf("dns settings manager description = %q", got)
	}
	if got := manager.DnsSettings().Servers(); !reflect.DeepEqual(got, []string{"9.9.9.9", "149.112.112.112"}) {
		return fmt.Errorf("dns settings servers = %v", got)
	}
	if got := manager.OnDemandRules(); len(got) != 1 {
		return fmt.Errorf("dns settings manager on-demand rules = %d, want 1", len(got))
	}
	return nil
}

func checkManagerArrayBlocks() error {
	invoker := objcbridge.NewBlockInvoker()
	emptyArray := objc.Send[objc.ID](objc.ID(objc.GetClass("NSArray")), objc.Sel("array"))
	if emptyArray == 0 {
		return fmt.Errorf("create empty NSArray")
	}

	if err := expectAppProxyManagers(invoker, emptyArray); err != nil {
		return err
	}
	if err := expectRelayManagers(invoker, emptyArray); err != nil {
		return err
	}
	if err := expectTransparentProxyManagers(invoker, emptyArray); err != nil {
		return err
	}
	if err := expectTunnelProviderManagers(invoker, emptyArray); err != nil {
		return err
	}
	return nil
}

func expectAppProxyManagers(invoker *objcbridge.BlockInvoker, array objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNEAppProxyProviderManagerArrayErrorBlock(func(managers *[]networkextension.NEAppProxyProviderManager, err error) {
		if err != nil {
			panic(err)
		}
		if managers == nil || len(*managers) != 0 {
			panic("unexpected app proxy managers")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectError(block, array, 0); err != nil {
		return fmt.Errorf("invoke app proxy manager array block: %w", err)
	}
	if !called {
		return fmt.Errorf("app proxy manager array block was not called")
	}
	return nil
}

func expectRelayManagers(invoker *objcbridge.BlockInvoker, array objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNERelayManagerArrayErrorBlock(func(managers *[]networkextension.NERelayManager, err error) {
		if err != nil {
			panic(err)
		}
		if managers == nil || len(*managers) != 0 {
			panic("unexpected relay managers")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectError(block, array, 0); err != nil {
		return fmt.Errorf("invoke relay manager array block: %w", err)
	}
	if !called {
		return fmt.Errorf("relay manager array block was not called")
	}
	return nil
}

func expectTransparentProxyManagers(invoker *objcbridge.BlockInvoker, array objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNETransparentProxyManagerArrayErrorBlock(func(managers *[]networkextension.NETransparentProxyManager, err error) {
		if err != nil {
			panic(err)
		}
		if managers == nil || len(*managers) != 0 {
			panic("unexpected transparent proxy managers")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectError(block, array, 0); err != nil {
		return fmt.Errorf("invoke transparent proxy manager array block: %w", err)
	}
	if !called {
		return fmt.Errorf("transparent proxy manager array block was not called")
	}
	return nil
}

func expectTunnelProviderManagers(invoker *objcbridge.BlockInvoker, array objc.ID) error {
	var called bool
	block, cleanup := networkextension.NewNETunnelProviderManagerArrayErrorBlock(func(managers *[]networkextension.NETunnelProviderManager, err error) {
		if err != nil {
			panic(err)
		}
		if managers == nil || len(*managers) != 0 {
			panic("unexpected tunnel provider managers")
		}
		called = true
	})
	defer cleanup()
	if err := invoker.ObjectError(block, array, 0); err != nil {
		return fmt.Errorf("invoke tunnel provider manager array block: %w", err)
	}
	if !called {
		return fmt.Errorf("tunnel provider manager array block was not called")
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-manager-config-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-manager-config-demo: NetworkExtension manager config smoke ok")
}
