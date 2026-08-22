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

// See: https://developer.apple.com/documentation/Accelerate/BNNSActivationFunction
type BNNSActivationFunction uint32

const (
	// BNNSActivationFunctionAbs: An activation function that returns the absolute value of its input.
	BNNSActivationFunctionAbs BNNSActivationFunction = 6
	// BNNSActivationFunctionCELU: An activation function that evaluates the continuously differentiable exponential linear units (CELU) on its input.
	BNNSActivationFunctionCELU BNNSActivationFunction = 24
	// BNNSActivationFunctionClamp: An activation function that returns its input clamped to the specified range.
	BNNSActivationFunctionClamp BNNSActivationFunction = 8
	// BNNSActivationFunctionClampedLeakyRectifiedLinear: An activation function that returns its input clamped to beta when that is greater than or equal to zero, otherwise it returns its input multiplied by alpha clamped to beta.
	BNNSActivationFunctionClampedLeakyRectifiedLinear BNNSActivationFunction = 19
	// BNNSActivationFunctionELU: An activation function that evaluates the exponential linear units (ELU) on its input.
	BNNSActivationFunctionELU  BNNSActivationFunction = 18
	BNNSActivationFunctionErf  BNNSActivationFunction = 33
	BNNSActivationFunctionGELU BNNSActivationFunction = 34
	// BNNSActivationFunctionGELUApproximation: An activation function that evaluates the Gaussian error linear units (GELU) approximation on its input.
	BNNSActivationFunctionGELUApproximation BNNSActivationFunction = 12
	// BNNSActivationFunctionGELUApproximation2: An activation function that provides a fast evaluation of the Gaussian error linear units (GELU) approximation on its input.
	BNNSActivationFunctionGELUApproximation2       BNNSActivationFunction = 30
	BNNSActivationFunctionGELUApproximationSigmoid BNNSActivationFunction = 35
	// BNNSActivationFunctionGumbel: An activation function that returns random numbers from the Gumbel distribution.
	BNNSActivationFunctionGumbel BNNSActivationFunction = 13
	// BNNSActivationFunctionGumbelMax: An activation function that returns random numbers from the Gumbel distribution.
	BNNSActivationFunctionGumbelMax BNNSActivationFunction = 14
	// BNNSActivationFunctionHardShrink: An activation function that returns zero when the absolute input is less than alpha, otherwise it returns its input.
	BNNSActivationFunctionHardShrink BNNSActivationFunction = 25
	// BNNSActivationFunctionHardSigmoid: An activation function that returns the hard sigmoid function of its input.
	BNNSActivationFunctionHardSigmoid BNNSActivationFunction = 15
	// BNNSActivationFunctionHardSwish: An activation function that returns the hard swish function of its input.
	BNNSActivationFunctionHardSwish BNNSActivationFunction = 30
	// BNNSActivationFunctionIdentity: An activation function that returns its input.
	BNNSActivationFunctionIdentity BNNSActivationFunction = 0
	// BNNSActivationFunctionIntegerLinearSaturate: An activation function that returns an arithmetic shift, preserving sign.
	BNNSActivationFunctionIntegerLinearSaturate BNNSActivationFunction = 9
	// BNNSActivationFunctionIntegerLinearSaturatePerChannel: An activation function that returns an arithmetic shift, preserving sign for each channel.
	BNNSActivationFunctionIntegerLinearSaturatePerChannel BNNSActivationFunction = 10
	// BNNSActivationFunctionLeakyRectifiedLinear: An activation function that returns its input when that is greater than or equal to zero, otherwise it returns its input multiplied by a specified value.
	BNNSActivationFunctionLeakyRectifiedLinear BNNSActivationFunction = 2
	// BNNSActivationFunctionLinear: An activation function that returns its input multiplied by a specified value.
	BNNSActivationFunctionLinear BNNSActivationFunction = 7
	// BNNSActivationFunctionLinearWithBias: An activation function that returns its input multiplied by a scale and added to a bias.
	BNNSActivationFunctionLinearWithBias BNNSActivationFunction = 20
	// BNNSActivationFunctionLogSigmoid: An activation function that returns the logarithm of the sigmoid function of its input.
	BNNSActivationFunctionLogSigmoid BNNSActivationFunction = 22
	// BNNSActivationFunctionLogSoftmax: An activation function that returns the logarithm of the softmax function of its input.
	BNNSActivationFunctionLogSoftmax BNNSActivationFunction = 21
	// BNNSActivationFunctionPReLUPerChannel: An activation function provides per-channel alpha values to Leaky Rectified Linear.
	BNNSActivationFunctionPReLUPerChannel BNNSActivationFunction = 29
	BNNSActivationFunctionReLU6           BNNSActivationFunction = 32
	// BNNSActivationFunctionRectifiedLinear: An activation function that returns its input when that is greater than or equal to zero, otherwise it returns zero.
	BNNSActivationFunctionRectifiedLinear BNNSActivationFunction = 1
	// BNNSActivationFunctionSELU: An activation function that evaluates the scaled exponential linear units (SELU) on its input.
	BNNSActivationFunctionSELU BNNSActivationFunction = 23
	// BNNSActivationFunctionScaledTanh: An activation function that returns the scaled hyperbolic tangent of its input.
	BNNSActivationFunctionScaledTanh BNNSActivationFunction = 5
	// BNNSActivationFunctionSiLU: An activation function that returns the sigmoid linear unit (SiLU) function of its input.
	BNNSActivationFunctionSiLU BNNSActivationFunction = 31
	// BNNSActivationFunctionSigmoid: An activation function that returns the sigmoid function of its input.
	BNNSActivationFunctionSigmoid BNNSActivationFunction = 3
	// BNNSActivationFunctionSoftShrink: An activation function that returns zero when the absolute input is less than alpha, otherwise it returns its input minus alpha.
	BNNSActivationFunctionSoftShrink BNNSActivationFunction = 26
	// BNNSActivationFunctionSoftmax: An activation function that returns the softmax function of its input.
	BNNSActivationFunctionSoftmax BNNSActivationFunction = 11
	// BNNSActivationFunctionSoftplus: An activation function that returns the softplus function of its input.
	BNNSActivationFunctionSoftplus BNNSActivationFunction = 16
	// BNNSActivationFunctionSoftsign: An activation function that returns the softsign function of its input.
	BNNSActivationFunctionSoftsign BNNSActivationFunction = 17
	// BNNSActivationFunctionTanh: An activation function that returns the hyperbolic tangent of its input.
	BNNSActivationFunctionTanh BNNSActivationFunction = 4
	// BNNSActivationFunctionTanhShrink: An activation function that returns its input minus the hyperbolic tangent of its input.
	BNNSActivationFunctionTanhShrink BNNSActivationFunction = 27
	// BNNSActivationFunctionThreshold: An activation function that returns beta if its input is less than a specified threshold, otherwise it returns its input.
	BNNSActivationFunctionThreshold BNNSActivationFunction = 28
)

