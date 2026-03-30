// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FileManager] class.
var (
	_FileManagerClass     FileManagerClass
	_FileManagerClassOnce sync.Once
)

func getFileManagerClass() FileManagerClass {
	_FileManagerClassOnce.Do(func() {
		_FileManagerClass = FileManagerClass{class: objc.GetClass("NSFileManager")}
	})
	return _FileManagerClass
}

// GetFileManagerClass returns the class object for NSFileManager.
func GetFileManagerClass() FileManagerClass {
	return getFileManagerClass()
}

type FileManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FileManagerClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FileManagerClass) Alloc() FileManager {
	rv := objc.Send[FileManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A convenient interface to the contents of the file system, and the primary
// means of interacting with it.
//
// # Overview
//
// A file manager object lets you examine the contents of the file system and
// make changes to it. The [NSFileManager] class provides convenient access to
// a shared file manager object that is suitable for most types of
// file-related manipulations. A file manager object is typically your primary
// mode of interaction with the file system. You use it to locate, create,
// copy, and move files and directories. You also use it to get information
// about a file or directory or change some of its attributes.
//
// When specifying the location of files, you can use either [NSURL] or
// [NSString] objects. The use of the [NSURL] class is generally preferred for
// specifying file-system items because URLs can convert path information to a
// more efficient representation internally. You can also obtain a bookmark
// from an [NSURL] object, which is similar to an alias and offers a more sure
// way of locating the file or directory later.
//
// If you are moving, copying, linking, or removing files or directories, you
// can use a delegate in conjunction with a file manager object to manage
// those operations. The delegate’s role is to affirm the operation and to
// decide whether to proceed when errors occur. In macOS 10.7 and later, the
// delegate must conform to the [NSFileManagerDelegate] protocol.
//
// In iOS 5.0 and later and in macOS 10.7 and later, [NSFileManager] includes
// methods for managing items stored in iCloud. Files and directories tagged
// for cloud storage are synced to iCloud so that they can be made available
// to the user’s iOS devices and Macintosh computers. Changes to an item in
// one location are propagated to all other locations to ensure the items stay
// in sync.
//
// # Sync control
//
// A [package] is a directory that the system presents as a single file to the
// person using the device. Apps with documents that contain multiple files
// can use packages to manage contents like media assets. In iOS 26 and macOS
// 26 and later, [NSFileManager] introduces methods for controlling how a file
// provider syncs these items. By pausing sync when your app opens a package
// and resuming when it closes, your app can prevent the file provider from
// changing the contents of the package in unexpected ways, which potentially
// leaves the document in an inconsistent state. You can also use this pause
// and resume API on regular “flat” files.
//
// # Threading considerations
//
// The methods of the shared [NSFileManager] object can be called from
// multiple threads safely. However, if you use a delegate to receive
// notifications about the status of move, copy, remove, and link operations,
// you should create a unique instance of the file manager object, assign your
// delegate to that object, and use that file manager to initiate your
// operations.
//
// # Accessing user directories
//
//   - [FileManager.HomeDirectoryForCurrentUser]: The home directory for the current user.
//   - [FileManager.HomeDirectoryForUser]: Returns the home directory for the specified user.
//   - [FileManager.TemporaryDirectory]: The temporary directory for the current user.
//
// # Locating system directories
//
//   - [FileManager.URLForDirectoryInDomainAppropriateForURLCreateError]: Locates and optionally creates the specified common directory in a domain.
//   - [FileManager.URLsForDirectoryInDomains]: Returns an array of URLs for the specified common directory in the requested domains.
//
// # Locating application group container directories
//
//   - [FileManager.ContainerURLForSecurityApplicationGroupIdentifier]: Returns the container directory associated with the specified security application group identifier.
//
// # Discovering directory contents
//
//   - [FileManager.ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError]: Performs a shallow search of the specified directory and returns URLs for the contained items.
//   - [FileManager.ContentsOfDirectoryAtPathError]: Performs a shallow search of the specified directory and returns the paths of any contained items.
//   - [FileManager.EnumeratorAtPath]: Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
//   - [FileManager.MountedVolumeURLsIncludingResourceValuesForKeysOptions]: Returns an array of URLs that identify the mounted volumes available on the device.
//   - [FileManager.SubpathsOfDirectoryAtPathError]: Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
//   - [FileManager.SubpathsAtPath]: Returns an array of strings identifying the paths for all items in the specified directory.
//
// # Creating and deleting items
//
//   - [FileManager.CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError]: Creates a directory with the given attributes at the specified URL.
//   - [FileManager.CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError]: Creates a directory with given attributes at the specified path.
//   - [FileManager.CreateFileAtPathContentsAttributes]: Creates a file with the specified content and attributes at the given location.
//   - [FileManager.RemoveItemAtURLError]: Removes the file or directory at the specified URL.
//   - [FileManager.RemoveItemAtPathError]: Removes the file or directory at the specified path.
//   - [FileManager.TrashItemAtURLResultingItemURLError]: Moves an item to the trash.
//
// # Replacing items
//
//   - [FileManager.ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError]: Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
//
// # Moving and copying items
//
//   - [FileManager.CopyItemAtURLToURLError]: Copies the file at the specified URL to a new location synchronously.
//   - [FileManager.CopyItemAtPathToPathError]: Copies the item at the specified path to a new location synchronously.
//   - [FileManager.MoveItemAtURLToURLError]: Moves the file or directory at the specified URL to a new location synchronously.
//   - [FileManager.MoveItemAtPathToPathError]: Moves the file or directory at the specified path to a new location synchronously.
//
// # Managing iCloud-based items
//
//   - [FileManager.UbiquityIdentityToken]: An opaque token that represents the current user’s iCloud Drive Documents identity.
//   - [FileManager.URLForUbiquityContainerIdentifier]: Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container.
//   - [FileManager.IsUbiquitousItemAtURL]: Returns a Boolean indicating whether the item is targeted for storage in iCloud.
//   - [FileManager.SetUbiquitousItemAtURLDestinationURLError]: Indicates whether the item at the specified URL should be stored in iCloud.
//   - [FileManager.StartDownloadingUbiquitousItemAtURLError]: Starts downloading (if necessary) the specified item to the local system.
//   - [FileManager.EvictUbiquitousItemAtURLError]: Removes the local copy of the specified item that’s stored in iCloud.
//   - [FileManager.URLForPublishingUbiquitousItemAtURLExpirationDateError]: Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud.
//
// # Controlling file provider synchronization
//
//   - [FileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]: Asynchronously pauses sync of an item at the given URL.
//   - [FileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]: Asynchronously resumes the sync on a paused item using the given resume behavior.
//   - [FileManager.FetchLatestRemoteVersionOfItemAtURLCompletionHandler]: Asynchronously fetches the latest remote version of a given item from the server.
//   - [FileManager.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler]: Asynchronously uploads the local version of the item using the provided conflict resolution policy.
//
// # Creating symbolic and hard links
//
//   - [FileManager.CreateSymbolicLinkAtURLWithDestinationURLError]: Creates a symbolic link at the specified URL that points to an item at the given URL.
//   - [FileManager.CreateSymbolicLinkAtPathWithDestinationPathError]: Creates a symbolic link that points to the specified destination.
//   - [FileManager.LinkItemAtURLToURLError]: Creates a hard link between the items at the specified URLs.
//   - [FileManager.LinkItemAtPathToPathError]: Creates a hard link between the items at the specified paths.
//   - [FileManager.DestinationOfSymbolicLinkAtPathError]: Returns the path of the item pointed to by a symbolic link.
//
// # Determining access to files
//
//   - [FileManager.FileExistsAtPath]: Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//   - [FileManager.FileExistsAtPathIsDirectory]: Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//   - [FileManager.IsReadableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
//   - [FileManager.IsWritableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
//   - [FileManager.IsExecutableFileAtPath]: Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
//   - [FileManager.IsDeletableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
//
// # Getting and setting attributes
//
//   - [FileManager.ComponentsToDisplayForPath]: Returns an array of strings representing the user-visible components of a given path.
//   - [FileManager.DisplayNameAtPath]: Returns the display name of the file or directory at a specified path.
//   - [FileManager.AttributesOfItemAtPathError]: Returns the attributes of the item at a given path.
//   - [FileManager.AttributesOfFileSystemForPathError]: Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
//   - [FileManager.SetAttributesOfItemAtPathError]: Sets the attributes of the specified file or directory.
//
// # Getting and comparing file contents
//
//   - [FileManager.ContentsAtPath]: Returns the contents of the file at the specified path.
//   - [FileManager.ContentsEqualAtPathAndPath]: Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents.
//
// # Getting the relationship between items
//
//   - [FileManager.GetRelationshipOfDirectoryAtURLToItemAtURLError]: Determines the type of relationship that exists between a directory and an item.
//   - [FileManager.GetRelationshipOfDirectoryInDomainToItemAtURLError]: Determines the type of relationship that exists between a system directory and the specified item.
//
// # Converting file paths to strings
//
//   - [FileManager.FileSystemRepresentationWithPath]: Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
//   - [FileManager.StringWithFileSystemRepresentationLength]: Returns an [NSString](<doc://com.apple.foundation/documentation/Foundation/NSString>) object whose contents are derived from the specified C-string path.
//
// # Managing the delegate
//
//   - [FileManager.Delegate]: The delegate of the file manager object.
//   - [FileManager.SetDelegate]
//
// # Managing the current directory
//
//   - [FileManager.ChangeCurrentDirectoryPath]: Changes the path of the current working directory to the specified path.
//   - [FileManager.CurrentDirectoryPath]: The path to the program’s current directory.
//
// # Unmounting volumes
//
//   - [FileManager.UnmountVolumeAtURLOptionsCompletionHandler]: Starts the process of unmounting the specified volume.
//   - [FileManager.NSFileManagerUnmountDissentingProcessIdentifierErrorKey]: The process identifier of the process that prevented a volume from unmounting.
//
// # Determining resource fork support
//
//   - [FileManager.NSFoundationVersionWithFileManagerResourceForkSupport]: The version of the Foundation framework in which
//   - [FileManager.SetNSFoundationVersionWithFileManagerResourceForkSupport]
//
// # Notifications
//
//   - [FileManager.NSUbiquityIdentityDidChange]: Sent after the iCloud (“ubiquity”) identity has changed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager
//
// [package]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/DocumentPackages/DocumentPackages.html#//apple_ref/doc/uid/10000123i-CH106-SW1
type FileManager struct {
	objectivec.Object
}

// FileManagerFromID constructs a [FileManager] from an objc.ID.
//
// A convenient interface to the contents of the file system, and the primary
// means of interacting with it.
func FileManagerFromID(id objc.ID) FileManager {
	return NSFileManager{objectivec.Object{ID: id}}
}

// NSFileManagerFromID is an alias for [FileManagerFromID] for cross-framework compatibility.
func NSFileManagerFromID(id objc.ID) FileManager { return FileManagerFromID(id) }

// NOTE: FileManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FileManager] class.
//
// # Accessing user directories
//
//   - [IFileManager.HomeDirectoryForCurrentUser]: The home directory for the current user.
//   - [IFileManager.HomeDirectoryForUser]: Returns the home directory for the specified user.
//   - [IFileManager.TemporaryDirectory]: The temporary directory for the current user.
//
// # Locating system directories
//
//   - [IFileManager.URLForDirectoryInDomainAppropriateForURLCreateError]: Locates and optionally creates the specified common directory in a domain.
//   - [IFileManager.URLsForDirectoryInDomains]: Returns an array of URLs for the specified common directory in the requested domains.
//
// # Locating application group container directories
//
//   - [IFileManager.ContainerURLForSecurityApplicationGroupIdentifier]: Returns the container directory associated with the specified security application group identifier.
//
// # Discovering directory contents
//
//   - [IFileManager.ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError]: Performs a shallow search of the specified directory and returns URLs for the contained items.
//   - [IFileManager.ContentsOfDirectoryAtPathError]: Performs a shallow search of the specified directory and returns the paths of any contained items.
//   - [IFileManager.EnumeratorAtPath]: Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
//   - [IFileManager.MountedVolumeURLsIncludingResourceValuesForKeysOptions]: Returns an array of URLs that identify the mounted volumes available on the device.
//   - [IFileManager.SubpathsOfDirectoryAtPathError]: Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
//   - [IFileManager.SubpathsAtPath]: Returns an array of strings identifying the paths for all items in the specified directory.
//
// # Creating and deleting items
//
//   - [IFileManager.CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError]: Creates a directory with the given attributes at the specified URL.
//   - [IFileManager.CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError]: Creates a directory with given attributes at the specified path.
//   - [IFileManager.CreateFileAtPathContentsAttributes]: Creates a file with the specified content and attributes at the given location.
//   - [IFileManager.RemoveItemAtURLError]: Removes the file or directory at the specified URL.
//   - [IFileManager.RemoveItemAtPathError]: Removes the file or directory at the specified path.
//   - [IFileManager.TrashItemAtURLResultingItemURLError]: Moves an item to the trash.
//
// # Replacing items
//
//   - [IFileManager.ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError]: Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
//
// # Moving and copying items
//
//   - [IFileManager.CopyItemAtURLToURLError]: Copies the file at the specified URL to a new location synchronously.
//   - [IFileManager.CopyItemAtPathToPathError]: Copies the item at the specified path to a new location synchronously.
//   - [IFileManager.MoveItemAtURLToURLError]: Moves the file or directory at the specified URL to a new location synchronously.
//   - [IFileManager.MoveItemAtPathToPathError]: Moves the file or directory at the specified path to a new location synchronously.
//
// # Managing iCloud-based items
//
//   - [IFileManager.UbiquityIdentityToken]: An opaque token that represents the current user’s iCloud Drive Documents identity.
//   - [IFileManager.URLForUbiquityContainerIdentifier]: Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container.
//   - [IFileManager.IsUbiquitousItemAtURL]: Returns a Boolean indicating whether the item is targeted for storage in iCloud.
//   - [IFileManager.SetUbiquitousItemAtURLDestinationURLError]: Indicates whether the item at the specified URL should be stored in iCloud.
//   - [IFileManager.StartDownloadingUbiquitousItemAtURLError]: Starts downloading (if necessary) the specified item to the local system.
//   - [IFileManager.EvictUbiquitousItemAtURLError]: Removes the local copy of the specified item that’s stored in iCloud.
//   - [IFileManager.URLForPublishingUbiquitousItemAtURLExpirationDateError]: Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud.
//
// # Controlling file provider synchronization
//
//   - [IFileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]: Asynchronously pauses sync of an item at the given URL.
//   - [IFileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]: Asynchronously resumes the sync on a paused item using the given resume behavior.
//   - [IFileManager.FetchLatestRemoteVersionOfItemAtURLCompletionHandler]: Asynchronously fetches the latest remote version of a given item from the server.
//   - [IFileManager.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler]: Asynchronously uploads the local version of the item using the provided conflict resolution policy.
//
// # Creating symbolic and hard links
//
//   - [IFileManager.CreateSymbolicLinkAtURLWithDestinationURLError]: Creates a symbolic link at the specified URL that points to an item at the given URL.
//   - [IFileManager.CreateSymbolicLinkAtPathWithDestinationPathError]: Creates a symbolic link that points to the specified destination.
//   - [IFileManager.LinkItemAtURLToURLError]: Creates a hard link between the items at the specified URLs.
//   - [IFileManager.LinkItemAtPathToPathError]: Creates a hard link between the items at the specified paths.
//   - [IFileManager.DestinationOfSymbolicLinkAtPathError]: Returns the path of the item pointed to by a symbolic link.
//
// # Determining access to files
//
//   - [IFileManager.FileExistsAtPath]: Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//   - [IFileManager.FileExistsAtPathIsDirectory]: Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//   - [IFileManager.IsReadableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
//   - [IFileManager.IsWritableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
//   - [IFileManager.IsExecutableFileAtPath]: Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
//   - [IFileManager.IsDeletableFileAtPath]: Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
//
// # Getting and setting attributes
//
//   - [IFileManager.ComponentsToDisplayForPath]: Returns an array of strings representing the user-visible components of a given path.
//   - [IFileManager.DisplayNameAtPath]: Returns the display name of the file or directory at a specified path.
//   - [IFileManager.AttributesOfItemAtPathError]: Returns the attributes of the item at a given path.
//   - [IFileManager.AttributesOfFileSystemForPathError]: Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
//   - [IFileManager.SetAttributesOfItemAtPathError]: Sets the attributes of the specified file or directory.
//
// # Getting and comparing file contents
//
//   - [IFileManager.ContentsAtPath]: Returns the contents of the file at the specified path.
//   - [IFileManager.ContentsEqualAtPathAndPath]: Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents.
//
// # Getting the relationship between items
//
//   - [IFileManager.GetRelationshipOfDirectoryAtURLToItemAtURLError]: Determines the type of relationship that exists between a directory and an item.
//   - [IFileManager.GetRelationshipOfDirectoryInDomainToItemAtURLError]: Determines the type of relationship that exists between a system directory and the specified item.
//
// # Converting file paths to strings
//
//   - [IFileManager.FileSystemRepresentationWithPath]: Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
//   - [IFileManager.StringWithFileSystemRepresentationLength]: Returns an [NSString](<doc://com.apple.foundation/documentation/Foundation/NSString>) object whose contents are derived from the specified C-string path.
//
// # Managing the delegate
//
//   - [IFileManager.Delegate]: The delegate of the file manager object.
//   - [IFileManager.SetDelegate]
//
// # Managing the current directory
//
//   - [IFileManager.ChangeCurrentDirectoryPath]: Changes the path of the current working directory to the specified path.
//   - [IFileManager.CurrentDirectoryPath]: The path to the program’s current directory.
//
// # Unmounting volumes
//
//   - [IFileManager.UnmountVolumeAtURLOptionsCompletionHandler]: Starts the process of unmounting the specified volume.
//   - [IFileManager.NSFileManagerUnmountDissentingProcessIdentifierErrorKey]: The process identifier of the process that prevented a volume from unmounting.
//
// # Determining resource fork support
//
//   - [IFileManager.NSFoundationVersionWithFileManagerResourceForkSupport]: The version of the Foundation framework in which
//   - [IFileManager.SetNSFoundationVersionWithFileManagerResourceForkSupport]
//
// # Notifications
//
//   - [IFileManager.NSUbiquityIdentityDidChange]: Sent after the iCloud (“ubiquity”) identity has changed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager
type IFileManager interface {
	objectivec.IObject

	// Topic: Accessing user directories

	// The home directory for the current user.
	HomeDirectoryForCurrentUser() INSURL
	// Returns the home directory for the specified user.
	HomeDirectoryForUser(userName string) INSURL
	// The temporary directory for the current user.
	TemporaryDirectory() INSURL

	// Topic: Locating system directories

	// Locates and optionally creates the specified common directory in a domain.
	URLForDirectoryInDomainAppropriateForURLCreateError(directory NSSearchPathDirectory, domain NSSearchPathDomainMask, url INSURL, shouldCreate bool) (INSURL, error)
	// Returns an array of URLs for the specified common directory in the requested domains.
	URLsForDirectoryInDomains(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask) []NSURL

	// Topic: Locating application group container directories

	// Returns the container directory associated with the specified security application group identifier.
	ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) INSURL

	// Topic: Discovering directory contents

	// Performs a shallow search of the specified directory and returns URLs for the contained items.
	ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url INSURL, keys []string, mask NSDirectoryEnumerationOptions) ([]NSURL, error)
	// Performs a shallow search of the specified directory and returns the paths of any contained items.
	ContentsOfDirectoryAtPathError(path string) ([]string, error)
	// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
	EnumeratorAtPath(path string) INSDirectoryEnumerator
	// Returns an array of URLs that identify the mounted volumes available on the device.
	MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string, options NSVolumeEnumerationOptions) []NSURL
	// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
	SubpathsOfDirectoryAtPathError(path string) ([]string, error)
	// Returns an array of strings identifying the paths for all items in the specified directory.
	SubpathsAtPath(path string) []string

	// Topic: Creating and deleting items

	// Creates a directory with the given attributes at the specified URL.
	CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url INSURL, createIntermediates bool, attributes INSDictionary) (bool, error)
	// Creates a directory with given attributes at the specified path.
	CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes INSDictionary) (bool, error)
	// Creates a file with the specified content and attributes at the given location.
	CreateFileAtPathContentsAttributes(path string, data INSData, attr INSDictionary) bool
	// Removes the file or directory at the specified URL.
	RemoveItemAtURLError(URL INSURL) (bool, error)
	// Removes the file or directory at the specified path.
	RemoveItemAtPathError(path string) (bool, error)
	// Moves an item to the trash.
	TrashItemAtURLResultingItemURLError(url INSURL, outResultingURL INSURL) (bool, error)

	// Topic: Replacing items

	// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
	ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL INSURL, newItemURL INSURL, backupItemName string, options NSFileManagerItemReplacementOptions, resultingURL INSURL) (bool, error)

	// Topic: Moving and copying items

	// Copies the file at the specified URL to a new location synchronously.
	CopyItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error)
	// Copies the item at the specified path to a new location synchronously.
	CopyItemAtPathToPathError(srcPath string, dstPath string) (bool, error)
	// Moves the file or directory at the specified URL to a new location synchronously.
	MoveItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error)
	// Moves the file or directory at the specified path to a new location synchronously.
	MoveItemAtPathToPathError(srcPath string, dstPath string) (bool, error)

	// Topic: Managing iCloud-based items

	// An opaque token that represents the current user’s iCloud Drive Documents identity.
	UbiquityIdentityToken() objectivec.IObject
	// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container.
	URLForUbiquityContainerIdentifier(containerIdentifier string) INSURL
	// Returns a Boolean indicating whether the item is targeted for storage in iCloud.
	IsUbiquitousItemAtURL(url INSURL) bool
	// Indicates whether the item at the specified URL should be stored in iCloud.
	SetUbiquitousItemAtURLDestinationURLError(flag bool, url INSURL, destinationURL INSURL) (bool, error)
	// Starts downloading (if necessary) the specified item to the local system.
	StartDownloadingUbiquitousItemAtURLError(url INSURL) (bool, error)
	// Removes the local copy of the specified item that’s stored in iCloud.
	EvictUbiquitousItemAtURLError(url INSURL) (bool, error)
	// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud.
	URLForPublishingUbiquitousItemAtURLExpirationDateError(url INSURL, outDate INSDate) (INSURL, error)

	// Topic: Controlling file provider synchronization

	// Asynchronously pauses sync of an item at the given URL.
	PauseSyncForUbiquitousItemAtURLCompletionHandler(url INSURL, completionHandler ErrorHandler)
	// Asynchronously resumes the sync on a paused item using the given resume behavior.
	ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url INSURL, behavior NSFileManagerResumeSyncBehavior, completionHandler ErrorHandler)
	// Asynchronously fetches the latest remote version of a given item from the server.
	FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url INSURL, completionHandler FileVersionErrorHandler)
	// Asynchronously uploads the local version of the item using the provided conflict resolution policy.
	UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url INSURL, conflictResolutionPolicy NSFileManagerUploadLocalVersionConflictPolicy, completionHandler FileVersionErrorHandler)

	// Topic: Creating symbolic and hard links

	// Creates a symbolic link at the specified URL that points to an item at the given URL.
	CreateSymbolicLinkAtURLWithDestinationURLError(url INSURL, destURL INSURL) (bool, error)
	// Creates a symbolic link that points to the specified destination.
	CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string) (bool, error)
	// Creates a hard link between the items at the specified URLs.
	LinkItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error)
	// Creates a hard link between the items at the specified paths.
	LinkItemAtPathToPathError(srcPath string, dstPath string) (bool, error)
	// Returns the path of the item pointed to by a symbolic link.
	DestinationOfSymbolicLinkAtPathError(path string) (string, error)

	// Topic: Determining access to files

	// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
	FileExistsAtPath(path string) bool
	// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
	FileExistsAtPathIsDirectory(path string) (bool, bool)
	// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
	IsReadableFileAtPath(path string) bool
	// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
	IsWritableFileAtPath(path string) bool
	// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
	IsExecutableFileAtPath(path string) bool
	// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
	IsDeletableFileAtPath(path string) bool

	// Topic: Getting and setting attributes

	// Returns an array of strings representing the user-visible components of a given path.
	ComponentsToDisplayForPath(path string) []string
	// Returns the display name of the file or directory at a specified path.
	DisplayNameAtPath(path string) string
	// Returns the attributes of the item at a given path.
	AttributesOfItemAtPathError(path string) (INSDictionary, error)
	// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
	AttributesOfFileSystemForPathError(path string) (INSDictionary, error)
	// Sets the attributes of the specified file or directory.
	SetAttributesOfItemAtPathError(attributes INSDictionary, path string) (bool, error)

	// Topic: Getting and comparing file contents

	// Returns the contents of the file at the specified path.
	ContentsAtPath(path string) INSData
	// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents.
	ContentsEqualAtPathAndPath(path1 string, path2 string) bool

	// Topic: Getting the relationship between items

	// Determines the type of relationship that exists between a directory and an item.
	GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship NSURLRelationship, directoryURL INSURL, otherURL INSURL) (bool, error)
	// Determines the type of relationship that exists between a system directory and the specified item.
	GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship NSURLRelationship, directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, url INSURL) (bool, error)

	// Topic: Converting file paths to strings

	// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
	FileSystemRepresentationWithPath(path string) string
	// Returns an [NSString](<doc://com.apple.foundation/documentation/Foundation/NSString>) object whose contents are derived from the specified C-string path.
	StringWithFileSystemRepresentationLength(str string, len_ uint) string

	// Topic: Managing the delegate

	// The delegate of the file manager object.
	Delegate() NSFileManagerDelegate
	SetDelegate(value NSFileManagerDelegate)

	// Topic: Managing the current directory

	// Changes the path of the current working directory to the specified path.
	ChangeCurrentDirectoryPath(path string) bool
	// The path to the program’s current directory.
	CurrentDirectoryPath() string

	// Topic: Unmounting volumes

	// Starts the process of unmounting the specified volume.
	UnmountVolumeAtURLOptionsCompletionHandler(url INSURL, mask NSFileManagerUnmountOptions, completionHandler ErrorHandler)
	// The process identifier of the process that prevented a volume from unmounting.
	NSFileManagerUnmountDissentingProcessIdentifierErrorKey() string

	// Topic: Determining resource fork support

	// The version of the Foundation framework in which
	NSFoundationVersionWithFileManagerResourceForkSupport() objectivec.IObject
	SetNSFoundationVersionWithFileManagerResourceForkSupport(value objectivec.IObject)

	// Topic: Notifications

	// Sent after the iCloud (“ubiquity”) identity has changed.
	NSUbiquityIdentityDidChange() NSNotificationName

	// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL.
	EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url INSURL, keys []string, mask NSDirectoryEnumerationOptions, handler URLErrorHandler) INSDirectoryEnumerator
}

