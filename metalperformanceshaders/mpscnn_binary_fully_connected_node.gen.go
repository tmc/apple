// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBinaryFullyConnectedNode] class.
var (
	_MPSCNNBinaryFullyConnectedNodeClass     MPSCNNBinaryFullyConnectedNodeClass
	_MPSCNNBinaryFullyConnectedNodeClassOnce sync.Once
)

func getMPSCNNBinaryFullyConnectedNodeClass() MPSCNNBinaryFullyConnectedNodeClass {
	_MPSCNNBinaryFullyConnectedNodeClassOnce.Do(func() {
		_MPSCNNBinaryFullyConnectedNodeClass = MPSCNNBinaryFullyConnectedNodeClass{class: objc.GetClass("MPSCNNBinaryFullyConnectedNode")}
	})
	return _MPSCNNBinaryFullyConnectedNodeClass
}

// GetMPSCNNBinaryFullyConnectedNodeClass returns the class object for MPSCNNBinaryFullyConnectedNode.
func GetMPSCNNBinaryFullyConnectedNodeClass() MPSCNNBinaryFullyConnectedNodeClass {
	return getMPSCNNBinaryFullyConnectedNodeClass()
}

type MPSCNNBinaryFullyConnectedNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBinaryFullyConnectedNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryFullyConnectedNodeClass) Alloc() MPSCNNBinaryFullyConnectedNode {
	rv := objc.Send[MPSCNNBinaryFullyConnectedNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a fully connected convolution layer with binary weights
// and optionally binarized input image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnectedNode
type MPSCNNBinaryFullyConnectedNode struct {
	MPSCNNBinaryConvolutionNode
}

// MPSCNNBinaryFullyConnectedNodeFromID constructs a [MPSCNNBinaryFullyConnectedNode] from an objc.ID.
//
// A representation of a fully connected convolution layer with binary weights
// and optionally binarized input image.
func MPSCNNBinaryFullyConnectedNodeFromID(id objc.ID) MPSCNNBinaryFullyConnectedNode {
	return MPSCNNBinaryFullyConnectedNode{MPSCNNBinaryConvolutionNode: MPSCNNBinaryConvolutionNodeFromID(id)}
}

// NOTE: MPSCNNBinaryFullyConnectedNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBinaryFullyConnectedNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnectedNode
type IMPSCNNBinaryFullyConnectedNode interface {
	IMPSCNNBinaryConvolutionNode
}

// Init initializes the instance.
func (c MPSCNNBinaryFullyConnectedNode) Init() MPSCNNBinaryFullyConnectedNode {
	rv := objc.Send[MPSCNNBinaryFullyConnectedNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryFullyConnectedNode) Autorelease() MPSCNNBinaryFullyConnectedNode {
	rv := objc.Send[MPSCNNBinaryFullyConnectedNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryFullyConnectedNode creates a new MPSCNNBinaryFullyConnectedNode instance.
func NewMPSCNNBinaryFullyConnectedNode() MPSCNNBinaryFullyConnectedNode {
	class := getMPSCNNBinaryFullyConnectedNodeClass()
	rv := objc.Send[MPSCNNBinaryFullyConnectedNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/init(source:weights:)
func NewCNNBinaryFullyConnectedNodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNBinaryFullyConnectedNode {
	instance := getMPSCNNBinaryFullyConnectedNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return MPSCNNBinaryFullyConnectedNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnectedNode/init(source:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryFullyConnectedNode {
	instance := getMPSCNNBinaryFullyConnectedNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return MPSCNNBinaryFullyConnectedNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnectedNode/init(source:weights:scaleValue:type:flags:)
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsScaleValueTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryFullyConnectedNode {
	instance := getMPSCNNBinaryFullyConnectedNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return MPSCNNBinaryFullyConnectedNodeFromID(rv)
}
