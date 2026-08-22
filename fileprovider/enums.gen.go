// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy
type NSFileProviderContentPolicy int

const (
	NSFileProviderContentPolicyDownloadEagerlyAndKeepDownloaded     NSFileProviderContentPolicy = 3
	NSFileProviderContentPolicyDownloadLazily                       NSFileProviderContentPolicy = 1
	NSFileProviderContentPolicyDownloadLazilyAndEvictOnRemoteUpdate NSFileProviderContentPolicy = 2
	NSFileProviderContentPolicyInherited                            NSFileProviderContentPolicy = 0
)

func (e NSFileProviderContentPolicy) String() string {
	switch e {
	case NSFileProviderContentPolicyDownloadEagerlyAndKeepDownloaded:
		return "NSFileProviderContentPolicyDownloadEagerlyAndKeepDownloaded"
	case NSFileProviderContentPolicyDownloadLazily:
		return "NSFileProviderContentPolicyDownloadLazily"
	case NSFileProviderContentPolicyDownloadLazilyAndEvictOnRemoteUpdate:
		return "NSFileProviderContentPolicyDownloadLazilyAndEvictOnRemoteUpdate"
	case NSFileProviderContentPolicyInherited:
		return "NSFileProviderContentPolicyInherited"
	default:
		return fmt.Sprintf("NSFileProviderContentPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions
type NSFileProviderCreateItemOptions uint

const (
	// NSFileProviderCreateItemDeletionConflicted: A value indicating a conflict for a deleted item.
	NSFileProviderCreateItemDeletionConflicted NSFileProviderCreateItemOptions = 2
	// NSFileProviderCreateItemMayAlreadyExist: An option indicating that the item may already exist in your remote storage.
	NSFileProviderCreateItemMayAlreadyExist NSFileProviderCreateItemOptions = 1
)

func (e NSFileProviderCreateItemOptions) String() string {
	switch e {
	case NSFileProviderCreateItemDeletionConflicted:
		return "NSFileProviderCreateItemDeletionConflicted"
	case NSFileProviderCreateItemMayAlreadyExist:
		return "NSFileProviderCreateItemMayAlreadyExist"
	default:
		return fmt.Sprintf("NSFileProviderCreateItemOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDeleteItemOptions
type NSFileProviderDeleteItemOptions uint

const (
	// NSFileProviderDeleteItemRecursive: A value indicating that the delete operation removes the item and all of its children.
	NSFileProviderDeleteItemRecursive NSFileProviderDeleteItemOptions = 1
)

func (e NSFileProviderDeleteItemOptions) String() string {
	switch e {
	case NSFileProviderDeleteItemRecursive:
		return "NSFileProviderDeleteItemRecursive"
	default:
		return fmt.Sprintf("NSFileProviderDeleteItemOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode
type NSFileProviderDomainRemovalMode int

const (
	// NSFileProviderDomainRemovalModePreserveDirtyUserData: Deletes the domain but keeps any items with unsynced, local changes.
	NSFileProviderDomainRemovalModePreserveDirtyUserData NSFileProviderDomainRemovalMode = 1
	// NSFileProviderDomainRemovalModePreserveDownloadedUserData: Deletes the domain, but keeps the downloaded user data.
	NSFileProviderDomainRemovalModePreserveDownloadedUserData NSFileProviderDomainRemovalMode = 2
	// NSFileProviderDomainRemovalModeRemoveAll: Deletes all items in the domain.
	NSFileProviderDomainRemovalModeRemoveAll NSFileProviderDomainRemovalMode = 0
)

func (e NSFileProviderDomainRemovalMode) String() string {
	switch e {
	case NSFileProviderDomainRemovalModePreserveDirtyUserData:
		return "NSFileProviderDomainRemovalModePreserveDirtyUserData"
	case NSFileProviderDomainRemovalModePreserveDownloadedUserData:
		return "NSFileProviderDomainRemovalModePreserveDownloadedUserData"
	case NSFileProviderDomainRemovalModeRemoveAll:
		return "NSFileProviderDomainRemovalModeRemoveAll"
	default:
		return fmt.Sprintf("NSFileProviderDomainRemovalMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct
type NSFileProviderDomainTestingModes uint

const (
	// NSFileProviderDomainTestingModeAlwaysEnabled: A testing mode that automatically enables the domain.
	NSFileProviderDomainTestingModeAlwaysEnabled NSFileProviderDomainTestingModes = 1
	// NSFileProviderDomainTestingModeInteractive: A testing mode where the extension can deterministically test asynchronous operations.
	NSFileProviderDomainTestingModeInteractive NSFileProviderDomainTestingModes = 2
)

func (e NSFileProviderDomainTestingModes) String() string {
	switch e {
	case NSFileProviderDomainTestingModeAlwaysEnabled:
		return "NSFileProviderDomainTestingModeAlwaysEnabled"
	case NSFileProviderDomainTestingModeInteractive:
		return "NSFileProviderDomainTestingModeInteractive"
	default:
		return fmt.Sprintf("NSFileProviderDomainTestingModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code
type NSFileProviderErrorCode int

const (
	// NSFileProviderErrorApplicationExtensionNotFound: An error indicating that there isn’t an app extension within the app bundle.
	NSFileProviderErrorApplicationExtensionNotFound NSFileProviderErrorCode = -2014
	// NSFileProviderErrorCannotSynchronize: An error indicating a failed sync attempt.
	NSFileProviderErrorCannotSynchronize NSFileProviderErrorCode = -2005
	// NSFileProviderErrorDeletionRejected: An error indicating a failed deletion action.
	NSFileProviderErrorDeletionRejected NSFileProviderErrorCode = -1006
	// NSFileProviderErrorDirectoryNotEmpty: An error indicating an attempt to nonrecursively delete a directory that isn’t empty.
	NSFileProviderErrorDirectoryNotEmpty NSFileProviderErrorCode = -1007
	NSFileProviderErrorDomainDisabled    NSFileProviderErrorCode = -2011
	NSFileProviderErrorExcludedFromSync  NSFileProviderErrorCode = -2010
	// NSFileProviderErrorFilenameCollision: An error indicating that an item with the same name already exists in the same directory.
	NSFileProviderErrorFilenameCollision NSFileProviderErrorCode = -1001
	// NSFileProviderErrorInsufficientQuota: An error indicating that the File Provider extension can’t upload the item because it would push the account over its quota.
	NSFileProviderErrorInsufficientQuota NSFileProviderErrorCode = -1003
	// NSFileProviderErrorLocalVersionConflictingWithServer: Returned by createItemBasedOnTemplate or modifyItem if the provider does not wish to sync the item.
	NSFileProviderErrorLocalVersionConflictingWithServer NSFileProviderErrorCode = -2015
	// NSFileProviderErrorNewerExtensionVersionFound: An error indicating that the registered provider in the system is a newer version than the one the app uses.
	NSFileProviderErrorNewerExtensionVersionFound NSFileProviderErrorCode = -2004
	// NSFileProviderErrorNoSuchItem: An error indicating that the specified item doesn’t exist.
	NSFileProviderErrorNoSuchItem NSFileProviderErrorCode = -1005
	// NSFileProviderErrorNonEvictable: An error indicating that the File Provider extension can’t evict an item.
	NSFileProviderErrorNonEvictable NSFileProviderErrorCode = -2008
	// NSFileProviderErrorNonEvictableChildren: An error indicating that the File Provider extension can’t evict a directory because it contains nonevictable items.
	NSFileProviderErrorNonEvictableChildren NSFileProviderErrorCode = -2006
	// NSFileProviderErrorNotAuthenticated: An error indicating that you can’t verify the user’s credentials.
	NSFileProviderErrorNotAuthenticated NSFileProviderErrorCode = -1000
	// NSFileProviderErrorOlderExtensionVersionRunning: An error indicating that the registered provider in the system is an older version than the one the app uses.
	NSFileProviderErrorOlderExtensionVersionRunning NSFileProviderErrorCode = -2003
	// NSFileProviderErrorPageExpired: An error indicating that the page is too old, and that the system must restart the enumeration operation from the beginning.
	NSFileProviderErrorPageExpired NSFileProviderErrorCode = -1002
	// NSFileProviderErrorProviderDomainNotFound: An error indicating that there isn’t a registered domain for the corresponding identifier.
	NSFileProviderErrorProviderDomainNotFound NSFileProviderErrorCode = -2013
	// NSFileProviderErrorProviderDomainTemporarilyUnavailable: An error indicating that the system is unable to service requests for the domain temporarily, and you can try again later.
	NSFileProviderErrorProviderDomainTemporarilyUnavailable NSFileProviderErrorCode = -2012
	// NSFileProviderErrorProviderNotFound: An error indicating that the File Provider manager can’t find the specified provider.
	NSFileProviderErrorProviderNotFound NSFileProviderErrorCode = -2001
	// NSFileProviderErrorProviderTranslocated: An error indicating the File Provider extension is in a disabled state due to Gatekeeper’s restrictions for apps from outside the App Store.
	NSFileProviderErrorProviderTranslocated NSFileProviderErrorCode = -2002
	// NSFileProviderErrorServerUnreachable: An error indicating that the File Provider extension can’t reach the remote server.
	NSFileProviderErrorServerUnreachable NSFileProviderErrorCode = -1004
	// NSFileProviderErrorSyncAnchorExpired: An error indicating that the sync anchor is too old, and that the system must restart the sync operation from the beginning.
	NSFileProviderErrorSyncAnchorExpired NSFileProviderErrorCode = -1002
	// NSFileProviderErrorUnsyncedEdits: An error indicating that the item contains unsynced changes.
	NSFileProviderErrorUnsyncedEdits NSFileProviderErrorCode = -2007
	// NSFileProviderErrorVersionNoLongerAvailable: An error indicating that the specified version is no longer available.
	NSFileProviderErrorVersionNoLongerAvailable NSFileProviderErrorCode = -2009
)

func (e NSFileProviderErrorCode) String() string {
	switch e {
	case NSFileProviderErrorApplicationExtensionNotFound:
		return "NSFileProviderErrorApplicationExtensionNotFound"
	case NSFileProviderErrorCannotSynchronize:
		return "NSFileProviderErrorCannotSynchronize"
	case NSFileProviderErrorDeletionRejected:
		return "NSFileProviderErrorDeletionRejected"
	case NSFileProviderErrorDirectoryNotEmpty:
		return "NSFileProviderErrorDirectoryNotEmpty"
	case NSFileProviderErrorDomainDisabled:
		return "NSFileProviderErrorDomainDisabled"
	case NSFileProviderErrorExcludedFromSync:
		return "NSFileProviderErrorExcludedFromSync"
	case NSFileProviderErrorFilenameCollision:
		return "NSFileProviderErrorFilenameCollision"
	case NSFileProviderErrorInsufficientQuota:
		return "NSFileProviderErrorInsufficientQuota"
	case NSFileProviderErrorLocalVersionConflictingWithServer:
		return "NSFileProviderErrorLocalVersionConflictingWithServer"
	case NSFileProviderErrorNewerExtensionVersionFound:
		return "NSFileProviderErrorNewerExtensionVersionFound"
	case NSFileProviderErrorNoSuchItem:
		return "NSFileProviderErrorNoSuchItem"
	case NSFileProviderErrorNonEvictable:
		return "NSFileProviderErrorNonEvictable"
	case NSFileProviderErrorNonEvictableChildren:
		return "NSFileProviderErrorNonEvictableChildren"
	case NSFileProviderErrorNotAuthenticated:
		return "NSFileProviderErrorNotAuthenticated"
	case NSFileProviderErrorOlderExtensionVersionRunning:
		return "NSFileProviderErrorOlderExtensionVersionRunning"
	case NSFileProviderErrorPageExpired:
		return "NSFileProviderErrorPageExpired"
	case NSFileProviderErrorProviderDomainNotFound:
		return "NSFileProviderErrorProviderDomainNotFound"
	case NSFileProviderErrorProviderDomainTemporarilyUnavailable:
		return "NSFileProviderErrorProviderDomainTemporarilyUnavailable"
	case NSFileProviderErrorProviderNotFound:
		return "NSFileProviderErrorProviderNotFound"
	case NSFileProviderErrorProviderTranslocated:
		return "NSFileProviderErrorProviderTranslocated"
	case NSFileProviderErrorServerUnreachable:
		return "NSFileProviderErrorServerUnreachable"
	case NSFileProviderErrorUnsyncedEdits:
		return "NSFileProviderErrorUnsyncedEdits"
	case NSFileProviderErrorVersionNoLongerAvailable:
		return "NSFileProviderErrorVersionNoLongerAvailable"
	default:
		return fmt.Sprintf("NSFileProviderErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderFetchContentsOptions
type NSFileProviderFetchContentsOptions uint

const (
	// NSFileProviderFetchContentsOptionsStrictVersioning: An option that indicates the system requires an exact match of the requested item’s version.
	NSFileProviderFetchContentsOptionsStrictVersioning NSFileProviderFetchContentsOptions = 1
)

func (e NSFileProviderFetchContentsOptions) String() string {
	switch e {
	case NSFileProviderFetchContentsOptionsStrictVersioning:
		return "NSFileProviderFetchContentsOptionsStrictVersioning"
	default:
		return fmt.Sprintf("NSFileProviderFetchContentsOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags
type NSFileProviderFileSystemFlags uint

const (
	// NSFileProviderFileSystemHidden: By default, the system hides the item when the user views the file system.
	NSFileProviderFileSystemHidden NSFileProviderFileSystemFlags = 8
	// NSFileProviderFileSystemPathExtensionHidden: By default, the system hides the item’s extension when showing its filename.
	NSFileProviderFileSystemPathExtensionHidden NSFileProviderFileSystemFlags = 16
	// NSFileProviderFileSystemUserExecutable: The user can execute the item.
	NSFileProviderFileSystemUserExecutable NSFileProviderFileSystemFlags = 1
	// NSFileProviderFileSystemUserReadable: The user can read the item.
	NSFileProviderFileSystemUserReadable NSFileProviderFileSystemFlags = 2
	// NSFileProviderFileSystemUserWritable: The user can modify the item.
	NSFileProviderFileSystemUserWritable NSFileProviderFileSystemFlags = 4
)

func (e NSFileProviderFileSystemFlags) String() string {
	switch e {
	case NSFileProviderFileSystemHidden:
		return "NSFileProviderFileSystemHidden"
	case NSFileProviderFileSystemPathExtensionHidden:
		return "NSFileProviderFileSystemPathExtensionHidden"
	case NSFileProviderFileSystemUserExecutable:
		return "NSFileProviderFileSystemUserExecutable"
	case NSFileProviderFileSystemUserReadable:
		return "NSFileProviderFileSystemUserReadable"
	case NSFileProviderFileSystemUserWritable:
		return "NSFileProviderFileSystemUserWritable"
	default:
		return fmt.Sprintf("NSFileProviderFileSystemFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities
type NSFileProviderItemCapabilities uint

const (
	// NSFileProviderItemCapabilitiesAllowsAddingSubItems: A value indicating that the user can add subitems.
	NSFileProviderItemCapabilitiesAllowsAddingSubItems NSFileProviderItemCapabilities = 2
	// NSFileProviderItemCapabilitiesAllowsContentEnumerating: A value indicating that the item’s contents can be enumerated.
	NSFileProviderItemCapabilitiesAllowsContentEnumerating NSFileProviderItemCapabilities = 1
	// NSFileProviderItemCapabilitiesAllowsDeleting: A value indicating that the item can be deleted.
	NSFileProviderItemCapabilitiesAllowsDeleting NSFileProviderItemCapabilities = 32
	// NSFileProviderItemCapabilitiesAllowsExcludingFromSync: A value indicating that the user can exclude the item from sync operations.
	NSFileProviderItemCapabilitiesAllowsExcludingFromSync NSFileProviderItemCapabilities = 128
	// NSFileProviderItemCapabilitiesAllowsReading: A value indicating that the value can be read from.
	NSFileProviderItemCapabilitiesAllowsReading NSFileProviderItemCapabilities = 1
	// NSFileProviderItemCapabilitiesAllowsRenaming: A value indicating that the item can be renamed.
	NSFileProviderItemCapabilitiesAllowsRenaming NSFileProviderItemCapabilities = 8
	// NSFileProviderItemCapabilitiesAllowsReparenting: A value indicating that the item can be moved.
	NSFileProviderItemCapabilitiesAllowsReparenting NSFileProviderItemCapabilities = 4
	// NSFileProviderItemCapabilitiesAllowsTrashing: A value indicating that the item can be moved to the trash.
	NSFileProviderItemCapabilitiesAllowsTrashing NSFileProviderItemCapabilities = 16
	// NSFileProviderItemCapabilitiesAllowsWriting: A value indicating that the item can be written to.
	NSFileProviderItemCapabilitiesAllowsWriting NSFileProviderItemCapabilities = 2
	// Deprecated.
	NSFileProviderItemCapabilitiesAllowsAll NSFileProviderItemCapabilities = 63
	// Deprecated.
	NSFileProviderItemCapabilitiesAllowsEvicting NSFileProviderItemCapabilities = 64
)

func (e NSFileProviderItemCapabilities) String() string {
	switch e {
	case NSFileProviderItemCapabilitiesAllowsAddingSubItems:
		return "NSFileProviderItemCapabilitiesAllowsAddingSubItems"
	case NSFileProviderItemCapabilitiesAllowsContentEnumerating:
		return "NSFileProviderItemCapabilitiesAllowsContentEnumerating"
	case NSFileProviderItemCapabilitiesAllowsDeleting:
		return "NSFileProviderItemCapabilitiesAllowsDeleting"
	case NSFileProviderItemCapabilitiesAllowsExcludingFromSync:
		return "NSFileProviderItemCapabilitiesAllowsExcludingFromSync"
	case NSFileProviderItemCapabilitiesAllowsRenaming:
		return "NSFileProviderItemCapabilitiesAllowsRenaming"
	case NSFileProviderItemCapabilitiesAllowsReparenting:
		return "NSFileProviderItemCapabilitiesAllowsReparenting"
	case NSFileProviderItemCapabilitiesAllowsTrashing:
		return "NSFileProviderItemCapabilitiesAllowsTrashing"
	case NSFileProviderItemCapabilitiesAllowsAll:
		return "NSFileProviderItemCapabilitiesAllowsAll"
	case NSFileProviderItemCapabilitiesAllowsEvicting:
		return "NSFileProviderItemCapabilitiesAllowsEvicting"
	default:
		return fmt.Sprintf("NSFileProviderItemCapabilities(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields
type NSFileProviderItemFields uint

const (
	// NSFileProviderItemContentModificationDate: The item’s modification date.
	NSFileProviderItemContentModificationDate NSFileProviderItemFields = 128
	// NSFileProviderItemContents: The item’s content.
	NSFileProviderItemContents NSFileProviderItemFields = 1
	// NSFileProviderItemCreationDate: The item’s creation date.
	NSFileProviderItemCreationDate NSFileProviderItemFields = 64
	// NSFileProviderItemExtendedAttributes: The item’s extended attributes.
	NSFileProviderItemExtendedAttributes NSFileProviderItemFields = 512
	// NSFileProviderItemFavoriteRank: The item’s favorite rank.
	NSFileProviderItemFavoriteRank NSFileProviderItemFields = 32
	// NSFileProviderItemFileSystemFlags: The flags describing the item’s on-disk representation.
	NSFileProviderItemFileSystemFlags NSFileProviderItemFields = 256
	// NSFileProviderItemFilename: The item’s filename.
	NSFileProviderItemFilename NSFileProviderItemFields = 2
	// NSFileProviderItemLastUsedDate: The date the item was last used.
	NSFileProviderItemLastUsedDate NSFileProviderItemFields = 8
	// NSFileProviderItemParentItemIdentifier: The identity of the directory that contains the item.
	NSFileProviderItemParentItemIdentifier NSFileProviderItemFields = 4
	// NSFileProviderItemTagData: The tags for the item.
	NSFileProviderItemTagData NSFileProviderItemFields = 16
	// NSFileProviderItemTypeAndCreator: The file type and creator codes for the item.
	NSFileProviderItemTypeAndCreator NSFileProviderItemFields = 1024
)

func (e NSFileProviderItemFields) String() string {
	switch e {
	case NSFileProviderItemContentModificationDate:
		return "NSFileProviderItemContentModificationDate"
	case NSFileProviderItemContents:
		return "NSFileProviderItemContents"
	case NSFileProviderItemCreationDate:
		return "NSFileProviderItemCreationDate"
	case NSFileProviderItemExtendedAttributes:
		return "NSFileProviderItemExtendedAttributes"
	case NSFileProviderItemFavoriteRank:
		return "NSFileProviderItemFavoriteRank"
	case NSFileProviderItemFileSystemFlags:
		return "NSFileProviderItemFileSystemFlags"
	case NSFileProviderItemFilename:
		return "NSFileProviderItemFilename"
	case NSFileProviderItemLastUsedDate:
		return "NSFileProviderItemLastUsedDate"
	case NSFileProviderItemParentItemIdentifier:
		return "NSFileProviderItemParentItemIdentifier"
	case NSFileProviderItemTagData:
		return "NSFileProviderItemTagData"
	case NSFileProviderItemTypeAndCreator:
		return "NSFileProviderItemTypeAndCreator"
	default:
		return fmt.Sprintf("NSFileProviderItemFields(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders
type NSFileProviderKnownFolders uint

const (
	NSFileProviderDesktop   NSFileProviderKnownFolders = 1
	NSFileProviderDocuments NSFileProviderKnownFolders = 2
)

func (e NSFileProviderKnownFolders) String() string {
	switch e {
	case NSFileProviderDesktop:
		return "NSFileProviderDesktop"
	case NSFileProviderDocuments:
		return "NSFileProviderDocuments"
	default:
		return fmt.Sprintf("NSFileProviderKnownFolders(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DisconnectionOptions
type NSFileProviderManagerDisconnectionOptions uint

const (
	// NSFileProviderManagerDisconnectionOptionsTemporary: A temporary disconnection.
	NSFileProviderManagerDisconnectionOptionsTemporary NSFileProviderManagerDisconnectionOptions = 1
)

func (e NSFileProviderManagerDisconnectionOptions) String() string {
	switch e {
	case NSFileProviderManagerDisconnectionOptionsTemporary:
		return "NSFileProviderManagerDisconnectionOptionsTemporary"
	default:
		return fmt.Sprintf("NSFileProviderManagerDisconnectionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderMaterializationFlags
type NSFileProviderMaterializationFlags uint

const (
	// NSFileProviderMaterializationFlagsKnownSparseRanges: A flag indicating that the system should consider the file fully materialized, even if it’s a sparse file.
	NSFileProviderMaterializationFlagsKnownSparseRanges NSFileProviderMaterializationFlags = 1
)

func (e NSFileProviderMaterializationFlags) String() string {
	switch e {
	case NSFileProviderMaterializationFlagsKnownSparseRanges:
		return "NSFileProviderMaterializationFlagsKnownSparseRanges"
	default:
		return fmt.Sprintf("NSFileProviderMaterializationFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions
type NSFileProviderModifyItemOptions uint

const (
	// NSFileProviderModifyItemFailOnConflict: An option to fail an upload in the event of a version conflict.
	NSFileProviderModifyItemFailOnConflict NSFileProviderModifyItemOptions = 2
	// NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication: An option to require the upload to complete before calling the completion handler.
	NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication NSFileProviderModifyItemOptions = 4
	// NSFileProviderModifyItemMayAlreadyExist: An option that indicates the changes may already exist in your remote storage.
	NSFileProviderModifyItemMayAlreadyExist NSFileProviderModifyItemOptions = 1
)

func (e NSFileProviderModifyItemOptions) String() string {
	switch e {
	case NSFileProviderModifyItemFailOnConflict:
		return "NSFileProviderModifyItemFailOnConflict"
	case NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication:
		return "NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication"
	case NSFileProviderModifyItemMayAlreadyExist:
		return "NSFileProviderModifyItemMayAlreadyExist"
	default:
		return fmt.Sprintf("NSFileProviderModifyItemOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide
type NSFileProviderTestingOperationSide uint

const (
	// NSFileProviderTestingOperationSideDisk: The File Provider extension’s local storage.
	NSFileProviderTestingOperationSideDisk NSFileProviderTestingOperationSide = 0
	// NSFileProviderTestingOperationSideFileProvider: The File Provider extension’s remote storage.
	NSFileProviderTestingOperationSideFileProvider NSFileProviderTestingOperationSide = 1
)

func (e NSFileProviderTestingOperationSide) String() string {
	switch e {
	case NSFileProviderTestingOperationSideDisk:
		return "NSFileProviderTestingOperationSideDisk"
	case NSFileProviderTestingOperationSideFileProvider:
		return "NSFileProviderTestingOperationSideFileProvider"
	default:
		return fmt.Sprintf("NSFileProviderTestingOperationSide(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType
type NSFileProviderTestingOperationType int

const (
	// NSFileProviderTestingOperationTypeChildrenEnumeration: Lists an item’s content.
	NSFileProviderTestingOperationTypeChildrenEnumeration NSFileProviderTestingOperationType = 6
	// NSFileProviderTestingOperationTypeCollisionResolution: Resolves a collision by renaming the new item.
	NSFileProviderTestingOperationTypeCollisionResolution NSFileProviderTestingOperationType = 7
	// NSFileProviderTestingOperationTypeContentFetch: Fetches an item’s content.
	NSFileProviderTestingOperationTypeContentFetch NSFileProviderTestingOperationType = 5
	// NSFileProviderTestingOperationTypeCreation: Propagates the creation of a source item to the target location.
	NSFileProviderTestingOperationTypeCreation NSFileProviderTestingOperationType = 2
	// NSFileProviderTestingOperationTypeDeletion: Propagates the deletion of the source item from the target location.
	NSFileProviderTestingOperationTypeDeletion NSFileProviderTestingOperationType = 4
	// NSFileProviderTestingOperationTypeIngestion: Alerts the system to changes to either the local or remote storage.
	NSFileProviderTestingOperationTypeIngestion NSFileProviderTestingOperationType = 0
	// NSFileProviderTestingOperationTypeLookup: Looks up an item.
	NSFileProviderTestingOperationTypeLookup NSFileProviderTestingOperationType = 1
	// NSFileProviderTestingOperationTypeModification: Propagates a change from the source item to the target location.
	NSFileProviderTestingOperationTypeModification NSFileProviderTestingOperationType = 3
)

func (e NSFileProviderTestingOperationType) String() string {
	switch e {
	case NSFileProviderTestingOperationTypeChildrenEnumeration:
		return "NSFileProviderTestingOperationTypeChildrenEnumeration"
	case NSFileProviderTestingOperationTypeCollisionResolution:
		return "NSFileProviderTestingOperationTypeCollisionResolution"
	case NSFileProviderTestingOperationTypeContentFetch:
		return "NSFileProviderTestingOperationTypeContentFetch"
	case NSFileProviderTestingOperationTypeCreation:
		return "NSFileProviderTestingOperationTypeCreation"
	case NSFileProviderTestingOperationTypeDeletion:
		return "NSFileProviderTestingOperationTypeDeletion"
	case NSFileProviderTestingOperationTypeIngestion:
		return "NSFileProviderTestingOperationTypeIngestion"
	case NSFileProviderTestingOperationTypeLookup:
		return "NSFileProviderTestingOperationTypeLookup"
	case NSFileProviderTestingOperationTypeModification:
		return "NSFileProviderTestingOperationTypeModification"
	default:
		return fmt.Sprintf("NSFileProviderTestingOperationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason
type NSFileProviderVolumeUnsupportedReason uint

const (
	NSFileProviderVolumeUnsupportedReasonNetwork      NSFileProviderVolumeUnsupportedReason = 16
	NSFileProviderVolumeUnsupportedReasonNonAPFS      NSFileProviderVolumeUnsupportedReason = 2
	NSFileProviderVolumeUnsupportedReasonNonEncrypted NSFileProviderVolumeUnsupportedReason = 4
	NSFileProviderVolumeUnsupportedReasonNone         NSFileProviderVolumeUnsupportedReason = 0
	NSFileProviderVolumeUnsupportedReasonQuarantined  NSFileProviderVolumeUnsupportedReason = 32
	NSFileProviderVolumeUnsupportedReasonReadOnly     NSFileProviderVolumeUnsupportedReason = 8
	NSFileProviderVolumeUnsupportedReasonUnknown      NSFileProviderVolumeUnsupportedReason = 1
)

func (e NSFileProviderVolumeUnsupportedReason) String() string {
	switch e {
	case NSFileProviderVolumeUnsupportedReasonNetwork:
		return "NSFileProviderVolumeUnsupportedReasonNetwork"
	case NSFileProviderVolumeUnsupportedReasonNonAPFS:
		return "NSFileProviderVolumeUnsupportedReasonNonAPFS"
	case NSFileProviderVolumeUnsupportedReasonNonEncrypted:
		return "NSFileProviderVolumeUnsupportedReasonNonEncrypted"
	case NSFileProviderVolumeUnsupportedReasonNone:
		return "NSFileProviderVolumeUnsupportedReasonNone"
	case NSFileProviderVolumeUnsupportedReasonQuarantined:
		return "NSFileProviderVolumeUnsupportedReasonQuarantined"
	case NSFileProviderVolumeUnsupportedReasonReadOnly:
		return "NSFileProviderVolumeUnsupportedReasonReadOnly"
	case NSFileProviderVolumeUnsupportedReasonUnknown:
		return "NSFileProviderVolumeUnsupportedReasonUnknown"
	default:
		return fmt.Sprintf("NSFileProviderVolumeUnsupportedReason(%d)", e)
	}
}
