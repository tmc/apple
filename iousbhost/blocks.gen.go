// Code generated from Apple documentation. DO NOT EDIT.

package iousbhost

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// IOUSBHostCompletionHandler handles The completion handler for asynchronous control, bulk, and interrupt transfers.

// NewIOUSBHostCompletionHandlerBlock wraps a Go [IOUSBHostCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOUSBHostCompletionHandlerBlock(handler IOUSBHostCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 uint32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOUSBHostControllerInterfaceCommandHandler is the signature for a completion handler block.

// IOUSBHostControllerInterfaceDoorbellHandler handles completion with primitive and object results.

// NewIOUSBHostControllerInterfaceDoorbellHandlerBlock wraps a Go [IOUSBHostControllerInterfaceDoorbellHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOUSBHostControllerInterfaceDoorbellHandlerBlock(handler IOUSBHostControllerInterfaceDoorbellHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive IOUSBHostControllerInterface, extra0 *uint32, extra1 uint32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOUSBHostInterestHandler handles The callback that handles underlying service-state changes.

// NewIOUSBHostInterestHandlerBlock wraps a Go [IOUSBHostInterestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOUSBHostInterestHandlerBlock(handler IOUSBHostInterestHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive IOUSBHostObject, extra0 uint32, extra1 unsafe.Pointer) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOUSBHostIsochronousCompletionHandler handles A completion handler for asynchronous isochronous transfers.

// NewIOUSBHostIsochronousCompletionHandlerBlock wraps a Go [IOUSBHostIsochronousCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOUSBHostIsochronousCompletionHandlerBlock(handler IOUSBHostIsochronousCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 *IOUSBHostIsochronousFrame) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOUSBHostIsochronousTransactionCompletionHandler handles completion with primitive and object results.

// NewIOUSBHostIsochronousTransactionCompletionHandlerBlock wraps a Go [IOUSBHostIsochronousTransactionCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOUSBHostIsochronousTransactionCompletionHandlerBlock(handler IOUSBHostIsochronousTransactionCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 *IOUSBHostIsochronousTransaction) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// Int32IOUSBHostIsochronousTransactionHandler handles completion with primitive and object results.
//
// Used by:
//   - [IOUSBHostPipe.EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler]
type Int32IOUSBHostIsochronousTransactionHandler = func(int32, *IOUSBHostIsochronousTransaction)

// NewInt32IOUSBHostIsochronousTransactionBlock wraps a Go [Int32IOUSBHostIsochronousTransactionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [IOUSBHostPipe.EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler]
func NewInt32IOUSBHostIsochronousTransactionBlock(handler Int32IOUSBHostIsochronousTransactionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 *IOUSBHostIsochronousTransaction) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// Int32Uint32Handler handles An IOUSBHostCompletionHandler that runs when the request completes, or times out after the call returns successfully.
//
// Used by:
//   - [IOUSBHostPipe.EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler]
//   - [IOUSBHostStream.EnqueueIORequestWithDataErrorCompletionHandler]
type Int32Uint32Handler = func(int32, uint32)

// NewInt32Uint32Block wraps a Go [Int32Uint32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [IOUSBHostPipe.EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler]
//   - [IOUSBHostStream.EnqueueIORequestWithDataErrorCompletionHandler]
func NewInt32Uint32Block(handler Int32Uint32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 uint32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
