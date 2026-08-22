// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBinaryFullyConnected] class.
var (
	_MPSCNNBinaryFullyConnectedClass     MPSCNNBinaryFullyConnectedClass
	_MPSCNNBinaryFullyConnectedClassOnce sync.Once
)

func getMPSCNNBinaryFullyConnectedClass() MPSCNNBinaryFullyConnectedClass {
	_MPSCNNBinaryFullyConnectedClassOnce.Do(func() {
		_MPSCNNBinaryFullyConnectedClass = MPSCNNBinaryFullyConnectedClass{class: objc.GetClass("MPSCNNBinaryFullyConnected")}
	})
	return _MPSCNNBinaryFullyConnectedClass
}

// GetMPSCNNBinaryFullyConnectedClass returns the class object for MPSCNNBinaryFullyConnected.
func GetMPSCNNBinaryFullyConnectedClass() MPSCNNBinaryFullyConnectedClass {
	return getMPSCNNBinaryFullyConnectedClass()
}

type MPSCNNBinaryFullyConnectedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBinaryFullyConnectedClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryFullyConnectedClass) Alloc() MPSCNNBinaryFullyConnected {
	rv := objc.Send[MPSCNNBinaryFullyConnected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A fully connected convolution layer with binary weights and optionally
// binarized input image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected
type MPSCNNBinaryFullyConnected struct {
	MPSCNNBinaryConvolution
}

// MPSCNNBinaryFullyConnectedFromID constructs a [MPSCNNBinaryFullyConnected] from an objc.ID.
//
// A fully connected convolution layer with binary weights and optionally
// binarized input image.
func MPSCNNBinaryFullyConnectedFromID(id objc.ID) MPSCNNBinaryFullyConnected {
	return MPSCNNBinaryFullyConnected{MPSCNNBinaryConvolution: MPSCNNBinaryConvolutionFromID(id)}
}

// NOTE: MPSCNNBinaryFullyConnected adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBinaryFullyConnected] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected
type IMPSCNNBinaryFullyConnected interface {
	IMPSCNNBinaryConvolution
}

// Init initializes the instance.
func (c MPSCNNBinaryFullyConnected) Init() MPSCNNBinaryFullyConnected {
	rv := objc.Send[MPSCNNBinaryFullyConnected](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryFullyConnected) Autorelease() MPSCNNBinaryFullyConnected {
	rv := objc.Send[MPSCNNBinaryFullyConnected](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryFullyConnected creates a new MPSCNNBinaryFullyConnected instance.
func NewMPSCNNBinaryFullyConnected() MPSCNNBinaryFullyConnected {
	class := getMPSCNNBinaryFullyConnectedClass()
	rv := objc.Send[MPSCNNBinaryFullyConnected](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBinaryFullyConnectedWithCoder(aDecoder foundation.INSCoder) MPSCNNBinaryFullyConnected {
	instance := getMPSCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBinaryFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer with binary weights.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected/init(coder:device:)
func NewCNNBinaryFullyConnectedWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBinaryFullyConnected {
	instance := getMPSCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBinaryFullyConnectedFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNBinaryFullyConnectedWithDevice(device metal.MTLDevice) MPSCNNBinaryFullyConnected {
	instance := getMPSCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBinaryFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer with binary weights.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected/init(device:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func NewCNNBinaryFullyConnectedWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryFullyConnected {
	instance := getMPSCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), device, convolutionData, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return MPSCNNBinaryFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer with binary weights.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected/init(device:convolutionData:scaleValue:type:flags:)
func NewCNNBinaryFullyConnectedWithDeviceConvolutionDataScaleValueTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryFullyConnected {
	instance := getMPSCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), device, convolutionData, scaleValue, type_, flags)
	return MPSCNNBinaryFullyConnectedFromID(rv)
}
