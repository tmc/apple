// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKOperationGroup] class.
var (
	_CKOperationGroupClass     CKOperationGroupClass
	_CKOperationGroupClassOnce sync.Once
)

func getCKOperationGroupClass() CKOperationGroupClass {
	_CKOperationGroupClassOnce.Do(func() {
		_CKOperationGroupClass = CKOperationGroupClass{class: objc.GetClass("CKOperationGroup")}
	})
	return _CKOperationGroupClass
}

// GetCKOperationGroupClass returns the class object for CKOperationGroup.
func GetCKOperationGroupClass() CKOperationGroupClass {
	return getCKOperationGroupClass()
}

type CKOperationGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKOperationGroupClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKOperationGroupClass) Alloc() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An explicit association between two or more operations.
//
// # Overview
//
// In certain situations, you might want to perform several CloudKit
// operations together. Grouping operations in CloudKit doesn’t ensure
// atomicity.
//
// For example, when building a Calendar app, you group the following actions:
//
// - Fetch records from CloudKit, which consists of numerous queries that
// fetch both new records and records with changes. - Perform incremental
// fetches of records in response to a push notification. - Update several
// records when the user saves a calendar event.
//
// Associate operation groups with operations by setting their
// [CKOperation.Group] property. Create a new operation group for each
// distinct user interaction.
//
// # Creating an Operation Group
//
//   - [CKOperationGroup.InitWithCoder]: Creates an operation group from a serialized instance.
//
// # Configuring an Operation Group
//
//   - [CKOperationGroup.DefaultConfiguration]: The default configuration for operations in the group.
//   - [CKOperationGroup.SetDefaultConfiguration]
//   - [CKOperationGroup.ExpectedReceiveSize]: The estimated size of traffic to download from CloudKit.
//   - [CKOperationGroup.SetExpectedReceiveSize]
//   - [CKOperationGroup.ExpectedSendSize]: The estimated size of traffic to upload to CloudKit.
//   - [CKOperationGroup.SetExpectedSendSize]
//   - [CKOperationGroup.Name]: The operation group’s name.
//   - [CKOperationGroup.SetName]
//   - [CKOperationGroup.OperationGroupID]: The operation group’s unique identifier.
//   - [CKOperationGroup.Quantity]: The number of operations in the operation group.
//   - [CKOperationGroup.SetQuantity]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup
type CKOperationGroup struct {
	objectivec.Object
}

// CKOperationGroupFromID constructs a [CKOperationGroup] from an objc.ID.
//
// An explicit association between two or more operations.
func CKOperationGroupFromID(id objc.ID) CKOperationGroup {
	return CKOperationGroup{objectivec.Object{ID: id}}
}

// NOTE: CKOperationGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKOperationGroup] class.
//
// # Creating an Operation Group
//
//   - [ICKOperationGroup.InitWithCoder]: Creates an operation group from a serialized instance.
//
// # Configuring an Operation Group
//
//   - [ICKOperationGroup.DefaultConfiguration]: The default configuration for operations in the group.
//   - [ICKOperationGroup.SetDefaultConfiguration]
//   - [ICKOperationGroup.ExpectedReceiveSize]: The estimated size of traffic to download from CloudKit.
//   - [ICKOperationGroup.SetExpectedReceiveSize]
//   - [ICKOperationGroup.ExpectedSendSize]: The estimated size of traffic to upload to CloudKit.
//   - [ICKOperationGroup.SetExpectedSendSize]
//   - [ICKOperationGroup.Name]: The operation group’s name.
//   - [ICKOperationGroup.SetName]
//   - [ICKOperationGroup.OperationGroupID]: The operation group’s unique identifier.
//   - [ICKOperationGroup.Quantity]: The number of operations in the operation group.
//   - [ICKOperationGroup.SetQuantity]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup
type ICKOperationGroup interface {
	objectivec.IObject

	// Topic: Creating an Operation Group

	// Creates an operation group from a serialized instance.
	InitWithCoder(aDecoder foundation.INSCoder) CKOperationGroup

	// Topic: Configuring an Operation Group

	// The default configuration for operations in the group.
	DefaultConfiguration() ICKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
	// The estimated size of traffic to download from CloudKit.
	ExpectedReceiveSize() CKOperationGroupTransferSize
	SetExpectedReceiveSize(value CKOperationGroupTransferSize)
	// The estimated size of traffic to upload to CloudKit.
	ExpectedSendSize() CKOperationGroupTransferSize
	SetExpectedSendSize(value CKOperationGroupTransferSize)
	// The operation group’s name.
	Name() string
	SetName(value string)
	// The operation group’s unique identifier.
	OperationGroupID() string
	// The number of operations in the operation group.
	Quantity() uint
	SetQuantity(value uint)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKOperationGroup) Init() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKOperationGroup) Autorelease() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperationGroup creates a new CKOperationGroup instance.
