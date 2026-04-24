// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedRemoveWindowsFromSpacesOperation] class.
var (
	_SLSBridgedRemoveWindowsFromSpacesOperationClass     SLSBridgedRemoveWindowsFromSpacesOperationClass
	_SLSBridgedRemoveWindowsFromSpacesOperationClassOnce sync.Once
)

func getSLSBridgedRemoveWindowsFromSpacesOperationClass() SLSBridgedRemoveWindowsFromSpacesOperationClass {
	_SLSBridgedRemoveWindowsFromSpacesOperationClassOnce.Do(func() {
		_SLSBridgedRemoveWindowsFromSpacesOperationClass = SLSBridgedRemoveWindowsFromSpacesOperationClass{class: objc.GetClass("SLSBridgedRemoveWindowsFromSpacesOperation")}
	})
	return _SLSBridgedRemoveWindowsFromSpacesOperationClass
}

// GetSLSBridgedRemoveWindowsFromSpacesOperationClass returns the class object for SLSBridgedRemoveWindowsFromSpacesOperation.
func GetSLSBridgedRemoveWindowsFromSpacesOperationClass() SLSBridgedRemoveWindowsFromSpacesOperationClass {
	return getSLSBridgedRemoveWindowsFromSpacesOperationClass()
}

type SLSBridgedRemoveWindowsFromSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedRemoveWindowsFromSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedRemoveWindowsFromSpacesOperationClass) Alloc() SLSBridgedRemoveWindowsFromSpacesOperation {
	rv := objc.Send[SLSBridgedRemoveWindowsFromSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedRemoveWindowsFromSpacesOperation.Spaces]
//   - [SLSBridgedRemoveWindowsFromSpacesOperation.Windows]
//   - [SLSBridgedRemoveWindowsFromSpacesOperation.InitWithWindowsSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation
type SLSBridgedRemoveWindowsFromSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedRemoveWindowsFromSpacesOperationFromID constructs a [SLSBridgedRemoveWindowsFromSpacesOperation] from an objc.ID.
func SLSBridgedRemoveWindowsFromSpacesOperationFromID(id objc.ID) SLSBridgedRemoveWindowsFromSpacesOperation {
	return SLSBridgedRemoveWindowsFromSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedRemoveWindowsFromSpacesOperation implements ISLSBridgedRemoveWindowsFromSpacesOperation.
var _ ISLSBridgedRemoveWindowsFromSpacesOperation = SLSBridgedRemoveWindowsFromSpacesOperation{}

// An interface definition for the [SLSBridgedRemoveWindowsFromSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedRemoveWindowsFromSpacesOperation.Spaces]
//   - [ISLSBridgedRemoveWindowsFromSpacesOperation.Windows]
//   - [ISLSBridgedRemoveWindowsFromSpacesOperation.InitWithWindowsSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation
type ISLSBridgedRemoveWindowsFromSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Spaces() foundation.INSArray
	Windows() foundation.INSArray
	InitWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedRemoveWindowsFromSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedRemoveWindowsFromSpacesOperation) Init() SLSBridgedRemoveWindowsFromSpacesOperation {
	rv := objc.Send[SLSBridgedRemoveWindowsFromSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedRemoveWindowsFromSpacesOperation) Autorelease() SLSBridgedRemoveWindowsFromSpacesOperation {
	rv := objc.Send[SLSBridgedRemoveWindowsFromSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedRemoveWindowsFromSpacesOperation creates a new SLSBridgedRemoveWindowsFromSpacesOperation instance.
func NewSLSBridgedRemoveWindowsFromSpacesOperation() SLSBridgedRemoveWindowsFromSpacesOperation {
	class := getSLSBridgedRemoveWindowsFromSpacesOperationClass()
	rv := objc.Send[SLSBridgedRemoveWindowsFromSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation/initWithCoder:
func NewSLSBridgedRemoveWindowsFromSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedRemoveWindowsFromSpacesOperation {
	instance := getSLSBridgedRemoveWindowsFromSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedRemoveWindowsFromSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation/initWithWindows:spaces:
func NewSLSBridgedRemoveWindowsFromSpacesOperationWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedRemoveWindowsFromSpacesOperation {
	instance := getSLSBridgedRemoveWindowsFromSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindows:spaces:"), windows, spaces)
	return SLSBridgedRemoveWindowsFromSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation/initWithWindows:spaces:
func (s SLSBridgedRemoveWindowsFromSpacesOperation) InitWithWindowsSpaces(windows objectivec.IObject, spaces objectivec.IObject) SLSBridgedRemoveWindowsFromSpacesOperation {
	rv := objc.Send[SLSBridgedRemoveWindowsFromSpacesOperation](s.ID, objc.Sel("initWithWindows:spaces:"), windows, spaces)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation/spaces
func (s SLSBridgedRemoveWindowsFromSpacesOperation) Spaces() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedRemoveWindowsFromSpacesOperation/windows
func (s SLSBridgedRemoveWindowsFromSpacesOperation) Windows() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
