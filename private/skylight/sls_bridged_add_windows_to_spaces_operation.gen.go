// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedAddWindowsToSpacesOperation] class.
var (
	_SLSBridgedAddWindowsToSpacesOperationClass     SLSBridgedAddWindowsToSpacesOperationClass
	_SLSBridgedAddWindowsToSpacesOperationClassOnce sync.Once
)

func getSLSBridgedAddWindowsToSpacesOperationClass() SLSBridgedAddWindowsToSpacesOperationClass {
	_SLSBridgedAddWindowsToSpacesOperationClassOnce.Do(func() {
		_SLSBridgedAddWindowsToSpacesOperationClass = SLSBridgedAddWindowsToSpacesOperationClass{class: objc.GetClass("SLSBridgedAddWindowsToSpacesOperation")}
	})
	return _SLSBridgedAddWindowsToSpacesOperationClass
}

// GetSLSBridgedAddWindowsToSpacesOperationClass returns the class object for SLSBridgedAddWindowsToSpacesOperation.
func GetSLSBridgedAddWindowsToSpacesOperationClass() SLSBridgedAddWindowsToSpacesOperationClass {
	return getSLSBridgedAddWindowsToSpacesOperationClass()
}

type SLSBridgedAddWindowsToSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedAddWindowsToSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedAddWindowsToSpacesOperationClass) Alloc() SLSBridgedAddWindowsToSpacesOperation {
	rv := objc.Send[SLSBridgedAddWindowsToSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedAddWindowsToSpacesOperation.Spaces]
//   - [SLSBridgedAddWindowsToSpacesOperation.Windows]
//   - [SLSBridgedAddWindowsToSpacesOperation.InitWithWindowsSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation
type SLSBridgedAddWindowsToSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedAddWindowsToSpacesOperationFromID constructs a [SLSBridgedAddWindowsToSpacesOperation] from an objc.ID.
func SLSBridgedAddWindowsToSpacesOperationFromID(id objc.ID) SLSBridgedAddWindowsToSpacesOperation {
	return SLSBridgedAddWindowsToSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedAddWindowsToSpacesOperation implements ISLSBridgedAddWindowsToSpacesOperation.
var _ ISLSBridgedAddWindowsToSpacesOperation = SLSBridgedAddWindowsToSpacesOperation{}

// An interface definition for the [SLSBridgedAddWindowsToSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedAddWindowsToSpacesOperation.Spaces]
//   - [ISLSBridgedAddWindowsToSpacesOperation.Windows]
//   - [ISLSBridgedAddWindowsToSpacesOperation.InitWithWindowsSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation
type ISLSBridgedAddWindowsToSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Spaces() foundation.INSArray
	Windows() foundation.INSArray
	InitWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedAddWindowsToSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedAddWindowsToSpacesOperation) Init() SLSBridgedAddWindowsToSpacesOperation {
	rv := objc.Send[SLSBridgedAddWindowsToSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedAddWindowsToSpacesOperation) Autorelease() SLSBridgedAddWindowsToSpacesOperation {
	rv := objc.Send[SLSBridgedAddWindowsToSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedAddWindowsToSpacesOperation creates a new SLSBridgedAddWindowsToSpacesOperation instance.
func NewSLSBridgedAddWindowsToSpacesOperation() SLSBridgedAddWindowsToSpacesOperation {
	class := getSLSBridgedAddWindowsToSpacesOperationClass()
	rv := objc.Send[SLSBridgedAddWindowsToSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation/initWithCoder:
func NewSLSBridgedAddWindowsToSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedAddWindowsToSpacesOperation {
	instance := getSLSBridgedAddWindowsToSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedAddWindowsToSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation/initWithWindows:spaces:
func NewSLSBridgedAddWindowsToSpacesOperationWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedAddWindowsToSpacesOperation {
	instance := getSLSBridgedAddWindowsToSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindows:spaces:"), windows, spaces)
	return SLSBridgedAddWindowsToSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation/initWithWindows:spaces:
func (s SLSBridgedAddWindowsToSpacesOperation) InitWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedAddWindowsToSpacesOperation {
	rv := objc.Send[SLSBridgedAddWindowsToSpacesOperation](s.ID, objc.Sel("initWithWindows:spaces:"), windows, spaces)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation/spaces
func (s SLSBridgedAddWindowsToSpacesOperation) Spaces() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedAddWindowsToSpacesOperation/windows
func (s SLSBridgedAddWindowsToSpacesOperation) Windows() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
