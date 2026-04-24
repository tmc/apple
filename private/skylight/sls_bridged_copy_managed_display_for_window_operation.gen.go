// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyManagedDisplayForWindowOperation] class.
var (
	_SLSBridgedCopyManagedDisplayForWindowOperationClass     SLSBridgedCopyManagedDisplayForWindowOperationClass
	_SLSBridgedCopyManagedDisplayForWindowOperationClassOnce sync.Once
)

func getSLSBridgedCopyManagedDisplayForWindowOperationClass() SLSBridgedCopyManagedDisplayForWindowOperationClass {
	_SLSBridgedCopyManagedDisplayForWindowOperationClassOnce.Do(func() {
		_SLSBridgedCopyManagedDisplayForWindowOperationClass = SLSBridgedCopyManagedDisplayForWindowOperationClass{class: objc.GetClass("SLSBridgedCopyManagedDisplayForWindowOperation")}
	})
	return _SLSBridgedCopyManagedDisplayForWindowOperationClass
}

// GetSLSBridgedCopyManagedDisplayForWindowOperationClass returns the class object for SLSBridgedCopyManagedDisplayForWindowOperation.
func GetSLSBridgedCopyManagedDisplayForWindowOperationClass() SLSBridgedCopyManagedDisplayForWindowOperationClass {
	return getSLSBridgedCopyManagedDisplayForWindowOperationClass()
}

type SLSBridgedCopyManagedDisplayForWindowOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyManagedDisplayForWindowOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyManagedDisplayForWindowOperationClass) Alloc() SLSBridgedCopyManagedDisplayForWindowOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForWindowOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyManagedDisplayForWindowOperation.MakeResultWithString]
//   - [SLSBridgedCopyManagedDisplayForWindowOperation.WindowID]
//   - [SLSBridgedCopyManagedDisplayForWindowOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation
type SLSBridgedCopyManagedDisplayForWindowOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyManagedDisplayForWindowOperationFromID constructs a [SLSBridgedCopyManagedDisplayForWindowOperation] from an objc.ID.
func SLSBridgedCopyManagedDisplayForWindowOperationFromID(id objc.ID) SLSBridgedCopyManagedDisplayForWindowOperation {
	return SLSBridgedCopyManagedDisplayForWindowOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyManagedDisplayForWindowOperation implements ISLSBridgedCopyManagedDisplayForWindowOperation.
var _ ISLSBridgedCopyManagedDisplayForWindowOperation = SLSBridgedCopyManagedDisplayForWindowOperation{}

// An interface definition for the [SLSBridgedCopyManagedDisplayForWindowOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyManagedDisplayForWindowOperation.MakeResultWithString]
//   - [ISLSBridgedCopyManagedDisplayForWindowOperation.WindowID]
//   - [ISLSBridgedCopyManagedDisplayForWindowOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation
type ISLSBridgedCopyManagedDisplayForWindowOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithString(string_ objectivec.IObject) objectivec.IObject
	WindowID() uint32
	InitWithWindowID(id uint32) SLSBridgedCopyManagedDisplayForWindowOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyManagedDisplayForWindowOperation) Init() SLSBridgedCopyManagedDisplayForWindowOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForWindowOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyManagedDisplayForWindowOperation) Autorelease() SLSBridgedCopyManagedDisplayForWindowOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForWindowOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyManagedDisplayForWindowOperation creates a new SLSBridgedCopyManagedDisplayForWindowOperation instance.
func NewSLSBridgedCopyManagedDisplayForWindowOperation() SLSBridgedCopyManagedDisplayForWindowOperation {
	class := getSLSBridgedCopyManagedDisplayForWindowOperationClass()
	rv := objc.Send[SLSBridgedCopyManagedDisplayForWindowOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation/initWithCoder:
func NewSLSBridgedCopyManagedDisplayForWindowOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyManagedDisplayForWindowOperation {
	instance := getSLSBridgedCopyManagedDisplayForWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyManagedDisplayForWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation/initWithWindowID:
func NewSLSBridgedCopyManagedDisplayForWindowOperationWithWindowID(id uint32) SLSBridgedCopyManagedDisplayForWindowOperation {
	instance := getSLSBridgedCopyManagedDisplayForWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSBridgedCopyManagedDisplayForWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation/makeResultWithString:
func (s SLSBridgedCopyManagedDisplayForWindowOperation) MakeResultWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation/initWithWindowID:
func (s SLSBridgedCopyManagedDisplayForWindowOperation) InitWithWindowID(id uint32) SLSBridgedCopyManagedDisplayForWindowOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForWindowOperation](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForWindowOperation/windowID
func (s SLSBridgedCopyManagedDisplayForWindowOperation) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
