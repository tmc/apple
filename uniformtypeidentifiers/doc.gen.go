// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

// Package uniformtypeidentifiers provides Go bindings for the UniformTypeIdentifiers framework.
//
// Provide uniform type identifiers that describe file types for storage or transfer.
//
// The [Uniform Type Identifiers](<doc://com.apple.uniformtypeidentifiers/documentation/UniformTypeIdentifiers>) framework provides a collection of common types that map to MIME and file types. Use these types in your project to describe the file types in your app. These descriptions help the system properly handle file storage formats or in-memory data for transfer — for example, transferring data to or from the pasteboard. The identifier types can also identify other resources, such as directories, volumes, or packages.
//
// # Essentials
//
//   - Defining file and data types for your app: Declare uniform type identifiers to support your app’s proprietary data formats.
//   - System-declared uniform type identifiers: Common types that the system declares.
//
// # Uniform type identifiers
//
//   - UTType: A structure that represents a type of data to load, send, or receive.
//   - UTTagClass: A type that represents tag classes.
//   - UTTypeReference: An object that represents a type of data to load, send, or receive.
//
// # Key Types
//
//   - [UTType] - An object that represents a type of data to load, send, or receive.
//
// [UniformTypeIdentifiers Documentation]: https://developer.apple.com/documentation/UniformTypeIdentifiers
package uniformtypeidentifiers

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the UniformTypeIdentifiers library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/UniformTypeIdentifiers.framework/UniformTypeIdentifiers",
	"/usr/lib/libUniformTypeIdentifiers.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: UniformTypeIdentifiers: failed to load framework from any known path\n")
}
