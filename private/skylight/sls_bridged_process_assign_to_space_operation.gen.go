// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedProcessAssignToSpaceOperation] class.
var (
	_SLSBridgedProcessAssignToSpaceOperationClass     SLSBridgedProcessAssignToSpaceOperationClass
	_SLSBridgedProcessAssignToSpaceOperationClassOnce sync.Once
)

func getSLSBridgedProcessAssignToSpaceOperationClass() SLSBridgedProcessAssignToSpaceOperationClass {
	_SLSBridgedProcessAssignToSpaceOperationClassOnce.Do(func() {
		_SLSBridgedProcessAssignToSpaceOperationClass = SLSBridgedProcessAssignToSpaceOperationClass{class: objc.GetClass("SLSBridgedProcessAssignToSpaceOperation")}
	})
	return _SLSBridgedProcessAssignToSpaceOperationClass
}

// GetSLSBridgedProcessAssignToSpaceOperationClass returns the class object for SLSBridgedProcessAssignToSpaceOperation.
func GetSLSBridgedProcessAssignToSpaceOperationClass() SLSBridgedProcessAssignToSpaceOperationClass {
	return getSLSBridgedProcessAssignToSpaceOperationClass()
}

type SLSBridgedProcessAssignToSpaceOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedProcessAssignToSpaceOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedProcessAssignToSpaceOperationClass) Alloc() SLSBridgedProcessAssignToSpaceOperation {
	rv := objc.Send[SLSBridgedProcessAssignToSpaceOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedProcessAssignToSpaceOperation.Process]
//   - [SLSBridgedProcessAssignToSpaceOperation.SpaceID]
//   - [SLSBridgedProcessAssignToSpaceOperation.InitWithProcessSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation
type SLSBridgedProcessAssignToSpaceOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedProcessAssignToSpaceOperationFromID constructs a [SLSBridgedProcessAssignToSpaceOperation] from an objc.ID.
func SLSBridgedProcessAssignToSpaceOperationFromID(id objc.ID) SLSBridgedProcessAssignToSpaceOperation {
	return SLSBridgedProcessAssignToSpaceOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedProcessAssignToSpaceOperation implements ISLSBridgedProcessAssignToSpaceOperation.
var _ ISLSBridgedProcessAssignToSpaceOperation = SLSBridgedProcessAssignToSpaceOperation{}

// An interface definition for the [SLSBridgedProcessAssignToSpaceOperation] class.
//
// # Methods
//
//   - [ISLSBridgedProcessAssignToSpaceOperation.Process]
//   - [ISLSBridgedProcessAssignToSpaceOperation.SpaceID]
//   - [ISLSBridgedProcessAssignToSpaceOperation.InitWithProcessSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation
type ISLSBridgedProcessAssignToSpaceOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Process() int
	SpaceID() uint64
	InitWithProcessSpaceID(process int, id uint64) SLSBridgedProcessAssignToSpaceOperation
}

// Init initializes the instance.
func (s SLSBridgedProcessAssignToSpaceOperation) Init() SLSBridgedProcessAssignToSpaceOperation {
	rv := objc.Send[SLSBridgedProcessAssignToSpaceOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedProcessAssignToSpaceOperation) Autorelease() SLSBridgedProcessAssignToSpaceOperation {
	rv := objc.Send[SLSBridgedProcessAssignToSpaceOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedProcessAssignToSpaceOperation creates a new SLSBridgedProcessAssignToSpaceOperation instance.
func NewSLSBridgedProcessAssignToSpaceOperation() SLSBridgedProcessAssignToSpaceOperation {
	class := getSLSBridgedProcessAssignToSpaceOperationClass()
	rv := objc.Send[SLSBridgedProcessAssignToSpaceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation/initWithCoder:
func NewSLSBridgedProcessAssignToSpaceOperationWithCoder(coder objectivec.IObject) SLSBridgedProcessAssignToSpaceOperation {
	instance := getSLSBridgedProcessAssignToSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedProcessAssignToSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation/initWithProcess:spaceID:
func NewSLSBridgedProcessAssignToSpaceOperationWithProcessSpaceID(process int, id uint64) SLSBridgedProcessAssignToSpaceOperation {
	instance := getSLSBridgedProcessAssignToSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcess:spaceID:"), process, id)
	return SLSBridgedProcessAssignToSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation/initWithProcess:spaceID:
func (s SLSBridgedProcessAssignToSpaceOperation) InitWithProcessSpaceID(process int, id uint64) SLSBridgedProcessAssignToSpaceOperation {
	rv := objc.Send[SLSBridgedProcessAssignToSpaceOperation](s.ID, objc.Sel("initWithProcess:spaceID:"), process, id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation/process
func (s SLSBridgedProcessAssignToSpaceOperation) Process() int {
	rv := objc.Send[int](s.ID, objc.Sel("process"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedProcessAssignToSpaceOperation/spaceID
func (s SLSBridgedProcessAssignToSpaceOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
