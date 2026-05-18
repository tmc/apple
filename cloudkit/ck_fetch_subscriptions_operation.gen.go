// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
// If you assign a handler to the [CKFetchSubscriptionsOperation.CompletionBlock] property, the operation
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
// # Instance Properties
//
//   - [CKFetchSubscriptionsOperation.FetchSubscriptionsResultBlock]: The closure to execute after CloudKit retrieves all of the subscriptions.
//   - [CKFetchSubscriptionsOperation.SetFetchSubscriptionsResultBlock]
//   - [CKFetchSubscriptionsOperation.PerSubscriptionResultBlock]: The closure to execute when a subscription becomes available.
//   - [CKFetchSubscriptionsOperation.SetPerSubscriptionResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation
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
// # Instance Properties
//
//   - [ICKFetchSubscriptionsOperation.FetchSubscriptionsResultBlock]: The closure to execute after CloudKit retrieves all of the subscriptions.
//   - [ICKFetchSubscriptionsOperation.SetFetchSubscriptionsResultBlock]
//   - [ICKFetchSubscriptionsOperation.PerSubscriptionResultBlock]: The closure to execute when a subscription becomes available.
//   - [ICKFetchSubscriptionsOperation.SetPerSubscriptionResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation
type ICKFetchSubscriptionsOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Fetch Subscriptions Operation

	// The IDs of the subscriptions to fetch.
	SubscriptionIDs() string
	SetSubscriptionIDs(value string)

	// Topic: Processing the Fetch Subscription Results

	// The block to execute with the fetch results.
	FetchSubscriptionCompletionBlock() unsafe.Pointer
	SetFetchSubscriptionCompletionBlock(value unsafe.Pointer)

	// Topic: Instance Properties

	// The closure to execute after CloudKit retrieves all of the subscriptions.
	FetchSubscriptionsResultBlock() unsafe.Pointer
	SetFetchSubscriptionsResultBlock(value unsafe.Pointer)
	// The closure to execute when a subscription becomes available.
	PerSubscriptionResultBlock() unsafe.Pointer
	SetPerSubscriptionResultBlock(value unsafe.Pointer)
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
func (c CKFetchSubscriptionsOperation) SubscriptionIDs() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subscriptionIDs"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchSubscriptionsOperation) SetSubscriptionIDs(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionIDs:"), objc.String(value))
}

// The block to execute with the fetch results.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptioncompletionblock-6hhpi
func (c CKFetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchSubscriptionCompletionBlock"))
	return rv
}
func (c CKFetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchSubscriptionCompletionBlock:"), value)
}

// The closure to execute after CloudKit retrieves all of the subscriptions.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptionsresultblock
func (c CKFetchSubscriptionsOperation) FetchSubscriptionsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchSubscriptionsResultBlock"))
	return rv
}
func (c CKFetchSubscriptionsOperation) SetFetchSubscriptionsResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchSubscriptionsResultBlock:"), value)
}

// The closure to execute when a subscription becomes available.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/persubscriptionresultblock
func (c CKFetchSubscriptionsOperation) PerSubscriptionResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perSubscriptionResultBlock"))
	return rv
}
func (c CKFetchSubscriptionsOperation) SetPerSubscriptionResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerSubscriptionResultBlock:"), value)
}
