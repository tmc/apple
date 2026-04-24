// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyAssociatedWindowsOperation] class.
var (
	_SLSBridgedCopyAssociatedWindowsOperationClass     SLSBridgedCopyAssociatedWindowsOperationClass
	_SLSBridgedCopyAssociatedWindowsOperationClassOnce sync.Once
)

func getSLSBridgedCopyAssociatedWindowsOperationClass() SLSBridgedCopyAssociatedWindowsOperationClass {
	_SLSBridgedCopyAssociatedWindowsOperationClassOnce.Do(func() {
		_SLSBridgedCopyAssociatedWindowsOperationClass = SLSBridgedCopyAssociatedWindowsOperationClass{class: objc.GetClass("SLSBridgedCopyAssociatedWindowsOperation")}
	})
	return _SLSBridgedCopyAssociatedWindowsOperationClass
}

// GetSLSBridgedCopyAssociatedWindowsOperationClass returns the class object for SLSBridgedCopyAssociatedWindowsOperation.
func GetSLSBridgedCopyAssociatedWindowsOperationClass() SLSBridgedCopyAssociatedWindowsOperationClass {
	return getSLSBridgedCopyAssociatedWindowsOperationClass()
}

type SLSBridgedCopyAssociatedWindowsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyAssociatedWindowsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyAssociatedWindowsOperationClass) Alloc() SLSBridgedCopyAssociatedWindowsOperation {
	rv := objc.Send[SLSBridgedCopyAssociatedWindowsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyAssociatedWindowsOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopyAssociatedWindowsOperation.WindowID]
//   - [SLSBridgedCopyAssociatedWindowsOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation
type SLSBridgedCopyAssociatedWindowsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyAssociatedWindowsOperationFromID constructs a [SLSBridgedCopyAssociatedWindowsOperation] from an objc.ID.
func SLSBridgedCopyAssociatedWindowsOperationFromID(id objc.ID) SLSBridgedCopyAssociatedWindowsOperation {
	return SLSBridgedCopyAssociatedWindowsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyAssociatedWindowsOperation implements ISLSBridgedCopyAssociatedWindowsOperation.
var _ ISLSBridgedCopyAssociatedWindowsOperation = SLSBridgedCopyAssociatedWindowsOperation{}

// An interface definition for the [SLSBridgedCopyAssociatedWindowsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyAssociatedWindowsOperation.MakeResultWithNumbers]
//   - [ISLSBridgedCopyAssociatedWindowsOperation.WindowID]
//   - [ISLSBridgedCopyAssociatedWindowsOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation
type ISLSBridgedCopyAssociatedWindowsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	WindowID() uint32
	InitWithWindowID(id uint32) SLSBridgedCopyAssociatedWindowsOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyAssociatedWindowsOperation) Init() SLSBridgedCopyAssociatedWindowsOperation {
	rv := objc.Send[SLSBridgedCopyAssociatedWindowsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyAssociatedWindowsOperation) Autorelease() SLSBridgedCopyAssociatedWindowsOperation {
	rv := objc.Send[SLSBridgedCopyAssociatedWindowsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyAssociatedWindowsOperation creates a new SLSBridgedCopyAssociatedWindowsOperation instance.
func NewSLSBridgedCopyAssociatedWindowsOperation() SLSBridgedCopyAssociatedWindowsOperation {
	class := getSLSBridgedCopyAssociatedWindowsOperationClass()
	rv := objc.Send[SLSBridgedCopyAssociatedWindowsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation/initWithCoder:
func NewSLSBridgedCopyAssociatedWindowsOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyAssociatedWindowsOperation {
	instance := getSLSBridgedCopyAssociatedWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyAssociatedWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation/initWithWindowID:
func NewSLSBridgedCopyAssociatedWindowsOperationWithWindowID(id uint32) SLSBridgedCopyAssociatedWindowsOperation {
	instance := getSLSBridgedCopyAssociatedWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSBridgedCopyAssociatedWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation/makeResultWithNumbers:
func (s SLSBridgedCopyAssociatedWindowsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation/initWithWindowID:
func (s SLSBridgedCopyAssociatedWindowsOperation) InitWithWindowID(id uint32) SLSBridgedCopyAssociatedWindowsOperation {
	rv := objc.Send[SLSBridgedCopyAssociatedWindowsOperation](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyAssociatedWindowsOperation/windowID
func (s SLSBridgedCopyAssociatedWindowsOperation) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
