// Code generated from Apple documentation. DO NOT EDIT.

package iokit

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// IOAsyncCallback handles standard callback function for asynchronous I/O requests with lots of extra arguments beyond a refcon and result code.

// NewIOAsyncCallbackBlock wraps a Go [IOAsyncCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOAsyncCallbackBlock(handler IOAsyncCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int, extra1 unsafe.Pointer, extra2 uint32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOAsyncCallback0 handles standard callback function for asynchronous I/O requests with no extra arguments beyond a refcon and result code.

// NewIOAsyncCallback0Block wraps a Go [IOAsyncCallback0] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOAsyncCallback0Block(handler IOAsyncCallback0) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOAsyncCallback1 handles standard callback function for asynchronous I/O requests with one extra argument beyond a refcon and result code.

// NewIOAsyncCallback1Block wraps a Go [IOAsyncCallback1] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOAsyncCallback1Block(handler IOAsyncCallback1) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int, extra1 unsafe.Pointer) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOAsyncCallback2 handles standard callback function for asynchronous I/O requests with two extra arguments beyond a refcon and result code.

// NewIOAsyncCallback2Block wraps a Go [IOAsyncCallback2] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOAsyncCallback2Block(handler IOAsyncCallback2) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int, extra1 unsafe.Pointer, extra2 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOServiceInterestCallback handles Callback function to be notified of changes in state of an IOService.

// NewIOServiceInterestCallbackBlock wraps a Go [IOServiceInterestCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOServiceInterestCallbackBlock(handler IOServiceInterestCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uintptr, extra1 uint32, extra2 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IOServiceMatchingCallback handles Callback function to be notified of IOService publication.

// NewIOServiceMatchingCallbackBlock wraps a Go [IOServiceMatchingCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIOServiceMatchingCallbackBlock(handler IOServiceMatchingCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uintptr) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
