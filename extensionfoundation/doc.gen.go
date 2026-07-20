// Code generated from Apple documentation for ExtensionFoundation. DO NOT EDIT.

// Package extensionfoundation provides Go bindings for the ExtensionFoundation framework.
//
// Create executable bundles to extend the functionality of other apps.
//
// An app extension is an executable code bundle that extends the capabilities
// of the system, or the capabilities of another app. At the system level, app
// extensions give you a way to add your custom capabilities to system
// features. For example, Creating a widget extension display your app’s
// content in specific locations like the iOS Home Screen and Lock screen. To
// build app extensions for system features, you typically use a different
// framework or a dedicated set of types instead of this framework.
//
// # Essentials
//
//   - [Adding support for app extensions to your app]: Create an app extension model by defining your code’s extension points and communicating with app extensions at runtime.
//
// # App-extension setup
//
//   - [Building an app extension to support a host app]: Create an app extension to perform tasks in a separate process from a host app.
//   - AppExtension: An interface you use to declare the content, structure, and behavior of an app extension.
//   - AppExtensionConfiguration: An interface you use to configure the XPC connection in your app extension.
//   - ConnectionHandler: A type that contains a custom closure that handles incoming XPC connections.
//
// # Host-app configuration
//
//   - [Discovering app extensions from your app]: Find the app extensions that match your host app’s extension points and are available to use.
//   - AppExtensionProcess: A type the host app creates to launch and manage an app extension.
//   - AppExtensionIdentity: A type that uniquely identifies an app extension on the system.
//
// # Extension-point management
//
//   - AppExtensionPoint: A type you use to declare your host app’s extension points and bind to them from app extensions.
//   - ExtensionPointDefining: An interface that extension point types adopt.//
//
// [Adding support for app extensions to your app]: https://developer.apple.com/documentation/extensionfoundation/adding-support-for-app-extensions-to-your-app
// [Building an app extension to support a host app]: https://developer.apple.com/documentation/extensionfoundation/building-an-app-extension-to-support-a-host-app
// [Discovering app extensions from your app]: https://developer.apple.com/documentation/extensionfoundation/discovering-app-extensions-from-your-app
package extensionfoundation

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ExtensionFoundation library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ExtensionFoundation.framework/ExtensionFoundation",
	"/usr/lib/libExtensionFoundation.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: ExtensionFoundation: failed to load framework from any known path\n")
	}
}
