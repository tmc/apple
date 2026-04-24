// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyBestManagedDisplayForRectOperation] class.
var (
	_SLSBridgedCopyBestManagedDisplayForRectOperationClass     SLSBridgedCopyBestManagedDisplayForRectOperationClass
	_SLSBridgedCopyBestManagedDisplayForRectOperationClassOnce sync.Once
)

func getSLSBridgedCopyBestManagedDisplayForRectOperationClass() SLSBridgedCopyBestManagedDisplayForRectOperationClass {
	_SLSBridgedCopyBestManagedDisplayForRectOperationClassOnce.Do(func() {
		_SLSBridgedCopyBestManagedDisplayForRectOperationClass = SLSBridgedCopyBestManagedDisplayForRectOperationClass{class: objc.GetClass("SLSBridgedCopyBestManagedDisplayForRectOperation")}
	})
	return _SLSBridgedCopyBestManagedDisplayForRectOperationClass
}

// GetSLSBridgedCopyBestManagedDisplayForRectOperationClass returns the class object for SLSBridgedCopyBestManagedDisplayForRectOperation.
func GetSLSBridgedCopyBestManagedDisplayForRectOperationClass() SLSBridgedCopyBestManagedDisplayForRectOperationClass {
	return getSLSBridgedCopyBestManagedDisplayForRectOperationClass()
}

type SLSBridgedCopyBestManagedDisplayForRectOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyBestManagedDisplayForRectOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyBestManagedDisplayForRectOperationClass) Alloc() SLSBridgedCopyBestManagedDisplayForRectOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForRectOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyBestManagedDisplayForRectOperation.MakeResultWithString]
//   - [SLSBridgedCopyBestManagedDisplayForRectOperation.Rect]
//   - [SLSBridgedCopyBestManagedDisplayForRectOperation.InitWithRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation
type SLSBridgedCopyBestManagedDisplayForRectOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyBestManagedDisplayForRectOperationFromID constructs a [SLSBridgedCopyBestManagedDisplayForRectOperation] from an objc.ID.
func SLSBridgedCopyBestManagedDisplayForRectOperationFromID(id objc.ID) SLSBridgedCopyBestManagedDisplayForRectOperation {
	return SLSBridgedCopyBestManagedDisplayForRectOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyBestManagedDisplayForRectOperation implements ISLSBridgedCopyBestManagedDisplayForRectOperation.
var _ ISLSBridgedCopyBestManagedDisplayForRectOperation = SLSBridgedCopyBestManagedDisplayForRectOperation{}

// An interface definition for the [SLSBridgedCopyBestManagedDisplayForRectOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyBestManagedDisplayForRectOperation.MakeResultWithString]
//   - [ISLSBridgedCopyBestManagedDisplayForRectOperation.Rect]
//   - [ISLSBridgedCopyBestManagedDisplayForRectOperation.InitWithRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation
type ISLSBridgedCopyBestManagedDisplayForRectOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithString(string_ objectivec.IObject) objectivec.IObject
	Rect() corefoundation.CGRect
	InitWithRect(rect corefoundation.CGRect) SLSBridgedCopyBestManagedDisplayForRectOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyBestManagedDisplayForRectOperation) Init() SLSBridgedCopyBestManagedDisplayForRectOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForRectOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyBestManagedDisplayForRectOperation) Autorelease() SLSBridgedCopyBestManagedDisplayForRectOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForRectOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyBestManagedDisplayForRectOperation creates a new SLSBridgedCopyBestManagedDisplayForRectOperation instance.
func NewSLSBridgedCopyBestManagedDisplayForRectOperation() SLSBridgedCopyBestManagedDisplayForRectOperation {
	class := getSLSBridgedCopyBestManagedDisplayForRectOperationClass()
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForRectOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation/initWithCoder:
func NewSLSBridgedCopyBestManagedDisplayForRectOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyBestManagedDisplayForRectOperation {
	instance := getSLSBridgedCopyBestManagedDisplayForRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyBestManagedDisplayForRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation/initWithRect:
func NewSLSBridgedCopyBestManagedDisplayForRectOperationWithRect(rect corefoundation.CGRect) SLSBridgedCopyBestManagedDisplayForRectOperation {
	instance := getSLSBridgedCopyBestManagedDisplayForRectOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRect:"), rect)
	return SLSBridgedCopyBestManagedDisplayForRectOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation/makeResultWithString:
func (s SLSBridgedCopyBestManagedDisplayForRectOperation) MakeResultWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation/initWithRect:
func (s SLSBridgedCopyBestManagedDisplayForRectOperation) InitWithRect(rect corefoundation.CGRect) SLSBridgedCopyBestManagedDisplayForRectOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForRectOperation](s.ID, objc.Sel("initWithRect:"), rect)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForRectOperation/rect
func (s SLSBridgedCopyBestManagedDisplayForRectOperation) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}
