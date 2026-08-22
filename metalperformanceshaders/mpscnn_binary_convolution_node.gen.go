// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBinaryConvolutionNode] class.
var (
	_MPSCNNBinaryConvolutionNodeClass     MPSCNNBinaryConvolutionNodeClass
	_MPSCNNBinaryConvolutionNodeClassOnce sync.Once
)

func getMPSCNNBinaryConvolutionNodeClass() MPSCNNBinaryConvolutionNodeClass {
	_MPSCNNBinaryConvolutionNodeClassOnce.Do(func() {
		_MPSCNNBinaryConvolutionNodeClass = MPSCNNBinaryConvolutionNodeClass{class: objc.GetClass("MPSCNNBinaryConvolutionNode")}
	})
	return _MPSCNNBinaryConvolutionNodeClass
}

// GetMPSCNNBinaryConvolutionNodeClass returns the class object for MPSCNNBinaryConvolutionNode.
func GetMPSCNNBinaryConvolutionNodeClass() MPSCNNBinaryConvolutionNodeClass {
	return getMPSCNNBinaryConvolutionNodeClass()
}

type MPSCNNBinaryConvolutionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBinaryConvolutionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryConvolutionNodeClass) Alloc() MPSCNNBinaryConvolutionNode {
	rv := objc.Send[MPSCNNBinaryConvolutionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a convolution kernel with binary weights and an input
// image using binary approximations.
//
// # Initializers
//
//   - [MPSCNNBinaryConvolutionNode.InitWithSourceWeightsScaleValueTypeFlags]
//   - [MPSCNNBinaryConvolutionNode.InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode
type MPSCNNBinaryConvolutionNode struct {
	MPSCNNConvolutionNode
}

// MPSCNNBinaryConvolutionNodeFromID constructs a [MPSCNNBinaryConvolutionNode] from an objc.ID.
//
// A representation of a convolution kernel with binary weights and an input
// image using binary approximations.
func MPSCNNBinaryConvolutionNodeFromID(id objc.ID) MPSCNNBinaryConvolutionNode {
	return MPSCNNBinaryConvolutionNode{MPSCNNConvolutionNode: MPSCNNConvolutionNodeFromID(id)}
}

// NOTE: MPSCNNBinaryConvolutionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBinaryConvolutionNode] class.
//
// # Initializers
//
//   - [IMPSCNNBinaryConvolutionNode.InitWithSourceWeightsScaleValueTypeFlags]
//   - [IMPSCNNBinaryConvolutionNode.InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode
type IMPSCNNBinaryConvolutionNode interface {
	IMPSCNNConvolutionNode

	// Topic: Initializers

	InitWithSourceWeightsScaleValueTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode
	InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode
}

// Init initializes the instance.
func (c MPSCNNBinaryConvolutionNode) Init() MPSCNNBinaryConvolutionNode {
	rv := objc.Send[MPSCNNBinaryConvolutionNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryConvolutionNode) Autorelease() MPSCNNBinaryConvolutionNode {
	rv := objc.Send[MPSCNNBinaryConvolutionNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryConvolutionNode creates a new MPSCNNBinaryConvolutionNode instance.
func NewMPSCNNBinaryConvolutionNode() MPSCNNBinaryConvolutionNode {
	class := getMPSCNNBinaryConvolutionNodeClass()
	rv := objc.Send[MPSCNNBinaryConvolutionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/init(source:weights:)
func NewCNNBinaryConvolutionNodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNBinaryConvolutionNode {
	instance := getMPSCNNBinaryConvolutionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return MPSCNNBinaryConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/init(source:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func NewCNNBinaryConvolutionNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	instance := getMPSCNNBinaryConvolutionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return MPSCNNBinaryConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/init(source:weights:scaleValue:type:flags:)
func NewCNNBinaryConvolutionNodeWithSourceWeightsScaleValueTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	instance := getMPSCNNBinaryConvolutionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return MPSCNNBinaryConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/init(source:weights:scaleValue:type:flags:)
func (c MPSCNNBinaryConvolutionNode) InitWithSourceWeightsScaleValueTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	rv := objc.Send[MPSCNNBinaryConvolutionNode](c.ID, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/init(source:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func (c MPSCNNBinaryConvolutionNode) InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	rv := objc.Send[MPSCNNBinaryConvolutionNode](c.ID, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/nodeWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:
func (_MPSCNNBinaryConvolutionNodeClass MPSCNNBinaryConvolutionNodeClass) NodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNBinaryConvolutionNodeClass.class), objc.Sel("nodeWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return MPSCNNBinaryConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode/nodeWithSource:weights:scaleValue:type:flags:
func (_MPSCNNBinaryConvolutionNodeClass MPSCNNBinaryConvolutionNodeClass) NodeWithSourceWeightsScaleValueTypeFlags(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolutionNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNBinaryConvolutionNodeClass.class), objc.Sel("nodeWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return MPSCNNBinaryConvolutionNodeFromID(rv)
}
