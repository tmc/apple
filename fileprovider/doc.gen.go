// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

// Package fileprovider provides Go bindings for the FileProvider framework.
//
// An extension other apps use to access files and folders managed by your app
// and synced with a remote storage.
//
// If your app focuses on providing and syncing user documents from remote
// storage, you can implement a File Provider extension to give users access
// to those documents when they’re using other apps. If you just need to
// share local documents, see Share files locally below.
//
// # Essentials
//
//   - [File Provider updates]: Learn about important changes to File Provider.
//
// # Extension types
//
//   - [Replicated File Provider extension]: Build a File Provider extension that syncs the local copies of your files with your remote storage. ([NSFileProviderReplicatedExtension], [NSFileProviderEnumerating], [NSFileProviderIncrementalContentFetching], [NSFileProviderPartialContentFetching], [NSFileProviderServicing])
//   - [Nonreplicated File Provider extension]: Build a File Provider extension that hosts and manages the user’s local files.
//
// # Extension management
//
//   - [NSFileProviderManager]: A manager object that you use to communicate with the file provider from either your app or your File Provider extension. ([NSFileProviderKnownFolders], [NSFileProviderKnownFolderLocations], [NSFileProviderKnownFolderSupporting], [NSFileProviderExternalVolumeHandling])
//
// # Provided items
//
//   - [NSFileProviderItem]: An item the File Provider extension manages.
//   - [NSFileProviderItem]: A protocol that defines the properties of an item managed by the File Provider extension. ([NSFileProviderContentPolicy], [NSFileProviderFileSystemFlags])
//   - [NSFileProviderItemIdentifier]: A unique identifier for an item managed by the File Provider extension.
//   - [NSFileProviderItemCapabilities]: An item’s capabilities, which define the actions that the user can perform in the document browser.
//   - [NSFileProviderTypeAndCreator]: A structure that contains the file type and file creator codes for an item.
//
// # Cloud search
//
//   - [NSFileProviderSearching]: A protocol you implement to support searching in your file provider. ([NSFileProviderStringSearchRequest], [NSFileProviderSearchEnumerator])
//
// # Domains
//
//   - [NSFileProviderDomain]: A File Provider extension’s domain. ([NSFileProviderDomainIdentifier], [NSFileProviderKnownFolders])
//
// # Errors
//
//   - [NSFileProviderErrorCode]: The error codes for the File Provider extension.
//   - [FileProviderErrorDomain]: The error domain for the File Provider extension.
//   - [FileProviderErrorItemKey]: The key for accessing information about sync-related errors.
//   - [FileProviderErrorNonExistentItemIdentifierKey]: The key for accessing the specified item’s identifier when the item doesn’t exist.
//
// # Data export
//
//   - [Exporting file provider metrics data]: Download and analyze usage, consistency, and error data.
//
// # Global variables and macros
//
//   - [Global variables and macros]
//
// # Variables
//
//   - [FileProviderUserInfoExperimentIDKey]: System interpreted user info key When setting a value to that user info on a domain, the system will ingest this value.
//
// # Type Aliases
//
//   - [NSFileProviderUserInfoKey]
//
// # Enumerations
//
//   - [NSFileProviderVolumeUnsupportedReason]: Constants that describe why an external volume might not be eligible for storing a domain.//
//
// # Key Types
//
//   - [NSFileProviderManager] - A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
//   - [NSFileProviderDomain] - A File Provider extension’s domain.
//   - [NSFileProviderItemVersion] - The version of the item’s content and its metadata.
//   - [NSFileProviderKnownFolderLocations] - A class for working with known-folder locations.
//   - [NSFileProviderRequest] - An object that provides information about the application requesting data from the File Provider extension.
//   - [NSFileProviderKnownFolderLocation]
//   - [NSFileProviderStringSearchRequest] - A type that contains details of a string-based search request.
//   - [NSFileProviderDomainVersion] - An opaque object that identifies a specific version of a domain.
//
// [Exporting file provider metrics data]: https://developer.apple.com/documentation/fileprovider/exporting-file-provider-metrics-data
// [File Provider updates]: https://developer.apple.com/documentation/Updates/FileProvider
// [Global variables and macros]: https://developer.apple.com/documentation/fileprovider/global-variables-and-macros
// [Nonreplicated File Provider extension]: https://developer.apple.com/documentation/fileprovider/nonreplicated-file-provider-extension
// [Replicated File Provider extension]: https://developer.apple.com/documentation/fileprovider/replicated-file-provider-extension
package fileprovider

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the FileProvider library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/FileProvider.framework/FileProvider",
	"/usr/lib/libFileProvider.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: FileProvider: failed to load framework from any known path\n")
	}
}
