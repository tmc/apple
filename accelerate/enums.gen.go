// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Accelerate/BLAS_THREADING
type BLAS_THREADING uint32

const (
	BLAS_THREADING_MAX_OPTIONS BLAS_THREADING = 2
	// BLAS_THREADING_MULTI_THREADED: A constant that specifies that the Accelerate framework decides whether BLAS and LAPACK execute on single or multiple threads.
	BLAS_THREADING_MULTI_THREADED BLAS_THREADING = 0
	// BLAS_THREADING_SINGLE_THREADED: A constant that specifies BLAS and LAPACK execute on a single thread only.
	BLAS_THREADING_SINGLE_THREADED BLAS_THREADING = 1
)

func (e BLAS_THREADING) String() string {
	switch e {
	case BLAS_THREADING_MAX_OPTIONS:
		return "BLAS_THREADING_MAX_OPTIONS"
	case BLAS_THREADING_MULTI_THREADED:
		return "BLAS_THREADING_MULTI_THREADED"
	case BLAS_THREADING_SINGLE_THREADED:
		return "BLAS_THREADING_SINGLE_THREADED"
	default:
		return fmt.Sprintf("BLAS_THREADING(%d)", e)
	}
}

type BNNSActivationFunction uint

const (
	// BNNSActivationFunctionAbs: An activation function that returns the absolute value of its input.
	BNNSActivationFunctionAbs BNNSActivationFunction = 4
	// BNNSActivationFunctionCELU: An activation function that evaluates the continuously differentiable exponential linear units (CELU) on its input.
	BNNSActivationFunctionCELU BNNSActivationFunction = 8
	// BNNSActivationFunctionClamp: An activation function that returns its input clamped to the specified range.
	BNNSActivationFunctionClamp BNNSActivationFunction = 5
	// BNNSActivationFunctionClampedLeakyRectifiedLinear: An activation function that returns its input clamped to beta when that is greater than or equal to zero, otherwise it returns its input multiplied by alpha clamped to beta.
	BNNSActivationFunctionClampedLeakyRectifiedLinear BNNSActivationFunction = 8
	// BNNSActivationFunctionELU: An activation function that evaluates the exponential linear units (ELU) on its input.
	BNNSActivationFunctionELU  BNNSActivationFunction = 8
	BNNSActivationFunctionErf  BNNSActivationFunction = 10
	BNNSActivationFunctionGELU BNNSActivationFunction = 10
	// BNNSActivationFunctionGELUApproximation: An activation function that evaluates the Gaussian error linear units (GELU) approximation on its input.
	BNNSActivationFunctionGELUApproximation BNNSActivationFunction = 5
	// BNNSActivationFunctionGELUApproximation2: An activation function that provides a fast evaluation of the Gaussian error linear units (GELU) approximation on its input.
	BNNSActivationFunctionGELUApproximation2       BNNSActivationFunction = 8
	BNNSActivationFunctionGELUApproximationSigmoid BNNSActivationFunction = 10
	// BNNSActivationFunctionGumbel: An activation function that returns random numbers from the Gumbel distribution.
	BNNSActivationFunctionGumbel BNNSActivationFunction = 8
	// BNNSActivationFunctionGumbelMax: An activation function that returns random numbers from the Gumbel distribution.
	BNNSActivationFunctionGumbelMax BNNSActivationFunction = 8
	// BNNSActivationFunctionHardShrink: An activation function that returns zero when the absolute input is less than alpha, otherwise it returns its input.
	BNNSActivationFunctionHardShrink BNNSActivationFunction = 8
	// BNNSActivationFunctionHardSigmoid: An activation function that returns the hard sigmoid function of its input.
	BNNSActivationFunctionHardSigmoid BNNSActivationFunction = 8
	// BNNSActivationFunctionHardSwish: An activation function that returns the hard swish function of its input.
	BNNSActivationFunctionHardSwish BNNSActivationFunction = 8
	// BNNSActivationFunctionIdentity: An activation function that returns its input.
	BNNSActivationFunctionIdentity BNNSActivationFunction = 0
	// BNNSActivationFunctionIntegerLinearSaturate: An activation function that returns an arithmetic shift, preserving sign.
	BNNSActivationFunctionIntegerLinearSaturate BNNSActivationFunction = 5
	// BNNSActivationFunctionIntegerLinearSaturatePerChannel: An activation function that returns an arithmetic shift, preserving sign for each channel.
	BNNSActivationFunctionIntegerLinearSaturatePerChannel BNNSActivationFunction = 5
	// BNNSActivationFunctionLeakyRectifiedLinear: An activation function that returns its input when that is greater than or equal to zero, otherwise it returns its input multiplied by a specified value.
	BNNSActivationFunctionLeakyRectifiedLinear BNNSActivationFunction = 4
	// BNNSActivationFunctionLinear: An activation function that returns its input multiplied by a specified value.
	BNNSActivationFunctionLinear BNNSActivationFunction = 4
	// BNNSActivationFunctionLinearWithBias: An activation function that returns its input multiplied by a scale and added to a bias.
	BNNSActivationFunctionLinearWithBias BNNSActivationFunction = 8
	// BNNSActivationFunctionLogSigmoid: An activation function that returns the logarithm of the sigmoid function of its input.
	BNNSActivationFunctionLogSigmoid BNNSActivationFunction = 8
	// BNNSActivationFunctionLogSoftmax: An activation function that returns the logarithm of the softmax function of its input.
	BNNSActivationFunctionLogSoftmax BNNSActivationFunction = 8
	// BNNSActivationFunctionPReLUPerChannel: An activation function provides per-channel alpha values to Leaky Rectified Linear.
	BNNSActivationFunctionPReLUPerChannel BNNSActivationFunction = 8
	BNNSActivationFunctionReLU6           BNNSActivationFunction = 9
	// BNNSActivationFunctionRectifiedLinear: An activation function that returns its input when that is greater than or equal to zero, otherwise it returns zero.
	BNNSActivationFunctionRectifiedLinear BNNSActivationFunction = 4
	// BNNSActivationFunctionSELU: An activation function that evaluates the scaled exponential linear units (SELU) on its input.
	BNNSActivationFunctionSELU BNNSActivationFunction = 8
	// BNNSActivationFunctionScaledTanh: An activation function that returns the scaled hyperbolic tangent of its input.
	BNNSActivationFunctionScaledTanh BNNSActivationFunction = 4
	// BNNSActivationFunctionSiLU: An activation function that returns the sigmoid linear unit (SiLU) function of its input.
	BNNSActivationFunctionSiLU BNNSActivationFunction = 9
	// BNNSActivationFunctionSigmoid: An activation function that returns the sigmoid function of its input.
	BNNSActivationFunctionSigmoid BNNSActivationFunction = 4
	// BNNSActivationFunctionSoftShrink: An activation function that returns zero when the absolute input is less than alpha, otherwise it returns its input minus alpha.
	BNNSActivationFunctionSoftShrink BNNSActivationFunction = 8
	// BNNSActivationFunctionSoftmax: An activation function that returns the softmax function of its input.
	BNNSActivationFunctionSoftmax BNNSActivationFunction = 5
	// BNNSActivationFunctionSoftplus: An activation function that returns the softplus function of its input.
	BNNSActivationFunctionSoftplus BNNSActivationFunction = 8
	// BNNSActivationFunctionSoftsign: An activation function that returns the softsign function of its input.
	BNNSActivationFunctionSoftsign BNNSActivationFunction = 8
	// BNNSActivationFunctionTanh: An activation function that returns the hyperbolic tangent of its input.
	BNNSActivationFunctionTanh BNNSActivationFunction = 4
	// BNNSActivationFunctionTanhShrink: An activation function that returns its input minus the hyperbolic tangent of its input.
	BNNSActivationFunctionTanhShrink BNNSActivationFunction = 8
	// BNNSActivationFunctionThreshold: An activation function that returns beta if its input is less than a specified threshold, otherwise it returns its input.
	BNNSActivationFunctionThreshold BNNSActivationFunction = 8
)

func (e BNNSActivationFunction) String() string {
	switch e {
	case BNNSActivationFunctionAbs:
		return "BNNSActivationFunctionAbs"
	case BNNSActivationFunctionCELU:
		return "BNNSActivationFunctionCELU"
	case BNNSActivationFunctionClamp:
		return "BNNSActivationFunctionClamp"
	case BNNSActivationFunctionErf:
		return "BNNSActivationFunctionErf"
	case BNNSActivationFunctionIdentity:
		return "BNNSActivationFunctionIdentity"
	case BNNSActivationFunctionReLU6:
		return "BNNSActivationFunctionReLU6"
	default:
		return fmt.Sprintf("BNNSActivationFunction(%d)", e)
	}
}

type BNNSArithmetic uint

const (
	// BNNSArithmeticAbs: An operation that calculates the element-wise absolute of its input.
	BNNSArithmeticAbs BNNSArithmetic = 9
	// BNNSArithmeticAcos: An operation that calculates the element-wise inverse cosine of its input.
	BNNSArithmeticAcos BNNSArithmetic = 8
	// BNNSArithmeticAcosh: An operation that calculates the element-wise inverse hyperbolic cosine of its input.
	BNNSArithmeticAcosh BNNSArithmetic = 8
	// BNNSArithmeticAdd: An operation that calculates the element-wise sum of its two inputs.
	BNNSArithmeticAdd BNNSArithmetic = 0
	// BNNSArithmeticAsin: An operation that calculates the element-wise inverse sine of its input.
	BNNSArithmeticAsin BNNSArithmetic = 8
	// BNNSArithmeticAsinh: An operation that calculates the element-wise inverse hyperbolic sine of its input.
	BNNSArithmeticAsinh BNNSArithmetic = 8
	// BNNSArithmeticAtan: An operation that calculates the element-wise inverse tangent of its input.
	BNNSArithmeticAtan BNNSArithmetic = 8
	// BNNSArithmeticAtanh: An operation that calculates the element-wise inverse hyperbolic tangent of its input.
	BNNSArithmeticAtanh BNNSArithmetic = 8
	// BNNSArithmeticCeil: An operation that calculates the element-wise ceiling of its input.
	BNNSArithmeticCeil BNNSArithmetic = 8
	// BNNSArithmeticCos: An operation that calculates the element-wise cosine of its input.
	BNNSArithmeticCos BNNSArithmetic = 8
	// BNNSArithmeticCosh: An operation that calculates the element-wise hyperbolic cosine of its input.
	BNNSArithmeticCosh BNNSArithmetic = 8
	// BNNSArithmeticDivide: An operation that calculates the element-wise division of its two inputs.
	BNNSArithmeticDivide BNNSArithmetic = 8
	// BNNSArithmeticDivideNoNaN: An operation that calculates the element-wise division of its two inputs and returns zero if the divisor is zero, even if the first input is NaN or infinity.
	BNNSArithmeticDivideNoNaN BNNSArithmetic = 8
	// BNNSArithmeticErf: An operation that calculates the element-wise error function of its input.
	BNNSArithmeticErf BNNSArithmetic = 10
	// BNNSArithmeticExp: An operation that calculates the element-wise result of  raised to the power of its input.
	BNNSArithmeticExp BNNSArithmetic = 8
	// BNNSArithmeticExp2: An operation that calculates the element-wise result of 2 raised to the power of its input.
	BNNSArithmeticExp2 BNNSArithmetic = 8
	// BNNSArithmeticFloor: An operation that calculates the element-wise floor of its input.
	BNNSArithmeticFloor BNNSArithmetic = 8
	// BNNSArithmeticFloorDivide: An operation that calculates the element-wise floor division of its inputs.
	BNNSArithmeticFloorDivide BNNSArithmetic = 10
	// BNNSArithmeticLog: An operation that calculates the element-wise natural logarithm of its input.
	BNNSArithmeticLog BNNSArithmetic = 8
	// BNNSArithmeticLog2: An operation that calculates the element-wise base 2 logarithm of its input.
	BNNSArithmeticLog2 BNNSArithmetic = 8
	// BNNSArithmeticMaximum: An operation that calculates the element-wise maximum of its two inputs.
	BNNSArithmeticMaximum BNNSArithmetic = 9
	// BNNSArithmeticMinimum: An operation that calculates the element-wise minimum of its two inputs.
	BNNSArithmeticMinimum BNNSArithmetic = 9
	// BNNSArithmeticMultiply: An operation that calculates the element-wise product of its two inputs.
	BNNSArithmeticMultiply BNNSArithmetic = 8
	// BNNSArithmeticMultiplyAdd: An operation that calculates the element-wise fused multiply-add of its three inputs.
	BNNSArithmeticMultiplyAdd BNNSArithmetic = 8
	// BNNSArithmeticMultiplyNoNaN: An operation that calculates the element-wise product of its two inputs and returns zero, even if the first input is NaN or infinity.
	BNNSArithmeticMultiplyNoNaN BNNSArithmetic = 8
	// BNNSArithmeticNegate: An operation that calculates the element-wise negation of its input.
	BNNSArithmeticNegate BNNSArithmetic = 10
	// BNNSArithmeticPow: An operation that calculates the element-wise first input raised to the power of its second input.
	BNNSArithmeticPow BNNSArithmetic = 8
	// BNNSArithmeticReciprocal: An operation that calculates the element-wise reciprocal of its input.
	BNNSArithmeticReciprocal BNNSArithmetic = 10
	// BNNSArithmeticReciprocalSquareRoot: An operation that calculates the element-wise reciprocal square root of its input.
	BNNSArithmeticReciprocalSquareRoot BNNSArithmetic = 8
	// BNNSArithmeticRound: An operation that calculates the element-wise rounding of its input.
	BNNSArithmeticRound BNNSArithmetic = 8
	// BNNSArithmeticSelect: An operation that selects elements from either its second or third input based on the corresponding value of its first input.
	BNNSArithmeticSelect BNNSArithmetic = 9
	// BNNSArithmeticSign: An operation that calculates the element-wise sign of its input.
	BNNSArithmeticSign BNNSArithmetic = 10
	// BNNSArithmeticSin: An operation that calculates the element-wise sine of its input.
	BNNSArithmeticSin BNNSArithmetic = 8
	// BNNSArithmeticSinh: An operation that calculates the element-wise hyperbolic sine of its input.
	BNNSArithmeticSinh BNNSArithmetic = 8
	// BNNSArithmeticSquare: An operation that calculates the element-wise square of its input.
	BNNSArithmeticSquare BNNSArithmetic = 10
	// BNNSArithmeticSquareRoot: An operation that calculates the element-wise square root of its input.
	BNNSArithmeticSquareRoot BNNSArithmetic = 8
	// BNNSArithmeticSubtract: An operation that calculates the element-wise difference of its two inputs.
	BNNSArithmeticSubtract BNNSArithmetic = 8
	// BNNSArithmeticTan: An operation that calculates the element-wise tangent of its input.
	BNNSArithmeticTan BNNSArithmetic = 8
	// BNNSArithmeticTanh: An operation that calculates the element-wise hyperbolic tangent of its input.
	BNNSArithmeticTanh BNNSArithmetic = 8
	// BNNSArithmeticTruncDivide: An operation that calculates the element-wise truncated division of its inputs.
	BNNSArithmeticTruncDivide BNNSArithmetic = 10
	// BNNSArithmeticTruncRemainder: An operation that calculates the element-wise remainder of truncated division of its inputs.
	BNNSArithmeticTruncRemainder BNNSArithmetic = 10
)

