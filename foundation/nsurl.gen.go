// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURL] class.
var (
	_NSURLClass     NSURLClass
	_NSURLClassOnce sync.Once
)

func getNSURLClass() NSURLClass {
	_NSURLClassOnce.Do(func() {
		_NSURLClass = NSURLClass{class: objc.GetClass("NSURL")}
	})
	return _NSURLClass
}

// GetNSURLClass returns the class object for NSURL.
func GetNSURLClass() NSURLClass {
	return getNSURLClass()
}

type NSURLClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLClass) Alloc() NSURL {
	rv := objc.Send[NSURL](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the location of a resource, such as an item on a
// remote server or the path to a local file.
//
// # Overview
// 
// In Swift, this object bridges to [URL]; use [NSURL] when you need reference
// semantics or other Foundation-specific behavior.
// 
// You can use URL objects to construct URLs and access their parts. For URLs
// that represent local files, you can also manipulate properties of those
// files directly, such as changing the file’s last modification date.
// Finally, you can pass URL objects to other APIs to retrieve the contents of
// those URLs. For example, you can use the [NSURLSession], [NSURLConnection],
// and [NSURLDownload] classes to access the contents of remote resources, as
// described in [URL Loading System].
// 
// URL objects are the preferred way to refer to local files. Most objects
// that read data from or write data to a file have methods that accept an
// [NSURL] object instead of a pathname as the file reference. For example,
// you can get the contents of a local file URL as an [NSString] object using
// the `NSString/init()-715fw` initializer, or as an [NSData] object using the
// `NSData/init()-5abi3` initializer.
// 
// You can also use URLs for interapplication communication. In macOS, the
// [NSWorkspace] class provides the [open(_:)] method to open a location
// specified by a URL. Similarly, in iOS, the [UIApplication] class provides
// the [open(_:options:completionHandler:)] method.
// 
// Additionally, you can use URLs when working with pasteboards, as described
// in NSURL Additions Reference (part of the AppKit framework).
// 
// The [NSURL] class is “toll-free bridged” with its Core Foundation
// counterpart, [CFURL]. See [Toll-Free Bridging] for more information on
// toll-free bridging.
// 
// # Structure of a URL
// 
// An [NSURL] object is composed of two parts—a potentially `nil` base URL
// and a string that is resolved relative to the base URL. An [NSURL] object
// is considered absolute if its string part is fully resolved without a base;
// all other URLs are considered relative.
// 
// For example, when constructing an [NSURL] object, you might specify
// `///path/to/user/` as the base URL and `folder/file.Html()` as the string
// part, as follows:
// 
// When fully resolved, the absolute URL is
// `///path/to/user/folder/file.Html()`.
// 
// A URL can be also be divided into pieces based on its structure. For
// example, the URL
// `//p4ssw0rd@www.Example().443/script.Ext();param=value?query=value#ref`
// contains the following URL components:
// 
// [Table data omitted]
// 
// The [NSURL] class provides properties that let you examine each of these
// components.
// 
// For apps linked before iOS 17, the [NSURL] class parses URLs according to
// [RFC 1808], [RFC 1738], and [RFC 2732].
// 
// # Bookmarks and Security Scope
// 
// Starting with OS X v10.6 and iOS 4.0, the [NSURL] class provides a facility
// for creating and using bookmark objects. A provides a persistent reference
// to a file-system resource. When you resolve a bookmark, you obtain a URL to
// the resource’s current location. A bookmark’s association with a
// file-system resource (typically a file or folder) usually continues to work
// if the user moves or renames the resource, or if the user relaunches your
// app or restarts the system.
// 
// For a general introduction to using bookmarks, read [Locating Files Using
// Bookmarks] in [File System Programming Guide].
// 
// In a macOS app that adopts App Sandbox, you can use to gain access to
// file-system resources outside your app’s sandbox. These bookmarks
// preserve the user’s intent to give your app access to a resource across
// app launches. For details on how this works, including information on the
// entitlements you need in your Xcode project, read [Security-Scoped
// Bookmarks and Persistent Resource Access] in [App Sandbox Design Guide].
// The methods for using security-scoped bookmarks are described in this
// document in Working with Bookmark Data.
// 
// When you resolve a security-scoped bookmark, you get a security-scoped URL.
// 
// # Security-Scoped URLs
// 
// Security-scoped URLs provide access to resources outside an app’s
// sandbox. In macOS, you get access to security-scoped URLs when you resolve
// a security-scoped bookmark. In iOS, apps that or documents using a
// [UIDocumentPickerViewController] also receive security-scoped URLs.
// 
// To gain access to a security-scoped URL, you must call the
// [NSURL.StartAccessingSecurityScopedResource] method (or its Core Foundation
// equivalent, the [CFURLStartAccessingSecurityScopedResource(_:)] function).
// For iOS apps, if you use a [UIDocument] to access the URL, it automatically
// manages the security-scoped URL for you.
// 
// If `startAccessingSecurityScopedResource` (or
// [CFUrLStartAccessingSecurityScopedResource]) returns [true], you must
// relinquish your access by calling the [NSURL.StopAccessingSecurityScopedResource]
// method (or its Core Foundation equivalent, the
// [CFURLStopAccessingSecurityScopedResource(_:)] function). You should
// relinquish your access as soon as you have finished using the file. After
// you call these methods, you immediately lose access to the resource in
// question.
// 
// # Security-Scoped URLs and String Paths
// 
// In a macOS app, when you copy a security-scoped URL, the copy has the
// security scope of the original. You gain access to the file-system resource
// (that the URL points to) just as you would with the original URL: by
// calling the [NSURL.StartAccessingSecurityScopedResource] method (or its Core
// Foundation equivalent).
// 
// If you need a security-scoped URL’s path as a string value (as provided
// by the [NSURL.Path] method), such as to provide to an API that requires a string
// value, obtain the path from the URL as needed. Note, however, that a
// string-based path obtained from a security-scoped URL have security scope
// and you cannot use that string to obtain access to a security-scoped
// resource.
// 
// # iCloud Document Thumbnails
// 
// With OS X v10.10 and iOS 8.0, the NSURL class includes the ability to get
// and set document thumbnails as a resource property for iCloud documents.
// You can get a dictionary of [NSImage] objects in macOS or [UIImage] objects
// in iOS using the [NSURL.GetResourceValueForKeyError] or
// [NSURL.GetPromisedItemResourceValueForKeyError] methods.
// 
// In macOS, you can set a dictionary of thumbnails using the
// [NSURL.SetResourceValueForKeyError] method. You can also get or set all the
// thumbnails as an [NSImage] object with multiple representations by using
// the [thumbnailKey].
//
// [App Sandbox Design Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/AppSandboxDesignGuide/AboutAppSandbox/AboutAppSandbox.html#//apple_ref/doc/uid/TP40011183
// [CFURLStartAccessingSecurityScopedResource(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
// [CFURLStopAccessingSecurityScopedResource(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
// [CFURL]: https://developer.apple.com/documentation/CoreFoundation/CFURL
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
// [Locating Files Using Bookmarks]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/AccessingFilesandDirectories/AccessingFilesandDirectories.html#//apple_ref/doc/uid/TP40010672-CH3-SW10
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [NSWorkspace]: https://developer.apple.com/documentation/AppKit/NSWorkspace
// [RFC 1738]: https://tools.ietf.org/html/rfc1738
// [RFC 1808]: https://tools.ietf.org/html/rfc1808
// [RFC 2732]: https://tools.ietf.org/html/rfc2732
// [Security-Scoped Bookmarks and Persistent Resource Access]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/AppSandboxDesignGuide/AppSandboxInDepth/AppSandboxInDepth.html#//apple_ref/doc/uid/TP40011183-CH3-SW16
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [UIApplication]: https://developer.apple.com/documentation/UIKit/UIApplication
// [UIDocumentPickerViewController]: https://developer.apple.com/documentation/UIKit/UIDocumentPickerViewController
// [UIDocument]: https://developer.apple.com/documentation/UIKit/UIDocument
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
// [URL Loading System]: https://developer.apple.com/documentation/Foundation/url-loading-system
// [URL]: https://developer.apple.com/documentation/Foundation/URL
// [open(_:)]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
// [open(_:options:completionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplication/open(_:options:completionHandler:)
// [thumbnailKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/thumbnailKey
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a URL object
//
//   - [NSURL.InitWithString]: Initializes an NSURL object with a provided URL string.
//   - [NSURL.InitWithStringEncodingInvalidCharacters]: Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//   - [NSURL.InitWithStringRelativeToURL]: Initializes an NSURL object with a base URL and a relative string.
//   - [NSURL.InitFileURLWithPathIsDirectory]: Initializes a newly created NSURL referencing the local file or directory at `path`.
//   - [NSURL.InitFileURLWithPathRelativeToURL]
//   - [NSURL.InitFileURLWithPathIsDirectoryRelativeToURL]
//   - [NSURL.InitFileURLWithPath]: Initializes a newly created NSURL referencing the local file or directory at `path`.
//   - [NSURL.InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError]: Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//   - [NSURL.GetFileSystemRepresentationMaxLength]: Fills the provided buffer with a C string representing a local file system path.
//   - [NSURL.InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL]: Initializes a URL object with a C string representing a local file system path.
//   - [NSURL.InitAbsoluteURLWithDataRepresentationRelativeToURL]
//   - [NSURL.InitWithDataRepresentationRelativeToURL]
//   - [NSURL.DataRepresentation]
//
// # Querying an NSURL
//
//   - [NSURL.CheckResourceIsReachableAndReturnError]: Returns whether the resource pointed to by a file URL can be reached.
//   - [NSURL.IsFileReferenceURL]: Returns whether the URL is a file reference URL.
//   - [NSURL.FileURL]: A boolean value that determines whether the receiver is a file URL.
//
// # Accessing the Parts of the URL
//
//   - [NSURL.AbsoluteString]: The URL string for the receiver as an absolute URL. (read-only)
//   - [NSURL.AbsoluteURL]: An absolute URL that refers to the same resource as the receiver. (read-only)
//   - [NSURL.BaseURL]: The base URL. (read-only)
//   - [NSURL.FileSystemRepresentation]: A C string containing the URL’s file system path. (read-only)
//   - [NSURL.Fragment]: The fragment identifier, conforming to RFC 1808. (read-only)
//   - [NSURL.Host]: The host, conforming to RFC 1808. (read-only)
//   - [NSURL.LastPathComponent]: The last path component. (read-only)
//   - [NSURL.Password]: The password conforming to RFC 1808. (read-only)
//   - [NSURL.Path]: The path, conforming to RFC 1808. (read-only)
//   - [NSURL.PathComponents]: An array containing the  path components. (read-only)
//   - [NSURL.PathExtension]: The path extension. (read-only)
//   - [NSURL.Port]: The port, conforming to RFC 1808.
//   - [NSURL.Query]: The query string, conforming to RFC 1808.
//   - [NSURL.RelativePath]: The relative path, conforming to RFC 1808. (read-only)
//   - [NSURL.RelativeString]: A string representation of the relative portion of the URL. (read-only)
//   - [NSURL.ResourceSpecifier]: The resource specifier. (read-only)
//   - [NSURL.Scheme]: The scheme. (read-only)
//   - [NSURL.StandardizedURL]: A copy of the URL with any instances of `".."` or `"."` removed from its path. (read-only)
//   - [NSURL.User]: The user name, conforming to RFC 1808.
//
// # Accessing Resource Values
//
//   - [NSURL.ResourceValuesForKeysError]: Returns the resource values for the properties identified by specified array of keys.
//   - [NSURL.GetResourceValueForKeyError]: Returns the value of the resource property for the specified key.
//   - [NSURL.SetResourceValueForKeyError]: Sets the URL’s resource property for a given key to a given value.
//   - [NSURL.SetResourceValuesError]: Sets the URL’s resource properties for a given set of keys to a given set of values.
//   - [NSURL.RemoveAllCachedResourceValues]: Removes all cached resource values and temporary resource values from the URL object.
//   - [NSURL.RemoveCachedResourceValueForKey]: Removes the cached resource value identified by a given key from the URL object.
//   - [NSURL.SetTemporaryResourceValueForKey]: Sets a temporary resource value on the URL object.
//
// # Modifying and Converting a File URL
//
//   - [NSURL.FilePathURL]: A file path URL that points to the same resource as the URL object. (read-only)
//   - [NSURL.FileReferenceURL]: Returns a new file reference URL that points to the same resource as the receiver.
//   - [NSURL.URLByAppendingPathComponent]: Returns a new URL by appending a path component to the original URL.
//   - [NSURL.URLByAppendingPathComponentIsDirectory]: Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//   - [NSURL.URLByAppendingPathComponentConformingToType]: Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//   - [NSURL.URLByAppendingPathExtension]: Returns a new URL by appending a path extension to the original URL.
//   - [NSURL.URLByAppendingPathExtensionForType]: Returns a URL by appending the path extension for a uniform type identifier.
//   - [NSURL.URLByDeletingLastPathComponent]: A URL you create by removing the last path component from the receiver. (read-only)
//   - [NSURL.URLByDeletingPathExtension]: A URL you create by removing the path extension from the receiver, if any. (read-only)
//   - [NSURL.URLByResolvingSymlinksInPath]: A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//   - [NSURL.URLByStandardizingPath]: A URL that points to the same resource as the original URL using an absolute path. (read-only)
//   - [NSURL.HasDirectoryPath]: A Boolean value that indicates whether the URL string’s path represents a directory.
//
// # Working with Bookmark Data
//
//   - [NSURL.BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError]: Returns a bookmark for the URL, created with specified options and resource values.
//   - [NSURL.StartAccessingSecurityScopedResource]: In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//   - [NSURL.StopAccessingSecurityScopedResource]: In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// # Working with Promised Items
//
//   - [NSURL.CheckPromisedItemIsReachableAndReturnError]: Returns whether the promised item can be reached.
//   - [NSURL.GetPromisedItemResourceValueForKeyError]: Returns the value of the resource property for the specified key.
//   - [NSURL.PromisedItemResourceValuesForKeysError]: Returns the resource values for the properties identified by specified array of keys.
//
// # Working with Pasteboards
//
//   - [NSURL.WriteToPasteboard]: Writes the URL to the specified pasteboard.
//
// # Using Quick Looks
//
//   - [NSURL.CustomPlaygroundQuickLook]: A custom playground Quick Look for this instance.
//   - [NSURL.SetCustomPlaygroundQuickLook]
//
// See: https://developer.apple.com/documentation/Foundation/NSURL
type NSURL struct {
	objectivec.Object
}

// NSURLFromID constructs a [NSURL] from an objc.ID.
//
// An object that represents the location of a resource, such as an item on a
// remote server or the path to a local file.
func NSURLFromID(id objc.ID) NSURL {
	return NSURL{objectivec.Object{ID: id}}
}
// NOTE: NSURL adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURL] class.
//
// # Creating a URL object
//
//   - [INSURL.InitWithString]: Initializes an NSURL object with a provided URL string.
//   - [INSURL.InitWithStringEncodingInvalidCharacters]: Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//   - [INSURL.InitWithStringRelativeToURL]: Initializes an NSURL object with a base URL and a relative string.
//   - [INSURL.InitFileURLWithPathIsDirectory]: Initializes a newly created NSURL referencing the local file or directory at `path`.
//   - [INSURL.InitFileURLWithPathRelativeToURL]
//   - [INSURL.InitFileURLWithPathIsDirectoryRelativeToURL]
//   - [INSURL.InitFileURLWithPath]: Initializes a newly created NSURL referencing the local file or directory at `path`.
//   - [INSURL.InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError]: Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//   - [INSURL.GetFileSystemRepresentationMaxLength]: Fills the provided buffer with a C string representing a local file system path.
//   - [INSURL.InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL]: Initializes a URL object with a C string representing a local file system path.
//   - [INSURL.InitAbsoluteURLWithDataRepresentationRelativeToURL]
//   - [INSURL.InitWithDataRepresentationRelativeToURL]
//   - [INSURL.DataRepresentation]
//
// # Querying an NSURL
//
//   - [INSURL.CheckResourceIsReachableAndReturnError]: Returns whether the resource pointed to by a file URL can be reached.
//   - [INSURL.IsFileReferenceURL]: Returns whether the URL is a file reference URL.
//   - [INSURL.FileURL]: A boolean value that determines whether the receiver is a file URL.
//
// # Accessing the Parts of the URL
//
//   - [INSURL.AbsoluteString]: The URL string for the receiver as an absolute URL. (read-only)
//   - [INSURL.AbsoluteURL]: An absolute URL that refers to the same resource as the receiver. (read-only)
//   - [INSURL.BaseURL]: The base URL. (read-only)
//   - [INSURL.FileSystemRepresentation]: A C string containing the URL’s file system path. (read-only)
//   - [INSURL.Fragment]: The fragment identifier, conforming to RFC 1808. (read-only)
//   - [INSURL.Host]: The host, conforming to RFC 1808. (read-only)
//   - [INSURL.LastPathComponent]: The last path component. (read-only)
//   - [INSURL.Password]: The password conforming to RFC 1808. (read-only)
//   - [INSURL.Path]: The path, conforming to RFC 1808. (read-only)
//   - [INSURL.PathComponents]: An array containing the  path components. (read-only)
//   - [INSURL.PathExtension]: The path extension. (read-only)
//   - [INSURL.Port]: The port, conforming to RFC 1808.
//   - [INSURL.Query]: The query string, conforming to RFC 1808.
//   - [INSURL.RelativePath]: The relative path, conforming to RFC 1808. (read-only)
//   - [INSURL.RelativeString]: A string representation of the relative portion of the URL. (read-only)
//   - [INSURL.ResourceSpecifier]: The resource specifier. (read-only)
//   - [INSURL.Scheme]: The scheme. (read-only)
//   - [INSURL.StandardizedURL]: A copy of the URL with any instances of `".."` or `"."` removed from its path. (read-only)
//   - [INSURL.User]: The user name, conforming to RFC 1808.
//
// # Accessing Resource Values
//
//   - [INSURL.ResourceValuesForKeysError]: Returns the resource values for the properties identified by specified array of keys.
//   - [INSURL.GetResourceValueForKeyError]: Returns the value of the resource property for the specified key.
//   - [INSURL.SetResourceValueForKeyError]: Sets the URL’s resource property for a given key to a given value.
//   - [INSURL.SetResourceValuesError]: Sets the URL’s resource properties for a given set of keys to a given set of values.
//   - [INSURL.RemoveAllCachedResourceValues]: Removes all cached resource values and temporary resource values from the URL object.
//   - [INSURL.RemoveCachedResourceValueForKey]: Removes the cached resource value identified by a given key from the URL object.
//   - [INSURL.SetTemporaryResourceValueForKey]: Sets a temporary resource value on the URL object.
//
// # Modifying and Converting a File URL
//
//   - [INSURL.FilePathURL]: A file path URL that points to the same resource as the URL object. (read-only)
//   - [INSURL.FileReferenceURL]: Returns a new file reference URL that points to the same resource as the receiver.
//   - [INSURL.URLByAppendingPathComponent]: Returns a new URL by appending a path component to the original URL.
//   - [INSURL.URLByAppendingPathComponentIsDirectory]: Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//   - [INSURL.URLByAppendingPathComponentConformingToType]: Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//   - [INSURL.URLByAppendingPathExtension]: Returns a new URL by appending a path extension to the original URL.
//   - [INSURL.URLByAppendingPathExtensionForType]: Returns a URL by appending the path extension for a uniform type identifier.
//   - [INSURL.URLByDeletingLastPathComponent]: A URL you create by removing the last path component from the receiver. (read-only)
//   - [INSURL.URLByDeletingPathExtension]: A URL you create by removing the path extension from the receiver, if any. (read-only)
//   - [INSURL.URLByResolvingSymlinksInPath]: A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//   - [INSURL.URLByStandardizingPath]: A URL that points to the same resource as the original URL using an absolute path. (read-only)
//   - [INSURL.HasDirectoryPath]: A Boolean value that indicates whether the URL string’s path represents a directory.
//
// # Working with Bookmark Data
//
//   - [INSURL.BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError]: Returns a bookmark for the URL, created with specified options and resource values.
//   - [INSURL.StartAccessingSecurityScopedResource]: In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//   - [INSURL.StopAccessingSecurityScopedResource]: In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// # Working with Promised Items
//
//   - [INSURL.CheckPromisedItemIsReachableAndReturnError]: Returns whether the promised item can be reached.
//   - [INSURL.GetPromisedItemResourceValueForKeyError]: Returns the value of the resource property for the specified key.
//   - [INSURL.PromisedItemResourceValuesForKeysError]: Returns the resource values for the properties identified by specified array of keys.
//
// # Working with Pasteboards
//
//   - [INSURL.WriteToPasteboard]: Writes the URL to the specified pasteboard.
//
// # Using Quick Looks
//
//   - [INSURL.CustomPlaygroundQuickLook]: A custom playground Quick Look for this instance.
//   - [INSURL.SetCustomPlaygroundQuickLook]
//
// See: https://developer.apple.com/documentation/Foundation/NSURL
type INSURL interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSItemProviderReading
	NSItemProviderWriting
	NSSecureCoding

	// Topic: Creating a URL object

	// Initializes an NSURL object with a provided URL string.
	InitWithString(URLString string) NSURL
	// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
	InitWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURL
	// Initializes an NSURL object with a base URL and a relative string.
	InitWithStringRelativeToURL(URLString string, baseURL INSURL) NSURL
	// Initializes a newly created NSURL referencing the local file or directory at `path`.
	InitFileURLWithPathIsDirectory(path string, isDir bool) NSURL
	InitFileURLWithPathRelativeToURL(path string, baseURL INSURL) NSURL
	InitFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL
	// Initializes a newly created NSURL referencing the local file or directory at `path`.
	InitFileURLWithPath(path string) NSURL
	// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
	InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData INSData, options NSURLBookmarkResolutionOptions, relativeURL INSURL, isStale unsafe.Pointer) (NSURL, error)
	// Fills the provided buffer with a C string representing a local file system path.
	GetFileSystemRepresentationMaxLength(buffer string, maxBufferLength uint) bool
	// Initializes a URL object with a C string representing a local file system path.
	InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL
	InitAbsoluteURLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL
	InitWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL
	DataRepresentation() INSData

	// Topic: Querying an NSURL

	// Returns whether the resource pointed to by a file URL can be reached.
	CheckResourceIsReachableAndReturnError() (bool, error)
	// Returns whether the URL is a file reference URL.
	IsFileReferenceURL() bool
	// A boolean value that determines whether the receiver is a file URL.
	FileURL() bool

	// Topic: Accessing the Parts of the URL

	// The URL string for the receiver as an absolute URL. (read-only)
	AbsoluteString() string
	// An absolute URL that refers to the same resource as the receiver. (read-only)
	AbsoluteURL() INSURL
	// The base URL. (read-only)
	BaseURL() INSURL
	// A C string containing the URL’s file system path. (read-only)
	FileSystemRepresentation() string
	// The fragment identifier, conforming to RFC 1808. (read-only)
	Fragment() string
	// The host, conforming to RFC 1808. (read-only)
	Host() string
	// The last path component. (read-only)
	LastPathComponent() string
	// The password conforming to RFC 1808. (read-only)
	Password() string
	// The path, conforming to RFC 1808. (read-only)
	Path() string
	// An array containing the  path components. (read-only)
	PathComponents() []string
	// The path extension. (read-only)
	PathExtension() string
	// The port, conforming to RFC 1808.
	Port() INSNumber
	// The query string, conforming to RFC 1808.
	Query() string
	// The relative path, conforming to RFC 1808. (read-only)
	RelativePath() string
	// A string representation of the relative portion of the URL. (read-only)
	RelativeString() string
	// The resource specifier. (read-only)
	ResourceSpecifier() string
	// The scheme. (read-only)
	Scheme() string
	// A copy of the URL with any instances of `".."` or `"."` removed from its path. (read-only)
	StandardizedURL() INSURL
	// The user name, conforming to RFC 1808.
	User() string

	// Topic: Accessing Resource Values

	// Returns the resource values for the properties identified by specified array of keys.
	ResourceValuesForKeysError(keys []string) (INSDictionary, error)
	// Returns the value of the resource property for the specified key.
	GetResourceValueForKeyError(value []objectivec.IObject, key NSURLResourceKey) (bool, error)
	// Sets the URL’s resource property for a given key to a given value.
	SetResourceValueForKeyError(value objectivec.IObject, key NSURLResourceKey) (bool, error)
	// Sets the URL’s resource properties for a given set of keys to a given set of values.
	SetResourceValuesError(keyedValues INSDictionary) (bool, error)
	// Removes all cached resource values and temporary resource values from the URL object.
	RemoveAllCachedResourceValues()
	// Removes the cached resource value identified by a given key from the URL object.
	RemoveCachedResourceValueForKey(key NSURLResourceKey)
	// Sets a temporary resource value on the URL object.
	SetTemporaryResourceValueForKey(value objectivec.IObject, key NSURLResourceKey)

	// Topic: Modifying and Converting a File URL

	// A file path URL that points to the same resource as the URL object. (read-only)
	FilePathURL() INSURL
	// Returns a new file reference URL that points to the same resource as the receiver.
	FileReferenceURL() INSURL
	// Returns a new URL by appending a path component to the original URL.
	URLByAppendingPathComponent(pathComponent string) INSURL
	// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
	URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) INSURL
	// Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
	URLByAppendingPathComponentConformingToType(partialName string, contentType objectivec.IObject) INSURL
	// Returns a new URL by appending a path extension to the original URL.
	URLByAppendingPathExtension(pathExtension string) INSURL
	// Returns a URL by appending the path extension for a uniform type identifier.
	URLByAppendingPathExtensionForType(contentType objectivec.IObject) INSURL
	// A URL you create by removing the last path component from the receiver. (read-only)
	URLByDeletingLastPathComponent() INSURL
	// A URL you create by removing the path extension from the receiver, if any. (read-only)
	URLByDeletingPathExtension() INSURL
	// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
	URLByResolvingSymlinksInPath() INSURL
	// A URL that points to the same resource as the original URL using an absolute path. (read-only)
	URLByStandardizingPath() INSURL
	// A Boolean value that indicates whether the URL string’s path represents a directory.
	HasDirectoryPath() bool

	// Topic: Working with Bookmark Data

	// Returns a bookmark for the URL, created with specified options and resource values.
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options NSURLBookmarkCreationOptions, keys []string, relativeURL INSURL) (INSData, error)
	// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
	StartAccessingSecurityScopedResource() bool
	// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
	StopAccessingSecurityScopedResource()

	// Topic: Working with Promised Items

	// Returns whether the promised item can be reached.
	CheckPromisedItemIsReachableAndReturnError() (bool, error)
	// Returns the value of the resource property for the specified key.
	GetPromisedItemResourceValueForKeyError(value []objectivec.IObject, key NSURLResourceKey) (bool, error)
	// Returns the resource values for the properties identified by specified array of keys.
	PromisedItemResourceValuesForKeysError(keys []string) (INSDictionary, error)

	// Topic: Working with Pasteboards

	// Writes the URL to the specified pasteboard.
	WriteToPasteboard(pasteBoard objectivec.IObject)

	// Topic: Using Quick Looks

	// A custom playground Quick Look for this instance.
	CustomPlaygroundQuickLook() objectivec.IObject
	SetCustomPlaygroundQuickLook(value objectivec.IObject)
}

