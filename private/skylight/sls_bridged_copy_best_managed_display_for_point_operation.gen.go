// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyBestManagedDisplayForPointOperation] class.
var (
	_SLSBridgedCopyBestManagedDisplayForPointOperationClass     SLSBridgedCopyBestManagedDisplayForPointOperationClass
	_SLSBridgedCopyBestManagedDisplayForPointOperationClassOnce sync.Once
)

func getSLSBridgedCopyBestManagedDisplayForPointOperationClass() SLSBridgedCopyBestManagedDisplayForPointOperationClass {
	_SLSBridgedCopyBestManagedDisplayForPointOperationClassOnce.Do(func() {
		_SLSBridgedCopyBestManagedDisplayForPointOperationClass = SLSBridgedCopyBestManagedDisplayForPointOperationClass{class: objc.GetClass("SLSBridgedCopyBestManagedDisplayForPointOperation")}
	})
	return _SLSBridgedCopyBestManagedDisplayForPointOperationClass
}

// GetSLSBridgedCopyBestManagedDisplayForPointOperationClass returns the class object for SLSBridgedCopyBestManagedDisplayForPointOperation.
func GetSLSBridgedCopyBestManagedDisplayForPointOperationClass() SLSBridgedCopyBestManagedDisplayForPointOperationClass {
	return getSLSBridgedCopyBestManagedDisplayForPointOperationClass()
}

type SLSBridgedCopyBestManagedDisplayForPointOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyBestManagedDisplayForPointOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyBestManagedDisplayForPointOperationClass) Alloc() SLSBridgedCopyBestManagedDisplayForPointOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForPointOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyBestManagedDisplayForPointOperation.MakeResultWithString]
//   - [SLSBridgedCopyBestManagedDisplayForPointOperation.Point]
//   - [SLSBridgedCopyBestManagedDisplayForPointOperation.InitWithPoint]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation
type SLSBridgedCopyBestManagedDisplayForPointOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyBestManagedDisplayForPointOperationFromID constructs a [SLSBridgedCopyBestManagedDisplayForPointOperation] from an objc.ID.
func SLSBridgedCopyBestManagedDisplayForPointOperationFromID(id objc.ID) SLSBridgedCopyBestManagedDisplayForPointOperation {
	return SLSBridgedCopyBestManagedDisplayForPointOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyBestManagedDisplayForPointOperation implements ISLSBridgedCopyBestManagedDisplayForPointOperation.
var _ ISLSBridgedCopyBestManagedDisplayForPointOperation = SLSBridgedCopyBestManagedDisplayForPointOperation{}

// An interface definition for the [SLSBridgedCopyBestManagedDisplayForPointOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyBestManagedDisplayForPointOperation.MakeResultWithString]
//   - [ISLSBridgedCopyBestManagedDisplayForPointOperation.Point]
//   - [ISLSBridgedCopyBestManagedDisplayForPointOperation.InitWithPoint]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation
type ISLSBridgedCopyBestManagedDisplayForPointOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithString(string_ objectivec.IObject) objectivec.IObject
	Point() corefoundation.CGPoint
	InitWithPoint(point corefoundation.CGPoint) SLSBridgedCopyBestManagedDisplayForPointOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyBestManagedDisplayForPointOperation) Init() SLSBridgedCopyBestManagedDisplayForPointOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForPointOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyBestManagedDisplayForPointOperation) Autorelease() SLSBridgedCopyBestManagedDisplayForPointOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForPointOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyBestManagedDisplayForPointOperation creates a new SLSBridgedCopyBestManagedDisplayForPointOperation instance.
func NewSLSBridgedCopyBestManagedDisplayForPointOperation() SLSBridgedCopyBestManagedDisplayForPointOperation {
	class := getSLSBridgedCopyBestManagedDisplayForPointOperationClass()
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForPointOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation/initWithCoder:
func NewSLSBridgedCopyBestManagedDisplayForPointOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyBestManagedDisplayForPointOperation {
	instance := getSLSBridgedCopyBestManagedDisplayForPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyBestManagedDisplayForPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation/initWithPoint:
func NewSLSBridgedCopyBestManagedDisplayForPointOperationWithPoint(point corefoundation.CGPoint) SLSBridgedCopyBestManagedDisplayForPointOperation {
	instance := getSLSBridgedCopyBestManagedDisplayForPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPoint:"), point)
	return SLSBridgedCopyBestManagedDisplayForPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation/makeResultWithString:
func (s SLSBridgedCopyBestManagedDisplayForPointOperation) MakeResultWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation/initWithPoint:
func (s SLSBridgedCopyBestManagedDisplayForPointOperation) InitWithPoint(point corefoundation.CGPoint) SLSBridgedCopyBestManagedDisplayForPointOperation {
	rv := objc.Send[SLSBridgedCopyBestManagedDisplayForPointOperation](s.ID, objc.Sel("initWithPoint:"), point)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyBestManagedDisplayForPointOperation/point
func (s SLSBridgedCopyBestManagedDisplayForPointOperation) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}
