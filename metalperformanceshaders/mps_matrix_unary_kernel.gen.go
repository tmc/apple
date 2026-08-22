// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixUnaryKernel] class.
var (
	_MPSMatrixUnaryKernelClass     MPSMatrixUnaryKernelClass
	_MPSMatrixUnaryKernelClassOnce sync.Once
)

func getMPSMatrixUnaryKernelClass() MPSMatrixUnaryKernelClass {
	_MPSMatrixUnaryKernelClassOnce.Do(func() {
		_MPSMatrixUnaryKernelClass = MPSMatrixUnaryKernelClass{class: objc.GetClass("MPSMatrixUnaryKernel")}
	})
	return _MPSMatrixUnaryKernelClass
}

// GetMPSMatrixUnaryKernelClass returns the class object for MPSMatrixUnaryKernel.
func GetMPSMatrixUnaryKernelClass() MPSMatrixUnaryKernelClass {
	return getMPSMatrixUnaryKernelClass()
}

type MPSMatrixUnaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixUnaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixUnaryKernelClass) Alloc() MPSMatrixUnaryKernel {
	rv := objc.Send[MPSMatrixUnaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that consumes one matrix and produces one matrix.
//
// # Instance Properties
//
//   - [MPSMatrixUnaryKernel.BatchSize]
//   - [MPSMatrixUnaryKernel.SetBatchSize]
//   - [MPSMatrixUnaryKernel.BatchStart]
//   - [MPSMatrixUnaryKernel.SetBatchStart]
//   - [MPSMatrixUnaryKernel.ResultMatrixOrigin]
//   - [MPSMatrixUnaryKernel.SetResultMatrixOrigin]
//   - [MPSMatrixUnaryKernel.SourceMatrixOrigin]
//   - [MPSMatrixUnaryKernel.SetSourceMatrixOrigin]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel
type MPSMatrixUnaryKernel struct {
	MPSKernel
}

// MPSMatrixUnaryKernelFromID constructs a [MPSMatrixUnaryKernel] from an objc.ID.
//
// A kernel that consumes one matrix and produces one matrix.
func MPSMatrixUnaryKernelFromID(id objc.ID) MPSMatrixUnaryKernel {
	return MPSMatrixUnaryKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixUnaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixUnaryKernel] class.
//
// # Instance Properties
//
//   - [IMPSMatrixUnaryKernel.BatchSize]
//   - [IMPSMatrixUnaryKernel.SetBatchSize]
//   - [IMPSMatrixUnaryKernel.BatchStart]
//   - [IMPSMatrixUnaryKernel.SetBatchStart]
//   - [IMPSMatrixUnaryKernel.ResultMatrixOrigin]
//   - [IMPSMatrixUnaryKernel.SetResultMatrixOrigin]
//   - [IMPSMatrixUnaryKernel.SourceMatrixOrigin]
//   - [IMPSMatrixUnaryKernel.SetSourceMatrixOrigin]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel
type IMPSMatrixUnaryKernel interface {
	IMPSKernel

	// Topic: Instance Properties

	BatchSize() uint
	SetBatchSize(value uint)
	BatchStart() uint
	SetBatchStart(value uint)
	ResultMatrixOrigin() metal.MTLOrigin
	SetResultMatrixOrigin(value metal.MTLOrigin)
	SourceMatrixOrigin() metal.MTLOrigin
	SetSourceMatrixOrigin(value metal.MTLOrigin)
}

// Init initializes the instance.
func (m MPSMatrixUnaryKernel) Init() MPSMatrixUnaryKernel {
	rv := objc.Send[MPSMatrixUnaryKernel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixUnaryKernel) Autorelease() MPSMatrixUnaryKernel {
	rv := objc.Send[MPSMatrixUnaryKernel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixUnaryKernel creates a new MPSMatrixUnaryKernel instance.
func NewMPSMatrixUnaryKernel() MPSMatrixUnaryKernel {
	class := getMPSMatrixUnaryKernelClass()
	rv := objc.Send[MPSMatrixUnaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixUnaryKernelWithCoder(aDecoder foundation.INSCoder) MPSMatrixUnaryKernel {
	instance := getMPSMatrixUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixUnaryKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixUnaryKernel {
	instance := getMPSMatrixUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixUnaryKernelFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewMatrixUnaryKernelWithDevice(device metal.MTLDevice) MPSMatrixUnaryKernel {
	instance := getMPSMatrixUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel/batchSize
func (m MPSMatrixUnaryKernel) BatchSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MPSMatrixUnaryKernel) SetBatchSize(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel/batchStart
func (m MPSMatrixUnaryKernel) BatchStart() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchStart"))
	return rv
}
func (m MPSMatrixUnaryKernel) SetBatchStart(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchStart:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel/resultMatrixOrigin
func (m MPSMatrixUnaryKernel) ResultMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("resultMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixUnaryKernel) SetResultMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setResultMatrixOrigin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel/sourceMatrixOrigin
func (m MPSMatrixUnaryKernel) SourceMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("sourceMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixUnaryKernel) SetSourceMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceMatrixOrigin:"), value)
}
