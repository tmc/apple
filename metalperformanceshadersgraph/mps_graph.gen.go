// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraph] class.
var (
	_MPSGraphClass     MPSGraphClass
	_MPSGraphClassOnce sync.Once
)

func getMPSGraphClass() MPSGraphClass {
	_MPSGraphClassOnce.Do(func() {
		_MPSGraphClass = MPSGraphClass{class: objc.GetClass("MPSGraph")}
	})
	return _MPSGraphClass
}

// GetMPSGraphClass returns the class object for MPSGraph.
func GetMPSGraphClass() MPSGraphClass {
	return getMPSGraphClass()
}

type MPSGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphClass) Alloc() MPSGraph {
	rv := objc.Send[MPSGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The optimized representation of a compute graph of operations and tensors.
//
// # Overview
//
// An MPSGraph is a symbolic representation of operations to be utilized to
// execute compute graphs on a device.
//
// # Instance Properties
//
//   - [MPSGraph.Options]: Options for the graph.
//   - [MPSGraph.SetOptions]
//   - [MPSGraph.PlaceholderTensors]: Array of all the placeholder tensors.
//
// # Instance Methods
//
//   - [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [MPSGraph.HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName]: Computes the hamming distance of two input tensors with support for broadcasting.
//   - [MPSGraph.HermiteanToRealFFTWithTensorAxesDescriptorName]: Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.HermiteanToRealFFTWithTensorAxesTensorDescriptorName]: Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.L2NormPooling4DWithSourceTensorDescriptorName]: Creates a 4D L2-norm pooling operation and returns the result tensor.
//   - [MPSGraph.L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a L2-Norm pooling gradient operation and returns the result tensor.
//   - [MPSGraph.LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [MPSGraph.AbsoluteWithTensorName]: Returns the absolute values of the input tensor elements.
//   - [MPSGraph.AbsoluteSquareWithTensorName]: Returns the absolute square of the input tensor elements.
//   - [MPSGraph.AcosWithTensorName]: Applies the inverse cosine operation to the input tensor elements.
//   - [MPSGraph.AcoshWithTensorName]: Applies the inverse hyperbolic cosine operation to the input tensor elements.
//   - [MPSGraph.AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName]: Creates operations to apply Adam optimization.
//   - [MPSGraph.AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName]: Creates operations to apply Adam optimization.
//   - [MPSGraph.AdditionWithPrimaryTensorSecondaryTensorName]: Adds two input tensors.
//   - [MPSGraph.ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName]: The Stochastic gradient descent performs a gradient descent `variable = variable - (learningRate * g)` where, `g` is gradient of error wrt variable this op directly writes to the variable
//   - [MPSGraph.ArgSortWithTensorAxisDescendingName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [MPSGraph.ArgSortWithTensorAxisName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [MPSGraph.ArgSortWithTensorAxisTensorDescendingName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [MPSGraph.ArgSortWithTensorAxisTensorName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [MPSGraph.AsinWithTensorName]: Applies the inverse sine operation to the input tensor elements.
//   - [MPSGraph.AsinhWithTensorName]: Applies the inverse hyperbolic sine operation to the input tensor elements.
//   - [MPSGraph.AssignVariableWithValueOfTensorName]: Creates an assign operation which writes at this point of execution of the graph.
//   - [MPSGraph.AtanWithTensorName]: Applies the inverse tangent operation to the input tensor elements.
//   - [MPSGraph.Atan2WithPrimaryTensorSecondaryTensorName]: Returns the elementwise two-argument arctangent of the input tensors.
//   - [MPSGraph.AtanhWithTensorName]: Applies the inverse hyperbolic tangent operation to the input tensor elements.
//   - [MPSGraph.AvgPooling2DWithSourceTensorDescriptorName]: Creates a 2D average-pooling operation and returns the result tensor.
//   - [MPSGraph.AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a 2D average pooling gradient operation and returns the result tensor.
//   - [MPSGraph.AvgPooling4DWithSourceTensorDescriptorName]: Creates a 4D average pooling operation and returns the result tensor.
//   - [MPSGraph.AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates an average pooling gradient operation and returns the result tensor.
//   - [MPSGraph.BandPartWithTensorNumLowerNumUpperName]: Computes the band part of an input tensor.
//   - [MPSGraph.BandPartWithTensorNumLowerTensorNumUpperTensorName]: Creates the band part operation and returns the result.
//   - [MPSGraph.BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName]: Creates a batch-to-space operation and returns the result tensor.
//   - [MPSGraph.BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName]: Creates a batch-to-space operation and returns the result tensor.
//   - [MPSGraph.BitwiseANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise AND of binary representations of two integer tensors.
//   - [MPSGraph.BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName]: Returns the elementwise left-shifted binary representations of the primary integer by the secondary tensor amount.
//   - [MPSGraph.BitwiseNOTWithTensorName]: Applies the bitwise NOT operation to the input tensor element.
//   - [MPSGraph.BitwiseORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise OR of binary representations of two integer tensors.
//   - [MPSGraph.BitwisePopulationCountWithTensorName]: Returns the population count of the input tensor elements.
//   - [MPSGraph.BitwiseRightShiftWithPrimaryTensorSecondaryTensorName]: Returns the elementwise right-shifted binary representations of the primary integer by the secondary tensor amount.
//   - [MPSGraph.BitwiseXORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise XOR of binary representations of two integer tensors.
//   - [MPSGraph.BottomKWithSourceTensorAxisKName]: Creates a BottomK operation and returns the value and indices tensors.
//   - [MPSGraph.BottomKWithSourceTensorAxisTensorKTensorName]: Creates a BottomK operation and returns the result tensor.
//   - [MPSGraph.BottomKWithGradientTensorSourceAxisKName]: Creates a BottomKGradient operation and returns the result tensor.
//   - [MPSGraph.BottomKWithGradientTensorSourceAxisTensorKTensorName]: Creates a BottomKGradient operation and returns the result tensor.
//   - [MPSGraph.BroadcastTensorToShapeName]: Creates a broadcast operation and returns the result tensor.
//   - [MPSGraph.BroadcastTensorToShapeTensorName]: Creates a broadcast operation and returns the result tensor.
//   - [MPSGraph.CallSymbolNameInputTensorsOutputTypesName]: Creates an operation which invokes another executable.
//   - [MPSGraph.CastTensorToTypeName]: Creates a cast operation and returns the result tensor.
//   - [MPSGraph.CeilWithTensorName]: Applies the ceiling operation to the input tensor elements.
//   - [MPSGraph.ClampWithTensorMinValueTensorMaxValueTensorName]: Clamps the values in the first tensor between the corresponding values in the minimum and maximum value tensor.
//   - [MPSGraph.ColToImWithSourceTensorOutputShapeDescriptorName]: Creates a column to image operation and returns the result tensor.
//   - [MPSGraph.CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor]: Compiles the graph for the given feeds to returns the target tensor values, ensuring all target operations would be executed.
//   - [MPSGraph.ConstantWithRealPartImaginaryPart]: Creates a complex constant op with the MPSDataTypeComplexFloat32 data type and returns the result tensor.
//   - [MPSGraph.ConstantWithRealPartImaginaryPartDataType]: Creates a complex constant operation and returns the result tensor.
//   - [MPSGraph.ConstantWithRealPartImaginaryPartShapeDataType]: Creates a complex constant op with a given shape and returns the result tensor.
//   - [MPSGraph.ComplexTensorWithRealTensorImaginaryTensorName]: Returns a complex tensor from the two input tensors.
//   - [MPSGraph.ConcatTensorWithTensorDimensionName]: Creates a concatenation operation and returns the result tensor.
//   - [MPSGraph.ConcatTensorsDimensionInterleaveName]: Creates a concatenation operation and returns the result tensor.
//   - [MPSGraph.ConcatTensorsDimensionName]: Creates a concatenation operation and returns the result tensor.
//   - [MPSGraph.ConjugateWithTensorName]: Returns the complex conjugate of the input tensor elements.
//   - [MPSGraph.ConstantWithScalarDataType]: Creates a constant operation and returns the result tensor.
//   - [MPSGraph.ConstantWithScalarShapeDataType]: Creates a constant op with a given shape and returns the result tensor.
//   - [MPSGraph.ConstantWithDataShapeDataType]: Creates a constant op with a given shape and data, and returns the result tensor.
//   - [MPSGraph.Convolution2DWithSourceTensorWeightsTensorDescriptorName]: Creates a 2D (forward) convolution operation and returns the result tensor.
//   - [MPSGraph.Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [MPSGraph.Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [MPSGraph.Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [MPSGraph.Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to weights tensor of forward convolution.
//   - [MPSGraph.Convolution3DWithSourceTensorWeightsTensorDescriptorName]: Creates a 3D forward convolution operation and returns the result tensor.
//   - [MPSGraph.Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [MPSGraph.Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [MPSGraph.Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [MPSGraph.Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [MPSGraph.ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName]: Creates a convolution transpose operation and returns the result tensor.
//   - [MPSGraph.ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName]: Creates a convolution transpose operation and returns the result tensor.
//   - [MPSGraph.ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//   - [MPSGraph.ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//   - [MPSGraph.ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//   - [MPSGraph.ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//   - [MPSGraph.CoordinateAlongAxisWithShapeName]: Creates a get-coordindate operation and returns the result tensor.
//   - [MPSGraph.CoordinateAlongAxisWithShapeTensorName]: Creates a get-coordindate operation and returns the result tensor.
//   - [MPSGraph.CoordinateAlongAxisTensorWithShapeName]: Creates a get-coordindate operation and returns the result tensor.
//   - [MPSGraph.CoordinateAlongAxisTensorWithShapeTensorName]: Creates a get-coordindate operation and returns the result tensor.
//   - [MPSGraph.CosWithTensorName]: Applies the cosine operation to the input tensor elements.
//   - [MPSGraph.CoshWithTensorName]: Applies the hyperbolic cosine operation to the input tensor elements.
//   - [MPSGraph.CumulativeMaximumWithTensorAxisExclusiveReverseName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMaximumWithTensorAxisName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMaximumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMaximumWithTensorAxisTensorName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMinimumWithTensorAxisExclusiveReverseName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMinimumWithTensorAxisName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMinimumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeMinimumWithTensorAxisTensorName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeProductWithTensorAxisExclusiveReverseName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeProductWithTensorAxisName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeProductWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeProductWithTensorAxisTensorName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeSumWithTensorAxisExclusiveReverseName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeSumWithTensorAxisName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeSumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [MPSGraph.CumulativeSumWithTensorAxisTensorName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [MPSGraph.DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName]: Creates a depth-to-space2D operation and returns the result tensor.
//   - [MPSGraph.DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName]: Creates a depth-to-space2D operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName]: Creates a 2D-depthwise convolution operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName]: Creates a 2D-depthwise convolution gradient for data operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]: Creates a 2D-depthwise convolution gradient for weights operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName]: Creates a 3D depthwise convolution operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName]: Creates a 3D depthwise convolution gradient for data operation and returns the result tensor.
//   - [MPSGraph.DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]: Creates a 3D depthwise convolution gradient for weights operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorLUTTensorAxisName]: Creates a vector lookup-table based quantization operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorLUTTensorName]: Creates a lookup-table based quantization operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorScaleZeroPointDataTypeName]: Creates Dequantize operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorScaleTensorDataTypeName]: Creates a dequantize operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorScaleTensorZeroPointDataTypeAxisName]: Creates Dequantize operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName]: Creates a dequantize operation and returns the result tensor.
//   - [MPSGraph.DequantizeTensorScaleTensorZeroPointTensorDataTypeName]: Creates a dequantize operation and returns the result tensor.
//   - [MPSGraph.DivisionWithPrimaryTensorSecondaryTensorName]: Divides the first input tensor by the second.
//   - [MPSGraph.DivisionNoNaNWithPrimaryTensorSecondaryTensorName]: Divides the first input tensor by the second, with the result being 0 if the denominator is 0.
//   - [MPSGraph.DropoutTensorRateTensorName]: Creates a dropout operation and returns the result
//   - [MPSGraph.DropoutTensorRateName]: Creates a dropout operation and returns the result
//   - [MPSGraph.EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//   - [MPSGraph.EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraph.EqualWithPrimaryTensorSecondaryTensorName]: Returns the elementwise equality check of the input tensors.
//   - [MPSGraph.ErfWithTensorName]: Applies the error function to the input tensor elements.
//   - [MPSGraph.ExpandDimsOfTensorAxesName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [MPSGraph.ExpandDimsOfTensorAxesTensorName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [MPSGraph.ExpandDimsOfTensorAxisName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [MPSGraph.ExponentWithTensorName]: Applies the natural exponent to the input tensor elements.
//   - [MPSGraph.ExponentBase10WithTensorName]: Applies an exponent with base 10 to the input tensor elements.
//   - [MPSGraph.ExponentBase2WithTensorName]: Applies an exponent with base 2 to the input tensor elements.
//   - [MPSGraph.FastFourierTransformWithTensorAxesDescriptorName]: Creates a fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.FastFourierTransformWithTensorAxesTensorDescriptorName]: Creates a fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.Flatten2DTensorAxisName]: Creates a flatten2D operation and returns the result tensor.
//   - [MPSGraph.Flatten2DTensorAxisTensorName]: Creates a flatten2D operation and returns the result tensor.
//   - [MPSGraph.FloorWithTensorName]: Applies the floor operation to the input tensor elements.
//   - [MPSGraph.FloorModuloWithPrimaryTensorSecondaryTensorName]: Returns the remainder of floor divison between the primary and secondary tensor.
//   - [MPSGraph.GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName]: Creates a Gather operation and returns the result tensor.
//   - [MPSGraph.GatherAlongAxisWithUpdatesTensorIndicesTensorName]: Creates a GatherAlongAxis operation and returns the result tensor.
//   - [MPSGraph.GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName]: Creates a GatherAlongAxis operation and returns the result tensor.
//   - [MPSGraph.GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName]: Creates a GatherND operation and returns the result tensor.
//   - [MPSGraph.GradientForPrimaryTensorWithTensorsName]: Calculates a partial derivative of primaryTensor with respect to the tensors.
//   - [MPSGraph.GreaterThanWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is greater than the second.
//   - [MPSGraph.GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is greater than or equal to the second.
//   - [MPSGraph.IdentityWithTensorName]: Copies the input tensor values into the output, behaving as an identity operation.
//   - [MPSGraph.ImToColWithSourceTensorDescriptorName]: Creates an imToCol operation and returns the result tensor.
//   - [MPSGraph.ImaginaryPartOfTensorName]: Returns the imaginary part of a tensor.
//   - [MPSGraph.InverseOfTensorName]: Computes the inverse of an input tensor.
//   - [MPSGraph.IsFiniteWithTensorName]: Checks if the input tensor elements are finite or not.
//   - [MPSGraph.IsInfiniteWithTensorName]: Checks if the input tensor elements are infinite or not.
//   - [MPSGraph.IsNaNWithTensorName]: Checks if the input tensor elements are [NaN] or not.
//   - [MPSGraph.LeakyReLUWithTensorAlphaName]: Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//   - [MPSGraph.LeakyReLUWithTensorAlphaTensorName]: Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//   - [MPSGraph.LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName]: Computes the gradient of the leaky rectified linear unit (ReLU) activation.
//   - [MPSGraph.LessThanWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is less than the second.
//   - [MPSGraph.LessThanOrEqualToWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is less than or equal to the second.
//   - [MPSGraph.LogarithmWithTensorName]: Computes the natural logarithm to the input tensor elements.
//   - [MPSGraph.LogarithmBase10WithTensorName]: Computes the logarithm with base 10 to the input tensor elements.
//   - [MPSGraph.LogarithmBase2WithTensorName]: Computes the logarithm with base 2 to the input tensor elements.
//   - [MPSGraph.LogicalANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical AND of the input tensors.
//   - [MPSGraph.LogicalNANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical NAND of the input tensors.
//   - [MPSGraph.LogicalNORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical NOR of the input tensors.
//   - [MPSGraph.LogicalORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical OR of the input tensors.
//   - [MPSGraph.LogicalXNORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical XNOR of the input tensors.
//   - [MPSGraph.LogicalXORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical XOR of the input tensors.
//   - [MPSGraph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName]: Computes the matrix multiplication of 2 input tensors with support for broadcasting.
//   - [MPSGraph.MaxPooling2DWithSourceTensorDescriptorName]: Creates a 2D max-pooling operation and returns the result tensor.
//   - [MPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName]: Creates a 2D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//   - [MPSGraph.MaxPooling4DWithSourceTensorDescriptorName]: Creates a 4D max-pooling operation and returns the result tensor.
//   - [MPSGraph.MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName]: Creates a 4D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//   - [MPSGraph.MaximumWithPrimaryTensorSecondaryTensorName]: Returns the elementwise maximum of the input tensors.
//   - [MPSGraph.MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName]: Returns the elementwise maximum of the input tensors, while propagating [NaN] values.
//   - [MPSGraph.MeanOfTensorAxesName]: Returns the mean of the first input along the specified axes.
//   - [MPSGraph.MinimumWithPrimaryTensorSecondaryTensorName]: Returns the elementwise minimum of the input tensors.
//   - [MPSGraph.MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName]: Returns the elementwise minimum of the input tensors, while propagating [NaN] values.
//   - [MPSGraph.ModuloWithPrimaryTensorSecondaryTensorName]: Returns the remainder obtained by dividing the first input tensor by the second.
//   - [MPSGraph.MultiplicationWithPrimaryTensorSecondaryTensorName]: Multiplies two input tensors.
//   - [MPSGraph.NegativeWithTensorName]: Applies negative to the input tensor elements.
//   - [MPSGraph.NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName]: Creates a nonMaximumumSuppression operation and returns the result tensor.
//   - [MPSGraph.NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName]: Creates a nonMaximumumSuppression operation and returns the result tensor.
//   - [MPSGraph.NonZeroIndicesOfTensorName]: Computes the indices of the non-zero elements of the input tensor.
//   - [MPSGraph.NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName]: Creates a normalization beta-gradient operation and returns the result tensor.
//   - [MPSGraph.NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName]: Creates a normalization gamma-gradient operation and returns the result tensor.
//   - [MPSGraph.NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName]: Creates a normalization input gradient operation and returns the result tensor.
//   - [MPSGraph.NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName]: Creates a batch normalization operation and returns the result tensor.
//   - [MPSGraph.NotWithTensorName]: Applies the logical NOT operation to the input tensor elements.
//   - [MPSGraph.NotEqualWithPrimaryTensorSecondaryTensorName]: Returns the elementwise inequality check of the input tensors.
//   - [MPSGraph.OneHotWithIndicesTensorDepthAxisDataTypeName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.OneHotWithIndicesTensorDepthAxisName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.OneHotWithIndicesTensorDepthDataTypeName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.OneHotWithIndicesTensorDepthName]: Creates a oneHot operation and returns the result tensor.
//   - [MPSGraph.PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName]: Creates a padding gradient operation and returns the result tensor.
//   - [MPSGraph.PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName]: Creates a padding operation and returns the result tensor.
//   - [MPSGraph.PlaceholderWithShapeDataTypeName]: Creates a placeholder operation and returns the result tensor.
//   - [MPSGraph.PlaceholderWithShapeName]: Creates a placeholder operation and returns the result tensor with the dataType of the placeholder tensor set to 32 bit float.
//   - [MPSGraph.PowerWithPrimaryTensorSecondaryTensorName]: Returns the elementwise result of raising the first tensor to the power of the second tensor.
//   - [MPSGraph.QuantizeTensorScaleZeroPointDataTypeName]: Creates a Quantize operation and returns the result tensor.
//   - [MPSGraph.QuantizeTensorScaleTensorZeroPointDataTypeAxisName]: Creates a Quantize operation and returns the result tensor.
//   - [MPSGraph.QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName]: Creates a Quantize operation and returns the result tensor.
//   - [MPSGraph.RandomPhiloxStateTensorWithCounterLowCounterHighKeyName]: Creates a tensor representing state using the Philox algorithm with given counter and key values.
//   - [MPSGraph.RandomPhiloxStateTensorWithSeedName]: Creates a tensor representing state using the Philox algorithm with given counter and key values.
//   - [MPSGraph.RandomTensorWithShapeDescriptorName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [MPSGraph.RandomTensorWithShapeDescriptorSeedName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [MPSGraph.RandomTensorWithShapeDescriptorStateTensorName]: Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//   - [MPSGraph.RandomTensorWithShapeTensorDescriptorName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [MPSGraph.RandomTensorWithShapeTensorDescriptorSeedName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [MPSGraph.RandomTensorWithShapeTensorDescriptorStateTensorName]: Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//   - [MPSGraph.RandomUniformTensorWithShapeName]: Creates a RandomUniform operation and returns random uniform values
//   - [MPSGraph.RandomUniformTensorWithShapeSeedName]: Creates a RandomUniform operation and returns random uniform values
//   - [MPSGraph.RandomUniformTensorWithShapeStateTensorName]: Creates a RandomUniform operation and returns random uniform values and updated state
//   - [MPSGraph.RandomUniformTensorWithShapeTensorName]: Creates a RandomUniform operation and returns random uniform values
//   - [MPSGraph.RandomUniformTensorWithShapeTensorSeedName]: Creates a RandomUniform operation and returns random uniform values
//   - [MPSGraph.RandomUniformTensorWithShapeTensorStateTensorName]: Creates a RandomUniform operation and returns random uniform values and updated state
//   - [MPSGraph.ReLUWithTensorName]: Computes the ReLU (rectified linear activation unit) function with the input tensor.
//   - [MPSGraph.ReLUGradientWithIncomingGradientSourceTensorName]: Computes the gradient of the ReLU  (rectified linear activation unit) function using the incoming gradient.
//   - [MPSGraph.ReadVariableName]: Creates a read op which reads at this point of execution of the graph and returns the result tensor.
//   - [MPSGraph.RealPartOfTensorName]: Returns the real part of a tensor.
//   - [MPSGraph.RealToHermiteanFFTWithTensorAxesDescriptorName]: Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.RealToHermiteanFFTWithTensorAxesTensorDescriptorName]: Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//   - [MPSGraph.ReciprocalWithTensorName]: Applies the reciprocal operation to the input tensor elements.
//   - [MPSGraph.ReciprocalSquareRootWithTensorName]: Applies the reciprocal square root operation to the input tensor elements.
//   - [MPSGraph.ReductionAndWithTensorAxesName]: Creates a reduction and operation and returns the result tensor.
//   - [MPSGraph.ReductionAndWithTensorAxisName]: Creates a reduction and operation and returns the result tensor.
//   - [MPSGraph.ReductionArgMaximumWithTensorAxisName]: Creates a reduction argMax operation and returns the result tensor.
//   - [MPSGraph.ReductionArgMinimumWithTensorAxisName]: Creates a reduction argMin operation and returns the result tensor.
//   - [MPSGraph.ReductionMaximumWithTensorAxesName]: Creates a reduction max operation and returns the result tensor.
//   - [MPSGraph.ReductionMaximumWithTensorAxisName]: Creates a reduction max operation and returns the result tensor.
//   - [MPSGraph.ReductionMaximumPropagateNaNWithTensorAxesName]: Creates a reduction max propagate NaN operation and returns the result tensor.
//   - [MPSGraph.ReductionMaximumPropagateNaNWithTensorAxisName]: Creates a reduction max propagate NaN operation and returns the result tensor.
//   - [MPSGraph.ReductionMinimumWithTensorAxesName]: Creates a reduction min operation and returns the result tensor.
//   - [MPSGraph.ReductionMinimumWithTensorAxisName]: Creates a reduction minimum operation and returns the result tensor.
//   - [MPSGraph.ReductionMinimumPropagateNaNWithTensorAxesName]: Creates a reduction min propagate NaN operation and returns the result tensor.
//   - [MPSGraph.ReductionMinimumPropagateNaNWithTensorAxisName]: Creates a reduction min propagate NaN operation and returns the result tensor.
//   - [MPSGraph.ReductionOrWithTensorAxesName]: Creates a reduction or operation and returns the result tensor.
//   - [MPSGraph.ReductionOrWithTensorAxisName]: Creates a reduction or operation and returns the result tensor.
//   - [MPSGraph.ReductionProductWithTensorAxesName]: Creates a reduction product operation and returns the result tensor.
//   - [MPSGraph.ReductionProductWithTensorAxisName]: Creates a reduction product operation and returns the result tensor.
//   - [MPSGraph.ReductionSumWithTensorAxesName]: Creates a reduction sum operation and returns the result tensor.
//   - [MPSGraph.ReductionSumWithTensorAxisName]: Creates a reduction sum operation and returns the result tensor.
//   - [MPSGraph.ReinterpretCastTensorToTypeName]: Creates a reinterpret cast operation and returns the result tensor.
//   - [MPSGraph.ReshapeTensorWithShapeName]: Creates a reshape operation and returns the result tensor.
//   - [MPSGraph.ReshapeTensorWithShapeTensorName]: Creates a reshape operation and returns the result tensor.
//   - [MPSGraph.ResizeTensorSizeModeCenterResultAlignCornersLayoutName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeTensorSizeTensorModeCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName]: Resamples input images to given size using the provided scale and offset. Destination indices are computed using
//   - [MPSGraph.ResizeTensorSizeTensorScaleTensorOffsetTensorModeName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName]: Resamples input images to given size using bilinear sampling.
//   - [MPSGraph.ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName]: Resamples input images to given size using the provided scale and offset and bilinear sampling See above discussion for more details.
//   - [MPSGraph.ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName]: Resamples input images to given size using nearest neighbor sampling.
//   - [MPSGraph.ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName]: Resamples input images to given size using the provided scale and offset and nearest neighbor sampling See above discussion for more details.
//   - [MPSGraph.ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName]: Creates a Resize operation and returns the result tensor.
//   - [MPSGraph.ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [MPSGraph.ReverseTensorAxesName]: Creates a reverse operation and returns the result tensor.
//   - [MPSGraph.ReverseTensorAxesTensorName]: Creates a reverse operation and returns the result tensor.
//   - [MPSGraph.ReverseTensorName]: Creates a reverse operation and returns the result tensor.
//   - [MPSGraph.RintWithTensorName]: Rounds the input tensor elements by rounding to nearest even.
//   - [MPSGraph.RoundWithTensorName]: Rounds the input tensor elements.
//   - [MPSGraph.RunWithFeedsTargetTensorsTargetOperations]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary]: Runs the graph for the given feeds and returns the target tensor values in the results dictionary provided by the user.
//   - [MPSGraph.RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraph.RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraph.RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//   - [MPSGraph.RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [MPSGraph.SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName]: Samples a tensor using the coordinates provided, using nearest neighbor sampling with specified rounding mode.
//   - [MPSGraph.SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName]: Samples a tensor using the coordinates provided.
//   - [MPSGraph.ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName]: Creates a scaled dot product attention (SDPA) operation and returns the result tensor.
//   - [MPSGraph.ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName]: Creates a scaled dot product attention (SDPA) operation (without a mask) and returns the result tensor.
//   - [MPSGraph.ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName]: Creates a Scatter operation and returns the result tensor.
//   - [MPSGraph.ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [MPSGraph.ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [MPSGraph.ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [MPSGraph.ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [MPSGraph.ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName]: Creates a ScatterND operation and returns the result tensor.
//   - [MPSGraph.ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName]: Creates a ScatterND operation and returns the result tensor.
//   - [MPSGraph.ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName]: Creates a ScatterND operation and returns the result tensor.
//   - [MPSGraph.ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName]: Creates a Scatter operation and returns the result tensor.
//   - [MPSGraph.SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName]: Selects values from either the true or false predicate tensor, depending on the values in the first input.
//   - [MPSGraph.ShapeOfTensorName]: Creates a shape-of operation and returns the result tensor.
//   - [MPSGraph.SigmoidWithTensorName]: Computes the sigmoid operation on an input tensor.
//   - [MPSGraph.SigmoidGradientWithIncomingGradientSourceTensorName]: Computes the gradient of the sigmoid function using the incoming gradient tensor.
//   - [MPSGraph.SignWithTensorName]: Returns the sign of the input tensor elements.
//   - [MPSGraph.SignbitWithTensorName]: Returns the sign bit of the input tensor elements.
//   - [MPSGraph.SinWithTensorName]: Applies the sine operation to the input tensor elements.
//   - [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [MPSGraph.SinhWithTensorName]: Applies the hyperbolic sine operation to the input tensor elements.
//   - [MPSGraph.SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [MPSGraph.SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName]: Creates a slice gradient operation and returns the result tensor.
//   - [MPSGraph.SliceGradientTensorFwdInShapeTensorStartsEndsStridesName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [MPSGraph.SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [MPSGraph.SliceTensorDimensionStartLengthName]: Creates a slice operation and returns the result tensor.
//   - [MPSGraph.SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice operation and returns the result tensor.
//   - [MPSGraph.SliceTensorStartTensorSizeTensorSqueezeMaskName]: Creates a slice operation and returns the result tensor.
//   - [MPSGraph.SliceTensorStartsEndsStridesName]: Creates a strided-slice operation and returns the result tensor.
//   - [MPSGraph.SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice operation and returns the result tensor.
//   - [MPSGraph.SliceUpdateDataTensorUpdateTensorStartsEndsStridesName]: Creates a strided-slice update operation with zero masks and returns the result tensor.
//   - [MPSGraph.SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice update operation and returns the result tensor.
//   - [MPSGraph.SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName]: Creates a strided-slice update operation with zero masks and returns the result tensor.
//   - [MPSGraph.SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice update operation and returns the result tensor.
//   - [MPSGraph.SoftMaxWithTensorAxisName]: Computes the softmax function on the input tensor along the specified axis.
//   - [MPSGraph.SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName]: Creates a softmax cross-entropy loss operation and returns the result tensor.
//   - [MPSGraph.SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName]: Creates the gradient of a softmax cross-entropy loss operation and returns the result tensor.
//   - [MPSGraph.SoftMaxGradientWithIncomingGradientSourceTensorAxisName]: Computes the gradient of the softmax function along the specified axis using the incoming gradient tensor.
//   - [MPSGraph.SortWithTensorAxisDescendingName]: Sorts the elements of the input tensor along the specified axis.
//   - [MPSGraph.SortWithTensorAxisName]: Sorts the elements of the input tensor along the specified axis.
//   - [MPSGraph.SortWithTensorAxisTensorDescendingName]: Sorts the elements of the input tensor along the specified axis.
//   - [MPSGraph.SortWithTensorAxisTensorName]: Sorts the elements of the input tensor along the specified axis.
//   - [MPSGraph.SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName]: Creates a space-to-depth2D operation and returns the result tensor.
//   - [MPSGraph.SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName]: Creates a space-to-depth2D operation and returns the result tensor.
//   - [MPSGraph.SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName]: Creates a space-to-batch operation and returns the result tensor.
//   - [MPSGraph.SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName]: Creates a space-to-batch operation and returns the result tensor.
//   - [MPSGraph.SparseTensorWithDescriptorTensorsShapeName]: Creates a sparse tensor representation.
//   - [MPSGraph.SparseTensorWithTypeTensorsShapeDataTypeName]: Creates a sparse tensor representation.
//   - [MPSGraph.SplitTensorNumSplitsAxisName]: Creates a split operation and returns the result tensor.
//   - [MPSGraph.SplitTensorSplitSizesAxisName]: Creates a split operation and returns the result tensor.
//   - [MPSGraph.SplitTensorSplitSizesTensorAxisName]: Creates a split operation and returns the result tensor.
//   - [MPSGraph.SquareWithTensorName]: Applies the square operation to the input tensor elements.
//   - [MPSGraph.SquareRootWithTensorName]: Applies the square root operation to the input tensor elements.
//   - [MPSGraph.SqueezeTensorAxesName]: Creates a squeeze operation and returns the result tensor.
//   - [MPSGraph.SqueezeTensorAxesTensorName]: Creates a squeeze operation and returns the result tensor.
//   - [MPSGraph.SqueezeTensorAxisName]: Creates a squeeze operation and returns the result tensor.
//   - [MPSGraph.SqueezeTensorName]: Creates a squeeze operation and returns the result tensor.
//   - [MPSGraph.StackTensorsAxisName]: Creates a stack operation and returns the result tensor.
//   - [MPSGraph.StencilWithSourceTensorWeightsTensorDescriptorName]: Creates a stencil operation and returns the result tensor.
//   - [MPSGraph.StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName]: The Stochastic gradient descent performs a gradient descent.
//   - [MPSGraph.SubtractionWithPrimaryTensorSecondaryTensorName]: Subtracts the second input tensor from the first.
//   - [MPSGraph.TanWithTensorName]: Applies the tangent operation to the input tensor elements.
//   - [MPSGraph.TanhWithTensorName]: Applies the hyperbolic tangent operation to the input tensor elements.
//   - [MPSGraph.TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName]: Creates a tile gradient operation and returns the result tensor.
//   - [MPSGraph.TileTensorWithMultiplierName]: Creates a tile operation and returns the result tensor.
//   - [MPSGraph.TopKWithSourceTensorAxisKName]: Creates a TopK operation and returns the value and indices tensors.
//   - [MPSGraph.TopKWithSourceTensorAxisTensorKTensorName]: Creates a TopK operation and returns the result tensor.
//   - [MPSGraph.TopKWithSourceTensorKName]: Creates a TopK operation and returns the value and indices tensors
//   - [MPSGraph.TopKWithSourceTensorKTensorName]: Creates a TopK operation and returns the result tensor.
//   - [MPSGraph.TopKWithGradientTensorSourceKName]: Creates a TopKGradient operation and returns the result tensor.
//   - [MPSGraph.TopKWithGradientTensorSourceKTensorName]: Creates a TopKGradient operation and returns the result tensor.
//   - [MPSGraph.TopKWithGradientTensorSourceAxisKName]: Creates a TopKGradient operation and returns the result tensor.
//   - [MPSGraph.TopKWithGradientTensorSourceAxisTensorKTensorName]: Creates a TopKGradient operation and returns the result tensor.
//   - [MPSGraph.TransposeTensorPermutationName]: Creates a permutation operation and returns the result tensor.
//   - [MPSGraph.TransposeTensorDimensionWithDimensionName]: Creates a transpose operation and returns the result tensor.
//   - [MPSGraph.TruncateWithTensorName]: Applies the truncate operation to the input tensor elements.
//   - [MPSGraph.VariableWithDataShapeDataTypeName]: Creates a variable operation and returns the result tensor.
//   - [MPSGraph.VariableFromTensorWithTensorName]: Creates a variable from an input tensor.
//   - [MPSGraph.VarianceOfTensorAxesName]: Returns the variance of the first input along the specified axes.
//   - [MPSGraph.VarianceOfTensorMeanTensorAxesName]: Returns the variance of the first input along the specified axes when the mean has been precomputed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph
type MPSGraph struct {
	MPSGraphObject
}

// MPSGraphFromID constructs a [MPSGraph] from an objc.ID.
//
// The optimized representation of a compute graph of operations and tensors.
func MPSGraphFromID(id objc.ID) MPSGraph {
	return MPSGraph{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraph] class.
//
// # Instance Properties
//
//   - [IMPSGraph.Options]: Options for the graph.
//   - [IMPSGraph.SetOptions]
//   - [IMPSGraph.PlaceholderTensors]: Array of all the placeholder tensors.
//
// # Instance Methods
//
//   - [IMPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]: Creates a GRU operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]: Creates a GRU gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName]: Computes the hamming distance of two input tensors with support for broadcasting.
//   - [IMPSGraph.HermiteanToRealFFTWithTensorAxesDescriptorName]: Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.HermiteanToRealFFTWithTensorAxesTensorDescriptorName]: Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.L2NormPooling4DWithSourceTensorDescriptorName]: Creates a 4D L2-norm pooling operation and returns the result tensor.
//   - [IMPSGraph.L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a L2-Norm pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [IMPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [IMPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]: Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
//   - [IMPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]: Creates an LSTM gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.AbsoluteWithTensorName]: Returns the absolute values of the input tensor elements.
//   - [IMPSGraph.AbsoluteSquareWithTensorName]: Returns the absolute square of the input tensor elements.
//   - [IMPSGraph.AcosWithTensorName]: Applies the inverse cosine operation to the input tensor elements.
//   - [IMPSGraph.AcoshWithTensorName]: Applies the inverse hyperbolic cosine operation to the input tensor elements.
//   - [IMPSGraph.AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName]: Creates operations to apply Adam optimization.
//   - [IMPSGraph.AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName]: Creates operations to apply Adam optimization.
//   - [IMPSGraph.AdditionWithPrimaryTensorSecondaryTensorName]: Adds two input tensors.
//   - [IMPSGraph.ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName]: The Stochastic gradient descent performs a gradient descent `variable = variable - (learningRate * g)` where, `g` is gradient of error wrt variable this op directly writes to the variable
//   - [IMPSGraph.ArgSortWithTensorAxisDescendingName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [IMPSGraph.ArgSortWithTensorAxisName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [IMPSGraph.ArgSortWithTensorAxisTensorDescendingName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [IMPSGraph.ArgSortWithTensorAxisTensorName]: Computes the indices that sort the elements of the input tensor along the specified axis.
//   - [IMPSGraph.AsinWithTensorName]: Applies the inverse sine operation to the input tensor elements.
//   - [IMPSGraph.AsinhWithTensorName]: Applies the inverse hyperbolic sine operation to the input tensor elements.
//   - [IMPSGraph.AssignVariableWithValueOfTensorName]: Creates an assign operation which writes at this point of execution of the graph.
//   - [IMPSGraph.AtanWithTensorName]: Applies the inverse tangent operation to the input tensor elements.
//   - [IMPSGraph.Atan2WithPrimaryTensorSecondaryTensorName]: Returns the elementwise two-argument arctangent of the input tensors.
//   - [IMPSGraph.AtanhWithTensorName]: Applies the inverse hyperbolic tangent operation to the input tensor elements.
//   - [IMPSGraph.AvgPooling2DWithSourceTensorDescriptorName]: Creates a 2D average-pooling operation and returns the result tensor.
//   - [IMPSGraph.AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a 2D average pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.AvgPooling4DWithSourceTensorDescriptorName]: Creates a 4D average pooling operation and returns the result tensor.
//   - [IMPSGraph.AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates an average pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.BandPartWithTensorNumLowerNumUpperName]: Computes the band part of an input tensor.
//   - [IMPSGraph.BandPartWithTensorNumLowerTensorNumUpperTensorName]: Creates the band part operation and returns the result.
//   - [IMPSGraph.BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName]: Creates a batch-to-space operation and returns the result tensor.
//   - [IMPSGraph.BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName]: Creates a batch-to-space operation and returns the result tensor.
//   - [IMPSGraph.BitwiseANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise AND of binary representations of two integer tensors.
//   - [IMPSGraph.BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName]: Returns the elementwise left-shifted binary representations of the primary integer by the secondary tensor amount.
//   - [IMPSGraph.BitwiseNOTWithTensorName]: Applies the bitwise NOT operation to the input tensor element.
//   - [IMPSGraph.BitwiseORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise OR of binary representations of two integer tensors.
//   - [IMPSGraph.BitwisePopulationCountWithTensorName]: Returns the population count of the input tensor elements.
//   - [IMPSGraph.BitwiseRightShiftWithPrimaryTensorSecondaryTensorName]: Returns the elementwise right-shifted binary representations of the primary integer by the secondary tensor amount.
//   - [IMPSGraph.BitwiseXORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise bitwise XOR of binary representations of two integer tensors.
//   - [IMPSGraph.BottomKWithSourceTensorAxisKName]: Creates a BottomK operation and returns the value and indices tensors.
//   - [IMPSGraph.BottomKWithSourceTensorAxisTensorKTensorName]: Creates a BottomK operation and returns the result tensor.
//   - [IMPSGraph.BottomKWithGradientTensorSourceAxisKName]: Creates a BottomKGradient operation and returns the result tensor.
//   - [IMPSGraph.BottomKWithGradientTensorSourceAxisTensorKTensorName]: Creates a BottomKGradient operation and returns the result tensor.
//   - [IMPSGraph.BroadcastTensorToShapeName]: Creates a broadcast operation and returns the result tensor.
//   - [IMPSGraph.BroadcastTensorToShapeTensorName]: Creates a broadcast operation and returns the result tensor.
//   - [IMPSGraph.CallSymbolNameInputTensorsOutputTypesName]: Creates an operation which invokes another executable.
//   - [IMPSGraph.CastTensorToTypeName]: Creates a cast operation and returns the result tensor.
//   - [IMPSGraph.CeilWithTensorName]: Applies the ceiling operation to the input tensor elements.
//   - [IMPSGraph.ClampWithTensorMinValueTensorMaxValueTensorName]: Clamps the values in the first tensor between the corresponding values in the minimum and maximum value tensor.
//   - [IMPSGraph.ColToImWithSourceTensorOutputShapeDescriptorName]: Creates a column to image operation and returns the result tensor.
//   - [IMPSGraph.CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor]: Compiles the graph for the given feeds to returns the target tensor values, ensuring all target operations would be executed.
//   - [IMPSGraph.ConstantWithRealPartImaginaryPart]: Creates a complex constant op with the MPSDataTypeComplexFloat32 data type and returns the result tensor.
//   - [IMPSGraph.ConstantWithRealPartImaginaryPartDataType]: Creates a complex constant operation and returns the result tensor.
//   - [IMPSGraph.ConstantWithRealPartImaginaryPartShapeDataType]: Creates a complex constant op with a given shape and returns the result tensor.
//   - [IMPSGraph.ComplexTensorWithRealTensorImaginaryTensorName]: Returns a complex tensor from the two input tensors.
//   - [IMPSGraph.ConcatTensorWithTensorDimensionName]: Creates a concatenation operation and returns the result tensor.
//   - [IMPSGraph.ConcatTensorsDimensionInterleaveName]: Creates a concatenation operation and returns the result tensor.
//   - [IMPSGraph.ConcatTensorsDimensionName]: Creates a concatenation operation and returns the result tensor.
//   - [IMPSGraph.ConjugateWithTensorName]: Returns the complex conjugate of the input tensor elements.
//   - [IMPSGraph.ConstantWithScalarDataType]: Creates a constant operation and returns the result tensor.
//   - [IMPSGraph.ConstantWithScalarShapeDataType]: Creates a constant op with a given shape and returns the result tensor.
//   - [IMPSGraph.ConstantWithDataShapeDataType]: Creates a constant op with a given shape and data, and returns the result tensor.
//   - [IMPSGraph.Convolution2DWithSourceTensorWeightsTensorDescriptorName]: Creates a 2D (forward) convolution operation and returns the result tensor.
//   - [IMPSGraph.Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [IMPSGraph.Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [IMPSGraph.Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [IMPSGraph.Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 2D convolution gradient operation with respect to weights tensor of forward convolution.
//   - [IMPSGraph.Convolution3DWithSourceTensorWeightsTensorDescriptorName]: Creates a 3D forward convolution operation and returns the result tensor.
//   - [IMPSGraph.Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [IMPSGraph.Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
//   - [IMPSGraph.Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [IMPSGraph.Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
//   - [IMPSGraph.ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName]: Creates a convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName]: Creates a convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName]: Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
//   - [IMPSGraph.CoordinateAlongAxisWithShapeName]: Creates a get-coordindate operation and returns the result tensor.
//   - [IMPSGraph.CoordinateAlongAxisWithShapeTensorName]: Creates a get-coordindate operation and returns the result tensor.
//   - [IMPSGraph.CoordinateAlongAxisTensorWithShapeName]: Creates a get-coordindate operation and returns the result tensor.
//   - [IMPSGraph.CoordinateAlongAxisTensorWithShapeTensorName]: Creates a get-coordindate operation and returns the result tensor.
//   - [IMPSGraph.CosWithTensorName]: Applies the cosine operation to the input tensor elements.
//   - [IMPSGraph.CoshWithTensorName]: Applies the hyperbolic cosine operation to the input tensor elements.
//   - [IMPSGraph.CumulativeMaximumWithTensorAxisExclusiveReverseName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMaximumWithTensorAxisName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMaximumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMaximumWithTensorAxisTensorName]: Computes the cumulative maximum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMinimumWithTensorAxisExclusiveReverseName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMinimumWithTensorAxisName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMinimumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeMinimumWithTensorAxisTensorName]: Computes the cumulative minimum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeProductWithTensorAxisExclusiveReverseName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeProductWithTensorAxisName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeProductWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeProductWithTensorAxisTensorName]: Computes the cumulative product of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeSumWithTensorAxisExclusiveReverseName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeSumWithTensorAxisName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeSumWithTensorAxisTensorExclusiveReverseName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [IMPSGraph.CumulativeSumWithTensorAxisTensorName]: Computes the cumulative sum of the input tensor along the specified axis.
//   - [IMPSGraph.DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName]: Creates a depth-to-space2D operation and returns the result tensor.
//   - [IMPSGraph.DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName]: Creates a depth-to-space2D operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName]: Creates a 2D-depthwise convolution operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName]: Creates a 2D-depthwise convolution gradient for data operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]: Creates a 2D-depthwise convolution gradient for weights operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName]: Creates a 3D depthwise convolution operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName]: Creates a 3D depthwise convolution gradient for data operation and returns the result tensor.
//   - [IMPSGraph.DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]: Creates a 3D depthwise convolution gradient for weights operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorLUTTensorAxisName]: Creates a vector lookup-table based quantization operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorLUTTensorName]: Creates a lookup-table based quantization operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorScaleZeroPointDataTypeName]: Creates Dequantize operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorScaleTensorDataTypeName]: Creates a dequantize operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorScaleTensorZeroPointDataTypeAxisName]: Creates Dequantize operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName]: Creates a dequantize operation and returns the result tensor.
//   - [IMPSGraph.DequantizeTensorScaleTensorZeroPointTensorDataTypeName]: Creates a dequantize operation and returns the result tensor.
//   - [IMPSGraph.DivisionWithPrimaryTensorSecondaryTensorName]: Divides the first input tensor by the second.
//   - [IMPSGraph.DivisionNoNaNWithPrimaryTensorSecondaryTensorName]: Divides the first input tensor by the second, with the result being 0 if the denominator is 0.
//   - [IMPSGraph.DropoutTensorRateTensorName]: Creates a dropout operation and returns the result
//   - [IMPSGraph.DropoutTensorRateName]: Creates a dropout operation and returns the result
//   - [IMPSGraph.EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//   - [IMPSGraph.EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraph.EqualWithPrimaryTensorSecondaryTensorName]: Returns the elementwise equality check of the input tensors.
//   - [IMPSGraph.ErfWithTensorName]: Applies the error function to the input tensor elements.
//   - [IMPSGraph.ExpandDimsOfTensorAxesName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [IMPSGraph.ExpandDimsOfTensorAxesTensorName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [IMPSGraph.ExpandDimsOfTensorAxisName]: Creates an expand-dimensions operation and returns the result tensor.
//   - [IMPSGraph.ExponentWithTensorName]: Applies the natural exponent to the input tensor elements.
//   - [IMPSGraph.ExponentBase10WithTensorName]: Applies an exponent with base 10 to the input tensor elements.
//   - [IMPSGraph.ExponentBase2WithTensorName]: Applies an exponent with base 2 to the input tensor elements.
//   - [IMPSGraph.FastFourierTransformWithTensorAxesDescriptorName]: Creates a fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.FastFourierTransformWithTensorAxesTensorDescriptorName]: Creates a fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.Flatten2DTensorAxisName]: Creates a flatten2D operation and returns the result tensor.
//   - [IMPSGraph.Flatten2DTensorAxisTensorName]: Creates a flatten2D operation and returns the result tensor.
//   - [IMPSGraph.FloorWithTensorName]: Applies the floor operation to the input tensor elements.
//   - [IMPSGraph.FloorModuloWithPrimaryTensorSecondaryTensorName]: Returns the remainder of floor divison between the primary and secondary tensor.
//   - [IMPSGraph.GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName]: Creates a Gather operation and returns the result tensor.
//   - [IMPSGraph.GatherAlongAxisWithUpdatesTensorIndicesTensorName]: Creates a GatherAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName]: Creates a GatherAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName]: Creates a GatherND operation and returns the result tensor.
//   - [IMPSGraph.GradientForPrimaryTensorWithTensorsName]: Calculates a partial derivative of primaryTensor with respect to the tensors.
//   - [IMPSGraph.GreaterThanWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is greater than the second.
//   - [IMPSGraph.GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is greater than or equal to the second.
//   - [IMPSGraph.IdentityWithTensorName]: Copies the input tensor values into the output, behaving as an identity operation.
//   - [IMPSGraph.ImToColWithSourceTensorDescriptorName]: Creates an imToCol operation and returns the result tensor.
//   - [IMPSGraph.ImaginaryPartOfTensorName]: Returns the imaginary part of a tensor.
//   - [IMPSGraph.InverseOfTensorName]: Computes the inverse of an input tensor.
//   - [IMPSGraph.IsFiniteWithTensorName]: Checks if the input tensor elements are finite or not.
//   - [IMPSGraph.IsInfiniteWithTensorName]: Checks if the input tensor elements are infinite or not.
//   - [IMPSGraph.IsNaNWithTensorName]: Checks if the input tensor elements are [NaN] or not.
//   - [IMPSGraph.LeakyReLUWithTensorAlphaName]: Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//   - [IMPSGraph.LeakyReLUWithTensorAlphaTensorName]: Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
//   - [IMPSGraph.LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName]: Computes the gradient of the leaky rectified linear unit (ReLU) activation.
//   - [IMPSGraph.LessThanWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is less than the second.
//   - [IMPSGraph.LessThanOrEqualToWithPrimaryTensorSecondaryTensorName]: Checks in an elementwise manner if the first input tensor is less than or equal to the second.
//   - [IMPSGraph.LogarithmWithTensorName]: Computes the natural logarithm to the input tensor elements.
//   - [IMPSGraph.LogarithmBase10WithTensorName]: Computes the logarithm with base 10 to the input tensor elements.
//   - [IMPSGraph.LogarithmBase2WithTensorName]: Computes the logarithm with base 2 to the input tensor elements.
//   - [IMPSGraph.LogicalANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical AND of the input tensors.
//   - [IMPSGraph.LogicalNANDWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical NAND of the input tensors.
//   - [IMPSGraph.LogicalNORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical NOR of the input tensors.
//   - [IMPSGraph.LogicalORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical OR of the input tensors.
//   - [IMPSGraph.LogicalXNORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical XNOR of the input tensors.
//   - [IMPSGraph.LogicalXORWithPrimaryTensorSecondaryTensorName]: Returns the elementwise logical XOR of the input tensors.
//   - [IMPSGraph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName]: Computes the matrix multiplication of 2 input tensors with support for broadcasting.
//   - [IMPSGraph.MaxPooling2DWithSourceTensorDescriptorName]: Creates a 2D max-pooling operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName]: Creates a 2D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//   - [IMPSGraph.MaxPooling4DWithSourceTensorDescriptorName]: Creates a 4D max-pooling operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]: Creates a max-pooling gradient operation and returns the result tensor.
//   - [IMPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName]: Creates a 4D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//   - [IMPSGraph.MaximumWithPrimaryTensorSecondaryTensorName]: Returns the elementwise maximum of the input tensors.
//   - [IMPSGraph.MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName]: Returns the elementwise maximum of the input tensors, while propagating [NaN] values.
//   - [IMPSGraph.MeanOfTensorAxesName]: Returns the mean of the first input along the specified axes.
//   - [IMPSGraph.MinimumWithPrimaryTensorSecondaryTensorName]: Returns the elementwise minimum of the input tensors.
//   - [IMPSGraph.MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName]: Returns the elementwise minimum of the input tensors, while propagating [NaN] values.
//   - [IMPSGraph.ModuloWithPrimaryTensorSecondaryTensorName]: Returns the remainder obtained by dividing the first input tensor by the second.
//   - [IMPSGraph.MultiplicationWithPrimaryTensorSecondaryTensorName]: Multiplies two input tensors.
//   - [IMPSGraph.NegativeWithTensorName]: Applies negative to the input tensor elements.
//   - [IMPSGraph.NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName]: Creates a nonMaximumumSuppression operation and returns the result tensor.
//   - [IMPSGraph.NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName]: Creates a nonMaximumumSuppression operation and returns the result tensor.
//   - [IMPSGraph.NonZeroIndicesOfTensorName]: Computes the indices of the non-zero elements of the input tensor.
//   - [IMPSGraph.NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName]: Creates a normalization beta-gradient operation and returns the result tensor.
//   - [IMPSGraph.NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName]: Creates a normalization gamma-gradient operation and returns the result tensor.
//   - [IMPSGraph.NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName]: Creates a normalization input gradient operation and returns the result tensor.
//   - [IMPSGraph.NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName]: Creates a batch normalization operation and returns the result tensor.
//   - [IMPSGraph.NotWithTensorName]: Applies the logical NOT operation to the input tensor elements.
//   - [IMPSGraph.NotEqualWithPrimaryTensorSecondaryTensorName]: Returns the elementwise inequality check of the input tensors.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthAxisDataTypeName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthAxisName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthDataTypeName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.OneHotWithIndicesTensorDepthName]: Creates a oneHot operation and returns the result tensor.
//   - [IMPSGraph.PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName]: Creates a padding gradient operation and returns the result tensor.
//   - [IMPSGraph.PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName]: Creates a padding operation and returns the result tensor.
//   - [IMPSGraph.PlaceholderWithShapeDataTypeName]: Creates a placeholder operation and returns the result tensor.
//   - [IMPSGraph.PlaceholderWithShapeName]: Creates a placeholder operation and returns the result tensor with the dataType of the placeholder tensor set to 32 bit float.
//   - [IMPSGraph.PowerWithPrimaryTensorSecondaryTensorName]: Returns the elementwise result of raising the first tensor to the power of the second tensor.
//   - [IMPSGraph.QuantizeTensorScaleZeroPointDataTypeName]: Creates a Quantize operation and returns the result tensor.
//   - [IMPSGraph.QuantizeTensorScaleTensorZeroPointDataTypeAxisName]: Creates a Quantize operation and returns the result tensor.
//   - [IMPSGraph.QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName]: Creates a Quantize operation and returns the result tensor.
//   - [IMPSGraph.RandomPhiloxStateTensorWithCounterLowCounterHighKeyName]: Creates a tensor representing state using the Philox algorithm with given counter and key values.
//   - [IMPSGraph.RandomPhiloxStateTensorWithSeedName]: Creates a tensor representing state using the Philox algorithm with given counter and key values.
//   - [IMPSGraph.RandomTensorWithShapeDescriptorName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [IMPSGraph.RandomTensorWithShapeDescriptorSeedName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [IMPSGraph.RandomTensorWithShapeDescriptorStateTensorName]: Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//   - [IMPSGraph.RandomTensorWithShapeTensorDescriptorName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [IMPSGraph.RandomTensorWithShapeTensorDescriptorSeedName]: Creates a Random op of type matching distribution in descriptor and returns random values.
//   - [IMPSGraph.RandomTensorWithShapeTensorDescriptorStateTensorName]: Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
//   - [IMPSGraph.RandomUniformTensorWithShapeName]: Creates a RandomUniform operation and returns random uniform values
//   - [IMPSGraph.RandomUniformTensorWithShapeSeedName]: Creates a RandomUniform operation and returns random uniform values
//   - [IMPSGraph.RandomUniformTensorWithShapeStateTensorName]: Creates a RandomUniform operation and returns random uniform values and updated state
//   - [IMPSGraph.RandomUniformTensorWithShapeTensorName]: Creates a RandomUniform operation and returns random uniform values
//   - [IMPSGraph.RandomUniformTensorWithShapeTensorSeedName]: Creates a RandomUniform operation and returns random uniform values
//   - [IMPSGraph.RandomUniformTensorWithShapeTensorStateTensorName]: Creates a RandomUniform operation and returns random uniform values and updated state
//   - [IMPSGraph.ReLUWithTensorName]: Computes the ReLU (rectified linear activation unit) function with the input tensor.
//   - [IMPSGraph.ReLUGradientWithIncomingGradientSourceTensorName]: Computes the gradient of the ReLU  (rectified linear activation unit) function using the incoming gradient.
//   - [IMPSGraph.ReadVariableName]: Creates a read op which reads at this point of execution of the graph and returns the result tensor.
//   - [IMPSGraph.RealPartOfTensorName]: Returns the real part of a tensor.
//   - [IMPSGraph.RealToHermiteanFFTWithTensorAxesDescriptorName]: Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.RealToHermiteanFFTWithTensorAxesTensorDescriptorName]: Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//   - [IMPSGraph.ReciprocalWithTensorName]: Applies the reciprocal operation to the input tensor elements.
//   - [IMPSGraph.ReciprocalSquareRootWithTensorName]: Applies the reciprocal square root operation to the input tensor elements.
//   - [IMPSGraph.ReductionAndWithTensorAxesName]: Creates a reduction and operation and returns the result tensor.
//   - [IMPSGraph.ReductionAndWithTensorAxisName]: Creates a reduction and operation and returns the result tensor.
//   - [IMPSGraph.ReductionArgMaximumWithTensorAxisName]: Creates a reduction argMax operation and returns the result tensor.
//   - [IMPSGraph.ReductionArgMinimumWithTensorAxisName]: Creates a reduction argMin operation and returns the result tensor.
//   - [IMPSGraph.ReductionMaximumWithTensorAxesName]: Creates a reduction max operation and returns the result tensor.
//   - [IMPSGraph.ReductionMaximumWithTensorAxisName]: Creates a reduction max operation and returns the result tensor.
//   - [IMPSGraph.ReductionMaximumPropagateNaNWithTensorAxesName]: Creates a reduction max propagate NaN operation and returns the result tensor.
//   - [IMPSGraph.ReductionMaximumPropagateNaNWithTensorAxisName]: Creates a reduction max propagate NaN operation and returns the result tensor.
//   - [IMPSGraph.ReductionMinimumWithTensorAxesName]: Creates a reduction min operation and returns the result tensor.
//   - [IMPSGraph.ReductionMinimumWithTensorAxisName]: Creates a reduction minimum operation and returns the result tensor.
//   - [IMPSGraph.ReductionMinimumPropagateNaNWithTensorAxesName]: Creates a reduction min propagate NaN operation and returns the result tensor.
//   - [IMPSGraph.ReductionMinimumPropagateNaNWithTensorAxisName]: Creates a reduction min propagate NaN operation and returns the result tensor.
//   - [IMPSGraph.ReductionOrWithTensorAxesName]: Creates a reduction or operation and returns the result tensor.
//   - [IMPSGraph.ReductionOrWithTensorAxisName]: Creates a reduction or operation and returns the result tensor.
//   - [IMPSGraph.ReductionProductWithTensorAxesName]: Creates a reduction product operation and returns the result tensor.
//   - [IMPSGraph.ReductionProductWithTensorAxisName]: Creates a reduction product operation and returns the result tensor.
//   - [IMPSGraph.ReductionSumWithTensorAxesName]: Creates a reduction sum operation and returns the result tensor.
//   - [IMPSGraph.ReductionSumWithTensorAxisName]: Creates a reduction sum operation and returns the result tensor.
//   - [IMPSGraph.ReinterpretCastTensorToTypeName]: Creates a reinterpret cast operation and returns the result tensor.
//   - [IMPSGraph.ReshapeTensorWithShapeName]: Creates a reshape operation and returns the result tensor.
//   - [IMPSGraph.ReshapeTensorWithShapeTensorName]: Creates a reshape operation and returns the result tensor.
//   - [IMPSGraph.ResizeTensorSizeModeCenterResultAlignCornersLayoutName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeTensorSizeTensorModeCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName]: Resamples input images to given size using the provided scale and offset. Destination indices are computed using
//   - [IMPSGraph.ResizeTensorSizeTensorScaleTensorOffsetTensorModeName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName]: Resamples input images to given size using bilinear sampling.
//   - [IMPSGraph.ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName]: Resamples input images to given size using the provided scale and offset and bilinear sampling See above discussion for more details.
//   - [IMPSGraph.ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName]: Resamples input images to given size using nearest neighbor sampling.
//   - [IMPSGraph.ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName]: Resamples input images to given size using the provided scale and offset and nearest neighbor sampling See above discussion for more details.
//   - [IMPSGraph.ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName]: Creates a Resize operation and returns the result tensor.
//   - [IMPSGraph.ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName]: Creates a Resize gradient operation and returns the result tensor.
//   - [IMPSGraph.ReverseTensorAxesName]: Creates a reverse operation and returns the result tensor.
//   - [IMPSGraph.ReverseTensorAxesTensorName]: Creates a reverse operation and returns the result tensor.
//   - [IMPSGraph.ReverseTensorName]: Creates a reverse operation and returns the result tensor.
//   - [IMPSGraph.RintWithTensorName]: Rounds the input tensor elements by rounding to nearest even.
//   - [IMPSGraph.RoundWithTensorName]: Rounds the input tensor elements.
//   - [IMPSGraph.RunWithFeedsTargetTensorsTargetOperations]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary]: Runs the graph for the given feeds and returns the target tensor values in the results dictionary provided by the user.
//   - [IMPSGraph.RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraph.RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraph.RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor]: Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
//   - [IMPSGraph.RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor]: Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//   - [IMPSGraph.SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName]: Samples a tensor using the coordinates provided, using nearest neighbor sampling with specified rounding mode.
//   - [IMPSGraph.SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName]: Samples a tensor using the coordinates provided.
//   - [IMPSGraph.ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName]: Creates a scaled dot product attention (SDPA) operation and returns the result tensor.
//   - [IMPSGraph.ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName]: Creates a scaled dot product attention (SDPA) operation (without a mask) and returns the result tensor.
//   - [IMPSGraph.ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName]: Creates a Scatter operation and returns the result tensor.
//   - [IMPSGraph.ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName]: Creates a ScatterAlongAxis operation and returns the result tensor.
//   - [IMPSGraph.ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName]: Creates a ScatterND operation and returns the result tensor.
//   - [IMPSGraph.ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName]: Creates a ScatterND operation and returns the result tensor.
//   - [IMPSGraph.ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName]: Creates a ScatterND operation and returns the result tensor.
//   - [IMPSGraph.ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName]: Creates a Scatter operation and returns the result tensor.
//   - [IMPSGraph.SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName]: Selects values from either the true or false predicate tensor, depending on the values in the first input.
//   - [IMPSGraph.ShapeOfTensorName]: Creates a shape-of operation and returns the result tensor.
//   - [IMPSGraph.SigmoidWithTensorName]: Computes the sigmoid operation on an input tensor.
//   - [IMPSGraph.SigmoidGradientWithIncomingGradientSourceTensorName]: Computes the gradient of the sigmoid function using the incoming gradient tensor.
//   - [IMPSGraph.SignWithTensorName]: Returns the sign of the input tensor elements.
//   - [IMPSGraph.SignbitWithTensorName]: Returns the sign bit of the input tensor elements.
//   - [IMPSGraph.SinWithTensorName]: Applies the sine operation to the input tensor elements.
//   - [IMPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//   - [IMPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName]: Creates a single-gate RNN gradient operation and returns the gradient tensor values.
//   - [IMPSGraph.SinhWithTensorName]: Applies the hyperbolic sine operation to the input tensor elements.
//   - [IMPSGraph.SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [IMPSGraph.SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName]: Creates a slice gradient operation and returns the result tensor.
//   - [IMPSGraph.SliceGradientTensorFwdInShapeTensorStartsEndsStridesName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [IMPSGraph.SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice gradient operation and returns the result tensor.
//   - [IMPSGraph.SliceTensorDimensionStartLengthName]: Creates a slice operation and returns the result tensor.
//   - [IMPSGraph.SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice operation and returns the result tensor.
//   - [IMPSGraph.SliceTensorStartTensorSizeTensorSqueezeMaskName]: Creates a slice operation and returns the result tensor.
//   - [IMPSGraph.SliceTensorStartsEndsStridesName]: Creates a strided-slice operation and returns the result tensor.
//   - [IMPSGraph.SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice operation and returns the result tensor.
//   - [IMPSGraph.SliceUpdateDataTensorUpdateTensorStartsEndsStridesName]: Creates a strided-slice update operation with zero masks and returns the result tensor.
//   - [IMPSGraph.SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice update operation and returns the result tensor.
//   - [IMPSGraph.SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName]: Creates a strided-slice update operation with zero masks and returns the result tensor.
//   - [IMPSGraph.SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName]: Creates a strided-slice update operation and returns the result tensor.
//   - [IMPSGraph.SoftMaxWithTensorAxisName]: Computes the softmax function on the input tensor along the specified axis.
//   - [IMPSGraph.SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName]: Creates a softmax cross-entropy loss operation and returns the result tensor.
//   - [IMPSGraph.SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName]: Creates the gradient of a softmax cross-entropy loss operation and returns the result tensor.
//   - [IMPSGraph.SoftMaxGradientWithIncomingGradientSourceTensorAxisName]: Computes the gradient of the softmax function along the specified axis using the incoming gradient tensor.
//   - [IMPSGraph.SortWithTensorAxisDescendingName]: Sorts the elements of the input tensor along the specified axis.
//   - [IMPSGraph.SortWithTensorAxisName]: Sorts the elements of the input tensor along the specified axis.
//   - [IMPSGraph.SortWithTensorAxisTensorDescendingName]: Sorts the elements of the input tensor along the specified axis.
//   - [IMPSGraph.SortWithTensorAxisTensorName]: Sorts the elements of the input tensor along the specified axis.
//   - [IMPSGraph.SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName]: Creates a space-to-depth2D operation and returns the result tensor.
//   - [IMPSGraph.SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName]: Creates a space-to-depth2D operation and returns the result tensor.
//   - [IMPSGraph.SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName]: Creates a space-to-batch operation and returns the result tensor.
//   - [IMPSGraph.SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName]: Creates a space-to-batch operation and returns the result tensor.
//   - [IMPSGraph.SparseTensorWithDescriptorTensorsShapeName]: Creates a sparse tensor representation.
//   - [IMPSGraph.SparseTensorWithTypeTensorsShapeDataTypeName]: Creates a sparse tensor representation.
//   - [IMPSGraph.SplitTensorNumSplitsAxisName]: Creates a split operation and returns the result tensor.
//   - [IMPSGraph.SplitTensorSplitSizesAxisName]: Creates a split operation and returns the result tensor.
//   - [IMPSGraph.SplitTensorSplitSizesTensorAxisName]: Creates a split operation and returns the result tensor.
//   - [IMPSGraph.SquareWithTensorName]: Applies the square operation to the input tensor elements.
//   - [IMPSGraph.SquareRootWithTensorName]: Applies the square root operation to the input tensor elements.
//   - [IMPSGraph.SqueezeTensorAxesName]: Creates a squeeze operation and returns the result tensor.
//   - [IMPSGraph.SqueezeTensorAxesTensorName]: Creates a squeeze operation and returns the result tensor.
//   - [IMPSGraph.SqueezeTensorAxisName]: Creates a squeeze operation and returns the result tensor.
//   - [IMPSGraph.SqueezeTensorName]: Creates a squeeze operation and returns the result tensor.
//   - [IMPSGraph.StackTensorsAxisName]: Creates a stack operation and returns the result tensor.
//   - [IMPSGraph.StencilWithSourceTensorWeightsTensorDescriptorName]: Creates a stencil operation and returns the result tensor.
//   - [IMPSGraph.StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName]: The Stochastic gradient descent performs a gradient descent.
//   - [IMPSGraph.SubtractionWithPrimaryTensorSecondaryTensorName]: Subtracts the second input tensor from the first.
//   - [IMPSGraph.TanWithTensorName]: Applies the tangent operation to the input tensor elements.
//   - [IMPSGraph.TanhWithTensorName]: Applies the hyperbolic tangent operation to the input tensor elements.
//   - [IMPSGraph.TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName]: Creates a tile gradient operation and returns the result tensor.
//   - [IMPSGraph.TileTensorWithMultiplierName]: Creates a tile operation and returns the result tensor.
//   - [IMPSGraph.TopKWithSourceTensorAxisKName]: Creates a TopK operation and returns the value and indices tensors.
//   - [IMPSGraph.TopKWithSourceTensorAxisTensorKTensorName]: Creates a TopK operation and returns the result tensor.
//   - [IMPSGraph.TopKWithSourceTensorKName]: Creates a TopK operation and returns the value and indices tensors
//   - [IMPSGraph.TopKWithSourceTensorKTensorName]: Creates a TopK operation and returns the result tensor.
//   - [IMPSGraph.TopKWithGradientTensorSourceKName]: Creates a TopKGradient operation and returns the result tensor.
//   - [IMPSGraph.TopKWithGradientTensorSourceKTensorName]: Creates a TopKGradient operation and returns the result tensor.
//   - [IMPSGraph.TopKWithGradientTensorSourceAxisKName]: Creates a TopKGradient operation and returns the result tensor.
//   - [IMPSGraph.TopKWithGradientTensorSourceAxisTensorKTensorName]: Creates a TopKGradient operation and returns the result tensor.
//   - [IMPSGraph.TransposeTensorPermutationName]: Creates a permutation operation and returns the result tensor.
//   - [IMPSGraph.TransposeTensorDimensionWithDimensionName]: Creates a transpose operation and returns the result tensor.
//   - [IMPSGraph.TruncateWithTensorName]: Applies the truncate operation to the input tensor elements.
//   - [IMPSGraph.VariableWithDataShapeDataTypeName]: Creates a variable operation and returns the result tensor.
//   - [IMPSGraph.VariableFromTensorWithTensorName]: Creates a variable from an input tensor.
//   - [IMPSGraph.VarianceOfTensorAxesName]: Returns the variance of the first input along the specified axes.
//   - [IMPSGraph.VarianceOfTensorMeanTensorAxesName]: Returns the variance of the first input along the specified axes when the mean has been precomputed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph
type IMPSGraph interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// Options for the graph.
	Options() MPSGraphOptions
	SetOptions(value MPSGraphOptions)
	// Array of all the placeholder tensors.
	PlaceholderTensors() []MPSGraphTensor

	// Topic: Instance Methods

	// Creates a GRU operation and returns the value and optionally the training state tensor.
	GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Creates a GRU operation and returns the value and optionally the training state tensor.
	GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Creates a GRU operation and returns the value and optionally the training state tensor.
	GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Creates a GRU gradient operation and returns the gradient tensor values.
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Creates a GRU gradient operation and returns the gradient tensor values.
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Creates a GRU gradient operation and returns the gradient tensor values.
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor
	// Computes the hamming distance of two input tensors with support for broadcasting.
	HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, resultDataType uint32, name string) IMPSGraphTensor
	// Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
	HermiteanToRealFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
	HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Creates a 4D L2-norm pooling operation and returns the result tensor.
	L2NormPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates a L2-Norm pooling gradient operation and returns the result tensor.
	L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
	LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
	LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and  the training state tensor.
	LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM gradient operation and returns the gradient tensor values.
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM gradient operation and returns the gradient tensor values.
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM gradient operation and returns the gradient tensor values.
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Creates an LSTM gradient operation and returns the gradient tensor values.
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, cellGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor
	// Returns the absolute values of the input tensor elements.
	AbsoluteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the absolute square of the input tensor elements.
	AbsoluteSquareWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the inverse cosine operation to the input tensor elements.
	AcosWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the inverse hyperbolic cosine operation to the input tensor elements.
	AcoshWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates operations to apply Adam optimization.
	AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(currentLearningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates operations to apply Adam optimization.
	AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(learningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, beta1PowerTensor IMPSGraphTensor, beta2PowerTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) []MPSGraphTensor
	// Adds two input tensors.
	AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// The Stochastic gradient descent performs a gradient descent `variable = variable - (learningRate * g)` where, `g` is gradient of error wrt variable this op directly writes to the variable
	ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor IMPSGraphTensor, variable IMPSGraphVariableOp, gradientTensor IMPSGraphTensor, name string) IMPSGraphOperation
	// Computes the indices that sort the elements of the input tensor along the specified axis.
	ArgSortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name string) IMPSGraphTensor
	// Computes the indices that sort the elements of the input tensor along the specified axis.
	ArgSortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Computes the indices that sort the elements of the input tensor along the specified axis.
	ArgSortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name string) IMPSGraphTensor
	// Computes the indices that sort the elements of the input tensor along the specified axis.
	ArgSortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the inverse sine operation to the input tensor elements.
	AsinWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the inverse hyperbolic sine operation to the input tensor elements.
	AsinhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates an assign operation which writes at this point of execution of the graph.
	AssignVariableWithValueOfTensorName(variable IMPSGraphTensor, tensor IMPSGraphTensor, name string) IMPSGraphOperation
	// Applies the inverse tangent operation to the input tensor elements.
	AtanWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise two-argument arctangent of the input tensors.
	Atan2WithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the inverse hyperbolic tangent operation to the input tensor elements.
	AtanhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a 2D average-pooling operation and returns the result tensor.
	AvgPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D average pooling gradient operation and returns the result tensor.
	AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 4D average pooling operation and returns the result tensor.
	AvgPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates an average pooling gradient operation and returns the result tensor.
	AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Computes the band part of an input tensor.
	BandPartWithTensorNumLowerNumUpperName(inputTensor IMPSGraphTensor, numLower int, numUpper int, name string) IMPSGraphTensor
	// Creates the band part operation and returns the result.
	BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor IMPSGraphTensor, numLowerTensor IMPSGraphTensor, numUpperTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a batch-to-space operation and returns the result tensor.
	BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.NSNumber, batchAxis int, blockDimensions []foundation.NSNumber, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a batch-to-space operation and returns the result tensor.
	BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Returns the elementwise bitwise AND of binary representations of two integer tensors.
	BitwiseANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise left-shifted binary representations of the primary integer by the secondary tensor amount.
	BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the bitwise NOT operation to the input tensor element.
	BitwiseNOTWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise bitwise OR of binary representations of two integer tensors.
	BitwiseORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the population count of the input tensor elements.
	BitwisePopulationCountWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise right-shifted binary representations of the primary integer by the secondary tensor amount.
	BitwiseRightShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise bitwise XOR of binary representations of two integer tensors.
	BitwiseXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a BottomK operation and returns the value and indices tensors.
	BottomKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name string) []MPSGraphTensor
	// Creates a BottomK operation and returns the result tensor.
	BottomKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a BottomKGradient operation and returns the result tensor.
	BottomKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name string) IMPSGraphTensor
	// Creates a BottomKGradient operation and returns the result tensor.
	BottomKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a broadcast operation and returns the result tensor.
	BroadcastTensorToShapeName(tensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a broadcast operation and returns the result tensor.
	BroadcastTensorToShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates an operation which invokes another executable.
	CallSymbolNameInputTensorsOutputTypesName(symbolName string, inputTensors []MPSGraphTensor, outputTypes []MPSGraphType, name string) []MPSGraphTensor
	// Creates a cast operation and returns the result tensor.
	CastTensorToTypeName(tensor IMPSGraphTensor, type_ uint32, name string) IMPSGraphTensor
	// Applies the ceiling operation to the input tensor elements.
	CeilWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Clamps the values in the first tensor between the corresponding values in the minimum and maximum value tensor.
	ClampWithTensorMinValueTensorMaxValueTensorName(tensor IMPSGraphTensor, minValueTensor IMPSGraphTensor, maxValueTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a column to image operation and returns the result tensor.
	ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphImToColOpDescriptor, name string) IMPSGraphTensor
	// Compiles the graph for the given feeds to returns the target tensor values, ensuring all target operations would be executed.
	CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IMPSGraphDevice, feeds MPSGraphTensorShapedTypeDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, compilationDescriptor IMPSGraphCompilationDescriptor) IMPSGraphExecutable
	// Creates a complex constant op with the MPSDataTypeComplexFloat32 data type and returns the result tensor.
	ConstantWithRealPartImaginaryPart(realPart float64, imaginaryPart float64) IMPSGraphTensor
	// Creates a complex constant operation and returns the result tensor.
	ConstantWithRealPartImaginaryPartDataType(realPart float64, imaginaryPart float64, dataType uint32) IMPSGraphTensor
	// Creates a complex constant op with a given shape and returns the result tensor.
	ConstantWithRealPartImaginaryPartShapeDataType(realPart float64, imaginaryPart float64, shape foundation.NSArray, dataType uint32) IMPSGraphTensor
	// Returns a complex tensor from the two input tensors.
	ComplexTensorWithRealTensorImaginaryTensorName(realTensor IMPSGraphTensor, imaginaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a concatenation operation and returns the result tensor.
	ConcatTensorWithTensorDimensionName(tensor IMPSGraphTensor, tensor2 IMPSGraphTensor, dimensionIndex int, name string) IMPSGraphTensor
	// Creates a concatenation operation and returns the result tensor.
	ConcatTensorsDimensionInterleaveName(tensors []MPSGraphTensor, dimensionIndex int, interleave bool, name string) IMPSGraphTensor
	// Creates a concatenation operation and returns the result tensor.
	ConcatTensorsDimensionName(tensors []MPSGraphTensor, dimensionIndex int, name string) IMPSGraphTensor
	// Returns the complex conjugate of the input tensor elements.
	ConjugateWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a constant operation and returns the result tensor.
	ConstantWithScalarDataType(scalar float64, dataType uint32) IMPSGraphTensor
	// Creates a constant op with a given shape and returns the result tensor.
	ConstantWithScalarShapeDataType(scalar float64, shape foundation.NSArray, dataType uint32) IMPSGraphTensor
	// Creates a constant op with a given shape and data, and returns the result tensor.
	ConstantWithDataShapeDataType(data foundation.NSData, shape foundation.NSArray, dataType uint32) IMPSGraphTensor
	// Creates a 2D (forward) convolution operation and returns the result tensor.
	Convolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
	Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D convolution gradient operation with respect to the source tensor of the forward convolution.
	Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D convolution gradient operation with respect to the weights tensor of the forward convolution.
	Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D convolution gradient operation with respect to weights tensor of forward convolution.
	Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D forward convolution operation and returns the result tensor.
	Convolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
	Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D convolution gradient operation with respect to the source tensor of the forward convolution.
	Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
	Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D convolution gradient operation with respect to the weights tensor of the forward convolution.
	Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose gradient operation with respect to the source tensor of convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a convolution transpose gradient operation with respect to the weights tensor of the convolution transpose operation and returns the result tensor.
	ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a get-coordindate operation and returns the result tensor.
	CoordinateAlongAxisWithShapeName(axis int, shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a get-coordindate operation and returns the result tensor.
	CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a get-coordindate operation and returns the result tensor.
	CoordinateAlongAxisTensorWithShapeName(axisTensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a get-coordindate operation and returns the result tensor.
	CoordinateAlongAxisTensorWithShapeTensorName(axisTensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the cosine operation to the input tensor elements.
	CosWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the hyperbolic cosine operation to the input tensor elements.
	CoshWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the cumulative maximum of the input tensor along the specified axis.
	CumulativeMaximumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative maximum of the input tensor along the specified axis.
	CumulativeMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Computes the cumulative maximum of the input tensor along the specified axis.
	CumulativeMaximumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative maximum of the input tensor along the specified axis.
	CumulativeMaximumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the cumulative minimum of the input tensor along the specified axis.
	CumulativeMinimumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative minimum of the input tensor along the specified axis.
	CumulativeMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Computes the cumulative minimum of the input tensor along the specified axis.
	CumulativeMinimumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative minimum of the input tensor along the specified axis.
	CumulativeMinimumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the cumulative product of the input tensor along the specified axis.
	CumulativeProductWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative product of the input tensor along the specified axis.
	CumulativeProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Computes the cumulative product of the input tensor along the specified axis.
	CumulativeProductWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative product of the input tensor along the specified axis.
	CumulativeProductWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the cumulative sum of the input tensor along the specified axis.
	CumulativeSumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative sum of the input tensor along the specified axis.
	CumulativeSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Computes the cumulative sum of the input tensor along the specified axis.
	CumulativeSumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor
	// Computes the cumulative sum of the input tensor along the specified axis.
	CumulativeSumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a depth-to-space2D operation and returns the result tensor.
	DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a depth-to-space2D operation and returns the result tensor.
	DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a 2D-depthwise convolution operation and returns the result tensor.
	DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D-depthwise convolution gradient for data operation and returns the result tensor.
	DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D-depthwise convolution gradient for weights operation and returns the result tensor.
	DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D depthwise convolution operation and returns the result tensor.
	DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D depthwise convolution gradient for data operation and returns the result tensor.
	DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 3D depthwise convolution gradient for weights operation and returns the result tensor.
	DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor
	// Creates a vector lookup-table based quantization operation and returns the result tensor.
	DequantizeTensorLUTTensorAxisName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a lookup-table based quantization operation and returns the result tensor.
	DequantizeTensorLUTTensorName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates Dequantize operation and returns the result tensor.
	DequantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType uint32, name string) IMPSGraphTensor
	// Creates a dequantize operation and returns the result tensor.
	DequantizeTensorScaleTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, dataType uint32, name string) IMPSGraphTensor
	// Creates Dequantize operation and returns the result tensor.
	DequantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType uint32, axis int, name string) IMPSGraphTensor
	// Creates a dequantize operation and returns the result tensor.
	DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, axis int, name string) IMPSGraphTensor
	// Creates a dequantize operation and returns the result tensor.
	DequantizeTensorScaleTensorZeroPointTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, name string) IMPSGraphTensor
	// Divides the first input tensor by the second.
	DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Divides the first input tensor by the second, with the result being 0 if the denominator is 0.
	DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a dropout operation and returns the result
	DropoutTensorRateTensorName(tensor IMPSGraphTensor, rate IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a dropout operation and returns the result
	DropoutTensorRateName(tensor IMPSGraphTensor, rate float64, name string) IMPSGraphTensor
	// Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
	EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary, executionDescriptor IMPSGraphExecutionDescriptor)
	// Encodes the graph for the given feeds to returns the target tensor values, ensuring all target operations also executed.
	EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary
	// Returns the elementwise equality check of the input tensors.
	EqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the error function to the input tensor elements.
	ErfWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates an expand-dimensions operation and returns the result tensor.
	ExpandDimsOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates an expand-dimensions operation and returns the result tensor.
	ExpandDimsOfTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates an expand-dimensions operation and returns the result tensor.
	ExpandDimsOfTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Applies the natural exponent to the input tensor elements.
	ExponentWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies an exponent with base 10 to the input tensor elements.
	ExponentBase10WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies an exponent with base 2 to the input tensor elements.
	ExponentBase2WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a fast Fourier transform operation and returns the result tensor.
	FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Creates a fast Fourier transform operation and returns the result tensor.
	FastFourierTransformWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Creates a flatten2D operation and returns the result tensor.
	Flatten2DTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a flatten2D operation and returns the result tensor.
	Flatten2DTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the floor operation to the input tensor elements.
	FloorWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the remainder of floor divison between the primary and secondary tensor.
	FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Gather operation and returns the result tensor.
	GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis uint, batchDimensions uint, name string) IMPSGraphTensor
	// Creates a GatherAlongAxis operation and returns the result tensor.
	GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a GatherAlongAxis operation and returns the result tensor.
	GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a GatherND operation and returns the result tensor.
	GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, name string) IMPSGraphTensor
	// Calculates a partial derivative of primaryTensor with respect to the tensors.
	GradientForPrimaryTensorWithTensorsName(primaryTensor IMPSGraphTensor, tensors []MPSGraphTensor, name string) foundation.INSDictionary
	// Checks in an elementwise manner if the first input tensor is greater than the second.
	GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks in an elementwise manner if the first input tensor is greater than or equal to the second.
	GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Copies the input tensor values into the output, behaving as an identity operation.
	IdentityWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates an imToCol operation and returns the result tensor.
	ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name string) IMPSGraphTensor
	// Returns the imaginary part of a tensor.
	ImaginaryPartOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the inverse of an input tensor.
	InverseOfTensorName(inputTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks if the input tensor elements are finite or not.
	IsFiniteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks if the input tensor elements are infinite or not.
	IsInfiniteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks if the input tensor elements are [NaN] or not.
	IsNaNWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
	LeakyReLUWithTensorAlphaName(tensor IMPSGraphTensor, alpha float64, name string) IMPSGraphTensor
	// Computes the leaky rectified linear unit (ReLU) activation function on the input tensor.
	LeakyReLUWithTensorAlphaTensorName(tensor IMPSGraphTensor, alphaTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the gradient of the leaky rectified linear unit (ReLU) activation.
	LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, alphaTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks in an elementwise manner if the first input tensor is less than the second.
	LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Checks in an elementwise manner if the first input tensor is less than or equal to the second.
	LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the natural logarithm to the input tensor elements.
	LogarithmWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the logarithm with base 10 to the input tensor elements.
	LogarithmBase10WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the logarithm with base 2 to the input tensor elements.
	LogarithmBase2WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical AND of the input tensors.
	LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical NAND of the input tensors.
	LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical NOR of the input tensors.
	LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical OR of the input tensors.
	LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical XNOR of the input tensors.
	LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise logical XOR of the input tensors.
	LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the matrix multiplication of 2 input tensors with support for broadcasting.
	MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a 2D max-pooling operation and returns the result tensor.
	MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 2D max-pooling operation and returns the result tensor and the corresponding indices tensor.
	MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) []MPSGraphTensor
	// Creates a 4D max-pooling operation and returns the result tensor.
	MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates a max-pooling gradient operation and returns the result tensor.
	MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor
	// Creates a 4D max-pooling operation and returns the result tensor and the corresponding indices tensor.
	MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) []MPSGraphTensor
	// Returns the elementwise maximum of the input tensors.
	MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise maximum of the input tensors, while propagating [NaN] values.
	MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the mean of the first input along the specified axes.
	MeanOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Returns the elementwise minimum of the input tensors.
	MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise minimum of the input tensors, while propagating [NaN] values.
	MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the remainder obtained by dividing the first input tensor by the second.
	ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Multiplies two input tensors.
	MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies negative to the input tensor elements.
	NegativeWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a nonMaximumumSuppression operation and returns the result tensor.
	NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, classIndicesTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode MPSGraphNonMaximumSuppressionCoordinateMode, name string) IMPSGraphTensor
	// Creates a nonMaximumumSuppression operation and returns the result tensor.
	NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode MPSGraphNonMaximumSuppressionCoordinateMode, name string) IMPSGraphTensor
	// Computes the indices of the non-zero elements of the input tensor.
	NonZeroIndicesOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a normalization beta-gradient operation and returns the result tensor.
	NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a normalization gamma-gradient operation and returns the result tensor.
	NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, axes []foundation.NSNumber, epsilon float32, name string) IMPSGraphTensor
	// Creates a normalization input gradient operation and returns the result tensor.
	NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, gamma IMPSGraphTensor, gammaGradient IMPSGraphTensor, betaGradient IMPSGraphTensor, axes []foundation.NSNumber, epsilon float32, name string) IMPSGraphTensor
	// Creates a batch normalization operation and returns the result tensor.
	NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor IMPSGraphTensor, mean IMPSGraphTensor, variance IMPSGraphTensor, gamma IMPSGraphTensor, beta IMPSGraphTensor, epsilon float32, name string) IMPSGraphTensor
	// Applies the logical NOT operation to the input tensor elements.
	NotWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the elementwise inequality check of the input tensors.
	NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthAxisDataTypeName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType uint32, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType uint32, onValue float64, offValue float64, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthAxisName(indicesTensor IMPSGraphTensor, depth uint, axis uint, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthDataTypeName(indicesTensor IMPSGraphTensor, depth uint, dataType uint32, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, dataType uint32, onValue float64, offValue float64, name string) IMPSGraphTensor
	// Creates a oneHot operation and returns the result tensor.
	OneHotWithIndicesTensorDepthName(indicesTensor IMPSGraphTensor, depth uint, name string) IMPSGraphTensor
	// Creates a padding gradient operation and returns the result tensor.
	PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, paddingMode MPSGraphPaddingMode, leftPadding foundation.NSArray, rightPadding foundation.NSArray, name string) IMPSGraphTensor
	// Creates a padding operation and returns the result tensor.
	PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor IMPSGraphTensor, paddingMode MPSGraphPaddingMode, leftPadding foundation.NSArray, rightPadding foundation.NSArray, constantValue float64, name string) IMPSGraphTensor
	// Creates a placeholder operation and returns the result tensor.
	PlaceholderWithShapeDataTypeName(shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor
	// Creates a placeholder operation and returns the result tensor with the dataType of the placeholder tensor set to 32 bit float.
	PlaceholderWithShapeName(shape foundation.NSArray, name string) IMPSGraphTensor
	// Returns the elementwise result of raising the first tensor to the power of the second tensor.
	PowerWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Quantize operation and returns the result tensor.
	QuantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType uint32, name string) IMPSGraphTensor
	// Creates a Quantize operation and returns the result tensor.
	QuantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType uint32, axis int, name string) IMPSGraphTensor
	// Creates a Quantize operation and returns the result tensor.
	QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, axis int, name string) IMPSGraphTensor
	// Creates a tensor representing state using the Philox algorithm with given counter and key values.
	RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name string) IMPSGraphTensor
	// Creates a tensor representing state using the Philox algorithm with given counter and key values.
	RandomPhiloxStateTensorWithSeedName(seed uint, name string) IMPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor and returns random values.
	RandomTensorWithShapeDescriptorName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, name string) IMPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor and returns random values.
	RandomTensorWithShapeDescriptorSeedName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, seed uint, name string) IMPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
	RandomTensorWithShapeDescriptorStateTensorName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor and returns random values.
	RandomTensorWithShapeTensorDescriptorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, name string) IMPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor and returns random values.
	RandomTensorWithShapeTensorDescriptorSeedName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, seed uint, name string) IMPSGraphTensor
	// Creates a Random op of type matching distribution in descriptor, and returns random values and updated state.
	RandomTensorWithShapeTensorDescriptorStateTensorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values
	RandomUniformTensorWithShapeName(shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values
	RandomUniformTensorWithShapeSeedName(shape foundation.NSArray, seed uint, name string) IMPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values and updated state
	RandomUniformTensorWithShapeStateTensorName(shape foundation.NSArray, state IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values
	RandomUniformTensorWithShapeTensorName(shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values
	RandomUniformTensorWithShapeTensorSeedName(shapeTensor IMPSGraphTensor, seed uint, name string) IMPSGraphTensor
	// Creates a RandomUniform operation and returns random uniform values and updated state
	RandomUniformTensorWithShapeTensorStateTensorName(shapeTensor IMPSGraphTensor, state IMPSGraphTensor, name string) []MPSGraphTensor
	// Computes the ReLU (rectified linear activation unit) function with the input tensor.
	ReLUWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the gradient of the ReLU  (rectified linear activation unit) function using the incoming gradient.
	ReLUGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a read op which reads at this point of execution of the graph and returns the result tensor.
	ReadVariableName(variable IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the real part of a tensor.
	RealPartOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
	RealToHermiteanFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
	RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor
	// Applies the reciprocal operation to the input tensor elements.
	ReciprocalWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the reciprocal square root operation to the input tensor elements.
	ReciprocalSquareRootWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a reduction and operation and returns the result tensor.
	ReductionAndWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction and operation and returns the result tensor.
	ReductionAndWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction argMax operation and returns the result tensor.
	ReductionArgMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction argMin operation and returns the result tensor.
	ReductionArgMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction max operation and returns the result tensor.
	ReductionMaximumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction max operation and returns the result tensor.
	ReductionMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction max propagate NaN operation and returns the result tensor.
	ReductionMaximumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction max propagate NaN operation and returns the result tensor.
	ReductionMaximumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction min operation and returns the result tensor.
	ReductionMinimumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction minimum operation and returns the result tensor.
	ReductionMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction min propagate NaN operation and returns the result tensor.
	ReductionMinimumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction min propagate NaN operation and returns the result tensor.
	ReductionMinimumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction or operation and returns the result tensor.
	ReductionOrWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction or operation and returns the result tensor.
	ReductionOrWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction product operation and returns the result tensor.
	ReductionProductWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction product operation and returns the result tensor.
	ReductionProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reduction sum operation and returns the result tensor.
	ReductionSumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reduction sum operation and returns the result tensor.
	ReductionSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a reinterpret cast operation and returns the result tensor.
	ReinterpretCastTensorToTypeName(tensor IMPSGraphTensor, type_ uint32, name string) IMPSGraphTensor
	// Creates a reshape operation and returns the result tensor.
	ReshapeTensorWithShapeName(tensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a reshape operation and returns the result tensor.
	ReshapeTensorWithShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeTensorSizeModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size foundation.NSArray, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeTensorSizeTensorModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, name string) IMPSGraphTensor
	// Resamples input images to given size using the provided scale and offset. Destination indices are computed using
	ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode MPSGraphResizeMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeTensorSizeTensorScaleTensorOffsetTensorModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode MPSGraphResizeMode, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode MPSGraphResizeMode, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode MPSGraphResizeMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Resamples input images to given size using bilinear sampling.
	ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, name string) IMPSGraphTensor
	// Resamples input images to given size using the provided scale and offset and bilinear sampling See above discussion for more details.
	ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Resamples input images to given size using nearest neighbor sampling.
	ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, name string) IMPSGraphTensor
	// Resamples input images to given size using the provided scale and offset and nearest neighbor sampling See above discussion for more details.
	ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize operation and returns the result tensor.
	ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, name string) IMPSGraphTensor
	// Creates a Resize gradient operation and returns the result tensor.
	ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor
	// Creates a reverse operation and returns the result tensor.
	ReverseTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a reverse operation and returns the result tensor.
	ReverseTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a reverse operation and returns the result tensor.
	ReverseTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Rounds the input tensor elements by rounding to nearest even.
	RintWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Rounds the input tensor elements.
	RoundWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
	RunWithFeedsTargetTensorsTargetOperations(feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation) MPSGraphTensorDataDictionary
	// Runs the graph for the given feeds and returns the target tensor values in the results dictionary provided by the user.
	RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary)
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
	RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation) MPSGraphTensorDataDictionary
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
	RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary
	// Encodes the graph for the given feeds to returns the target tensor values in the results dictionary provided by the user.
	RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary, executionDescriptor IMPSGraphExecutionDescriptor)
	// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
	RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary
	// Samples a tensor using the coordinates provided, using nearest neighbor sampling with specified rounding mode.
	SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode MPSGraphPaddingMode, nearestRoundingMode MPSGraphResizeNearestRoundingMode, constantValue float64, name string) IMPSGraphTensor
	// Samples a tensor using the coordinates provided.
	SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode MPSGraphPaddingMode, samplingMode MPSGraphResizeMode, constantValue float64, name string) IMPSGraphTensor
	// Creates a scaled dot product attention (SDPA) operation and returns the result tensor.
	ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, maskTensor IMPSGraphTensor, scale float32, name string) IMPSGraphTensor
	// Creates a scaled dot product attention (SDPA) operation (without a mask) and returns the result tensor.
	ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, scale float32, name string) IMPSGraphTensor
	// Creates a Scatter operation and returns the result tensor.
	ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, axis int, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterAlongAxis operation and returns the result tensor.
	ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterAlongAxis operation and returns the result tensor.
	ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterAlongAxis operation and returns the result tensor.
	ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor IMPSGraphTensor, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterAlongAxis operation and returns the result tensor.
	ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterND operation and returns the result tensor.
	ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, batchDimensions uint, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a ScatterND operation and returns the result tensor.
	ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, batchDimensions uint, name string) IMPSGraphTensor
	// Creates a ScatterND operation and returns the result tensor.
	ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Creates a Scatter operation and returns the result tensor.
	ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis int, mode MPSGraphScatterMode, name string) IMPSGraphTensor
	// Selects values from either the true or false predicate tensor, depending on the values in the first input.
	SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor IMPSGraphTensor, truePredicateTensor IMPSGraphTensor, falseSelectTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a shape-of operation and returns the result tensor.
	ShapeOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the sigmoid operation on an input tensor.
	SigmoidWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Computes the gradient of the sigmoid function using the incoming gradient tensor.
	SigmoidGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the sign of the input tensor elements.
	SignWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the sign bit of the input tensor elements.
	SignbitWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the sine operation to the input tensor elements.
	SinWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
	SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
	SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
	SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Creates a single-gate RNN gradient operation and returns the gradient tensor values.
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor
	// Applies the hyperbolic sine operation to the input tensor elements.
	SinhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a strided-slice gradient operation and returns the result tensor.
	SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a slice gradient operation and returns the result tensor.
	SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a strided-slice gradient operation and returns the result tensor.
	SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a strided-slice gradient operation and returns the result tensor.
	SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a slice operation and returns the result tensor.
	SliceTensorDimensionStartLengthName(tensor IMPSGraphTensor, dimensionIndex uint, start int, length int, name string) IMPSGraphTensor
	// Creates a strided-slice operation and returns the result tensor.
	SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a slice operation and returns the result tensor.
	SliceTensorStartTensorSizeTensorSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a strided-slice operation and returns the result tensor.
	SliceTensorStartsEndsStridesName(tensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a strided-slice operation and returns the result tensor.
	SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a strided-slice update operation with zero masks and returns the result tensor.
	SliceUpdateDataTensorUpdateTensorStartsEndsStridesName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a strided-slice update operation and returns the result tensor.
	SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Creates a strided-slice update operation with zero masks and returns the result tensor.
	SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a strided-slice update operation and returns the result tensor.
	SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor
	// Computes the softmax function on the input tensor along the specified axis.
	SoftMaxWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a softmax cross-entropy loss operation and returns the result tensor.
	SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType MPSGraphLossReductionType, name string) IMPSGraphTensor
	// Creates the gradient of a softmax cross-entropy loss operation and returns the result tensor.
	SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType MPSGraphLossReductionType, name string) IMPSGraphTensor
	// Computes the gradient of the softmax function along the specified axis using the incoming gradient tensor.
	SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Sorts the elements of the input tensor along the specified axis.
	SortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name string) IMPSGraphTensor
	// Sorts the elements of the input tensor along the specified axis.
	SortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Sorts the elements of the input tensor along the specified axis.
	SortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name string) IMPSGraphTensor
	// Sorts the elements of the input tensor along the specified axis.
	SortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a space-to-depth2D operation and returns the result tensor.
	SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a space-to-depth2D operation and returns the result tensor.
	SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a space-to-batch operation and returns the result tensor.
	SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.NSNumber, batchAxis int, blockDimensions []foundation.NSNumber, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a space-to-batch operation and returns the result tensor.
	SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name string) IMPSGraphTensor
	// Creates a sparse tensor representation.
	SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor IMPSGraphCreateSparseOpDescriptor, inputTensorArray []MPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor
	// Creates a sparse tensor representation.
	SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType MPSGraphSparseStorageType, inputTensorArray []MPSGraphTensor, shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor
	// Creates a split operation and returns the result tensor.
	SplitTensorNumSplitsAxisName(tensor IMPSGraphTensor, numSplits uint, axis int, name string) []MPSGraphTensor
	// Creates a split operation and returns the result tensor.
	SplitTensorSplitSizesAxisName(tensor IMPSGraphTensor, splitSizes []foundation.NSNumber, axis int, name string) []MPSGraphTensor
	// Creates a split operation and returns the result tensor.
	SplitTensorSplitSizesTensorAxisName(tensor IMPSGraphTensor, splitSizesTensor IMPSGraphTensor, axis int, name string) []MPSGraphTensor
	// Applies the square operation to the input tensor elements.
	SquareWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the square root operation to the input tensor elements.
	SquareRootWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a squeeze operation and returns the result tensor.
	SqueezeTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a squeeze operation and returns the result tensor.
	SqueezeTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a squeeze operation and returns the result tensor.
	SqueezeTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a squeeze operation and returns the result tensor.
	SqueezeTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a stack operation and returns the result tensor.
	StackTensorsAxisName(inputTensors []MPSGraphTensor, axis int, name string) IMPSGraphTensor
	// Creates a stencil operation and returns the result tensor.
	StencilWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphStencilOpDescriptor, name string) IMPSGraphTensor
	// The Stochastic gradient descent performs a gradient descent.
	StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Subtracts the second input tensor from the first.
	SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the tangent operation to the input tensor elements.
	TanWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Applies the hyperbolic tangent operation to the input tensor elements.
	TanhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a tile gradient operation and returns the result tensor.
	TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, multiplier foundation.NSArray, name string) IMPSGraphTensor
	// Creates a tile operation and returns the result tensor.
	TileTensorWithMultiplierName(tensor IMPSGraphTensor, multiplier foundation.NSArray, name string) IMPSGraphTensor
	// Creates a TopK operation and returns the value and indices tensors.
	TopKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name string) []MPSGraphTensor
	// Creates a TopK operation and returns the result tensor.
	TopKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a TopK operation and returns the value and indices tensors
	TopKWithSourceTensorKName(source IMPSGraphTensor, k uint, name string) []MPSGraphTensor
	// Creates a TopK operation and returns the result tensor.
	TopKWithSourceTensorKTensorName(source IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor
	// Creates a TopKGradient operation and returns the result tensor.
	TopKWithGradientTensorSourceKName(gradient IMPSGraphTensor, source IMPSGraphTensor, k uint, name string) IMPSGraphTensor
	// Creates a TopKGradient operation and returns the result tensor.
	TopKWithGradientTensorSourceKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a TopKGradient operation and returns the result tensor.
	TopKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name string) IMPSGraphTensor
	// Creates a TopKGradient operation and returns the result tensor.
	TopKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a permutation operation and returns the result tensor.
	TransposeTensorPermutationName(tensor IMPSGraphTensor, permutation []foundation.NSNumber, name string) IMPSGraphTensor
	// Creates a transpose operation and returns the result tensor.
	TransposeTensorDimensionWithDimensionName(tensor IMPSGraphTensor, dimensionIndex uint, dimensionIndex2 uint, name string) IMPSGraphTensor
	// Applies the truncate operation to the input tensor elements.
	TruncateWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Creates a variable operation and returns the result tensor.
	VariableWithDataShapeDataTypeName(data foundation.NSData, shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor
	// Creates a variable from an input tensor.
	VariableFromTensorWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor
	// Returns the variance of the first input along the specified axes.
	VarianceOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor
	// Returns the variance of the first input along the specified axes when the mean has been precomputed.
	VarianceOfTensorMeanTensorAxesName(tensor IMPSGraphTensor, meanTensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor

	// Creates a scaled dot product attention (SDPA) operation using a descriptor and returns the result tensor.
	ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorDescriptorName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, descriptor IMPSGraphSDPADescriptor, name string) IMPSGraphTensor
}

