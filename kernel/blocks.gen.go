// Code generated from Apple documentation. DO NOT EDIT.

package kernel

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// IODataQueueClientDequeueEntryBlock handles completion with primitive and object results.

// NewIODataQueueClientDequeueEntryBlock wraps a Go [IODataQueueClientDequeueEntryBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIODataQueueClientDequeueEntryBlock(handler IODataQueueClientDequeueEntryBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uintptr) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IODataQueueClientEnqueueEntryBlock handles completion with primitive and object results.

// NewIODataQueueClientEnqueueEntryBlock wraps a Go [IODataQueueClientEnqueueEntryBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIODataQueueClientEnqueueEntryBlock(handler IODataQueueClientEnqueueEntryBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uintptr) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IODispatchBlock is the signature for a completion handler block.

// NewIODispatchBlock wraps a Go [IODispatchBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIODispatchBlock(handler IODispatchBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// IODispatchQueueCancelHandler is the signature for a completion handler block.

// NewIODispatchQueueCancelHandlerBlock wraps a Go [IODispatchQueueCancelHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIODispatchQueueCancelHandlerBlock(handler IODispatchQueueCancelHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// IODispatchSourceCancelHandler is the signature for a completion handler block.

// NewIODispatchSourceCancelHandlerBlock wraps a Go [IODispatchSourceCancelHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIODispatchSourceCancelHandlerBlock(handler IODispatchSourceCancelHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// IOInterruptActionBlock handles completion with primitive and object results.

// NewIOInterruptActionBlock wraps a Go [IOInterruptActionBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOInterruptActionBlock(handler IOInterruptActionBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *IOService, extra0 int32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOServiceApplierBlock handles completion with a primitive value.

// NewIOServiceApplierBlock wraps a Go [IOServiceApplierBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOServiceApplierBlock(handler IOServiceApplierBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *IOService) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOServiceNotificationBlock handles completion with primitive and object results.

// NewIOServiceNotificationBlock wraps a Go [IOServiceNotificationBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOServiceNotificationBlock(handler IOServiceNotificationBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint64, extra0 *IOService, extra1 uint64) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// OSActionAbortedHandler is the signature for a completion handler block.

// NewOSActionAbortedHandlerBlock wraps a Go [OSActionAbortedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSActionAbortedHandlerBlock(handler OSActionAbortedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// OSActionCancelHandler is the signature for a completion handler block.

// NewOSActionCancelHandlerBlock wraps a Go [OSActionCancelHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSActionCancelHandlerBlock(handler OSActionCancelHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// OSBlock is the signature for a completion handler block.

// NewOSBlock wraps a Go [OSBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSBlock(handler OSBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// OSObjectApplierBlock handles completion with a primitive value.

// NewOSObjectApplierBlock wraps a Go [OSObjectApplierBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSObjectApplierBlock(handler OSObjectApplierBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *OSObject) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// OSSerializerBlock handles a primitive value and returns a primitive value.

// NewOSSerializerBlock wraps a Go [OSSerializerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSSerializerBlock(handler OSSerializerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *OSSerialize) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
