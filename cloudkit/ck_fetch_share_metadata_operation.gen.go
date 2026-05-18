// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchShareMetadataOperation] class.
var (
	_CKFetchShareMetadataOperationClass     CKFetchShareMetadataOperationClass
	_CKFetchShareMetadataOperationClassOnce sync.Once
)

func getCKFetchShareMetadataOperationClass() CKFetchShareMetadataOperationClass {
	_CKFetchShareMetadataOperationClassOnce.Do(func() {
		_CKFetchShareMetadataOperationClass = CKFetchShareMetadataOperationClass{class: objc.GetClass("CKFetchShareMetadataOperation")}
	})
	return _CKFetchShareMetadataOperationClass
}

// GetCKFetchShareMetadataOperationClass returns the class object for CKFetchShareMetadataOperation.
func GetCKFetchShareMetadataOperationClass() CKFetchShareMetadataOperationClass {
	return getCKFetchShareMetadataOperationClass()
}

type CKFetchShareMetadataOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchShareMetadataOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchShareMetadataOperationClass) Alloc() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that fetches metadata for one or more shares.
//
// # Overview
//
// Use this operation to fetch the metadata for one or more shares. A
// share’s metadata contains the share and details about the user’s
// participation. Fetch metadata when you want to manually accept
// participation in a share using [CKAcceptSharesOperation].
//
// For a shared record hierarchy, the fetched metadata includes the record ID
// of the share’s root record. Set [CKFetchShareMetadataOperation.ShouldFetchRootRecord] to true to fetch
// the entire root record. You can further customize this behavior using
// [CKFetchShareMetadataOperation.RootRecordDesiredKeys] to specify which fields you want to include in your
// fetch. This functionality isn’t applicable for a shared record zone
// because, unlike a shared record hierarchy, it doesn’t have a nominated
// root record.
//
// To run the operation, add it to any container’s operation queue. Returned
// metadata includes the ID of the container that stores the share. The
// operation executes its callbacks on a private serial queue.
//
// The operation calls [perShareMetadataBlock] once for each URL you provide,
// and CloudKit returns the metadata, or an error if the fetch fails. CloudKit
// also batches per-URL errors. If the operation completes with errors, it
// returns a [CKFetchShareMetadataOperation.PartialFailure] error. The error stores individual errors in its
// [CKFetchShareMetadataOperation.UserInfo] dictionary. Use the [CKFetchShareMetadataOperation.CKPartialErrorsByItemIDKey] key to extract
// them.
//
// When all of the following conditions are true, CloudKit returns a
// [CKFetchShareMetadataOperation.ParticipantMayNeedVerification] error:
//
// - There are pending participants that don’t have matched iCloud accounts.
// - The current user has an active iCloud account and isn’t an existing
// participant (pending or otherwise).
//
// On receipt of this error, call [open(_:options:completionHandler:)] with
// the share’s URL to allow CloudKit to verify the user.
//
// The following example demonstrates how to create the operation, configure
// it, and then execute it using the default container’s operation queue:
//
// # Creating an Operation
//
//   - [CKFetchShareMetadataOperation.InitWithShareURLs]: Creates an operation for fetching the metadata for the specified shares.
//
// # Configuring the Operation
//
//   - [CKFetchShareMetadataOperation.ShareURLs]: The URLs of the shares to fetch.
//   - [CKFetchShareMetadataOperation.SetShareURLs]
//   - [CKFetchShareMetadataOperation.ShouldFetchRootRecord]: A Boolean value that indicates whether to retrieve the root record.
//   - [CKFetchShareMetadataOperation.SetShouldFetchRootRecord]
//
// # Instance Properties
//
//   - [CKFetchShareMetadataOperation.FetchShareMetadataResultBlock]: The closure to execute when the operation finishes.
//   - [CKFetchShareMetadataOperation.SetFetchShareMetadataResultBlock]
//   - [CKFetchShareMetadataOperation.PerShareMetadataResultBlock]: The closure to execute as the operation fetches individual shares.
//   - [CKFetchShareMetadataOperation.SetPerShareMetadataResultBlock]
//   - [CKFetchShareMetadataOperation.RootRecordDesiredKeys]: The fields to return when fetching the root record.
//   - [CKFetchShareMetadataOperation.SetRootRecordDesiredKeys]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation
//
// [open(_:options:completionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplication/open(_:options:completionHandler:)
// [perShareMetadataBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/perShareMetadataBlock
type CKFetchShareMetadataOperation struct {
	CKOperation
}

