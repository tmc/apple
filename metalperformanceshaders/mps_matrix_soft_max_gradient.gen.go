// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixSoftMaxGradient] class.
var (
	_MPSMatrixSoftMaxGradientClass     MPSMatrixSoftMaxGradientClass
	_MPSMatrixSoftMaxGradientClassOnce sync.Once
)

func getMPSMatrixSoftMaxGradientClass() MPSMatrixSoftMaxGradientClass {
	_MPSMatrixSoftMaxGradientClassOnce.Do(func() {
		_MPSMatrixSoftMaxGradientClass = MPSMatrixSoftMaxGradientClass{class: objc.GetClass("MPSMatrixSoftMaxGradient")}
	})
	return _MPSMatrixSoftMaxGradientClass
}

// GetMPSMatrixSoftMaxGradientClass returns the class object for MPSMatrixSoftMaxGradient.
func GetMPSMatrixSoftMaxGradientClass() MPSMatrixSoftMaxGradientClass {
	return getMPSMatrixSoftMaxGradientClass()
}

type MPSMatrixSoftMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSoftMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSoftMaxGradientClass) Alloc() MPSMatrixSoftMaxGradient {
	rv := objc.Send[MPSMatrixSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient softmax kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixSoftMaxGradient.SourceColumns]
//   - [MPSMatrixSoftMaxGradient.SetSourceColumns]
//   - [MPSMatrixSoftMaxGradient.SourceRows]
//   - [MPSMatrixSoftMaxGradient.SetSourceRows]
//
// # Instance Methods
//
//   - [MPSMatrixSoftMaxGradient.EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient
type MPSMatrixSoftMaxGradient struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixSoftMaxGradientFromID constructs a [MPSMatrixSoftMaxGradient] from an objc.ID.
//
// A gradient softmax kernel that operates on matrices.
func MPSMatrixSoftMaxGradientFromID(id objc.ID) MPSMatrixSoftMaxGradient {
	return MPSMatrixSoftMaxGradient{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixSoftMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSoftMaxGradient] class.
//
// # Instance Properties
//
//   - [IMPSMatrixSoftMaxGradient.SourceColumns]
//   - [IMPSMatrixSoftMaxGradient.SetSourceColumns]
//   - [IMPSMatrixSoftMaxGradient.SourceRows]
//   - [IMPSMatrixSoftMaxGradient.SetSourceRows]
//
// # Instance Methods
//
//   - [IMPSMatrixSoftMaxGradient.EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient
type IMPSMatrixSoftMaxGradient interface {
	IMPSMatrixBinaryKernel

	// Topic: Instance Properties

	SourceColumns() uint
	SetSourceColumns(value uint)
	SourceRows() uint
	SetSourceRows(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, forwardOutputMatrix IMPSMatrix, resultMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixSoftMaxGradient) Init() MPSMatrixSoftMaxGradient {
	rv := objc.Send[MPSMatrixSoftMaxGradient](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSoftMaxGradient) Autorelease() MPSMatrixSoftMaxGradient {
	rv := objc.Send[MPSMatrixSoftMaxGradient](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSoftMaxGradient creates a new MPSMatrixSoftMaxGradient instance.
func NewMPSMatrixSoftMaxGradient() MPSMatrixSoftMaxGradient {
	class := getMPSMatrixSoftMaxGradientClass()
	rv := objc.Send[MPSMatrixSoftMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSoftMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSMatrixSoftMaxGradient {
	instance := getMPSMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/init(coder:device:)
func NewMatrixSoftMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSoftMaxGradient {
	instance := getMPSMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/init(device:)
func NewMatrixSoftMaxGradientWithDevice(device metal.MTLDevice) MPSMatrixSoftMaxGradient {
	instance := getMPSMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/encode(to:gradientMatrix:forwardOutputMatrix:resultMatrix:)
func (m MPSMatrixSoftMaxGradient) EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, forwardOutputMatrix IMPSMatrix, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:forwardOutputMatrix:resultMatrix:"), commandBuffer, gradientMatrix, forwardOutputMatrix, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/sourceColumns
func (m MPSMatrixSoftMaxGradient) SourceColumns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceColumns"))
	return rv
}
func (m MPSMatrixSoftMaxGradient) SetSourceColumns(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceColumns:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/sourceRows
func (m MPSMatrixSoftMaxGradient) SourceRows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceRows"))
	return rv
}
func (m MPSMatrixSoftMaxGradient) SetSourceRows(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceRows:"), value)
}
