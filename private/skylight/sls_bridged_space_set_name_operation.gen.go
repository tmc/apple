// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetNameOperation] class.
var (
	_SLSBridgedSpaceSetNameOperationClass     SLSBridgedSpaceSetNameOperationClass
	_SLSBridgedSpaceSetNameOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetNameOperationClass() SLSBridgedSpaceSetNameOperationClass {
	_SLSBridgedSpaceSetNameOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetNameOperationClass = SLSBridgedSpaceSetNameOperationClass{class: objc.GetClass("SLSBridgedSpaceSetNameOperation")}
	})
	return _SLSBridgedSpaceSetNameOperationClass
}

// GetSLSBridgedSpaceSetNameOperationClass returns the class object for SLSBridgedSpaceSetNameOperation.
func GetSLSBridgedSpaceSetNameOperationClass() SLSBridgedSpaceSetNameOperationClass {
	return getSLSBridgedSpaceSetNameOperationClass()
}

type SLSBridgedSpaceSetNameOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetNameOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetNameOperationClass) Alloc() SLSBridgedSpaceSetNameOperation {
	rv := objc.Send[SLSBridgedSpaceSetNameOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetNameOperation.Name]
//   - [SLSBridgedSpaceSetNameOperation.SpaceID]
//   - [SLSBridgedSpaceSetNameOperation.InitWithSpaceIDName]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation
type SLSBridgedSpaceSetNameOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetNameOperationFromID constructs a [SLSBridgedSpaceSetNameOperation] from an objc.ID.
func SLSBridgedSpaceSetNameOperationFromID(id objc.ID) SLSBridgedSpaceSetNameOperation {
	return SLSBridgedSpaceSetNameOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetNameOperation implements ISLSBridgedSpaceSetNameOperation.
var _ ISLSBridgedSpaceSetNameOperation = SLSBridgedSpaceSetNameOperation{}

// An interface definition for the [SLSBridgedSpaceSetNameOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetNameOperation.Name]
//   - [ISLSBridgedSpaceSetNameOperation.SpaceID]
//   - [ISLSBridgedSpaceSetNameOperation.InitWithSpaceIDName]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation
type ISLSBridgedSpaceSetNameOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Name() string
	SpaceID() uint64
	InitWithSpaceIDName(id uint64, name objectivec.IObject) SLSBridgedSpaceSetNameOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetNameOperation) Init() SLSBridgedSpaceSetNameOperation {
	rv := objc.Send[SLSBridgedSpaceSetNameOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetNameOperation) Autorelease() SLSBridgedSpaceSetNameOperation {
	rv := objc.Send[SLSBridgedSpaceSetNameOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetNameOperation creates a new SLSBridgedSpaceSetNameOperation instance.
func NewSLSBridgedSpaceSetNameOperation() SLSBridgedSpaceSetNameOperation {
	class := getSLSBridgedSpaceSetNameOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetNameOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation/initWithCoder:
func NewSLSBridgedSpaceSetNameOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetNameOperation {
	instance := getSLSBridgedSpaceSetNameOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetNameOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation/initWithSpaceID:name:
func NewSLSBridgedSpaceSetNameOperationWithSpaceIDName(id uint64, name objectivec.IObject) SLSBridgedSpaceSetNameOperation {
	instance := getSLSBridgedSpaceSetNameOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:name:"), id, name)
	return SLSBridgedSpaceSetNameOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation/initWithSpaceID:name:
func (s SLSBridgedSpaceSetNameOperation) InitWithSpaceIDName(id uint64, name objectivec.IObject) SLSBridgedSpaceSetNameOperation {
	rv := objc.Send[SLSBridgedSpaceSetNameOperation](s.ID, objc.Sel("initWithSpaceID:name:"), id, name)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation/name
func (s SLSBridgedSpaceSetNameOperation) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetNameOperation/spaceID
func (s SLSBridgedSpaceSetNameOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
