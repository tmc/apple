// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKOperationConfiguration] class.
var (
	_CKOperationConfigurationClass     CKOperationConfigurationClass
	_CKOperationConfigurationClassOnce sync.Once
)

func getCKOperationConfigurationClass() CKOperationConfigurationClass {
	_CKOperationConfigurationClassOnce.Do(func() {
		_CKOperationConfigurationClass = CKOperationConfigurationClass{class: objc.GetClass("CKOperationConfiguration")}
	})
	return _CKOperationConfigurationClass
}

// GetCKOperationConfigurationClass returns the class object for CKOperationConfiguration.
func GetCKOperationConfigurationClass() CKOperationConfigurationClass {
	return getCKOperationConfigurationClass()
}

type CKOperationConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKOperationConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKOperationConfigurationClass) Alloc() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes how a CloudKit operation behaves.
//
// # Overview
//
// All of the properties in [CKOperationConfiguration] have a default value.
// When determining which properties to apply to a CloudKit operation, consult
// the operation’s configuration property, as well as the
// [CKOperationConfiguration.DefaultConfiguration] property of the group that the operation belongs to.
// These properties combine through the following rules:
//
// [Table data omitted]
//
// # Preparing a Configuration
//
//   - [CKOperationConfiguration.AllowsCellularAccess]: A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//   - [CKOperationConfiguration.SetAllowsCellularAccess]
//   - [CKOperationConfiguration.Container]: The configuration’s container.
//   - [CKOperationConfiguration.SetContainer]
//   - [CKOperationConfiguration.LongLived]: A Boolean value that indicates whether the operations that use this configuration are long-lived.
//   - [CKOperationConfiguration.SetLongLived]
//   - [CKOperationConfiguration.QualityOfService]: The priority that the system uses when it allocates resources to the operations that use this configuration.
//   - [CKOperationConfiguration.SetQualityOfService]
//   - [CKOperationConfiguration.TimeoutIntervalForRequest]: The maximum amount of time that a request can take.
//   - [CKOperationConfiguration.SetTimeoutIntervalForRequest]
//   - [CKOperationConfiguration.TimeoutIntervalForResource]: The maximum amount of time that a resource request can take.
//   - [CKOperationConfiguration.SetTimeoutIntervalForResource]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class
type CKOperationConfiguration struct {
	objectivec.Object
}

// CKOperationConfigurationFromID constructs a [CKOperationConfiguration] from an objc.ID.
//
// An object that describes how a CloudKit operation behaves.
func CKOperationConfigurationFromID(id objc.ID) CKOperationConfiguration {
	return CKOperationConfiguration{objectivec.Object{ID: id}}
}

// NOTE: CKOperationConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKOperationConfiguration] class.
//
// # Preparing a Configuration
//
//   - [ICKOperationConfiguration.AllowsCellularAccess]: A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//   - [ICKOperationConfiguration.SetAllowsCellularAccess]
//   - [ICKOperationConfiguration.Container]: The configuration’s container.
//   - [ICKOperationConfiguration.SetContainer]
//   - [ICKOperationConfiguration.LongLived]: A Boolean value that indicates whether the operations that use this configuration are long-lived.
//   - [ICKOperationConfiguration.SetLongLived]
//   - [ICKOperationConfiguration.QualityOfService]: The priority that the system uses when it allocates resources to the operations that use this configuration.
//   - [ICKOperationConfiguration.SetQualityOfService]
//   - [ICKOperationConfiguration.TimeoutIntervalForRequest]: The maximum amount of time that a request can take.
//   - [ICKOperationConfiguration.SetTimeoutIntervalForRequest]
//   - [ICKOperationConfiguration.TimeoutIntervalForResource]: The maximum amount of time that a resource request can take.
//   - [ICKOperationConfiguration.SetTimeoutIntervalForResource]
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class
type ICKOperationConfiguration interface {
	objectivec.IObject

	// Topic: Preparing a Configuration

	// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	// The configuration’s container.
	Container() ICKContainer
	SetContainer(value ICKContainer)
	// A Boolean value that indicates whether the operations that use this configuration are long-lived.
	LongLived() bool
	SetLongLived(value bool)
	// The priority that the system uses when it allocates resources to the operations that use this configuration.
	QualityOfService() foundation.NSQualityOfService
	SetQualityOfService(value foundation.NSQualityOfService)
	// The maximum amount of time that a request can take.
	TimeoutIntervalForRequest() float64
	SetTimeoutIntervalForRequest(value float64)
	// The maximum amount of time that a resource request can take.
	TimeoutIntervalForResource() float64
	SetTimeoutIntervalForResource(value float64)

	// The operation’s configuration.
	Configuration() ICKOperationConfiguration
	SetConfiguration(value ICKOperationConfiguration)
	// The default configuration for operations in the group.
	DefaultConfiguration() ICKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
	// The operation’s group.
	Group() ICKOperationGroup
	SetGroup(value ICKOperationGroup)
	// The closure to execute when the server begins to store callbacks for the long-lived operation.
	LongLivedOperationWasPersistedBlock() unsafe.Pointer
	SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKOperationConfiguration) Init() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKOperationConfiguration) Autorelease() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperationConfiguration creates a new CKOperationConfiguration instance.
