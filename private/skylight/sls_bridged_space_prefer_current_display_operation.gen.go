// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpacePreferCurrentDisplayOperation] class.
var (
	_SLSBridgedSpacePreferCurrentDisplayOperationClass     SLSBridgedSpacePreferCurrentDisplayOperationClass
	_SLSBridgedSpacePreferCurrentDisplayOperationClassOnce sync.Once
)

func getSLSBridgedSpacePreferCurrentDisplayOperationClass() SLSBridgedSpacePreferCurrentDisplayOperationClass {
	_SLSBridgedSpacePreferCurrentDisplayOperationClassOnce.Do(func() {
		_SLSBridgedSpacePreferCurrentDisplayOperationClass = SLSBridgedSpacePreferCurrentDisplayOperationClass{class: objc.GetClass("SLSBridgedSpacePreferCurrentDisplayOperation")}
	})
	return _SLSBridgedSpacePreferCurrentDisplayOperationClass
}

// GetSLSBridgedSpacePreferCurrentDisplayOperationClass returns the class object for SLSBridgedSpacePreferCurrentDisplayOperation.
func GetSLSBridgedSpacePreferCurrentDisplayOperationClass() SLSBridgedSpacePreferCurrentDisplayOperationClass {
	return getSLSBridgedSpacePreferCurrentDisplayOperationClass()
}

type SLSBridgedSpacePreferCurrentDisplayOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpacePreferCurrentDisplayOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpacePreferCurrentDisplayOperationClass) Alloc() SLSBridgedSpacePreferCurrentDisplayOperation {
	rv := objc.Send[SLSBridgedSpacePreferCurrentDisplayOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpacePreferCurrentDisplayOperation.SpaceID]
//   - [SLSBridgedSpacePreferCurrentDisplayOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation
type SLSBridgedSpacePreferCurrentDisplayOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpacePreferCurrentDisplayOperationFromID constructs a [SLSBridgedSpacePreferCurrentDisplayOperation] from an objc.ID.
func SLSBridgedSpacePreferCurrentDisplayOperationFromID(id objc.ID) SLSBridgedSpacePreferCurrentDisplayOperation {
	return SLSBridgedSpacePreferCurrentDisplayOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpacePreferCurrentDisplayOperation implements ISLSBridgedSpacePreferCurrentDisplayOperation.
var _ ISLSBridgedSpacePreferCurrentDisplayOperation = SLSBridgedSpacePreferCurrentDisplayOperation{}

// An interface definition for the [SLSBridgedSpacePreferCurrentDisplayOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpacePreferCurrentDisplayOperation.SpaceID]
//   - [ISLSBridgedSpacePreferCurrentDisplayOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation
type ISLSBridgedSpacePreferCurrentDisplayOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpacePreferCurrentDisplayOperation
}

// Init initializes the instance.
func (s SLSBridgedSpacePreferCurrentDisplayOperation) Init() SLSBridgedSpacePreferCurrentDisplayOperation {
	rv := objc.Send[SLSBridgedSpacePreferCurrentDisplayOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpacePreferCurrentDisplayOperation) Autorelease() SLSBridgedSpacePreferCurrentDisplayOperation {
	rv := objc.Send[SLSBridgedSpacePreferCurrentDisplayOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpacePreferCurrentDisplayOperation creates a new SLSBridgedSpacePreferCurrentDisplayOperation instance.
func NewSLSBridgedSpacePreferCurrentDisplayOperation() SLSBridgedSpacePreferCurrentDisplayOperation {
	class := getSLSBridgedSpacePreferCurrentDisplayOperationClass()
	rv := objc.Send[SLSBridgedSpacePreferCurrentDisplayOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation/initWithCoder:
func NewSLSBridgedSpacePreferCurrentDisplayOperationWithCoder(coder objectivec.IObject) SLSBridgedSpacePreferCurrentDisplayOperation {
	instance := getSLSBridgedSpacePreferCurrentDisplayOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpacePreferCurrentDisplayOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation/initWithSpaceID:
func NewSLSBridgedSpacePreferCurrentDisplayOperationWithSpaceID(id uint64) SLSBridgedSpacePreferCurrentDisplayOperation {
	instance := getSLSBridgedSpacePreferCurrentDisplayOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpacePreferCurrentDisplayOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation/initWithSpaceID:
func (s SLSBridgedSpacePreferCurrentDisplayOperation) InitWithSpaceID(id uint64) SLSBridgedSpacePreferCurrentDisplayOperation {
	rv := objc.Send[SLSBridgedSpacePreferCurrentDisplayOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpacePreferCurrentDisplayOperation/spaceID
func (s SLSBridgedSpacePreferCurrentDisplayOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
