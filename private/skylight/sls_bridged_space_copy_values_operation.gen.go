// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyValuesOperation] class.
var (
	_SLSBridgedSpaceCopyValuesOperationClass     SLSBridgedSpaceCopyValuesOperationClass
	_SLSBridgedSpaceCopyValuesOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyValuesOperationClass() SLSBridgedSpaceCopyValuesOperationClass {
	_SLSBridgedSpaceCopyValuesOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyValuesOperationClass = SLSBridgedSpaceCopyValuesOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyValuesOperation")}
	})
	return _SLSBridgedSpaceCopyValuesOperationClass
}

// GetSLSBridgedSpaceCopyValuesOperationClass returns the class object for SLSBridgedSpaceCopyValuesOperation.
func GetSLSBridgedSpaceCopyValuesOperationClass() SLSBridgedSpaceCopyValuesOperationClass {
	return getSLSBridgedSpaceCopyValuesOperationClass()
}

type SLSBridgedSpaceCopyValuesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyValuesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyValuesOperationClass) Alloc() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyValuesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyValuesOperation.MakeResultWithPropertyListDictionary]
//   - [SLSBridgedSpaceCopyValuesOperation.SpaceID]
//   - [SLSBridgedSpaceCopyValuesOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation
type SLSBridgedSpaceCopyValuesOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyValuesOperationFromID constructs a [SLSBridgedSpaceCopyValuesOperation] from an objc.ID.
func SLSBridgedSpaceCopyValuesOperationFromID(id objc.ID) SLSBridgedSpaceCopyValuesOperation {
	return SLSBridgedSpaceCopyValuesOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyValuesOperation implements ISLSBridgedSpaceCopyValuesOperation.
var _ ISLSBridgedSpaceCopyValuesOperation = SLSBridgedSpaceCopyValuesOperation{}

// An interface definition for the [SLSBridgedSpaceCopyValuesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyValuesOperation.MakeResultWithPropertyListDictionary]
//   - [ISLSBridgedSpaceCopyValuesOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyValuesOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation
type ISLSBridgedSpaceCopyValuesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithPropertyListDictionary(dictionary objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyValuesOperation) Init() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyValuesOperation) Autorelease() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyValuesOperation creates a new SLSBridgedSpaceCopyValuesOperation instance.
func NewSLSBridgedSpaceCopyValuesOperation() SLSBridgedSpaceCopyValuesOperation {
	class := getSLSBridgedSpaceCopyValuesOperationClass()
	rv := objc.Send[SLSBridgedSpaceCopyValuesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation/initWithCoder:
func NewSLSBridgedSpaceCopyValuesOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyValuesOperation {
	instance := getSLSBridgedSpaceCopyValuesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyValuesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation/initWithSpaceID:
func NewSLSBridgedSpaceCopyValuesOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation {
	instance := getSLSBridgedSpaceCopyValuesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyValuesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation/makeResultWithPropertyListDictionary:
func (s SLSBridgedSpaceCopyValuesOperation) MakeResultWithPropertyListDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithPropertyListDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation/initWithSpaceID:
func (s SLSBridgedSpaceCopyValuesOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyValuesOperation/spaceID
func (s SLSBridgedSpaceCopyValuesOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
