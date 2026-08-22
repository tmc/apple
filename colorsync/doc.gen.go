// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

// Package colorsync provides Go bindings for the ColorSync framework.
//
// Reproduce colors accurately across a range of input, output, and display
// devices.
//
// ColorSync is the color-management engine on Apple platforms. For most apps,
// color management happens automatically through higher-level frameworks such
// as Core Graphics and Core Image. Use ColorSync directly when your app
// needs to manage color itself; for example, a professional photo, print, or
// video app that builds custom transforms, or a tool that inspects and
// calibrates the profiles assigned to devices and displays.
//
// # Color conversion
//
//   - [Color transforms]: Convert color from one profile’s color space to another.
//   - [Pixel format and data layout]: Describe the memory layout of the pixel buffers a color transform reads and writes. ([ColorSyncAlphaInfo], [ColorSyncDataDepth])
//
// # Profile and HDR metadata
//
//   - [Color profiles]: Work with the ICC profiles that describe device and working color spaces. ([ColorSyncProfileIterateCallback], [ColorSyncMD5])
//   - [Headroom Adaptive Gain Curve]: Work with SMPTE ST 2094-50 tone-mapping metadata shared between HDR stills and video.
//
// # System color management
//
//   - [Color devices]: Manage the color profiles assigned to displays, printers, scanners, and cameras. ([ColorSyncDeviceProfileIterateCallback])
//   - [Color management modules]: Work with the Color Management Modules that perform color conversions. ([ColorSyncCMMIterateCallback], [CMMApplyTransformProc], [CMMCreateTransformPropertyProc], [CMMInitializeLinkProfileProc], [CMMInitializeTransformProc])
//
// # Supporting types and conventions
//
//   - [Supporting types and conventions]: Reference the signatures and conventions that support the color-management APIs.//
//
// [Color devices]: https://developer.apple.com/documentation/colorsync/color-devices
// [Color management modules]: https://developer.apple.com/documentation/colorsync/color-management-modules
// [Color profiles]: https://developer.apple.com/documentation/colorsync/color-profiles
// [Color transforms]: https://developer.apple.com/documentation/colorsync/color-transforms
// [Headroom Adaptive Gain Curve]: https://developer.apple.com/documentation/colorsync/headroom-adaptive-gain-curve
// [Pixel format and data layout]: https://developer.apple.com/documentation/colorsync/pixel-format
// [Supporting types and conventions]: https://developer.apple.com/documentation/colorsync/supporting-types-and-conventions
package colorsync

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ColorSync library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ColorSync.framework/ColorSync",
	"/usr/lib/libColorSync.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: ColorSync: failed to load framework from any known path\n")
	}
}
