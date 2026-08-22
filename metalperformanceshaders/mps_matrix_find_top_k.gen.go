// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixFindTopK] class.
var (
	_MPSMatrixFindTopKClass     MPSMatrixFindTopKClass
	_MPSMatrixFindTopKClassOnce sync.Once
)

func getMPSMatrixFindTopKClass() MPSMatrixFindTopKClass {
	_MPSMatrixFindTopKClassOnce.Do(func() {
		_MPSMatrixFindTopKClass = MPSMatrixFindTopKClass{class: objc.GetClass("MPSMatrixFindTopK")}
	})
	return _MPSMatrixFindTopKClass
}

// GetMPSMatrixFindTopKClass returns the class object for MPSMatrixFindTopK.
func GetMPSMatrixFindTopKClass() MPSMatrixFindTopKClass {
	return getMPSMatrixFindTopKClass()
}

type MPSMatrixFindTopKClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixFindTopKClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixFindTopKClass) Alloc() MPSMatrixFindTopK {
	rv := objc.Send[MPSMatrixFindTopK](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the top-K values and their corresponding indices in
// a matrix.
//
// # Initializers
//
//   - [MPSMatrixFindTopK.InitWithDeviceNumberOfTopKValues]
//
// # Instance Properties
//
//   - [MPSMatrixFindTopK.IndexOffset]
//   - [MPSMatrixFindTopK.SetIndexOffset]
//   - [MPSMatrixFindTopK.NumberOfTopKValues]
//   - [MPSMatrixFindTopK.SetNumberOfTopKValues]
//   - [MPSMatrixFindTopK.SourceColumns]
//   - [MPSMatrixFindTopK.SetSourceColumns]
//   - [MPSMatrixFindTopK.SourceRows]
//   - [MPSMatrixFindTopK.SetSourceRows]
//
// # Instance Methods
//
//   - [MPSMatrixFindTopK.EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK
type MPSMatrixFindTopK struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixFindTopKFromID constructs a [MPSMatrixFindTopK] from an objc.ID.
//
// A kernel for computing the top-K values and their corresponding indices in
// a matrix.
func MPSMatrixFindTopKFromID(id objc.ID) MPSMatrixFindTopK {
	return MPSMatrixFindTopK{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixFindTopK adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixFindTopK] class.
//
// # Initializers
//
//   - [IMPSMatrixFindTopK.InitWithDeviceNumberOfTopKValues]
//
// # Instance Properties
//
//   - [IMPSMatrixFindTopK.IndexOffset]
//   - [IMPSMatrixFindTopK.SetIndexOffset]
//   - [IMPSMatrixFindTopK.NumberOfTopKValues]
//   - [IMPSMatrixFindTopK.SetNumberOfTopKValues]
//   - [IMPSMatrixFindTopK.SourceColumns]
//   - [IMPSMatrixFindTopK.SetSourceColumns]
//   - [IMPSMatrixFindTopK.SourceRows]
//   - [IMPSMatrixFindTopK.SetSourceRows]
//
// # Instance Methods
//
//   - [IMPSMatrixFindTopK.EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK
type IMPSMatrixFindTopK interface {
	IMPSMatrixUnaryKernel

	// Topic: Initializers

	InitWithDeviceNumberOfTopKValues(device metal.MTLDevice, numberOfTopKValues uint) MPSMatrixFindTopK

	// Topic: Instance Properties

	IndexOffset() uint
	SetIndexOffset(value uint)
	NumberOfTopKValues() uint
	SetNumberOfTopKValues(value uint)
	SourceColumns() uint
	SetSourceColumns(value uint)
	SourceRows() uint
	SetSourceRows(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, resultIndexMatrix IMPSMatrix, resultValueMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixFindTopK) Init() MPSMatrixFindTopK {
	rv := objc.Send[MPSMatrixFindTopK](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixFindTopK) Autorelease() MPSMatrixFindTopK {
	rv := objc.Send[MPSMatrixFindTopK](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixFindTopK creates a new MPSMatrixFindTopK instance.
func NewMPSMatrixFindTopK() MPSMatrixFindTopK {
	class := getMPSMatrixFindTopKClass()
	rv := objc.Send[MPSMatrixFindTopK](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixFindTopKWithCoder(aDecoder foundation.INSCoder) MPSMatrixFindTopK {
	instance := getMPSMatrixFindTopKClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixFindTopKFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/init(coder:device:)
func NewMatrixFindTopKWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixFindTopK {
	instance := getMPSMatrixFindTopKClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixFindTopKFromID(rv)
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
func NewMatrixFindTopKWithDevice(device metal.MTLDevice) MPSMatrixFindTopK {
	instance := getMPSMatrixFindTopKClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixFindTopKFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/init(device:numberOfTopKValues:)
func NewMatrixFindTopKWithDeviceNumberOfTopKValues(device metal.MTLDevice, numberOfTopKValues uint) MPSMatrixFindTopK {
	instance := getMPSMatrixFindTopKClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:numberOfTopKValues:"), device, numberOfTopKValues)
	return MPSMatrixFindTopKFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/init(device:numberOfTopKValues:)
func (m MPSMatrixFindTopK) InitWithDeviceNumberOfTopKValues(device metal.MTLDevice, numberOfTopKValues uint) MPSMatrixFindTopK {
	rv := objc.Send[MPSMatrixFindTopK](m.ID, objc.Sel("initWithDevice:numberOfTopKValues:"), device, numberOfTopKValues)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/encode(commandBuffer:inputMatrix:resultIndexMatrix:resultValueMatrix:)
func (m MPSMatrixFindTopK) EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, resultIndexMatrix IMPSMatrix, resultValueMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:resultIndexMatrix:resultValueMatrix:"), commandBuffer, inputMatrix, resultIndexMatrix, resultValueMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/indexOffset
func (m MPSMatrixFindTopK) IndexOffset() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("indexOffset"))
	return rv
}
func (m MPSMatrixFindTopK) SetIndexOffset(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/numberOfTopKValues
func (m MPSMatrixFindTopK) NumberOfTopKValues() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("numberOfTopKValues"))
	return rv
}
func (m MPSMatrixFindTopK) SetNumberOfTopKValues(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberOfTopKValues:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/sourceColumns
func (m MPSMatrixFindTopK) SourceColumns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceColumns"))
	return rv
}
func (m MPSMatrixFindTopK) SetSourceColumns(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceColumns:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK/sourceRows
func (m MPSMatrixFindTopK) SourceRows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceRows"))
	return rv
}
func (m MPSMatrixFindTopK) SetSourceRows(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceRows:"), value)
}
