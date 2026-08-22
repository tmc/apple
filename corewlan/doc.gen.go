// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

// Package corewlan provides Go bindings for the CoreWLAN framework.
//
// Query AirPort interfaces and choose wireless networks.
//
// The CoreWLAN framework provides APIs for querying AirPort interfaces and
// choosing networks.
//
// # Classes
//
//   - [CWChannel]: Encapsulates an IEEE 802.11 channel.
//   - [CWConfiguration]: Encapsulates an immutable configuration for an AirPort WLAN interface.
//   - [CWInterface]: Encapsulates an IEEE 802.11 interface.
//   - [CWMutableConfiguration]: Encapsulates a mutable configuration for an AirPort WLAN interface.
//   - [CWMutableNetworkProfile]: Encapsulates a mutable network profile entry.
//   - [CWNetwork]: Encapsulates an IEEE 802.11 network, providing read-only accessors to various properties of the network.
//   - [CWNetworkProfile]: Encapsulates an immutable network profile entry.
//   - [CWWiFiClient]: A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
//
// # Protocols
//
//   - [CWEventDelegate]: The interface a Wi-Fi client object uses to notify its delegate about Wi-Fi events.
//
// # Key Types
//
//   - [CWInterface] - Encapsulates an IEEE 802.11 interface.
//   - [CWNetwork] - Encapsulates an IEEE 802.11 network, providing read-only accessors to various properties of the network.
//   - [CWWiFiClient] - A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
//   - [CWConfiguration] - Encapsulates an immutable configuration for an AirPort WLAN interface.
//   - [CWMutableConfiguration] - Encapsulates a mutable configuration for an AirPort WLAN interface.
//   - [CWNetworkProfile] - Encapsulates an immutable network profile entry.
//   - [CWChannel] - Encapsulates an IEEE 802.11 channel.
//   - [CWMutableNetworkProfile] - Encapsulates a mutable network profile entry.
package corewlan

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreWLAN library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreWLAN.framework/CoreWLAN",
	"/usr/lib/libCoreWLAN.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreWLAN: failed to load framework from any known path\n")
	}
}