// Init initializes the instance.
func (f FileManager) Init() FileManager {
	rv := objc.Send[FileManager](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FileManager) Autorelease() FileManager {
	rv := objc.Send[FileManager](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileManager creates a new FileManager instance.
func NewFileManager() FileManager {
	class := getFileManagerClass()
	rv := objc.Send[FileManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a file manager object that is authorized to perform privileged
// file system operations.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
// authorization is a [appkit.NSWorkspaceAuthorization].
func NewFileManagerWithAuthorization(authorization objectivec.IObject) FileManager {
	rv := objc.Send[objc.ID](objc.ID(getFileManagerClass().class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return FileManagerFromID(rv)
}

// Returns the home directory for the specified user.
//
// userName: The username of the owner of the desired home directory.
//
// # Return Value
//
// A URL object containing the location of the specified user’s home
// directory, or `nil` if no such user exists or the user’s home directory
// is not available.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f FileManager) HomeDirectoryForUser(userName string) INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("homeDirectoryForUser:"), objc.String(userName))
	return NSURLFromID(rv)
}

// Locates and optionally creates the specified common directory in a domain.
//
// directory: The search path directory. The supported values are described in
// [FileManager.SearchPathDirectory].
//
// domain: The file system domain to search. The value for this parameter is one of
// the constants described in [FileManager.SearchPathDomainMask]. You should
// specify only one domain for your search and you may not specify the
// [NSAllDomainsMask] constant for this parameter.
//
// url: The file URL used to determine the location of the returned URL. Only the
// volume of this parameter is used.
//
// This parameter is ignored unless the `directory` parameter contains the
// value [NSItemReplacementDirectory] and the `domain` parameter contains the
// value [NSUserDomainMask].
//
// shouldCreate: Whether to create the directory if it does not already exist.
//
// When creating a temporary directory, this parameter is ignored and the
// directory is always created.
//
// # Return Value
//
// The [NSURL] for the requested directory. When using Objective-C, if an
// error occurs, this method returns `nil` and assigns an appropriate error
// object to the `error` parameter.
//
// # Discussion
//
// You typically use this method to locate one of the standard system
// directories, such as the [Documents], `Application Support` or [Caches]
// directories. After locating (or creating) the desired directory, this
// method returns the URL for that directory. If more than one appropriate
// directory exists in the specified domain, this method returns only the
// first one it finds.
//
// You can use this method to create a new temporary directory. To do so,
// specify [NSItemReplacementDirectory] for the `directory` parameter,
// [NSUserDomainMask] for the `domain` parameter, and a URL for the `url`
// parameter which determines the volume of the returned URL.
//
// For example, the following code results in a new temporary directory with a
// path in the form of
// `/private/var/folders/d0/h37cw8ns3h1bfr_2gnwq2yyc0000gn/T/TemporaryItems/Untitled/`:
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/url(for:in:appropriateFor:create:)
//
// [FileManager.SearchPathDirectory]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
// [FileManager.SearchPathDomainMask]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
func (f FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory NSSearchPathDirectory, domain NSSearchPathDomainMask, url INSURL, shouldCreate bool) (INSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URLForDirectory:inDomain:appropriateForURL:create:error:"), directory, domain, url, shouldCreate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}

// Returns an array of URLs for the specified common directory in the
// requested domains.
//
// directory: The search path directory. The supported values are described in
// [FileManager.SearchPathDirectory].
//
// domainMask: The file system domain to search. The value for this parameter is one or
// more of the constants described in [FileManager.SearchPathDomainMask].
//
// # Return Value
//
// An array of [NSURL] objects identifying the requested directories. The
// directories are ordered according to the order of the domain mask
// constants, with items in the user domain first and items in the system
// domain last.
//
// # Discussion
//
// This method is intended to locate known and common directories in the
// system. For example, setting the directory to [NSApplicationDirectory],
// will return the Applications directories in the requested domain. There are
// a number of common directories available in the
// [FileManager.SearchPathDirectory], including: [NSDesktopDirectory],
// [NSApplicationSupportDirectory], and many more.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/urls(for:in:)
//
// [FileManager.SearchPathDirectory]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
// [FileManager.SearchPathDomainMask]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
//
// [FileManager.SearchPathDirectory]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
func (f FileManager) URLsForDirectoryInDomains(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask) []NSURL {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	})
}

// Returns the container directory associated with the specified security
// application group identifier.
//
// groupIdentifier: A string that names the group whose shared directory you want to obtain.
// This input should exactly match one of the strings in the app’s [App
// Groups Entitlement].
//
// # Return Value
//
// A URL indicating the location of the group’s shared directory in the file
// system. In iOS, the value is `nil` when the group identifier is invalid. In
// macOS, a URL of the expected form is always returned, even if the app group
// is invalid, so be sure to test that you can access the underlying directory
// before attempting to use it.
//
// # Discussion
//
// Sandboxed apps in macOS and all apps in iOS that need to share files with
// other apps from the same developer on a given device use the [App Groups
// Entitlement] to join one or more application groups. The entitlement
// consists of an array of group identifier strings that indicate the groups
// to which the app belongs, as described in [Adding an App to an App Group]
// in [Entitlement Key Reference].
//
// You use one of these group identifier strings to locate the corresponding
// group’s shared directory. When you call
// [ContainerURLForSecurityApplicationGroupIdentifier] with one of your
// app’s group identifiers, the method returns an [NSURL] instance
// specifying the location in the file system of that group’s shared
// directory. The behavior of application groups differs between macOS and
// iOS.
//
// # App Groups in macOS
//
// For a sandboxed app in macOS, the group directory is located at
// `~/Library/Group Containers/`, where the application group identifier
// begins with the developer’s team identifier followed by a dot, followed
// by the specific group name. The system creates this directory automatically
// the first time your app needs it and never removes it.
//
// The system also creates the `Library/Application Support`,
// `Library/Caches`, and `Library/Preferences` subdirectories inside the group
// directory the first time you use it. You are free to add or remove
// subdirectories as you see fit, but you are encouraged to use these
// standardized locations as you would in the app’s usual container.
//
// If you call the method with an invalid group identifier, namely one for
// which you do not have an entitlement, the method still returns a URL of the
// expected form, but the corresponding group directory does not actually
// exist, nor can your sandboxed app create it. Therefore be sure to test that
// you can successfully access the returned URL before using it.
//
// # App Groups in iOS
//
// In iOS, the group identifier starts with the word `group` and a dot,
// followed by the group name. However, the system makes no guarantee about
// the group directory’s name or location in the file system. Indeed, the
// directory is accessible only with the file URL returned by this method. As
// in macOS, the system creates the directory when you need it. Unlike in
// macOS, when all the apps in a given app group are removed from the device,
// the system detects this condition and removes the corresponding group
// directory as well.
//
// The system creates only the `Library/Caches` subdirectory automatically,
// but you can create others yourself if you need them. You are free to use
// the group directory as you see fit, but take care to coordinate its
// structure among all the group’s apps.
//
// If you call the method with an invalid group identifier in iOS, the method
// returns a `nil` value.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
//
// [App Groups Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.application-groups
// [Adding an App to an App Group]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/EntitlementKeyReference/Chapters/EnablingAppSandbox.html#//apple_ref/doc/uid/TP40011195-CH4-SW19
// [Entitlement Key Reference]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/EntitlementKeyReference/Chapters/AboutEntitlements.html#//apple_ref/doc/uid/TP40011195
//
// [App Groups Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.application-groups
func (f FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("containerURLForSecurityApplicationGroupIdentifier:"), objc.String(groupIdentifier))
	return NSURLFromID(rv)
}

// Performs a shallow search of the specified directory and returns URLs for
// the contained items.
//
// url: The URL for the directory whose contents you want to enumerate.
//
// keys: An array of keys that identify the file properties that you want
// pre-fetched for each item in the directory. For each returned URL, the
// specified properties are fetched and cached in the [NSURL] object. For a
// list of keys you can specify, see [Common File System Resource Keys].
//
// If you want directory contents to have no pre-fetched file properties, pass
// an empty array to this parameter. If you want directory contents to have
// default set of pre-fetched file properties, pass `nil` to this parameter.
//
// mask: Options for the enumeration. Because this method performs only shallow
// enumerations, options that prevent descending into subdirectories or
// packages are not allowed; the only supported option is
// [NSDirectoryEnumerationSkipsHiddenFiles].
//
// # Return Value
//
// An array of [NSURL] objects, each of which identifies a file, directory, or
// symbolic link contained in `url`. If the directory contains no entries,
// this method returns an empty array. When using Objective-C, if an error
// occurs, this method returns `nil` and assigns an appropriate error object
// to the `error` parameter.
//
// # Discussion
//
// This method performs a shallow search of the directory and therefore does
// not traverse symbolic links or return the contents of any subdirectories.
// This method also does not return URLs for the current directory
// (”`.`”), parent directory (”`..`”), or resource forks (files that
// begin with “`._`”) but it does return other hidden files. If you need
// to perform a deep enumeration, use the
// [EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler] method
// instead.
//
// The order of the files in the returned array is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
//
// [Common File System Resource Keys]: https://developer.apple.com/documentation/CoreFoundation/common-file-system-resource-keys
func (f FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url INSURL, keys []string, mask NSDirectoryEnumerationOptions) ([]NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), url, objectivec.StringSliceToNSArray(keys), mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	}), nil

}

// Performs a shallow search of the specified directory and returns the paths
// of any contained items.
//
// path: The path to the directory whose contents you want to enumerate.
//
// # Return Value
//
// An array of [NSString] objects, each of which identifies a file, directory,
// or symbolic link contained in `path`. Returns an empty array if the
// directory exists but has no contents. In Objective-C, if an error occurs,
// this method returns `nil` and assigns an appropriate error object to the
// `error` parameter.
//
// # Discussion
//
// This method performs a shallow search of the directory and therefore does
// not traverse symbolic links or return the contents of any subdirectories.
// This method also does not return URLs for the current directory
// (”`.`”), parent directory (”`..`”), or resource forks (files that
// begin with “`._`”) but it does return other hidden files (files that
// begin with a period character). If you need to perform a deep enumeration,
// use the [EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
// method instead.
//
// The order of the files in the returned array is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f FileManager) ContentsOfDirectoryAtPathError(path string) ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("contentsOfDirectoryAtPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Returns a directory enumerator object that can be used to perform a deep
// enumeration of the directory at the specified path.
//
// path: The path of the directory to enumerate.
//
// # Return Value
//
// A [NSDirectoryEnumerator] object that enumerates the contents of the
// directory at `path`.
//
// # Discussion
//
// If `path` is a filename, the method returns an enumerator object that
// enumerates no files—the first call to [NextObject] will return `nil`.
//
// Because the enumeration is deep—that is, it lists the contents of all
// subdirectories—this enumerator object is useful for performing actions
// that involve large file-system subtrees. This method does not resolve
// symbolic links encountered in the traversal process, nor does it recurse
// through them if they point to a directory.
//
// This code fragment enumerates the subdirectories and files under a user’s
// [Documents] directory and processes all files with an extension of
// `XCUIElementTypeDoc`:
//
// The [NSDirectoryEnumerator] class has methods for obtaining the attributes
// of the existing path and of the parent directory and for skipping
// descendants of the existing path.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/enumerator(atPath:)
func (f FileManager) EnumeratorAtPath(path string) INSDirectoryEnumerator {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("enumeratorAtPath:"), objc.String(path))
	return NSDirectoryEnumeratorFromID(rv)
}

