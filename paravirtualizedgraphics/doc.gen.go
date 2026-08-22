// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

// Package paravirtualizedgraphics provides Go bindings for the ParavirtualizedGraphics framework.
//
// Add graphics acceleration to your guest driver stack.
//
// If you have an app that implements hardware-level virtualization,
// performance inside the virtual machine is critical, particularly for
// graphics. The ParavirtualizedGraphics framework implements
// hardware-accelerated graphics for macOS running in a virtual machine,
// hereafter known as the guest. The operating system provides a graphics
// driver that runs inside the guest, communicating with the framework in the
// host operating system to take advantage of Metal-accelerated graphics.
//
// # PCI Device Characteristics
//
//   - [PGCopyOptionROMURL]: Copies the URL of the ROM image to use on the guest graphics device.
//
// # Devices
//
//   - [PGDeviceDescriptor]: A description of the paravirtualized graphics device to create. ([PGCreateTask], [PGDestroyTask], [PGMapMemory], [PGUnmapMemory], [PGReadMemory])
//   - [PGDevice]: A paravirtualized GPU device object. ([PGResumeErrorCode])
//
// # Displays
//
//   - [PGDisplayDescriptor]: A descriptor for a virtual display. ([PGDisplayCursorGlyphHandler], [PGDisplayCursorShowHandler], [PGDisplayModeChangeHandler], [PGDisplayNewFrameEventHandler])
//   - [PGDisplay]: An object that provides display functionality to the guest operating system in a way that the host-side virtual machine app can intercept.
//   - [PGDisplayMode]: A description of a supported display mode.
//   - [PGDisplayCoord_t]: Coordinates that describe sizes or offsets within a 2D array of pixels.
//
// # Functions
//
//   - [PGCreateDeviceWithDescriptor]
//   - [PGNewDeviceWithDescriptor]
//
// # Key Types
//
//   - [PGDeviceDescriptor] - A description of the paravirtualized graphics device to create.
//   - [PGDisplayDescriptor] - A descriptor for a virtual display.
//   - [PGDisplayMode] - A description of a supported display mode.
package paravirtualizedgraphics

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ParavirtualizedGraphics library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ParavirtualizedGraphics.framework/ParavirtualizedGraphics",
	"/usr/lib/libParavirtualizedGraphics.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: ParavirtualizedGraphics: failed to load framework from any known path\n")
	}
}
