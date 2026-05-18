// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineSendChangesOptions] class.
var (
	_CKSyncEngineSendChangesOptionsClass     CKSyncEngineSendChangesOptionsClass
	_CKSyncEngineSendChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineSendChangesOptionsClass() CKSyncEngineSendChangesOptionsClass {
	_CKSyncEngineSendChangesOptionsClassOnce.Do(func() {
		_CKSyncEngineSendChangesOptionsClass = CKSyncEngineSendChangesOptionsClass{class: objc.GetClass("CKSyncEngineSendChangesOptions")}
	})
	return _CKSyncEngineSendChangesOptionsClass
}

// GetCKSyncEngineSendChangesOptionsClass returns the class object for CKSyncEngineSendChangesOptions.
func GetCKSyncEngineSendChangesOptionsClass() CKSyncEngineSendChangesOptionsClass {
	return getCKSyncEngineSendChangesOptionsClass()
}

type CKSyncEngineSendChangesOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineSendChangesOptionsClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineSendChangesOptionsClass) Alloc() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A set of options to use with a send operation.
//
// # Managing attributes
//
//   - [CKSyncEngineSendChangesOptions.OperationGroup]: The operation group to use for the underlying CloudKit operations.
//   - [CKSyncEngineSendChangesOptions.SetOperationGroup]
//
// # Instance Properties
//
//   - [CKSyncEngineSendChangesOptions.Scope]
//   - [CKSyncEngineSendChangesOptions.SetScope]
//
// # Instance Methods
//
//   - [CKSyncEngineSendChangesOptions.InitWithScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions
type CKSyncEngineSendChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineSendChangesOptionsFromID constructs a [CKSyncEngineSendChangesOptions] from an objc.ID.
//
// A set of options to use with a send operation.
func CKSyncEngineSendChangesOptionsFromID(id objc.ID) CKSyncEngineSendChangesOptions {
	return CKSyncEngineSendChangesOptions{objectivec.Object{ID: id}}
}

// NOTE: CKSyncEngineSendChangesOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSyncEngineSendChangesOptions] class.
//
// # Managing attributes
//
//   - [ICKSyncEngineSendChangesOptions.OperationGroup]: The operation group to use for the underlying CloudKit operations.
//   - [ICKSyncEngineSendChangesOptions.SetOperationGroup]
//
// # Instance Properties
//
//   - [ICKSyncEngineSendChangesOptions.Scope]
//   - [ICKSyncEngineSendChangesOptions.SetScope]
//
// # Instance Methods
//
//   - [ICKSyncEngineSendChangesOptions.InitWithScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions
type ICKSyncEngineSendChangesOptions interface {
	objectivec.IObject

	// Topic: Managing attributes

	// The operation group to use for the underlying CloudKit operations.
	OperationGroup() ICKOperationGroup
	SetOperationGroup(value ICKOperationGroup)

	// Topic: Instance Properties

	Scope() ICKSyncEngineSendChangesScope
	SetScope(value ICKSyncEngineSendChangesScope)

	// Topic: Instance Methods

	InitWithScope(scope ICKSyncEngineSendChangesScope) CKSyncEngineSendChangesOptions
}

// Init initializes the instance.
func (c CKSyncEngineSendChangesOptions) Init() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineSendChangesOptions) Autorelease() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesOptions creates a new CKSyncEngineSendChangesOptions instance.
func NewCKSyncEngineSendChangesOptions() CKSyncEngineSendChangesOptions {
	class := getCKSyncEngineSendChangesOptionsClass()
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/initWithScope:
func NewCKSyncEngineSendChangesOptionsWithScope(scope ICKSyncEngineSendChangesScope) CKSyncEngineSendChangesOptions {
	instance := getCKSyncEngineSendChangesOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScope:"), scope)
	return CKSyncEngineSendChangesOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/initWithScope:
func (c CKSyncEngineSendChangesOptions) InitWithScope(scope ICKSyncEngineSendChangesScope) CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c.ID, objc.Sel("initWithScope:"), scope)
	return rv
}

// The operation group to use for the underlying CloudKit operations.
//
// # Discussion
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/operationGroup
func (c CKSyncEngineSendChangesOptions) OperationGroup() ICKOperationGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("operationGroup"))
	return CKOperationGroupFromID(objc.ID(rv))
}
func (c CKSyncEngineSendChangesOptions) SetOperationGroup(value ICKOperationGroup) {
	objc.Send[struct{}](c.ID, objc.Sel("setOperationGroup:"), value)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/scope
func (c CKSyncEngineSendChangesOptions) Scope() ICKSyncEngineSendChangesScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("scope"))
	return CKSyncEngineSendChangesScopeFromID(objc.ID(rv))
}
func (c CKSyncEngineSendChangesOptions) SetScope(value ICKSyncEngineSendChangesScope) {
	objc.Send[struct{}](c.ID, objc.Sel("setScope:"), value)
}
