// Code generated from Apple documentation. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NWBrowseResultEnumerateInterface handles A handler that enumerates the interfaces associated with a discovered service.

// NewNWBrowseResultEnumerateInterfaceBlock wraps a Go [NWBrowseResultEnumerateInterface] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWBrowseResultEnumerateInterfaceBlock(handler NWBrowseResultEnumerateInterface) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWBrowserBrowseResultsChangedHandler handles A handler that delivers updates about discovered services.

// NewNWBrowserBrowseResultsChangedHandlerBlock wraps a Go [NWBrowserBrowseResultsChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWBrowserBrowseResultsChangedHandlerBlock(handler NWBrowserBrowseResultsChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 objectivec.Object, extra1 bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWBrowserStateChangedHandler handles A handler that delivers browser state updates with associated errors.

// NewNWBrowserStateChangedHandlerBlock wraps a Go [NWBrowserStateChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWBrowserStateChangedHandlerBlock(handler NWBrowserStateChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NWBrowserState, extra0 NWError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionBooleanEventHandler handles A handler that receives Boolean state updates from a connection, such as viability and better path state.

// NewNWConnectionBooleanEventHandlerBlock wraps a Go [NWConnectionBooleanEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionBooleanEventHandlerBlock(handler NWConnectionBooleanEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionGroupNewConnectionHandler handles completion with a primitive value.

// NewNWConnectionGroupNewConnectionHandlerBlock wraps a Go [NWConnectionGroupNewConnectionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionGroupNewConnectionHandlerBlock(handler NWConnectionGroupNewConnectionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionGroupReceiveHandler handles A handler that receives inbound messages from members of the group.

// NewNWConnectionGroupReceiveHandlerBlock wraps a Go [NWConnectionGroupReceiveHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionGroupReceiveHandlerBlock(handler NWConnectionGroupReceiveHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 objectivec.Object, extra1 bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionGroupSendCompletion handles A completion to notify you when data has been processed and sent.

// NewNWConnectionGroupSendCompletionBlock wraps a Go [NWConnectionGroupSendCompletion] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionGroupSendCompletionBlock(handler NWConnectionGroupSendCompletion) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NWError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionGroupStateChangedHandler handles A handler that receives connection group state updates.

// NewNWConnectionGroupStateChangedHandlerBlock wraps a Go [NWConnectionGroupStateChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionGroupStateChangedHandlerBlock(handler NWConnectionGroupStateChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NWConnectionGroupState, extra0 NWError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionPathEventHandler handles A handler that delivers network path updates.

// NewNWConnectionPathEventHandlerBlock wraps a Go [NWConnectionPathEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionPathEventHandlerBlock(handler NWConnectionPathEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionReceiveCompletion handles A completion handler that indicates when content has been received by the connection, or that an error was encountered.

// NewNWConnectionReceiveCompletionBlock wraps a Go [NWConnectionReceiveCompletion] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionReceiveCompletionBlock(handler NWConnectionReceiveCompletion) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 objectivec.Object, extra1 bool, extra2 NWError) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionSendCompletion handles A completion handler that indicates when the connection has finished processing sent content.

// NewNWConnectionSendCompletionBlock wraps a Go [NWConnectionSendCompletion] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionSendCompletionBlock(handler NWConnectionSendCompletion) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NWError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWConnectionStateChangedHandler handles A handler that delivers connection state updates with associated errors.

// NewNWConnectionStateChangedHandlerBlock wraps a Go [NWConnectionStateChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWConnectionStateChangedHandlerBlock(handler NWConnectionStateChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NWConnectionState, extra0 NWError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWDataTransferReportCollectBlock handles A block that is delivered when a data transfer report is fully collected.

// NewNWDataTransferReportCollectBlock wraps a Go [NWDataTransferReportCollectBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWDataTransferReportCollectBlock(handler NWDataTransferReportCollectBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWEstablishmentReportAccessBlock handles A block that delivers a connection’s establishment report when it’s in the ready state.

// NewNWEstablishmentReportAccessBlock wraps a Go [NWEstablishmentReportAccessBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWEstablishmentReportAccessBlock(handler NWEstablishmentReportAccessBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWEthernetChannelReceiveHandler handles A handler that delivers inbound Ethernet frames.

