// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

// Package systemextensions provides Go bindings for the SystemExtensions framework.
//
// Install and manage user space code that extends the capabilities of macOS.
//
// Extend the capabilities of macOS by installing and managing system
// extensions—drivers and other low-level code—in user space rather than
// in the kernel. By running in user space, system extensions can’t
// compromise the security or stability of macOS. The system grants these
// extensions a high level of privilege, so they can perform the kinds of
// tasks previously reserved for kernel extensions (KEXTs).
//
// # Essentials
//
//   - [Implementing drivers, system extensions, and kexts]: Create drivers and system extensions to communicate with hardware and provide low-level services, and only use kernel extensions for a few tasks.
//   - [Debugging and testing system extensions]: Debug your system extensions by temporarily disabling the security checks that macOS performs during the installation process.
//   - [System Extension Entitlement]: A Boolean value that indicates whether your app has permission to activate or deactivate system extensions.
//
// # Usage descriptions
//
//   - [SystemExtensionUsageDescriptionKey]: A message that tells the user why the app is trying to install a system extension bundle.
//   - [OSBundleUsageDescriptionKey]: A message that tells the user why the app is trying to install a driver extension bundle.
//
// # Extension activation and deactivation
//
//   - [Installing System Extensions and Drivers]: Activate system extensions and drivers to make them available to the system, and update or deactivate them as needed.
//   - [OSSystemExtensionManager]: A type that facilitates activation and deactivation of system extensions. ([OSSystemExtensionRequest])
//   - [OSSystemExtensionRequest]: A request to activate or deactivate a system extension. ([OSSystemExtensionRequestDelegate])
//   - [System Extension Redistributable Entitlement]: A Boolean value that indicates whether other development teams may distribute a system extension you create.
//
// # Errors
//
//   - [OSSystemExtensionErrorCode]: Error codes for system extensions.
//   - [OSSystemExtensionErrorDomain]: The error domain identifying system extension errors.
//
// # Classes
//
//   - [OSSystemExtensionInfo]
//   - [OSSystemExtensionsWorkspace]
//
// # Protocols
//
//   - [OSSystemExtensionsWorkspaceObserver]//
//
// # Key Types
//
//   - [OSSystemExtensionProperties] - Properties that identify a specific version of a system extension.
//   - [OSSystemExtensionRequest] - A request to activate or deactivate a system extension.
//   - [OSSystemExtensionsWorkspace]
//   - [OSSystemExtensionInfo]
//   - [OSSystemExtensionManager] - A type that facilitates activation and deactivation of system extensions.
//
// [Debugging and testing system extensions]: https://developer.apple.com/documentation/DriverKit/debugging-and-testing-system-extensions
// [Implementing drivers, system extensions, and kexts]: https://developer.apple.com/documentation/systemextensions/implementing-drivers-system-extensions-and-kexts
// [Installing System Extensions and Drivers]: https://developer.apple.com/documentation/systemextensions/installing-system-extensions-and-drivers
// [System Extension Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.system-extension.install
// [System Extension Redistributable Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.system-extension.redistributable
package systemextensions

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the SystemExtensions library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/SystemExtensions.framework/SystemExtensions",
	"/usr/lib/libSystemExtensions.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: SystemExtensions: failed to load framework from any known path\n")
	}
}
