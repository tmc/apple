// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceFinishedResizeForRectOperation] class.
var (
	_SLSBridgedSpaceFinishedResizeForRectOperationClass     SLSBridgedSpaceFinishedResizeForRectOperationClass
	_SLSBridgedSpaceFinishedResizeForRectOperationClassOnce sync.Once
)

func getSLSBridgedSpaceFinishedResizeForRectOperationClass() SLSBridgedSpaceFinishedResizeForRectOperationClass {
	_SLSBridgedSpaceFinishedResizeForRectOperationClassOnce.Do(func() {
		_SLSBridgedSpaceFinishedResizeForRectOperationClass = SLSBridgedSpaceFinishedResizeForRectOperationClass{class: objc.GetClass("SLSBridgedSpaceFinishedResizeForRectOperation")}
	})
	return _SLSBridgedSpaceFinishedResizeForRectOperationClass
}

// GetSLSBridgedSpaceFinishedResizeForRectOperationClass returns the class object for SLSBridgedSpaceFinishedResizeForRectOperation.
func GetSLSBridgedSpaceFinishedResizeForRectOperationClass() SLSBridgedSpaceFinishedResizeForRectOperationClass {
	return getSLSBridgedSpaceFinishedResizeForRectOperationClass()
}

type SLSBridgedSpaceFinishedResizeForRectOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceFinishedResizeForRectOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceFinishedResizeForRectOperationClass) Alloc() SLSBridgedSpaceFinishedResizeForRectOperation {
	rv := objc.Send[SLSBridgedSpaceFinishedResizeForRectOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceFinishedResizeForRectOperation.Rect]
//   - [SLSBridgedSpaceFinishedResizeForRectOperation.SpaceID]
//   - [SLSBridgedSpaceFinishedResizeForRectOperation.InitWithSpaceIDRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation
type SLSBridgedSpaceFinishedResizeForRectOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceFinishedResizeForRectOperationFromID constructs a [SLSBridgedSpaceFinishedResizeForRectOperation] from an objc.ID.
func SLSBridgedSpaceFinishedResizeForRectOperationFromID(id objc.ID) SLSBridgedSpaceFinishedResizeForRectOperation {
	return SLSBridgedSpaceFinishedResizeForRectOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceFinishedResizeForRectOperation implements ISLSBridgedSpaceFinishedResizeForRectOperation.
var _ ISLSBridgedSpaceFinishedResizeForRectOperation = SLSBridgedSpaceFinishedResizeForRectOperation{}

// An interface definition for the [SLSBridgedSpaceFinishedResizeForRectOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceFinishedResizeForRectOperation.Rect]
//   - [ISLSBridgedSpaceFinishedResizeForRectOperation.SpaceID]
//   - [ISLSBridgedSpaceFinishedResizeForRectOperation.InitWithSpaceIDRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation
type ISLSBridgedSpaceFinishedResizeForRectOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Rect() corefoundation.CGRect
	SpaceID() uint64
	InitWithSpaceIDRect(id uint64, rect corefoundation.CGRect) SLSBridgedSpaceFinishedResizeForRectOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceFinishedResizeForRectOperation) Init() SLSBridgedSpaceFinishedResizeForRectOperation {
	rv := objc.Send[SLSBridgedSpaceFinishedResizeForRectOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceFinishedResizeForRectOperation) Autorelease() SLSBridgedSpaceFinishedResizeForRectOperation {
	rv := objc.Send[SLSBridgedSpaceFinishedResizeForRectOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceFinishedResizeForRectOperation creates a new SLSBridgedSpaceFinishedResizeForRectOperation instance.
func NewSLSBridgedSpaceFinishedResizeForRectOperation() SLSBridgedSpaceFinishedResizeForRectOperation {
	class := getSLSBridgedSpaceFinishedResizeForRectOperationClass()
	rv := objc.Send[SLSBridgedSpaceFinishedResizeForRectOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation/initWithCoder:
func NewSLSBridgedSpaceFinishedResizeForRectOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceFinishedResizeForRectOperation {
	instance := getSLSBridgedSpaceFinishedResizeForRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceFinishedResizeForRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation/initWithSpaceID:rect:
func NewSLSBridgedSpaceFinishedResizeForRectOperationWithSpaceIDRect(id uint64, rect corefoundation.CGRect) SLSBridgedSpaceFinishedResizeForRectOperation {
	instance := getSLSBridgedSpaceFinishedResizeForRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:rect:"), id, rect)
	return SLSBridgedSpaceFinishedResizeForRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation/initWithSpaceID:rect:
func (s SLSBridgedSpaceFinishedResizeForRectOperation) InitWithSpaceIDRect(id uint64, rect corefoundation.CGRect) SLSBridgedSpaceFinishedResizeForRectOperation {
	rv := objc.Send[SLSBridgedSpaceFinishedResizeForRectOperation](s.ID, objc.Sel("initWithSpaceID:rect:"), id, rect)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation/rect
func (s SLSBridgedSpaceFinishedResizeForRectOperation) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceFinishedResizeForRectOperation/spaceID
func (s SLSBridgedSpaceFinishedResizeForRectOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