// Init initializes the instance.
func (u NSURL) Init() NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURL) Autorelease() NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURL creates a new NSURL instance.
func NewNSURL() NSURL {
	class := getNSURLClass()
	rv := objc.Send[NSURL](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return NSURLFromID(rv)
}

// Returns a new URL made by resolving the alias file at `url`.
//
// url: The URL pointing to the alias file.
//
// options: Options taken into account when resolving the bookmark data. The
// [URLBookmarkResolutionWithSecurityScope] option is not supported by this
// method.
//
// # Return Value
// 
// A new URL created by resolving the bookmark data derived from the provided
// alias file. If an error occurs, this method returns `nil`.
//
// # Discussion
// 
// Creates and initializes a new URL based on the alias file at `url`. Use
// this method to resolve bookmark data that was saved using
// [WriteBookmarkDataToURLOptionsError] and resolves that data in one step.
// 
// If the `url` argument does not refer to an alias file as defined by the
// [NSURLIsAliasFileKey] property, this method returns the `url` argument.
// 
// If the `url` argument is unreachable, this method returns `nil` and the
// optional error argument is populated.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func NewURLByResolvingAliasFileAtURLOptionsError(url INSURL, options NSURLBookmarkResolutionOptions) (NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getNSURLClass().class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil
}

// Initializes a newly created NSURL that points to a location specified by
// resolving bookmark data.
//
// bookmarkData: The bookmark data the URL is derived from.
//
// options: Options taken into account when resolving the bookmark data.
//
// relativeURL: The base URL that the bookmark data is relative to.
//
// isStale: If [true], the bookmark data is stale.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An NSURL initialized by resolving `bookmarkData`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData INSData, options NSURLBookmarkResolutionOptions, relativeURL INSURL, isStale unsafe.Pointer) (NSURL, error) {
	var errorPtr objc.ID
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil
}