// NewNWEthernetChannelReceiveHandlerBlock wraps a Go [NWEthernetChannelReceiveHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWEthernetChannelReceiveHandlerBlock(handler NWEthernetChannelReceiveHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 uint32, extra1 *uint8, extra2 *uint8) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWEthernetChannelSendCompletion handles A handler that indicates when an Ethernet frame has been sent, or if an error was encountered.

// NewNWEthernetChannelSendCompletionBlock wraps a Go [NWEthernetChannelSendCompletion] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWEthernetChannelSendCompletionBlock(handler NWEthernetChannelSendCompletion) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NWError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWEthernetChannelStateChangedHandler handles A handler that delivers Ethernet channel state updates with associated errors.

// NewNWEthernetChannelStateChangedHandlerBlock wraps a Go [NWEthernetChannelStateChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWEthernetChannelStateChangedHandlerBlock(handler NWEthernetChannelStateChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NWEthernetChannelState, extra0 NWError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerBlock handles A block to be invoked asynchronously on your framer protocol’s scheduling context.

// NewNWFramerBlock wraps a Go [NWFramerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerBlock(handler NWFramerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerCleanupHandler handles A handler that tells your protocol to clean up all allocations before being deallocated.

// NewNWFramerCleanupHandlerBlock wraps a Go [NWFramerCleanupHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerCleanupHandlerBlock(handler NWFramerCleanupHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerInputHandler handles A handler that notifies your protocol that new inbound data is available to parse.

// NewNWFramerInputHandlerBlock wraps a Go [NWFramerInputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerInputHandlerBlock(handler NWFramerInputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) uint64 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerMessageDisposeValue handles A handler that’s invoked when your custom value needs to be released due to a message being released or the value being replaced.

// NewNWFramerMessageDisposeValueBlock wraps a Go [NWFramerMessageDisposeValue] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerMessageDisposeValueBlock(handler NWFramerMessageDisposeValue) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal kernel.Pointer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerOutputHandler handles A handler that notifies your protocol about a new outbound message.

// NewNWFramerOutputHandlerBlock wraps a Go [NWFramerOutputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerOutputHandlerBlock(handler NWFramerOutputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 objectivec.Object, extra1 uint32, extra2 bool) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerStartHandler handles A handler that represents the entry point into your custom protocol.

// NewNWFramerStartHandlerBlock wraps a Go [NWFramerStartHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerStartHandlerBlock(handler NWFramerStartHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) NWFramerStartResult {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerStopHandler handles A handler that requests that your protocol send any final messages to close the connection.

// NewNWFramerStopHandlerBlock wraps a Go [NWFramerStopHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerStopHandlerBlock(handler NWFramerStopHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWFramerWakeupHandler handles A handler that delivers a scheduled wakeup event.

// NewNWFramerWakeupHandlerBlock wraps a Go [NWFramerWakeupHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWFramerWakeupHandlerBlock(handler NWFramerWakeupHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWGroupDescriptorEnumerateEndpointsBlock handles A handler that lists all endpoints added to the group descriptor.

// NewNWGroupDescriptorEnumerateEndpointsBlock wraps a Go [NWGroupDescriptorEnumerateEndpointsBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWGroupDescriptorEnumerateEndpointsBlock(handler NWGroupDescriptorEnumerateEndpointsBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWListenerAdvertisedEndpointChangedHandler handles A handler that indicates changes to the service endpoints being advertised as they are added and removed.

// NewNWListenerAdvertisedEndpointChangedHandlerBlock wraps a Go [NWListenerAdvertisedEndpointChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWListenerAdvertisedEndpointChangedHandlerBlock(handler NWListenerAdvertisedEndpointChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.Object, extra0 bool) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWListenerNewConnectionGroupHandler handles completion with a primitive value.

// NewNWListenerNewConnectionGroupHandlerBlock wraps a Go [NWListenerNewConnectionGroupHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWListenerNewConnectionGroupHandlerBlock(handler NWListenerNewConnectionGroupHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWListenerNewConnectionHandler handles A handler that delivers inbound connections.

