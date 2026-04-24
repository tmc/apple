// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplaySetIsAnimatingOperation] class.
var (
	_SLSBridgedManagedDisplaySetIsAnimatingOperationClass     SLSBridgedManagedDisplaySetIsAnimatingOperationClass
	_SLSBridgedManagedDisplaySetIsAnimatingOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplaySetIsAnimatingOperationClass() SLSBridgedManagedDisplaySetIsAnimatingOperationClass {
	_SLSBridgedManagedDisplaySetIsAnimatingOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplaySetIsAnimatingOperationClass = SLSBridgedManagedDisplaySetIsAnimatingOperationClass{class: objc.GetClass("SLSBridgedManagedDisplaySetIsAnimatingOperation")}
	})
	return _SLSBridgedManagedDisplaySetIsAnimatingOperationClass
}

// GetSLSBridgedManagedDisplaySetIsAnimatingOperationClass returns the class object for SLSBridgedManagedDisplaySetIsAnimatingOperation.
func GetSLSBridgedManagedDisplaySetIsAnimatingOperationClass() SLSBridgedManagedDisplaySetIsAnimatingOperationClass {
	return getSLSBridgedManagedDisplaySetIsAnimatingOperationClass()
}

type SLSBridgedManagedDisplaySetIsAnimatingOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplaySetIsAnimatingOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplaySetIsAnimatingOperationClass) Alloc() SLSBridgedManagedDisplaySetIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetIsAnimatingOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplaySetIsAnimatingOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplaySetIsAnimatingOperation.IsAnimating]
//   - [SLSBridgedManagedDisplaySetIsAnimatingOperation.InitWithDisplayIdentifierIsAnimating]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation
type SLSBridgedManagedDisplaySetIsAnimatingOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplaySetIsAnimatingOperationFromID constructs a [SLSBridgedManagedDisplaySetIsAnimatingOperation] from an objc.ID.
func SLSBridgedManagedDisplaySetIsAnimatingOperationFromID(id objc.ID) SLSBridgedManagedDisplaySetIsAnimatingOperation {
	return SLSBridgedManagedDisplaySetIsAnimatingOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplaySetIsAnimatingOperation implements ISLSBridgedManagedDisplaySetIsAnimatingOperation.
var _ ISLSBridgedManagedDisplaySetIsAnimatingOperation = SLSBridgedManagedDisplaySetIsAnimatingOperation{}

// An interface definition for the [SLSBridgedManagedDisplaySetIsAnimatingOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplaySetIsAnimatingOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplaySetIsAnimatingOperation.IsAnimating]
//   - [ISLSBridgedManagedDisplaySetIsAnimatingOperation.InitWithDisplayIdentifierIsAnimating]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation
type ISLSBridgedManagedDisplaySetIsAnimatingOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	IsAnimating() bool
	InitWithDisplayIdentifierIsAnimating(identifier objectivec.IObject, animating bool) SLSBridgedManagedDisplaySetIsAnimatingOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplaySetIsAnimatingOperation) Init() SLSBridgedManagedDisplaySetIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetIsAnimatingOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplaySetIsAnimatingOperation) Autorelease() SLSBridgedManagedDisplaySetIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetIsAnimatingOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplaySetIsAnimatingOperation creates a new SLSBridgedManagedDisplaySetIsAnimatingOperation instance.
func NewSLSBridgedManagedDisplaySetIsAnimatingOperation() SLSBridgedManagedDisplaySetIsAnimatingOperation {
	class := getSLSBridgedManagedDisplaySetIsAnimatingOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplaySetIsAnimatingOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation/initWithCoder:
func NewSLSBridgedManagedDisplaySetIsAnimatingOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplaySetIsAnimatingOperation {
	instance := getSLSBridgedManagedDisplaySetIsAnimatingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplaySetIsAnimatingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation/initWithDisplayIdentifier:isAnimating:
func NewSLSBridgedManagedDisplaySetIsAnimatingOperationWithDisplayIdentifierIsAnimating(identifier objectivec.IObject, animating bool) SLSBridgedManagedDisplaySetIsAnimatingOperation {
	instance := getSLSBridgedManagedDisplaySetIsAnimatingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:isAnimating:"), identifier, animating)
	return SLSBridgedManagedDisplaySetIsAnimatingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation/initWithDisplayIdentifier:isAnimating:
func (s SLSBridgedManagedDisplaySetIsAnimatingOperation) InitWithDisplayIdentifierIsAnimating(identifier objectivec.IObject, animating bool) SLSBridgedManagedDisplaySetIsAnimatingOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetIsAnimatingOperation](s.ID, objc.Sel("initWithDisplayIdentifier:isAnimating:"), identifier, animating)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation/displayIdentifier
func (s SLSBridgedManagedDisplaySetIsAnimatingOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetIsAnimatingOperation/isAnimating
func (s SLSBridgedManagedDisplaySetIsAnimatingOperation) IsAnimating() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAnimating"))
	return rv
}
