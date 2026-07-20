// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKRecordZoneSubscription] class.
var (
	_CKRecordZoneSubscriptionClass     CKRecordZoneSubscriptionClass
	_CKRecordZoneSubscriptionClassOnce sync.Once
)

func getCKRecordZoneSubscriptionClass() CKRecordZoneSubscriptionClass {
	_CKRecordZoneSubscriptionClassOnce.Do(func() {
		_CKRecordZoneSubscriptionClass = CKRecordZoneSubscriptionClass{class: objc.GetClass("CKRecordZoneSubscription")}
	})
	return _CKRecordZoneSubscriptionClass
}

// GetCKRecordZoneSubscriptionClass returns the class object for CKRecordZoneSubscription.
func GetCKRecordZoneSubscriptionClass() CKRecordZoneSubscriptionClass {
	return getCKRecordZoneSubscriptionClass()
}

type CKRecordZoneSubscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordZoneSubscriptionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordZoneSubscriptionClass) Alloc() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A subscription that generates push notifications when CloudKit modifies
// records in a specific record zone.
//
// # Overview
//
// Subscriptions track the creation, modification, and deletion of records in
// a database, and are fundamental in keeping data on the user’s device up
// to date. A subscription applies only to the user that creates it. When a
// subscription registers a change, such as CloudKit saving a new record, it
// sends push notifications to the user’s devices to inform your app about
// the change. You can then fetch the changes and cache them on-device. When
// appropriate, the server excludes the device where the change originates.
//
// Record zone subscriptions execute whenever a change happens in the record
// zone you specify when you create the subscription. You can further
// specialize the subscription by setting its
// [CKDatabaseSubscription.RecordType] property to a specific record type.
// This limits the scope of the subscription to only track changes to records
// of that type and reduces the number of notifications it generates.
//
// Create any subscriptions on your app’s first launch. After you initialize
// a subscription, save it to the server using
// [CKModifySubscriptionsOperation]. When the operation completes, record that
// state on-device (in [UserDefaults], for example). You can then check that
// state on subsequent launches to prevent unnecessary trips to the server.
//
// To configure the notification that the subscription generates, set the
// subscription’s [CKSubscription.NotificationInfo] property. Because the
// system coalesces notifications, don’t rely on them for specific changes.
// CloudKit can omit data to keep the payload size under the APNs size limit.
// Consider notifications an indication of remote changes and use
// [CKFetchRecordZoneChangesOperation] to fetch the changed records. Server
// change tokens allow you to limit the fetch results to just the changes
// since your previous fetch.
//
// The example below shows how to create a record zone subscription in the
// user’s private database, configure the notifications it generates — in
// this case, silent push notifications — and then save that subscription to
// the server:
//
// # Accessing the Subscription Metadata
//
//   - [CKRecordZoneSubscription.RecordType]: The type of record that the subscription queries.
//   - [CKRecordZoneSubscription.SetRecordType]
//   - [CKRecordZoneSubscription.ZoneID]: The ID of the record zone that the subscription queries.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription
//
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
type CKRecordZoneSubscription struct {
	CKSubscription
}

// CKRecordZoneSubscriptionFromID constructs a [CKRecordZoneSubscription] from an objc.ID.
//
// A subscription that generates push notifications when CloudKit modifies
// records in a specific record zone.
func CKRecordZoneSubscriptionFromID(id objc.ID) CKRecordZoneSubscription {
	return CKRecordZoneSubscription{CKSubscription: CKSubscriptionFromID(id)}
}

// NOTE: CKRecordZoneSubscription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecordZoneSubscription] class.
//
// # Accessing the Subscription Metadata
//
//   - [ICKRecordZoneSubscription.RecordType]: The type of record that the subscription queries.
//   - [ICKRecordZoneSubscription.SetRecordType]
//   - [ICKRecordZoneSubscription.ZoneID]: The ID of the record zone that the subscription queries.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription
type ICKRecordZoneSubscription interface {
	ICKSubscription

	// Topic: Accessing the Subscription Metadata

	// The type of record that the subscription queries.
	RecordType() unsafe.Pointer
	SetRecordType(value unsafe.Pointer)
	// The ID of the record zone that the subscription queries.
	ZoneID() ICKRecordZoneID
}

// Init initializes the instance.
func (c CKRecordZoneSubscription) Init() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecordZoneSubscription) Autorelease() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneSubscription creates a new CKRecordZoneSubscription instance.
func NewCKRecordZoneSubscription() CKRecordZoneSubscription {
	class := getCKRecordZoneSubscriptionClass()
	rv := objc.Send[CKRecordZoneSubscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a zone-based subscription from a serialized instance.
//
// aDecoder: The coder for decoding the serialized record zone subscription.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/init(coder:)
func NewCKRecordZoneSubscriptionWithCoder(aDecoder foundation.INSCoder) CKRecordZoneSubscription {
	instance := getCKRecordZoneSubscriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKRecordZoneSubscriptionFromID(rv)
}

// The type of record that the subscription queries.
//
// See: https://developer.apple.com/documentation/cloudkit/ckrecordzonesubscription/recordtype-1fuqo
func (c CKRecordZoneSubscription) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordType"))
	return rv
}
func (c CKRecordZoneSubscription) SetRecordType(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordType:"), value)
}

// The ID of the record zone that the subscription queries.
//
// # Discussion
//
// This property applies to query-based subscriptions and zone-based
// subscriptions. Specifying a record zone ID limits the scope of the query to
// only the records in that zone. For zone-based subscriptions, the query
// includes all records in the specified record zone. For a query-based
// subscription, the query includes only records of a specific type in the
// specified record zone.
//
// For zone-based subscriptions, CloudKit sets this property’s value
// automatically. For all other subscription types, the default value is
// `nil`. If you want to scope your query-based subscription to a specific
// record zone, you must assign a value explicitly.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/zoneID
func (c CKRecordZoneSubscription) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
