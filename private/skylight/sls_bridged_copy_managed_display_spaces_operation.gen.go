// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyManagedDisplaySpacesOperation] class.
var (
	_SLSBridgedCopyManagedDisplaySpacesOperationClass     SLSBridgedCopyManagedDisplaySpacesOperationClass
	_SLSBridgedCopyManagedDisplaySpacesOperationClassOnce sync.Once
)

func getSLSBridgedCopyManagedDisplaySpacesOperationClass() SLSBridgedCopyManagedDisplaySpacesOperationClass {
	_SLSBridgedCopyManagedDisplaySpacesOperationClassOnce.Do(func() {
		_SLSBridgedCopyManagedDisplaySpacesOperationClass = SLSBridgedCopyManagedDisplaySpacesOperationClass{class: objc.GetClass("SLSBridgedCopyManagedDisplaySpacesOperation")}
	})
	return _SLSBridgedCopyManagedDisplaySpacesOperationClass
}

// GetSLSBridgedCopyManagedDisplaySpacesOperationClass returns the class object for SLSBridgedCopyManagedDisplaySpacesOperation.
func GetSLSBridgedCopyManagedDisplaySpacesOperationClass() SLSBridgedCopyManagedDisplaySpacesOperationClass {
	return getSLSBridgedCopyManagedDisplaySpacesOperationClass()
}

type SLSBridgedCopyManagedDisplaySpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyManagedDisplaySpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyManagedDisplaySpacesOperationClass) Alloc() SLSBridgedCopyManagedDisplaySpacesOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaySpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyManagedDisplaySpacesOperation.MakeResultWithPropertyListArray]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaySpacesOperation
type SLSBridgedCopyManagedDisplaySpacesOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyManagedDisplaySpacesOperationFromID constructs a [SLSBridgedCopyManagedDisplaySpacesOperation] from an objc.ID.
func SLSBridgedCopyManagedDisplaySpacesOperationFromID(id objc.ID) SLSBridgedCopyManagedDisplaySpacesOperation {
	return SLSBridgedCopyManagedDisplaySpacesOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyManagedDisplaySpacesOperation implements ISLSBridgedCopyManagedDisplaySpacesOperation.
var _ ISLSBridgedCopyManagedDisplaySpacesOperation = SLSBridgedCopyManagedDisplaySpacesOperation{}

// An interface definition for the [SLSBridgedCopyManagedDisplaySpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyManagedDisplaySpacesOperation.MakeResultWithPropertyListArray]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaySpacesOperation
type ISLSBridgedCopyManagedDisplaySpacesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithPropertyListArray(array objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (s SLSBridgedCopyManagedDisplaySpacesOperation) Init() SLSBridgedCopyManagedDisplaySpacesOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaySpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyManagedDisplaySpacesOperation) Autorelease() SLSBridgedCopyManagedDisplaySpacesOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaySpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyManagedDisplaySpacesOperation creates a new SLSBridgedCopyManagedDisplaySpacesOperation instance.
func NewSLSBridgedCopyManagedDisplaySpacesOperation() SLSBridgedCopyManagedDisplaySpacesOperation {
	class := getSLSBridgedCopyManagedDisplaySpacesOperationClass()
	rv := objc.Send[SLSBridgedCopyManagedDisplaySpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaySpacesOperation/initWithCoder:
func NewSLSBridgedCopyManagedDisplaySpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyManagedDisplaySpacesOperation {
	instance := getSLSBridgedCopyManagedDisplaySpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyManagedDisplaySpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaySpacesOperation/makeResultWithPropertyListArray:
func (s SLSBridgedCopyManagedDisplaySpacesOperation) MakeResultWithPropertyListArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithPropertyListArray:"), array)
	return objectivec.Object{ID: rv}
}