func (e BNNSArithmetic) String() string {
	switch e {
	case BNNSArithmeticAbs:
		return "BNNSArithmeticAbs"
	case BNNSArithmeticAcos:
		return "BNNSArithmeticAcos"
	case BNNSArithmeticAdd:
		return "BNNSArithmeticAdd"
	case BNNSArithmeticErf:
		return "BNNSArithmeticErf"
	default:
		return fmt.Sprintf("BNNSArithmetic(%d)", e)
	}
}

type BNNSData uint

const (
	// BNNSDataLayout1DFirstMajor: A constant that represents a 1D first-major vector.
	BNNSDataLayout1DFirstMajor BNNSData = 8
	// BNNSDataLayout1DLastMajor: A constant that represents a 1D last-major vector.
	BNNSDataLayout1DLastMajor BNNSData = 8
	// BNNSDataLayout2DFirstMajor: A constant that represents a 2D first-major matrix.
	BNNSDataLayout2DFirstMajor BNNSData = 8
	// BNNSDataLayout2DLastMajor: A constant that represents a 2D last-major matrix.
	BNNSDataLayout2DLastMajor BNNSData = 8
	// BNNSDataLayout3DFirstMajor: A constant that represents a 3D first-major tensor.
	BNNSDataLayout3DFirstMajor BNNSData = 8
	// BNNSDataLayout3DLastMajor: A constant that represents a 3D last-major tensor.
	BNNSDataLayout3DLastMajor BNNSData = 10
	// BNNSDataLayout4DFirstMajor: A constant that represents a 4D first-major tensor.
	BNNSDataLayout4DFirstMajor BNNSData = 8
	// BNNSDataLayout4DLastMajor: A constant that represents a 4D last-major tensor.
	BNNSDataLayout4DLastMajor BNNSData = 8
	// BNNSDataLayout5DFirstMajor: A constant that represents a 5D first-major tensor.
	BNNSDataLayout5DFirstMajor BNNSData = 8
	// BNNSDataLayout5DLastMajor: A constant that represents a 5D last-major tensor.
	BNNSDataLayout5DLastMajor BNNSData = 8
	// BNNSDataLayout6DFirstMajor: A constant that represents a 6D first-major tensor.
	BNNSDataLayout6DFirstMajor BNNSData = 8
	// BNNSDataLayout6DLastMajor: A constant that represents a 6D last-major tensor.
	BNNSDataLayout6DLastMajor BNNSData = 8
	// BNNSDataLayout7DFirstMajor: A constant that represents a 7D first-major tensor.
	BNNSDataLayout7DFirstMajor BNNSData = 8
	// BNNSDataLayout7DLastMajor: A constant that represents a 7D last-major tensor.
	BNNSDataLayout7DLastMajor BNNSData = 8
	// BNNSDataLayout8DFirstMajor: A constant that represents a 8D first-major tensor.
	BNNSDataLayout8DFirstMajor BNNSData = 8
	// BNNSDataLayout8DLastMajor: A constant that represents a 8D last-major tensor.
	BNNSDataLayout8DLastMajor BNNSData = 8
	// BNNSDataLayoutColumnMajorMatrix: A constant that represents a 2D column-major matrix.
	BNNSDataLayoutColumnMajorMatrix BNNSData = 8
	// BNNSDataLayoutConvolutionWeightsIOHrWr: A constant that represents a 4D array of rotated convolution weights.
	BNNSDataLayoutConvolutionWeightsIOHrWr BNNSData = 8
	// BNNSDataLayoutConvolutionWeightsOIHW: A constant that represents a 4D array of convolution weights.
	BNNSDataLayoutConvolutionWeightsOIHW BNNSData = 8
	// BNNSDataLayoutConvolutionWeightsOIHW_Pack32: A constant that represents a 4D array of packed convolution weights with 32-output channel packing and 128-byte array address alignment.
	BNNSDataLayoutConvolutionWeightsOIHW_Pack32 BNNSData = 8
	// BNNSDataLayoutConvolutionWeightsOIHrWr: A constant that represents a 4D array of rotated convolution weights.
	BNNSDataLayoutConvolutionWeightsOIHrWr BNNSData = 8
	BNNSDataLayoutFullyConnectedSparse     BNNSData = 8
	// BNNSDataLayoutImageCHW: A constant that represents a 3D image stack.
	BNNSDataLayoutImageCHW BNNSData = 10
	BNNSDataLayoutMHA_DHK  BNNSData = 9
	// BNNSDataLayoutNSE: A constant that represents a 3D tensor with the size elements embedding dimension, sequence length, and batch size.
	BNNSDataLayoutNSE BNNSData = 9
	// BNNSDataLayoutRowMajorMatrix: A constant that represents a 2D row-major matrix.
	BNNSDataLayoutRowMajorMatrix BNNSData = 8
	// BNNSDataLayoutSNE: A constant that represents a 3D tensor with the size elements embedding dimension, batch size, and sequence length.
	BNNSDataLayoutSNE BNNSData = 7
	// BNNSDataLayoutVector: A constant that represents a 1D vector.
	BNNSDataLayoutVector BNNSData = 0
)

func (e BNNSData) String() string {
	switch e {
	case BNNSDataLayout1DFirstMajor:
		return "BNNSDataLayout1DFirstMajor"
	case BNNSDataLayout3DLastMajor:
		return "BNNSDataLayout3DLastMajor"
	case BNNSDataLayoutMHA_DHK:
		return "BNNSDataLayoutMHA_DHK"
	case BNNSDataLayoutSNE:
		return "BNNSDataLayoutSNE"
	case BNNSDataLayoutVector:
		return "BNNSDataLayoutVector"
	default:
		return fmt.Sprintf("BNNSData(%d)", e)
	}
}

type BNNSDataType uint

const (
	// BNNSDataTypeBFloat16: # Discussion
	BNNSDataTypeBFloat16 BNNSDataType = 11
	// BNNSDataTypeBoolean: # Discussion
	BNNSDataTypeBoolean BNNSDataType = 8
	// BNNSDataTypeFloat16: # Discussion
	BNNSDataTypeFloat16 BNNSDataType = 4
	// BNNSDataTypeFloat32: # Discussion
	BNNSDataTypeFloat32 BNNSDataType = 11
	// BNNSDataTypeFloatBit: # Discussion
	BNNSDataTypeFloatBit BNNSDataType = 0
	// BNNSDataTypeIndexed1: # Discussion
	BNNSDataTypeIndexed1 BNNSDataType = 4
	// BNNSDataTypeIndexed2: # Discussion
	BNNSDataTypeIndexed2 BNNSDataType = 15
	// BNNSDataTypeIndexed4: # Discussion
	BNNSDataTypeIndexed4 BNNSDataType = 15
	// BNNSDataTypeIndexed8: # Discussion
	BNNSDataTypeIndexed8 BNNSDataType = 15
	// BNNSDataTypeIndexedBit: # Discussion
	BNNSDataTypeIndexedBit BNNSDataType = 16
	// BNNSDataTypeInt1: # Discussion
	BNNSDataTypeInt1 BNNSDataType = 4
	// BNNSDataTypeInt16: # Discussion
	BNNSDataTypeInt16 BNNSDataType = 11
	// BNNSDataTypeInt2: # Discussion
	BNNSDataTypeInt2 BNNSDataType = 15
	// BNNSDataTypeInt32: # Discussion
	BNNSDataTypeInt32 BNNSDataType = 11
	// BNNSDataTypeInt4: # Discussion
	BNNSDataTypeInt4 BNNSDataType = 15
	// BNNSDataTypeInt64: # Discussion
	BNNSDataTypeInt64 BNNSDataType = 11
	// BNNSDataTypeInt8: # Discussion
	BNNSDataTypeInt8 BNNSDataType = 15
	// BNNSDataTypeIntBit: # Discussion
	BNNSDataTypeIntBit BNNSDataType = 16
	// BNNSDataTypeMiscellaneousBit: # Discussion
	BNNSDataTypeMiscellaneousBit BNNSDataType = 11
	// BNNSDataTypeUInt1: # Discussion
	BNNSDataTypeUInt1 BNNSDataType = 5
	// BNNSDataTypeUInt16: # Discussion
	BNNSDataTypeUInt16 BNNSDataType = 12
	// BNNSDataTypeUInt2: # Discussion
	BNNSDataTypeUInt2 BNNSDataType = 15
	// BNNSDataTypeUInt3: # Discussion
	BNNSDataTypeUInt3 BNNSDataType = 15
	// BNNSDataTypeUInt32: # Discussion
	BNNSDataTypeUInt32 BNNSDataType = 12
	// BNNSDataTypeUInt4: # Discussion
	BNNSDataTypeUInt4 BNNSDataType = 19
	// BNNSDataTypeUInt6: # Discussion
	BNNSDataTypeUInt6 BNNSDataType = 15
	// BNNSDataTypeUInt64: # Discussion
	BNNSDataTypeUInt64 BNNSDataType = 12
	// BNNSDataTypeUInt8: # Discussion
	BNNSDataTypeUInt8 BNNSDataType = 19
	// BNNSDataTypeUIntBit: # Discussion
	BNNSDataTypeUIntBit BNNSDataType = 16
)

func (e BNNSDataType) String() string {
	switch e {
	case BNNSDataTypeBFloat16:
		return "BNNSDataTypeBFloat16"
	case BNNSDataTypeBoolean:
		return "BNNSDataTypeBoolean"
	case BNNSDataTypeFloat16:
		return "BNNSDataTypeFloat16"
	case BNNSDataTypeFloatBit:
		return "BNNSDataTypeFloatBit"
	case BNNSDataTypeIndexed2:
		return "BNNSDataTypeIndexed2"
	case BNNSDataTypeIndexedBit:
		return "BNNSDataTypeIndexedBit"
	case BNNSDataTypeUInt1:
		return "BNNSDataTypeUInt1"
	case BNNSDataTypeUInt16:
		return "BNNSDataTypeUInt16"
	case BNNSDataTypeUInt4:
		return "BNNSDataTypeUInt4"
	default:
		return fmt.Sprintf("BNNSDataType(%d)", e)
	}
}

type BNNSEmbeddingFlagScaleGradientBy uint

const (
	// BNNSEmbeddingFlagScaleGradientByFrequency: A flag that specifies that the operation scales calculated gradients based on the number of occurrence of the corresponding index in the input.
	BNNSEmbeddingFlagScaleGradientByFrequency BNNSEmbeddingFlagScaleGradientBy = 0
)

func (e BNNSEmbeddingFlagScaleGradientBy) String() string {
	switch e {
	case BNNSEmbeddingFlagScaleGradientByFrequency:
		return "BNNSEmbeddingFlagScaleGradientByFrequency"
	default:
		return fmt.Sprintf("BNNSEmbeddingFlagScaleGradientBy(%d)", e)
	}
}

type BNNSFlagsUseClient uint

const (
	// BNNSFlagsUseClientPtr: A flag that instructs the filter to use pointers to data you provide at creation time.
	BNNSFlagsUseClientPtr BNNSFlagsUseClient = 0
)

