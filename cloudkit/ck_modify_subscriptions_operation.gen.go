// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKModifySubscriptionsOperation] class.
var (
	_CKModifySubscriptionsOperationClass     CKModifySubscriptionsOperationClass
	_CKModifySubscriptionsOperationClassOnce sync.Once
)

func getCKModifySubscriptionsOperationClass() CKModifySubscriptionsOperationClass {
	_CKModifySubscriptionsOperationClassOnce.Do(func() {
		_CKModifySubscriptionsOperationClass = CKModifySubscriptionsOperationClass{class: objc.GetClass("CKModifySubscriptionsOperation")}
	})
	return _CKModifySubscriptionsOperationClass
}

// GetCKModifySubscriptionsOperationClass returns the class object for CKModifySubscriptionsOperation.
func GetCKModifySubscriptionsOperationClass() CKModifySubscriptionsOperationClass {
	return getCKModifySubscriptionsOperationClass()
}

type CKModifySubscriptionsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKModifySubscriptionsOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKModifySubscriptionsOperationClass) Alloc() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation for modifying one or more subscriptions.
//
// # Overview
//
// After you create or change the configuration of a subscription, use this
// operation to save those changes to the server. You can also use this
// operation to permanently delete subscriptions.
//
// If you assign a handler to the [completionBlock] property, the operation
// calls it after it executes and passes it the results. Use the handler to
// perform any housekeeping tasks for the operation. The handler you specify
// should manage any failures, whether due to an error or an explicit
// cancellation.
//
// # Configuring the Modify Subscriptions Operation
//
//   - [CKModifySubscriptionsOperation.SubscriptionsToSave]: The subscriptions to save to the database.
//   - [CKModifySubscriptionsOperation.SetSubscriptionsToSave]
//   - [CKModifySubscriptionsOperation.SubscriptionIDsToDelete]: The IDs of the subscriptions that you want to delete.
//   - [CKModifySubscriptionsOperation.SetSubscriptionIDsToDelete]
//
// # Processing the Modify Subscription Results
//
//   - [CKModifySubscriptionsOperation.ModifySubscriptionsCompletionBlock]: The closure to execute after the operation modifies the subscriptions.
//   - [CKModifySubscriptionsOperation.SetModifySubscriptionsCompletionBlock]
//
// # Instance Properties
//
//   - [CKModifySubscriptionsOperation.PerSubscriptionDeleteBlock]: The closure to execute when CloudKit deletes a subscription.
//   - [CKModifySubscriptionsOperation.SetPerSubscriptionDeleteBlock]
//   - [CKModifySubscriptionsOperation.PerSubscriptionSaveBlock]: The closure to execute when CloudKit saves a subscription.
//   - [CKModifySubscriptionsOperation.SetPerSubscriptionSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation
//
// [completionBlock]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
type CKModifySubscriptionsOperation struct {
	CKDatabaseOperation
}

// CKModifySubscriptionsOperationFromID constructs a [CKModifySubscriptionsOperation] from an objc.ID.
//
// An operation for modifying one or more subscriptions.
func CKModifySubscriptionsOperationFromID(id objc.ID) CKModifySubscriptionsOperation {
	return CKModifySubscriptionsOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKModifySubscriptionsOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKModifySubscriptionsOperation] class.
