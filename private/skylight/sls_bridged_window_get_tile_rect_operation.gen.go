// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowGetTileRectOperation] class.
var (
	_SLSBridgedWindowGetTileRectOperationClass     SLSBridgedWindowGetTileRectOperationClass
	_SLSBridgedWindowGetTileRectOperationClassOnce sync.Once
)

func getSLSBridgedWindowGetTileRectOperationClass() SLSBridgedWindowGetTileRectOperationClass {
	_SLSBridgedWindowGetTileRectOperationClassOnce.Do(func() {
		_SLSBridgedWindowGetTileRectOperationClass = SLSBridgedWindowGetTileRectOperationClass{class: objc.GetClass("SLSBridgedWindowGetTileRectOperation")}
	})
	return _SLSBridgedWindowGetTileRectOperationClass
}

// GetSLSBridgedWindowGetTileRectOperationClass returns the class object for SLSBridgedWindowGetTileRectOperation.
func GetSLSBridgedWindowGetTileRectOperationClass() SLSBridgedWindowGetTileRectOperationClass {
	return getSLSBridgedWindowGetTileRectOperationClass()
}

type SLSBridgedWindowGetTileRectOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowGetTileRectOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowGetTileRectOperationClass) Alloc() SLSBridgedWindowGetTileRectOperation {
	rv := objc.Send[SLSBridgedWindowGetTileRectOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowGetTileRectOperation.MakeResultWithRect]
//   - [SLSBridgedWindowGetTileRectOperation.WindowID]
//   - [SLSBridgedWindowGetTileRectOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation
type SLSBridgedWindowGetTileRectOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedWindowGetTileRectOperationFromID constructs a [SLSBridgedWindowGetTileRectOperation] from an objc.ID.
func SLSBridgedWindowGetTileRectOperationFromID(id objc.ID) SLSBridgedWindowGetTileRectOperation {
	return SLSBridgedWindowGetTileRectOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedWindowGetTileRectOperation implements ISLSBridgedWindowGetTileRectOperation.
var _ ISLSBridgedWindowGetTileRectOperation = SLSBridgedWindowGetTileRectOperation{}

// An interface definition for the [SLSBridgedWindowGetTileRectOperation] class.
//
// # Methods
//
//   - [ISLSBridgedWindowGetTileRectOperation.MakeResultWithRect]
//   - [ISLSBridgedWindowGetTileRectOperation.WindowID]
//   - [ISLSBridgedWindowGetTileRectOperation.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation
type ISLSBridgedWindowGetTileRectOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithRect(rect corefoundation.CGRect) objectivec.IObject
	WindowID() uint32
	InitWithWindowID(id uint32) SLSBridgedWindowGetTileRectOperation
}

// Init initializes the instance.
func (s SLSBridgedWindowGetTileRectOperation) Init() SLSBridgedWindowGetTileRectOperation {
	rv := objc.Send[SLSBridgedWindowGetTileRectOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowGetTileRectOperation) Autorelease() SLSBridgedWindowGetTileRectOperation {
	rv := objc.Send[SLSBridgedWindowGetTileRectOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowGetTileRectOperation creates a new SLSBridgedWindowGetTileRectOperation instance.
func NewSLSBridgedWindowGetTileRectOperation() SLSBridgedWindowGetTileRectOperation {
	class := getSLSBridgedWindowGetTileRectOperationClass()
	rv := objc.Send[SLSBridgedWindowGetTileRectOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation/initWithCoder:
func NewSLSBridgedWindowGetTileRectOperationWithCoder(coder objectivec.IObject) SLSBridgedWindowGetTileRectOperation {
	instance := getSLSBridgedWindowGetTileRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowGetTileRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation/initWithWindowID:
func NewSLSBridgedWindowGetTileRectOperationWithWindowID(id uint32) SLSBridgedWindowGetTileRectOperation {
	instance := getSLSBridgedWindowGetTileRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSBridgedWindowGetTileRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation/makeResultWithRect:
func (s SLSBridgedWindowGetTileRectOperation) MakeResultWithRect(rect corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithRect:"), rect)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation/initWithWindowID:
func (s SLSBridgedWindowGetTileRectOperation) InitWithWindowID(id uint32) SLSBridgedWindowGetTileRectOperation {
	rv := objc.Send[SLSBridgedWindowGetTileRectOperation](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowGetTileRectOperation/windowID
func (s SLSBridgedWindowGetTileRectOperation) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