func (e BNNSFlagsUseClient) String() string {
	switch e {
	case BNNSFlagsUseClientPtr:
		return "BNNSFlagsUseClientPtr"
	default:
		return fmt.Sprintf("BNNSFlagsUseClient(%d)", e)
	}
}

type BNNSGraphArgumentIntent uint

const (
	// BNNSGraphArgumentIntentIn: A constant that specifies the argument provides an input tensor.
	BNNSGraphArgumentIntentIn BNNSGraphArgumentIntent = 0
	// BNNSGraphArgumentIntentInOut: A constant that specifies the argument is an in-place input and output tensor.
	BNNSGraphArgumentIntentInOut BNNSGraphArgumentIntent = 12
	// BNNSGraphArgumentIntentOut: A constant that specifies the argument provides an output tensor.
	BNNSGraphArgumentIntentOut BNNSGraphArgumentIntent = 12
)

func (e BNNSGraphArgumentIntent) String() string {
	switch e {
	case BNNSGraphArgumentIntentIn:
		return "BNNSGraphArgumentIntentIn"
	case BNNSGraphArgumentIntentInOut:
		return "BNNSGraphArgumentIntentInOut"
	default:
		return fmt.Sprintf("BNNSGraphArgumentIntent(%d)", e)
	}
}

type BNNSGraphArgumentType uint

const (
	// BNNSGraphArgumentTypePointer: A pointer to the raw data for the tensor.
	BNNSGraphArgumentTypePointer BNNSGraphArgumentType = 0
	// BNNSGraphArgumentTypeTensor: A tensor structure.
	BNNSGraphArgumentTypeTensor BNNSGraphArgumentType = 12
)

func (e BNNSGraphArgumentType) String() string {
	switch e {
	case BNNSGraphArgumentTypePointer:
		return "BNNSGraphArgumentTypePointer"
	case BNNSGraphArgumentTypeTensor:
		return "BNNSGraphArgumentTypeTensor"
	default:
		return fmt.Sprintf("BNNSGraphArgumentType(%d)", e)
	}
}

type BNNSGraphMessageLevel uint

const (
	// BNNSGraphMessageLevelError: A constant that specifies error message types.
	BNNSGraphMessageLevelError BNNSGraphMessageLevel = 12
	// BNNSGraphMessageLevelInfo: A constant that specifies information message types.
	BNNSGraphMessageLevelInfo BNNSGraphMessageLevel = 0
	// BNNSGraphMessageLevelUnsupported: A constant that specifies unsupported function message types.
	BNNSGraphMessageLevelUnsupported BNNSGraphMessageLevel = 12
	// BNNSGraphMessageLevelWarning: A constant that specifies warning message types.
	BNNSGraphMessageLevelWarning BNNSGraphMessageLevel = 12
)

func (e BNNSGraphMessageLevel) String() string {
	switch e {
	case BNNSGraphMessageLevelError:
		return "BNNSGraphMessageLevelError"
	case BNNSGraphMessageLevelInfo:
		return "BNNSGraphMessageLevelInfo"
	default:
		return fmt.Sprintf("BNNSGraphMessageLevel(%d)", e)
	}
}

type BNNSGraphOptimizationPreference uint

const (
	// BNNSGraphOptimizationPreferenceIRSize: A constant that specifies compilation optimization for smallest graph size on disk.
	BNNSGraphOptimizationPreferenceIRSize BNNSGraphOptimizationPreference = 12
	// BNNSGraphOptimizationPreferencePerformance: A constant that specifies compilation optimization for best execution performance.
	BNNSGraphOptimizationPreferencePerformance BNNSGraphOptimizationPreference = 0
)

func (e BNNSGraphOptimizationPreference) String() string {
	switch e {
	case BNNSGraphOptimizationPreferenceIRSize:
		return "BNNSGraphOptimizationPreferenceIRSize"
	case BNNSGraphOptimizationPreferencePerformance:
		return "BNNSGraphOptimizationPreferencePerformance"
	default:
		return fmt.Sprintf("BNNSGraphOptimizationPreference(%d)", e)
	}
}

type BNNSInterpolationMethod uint

const (
	// BNNSInterpolationMethodLinear: Interpolation that is linear or bilinear depending on the number of resized dimensions.
	BNNSInterpolationMethodLinear BNNSInterpolationMethod = 8
	// BNNSInterpolationMethodNearest: Nearest-neighbor interpolation.
	BNNSInterpolationMethodNearest BNNSInterpolationMethod = 0
)

func (e BNNSInterpolationMethod) String() string {
	switch e {
	case BNNSInterpolationMethodLinear:
		return "BNNSInterpolationMethodLinear"
	case BNNSInterpolationMethodNearest:
		return "BNNSInterpolationMethodNearest"
	default:
		return fmt.Sprintf("BNNSInterpolationMethod(%d)", e)
	}
}

type BNNSLayerFlagsLSTM uint

const (
	// BNNSLayerFlagsLSTMBidirectional: A flag that enables bidirectional long short-term memory (LSTM).
	BNNSLayerFlagsLSTMBidirectional BNNSLayerFlagsLSTM = 0
	// BNNSLayerFlagsLSTMDefaultActivations: A flag that ignores the specified gate activations and instructs the operation to use default activations.
	BNNSLayerFlagsLSTMDefaultActivations BNNSLayerFlagsLSTM = 8
)

func (e BNNSLayerFlagsLSTM) String() string {
	switch e {
	case BNNSLayerFlagsLSTMBidirectional:
		return "BNNSLayerFlagsLSTMBidirectional"
	case BNNSLayerFlagsLSTMDefaultActivations:
		return "BNNSLayerFlagsLSTMDefaultActivations"
	default:
		return fmt.Sprintf("BNNSLayerFlagsLSTM(%d)", e)
	}
}

type BNNSLinearSampling uint

const (
	// BNNSLinearSamplingAlignCorners: The align corners sampling mode.
	BNNSLinearSamplingAlignCorners BNNSLinearSampling = 10
	// BNNSLinearSamplingDefault: The default linear sampling mode.
	BNNSLinearSamplingDefault BNNSLinearSampling = 0
	// BNNSLinearSamplingOffsetCorners: The offset corners sampling mode.
	BNNSLinearSamplingOffsetCorners BNNSLinearSampling = 10
	// BNNSLinearSamplingStrictAlignCorners: The strict align corners sampling mode.
	BNNSLinearSamplingStrictAlignCorners BNNSLinearSampling = 10
	// BNNSLinearSamplingUnalignCorners: The unalign corners sampling mode.
	BNNSLinearSamplingUnalignCorners BNNSLinearSampling = 10
)

func (e BNNSLinearSampling) String() string {
	switch e {
	case BNNSLinearSamplingAlignCorners:
		return "BNNSLinearSamplingAlignCorners"
	case BNNSLinearSamplingDefault:
		return "BNNSLinearSamplingDefault"
	default:
		return fmt.Sprintf("BNNSLinearSampling(%d)", e)
	}
}

type BNNSLossFunction uint

const (
	// BNNSLossFunctionCategoricalCrossEntropy: Performs categorical cross entropy computation between input prediction and labels.
	BNNSLossFunctionCategoricalCrossEntropy BNNSLossFunction = 8
	// BNNSLossFunctionCosineDistance: Performs cosine distance loss computation between input predictions and labels.
	BNNSLossFunctionCosineDistance BNNSLossFunction = 8
	// BNNSLossFunctionHinge: Performs Hinge loss computation between labels and unbounded zero-centered binary predictions.
	BNNSLossFunctionHinge BNNSLossFunction = 8
	// BNNSLossFunctionHuber: Huber loss computation between input logits and one-hot encoded labels.
	BNNSLossFunctionHuber BNNSLossFunction = 8
	// BNNSLossFunctionLog: Log loss computation between labels and predictions.
	BNNSLossFunctionLog BNNSLossFunction = 8
	// BNNSLossFunctionMeanAbsoluteError: Mean absolute error (MAE) computation between input prediction and labels.
	BNNSLossFunctionMeanAbsoluteError BNNSLossFunction = 8
	// BNNSLossFunctionMeanSquareError: Mean square error (MSE) computation between input logits and one-hot encoded labels.
	BNNSLossFunctionMeanSquareError BNNSLossFunction = 8
	// BNNSLossFunctionSigmoidCrossEntropy: Sigmoid activation on input logits, and independent computation of cross-entropy loss for each class.
	BNNSLossFunctionSigmoidCrossEntropy BNNSLossFunction = 8
	// BNNSLossFunctionSoftmaxCrossEntropy: Softmax activation on input logits, and computation of cross-entropy loss with one-hot encoded labels.
	BNNSLossFunctionSoftmaxCrossEntropy BNNSLossFunction = 0
	// BNNSLossFunctionYolo: You Only Look Once (YOLO) loss computation between prediction and ground truth labels.
	BNNSLossFunctionYolo BNNSLossFunction = 8
)

func (e BNNSLossFunction) String() string {
	switch e {
	case BNNSLossFunctionCategoricalCrossEntropy:
		return "BNNSLossFunctionCategoricalCrossEntropy"
	case BNNSLossFunctionSoftmaxCrossEntropy:
		return "BNNSLossFunctionSoftmaxCrossEntropy"
	default:
		return fmt.Sprintf("BNNSLossFunction(%d)", e)
	}
}

type BNNSLossReduction uint

const (
	// BNNSLossReductionMean: Sums the loss of all samples in the batch and divides by the number of samples.
	BNNSLossReductionMean BNNSLossReduction = 8
	// BNNSLossReductionNonZeroWeightMean: Sums the loss of all samples in the batch and divides by the number of non-zero weights.
	BNNSLossReductionNonZeroWeightMean BNNSLossReduction = 8
	// BNNSLossReductionNone: Returns the loss without any reduction.
	BNNSLossReductionNone BNNSLossReduction = 0
	// BNNSLossReductionSum: Sums the loss of all samples in the batch.
	BNNSLossReductionSum BNNSLossReduction = 8
	// BNNSLossReductionWeightedMean: Sums the loss of all samples in the batch and divides by the sum of all weights.
	BNNSLossReductionWeightedMean BNNSLossReduction = 8
)

func (e BNNSLossReduction) String() string {
	switch e {
	case BNNSLossReductionMean:
		return "BNNSLossReductionMean"
	case BNNSLossReductionNone:
		return "BNNSLossReductionNone"
	default:
		return fmt.Sprintf("BNNSLossReduction(%d)", e)
	}
}

type BNNSNDArrayFlagBackprop uint

const (
	// BNNSNDArrayFlagBackpropAccumulate: A flag that indicates backpropagation adds the value of the Jacobian to the elements of this n-dimensional array.
	BNNSNDArrayFlagBackpropAccumulate BNNSNDArrayFlagBackprop = 8
	// BNNSNDArrayFlagBackpropSet: A flag that indicates the elements of this n-dimensional array are overwritten by the Jacobian during backpropagation.
	BNNSNDArrayFlagBackpropSet BNNSNDArrayFlagBackprop = 0
)

func (e BNNSNDArrayFlagBackprop) String() string {
	switch e {
	case BNNSNDArrayFlagBackpropAccumulate:
		return "BNNSNDArrayFlagBackpropAccumulate"
	case BNNSNDArrayFlagBackpropSet:
		return "BNNSNDArrayFlagBackpropSet"
	default:
		return fmt.Sprintf("BNNSNDArrayFlagBackprop(%d)", e)
	}
}

type BNNSOptimizerClipping uint

const (
	// BNNSOptimizerClippingByGlobalNorm: A constant that specifes clipping to a maximum global Euclidean norm.
	BNNSOptimizerClippingByGlobalNorm BNNSOptimizerClipping = 9
	// BNNSOptimizerClippingByNorm: A constant that specifes clipping to a maximum Euclidean norm.
	BNNSOptimizerClippingByNorm BNNSOptimizerClipping = 9
	// BNNSOptimizerClippingByValue: A constant that specifes clipping to minimum and maximum values.
	BNNSOptimizerClippingByValue BNNSOptimizerClipping = 9
	// BNNSOptimizerClippingNone: A constant that specifes no clipping.
	BNNSOptimizerClippingNone BNNSOptimizerClipping = 0
)

func (e BNNSOptimizerClipping) String() string {
	switch e {
	case BNNSOptimizerClippingByGlobalNorm:
		return "BNNSOptimizerClippingByGlobalNorm"
	case BNNSOptimizerClippingNone:
		return "BNNSOptimizerClippingNone"
	default:
		return fmt.Sprintf("BNNSOptimizerClipping(%d)", e)
	}
}

type BNNSOptimizerFunction uint

