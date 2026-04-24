// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation] class.
var (
	_SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass     SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass
	_SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClassOnce sync.Once
)

func getSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass {
	_SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClassOnce.Do(func() {
		_SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass = SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass{class: objc.GetClass("SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation")}
	})
	return _SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass
}

// GetSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass returns the class object for SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.
func GetSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass {
	return getSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass()
}

type SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass) Alloc() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.Options]
//   - [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.SpaceID]
//   - [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.Windows]
//   - [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.InitWithSpaceIDWindowsOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation
type SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationFromID constructs a [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation] from an objc.ID.
func SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationFromID(id objc.ID) SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	return SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation implements ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.
var _ ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation = SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation{}

// An interface definition for the [SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.Options]
//   - [ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.SpaceID]
//   - [ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.Windows]
//   - [ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation.InitWithSpaceIDWindowsOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation
type ISLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Options() uint32
	SpaceID() uint64
	Windows() foundation.INSArray
	InitWithSpaceIDWindowsOptions(id uint64, windows objectivec.IObject, options uint32) SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) Init() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) Autorelease() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation creates a new SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation instance.
func NewSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation() SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	class := getSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass()
	rv := objc.Send[SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/initWithCoder:
func NewSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	instance := getSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/initWithSpaceID:windows:options:
func NewSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationWithSpaceIDWindowsOptions(id uint64, windows objectivec.IObject, options uint32) SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	instance := getSLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:windows:options:"), id, windows, options)
	return SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/initWithSpaceID:windows:options:
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) InitWithSpaceIDWindowsOptions(id uint64, windows objectivec.IObject, options uint32) SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation](s.ID, objc.Sel("initWithSpaceID:windows:options:"), id, windows, options)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/options
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/spaceID
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation/windows
func (s SLSBridgedSpaceAddWindowsAndRemoveFromSpacesOperation) Windows() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
