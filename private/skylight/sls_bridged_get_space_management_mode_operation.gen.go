// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedGetSpaceManagementModeOperation] class.
var (
	_SLSBridgedGetSpaceManagementModeOperationClass     SLSBridgedGetSpaceManagementModeOperationClass
	_SLSBridgedGetSpaceManagementModeOperationClassOnce sync.Once
)

func getSLSBridgedGetSpaceManagementModeOperationClass() SLSBridgedGetSpaceManagementModeOperationClass {
	_SLSBridgedGetSpaceManagementModeOperationClassOnce.Do(func() {
		_SLSBridgedGetSpaceManagementModeOperationClass = SLSBridgedGetSpaceManagementModeOperationClass{class: objc.GetClass("SLSBridgedGetSpaceManagementModeOperation")}
	})
	return _SLSBridgedGetSpaceManagementModeOperationClass
}

// GetSLSBridgedGetSpaceManagementModeOperationClass returns the class object for SLSBridgedGetSpaceManagementModeOperation.
func GetSLSBridgedGetSpaceManagementModeOperationClass() SLSBridgedGetSpaceManagementModeOperationClass {
	return getSLSBridgedGetSpaceManagementModeOperationClass()
}

type SLSBridgedGetSpaceManagementModeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedGetSpaceManagementModeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedGetSpaceManagementModeOperationClass) Alloc() SLSBridgedGetSpaceManagementModeOperation {
	rv := objc.Send[SLSBridgedGetSpaceManagementModeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedGetSpaceManagementModeOperation.MakeResultWithSpaceManagementMode]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceManagementModeOperation
type SLSBridgedGetSpaceManagementModeOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedGetSpaceManagementModeOperationFromID constructs a [SLSBridgedGetSpaceManagementModeOperation] from an objc.ID.
func SLSBridgedGetSpaceManagementModeOperationFromID(id objc.ID) SLSBridgedGetSpaceManagementModeOperation {
	return SLSBridgedGetSpaceManagementModeOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedGetSpaceManagementModeOperation implements ISLSBridgedGetSpaceManagementModeOperation.
var _ ISLSBridgedGetSpaceManagementModeOperation = SLSBridgedGetSpaceManagementModeOperation{}

// An interface definition for the [SLSBridgedGetSpaceManagementModeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedGetSpaceManagementModeOperation.MakeResultWithSpaceManagementMode]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceManagementModeOperation
type ISLSBridgedGetSpaceManagementModeOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSpaceManagementMode(mode uint64) objectivec.IObject
}

// Init initializes the instance.
func (s SLSBridgedGetSpaceManagementModeOperation) Init() SLSBridgedGetSpaceManagementModeOperation {
	rv := objc.Send[SLSBridgedGetSpaceManagementModeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedGetSpaceManagementModeOperation) Autorelease() SLSBridgedGetSpaceManagementModeOperation {
	rv := objc.Send[SLSBridgedGetSpaceManagementModeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedGetSpaceManagementModeOperation creates a new SLSBridgedGetSpaceManagementModeOperation instance.
func NewSLSBridgedGetSpaceManagementModeOperation() SLSBridgedGetSpaceManagementModeOperation {
	class := getSLSBridgedGetSpaceManagementModeOperationClass()
	rv := objc.Send[SLSBridgedGetSpaceManagementModeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceManagementModeOperation/initWithCoder:
func NewSLSBridgedGetSpaceManagementModeOperationWithCoder(coder objectivec.IObject) SLSBridgedGetSpaceManagementModeOperation {
	instance := getSLSBridgedGetSpaceManagementModeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedGetSpaceManagementModeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceManagementModeOperation/makeResultWithSpaceManagementMode:
func (s SLSBridgedGetSpaceManagementModeOperation) MakeResultWithSpaceManagementMode(mode uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceManagementMode:"), mode)
	return objectivec.Object{ID: rv}
}
