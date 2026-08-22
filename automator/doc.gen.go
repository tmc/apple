// Code generated from Apple documentation for Automator. DO NOT EDIT.

// Package automator provides Go bindings for the Automator framework.
//
// Develop actions that the Automator app can load and run. View, edit, and
// run Automator workflows in your app.
//
// The Automator framework supports the development of actions for the
// Automator app, as well as the ability to run a workflow in developer apps.
// An action is a bundle that, when loaded and run, performs a specific task,
// such as copying a file or cropping an image. Using Automator, users can
// construct and execute workflows consisting of a sequence of actions.
// Developers can also load and execute workflows in their apps. As a workflow
// executes, the output of one action is typically passed as the input to the
// next action. Automator loads action bundles from standard locations in the
// file system: `/System/Library/Automator`, `/Library/Automator`, and
// `~/Library/Automator`.
//
// # Actions
//
//   - [AMBundleAction]: An object that represents an Automator action that’s a loadable bundle.
//   - [AMShellScriptAction]: An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
//   - [AMAction]: An abstract class that defines the interface and general characteristics of Automator actions. ([AMLogLevel])
//
// # Workflows
//
//   - [AMWorkflow]: An object that lets you use an Automator workflow in your app.
//   - [AMWorkflowController]: An object that lets you manage an Automator workflow in your app. ([AMWorkflowControllerDelegate])
//   - [AMWorkflowView]: An object that lets you view and edit Automator workflows in your app.
//   - [AMWorkspace]: A workspace for running an Automator workflow.
//
// # Errors
//
//   - [AMErrorCode]: Automator error codes.
//
// # Key Types
//
//   - [AMAction] - An abstract class that defines the interface and general characteristics of Automator actions.
//   - [AMWorkflow] - An object that lets you use an Automator workflow in your app.
//   - [AMWorkflowController] - An object that lets you manage an Automator workflow in your app.
//   - [AMBundleAction] - An object that represents an Automator action that’s a loadable bundle.
//   - [AMShellScriptAction] - An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
//   - [AMWorkflowView] - An object that lets you view and edit Automator workflows in your app.
//   - [AMWorkspace] - A workspace for running an Automator workflow.
package automator

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Automator library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Automator.framework/Automator",
	"/usr/lib/libAutomator.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: Automator: failed to load framework from any known path\n")
	}
}