// CKFetchShareMetadataOperationFromID constructs a [CKFetchShareMetadataOperation] from an objc.ID.
//
// An operation that fetches metadata for one or more shares.
func CKFetchShareMetadataOperationFromID(id objc.ID) CKFetchShareMetadataOperation {
	return CKFetchShareMetadataOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKFetchShareMetadataOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchShareMetadataOperation] class.
//
// # Creating an Operation
//
//   - [ICKFetchShareMetadataOperation.InitWithShareURLs]: Creates an operation for fetching the metadata for the specified shares.
//
// # Configuring the Operation
//
//   - [ICKFetchShareMetadataOperation.ShareURLs]: The URLs of the shares to fetch.
//   - [ICKFetchShareMetadataOperation.SetShareURLs]
//   - [ICKFetchShareMetadataOperation.ShouldFetchRootRecord]: A Boolean value that indicates whether to retrieve the root record.
//   - [ICKFetchShareMetadataOperation.SetShouldFetchRootRecord]
//
// # Instance Properties
//
//   - [ICKFetchShareMetadataOperation.FetchShareMetadataResultBlock]: The closure to execute when the operation finishes.
//   - [ICKFetchShareMetadataOperation.SetFetchShareMetadataResultBlock]
//   - [ICKFetchShareMetadataOperation.PerShareMetadataResultBlock]: The closure to execute as the operation fetches individual shares.
//   - [ICKFetchShareMetadataOperation.SetPerShareMetadataResultBlock]
//   - [ICKFetchShareMetadataOperation.RootRecordDesiredKeys]: The fields to return when fetching the root record.
//   - [ICKFetchShareMetadataOperation.SetRootRecordDesiredKeys]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation
type ICKFetchShareMetadataOperation interface {
	ICKOperation

	// Topic: Creating an Operation

	// Creates an operation for fetching the metadata for the specified shares.
	InitWithShareURLs(shareURLs []foundation.NSURL) CKFetchShareMetadataOperation

	// Topic: Configuring the Operation

	// The URLs of the shares to fetch.
	ShareURLs() []foundation.NSURL
	SetShareURLs(value []foundation.NSURL)
	// A Boolean value that indicates whether to retrieve the root record.
	ShouldFetchRootRecord() bool
	SetShouldFetchRootRecord(value bool)

	// Topic: Instance Properties

	// The closure to execute when the operation finishes.
	FetchShareMetadataResultBlock() unsafe.Pointer
	SetFetchShareMetadataResultBlock(value unsafe.Pointer)
	// The closure to execute as the operation fetches individual shares.
	PerShareMetadataResultBlock() unsafe.Pointer
	SetPerShareMetadataResultBlock(value unsafe.Pointer)
	// The fields to return when fetching the root record.
	RootRecordDesiredKeys() string
	SetRootRecordDesiredKeys(value string)

	// The key to retrieve partial errors.
	CKPartialErrorsByItemIDKey() string
	// An error that occurs when an operation completes with partial failures.
	PartialFailure() CKErrorCode
	SetPartialFailure(value CKErrorCode)
	// An error that occurs when the user isn’t a participant of the share.
	ParticipantMayNeedVerification() CKErrorCode
	SetParticipantMayNeedVerification(value CKErrorCode)
	// The user info dictionary.
	UserInfo() string
	SetUserInfo(value string)
}