// Initializes a URL object with a C string representing a local file system
// path.
//
// path: A null-terminated C string in file system representation containing the
// path to represent as a URL. If this path is a relative path, it is treated
// as being relative to the current working directory.
//
// isDir: [true] if the last path part is a directory, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// baseURL: The base URL for the new URL object. This must be a file URL. If `path` is
// absolute, this URL is ignored.
//
// # Return Value
// 
// Returns the initialized object.
//
// # Discussion
// 
// The file system representation format is described in File Encodings and
// Fonts.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), unsafe.Pointer(unsafe.StringData(path + "\x00")), isDir, baseURL)
	return NSURLFromID(rv)
}

// Initializes a newly created NSURL referencing the local file or directory
// at `path`.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// # Discussion
// 
// Invoking this method is equivalent to invoking [init(scheme:host:path:)]
// with scheme [NSURLFileScheme], a `nil` host, and `path`.
// 
// This method assumes that `path` is a directory if it ends with a slash. If
// `path` does not end with a slash, the method examines the file system to
// determine if `path` is a file or a directory. If `path` exists in the file
// system and is a directory, the method appends a trailing slash. If `path`
// does not exist in the file system, the method assumes that it represents a
// file and does not append a trailing slash.
// 
// As an alternative, consider using [InitFileURLWithPathIsDirectory], which
// allows you to explicitly specify whether the returned [NSURL] object
// represents a file or directory.
//
// [NSURLFileScheme]: https://developer.apple.com/documentation/Foundation/NSURLFileScheme
// [init(scheme:host:path:)]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path string) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFileURLWithPath:"), objc.String(path))
	return NSURLFromID(rv)
}

// Initializes a newly created NSURL referencing the local file or directory
// at `path`.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// isDir: A Boolean value that specifies whether `path` is treated as a directory
// path when resolving against relative path components. Pass [true] if the
// `path` indicates a directory, [false] otherwise
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// # Discussion
// 
// Invoking this method is equivalent to invoking [init(scheme:host:path:)]
// with scheme [NSURLFileScheme], a `nil` host, and `path`.
//
// [NSURLFileScheme]: https://developer.apple.com/documentation/Foundation/NSURLFileScheme
// [init(scheme:host:path:)]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path string, isDir bool) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return NSURLFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return NSURLFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path string, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return NSURLFromID(rv)
}

// Reads an NSURL object off of the specified pasteboard.
//
// pasteBoard: The target pasteboard.
//
// # Return Value
// 
// A [NSURL] object, or `nil` if the pasteboard does not contain
// [NSURLPboardType] data.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
// pasteBoard is a [appkit.NSPasteboard].
func NewURLFromPasteboard(pasteBoard objectivec.IObject) NSURL {
	rv := objc.Send[objc.ID](objc.ID(getNSURLClass().class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return NSURLFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLWithCoder(coder INSCoder) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSURLFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func NewURLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, baseURL)
	return NSURLFromID(rv)
}

// Initializes an NSURL object with a provided URL string.
//
// URLString: The URL string with which to initialize the NSURL object. Linked on or
// after iOS 17, this method parses [URLString] according to RFC 3986. Linked
// before iOS 17, this method parses [URLString] according to RFCs 1738 and
// 1808.
//
// # Return Value
// 
// An NSURL object initialized with [URLString]. If the URL string was
// malformed, returns `nil`.
//
// # Discussion
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)
func NewURLWithString(URLString string) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	return NSURLFromID(rv)
}

// Creates an instance from the provided string, optionally IDNA- and
// percent-encoding any invalid characters.
//
// URLString: A URL location.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in `string`.
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the initializer returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func NewURLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return NSURLFromID(rv)
}