// Returns an array of URLs that identify the mounted volumes available on the
// device.
//
// propertyKeys: An array of keys that identify the file properties that you want
// pre-fetched for each volume. For each returned URL, the values for these
// keys are cached in the corresponding [NSURL] objects. You may specify `nil`
// for this parameter. For a list of keys you can specify, see Common File
// System Resource Keys.
//
// options: Option flags for the enumeration. For a list of possible values, see
// [FileManager.VolumeEnumerationOptions].
//
// # Return Value
//
// An array of [NSURL] objects identifying the mounted volumes.
//
// # Discussion
//
// This call may block if I/O is required to determine values for the
// requested `propertyKeys`.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
//
// [FileManager.VolumeEnumerationOptions]: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
func (f FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string, options NSVolumeEnumerationOptions) []NSURL {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), objectivec.StringSliceToNSArray(propertyKeys), options)
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	})
}

// Performs a deep enumeration of the specified directory and returns the
// paths of all of the contained subdirectories.
//
// path: The path of the directory to list.
//
// # Return Value
//
// An array of strings, each containing the path of an item in the directory
// specified by `path`. When using Objective-C, returns `nil` if an error
// occurred.
//
// # Discussion
//
// This method recurses the specified directory and its subdirectories. The
// method skips the “`.`” and “`..`” directories at each level of the
// recursion.
//
// Because this method recurses the directory’s contents, you might not want
// to use it in performance-critical code. Instead, consider using the
// [EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler] or
// [EnumeratorAtPath] method to enumerate the directory contents yourself.
// Doing so gives you more control over the retrieval of items and more
// opportunities to complete the enumeration or perform other tasks at the
// same time.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f FileManager) SubpathsOfDirectoryAtPathError(path string) ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("subpathsOfDirectoryAtPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Returns an array of strings identifying the paths for all items in the
// specified directory.
//
// path: The path of the directory to list.
//
// # Return Value
//
// An array of [NSString] objects, each of which contains the path of an item
// in the directory specified by `path`. If `path` is a symbolic link, this
// method traverses the link. This method returns `nil` if it cannot retrieve
// the device of the linked-to file.
//
// # Discussion
//
// This method recurses the specified directory and its subdirectories. The
// method skips the “`.`” and “`..`” directories at each level of the
// recursion.
//
// This method reveals every element of the subtree at `path`, including the
// contents of file packages (such as apps, nib files, and RTFD files). This
// code fragment gets the contents of `/System/Library/Fonts` after verifying
// that the directory exists:
//
// # Special Considerations
//
// In macOS 10.5 and later, use [SubpathsOfDirectoryAtPathError] instead.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/subpaths(atPath:)
func (f FileManager) SubpathsAtPath(path string) []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("subpathsAtPath:"), objc.String(path))
	return objc.ConvertSliceToStrings(rv)
}

