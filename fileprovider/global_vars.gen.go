// Code generated from Apple documentation. DO NOT EDIT.

package fileprovider

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// FileProviderDomainDidChange is a notification that posts when a file provider’s domain changes.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainDidChange
	FileProviderDomainDidChange foundation.NSNotificationName
	// FileProviderMaterializedSetDidChange is a notification that the system posts when the set of materialized items changes for your file provider extension.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderMaterializedSetDidChange
	FileProviderMaterializedSetDidChange foundation.NSNotificationName
	// FileProviderPendingSetDidChange is a notification that the system posts when the set of pending items changes for your file provider extension.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetDidChange
	FileProviderPendingSetDidChange foundation.NSNotificationName
)

var (
	// FileProviderErrorDomain is the error domain for the File Provider extension.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderErrorDomain
	FileProviderErrorDomain foundation.NSErrorDomain
)

var (
	// FileProviderErrorItemKey is the key for accessing information about sync-related errors.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderErrorItemKey
	FileProviderErrorItemKey foundation.NSErrorUserInfoKey
	// FileProviderErrorNonExistentItemIdentifierKey is the key for accessing the specified item’s identifier when the item doesn’t exist.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderErrorNonExistentItemIdentifierKey
	FileProviderErrorNonExistentItemIdentifierKey foundation.NSErrorUserInfoKey
)

var (
	// FileProviderFavoriteRankUnranked is a value that indicates that the item is not ranked.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderFavoriteRankUnranked
	FileProviderFavoriteRankUnranked uint64
)

var (
	// FileProviderInitialPageSortedByDate is the initial batch of items when sorted by date.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPage/initialPageSortedByDate
	FileProviderInitialPageSortedByDate NSFileProviderPage
	// FileProviderInitialPageSortedByName is the initial batch of items when sorted by name.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPage/initialPageSortedByName
	FileProviderInitialPageSortedByName NSFileProviderPage
)

var (
	// FileProviderRootContainerItemIdentifier is the persistent identifier for the root directory of the file provider’s shared file hierarchy.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/rootContainer
	FileProviderRootContainerItemIdentifier NSFileProviderItemIdentifier
	// FileProviderTrashContainerItemIdentifier is the persistent identifier for the parent of all trashed items.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/trashContainer
	FileProviderTrashContainerItemIdentifier NSFileProviderItemIdentifier
	// FileProviderWorkingSetContainerItemIdentifier is the persistent identifier representing the working set of documents and directories.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/workingSet
	FileProviderWorkingSetContainerItemIdentifier NSFileProviderItemIdentifier
)

var (
	// FileProviderUserInfoExperimentIDKey is system interpreted user info key When setting a value to that user info on a domain, the system will ingest this value. If user has given their consent for telemetry, this value will be used to decorate telemetry messages sent by the FileProvider subsystem. The telemetry messages can be then later on retrieved by developers along with the other metrics through the CloudKit console as detailed here: https://developer.apple.com/documentation/fileprovider/exporting-file-provider-metrics-data?language=objc This will help developers triaging data they receive from testing population compared to regular users The value must either be a NSNumber between [0 - 31]. If it’s not in that range, or if it is not a NSNumber, any call to addDomain with that invalid UserInfo dictionary will fail with a EINVAL POSIX NSError. To update this value, the provider must call addDomain with an updated userInfo dictionary.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInfoKey/experimentID
	FileProviderUserInfoExperimentIDKey NSFileProviderUserInfoKey
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderDomainDidChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderDomainDidChange = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderErrorItemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderErrorItemKey = foundation.NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderErrorNonExistentItemIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderErrorNonExistentItemIdentifierKey = foundation.NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderFavoriteRankUnranked"); err == nil && ptr != 0 {
		FileProviderFavoriteRankUnranked = objc.ValueAt[uint64](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderInitialPageSortedByDate"); err == nil && ptr != 0 {
		FileProviderInitialPageSortedByDate = objc.ValueAt[NSFileProviderPage](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderInitialPageSortedByName"); err == nil && ptr != 0 {
		FileProviderInitialPageSortedByName = objc.ValueAt[NSFileProviderPage](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderMaterializedSetDidChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderMaterializedSetDidChange = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderPendingSetDidChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderPendingSetDidChange = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderRootContainerItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderRootContainerItemIdentifier = NSFileProviderItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderTrashContainerItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderTrashContainerItemIdentifier = NSFileProviderItemIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderUserInfoExperimentIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderUserInfoExperimentIDKey = NSFileProviderUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProviderWorkingSetContainerItemIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProviderWorkingSetContainerItemIdentifier = NSFileProviderItemIdentifier(objc.GoString(cstr))
			}
		}
	}

}
