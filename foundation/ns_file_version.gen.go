// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileVersion] class.
var (
	_NSFileVersionClass     NSFileVersionClass
	_NSFileVersionClassOnce sync.Once
)

func getNSFileVersionClass() NSFileVersionClass {
	_NSFileVersionClassOnce.Do(func() {
		_NSFileVersionClass = NSFileVersionClass{class: objc.GetClass("NSFileVersion")}
	})
	return _NSFileVersionClass
}

// GetNSFileVersionClass returns the class object for NSFileVersion.
func GetNSFileVersionClass() NSFileVersionClass {
	return getNSFileVersionClass()
}

type NSFileVersionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileVersionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileVersionClass) Alloc() NSFileVersion {
	rv := objc.Send[NSFileVersion](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A snapshot of a file at a specific point in time.
//
// # Overview
// 
// Use the methods of this class to access, create, and manage file revisions
// in your app.
// 
// Each file version instance contains metadata about a single revision,
// including the location of the associated file, the modification date of the
// revision, and whether the revision is discardable.
// 
// In Mac apps, you can use file version objects to track changes to a local
// file over time and to prevent the loss of data during editing. When
// managing local versions, the document architecture creates versions at
// specific points in the lifetime of your application. Your application can
// also create versions explicitly at times that your application designates
// as appropriate.
// 
// In addition to managing local files, the system also uses this class to
// manage cloud-based files. For files in the cloud, there is usually only one
// version of the file at any given time. However, additional file versions
// may be created in cases where two different computers attempt to save the
// file to the cloud at the same time. In that case, one file is chosen as the
// current version and any other versions are tagged as being in conflict with
// the original. Conflict versions are reported to the appropriate file
// presenter objects and should be resolved as soon as possible so that the
// corresponding files can be removed from the cloud.
//
// # Accessing the Version Information
//
//   - [NSFileVersion.URL]: The URL identifying the location of the file associated with the file version object.
//   - [NSFileVersion.LocalizedName]: The string containing the user-presentable name of the file version.
//   - [NSFileVersion.LocalizedNameOfSavingComputer]: The user-presentable name of the computer on which the revision was saved.
//   - [NSFileVersion.ModificationDate]: The modification date of the version.
//   - [NSFileVersion.PersistentIdentifier]: The identifier for this version of the file.
//   - [NSFileVersion.Discardable]: A Boolean value that specifies whether the system can delete the associated file at some future time.
//   - [NSFileVersion.SetDiscardable]
//
// # Handling Version Conflicts
//
//   - [NSFileVersion.Conflict]: A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//   - [NSFileVersion.Resolved]: A Boolean value that indicates if the version object is in conflict or not.
//   - [NSFileVersion.SetResolved]
//
// # Replacing and Deleting Versions
//
//   - [NSFileVersion.ReplaceItemAtURLOptionsError]: Replace the contents of the specified file with the contents of the current version’s file.
//   - [NSFileVersion.RemoveAndReturnError]: Remove this version object and its associated file from the version store.
//
// # Instance Properties
//
//   - [NSFileVersion.HasLocalContents]
//   - [NSFileVersion.HasThumbnail]
//   - [NSFileVersion.OriginatorNameComponents]
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion
type NSFileVersion struct {
	objectivec.Object
}

// NSFileVersionFromID constructs a [NSFileVersion] from an objc.ID.
//
// A snapshot of a file at a specific point in time.
func NSFileVersionFromID(id objc.ID) NSFileVersion {
	return NSFileVersion{objectivec.Object{ID: id}}
}
// NOTE: NSFileVersion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileVersion] class.
//
// # Accessing the Version Information
//
//   - [INSFileVersion.URL]: The URL identifying the location of the file associated with the file version object.
//   - [INSFileVersion.LocalizedName]: The string containing the user-presentable name of the file version.
//   - [INSFileVersion.LocalizedNameOfSavingComputer]: The user-presentable name of the computer on which the revision was saved.
//   - [INSFileVersion.ModificationDate]: The modification date of the version.
//   - [INSFileVersion.PersistentIdentifier]: The identifier for this version of the file.
//   - [INSFileVersion.Discardable]: A Boolean value that specifies whether the system can delete the associated file at some future time.
//   - [INSFileVersion.SetDiscardable]
//
// # Handling Version Conflicts
//
//   - [INSFileVersion.Conflict]: A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//   - [INSFileVersion.Resolved]: A Boolean value that indicates if the version object is in conflict or not.
//   - [INSFileVersion.SetResolved]
//
// # Replacing and Deleting Versions
//
//   - [INSFileVersion.ReplaceItemAtURLOptionsError]: Replace the contents of the specified file with the contents of the current version’s file.
//   - [INSFileVersion.RemoveAndReturnError]: Remove this version object and its associated file from the version store.
//
// # Instance Properties
//
//   - [INSFileVersion.HasLocalContents]
//   - [INSFileVersion.HasThumbnail]
//   - [INSFileVersion.OriginatorNameComponents]
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion
type INSFileVersion interface {
	objectivec.IObject

	// Topic: Accessing the Version Information

	// The URL identifying the location of the file associated with the file version object.
	URL() INSURL
	// The string containing the user-presentable name of the file version.
	LocalizedName() string
	// The user-presentable name of the computer on which the revision was saved.
	LocalizedNameOfSavingComputer() string
	// The modification date of the version.
	ModificationDate() INSDate
	// The identifier for this version of the file.
	PersistentIdentifier() NSCoding
	// A Boolean value that specifies whether the system can delete the associated file at some future time.
	Discardable() bool
	SetDiscardable(value bool)

	// Topic: Handling Version Conflicts

	// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
	Conflict() bool
	// A Boolean value that indicates if the version object is in conflict or not.
	Resolved() bool
	SetResolved(value bool)

	// Topic: Replacing and Deleting Versions

	// Replace the contents of the specified file with the contents of the current version’s file.
	ReplaceItemAtURLOptionsError(url INSURL, options NSFileVersionReplacingOptions) (INSURL, error)
	// Remove this version object and its associated file from the version store.
	RemoveAndReturnError() (bool, error)

	// Topic: Instance Properties

	HasLocalContents() bool
	HasThumbnail() bool
	OriginatorNameComponents() INSPersonNameComponents
}

// Init initializes the instance.
func (f NSFileVersion) Init() NSFileVersion {
	rv := objc.Send[NSFileVersion](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileVersion) Autorelease() NSFileVersion {
	rv := objc.Send[NSFileVersion](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileVersion creates a new NSFileVersion instance.
func NewNSFileVersion() NSFileVersion {
	class := getNSFileVersionClass()
	rv := objc.Send[NSFileVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Replace the contents of the specified file with the contents of the current
// version’s file.
//
// url: The file whose contents you want to replace. If the file at this URL does
// not exist, a new file is created at the location.
//
// options: Specify `0` to overwrite the file in place; otherwise, specify one of the
// constants described in [NSFileVersion.ReplacingOptions].
// //
// [NSFileVersion.ReplacingOptions]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions
//
// # Return Value
// 
// The URL of the file that was written, which may be different than the one
// specified in the `url` parameter.
//
// # Discussion
// 
// When replacing the contents of the file, this method does not normally
// replace the display name associated with the file. The only exception is
// when the file at `url` is of a different type than the file associated with
// this version object. In such a case, the file name remains the same but its
// filename extension changes to match the type of the new contents. (Of
// course, if filename extension hiding is enabled, this change is not
// noticeable to users.)
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/replaceItem(at:options:)
func (f NSFileVersion) ReplaceItemAtURLOptionsError(url INSURL, options NSFileVersionReplacingOptions) (INSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("replaceItemAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSURL{}, NSErrorFrom(errorPtr)
	}
	return NSURLFromID(rv), nil

}
// Remove this version object and its associated file from the version store.
//
// # Discussion
// 
// This method removes this version object and its file from the version
// store, freeing up the associated storage space. You must not call this
// method for the current file version—that is, the version object returned
// by the [CurrentVersionOfItemAtURL] method.
// 
// You should always remove file versions as part of a coordinated write
// operation to a file. In other words, always call this method from a block
// passed to a file coordinator object to initiate a write operation. Doing so
// ensures that no other processes are operating on the file while you remove
// the version information.
// 
// If successful, subsequent requests for the versions of the file do not
// include this version object (or any object with the same information). You
// can use this method to free up disk space by removing versions that are no
// longer needed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/remove()
func (f NSFileVersion) RemoveAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("removeAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the most recent version object for the file at the specified URL.
//
// url: The URL of the file whose version object you want.
//
// # Return Value
// 
// The version object representing the current version of the file or `nil` if
// there is no such file.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/currentVersionOfItem(at:)
func (_NSFileVersionClass NSFileVersionClass) CurrentVersionOfItemAtURL(url INSURL) NSFileVersion {
	rv := objc.Send[objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("currentVersionOfItemAtURL:"), url)
	return NSFileVersionFromID(rv)
}
// Returns all versions of the specified file except the current version.
//
// url: The URL of the file whose versions you want.
//
// # Return Value
// 
// An array of file version objects or `nil` if there is no such file. The
// array does not contain the version object returned by the
// [CurrentVersionOfItemAtURL] method.
//
// # Discussion
// 
// For locally based files, this property typically contains versions of the
// file that you saved explicitly or that were saved at appropriate times
// while the file was being edited. For documents residing in the cloud, this
// property typically returns zero or more file versions representing
// conflicting versions of a file that need to be resolved with the current
// version.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/otherVersionsOfItem(at:)
func (_NSFileVersionClass NSFileVersionClass) OtherVersionsOfItemAtURL(url INSURL) []NSFileVersion {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("otherVersionsOfItemAtURL:"), url)
	return objc.ConvertSlice(rv, func(id objc.ID) NSFileVersion {
		return NSFileVersionFromID(id)
	})
}
// Returns the version of the file that has the specified persistent ID.
//
// url: The URL of the file whose version you want.
//
// persistentIdentifier: The persistent ID of the [NSFileVersion] object you want.
//
// # Return Value
// 
// The file version object with the specified ID or `nil` if no such version
// object exists.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/version(itemAt:forPersistentIdentifier:)
func (_NSFileVersionClass NSFileVersionClass) VersionOfItemAtURLForPersistentIdentifier(url INSURL, persistentIdentifier objectivec.IObject) NSFileVersion {
	rv := objc.Send[objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("versionOfItemAtURL:forPersistentIdentifier:"), url, persistentIdentifier)
	return NSFileVersionFromID(rv)
}
// Creates and returns a temporary directory to use for saving the contents of
// the file.
//
// url: The URL of the file whose contents you want to save.
//
// # Return Value
// 
// A URL identifying the temporary directory in which to create a the new
// file. You must delete the directory specified by this URL after you have
// created the file and moved it to its proper location.
//
// # Discussion
// 
// You can use this method in situations where you want to create a file in a
// temporary location. For example, you might use this method when saving the
// contents of a file to disk for the first time. When you finish creating the
// temporary file, move it to a more appropriate location, such as the
// user’s [Documents] directory. You must delete the directory returned by
// this method when you are done with it.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/temporaryDirectoryURLForNewVersionOfItem(at:)
func (_NSFileVersionClass NSFileVersionClass) TemporaryDirectoryURLForNewVersionOfItemAtURL(url INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("temporaryDirectoryURLForNewVersionOfItemAtURL:"), url)
	return NSURLFromID(rv)
}
// Creates a version of the file at the specified location.
//
// url: The location at which to store the new file version.
//
// contentsURL: The URL that specifies the file to use for the version contents.
//
// options: Specify 0 for this parameter if you want to create a copy of the file at
// the location specified by the `url` parameter. Alternatively, specify one
// of the constants described in [NSFileVersion.AddingOptions].
// //
// [NSFileVersion.AddingOptions]: https://developer.apple.com/documentation/Foundation/NSFileVersion/AddingOptions
//
// # Return Value
// 
// The file version object representing the new version or `nil` if an error
// occurred.
//
// # Discussion
// 
// You can use this method to save a version of your file to the location
// specified by the `url` parameter. The contents of the file are taken from
// the `contentsURL` parameter, whose value may be the same as the `url`
// parameter.
// 
// You should always add file versions as part of a coordinated write
// operation to a file. In other words, always call this method from a block
// passed to a file coordinator object to initiate a write operation. Doing so
// ensures that no other processes are operating on the file while you save
// the version to its new location.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/addOfItem(at:withContentsOf:options:)
func (_NSFileVersionClass NSFileVersionClass) AddVersionOfItemAtURLWithContentsOfURLOptionsError(url INSURL, contentsURL INSURL, options NSFileVersionAddingOptions) (NSFileVersion, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("addVersionOfItemAtURL:withContentsOfURL:options:error:"), url, contentsURL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSFileVersion{}, NSErrorFrom(errorPtr)
	}
	return NSFileVersionFromID(rv), nil

}
// Returns an array of version objects that are currently in conflict for the
// specified URL.
//
// url: The URL of the file that has associated version objects.
//
// # Return Value
// 
// An array of [NSFileVersion] objects that represent the versions in conflict
// or `nil` if the file at URL does not exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/unresolvedConflictVersionsOfItem(at:)
func (_NSFileVersionClass NSFileVersionClass) UnresolvedConflictVersionsOfItemAtURL(url INSURL) []NSFileVersion {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFileVersionClass.class), objc.Sel("unresolvedConflictVersionsOfItemAtURL:"), url)
	return objc.ConvertSlice(rv, func(id objc.ID) NSFileVersion {
		return NSFileVersionFromID(id)
	})
}
// Removes all versions of a file, except the current one, from the version
// store.
//
// url: The file whose older versions you want to delete. If the file at this URL
// does not exist, a new file is created at the location.
//
// # Discussion
// 
// This method removes all versions except the current one from the version
// store, freeing up the associated storage space.
// 
// You should always remove file versions as part of a coordinated write
// operation to a file. In other words, always call this method from a block
// passed to a file coordinator object to initiate a write operation. Doing so
// ensures that no other processes are operating on the file while you remove
// the version information.
// 
// If successful, subsequent requests for the versions of the file reflect
// that only the current version is available. You can use this method to free
// up disk space by removing versions that are no longer needed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/removeOtherVersionsOfItem(at:)
func (_NSFileVersionClass NSFileVersionClass) RemoveOtherVersionsOfItemAtURLError(url INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSFileVersionClass.class), objc.Sel("removeOtherVersionsOfItemAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeOtherVersionsOfItemAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The URL identifying the location of the file associated with the file
// version object.
//
// # Discussion
// 
// The URL identifies the location of the file associated with this version.
// If this version of the file has been deleted, the value in this property is
// `nil`.
// 
// Do not display any part of this URL to the user. The location of file
// versions is managed by the system and should not be exposed to the user. If
// you want to present the name of a file version, use the [LocalizedName]
// property.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/url
func (f NSFileVersion) URL() INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}
// The string containing the user-presentable name of the file version.
//
// # Discussion
// 
// When displaying different versions of a file to the user, you should
// present this string to the user instead of the version’s URL.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedName
func (f NSFileVersion) LocalizedName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localizedName"))
	return NSStringFromID(rv).String()
}
// The user-presentable name of the computer on which the revision was saved.
//
// # Discussion
// 
// If the current revision has been deleted from disk, or if no computer name
// was recorded, the value in this property is `nil`. The computer name is
// guaranteed to be recorded only when the current version is in conflict with
// another version. The version object does not track changes to the computer
// name itself. Thus, if the computer name changed, the value in this string
// might be an old value.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedNameOfSavingComputer
func (f NSFileVersion) LocalizedNameOfSavingComputer() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localizedNameOfSavingComputer"))
	return NSStringFromID(rv).String()
}
// The modification date of the version.
//
// # Discussion
// 
// If the version has been deleted, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/modificationDate
func (f NSFileVersion) ModificationDate() INSDate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("modificationDate"))
	return NSDateFromID(objc.ID(rv))
}
// The identifier for this version of the file.
//
// # Discussion
// 
// You can save the value of this property persistently and use it to recreate
// the version object later. When recreating the version object using the
// [VersionOfItemAtURLForPersistentIdentifier] method, the version object
// returned is equivalent to the current object.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/persistentIdentifier
func (f NSFileVersion) PersistentIdentifier() NSCoding {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("persistentIdentifier"))
	return NSCodingObjectFromID(rv)
}
// A Boolean value that specifies whether the system can delete the associated
// file at some future time.
//
// # Discussion
// 
// Marking a file version as discardable gives the system the flexibility to
// reclaim the space, occupied by the associated file, at some future time. Do
// not, however, depend on the file being discarded.
// 
// After setting this property to [true], do not set this property to [false]
// again. Doing so causes the system to raise an exception. In addition, if
// you set this property to [true] for the version of the file returned by the
// [CurrentVersionOfItemAtURL] method, the system raises an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/isDiscardable
func (f NSFileVersion) Discardable() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isDiscardable"))
	return rv
}
func (f NSFileVersion) SetDiscardable(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setDiscardable:"), value)
}
// A Boolean value indicating whether the contents of the version are in
// conflict with the contents of another version.
//
// # Discussion
// 
// When two or more versions of a file are written at the same time, perhaps
// because the file is saved in the cloud and one or more of the writers were
// offline when they were writing, the system attempts to resolve the conflict
// automatically. It does this by picking one of the file versions to be the
// current file and setting this property to [true] for the other file
// versions that are in conflict.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/isConflict
func (f NSFileVersion) Conflict() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isConflict"))
	return rv
}
// A Boolean value that indicates if the version object is in conflict or not.
//
// # Discussion
// 
// When the system detects a conflict involving versions of a file, it sets
// this property to [false] to indicate an unresolved conflict. After you
// resolve the conflict, set this property to [true] to tell the system it is
// resolved; you must then remove any versions of the file that are no longer
// useful.
// 
// To remove an unused version of a file, call the [RemoveAndReturnError]
// method. To remove all unused versions of a file, call the
// [RemoveOtherVersionsOfItemAtURLError] method.
// 
// Resolving a conflict causes the file version object to be removed from any
// reports about conflicting versions, such as those returned by the
// [UnresolvedConflictVersionsOfItemAtURL] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/isResolved
func (f NSFileVersion) Resolved() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isResolved"))
	return rv
}
func (f NSFileVersion) SetResolved(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setResolved:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasLocalContents
func (f NSFileVersion) HasLocalContents() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasLocalContents"))
	return rv
}
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasThumbnail
func (f NSFileVersion) HasThumbnail() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasThumbnail"))
	return rv
}
// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/originatorNameComponents
func (f NSFileVersion) OriginatorNameComponents() INSPersonNameComponents {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("originatorNameComponents"))
	return NSPersonNameComponentsFromID(objc.ID(rv))
}