// Creates a directory with the given attributes at the specified URL.
//
// url: A file URL that specifies the directory to create. If you want to specify a
// relative path, you must set the current working directory before creating
// the corresponding [NSURL] object. This parameter must not be `nil`.
//
// createIntermediates: If true, this method creates any nonexistent parent directories as part of
// creating the directory in `url`. If false, this method fails if any of the
// intermediate parent directories does not exist.
//
// attributes: The file attributes for the new directory. You can set the owner and group
// numbers, file permissions, and modification date. If you specify `nil` for
// this parameter, the directory is created according to the umask(2) macOS
// Developer Tools Manual Page of the process. The Supporting Types section
// lists the global constants used as keys in the `attributes` dictionary.
// Some of the keys, such as [hfsCreatorCode] and [hfsTypeCode], do not apply
// to directories.
//
// # Discussion
//
// If you specify `nil` for the `attributes` parameter, this method uses a
// default set of values for the owner, group, and permissions of any newly
// created directories in the path. Similarly, if you omit a specific
// attribute, the default value is used. The default values for newly created
// directories are as follows:
//
// - Permissions are set according to the umask of the current process. For
// more information, see umask. - The owner ID is set to the effective user ID
// of the process. - The group ID is set to that of the parent directory.
//
// If you want to specify a relative path for url, you must set the current
// working directory before creating the corresponding [NSURL] object.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(at:withIntermediateDirectories:attributes:)
//
// [hfsCreatorCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
// [hfsTypeCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
func (f FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url INSURL, createIntermediates bool, attributes INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("createDirectoryAtURL:withIntermediateDirectories:attributes:error:"), url, createIntermediates, attributes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createDirectoryAtURL:withIntermediateDirectories:attributes:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a directory with given attributes at the specified path.
//
// path: A path string identifying the directory to create. You may specify a full
// path or a path that is relative to the current working directory. This
// parameter must not be `nil`.
//
// createIntermediates: If true, this method creates any nonexistent parent directories as part of
// creating the directory in `path`. If false, this method fails if any of the
// intermediate parent directories does not exist. This method also fails if
// any of the intermediate path elements corresponds to a file and not a
// directory.
//
// attributes: The file attributes for the new directory and any newly created
// intermediate directories. You can set the owner and group numbers, file
// permissions, and modification date. If you specify `nil` for this parameter
// or omit a particular value, one or more default values are used as
// described in the discussion. For a list of keys you can include in this
// dictionary, see Supporting Types. Some of the keys, such as
// [hfsCreatorCode] and [hfsTypeCode], do not apply to directories.
//
// # Discussion
//
// If you specify `nil` for the `attributes` parameter, this method uses a
// default set of values for the owner, group, and permissions of any newly
// created directories in the path. Similarly, if you omit a specific
// attribute, the default value is used. The default values for newly created
// directories are as follows:
//
// - Permissions are set according to the umask of the current process. For
// more information, see umask. - The owner ID is set to the effective user ID
// of the process. - The group ID is set to that of the parent directory.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:withIntermediateDirectories:attributes:)
//
// [hfsCreatorCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
// [hfsTypeCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
func (f FileManager) CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("createDirectoryAtPath:withIntermediateDirectories:attributes:error:"), objc.String(path), createIntermediates, attributes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createDirectoryAtPath:withIntermediateDirectories:attributes:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a file with the specified content and attributes at the given
// location.
//
// path: The path for the new file.
//
// data: A data object containing the contents of the new file.
//
// attr: A dictionary containing the attributes to associate with the new file. You
// can use these attributes to set the owner and group numbers, file
// permissions, and modification date. For a list of keys, see
// [NSFileAttributeKey]. If you specify `nil` for `attributes`, the file is
// created with a set of default attributes.
//
// # Return Value
//
// true if the operation was successful or if the item already exists,
// otherwise false.
//
// # Discussion
//
// If you specify `nil` for the `attributes` parameter, this method uses a
// default set of values for the owner, group, and permissions of any newly
// created directories in the path. Similarly, if you omit a specific
// attribute, the default value is used. The default values for newly created
// files are as follows:
//
// - Permissions are set according to the umask of the current process. For
// more information, see umask. - The owner ID is set to the effective user ID
// of the process. - The group ID is set to that of the parent directory.
//
// If a file already exists at `path`, this method overwrites the contents of
// that file if the current process has the appropriate privileges to do so.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/createFile(atPath:contents:attributes:)
func (f FileManager) CreateFileAtPathContentsAttributes(path string, data INSData, attr INSDictionary) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("createFileAtPath:contents:attributes:"), objc.String(path), data, attr)
	return rv
}

