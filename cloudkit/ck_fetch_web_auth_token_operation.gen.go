// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKFetchWebAuthTokenOperation] class.
var (
	_CKFetchWebAuthTokenOperationClass     CKFetchWebAuthTokenOperationClass
	_CKFetchWebAuthTokenOperationClassOnce sync.Once
)

func getCKFetchWebAuthTokenOperationClass() CKFetchWebAuthTokenOperationClass {
	_CKFetchWebAuthTokenOperationClassOnce.Do(func() {
		_CKFetchWebAuthTokenOperationClass = CKFetchWebAuthTokenOperationClass{class: objc.GetClass("CKFetchWebAuthTokenOperation")}
	})
	return _CKFetchWebAuthTokenOperationClass
}

// GetCKFetchWebAuthTokenOperationClass returns the class object for CKFetchWebAuthTokenOperation.
func GetCKFetchWebAuthTokenOperationClass() CKFetchWebAuthTokenOperationClass {
	return getCKFetchWebAuthTokenOperationClass()
}

type CKFetchWebAuthTokenOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchWebAuthTokenOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchWebAuthTokenOperationClass) Alloc() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that creates an authentication token for use with CloudKit web
// services.
//
// # Overview
//
// CloudKit web services provides an HTTP interface to fetch, create, update,
// and delete records, zones, and subscriptions. Each request you send
// requires an API token, which you configure in [CloudKit Dashboard]. You
// must create an API token for each container in each environment.
//
// If you want to send a request to an endpoint that requires an authenticated
// user, use this operation to fetch an authentication token. Append the
// authentication token, along with the API token, to the endpoint’s URL.
// That request then acts on behalf of the current user. Authentication tokens
// are short-lived and expire after a single use.
//
// For an example of using a web authentication token with a CloudKit web
// service, see [Changing Access Controls on User Data].
//
// This operation executes the handlers you provide on a background queue.
// Tasks that need access to the main queue must redirect as appropriate.
//
// The operation calls [fetchWebAuthTokenCompletionBlock] after it executes to
// provide the fetched token. Use the completion handler to perform
// housekeeping tasks for the operation. It should also manage any failures,
// whether due to an error or an explicit cancellation.
//
// CloudKit operations have a default QoS of [QualityOfService.default].
// Operations with this service level are discretionary. The system schedules
// their execution at an optimal time according to battery level and network
// conditions, among other factors. Use the [CKFetchWebAuthTokenOperation.QualityOfService] property to set
// a more appropriate QoS for the operation.
//
// The following example shows how to create the operation, configure its
// callbacks, and execute it in the user’s private database:
//
// # Creating a Fetch Token Operation
//
//   - [CKFetchWebAuthTokenOperation.InitWithAPIToken]: Creates a fetch operation for the specified API token.
//
// # Managing the Operation’s Configuration
//
//   - [CKFetchWebAuthTokenOperation.APIToken]: The API token that allows access to an app’s container.
//   - [CKFetchWebAuthTokenOperation.SetAPIToken]
//
// # Instance Properties
//
//   - [CKFetchWebAuthTokenOperation.FetchWebAuthTokenResultBlock]: The closure to execute when the operation finishes.
//   - [CKFetchWebAuthTokenOperation.SetFetchWebAuthTokenResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation
//
// [Changing Access Controls on User Data]: https://developer.apple.com/documentation/CloudKit/changing-access-controls-on-user-data
// [CloudKit Dashboard]: https://icloud.developer.apple.com
// [QualityOfService.default]: https://developer.apple.com/documentation/Foundation/QualityOfService/default
// [fetchWebAuthTokenCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/fetchWebAuthTokenCompletionBlock
type CKFetchWebAuthTokenOperation struct {
	CKDatabaseOperation
}

// CKFetchWebAuthTokenOperationFromID constructs a [CKFetchWebAuthTokenOperation] from an objc.ID.
//
// An operation that creates an authentication token for use with CloudKit web
// services.
func CKFetchWebAuthTokenOperationFromID(id objc.ID) CKFetchWebAuthTokenOperation {
	return CKFetchWebAuthTokenOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchWebAuthTokenOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchWebAuthTokenOperation] class.
//
// # Creating a Fetch Token Operation
//
//   - [ICKFetchWebAuthTokenOperation.InitWithAPIToken]: Creates a fetch operation for the specified API token.
//
// # Managing the Operation’s Configuration
//
//   - [ICKFetchWebAuthTokenOperation.APIToken]: The API token that allows access to an app’s container.
//   - [ICKFetchWebAuthTokenOperation.SetAPIToken]
//
// # Instance Properties
//
//   - [ICKFetchWebAuthTokenOperation.FetchWebAuthTokenResultBlock]: The closure to execute when the operation finishes.
//   - [ICKFetchWebAuthTokenOperation.SetFetchWebAuthTokenResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation
type ICKFetchWebAuthTokenOperation interface {
	ICKDatabaseOperation

	// Topic: Creating a Fetch Token Operation

	// Creates a fetch operation for the specified API token.
	InitWithAPIToken(APIToken string) CKFetchWebAuthTokenOperation

	// Topic: Managing the Operation’s Configuration

	// The API token that allows access to an app’s container.
	APIToken() string
	SetAPIToken(value string)

	// Topic: Instance Properties

	// The closure to execute when the operation finishes.
	FetchWebAuthTokenResultBlock() unsafe.Pointer
	SetFetchWebAuthTokenResultBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKFetchWebAuthTokenOperation) Init() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchWebAuthTokenOperation) Autorelease() CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchWebAuthTokenOperation creates a new CKFetchWebAuthTokenOperation instance.
func NewCKFetchWebAuthTokenOperation() CKFetchWebAuthTokenOperation {
	class := getCKFetchWebAuthTokenOperationClass()
	rv := objc.Send[CKFetchWebAuthTokenOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fetch operation for the specified API token.
//
// APIToken: The API token that allows access to an app’s container.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/init(apiToken:)-14712
func NewCKFetchWebAuthTokenOperationWithAPIToken(APIToken string) CKFetchWebAuthTokenOperation {
	instance := getCKFetchWebAuthTokenOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAPIToken:"), objc.String(APIToken))
	return CKFetchWebAuthTokenOperationFromID(rv)
}

// Creates a fetch operation for the specified API token.
//
// APIToken: The API token that allows access to an app’s container.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/init(apiToken:)-14712
func (c CKFetchWebAuthTokenOperation) InitWithAPIToken(APIToken string) CKFetchWebAuthTokenOperation {
	rv := objc.Send[CKFetchWebAuthTokenOperation](c.ID, objc.Sel("initWithAPIToken:"), objc.String(APIToken))
	return rv
}

// The API token that allows access to an app’s container.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchWebAuthTokenOperation/apiToken
func (c CKFetchWebAuthTokenOperation) APIToken() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("APIToken"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchWebAuthTokenOperation) SetAPIToken(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAPIToken:"), objc.String(value))
}

// The closure to execute when the operation finishes.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/fetchwebauthtokenresultblock
func (c CKFetchWebAuthTokenOperation) FetchWebAuthTokenResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchWebAuthTokenResultBlock"))
	return rv
}
func (c CKFetchWebAuthTokenOperation) SetFetchWebAuthTokenResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchWebAuthTokenResultBlock:"), value)
}
