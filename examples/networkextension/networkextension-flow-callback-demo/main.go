// Command networkextension-flow-callback-demo exercises NetworkExtension
// flow/session callback block signatures without provider-owned flow objects.
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
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

type flowCounts struct {
	open            uint64
	writeData       uint64
	readData        uint64
	readDatagrams   uint64
	writeDatagrams  uint64
	providerMessage uint64
	remoteEndpoint  uint64
	remoteFlow      uint64
	localEndpoint   uint64
	localFlow       uint64
	readPackets     uint64
	readPacketObjs  uint64
	writePackets    uint64
	writePacketObjs uint64
	startTunnel     uint64
	stopTunnel      uint64
	disconnectError uint64
	urlVerdict      uint64
	identity        uint64
}

func registerFlowCallbackClass(name string, invoker *objcbridge.BlockInvoker, counts *flowCounts) (objc.Class, error) {
	super := objc.GetClass("NSObject")
	if super == 0 {
		return 0, errors.New("lookup NSObject")
	}

	cls, err := objc.RegisterClass(name, super, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("register %s: %w", name, err)
	}

	response := foundation.NewDataWithBytesLength([]byte("response"))
	payload := foundation.NewDataWithBytesLength([]byte("flow"))
	datagramArray := objectivec.IObjectSliceToNSArray([]foundation.NSData{payload})
	protocolArray := objectivec.IObjectSliceToNSArray([]foundation.NSNumber{foundation.NewNumberWithInt(2)})
	packet := networkextension.NewPacketWithDataProtocolFamily(payload, 2)
	packetArray := objectivec.IObjectSliceToNSArray([]networkextension.NEPacket{packet})
	methods := []objc.MethodDef{
		{
			Cmd: objc.RegisterName("openWithLocalFlowEndpoint:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, endpoint objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.open, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("writeData:withCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, data objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.writeData, 1)
				if data == 0 {
					panic("nil write data")
				}
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("readDataWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.readData, 1)
				if err := invoker.ObjectError(completion, payload.GetID(), 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("readDatagramsAndFlowEndpointsWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.readDatagrams, 1)
				if err := invoker.ObjectObjectError(completion, datagramArray, 0, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("writeDatagrams:sentByFlowEndpoints:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, datagrams objc.ID, endpoints objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.writeDatagrams, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("sendProviderMessage:returnError:responseHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, message objc.ID, errp objc.ID, completion objc.ID) bool {
				atomic.AddUint64(&counts.providerMessage, 1)
				if message == 0 {
					panic("nil provider message")
				}
				if err := invoker.Object(completion, response.GetID()); err != nil {
					panic(err)
				}
				return true
			},
		},
		{
			Cmd: objc.RegisterName("remoteEndpoint"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.remoteEndpoint, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("remoteFlowEndpoint"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.remoteFlow, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("localEndpoint"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.localEndpoint, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("localFlowEndpoint"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.ID {
				atomic.AddUint64(&counts.localFlow, 1)
				return 0
			},
		},
		{
			Cmd: objc.RegisterName("readPacketObjectsWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.readPacketObjs, 1)
				if err := invoker.Object(completion, packetArray); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("writePacketObjects:"),
			Fn: func(self objc.ID, _cmd objc.SEL, packets objc.ID) bool {
				atomic.AddUint64(&counts.writePacketObjs, 1)
				if packets == 0 {
					panic("nil packet objects")
				}
				return true
			},
		},
		{
			Cmd: objc.RegisterName("readPacketsWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.readPackets, 1)
				if err := invoker.ObjectObject(completion, datagramArray, protocolArray); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("writePackets:withProtocols:"),
			Fn: func(self objc.ID, _cmd objc.SEL, packets objc.ID, protocols objc.ID) bool {
				atomic.AddUint64(&counts.writePackets, 1)
				if packets == 0 || protocols == 0 {
					panic("nil packet arrays")
				}
				return true
			},
		},
		{
			Cmd: objc.RegisterName("startTunnelWithOptions:andReturnError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, options objc.ID, errp uintptr) bool {
				atomic.AddUint64(&counts.startTunnel, 1)
				return true
			},
		},
		{
			Cmd: objc.RegisterName("stopTunnel"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				atomic.AddUint64(&counts.stopTunnel, 1)
			},
		},
		{
			Cmd: objc.RegisterName("fetchLastDisconnectErrorWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.disconnectError, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("verdictForURL:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, url objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.urlVerdict, 1)
				if url == 0 {
					panic("nil URL")
				}
				if err := invoker.Uint(completion, uint(networkextension.NEURLFilterVerdictDeny)); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("provideIdentityForConnection:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connection objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.identity, 1)
				if err := invoker.ObjectError(completion, 0, 0); err != nil {
					panic(err)
				}
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func runSmoke() error {
	var counts flowCounts
	name := fmt.Sprintf("GoNetworkExtensionFlowCallbacks_%d_%d", os.Getpid(), time.Now().UnixNano())
	cls, err := registerFlowCallbackClass(name, objcbridge.NewBlockInvoker(), &counts)
	if err != nil {
		return err
	}

	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("new"))
	if id == 0 {
		return fmt.Errorf("instantiate %s", name)
	}
	defer objc.Send[struct{}](id, objc.Sel("release"))

	for _, sel := range []string{
		"openWithLocalFlowEndpoint:completionHandler:",
		"writeData:withCompletionHandler:",
		"readDataWithCompletionHandler:",
		"readDatagramsAndFlowEndpointsWithCompletionHandler:",
		"writeDatagrams:sentByFlowEndpoints:completionHandler:",
		"sendProviderMessage:returnError:responseHandler:",
		"remoteEndpoint",
		"remoteFlowEndpoint",
		"localEndpoint",
		"localFlowEndpoint",
		"readPacketObjectsWithCompletionHandler:",
		"writePacketObjects:",
		"readPacketsWithCompletionHandler:",
		"writePackets:withProtocols:",
		"startTunnelWithOptions:andReturnError:",
		"stopTunnel",
		"fetchLastDisconnectErrorWithCompletionHandler:",
		"verdictForURL:completionHandler:",
		"provideIdentityForConnection:completionHandler:",
	} {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("callback object does not respond to %s", sel)
		}
	}

	if err := callErrorSelector(id, "openWithLocalFlowEndpoint:completionHandler:", func(block objc.ID) {
		objc.Send[struct{}](id, objc.Sel("openWithLocalFlowEndpoint:completionHandler:"), objc.ID(0), block)
	}); err != nil {
		return err
	}
	if err := callWriteData(id); err != nil {
		return err
	}
	if err := callReadData(id); err != nil {
		return err
	}
	if err := callReadDatagrams(id); err != nil {
		return err
	}
	if err := callWriteDatagrams(id); err != nil {
		return err
	}
	if err := callProviderMessage(id); err != nil {
		return err
	}
	if err := callErrorSelector(id, "fetchLastDisconnectErrorWithCompletionHandler:", nil); err != nil {
		return err
	}
	if err := callURLVerdict(id); err != nil {
		return err
	}
	if err := callIdentity(id); err != nil {
		return err
	}
	if err := callGeneratedWrappers(id); err != nil {
		return err
	}

	checks := []struct {
		name string
		got  uint64
		want uint64
	}{
		{"open", atomic.LoadUint64(&counts.open), 3},
		{"write data", atomic.LoadUint64(&counts.writeData), 3},
		{"read data", atomic.LoadUint64(&counts.readData), 3},
		{"read datagrams", atomic.LoadUint64(&counts.readDatagrams), 2},
		{"write datagrams", atomic.LoadUint64(&counts.writeDatagrams), 2},
		{"provider message", atomic.LoadUint64(&counts.providerMessage), 3},
		{"remote endpoint", atomic.LoadUint64(&counts.remoteEndpoint), 1},
		{"remote flow endpoint", atomic.LoadUint64(&counts.remoteFlow), 1},
		{"local endpoint", atomic.LoadUint64(&counts.localEndpoint), 1},
		{"local flow endpoint", atomic.LoadUint64(&counts.localFlow), 1},
		{"read packet objects", atomic.LoadUint64(&counts.readPacketObjs), 1},
		{"write packet objects", atomic.LoadUint64(&counts.writePacketObjs), 1},
		{"read packets", atomic.LoadUint64(&counts.readPackets), 1},
		{"write packets", atomic.LoadUint64(&counts.writePackets), 1},
		{"start tunnel", atomic.LoadUint64(&counts.startTunnel), 1},
		{"stop tunnel", atomic.LoadUint64(&counts.stopTunnel), 1},
		{"disconnect error", atomic.LoadUint64(&counts.disconnectError), 1},
		{"URL verdict", atomic.LoadUint64(&counts.urlVerdict), 1},
		{"identity", atomic.LoadUint64(&counts.identity), 1},
	}
	for _, check := range checks {
		if check.got != check.want {
			return fmt.Errorf("%s callbacks = %d, want %d", check.name, check.got, check.want)
		}
	}
	return nil
}

func callErrorSelector(id objc.ID, sel string, send func(block objc.ID)) error {
	var completed uint64
	block, cleanup := networkextension.NewErrorBlock(func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	if send != nil {
		send(block)
	} else {
		objc.Send[struct{}](id, objc.Sel(sel), block)
	}
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", sel, got)
	}
	return nil
}

func callWriteData(id objc.ID) error {
	data := foundation.NewDataWithBytesLength([]byte("write"))
	return callErrorSelector(id, "writeData:withCompletionHandler:", func(block objc.ID) {
		objc.Send[struct{}](id, objc.Sel("writeData:withCompletionHandler:"), data, block)
	})
}

func callReadData(id objc.ID) error {
	var completed uint64
	block, cleanup := networkextension.NewDataErrorBlock(func(data *foundation.NSData, err error) {
		if err != nil {
			panic(err)
		}
		if data == nil {
			panic("nil read data")
		}
		if got := data.Length(); got != 4 {
			panic(fmt.Sprintf("read data length = %d, want 4", got))
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	objc.Send[struct{}](id, objc.Sel("readDataWithCompletionHandler:"), block)
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("read data completions = %d, want 1", got)
	}
	return nil
}

func callReadDatagrams(id objc.ID) error {
	var completed uint64
	block, cleanup := networkextension.NewNSDataArrayObjectArrayErrorBlock(func(datagrams *[]foundation.NSData, endpoints *[]objectivec.Object, err error) {
		if err != nil {
			panic(err)
		}
		if datagrams == nil || len(*datagrams) != 1 {
			panic("unexpected datagrams")
		}
		if endpoints != nil {
			panic("unexpected flow endpoints")
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	objc.Send[struct{}](id, objc.Sel("readDatagramsAndFlowEndpointsWithCompletionHandler:"), block)
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("read datagrams completions = %d, want 1", got)
	}
	return nil
}

func callWriteDatagrams(id objc.ID) error {
	return callErrorSelector(id, "writeDatagrams:sentByFlowEndpoints:completionHandler:", func(block objc.ID) {
		objc.Send[struct{}](id, objc.Sel("writeDatagrams:sentByFlowEndpoints:completionHandler:"), objc.ID(0), objc.ID(0), block)
	})
}

func callProviderMessage(id objc.ID) error {
	message := foundation.NewDataWithBytesLength([]byte("message"))
	var completed uint64
	block, cleanup := networkextension.NewDataBlock(func(data *foundation.NSData) {
		if data == nil {
			panic("nil provider response")
		}
		if got := data.Length(); got != 8 {
			panic(fmt.Sprintf("provider response length = %d, want 8", got))
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	ok := objc.Send[bool](id, objc.Sel("sendProviderMessage:returnError:responseHandler:"), message, objc.ID(0), block)
	if !ok {
		return errors.New("provider message returned false")
	}
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("provider message completions = %d, want 1", got)
	}
	return nil
}

func callURLVerdict(id objc.ID) error {
	var completed uint64
	block, cleanup := networkextension.NewNEURLFilterVerdictBlock(func(verdict networkextension.NEURLFilterVerdict) {
		if verdict != networkextension.NEURLFilterVerdictDeny {
			panic(fmt.Sprintf("URL verdict = %v, want %v", verdict, networkextension.NEURLFilterVerdictDeny))
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	url := foundation.NewURLWithString("https://example.invalid/")
	objc.Send[struct{}](id, objc.Sel("verdictForURL:completionHandler:"), url, block)
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("URL verdict completions = %d, want 1", got)
	}
	return nil
}

func callIdentity(id objc.ID) error {
	var completed uint64
	block, cleanup := networkextension.NewSecIdentityRefArrayBlock(func(identity security.SecIdentity, certificateChain *foundation.NSArray) {
		if identity.GetID() != 0 {
			panic("unexpected identity")
		}
		if certificateChain != nil {
			panic("unexpected certificate chain")
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	objc.Send[struct{}](id, objc.Sel("provideIdentityForConnection:completionHandler:"), objc.ID(0), block)
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("identity completions = %d, want 1", got)
	}
	return nil
}

func callGeneratedWrappers(id objc.ID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	flow := networkextension.NEAppProxyFlowFromID(id)
	if err := expectErrorCompletion("generated open", func(handler networkextension.ErrorHandler) {
		flow.OpenWithLocalFlowEndpointCompletionHandler(network.NWEndpointFromID(0), handler)
	}); err != nil {
		return err
	}
	if err := flow.OpenWithLocalFlowEndpoint(ctx, network.NWEndpointFromID(0)); err != nil {
		return fmt.Errorf("generated open sync: %w", err)
	}

	data := foundation.NewDataWithBytesLength([]byte("write"))
	tcp := networkextension.NEAppProxyTCPFlowFromID(id)
	if err := expectErrorCompletion("generated tcp write", func(handler networkextension.ErrorHandler) {
		tcp.WriteDataWithCompletionHandler(data, handler)
	}); err != nil {
		return err
	}
	if err := tcp.WriteData(ctx, data); err != nil {
		return fmt.Errorf("generated tcp write sync: %w", err)
	}
	if err := expectDataErrorCompletion("generated tcp read", func(handler networkextension.DataErrorHandler) {
		tcp.ReadDataWithCompletionHandler(handler)
	}); err != nil {
		return err
	}
	readData, err := tcp.ReadData(ctx)
	if err != nil {
		return fmt.Errorf("generated tcp read sync: %w", err)
	}
	if readData == nil || readData.Length() != 4 {
		return errors.New("generated tcp read sync returned unexpected data")
	}
	if endpoint := tcp.RemoteEndpoint(); endpoint.GetID() != 0 {
		return fmt.Errorf("generated tcp remote endpoint = %#x, want nil", uintptr(endpoint.GetID()))
	}
	if endpoint := tcp.RemoteFlowEndpoint(); endpoint.GetID() != 0 {
		return fmt.Errorf("generated tcp remote flow endpoint = %#x, want nil", uintptr(endpoint.GetID()))
	}

	udp := networkextension.NEAppProxyUDPFlowFromID(id)
	if err := expectDatagramsCompletion("generated udp read", func(handler networkextension.NSDataArrayObjectArrayErrorHandler) {
		udp.ReadDatagramsAndFlowEndpointsWithCompletionHandler(handler)
	}); err != nil {
		return err
	}
	if err := expectErrorCompletion("generated udp write", func(handler networkextension.ErrorHandler) {
		udp.WriteDatagramsSentByFlowEndpointsCompletionHandler([]foundation.NSData{data}, nil, handler)
	}); err != nil {
		return err
	}
	if endpoint := udp.LocalEndpoint(); endpoint.GetID() != 0 {
		return fmt.Errorf("generated udp local endpoint = %#x, want nil", uintptr(endpoint.GetID()))
	}
	if endpoint := udp.LocalFlowEndpoint(); endpoint.GetID() != 0 {
		return fmt.Errorf("generated udp local flow endpoint = %#x, want nil", uintptr(endpoint.GetID()))
	}

	packetFlow := networkextension.NEPacketTunnelFlowFromID(id)
	if err := expectPacketObjectsCompletion("generated packet object read", func(handler networkextension.NEPacketArrayHandler) {
		packetFlow.ReadPacketObjectsWithCompletionHandler(handler)
	}); err != nil {
		return err
	}
	packet := networkextension.NewPacketWithDataProtocolFamily(data, 2)
	if !packetFlow.WritePacketObjects([]networkextension.NEPacket{packet}) {
		return errors.New("generated packet object write returned false")
	}
	if err := expectPacketsCompletion("generated packet read", func(handler networkextension.NSDataArrayNSNumberArrayHandler) {
		packetFlow.ReadPacketsWithCompletionHandler(handler)
	}); err != nil {
		return err
	}
	if !packetFlow.WritePacketsWithProtocols([]foundation.NSData{data}, []foundation.NSNumber{foundation.NewNumberWithInt(2)}) {
		return errors.New("generated packet write returned false")
	}

	session := networkextension.NETunnelProviderSessionFromID(id)
	message := foundation.NewDataWithBytesLength([]byte("message"))
	if err := expectDataCompletion("generated provider message", func(handler networkextension.DataHandler) {
		if !session.SendProviderMessageReturnErrorResponseHandler(message, foundation.NSErrorFromID(0), handler) {
			panic("provider message returned false")
		}
	}); err != nil {
		return err
	}
	response, err := session.SendProviderMessageReturnErrorResponseHandlerSync(ctx, message, foundation.NSErrorFromID(0))
	if err != nil {
		return fmt.Errorf("generated provider message sync: %w", err)
	}
	if response == nil || response.Length() != 8 {
		return errors.New("generated provider message sync returned unexpected data")
	}
	ok, err := session.StartTunnelWithOptionsAndReturnError(nil)
	if err != nil {
		return fmt.Errorf("generated tunnel start: %w", err)
	}
	if !ok {
		return errors.New("generated tunnel start returned false")
	}
	session.StopTunnel()
	return nil
}

func expectErrorCompletion(name string, call func(networkextension.ErrorHandler)) error {
	var completed uint64
	call(func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func expectDataCompletion(name string, call func(networkextension.DataHandler)) error {
	var completed uint64
	call(func(data *foundation.NSData) {
		if data == nil || data.Length() != 8 {
			panic("unexpected data")
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func expectDataErrorCompletion(name string, call func(networkextension.DataErrorHandler)) error {
	var completed uint64
	call(func(data *foundation.NSData, err error) {
		if err != nil {
			panic(err)
		}
		if data == nil || data.Length() != 4 {
			panic("unexpected data")
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func expectDatagramsCompletion(name string, call func(networkextension.NSDataArrayObjectArrayErrorHandler)) error {
	var completed uint64
	call(func(datagrams *[]foundation.NSData, endpoints *[]objectivec.Object, err error) {
		if err != nil {
			panic(err)
		}
		if datagrams == nil || len(*datagrams) != 1 {
			panic("unexpected datagrams")
		}
		if endpoints != nil {
			panic("unexpected endpoints")
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func expectPacketObjectsCompletion(name string, call func(networkextension.NEPacketArrayHandler)) error {
	var completed uint64
	call(func(packets *[]networkextension.NEPacket) {
		if packets == nil || len(*packets) != 1 {
			panic("unexpected packet objects")
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func expectPacketsCompletion(name string, call func(networkextension.NSDataArrayNSNumberArrayHandler)) error {
	var completed uint64
	call(func(packets *[]foundation.NSData, protocols *[]foundation.NSNumber) {
		if packets == nil || len(*packets) != 1 {
			panic("unexpected packets")
		}
		if protocols == nil || len(*protocols) != 1 {
			panic("unexpected protocols")
		}
		atomic.AddUint64(&completed, 1)
	})
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("%s completions = %d, want 1", name, got)
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-flow-callback-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-flow-callback-demo: NetworkExtension flow callback smoke ok")
}
