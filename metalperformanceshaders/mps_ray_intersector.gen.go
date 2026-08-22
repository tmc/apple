// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSRayIntersector] class.
var (
	_MPSRayIntersectorClass     MPSRayIntersectorClass
	_MPSRayIntersectorClassOnce sync.Once
)

func getMPSRayIntersectorClass() MPSRayIntersectorClass {
	_MPSRayIntersectorClassOnce.Do(func() {
		_MPSRayIntersectorClass = MPSRayIntersectorClass{class: objc.GetClass("MPSRayIntersector")}
	})
	return _MPSRayIntersectorClass
}

// GetMPSRayIntersectorClass returns the class object for MPSRayIntersector.
func GetMPSRayIntersectorClass() MPSRayIntersectorClass {
	return getMPSRayIntersectorClass()
}

type MPSRayIntersectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRayIntersectorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRayIntersectorClass) Alloc() MPSRayIntersector {
	rv := objc.Send[MPSRayIntersector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that performs intersection tests between rays and geometry.
//
// # Instance Properties
//
//   - [MPSRayIntersector.BoundingBoxIntersectionTestType]
//   - [MPSRayIntersector.SetBoundingBoxIntersectionTestType]
//   - [MPSRayIntersector.CullMode]
//   - [MPSRayIntersector.SetCullMode]
//   - [MPSRayIntersector.FrontFacingWinding]
//   - [MPSRayIntersector.SetFrontFacingWinding]
//   - [MPSRayIntersector.IntersectionDataType]
//   - [MPSRayIntersector.SetIntersectionDataType]
//   - [MPSRayIntersector.IntersectionStride]
//   - [MPSRayIntersector.SetIntersectionStride]
//   - [MPSRayIntersector.RayDataType]
//   - [MPSRayIntersector.SetRayDataType]
//   - [MPSRayIntersector.RayIndexDataType]
//   - [MPSRayIntersector.SetRayIndexDataType]
//   - [MPSRayIntersector.RayMask]
//   - [MPSRayIntersector.SetRayMask]
//   - [MPSRayIntersector.RayMaskOperator]
//   - [MPSRayIntersector.SetRayMaskOperator]
//   - [MPSRayIntersector.RayMaskOptions]
//   - [MPSRayIntersector.SetRayMaskOptions]
//   - [MPSRayIntersector.RayStride]
//   - [MPSRayIntersector.SetRayStride]
//   - [MPSRayIntersector.TriangleIntersectionTestType]
//   - [MPSRayIntersector.SetTriangleIntersectionTestType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector
type MPSRayIntersector struct {
	MPSKernel
}

// MPSRayIntersectorFromID constructs a [MPSRayIntersector] from an objc.ID.
//
// A kernel that performs intersection tests between rays and geometry.
func MPSRayIntersectorFromID(id objc.ID) MPSRayIntersector {
	return MPSRayIntersector{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSRayIntersector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRayIntersector] class.
//
// # Instance Properties
//
//   - [IMPSRayIntersector.BoundingBoxIntersectionTestType]
//   - [IMPSRayIntersector.SetBoundingBoxIntersectionTestType]
//   - [IMPSRayIntersector.CullMode]
//   - [IMPSRayIntersector.SetCullMode]
//   - [IMPSRayIntersector.FrontFacingWinding]
//   - [IMPSRayIntersector.SetFrontFacingWinding]
//   - [IMPSRayIntersector.IntersectionDataType]
//   - [IMPSRayIntersector.SetIntersectionDataType]
//   - [IMPSRayIntersector.IntersectionStride]
//   - [IMPSRayIntersector.SetIntersectionStride]
//   - [IMPSRayIntersector.RayDataType]
//   - [IMPSRayIntersector.SetRayDataType]
//   - [IMPSRayIntersector.RayIndexDataType]
//   - [IMPSRayIntersector.SetRayIndexDataType]
//   - [IMPSRayIntersector.RayMask]
//   - [IMPSRayIntersector.SetRayMask]
//   - [IMPSRayIntersector.RayMaskOperator]
//   - [IMPSRayIntersector.SetRayMaskOperator]
//   - [IMPSRayIntersector.RayMaskOptions]
//   - [IMPSRayIntersector.SetRayMaskOptions]
//   - [IMPSRayIntersector.RayStride]
//   - [IMPSRayIntersector.SetRayStride]
//   - [IMPSRayIntersector.TriangleIntersectionTestType]
//   - [IMPSRayIntersector.SetTriangleIntersectionTestType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector
type IMPSRayIntersector interface {
	IMPSKernel

	// Topic: Instance Properties

	BoundingBoxIntersectionTestType() MPSBoundingBoxIntersectionTestType
	SetBoundingBoxIntersectionTestType(value MPSBoundingBoxIntersectionTestType)
	CullMode() metal.MTLCullMode
	SetCullMode(value metal.MTLCullMode)
	FrontFacingWinding() metal.MTLWinding
	SetFrontFacingWinding(value metal.MTLWinding)
	IntersectionDataType() MPSIntersectionDataType
	SetIntersectionDataType(value MPSIntersectionDataType)
	IntersectionStride() uint
	SetIntersectionStride(value uint)
	RayDataType() MPSRayDataType
	SetRayDataType(value MPSRayDataType)
	RayIndexDataType() MPSDataType
	SetRayIndexDataType(value MPSDataType)
	RayMask() uint32
	SetRayMask(value uint32)
	RayMaskOperator() MPSRayMaskOperator
	SetRayMaskOperator(value MPSRayMaskOperator)
	RayMaskOptions() MPSRayMaskOptions
	SetRayMaskOptions(value MPSRayMaskOptions)
	RayStride() uint
	SetRayStride(value uint)
	TriangleIntersectionTestType() MPSTriangleIntersectionTestType
	SetTriangleIntersectionTestType(value MPSTriangleIntersectionTestType)
}

// Init initializes the instance.
func (r MPSRayIntersector) Init() MPSRayIntersector {
	rv := objc.Send[MPSRayIntersector](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRayIntersector) Autorelease() MPSRayIntersector {
	rv := objc.Send[MPSRayIntersector](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRayIntersector creates a new MPSRayIntersector instance.
func NewMPSRayIntersector() MPSRayIntersector {
	class := getMPSRayIntersectorClass()
	rv := objc.Send[MPSRayIntersector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewRayIntersectorWithCoder(aDecoder foundation.INSCoder) MPSRayIntersector {
	instance := getMPSRayIntersectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSRayIntersectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/init(coder:device:)
func NewRayIntersectorWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSRayIntersector {
	instance := getMPSRayIntersectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSRayIntersectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/init(device:)
func NewRayIntersectorWithDevice(device metal.MTLDevice) MPSRayIntersector {
	instance := getMPSRayIntersectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSRayIntersectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/boundingBoxIntersectionTestType
func (r MPSRayIntersector) BoundingBoxIntersectionTestType() MPSBoundingBoxIntersectionTestType {
	rv := objc.Send[MPSBoundingBoxIntersectionTestType](r.ID, objc.Sel("boundingBoxIntersectionTestType"))
	return MPSBoundingBoxIntersectionTestType(rv)
}
func (r MPSRayIntersector) SetBoundingBoxIntersectionTestType(value MPSBoundingBoxIntersectionTestType) {
	objc.Send[struct{}](r.ID, objc.Sel("setBoundingBoxIntersectionTestType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/cullMode
func (r MPSRayIntersector) CullMode() metal.MTLCullMode {
	rv := objc.Send[metal.MTLCullMode](r.ID, objc.Sel("cullMode"))
	return metal.MTLCullMode(rv)
}
func (r MPSRayIntersector) SetCullMode(value metal.MTLCullMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setCullMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/frontFacingWinding
func (r MPSRayIntersector) FrontFacingWinding() metal.MTLWinding {
	rv := objc.Send[metal.MTLWinding](r.ID, objc.Sel("frontFacingWinding"))
	return metal.MTLWinding(rv)
}
func (r MPSRayIntersector) SetFrontFacingWinding(value metal.MTLWinding) {
	objc.Send[struct{}](r.ID, objc.Sel("setFrontFacingWinding:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/intersectionDataType
func (r MPSRayIntersector) IntersectionDataType() MPSIntersectionDataType {
	rv := objc.Send[MPSIntersectionDataType](r.ID, objc.Sel("intersectionDataType"))
	return MPSIntersectionDataType(rv)
}
func (r MPSRayIntersector) SetIntersectionDataType(value MPSIntersectionDataType) {
	objc.Send[struct{}](r.ID, objc.Sel("setIntersectionDataType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/intersectionStride
func (r MPSRayIntersector) IntersectionStride() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("intersectionStride"))
	return rv
}
func (r MPSRayIntersector) SetIntersectionStride(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setIntersectionStride:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayDataType
func (r MPSRayIntersector) RayDataType() MPSRayDataType {
	rv := objc.Send[MPSRayDataType](r.ID, objc.Sel("rayDataType"))
	return MPSRayDataType(rv)
}
func (r MPSRayIntersector) SetRayDataType(value MPSRayDataType) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayDataType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayIndexDataType
func (r MPSRayIntersector) RayIndexDataType() MPSDataType {
	rv := objc.Send[MPSDataType](r.ID, objc.Sel("rayIndexDataType"))
	return MPSDataType(rv)
}
func (r MPSRayIntersector) SetRayIndexDataType(value MPSDataType) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayIndexDataType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayMask
func (r MPSRayIntersector) RayMask() uint32 {
	rv := objc.Send[uint32](r.ID, objc.Sel("rayMask"))
	return rv
}
func (r MPSRayIntersector) SetRayMask(value uint32) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayMask:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayMaskOperator
func (r MPSRayIntersector) RayMaskOperator() MPSRayMaskOperator {
	rv := objc.Send[MPSRayMaskOperator](r.ID, objc.Sel("rayMaskOperator"))
	return MPSRayMaskOperator(rv)
}
func (r MPSRayIntersector) SetRayMaskOperator(value MPSRayMaskOperator) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayMaskOperator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayMaskOptions
func (r MPSRayIntersector) RayMaskOptions() MPSRayMaskOptions {
	rv := objc.Send[MPSRayMaskOptions](r.ID, objc.Sel("rayMaskOptions"))
	return MPSRayMaskOptions(rv)
}
func (r MPSRayIntersector) SetRayMaskOptions(value MPSRayMaskOptions) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayMaskOptions:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/rayStride
func (r MPSRayIntersector) RayStride() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("rayStride"))
	return rv
}
func (r MPSRayIntersector) SetRayStride(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setRayStride:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/triangleIntersectionTestType
func (r MPSRayIntersector) TriangleIntersectionTestType() MPSTriangleIntersectionTestType {
	rv := objc.Send[MPSTriangleIntersectionTestType](r.ID, objc.Sel("triangleIntersectionTestType"))
	return MPSTriangleIntersectionTestType(rv)
}
func (r MPSRayIntersector) SetTriangleIntersectionTestType(value MPSTriangleIntersectionTestType) {
	objc.Send[struct{}](r.ID, objc.Sel("setTriangleIntersectionTestType:"), value)
}
