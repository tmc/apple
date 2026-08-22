// Code generated from Apple documentation. DO NOT EDIT.

package vmnet

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VmnetInterfaceCompletionHandler handles completion with a primitive value.

// NewVmnetInterfaceCompletionHandlerBlock wraps a Go [VmnetInterfaceCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVmnetInterfaceCompletionHandlerBlock(handler VmnetInterfaceCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal Vmnet_return_t) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VmnetInterfaceEventCallback handles completion with primitive and object results.

// NewVmnetInterfaceEventCallbackBlock wraps a Go [VmnetInterfaceEventCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVmnetInterfaceEventCallbackBlock(handler VmnetInterfaceEventCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive Interface_event_t, extra0 objectivec.Object) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VmnetInterfaceGetIPPortForwardingRulesHandler handles completion with a primitive value.

// NewVmnetInterfaceGetIPPortForwardingRulesHandlerBlock wraps a Go [VmnetInterfaceGetIPPortForwardingRulesHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVmnetInterfaceGetIPPortForwardingRulesHandlerBlock(handler VmnetInterfaceGetIPPortForwardingRulesHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VmnetInterfaceGetPortForwardingRulesHandler handles completion with a primitive value.

// NewVmnetInterfaceGetPortForwardingRulesHandlerBlock wraps a Go [VmnetInterfaceGetPortForwardingRulesHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVmnetInterfaceGetPortForwardingRulesHandlerBlock(handler VmnetInterfaceGetPortForwardingRulesHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VmnetStartInterfaceCompletionHandler handles completion with primitive and object results.

// NewVmnetStartInterfaceCompletionHandlerBlock wraps a Go [VmnetStartInterfaceCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVmnetStartInterfaceCompletionHandlerBlock(handler VmnetStartInterfaceCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive Vmnet_return_t, extra0 objectivec.Object) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
