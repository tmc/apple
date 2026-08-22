// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// A protocol that defines the properties of an item managed by the File Provider extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol
type NSFileProviderItem interface {
	objectivec.IObject

	// The item’s persistent identifier.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// The item’s filename.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/filename
	Filename() string

	// The item’s Uniform Type Identifier (UTI).
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/typeIdentifier
	TypeIdentifier() string

	// The item’s Uniform Type Identifier (UTI).
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentType
	ContentType() uniformtypeidentifiers.UTType

	// The item’s capabilities.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/capabilities
	Capabilities() NSFileProviderItemCapabilities

	// The number of items contained by this item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/childItemCount
	ChildItemCount() foundation.NSNumber

	// The document’s size, in bytes.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/documentSize
	DocumentSize() foundation.NSNumber

	// contentPolicy protocol.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentPolicy
	ContentPolicy() NSFileProviderContentPolicy

	// The persistent identifier of the item’s parent folder.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/parentItemIdentifier
	ParentItemIdentifier() NSFileProviderItemIdentifier

	// A Boolean value that indicates whether an item is in the trash.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isTrashed
	IsTrashed() bool

	// The target of the symlink.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/symlinkTargetPath
	SymlinkTargetPath() string

	// The date the item was last modified.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentModificationDate
	ContentModificationDate() foundation.NSDate

	// The date the item was created.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/creationDate
	CreationDate() foundation.NSDate

	// The date the item was last used.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/lastUsedDate
	LastUsedDate() foundation.NSDate

	// A version object that tracks changes to an item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/itemVersion
	ItemVersion() INSFileProviderItemVersion

	// A data value used to determine when the item changes.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/versionIdentifier
	VersionIdentifier() foundation.NSData

	// A Boolean value that indicates whether the item is the most recent version downloaded from the server.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isMostRecentVersionDownloaded
	IsMostRecentVersionDownloaded() bool

	// A Boolean value that indicates whether the item is currently uploading to your remote server.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isUploading
	IsUploading() bool

	// A Boolean value that indicates whether the item has been uploaded to your remote server.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isUploaded
	IsUploaded() bool

	// An object describing an error that occurred while uploading the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/uploadingError
	UploadingError() foundation.NSError

	// A Boolean value that indicates whether the item is currently downloading from your remote server.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isDownloading
	IsDownloading() bool

	// A Boolean value that indicates whether the item has been downloaded from your remote server.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isDownloaded
	IsDownloaded() bool

	// An object describing an error that occurred while downloading the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/downloadingError
	DownloadingError() foundation.NSError

	// A Boolean value that indicates whether the item is shared with other users.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isShared
	IsShared() bool

	// A Boolean value that indicates whether the item was shared by the current user.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isSharedByCurrentUser
	IsSharedByCurrentUser() bool

	// The most recent editor’s name.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/mostRecentEditorNameComponents
	MostRecentEditorNameComponents() foundation.NSPersonNameComponents

	// The name of the item’s owner.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/ownerNameComponents
	OwnerNameComponents() foundation.NSPersonNameComponents

	// The extended file attributes synced by the File Provider extension.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/extendedAttributes
	ExtendedAttributes() foundation.INSDictionary

	// Flags that define an item’s on-disk properties and its appearance in the user interface.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/fileSystemFlags
	FileSystemFlags() NSFileProviderFileSystemFlags

	// An abstract data blob representing the tags associated with the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/tagData
	TagData() foundation.NSData

	// A property list that contains additional data about the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/userInfo
	UserInfo() foundation.INSDictionary

	// A 64-bit, unsigned integer indicating the order of the favorite item in the Favorites list.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/favoriteRank
	FavoriteRank() foundation.NSNumber

	// The file type and creator codes for the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/typeAndCreator
	TypeAndCreator() NSFileProviderTypeAndCreator
}

// NSFileProviderItemObject wraps an existing Objective-C object that conforms to the NSFileProviderItem protocol.
type NSFileProviderItemObject struct {
	objectivec.Object
}

