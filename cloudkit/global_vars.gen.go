// Code generated from Apple documentation. DO NOT EDIT.

package cloudkit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// CKAccountChangedNotification is a notification that a container posts when the status of an iCloud account changes.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKAccountChangedNotification
	CKAccountChangedNotification string
	// CKCurrentUserDefaultName is a constant that provides the current user’s default name.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKCurrentUserDefaultName
	CKCurrentUserDefaultName string
	// CKErrorDomain is the error domain for CloudKit errors.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKErrorDomain
	CKErrorDomain string
	// CKErrorRetryAfterKey is the key to retrieve the number of seconds to wait before you retry a request.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKErrorRetryAfterKey
	CKErrorRetryAfterKey string
	// CKErrorUserDidResetEncryptedDataKey is the key that determines whether CloudKit deletes a record zone because of a user action.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKErrorUserDidResetEncryptedDataKey
	CKErrorUserDidResetEncryptedDataKey string
	// CKPartialErrorsByItemIDKey is the key to retrieve partial errors.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKPartialErrorsByItemIDKey
	CKPartialErrorsByItemIDKey string
	// CKRecordChangedErrorAncestorRecordKey is the key to retrieve the original version of the record.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordChangedErrorAncestorRecordKey
	CKRecordChangedErrorAncestorRecordKey string
	// CKRecordChangedErrorClientRecordKey is the key to retrieve the local version of the record.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordChangedErrorClientRecordKey
	CKRecordChangedErrorClientRecordKey string
	// CKRecordChangedErrorServerRecordKey is the key to retrieve the server’s version of the record.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordChangedErrorServerRecordKey
	CKRecordChangedErrorServerRecordKey string
	// CKRecordNameZoneWideShare is the name of a share record that manages a shared record zone.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordNameZoneWideShare
	CKRecordNameZoneWideShare string
	// CKRecordZoneDefaultName is the default record zone’s name.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneDefaultName-8mfij
	CKRecordZoneDefaultName string
)

var (
	// CKQueryOperationMaximumResults is a constant value that represents the maximum number of results CloudKit retrieves.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/maximumResults
	CKQueryOperationMaximumResults uint
)

var (
	// CKRecordCreationDateKey is a key constant that a record uses for its CKRecord.creationDate.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordCreationDateKey
	CKRecordCreationDateKey CKRecordFieldKey
	// CKRecordCreatorUserRecordIDKey is a key constant that a record uses for its CKRecord.creatorUserRecordID.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordCreatorUserRecordIDKey
	CKRecordCreatorUserRecordIDKey CKRecordFieldKey
	// CKRecordLastModifiedUserRecordIDKey is a key constant that a record uses for its CKRecord.lastModifiedUserRecordID.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordLastModifiedUserRecordIDKey
	CKRecordLastModifiedUserRecordIDKey CKRecordFieldKey
	// CKRecordModificationDateKey is a key constant that a record uses for its CKRecord.modificationDate.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordModificationDateKey
	CKRecordModificationDateKey CKRecordFieldKey
	// CKRecordParentKey is the key constant that a record uses for its CKRecord.parent.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordParentKey-2kx8l
	CKRecordParentKey CKRecordFieldKey
	// CKRecordRecordIDKey is a key constant that a record uses for its CKRecord.recordID.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordRecordIDKey
	CKRecordRecordIDKey CKRecordFieldKey
	// CKRecordShareKey is the key constant that a record uses for its CKRecord.share.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordShareKey-rrat
	CKRecordShareKey CKRecordFieldKey
	// CKShareThumbnailImageDataKey is the system field key for the share’s thumbnail image data.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKShareThumbnailImageDataKey-1rxdx
	CKShareThumbnailImageDataKey CKRecordFieldKey
	// CKShareTitleKey is the system field key for the share’s title.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKShareTitleKey-9yavd
	CKShareTitleKey CKRecordFieldKey
	// CKShareTypeKey is the system field key for the share’s type.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKShareTypeKey-204gl
	CKShareTypeKey CKRecordFieldKey
)

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKAccountChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKAccountChangedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKCurrentUserDefaultName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKCurrentUserDefaultName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKErrorRetryAfterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKErrorRetryAfterKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKErrorUserDidResetEncryptedDataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKErrorUserDidResetEncryptedDataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKPartialErrorsByItemIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKPartialErrorsByItemIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKQueryOperationMaximumResults"); err == nil && ptr != 0 {
		CKQueryOperationMaximumResults = objc.ValueAt[uint](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordChangedErrorAncestorRecordKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordChangedErrorAncestorRecordKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordChangedErrorClientRecordKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordChangedErrorClientRecordKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordChangedErrorServerRecordKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordChangedErrorServerRecordKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordCreationDateKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordCreatorUserRecordIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordCreatorUserRecordIDKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordLastModifiedUserRecordIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordLastModifiedUserRecordIDKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordModificationDateKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordNameZoneWideShare"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordNameZoneWideShare = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordParentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordParentKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordRecordIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordRecordIDKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordShareKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordShareKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordTypeShare"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordTypes.Share = CKRecordType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordTypeUserRecord"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordTypes.UserRecord = CKRecordType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKRecordZoneDefaultName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKRecordZoneDefaultName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKShareThumbnailImageDataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKShareThumbnailImageDataKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKShareTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKShareTitleKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CKShareTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CKShareTypeKey = CKRecordFieldKey(objc.GoString(cstr))
			}
		}
	}

}

// CKRecordTypes provides typed accessors for [CKRecordType] constants.
var CKRecordTypes struct {
	// Share: The system type that identifies a share record.
	Share CKRecordType
	// UserRecord: The system type that identifies a user record.
	UserRecord CKRecordType
}
