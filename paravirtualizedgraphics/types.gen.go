// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

// C struct types

// PGDisplayCoord_t - Coordinates that describe sizes or offsets within a 2D array of pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayCoord_t
type PGDisplayCoord_t struct {
	X uint16 // The horizontal coordinate value.
	Y uint16 // The vertical coordinate value.

}

// PGPhysicalMemoryRange_s - A range in the guest virtual machine’s physical memory address space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGPhysicalMemoryRange_s
type PGPhysicalMemoryRange_s struct {
	PhysicalAddress uint64 // The starting address of the range in physical memory.
	PhysicalLength  uint64 // The length of the range.

}
