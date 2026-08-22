// Code generated from Apple documentation. DO NOT EDIT.

package security

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AuthorizationAsyncCallback handles A block used as a callback for the asynchronous version of copying authorization rights.

// NewAuthorizationAsyncCallbackBlock wraps a Go [AuthorizationAsyncCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAuthorizationAsyncCallbackBlock(handler AuthorizationAsyncCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 *AuthorizationItemSet) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecKeyGeneratePairBlock handles A block called with the results of a call to [SecKeyGeneratePairAsync(_:_:_:)].

// NewSecKeyGeneratePairBlock wraps a Go [SecKeyGeneratePairBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecKeyGeneratePairBlock(handler SecKeyGeneratePairBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive string, extra0 string, extra1 corefoundation.CFErrorRef) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecMessageBlock handles A block that delivers messages during asynchronous operations.

// NewSecMessageBlock wraps a Go [SecMessageBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecMessageBlock(handler SecMessageBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 corefoundation.CFErrorRef, extra1 uint32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecProtocolChallenge is the signature for a completion handler block.

// SecProtocolChallengeComplete handles completion with a primitive value.

// NewSecProtocolChallengeCompleteBlock wraps a Go [SecProtocolChallengeComplete] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecProtocolChallengeCompleteBlock(handler SecProtocolChallengeComplete) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecProtocolKeyUpdate is the signature for a completion handler block.

// SecProtocolKeyUpdateComplete is the signature for a completion handler block.

// NewSecProtocolKeyUpdateCompleteBlock wraps a Go [SecProtocolKeyUpdateComplete] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecProtocolKeyUpdateCompleteBlock(handler SecProtocolKeyUpdateComplete) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// SecProtocolPreSharedKeySelection is the signature for a completion handler block.

// SecProtocolPreSharedKeySelectionComplete handles completion with a primitive value.

// NewSecProtocolPreSharedKeySelectionCompleteBlock wraps a Go [SecProtocolPreSharedKeySelectionComplete] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecProtocolPreSharedKeySelectionCompleteBlock(handler SecProtocolPreSharedKeySelectionComplete) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Object) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecProtocolVerify is the signature for a completion handler block.

// SecProtocolVerifyComplete handles completion with a primitive value.

// NewSecProtocolVerifyCompleteBlock wraps a Go [SecProtocolVerifyComplete] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecProtocolVerifyCompleteBlock(handler SecProtocolVerifyComplete) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecTransformDataBlock handles A block used to override the default data handling for a transform.

// NewSecTransformDataBlock wraps a Go [SecTransformDataBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecTransformDataBlock(handler SecTransformDataBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) unsafe.Pointer {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecTrustCallback handles A block called with the results of an asynchronous trust evaluation.

// NewSecTrustCallbackBlock wraps a Go [SecTrustCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecTrustCallbackBlock(handler SecTrustCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SecTrustRef, extra0 SecTrustResultType) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecTrustWithErrorCallback handles A block called with the results of an asynchronous trust evaluation.

// NewSecTrustWithErrorCallbackBlock wraps a Go [SecTrustWithErrorCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSecTrustWithErrorCallbackBlock(handler SecTrustWithErrorCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SecTrustRef, extra0 bool, extra1 corefoundation.CFErrorRef) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}
