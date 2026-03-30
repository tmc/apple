// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// BNNSActivation - A set of parameters that describe common activation functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSActivation
type BNNSActivation struct {
	Function            BNNSActivationFunction // The activation function that the layer applies to its output.
	Alpha               float32                // The parameter for the alpha of the activation function.
	Beta                float32                // The parameter for the beta of the activation function.
	Iscale              int32                  // Scale for integer functions.
	Ioffset             int32                  // Offset for integer functions.
	Ishift              int32                  // Shift for integer functions.
	Iscale_per_channel  *int32                 // Scale per channel for integer functions.
	Ioffset_per_channel *int32                 // Offset per channel for integer functions.
	Ishift_per_channel  *int32                 // Shift per channel for integer functions.

}

// BNNSArithmeticBinary - A structure that contains the inputs and output of an arithmetic operation with two inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticBinary
type BNNSArithmeticBinary struct {
	In1      BNNSNDArrayDescriptor // The descriptor of the first input.
	In1_type unsafe.Pointer        // The descriptor type of the first input.
	In2      BNNSNDArrayDescriptor // The descriptor of the second input.
	In2_type unsafe.Pointer        // The descriptor type of the second input.
	Out      BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer        // The descriptor type of the output.

}

// BNNSArithmeticTernary - A structure that contains the inputs and output of an arithmetic operation with three inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticTernary
type BNNSArithmeticTernary struct {
	In1      BNNSNDArrayDescriptor // The descriptor of the first input.
	In1_type unsafe.Pointer        // The descriptor type of the first input.
	In2      BNNSNDArrayDescriptor // The descriptor of the second input.
	In2_type unsafe.Pointer        // The descriptor type of the second input.
	In3      BNNSNDArrayDescriptor // The descriptor of the third input.
	In3_type unsafe.Pointer        // The descriptor type of the third input.
	Out      BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer        // The descriptor type of the output.

}

// BNNSArithmeticUnary - A structure that contains the input and output of an arithmetic operation with a single input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticUnary
type BNNSArithmeticUnary struct {
	In       BNNSNDArrayDescriptor // The descriptor of the input.
	In_type  unsafe.Pointer        // The descriptor type of the input.
	Out      BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer        // The descriptor type of the output.

}

// BNNSConvolutionLayerParameters - A structure containing convolution parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSConvolutionLayerParameters
type BNNSConvolutionLayerParameters struct {
	Activation   BNNSActivation // The layer activation function.
	Bias         BNNSLayerData  // Layer bias, one for each output channel.
	In_channels  uintptr        // The number of input channels.
	K_height     uintptr        // The height of the convolution kernel.
	K_width      uintptr        // The width of the convolution kernel.
	Out_channels uintptr        // The number of output channels.
	Weights      BNNSLayerData  // Convolution weights.
	X_padding    uintptr        // The X padding.
	X_stride     uintptr        // The X increment in the input image.
	Y_padding    uintptr        // The Y padding.
	Y_stride     uintptr        // The Y increment in the input image.

}

// BNNSFilterParameters - A structure that contains common filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterParameters
type BNNSFilterParameters struct {
	Flags        uint32    // A logical OR of zero or more values from BNNS flags.
	N_threads    uintptr   // The number of worker threads to execute.
	Alloc_memory BNNSAlloc // The function called to allocate memory.
	Free_memory  BNNSFree  // The function called to deallocate memory.

}

// BNNSFullyConnectedLayerParameters - A structure containing fully connected layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFullyConnectedLayerParameters
type BNNSFullyConnectedLayerParameters struct {
	Activation BNNSActivation // The layer activation function.
	Bias       BNNSLayerData  // Layer bias, one for each output component.
	In_size    uintptr        // The size of the input vector.
	Out_size   uintptr        // The size of the output vector.
	Weights    BNNSLayerData  // Matrix coefficients.

}

// BNNSImageStackDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSImageStackDescriptor
type BNNSImageStackDescriptor struct {
	Channels     uintptr
	Data_bias    float32
	Data_scale   float32
	Data_type    BNNSDataType
	Height       uintptr
	Image_stride uintptr
	Row_stride   uintptr
	Width        uintptr
}

// BNNSLSTMDataDescriptor - A structure that contains the input-output, hidden, and cell state n-dimensional array descriptors for a long short-term memory (LSTM) layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLSTMDataDescriptor
type BNNSLSTMDataDescriptor struct {
	Data_desc       BNNSNDArrayDescriptor // The descriptor of the input-output.
	Hidden_desc     BNNSNDArrayDescriptor // The descriptor of the hidden input-output.
	Cell_state_desc BNNSNDArrayDescriptor // The descriptor of the cell-state input-output.

}

// BNNSLSTMGateDescriptor - A structure that describes a long short-term memory (LSTM) gate layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLSTMGateDescriptor
type BNNSLSTMGateDescriptor struct {
	Iw_desc    BNNSNDArrayDescriptor // The descriptor of the input weights.
	Hw_desc    BNNSNDArrayDescriptor // The descriptor of the hidden weights.
	Cw_desc    BNNSNDArrayDescriptor // The descriptor of the cell weights.
	B_desc     BNNSNDArrayDescriptor // The descriptor of the bias.
	Activation BNNSActivation        // The activation function that the layer applies to the output.

}

// BNNSLayerData - A structure containing common layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerData
type BNNSLayerData struct {
	Data       unsafe.Pointer // Pointer to layer values (weights, bias), layout and size are specific to each layer.
	Data_bias  float32        // Conversion bias for values, used for integer data types only, ignored for indexed and float data types.
	Data_scale float32        // Conversion scale for values, used for integer data types only, ignored for indexed and float data types.
	Data_table []float32      // Conversion table (256 values) for indexed floating point data, used for indexed data types only.
	Data_type  BNNSDataType   // Storage data type for the values stored in data.

}

// BNNSLayerParametersActivation - A set of parameters that define an activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersActivation
type BNNSLayerParametersActivation struct {
	I_desc     BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc     BNNSNDArrayDescriptor // The descriptor of the output.
	Activation BNNSActivation        // The activation function that the layer applies to the output.
	Axis_flags uint32                // Flags that indicate axes on which to apply certain activation functions.

}

// BNNSLayerParametersArithmetic - A structure that contains the parameters of an arithmetic layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersArithmetic
type BNNSLayerParametersArithmetic struct {
	Arithmetic_function        unsafe.Pointer // The arithmetic operation of the layer.
	Arithmetic_function_fields unsafe.Pointer // A pointer to an arithmetic function field structure.
	Activation                 BNNSActivation // The activation function that the layer applies to the output.

}

// BNNSLayerParametersBroadcastMatMul - A set of parameters that define a broadcast matrix multiply layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersBroadcastMatMul
type BNNSLayerParametersBroadcastMatMul struct {
	Alpha        float32               // A value to scale the result.
	Beta         float32               // A value, that must be either 0.0 or 1.0, you use to scale the existing output before the operation adds it to the result.
	TransA       bool                  // A Boolean value that transposes the last two dimensions of matrix .
	TransB       bool                  // A Boolean value that transposes the last two dimensions of matrix .
	Quadratic    bool                  // A Boolean value that determines whether the operation multiplies matrix  by itself.
	A_is_weights bool                  // A Boolean value that determines whether to treat matrix  as weights.
	B_is_weights bool                  // A Boolean value that determines whether to treat matrix  as weights.
	IA_desc      BNNSNDArrayDescriptor // The descriptor of matrix .
	IB_desc      BNNSNDArrayDescriptor // The descriptor of matrix .
	O_desc       BNNSNDArrayDescriptor // The descriptor of the output.

}

