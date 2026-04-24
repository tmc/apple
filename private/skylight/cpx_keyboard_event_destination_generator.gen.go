// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyboardEventDestinationGenerator] class.
var (
	_CPXKeyboardEventDestinationGeneratorClass     CPXKeyboardEventDestinationGeneratorClass
	_CPXKeyboardEventDestinationGeneratorClassOnce sync.Once
)

func getCPXKeyboardEventDestinationGeneratorClass() CPXKeyboardEventDestinationGeneratorClass {
	_CPXKeyboardEventDestinationGeneratorClassOnce.Do(func() {
		_CPXKeyboardEventDestinationGeneratorClass = CPXKeyboardEventDestinationGeneratorClass{class: objc.GetClass("CPXKeyboardEventDestinationGenerator")}
	})
	return _CPXKeyboardEventDestinationGeneratorClass
}

// GetCPXKeyboardEventDestinationGeneratorClass returns the class object for CPXKeyboardEventDestinationGenerator.
func GetCPXKeyboardEventDestinationGeneratorClass() CPXKeyboardEventDestinationGeneratorClass {
	return getCPXKeyboardEventDestinationGeneratorClass()
}

type CPXKeyboardEventDestinationGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyboardEventDestinationGeneratorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyboardEventDestinationGeneratorClass) Alloc() CPXKeyboardEventDestinationGenerator {
	rv := objc.Send[CPXKeyboardEventDestinationGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXKeyboardEventDestinationGenerator.DestinationForEventContextTargetExtras]
//   - [CPXKeyboardEventDestinationGenerator.InitWithDeliveryManagerFocusManagerSequenceTracker]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventDestinationGenerator
type CPXKeyboardEventDestinationGenerator struct {
	objectivec.Object
}

// CPXKeyboardEventDestinationGeneratorFromID constructs a [CPXKeyboardEventDestinationGenerator] from an objc.ID.
func CPXKeyboardEventDestinationGeneratorFromID(id objc.ID) CPXKeyboardEventDestinationGenerator {
	return CPXKeyboardEventDestinationGenerator{objectivec.Object{ID: id}}
}

// Ensure CPXKeyboardEventDestinationGenerator implements ICPXKeyboardEventDestinationGenerator.
var _ ICPXKeyboardEventDestinationGenerator = CPXKeyboardEventDestinationGenerator{}

// An interface definition for the [CPXKeyboardEventDestinationGenerator] class.
//
// # Methods
//
//   - [ICPXKeyboardEventDestinationGenerator.DestinationForEventContextTargetExtras]
//   - [ICPXKeyboardEventDestinationGenerator.InitWithDeliveryManagerFocusManagerSequenceTracker]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventDestinationGenerator
type ICPXKeyboardEventDestinationGenerator interface {
	objectivec.IObject

	// Topic: Methods

	DestinationForEventContextTargetExtras(event *SLSEventRecordRef, target *CPSProcessRecRef, extras []objectivec.IObject) objectivec.IObject
	InitWithDeliveryManagerFocusManagerSequenceTracker(manager objectivec.IObject, manager2 objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventDestinationGenerator
}

// Init initializes the instance.
func (c CPXKeyboardEventDestinationGenerator) Init() CPXKeyboardEventDestinationGenerator {
	rv := objc.Send[CPXKeyboardEventDestinationGenerator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyboardEventDestinationGenerator) Autorelease() CPXKeyboardEventDestinationGenerator {
	rv := objc.Send[CPXKeyboardEventDestinationGenerator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyboardEventDestinationGenerator creates a new CPXKeyboardEventDestinationGenerator instance.
func NewCPXKeyboardEventDestinationGenerator() CPXKeyboardEventDestinationGenerator {
	class := getCPXKeyboardEventDestinationGeneratorClass()
	rv := objc.Send[CPXKeyboardEventDestinationGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventDestinationGenerator/initWithDeliveryManager:focusManager:sequenceTracker:
func NewCPXKeyboardEventDestinationGeneratorWithDeliveryManagerFocusManagerSequenceTracker(manager objectivec.IObject, manager2 objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventDestinationGenerator {
	instance := getCPXKeyboardEventDestinationGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:focusManager:sequenceTracker:"), manager, manager2, tracker)
	return CPXKeyboardEventDestinationGeneratorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventDestinationGenerator/destinationForEvent:contextTarget:extras:
func (c CPXKeyboardEventDestinationGenerator) DestinationForEventContextTargetExtras(event *SLSEventRecordRef, target *CPSProcessRecRef, extras []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationForEvent:contextTarget:extras:"), event, target, objectivec.IObjectSliceToNSArray(extras))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventDestinationGenerator/initWithDeliveryManager:focusManager:sequenceTracker:
func (c CPXKeyboardEventDestinationGenerator) InitWithDeliveryManagerFocusManagerSequenceTracker(manager objectivec.IObject, manager2 objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventDestinationGenerator {
	rv := objc.Send[CPXKeyboardEventDestinationGenerator](c.ID, objc.Sel("initWithDeliveryManager:focusManager:sequenceTracker:"), manager, manager2, tracker)
	return rv
}