// Removes the file or directory at the specified URL.
//
// URL: A file URL specifying the file or directory to remove. If the URL specifies
// a directory, the contents of that directory are recursively removed. You
// may specify `nil` for this parameter in Objective-C.
//
// # Discussion
//
// Prior to removing each item, the file manager asks its delegate if it
// should actually do so. It does this by calling the
// [FileManagerShouldRemoveItemAtURL] method; if that method is not
// implemented (or the process is running in OS X 10.5 or earlier) it calls
// the [FileManagerShouldRemoveItemAtPath] method instead. If the delegate
// method returns true, or if the delegate does not implement the appropriate
// methods, the file manager proceeds to remove the file or directory. If
// there is an error removing an item, the file manager may also call the
// delegate’s [FileManagerShouldProceedAfterErrorRemovingItemAtURL] or
// [FileManagerShouldProceedAfterErrorRemovingItemAtPath] method to determine
// how to proceed.
//
// Removing an item also removes all old versions of that item, invalidating
// any URLs returned by the
// [URLForPublishingUbiquitousItemAtURLExpirationDateError] method to old
// versions.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(at:)
func (f FileManager) RemoveItemAtURLError(URL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("removeItemAtURL:error:"), URL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes the file or directory at the specified path.
//
// path: A path string indicating the file or directory to remove. If the path
// specifies a directory, the contents of that directory are recursively
// removed. You may specify `nil` for this parameter in Objective-C.
//
// # Discussion
//
// Prior to removing each item, the file manager asks its delegate if it
// should actually do so. It does this by calling the
// [FileManagerShouldRemoveItemAtURL] method; if that method is not
// implemented (or the process is running in OS X 10.5 or earlier) it calls
// the [FileManagerShouldRemoveItemAtPath] method instead. If the delegate
// method returns true, or if the delegate does not implement the appropriate
// methods, the file manager proceeds to remove the file or directory. If
// there is an error removing an item, the file manager may also call the
// delegate’s [FileManagerShouldProceedAfterErrorRemovingItemAtURL] or
// [FileManagerShouldProceedAfterErrorRemovingItemAtPath] method to determine
// how to proceed.
//
// Removing an item also removes all old versions of that item, invalidating
// any URLs returned by the
// [URLForPublishingUbiquitousItemAtURLExpirationDateError] method to old
// versions.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(atPath:)
func (f FileManager) RemoveItemAtPathError(path string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("removeItemAtPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeItemAtPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Moves an item to the trash.
//
// url: The item to move to the trash.
//
// outResultingURL: On input, a pointer to a URL object. On output, this pointer is set to the
// item’s location in the trash. The actual name of the item may be changed
// when moving it to the trash, so use this URL to access it. You may specify
// `nil` for this parameter if you do not want the information.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/trashItem(at:resultingItemURL:)
func (f FileManager) TrashItemAtURLResultingItemURLError(url INSURL, outResultingURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("trashItemAtURL:resultingItemURL:error:"), url, outResultingURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("trashItemAtURL:resultingItemURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Replaces the contents of the item at the specified URL in a manner that
// ensures no data loss occurs.
//
// originalItemURL: The item containing the content you want to replace.
//
// newItemURL: The item containing the new content for `originalItemURL`. It is
// recommended that you put this item in a temporary directory as provided by
// the OS. If a temporary directory is not available, put this item in a
// uniquely named directory that is in the same directory as the original
// item.
//
// backupItemName: If provided, the name used to create a backup of the original item. The
// backup is placed in the same directory as the original item. If an error
// occurs during the creation of the backup item, the operation fails. If
// there is already an item with the same name as the backup item, that item
// will be removed.
//
// The backup item will be removed in the event of success unless the
// [NSFileManagerItemReplacementWithoutDeletingBackupItem] option is provided
// in `options`.
//
// options: The options to use during the replacement. Typically, you pass
// [NSFileManagerItemReplacementUsingNewMetadataOnly] for this parameter,
// which uses only the metadata from the new item. You can also combine the
// options described in [FileManager.ItemReplacementOptions] using the
// C-bitwise OR operator.
//
// resultingURL: On input, a pointer for a URL object. When the item is replaced, this
// pointer is set to the URL of the new item. If no new file system object is
// required, the URL object in this parameter may be the same passed to the
// `originalItemURL` parameter. However, if a new file system object is
// required, the URL object may be different. For example, replacing an RTF
// document with an RTFD document requires the creation of a new file.
//
// # Discussion
//
// By default, the creation date, permissions, Finder label and color, and
// Spotlight comments of the original item are preserved on the new item. You
// can configure which metadata is preserved using the `options` parameter.
//
// This method works only when the `originalItemURL` and `newItemURL`
// parameters are located on the same volume. Attempting to call this method
// by passing `originalItemURL` and `newItemURL` parameters that have
// locations on different volumes results in an error. Instead, you can call
// the [URLForDirectoryInDomainAppropriateForURLCreateError] method, passing
// [NSItemReplacementDirectory] as the search path directory, to get a
// temporary URL on the destination’s volume that is suitable for use with
// this method.
//
// If an error occurs and the original item is not in the original location or
// a temporary location, the resulting error object contains a user info
// dictionary with the key `"NSFileOriginalItemLocationKey"`. The value
// assigned to that key is an [NSURL] object with the location of the item.
// The error code is one of the file-related errors described in [NSError
// Codes].
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
//
// [FileManager.ItemReplacementOptions]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
// [NSError Codes]: https://developer.apple.com/documentation/Foundation/1448136-nserror-codes
func (f FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL INSURL, newItemURL INSURL, backupItemName string, options NSFileManagerItemReplacementOptions, resultingURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), originalItemURL, newItemURL, objc.String(backupItemName), options, resultingURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Copies the file at the specified URL to a new location synchronously.
//
// srcURL: The file URL that identifies the file you want to copy. The URL in this
// parameter must not be a file reference URL. This parameter must not be
// `nil`.
//
// dstURL: The URL at which to place the copy of `srcURL`. The URL in this parameter
// must not be a file reference URL and must include the name of the file in
// its new location. This parameter must not be `nil`.
//
// # Discussion
//
// When copying items, the current process must have permission to read the
// file or directory at `srcURL` and write the parent directory of `dstURL`.
// If the item at `srcURL` is a directory, this method copies the directory
// and all of its contents, including any hidden files. If a file with the
// same name already exists at `dstURL`, this method stops the copy attempt
// and returns an appropriate error. If the last component of `srcURL` is a
// symbolic link, only the link is copied to the new path.
//
// Prior to copying each item, the file manager asks its delegate if it should
// actually do so. It does this by calling the
// [FileManagerShouldCopyItemAtURLToURL] method; if that method is not
// implemented (or the process is running in OS X 10.5 or earlier) it calls
// the [FileManagerShouldCopyItemAtPathToPath] method instead. If the delegate
// method returns true, or if the delegate does not implement the appropriate
// methods, the file manager proceeds to copy the file or directory. If there
// is an error copying an item, the file manager may also call the
// delegate’s [FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath] method to
// determine how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(at:to:)
func (f FileManager) CopyItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("copyItemAtURL:toURL:error:"), srcURL, dstURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyItemAtURL:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Copies the item at the specified path to a new location synchronously.
//
// srcPath: The path to the file or directory you want to move. This parameter must not
// be `nil`.
//
// dstPath: The path at which to place the copy of `srcPath`. This path must include
// the name of the file or directory in its new location. This parameter must
// not be `nil`.
//
// # Discussion
//
// When copying items, the current process must have permission to read the
// file or directory at `srcPath` and write the parent directory of `dstPath`.
// If the item at `srcPath` is a directory, this method copies the directory
// and all of its contents, including any hidden files. If a file with the
// same name already exists at `dstPath`, this method stops the copy attempt
// and returns an appropriate error. If the last component of `srcPath` is a
// symbolic link, only the link is copied to the new path.
//
// Prior to copying an item, the file manager asks its delegate if it should
// actually do so for each item. It does this by calling the
// [FileManagerShouldCopyItemAtURLToURL] method; if that method is not
// implemented it calls the [FileManagerShouldCopyItemAtPathToPath] method
// instead. If the delegate method returns true, or if the delegate does not
// implement the appropriate methods, the file manager copies the given file
// or directory. If there is an error copying an item, the file manager may
// also call the delegate’s
// [FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath] method to
// determine how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(atPath:toPath:)
func (f FileManager) CopyItemAtPathToPathError(srcPath string, dstPath string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("copyItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyItemAtPath:toPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Moves the file or directory at the specified URL to a new location
// synchronously.
//
// srcURL: The file URL that identifies the file or directory you want to move. The
// URL in this parameter must not be a file reference URL. This parameter must
// not be `nil`.
//
// dstURL: The new location for the item in `srcURL`. The URL in this parameter must
// not be a file reference URL and must include the name of the file or
// directory in its new location. This parameter must not be `nil`.
//
// # Discussion
//
// When moving items, the current process must have permission to read the
// item at `srcURL` and write the parent directory of `dstURL`. If the item at
// `srcURL` is a directory, this method moves the directory and all of its
// contents, including any hidden files. If an item with the same name already
// exists at `dstURL`, this method stops the move attempt and returns an
// appropriate error. If the last component of `srcURL` is a symbolic link,
// only the link is moved to the new path; the item pointed to by the link
// remains at its current location.
//
// Prior to moving the item, the file manager asks its delegate if it should
// actually move it. It does this by calling the
// [FileManagerShouldMoveItemAtURLToURL] method; if that method is not
// implemented it calls the [FileManagerShouldMoveItemAtPathToPath] method
// instead. If the item being moved is a directory, the file manager notifies
// the delegate only for the directory itself and not for any of its contents.
// If the delegate method returns true, or if the delegate does not implement
// the appropriate methods, the file manager moves the file. If there is an
// error moving one out of several items, the file manager may also call the
// delegate’s [FileManagerShouldProceedAfterErrorMovingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorMovingItemAtPathToPath] method to
// determine how to proceed.
//
// If the source and destination of the move operation are not on the same
// volume, this method copies the item first and then removes it from its
// current location. This behavior may trigger additional delegate
// notifications related to copying and removing individual items.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(at:to:)
func (f FileManager) MoveItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("moveItemAtURL:toURL:error:"), srcURL, dstURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("moveItemAtURL:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Moves the file or directory at the specified path to a new location
// synchronously.
//
// srcPath: The path to the file or directory you want to move. This parameter must not
// be `nil`.
//
// dstPath: The new path for the item in `srcPath`. This path must include the name of
// the file or directory in its new location. This parameter must not be
// `nil`.
//
// # Discussion
//
// When moving items, the current process must have permission to read the
// item at `srcPath` and write the parent directory of `dstPath`. If the item
// at `srcPath` is a directory, this method moves the directory and all of its
// contents, including any hidden files. If an item with the same name already
// exists at `dstPath`, this method stops the move attempt and returns an
// appropriate error. If the last component of `srcPath` is a symbolic link,
// only the link is moved to the new path; the item pointed to by the link
// remains at its current location.
//
// Prior to moving the item, the file manager asks its delegate if it should
// actually move it. It does this by calling the
// [FileManagerShouldMoveItemAtURLToURL] method; if that method is not
// implemented it calls the [FileManagerShouldMoveItemAtPathToPath] method
// instead. If the item being moved is a directory, the file manager notifies
// the delegate only for the directory itself and not for any of its contents.
// If the delegate method returns true, or if the delegate does not implement
// the appropriate methods, the file manager moves the file. If there is an
// error moving one out of several items, the file manager may also call the
// delegate’s [FileManagerShouldProceedAfterErrorMovingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorMovingItemAtPathToPath] method to
// determine how to proceed.
//
// If the source and destination of the move operation are not on the same
// volume, this method copies the item first and then removes it from its
// current location. This behavior may trigger additional delegate
// notifications related to copying and removing individual items.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(atPath:toPath:)
func (f FileManager) MoveItemAtPathToPathError(srcPath string, dstPath string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("moveItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("moveItemAtPath:toPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the URL for the iCloud container associated with the specified
// identifier and establishes access to that container.
//
// containerIdentifier: The fully-qualified container identifier for an iCloud container directory.
// The string you specify must not contain wildcards and must be of the form
// `.`, where is your development team ID and is the bundle identifier of the
// container you want to access.
//
// The container identifiers for your app must be declared in the
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeUbiquity()-container-identifiers`
// array of the `XCUIElementTypeEntitlements` property list file in your Xcode
// project.
//
// If you specify `nil` for this parameter, this method returns the first
// container listed in the
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeUbiquity()-container-identifiers`
// entitlement array.
//
// # Return Value
//
// A URL pointing to the specified ubiquity container, or `nil` if the
// container could not be located or if iCloud storage is unavailable for the
// current user or device.
//
// # Discussion
//
// You use this method to determine the location of your app’s ubiquity
// container directories and to configure your app’s initial iCloud access.
// The first time you call this method for a given ubiquity container, the
// system extends your app’s sandbox to include that container. In iOS, you
// must call this method at least once before trying to search for cloud-based
// files in the ubiquity container. If your app accesses multiple ubiquity
// containers, call this method once for each container. In macOS, you do not
// need to call this method if you use [NSDocument]-based objects, because the
// system then calls this method automatically.
//
// You can use the URL returned by this method to build paths to files and
// directories within your app’s ubiquity container. Each app that syncs
// documents to the cloud must have at least one associated ubiquity container
// in which to put those files. This container can be unique to the app or
// shared by multiple apps.
//
// In addition to writing to its own ubiquity container, an app can write to
// any container directory for which it has the appropriate permission. Each
// additional ubiquity container should be listed as an additional value in
// the
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeUbiquity()-container-identifiers`
// entitlement array.
//
// To learn how to view your development team’s unique value, read To view
// the team ID in Tools Workflow Guide for Mac.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/url(forUbiquityContainerIdentifier:)
//
// [NSDocument]: https://developer.apple.com/documentation/AppKit/NSDocument
func (f FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string) INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URLForUbiquityContainerIdentifier:"), objc.String(containerIdentifier))
	return NSURLFromID(rv)
}

// Returns a Boolean indicating whether the item is targeted for storage in
// iCloud.
//
// url: Specify the URL for the file or directory whose status you want to check.
//
// # Return Value
//
// true if the item is targeted for iCloud storage or false if it is not. This
// method also returns false if no item exists at `url`.
//
// # Discussion
//
// This method reflects only whether the item should be stored in iCloud
// because a call was made to the [SetUbiquitousItemAtURLDestinationURLError]
// method with a value of true for its `flag` parameter. This method does not
// reflect whether the file has actually been uploaded to any iCloud servers.
// To determine a file’s upload status, check the
// [NSURLUbiquitousItemIsUploadedKey] attribute of the corresponding [NSURL]
// object.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f FileManager) IsUbiquitousItemAtURL(url INSURL) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isUbiquitousItemAtURL:"), url)
	return rv
}

// Indicates whether the item at the specified URL should be stored in iCloud.
//
// flag: true to move the item to iCloud or false to remove it from iCloud (if it is
// there currently).
//
// url: The URL of the item (file or directory) that you want to store in iCloud.
//
// destinationURL: When moving a file into iCloud, this is the location in iCloud at which to
// store the file or directory. This URL must be constructed from a URL
// returned by the [URLForUbiquityContainerIdentifier] method, which you use
// to retrieve the desired iCloud container directory. The URL you specify may
// contain additional subdirectories so that you can organize your files
// hierarchically in iCloud. However, you are responsible for creating those
// intermediate subdirectories (using the [NSFileManager] class) in your
// iCloud container directory. When moving a file , this is the location on
// the local device.
//
// # Discussion
//
// Use this method to move a file from its current location to iCloud. For
// files located in an app’s sandbox, this involves physically removing the
// file from the sandbox container. (The system extends your app’s sandbox
// privileges to give it access to files it moves to iCloud.) You can also use
// this method to move files out of iCloud and back into a local directory.
//
// If your app is presenting the file’s contents to the user, it must have
// an active file presenter object configured to monitor the specified file or
// directory before calling this method. When you specify true for the `flag`
// parameter, this method attempts to move the file or directory to the cloud
// and returns true if it is successful. Calling this method also notifies
// your file presenter of the new location of the file so that your app can
// continue to operate on it.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url INSURL, destinationURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, url, destinationURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setUbiquitous:itemAtURL:destinationURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Starts downloading (if necessary) the specified item to the local system.
//
// url: The URL for the file or directory in the cloud that you want to download.
//
// # Discussion
//
// If a cloud-based file or directory has not been downloaded yet, calling
// this method starts the download process. If the item exists locally,
// calling this method synchronizes the local copy with the version in the
// cloud.
//
// For a given URL, you can determine if a file is downloaded by getting the
// value of the [NSMetadataUbiquitousItemDownloadingStatusKey] key. You can
// also use related keys to determine the current progress in downloading the
// file.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f FileManager) StartDownloadingUbiquitousItemAtURLError(url INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startDownloadingUbiquitousItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes the local copy of the specified item that’s stored in iCloud.
//
// url: The URL to a file or directory in iCloud storage.
//
// # Discussion
//
// Don’t use a coordinated write to perform this operation. In macOS 10.7 or
// earlier, the framework takes a coordinated write in its implementation of
// this method, so taking one in your app causes a deadlock. In macOS 10.8 and
// later, the framework detects coordination done by your app on the same
// thread as the call to this method, to avoid deadlocking.
//
// This method doesn’t remove the item from iCloud. It removes only the
// local version. You can then use [StartDownloadingUbiquitousItemAtURLError]
// to force iCloud to download a new version of the file or directory from the
// server.
//
// To delete a file permanently from the user’s iCloud storage, use the
// regular [NSFileManager] routines for deleting files and directories.
// Remember that deleting items from iCloud can’t be undone. Once deleted,
// the item is gone forever.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f FileManager) EvictUbiquitousItemAtURLError(url INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("evictUbiquitousItemAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("evictUbiquitousItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a URL that can be emailed to users to allow them to download a copy
// of a flat file item from iCloud.
//
// url: The URL of the item in the cloud that you want to share. The URL must be
// prefixed with the base URL returned from the
// [URLForUbiquityContainerIdentifier] method that corresponds to the item’s
// location. The file must be a flat file, not a bundle. The file at the
// specified URL must already be uploaded to iCloud when you call this method.
//
// outDate: On input, a pointer to a variable for a date object. On output, this
// parameter contains the date after which the item is no longer available at
// the returned URL. You may specify `nil` for this parameter if you are not
// interested in the date.
//
// # Return Value
//
// A URL with which users can download a copy of the item at `url`. In
// Objective-C, returns `nil` if the URL could not be created for any reason.
//
// # Discussion
//
// This method creates a snapshot of the specified flat file and places that
// copy in a temporary iCloud location where it can be accessed by other users
// using the returned URL. The snapshot reflects the contents of the file at
// the time the URL was generated and is not updated when subsequent changes
// are made to the original file in the user’s iCloud storage. The snapshot
// file remains available at the specified URL until the date specified in the
// `outDate` parameter, after which it is automatically deleted. Explicitly
// deleting the item by calling the [RemoveItemAtURLError] or
// [RemoveItemAtPathError] method also deletes all old versions of the item,
// invalidating URLs to those versions returned by this method.
//
// Your app must have access to the network for this call to succeed. If the
// specified file is in the process of being uploaded to iCloud, you must not
// call this method until the upload has finished.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/url(forPublishingUbiquitousItemAt:expiration:)
func (f FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url INSURL, outDate INSDate) (INSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URLForPublishingUbiquitousItemAtURL:expirationDate:error:"), url, outDate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}

// Asynchronously pauses sync of an item at the given URL.
//
// url: The URL of the item for which to pause sync.
//
// completionHandler: A closure or block that the framework calls when the pause action
// completes. It receives a single [NSError] parameter to indicate an error
// that prevented pausing; this value is `nil` if the pause succeeded. In
// Swift, you can omit the completion handler and catch the thrown error
// instead.
//
// # Discussion
//
// Call this when opening an item to prevent sync from altering the contents
// of the URL. Once paused, the file provider will not upload local changes
// nor download remote changes.
//
// While paused, call
// [UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler]
// when the document is in a stable state. This action keeps the server
// version as up-to-date as possible.
//
// If the item is already paused, a second call to this method reports
// success. If the file provider is already applying changes to the item, the
// pause fails with an [NSFileWriteUnknownError], with an underlying error
// that has domain [NSPOSIXErrorDomain] and code [EBUSY]. If the pause fails,
// wait for the state to stabilize before retrying. Pausing also fails with
// [featureUnsupported] if `url` refers to a regular (non-package) directory.
//
// Pausing sync is independent of the calling app’s lifecycle; sync
// doesn’t automatically resume if the app closes or crashes and relaunches
// later. To resume syncing, explicitly call
// [ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]. Always be
// sure to resume syncing before you close the item.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/pauseSyncForUbiquitousItem(at:completionHandler:)
//
// [EBUSY]: https://developer.apple.com/documentation/Foundation/POSIXError/EBUSY
// [featureUnsupported]: https://developer.apple.com/documentation/Foundation/CocoaError/featureUnsupported
func (f FileManager) PauseSyncForUbiquitousItemAtURLCompletionHandler(url INSURL, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("pauseSyncForUbiquitousItemAtURL:completionHandler:"), url, _block1)
}

// Asynchronously resumes the sync on a paused item using the given resume
// behavior.
//
// url: The URL of the item for which to resume sync.
//
// behavior: A [NSFileManagerResumeSyncBehavior] value that tells the file manager how
// to handle conflicts between local and remote versions of files.
//
// completionHandler: A closure or block that the framework calls when the resume action
// completes. It receives a single [NSError] parameter to indicate an error
// that prevented the resume action; the value is `nil` if the resume
// succeeded. In Swift, you can omit the completion handler and catch the
// thrown error instead.
//
// # Discussion
//
// Always call this method when your app closes an item to allow the file
// provider to sync local changes back to the server.
//
// In most situations, the
// [NSFileManagerResumeSyncBehaviorPreserveLocalChanges] behavior is the best
// choice to avoid any risk of data loss.
//
// The resume call fails with [featureUnsupported] if `url` isn’t currently
// paused. If the device isn’t connected to the network, the call may fail
// with [NSFileWriteUnknownError], with the underlying error of
// [serverUnreachable].
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/resumeSyncForUbiquitousItem(at:with:completionHandler:)
//
// [NSFileManagerResumeSyncBehavior]: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
// [featureUnsupported]: https://developer.apple.com/documentation/Foundation/CocoaError/featureUnsupported
// [serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/serverUnreachable
func (f FileManager) ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url INSURL, behavior NSFileManagerResumeSyncBehavior, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("resumeSyncForUbiquitousItemAtURL:withBehavior:completionHandler:"), url, behavior, _block2)
}

// Asynchronously fetches the latest remote version of a given item from the
// server.
//
// url: The URL of the item for which to check the version.
//
// completionHandler: A closure or block that the framework calls when the fetch action
// completes. It receives parameters of types [NSFileVersion] and [NSError].
// The error is `nil` if fetching the remote version succeeded; otherwise it
// indicates the error that caused the call to fail. In Swift, you can omit
// the completion handler, catching any error in a `do`-`catch` block and
// receiving the version as the return value.
//
// # Discussion
//
// Use this method if uploading fails due to a version conflict and sync is
// paused. In this case, fetching the latest remote version allows you to
// inspect the newer item from the server, resolve the conflict, and resume
// uploading.
//
// The version provided by this call depends on several factors:
//
// - If there is no newer version of the file on the server, the caller
// receives the current version of the file. - If the server has a newer
// version and sync isn’t paused, this call replaces the local item and
// provides the version of the new item. - If the server has a newer version
// but sync is paused, the returned version points to a side location. In this
// case, call [ReplaceItemAtURLOptionsError] on the provided version object to
// replace the local item with the newer item from the server.
//
// If the device isn’t connected to the network, the call may fail with
// [NSFileReadUnknownError], with the underlying error of [serverUnreachable].
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/fetchLatestRemoteVersionOfItem(at:completionHandler:)
//
// [serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/serverUnreachable
func (f FileManager) FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url INSURL, completionHandler FileVersionErrorHandler) {
	_block1, _ := NewFileVersionErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("fetchLatestRemoteVersionOfItemAtURL:completionHandler:"), url, _block1)
}

// Asynchronously uploads the local version of the item using the provided
// conflict resolution policy.
//
// url: The URL of the item for which to check the version.
//
// conflictResolutionPolicy: The policy the file manager applies if the local and server versions
// conflict.
//
// completionHandler: A closure or block that the framework calls when the upload completes. It
// receives parameters of types [NSFileVersion] and [NSError]. The error is
// `nil` if fetching the remote version succeeded; otherwise it indicates the
// error that caused the call to fail. In Swift, you can omit the completion
// handler, catching any error in a `do`-`catch` block and receiving the
// version as the return value.
//
// # Discussion
//
// Once your app pauses a sync for an item, call this method every time your
// document is in a stable state. This action keeps the server version as
// up-to-date as possible.
//
// If the server has a newer version than the one to which the app made
// changes, uploading fails with [NSFileWriteUnknownError], with an underlying
// error of [localVersionConflictingWithServer]. In this case, call
// [FetchLatestRemoteVersionOfItemAtURLCompletionHandler], rebase local
// changes on top of that version, and retry the upload.
//
// If the device isn’t connected to the network, the call may fail with
// [NSFileWriteUnknownError], with the underlying error of
// [serverUnreachable].
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/uploadLocalVersionOfUbiquitousItem(at:withConflictResolutionPolicy:completionHandler:)
//
// [localVersionConflictingWithServer]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/localVersionConflictingWithServer
// [serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/serverUnreachable
func (f FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url INSURL, conflictResolutionPolicy NSFileManagerUploadLocalVersionConflictPolicy, completionHandler FileVersionErrorHandler) {
	_block2, _ := NewFileVersionErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("uploadLocalVersionOfUbiquitousItemAtURL:withConflictResolutionPolicy:completionHandler:"), url, conflictResolutionPolicy, _block2)
}

// Creates a symbolic link at the specified URL that points to an item at the
// given URL.
//
// url: The file URL at which to create the new symbolic link. The last path
// component of the URL issued as the name of the link.
//
// destURL: The file URL that contains the item to be pointed to by the link. In other
// words, this is the destination of the link.
//
// # Discussion
//
// This method does not traverse symbolic links contained in `destURL`, making
// it possible to create symbolic links to locations that do not yet exist.
// Also, if the final path component in `url` is a symbolic link, that link is
// not followed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url INSURL, destURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), url, destURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createSymbolicLinkAtURL:withDestinationURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a symbolic link that points to the specified destination.
//
// path: The path at which to create the new symbolic link. The last path component
// is used as the name of the link.
//
// destPath: The path that contains the item to be pointed to by the link. In other
// words, this is the destination of the link.
//
// # Discussion
//
// This method does not traverse symbolic links contained in `destPath`,
// making it possible to create symbolic links to locations that do not yet
// exist. Also, if the final path component in `path` is a symbolic link, that
// link is not followed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:withDestinationPath:)
func (f FileManager) CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("createSymbolicLinkAtPath:withDestinationPath:error:"), objc.String(path), objc.String(destPath), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createSymbolicLinkAtPath:withDestinationPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a hard link between the items at the specified URLs.
//
// srcURL: The file URL that identifies the source of the link. The URL in this
// parameter must not be a file reference URL; it must specify the actual path
// to the item. The value in this parameter must not be `nil`.
//
// dstURL: The file URL that specifies where you want to create the hard link. The URL
// in this parameter must not be a file reference URL; it must specify the
// actual path to the item. The value in this parameter must not be `nil`.
//
// # Discussion
//
// Use this method to create hard links between files in the current file
// system. If `srcURL` is a directory, this method creates a new directory at
// `dstURL` and then creates hard links for the items in that directory. If
// `srcURL` is (or contains) a symbolic link, the symbolic link is copied and
// not converted to a hard link at `dstURL`.
//
// Prior to linking each item, the file manager asks its delegate if it should
// actually create the link. It does this by calling the
// [FileManagerShouldLinkItemAtURLToURL] method; if that method is not
// implemented it calls the [FileManagerShouldLinkItemAtPathToPath] method
// instead. If the delegate method returns true, or if the delegate does not
// implement the appropriate methods, the file manager creates the hard link.
// If there is an error linking one out of several items, the file manager may
// also call the delegate’s
// [FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath] method to
// determine how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(at:to:)
func (f FileManager) LinkItemAtURLToURLError(srcURL INSURL, dstURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("linkItemAtURL:toURL:error:"), srcURL, dstURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("linkItemAtURL:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a hard link between the items at the specified paths.
//
// srcPath: The path that specifies the item you wish to link to. The value in this
// parameter must not be `nil`.
//
// dstPath: The path that identifies the location where the link will be created. The
// value in this parameter must not be `nil`.
//
// # Discussion
//
// Use this method to create hard links between files in the current file
// system. If `srcPath` is a directory, this method creates a new directory at
// `dstPath` and then creates hard links for the items in that directory. If
// `srcPath` is (or contains) a symbolic link, the symbolic link is copied to
// the new location and not converted to a hard link.
//
// Prior to linking each item, the file manager asks its delegate if it should
// actually create the link. It does this by calling the
// [FileManagerShouldLinkItemAtURLToURL] method; if that method is not
// implemented it calls the [FileManagerShouldLinkItemAtPathToPath] method
// instead. If the delegate method returns true, or if the delegate does not
// implement the appropriate methods, the file manager creates the hard link.
// If there is an error linking one out of several items, the file manager may
// also call the delegate’s
// [FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL] or
// [FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath] method to
// determine how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(atPath:toPath:)
func (f FileManager) LinkItemAtPathToPathError(srcPath string, dstPath string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("linkItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("linkItemAtPath:toPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the path of the item pointed to by a symbolic link.
//
// path: The path of a file or directory.
//
// # Return Value
//
// An [NSString] object containing the path of the directory or file to which
// the symbolic link `path` refers. When using Objective-C, returns `nil` upon
// failure. If the symbolic link is specified as a relative path, that
// relative path is returned.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/destinationOfSymbolicLink(atPath:)
func (f FileManager) DestinationOfSymbolicLinkAtPathError(path string) (string, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("destinationOfSymbolicLinkAtPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return "", NSErrorFrom(errorPtr)
	}
	return objc.IDToString(rv), nil

}

// Returns a Boolean value that indicates whether a file or directory exists
// at a specified path.
//
// path: The path of the file or directory. If `path` begins with a tilde (`~`), it
// must first be expanded with [StringByExpandingTildeInPath]; otherwise, this
// method returns false.
//
// App Sandbox does not restrict which path values may be passed to this
// parameter.
//
// # Return Value
//
// true if a file at the specified path exists, or false if the file does not
// exist or its existence could not be determined.
//
// # Discussion
//
// If the file at `path` is inaccessible to your app, perhaps because one or
// more parent directories are inaccessible, this method returns false. If the
// final element in `path` specifies a symbolic link, this method traverses
// the link and returns true or false based on the existence of the file at
// the link destination.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:)
func (f FileManager) FileExistsAtPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("fileExistsAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether a file or directory exists
// at a specified path.
//
// path: The path of a file or directory. If `path` begins with a tilde (`~`), it
// must first be expanded with [StringByExpandingTildeInPath], or this method
// will return false.
//
// isDirectory: Upon return, contains true if `path` is a directory or if the final path
// element is a symbolic link that points to a directory; otherwise, contains
// false. If `path` doesn’t exist, this value is undefined upon return. Pass
// [NULL] if you do not need this information.
//
// # Return Value
//
// true if a file at the specified path exists, or false if the file’s does
// not exist or its existence could not be determined.
//
// # Discussion
//
// If the file at `path` is inaccessible to your app, perhaps because one or
// more parent directories are inaccessible, this method returns false. If the
// final element in `path` specifies a symbolic link, this method traverses
// the link and returns true or false based on the existence of the file at
// the link destination.
//
// If you need to further determine whether `path` is a package, use the
// [isFilePackage(atPath:)] method of [NSWorkspace].
//
// The following example code gets an array that identifies the fonts in the
// user’s fonts directory:
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
//
// [NSWorkspace]: https://developer.apple.com/documentation/AppKit/NSWorkspace
// [isFilePackage(atPath:)]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isFilePackage(atPath:)
func (f FileManager) FileExistsAtPathIsDirectory(path string) (bool, bool) {
	var isDirectory bool
	rv := objc.Send[bool](f.ID, objc.Sel("fileExistsAtPath:isDirectory:"), objc.String(path), unsafe.Pointer(&isDirectory))
	return isDirectory, rv
}

// Returns a Boolean value that indicates whether the invoking object appears
// able to read a specified file.
//
// path: A file path.
//
// # Return Value
//
// true if the current process has read privileges for the file at `path`;
// otherwise false if the process does not have read privileges or the
// existence of the file could not be determined.
//
// # Discussion
//
// If the file at `path` is inaccessible to your app, perhaps because it does
// not have search privileges for one or more parent directories, this method
// returns false. This method traverses symbolic links in the path. This
// method also uses the real user ID and group ID, as opposed to the effective
// user and group IDs, to determine if the file is readable.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f FileManager) IsReadableFileAtPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isReadableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears
// able to write to a specified file.
//
// path: A file path.
//
// # Return Value
//
// true if the current process has write privileges for the file at `path`;
// otherwise false if the process does not have write privileges or the
// existence of the file could not be determined.
//
// # Discussion
//
// If the file at `path` is inaccessible to your app, perhaps because it does
// not have search privileges for one or more parent directories, this method
// returns false. This method traverses symbolic links in the path. This
// method also uses the real user ID and group ID, as opposed to the effective
// user and group IDs, to determine if the file is writable.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f FileManager) IsWritableFileAtPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isWritableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the operating system appears
// able to execute a specified file.
//
// path: A file path.
//
// # Return Value
//
// true if the current process has execute privileges for the file at `path`;
// otherwise false if the process does not have execute privileges or the
// existence of the file could not be determined.
//
// # Discussion
//
// If the file at `path` is inaccessible to your app, perhaps because it does
// not have search privileges for one or more parent directories, this method
// returns false. This method traverses symbolic links in the path. This
// method also uses the real user ID and group ID, as opposed to the effective
// user and group IDs, to determine if the file is executable.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f FileManager) IsExecutableFileAtPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isExecutableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears
// able to delete a specified file.
//
// path: A file path.
//
// # Return Value
//
// true if the current process has delete privileges for the file at `path`;
// otherwise false if the process does not have delete privileges or the
// existence of the file could not be determined.
//
// # Discussion
//
// For a directory or file to be deletable, the current process must either be
// able to write to the parent directory of `path` or it must have the same
// owner as the item at `path`. If `path` is a directory, every item contained
// in `path` must be deletable by the current process.
//
// If the file at `path` is inaccessible to your app, perhaps because it does
// not have search privileges for one or more parent directories, this method
// returns false. If the item at `path` is a symbolic link, it is not
// traversed.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f FileManager) IsDeletableFileAtPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isDeletableFileAtPath:"), objc.String(path))
	return rv
}

// Returns an array of strings representing the user-visible components of a
// given path.
//
// path: A pathname.
//
// # Return Value
//
// An array of [NSString] objects representing the user-visible (for the
// Finder, Open and Save panels, and so on) components of `path`. Returns
// `nil` if path does not exist.
//
// # Discussion
//
// These components cannot be used for path operations and are only suitable
// for display to the user.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/componentsToDisplay(forPath:)
func (f FileManager) ComponentsToDisplayForPath(path string) []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("componentsToDisplayForPath:"), objc.String(path))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the display name of the file or directory at a specified path.
//
// path: The path of a file or directory.
//
// # Return Value
//
// The name of the file or directory at `path` in a localized form appropriate
// for presentation to the user. If there is no file or directory at `path`,
// or if an error occurs, returns `path` as is.
//
// # Discussion
//
// Display names are user-friendly names for files. They are typically used to
// localize standard file and directory names according to the user’s
// language settings. They may also reflect other modifications, such as the
// removal of filename extensions. Such modifications are used only when
// displaying the file or directory to the user and do not reflect the actual
// path to the item in the file system. For example, if the current user’s
// preferred language is French, the following code fragment logs the name
// `Bibliothèque` and not the name [Library], which is the actual name of the
// directory.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/displayName(atPath:)
func (f FileManager) DisplayNameAtPath(path string) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("displayNameAtPath:"), objc.String(path))
	return NSStringFromID(rv).String()
}

// Returns the attributes of the item at a given path.
//
// path: The path of a file or directory.
//
// # Return Value
//
// A dictionary object that describes the attributes (file, directory,
// symlink, and so on) of the file specified by `path` (or `nil` if an error
// occurred in Objective-C). The keys in the dictionary are described in `File
// Attribute Keys`.
//
// # Discussion
//
// If the item at the path is a symbolic link—that is, the value of the
// [type] key in the attributes dictionary is [typeSymbolicLink]—you can use
// the [DestinationOfSymbolicLinkAtPathError] method to retrieve the path of
// the item pointed to by the link. You can also use the
// [StringByResolvingSymlinksInPath] method of [NSString] to resolve links in
// the path before retrieving the item’s attributes.
//
// As a convenience, [NSDictionary] provides a set of methods (declared as a
// category on [NSDictionary]) for quickly and efficiently obtaining attribute
// information from the returned dictionary: [FileGroupOwnerAccountName],
// [FileModificationDate], [FileOwnerAccountName], [FilePosixPermissions],
// [FileSize], [FileSystemFileNumber], [FileSystemNumber], and [FileType].
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfItem(atPath:)
//
// [typeSymbolicLink]: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeSymbolicLink
// [type]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/type
func (f FileManager) AttributesOfItemAtPathError(path string) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributesOfItemAtPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}

// Returns a dictionary that describes the attributes of the mounted file
// system on which a given path resides.
//
// path: Any pathname within the mounted file system.
//
// # Return Value
//
// A dictionary object that describes the attributes of the mounted file
// system on which `path` resides. See `File-System Attribute Keys` for a
// description of the keys available in the dictionary.
//
// # Discussion
//
// This method does not traverse a terminal symbolic link.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfFileSystem(forPath:)
func (f FileManager) AttributesOfFileSystemForPathError(path string) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributesOfFileSystemForPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}

// Sets the attributes of the specified file or directory.
//
// attributes: A dictionary containing as keys the attributes to set for `path` and as
// values the corresponding value for the attribute. You can set the following
// attributes: [busy], [creationDate], [extensionHidden],
// [groupOwnerAccountID], [groupOwnerAccountName], [hfsCreatorCode],
// [hfsTypeCode], [immutable], [modificationDate], [ownerAccountID],
// [ownerAccountName], [posixPermissions]. You can change single attributes or
// any combination of attributes; you need not specify keys for all
// attributes.
//
// path: The path of a file or directory.
//
// # Discussion
//
// As in the POSIX standard, the app either must own the file or directory or
// must be running as superuser for attribute changes to take effect. The
// method attempts to make all changes specified in attributes and ignores any
// rejection of an attempted modification. If the last component of the path
// is a symbolic link, the system traverses it.
//
// You must initialize the [posixPermissions] value with the code representing
// the POSIX file-permissions bit pattern. The system sets [hfsCreatorCode]
// and [hfsTypeCode] only when `path` specifies a file.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/setAttributes(_:ofItemAtPath:)
//
// [busy]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/busy
// [creationDate]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/creationDate
// [extensionHidden]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/extensionHidden
// [groupOwnerAccountID]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountID
// [groupOwnerAccountName]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountName
// [hfsCreatorCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
// [hfsTypeCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
// [immutable]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/immutable
// [modificationDate]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/modificationDate
// [ownerAccountID]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountID
// [ownerAccountName]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountName
// [posixPermissions]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/posixPermissions
//
// [hfsCreatorCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
// [hfsTypeCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
// [posixPermissions]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/posixPermissions
func (f FileManager) SetAttributesOfItemAtPathError(attributes INSDictionary, path string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setAttributes:ofItemAtPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the contents of the file at the specified path.
//
// path: The path of the file whose contents you want.
//
// # Return Value
//
// An [NSData] object with the contents of the file. If `path` specifies a
// directory, or if some other error occurs, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/contents(atPath:)
func (f FileManager) ContentsAtPath(path string) INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("contentsAtPath:"), objc.String(path))
	return NSDataFromID(rv)
}

// Returns a Boolean value that indicates whether the files or directories in
// specified paths have the same contents.
//
// path1: The path of a file or directory to compare with the contents of `path2`.
//
// path2: The path of a file or directory to compare with the contents of `path1`.
//
// # Return Value
//
// true if file or directory specified in `path1` has the same contents as
// that specified in `path2`, otherwise false.
//
// # Discussion
//
// If `path1` and `path2` are directories, the contents are the list of files
// and subdirectories each contains—contents of subdirectories are also
// compared. For files, this method checks to see if they’re the same file,
// then compares their size, and finally compares their contents. This method
// does not traverse symbolic links, but compares the links themselves.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/contentsEqual(atPath:andPath:)
func (f FileManager) ContentsEqualAtPathAndPath(path1 string, path2 string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("contentsEqualAtPath:andPath:"), objc.String(path1), objc.String(path2))
	return rv
}

// Determines the type of relationship that exists between a directory and an
// item.
//
// outRelationship: A pointer to a variable in which to put the relationship between
// `directoryURL` and `otherURL`. For a list of possible values, see
// [FileManager.URLRelationship].
//
// directoryURL: The URL of the directory that potentially contains the item in `otherURL`.
// The URL in this parameter must specify a directory. This parameter must not
// be `nil`.
//
// otherURL: The URL of the file or directory whose relationship to `directoryURL` is
// being tested. This parameter must not be `nil`.
//
// # Discussion
//
// Use this method to determine the relationship between an item and a
// directory whose location you already know. If the relationship between the
// items is determined successfully, this method sets the value of the
// `outRelationship` parameter to an appropriate value. The directory may
// contain the item, it may be the same as the item, or it may not have a
// direct relationship to the item.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:ofDirectoryAt:toItemAt:)
//
// [FileManager.URLRelationship]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
func (f FileManager) GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship NSURLRelationship, directoryURL INSURL, otherURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("getRelationship:ofDirectoryAtURL:toItemAtURL:error:"), outRelationship, directoryURL, otherURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getRelationship:ofDirectoryAtURL:toItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines the type of relationship that exists between a system directory
// and the specified item.
//
// outRelationship: A pointer to a variable in which to put the relationship between
// `directoryURL` and `otherURL`. For a list of possible values, see
// [FileManager.URLRelationship].
//
// directory: The search path directory. For a list of possible values, see
// [FileManager.SearchPathDirectory].
//
// domainMask: The file system domain to search. Specify `0` for this parameter if you
// want the file manager to choose the domain that is most appropriate for the
// specified `url`.
//
// url: The URL of the file or directory whose relationship to `directoryURL` is
// being tested. This parameter must not be `nil`.
//
// # Discussion
//
// Use this method to determine the relationship between an item and one of
// the system-specific directories. For example, you might use this method to
// determine if the specified item is in the user’s [Documents] directory or
// is in the trash. If the relationship between the items is determined
// successfully, this method sets the value of the `outRelationship` parameter
// to an appropriate value. The directory may contain the item, it may be the
// same as the item, or it may not have a direct relationship to the item.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:of:in:toItemAt:)
//
// [FileManager.URLRelationship]: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
// [FileManager.SearchPathDirectory]: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
func (f FileManager) GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship NSURLRelationship, directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, url INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("getRelationship:ofDirectory:inDomain:toItemAtURL:error:"), outRelationship, directory, domainMask, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getRelationship:ofDirectory:inDomain:toItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a C-string representation of a given path that properly encodes
// Unicode strings for use by the file system.
//
// path: A string object containing a path to a file. This parameter must not be
// `nil` or contain the empty string.
//
// # Return Value
//
// A C-string representation of `path` that properly encodes Unicode strings
// for use by the file system.
//
// # Discussion
//
// Use this method if your code calls system routines that expect C-string
// path arguments. If you use the C string beyond the scope of the current
// autorelease pool, you must copy it.
//
// This method raises an exception if `path` is `nil` or contains the empty
// string. This method also throws an exception if the conversion of the
// string fails.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f FileManager) FileSystemRepresentationWithPath(path string) string {
	rv := objc.Send[*byte](f.ID, objc.Sel("fileSystemRepresentationWithPath:"), objc.String(path))
	return objc.GoString(rv)
}

// Returns an [NSString] object whose contents are derived from the specified
// C-string path.
//
// str: A C string representation of a pathname.
//
// len: The number of characters in `string`.
//
// # Return Value
//
// An [NSString] object converted from the C-string representation `string`
// with length `len` of a pathname in the current file system.
//
// # Discussion
//
// Use this method if your code receives paths as C strings from system
// routines.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f FileManager) StringWithFileSystemRepresentationLength(str string, len_ uint) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stringWithFileSystemRepresentation:length:"), unsafe.Pointer(unsafe.StringData(str+"\x00")), len_)
	return NSStringFromID(rv).String()
}

// Changes the path of the current working directory to the specified path.
//
// path: The path of the directory to which to change.
//
// # Return Value
//
// true if successful, otherwise false.
//
// # Discussion
//
// All relative pathnames refer implicitly to the current working directory.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/changeCurrentDirectoryPath(_:)
func (f FileManager) ChangeCurrentDirectoryPath(path string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("changeCurrentDirectoryPath:"), objc.String(path))
	return rv
}

// Starts the process of unmounting the specified volume.
//
// url: A file URL specifying the volume to be unmounted.
//
// mask: A bitmask of [FileManager.UnmountOptions] that you can use to customize the
// unmount operation’s behavior.
//
// completionHandler: A block executed when the unmount operation completes. The block receives
// an error parameter which is `nil` if unmounting was successful. Otherwise,
// it indicates why unmounting failed.
//
// # Discussion
//
// If the volume is encrypted, it is relocked after being unmounted.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/unmountVolume(at:options:completionHandler:)
//
// [FileManager.UnmountOptions]: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
func (f FileManager) UnmountVolumeAtURLOptionsCompletionHandler(url INSURL, mask NSFileManagerUnmountOptions, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("unmountVolumeAtURL:options:completionHandler:"), url, mask, _block2)
}

// Returns a directory enumerator object that can be used to perform a deep
// enumeration of the directory at the specified URL.
//
// url: The location of the directory for which you want an enumeration. This URL
// must not be a symbolic link that points to the desired directory. You can
// use the [URLByResolvingSymlinksInPath] method to resolve any symlinks in
// the URL.
//
// keys: An array of keys that identify the properties that you want pre-fetched for
// each item in the enumeration. The values for these keys are cached in the
// corresponding [NSURL] objects. You may specify `nil` for this parameter.
// For a list of keys you can specify, see [NSURLResourceKey].
//
// mask: Options for the enumeration. For a list of valid options, see
// [FileManager.DirectoryEnumerationOptions].
//
// handler: An optional error handler block for the file manager to call when an error
// occurs. The handler block should return true if you want the enumeration to
// continue or false if you want the enumeration to stop. The block takes the
// following parameters:
//
// url: An [NSURL] object that identifies the item for which the error
// occurred. error: An [NSError] object that contains information about the
// error.
//
// If you specify `nil` for this parameter, the enumerator object continues to
// enumerate items as if you had specified a block that returned true.
//
// # Return Value
//
// An directory enumerator object that enumerates the contents of the
// directory at `url`. If `url` is a filename, the method returns an
// enumerator object that enumerates no files—the first call to [NextObject]
// returns `nil`.
//
// # Discussion
//
// Because the enumeration is deep—that is, it lists the contents of all
// subdirectories—this enumerator object is useful for performing actions
// that involve large file-system subtrees. If the method is passed a
// directory on which another file system is mounted (a mount point), it
// traverses the mount point. This method does not resolve symbolic links or
// mount points encountered in the enumeration process, nor does it recurse
// through them if they point to a directory.
//
// For example, if you pass a URL that points to
// `/Volumes/MyMountedFileSystem`, the returned enumerator will include the
// entire directory structure for the file system mounted at that location.
// If, on the other hand, you pass `/Volumes`, the returned enumerator will
// include `/Volumes/MyMountedFileSystem` as one of its results, but will not
// traverse into the file system mounted there.
//
// The [NSDirectoryEnumerator] class has methods for skipping descendants of
// the existing path and for returning the number of levels deep the current
// object is in the directory hierarchy being enumerated (where the directory
// passed to [EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler] is
// considered to be level 0).
//
// This code fragment enumerates a URL and its subdirectories, collecting the
// URLs of files (skips directories). It also demonstrates how to ignore the
// contents of specified directories, in this case directories named
// “`_extras`”.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
//
// [FileManager.DirectoryEnumerationOptions]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
func (f FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url INSURL, keys []string, mask NSDirectoryEnumerationOptions, handler URLErrorHandler) INSDirectoryEnumerator {
	_block3, _ := NewURLErrorBlock(handler)
	rv := objc.Send[objc.ID](f.ID, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), url, keys, mask, _block3)
	return NSDirectoryEnumeratorFromID(rv)
}

// The home directory for the current user.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectoryForCurrentUser
func (f FileManager) HomeDirectoryForCurrentUser() INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("homeDirectoryForCurrentUser"))
	return NSURLFromID(objc.ID(rv))
}