// BNNSLayerParametersConvolution - A structure that contains the parameters of a convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersConvolution
type BNNSLayerParametersConvolution struct {
	I_desc            BNNSNDArrayDescriptor // The descriptor of the input.
	W_desc            BNNSNDArrayDescriptor // The descriptor of the weights.
	O_desc            BNNSNDArrayDescriptor // The descriptor of the output.
	Bias              BNNSNDArrayDescriptor // The bias descriptor.
	Activation        BNNSActivation        // The activation function that the layer applies to the output.
	X_stride          uintptr               // The width increment of the input image.
	Y_stride          uintptr               // The height increment of the input image.
	X_dilation_stride uintptr               // The width increment between elements in the input image during convolution.
	Y_dilation_stride uintptr               // The height increment between elements in the input image during convolution.
	X_padding         uintptr               // The width padding, which is the number of virtual zeros added to the left and right of each channel.
	Y_padding         uintptr               // The height padding, which is the number of virtual zeros added to the top and bottom of each channel.
	Groups            uintptr               // Convolution group size.
	Pad               uintptr               // Padding which is asymmetric and ignored if the width or height padding values are greater than zero.

}

// BNNSLayerParametersCropResize - A set of parameters that describe a crop-resize operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersCropResize
type BNNSLayerParametersCropResize struct {
	Normalized_coordinates bool                    // A Boolean value that specifies whether the operation treats the coordinates as normalized to `0...1`.
	Spatial_scale          float32                 // An additional spatial scale that mutliplies the bounding box coordinates.
	Extrapolation_value    float32                 // A value that the operation uses for extrapolation. Default value is `0`.
	Sampling_mode          unsafe.Pointer          // The sampling mode that the operation uses to select sample points.
	Box_coordinate_mode    unsafe.Pointer          // A constant that defines the convention that the operation uses to specify the four bounding box coordinates.
	Method                 BNNSInterpolationMethod // The interpolation method.

}

// BNNSLayerParametersDropout - A structure that contains the parameters of a dropout layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersDropout
type BNNSLayerParametersDropout struct {
	I_desc  BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc  BNNSNDArrayDescriptor // The descriptor of the output.
	Rate    float32               // The probability that the layer drops out an element or a group of elements.
	Seed    uint32                // The seed for the random number generator which is ignored if zero.
	Control uint8                 // An 8-bit bit mask that indicates the dimension of the grouping of the dropout decision.

}

// BNNSLayerParametersEmbedding - A structure that contains the parameters of an embedding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersEmbedding
type BNNSLayerParametersEmbedding struct {
	Flags       unsafe.Pointer        // A bit field for flags that specify additional behavior, such as scaling gradient by frequency.
	I_desc      BNNSNDArrayDescriptor // The signed or unsigned integer descriptor of the input.
	O_desc      BNNSNDArrayDescriptor // The descriptor of the output.
	Dictionary  BNNSNDArrayDescriptor // The descriptor of the dictionary.
	Padding_idx uintptr               // The padding index.
	Max_norm    float32               // The maximum norm.
	Norm_type   float32               // The norm type.

}

// BNNSLayerParametersFullyConnected - A structure that contains the parameters of a fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersFullyConnected
type BNNSLayerParametersFullyConnected struct {
	I_desc     BNNSNDArrayDescriptor // The descriptor of the input.
	W_desc     BNNSNDArrayDescriptor // The descriptor of the weights.
	O_desc     BNNSNDArrayDescriptor // The descriptor of the output.
	Bias       BNNSNDArrayDescriptor // The descriptor of the bias.
	Activation BNNSActivation        // The activation function that the layer applies to the output.

}

// BNNSLayerParametersGram - A set of parameters that define a Gram matrix layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersGram
type BNNSLayerParametersGram struct {
	Alpha  float32               // A value to scale the result.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.

}

// BNNSLayerParametersLSTM - A structure that contains the parameters of a long short-term memory (LSTM) layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLSTM
type BNNSLayerParametersLSTM struct {
	Input_size          uintptr                // The number of elements in the input.
	Hidden_size         uintptr                // The number of elements in the hidden state.
	Batch_size          uintptr                // The number of input and output samples.
	Num_layers          uintptr                // The number of stacked long short-term memory (LSTM) layers.
	Seq_len             uintptr                // The size of the sequential input.
	Dropout             float32                // The dropout ratio to apply between long short-term memory (LSTM) layers.
	Lstm_flags          uint32                 // Flags that control the behavior of a long short-term memory (LSTM) layer.
	Sequence_descriptor BNNSNDArrayDescriptor  // A 1D array of unsigned-integer elements that determines the batch size for each step.
	Input_descriptor    BNNSLSTMDataDescriptor // Descriptors of the input, hidden input, and cell-state input data.
	Output_descriptor   BNNSLSTMDataDescriptor // Descriptors of the output, hidden output, and cell-state output data.
	Input_gate          BNNSLSTMGateDescriptor // The descriptor of the input gate, which uses default sigmoid activation.
	Forget_gate         BNNSLSTMGateDescriptor // The descriptor of the forget gate, which uses default sigmoid activation.
	Candidate_gate      BNNSLSTMGateDescriptor // The descriptor of the candidate gate, which uses default tanh activation.
	Output_gate         BNNSLSTMGateDescriptor // The descriptor of the output gate, which uses default sigmoid activation.
	Hidden_activation   BNNSActivation         // Hidden activation function, which uses default tanh activation.

}

// BNNSLayerParametersLossBase - A structure that contains the parameters of a loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossBase
type BNNSLayerParametersLossBase struct {
	Function  BNNSLossFunction      // The function that’s used to compute loss.
	I_desc    BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc    BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction unsafe.Pointer        // The function that’s used to reduce the computed loss.

}

// BNNSLayerParametersLossHuber - A structure that contains the parameters of a Huber loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossHuber
type BNNSLayerParametersLossHuber struct {
	Function    BNNSLossFunction      // The function that’s used to compute loss.
	I_desc      BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc      BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction   unsafe.Pointer        // The function that’s used to reduce the computed loss.
	Huber_delta float32               // The boundary value that defines where Huber loss returns mean absolute error or mean square error.

}

// BNNSLayerParametersLossSigmoidCrossEntropy - A structure that contains the parameters of a sigmoid cross entropy loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossSigmoidCrossEntropy
type BNNSLayerParametersLossSigmoidCrossEntropy struct {
	Function     BNNSLossFunction      // The function that’s used to compute loss.
	I_desc       BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc       BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction    unsafe.Pointer        // The function that’s used to reduce the computed loss.
	Label_smooth float32               // A value that defines the smoothing that the loss function applies to the labels.

}

// BNNSLayerParametersLossSoftmaxCrossEntropy - A structure that contains the parameters of a softmax cross entropy loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossSoftmaxCrossEntropy
type BNNSLayerParametersLossSoftmaxCrossEntropy struct {
	Function     BNNSLossFunction      // The function that’s used to compute loss.
	I_desc       BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc       BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction    unsafe.Pointer        // The function that’s used to reduce the computed loss.
	Label_smooth float32               // A value that defines the smoothing that the loss function applies to the labels.

}

// BNNSLayerParametersLossYolo - A structure that contains the parameters of a You Only Look Once (YOLO) loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossYolo
type BNNSLayerParametersLossYolo struct {
	Function               BNNSLossFunction      // The function that’s used to compute loss.
	I_desc                 BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc                 BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction              unsafe.Pointer        // The function that’s used to reduce the computed loss (must be sum reduction for YOLO).
	Huber_delta            float32               // A value that’s interpreted as width-height loss.
	Number_of_grid_columns uintptr               // The number of columns in the grid.
	Number_of_grid_rows    uintptr               // The number of rows in the grid.
	Number_of_anchor_boxes uintptr               // The number of anchor boxes in each cell.
	Anchor_box_size        uintptr               // The size of the anchor box.
	Rescore                bool                  // A Boolean value that determines whether to rescore confidence according to prediction verus ground truth Intersection Over Union (IOU).
	Scale_xy               float32               // The value that specifies the x, y loss-scaling factor.
	Scale_wh               float32               // A Boolean value that determines whether to rescore confidence according to prediction verus ground truth Intersection Over Union (IOU).
	Scale_object           float32               // The value that specifies the object confidence loss-scaling factor.
	Scale_no_object        float32               // The value that specifies the no-object confidence scaling factor.
	Object_minimum_iou     float32               // The value that specifies intersection over union (IOU) that’s the minimum the function treats as an object.
	No_object_maximum_iou  float32               // The value that specifies intersection over union (IOU) that’s the maximum the function treats as not an object.
	Scale_classification   float32               // The value that specifies the classification scaling factor.
	Anchors_data           []float32             // Maximum IOU for treating as no object.

}