func (o NSFileProviderItemObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderItemObjectFromID constructs a [NSFileProviderItemObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderItemObjectFromID(id objc.ID) NSFileProviderItemObject {
	return NSFileProviderItemObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The item’s persistent identifier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/itemIdentifier
func (o NSFileProviderItemObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The item’s filename.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/filename
func (o NSFileProviderItemObject) Filename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filename"))
	return foundation.NSStringFromID(rv).String()
}

// The item’s Uniform Type Identifier (UTI).
//
// # Discussion
//
// Your extension must provide either the [TypeIdentifier] or the
// [ContentType] property. Use the [TypeIdentifier] property only in iOS 13 or
// earlier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/typeIdentifier
func (o NSFileProviderItemObject) TypeIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("typeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The item’s Uniform Type Identifier (UTI).
//
// # Discussion
//
// Your extension must provide either the [TypeIdentifier] or the
// [ContentType] property. Where possible, use the [ContentType] property. Use
// the [TypeIdentifier] property only in iOS 13 or earlier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentType
func (o NSFileProviderItemObject) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(rv)
}

// The item’s capabilities.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/capabilities
func (o NSFileProviderItemObject) Capabilities() NSFileProviderItemCapabilities {
	rv := objc.Send[NSFileProviderItemCapabilities](o.ID, objc.Sel("capabilities"))
	return NSFileProviderItemCapabilities(rv)
}

// The number of items contained by this item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/childItemCount
func (o NSFileProviderItemObject) ChildItemCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("childItemCount"))
	return foundation.NSNumberFromID(rv)
}

// The document’s size, in bytes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/documentSize
func (o NSFileProviderItemObject) DocumentSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("documentSize"))
	return foundation.NSNumberFromID(rv)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentPolicy
func (o NSFileProviderItemObject) ContentPolicy() NSFileProviderContentPolicy {
	rv := objc.Send[NSFileProviderContentPolicy](o.ID, objc.Sel("contentPolicy"))
	return NSFileProviderContentPolicy(rv)
}

// The persistent identifier of the item’s parent folder.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/parentItemIdentifier
func (o NSFileProviderItemObject) ParentItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parentItemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// A Boolean value that indicates whether an item is in the trash.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isTrashed
func (o NSFileProviderItemObject) IsTrashed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isTrashed"))
	return bool(rv)
}

// The target of the symlink.
//
// # Discussion
//
// If the extension contains an item with a [TypeIdentifier] of
// `public.Symlink()` (`kUTTypeSymLink`), this property contains the
// symlink’s target. Otherwise, it’s `nil`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/symlinkTargetPath
func (o NSFileProviderItemObject) SymlinkTargetPath() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("symlinkTargetPath"))
	return foundation.NSStringFromID(rv).String()
}

// The date the item was last modified.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/contentModificationDate
func (o NSFileProviderItemObject) ContentModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentModificationDate"))
	return foundation.NSDateFromID(rv)
}

// The date the item was created.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/creationDate
func (o NSFileProviderItemObject) CreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("creationDate"))
	return foundation.NSDateFromID(rv)
}

// The date the item was last used.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/lastUsedDate
func (o NSFileProviderItemObject) LastUsedDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lastUsedDate"))
	return foundation.NSDateFromID(rv)
}

// A version object that tracks changes to an item.
//
// # Discussion
//
// The version object lets you track changes to an item’s content and
// metadata separately. Updating the version also invalidates the thumbnail
// cache. For more information, see [NSFileProviderItemVersion].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/itemVersion
func (o NSFileProviderItemObject) ItemVersion() INSFileProviderItemVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemVersion"))
	return NSFileProviderItemVersionFromID(rv)
}

// A data value used to determine when the item changes.
//
// # Discussion
//
// This property contains a data object that can uniquely identify each
// version of the item; for example, the hash of a document’s contents.
//
// Version identifiers are limited to 1000 bytes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/versionIdentifier
func (o NSFileProviderItemObject) VersionIdentifier() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("versionIdentifier"))
	return foundation.NSDataFromID(rv)
}

// A Boolean value that indicates whether the item is the most recent version
// downloaded from the server.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isMostRecentVersionDownloaded
func (o NSFileProviderItemObject) IsMostRecentVersionDownloaded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isMostRecentVersionDownloaded"))
	return bool(rv)
}

// A Boolean value that indicates whether the item is currently uploading to
// your remote server.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isUploading
func (o NSFileProviderItemObject) IsUploading() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUploading"))
	return bool(rv)
}

// A Boolean value that indicates whether the item has been uploaded to your
// remote server.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isUploaded
func (o NSFileProviderItemObject) IsUploaded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUploaded"))
	return bool(rv)
}

// An object describing an error that occurred while uploading the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/uploadingError
func (o NSFileProviderItemObject) UploadingError() foundation.NSError {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uploadingError"))
	return foundation.NSErrorFromID(rv)
}

// A Boolean value that indicates whether the item is currently downloading
// from your remote server.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isDownloading
func (o NSFileProviderItemObject) IsDownloading() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDownloading"))
	return bool(rv)
}

// A Boolean value that indicates whether the item has been downloaded from
// your remote server.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isDownloaded
func (o NSFileProviderItemObject) IsDownloaded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDownloaded"))
	return bool(rv)
}

// An object describing an error that occurred while downloading the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/downloadingError
func (o NSFileProviderItemObject) DownloadingError() foundation.NSError {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("downloadingError"))
	return foundation.NSErrorFromID(rv)
}

// A Boolean value that indicates whether the item is shared with other users.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isShared
func (o NSFileProviderItemObject) IsShared() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isShared"))
	return bool(rv)
}

