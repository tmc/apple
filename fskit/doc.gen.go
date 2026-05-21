// Code generated from Apple documentation for FSKit. DO NOT EDIT.

// Package fskit provides Go bindings for the FSKit framework.
//
// Implement a file system that runs in user space.
//
// With FSKit, you can extend macOS by enabling access to new types of file
// systems. You do this by developing an FSKit module ([FSModule]), which you
// deliver as an app extension that runs in user space, and is compatible with
// Mac App Store distribution. FSKit connects your module to the system’s
// existing frameworks and tools, like Disk Arbitration, NetFS, and the
// `mount(8)` command.
//
// # File systems
//
//   - [FSFileSystem]: An abstract base class for implementing a full-featured file system.
//   - [FSUnaryFileSystem]: An abstract base class for implementing a minimal file system. ([FSUnaryFileSystemOperations])
//   - [FSFileSystemBase]: A protocol containing functionality supplied by FSKit to file system implementations.
//   - [FSFileName]: The name of a file, expressed as a data buffer.
//
// # Containers
//
//   - [FSContainerIdentifier]: A type that identifies a container.
//   - [FSContainerStatus]: A type that represents a container’s status. ([FSContainerState])
//
// # Resources
//
//   - [FSResource]: An abstract resource a file system uses to provide data for a volume.
//   - [FSBlockDeviceResource]: A resource that represents a block storage disk partition. ([FSMetadataRange])
//   - [FSPathURLResource]: A resource that represents a path in the system file space.
//   - [FSGenericURLResource]: A resource that represents an abstract URL.
//
// # Volumes
//
//   - [FSVolume]: A directory structure for files and folders. ([FSFileName], [FSVolumeKernelOffloadedIOOperations])
//
// # Items
//
//   - [FSItem]: A distinct object in a file hierarchy, such as a file, directory, symlink, socket, and more.
//
// # Maintenance and management
//
//   - [FSManageableResourceMaintenanceOperations]: Maintenance operations for a file system’s resources.
//
// # Operations
//
//   - [FSOperationID]: A unique identifier for an operation.
//
// # Tasks
//
//   - [FSTask]: A class that enables a file system module to pass log messages and completion notifications to clients.
//   - [FSTaskOptions]: A class that passes command options to a task, optionally providing security-scoped URLs.
//
// # Errors and logging
//
//   - [Fs_errorForCocoaError]: Creates an error object for the given Cocoa error code.
//   - [Fs_errorForMachError]: Creates an error object for the given Mach error code.
//   - [Fs_errorForPOSIXError]: Creates an error object for the given POSIX error code.
//   - [FSErrorCode]: A code that indicates a specific FSKit error.
//   - [FSKitErrorDomain]: An error domain for FSKit errors.
//
// # FSKit interactions
//
//   - [FSClient]: An interface for apps and daemons to interact with FSKit. ([FSModuleIdentity])
//
// # Utilities
//
//   - [FSKitVersionNumber]: Project version number for FSKit.
//   - [FSKitVersionString]: Project version string for FSKit.
//
// # Supporting types
//
//   - [FSBlockmapFlags]: Flags that describe the behavior of a blockmap operation.
//   - [FSCompleteIOFlags]: Flags that describe the behavior of an I/O completion operation.
//   - [FSEntityIdentifier]: A base type that identifies containers and volumes.
//   - [FSExtentPacker]: A type that directs the kernel to map space on disk to a specific file managed by this file system. ([FSExtentType])
//   - [FSExtentType]: An enumeration of types of extents.
//   - [FSMatchResult]: A type that represents the recognition and usability of a probed resource.
//   - [FSMetadataRange]: A range that describes contiguous metadata segments on disk.
//   - [FSProbeResult]: An object that represents the results of a specific probe. ([FSMatchResult])
//
// # Macros
//
//   - FSKIT_API_AVAILABILITY_V1
//   - FSKIT_API_UNAVAILABLE_V1
//   - FS_ALWAYS_EXPORT
//   - FS_EXPORT
//   - FS_EXPORT_INTERNAL
//   - FS_EXTERN
//   - FS_SUPPORTED_VISIBILITY
//   - FSKIT_API_AVAILABILITY_V2
//   - FSKIT_API_AVAILABILITY_V2_4
//
// # Key Types
//
//   - [FSItemAttributes] - Attributes of an item, such as size, creation and modification times, and user and group identifiers.
//   - [FSVolumeSupportedCapabilities] - A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.
//   - [FSBlockDeviceResource] - A resource that represents a block storage disk partition.
//   - [FSStatFSResult] - A type used to report a volume’s statistics.
//   - [FSFileName] - The name of a file, expressed as a data buffer.
//   - [FSProbeResult] - An object that represents the results of a specific probe.
//   - [FSContainerStatus] - A type that represents a container’s status.
//   - [FSEntityIdentifier] - A base type that identifies containers and volumes.
//   - [FSContainerIdentifier] - A type that identifies a container.
//   - [FSMetadataRange] - A range that describes contiguous metadata segments on disk.
package fskit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the FSKit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/FSKit.framework/FSKit",
	"/usr/lib/libFSKit.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: FSKit: failed to load framework from any known path\n")
}