// Init initializes the instance.
func (g MPSGraph) Init() MPSGraph {
	rv := objc.Send[MPSGraph](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraph) Autorelease() MPSGraph {
	rv := objc.Send[MPSGraph](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraph creates a new MPSGraph instance.
func NewMPSGraph() MPSGraph {
	class := getMPSGraphClass()
	rv := objc.Send[MPSGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a GRU operation and returns the value and optionally the training
// state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 depending on value of
// `descriptor.Training()`. The layout of the state output is [T,N,H] or
// [T,N,2H] for bidirectional, and the layout of the `trainingState` output is
// [T,N,3H] or [T,N,6H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// If `resetAfter = YES` then `c[t]` is replaced by
//
// If `flipZ = YES` then `h[t]` is replaced by
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `z[t]` is the second output (optional) and `h[-1]` is `initState`.
// `b2` is an optional `resetBias` vector, only used when `resetAfter = YES`.
// See [MPSGraphGRUDescriptor] for different `activation` options for `f()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:descriptor:name:)
func (g MPSGraph) GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:descriptor:name:"), source, recurrentWeight, inputWeight, bias, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a GRU operation and returns the value and optionally the training
// state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 depending on value of
// `descriptor.Training()`. The layout of the state output is [T,N,H] or
// [T,N,2H] for bidirectional, and the layout of the `trainingState` output is
// [T,N,3H] or [T,N,6H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// If `resetAfter = YES` then `c[t]` is replaced by
//
// If `flipZ = YES` then `h[t]` is replaced by
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `z[t]` is the second output (optional) and `h[-1]` is `initState`.
// `b2` is an optional `resetBias` vector, only used when `resetAfter = YES`.
// See [MPSGraphGRUDescriptor] for different `activation` options for `f()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:initState:descriptor:name:)
func (g MPSGraph) GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a GRU operation and returns the value and optionally the training
// state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. Useful for dropout.
//
// secondaryBias: A tensor containing the secondary bias vector `b2` - optional, if missing
// the operation assumes zeroes. Only used with `reset_after = YES`. Shape is
// [H], ie. a vector for each gate, or [2H] for bidirectional.
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 depending on value of
// `descriptor.Training()`. The layout of the state output is [T,N,H] or
// [T,N,2H] for bidirectional, and the layout of the `trainingState` output is
// [T,N,3H] or [T,N,6H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// If `resetAfter = YES` then `c[t]` is replaced by
//
// If `flipZ = YES` then `h[t]` is replaced by
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `z[t]` is the second output (optional) and `h[-1]` is `initState`.
// `b2` is an optional `resetBias` vector, only used when `resetAfter = YES`.
// See [MPSGraphGRUDescriptor] for different `activation` options for `f()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:)
func (g MPSGraph) GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, mask, secondaryBias, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a GRU gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// outputFwd: The first output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight` and for `bias`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:)
func (g MPSGraph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, inputWeight, bias, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a GRU gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// outputFwd: The first output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias` and for
// `initState`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:initState:descriptor:name:)
func (g MPSGraph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, inputWeight, bias, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a GRU gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,3H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,6H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,3H,H] and otherwise it is [3H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// outputFwd: The first output of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// with `descriptor.Training() = YES`.
//
// stateGradient: The input gradient for state coming from the future timestep - optional, if
// missing the operation assumes zeroes.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [6H,I] and otherwise it is [3H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [6H] and otherwise it is
// [3H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. Useful for dropout.
//
// secondaryBias: A tensor containing the secondary bias vector `b2` - optional, if missing
// the operation assumes zeroes. Only used with `reset_after = YES`. Shape is
// [H], ie. a vector for each gate, or [2H] for bidirectional.
//
// descriptor: A descriptor that defines the parameters for the GRU operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias`, for
// `initState` and for `secondaryBias`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:stateGradient:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:)
func (g MPSGraph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, secondaryBias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:stateGradient:inputWeight:bias:initState:mask:secondaryBias:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, stateGradient, inputWeight, bias, initState, mask, secondaryBias, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Computes the hamming distance of two input tensors with support for
// broadcasting.
//
// primaryTensor: The first input tensor.
//
// secondaryTensor: The second input tensor.
//
// resultDataType: The datatype of the return MPSGraphTensor. Must be either
// [MPSDataTypeUInt32] or [MPSDataTypeUInt16].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid tensor containing the hamming distance between the input tensors.
//
// # Discussion
//
// The hamming distance is computed between 2 sets of vectors and the last
// dimension(s) of each input tensor is considered a vector.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HammingDistance(primary:secondary:resultDataType:name:)
func (g MPSGraph) HammingDistanceWithPrimaryTensorSecondaryTensorResultDataTypeName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, resultDataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("HammingDistanceWithPrimaryTensor:secondaryTensor:resultDataType:name:"), primaryTensor, secondaryTensor, resultDataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Hermitean-to-real fast Fourier transform operation and returns
// the result tensor.
//
// tensor: A complex-valued input tensor with reduced size (see Discussion). Must have
// datatype [MPSDataTypeComplexFloat32] or [MPSDataTypeComplexFloat16].
//
// axes: An array of numbers that specifies over which axes MPSGraph performs the
// Fourier transform - all axes must be contained within last four dimensions
// of the input tensor.
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor of type [MPSDataTypeFloat32] or [MPSDataTypeFloat16]
// (full size).
//
// # Discussion
//
// This operation computes the fast Fourier transform of a complex-valued
// input tensor according to the following formulae.
//
// `in'[nu] = conjugate(in[n - nu])`, for the last dimension defined by `axes`
// when `nu` is out of range of the input dimension. `scale = 1` for
// `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode = size`, `scale =
// 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is the volume of the
// transformation defined by the dimensions included in `axes` (`V_f = prod_{i
// \in axes} shape(input)[i]`) (see [MPSGraphFFTDescriptor.ScalingMode]), `+`
// is selected in `+/-` when `inverse` is specified, otherwise `-` is used and
// the sum is done separately over each dimension in `axes` and `n` is the
// dimension length of that axis. With this API MPSGraph treats the input
// tensor to have only the unique frequencies, which means that the resulting
// tensor has size `(inSize-1)*2 + x` in the last dimension defined by `axes`,
// where `inSize = shape(input)[axis] ( = (n/2)+1 )` is the size of the input
// `tensor` in the last transformed dimension and `x = 1` when
// [MPSGraphFFTDescriptor.RoundToOddHermitean] = [YES] and `x = 0` otherwise.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HermiteanToRealFFT(_:axes:descriptor:name:)
func (g MPSGraph) HermiteanToRealFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("HermiteanToRealFFTWithTensor:axes:descriptor:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Hermitean-to-real fast Fourier transform operation and returns
// the result tensor.
//
// tensor: A complex-valued input tensor with reduced size (see Discussion). Must have
// datatype [MPSDataTypeComplexFloat32] or [MPSDataTypeComplexFloat16].
//
// axesTensor: A tensor of rank one containing the axes over which MPSGraph performs the
// transformation. See
// [MPSGraph.FastFourierTransformWithTensorAxesDescriptorName].
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor of type [MPSDataTypeFloat32] or [MPSDataTypeFloat16]
// (full size).
//
// # Discussion
//
// This operation computes the fast Fourier transform of a complex-valued
// input tensor according to the following formulae.
//
// `in'[nu] = conjugate(in[n - nu])`, for the last dimension defined by `axes`
// when `nu` is out of range of the input dimension. `scale = 1` for
// `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode = size`, `scale =
// 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is the volume of the
// transformation defined by the dimensions included in `axes` (`V_f = prod_{i
// \in axes} shape(input)[i]`) (see [MPSGraphFFTDescriptor.ScalingMode]), `+`
// is selected in `+/-` when `inverse` is specified, otherwise `-` is used and
// the sum is done separately over each dimension in `axes` and `n` is the
// dimension length of that axis. With this API MPSGraph treats the input
// tensor to have only the unique frequencies, which means that the resulting
// tensor has size `(inSize-1)*2 + x` in the last dimension defined by `axes`,
// where `inSize = shape(input)[axis] ( = (n/2)+1 )` is the size of the input
// `tensor` in the last transformed dimension and `x = 1` when
// [MPSGraphFFTDescriptor.RoundToOddHermitean] = [YES] and `x = 0` otherwise.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HermiteanToRealFFT(_:axesTensor:descriptor:name:)
func (g MPSGraph) HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("HermiteanToRealFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 4D L2-norm pooling operation and returns the result tensor.
//
// source: A source tensor.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/L2NormPooling4D(_:descriptor:name:)
func (g MPSGraph) L2NormPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("L2NormPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a L2-Norm pooling gradient operation and returns the result tensor.
//
// gradient: An input gradient tensor.
//
// source: The input tensor for the forward pass.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/L2NormPooling4DGradient(_:source:descriptor:name:)
func (g MPSGraph) L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("L2NormPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an LSTM operation and returns the value tensor and optionally the
// cell state tensor and the training state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 or 3, depending on values of
// `descriptor.ProduceCell()` and `descriptor.Training()`. The layout of the
// both state and cell outputs are [T,N,H] or [T,N,2H] for bidirectional, and
// the layout of the trainingState output is [T,N,4H] or [T,N,8H] for
// bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `c` and optionally `z`
// that are defined recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `c[t]` is the second output (optional), `z[t]` is either the second
// or third output (optional), `h[-1]` is `initCell`. and `h[-1]` is
// `initState`. `p` is an optional peephole vector. See
// [MPSGraphLSTMDescriptor] for different `activation` options for `f()` and
// `g()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:initState:initCell:descriptor:name:)
func (g MPSGraph) LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:initState:initCell:descriptor:name:"), source, recurrentWeight, initState, initCell, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM operation and returns the value tensor and optionally the
// cell state tensor and the training state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [8H,I] and otherwise it is [4H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [8H] and otherwise it is
// [4H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 or 3, depending on values of
// `descriptor.ProduceCell()` and `descriptor.Training()`. The layout of the
// both state and cell outputs are [T,N,H] or [T,N,2H] for bidirectional, and
// the layout of the trainingState output is [T,N,4H] or [T,N,8H] for
// bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `c` and optionally `z`
// that are defined recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `c[t]` is the second output (optional), `z[t]` is either the second
// or third output (optional), `h[-1]` is `initCell`. and `h[-1]` is
// `initState`. `p` is an optional peephole vector. See
// [MPSGraphLSTMDescriptor] for different `activation` options for `f()` and
// `g()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:inputWeight:bias:initState:initCell:descriptor:name:)
func (g MPSGraph) LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:inputWeight:bias:initState:initCell:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, initCell, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM operation and returns the value tensor and optionally the
// cell state tensor and the training state tensor.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [8H,I] and otherwise it is [4H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [8H] and otherwise it is
// [4H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. Useful for dropout.
//
// peephole: A tensor containing the peephole vector `v` - optional, if missing the
// operation assumes zeroes. Shape is [4H], ie. a vector for each gate, or
// [2,4H] for bidirectional.
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of size 1 or 2 or 3, depending on values of
// `descriptor.ProduceCell()` and `descriptor.Training()`. The layout of the
// both state and cell outputs are [T,N,H] or [T,N,2H] for bidirectional, and
// the layout of the trainingState output is [T,N,4H] or [T,N,8H] for
// bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `c` and optionally `z`
// that are defined recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is optional
// `bias`, `m` is optional `mask`, `x[t]` is `source` `h[t]` is the first
// output, `c[t]` is the second output (optional), `z[t]` is either the second
// or third output (optional), `h[-1]` is `initCell`. and `h[-1]` is
// `initState`. `p` is an optional peephole vector. See
// [MPSGraphLSTMDescriptor] for different `activation` options for `f()` and
// `g()`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:)
func (g MPSGraph) LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, initCell, mask, peephole, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The third output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES`.
//
// cellOutputFwd: The second output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES` or `descriptor.ProduceCell() = YES`.
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias`, for
// `initState` and for `initCell`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:descriptor:name:)
func (g MPSGraph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The third output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES`.
//
// cellOutputFwd: The second output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES` or `descriptor.ProduceCell() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [8H,I] and otherwise it is [4H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [8H] and otherwise it is
// [4H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias`, for
// `initState` and for `initCell`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:descriptor:name:)
func (g MPSGraph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, inputWeight, bias, initState, initCell, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The third output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES`.
//
// cellOutputFwd: The second output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES` or `descriptor.ProduceCell() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [8H,I] and otherwise it is [4H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [8H] and otherwise it is
// [4H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. Useful for dropout.
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias`, for
// `peephole`, for `initState` and for `initCell`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:mask:descriptor:name:)
func (g MPSGraph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:inputWeight:bias:initState:initCell:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, inputWeight, bias, initState, initCell, mask, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates an LSTM gradient operation and returns the gradient tensor values.
//
// source: A tensor containing the source data `x[t]` with the data layout [T,N,I]. In
// case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,4H] and for `inputWeight = nil` and `bidirectional = YES` the layout
// is [T,N,8H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,4H,H] and otherwise it is [4H,H].
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The third output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES`.
//
// cellOutputFwd: The second output of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// with `descriptor.Training() = YES` or `descriptor.ProduceCell() = YES`.
//
// stateGradient: The input gradient for state coming from the future timestep - optional, if
// missing the operation assumes zeroes.
//
// cellGradient: Input gradient for cell coming from the future timestep - optional, if
// missing the operation assumes zeroes.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [8H,I] and otherwise it is [4H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [8H] and otherwise it is
// [4H].
//
// initState: The initial internal state of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// initCell: The initial internal cell of the LSTM `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. Useful for dropout.
//
// peephole: A tensor containing the peephole vector `v` - optional, if missing the
// operation assumes zeroes. Shape is [4H], ie. a vector for each gate, or
// [2,4H] for bidirectional.
//
// descriptor: A descriptor that defines the parameters for the LSTM operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is nil, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias`, for
// `peephole`, for `initState` and for `initCell`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTMGradients(_:recurrentWeight:sourceGradient:zState:cellOutputFwd:stateGradient:cellGradient:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:)
func (g MPSGraph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, cellOutputFwd IMPSGraphTensor, stateGradient IMPSGraphTensor, cellGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, mask IMPSGraphTensor, peephole IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:stateGradient:cellGradient:inputWeight:bias:initState:initCell:mask:peephole:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, cellOutputFwd, stateGradient, cellGradient, inputWeight, bias, initState, initCell, mask, peephole, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Returns the absolute values of the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/absolute(with:name:)
func (g MPSGraph) AbsoluteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("absoluteWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the absolute square of the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation..
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/absoluteSquare(tensor:name:)
func (g MPSGraph) AbsoluteSquareWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("absoluteSquareWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the inverse cosine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/acos(with:name:)
func (g MPSGraph) AcosWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("acosWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the inverse hyperbolic cosine operation to the input tensor
// elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/acosh(with:name:)
func (g MPSGraph) AcoshWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("acoshWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates operations to apply Adam optimization.
//
// beta1Tensor: beta1Tensor
//
// beta2Tensor: beta2Tensor
//
// epsilonTensor: Epsilon tensor
//
// valuesTensor: Values to update with optimization
//
// momentumTensor: Momentum tensor
//
// velocityTensor: Velocity tensor
//
// maximumVelocityTensor: Optional maximum velocity tensor
//
// gradientTensor: Partial gradient of the trainable parameters with respect to loss
//
// name: Name for the operation
//
// # Return Value
//
// If maximumVelocity is nil array of 3 tensors (update, newMomentum,
// newVelocity) else array of 4 tensors (update, newMomentum, newVelocity,
// newMaximumVelocity)
//
// # Discussion
//
// # The adam update ops are added
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/adam(currentLearningRate:beta1:beta2:epsilon:values:momentum:velocity:maximumVelocity:gradient:name:)
func (g MPSGraph) AdamWithCurrentLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(currentLearningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("adamWithCurrentLearningRateTensor:beta1Tensor:beta2Tensor:epsilonTensor:valuesTensor:momentumTensor:velocityTensor:maximumVelocityTensor:gradientTensor:name:"), currentLearningRateTensor, beta1Tensor, beta2Tensor, epsilonTensor, valuesTensor, momentumTensor, velocityTensor, maximumVelocityTensor, gradientTensor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates operations to apply Adam optimization.
//
// learningRateTensor: Scalar tensor which indicates the learning rate to use with the optimizer
//
// beta1Tensor: beta1Tensor
//
// beta2Tensor: beta2Tensor
//
// beta1PowerTensor: `beta1^t` beta1 power tensor
//
// beta2PowerTensor: `beta2^t` beta2 power tensor
//
// valuesTensor: Values to update with optimization
//
// momentumTensor: Momentum tensor
//
// velocityTensor: Velocity tensor
//
// maximumVelocityTensor: Optional maximum velocity tensor
//
// gradientTensor: Partial gradient of the trainable parameters with respect to loss
//
// name: Name for the operation
//
// # Return Value
//
// If maximumVelocity is nil array of 3 tensors (update, newMomentum,
// newVelocity) else array of 4 tensors (update, newMomentum, newVelocity,
// newMaximumVelocity)
//
// # Discussion
//
// The adam update ops are added current learning rate:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/adam(learningRate:beta1:beta2:epsilon:beta1Power:beta2Power:values:momentum:velocity:maximumVelocity:gradient:name:)
func (g MPSGraph) AdamWithLearningRateTensorBeta1TensorBeta2TensorEpsilonTensorBeta1PowerTensorBeta2PowerTensorValuesTensorMomentumTensorVelocityTensorMaximumVelocityTensorGradientTensorName(learningRateTensor IMPSGraphTensor, beta1Tensor IMPSGraphTensor, beta2Tensor IMPSGraphTensor, epsilonTensor IMPSGraphTensor, beta1PowerTensor IMPSGraphTensor, beta2PowerTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, momentumTensor IMPSGraphTensor, velocityTensor IMPSGraphTensor, maximumVelocityTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("adamWithLearningRateTensor:beta1Tensor:beta2Tensor:epsilonTensor:beta1PowerTensor:beta2PowerTensor:valuesTensor:momentumTensor:velocityTensor:maximumVelocityTensor:gradientTensor:name:"), learningRateTensor, beta1Tensor, beta2Tensor, epsilonTensor, beta1PowerTensor, beta2PowerTensor, valuesTensor, momentumTensor, velocityTensor, maximumVelocityTensor, gradientTensor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Adds two input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates an add operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/addition(_:_:name:)
func (g MPSGraph) AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("additionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// The Stochastic gradient descent performs a gradient descent `variable =
// variable - (learningRate * g)` where, `g` is gradient of error wrt variable
// this op directly writes to the variable
//
// learningRateTensor: Scalar tensor which indicates the learning rate to use with the optimizer
//
// variable: Variable operation with trainable parameters
//
// gradientTensor: Partial gradient of the trainable parameters with respect to loss
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/applyStochasticGradientDescent(learningRate:variable:gradient:name:)
func (g MPSGraph) ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor IMPSGraphTensor, variable IMPSGraphVariableOp, gradientTensor IMPSGraphTensor, name string) IMPSGraphOperation {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("applyStochasticGradientDescentWithLearningRateTensor:variable:gradientTensor:name:"), learningRateTensor, variable, gradientTensor, objc.String(name))
	return MPSGraphOperationFromID(rv)
}

// Computes the indices that sort the elements of the input tensor along the
// specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension over which you sort the tensor
//
// descending: If true, reverse the sort direction
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object with 32-bit integer data type
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axis:descending:name:)
func (g MPSGraph) ArgSortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("argSortWithTensor:axis:descending:name:"), tensor, axis, descending, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the indices that sort the elements of the input tensor along the
// specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension over which you sort the tensor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object with 32-bit integer data type
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axis:name:)
func (g MPSGraph) ArgSortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("argSortWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the indices that sort the elements of the input tensor along the
// specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension over which you sort the tensor
//
// descending: If true, reverse the sort direction
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object with 32-bit integer data type
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axisTensor:descending:name:)
func (g MPSGraph) ArgSortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("argSortWithTensor:axisTensor:descending:name:"), tensor, axisTensor, descending, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the indices that sort the elements of the input tensor along the
// specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension over which you sort the tensor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object with 32-bit integer data type
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/argSort(_:axisTensor:name:)
func (g MPSGraph) ArgSortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("argSortWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the inverse sine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/asin(with:name:)
func (g MPSGraph) AsinWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("asinWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the inverse hyperbolic sine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/asinh(with:name:)
func (g MPSGraph) AsinhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("asinhWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an assign operation which writes at this point of execution of the
// graph.
//
// variable: The variable resource tensor to assign to.
//
// tensor: The tensor to assign to the variable.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/assign(_:tensor:name:)
func (g MPSGraph) AssignVariableWithValueOfTensorName(variable IMPSGraphTensor, tensor IMPSGraphTensor, name string) IMPSGraphOperation {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("assignVariable:withValueOfTensor:name:"), variable, tensor, objc.String(name))
	return MPSGraphOperationFromID(rv)
}

// Applies the inverse tangent operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/atan(with:name:)
func (g MPSGraph) AtanWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("atanWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise two-argument arctangent of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a `atan2` operation and returns the result tensor.
// It supports broadcasting as well. Graph computes arc tangent of
// primaryTensor over secondaryTensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/atan2(withPrimaryTensor:secondaryTensor:name:)
func (g MPSGraph) Atan2WithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("atan2WithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the inverse hyperbolic tangent operation to the input tensor
// elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/atanh(with:name:)
func (g MPSGraph) AtanhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("atanhWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D average-pooling operation and returns the result tensor.
//
// source: A 2D Image source as tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling2D(withSourceTensor:descriptor:name:)
func (g MPSGraph) AvgPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("avgPooling2DWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D average pooling gradient operation and returns the result
// tensor.
//
// gradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// source: The input tensor for the forward pass.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling2DGradient(withGradientTensor:sourceTensor:descriptor:name:)
func (g MPSGraph) AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("avgPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 4D average pooling operation and returns the result tensor.
//
// source: A source tensor.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling4D(_:descriptor:name:)
func (g MPSGraph) AvgPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("avgPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an average pooling gradient operation and returns the result
// tensor.
//
// gradient: An input gradient tensor.
//
// source: The input tensor for the forward pass.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/avgPooling4DGradient(_:source:descriptor:name:)
func (g MPSGraph) AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("avgPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the band part of an input tensor.
//
// inputTensor: Input tensor
//
// numLower: The number of diagonals in the lower triangle to keep. If -1, the framework
// returns all sub diagnols.
//
// numUpper: The number of diagonals in the upper triangle to keep. If -1, the framework
// returns all super diagnols.
//
// name: Name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation copies a diagonal band of values from input tensor to a
// result tensor of the same size. A coordinate `[..., i, j]` is in the band
// if
//
// The values outside of the band are set to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bandPart(_:numLower:numUpper:name:)
func (g MPSGraph) BandPartWithTensorNumLowerNumUpperName(inputTensor IMPSGraphTensor, numLower int, numUpper int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bandPartWithTensor:numLower:numUpper:name:"), inputTensor, numLower, numUpper, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates the band part operation and returns the result.
//
// inputTensor: The source tensor to copy.
//
// numLowerTensor: Scalar Int32 tensor. The number of diagonals in the lower triangle to keep.
// If -1, keep all.
//
// numUpperTensor: Scalar Int32 tensor. The number of diagonals in the upper triangle to keep.
// If -1, keep all.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// See above discussion of bandPartWithTensor: numLower: numUpper: name:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bandPart(_:numLowerTensor:numUpperTensor:name:)
func (g MPSGraph) BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor IMPSGraphTensor, numLowerTensor IMPSGraphTensor, numUpperTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bandPartWithTensor:numLowerTensor:numUpperTensor:name:"), inputTensor, numLowerTensor, numUpperTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a batch-to-space operation and returns the result tensor.
//
// tensor: The input tensor.
//
// spatialAxes: The axes that define the dimensions containing the spatial blocks.
//
// batchAxis: The axis that defines the destination dimension, where to copy the blocks.
//
// blockDimensions: An array of numbers that defines the size of the rectangular spatial
// sub-block.
//
// usePixelShuffleOrder: A parameter that controls layout of the sub-blocks within the batch
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the input tensor, where values from the
// `batchAxis` dimension are moved in spatial blocks of size `blockDimensions`
// to the `spatialAxes` dimensions (for `usePixelShuffleOrder=YES` 1,2 or 3
// axes supported, otherwise limited only by [MPSNDArray] rank limitations).
// Use the `usePixelShuffleOrder` parameter to control how the data within
// spatial blocks is ordered in the `batchAxis` dimension: with
// `usePixelShuffleOrder = YES` MPSGraph stores the values of the spatial
// block contiguosly within the `batchAxis` dimension whereas without it they
// are stored interleaved with existing values in the `batchAxis` dimension.
// Note: This operation is the inverse of
// [MPSGraph.SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName].
// Note: This operation is a generalization of
// [MPSGraph.DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/batchToSpace(_:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:)
func (g MPSGraph) BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.NSNumber, batchAxis int, blockDimensions []foundation.NSNumber, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("batchToSpaceTensor:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:"), tensor, objectivec.IObjectSliceToNSArray(spatialAxes), batchAxis, objectivec.IObjectSliceToNSArray(blockDimensions), usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a batch-to-space operation and returns the result tensor.
//
// tensor: The input tensor.
//
// spatialAxesTensor: A tensor that contains the axes that define the dimensions containing the
// spatial blocks.
//
// batchAxisTensor: A tensor that contains the axis that defines the destination dimension,
// where to copy the blocks.
//
// blockDimensionsTensor: A tensor that defines the size of the rectangular spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls layout of the sub-blocks within the batch
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the input tensor, where values from the
// `batchAxisTensor` dimension are moved in spatial blocks of size
// `blockDimensionsTensor` to the `spatialAxesTensor` dimensions (for
// `usePixelShuffleOrder=YES` 1,2 or 3 axes supported, otherwise limited only
// by [MPSNDArray] rank limitations). Use the `usePixelShuffleOrder` parameter
// to control how the data within spatial blocks is ordered in the
// `batchAxisTensor` dimension: with `usePixelShuffleOrder = YES` MPSGraph
// stores the values of the spatial block contiguosly within the
// `batchAxisTensor` dimension whereas without it they are stored interleaved
// with existing values in the `batchAxisTensor` dimension. Note: This
// operation is the inverse of
// [MPSGraph.SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName].
// Note: This operation is a generalization of
// [MPSGraph.DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/batchToSpace(_:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:)
func (g MPSGraph) BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("batchToSpaceTensor:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:"), tensor, spatialAxesTensor, batchAxisTensor, blockDimensionsTensor, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise bitwise AND of binary representations of two
// integer tensors.
//
// primaryTensor: The primary input tensor, must be of integer type.
//
// secondaryTensor: The secondary input tensor, must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseAND(_:_:name:)
func (g MPSGraph) BitwiseANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise left-shifted binary representations of the primary
// integer by the secondary tensor amount.
//
// primaryTensor: The primary input tensor, must be of integer type.
//
// secondaryTensor: The secondary input tensor, must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseLeftShift(_:_:name:)
func (g MPSGraph) BitwiseLeftShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseLeftShiftWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the bitwise NOT operation to the input tensor element.
//
// tensor: The input tensor, which must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation only accepts integer tensors.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseNOT(_:name:)
func (g MPSGraph) BitwiseNOTWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseNOTWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise bitwise OR of binary representations of two integer
// tensors.
//
// primaryTensor: The primary input tensor, must be of integer type.
//
// secondaryTensor: The secondary input tensor, must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseOR(_:_:name:)
func (g MPSGraph) BitwiseORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the population count of the input tensor elements.
//
// tensor: The input tensor, which must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation only accepts integer tensors, and returns the number of bits
// set in the input element.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwisePopulationCount(_:name:)
func (g MPSGraph) BitwisePopulationCountWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwisePopulationCountWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise right-shifted binary representations of the primary
// integer by the secondary tensor amount.
//
// primaryTensor: The primary input tensor, must be of integer type.
//
// secondaryTensor: The secondary input tensor, must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseRightShift(_:_:name:)
func (g MPSGraph) BitwiseRightShiftWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseRightShiftWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise bitwise XOR of binary representations of two
// integer tensors.
//
// primaryTensor: The primary input tensor, must be of integer type.
//
// secondaryTensor: The secondary input tensor, must be of integer type.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bitwiseXOR(_:_:name:)
func (g MPSGraph) BitwiseXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bitwiseXORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a BottomK operation and returns the value and indices tensors.
//
// source: Tensor containing source data.
//
// axis: The dimension along which to compute the BottomK values.
//
// k: The number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 2.
//
// # Discussion
//
// Finds the k smallest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the bottom values, and the
// second array corresponds to the indices of the bottom values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomK(_:axis:k:name:)
func (g MPSGraph) BottomKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("bottomKWithSourceTensor:axis:k:name:"), source, axis, k, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a BottomK operation and returns the result tensor.
//
// source: Tensor containing source data.
//
// axisTensor: Tensor containing the dimension along which to compute the BottomK values.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 2.
//
// # Discussion
//
// Finds the k smallest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the bottom values, and the
// second array corresponds to the indices of the bottom values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomK(_:axisTensor:kTensor:name:)
func (g MPSGraph) BottomKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("bottomKWithSourceTensor:axisTensor:kTensor:name:"), source, axisTensor, kTensor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a BottomKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// axis: The dimension along which to compute the BottomK values.
//
// k: The number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K smallest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomKGradient(_:source:axis:k:name:)
func (g MPSGraph) BottomKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bottomKWithGradientTensor:source:axis:k:name:"), gradient, source, axis, k, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a BottomKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// axisTensor: Tensor containing the dimension along which to compute the BottomK values.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K smallest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/bottomKGradient(_:source:axisTensor:kTensor:name:)
func (g MPSGraph) BottomKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("bottomKWithGradientTensor:source:axisTensor:kTensor:name:"), gradient, source, axisTensor, kTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a broadcast operation and returns the result tensor.
//
// tensor: The tensor to be broadcasted
//
// shape: The shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Broadcasts values inside the tensor, starting from the trailing dimensions,
// to give it the correct shape. This is equivalent to the broadcasting for
// arithmetic operations when operands have different shapes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/broadcast(_:shape:name:)
func (g MPSGraph) BroadcastTensorToShapeName(tensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("broadcastTensor:toShape:name:"), tensor, shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a broadcast operation and returns the result tensor.
//
// tensor: The Tensor to be broadcasted.
//
// shapeTensor: A rank-1 tensor of type [MPSDataTypeInt32] or [MPSDataTypeInt64] that
// defines the shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Broadcasts values inside the tensor, starting from the trailing dimensions,
// to give it the correct shape. This is equivalent to the broadcasting for
// arithmetic operations when operands have different shapes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/broadcast(_:shapeTensor:name:)
func (g MPSGraph) BroadcastTensorToShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("broadcastTensor:toShapeTensor:name:"), tensor, shapeTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an operation which invokes another executable.
//
// symbolName: The unique identifier used to find the executable in the
// `MPSGraphCompilationDescriptor.Callables()` directory.
//
// inputTensors: The tensors which are passed as inputs to the executable being invoked.
//
// outputTypes: The expected return types of the executable being invoked.
//
// name: Name of operation.
//
// # Return Value
//
// An array of valid [MPSGraphTensor] objects representing the return tensors
// of the invoked executable.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/call(symbolName:inputTensors:outputTypes:name:)
func (g MPSGraph) CallSymbolNameInputTensorsOutputTypesName(symbolName string, inputTensors []MPSGraphTensor, outputTypes []MPSGraphType, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("callSymbolName:inputTensors:outputTypes:name:"), objc.String(symbolName), objectivec.IObjectSliceToNSArray(inputTensors), objectivec.IObjectSliceToNSArray(outputTypes), objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a cast operation and returns the result tensor.
//
// tensor: The input tensor.
//
// type: The datatype to which MPSGraph casts the input.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Returns the input tensor casted to the specied data type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cast(_:to:name:)
func (g MPSGraph) CastTensorToTypeName(tensor IMPSGraphTensor, type_ uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("castTensor:toType:name:"), tensor, type_, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the ceiling operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/ceil(with:name:)
func (g MPSGraph) CeilWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("ceilWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Clamps the values in the first tensor between the corresponding values in
// the minimum and maximum value tensor.
//
// tensor: The tensor to be clamped.
//
// minValueTensor: The tensor with min values to clamp to.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a clamp operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/clamp(_:min:max:name:)
func (g MPSGraph) ClampWithTensorMinValueTensorMaxValueTensorName(tensor IMPSGraphTensor, minValueTensor IMPSGraphTensor, maxValueTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("clampWithTensor:minValueTensor:maxValueTensor:name:"), tensor, minValueTensor, maxValueTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a column to image operation and returns the result tensor.
//
// source: The tensor containing the source data. Must be of rank 4. The layout is
// defined by `descriptor.DataLayout()`.
//
// outputShape: The result tensor shape.
//
// descriptor: The descriptor object that specifies the parameters of the operation.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/colToIm(_:outputShape:descriptor:name:)
func (g MPSGraph) ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphImToColOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("colToImWithSourceTensor:outputShape:descriptor:name:"), source, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Compiles the graph for the given feeds to returns the target tensor values,
// ensuring all target operations would be executed.
//
// device: MPSGraph device to optimize for.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// compilationDescriptor: Compilation descriptor to set different compilation parameters.
//
// # Return Value
//
// # A valid MPSGraphExecutable object
//
// # Discussion
//
// This call blocks until execution has completed. The compilation descriptor
// helps specialize the executable returned.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/compile(with:feeds:targetTensors:targetOperations:compilationDescriptor:)
func (g MPSGraph) CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IMPSGraphDevice, feeds MPSGraphTensorShapedTypeDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, compilationDescriptor IMPSGraphCompilationDescriptor) IMPSGraphExecutable {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("compileWithDevice:feeds:targetTensors:targetOperations:compilationDescriptor:"), device, feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations), compilationDescriptor)
	return MPSGraphExecutableFromID(rv)
}

// Creates a complex constant op with the MPSDataTypeComplexFloat32 data type
// and returns the result tensor.
//
// realPart: The real part of the complex scalar to fill the entire tensor values with.
//
// imaginaryPart: The imaginary part of the complex scalar to fill the entire tensor values
// with.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:)
func (g MPSGraph) ConstantWithRealPartImaginaryPart(realPart float64, imaginaryPart float64) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithRealPart:imaginaryPart:"), realPart, imaginaryPart)
	return MPSGraphTensorFromID(rv)
}

// Creates a complex constant operation and returns the result tensor.
//
// realPart: The real part of the complex scalar to fill the entire tensor values with.
//
// imaginaryPart: The imaginary part of the complex scalar to fill the entire tensor values
// with.
//
// dataType: The dataType of the constant tensor.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:dataType:)
func (g MPSGraph) ConstantWithRealPartImaginaryPartDataType(realPart float64, imaginaryPart float64, dataType uint32) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithRealPart:imaginaryPart:dataType:"), realPart, imaginaryPart, dataType)
	return MPSGraphTensorFromID(rv)
}

// Creates a complex constant op with a given shape and returns the result
// tensor.
//
// realPart: The real part of the complex scalar to fill the entire tensor values with.
//
// imaginaryPart: The imaginary part of the complex scalar to fill the entire tensor values
// with.
//
// shape: The shape of the output tensor. This has to be statically shaped.
//
// dataType: The dataType of the constant tensor.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexConstant(realPart:imaginaryPart:shape:dataType:)
func (g MPSGraph) ConstantWithRealPartImaginaryPartShapeDataType(realPart float64, imaginaryPart float64, shape foundation.NSArray, dataType uint32) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithRealPart:imaginaryPart:shape:dataType:"), realPart, imaginaryPart, shape, dataType)
	return MPSGraphTensorFromID(rv)
}

// Returns a complex tensor from the two input tensors.
//
// realTensor: The real part of the complex tensor.
//
// imaginaryTensor: The imaginary part of the complex tensor.
//
// name: An optional string which serves as an identifier for the operation..
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/complexTensor(realTensor:imaginaryTensor:name:)
func (g MPSGraph) ComplexTensorWithRealTensorImaginaryTensorName(realTensor IMPSGraphTensor, imaginaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("complexTensorWithRealTensor:imaginaryTensor:name:"), realTensor, imaginaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a concatenation operation and returns the result tensor.
//
// tensor: The first tensor to concatenate.
//
// tensor2: The second tensor to concatenate.
//
// dimensionIndex: The dimension to concatenate across, must be in range: `-rank <= dimension
// < rank`.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Concatenates two input tensors along the specified dimension. Tensors must
// be broadcast compatible along all other dimensions, and have the same
// datatype.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensor(_:with:dimension:name:)
func (g MPSGraph) ConcatTensorWithTensorDimensionName(tensor IMPSGraphTensor, tensor2 IMPSGraphTensor, dimensionIndex int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("concatTensor:withTensor:dimension:name:"), tensor, tensor2, dimensionIndex, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a concatenation operation and returns the result tensor.
//
// tensors: The tensors to concatenate.
//
// dimensionIndex: The dimension to concatenate across, must be in range: `-rank <= dimension
// < rank`.
//
// interleave: A boolean value that specifies whether the operation interleaves input
// tensors.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Concatenates all input tensors along specified dimension. All inputs must
// be broadcast compatible along all other dimensions, and have the same type.
// When interleave is specified, all tensors will be interleaved. To
// interleave, make sure to provide broadcast compatible inputs along the
// specified dimension as well. For example:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensors(_:dimension:interleave:name:)
func (g MPSGraph) ConcatTensorsDimensionInterleaveName(tensors []MPSGraphTensor, dimensionIndex int, interleave bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("concatTensors:dimension:interleave:name:"), objectivec.IObjectSliceToNSArray(tensors), dimensionIndex, interleave, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a concatenation operation and returns the result tensor.
//
// tensors: The tensors to concatenate.
//
// dimensionIndex: The dimension to concatenate across, must be in range: `-rank <= dimension
// < rank`.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Concatenates all input tensors along the specified dimension. All inputs
// must be broadcast compatible along all other dimensions, and have the same
// datatype.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/concatTensors(_:dimension:name:)
func (g MPSGraph) ConcatTensorsDimensionName(tensors []MPSGraphTensor, dimensionIndex int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("concatTensors:dimension:name:"), objectivec.IObjectSliceToNSArray(tensors), dimensionIndex, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the complex conjugate of the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation..
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/conjugate(tensor:name:)
func (g MPSGraph) ConjugateWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("conjugateWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a constant operation and returns the result tensor.
//
// scalar: The scalar value to fill the entire tensor values with.
//
// dataType: The dataType of the constant tensor.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:dataType:)
func (g MPSGraph) ConstantWithScalarDataType(scalar float64, dataType uint32) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithScalar:dataType:"), scalar, dataType)
	return MPSGraphTensorFromID(rv)
}

// Creates a constant op with a given shape and returns the result tensor.
//
// scalar: The scalar value to fill the entire tensor values with.
//
// shape: The shape of the output tensor.
//
// dataType: The dataType of the constant tensor.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:shape:dataType:)-3wa0e
func (g MPSGraph) ConstantWithScalarShapeDataType(scalar float64, shape foundation.NSArray, dataType uint32) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithScalar:shape:dataType:"), scalar, shape, dataType)
	return MPSGraphTensorFromID(rv)
}

// Creates a constant op with a given shape and data, and returns the result
// tensor.
//
// data: The data for the tensor. The number of bytes should be
// sizeof(dataType)numberOfElements.
//
// shape: The shape of the output tensor. This has to be statically shaped.
//
// dataType: The dataType of theconstant tensor.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/constant(_:shape:dataType:)-ylr4
func (g MPSGraph) ConstantWithDataShapeDataType(data foundation.NSData, shape foundation.NSArray, dataType uint32) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("constantWithData:shape:dataType:"), data, shape, dataType)
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D (forward) convolution operation and returns the result tensor.
//
// source: Source tensor - must be a rank 4 tensor. The layout is defined by
// `descriptor.DataLayout()`.
//
// weights: Weights tensor, must be rank 4. The layout is defined by
// `descriptor.WeightsLayout()`.
//
// descriptor: Specifies strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2D(_:weights:descriptor:name:)
func (g MPSGraph) Convolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution2DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D convolution gradient operation with respect to the source
// tensor of the forward convolution.
//
// incomingGradient: Incoming loss gradient tensor
//
// weights: Forward pass weights tensor
//
// outputShape: Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [S] is source tensor to forward convolution, [R] is the result/returned
// tensor from forward convolution, and [L] is the loss function,
// `convolution2DDataGradientWithIncomingGradientTensor` returns tensor `dL/dS
// = dL/dR * dR/dS`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D convolution gradient operation with respect to the source
// tensor of the forward convolution.
//
// weights: Forward pass weights tensor
//
// outputShapeTensor: 4D Int32 or Int64 tensor. Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [S] is source tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution2DDataGradientWithIncomingGradientTensor returns tensor `dL/dS =
// dL/dR * dR/dS`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, weights, outputShapeTensor, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D convolution gradient operation with respect to the weights
// tensor of the forward convolution.
//
// incomingGradient: Incoming loss gradient tensor
//
// outputShape: Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [W] is weights tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution2DWeightsGradientWithIncomingGradientTensor returns tensor
// `dL/dW = dL/dR * dR/dW`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DWeightsGradient(_:source:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, source, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D convolution gradient operation with respect to weights tensor
// of forward convolution.
//
// outputShapeTensor: 4D int32 or Int64 Tensor. Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [W] is weights tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution2DWeightsGradientWithIncomingGradientTensor returns tensor
// `dL/dW = dL/dR * dR/dW`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution2DWeightsGradient(_:source:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, source, outputShapeTensor, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D forward convolution operation and returns the result tensor.
//
// source: Source tensor - must be of rank 5. The layout is defined by
// `descriptor.DataLayout()`.
//
// weights: Weights tensor, must be rank 5. The layout is defined by
// `descriptor.WeightsLayout()`.
//
// descriptor: Specifies strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3D(_:weights:descriptor:name:)
func (g MPSGraph) Convolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution3DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D convolution gradient operation with respect to the source
// tensor of the forward convolution.
//
// incomingGradient: Incoming loss gradient tensor
//
// weights: Forward pass weights tensor
//
// outputShape: Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [S] is source tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution3DDataGradientWithIncomingGradientTensor returns tensor `dL/dS =
// dL/dR * dR/dS`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D convolution gradient operation with respect to the source
// tensor of the forward convolution.
//
// weights: Forward pass weights tensor
//
// outputShapeTensor: 4D Int32 or Int64 tensor. Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [S] is source tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution3DDataGradientWithIncomingGradientTensor returns tensor `dL/dS =
// dL/dR * dR/dS`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, weights IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, weights, outputShapeTensor, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D convolution gradient operation with respect to the weights
// tensor of the forward convolution.
//
// incomingGradient: Incoming loss gradient tensor
//
// outputShape: Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [W] is weights tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution3DWeightsGradientWithIncomingGradientTensor returns tensor
// `dL/dW = dL/dR * dR/dW`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DWeightsGradient(_:source:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, source, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D convolution gradient operation with respect to the weights
// tensor of the forward convolution.
//
// outputShapeTensor: 4D int32 or Int64 Tensor. Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward convolution 2D op `descriptor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// If [W] is weights tensor to forward convolution, [R] is the result/returned
// tensor of forward convolution, and [L] is the loss function,
// convolution3DWeightsGradientWithIncomingGradientTensor returns tensor
// `dL/dW = dL/dR * dR/dW`, where `dL/dR` is the incomingGradient parameter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolution3DWeightsGradient(_:source:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) Convolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, outputShapeTensor IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), gradient, source, outputShapeTensor, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose operation and returns the result tensor.
//
// source: Input tensor
//
// weights: Weights tensor
//
// outputShape: Shape of the result tensor.
//
// descriptor: Descriptor for the corresponding forward 2D-convolution operation
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Convolution Tranpose operation is exactly the same as convolution gradint
// with respect to input image
// `convolution2DDataGradientWithIncomingGradient`. Weights tensor and source
// tensors are interpreted as they are in
// `convolution2DDataGradientWithIncomingGradient`. Convolution with stride
// `s` downsamples source tensor by factor `s` in spatial dimensions whereas
// convolution tranpose with stride `s` upsamples source tensor by factor `s`.
// Convolution transpose can map the same source size to multiple destination
// sizes. The relationship between the width of the source and the width of
// the destination is `(sourceWidth - 1)stride + 1 + (kernelWidth -
// 1)dilationRate <= destinationWidth + paddingLeft + paddingRight` so there
// are stride -1 values of the width of the destination that give same width
// of the source. In order to disambiguate, outputShape parameter is used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2D(_:weights:outputShape:descriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DWithSourceTensor:weightsTensor:outputShape:descriptor:name:"), source, weights, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose operation and returns the result tensor.
//
// source: Input tensor
//
// weights: Weights tensor
//
// outputShape: 1D Int32 or Int64 tensor. shape of the result tensor.
//
// descriptor: Descriptor for the corresponding forward Conv2D operation
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2D(_:weights:outputShapeTensor:descriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DWithSourceTensor:weightsTensor:outputShapeTensor:descriptor:name:"), source, weights, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose gradient operation with respect to the
// source tensor of convolution transpose operation and returns the result
// tensor.
//
// incomingGradient: Incoming gradient tensor
//
// weights: Forward pass weights tensor
//
// outputShape: Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward pass op descriptor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Inserts an operation in graph to compute gradient of convolution transpose
// with respect to source tensor of the corresponding convolution transpose
// operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DDataGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose gradient operation with respect to the
// source tensor of convolution transpose operation and returns the result
// tensor.
//
// incomingGradient: Incoming gradient tensor
//
// weights: Forward pass weights tensor
//
// outputShape: 1D Int32 or Int64 Tensor. Shape of the forward pass source tensor
//
// forwardConvolutionDescriptor: Forward pass op descriptor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Inserts an operation in graph to compute gradient of convolution transpose
// with respect to source tensor of the corresponding convolution transpose
// operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DDataGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), incomingGradient, weights, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose gradient operation with respect to the
// weights tensor of the convolution transpose operation and returns the
// result tensor.
//
// incomingGradientTensor: Incoming gradient tensor
//
// source: Forward pass source tensor
//
// outputShape: Shape of the forward pass source weights tensor
//
// forwardConvolutionDescriptor: Forward pass op descriptor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Inserts an operation in graph to compute gradient of convolution transpose
// with respect to the weights tensor of the corresponding convolution
// transpose operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DWeightsGradient(_:weights:outputShape:forwardConvolutionDescriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:forwardConvolutionDescriptor:name:"), incomingGradientTensor, source, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a convolution transpose gradient operation with respect to the
// weights tensor of the convolution transpose operation and returns the
// result tensor.
//
// incomingGradientTensor: Incoming gradient tensor
//
// source: Forward pass source tensor
//
// outputShape: 1D Int32 or Int64 Tensor. Shape of the forward pass source weights tensor
//
// forwardConvolutionDescriptor: Forward pass op descriptor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Inserts an operation in graph to compute gradient of convolution transpose
// with respect to the weights tensor of the corresponding convolution
// transpose operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/convolutionTranspose2DWeightsGradient(_:weights:outputShapeTensor:forwardConvolutionDescriptor:name:)
func (g MPSGraph) ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor IMPSGraphTensor, source IMPSGraphTensor, outputShape IMPSGraphTensor, forwardConvolutionDescriptor IMPSGraphConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("convolutionTranspose2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), incomingGradientTensor, source, outputShape, forwardConvolutionDescriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a get-coordindate operation and returns the result tensor.
//
// axis: The coordinate axis an element’s value is set to. Negative values wrap
// around.
//
// shape: The shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of specified shape with value at index `[i_0, i_1, ... ,
// i_N] = i_axis` For example,
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxis:withShape:name:)
func (g MPSGraph) CoordinateAlongAxisWithShapeName(axis int, shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("coordinateAlongAxis:withShape:name:"), axis, shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a get-coordindate operation and returns the result tensor.
//
// axis: The coordinate axis an element’s value is set to. Negative values wrap
// around.
//
// shapeTensor: A rank-1 tensor of type [MPSDataTypeInt32] or [MPSDataTypeInt64] that
// defines the shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// See [MPSGraph.CoordinateAlongAxisWithShapeName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxis:withShapeTensor:name:)
func (g MPSGraph) CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("coordinateAlongAxis:withShapeTensor:name:"), axis, shapeTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a get-coordindate operation and returns the result tensor.
//
// axisTensor: A Scalar tensor of type [MPSDataTypeInt32], that specifies the coordinate
// axis an element’s value is set to. Negative values wrap around.
//
// shape: The shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// See [MPSGraph.CoordinateAlongAxisWithShapeName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxisTensor:withShape:name:)
func (g MPSGraph) CoordinateAlongAxisTensorWithShapeName(axisTensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("coordinateAlongAxisTensor:withShape:name:"), axisTensor, shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a get-coordindate operation and returns the result tensor.
//
// axisTensor: A Scalar tensor of type [MPSDataTypeInt32], that specifies the coordinate
// axis an element’s value is set to. Negative values wrap around.
//
// shapeTensor: A rank-1 tensor of type [MPSDataTypeInt32] or [MPSDataTypeInt64] that
// defines the shape of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// See [MPSGraph.CoordinateAlongAxisWithShapeName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/coordinate(alongAxisTensor:withShapeTensor:name:)
func (g MPSGraph) CoordinateAlongAxisTensorWithShapeTensorName(axisTensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("coordinateAlongAxisTensor:withShapeTensor:name:"), axisTensor, shapeTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the cosine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cos(with:name:)
func (g MPSGraph) CosWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cosWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the hyperbolic cosine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cosh(with:name:)
func (g MPSGraph) CoshWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("coshWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative maximum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to the lowest value of the tensor data type
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axis:exclusive:reverse:name:)
func (g MPSGraph) CumulativeMaximumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMaximumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative maximum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axis:name:)
func (g MPSGraph) CumulativeMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMaximumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative maximum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to the lowest value of the tensor data type
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axisTensor:exclusive:reverse:name:)
func (g MPSGraph) CumulativeMaximumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMaximumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative maximum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMaximum(_:axisTensor:name:)
func (g MPSGraph) CumulativeMaximumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMaximumWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative minimum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to the largest value of the tensor data type
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axis:exclusive:reverse:name:)
func (g MPSGraph) CumulativeMinimumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMinimumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative minimum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axis:name:)
func (g MPSGraph) CumulativeMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMinimumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative minimum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to the largest value of the tensor data type
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axisTensor:exclusive:reverse:name:)
func (g MPSGraph) CumulativeMinimumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMinimumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative minimum of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeMinimum(_:axisTensor:name:)
func (g MPSGraph) CumulativeMinimumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeMinimumWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative product of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to one
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axis:exclusive:reverse:name:)
func (g MPSGraph) CumulativeProductWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeProductWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative product of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axis:name:)
func (g MPSGraph) CumulativeProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeProductWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative product of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to one
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axisTensor:exclusive:reverse:name:)
func (g MPSGraph) CumulativeProductWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeProductWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative product of the input tensor along the specified
// axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeProduct(_:axisTensor:name:)
func (g MPSGraph) CumulativeProductWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeProductWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative sum of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to zero
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axis:exclusive:reverse:name:)
func (g MPSGraph) CumulativeSumWithTensorAxisExclusiveReverseName(tensor IMPSGraphTensor, axis int, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeSumWithTensor:axis:exclusive:reverse:name:"), tensor, axis, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative sum of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axis:name:)
func (g MPSGraph) CumulativeSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeSumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative sum of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// exclusive: If true, perform the exclusive cumulative operation, and the first element
// will be equal to zero
//
// reverse: If true, reverse the direction of the cumulative operation along the
// specified axis
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axisTensor:exclusive:reverse:name:)
func (g MPSGraph) CumulativeSumWithTensorAxisTensorExclusiveReverseName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, exclusive bool, reverse bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeSumWithTensor:axisTensor:exclusive:reverse:name:"), tensor, axisTensor, exclusive, reverse, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the cumulative sum of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension where you compute the cumulative operation
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/cumulativeSum(_:axisTensor:name:)
func (g MPSGraph) CumulativeSumWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cumulativeSumWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a depth-to-space2D operation and returns the result tensor.
//
// tensor: The input tensor.
//
// widthAxis: The axis that defines the fastest running dimension within the block.
//
// heightAxis: The axis that defines the 2nd fastest running dimension within the block.
//
// depthAxis: The axis that defines the destination dimension, where to copy the blocks.
//
// blockSize: The size of the square spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls the layout of the sub-blocks within the depth
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the input tensor, where values from the
// `depthAxis` dimension are moved in spatial blocks of size `blockSize` to
// the `heightAxis` and `widthAxis` dimensions. Use the `usePixelShuffleOrder`
// parameter to control how the data within spatial blocks is ordered in the
// `depthAxis` dimension: with `usePixelShuffleOrder = YES` MPSGraph stores
// the values of the spatial block contiguosly within the `depthAxis`
// dimension, whereas without it they are stored interleaved with existing
// values in the `depthAxisTensor` dimension. This operation is the inverse of
// [MPSGraph.SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depth(toSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:)
func (g MPSGraph) DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthToSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a depth-to-space2D operation and returns the result tensor.
//
// tensor: The input tensor.
//
// widthAxisTensor: A scalar tensor that contains the axis that defines the fastest running
// dimension within the block.
//
// heightAxisTensor: A scalar tensor that contains the axis that defines the 2nd fastest running
// dimension within the block.
//
// depthAxisTensor: A scalar tensor that contains the axis that defines the destination
// dimension, where to copy the blocks.
//
// blockSize: The size of the square spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls the layout of the sub-blocks within the depth
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the input tensor, where values from the
// `depthAxisTensor` dimension are moved in spatial blocks of size `blockSize`
// to the `heightAxisTensor` and `widthAxisTensor` dimensions. Use the
// `usePixelShuffleOrder` parameter to control how the data within spatial
// blocks is ordered in the `depthAxisTensor` dimension: with
// `usePixelShuffleOrder = YES` MPSGraph stores the values of the spatial
// block contiguosly within the `depthAxisTensor` dimension, whereas without
// it they are stored interleaved with existing values in the
// `depthAxisTensor` dimension. This operation is the inverse of
// [MPSGraph.SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depth(toSpace2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:)
func (g MPSGraph) DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthToSpace2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxisTensor, heightAxisTensor, depthAxisTensor, blockSize, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D-depthwise convolution operation and returns the result tensor.
//
// source: A 2D Image source as tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// weights: The weights tensor, must be rank=4. The layout is defined by
// `descriptor.WeightsLayout()`.
//
// descriptor: The descriptor object that specifies strides, dilation rates, paddings and
// layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2D(_:weights:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution2DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D-depthwise convolution gradient for data operation and returns
// the result tensor.
//
// incomingGradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// weights: The weights tensor, must be rank=4. The layout is defined by
// `descriptor.WeightsLayout()`.
//
// outputShape: The shape of the οutput tensor (and therefore input tensor of forward
// pass).
//
// descriptor: The descriptor object that specifies strides, dilation rates, paddings and
// layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2DDataGradient(_:weights:outputShape:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), incomingGradient, weights, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D-depthwise convolution gradient for weights operation and
// returns the result tensor.
//
// incomingGradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// source: A 2D Image source as tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// outputShape: The shape of the οutput tensor (and therefore weight tensor of forward
// pass).
//
// descriptor: The descriptor object that specifies strides, dilation rates, paddings and
// layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution2DWeightsGradient(_:source:outputShape:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), incomingGradient, source, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D depthwise convolution operation and returns the result tensor.
//
// source: A 3D Image source as tensor - must be at least rank=4 (CDHW when
// channelDimensionIndex = -4).
//
// weights: The weights tensor, must be rank=4 - axes are interpreted as CDHW when
// channelDimensionIndex = -4 .
//
// descriptor: The descriptor object that specifies strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Works exactly like depthwise convolution2D, but in three dimensions.
// Supports different layouts with the
// [MPSGraphDepthwiseConvolution3DOpDescriptor.ChannelDimensionIndex]
// property. If your weights need a different layout add a permute operation
// on them before this operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3D(_:weights:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution3DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D depthwise convolution gradient for data operation and returns
// the result tensor.
//
// incomingGradient: A 3D input gradient tensor - must be at least rank=4 (CDHW).
//
// weights: The weights tensor, must be rank=4 - axes are interpreted as CDHW.
//
// outputShape: The shape of the οutput tensor (and therefore input tensor of forward
// pass).
//
// descriptor: The descriptor object that specifies strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3DDataGradient(_:weights:outputShape:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), incomingGradient, weights, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 3D depthwise convolution gradient for weights operation and
// returns the result tensor.
//
// incomingGradient: A 3D input gradient tensor - must be at least rank=4 (NCDHW).
//
// source: The forward pass 3D Image source as tensor - must be at least rank=4
// (NCDHW).
//
// outputShape: The shape of the οutput tensor (and therefore weight tensor of forward
// pass).
//
// descriptor: The descriptor object that specifies strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3DWeightsGradient(_:source:outputShape:descriptor:name:)
func (g MPSGraph) DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, source IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("depthwiseConvolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), incomingGradient, source, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a vector lookup-table based quantization operation and returns the
// result tensor.
//
// tensor: Input tensor to be dequantized.
//
// LUTTensor: The lookup table to use - for u4 the second to last dimension should have
// 16 elements, and for u8 256 elements.
//
// axis: Axis on which the scale 1D value is being broadcasted.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// Converts a u8 or u4 `tensor` to a float tensor by applying a lookup
// operation, where each input index defines a vector of values. The operation
// reads the vector values from the last dimension of the lookup table tensor
// and stores them into the dimension defined by `axis` on the result tensor.
//
// Note: The operation supports LUT groups up to the last 2 dimensions for
// `tensor`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:LUTTensor:axis:name:)
func (g MPSGraph) DequantizeTensorLUTTensorAxisName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:LUTTensor:axis:name:"), tensor, LUTTensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a lookup-table based quantization operation and returns the result
// tensor.
//
// tensor: Input tensor to be dequantized.
//
// LUTTensor: The lookup table to use - for u4 the last dimension should have 16
// elements, and for u8 256 elements.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// Converts a u8 or u4 `tensor` to a float tensor by applying a lookup
// operation:
//
// Note: The operation supports LUT groups up to the last 3 dimensions for
// `tensor`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:LUTTensor:name:)
func (g MPSGraph) DequantizeTensorLUTTensorName(tensor IMPSGraphTensor, LUTTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:LUTTensor:name:"), tensor, LUTTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates Dequantize operation and returns the result tensor.
//
// tensor: Input tensor to be dequantized
//
// scale: Scale scalar parameter
//
// zeroPoint: Bias scalar parameter (converted to dataType of tensor)
//
// dataType: Float data type of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the i8 or u8 `tensor` to a float tensor by applying a scale + bias
// transform: result = scale(tensor - zeroPoint)
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scale:zeroPoint:dataType:name:)
func (g MPSGraph) DequantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:scale:zeroPoint:dataType:name:"), tensor, scale, zeroPoint, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a dequantize operation and returns the result tensor.
//
// tensor: Input tensor to be dequantized.
//
// scaleTensor: Scale Tensor parameter with groups support.
//
// dataType: Float data type of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of datatype `dataType`.
//
// # Discussion
//
// Converts the i8, u8, i4 or u4 `tensor` to a float tensor by applying a
// scale and bias transform:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:dataType:name:)
func (g MPSGraph) DequantizeTensorScaleTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:scaleTensor:dataType:name:"), tensor, scaleTensor, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates Dequantize operation and returns the result tensor.
//
// tensor: Input tensor to be dequantized
//
// scaleTensor: Scale scalar or 1D Tensor parameter with size == tensor.shape[axis]
//
// zeroPoint: Bias scalar parameter (converted to dataType of tensor)
//
// dataType: Float data type of the result tensor.
//
// axis: Axis on which the scale 1D value is being broadcasted
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the i8 or u8 `tensor` to a float tensor by applying a scale + bias
// transform: result = scaleTensor(tensor - zeroPoint)
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPoint:dataType:axis:name:)
func (g MPSGraph) DequantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType uint32, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPoint:dataType:axis:name:"), tensor, scaleTensor, zeroPoint, dataType, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a dequantize operation and returns the result tensor.
//
// tensor: Input tensor to be dequantized
//
// scaleTensor: Scale scalar or 1D Tensor parameter with size == tensor.shape[axis]
//
// zeroPointTensor: Bias scalar or 1D Tensor parameter with size == tensor.shape[axis]
//
// dataType: Float data type of the result tensor.
//
// axis: Axis on which the scale 1D value is being broadcasted
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the i8 or u8 `tensor` to a float tensor by applying a scale + bias
// transform: result = scaleTensor(tensor - zeroPointTensor)
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPointTensor:dataType:axis:name:)
func (g MPSGraph) DequantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPointTensor:dataType:axis:name:"), tensor, scaleTensor, zeroPointTensor, dataType, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a dequantize operation and returns the result tensor.
//
// tensor: Input tensor to be dequantized.
//
// scaleTensor: The scale tensor with groups support.
//
// zeroPointTensor: The bias tensor with groups support.
//
// dataType: Float data type of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array of datatype `dataType`.
//
// # Discussion
//
// Convert the i8, u8, i4 or u4 `tensor` to a float tensor by applying a scale
// and bias transform:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dequantize(_:scaleTensor:zeroPointTensor:dataType:name:)
func (g MPSGraph) DequantizeTensorScaleTensorZeroPointTensorDataTypeName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dequantizeTensor:scaleTensor:zeroPointTensor:dataType:name:"), tensor, scaleTensor, zeroPointTensor, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Divides the first input tensor by the second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a divide operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/division(_:_:name:)
func (g MPSGraph) DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("divisionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Divides the first input tensor by the second, with the result being 0 if
// the denominator is 0.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/divisionNoNaN(_:_:name:)
func (g MPSGraph) DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("divisionNoNaNWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a dropout operation and returns the result
//
// tensor: Input tensor
//
// rate: The rate of values to be set to 0
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Removes values in the `tensor` with a percentage chance equal to `rate`.
// Removed values are set to 0
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dropout(_:rate:name:)-16cq4
func (g MPSGraph) DropoutTensorRateTensorName(tensor IMPSGraphTensor, rate IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dropoutTensor:rateTensor:name:"), tensor, rate, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a dropout operation and returns the result
//
// tensor: Input tensor
//
// rate: The rate of values to be set to 0
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Removes values in the `tensor` with a percentage chance equal to `rate`.
// Removed values are set to 0
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/dropout(_:rate:name:)-6hvf3
func (g MPSGraph) DropoutTensorRateName(tensor IMPSGraphTensor, rate float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dropoutTensor:rate:name:"), tensor, rate, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Encodes the graph for the given feeds to returns the target tensor values
// in the results dictionary provided by the user.
//
// commandBuffer: commandBuffer passed to execute the graph on, commitAndContinue might be
// called, please don’t rely on underlying MTLCommandBuffer to remain
// uncommitted.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetOperations: Operations to be completed at the end of the run.
//
// resultsDictionary: MPSGraphTensors dictionary passed by user, these will be filled with graph
// output data.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Discussion
//
// It ensures all target operations also executed. This call is asynchronous
// and will return immediately if a completionHandler is set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/encode(to:feeds:targetOperations:resultsDictionary:executionDescriptor:)
func (g MPSGraph) EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary, executionDescriptor IMPSGraphExecutionDescriptor) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:feeds:targetOperations:resultsDictionary:executionDescriptor:"), commandBuffer.ID, feeds, objectivec.IObjectSliceToNSArray(targetOperations), resultsDictionary, executionDescriptor)
}

// Encodes the graph for the given feeds to returns the target tensor values,
// ensuring all target operations also executed.
//
// commandBuffer: commandBuffer passed to exectute the graph on, it is an MPSCommandBuffer,
// commitAndContinue might be called, please don’t rely on underlying
// MTLCommandBuffer to remain uncommitted.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Return Value
//
// A valid MPSGraphTensor : MPSGraphTensorData dictionary with results
// synchronized to the CPU memory if MPSGraphOptionsSynchronizeResults set.
//
// # Discussion
//
// This call is asynchronous and will return immediately if a
// completionHandler is set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/encode(to:feeds:targetTensors:targetOperations:executionDescriptor:)
func (g MPSGraph) EncodeToCommandBufferFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandBuffer *metalperformanceshaders.MPSCommandBuffer, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary {
	rv := objc.Send[MPSGraphTensorDataDictionary](g.ID, objc.Sel("encodeToCommandBuffer:feeds:targetTensors:targetOperations:executionDescriptor:"), commandBuffer.ID, feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations), executionDescriptor)
	return MPSGraphTensorDataDictionary(rv)
}

// Returns the elementwise equality check of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a equal operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/equal(_:_:name:)
func (g MPSGraph) EqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("equalWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the error function to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/erf(with:name:)
func (g MPSGraph) ErfWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("erfWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an expand-dimensions operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axes: The axes to expand.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Expands the tensor, inserting dimensions with size 1 at specified axes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axes:name:)
func (g MPSGraph) ExpandDimsOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("expandDimsOfTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an expand-dimensions operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axesTensor: The tensor containing the axes to expand.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Expands the tensor, inserting dimensions with size 1 at specified axes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axesTensor:name:)
func (g MPSGraph) ExpandDimsOfTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("expandDimsOfTensor:axesTensor:name:"), tensor, axesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an expand-dimensions operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axis: The axis to expand.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Expands the tensor, inserting a dimension with size 1 at the specified
// axis.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/expandDims(_:axis:name:)
func (g MPSGraph) ExpandDimsOfTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("expandDimsOfTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the natural exponent to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponent(with:name:)
func (g MPSGraph) ExponentWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("exponentWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies an exponent with base 10 to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponentBase10(with:name:)
func (g MPSGraph) ExponentBase10WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("exponentBase10WithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies an exponent with base 2 to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/exponentBase2(with:name:)
func (g MPSGraph) ExponentBase2WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("exponentBase2WithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a fast Fourier transform operation and returns the result tensor.
//
// tensor: A complex or real-valued input tensor.
//
// axes: An array of numbers that specifies over which axes MPSGraph performs the
// Fourier transform - all axes must be contained within last four dimensions
// of the input tensor.
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid complex-valued MPSGraphTensor of the same shape as `tensor`.
//
// # Discussion
//
// This operation computes the fast Fourier transform of the input tensor
// according to the following formulae.
//
// `scale = 1` for `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode =
// size`, `scale = 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is
// the volume of the transformation defined by the dimensions included in
// `axes` (`V_f = prod_{i \in axes} shape(input)[i]`) (see
// [MPSGraphFFTDescriptor.ScalingMode]), `+` is selected in `+/-` when
// `inverse` is specified, otherwise `-` is used and the sum is done
// separately over each dimension in `axes` and `n` is the dimension length of
// that axis.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/fastFourierTransform(_:axes:descriptor:name:)
func (g MPSGraph) FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("fastFourierTransformWithTensor:axes:descriptor:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a fast Fourier transform operation and returns the result tensor.
//
// tensor: A complex or real-valued input tensor.
//
// axesTensor: A tensor of rank one containing the axes over which MPSGraph performs the
// transformation. See
// [MPSGraph.FastFourierTransformWithTensorAxesDescriptorName].
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid complex-valued MPSGraphTensor of the same shape as `tensor`.
//
// # Discussion
//
// This operation computes the fast Fourier transform of the input tensor
// according to the following formulae.
//
// `scale = 1` for `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode =
// size`, `scale = 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is
// the volume of the transformation defined by the dimensions included in
// `axes` (`V_f = prod_{i \in axes} shape(input)[i]`) (see
// [MPSGraphFFTDescriptor.ScalingMode]), `+` is selected in `+/-` when
// `inverse` is specified, otherwise `-` is used and the sum is done
// separately over each dimension in `axes` and `n` is the dimension length of
// that axis.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/fastFourierTransform(_:axesTensor:descriptor:name:)
func (g MPSGraph) FastFourierTransformWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("fastFourierTransformWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a flatten2D operation and returns the result tensor.
//
// tensor: The tensor to be flattened.
//
// axis: The axis around which to flatten.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Flattens dimensions before `axis` to `result[0]` and dimensions starting
// from `axis` to `result[1]` and returns a rank-2 tensor as result.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/flatten2D(_:axis:name:)
func (g MPSGraph) Flatten2DTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("flatten2DTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a flatten2D operation and returns the result tensor.
//
// tensor: The tensor to be flattened.
//
// axisTensor: A scalar tensor that contains the axis around which to flatten.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Flattens dimensions before `axis` to `result[0]` and dimensions starting
// from `axis` to `result[1]` and returns a rank-2 tensor as result.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/flatten2D(_:axisTensor:name:)
func (g MPSGraph) Flatten2DTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("flatten2DTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the floor operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/floor(with:name:)
func (g MPSGraph) FloorWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("floorWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the remainder of floor divison between the primary and secondary
// tensor.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// Creates a floorModulo operation and returns the result tensor, it supports
// broadcasting as well, returns 0 if divisor is 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/floorModulo(_:_:name:)
func (g MPSGraph) FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("floorModuloWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Gather operation and returns the result tensor.
//
// updatesTensor: Tensor containing slices to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the updates indices to read slices from
//
// axis: The dimension on which to perform the gather
//
// batchDimensions: The number of batch dimensions
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Gathers the values in updatesTensor to the result tensor along the indices
// in indicesTensor. The gather is defined as
//
// # The tensors have the following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gather(withUpdatesTensor:indicesTensor:axis:batchDimensions:name:)
func (g MPSGraph) GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis uint, batchDimensions uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gatherWithUpdatesTensor:indicesTensor:axis:batchDimensions:name:"), updatesTensor, indicesTensor, axis, batchDimensions, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a GatherAlongAxis operation and returns the result tensor.
//
// axis: The axis to gather from. Negative values wrap around
//
// updatesTensor: The input tensor to gather values from
//
// indicesTensor: Int32 or Int64 tensor used to index `updatesTensor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Gather values from `updatesTensor` along the specified `axis` at indices in
// `indicesTensor`. The shape of `updatesTensor` and `indicesTensor` must
// match except at `axis`. The shape of the result tensor is equal to the
// shape of `indicesTensor`. If an index is out of bounds of the
// `updatesTensor` along `axis` a 0 is inserted.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherAlongAxis(_:updates:indices:name:)
func (g MPSGraph) GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gatherAlongAxis:withUpdatesTensor:indicesTensor:name:"), axis, updatesTensor, indicesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a GatherAlongAxis operation and returns the result tensor.
//
// axisTensor: Scalar Int32 tensor. The axis to gather from. Negative values wrap around
//
// updatesTensor: The input tensor to gather values from
//
// indicesTensor: Int32 or Int64 tensor used to index `updatesTensor`
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Gather values from `updatesTensor` along the specified `axis` at indices in
// `indicesTensor`. The shape of `updatesTensor` and `indicesTensor` must
// match except at `axis`. The shape of the result tensor is equal to the
// shape of `indicesTensor`. If an index is out of bounds of the
// `updatesTensor` along `axis` a 0 is inserted.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherAlongAxisTensor(_:updates:indices:name:)
func (g MPSGraph) GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gatherAlongAxisTensor:withUpdatesTensor:indicesTensor:name:"), axisTensor, updatesTensor, indicesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a GatherND operation and returns the result tensor.
//
// updatesTensor: Tensor containing slices to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the updates indices to read slices from
//
// batchDimensions: The number of batch dimensions
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Gathers the slices in updatesTensor to the result tensor along the indices
// in indicesTensor. The gather is defined as
//
// # The tensors have the following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gatherND(withUpdatesTensor:indicesTensor:batchDimensions:name:)
func (g MPSGraph) GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gatherNDWithUpdatesTensor:indicesTensor:batchDimensions:name:"), updatesTensor, indicesTensor, batchDimensions, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Calculates a partial derivative of primaryTensor with respect to the
// tensors.
//
// primaryTensor: Tensor to be differentiated (numerator).
//
// tensors: Tensors to do the differentiation with (denominator).
//
// name: Name for the gradient operation.
//
// # Return Value
//
// A valid MPSGraphTensor dictionary object containing partial derivative
// d(primaryTensor)/d(secondaryTensor) for each tensor as key.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/gradients(of:with:name:)
func (g MPSGraph) GradientForPrimaryTensorWithTensorsName(primaryTensor IMPSGraphTensor, tensors []MPSGraphTensor, name string) foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gradientForPrimaryTensor:withTensors:name:"), primaryTensor, objectivec.IObjectSliceToNSArray(tensors), objc.String(name))
	return foundation.NSDictionaryFromID(rv)
}

// Checks in an elementwise manner if the first input tensor is greater than
// the second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a `greaterThan` operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/greaterThan(_:_:name:)
func (g MPSGraph) GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("greaterThanWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks in an elementwise manner if the first input tensor is greater than
// or equal to the second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a `greaterThanOrEqual` operation and returns the
// result tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/greaterThanOrEqualTo(_:_:name:)
func (g MPSGraph) GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("greaterThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Copies the input tensor values into the output, behaving as an identity
// operation.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object which is a copy of the input.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/identity(with:name:)
func (g MPSGraph) IdentityWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("identityWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates an imToCol operation and returns the result tensor.
//
// source: The tensor containing the source data. Must be of rank 4. The layout is
// defined by `descriptor.DataLayout()`.
//
// descriptor: The descriptor object that specifies the parameters of the operation.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/imToCol(_:descriptor:name:)
func (g MPSGraph) ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("imToColWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the imaginary part of a tensor.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation..
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/imaginaryPartOfTensor(tensor:name:)
func (g MPSGraph) ImaginaryPartOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("imaginaryPartOfTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the inverse of an input tensor.
//
// inputTensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the inverse of the input tensor.
//
// # Discussion
//
// The framework computes the inverse of a square matrix by calling LU
// decomposition and LU solver. All dimensions after the first 2 are treated
// as batch dimensions and the inverse for each batch is computed. Results are
// undefined for ill conditioned matrices.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/inverse(input:name:)
func (g MPSGraph) InverseOfTensorName(inputTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inverseOfTensor:name:"), inputTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks if the input tensor elements are finite or not.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// If the input tensor element is finite, the operation returns `true`, else
// it returns `false`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isFinite(with:name:)
func (g MPSGraph) IsFiniteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isFiniteWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks if the input tensor elements are infinite or not.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// If the input tensor element is infinite, the operation returns `true`, else
// it returns `false`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isInfinite(with:name:)
func (g MPSGraph) IsInfiniteWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isInfiniteWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks if the input tensor elements are [NaN] or not.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// If the input tensor element is [NaN], the operation returns `true`, else it
// returns `false`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/isNaN(with:name:)
func (g MPSGraph) IsNaNWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isNaNWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the leaky rectified linear unit (ReLU) activation function on the
// input tensor.
//
// tensor: An input tensor.
//
// alpha: The scalar value alpha used by all elements in the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// # Discussion
//
// The operation is: f(x) = max(x, alpha).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLU(with:alpha:name:)
func (g MPSGraph) LeakyReLUWithTensorAlphaName(tensor IMPSGraphTensor, alpha float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leakyReLUWithTensor:alpha:name:"), tensor, alpha, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the leaky rectified linear unit (ReLU) activation function on the
// input tensor.
//
// tensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// # Discussion
//
// The operation is: f(x) = max(x, alpha). This operation supports
// broadcasting with the alpha tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLU(with:alphaTensor:name:)
func (g MPSGraph) LeakyReLUWithTensorAlphaTensorName(tensor IMPSGraphTensor, alphaTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leakyReLUWithTensor:alphaTensor:name:"), tensor, alphaTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the gradient of the leaky rectified linear unit (ReLU) activation.
//
// gradient: The incoming gradient tensor.
//
// source: The input tensor in forward pass.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// # Discussion
//
// This operation supports broadcasting with the alpha tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/leakyReLUGradient(withIncomingGradient:sourceTensor:alphaTensor:name:)
func (g MPSGraph) LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, alphaTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leakyReLUGradientWithIncomingGradient:sourceTensor:alphaTensor:name:"), gradient, source, alphaTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks in an elementwise manner if the first input tensor is less than the
// second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a `lessThan` operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/lessThan(_:_:name:)
func (g MPSGraph) LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("lessThanWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Checks in an elementwise manner if the first input tensor is less than or
// equal to the second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a `lessThanOrEqualTo` operation and returns the
// result tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/lessThanOrEqualTo(_:_:name:)
func (g MPSGraph) LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("lessThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the natural logarithm to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithm(with:name:)
func (g MPSGraph) LogarithmWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logarithmWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the logarithm with base 10 to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithmBase10(with:name:)
func (g MPSGraph) LogarithmBase10WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logarithmBase10WithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the logarithm with base 2 to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logarithmBase2(with:name:)
func (g MPSGraph) LogarithmBase2WithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logarithmBase2WithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical AND of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical AND operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalAND(_:_:name:)
func (g MPSGraph) LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical NAND of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical NAND operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalNAND(_:_:name:)
func (g MPSGraph) LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalNANDWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical NOR of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical NOR operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalNOR(_:_:name:)
func (g MPSGraph) LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalNORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical OR of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical OR operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalOR(_:_:name:)
func (g MPSGraph) LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical XNOR of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical XNOR operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalXNOR(_:_:name:)
func (g MPSGraph) LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalXNORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise logical XOR of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a logical XOR operation and returns the result
// tensor. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/logicalXOR(_:_:name:)
func (g MPSGraph) LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("logicalXORWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the matrix multiplication of 2 input tensors with support for
// broadcasting.
//
// primaryTensor: The left-hand side tensor.
//
// secondaryTensor: The right-hand side tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid tensor containing the product of the input matrices.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/matrixMultiplication(primary:secondary:name:)
func (g MPSGraph) MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("matrixMultiplicationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D max-pooling operation and returns the result tensor.
//
// source: A 2D Image source as tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2D(withSourceTensor:descriptor:name:)
func (g MPSGraph) MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling2DWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// indices: The indices tensor returned from
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName].
//
// outputShape: The shape of the destination gradient.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// Destination gradient tensor.
//
// # Discussion
//
// With this API MPSGraph computes the max-pooling gradient efficiently by
// reusing the indices from the forward API instead of recomputing them. The
// descriptor must set `returnIndicesMode` and `returnIndicesDataType` to the
// same value as that set by the forward pass.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:indicesTensor:outputShape:descriptor:name:)
func (g MPSGraph) MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:indicesTensor:outputShape:descriptor:name:"), gradient, indices, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// indices: The indices tensor returned from
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName].
//
// outputShape: A tensor containing the shape of the destination gradient.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// Destination gradient tensor.
//
// # Discussion
//
// With this API MPSGraph computes the max-pooling gradient efficiently by
// reusing the indices from the forward API instead of recomputing them. The
// descriptor must set `returnIndicesMode` and `returnIndicesDataType` to the
// same value as that set by the forward pass.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:)
func (g MPSGraph) MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:"), gradient, indices, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: A 2D input gradient tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// source: The input tensor for the forward pass.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:sourceTensor:descriptor:name:)
func (g MPSGraph) MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 2D max-pooling operation and returns the result tensor and the
// corresponding indices tensor.
//
// source: A 2D Image source as tensor - must be of rank=4. The layout is defined by
// `descriptor.DataLayout()`.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of two MPSGraphTensors. The first tensor holds the result of max
// pool and the second tensor holds the corresponding indices
//
// # Discussion
//
// In order to Computes the indices, `returnIndicesMode` of the descriptor
// must be set. The datatype of indices tensor can be set using
// `returnIndicesDataType`. If `returnIndicesMode =
// MPSGraphPoolingReturnIndicesNone` then only the first result MPSGraph
// returns will be valid and using the second result will assert.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)
func (g MPSGraph) MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("maxPooling2DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a 4D max-pooling operation and returns the result tensor.
//
// source: A source tensor.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4D(_:descriptor:name:)
func (g MPSGraph) MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: An input gradient tensor.
//
// source: The input tensor for the forward pass.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(_:source:descriptor:name:)
func (g MPSGraph) MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: An input gradient tensor.
//
// indices: Indices tensor returned from
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName].
//
// outputShape: The shape of the destination gradient.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// Destination gradient tensor.
//
// # Discussion
//
// With this API MPSGraph computes the max-pooling gradient efficiently by
// reusing the indices from the forward API instead of recomputing them. The
// descriptor must set `returnIndicesMode` and `returnIndicesDataType` to the
// same value as that set by the forward pass.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(withGradientTensor:indicesTensor:outputShape:descriptor:name:)
func (g MPSGraph) MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape foundation.NSArray, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:indicesTensor:outputShape:descriptor:name:"), gradient, indices, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// gradient: An input gradient tensor.
//
// indices: The indices tensor returned from
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName].
//
// outputShape: A tensor containing the shape of the destination gradient.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates, paddings and layouts.
//
// name: The name for the operation.
//
// # Return Value
//
// Destination gradient tensor.
//
// # Discussion
//
// With this API MPSGraph computes the max-pooling gradient efficiently by
// reusing the indices from the forward API instead of recomputing them. The
// descriptor must set `returnIndicesMode` and `returnIndicesDataType` to the
// same value as that set by the forward pass.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DGradient(withGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:)
func (g MPSGraph) MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName(gradient IMPSGraphTensor, indices IMPSGraphTensor, outputShape IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maxPooling4DGradientWithGradientTensor:indicesTensor:outputShapeTensor:descriptor:name:"), gradient, indices, outputShape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a 4D max-pooling operation and returns the result tensor and the
// corresponding indices tensor.
//
// source: The source tensor on which pooling will be performed.
//
// descriptor: A pooling operation descriptor that specifies pooling window sizes,
// strides, dilation rates and paddings.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of two MPSGraphTensors. The first tensor holds the result of max
// pool and the second tensor holds the corresponding indices.
//
// # Discussion
//
// In order to Computes the indices, `returnIndicesMode` of the descriptor
// must be set. The datatype of indices tensor can be set using
// `returnIndicesDataType`. If `returnIndicesMode =
// MPSGraphPoolingReturnIndicesNone` then only the first result MPSGraph
// returns will be valid and using the second result will assert.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DReturnIndices(_:descriptor:name:)
func (g MPSGraph) MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("maxPooling4DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Returns the elementwise maximum of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a maximum operation and returns the result tensor.
// It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maximum(_:_:name:)
func (g MPSGraph) MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maximumWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise maximum of the input tensors, while propagating
// [NaN] values.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a maximum with [NaN] propagation operation and
// returns the result tensor. This means that if any of the elementwise
// operands is [NaN], the result is [NaN]. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maximumWithNaNPropagation(_:_:name:)
func (g MPSGraph) MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maximumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the mean of the first input along the specified axes.
//
// axes: A list of axes over which to perform the reduction. The order of dimensions
// goes from the slowest moving at axis=0 to the fastest moving dimension.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/mean(of:axes:name:)
func (g MPSGraph) MeanOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("meanOfTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise minimum of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a minimum operation and returns the result tensor.
// It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/minimum(_:_:name:)
func (g MPSGraph) MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("minimumWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise minimum of the input tensors, while propagating
// [NaN] values.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a minimum with [NaN] propagation operation and
// returns the result tensor. This means that if any of the elementwise
// operands is [NaN], the result is [NaN]. It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/minimumWithNaNPropagation(_:_:name:)
func (g MPSGraph) MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("minimumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the remainder obtained by dividing the first input tensor by the
// second.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a modulo operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/modulo(_:_:name:)
func (g MPSGraph) ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("moduloWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Multiplies two input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a multiply operation and returns the result tensor.
// It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/multiplication(_:_:name:)
func (g MPSGraph) MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("multiplicationWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies negative to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/negative(with:name:)
func (g MPSGraph) NegativeWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("negativeWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a nonMaximumumSuppression operation and returns the result tensor.
//
// boxesTensor: A tensor containing the coordinates of the input boxes. Must be a rank 3
// tensor of shape [N,B,4] of type [MPSDataTypeFloat32]
//
// scoresTensor: A tensor containing the scores of the input boxes. Must be a rank 3 tensor
// of shape [N,B,1] of type [MPSDataTypeFloat32]
//
// classIndicesTensor: A tensor containing the class indices of the input boxes. Must be a rank 2
// tensor of shape [N,B] of type [MPSDataTypeInt32]
//
// IOUThreshold: The threshold for when to reject boxes based on their Intersection Over
// Union. Valid range is [0,1].
//
// scoreThreshold: The threshold for when to reject boxes based on their score, before IOU
// suppression.
//
// perClassSuppression: When this is specified a box will only suppress another box if they have
// the same class.
//
// coordinateMode: The coordinate mode the box coordinates are provided in.
//
// name: The name for the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonMaximumSuppression(withBoxesTensor:scoresTensor:classIndicesTensor:iouThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:)
func (g MPSGraph) NonMaximumSuppressionWithBoxesTensorScoresTensorClassIndicesTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, classIndicesTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode MPSGraphNonMaximumSuppressionCoordinateMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nonMaximumSuppressionWithBoxesTensor:scoresTensor:classIndicesTensor:IOUThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:"), boxesTensor, scoresTensor, classIndicesTensor, IOUThreshold, scoreThreshold, perClassSuppression, coordinateMode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a nonMaximumumSuppression operation and returns the result tensor.
//
// boxesTensor: A tensor containing the coordinates of the input boxes. Must be a rank 3
// tensor of shape [N,B,4] of type [MPSDataTypeFloat32]
//
// scoresTensor: A tensor containing the scores of the input boxes. Must be a rank 3 tensor
// of shape [N,B,K] of type [MPSDataTypeFloat32]
//
// IOUThreshold: The threshold for when to reject boxes based on their Intersection Over
// Union. Valid range is [0,1].
//
// scoreThreshold: The threshold for when to reject boxes based on their score, before IOU
// suppression.
//
// perClassSuppression: When this is specified a box will only suppress another box if they have
// the same class.
//
// coordinateMode: The coordinate mode the box coordinates are provided in.
//
// name: The name for the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonMaximumSuppression(withBoxesTensor:scoresTensor:iouThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:)
func (g MPSGraph) NonMaximumSuppressionWithBoxesTensorScoresTensorIOUThresholdScoreThresholdPerClassSuppressionCoordinateModeName(boxesTensor IMPSGraphTensor, scoresTensor IMPSGraphTensor, IOUThreshold float32, scoreThreshold float32, perClassSuppression bool, coordinateMode MPSGraphNonMaximumSuppressionCoordinateMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nonMaximumSuppressionWithBoxesTensor:scoresTensor:IOUThreshold:scoreThreshold:perClassSuppression:coordinateMode:name:"), boxesTensor, scoresTensor, IOUThreshold, scoreThreshold, perClassSuppression, coordinateMode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the indices of the non-zero elements of the input tensor.
//
// tensor: An MPSGraphTensor of which to compute the non-zero indices.
//
// # Return Value
//
// A valid MPSGraphTensor containing indices in signed int32 data type.
//
// # Discussion
//
// The indices are returned as a two-dimensional tensor of size
// `[number_of_nonzeros, input_rank]`. Each row in the result contains indices
// of a nonzero elements in input. For example:
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/nonZeroIndices(_:name:)
func (g MPSGraph) NonZeroIndicesOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nonZeroIndicesOfTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a normalization beta-gradient operation and returns the result
// tensor.
//
// incomingGradientTensor: The incoming original `resultTensor` gradient.
//
// sourceTensor: The original input source in forward direction.
//
// axes: The axes of normalization.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// The mean and variance tensors should be outputs of `name` and `name`. Use
// the axes parameter to achieve different types of normalizations. For
// example (assuming your data is in [NxHxWxC] format) Batch normalization:
// axes = [0, 1, 2] Instance normalization: axes = [1, 2]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationBetaGradient(withIncomingGradientTensor:sourceTensor:reductionAxes:name:)
func (g MPSGraph) NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("normalizationBetaGradientWithIncomingGradientTensor:sourceTensor:reductionAxes:name:"), incomingGradientTensor, sourceTensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a normalization gamma-gradient operation and returns the result
// tensor.
//
// incomingGradientTensor: The incoming original `resultTensor` gradient.
//
// sourceTensor: The original input source in forward direction.
//
// meanTensor: The mean tensor.
//
// varianceTensor: The variance tensor.
//
// axes: The axes of normalization.
//
// epsilon: A small value to add to the variance when normalizing the inputs.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// The mean and variance tensors should be outputs of `name` and `name`. Use
// the axes parameter to achieve different types of normalizations. For
// example (assuming your data is in [NxHxWxC] format) Batch normalization:
// axes = [0, 1, 2] Instance normalization: axes = [1, 2]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationGammaGradient(withIncomingGradientTensor:sourceTensor:mean:varianceTensor:reductionAxes:epsilon:name:)
func (g MPSGraph) NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, axes []foundation.NSNumber, epsilon float32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("normalizationGammaGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:reductionAxes:epsilon:name:"), incomingGradientTensor, sourceTensor, meanTensor, varianceTensor, objectivec.IObjectSliceToNSArray(axes), epsilon, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a normalization input gradient operation and returns the result
// tensor.
//
// incomingGradientTensor: The incoming original `resultTensor` gradient.
//
// sourceTensor: The original input source in forward direction.
//
// meanTensor: The mean tensor.
//
// varianceTensor: The variance tensor.
//
// gamma: The gamma tensor.
//
// gammaGradient: The `gammaGradient` tensor.
//
// betaGradient: The `betaGradient` tensor
//
// axes: The axes of normalization.
//
// epsilon: A small value to add to the variance when normalizing the inputs.
//
// name: An optional name for the operation.
//
// # Discussion
//
// The mean and variance tensors should be outputs of `name` and `name`. Use
// the axes parameter to achieve different types of normalizations. For
// example (assuming your data is in [NxHxWxC] format) Batch normalization:
// axes = [0, 1, 2] Instance normalization: axes = [1, 2]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalizationGradient(withIncomingGradientTensor:sourceTensor:mean:varianceTensor:gammaTensor:gammaGradientTensor:betaGradientTensor:reductionAxes:epsilon:name:)
func (g MPSGraph) NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, meanTensor IMPSGraphTensor, varianceTensor IMPSGraphTensor, gamma IMPSGraphTensor, gammaGradient IMPSGraphTensor, betaGradient IMPSGraphTensor, axes []foundation.NSNumber, epsilon float32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("normalizationGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:gammaTensor:gammaGradientTensor:betaGradientTensor:reductionAxes:epsilon:name:"), incomingGradientTensor, sourceTensor, meanTensor, varianceTensor, gamma, gammaGradient, betaGradient, objectivec.IObjectSliceToNSArray(axes), epsilon, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a batch normalization operation and returns the result tensor.
//
// tensor: The input tensor.
//
// mean: The mean tensor.
//
// variance: The variance tensor.
//
// gamma: The tensor used to scale the normalized result.
//
// beta: The tensor used to bias the normalized result.
//
// epsilon: A small value to add to the variance when normalizing the inputs.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// The mean and variance tensors should be outputs of `name` and `name`. Use
// the axes parameter to achieve different types of normalizations. For
// example (assuming your data is in NxHxWxC format) Batch normalization: axes
// = [0, 1, 2] Instance normalization: axes = [1, 2] Shapes for gamma and beta
// must match the input data along at least one dimension and will be
// broadcast along the rest. For batch normalization, gamma and beta would
// typically be 1x1x1xC i.e. one value per channel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/normalize(_:mean:variance:gamma:beta:epsilon:name:)
func (g MPSGraph) NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor IMPSGraphTensor, mean IMPSGraphTensor, variance IMPSGraphTensor, gamma IMPSGraphTensor, beta IMPSGraphTensor, epsilon float32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("normalizationWithTensor:meanTensor:varianceTensor:gammaTensor:betaTensor:epsilon:name:"), tensor, mean, variance, gamma, beta, epsilon, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the logical NOT operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/not(with:name:)
func (g MPSGraph) NotWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("notWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise inequality check of the input tensors.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a not equal operation and returns the result tensor.
// It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/notEqual(_:_:name:)
func (g MPSGraph) NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("notEqualWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// axis: The axis to insert the new oneHot vector at
//
// dataType: MPSDataType of the result tensor.
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the rank of `indicesTensor` + 1. Inserts
// a new axis at the axis specified, or the minor axis if `axis` is -1. The
// values at the indices in the indicesTensor will be set to 1, and all other
// values will be set to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:dataType:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthAxisDataTypeName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:dataType:name:"), indicesTensor, depth, axis, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// axis: The axis to insert the new oneHot vector at. Defaults to -1, the minor axis
//
// dataType: MPSDataType of the result tensor Defaults to MPSDataTypeFloat
//
// onValue: The value for indices designated by the indicesTensor. This value must
// match the specified data type. Defaults to 1.0f
//
// offValue: The value for indices not designated by the indicesTensor. This value must
// match the specified data type. Defaults to 0.0f
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the indicesTensor rank + 1. Inserts a new
// axis at the axis specified, or the minor axis if axis is -1. The values at
// the indices in the indicesTensor will have the onValue, and all other
// values will be set to the offValue.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:dataType:onValue:offValue:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthAxisDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, axis uint, dataType uint32, onValue float64, offValue float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:dataType:onValue:offValue:name:"), indicesTensor, depth, axis, dataType, onValue, offValue, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// axis: The axis to insert the new oneHot vector at
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the rank of `indicesTensor` + 1, of type
// MPSDataTypeFloat32. Inserts a new axis at the axis specified, or the minor
// axis if `axis` is -1. The values at the indices in the indicesTensor will
// be set to 1, and all other values will be set to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:axis:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthAxisName(indicesTensor IMPSGraphTensor, depth uint, axis uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:axis:name:"), indicesTensor, depth, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// dataType: MPSDataType of the result tensor.
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the rank of `indicesTensor` + 1. Inserts
// a new axis at the minor dimension. The values at the indices in the
// indicesTensor will be set to 1, and all other values will be set to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:dataType:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthDataTypeName(indicesTensor IMPSGraphTensor, depth uint, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:dataType:name:"), indicesTensor, depth, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// dataType: MPSDataType of the result tensor.
//
// onValue: The value for indices designated by the indicesTensor. This value must
// match the specified data type.
//
// offValue: The value for indices not designated by the indicesTensor. This value must
// match the specified data type.
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the rank of `indicesTensor` + 1. Inserts
// a new axis at the minor dimension. The values at the indices in the
// indicesTensor will have the onValue, and all other values will be set to
// the offValue.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:dataType:onValue:offValue:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthDataTypeOnValueOffValueName(indicesTensor IMPSGraphTensor, depth uint, dataType uint32, onValue float64, offValue float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:dataType:onValue:offValue:name:"), indicesTensor, depth, dataType, onValue, offValue, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a oneHot operation and returns the result tensor.
//
// indicesTensor: Tensor of indices for on values
//
// depth: Depth of the oneHot vector along the axis
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor of rank equal to the rank of `indicesTensor` + 1, of type
// MPSDataTypeFloat32. Inserts a new axis at the minor dimension. The values
// at the indices in the indicesTensor will be set to 1, and all other values
// will be set to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/oneHot(withIndicesTensor:depth:name:)
func (g MPSGraph) OneHotWithIndicesTensorDepthName(indicesTensor IMPSGraphTensor, depth uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("oneHotWithIndicesTensor:depth:name:"), indicesTensor, depth, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a padding gradient operation and returns the result tensor.
//
// incomingGradientTensor: The input gradient tensor.
//
// sourceTensor: The input tensor of the forward pass.
//
// paddingMode: The parameter that defines the padding mode.
//
// leftPadding: The parameter that defines how much padding the operation applies to the
// input tensor before each dimension - must be of size `rank(tensor)`.
//
// rightPadding: The parameter that defines how much padding the operation applies to the
// input tensor after each dimension - must be of size `rank(tensor)`.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/padGradient(withIncomingGradientTensor:sourceTensor:paddingMode:leftPadding:rightPadding:name:)
func (g MPSGraph) PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, paddingMode MPSGraphPaddingMode, leftPadding foundation.NSArray, rightPadding foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("padGradientWithIncomingGradientTensor:sourceTensor:paddingMode:leftPadding:rightPadding:name:"), incomingGradientTensor, sourceTensor, paddingMode, leftPadding, rightPadding, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a padding operation and returns the result tensor.
//
// tensor: The input tensor.
//
// paddingMode: The parameter that defines the padding mode.
//
// leftPadding: The parameter that defines how much padding the operation applies to the
// input tensor before each dimension - must be of size `rank(tensor)`.
//
// rightPadding: The parameter that defines how much padding the operation applies to the
// input tensor after each dimension - must be of size `rank(tensor)`.
//
// constantValue: The constant value the operation uses when `paddingMode =
// MPSGraphPaddingModeConstant`.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/padTensor(_:with:leftPadding:rightPadding:constantValue:name:)
func (g MPSGraph) PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor IMPSGraphTensor, paddingMode MPSGraphPaddingMode, leftPadding foundation.NSArray, rightPadding foundation.NSArray, constantValue float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("padTensor:withPaddingMode:leftPadding:rightPadding:constantValue:name:"), tensor, paddingMode, leftPadding, rightPadding, constantValue, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a placeholder operation and returns the result tensor.
//
// shape: The shape of the output tensor. A nil shape will result in an unranked
// tensor.
//
// dataType: The dataType of the placeholder tensor.
//
// name: The name for the placeholder operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholder(shape:dataType:name:)
func (g MPSGraph) PlaceholderWithShapeDataTypeName(shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("placeholderWithShape:dataType:name:"), shape, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a placeholder operation and returns the result tensor with the
// dataType of the placeholder tensor set to 32 bit float.
//
// shape: The shape of the output tensor. A nil shape will result in an unranked
// tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholder(shape:name:)
func (g MPSGraph) PlaceholderWithShapeName(shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("placeholderWithShape:name:"), shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the elementwise result of raising the first tensor to the power of
// the second tensor.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a power operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/power(_:_:name:)
func (g MPSGraph) PowerWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("powerWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Quantize operation and returns the result tensor.
//
// tensor: Input tensor to be quantized
//
// scale: Scale scalar parameter
//
// zeroPoint: Bias scalar parameter (converted to dataType of resultTensor)
//
// dataType: Integer data type of the result tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the float `tensor` to an i8 or u8 tensor by applying a scale + bias
// transform: result = (tensor / scale) + zeroPoint
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scale:zeroPoint:dataType:name:)
func (g MPSGraph) QuantizeTensorScaleZeroPointDataTypeName(tensor IMPSGraphTensor, scale float64, zeroPoint float64, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("quantizeTensor:scale:zeroPoint:dataType:name:"), tensor, scale, zeroPoint, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Quantize operation and returns the result tensor.
//
// tensor: Input tensor to be quantized
//
// scaleTensor: Scale 1D Tensor parameter with size == tensor.shape[axis]
//
// zeroPoint: Bias scalar parameter (converted to dataType of resultTensor)
//
// dataType: Integer data type of the result tensor.
//
// axis: Axis on which the scale 1D value is being broadcasted
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the float `tensor` to an i8 or u8 tensor by applying a scale + bias
// transform: result = (tensor / scaleTensor) + zeroPoint
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scaleTensor:zeroPoint:dataType:axis:name:)
func (g MPSGraph) QuantizeTensorScaleTensorZeroPointDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPoint float64, dataType uint32, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("quantizeTensor:scaleTensor:zeroPoint:dataType:axis:name:"), tensor, scaleTensor, zeroPoint, dataType, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Quantize operation and returns the result tensor.
//
// tensor: Input tensor to be quantized
//
// scaleTensor: Scale scalar or 1D Tensor parameter with size == tensor.shape[axis]
//
// zeroPointTensor: Bias scalar or 1D Tensor parameter with size == tensor.shape[axis]
//
// dataType: Integer data type of the result tensor.
//
// axis: Axis on which the scale 1D value is being broadcasted
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of datatype dataType
//
// # Discussion
//
// Convert the float `tensor` to an i8 or u8 tensor by applying a scale + bias
// transform: result = (tensor / scaleTensor) + zeroPointTensor
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/quantize(_:scaleTensor:zeroPointTensor:dataType:axis:name:)
func (g MPSGraph) QuantizeTensorScaleTensorZeroPointTensorDataTypeAxisName(tensor IMPSGraphTensor, scaleTensor IMPSGraphTensor, zeroPointTensor IMPSGraphTensor, dataType uint32, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("quantizeTensor:scaleTensor:zeroPointTensor:dataType:axis:name:"), tensor, scaleTensor, zeroPointTensor, dataType, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a tensor representing state using the Philox algorithm with given
// counter and key values.
//
// counterLow: The value to initilaize lower 64 bits of counter to. Philox utilizes a 128
// bit counter
//
// counterHigh: The value to initilaize upper 64 bits of counter to. Philox utilizes a 128
// bit counter
//
// key: The value to initialize the key to in Philox algorithm.
//
// name: Name for the operation
//
// # Return Value
//
// An MPSGraphTensor representing a random state, to be passed as an input to
// a random op.
//
// # Discussion
//
// See randomPhiloxStateTensorWithSeed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomPhiloxStateTensor(withCounterLow:counterHigh:key:name:)
func (g MPSGraph) RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomPhiloxStateTensorWithCounterLow:counterHigh:key:name:"), counterLow, counterHigh, key, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a tensor representing state using the Philox algorithm with given
// counter and key values.
//
// # Discussion
//
// Generates random numbers using the Philox counter-based algorithm, for
// further details see: John K. Salmon, Mark A. Moraes, Ron O. Dror, and David
// E. Shaw. Parallel Random Numbers: As Easy as 1, 2, 3. A stateTensor
// generated with this API can be used in MPSGraph Random APIs which accept a
// stateTensor. The updated stateTensor is returned alongside the random
// values, and can be fed to the following random layer. In most use cases, a
// stateTensor should only need to be initialized once at the start of the
// graph. A stateTensor can be set as a target tensor of an MPSGraph execution
// to obtain a stateTensor serialized as an NDArray. This can be used as input
// to a placeholder in the graph to avoid ever needing to have a state
// intilization layer in an MPSGraph. This can allow for a continued stream
// through multiple executions of a single MPSGraph by having the final
// stateTensor as a target tensor passed into the following MPSGraph execution
// as a placeholder input. This may be helpful for training graphs in
// particular.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomPhiloxStateTensor(withSeed:name:)
func (g MPSGraph) RandomPhiloxStateTensorWithSeedName(seed uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomPhiloxStateTensorWithSeed:name:"), seed, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Random op of type matching distribution in descriptor and returns
// random values.
//
// shape: The shape of the tensor generated
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random values in the distribution
// specified. Uses a random seed value to initalize state. No state is
// preserved, and subsequent calls are not guaranteed to result in a unique
// stream of random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:name:)
func (g MPSGraph) RandomTensorWithShapeDescriptorName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomTensorWithShape:descriptor:name:"), shape, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Random op of type matching distribution in descriptor and returns
// random values.
//
// shape: The shape of the tensor generated
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// seed: The seed to use to initialize state. All calls with equal seed yield an
// identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random values in the distribution
// specified. Uses the provided seed value to initalize state. No state is
// preserved, and all calls with equal seed yield an identical stream of
// random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:seed:name:)
func (g MPSGraph) RandomTensorWithShapeDescriptorSeedName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, seed uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomTensorWithShape:descriptor:seed:name:"), shape, descriptor, seed, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Random op of type matching distribution in descriptor, and
// returns random values and updated state.
//
// shape: The shape of the tensor generated
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// state: The state to define a stream of random values. All calls with equal state
// yield an identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of MPSGraphTensor of size 2. The first MPSGraphTensor is of shape
// containing random values in the defined range. The second MPSGraphTensor is
// the updated state tensor.
//
// # Discussion
//
// Returns an array of 2 tensors, where the first is of provided shape of
// random values in the distribution specified, and the second is the updated
// state tensor. Uses the provided state to define a stream of random values.
// No state is preserved, and all calls with equal state yield an identical
// stream of random values. The initial stateTensor provided should be created
// using the MPSGraph randomPhiloxStateTensor APIs. The resulting stateTensor
// from this op can be passed as an argument to the following random calls to
// continue sampling from the stream.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShape:descriptor:stateTensor:name:)
func (g MPSGraph) RandomTensorWithShapeDescriptorStateTensorName(shape foundation.NSArray, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("randomTensorWithShape:descriptor:stateTensor:name:"), shape, descriptor, state, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a Random op of type matching distribution in descriptor and returns
// random values.
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random values in the distribution
// specified. Uses a random seed value to initalize state. No state is
// preserved, and subsequent calls are not guaranteed to result in a unique
// stream of random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:name:)
func (g MPSGraph) RandomTensorWithShapeTensorDescriptorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:name:"), shapeTensor, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Random op of type matching distribution in descriptor and returns
// random values.
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// seed: The seed to use to initialize state. All calls with equal seed yield an
// identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random values in the distribution
// specified. Uses the provided seed value to initalize state. No state is
// preserved, and all calls with equal seed yield an identical stream of
// random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:seed:name:)
func (g MPSGraph) RandomTensorWithShapeTensorDescriptorSeedName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, seed uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:seed:name:"), shapeTensor, descriptor, seed, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Random op of type matching distribution in descriptor, and
// returns random values and updated state.
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated.
//
// descriptor: The descriptor of the distribution. See MPSGraphRandomOpDescriptor.
//
// state: The state to define a stream of random values. All calls with equal state
// yield an identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of MPSGraphTensor of size 2. The first MPSGraphTensor is of shape
// containing random values in the defined range. The second MPSGraphTensor is
// the updated state tensor.
//
// # Discussion
//
// Returns an array of 2 tensors, where the first is of provided shape of
// random values in the distribution specified, and the second is the updated
// state tensor. Uses the provided state to define a stream of random values.
// No state is preserved, and all calls with equal state yield an identical
// stream of random values. The initial stateTensor provided should be created
// using the MPSGraph randomPhiloxStateTensor APIs. The resulting stateTensor
// from this op can be passed as an argument to the following random calls to
// continue sampling from the stream.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomTensor(withShapeTensor:descriptor:stateTensor:name:)
func (g MPSGraph) RandomTensorWithShapeTensorDescriptorStateTensorName(shapeTensor IMPSGraphTensor, descriptor IMPSGraphRandomOpDescriptor, state IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("randomTensorWithShapeTensor:descriptor:stateTensor:name:"), shapeTensor, descriptor, state, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a RandomUniform operation and returns random uniform values
//
// shape: The shape of the tensor generated
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random uniform values in the range
// [0.0, 1.0). Uses a random seed value to initalize state. No state is
// preserved, and subsequent calls are not guaranteed to result in a unique
// stream of random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:name:)
func (g MPSGraph) RandomUniformTensorWithShapeName(shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomUniformTensorWithShape:name:"), shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a RandomUniform operation and returns random uniform values
//
// shape: The shape of the tensor generated
//
// seed: The seed to use to initialize state. All calls with equal seed yield an
// identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random uniform values in the range
// [0.0, 1.0). Uses the provided seed value to initalize state. No state is
// preserved, and all calls with equal seed yield an identical stream of
// random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:seed:name:)
func (g MPSGraph) RandomUniformTensorWithShapeSeedName(shape foundation.NSArray, seed uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomUniformTensorWithShape:seed:name:"), shape, seed, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a RandomUniform operation and returns random uniform values and
// updated state
//
// shape: The shape of the tensor generated
//
// state: The state to define a stream of random values. All calls with equal state
// yield an identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of MPSGraphTensor of size 2. The first MPSGraphTensor is of shape
// containing random values in the defined range. The second MPSGraphTensor is
// the updated state tensor.
//
// # Discussion
//
// Returns an array of 2 tensors, where the first is a tensor of provided
// shape of random uniform values in the range [0.0, 1.0), and the second is
// the updated state tensor. The provided state is used to define a stream of
// random values. No state is preserved, and all calls with equal state yield
// an identical stream of random values. The initial stateTensor provided
// should be created using the MPSGraph randomPhiloxStateTensor APIs. The
// resulting stateTensor from this op can be passed as an argument to the
// following random calls to continue sampling from the stream.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShape:stateTensor:name:)
func (g MPSGraph) RandomUniformTensorWithShapeStateTensorName(shape foundation.NSArray, state IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("randomUniformTensorWithShape:stateTensor:name:"), shape, state, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a RandomUniform operation and returns random uniform values
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random uniform values in the range
// [0.0, 1.0). Uses a random seed value to initalize state. No state is
// preserved, and subsequent calls are not guaranteed to result in a unique
// stream of random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:name:)
func (g MPSGraph) RandomUniformTensorWithShapeTensorName(shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomUniformTensorWithShapeTensor:name:"), shapeTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a RandomUniform operation and returns random uniform values
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated
//
// seed: The seed to use to initialize state. All calls with equal seed yield an
// identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An MPSGraphTensor of shape containing random values in the defined range.
//
// # Discussion
//
// Returns a tensor of provided shape of random uniform values in the range
// [0.0, 1.0). Uses the provided seed value to initalize state. No state is
// preserved, and all calls with equal seed yield an identical stream of
// random values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:seed:name:)
func (g MPSGraph) RandomUniformTensorWithShapeTensorSeedName(shapeTensor IMPSGraphTensor, seed uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("randomUniformTensorWithShapeTensor:seed:name:"), shapeTensor, seed, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a RandomUniform operation and returns random uniform values and
// updated state
//
// shapeTensor: 1D Int32 or Int64 tensor. The shape of the tensor generated
//
// state: The state to define a stream of random values. All calls with equal state
// yield an identical stream of random values.
//
// name: The name for the operation.
//
// # Return Value
//
// An array of MPSGraphTensor of size 2. The first MPSGraphTensor is of shape
// containing random values in the defined range. The second MPSGraphTensor is
// the updated state tensor.
//
// # Discussion
//
// Returns an array of 2 tensors, where the first is a tensor of provided
// shape of random uniform values in the range [0.0, 1.0), and the second is
// the updated state tensor. The provided state is used to define a stream of
// random values. No state is preserved, and all calls with equal state yield
// an identical stream of random values. The initial stateTensor provided
// should be created using the MPSGraph randomPhiloxStateTensor APIs. The
// resulting stateTensor from this op can be passed as an argument to the
// following random calls to continue sampling from the stream.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/randomUniformTensor(withShapeTensor:stateTensor:name:)
func (g MPSGraph) RandomUniformTensorWithShapeTensorStateTensorName(shapeTensor IMPSGraphTensor, state IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("randomUniformTensorWithShapeTensor:stateTensor:name:"), shapeTensor, state, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Computes the ReLU (rectified linear activation unit) function with the
// input tensor.
//
// tensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// The operation is: f(x) = max(x, 0).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reLU(with:name:)
func (g MPSGraph) ReLUWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reLUWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the gradient of the ReLU (rectified linear activation unit)
// function using the incoming gradient.
//
// gradient: The incoming gradient tensor.
//
// source: The input tensor from forward pass.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reLUGradient(withIncomingGradient:sourceTensor:name:)
func (g MPSGraph) ReLUGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reLUGradientWithIncomingGradient:sourceTensor:name:"), gradient, source, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a read op which reads at this point of execution of the graph and
// returns the result tensor.
//
// variable: The variable resource tensor to read from.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/read(_:name:)
func (g MPSGraph) ReadVariableName(variable IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("readVariable:name:"), variable, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the real part of a tensor.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation..
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realPartOfTensor(tensor:name:)
func (g MPSGraph) RealPartOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("realPartOfTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Real-to-Hermitean fast Fourier transform operation and returns
// the result tensor.
//
// tensor: A Real-valued input tensor. Must have datatype [MPSDataTypeFloat32] or
// [MPSDatatypeFloat16].
//
// axes: An array of numbers that specifies over which axes MPSGraph performs the
// Fourier transform - all axes must be contained within last four dimensions
// of the input tensor.
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor of type [MPSDataTypeComplexFloat32] or
// [MPSDataTypeComplexFloat16] with reduced size (see Discussion).
//
// # Discussion
//
// This operation computes the fast Fourier transform of a real-valued input
// tensor according to the following formulae.
//
// `scale = 1` for `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode =
// size`, `scale = 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is
// the volume of the transformation defined by the dimensions included in
// `axes` (`V_f = prod_{i \in axes} shape(input)[i]`) (see
// [MPSGraphFFTDescriptor.ScalingMode]), `+` is selected in `+/-` when
// `inverse` is specified, otherwise `-` is used and the sum is done
// separately over each dimension in `axes` and `n` is the dimension length of
// that axis. With this API MPSGraph writes out only the results for the
// unique frequencies, resulting in a tensor which has size `(n/2)+1` in the
// last dimension defined by `axes`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realToHermiteanFFT(_:axes:descriptor:name:)
func (g MPSGraph) RealToHermiteanFFTWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.NSNumber, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("realToHermiteanFFTWithTensor:axes:descriptor:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Real-to-Hermitean fast Fourier transform operation and returns
// the result tensor.
//
// tensor: A real-valued input tensor. Must have datatype [MPSDataTypeFloat32] or
// [MPSDatatypeFloat16].
//
// axesTensor: A tensor of rank one containing the axes over which MPSGraph performs the
// transformation. See
// [MPSGraph.FastFourierTransformWithTensorAxesDescriptorName].
//
// descriptor: A descriptor that defines the parameters of the Fourier transform operation
// - see [MPSGraphFFTDescriptor].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor of type [MPSDataTypeComplexFloat32] or
// [MPSDataTypeComplexFloat16] with reduced size (see Discussion).
//
// # Discussion
//
// This operation computes the fast Fourier transform of a real-valued input
// tensor according to the following formulae.
//
// `scale = 1` for `scaling_mode = none`, `scale = 1/V_f` for `scaling_mode =
// size`, `scale = 1/sqrt(V_f)` for `scaling_mode = unitary`, where `V_f` is
// the volume of the transformation defined by the dimensions included in
// `axes` (`V_f = prod_{i \in axes} shape(input)[i]`) (see
// [MPSGraphFFTDescriptor.ScalingMode]), `+` is selected in `+/-` when
// `inverse` is specified, otherwise `-` is used and the sum is done
// separately over each dimension in `axes` and `n` is the dimension length of
// that axis. With this API MPSGraph writes out only the results for the
// unique frequencies, resulting in a tensor which has size `(n/2)+1` in the
// last dimension defined by `axes`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realToHermiteanFFT(_:axesTensor:descriptor:name:)
func (g MPSGraph) RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("realToHermiteanFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the reciprocal operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reciprocal(with:name:)
func (g MPSGraph) ReciprocalWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reciprocalWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the reciprocal square root operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reciprocalSquareRoot(_:name:)
func (g MPSGraph) ReciprocalSquareRootWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reciprocalSquareRootWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction and operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionAnd(with:axes:name:)
func (g MPSGraph) ReductionAndWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionAndWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction and operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionAnd(with:axis:name:)
func (g MPSGraph) ReductionAndWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionAndWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction argMax operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionArgMaximum(with:axis:name:)
func (g MPSGraph) ReductionArgMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionArgMaximumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction argMin operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionArgMinimum(with:axis:name:)
func (g MPSGraph) ReductionArgMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionArgMinimumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction max operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximum(with:axes:name:)
func (g MPSGraph) ReductionMaximumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMaximumWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction max operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximum(with:axis:name:)
func (g MPSGraph) ReductionMaximumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMaximumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction max propagate NaN operation and returns the result
// tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximumPropagateNaN(with:axes:name:)
func (g MPSGraph) ReductionMaximumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMaximumPropagateNaNWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction max propagate NaN operation and returns the result
// tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMaximumPropagateNaN(with:axis:name:)
func (g MPSGraph) ReductionMaximumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMaximumPropagateNaNWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction min operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimum(with:axes:name:)
func (g MPSGraph) ReductionMinimumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMinimumWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction minimum operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimum(with:axis:name:)
func (g MPSGraph) ReductionMinimumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMinimumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction min propagate NaN operation and returns the result
// tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimumPropagateNaN(with:axes:name:)
func (g MPSGraph) ReductionMinimumPropagateNaNWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMinimumPropagateNaNWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction min propagate NaN operation and returns the result
// tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionMinimumPropagateNaN(with:axis:name:)
func (g MPSGraph) ReductionMinimumPropagateNaNWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionMinimumPropagateNaNWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction or operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionOr(with:axes:name:)
func (g MPSGraph) ReductionOrWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionOrWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction or operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionOr(with:axis:name:)
func (g MPSGraph) ReductionOrWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionOrWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction product operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionProduct(with:axes:name:)
func (g MPSGraph) ReductionProductWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionProductWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction product operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionProduct(with:axis:name:)
func (g MPSGraph) ReductionProductWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionProductWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction sum operation and returns the result tensor.
//
// tensor: Input tensor
//
// axes: Axes of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionSum(with:axes:name:)
func (g MPSGraph) ReductionSumWithTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionSumWithTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reduction sum operation and returns the result tensor.
//
// tensor: Input tensor
//
// axis: Axis of reduction
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reductionSum(with:axis:name:)
func (g MPSGraph) ReductionSumWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reductionSumWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reinterpret cast operation and returns the result tensor.
//
// tensor: The input tensor.
//
// type: The element type of the returned tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Returns input tensor (with element type `tensor_type`) reinterpreted to
// element type passed in with the last dimension scaled by
// `sizeof(tensor_type) / sizeof(type)`. This operation is endianness agnostic
// and MPSGraph reinterprets the data with the endianness of the system.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reinterpretCast(_:to:name:)
func (g MPSGraph) ReinterpretCastTensorToTypeName(tensor IMPSGraphTensor, type_ uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reinterpretCastTensor:toType:name:"), tensor, type_, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reshape operation and returns the result tensor.
//
// tensor: The tensor to be reshaped.
//
// shape: The result tensor shape.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation reshapes the input tensor to the target shape. The shape
// must be compatible with the input tensor shape, specifically the volume of
// the input tensor has to match the volume defined by the shape. The shape is
// allowed to contain dynamic dimensions (-1) when the result type can be
// inferred unambiguously.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reshape(_:shape:name:)
func (g MPSGraph) ReshapeTensorWithShapeName(tensor IMPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reshapeTensor:withShape:name:"), tensor, shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reshape operation and returns the result tensor.
//
// tensor: The tensor to be reshaped.
//
// shapeTensor: A 1D tensor of type [MPSDataTypeInt32] or [MPSDataTypeInt64], that contains
// the target shape values.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation reshapes the input tensor to the target shape. The shape
// tensor must be compatible with the input tensor shape, specifically the
// volume of the input tensor has to match the volume defined by the shape
// tensor. The shape tensor is allowed to contain dynamic dimensions (-1) when
// the result type can be inferred unambiguously.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reshape(_:shapeTensor:name:)
func (g MPSGraph) ReshapeTensorWithShapeTensorName(tensor IMPSGraphTensor, shapeTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reshapeTensor:withShapeTensor:name:"), tensor, shapeTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: A 2-element shape as [newHeight, newWidth]
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size. Result images will be distorted if
// size is of different aspect ratio. Resize supports the following modes:
// Nearest Neighbor - values are interpolated using the closest neighbor pixel
// Bilinear - values are computed using bilinear interpolation of 4
// neighboring pixels Destination indices are computed using direct index
// scaling by default, with no offset added. If the centerResult parameter is
// true, the destination indices will be scaled and shifted to be centered on
// the input image. If the alignCorners parameter is true, the corners of the
// result images will match the input images. Scaling will be modified to a
// factor of (size - 1) / (inputSize - 1). When alignCorners is true, the
// centerResult parameter does nothing. In order to achieve the same behavior
// as OpenCV’s resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:size:mode:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeTensorSizeModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size foundation.NSArray, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeTensor:size:mode:centerResult:alignCorners:layout:name:"), imagesTensor, size, mode, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size. Result images will be distorted if
// size is of different aspect ratio. Resize supports the following modes:
// Nearest Neighbor - values are interpolated using the closest neighbor pixel
// Bilinear - values are computed using bilinear interpolation of 4
// neighboring pixels Destination indices are computed using direct index
// scaling by default, with no offset added. If the centerResult parameter is
// true, the destination indices will be scaled and shifted to be centered on
// the input image. If the alignCorners parameter is true, the corners of the
// result images will match the input images. Scaling will be modified to a
// factor of (size - 1) / (inputSize - 1). When alignCorners is true, the
// centerResult parameter does nothing. In order to achieve the same behavior
// as OpenCV’s resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:mode:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeTensor:sizeTensor:mode:centerResult:alignCorners:layout:name:"), imagesTensor, size, mode, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size. Result images will be distorted if
// size is of different aspect ratio. Resize supports the following modes:
// Nearest Neighbor - values are interpolated using the closest neighbor pixel
// Bilinear - values are computed using bilinear interpolation of 4
// neighboring pixels Destination indices are computed using direct index
// scaling by default, with no offset added. If the centerResult parameter is
// true, the destination indices will be scaled and shifted to be centered on
// the input image. If the alignCorners parameter is true, the corners of the
// result images will match the input images. Scaling will be modified to a
// factor of (size - 1) / (inputSize - 1). When alignCorners is true, the
// centerResult parameter does nothing. In order to achieve the same behavior
// as OpenCV’s resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:mode:centerResult:alignCorners:name:)
func (g MPSGraph) ResizeTensorSizeTensorModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeTensor:sizeTensor:mode:centerResult:alignCorners:name:"), imagesTensor, size, mode, centerResult, alignCorners, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Resamples input images to given size using the provided scale and offset.
// Destination indices are computed using
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// For most use cases passing the scale and offset directly is unnecessary,
// and it is preferable to use the API specifying centerResult and
// alignCorners.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:scaleOffsetTensor:mode:layout:name:)
func (g MPSGraph) ResizeTensorSizeTensorScaleOffsetTensorModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode MPSGraphResizeMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeTensor:sizeTensor:scaleOffsetTensor:mode:layout:name:"), imagesTensor, size, scaleOffset, mode, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using the provided scale and offset.
// Destination indices are computed using
//
// For most use cases passing the scale and offset directly is unnecessary,
// and it is preferable to use the API specifying centerResult and
// alignCorners.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(_:sizeTensor:scaleTensor:offsetTenor:mode:name:)
func (g MPSGraph) ResizeTensorSizeTensorScaleTensorOffsetTensorModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode MPSGraphResizeMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeTensor:sizeTensor:scaleTensor:offsetTensor:mode:name:"), imagesTensor, size, scale, offset, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:mode:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, mode MPSGraphResizeMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeWithGradientTensor:input:mode:centerResult:alignCorners:layout:name:"), gradient, input, mode, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:scale:offsetTensor:mode:name:)
func (g MPSGraph) ResizeWithGradientTensorInputScaleTensorOffsetTensorModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, mode MPSGraphResizeMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeWithGradientTensor:input:scaleTensor:offsetTensor:mode:name:"), gradient, input, scale, offset, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// mode: The resampling mode to use. If nearest sampling is specifed,
// RoundPreferCeil mode will be used.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resize(withGradientTensor:input:scaleOffsetTensor:mode:layout:name:)
func (g MPSGraph) ResizeWithGradientTensorInputScaleOffsetTensorModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, mode MPSGraphResizeMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeWithGradientTensor:input:scaleOffsetTensor:mode:layout:name:"), gradient, input, scaleOffset, mode, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Resamples input images to given size using bilinear sampling.
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using nearest neighbor sampling.
// Result images will be distorted if size is of different aspect ratio.
// Destination indices are computed using direct index scaling by default,
// with no offset added. If the centerResult parameter is true, the
// destination indices will be scaled and shifted to be centered on the input
// image. If the alignCorners parameter is true, the corners of the result
// images will match the input images. Scaling will be modified to a factor of
// (size - 1) / (inputSize - 1). When alignCorners is true, the centerResult
// parameter does nothing. In order to achieve the same behavior as OpenCV’s
// resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:centerResult:alignCorners:layout:name:"), imagesTensor, size, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using bilinear sampling. Result images
// will be distorted if size is of different aspect ratio. Destination indices
// are computed using direct index scaling by default, with no offset added.
// If the centerResult parameter is true, the destination indices will be
// scaled and shifted to be centered on the input image. If the alignCorners
// parameter is true, the corners of the result images will match the input
// images. Scaling will be modified to a factor of (size - 1) / (inputSize -
// 1). When alignCorners is true, the centerResult parameter does nothing. In
// order to achieve the same behavior as OpenCV’s resize and
// TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:centerResult:alignCorners:name:)
func (g MPSGraph) ResizeBilinearWithTensorSizeTensorCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, centerResult bool, alignCorners bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:centerResult:alignCorners:name:"), imagesTensor, size, centerResult, alignCorners, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Resamples input images to given size using the provided scale and offset
// and bilinear sampling See above discussion for more details.
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:scaleOffsetTensor:layout:name:)
func (g MPSGraph) ResizeBilinearWithTensorSizeTensorScaleOffsetTensorLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:scaleOffsetTensor:layout:name:"), imagesTensor, size, scaleOffset, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using the provided scale and offset
// and bilinear sampling. Destination indices are computed using
//
// For most use cases passing the scale and offset directly is unnecessary,
// and it is preferable to use the API specifying centerResult and
// alignCorners.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(_:sizeTensor:scaleTensor:offsetTensor:name:)
func (g MPSGraph) ResizeBilinearWithTensorSizeTensorScaleTensorOffsetTensorName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithTensor:sizeTensor:scaleTensor:offsetTensor:name:"), imagesTensor, size, scale, offset, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeBilinearWithGradientTensorInputCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithGradientTensor:input:centerResult:alignCorners:layout:name:"), gradient, input, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with bilinear sampling
// and identical parameters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:scale:offsetTensor:name:)
func (g MPSGraph) ResizeBilinearWithGradientTensorInputScaleTensorOffsetTensorName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithGradientTensor:input:scaleTensor:offsetTensor:name:"), gradient, input, scale, offset, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with bilinear sampling
// and identical parameters. See discussion of resizeTensor for more in depth
// description of resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeBilinear(withGradientTensor:input:scaleOffsetTensor:layout:name:)
func (g MPSGraph) ResizeBilinearWithGradientTensorInputScaleOffsetTensorLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeBilinearWithGradientTensor:input:scaleOffsetTensor:layout:name:"), gradient, input, scaleOffset, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Resamples input images to given size using nearest neighbor sampling.
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling. Default is
// roundPreferCeil.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// This API allows for the rounding mode to be specified. Resamples input
// images to given size. Result images will be distorted if size is of
// different aspect ratio. Resize supports the following modes: Nearest
// Neighbor - values are interpolated using the closest neighbor pixel
// Bilinear - values are computed using bilinear interpolation of 4
// neighboring pixels Destination indices are computed using direct index
// scaling by default, with no offset added. If the centerResult parameter is
// true, the destination indices will be scaled and shifted to be centered on
// the input image. If the alignCorners parameter is true, the corners of the
// result images will match the input images. Scaling will be modified to a
// factor of (size - 1) / (inputSize - 1). When alignCorners is true, the
// centerResult parameter does nothing. In order to achieve the same behavior
// as OpenCV’s resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:nearestRoundingMode:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:nearestRoundingMode:centerResult:alignCorners:layout:name:"), imagesTensor, size, nearestRoundingMode, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling. Default is
// roundPreferCeil.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using nearest neighbor sampling.
// Result images will be distorted if size is of different aspect ratio.
// Destination indices are computed using direct index scaling by default,
// with no offset added. If the centerResult parameter is true, the
// destination indices will be scaled and shifted to be centered on the input
// image. If the alignCorners parameter is true, the corners of the result
// images will match the input images. Scaling will be modified to a factor of
// (size - 1) / (inputSize - 1). When alignCorners is true, the centerResult
// parameter does nothing. In order to achieve the same behavior as OpenCV’s
// resize and TensorFlowV2’s resize,
//
// # To achieve the same behavior as TensorFlowV1 resize
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:nearestRoundingMode:centerResult:alignCorners:name:)
func (g MPSGraph) ResizeNearestWithTensorSizeTensorNearestRoundingModeCenterResultAlignCornersName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:nearestRoundingMode:centerResult:alignCorners:name:"), imagesTensor, size, nearestRoundingMode, centerResult, alignCorners, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Resamples input images to given size using the provided scale and offset
// and nearest neighbor sampling See above discussion for more details.
//
// imagesTensor: Tensor containing input images.
//
// size: 1D Int32 or Int64 tensor. A 2-element shape as [newHeight, newWidth]
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:scaleOffsetTensor:nearestRoundingMode:layout:name:)
func (g MPSGraph) ResizeNearestWithTensorSizeTensorScaleOffsetTensorNearestRoundingModeLayoutName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:scaleOffsetTensor:nearestRoundingMode:layout:name:"), imagesTensor, size, scaleOffset, nearestRoundingMode, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize operation and returns the result tensor.
//
// imagesTensor: Tensor containing input images.
//
// size: The target size of the result tensor. 1D Int32 or Int64 tensor of size
// equal to rank of input.
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling. Default is
// roundPreferCeil.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Resamples input images to given size using the provided scale and offset
// and nearest neighbor sampling. Destination indices are computed using
//
// For most use cases passing the scale and offset directly is unnecessary,
// and it is preferable to use the API specifying centerResult and
// alignCorners.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(_:sizeTensor:scaleTensor:offsetTensor:nearestRoundingMode:name:)
func (g MPSGraph) ResizeNearestWithTensorSizeTensorScaleTensorOffsetTensorNearestRoundingModeName(imagesTensor IMPSGraphTensor, size IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithTensor:sizeTensor:scaleTensor:offsetTensor:nearestRoundingMode:name:"), imagesTensor, size, scale, offset, nearestRoundingMode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling.
//
// centerResult: Controls if the result image is centered on the input image. When NO, the
// result will have the top left corner aligned
//
// alignCorners: When YES, the result image will have the same value as the input image in
// the corners
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:nearestRoundingMode:centerResult:alignCorners:layout:name:)
func (g MPSGraph) ResizeNearestWithGradientTensorInputNearestRoundingModeCenterResultAlignCornersLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, centerResult bool, alignCorners bool, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithGradientTensor:input:nearestRoundingMode:centerResult:alignCorners:layout:name:"), gradient, input, nearestRoundingMode, centerResult, alignCorners, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scale: 1D float tensor of size equal to rank of input.
//
// offset: 1D float tensor of size equal to rank of input.
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling. Default is
// roundPreferCeil.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with nearest neighbor
// sampling and identical parameters. See discussion of resizeTensor for more
// in depth description of resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:scale:offsetTensor:nearestRoundingMode:name:)
func (g MPSGraph) ResizeNearestWithGradientTensorInputScaleTensorOffsetTensorNearestRoundingModeName(gradient IMPSGraphTensor, input IMPSGraphTensor, scale IMPSGraphTensor, offset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithGradientTensor:input:scaleTensor:offsetTensor:nearestRoundingMode:name:"), gradient, input, scale, offset, nearestRoundingMode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Resize gradient operation and returns the result tensor.
//
// gradient: Incoming gradient tensor
//
// input: Forward pass input tensor
//
// scaleOffset: 1D float tensor. A 4-element shape as [scaleY, scaleX, offsetY, offsetX]
//
// nearestRoundingMode: The rounding mode to use when using nearest resampling.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC, NCHW, HWC, CHW, and HW.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Computes the gradient for the forward pass Resize op with identical
// parameters. See discussion of resizeTensor for more in depth description of
// resize paramters.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/resizeNearest(withGradientTensor:input:scaleOffsetTensor:nearestRoundingMode:layout:name:)
func (g MPSGraph) ResizeNearestWithGradientTensorInputScaleOffsetTensorNearestRoundingModeLayoutName(gradient IMPSGraphTensor, input IMPSGraphTensor, scaleOffset IMPSGraphTensor, nearestRoundingMode MPSGraphResizeNearestRoundingMode, layout MPSGraphTensorNamedDataLayout, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resizeNearestWithGradientTensor:input:scaleOffsetTensor:nearestRoundingMode:layout:name:"), gradient, input, scaleOffset, nearestRoundingMode, layout, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reverse operation and returns the result tensor.
//
// tensor: The tensor to be reversed.
//
// axes: A tensor that specifies axes to be reversed (Axes must be unique and within
// normal axis range).
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Reverses a tensor on given axes. Semantics based on [TensorFlow reverse
// op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:axes:name:)
//
// [TensorFlow reverse op]: https://www.tensorflow.org/api_docs/python/tf/reverse
func (g MPSGraph) ReverseTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reverseTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reverse operation and returns the result tensor.
//
// tensor: The tensor to be reversed.
//
// axesTensor: A tensor that specifies axes to be reversed (Axes must be unique and within
// normal axis range).
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Reverses a tensor on given axes. Semantics based on [TensorFlow reverse
// op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:axesTensor:name:)
//
// [TensorFlow reverse op]: https://www.tensorflow.org/api_docs/python/tf/reverse
func (g MPSGraph) ReverseTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reverseTensor:axesTensor:name:"), tensor, axesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a reverse operation and returns the result tensor.
//
// tensor: The tensor to be reversed.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Reverses a tensor on all axes. Semantics based on [TensorFlow reverse op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/reverse(_:name:)
//
// [TensorFlow reverse op]: https://www.tensorflow.org/api_docs/python/tf/reverse
func (g MPSGraph) ReverseTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("reverseTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Rounds the input tensor elements by rounding to nearest even.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/rint(with:name:)
func (g MPSGraph) RintWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rintWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Rounds the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/round(with:name:)
func (g MPSGraph) RoundWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("roundWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// # Return Value
//
// A valid MPSGraphTensor : MPSGraphTensorData dictionary with results
// synchronized to the CPU memory.
//
// # Discussion
//
// This call blocks until execution has completed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(feeds:targetTensors:targetOperations:)
func (g MPSGraph) RunWithFeedsTargetTensorsTargetOperations(feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation) MPSGraphTensorDataDictionary {
	rv := objc.Send[MPSGraphTensorDataDictionary](g.ID, objc.Sel("runWithFeeds:targetTensors:targetOperations:"), feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations))
	return MPSGraphTensorDataDictionary(rv)
}

// Runs the graph for the given feeds and returns the target tensor values in
// the results dictionary provided by the user.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetOperations: Operations to be completed at the end of the run.
//
// resultsDictionary: MPSGraphTensors dictionary passed by user, these will be filled with graph
// output data.
//
// # Discussion
//
// It also ensures all target operations also executed. This call blocks until
// execution has completed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(with:feeds:targetOperations:resultsDictionary:)
func (g MPSGraph) RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary) {
	objc.Send[objc.ID](g.ID, objc.Sel("runWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:"), commandQueue, feeds, objectivec.IObjectSliceToNSArray(targetOperations), resultsDictionary)
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// # Return Value
//
// A valid MPSGraphTensor : MPSGraphTensorData dictionary with results
// synchronized to the CPU memory.
//
// # Discussion
//
// This call blocks until execution has completed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/run(with:feeds:targetTensors:targetOperations:)
func (g MPSGraph) RunWithMTLCommandQueueFeedsTargetTensorsTargetOperations(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation) MPSGraphTensorDataDictionary {
	rv := objc.Send[MPSGraphTensorDataDictionary](g.ID, objc.Sel("runWithMTLCommandQueue:feeds:targetTensors:targetOperations:"), commandQueue, feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations))
	return MPSGraphTensorDataDictionary(rv)
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Return Value
//
// A valid MPSGraphTensor : MPSGraphTensorData dictionary with results
// synchronized to the CPU memory.
//
// # Discussion
//
// This call is asynchronous and will return immediately if a
// completionHandler is set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(feeds:targetTensors:targetOperations:executionDescriptor:)
func (g MPSGraph) RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary {
	rv := objc.Send[MPSGraphTensorDataDictionary](g.ID, objc.Sel("runAsyncWithFeeds:targetTensors:targetOperations:executionDescriptor:"), feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations), executionDescriptor)
	return MPSGraphTensorDataDictionary(rv)
}

// Encodes the graph for the given feeds to returns the target tensor values
// in the results dictionary provided by the user.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetOperations: Operations to be completed at the end of the run.
//
// resultsDictionary: MPSGraphTensors dictionary passed by user, these will be filled with graph
// output data.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Discussion
//
// It ensures all target operations also executed. This call is asynchronous
// and will return immediately if a completionHandler is set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(with:feeds:targetOperations:resultsDictionary:executionDescriptor:)
func (g MPSGraph) RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetOperations []MPSGraphOperation, resultsDictionary MPSGraphTensorDataDictionary, executionDescriptor IMPSGraphExecutionDescriptor) {
	objc.Send[objc.ID](g.ID, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:executionDescriptor:"), commandQueue, feeds, objectivec.IObjectSliceToNSArray(targetOperations), resultsDictionary, executionDescriptor)
}

// Runs the graph for the given feeds and returns the target tensor values,
// ensuring all target operations also executed.
//
// commandQueue: CommandQueue passed to exectute the graph on.
//
// feeds: Feeds dictionary for the placeholder tensors.
//
// targetTensors: Tensors for which the caller wishes MPSGraphTensorData to be returned.
//
// targetOperations: Operations to be completed at the end of the run.
//
// executionDescriptor: ExecutionDescriptor to be passed in and used.
//
// # Return Value
//
// A valid MPSGraphTensor : MPSGraphTensorData dictionary with results
// synchronized to the CPU memory if MPSGraphOptionsSynchronizeResults set.
//
// # Discussion
//
// This call is asynchronous and will return immediately if a
// completionHandler is set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/runAsync(with:feeds:targetTensors:targetOperations:executionDescriptor:)
func (g MPSGraph) RunAsyncWithMTLCommandQueueFeedsTargetTensorsTargetOperationsExecutionDescriptor(commandQueue metal.MTLCommandQueue, feeds MPSGraphTensorDataDictionary, targetTensors []MPSGraphTensor, targetOperations []MPSGraphOperation, executionDescriptor IMPSGraphExecutionDescriptor) MPSGraphTensorDataDictionary {
	rv := objc.Send[MPSGraphTensorDataDictionary](g.ID, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetTensors:targetOperations:executionDescriptor:"), commandQueue, feeds, objectivec.IObjectSliceToNSArray(targetTensors), objectivec.IObjectSliceToNSArray(targetOperations), executionDescriptor)
	return MPSGraphTensorDataDictionary(rv)
}

// Samples a tensor using the coordinates provided, using nearest neighbor
// sampling with specified rounding mode.
//
// source: Tensor containing source data
//
// coordinates: A tensor (N, Hout, Wout, 2) that contains the coordinates of the samples in
// the source tensor that constitute the output tensor.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC and NCHW.
//
// normalizeCoordinates: If true, coordinates are within [-1, 1] x [-1, 1] otherwise they are in
// pixels in the input tensor.
//
// relativeCoordinates: If true, coordinates are relative to the postion of the pixel in the output
// tensor and scaled back to the input tensor size
//
// alignCorners: If true, coordinate extrema are equal to the center of edge pixels,
// otherwise extrema are equal to outer edge of edge pixels
//
// paddingMode: Determines how samples outside the inputTensor are evaluated (only
// constant, reflect, symmetric and clampToEdge are supported)
//
// nearestRoundingMode: The rounding mode to use for determining the nearest neighbor. Valid modes
// are roundPreferCeil, roundPreferFloor, ceil, and floor.
//
// constantValue: If paddingMode is MPSGraphPaddingModeConstant, then this constant is used
// for samples outside the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Given an input tensor (N, H1, W1, C) or (N, C, H1, W1) and coordinates
// tensor (N, H2, W2, 2) this operation outputs a tensor of size (N, H2, W2,
// C) or (N, C, H2, W2) by sampling the input tensor at the coordinates
// provided by the coordinates tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sampleGrid(withSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:nearestRoundingMode:constantValue:name:)
func (g MPSGraph) SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeNearestRoundingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode MPSGraphPaddingMode, nearestRoundingMode MPSGraphResizeNearestRoundingMode, constantValue float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sampleGridWithSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:nearestRoundingMode:constantValue:name:"), source, coordinates, layout, normalizeCoordinates, relativeCoordinates, alignCorners, paddingMode, nearestRoundingMode, constantValue, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Samples a tensor using the coordinates provided.
//
// source: Tensor containing source data
//
// coordinates: A tensor (N, Hout, Wout, 2) that contains the coordinates of the samples in
// the source tensor that constitute the output tensor.
//
// layout: Specifies what layout the provided tensor is in. The returned tensor will
// follow the same layout. Valid layouts are NHWC and NCHW.
//
// normalizeCoordinates: If true, coordinates are within [-1, 1] x [-1, 1] otherwise they are in
// pixels in the input tensor.
//
// relativeCoordinates: If true, coordinates are relative to the postion of the pixel in the output
// tensor and scaled back to the input tensor size
//
// alignCorners: If true, coordinate extrema are equal to the center of edge pixels,
// otherwise extrema are equal to outer edge of edge pixels
//
// paddingMode: Determines how samples outside the inputTensor are evaluated (only
// constant, reflect, symmetric and clampToEdge are supported)
//
// samplingMode: Can be either MPSGraphResizeNearest or MPSGraphResizeBilinear. Nearest
// sampling will use roundPreferCeil.
//
// constantValue: If paddingMode is MPSGraphPaddingModeConstant, then this constant is used
// for samples outside the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Given an input tensor (N, H1, W1, C) or (N, C, H1, W1) and coordinates
// tensor (N, H2, W2, 2) this operation outputs a tensor of size (N, H2, W2,
// C) or (N, C, H2, W2) by sampling the input tensor at the coordinates
// provided by the coordinates tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sampleGrid(withSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:samplingMode:constantValue:name:)
func (g MPSGraph) SampleGridWithSourceTensorCoordinateTensorLayoutNormalizeCoordinatesRelativeCoordinatesAlignCornersPaddingModeSamplingModeConstantValueName(source IMPSGraphTensor, coordinates IMPSGraphTensor, layout MPSGraphTensorNamedDataLayout, normalizeCoordinates bool, relativeCoordinates bool, alignCorners bool, paddingMode MPSGraphPaddingMode, samplingMode MPSGraphResizeMode, constantValue float64, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sampleGridWithSourceTensor:coordinateTensor:layout:normalizeCoordinates:relativeCoordinates:alignCorners:paddingMode:samplingMode:constantValue:name:"), source, coordinates, layout, normalizeCoordinates, relativeCoordinates, alignCorners, paddingMode, samplingMode, constantValue, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a scaled dot product attention (SDPA) operation and returns the
// result tensor.
//
// queryTensor: A tensor that represents the query projection.
//
// keyTensor: A tensor that represents the key projection.
//
// valueTensor: A tensor that represents the value projection.
//
// maskTensor: An optional tensor that contains a mask that is applied to the scaled,
// matrix multiplied query and value matrices. If mask tensor is nil, the QK^T
// is not element-wise masked.
//
// scale: A scale that is applied to the result of query and value matrix multiply.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// SDPA Op computes attention by computing softmax(scale * QK^T + M)V.
// queryTensor Q with shape [B, Hq, Nq, F] and keyTensor K with shape [B, Hq,
// Nkv, F], with Q’s H dimension expandable to satisfy matmul QK^T.
// maskTensor M’s shape should be broadcast compatible to satisfy (QK^T +
// M). valueTensor V with shape [B, Hv, Nkv, F] should satisfy the matmul
// (QK^T + M)V.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scaledDotProductAttention(query:key:value:mask:scale:name:)
func (g MPSGraph) ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorMaskTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, maskTensor IMPSGraphTensor, scale float32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scaledDotProductAttentionWithQueryTensor:keyTensor:valueTensor:maskTensor:scale:name:"), queryTensor, keyTensor, valueTensor, maskTensor, scale, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a scaled dot product attention (SDPA) operation (without a mask)
// and returns the result tensor.
//
// queryTensor: A tensor that represents the query projection.
//
// keyTensor: A tensor that represents the key projection.
//
// valueTensor: A tensor that represents the value projection.
//
// scale: A scale that is applied on the result of query and value matrix multiply.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scaledDotProductAttention(query:key:value:scale:name:)
func (g MPSGraph) ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorScaleName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, scale float32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scaledDotProductAttentionWithQueryTensor:keyTensor:valueTensor:scale:name:"), queryTensor, keyTensor, valueTensor, scale, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Scatter operation and returns the result tensor.
//
// updatesTensor: Tensor containing values to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the result indices to insert values at.
//
// shape: The shape of the result tensor.
//
// axis: The axis of the result tensor to scatter values along.
//
// mode: The type of update to use on the destination.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Scatters the slices in updatesTensor to the result tensor along the indices
// in indicesTensor. The scatter is defined as
//
// Collisions will be updated according to mode. The tensors have the
// following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatter(_:indices:shape:axis:mode:name:)
func (g MPSGraph) ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, axis int, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterWithUpdatesTensor:indicesTensor:shape:axis:mode:name:"), updatesTensor, indicesTensor, shape, axis, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// axis: The axis to scatter to. Negative values wrap around
//
// dataTensor: The input tensor to scatter values onto
//
// updatesTensor: The input tensor to scatter values from
//
// indicesTensor: Int32 or Int64 tensor used to index the result tensor.
//
// mode: The type of update to use
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatter values from `updatesTensor` along the specified `axis` at indices
// in `indicesTensor` onto `dataTensor`. Values in `dataTensor` are updated
// following `mode`. See MPSGraphScatterMode. The shape of `updatesTensor` and
// `indicesTensor` must match. The shape of `dataTensor` must match except at
// `axis`. If an index is out of bounds of `shape` along `axis` the update
// value is skipped. For example,
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxis(_:data:updates:indices:mode:name:)
func (g MPSGraph) ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterAlongAxis:withDataTensor:updatesTensor:indicesTensor:mode:name:"), axis, dataTensor, updatesTensor, indicesTensor, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// axis: The axis to scatter to. Negative values wrap around
//
// updatesTensor: The input tensor to scatter values from
//
// indicesTensor: Int32 or Int64 tensor used to index the result tensor.
//
// mode: The type of update to use
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatter values from `updatesTensor` along the specified `axis` at indices
// in `indicesTensor` into a result tensor. Values are updated following
// `mode`. See MPSGraphScatterMode. The shape of `updatesTensor` and
// `indicesTensor` must match. `shape` must match except at `axis`. The shape
// of the result tensor is equal to `shape` and initialized with an initial
// value corresponding to `mode`. If an index is out of bounds of `shape`
// along `axis` the update value is skipped.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxis(_:updates:indices:shape:mode:name:)
func (g MPSGraph) ScatterAlongAxisWithUpdatesTensorIndicesTensorShapeModeName(axis int, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterAlongAxis:withUpdatesTensor:indicesTensor:shape:mode:name:"), axis, updatesTensor, indicesTensor, shape, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// axisTensor: Scalar Int32 tensor. The axis to scatter to. Negative values wrap around
//
// dataTensor: The input tensor to scatter values onto
//
// updatesTensor: The input tensor to scatter values from
//
// indicesTensor: Int32 or Int64 tensor used to index the result tensor.
//
// mode: The type of update to use
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatter values from `updatesTensor` along the specified `axis` at indices
// in `indicesTensor` onto `dataTensor`. Values in `dataTensor` are updated
// following `mode`. See MPSGraphScatterMode. The shape of `updatesTensor` and
// `indicesTensor` must match. The shape of `dataTensor` must match except at
// `axis`. If an index is out of bounds of `shape` along `axis` the update
// value is skipped. For example,
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxisTensor(_:data:updates:indices:mode:name:)
func (g MPSGraph) ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor IMPSGraphTensor, dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterAlongAxisTensor:withDataTensor:updatesTensor:indicesTensor:mode:name:"), axisTensor, dataTensor, updatesTensor, indicesTensor, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterAlongAxis operation and returns the result tensor.
//
// axisTensor: Scalar Int32 tensor. The axis to scatter to. Negative values wrap around
//
// updatesTensor: The input tensor to scatter values from
//
// indicesTensor: Int32 or Int64 tensor used to index the result tensor.
//
// mode: The type of update to use
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatter values from `updatesTensor` along the specified `axis` at indices
// in `indicesTensor` into a result tensor. Values are updated following
// `mode`. See MPSGraphScatterMode. The shape of `updatesTensor` and
// `indicesTensor` must match. `shape` must match except at `axis`. The shape
// of the result tensor is equal to `shape` and initialized with an initial
// value corresponding to `mode`. If an index is out of bounds of `shape`
// along `axis` the update value is skipped.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterAlongAxisTensor(_:updates:indices:shape:mode:name:)
func (g MPSGraph) ScatterAlongAxisTensorWithUpdatesTensorIndicesTensorShapeModeName(axisTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterAlongAxisTensor:withUpdatesTensor:indicesTensor:shape:mode:name:"), axisTensor, updatesTensor, indicesTensor, shape, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterND operation and returns the result tensor.
//
// updatesTensor: Tensor containing slices to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the result indices to insert slices at
//
// shape: The shape of the result tensor.
//
// batchDimensions: The number of batch dimensions
//
// mode: The type of update to use on the destination
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatters the slices in updatesTensor to the result tensor along the indices
// in indicesTensor. The scatter is defined as
//
// Collisions will be summed, and slices not set by indices are set to 0. The
// tensors have the following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterND(withUpdatesTensor:indicesTensor:shape:batchDimensions:mode:name:)
func (g MPSGraph) ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, batchDimensions uint, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterNDWithUpdatesTensor:indicesTensor:shape:batchDimensions:mode:name:"), updatesTensor, indicesTensor, shape, batchDimensions, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterND operation and returns the result tensor.
//
// updatesTensor: Tensor containing slices to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the result indices to insert slices at
//
// shape: The shape of the result tensor.
//
// batchDimensions: The number of batch dimensions
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatters the slices in updatesTensor to the result tensor along the indices
// in indicesTensor. The scatter is defined as
//
// Collisions will be summed, and slices not set by indices are set to 0. The
// tensors have the following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterND(withUpdatesTensor:indicesTensor:shape:batchDimensions:name:)
func (g MPSGraph) ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsName(updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, shape foundation.NSArray, batchDimensions uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterNDWithUpdatesTensor:indicesTensor:shape:batchDimensions:name:"), updatesTensor, indicesTensor, shape, batchDimensions, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a ScatterND operation and returns the result tensor.
//
// dataTensor: Tensor containing inital values of same shape as result tensor
//
// updatesTensor: Tensor containing slices to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the result indices to insert slices at
//
// batchDimensions: The number of batch dimensions
//
// mode: The type of update to use on the destination
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatters the slices in updatesTensor to the result tensor along the indices
// in indicesTensor, on top of dataTensor. The scatter is defined as
//
// Collisions will be updated according to mode, and slices not set by indices
// are set to 0. The tensors have the following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterNDWithData(_:updates:indices:batchDimensions:mode:name:)
func (g MPSGraph) ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, batchDimensions uint, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterNDWithDataTensor:updatesTensor:indicesTensor:batchDimensions:mode:name:"), dataTensor, updatesTensor, indicesTensor, batchDimensions, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a Scatter operation and returns the result tensor.
//
// dataTensor: Tensor containing inital values of same shape as result tensor
//
// updatesTensor: Tensor containing values to be inserted into the result tensor.
//
// indicesTensor: Tensor containg the result indices to insert values at
//
// axis: The axis of the result tensor to scatter values along
//
// mode: The type of update to use on the destination
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Scatters the slices in updatesTensor to the result tensor along the indices
// in indicesTensor, on top of dataTensor. The scatter is defined as
//
// Collisions will be updated according to mode. The tensors have the
// following shape requirements
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scatterWithData(_:updates:indices:axis:mode:name:)
func (g MPSGraph) ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor IMPSGraphTensor, updatesTensor IMPSGraphTensor, indicesTensor IMPSGraphTensor, axis int, mode MPSGraphScatterMode, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scatterWithDataTensor:updatesTensor:indicesTensor:axis:mode:name:"), dataTensor, updatesTensor, indicesTensor, axis, mode, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Selects values from either the true or false predicate tensor, depending on
// the values in the first input.
//
// predicateTensor: The predicate tensor.
//
// truePredicateTensor: The tensor to select values from if predicate is true.
//
// falseSelectTensor: The tensor to select values from if predicate is false.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a select operation and returns the result tensor. It
// supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/select(predicate:trueTensor:falseTensor:name:)
func (g MPSGraph) SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor IMPSGraphTensor, truePredicateTensor IMPSGraphTensor, falseSelectTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("selectWithPredicateTensor:truePredicateTensor:falsePredicateTensor:name:"), predicateTensor, truePredicateTensor, falseSelectTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a shape-of operation and returns the result tensor.
//
// tensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Returns a rank-1 tensor of type [MPSDataTypeInt32] with the values of the
// static shape of the input tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/shapeOf(_:name:)
func (g MPSGraph) ShapeOfTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shapeOfTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the sigmoid operation on an input tensor.
//
// tensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sigmoid(with:name:)
func (g MPSGraph) SigmoidWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sigmoidWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the gradient of the sigmoid function using the incoming gradient
// tensor.
//
// gradient: The incoming gradient tensor.
//
// source: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sigmoidGradient(withIncomingGradient:sourceTensor:name:)
func (g MPSGraph) SigmoidGradientWithIncomingGradientSourceTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sigmoidGradientWithIncomingGradient:sourceTensor:name:"), gradient, source, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the sign of the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation returns 1.0 if the correspnding input element is greater
// than 0, -1.0 if it is lesser than 0, -0.0 if it is equal to -0.0, and +0.0
// if it is equal to +0.0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sign(with:name:)
func (g MPSGraph) SignWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("signWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the sign bit of the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation returns `true` if the sign bit is set for the correspnding
// floating-point input element, otherwise it returns `false`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/signbit(with:name:)
func (g MPSGraph) SignbitWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("signbitWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the sine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sin(with:name:)
func (g MPSGraph) SinWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sinWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a single-gate RNN operation and returns the value and optionally
// the training state tensor.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 1 or 2, depending on value of
// `descriptor.Training()`. The layout of the both outputs is [T,N,H] or
// [T,N,2H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is `bias`, `m`
// is optional `mask`, `x[t]` is `source` `h[t]` is the first output, `z[t]`
// is the second output (optional) and `h[-1]` is `initState`. See
// [MPSGraphSingleGateRNNDescriptor] for different `activation` options.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:initState:descriptor:name:)
func (g MPSGraph) SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:initState:descriptor:name:"), source, recurrentWeight, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN operation and returns the value and optionally
// the training state tensor.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [2H,I] and otherwise it is [H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [2H] and otherwise it is
// [H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 1 or 2, depending on value of
// `descriptor.Training()`. The layout of the both outputs is [T,N,H] or
// [T,N,2H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is `bias`, `m`
// is optional `mask`, `x[t]` is `source` `h[t]` is the first output, `z[t]`
// is the second output (optional) and `h[-1]` is `initState`. See
// [MPSGraphSingleGateRNNDescriptor] for different `activation` options.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:inputWeight:bias:initState:descriptor:name:)
func (g MPSGraph) SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN operation and returns the value and optionally
// the training state tensor.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H].
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [2H,I] and otherwise it is [H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [2H] and otherwise it is
// [H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. This is useful for dropout support.
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 1 or 2, depending on value of
// `descriptor.Training()`. The layout of the both outputs is [T,N,H] or
// [T,N,2H] for bidirectional.
//
// # Discussion
//
// This operation returns tensors `h` and optionally `z` that are defined
// recursively as follows:
//
// [W] is optional `inputWeight`, [R] is `recurrentWeight`, `b` is `bias`, `m`
// is optional `mask`, `x[t]` is `source` `h[t]` is the first output, `z[t]`
// is the second output (optional) and `h[-1]` is `initState`. See
// [MPSGraphSingleGateRNNDescriptor] for different `activation` options.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:)
func (g MPSGraph) SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, mask, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN gradient operation and returns the gradient
// tensor values.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H]. Note: For `bidirectional` this
// tensor must have a static shape.
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]
// with `descriptor.Training() = YES`.
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is `nil`, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias` and finally
// for `initState`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:initState:descriptor:name:)
func (g MPSGraph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN gradient operation and returns the gradient
// tensor values.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H]. Note: For `bidirectional` this
// tensor must have a static shape.
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]
// with `descriptor.Training() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [2H,I] and otherwise it is [H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [2H] and otherwise it is
// [H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is `nil`, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias` and finally
// for `initState`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:descriptor:name:)
func (g MPSGraph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, inputWeight, bias, initState, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN gradient operation and returns the gradient
// tensor values.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H]. Note: For `bidirectional` this
// tensor must have a static shape.
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]
// with `descriptor.Training() = YES`.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [2H,I] and otherwise it is [H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [2H] and otherwise it is
// [H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. This is useful for dropout support.
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is `nil`, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias` and finally
// for `initState`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:mask:descriptor:name:)
func (g MPSGraph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, inputWeight, bias, initState, mask, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a single-gate RNN gradient operation and returns the gradient
// tensor values.
//
// source: A tensor that contains the source data `x[t]` with the data layout [T,N,I].
// In case `inputWeight = nil` and `bidirectional = NO` then the layout is
// [T,N,H] and for `inputWeight = nil` and `bidirectional = YES` the layout is
// [T,N,2H].
//
// recurrentWeight: A tensor containing the recurrent weights [R]. For `bidirectional` the
// layout is [2,H,H] and otherwise it is [H,H]. Note: For `bidirectional` this
// tensor must have a static shape.
//
// sourceGradient: The input gradient, that is the gradient of a tensor with respect to the
// first output of the forward pass.
//
// zState: The second output of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]
// with `descriptor.Training() = YES`.
//
// stateGradient: The input gradient coming from the future timestep - optional, if missing
// the operation assumes zeroes.
//
// inputWeight: A tensor containing the input weights matrix [W] - optional, if missing the
// operation assumes a diagonal unit-matrix. For `bidirectional` the layout is
// [2H,I] and otherwise it is [H,I].
//
// bias: A tensor containing the bias `b` - optional, if missing the operation
// assumes zeroes. For `bidirectional` the layout is [2H] and otherwise it is
// [H].
//
// initState: The initial internal state of the RNN `h[-1]` - optional, if missing the
// operation assumes zeroes. For `bidirectional` the layout is [N,2H] and
// otherwise it is [N,H].
//
// mask: A tensor containing the mask `m` - optional, if missing the operation
// assumes ones. This is useful for dropout support.
//
// descriptor: A descriptor that defines the parameters for the RNN operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] array containing gradients for each input tensor,
// except for `sourceGradient` and `mask`. In case an input is `nil`, no
// gradient will be returned for it. The order of the gradients will be: for
// `source`, for `recurrentWeight`, for `inputWeight`, for `bias` and finally
// for `initState`.
//
// # Discussion
//
// For details of this operation and parameters, refer to documentation of
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNNGradients(_:recurrentWeight:sourceGradient:zState:stateGradient:inputWeight:bias:initState:mask:descriptor:name:)
func (g MPSGraph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, stateGradient IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:stateGradient:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, stateGradient, inputWeight, bias, initState, mask, descriptor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Applies the hyperbolic sine operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sinh(with:name:)
func (g MPSGraph) SinhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sinhWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice gradient operation and returns the result tensor.
//
// inputGradientTensor: The input gradient.
//
// fwdInShapeTensor: The shape of the forward pass input, that is the shape of the gradient
// output.
//
// startTensor: The tensor that specifies the starting points for each dimension.
//
// endTensor: The tensor that specifies the ending points for each dimension.
//
// strideTensor: The tensor that specifies the strides for each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:start:end:strideTensor:startMask:endMask:squeezeMask:name:)
func (g MPSGraph) SliceGradientTensorFwdInShapeTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:startTensor:endTensor:strideTensor:startMask:endMask:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, startTensor, endTensor, strideTensor, startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a slice gradient operation and returns the result tensor.
//
// inputGradientTensor: The input gradient.
//
// fwdInShapeTensor: The shape of the forward pass input, that is the shape of the gradient
// output.
//
// startTensor: The tensor that specifies the starting points for each dimension.
//
// sizeTensor: The tensor that specifies the size of the forward result for each
// dimension.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:start:sizeTensor:squeezeMask:name:)
func (g MPSGraph) SliceGradientTensorFwdInShapeTensorStartTensorSizeTensorSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:startTensor:sizeTensor:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, startTensor, sizeTensor, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice gradient operation and returns the result tensor.
//
// inputGradientTensor: The input gradient.
//
// fwdInShapeTensor: The shape of the forward pass input, that is the shape of the gradient
// output.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:starts:ends:strides:name:)
func (g MPSGraph) SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:starts:ends:strides:name:"), inputGradientTensor, fwdInShapeTensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice gradient operation and returns the result tensor.
//
// inputGradientTensor: The input gradient.
//
// fwdInShapeTensor: The shape of the forward pass input, that is the shape of the gradient
// output.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceGradientTensor(_:fwdInShapeTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:)
func (g MPSGraph) SliceGradientTensorFwdInShapeTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(inputGradientTensor IMPSGraphTensor, fwdInShapeTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceGradientTensor:fwdInShapeTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), inputGradientTensor, fwdInShapeTensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a slice operation and returns the result tensor.
//
// tensor: The tensor to be sliced.
//
// dimensionIndex: The dimension to slice.
//
// start: The starting index of the slice, can be negative to count from the end of
// the tensor dimension.
//
// length: The length of the slice.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:dimension:start:length:name:)
func (g MPSGraph) SliceTensorDimensionStartLengthName(tensor IMPSGraphTensor, dimensionIndex uint, start int, length int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceTensor:dimension:start:length:name:"), tensor, dimensionIndex, start, length, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice operation and returns the result tensor.
//
// tensor: The Tensor to be sliced.
//
// startTensor: The tensor that specifies the starting points for each dimension.
//
// endTensor: The tensor that specifies the ending points for each dimension.
//
// strideTensor: The tensor that specifies the strides for each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Slices a tensor starting from `startTensor`, stopping short before
// `endTensor` stepping `strideTensor` paces between each value. Semantics
// based on [TensorFlow Strided Slice Op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:start:end:strideTensor:startMask:endMask:squeezeMask:name:)
//
// [TensorFlow Strided Slice Op]: https://www.tensorflow.org/api_docs/python/tf/strided_slice
func (g MPSGraph) SliceTensorStartTensorEndTensorStrideTensorStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, endTensor IMPSGraphTensor, strideTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceTensor:startTensor:endTensor:strideTensor:startMask:endMask:squeezeMask:name:"), tensor, startTensor, endTensor, strideTensor, startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a slice operation and returns the result tensor.
//
// tensor: The Tensor to be sliced.
//
// startTensor: The tensor that specifies the starting points for each dimension.
//
// sizeTensor: The tensor that specifies the size of the result for each dimension.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Slices a tensor starting from `startTensor`, stopping short before
// `startTensor + endTensor` stepping a single pace between each value.
// Semantics based on [TensorFlow Strided Slice Op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:start:sizeTensor:squeezeMask:name:)
//
// [TensorFlow Strided Slice Op]: https://www.tensorflow.org/api_docs/python/tf/strided_slice
func (g MPSGraph) SliceTensorStartTensorSizeTensorSqueezeMaskName(tensor IMPSGraphTensor, startTensor IMPSGraphTensor, sizeTensor IMPSGraphTensor, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceTensor:startTensor:sizeTensor:squeezeMask:name:"), tensor, startTensor, sizeTensor, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice operation and returns the result tensor.
//
// tensor: The tensor to be sliced.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Slices a tensor starting from `starts`, stopping short before `ends`
// stepping `strides` paces between each value. Semantics based on [TensorFlow
// Strided Slice Op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:starts:ends:strides:name:)
//
// [TensorFlow Strided Slice Op]: https://www.tensorflow.org/api_docs/python/tf/strided_slice
func (g MPSGraph) SliceTensorStartsEndsStridesName(tensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceTensor:starts:ends:strides:name:"), tensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice operation and returns the result tensor.
//
// tensor: The Tensor to be sliced.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Slices a tensor starting from `starts`, stopping short before `ends`
// stepping `strides` paces between each value. Semantics based on [TensorFlow
// Strided Slice Op].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceTensor(_:starts:ends:strides:startMask:endMask:squeezeMask:name:)
//
// [TensorFlow Strided Slice Op]: https://www.tensorflow.org/api_docs/python/tf/strided_slice
func (g MPSGraph) SliceTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(tensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), tensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice update operation with zero masks and returns the
// result tensor.
//
// dataTensor: The large tensor that will receive the update.
//
// updateTensor: The tensor with the new values that will replace values in the dataTensor.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:starts:ends:strides:name:)
func (g MPSGraph) SliceUpdateDataTensorUpdateTensorStartsEndsStridesName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:starts:ends:strides:name:"), dataTensor, updateTensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice update operation and returns the result tensor.
//
// dataTensor: The large tensor that will receive the update.
//
// updateTensor: The tensor with the new values that will replace values in the dataTensor.
//
// starts: An array of numbers that specify the starting points for each dimension.
//
// ends: An array of numbers that specify the ending points for each dimension.
//
// strides: An array of numbers that specify the strides for each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:starts:ends:strides:startMask:endMask:squeezeMask:name:)
func (g MPSGraph) SliceUpdateDataTensorUpdateTensorStartsEndsStridesStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, starts []foundation.NSNumber, ends []foundation.NSNumber, strides []foundation.NSNumber, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:starts:ends:strides:startMask:endMask:squeezeMask:name:"), dataTensor, updateTensor, objectivec.IObjectSliceToNSArray(starts), objectivec.IObjectSliceToNSArray(ends), objectivec.IObjectSliceToNSArray(strides), startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice update operation with zero masks and returns the
// result tensor.
//
// dataTensor: The large tensor that will receive the update.
//
// updateTensor: The tensor with the new values that will replace values in the dataTensor.
//
// startsTensor: A Tensor that contains an array of numbers that specify the starting points
// for each dimension.
//
// endsTensor: A Tensor that contains an array of numbers that specify the ending points
// for each dimension.
//
// stridesTensor: A Tensor that contains an array of numbers that specify the strides for
// each dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:startsTensor:endsTensor:stridesTensor:name:)
func (g MPSGraph) SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:startsTensor:endsTensor:stridesTensor:name:"), dataTensor, updateTensor, startsTensor, endsTensor, stridesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a strided-slice update operation and returns the result tensor.
//
// dataTensor: The large tensor that will receive the update.
//
// updateTensor: The tensor with the new values that will replace values in the dataTensor.
//
// startsTensor: A Tensor that contains an array of numbers that specify the starting points
// for each dimension.
//
// endsTensor: A Tensor that contains an array of numbers that specify the ending points
// for each dimension.
//
// stridesTensor: A Tensor that contains an array of numbers that specify the strides for
// each dimension.
//
// startMask: A bitmask that indicates dimensions whose `starts` values the operation
// should ignore.
//
// endMask: A bitmask that indicates dimensions whose `ends` values the operation
// should ignore.
//
// squeezeMask: A bitmask that indicates dimensions the operation will squeeze out from the
// result.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sliceUpdateDataTensor(_:update:startsTensor:endsTensor:stridesTensor:startMask:endMask:squeezeMask:name:)
func (g MPSGraph) SliceUpdateDataTensorUpdateTensorStartsTensorEndsTensorStridesTensorStartMaskEndMaskSqueezeMaskName(dataTensor IMPSGraphTensor, updateTensor IMPSGraphTensor, startsTensor IMPSGraphTensor, endsTensor IMPSGraphTensor, stridesTensor IMPSGraphTensor, startMask uint32, endMask uint32, squeezeMask uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sliceUpdateDataTensor:updateTensor:startsTensor:endsTensor:stridesTensor:startMask:endMask:squeezeMask:name:"), dataTensor, updateTensor, startsTensor, endsTensor, stridesTensor, startMask, endMask, squeezeMask, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the softmax function on the input tensor along the specified axis.
//
// tensor: The input tensor.
//
// axis: The axis along which softmax is computed.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMax(with:axis:name:)
func (g MPSGraph) SoftMaxWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("softMaxWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a softmax cross-entropy loss operation and returns the result
// tensor.
//
// sourceTensor: The source tensor.
//
// labelsTensor: The labels tensor.
//
// axis: The axis over which the operation computes the softmax reduction.
//
// reductionType: The type of reduction MPSGraph uses to reduce across all other axes than
// `axis`. See: [MPSGraphLossReductionType].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// The softmax cross-entropy operation computes:
//
// the operation performs the reduction over the `axis` dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxCrossEntropy(_:labels:axis:reuctionType:name:)
//
// [MPSGraphLossReductionType]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType
func (g MPSGraph) SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType MPSGraphLossReductionType, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("softMaxCrossEntropyWithSourceTensor:labelsTensor:axis:reductionType:name:"), sourceTensor, labelsTensor, axis, reductionType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates the gradient of a softmax cross-entropy loss operation and returns
// the result tensor.
//
// gradientTensor: The input gradientTensor. Note: in most cases this is the initial gradient
// tensor, which is a constant tensor with value one.
//
// sourceTensor: The source tensor.
//
// labelsTensor: The labels tensor.
//
// axis: The axis over which the operation computes the softmax reduction.
//
// reductionType: The type of reduction MPSGraph uses to reduce across all other axes than
// `axis`. See: [MPSGraphLossReductionType].
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxCrossEntropyGradient(_:source:labels:axis:reuctionType:name:)
//
// [MPSGraphLossReductionType]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLossReductionType
func (g MPSGraph) SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, labelsTensor IMPSGraphTensor, axis int, reductionType MPSGraphLossReductionType, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("softMaxCrossEntropyGradientWithIncomingGradientTensor:sourceTensor:labelsTensor:axis:reductionType:name:"), gradientTensor, sourceTensor, labelsTensor, axis, reductionType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Computes the gradient of the softmax function along the specified axis
// using the incoming gradient tensor.
//
// gradient: The incoming gradient tensor.
//
// source: The input tensor.
//
// axis: The axis along which softmax is computed.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/softMaxGradient(withIncomingGradient:sourceTensor:axis:name:)
func (g MPSGraph) SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("softMaxGradientWithIncomingGradient:sourceTensor:axis:name:"), gradient, source, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Sorts the elements of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension over which you sort the tensor
//
// descending: If true, reverse the sort direction
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axis:descending:name:)
func (g MPSGraph) SortWithTensorAxisDescendingName(tensor IMPSGraphTensor, axis int, descending bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sortWithTensor:axis:descending:name:"), tensor, axis, descending, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Sorts the elements of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axis: The tensor dimension over which you sort the tensor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axis:name:)
func (g MPSGraph) SortWithTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sortWithTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Sorts the elements of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension over which you sort the tensor
//
// descending: If true, reverse the sort direction
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axisTensor:descending:name:)
func (g MPSGraph) SortWithTensorAxisTensorDescendingName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, descending bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sortWithTensor:axisTensor:descending:name:"), tensor, axisTensor, descending, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Sorts the elements of the input tensor along the specified axis.
//
// tensor: The input tensor
//
// axisTensor: The tensor dimension over which you sort the tensor
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sort(_:axisTensor:name:)
func (g MPSGraph) SortWithTensorAxisTensorName(tensor IMPSGraphTensor, axisTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sortWithTensor:axisTensor:name:"), tensor, axisTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a space-to-depth2D operation and returns the result tensor.
//
// tensor: The input tensor.
//
// widthAxis: The axis that defines the fastest running dimension within the block.
//
// heightAxis: The axis that defines the 2nd fastest running dimension within the block.
//
// depthAxis: The axis that defines the destination dimension, where to copy the blocks.
//
// blockSize: The size of the square spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls the layout of the sub-blocks within the depth
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// This operation outputs a copy of the `input` tensor, where values from the
// `widthAxis` and `heightAxis` dimensions are moved in spatial blocks of size
// `blockSize` to the `depthAxis` dimension. Use the `usePixelShuffleOrder`
// parameter to control how the data within spatial blocks is ordered in the
// `depthAxis` dimension: with `usePixelShuffleOrder=YES` MPSGraph stores the
// values of the spatial blocks contiguosly within the `depthAxis` dimension,
// whereas otherwise they are stored interleaved with existing values in the
// `depthAxis` dimension. This operation is the inverse of
// `MPSGraph/depthToSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/space(toDepth2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:)
func (g MPSGraph) SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("spaceToDepth2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a space-to-depth2D operation and returns the result tensor.
//
// tensor: The input tensor.
//
// widthAxisTensor: A scalar tensor that contains the axis that defines the fastest running
// dimension within the block.
//
// heightAxisTensor: A scalar tensor that contains the axis that defines the 2nd fastest running
// dimension within the block.
//
// depthAxisTensor: A scalar tensor that contains the axis that defines the destination
// dimension, where to copy the blocks.
//
// blockSize: The size of the square spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls the layout of the sub-blocks within the depth
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// This operation outputs a copy of the `input` tensor, where values from the
// `widthAxisTensor` and `heightAxisTensor` dimensions are moved in spatial
// blocks of size `blockSize` to the `depthAxisTensor` dimension. Use the
// `usePixelShuffleOrder` parameter to control how the data within spatial
// blocks is ordered in the `depthAxisTensor` dimension: with
// `usePixelShuffleOrder=YES` MPSGraph stores the values of the spatial blocks
// contiguosly within the `depthAxisTensor` dimension, whereas otherwise they
// are stored interleaved with existing values in the `depthAxisTensor`
// dimension. This operation is the inverse of
// [MPSGraph.DepthToSpace2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/space(toDepth2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:)
func (g MPSGraph) SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName(tensor IMPSGraphTensor, widthAxisTensor IMPSGraphTensor, heightAxisTensor IMPSGraphTensor, depthAxisTensor IMPSGraphTensor, blockSize uint, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("spaceToDepth2DTensor:widthAxisTensor:heightAxisTensor:depthAxisTensor:blockSize:usePixelShuffleOrder:name:"), tensor, widthAxisTensor, heightAxisTensor, depthAxisTensor, blockSize, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a space-to-batch operation and returns the result tensor.
//
// tensor: The input tensor.
//
// spatialAxes: The axes that define the dimensions containing the spatial blocks.
//
// batchAxis: The axis that defines the destination dimension, where to copy the blocks.
//
// blockDimensions: An array of numbers that defines the size of the rectangular spatial
// sub-block.
//
// usePixelShuffleOrder: A parameter that controls layout of the sub-blocks within the batch
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the `input` tensor, where values from the
// `spatialAxes` (for `usePixelShuffleOrder=YES` 1,2 or 3 axes supported,
// otherwise limited only by [MPSNDArray] rank limitations) dimensions are
// moved in spatial blocks with rectangular size defined by `blockDimensions`
// to the `batchAxis` dimension. Use the `usePixelShuffleOrder` parameter to
// control how the data within spatial blocks is ordered in the `batchAxis`
// dimension: with `usePixelShuffleOrder=YES` MPSGraph stores the values of
// the spatial blocks contiguosly within the `batchAxis` dimension, whereas
// otherwise they are stored interleaved with existing values in the
// `batchAxis` dimension. Note: This operation is the inverse of
// [MPSGraph.BatchToSpaceTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName].
// Note: This operation is a generalization of
// [MPSGraph.SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/spaceToBatch(_:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:)
func (g MPSGraph) SpaceToBatchTensorSpatialAxesBatchAxisBlockDimensionsUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxes []foundation.NSNumber, batchAxis int, blockDimensions []foundation.NSNumber, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("spaceToBatchTensor:spatialAxes:batchAxis:blockDimensions:usePixelShuffleOrder:name:"), tensor, objectivec.IObjectSliceToNSArray(spatialAxes), batchAxis, objectivec.IObjectSliceToNSArray(blockDimensions), usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a space-to-batch operation and returns the result tensor.
//
// tensor: The input tensor.
//
// spatialAxesTensor: A tensor that contains the axes that define the dimensions containing the
// spatial blocks.
//
// batchAxisTensor: A tensor that contains the axis that defines the destination dimension,
// where to copy the blocks.
//
// blockDimensionsTensor: A tensor that defines the size of the rectangular spatial sub-block.
//
// usePixelShuffleOrder: A parameter that controls layout of the sub-blocks within the batch
// dimension.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// This operation outputs a copy of the `input` tensor, where values from the
// `spatialAxesTensor` (for `usePixelShuffleOrder=YES` 1,2 or 3 axes
// supported, otherwise limited only by [MPSNDArray] rank limitations)
// dimensions are moved in spatial blocks with rectangular size defined by
// `blockDimensionsTensor` to the `batchAxisTensor` dimension. Use the
// `usePixelShuffleOrder` parameter to control how the data within spatial
// blocks is ordered in the `batchAxisTensor` dimension: with
// `usePixelShuffleOrder=YES` MPSGraph stores the values of the spatial blocks
// contiguosly within the `batchAxisTensor` dimension, whereas otherwise they
// are stored interleaved with existing values in the `batchAxisTensor`
// dimension. Note: This operation is the inverse of
// [MPSGraph.BatchToSpaceTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName].
// Note: This operation is a generalization of
// [MPSGraph.SpaceToDepth2DTensorWidthAxisTensorHeightAxisTensorDepthAxisTensorBlockSizeUsePixelShuffleOrderName].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/spaceToBatch(_:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:)
func (g MPSGraph) SpaceToBatchTensorSpatialAxesTensorBatchAxisTensorBlockDimensionsTensorUsePixelShuffleOrderName(tensor IMPSGraphTensor, spatialAxesTensor IMPSGraphTensor, batchAxisTensor IMPSGraphTensor, blockDimensionsTensor IMPSGraphTensor, usePixelShuffleOrder bool, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("spaceToBatchTensor:spatialAxesTensor:batchAxisTensor:blockDimensionsTensor:usePixelShuffleOrder:name:"), tensor, spatialAxesTensor, batchAxisTensor, blockDimensionsTensor, usePixelShuffleOrder, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a sparse tensor representation.
//
// sparseDescriptor: A sparseDescriptor.
//
// inputTensorArray: An array of input tensors as [sparseVals, indexTensor0, indexTensor1].
//
// shape: The shape of the sparse tensor.
//
// name: A name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object
//
// # Discussion
//
// sparseVals corresponds to non zero values in matrix. indexTensor0 and
// indexTensor1 are indices used for indexing into sparse data structure. For
// COO, indexTensor0 is x index and indexTensor1 is y index . For CSC,
// indexTensor0 and indexTensor1 correspond to rowIndex and colStarts
// respectively. For CSR, indexTensor0 and indexTensor1 correspond to colIndex
// and rowStarts respectively. You must set input tensors appropriately for
// each sparse storage type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sparseTensor(sparseTensorWithDescriptor:tensors:shape:name:)
func (g MPSGraph) SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor IMPSGraphCreateSparseOpDescriptor, inputTensorArray []MPSGraphTensor, shape foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sparseTensorWithDescriptor:tensors:shape:name:"), sparseDescriptor, objectivec.IObjectSliceToNSArray(inputTensorArray), shape, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a sparse tensor representation.
//
// sparseStorageType: A sparseStorageType.
//
// inputTensorArray: An array of input tensors as [sparseVals, indexTensor0, indexTensor1].
//
// shape: The shape of the sparse tensor.
//
// dataType: The dataType of the sparse tensor.
//
// name: A name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// # Discussion
//
// sparseVals corresponds to non zero values in matrix. indexTensor0 and
// indexTensor1 are indices used for indexing into sparse data structure. For
// COO, indexTensor0 is x index and indexTensor1 is y index. For CSC,
// indexTensor0 and indexTensor1 correspond to rowIndex and colStarts
// respectively. For CSR, indexTensor0 and indexTensor1 correspond to colIndex
// and rowStarts respectively. You must set input tensors appropriately for
// each sparse storage type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/sparseTensor(sparseTensorWithType:tensors:shape:dataType:name:)
func (g MPSGraph) SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType MPSGraphSparseStorageType, inputTensorArray []MPSGraphTensor, shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sparseTensorWithType:tensors:shape:dataType:name:"), sparseStorageType, objectivec.IObjectSliceToNSArray(inputTensorArray), shape, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a split operation and returns the result tensor.
//
// tensor: The input tensor.
//
// numSplits: The number of result tensors to split to.
//
// axis: The dimension along which MPSGraph splits the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Splits the input tensor along `axis` into `numsplits` result tensors of
// equal size. Requires that the lenth of the input along `axis` is divisible
// by `numSplits`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:numSplits:axis:name:)
func (g MPSGraph) SplitTensorNumSplitsAxisName(tensor IMPSGraphTensor, numSplits uint, axis int, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("splitTensor:numSplits:axis:name:"), tensor, numSplits, axis, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a split operation and returns the result tensor.
//
// tensor: The input tensor.
//
// splitSizes: The lengths of the result tensors along the split axis.
//
// axis: The dimension along which MPSGraph splits the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Splits the input tensor along `axis` into multiple result tensors of size
// determined by `splitSizes`. Requires that the sum of `splitSizes` is equal
// to the lenth of the input along `axis`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:splitSizes:axis:name:)
func (g MPSGraph) SplitTensorSplitSizesAxisName(tensor IMPSGraphTensor, splitSizes []foundation.NSNumber, axis int, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("splitTensor:splitSizes:axis:name:"), tensor, objectivec.IObjectSliceToNSArray(splitSizes), axis, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a split operation and returns the result tensor.
//
// tensor: The input tensor
//
// splitSizesTensor: The lengths of the result tensors along the split axis.
//
// axis: The dimension along which MPSGraph splits the input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Splits the input tensor along `axis` into multiple result tensors of size
// determined by `splitSizesTensor`. Requires that the sum of the elements of
// `splitSizesTensor` is equal to the lenth of the input along `axis`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/split(_:splitSizesTensor:axis:name:)
func (g MPSGraph) SplitTensorSplitSizesTensorAxisName(tensor IMPSGraphTensor, splitSizesTensor IMPSGraphTensor, axis int, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("splitTensor:splitSizesTensor:axis:name:"), tensor, splitSizesTensor, axis, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Applies the square operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/square(with:name:)
func (g MPSGraph) SquareWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squareWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the square root operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squareRoot(with:name:)
func (g MPSGraph) SquareRootWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squareRootWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a squeeze operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axes: The axes to squeeze.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Squeezes the tensor, removing dimensions with size 1 at specified axes. The
// size of the input tensor must be 1 at all specified axes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axes:name:)
func (g MPSGraph) SqueezeTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squeezeTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a squeeze operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axesTensor: The tensor containing the axes to squeeze.
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor object
//
// # Discussion
//
// Squeezes the tensor, removing dimensions with size 1 at specified axes. The
// size of the input tensor must be 1 at all specified axes.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axesTensor:name:)
func (g MPSGraph) SqueezeTensorAxesTensorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squeezeTensor:axesTensor:name:"), tensor, axesTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a squeeze operation and returns the result tensor.
//
// tensor: The input tensor.
//
// axis: The axis to squeeze.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Squeezes the tensor, removing a dimension with size 1 at the specified
// axis. The size of the input tensor must be 1 at the specified axis.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:axis:name:)
func (g MPSGraph) SqueezeTensorAxisName(tensor IMPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squeezeTensor:axis:name:"), tensor, axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a squeeze operation and returns the result tensor.
//
// tensor: The input tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Squeezes the tensor, removing all dimensions with size 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/squeeze(_:name:)
func (g MPSGraph) SqueezeTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("squeezeTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a stack operation and returns the result tensor.
//
// inputTensors: The input tensors.
//
// axis: The dimension to stack tensors into result. Must be in range: `-rank + 1 <=
// dimension < rank + 1`.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Stacks all input tensors along `axis` into a result tensor of `rank + 1`.
// Tensors must be broadcast compatible along all dimensions except `axis`,
// and have the same type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stack(_:axis:name:)
func (g MPSGraph) StackTensorsAxisName(inputTensors []MPSGraphTensor, axis int, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stackTensors:axis:name:"), objectivec.IObjectSliceToNSArray(inputTensors), axis, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a stencil operation and returns the result tensor.
//
// source: The tensor containing the source data. Must be of rank 4 or greater.
//
// weights: A 4-D tensor containing the weights data.
//
// descriptor: The descriptor object that specifies the parameters for the stencil
// operation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Performs a weighted reduction operation (See
// [MPSGraphStencilOpDescriptor.ReductionMode]) on the last 4 dimensions of
// the `source` over the window determined by `weights`, according to the
// value defined in `descriptor`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stencil(withSourceTensor:weightsTensor:descriptor:name:)
func (g MPSGraph) StencilWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphStencilOpDescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stencilWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// The Stochastic gradient descent performs a gradient descent.
//
// learningRateTensor: Scalar tensor which indicates the learning rate to use with the optimizer
//
// valuesTensor: Values tensor, usually representing the trainable parameters
//
// gradientTensor: Partial gradient of the trainable parameters with respect to loss
//
// name: Name for the operation
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// `variable = variable - (learningRate * g)` where, `g` is gradient of error
// wrt variable
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/stochasticGradientDescent(learningRate:values:gradient:name:)
func (g MPSGraph) StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor IMPSGraphTensor, valuesTensor IMPSGraphTensor, gradientTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stochasticGradientDescentWithLearningRateTensor:valuesTensor:gradientTensor:name:"), learningRateTensor, valuesTensor, gradientTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Subtracts the second input tensor from the first.
//
// primaryTensor: The LHS tensor of the binary Op.
//
// secondaryTensor: The RHS tensor of the binary Op.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation creates a subtract operation and returns the result tensor.
// It supports broadcasting as well.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/subtraction(_:_:name:)
func (g MPSGraph) SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor IMPSGraphTensor, secondaryTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("subtractionWithPrimaryTensor:secondaryTensor:name:"), primaryTensor, secondaryTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the tangent operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tan(with:name:)
func (g MPSGraph) TanWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("tanWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the hyperbolic tangent operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tanh(with:name:)
func (g MPSGraph) TanhWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("tanhWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a tile gradient operation and returns the result tensor.
//
// incomingGradientTensor: The input gradient tensor.
//
// sourceTensor: The input tensor of the forward pass.
//
// multiplier: An array of numbers that specifies how many copies per dimension MPSGraph
// produced in the forward pass.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tileGradient(withIncomingGradientTensor:sourceTensor:withMultiplier:name:)
func (g MPSGraph) TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor IMPSGraphTensor, sourceTensor IMPSGraphTensor, multiplier foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("tileGradientWithIncomingGradientTensor:sourceTensor:withMultiplier:name:"), incomingGradientTensor, sourceTensor, multiplier, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a tile operation and returns the result tensor.
//
// tensor: The input tensor
//
// multiplier: An array of numbers that specifies how many copies per dimension MPSGraph
// produces.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Creates a tensor which contains multiple copies of the input tensor along
// each dimension of the tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/tileTensor(_:withMultiplier:name:)
func (g MPSGraph) TileTensorWithMultiplierName(tensor IMPSGraphTensor, multiplier foundation.NSArray, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("tileTensor:withMultiplier:name:"), tensor, multiplier, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a TopK operation and returns the value and indices tensors.
//
// source: Tensor containing source data.
//
// axis: The dimension along which to compute the TopK values.
//
// k: The number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 2.
//
// # Discussion
//
// Finds the k largest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the top values, and the second
// array corresponds to the indices of the top values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:axis:k:name:)
func (g MPSGraph) TopKWithSourceTensorAxisKName(source IMPSGraphTensor, axis int, k uint, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("topKWithSourceTensor:axis:k:name:"), source, axis, k, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a TopK operation and returns the result tensor.
//
// source: Tensor containing source data.
//
// axisTensor: Tensor containing the dimension along which to compute the TopK values.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 2.
//
// # Discussion
//
// Finds the k largest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the top values, and the second
// array corresponds to the indices of the top values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:axisTensor:kTensor:name:)
func (g MPSGraph) TopKWithSourceTensorAxisTensorKTensorName(source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("topKWithSourceTensor:axisTensor:kTensor:name:"), source, axisTensor, kTensor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a TopK operation and returns the value and indices tensors
//
// source: Tensor containing source data
//
// k: The number of largest values to return
//
// name: The name for the operation.
//
// # Return Value
//
// # A valid MPSGraphTensor array of size 2
//
// # Discussion
//
// Finds the k largest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the top values, and the second
// element of the result array corresponds to the indices of the top values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:k:name:)
func (g MPSGraph) TopKWithSourceTensorKName(source IMPSGraphTensor, k uint, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("topKWithSourceTensor:k:name:"), source, k, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a TopK operation and returns the result tensor.
//
// source: Tensor containing source data.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor array of size 2.
//
// # Discussion
//
// Finds the k largest values along the minor dimension of the input. The
// source must have at least k elements along its minor dimension. The first
// element of the result array corresponds to the top values, and the second
// element of the result array corresponds to the indices of the top values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topK(_:kTensor:name:)
func (g MPSGraph) TopKWithSourceTensorKTensorName(source IMPSGraphTensor, kTensor IMPSGraphTensor, name string) []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("topKWithSourceTensor:kTensor:name:"), source, kTensor, objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Creates a TopKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// k: The number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K largest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:input:k:name:)
func (g MPSGraph) TopKWithGradientTensorSourceKName(gradient IMPSGraphTensor, source IMPSGraphTensor, k uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("topKWithGradientTensor:source:k:name:"), gradient, source, k, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a TopKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K largest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:input:kTensor:name:)
func (g MPSGraph) TopKWithGradientTensorSourceKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("topKWithGradientTensor:source:kTensor:name:"), gradient, source, kTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a TopKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// axis: The dimension along which to compute the TopK values..
//
// k: The number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K largest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:source:axis:k:name:)
func (g MPSGraph) TopKWithGradientTensorSourceAxisKName(gradient IMPSGraphTensor, source IMPSGraphTensor, axis int, k uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("topKWithGradientTensor:source:axis:k:name:"), gradient, source, axis, k, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a TopKGradient operation and returns the result tensor.
//
// gradient: Tensor containing the incoming gradient.
//
// source: Tensor containing source data.
//
// axisTensor: Tensor containing the dimension along which to compute the TopK values.
//
// kTensor: Tensor of the number of largest values to return.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Finds the K largest values along the minor dimension of the input. The
// input must have at least K elements along its minor dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/topKGradient(_:source:axisTensor:kTensor:name:)
func (g MPSGraph) TopKWithGradientTensorSourceAxisTensorKTensorName(gradient IMPSGraphTensor, source IMPSGraphTensor, axisTensor IMPSGraphTensor, kTensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("topKWithGradientTensor:source:axisTensor:kTensor:name:"), gradient, source, axisTensor, kTensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a permutation operation and returns the result tensor.
//
// tensor: The tensor to be permuted.
//
// permutation: An array of numbers defining the permutation, must be of length
// `rank(tensor)` and define a valid permutation.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Permutes the dimensions of the input tensor according to values in
// `permutation`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/transpose(_:permutation:name:)
func (g MPSGraph) TransposeTensorPermutationName(tensor IMPSGraphTensor, permutation []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("transposeTensor:permutation:name:"), tensor, objectivec.IObjectSliceToNSArray(permutation), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a transpose operation and returns the result tensor.
//
// tensor: The tensor to be transposed.
//
// dimensionIndex: The first dimension index to be transposed.
//
// dimensionIndex2: The second dimension index to be transposed.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// Transposes the dimensions `dimensionIndex` and `dimensionIndex2` of the
// input tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/transposeTensor(_:dimension:withDimension:name:)
func (g MPSGraph) TransposeTensorDimensionWithDimensionName(tensor IMPSGraphTensor, dimensionIndex uint, dimensionIndex2 uint, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("transposeTensor:dimension:withDimension:name:"), tensor, dimensionIndex, dimensionIndex2, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Applies the truncate operation to the input tensor elements.
//
// tensor: The input tensor.
//
// name: An optional string which serves as an identifier for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object containing the elementwise result of the
// applied operation.
//
// # Discussion
//
// This operation applies the floor operation to positive inputs and ceiling
// operation to negative inputs.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/truncate(_:name:)
func (g MPSGraph) TruncateWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("truncateWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a variable operation and returns the result tensor.
//
// data: The data for the tensor. The number of bytes should be
// sizeof(dataType)numberOfElements.
//
// shape: The shape of the output tensor. This has to be statically shaped.
//
// dataType: The dataType of the constant tensor.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variable(with:shape:dataType:name:)
func (g MPSGraph) VariableWithDataShapeDataTypeName(data foundation.NSData, shape foundation.NSArray, dataType uint32, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("variableWithData:shape:dataType:name:"), data, shape, dataType, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a variable from an input tensor.
//
// tensor: The tensor from which to form the variable.
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variableFromTensor(_:name:)
func (g MPSGraph) VariableFromTensorWithTensorName(tensor IMPSGraphTensor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("variableFromTensorWithTensor:name:"), tensor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the variance of the first input along the specified axes.
//
// axes: A list of axes over which to perform the reduction. Tthe order of
// dimensions goes from the slowest moving at axis=0 to the fastest moving
// dimension.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variance(of:axes:name:)
func (g MPSGraph) VarianceOfTensorAxesName(tensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("varianceOfTensor:axes:name:"), tensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Returns the variance of the first input along the specified axes when the
// mean has been precomputed.
//
// axes: A list of axes over which to perform the reduction such that the order of
// dimensions goes from the slowest moving at axis=0 to the fastest moving
// dimension.
//
// name: An optional name for the operation.
//
// # Return Value
//
// A valid [MPSGraphTensor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/variance(of:mean:axes:name:)
func (g MPSGraph) VarianceOfTensorMeanTensorAxesName(tensor IMPSGraphTensor, meanTensor IMPSGraphTensor, axes []foundation.NSNumber, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("varianceOfTensor:meanTensor:axes:name:"), tensor, meanTensor, objectivec.IObjectSliceToNSArray(axes), objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Creates a scaled dot product attention (SDPA) operation using a descriptor
// and returns the result tensor.
//
// queryTensor: A tensor that represents the query projection.
//
// keyTensor: A tensor that represents the key projection.
//
// valueTensor: A tensor that represents the value projection.
//
// descriptor: A descriptor specifying scale and optional features (mask, isCausal,
// sinks).
//
// name: The name for the operation.
//
// # Return Value
//
// A valid MPSGraphTensor object.
//
// # Discussion
//
// The descriptor allows configuring an optional attention mask, causal
// masking, and attention sinks without requiring a separate API method for
// each combination of features.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/scaledDotProductAttention(query:key:value:descriptor:name:)
func (g MPSGraph) ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorDescriptorName(queryTensor IMPSGraphTensor, keyTensor IMPSGraphTensor, valueTensor IMPSGraphTensor, descriptor IMPSGraphSDPADescriptor, name string) IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scaledDotProductAttentionWithQueryTensor:keyTensor:valueTensor:descriptor:name:"), queryTensor, keyTensor, valueTensor, descriptor, objc.String(name))
	return MPSGraphTensorFromID(rv)
}

// Options for the graph.
//
// # Discussion
//
// The default value is [MPSGraphOptionsDefault].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/options
func (g MPSGraph) Options() MPSGraphOptions {
	rv := objc.Send[MPSGraphOptions](g.ID, objc.Sel("options"))
	return MPSGraphOptions(rv)
}
func (g MPSGraph) SetOptions(value MPSGraphOptions) {
	objc.Send[struct{}](g.ID, objc.Sel("setOptions:"), value)
}

// Array of all the placeholder tensors.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/placeholderTensors
func (g MPSGraph) PlaceholderTensors() []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("placeholderTensors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}