// Initializes an NSURL object with a base URL and a relative string.
//
// URLString: The URL string with which to initialize the NSURL object. Linked on or
// after iOS 17, this method parses [URLString] according to RFC 3986. Linked
// before iOS 17, this method parses [URLString] according to RFCs 1738 and
// 1808. [URLString] is interpreted relative to `baseURL`.
//
// baseURL: The base URL for the NSURL object.
//
// # Return Value
// 
// An NSURL object initialized with [URLString] and `baseURL`. If [URLString]
// was malformed, returns `nil`.
//
// # Discussion
// 
// This method allows you to create a URL relative to a base path or URL. For
// example, if you have the URL for a folder on disk and the name of a file
// within that folder, you can construct a URL for the file by providing the
// folder’s URL as the base path (with a trailing slash) and the filename as
// the string part.
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
// 
// [InitWithStringRelativeToURL] is the designated initializer for NSURL.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)
func NewURLWithStringRelativeToURL(URLString string, baseURL INSURL) NSURL {
	instance := getNSURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return NSURLFromID(rv)
}

// Initializes an NSURL object with a provided URL string.
//
// URLString: The URL string with which to initialize the NSURL object. Linked on or
// after iOS 17, this method parses [URLString] according to RFC 3986. Linked
// before iOS 17, this method parses [URLString] according to RFCs 1738 and
// 1808.
//
// # Return Value
// 
// An NSURL object initialized with [URLString]. If the URL string was
// malformed, returns `nil`.
//
// # Discussion
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)
func (u NSURL) InitWithString(URLString string) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initWithString:"), objc.String(URLString))
	return rv
}
// Creates an instance from the provided string, optionally IDNA- and
// percent-encoding any invalid characters.
//
// URLString: A URL location.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in `string`.
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the initializer returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func (u NSURL) InitWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}
// Initializes an NSURL object with a base URL and a relative string.
//
// URLString: The URL string with which to initialize the NSURL object. Linked on or
// after iOS 17, this method parses [URLString] according to RFC 3986. Linked
// before iOS 17, this method parses [URLString] according to RFCs 1738 and
// 1808. [URLString] is interpreted relative to `baseURL`.
//
// baseURL: The base URL for the NSURL object.
//
// # Return Value
// 
// An NSURL object initialized with [URLString] and `baseURL`. If [URLString]
// was malformed, returns `nil`.
//
// # Discussion
// 
// This method allows you to create a URL relative to a base path or URL. For
// example, if you have the URL for a folder on disk and the name of a file
// within that folder, you can construct a URL for the file by providing the
// folder’s URL as the base path (with a trailing slash) and the filename as
// the string part.
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
// 
// [InitWithStringRelativeToURL] is the designated initializer for NSURL.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)
func (u NSURL) InitWithStringRelativeToURL(URLString string, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return rv
}
// Initializes a newly created NSURL referencing the local file or directory
// at `path`.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// isDir: A Boolean value that specifies whether `path` is treated as a directory
// path when resolving against relative path components. Pass [true] if the
// `path` indicates a directory, [false] otherwise
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// # Discussion
// 
// Invoking this method is equivalent to invoking [init(scheme:host:path:)]
// with scheme [NSURLFileScheme], a `nil` host, and `path`.
//
// [NSURLFileScheme]: https://developer.apple.com/documentation/Foundation/NSURLFileScheme
// [init(scheme:host:path:)]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func (u NSURL) InitFileURLWithPathIsDirectory(path string, isDir bool) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initFileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func (u NSURL) InitFileURLWithPathRelativeToURL(path string, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func (u NSURL) InitFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return rv
}
// Initializes a newly created NSURL referencing the local file or directory
// at `path`.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// # Discussion
// 
// Invoking this method is equivalent to invoking [init(scheme:host:path:)]
// with scheme [NSURLFileScheme], a `nil` host, and `path`.
// 
// This method assumes that `path` is a directory if it ends with a slash. If
// `path` does not end with a slash, the method examines the file system to
// determine if `path` is a file or a directory. If `path` exists in the file
// system and is a directory, the method appends a trailing slash. If `path`
// does not exist in the file system, the method assumes that it represents a
// file and does not append a trailing slash.
// 
// As an alternative, consider using [InitFileURLWithPathIsDirectory], which
// allows you to explicitly specify whether the returned [NSURL] object
// represents a file or directory.
//
// [NSURLFileScheme]: https://developer.apple.com/documentation/Foundation/NSURLFileScheme
// [init(scheme:host:path:)]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func (u NSURL) InitFileURLWithPath(path string) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initFileURLWithPath:"), objc.String(path))
	return rv
}
// Initializes a newly created NSURL that points to a location specified by
// resolving bookmark data.
//
// bookmarkData: The bookmark data the URL is derived from.
//
// options: Options taken into account when resolving the bookmark data.
//
// relativeURL: The base URL that the bookmark data is relative to.
//
// isStale: If [true], the bookmark data is stale.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An NSURL initialized by resolving `bookmarkData`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func (u NSURL) InitByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData INSData, options NSURLBookmarkResolutionOptions, relativeURL INSURL, isStale unsafe.Pointer) (NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}
// Fills the provided buffer with a C string representing a local file system
// path.
//
// buffer: A buffer large enough to hold the path. On return, contains a
// null-terminated C string in file system representation.
//
// maxBufferLength: The size of `buffer` in bytes (typically [MAXPATHLEN] or `PATH_MAX`).
//
// # Return Value
// 
// Returns [true] if the URL could be converted into a file system
// representation, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The file system representation format is described in File Encodings and
// Fonts.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)
func (u NSURL) GetFileSystemRepresentationMaxLength(buffer string, maxBufferLength uint) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), unsafe.Pointer(unsafe.StringData(buffer + "\x00")), maxBufferLength)
	return rv
}
// Initializes a URL object with a C string representing a local file system
// path.
//
// path: A null-terminated C string in file system representation containing the
// path to represent as a URL. If this path is a relative path, it is treated
// as being relative to the current working directory.
//
// isDir: [true] if the last path part is a directory, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// baseURL: The base URL for the new URL object. This must be a file URL. If `path` is
// absolute, this URL is ignored.
//
// # Return Value
// 
// Returns the initialized object.
//
// # Discussion
// 
// The file system representation format is described in File Encodings and
// Fonts.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func (u NSURL) InitFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), unsafe.Pointer(unsafe.StringData(path + "\x00")), isDir, baseURL)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func (u NSURL) InitAbsoluteURLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func (u NSURL) InitWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}
// Returns whether the resource pointed to by a file URL can be reached.
//
// error: The error that occurred when the resource could not be reached.
//
// # Return Value
// 
// [true] if the resource is reachable; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method synchronously checks if the file at the provided URL is
// reachable. Checking reachability is appropriate when making decisions that
// do not require other immediate operations on the resource, such as periodic
// maintenance of user interface state that depends on the existence of a
// specific document. For example, you might remove an item from a download
// list if the user deletes the file.
// 
// If your app must perform operations on the file, such as opening it or
// copying resource properties, it is more efficient to attempt the operation
// and handle any failure that may occur.
// 
// If this method returns [false], the object pointer referenced by `error` is
// populated with additional information.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)
func (u NSURL) CheckResourceIsReachableAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("checkResourceIsReachableAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkResourceIsReachableAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Returns whether the URL is a file reference URL.
//
// # Return Value
// 
// [true] if the URL is a file reference URL; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/isFileReferenceURL()
func (u NSURL) IsFileReferenceURL() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isFileReferenceURL"))
	return rv
}
// Returns the resource values for the properties identified by specified
// array of keys.
//
// keys: An array of property keys for the desired resource properties.
//
// # Return Value
// 
// A dictionary of resource values indexed by key.
//
// # Discussion
// 
// This method first checks if the URL object already caches the specified
// resource values. If so, it returns the cached resource values to the
// caller. If not, then this method synchronously obtains the resource values
// from the backing store, adds the resource values to the URL object’s
// cache, and returns the resource values to the caller.
// 
// The type of the returned resource value varies by resource property; for
// details, see the documentation for the key you want to access.
// 
// If the result dictionary does not contain a resource value for one or more
// of the requested resource keys, it means those resource properties are not
// available for the URL, and no errors occurred when determining those
// resource properties were not available.
// 
// If an error occurs, this method returns `nil` and populates the object
// pointer referenced by `error` with additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u NSURL) ResourceValuesForKeysError(keys []string) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("resourceValuesForKeys:error:"), objectivec.StringSliceToNSArray(keys), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}
// Returns the value of the resource property for the specified key.
//
// value: The location where the value for the resource property identified by `key`
// should be stored.
//
// key: The name of one of the URL’s resource properties.
//
// # Discussion
// 
// This method first checks if the URL object already caches the resource
// value. If so, it returns the cached resource value to the caller. If not,
// then this method synchronously obtains the resource value from the backing
// store, adds the resource value to the URL object’s cache, and returns the
// resource value to the caller.
// 
// The type of the returned resource value varies by resource property; for
// details, see the documentation for the key you want to access.
// 
// If this method returns [true] and the value is populated with `nil`, it
// means that the resource property is not available for the specified
// resource, and that no errors occurred when determining that the resource
// property was unavailable.
// 
// If this method returns [false], an error occurred. The object pointer
// referenced by `error` is populated with additional information.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u NSURL) GetResourceValueForKeyError(value []objectivec.IObject, key NSURLResourceKey) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("getResourceValue:forKey:error:"), objectivec.IObjectSliceToNSArray(value), objc.String(string(key)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getResourceValue:forKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the URL’s resource property for a given key to a given value.
//
// value: The value for the resource property defined by `key`.
//
// key: The name of one of the URL’s resource properties.
//
// # Discussion
// 
// This method synchronously writes the new resource value out to disk.
// Attempts to set a read-only resource property or to set a resource property
// that is not supported by the resource are ignored and are not considered
// errors.
// 
// If an error occurs, this method returns [false] and populates the object
// pointer referenced by `error` with additional information.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u NSURL) SetResourceValueForKeyError(value objectivec.IObject, key NSURLResourceKey) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("setResourceValue:forKey:error:"), value, objc.String(string(key)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setResourceValue:forKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the URL’s resource properties for a given set of keys to a given set
// of values.
//
// keyedValues: A dictionary of resource values to be set.
//
// # Discussion
// 
// This method synchronously writes the new resource value out to disk. If an
// error occurs after some resource properties have been successfully changed,
// the `userInfo` dictionary in the returned error object contains a
// `kCFURLKeysOfUnsetValuesKey` key whose value is an array of the resource
// values that were not successfully set.
// 
// Attempts to set a read-only resource property or to set a resource property
// that is not supported by the resource are ignored and are not considered
// errors.
// 
// The order in which the resource values are set is not defined. If you need
// to guarantee the order in which resource values are set, you should make
// multiple requests to this method or [SetResourceValueForKeyError].
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)
func (u NSURL) SetResourceValuesError(keyedValues INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("setResourceValues:error:"), keyedValues, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setResourceValues:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Removes all cached resource values and temporary resource values from the
// URL object.
//
// # Discussion
// 
// This method is applicable only to URLs that represent file system
// resources.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/removeAllCachedResourceValues()
func (u NSURL) RemoveAllCachedResourceValues() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllCachedResourceValues"))
}
// Removes the cached resource value identified by a given key from the URL
// object.
//
// key: The resource value key whose cached values you want to remove.
//
// # Discussion
// 
// Removing a cached resource value may remove other cached resource values
// because some resource values are cached as a set of values, and because
// some resource values depend on other resource values. (Temporary resource
// values have no dependencies.)
// 
// This method is currently applicable only to URLs for file system resources.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/removeCachedResourceValue(forKey:)
func (u NSURL) RemoveCachedResourceValueForKey(key NSURLResourceKey) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCachedResourceValueForKey:"), objc.String(string(key)))
}
// Sets a temporary resource value on the URL object.
//
// value: The value to store.
//
// key: The key where the value should be stored. This key must be unique and must
// not conflict with any system-defined keys. Reverse-domain-name notation is
// recommended.
//
// # Discussion
// 
// Your app can use a temporary resource value to temporarily store a value
// for an app-defined resource value key in memory without modifying the
// actual resource that the URL represents. Once set, you can copy the
// temporary resource value from the URL object just as you would copy
// system-defined keys—by calling [GetResourceValueForKeyError] or
// [ResourceValuesForKeysError].
// 
// Your app can remove a temporary resource value from the URL object by
// calling [RemoveCachedResourceValueForKey] or
// [RemoveAllCachedResourceValues] (to remove all temporary values).
// 
// This method is applicable only to URLs for file system resources.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u NSURL) SetTemporaryResourceValueForKey(value objectivec.IObject, key NSURLResourceKey) {
	objc.Send[objc.ID](u.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, objc.String(string(key)))
}
// Returns a new file reference URL that points to the same resource as the
// receiver.
//
// # Return Value
// 
// The new file reference URL.
//
// # Discussion
// 
// File reference URLs use a URL path syntax that identifies a file system
// object by reference, not by path. This form of file URL remains valid when
// the file system path of the URL’s underlying resource changes.
// 
// If the original URL is a file path URL, this property contains a copy of
// the URL converted into a file reference URL. If the original URL is a file
// reference URL, this property contains the original. If the original URL is
// not a file URL, this property contains `nil`.
// 
// File reference URLs cannot be created to file system objects which do not
// exist or are not reachable. This property contains `nil` instead.
// 
// In some areas of the file system hierarchy, file reference URLs cannot be
// generated to the leaf node of the URL path.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileReferenceURL()
func (u NSURL) FileReferenceURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("fileReferenceURL"))
	return NSURLFromID(rv)
}
// Returns a new URL by appending a path component to the original URL.
//
// pathComponent: The path component to add to the URL, in its original form (not URL
// encoded).
//
// # Return Value
// 
// A new URL with `pathComponent` appended.
//
// # Discussion
// 
// If the original URL does not end with a forward slash and `pathComponent`
// does not begin with a forward slash, a forward slash is inserted between
// the two parts of the returned URL, unless the original URL is the empty
// string.
// 
// If the receiver is a file URL and `pathComponent` does not end with a
// trailing slash, this method may read file metadata to determine whether the
// resulting path is a directory. This is done synchronously, and may have
// significant performance costs if the receiver is a location on a network
// mounted filesystem. You can instead call the
// [URLByAppendingPathComponentIsDirectory] method if you know whether the
// resulting path is a directory to avoid this file metadata operation.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u NSURL) URLByAppendingPathComponent(pathComponent string) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByAppendingPathComponent:"), objc.String(pathComponent))
	return NSURLFromID(rv)
}
// Returns a new URL by appending a path component to the original URL, along
// with a trailing slash if the component is a directory.
//
// pathComponent: The path component to add to the URL.
//
// isDirectory: If [true], a trailing slash is appended after `pathComponent`.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new URL with `pathComponent` appended.
//
// # Discussion
// 
// If the original URL does not end with a forward slash and `pathComponent`
// does not begin with a forward slash, a forward slash is inserted between
// the two parts of the returned URL, unless the original URL is the empty
// string.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u NSURL) URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), objc.String(pathComponent), isDirectory)
	return NSURLFromID(rv)
}
// Returns a URL by appending the specified path component with the file
// extension for a uniform type identifier.
//
// partialName: The partial name to append.
//
// contentType: A uniform type identifier the resulting conforms to.
//
// # Return Value
// 
// A new URL with the partial name and the type’s preferred extension
// appended.
//
// # Discussion
// 
// Use this method when you want to mix partial input from a user or other
// source, and need to produce a complete filename suitable for that input.
// For example, if you download a file from the internet and know its MIME
// type, you can use this method to ensure the URL has the correct filename
// extension where you save the file.
// 
// If `partialName` already has a path extension, and that path extension is
// valid for file system objects of type `contentType`, the function doesn’t
// add an extension before appending it to the URL. For example, if the inputs
// are `puppy.Jpg()` and [image], respectively, the function returns a URL
// with an appended path component of `puppy.Jpg()`. However, if the inputs
// are `puppy.Jpg()` and [plainText], respectively, the function returns a URL
// with an appended path component of `puppy.JpgXCUIElementTypeTxt()`. If you
// want to replace any existing path extension, use the
// [deletePathExtension()] method first.
// 
// If the function can’t append the path component, it returns an unchanged
// URL.
// 
// For more information about types, see [Uniform Type Identifiers].
//
// [Uniform Type Identifiers]: https://developer.apple.com/documentation/UniformTypeIdentifiers
// [deletePathExtension()]: https://developer.apple.com/documentation/Foundation/URL/deletePathExtension()
// [image]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/image
// [plainText]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/plainText
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)
func (u NSURL) URLByAppendingPathComponentConformingToType(partialName string, contentType objectivec.IObject) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return NSURLFromID(rv)
}
// Returns a new URL by appending a path extension to the original URL.
//
// pathExtension: The path extension to add to the URL.
//
// # Return Value
// 
// A new URL with `pathExtension` appended.
//
// # Discussion
// 
// If the original URL ends with one or more forward slashes, these are
// removed from the returned URL. A period is inserted between the two parts
// of the new URL.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(_:)
func (u NSURL) URLByAppendingPathExtension(pathExtension string) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByAppendingPathExtension:"), objc.String(pathExtension))
	return NSURLFromID(rv)
}
// Returns a URL by appending the path extension for a uniform type
// identifier.
//
// contentType: The uniform type identifer to use for the extension.
//
// # Return Value
// 
// A new URL with the type’s preferred extension appended.
//
// # Discussion
// 
// For more information about types, see [Uniform Type Identifiers].
//
// [Uniform Type Identifiers]: https://developer.apple.com/documentation/UniformTypeIdentifiers
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(for:)
func (u NSURL) URLByAppendingPathExtensionForType(contentType objectivec.IObject) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByAppendingPathExtensionForType:"), contentType)
	return NSURLFromID(rv)
}
// Returns a bookmark for the URL, created with specified options and resource
// values.
//
// options: Options taken into account when creating the bookmark for the URL. The
// possible flags (which can be combined with bitwise [OR] operations) are
// described in [NSURL.BookmarkCreationOptions].
// 
// To create a security-scoped bookmark to support App Sandbox, include the
// [URLBookmarkCreationWithSecurityScope] flag. When you later resolve the
// bookmark, you can use the resulting security-scoped URL to obtain
// read/write access to the file-system resource pointed to by the URL.
// 
// If you instead want to create a security-scoped bookmark that, when
// resolved, enables you to obtain read-only access to a file-system resource,
// bitwise [OR] this parameter’s value with both the
// [URLBookmarkCreationWithSecurityScope] option and the
// [URLBookmarkCreationSecurityScopeAllowOnlyReadAccess] option.
// //
// [NSURL.BookmarkCreationOptions]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions
//
// keys: An array of names of URL resource properties to store as part of the
// bookmark. You can later access these values (without resolving the
// bookmark) by calling the [ResourceValuesForKeysFromBookmarkData] method.
// 
// The values of these properties must be of a type that the bookmark
// generation code can serialize. Specifically, the values can contain any of
// the following primitive types:
// 
// - [NSString] or [CFString] - [NSData] or [CFData] - [NSDate] or [CFDate] -
// [NSNumber] or [CFNumber] - [CFBoolean] - [NSURL] or [CFURL] - `kCFNull` or
// [NSNull] - [CFUUID]
// 
// In addition, the properties can contain the following collection classes:
// 
// - [NSArray] or [CFArray] containing only the above primitive types -
// [NSDictionary] or [CFDictionary] with [NSString] or [CFString] keys, in
// which all values contain only the above primitive types
//
// relativeURL: The URL that the bookmark data will be relative to.
// 
// If you are creating a security-scoped bookmark to support App Sandbox, use
// this parameter as follows:
// 
// - To create an app-scoped bookmark, use a value of `nil`. - To create a
// document-scoped bookmark, use the path (despite this parameter’s name) to
// the document file that is to own the new security-scoped bookmark.
// 
// App Sandbox does not restrict which URL values may be passed to this
// parameter.
//
// # Return Value
// 
// A bookmark for the URL.
//
// # Discussion
// 
// This method returns bookmark data that can later be resolved into a URL
// object for a file even if the user moves or renames it (if the volume
// format on which the file resides supports doing so).
// 
// You can also use this method to create a security-scoped bookmark to
// support App Sandbox. Before you do so, you must first enable the
// appropriate entitlements for your app, as described in [Enabling
// Security-Scoped Bookmark and URL Access]. In addition, be sure to
// understand the behavior of the `options` and `relativeURL` parameters.
// 
// For an app-scoped bookmark, no sandboxed app other than the one that
// created the bookmark can obtain access to the file-system resource that the
// URL (obtained from the bookmark) points to. Specifically, a bookmark
// created with security scope fails to resolve if the caller does not have
// the same code signing identity as the caller that created the bookmark.
// 
// For a document-scoped bookmark, any sandboxed app that has access to the
// bookmark data itself, and has access to the document that owns the
// bookmark, can obtain access to the resource.
//
// [Enabling Security-Scoped Bookmark and URL Access]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/EntitlementKeyReference/Chapters/EnablingAppSandbox.html#//apple_ref/doc/uid/TP40011195-CH4-SW18
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u NSURL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options NSURLBookmarkCreationOptions, keys []string, relativeURL INSURL) (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, objectivec.StringSliceToNSArray(keys), relativeURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// In an app that has adopted App Sandbox, makes the resource pointed to by a
// security-scoped URL available to the app.
//
// # Return Value
// 
// [true] if the request to access the resource succeeded; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When you obtain a security-scoped URL, such as by resolving a
// security-scoped bookmark, you can’t immediately use the resource it
// points to. To make the resource available to your app, by way of adding its
// location to your app’s sandbox, call this method on the security-scoped
// URL. You can also use Core Foundation equivalent, the
// [CFURLStartAccessingSecurityScopedResource(_:)] function.
// 
// If this method returns [true], then you must relinquish access as soon as
// you finish using the resource. Call the
// [StopAccessingSecurityScopedResource] method to relinquish access. You must
// balance each call to [StartAccessingSecurityScopedResource] for a given
// security-scoped URL with a call to [StopAccessingSecurityScopedResource].
// When you make the last balanced call to
// [StopAccessingSecurityScopedResource], you immediately lose access to the
// resource in question.
//
// [CFURLStartAccessingSecurityScopedResource(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
func (u NSURL) StartAccessingSecurityScopedResource() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("startAccessingSecurityScopedResource"))
	return rv
}
// In an app that adopts App Sandbox, revokes access to the resource pointed
// to by a security-scoped URL.
//
// # Discussion
// 
// When you no longer need access to a file or directory pointed to by a
// security-scoped URL, such as one returned by resolving a security-scoped
// bookmark, call this method on the URL to relinquish access. You can also
// use its Core Foundation equivalent, the
// [CFURLStopAccessingSecurityScopedResource(_:)] function.
// 
// You must balance each call to [StartAccessingSecurityScopedResource] for a
// given security-scoped URL with a call to
// [StopAccessingSecurityScopedResource]. When you make the last balanced call
// to [StopAccessingSecurityScopedResource], you immediately lose access to
// the resource in question.
// 
// If you call this method on a URL whose referenced resource you don’t have
// access to, nothing happens.
//
// [CFURLStopAccessingSecurityScopedResource(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
func (u NSURL) StopAccessingSecurityScopedResource() {
	objc.Send[objc.ID](u.ID, objc.Sel("stopAccessingSecurityScopedResource"))
}
// Returns whether the promised item can be reached.
//
// error: The error that occurred when the promised item could not be reached.
//
// # Return Value
// 
// [true] if the promised item is reachable; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method behaves identically to
// [CheckResourceIsReachableAndReturnError], but works on promised items. A
// promised item is not guaranteed to have its contents in the file system
// until you use a file coordinator to perform a coordinated read on its URL,
// which causes the contents to be downloaded or otherwise generated. Promised
// item URLs are returned by various APIs, including:
// 
// - A metadata query using either the [NSMetadataQueryUbiquitousDataScope] or
// [NSMetadataQueryUbiquitousDocumentsScope] scopes - The contents of the
// directory returned by the file manager’s
// `URLForUbiquitousContainerIdentifier:` - The URL inside the accessor block
// of a coordinated read or write operation that used the
// [FileCoordinatorReadingImmediatelyAvailableMetadataOnly],
// [FileCoordinatorWritingForDeleting], [FileCoordinatorWritingForMoving], or
// [FileCoordinatorWritingContentIndependentMetadataOnly] options
// 
// You must use this method instead of
// `checkResourceIsReachableAndReturnError` for any URLs returned by these
// methods.
//
// [NSMetadataQueryUbiquitousDataScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDataScope
// [NSMetadataQueryUbiquitousDocumentsScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDocumentsScope
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)
func (u NSURL) CheckPromisedItemIsReachableAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkPromisedItemIsReachableAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Returns the value of the resource property for the specified key.
//
// value: The location where the value for the resource property identified by `key`
// should be stored.
//
// key: The name of one of the URL’s resource properties.
//
// # Discussion
// 
// This method behaves identically to [GetResourceValueForKeyError], but works
// on promised items. A promised item is not guaranteed to have its contents
// in the file system until you use a file coordinator to perform a
// coordinated read on its URL, which causes the contents to be downloaded or
// otherwise generated. Promised item URLs are returned by various APIs,
// including:
// 
// - A metadata query using either the [NSMetadataQueryUbiquitousDataScope] or
// [NSMetadataQueryUbiquitousDocumentsScope] scopes - The contents of the
// directory returned by the file manager’s
// `URLForUbiquitousContainerIdentifier:` - The URL inside the accessor block
// of a coordinated read or write operation that used the
// [FileCoordinatorReadingImmediatelyAvailableMetadataOnly],
// [FileCoordinatorWritingForDeleting], [FileCoordinatorWritingForMoving], or
// [FileCoordinatorWritingContentIndependentMetadataOnly] options
// 
// You must use this method instead of `` for any URLs returned by these
// methods.
// 
// This method works for any resource value that is not tied to the item’s
// contents. Some keys, like [contentAccessDateKey] or
// [generationIdentifierKey], do not return valid values. If you use one of
// these keys, the method returns [true], but the value returns `nil`.
//
// [NSMetadataQueryUbiquitousDataScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDataScope
// [NSMetadataQueryUbiquitousDocumentsScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDocumentsScope
// [contentAccessDateKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/contentAccessDateKey
// [generationIdentifierKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/generationIdentifierKey
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)
func (u NSURL) GetPromisedItemResourceValueForKeyError(value []objectivec.IObject, key NSURLResourceKey) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("getPromisedItemResourceValue:forKey:error:"), objectivec.IObjectSliceToNSArray(value), objc.String(string(key)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getPromisedItemResourceValue:forKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Returns the resource values for the properties identified by specified
// array of keys.
//
// keys: An array of names of URL resource properties.
//
// # Return Value
// 
// A dictionary of resource values indexed by key.
//
// # Discussion
// 
// This method behaves identically to [ResourceValuesForKeysError], but works
// on promised items. A promised item is not guaranteed to have its contents
// in the file system until you use a file coordinator to perform a
// coordinated read on its URL, which causes the contents to be downloaded or
// otherwise generated. Promised item URLs are returned by various APIs,
// including:
// 
// - A metadata query using either the [NSMetadataQueryUbiquitousDataScope] or
// [NSMetadataQueryUbiquitousDocumentsScope] scopes - The contents of the
// directory returned by the file manager’s
// `URLForUbiquitousContainerIdentifier:` - The URL inside the accessor block
// of a coordinated read or write operation that used the
// [FileCoordinatorReadingImmediatelyAvailableMetadataOnly],
// [FileCoordinatorWritingForDeleting], [FileCoordinatorWritingForMoving], or
// [FileCoordinatorWritingContentIndependentMetadataOnly] options
// 
// You must use this method instead of `` for any URLs returned by these
// methods.
// 
// This method works for any resource value that is not tied to the item’s
// contents. Some keys, like [contentAccessDateKey] or
// [generationIdentifierKey], do not return valid values. If you use one of
// these keys, the method returns [true], but the value returns `nil`.
//
// [NSMetadataQueryUbiquitousDataScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDataScope
// [NSMetadataQueryUbiquitousDocumentsScope]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDocumentsScope
// [contentAccessDateKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/contentAccessDateKey
// [generationIdentifierKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/generationIdentifierKey
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)
func (u NSURL) PromisedItemResourceValuesForKeysError(keys []string) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("promisedItemResourceValuesForKeys:error:"), objectivec.StringSliceToNSArray(keys), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}
// Writes the URL to the specified pasteboard.
//
// pasteBoard: The target pasteboard.
//
// pasteBoard is a [appkit.NSPasteboard].
//
// # Discussion
// 
// You must declare an [NSURLPboardType] data type for the pasteboard before
// invoking this method. Otherwise, the method returns without doing anything.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/write(to:)
func (u NSURL) WriteToPasteboard(pasteBoard objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u NSURL) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u NSURL) InitWithCoder(coder INSCoder) NSURL {
	rv := objc.Send[NSURL](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Asks the item provider for the representation visibility specification for
// the given UTI.
//
// typeIdentifier: A uniform type identifier (UTI).
//
// # Return Value
// 
// A representation visibility specification for the given UTI.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/itemProviderVisibilityForRepresentation(withTypeIdentifier:)-swift.method
func (u NSURL) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) NSItemProviderRepresentationVisibility {
	rv := objc.Send[NSItemProviderRepresentationVisibility](u.ID, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), objc.String(typeIdentifier))
	return NSItemProviderRepresentationVisibility(rv)
}
// Loads data of a particular type, identified by the given UTI.
//
// typeIdentifier: The uniform type identifier (UTI) identifying the type of data to load.
//
// completionHandler: The handler that’s called after the data is loaded.
//
// # Return Value
// 
// The progress of the data load process.
//
// # Discussion
// 
// When the system calls this method, the `typeIdentifier` parameter is set to
// one of the elements in the `writableTypeIdentifiersForItemProvider` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/loadData(withTypeIdentifier:forItemProviderCompletionHandler:)
func (u NSURL) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
_block1, _cleanup1 := NewDataErrorBlock(completionHandler)
	defer _cleanup1()
	rv := objc.Send[objc.ID](u.ID, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Initializes and returns a newly created NSURL object as a file URL with a
// specified path.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// isDir: A Boolean value that specifies whether `path` is treated as a directory
// path when resolving against relative path components. Pass [true] if the
// `path` indicates a directory, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (_NSURLClass NSURLClass) FileURLWithPathIsDirectory(path string, isDir bool) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return NSURLFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (_NSURLClass NSURLClass) FileURLWithPathRelativeToURL(path string, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return NSURLFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (_NSURLClass NSURLClass) FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return NSURLFromID(rv)
}
// Initializes and returns a newly created NSURL object as a file URL with a
// specified path.
//
// path: The path that the NSURL object will represent. `path` should be a valid
// system path, and must not be an empty path. If `path` begins with a tilde,
// it must first be expanded with [StringByExpandingTildeInPath]. If `path` is
// a relative path, it is treated as being relative to the current working
// directory.
//
// # Return Value
// 
// An NSURL object initialized with `path`.
//
// # Discussion
// 
// This method assumes that `path` is a directory if it ends with a slash. If
// `path` does not end with a slash, the method examines the file system to
// determine if `path` is a file or a directory. If `path` exists in the file
// system and is a directory, the method appends a trailing slash. If `path`
// does not exist in the file system, the method assumes that it represents a
// file and does not append a trailing slash.
// 
// As an alternative, consider using [FileURLWithPathIsDirectory], which
// allows you to explicitly specify whether the returned [NSURL] object
// represents a file or directory.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)
func (_NSURLClass NSURLClass) FileURLWithPath(path string) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithPath:"), objc.String(path))
	return NSURLFromID(rv)
}
// Initializes and returns a newly created NSURL object as a file URL with
// specified path components.
//
// components: An array of path components.
//
// # Return Value
// 
// An NSURL object initialized with `components`.
//
// # Discussion
// 
// The path components are separated by a forward slash in the returned URL.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPathComponents:)
func (_NSURLClass NSURLClass) FileURLWithPathComponents(components []string) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithPathComponents:"), objectivec.StringSliceToNSArray(components))
	return NSURLFromID(rv)
}
// Returns a new URL object initialized with a C string representing a local
// file system path.
//
// path: A null-terminated C string in file system representation containing the
// path to represent as a URL. If this path is a relative path, it is treated
// as being relative to the current working directory.
//
// isDir: [true] if the last path part is a directory, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// baseURL: The base URL for the new URL object. This must be a file URL. If `path` is
// absolute, this URL is ignored.
//
// # Return Value
// 
// Returns the new object.
//
// # Discussion
// 
// The file system representation format is described in [File System
// Programming Guide].
//
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)
func (_NSURLClass NSURLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path string, isDir bool, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), unsafe.Pointer(unsafe.StringData(path + "\x00")), isDir, baseURL)
	return NSURLFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)