// The temporary directory for the current user.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/temporaryDirectory
func (f FileManager) TemporaryDirectory() INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("temporaryDirectory"))
	return NSURLFromID(objc.ID(rv))
}

// An opaque token that represents the current user’s iCloud Drive Documents
// identity.
//
// # Discussion
//
// In iCloud Drive Documents, when iCloud is available, this property contains
// an opaque object representing the identity of the current user. If iCloud
// is unavailable or there is no logged-in user, the value of this property is
// `nil`. Accessing the value of this property is relatively fast, so you can
// check the value at launch time from your app’s main thread.
//
// You can use the token in this property, together with the
// [NSUbiquityIdentityDidChange] notification, to detect when the user logs in
// or out of iCloud and to detect changes to the active iCloud account. When
// the user logs in with a different iCloud account, the identity token
// changes, and the system posts the notification. If you stored or archived
// the previous token, compare that token to the newly obtained one using the
// [isEqual(_:)] method to determine if the users are the same or different.
//
// Accessing the token in this property doesn’t connect your app to its
// ubiquity containers. To establish access to a ubiquity container, call the
// [URLForUbiquityContainerIdentifier] method. In macOS, you can instead use
// an [NSDocument] object, which establishes access automatically.
//
// CloudKit clients should not use this token as a way to identify whether the
// iCloud account is logged in. Instead, use
// [accountStatus(completionHandler:)] or
// [fetchUserRecordID(completionHandler:)].
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/ubiquityIdentityToken
//
// [NSDocument]: https://developer.apple.com/documentation/AppKit/NSDocument
// [accountStatus(completionHandler:)]: https://developer.apple.com/documentation/CloudKit/CKContainer/accountStatus(completionHandler:)
// [fetchUserRecordID(completionHandler:)]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchUserRecordID(completionHandler:)
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
func (f FileManager) UbiquityIdentityToken() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("ubiquityIdentityToken"))
	return objectivec.Object{ID: rv}
}

