// Command networkextension-provider-wrapper-demo exercises generated provider
// wrapper methods against local Objective-C callback objects.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/network"
	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

type providerCounts struct {
	packetStart     uint64
	packetStop      uint64
	packetCancel    uint64
	packetFlow      uint64
	packetInterface uint64
	appStart        uint64
	appStop         uint64
	appCancel       uint64
	appFlow         uint64
	appUDPFlow      uint64
}

func registerPacketTunnelClass(name string, invoker *objcbridge.BlockInvoker, counts *providerCounts) (objc.Class, error) {
	super := objc.GetClass("NSObject")
	if super == 0 {
		return 0, errors.New("lookup NSObject")
	}
	cls, err := objc.RegisterClass(name, super, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("register %s: %w", name, err)
	}
	methods := []objc.MethodDef{
		{
			Cmd: objc.RegisterName("startTunnelWithOptions:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, options objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.packetStart, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("stopTunnelWithReason:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, reason networkextension.NEProviderStopReason, completion objc.ID) {
				atomic.AddUint64(&counts.packetStop, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("cancelTunnelWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, err objc.ID) {
				atomic.AddUint64(&counts.packetCancel, 1)
			},
		},
		{
			Cmd: objc.RegisterName("packetFlow"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.packetFlow, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("virtualInterface"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.packetInterface, 1)
				return 0
			},
		},
	}
	return cls, objcbridge.AddMethods(cls, name, methods)
}

func registerAppProxyClass(name string, invoker *objcbridge.BlockInvoker, counts *providerCounts) (objc.Class, error) {
	super := objc.GetClass("NSObject")
	if super == 0 {
		return 0, errors.New("lookup NSObject")
	}
	cls, err := objc.RegisterClass(name, super, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("register %s: %w", name, err)
	}
	methods := []objc.MethodDef{
		{
			Cmd: objc.RegisterName("startProxyWithOptions:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, options objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.appStart, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("stopProxyWithReason:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, reason networkextension.NEProviderStopReason, completion objc.ID) {
				atomic.AddUint64(&counts.appStop, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("cancelProxyWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, err objc.ID) {
				atomic.AddUint64(&counts.appCancel, 1)
			},
		},
		{
			Cmd: objc.RegisterName("handleNewFlow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID) bool {
				atomic.AddUint64(&counts.appFlow, 1)
				return true
			},
		},
		{
			Cmd: objc.RegisterName("handleNewUDPFlow:initialRemoteFlowEndpoint:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, endpoint objc.ID) bool {
				atomic.AddUint64(&counts.appUDPFlow, 1)
				return true
			},
		},
	}
	return cls, objcbridge.AddMethods(cls, name, methods)
}

func runSmoke() error {
	var counts providerCounts
	if err := exercisePacketTunnelProvider(&counts); err != nil {
		return err
	}
	if err := exerciseAppProxyProvider(&counts); err != nil {
		return err
	}
	if err := checkProviderSubclassAvailability(); err != nil {
		return err
	}
	return checkCounts(&counts)
}

func exercisePacketTunnelProvider(counts *providerCounts) error {
	name := fmt.Sprintf("GoPacketTunnelProviderWrappers_%d_%d", os.Getpid(), time.Now().UnixNano())
	cls, err := registerPacketTunnelClass(name, objcbridge.NewBlockInvoker(), counts)
	if err != nil {
		return err
	}
	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("new"))
	if id == 0 {
		return fmt.Errorf("instantiate %s", name)
	}
	defer objc.Send[struct{}](id, objc.Sel("release"))

	if err := checkSelectors(id, []string{
		"startTunnelWithOptions:completionHandler:",
		"stopTunnelWithReason:completionHandler:",
		"cancelTunnelWithError:",
		"packetFlow",
		"virtualInterface",
	}); err != nil {
		return err
	}

	provider := networkextension.NEPacketTunnelProviderFromID(id)
	var started uint64
	provider.StartTunnelWithOptionsCompletionHandler(nil, func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&started, 1)
	})
	if got := atomic.LoadUint64(&started); got != 1 {
		return fmt.Errorf("packet tunnel direct starts = %d, want 1", got)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := provider.StartTunnelWithOptions(ctx, nil); err != nil {
		return fmt.Errorf("packet tunnel context start: %w", err)
	}

	var stopped uint64
	provider.StopTunnelWithReasonCompletionHandler(networkextension.NEProviderStopReasonUserInitiated, func() {
		atomic.AddUint64(&stopped, 1)
	})
	if got := atomic.LoadUint64(&stopped); got != 1 {
		return fmt.Errorf("packet tunnel direct stops = %d, want 1", got)
	}
	if err := provider.StopTunnelWithReason(ctx, networkextension.NEProviderStopReasonUserInitiated); err != nil {
		return fmt.Errorf("packet tunnel context stop: %w", err)
	}

	provider.CancelTunnelWithError(foundation.NSErrorFromID(0))
	if flow := provider.PacketFlow(); flow.GetID() != 0 {
		return fmt.Errorf("packet flow id = %v, want nil", flow.GetID())
	}
	if iface := provider.VirtualInterface(); iface.GetID() != 0 {
		return fmt.Errorf("virtual interface id = %v, want nil", iface.GetID())
	}
	return nil
}

func exerciseAppProxyProvider(counts *providerCounts) error {
	name := fmt.Sprintf("GoAppProxyProviderWrappers_%d_%d", os.Getpid(), time.Now().UnixNano())
	cls, err := registerAppProxyClass(name, objcbridge.NewBlockInvoker(), counts)
	if err != nil {
		return err
	}
	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("new"))
	if id == 0 {
		return fmt.Errorf("instantiate %s", name)
	}
	defer objc.Send[struct{}](id, objc.Sel("release"))

	if err := checkSelectors(id, []string{
		"startProxyWithOptions:completionHandler:",
		"stopProxyWithReason:completionHandler:",
		"cancelProxyWithError:",
		"handleNewFlow:",
		"handleNewUDPFlow:initialRemoteFlowEndpoint:",
	}); err != nil {
		return err
	}

	provider := networkextension.NEAppProxyProviderFromID(id)
	var started uint64
	provider.StartProxyWithOptionsCompletionHandler(nil, func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&started, 1)
	})
	if got := atomic.LoadUint64(&started); got != 1 {
		return fmt.Errorf("app proxy direct starts = %d, want 1", got)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := provider.StartProxyWithOptions(ctx, nil); err != nil {
		return fmt.Errorf("app proxy context start: %w", err)
	}

	var stopped uint64
	provider.StopProxyWithReasonCompletionHandler(networkextension.NEProviderStopReasonUserInitiated, func() {
		atomic.AddUint64(&stopped, 1)
	})
	if got := atomic.LoadUint64(&stopped); got != 1 {
		return fmt.Errorf("app proxy direct stops = %d, want 1", got)
	}
	if err := provider.StopProxyWithReason(ctx, networkextension.NEProviderStopReasonUserInitiated); err != nil {
		return fmt.Errorf("app proxy context stop: %w", err)
	}

	provider.CancelProxyWithError(foundation.NSErrorFromID(0))
	if !provider.HandleNewFlow(networkextension.NEAppProxyFlowFromID(0)) {
		return errors.New("app proxy handle flow returned false")
	}
	if !provider.HandleNewUDPFlowInitialRemoteFlowEndpoint(networkextension.NEAppProxyUDPFlowFromID(0), network.NWEndpoint{}) {
		return errors.New("app proxy handle udp flow returned false")
	}
	return nil
}