// NewNWListenerNewConnectionHandlerBlock wraps a Go [NWListenerNewConnectionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWListenerNewConnectionHandlerBlock(handler NWListenerNewConnectionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWListenerStateChangedHandler handles A handler that delivers listener state updates with associated errors.

// NewNWListenerStateChangedHandlerBlock wraps a Go [NWListenerStateChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWListenerStateChangedHandlerBlock(handler NWListenerStateChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NWListenerState, extra0 NWError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWParametersConfigureProtocolBlock handles A block to configure protocol options during the creation of a parameters object.

// NewNWParametersConfigureProtocolBlock wraps a Go [NWParametersConfigureProtocolBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWParametersConfigureProtocolBlock(handler NWParametersConfigureProtocolBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWParametersIterateInterfaceTypesBlock handles A block that allows inspection of a list of interface types.

// NewNWParametersIterateInterfaceTypesBlock wraps a Go [NWParametersIterateInterfaceTypesBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWParametersIterateInterfaceTypesBlock(handler NWParametersIterateInterfaceTypesBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NWInterfaceType) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWParametersIterateInterfacesBlock handles A block that allows inspection of a list of interfaces.

// NewNWParametersIterateInterfacesBlock wraps a Go [NWParametersIterateInterfacesBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWParametersIterateInterfacesBlock(handler NWParametersIterateInterfacesBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWPathEnumerateGatewaysBlock handles A block that enumerates the gateways configured on the interfaces available to a path.

// NewNWPathEnumerateGatewaysBlock wraps a Go [NWPathEnumerateGatewaysBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWPathEnumerateGatewaysBlock(handler NWPathEnumerateGatewaysBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWPathEnumerateInterfacesBlock handles A block that enumerates the interfaces available to a path.

// NewNWPathEnumerateInterfacesBlock wraps a Go [NWPathEnumerateInterfacesBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWPathEnumerateInterfacesBlock(handler NWPathEnumerateInterfacesBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWPathMonitorCancelHandler handles A handler that indicates when a monitor has been cancelled.

// NewNWPathMonitorCancelHandlerBlock wraps a Go [NWPathMonitorCancelHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWPathMonitorCancelHandlerBlock(handler NWPathMonitorCancelHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// NWPathMonitorUpdateHandler handles A handler that delivers network path updates.

// NewNWPathMonitorUpdateHandlerBlock wraps a Go [NWPathMonitorUpdateHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWPathMonitorUpdateHandlerBlock(handler NWPathMonitorUpdateHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWProtocolStackIterateProtocolsBlock handles A block that allows you to inspect or modify a single protocol’s options.

// NewNWProtocolStackIterateProtocolsBlock wraps a Go [NWProtocolStackIterateProtocolsBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWProtocolStackIterateProtocolsBlock(handler NWProtocolStackIterateProtocolsBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWProxyDomainEnumerator handles completion with a primitive value.

// NewNWProxyDomainEnumeratorBlock wraps a Go [NWProxyDomainEnumerator] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWProxyDomainEnumeratorBlock(handler NWProxyDomainEnumerator) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal string) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWReportResolutionReportEnumerator handles Iterates a list of resolution steps, as [nw_resolution_report_t] objects, performed during connection establishment, in order from first resolved to last resolved.

// NewNWReportResolutionReportEnumeratorBlock wraps a Go [NWReportResolutionReportEnumerator] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWReportResolutionReportEnumeratorBlock(handler NWReportResolutionReportEnumerator) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWWsClientRequestHandler handles A handler that delivers inbound client handshake requests.

// NewNWWsClientRequestHandlerBlock wraps a Go [NWWsClientRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWWsClientRequestHandlerBlock(handler NWWsClientRequestHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) objectivec.Object {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWWsPongHandler handles A handler that indicates that a Pong message has been received for a previously sent Ping message, or that an error was encountered.

// NewNWWsPongHandlerBlock wraps a Go [NWWsPongHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWWsPongHandlerBlock(handler NWWsPongHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NWError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NWWsSubprotocolEnumerator handles A block that enumerates the supported subprotocols in a WebSocket client request.

// NewNWWsSubprotocolEnumeratorBlock wraps a Go [NWWsSubprotocolEnumerator] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNWWsSubprotocolEnumeratorBlock(handler NWWsSubprotocolEnumerator) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal string) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
