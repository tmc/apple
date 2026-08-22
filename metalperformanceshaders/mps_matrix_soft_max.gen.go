// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixSoftMax] class.
var (
	_MPSMatrixSoftMaxClass     MPSMatrixSoftMaxClass
	_MPSMatrixSoftMaxClassOnce sync.Once
)

func getMPSMatrixSoftMaxClass() MPSMatrixSoftMaxClass {
	_MPSMatrixSoftMaxClassOnce.Do(func() {
		_MPSMatrixSoftMaxClass = MPSMatrixSoftMaxClass{class: objc.GetClass("MPSMatrixSoftMax")}
	})
	return _MPSMatrixSoftMaxClass
}

// GetMPSMatrixSoftMaxClass returns the class object for MPSMatrixSoftMax.
func GetMPSMatrixSoftMaxClass() MPSMatrixSoftMaxClass {
	return getMPSMatrixSoftMaxClass()
}

type MPSMatrixSoftMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSoftMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSoftMaxClass) Alloc() MPSMatrixSoftMax {
	rv := objc.Send[MPSMatrixSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A softmax kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixSoftMax.SourceColumns]
//   - [MPSMatrixSoftMax.SetSourceColumns]
//   - [MPSMatrixSoftMax.SourceRows]
//   - [MPSMatrixSoftMax.SetSourceRows]
//
// # Instance Methods
//
//   - [MPSMatrixSoftMax.EncodeToCommandBufferInputMatrixResultMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax
type MPSMatrixSoftMax struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixSoftMaxFromID constructs a [MPSMatrixSoftMax] from an objc.ID.
//
// A softmax kernel that operates on matrices.
func MPSMatrixSoftMaxFromID(id objc.ID) MPSMatrixSoftMax {
	return MPSMatrixSoftMax{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixSoftMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSoftMax] class.
//
// # Instance Properties
//
//   - [IMPSMatrixSoftMax.SourceColumns]
//   - [IMPSMatrixSoftMax.SetSourceColumns]
//   - [IMPSMatrixSoftMax.SourceRows]
//   - [IMPSMatrixSoftMax.SetSourceRows]
//
// # Instance Methods
//
//   - [IMPSMatrixSoftMax.EncodeToCommandBufferInputMatrixResultMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax
type IMPSMatrixSoftMax interface {
	IMPSMatrixUnaryKernel

	// Topic: Instance Properties

	SourceColumns() uint
	SetSourceColumns(value uint)
	SourceRows() uint
	SetSourceRows(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, resultMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixSoftMax) Init() MPSMatrixSoftMax {
	rv := objc.Send[MPSMatrixSoftMax](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSoftMax) Autorelease() MPSMatrixSoftMax {
	rv := objc.Send[MPSMatrixSoftMax](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSoftMax creates a new MPSMatrixSoftMax instance.
func NewMPSMatrixSoftMax() MPSMatrixSoftMax {
	class := getMPSMatrixSoftMaxClass()
	rv := objc.Send[MPSMatrixSoftMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSoftMaxWithCoder(aDecoder foundation.INSCoder) MPSMatrixSoftMax {
	instance := getMPSMatrixSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/init(coder:device:)
func NewMatrixSoftMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSoftMax {
	instance := getMPSMatrixSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/init(device:)
func NewMatrixSoftMaxWithDevice(device metal.MTLDevice) MPSMatrixSoftMax {
	instance := getMPSMatrixSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/encode(commandBuffer:inputMatrix:resultMatrix:)
func (m MPSMatrixSoftMax) EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:resultMatrix:"), commandBuffer, inputMatrix, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/sourceColumns
func (m MPSMatrixSoftMax) SourceColumns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceColumns"))
	return rv
}
func (m MPSMatrixSoftMax) SetSourceColumns(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceColumns:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/sourceRows
func (m MPSMatrixSoftMax) SourceRows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceRows"))
	return rv
}
func (m MPSMatrixSoftMax) SetSourceRows(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceRows:"), value)
}
