// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixBinaryKernel] class.
var (
	_MPSMatrixBinaryKernelClass     MPSMatrixBinaryKernelClass
	_MPSMatrixBinaryKernelClassOnce sync.Once
)

func getMPSMatrixBinaryKernelClass() MPSMatrixBinaryKernelClass {
	_MPSMatrixBinaryKernelClassOnce.Do(func() {
		_MPSMatrixBinaryKernelClass = MPSMatrixBinaryKernelClass{class: objc.GetClass("MPSMatrixBinaryKernel")}
	})
	return _MPSMatrixBinaryKernelClass
}

// GetMPSMatrixBinaryKernelClass returns the class object for MPSMatrixBinaryKernel.
func GetMPSMatrixBinaryKernelClass() MPSMatrixBinaryKernelClass {
	return getMPSMatrixBinaryKernelClass()
}

type MPSMatrixBinaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixBinaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixBinaryKernelClass) Alloc() MPSMatrixBinaryKernel {
	rv := objc.Send[MPSMatrixBinaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that consumes two matrices and produces one matrix.
//
// # Instance Properties
//
//   - [MPSMatrixBinaryKernel.BatchSize]
//   - [MPSMatrixBinaryKernel.SetBatchSize]
//   - [MPSMatrixBinaryKernel.BatchStart]
//   - [MPSMatrixBinaryKernel.SetBatchStart]
//   - [MPSMatrixBinaryKernel.PrimarySourceMatrixOrigin]
//   - [MPSMatrixBinaryKernel.SetPrimarySourceMatrixOrigin]
//   - [MPSMatrixBinaryKernel.ResultMatrixOrigin]
//   - [MPSMatrixBinaryKernel.SetResultMatrixOrigin]
//   - [MPSMatrixBinaryKernel.SecondarySourceMatrixOrigin]
//   - [MPSMatrixBinaryKernel.SetSecondarySourceMatrixOrigin]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel
type MPSMatrixBinaryKernel struct {
	MPSKernel
}

// MPSMatrixBinaryKernelFromID constructs a [MPSMatrixBinaryKernel] from an objc.ID.
//
// A kernel that consumes two matrices and produces one matrix.
func MPSMatrixBinaryKernelFromID(id objc.ID) MPSMatrixBinaryKernel {
	return MPSMatrixBinaryKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixBinaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixBinaryKernel] class.
//
// # Instance Properties
//
//   - [IMPSMatrixBinaryKernel.BatchSize]
//   - [IMPSMatrixBinaryKernel.SetBatchSize]
//   - [IMPSMatrixBinaryKernel.BatchStart]
//   - [IMPSMatrixBinaryKernel.SetBatchStart]
//   - [IMPSMatrixBinaryKernel.PrimarySourceMatrixOrigin]
//   - [IMPSMatrixBinaryKernel.SetPrimarySourceMatrixOrigin]
//   - [IMPSMatrixBinaryKernel.ResultMatrixOrigin]
//   - [IMPSMatrixBinaryKernel.SetResultMatrixOrigin]
//   - [IMPSMatrixBinaryKernel.SecondarySourceMatrixOrigin]
//   - [IMPSMatrixBinaryKernel.SetSecondarySourceMatrixOrigin]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel
type IMPSMatrixBinaryKernel interface {
	IMPSKernel

	// Topic: Instance Properties

	BatchSize() uint
	SetBatchSize(value uint)
	BatchStart() uint
	SetBatchStart(value uint)
	PrimarySourceMatrixOrigin() metal.MTLOrigin
	SetPrimarySourceMatrixOrigin(value metal.MTLOrigin)
	ResultMatrixOrigin() metal.MTLOrigin
	SetResultMatrixOrigin(value metal.MTLOrigin)
	SecondarySourceMatrixOrigin() metal.MTLOrigin
	SetSecondarySourceMatrixOrigin(value metal.MTLOrigin)
}

// Init initializes the instance.
func (m MPSMatrixBinaryKernel) Init() MPSMatrixBinaryKernel {
	rv := objc.Send[MPSMatrixBinaryKernel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixBinaryKernel) Autorelease() MPSMatrixBinaryKernel {
	rv := objc.Send[MPSMatrixBinaryKernel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixBinaryKernel creates a new MPSMatrixBinaryKernel instance.
func NewMPSMatrixBinaryKernel() MPSMatrixBinaryKernel {
	class := getMPSMatrixBinaryKernelClass()
	rv := objc.Send[MPSMatrixBinaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixBinaryKernelWithCoder(aDecoder foundation.INSCoder) MPSMatrixBinaryKernel {
	instance := getMPSMatrixBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixBinaryKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixBinaryKernel {
	instance := getMPSMatrixBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixBinaryKernelFromID(rv)
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
func NewMatrixBinaryKernelWithDevice(device metal.MTLDevice) MPSMatrixBinaryKernel {
	instance := getMPSMatrixBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel/batchSize
func (m MPSMatrixBinaryKernel) BatchSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MPSMatrixBinaryKernel) SetBatchSize(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel/batchStart
func (m MPSMatrixBinaryKernel) BatchStart() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchStart"))
	return rv
}
func (m MPSMatrixBinaryKernel) SetBatchStart(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchStart:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel/primarySourceMatrixOrigin
func (m MPSMatrixBinaryKernel) PrimarySourceMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("primarySourceMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixBinaryKernel) SetPrimarySourceMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrimarySourceMatrixOrigin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel/resultMatrixOrigin
func (m MPSMatrixBinaryKernel) ResultMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("resultMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixBinaryKernel) SetResultMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setResultMatrixOrigin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel/secondarySourceMatrixOrigin
func (m MPSMatrixBinaryKernel) SecondarySourceMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("secondarySourceMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixBinaryKernel) SetSecondarySourceMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setSecondarySourceMatrixOrigin:"), value)
}
