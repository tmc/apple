// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedResetWindowsOperation] class.
var (
	_SLSBridgedResetWindowsOperationClass     SLSBridgedResetWindowsOperationClass
	_SLSBridgedResetWindowsOperationClassOnce sync.Once
)

func getSLSBridgedResetWindowsOperationClass() SLSBridgedResetWindowsOperationClass {
	_SLSBridgedResetWindowsOperationClassOnce.Do(func() {
		_SLSBridgedResetWindowsOperationClass = SLSBridgedResetWindowsOperationClass{class: objc.GetClass("SLSBridgedResetWindowsOperation")}
	})
	return _SLSBridgedResetWindowsOperationClass
}

// GetSLSBridgedResetWindowsOperationClass returns the class object for SLSBridgedResetWindowsOperation.
func GetSLSBridgedResetWindowsOperationClass() SLSBridgedResetWindowsOperationClass {
	return getSLSBridgedResetWindowsOperationClass()
}

type SLSBridgedResetWindowsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedResetWindowsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedResetWindowsOperationClass) Alloc() SLSBridgedResetWindowsOperation {
	rv := objc.Send[SLSBridgedResetWindowsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedResetWindowsOperation.Windows]
//   - [SLSBridgedResetWindowsOperation.InitWithWindows]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation
type SLSBridgedResetWindowsOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedResetWindowsOperationFromID constructs a [SLSBridgedResetWindowsOperation] from an objc.ID.
func SLSBridgedResetWindowsOperationFromID(id objc.ID) SLSBridgedResetWindowsOperation {
	return SLSBridgedResetWindowsOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedResetWindowsOperation implements ISLSBridgedResetWindowsOperation.
var _ ISLSBridgedResetWindowsOperation = SLSBridgedResetWindowsOperation{}

// An interface definition for the [SLSBridgedResetWindowsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedResetWindowsOperation.Windows]
//   - [ISLSBridgedResetWindowsOperation.InitWithWindows]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation
type ISLSBridgedResetWindowsOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Windows() foundation.INSArray
	InitWithWindows(windows objectivec.IObject) SLSBridgedResetWindowsOperation
}

// Init initializes the instance.
func (s SLSBridgedResetWindowsOperation) Init() SLSBridgedResetWindowsOperation {
	rv := objc.Send[SLSBridgedResetWindowsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedResetWindowsOperation) Autorelease() SLSBridgedResetWindowsOperation {
	rv := objc.Send[SLSBridgedResetWindowsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedResetWindowsOperation creates a new SLSBridgedResetWindowsOperation instance.
func NewSLSBridgedResetWindowsOperation() SLSBridgedResetWindowsOperation {
	class := getSLSBridgedResetWindowsOperationClass()
	rv := objc.Send[SLSBridgedResetWindowsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation/initWithCoder:
func NewSLSBridgedResetWindowsOperationWithCoder(coder objectivec.IObject) SLSBridgedResetWindowsOperation {
	instance := getSLSBridgedResetWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedResetWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation/initWithWindows:
func NewSLSBridgedResetWindowsOperationWithWindows(windows objectivec.IObject) SLSBridgedResetWindowsOperation {
	instance := getSLSBridgedResetWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindows:"), windows)
	return SLSBridgedResetWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation/initWithWindows:
func (s SLSBridgedResetWindowsOperation) InitWithWindows(windows objectivec.IObject) SLSBridgedResetWindowsOperation {
	rv := objc.Send[SLSBridgedResetWindowsOperation](s.ID, objc.Sel("initWithWindows:"), windows)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedResetWindowsOperation/windows
func (s SLSBridgedResetWindowsOperation) Windows() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