// BNNSLayerParametersMultiheadAttention - A structure that contains the parameters of a multihead attention layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersMultiheadAttention
type BNNSLayerParametersMultiheadAttention struct {
	Query           BNNSMHAProjectionParameters // A projection parameter structure that describes the query-related input parameters and projection.
	Key             BNNSMHAProjectionParameters // A projection parameter structure that describes the key-related input parameters and projection.
	Value           BNNSMHAProjectionParameters // A projection parameter structure that describes the value-related input parameters and projection.
	Add_zero_attn   bool                        // A Boolean value that, if true, adds a row of zeroes to the projected  and  inputs to the calculation.
	Key_attn_bias   BNNSNDArrayDescriptor       // A 2D tensor that’s added to the value as part of the attention calculation.
	Value_attn_bias BNNSNDArrayDescriptor       // An optional `d_value` x `num_heads` 2D tensor that’s added as part of the attention calculation.
	Output          BNNSMHAProjectionParameters // A projection parameter structure that describes the output tensor and associated projection.
	Dropout         float32                     // The seed for the dropout layer’s random number generator.
	Seed            uint32                      // A random seed for the dropout layer.

}

// BNNSLayerParametersNormalization - A structure that contains the parameters of a normalization layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersNormalization
type BNNSLayerParametersNormalization struct {
	I_desc               BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc               BNNSNDArrayDescriptor // The descriptor of the output.
	Beta_desc            BNNSNDArrayDescriptor // The descriptor of the beta or bias.
	Gamma_desc           BNNSNDArrayDescriptor // The descriptor of the gamma or scale.
	Moving_mean_desc     BNNSNDArrayDescriptor // The descriptor of the moving mean.
	Moving_variance_desc BNNSNDArrayDescriptor // The descriptor of the moving variance.
	Momentum             float32               // A value, between 0 and 1, the normalization operation uses to update the moving mean and moving variance during training.
	Epsilon              float32               // The epsilon in the computation of the standard deviation.
	Activation           BNNSActivation        // The activation function that the layer applies to the output.
	Num_groups           uintptr               // The number of groups over which the layer computes normalization statistics.
	Normalization_axis   uintptr               // The axis on which a layer normalization operation starts normalization.

}

// BNNSLayerParametersPadding - A structure that contains the parameters of a padding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPadding
type BNNSLayerParametersPadding struct {
	I_desc        BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc        BNNSNDArrayDescriptor // The descriptor of the output.
	Padding_size  uintptr               // The number of padding elements to add before and after the original data.
	Padding_mode  BNNSPaddingMode       // The mode the operation uses to pad.
	Padding_value uint32                // The value the operation uses to fill the padding area when the mode is constant.

}

// BNNSLayerParametersPermute - A structure that contains the parameters of a permute layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPermute
type BNNSLayerParametersPermute struct {
	I_desc      BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc      BNNSNDArrayDescriptor // The descriptor of the output.
	Permutation uintptr               // The tuple that defines the permutation.

}

// BNNSLayerParametersPooling - A structure that contains the parameters of a pooling layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPooling
type BNNSLayerParametersPooling struct {
	I_desc            BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc            BNNSNDArrayDescriptor // The descriptor of the output.
	Bias              BNNSNDArrayDescriptor // The descriptor of the bias.
	Activation        BNNSActivation        // The activation function that the layer applies to the output.
	Pooling_function  BNNSPoolingFunction   // The variable that specifies the pooling function.
	K_width           uintptr               // The width of the kernel.
	K_height          uintptr               // The height of the kernel.
	X_stride          uintptr               // The width increment of the input image.
	Y_stride          uintptr               // The height increment of the input image.
	X_dilation_stride uintptr               // The width increment between elements in the input image during convolution.
	Y_dilation_stride uintptr               // The height increment between elements in the input image during convolution.
	X_padding         uintptr               // The width padding, which is the number of virtual zeros added to the left and right of each channel.
	Y_padding         uintptr               // The height padding, which is the number of virtual zeros added to the top and bottom of each channel.
	Pad               uintptr               // Asymmetric padding, ignored if `x_padding` or `y_padding` are greater than zero.

}

// BNNSLayerParametersQuantization - A structure that contains the parameters of a quantization layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersQuantization
type BNNSLayerParametersQuantization struct {
	Axis_mask uintptr               // A bitmask that defines the axis  to which the function applies scale and bias.
	Function  BNNSQuantizerFunction // The quantize function.
	I_desc    BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc    BNNSNDArrayDescriptor // The descriptor of the output.
	Scale     BNNSNDArrayDescriptor // The descriptor of the scale.
	Bias      BNNSNDArrayDescriptor // The descriptor of the bias.

}

// BNNSLayerParametersReduction - A set of parameters that define a reduction layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersReduction
type BNNSLayerParametersReduction struct {
	I_desc      BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc      BNNSNDArrayDescriptor // The descriptor of the output.
	W_desc      BNNSNDArrayDescriptor // The descriptor of the weights.
	Reduce_func BNNSReduceFunction    // The variable that specifies the reduction function.
	Epsilon     float32               // A value that the operation adds to each element when computing the sum of logarithms.

}

// BNNSLayerParametersResize - A structure that contains the parameters of a resize layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersResize
type BNNSLayerParametersResize struct {
	Method        BNNSInterpolationMethod // The interpolation method for resizing.
	I_desc        BNNSNDArrayDescriptor   // The descriptor of the input.
	O_desc        BNNSNDArrayDescriptor   // The descriptor of the output.
	Align_corners bool                    // A Boolean value that specifies whether to align the corners of the upscaling grid to the center of scaling dimensions instead of to the edges.

}

// BNNSLayerParametersTensorContraction - A structure that contains the parameters of a tensor-contraction layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersTensorContraction
type BNNSLayerParametersTensorContraction struct {
	Operation *byte                 // The string that describes the operation.
	Alpha     float32               // Scaling that the operation applies to the result.
	Beta      float32               // A value, that must be either 0.0 or 1.0, you use to scale the existing output before the operation adds it to the result.
	IA_desc   BNNSNDArrayDescriptor // The descriptor of input matrix .
	IB_desc   BNNSNDArrayDescriptor // The descriptor of input matrix .
	O_desc    BNNSNDArrayDescriptor // The descriptor of the output.

}

// BNNSMHAProjectionParameters - A structure that contains multihead attention projection parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSMHAProjectionParameters
type BNNSMHAProjectionParameters struct {
	Target_desc BNNSNDArrayDescriptor // The descriptor—which is either an input query, key, or value, or an output—of the main target of the operation.
	Weights     BNNSNDArrayDescriptor // The descriptor of the initial projection’s weights.
	Bias        BNNSNDArrayDescriptor // The descriptor of the initial projection’s bias.

}

// BNNSNDArrayDescriptor - A structure that describes the shape, stride, data type, and, optionally, the memory location of an n-dimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayDescriptor
type BNNSNDArrayDescriptor struct {
	Flags           unsafe.Pointer // Flags that control some behaviors of the n-dimensional array.
	Layout          unsafe.Pointer // The dimension of the n-dimensional array.
	Data            unsafe.Pointer // A pointer that is optional and points to the underlying data.
	Data_type       BNNSDataType   // The data type of the n-dimensional array.
	Table_data      unsafe.Pointer // The lookup table for indexed data types.
	Table_data_type BNNSDataType   // The data type of the lookup table.
	Data_scale      float32        // The scale you use to convert integer and unsigned integer data to floating point.
	Data_bias       float32        // The bias you use to convert integer and unsigned integer data to floating point.
	Size            uintptr        // The number of values in each dimension.
	Stride          uintptr        // The increment, in values, between consecutive elements in each dimension.

}