func (_NSURLClass NSURLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return NSURLFromID(rv)
}
// Initializes and returns bookmark data derived from an alias file pointed to
// by a specified URL.
//
// bookmarkFileURL: The URL that points to a file containing bookmark data.
//
// # Return Value
// 
// The bookmark data for the alias file.
//
// # Discussion
// 
// This method doesn’t check to see if `bookmarkFileURL` points to an alias
// file. This allows this method to work with any file containing bookmark
// data. If `bookmarkFileURL` refers to a file which does not contain bookmark
// data or to a non-file object, such as a directory or symbolic link, this
// method returns `nil` produces an error.
// 
// This method returns `nil` if bookmark data cannot be created.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)
func (_NSURLClass NSURLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL INSURL) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("bookmarkDataWithContentsOfURL:error:"), bookmarkFileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Returns the resource values for properties identified by a specified array
// of keys contained in specified bookmark data.
//
// keys: An array of names of URL resource properties. In addition to the standard,
// system-defined resource properties, you can also request any custom
// properties that you provided when you created the bookmark. See the
// [BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError]
// method for details.
//
// bookmarkData: The bookmark data from which you want to retrieve resource values.
//
// # Return Value
// 
// A dictionary of the requested resource values contained in `bookmarkData`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)
func (_NSURLClass NSURLClass) ResourceValuesForKeysFromBookmarkData(keys []string, bookmarkData INSData) INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("resourceValuesForKeys:fromBookmarkData:"), objectivec.StringSliceToNSArray(keys), bookmarkData)
	return NSDictionaryFromID(rv)
}
// Creates an alias file on disk at a specified location with specified
// bookmark data.
//
// bookmarkData: The bookmark data containing information for the alias file.
//
// bookmarkFileURL: The desired location of the alias file.
//
// options: Options taken into account when creating the alias file.
//
// # Discussion
// 
// This method will produce an error if `bookmarkData` was not created with
// the [NSURLBookmarkCreationSuitableForBookmarkFile] option.
// 
// If `bookmarkFileURL` points to a directory, the alias file will be created
// in that directory with its name derived from the information in
// `bookmarkData`. If `bookmarkFileURL` points to a file, the alias file will
// be created with the location and name indicated by `bookmarkFileURL`, and
// its extension will be changed to `XCUIElementTypeAlias` if it is not
// already.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)
func (_NSURLClass NSURLClass) WriteBookmarkDataToURLOptionsError(bookmarkData INSData, bookmarkFileURL INSURL, options NSURLBookmarkFileCreationOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSURLClass.class), objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, bookmarkFileURL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeBookmarkData:toURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Creates a new instance of a class using the given data and UTI string.
//
// data: The data used to create the object.
//
// typeIdentifier: The uniform type identifier (UTI) representing the data type of `data`.
//
// # Return Value
// 
// An object created from the given data.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderReading/object(withItemProviderData:typeIdentifier:)
func (_NSURLClass NSURLClass) ObjectWithItemProviderDataTypeIdentifierError(data INSData, typeIdentifier string) (NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("objectWithItemProviderData:typeIdentifier:error:"), data, objc.String(typeIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}
// Returns a new URL made by resolving bookmark data.
//
// bookmarkData: The bookmark data the URL is derived from.
//
// options: Options taken into account when resolving the bookmark data.
// 
// To resolve a security-scoped bookmark to support App Sandbox, you must
// include (by way of bitwise [OR] operators with any other options in this
// parameter) the [URLBookmarkResolutionWithSecurityScope] option.
//
// relativeURL: The base URL that the bookmark data is relative to.
// 
// If you are resolving a security-scoped bookmark to obtain a security-scoped
// URL, use this parameter as follows:
// 
// - To resolve an app-scoped bookmark, use a value of `nil`. - To resolve a
// document-scoped bookmark, use the path (despite this parameter’s name) to
// the document from which you retrieved the bookmark.
// 
// App Sandbox does not restrict which URL values may be passed to this
// parameter.
//
// isStale: On return, if [true], the bookmark data is stale. Your app should create a
// new bookmark using the returned URL and use it in place of any stored
// copies of the existing bookmark.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// error: The error that occurred in the case that the URL cannot be created.
//
// # Return Value
// 
// A new URL made by resolving `bookmarkData`.
//
// # Discussion
// 
// This method fails if the original file or directory could not be located or
// is on a volume that could not be mounted. If this method fails, you can use
// the [ResourceValuesForKeysFromBookmarkData] method to obtain information
// about the bookmark, such as the last known path ([pathKey]) to help the
// user decide how to proceed.
// 
// To obtain a security-scoped URL from a security-scoped bookmark, call this
// method using the [URLBookmarkResolutionWithSecurityScope] option. In
// addition, to use security scope, you must first have enabled the
// appropriate entitlements for your app, as described in [Enabling
// Security-Scoped Bookmark and URL Access].
// 
// To then obtain access to the file-system resource pointed to by a
// security-scoped URL (in other words, to bring the resource into your
// app’s sandbox), call the [StartAccessingSecurityScopedResource] method
// (or its Core Foundation equivalent, the
// [CFURLStartAccessingSecurityScopedResource(_:)] function) on the URL.
// 
// For an app-scoped bookmark, no sandboxed app other than the one that
// created the bookmark can obtain access to the file-system resource that the
// URL (obtained from the bookmark) points to.
// 
// For a document-scoped bookmark, any sandboxed app that has access to the
// bookmark data itself, and has access to the document that owns the
// bookmark, can obtain access to the resource.
//
// [CFURLStartAccessingSecurityScopedResource(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
// [Enabling Security-Scoped Bookmark and URL Access]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/EntitlementKeyReference/Chapters/EnablingAppSandbox.html#//apple_ref/doc/uid/TP40011195-CH4-SW18
// [pathKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/pathKey
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (_NSURLClass NSURLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData INSData, options NSURLBookmarkResolutionOptions, relativeURL INSURL, isStale unsafe.Pointer) (NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:
func (_NSURLClass NSURLClass) URLWithDataRepresentationRelativeToURL(data INSData, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return NSURLFromID(rv)
}
// Creates and returns an NSURL object initialized with a provided URL string.
//
// URLString: The URL string with which to initialize the NSURL object. Linked on or
// after iOS 17, this method parses [URLString] according to RFC 3986. Linked
// before iOS 17, this method parses [URLString] according to RFCs 1738 and
// 1808.
//
// # Return Value
// 
// An NSURL object initialized with [URLString]. If the URL string was
// malformed or `nil`, returns `nil`.
//
// # Discussion
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:
func (_NSURLClass NSURLClass) URLWithString(URLString string) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("URLWithString:"), objc.String(URLString))
	return NSURLFromID(rv)
}
// Creates and returns an instance from the provided string, optionally IDNA-
// and percent-encoding any invalid characters.
//
// URLString: A URL location.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in `string`.
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the method returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:
func (_NSURLClass NSURLClass) URLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("URLWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return NSURLFromID(rv)
}
// Creates and returns an NSURL object initialized with a base URL and a
// relative string.
//
// URLString: The URL string with which to initialize the NSURL object. May not be `nil`.
// Linked on or after iOS 17, this method parses [URLString] according to RFC
// 3986. Linked before iOS 17, this method parses [URLString] according to
// RFCs 1738 and 1808. [URLString] is interpreted relative to `baseURL`.
//
// baseURL: The base URL for the NSURL object.
//
// # Return Value
// 
// An NSURL object initialized with [URLString] and `baseURL`. If [URLString]
// was malformed or `nil`, returns `nil`.
//
// # Discussion
// 
// This method allows you to create a URL relative to a base path or URL. For
// example, if you have the URL for a folder on disk and the name of a file
// within that folder, you can construct a URL for the file by providing the
// folder’s URL as the base path (with a trailing slash) and the filename as
// the string part.
// 
// To check if [URLString] is strictly valid according to the RFC, use the new
// `[NSURL URLWithString:URLString NO]` method. This method leaves all
// characters as they are and returns `nil` if [URLString] is explicitly
// invalid.
// 
// For apps linked before iOS 17, this method expects [URLString] to contain
// only characters that are allowed in a properly formed URL. All other
// characters must be properly percent encoded. Any percent-encoded characters
// are interpreted using UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:relativeToURL:
func (_NSURLClass NSURLClass) URLWithStringRelativeToURL(URLString string, baseURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSURLClass.class), objc.Sel("URLWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return NSURLFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSURL/dataRepresentation
func (u NSURL) DataRepresentation() INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataRepresentation"))
	return NSDataFromID(objc.ID(rv))
}
// A boolean value that determines whether the receiver is a file URL.
//
// # Discussion
// 
// The property’s value is [true] if the receiver uses the file scheme,
// [false] otherwise. Both file path and file reference URLs are considered to
// be file URLs.
// 
// If this property’s value is [true], then the receiver’s [Path] property
// contains a suitable value for input into [NSFileManager] or
// [NSPathUtilities].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
func (u NSURL) FileURL() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isFileURL"))
	return rv
}
// The URL string for the receiver as an absolute URL. (read-only)
//
// # Discussion
// 
// This property’s value is calculated by resolving the receiver’s string
// against its base according to the algorithm given in RFC 1808.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/absoluteString
func (u NSURL) AbsoluteString() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("absoluteString"))
	return NSStringFromID(rv).String()
}
// An absolute URL that refers to the same resource as the receiver.
// (read-only)
//
// # Discussion
// 
// If the URL is already absolute, this property contains a copy of the
// receiver. Resolution is performed per RFC 1808.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL
func (u NSURL) AbsoluteURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("absoluteURL"))
	return NSURLFromID(objc.ID(rv))
}
// The base URL. (read-only)
//
// # Discussion
// 
// This property contains the base URL. If the receiver is an absolute URL,
// this property contains `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/baseURL
func (u NSURL) BaseURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("baseURL"))
	return NSURLFromID(objc.ID(rv))
}
// A C string containing the URL’s file system path. (read-only)
//
// # Discussion
// 
// This returns a null-terminated C string in file system representation.
// 
// This string is automatically freed in the same way that a returned object
// would be released. The caller must either copy the string or use
// [GetFileSystemRepresentationMaxLength] if it needs to store the
// representation outside of the autorelease context in which the value was
// obtained.
// 
// The file system representation format is described in File Encodings and
// Fonts.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fileSystemRepresentation
func (u NSURL) FileSystemRepresentation() string {
	rv := objc.Send[*byte](u.ID, objc.Sel("fileSystemRepresentation"))
	return objc.GoString(rv)
}
// The fragment identifier, conforming to RFC 1808. (read-only)
//
// # Discussion
// 
// This property contains the URL’s fragment. Any percent-encoded characters
// are not unescaped. If the receiver does not conform to RFC 1808, this
// property contains `nil`. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()#jumpLocation`, the fragment
// identifier is `jumpLocation`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/fragment
func (u NSURL) Fragment() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("fragment"))
	return NSStringFromID(rv).String()
}
// The host, conforming to RFC 1808. (read-only)
//
// # Discussion
// 
// This property contains the host, unescaped using the
// [replacingPercentEscapes(using:)] method. If the receiver does not conform
// to RFC 1808, this property contains `nil`. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()`, the host is
// `www.ExampleXCUIElementTypeCom()`.
//
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/host
func (u NSURL) Host() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("host"))
	return NSStringFromID(rv).String()
}
// The last path component. (read-only)
//
// # Discussion
// 
// This property contains the last path component, unescaped using the
// [replacingPercentEscapes(using:)] method. For example, in the URL
// `///path/to/file`, the last path component is `file`.
//
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/lastPathComponent
func (u NSURL) LastPathComponent() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("lastPathComponent"))
	return NSStringFromID(rv).String()
}
// The password conforming to RFC 1808. (read-only)
//
// # Discussion
// 
// This property contains the password. Any percent-encoded characters are not
// unescaped. If the receiver does not conform to RFC 1808, it contains `nil`.
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the password is
// `password`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/password
func (u NSURL) Password() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("password"))
	return NSStringFromID(rv).String()
}
// The path, conforming to RFC 1808. (read-only)
//
// # Discussion
// 
// This property contains the path, unescaped using the
// [replacingPercentEscapes(using:)] method. If the receiver does not conform
// to RFC 1808, this property contains `nil`.
// 
// If the receiver contains a file or file reference URL (as determined with
// [FileURL]), this property’s value is suitable for input into methods of
// [NSFileManager] or [NSPathUtilities]. If the path has a trailing slash, it
// is stripped.
// 
// If the receiver contains a file reference URL, this property’s value
// provides the for the referenced resource, which may be `nil` if the
// resource no longer exists.
// 
// If the [parameterString] property contains a non-`nil` value, the path may
// be incomplete. If the receiver contains an unencoded semicolon, the path
// property ends at the character before the semicolon. The remainder of the
// URL is provided in the [parameterString] property.
// 
// To obtain the complete path, if [parameterString] contains a non-`nil`
// value, append a semicolon, followed by the parameter string.
// 
// Per RFC 3986, the leading slash after the authority (host name and port)
// portion is treated as part of the path. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()`, the path is
// `/index.Html()`.
//
// [parameterString]: https://developer.apple.com/documentation/Foundation/NSURL/parameterString
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/path
func (u NSURL) Path() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("path"))
	return NSStringFromID(rv).String()
}
// An array containing the path components. (read-only)
//
// # Discussion
// 
// This property contains an array containing the individual path components
// of the URL, each unescaped using the [replacingPercentEscapes(using:)]
// method. For example, in the URL `///directory/directory%202/file`, the path
// components array would be `@[@"/", @"directory", @"directory 2", @"file"]`.
//
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/pathComponents
func (u NSURL) PathComponents() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("pathComponents"))
	return objc.ConvertSliceToStrings(rv)
}
// The path extension. (read-only)
//
// # Return Value
// 
// The path extension of the receiver, or `nil` if [Path] is `nil`.
// 
// # Discussion
// 
// This property contains the path extension, unescaped using the
// [replacingPercentEscapes(using:)] method. For example, in the URL
// `///path/to/file.Txt()`, the path extension is `txt`.
//
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/pathExtension
func (u NSURL) PathExtension() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("pathExtension"))
	return NSStringFromID(rv).String()
}
// The port, conforming to RFC 1808.
//
// # Discussion
// 
// This property contains the port number. If the receiver does not conform to
// RFC 1808, this property contains `nil`. For example, in the URL
// `//www.Example().8080/index.Php()`, the port number is `8080`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/port
func (u NSURL) Port() INSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("port"))
	return NSNumberFromID(objc.ID(rv))
}
// The query string, conforming to RFC 1808.
//
// # Discussion
// 
// This property contains the query string. Any percent-encoded characters are
// not unescaped. If the receiver does not conform to RFC 1808, this property
// contains `nil`. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1&key2=value2`,
// the query string is `key1=value1&key2=value2`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/query
func (u NSURL) Query() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("query"))
	return NSStringFromID(rv).String()
}
// The relative path, conforming to RFC 1808. (read-only)
//
// # Discussion
// 
// This property contains the relative path of the receiver’s URL without
// resolving against its base URL. Any percent-encoded characters are not
// unescaped. If the path has a trailing slash it is stripped. If the receiver
// is an absolute URL, this property contains the same value as [Path]. If the
// receiver does not conform to RFC 1808, it contains `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/relativePath
func (u NSURL) RelativePath() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("relativePath"))
	return NSStringFromID(rv).String()
}
// A string representation of the relative portion of the URL. (read-only)
//
// # Discussion
// 
// This property contains a string representation of the relative portion of
// the URL. Any percent-encoded characters are not unescaped. If the receiver
// is an absolute URL this method returns the same value as [AbsoluteString].
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/relativeString
func (u NSURL) RelativeString() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("relativeString"))
	return NSStringFromID(rv).String()
}
// The resource specifier. (read-only)
//
// # Discussion
// 
// This property contains the resource specifier. Any percent-encoded
// characters are not unescaped. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()?key1=value1#jumplink`, the
// resource specifier is
// `//www.ExampleXCUIElementTypeCom()/index.Html()?key1=value1#jumplink`
// (everything after the colon).
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/resourceSpecifier
func (u NSURL) ResourceSpecifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("resourceSpecifier"))
	return NSStringFromID(rv).String()
}
// The scheme. (read-only)
//
// # Discussion
// 
// This property contains the scheme. Any percent-encoded characters are not
// unescaped. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()`, the scheme is `http`.
// 
// The full URL is the concatenation of the scheme, a colon (`:`), and the
// value of [ResourceSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/scheme
func (u NSURL) Scheme() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("scheme"))
	return NSStringFromID(rv).String()
}
// A copy of the URL with any instances of `".."` or `"."` removed from its
// path. (read-only)
//
// # Discussion
// 
// This property contains a new [NSURL] object, initialized using the
// receiver’s path with any instances of `".."` or `"."` removed. If the
// receiver does not conform to RFC 1808, this property contains `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/standardized
func (u NSURL) StandardizedURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("standardizedURL"))
	return NSURLFromID(objc.ID(rv))
}
// The user name, conforming to RFC 1808.
//
// # Discussion
// 
// This property contains the user name, unescaped using the
// [replacingPercentEscapes(using:)] method. If the receiver’s URL does not
// conform to RFC 1808, this property returns `nil`. For example, in the URL
// `//username@www.ExampleXCUIElementTypeCom()/`, the user name is `username`.
//
// [replacingPercentEscapes(using:)]: https://developer.apple.com/documentation/Foundation/NSString/replacingPercentEscapes(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/user
func (u NSURL) User() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("user"))
	return NSStringFromID(rv).String()
}
// A file path URL that points to the same resource as the URL object.
// (read-only)
//
// # Discussion
// 
// If the receiver is a file reference URL, this property contains a copy of
// the URL converted to a file path URL. If the receiver’s URL is a file
// path URL, this property contains the original URL. If the original URL is
// not a file URL, or if the resource is not reachable or no longer exists,
// this property contains `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/filePathURL
func (u NSURL) FilePathURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("filePathURL"))
	return NSURLFromID(objc.ID(rv))
}
// A URL you create by removing the last path component from the receiver.
// (read-only)
//
// # Discussion
// 
// If the receiver’s URL represents the root path, this property contains a
// copy of the original URL. Otherwise, if the original URL has only one path
// component, this property contains the empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/deletingLastPathComponent
func (u NSURL) URLByDeletingLastPathComponent() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByDeletingLastPathComponent"))
	return NSURLFromID(objc.ID(rv))
}
// A URL you create by removing the path extension from the receiver, if any.
// (read-only)
//
// # Discussion
// 
// If the receiver represents the root path, this property contains a copy of
// the original URL. If the URL has multiple path extensions, only the last
// one is removed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/deletingPathExtension
func (u NSURL) URLByDeletingPathExtension() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByDeletingPathExtension"))
	return NSURLFromID(objc.ID(rv))
}
// A URL that points to the same resource as the receiver and includes no
// symbolic links. (read-only)
//
// # Discussion
// 
// If the receiver has no symbolic links, this property contains a copy of the
// original URL.
// 
// If some symbolic links cannot be resolved, this property contains those
// broken symbolic links.
// 
// If the name of the receiving path begins with `/private`, this property
// strips off the `/private` designator, provided the result is the name of an
// existing file.
// 
// This property only works on URLs with the `` path scheme. For all other
// URLs, it contains a copy of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/resolvingSymlinksInPath
func (u NSURL) URLByResolvingSymlinksInPath() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByResolvingSymlinksInPath"))
	return NSURLFromID(objc.ID(rv))
}
// A URL that points to the same resource as the original URL using an
// absolute path. (read-only)
//
// # Discussion
// 
// This property only works on URLs with the `` path scheme. For all other
// URLs, it returns a copy of the original URL.
// 
// Like [StringByStandardizingPath], this property can make the following
// changes in the provided URL:
// 
// - Expand an initial tilde expression using [StringByExpandingTildeInPath].
// - Reduce empty components and references to the current directory (that is,
// the sequences “//” and “/./”) to single path separators. - In
// absolute paths only, resolve references to the parent directory (that is,
// the component “..”) to the real parent directory if possible using
// [StringByResolvingSymlinksInPath], which consults the file system to
// resolve each potential symbolic link.
// 
// In relative paths, because symbolic links can’t be resolved, references
// to the parent directory are left in place.
// 
// - Remove an initial component of “/private” from the path if the result
// still indicates an existing file or directory (checked by consulting the
// file system).
// 
// Note that the path contained by this property may still have symbolic link
// components in it. Note also that this property only works with file paths
// (not, for example, string representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/standardizingPath
func (u NSURL) URLByStandardizingPath() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLByStandardizingPath"))
	return NSURLFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the URL string’s path represents a
