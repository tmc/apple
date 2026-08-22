// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

// Package iousbhost provides Go bindings for the IOUSBHost framework.
//
// Create host-mode user space drivers for USB devices.
//
// With the [IOUSBHost] framework, you can access custom and
// non–class-compliant USB devices from within your apps. Use this framework
// to connect to cameras, audio devices, scanners, printers, keyboards, mouse
// devices, MIDI keyboards, and USB hubs.
//
// # Function Drivers
//
//   - [IOUSBHostInterface]: The class for accessing USB-related services.
//   - [IOUSBHostPipe]: The class that sends control, bulk, interrupt, and isochronous input/output requests for function drivers, and manages stream capabilities. ([IOUSBHostCompletionHandler], [IOUSBHostIsochronousCompletionHandler], [IOUSBHostTime], [IOUSBHostIsochronousFrame], [IOUSBHostCompletionHandler])
//   - [IOUSBHostStream]: The class responsible for sending stream data for function drivers. ([IOUSBHostCompletionHandler])
//
// # Device Drivers
//
//   - [IOUSBHostDevice]: The class that claims and configures devices, retrieves descriptors, and sends device requests.
//
// # Base Classes
//
//   - [IOUSBHostObject]: This class provides basic functionality for sending device requests and retrieving descriptors. ([IOUSBHostObjectInitOptions], [IOUSBHostInterestHandler], [IOUSBHostCompletionHandler], [IOUSBHostAbortOption], [IOUSBHostDevice])
//   - [IOUSBHostIOSource]: This class provides basic functionality for deriving pipe and stream classes. ([IOUSBHostPipe], [IOUSBHostStream])
//
// # IOServicePlane Properties
//
//   - [IOUSBHostInterfacePropertyKey]: Properties of a USB interface that describe its state.
//   - [IOUSBHostDevicePropertyKey]: Properties of a USB device that describe its state. ([IOUSBHostPropertyKey])
//   - [IOUSBHostMatchingPropertyKey]: Properties for implementing the matching service.
//   - [IOUSBHostPropertyKey]: Properties that the USB host device and interface classes share.
//
// # Version Number
//
//   - [IOUSBHostVersionNumber]: The version number of the framework.
//   - [IOUSBHostVersionString]: A string representation of the framework’s version number.
//
// # Error Domain
//
//   - [IOUSBHostErrorDomain]: The error domain for the framework.
//
// # Classes
//
//   - [IOUSBHostCIControllerStateMachine]
//   - [IOUSBHostCIDeviceStateMachine]
//   - [IOUSBHostCIEndpointStateMachine]
//   - [IOUSBHostCIPortStateMachine]
//   - [IOUSBHostControllerInterface]
//
// # Enumerations
//
//   - [IOUSBHostObjectDataOptions]
//
// # Key Types
//
//   - [IOUSBHostControllerInterface]
//   - [IOUSBHostCIPortStateMachine]
//   - [IOUSBHostPipe] - The class that sends control, bulk, interrupt, and isochronous input/output requests for function drivers, and manages stream capabilities.
//   - [IOUSBHostObject] - This class provides basic functionality for sending device requests and retrieving descriptors.
//   - [IOUSBHostCIEndpointStateMachine]
//   - [IOUSBHostCIDeviceStateMachine]
//   - [IOUSBHostCIControllerStateMachine]
//   - [IOUSBHostInterface] - The class for accessing USB-related services.
//   - [IOUSBHostStream] - The class responsible for sending stream data for function drivers.
//   - [IOUSBHostIOSource] - This class provides basic functionality for deriving pipe and stream classes.
package iousbhost

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the IOUSBHost library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/IOUSBHost.framework/IOUSBHost",
	"/usr/lib/libIOUSBHost.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: IOUSBHost: failed to load framework from any known path\n")
	}
}
