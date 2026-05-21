// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSubscription] class.
var (
	_CKSubscriptionClass     CKSubscriptionClass
	_CKSubscriptionClassOnce sync.Once
)

func getCKSubscriptionClass() CKSubscriptionClass {
	_CKSubscriptionClassOnce.Do(func() {
		_CKSubscriptionClass = CKSubscriptionClass{class: objc.GetClass("CKSubscription")}
	})
	return _CKSubscriptionClass
}

// GetCKSubscriptionClass returns the class object for CKSubscription.
func GetCKSubscriptionClass() CKSubscriptionClass {
	return getCKSubscriptionClass()
}

type CKSubscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSubscriptionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSubscriptionClass) Alloc() CKSubscription {
	rv := objc.Send[CKSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for subscriptions.
//
// # Overview
//
// A subscription acts like a persistent query on the server that can track
// the creation, deletion, and modification of records. When changes occur,
// they trigger the delivery of push notifications so that your app can
// respond appropriately.
//
// Subscriptions don’t become active until you save them to the server and
// the server has time to index them. To save a subscription, use an instance
// of [CKModifySubscriptionsOperation] or the
// [CKDatabase.SaveSubscriptionCompletionHandler] method of [CKDatabase]. To
// cancel a subscription, delete the corresponding subscription from the
// server.
//
// Most of a subscription’s configuration happens at initialization time.
// You must, however, specify how to deliver push notifications to the
// user’s device. Use the [CKSubscription.NotificationInfo] property to
// configure the delivery options. You must save the subscription before the
// changes take effect.
//
// # Handling the Resulting Push Notifications
//
// When CloudKit modifies a record and triggers a subscription, the server
// sends push notifications to all devices with that subscription except for
// the one that makes the original changes. For subscription-based push
// notifications, the server can add data to the notification payload that
// indicates the condition that triggers the notification. In the
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]
// method of your app delegate, create a [CKNotification] object from the
// provided `userInfo` dictionary. You can then query it for the information
// that’s relevant to the notification.
//
// In addition to sending a record ID with a push notification, you can ask
// the server to send a limited amount of data from the record that triggers
// the notification. Use the [CKSubscription.DesiredKeys] property of the
// object you assign to [CKSubscription.NotificationInfo] to specify the keys
// to include.
//
// APNs limits the size of a push notification’s payload and CloudKit may
// omit keys and other pieces of data to keep the payload’s size under that
// limit. If this happens, you can fetch the entire payload from the server
// using an instance of [CKFetchNotificationChangesOperation]. This operation
// provides instances of [CKQueryNotification] or [CKRecordZoneNotification],
// which contain information about the push notifications that CloudKit
// delivers to your app.
//
// # Specifying the Push Notification Data
//
//   - [CKSubscription.NotificationInfo]: The configuration for a subscription’s push notifications.
//   - [CKSubscription.SetNotificationInfo]
//
// # Accessing the Subscription Metadata
//
//   - [CKSubscription.SubscriptionID]: The subscription’s unique identifier.
//   - [CKSubscription.SetSubscriptionID]
//   - [CKSubscription.SubscriptionType]: The behavior that a subscription provides.
//
// # Initializers
//
//   - [CKSubscription.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription
//
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
type CKSubscription struct {
	objectivec.Object
}

// CKSubscriptionFromID constructs a [CKSubscription] from an objc.ID.
//
// An abstract base class for subscriptions.
func CKSubscriptionFromID(id objc.ID) CKSubscription {
	return CKSubscription{objectivec.Object{ID: id}}
}

// NOTE: CKSubscription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSubscription] class.
//
// # Specifying the Push Notification Data
//
//   - [ICKSubscription.NotificationInfo]: The configuration for a subscription’s push notifications.
//   - [ICKSubscription.SetNotificationInfo]
//
// # Accessing the Subscription Metadata
//
//   - [ICKSubscription.SubscriptionID]: The subscription’s unique identifier.
//   - [ICKSubscription.SetSubscriptionID]
//   - [ICKSubscription.SubscriptionType]: The behavior that a subscription provides.
//
// # Initializers
//
//   - [ICKSubscription.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription
type ICKSubscription interface {
	objectivec.IObject

	// Topic: Specifying the Push Notification Data

	// The configuration for a subscription’s push notifications.
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)

	// Topic: Accessing the Subscription Metadata

	// The subscription’s unique identifier.
	SubscriptionID() CKSubscriptionID
	SetSubscriptionID(value CKSubscriptionID)
	// The behavior that a subscription provides.
	SubscriptionType() CKSubscriptionType

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CKSubscription

	// The names of fields to include in the push notification’s payload.
	DesiredKeys() unsafe.Pointer
	SetDesiredKeys(value kernel.Pointer)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKSubscription) Init() CKSubscription {
	rv := objc.Send[CKSubscription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSubscription) Autorelease() CKSubscription {
	rv := objc.Send[CKSubscription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSubscription creates a new CKSubscription instance.
func NewCKSubscription() CKSubscription {
	class := getCKSubscriptionClass()
	rv := objc.Send[CKSubscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/init(coder:)
func NewCKSubscriptionWithCoder(coder foundation.INSCoder) CKSubscription {
	instance := getCKSubscriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CKSubscriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/init(coder:)
func (c CKSubscription) InitWithCoder(coder foundation.INSCoder) CKSubscription {
	rv := objc.Send[CKSubscription](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CKSubscription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The configuration for a subscription’s push notifications.
//
// # Discussion
//
// If you want the system to display your subscription’s push notifications,
// assign a value to this property. The server uses the configuration you
// provide to determine the delivery options for notifications. For example,
// you can specify the text to display to the user, and the sound to play. You
// can also specify which fields of the record to include in the
// notification’s payload.
//
// If you don’t assign a value to this property, CloudKit still sends push
// notifications, but the system doesn’t display them to the user. The
// default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/notificationInfo-swift.property
func (c CKSubscription) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("notificationInfo"))
	return CKNotificationInfoFromID(objc.ID(rv))
}
func (c CKSubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[struct{}](c.ID, objc.Sel("setNotificationInfo:"), value)
}

// The subscription’s unique identifier.
//
// See: https://developer.apple.com/documentation/cloudkit/cksubscription/subscriptionid-6fp3j
func (c CKSubscription) SubscriptionID() CKSubscriptionID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subscriptionID"))
	return CKSubscriptionID(foundation.NSStringFromID(rv).String())
}
func (c CKSubscription) SetSubscriptionID(value CKSubscriptionID) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionID:"), objc.String(string(value)))
}

// The behavior that a subscription provides.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/subscriptionType-swift.property
func (c CKSubscription) SubscriptionType() CKSubscriptionType {
	rv := objc.Send[CKSubscriptionType](c.ID, objc.Sel("subscriptionType"))
	return CKSubscriptionType(rv)
}

// The names of fields to include in the push notification’s payload.
//
// See: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c CKSubscription) DesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("desiredKeys"))
	return rv
}
func (c CKSubscription) SetDesiredKeys(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setDesiredKeys:"), value)
}
