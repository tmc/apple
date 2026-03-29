
// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

// Package coremediaio provides Go bindings for the CoreMediaIO framework.
//
// Securely support custom camera devices in macOS.
//
// Use the Core Media I/O framework to enable support for custom camera devices in macOS. Starting in macOS 12.3, the framework builds on [System Extensions](<doc://com.apple.documentation/documentation/SystemExtensions>) to enable you to support custom devices while maintaining system privacy and security protections. The system prevents apps from loading extension code into their process to ensure that they can’t bypass macOS privacy protections or mask their identity.
//
// # Providers
//
//   - Creating a camera extension with Core Media I/O: Build high-performance camera drivers that are secure and simple to deploy.
//   - Overriding the default USB video class extension: Create a simple DriverKit extension to override the default driver-matching behavior for USB devices.
//   - CMIOExtensionProvider: An object that manages device connections for a provider.
//   - CMIOExtensionProviderSource: A protocol for objects that act as provider sources.
//   - CMIOExtensionProviderProperties: An object that manages the properties of an extension provider.
//
// # Devices
//
//   - CMIOExtensionDevice: An object that represents a physical or virtual device.
//   - CMIOExtensionDeviceSource: A protocol for objects that act as device sources.
//   - CMIOExtensionDeviceProperties: An object that defines the properties of a device.
//
// # Streams
//
//   - CMIOExtensionStream: An object that represents a stream of media data.
//   - CMIOExtensionStreamSource: A protocol for objects that act as stream sources.
//   - CMIOExtensionStreamProperties: An object that describes the properties of an extension stream.
//   - CMIOExtensionClient: An object that represents a client of the extension.
//
// # Properties
//
//   - CMIOExtensionProperty: A structure that defines the properties that providers, devices, and streams support.
//   - CMIOExtensionPropertyState: An object that describes the state of a property.
//   - CMIOExtensionPropertyAttributes: An object that describes the attributes of a property.
//   - CMIOExtensionInfoDictionaryKey: A key that specifies the extension information dictionary.
//   - CMIOExtensionMachServiceNameKey: A key that specifies the mach service name.
//
// # DAL Plug-Ins
//
//   - Device Abstraction Layer (DAL) Plug-Ins: API reference for legacy DAL plug-ins.
//
// [CoreMediaIO Documentation]: https://developer.apple.com/documentation/CoreMediaIO
package coremediaio

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreMediaIO.framework/CoreMediaIO"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreMediaIO: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

