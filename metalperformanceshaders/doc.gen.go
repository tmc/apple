// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

// Package metalperformanceshaders provides Go bindings for the MetalPerformanceShaders framework.
//
// Optimize graphics and compute performance with kernels that are fine-tuned
// for the unique characteristics of each Metal GPU family.
//
// The Metal Performance Shaders framework contains a collection of highly
// optimized compute and graphics shaders that are designed to integrate
// easily and efficiently into your Metal app. These data-parallel primitives
// are specially tuned to take advantage of the unique hardware
// characteristics of each GPU family to ensure optimal performance.
//
// # Fundamentals
//
//   - [The MPSKernel Class]
//   - [Tuning Hints]
//
// # Device Support
//
//   - [MPSSupportsMTLDevice]: Determines whether the Metal Performance Shaders framework supports a Metal device.
//
// # Image Filters
//
//   - [Image Filters]: Apply high-performance filters to, and extract statistical and histogram data from images. ([MPSImageAreaMax], [MPSImageDilate], [MPSImageAreaMin], [MPSImageErode], [MPSImageConvolution])
//
// # Neural Networks
//
//   - [Training a Neural Network with Metal Performance Shaders]: Use an MPS neural network graph to train a simple neural network digit classifier.
//   - [MPSImage]: A texture that may have more than four channels for use in convolutional neural networks. ([MPSImageDescriptor], [MPSPurgeableState], [MPSImageReadWriteParams], [MPSDataLayout], [MPSImageAllocator])
//   - [MPSTemporaryImage]: A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly. ([MPSImageDescriptor], [MPSImageAllocator])
//   - [Objects that Simplify the Creation of Neural Networks]: Simplify the creation of neural networks using networks of filter, image, and state nodes. ([MPSNNGraph], [MPSNNImageNode], [MPSHandle], [MPSNNAdditionNode], [MPSNNAdditionGradientNode])
//   - [Convolutional Neural Network Kernels]: Build neural networks with layers. ([MPSCNNAdd], [MPSCNNAddGradient], [MPSCNNSubtract], [MPSCNNSubtractGradient], [MPSCNNMultiply])
//   - [Recurrent Neural Networks]: Create recurrent neural networks. ([MPSRNNImageInferenceLayer], [MPSRNNMatrixInferenceLayer], [MPSRNNSingleGateDescriptor], [MPSGRUDescriptor], [MPSLSTMDescriptor])
//
// # Matrices and Vectors
//
//   - [Matrices and Vectors]: Solve systems of equations, factorize matrices and multiply matrices and vectors. ([MPSMatrix], [MPSMatrixDescriptor], [MPSTemporaryMatrix], [MPSVector], [MPSVectorDescriptor])
//
// # Kernel Base Classes
//
//   - [MPSKernel]: A standard interface for Metal Performance Shaders kernels. ([MPSKernelOptions])
//
// # Keyed Archivers
//
//   - [MPSKeyedUnarchiver]: A keyed archiver that supports Metal Performance Shaders kernel decoding.
//   - [MPSDeviceProvider]: An interface that enables the setting of a Metal device for unarchived objects.
//
// # Ray Tracing
//
//   - [Accelerating ray tracing and motion blur using Metal]: Generate ray-traced images with motion blur using GPU-based parallel processing.
//
// # Classes
//
//   - [MPSCNNConvolutionTransposeGradient]
//   - [MPSCNNConvolutionTransposeGradientNode]
//   - [MPSCNNConvolutionTransposeGradientState]
//   - [MPSCNNConvolutionTransposeGradientStateNode]
//   - [MPSCNNFullyConnectedGradientNode]
//   - [MPSCNNGroupNormalization]
//   - [MPSCNNGroupNormalizationGradient]
//   - [MPSCNNGroupNormalizationGradientNode]
//   - [MPSCNNGroupNormalizationGradientState]
//   - [MPSCNNGroupNormalizationNode]
//   - [MPSCNNMultiaryKernel]
//   - [MPSCNNNeuronGeLUNode]
//   - [MPSCommandBuffer]
//   - [MPSFColorConversion]
//   - [MPSFunction]
//   - [MPSImageCanny]
//   - [MPSImageEDLines]
//   - [MPSImageNormalizedHistogram]: A filter that computes the normalized histogram of an image.
//   - [MPSMatrixRandom]
//   - [MPSMatrixRandomDistributionDescriptor]
//   - [MPSMatrixRandomMTGP32]
//   - [MPSMatrixRandomPhilox]
//   - [MPSNDArray]
//   - [MPSNDArrayAffineInt4Dequantize]
//   - [MPSNDArrayAffineQuantizationDescriptor]
//   - [MPSNDArrayBinaryKernel]
//   - [MPSNDArrayBinaryPrimaryGradientKernel]
//   - [MPSNDArrayBinarySecondaryGradientKernel]
//   - [MPSNDArrayDescriptor]
//   - [MPSNDArrayGather]
//   - [MPSNDArrayGatherGradient]
//   - [MPSNDArrayGatherGradientState]
//   - [MPSNDArrayGradientState]
//   - [MPSNDArrayIdentity]
//   - [MPSNDArrayLUTDequantize]
//   - [MPSNDArrayLUTQuantizationDescriptor]
//   - [MPSNDArrayMatrixMultiplication]
//   - [MPSNDArrayMultiaryBase]
//   - [MPSNDArrayMultiaryGradientKernel]
//   - [MPSNDArrayMultiaryKernel]
//   - [MPSNDArrayQuantizationDescriptor]
//   - [MPSNDArrayQuantizedMatrixMultiplication]
//   - [MPSNDArrayStridedSlice]
//   - [MPSNDArrayStridedSliceGradient]
//   - [MPSNDArrayUnaryGradientKernel]
//   - [MPSNDArrayUnaryKernel]
//   - [MPSNDArrayVectorLUTDequantize]
//   - [MPSNNCompare]
//   - [MPSNNComparisonNode]
//   - [MPSNNCropAndResizeBilinear]: A cropping and bilinear resizing filter.
//   - [MPSNNForwardLoss]
//   - [MPSNNForwardLossNode]
//   - [MPSNNGramMatrixCalculation]
//   - [MPSNNGramMatrixCalculationGradient]
//   - [MPSNNGramMatrixCalculationGradientNode]
//   - [MPSNNGramMatrixCalculationNode]
//   - [MPSNNGridSample]
//   - [MPSNNInitialGradient]
//   - [MPSNNInitialGradientNode]
//   - [MPSNNLocalCorrelation]
//   - [MPSNNLossGradient]
//   - [MPSNNLossGradientNode]
//   - [MPSNNMultiaryGradientState]
//   - [MPSNNMultiaryGradientStateNode]
//   - [MPSNNPad]
//   - [MPSNNPadGradient]
//   - [MPSNNPadGradientNode]
//   - [MPSNNPadNode]
//   - [MPSNNReductionColumnMaxNode]
//   - [MPSNNReductionColumnMeanNode]
//   - [MPSNNReductionColumnMinNode]
//   - [MPSNNReductionColumnSumNode]
//   - [MPSNNReductionFeatureChannelsArgumentMaxNode]
//   - [MPSNNReductionFeatureChannelsArgumentMinNode]
//   - [MPSNNReductionFeatureChannelsMaxNode]
//   - [MPSNNReductionFeatureChannelsMeanNode]
//   - [MPSNNReductionFeatureChannelsMinNode]
//   - [MPSNNReductionFeatureChannelsSumNode]
//   - [MPSNNReductionRowMaxNode]
//   - [MPSNNReductionRowMeanNode]
//   - [MPSNNReductionRowMinNode]
//   - [MPSNNReductionRowSumNode]
//   - [MPSNNReductionSpatialMeanGradientNode]
//   - [MPSNNReductionSpatialMeanNode]
//   - [MPSNNReshapeGradient]
//   - [MPSNNReshapeGradientNode]
//   - [MPSNNReshapeNode]
//   - [MPSNNResizeBilinear]: A bilinear resizing filter.
//   - [MPSNNUnaryReductionNode]
//   - [MPSPredicate]
//   - [MPSSVGF]
//   - [MPSSVGFDefaultTextureAllocator]
//   - [MPSSVGFDenoiser]
//   - [MPSStateResourceList]: An interface for objects that define resources for Metal Performance Shaders state containers.
//   - [MPSTemporalAA]
//   - [MPSTemporaryNDArray]
//
// # Protocols
//
//   - [MPSCNNGroupNormalizationDataSource]
//   - [MPSHeapProvider]
//   - [MPSNDArrayAllocator]
//   - [MPSNNGramMatrixCallback]
//   - [MPSNNLossCallback]
//   - [MPSSVGFTextureAllocator]
//
// # Variables
//
//   - [MPSRectNoClip]
//
// # Enumerations
//
//   - [MPSFColorConversionOptions]//
//
// # Key Types
//
//   - [MPSCNNBinaryKernel] - A convolution neural network kernel.
//   - [MPSCNNMultiaryKernel]
//   - [MPSCNNKernel] - Base class for neural network layers.
//   - [MPSImage] - A texture that may have more than four channels for use in convolutional neural networks.
//   - [MPSNDArray]
//   - [MPSCNNConvolutionTranspose] - A transposed convolution kernel.
//   - [MPSNNGraph] - An optimized representation of a graph of neural network image and filter nodes.
//   - [MPSSVGF]
//   - [MPSCNNYOLOLoss] - A kernel that computes the YOLO loss and loss gradient between specified predictions and labels.
//   - [MPSRayIntersector] - A kernel that performs intersection tests between rays and geometry.
//
// [Accelerating ray tracing and motion blur using Metal]: https://developer.apple.com/documentation/metal/accelerating-ray-tracing-and-motion-blur-using-metal
// [Convolutional Neural Network Kernels]: https://developer.apple.com/documentation/metalperformanceshaders/convolutional-neural-network-kernels
// [Image Filters]: https://developer.apple.com/documentation/metalperformanceshaders/image-filters
// [Matrices and Vectors]: https://developer.apple.com/documentation/metalperformanceshaders/matrices-and-vectors
// [Objects that Simplify the Creation of Neural Networks]: https://developer.apple.com/documentation/metalperformanceshaders/objects-that-simplify-the-creation-of-neural-networks
// [Recurrent Neural Networks]: https://developer.apple.com/documentation/metalperformanceshaders/recurrent-neural-networks
// [The MPSKernel Class]: https://developer.apple.com/documentation/metalperformanceshaders/the-mpskernel-class
// [Training a Neural Network with Metal Performance Shaders]: https://developer.apple.com/documentation/metalperformanceshaders/training-a-neural-network-with-metal-performance-shaders
// [Tuning Hints]: https://developer.apple.com/documentation/metalperformanceshaders/tuning-hints
package metalperformanceshaders

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the MetalPerformanceShaders library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/MetalPerformanceShaders.framework/MetalPerformanceShaders",
	"/usr/lib/libMetalPerformanceShaders.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: MetalPerformanceShaders: failed to load framework from any known path\n")
	}
}
