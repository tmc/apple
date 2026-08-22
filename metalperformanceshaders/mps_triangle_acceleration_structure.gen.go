// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSTriangleAccelerationStructure] class.
var (
	_MPSTriangleAccelerationStructureClass     MPSTriangleAccelerationStructureClass
	_MPSTriangleAccelerationStructureClassOnce sync.Once
)

func getMPSTriangleAccelerationStructureClass() MPSTriangleAccelerationStructureClass {
	_MPSTriangleAccelerationStructureClassOnce.Do(func() {
		_MPSTriangleAccelerationStructureClass = MPSTriangleAccelerationStructureClass{class: objc.GetClass("MPSTriangleAccelerationStructure")}
	})
	return _MPSTriangleAccelerationStructureClass
}

// GetMPSTriangleAccelerationStructureClass returns the class object for MPSTriangleAccelerationStructure.
func GetMPSTriangleAccelerationStructureClass() MPSTriangleAccelerationStructureClass {
	return getMPSTriangleAccelerationStructureClass()
}

type MPSTriangleAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTriangleAccelerationStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTriangleAccelerationStructureClass) Alloc() MPSTriangleAccelerationStructure {
	rv := objc.Send[MPSTriangleAccelerationStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An acceleration structure built over triangles.
//
// # Instance Properties
//
//   - [MPSTriangleAccelerationStructure.TriangleCount]
//   - [MPSTriangleAccelerationStructure.SetTriangleCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure
type MPSTriangleAccelerationStructure struct {
	MPSPolygonAccelerationStructure
}

// MPSTriangleAccelerationStructureFromID constructs a [MPSTriangleAccelerationStructure] from an objc.ID.
//
// An acceleration structure built over triangles.
func MPSTriangleAccelerationStructureFromID(id objc.ID) MPSTriangleAccelerationStructure {
	return MPSTriangleAccelerationStructure{MPSPolygonAccelerationStructure: MPSPolygonAccelerationStructureFromID(id)}
}

// NOTE: MPSTriangleAccelerationStructure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTriangleAccelerationStructure] class.
//
// # Instance Properties
//
//   - [IMPSTriangleAccelerationStructure.TriangleCount]
//   - [IMPSTriangleAccelerationStructure.SetTriangleCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure
type IMPSTriangleAccelerationStructure interface {
	IMPSPolygonAccelerationStructure

	// Topic: Instance Properties

	TriangleCount() uint
	SetTriangleCount(value uint)
}

// Init initializes the instance.
func (t MPSTriangleAccelerationStructure) Init() MPSTriangleAccelerationStructure {
	rv := objc.Send[MPSTriangleAccelerationStructure](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTriangleAccelerationStructure) Autorelease() MPSTriangleAccelerationStructure {
	rv := objc.Send[MPSTriangleAccelerationStructure](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTriangleAccelerationStructure creates a new MPSTriangleAccelerationStructure instance.
func NewMPSTriangleAccelerationStructure() MPSTriangleAccelerationStructure {
	class := getMPSTriangleAccelerationStructureClass()
	rv := objc.Send[MPSTriangleAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewTriangleAccelerationStructureWithCoder(aDecoder foundation.INSCoder) MPSTriangleAccelerationStructure {
	instance := getMPSTriangleAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSTriangleAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:device:)
func NewTriangleAccelerationStructureWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSTriangleAccelerationStructure {
	instance := getMPSTriangleAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSTriangleAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:group:)
func NewTriangleAccelerationStructureWithCoderGroup(aDecoder foundation.INSCoder, group IMPSAccelerationStructureGroup) MPSTriangleAccelerationStructure {
	instance := getMPSTriangleAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	return MPSTriangleAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(device:)
func NewTriangleAccelerationStructureWithDevice(device metal.MTLDevice) MPSTriangleAccelerationStructure {
	instance := getMPSTriangleAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSTriangleAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewTriangleAccelerationStructureWithGroup(group IMPSAccelerationStructureGroup) MPSTriangleAccelerationStructure {
	instance := getMPSTriangleAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroup:"), group)
	return MPSTriangleAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure/triangleCount
func (t MPSTriangleAccelerationStructure) TriangleCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("triangleCount"))
	return rv
}
func (t MPSTriangleAccelerationStructure) SetTriangleCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setTriangleCount:"), value)
}
