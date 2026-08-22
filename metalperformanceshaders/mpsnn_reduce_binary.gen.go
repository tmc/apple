// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceBinary] class.
var (
	_MPSNNReduceBinaryClass     MPSNNReduceBinaryClass
	_MPSNNReduceBinaryClassOnce sync.Once
)

func getMPSNNReduceBinaryClass() MPSNNReduceBinaryClass {
	_MPSNNReduceBinaryClassOnce.Do(func() {
		_MPSNNReduceBinaryClass = MPSNNReduceBinaryClass{class: objc.GetClass("MPSNNReduceBinary")}
	})
	return _MPSNNReduceBinaryClass
}

// GetMPSNNReduceBinaryClass returns the class object for MPSNNReduceBinary.
func GetMPSNNReduceBinaryClass() MPSNNReduceBinaryClass {
	return getMPSNNReduceBinaryClass()
}

type MPSNNReduceBinaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceBinaryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceBinaryClass) Alloc() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for binary reduction filters.
//
// # Instance Properties
//
//   - [MPSNNReduceBinary.PrimarySourceClipRect]
//   - [MPSNNReduceBinary.SetPrimarySourceClipRect]
//   - [MPSNNReduceBinary.SecondarySourceClipRect]
//   - [MPSNNReduceBinary.SetSecondarySourceClipRect]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceBinary
type MPSNNReduceBinary struct {
	MPSCNNBinaryKernel
}

// MPSNNReduceBinaryFromID constructs a [MPSNNReduceBinary] from an objc.ID.
//
// The base class for binary reduction filters.
func MPSNNReduceBinaryFromID(id objc.ID) MPSNNReduceBinary {
	return MPSNNReduceBinary{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}

// NOTE: MPSNNReduceBinary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceBinary] class.
//
// # Instance Properties
//
//   - [IMPSNNReduceBinary.PrimarySourceClipRect]
//   - [IMPSNNReduceBinary.SetPrimarySourceClipRect]
//   - [IMPSNNReduceBinary.SecondarySourceClipRect]
//   - [IMPSNNReduceBinary.SetSecondarySourceClipRect]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceBinary
type IMPSNNReduceBinary interface {
	IMPSCNNBinaryKernel

	// Topic: Instance Properties

	PrimarySourceClipRect() metal.MTLRegion
	SetPrimarySourceClipRect(value metal.MTLRegion)
	SecondarySourceClipRect() metal.MTLRegion
	SetSecondarySourceClipRect(value metal.MTLRegion)
}

// Init initializes the instance.
func (r MPSNNReduceBinary) Init() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceBinary) Autorelease() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceBinary creates a new MPSNNReduceBinary instance.
func NewMPSNNReduceBinary() MPSNNReduceBinary {
	class := getMPSNNReduceBinaryClass()
	rv := objc.Send[MPSNNReduceBinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceBinaryWithCoder(aDecoder foundation.INSCoder) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceBinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewReduceBinaryWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceBinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(device:)
func NewReduceBinaryWithDevice(device metal.MTLDevice) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceBinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceBinary/primarySourceClipRect
func (r MPSNNReduceBinary) PrimarySourceClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](r.ID, objc.Sel("primarySourceClipRect"))
	return metal.MTLRegion(rv)
}
func (r MPSNNReduceBinary) SetPrimarySourceClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](r.ID, objc.Sel("setPrimarySourceClipRect:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceBinary/secondarySourceClipRect
func (r MPSNNReduceBinary) SecondarySourceClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](r.ID, objc.Sel("secondarySourceClipRect"))
	return metal.MTLRegion(rv)
}
func (r MPSNNReduceBinary) SetSecondarySourceClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](r.ID, objc.Sel("setSecondarySourceClipRect:"), value)
}
