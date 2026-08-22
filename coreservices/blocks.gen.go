// Code generated from Apple documentation. DO NOT EDIT.

package coreservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// AEDisposeExternalProcPtr handles Defines a pointer to a function the Apple Event Manager calls to dispose of a descriptor created by the [AECreateDescFromExternalPtr] function.

// NewAEDisposeExternalProcPtrBlock wraps a Go [AEDisposeExternalProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAEDisposeExternalProcPtrBlock(handler AEDisposeExternalProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int, extra1 uintptr) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// AERemoteProcessResolverCallback handles Defines a pointer to a function the Apple Event Manager calls when the asynchronous execution of a remote process resolver completes, either due to success or failure, after a call to the [AERemoteProcessResolverScheduleWithRunLoop] function.

// NewAERemoteProcessResolverCallbackBlock wraps a Go [AERemoteProcessResolverCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAERemoteProcessResolverCallbackBlock(handler AERemoteProcessResolverCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive AERemoteProcessResolverRef, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CSDiskSpaceRecoveryCallback handles completion with primitive and object results.

// NewCSDiskSpaceRecoveryCallbackBlock wraps a Go [CSDiskSpaceRecoveryCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCSDiskSpaceRecoveryCallbackBlock(handler CSDiskSpaceRecoveryCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive bool, extra0 uint64, extra1 corefoundation.CFErrorRef) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// CSIdentityQueryReceiveEventCallback handles completion with primitive and object results.

// NewCSIdentityQueryReceiveEventCallbackBlock wraps a Go [CSIdentityQueryReceiveEventCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCSIdentityQueryReceiveEventCallbackBlock(handler CSIdentityQueryReceiveEventCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive CSIdentityQueryRef, extra0 CSIdentityQueryEvent, extra1 corefoundation.CFArrayRef, extra2 corefoundation.CFErrorRef, extra3 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// CSIdentityStatusUpdatedCallback handles completion with primitive and object results.

// NewCSIdentityStatusUpdatedCallbackBlock wraps a Go [CSIdentityStatusUpdatedCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCSIdentityStatusUpdatedCallbackBlock(handler CSIdentityStatusUpdatedCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive CSIdentityRef, extra0 corefoundation.CFIndex, extra1 corefoundation.CFErrorRef, extra2 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// FSEventStreamCallback handles completion with primitive and object results.

// NewFSEventStreamCallbackBlock wraps a Go [FSEventStreamCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFSEventStreamCallbackBlock(handler FSEventStreamCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive ConstFSEventStreamRef, extra0 unsafe.Pointer, extra1 int32, extra2 unsafe.Pointer, extra3 unsafe.Pointer, extra4 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2, extra3, extra4)
	})
	return objc.ID(block), func() { block.Release() }
}

// LSSharedFileListChangedProcPtr handles completion with primitive and object results.

// NewLSSharedFileListChangedProcPtrBlock wraps a Go [LSSharedFileListChangedProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewLSSharedFileListChangedProcPtrBlock(handler LSSharedFileListChangedProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive LSSharedFileListRef, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// OSLDisposeTokenProcPtr handles Defines a pointer to a dispose token callback function.

// NewOSLDisposeTokenProcPtrBlock wraps a Go [OSLDisposeTokenProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSLDisposeTokenProcPtrBlock(handler OSLDisposeTokenProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int16 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// OSLGetErrDescProcPtr handles Defines a pointer to an error descriptor callback function.

// NewOSLGetErrDescProcPtrBlock wraps a Go [OSLGetErrDescProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSLGetErrDescProcPtrBlock(handler OSLGetErrDescProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int16 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
