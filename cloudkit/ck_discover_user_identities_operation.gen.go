// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKDiscoverUserIdentitiesOperation] class.
var (
	_CKDiscoverUserIdentitiesOperationClass     CKDiscoverUserIdentitiesOperationClass
	_CKDiscoverUserIdentitiesOperationClassOnce sync.Once
)

func getCKDiscoverUserIdentitiesOperationClass() CKDiscoverUserIdentitiesOperationClass {
	_CKDiscoverUserIdentitiesOperationClassOnce.Do(func() {
		_CKDiscoverUserIdentitiesOperationClass = CKDiscoverUserIdentitiesOperationClass{class: objc.GetClass("CKDiscoverUserIdentitiesOperation")}
	})
	return _CKDiscoverUserIdentitiesOperationClass
}

// GetCKDiscoverUserIdentitiesOperationClass returns the class object for CKDiscoverUserIdentitiesOperation.
func GetCKDiscoverUserIdentitiesOperationClass() CKDiscoverUserIdentitiesOperationClass {
	return getCKDiscoverUserIdentitiesOperationClass()
}

type CKDiscoverUserIdentitiesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDiscoverUserIdentitiesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDiscoverUserIdentitiesOperationClass) Alloc() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that uses the provided criteria to search for discoverable
// iCloud users.
//
// # Overview
//
// Use this operation to discover one or more iCloud users that match identity
// information you provide, such as email addresses and phone numbers.
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
// The operation calls [discoverUserIdentitiesCompletionBlock] after it
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
// # Configuring the Operation
//
//   - [CKDiscoverUserIdentitiesOperation.UserIdentityLookupInfos]: The lookup info for discovering user identities.
//   - [CKDiscoverUserIdentitiesOperation.SetUserIdentityLookupInfos]
//
// # Processing the Results
//
//   - [CKDiscoverUserIdentitiesOperation.UserIdentityDiscoveredBlock]: The closure to execute for each user identity.
//   - [CKDiscoverUserIdentitiesOperation.SetUserIdentityDiscoveredBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation
//
// [QualityOfService.default]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
// [discoverUserIdentitiesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/discoverUserIdentitiesCompletionBlock
// [qualityOfService]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
type CKDiscoverUserIdentitiesOperation struct {
	CKOperation
}

// CKDiscoverUserIdentitiesOperationFromID constructs a [CKDiscoverUserIdentitiesOperation] from an objc.ID.
//
// An operation that uses the provided criteria to search for discoverable
// iCloud users.
func CKDiscoverUserIdentitiesOperationFromID(id objc.ID) CKDiscoverUserIdentitiesOperation {
	return CKDiscoverUserIdentitiesOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKDiscoverUserIdentitiesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDiscoverUserIdentitiesOperation] class.
//
// # Configuring the Operation
//
//   - [ICKDiscoverUserIdentitiesOperation.UserIdentityLookupInfos]: The lookup info for discovering user identities.
//   - [ICKDiscoverUserIdentitiesOperation.SetUserIdentityLookupInfos]
//
// # Processing the Results
//
//   - [ICKDiscoverUserIdentitiesOperation.UserIdentityDiscoveredBlock]: The closure to execute for each user identity.
//   - [ICKDiscoverUserIdentitiesOperation.SetUserIdentityDiscoveredBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation
type ICKDiscoverUserIdentitiesOperation interface {
	ICKOperation

	// Topic: Configuring the Operation

	// The lookup info for discovering user identities.
	UserIdentityLookupInfos() []CKUserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo)

	// Topic: Processing the Results

	// The closure to execute for each user identity.
	UserIdentityDiscoveredBlock() CKUserIdentityCKUserIdentityLookupInfoHandler
	SetUserIdentityDiscoveredBlock(value CKUserIdentityCKUserIdentityLookupInfoHandler)
}

// Init initializes the instance.
func (c CKDiscoverUserIdentitiesOperation) Init() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDiscoverUserIdentitiesOperation) Autorelease() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDiscoverUserIdentitiesOperation creates a new CKDiscoverUserIdentitiesOperation instance.
func NewCKDiscoverUserIdentitiesOperation() CKDiscoverUserIdentitiesOperation {
	class := getCKDiscoverUserIdentitiesOperationClass()
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for discovering the user identities of the specified
// lookup infos.
//
// userIdentityLookupInfos: An array that contains instances of [CKUserIdentityLookupInfo]. CloudKit
// uses this parameter as the default value for the
// [CKDiscoverUserIdentitiesOperation.UserIdentityLookupInfos] property. If
// you specify `nil`, you must assign a value to that property before you
// execute the operation.
//
// # Discussion
//
// After you create the operation, assign a handler to
// [discoverUserIdentitiesCompletionBlock] so that you can process the search
// results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/init(userIdentityLookupInfos:)
//
// [discoverUserIdentitiesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/discoverUserIdentitiesCompletionBlock
func NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKDiscoverUserIdentitiesOperation {
	instance := getCKDiscoverUserIdentitiesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), objectivec.IObjectSliceToNSArray(userIdentityLookupInfos))
	return CKDiscoverUserIdentitiesOperationFromID(rv)
}

// The lookup info for discovering user identities.
//
// # Discussion
//
// Use this property to view or change the lookup info that CloudKit uses to
// discover user identities. If you intend to modify this property’s value,
// do so before you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityLookupInfos
func (c CKDiscoverUserIdentitiesOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("userIdentityLookupInfos"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKUserIdentityLookupInfo {
		return CKUserIdentityLookupInfoFromID(id)
	})
}
func (c CKDiscoverUserIdentitiesOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentityLookupInfos:"), objectivec.IObjectSliceToNSArray(value))
}

// The closure to execute for each user identity.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The user identity. - The lookup info that corresponds to the user
// identity.
//
// The operation executes this closure one or more times for each user
// identity it discovers. Each time the closure executes, it executes serially
// with respect to the other closures of the operation.
//
// If you intend to use this closure to process results, set it before you
// execute the operation or add the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c CKDiscoverUserIdentitiesOperation) UserIdentityDiscoveredBlock() CKUserIdentityCKUserIdentityLookupInfoHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentityDiscoveredBlock"))
	_ = rv
	return nil
}
func (c CKDiscoverUserIdentitiesOperation) SetUserIdentityDiscoveredBlock(value CKUserIdentityCKUserIdentityLookupInfoHandler) {
	block, cleanup := NewCKUserIdentityCKUserIdentityLookupInfoBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentityDiscoveredBlock:"), block)
}
