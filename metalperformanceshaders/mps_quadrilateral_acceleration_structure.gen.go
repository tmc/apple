// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSQuadrilateralAccelerationStructure] class.
var (
	_MPSQuadrilateralAccelerationStructureClass     MPSQuadrilateralAccelerationStructureClass
	_MPSQuadrilateralAccelerationStructureClassOnce sync.Once
)

func getMPSQuadrilateralAccelerationStructureClass() MPSQuadrilateralAccelerationStructureClass {
	_MPSQuadrilateralAccelerationStructureClassOnce.Do(func() {
		_MPSQuadrilateralAccelerationStructureClass = MPSQuadrilateralAccelerationStructureClass{class: objc.GetClass("MPSQuadrilateralAccelerationStructure")}
	})
	return _MPSQuadrilateralAccelerationStructureClass
}

// GetMPSQuadrilateralAccelerationStructureClass returns the class object for MPSQuadrilateralAccelerationStructure.
func GetMPSQuadrilateralAccelerationStructureClass() MPSQuadrilateralAccelerationStructureClass {
	return getMPSQuadrilateralAccelerationStructureClass()
}

type MPSQuadrilateralAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSQuadrilateralAccelerationStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSQuadrilateralAccelerationStructureClass) Alloc() MPSQuadrilateralAccelerationStructure {
	rv := objc.Send[MPSQuadrilateralAccelerationStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSQuadrilateralAccelerationStructure.QuadrilateralCount]
//   - [MPSQuadrilateralAccelerationStructure.SetQuadrilateralCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure
type MPSQuadrilateralAccelerationStructure struct {
	MPSPolygonAccelerationStructure
}

// MPSQuadrilateralAccelerationStructureFromID constructs a [MPSQuadrilateralAccelerationStructure] from an objc.ID.
func MPSQuadrilateralAccelerationStructureFromID(id objc.ID) MPSQuadrilateralAccelerationStructure {
	return MPSQuadrilateralAccelerationStructure{MPSPolygonAccelerationStructure: MPSPolygonAccelerationStructureFromID(id)}
}

// NOTE: MPSQuadrilateralAccelerationStructure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSQuadrilateralAccelerationStructure] class.
//
// # Instance Properties
//
//   - [IMPSQuadrilateralAccelerationStructure.QuadrilateralCount]
//   - [IMPSQuadrilateralAccelerationStructure.SetQuadrilateralCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure
type IMPSQuadrilateralAccelerationStructure interface {
	IMPSPolygonAccelerationStructure

	// Topic: Instance Properties

	QuadrilateralCount() uint
	SetQuadrilateralCount(value uint)
}

// Init initializes the instance.
func (q MPSQuadrilateralAccelerationStructure) Init() MPSQuadrilateralAccelerationStructure {
	rv := objc.Send[MPSQuadrilateralAccelerationStructure](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q MPSQuadrilateralAccelerationStructure) Autorelease() MPSQuadrilateralAccelerationStructure {
	rv := objc.Send[MPSQuadrilateralAccelerationStructure](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSQuadrilateralAccelerationStructure creates a new MPSQuadrilateralAccelerationStructure instance.
func NewMPSQuadrilateralAccelerationStructure() MPSQuadrilateralAccelerationStructure {
	class := getMPSQuadrilateralAccelerationStructureClass()
	rv := objc.Send[MPSQuadrilateralAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewQuadrilateralAccelerationStructureWithCoder(aDecoder foundation.INSCoder) MPSQuadrilateralAccelerationStructure {
	instance := getMPSQuadrilateralAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSQuadrilateralAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:device:)
func NewQuadrilateralAccelerationStructureWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSQuadrilateralAccelerationStructure {
	instance := getMPSQuadrilateralAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSQuadrilateralAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:group:)
func NewQuadrilateralAccelerationStructureWithCoderGroup(aDecoder foundation.INSCoder, group IMPSAccelerationStructureGroup) MPSQuadrilateralAccelerationStructure {
	instance := getMPSQuadrilateralAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	return MPSQuadrilateralAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(device:)
func NewQuadrilateralAccelerationStructureWithDevice(device metal.MTLDevice) MPSQuadrilateralAccelerationStructure {
	instance := getMPSQuadrilateralAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSQuadrilateralAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewQuadrilateralAccelerationStructureWithGroup(group IMPSAccelerationStructureGroup) MPSQuadrilateralAccelerationStructure {
	instance := getMPSQuadrilateralAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroup:"), group)
	return MPSQuadrilateralAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure/quadrilateralCount
func (q MPSQuadrilateralAccelerationStructure) QuadrilateralCount() uint {
	rv := objc.Send[uint](q.ID, objc.Sel("quadrilateralCount"))
	return rv
}
func (q MPSQuadrilateralAccelerationStructure) SetQuadrilateralCount(value uint) {
	objc.Send[struct{}](q.ID, objc.Sel("setQuadrilateralCount:"), value)
}
