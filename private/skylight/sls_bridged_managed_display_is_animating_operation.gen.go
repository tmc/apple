// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplayIsAnimatingOperation] class.
var (
	_SLSBridgedManagedDisplayIsAnimatingOperationClass     SLSBridgedManagedDisplayIsAnimatingOperationClass
	_SLSBridgedManagedDisplayIsAnimatingOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplayIsAnimatingOperationClass() SLSBridgedManagedDisplayIsAnimatingOperationClass {
	_SLSBridgedManagedDisplayIsAnimatingOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplayIsAnimatingOperationClass = SLSBridgedManagedDisplayIsAnimatingOperationClass{class: objc.GetClass("SLSBridgedManagedDisplayIsAnimatingOperation")}
	})
	return _SLSBridgedManagedDisplayIsAnimatingOperationClass
}

// GetSLSBridgedManagedDisplayIsAnimatingOperationClass returns the class object for SLSBridgedManagedDisplayIsAnimatingOperation.
func GetSLSBridgedManagedDisplayIsAnimatingOperationClass() SLSBridgedManagedDisplayIsAnimatingOperationClass {
	return getSLSBridgedManagedDisplayIsAnimatingOperationClass()
}

type SLSBridgedManagedDisplayIsAnimatingOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplayIsAnimatingOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplayIsAnimatingOperationClass) Alloc() SLSBridgedManagedDisplayIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplayIsAnimatingOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplayIsAnimatingOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplayIsAnimatingOperation.MakeResultWithBoolValue]
//   - [SLSBridgedManagedDisplayIsAnimatingOperation.InitWithDisplayIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation
type SLSBridgedManagedDisplayIsAnimatingOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplayIsAnimatingOperationFromID constructs a [SLSBridgedManagedDisplayIsAnimatingOperation] from an objc.ID.
func SLSBridgedManagedDisplayIsAnimatingOperationFromID(id objc.ID) SLSBridgedManagedDisplayIsAnimatingOperation {
	return SLSBridgedManagedDisplayIsAnimatingOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplayIsAnimatingOperation implements ISLSBridgedManagedDisplayIsAnimatingOperation.
var _ ISLSBridgedManagedDisplayIsAnimatingOperation = SLSBridgedManagedDisplayIsAnimatingOperation{}

// An interface definition for the [SLSBridgedManagedDisplayIsAnimatingOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplayIsAnimatingOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplayIsAnimatingOperation.MakeResultWithBoolValue]
//   - [ISLSBridgedManagedDisplayIsAnimatingOperation.InitWithDisplayIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation
type ISLSBridgedManagedDisplayIsAnimatingOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	MakeResultWithBoolValue(value bool) objectivec.IObject
	InitWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayIsAnimatingOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplayIsAnimatingOperation) Init() SLSBridgedManagedDisplayIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplayIsAnimatingOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplayIsAnimatingOperation) Autorelease() SLSBridgedManagedDisplayIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplayIsAnimatingOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplayIsAnimatingOperation creates a new SLSBridgedManagedDisplayIsAnimatingOperation instance.
func NewSLSBridgedManagedDisplayIsAnimatingOperation() SLSBridgedManagedDisplayIsAnimatingOperation {
	class := getSLSBridgedManagedDisplayIsAnimatingOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplayIsAnimatingOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation/initWithCoder:
func NewSLSBridgedManagedDisplayIsAnimatingOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplayIsAnimatingOperation {
	instance := getSLSBridgedManagedDisplayIsAnimatingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplayIsAnimatingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation/initWithDisplayIdentifier:
func NewSLSBridgedManagedDisplayIsAnimatingOperationWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayIsAnimatingOperation {
	instance := getSLSBridgedManagedDisplayIsAnimatingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:"), identifier)
	return SLSBridgedManagedDisplayIsAnimatingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation/makeResultWithBoolValue:
func (s SLSBridgedManagedDisplayIsAnimatingOperation) MakeResultWithBoolValue(value bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithBoolValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation/initWithDisplayIdentifier:
func (s SLSBridgedManagedDisplayIsAnimatingOperation) InitWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplayIsAnimatingOperation](s.ID, objc.Sel("initWithDisplayIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayIsAnimatingOperation/displayIdentifier
func (s SLSBridgedManagedDisplayIsAnimatingOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
