// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSPredicate] class.
var (
	_MPSPredicateClass     MPSPredicateClass
	_MPSPredicateClassOnce sync.Once
)

func getMPSPredicateClass() MPSPredicateClass {
	_MPSPredicateClassOnce.Do(func() {
		_MPSPredicateClass = MPSPredicateClass{class: objc.GetClass("MPSPredicate")}
	})
	return _MPSPredicateClass
}

// GetMPSPredicateClass returns the class object for MPSPredicate.
func GetMPSPredicateClass() MPSPredicateClass {
	return getMPSPredicateClass()
}

type MPSPredicateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSPredicateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSPredicateClass) Alloc() MPSPredicate {
	rv := objc.Send[MPSPredicate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSPredicate.InitWithBufferOffset]
//   - [MPSPredicate.InitWithDevice]
//
// # Instance Properties
//
//   - [MPSPredicate.PredicateBuffer]
//   - [MPSPredicate.PredicateOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate
type MPSPredicate struct {
	objectivec.Object
}

// MPSPredicateFromID constructs a [MPSPredicate] from an objc.ID.
func MPSPredicateFromID(id objc.ID) MPSPredicate {
	return MPSPredicate{objectivec.Object{ID: id}}
}

// NOTE: MPSPredicate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSPredicate] class.
//
// # Initializers
//
//   - [IMPSPredicate.InitWithBufferOffset]
//   - [IMPSPredicate.InitWithDevice]
//
// # Instance Properties
//
//   - [IMPSPredicate.PredicateBuffer]
//   - [IMPSPredicate.PredicateOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate
type IMPSPredicate interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithBufferOffset(buffer metal.MTLBuffer, offset uint) MPSPredicate
	InitWithDevice(device metal.MTLDevice) MPSPredicate

	// Topic: Instance Properties

	PredicateBuffer() metal.MTLBuffer
	PredicateOffset() uint
}

// Init initializes the instance.
func (p MPSPredicate) Init() MPSPredicate {
	rv := objc.Send[MPSPredicate](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSPredicate) Autorelease() MPSPredicate {
	rv := objc.Send[MPSPredicate](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSPredicate creates a new MPSPredicate instance.
func NewMPSPredicate() MPSPredicate {
	class := getMPSPredicateClass()
	rv := objc.Send[MPSPredicate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(buffer:offset:)
func NewPredicateWithBufferOffset(buffer metal.MTLBuffer, offset uint) MPSPredicate {
	instance := getMPSPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:"), buffer, offset)
	return MPSPredicateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(device:)
func NewPredicateWithDevice(device metal.MTLDevice) MPSPredicate {
	instance := getMPSPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSPredicateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(buffer:offset:)
func (p MPSPredicate) InitWithBufferOffset(buffer metal.MTLBuffer, offset uint) MPSPredicate {
	rv := objc.Send[MPSPredicate](p.ID, objc.Sel("initWithBuffer:offset:"), buffer, offset)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(device:)
func (p MPSPredicate) InitWithDevice(device metal.MTLDevice) MPSPredicate {
	rv := objc.Send[MPSPredicate](p.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/predicateWithBuffer:offset:
func (_MPSPredicateClass MPSPredicateClass) PredicateWithBufferOffset(buffer metal.MTLBuffer, offset uint) MPSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_MPSPredicateClass.class), objc.Sel("predicateWithBuffer:offset:"), buffer, offset)
	return MPSPredicateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/predicateBuffer
func (p MPSPredicate) PredicateBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("predicateBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/predicateOffset
func (p MPSPredicate) PredicateOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("predicateOffset"))
	return rv
}
