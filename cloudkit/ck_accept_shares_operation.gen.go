// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

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
// provides to your app when the user taps or clicks a share’s
// [CKShare.URL]. The method CloudKit calls varies by platform and app
// configuration. For more information, see [CKShareMetadata]. You can also
// fetch a share’s metadata using [CKFetchShareMetadataOperation].
//
// If there are several metadatas, group them by their
// [CKShareMetadata.ContainerIdentifier] and create an operation for each
// container. Then add the operation to each container’s operation queue to
// run it. The operation executes its callbacks on a private serial queue.
//
// The operation calls [perShareCompletionBlock] once for each metadata you
// provide. CloudKit returns the metadata and its related share, or an error
// if it can’t accept the share. CloudKit also batches per-metadata errors.
// If the operation completes with errors, it returns a [partialFailure]
// error. The error stores individual errors in its [userInfo] dictionary. Use
// the [CKPartialErrorsByItemIDKey] key to extract them.
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
// See: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
//
// [CKPartialErrorsByItemIDKey]: https://developer.apple.com/documentation/CloudKit/CKPartialErrorsByItemIDKey
// [acceptSharesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/acceptSharesCompletionBlock
// [partialFailure]: https://developer.apple.com/documentation/CloudKit/CKError/partialFailure
// [perShareCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation/perShareCompletionBlock
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
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
// value to the [CKAcceptSharesOperation.ShareMetadatas] property before you
// execute the operation.
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
// value to the [CKAcceptSharesOperation.ShareMetadatas] property before you
// execute the operation.
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