const (
	// BNNSOptimizerFunctionAdam: An optimizer function that updates parameters according to the Adam algorithm.
	BNNSOptimizerFunctionAdam BNNSOptimizerFunction = 8
	// BNNSOptimizerFunctionAdamAMSGrad: An optimizer function that updates parameters according to the AMSGrad variant of the Adam algorithm.
	BNNSOptimizerFunctionAdamAMSGrad BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionAdamAMSGradWithClipping: An optimizer function that updates parameters according to the AMSGrad variant of the Adam algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamAMSGradWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionAdamW: An optimizer function that updates parameters according to the AdamW algorithm.
	BNNSOptimizerFunctionAdamW BNNSOptimizerFunction = 8
	// BNNSOptimizerFunctionAdamWAMSGrad: An optimizer function that updates parameters according to the AMSGrad variant of the AdamW algorithm.
	BNNSOptimizerFunctionAdamWAMSGrad BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionAdamWAMSGradWithClipping: An optimizer function that updates parameters according to the AMSGrad variant of the AdamW algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWAMSGradWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionAdamWWithClipping: An optimizer function that updates parameters according to the AdamW algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionAdamWithClipping: An optimizer function that updates parameters according to the Adam algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionRMSProp: An optimizer function that updates parameters according to the root mean square propagation (RMSProp) algorithm.
	BNNSOptimizerFunctionRMSProp BNNSOptimizerFunction = 8
	// BNNSOptimizerFunctionRMSPropWithClipping: An optimizer function that updates parameters according to the root mean square propagation (RMSProp) algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionRMSPropWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionSGDMomentum: An optimizer function that updates parameters according to the stochastic gradient descent (SGD) with momentum algorithm.
	BNNSOptimizerFunctionSGDMomentum BNNSOptimizerFunction = 0
	// BNNSOptimizerFunctionSGDMomentumWithClipping: An optimizer function that updates parameters according to the stochastic gradient descent (SGD) with momentum algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionSGDMomentumWithClipping BNNSOptimizerFunction = 9
)

func (e BNNSOptimizerFunction) String() string {
	switch e {
	case BNNSOptimizerFunctionAdam:
		return "BNNSOptimizerFunctionAdam"
	case BNNSOptimizerFunctionAdamAMSGrad:
		return "BNNSOptimizerFunctionAdamAMSGrad"
	case BNNSOptimizerFunctionSGDMomentum:
		return "BNNSOptimizerFunctionSGDMomentum"
	default:
		return fmt.Sprintf("BNNSOptimizerFunction(%d)", e)
	}
}

type BNNSOptimizerRegularization uint

const (
	// BNNSOptimizerRegularizationL1: A regularization function that applies L1 regularization.
	BNNSOptimizerRegularizationL1 BNNSOptimizerRegularization = 8
	// BNNSOptimizerRegularizationL2: A regularization function that applies L2 regularization.
	BNNSOptimizerRegularizationL2 BNNSOptimizerRegularization = 8
	// BNNSOptimizerRegularizationNone: A regularization function that adoesn’t apply any regularization.
	BNNSOptimizerRegularizationNone BNNSOptimizerRegularization = 0
)

func (e BNNSOptimizerRegularization) String() string {
	switch e {
	case BNNSOptimizerRegularizationL1:
		return "BNNSOptimizerRegularizationL1"
	case BNNSOptimizerRegularizationNone:
		return "BNNSOptimizerRegularizationNone"
	default:
		return fmt.Sprintf("BNNSOptimizerRegularization(%d)", e)
	}
}

type BNNSPaddingMode uint

const (
	// BNNSPaddingModeConstant: A constant that indicates that a padding operation fills the padded area with a specified constant.
	BNNSPaddingModeConstant BNNSPaddingMode = 0
	// BNNSPaddingModeReflect: A constant that indicates that a padding operation fills the padded area to form an odd-symmetric pattern.
	BNNSPaddingModeReflect BNNSPaddingMode = 8
	// BNNSPaddingModeSymmetric: A constant that indicates that a padding operation fills the padded area to form an even-symmetric pattern.
	BNNSPaddingModeSymmetric BNNSPaddingMode = 8
)

func (e BNNSPaddingMode) String() string {
	switch e {
	case BNNSPaddingModeConstant:
		return "BNNSPaddingModeConstant"
	case BNNSPaddingModeReflect:
		return "BNNSPaddingModeReflect"
	default:
		return fmt.Sprintf("BNNSPaddingMode(%d)", e)
	}
}

type BNNSPointerSpecifier uint

const (
	// BNNSPointerSpecifierAlpha: A constant that specifies the alpha pointer.
	BNNSPointerSpecifierAlpha BNNSPointerSpecifier = 0
	// BNNSPointerSpecifierBeta: A constant that specifies the beta pointer.
	BNNSPointerSpecifierBeta BNNSPointerSpecifier = 8
)

func (e BNNSPointerSpecifier) String() string {
	switch e {
	case BNNSPointerSpecifierAlpha:
		return "BNNSPointerSpecifierAlpha"
	case BNNSPointerSpecifierBeta:
		return "BNNSPointerSpecifierBeta"
	default:
		return fmt.Sprintf("BNNSPointerSpecifier(%d)", e)
	}
}

type BNNSPoolingFunction uint

const (
	// BNNSPoolingFunctionAverageCountExcludePadding: A function for pooling that computes the average of each element in the pooling kernel, excluding zero-padding.
	BNNSPoolingFunctionAverageCountExcludePadding BNNSPoolingFunction = 8
	// BNNSPoolingFunctionAverageCountIncludePadding: A function for pooling that computes the average of each element in the pooling kernel, including zero-padding.
	BNNSPoolingFunctionAverageCountIncludePadding BNNSPoolingFunction = 4
	// BNNSPoolingFunctionL2Norm: A function for pooling that computes the square root of the sum of squares of each element in the pooling kernel.
	BNNSPoolingFunctionL2Norm BNNSPoolingFunction = 8
	// BNNSPoolingFunctionMax: A function for pooling that computes the maximum of each element in the pooling kernel.
	BNNSPoolingFunctionMax BNNSPoolingFunction = 0
	// BNNSPoolingFunctionUnMax: A function for pooling that’s the partial inverse of max pooling and sets all nonmaximal values to zero.
	BNNSPoolingFunctionUnMax BNNSPoolingFunction = 8
	// Deprecated.
	BNNSPoolingFunctionAverage BNNSPoolingFunction = 8
)

func (e BNNSPoolingFunction) String() string {
	switch e {
	case BNNSPoolingFunctionAverageCountExcludePadding:
		return "BNNSPoolingFunctionAverageCountExcludePadding"
	case BNNSPoolingFunctionAverageCountIncludePadding:
		return "BNNSPoolingFunctionAverageCountIncludePadding"
	case BNNSPoolingFunctionMax:
		return "BNNSPoolingFunctionMax"
	default:
		return fmt.Sprintf("BNNSPoolingFunction(%d)", e)
	}
}

type BNNSQuantizerFunction uint

const (
	// BNNSQuantizerFunctionDequantize: A constant that specifes conversion to a higher precision.
	BNNSQuantizerFunctionDequantize BNNSQuantizerFunction = 9
	// BNNSQuantizerFunctionQuantize: A constant that specifes conversion to a lower precision.
	BNNSQuantizerFunctionQuantize BNNSQuantizerFunction = 0
)

func (e BNNSQuantizerFunction) String() string {
	switch e {
	case BNNSQuantizerFunctionDequantize:
		return "BNNSQuantizerFunctionDequantize"
	case BNNSQuantizerFunctionQuantize:
		return "BNNSQuantizerFunctionQuantize"
	default:
		return fmt.Sprintf("BNNSQuantizerFunction(%d)", e)
	}
}

type BNNSRandomGeneratorMethodAES uint

const (
	// BNNSRandomGeneratorMethodAES_CTR: A constant that specifes an implementation that’s based on the Advanced Encryption Standard (AES) hash of a counter.
	BNNSRandomGeneratorMethodAES_CTR BNNSRandomGeneratorMethodAES = 0
)

func (e BNNSRandomGeneratorMethodAES) String() string {
	switch e {
	case BNNSRandomGeneratorMethodAES_CTR:
		return "BNNSRandomGeneratorMethodAES_CTR"
	default:
		return fmt.Sprintf("BNNSRandomGeneratorMethodAES(%d)", e)
	}
}

type BNNSReduceFunction uint

const (
	// BNNSReduceFunctionAll: An alias of the logical AND reduction function.
	BNNSReduceFunctionAll BNNSReduceFunction = 8
	// BNNSReduceFunctionAny: An alias of the logical OR reduction function.
	BNNSReduceFunctionAny BNNSReduceFunction = 10
	// BNNSReduceFunctionArgMax: A reduction function that computes the index of the maximum value.
	BNNSReduceFunctionArgMax BNNSReduceFunction = 8
	// BNNSReduceFunctionArgMin: A reduction function that computes the index of the minimum value.
	BNNSReduceFunctionArgMin BNNSReduceFunction = 8
	// BNNSReduceFunctionL1Norm: A reduction function that computes the sum of the absolute value of each element.
	BNNSReduceFunctionL1Norm BNNSReduceFunction = 8
	// BNNSReduceFunctionL2Norm: A reduction function that computes the Euclidean norm.
	BNNSReduceFunctionL2Norm BNNSReduceFunction = 8
	BNNSReduceFunctionLogSum BNNSReduceFunction = 10
	// BNNSReduceFunctionLogSumExp: A reduction function that computes the logarithm of the sum of the exponentials of each element.
	BNNSReduceFunctionLogSumExp BNNSReduceFunction = 10
	// BNNSReduceFunctionLogicalAnd: A reduction function that reduces a tensor to true if all elements are true.
	BNNSReduceFunctionLogicalAnd BNNSReduceFunction = 8
	// BNNSReduceFunctionLogicalOr: A reduction function that reduces a tensor to true if any element is true.
	BNNSReduceFunctionLogicalOr BNNSReduceFunction = 8
	// BNNSReduceFunctionMax: A reduction function that computes the maximum value.
	BNNSReduceFunctionMax BNNSReduceFunction = 0
	// BNNSReduceFunctionMean: A reduction function that computes the mean value.
	BNNSReduceFunctionMean BNNSReduceFunction = 8
	// BNNSReduceFunctionMeanNonZero: A reduction function that computes the mean value of nonzero elements.
	BNNSReduceFunctionMeanNonZero BNNSReduceFunction = 8
	// BNNSReduceFunctionMin: A reduction function that computes the minimum value.
	BNNSReduceFunctionMin BNNSReduceFunction = 8
	// BNNSReduceFunctionNone: A reduction function that copies the input to the output.
	BNNSReduceFunctionNone BNNSReduceFunction = 10
	// BNNSReduceFunctionProduct: A reduction function that computes the product of all values.
	BNNSReduceFunctionProduct BNNSReduceFunction = 10
	// BNNSReduceFunctionSum: A reduction function that computes the sum of all values.
	BNNSReduceFunctionSum BNNSReduceFunction = 8
	// BNNSReduceFunctionSumLog: A reduction function that computes the sum of the natural logarithm of all values.
	BNNSReduceFunctionSumLog BNNSReduceFunction = 8
	// BNNSReduceFunctionSumSquare: A reduction function that computes the sum of the square of all values.
	BNNSReduceFunctionSumSquare BNNSReduceFunction = 8
)

func (e BNNSReduceFunction) String() string {
	switch e {
	case BNNSReduceFunctionAll:
		return "BNNSReduceFunctionAll"
	case BNNSReduceFunctionAny:
		return "BNNSReduceFunctionAny"
	case BNNSReduceFunctionMax:
		return "BNNSReduceFunctionMax"
	default:
		return fmt.Sprintf("BNNSReduceFunction(%d)", e)
	}
}

type BNNSRelationalOperator uint

const (
	// BNNSRelationalOperatorEqual: The operator that indicates the equal-to relationship.
	BNNSRelationalOperatorEqual BNNSRelationalOperator = 0
	// BNNSRelationalOperatorGreater: The operator that indicates the greater-than relationship.
	BNNSRelationalOperatorGreater BNNSRelationalOperator = 8
	// BNNSRelationalOperatorGreaterEqual: The operator that indicates the greater-than or equal-to relationship.
	BNNSRelationalOperatorGreaterEqual BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLess: The operator that indicates the less-than relationship.
	BNNSRelationalOperatorLess BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLessEqual: The operator that indicates the less-than or equal-to relationship.
	BNNSRelationalOperatorLessEqual BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalAND: The operator that indicates the logical AND relationship.
	BNNSRelationalOperatorLogicalAND BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalNAND: The operator that indicates the logical NAND relationship.
	BNNSRelationalOperatorLogicalNAND BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalNOR: The operator that indicates the logical NOR relationship.
	BNNSRelationalOperatorLogicalNOR BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalNOT: The operator that indicates the logical NOT relationship.
	BNNSRelationalOperatorLogicalNOT BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalOR: The operator that indicates the logical OR relationship.
	BNNSRelationalOperatorLogicalOR BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalXOR: The operator that indicates the logical XOR relationship.
	BNNSRelationalOperatorLogicalXOR BNNSRelationalOperator = 8
	// BNNSRelationalOperatorNotEqual: The operator that indicates the not-equal relationship.
	BNNSRelationalOperatorNotEqual BNNSRelationalOperator = 8
)

func (e BNNSRelationalOperator) String() string {
	switch e {
	case BNNSRelationalOperatorEqual:
		return "BNNSRelationalOperatorEqual"
	case BNNSRelationalOperatorGreater:
		return "BNNSRelationalOperatorGreater"
	default:
		return fmt.Sprintf("BNNSRelationalOperator(%d)", e)
	}
}

type BNNSSGDMomentum uint

const (
	// BNNSSGDMomentumVariant0: A constant that indicates SGD momentum variant 0.
	BNNSSGDMomentumVariant0 BNNSSGDMomentum = 0
	// BNNSSGDMomentumVariant1: A constant that indicates SGD momentum variant 1.
	BNNSSGDMomentumVariant1 BNNSSGDMomentum = 8
	// BNNSSGDMomentumVariant2: A constant that indicates SGD momentum variant 2.
	BNNSSGDMomentumVariant2 BNNSSGDMomentum = 8
)

