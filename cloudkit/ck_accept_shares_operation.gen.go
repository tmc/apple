// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKAcceptSharesOperation] class.
var (
	_CKAcceptSharesOperationClass     CKAcceptSharesOperationClass
	_CKAcceptSharesOperationClassOnce sync.Once
)

func getCKAcceptSharesOperationClass() CKAcceptSharesOperationClass {
	_CKAcceptSharesOperationClassOnce.Do(func() {
		_CKAcceptSharesOperationClass = CKAcceptSharesOperationClass{class: objc.GetClass("CKAcceptSharesOperation")}
	})
	return _CKAcceptSharesOperationClass
}

// GetCKAcceptSharesOperationClass returns the class object for CKAcceptSharesOperation.
func GetCKAcceptSharesOperationClass() CKAcceptSharesOperationClass {
	return getCKAcceptSharesOperationClass()
}

type CKAcceptSharesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKAcceptSharesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKAcceptSharesOperationClass) Alloc() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that confirms a user’s participation in a share.
//
// # Overview
//
// Use this operation to accept participation in one or more shares. You
// create the operation with an array of share metadatas, which CloudKit
// provides to your app when the user taps or clicks a share’s [URL]. The
// method CloudKit calls varies by platform and app configuration. For more
// information, see [CKShareMetadata]. You can also fetch a share’s metadata
// using [CKFetchShareMetadataOperation].
//
// If there are several metadatas, group them by their [CKAcceptSharesOperation.ContainerIdentifier]
// and create an operation for each container. Then add the operation to each
// container’s operation queue to run it. The operation executes its
// callbacks on a private serial queue.
//
// The operation calls [perShareCompletionBlock] once for each metadata you
// provide. CloudKit returns the metadata and its related share, or an error
// if it can’t accept the share. CloudKit also batches per-metadata errors.
// If the operation completes with errors, it returns a [CKAcceptSharesOperation.PartialFailure]
// error. The error stores individual errors in its [CKAcceptSharesOperation.UserInfo] dictionary. Use
// the [CKAcceptSharesOperation.CKPartialErrorsByItemIDKey] key to extract them.
//
// After CloudKit applies all record changes, the operation calls
// [acceptSharesCompletionBlock]. When the closure executes, the server may
// continue processing residual tasks of the operation, such as creating the
// record zone in the user’s private database.
//
// The following example demonstrates how to accept a share that CloudKit
// provides to your window scene delegate. It shows how to create the
// operation, configure it, and execute it in the correct container:
//
// # Creating a Share Accept Operation
//
//   - [CKAcceptSharesOperation.InitWithShareMetadatas]: Creates an operation for accepting the specified shares.
//
// # Processing the Share Accept Results
//
//   - [CKAcceptSharesOperation.ShareMetadatas]: The share metadatas to process.
//   - [CKAcceptSharesOperation.SetShareMetadatas]
//
// # Instance Properties
//
//   - [CKAcceptSharesOperation.AcceptSharesResultBlock]: The closure to execute when the operation finishes.
//   - [CKAcceptSharesOperation.SetAcceptSharesResultBlock]
//   - [CKAcceptSharesOperation.PerShareResultBlock]: The block to execute as CloudKit processes individual shares.
//   - [CKAcceptSharesOperation.SetPerShareResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
//
// [acceptSharesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
// [perShareCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/perShareCompletionBlock
type CKAcceptSharesOperation struct {
	CKOperation
}

// CKAcceptSharesOperationFromID constructs a [CKAcceptSharesOperation] from an objc.ID.
//
// An operation that confirms a user’s participation in a share.
func CKAcceptSharesOperationFromID(id objc.ID) CKAcceptSharesOperation {
	return CKAcceptSharesOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKAcceptSharesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKAcceptSharesOperation] class.
//
// # Creating a Share Accept Operation
//
//   - [ICKAcceptSharesOperation.InitWithShareMetadatas]: Creates an operation for accepting the specified shares.
//
// # Processing the Share Accept Results
//
//   - [ICKAcceptSharesOperation.ShareMetadatas]: The share metadatas to process.
//   - [ICKAcceptSharesOperation.SetShareMetadatas]
//
// # Instance Properties
//
//   - [ICKAcceptSharesOperation.AcceptSharesResultBlock]: The closure to execute when the operation finishes.
//   - [ICKAcceptSharesOperation.SetAcceptSharesResultBlock]
//   - [ICKAcceptSharesOperation.PerShareResultBlock]: The block to execute as CloudKit processes individual shares.
//   - [ICKAcceptSharesOperation.SetPerShareResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
type ICKAcceptSharesOperation interface {
	ICKOperation

	// Topic: Creating a Share Accept Operation

	// Creates an operation for accepting the specified shares.
	InitWithShareMetadatas(shareMetadatas []CKShareMetadata) CKAcceptSharesOperation

	// Topic: Processing the Share Accept Results

	// The share metadatas to process.
	ShareMetadatas() []CKShareMetadata
	SetShareMetadatas(value []CKShareMetadata)

	// Topic: Instance Properties

	// The closure to execute when the operation finishes.
	AcceptSharesResultBlock() unsafe.Pointer
	SetAcceptSharesResultBlock(value unsafe.Pointer)
	// The block to execute as CloudKit processes individual shares.
	PerShareResultBlock() unsafe.Pointer
	SetPerShareResultBlock(value unsafe.Pointer)

	// The key to retrieve partial errors.
	CKPartialErrorsByItemIDKey() string
	// The ID of the share’s container.
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	// An error that occurs when an operation completes with partial failures.
	PartialFailure() CKErrorCode
	SetPartialFailure(value CKErrorCode)
	// The Uniform Resource Locator (URL) for inviting participants to the share.
	Url() foundation.NSURL
	SetURL(value foundation.NSURL)
	// The user info dictionary.
	UserInfo() string
	SetUserInfo(value string)
}