// BNNSOptimizerAdamFields - A structure that contains the fields of an Adam optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerAdamFields
type BNNSOptimizerAdamFields struct {
	Learning_rate        float32        // A value that specifies the learning rate.
	Beta1                float32        // A value that specifies the first moment constant in the range 0 to 1.
	Beta2                float32        // A value that specifies the second moment constant in the range 0 to 1.
	Time_step            float32        // A value that represents the optimizer’s current time and you’re responsible for updating after optimizing all the layer parameters in your network.
	Epsilon              float32        // An addition for the division in the parameter update stage.
	Gradient_scale       float32        // A value that specifies the gradient scaling factor.
	Regularization_scale float32        // A value that specifies the regularization scaling factor.
	Clip_gradients       bool           // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_min   float32        // The values for the minimum gradient.
	Clip_gradients_max   float32        // The values for the maximum gradient.
	Regularization_func  unsafe.Pointer // The variable that specifies the regularization function.

}

// BNNSOptimizerAdamWithClippingFields - A structure that contains the fields of an Adam or AdamW optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerAdamWithClippingFields
type BNNSOptimizerAdamWithClippingFields struct {
	Learning_rate           float32        // A value that specifies the learning rate.
	Beta1                   float32        // A value that specifies the first moment constant in the range 0 to 1.
	Beta2                   float32        // A value that specifies the second moment constant in the range 0 to 1.
	Time_step               float32        // A value that’s at least 1 and represents the optimizer’s current time.
	Epsilon                 float32        // An addition for the division in the parameter update stage.
	Gradient_scale          float32        // A value that specifies the gradient scaling factor.
	Regularization_scale    float32        // A value that specifies the regularization scaling factor.
	Regularization_func     unsafe.Pointer // The variable that specifies the regularization function.
	Clipping_func           unsafe.Pointer // The clipping function.
	Clip_gradients_min      float32        // The minimum clipping value for clipping by value.
	Clip_gradients_max      float32        // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32        // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_use_norm float32        // An optional value for a known Euclidean norm for clipping by global norm.

}

// BNNSOptimizerRMSPropFields - A structure that contains the fields of a root mean square propagation (RMSProp) optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerRMSPropFields
type BNNSOptimizerRMSPropFields struct {
	Learning_rate        float32        // A value that specifies the learning rate.
	Alpha                float32        // A constant that specifies smoothing.
	Epsilon              float32        // A term that the optimizer adds to the denominator.
	Centered             bool           // A Boolean value that specifies whether to use the centered variant.
	Momentum             float32        // The rate of momentum decay.
	Gradient_scale       float32        // A value that specifies the gradient scaling factor.
	Regularization_scale float32        // A value that specifies the regularization scaling factor.
	Clip_gradients       bool           // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_min   float32        // The values for the minimum gradient.
	Clip_gradients_max   float32        // The values for the maximum gradient.
	Regularization_func  unsafe.Pointer // The variable that specifies the regularization function.

}

// BNNSOptimizerRMSPropWithClippingFields - A structure that contains the fields of a root mean square propagation (RMSProp) optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerRMSPropWithClippingFields
type BNNSOptimizerRMSPropWithClippingFields struct {
	Learning_rate           float32        // A value that specifies the learning rate.
	Alpha                   float32        // A constant that specifies smoothing.
	Epsilon                 float32        // A term that the optimizer adds to the denominator.
	Centered                bool           // A Boolean value that specifies whether to use the centered variant.
	Momentum                float32        // The rate of momentum decay.
	Gradient_scale          float32        // A value that specifies the gradient scaling factor.
	Regularization_scale    float32        // A value that specifies the regularization scaling factor.
	Regularization_func     unsafe.Pointer // The variable that specifies the regularization function.
	Clipping_func           unsafe.Pointer // The clipping function.
	Clip_gradients_min      float32        // The minimum clipping value for clipping by value.
	Clip_gradients_max      float32        // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32        // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_use_norm float32        // An optional value for a known Euclidean norm for clipping by global norm.

}

// BNNSOptimizerSGDMomentumFields - A structure that contains the fields of a stochastic gradient descent (SGD) with momentum optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerSGDMomentumFields
type BNNSOptimizerSGDMomentumFields struct {
	Learning_rate        float32        // A value that specifies the learning rate.
	Momentum             float32        // The rate of momentum decay.
	Gradient_scale       float32        // A value that specifies the gradient scaling factor.
	Regularization_scale float32        // A value that specifies the regularization scaling factor.
	Clip_gradients       bool           // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_min   float32        // The values for the minimum gradient.
	Clip_gradients_max   float32        // The values for the maximum gradient.
	Nesterov             bool           // A Boolean value that specifies whether to use Nesterov momentum update.
	Regularization_func  unsafe.Pointer // The variable that specifies the regularization function.
	Sgd_momentum_variant unsafe.Pointer // The variable that specifies the momentum variant.

}

// BNNSOptimizerSGDMomentumWithClippingFields - A structure that contains the fields of a stochastic gradient descent (SGD) with momentum optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerSGDMomentumWithClippingFields
type BNNSOptimizerSGDMomentumWithClippingFields struct {
	Learning_rate           float32        // A value that specifies the learning rate.
	Momentum                float32        // The rate of momentum decay.
	Gradient_scale          float32        // A value that specifies the gradient scaling factor.
	Regularization_scale    float32        // A value that specifies the regularization scaling factor.
	Nesterov                bool           // A Boolean value that specifies whether to use Nesterov momentum update.
	Regularization_func     unsafe.Pointer // The variable that specifies the regularization function.
	Sgd_momentum_variant    unsafe.Pointer // The variable that specifies the momentum variant.
	Clipping_func           unsafe.Pointer // The clipping function.
	Clip_gradients_min      float32        // The minimum clipping value for clipping by value.
	Clip_gradients_max      float32        // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32        // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_use_norm float32        // An optional value for a known Euclidean norm for clipping by global norm.

}

// BNNSPoolingLayerParameters - A structure containing pooling layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSPoolingLayerParameters
type BNNSPoolingLayerParameters struct {
	Activation       BNNSActivation      // The layer activation function.
	Bias             BNNSLayerData       // Layer bias, one for each output channel.
	In_channels      uintptr             // The number of input channels.
	K_height         uintptr             // The height of the convolution kernel.
	K_width          uintptr             // The width of the convolution kernel.
	Out_channels     uintptr             // The number of output channels.
	Pooling_function BNNSPoolingFunction // The pooling function to apply to each sample.
	X_padding        uintptr             // The X padding.
	X_stride         uintptr             // The X increment in the input image.
	Y_padding        uintptr             // The Y padding.
	Y_stride         uintptr             // The Y increment in the input image.

}

// BNNSSparsityParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSSparsityParameters
type BNNSSparsityParameters struct {
	Flags          uint64
	Sparsity_ratio uint32
	Sparsity_type  BNNSSparsityType
	Target_system  BNNSTargetSystem
}

// BNNSTensor - A structure that describes the shape, stride, data type, and, optionally, the memory location of an n-dimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTensor
type BNNSTensor struct {
	Data_type          BNNSDataType   // The data type of the tensor.
	Rank               uint8          // The rank of the tensor.
	Data               unsafe.Pointer // A pointer to the memory that contains the tensor values.
	Data_size_in_bytes uintptr        // The size, in bytes, of the memory that contains the tensor values.
	Name               *byte          // An optional name for the tensor that you can use for debugging.
	Shape              int            // A tuple of unsigned-integer elements that specify the size of the tensor.
	Stride             int            // A tuple of unsigned-integer elements that specify the stride of the tensor.

}

// BNNSVectorDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSVectorDescriptor
type BNNSVectorDescriptor struct {
	Data_bias  float32
	Data_scale float32
	Data_type  BNNSDataType
	Size       uintptr
}

// DSPComplex - A structure that represents a single-precision complex value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPComplex
type DSPComplex struct {
	Real float32 // The real part of the value.
	Imag float32 // The imaginary part of the value.

}

// DSPDoubleComplex - A structure that represents a double-precision complex value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPDoubleComplex
type DSPDoubleComplex struct {
	Real float64 // The real part of the value.
	Imag float64 // The imaginary part of the value.

}

// DSPDoubleSplitComplex - A structure that represents a double-precision complex vector with the real and imaginary parts stored in separate arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPDoubleSplitComplex
type DSPDoubleSplitComplex struct {
	Realp []float64 // An array of real parts of the complex numbers.
	Imagp []float64 // An array of imaginary parts of the complex numbers.

}

