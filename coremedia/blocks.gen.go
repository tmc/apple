// Code generated from Apple documentation. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// CMBufferGetBooleanHandler handles a primitive value and returns a primitive value.

// NewCMBufferGetBooleanHandlerBlock wraps a Go [CMBufferGetBooleanHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMBufferGetBooleanHandlerBlock(handler CMBufferGetBooleanHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) byte {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CMBufferGetSizeHandler handles a primitive value and returns a primitive value.

// NewCMBufferGetSizeHandlerBlock wraps a Go [CMBufferGetSizeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMBufferGetSizeHandlerBlock(handler CMBufferGetSizeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) uint64 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CMBufferGetTimeHandler handles a primitive value and returns a primitive value.

// NewCMBufferGetTimeHandlerBlock wraps a Go [CMBufferGetTimeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMBufferGetTimeHandlerBlock(handler CMBufferGetTimeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) CMTime {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CMBufferQueueTriggerHandler handles A type alias for a trigger handler.

// NewCMBufferQueueTriggerHandlerBlock wraps a Go [CMBufferQueueTriggerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMBufferQueueTriggerHandlerBlock(handler CMBufferQueueTriggerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CMSampleBufferInvalidateHandler handles Client callback called by [CMSampleBufferInvalidate(_:)].

// NewCMSampleBufferInvalidateHandlerBlock wraps a Go [CMSampleBufferInvalidateHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMSampleBufferInvalidateHandlerBlock(handler CMSampleBufferInvalidateHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CMSampleBufferMakeDataReadyHandler handles A block the system calls to make the sample buffer ready for use.

// NewCMSampleBufferMakeDataReadyHandlerBlock wraps a Go [CMSampleBufferMakeDataReadyHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMSampleBufferMakeDataReadyHandlerBlock(handler CMSampleBufferMakeDataReadyHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
