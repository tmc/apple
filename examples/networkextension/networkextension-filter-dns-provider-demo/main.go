// Command networkextension-filter-dns-provider-demo exercises DNS proxy and
// content-filter provider callbacks without installing an extension.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/network"
	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

type callbackCounts struct {
	dnsStart      uint64
	dnsStop       uint64
	dnsCancel     uint64
	dnsFlow       uint64
	dnsUDPFlow    uint64
	dnsSettings   uint64
	filterStart   uint64
	filterStop    uint64
	filterReport  uint64
	filterConfig  uint64
	newFlow       uint64
	inboundData   uint64
	outboundData  uint64
	inboundDone   uint64
	outboundDone  uint64
	applySettings uint64
	resumeFlow    uint64
	updateFlow    uint64
	packetSet     uint64
	packetGet     uint64
	packetBlock   uint64
	packetDelay   uint64
	packetAllow   uint64
}

type callbackObjects struct {
	filterConfig   networkextension.NEFilterProviderConfiguration
	newFlowVerdict networkextension.NEFilterNewFlowVerdict
	dataVerdict    networkextension.NEFilterDataVerdict
	packetBlock    objc.Block
}

func registerDNSProxyClass(name string, invoker *objcbridge.BlockInvoker, counts *callbackCounts) (objc.Class, error) {
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
				atomic.AddUint64(&counts.dnsStart, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("stopProxyWithReason:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, reason networkextension.NEProviderStopReason, completion objc.ID) {
				atomic.AddUint64(&counts.dnsStop, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("cancelProxyWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, err objc.ID) {
				atomic.AddUint64(&counts.dnsCancel, 1)
			},
		},
		{
			Cmd: objc.RegisterName("handleNewFlow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID) bool {
				atomic.AddUint64(&counts.dnsFlow, 1)
				return true
			},
		},
		{
			Cmd: objc.RegisterName("handleNewUDPFlow:initialRemoteFlowEndpoint:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, endpoint objc.ID) bool {
				atomic.AddUint64(&counts.dnsUDPFlow, 1)
				return true
			},
		},
		{
			Cmd: objc.RegisterName("systemDNSSettings"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.dnsSettings, 1)
				return 0
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func registerFilterClass(name string, invoker *objcbridge.BlockInvoker, counts *callbackCounts, objects callbackObjects) (objc.Class, error) {
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
			Cmd: objc.RegisterName("startFilterWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.filterStart, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("stopFilterWithReason:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, reason networkextension.NEProviderStopReason, completion objc.ID) {
				atomic.AddUint64(&counts.filterStop, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("handleReport:"),
			Fn: func(self objc.ID, _cmd objc.SEL, report objc.ID) {
				atomic.AddUint64(&counts.filterReport, 1)
			},
		},
		{
			Cmd: objc.RegisterName("filterConfiguration"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.filterConfig, 1)
				return objects.filterConfig.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("handleNewFlow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID) objc.ID {
				atomic.AddUint64(&counts.newFlow, 1)
				return objects.newFlowVerdict.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("handleInboundDataFromFlow:readBytesStartOffset:readBytes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, offset uint, readBytes objc.ID) objc.ID {
				atomic.AddUint64(&counts.inboundData, 1)
				if readBytes == 0 || offset != 7 {
					panic("unexpected inbound data callback arguments")
				}
				return objects.dataVerdict.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("handleOutboundDataFromFlow:readBytesStartOffset:readBytes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, offset uint, readBytes objc.ID) objc.ID {
				atomic.AddUint64(&counts.outboundData, 1)
				if readBytes == 0 || offset != 11 {
					panic("unexpected outbound data callback arguments")
				}
				return objects.dataVerdict.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("handleInboundDataCompleteForFlow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID) objc.ID {
				atomic.AddUint64(&counts.inboundDone, 1)
				return objects.dataVerdict.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("handleOutboundDataCompleteForFlow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID) objc.ID {
				atomic.AddUint64(&counts.outboundDone, 1)
				return objects.dataVerdict.GetID()
			},
		},
		{
			Cmd: objc.RegisterName("applySettings:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, settings objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.applySettings, 1)
				if settings == 0 {
					panic("nil filter settings")
				}
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("resumeFlow:withVerdict:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, verdict objc.ID) {
				atomic.AddUint64(&counts.resumeFlow, 1)
				if verdict == 0 {
					panic("nil resume verdict")
				}
			},
		},
		{
			Cmd: objc.RegisterName("updateFlow:usingVerdict:forDirection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, flow objc.ID, verdict objc.ID, direction networkextension.NETrafficDirection) {
				atomic.AddUint64(&counts.updateFlow, 1)
				if verdict == 0 || direction != networkextension.NETrafficDirectionOutbound {
					panic("unexpected update verdict arguments")
				}
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func registerFilterPacketClass(name string, counts *callbackCounts, objects callbackObjects) (objc.Class, error) {
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
			Cmd: objc.RegisterName("packetHandler"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.packetGet, 1)
				return objc.ID(objects.packetBlock)
			},
		},
		{
			Cmd: objc.RegisterName("setPacketHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, block objc.ID) {
				atomic.AddUint64(&counts.packetSet, 1)
				if block == 0 {
					panic("nil packet handler")
				}
				packet := [...]byte{1, 2, 3, 4}
				objc.Block(block).Invoke(
					objc.ID(0),
					objc.ID(0),
					networkextension.NETrafficDirectionInbound,
					unsafe.Pointer(&packet[0]),
					uint32(len(packet)),
				)
			},
		},
		{
			Cmd: objc.RegisterName("delayCurrentPacket:"),
			Fn: func(self objc.ID, _cmd objc.SEL, context objc.ID) objc.ID {
				atomic.AddUint64(&counts.packetDelay, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("allowPacket:"),
			Fn: func(self objc.ID, _cmd objc.SEL, packet objc.ID) {
				atomic.AddUint64(&counts.packetAllow, 1)
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func runSmoke() error {
	var counts callbackCounts
	objects := callbackObjects{
		filterConfig:   networkextension.NewNEFilterProviderConfiguration(),
		newFlowVerdict: networkextension.GetNEFilterNewFlowVerdictClass().AllowVerdict(),
		dataVerdict:    networkextension.GetNEFilterDataVerdictClass().AllowVerdict(),
	}
	objects.filterConfig.SetServerAddress("filter-provider.example")
	objects.packetBlock = objc.NewBlock(func(b objc.Block, context objc.ID, iface objc.ID, direction networkextension.NETrafficDirection, packet unsafe.Pointer, packetLen uint32) networkextension.NEFilterPacketProviderVerdict {
		atomic.AddUint64(&counts.packetBlock, 1)
		if context != 0 || direction != networkextension.NETrafficDirectionInbound || packet == nil || packetLen != 4 {
			panic("unexpected packet handler arguments")
		}
		return networkextension.NEFilterPacketProviderVerdictAllow
	})
	defer objects.packetBlock.Release()

	dnsName := fmt.Sprintf("GoDNSProxyProviderCallbacks_%d_%d", os.Getpid(), time.Now().UnixNano())
	dnsClass, err := registerDNSProxyClass(dnsName, objcbridge.NewBlockInvoker(), &counts)
	if err != nil {
		return err
	}
	dnsID := objc.Send[objc.ID](objc.ID(dnsClass), objc.Sel("new"))
	if dnsID == 0 {
		return fmt.Errorf("instantiate %s", dnsName)
	}
	defer objc.Send[struct{}](dnsID, objc.Sel("release"))

	filterName := fmt.Sprintf("GoFilterProviderCallbacks_%d_%d", os.Getpid(), time.Now().UnixNano())
	filterClass, err := registerFilterClass(filterName, objcbridge.NewBlockInvoker(), &counts, objects)
	if err != nil {
		return err
	}
	filterID := objc.Send[objc.ID](objc.ID(filterClass), objc.Sel("new"))
	if filterID == 0 {
		return fmt.Errorf("instantiate %s", filterName)
	}
	defer objc.Send[struct{}](filterID, objc.Sel("release"))

	packetName := fmt.Sprintf("GoFilterPacketProviderCallbacks_%d_%d", os.Getpid(), time.Now().UnixNano())
	packetClass, err := registerFilterPacketClass(packetName, &counts, objects)
	if err != nil {
		return err
	}
	packetID := objc.Send[objc.ID](objc.ID(packetClass), objc.Sel("new"))
	if packetID == 0 {
		return fmt.Errorf("instantiate %s", packetName)
	}
	defer objc.Send[struct{}](packetID, objc.Sel("release"))

	if err := checkDNSSelectors(dnsID); err != nil {
		return err
	}
	if err := checkFilterSelectors(filterID); err != nil {
		return err
	}
	if err := checkFilterPacketSelectors(packetID); err != nil {
		return err
	}
	if err := exerciseDNSProxyProvider(dnsID); err != nil {
		return err
	}
	if err := exerciseFilterProvider(filterID); err != nil {
		return err
	}
	if err := exerciseFilterDataProvider(filterID); err != nil {
		return err
	}
	if err := exerciseFilterPacketProvider(packetID, objects.packetBlock); err != nil {
		return err
	}
	if err := checkFilterObjectDefaults(); err != nil {
		return err
	}
	return checkCounts(&counts)
}

func checkFilterObjectDefaults() error {
	flow := networkextension.NewNEFilterFlow()
	if flow.GetID() == 0 {
		return errors.New("create filter flow")
	}
	if flow.URL().GetID() != 0 || flow.Identifier().GetID() != 0 {
		return errors.New("new filter flow has URL or identifier")
	}
	if flow.SourceAppAuditToken().GetID() != 0 || flow.SourceProcessAuditToken().GetID() != 0 {
		return errors.New("new filter flow has audit token")
	}
	if got := flow.Direction(); got != networkextension.NETrafficDirectionAny {
		return fmt.Errorf("filter flow direction = %v, want any", got)
	}

	socketFlow := networkextension.NewNEFilterSocketFlow()
	if socketFlow.GetID() == 0 {
		return errors.New("create filter socket flow")
	}
	if socketFlow.RemoteEndpoint().GetID() != 0 || socketFlow.LocalEndpoint().GetID() != 0 {
		return errors.New("new filter socket flow has endpoint")
	}
	if socketFlow.RemoteFlowEndpoint().GetID() != 0 || socketFlow.LocalFlowEndpoint().GetID() != 0 {
		return errors.New("new filter socket flow has Network.framework endpoint")
	}
	if socketFlow.RemoteHostname() != "" ||
		socketFlow.SocketFamily() != 0 ||
		socketFlow.SocketType() != 0 ||
		socketFlow.SocketProtocol() != 0 {
		return errors.New("new filter socket flow has socket metadata")
	}

	report := networkextension.NewNEFilterReport()
	if report.GetID() == 0 {
		return errors.New("create filter report")
	}
	if report.Flow().GetID() != 0 ||
		report.Action() != networkextension.NEFilterActionInvalid ||
		report.Event() != 0 ||
		report.BytesInboundCount() != 0 ||
		report.BytesOutboundCount() != 0 {
		return errors.New("new filter report has nonzero fields")
	}

	context := networkextension.NewNEFilterPacketContext()
	if context.GetID() == 0 {
		return errors.New("create filter packet context")
	}
	return nil
}

func checkDNSSelectors(id objc.ID) error {
	for _, sel := range []string{
		"startProxyWithOptions:completionHandler:",
		"stopProxyWithReason:completionHandler:",
		"cancelProxyWithError:",
		"handleNewFlow:",
		"handleNewUDPFlow:initialRemoteFlowEndpoint:",
		"systemDNSSettings",
	} {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("dns callback object does not respond to %s", sel)
		}
	}
	return nil
}

func checkFilterSelectors(id objc.ID) error {
	for _, sel := range []string{
		"startFilterWithCompletionHandler:",
		"stopFilterWithReason:completionHandler:",
		"handleReport:",
		"filterConfiguration",
		"handleNewFlow:",
		"handleInboundDataFromFlow:readBytesStartOffset:readBytes:",
		"handleOutboundDataFromFlow:readBytesStartOffset:readBytes:",
		"handleInboundDataCompleteForFlow:",
		"handleOutboundDataCompleteForFlow:",
		"applySettings:completionHandler:",
		"resumeFlow:withVerdict:",
		"updateFlow:usingVerdict:forDirection:",
	} {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("filter callback object does not respond to %s", sel)
		}
	}
	return nil
}

func checkFilterPacketSelectors(id objc.ID) error {
	for _, sel := range []string{
		"packetHandler",
		"setPacketHandler:",
		"delayCurrentPacket:",
		"allowPacket:",
	} {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("filter packet callback object does not respond to %s", sel)
		}
	}
	return nil
}

func exerciseDNSProxyProvider(id objc.ID) error {
	provider := networkextension.NEDNSProxyProviderFromID(id)

	var started uint64
	provider.StartProxyWithOptionsCompletionHandler(nil, func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&started, 1)
	})
	if got := atomic.LoadUint64(&started); got != 1 {
		return fmt.Errorf("dns proxy direct starts = %d, want 1", got)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := provider.StartProxyWithOptions(ctx, nil); err != nil {
		return fmt.Errorf("dns proxy context start: %w", err)
	}

	var stopped uint64
	provider.StopProxyWithReasonCompletionHandler(networkextension.NEProviderStopReasonUserInitiated, func() {
		atomic.AddUint64(&stopped, 1)
	})
	if got := atomic.LoadUint64(&stopped); got != 1 {
		return fmt.Errorf("dns proxy direct stops = %d, want 1", got)
	}
	if err := provider.StopProxyWithReason(ctx, networkextension.NEProviderStopReasonUserInitiated); err != nil {
		return fmt.Errorf("dns proxy context stop: %w", err)
	}

	provider.CancelProxyWithError(foundation.NSErrorFromID(0))
	if !provider.HandleNewFlow(networkextension.NEAppProxyFlowFromID(0)) {
		return errors.New("dns proxy handle flow returned false")
	}
	if !provider.HandleNewUDPFlowInitialRemoteFlowEndpoint(networkextension.NEAppProxyUDPFlowFromID(0), network.NWEndpoint{}) {
		return errors.New("dns proxy handle udp flow returned false")
	}
	if settings := provider.SystemDNSSettings(); settings.GetID() != 0 {
		return fmt.Errorf("system dns settings id = %v, want nil", settings.GetID())
	}
	return nil
}

func exerciseFilterProvider(id objc.ID) error {
	provider := networkextension.NEFilterProviderFromID(id)

	var started uint64
	provider.StartFilterWithCompletionHandler(func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&started, 1)
	})
	if got := atomic.LoadUint64(&started); got != 1 {
		return fmt.Errorf("filter direct starts = %d, want 1", got)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := provider.StartFilter(ctx); err != nil {
		return fmt.Errorf("filter context start: %w", err)
	}

	var stopped uint64
	provider.StopFilterWithReasonCompletionHandler(networkextension.NEProviderStopReasonUserInitiated, func() {
		atomic.AddUint64(&stopped, 1)
	})
	if got := atomic.LoadUint64(&stopped); got != 1 {
		return fmt.Errorf("filter direct stops = %d, want 1", got)
	}
	if err := provider.StopFilterWithReason(ctx, networkextension.NEProviderStopReasonUserInitiated); err != nil {
		return fmt.Errorf("filter context stop: %w", err)
	}

	provider.HandleReport(networkextension.NEFilterReportFromID(0))
	if got := provider.FilterConfiguration().ServerAddress(); got != "filter-provider.example" {
		return fmt.Errorf("filter configuration server = %q", got)
	}
	return nil
}

func exerciseFilterDataProvider(id objc.ID) error {
	provider := networkextension.NEFilterDataProviderFromID(id)
	data := foundation.NewDataWithBytesLength([]byte("packet"))
	settings := newFilterSettings()
	flow := networkextension.NEFilterFlowFromID(0)
	socketFlow := networkextension.NEFilterSocketFlowFromID(0)
	verdict := networkextension.GetNEFilterDataVerdictClass().AllowVerdict()

	if got := provider.HandleNewFlow(flow); got.GetID() == 0 {
		return fmt.Errorf("filter data new-flow verdict is nil")
	}
	if got := provider.HandleInboundDataFromFlowReadBytesStartOffsetReadBytes(flow, 7, data); got.GetID() == 0 {
		return fmt.Errorf("filter inbound verdict is nil")
	}
	if got := provider.HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes(flow, 11, data); got.GetID() == 0 {
		return fmt.Errorf("filter outbound verdict is nil")
	}
	if got := provider.HandleInboundDataCompleteForFlow(flow); got.GetID() == 0 {
		return fmt.Errorf("filter inbound complete verdict is nil")
	}
	if got := provider.HandleOutboundDataCompleteForFlow(flow); got.GetID() == 0 {
		return fmt.Errorf("filter outbound complete verdict is nil")
	}

	var applied uint64
	provider.ApplySettingsCompletionHandler(settings, func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&applied, 1)
	})
	if got := atomic.LoadUint64(&applied); got != 1 {
		return fmt.Errorf("filter direct apply completions = %d, want 1", got)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := provider.ApplySettings(ctx, settings); err != nil {
		return fmt.Errorf("filter context apply: %w", err)
	}

	provider.ResumeFlowWithVerdict(flow, verdict)
	provider.UpdateFlowUsingVerdictForDirection(socketFlow, verdict, networkextension.NETrafficDirectionOutbound)
	return nil
}

func exerciseFilterPacketProvider(id objc.ID, block objc.Block) error {
	provider := networkextension.NEFilterPacketProviderFromID(id)

	// The generated block-typed property setter currently has a Go func
	// shape, so this example uses the public raw Objective-C block API for the
	// packetHandler property and the generated wrappers for real methods.
	objc.Send[struct{}](id, objc.Sel("setPacketHandler:"), objc.ID(block))
	got := objc.Send[objc.ID](id, objc.Sel("packetHandler"))
	if got != objc.ID(block) {
		return fmt.Errorf("packet handler id = %v, want %v", got, objc.ID(block))
	}
	packet := [...]byte{5, 6, 7, 8}
	objc.Block(got).Invoke(
		objc.ID(0),
		objc.ID(0),
		networkextension.NETrafficDirectionInbound,
		unsafe.Pointer(&packet[0]),
		uint32(len(packet)),
	)
	if got := provider.DelayCurrentPacket(networkextension.NEFilterPacketContextFromID(0)); got.GetID() != 0 {
		return fmt.Errorf("delayed packet id = %v, want nil", got.GetID())
	}
	provider.AllowPacket(networkextension.NEPacketFromID(0))
	return nil
}

func newFilterSettings() networkextension.NEFilterSettings {
	host := networkextension.NewNWHostEndpointWithHostnamePort("provider-filter.example", "443")
	rule := networkextension.NewNetworkRuleWithDestinationHostProtocol(host, networkextension.NENetworkRuleProtocolTCP)
	filterRule := networkextension.NewFilterRuleWithNetworkRuleAction(rule, networkextension.NEFilterActionFilterData)
	return networkextension.NewFilterSettingsWithRulesDefaultAction(
		[]networkextension.NEFilterRule{filterRule},
		networkextension.NEFilterActionAllow,
	)
}

func checkCounts(counts *callbackCounts) error {
	checks := []struct {
		name string
		got  uint64
		want uint64
	}{
		{"dns starts", atomic.LoadUint64(&counts.dnsStart), 2},
		{"dns stops", atomic.LoadUint64(&counts.dnsStop), 2},
		{"dns cancels", atomic.LoadUint64(&counts.dnsCancel), 1},
		{"dns flows", atomic.LoadUint64(&counts.dnsFlow), 1},
		{"dns udp flows", atomic.LoadUint64(&counts.dnsUDPFlow), 1},
		{"dns settings", atomic.LoadUint64(&counts.dnsSettings), 1},
		{"filter starts", atomic.LoadUint64(&counts.filterStart), 2},
		{"filter stops", atomic.LoadUint64(&counts.filterStop), 2},
		{"filter reports", atomic.LoadUint64(&counts.filterReport), 1},
		{"filter configs", atomic.LoadUint64(&counts.filterConfig), 1},
		{"new flow verdicts", atomic.LoadUint64(&counts.newFlow), 1},
		{"inbound data verdicts", atomic.LoadUint64(&counts.inboundData), 1},
		{"outbound data verdicts", atomic.LoadUint64(&counts.outboundData), 1},
		{"inbound complete verdicts", atomic.LoadUint64(&counts.inboundDone), 1},
		{"outbound complete verdicts", atomic.LoadUint64(&counts.outboundDone), 1},
		{"apply settings", atomic.LoadUint64(&counts.applySettings), 2},
		{"resume flow", atomic.LoadUint64(&counts.resumeFlow), 1},
		{"update flow", atomic.LoadUint64(&counts.updateFlow), 1},
		{"packet handler sets", atomic.LoadUint64(&counts.packetSet), 1},
		{"packet handler gets", atomic.LoadUint64(&counts.packetGet), 1},
		{"packet handler invokes", atomic.LoadUint64(&counts.packetBlock), 2},
		{"packet delays", atomic.LoadUint64(&counts.packetDelay), 1},
		{"packet allows", atomic.LoadUint64(&counts.packetAllow), 1},
	}
	for _, check := range checks {
		if check.got != check.want {
			return fmt.Errorf("%s = %d, want %d", check.name, check.got, check.want)
		}
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-filter-dns-provider-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-filter-dns-provider-demo: NetworkExtension filter/DNS provider smoke ok")
}
