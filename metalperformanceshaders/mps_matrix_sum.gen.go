// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSMatrixSum] class.
var (
	_MPSMatrixSumClass     MPSMatrixSumClass
	_MPSMatrixSumClassOnce sync.Once
)

func getMPSMatrixSumClass() MPSMatrixSumClass {
	_MPSMatrixSumClassOnce.Do(func() {
		_MPSMatrixSumClass = MPSMatrixSumClass{class: objc.GetClass("MPSMatrixSum")}
	})
	return _MPSMatrixSumClass
}

// GetMPSMatrixSumClass returns the class object for MPSMatrixSum.
func GetMPSMatrixSumClass() MPSMatrixSumClass {
	return getMPSMatrixSumClass()
}

type MPSMatrixSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSumClass) Alloc() MPSMatrixSum {
	rv := objc.Send[MPSMatrixSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for performing a pointwise summation of a matrix.
//
// # Initializers
//
//   - [MPSMatrixSum.InitWithDeviceCountRowsColumnsTranspose]
//
// # Instance Properties
//
//   - [MPSMatrixSum.Columns]
//   - [MPSMatrixSum.Count]
//   - [MPSMatrixSum.NeuronParameterA]
//   - [MPSMatrixSum.NeuronParameterB]
//   - [MPSMatrixSum.NeuronParameterC]
//   - [MPSMatrixSum.ResultMatrixOrigin]
//   - [MPSMatrixSum.SetResultMatrixOrigin]
//   - [MPSMatrixSum.Rows]
//   - [MPSMatrixSum.Transpose]
//
// # Instance Methods
//
//   - [MPSMatrixSum.EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex]
//   - [MPSMatrixSum.NeuronType]
//   - [MPSMatrixSum.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum
type MPSMatrixSum struct {
	MPSKernel
}

// MPSMatrixSumFromID constructs a [MPSMatrixSum] from an objc.ID.
//
// A kernel for performing a pointwise summation of a matrix.
func MPSMatrixSumFromID(id objc.ID) MPSMatrixSum {
	return MPSMatrixSum{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSum] class.
//
// # Initializers
//
//   - [IMPSMatrixSum.InitWithDeviceCountRowsColumnsTranspose]
//
// # Instance Properties
//
//   - [IMPSMatrixSum.Columns]
//   - [IMPSMatrixSum.Count]
//   - [IMPSMatrixSum.NeuronParameterA]
//   - [IMPSMatrixSum.NeuronParameterB]
//   - [IMPSMatrixSum.NeuronParameterC]
//   - [IMPSMatrixSum.ResultMatrixOrigin]
//   - [IMPSMatrixSum.SetResultMatrixOrigin]
//   - [IMPSMatrixSum.Rows]
//   - [IMPSMatrixSum.Transpose]
//
// # Instance Methods
//
//   - [IMPSMatrixSum.EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex]
//   - [IMPSMatrixSum.NeuronType]
//   - [IMPSMatrixSum.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum
type IMPSMatrixSum interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceCountRowsColumnsTranspose(device metal.MTLDevice, count uint, rows uint, columns uint, transpose bool) MPSMatrixSum

	// Topic: Instance Properties

	Columns() uint
	Count() uint
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	ResultMatrixOrigin() metal.MTLOrigin
	SetResultMatrixOrigin(value metal.MTLOrigin)
	Rows() uint
	Transpose() bool

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, resultMatrix IMPSMatrix, scaleVector IMPSVector, offsetVector IMPSVector, biasVector IMPSVector, startIndex uint)
	NeuronType() MPSCNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixSum) Init() MPSMatrixSum {
	rv := objc.Send[MPSMatrixSum](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSum) Autorelease() MPSMatrixSum {
	rv := objc.Send[MPSMatrixSum](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSum creates a new MPSMatrixSum instance.
func NewMPSMatrixSum() MPSMatrixSum {
	class := getMPSMatrixSumClass()
	rv := objc.Send[MPSMatrixSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSumWithCoder(aDecoder foundation.INSCoder) MPSMatrixSum {
	instance := getMPSMatrixSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/init(coder:device:)
func NewMatrixSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSum {
	instance := getMPSMatrixSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSumFromID(rv)
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
func NewMatrixSumWithDevice(device metal.MTLDevice) MPSMatrixSum {
	instance := getMPSMatrixSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/init(device:count:rows:columns:transpose:)
func NewMatrixSumWithDeviceCountRowsColumnsTranspose(device metal.MTLDevice, count uint, rows uint, columns uint, transpose bool) MPSMatrixSum {
	instance := getMPSMatrixSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:count:rows:columns:transpose:"), device, count, rows, columns, transpose)
	return MPSMatrixSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/init(device:count:rows:columns:transpose:)
func (m MPSMatrixSum) InitWithDeviceCountRowsColumnsTranspose(device metal.MTLDevice, count uint, rows uint, columns uint, transpose bool) MPSMatrixSum {
	rv := objc.Send[MPSMatrixSum](m.ID, objc.Sel("initWithDevice:count:rows:columns:transpose:"), device, count, rows, columns, transpose)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/encode(to:sourceMatrices:resultMatrix:scale:offsetVector:biasVector:start:)
func (m MPSMatrixSum) EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, resultMatrix IMPSMatrix, scaleVector IMPSVector, offsetVector IMPSVector, biasVector IMPSVector, startIndex uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrices:resultMatrix:scaleVector:offsetVector:biasVector:startIndex:"), buffer, objectivec.IObjectSliceToNSArray(sourceMatrices), resultMatrix, scaleVector, offsetVector, biasVector, startIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/neuronType()
func (m MPSMatrixSum) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixSum) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/columns
func (m MPSMatrixSum) Columns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("columns"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/count
func (m MPSMatrixSum) Count() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("count"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/neuronParameterA
func (m MPSMatrixSum) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/neuronParameterB
func (m MPSMatrixSum) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/neuronParameterC
func (m MPSMatrixSum) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/resultMatrixOrigin
func (m MPSMatrixSum) ResultMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("resultMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixSum) SetResultMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setResultMatrixOrigin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/rows
func (m MPSMatrixSum) Rows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rows"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/transpose
func (m MPSMatrixSum) Transpose() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("transpose"))
	return rv
}
