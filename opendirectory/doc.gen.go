// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

// Package opendirectory provides Go bindings for the OpenDirectory framework.
//
// Authenticate users, and search for contact information in Open Directory
// and LDAP directories.
//
// # Classes
//
//   - [ODAttributeMap]
//   - [ODConfiguration]
//   - [ODMappings]
//   - [ODModuleEntry]
//   - [ODNode]: An object serves as a Cocoa wrapper for an Open Directory node.
//   - [ODNodeRef]: An Open Directory node type.
//   - [ODQuery]: An object serves as a Cocoa wrapper for an Open Directory query.
//   - [ODQueryRef]: An Open Directory query type.
//   - [ODRecord]: An object serves as a Cocoa wrapper for an Open Directory record.
//   - [ODRecordMap]
//   - [ODRecordRef]: An Open Directory record type.
//   - [ODSession]: An object serves as a Cocoa wrapper for an Open Directory session.
//   - [ODSessionRef]: An Open Directory session type.
//
// # Protocols
//
//   - [ODQueryDelegate]: The protocol defines methods for receiving results returned from an Open Directory query.
//
// # Type Aliases
//
//   - [ODContextRef]
//
// # Key Types
//
//   - [ODConfiguration]
//   - [ODRecord] - An [ODRecord] object serves as a Cocoa wrapper for an Open Directory record.
//   - [ODNode] - An [ODNode] object serves as a Cocoa wrapper for an Open Directory node.
//   - [ODSession] - An [ODSession] object serves as a Cocoa wrapper for an Open Directory session.
//   - [ODAttributeMap]
//   - [ODMappings]
//   - [ODModuleEntry]
//   - [ODQuery] - An [ODQuery] object serves as a Cocoa wrapper for an Open Directory query.
//   - [ODRecordMap]
package opendirectory

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the OpenDirectory library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/OpenDirectory.framework/OpenDirectory",
	"/usr/lib/libOpenDirectory.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: OpenDirectory: failed to load framework from any known path\n")
	}
}
