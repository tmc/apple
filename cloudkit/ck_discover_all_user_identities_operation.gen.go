// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKDiscoverAllUserIdentitiesOperation] class.
var (
	_CKDiscoverAllUserIdentitiesOperationClass     CKDiscoverAllUserIdentitiesOperationClass
	_CKDiscoverAllUserIdentitiesOperationClassOnce sync.Once
)

func getCKDiscoverAllUserIdentitiesOperationClass() CKDiscoverAllUserIdentitiesOperationClass {
	_CKDiscoverAllUserIdentitiesOperationClassOnce.Do(func() {
		_CKDiscoverAllUserIdentitiesOperationClass = CKDiscoverAllUserIdentitiesOperationClass{class: objc.GetClass("CKDiscoverAllUserIdentitiesOperation")}
	})
	return _CKDiscoverAllUserIdentitiesOperationClass
}

// GetCKDiscoverAllUserIdentitiesOperationClass returns the class object for CKDiscoverAllUserIdentitiesOperation.
func GetCKDiscoverAllUserIdentitiesOperationClass() CKDiscoverAllUserIdentitiesOperationClass {
	return getCKDiscoverAllUserIdentitiesOperationClass()
}

type CKDiscoverAllUserIdentitiesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDiscoverAllUserIdentitiesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDiscoverAllUserIdentitiesOperationClass) Alloc() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that uses the device’s contacts to search for discoverable
// iCloud users.
//
// # Overview
//
// Use this operation to discover iCloud users that match entries in the
// device’s Contacts database. CloudKit uses the email addresses and phone
// numbers in each Contact record to search for a matching iCloud account.
//
// Although your app doesn’t need authorization to use the Contacts database
// to execute this operation, if it has authorization, you can use the
// [CKUserIdentity.ContactIdentifiers] property on any returned user identity
// to fetch the corresponding Contact record from the database.
//
// Before CloudKit can return a user’s identity, you must ask for their
// permission by calling
// [CKContainer.RequestApplicationPermissionCompletionHandler]. Do this as
// part of any onboarding where you can highlight the benefits of being
// discoverable within the context of your app.
//
// The operation executes the handlers you provide on an internal queue it
// manages. You must provide handlers capable of executing on a background
// queue. Tasks that need access to the main queue must redirect as
// appropriate.
//
// The operation calls [discoverAllUserIdentitiesCompletionBlock] after it
// executes and returns results. Use the completion handler to perform
// housekeeping tasks for the operation. It should also manage any failures,
// whether due to an error or an explicit cancellation.
//
// CloudKit operations have a default QoS of [QualityOfService.default].
// Operations with this service level are discretionary. The system schedules
// their execution at an optimal time according to battery level and network
// conditions, among other factors. Use the [qualityOfService] property to set
// a more appropriate QoS for the operation.
//
// The following example shows how to create the operation, configure its
// callbacks, and execute it using the default container’s queue:
//
// # Processing the Operation Results
//
//   - [CKDiscoverAllUserIdentitiesOperation.UserIdentityDiscoveredBlock]: The closure to execute for each user identity.
//   - [CKDiscoverAllUserIdentitiesOperation.SetUserIdentityDiscoveredBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation
//
// [QualityOfService.default]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
// [discoverAllUserIdentitiesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/discoverAllUserIdentitiesCompletionBlock
// [qualityOfService]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
type CKDiscoverAllUserIdentitiesOperation struct {
	CKOperation
}

// CKDiscoverAllUserIdentitiesOperationFromID constructs a [CKDiscoverAllUserIdentitiesOperation] from an objc.ID.
//
// An operation that uses the device’s contacts to search for discoverable
// iCloud users.
func CKDiscoverAllUserIdentitiesOperationFromID(id objc.ID) CKDiscoverAllUserIdentitiesOperation {
	return CKDiscoverAllUserIdentitiesOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKDiscoverAllUserIdentitiesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDiscoverAllUserIdentitiesOperation] class.
//
// # Processing the Operation Results
//
//   - [ICKDiscoverAllUserIdentitiesOperation.UserIdentityDiscoveredBlock]: The closure to execute for each user identity.
//   - [ICKDiscoverAllUserIdentitiesOperation.SetUserIdentityDiscoveredBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation
type ICKDiscoverAllUserIdentitiesOperation interface {
	ICKOperation

	// Topic: Processing the Operation Results

	// The closure to execute for each user identity.
	UserIdentityDiscoveredBlock() CKUserIdentityHandler
	SetUserIdentityDiscoveredBlock(value CKUserIdentityHandler)
}

// Init initializes the instance.
func (c CKDiscoverAllUserIdentitiesOperation) Init() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDiscoverAllUserIdentitiesOperation) Autorelease() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDiscoverAllUserIdentitiesOperation creates a new CKDiscoverAllUserIdentitiesOperation instance.
func NewCKDiscoverAllUserIdentitiesOperation() CKDiscoverAllUserIdentitiesOperation {
	class := getCKDiscoverAllUserIdentitiesOperationClass()
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The closure to execute for each user identity.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameter:
//
// - The user identity that matches an entry in the device’s Contacts.
//
// The operation executes this closure one or more times for each user
// identity it discovers. Each time the closure executes, it executes serially
// with respect to the other closures of the operation.
//
// If you intend to use this closure to process results, set it before you
// execute the operation or add the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c CKDiscoverAllUserIdentitiesOperation) UserIdentityDiscoveredBlock() CKUserIdentityHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentityDiscoveredBlock"))
	_ = rv
	return nil
}
func (c CKDiscoverAllUserIdentitiesOperation) SetUserIdentityDiscoveredBlock(value CKUserIdentityHandler) {
	block, cleanup := NewCKUserIdentityBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentityDiscoveredBlock:"), block)
}
