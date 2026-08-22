// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation] class.
var (
	_SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass     SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass
	_SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass {
	_SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass = SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass{class: objc.GetClass("SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation")}
	})
	return _SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass
}

// GetSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass returns the class object for SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.
func GetSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass {
	return getSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass()
}

type SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass) Alloc() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.MakeResultWithBoolValue]
//   - [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.WindowID]
//   - [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.InitWithDisplayIdentifierWindowID]
type SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationFromID constructs a [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation] from an objc.ID.
func SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationFromID(id objc.ID) SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	return SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation implements ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.
var _ ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation = SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation{}

// An interface definition for the [SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.MakeResultWithBoolValue]
//   - [ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.WindowID]
//   - [ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation.InitWithDisplayIdentifierWindowID]
type ISLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	MakeResultWithBoolValue(value bool) objectivec.IObject
	WindowID() uint32
	InitWithDisplayIdentifierWindowID(identifier objectivec.IObject, id uint32) SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) Init() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) Autorelease() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation creates a new SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation instance.
func NewSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation() SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	class := getSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass()
	rv := objc.SendIfResponds[SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	instance := getSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationFromID(rv)
}

func NewSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationWithDisplayIdentifierWindowID(identifier objectivec.IObject, id uint32) SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	instance := getSLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:windowID:"), identifier, id)
	return SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperationFromID(rv)
}

func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) MakeResultWithBoolValue(value bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithBoolValue:"), value)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) InitWithDisplayIdentifierWindowID(identifier objectivec.IObject, id uint32) SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation](s.ID, objc.Sel("initWithDisplayIdentifier:windowID:"), identifier, id)
	return rv
}

func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) DisplayIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSBridgedManagedDisplayCurrentSpaceAllowsWindowOperation) WindowID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
