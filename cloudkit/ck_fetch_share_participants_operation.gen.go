// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
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
// participant’s [CKFetchShareParticipantsOperation.UserIdentity] to reflect that by setting the
// [CKFetchShareParticipantsOperation.HasiCloudAccount] property to false. CloudKit associates a participant
// with their iCloud account when they accept the share.
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
// operation completes with errors, it returns a [CKFetchShareParticipantsOperation.PartialFailure] error. The
// error stores the individual errors in its [CKFetchShareParticipantsOperation.UserInfo] dictionary. Use the
// [CKFetchShareParticipantsOperation.CKPartialErrorsByItemIDKey] key to extract them.
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
// # Instance Properties
//
//   - [CKFetchShareParticipantsOperation.FetchShareParticipantsResultBlock]: The closure to execute when the operation finishes.
//   - [CKFetchShareParticipantsOperation.SetFetchShareParticipantsResultBlock]
//   - [CKFetchShareParticipantsOperation.PerShareParticipantResultBlock]: The closure to execute as the operation generates individual participants.
//   - [CKFetchShareParticipantsOperation.SetPerShareParticipantResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation
//
// [shareParticipantFetchedBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
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
// # Instance Properties
//
//   - [ICKFetchShareParticipantsOperation.FetchShareParticipantsResultBlock]: The closure to execute when the operation finishes.
//   - [ICKFetchShareParticipantsOperation.SetFetchShareParticipantsResultBlock]
//   - [ICKFetchShareParticipantsOperation.PerShareParticipantResultBlock]: The closure to execute as the operation generates individual participants.
//   - [ICKFetchShareParticipantsOperation.SetPerShareParticipantResultBlock]
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

	// Topic: Instance Properties

	// The closure to execute when the operation finishes.
	FetchShareParticipantsResultBlock() unsafe.Pointer
	SetFetchShareParticipantsResultBlock(value unsafe.Pointer)
	// The closure to execute as the operation generates individual participants.
	PerShareParticipantResultBlock() unsafe.Pointer
	SetPerShareParticipantResultBlock(value unsafe.Pointer)

	// The key to retrieve partial errors.
	CKPartialErrorsByItemIDKey() string
	// A Boolean value that indicates whether the user has an iCloud account.
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
	// An error that occurs when an operation completes with partial failures.
	PartialFailure() CKErrorCode
	SetPartialFailure(value CKErrorCode)
	// The identity of the participant.
	UserIdentity() ICKUserIdentity
	SetUserIdentity(value ICKUserIdentity)
	// The user info dictionary.
	UserInfo() string
	SetUserInfo(value string)
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
// value to the [UserIdentityLookupInfos] property before you execute this
// operation.
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
// value to the [UserIdentityLookupInfos] property before you execute this
// operation.
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

// The closure to execute when the operation finishes.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/fetchshareparticipantsresultblock
func (c CKFetchShareParticipantsOperation) FetchShareParticipantsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchShareParticipantsResultBlock"))
	return rv
}
func (c CKFetchShareParticipantsOperation) SetFetchShareParticipantsResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchShareParticipantsResultBlock:"), value)
}

// The closure to execute as the operation generates individual participants.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/pershareparticipantresultblock
func (c CKFetchShareParticipantsOperation) PerShareParticipantResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perShareParticipantResultBlock"))
	return rv
}
func (c CKFetchShareParticipantsOperation) SetPerShareParticipantResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerShareParticipantResultBlock:"), value)
}

// The key to retrieve partial errors.
//
// See: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c CKFetchShareParticipantsOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the user has an iCloud account.
//
// See: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c CKFetchShareParticipantsOperation) HasiCloudAccount() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasiCloudAccount"))
	return rv
}
func (c CKFetchShareParticipantsOperation) SetHasiCloudAccount(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasiCloudAccount:"), value)
}

// An error that occurs when an operation completes with partial failures.
//
// See: https://developer.apple.com/documentation/cloudkit/ckerror/partialfailure
func (c CKFetchShareParticipantsOperation) PartialFailure() CKErrorCode {
	rv := objc.Send[CKErrorCode](c.ID, objc.Sel("partialFailure"))
	return CKErrorCode(rv)
}
func (c CKFetchShareParticipantsOperation) SetPartialFailure(value CKErrorCode) {
	objc.Send[struct{}](c.ID, objc.Sel("setPartialFailure:"), value)
}

// The identity of the participant.
//
// See: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c CKFetchShareParticipantsOperation) UserIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}
func (c CKFetchShareParticipantsOperation) SetUserIdentity(value ICKUserIdentity) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentity:"), value)
}

// The user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c CKFetchShareParticipantsOperation) UserInfo() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userInfo"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchShareParticipantsOperation) SetUserInfo(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserInfo:"), objc.String(value))
}
