// Package custommmio provides experimental custom MMIO device helpers.
//
// It is intended to wrap private custom MMIO device configuration types,
// including region layout, IRQ wiring, provider hookup, and save or restore
// support flags for plugin-backed devices.
//
// This package should stay small and explicit about ownership of raw MMIO
// details such as regions and interrupts.
package custommmio
