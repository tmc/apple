// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchShareParticipantsOperation] class.
var (
	_CKFetchShareParticipantsOperationClass     CKFetchShareParticipantsOperationClass
	_CKFetchShareParticipantsOperationClassOnce sync.Once
)

func getCKFetchShareParticipantsOperationClass() CKFetchShareParticipantsOperationClass {
	_CKFetchShareParticipantsOperationClassOnce.Do(func() {
		_CKFetchShareParticipantsOperationClass = CKFetchShareParticipantsOperationClass{class: objc.GetClass("CKFetchShareParticipantsOperation")}
	})
	return _CKFetchShareParticipantsOperationClass
}

// GetCKFetchShareParticipantsOperationClass returns the class object for CKFetchShareParticipantsOperation.
func GetCKFetchShareParticipantsOperationClass() CKFetchShareParticipantsOperationClass {
	return getCKFetchShareParticipantsOperationClass()
}

type CKFetchShareParticipantsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchShareParticipantsOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchShareParticipantsOperationClass) Alloc() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that converts user identities into share participants.
//
// # Overview
//
// Participants are a fundamental part of sharing in CloudKit. A participant
// provides information about a user and their participation in a share, which
// includes their identity, acceptance status, role, and permissions. The
// acceptance status manages the user’s visibilty of the shared records. The
// role and permissions control what actions the user can perform on those
// records.
//
// You don’t create participants. Instead, create an instance of
// [CKUserIdentityLookupInfo] for each user. Provide the user’s email
// address or phone number, and then use this operation to convert them into
// participants that you can add to a share. CloudKit limits the number of
// participants in a share to 100, and each participant must have an active
// iCloud account.
//
// CloudKit queries iCloud for corresponding accounts as part of the
// operation. If it doesn’t find an account, the server updates the
// participant’s [CKShareParticipant.UserIdentity] to reflect that by
// setting the [CKUserIdentity.HasiCloudAccount] property to false. CloudKit
// associates a participant with their iCloud account when they accept the
// share.
//
// Anyone with the URL of a public share can become a participant in that
// share. For a private share, the owner manages its participants. A
// participant can’t accept a private share unless the owner adds them
// first.
//
// To run the operation, add it to the container’s operation queue. The
// operation executes its callbacks on a private serial queue.
//
// The following example demonstrates how to create the operation, configure
// it, and then execute it using the default container’s operation queue:
//
// The operation calls [shareParticipantFetchedBlock] once for each item you
// provide, and CloudKit returns the participant, or an error if it can’t
// generate a particpant. CloudKit also batches per-participant errors. If the
// operation completes with errors, it returns a [partialFailure] error. The
// error stores the individual errors in its [userInfo] dictionary. Use the
// [CKPartialErrorsByItemIDKey] key to extract them.
//
// # Creating an Operation
//
//   - [CKFetchShareParticipantsOperation.InitWithUserIdentityLookupInfos]: Creates an operation for generating share participants from the specified user data.
//
// # Configuring the Operation
//
//   - [CKFetchShareParticipantsOperation.UserIdentityLookupInfos]: The user data for the participants.
//   - [CKFetchShareParticipantsOperation.SetUserIdentityLookupInfos]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation
//
// [CKPartialErrorsByItemIDKey]: https://developer.apple.com/documentation/CloudKit/CKPartialErrorsByItemIDKey
// [partialFailure]: https://developer.apple.com/documentation/CloudKit/CKError/partialFailure
// [shareParticipantFetchedBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
type CKFetchShareParticipantsOperation struct {
	CKOperation
}

// CKFetchShareParticipantsOperationFromID constructs a [CKFetchShareParticipantsOperation] from an objc.ID.
//
// An operation that converts user identities into share participants.
func CKFetchShareParticipantsOperationFromID(id objc.ID) CKFetchShareParticipantsOperation {
	return CKFetchShareParticipantsOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKFetchShareParticipantsOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchShareParticipantsOperation] class.
//
// # Creating an Operation
//
//   - [ICKFetchShareParticipantsOperation.InitWithUserIdentityLookupInfos]: Creates an operation for generating share participants from the specified user data.
//
// # Configuring the Operation
//
//   - [ICKFetchShareParticipantsOperation.UserIdentityLookupInfos]: The user data for the participants.
//   - [ICKFetchShareParticipantsOperation.SetUserIdentityLookupInfos]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation
type ICKFetchShareParticipantsOperation interface {
	ICKOperation

	// Topic: Creating an Operation

	// Creates an operation for generating share participants from the specified user data.
	InitWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKFetchShareParticipantsOperation

	// Topic: Configuring the Operation

	// The user data for the participants.
	UserIdentityLookupInfos() []CKUserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo)
}

// Init initializes the instance.
func (c CKFetchShareParticipantsOperation) Init() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchShareParticipantsOperation) Autorelease() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchShareParticipantsOperation creates a new CKFetchShareParticipantsOperation instance.
func NewCKFetchShareParticipantsOperation() CKFetchShareParticipantsOperation {
	class := getCKFetchShareParticipantsOperationClass()
	rv := objc.Send[CKFetchShareParticipantsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for generating share participants from the specified
// user data.
//
// userIdentityLookupInfos: The user data for the participants. If you specify `nil`, you must assign a
// value to the [CKFetchShareParticipantsOperation.UserIdentityLookupInfos]
// property before you execute this operation.
//
// # Discussion
//
// After you create the operation, assign a handler to the
// [fetchShareParticipantsCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/init(userIdentityLookupInfos:)
//
// [fetchShareParticipantsCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKFetchShareParticipantsOperation {
	instance := getCKFetchShareParticipantsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), objectivec.IObjectSliceToNSArray(userIdentityLookupInfos))
	return CKFetchShareParticipantsOperationFromID(rv)
}

// Creates an operation for generating share participants from the specified
// user data.
//
// userIdentityLookupInfos: The user data for the participants. If you specify `nil`, you must assign a
// value to the [CKFetchShareParticipantsOperation.UserIdentityLookupInfos]
// property before you execute this operation.
//
// # Discussion
//
// After you create the operation, assign a handler to the
// [fetchShareParticipantsCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/init(userIdentityLookupInfos:)
//
// [fetchShareParticipantsCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func (c CKFetchShareParticipantsOperation) InitWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](c.ID, objc.Sel("initWithUserIdentityLookupInfos:"), objectivec.IObjectSliceToNSArray(userIdentityLookupInfos))
	return rv
}

// The user data for the participants.
//
// # Discussion
//
// Use this property to view or change the participants user data. If you
// intend to specify or change the value of this property, do so before you
// execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c CKFetchShareParticipantsOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("userIdentityLookupInfos"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKUserIdentityLookupInfo {
		return CKUserIdentityLookupInfoFromID(id)
	})
}
func (c CKFetchShareParticipantsOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentityLookupInfos:"), objectivec.IObjectSliceToNSArray(value))
}
