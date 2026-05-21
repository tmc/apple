// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKRecordZoneNotification] class.
var (
	_CKRecordZoneNotificationClass     CKRecordZoneNotificationClass
	_CKRecordZoneNotificationClassOnce sync.Once
)

func getCKRecordZoneNotificationClass() CKRecordZoneNotificationClass {
	_CKRecordZoneNotificationClassOnce.Do(func() {
		_CKRecordZoneNotificationClass = CKRecordZoneNotificationClass{class: objc.GetClass("CKRecordZoneNotification")}
	})
	return _CKRecordZoneNotificationClass
}

// GetCKRecordZoneNotificationClass returns the class object for CKRecordZoneNotification.
func GetCKRecordZoneNotificationClass() CKRecordZoneNotificationClass {
	return getCKRecordZoneNotificationClass()
}

type CKRecordZoneNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordZoneNotificationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordZoneNotificationClass) Alloc() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A notification that triggers when the contents of a record zone change.
//
// # Overview
//
// A record zone subscription executes when a user, or in certain scenarios,
// CloudKit, modifies a record in that zone, for example, when a field’s
// value changes in a record. When CloudKit registers the change, it sends
// push notifications to the user’s devices to inform your app about the
// change. You can then fetch the changes and cache them on-device. When
// appropriate, CloudKit excludes the device where the change originates.
//
// You configure a subscription’s notifications by setting it’s
// [CKSubscription.NotificationInfo] property. Do this before you save it to
// the server. A subscription generates either high-priority or
// medium-priority push notifications. CloudKit delivers medium-priority
// notifications to your app in the background. High-priority notifications
// are visual and the system displays them to the user. Visual notifications
// need the user’s permission. For more information, see [Asking permission
// to use notifications].
//
// A subscription uses [CKNotificationInfo] to configure its notifications.
// For background delivery, set only its
// [CKNotificationInfo.ShouldSendContentAvailable] property to true. If you
// set any other property, CloudKit treats the notification as high-priority.
//
// Don’t rely on push notifications for specific changes to records because
// the system can coalesce them. CloudKit can omit data to keep the
// notification’s payload size under the APNs size limit. Consider
// notifications an indication of remote changes. Use
// [CKRecordZoneNotification.DatabaseScope] to determine which database
// contains the changed record zone, and
// [CKRecordZoneNotification.RecordZoneID] to determine which zone contains
// changed records. You can then fetch just those changes using
// [CKFetchRecordZoneChangesOperation]. A notification’s
// [CKNotification.IsPruned] property is true if CloudKit omits data.
//
// You don’t instantiate this class. Instead, implement
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)] in
// your app delegate. Initialize [CKNotification] with the `userInfo`
// dictionary that CloudKit passes to the method. This returns an instance of
// the appropriate subclass. Use the [CKNotification.NotificationType]
// property to determine the type. Then cast to that type to access
// type-specific properties and methods.
//
// # Getting the Record Zone ID
//
//   - [CKRecordZoneNotification.RecordZoneID]: The ID of the record zone that has changes.
//
// # Getting the Database Scope
//
//   - [CKRecordZoneNotification.DatabaseScope]: The type of database for the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification
//
// [Asking permission to use notifications]: https://developer.apple.com/documentation/UserNotifications/asking-permission-to-use-notifications
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
type CKRecordZoneNotification struct {
	CKNotification
}

// CKRecordZoneNotificationFromID constructs a [CKRecordZoneNotification] from an objc.ID.
//
// A notification that triggers when the contents of a record zone change.
func CKRecordZoneNotificationFromID(id objc.ID) CKRecordZoneNotification {
	return CKRecordZoneNotification{CKNotification: CKNotificationFromID(id)}
}

// NOTE: CKRecordZoneNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecordZoneNotification] class.
//
// # Getting the Record Zone ID
//
//   - [ICKRecordZoneNotification.RecordZoneID]: The ID of the record zone that has changes.
//
// # Getting the Database Scope
//
//   - [ICKRecordZoneNotification.DatabaseScope]: The type of database for the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification
type ICKRecordZoneNotification interface {
	ICKNotification

	// Topic: Getting the Record Zone ID

	// The ID of the record zone that has changes.
	RecordZoneID() ICKRecordZoneID

	// Topic: Getting the Database Scope

	// The type of database for the record zone.
	DatabaseScope() CKDatabaseScope
}

// Init initializes the instance.
func (c CKRecordZoneNotification) Init() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecordZoneNotification) Autorelease() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneNotification creates a new CKRecordZoneNotification instance.
func NewCKRecordZoneNotification() CKRecordZoneNotification {
	class := getCKRecordZoneNotificationClass()
	rv := objc.Send[CKRecordZoneNotification](objc.ID(class.class), objc.Sel("new"))
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
func NewCKRecordZoneNotificationFromRemoteNotificationDictionary(notificationDictionary foundation.INSDictionary) CKRecordZoneNotification {
	rv := objc.Send[objc.ID](objc.ID(getCKRecordZoneNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return CKRecordZoneNotificationFromID(rv)
}

// The ID of the record zone that has changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification/recordZoneID
func (c CKRecordZoneNotification) RecordZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}

// The type of database for the record zone.
//
// # Discussion
//
// This property’s value is one of the constants that [CKDatabase.Scope]
// defines.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification/databaseScope
//
// [CKDatabase.Scope]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
func (c CKRecordZoneNotification) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c.ID, objc.Sel("databaseScope"))
	return CKDatabaseScope(rv)
}
