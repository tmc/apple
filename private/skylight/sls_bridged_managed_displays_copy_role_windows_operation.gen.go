// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplaysCopyRoleWindowsOperation] class.
var (
	_SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass     SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass
	_SLSBridgedManagedDisplaysCopyRoleWindowsOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass() SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass {
	_SLSBridgedManagedDisplaysCopyRoleWindowsOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass = SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass{class: objc.GetClass("SLSBridgedManagedDisplaysCopyRoleWindowsOperation")}
	})
	return _SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass
}

// GetSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass returns the class object for SLSBridgedManagedDisplaysCopyRoleWindowsOperation.
func GetSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass() SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass {
	return getSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass()
}

type SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplaysCopyRoleWindowsOperationClass) Alloc() SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	rv := objc.Send[SLSBridgedManagedDisplaysCopyRoleWindowsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplaysCopyRoleWindowsOperation.DisplayIdentifiers]
//   - [SLSBridgedManagedDisplaysCopyRoleWindowsOperation.MakeResultWithNumbers]
//   - [SLSBridgedManagedDisplaysCopyRoleWindowsOperation.Role]
//   - [SLSBridgedManagedDisplaysCopyRoleWindowsOperation.InitWithDisplayIdentifiersRole]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation
type SLSBridgedManagedDisplaysCopyRoleWindowsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplaysCopyRoleWindowsOperationFromID constructs a [SLSBridgedManagedDisplaysCopyRoleWindowsOperation] from an objc.ID.
func SLSBridgedManagedDisplaysCopyRoleWindowsOperationFromID(id objc.ID) SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	return SLSBridgedManagedDisplaysCopyRoleWindowsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplaysCopyRoleWindowsOperation implements ISLSBridgedManagedDisplaysCopyRoleWindowsOperation.
var _ ISLSBridgedManagedDisplaysCopyRoleWindowsOperation = SLSBridgedManagedDisplaysCopyRoleWindowsOperation{}

// An interface definition for the [SLSBridgedManagedDisplaysCopyRoleWindowsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplaysCopyRoleWindowsOperation.DisplayIdentifiers]
//   - [ISLSBridgedManagedDisplaysCopyRoleWindowsOperation.MakeResultWithNumbers]
//   - [ISLSBridgedManagedDisplaysCopyRoleWindowsOperation.Role]
//   - [ISLSBridgedManagedDisplaysCopyRoleWindowsOperation.InitWithDisplayIdentifiersRole]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation
type ISLSBridgedManagedDisplaysCopyRoleWindowsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifiers() foundation.INSArray
	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Role() uint64
	InitWithDisplayIdentifiersRole(identifiers objectivec.IObject, role uint64) SLSBridgedManagedDisplaysCopyRoleWindowsOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) Init() SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	rv := objc.Send[SLSBridgedManagedDisplaysCopyRoleWindowsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) Autorelease() SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	rv := objc.Send[SLSBridgedManagedDisplaysCopyRoleWindowsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplaysCopyRoleWindowsOperation creates a new SLSBridgedManagedDisplaysCopyRoleWindowsOperation instance.
func NewSLSBridgedManagedDisplaysCopyRoleWindowsOperation() SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	class := getSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplaysCopyRoleWindowsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/initWithCoder:
func NewSLSBridgedManagedDisplaysCopyRoleWindowsOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	instance := getSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplaysCopyRoleWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/initWithDisplayIdentifiers:role:
func NewSLSBridgedManagedDisplaysCopyRoleWindowsOperationWithDisplayIdentifiersRole(identifiers objectivec.IObject, role uint64) SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	instance := getSLSBridgedManagedDisplaysCopyRoleWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifiers:role:"), identifiers, role)
	return SLSBridgedManagedDisplaysCopyRoleWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/makeResultWithNumbers:
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/initWithDisplayIdentifiers:role:
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) InitWithDisplayIdentifiersRole(identifiers objectivec.IObject, role uint64) SLSBridgedManagedDisplaysCopyRoleWindowsOperation {
	rv := objc.Send[SLSBridgedManagedDisplaysCopyRoleWindowsOperation](s.ID, objc.Sel("initWithDisplayIdentifiers:role:"), identifiers, role)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/displayIdentifiers
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) DisplayIdentifiers() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifiers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaysCopyRoleWindowsOperation/role
func (s SLSBridgedManagedDisplaysCopyRoleWindowsOperation) Role() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("role"))
	return rv
}
