// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

// Package executionpolicy provides Go bindings for the ExecutionPolicy framework.
//
// Provide functionality so developer tools can manage execution policy
// exceptions.
//
// This framework provides the ability for developer tools to manage execution
// policy exceptions for software that’s built locally.
//
// # Classes
//
//   - [EPDeveloperTool]
//   - [EPExecutionPolicy]
//
// # Key Types
//
//   - [EPDeveloperTool]
//   - [EPExecutionPolicy]
package executionpolicy

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ExecutionPolicy library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ExecutionPolicy.framework/ExecutionPolicy",
	"/usr/lib/libExecutionPolicy.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: ExecutionPolicy: failed to load framework from any known path\n")
	}
}