// Init initializes the instance.
func (c CKAcceptSharesOperation) Init() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKAcceptSharesOperation) Autorelease() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAcceptSharesOperation creates a new CKAcceptSharesOperation instance.
func NewCKAcceptSharesOperation() CKAcceptSharesOperation {
	class := getCKAcceptSharesOperationClass()
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for accepting the specified shares.
//
// shareMetadatas: The share metadatas to accept. If you specify `nil`, you must assign a
// value to the [ShareMetadatas] property before you execute the operation.
//
// # Discussion
//
// After initializing the operation, assign a handler to the
// [acceptSharesCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/init(shareMetadatas:)
//
// [acceptSharesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
func NewCKAcceptSharesOperationWithShareMetadatas(shareMetadatas []CKShareMetadata) CKAcceptSharesOperation {
	instance := getCKAcceptSharesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShareMetadatas:"), objectivec.IObjectSliceToNSArray(shareMetadatas))
	return CKAcceptSharesOperationFromID(rv)
}

// Creates an operation for accepting the specified shares.
//
// shareMetadatas: The share metadatas to accept. If you specify `nil`, you must assign a
// value to the [ShareMetadatas] property before you execute the operation.
//
// # Discussion
//
// After initializing the operation, assign a handler to the
// [acceptSharesCompletionBlock] property to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/init(shareMetadatas:)
//
// [acceptSharesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
func (c CKAcceptSharesOperation) InitWithShareMetadatas(shareMetadatas []CKShareMetadata) CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c.ID, objc.Sel("initWithShareMetadatas:"), objectivec.IObjectSliceToNSArray(shareMetadatas))
	return rv
}

// The share metadatas to process.
//
// # Discussion
//
// Use this property to view or change the metadata of the shares you want to
// process. If you intend to specify or change the value of this property, do
// so before you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/shareMetadatas
func (c CKAcceptSharesOperation) ShareMetadatas() []CKShareMetadata {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shareMetadatas"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKShareMetadata {
		return CKShareMetadataFromID(id)
	})
}
func (c CKAcceptSharesOperation) SetShareMetadatas(value []CKShareMetadata) {
	objc.Send[struct{}](c.ID, objc.Sel("setShareMetadatas:"), objectivec.IObjectSliceToNSArray(value))
}

// The closure to execute when the operation finishes.
//
// See: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/acceptsharesresultblock
func (c CKAcceptSharesOperation) AcceptSharesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("acceptSharesResultBlock"))
	return rv
}
func (c CKAcceptSharesOperation) SetAcceptSharesResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setAcceptSharesResultBlock:"), value)
}

// The block to execute as CloudKit processes individual shares.
//
// See: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/pershareresultblock
func (c CKAcceptSharesOperation) PerShareResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perShareResultBlock"))
	return rv
}
func (c CKAcceptSharesOperation) SetPerShareResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerShareResultBlock:"), value)
}

// The key to retrieve partial errors.
//
// See: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c CKAcceptSharesOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return foundation.NSStringFromID(rv).String()
}

// The ID of the share’s container.
//
// See: https://developer.apple.com/documentation/cloudkit/ckshare/metadata/containeridentifier
func (c CKAcceptSharesOperation) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKAcceptSharesOperation) SetContainerIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}

// An error that occurs when an operation completes with partial failures.
//
// See: https://developer.apple.com/documentation/cloudkit/ckerror/partialfailure
func (c CKAcceptSharesOperation) PartialFailure() CKErrorCode {
	rv := objc.Send[CKErrorCode](c.ID, objc.Sel("partialFailure"))
	return CKErrorCode(rv)
}
func (c CKAcceptSharesOperation) SetPartialFailure(value CKErrorCode) {
	objc.Send[struct{}](c.ID, objc.Sel("setPartialFailure:"), value)
}

// The Uniform Resource Locator (URL) for inviting participants to the share.
//
// See: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c CKAcceptSharesOperation) Url() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c CKAcceptSharesOperation) SetURL(value foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setURL:"), value)
}

// The user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c CKAcceptSharesOperation) UserInfo() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userInfo"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKAcceptSharesOperation) SetUserInfo(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserInfo:"), objc.String(value))
}
