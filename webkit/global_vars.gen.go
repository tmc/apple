// Code generated from Apple documentation. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/WebKit/NSReadAccessURLDocumentOption
	ReadAccessURLDocumentOption appkit.NSAttributedStringDocumentReadingOptionKey
	// WKWebExtensionContextDeniedPermissionMatchPatternsWereRemovedNotification is a notification the system sends whenever a web extension context has newly removed denied permission match patterns.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissionMatchPatternsWereRemovedNotification
	WKWebExtensionContextDeniedPermissionMatchPatternsWereRemovedNotification foundation.NSNotificationName
	// WKWebExtensionContextDeniedPermissionsWereRemovedNotification is a notification the system sends whenever a web extension context has newly removed denied permissions.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissionsWereRemovedNotification
	WKWebExtensionContextDeniedPermissionsWereRemovedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/errorDomain
	WKWebExtensionContextErrorDomain foundation.NSErrorDomain
	// WKWebExtensionContextErrorsDidUpdateNotification is this notification is sent whenever a web extension context has new errors or errors were cleared.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/errorsDidUpdateNotification
	WKWebExtensionContextErrorsDidUpdateNotification foundation.NSNotificationName
	// WKWebExtensionContextGrantedPermissionMatchPatternsWereRemovedNotification is this notification is sent whenever a web extension context has newly removed granted permission match patterns.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissionMatchPatternsWereRemovedNotification
	WKWebExtensionContextGrantedPermissionMatchPatternsWereRemovedNotification foundation.NSNotificationName
	// WKWebExtensionContextGrantedPermissionsWereRemovedNotification is this notification is sent whenever a web extension context has newly removed granted permissions.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissionsWereRemovedNotification
	WKWebExtensionContextGrantedPermissionsWereRemovedNotification foundation.NSNotificationName
	// WKWebExtensionContextPermissionMatchPatternsWereDeniedNotification is this notification is sent whenever a web extension context has newly denied permission match patterns.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionMatchPatternsWereDeniedNotification
	WKWebExtensionContextPermissionMatchPatternsWereDeniedNotification foundation.NSNotificationName
	// WKWebExtensionContextPermissionMatchPatternsWereGrantedNotification is this notification is sent whenever a web extension context has newly granted permission match patterns.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionMatchPatternsWereGrantedNotification
	WKWebExtensionContextPermissionMatchPatternsWereGrantedNotification foundation.NSNotificationName
	// WKWebExtensionContextPermissionsWereDeniedNotification is this notification is sent whenever a web extension context has newly denied permissions.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionsWereDeniedNotification
	WKWebExtensionContextPermissionsWereDeniedNotification foundation.NSNotificationName
	// WKWebExtensionContextPermissionsWereGrantedNotification is this notification is sent whenever a web extension context has newly granted permissions.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionsWereGrantedNotification
	WKWebExtensionContextPermissionsWereGrantedNotification foundation.NSNotificationName
	// WKWebExtensionDataRecordErrorDomain is indicates a [WKWebExtension.DataRecord] error.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/errorDomain
	WKWebExtensionDataRecordErrorDomain foundation.NSErrorDomain
	// WKWebExtensionErrorDomain is indicates a web extension error.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/errorDomain
	WKWebExtensionErrorDomain foundation.NSErrorDomain
	// WKWebExtensionMatchPatternErrorDomain is a string that identifies the error domain.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/errorDomain
	WKWebExtensionMatchPatternErrorDomain foundation.NSErrorDomain
	// WKWebExtensionMessagePortErrorDomain is a string that identifies the error domain.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/errorDomain
	WKWebExtensionMessagePortErrorDomain foundation.NSErrorDomain
)

