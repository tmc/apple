
// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

// Package servicemanagement provides Go bindings for the ServiceManagement framework.
//
// Manage startup items, launch agents, and launch daemons from within an app.
//
// Use Service Management to install and observe the permission settings of three supplemental helper executables that macOS supports. You can use all three of these to provide additional functionality related to your app, from inside your app’s bundle:
//
// # Essentials
//
//   - Updating helper executables from earlier versions of macOS: Simplify your app’s helper executables and support new authorization controls.
//   - Updating your app package installer to use the new Service Management API: Learn about the Service Management API with a GUI-less agent app.
//
// # Management
//
//   - SMAppService: An object the framework uses to control helper executables that live inside an app’s main bundle.
//   - SMJobBless(_:_:_:_:): Submits the executable for the given label as a job to .
//   - Authorization Constants: Constants that describe the ability to authorize helper executables or modify daemon applications.
//   - Property List Keys: Property list keys that describe the kinds of applications, daemons, and helper executables the framework manages.
//
// # Enablement
//
//   - SMLoginItemSetEnabled(_:_:): Enables a helper executable in the main app-bundle directory.
//
// # Status
//
//   - SMAppService.Status: Constants that describe the registration or authorization status of a helper executable.
//
// # Errors
//
//   - Service Management Errors: Errors that the framework returns.
//
// # Deprecated
//
//   - Deprecated Symbols
//
// # Variables
//
//   - SMAppServiceErrorDomain
//
// # Key Types
//
//   - [SMAppService] - An object the framework uses to control helper executables that live inside an app’s main bundle.
//
// [ServiceManagement Documentation]: https://developer.apple.com/documentation/ServiceManagement
package servicemanagement

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ServiceManagement.framework/ServiceManagement"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ServiceManagement: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

