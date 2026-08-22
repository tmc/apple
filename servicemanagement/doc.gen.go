// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

// Package servicemanagement provides Go bindings for the ServiceManagement framework.
//
// Manage startup items, launch agents, and launch daemons from within an app.
//
// Use Service Management to install and observe the permission settings of
// three supplemental helper executables that macOS supports. You can use all
// three of these to provide additional functionality related to your app,
// from inside your app’s bundle:
//
// # Essentials
//
//   - [Updating helper executables from earlier versions of macOS]: Simplify your app’s helper executables and support new authorization controls.
//   - [Updating your app package installer to use the new Service Management API]: Learn about the Service Management API with a GUI-less agent app.
//
// # Management
//
//   - [SMAppService]: An object the framework uses to control helper executables that live inside an app’s main bundle.
//   - [Authorization Constants]: Constants that describe the ability to authorize helper executables or modify daemon applications.
//   - [Property List Keys]: Property list keys that describe the kinds of applications, daemons, and helper executables the framework manages.
//
// # Status
//
//   - [SMAppServiceStatus]: Constants that describe the registration or authorization status of a helper executable.
//
// # Errors
//
//   - [Service Management Errors]: Errors that the framework returns.
//
// # Variables
//
//   - [SMAppServiceErrorDomain]//
//
// # Key Types
//
//   - [SMAppService] - An object the framework uses to control helper executables that live inside an app’s main bundle.
//
// [Authorization Constants]: https://developer.apple.com/documentation/servicemanagement/authorization-constants
// [Property List Keys]: https://developer.apple.com/documentation/servicemanagement/property-list-keys
// [Service Management Errors]: https://developer.apple.com/documentation/servicemanagement/service-management-errors
// [Updating helper executables from earlier versions of macOS]: https://developer.apple.com/documentation/servicemanagement/updating-helper-executables-from-earlier-versions-of-macos
// [Updating your app package installer to use the new Service Management API]: https://developer.apple.com/documentation/servicemanagement/updating-your-app-package-installer-to-use-the-new-service-management-api
package servicemanagement

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ServiceManagement library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ServiceManagement.framework/ServiceManagement",
	"/usr/lib/libServiceManagement.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: ServiceManagement: failed to load framework from any known path\n")
	}
}
