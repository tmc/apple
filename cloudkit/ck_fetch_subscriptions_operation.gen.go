// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchSubscriptionsOperation] class.
var (
	_CKFetchSubscriptionsOperationClass     CKFetchSubscriptionsOperationClass
	_CKFetchSubscriptionsOperationClassOnce sync.Once
)

func getCKFetchSubscriptionsOperationClass() CKFetchSubscriptionsOperationClass {
	_CKFetchSubscriptionsOperationClassOnce.Do(func() {
		_CKFetchSubscriptionsOperationClass = CKFetchSubscriptionsOperationClass{class: objc.GetClass("CKFetchSubscriptionsOperation")}
	})
	return _CKFetchSubscriptionsOperationClass
}

// GetCKFetchSubscriptionsOperationClass returns the class object for CKFetchSubscriptionsOperation.
func GetCKFetchSubscriptionsOperationClass() CKFetchSubscriptionsOperationClass {
	return getCKFetchSubscriptionsOperationClass()
}

type CKFetchSubscriptionsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchSubscriptionsOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchSubscriptionsOperationClass) Alloc() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation for fetching subscriptions.
//
// # Overview
//
// A fetch subscriptions operation retrieves subscriptions (with IDs you
// already know) from iCloud and can fetch all subscriptions for the current
// user.
//
// You might fetch subscriptions so you can examine or modify their parameters
// — for example, to adjust the delivery options for push notifications that
// the subscription generates.
//
// If you assign a handler to the [completionBlock] property, the operation
// calls it after it executes and passes it the results. Use the handler to
// perform any housekeeping tasks for the operation. The handler you specify
// should manage any failures, whether due to an error or an explicit
// cancellation.
//
// # Configuring the Fetch Subscriptions Operation
//
//   - [CKFetchSubscriptionsOperation.SubscriptionIDs]: The IDs of the subscriptions to fetch.
//   - [CKFetchSubscriptionsOperation.SetSubscriptionIDs]
//
// # Processing the Fetch Subscription Results
//
//   - [CKFetchSubscriptionsOperation.FetchSubscriptionCompletionBlock]: The block to execute with the fetch results.
//   - [CKFetchSubscriptionsOperation.SetFetchSubscriptionCompletionBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation
//
// [completionBlock]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
type CKFetchSubscriptionsOperation struct {
	CKDatabaseOperation
}

// CKFetchSubscriptionsOperationFromID constructs a [CKFetchSubscriptionsOperation] from an objc.ID.
//
// An operation for fetching subscriptions.
func CKFetchSubscriptionsOperationFromID(id objc.ID) CKFetchSubscriptionsOperation {
	return CKFetchSubscriptionsOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchSubscriptionsOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchSubscriptionsOperation] class.
//
// # Configuring the Fetch Subscriptions Operation
//
//   - [ICKFetchSubscriptionsOperation.SubscriptionIDs]: The IDs of the subscriptions to fetch.
//   - [ICKFetchSubscriptionsOperation.SetSubscriptionIDs]
//
// # Processing the Fetch Subscription Results
//
//   - [ICKFetchSubscriptionsOperation.FetchSubscriptionCompletionBlock]: The block to execute with the fetch results.
//   - [ICKFetchSubscriptionsOperation.SetFetchSubscriptionCompletionBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation
type ICKFetchSubscriptionsOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Fetch Subscriptions Operation

	// The IDs of the subscriptions to fetch.
	SubscriptionIDs() []CKSubscriptionID
	SetSubscriptionIDs(value []CKSubscriptionID)

	// Topic: Processing the Fetch Subscription Results

	// The block to execute with the fetch results.
	FetchSubscriptionCompletionBlock() objectivec.IObject
	SetFetchSubscriptionCompletionBlock(value objectivec.IObject)
}

// Init initializes the instance.
func (c CKFetchSubscriptionsOperation) Init() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchSubscriptionsOperation) Autorelease() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchSubscriptionsOperation creates a new CKFetchSubscriptionsOperation instance.
func NewCKFetchSubscriptionsOperation() CKFetchSubscriptionsOperation {
	class := getCKFetchSubscriptionsOperationClass()
	rv := objc.Send[CKFetchSubscriptionsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an operation that fetches all of the user’s subscriptions.
//
// # Discussion
//
// After creating the operation, set the [fetchSubscriptionCompletionBlock]
// property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/fetchAllSubscriptionsOperation()
//
// [fetchSubscriptionCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/fetchSubscriptionCompletionBlock-207ep
func (_CKFetchSubscriptionsOperationClass CKFetchSubscriptionsOperationClass) FetchAllSubscriptionsOperation() CKFetchSubscriptionsOperation {
	rv := objc.Send[objc.ID](objc.ID(_CKFetchSubscriptionsOperationClass.class), objc.Sel("fetchAllSubscriptionsOperation"))
	return CKFetchSubscriptionsOperationFromID(rv)
}

// The IDs of the subscriptions to fetch.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/subscriptionids-17f4q
func (c CKFetchSubscriptionsOperation) SubscriptionIDs() []CKSubscriptionID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("subscriptionIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSubscriptionID {
		return CKSubscriptionID(foundation.NSStringFromID(id).String())
	})
}
func (c CKFetchSubscriptionsOperation) SetSubscriptionIDs(value []CKSubscriptionID) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionIDs:"), objectivec.StringSliceToNSArray(value))
}

// The block to execute with the fetch results.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptioncompletionblock-6hhpi
func (c CKFetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fetchSubscriptionCompletionBlock"))
	return objectivec.Object{ID: rv}
}
func (c CKFetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchSubscriptionCompletionBlock:"), value)
}