// The delegate of the file manager object.
//
// # Discussion
//
// It is recommended that you assign a delegate to the file manager object
// only if you allocated and initialized the object yourself. Avoid assigning
// a delegate to the shared file manager obtained from the [DefaultManager]
// method.
//
// The default value of this property is `nil`. When assigning a delegate to
// this property, your object must conform to the [NSFileManagerDelegate]
// protocol.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/delegate
func (f FileManager) Delegate() NSFileManagerDelegate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("delegate"))
	return NSFileManagerDelegateObjectFromID(rv)
}
func (f FileManager) SetDelegate(value NSFileManagerDelegate) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelegate:"), value)
}

// The path to the program’s current directory.
//
// # Discussion
//
// The current directory path is the starting point for any relative paths you
// specify. For example, if the current directory is `/tmp` and you specify a
// relative pathname of `reports/info.Txt()`, the resulting full path for the
// item is `/tmp/reports/info.Txt()`.
//
// When an app is launched, this property is initially set to the app’s
// current working directory. If the current working directory is not
// accessible for any reason, the value of this property is `nil`. You can
// change the value of this property by calling the
// [ChangeCurrentDirectoryPath] method.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/currentDirectoryPath
func (f FileManager) CurrentDirectoryPath() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("currentDirectoryPath"))
	return NSStringFromID(rv).String()
}

