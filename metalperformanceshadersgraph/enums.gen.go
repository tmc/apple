// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeploymentPlatform
type MPSGraphDeploymentPlatform uint64

const (
	// MPSGraphDeploymentPlatformIOS: Deployment target for iOS.
	MPSGraphDeploymentPlatformIOS MPSGraphDeploymentPlatform = 1
	// MPSGraphDeploymentPlatformMacOS: Deployment platofmr for macOS.
	MPSGraphDeploymentPlatformMacOS MPSGraphDeploymentPlatform = 0
	// MPSGraphDeploymentPlatformTvOS: Deployment target for tvOS.
	MPSGraphDeploymentPlatformTvOS MPSGraphDeploymentPlatform = 2
	// MPSGraphDeploymentPlatformVisionOS: Deployment target for visionOS.
	MPSGraphDeploymentPlatformVisionOS MPSGraphDeploymentPlatform = 3
)

func (e MPSGraphDeploymentPlatform) String() string {
	switch e {
	case MPSGraphDeploymentPlatformIOS:
		return "MPSGraphDeploymentPlatformIOS"
	case MPSGraphDeploymentPlatformMacOS:
		return "MPSGraphDeploymentPlatformMacOS"
	case MPSGraphDeploymentPlatformTvOS:
		return "MPSGraphDeploymentPlatformTvOS"
	case MPSGraphDeploymentPlatformVisionOS:
		return "MPSGraphDeploymentPlatformVisionOS"
	default:
		return fmt.Sprintf("MPSGraphDeploymentPlatform(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDeviceType
type MPSGraphDeviceType uint32

const (
	// MPSGraphDeviceTypeMetal: Device of type Metal
	MPSGraphDeviceTypeMetal MPSGraphDeviceType = 0
)

func (e MPSGraphDeviceType) String() string {
	switch e {
	case MPSGraphDeviceTypeMetal:
		return "MPSGraphDeviceTypeMetal"
	default:
		return fmt.Sprintf("MPSGraphDeviceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionStage
type MPSGraphExecutionStage uint64

const (
	// MPSGraphExecutionStageCompleted: stage when execution of the graph completes.
	MPSGraphExecutionStageCompleted MPSGraphExecutionStage = 0
)

func (e MPSGraphExecutionStage) String() string {
	switch e {
	case MPSGraphExecutionStageCompleted:
		return "MPSGraphExecutionStageCompleted"
	default:
		return fmt.Sprintf("MPSGraphExecutionStage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTScalingMode
type MPSGraphFFTScalingMode uint

const (
	// MPSGraphFFTScalingModeNone: Computes the FFT result with no scaling.
	MPSGraphFFTScalingModeNone MPSGraphFFTScalingMode = 0
	// MPSGraphFFTScalingModeSize: Scales the FFT result with reciprocal of the total FFT size over all transformed dimensions.
	MPSGraphFFTScalingModeSize MPSGraphFFTScalingMode = 1
	// MPSGraphFFTScalingModeUnitary: Scales the FFT result with reciprocal square root of the total FFT size over all transformed dimensions, resulting in signal strength conserving transformation.
	MPSGraphFFTScalingModeUnitary MPSGraphFFTScalingMode = 2
)

func (e MPSGraphFFTScalingMode) String() string {
	switch e {
	case MPSGraphFFTScalingModeNone:
		return "MPSGraphFFTScalingModeNone"
	case MPSGraphFFTScalingModeSize:
		return "MPSGraphFFTScalingModeSize"
	case MPSGraphFFTScalingModeUnitary:
		return "MPSGraphFFTScalingModeUnitary"
	default:
		return fmt.Sprintf("MPSGraphFFTScalingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType
type MPSGraphLossReductionType uint64

const (
	// MPSGraphLossReductionTypeAxis: Computes the loss without reduction.
	MPSGraphLossReductionTypeAxis MPSGraphLossReductionType = 0
	// MPSGraphLossReductionTypeMean: Reduces the loss down to a scalar with a mean operation.
	MPSGraphLossReductionTypeMean MPSGraphLossReductionType = 2
	// MPSGraphLossReductionTypeNone: Computes the loss without reduction.
	MPSGraphLossReductionTypeNone MPSGraphLossReductionType = 0
	// MPSGraphLossReductionTypeSum: Reduces the loss down to a scalar with a sum operation.
	MPSGraphLossReductionTypeSum MPSGraphLossReductionType = 1
)

func (e MPSGraphLossReductionType) String() string {
	switch e {
	case MPSGraphLossReductionTypeAxis:
		return "MPSGraphLossReductionTypeAxis"
	case MPSGraphLossReductionTypeMean:
		return "MPSGraphLossReductionTypeMean"
	case MPSGraphLossReductionTypeSum:
		return "MPSGraphLossReductionTypeSum"
	default:
		return fmt.Sprintf("MPSGraphLossReductionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphNonMaximumSuppressionCoordinateMode
type MPSGraphNonMaximumSuppressionCoordinateMode uint

const (
	MPSGraphNonMaximumSuppressionCoordinateModeCentersHeightFirst MPSGraphNonMaximumSuppressionCoordinateMode = 2
	MPSGraphNonMaximumSuppressionCoordinateModeCentersWidthFirst  MPSGraphNonMaximumSuppressionCoordinateMode = 3
	MPSGraphNonMaximumSuppressionCoordinateModeCornersHeightFirst MPSGraphNonMaximumSuppressionCoordinateMode = 0
	MPSGraphNonMaximumSuppressionCoordinateModeCornersWidthFirst  MPSGraphNonMaximumSuppressionCoordinateMode = 1
)

func (e MPSGraphNonMaximumSuppressionCoordinateMode) String() string {
	switch e {
	case MPSGraphNonMaximumSuppressionCoordinateModeCentersHeightFirst:
		return "MPSGraphNonMaximumSuppressionCoordinateModeCentersHeightFirst"
	case MPSGraphNonMaximumSuppressionCoordinateModeCentersWidthFirst:
		return "MPSGraphNonMaximumSuppressionCoordinateModeCentersWidthFirst"
	case MPSGraphNonMaximumSuppressionCoordinateModeCornersHeightFirst:
		return "MPSGraphNonMaximumSuppressionCoordinateModeCornersHeightFirst"
	case MPSGraphNonMaximumSuppressionCoordinateModeCornersWidthFirst:
		return "MPSGraphNonMaximumSuppressionCoordinateModeCornersWidthFirst"
	default:
		return fmt.Sprintf("MPSGraphNonMaximumSuppressionCoordinateMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimization
type MPSGraphOptimization uint64

const (
	// MPSGraphOptimizationLevel0: Graph performs core optimizations only.
	MPSGraphOptimizationLevel0 MPSGraphOptimization = 0
	// MPSGraphOptimizationLevel1: Graph performs additional Optimizations, like using the placement pass to dispatch across different HW blocks like the NeuralEngine and CPU along with the GPU.
	MPSGraphOptimizationLevel1 MPSGraphOptimization = 1
)

func (e MPSGraphOptimization) String() string {
	switch e {
	case MPSGraphOptimizationLevel0:
		return "MPSGraphOptimizationLevel0"
	case MPSGraphOptimizationLevel1:
		return "MPSGraphOptimizationLevel1"
	default:
		return fmt.Sprintf("MPSGraphOptimization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptimizationProfile
type MPSGraphOptimizationProfile uint64

const (
	// MPSGraphOptimizationProfilePerformance: Default, graph optimized for performance.
	MPSGraphOptimizationProfilePerformance MPSGraphOptimizationProfile = 0
	// MPSGraphOptimizationProfilePowerEfficiency: Graph optimized for power efficiency.
	MPSGraphOptimizationProfilePowerEfficiency MPSGraphOptimizationProfile = 1
)

func (e MPSGraphOptimizationProfile) String() string {
	switch e {
	case MPSGraphOptimizationProfilePerformance:
		return "MPSGraphOptimizationProfilePerformance"
	case MPSGraphOptimizationProfilePowerEfficiency:
		return "MPSGraphOptimizationProfilePowerEfficiency"
	default:
		return fmt.Sprintf("MPSGraphOptimizationProfile(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOptions
type MPSGraphOptions uint64

const (
	// MPSGraphOptionsDefault: The framework uses these options as default if not overriden.
	MPSGraphOptionsDefault MPSGraphOptions = 1
	// MPSGraphOptionsNone: No Options.
	MPSGraphOptionsNone MPSGraphOptions = 0
	// MPSGraphOptionsSynchronizeResults: The graph synchronizes results to the CPU using a blit encoder if on a discrete GPU at the end of execution.
	MPSGraphOptionsSynchronizeResults MPSGraphOptions = 1
	// MPSGraphOptionsVerbose: The framework prints more logging info.
	MPSGraphOptionsVerbose MPSGraphOptions = 2
)

func (e MPSGraphOptions) String() string {
	switch e {
	case MPSGraphOptionsDefault:
		return "MPSGraphOptionsDefault"
	case MPSGraphOptionsNone:
		return "MPSGraphOptionsNone"
	case MPSGraphOptionsVerbose:
		return "MPSGraphOptionsVerbose"
	default:
		return fmt.Sprintf("MPSGraphOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingMode
type MPSGraphPaddingMode int

const (
	// MPSGraphPaddingModeAntiPeriodic: Anti Periodic `x[-2] -> -x[L-3]`
	MPSGraphPaddingModeAntiPeriodic MPSGraphPaddingMode = 6
	// MPSGraphPaddingModeClampToEdge: ClampToEdge (PyTorch ReplicationPad)
	MPSGraphPaddingModeClampToEdge MPSGraphPaddingMode = 3
	// MPSGraphPaddingModeConstant: Constant
	MPSGraphPaddingModeConstant MPSGraphPaddingMode = 0
	// MPSGraphPaddingModePeriodic: Periodic `x[-2] -> x[L-3], where L is size of x.`
	MPSGraphPaddingModePeriodic MPSGraphPaddingMode = 5
	// MPSGraphPaddingModeReflect: Reflect
	MPSGraphPaddingModeReflect MPSGraphPaddingMode = 1
	// MPSGraphPaddingModeSymmetric: Symmetric
	MPSGraphPaddingModeSymmetric MPSGraphPaddingMode = 2
	// MPSGraphPaddingModeZero: Zero
	MPSGraphPaddingModeZero MPSGraphPaddingMode = 4
)

func (e MPSGraphPaddingMode) String() string {
	switch e {
	case MPSGraphPaddingModeAntiPeriodic:
		return "MPSGraphPaddingModeAntiPeriodic"
	case MPSGraphPaddingModeClampToEdge:
		return "MPSGraphPaddingModeClampToEdge"
	case MPSGraphPaddingModeConstant:
		return "MPSGraphPaddingModeConstant"
	case MPSGraphPaddingModePeriodic:
		return "MPSGraphPaddingModePeriodic"
	case MPSGraphPaddingModeReflect:
		return "MPSGraphPaddingModeReflect"
	case MPSGraphPaddingModeSymmetric:
		return "MPSGraphPaddingModeSymmetric"
	case MPSGraphPaddingModeZero:
		return "MPSGraphPaddingModeZero"
	default:
		return fmt.Sprintf("MPSGraphPaddingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle
type MPSGraphPaddingStyle uint

const (
	// MPSGraphPaddingStyleExplicit: Explicit
	MPSGraphPaddingStyleExplicit MPSGraphPaddingStyle = 0
	// MPSGraphPaddingStyleExplicitOffset: TF_VALID
	MPSGraphPaddingStyleExplicitOffset MPSGraphPaddingStyle = 3
	// MPSGraphPaddingStyleONNX_SAME_LOWER: Explicit offsets
	MPSGraphPaddingStyleONNX_SAME_LOWER MPSGraphPaddingStyle = 4
	// MPSGraphPaddingStyleTF_SAME: TF_SAME
	MPSGraphPaddingStyleTF_SAME MPSGraphPaddingStyle = 2
	// MPSGraphPaddingStyleTF_VALID: ONNX_SAME_LOWER
	MPSGraphPaddingStyleTF_VALID MPSGraphPaddingStyle = 1
)

func (e MPSGraphPaddingStyle) String() string {
	switch e {
	case MPSGraphPaddingStyleExplicit:
		return "MPSGraphPaddingStyleExplicit"
	case MPSGraphPaddingStyleExplicitOffset:
		return "MPSGraphPaddingStyleExplicitOffset"
	case MPSGraphPaddingStyleONNX_SAME_LOWER:
		return "MPSGraphPaddingStyleONNX_SAME_LOWER"
	case MPSGraphPaddingStyleTF_SAME:
		return "MPSGraphPaddingStyleTF_SAME"
	case MPSGraphPaddingStyleTF_VALID:
		return "MPSGraphPaddingStyleTF_VALID"
	default:
		return fmt.Sprintf("MPSGraphPaddingStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPoolingReturnIndicesMode
type MPSGraphPoolingReturnIndicesMode uint

const (
	// MPSGraphPoolingReturnIndicesGlobalFlatten1D: Returns indices flattened in inner most (last) dimension.
	MPSGraphPoolingReturnIndicesGlobalFlatten1D MPSGraphPoolingReturnIndicesMode = 1
	// MPSGraphPoolingReturnIndicesGlobalFlatten2D: Returns indices flattened in 2 innermost dimensions.
	MPSGraphPoolingReturnIndicesGlobalFlatten2D MPSGraphPoolingReturnIndicesMode = 2
	// MPSGraphPoolingReturnIndicesGlobalFlatten3D: Returns indices flattened in 3 innernost dimensions.
	MPSGraphPoolingReturnIndicesGlobalFlatten3D MPSGraphPoolingReturnIndicesMode = 3
	// MPSGraphPoolingReturnIndicesGlobalFlatten4D: Returns indices flattened in 4 innermost dimensions.
	MPSGraphPoolingReturnIndicesGlobalFlatten4D MPSGraphPoolingReturnIndicesMode = 4
	// MPSGraphPoolingReturnIndicesLocalFlatten1D: Returns indices within pooling window, flattened in inner most dimension.
	MPSGraphPoolingReturnIndicesLocalFlatten1D MPSGraphPoolingReturnIndicesMode = 5
	// MPSGraphPoolingReturnIndicesLocalFlatten2D: Returns indices within pooling window, flattened in 2 innermost dimensions.
	MPSGraphPoolingReturnIndicesLocalFlatten2D MPSGraphPoolingReturnIndicesMode = 6
	// MPSGraphPoolingReturnIndicesLocalFlatten3D: Returns indices within pooling window, flattened in 3 innernost dimensions.
	MPSGraphPoolingReturnIndicesLocalFlatten3D MPSGraphPoolingReturnIndicesMode = 7
	// MPSGraphPoolingReturnIndicesLocalFlatten4D: Returns indices within pooling window, flattened in 4 innermost dimensions.
	MPSGraphPoolingReturnIndicesLocalFlatten4D MPSGraphPoolingReturnIndicesMode = 8
	// MPSGraphPoolingReturnIndicesNone: No indices returned.
	MPSGraphPoolingReturnIndicesNone MPSGraphPoolingReturnIndicesMode = 0
)

func (e MPSGraphPoolingReturnIndicesMode) String() string {
	switch e {
	case MPSGraphPoolingReturnIndicesGlobalFlatten1D:
		return "MPSGraphPoolingReturnIndicesGlobalFlatten1D"
	case MPSGraphPoolingReturnIndicesGlobalFlatten2D:
		return "MPSGraphPoolingReturnIndicesGlobalFlatten2D"
	case MPSGraphPoolingReturnIndicesGlobalFlatten3D:
		return "MPSGraphPoolingReturnIndicesGlobalFlatten3D"
	case MPSGraphPoolingReturnIndicesGlobalFlatten4D:
		return "MPSGraphPoolingReturnIndicesGlobalFlatten4D"
	case MPSGraphPoolingReturnIndicesLocalFlatten1D:
		return "MPSGraphPoolingReturnIndicesLocalFlatten1D"
	case MPSGraphPoolingReturnIndicesLocalFlatten2D:
		return "MPSGraphPoolingReturnIndicesLocalFlatten2D"
	case MPSGraphPoolingReturnIndicesLocalFlatten3D:
		return "MPSGraphPoolingReturnIndicesLocalFlatten3D"
	case MPSGraphPoolingReturnIndicesLocalFlatten4D:
		return "MPSGraphPoolingReturnIndicesLocalFlatten4D"
	case MPSGraphPoolingReturnIndicesNone:
		return "MPSGraphPoolingReturnIndicesNone"
	default:
		return fmt.Sprintf("MPSGraphPoolingReturnIndicesMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRNNActivation
type MPSGraphRNNActivation uint

const (
	// MPSGraphRNNActivationHardSigmoid: Defines a Hard sigmoid activation.
	MPSGraphRNNActivationHardSigmoid MPSGraphRNNActivation = 4
	// MPSGraphRNNActivationNone: Defines a pass through activation.
	MPSGraphRNNActivationNone MPSGraphRNNActivation = 0
	// MPSGraphRNNActivationRelu: Defines a ReLU activation.
	MPSGraphRNNActivationRelu MPSGraphRNNActivation = 1
	// MPSGraphRNNActivationSigmoid: Defines a Sigmoid activation.
	MPSGraphRNNActivationSigmoid MPSGraphRNNActivation = 3
	// MPSGraphRNNActivationTanh: Defines a Tanh activation.
	MPSGraphRNNActivationTanh MPSGraphRNNActivation = 2
)

func (e MPSGraphRNNActivation) String() string {
	switch e {
	case MPSGraphRNNActivationHardSigmoid:
		return "MPSGraphRNNActivationHardSigmoid"
	case MPSGraphRNNActivationNone:
		return "MPSGraphRNNActivationNone"
	case MPSGraphRNNActivationRelu:
		return "MPSGraphRNNActivationRelu"
	case MPSGraphRNNActivationSigmoid:
		return "MPSGraphRNNActivationSigmoid"
	case MPSGraphRNNActivationTanh:
		return "MPSGraphRNNActivationTanh"
	default:
		return fmt.Sprintf("MPSGraphRNNActivation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomDistribution
type MPSGraphRandomDistribution uint64

const (
	// MPSGraphRandomDistributionNormal: The normal distribution defined by mean and standard deviation.
	MPSGraphRandomDistributionNormal MPSGraphRandomDistribution = 1
	// MPSGraphRandomDistributionTruncatedNormal: The normal distribution defined by mean and standard deviation, truncated to the range [min, max)
	MPSGraphRandomDistributionTruncatedNormal MPSGraphRandomDistribution = 2
	// MPSGraphRandomDistributionUniform: The uniform distribution, with samples drawn uniformly from [min, max) for float types, and [min, max] for integer types.
	MPSGraphRandomDistributionUniform MPSGraphRandomDistribution = 0
)

func (e MPSGraphRandomDistribution) String() string {
	switch e {
	case MPSGraphRandomDistributionNormal:
		return "MPSGraphRandomDistributionNormal"
	case MPSGraphRandomDistributionTruncatedNormal:
		return "MPSGraphRandomDistributionTruncatedNormal"
	case MPSGraphRandomDistributionUniform:
		return "MPSGraphRandomDistributionUniform"
	default:
		return fmt.Sprintf("MPSGraphRandomDistribution(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomNormalSamplingMethod
type MPSGraphRandomNormalSamplingMethod uint64

const (
	// MPSGraphRandomNormalSamplingBoxMuller: Use Box Muller transform to convert uniform values to values in the normal distribution.
	MPSGraphRandomNormalSamplingBoxMuller MPSGraphRandomNormalSamplingMethod = 1
	// MPSGraphRandomNormalSamplingInvCDF: Use inverse erf to convert uniform values to values in the normal distribution
	MPSGraphRandomNormalSamplingInvCDF MPSGraphRandomNormalSamplingMethod = 0
)

func (e MPSGraphRandomNormalSamplingMethod) String() string {
	switch e {
	case MPSGraphRandomNormalSamplingBoxMuller:
		return "MPSGraphRandomNormalSamplingBoxMuller"
	case MPSGraphRandomNormalSamplingInvCDF:
		return "MPSGraphRandomNormalSamplingInvCDF"
	default:
		return fmt.Sprintf("MPSGraphRandomNormalSamplingMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReducedPrecisionFastMath
type MPSGraphReducedPrecisionFastMath uint

const (
	// MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate: Execute winograd transform intermediate as FP16.
	MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate MPSGraphReducedPrecisionFastMath = 2
	// MPSGraphReducedPrecisionFastMathAllowFP16Intermediates: Curated list allowing intermediates for multi-pass GPU kernels to be FP16.
	MPSGraphReducedPrecisionFastMathAllowFP16Intermediates MPSGraphReducedPrecisionFastMath = 2
	// MPSGraphReducedPrecisionFastMathDefault: Default selection.
	MPSGraphReducedPrecisionFastMathDefault MPSGraphReducedPrecisionFastMath = 0
	// MPSGraphReducedPrecisionFastMathNone: Full precision math with maximum accuracy.
	MPSGraphReducedPrecisionFastMathNone MPSGraphReducedPrecisionFastMath = 0
)

func (e MPSGraphReducedPrecisionFastMath) String() string {
	switch e {
	case MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate:
		return "MPSGraphReducedPrecisionFastMathAllowFP16Conv2DWinogradTransformIntermediate"
	case MPSGraphReducedPrecisionFastMathDefault:
		return "MPSGraphReducedPrecisionFastMathDefault"
	default:
		return fmt.Sprintf("MPSGraphReducedPrecisionFastMath(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphReductionMode
type MPSGraphReductionMode uint

const (
	// MPSGraphReductionModeArgumentMax: Argument Max
	MPSGraphReductionModeArgumentMax MPSGraphReductionMode = 5
	// MPSGraphReductionModeArgumentMin: Argument Min
	MPSGraphReductionModeArgumentMin MPSGraphReductionMode = 4
	// MPSGraphReductionModeMax: Max
	MPSGraphReductionModeMax MPSGraphReductionMode = 1
	// MPSGraphReductionModeMin: Min
	MPSGraphReductionModeMin MPSGraphReductionMode = 0
	// MPSGraphReductionModeProduct: Product
	MPSGraphReductionModeProduct MPSGraphReductionMode = 3
	// MPSGraphReductionModeSum: Sum
	MPSGraphReductionModeSum MPSGraphReductionMode = 2
)

func (e MPSGraphReductionMode) String() string {
	switch e {
	case MPSGraphReductionModeArgumentMax:
		return "MPSGraphReductionModeArgumentMax"
	case MPSGraphReductionModeArgumentMin:
		return "MPSGraphReductionModeArgumentMin"
	case MPSGraphReductionModeMax:
		return "MPSGraphReductionModeMax"
	case MPSGraphReductionModeMin:
		return "MPSGraphReductionModeMin"
	case MPSGraphReductionModeProduct:
		return "MPSGraphReductionModeProduct"
	case MPSGraphReductionModeSum:
		return "MPSGraphReductionModeSum"
	default:
		return fmt.Sprintf("MPSGraphReductionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeMode
type MPSGraphResizeMode uint

const (
	// MPSGraphResizeBilinear: Samples the 4 neighbors to the pixel coordinate and uses bilinear interpolation.
	MPSGraphResizeBilinear MPSGraphResizeMode = 1
	// MPSGraphResizeNearest: Samples the nearest neighbor to the pixel coordinate.
	MPSGraphResizeNearest MPSGraphResizeMode = 0
)

func (e MPSGraphResizeMode) String() string {
	switch e {
	case MPSGraphResizeBilinear:
		return "MPSGraphResizeBilinear"
	case MPSGraphResizeNearest:
		return "MPSGraphResizeNearest"
	default:
		return fmt.Sprintf("MPSGraphResizeMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphResizeNearestRoundingMode
type MPSGraphResizeNearestRoundingMode uint

const (
	// MPSGraphResizeNearestRoundingModeCeil: Rounds values toward +inf.
	MPSGraphResizeNearestRoundingModeCeil MPSGraphResizeNearestRoundingMode = 2
	// MPSGraphResizeNearestRoundingModeFloor: Rounds values toward -inf.
	MPSGraphResizeNearestRoundingModeFloor MPSGraphResizeNearestRoundingMode = 3
	// MPSGraphResizeNearestRoundingModeRoundPreferCeil: Rounds values to the nearest integer value, with 0.5f offset rounding toward +inf.
	MPSGraphResizeNearestRoundingModeRoundPreferCeil MPSGraphResizeNearestRoundingMode = 0
	// MPSGraphResizeNearestRoundingModeRoundPreferFloor: Rounds values to the nearest integer value, with 0.5f rounding toward -inf.
	MPSGraphResizeNearestRoundingModeRoundPreferFloor MPSGraphResizeNearestRoundingMode = 1
	// MPSGraphResizeNearestRoundingModeRoundToEven: Rounds values to the nearest integer value, with 0.5f rounding toward the closest even value.
	MPSGraphResizeNearestRoundingModeRoundToEven MPSGraphResizeNearestRoundingMode = 4
	// MPSGraphResizeNearestRoundingModeRoundToOdd: Rounds values to the nearest integer value, with 0.5f rounding toward the closest odd value.
	MPSGraphResizeNearestRoundingModeRoundToOdd MPSGraphResizeNearestRoundingMode = 5
)

func (e MPSGraphResizeNearestRoundingMode) String() string {
	switch e {
	case MPSGraphResizeNearestRoundingModeCeil:
		return "MPSGraphResizeNearestRoundingModeCeil"
	case MPSGraphResizeNearestRoundingModeFloor:
		return "MPSGraphResizeNearestRoundingModeFloor"
	case MPSGraphResizeNearestRoundingModeRoundPreferCeil:
		return "MPSGraphResizeNearestRoundingModeRoundPreferCeil"
	case MPSGraphResizeNearestRoundingModeRoundPreferFloor:
		return "MPSGraphResizeNearestRoundingModeRoundPreferFloor"
	case MPSGraphResizeNearestRoundingModeRoundToEven:
		return "MPSGraphResizeNearestRoundingModeRoundToEven"
	case MPSGraphResizeNearestRoundingModeRoundToOdd:
		return "MPSGraphResizeNearestRoundingModeRoundToOdd"
	default:
		return fmt.Sprintf("MPSGraphResizeNearestRoundingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphScatterMode
type MPSGraphScatterMode int

const (
	// MPSGraphScatterModeAdd: Add
	MPSGraphScatterModeAdd MPSGraphScatterMode = 0
	// MPSGraphScatterModeDiv: Divide
	MPSGraphScatterModeDiv MPSGraphScatterMode = 3
	// MPSGraphScatterModeMax: Maximum
	MPSGraphScatterModeMax MPSGraphScatterMode = 5
	// MPSGraphScatterModeMin: Minimum
	MPSGraphScatterModeMin MPSGraphScatterMode = 4
	// MPSGraphScatterModeMul: Multiply
	MPSGraphScatterModeMul MPSGraphScatterMode = 2
	// MPSGraphScatterModeSet: Set
	MPSGraphScatterModeSet MPSGraphScatterMode = 6
	// MPSGraphScatterModeSub: Sub
	MPSGraphScatterModeSub MPSGraphScatterMode = 1
)

func (e MPSGraphScatterMode) String() string {
	switch e {
	case MPSGraphScatterModeAdd:
		return "MPSGraphScatterModeAdd"
	case MPSGraphScatterModeDiv:
		return "MPSGraphScatterModeDiv"
	case MPSGraphScatterModeMax:
		return "MPSGraphScatterModeMax"
	case MPSGraphScatterModeMin:
		return "MPSGraphScatterModeMin"
	case MPSGraphScatterModeMul:
		return "MPSGraphScatterModeMul"
	case MPSGraphScatterModeSet:
		return "MPSGraphScatterModeSet"
	case MPSGraphScatterModeSub:
		return "MPSGraphScatterModeSub"
	default:
		return fmt.Sprintf("MPSGraphScatterMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSparseStorageType
type MPSGraphSparseStorageType uint64

const (
	// MPSGraphSparseStorageCOO: COO Storage
	MPSGraphSparseStorageCOO MPSGraphSparseStorageType = 0
	// MPSGraphSparseStorageCSC: CSC Storage
	MPSGraphSparseStorageCSC MPSGraphSparseStorageType = 1
	// MPSGraphSparseStorageCSR: CSR Storage
	MPSGraphSparseStorageCSR MPSGraphSparseStorageType = 2
)

func (e MPSGraphSparseStorageType) String() string {
	switch e {
	case MPSGraphSparseStorageCOO:
		return "MPSGraphSparseStorageCOO"
	case MPSGraphSparseStorageCSC:
		return "MPSGraphSparseStorageCSC"
	case MPSGraphSparseStorageCSR:
		return "MPSGraphSparseStorageCSR"
	default:
		return fmt.Sprintf("MPSGraphSparseStorageType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
type MPSGraphTensorNamedDataLayout uint

const (
	// MPSGraphTensorNamedDataLayoutCHW: LayoutCHW
	MPSGraphTensorNamedDataLayoutCHW MPSGraphTensorNamedDataLayout = 4
	// MPSGraphTensorNamedDataLayoutDHWIO: LayoutDHWIO
	MPSGraphTensorNamedDataLayoutDHWIO MPSGraphTensorNamedDataLayout = 10
	// MPSGraphTensorNamedDataLayoutHW: LayoutHW
	MPSGraphTensorNamedDataLayoutHW MPSGraphTensorNamedDataLayout = 6
	// MPSGraphTensorNamedDataLayoutHWC: LayoutHWC
	MPSGraphTensorNamedDataLayoutHWC MPSGraphTensorNamedDataLayout = 5
	// MPSGraphTensorNamedDataLayoutHWIO: LayoutHWIO
	MPSGraphTensorNamedDataLayoutHWIO MPSGraphTensorNamedDataLayout = 3
	// MPSGraphTensorNamedDataLayoutNCDHW: LayoutNCDHW
	MPSGraphTensorNamedDataLayoutNCDHW MPSGraphTensorNamedDataLayout = 7
	// MPSGraphTensorNamedDataLayoutNCHW: LayoutNCHW
	MPSGraphTensorNamedDataLayoutNCHW MPSGraphTensorNamedDataLayout = 0
	// MPSGraphTensorNamedDataLayoutNDHWC: LayoutNDHWC
	MPSGraphTensorNamedDataLayoutNDHWC MPSGraphTensorNamedDataLayout = 8
	// MPSGraphTensorNamedDataLayoutNHWC: LayoutNHWC
	MPSGraphTensorNamedDataLayoutNHWC MPSGraphTensorNamedDataLayout = 1
	// MPSGraphTensorNamedDataLayoutOIDHW: LayoutOIDHW
	MPSGraphTensorNamedDataLayoutOIDHW MPSGraphTensorNamedDataLayout = 9
	// MPSGraphTensorNamedDataLayoutOIHW: LayoutOIHW
	MPSGraphTensorNamedDataLayoutOIHW MPSGraphTensorNamedDataLayout = 2
)

func (e MPSGraphTensorNamedDataLayout) String() string {
	switch e {
	case MPSGraphTensorNamedDataLayoutCHW:
		return "MPSGraphTensorNamedDataLayoutCHW"
	case MPSGraphTensorNamedDataLayoutDHWIO:
		return "MPSGraphTensorNamedDataLayoutDHWIO"
	case MPSGraphTensorNamedDataLayoutHW:
		return "MPSGraphTensorNamedDataLayoutHW"
	case MPSGraphTensorNamedDataLayoutHWC:
		return "MPSGraphTensorNamedDataLayoutHWC"
	case MPSGraphTensorNamedDataLayoutHWIO:
		return "MPSGraphTensorNamedDataLayoutHWIO"
	case MPSGraphTensorNamedDataLayoutNCDHW:
		return "MPSGraphTensorNamedDataLayoutNCDHW"
	case MPSGraphTensorNamedDataLayoutNCHW:
		return "MPSGraphTensorNamedDataLayoutNCHW"
	case MPSGraphTensorNamedDataLayoutNDHWC:
		return "MPSGraphTensorNamedDataLayoutNDHWC"
	case MPSGraphTensorNamedDataLayoutNHWC:
		return "MPSGraphTensorNamedDataLayoutNHWC"
	case MPSGraphTensorNamedDataLayoutOIDHW:
		return "MPSGraphTensorNamedDataLayoutOIDHW"
	case MPSGraphTensorNamedDataLayoutOIHW:
		return "MPSGraphTensorNamedDataLayoutOIHW"
	default:
		return fmt.Sprintf("MPSGraphTensorNamedDataLayout(%d)", e)
	}
}
