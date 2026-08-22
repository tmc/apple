// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSAccelerationStructure] class.
var (
	_MPSAccelerationStructureClass     MPSAccelerationStructureClass
	_MPSAccelerationStructureClassOnce sync.Once
)

func getMPSAccelerationStructureClass() MPSAccelerationStructureClass {
	_MPSAccelerationStructureClassOnce.Do(func() {
		_MPSAccelerationStructureClass = MPSAccelerationStructureClass{class: objc.GetClass("MPSAccelerationStructure")}
	})
	return _MPSAccelerationStructureClass
}

// GetMPSAccelerationStructureClass returns the class object for MPSAccelerationStructure.
func GetMPSAccelerationStructureClass() MPSAccelerationStructureClass {
	return getMPSAccelerationStructureClass()
}

type MPSAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSAccelerationStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSAccelerationStructureClass) Alloc() MPSAccelerationStructure {
	rv := objc.Send[MPSAccelerationStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for data structures that are built over geometry and used to
// accelerate ray tracing.
//
// # Instance Properties
//
//   - [MPSAccelerationStructure.BoundingBox]
//   - [MPSAccelerationStructure.Group]
//   - [MPSAccelerationStructure.Status]
//   - [MPSAccelerationStructure.Usage]
//   - [MPSAccelerationStructure.SetUsage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure
type MPSAccelerationStructure struct {
	MPSKernel
}

// MPSAccelerationStructureFromID constructs a [MPSAccelerationStructure] from an objc.ID.
//
// The base class for data structures that are built over geometry and used to
// accelerate ray tracing.
func MPSAccelerationStructureFromID(id objc.ID) MPSAccelerationStructure {
	return MPSAccelerationStructure{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSAccelerationStructure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSAccelerationStructure] class.
//
// # Instance Properties
//
//   - [IMPSAccelerationStructure.BoundingBox]
//   - [IMPSAccelerationStructure.Group]
//   - [IMPSAccelerationStructure.Status]
//   - [IMPSAccelerationStructure.Usage]
//   - [IMPSAccelerationStructure.SetUsage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure
type IMPSAccelerationStructure interface {
	IMPSKernel

	// Topic: Instance Properties

	BoundingBox() MPSAxisAlignedBoundingBox
	Group() IMPSAccelerationStructureGroup
	Status() MPSAccelerationStructureStatus
	Usage() MPSAccelerationStructureUsage
	SetUsage(value MPSAccelerationStructureUsage)
}

// Init initializes the instance.
func (a MPSAccelerationStructure) Init() MPSAccelerationStructure {
	rv := objc.Send[MPSAccelerationStructure](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSAccelerationStructure) Autorelease() MPSAccelerationStructure {
	rv := objc.Send[MPSAccelerationStructure](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSAccelerationStructure creates a new MPSAccelerationStructure instance.
func NewMPSAccelerationStructure() MPSAccelerationStructure {
	class := getMPSAccelerationStructureClass()
	rv := objc.Send[MPSAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewAccelerationStructureWithCoder(aDecoder foundation.INSCoder) MPSAccelerationStructure {
	instance := getMPSAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:device:)
func NewAccelerationStructureWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSAccelerationStructure {
	instance := getMPSAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:group:)
func NewAccelerationStructureWithCoderGroup(aDecoder foundation.INSCoder, group IMPSAccelerationStructureGroup) MPSAccelerationStructure {
	instance := getMPSAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	return MPSAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(device:)
func NewAccelerationStructureWithDevice(device metal.MTLDevice) MPSAccelerationStructure {
	instance := getMPSAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewAccelerationStructureWithGroup(group IMPSAccelerationStructureGroup) MPSAccelerationStructure {
	instance := getMPSAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroup:"), group)
	return MPSAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/boundingBox
func (a MPSAccelerationStructure) BoundingBox() MPSAxisAlignedBoundingBox {
	rv := objc.Send[MPSAxisAlignedBoundingBox](a.ID, objc.Sel("boundingBox"))
	return MPSAxisAlignedBoundingBox(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/group
func (a MPSAccelerationStructure) Group() IMPSAccelerationStructureGroup {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("group"))
	return MPSAccelerationStructureGroupFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/status
func (a MPSAccelerationStructure) Status() MPSAccelerationStructureStatus {
	rv := objc.Send[MPSAccelerationStructureStatus](a.ID, objc.Sel("status"))
	return MPSAccelerationStructureStatus(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/usage
func (a MPSAccelerationStructure) Usage() MPSAccelerationStructureUsage {
	rv := objc.Send[MPSAccelerationStructureUsage](a.ID, objc.Sel("usage"))
	return MPSAccelerationStructureUsage(rv)
}
func (a MPSAccelerationStructure) SetUsage(value MPSAccelerationStructureUsage) {
	objc.Send[struct{}](a.ID, objc.Sel("setUsage:"), value)
}
