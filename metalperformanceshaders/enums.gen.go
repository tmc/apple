// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureStatus
type MPSAccelerationStructureStatus uint

const (
	// Deprecated.
	MPSAccelerationStructureStatusBuilt MPSAccelerationStructureStatus = 1
	// Deprecated.
	MPSAccelerationStructureStatusUnbuilt MPSAccelerationStructureStatus = 0
)

func (e MPSAccelerationStructureStatus) String() string {
	switch e {
	case MPSAccelerationStructureStatusBuilt:
		return "MPSAccelerationStructureStatusBuilt"
	case MPSAccelerationStructureStatusUnbuilt:
		return "MPSAccelerationStructureStatusUnbuilt"
	default:
		return fmt.Sprintf("MPSAccelerationStructureStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureUsage
type MPSAccelerationStructureUsage uint

const (
	// Deprecated.
	MPSAccelerationStructureUsageFrequentRebuild MPSAccelerationStructureUsage = 2
	// Deprecated.
	MPSAccelerationStructureUsageNone MPSAccelerationStructureUsage = 0
	// Deprecated.
	MPSAccelerationStructureUsagePreferCPUBuild MPSAccelerationStructureUsage = 8
	// Deprecated.
	MPSAccelerationStructureUsagePreferGPUBuild MPSAccelerationStructureUsage = 4
	// Deprecated.
	MPSAccelerationStructureUsageRefit MPSAccelerationStructureUsage = 1
)

func (e MPSAccelerationStructureUsage) String() string {
	switch e {
	case MPSAccelerationStructureUsageFrequentRebuild:
		return "MPSAccelerationStructureUsageFrequentRebuild"
	case MPSAccelerationStructureUsageNone:
		return "MPSAccelerationStructureUsageNone"
	case MPSAccelerationStructureUsagePreferCPUBuild:
		return "MPSAccelerationStructureUsagePreferCPUBuild"
	case MPSAccelerationStructureUsagePreferGPUBuild:
		return "MPSAccelerationStructureUsagePreferGPUBuild"
	case MPSAccelerationStructureUsageRefit:
		return "MPSAccelerationStructureUsageRefit"
	default:
		return fmt.Sprintf("MPSAccelerationStructureUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAliasingStrategy
type MPSAliasingStrategy uint

const (
	MPSAliasingStrategyAliasingReserved         MPSAliasingStrategy = 3
	MPSAliasingStrategyDefault                  MPSAliasingStrategy = 0
	MPSAliasingStrategyDontCare                 MPSAliasingStrategy = 0
	MPSAliasingStrategyPreferNonTemporaryMemory MPSAliasingStrategy = 8
	MPSAliasingStrategyPreferTemporaryMemory    MPSAliasingStrategy = 4
	MPSAliasingStrategyShallAlias               MPSAliasingStrategy = 1
	MPSAliasingStrategyShallNotAlias            MPSAliasingStrategy = 2
)

func (e MPSAliasingStrategy) String() string {
	switch e {
	case MPSAliasingStrategyAliasingReserved:
		return "MPSAliasingStrategyAliasingReserved"
	case MPSAliasingStrategyDefault:
		return "MPSAliasingStrategyDefault"
	case MPSAliasingStrategyPreferNonTemporaryMemory:
		return "MPSAliasingStrategyPreferNonTemporaryMemory"
	case MPSAliasingStrategyPreferTemporaryMemory:
		return "MPSAliasingStrategyPreferTemporaryMemory"
	case MPSAliasingStrategyShallAlias:
		return "MPSAliasingStrategyShallAlias"
	case MPSAliasingStrategyShallNotAlias:
		return "MPSAliasingStrategyShallNotAlias"
	default:
		return fmt.Sprintf("MPSAliasingStrategy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType
type MPSAlphaType uint

const (
	// MPSAlphaTypeAlphaIsOne: Alpha is guaranteed to be 1.
	MPSAlphaTypeAlphaIsOne MPSAlphaType = 1
	// MPSAlphaTypeNonPremultiplied: The image is not premultiplied by alpha.
	MPSAlphaTypeNonPremultiplied MPSAlphaType = 0
	// MPSAlphaTypePremultiplied: The image is premultiplied by alpha.
	MPSAlphaTypePremultiplied MPSAlphaType = 2
)

func (e MPSAlphaType) String() string {
	switch e {
	case MPSAlphaTypeAlphaIsOne:
		return "MPSAlphaTypeAlphaIsOne"
	case MPSAlphaTypeNonPremultiplied:
		return "MPSAlphaTypeNonPremultiplied"
	case MPSAlphaTypePremultiplied:
		return "MPSAlphaTypePremultiplied"
	default:
		return fmt.Sprintf("MPSAlphaType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBoundingBoxIntersectionTestType
type MPSBoundingBoxIntersectionTestType uint

const (
	// Deprecated.
	MPSBoundingBoxIntersectionTestTypeAxisAligned MPSBoundingBoxIntersectionTestType = 1
	// Deprecated.
	MPSBoundingBoxIntersectionTestTypeDefault MPSBoundingBoxIntersectionTestType = 0
	// Deprecated.
	MPSBoundingBoxIntersectionTestTypeFast MPSBoundingBoxIntersectionTestType = 2
)

func (e MPSBoundingBoxIntersectionTestType) String() string {
	switch e {
	case MPSBoundingBoxIntersectionTestTypeAxisAligned:
		return "MPSBoundingBoxIntersectionTestTypeAxisAligned"
	case MPSBoundingBoxIntersectionTestTypeDefault:
		return "MPSBoundingBoxIntersectionTestTypeDefault"
	case MPSBoundingBoxIntersectionTestTypeFast:
		return "MPSBoundingBoxIntersectionTestTypeFast"
	default:
		return fmt.Sprintf("MPSBoundingBoxIntersectionTestType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationFlags
type MPSCNNBatchNormalizationFlags uint

const (
	MPSCNNBatchNormalizationFlagsCalculateStatisticsAlways    MPSCNNBatchNormalizationFlags = 1
	MPSCNNBatchNormalizationFlagsCalculateStatisticsAutomatic MPSCNNBatchNormalizationFlags = 0
	MPSCNNBatchNormalizationFlagsCalculateStatisticsMask      MPSCNNBatchNormalizationFlags = 3
	MPSCNNBatchNormalizationFlagsCalculateStatisticsNever     MPSCNNBatchNormalizationFlags = 2
	MPSCNNBatchNormalizationFlagsDefault                      MPSCNNBatchNormalizationFlags = 0
)

func (e MPSCNNBatchNormalizationFlags) String() string {
	switch e {
	case MPSCNNBatchNormalizationFlagsCalculateStatisticsAlways:
		return "MPSCNNBatchNormalizationFlagsCalculateStatisticsAlways"
	case MPSCNNBatchNormalizationFlagsCalculateStatisticsAutomatic:
		return "MPSCNNBatchNormalizationFlagsCalculateStatisticsAutomatic"
	case MPSCNNBatchNormalizationFlagsCalculateStatisticsMask:
		return "MPSCNNBatchNormalizationFlagsCalculateStatisticsMask"
	case MPSCNNBatchNormalizationFlagsCalculateStatisticsNever:
		return "MPSCNNBatchNormalizationFlagsCalculateStatisticsNever"
	default:
		return fmt.Sprintf("MPSCNNBatchNormalizationFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionFlags
type MPSCNNBinaryConvolutionFlags uint

const (
	MPSCNNBinaryConvolutionFlagsNone           MPSCNNBinaryConvolutionFlags = 0
	MPSCNNBinaryConvolutionFlagsUseBetaScaling MPSCNNBinaryConvolutionFlags = 1
)

func (e MPSCNNBinaryConvolutionFlags) String() string {
	switch e {
	case MPSCNNBinaryConvolutionFlagsNone:
		return "MPSCNNBinaryConvolutionFlagsNone"
	case MPSCNNBinaryConvolutionFlagsUseBetaScaling:
		return "MPSCNNBinaryConvolutionFlagsUseBetaScaling"
	default:
		return fmt.Sprintf("MPSCNNBinaryConvolutionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionType
type MPSCNNBinaryConvolutionType uint

const (
	// MPSCNNBinaryConvolutionTypeAND: A convolution type that uses input image binarization and the AND-operation.
	MPSCNNBinaryConvolutionTypeAND MPSCNNBinaryConvolutionType = 2
	// MPSCNNBinaryConvolutionTypeBinaryWeights: A convolution type that operates as a normal convolution, except that the weights are binary values.
	MPSCNNBinaryConvolutionTypeBinaryWeights MPSCNNBinaryConvolutionType = 0
	// MPSCNNBinaryConvolutionTypeXNOR: A convolution type that uses input image binarization and the XNOR-operation.
	MPSCNNBinaryConvolutionTypeXNOR MPSCNNBinaryConvolutionType = 1
)

func (e MPSCNNBinaryConvolutionType) String() string {
	switch e {
	case MPSCNNBinaryConvolutionTypeAND:
		return "MPSCNNBinaryConvolutionTypeAND"
	case MPSCNNBinaryConvolutionTypeBinaryWeights:
		return "MPSCNNBinaryConvolutionTypeBinaryWeights"
	case MPSCNNBinaryConvolutionTypeXNOR:
		return "MPSCNNBinaryConvolutionTypeXNOR"
	default:
		return fmt.Sprintf("MPSCNNBinaryConvolutionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionFlags
type MPSCNNConvolutionFlags uint

const (
	// Deprecated.
	MPSCNNConvolutionFlagsNone MPSCNNConvolutionFlags = 0
)

func (e MPSCNNConvolutionFlags) String() string {
	switch e {
	case MPSCNNConvolutionFlagsNone:
		return "MPSCNNConvolutionFlagsNone"
	default:
		return fmt.Sprintf("MPSCNNConvolutionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientOption
type MPSCNNConvolutionGradientOption uint

const (
	MPSCNNConvolutionGradientOptionAll                        MPSCNNConvolutionGradientOption = 3
	MPSCNNConvolutionGradientOptionGradientWithData           MPSCNNConvolutionGradientOption = 1
	MPSCNNConvolutionGradientOptionGradientWithWeightsAndBias MPSCNNConvolutionGradientOption = 2
)

func (e MPSCNNConvolutionGradientOption) String() string {
	switch e {
	case MPSCNNConvolutionGradientOptionAll:
		return "MPSCNNConvolutionGradientOptionAll"
	case MPSCNNConvolutionGradientOptionGradientWithData:
		return "MPSCNNConvolutionGradientOptionGradientWithData"
	case MPSCNNConvolutionGradientOptionGradientWithWeightsAndBias:
		return "MPSCNNConvolutionGradientOptionGradientWithWeightsAndBias"
	default:
		return fmt.Sprintf("MPSCNNConvolutionGradientOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsLayout
type MPSCNNConvolutionWeightsLayout uint32

const (
	MPSCNNConvolutionWeightsLayoutOHWI MPSCNNConvolutionWeightsLayout = 0
)

func (e MPSCNNConvolutionWeightsLayout) String() string {
	switch e {
	case MPSCNNConvolutionWeightsLayoutOHWI:
		return "MPSCNNConvolutionWeightsLayoutOHWI"
	default:
		return fmt.Sprintf("MPSCNNConvolutionWeightsLayout(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossType
type MPSCNNLossType uint32

const (
	MPSCNNLossTypeCategoricalCrossEntropy   MPSCNNLossType = 4
	MPSCNNLossTypeCosineDistance            MPSCNNLossType = 7
	MPSCNNLossTypeCount                     MPSCNNLossType = 10
	MPSCNNLossTypeHinge                     MPSCNNLossType = 5
	MPSCNNLossTypeHuber                     MPSCNNLossType = 6
	MPSCNNLossTypeKullbackLeiblerDivergence MPSCNNLossType = 9
	MPSCNNLossTypeLog                       MPSCNNLossType = 8
	MPSCNNLossTypeMeanAbsoluteError         MPSCNNLossType = 0
	MPSCNNLossTypeMeanSquaredError          MPSCNNLossType = 1
	MPSCNNLossTypeSigmoidCrossEntropy       MPSCNNLossType = 3
	MPSCNNLossTypeSoftMaxCrossEntropy       MPSCNNLossType = 2
)

func (e MPSCNNLossType) String() string {
	switch e {
	case MPSCNNLossTypeCategoricalCrossEntropy:
		return "MPSCNNLossTypeCategoricalCrossEntropy"
	case MPSCNNLossTypeCosineDistance:
		return "MPSCNNLossTypeCosineDistance"
	case MPSCNNLossTypeCount:
		return "MPSCNNLossTypeCount"
	case MPSCNNLossTypeHinge:
		return "MPSCNNLossTypeHinge"
	case MPSCNNLossTypeHuber:
		return "MPSCNNLossTypeHuber"
	case MPSCNNLossTypeKullbackLeiblerDivergence:
		return "MPSCNNLossTypeKullbackLeiblerDivergence"
	case MPSCNNLossTypeLog:
		return "MPSCNNLossTypeLog"
	case MPSCNNLossTypeMeanAbsoluteError:
		return "MPSCNNLossTypeMeanAbsoluteError"
	case MPSCNNLossTypeMeanSquaredError:
		return "MPSCNNLossTypeMeanSquaredError"
	case MPSCNNLossTypeSigmoidCrossEntropy:
		return "MPSCNNLossTypeSigmoidCrossEntropy"
	case MPSCNNLossTypeSoftMaxCrossEntropy:
		return "MPSCNNLossTypeSoftMaxCrossEntropy"
	default:
		return fmt.Sprintf("MPSCNNLossType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronType
type MPSCNNNeuronType int32

const (
	// MPSCNNNeuronTypeAbsolute: A neuron type indicating an absolute neuron filter.
	MPSCNNNeuronTypeAbsolute MPSCNNNeuronType = 6
	MPSCNNNeuronTypeCount    MPSCNNNeuronType = 16
	// MPSCNNNeuronTypeELU: A neuron type indicating a parametric exponential linear unit neuron filter.
	MPSCNNNeuronTypeELU         MPSCNNNeuronType = 9
	MPSCNNNeuronTypeExponential MPSCNNNeuronType = 13
	MPSCNNNeuronTypeGeLU        MPSCNNNeuronType = 15
	// MPSCNNNeuronTypeHardSigmoid: A neuron type indicating a hard sigmoid neuron filter.
	MPSCNNNeuronTypeHardSigmoid MPSCNNNeuronType = 4
	// MPSCNNNeuronTypeLinear: A neuron type indicating a linear neuron filter.
	MPSCNNNeuronTypeLinear    MPSCNNNeuronType = 2
	MPSCNNNeuronTypeLogarithm MPSCNNNeuronType = 14
	// MPSCNNNeuronTypeNone: A neuron type indicating no neuron filter.
	MPSCNNNeuronTypeNone  MPSCNNNeuronType = 0
	MPSCNNNeuronTypePReLU MPSCNNNeuronType = 10
	MPSCNNNeuronTypePower MPSCNNNeuronType = 12
	// MPSCNNNeuronTypeReLU: A neuron type indicating a rectified linear unit neuron filter.
	MPSCNNNeuronTypeReLU  MPSCNNNeuronType = 1
	MPSCNNNeuronTypeReLUN MPSCNNNeuronType = 11
	// MPSCNNNeuronTypeSigmoid: A neuron type indicating a sigmoid neuron filter.
	MPSCNNNeuronTypeSigmoid MPSCNNNeuronType = 3
	// MPSCNNNeuronTypeSoftPlus: A neuron type indicating a parametric softplus neuron filter.
	MPSCNNNeuronTypeSoftPlus MPSCNNNeuronType = 7
	// MPSCNNNeuronTypeSoftSign: A neuron type indicating a softsign neuron filter.
	MPSCNNNeuronTypeSoftSign MPSCNNNeuronType = 8
	// MPSCNNNeuronTypeTanH: A neuron type indicating a hyperbolic tangent neuron filter.
	MPSCNNNeuronTypeTanH MPSCNNNeuronType = 5
)

func (e MPSCNNNeuronType) String() string {
	switch e {
	case MPSCNNNeuronTypeAbsolute:
		return "MPSCNNNeuronTypeAbsolute"
	case MPSCNNNeuronTypeCount:
		return "MPSCNNNeuronTypeCount"
	case MPSCNNNeuronTypeELU:
		return "MPSCNNNeuronTypeELU"
	case MPSCNNNeuronTypeExponential:
		return "MPSCNNNeuronTypeExponential"
	case MPSCNNNeuronTypeGeLU:
		return "MPSCNNNeuronTypeGeLU"
	case MPSCNNNeuronTypeHardSigmoid:
		return "MPSCNNNeuronTypeHardSigmoid"
	case MPSCNNNeuronTypeLinear:
		return "MPSCNNNeuronTypeLinear"
	case MPSCNNNeuronTypeLogarithm:
		return "MPSCNNNeuronTypeLogarithm"
	case MPSCNNNeuronTypeNone:
		return "MPSCNNNeuronTypeNone"
	case MPSCNNNeuronTypePReLU:
		return "MPSCNNNeuronTypePReLU"
	case MPSCNNNeuronTypePower:
		return "MPSCNNNeuronTypePower"
	case MPSCNNNeuronTypeReLU:
		return "MPSCNNNeuronTypeReLU"
	case MPSCNNNeuronTypeReLUN:
		return "MPSCNNNeuronTypeReLUN"
	case MPSCNNNeuronTypeSigmoid:
		return "MPSCNNNeuronTypeSigmoid"
	case MPSCNNNeuronTypeSoftPlus:
		return "MPSCNNNeuronTypeSoftPlus"
	case MPSCNNNeuronTypeSoftSign:
		return "MPSCNNNeuronTypeSoftSign"
	case MPSCNNNeuronTypeTanH:
		return "MPSCNNNeuronTypeTanH"
	default:
		return fmt.Sprintf("MPSCNNNeuronType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNReductionType
type MPSCNNReductionType int32

const (
	MPSCNNReductionTypeCount               MPSCNNReductionType = 4
	MPSCNNReductionTypeMean                MPSCNNReductionType = 2
	MPSCNNReductionTypeNone                MPSCNNReductionType = 0
	MPSCNNReductionTypeSum                 MPSCNNReductionType = 1
	MPSCNNReductionTypeSumByNonZeroWeights MPSCNNReductionType = 3
)

func (e MPSCNNReductionType) String() string {
	switch e {
	case MPSCNNReductionTypeCount:
		return "MPSCNNReductionTypeCount"
	case MPSCNNReductionTypeMean:
		return "MPSCNNReductionTypeMean"
	case MPSCNNReductionTypeNone:
		return "MPSCNNReductionTypeNone"
	case MPSCNNReductionTypeSum:
		return "MPSCNNReductionTypeSum"
	case MPSCNNReductionTypeSumByNonZeroWeights:
		return "MPSCNNReductionTypeSumByNonZeroWeights"
	default:
		return fmt.Sprintf("MPSCNNReductionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNWeightsQuantizationType
type MPSCNNWeightsQuantizationType uint32

const (
	MPSCNNWeightsQuantizationTypeLinear      MPSCNNWeightsQuantizationType = 1
	MPSCNNWeightsQuantizationTypeLookupTable MPSCNNWeightsQuantizationType = 2
	MPSCNNWeightsQuantizationTypeNone        MPSCNNWeightsQuantizationType = 0
)

func (e MPSCNNWeightsQuantizationType) String() string {
	switch e {
	case MPSCNNWeightsQuantizationTypeLinear:
		return "MPSCNNWeightsQuantizationTypeLinear"
	case MPSCNNWeightsQuantizationTypeLookupTable:
		return "MPSCNNWeightsQuantizationTypeLookupTable"
	case MPSCNNWeightsQuantizationTypeNone:
		return "MPSCNNWeightsQuantizationTypeNone"
	default:
		return fmt.Sprintf("MPSCNNWeightsQuantizationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelIndex
type MPSCustomKernelIndex uint32

const (
	MPSCustomKernelIndexDestIndex     MPSCustomKernelIndex = 0
	MPSCustomKernelIndexSrc0Index     MPSCustomKernelIndex = 0
	MPSCustomKernelIndexSrc1Index     MPSCustomKernelIndex = 1
	MPSCustomKernelIndexSrc2Index     MPSCustomKernelIndex = 2
	MPSCustomKernelIndexSrc3Index     MPSCustomKernelIndex = 3
	MPSCustomKernelIndexSrc4Index     MPSCustomKernelIndex = 4
	MPSCustomKernelIndexUserDataIndex MPSCustomKernelIndex = 30
)

func (e MPSCustomKernelIndex) String() string {
	switch e {
	case MPSCustomKernelIndexDestIndex:
		return "MPSCustomKernelIndexDestIndex"
	case MPSCustomKernelIndexSrc1Index:
		return "MPSCustomKernelIndexSrc1Index"
	case MPSCustomKernelIndexSrc2Index:
		return "MPSCustomKernelIndexSrc2Index"
	case MPSCustomKernelIndexSrc3Index:
		return "MPSCustomKernelIndexSrc3Index"
	case MPSCustomKernelIndexSrc4Index:
		return "MPSCustomKernelIndexSrc4Index"
	case MPSCustomKernelIndexUserDataIndex:
		return "MPSCustomKernelIndexUserDataIndex"
	default:
		return fmt.Sprintf("MPSCustomKernelIndex(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataLayout
type MPSDataLayout uint

const (
	MPSDataLayoutFeatureChannelsxHeightxWidth MPSDataLayout = 1
	MPSDataLayoutHeightxWidthxFeatureChannels MPSDataLayout = 0
)

func (e MPSDataLayout) String() string {
	switch e {
	case MPSDataLayoutFeatureChannelsxHeightxWidth:
		return "MPSDataLayoutFeatureChannelsxHeightxWidth"
	case MPSDataLayoutHeightxWidthxFeatureChannels:
		return "MPSDataLayoutHeightxWidthxFeatureChannels"
	default:
		return fmt.Sprintf("MPSDataLayout(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDataType
type MPSDataType uint32

const (
	MPSDataTypeAlternateEncodingBit MPSDataType = 0x80000000
	MPSDataTypeBFloat16             MPSDataType = 2415919120
	MPSDataTypeBool                 MPSDataType = 2147483656
	MPSDataTypeComplexBit           MPSDataType = 0x1000000
	MPSDataTypeComplexFloat16       MPSDataType = 285212704
	MPSDataTypeComplexFloat32       MPSDataType = 285212736
	MPSDataTypeFloat16              MPSDataType = 268435472
	// MPSDataTypeFloat32: A 32-bit floating point type (single precision).
	MPSDataTypeFloat32 MPSDataType = 268435488
	// MPSDataTypeFloatBit: A common bit for all floating point data types.
	MPSDataTypeFloatBit      MPSDataType = 0x10000000
	MPSDataTypeInt16         MPSDataType = 536870928
	MPSDataTypeInt2          MPSDataType = 536870914
	MPSDataTypeInt32         MPSDataType = 536870944
	MPSDataTypeInt4          MPSDataType = 536870916
	MPSDataTypeInt64         MPSDataType = 536870976
	MPSDataTypeInt8          MPSDataType = 536870920
	MPSDataTypeIntBit        MPSDataType = 536870912
	MPSDataTypeInvalid       MPSDataType = 0
	MPSDataTypeNormalizedBit MPSDataType = 0x40000000
	MPSDataTypeSignedBit     MPSDataType = 0x20000000
	MPSDataTypeUInt16        MPSDataType = 16
	MPSDataTypeUInt2         MPSDataType = 2
	MPSDataTypeUInt32        MPSDataType = 32
	MPSDataTypeUInt4         MPSDataType = 4
	MPSDataTypeUInt64        MPSDataType = 64
	MPSDataTypeUInt8         MPSDataType = 8
	MPSDataTypeUnorm1        MPSDataType = 1073741825
	MPSDataTypeUnorm8        MPSDataType = 1073741832
)

func (e MPSDataType) String() string {
	switch e {
	case MPSDataTypeAlternateEncodingBit:
		return "MPSDataTypeAlternateEncodingBit"
	case MPSDataTypeBFloat16:
		return "MPSDataTypeBFloat16"
	case MPSDataTypeBool:
		return "MPSDataTypeBool"
	case MPSDataTypeComplexBit:
		return "MPSDataTypeComplexBit"
	case MPSDataTypeComplexFloat16:
		return "MPSDataTypeComplexFloat16"
	case MPSDataTypeComplexFloat32:
		return "MPSDataTypeComplexFloat32"
	case MPSDataTypeFloat16:
		return "MPSDataTypeFloat16"
	case MPSDataTypeFloat32:
		return "MPSDataTypeFloat32"
	case MPSDataTypeFloatBit:
		return "MPSDataTypeFloatBit"
	case MPSDataTypeInt16:
		return "MPSDataTypeInt16"
	case MPSDataTypeInt2:
		return "MPSDataTypeInt2"
	case MPSDataTypeInt32:
		return "MPSDataTypeInt32"
	case MPSDataTypeInt4:
		return "MPSDataTypeInt4"
	case MPSDataTypeInt64:
		return "MPSDataTypeInt64"
	case MPSDataTypeInt8:
		return "MPSDataTypeInt8"
	case MPSDataTypeIntBit:
		return "MPSDataTypeIntBit"
	case MPSDataTypeInvalid:
		return "MPSDataTypeInvalid"
	case MPSDataTypeNormalizedBit:
		return "MPSDataTypeNormalizedBit"
	case MPSDataTypeUInt16:
		return "MPSDataTypeUInt16"
	case MPSDataTypeUInt2:
		return "MPSDataTypeUInt2"
	case MPSDataTypeUInt32:
		return "MPSDataTypeUInt32"
	case MPSDataTypeUInt4:
		return "MPSDataTypeUInt4"
	case MPSDataTypeUInt64:
		return "MPSDataTypeUInt64"
	case MPSDataTypeUInt8:
		return "MPSDataTypeUInt8"
	case MPSDataTypeUnorm1:
		return "MPSDataTypeUnorm1"
	case MPSDataTypeUnorm8:
		return "MPSDataTypeUnorm8"
	default:
		return fmt.Sprintf("MPSDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceCapsValues
type MPSDeviceCapsValues uint32

const (
	MPSDeviceCapsLast                        MPSDeviceCapsValues = 8192
	MPSDeviceCapsNull                        MPSDeviceCapsValues = 0
	MPSDeviceIsAppleDevice                   MPSDeviceCapsValues = 1024
	MPSDeviceSupportsBFloat16Arithmetic      MPSDeviceCapsValues = 4096
	MPSDeviceSupportsFloat16BicubicFiltering MPSDeviceCapsValues = 512
	MPSDeviceSupportsFloat32Filtering        MPSDeviceCapsValues = 128
	MPSDeviceSupportsNorm16BicubicFiltering  MPSDeviceCapsValues = 256
	MPSDeviceSupportsQuadShuffle             MPSDeviceCapsValues = 16
	MPSDeviceSupportsReadWriteTextures       MPSDeviceCapsValues = 4
	MPSDeviceSupportsReadableArrayOfTextures MPSDeviceCapsValues = 1
	MPSDeviceSupportsSimdReduction           MPSDeviceCapsValues = 64
	MPSDeviceSupportsSimdShuffle             MPSDeviceCapsValues = 32
	MPSDeviceSupportsSimdShuffleAndFill      MPSDeviceCapsValues = 2048
	MPSDeviceSupportsSimdgroupBarrier        MPSDeviceCapsValues = 8
	MPSDeviceSupportsWritableArrayOfTextures MPSDeviceCapsValues = 2
)

func (e MPSDeviceCapsValues) String() string {
	switch e {
	case MPSDeviceCapsLast:
		return "MPSDeviceCapsLast"
	case MPSDeviceCapsNull:
		return "MPSDeviceCapsNull"
	case MPSDeviceIsAppleDevice:
		return "MPSDeviceIsAppleDevice"
	case MPSDeviceSupportsBFloat16Arithmetic:
		return "MPSDeviceSupportsBFloat16Arithmetic"
	case MPSDeviceSupportsFloat16BicubicFiltering:
		return "MPSDeviceSupportsFloat16BicubicFiltering"
	case MPSDeviceSupportsFloat32Filtering:
		return "MPSDeviceSupportsFloat32Filtering"
	case MPSDeviceSupportsNorm16BicubicFiltering:
		return "MPSDeviceSupportsNorm16BicubicFiltering"
	case MPSDeviceSupportsQuadShuffle:
		return "MPSDeviceSupportsQuadShuffle"
	case MPSDeviceSupportsReadWriteTextures:
		return "MPSDeviceSupportsReadWriteTextures"
	case MPSDeviceSupportsReadableArrayOfTextures:
		return "MPSDeviceSupportsReadableArrayOfTextures"
	case MPSDeviceSupportsSimdReduction:
		return "MPSDeviceSupportsSimdReduction"
	case MPSDeviceSupportsSimdShuffle:
		return "MPSDeviceSupportsSimdShuffle"
	case MPSDeviceSupportsSimdShuffleAndFill:
		return "MPSDeviceSupportsSimdShuffleAndFill"
	case MPSDeviceSupportsSimdgroupBarrier:
		return "MPSDeviceSupportsSimdgroupBarrier"
	case MPSDeviceSupportsWritableArrayOfTextures:
		return "MPSDeviceSupportsWritableArrayOfTextures"
	default:
		return fmt.Sprintf("MPSDeviceCapsValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceOptions
type MPSDeviceOptions uint

const (
	MPSDeviceOptionsDefault       MPSDeviceOptions = 0
	MPSDeviceOptionsLowPower      MPSDeviceOptions = 1
	MPSDeviceOptionsSkipRemovable MPSDeviceOptions = 2
)

func (e MPSDeviceOptions) String() string {
	switch e {
	case MPSDeviceOptionsDefault:
		return "MPSDeviceOptionsDefault"
	case MPSDeviceOptionsLowPower:
		return "MPSDeviceOptionsLowPower"
	case MPSDeviceOptionsSkipRemovable:
		return "MPSDeviceOptionsSkipRemovable"
	default:
		return fmt.Sprintf("MPSDeviceOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversionOptions
type MPSFColorConversionOptions uint64

const (
	MPSFColorConversionOptionsPrecisionDefault MPSFColorConversionOptions = 0
)

func (e MPSFColorConversionOptions) String() string {
	switch e {
	case MPSFColorConversionOptionsPrecisionDefault:
		return "MPSFColorConversionOptionsPrecisionDefault"
	default:
		return fmt.Sprintf("MPSFColorConversionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeBit
type MPSFloatDataTypeBit uint32

const (
	MPSFloatDataTypeExponentBit MPSFloatDataTypeBit = 0x7c0000
	MPSFloatDataTypeMantissaBit MPSFloatDataTypeBit = 0x3fc00
	MPSFloatDataTypeSignBit     MPSFloatDataTypeBit = 0x800000
)

func (e MPSFloatDataTypeBit) String() string {
	switch e {
	case MPSFloatDataTypeExponentBit:
		return "MPSFloatDataTypeExponentBit"
	case MPSFloatDataTypeMantissaBit:
		return "MPSFloatDataTypeMantissaBit"
	case MPSFloatDataTypeSignBit:
		return "MPSFloatDataTypeSignBit"
	default:
		return fmt.Sprintf("MPSFloatDataTypeBit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFloatDataTypeShift
type MPSFloatDataTypeShift uint32

const (
	MPSFloatDataTypeExponentShift MPSFloatDataTypeShift = 18
	MPSFloatDataTypeMantissaShift MPSFloatDataTypeShift = 10
	MPSFloatDataTypeSignShift     MPSFloatDataTypeShift = 23
)

func (e MPSFloatDataTypeShift) String() string {
	switch e {
	case MPSFloatDataTypeExponentShift:
		return "MPSFloatDataTypeExponentShift"
	case MPSFloatDataTypeMantissaShift:
		return "MPSFloatDataTypeMantissaShift"
	case MPSFloatDataTypeSignShift:
		return "MPSFloatDataTypeSignShift"
	default:
		return fmt.Sprintf("MPSFloatDataTypeShift(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode
type MPSImageEdgeMode uint

const (
	// MPSImageEdgeModeClamp: Out-of-bound pixels are clamped to the nearest edge pixel.
	MPSImageEdgeModeClamp          MPSImageEdgeMode = 1
	MPSImageEdgeModeConstant       MPSImageEdgeMode = 4
	MPSImageEdgeModeMirror         MPSImageEdgeMode = 2
	MPSImageEdgeModeMirrorWithEdge MPSImageEdgeMode = 3
	// MPSImageEdgeModeZero: Out-of-bound pixels are set to `(0.0, 0.0, 0.0, 1.0)` for images without an alpha channel or `(0.0, 0.0, 0.0, 0.0)` for images with an alpha channel, as defined by their pixel format.
	MPSImageEdgeModeZero MPSImageEdgeMode = 0
)

func (e MPSImageEdgeMode) String() string {
	switch e {
	case MPSImageEdgeModeClamp:
		return "MPSImageEdgeModeClamp"
	case MPSImageEdgeModeConstant:
		return "MPSImageEdgeModeConstant"
	case MPSImageEdgeModeMirror:
		return "MPSImageEdgeModeMirror"
	case MPSImageEdgeModeMirrorWithEdge:
		return "MPSImageEdgeModeMirrorWithEdge"
	case MPSImageEdgeModeZero:
		return "MPSImageEdgeModeZero"
	default:
		return fmt.Sprintf("MPSImageEdgeMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFeatureChannelFormat
type MPSImageFeatureChannelFormat uint

const (
	MPSImageFeatureChannelFormatCount MPSImageFeatureChannelFormat = 6
	// MPSImageFeatureChannelFormatFloat16: IEEE-754 16-bit floating-point type (half precision).
	MPSImageFeatureChannelFormatFloat16 MPSImageFeatureChannelFormat = 3
	// MPSImageFeatureChannelFormatFloat32: IEEE-754 32-bit floating-point type (single precision, standard `float` type in C).
	MPSImageFeatureChannelFormatFloat32 MPSImageFeatureChannelFormat = 4
	MPSImageFeatureChannelFormatNone    MPSImageFeatureChannelFormat = 0
	// MPSImageFeatureChannelFormatUnorm16: `uint16_t` type with value `[0,65535]` and encoding `[0,1.0]`.
	MPSImageFeatureChannelFormatUnorm16 MPSImageFeatureChannelFormat = 2
	// MPSImageFeatureChannelFormatUnorm8: `uint8_t` type with value `[0,255]` and encoding `[0,1.0]`.
	MPSImageFeatureChannelFormatUnorm8     MPSImageFeatureChannelFormat = 1
	MPSImageFeatureChannelFormat_reserved0 MPSImageFeatureChannelFormat = 5
)

func (e MPSImageFeatureChannelFormat) String() string {
	switch e {
	case MPSImageFeatureChannelFormatCount:
		return "MPSImageFeatureChannelFormatCount"
	case MPSImageFeatureChannelFormatFloat16:
		return "MPSImageFeatureChannelFormatFloat16"
	case MPSImageFeatureChannelFormatFloat32:
		return "MPSImageFeatureChannelFormatFloat32"
	case MPSImageFeatureChannelFormatNone:
		return "MPSImageFeatureChannelFormatNone"
	case MPSImageFeatureChannelFormatUnorm16:
		return "MPSImageFeatureChannelFormatUnorm16"
	case MPSImageFeatureChannelFormatUnorm8:
		return "MPSImageFeatureChannelFormatUnorm8"
	case MPSImageFeatureChannelFormat_reserved0:
		return "MPSImageFeatureChannelFormat_reserved0"
	default:
		return fmt.Sprintf("MPSImageFeatureChannelFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageType
type MPSImageType uint32

const (
	MPSImageType2d                    MPSImageType = 0
	MPSImageType2d_array              MPSImageType = 1
	MPSImageType2d_array_noAlpha      MPSImageType = 5
	MPSImageType2d_noAlpha            MPSImageType = 4
	MPSImageTypeArray2d               MPSImageType = 2
	MPSImageTypeArray2d_array         MPSImageType = 3
	MPSImageTypeArray2d_array_noAlpha MPSImageType = 7
	MPSImageTypeArray2d_noAlpha       MPSImageType = 6
	MPSImageType_ArrayMask            MPSImageType = 1
	MPSImageType_BatchMask            MPSImageType = 2
	MPSImageType_bitCount             MPSImageType = 6
	MPSImageType_mask                 MPSImageType = 63
	MPSImageType_noAlpha              MPSImageType = 4
	MPSImageType_texelFormatBFloat16  MPSImageType = 24
	MPSImageType_texelFormatFloat16   MPSImageType = 16
	MPSImageType_texelFormatMask      MPSImageType = 0x38
	MPSImageType_texelFormatShift     MPSImageType = 3
	MPSImageType_texelFormatStandard  MPSImageType = 0
	MPSImageType_texelFormatUnorm8    MPSImageType = 8
	MPSImageType_typeMask             MPSImageType = 3
)

func (e MPSImageType) String() string {
	switch e {
	case MPSImageType2d:
		return "MPSImageType2d"
	case MPSImageType2d_array:
		return "MPSImageType2d_array"
	case MPSImageType2d_array_noAlpha:
		return "MPSImageType2d_array_noAlpha"
	case MPSImageType2d_noAlpha:
		return "MPSImageType2d_noAlpha"
	case MPSImageTypeArray2d:
		return "MPSImageTypeArray2d"
	case MPSImageTypeArray2d_array:
		return "MPSImageTypeArray2d_array"
	case MPSImageTypeArray2d_array_noAlpha:
		return "MPSImageTypeArray2d_array_noAlpha"
	case MPSImageTypeArray2d_noAlpha:
		return "MPSImageTypeArray2d_noAlpha"
	case MPSImageType_mask:
		return "MPSImageType_mask"
	case MPSImageType_texelFormatBFloat16:
		return "MPSImageType_texelFormatBFloat16"
	case MPSImageType_texelFormatFloat16:
		return "MPSImageType_texelFormatFloat16"
	case MPSImageType_texelFormatMask:
		return "MPSImageType_texelFormatMask"
	case MPSImageType_texelFormatUnorm8:
		return "MPSImageType_texelFormatUnorm8"
	default:
		return fmt.Sprintf("MPSImageType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDataType
type MPSIntersectionDataType uint

const (
	MPSIntersectionDataTypeDistance                                                  MPSIntersectionDataType = 0
	MPSIntersectionDataTypeDistancePrimitiveIndex                                    MPSIntersectionDataType = 1
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndex                         MPSIntersectionDataType = 5
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexCoordinates              MPSIntersectionDataType = 6
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndex            MPSIntersectionDataType = 7
	MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates MPSIntersectionDataType = 8
	MPSIntersectionDataTypeDistancePrimitiveIndexCoordinates                         MPSIntersectionDataType = 2
	MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndex                       MPSIntersectionDataType = 3
	MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndexCoordinates            MPSIntersectionDataType = 4
)

func (e MPSIntersectionDataType) String() string {
	switch e {
	case MPSIntersectionDataTypeDistance:
		return "MPSIntersectionDataTypeDistance"
	case MPSIntersectionDataTypeDistancePrimitiveIndex:
		return "MPSIntersectionDataTypeDistancePrimitiveIndex"
	case MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndex:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndex"
	case MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexCoordinates:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexCoordinates"
	case MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndex:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndex"
	case MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates"
	case MPSIntersectionDataTypeDistancePrimitiveIndexCoordinates:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexCoordinates"
	case MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndex:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndex"
	case MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndexCoordinates:
		return "MPSIntersectionDataTypeDistancePrimitiveIndexInstanceIndexCoordinates"
	default:
		return fmt.Sprintf("MPSIntersectionDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionType
type MPSIntersectionType uint

const (
	MPSIntersectionTypeAny     MPSIntersectionType = 1
	MPSIntersectionTypeNearest MPSIntersectionType = 0
)

func (e MPSIntersectionType) String() string {
	switch e {
	case MPSIntersectionTypeAny:
		return "MPSIntersectionTypeAny"
	case MPSIntersectionTypeNearest:
		return "MPSIntersectionTypeNearest"
	default:
		return fmt.Sprintf("MPSIntersectionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernelOptions
type MPSKernelOptions uint

const (
	// MPSKernelOptionsAllowReducedPrecision: When possible, kernels use a higher-precision data representation internally than the destination storage format to avoid excessive accumulation of computational rounding error in the result.
	MPSKernelOptionsAllowReducedPrecision MPSKernelOptions = 2
	// MPSKernelOptionsDisableInternalTiling: Some kernels may automatically split up their work internally into multiple tiles.
	MPSKernelOptionsDisableInternalTiling MPSKernelOptions = 4
	// MPSKernelOptionsInsertDebugGroups: Enables calling kernel encode methods.
	MPSKernelOptionsInsertDebugGroups MPSKernelOptions = 8
	// MPSKernelOptionsNone: The default option for the kernel.
	MPSKernelOptionsNone MPSKernelOptions = 0
	// MPSKernelOptionsSkipAPIValidation: A property that directs the kernel to perform or skip argument validation.
	MPSKernelOptionsSkipAPIValidation MPSKernelOptions = 1
	MPSKernelOptionsVerbose           MPSKernelOptions = 16
)

func (e MPSKernelOptions) String() string {
	switch e {
	case MPSKernelOptionsAllowReducedPrecision:
		return "MPSKernelOptionsAllowReducedPrecision"
	case MPSKernelOptionsDisableInternalTiling:
		return "MPSKernelOptionsDisableInternalTiling"
	case MPSKernelOptionsInsertDebugGroups:
		return "MPSKernelOptionsInsertDebugGroups"
	case MPSKernelOptionsNone:
		return "MPSKernelOptionsNone"
	case MPSKernelOptionsSkipAPIValidation:
		return "MPSKernelOptionsSkipAPIValidation"
	case MPSKernelOptionsVerbose:
		return "MPSKernelOptionsVerbose"
	default:
		return fmt.Sprintf("MPSKernelOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionStatus
type MPSMatrixDecompositionStatus int32

const (
	// MPSMatrixDecompositionStatusFailure: A status indicating the decomposition was not able to be completed.
	MPSMatrixDecompositionStatusFailure MPSMatrixDecompositionStatus = -1
	// MPSMatrixDecompositionStatusNonPositiveDefinite: A status indicating a non-positive-definite pivot value was calculated.
	MPSMatrixDecompositionStatusNonPositiveDefinite MPSMatrixDecompositionStatus = -3
	// MPSMatrixDecompositionStatusSingular: A status indicating the resulting decomposition is not suitable for use in a subsequent system solve.
	MPSMatrixDecompositionStatusSingular MPSMatrixDecompositionStatus = -2
	// MPSMatrixDecompositionStatusSuccess: A status indicating the decomposition was performed successfully.
	MPSMatrixDecompositionStatusSuccess MPSMatrixDecompositionStatus = 0
)

func (e MPSMatrixDecompositionStatus) String() string {
	switch e {
	case MPSMatrixDecompositionStatusFailure:
		return "MPSMatrixDecompositionStatusFailure"
	case MPSMatrixDecompositionStatusNonPositiveDefinite:
		return "MPSMatrixDecompositionStatusNonPositiveDefinite"
	case MPSMatrixDecompositionStatusSingular:
		return "MPSMatrixDecompositionStatusSingular"
	case MPSMatrixDecompositionStatusSuccess:
		return "MPSMatrixDecompositionStatusSuccess"
	default:
		return fmt.Sprintf("MPSMatrixDecompositionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistribution
type MPSMatrixRandomDistribution uint

const (
	MPSMatrixRandomDistributionDefault MPSMatrixRandomDistribution = 1
	MPSMatrixRandomDistributionNormal  MPSMatrixRandomDistribution = 3
	MPSMatrixRandomDistributionUniform MPSMatrixRandomDistribution = 2
)

func (e MPSMatrixRandomDistribution) String() string {
	switch e {
	case MPSMatrixRandomDistributionDefault:
		return "MPSMatrixRandomDistributionDefault"
	case MPSMatrixRandomDistributionNormal:
		return "MPSMatrixRandomDistributionNormal"
	case MPSMatrixRandomDistributionUniform:
		return "MPSMatrixRandomDistributionUniform"
	default:
		return fmt.Sprintf("MPSMatrixRandomDistribution(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationScheme
type MPSNDArrayQuantizationScheme uint

const (
	MPSNDArrayQuantizationTypeAffine MPSNDArrayQuantizationScheme = 1
	MPSNDArrayQuantizationTypeLUT    MPSNDArrayQuantizationScheme = 2
	MPSNDArrayQuantizationTypeNone   MPSNDArrayQuantizationScheme = 0
)

func (e MPSNDArrayQuantizationScheme) String() string {
	switch e {
	case MPSNDArrayQuantizationTypeAffine:
		return "MPSNDArrayQuantizationTypeAffine"
	case MPSNDArrayQuantizationTypeLUT:
		return "MPSNDArrayQuantizationTypeLUT"
	case MPSNDArrayQuantizationTypeNone:
		return "MPSNDArrayQuantizationTypeNone"
	default:
		return fmt.Sprintf("MPSNDArrayQuantizationScheme(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType
type MPSNNComparisonType uint

const (
	MPSNNComparisonTypeEqual          MPSNNComparisonType = 0
	MPSNNComparisonTypeGreater        MPSNNComparisonType = 4
	MPSNNComparisonTypeGreaterOrEqual MPSNNComparisonType = 5
	MPSNNComparisonTypeLess           MPSNNComparisonType = 2
	MPSNNComparisonTypeLessOrEqual    MPSNNComparisonType = 3
	MPSNNComparisonTypeNotEqual       MPSNNComparisonType = 1
)

func (e MPSNNComparisonType) String() string {
	switch e {
	case MPSNNComparisonTypeEqual:
		return "MPSNNComparisonTypeEqual"
	case MPSNNComparisonTypeGreater:
		return "MPSNNComparisonTypeGreater"
	case MPSNNComparisonTypeGreaterOrEqual:
		return "MPSNNComparisonTypeGreaterOrEqual"
	case MPSNNComparisonTypeLess:
		return "MPSNNComparisonTypeLess"
	case MPSNNComparisonTypeLessOrEqual:
		return "MPSNNComparisonTypeLessOrEqual"
	case MPSNNComparisonTypeNotEqual:
		return "MPSNNComparisonTypeNotEqual"
	default:
		return fmt.Sprintf("MPSNNComparisonType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConvolutionAccumulatorPrecisionOption
type MPSNNConvolutionAccumulatorPrecisionOption uint

const (
	MPSNNConvolutionAccumulatorPrecisionOptionFloat MPSNNConvolutionAccumulatorPrecisionOption = 1
	MPSNNConvolutionAccumulatorPrecisionOptionHalf  MPSNNConvolutionAccumulatorPrecisionOption = 0
)

func (e MPSNNConvolutionAccumulatorPrecisionOption) String() string {
	switch e {
	case MPSNNConvolutionAccumulatorPrecisionOptionFloat:
		return "MPSNNConvolutionAccumulatorPrecisionOptionFloat"
	case MPSNNConvolutionAccumulatorPrecisionOptionHalf:
		return "MPSNNConvolutionAccumulatorPrecisionOptionHalf"
	default:
		return fmt.Sprintf("MPSNNConvolutionAccumulatorPrecisionOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPaddingMethod
type MPSNNPaddingMethod uint

const (
	MPSNNPaddingMethodAddRemainderToBottomLeft  MPSNNPaddingMethod = 8
	MPSNNPaddingMethodAddRemainderToBottomRight MPSNNPaddingMethod = 12
	MPSNNPaddingMethodAddRemainderToMask        MPSNNPaddingMethod = 12
	// MPSNNPaddingMethodAddRemainderToTopLeft: A padding method where leftover padding is added to the top or left side of image as appropriate.
	MPSNNPaddingMethodAddRemainderToTopLeft    MPSNNPaddingMethod = 0
	MPSNNPaddingMethodAddRemainderToTopRight   MPSNNPaddingMethod = 4
	MPSNNPaddingMethodAlignBottomRight         MPSNNPaddingMethod = 2
	MPSNNPaddingMethodAlignCentered            MPSNNPaddingMethod = 0
	MPSNNPaddingMethodAlignMask                MPSNNPaddingMethod = 3
	MPSNNPaddingMethodAlignTopLeft             MPSNNPaddingMethod = 1
	MPSNNPaddingMethodAlign_reserved           MPSNNPaddingMethod = 3
	MPSNNPaddingMethodCustom                   MPSNNPaddingMethod = 16384
	MPSNNPaddingMethodCustomAllowForNodeFusion MPSNNPaddingMethod = 8192
	MPSNNPaddingMethodExcludeEdges             MPSNNPaddingMethod = 32768
	MPSNNPaddingMethodSizeFull                 MPSNNPaddingMethod = 32
	MPSNNPaddingMethodSizeMask                 MPSNNPaddingMethod = 0x7f0
	MPSNNPaddingMethodSizeSame                 MPSNNPaddingMethod = 16
	// MPSNNPaddingMethodSizeValidOnly: A padding method where result values are only produced for the area that is guaranteed to have all of its input values defined
	MPSNNPaddingMethodSizeValidOnly MPSNNPaddingMethod = 0
	MPSNNPaddingMethodSize_reserved MPSNNPaddingMethod = 48
	// Deprecated.
	MPSNNPaddingMethodCustomWhitelistForNodeFusion MPSNNPaddingMethod = 8192
)

func (e MPSNNPaddingMethod) String() string {
	switch e {
	case MPSNNPaddingMethodAddRemainderToBottomLeft:
		return "MPSNNPaddingMethodAddRemainderToBottomLeft"
	case MPSNNPaddingMethodAddRemainderToBottomRight:
		return "MPSNNPaddingMethodAddRemainderToBottomRight"
	case MPSNNPaddingMethodAddRemainderToTopLeft:
		return "MPSNNPaddingMethodAddRemainderToTopLeft"
	case MPSNNPaddingMethodAddRemainderToTopRight:
		return "MPSNNPaddingMethodAddRemainderToTopRight"
	case MPSNNPaddingMethodAlignBottomRight:
		return "MPSNNPaddingMethodAlignBottomRight"
	case MPSNNPaddingMethodAlignMask:
		return "MPSNNPaddingMethodAlignMask"
	case MPSNNPaddingMethodAlignTopLeft:
		return "MPSNNPaddingMethodAlignTopLeft"
	case MPSNNPaddingMethodCustom:
		return "MPSNNPaddingMethodCustom"
	case MPSNNPaddingMethodCustomAllowForNodeFusion:
		return "MPSNNPaddingMethodCustomAllowForNodeFusion"
	case MPSNNPaddingMethodExcludeEdges:
		return "MPSNNPaddingMethodExcludeEdges"
	case MPSNNPaddingMethodSizeFull:
		return "MPSNNPaddingMethodSizeFull"
	case MPSNNPaddingMethodSizeMask:
		return "MPSNNPaddingMethodSizeMask"
	case MPSNNPaddingMethodSizeSame:
		return "MPSNNPaddingMethodSizeSame"
	case MPSNNPaddingMethodSize_reserved:
		return "MPSNNPaddingMethodSize_reserved"
	default:
		return fmt.Sprintf("MPSNNPaddingMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNRegularizationType
type MPSNNRegularizationType uint

const (
	MPSNNRegularizationTypeL1   MPSNNRegularizationType = 1
	MPSNNRegularizationTypeL2   MPSNNRegularizationType = 2
	MPSNNRegularizationTypeNone MPSNNRegularizationType = 0
)

func (e MPSNNRegularizationType) String() string {
	switch e {
	case MPSNNRegularizationTypeL1:
		return "MPSNNRegularizationTypeL1"
	case MPSNNRegularizationTypeL2:
		return "MPSNNRegularizationTypeL2"
	case MPSNNRegularizationTypeNone:
		return "MPSNNRegularizationTypeNone"
	default:
		return fmt.Sprintf("MPSNNRegularizationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainingStyle
type MPSNNTrainingStyle uint

const (
	MPSNNTrainingStyleUpdateDeviceCPU  MPSNNTrainingStyle = 1
	MPSNNTrainingStyleUpdateDeviceGPU  MPSNNTrainingStyle = 2
	MPSNNTrainingStyleUpdateDeviceNone MPSNNTrainingStyle = 0
)

func (e MPSNNTrainingStyle) String() string {
	switch e {
	case MPSNNTrainingStyleUpdateDeviceCPU:
		return "MPSNNTrainingStyleUpdateDeviceCPU"
	case MPSNNTrainingStyleUpdateDeviceGPU:
		return "MPSNNTrainingStyleUpdateDeviceGPU"
	case MPSNNTrainingStyleUpdateDeviceNone:
		return "MPSNNTrainingStyleUpdateDeviceNone"
	default:
		return fmt.Sprintf("MPSNNTrainingStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonType
type MPSPolygonType uint

const (
	// Deprecated.
	MPSPolygonTypeQuadrilateral MPSPolygonType = 1
	// Deprecated.
	MPSPolygonTypeTriangle MPSPolygonType = 0
)

func (e MPSPolygonType) String() string {
	switch e {
	case MPSPolygonTypeQuadrilateral:
		return "MPSPolygonTypeQuadrilateral"
	case MPSPolygonTypeTriangle:
		return "MPSPolygonTypeTriangle"
	default:
		return fmt.Sprintf("MPSPolygonType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPurgeableState
type MPSPurgeableState uint

const (
	// MPSPurgeableStateAllocationDeferred: The image’s underlying texture hasn’t been allocated yet.
	MPSPurgeableStateAllocationDeferred MPSPurgeableState = 0
	// MPSPurgeableStateEmpty: The contents of the resource are or will be discarded.
	MPSPurgeableStateEmpty MPSPurgeableState = 4
	// MPSPurgeableStateKeepCurrent: The current state is queried but doesn’t change.
	MPSPurgeableStateKeepCurrent MPSPurgeableState = 1
	// MPSPurgeableStateNonVolatile: The contents of the resource aren’t allowed to be discarded.
	MPSPurgeableStateNonVolatile MPSPurgeableState = 2
	// MPSPurgeableStateVolatile: The system is allowed to discard the resource to free up memory.
	MPSPurgeableStateVolatile MPSPurgeableState = 3
)

func (e MPSPurgeableState) String() string {
	switch e {
	case MPSPurgeableStateAllocationDeferred:
		return "MPSPurgeableStateAllocationDeferred"
	case MPSPurgeableStateEmpty:
		return "MPSPurgeableStateEmpty"
	case MPSPurgeableStateKeepCurrent:
		return "MPSPurgeableStateKeepCurrent"
	case MPSPurgeableStateNonVolatile:
		return "MPSPurgeableStateNonVolatile"
	case MPSPurgeableStateVolatile:
		return "MPSPurgeableStateVolatile"
	default:
		return fmt.Sprintf("MPSPurgeableState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNBidirectionalCombineMode
type MPSRNNBidirectionalCombineMode uint

const (
	// MPSRNNBidirectionalCombineModeAdd: A mode in which two sequences are summed to form a single output.
	MPSRNNBidirectionalCombineModeAdd MPSRNNBidirectionalCombineMode = 1
	// MPSRNNBidirectionalCombineModeConcatenate: A mode in which two sequences are concatenated along the feature channels to form a single output.
	MPSRNNBidirectionalCombineModeConcatenate MPSRNNBidirectionalCombineMode = 2
	// MPSRNNBidirectionalCombineModeNone: A mode in which two sequences are kept separate.
	MPSRNNBidirectionalCombineModeNone MPSRNNBidirectionalCombineMode = 0
)

func (e MPSRNNBidirectionalCombineMode) String() string {
	switch e {
	case MPSRNNBidirectionalCombineModeAdd:
		return "MPSRNNBidirectionalCombineModeAdd"
	case MPSRNNBidirectionalCombineModeConcatenate:
		return "MPSRNNBidirectionalCombineModeConcatenate"
	case MPSRNNBidirectionalCombineModeNone:
		return "MPSRNNBidirectionalCombineModeNone"
	default:
		return fmt.Sprintf("MPSRNNBidirectionalCombineMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixId
type MPSRNNMatrixId uint

const (
	MPSRNNMatrixIdGRUInputGateBiasTerms            MPSRNNMatrixId = 21
	MPSRNNMatrixIdGRUInputGateInputWeights         MPSRNNMatrixId = 19
	MPSRNNMatrixIdGRUInputGateRecurrentWeights     MPSRNNMatrixId = 20
	MPSRNNMatrixIdGRUOutputGateBiasTerms           MPSRNNMatrixId = 28
	MPSRNNMatrixIdGRUOutputGateInputGateWeights    MPSRNNMatrixId = 27
	MPSRNNMatrixIdGRUOutputGateInputWeights        MPSRNNMatrixId = 25
	MPSRNNMatrixIdGRUOutputGateRecurrentWeights    MPSRNNMatrixId = 26
	MPSRNNMatrixIdGRURecurrentGateBiasTerms        MPSRNNMatrixId = 24
	MPSRNNMatrixIdGRURecurrentGateInputWeights     MPSRNNMatrixId = 22
	MPSRNNMatrixIdGRURecurrentGateRecurrentWeights MPSRNNMatrixId = 23
	MPSRNNMatrixIdLSTMForgetGateBiasTerms          MPSRNNMatrixId = 10
	MPSRNNMatrixIdLSTMForgetGateInputWeights       MPSRNNMatrixId = 7
	MPSRNNMatrixIdLSTMForgetGateMemoryWeights      MPSRNNMatrixId = 9
	MPSRNNMatrixIdLSTMForgetGateRecurrentWeights   MPSRNNMatrixId = 8
	MPSRNNMatrixIdLSTMInputGateBiasTerms           MPSRNNMatrixId = 6
	MPSRNNMatrixIdLSTMInputGateInputWeights        MPSRNNMatrixId = 3
	MPSRNNMatrixIdLSTMInputGateMemoryWeights       MPSRNNMatrixId = 5
	MPSRNNMatrixIdLSTMInputGateRecurrentWeights    MPSRNNMatrixId = 4
	MPSRNNMatrixIdLSTMMemoryGateBiasTerms          MPSRNNMatrixId = 14
	MPSRNNMatrixIdLSTMMemoryGateInputWeights       MPSRNNMatrixId = 11
	MPSRNNMatrixIdLSTMMemoryGateMemoryWeights      MPSRNNMatrixId = 13
	MPSRNNMatrixIdLSTMMemoryGateRecurrentWeights   MPSRNNMatrixId = 12
	MPSRNNMatrixIdLSTMOutputGateBiasTerms          MPSRNNMatrixId = 18
	MPSRNNMatrixIdLSTMOutputGateInputWeights       MPSRNNMatrixId = 15
	MPSRNNMatrixIdLSTMOutputGateMemoryWeights      MPSRNNMatrixId = 17
	MPSRNNMatrixIdLSTMOutputGateRecurrentWeights   MPSRNNMatrixId = 16
	MPSRNNMatrixIdSingleGateBiasTerms              MPSRNNMatrixId = 2
	MPSRNNMatrixIdSingleGateInputWeights           MPSRNNMatrixId = 0
	MPSRNNMatrixIdSingleGateRecurrentWeights       MPSRNNMatrixId = 1
	MPSRNNMatrixId_count                           MPSRNNMatrixId = 29
)

func (e MPSRNNMatrixId) String() string {
	switch e {
	case MPSRNNMatrixIdGRUInputGateBiasTerms:
		return "MPSRNNMatrixIdGRUInputGateBiasTerms"
	case MPSRNNMatrixIdGRUInputGateInputWeights:
		return "MPSRNNMatrixIdGRUInputGateInputWeights"
	case MPSRNNMatrixIdGRUInputGateRecurrentWeights:
		return "MPSRNNMatrixIdGRUInputGateRecurrentWeights"
	case MPSRNNMatrixIdGRUOutputGateBiasTerms:
		return "MPSRNNMatrixIdGRUOutputGateBiasTerms"
	case MPSRNNMatrixIdGRUOutputGateInputGateWeights:
		return "MPSRNNMatrixIdGRUOutputGateInputGateWeights"
	case MPSRNNMatrixIdGRUOutputGateInputWeights:
		return "MPSRNNMatrixIdGRUOutputGateInputWeights"
	case MPSRNNMatrixIdGRUOutputGateRecurrentWeights:
		return "MPSRNNMatrixIdGRUOutputGateRecurrentWeights"
	case MPSRNNMatrixIdGRURecurrentGateBiasTerms:
		return "MPSRNNMatrixIdGRURecurrentGateBiasTerms"
	case MPSRNNMatrixIdGRURecurrentGateInputWeights:
		return "MPSRNNMatrixIdGRURecurrentGateInputWeights"
	case MPSRNNMatrixIdGRURecurrentGateRecurrentWeights:
		return "MPSRNNMatrixIdGRURecurrentGateRecurrentWeights"
	case MPSRNNMatrixIdLSTMForgetGateBiasTerms:
		return "MPSRNNMatrixIdLSTMForgetGateBiasTerms"
	case MPSRNNMatrixIdLSTMForgetGateInputWeights:
		return "MPSRNNMatrixIdLSTMForgetGateInputWeights"
	case MPSRNNMatrixIdLSTMForgetGateMemoryWeights:
		return "MPSRNNMatrixIdLSTMForgetGateMemoryWeights"
	case MPSRNNMatrixIdLSTMForgetGateRecurrentWeights:
		return "MPSRNNMatrixIdLSTMForgetGateRecurrentWeights"
	case MPSRNNMatrixIdLSTMInputGateBiasTerms:
		return "MPSRNNMatrixIdLSTMInputGateBiasTerms"
	case MPSRNNMatrixIdLSTMInputGateInputWeights:
		return "MPSRNNMatrixIdLSTMInputGateInputWeights"
	case MPSRNNMatrixIdLSTMInputGateMemoryWeights:
		return "MPSRNNMatrixIdLSTMInputGateMemoryWeights"
	case MPSRNNMatrixIdLSTMInputGateRecurrentWeights:
		return "MPSRNNMatrixIdLSTMInputGateRecurrentWeights"
	case MPSRNNMatrixIdLSTMMemoryGateBiasTerms:
		return "MPSRNNMatrixIdLSTMMemoryGateBiasTerms"
	case MPSRNNMatrixIdLSTMMemoryGateInputWeights:
		return "MPSRNNMatrixIdLSTMMemoryGateInputWeights"
	case MPSRNNMatrixIdLSTMMemoryGateMemoryWeights:
		return "MPSRNNMatrixIdLSTMMemoryGateMemoryWeights"
	case MPSRNNMatrixIdLSTMMemoryGateRecurrentWeights:
		return "MPSRNNMatrixIdLSTMMemoryGateRecurrentWeights"
	case MPSRNNMatrixIdLSTMOutputGateBiasTerms:
		return "MPSRNNMatrixIdLSTMOutputGateBiasTerms"
	case MPSRNNMatrixIdLSTMOutputGateInputWeights:
		return "MPSRNNMatrixIdLSTMOutputGateInputWeights"
	case MPSRNNMatrixIdLSTMOutputGateMemoryWeights:
		return "MPSRNNMatrixIdLSTMOutputGateMemoryWeights"
	case MPSRNNMatrixIdLSTMOutputGateRecurrentWeights:
		return "MPSRNNMatrixIdLSTMOutputGateRecurrentWeights"
	case MPSRNNMatrixIdSingleGateBiasTerms:
		return "MPSRNNMatrixIdSingleGateBiasTerms"
	case MPSRNNMatrixIdSingleGateInputWeights:
		return "MPSRNNMatrixIdSingleGateInputWeights"
	case MPSRNNMatrixIdSingleGateRecurrentWeights:
		return "MPSRNNMatrixIdSingleGateRecurrentWeights"
	case MPSRNNMatrixId_count:
		return "MPSRNNMatrixId_count"
	default:
		return fmt.Sprintf("MPSRNNMatrixId(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSequenceDirection
type MPSRNNSequenceDirection uint

const (
	MPSRNNSequenceDirectionBackward MPSRNNSequenceDirection = 1
	MPSRNNSequenceDirectionForward  MPSRNNSequenceDirection = 0
)

func (e MPSRNNSequenceDirection) String() string {
	switch e {
	case MPSRNNSequenceDirectionBackward:
		return "MPSRNNSequenceDirectionBackward"
	case MPSRNNSequenceDirectionForward:
		return "MPSRNNSequenceDirectionForward"
	default:
		return fmt.Sprintf("MPSRNNSequenceDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayDataType
type MPSRayDataType uint

const (
	MPSRayDataTypeOriginDirection                       MPSRayDataType = 0
	MPSRayDataTypeOriginMaskDirectionMaxDistance        MPSRayDataType = 2
	MPSRayDataTypeOriginMinDistanceDirectionMaxDistance MPSRayDataType = 1
	MPSRayDataTypePackedOriginDirection                 MPSRayDataType = 3
)

func (e MPSRayDataType) String() string {
	switch e {
	case MPSRayDataTypeOriginDirection:
		return "MPSRayDataTypeOriginDirection"
	case MPSRayDataTypeOriginMaskDirectionMaxDistance:
		return "MPSRayDataTypeOriginMaskDirectionMaxDistance"
	case MPSRayDataTypeOriginMinDistanceDirectionMaxDistance:
		return "MPSRayDataTypeOriginMinDistanceDirectionMaxDistance"
	case MPSRayDataTypePackedOriginDirection:
		return "MPSRayDataTypePackedOriginDirection"
	default:
		return fmt.Sprintf("MPSRayDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOperator
type MPSRayMaskOperator uint

const (
	// Deprecated.
	MPSRayMaskOperatorAnd MPSRayMaskOperator = 0
	// Deprecated.
	MPSRayMaskOperatorEqual MPSRayMaskOperator = 10
	// Deprecated.
	MPSRayMaskOperatorGreaterThan MPSRayMaskOperator = 8
	// Deprecated.
	MPSRayMaskOperatorGreaterThanOrEqualTo MPSRayMaskOperator = 9
	// Deprecated.
	MPSRayMaskOperatorLessThan MPSRayMaskOperator = 6
	// Deprecated.
	MPSRayMaskOperatorLessThanOrEqualTo MPSRayMaskOperator = 7
	// Deprecated.
	MPSRayMaskOperatorNotAnd MPSRayMaskOperator = 1
	// Deprecated.
	MPSRayMaskOperatorNotEqual MPSRayMaskOperator = 11
	// Deprecated.
	MPSRayMaskOperatorNotOr MPSRayMaskOperator = 3
	// Deprecated.
	MPSRayMaskOperatorNotXor MPSRayMaskOperator = 5
	// Deprecated.
	MPSRayMaskOperatorOr MPSRayMaskOperator = 2
	// Deprecated.
	MPSRayMaskOperatorXor MPSRayMaskOperator = 4
)

func (e MPSRayMaskOperator) String() string {
	switch e {
	case MPSRayMaskOperatorAnd:
		return "MPSRayMaskOperatorAnd"
	case MPSRayMaskOperatorEqual:
		return "MPSRayMaskOperatorEqual"
	case MPSRayMaskOperatorGreaterThan:
		return "MPSRayMaskOperatorGreaterThan"
	case MPSRayMaskOperatorGreaterThanOrEqualTo:
		return "MPSRayMaskOperatorGreaterThanOrEqualTo"
	case MPSRayMaskOperatorLessThan:
		return "MPSRayMaskOperatorLessThan"
	case MPSRayMaskOperatorLessThanOrEqualTo:
		return "MPSRayMaskOperatorLessThanOrEqualTo"
	case MPSRayMaskOperatorNotAnd:
		return "MPSRayMaskOperatorNotAnd"
	case MPSRayMaskOperatorNotEqual:
		return "MPSRayMaskOperatorNotEqual"
	case MPSRayMaskOperatorNotOr:
		return "MPSRayMaskOperatorNotOr"
	case MPSRayMaskOperatorNotXor:
		return "MPSRayMaskOperatorNotXor"
	case MPSRayMaskOperatorOr:
		return "MPSRayMaskOperatorOr"
	case MPSRayMaskOperatorXor:
		return "MPSRayMaskOperatorXor"
	default:
		return fmt.Sprintf("MPSRayMaskOperator(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayMaskOptions
type MPSRayMaskOptions uint

const (
	// Deprecated.
	MPSRayMaskOptionInstance MPSRayMaskOptions = 2
	// Deprecated.
	MPSRayMaskOptionNone MPSRayMaskOptions = 0
	// Deprecated.
	MPSRayMaskOptionPrimitive MPSRayMaskOptions = 1
)

func (e MPSRayMaskOptions) String() string {
	switch e {
	case MPSRayMaskOptionInstance:
		return "MPSRayMaskOptionInstance"
	case MPSRayMaskOptionNone:
		return "MPSRayMaskOptionNone"
	case MPSRayMaskOptionPrimitive:
		return "MPSRayMaskOptionPrimitive"
	default:
		return fmt.Sprintf("MPSRayMaskOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceType
type MPSStateResourceType uint

const (
	MPSStateResourceTypeBuffer  MPSStateResourceType = 1
	MPSStateResourceTypeNone    MPSStateResourceType = 0
	MPSStateResourceTypeTexture MPSStateResourceType = 2
)

func (e MPSStateResourceType) String() string {
	switch e {
	case MPSStateResourceTypeBuffer:
		return "MPSStateResourceTypeBuffer"
	case MPSStateResourceTypeNone:
		return "MPSStateResourceTypeNone"
	case MPSStateResourceTypeTexture:
		return "MPSStateResourceTypeTexture"
	default:
		return fmt.Sprintf("MPSStateResourceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalWeighting
type MPSTemporalWeighting uint

const (
	MPSTemporalWeightingAverage                  MPSTemporalWeighting = 0
	MPSTemporalWeightingExponentialMovingAverage MPSTemporalWeighting = 1
)

func (e MPSTemporalWeighting) String() string {
	switch e {
	case MPSTemporalWeightingAverage:
		return "MPSTemporalWeightingAverage"
	case MPSTemporalWeightingExponentialMovingAverage:
		return "MPSTemporalWeightingExponentialMovingAverage"
	default:
		return fmt.Sprintf("MPSTemporalWeighting(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTransformType
type MPSTransformType uint

const (
	MPSTransformTypeFloat4x4 MPSTransformType = 0
	MPSTransformTypeIdentity MPSTransformType = 1
)

func (e MPSTransformType) String() string {
	switch e {
	case MPSTransformTypeFloat4x4:
		return "MPSTransformTypeFloat4x4"
	case MPSTransformTypeIdentity:
		return "MPSTransformTypeIdentity"
	default:
		return fmt.Sprintf("MPSTransformType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleIntersectionTestType
type MPSTriangleIntersectionTestType uint

const (
	// Deprecated.
	MPSTriangleIntersectionTestTypeDefault MPSTriangleIntersectionTestType = 0
	// Deprecated.
	MPSTriangleIntersectionTestTypeWatertight MPSTriangleIntersectionTestType = 1
)

func (e MPSTriangleIntersectionTestType) String() string {
	switch e {
	case MPSTriangleIntersectionTestTypeDefault:
		return "MPSTriangleIntersectionTestTypeDefault"
	case MPSTriangleIntersectionTestTypeWatertight:
		return "MPSTriangleIntersectionTestTypeWatertight"
	default:
		return fmt.Sprintf("MPSTriangleIntersectionTestType(%d)", e)
	}
}