// directory.
//
// # Discussion
// 
// This property doesn’t check the resource the URL refers to.
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/hasDirectoryPath
func (u NSURL) HasDirectoryPath() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("hasDirectoryPath"))
	return rv
}
// A custom playground Quick Look for this instance.
//
// See: https://developer.apple.com/documentation/foundation/nsurl/customplaygroundquicklook
func (u NSURL) CustomPlaygroundQuickLook() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("customPlaygroundQuickLook"))
	return objectivec.Object{ID: rv}
}
func (u NSURL) SetCustomPlaygroundQuickLook(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}
// An array of UTI strings representing the types of data that can be loaded
// for an item provider.
//
// # Discussion
// 
// Provide uniform type identifiers (UTIs) in order from highest fidelity to
// lowest. If your app employs a native data representation, place that first
// in the array.
// 
// Use the instance version of this property when you initialize an item
// provider with an object. As possible, implement this property to provide an
// extended array of UTIs based on the object. For example, for an [NSURL]
// object, your implementation could offer the `public.File()-url` UTI, in
// addition to the `public.Url()` UTI, if your implementation detects that the
// stored URL uses the `//` scheme.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.property
func (u NSURL) WritableTypeIdentifiersForItemProvider() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSItemProviderReading
			

			// Protocol methods for NSItemProviderWriting
			

			// Protocol methods for NSSecureCoding
			

// LoadDataWithTypeIdentifierForItemProvider is a synchronous wrapper around [NSURL.LoadDataWithTypeIdentifierForItemProviderCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u NSURL) LoadDataWithTypeIdentifierForItemProvider(ctx context.Context, typeIdentifier string) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	u.LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

