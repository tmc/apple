// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplaySetRoleWindowOperation] class.
var (
	_SLSBridgedManagedDisplaySetRoleWindowOperationClass     SLSBridgedManagedDisplaySetRoleWindowOperationClass
	_SLSBridgedManagedDisplaySetRoleWindowOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplaySetRoleWindowOperationClass() SLSBridgedManagedDisplaySetRoleWindowOperationClass {
	_SLSBridgedManagedDisplaySetRoleWindowOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplaySetRoleWindowOperationClass = SLSBridgedManagedDisplaySetRoleWindowOperationClass{class: objc.GetClass("SLSBridgedManagedDisplaySetRoleWindowOperation")}
	})
	return _SLSBridgedManagedDisplaySetRoleWindowOperationClass
}

// GetSLSBridgedManagedDisplaySetRoleWindowOperationClass returns the class object for SLSBridgedManagedDisplaySetRoleWindowOperation.
func GetSLSBridgedManagedDisplaySetRoleWindowOperationClass() SLSBridgedManagedDisplaySetRoleWindowOperationClass {
	return getSLSBridgedManagedDisplaySetRoleWindowOperationClass()
}

type SLSBridgedManagedDisplaySetRoleWindowOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplaySetRoleWindowOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplaySetRoleWindowOperationClass) Alloc() SLSBridgedManagedDisplaySetRoleWindowOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetRoleWindowOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplaySetRoleWindowOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplaySetRoleWindowOperation.Role]
//   - [SLSBridgedManagedDisplaySetRoleWindowOperation.WindowID]
//   - [SLSBridgedManagedDisplaySetRoleWindowOperation.InitWithDisplayIdentifierRoleWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation
type SLSBridgedManagedDisplaySetRoleWindowOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplaySetRoleWindowOperationFromID constructs a [SLSBridgedManagedDisplaySetRoleWindowOperation] from an objc.ID.
func SLSBridgedManagedDisplaySetRoleWindowOperationFromID(id objc.ID) SLSBridgedManagedDisplaySetRoleWindowOperation {
	return SLSBridgedManagedDisplaySetRoleWindowOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplaySetRoleWindowOperation implements ISLSBridgedManagedDisplaySetRoleWindowOperation.
var _ ISLSBridgedManagedDisplaySetRoleWindowOperation = SLSBridgedManagedDisplaySetRoleWindowOperation{}

// An interface definition for the [SLSBridgedManagedDisplaySetRoleWindowOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplaySetRoleWindowOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplaySetRoleWindowOperation.Role]
//   - [ISLSBridgedManagedDisplaySetRoleWindowOperation.WindowID]
//   - [ISLSBridgedManagedDisplaySetRoleWindowOperation.InitWithDisplayIdentifierRoleWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation
type ISLSBridgedManagedDisplaySetRoleWindowOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	Role() uint64
	WindowID() uint32
	InitWithDisplayIdentifierRoleWindowID(identifier objectivec.IObject, role uint64, id uint32) SLSBridgedManagedDisplaySetRoleWindowOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) Init() SLSBridgedManagedDisplaySetRoleWindowOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetRoleWindowOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) Autorelease() SLSBridgedManagedDisplaySetRoleWindowOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetRoleWindowOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplaySetRoleWindowOperation creates a new SLSBridgedManagedDisplaySetRoleWindowOperation instance.
func NewSLSBridgedManagedDisplaySetRoleWindowOperation() SLSBridgedManagedDisplaySetRoleWindowOperation {
	class := getSLSBridgedManagedDisplaySetRoleWindowOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplaySetRoleWindowOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/initWithCoder:
func NewSLSBridgedManagedDisplaySetRoleWindowOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplaySetRoleWindowOperation {
	instance := getSLSBridgedManagedDisplaySetRoleWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplaySetRoleWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/initWithDisplayIdentifier:role:windowID:
func NewSLSBridgedManagedDisplaySetRoleWindowOperationWithDisplayIdentifierRoleWindowID(identifier objectivec.IObject, role uint64, id uint32) SLSBridgedManagedDisplaySetRoleWindowOperation {
	instance := getSLSBridgedManagedDisplaySetRoleWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:role:windowID:"), identifier, role, id)
	return SLSBridgedManagedDisplaySetRoleWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/initWithDisplayIdentifier:role:windowID:
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) InitWithDisplayIdentifierRoleWindowID(identifier objectivec.IObject, role uint64, id uint32) SLSBridgedManagedDisplaySetRoleWindowOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetRoleWindowOperation](s.ID, objc.Sel("initWithDisplayIdentifier:role:windowID:"), identifier, role, id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/displayIdentifier
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/role
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) Role() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("role"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetRoleWindowOperation/windowID
func (s SLSBridgedManagedDisplaySetRoleWindowOperation) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
