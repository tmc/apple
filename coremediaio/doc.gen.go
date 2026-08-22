// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

// Package coremediaio provides Go bindings for the CoreMediaIO framework.
//
// Securely support custom camera devices in macOS.
//
// Use the Core Media I/O framework to enable support for custom camera
// devices in macOS. Starting in macOS 12.3, the framework builds on [System
// Extensions] to enable you to support custom devices while maintaining
// system privacy and security protections. The system prevents apps from
// loading extension code into their process to ensure that they can’t
// bypass macOS privacy protections or mask their identity.
//
// # Providers
//
//   - [Creating a camera extension with Core Media I/O]: Build high-performance camera drivers that are secure and simple to deploy.
//   - [Overriding the default USB video class extension]: Create a simple DriverKit extension to override the default driver-matching behavior for USB devices.
//   - [CMIOExtensionProviderSource]: A protocol for objects that act as provider sources.
//
// # Devices
//
//   - [CMIOExtensionDeviceSource]: A protocol for objects that act as device sources.
//
// # Streams
//
//   - [CMIOExtensionStreamSource]: A protocol for objects that act as stream sources.
//
// # Properties
//
//   - [CMIOExtensionProperty]: A structure that defines the properties that providers, devices, and streams support.
//   - [CMIOExtensionInfoDictionaryKey]: A key that specifies the extension information dictionary.
//   - [CMIOExtensionMachServiceNameKey]: A key that specifies the mach service name.//
//
// [Creating a camera extension with Core Media I/O]: https://developer.apple.com/documentation/coremediaio/creating-a-camera-extension-with-core-media-i-o
// [Overriding the default USB video class extension]: https://developer.apple.com/documentation/coremediaio/overriding-the-default-usb-video-class-extension
package coremediaio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreMediaIO library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreMediaIO.framework/CoreMediaIO",
	"/usr/lib/libCoreMediaIO.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreMediaIO: failed to load framework from any known path\n")
	}
}