// DSPSplitComplex - A structure that represents a single-precision complex vector with the real and imaginary parts stored in separate arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPSplitComplex
type DSPSplitComplex struct {
	Realp []float32 // An array of real parts of the complex numbers.
	Imagp []float32 // An array of imaginary parts of the complex numbers.

}

// DenseMatrix_Complex_Double - Contains a dense `rowCount` x `columnCount` matrix of complex double values stored in column-major order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Complex_Double
type DenseMatrix_Complex_Double struct {
	RowCount     int
	ColumnCount  int
	ColumnStride int
	Attributes   SparseAttributesComplex_t // A type representing the attributes of a matrix.
	Data         objectivec.IObject
}

// DenseMatrix_Complex_Float - Contains a dense `rowCount` x `columnCount` matrix of complex float values stored in column-major order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Complex_Float
type DenseMatrix_Complex_Float struct {
	RowCount     int
	ColumnCount  int
	ColumnStride int
	Attributes   SparseAttributesComplex_t // A type representing the attributes of a matrix.
	Data         objectivec.IObject
}

// DenseMatrix_Double - A structure that contains a dense matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Double
type DenseMatrix_Double struct {
	RowCount     int                // The number of rows in the matrix.
	ColumnCount  int                // The number of columns in the matrix.
	ColumnStride int                // The stride between matrix columns, in elements.
	Attributes   SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	Data         []float64          // The array of double-precision, floating-point values in column-major order.

}

// DenseMatrix_Float - A structure that contains a dense matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Float
type DenseMatrix_Float struct {
	RowCount     int                // The number of rows in the matrix.
	ColumnCount  int                // The number of columns in the matrix.
	ColumnStride int                // The stride between matrix columns, in elements.
	Attributes   SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	Data         []float32          // The array of single-precision, floating-point values in column-major order.

}

// DenseVector_Complex_Double - Contains a dense vector of double complex values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Complex_Double
type DenseVector_Complex_Double struct {
	Count int
	Data  objectivec.IObject
}

// DenseVector_Complex_Float - Contains a dense vector of float complex values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Complex_Float
type DenseVector_Complex_Float struct {
	Count int
	Data  objectivec.IObject
}

// DenseVector_Double - A structure that contains a dense vector of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Double
type DenseVector_Double struct {
	Count int       // The number of items in the vector.
	Data  []float64 // The array of double-precision, floating-point values.

}

// DenseVector_Float - A structure that contains a dense vector of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Float
type DenseVector_Float struct {
	Count int       // The number of items in the vector.
	Data  []float32 // The array of single-precision, floating-point values.

}

// SparseAttributesComplex_t - A type representing the attributes of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseAttributesComplex_t
type SparseAttributesComplex_t struct {
	Conjugate_transpose bool
	Kind                unsafe.Pointer // A flag to describe the type of matrix represented.
	Transpose           bool
	Triangle            unsafe.Pointer // A flag to indicate which triangle of a matrix is used.

}

// SparseAttributes_t - A structure that represents the attributes of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseAttributes_t
type SparseAttributes_t struct {
	Transpose bool           // A Boolean value that specifies whether to implicitly transpose the matrix.
	Triangle  unsafe.Pointer // An enumeration that specifies which triangle unit-triangular, triangular, and symmetric matrices need to use.
	Kind      unsafe.Pointer // An eumeration that specifies whether the matrix is ordinary, unit-triangular, triangular, or symmetric.

}

// SparseCGOptions - Options for creating a conjugate gradient (CG) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseCGOptions
type SparseCGOptions struct {
	MaxIterations int         // The maximum number of iterations to perform.
	Atol          float64     // The absolute convergence tolerance.
	Rtol          float64     // The relative convergence tolerance.
	ReportError   func(*byte) // An optional error-reporting routine.
	ReportStatus  func(*byte) // The function to report status.

}

// SparseGMRESOptions - Options for creating a generalized minimal residual (GMRES) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseGMRESOptions
type SparseGMRESOptions struct {
	Variant       unsafe.Pointer // The exact variant of GMRES to implement.
	Nvec          int            // The number of orthogonal vectors the operation maintains.
	MaxIterations int            // The maximum number of iterations to perform.
	Atol          float64        // The absolute convergence tolerance.
	Rtol          float64        // The relative convergence tolerance.
	ReportError   func(*byte)    // An optional error-reporting routine.
	ReportStatus  func(*byte)    // The function to report status.

}

// SparseIterativeMethod - The base type for all iterative methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseIterativeMethod
type SparseIterativeMethod struct {
	Method  int       // The iterative method this structure represents.
	Options [256]byte // The options for the method.
	Base    uint
	Cg      SparseCGOptions    // Conjugate Gradient Options.
	Gmres   SparseGMRESOptions // Right-preconditioned (F/DQ)GMRES Parameters Options.
	Lsmr    SparseLSMROptions  // LSMR is MINRES specialised for solving least squares.
	Padding int8
}

// SparseLSMROptions - Options for creating a least squares minimum residual method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseLSMROptions
type SparseLSMROptions struct {
	Lambda          float64        // The damping parameter lambda for regularized least squares.
	Nvec            int            // The number of vectors to use for local reorthogonalization.
	ConvergenceTest unsafe.Pointer // The convergence test to use for iterative solve methods.
	Atol            float64        // The absolute tolerance (default test) or  tolerance (Fong-Saunders test).
	Rtol            float64        // The relative convergence tolerance (default test only).
	Btol            float64        // The  tolerance (Fong-Saunders test only).
	ConditionLimit  float64        // The condition number limit (Fong-Saunders test only).
	MaxIterations   int            // The maximum number of iterations.
	ReportError     func(*byte)    // An optional error-reporting routine.
	ReportStatus    func(*byte)    // An optional status-reporting routine.

}

// SparseMatrixStructure - A description of the sparsity structure of a sparse matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrixStructure
type SparseMatrixStructure struct {
	RowCount     int                // The number of rows in the matrix.
	ColumnCount  int                // The number of columns in the matrix.
	Attributes   SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	BlockSize    uint8              // The block size of the matrix.
	ColumnStarts *int               // The starting index for each column in the row indices array.
	RowIndices   []int              // The row indices of the matrix.

}

// SparseMatrixStructureComplex - A type representing the sparsity structure of a sparse complex matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrixStructureComplex
type SparseMatrixStructureComplex struct {
	RowCount     int
	ColumnCount  int
	Attributes   SparseAttributesComplex_t // A type representing the attributes of a matrix.
	BlockSize    uint8
	ColumnStarts *int
	RowIndices   []int
}

// SparseMatrix_Complex_Double - A type representing a sparse complex matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Complex_Double
type SparseMatrix_Complex_Double struct {
	Structure SparseMatrixStructureComplex // A type representing the sparsity structure of a sparse complex matrix.
	Data      objectivec.IObject
}

// SparseMatrix_Complex_Float - A type representing a sparse complex matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Complex_Float
type SparseMatrix_Complex_Float struct {
	Structure SparseMatrixStructureComplex // A type representing the sparsity structure of a sparse complex matrix.
	Data      objectivec.IObject
}

// SparseMatrix_Double - A structure that contains a sparse matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Double
type SparseMatrix_Double struct {
	Structure SparseMatrixStructure // The sparsity structure of the matrix.
	Data      []float64             // The array of contiguous values in the nonzero blocks of the matrix.

}

// SparseMatrix_Float - A structure that contains a sparse matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Float
type SparseMatrix_Float struct {
	Structure SparseMatrixStructure // The sparsity structure of the matrix.
	Data      []float32             // The array of contiguous values in the nonzero blocks of the matrix.

}

// SparseNumericFactorOptions - A structure that contains options that affect the numerical stage of a sparse factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseNumericFactorOptions
type SparseNumericFactorOptions struct {
	Control        unsafe.Pointer // The flags that control the computation.
	ScalingMethod  unsafe.Pointer // The scaling method.
	Scaling        unsafe.Pointer // An array that scales the matrix before factorization.
	PivotTolerance float64        // The pivot tolerance that threshold partial pivoting uses.
	ZeroTolerance  float64        // The zero tolerance that some pivoting modes use.

}

