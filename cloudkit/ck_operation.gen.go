// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKOperation] class.
var (
	_CKOperationClass     CKOperationClass
	_CKOperationClassOnce sync.Once
)

func getCKOperationClass() CKOperationClass {
	_CKOperationClassOnce.Do(func() {
		_CKOperationClass = CKOperationClass{class: objc.GetClass("CKOperation")}
	})
	return _CKOperationClass
}

// GetCKOperationClass returns the class object for CKOperation.
func GetCKOperationClass() CKOperationClass {
	return getCKOperationClass()
}

type CKOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKOperationClass) Alloc() CKOperation {
	rv := objc.Send[CKOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for all operations that execute in a database.
//
// # Overview
//
// All CloudKit operations descend from [CKOperation], which provides the
// infrastructure for executing tasks in one of your app’s containers.
// Don’t subclass or create instances of this class directly. Instead,
// create instances of one of its concrete subclasses.
//
// Use the properties of this class to configure the behavior of the operation
// before submitting it to a queue or executing it directly. CloudKit
// operations involve communicating with the iCloud servers to send and
// receive data. You can use the properties of this class to configure the
// behavior of those network requests to ensure the best performance for your
// app.
//
// # Long-Lived Operations
//
// A is an operation that continues to run after the user closes the app. To
// specify a long-lived operation, set [CKOperation.IsLongLived] to true, provide a
// completion handler, and execute the operation. To get the identifiers of
// all running long-lived operations, use the [allLongLivedOperationIDs()]
// method that [CKContainer] provides. To get a specific long-lived operation,
// use the [longLivedOperation(for:)] method. Make sure you set the completion
// handler of a long-lived operation before you execute it so that the system
// can notify you when it completes and you can process the results. Do not
// execute an operation, change it to long-lived, and execute it again as a
// long-lived operation.
//
// The following is the typical life cycle of a long-lived operation:
//
// - The app creates a long-lived operation and executes it.
//
// The daemon starts saving and sending the callbacks to the running app. 2.
// The app exits.
//
// The daemon continues running the long-lived operation and saves the
// callbacks. 3. The app launches and fetches the long-lived operation.
//
// If the operation is running or if it completed within the previous 24
// hours, the daemon returns a proxy for the long-lived operation. If the
// operation completed more than 24 hours previously, the daemon may stop
// returning it in fetch requests. 4. The app runs the long-lived operation
// again.
//
// The daemon sends the app all the saved callbacks (it doesn’t actually
// rerun the operation), and continues saving the callbacks and sending them
// to the running app. 5. The app receives the completion callback or the app
// cancels the operation.
//
// The daemon stops including the operation in future fetch results.
//
// # Identifying the Operation
//
//   - [CKOperation.OperationID]: A unique identifier for a long-lived operation.
//   - [CKOperation.SetOperationID]
//
// # Managing the Operation’s Configuration
//
//   - [CKOperation.Configuration]: The operation’s configuration.
//   - [CKOperation.SetConfiguration]
//   - [CKOperation.Group]: The operation’s group.
//   - [CKOperation.SetGroup]
//   - [CKOperation.LongLivedOperationWasPersistedBlock]: The closure to execute when the server begins to store callbacks for the long-lived operation.
//   - [CKOperation.SetLongLivedOperationWasPersistedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation
//
// [allLongLivedOperationIDs()]: https://developer.apple.com/documentation/CloudKit/CKContainer/allLongLivedOperationIDs()
// [longLivedOperation(for:)]: https://developer.apple.com/documentation/CloudKit/CKContainer/longLivedOperation(for:)
type CKOperation struct {
	foundation.Operation
}

// CKOperationFromID constructs a [CKOperation] from an objc.ID.
//
// The abstract base class for all operations that execute in a database.
func CKOperationFromID(id objc.ID) CKOperation {
	return CKOperation{Operation: foundation.OperationFromID(id)}
}

// NOTE: CKOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKOperation] class.
//
// # Identifying the Operation
//
//   - [ICKOperation.OperationID]: A unique identifier for a long-lived operation.
//   - [ICKOperation.SetOperationID]
//
// # Managing the Operation’s Configuration
//
//   - [ICKOperation.Configuration]: The operation’s configuration.
//   - [ICKOperation.SetConfiguration]
//   - [ICKOperation.Group]: The operation’s group.
//   - [ICKOperation.SetGroup]
//   - [ICKOperation.LongLivedOperationWasPersistedBlock]: The closure to execute when the server begins to store callbacks for the long-lived operation.
//   - [ICKOperation.SetLongLivedOperationWasPersistedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation
type ICKOperation interface {
	foundation.IOperation

	// Topic: Identifying the Operation

	// A unique identifier for a long-lived operation.
	OperationID() string
	SetOperationID(value string)

	// Topic: Managing the Operation’s Configuration

	// The operation’s configuration.
	Configuration() ICKOperationConfiguration
	SetConfiguration(value ICKOperationConfiguration)
	// The operation’s group.
	Group() ICKOperationGroup
	SetGroup(value ICKOperationGroup)
	// The closure to execute when the server begins to store callbacks for the long-lived operation.
	LongLivedOperationWasPersistedBlock() VoidHandler
	SetLongLivedOperationWasPersistedBlock(value VoidHandler)

	// A Boolean value that indicates whether the operation can send data over the cellular network.
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	// The operation’s container.
	Container() ICKContainer
	SetContainer(value ICKContainer)
	// A Boolean value that indicates whether the operation is long-lived.
	IsLongLived() bool
	SetLongLived(value bool)
	// The timeout interval when waiting for additional data.
	TimeoutIntervalForRequest() float64
	SetTimeoutIntervalForRequest(value float64)
	// The maximum amount of time that a resource request can use.
	TimeoutIntervalForResource() float64
	SetTimeoutIntervalForResource(value float64)
}

// Init initializes the instance.
func (c CKOperation) Init() CKOperation {
	rv := objc.Send[CKOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKOperation) Autorelease() CKOperation {
	rv := objc.Send[CKOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperation creates a new CKOperation instance.
func NewCKOperation() CKOperation {
	class := getCKOperationClass()
	rv := objc.Send[CKOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A unique identifier for a long-lived operation.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/operationid-8auuc
func (c CKOperation) OperationID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("operationID"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKOperation) SetOperationID(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setOperationID:"), objc.String(value))
}

// The operation’s configuration.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/configuration-swift.property
func (c CKOperation) Configuration() ICKOperationConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return CKOperationConfigurationFromID(objc.ID(rv))
}
func (c CKOperation) SetConfiguration(value ICKOperationConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfiguration:"), value)
}

// The operation’s group.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/group
func (c CKOperation) Group() ICKOperationGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("group"))
	return CKOperationGroupFromID(objc.ID(rv))
}
func (c CKOperation) SetGroup(value ICKOperationGroup) {
	objc.Send[struct{}](c.ID, objc.Sel("setGroup:"), value)
}

// The closure to execute when the server begins to store callbacks for the
// long-lived operation.
//
// # Discussion
//
// If your app exits before CloudKit calls this property’s value, the system
// doesn’t include the operation’s ID in the results of calls to the
// [fetchAllLongLivedOperationIDsWithCompletionHandler:] method.
//
// For more information, see [CKOperation].
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/longLivedOperationWasPersistedBlock
//
// [fetchAllLongLivedOperationIDsWithCompletionHandler:]: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchAllLongLivedOperationIDsWithCompletionHandler:
func (c CKOperation) LongLivedOperationWasPersistedBlock() VoidHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	_ = rv
	return nil
}
func (c CKOperation) SetLongLivedOperationWasPersistedBlock(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), block)
}

// A Boolean value that indicates whether the operation can send data over the
// cellular network.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/allowscellularaccess
func (c CKOperation) AllowsCellularAccess() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsCellularAccess"))
	return rv
}
func (c CKOperation) SetAllowsCellularAccess(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// The operation’s container.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/container
func (c CKOperation) Container() ICKContainer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("container"))
	return CKContainerFromID(objc.ID(rv))
}
func (c CKOperation) SetContainer(value ICKContainer) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainer:"), value)
}

// A Boolean value that indicates whether the operation is long-lived.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/islonglived
func (c CKOperation) IsLongLived() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("longLived"))
	return rv
}
func (c CKOperation) SetLongLived(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setLongLived:"), value)
}

// The timeout interval when waiting for additional data.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/timeoutintervalforrequest
func (c CKOperation) TimeoutIntervalForRequest() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}
func (c CKOperation) SetTimeoutIntervalForRequest(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// The maximum amount of time that a resource request can use.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/timeoutintervalforresource
func (c CKOperation) TimeoutIntervalForResource() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}
func (c CKOperation) SetTimeoutIntervalForResource(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}
