// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

// Package systemconfiguration provides Go bindings for the SystemConfiguration framework.
//
// Allow applications to access a device’s network configuration settings.
// Determine the reachability of the device, such as whether Wi-Fi or cell
// connectivity are active.
//
// This collection of documents describes the programming interfaces of the
// System Configuration framework. The System Configuration framework provides
// functions that determine the reachability of target hosts in both a
// synchronous and an asynchronous manner. It also provides error detection
// facilities.
//
// # Entitlements
//
//   - [Access Wi-Fi Information Entitlement]: A Boolean value indicating whether your app can access information about the connected Wi-Fi network.//
//
// [Access Wi-Fi Information Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.networking.wifi-info
package systemconfiguration

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the SystemConfiguration library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/SystemConfiguration.framework/SystemConfiguration",
	"/usr/lib/libSystemConfiguration.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: SystemConfiguration: failed to load framework from any known path\n")
	}
}
