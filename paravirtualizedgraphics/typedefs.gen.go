// Code generated from Apple documentation. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"unsafe"

	"github.com/tmc/apple/appkit"
)

// PGAddTraceRange is the block signature for a routine that adds a trace range.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGAddTraceRange
type PGAddTraceRange = func(range_ *PGPhysicalMemoryRange_s, handler func(*PGPhysicalMemoryRange_s)) *uintptr

// PGCreateTask is the block signature for a routine that creates a task.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGCreateTask
type PGCreateTask = func(vmSize uint64, baseAddress unsafe.Pointer) *uintptr

// PGDestroyTask is the block signature for a routine that destroys a task.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDestroyTask
type PGDestroyTask = func(task *uintptr)

// PGDisplayCursorGlyphHandler is the block signature for a routine that handles changes to the cursor’s appearance.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayCursorGlyphHandler
type PGDisplayCursorGlyphHandler = func(glyph appkit.NSBitmapImageRep, hotSpot PGDisplayCoord_t)

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayCursorMoveHandler
type PGDisplayCursorMoveHandler = func()

// PGDisplayCursorShowHandler is the block signature for a routine that handles changes to the cursor’s visibility.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayCursorShowHandler
type PGDisplayCursorShowHandler = func(show bool)

// PGDisplayModeChangeHandler is the block signature for a routine that handles changes to the display’s graphics mode.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayModeChangeHandler
type PGDisplayModeChangeHandler = func(sizeInPixels PGDisplayCoord_t, pixelFormat uint32)

// PGDisplayNewFrameEventHandler is the block signature for a routine that handles frame updates from the guest.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayNewFrameEventHandler
type PGDisplayNewFrameEventHandler = func()

// PGMapMemory is the block signature for a routine that maps guest physical memory into a task.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGMapMemory
type PGMapMemory = func(task *uintptr, rangeCount uint32, virtualOffset uint64, readOnly bool, ranges *PGPhysicalMemoryRange_s) bool

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGPhysicalMemoryRange_t
type PGPhysicalMemoryRange_t = PGPhysicalMemoryRange_s

// PGRaiseInterrupt is the block signature for a routine that raises interrupts in the guest environment.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGRaiseInterrupt
type PGRaiseInterrupt = func(vector uint32)

// PGReadMemory is the block signature for a routine that copies data from guest physical memory into host memory.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGReadMemory
type PGReadMemory = func(physicalAddress uint64, length uint64, dst unsafe.Pointer) bool

// PGRemoveTraceRange is the block signature for a routine that removes a trace range.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGRemoveTraceRange
type PGRemoveTraceRange = func(range_ *uintptr)

// PGTask_t is an opaque data pointer representing a specific virtual task.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGTask_t
// PGTask_t is an unresolved C aggregate typedef.
type PGTask_t unsafe.Pointer

// PGTraceRangeHandler is the block signature for a routine that handles trace requests.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGTraceRangeHandler
type PGTraceRangeHandler = func(dirty *PGPhysicalMemoryRange_s)

// PGTraceRange_t is an opaque data pointer representing a specific trace.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGTraceRange_t
// PGTraceRange_t is an unresolved C aggregate typedef.
type PGTraceRange_t unsafe.Pointer

// PGUnmapMemory is the block signature for a routine that unmaps guest physical memory from a task.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGUnmapMemory
type PGUnmapMemory = func(task *uintptr, virtualOffset uint64, length uint64) bool

// PGPhysicalMemoryRange is a Go-name alias for PGPhysicalMemoryRange_t.
type PGPhysicalMemoryRange = PGPhysicalMemoryRange_t

// PGTask is a Go-name alias for PGTask_t.
type PGTask = PGTask_t

// PGTraceRange is a Go-name alias for PGTraceRange_t.
type PGTraceRange = PGTraceRange_t
