// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

// Package hypervisor provides Go bindings for the Hypervisor framework.
//
// Build virtualization solutions on top of a lightweight hypervisor, without
// third-party kernel extensions.
//
// Hypervisor provides C APIs so you can interact with virtualization
// technologies in user space, without writing kernel extensions (KEXTs). As a
// result, the apps you create using this framework are suitable for
// distribution on the Mac App Store.
//
// Mac App Store: https://www.appstore.com/
//
// # Platforms
//
//   - [Apple Silicon]: Create and run virtual machines on Apple silicon. ([OSHVVmConfig], [HVVmConfig], [HVReturn])
//   - [Intel-based Mac]: Create and run virtual machines on Intel-based Mac computers. ([HVVmOptions], [HVCapability], [HVIonMessage], [HVIonFlags], [HVReturn])
//
// # Entitlements
//
//   - [com.apple.security.hypervisor]: A Boolean value that indicates whether the app creates and manages virtual machines.
//   - [com.apple.vm.networking]: A Boolean that indicates whether the app manages virtual network interfaces without escalating privileges to the root user.
//   - [com.apple.vm.device-access]: A Boolean value that indicates whether the app captures USB devices and uses them in the guest-operating system.
//
// # Functions
//
//   - [HVVmConfigGetDefaultIPAGranule]
//   - [HVVmConfigGetIPAGranule]
//   - [HVVmConfigSetIPAGranule]
//
// # Enumerations
//
//   - [HVIPAGranule]//
//
// [Apple Silicon]: https://developer.apple.com/documentation/hypervisor/apple-silicon
// [Intel-based Mac]: https://developer.apple.com/documentation/hypervisor/intel-based-mac
// [com.apple.security.hypervisor]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.hypervisor
// [com.apple.vm.device-access]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.vm.device-access
// [com.apple.vm.networking]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.vm.networking
package hypervisor

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Hypervisor library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Hypervisor.framework/Hypervisor",
	"/usr/lib/libHypervisor.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: Hypervisor: failed to load framework from any known path\n")
}