// SparseOpaqueFactorization_Complex_Double - A semi-opaque type representing a matrix factorization in complex double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Complex_Double
type SparseOpaqueFactorization_Complex_Double struct {
	Status                       unsafe.Pointer                    // Status field for a factorization.
	Attributes                   SparseAttributesComplex_t         // A type representing the attributes of a matrix.
	SymbolicFactorization        SparseOpaqueSymbolicFactorization // A semi-opaque type representing symbolic matrix factorization.
	UserFactorStorage            bool
	NumericFactorization         unsafe.Pointer
	SolveWorkspaceRequiredStatic uintptr
	SolveWorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueFactorization_Complex_Float - A semi-opaque type representing a matrix factorization in complex float.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Complex_Float
type SparseOpaqueFactorization_Complex_Float struct {
	Status                       unsafe.Pointer                    // Status field for a factorization.
	Attributes                   SparseAttributesComplex_t         // A type representing the attributes of a matrix.
	SymbolicFactorization        SparseOpaqueSymbolicFactorization // A semi-opaque type representing symbolic matrix factorization.
	UserFactorStorage            bool
	NumericFactorization         unsafe.Pointer
	SolveWorkspaceRequiredStatic uintptr
	SolveWorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueFactorization_Double - A structure that represents the factorization of a matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Double
type SparseOpaqueFactorization_Double struct {
	Status                       unsafe.Pointer                    // The status of the factorization object.
	Attributes                   SparseAttributes_t                // The attributes of a factorization object.
	SymbolicFactorization        SparseOpaqueSymbolicFactorization // The symbolic factorization that this numeric factorization depends on.
	UserFactorStorage            bool                              // A Boolean value that indicates whether user-provided storage backs this object.
	NumericFactorization         unsafe.Pointer                    // The pointer to a private internal representation of a numeric factor.
	SolveWorkspaceRequiredStatic uintptr                           // The required size of the static workspace for a call to a sparse solve function.
	SolveWorkspaceRequiredPerRHS uintptr                           // The required size of the per-right-hand-side workspace for a call to a sparse solve function.

}

// SparseOpaqueFactorization_Float - A structure that represents the factorization of a matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Float
type SparseOpaqueFactorization_Float struct {
	Status                       unsafe.Pointer                    // The status of the factorization object.
	Attributes                   SparseAttributes_t                // The attributes of a factorization object.
	SymbolicFactorization        SparseOpaqueSymbolicFactorization // The symbolic factorization that this numeric factorization depends on.
	UserFactorStorage            bool                              // A Boolean value that indicates whether user-provided storage backs this object.
	NumericFactorization         unsafe.Pointer                    // The pointer to a private internal representation of a numeric factor.
	SolveWorkspaceRequiredStatic uintptr                           // The required size of the static workspace for a call to a sparse solve function.
	SolveWorkspaceRequiredPerRHS uintptr                           // The required size of the per-right-hand-side workspace for a call to a sparse solve function.

}

// SparseOpaquePreconditioner_Complex_Double - Represents a preconditioner for matrices of complex double values .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Complex_Double
type SparseOpaquePreconditioner_Complex_Double struct {
	Type  unsafe.Pointer // Types of preconditioner.
	Apply func(unsafe.Pointer, CBLAS_TRANSPOSE, DenseMatrix_Complex_Double, DenseMatrix_Complex_Double)
	Mem   unsafe.Pointer
}

// SparseOpaquePreconditioner_Complex_Float - Represents a preconditioner for matrices of complex float values .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Complex_Float
type SparseOpaquePreconditioner_Complex_Float struct {
	Type  unsafe.Pointer // Types of preconditioner.
	Apply func(unsafe.Pointer, CBLAS_TRANSPOSE, DenseMatrix_Complex_Float, DenseMatrix_Complex_Float)
	Mem   unsafe.Pointer
}

// SparseOpaquePreconditioner_Double - A structure that represents a double-precision preconditioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Double
type SparseOpaquePreconditioner_Double struct {
	Type  unsafe.Pointer                                                                // The preconditioner type.
	Apply func(unsafe.Pointer, CBLAS_TRANSPOSE, DenseMatrix_Double, DenseMatrix_Double) // A function that calculates , where  is the preconditioner.
	Mem   unsafe.Pointer                                                                // The unaltered memory pointer that passes as the first parameter of the apply function.

}

// SparseOpaquePreconditioner_Float - A structure that represents a single-precision preconditioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Float
type SparseOpaquePreconditioner_Float struct {
	Type  unsafe.Pointer                                                              // The preconditioner type.
	Apply func(unsafe.Pointer, CBLAS_TRANSPOSE, DenseMatrix_Float, DenseMatrix_Float) // A function that calculates , where  is the preconditioner.
	Mem   unsafe.Pointer                                                              // The unaltered memory pointer that passes as the first parameter of the apply function.

}

// SparseOpaqueSubfactor_Complex_Double - Represents a sub-factor of the factorization (for example,  [L] from `LDL^T`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Complex_Double
type SparseOpaqueSubfactor_Complex_Double struct {
	Attributes              SparseAttributesComplex_t                // A type representing the attributes of a matrix.
	Contents                unsafe.Pointer                           // Types of sub-factor object.
	Factor                  SparseOpaqueFactorization_Complex_Double // A semi-opaque type representing a matrix factorization in complex double.
	WorkspaceRequiredStatic uintptr
	WorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueSubfactor_Complex_Float - Represents a sub-factor of the factorization (for example,  [L] from `LDL^T`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Complex_Float
type SparseOpaqueSubfactor_Complex_Float struct {
	Attributes              SparseAttributesComplex_t               // A type representing the attributes of a matrix.
	Contents                unsafe.Pointer                          // Types of sub-factor object.
	Factor                  SparseOpaqueFactorization_Complex_Float // A semi-opaque type representing a matrix factorization in complex float.
	WorkspaceRequiredStatic uintptr
	WorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueSubfactor_Double - Represents a sub-factor of the factorization (for example,  [L] from `LDL^T`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Double
type SparseOpaqueSubfactor_Double struct {
	Attributes              SparseAttributes_t               // A type representing the attributes of a matrix.
	Contents                unsafe.Pointer                   // Types of sub-factor object.
	Factor                  SparseOpaqueFactorization_Double // A semi-opaque type representing a matrix factorization in double.
	WorkspaceRequiredStatic uintptr
	WorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueSubfactor_Float - Represents a sub-factor of the factorization (for example,  [L] from `LDL^T`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Float
type SparseOpaqueSubfactor_Float struct {
	Attributes              SparseAttributes_t              // A type representing the attributes of a matrix.
	Contents                unsafe.Pointer                  // Types of sub-factor object.
	Factor                  SparseOpaqueFactorization_Float // A semi-opaque type representing a matrix factorization in float.
	WorkspaceRequiredStatic uintptr
	WorkspaceRequiredPerRHS uintptr
}

// SparseOpaqueSymbolicFactorization - A semi-opaque type that represents symbolic matrix factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSymbolicFactorization
type SparseOpaqueSymbolicFactorization struct {
	Status               unsafe.Pointer     // The status of the factorization.
	RowCount             int                // The number of rows.
	ColumnCount          int                // The number of columns.
	Attributes           SparseAttributes_t // The attributes of the factorization.
	BlockSize            uint8              // The block size.
	Type                 unsafe.Pointer     // The factorization type.
	Factorization        unsafe.Pointer     // A pointer to a private internal representation of the symbolic factor.
	WorkspaceSize_Float  uintptr            // Size, in bytes, of workspace required to perform numerical factorization in floats.
	WorkspaceSize_Double uintptr            // Size, in bytes, of workspace required to perform numerical factorization in doubles.
	FactorSize_Float     uintptr            // Minimum size, in bytes, required to store numerical factors in float.
	FactorSize_Double    uintptr            // Minimum size, in bytes, required to store numerical factors in doubles.

}

// SparseSymbolicFactorOptions - A structure that contains options that affect the symbolic stage of a sparse factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseSymbolicFactorOptions
type SparseSymbolicFactorOptions struct {
	Control              unsafe.Pointer            // The flags that control the computation.
	OrderMethod          unsafe.Pointer            // The ordering algorithm.
	Order                []int                     // The user-supplied array for ordering.
	IgnoreRowsAndColumns []int                     // An array that contains row and column indices to ignore.
	Malloc               func(uint) unsafe.Pointer // The function for allocating any necessary storage.
	Free                 func(unsafe.Pointer)      // The function for freeing allocated storage.
	ReportError          func(*byte)               // The function for reporting parameter errors.

}

// Bnns_graph_argument_t - Describes data associated with an input or output argument
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_argument_t
type Bnns_graph_argument_t struct {
	Data_ptr_size uintptr                // size in bytes of `data_ptr`, if set
	Data_ptr      unsafe.Pointer         // Direct pointer to numerical data
	Descriptor    *BNNSNDArrayDescriptor // Pointer to BNNSNDArrayDescriptor (deprecated, use BNNSTensor instead)
	Tensor        *BNNSTensor            // Pointer to BNNSTensor

}

// Bnns_graph_compile_options_t - The compilation options that BNNS uses when compiling a source mlmodelc file to a graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_compile_options_t
type Bnns_graph_compile_options_t struct {
	Data unsafe.Pointer // A pointer to the opaque compilation options object.
	Size uintptr        // The size, in bytes, of the opaque compilation options object.

}

// Bnns_graph_context_t - An object that wraps a compiled graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_context_t
type Bnns_graph_context_t struct {
	Data unsafe.Pointer // A pointer to the opaque graph context object.
	Size uintptr        // The size, in bytes, of the opaque graph context object.

}

// Bnns_graph_shape_t - The specification of the shape of an argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_shape_t
type Bnns_graph_shape_t struct {
	Rank  uintptr // The rank of the shape.
	Shape *uint64 // An array of unsigned-integer elements that specify the size of the shape.

}

// Bnns_graph_t - The compiled graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_t
type Bnns_graph_t struct {
	Data unsafe.Pointer // A pointer to opaque graph object.
	Size uintptr        // The size, in bytes, of the opaque graph object.

}

// Bnns_user_message_data_t - Additional user-defined logging argument for message-logging callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_user_message_data_t
type Bnns_user_message_data_t struct {
	Data unsafe.Pointer // A pointer to the additional logging data.
	Size uintptr        // The size of the additional logging data.

}

// Quadrature_integrate_function
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/quadrature_integrate_function
type Quadrature_integrate_function struct {
	Fun     Quadrature_function_array
	Fun_arg unsafe.Pointer
}

// Quadrature_integrate_options
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/quadrature_integrate_options
type Quadrature_integrate_options struct {
	Integrator              unsafe.Pointer
	Abs_tolerance           float64
	Rel_tolerance           float64
	Qag_points_per_interval uintptr
	Max_intervals           uintptr
}

// Simd_double2x2 - A matrix of two columns and two rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double2x2
type Simd_double2x2 struct {
	Determinant float64      // The determinant of the matrix.
	Columns     Simd_double2 // The columns of the matrix.

}

// Simd_double2x3 - A matrix of two columns and three rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double2x3
type Simd_double2x3 struct {
	Columns Simd_double3 // The columns of the matrix.

}

// Simd_double2x4 - A matrix of two columns and four rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double2x4
type Simd_double2x4 struct {
	Columns Simd_double4 // The columns of the matrix.

}

// Simd_double3x2 - A matrix of three columns and two rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double3x2
type Simd_double3x2 struct {
	Columns Simd_double2 // The columns of the matrix.

}

// Simd_double3x3 - A matrix of three columns and three rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double3x3
type Simd_double3x3 struct {
	Determinant float64      // The determinant of the matrix.
	Columns     Simd_double3 // The columns of the matrix.

}

// Simd_double3x4 - A matrix of three columns and four rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double3x4
type Simd_double3x4 struct {
	Columns Simd_double4 // The columns of the matrix.

}

// Simd_double4x2 - A matrix of four columns and two rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double4x2
type Simd_double4x2 struct {
	Columns Simd_double2 // The columns of the matrix.

}

// Simd_double4x3 - A matrix of four columns and three rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double4x3
type Simd_double4x3 struct {
	Columns Simd_double3 // The columns of the matrix.

}

// Simd_double4x4 - A matrix of four columns and four rows that contains double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_double4x4
type Simd_double4x4 struct {
	Determinant float64      // The determinant of the matrix.
	Columns     Simd_double4 // The columns of the matrix.

}

// Simd_float2x2 - A matrix of two columns and two rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float2x2
type Simd_float2x2 struct {
	Determinant float32     // The determinant of the matrix.
	Columns     Simd_float2 // The columns of the matrix.

}

// Simd_float2x3 - A matrix of two columns and three rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float2x3
type Simd_float2x3 struct {
	Columns Simd_float3 // The columns of the matrix.

}

// Simd_float2x4 - A matrix of two columns and four rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float2x4
type Simd_float2x4 struct {
	Columns Simd_float4 // The columns of the matrix.

}

// Simd_float3x2 - A matrix of three columns and two rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float3x2
type Simd_float3x2 struct {
	Columns Simd_float2 // The columns of the matrix.

}

// Simd_float3x3 - A matrix of three columns and three rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float3x3
type Simd_float3x3 struct {
	Determinant float32     // The determinant of the matrix.
	Columns     Simd_float3 // The columns of the matrix.

}

// Simd_float3x4 - A matrix of three columns and four rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float3x4
type Simd_float3x4 struct {
	Columns Simd_float4 // The columns of the matrix.

}

// Simd_float4x2 - A matrix of four columns and two rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float4x2
type Simd_float4x2 struct {
	Columns Simd_float2 // The columns of the matrix.

}

// Simd_float4x3 - A matrix of four columns and three rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float4x3
type Simd_float4x3 struct {
	Columns Simd_float3 // The columns of the matrix.

}

// Simd_float4x4 - A matrix of four columns and four rows that contains single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_float4x4
type Simd_float4x4 struct {
	Determinant float32     // The determinant of the matrix.
	Columns     Simd_float4 // The columns of the matrix.

}

// Simd_half2x2 - A matrix of two columns and two rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half2x2
type Simd_half2x2 struct {
	Columns Simd_half2 // The columns of the matrix.

}

// Simd_half2x3 - A matrix of two columns and three rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half2x3
type Simd_half2x3 struct {
	Columns Simd_half3 // The columns of the matrix.

}

// Simd_half2x4 - A matrix of two columns and four rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half2x4
type Simd_half2x4 struct {
	Columns Simd_half4 // The columns of the matrix.

}

// Simd_half3x2 - A matrix of three columns and two rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half3x2
type Simd_half3x2 struct {
	Columns Simd_half2 // The columns of the matrix.

}

// Simd_half3x3 - A matrix of three columns and three rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half3x3
type Simd_half3x3 struct {
	Columns Simd_half3 // The columns of the matrix.

}

// Simd_half3x4 - A matrix of three columns and four rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half3x4
type Simd_half3x4 struct {
	Columns Simd_half4 // The columns of the matrix.

}

// Simd_half4x2 - A matrix of four columns and two rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half4x2
type Simd_half4x2 struct {
	Columns Simd_half2 // The columns of the matrix.

}

// Simd_half4x3 - A matrix of four columns and three rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half4x3
type Simd_half4x3 struct {
	Columns Simd_half3 // The columns of the matrix.

}

// Simd_half4x4 - A matrix of four columns and four rows that contains half-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_half4x4
type Simd_half4x4 struct {
	Columns Simd_half4 // The columns of the matrix.

}

// Simd_quatd - A double-precision quaternion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_quatd
type Simd_quatd struct {
	Angle  float64      // The angle, in radians, by which the quaternion’s action rotates.
	Axis   float64      // The normalized axis about which the quaternion’s action rotates.
	Imag   float64      // The imaginary part of the quaternion.
	Real   float64      // The real part of the quaternion.
	Length float64      // The length of the quaternion.
	Vector Simd_double4 // The underlying vector of the quaternion.

}

// Simd_quatf - A single-precision quaternion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/simd/simd_quatf
type Simd_quatf struct {
	Angle  float32     // The angle, in radians, by which the quaternion’s action rotates.
	Axis   float32     // The normalized axis about which the quaternion’s action rotates.
	Imag   float32     // The imaginary part of the quaternion.
	Real   float32     // The real part of the quaternion.
	Length float32     // The length of the quaternion.
	Vector Simd_float4 // The underlying vector of the quaternion.

}

// VDSP_int24 - A data structure that holds a 24-bit signed integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_int24
type VDSP_int24 struct {
	Bytes uint8 // The bytes that represent the value.

}

// VDSP_uint24 - A data structure that holds a 24-bit unsigned integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_uint24
type VDSP_uint24 struct {
	Bytes uint8 // The bytes that represent the value.

}

// VImageChannelDescription - A description of the range and clamp limits for a pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageChannelDescription
type VImageChannelDescription struct {
	Min  float64 // The minimum encoded value.
	Zero float64 // The encoding for the value zero.
	Full float64 // The encoding for the value one.
	Max  float64 // The maximum encoded value.

}

// VImageRGBPrimaries - A representation of the chromaticity of primaries that define a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageRGBPrimaries
type VImageRGBPrimaries struct {
	Red_x   float32 // The red `x` value according to the CIE 1931 color space.
	Green_x float32 // The green `x` value according to the CIE 1931 color space.
	Blue_x  float32 // The blue `x` value according to the CIE 1931 color space.
	White_x float32 // The white point `x` value according to the CIE 1931 color space.
	Red_y   float32 // The red `y` value according to the CIE 1931 color space.
	Green_y float32 // The green_ _`y` value according to the CIE 1931 color space.
	Blue_y  float32 // The blue `y` value according to the CIE 1931 color space.
	White_y float32 // The white point `y` value according to the CIE 1931 color space.

}

// VImageTransferFunction - A transfer function to convert from linear to nonlinear RGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageTransferFunction
type VImageTransferFunction struct {
	Cutoff float64 // The `cutoff` in the transfer function.
	C0     float64 // The `c0` in the transfer function.
	C1     float64 // The `c1` in the transfer function.
	C2     float64 // The `c2` in the transfer function.
	C3     float64 // The `c3` in the transfer function.
	Gamma  float64 // The `gamma` in the transfer function.
	C4     float64 // The `c4` in the transfer function.
	C5     float64 // The `c5` in the transfer function.

}

// VImageWhitePoint - A representation of a white point according to the CIE 1931 color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageWhitePoint
type VImageWhitePoint struct {
	White_x float32 // The white point `x` value according to the CIE 1931 color space.
	White_y float32 // The white point `y` value according to the CIE 1931 color space.

}

// VImage_ARGBToYpCbCr - The information that describes the conversion from ARGB to YpCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_ARGBToYpCbCr
type VImage_ARGBToYpCbCr struct {
	Opaque uint8
}

// VImage_ARGBToYpCbCrMatrix - The 3 x 3 matrix that the vImage library uses to convert from RGB to YpCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_ARGBToYpCbCrMatrix
type VImage_ARGBToYpCbCrMatrix struct {
	R_Yp      float32 // The  value in the conversion matrix.
	G_Yp      float32 // The  value in the conversion matrix.
	B_Yp      float32 // The  value in the conversion matrix.
	R_Cb      float32 // The  value in the conversion matrix.
	G_Cb      float32 // The  value in the conversion matrix.
	B_Cb_R_Cr float32 // The  value in the conversion matrix.
	G_Cr      float32 // The  value in the conversion matrix.
	B_Cr      float32 // The  value in the conversion matrix.

}

// VImage_AffineTransform - A structure for values that represent an affine transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_AffineTransform
type VImage_AffineTransform struct {
	A  float32 // The entry at position `[1,1]` in the matrix.
	B  float32 // The entry at position `[1,2]` in the matrix.
	C  float32 // The entry at position `[2,1]` in the matrix.
	D  float32 // The entry at position `[2,2]` in the matrix.
	Tx float32 // The entry at position `[3,1]` in the matrix.
	Ty float32 // The entry at position `[3,2]` in the matrix.

}

// VImage_AffineTransform_Double - A structure for values that represent a double-precision affine transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_AffineTransform_Double
type VImage_AffineTransform_Double struct {
	A  float64 // The entry at position `[1,1]` in the matrix.
	B  float64 // The entry at position `[1,2]` in the matrix.
	C  float64 // The entry at position `[2,1]` in the matrix.
	D  float64 // The entry at position `[2,2]` in the matrix.
	Tx float64 // The entry at position `[3,1]` in the matrix.
	Ty float64 // The entry at position `[3,2]` in the matrix.

}

// VImage_Buffer - An image buffer that stores an image’s pixel data, dimensions, and row stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_Buffer
type VImage_Buffer struct {
	Height   uint           // The height of the image, in pixels.
	Width    uint           // The width of the image, in pixels.
	RowBytes uintptr        // The distance, in bytes, between the start of one pixel row and the next in an image, including any unused space between them.
	Data     unsafe.Pointer // A pointer to the top-left pixel of the image.

}

// VImage_CGImageFormat - The description of a Core Graphics image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_CGImageFormat
type VImage_CGImageFormat struct {
	BitsPerComponent uint32                              // The number of bits that represents one channel of data in one pixel.
	BitsPerPixel     uint32                              // The number of bits that represents one pixel.
	ColorSpace       coregraphics.CGColorSpaceRef        // A description of the position of the pixel data in the image, relative to a reference XYZ color space.
	BitmapInfo       coregraphics.CGBitmapInfo           // The component information that describes the color channels.
	Version          uint32                              // The version number.
	Decode           *float64                            // The decode array for the image.
	RenderingIntent  coregraphics.CGColorRenderingIntent // A rendering intent constant that specifies how Core Graphics handles colors that aren’t within the destination color space gamut.

}

// VImage_PerpsectiveTransform - A projective-transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_PerpsectiveTransform
type VImage_PerpsectiveTransform struct {
	A  float32 // The top-left cell in the 3 x 3 transformation matrix.
	B  float32 // The top-middle cell in the 3 x 3 transformation matrix.
	C  float32 // The middle-left cell in the 3 x 3 transformation matrix.
	D  float32 // The middle-middle cell in the 3 x 3 transformation matrix.
	Tx float32 // The x-coordinate translation.
	Ty float32 // The y-coordinate translation.
	Vx float32 // The x-component of the projective vector.
	Vy float32 // The y-component of the projective vector.
	V  float32 // The homogeneous scale factor.

}

// VImage_YpCbCrPixelRange - The description of range and clamping information for YpCbCr pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrPixelRange
type VImage_YpCbCrPixelRange struct {
	Yp_bias      int32 // The encoding for `Y' = 0.0` for this video format (varies by bit depth).
	CbCr_bias    int32 // The encoding for `{Cb, Cr} = 0.0` for this video format.
	YpRangeMax   int32 // The encoding for `Y' = 1.0` for this video format.
	CbCrRangeMax int32 // The encoding for `{Cb, Cr} = 0.5` for this video format.
	YpMax        int32 // The encoding for the maximum allowed Y’ value.
	YpMin        int32 // The encoding of the minimum allowed Y’ value.
	CbCrMax      int32 // The encoding of the maximum allowed `{Cb, Cr}` value.
	CbCrMin      int32 // The encoding of the minimum allowed `{Cb, Cr}` value.

}

// VImage_YpCbCrToARGB - The information that describes the conversion from YpCbCr to ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGB
type VImage_YpCbCrToARGB struct {
	Opaque uint8 // The bytes of the opaque representation.

}

// VImage_YpCbCrToARGBMatrix - The 3 x 3 matrix that the vImage library uses to convert from YpCbCr to RGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGBMatrix
type VImage_YpCbCrToARGBMatrix struct {
	Yp   float32 // The  value in the conversion matrix.
	Cr_R float32 // The  value in the conversion matrix.
	Cr_G float32 // The  value in the conversion matrix.
	Cb_G float32 // The  value in the conversion matrix.
	Cb_B float32 // The  value in the conversion matrix.

}
