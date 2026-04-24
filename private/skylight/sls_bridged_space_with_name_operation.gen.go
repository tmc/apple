// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceWithNameOperation] class.
var (
	_SLSBridgedSpaceWithNameOperationClass     SLSBridgedSpaceWithNameOperationClass
	_SLSBridgedSpaceWithNameOperationClassOnce sync.Once
)

func getSLSBridgedSpaceWithNameOperationClass() SLSBridgedSpaceWithNameOperationClass {
	_SLSBridgedSpaceWithNameOperationClassOnce.Do(func() {
		_SLSBridgedSpaceWithNameOperationClass = SLSBridgedSpaceWithNameOperationClass{class: objc.GetClass("SLSBridgedSpaceWithNameOperation")}
	})
	return _SLSBridgedSpaceWithNameOperationClass
}

// GetSLSBridgedSpaceWithNameOperationClass returns the class object for SLSBridgedSpaceWithNameOperation.
func GetSLSBridgedSpaceWithNameOperationClass() SLSBridgedSpaceWithNameOperationClass {
	return getSLSBridgedSpaceWithNameOperationClass()
}

type SLSBridgedSpaceWithNameOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceWithNameOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceWithNameOperationClass) Alloc() SLSBridgedSpaceWithNameOperation {
	rv := objc.Send[SLSBridgedSpaceWithNameOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceWithNameOperation.MakeResultWithSpaceID]
//   - [SLSBridgedSpaceWithNameOperation.Name]
//   - [SLSBridgedSpaceWithNameOperation.InitWithName]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation
type SLSBridgedSpaceWithNameOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceWithNameOperationFromID constructs a [SLSBridgedSpaceWithNameOperation] from an objc.ID.
func SLSBridgedSpaceWithNameOperationFromID(id objc.ID) SLSBridgedSpaceWithNameOperation {
	return SLSBridgedSpaceWithNameOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceWithNameOperation implements ISLSBridgedSpaceWithNameOperation.
var _ ISLSBridgedSpaceWithNameOperation = SLSBridgedSpaceWithNameOperation{}

// An interface definition for the [SLSBridgedSpaceWithNameOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceWithNameOperation.MakeResultWithSpaceID]
//   - [ISLSBridgedSpaceWithNameOperation.Name]
//   - [ISLSBridgedSpaceWithNameOperation.InitWithName]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation
type ISLSBridgedSpaceWithNameOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSpaceID(id uint64) objectivec.IObject
	Name() string
	InitWithName(name objectivec.IObject) SLSBridgedSpaceWithNameOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceWithNameOperation) Init() SLSBridgedSpaceWithNameOperation {
	rv := objc.Send[SLSBridgedSpaceWithNameOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceWithNameOperation) Autorelease() SLSBridgedSpaceWithNameOperation {
	rv := objc.Send[SLSBridgedSpaceWithNameOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceWithNameOperation creates a new SLSBridgedSpaceWithNameOperation instance.
func NewSLSBridgedSpaceWithNameOperation() SLSBridgedSpaceWithNameOperation {
	class := getSLSBridgedSpaceWithNameOperationClass()
	rv := objc.Send[SLSBridgedSpaceWithNameOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation/initWithCoder:
func NewSLSBridgedSpaceWithNameOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceWithNameOperation {
	instance := getSLSBridgedSpaceWithNameOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceWithNameOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation/initWithName:
func NewSLSBridgedSpaceWithNameOperationWithName(name objectivec.IObject) SLSBridgedSpaceWithNameOperation {
	instance := getSLSBridgedSpaceWithNameOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), name)
	return SLSBridgedSpaceWithNameOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation/makeResultWithSpaceID:
func (s SLSBridgedSpaceWithNameOperation) MakeResultWithSpaceID(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation/initWithName:
func (s SLSBridgedSpaceWithNameOperation) InitWithName(name objectivec.IObject) SLSBridgedSpaceWithNameOperation {
	rv := objc.Send[SLSBridgedSpaceWithNameOperation](s.ID, objc.Sel("initWithName:"), name)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceWithNameOperation/name
func (s SLSBridgedSpaceWithNameOperation) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
