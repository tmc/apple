// Code generated from Apple documentation. DO NOT EDIT.

package applicationservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
)

// ATSCubicClosePathProcPtr handles a primitive value and returns a primitive value.

// NewATSCubicClosePathProcPtrBlock wraps a Go [ATSCubicClosePathProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewATSCubicClosePathProcPtrBlock(handler ATSCubicClosePathProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int32 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// ATSNotificationCallback handles completion with primitive and object results.

// NewATSNotificationCallbackBlock wraps a Go [ATSNotificationCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewATSNotificationCallbackBlock(handler ATSNotificationCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive ATSFontNotificationInfoRef, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// ATSQuadraticClosePathProcPtr handles a primitive value and returns a primitive value.

// NewATSQuadraticClosePathProcPtrBlock wraps a Go [ATSQuadraticClosePathProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewATSQuadraticClosePathProcPtrBlock(handler ATSQuadraticClosePathProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int32 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// ATSQuadraticNewPathProcPtr handles a primitive value and returns a primitive value.

// NewATSQuadraticNewPathProcPtrBlock wraps a Go [ATSQuadraticNewPathProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewATSQuadraticNewPathProcPtrBlock(handler ATSQuadraticNewPathProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) int32 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AXObserverCallback handles completion with primitive and object results.

// NewAXObserverCallbackBlock wraps a Go [AXObserverCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAXObserverCallbackBlock(handler AXObserverCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive AXObserverRef, extra0 AXUIElementRef, extra1 corefoundation.CFStringRef, extra2 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AXObserverCallbackWithInfo handles completion with primitive and object results.

// NewAXObserverCallbackWithInfoBlock wraps a Go [AXObserverCallbackWithInfo] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAXObserverCallbackWithInfoBlock(handler AXObserverCallbackWithInfo) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive AXObserverRef, extra0 AXUIElementRef, extra1 corefoundation.CFStringRef, extra2 corefoundation.CFDictionaryRef, extra3 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// ColorComplementProcPtr handles a primitive value and returns a primitive value.

// NewColorComplementProcPtrBlock wraps a Go [ColorComplementProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewColorComplementProcPtrBlock(handler ColorComplementProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// DragGrayRgnProcPtr is the signature for a completion handler block.

// NewDragGrayRgnProcPtrBlock wraps a Go [DragGrayRgnProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewDragGrayRgnProcPtrBlock(handler DragGrayRgnProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// QDArcProcPtr handles completion with primitive and object results.

// NewQDArcProcPtrBlock wraps a Go [QDArcProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDArcProcPtrBlock(handler QDArcProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 unsafe.Pointer, extra1 int16, extra2 int16) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDBitsProcPtr handles completion with primitive and object results.

// NewQDBitsProcPtrBlock wraps a Go [QDBitsProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDBitsProcPtrBlock(handler QDBitsProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 unsafe.Pointer, extra1 unsafe.Pointer, extra2 int16, extra3 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDCommentProcPtr handles completion with primitive and object results.

// NewQDCommentProcPtrBlock wraps a Go [QDCommentProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDCommentProcPtrBlock(handler QDCommentProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int16, extra0 int16, extra1 kernel.Handle) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDGetPicProcPtr handles completion with primitive and object results.

// NewQDGetPicProcPtrBlock wraps a Go [QDGetPicProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDGetPicProcPtrBlock(handler QDGetPicProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int16) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDJShieldCursorProcPtr handles completion with primitive and object results.

// NewQDJShieldCursorProcPtrBlock wraps a Go [QDJShieldCursorProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDJShieldCursorProcPtrBlock(handler QDJShieldCursorProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int16, extra0 int16, extra1 int16, extra2 int16) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDLineProcPtr handles completion with a primitive value.

// NewQDLineProcPtrBlock wraps a Go [QDLineProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDLineProcPtrBlock(handler QDLineProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal Point) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDOpcodeProcPtr handles completion with primitive and object results.

// NewQDOpcodeProcPtrBlock wraps a Go [QDOpcodeProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDOpcodeProcPtrBlock(handler QDOpcodeProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 unsafe.Pointer, extra1 uint16, extra2 int16) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDOvalProcPtr handles completion with primitive and object results.