func (e BNNSSGDMomentum) String() string {
	switch e {
	case BNNSSGDMomentumVariant0:
		return "BNNSSGDMomentumVariant0"
	case BNNSSGDMomentumVariant1:
		return "BNNSSGDMomentumVariant1"
	default:
		return fmt.Sprintf("BNNSSGDMomentum(%d)", e)
	}
}

type BNNSShuffleType uint

const (
	BNNSShuffleTypeDepthToSpaceNCHW BNNSShuffleType = 10
	// BNNSShuffleTypePixelShuffleNCHW: The pixel shuffle for the NCHW (batch, channels, height, width) format, equivalent to depth-to-space in Column Row Depth (CRD) mode.
	BNNSShuffleTypePixelShuffleNCHW BNNSShuffleType = 0
	// BNNSShuffleTypePixelUnshuffleNCHW: The pixel unshuffle for the NCHW (batch, channels, height, width) format, equivalent to space-to-depth in Column Row Depth (CRD) mode.
	BNNSShuffleTypePixelUnshuffleNCHW BNNSShuffleType = 10
	BNNSShuffleTypeSpaceToDepthNCHW   BNNSShuffleType = 10
)

func (e BNNSShuffleType) String() string {
	switch e {
	case BNNSShuffleTypeDepthToSpaceNCHW:
		return "BNNSShuffleTypeDepthToSpaceNCHW"
	case BNNSShuffleTypePixelShuffleNCHW:
		return "BNNSShuffleTypePixelShuffleNCHW"
	default:
		return fmt.Sprintf("BNNSShuffleType(%d)", e)
	}
}

type BNNSSparsityType uint

const (
	BNNSSparsityTypeUnstructured BNNSSparsityType = 0
)

func (e BNNSSparsityType) String() string {
	switch e {
	case BNNSSparsityTypeUnstructured:
		return "BNNSSparsityTypeUnstructured"
	default:
		return fmt.Sprintf("BNNSSparsityType(%d)", e)
	}
}

type BNNSTargetSystem uint

const (
	BNNSTargetSystemGeneric BNNSTargetSystem = 0
)

func (e BNNSTargetSystem) String() string {
	switch e {
	case BNNSTargetSystemGeneric:
		return "BNNSTargetSystemGeneric"
	default:
		return fmt.Sprintf("BNNSTargetSystem(%d)", e)
	}
}

type Bnns uint

const (
	// BnnsArithmetic: An arithmetic filter.
	BnnsArithmetic Bnns = 9
	// BNNSBatchNorm: A batch normalization filter.
	BNNSBatchNorm Bnns = 8
	// BNNSCenterSizeHeightFirst: Specifies coordinates as corners with the order: height start, width start, height end, width end.
	BNNSCenterSizeHeightFirst Bnns = 10
	// BNNSCenterSizeWidthFirst: Specifies coordinates as corners with the order: width start, height start, width end, height end.
	BNNSCenterSizeWidthFirst Bnns = 10
	// BNNSConstant: A constant that doesn’t have a gradient.
	BNNSConstant Bnns = 0
	// BNNSConvolution: A convolution filter.
	BNNSConvolution Bnns = 0
	// BNNSCornersHeightFirst: Specifies coordinates as center and size with the order: height center, width center, height, width.
	BNNSCornersHeightFirst Bnns = 0
	// BNNSCornersWidthFirst: Specifies coordinates as center and size with the order: width center, height center, width, height.
	BNNSCornersWidthFirst Bnns = 10
	// BNNSFullyConnected: A fully connected filter.
	BNNSFullyConnected Bnns = 8
	// BNNSGroupNorm: A group normalization filter.
	BNNSGroupNorm Bnns = 8
	// BNNSInstanceNorm: An instance normalization filter.
	BNNSInstanceNorm Bnns = 8
	// BNNSLayerNorm: A layer normalization filter.
	BNNSLayerNorm Bnns = 8
	// BNNSParameter: A parameter that’s trainable, such as weights or bias.
	BNNSParameter Bnns = 8
	// BNNSQuantization: A quantization filter.
	BNNSQuantization Bnns = 8
	// BNNSSample: A sample such as input or output.
	BNNSSample Bnns = 8
	// BNNSTransposedConvolution: A transposed convolution filter.
	BNNSTransposedConvolution Bnns = 8
)

func (e Bnns) String() string {
	switch e {
	case BnnsArithmetic:
		return "BnnsArithmetic"
	case BNNSBatchNorm:
		return "BNNSBatchNorm"
	case BNNSCenterSizeHeightFirst:
		return "BNNSCenterSizeHeightFirst"
	case BNNSConstant:
		return "BNNSConstant"
	default:
		return fmt.Sprintf("Bnns(%d)", e)
	}
}

type Bnnsl2 uint

const (
	// BNNSL2Norm: A constant that represents the L2 norm.
	BNNSL2Norm Bnnsl2 = 0
)

