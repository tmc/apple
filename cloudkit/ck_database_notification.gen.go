// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKDatabaseNotification] class.
var (
	_CKDatabaseNotificationClass     CKDatabaseNotificationClass
	_CKDatabaseNotificationClassOnce sync.Once
)

func getCKDatabaseNotificationClass() CKDatabaseNotificationClass {
	_CKDatabaseNotificationClassOnce.Do(func() {
		_CKDatabaseNotificationClass = CKDatabaseNotificationClass{class: objc.GetClass("CKDatabaseNotification")}
	})
	return _CKDatabaseNotificationClass
}

// GetCKDatabaseNotificationClass returns the class object for CKDatabaseNotification.
func GetCKDatabaseNotificationClass() CKDatabaseNotificationClass {
	return getCKDatabaseNotificationClass()
}

type CKDatabaseNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDatabaseNotificationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDatabaseNotificationClass) Alloc() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A notification that triggers when the contents of a database change.
//
// # Overview
//
// Database subscriptions execute when changes happen in any of a database’s
// record zones, for example, when CloudKit saves a new record. When the
// subscription registers a change, it sends push notifications to the
// user’s devices to inform your app about the change. You can then fetch
// the changes and cache them on-device. When appropriate, CloudKit excludes
// the device where the change originates.
//
// You configure a subscription’s notifications by setting it’s
// [CKDatabaseNotification.NotificationInfo] property. Do this before you save it to the server. A
// subscription generates either high-priority or medium-priority push
// notifications. CloudKit delivers medium-priority notifications to your app
// in the background. High-priority notifications are visual and the system
// displays them to the user. Visual notifications need the user’s
// permission. For more information, see [Asking permission to use
// notifications].
//
// A subscription uses [CKNotificationInfo] to configure its notifications.
// For background delivery, set only its [CKDatabaseNotification.ShouldSendContentAvailable] property
// to true. If you set any other property, CloudKit treats the notification as
// high-priority.
//
// Don’t rely on push notifications for specific changes because the system
// can coalesce them. CloudKit can omit data to keep the notification’s
// payload size under the APNs size limit. Consider notifications an
// indication of remote changes. Use [CKDatabaseNotification.DatabaseScope] to determine which
// database has changes, and then [CKFetchDatabaseChangesOperation] to fetch
// those changes. A notification’s [CKDatabaseNotification.IsPruned] property is true if CloudKit
// omits data.
//
// You don’t instantiate this class. Instead, implement
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)] in
// your app delegate. Initialize [CKNotification] with the `userInfo`
// dictionary that CloudKit passes to the method. This returns an instance of
// the appropriate subclass. Use the [CKDatabaseNotification.NotificationType] property to determine
// the type. Then cast to that type to access type-specific properties and
// methods.
//
// # Getting the Database Scope
//
//   - [CKDatabaseNotification.DatabaseScope]: The type of database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseNotification
//
// [Asking permission to use notifications]: https://developer.apple.com/documentation/UserNotifications/asking-permission-to-use-notifications
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
type CKDatabaseNotification struct {
	CKNotification
}

// CKDatabaseNotificationFromID constructs a [CKDatabaseNotification] from an objc.ID.
//
// A notification that triggers when the contents of a database change.
func CKDatabaseNotificationFromID(id objc.ID) CKDatabaseNotification {
	return CKDatabaseNotification{CKNotification: CKNotificationFromID(id)}
}

// NOTE: CKDatabaseNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDatabaseNotification] class.
//
// # Getting the Database Scope
//
//   - [ICKDatabaseNotification.DatabaseScope]: The type of database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseNotification
type ICKDatabaseNotification interface {
	ICKNotification

	// Topic: Getting the Database Scope

	// The type of database.
	DatabaseScope() CKDatabaseScope

	// The configuration for a subscription’s push notifications.
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
	// A Boolean value that indicates whether the push notification includes the content available flag.
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
}

// Init initializes the instance.
func (c CKDatabaseNotification) Init() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDatabaseNotification) Autorelease() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseNotification creates a new CKDatabaseNotification instance.
func NewCKDatabaseNotification() CKDatabaseNotification {
	class := getCKDatabaseNotificationClass()
	rv := objc.Send[CKDatabaseNotification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new notification using the specified payload data.
//
// notificationDictionary: The push notification’s payload data. Use the dictionary that the system
// provides to your app delegate’s
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]
// method. This parameter must not be `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
//
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
func NewCKDatabaseNotificationFromRemoteNotificationDictionary(notificationDictionary foundation.INSDictionary) CKDatabaseNotification {
	rv := objc.Send[objc.ID](objc.ID(getCKDatabaseNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return CKDatabaseNotificationFromID(rv)
}

// The type of database.
//
// # Discussion
//
// This property’s value is one of the constants that [CKDatabase.Scope]
// defines.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseNotification/databaseScope
//
// [CKDatabase.Scope]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
func (c CKDatabaseNotification) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c.ID, objc.Sel("databaseScope"))
	return CKDatabaseScope(rv)
}

// The configuration for a subscription’s push notifications.
//
// See: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c CKDatabaseNotification) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("notificationInfo"))
	return CKNotificationInfoFromID(objc.ID(rv))
}
func (c CKDatabaseNotification) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[struct{}](c.ID, objc.Sel("setNotificationInfo:"), value)
}

// A Boolean value that indicates whether the push notification includes the
// content available flag.
//
// See: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c CKDatabaseNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}
func (c CKDatabaseNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}