// NewQDOvalProcPtrBlock wraps a Go [QDOvalProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDOvalProcPtrBlock(handler QDOvalProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDPolyProcPtr handles completion with primitive and object results.

// NewQDPolyProcPtrBlock wraps a Go [QDPolyProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDPolyProcPtrBlock(handler QDPolyProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 PolyHandle) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDPutPicProcPtr handles completion with primitive and object results.

// NewQDPutPicProcPtrBlock wraps a Go [QDPutPicProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDPutPicProcPtrBlock(handler QDPutPicProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int16) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDRRectProcPtr handles completion with primitive and object results.

// NewQDRRectProcPtrBlock wraps a Go [QDRRectProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDRRectProcPtrBlock(handler QDRRectProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 unsafe.Pointer, extra1 int16, extra2 int16) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDRectProcPtr handles completion with primitive and object results.

// NewQDRectProcPtrBlock wraps a Go [QDRectProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDRectProcPtrBlock(handler QDRectProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDRgnProcPtr handles completion with primitive and object results.

// NewQDRgnProcPtrBlock wraps a Go [QDRgnProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDRgnProcPtrBlock(handler QDRgnProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GrafVerb, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// QDTextProcPtr handles completion with primitive and object results.

// NewQDTextProcPtrBlock wraps a Go [QDTextProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewQDTextProcPtrBlock(handler QDTextProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int16, extra0 unsafe.Pointer, extra1 Point, extra2 Point) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechDoneProcPtr handles Defines a pointer to a speech-done callback functionwhich is called when the Speech Synthesis Manager finishes speakinga buffer of text.

// NewSpeechDoneProcPtrBlock wraps a Go [SpeechDoneProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechDoneProcPtrBlock(handler SpeechDoneProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechErrorCFProcPtr handles Defines a pointer to an error callback function that handles syntax errors within commands embedded in a [CFString] object being processed by the Speech Synthesis Manager.

// NewSpeechErrorCFProcPtrBlock wraps a Go [SpeechErrorCFProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechErrorCFProcPtrBlock(handler SpeechErrorCFProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 corefoundation.CFErrorRef) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechErrorProcPtr handles Defines a pointer to an error callback functionthat handles syntax errors within commands embedded in a text bufferbeing processed by the Speech Synthesis Manager.

// NewSpeechErrorProcPtrBlock wraps a Go [SpeechErrorProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechErrorProcPtrBlock(handler SpeechErrorProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 int16, extra2 int32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechPhonemeProcPtr handles Defines a pointer to a phoneme callback functionthat is called by the Speech Synthesis Manager before it pronouncesa phoneme.

// NewSpeechPhonemeProcPtrBlock wraps a Go [SpeechPhonemeProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechPhonemeProcPtrBlock(handler SpeechPhonemeProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 int16) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechSyncProcPtr handles Defines a pointer to a synchronization callbackfunction that is called when the Speech Synthesis Manager encountersa synchronization command embedded in a text buffer.

// NewSpeechSyncProcPtrBlock wraps a Go [SpeechSyncProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechSyncProcPtrBlock(handler SpeechSyncProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 uint32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechTextDoneProcPtr handles Defines a pointer to a text-done callback functionthat is called when the Speech Synthesis Manager has finished processinga buffer of text.

// NewSpeechTextDoneProcPtrBlock wraps a Go [SpeechTextDoneProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechTextDoneProcPtrBlock(handler SpeechTextDoneProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 unsafe.Pointer, extra2 unsafe.Pointer, extra3 unsafe.Pointer) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechWordCFProcPtr handles Defines a pointer to a Core Foundation-based word callback function that is called by the Speech Synthesis Manager before it pronounces a word.

// NewSpeechWordCFProcPtrBlock wraps a Go [SpeechWordCFProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechWordCFProcPtrBlock(handler SpeechWordCFProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 corefoundation.CFStringRef, extra2 corefoundation.CFRange) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// SpeechWordProcPtr handles Defines a pointer to a word callback functionthat is called by the Speech Synthesis Manager before it pronouncesa word.

// NewSpeechWordProcPtrBlock wraps a Go [SpeechWordProcPtr] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSpeechWordProcPtrBlock(handler SpeechWordProcPtr) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *SpeechChannelRecord, extra0 uintptr, extra1 uint, extra2 uint16) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}
