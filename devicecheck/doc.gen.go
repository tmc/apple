// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

// Package devicecheck provides Go bindings for the DeviceCheck framework.
//
// Reduce fraudulent use of your services by managing device state and
// asserting app integrity.
//
// The DeviceCheck services consist of both a framework interface that you
// access from your app and an Apple server interface that you access from
// your own server.
//
// # Device identification
//
//   - [Accessing and modifying per-device data]: Use a token from your app to query and modify two per-device binary digits stored on an Apple server.
//   - [DCDevice]: A representation of a device that provides a unique, authenticated token.
//
// # App Attest
//
//   - [Establishing your app’s integrity]: Ensure that requests your server receives come from legitimate instances of your app.
//   - [Validating apps that connect to your server]: Verify that connections to your server come from legitimate instances of your app.
//   - [Assessing fraud risk]: Request and analyze risk data using server-to-server calls.
//   - [Preparing to use the app attest service]: Test your implementation in a development environment and onboard users gradually.
//   - [Attestation Object Validation Guide]: Use this guide to validate your implementation of verifying the attestation object verification process.
//   - [DCAppAttestService]: A service that you use to validate the instance of your app running on a device.
//   - [App Attest Environment]: The environment for an app that uses the App Attest service to validate itself.
//
// # Errors
//
//   - [DCError]: DeviceCheck error codes.
//   - [DCErrorDomain]: The error domain for errors associated with DeviceCheck APIs.//
//
// # Key Types
//
//   - [DCAppAttestService] - A service that you use to validate the instance of your app running on a device.
//   - [DCDevice] - A representation of a device that provides a unique, authenticated token.
//
// [Accessing and modifying per-device data]: https://developer.apple.com/documentation/devicecheck/accessing-and-modifying-per-device-data
// [App Attest Environment]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.devicecheck.appattest-environment
// [Assessing fraud risk]: https://developer.apple.com/documentation/devicecheck/assessing-fraud-risk
// [Attestation Object Validation Guide]: https://developer.apple.com/documentation/devicecheck/attestation-object-validation-guide
// [Establishing your app’s integrity]: https://developer.apple.com/documentation/devicecheck/establishing-your-app-s-integrity
// [Preparing to use the app attest service]: https://developer.apple.com/documentation/devicecheck/preparing-to-use-the-app-attest-service
// [Validating apps that connect to your server]: https://developer.apple.com/documentation/devicecheck/validating-apps-that-connect-to-your-server
package devicecheck

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the DeviceCheck library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/DeviceCheck.framework/DeviceCheck",
	"/usr/lib/libDeviceCheck.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: DeviceCheck: failed to load framework from any known path\n")
	}
}