func (e BNNSActivationFunction) String() string {
	switch e {
	case BNNSActivationFunctionAbs:
		return "BNNSActivationFunctionAbs"
	case BNNSActivationFunctionCELU:
		return "BNNSActivationFunctionCELU"
	case BNNSActivationFunctionClamp:
		return "BNNSActivationFunctionClamp"
	case BNNSActivationFunctionClampedLeakyRectifiedLinear:
		return "BNNSActivationFunctionClampedLeakyRectifiedLinear"
	case BNNSActivationFunctionELU:
		return "BNNSActivationFunctionELU"
	case BNNSActivationFunctionErf:
		return "BNNSActivationFunctionErf"
	case BNNSActivationFunctionGELU:
		return "BNNSActivationFunctionGELU"
	case BNNSActivationFunctionGELUApproximation:
		return "BNNSActivationFunctionGELUApproximation"
	case BNNSActivationFunctionGELUApproximation2:
		return "BNNSActivationFunctionGELUApproximation2"
	case BNNSActivationFunctionGELUApproximationSigmoid:
		return "BNNSActivationFunctionGELUApproximationSigmoid"
	case BNNSActivationFunctionGumbel:
		return "BNNSActivationFunctionGumbel"
	case BNNSActivationFunctionGumbelMax:
		return "BNNSActivationFunctionGumbelMax"
	case BNNSActivationFunctionHardShrink:
		return "BNNSActivationFunctionHardShrink"
	case BNNSActivationFunctionHardSigmoid:
		return "BNNSActivationFunctionHardSigmoid"
	case BNNSActivationFunctionIdentity:
		return "BNNSActivationFunctionIdentity"
	case BNNSActivationFunctionIntegerLinearSaturate:
		return "BNNSActivationFunctionIntegerLinearSaturate"
	case BNNSActivationFunctionIntegerLinearSaturatePerChannel:
		return "BNNSActivationFunctionIntegerLinearSaturatePerChannel"
	case BNNSActivationFunctionLeakyRectifiedLinear:
		return "BNNSActivationFunctionLeakyRectifiedLinear"
	case BNNSActivationFunctionLinear:
		return "BNNSActivationFunctionLinear"
	case BNNSActivationFunctionLinearWithBias:
		return "BNNSActivationFunctionLinearWithBias"
	case BNNSActivationFunctionLogSigmoid:
		return "BNNSActivationFunctionLogSigmoid"
	case BNNSActivationFunctionLogSoftmax:
		return "BNNSActivationFunctionLogSoftmax"
	case BNNSActivationFunctionPReLUPerChannel:
		return "BNNSActivationFunctionPReLUPerChannel"
	case BNNSActivationFunctionReLU6:
		return "BNNSActivationFunctionReLU6"
	case BNNSActivationFunctionRectifiedLinear:
		return "BNNSActivationFunctionRectifiedLinear"
	case BNNSActivationFunctionSELU:
		return "BNNSActivationFunctionSELU"
	case BNNSActivationFunctionScaledTanh:
		return "BNNSActivationFunctionScaledTanh"
	case BNNSActivationFunctionSiLU:
		return "BNNSActivationFunctionSiLU"
	case BNNSActivationFunctionSigmoid:
		return "BNNSActivationFunctionSigmoid"
	case BNNSActivationFunctionSoftShrink:
		return "BNNSActivationFunctionSoftShrink"
	case BNNSActivationFunctionSoftmax:
		return "BNNSActivationFunctionSoftmax"
	case BNNSActivationFunctionSoftplus:
		return "BNNSActivationFunctionSoftplus"
	case BNNSActivationFunctionSoftsign:
		return "BNNSActivationFunctionSoftsign"
	case BNNSActivationFunctionTanh:
		return "BNNSActivationFunctionTanh"
	case BNNSActivationFunctionTanhShrink:
		return "BNNSActivationFunctionTanhShrink"
	case BNNSActivationFunctionThreshold:
		return "BNNSActivationFunctionThreshold"
	default:
		return fmt.Sprintf("BNNSActivationFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticFunction
type BNNSArithmeticFunction uint32

const (
	// BNNSArithmeticAbs: An operation that calculates the element-wise absolute of its input.
	BNNSArithmeticAbs BNNSArithmeticFunction = 32
	// BNNSArithmeticAcos: An operation that calculates the element-wise inverse cosine of its input.
	BNNSArithmeticAcos BNNSArithmeticFunction = 13
	// BNNSArithmeticAcosh: An operation that calculates the element-wise inverse hyperbolic cosine of its input.
	BNNSArithmeticAcosh BNNSArithmeticFunction = 19
	// BNNSArithmeticAdd: An operation that calculates the element-wise sum of its two inputs.
	BNNSArithmeticAdd BNNSArithmeticFunction = 0
	// BNNSArithmeticAsin: An operation that calculates the element-wise inverse sine of its input.
	BNNSArithmeticAsin BNNSArithmeticFunction = 12
	// BNNSArithmeticAsinh: An operation that calculates the element-wise inverse hyperbolic sine of its input.
	BNNSArithmeticAsinh BNNSArithmeticFunction = 18
	// BNNSArithmeticAtan: An operation that calculates the element-wise inverse tangent of its input.
	BNNSArithmeticAtan BNNSArithmeticFunction = 14
	// BNNSArithmeticAtanh: An operation that calculates the element-wise inverse hyperbolic tangent of its input.
	BNNSArithmeticAtanh BNNSArithmeticFunction = 20
	// BNNSArithmeticCeil: An operation that calculates the element-wise ceiling of its input.
	BNNSArithmeticCeil BNNSArithmeticFunction = 6
	// BNNSArithmeticCos: An operation that calculates the element-wise cosine of its input.
	BNNSArithmeticCos BNNSArithmeticFunction = 10
	// BNNSArithmeticCosh: An operation that calculates the element-wise hyperbolic cosine of its input.
	BNNSArithmeticCosh BNNSArithmeticFunction = 16
	// BNNSArithmeticDivide: An operation that calculates the element-wise division of its two inputs.
	BNNSArithmeticDivide BNNSArithmeticFunction = 3
	// BNNSArithmeticDivideNoNaN: An operation that calculates the element-wise division of its two inputs and returns zero if the divisor is zero, even if the first input is NaN or infinity.
	BNNSArithmeticDivideNoNaN BNNSArithmeticFunction = 27
	// BNNSArithmeticErf: An operation that calculates the element-wise error function of its input.
	BNNSArithmeticErf BNNSArithmeticFunction = 40
	// BNNSArithmeticExp: An operation that calculates the element-wise result of e raised to the power of its input.
	BNNSArithmeticExp BNNSArithmeticFunction = 22
	// BNNSArithmeticExp2: An operation that calculates the element-wise result of 2 raised to the power of its input.
	BNNSArithmeticExp2 BNNSArithmeticFunction = 23
	// BNNSArithmeticFloor: An operation that calculates the element-wise floor of its input.
	BNNSArithmeticFloor BNNSArithmeticFunction = 7
	// BNNSArithmeticFloorDivide: An operation that calculates the element-wise floor division of its inputs.
	BNNSArithmeticFloorDivide BNNSArithmeticFunction = 37
	// BNNSArithmeticLog: An operation that calculates the element-wise natural logarithm of its input.
	BNNSArithmeticLog BNNSArithmeticFunction = 24
	// BNNSArithmeticLog2: An operation that calculates the element-wise base 2 logarithm of its input.
	BNNSArithmeticLog2 BNNSArithmeticFunction = 25
	// BNNSArithmeticMaximum: An operation that calculates the element-wise maximum of its two inputs.
	BNNSArithmeticMaximum BNNSArithmeticFunction = 30
	// BNNSArithmeticMinimum: An operation that calculates the element-wise minimum of its two inputs.
	BNNSArithmeticMinimum BNNSArithmeticFunction = 29
	// BNNSArithmeticMultiply: An operation that calculates the element-wise product of its two inputs.
	BNNSArithmeticMultiply BNNSArithmeticFunction = 2
	// BNNSArithmeticMultiplyAdd: An operation that calculates the element-wise fused multiply-add of its three inputs.
	BNNSArithmeticMultiplyAdd BNNSArithmeticFunction = 28
	// BNNSArithmeticMultiplyNoNaN: An operation that calculates the element-wise product of its two inputs and returns zero, even if the first input is NaN or infinity.
	BNNSArithmeticMultiplyNoNaN BNNSArithmeticFunction = 26
	// BNNSArithmeticNegate: An operation that calculates the element-wise negation of its input.
	BNNSArithmeticNegate BNNSArithmeticFunction = 34
	// BNNSArithmeticPow: An operation that calculates the element-wise first input raised to the power of its second input.
	BNNSArithmeticPow BNNSArithmeticFunction = 21
	// BNNSArithmeticReciprocal: An operation that calculates the element-wise reciprocal of its input.
	BNNSArithmeticReciprocal BNNSArithmeticFunction = 35
	// BNNSArithmeticReciprocalSquareRoot: An operation that calculates the element-wise reciprocal square root of its input.
	BNNSArithmeticReciprocalSquareRoot BNNSArithmeticFunction = 5
	// BNNSArithmeticRound: An operation that calculates the element-wise rounding of its input.
	BNNSArithmeticRound BNNSArithmeticFunction = 8
	// BNNSArithmeticSelect: An operation that selects elements from either its second or third input based on the corresponding value of its first input.
	BNNSArithmeticSelect BNNSArithmeticFunction = 31
	// BNNSArithmeticSign: An operation that calculates the element-wise sign of its input.
	BNNSArithmeticSign BNNSArithmeticFunction = 33
	// BNNSArithmeticSin: An operation that calculates the element-wise sine of its input.
	BNNSArithmeticSin BNNSArithmeticFunction = 9
	// BNNSArithmeticSinh: An operation that calculates the element-wise hyperbolic sine of its input.
	BNNSArithmeticSinh BNNSArithmeticFunction = 15
	// BNNSArithmeticSquare: An operation that calculates the element-wise square of its input.
	BNNSArithmeticSquare BNNSArithmeticFunction = 36
	// BNNSArithmeticSquareRoot: An operation that calculates the element-wise square root of its input.
	BNNSArithmeticSquareRoot BNNSArithmeticFunction = 4
	// BNNSArithmeticSubtract: An operation that calculates the element-wise difference of its two inputs.
	BNNSArithmeticSubtract BNNSArithmeticFunction = 1
	// BNNSArithmeticTan: An operation that calculates the element-wise tangent of its input.
	BNNSArithmeticTan BNNSArithmeticFunction = 11
	// BNNSArithmeticTanh: An operation that calculates the element-wise hyperbolic tangent of its input.
	BNNSArithmeticTanh BNNSArithmeticFunction = 17
	// BNNSArithmeticTruncDivide: An operation that calculates the element-wise truncated division of its inputs.
	BNNSArithmeticTruncDivide BNNSArithmeticFunction = 38
	// BNNSArithmeticTruncRemainder: An operation that calculates the element-wise remainder of truncated division of its inputs.
	BNNSArithmeticTruncRemainder BNNSArithmeticFunction = 39
)

func (e BNNSArithmeticFunction) String() string {
	switch e {
	case BNNSArithmeticAbs:
		return "BNNSArithmeticAbs"
	case BNNSArithmeticAcos:
		return "BNNSArithmeticAcos"
	case BNNSArithmeticAcosh:
		return "BNNSArithmeticAcosh"
	case BNNSArithmeticAdd:
		return "BNNSArithmeticAdd"
	case BNNSArithmeticAsin:
		return "BNNSArithmeticAsin"
	case BNNSArithmeticAsinh:
		return "BNNSArithmeticAsinh"
	case BNNSArithmeticAtan:
		return "BNNSArithmeticAtan"
	case BNNSArithmeticAtanh:
		return "BNNSArithmeticAtanh"
	case BNNSArithmeticCeil:
		return "BNNSArithmeticCeil"
	case BNNSArithmeticCos:
		return "BNNSArithmeticCos"
	case BNNSArithmeticCosh:
		return "BNNSArithmeticCosh"
	case BNNSArithmeticDivide:
		return "BNNSArithmeticDivide"
	case BNNSArithmeticDivideNoNaN:
		return "BNNSArithmeticDivideNoNaN"
	case BNNSArithmeticErf:
		return "BNNSArithmeticErf"
	case BNNSArithmeticExp:
		return "BNNSArithmeticExp"
	case BNNSArithmeticExp2:
		return "BNNSArithmeticExp2"
	case BNNSArithmeticFloor:
		return "BNNSArithmeticFloor"
	case BNNSArithmeticFloorDivide:
		return "BNNSArithmeticFloorDivide"
	case BNNSArithmeticLog:
		return "BNNSArithmeticLog"
	case BNNSArithmeticLog2:
		return "BNNSArithmeticLog2"
	case BNNSArithmeticMaximum:
		return "BNNSArithmeticMaximum"
	case BNNSArithmeticMinimum:
		return "BNNSArithmeticMinimum"
	case BNNSArithmeticMultiply:
		return "BNNSArithmeticMultiply"
	case BNNSArithmeticMultiplyAdd:
		return "BNNSArithmeticMultiplyAdd"
	case BNNSArithmeticMultiplyNoNaN:
		return "BNNSArithmeticMultiplyNoNaN"
	case BNNSArithmeticNegate:
		return "BNNSArithmeticNegate"
	case BNNSArithmeticPow:
		return "BNNSArithmeticPow"
	case BNNSArithmeticReciprocal:
		return "BNNSArithmeticReciprocal"
	case BNNSArithmeticReciprocalSquareRoot:
		return "BNNSArithmeticReciprocalSquareRoot"
	case BNNSArithmeticRound:
		return "BNNSArithmeticRound"
	case BNNSArithmeticSelect:
		return "BNNSArithmeticSelect"
	case BNNSArithmeticSign:
		return "BNNSArithmeticSign"
	case BNNSArithmeticSin:
		return "BNNSArithmeticSin"
	case BNNSArithmeticSinh:
		return "BNNSArithmeticSinh"
	case BNNSArithmeticSquare:
		return "BNNSArithmeticSquare"
	case BNNSArithmeticSquareRoot:
		return "BNNSArithmeticSquareRoot"
	case BNNSArithmeticSubtract:
		return "BNNSArithmeticSubtract"
	case BNNSArithmeticTan:
		return "BNNSArithmeticTan"
	case BNNSArithmeticTanh:
		return "BNNSArithmeticTanh"
	case BNNSArithmeticTruncDivide:
		return "BNNSArithmeticTruncDivide"
	case BNNSArithmeticTruncRemainder:
		return "BNNSArithmeticTruncRemainder"
	default:
		return fmt.Sprintf("BNNSArithmeticFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSBoxCoordinateMode
type BNNSBoxCoordinateMode uint32

const (
	// BNNSCenterSizeHeightFirst: Specifies coordinates as corners with the order: height start, width start, height end, width end.
	BNNSCenterSizeHeightFirst BNNSBoxCoordinateMode = 2
	// BNNSCenterSizeWidthFirst: Specifies coordinates as corners with the order: width start, height start, width end, height end.
	BNNSCenterSizeWidthFirst BNNSBoxCoordinateMode = 3
	// BNNSCornersHeightFirst: Specifies coordinates as center and size with the order: height center, width center, height, width.
	BNNSCornersHeightFirst BNNSBoxCoordinateMode = 0
	// BNNSCornersWidthFirst: Specifies coordinates as center and size with the order: width center, height center, width, height.
	BNNSCornersWidthFirst BNNSBoxCoordinateMode = 1
)

func (e BNNSBoxCoordinateMode) String() string {
	switch e {
	case BNNSCenterSizeHeightFirst:
		return "BNNSCenterSizeHeightFirst"
	case BNNSCenterSizeWidthFirst:
		return "BNNSCenterSizeWidthFirst"
	case BNNSCornersHeightFirst:
		return "BNNSCornersHeightFirst"
	case BNNSCornersWidthFirst:
		return "BNNSCornersWidthFirst"
	default:
		return fmt.Sprintf("BNNSBoxCoordinateMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSDataLayout
type BNNSDataLayout uint32

const (
	// BNNSDataLayout1DFirstMajor: A constant that represents a 1D first-major vector.
	BNNSDataLayout1DFirstMajor BNNSDataLayout = 0x18001
	// BNNSDataLayout1DLastMajor: A constant that represents a 1D last-major vector.
	BNNSDataLayout1DLastMajor BNNSDataLayout = 0x18000
	// BNNSDataLayout2DFirstMajor: A constant that represents a 2D first-major matrix.
	BNNSDataLayout2DFirstMajor BNNSDataLayout = 0x28001
	// BNNSDataLayout2DLastMajor: A constant that represents a 2D last-major matrix.
	BNNSDataLayout2DLastMajor BNNSDataLayout = 0x28000
	// BNNSDataLayout3DFirstMajor: A constant that represents a 3D first-major tensor.
	BNNSDataLayout3DFirstMajor BNNSDataLayout = 0x38001
	// BNNSDataLayout3DLastMajor: A constant that represents a 3D last-major tensor.
	BNNSDataLayout3DLastMajor BNNSDataLayout = 0x38000
	// BNNSDataLayout4DFirstMajor: A constant that represents a 4D first-major tensor.
	BNNSDataLayout4DFirstMajor BNNSDataLayout = 0x48001
	// BNNSDataLayout4DLastMajor: A constant that represents a 4D last-major tensor.
	BNNSDataLayout4DLastMajor BNNSDataLayout = 0x48000
	// BNNSDataLayout5DFirstMajor: A constant that represents a 5D first-major tensor.
	BNNSDataLayout5DFirstMajor BNNSDataLayout = 0x58001
	// BNNSDataLayout5DLastMajor: A constant that represents a 5D last-major tensor.
	BNNSDataLayout5DLastMajor BNNSDataLayout = 0x58000
	// BNNSDataLayout6DFirstMajor: A constant that represents a 6D first-major tensor.
	BNNSDataLayout6DFirstMajor BNNSDataLayout = 0x68001
	// BNNSDataLayout6DLastMajor: A constant that represents a 6D last-major tensor.
	BNNSDataLayout6DLastMajor BNNSDataLayout = 0x68000
	// BNNSDataLayout7DFirstMajor: A constant that represents a 7D first-major tensor.
	BNNSDataLayout7DFirstMajor BNNSDataLayout = 0x78001
	// BNNSDataLayout7DLastMajor: A constant that represents a 7D last-major tensor.
	BNNSDataLayout7DLastMajor BNNSDataLayout = 0x78000
	// BNNSDataLayout8DFirstMajor: A constant that represents a 8D first-major tensor.
	BNNSDataLayout8DFirstMajor BNNSDataLayout = 0x88001
	// BNNSDataLayout8DLastMajor: A constant that represents a 8D last-major tensor.
	BNNSDataLayout8DLastMajor BNNSDataLayout = 0x88000
	// BNNSDataLayoutColumnMajorMatrix: A constant that represents a 2D column-major matrix.
	BNNSDataLayoutColumnMajorMatrix BNNSDataLayout = 0x20001
	// BNNSDataLayoutConvolutionWeightsIOHrWr: A constant that represents a 4D array of rotated convolution weights.
	BNNSDataLayoutConvolutionWeightsIOHrWr BNNSDataLayout = 0x40002
	// BNNSDataLayoutConvolutionWeightsOIHW: A constant that represents a 4D array of convolution weights.
	BNNSDataLayoutConvolutionWeightsOIHW BNNSDataLayout = 0x40000
	// BNNSDataLayoutConvolutionWeightsOIHW_Pack32: A constant that represents a 4D array of packed convolution weights with 32-output channel packing and 128-byte array address alignment.
	BNNSDataLayoutConvolutionWeightsOIHW_Pack32 BNNSDataLayout = 0x40010
	// BNNSDataLayoutConvolutionWeightsOIHrWr: A constant that represents a 4D array of rotated convolution weights.
	BNNSDataLayoutConvolutionWeightsOIHrWr BNNSDataLayout = 0x40001
	BNNSDataLayoutFullyConnectedSparse     BNNSDataLayout = 0x21001
	// BNNSDataLayoutImageCHW: A constant that represents a 3D image stack.
	BNNSDataLayoutImageCHW BNNSDataLayout = 0x30000
	BNNSDataLayoutMHA_DHK  BNNSDataLayout = 0x30003
	// BNNSDataLayoutNSE: A constant that represents a 3D tensor with the size elements embedding dimension, sequence length, and batch size.
	BNNSDataLayoutNSE BNNSDataLayout = 0x30002
	// BNNSDataLayoutRowMajorMatrix: A constant that represents a 2D row-major matrix.
	BNNSDataLayoutRowMajorMatrix BNNSDataLayout = 0x20000
	// BNNSDataLayoutSNE: A constant that represents a 3D tensor with the size elements embedding dimension, batch size, and sequence length.
	BNNSDataLayoutSNE BNNSDataLayout = 0x30001
	// BNNSDataLayoutVector: A constant that represents a 1D vector.
	BNNSDataLayoutVector BNNSDataLayout = 0x10000
)

func (e BNNSDataLayout) String() string {
	switch e {
	case BNNSDataLayout1DFirstMajor:
		return "BNNSDataLayout1DFirstMajor"
	case BNNSDataLayout1DLastMajor:
		return "BNNSDataLayout1DLastMajor"
	case BNNSDataLayout2DFirstMajor:
		return "BNNSDataLayout2DFirstMajor"
	case BNNSDataLayout2DLastMajor:
		return "BNNSDataLayout2DLastMajor"
	case BNNSDataLayout3DFirstMajor:
		return "BNNSDataLayout3DFirstMajor"
	case BNNSDataLayout3DLastMajor:
		return "BNNSDataLayout3DLastMajor"
	case BNNSDataLayout4DFirstMajor:
		return "BNNSDataLayout4DFirstMajor"
	case BNNSDataLayout4DLastMajor:
		return "BNNSDataLayout4DLastMajor"
	case BNNSDataLayout5DFirstMajor:
		return "BNNSDataLayout5DFirstMajor"
	case BNNSDataLayout5DLastMajor:
		return "BNNSDataLayout5DLastMajor"
	case BNNSDataLayout6DFirstMajor:
		return "BNNSDataLayout6DFirstMajor"
	case BNNSDataLayout6DLastMajor:
		return "BNNSDataLayout6DLastMajor"
	case BNNSDataLayout7DFirstMajor:
		return "BNNSDataLayout7DFirstMajor"
	case BNNSDataLayout7DLastMajor:
		return "BNNSDataLayout7DLastMajor"
	case BNNSDataLayout8DFirstMajor:
		return "BNNSDataLayout8DFirstMajor"
	case BNNSDataLayout8DLastMajor:
		return "BNNSDataLayout8DLastMajor"
	case BNNSDataLayoutColumnMajorMatrix:
		return "BNNSDataLayoutColumnMajorMatrix"
	case BNNSDataLayoutConvolutionWeightsIOHrWr:
		return "BNNSDataLayoutConvolutionWeightsIOHrWr"
	case BNNSDataLayoutConvolutionWeightsOIHW:
		return "BNNSDataLayoutConvolutionWeightsOIHW"
	case BNNSDataLayoutConvolutionWeightsOIHW_Pack32:
		return "BNNSDataLayoutConvolutionWeightsOIHW_Pack32"
	case BNNSDataLayoutConvolutionWeightsOIHrWr:
		return "BNNSDataLayoutConvolutionWeightsOIHrWr"
	case BNNSDataLayoutFullyConnectedSparse:
		return "BNNSDataLayoutFullyConnectedSparse"
	case BNNSDataLayoutImageCHW:
		return "BNNSDataLayoutImageCHW"
	case BNNSDataLayoutMHA_DHK:
		return "BNNSDataLayoutMHA_DHK"
	case BNNSDataLayoutNSE:
		return "BNNSDataLayoutNSE"
	case BNNSDataLayoutRowMajorMatrix:
		return "BNNSDataLayoutRowMajorMatrix"
	case BNNSDataLayoutSNE:
		return "BNNSDataLayoutSNE"
	case BNNSDataLayoutVector:
		return "BNNSDataLayoutVector"
	default:
		return fmt.Sprintf("BNNSDataLayout(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSDataType
type BNNSDataType uint32

const (
	// BNNSDataTypeBFloat16: # Discussion
	BNNSDataTypeBFloat16 BNNSDataType = 98320
	// BNNSDataTypeBoolean: # Discussion
	BNNSDataTypeBoolean BNNSDataType = 1048584
	// BNNSDataTypeFloat16: # Discussion
	BNNSDataTypeFloat16 BNNSDataType = 65552
	// BNNSDataTypeFloat32: # Discussion
	BNNSDataTypeFloat32 BNNSDataType = 65568
	// BNNSDataTypeFloatBit: # Discussion
	BNNSDataTypeFloatBit BNNSDataType = 0x10000
	// BNNSDataTypeIndexed1: # Discussion
	BNNSDataTypeIndexed1 BNNSDataType = 524289
	// BNNSDataTypeIndexed2: # Discussion
	BNNSDataTypeIndexed2 BNNSDataType = 524290
	// BNNSDataTypeIndexed4: # Discussion
	BNNSDataTypeIndexed4 BNNSDataType = 524292
	// BNNSDataTypeIndexed8: # Discussion
	BNNSDataTypeIndexed8 BNNSDataType = 524296
	// BNNSDataTypeIndexedBit: # Discussion
	BNNSDataTypeIndexedBit BNNSDataType = 0x80000
	// BNNSDataTypeInt1: # Discussion
	BNNSDataTypeInt1 BNNSDataType = 131073
	// BNNSDataTypeInt16: # Discussion
	BNNSDataTypeInt16 BNNSDataType = 131088
	// BNNSDataTypeInt2: # Discussion
	BNNSDataTypeInt2 BNNSDataType = 131074
	// BNNSDataTypeInt32: # Discussion
	BNNSDataTypeInt32 BNNSDataType = 131104
	// BNNSDataTypeInt4: # Discussion
	BNNSDataTypeInt4 BNNSDataType = 131076
	// BNNSDataTypeInt64: # Discussion
	BNNSDataTypeInt64 BNNSDataType = 131136
	// BNNSDataTypeInt8: # Discussion
	BNNSDataTypeInt8 BNNSDataType = 131080
	// BNNSDataTypeIntBit: # Discussion
	BNNSDataTypeIntBit BNNSDataType = 0x20000
	// BNNSDataTypeMiscellaneousBit: # Discussion
	BNNSDataTypeMiscellaneousBit BNNSDataType = 0x100000
	// BNNSDataTypeUInt1: # Discussion
	BNNSDataTypeUInt1 BNNSDataType = 262145
	// BNNSDataTypeUInt16: # Discussion
	BNNSDataTypeUInt16 BNNSDataType = 262160
	// BNNSDataTypeUInt2: # Discussion
	BNNSDataTypeUInt2 BNNSDataType = 262146
	// BNNSDataTypeUInt3: # Discussion
	BNNSDataTypeUInt3 BNNSDataType = 262147
	// BNNSDataTypeUInt32: # Discussion
	BNNSDataTypeUInt32 BNNSDataType = 262176
	// BNNSDataTypeUInt4: # Discussion
	BNNSDataTypeUInt4 BNNSDataType = 262148
	// BNNSDataTypeUInt6: # Discussion
	BNNSDataTypeUInt6 BNNSDataType = 262150
	// BNNSDataTypeUInt64: # Discussion
	BNNSDataTypeUInt64 BNNSDataType = 262208
	// BNNSDataTypeUInt8: # Discussion
	BNNSDataTypeUInt8 BNNSDataType = 262152
	// BNNSDataTypeUIntBit: # Discussion
	BNNSDataTypeUIntBit BNNSDataType = 0x40000
)

func (e BNNSDataType) String() string {
	switch e {
	case BNNSDataTypeBFloat16:
		return "BNNSDataTypeBFloat16"
	case BNNSDataTypeBoolean:
		return "BNNSDataTypeBoolean"
	case BNNSDataTypeFloat16:
		return "BNNSDataTypeFloat16"
	case BNNSDataTypeFloat32:
		return "BNNSDataTypeFloat32"
	case BNNSDataTypeFloatBit:
		return "BNNSDataTypeFloatBit"
	case BNNSDataTypeIndexed1:
		return "BNNSDataTypeIndexed1"
	case BNNSDataTypeIndexed2:
		return "BNNSDataTypeIndexed2"
	case BNNSDataTypeIndexed4:
		return "BNNSDataTypeIndexed4"
	case BNNSDataTypeIndexed8:
		return "BNNSDataTypeIndexed8"
	case BNNSDataTypeIndexedBit:
		return "BNNSDataTypeIndexedBit"
	case BNNSDataTypeInt1:
		return "BNNSDataTypeInt1"
	case BNNSDataTypeInt16:
		return "BNNSDataTypeInt16"
	case BNNSDataTypeInt2:
		return "BNNSDataTypeInt2"
	case BNNSDataTypeInt32:
		return "BNNSDataTypeInt32"
	case BNNSDataTypeInt4:
		return "BNNSDataTypeInt4"
	case BNNSDataTypeInt64:
		return "BNNSDataTypeInt64"
	case BNNSDataTypeInt8:
		return "BNNSDataTypeInt8"
	case BNNSDataTypeIntBit:
		return "BNNSDataTypeIntBit"
	case BNNSDataTypeMiscellaneousBit:
		return "BNNSDataTypeMiscellaneousBit"
	case BNNSDataTypeUInt1:
		return "BNNSDataTypeUInt1"
	case BNNSDataTypeUInt16:
		return "BNNSDataTypeUInt16"
	case BNNSDataTypeUInt2:
		return "BNNSDataTypeUInt2"
	case BNNSDataTypeUInt3:
		return "BNNSDataTypeUInt3"
	case BNNSDataTypeUInt32:
		return "BNNSDataTypeUInt32"
	case BNNSDataTypeUInt4:
		return "BNNSDataTypeUInt4"
	case BNNSDataTypeUInt6:
		return "BNNSDataTypeUInt6"
	case BNNSDataTypeUInt64:
		return "BNNSDataTypeUInt64"
	case BNNSDataTypeUInt8:
		return "BNNSDataTypeUInt8"
	case BNNSDataTypeUIntBit:
		return "BNNSDataTypeUIntBit"
	default:
		return fmt.Sprintf("BNNSDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSDescriptorType
type BNNSDescriptorType uint32

const (
	// BNNSConstant: A constant that doesn’t have a gradient.
	BNNSConstant BNNSDescriptorType = 0
	// BNNSParameter: A parameter that’s trainable, such as weights or bias.
	BNNSParameter BNNSDescriptorType = 2
	// BNNSSample: A sample such as input or output.
	BNNSSample BNNSDescriptorType = 1
)

func (e BNNSDescriptorType) String() string {
	switch e {
	case BNNSConstant:
		return "BNNSConstant"
	case BNNSParameter:
		return "BNNSParameter"
	case BNNSSample:
		return "BNNSSample"
	default:
		return fmt.Sprintf("BNNSDescriptorType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSEmbeddingFlags
type BNNSEmbeddingFlags uint32

const (
	// BNNSEmbeddingFlagScaleGradientByFrequency: A flag that specifies that the operation scales calculated gradients based on the number of occurrence of the corresponding index in the input.
	BNNSEmbeddingFlagScaleGradientByFrequency BNNSEmbeddingFlags = 1
)

func (e BNNSEmbeddingFlags) String() string {
	switch e {
	case BNNSEmbeddingFlagScaleGradientByFrequency:
		return "BNNSEmbeddingFlagScaleGradientByFrequency"
	default:
		return fmt.Sprintf("BNNSEmbeddingFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterType
type BNNSFilterType uint32

const (
	// BNNSArithmetic: An arithmetic filter.
	BNNSArithmetic BNNSFilterType = 8
	// BNNSBatchNorm: A batch normalization filter.
	BNNSBatchNorm BNNSFilterType = 2
	// BNNSConvolution: A convolution filter.
	BNNSConvolution BNNSFilterType = 0
	// BNNSFullyConnected: A fully connected filter.
	BNNSFullyConnected BNNSFilterType = 1
	// BNNSGroupNorm: A group normalization filter.
	BNNSGroupNorm BNNSFilterType = 5
	// BNNSInstanceNorm: An instance normalization filter.
	BNNSInstanceNorm BNNSFilterType = 3
	// BNNSLayerNorm: A layer normalization filter.
	BNNSLayerNorm BNNSFilterType = 4
	// BNNSQuantization: A quantization filter.
	BNNSQuantization BNNSFilterType = 7
	// BNNSTransposedConvolution: A transposed convolution filter.
	BNNSTransposedConvolution BNNSFilterType = 6
)

func (e BNNSFilterType) String() string {
	switch e {
	case BNNSArithmetic:
		return "BNNSArithmetic"
	case BNNSBatchNorm:
		return "BNNSBatchNorm"
	case BNNSConvolution:
		return "BNNSConvolution"
	case BNNSFullyConnected:
		return "BNNSFullyConnected"
	case BNNSGroupNorm:
		return "BNNSGroupNorm"
	case BNNSInstanceNorm:
		return "BNNSInstanceNorm"
	case BNNSLayerNorm:
		return "BNNSLayerNorm"
	case BNNSQuantization:
		return "BNNSQuantization"
	case BNNSTransposedConvolution:
		return "BNNSTransposedConvolution"
	default:
		return fmt.Sprintf("BNNSFilterType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSFlags
type BNNSFlags uint32

const (
	// BNNSFlagsUseClientPtr: A flag that instructs the filter to use pointers to data you provide at creation time.
	BNNSFlagsUseClientPtr BNNSFlags = 0x1
)

func (e BNNSFlags) String() string {
	switch e {
	case BNNSFlagsUseClientPtr:
		return "BNNSFlagsUseClientPtr"
	default:
		return fmt.Sprintf("BNNSFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphArgumentIntent
type BNNSGraphArgumentIntent uint32

const (
	// BNNSGraphArgumentIntentIn: A constant that specifies the argument provides an input tensor.
	BNNSGraphArgumentIntentIn BNNSGraphArgumentIntent = 1
	// BNNSGraphArgumentIntentInOut: A constant that specifies the argument is an in-place input and output tensor.
	BNNSGraphArgumentIntentInOut BNNSGraphArgumentIntent = 3
	// BNNSGraphArgumentIntentOut: A constant that specifies the argument provides an output tensor.
	BNNSGraphArgumentIntentOut BNNSGraphArgumentIntent = 2
)

func (e BNNSGraphArgumentIntent) String() string {
	switch e {
	case BNNSGraphArgumentIntentIn:
		return "BNNSGraphArgumentIntentIn"
	case BNNSGraphArgumentIntentInOut:
		return "BNNSGraphArgumentIntentInOut"
	case BNNSGraphArgumentIntentOut:
		return "BNNSGraphArgumentIntentOut"
	default:
		return fmt.Sprintf("BNNSGraphArgumentIntent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphArgumentType
type BNNSGraphArgumentType uint32

const (
	// BNNSGraphArgumentTypePointer: A pointer to the raw data for the tensor.
	BNNSGraphArgumentTypePointer BNNSGraphArgumentType = 0
	// BNNSGraphArgumentTypeTensor: A tensor structure.
	BNNSGraphArgumentTypeTensor BNNSGraphArgumentType = 2
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

// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphMessageLevel
type BNNSGraphMessageLevel uint32

const (
	// BNNSGraphMessageLevelError: A constant that specifies error message types.
	BNNSGraphMessageLevelError BNNSGraphMessageLevel = 8
	// BNNSGraphMessageLevelInfo: A constant that specifies information message types.
	BNNSGraphMessageLevelInfo BNNSGraphMessageLevel = 1
	// BNNSGraphMessageLevelUnsupported: A constant that specifies unsupported function message types.
	BNNSGraphMessageLevelUnsupported BNNSGraphMessageLevel = 2
	// BNNSGraphMessageLevelWarning: A constant that specifies warning message types.
	BNNSGraphMessageLevelWarning BNNSGraphMessageLevel = 4
)

func (e BNNSGraphMessageLevel) String() string {
	switch e {
	case BNNSGraphMessageLevelError:
		return "BNNSGraphMessageLevelError"
	case BNNSGraphMessageLevelInfo:
		return "BNNSGraphMessageLevelInfo"
	case BNNSGraphMessageLevelUnsupported:
		return "BNNSGraphMessageLevelUnsupported"
	case BNNSGraphMessageLevelWarning:
		return "BNNSGraphMessageLevelWarning"
	default:
		return fmt.Sprintf("BNNSGraphMessageLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphOptimizationPreference
type BNNSGraphOptimizationPreference uint32

const (
	// BNNSGraphOptimizationPreferenceIRSize: A constant that specifies compilation optimization for smallest graph size on disk.
	BNNSGraphOptimizationPreferenceIRSize BNNSGraphOptimizationPreference = 1
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

// See: https://developer.apple.com/documentation/Accelerate/BNNSInterpolationMethod
type BNNSInterpolationMethod uint32

const (
	// BNNSInterpolationMethodLinear: Interpolation that is linear or bilinear depending on the number of resized dimensions.
	BNNSInterpolationMethodLinear BNNSInterpolationMethod = 1
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

// See: https://developer.apple.com/documentation/Accelerate/BNNSLayerFlags
type BNNSLayerFlags uint32

const (
	// BNNSLayerFlagsLSTMBidirectional: A flag that enables bidirectional long short-term memory (LSTM).
	BNNSLayerFlagsLSTMBidirectional BNNSLayerFlags = 0x1
	// BNNSLayerFlagsLSTMDefaultActivations: A flag that ignores the specified gate activations and instructs the operation to use default activations.
	BNNSLayerFlagsLSTMDefaultActivations BNNSLayerFlags = 0x2
)

func (e BNNSLayerFlags) String() string {
	switch e {
	case BNNSLayerFlagsLSTMBidirectional:
		return "BNNSLayerFlagsLSTMBidirectional"
	case BNNSLayerFlagsLSTMDefaultActivations:
		return "BNNSLayerFlagsLSTMDefaultActivations"
	default:
		return fmt.Sprintf("BNNSLayerFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSLinearSamplingMode
type BNNSLinearSamplingMode uint32

const (
	// BNNSLinearSamplingAlignCorners: The align corners sampling mode.
	BNNSLinearSamplingAlignCorners BNNSLinearSamplingMode = 1
	// BNNSLinearSamplingDefault: The default linear sampling mode.
	BNNSLinearSamplingDefault BNNSLinearSamplingMode = 0
	// BNNSLinearSamplingOffsetCorners: The offset corners sampling mode.
	BNNSLinearSamplingOffsetCorners BNNSLinearSamplingMode = 4
	// BNNSLinearSamplingStrictAlignCorners: The strict align corners sampling mode.
	BNNSLinearSamplingStrictAlignCorners BNNSLinearSamplingMode = 3
	// BNNSLinearSamplingUnalignCorners: The unalign corners sampling mode.
	BNNSLinearSamplingUnalignCorners BNNSLinearSamplingMode = 2
)

func (e BNNSLinearSamplingMode) String() string {
	switch e {
	case BNNSLinearSamplingAlignCorners:
		return "BNNSLinearSamplingAlignCorners"
	case BNNSLinearSamplingDefault:
		return "BNNSLinearSamplingDefault"
	case BNNSLinearSamplingOffsetCorners:
		return "BNNSLinearSamplingOffsetCorners"
	case BNNSLinearSamplingStrictAlignCorners:
		return "BNNSLinearSamplingStrictAlignCorners"
	case BNNSLinearSamplingUnalignCorners:
		return "BNNSLinearSamplingUnalignCorners"
	default:
		return fmt.Sprintf("BNNSLinearSamplingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSLossFunction
type BNNSLossFunction uint32

const (
	// BNNSLossFunctionCategoricalCrossEntropy: Performs categorical cross entropy computation between input prediction and labels.
	BNNSLossFunctionCategoricalCrossEntropy BNNSLossFunction = 10
	// BNNSLossFunctionCosineDistance: Performs cosine distance loss computation between input predictions and labels.
	BNNSLossFunctionCosineDistance BNNSLossFunction = 7
	// BNNSLossFunctionHinge: Performs Hinge loss computation between labels and unbounded zero-centered binary predictions.
	BNNSLossFunctionHinge BNNSLossFunction = 8
	// BNNSLossFunctionHuber: Huber loss computation between input logits and one-hot encoded labels.
	BNNSLossFunctionHuber BNNSLossFunction = 4
	// BNNSLossFunctionLog: Log loss computation between labels and predictions.
	BNNSLossFunctionLog BNNSLossFunction = 6
	// BNNSLossFunctionMeanAbsoluteError: Mean absolute error (MAE) computation between input prediction and labels.
	BNNSLossFunctionMeanAbsoluteError BNNSLossFunction = 9
	// BNNSLossFunctionMeanSquareError: Mean square error (MSE) computation between input logits and one-hot encoded labels.
	BNNSLossFunctionMeanSquareError BNNSLossFunction = 3
	// BNNSLossFunctionSigmoidCrossEntropy: Sigmoid activation on input logits, and independent computation of cross-entropy loss for each class.
	BNNSLossFunctionSigmoidCrossEntropy BNNSLossFunction = 2
	// BNNSLossFunctionSoftmaxCrossEntropy: Softmax activation on input logits, and computation of cross-entropy loss with one-hot encoded labels.
	BNNSLossFunctionSoftmaxCrossEntropy BNNSLossFunction = 1
	// BNNSLossFunctionYolo: You Only Look Once (YOLO) loss computation between prediction and ground truth labels.
	BNNSLossFunctionYolo BNNSLossFunction = 5
)

func (e BNNSLossFunction) String() string {
	switch e {
	case BNNSLossFunctionCategoricalCrossEntropy:
		return "BNNSLossFunctionCategoricalCrossEntropy"
	case BNNSLossFunctionCosineDistance:
		return "BNNSLossFunctionCosineDistance"
	case BNNSLossFunctionHinge:
		return "BNNSLossFunctionHinge"
	case BNNSLossFunctionHuber:
		return "BNNSLossFunctionHuber"
	case BNNSLossFunctionLog:
		return "BNNSLossFunctionLog"
	case BNNSLossFunctionMeanAbsoluteError:
		return "BNNSLossFunctionMeanAbsoluteError"
	case BNNSLossFunctionMeanSquareError:
		return "BNNSLossFunctionMeanSquareError"
	case BNNSLossFunctionSigmoidCrossEntropy:
		return "BNNSLossFunctionSigmoidCrossEntropy"
	case BNNSLossFunctionSoftmaxCrossEntropy:
		return "BNNSLossFunctionSoftmaxCrossEntropy"
	case BNNSLossFunctionYolo:
		return "BNNSLossFunctionYolo"
	default:
		return fmt.Sprintf("BNNSLossFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSLossReductionFunction
type BNNSLossReductionFunction uint32

const (
	// BNNSLossReductionMean: Sums the loss of all samples in the batch and divides by the number of samples.
	BNNSLossReductionMean BNNSLossReductionFunction = 3
	// BNNSLossReductionNonZeroWeightMean: Sums the loss of all samples in the batch and divides by the number of non-zero weights.
	BNNSLossReductionNonZeroWeightMean BNNSLossReductionFunction = 4
	// BNNSLossReductionNone: Returns the loss without any reduction.
	BNNSLossReductionNone BNNSLossReductionFunction = 0
	// BNNSLossReductionSum: Sums the loss of all samples in the batch.
	BNNSLossReductionSum BNNSLossReductionFunction = 1
	// BNNSLossReductionWeightedMean: Sums the loss of all samples in the batch and divides by the sum of all weights.
	BNNSLossReductionWeightedMean BNNSLossReductionFunction = 2
)

func (e BNNSLossReductionFunction) String() string {
	switch e {
	case BNNSLossReductionMean:
		return "BNNSLossReductionMean"
	case BNNSLossReductionNonZeroWeightMean:
		return "BNNSLossReductionNonZeroWeightMean"
	case BNNSLossReductionNone:
		return "BNNSLossReductionNone"
	case BNNSLossReductionSum:
		return "BNNSLossReductionSum"
	case BNNSLossReductionWeightedMean:
		return "BNNSLossReductionWeightedMean"
	default:
		return fmt.Sprintf("BNNSLossReductionFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayFlags
type BNNSNDArrayFlags uint32

const (
	// BNNSNDArrayFlagBackpropAccumulate: A flag that indicates backpropagation adds the value of the Jacobian to the elements of this n-dimensional array.
	BNNSNDArrayFlagBackpropAccumulate BNNSNDArrayFlags = 1
	// BNNSNDArrayFlagBackpropSet: A flag that indicates the elements of this n-dimensional array are overwritten by the Jacobian during backpropagation.
	BNNSNDArrayFlagBackpropSet BNNSNDArrayFlags = 0
)

func (e BNNSNDArrayFlags) String() string {
	switch e {
	case BNNSNDArrayFlagBackpropAccumulate:
		return "BNNSNDArrayFlagBackpropAccumulate"
	case BNNSNDArrayFlagBackpropSet:
		return "BNNSNDArrayFlagBackpropSet"
	default:
		return fmt.Sprintf("BNNSNDArrayFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSNormType
type BNNSNormType uint32

const (
	// BNNSL2Norm: A constant that represents the L2 norm.
	BNNSL2Norm BNNSNormType = 1
)

func (e BNNSNormType) String() string {
	switch e {
	case BNNSL2Norm:
		return "BNNSL2Norm"
	default:
		return fmt.Sprintf("BNNSNormType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerClippingFunction
type BNNSOptimizerClippingFunction uint32

const (
	// BNNSOptimizerClippingByGlobalNorm: A constant that specifes clipping to a maximum global Euclidean norm.
	BNNSOptimizerClippingByGlobalNorm BNNSOptimizerClippingFunction = 3
	// BNNSOptimizerClippingByNorm: A constant that specifes clipping to a maximum Euclidean norm.
	BNNSOptimizerClippingByNorm BNNSOptimizerClippingFunction = 2
	// BNNSOptimizerClippingByValue: A constant that specifes clipping to minimum and maximum values.
	BNNSOptimizerClippingByValue BNNSOptimizerClippingFunction = 1
	// BNNSOptimizerClippingNone: A constant that specifes no clipping.
	BNNSOptimizerClippingNone BNNSOptimizerClippingFunction = 0
)

func (e BNNSOptimizerClippingFunction) String() string {
	switch e {
	case BNNSOptimizerClippingByGlobalNorm:
		return "BNNSOptimizerClippingByGlobalNorm"
	case BNNSOptimizerClippingByNorm:
		return "BNNSOptimizerClippingByNorm"
	case BNNSOptimizerClippingByValue:
		return "BNNSOptimizerClippingByValue"
	case BNNSOptimizerClippingNone:
		return "BNNSOptimizerClippingNone"
	default:
		return fmt.Sprintf("BNNSOptimizerClippingFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerFunction
type BNNSOptimizerFunction uint32

const (
	// BNNSOptimizerFunctionAdam: An optimizer function that updates parameters according to the Adam algorithm.
	BNNSOptimizerFunctionAdam BNNSOptimizerFunction = 2
	// BNNSOptimizerFunctionAdamAMSGrad: An optimizer function that updates parameters according to the AMSGrad variant of the Adam algorithm.
	BNNSOptimizerFunctionAdamAMSGrad BNNSOptimizerFunction = 5
	// BNNSOptimizerFunctionAdamAMSGradWithClipping: An optimizer function that updates parameters according to the AMSGrad variant of the Adam algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamAMSGradWithClipping BNNSOptimizerFunction = 11
	// BNNSOptimizerFunctionAdamW: An optimizer function that updates parameters according to the AdamW algorithm.
	BNNSOptimizerFunctionAdamW BNNSOptimizerFunction = 4
	// BNNSOptimizerFunctionAdamWAMSGrad: An optimizer function that updates parameters according to the AMSGrad variant of the AdamW algorithm.
	BNNSOptimizerFunctionAdamWAMSGrad BNNSOptimizerFunction = 6
	// BNNSOptimizerFunctionAdamWAMSGradWithClipping: An optimizer function that updates parameters according to the AMSGrad variant of the AdamW algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWAMSGradWithClipping BNNSOptimizerFunction = 12
	// BNNSOptimizerFunctionAdamWWithClipping: An optimizer function that updates parameters according to the AdamW algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWWithClipping BNNSOptimizerFunction = 10
	// BNNSOptimizerFunctionAdamWithClipping: An optimizer function that updates parameters according to the Adam algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionAdamWithClipping BNNSOptimizerFunction = 8
	// BNNSOptimizerFunctionRMSProp: An optimizer function that updates parameters according to the root mean square propagation (RMSProp) algorithm.
	BNNSOptimizerFunctionRMSProp BNNSOptimizerFunction = 3
	// BNNSOptimizerFunctionRMSPropWithClipping: An optimizer function that updates parameters according to the root mean square propagation (RMSProp) algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionRMSPropWithClipping BNNSOptimizerFunction = 9
	// BNNSOptimizerFunctionSGDMomentum: An optimizer function that updates parameters according to the stochastic gradient descent (SGD) with momentum algorithm.
	BNNSOptimizerFunctionSGDMomentum BNNSOptimizerFunction = 1
	// BNNSOptimizerFunctionSGDMomentumWithClipping: An optimizer function that updates parameters according to the stochastic gradient descent (SGD) with momentum algorithm and optionally clips the gradient by value or by norm.
	BNNSOptimizerFunctionSGDMomentumWithClipping BNNSOptimizerFunction = 7
)

func (e BNNSOptimizerFunction) String() string {
	switch e {
	case BNNSOptimizerFunctionAdam:
		return "BNNSOptimizerFunctionAdam"
	case BNNSOptimizerFunctionAdamAMSGrad:
		return "BNNSOptimizerFunctionAdamAMSGrad"
	case BNNSOptimizerFunctionAdamAMSGradWithClipping:
		return "BNNSOptimizerFunctionAdamAMSGradWithClipping"
	case BNNSOptimizerFunctionAdamW:
		return "BNNSOptimizerFunctionAdamW"
	case BNNSOptimizerFunctionAdamWAMSGrad:
		return "BNNSOptimizerFunctionAdamWAMSGrad"
	case BNNSOptimizerFunctionAdamWAMSGradWithClipping:
		return "BNNSOptimizerFunctionAdamWAMSGradWithClipping"
	case BNNSOptimizerFunctionAdamWWithClipping:
		return "BNNSOptimizerFunctionAdamWWithClipping"
	case BNNSOptimizerFunctionAdamWithClipping:
		return "BNNSOptimizerFunctionAdamWithClipping"
	case BNNSOptimizerFunctionRMSProp:
		return "BNNSOptimizerFunctionRMSProp"
	case BNNSOptimizerFunctionRMSPropWithClipping:
		return "BNNSOptimizerFunctionRMSPropWithClipping"
	case BNNSOptimizerFunctionSGDMomentum:
		return "BNNSOptimizerFunctionSGDMomentum"
	case BNNSOptimizerFunctionSGDMomentumWithClipping:
		return "BNNSOptimizerFunctionSGDMomentumWithClipping"
	default:
		return fmt.Sprintf("BNNSOptimizerFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerRegularizationFunction
type BNNSOptimizerRegularizationFunction uint32

const (
	// BNNSOptimizerRegularizationL1: A regularization function that applies L1 regularization.
	BNNSOptimizerRegularizationL1 BNNSOptimizerRegularizationFunction = 1
	// BNNSOptimizerRegularizationL2: A regularization function that applies L2 regularization.
	BNNSOptimizerRegularizationL2 BNNSOptimizerRegularizationFunction = 2
	// BNNSOptimizerRegularizationNone: A regularization function that adoesn’t apply any regularization.
	BNNSOptimizerRegularizationNone BNNSOptimizerRegularizationFunction = 0
)

func (e BNNSOptimizerRegularizationFunction) String() string {
	switch e {
	case BNNSOptimizerRegularizationL1:
		return "BNNSOptimizerRegularizationL1"
	case BNNSOptimizerRegularizationL2:
		return "BNNSOptimizerRegularizationL2"
	case BNNSOptimizerRegularizationNone:
		return "BNNSOptimizerRegularizationNone"
	default:
		return fmt.Sprintf("BNNSOptimizerRegularizationFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerSGDMomentumVariant
type BNNSOptimizerSGDMomentumVariant uint32

const (
	// BNNSSGDMomentumVariant0: A constant that indicates SGD momentum variant 0.
	BNNSSGDMomentumVariant0 BNNSOptimizerSGDMomentumVariant = 0
	// BNNSSGDMomentumVariant1: A constant that indicates SGD momentum variant 1.
	BNNSSGDMomentumVariant1 BNNSOptimizerSGDMomentumVariant = 1
	// BNNSSGDMomentumVariant2: A constant that indicates SGD momentum variant 2.
	BNNSSGDMomentumVariant2 BNNSOptimizerSGDMomentumVariant = 2
)

func (e BNNSOptimizerSGDMomentumVariant) String() string {
	switch e {
	case BNNSSGDMomentumVariant0:
		return "BNNSSGDMomentumVariant0"
	case BNNSSGDMomentumVariant1:
		return "BNNSSGDMomentumVariant1"
	case BNNSSGDMomentumVariant2:
		return "BNNSSGDMomentumVariant2"
	default:
		return fmt.Sprintf("BNNSOptimizerSGDMomentumVariant(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSPaddingMode
type BNNSPaddingMode uint32

const (
	// BNNSPaddingModeConstant: A constant that indicates that a padding operation fills the padded area with a specified constant.
	BNNSPaddingModeConstant BNNSPaddingMode = 0
	// BNNSPaddingModeReflect: A constant that indicates that a padding operation fills the padded area to form an odd-symmetric pattern.
	BNNSPaddingModeReflect BNNSPaddingMode = 1
	// BNNSPaddingModeSymmetric: A constant that indicates that a padding operation fills the padded area to form an even-symmetric pattern.
	BNNSPaddingModeSymmetric BNNSPaddingMode = 2
)

func (e BNNSPaddingMode) String() string {
	switch e {
	case BNNSPaddingModeConstant:
		return "BNNSPaddingModeConstant"
	case BNNSPaddingModeReflect:
		return "BNNSPaddingModeReflect"
	case BNNSPaddingModeSymmetric:
		return "BNNSPaddingModeSymmetric"
	default:
		return fmt.Sprintf("BNNSPaddingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSPointerSpecifier
type BNNSPointerSpecifier uint32

const (
	// BNNSPointerSpecifierAlpha: A constant that specifies the alpha pointer.
	BNNSPointerSpecifierAlpha BNNSPointerSpecifier = 0
	// BNNSPointerSpecifierBeta: A constant that specifies the beta pointer.
	BNNSPointerSpecifierBeta BNNSPointerSpecifier = 1
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

// See: https://developer.apple.com/documentation/Accelerate/BNNSPoolingFunction
type BNNSPoolingFunction uint32

const (
	// BNNSPoolingFunctionAverageCountExcludePadding: A function for pooling that computes the average of each element in the pooling kernel, excluding zero-padding.
	BNNSPoolingFunctionAverageCountExcludePadding BNNSPoolingFunction = 2
	// BNNSPoolingFunctionAverageCountIncludePadding: A function for pooling that computes the average of each element in the pooling kernel, including zero-padding.
	BNNSPoolingFunctionAverageCountIncludePadding BNNSPoolingFunction = 1
	// BNNSPoolingFunctionL2Norm: A function for pooling that computes the square root of the sum of squares of each element in the pooling kernel.
	BNNSPoolingFunctionL2Norm BNNSPoolingFunction = 4
	// BNNSPoolingFunctionMax: A function for pooling that computes the maximum of each element in the pooling kernel.
	BNNSPoolingFunctionMax BNNSPoolingFunction = 0
	// BNNSPoolingFunctionUnMax: A function for pooling that’s the partial inverse of max pooling and sets all nonmaximal values to zero.
	BNNSPoolingFunctionUnMax BNNSPoolingFunction = 3
	// Deprecated.
	BNNSPoolingFunctionAverage BNNSPoolingFunction = 1
)

func (e BNNSPoolingFunction) String() string {
	switch e {
	case BNNSPoolingFunctionAverageCountExcludePadding:
		return "BNNSPoolingFunctionAverageCountExcludePadding"
	case BNNSPoolingFunctionAverageCountIncludePadding:
		return "BNNSPoolingFunctionAverageCountIncludePadding"
	case BNNSPoolingFunctionL2Norm:
		return "BNNSPoolingFunctionL2Norm"
	case BNNSPoolingFunctionMax:
		return "BNNSPoolingFunctionMax"
	case BNNSPoolingFunctionUnMax:
		return "BNNSPoolingFunctionUnMax"
	default:
		return fmt.Sprintf("BNNSPoolingFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSQuantizerFunction
type BNNSQuantizerFunction uint32

const (
	// BNNSQuantizerFunctionDequantize: A constant that specifes conversion to a higher precision.
	BNNSQuantizerFunctionDequantize BNNSQuantizerFunction = 1
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

// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomGeneratorMethod
type BNNSRandomGeneratorMethod uint32

const (
	// BNNSRandomGeneratorMethodAES_CTR: A constant that specifes an implementation that’s based on the Advanced Encryption Standard (AES) hash of a counter.
	BNNSRandomGeneratorMethodAES_CTR BNNSRandomGeneratorMethod = 0
)

func (e BNNSRandomGeneratorMethod) String() string {
	switch e {
	case BNNSRandomGeneratorMethodAES_CTR:
		return "BNNSRandomGeneratorMethodAES_CTR"
	default:
		return fmt.Sprintf("BNNSRandomGeneratorMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSReduceFunction
type BNNSReduceFunction uint32

const (
	// BNNSReduceFunctionAll: An alias of the logical AND reduction function.
	BNNSReduceFunctionAll BNNSReduceFunction = 11
	// BNNSReduceFunctionAny: An alias of the logical OR reduction function.
	BNNSReduceFunctionAny BNNSReduceFunction = 10
	// BNNSReduceFunctionArgMax: A reduction function that computes the index of the maximum value.
	BNNSReduceFunctionArgMax BNNSReduceFunction = 2
	// BNNSReduceFunctionArgMin: A reduction function that computes the index of the minimum value.
	BNNSReduceFunctionArgMin BNNSReduceFunction = 3
	// BNNSReduceFunctionL1Norm: A reduction function that computes the sum of the absolute value of each element.
	BNNSReduceFunctionL1Norm BNNSReduceFunction = 9
	// BNNSReduceFunctionL2Norm: A reduction function that computes the Euclidean norm.
	BNNSReduceFunctionL2Norm BNNSReduceFunction = 12
	BNNSReduceFunctionLogSum BNNSReduceFunction = 16
	// BNNSReduceFunctionLogSumExp: A reduction function that computes the logarithm of the sum of the exponentials of each element.
	BNNSReduceFunctionLogSumExp BNNSReduceFunction = 13
	// BNNSReduceFunctionLogicalAnd: A reduction function that reduces a tensor to true if all elements are true.
	BNNSReduceFunctionLogicalAnd BNNSReduceFunction = 11
	// BNNSReduceFunctionLogicalOr: A reduction function that reduces a tensor to true if any element is true.
	BNNSReduceFunctionLogicalOr BNNSReduceFunction = 10
	// BNNSReduceFunctionMax: A reduction function that computes the maximum value.
	BNNSReduceFunctionMax BNNSReduceFunction = 0
	// BNNSReduceFunctionMean: A reduction function that computes the mean value.
	BNNSReduceFunctionMean BNNSReduceFunction = 4
	// BNNSReduceFunctionMeanNonZero: A reduction function that computes the mean value of nonzero elements.
	BNNSReduceFunctionMeanNonZero BNNSReduceFunction = 5
	// BNNSReduceFunctionMin: A reduction function that computes the minimum value.
	BNNSReduceFunctionMin BNNSReduceFunction = 1
	// BNNSReduceFunctionNone: A reduction function that copies the input to the output.
	BNNSReduceFunctionNone BNNSReduceFunction = 15
	// BNNSReduceFunctionProduct: A reduction function that computes the product of all values.
	BNNSReduceFunctionProduct BNNSReduceFunction = 14
	// BNNSReduceFunctionSum: A reduction function that computes the sum of all values.
	BNNSReduceFunctionSum BNNSReduceFunction = 6
	// BNNSReduceFunctionSumLog: A reduction function that computes the sum of the natural logarithm of all values.
	BNNSReduceFunctionSumLog BNNSReduceFunction = 8
	// BNNSReduceFunctionSumSquare: A reduction function that computes the sum of the square of all values.
	BNNSReduceFunctionSumSquare BNNSReduceFunction = 7
)

func (e BNNSReduceFunction) String() string {
	switch e {
	case BNNSReduceFunctionAll:
		return "BNNSReduceFunctionAll"
	case BNNSReduceFunctionAny:
		return "BNNSReduceFunctionAny"
	case BNNSReduceFunctionArgMax:
		return "BNNSReduceFunctionArgMax"
	case BNNSReduceFunctionArgMin:
		return "BNNSReduceFunctionArgMin"
	case BNNSReduceFunctionL1Norm:
		return "BNNSReduceFunctionL1Norm"
	case BNNSReduceFunctionL2Norm:
		return "BNNSReduceFunctionL2Norm"
	case BNNSReduceFunctionLogSum:
		return "BNNSReduceFunctionLogSum"
	case BNNSReduceFunctionLogSumExp:
		return "BNNSReduceFunctionLogSumExp"
	case BNNSReduceFunctionMax:
		return "BNNSReduceFunctionMax"
	case BNNSReduceFunctionMean:
		return "BNNSReduceFunctionMean"
	case BNNSReduceFunctionMeanNonZero:
		return "BNNSReduceFunctionMeanNonZero"
	case BNNSReduceFunctionMin:
		return "BNNSReduceFunctionMin"
	case BNNSReduceFunctionNone:
		return "BNNSReduceFunctionNone"
	case BNNSReduceFunctionProduct:
		return "BNNSReduceFunctionProduct"
	case BNNSReduceFunctionSum:
		return "BNNSReduceFunctionSum"
	case BNNSReduceFunctionSumLog:
		return "BNNSReduceFunctionSumLog"
	case BNNSReduceFunctionSumSquare:
		return "BNNSReduceFunctionSumSquare"
	default:
		return fmt.Sprintf("BNNSReduceFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSRelationalOperator
type BNNSRelationalOperator uint32

const (
	// BNNSRelationalOperatorEqual: The operator that indicates the equal-to relationship.
	BNNSRelationalOperatorEqual BNNSRelationalOperator = 0
	// BNNSRelationalOperatorGreater: The operator that indicates the greater-than relationship.
	BNNSRelationalOperatorGreater BNNSRelationalOperator = 3
	// BNNSRelationalOperatorGreaterEqual: The operator that indicates the greater-than or equal-to relationship.
	BNNSRelationalOperatorGreaterEqual BNNSRelationalOperator = 4
	// BNNSRelationalOperatorLess: The operator that indicates the less-than relationship.
	BNNSRelationalOperatorLess BNNSRelationalOperator = 1
	// BNNSRelationalOperatorLessEqual: The operator that indicates the less-than or equal-to relationship.
	BNNSRelationalOperatorLessEqual BNNSRelationalOperator = 2
	// BNNSRelationalOperatorLogicalAND: The operator that indicates the logical AND relationship.
	BNNSRelationalOperatorLogicalAND BNNSRelationalOperator = 6
	// BNNSRelationalOperatorLogicalNAND: The operator that indicates the logical NAND relationship.
	BNNSRelationalOperatorLogicalNAND BNNSRelationalOperator = 9
	// BNNSRelationalOperatorLogicalNOR: The operator that indicates the logical NOR relationship.
	BNNSRelationalOperatorLogicalNOR BNNSRelationalOperator = 10
	// BNNSRelationalOperatorLogicalNOT: The operator that indicates the logical NOT relationship.
	BNNSRelationalOperatorLogicalNOT BNNSRelationalOperator = 8
	// BNNSRelationalOperatorLogicalOR: The operator that indicates the logical OR relationship.
	BNNSRelationalOperatorLogicalOR BNNSRelationalOperator = 7
	// BNNSRelationalOperatorLogicalXOR: The operator that indicates the logical XOR relationship.
	BNNSRelationalOperatorLogicalXOR BNNSRelationalOperator = 11
	// BNNSRelationalOperatorNotEqual: The operator that indicates the not-equal relationship.
	BNNSRelationalOperatorNotEqual BNNSRelationalOperator = 5
)

func (e BNNSRelationalOperator) String() string {
	switch e {
	case BNNSRelationalOperatorEqual:
		return "BNNSRelationalOperatorEqual"
	case BNNSRelationalOperatorGreater:
		return "BNNSRelationalOperatorGreater"
	case BNNSRelationalOperatorGreaterEqual:
		return "BNNSRelationalOperatorGreaterEqual"
	case BNNSRelationalOperatorLess:
		return "BNNSRelationalOperatorLess"
	case BNNSRelationalOperatorLessEqual:
		return "BNNSRelationalOperatorLessEqual"
	case BNNSRelationalOperatorLogicalAND:
		return "BNNSRelationalOperatorLogicalAND"
	case BNNSRelationalOperatorLogicalNAND:
		return "BNNSRelationalOperatorLogicalNAND"
	case BNNSRelationalOperatorLogicalNOR:
		return "BNNSRelationalOperatorLogicalNOR"
	case BNNSRelationalOperatorLogicalNOT:
		return "BNNSRelationalOperatorLogicalNOT"
	case BNNSRelationalOperatorLogicalOR:
		return "BNNSRelationalOperatorLogicalOR"
	case BNNSRelationalOperatorLogicalXOR:
		return "BNNSRelationalOperatorLogicalXOR"
	case BNNSRelationalOperatorNotEqual:
		return "BNNSRelationalOperatorNotEqual"
	default:
		return fmt.Sprintf("BNNSRelationalOperator(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSShuffleType
type BNNSShuffleType uint32

const (
	BNNSShuffleTypeDepthToSpaceNCHW BNNSShuffleType = 2
	// BNNSShuffleTypePixelShuffleNCHW: The pixel shuffle for the NCHW (batch, channels, height, width) format, equivalent to depth-to-space in Column Row Depth (CRD) mode.
	BNNSShuffleTypePixelShuffleNCHW BNNSShuffleType = 0
	// BNNSShuffleTypePixelUnshuffleNCHW: The pixel unshuffle for the NCHW (batch, channels, height, width) format, equivalent to space-to-depth in Column Row Depth (CRD) mode.
	BNNSShuffleTypePixelUnshuffleNCHW BNNSShuffleType = 1
	BNNSShuffleTypeSpaceToDepthNCHW   BNNSShuffleType = 3
)

func (e BNNSShuffleType) String() string {
	switch e {
	case BNNSShuffleTypeDepthToSpaceNCHW:
		return "BNNSShuffleTypeDepthToSpaceNCHW"
	case BNNSShuffleTypePixelShuffleNCHW:
		return "BNNSShuffleTypePixelShuffleNCHW"
	case BNNSShuffleTypePixelUnshuffleNCHW:
		return "BNNSShuffleTypePixelUnshuffleNCHW"
	case BNNSShuffleTypeSpaceToDepthNCHW:
		return "BNNSShuffleTypeSpaceToDepthNCHW"
	default:
		return fmt.Sprintf("BNNSShuffleType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/BNNSSparsityType
type BNNSSparsityType uint32

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

// See: https://developer.apple.com/documentation/Accelerate/BNNSTargetSystem
type BNNSTargetSystem uint32

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
	AtlasConj      CBLAS_TRANSPOSE = 114
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

type FftForward int32

const (
	// FFT_FORWARD: Forward FFT.
	FFT_FORWARD FftForward = 1
	// FFT_INVERSE: Inverse FFT.
	FFT_INVERSE FftForward = -1
)

func (e FftForward) String() string {
	switch e {
	case FFT_FORWARD:
		return "FFT_FORWARD"
	case FFT_INVERSE:
		return "FFT_INVERSE"
	default:
		return fmt.Sprintf("FftForward(%d)", e)
	}
}

type FftRadix2 uint32

const (
	FFT_RADIX2 FftRadix2 = 0
	FFT_RADIX3 FftRadix2 = 1
	FFT_RADIX5 FftRadix2 = 2
)

func (e FftRadix2) String() string {
	switch e {
	case FFT_RADIX2:
		return "FFT_RADIX2"
	case FFT_RADIX3:
		return "FFT_RADIX3"
	case FFT_RADIX5:
		return "FFT_RADIX5"
	default:
		return fmt.Sprintf("FftRadix2(%d)", e)
	}
}

type KFFT uint32

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

type KFFTDirection int32

const (
	KFFTDirection_Forward KFFTDirection = 1
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

type KRotate0DegreesClockwise uint32

const (
	// KRotate0DegreesClockwiseValue: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesClockwiseValue KRotate0DegreesClockwise = 0
	// KRotate0DegreesCounterClockwise: A constant that specifies rotation by 0° (that is, copy without rotating).
	KRotate0DegreesCounterClockwise KRotate0DegreesClockwise = 0
	// KRotate180DegreesClockwise: A constant that specifies rotation by 180° clockwise.
	KRotate180DegreesClockwise KRotate0DegreesClockwise = 2
	// KRotate180DegreesCounterClockwise: A constant that specifies rotation by 180° counterclockwise.
	KRotate180DegreesCounterClockwise KRotate0DegreesClockwise = 2
	// KRotate270DegreesClockwise: A constant that specifies rotation by 270° clockwise.
	KRotate270DegreesClockwise KRotate0DegreesClockwise = 1
	// KRotate270DegreesCounterClockwise: A constant that specifies rotation by 270° counterclockwise.
	KRotate270DegreesCounterClockwise KRotate0DegreesClockwise = 3
	// KRotate90DegreesClockwise: A constant that specifies rotation by 90° clockwise.
	KRotate90DegreesClockwise KRotate0DegreesClockwise = 3
	// KRotate90DegreesCounterClockwise: A constant that specifies rotation by 90° counterclockwise.
	KRotate90DegreesCounterClockwise KRotate0DegreesClockwise = 1
)

func (e KRotate0DegreesClockwise) String() string {
	switch e {
	case KRotate0DegreesClockwiseValue:
		return "KRotate0DegreesClockwiseValue"
	case KRotate180DegreesClockwise:
		return "KRotate180DegreesClockwise"
	case KRotate270DegreesClockwise:
		return "KRotate270DegreesClockwise"
	case KRotate270DegreesCounterClockwise:
		return "KRotate270DegreesCounterClockwise"
	default:
		return fmt.Sprintf("KRotate0DegreesClockwise(%d)", e)
	}
}

type KvImageBufferTypeCode uint32

const (
	// KvImageBufferTypeCode_Alpha: The buffer contains the alpha channel.
	KvImageBufferTypeCode_Alpha KvImageBufferTypeCode = 17
	// KvImageBufferTypeCode_CGFormat: The buffer contains data describable as a vImage Core Graphics image format as a single buffer.
	KvImageBufferTypeCode_CGFormat KvImageBufferTypeCode = 24
	// KvImageBufferTypeCode_CMYK_Black: If the image has a CMYK color model, the buffer contains the black channel.
	KvImageBufferTypeCode_CMYK_Black KvImageBufferTypeCode = 4
	// KvImageBufferTypeCode_CMYK_Cyan: If the image has a CMYK color model, the buffer contains the cyan channel.
	KvImageBufferTypeCode_CMYK_Cyan KvImageBufferTypeCode = 1
	// KvImageBufferTypeCode_CMYK_Magenta: If the image has a CMYK color model, the buffer contains the magenta channel.
	KvImageBufferTypeCode_CMYK_Magenta KvImageBufferTypeCode = 2
	// KvImageBufferTypeCode_CMYK_Yellow: If the image has a CMYK color model, the buffer contains the yellow channel.
	KvImageBufferTypeCode_CMYK_Yellow KvImageBufferTypeCode = 3
	// KvImageBufferTypeCode_CVPixelBuffer_YCbCr: The buffer contains luminance and both chroma channels interleaved according to the vImageConstCVImageFormat image type.
	KvImageBufferTypeCode_CVPixelBuffer_YCbCr KvImageBufferTypeCode = 19
	// KvImageBufferTypeCode_Cb: The buffer contains the blue chrominance channel.
	KvImageBufferTypeCode_Cb KvImageBufferTypeCode = 22
	// KvImageBufferTypeCode_Chroma: The buffer contains both chrominance channels, interleaved.
	KvImageBufferTypeCode_Chroma KvImageBufferTypeCode = 21
	// KvImageBufferTypeCode_Chunky: The buffer contains chunky data not describable as a vImage Core Graphics image format.
	KvImageBufferTypeCode_Chunky              KvImageBufferTypeCode = 25
	KvImageBufferTypeCode_ColorSpaceChannel1  KvImageBufferTypeCode = 1
	KvImageBufferTypeCode_ColorSpaceChannel10 KvImageBufferTypeCode = 10
	KvImageBufferTypeCode_ColorSpaceChannel11 KvImageBufferTypeCode = 11
	KvImageBufferTypeCode_ColorSpaceChannel12 KvImageBufferTypeCode = 12
	KvImageBufferTypeCode_ColorSpaceChannel13 KvImageBufferTypeCode = 13
	KvImageBufferTypeCode_ColorSpaceChannel14 KvImageBufferTypeCode = 14
	KvImageBufferTypeCode_ColorSpaceChannel15 KvImageBufferTypeCode = 15
	KvImageBufferTypeCode_ColorSpaceChannel16 KvImageBufferTypeCode = 16
	KvImageBufferTypeCode_ColorSpaceChannel2  KvImageBufferTypeCode = 2
	KvImageBufferTypeCode_ColorSpaceChannel3  KvImageBufferTypeCode = 3
	KvImageBufferTypeCode_ColorSpaceChannel4  KvImageBufferTypeCode = 4
	KvImageBufferTypeCode_ColorSpaceChannel5  KvImageBufferTypeCode = 5
	KvImageBufferTypeCode_ColorSpaceChannel6  KvImageBufferTypeCode = 6
	KvImageBufferTypeCode_ColorSpaceChannel7  KvImageBufferTypeCode = 7
	KvImageBufferTypeCode_ColorSpaceChannel8  KvImageBufferTypeCode = 8
	KvImageBufferTypeCode_ColorSpaceChannel9  KvImageBufferTypeCode = 9
	// KvImageBufferTypeCode_Cr: The buffer contains the red chrominance channel.
	KvImageBufferTypeCode_Cr KvImageBufferTypeCode = 23
	// KvImageBufferTypeCode_EndOfList: End of list marker.
	KvImageBufferTypeCode_EndOfList KvImageBufferTypeCode = 0
	// KvImageBufferTypeCode_Indexed: The buffer contains data in an indexed colorspace.
	KvImageBufferTypeCode_Indexed KvImageBufferTypeCode = 18
	// KvImageBufferTypeCode_LAB_A: If the image has a LAB color model, the buffer contains the a* channel.
	KvImageBufferTypeCode_LAB_A KvImageBufferTypeCode = 2
	// KvImageBufferTypeCode_LAB_B: If the image has a LAB color model, the buffer contains the b* channel.
	KvImageBufferTypeCode_LAB_B KvImageBufferTypeCode = 3
	// KvImageBufferTypeCode_LAB_L: If the image has a LAB color model, the buffer contains the L* channel.
	KvImageBufferTypeCode_LAB_L KvImageBufferTypeCode = 1
	// KvImageBufferTypeCode_Luminance: The buffer contains only luminance data.
	KvImageBufferTypeCode_Luminance KvImageBufferTypeCode = 20
	// KvImageBufferTypeCode_Monochrome: The buffer contains a single color channel.
	KvImageBufferTypeCode_Monochrome KvImageBufferTypeCode = 1
	// KvImageBufferTypeCode_RGB_Blue: If the image has a RGB color model, the buffer contains the blue channel.
	KvImageBufferTypeCode_RGB_Blue KvImageBufferTypeCode = 3
	// KvImageBufferTypeCode_RGB_Green: If the image has a RGB color model, the buffer contains the green channel.
	KvImageBufferTypeCode_RGB_Green KvImageBufferTypeCode = 2
	// KvImageBufferTypeCode_RGB_Red: If the image has a RGB color model, the buffer contains the red channel.
	KvImageBufferTypeCode_RGB_Red           KvImageBufferTypeCode = 1
	KvImageBufferTypeCode_UniqueFormatCount KvImageBufferTypeCode = 26
	// KvImageBufferTypeCode_XYZ_X: If the image has a XYZ color model, the buffer contains the X channel.
	KvImageBufferTypeCode_XYZ_X KvImageBufferTypeCode = 1
	// KvImageBufferTypeCode_XYZ_Y: If the image has a XYZ color model, the buffer contains the Y channel.
	KvImageBufferTypeCode_XYZ_Y KvImageBufferTypeCode = 2
	// KvImageBufferTypeCode_XYZ_Z: If the image has a XYZ color model, the buffer contains the Z channel.
	KvImageBufferTypeCode_XYZ_Z KvImageBufferTypeCode = 3
)

func (e KvImageBufferTypeCode) String() string {
	switch e {
	case KvImageBufferTypeCode_Alpha:
		return "KvImageBufferTypeCode_Alpha"
	case KvImageBufferTypeCode_CGFormat:
		return "KvImageBufferTypeCode_CGFormat"
	case KvImageBufferTypeCode_CMYK_Black:
		return "KvImageBufferTypeCode_CMYK_Black"
	case KvImageBufferTypeCode_CMYK_Cyan:
		return "KvImageBufferTypeCode_CMYK_Cyan"
	case KvImageBufferTypeCode_CMYK_Magenta:
		return "KvImageBufferTypeCode_CMYK_Magenta"
	case KvImageBufferTypeCode_CMYK_Yellow:
		return "KvImageBufferTypeCode_CMYK_Yellow"
	case KvImageBufferTypeCode_CVPixelBuffer_YCbCr:
		return "KvImageBufferTypeCode_CVPixelBuffer_YCbCr"
	case KvImageBufferTypeCode_Cb:
		return "KvImageBufferTypeCode_Cb"
	case KvImageBufferTypeCode_Chroma:
		return "KvImageBufferTypeCode_Chroma"
	case KvImageBufferTypeCode_Chunky:
		return "KvImageBufferTypeCode_Chunky"
	case KvImageBufferTypeCode_ColorSpaceChannel10:
		return "KvImageBufferTypeCode_ColorSpaceChannel10"
	case KvImageBufferTypeCode_ColorSpaceChannel11:
		return "KvImageBufferTypeCode_ColorSpaceChannel11"
	case KvImageBufferTypeCode_ColorSpaceChannel12:
		return "KvImageBufferTypeCode_ColorSpaceChannel12"
	case KvImageBufferTypeCode_ColorSpaceChannel13:
		return "KvImageBufferTypeCode_ColorSpaceChannel13"
	case KvImageBufferTypeCode_ColorSpaceChannel14:
		return "KvImageBufferTypeCode_ColorSpaceChannel14"
	case KvImageBufferTypeCode_ColorSpaceChannel15:
		return "KvImageBufferTypeCode_ColorSpaceChannel15"
	case KvImageBufferTypeCode_ColorSpaceChannel16:
		return "KvImageBufferTypeCode_ColorSpaceChannel16"
	case KvImageBufferTypeCode_ColorSpaceChannel5:
		return "KvImageBufferTypeCode_ColorSpaceChannel5"
	case KvImageBufferTypeCode_ColorSpaceChannel6:
		return "KvImageBufferTypeCode_ColorSpaceChannel6"
	case KvImageBufferTypeCode_ColorSpaceChannel7:
		return "KvImageBufferTypeCode_ColorSpaceChannel7"
	case KvImageBufferTypeCode_ColorSpaceChannel8:
		return "KvImageBufferTypeCode_ColorSpaceChannel8"
	case KvImageBufferTypeCode_ColorSpaceChannel9:
		return "KvImageBufferTypeCode_ColorSpaceChannel9"
	case KvImageBufferTypeCode_Cr:
		return "KvImageBufferTypeCode_Cr"
	case KvImageBufferTypeCode_EndOfList:
		return "KvImageBufferTypeCode_EndOfList"
	case KvImageBufferTypeCode_Indexed:
		return "KvImageBufferTypeCode_Indexed"
	case KvImageBufferTypeCode_Luminance:
		return "KvImageBufferTypeCode_Luminance"
	case KvImageBufferTypeCode_UniqueFormatCount:
		return "KvImageBufferTypeCode_UniqueFormatCount"
	default:
		return fmt.Sprintf("KvImageBufferTypeCode(%d)", e)
	}
}

type KvImageCVImageFormat int32

const (
	// KvImageCVImageFormat_AlphaIsOneHint: A hint that indicates the alpha channel is opaque.
	KvImageCVImageFormat_AlphaIsOneHint KvImageCVImageFormat = -21604
	// KvImageCVImageFormat_ChromaSiting: An error code that indicates the chroma siting information is absent.
	KvImageCVImageFormat_ChromaSiting KvImageCVImageFormat = -21601
	// KvImageCVImageFormat_ColorSpace: An error code that indicates the image’s color space is missing.
	KvImageCVImageFormat_ColorSpace KvImageCVImageFormat = -21602
	// KvImageCVImageFormat_ConversionMatrix: An error code that indicates the required conversion matrix is absent.
	KvImageCVImageFormat_ConversionMatrix KvImageCVImageFormat = -21600
	// KvImageCVImageFormat_NoError: An error code that indicates the conversion completed without error.
	KvImageCVImageFormat_NoError KvImageCVImageFormat = 0
	// KvImageCVImageFormat_VideoChannelDescription: An error code that indicates the range and clipping information is missing.
	KvImageCVImageFormat_VideoChannelDescription KvImageCVImageFormat = -21603
)

func (e KvImageCVImageFormat) String() string {
	switch e {
	case KvImageCVImageFormat_AlphaIsOneHint:
		return "KvImageCVImageFormat_AlphaIsOneHint"
	case KvImageCVImageFormat_ChromaSiting:
		return "KvImageCVImageFormat_ChromaSiting"
	case KvImageCVImageFormat_ColorSpace:
		return "KvImageCVImageFormat_ColorSpace"
	case KvImageCVImageFormat_ConversionMatrix:
		return "KvImageCVImageFormat_ConversionMatrix"
	case KvImageCVImageFormat_NoError:
		return "KvImageCVImageFormat_NoError"
	case KvImageCVImageFormat_VideoChannelDescription:
		return "KvImageCVImageFormat_VideoChannelDescription"
	default:
		return fmt.Sprintf("KvImageCVImageFormat(%d)", e)
	}
}

type KvImageConvert uint32

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
	KvImageConvert_OrderedNoiseShapeMask KvImageConvert = 0xf0000000
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
	case KvImageConvert_OrderedNoiseShapeMask:
		return "KvImageConvert_OrderedNoiseShapeMask"
	case KvImageConvert_OrderedUniformBlue:
		return "KvImageConvert_OrderedUniformBlue"
	default:
		return fmt.Sprintf("KvImageConvert(%d)", e)
	}
}

type KvImageGamma uint32

const (
	// KvImageGamma_11_over_5_half_precision: A half-precision calculation using a gamma value of 11/5 or 2.2.
	KvImageGamma_11_over_5_half_precision KvImageGamma = 5
	// KvImageGamma_11_over_9_half_precision: A half-precision calculation using a gamma value of 11/9 or (11/5)/(9/5).
	KvImageGamma_11_over_9_half_precision KvImageGamma = 8
	// KvImageGamma_5_over_11_half_precision: A half-precision calculation using a gamma value of 5/11 or 1/2.2.
	KvImageGamma_5_over_11_half_precision KvImageGamma = 4
	// KvImageGamma_5_over_9_half_precision: A half-precision calculation using a gamma value of 5/9 or 1/1.8.
	KvImageGamma_5_over_9_half_precision KvImageGamma = 2
	// KvImageGamma_9_over_11_half_precision: A half-precision calculation using a gamma value of 9/11 or (9/5)/(11/5).
	KvImageGamma_9_over_11_half_precision KvImageGamma = 9
	// KvImageGamma_9_over_5_half_precision: A half-precision calculation using a gamma value of 9/5 or 1.8.
	KvImageGamma_9_over_5_half_precision KvImageGamma = 3
	// KvImageGamma_BT709_forward_half_precision: The ITU-R BT.709 standard.
	KvImageGamma_BT709_forward_half_precision KvImageGamma = 10
	// KvImageGamma_BT709_reverse_half_precision: The ITU-R BT.709 standard reverse.
	KvImageGamma_BT709_reverse_half_precision KvImageGamma = 11
	// KvImageGamma_UseGammaValue: A user-defined gamma value with full-precision calculation.
	KvImageGamma_UseGammaValue KvImageGamma = 0
	// KvImageGamma_UseGammaValue_half_precision: A user-defined gamma value with half-precision calculation.
	KvImageGamma_UseGammaValue_half_precision KvImageGamma = 1
	// KvImageGamma_sRGB_forward_half_precision: A half-precision calculation using the sRGB standard gamma value of 2.2.
	KvImageGamma_sRGB_forward_half_precision KvImageGamma = 6
	// KvImageGamma_sRGB_reverse_half_precision: A half-precision calculation using the sRGB standard gamma value of 1/2.2.
	KvImageGamma_sRGB_reverse_half_precision KvImageGamma = 7
)

func (e KvImageGamma) String() string {
	switch e {
	case KvImageGamma_11_over_5_half_precision:
		return "KvImageGamma_11_over_5_half_precision"
	case KvImageGamma_11_over_9_half_precision:
		return "KvImageGamma_11_over_9_half_precision"
	case KvImageGamma_5_over_11_half_precision:
		return "KvImageGamma_5_over_11_half_precision"
	case KvImageGamma_5_over_9_half_precision:
		return "KvImageGamma_5_over_9_half_precision"
	case KvImageGamma_9_over_11_half_precision:
		return "KvImageGamma_9_over_11_half_precision"
	case KvImageGamma_9_over_5_half_precision:
		return "KvImageGamma_9_over_5_half_precision"
	case KvImageGamma_BT709_forward_half_precision:
		return "KvImageGamma_BT709_forward_half_precision"
	case KvImageGamma_BT709_reverse_half_precision:
		return "KvImageGamma_BT709_reverse_half_precision"
	case KvImageGamma_UseGammaValue:
		return "KvImageGamma_UseGammaValue"
	case KvImageGamma_UseGammaValue_half_precision:
		return "KvImageGamma_UseGammaValue_half_precision"
	case KvImageGamma_sRGB_forward_half_precision:
		return "KvImageGamma_sRGB_forward_half_precision"
	case KvImageGamma_sRGB_reverse_half_precision:
		return "KvImageGamma_sRGB_reverse_half_precision"
	default:
		return fmt.Sprintf("KvImageGamma(%d)", e)
	}
}

type KvImageInterpolation uint32

const (
	// KvImageInterpolationLinear: Linear interpoation
	KvImageInterpolationLinear KvImageInterpolation = 1
	// KvImageInterpolationNearest: Nearest neigborhood
	KvImageInterpolationNearest KvImageInterpolation = 0
)

func (e KvImageInterpolation) String() string {
	switch e {
	case KvImageInterpolationLinear:
		return "KvImageInterpolationLinear"
	case KvImageInterpolationNearest:
		return "KvImageInterpolationNearest"
	default:
		return fmt.Sprintf("KvImageInterpolation(%d)", e)
	}
}

type KvImageMatrixType uint32

const (
	KvImageMatrixType_ARGBToYpCbCrMatrix KvImageMatrixType = 1
	KvImageMatrixType_None               KvImageMatrixType = 0
)

func (e KvImageMatrixType) String() string {
	switch e {
	case KvImageMatrixType_ARGBToYpCbCrMatrix:
		return "KvImageMatrixType_ARGBToYpCbCrMatrix"
	case KvImageMatrixType_None:
		return "KvImageMatrixType_None"
	default:
		return fmt.Sprintf("KvImageMatrixType(%d)", e)
	}
}

type KvImageNoError int32

const (
	// KvImageBufferSizeMismatch: The function requires the source and destination buffers to have the same height and the same width, but they do not.
	KvImageBufferSizeMismatch KvImageNoError = -21774
	KvImageColorSyncIsAbsent  KvImageNoError = -21779
	KvImageCoreVideoIsAbsent  KvImageNoError = -21784
	// KvImageInternalError: A serious error occured inside vImage, which prevented vImage from continuing.
	KvImageInternalError        KvImageNoError = -21776
	KvImageInvalidCVImageFormat KvImageNoError = -21782
	// KvImageInvalidEdgeStyle: The edge style specified is invalid.
	KvImageInvalidEdgeStyle   KvImageNoError = -21768
	KvImageInvalidImageFormat KvImageNoError = -21778
	KvImageInvalidImageObject KvImageNoError = -21781
	// KvImageInvalidKernelSize: Either the kernel height, the kernel width, or both, are even.
	KvImageInvalidKernelSize KvImageNoError = -21767
	// KvImageInvalidOffset_X: The `srcOffsetToROI_X` parameter that specifies the left edge of the region of interest is greater than the width of the source image.
	KvImageInvalidOffset_X KvImageNoError = -21769
	// KvImageInvalidOffset_Y: The `srcOffsetToROI_Y` parameter that specifies the top edge of the region of interest is greater than the height of the source image.
	KvImageInvalidOffset_Y KvImageNoError = -21770
	// KvImageInvalidParameter: Invalid parameter.
	KvImageInvalidParameter KvImageNoError = -21773
	KvImageInvalidRowBytes  KvImageNoError = -21777
	// KvImageMemoryAllocationError: An attempt to allocate memory failed.
	KvImageMemoryAllocationError KvImageNoError = -21771
	// KvImageNoErrorValue: The vImage function completed without error.
	KvImageNoErrorValue KvImageNoError = 0
	// KvImageNullPointerArgument: A pointer parameter is [NULL] and it must not be.
	KvImageNullPointerArgument         KvImageNoError = -21772
	KvImageOutOfPlaceOperationRequired KvImageNoError = -21780
	// KvImageRoiLargerThanInputBuffer: The region of interest, as specified by the `srcOffsetToROI_X` and `srcOffsetToROI_Y` parameters and the height and width of the destination buffer, extends beyond the bottom edge or right edge of the source buffer.
	KvImageRoiLargerThanInputBuffer KvImageNoError = -21766
	// KvImageUnknownFlagsBit: The flag is not recognized.
	KvImageUnknownFlagsBit KvImageNoError = -21775
	// KvImageUnsupportedConversion: Some lower level conversion APIs only support conversion among a sparse matrix of image formats.
	KvImageUnsupportedConversion KvImageNoError = -21783
)

func (e KvImageNoError) String() string {
	switch e {
	case KvImageBufferSizeMismatch:
		return "KvImageBufferSizeMismatch"
	case KvImageColorSyncIsAbsent:
		return "KvImageColorSyncIsAbsent"
	case KvImageCoreVideoIsAbsent:
		return "KvImageCoreVideoIsAbsent"
	case KvImageInternalError:
		return "KvImageInternalError"
	case KvImageInvalidCVImageFormat:
		return "KvImageInvalidCVImageFormat"
	case KvImageInvalidEdgeStyle:
		return "KvImageInvalidEdgeStyle"
	case KvImageInvalidImageFormat:
		return "KvImageInvalidImageFormat"
	case KvImageInvalidImageObject:
		return "KvImageInvalidImageObject"
	case KvImageInvalidKernelSize:
		return "KvImageInvalidKernelSize"
	case KvImageInvalidOffset_X:
		return "KvImageInvalidOffset_X"
	case KvImageInvalidOffset_Y:
		return "KvImageInvalidOffset_Y"
	case KvImageInvalidParameter:
		return "KvImageInvalidParameter"
	case KvImageInvalidRowBytes:
		return "KvImageInvalidRowBytes"
	case KvImageMemoryAllocationError:
		return "KvImageMemoryAllocationError"
	case KvImageNoErrorValue:
		return "KvImageNoErrorValue"
	case KvImageNullPointerArgument:
		return "KvImageNullPointerArgument"
	case KvImageOutOfPlaceOperationRequired:
		return "KvImageOutOfPlaceOperationRequired"
	case KvImageRoiLargerThanInputBuffer:
		return "KvImageRoiLargerThanInputBuffer"
	case KvImageUnknownFlagsBit:
		return "KvImageUnknownFlagsBit"
	case KvImageUnsupportedConversion:
		return "KvImageUnsupportedConversion"
	default:
		return fmt.Sprintf("KvImageNoError(%d)", e)
	}
}

type KvImageNoFlags uint32

const (
	// KvImageBackgroundColorFill: A flag that uses the background color for missing pixels.
	KvImageBackgroundColorFill KvImageNoFlags = 4
	// KvImageCopyInPlace: A flag that copies the value of the edge pixel in the source to the destination.
	KvImageCopyInPlace KvImageNoFlags = 2
	// KvImageDoNotClamp: A flag that disables clamping in some conversions to floating-point formats.
	KvImageDoNotClamp KvImageNoFlags = 2048
	// KvImageDoNotTile: A flag that disables vImage internal tiling routines.
	KvImageDoNotTile KvImageNoFlags = 16
	// KvImageEdgeExtend: A flag that extends the edges of the image infinitely.
	KvImageEdgeExtend KvImageNoFlags = 8
	// KvImageGetTempBufferSize: A flag that returns the minimum temporary buffer size for the operation, given the parameters provided.
	KvImageGetTempBufferSize KvImageNoFlags = 128
	// KvImageHDRContent: A flag that uses HDR-aware methods.
	KvImageHDRContent KvImageNoFlags = 1024
	// KvImageHighQualityResampling: A flag that uses a higher-quality, slower resampling filter for geometry operations.
	KvImageHighQualityResampling KvImageNoFlags = 32
	// KvImageLeaveAlphaUnchanged: A flag that restricts the operation to red, green, and blue channels only.
	KvImageLeaveAlphaUnchanged KvImageNoFlags = 1
	// KvImageNoAllocate: A flag that prevents vImage from allocating additional storage.
	KvImageNoAllocate KvImageNoFlags = 512
	// KvImageNoFlagsValue: A flag that sets the behavior to the default.
	KvImageNoFlagsValue KvImageNoFlags = 0
	// KvImagePrintDiagnosticsToConsole: A flag that prints a debug message if the operation fails.
	KvImagePrintDiagnosticsToConsole KvImageNoFlags = 256
	// KvImageTruncateKernel: A flag that uses only the part of the kernel that overlaps the image.
	KvImageTruncateKernel KvImageNoFlags = 64
	// KvImageUseFP16Accumulator: A flag that specifies vImage uses faster but lower-precision internal arithmetic for floating-point 16-bit operations.
	KvImageUseFP16Accumulator KvImageNoFlags = 4096
)

func (e KvImageNoFlags) String() string {
	switch e {
	case KvImageBackgroundColorFill:
		return "KvImageBackgroundColorFill"
	case KvImageCopyInPlace:
		return "KvImageCopyInPlace"
	case KvImageDoNotClamp:
		return "KvImageDoNotClamp"
	case KvImageDoNotTile:
		return "KvImageDoNotTile"
	case KvImageEdgeExtend:
		return "KvImageEdgeExtend"
	case KvImageGetTempBufferSize:
		return "KvImageGetTempBufferSize"
	case KvImageHDRContent:
		return "KvImageHDRContent"
	case KvImageHighQualityResampling:
		return "KvImageHighQualityResampling"
	case KvImageLeaveAlphaUnchanged:
		return "KvImageLeaveAlphaUnchanged"
	case KvImageNoAllocate:
		return "KvImageNoAllocate"
	case KvImageNoFlagsValue:
		return "KvImageNoFlagsValue"
	case KvImagePrintDiagnosticsToConsole:
		return "KvImagePrintDiagnosticsToConsole"
	case KvImageTruncateKernel:
		return "KvImageTruncateKernel"
	case KvImageUseFP16Accumulator:
		return "KvImageUseFP16Accumulator"
	default:
		return fmt.Sprintf("KvImageNoFlags(%d)", e)
	}
}

type KvimagePNGFilterValue uint32

const (
	// KvImage_PNG_FILTER_VALUE_AVG: A filter that predicts a pixel value from the average of the pixels to the left and above the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_AVG KvimagePNGFilterValue = 3
	// KvImage_PNG_FILTER_VALUE_NONE: No filtering.
	KvImage_PNG_FILTER_VALUE_NONE KvimagePNGFilterValue = 0
	// KvImage_PNG_FILTER_VALUE_PAETH: A filter that predicts a pixel value by applying a linear function to the pixels located to the left, above, and to the upper-left of the predicted pixel location.
	KvImage_PNG_FILTER_VALUE_PAETH KvimagePNGFilterValue = 4
	// KvImage_PNG_FILTER_VALUE_SUB: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located to the left.
	KvImage_PNG_FILTER_VALUE_SUB KvimagePNGFilterValue = 1
	// KvImage_PNG_FILTER_VALUE_UP: A filter that computes the difference between each byte of a pixel and the value of the corresponding byte of the pixel located above.
	KvImage_PNG_FILTER_VALUE_UP KvimagePNGFilterValue = 2
)

func (e KvimagePNGFilterValue) String() string {
	switch e {
	case KvImage_PNG_FILTER_VALUE_AVG:
		return "KvImage_PNG_FILTER_VALUE_AVG"
	case KvImage_PNG_FILTER_VALUE_NONE:
		return "KvImage_PNG_FILTER_VALUE_NONE"
	case KvImage_PNG_FILTER_VALUE_PAETH:
		return "KvImage_PNG_FILTER_VALUE_PAETH"
	case KvImage_PNG_FILTER_VALUE_SUB:
		return "KvImage_PNG_FILTER_VALUE_SUB"
	case KvImage_PNG_FILTER_VALUE_UP:
		return "KvImage_PNG_FILTER_VALUE_UP"
	default:
		return fmt.Sprintf("KvimagePNGFilterValue(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseControl_t
type SparseControl_t uint32

const (
	// SparseDefaultControl: A flag that indicates default values.
	SparseDefaultControl SparseControl_t = 0
)

func (e SparseControl_t) String() string {
	switch e {
	case SparseDefaultControl:
		return "SparseDefaultControl"
	default:
		return fmt.Sprintf("SparseControl_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseFactorization_t
type SparseFactorization_t uint8

const (
	// SparseFactorizationCholesky: A constant that represents Cholesky (LLᵀ) factorization.
	SparseFactorizationCholesky SparseFactorization_t = 0
	// SparseFactorizationCholeskyAtA: A constant that represents QR factorization without storing Q.
	SparseFactorizationCholeskyAtA SparseFactorization_t = 41
	// SparseFactorizationLDLT: A constant that represents the default LDLᵀ factorization.
	SparseFactorizationLDLT SparseFactorization_t = 1
	// SparseFactorizationLDLTSBK: A constant that represents LDLᵀ factorization with Supernode-Bunch-Kaufman and static pivoting.
	SparseFactorizationLDLTSBK SparseFactorization_t = 3
	// SparseFactorizationLDLTTPP: A constant that represents LDLᵀ factorization with full-threshold partial pivoting.
	SparseFactorizationLDLTTPP SparseFactorization_t = 4
	// SparseFactorizationLDLTUnpivoted: A constant that represents Cholesky-like LDLᵀ factorization with only one-by-one pivots and no pivoting.
	SparseFactorizationLDLTUnpivoted SparseFactorization_t = 2
	// SparseFactorizationLU: Default LU factorization, currently LU with TPP.
	SparseFactorizationLU SparseFactorization_t = 80
	// SparseFactorizationLUSPP: LU factorization with partial pivoting restricted to within supernodes only.
	SparseFactorizationLUSPP SparseFactorization_t = 82
	// SparseFactorizationLUTPP: LU factorization with threshold partial pivoting.
	SparseFactorizationLUTPP SparseFactorization_t = 83
	// SparseFactorizationLUUnpivoted: LU factorization with no numerical pivoting.
	SparseFactorizationLUUnpivoted SparseFactorization_t = 81
	// SparseFactorizationQR: A constant that represents QR factorization.
	SparseFactorizationQR SparseFactorization_t = 40
)

func (e SparseFactorization_t) String() string {
	switch e {
	case SparseFactorizationCholesky:
		return "SparseFactorizationCholesky"
	case SparseFactorizationCholeskyAtA:
		return "SparseFactorizationCholeskyAtA"
	case SparseFactorizationLDLT:
		return "SparseFactorizationLDLT"
	case SparseFactorizationLDLTSBK:
		return "SparseFactorizationLDLTSBK"
	case SparseFactorizationLDLTTPP:
		return "SparseFactorizationLDLTTPP"
	case SparseFactorizationLDLTUnpivoted:
		return "SparseFactorizationLDLTUnpivoted"
	case SparseFactorizationLU:
		return "SparseFactorizationLU"
	case SparseFactorizationLUSPP:
		return "SparseFactorizationLUSPP"
	case SparseFactorizationLUTPP:
		return "SparseFactorizationLUTPP"
	case SparseFactorizationLUUnpivoted:
		return "SparseFactorizationLUUnpivoted"
	case SparseFactorizationQR:
		return "SparseFactorizationQR"
	default:
		return fmt.Sprintf("SparseFactorization_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseGMRESVariant_t
type SparseGMRESVariant_t uint8

const (
	// SparseVariantDQGMRES: A constant that specifies the DQGMRES variant.
	SparseVariantDQGMRES SparseGMRESVariant_t = 0
	// SparseVariantFGMRES: A constant that specifies the flexible GMRES variant.
	SparseVariantFGMRES SparseGMRESVariant_t = 2
	// SparseVariantGMRES: A constant that specifies the standard restarted GMRES variant.
	SparseVariantGMRES SparseGMRESVariant_t = 1
)

func (e SparseGMRESVariant_t) String() string {
	switch e {
	case SparseVariantDQGMRES:
		return "SparseVariantDQGMRES"
	case SparseVariantFGMRES:
		return "SparseVariantFGMRES"
	case SparseVariantGMRES:
		return "SparseVariantGMRES"
	default:
		return fmt.Sprintf("SparseGMRESVariant_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseIterativeStatus_t
type SparseIterativeStatus_t int32

const (
	// SparseIterativeConverged: A status that indicates the convergence of all solutions.
	SparseIterativeConverged SparseIterativeStatus_t = 0
	// SparseIterativeIllConditioned: A status that indicates the operation determines the problem is sufficiently ill-conditioned that convergence is unlikely.
	SparseIterativeIllConditioned SparseIterativeStatus_t = -2
	// SparseIterativeInternalError: A status that indicates an internal failure.
	SparseIterativeInternalError SparseIterativeStatus_t = -99
	// SparseIterativeMaxIterations: A status that indicates a failure to converge one or more solutions in the maximum number of iterations.
	SparseIterativeMaxIterations SparseIterativeStatus_t = 1
	// SparseIterativeParameterError: A status that indicates an error with one or more parameters.
	SparseIterativeParameterError SparseIterativeStatus_t = -1
)

func (e SparseIterativeStatus_t) String() string {
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
		return fmt.Sprintf("SparseIterativeStatus_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseKind_t
type SparseKind_t uint32

const (
	// SparseHermitian: A flag to describe the type of matrix represented.
	SparseHermitian SparseKind_t = 7
	// SparseOrdinary: An unsymmetric sparse matrix without special structure.
	SparseOrdinary SparseKind_t = 0
	// SparseSymmetric: A symmetric sparse matrix.
	SparseSymmetric SparseKind_t = 3
	// SparseTriangular: A triangular sparse matrix with a nonunit diagonal.
	SparseTriangular SparseKind_t = 1
	// SparseUnitTriangular: A triangular sparse matrix with a unit diagonal.
	SparseUnitTriangular SparseKind_t = 2
)

func (e SparseKind_t) String() string {
	switch e {
	case SparseHermitian:
		return "SparseHermitian"
	case SparseOrdinary:
		return "SparseOrdinary"
	case SparseSymmetric:
		return "SparseSymmetric"
	case SparseTriangular:
		return "SparseTriangular"
	case SparseUnitTriangular:
		return "SparseUnitTriangular"
	default:
		return fmt.Sprintf("SparseKind_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseLSMRConvergenceTest_t
type SparseLSMRConvergenceTest_t int32

const (
	// SparseLSMRCTDefault: The default convergence test.
	SparseLSMRCTDefault SparseLSMRConvergenceTest_t = 0
	// SparseLSMRCTFongSaunders: Fong and Saunder’s original convergence test.
	SparseLSMRCTFongSaunders SparseLSMRConvergenceTest_t = 1
)

func (e SparseLSMRConvergenceTest_t) String() string {
	switch e {
	case SparseLSMRCTDefault:
		return "SparseLSMRCTDefault"
	case SparseLSMRCTFongSaunders:
		return "SparseLSMRCTFongSaunders"
	default:
		return fmt.Sprintf("SparseLSMRConvergenceTest_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseOrder_t
type SparseOrder_t uint8

const (
	// SparseOrderAMD: Approximate minimum degree (AMD) ordering.
	SparseOrderAMD SparseOrder_t = 2
	// SparseOrderCOLAMD: The column AMD ordering for AᵀA.
	SparseOrderCOLAMD SparseOrder_t = 4
	// SparseOrderDefault: The default ordering.
	SparseOrderDefault SparseOrder_t = 0
	// SparseOrderMTMetis: Specifies type of fill-reducing ordering.
	SparseOrderMTMetis SparseOrder_t = 5
	// SparseOrderMetis: METIS nested dissection ordering.
	SparseOrderMetis SparseOrder_t = 3
	// SparseOrderUser: The user-supplied ordering, or identity if the order parameter is null.
	SparseOrderUser SparseOrder_t = 1
)

func (e SparseOrder_t) String() string {
	switch e {
	case SparseOrderAMD:
		return "SparseOrderAMD"
	case SparseOrderCOLAMD:
		return "SparseOrderCOLAMD"
	case SparseOrderDefault:
		return "SparseOrderDefault"
	case SparseOrderMTMetis:
		return "SparseOrderMTMetis"
	case SparseOrderMetis:
		return "SparseOrderMetis"
	case SparseOrderUser:
		return "SparseOrderUser"
	default:
		return fmt.Sprintf("SparseOrder_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparsePreconditioner_t
type SparsePreconditioner_t int32

const (
	// SparsePreconditionerDiagScaling: A diagonal scaling preconditioner.
	SparsePreconditionerDiagScaling SparsePreconditioner_t = 3
	// SparsePreconditionerDiagonal: A Jacobi preconditioner.
	SparsePreconditionerDiagonal SparsePreconditioner_t = 2
	// SparsePreconditionerNone: No preconditioner.
	SparsePreconditionerNone SparsePreconditioner_t = 0
	// SparsePreconditionerUser: A user-provided preconditioner.
	SparsePreconditionerUser SparsePreconditioner_t = 1
)

func (e SparsePreconditioner_t) String() string {
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
		return fmt.Sprintf("SparsePreconditioner_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseScaling_t
type SparseScaling_t uint8

const (
	// SparseScalingDefault: Default scaling.
	SparseScalingDefault SparseScaling_t = 0
	// SparseScalingEquilibriationInf: The norm equilibration scaling using infinity norm.
	SparseScalingEquilibriationInf SparseScaling_t = 2
	// SparseScalingHungarianScalingAndOrdering: Scaling and ordering using the Hungarian algorithm.
	SparseScalingHungarianScalingAndOrdering SparseScaling_t = 4
	// SparseScalingHungarianScalingOnly: Scaling using the Hungarian algorithm.
	SparseScalingHungarianScalingOnly SparseScaling_t = 3
	// SparseScalingUser: User scaling.
	SparseScalingUser SparseScaling_t = 1
)

func (e SparseScaling_t) String() string {
	switch e {
	case SparseScalingDefault:
		return "SparseScalingDefault"
	case SparseScalingEquilibriationInf:
		return "SparseScalingEquilibriationInf"
	case SparseScalingHungarianScalingAndOrdering:
		return "SparseScalingHungarianScalingAndOrdering"
	case SparseScalingHungarianScalingOnly:
		return "SparseScalingHungarianScalingOnly"
	case SparseScalingUser:
		return "SparseScalingUser"
	default:
		return fmt.Sprintf("SparseScaling_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseStatus_t
type SparseStatus_t int32

const (
	// SparseFactorizationFailed: The factorization failed due to a numerical issue.
	SparseFactorizationFailed SparseStatus_t = -1
	// SparseInternalError: The factorization encountered an internal error, such as failing to allocate memory.
	SparseInternalError SparseStatus_t = -3
	// SparseMatrixIsSingular: The factorization aborted because the matrix is singular.
	SparseMatrixIsSingular SparseStatus_t = -2
	// SparseParameterError: An error in a user-supplied parameter.
	SparseParameterError SparseStatus_t = -4
	// SparseStatusOK: The factorization was successful.
	SparseStatusOK SparseStatus_t = 0
	// SparseStatusReleased: The system freed the factorization object.
	SparseStatusReleased SparseStatus_t = -2147483647
)

func (e SparseStatus_t) String() string {
	switch e {
	case SparseFactorizationFailed:
		return "SparseFactorizationFailed"
	case SparseInternalError:
		return "SparseInternalError"
	case SparseMatrixIsSingular:
		return "SparseMatrixIsSingular"
	case SparseParameterError:
		return "SparseParameterError"
	case SparseStatusOK:
		return "SparseStatusOK"
	case SparseStatusReleased:
		return "SparseStatusReleased"
	default:
		return fmt.Sprintf("SparseStatus_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseSubfactor_t
type SparseSubfactor_t uint8

const (
	// SparseSubfactorD: A D factor subfactor that’s valid for LDLᵀ` `only.
	SparseSubfactorD SparseSubfactor_t = 4
	// SparseSubfactorInvalid: An invalid subfactor that indicates the requested type is incompatible with the supplied factorization or the system has destroyed it.
	SparseSubfactorInvalid SparseSubfactor_t = 0
	// SparseSubfactorL: An L factor subfactor that’s valid for Cholesky and LDLᵀ only.
	SparseSubfactorL SparseSubfactor_t = 3
	// SparseSubfactorP: A permutation subfactor that’s valid for all factorization types.
	SparseSubfactorP SparseSubfactor_t = 1
	// SparseSubfactorPLPS: A half-solve subfactor that’s valid for Cholesky and LDLᵀ only.
	SparseSubfactorPLPS SparseSubfactor_t = 5
	// SparseSubfactorQ: A Q factor subfactor that’s valid for QR only.
	SparseSubfactorQ SparseSubfactor_t = 6
	// SparseSubfactorR: An R factor subfactor that’s valid for QR and Cholesky AᵀA only.
	SparseSubfactorR SparseSubfactor_t = 7
	// SparseSubfactorRP: A half-solve subfactor that’s valid for QR and Cholesky AᵀA only.
	SparseSubfactorRP SparseSubfactor_t = 8
	// SparseSubfactorS: A diagonal scaling subfactor that’s valid for Cholesky and LDLᵀ only.
	SparseSubfactorS SparseSubfactor_t = 2
	// SparseSubfactorSc: Types of sub-factor object.
	SparseSubfactorSc SparseSubfactor_t = 10
	// SparseSubfactorSr: Types of sub-factor object.
	SparseSubfactorSr SparseSubfactor_t = 9
)

func (e SparseSubfactor_t) String() string {
	switch e {
	case SparseSubfactorD:
		return "SparseSubfactorD"
	case SparseSubfactorInvalid:
		return "SparseSubfactorInvalid"
	case SparseSubfactorL:
		return "SparseSubfactorL"
	case SparseSubfactorP:
		return "SparseSubfactorP"
	case SparseSubfactorPLPS:
		return "SparseSubfactorPLPS"
	case SparseSubfactorQ:
		return "SparseSubfactorQ"
	case SparseSubfactorR:
		return "SparseSubfactorR"
	case SparseSubfactorRP:
		return "SparseSubfactorRP"
	case SparseSubfactorS:
		return "SparseSubfactorS"
	case SparseSubfactorSc:
		return "SparseSubfactorSc"
	case SparseSubfactorSr:
		return "SparseSubfactorSr"
	default:
		return fmt.Sprintf("SparseSubfactor_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseTriangle_t
type SparseTriangle_t uint8

const (
	// SparseLowerTriangle: A constant that specifies the lower triangle.
	SparseLowerTriangle SparseTriangle_t = 1
	// SparseUpperTriangle: A constant that specifies the upper triangle.
	SparseUpperTriangle SparseTriangle_t = 0
)

func (e SparseTriangle_t) String() string {
	switch e {
	case SparseLowerTriangle:
		return "SparseLowerTriangle"
	case SparseUpperTriangle:
		return "SparseUpperTriangle"
	default:
		return fmt.Sprintf("SparseTriangle_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/SparseUpdate_t
type SparseUpdate_t uint8

const (
	// SparseUpdatePartialRefactor: Low-rank update algorithm selector
	SparseUpdatePartialRefactor SparseUpdate_t = 0
)

func (e SparseUpdate_t) String() string {
	switch e {
	case SparseUpdatePartialRefactor:
		return "SparseUpdatePartialRefactor"
	default:
		return fmt.Sprintf("SparseUpdate_t(%d)", e)
	}
}

type VdspHa uint32

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

// See: https://developer.apple.com/documentation/Accelerate/quadrature_integrator
type Quadrature_integrator uint32

const (
	// QUADRATURE_INTEGRATE_QAG: A constant that specifies a simple globally adaptive integrator.
	QUADRATURE_INTEGRATE_QAG Quadrature_integrator = 1
	// QUADRATURE_INTEGRATE_QAGS: A constant that specifies global adaptive quadrature.
	QUADRATURE_INTEGRATE_QAGS Quadrature_integrator = 2
	// QUADRATURE_INTEGRATE_QNG: A constant that specifies a simple non-adaptive automatic integrator.
	QUADRATURE_INTEGRATE_QNG Quadrature_integrator = 0
)

func (e Quadrature_integrator) String() string {
	switch e {
	case QUADRATURE_INTEGRATE_QAG:
		return "QUADRATURE_INTEGRATE_QAG"
	case QUADRATURE_INTEGRATE_QAGS:
		return "QUADRATURE_INTEGRATE_QAGS"
	case QUADRATURE_INTEGRATE_QNG:
		return "QUADRATURE_INTEGRATE_QNG"
	default:
		return fmt.Sprintf("Quadrature_integrator(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/quadrature_status
type Quadrature_status int32

const (
	// QUADRATURE_ALLOC_ERROR: A constant that indicates that memory allocation failed.
	QUADRATURE_ALLOC_ERROR Quadrature_status = -3
	// QUADRATURE_ERROR: A constant that indicates that a generic error occurred.
	QUADRATURE_ERROR Quadrature_status = -1
	// QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR: A constant that indicates bad integrand behaviour, or that an excessive roundoff error occurred.
	QUADRATURE_INTEGRATE_BAD_BEHAVIOUR_ERROR Quadrature_status = -102
	// QUADRATURE_INTEGRATE_MAX_EVAL_ERROR: A constant that indicates that the requested accuracy limit could not be reached.
	QUADRATURE_INTEGRATE_MAX_EVAL_ERROR Quadrature_status = -101
	// QUADRATURE_INTERNAL_ERROR: A constant that indicates that an internal error occurred.
	QUADRATURE_INTERNAL_ERROR Quadrature_status = -99
	// QUADRATURE_INVALID_ARG_ERROR: A constant that indicates that an invalid argument was passed to the operation.
	QUADRATURE_INVALID_ARG_ERROR Quadrature_status = -2
	// QUADRATURE_SUCCESS: A constant that indicates that the Quadrature operation was successful.
	QUADRATURE_SUCCESS Quadrature_status = 0
)

func (e Quadrature_status) String() string {
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
		return fmt.Sprintf("Quadrature_status(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_property
type Sparse_matrix_property uint32

const (
	// SPARSE_LOWER_SYMMETRIC: A symmetric matrix with values derived from the lower triangle.
	SPARSE_LOWER_SYMMETRIC Sparse_matrix_property = 8
	// SPARSE_LOWER_TRIANGULAR: A lower triangular matrix.
	SPARSE_LOWER_TRIANGULAR Sparse_matrix_property = 2
	// SPARSE_UPPER_SYMMETRIC: A symmetric matrix with values derived from the upper triangle.
	SPARSE_UPPER_SYMMETRIC Sparse_matrix_property = 4
	// SPARSE_UPPER_TRIANGULAR: An upper triangular matrix.
	SPARSE_UPPER_TRIANGULAR Sparse_matrix_property = 1
)

func (e Sparse_matrix_property) String() string {
	switch e {
	case SPARSE_LOWER_SYMMETRIC:
		return "SPARSE_LOWER_SYMMETRIC"
	case SPARSE_LOWER_TRIANGULAR:
		return "SPARSE_LOWER_TRIANGULAR"
	case SPARSE_UPPER_SYMMETRIC:
		return "SPARSE_UPPER_SYMMETRIC"
	case SPARSE_UPPER_TRIANGULAR:
		return "SPARSE_UPPER_TRIANGULAR"
	default:
		return fmt.Sprintf("Sparse_matrix_property(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/sparse_norm
type Sparse_norm uint32

const (
	// SPARSE_NORM_INF: Norm Inf
	SPARSE_NORM_INF Sparse_norm = 175
	// SPARSE_NORM_ONE: Norm One
	SPARSE_NORM_ONE Sparse_norm = 171
	// SPARSE_NORM_R1: Norm R1
	SPARSE_NORM_R1 Sparse_norm = 179
	// SPARSE_NORM_TWO: Norm Two
	SPARSE_NORM_TWO Sparse_norm = 173
)

func (e Sparse_norm) String() string {
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
		return fmt.Sprintf("Sparse_norm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/sparse_status
type Sparse_status int32

const (
	// SPARSE_CANNOT_SET_PROPERTY: A property was set after values were inserted into the matrix.
	SPARSE_CANNOT_SET_PROPERTY Sparse_status = -1001
	// SPARSE_ILLEGAL_PARAMETER: Operation was not completed because one or more of the arguments had an illegal value.
	SPARSE_ILLEGAL_PARAMETER Sparse_status = -1000
	// SPARSE_SUCCESS: Operation was a success.
	SPARSE_SUCCESS Sparse_status = 0
	// SPARSE_SYSTEM_ERROR: An internal error has occured, such as non enough memory.
	SPARSE_SYSTEM_ERROR Sparse_status = -1002
)

func (e Sparse_status) String() string {
	switch e {
	case SPARSE_CANNOT_SET_PROPERTY:
		return "SPARSE_CANNOT_SET_PROPERTY"
	case SPARSE_ILLEGAL_PARAMETER:
		return "SPARSE_ILLEGAL_PARAMETER"
	case SPARSE_SUCCESS:
		return "SPARSE_SUCCESS"
	case SPARSE_SYSTEM_ERROR:
		return "SPARSE_SYSTEM_ERROR"
	default:
		return fmt.Sprintf("Sparse_status(%d)", e)
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
	VDSP_DFT_FORWARD VDSP_DFT_Direction = 1
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
	VDSP_DFT_Interleaved_RealtoComplex    VDSP_DFT_RealtoComplex = -1
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

// See: https://developer.apple.com/documentation/Accelerate/vImageARGBType
type VImageARGBType uint32

const (
	// KvImageARGB16Q12: Any 8-bit four-channel interleaved buffer.
	KvImageARGB16Q12 VImageARGBType = 2
	// KvImageARGB16U: Any 16-bit unsigned, four-channel interleaved buffer.
	KvImageARGB16U VImageARGBType = 1
	// KvImageARGB8888: Any 16-bit signed fixed-point, four-channel interleaved buffer.
	KvImageARGB8888 VImageARGBType = 0
)

func (e VImageARGBType) String() string {
	switch e {
	case KvImageARGB16Q12:
		return "KvImageARGB16Q12"
	case KvImageARGB16U:
		return "KvImageARGB16U"
	case KvImageARGB8888:
		return "KvImageARGB8888"
	default:
		return fmt.Sprintf("VImageARGBType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vImageMDTableUsageHint
type VImageMDTableUsageHint uint32

const (
	// KvImageMDTableHint_16Q12: A table for transforming 16Q12 data.
	KvImageMDTableHint_16Q12 VImageMDTableUsageHint = 1
	// KvImageMDTableHint_Float: A table for transforming floating-point data.
	KvImageMDTableHint_Float VImageMDTableUsageHint = 2
)

func (e VImageMDTableUsageHint) String() string {
	switch e {
	case KvImageMDTableHint_16Q12:
		return "KvImageMDTableHint_16Q12"
	case KvImageMDTableHint_Float:
		return "KvImageMDTableHint_Float"
	default:
		return fmt.Sprintf("VImageMDTableUsageHint(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vImageYpCbCrType
type VImageYpCbCrType uint32

const (
	// KvImage420Yp8_Cb8_Cr8: Any y420 or f420 (planar component Y’CbCr 8-bit 4:2:0) buffer.
	KvImage420Yp8_Cb8_Cr8 VImageYpCbCrType = 3
	// KvImage420Yp8_CbCr8: Any 420v or 420f (biplanar component Y’CbCr 8-bit 4:2:0, video-range) buffer.
	KvImage420Yp8_CbCr8 VImageYpCbCrType = 4
	// KvImage422CbYpCrYp16: Any v216 (component Y’CbCr 10,12,14,16-bit 4:2:2) buffer.
	KvImage422CbYpCrYp16 VImageYpCbCrType = 13
	// KvImage422CbYpCrYp8: Any 2vuy (component Y’CbCr 8-bit 4:2:2) buffer.
	KvImage422CbYpCrYp8 VImageYpCbCrType = 0
	// KvImage422CbYpCrYp8_AA8: Any a2vy (first plane: video-range component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1; second plane: alpha 8-bit) buffer.
	KvImage422CbYpCrYp8_AA8 VImageYpCbCrType = 2
	// KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10: Any v210 (component Y’CbCr 10-bit 4:2:2) buffer.
	KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10 VImageYpCbCrType = 9
	// KvImage422YpCbYpCr8: Any yuvs or yuvf (component Y’CbCr 8-bit 4:2:2, ordered Y’0 Cb Y’1 Cr) buffer.
	KvImage422YpCbYpCr8 VImageYpCbCrType = 1
	// KvImage444AYpCbCr16: Any y416 (component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr) buffer.
	KvImage444AYpCbCr16 VImageYpCbCrType = 14
	// KvImage444AYpCbCr8: Any r408 or y408 (component Y’CbCrA 8-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr) buffer.
	KvImage444AYpCbCr8 VImageYpCbCrType = 5
	// KvImage444CbYpCrA8: Any v408 (component Y’CbCrA 8-bit 4:4:4:4) buffer.
	KvImage444CbYpCrA8 VImageYpCbCrType = 7
	// KvImage444CrYpCb10: Any v410 (component Y’CbCr 10-bit 4:4:4) buffer.
	KvImage444CrYpCb10 VImageYpCbCrType = 8
	// KvImage444CrYpCb8: Any v308 (component Y’CbCr 8-bit 4:4:4) buffer.
	KvImage444CrYpCb8 VImageYpCbCrType = 6
)

func (e VImageYpCbCrType) String() string {
	switch e {
	case KvImage420Yp8_Cb8_Cr8:
		return "KvImage420Yp8_Cb8_Cr8"
	case KvImage420Yp8_CbCr8:
		return "KvImage420Yp8_CbCr8"
	case KvImage422CbYpCrYp16:
		return "KvImage422CbYpCrYp16"
	case KvImage422CbYpCrYp8:
		return "KvImage422CbYpCrYp8"
	case KvImage422CbYpCrYp8_AA8:
		return "KvImage422CbYpCrYp8_AA8"
	case KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10:
		return "KvImage422CrYpCbYpCbYpCbYpCrYpCrYp10"
	case KvImage422YpCbYpCr8:
		return "KvImage422YpCbYpCr8"
	case KvImage444AYpCbCr16:
		return "KvImage444AYpCbCr16"
	case KvImage444AYpCbCr8:
		return "KvImage444AYpCbCr8"
	case KvImage444CbYpCrA8:
		return "KvImage444CbYpCrA8"
	case KvImage444CrYpCb10:
		return "KvImage444CrYpCb10"
	case KvImage444CrYpCb8:
		return "KvImage444CrYpCb8"
	default:
		return fmt.Sprintf("VImageYpCbCrType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accelerate/vImage_InterpolationMethod
type VImage_InterpolationMethod uint32

const (
	// KvImageFullInterpolation: Full linear interpolation.
	KvImageFullInterpolation VImage_InterpolationMethod = 1
	// KvImageHalfInterpolation: Partial linear interpolation.
	KvImageHalfInterpolation VImage_InterpolationMethod = 2
	// KvImageNoInterpolation: Nearest neighbor interpolation.
	KvImageNoInterpolation VImage_InterpolationMethod = 0
)

func (e VImage_InterpolationMethod) String() string {
	switch e {
	case KvImageFullInterpolation:
		return "KvImageFullInterpolation"
	case KvImageHalfInterpolation:
		return "KvImageHalfInterpolation"
	case KvImageNoInterpolation:
		return "KvImageNoInterpolation"
	default:
		return fmt.Sprintf("VImage_InterpolationMethod(%d)", e)
	}
}

// BlasThreading is a Go-name alias for BLAS_THREADING.
type BlasThreading = BLAS_THREADING

// CblasDiag is a Go-name alias for CBLAS_DIAG.
type CblasDiag = CBLAS_DIAG

// CblasOrder is a Go-name alias for CBLAS_ORDER.
type CblasOrder = CBLAS_ORDER

// CblasSide is a Go-name alias for CBLAS_SIDE.
type CblasSide = CBLAS_SIDE

// CblasTranspose is a Go-name alias for CBLAS_TRANSPOSE.
type CblasTranspose = CBLAS_TRANSPOSE

// CblasUplo is a Go-name alias for CBLAS_UPLO.
type CblasUplo = CBLAS_UPLO

// SparseControl is a Go-name alias for SparseControl_t.
type SparseControl = SparseControl_t

// SparseFactorization is a Go-name alias for SparseFactorization_t.
type SparseFactorization = SparseFactorization_t

// SparseGMRESVariant is a Go-name alias for SparseGMRESVariant_t.
type SparseGMRESVariant = SparseGMRESVariant_t

// SparseIterativeStatus is a Go-name alias for SparseIterativeStatus_t.
type SparseIterativeStatus = SparseIterativeStatus_t

// SparseKind is a Go-name alias for SparseKind_t.
type SparseKind = SparseKind_t

// SparseLSMRConvergenceTest is a Go-name alias for SparseLSMRConvergenceTest_t.
type SparseLSMRConvergenceTest = SparseLSMRConvergenceTest_t

// SparseOrder is a Go-name alias for SparseOrder_t.
type SparseOrder = SparseOrder_t

// SparsePreconditioner is a Go-name alias for SparsePreconditioner_t.
type SparsePreconditioner = SparsePreconditioner_t

// SparseScaling is a Go-name alias for SparseScaling_t.
type SparseScaling = SparseScaling_t

// SparseStatus is a Go-name alias for SparseStatus_t.
type SparseStatus = SparseStatus_t

// SparseSubfactor is a Go-name alias for SparseSubfactor_t.
type SparseSubfactor = SparseSubfactor_t

// SparseTriangle is a Go-name alias for SparseTriangle_t.
type SparseTriangle = SparseTriangle_t

// SparseUpdate is a Go-name alias for SparseUpdate_t.
type SparseUpdate = SparseUpdate_t

// QuadratureIntegrator is a Go-name alias for Quadrature_integrator.
type QuadratureIntegrator = Quadrature_integrator

// QuadratureStatus is a Go-name alias for Quadrature_status.
type QuadratureStatus = Quadrature_status

// SparseMatrixProperty is a Go-name alias for Sparse_matrix_property.
type SparseMatrixProperty = Sparse_matrix_property

// SparseNorm is a Go-name alias for Sparse_norm.
type SparseNorm = Sparse_norm

// VdspDctType is a Go-name alias for VDSP_DCT_Type.
type VdspDctType = VDSP_DCT_Type

// VdspDftDirection is a Go-name alias for VDSP_DFT_Direction.
type VdspDftDirection = VDSP_DFT_Direction

// VdspDftRealtoComplex is a Go-name alias for VDSP_DFT_RealtoComplex.
type VdspDftRealtoComplex = VDSP_DFT_RealtoComplex

// VImageInterpolationMethod is a Go-name alias for VImage_InterpolationMethod.
type VImageInterpolationMethod = VImage_InterpolationMethod