var (
	// WKErrorDomain is string that identifies the WebKit error domain.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKErrorDomain
	WKErrorDomain string
	// WKWebsiteDataTypeCookies is cookies.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeCookies
	WKWebsiteDataTypeCookies string
	// WKWebsiteDataTypeDiskCache is on-disk caches.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeDiskCache
	WKWebsiteDataTypeDiskCache string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeFetchCache
	WKWebsiteDataTypeFetchCache string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeFileSystem
	WKWebsiteDataTypeFileSystem string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeHashSalt
	WKWebsiteDataTypeHashSalt string
	// WKWebsiteDataTypeIndexedDBDatabases is indexedDB databases.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeIndexedDBDatabases
	WKWebsiteDataTypeIndexedDBDatabases string
	// WKWebsiteDataTypeLocalStorage is hTML local storage.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeLocalStorage
	WKWebsiteDataTypeLocalStorage string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeMediaKeys
	WKWebsiteDataTypeMediaKeys string
	// WKWebsiteDataTypeMemoryCache is in-memory caches.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeMemoryCache
	WKWebsiteDataTypeMemoryCache string
	// WKWebsiteDataTypeOfflineWebApplicationCache is hTML offline web app caches.
	//
	// Deprecated: Deprecated since macOS 26.2. WebApplicationCache is no longer supported
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeOfflineWebApplicationCache
	WKWebsiteDataTypeOfflineWebApplicationCache string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeScreenTime
	WKWebsiteDataTypeScreenTime string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeSearchFieldRecentSearches
	WKWebsiteDataTypeSearchFieldRecentSearches string
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeServiceWorkerRegistrations
	WKWebsiteDataTypeServiceWorkerRegistrations string
	// WKWebsiteDataTypeSessionStorage is hTML session storage.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeSessionStorage
	WKWebsiteDataTypeSessionStorage string
	// WKWebsiteDataTypeWebSQLDatabases is webSQL databases.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataTypeWebSQLDatabases
	WKWebsiteDataTypeWebSQLDatabases string
)

var ()

var ()

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSReadAccessURLDocumentOption"); err == nil && ptr != 0 {
		ReadAccessURLDocumentOption = *(*appkit.NSAttributedStringDocumentReadingOptionKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextDeniedPermissionMatchPatternsWereRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextDeniedPermissionMatchPatternsWereRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextDeniedPermissionsWereRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextDeniedPermissionsWereRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextErrorsDidUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextErrorsDidUpdateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextGrantedPermissionMatchPatternsWereRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextGrantedPermissionMatchPatternsWereRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextGrantedPermissionsWereRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextGrantedPermissionsWereRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextNotificationUserInfoKeyMatchPatterns"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextNotificationUserInfoKeys.MatchPatterns = WKWebExtensionContextNotificationUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextNotificationUserInfoKeyPermissions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextNotificationUserInfoKeys.Permissions = WKWebExtensionContextNotificationUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextPermissionMatchPatternsWereDeniedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextPermissionMatchPatternsWereDeniedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextPermissionMatchPatternsWereGrantedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextPermissionMatchPatternsWereGrantedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextPermissionsWereDeniedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextPermissionsWereDeniedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionContextPermissionsWereGrantedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionContextPermissionsWereGrantedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionDataRecordErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionDataRecordErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionDataTypeLocal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionDataTypes.Local = WKWebExtensionDataType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionDataTypeSession"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionDataTypes.Session = WKWebExtensionDataType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionDataTypeSynchronized"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionDataTypes.Synchronized = WKWebExtensionDataType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionMatchPatternErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionMatchPatternErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionMessagePortErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionMessagePortErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionActiveTab"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.ActiveTab = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionAlarms"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Alarms = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionClipboardWrite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.ClipboardWrite = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionContextMenus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.ContextMenus = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionCookies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Cookies = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionDeclarativeNetRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.DeclarativeNetRequest = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionDeclarativeNetRequestFeedback"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.DeclarativeNetRequestFeedback = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionDeclarativeNetRequestWithHostAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.DeclarativeNetRequestWithHostAccess = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionMenus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Menus = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionNativeMessaging"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.NativeMessaging = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionScripting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Scripting = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Storage = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionTabs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.Tabs = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionUnlimitedStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.UnlimitedStorage = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionWebNavigation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.WebNavigation = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebExtensionPermissionWebRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebExtensionPermissions.WebRequest = WKWebExtensionPermission(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeCookies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeCookies = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeDiskCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeDiskCache = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeFetchCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeFetchCache = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeFileSystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeFileSystem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeHashSalt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeHashSalt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeIndexedDBDatabases"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeIndexedDBDatabases = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeLocalStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeLocalStorage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeMediaKeys"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeMediaKeys = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeMemoryCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeMemoryCache = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeOfflineWebApplicationCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeOfflineWebApplicationCache = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeScreenTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeScreenTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeSearchFieldRecentSearches"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeSearchFieldRecentSearches = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeServiceWorkerRegistrations"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeServiceWorkerRegistrations = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeSessionStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeSessionStorage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "WKWebsiteDataTypeWebSQLDatabases"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WKWebsiteDataTypeWebSQLDatabases = objc.GoString(cstr)
			}
		}
	}

}