// Init initializes the instance.
func (c CKFetchShareMetadataOperation) Init() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchShareMetadataOperation) Autorelease() CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchShareMetadataOperation creates a new CKFetchShareMetadataOperation instance.
func NewCKFetchShareMetadataOperation() CKFetchShareMetadataOperation {
	class := getCKFetchShareMetadataOperationClass()
	rv := objc.Send[CKFetchShareMetadataOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for fetching the metadata for the specified shares.
//
// shareURLs: The URLs of the shares. If you specify `nil`, you must assign a value to
// the [ShareURLs] property before you execute the operation.
//
// # Discussion
//
// After creating the operation, assign a handler to the
// [fetchShareMetadataCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/init(shareURLs:)
//
// [fetchShareMetadataCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/fetchShareMetadataCompletionBlock
func NewCKFetchShareMetadataOperationWithShareURLs(shareURLs []foundation.NSURL) CKFetchShareMetadataOperation {
	instance := getCKFetchShareMetadataOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShareURLs:"), objectivec.IObjectSliceToNSArray(shareURLs))
	return CKFetchShareMetadataOperationFromID(rv)
}

// Creates an operation for fetching the metadata for the specified shares.
//
// shareURLs: The URLs of the shares. If you specify `nil`, you must assign a value to
// the [ShareURLs] property before you execute the operation.
//
// # Discussion
//
// After creating the operation, assign a handler to the
// [fetchShareMetadataCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/init(shareURLs:)
//
// [fetchShareMetadataCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/fetchShareMetadataCompletionBlock
func (c CKFetchShareMetadataOperation) InitWithShareURLs(shareURLs []foundation.NSURL) CKFetchShareMetadataOperation {
	rv := objc.Send[CKFetchShareMetadataOperation](c.ID, objc.Sel("initWithShareURLs:"), objectivec.IObjectSliceToNSArray(shareURLs))
	return rv
}

// The URLs of the shares to fetch.
//
// # Discussion
//
// Use this property to view or change the URLs of the shares to fetch. If you
// intend to specify or change this property’s value, do so before you
// execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shareURLs
func (c CKFetchShareMetadataOperation) ShareURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shareURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}
func (c CKFetchShareMetadataOperation) SetShareURLs(value []foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setShareURLs:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that indicates whether to retrieve the root record.
//
// # Discussion
//
// For a shared record hierarchy, set this property to true to include the
// root record in the fetched share metadata. CloudKit ignores this property
// for a shared record zone because, unlike a shared record hierarchy, it
// doesn’t have a nominated root record.
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation/shouldFetchRootRecord
func (c CKFetchShareMetadataOperation) ShouldFetchRootRecord() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldFetchRootRecord"))
	return rv
}
func (c CKFetchShareMetadataOperation) SetShouldFetchRootRecord(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldFetchRootRecord:"), value)
}

// The closure to execute when the operation finishes.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/fetchsharemetadataresultblock
func (c CKFetchShareMetadataOperation) FetchShareMetadataResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchShareMetadataResultBlock"))
	return rv
}
func (c CKFetchShareMetadataOperation) SetFetchShareMetadataResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchShareMetadataResultBlock:"), value)
}

// The closure to execute as the operation fetches individual shares.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/persharemetadataresultblock
func (c CKFetchShareMetadataOperation) PerShareMetadataResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perShareMetadataResultBlock"))
	return rv
}
func (c CKFetchShareMetadataOperation) SetPerShareMetadataResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerShareMetadataResultBlock:"), value)
}

// The fields to return when fetching the root record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c CKFetchShareMetadataOperation) RootRecordDesiredKeys() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rootRecordDesiredKeys"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchShareMetadataOperation) SetRootRecordDesiredKeys(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRootRecordDesiredKeys:"), objc.String(value))
}

// The key to retrieve partial errors.
//
// See: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c CKFetchShareMetadataOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return foundation.NSStringFromID(rv).String()
}

// An error that occurs when an operation completes with partial failures.
//
// See: https://developer.apple.com/documentation/cloudkit/ckerror/partialfailure
func (c CKFetchShareMetadataOperation) PartialFailure() CKErrorCode {
	rv := objc.Send[CKErrorCode](c.ID, objc.Sel("partialFailure"))
	return CKErrorCode(rv)
}
func (c CKFetchShareMetadataOperation) SetPartialFailure(value CKErrorCode) {
	objc.Send[struct{}](c.ID, objc.Sel("setPartialFailure:"), value)
}

// An error that occurs when the user isn’t a participant of the share.
//
// See: https://developer.apple.com/documentation/cloudkit/ckerror/participantmayneedverification
func (c CKFetchShareMetadataOperation) ParticipantMayNeedVerification() CKErrorCode {
	rv := objc.Send[CKErrorCode](c.ID, objc.Sel("participantMayNeedVerification"))
	return CKErrorCode(rv)
}
func (c CKFetchShareMetadataOperation) SetParticipantMayNeedVerification(value CKErrorCode) {
	objc.Send[struct{}](c.ID, objc.Sel("setParticipantMayNeedVerification:"), value)
}

// The user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c CKFetchShareMetadataOperation) UserInfo() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userInfo"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchShareMetadataOperation) SetUserInfo(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserInfo:"), objc.String(value))
}