//
// # Configuring the Modify Subscriptions Operation
//
//   - [ICKModifySubscriptionsOperation.SubscriptionsToSave]: The subscriptions to save to the database.
//   - [ICKModifySubscriptionsOperation.SetSubscriptionsToSave]
//   - [ICKModifySubscriptionsOperation.SubscriptionIDsToDelete]: The IDs of the subscriptions that you want to delete.
//   - [ICKModifySubscriptionsOperation.SetSubscriptionIDsToDelete]
//
// # Processing the Modify Subscription Results
//
//   - [ICKModifySubscriptionsOperation.ModifySubscriptionsCompletionBlock]: The closure to execute after the operation modifies the subscriptions.
//   - [ICKModifySubscriptionsOperation.SetModifySubscriptionsCompletionBlock]
//
// # Instance Properties
//
//   - [ICKModifySubscriptionsOperation.PerSubscriptionDeleteBlock]: The closure to execute when CloudKit deletes a subscription.
//   - [ICKModifySubscriptionsOperation.SetPerSubscriptionDeleteBlock]
//   - [ICKModifySubscriptionsOperation.PerSubscriptionSaveBlock]: The closure to execute when CloudKit saves a subscription.
//   - [ICKModifySubscriptionsOperation.SetPerSubscriptionSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation
type ICKModifySubscriptionsOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Modify Subscriptions Operation

	// The subscriptions to save to the database.
	SubscriptionsToSave() []CKSubscription
	SetSubscriptionsToSave(value []CKSubscription)
	// The IDs of the subscriptions that you want to delete.
	SubscriptionIDsToDelete() unsafe.Pointer
	SetSubscriptionIDsToDelete(value kernel.Pointer)

	// Topic: Processing the Modify Subscription Results

	// The closure to execute after the operation modifies the subscriptions.
	ModifySubscriptionsCompletionBlock() func(kernel.Pointer, kernel.Pointer, kernel.Pointer)
	SetModifySubscriptionsCompletionBlock(value func(kernel.Pointer, kernel.Pointer, kernel.Pointer))

	// Topic: Instance Properties

	// The closure to execute when CloudKit deletes a subscription.
	PerSubscriptionDeleteBlock() unsafe.Pointer
	SetPerSubscriptionDeleteBlock(value kernel.Pointer)
	// The closure to execute when CloudKit saves a subscription.
	PerSubscriptionSaveBlock() unsafe.Pointer
	SetPerSubscriptionSaveBlock(value kernel.Pointer)
}

// Init initializes the instance.
func (c CKModifySubscriptionsOperation) Init() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKModifySubscriptionsOperation) Autorelease() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifySubscriptionsOperation creates a new CKModifySubscriptionsOperation instance.
func NewCKModifySubscriptionsOperation() CKModifySubscriptionsOperation {
	class := getCKModifySubscriptionsOperationClass()
	rv := objc.Send[CKModifySubscriptionsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The subscriptions to save to the database.
//
// # Discussion
//
// This property contains the subscriptions that you want to save. Its initial
// value is the array that you pass to the
// [init(subscriptionsToSave:subscriptionIDsToDelete:)] method. Modify this
// property as necessary before you execute the operation or submit it to a
// queue. After CloudKit saves the subscriptions, it begins generating push
// notifications according to their criteria.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/subscriptionsToSave
//
// [init(subscriptionsToSave:subscriptionIDsToDelete:)]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/init(subscriptionsToSave:subscriptionIDsToDelete:)
func (c CKModifySubscriptionsOperation) SubscriptionsToSave() []CKSubscription {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("subscriptionsToSave"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSubscription {
		return CKSubscriptionFromID(id)
	})
}
func (c CKModifySubscriptionsOperation) SetSubscriptionsToSave(value []CKSubscription) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionsToSave:"), objectivec.IObjectSliceToNSArray(value))
}

// The IDs of the subscriptions that you want to delete.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/subscriptionidstodelete-3534e
func (c CKModifySubscriptionsOperation) SubscriptionIDsToDelete() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("subscriptionIDsToDelete"))
	return rv
}
func (c CKModifySubscriptionsOperation) SetSubscriptionIDsToDelete(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionIDsToDelete:"), value)
}

// The closure to execute after the operation modifies the subscriptions.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionscompletionblock-7l56
func (c CKModifySubscriptionsOperation) ModifySubscriptionsCompletionBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modifySubscriptionsCompletionBlock"))
	return objectivec.Object{ID: rv}
}
func (c CKModifySubscriptionsOperation) SetModifySubscriptionsCompletionBlock(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setModifySubscriptionsCompletionBlock:"), value)
}

// The closure to execute when CloudKit deletes a subscription.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptiondeleteblock-5ke2l
func (c CKModifySubscriptionsOperation) PerSubscriptionDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perSubscriptionDeleteBlock"))
	return rv
}
func (c CKModifySubscriptionsOperation) SetPerSubscriptionDeleteBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerSubscriptionDeleteBlock:"), value)
}

// The closure to execute when CloudKit saves a subscription.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptionsaveblock-8y9zn
func (c CKModifySubscriptionsOperation) PerSubscriptionSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perSubscriptionSaveBlock"))
	return rv
}
func (c CKModifySubscriptionsOperation) SetPerSubscriptionSaveBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerSubscriptionSaveBlock:"), value)
}