// WKWebExtensionContextNotificationUserInfoKeys provides typed accessors for [WKWebExtensionContextNotificationUserInfoKey] constants.
var WKWebExtensionContextNotificationUserInfoKeys struct {
	// MatchPatterns: The corresponding value represents the affected permission match patterns in [WKWebExtensionContext](<doc://com.apple.webkit/documentation/WebKit/WKWebExtensionContext>) notifications.
	MatchPatterns WKWebExtensionContextNotificationUserInfoKey
	// Permissions: The corresponding value represents the affected permissions in [WKWebExtensionContext](<doc://com.apple.webkit/documentation/WebKit/WKWebExtensionContext>) notifications.
	Permissions WKWebExtensionContextNotificationUserInfoKey
}

// WKWebExtensionDataTypes provides typed accessors for [WKWebExtensionDataType] constants.
var WKWebExtensionDataTypes struct {
	// Local: Specifies local storage, including `browser.StorageXCUIElementTypeLocal()`.
	Local WKWebExtensionDataType
	// Session: Specifies session storage, including `browser.StorageXCUIElementTypeSession()`.
	Session WKWebExtensionDataType
	// Synchronized: Specifies synchronized storage, including `browser.StorageXCUIElementTypeSync()`.
	Synchronized WKWebExtensionDataType
}

// WKWebExtensionPermissions provides typed accessors for [WKWebExtensionPermission] constants.
var WKWebExtensionPermissions struct {
	// ActiveTab: A request indicating that when a person interacts with the extension, the system grants extra permissions for the active tab only.
	ActiveTab WKWebExtensionPermission
	// Alarms: A request for access to the `browser.Alarms()` APIs.
	Alarms WKWebExtensionPermission
	// ClipboardWrite: A request for access to write to the clipboard.
	ClipboardWrite WKWebExtensionPermission
	// ContextMenus: A request for access to the `browser.ContextMenus()` APIs.
	ContextMenus WKWebExtensionPermission
	// Cookies: A request for access to the `browser.Cookies()` APIs.
	Cookies WKWebExtensionPermission
	// DeclarativeNetRequest: A request for access to the `browser.DeclarativeNetRequest()` APIs.
	DeclarativeNetRequest WKWebExtensionPermission
	// DeclarativeNetRequestFeedback: A request for access to the `browser.DeclarativeNetRequest()` APIs with extra information on matched rules.
	DeclarativeNetRequestFeedback WKWebExtensionPermission
	// DeclarativeNetRequestWithHostAccess: A request for access to the `browser.DeclarativeNetRequest()` APIs with the ability to modify or redirect requests.
	DeclarativeNetRequestWithHostAccess WKWebExtensionPermission
	// Menus: A request for access to the `browser.Menus()` APIs.
	Menus WKWebExtensionPermission
	// NativeMessaging: A request for access to send messages to the app extension bundle.
	NativeMessaging WKWebExtensionPermission
	// Scripting: A request for access to the `browser.Scripting()` APIs.
	Scripting WKWebExtensionPermission
	// Storage: A request for access to the `browser.Storage()` APIs.
	Storage WKWebExtensionPermission
	// Tabs: A request for access to extra information on the `browser.Tabs()` APIs.
	Tabs WKWebExtensionPermission
	// UnlimitedStorage: A request for access to an unlimited quota on the `browser.StorageXCUIElementTypeLocal()` APIs.
	UnlimitedStorage WKWebExtensionPermission
	// WebNavigation: A request for access to the `browser.WebNavigation()` APIs.
	WebNavigation WKWebExtensionPermission
	// WebRequest: A request for access to the `browser.WebRequest()` APIs.
	WebRequest WKWebExtensionPermission
}
