// Code generated from Apple documentation. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
)

// MPSAccelerationStructureCompletionHandler is a block of code that’s invoked when an operation on an acceleration structure has completed.
//
// Deprecated: Deprecated since macOS 14.0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureCompletionHandler
type MPSAccelerationStructureCompletionHandler = func(MPSAccelerationStructure)

// MPSCNNArithmeticGradientStateBatch is a batch of arithmetic gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradientStateBatch
type MPSCNNArithmeticGradientStateBatch = []*MPSCNNArithmeticGradientState

// MPSCNNConvolutionGradientStateBatch is a batch of convolution gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientStateBatch
type MPSCNNConvolutionGradientStateBatch = []*MPSCNNConvolutionGradientState

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientStateBatch
type MPSCNNConvolutionTransposeGradientStateBatch = []*MPSCNNConvolutionTransposeGradientState

// MPSCNNDropoutGradientStateBatch is a batch of dropout gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientStateBatch
type MPSCNNDropoutGradientStateBatch = []*MPSCNNDropoutGradientState

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientStateBatch
type MPSCNNGroupNormalizationGradientStateBatch = []*MPSCNNGroupNormalizationGradientState

// MPSCNNInstanceNormalizationGradientStateBatch is a batch of instance normalization gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientStateBatch
type MPSCNNInstanceNormalizationGradientStateBatch = []*MPSCNNInstanceNormalizationGradientState

// MPSCNNLossLabelsBatch is a batch of loss labels instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabelsBatch
type MPSCNNLossLabelsBatch = []*MPSCNNLossLabels

// MPSCopyAllocator is a block to make a copy of a source texture for filters that can only execute out of place.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCopyAllocator
type MPSCopyAllocator = func(MPSKernel, metal.MTLCommandBuffer, metal.MTLTexture) metal.MTLTexture

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceCaps
type MPSDeviceCaps = uint32

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunctionConstant
type MPSFunctionConstant = int64

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunctionConstantInMetal
type MPSFunctionConstantInMetal = uint32

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGradientNodeBlock
type MPSGradientNodeBlock = func(MPSNNFilterNode, MPSNNFilterNode, MPSNNImageNode, MPSNNImageNode)

// MPSImageBatch is a batch of Metal Performance Shader image instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBatch
type MPSImageBatch = []*MPSImage

// MPSNNBinaryGradientStateBatch is a batch of binary gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientStateBatch
type MPSNNBinaryGradientStateBatch = []*MPSNNBinaryGradientState

// MPSNNGradientStateBatch is a batch of gradient state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientStateBatch
type MPSNNGradientStateBatch = []*MPSNNGradientState

// MPSNNGraphCompletionHandler is a notification when an asynchronous graph execution has finished.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraphCompletionHandler
type MPSNNGraphCompletionHandler = func(MPSImage, foundation.NSError)

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientStateBatch
type MPSNNMultiaryGradientStateBatch = []*MPSNNMultiaryGradientState

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSShape
type MPSShape = []*foundation.NSNumber

// MPSStateBatch is a batch of Metal Performance Shader state instances.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateBatch
type MPSStateBatch = []*MPSState
