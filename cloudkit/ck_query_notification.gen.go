// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKQueryNotification] class.
var (
	_CKQueryNotificationClass     CKQueryNotificationClass
	_CKQueryNotificationClassOnce sync.Once
)

func getCKQueryNotificationClass() CKQueryNotificationClass {
	_CKQueryNotificationClassOnce.Do(func() {
		_CKQueryNotificationClass = CKQueryNotificationClass{class: objc.GetClass("CKQueryNotification")}
	})
	return _CKQueryNotificationClass
}

// GetCKQueryNotificationClass returns the class object for CKQueryNotification.
func GetCKQueryNotificationClass() CKQueryNotificationClass {
	return getCKQueryNotificationClass()
}

type CKQueryNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKQueryNotificationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKQueryNotificationClass) Alloc() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A notification that triggers when a record that matches the
// subscription’s predicate changes.
//
// # Overview
//
// Query subscriptions execute when a record that matches the subscription’s
// predicate changes, for example, when the user modifies a field’s value in
// the record. When CloudKit registers the change, it sends push notifications
// to the user’s devices to inform your app about the change. You can then
// fetch the changes and cache them on-device. When appropriate, CloudKit
// excludes the device where the change originates.
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
// Don’t rely on push notifications for changes because the system can
// coalesce them. CloudKit can omit data to keep the notification’s payload
// size under the APNs size limit. If you use [CKSubscription.DesiredKeys] to
// include extra data in the payload, the server removes that first. A
// notification’s [CKNotification.IsPruned] property is true if CloudKit
// omits data.
//
// Consider notifications an indication of remote changes. Use
// [CKDatabaseNotification.DatabaseScope] to determine which database contains
// the changed record. To fetch the changes, configure an instance of
// [CKQueryOperation] to match the subscription and then execute it in the
// database. CloudKit returns all records that match the predicate, including
// the changed record. Dispose of any records you cache on-device and use the
// operation’s results instead.
//
// You don’t instantiate this class. Instead, implement
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)] in
// your app delegate. Initialize [CKNotification] with the `userInfo`
// dictionary that CloudKit passes to the method. This returns an instance of
// the appropriate subclass. Use the [CKNotification.NotificationType]
// property to determine the type. Then cast to that type to access
// type-specific properties and methods.
//
// # Getting the Database Scope
//
//   - [CKQueryNotification.DatabaseScope]: The type of database for the record zone.
//
// # Getting the Notification Attributes
//
//   - [CKQueryNotification.QueryNotificationReason]: The event that triggers the push notification.
//
// # Getting the Record Information
//
//   - [CKQueryNotification.RecordID]: The ID of the record that CloudKit creates, updates, or deletes.
//   - [CKQueryNotification.RecordFields]: A dictionary of fields that have changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification
//
// [Asking permission to use notifications]: https://developer.apple.com/documentation/UserNotifications/asking-permission-to-use-notifications
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
type CKQueryNotification struct {
	CKNotification
}

// CKQueryNotificationFromID constructs a [CKQueryNotification] from an objc.ID.
//
// A notification that triggers when a record that matches the
// subscription’s predicate changes.
func CKQueryNotificationFromID(id objc.ID) CKQueryNotification {
	return CKQueryNotification{CKNotification: CKNotificationFromID(id)}
}

// NOTE: CKQueryNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKQueryNotification] class.
//
// # Getting the Database Scope
//
//   - [ICKQueryNotification.DatabaseScope]: The type of database for the record zone.
//
// # Getting the Notification Attributes
//
//   - [ICKQueryNotification.QueryNotificationReason]: The event that triggers the push notification.
//
// # Getting the Record Information
//
//   - [ICKQueryNotification.RecordID]: The ID of the record that CloudKit creates, updates, or deletes.
//   - [ICKQueryNotification.RecordFields]: A dictionary of fields that have changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification
type ICKQueryNotification interface {
	ICKNotification

	// Topic: Getting the Database Scope

	// The type of database for the record zone.
	DatabaseScope() CKDatabaseScope

	// Topic: Getting the Notification Attributes

	// The event that triggers the push notification.
	QueryNotificationReason() CKQueryNotificationReason

	// Topic: Getting the Record Information

	// The ID of the record that CloudKit creates, updates, or deletes.
	RecordID() ICKRecordID
	// A dictionary of fields that have changes.
	RecordFields() foundation.INSDictionary
}

// Init initializes the instance.
func (c CKQueryNotification) Init() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKQueryNotification) Autorelease() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryNotification creates a new CKQueryNotification instance.
func NewCKQueryNotification() CKQueryNotification {
	class := getCKQueryNotificationClass()
	rv := objc.Send[CKQueryNotification](objc.ID(class.class), objc.Sel("new"))
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
func NewCKQueryNotificationFromRemoteNotificationDictionary(notificationDictionary foundation.INSDictionary) CKQueryNotification {
	rv := objc.Send[objc.ID](objc.ID(getCKQueryNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return CKQueryNotificationFromID(rv)
}

// The type of database for the record zone.
//
// # Discussion
//
// This property’s value is one of the constants that [CKDatabase.Scope]
// defines.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/databaseScope
//
// [CKDatabase.Scope]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
func (c CKQueryNotification) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c.ID, objc.Sel("databaseScope"))
	return CKDatabaseScope(rv)
}

// The event that triggers the push notification.
//
// # Discussion
//
// Subscription notifications result from the creation, deletion, or updating
// of a single record. The record in question must match the subscription’s
// predicate for an event to trigger.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/queryNotificationReason
func (c CKQueryNotification) QueryNotificationReason() CKQueryNotificationReason {
	rv := objc.Send[CKQueryNotificationReason](c.ID, objc.Sel("queryNotificationReason"))
	return CKQueryNotificationReason(rv)
}

// The ID of the record that CloudKit creates, updates, or deletes.
//
// # Discussion
//
// Use this value to fetch the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/recordID
func (c CKQueryNotification) RecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// A dictionary of fields that have changes.
//
// # Discussion
//
// For updated and created records, this property contains the
// subscription’s desired keys. When you configure the notification info of
// a subscription, you specify the names of one or more fields in the
// [CKSubscription.DesiredKeys] property. When a push notification triggers,
// CloudKit retrieves the values for each of those keys from the record and
// includes them in the notification’s payload.
//
// For query notifications that you fetch from a container, all keys and
// values are present. For query notifications that you create from push
// notifications, one or more keys and values may be missing. Push
// notification payloads have a size limit, and CloudKit can exclude record
// fields when a payload exceeds that limit. For information about the order,
// see the overview of this class.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/recordFields
func (c CKQueryNotification) RecordFields() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordFields"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