// A Boolean value that indicates whether the item was shared by the current
// user.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/isSharedByCurrentUser
func (o NSFileProviderItemObject) IsSharedByCurrentUser() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSharedByCurrentUser"))
	return bool(rv)
}

// The most recent editor’s name.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/mostRecentEditorNameComponents
func (o NSFileProviderItemObject) MostRecentEditorNameComponents() foundation.NSPersonNameComponents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mostRecentEditorNameComponents"))
	return foundation.NSPersonNameComponentsFromID(rv)
}

// The name of the item’s owner.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/ownerNameComponents
func (o NSFileProviderItemObject) OwnerNameComponents() foundation.NSPersonNameComponents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("ownerNameComponents"))
	return foundation.NSPersonNameComponentsFromID(rv)
}

// The extended file attributes synced by the File Provider extension.
//
// # Discussion
//
// The extended file attributes are part of the item’s metadata. The system
// sets extended attributes on dataless files, and preserves them on files
// that it renders dataless. The system decides which attributes to sync. To
// sync an attribute, it calls the `xattr_name_with_flags(_:_:)` method and
// passes the `XATTR_FLAG_SYNCABLE` flag. Some older attributes are also
// synced.
//
// The system caps the syncable extended attributes to about 32KiB total for
// each item. If the extended attributes exceed this limit, the system
// automatically makes some of the attributes nonsyncable.
//
// The system also decides which nonsyncable attributes it preserves on the
// local copy when a remote item changes. For example, it preserves attributes
// created by calling `xattr_name_with_flags(_:_:)` and passing
// `XATTR_FLAG_CONTENT_DEPENDENT` as long as the remote change didn’t modify
// the [NSFileProviderItemVersion.ContentVersion] property for the item’s
// version.
//
// This dictionary doesn’t list extended attributes that are already covered
// by [NSFileProviderItemFields] values, like [NSFileProviderItemLastUsedDate]
// or [NSFileProviderItemTagData]. Similarly, the dictionary doesn’t include
// the 32 bits of Finder info stored in an extended attribute named
// `com.Apple().FinderInfo`, because this information exists in other
// [NSFileProviderItem] properties.
//
// Also, the resource fork is content and isn’t included in the extended
// attributes dictionary. If your extension detects remote changes to the
// resource fork, report these changes by modifying the item version’s
// [NSFileProviderItemVersion.ContentVersion] property.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/extendedAttributes
//
// [NSFileProviderItemFields]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields
func (o NSFileProviderItemObject) ExtendedAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("extendedAttributes"))
	return foundation.NSDictionaryFromID(rv)
}

// Flags that define an item’s on-disk properties and its appearance in the
// user interface.
//
// # Discussion
//
// The flags define the on-disk properties of the item. The system modifies
// the item’s appearance based on these flags.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/fileSystemFlags
func (o NSFileProviderItemObject) FileSystemFlags() NSFileProviderFileSystemFlags {
	rv := objc.Send[NSFileProviderFileSystemFlags](o.ID, objc.Sel("fileSystemFlags"))
	return NSFileProviderFileSystemFlags(rv)
}

// An abstract data blob representing the tags associated with the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/tagData
func (o NSFileProviderItemObject) TagData() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tagData"))
	return foundation.NSDataFromID(rv)
}

// A property list that contains additional data about the item.
//
// # Discussion
//
// The `userInfo` data is often used by the predicate for actions defined by
// the File Provider UI extension. For more information, see `Adding Actions
// to the Context Menu`.
//
// The `userInfo` dictionary can only accept entries with numbers (including
// Boolean values), dates, or strings as either the key or the value.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/userInfo
func (o NSFileProviderItemObject) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(rv)
}

// A 64-bit, unsigned integer indicating the order of the favorite item in the
// Favorites list.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/favoriteRank
func (o NSFileProviderItemObject) FavoriteRank() foundation.NSNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("favoriteRank"))
	return foundation.NSNumberFromID(rv)
}

// The file type and creator codes for the item.
//
// # Discussion
//
// This property contains two values: the file type code and the creator code.
// The system synchronizes both codes at the same time, so define both, even
// if you’re just changing one.
//
// If you modify this property, the system sets the
// [NSFileProviderTypeAndCreator] value passed to the
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler] or
// [ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
// methods. The system also writes the type and creator codes in the
// [FileInfo] structure, if relevant.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemProtocol/typeAndCreator
//
// [NSFileProviderTypeAndCreator]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTypeAndCreator
func (o NSFileProviderItemObject) TypeAndCreator() NSFileProviderTypeAndCreator {
	rv := objc.Send[NSFileProviderTypeAndCreator](o.ID, objc.Sel("typeAndCreator"))
	return NSFileProviderTypeAndCreator(rv)
}
