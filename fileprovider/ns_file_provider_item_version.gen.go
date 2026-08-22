// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderItemVersion] class.
var (
	_NSFileProviderItemVersionClass     NSFileProviderItemVersionClass
	_NSFileProviderItemVersionClassOnce sync.Once
)

func getNSFileProviderItemVersionClass() NSFileProviderItemVersionClass {
	_NSFileProviderItemVersionClassOnce.Do(func() {
		_NSFileProviderItemVersionClass = NSFileProviderItemVersionClass{class: objc.GetClass("NSFileProviderItemVersion")}
	})
	return _NSFileProviderItemVersionClass
}

// GetNSFileProviderItemVersionClass returns the class object for NSFileProviderItemVersion.
func GetNSFileProviderItemVersionClass() NSFileProviderItemVersionClass {
	return getNSFileProviderItemVersionClass()
}

type NSFileProviderItemVersionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderItemVersionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderItemVersionClass) Alloc() NSFileProviderItemVersion {
	rv := objc.Send[NSFileProviderItemVersion](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The version of the item’s content and its metadata.
//
// # Overview
//
// Each item has a separate version object for its metadata and its content.
// As a result, the file provider can update an item’s metadata without
// uploading or downloading a new copy of its content.
//
// # Creating Version Instances
//
//   - [NSFileProviderItemVersion.InitWithContentVersionMetadataVersion]: Creates a new version object.
//
// # Accessing Version Data
//
//   - [NSFileProviderItemVersion.ContentVersion]: An opaque object used to track versions of the item’s content.
//   - [NSFileProviderItemVersion.MetadataVersion]: An opaque object used to track versions of the item’s metadata.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion
type NSFileProviderItemVersion struct {
	objectivec.Object
}

// NSFileProviderItemVersionFromID constructs a [NSFileProviderItemVersion] from an objc.ID.
//
// The version of the item’s content and its metadata.
func NSFileProviderItemVersionFromID(id objc.ID) NSFileProviderItemVersion {
	return NSFileProviderItemVersion{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderItemVersion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderItemVersion] class.
//
// # Creating Version Instances
//
//   - [INSFileProviderItemVersion.InitWithContentVersionMetadataVersion]: Creates a new version object.
//
// # Accessing Version Data
//
//   - [INSFileProviderItemVersion.ContentVersion]: An opaque object used to track versions of the item’s content.
//   - [INSFileProviderItemVersion.MetadataVersion]: An opaque object used to track versions of the item’s metadata.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion
type INSFileProviderItemVersion interface {
	objectivec.IObject

	// Topic: Creating Version Instances

	// Creates a new version object.
	InitWithContentVersionMetadataVersion(contentVersion foundation.NSData, metadataVersion foundation.NSData) NSFileProviderItemVersion

	// Topic: Accessing Version Data

	// An opaque object used to track versions of the item’s content.
	ContentVersion() foundation.NSData
	// An opaque object used to track versions of the item’s metadata.
	MetadataVersion() foundation.NSData
}

// Init initializes the instance.
func (f NSFileProviderItemVersion) Init() NSFileProviderItemVersion {
	rv := objc.Send[NSFileProviderItemVersion](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderItemVersion) Autorelease() NSFileProviderItemVersion {
	rv := objc.Send[NSFileProviderItemVersion](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderItemVersion creates a new NSFileProviderItemVersion instance.
func NewNSFileProviderItemVersion() NSFileProviderItemVersion {
	class := getNSFileProviderItemVersionClass()
	rv := objc.Send[NSFileProviderItemVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new version object.
//
// contentVersion: An opaque version object for the item’s content.
//
// metadataVersion: An opaque version object for the item’s metadata.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/init(contentVersion:metadataVersion:)
func NewFileProviderItemVersionWithContentVersionMetadataVersion(contentVersion foundation.NSData, metadataVersion foundation.NSData) NSFileProviderItemVersion {
	instance := getNSFileProviderItemVersionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentVersion:metadataVersion:"), contentVersion, metadataVersion)
	return NSFileProviderItemVersionFromID(rv)
}

// Creates a new version object.
//
// contentVersion: An opaque version object for the item’s content.
//
// metadataVersion: An opaque version object for the item’s metadata.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/init(contentVersion:metadataVersion:)
func (f NSFileProviderItemVersion) InitWithContentVersionMetadataVersion(contentVersion foundation.NSData, metadataVersion foundation.NSData) NSFileProviderItemVersion {
	rv := objc.Send[NSFileProviderItemVersion](f.ID, objc.Sel("initWithContentVersion:metadataVersion:"), contentVersion, metadataVersion)
	return rv
}

// An opaque object used to track versions of the item’s content.
//
// # Discussion
//
// If the system stores a local copy of an item’s content, it downloads a
// new copy when the [NSFileProviderItemVersion.ContentVersion] changes. The
// content version also invalidates the system’s thumbnail cache.
//
// The system considers the file’s resource fork part of the file’s
// content. The version changes when either the data fork or the resource fork
// changes.
//
// The version data object must be no longer than 128 bytes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/contentVersion
func (f NSFileProviderItemVersion) ContentVersion() foundation.NSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("contentVersion"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// An opaque object used to track versions of the item’s metadata.
//
// # Discussion
//
// When the [NSFileProviderItemVersion.MetadataVersion] changes, system
// updates the dataless representation of the item on disk, but it doesn’t
// attempt to download the content.
//
// The version data object must be no longer than 128 bytes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/metadataVersion
func (f NSFileProviderItemVersion) MetadataVersion() foundation.NSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("metadataVersion"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A Boolean value indicating that this version predates the version returned
// by the file provider extension.
//
// # Discussion
//
// The system uses this property to represent an item that doesn’t have a
// corresponding version provided by the file provider extension.
//
// When creating an item by calling the
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// method, if your file provider extension returns an item that doesn’t
// match the template, the system tries to apply the necessary changes before
// saving the item to disk. However, if the system detects conflicts with the
// version on disk, it sends the new item back to your file provider extension
// by calling either the
// [ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
// of [DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler]
// methods with a `baseVersion` property that represents the item passed to
// the
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/beforeFirstSyncComponent
func (_NSFileProviderItemVersionClass NSFileProviderItemVersionClass) BeforeFirstSyncComponent() foundation.NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSFileProviderItemVersionClass.class), objc.Sel("beforeFirstSyncComponent"))
	return foundation.NSDataFromID(objc.ID(rv))
}
