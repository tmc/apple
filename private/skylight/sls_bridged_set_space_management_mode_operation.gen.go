// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSetSpaceManagementModeOperation] class.
var (
	_SLSBridgedSetSpaceManagementModeOperationClass     SLSBridgedSetSpaceManagementModeOperationClass
	_SLSBridgedSetSpaceManagementModeOperationClassOnce sync.Once
)

func getSLSBridgedSetSpaceManagementModeOperationClass() SLSBridgedSetSpaceManagementModeOperationClass {
	_SLSBridgedSetSpaceManagementModeOperationClassOnce.Do(func() {
		_SLSBridgedSetSpaceManagementModeOperationClass = SLSBridgedSetSpaceManagementModeOperationClass{class: objc.GetClass("SLSBridgedSetSpaceManagementModeOperation")}
	})
	return _SLSBridgedSetSpaceManagementModeOperationClass
}

// GetSLSBridgedSetSpaceManagementModeOperationClass returns the class object for SLSBridgedSetSpaceManagementModeOperation.
func GetSLSBridgedSetSpaceManagementModeOperationClass() SLSBridgedSetSpaceManagementModeOperationClass {
	return getSLSBridgedSetSpaceManagementModeOperationClass()
}

type SLSBridgedSetSpaceManagementModeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSetSpaceManagementModeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSetSpaceManagementModeOperationClass) Alloc() SLSBridgedSetSpaceManagementModeOperation {
	rv := objc.SendIfResponds[SLSBridgedSetSpaceManagementModeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSetSpaceManagementModeOperation.Mode]
//   - [SLSBridgedSetSpaceManagementModeOperation.InitWithMode]
type SLSBridgedSetSpaceManagementModeOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSetSpaceManagementModeOperationFromID constructs a [SLSBridgedSetSpaceManagementModeOperation] from an objc.ID.
func SLSBridgedSetSpaceManagementModeOperationFromID(id objc.ID) SLSBridgedSetSpaceManagementModeOperation {
	return SLSBridgedSetSpaceManagementModeOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSetSpaceManagementModeOperation implements ISLSBridgedSetSpaceManagementModeOperation.
var _ ISLSBridgedSetSpaceManagementModeOperation = SLSBridgedSetSpaceManagementModeOperation{}

// An interface definition for the [SLSBridgedSetSpaceManagementModeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSetSpaceManagementModeOperation.Mode]
//   - [ISLSBridgedSetSpaceManagementModeOperation.InitWithMode]
type ISLSBridgedSetSpaceManagementModeOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Mode() uint64
	InitWithMode(mode uint64) SLSBridgedSetSpaceManagementModeOperation
}

// Init initializes the instance.
func (s SLSBridgedSetSpaceManagementModeOperation) Init() SLSBridgedSetSpaceManagementModeOperation {
	rv := objc.SendIfResponds[SLSBridgedSetSpaceManagementModeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSetSpaceManagementModeOperation) Autorelease() SLSBridgedSetSpaceManagementModeOperation {
	rv := objc.SendIfResponds[SLSBridgedSetSpaceManagementModeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSetSpaceManagementModeOperation creates a new SLSBridgedSetSpaceManagementModeOperation instance.
func NewSLSBridgedSetSpaceManagementModeOperation() SLSBridgedSetSpaceManagementModeOperation {
	class := getSLSBridgedSetSpaceManagementModeOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSetSpaceManagementModeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSetSpaceManagementModeOperationWithCoder(coder objectivec.IObject) SLSBridgedSetSpaceManagementModeOperation {
	instance := getSLSBridgedSetSpaceManagementModeOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSetSpaceManagementModeOperationFromID(rv)
}

func NewSLSBridgedSetSpaceManagementModeOperationWithMode(mode uint64) SLSBridgedSetSpaceManagementModeOperation {
	instance := getSLSBridgedSetSpaceManagementModeOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMode:"), mode)
	return SLSBridgedSetSpaceManagementModeOperationFromID(rv)
}

func (s SLSBridgedSetSpaceManagementModeOperation) InitWithMode(mode uint64) SLSBridgedSetSpaceManagementModeOperation {
	rv := objc.SendIfResponds[SLSBridgedSetSpaceManagementModeOperation](s.ID, objc.Sel("initWithMode:"), mode)
	return rv
}

func (s SLSBridgedSetSpaceManagementModeOperation) Mode() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("mode"))
	return rv
}
