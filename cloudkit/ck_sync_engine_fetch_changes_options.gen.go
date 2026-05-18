// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFetchChangesOptions] class.
var (
	_CKSyncEngineFetchChangesOptionsClass     CKSyncEngineFetchChangesOptionsClass
	_CKSyncEngineFetchChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineFetchChangesOptionsClass() CKSyncEngineFetchChangesOptionsClass {
	_CKSyncEngineFetchChangesOptionsClassOnce.Do(func() {
		_CKSyncEngineFetchChangesOptionsClass = CKSyncEngineFetchChangesOptionsClass{class: objc.GetClass("CKSyncEngineFetchChangesOptions")}
	})
	return _CKSyncEngineFetchChangesOptionsClass
}

// GetCKSyncEngineFetchChangesOptionsClass returns the class object for CKSyncEngineFetchChangesOptions.
func GetCKSyncEngineFetchChangesOptionsClass() CKSyncEngineFetchChangesOptionsClass {
	return getCKSyncEngineFetchChangesOptionsClass()
}

type CKSyncEngineFetchChangesOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchChangesOptionsClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchChangesOptionsClass) Alloc() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A set of options to use with a fetch operation.
//
// # Managing attributes
//
//   - [CKSyncEngineFetchChangesOptions.OperationGroup]: The operation group to use for the underlying CloudKit operations.
//   - [CKSyncEngineFetchChangesOptions.SetOperationGroup]
//
// # Instance Properties
//
//   - [CKSyncEngineFetchChangesOptions.PrioritizedZoneIDs]
//   - [CKSyncEngineFetchChangesOptions.SetPrioritizedZoneIDs]
//   - [CKSyncEngineFetchChangesOptions.Scope]
//   - [CKSyncEngineFetchChangesOptions.SetScope]
//
// # Instance Methods
//
//   - [CKSyncEngineFetchChangesOptions.InitWithScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions
type CKSyncEngineFetchChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesOptionsFromID constructs a [CKSyncEngineFetchChangesOptions] from an objc.ID.
//
// A set of options to use with a fetch operation.
func CKSyncEngineFetchChangesOptionsFromID(id objc.ID) CKSyncEngineFetchChangesOptions {
	return CKSyncEngineFetchChangesOptions{objectivec.Object{ID: id}}
}

// NOTE: CKSyncEngineFetchChangesOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSyncEngineFetchChangesOptions] class.
//
// # Managing attributes
//
//   - [ICKSyncEngineFetchChangesOptions.OperationGroup]: The operation group to use for the underlying CloudKit operations.
//   - [ICKSyncEngineFetchChangesOptions.SetOperationGroup]
//
// # Instance Properties
//
//   - [ICKSyncEngineFetchChangesOptions.PrioritizedZoneIDs]
//   - [ICKSyncEngineFetchChangesOptions.SetPrioritizedZoneIDs]
//   - [ICKSyncEngineFetchChangesOptions.Scope]
//   - [ICKSyncEngineFetchChangesOptions.SetScope]
//
// # Instance Methods
//
//   - [ICKSyncEngineFetchChangesOptions.InitWithScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions
type ICKSyncEngineFetchChangesOptions interface {
	objectivec.IObject

	// Topic: Managing attributes

	// The operation group to use for the underlying CloudKit operations.
	OperationGroup() ICKOperationGroup
	SetOperationGroup(value ICKOperationGroup)

	// Topic: Instance Properties

	PrioritizedZoneIDs() []CKRecordZoneID
	SetPrioritizedZoneIDs(value []CKRecordZoneID)
	Scope() ICKSyncEngineFetchChangesScope
	SetScope(value ICKSyncEngineFetchChangesScope)

	// Topic: Instance Methods

	InitWithScope(scope ICKSyncEngineFetchChangesScope) CKSyncEngineFetchChangesOptions
}

// Init initializes the instance.
func (c CKSyncEngineFetchChangesOptions) Init() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchChangesOptions) Autorelease() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesOptions creates a new CKSyncEngineFetchChangesOptions instance.
func NewCKSyncEngineFetchChangesOptions() CKSyncEngineFetchChangesOptions {
	class := getCKSyncEngineFetchChangesOptionsClass()
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/initWithScope:
func NewCKSyncEngineFetchChangesOptionsWithScope(scope ICKSyncEngineFetchChangesScope) CKSyncEngineFetchChangesOptions {
	instance := getCKSyncEngineFetchChangesOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScope:"), scope)
	return CKSyncEngineFetchChangesOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/initWithScope:
func (c CKSyncEngineFetchChangesOptions) InitWithScope(scope ICKSyncEngineFetchChangesScope) CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c.ID, objc.Sel("initWithScope:"), scope)
	return rv
}

// The operation group to use for the underlying CloudKit operations.
//
// # Discussion
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/operationGroup
func (c CKSyncEngineFetchChangesOptions) OperationGroup() ICKOperationGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("operationGroup"))
	return CKOperationGroupFromID(objc.ID(rv))
}
func (c CKSyncEngineFetchChangesOptions) SetOperationGroup(value ICKOperationGroup) {
	objc.Send[struct{}](c.ID, objc.Sel("setOperationGroup:"), value)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/prioritizedZoneIDs
func (c CKSyncEngineFetchChangesOptions) PrioritizedZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("prioritizedZoneIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}
func (c CKSyncEngineFetchChangesOptions) SetPrioritizedZoneIDs(value []CKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrioritizedZoneIDs:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/scope
func (c CKSyncEngineFetchChangesOptions) Scope() ICKSyncEngineFetchChangesScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scope"))
	return CKSyncEngineFetchChangesScopeFromID(objc.ID(rv))
}
func (c CKSyncEngineFetchChangesOptions) SetScope(value ICKSyncEngineFetchChangesScope) {
	objc.Send[struct{}](c.ID, objc.Sel("setScope:"), value)
}