func NewCKOperationGroup() CKOperationGroup {
	class := getCKOperationGroupClass()
	rv := objc.Send[CKOperationGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation group from a serialized instance.
//
// aDecoder: The coder to use when deserializing the group.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/init(coder:)
func NewCKOperationGroupWithCoder(aDecoder foundation.INSCoder) CKOperationGroup {
	instance := getCKOperationGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKOperationGroupFromID(rv)
}

// Creates an operation group from a serialized instance.
//
// aDecoder: The coder to use when deserializing the group.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/init(coder:)
func (c CKOperationGroup) InitWithCoder(aDecoder foundation.INSCoder) CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}
func (c CKOperationGroup) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The default configuration for operations in the group.
//
// # Discussion
//
// If an operation in the group has its own configuration, that
// configuration’s values override the default configuration’s values. For
// more information, see [CKOperationConfiguration].
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/defaultConfiguration
func (c CKOperationGroup) DefaultConfiguration() ICKOperationConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("defaultConfiguration"))
	return CKOperationConfigurationFromID(objc.ID(rv))
}
func (c CKOperationGroup) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setDefaultConfiguration:"), value)
}

// The estimated size of traffic to download from CloudKit.
//
// # Discussion
//
// This property informs the system about the amount of data your app can
// transfer. An order-of-magnitude estimate is better than no estimate, and
// accuracy helps performance. The system checks this value when it schedules
// discretionary network requests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedReceiveSize
func (c CKOperationGroup) ExpectedReceiveSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c.ID, objc.Sel("expectedReceiveSize"))
	return CKOperationGroupTransferSize(rv)
}
func (c CKOperationGroup) SetExpectedReceiveSize(value CKOperationGroupTransferSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setExpectedReceiveSize:"), value)
}

// The estimated size of traffic to upload to CloudKit.
//
// # Discussion
//
// This property informs the system about the amount of data your app can
// transfer. An order-of-magnitude estimate is better than no estimate, and
// accuracy helps performance. The system checks this value when it schedules
// discretionary network requests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedSendSize
func (c CKOperationGroup) ExpectedSendSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c.ID, objc.Sel("expectedSendSize"))
	return CKOperationGroupTransferSize(rv)
}
func (c CKOperationGroup) SetExpectedSendSize(value CKOperationGroupTransferSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setExpectedSendSize:"), value)
}

// The operation group’s name.
//
// # Discussion
//
// The system sends the name of the operation group to CloudKit to provide
// aggregate reporting for [CKOperationGroup]. The name must not include any
// personal data.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/name
func (c CKOperationGroup) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKOperationGroup) SetName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setName:"), objc.String(value))
}

// The operation group’s unique identifier.
//
// # Discussion
//
// The framework generates this value and it’s unique to this operation
// group. The system sends this identifier to CloudKit, which can use it to
// identify server-side logs for [CKOperationGroup].
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/operationGroupID
func (c CKOperationGroup) OperationGroupID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("operationGroupID"))
	return foundation.NSStringFromID(rv).String()
}

// The number of operations in the operation group.
//
// # Discussion
//
// This property shows the number of operations that you expect to be in this
// operation group. It’s the developer’s responsibility to set this value.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/quantity
func (c CKOperationGroup) Quantity() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("quantity"))
	return rv
}
func (c CKOperationGroup) SetQuantity(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setQuantity:"), value)
}