func (e Bnnsl2) String() string {
	switch e {
	case BNNSL2Norm:
		return "BNNSL2Norm"
	default:
		return fmt.Sprintf("Bnnsl2(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/CBLAS_DIAG
type CBLAS_DIAG uint32

const (
	CblasNonUnit CBLAS_DIAG = 131
	CblasUnit    CBLAS_DIAG = 132
)

func (e CBLAS_DIAG) String() string {
	switch e {
	case CblasNonUnit:
		return "CblasNonUnit"
	case CblasUnit:
		return "CblasUnit"
	default:
		return fmt.Sprintf("CBLAS_DIAG(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/CBLAS_ORDER
type CBLAS_ORDER uint32

const (
	CblasColMajor CBLAS_ORDER = 102
	CblasRowMajor CBLAS_ORDER = 101
)

func (e CBLAS_ORDER) String() string {
	switch e {
	case CblasColMajor:
		return "CblasColMajor"
	case CblasRowMajor:
		return "CblasRowMajor"
	default:
		return fmt.Sprintf("CBLAS_ORDER(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/CBLAS_SIDE
type CBLAS_SIDE uint32

const (
	CblasLeft  CBLAS_SIDE = 141
	CblasRight CBLAS_SIDE = 142
)

func (e CBLAS_SIDE) String() string {
	switch e {
	case CblasLeft:
		return "CblasLeft"
	case CblasRight:
		return "CblasRight"
	default:
		return fmt.Sprintf("CBLAS_SIDE(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/CBLAS_TRANSPOSE
type CBLAS_TRANSPOSE uint32

const (
	AtlasConj      CBLAS_TRANSPOSE = 0
	CblasConjTrans CBLAS_TRANSPOSE = 113
	CblasNoTrans   CBLAS_TRANSPOSE = 111
	CblasTrans     CBLAS_TRANSPOSE = 112
)

func (e CBLAS_TRANSPOSE) String() string {
	switch e {
	case AtlasConj:
		return "AtlasConj"
	case CblasConjTrans:
		return "CblasConjTrans"
	case CblasNoTrans:
		return "CblasNoTrans"
	case CblasTrans:
		return "CblasTrans"
	default:
		return fmt.Sprintf("CBLAS_TRANSPOSE(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/CBLAS_UPLO
type CBLAS_UPLO uint32

const (
	CblasLower CBLAS_UPLO = 122
	CblasUpper CBLAS_UPLO = 121
)

func (e CBLAS_UPLO) String() string {
	switch e {
	case CblasLower:
		return "CblasLower"
	case CblasUpper:
		return "CblasUpper"
	default:
		return fmt.Sprintf("CBLAS_UPLO(%d)", e)
	}
}

type Fft int

const (
	// FFT_FORWARD: Forward FFT.
	FFT_FORWARD Fft = 0
	// FFT_INVERSE: Inverse FFT.
	FFT_INVERSE Fft = -1
	FFT_RADIX2  Fft = 0
	FFT_RADIX3  Fft = 1
	FFT_RADIX5  Fft = 2
)

func (e Fft) String() string {
	switch e {
	case FFT_FORWARD:
		return "FFT_FORWARD"
	case FFT_INVERSE:
		return "FFT_INVERSE"
	case FFT_RADIX3:
		return "FFT_RADIX3"
	case FFT_RADIX5:
		return "FFT_RADIX5"
	default:
		return fmt.Sprintf("Fft(%d)", e)
	}
}

type KFFT uint

const (
	KFFTRadix2 KFFT = 0
	KFFTRadix3 KFFT = 1
	KFFTRadix5 KFFT = 2
)

func (e KFFT) String() string {
	switch e {
	case KFFTRadix2:
		return "KFFTRadix2"
	case KFFTRadix3:
		return "KFFTRadix3"
	case KFFTRadix5:
		return "KFFTRadix5"
	default:
		return fmt.Sprintf("KFFT(%d)", e)
	}
}

type KFFTDirection int

const (
	KFFTDirection_Forward KFFTDirection = 0
	KFFTDirection_Inverse KFFTDirection = -1
)

func (e KFFTDirection) String() string {
	switch e {
	case KFFTDirection_Forward:
		return "KFFTDirection_Forward"
	case KFFTDirection_Inverse:
		return "KFFTDirection_Inverse"
	default:
		return fmt.Sprintf("KFFTDirection(%d)", e)
	}
}

type KRotate0DegreesClockwise uint

const (
	// KRotate0DegreesClockwiseValue: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesClockwiseValue KRotate0DegreesClockwise = 0
	// KRotate0DegreesCounterClockwise: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate180DegreesClockwise: A constant that specifies rotation by 180° clockwise.
	KRotate180DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate180DegreesCounterClockwise: A constant that specifies rotation by 180° counterclockwise.
	KRotate180DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate270DegreesClockwise: A constant that specifies rotation by 270° clockwise.
	KRotate270DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate270DegreesCounterClockwise: A constant that specifies rotation by 270° counterclockwise.
	KRotate270DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate90DegreesClockwise: A constant that specifies rotation by 90° clockwise.
	KRotate90DegreesClockwise KRotate0DegreesClockwise = 0
	// KRotate90DegreesCounterClockwise: A constant that specifies rotation by 90° counterclockwise.
	KRotate90DegreesCounterClockwise KRotate0DegreesClockwise = 0
)

func (e KRotate0DegreesClockwise) String() string {
	switch e {
	case KRotate0DegreesClockwiseValue:
		return "KRotate0DegreesClockwiseValue"
	default:
		return fmt.Sprintf("KRotate0DegreesClockwise(%d)", e)
	}
}

type Kv uint

const (
	// KvImage420Yp8_Cb8_Cr8: Any y420 or f420 (planar component Y’CbCr 8-bit 4:2:0) buffer.
	KvImage420Yp8_Cb8_Cr8 Kv = 0
	// KvImage420Yp8_CbCr8: Any 420v or 420f (biplanar component Y’CbCr 8-bit 4:2:0, video-range) buffer.
	KvImage420Yp8_CbCr8 Kv = 0
	// KvImage422CbYpCrYp16: Any v216 (component Y’CbCr 10,12,14,16-bit 4:2:2) buffer.
	KvImage422CbYpCrYp16 Kv = 0
	// KvImage422CbYpCrYp8: Any 2vuy (component Y’CbCr 8-bit 4:2:2) buffer.
	KvImage422CbYpCrYp8 Kv = 0
	// KvImage422CbYpCrYp8_AA8: Any a2vy (first plane: video-range component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1; second plane: alpha 8-bit) buffer.
	KvImage422CbYpCrYp8_AA8 Kv = 0
	// KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10: Any v210 (component Y’CbCr 10-bit 4:2:2) buffer.
	KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10 Kv = 0
	// KvImage422YpCbYpCr8: Any yuvs or yuvf (component Y’CbCr 8-bit 4:2:2, ordered Y’0 Cb Y’1 Cr) buffer.
	KvImage422YpCbYpCr8 Kv = 0
	// KvImage444AYpCbCr16: Any y416 (component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr) buffer.
	KvImage444AYpCbCr16 Kv = 0
	// KvImage444AYpCbCr8: Any r408 or y408 (component Y’CbCrA 8-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr) buffer.
	KvImage444AYpCbCr8 Kv = 0
	// KvImage444CbYpCrA8: Any v408 (component Y’CbCrA 8-bit 4:4:4:4) buffer.
	KvImage444CbYpCrA8 Kv = 0
	// KvImage444CrYpCb10: Any v410 (component Y’CbCr 10-bit 4:4:4) buffer.
	KvImage444CrYpCb10 Kv = 0
	// KvImage444CrYpCb8: Any v308 (component Y’CbCr 8-bit 4:4:4) buffer.
	KvImage444CrYpCb8 Kv = 0
)

func (e Kv) String() string {
	switch e {
	case KvImage420Yp8_Cb8_Cr8:
		return "KvImage420Yp8_Cb8_Cr8"
	default:
		return fmt.Sprintf("Kv(%d)", e)
	}
}

type KvImage uint

const (
	// KvImageBackgroundColorFill: A flag that uses the background color for missing pixels.
	KvImageBackgroundColorFill KvImage = 0
	// KvImageBufferSizeMismatch: The function requires the source and destination buffers to have the same height and the same width, but they do not.
	KvImageBufferSizeMismatch KvImage = 0
	KvImageColorSyncIsAbsent  KvImage = 0
	// KvImageCopyInPlace: A flag that copies the value of the edge pixel in the source to the destination.
	KvImageCopyInPlace       KvImage = 0
	KvImageCoreVideoIsAbsent KvImage = 0
	// KvImageDoNotClamp: A flag that disables clamping in some conversions to floating-point formats.
	KvImageDoNotClamp KvImage = 0
	// KvImageDoNotTile: A flag that disables vImage internal tiling routines.
	KvImageDoNotTile KvImage = 0
	// KvImageEdgeExtend: A flag that extends the edges of the image infinitely.
	KvImageEdgeExtend KvImage = 0
	// KvImageFullInterpolation: Full linear interpolation.
	KvImageFullInterpolation KvImage = 0
	// KvImageGetTempBufferSize: A flag that returns the minimum temporary buffer size for the operation, given the parameters provided.
	KvImageGetTempBufferSize KvImage = 0
	// KvImageHDRContent: A flag that uses HDR-aware methods.
	KvImageHDRContent KvImage = 0
	// KvImageHalfInterpolation: Partial linear interpolation.
	KvImageHalfInterpolation KvImage = 0
	// KvImageHighQualityResampling: A flag that uses a higher-quality, slower resampling filter for geometry operations.
	KvImageHighQualityResampling KvImage = 0
	// KvImageInternalError: A serious error occured inside vImage, which prevented vImage from continuing.
	KvImageInternalError        KvImage = 0
	KvImageInvalidCVImageFormat KvImage = 0
	// KvImageInvalidEdgeStyle: The edge style specified is invalid.
	KvImageInvalidEdgeStyle   KvImage = 0
	KvImageInvalidImageFormat KvImage = 0
	KvImageInvalidImageObject KvImage = 0
	// KvImageInvalidKernelSize: Either the kernel height, the kernel width, or both, are even.
	KvImageInvalidKernelSize KvImage = 0
	// KvImageInvalidOffset_X: The `srcOffsetToROI_X` parameter that specifies the left edge of the region of interest is greater than the width of the source image.
	KvImageInvalidOffset_X KvImage = 0
	// KvImageInvalidOffset_Y: The `srcOffsetToROI_Y` parameter that specifies the top edge of the region of interest is greater than the height of the source image.
	KvImageInvalidOffset_Y KvImage = 0
	// KvImageInvalidParameter: Invalid parameter.
	KvImageInvalidParameter KvImage = 0
	KvImageInvalidRowBytes  KvImage = 0
	// KvImageLeaveAlphaUnchanged: A flag that restricts the operation to red, green, and blue channels only.
	KvImageLeaveAlphaUnchanged KvImage = 0
	// KvImageMemoryAllocationError: An attempt to allocate memory failed.
	KvImageMemoryAllocationError KvImage = 0
	// KvImageNoAllocate: A flag that prevents vImage from allocating additional storage.
	KvImageNoAllocate KvImage = 0
	// KvImageNoError: The vImage function completed without error.
	KvImageNoError KvImage = 0
	// KvImageNoFlags: A flag that sets the behavior to the default.
	KvImageNoFlags KvImage = 0
	// KvImageNoInterpolation: Nearest neighbor interpolation.
	KvImageNoInterpolation KvImage = 0
	// KvImageNullPointerArgument: A pointer parameter is [NULL] and it must not be.
	KvImageNullPointerArgument         KvImage = 0
	KvImageOutOfPlaceOperationRequired KvImage = 0
	// KvImagePrintDiagnosticsToConsole: A flag that prints a debug message if the operation fails.
	KvImagePrintDiagnosticsToConsole KvImage = 0
	// KvImageRoiLargerThanInputBuffer: The region of interest, as specified by the `srcOffsetToROI_X` and `srcOffsetToROI_Y` parameters and the height and width of the destination buffer, extends beyond the bottom edge or right edge of the source buffer.
	KvImageRoiLargerThanInputBuffer KvImage = 0
	// KvImageTruncateKernel: A flag that uses only the part of the kernel that overlaps the image.
	KvImageTruncateKernel KvImage = 0
	// KvImageUnknownFlagsBit: The flag is not recognized.
	KvImageUnknownFlagsBit KvImage = 0
	// KvImageUnsupportedConversion: Some lower level conversion APIs only support conversion among a sparse matrix of image formats.
	KvImageUnsupportedConversion KvImage = 0
	// KvImageUseFP16Accumulator: A flag that specifies vImage uses faster but lower-precision internal arithmetic for floating-point 16-bit operations.
	KvImageUseFP16Accumulator KvImage = 0
)

func (e KvImage) String() string {
	switch e {
	case KvImageBackgroundColorFill:
		return "KvImageBackgroundColorFill"
	default:
		return fmt.Sprintf("KvImage(%d)", e)
	}
}

type KvImageARG uint

const (
	// KvImageARGB16Q12: Any 8-bit four-channel interleaved buffer.
	KvImageARGB16Q12 KvImageARG = 0
	// KvImageARGB16U: Any 16-bit unsigned, four-channel interleaved buffer.
	KvImageARGB16U KvImageARG = 0
	// KvImageARGB8888: Any 16-bit signed fixed-point, four-channel interleaved buffer.
	KvImageARGB8888 KvImageARG = 0
)

func (e KvImageARG) String() string {
	switch e {
	case KvImageARGB16Q12:
		return "KvImageARGB16Q12"
	default:
		return fmt.Sprintf("KvImageARG(%d)", e)
	}
}

type KvImageBufferTypeCode uint

const (
	// KvImageBufferTypeCode_Alpha: The buffer contains the alpha channel.
	KvImageBufferTypeCode_Alpha KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CGFormat: The buffer contains data describable as a vImage Core Graphics image format as a single buffer.
	KvImageBufferTypeCode_CGFormat KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Black: If the image has a CMYK color model, the buffer contains the black channel.
	KvImageBufferTypeCode_CMYK_Black KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Cyan: If the image has a CMYK color model, the buffer contains the cyan channel.
	KvImageBufferTypeCode_CMYK_Cyan KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Magenta: If the image has a CMYK color model, the buffer contains the magenta channel.
	KvImageBufferTypeCode_CMYK_Magenta KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CMYK_Yellow: If the image has a CMYK color model, the buffer contains the yellow channel.
	KvImageBufferTypeCode_CMYK_Yellow KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_CVPixelBuffer_YCbCr: The buffer contains luminance and both chroma channels interleaved according to the vImageConstCVImageFormat image type.
	KvImageBufferTypeCode_CVPixelBuffer_YCbCr KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Cb: The buffer contains the blue chrominance channel.
	KvImageBufferTypeCode_Cb KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Chroma: The buffer contains both chrominance channels, interleaved.
	KvImageBufferTypeCode_Chroma KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Chunky: The buffer contains chunky data not describable as a vImage Core Graphics image format.
	KvImageBufferTypeCode_Chunky              KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel1  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel10 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel11 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel12 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel13 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel14 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel15 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel16 KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel2  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel3  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel4  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel5  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel6  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel7  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel8  KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_ColorSpaceChannel9  KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Cr: The buffer contains the red chrominance channel.
	KvImageBufferTypeCode_Cr KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_EndOfList: End of list marker.
	KvImageBufferTypeCode_EndOfList KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Indexed: The buffer contains data in an indexed colorspace.
	KvImageBufferTypeCode_Indexed KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_A: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_A KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_B: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_B KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_LAB_L: If the image has a LAB color model, the buffer contains the  channel.
	KvImageBufferTypeCode_LAB_L KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Luminance: The buffer contains only luminance data.
	KvImageBufferTypeCode_Luminance KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Monochrome: The buffer contains a single color channel.
	KvImageBufferTypeCode_Monochrome KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Blue: If the image has a RGB color model, the buffer contains the blue channel.
	KvImageBufferTypeCode_RGB_Blue KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Green: If the image has a RGB color model, the buffer contains the green channel.
	KvImageBufferTypeCode_RGB_Green KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_RGB_Red: If the image has a RGB color model, the buffer contains the red channel.
	KvImageBufferTypeCode_RGB_Red           KvImageBufferTypeCode = 0
	KvImageBufferTypeCode_UniqueFormatCount KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_X: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_X KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_Y: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_Y KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_XYZ_Z: If the image has a XYZ color model, the buffer contains the  channel.
	KvImageBufferTypeCode_XYZ_Z KvImageBufferTypeCode = 0
)

func (e KvImageBufferTypeCode) String() string {
	switch e {
	case KvImageBufferTypeCode_Alpha:
		return "KvImageBufferTypeCode_Alpha"
	default:
		return fmt.Sprintf("KvImageBufferTypeCode(%d)", e)
	}
}

type KvImageCVImageFormat uint

const (
	// KvImageCVImageFormat_AlphaIsOneHint: A hint that indicates the alpha channel is opaque.
	KvImageCVImageFormat_AlphaIsOneHint KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ChromaSiting: An error code that indicates the chroma siting information is absent.
	KvImageCVImageFormat_ChromaSiting KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ColorSpace: An error code that indicates the image’s color space is missing.
	KvImageCVImageFormat_ColorSpace KvImageCVImageFormat = 0
	// KvImageCVImageFormat_ConversionMatrix: An error code that indicates the required conversion matrix is absent.
	KvImageCVImageFormat_ConversionMatrix KvImageCVImageFormat = 0
	// KvImageCVImageFormat_NoError: An error code that indicates the conversion completed without error.
	KvImageCVImageFormat_NoError KvImageCVImageFormat = 0
	// KvImageCVImageFormat_VideoChannelDescription: An error code that indicates the range and clipping information is missing.
	KvImageCVImageFormat_VideoChannelDescription KvImageCVImageFormat = 0
)

func (e KvImageCVImageFormat) String() string {
	switch e {
	case KvImageCVImageFormat_AlphaIsOneHint:
		return "KvImageCVImageFormat_AlphaIsOneHint"
	default:
		return fmt.Sprintf("KvImageCVImageFormat(%d)", e)
	}
}

type KvImageConvert uint

const (
	// KvImageConvert_DitherAtkinson: A constant that indicates the conversion will add Atkinson dithering to the image.
	KvImageConvert_DitherAtkinson KvImageConvert = 4
	// KvImageConvert_DitherFloydSteinberg: A constant that indicates the conversion will add Floyd-Steinberg dithering to the image.
	KvImageConvert_DitherFloydSteinberg KvImageConvert = 3
	// KvImageConvert_DitherNone: A constant that indicates the conversion will not apply dithering.
	KvImageConvert_DitherNone KvImageConvert = 0
	// KvImageConvert_DitherOrdered: A constant that indicates the conversion will add randomized, pre-computed blue noise to the image.
	KvImageConvert_DitherOrdered KvImageConvert = 1
	// KvImageConvert_DitherOrderedReproducible: A constant that indicates the conversion will add reproducible, pre-computed blue noise to the image.
	KvImageConvert_DitherOrderedReproducible KvImageConvert = 2
	// KvImageConvert_OrderedGaussianBlue: A constant that indicates the conversion will distribute the noise according to a Gaussian distribution.
	KvImageConvert_OrderedGaussianBlue   KvImageConvert = 0
	KvImageConvert_OrderedNoiseShapeMask KvImageConvert = 0
	// KvImageConvert_OrderedUniformBlue: A constant that indicates the conversion will distribute the noise uniformly.
	KvImageConvert_OrderedUniformBlue KvImageConvert = 268435456
)

func (e KvImageConvert) String() string {
	switch e {
	case KvImageConvert_DitherAtkinson:
		return "KvImageConvert_DitherAtkinson"
	case KvImageConvert_DitherFloydSteinberg:
		return "KvImageConvert_DitherFloydSteinberg"
	case KvImageConvert_DitherNone:
		return "KvImageConvert_DitherNone"
	case KvImageConvert_DitherOrdered:
		return "KvImageConvert_DitherOrdered"
	case KvImageConvert_DitherOrderedReproducible:
		return "KvImageConvert_DitherOrderedReproducible"
	case KvImageConvert_OrderedUniformBlue:
		return "KvImageConvert_OrderedUniformBlue"
	default:
		return fmt.Sprintf("KvImageConvert(%d)", e)
	}
}

type KvImageGamma uint

const (
	// KvImageGamma_11_over_5_half_precision: A half-precision calculation using a gamma value of 11/5 or 2.2.
	KvImageGamma_11_over_5_half_precision KvImageGamma = 0
	// KvImageGamma_11_over_9_half_precision: A half-precision calculation using a gamma value of 11/9 or (11/5)/(9/5).
	KvImageGamma_11_over_9_half_precision KvImageGamma = 0
	// KvImageGamma_5_over_11_half_precision: A half-precision calculation using a gamma value of 5/11 or 1/2.2.
	KvImageGamma_5_over_11_half_precision KvImageGamma = 0
	// KvImageGamma_5_over_9_half_precision: A half-precision calculation using a gamma value of 5/9 or 1/1.8.
	KvImageGamma_5_over_9_half_precision KvImageGamma = 0
	// KvImageGamma_9_over_11_half_precision: A half-precision calculation using a gamma value of 9/11 or (9/5)/(11/5).
	KvImageGamma_9_over_11_half_precision KvImageGamma = 0
	// KvImageGamma_9_over_5_half_precision: A half-precision calculation using a gamma value of 9/5 or 1.8.
	KvImageGamma_9_over_5_half_precision KvImageGamma = 0
	// KvImageGamma_BT709_forward_half_precision: The ITU-R BT.709 standard.
	KvImageGamma_BT709_forward_half_precision KvImageGamma = 0
	// KvImageGamma_BT709_reverse_half_precision: The ITU-R BT.709 standard reverse.
	KvImageGamma_BT709_reverse_half_precision KvImageGamma = 0
	// KvImageGamma_UseGammaValue: A user-defined gamma value with full-precision calculation.
	KvImageGamma_UseGammaValue KvImageGamma = 0
	// KvImageGamma_UseGammaValue_half_precision: A user-defined gamma value with half-precision calculation.
	KvImageGamma_UseGammaValue_half_precision KvImageGamma = 0
	// KvImageGamma_sRGB_forward_half_precision: A half-precision calculation using the sRGB standard gamma value of 2.2.
	KvImageGamma_sRGB_forward_half_precision KvImageGamma = 0
	// KvImageGamma_sRGB_reverse_half_precision: A half-precision calculation using the sRGB standard gamma value of 1/2.2.
	KvImageGamma_sRGB_reverse_half_precision KvImageGamma = 0
)

func (e KvImageGamma) String() string {
	switch e {
	case KvImageGamma_11_over_5_half_precision:
		return "KvImageGamma_11_over_5_half_precision"
	default:
		return fmt.Sprintf("KvImageGamma(%d)", e)
	}
}

type KvImageInterpolation uint

const (
	// KvImageInterpolationLinear: Linear interpoation
	KvImageInterpolationLinear KvImageInterpolation = 0
	// KvImageInterpolationNearest: Nearest neigborhood
	KvImageInterpolationNearest KvImageInterpolation = 0
)

func (e KvImageInterpolation) String() string {
	switch e {
	case KvImageInterpolationLinear:
		return "KvImageInterpolationLinear"
	default:
		return fmt.Sprintf("KvImageInterpolation(%d)", e)
	}
}

type KvImageMDTableHint uint

const (
	// KvImageMDTableHint_16Q12: A table for transforming 16Q12 data.
	KvImageMDTableHint_16Q12 KvImageMDTableHint = 0
	// KvImageMDTableHint_Float: A table for transforming floating-point data.
	KvImageMDTableHint_Float KvImageMDTableHint = 0
)

func (e KvImageMDTableHint) String() string {
	switch e {
	case KvImageMDTableHint_16Q12:
		return "KvImageMDTableHint_16Q12"
	default:
		return fmt.Sprintf("KvImageMDTableHint(%d)", e)
	}
}

type KvImageMatrixType uint

const (
	KvImageMatrixType_ARGBToYpCbCrMatrix KvImageMatrixType = 0
	KvImageMatrixType_None               KvImageMatrixType = 0
)

func (e KvImageMatrixType) String() string {
	switch e {
	case KvImageMatrixType_ARGBToYpCbCrMatrix:
		return "KvImageMatrixType_ARGBToYpCbCrMatrix"
	default:
		return fmt.Sprintf("KvImageMatrixType(%d)", e)
	}
}

type KvimagePNGFilterValue uint

const (
	// KvImage_PNG_FILTER_VALUE_AVG: A filter that predicts a pixel value from the average of the pixels to the left and above the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_AVG KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_NONE: No filtering.
	KvImage_PNG_FILTER_VALUE_NONE KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_PAETH: A filter that predicts a pixel value by applying a linear function to the pixels located to the left, above, and to the upper-left of the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_PAETH KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_SUB: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located to the left.
	KvImage_PNG_FILTER_VALUE_SUB KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_UP: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located above.
	KvImage_PNG_FILTER_VALUE_UP KvimagePNGFilterValue = 0
)

func (e KvimagePNGFilterValue) String() string {
	switch e {
	case KvImage_PNG_FILTER_VALUE_AVG:
		return "KvImage_PNG_FILTER_VALUE_AVG"
	default:
		return fmt.Sprintf("KvimagePNGFilterValue(%d)", e)
	}
}

type Quadrature int

const (
	// QUADRATURE_ALLOC_ERROR: A constant that indicates that memory allocation failed.
	QUADRATURE_ALLOC_ERROR Quadrature = -3
	// QUADRATURE_ERROR: A constant that indicates that a generic error occurred.
	QUADRATURE_ERROR Quadrature = -1
	// QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR: A constant that indicates bad integrand behaviour, or that an excessive roundoff error occurred.
	QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR Quadrature = -102
	// QUADRATURE_INTEGRATE_MAX_EVAL_ERROR: A constant that indicates that the requested accuracy limit could not be reached.
	QUADRATURE_INTEGRATE_MAX_EVAL_ERROR Quadrature = -101
	// QUADRATURE_INTERNAL_ERROR: A constant that indicates that an internal error occurred.
	QUADRATURE_INTERNAL_ERROR Quadrature = -99
	// QUADRATURE_INVALID_ARG_ERROR: A constant that indicates that an invalid argument was passed to the operation.
	QUADRATURE_INVALID_ARG_ERROR Quadrature = -2
	// QUADRATURE_SUCCESS: A constant that indicates that the Quadrature operation was successful.
	QUADRATURE_SUCCESS Quadrature = 0
)

func (e Quadrature) String() string {
	switch e {
	case QUADRATURE_ALLOC_ERROR:
		return "QUADRATURE_ALLOC_ERROR"
	case QUADRATURE_ERROR:
		return "QUADRATURE_ERROR"
	case QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR:
		return "QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR"
	case QUADRATURE_INTEGRATE_MAX_EVAL_ERROR:
		return "QUADRATURE_INTEGRATE_MAX_EVAL_ERROR"
	case QUADRATURE_INTERNAL_ERROR:
		return "QUADRATURE_INTERNAL_ERROR"
	case QUADRATURE_INVALID_ARG_ERROR:
		return "QUADRATURE_INVALID_ARG_ERROR"
	case QUADRATURE_SUCCESS:
		return "QUADRATURE_SUCCESS"
	default:
		return fmt.Sprintf("Quadrature(%d)", e)
	}
}

type QuadratureIntegrateQ uint

const (
	// QUADRATURE_INTEGRATE_QAG: A constant that specifies a simple globally adaptive integrator.
	QUADRATURE_INTEGRATE_QAG QuadratureIntegrateQ = 1
	// QUADRATURE_INTEGRATE_QAGS: A constant that specifies global adaptive quadrature.
	QUADRATURE_INTEGRATE_QAGS QuadratureIntegrateQ = 2
	// QUADRATURE_INTEGRATE_QNG: A constant that specifies a simple non-adaptive automatic integrator.
	QUADRATURE_INTEGRATE_QNG QuadratureIntegrateQ = 0
)

func (e QuadratureIntegrateQ) String() string {
	switch e {
	case QUADRATURE_INTEGRATE_QAG:
		return "QUADRATURE_INTEGRATE_QAG"
	case QUADRATURE_INTEGRATE_QAGS:
		return "QUADRATURE_INTEGRATE_QAGS"
	case QUADRATURE_INTEGRATE_QNG:
		return "QUADRATURE_INTEGRATE_QNG"
	default:
		return fmt.Sprintf("QuadratureIntegrateQ(%d)", e)
	}
}

type Sparse int

const (
	// SPARSE_CANNOT_SET_PROPERTY: A property was set after values were inserted into the matrix.
	SPARSE_CANNOT_SET_PROPERTY Sparse = -1001
	// SPARSE_ILLEGAL_PARAMETER: Operation was not completed because one or more of the arguments had an illegal value.
	SPARSE_ILLEGAL_PARAMETER Sparse = -1000
	// SPARSE_LOWER_SYMMETRIC: A symmetric matrix with values derived from the lower triangle.
	SPARSE_LOWER_SYMMETRIC Sparse = 8
	// SPARSE_LOWER_TRIANGULAR: A lower triangular matrix.
	SPARSE_LOWER_TRIANGULAR Sparse = 2
	// SPARSE_SUCCESS: Operation was a success.
	SPARSE_SUCCESS Sparse = 0
	// SPARSE_SYSTEM_ERROR: An internal error has occured, such as non enough memory.
	SPARSE_SYSTEM_ERROR Sparse = -1002
	// SPARSE_UPPER_SYMMETRIC: A symmetric matrix with values derived from the upper triangle.
	SPARSE_UPPER_SYMMETRIC Sparse = 4
	// SPARSE_UPPER_TRIANGULAR: An upper triangular matrix.
	SPARSE_UPPER_TRIANGULAR Sparse = 1
	// SparseFactorizationFailed: The factorization failed due to a numerical issue.
	SparseFactorizationFailed Sparse = -1
	// SparseHermitian: A flag to describe the type of matrix represented.
	SparseHermitian Sparse = 5
	// SparseInternalError: The factorization encountered an internal error, such as failing to allocate memory.
	SparseInternalError Sparse = -3
	// SparseLowerTriangle: A constant that specifies the lower triangle.
	SparseLowerTriangle Sparse = 5
	// SparseMatrixIsSingular: The factorization aborted because the matrix is singular.
	SparseMatrixIsSingular Sparse = -2
	// SparseOrdinary: An unsymmetric sparse matrix without special structure.
	SparseOrdinary Sparse = 0
	// SparseParameterError: An error in a user-supplied parameter.
	SparseParameterError Sparse = -4
	// SparseStatusOK: The factorization was successful.
	SparseStatusOK Sparse = 0
	// SparseStatusReleased: The system freed the factorization object.
	SparseStatusReleased Sparse = -2147483647
	// SparseSymmetric: A symmetric sparse matrix.
	SparseSymmetric Sparse = 5
	// SparseTriangular: A triangular sparse matrix with a nonunit diagonal.
	SparseTriangular Sparse = 5
	// SparseUnitTriangular: A triangular sparse matrix with a unit diagonal.
	SparseUnitTriangular Sparse = 5
	// SparseUpperTriangle: A constant that specifies the upper triangle.
	SparseUpperTriangle Sparse = 0
)

func (e Sparse) String() string {
	switch e {
	case SPARSE_CANNOT_SET_PROPERTY:
		return "SPARSE_CANNOT_SET_PROPERTY"
	case SPARSE_ILLEGAL_PARAMETER:
		return "SPARSE_ILLEGAL_PARAMETER"
	case SPARSE_LOWER_SYMMETRIC:
		return "SPARSE_LOWER_SYMMETRIC"
	case SPARSE_LOWER_TRIANGULAR:
		return "SPARSE_LOWER_TRIANGULAR"
	case SPARSE_SUCCESS:
		return "SPARSE_SUCCESS"
	case SPARSE_SYSTEM_ERROR:
		return "SPARSE_SYSTEM_ERROR"
	case SPARSE_UPPER_SYMMETRIC:
		return "SPARSE_UPPER_SYMMETRIC"
	case SPARSE_UPPER_TRIANGULAR:
		return "SPARSE_UPPER_TRIANGULAR"
	case SparseFactorizationFailed:
		return "SparseFactorizationFailed"
	case SparseHermitian:
		return "SparseHermitian"
	case SparseInternalError:
		return "SparseInternalError"
	case SparseMatrixIsSingular:
		return "SparseMatrixIsSingular"
	case SparseParameterError:
		return "SparseParameterError"
	case SparseStatusReleased:
		return "SparseStatusReleased"
	default:
		return fmt.Sprintf("Sparse(%d)", e)
	}
}

type SparseDefault uint

const (
	// SparseDefaultControl: A flag that indicates default values.
	SparseDefaultControl SparseDefault = 0
)

func (e SparseDefault) String() string {
	switch e {
	case SparseDefaultControl:
		return "SparseDefaultControl"
	default:
		return fmt.Sprintf("SparseDefault(%d)", e)
	}
}

type SparseFactorization uint

const (
	// SparseFactorizationCholesky: A constant that represents Cholesky () factorization.
	SparseFactorizationCholesky SparseFactorization = 0
	// SparseFactorizationCholeskyAtA: A constant that represents  factorization without storing .
	SparseFactorizationCholeskyAtA SparseFactorization = 5
	// SparseFactorizationLDLT: A constant that represents the default  factorization.
	SparseFactorizationLDLT SparseFactorization = 5
	// SparseFactorizationLDLTSBK: A constant that represents  factorization with Supernode-Bunch-Kaufman and static pivoting.
	SparseFactorizationLDLTSBK SparseFactorization = 5
	// SparseFactorizationLDLTTPP: A constant that represents  factorization with full-threshold partial pivoting.
	SparseFactorizationLDLTTPP SparseFactorization = 5
	// SparseFactorizationLDLTUnpivoted: A constant that represents Cholesky-like  factorization with only one-by-one pivots and no pivoting.
	SparseFactorizationLDLTUnpivoted SparseFactorization = 5
	// SparseFactorizationLU: Default LU factorization, currently LU with TPP.
	SparseFactorizationLU SparseFactorization = 5
	// SparseFactorizationLUSPP: LU factorization with partial pivoting restricted to within supernodes only.
	SparseFactorizationLUSPP SparseFactorization = 12
	// SparseFactorizationLUTPP: LU factorization with threshold partial pivoting.
	SparseFactorizationLUTPP SparseFactorization = 12
	// SparseFactorizationLUUnpivoted: LU factorization with no numerical pivoting.
	SparseFactorizationLUUnpivoted SparseFactorization = 12
	// SparseFactorizationQR: A constant that represents QR factorization.
	SparseFactorizationQR SparseFactorization = 5
)

func (e SparseFactorization) String() string {
	switch e {
	case SparseFactorizationCholesky:
		return "SparseFactorizationCholesky"
	case SparseFactorizationCholeskyAtA:
		return "SparseFactorizationCholeskyAtA"
	case SparseFactorizationLUSPP:
		return "SparseFactorizationLUSPP"
	default:
		return fmt.Sprintf("SparseFactorization(%d)", e)
	}
}

type SparseIterative int

const (
	// SparseIterativeConverged: A status that indicates the convergence of all solutions.
	SparseIterativeConverged SparseIterative = 0
	// SparseIterativeIllConditioned: A status that indicates the operation determines the problem is sufficiently ill-conditioned that convergence is unlikely.
	SparseIterativeIllConditioned SparseIterative = -2
	// SparseIterativeInternalError: A status that indicates an internal failure.
	SparseIterativeInternalError SparseIterative = -99
	// SparseIterativeMaxIterations: A status that indicates a failure to converge one or more solutions in the maximum number of iterations.
	SparseIterativeMaxIterations SparseIterative = 1
	// SparseIterativeParameterError: A status that indicates an error with one or more parameters.
	SparseIterativeParameterError SparseIterative = -1
)

func (e SparseIterative) String() string {
	switch e {
	case SparseIterativeConverged:
		return "SparseIterativeConverged"
	case SparseIterativeIllConditioned:
		return "SparseIterativeIllConditioned"
	case SparseIterativeInternalError:
		return "SparseIterativeInternalError"
	case SparseIterativeMaxIterations:
		return "SparseIterativeMaxIterations"
	case SparseIterativeParameterError:
		return "SparseIterativeParameterError"
	default:
		return fmt.Sprintf("SparseIterative(%d)", e)
	}
}

type SparseLSMRCT uint

const (
	// SparseLSMRCTDefault: The default convergence test.
	SparseLSMRCTDefault SparseLSMRCT = 0
	// SparseLSMRCTFongSaunders: Fong and Saunder’s original convergence test.
	SparseLSMRCTFongSaunders SparseLSMRCT = 1
)

func (e SparseLSMRCT) String() string {
	switch e {
	case SparseLSMRCTDefault:
		return "SparseLSMRCTDefault"
	case SparseLSMRCTFongSaunders:
		return "SparseLSMRCTFongSaunders"
	default:
		return fmt.Sprintf("SparseLSMRCT(%d)", e)
	}
}

type SparseNorm uint

const (
	// SPARSE_NORM_INF: Norm Inf
	SPARSE_NORM_INF SparseNorm = 175
	// SPARSE_NORM_ONE: Norm One
	SPARSE_NORM_ONE SparseNorm = 171
	// SPARSE_NORM_R1: Norm R1
	SPARSE_NORM_R1 SparseNorm = 179
	// SPARSE_NORM_TWO: Norm Two
	SPARSE_NORM_TWO SparseNorm = 173
)

func (e SparseNorm) String() string {
	switch e {
	case SPARSE_NORM_INF:
		return "SPARSE_NORM_INF"
	case SPARSE_NORM_ONE:
		return "SPARSE_NORM_ONE"
	case SPARSE_NORM_R1:
		return "SPARSE_NORM_R1"
	case SPARSE_NORM_TWO:
		return "SPARSE_NORM_TWO"
	default:
		return fmt.Sprintf("SparseNorm(%d)", e)
	}
}

type SparseOrder uint

const (
	// SparseOrderAMD: Approximate minimum degree (AMD) ordering.
	SparseOrderAMD SparseOrder = 5
	// SparseOrderCOLAMD: The column AMD ordering for .
	SparseOrderCOLAMD SparseOrder = 5
	// SparseOrderDefault: The default ordering.
	SparseOrderDefault SparseOrder = 0
	// SparseOrderMTMetis: Specifies type of fill-reducing ordering.
	SparseOrderMTMetis SparseOrder = 5
	// SparseOrderMetis: METIS nested dissection ordering.
	SparseOrderMetis SparseOrder = 5
	// SparseOrderUser: The user-supplied ordering, or identity if the order parameter is null.
	SparseOrderUser SparseOrder = 5
)

func (e SparseOrder) String() string {
	switch e {
	case SparseOrderAMD:
		return "SparseOrderAMD"
	case SparseOrderDefault:
		return "SparseOrderDefault"
	default:
		return fmt.Sprintf("SparseOrder(%d)", e)
	}
}

type SparsePreconditioner uint

const (
	// SparsePreconditionerDiagScaling: A diagonal scaling preconditioner.
	SparsePreconditionerDiagScaling SparsePreconditioner = 3
	// SparsePreconditionerDiagonal: A Jacobi preconditioner.
	SparsePreconditionerDiagonal SparsePreconditioner = 2
	// SparsePreconditionerNone: No preconditioner.
	SparsePreconditionerNone SparsePreconditioner = 0
	// SparsePreconditionerUser: A user-provided preconditioner.
	SparsePreconditionerUser SparsePreconditioner = 1
)

func (e SparsePreconditioner) String() string {
	switch e {
	case SparsePreconditionerDiagScaling:
		return "SparsePreconditionerDiagScaling"
	case SparsePreconditionerDiagonal:
		return "SparsePreconditionerDiagonal"
	case SparsePreconditionerNone:
		return "SparsePreconditionerNone"
	case SparsePreconditionerUser:
		return "SparsePreconditionerUser"
	default:
		return fmt.Sprintf("SparsePreconditioner(%d)", e)
	}
}

type SparseScaling uint

const (
	// SparseScalingDefault: Default scaling.
	SparseScalingDefault SparseScaling = 0
	// SparseScalingEquilibriationInf: The norm equilibration scaling using infinity norm.
	SparseScalingEquilibriationInf SparseScaling = 5
	// SparseScalingHungarianScalingAndOrdering: Scaling and ordering using the Hungarian algorithm.
	SparseScalingHungarianScalingAndOrdering SparseScaling = 27
	// SparseScalingHungarianScalingOnly: Scaling using the Hungarian algorithm.
	SparseScalingHungarianScalingOnly SparseScaling = 5
	// SparseScalingUser: User scaling.
	SparseScalingUser SparseScaling = 5
)

func (e SparseScaling) String() string {
	switch e {
	case SparseScalingDefault:
		return "SparseScalingDefault"
	case SparseScalingEquilibriationInf:
		return "SparseScalingEquilibriationInf"
	case SparseScalingHungarianScalingAndOrdering:
		return "SparseScalingHungarianScalingAndOrdering"
	default:
		return fmt.Sprintf("SparseScaling(%d)", e)
	}
}

type SparseSubfactor uint

const (
	// SparseSubfactorD: A  factor subfactor that’s valid for ` `only.
	SparseSubfactorD SparseSubfactor = 5
	// SparseSubfactorInvalid: An invalid subfactor that indicates the requested type is incompatible with the supplied factorization or the system has destroyed it.
	SparseSubfactorInvalid SparseSubfactor = 0
	// SparseSubfactorL: An  factor subfactor that’s valid for Cholesky and  only.
	SparseSubfactorL SparseSubfactor = 5
	// SparseSubfactorP: A permutation subfactor that’s valid for all factorization types.
	SparseSubfactorP SparseSubfactor = 5
	// SparseSubfactorPLPS: A half-solve subfactor that’s valid for Cholesky and  only.
	SparseSubfactorPLPS SparseSubfactor = 5
	// SparseSubfactorQ: A  factor subfactor that’s valid for QR only.
	SparseSubfactorQ SparseSubfactor = 5
	// SparseSubfactorR: An  factor subfactor that’s valid for QR and Cholesky  only.
	SparseSubfactorR SparseSubfactor = 5
	// SparseSubfactorRP: A half-solve subfactor that’s valid for QR and Cholesky  only.
	SparseSubfactorRP SparseSubfactor = 5
	// SparseSubfactorS: A diagonal scaling subfactor that’s valid for Cholesky and  only.
	SparseSubfactorS SparseSubfactor = 5
	// SparseSubfactorSc: Types of sub-factor object.
	SparseSubfactorSc SparseSubfactor = 12
	// SparseSubfactorSr: Types of sub-factor object.
	SparseSubfactorSr SparseSubfactor = 5
)

func (e SparseSubfactor) String() string {
	switch e {
	case SparseSubfactorD:
		return "SparseSubfactorD"
	case SparseSubfactorInvalid:
		return "SparseSubfactorInvalid"
	case SparseSubfactorSc:
		return "SparseSubfactorSc"
	default:
		return fmt.Sprintf("SparseSubfactor(%d)", e)
	}
}

type SparseUpdatePartial uint

const (
	// SparseUpdatePartialRefactor: Low-rank update algorithm selector
	SparseUpdatePartialRefactor SparseUpdatePartial = 0
)

func (e SparseUpdatePartial) String() string {
	switch e {
	case SparseUpdatePartialRefactor:
		return "SparseUpdatePartialRefactor"
	default:
		return fmt.Sprintf("SparseUpdatePartial(%d)", e)
	}
}

type SparseVariant uint

const (
	// SparseVariantDQGMRES: A constant that specifies the DQGMRES variant.
	SparseVariantDQGMRES SparseVariant = 0
	// SparseVariantFGMRES: A constant that specifies the flexible GMRES variant.
	SparseVariantFGMRES SparseVariant = 2
	// SparseVariantGMRES: A constant that specifies the standard restarted GMRES variant.
	SparseVariantGMRES SparseVariant = 1
)

func (e SparseVariant) String() string {
	switch e {
	case SparseVariantDQGMRES:
		return "SparseVariantDQGMRES"
	case SparseVariantFGMRES:
		return "SparseVariantFGMRES"
	case SparseVariantGMRES:
		return "SparseVariantGMRES"
	default:
		return fmt.Sprintf("SparseVariant(%d)", e)
	}
}

type VdspHa uint

const (
	// VDSP_HALF_WINDOW: Specifies that the window should only contain the bottom half of the values (`0` to `(N+1)/2`).
	VDSP_HALF_WINDOW VdspHa = 1
	// VDSP_HANN_DENORM: Specifies a denormalized Hann window.
	VDSP_HANN_DENORM VdspHa = 0
	// VDSP_HANN_NORM: Specifies a normalized Hann window
	VDSP_HANN_NORM VdspHa = 2
)

func (e VdspHa) String() string {
	switch e {
	case VDSP_HALF_WINDOW:
		return "VDSP_HALF_WINDOW"
	case VDSP_HANN_DENORM:
		return "VDSP_HANN_DENORM"
	case VDSP_HANN_NORM:
		return "VDSP_HANN_NORM"
	default:
		return fmt.Sprintf("VdspHa(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type
type VDSP_DCT_Type int32

const (
	// VDSP_DCT_II: A constant that specifies a type II discrete cosine transform.
	VDSP_DCT_II VDSP_DCT_Type = 2
	// VDSP_DCT_III: A constant that specifies a type III discrete cosine transform.
	VDSP_DCT_III VDSP_DCT_Type = 3
	// VDSP_DCT_IV: A constant that specifies a type IV discrete cosine transform.
	VDSP_DCT_IV VDSP_DCT_Type = 4
)

func (e VDSP_DCT_Type) String() string {
	switch e {
	case VDSP_DCT_II:
		return "VDSP_DCT_II"
	case VDSP_DCT_III:
		return "VDSP_DCT_III"
	case VDSP_DCT_IV:
		return "VDSP_DCT_IV"
	default:
		return fmt.Sprintf("VDSP_DCT_Type(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Direction
type VDSP_DFT_Direction int32

const (
	// VDSP_DFT_FORWARD: A constant that specifies a forward transform.
	VDSP_DFT_FORWARD VDSP_DFT_Direction = 0
	// VDSP_DFT_INVERSE: A constant that specifies an inverse transform.
	VDSP_DFT_INVERSE VDSP_DFT_Direction = -1
)

func (e VDSP_DFT_Direction) String() string {
	switch e {
	case VDSP_DFT_FORWARD:
		return "VDSP_DFT_FORWARD"
	case VDSP_DFT_INVERSE:
		return "VDSP_DFT_INVERSE"
	default:
		return fmt.Sprintf("VDSP_DFT_Direction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_RealtoComplex
type VDSP_DFT_RealtoComplex int

const (
	VDSP_DFT_Interleaved_ComplextoComplex VDSP_DFT_RealtoComplex = 0
	VDSP_DFT_Interleaved_RealtoComplex    VDSP_DFT_RealtoComplex = 1
)

func (e VDSP_DFT_RealtoComplex) String() string {
	switch e {
	case VDSP_DFT_Interleaved_ComplextoComplex:
		return "VDSP_DFT_Interleaved_ComplextoComplex"
	case VDSP_DFT_Interleaved_RealtoComplex:
		return "VDSP_DFT_Interleaved_RealtoComplex"
	default:
		return fmt.Sprintf("VDSP_DFT_RealtoComplex(%d)", e)
	}
}
