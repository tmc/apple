// Code generated from Apple documentation. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
)

// BitmapImageRepPGDisplayCoord_tHandler is the signature for a completion handler block.
type BitmapImageRepPGDisplayCoord_tHandler = func(*appkit.NSBitmapImageRep, PGDisplayCoord_t)

// NewBitmapImageRepPGDisplayCoord_tBlock wraps a Go [BitmapImageRepPGDisplayCoord_tHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewBitmapImageRepPGDisplayCoord_tBlock(handler BitmapImageRepPGDisplayCoord_tHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 PGDisplayCoord_t) {
		var result *appkit.NSBitmapImageRep
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := appkit.NSBitmapImageRepFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolUint64Handler handles completion with primitive and object results.
type BoolUint64Handler = func(uint64, uint64, unsafe.Pointer) bool

// NewBoolUint64Block wraps a Go [BoolUint64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewBoolUint64Block(handler BoolUint64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint64, extra0 uint64, extra1 unsafe.Pointer) bool {
		return handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolstructPGTask_sUint32Uint64BoolPGPhysicalMemoryRange_sHandler is the signature for a completion handler block.
type BoolstructPGTask_sUint32Uint64BoolPGPhysicalMemoryRange_sHandler = func(*uintptr, uint32, uint64, bool, *PGPhysicalMemoryRange_s) bool

// BoolstructPGTask_sUint64Uint64Handler is the signature for a completion handler block.
type BoolstructPGTask_sUint64Uint64Handler = func(*uintptr, uint64, uint64) bool

// PGDestroyTask handles The block signature for a routine that destroys a task.

// NewPGDestroyTaskBlock wraps a Go [PGDestroyTask] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDestroyTaskBlock(handler PGDestroyTask) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *uintptr) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayCoord_tUint32Handler handles completion with primitive and object results.
type PGDisplayCoord_tUint32Handler = func(PGDisplayCoord_t, uint32)

// NewPGDisplayCoord_tUint32Block wraps a Go [PGDisplayCoord_tUint32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayCoord_tUint32Block(handler PGDisplayCoord_tUint32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive PGDisplayCoord_t, extra0 uint32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayCursorGlyphHandler handles The block signature for a routine that handles changes to the cursor’s appearance.

// NewPGDisplayCursorGlyphHandlerBlock wraps a Go [PGDisplayCursorGlyphHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayCursorGlyphHandlerBlock(handler PGDisplayCursorGlyphHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive appkit.NSBitmapImageRep, extra0 PGDisplayCoord_t) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayCursorMoveHandler is the signature for a completion handler block.

// NewPGDisplayCursorMoveHandlerBlock wraps a Go [PGDisplayCursorMoveHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayCursorMoveHandlerBlock(handler PGDisplayCursorMoveHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayCursorShowHandler handles The block signature for a routine that handles changes to the cursor’s visibility.

// NewPGDisplayCursorShowBlock wraps a Go [PGDisplayCursorShowHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayCursorShowBlock(handler PGDisplayCursorShowHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayModeChangeHandler handles The block signature for a routine that handles changes to the display’s graphics mode.

// NewPGDisplayModeChangeHandlerBlock wraps a Go [PGDisplayModeChangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayModeChangeHandlerBlock(handler PGDisplayModeChangeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive PGDisplayCoord_t, extra0 uint32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGDisplayNewFrameEventHandler handles The block signature for a routine that handles frame updates from the guest.

// NewPGDisplayNewFrameEventHandlerBlock wraps a Go [PGDisplayNewFrameEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGDisplayNewFrameEventHandlerBlock(handler PGDisplayNewFrameEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// PGRaiseInterrupt handles The block signature for a routine that raises interrupts in the guest environment.

// NewPGRaiseInterruptBlock wraps a Go [PGRaiseInterrupt] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGRaiseInterruptBlock(handler PGRaiseInterrupt) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint32) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGRemoveTraceRange handles The block signature for a routine that removes a trace range.

// NewPGRemoveTraceRangeBlock wraps a Go [PGRemoveTraceRange] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGRemoveTraceRangeBlock(handler PGRemoveTraceRange) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *uintptr) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PGTraceRangeHandler handles The block signature for a routine that handles trace requests.

// NewPGTraceRangeHandlerBlock wraps a Go [PGTraceRangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewPGTraceRangeHandlerBlock(handler PGTraceRangeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *PGPhysicalMemoryRange_s) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// UintptrPGPhysicalMemoryRange_sHandler handles a primitive value and returns a primitive value.
type UintptrPGPhysicalMemoryRange_sHandler = func(*PGPhysicalMemoryRange_s) *uintptr

// NewUintptrPGPhysicalMemoryRange_sBlock wraps a Go [UintptrPGPhysicalMemoryRange_sHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewUintptrPGPhysicalMemoryRange_sBlock(handler UintptrPGPhysicalMemoryRange_sHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *PGPhysicalMemoryRange_s) *uintptr {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// UintptrUint64Handler handles a primitive value and returns a primitive value.
type UintptrUint64Handler = func(uint64) *uintptr

// NewUintptrUint64Block wraps a Go [UintptrUint64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewUintptrUint64Block(handler UintptrUint64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint64) *uintptr {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// structPGTask_sHandler is the signature for a completion handler block.
type structPGTask_sHandler = func(*uintptr)

// structPGTraceRange_sHandler is the signature for a completion handler block.
type structPGTraceRange_sHandler = func(*uintptr)
