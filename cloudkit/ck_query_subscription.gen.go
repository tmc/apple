// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKQuerySubscription] class.
var (
	_CKQuerySubscriptionClass     CKQuerySubscriptionClass
	_CKQuerySubscriptionClassOnce sync.Once
)

func getCKQuerySubscriptionClass() CKQuerySubscriptionClass {
	_CKQuerySubscriptionClassOnce.Do(func() {
		_CKQuerySubscriptionClass = CKQuerySubscriptionClass{class: objc.GetClass("CKQuerySubscription")}
	})
	return _CKQuerySubscriptionClass
}

// GetCKQuerySubscriptionClass returns the class object for CKQuerySubscription.
func GetCKQuerySubscriptionClass() CKQuerySubscriptionClass {
	return getCKQuerySubscriptionClass()
}

type CKQuerySubscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKQuerySubscriptionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKQuerySubscriptionClass) Alloc() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A subscription that generates push notifications when CloudKit modifies
// records that match a predicate.
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
// Query subscriptions execute whenever a change occurs in a database that
// matches the predicate and options you specify. You scope a query
// subscription to an individual record type that you provide during
// initialization. You can set the subscription’s
// [CKQuerySubscription.ZoneID] property to further specialize the
// subscription to a specific record zone in the database. This limits the
// scope of the subscription to only track changes in that record zone and
// reduces the number of notifications it generates. For more information
// about defining CloudKit-compatible predicates, see [CKQuery].
//
// Create any subscriptions on your app’s first launch. After you initialize
// a subscription, save it to the server using
// [CKModifySubscriptionsOperation]. When the operation completes, record that
// state on-device (in [UserDefaults], for example). You can then check that
// state on subsequent launches to prevent unnecessary trips to the server.
//
// To configure the notification the subscription generates, set the
// subscription’s [CKSubscription.NotificationInfo] property. Because the
// system coalesces notifications, don’t rely on them for specific changes.
// CloudKit can omit data to keep the payload size under the APNs size limit.
// Consider notifications an indication of remote changes and use
// [CKQueryOperation] to fetch the changed records. Create the operation with
// an instance of [CKQuery] that you configure with the same record type and
// predicate as the subscription. If you limit the subscription to a specific
// record zone, set the operation’s [CKQueryOperation.ZoneID] property to
// that record zone’s ID. Because [CKQueryOperation] doesn’t employ server
// change tokens, dispose of any records you cache on-device and use the
// query’s results instead.
//
// The example below shows how to create a query subscription in the user’s
// private database, configure the notifications it generates — in this
// case, silent push notifications — and then save that subscription to the
// server:
//
// # Accessing the Subscription Search Parameters
//
//   - [CKQuerySubscription.Predicate]: The matching criteria to apply to records.
//   - [CKQuerySubscription.QuerySubscriptionOptions]: Options that define the behavior of the subscription.
//
// # Accessing the Subscription Metadata
//
//   - [CKQuerySubscription.RecordType]: The type of record that the subscription queries.
//   - [CKQuerySubscription.SetRecordType]
//   - [CKQuerySubscription.ZoneID]: The ID of the record zone that the subscription queries.
//   - [CKQuerySubscription.SetZoneID]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription
//
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
type CKQuerySubscription struct {
	CKSubscription
}

// CKQuerySubscriptionFromID constructs a [CKQuerySubscription] from an objc.ID.
//
// A subscription that generates push notifications when CloudKit modifies
// records that match a predicate.
func CKQuerySubscriptionFromID(id objc.ID) CKQuerySubscription {
	return CKQuerySubscription{CKSubscription: CKSubscriptionFromID(id)}
}

