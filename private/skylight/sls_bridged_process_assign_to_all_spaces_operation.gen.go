// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedProcessAssignToAllSpacesOperation] class.
var (
	_SLSBridgedProcessAssignToAllSpacesOperationClass     SLSBridgedProcessAssignToAllSpacesOperationClass
	_SLSBridgedProcessAssignToAllSpacesOperationClassOnce sync.Once
)

func getSLSBridgedProcessAssignToAllSpacesOperationClass() SLSBridgedProcessAssignToAllSpacesOperationClass {
	_SLSBridgedProcessAssignToAllSpacesOperationClassOnce.Do(func() {
		_SLSBridgedProcessAssignToAllSpacesOperationClass = SLSBridgedProcessAssignToAllSpacesOperationClass{class: objc.GetClass("SLSBridgedProcessAssignToAllSpacesOperation")}
	})
	return _SLSBridgedProcessAssignToAllSpacesOperationClass
}

// GetSLSBridgedProcessAssignToAllSpacesOperationClass returns the class object for SLSBridgedProcessAssignToAllSpacesOperation.
func GetSLSBridgedProcessAssignToAllSpacesOperationClass() SLSBridgedProcessAssignToAllSpacesOperationClass {
	return getSLSBridgedProcessAssignToAllSpacesOperationClass()
}

type SLSBridgedProcessAssignToAllSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedProcessAssignToAllSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedProcessAssignToAllSpacesOperationClass) Alloc() SLSBridgedProcessAssignToAllSpacesOperation {
	rv := objc.Send[SLSBridgedProcessAssignToAllSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedProcessAssignToAllSpacesOperation.Process]
//   - [SLSBridgedProcessAssignToAllSpacesOperation.InitWithProcess]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation
type SLSBridgedProcessAssignToAllSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedProcessAssignToAllSpacesOperationFromID constructs a [SLSBridgedProcessAssignToAllSpacesOperation] from an objc.ID.
func SLSBridgedProcessAssignToAllSpacesOperationFromID(id objc.ID) SLSBridgedProcessAssignToAllSpacesOperation {
	return SLSBridgedProcessAssignToAllSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedProcessAssignToAllSpacesOperation implements ISLSBridgedProcessAssignToAllSpacesOperation.
var _ ISLSBridgedProcessAssignToAllSpacesOperation = SLSBridgedProcessAssignToAllSpacesOperation{}

// An interface definition for the [SLSBridgedProcessAssignToAllSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedProcessAssignToAllSpacesOperation.Process]
//   - [ISLSBridgedProcessAssignToAllSpacesOperation.InitWithProcess]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation
type ISLSBridgedProcessAssignToAllSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Process() int
	InitWithProcess(process int) SLSBridgedProcessAssignToAllSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedProcessAssignToAllSpacesOperation) Init() SLSBridgedProcessAssignToAllSpacesOperation {
	rv := objc.Send[SLSBridgedProcessAssignToAllSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedProcessAssignToAllSpacesOperation) Autorelease() SLSBridgedProcessAssignToAllSpacesOperation {
	rv := objc.Send[SLSBridgedProcessAssignToAllSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedProcessAssignToAllSpacesOperation creates a new SLSBridgedProcessAssignToAllSpacesOperation instance.
func NewSLSBridgedProcessAssignToAllSpacesOperation() SLSBridgedProcessAssignToAllSpacesOperation {
	class := getSLSBridgedProcessAssignToAllSpacesOperationClass()
	rv := objc.Send[SLSBridgedProcessAssignToAllSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation/initWithCoder:
func NewSLSBridgedProcessAssignToAllSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedProcessAssignToAllSpacesOperation {
	instance := getSLSBridgedProcessAssignToAllSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedProcessAssignToAllSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation/initWithProcess:
func NewSLSBridgedProcessAssignToAllSpacesOperationWithProcess(process int) SLSBridgedProcessAssignToAllSpacesOperation {
	instance := getSLSBridgedProcessAssignToAllSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcess:"), process)
	return SLSBridgedProcessAssignToAllSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation/initWithProcess:
func (s SLSBridgedProcessAssignToAllSpacesOperation) InitWithProcess(process int) SLSBridgedProcessAssignToAllSpacesOperation {
	rv := objc.Send[SLSBridgedProcessAssignToAllSpacesOperation](s.ID, objc.Sel("initWithProcess:"), process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToAllSpacesOperation/process
func (s SLSBridgedProcessAssignToAllSpacesOperation) Process() int {
	rv := objc.Send[int](s.ID, objc.Sel("process"))
	return rv
}
