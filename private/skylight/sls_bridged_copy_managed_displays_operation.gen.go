// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyManagedDisplaysOperation] class.
var (
	_SLSBridgedCopyManagedDisplaysOperationClass     SLSBridgedCopyManagedDisplaysOperationClass
	_SLSBridgedCopyManagedDisplaysOperationClassOnce sync.Once
)

func getSLSBridgedCopyManagedDisplaysOperationClass() SLSBridgedCopyManagedDisplaysOperationClass {
	_SLSBridgedCopyManagedDisplaysOperationClassOnce.Do(func() {
		_SLSBridgedCopyManagedDisplaysOperationClass = SLSBridgedCopyManagedDisplaysOperationClass{class: objc.GetClass("SLSBridgedCopyManagedDisplaysOperation")}
	})
	return _SLSBridgedCopyManagedDisplaysOperationClass
}

// GetSLSBridgedCopyManagedDisplaysOperationClass returns the class object for SLSBridgedCopyManagedDisplaysOperation.
func GetSLSBridgedCopyManagedDisplaysOperationClass() SLSBridgedCopyManagedDisplaysOperationClass {
	return getSLSBridgedCopyManagedDisplaysOperationClass()
}

type SLSBridgedCopyManagedDisplaysOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyManagedDisplaysOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyManagedDisplaysOperationClass) Alloc() SLSBridgedCopyManagedDisplaysOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaysOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyManagedDisplaysOperation.MakeResultWithStrings]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaysOperation
type SLSBridgedCopyManagedDisplaysOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyManagedDisplaysOperationFromID constructs a [SLSBridgedCopyManagedDisplaysOperation] from an objc.ID.
func SLSBridgedCopyManagedDisplaysOperationFromID(id objc.ID) SLSBridgedCopyManagedDisplaysOperation {
	return SLSBridgedCopyManagedDisplaysOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyManagedDisplaysOperation implements ISLSBridgedCopyManagedDisplaysOperation.
var _ ISLSBridgedCopyManagedDisplaysOperation = SLSBridgedCopyManagedDisplaysOperation{}

// An interface definition for the [SLSBridgedCopyManagedDisplaysOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyManagedDisplaysOperation.MakeResultWithStrings]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaysOperation
type ISLSBridgedCopyManagedDisplaysOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithStrings(strings objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (s SLSBridgedCopyManagedDisplaysOperation) Init() SLSBridgedCopyManagedDisplaysOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaysOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyManagedDisplaysOperation) Autorelease() SLSBridgedCopyManagedDisplaysOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplaysOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyManagedDisplaysOperation creates a new SLSBridgedCopyManagedDisplaysOperation instance.
func NewSLSBridgedCopyManagedDisplaysOperation() SLSBridgedCopyManagedDisplaysOperation {
	class := getSLSBridgedCopyManagedDisplaysOperationClass()
	rv := objc.Send[SLSBridgedCopyManagedDisplaysOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaysOperation/initWithCoder:
func NewSLSBridgedCopyManagedDisplaysOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyManagedDisplaysOperation {
	instance := getSLSBridgedCopyManagedDisplaysOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyManagedDisplaysOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplaysOperation/makeResultWithStrings:
func (s SLSBridgedCopyManagedDisplaysOperation) MakeResultWithStrings(strings objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithStrings:"), strings)
	return objectivec.Object{ID: rv}
}