// The process identifier of the process that prevented a volume from
// unmounting.
//
// See: https://developer.apple.com/documentation/foundation/nsfilemanagerunmountdissentingprocessidentifiererrorkey
func (f FileManager) NSFileManagerUnmountDissentingProcessIdentifierErrorKey() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFileManagerUnmountDissentingProcessIdentifierErrorKey"))
	return NSStringFromID(rv).String()
}

// The version of the Foundation framework in which
//
// See: https://developer.apple.com/documentation/foundation/nsfoundationversionwithfilemanagerresourceforksupport
func (f FileManager) NSFoundationVersionWithFileManagerResourceForkSupport() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFoundationVersionWithFileManagerResourceForkSupport"))
	return objectivec.Object{ID: rv}
}
func (f FileManager) SetNSFoundationVersionWithFileManagerResourceForkSupport(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setNSFoundationVersionWithFileManagerResourceForkSupport:"), value)
}

// Sent after the iCloud (“ubiquity”) identity has changed.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsubiquityidentitydidchange
func (f FileManager) NSUbiquityIdentityDidChange() NSNotificationName {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSUbiquityIdentityDidChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// The shared file manager object for the process.
//
// # Discussion
//
// This method always represents the same file manager object. If you plan to
// use a delegate with the file manager to receive notifications about the
// completion of file-based operations, you should create a new instance of
// [NSFileManager] rather than using the shared object.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/default
func (_FileManagerClass FileManagerClass) DefaultManager() FileManager {
	rv := objc.Send[objc.ID](objc.ID(_FileManagerClass.class), objc.Sel("defaultManager"))
	return NSFileManagerFromID(objc.ID(rv))
}

// PauseSyncForUbiquitousItemAtURL is a synchronous wrapper around [FileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FileManager) PauseSyncForUbiquitousItemAtURL(ctx context.Context, url INSURL) error {
	done := make(chan error, 1)
	f.PauseSyncForUbiquitousItemAtURLCompletionHandler(url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ResumeSyncForUbiquitousItemAtURLWithBehavior is a synchronous wrapper around [FileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FileManager) ResumeSyncForUbiquitousItemAtURLWithBehavior(ctx context.Context, url INSURL, behavior NSFileManagerResumeSyncBehavior) error {
	done := make(chan error, 1)
	f.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url, behavior, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FetchLatestRemoteVersionOfItemAtURL is a synchronous wrapper around [FileManager.FetchLatestRemoteVersionOfItemAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FileManager) FetchLatestRemoteVersionOfItemAtURL(ctx context.Context, url INSURL) (*NSFileVersion, error) {
	type result struct {
		val *NSFileVersion
		err error
	}
	done := make(chan result, 1)
	f.FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url, func(val *NSFileVersion, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicy is a synchronous wrapper around [FileManager.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicy(ctx context.Context, url INSURL, conflictResolutionPolicy NSFileManagerUploadLocalVersionConflictPolicy) (*NSFileVersion, error) {
	type result struct {
		val *NSFileVersion
		err error
	}
	done := make(chan result, 1)
	f.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url, conflictResolutionPolicy, func(val *NSFileVersion, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// UnmountVolumeAtURLOptions is a synchronous wrapper around [FileManager.UnmountVolumeAtURLOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FileManager) UnmountVolumeAtURLOptions(ctx context.Context, url INSURL, mask NSFileManagerUnmountOptions) error {
	done := make(chan error, 1)
	f.UnmountVolumeAtURLOptionsCompletionHandler(url, mask, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