// NOTE: CKQuerySubscription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKQuerySubscription] class.
//
// # Accessing the Subscription Search Parameters
//
//   - [ICKQuerySubscription.Predicate]: The matching criteria to apply to records.
//   - [ICKQuerySubscription.QuerySubscriptionOptions]: Options that define the behavior of the subscription.
//
// # Accessing the Subscription Metadata
//
//   - [ICKQuerySubscription.RecordType]: The type of record that the subscription queries.
//   - [ICKQuerySubscription.SetRecordType]
//   - [ICKQuerySubscription.ZoneID]: The ID of the record zone that the subscription queries.
//   - [ICKQuerySubscription.SetZoneID]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription
type ICKQuerySubscription interface {
	ICKSubscription

	// Topic: Accessing the Subscription Search Parameters

	// The matching criteria to apply to records.
	Predicate() foundation.NSPredicate
	// Options that define the behavior of the subscription.
	QuerySubscriptionOptions() CKQuerySubscriptionOptions

	// Topic: Accessing the Subscription Metadata

	// The type of record that the subscription queries.
	RecordType() CKRecordType
	SetRecordType(value CKRecordType)
	// The ID of the record zone that the subscription queries.
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)

	InitWithRecordTypePredicateSubscriptionIDOptions(recordType CKRecordType, predicate foundation.NSPredicate, subscriptionID CKSubscriptionID, querySubscriptionOptions CKQuerySubscriptionOptions) CKQuerySubscription
}

// Init initializes the instance.
func (c CKQuerySubscription) Init() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKQuerySubscription) Autorelease() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQuerySubscription creates a new CKQuerySubscription instance.
func NewCKQuerySubscription() CKQuerySubscription {
	class := getCKQuerySubscriptionClass()
	rv := objc.Send[CKQuerySubscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a query-based subscription from a serialized instance.
//
// aDecoder: The coder for decoding the serialized query subscription.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/init(coder:)
func NewCKQuerySubscriptionWithCoder(aDecoder foundation.INSCoder) CKQuerySubscription {
	instance := getCKQuerySubscriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKQuerySubscriptionFromID(rv)
}

func (c CKQuerySubscription) InitWithRecordTypePredicateSubscriptionIDOptions(recordType CKRecordType, predicate foundation.NSPredicate, subscriptionID CKSubscriptionID, querySubscriptionOptions CKQuerySubscriptionOptions) CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](c.ID, objc.Sel("initWithRecordType:predicate:subscriptionID:options:"), objc.String(string(recordType)), predicate, objc.String(string(subscriptionID)), querySubscriptionOptions)
	return rv
}

// The matching criteria to apply to records.
//
// # Discussion
//
// A query-based subscription uses its search predicate to identify potential
// matches for records. It combines the predicate information with the value
// in the [CKQuerySubscription.QuerySubscriptionOptions] property to determine
// when to send a push notification to the app.
//
// The search predicate defines the records that the subscription object
// monitors for changes. The system only uses the property’s value when the
// [CKSubscription.SubscriptionType] property is
// [CKSubscription.SubscriptionType.query]. Otherwise, the system ignores it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/predicate
//
// [CKSubscription.SubscriptionType.query]: https://developer.apple.com/documentation/CloudKit/CKSubscription/SubscriptionType-swift.enum/query
func (c CKQuerySubscription) Predicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}

// Options that define the behavior of the subscription.
//
// # Discussion
//
// Set the value of this property at initialization time. When you configure a
// query-based subscription, use one of the following values:
//
// - [CKQuerySubscriptionOptionsFiresOnRecordCreation] -
// [CKQuerySubscriptionOptionsFiresOnRecordUpdate] -
// [CKQuerySubscriptionOptionsFiresOnRecordDeletion]
//
// If you don’t set an option, the system throws an
// [invalidArgumentException].
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/querySubscriptionOptions
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (c CKQuerySubscription) QuerySubscriptionOptions() CKQuerySubscriptionOptions {
	rv := objc.Send[CKQuerySubscriptionOptions](c.ID, objc.Sel("querySubscriptionOptions"))
	return CKQuerySubscriptionOptions(rv)
}

// The type of record that the subscription queries.
//
// See: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/recordtype-4qgdo
func (c CKQuerySubscription) RecordType() CKRecordType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordType"))
	return CKRecordType(foundation.NSStringFromID(rv).String())
}
func (c CKQuerySubscription) SetRecordType(value CKRecordType) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordType:"), objc.String(string(value)))
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
// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/zoneID
func (c CKQuerySubscription) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
func (c CKQuerySubscription) SetZoneID(value ICKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setZoneID:"), value)
}