func checkSelectors(id objc.ID, selectors []string) error {
	for _, sel := range selectors {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("callback object does not respond to %s", sel)
		}
	}
	return nil
}

func checkCounts(counts *providerCounts) error {
	checks := []struct {
		name string
		got  uint64
		want uint64
	}{
		{"packet starts", atomic.LoadUint64(&counts.packetStart), 2},
		{"packet stops", atomic.LoadUint64(&counts.packetStop), 2},
		{"packet cancels", atomic.LoadUint64(&counts.packetCancel), 1},
		{"packet flows", atomic.LoadUint64(&counts.packetFlow), 1},
		{"packet interfaces", atomic.LoadUint64(&counts.packetInterface), 1},
		{"app starts", atomic.LoadUint64(&counts.appStart), 2},
		{"app stops", atomic.LoadUint64(&counts.appStop), 2},
		{"app cancels", atomic.LoadUint64(&counts.appCancel), 1},
		{"app flows", atomic.LoadUint64(&counts.appFlow), 1},
		{"app udp flows", atomic.LoadUint64(&counts.appUDPFlow), 1},
	}
	for _, check := range checks {
		if check.got != check.want {
			return fmt.Errorf("%s = %d, want %d", check.name, check.got, check.want)
		}
	}
	return nil
}

func checkProviderSubclassAvailability() error {
	if networkextension.GetNEEthernetTunnelProviderClass().Class() == 0 {
		return errors.New("ethernet tunnel provider class unavailable")
	}
	if networkextension.GetNETransparentProxyProviderClass().Class() == 0 {
		return errors.New("transparent proxy provider class unavailable")
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-provider-wrapper-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-provider-wrapper-demo: NetworkExtension provider wrapper smoke ok")
}
