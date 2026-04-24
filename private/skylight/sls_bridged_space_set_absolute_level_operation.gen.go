// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetAbsoluteLevelOperation] class.
var (
	_SLSBridgedSpaceSetAbsoluteLevelOperationClass     SLSBridgedSpaceSetAbsoluteLevelOperationClass
	_SLSBridgedSpaceSetAbsoluteLevelOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetAbsoluteLevelOperationClass() SLSBridgedSpaceSetAbsoluteLevelOperationClass {
	_SLSBridgedSpaceSetAbsoluteLevelOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetAbsoluteLevelOperationClass = SLSBridgedSpaceSetAbsoluteLevelOperationClass{class: objc.GetClass("SLSBridgedSpaceSetAbsoluteLevelOperation")}
	})
	return _SLSBridgedSpaceSetAbsoluteLevelOperationClass
}

// GetSLSBridgedSpaceSetAbsoluteLevelOperationClass returns the class object for SLSBridgedSpaceSetAbsoluteLevelOperation.
func GetSLSBridgedSpaceSetAbsoluteLevelOperationClass() SLSBridgedSpaceSetAbsoluteLevelOperationClass {
	return getSLSBridgedSpaceSetAbsoluteLevelOperationClass()
}

type SLSBridgedSpaceSetAbsoluteLevelOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetAbsoluteLevelOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetAbsoluteLevelOperationClass) Alloc() SLSBridgedSpaceSetAbsoluteLevelOperation {
	rv := objc.Send[SLSBridgedSpaceSetAbsoluteLevelOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetAbsoluteLevelOperation.Level]
//   - [SLSBridgedSpaceSetAbsoluteLevelOperation.SpaceID]
//   - [SLSBridgedSpaceSetAbsoluteLevelOperation.InitWithSpaceIDLevel]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation
type SLSBridgedSpaceSetAbsoluteLevelOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetAbsoluteLevelOperationFromID constructs a [SLSBridgedSpaceSetAbsoluteLevelOperation] from an objc.ID.
func SLSBridgedSpaceSetAbsoluteLevelOperationFromID(id objc.ID) SLSBridgedSpaceSetAbsoluteLevelOperation {
	return SLSBridgedSpaceSetAbsoluteLevelOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetAbsoluteLevelOperation implements ISLSBridgedSpaceSetAbsoluteLevelOperation.
var _ ISLSBridgedSpaceSetAbsoluteLevelOperation = SLSBridgedSpaceSetAbsoluteLevelOperation{}

// An interface definition for the [SLSBridgedSpaceSetAbsoluteLevelOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetAbsoluteLevelOperation.Level]
//   - [ISLSBridgedSpaceSetAbsoluteLevelOperation.SpaceID]
//   - [ISLSBridgedSpaceSetAbsoluteLevelOperation.InitWithSpaceIDLevel]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation
type ISLSBridgedSpaceSetAbsoluteLevelOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Level() int
	SpaceID() uint64
	InitWithSpaceIDLevel(id uint64, level int) SLSBridgedSpaceSetAbsoluteLevelOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetAbsoluteLevelOperation) Init() SLSBridgedSpaceSetAbsoluteLevelOperation {
	rv := objc.Send[SLSBridgedSpaceSetAbsoluteLevelOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetAbsoluteLevelOperation) Autorelease() SLSBridgedSpaceSetAbsoluteLevelOperation {
	rv := objc.Send[SLSBridgedSpaceSetAbsoluteLevelOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetAbsoluteLevelOperation creates a new SLSBridgedSpaceSetAbsoluteLevelOperation instance.
func NewSLSBridgedSpaceSetAbsoluteLevelOperation() SLSBridgedSpaceSetAbsoluteLevelOperation {
	class := getSLSBridgedSpaceSetAbsoluteLevelOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetAbsoluteLevelOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation/initWithCoder:
func NewSLSBridgedSpaceSetAbsoluteLevelOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetAbsoluteLevelOperation {
	instance := getSLSBridgedSpaceSetAbsoluteLevelOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetAbsoluteLevelOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation/initWithSpaceID:level:
func NewSLSBridgedSpaceSetAbsoluteLevelOperationWithSpaceIDLevel(id uint64, level int) SLSBridgedSpaceSetAbsoluteLevelOperation {
	instance := getSLSBridgedSpaceSetAbsoluteLevelOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:level:"), id, level)
	return SLSBridgedSpaceSetAbsoluteLevelOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation/initWithSpaceID:level:
func (s SLSBridgedSpaceSetAbsoluteLevelOperation) InitWithSpaceIDLevel(id uint64, level int) SLSBridgedSpaceSetAbsoluteLevelOperation {
	rv := objc.Send[SLSBridgedSpaceSetAbsoluteLevelOperation](s.ID, objc.Sel("initWithSpaceID:level:"), id, level)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation/level
func (s SLSBridgedSpaceSetAbsoluteLevelOperation) Level() int {
	rv := objc.Send[int](s.ID, objc.Sel("level"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAbsoluteLevelOperation/spaceID
func (s SLSBridgedSpaceSetAbsoluteLevelOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