func NewCKOperationConfiguration() CKOperationConfiguration {
	class := getCKOperationConfigurationClass()
	rv := objc.Send[CKOperationConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether operations that use this
// configuration can send data over the cellular network.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/allowsCellularAccess
func (c CKOperationConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsCellularAccess"))
	return rv
}
func (c CKOperationConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// The configuration’s container.
//
// # Discussion
//
// If you don’t provide a container, CloudKit uses the default container
// that [CKContainer] provides.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/container
func (c CKOperationConfiguration) Container() ICKContainer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("container"))
	return CKContainerFromID(objc.ID(rv))
}
func (c CKOperationConfiguration) SetContainer(value ICKContainer) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainer:"), value)
}

// A Boolean value that indicates whether the operations that use this
// configuration are long-lived.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/isLongLived
func (c CKOperationConfiguration) LongLived() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLongLived"))
	return rv
}
func (c CKOperationConfiguration) SetLongLived(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setLongLived:"), value)
}

// The priority that the system uses when it allocates resources to the
// operations that use this configuration.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c CKOperationConfiguration) QualityOfService() foundation.NSQualityOfService {
	rv := objc.Send[foundation.NSQualityOfService](c.ID, objc.Sel("qualityOfService"))
	return foundation.NSQualityOfService(rv)
}
func (c CKOperationConfiguration) SetQualityOfService(value foundation.NSQualityOfService) {
	objc.Send[struct{}](c.ID, objc.Sel("setQualityOfService:"), value)
}

// The maximum amount of time that a request can take.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForRequest
func (c CKOperationConfiguration) TimeoutIntervalForRequest() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}
func (c CKOperationConfiguration) SetTimeoutIntervalForRequest(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// The maximum amount of time that a resource request can take.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForResource
func (c CKOperationConfiguration) TimeoutIntervalForResource() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}
func (c CKOperationConfiguration) SetTimeoutIntervalForResource(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}

// The operation’s configuration.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c CKOperationConfiguration) Configuration() ICKOperationConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return CKOperationConfigurationFromID(objc.ID(rv))
}
func (c CKOperationConfiguration) SetConfiguration(value ICKOperationConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfiguration:"), value)
}

// The default configuration for operations in the group.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c CKOperationConfiguration) DefaultConfiguration() ICKOperationConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("defaultConfiguration"))
	return CKOperationConfigurationFromID(objc.ID(rv))
}
func (c CKOperationConfiguration) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setDefaultConfiguration:"), value)
}

// The operation’s group.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c CKOperationConfiguration) Group() ICKOperationGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("group"))
	return CKOperationGroupFromID(objc.ID(rv))
}
func (c CKOperationConfiguration) SetGroup(value ICKOperationGroup) {
	objc.Send[struct{}](c.ID, objc.Sel("setGroup:"), value)
}

// The closure to execute when the server begins to store callbacks for the
// long-lived operation.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c CKOperationConfiguration) LongLivedOperationWasPersistedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}
func (c CKOperationConfiguration) SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}
