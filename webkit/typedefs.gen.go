// Code generated from Apple documentation. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/foundation"
)

// See: https://developer.apple.com/documentation/WebKit/DOMTimeStamp
type DOMTimeStamp = uint64

// See: https://developer.apple.com/documentation/WebKit/NSAttributedStringCompletionHandler
type NSAttributedStringCompletionHandler = func(foundation.NSAttributedString, foundation.INSDictionary, foundation.NSError)

// WKWebExtensionContextNotificationUserInfoKey is constants for specifying web extension context information in notifications.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/NotificationUserInfoKey
type WKWebExtensionContextNotificationUserInfoKey = string

// WKWebExtensionDataType is constants for specifying data types for a [WKWebExtension.DataRecord].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataType
type WKWebExtensionDataType = string

// WKWebExtensionPermission is constants for specifying permission in a [WKWebExtensionContext].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Permission
type WKWebExtensionPermission = string
