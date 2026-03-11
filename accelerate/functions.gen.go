// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Accelerate: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Accelerate: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Accelerate: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _bLASGetThreading func() BLAS_THREADING

// BLASGetThreading returns the current BLAS and LAPACK threading model.
//
// See: https://developer.apple.com/documentation/Accelerate/BLASGetThreading()
func BLASGetThreading() BLAS_THREADING {
	if _bLASGetThreading == nil {
		panic("Accelerate: symbol BLASGetThreading not loaded")
	}
	return _bLASGetThreading()
}


var _bLASSetThreading func(threading BLAS_THREADING) int

// BLASSetThreading sets the BLAS and LAPACK threading model.
//
// See: https://developer.apple.com/documentation/Accelerate/BLASSetThreading(_:)
func BLASSetThreading(threading BLAS_THREADING) int {
	if _bLASSetThreading == nil {
		panic("Accelerate: symbol BLASSetThreading not loaded")
	}
	return _bLASSetThreading(threading)
}


var _bNNSApplyMultiheadAttention func(F BNNSFilter, batch_size uintptr, query unsafe.Pointer, query_stride uintptr, key unsafe.Pointer, key_stride uintptr, key_mask *BNNSNDArrayDescriptor, key_mask_stride uintptr, value unsafe.Pointer, value_stride uintptr, output unsafe.Pointer, output_stride uintptr, add_to_attention *BNNSNDArrayDescriptor, backprop_cache_size *uintptr, backprop_cache unsafe.Pointer, workspace_size *uintptr, workspace unsafe.Pointer) int

// BNNSApplyMultiheadAttention applies a mutihead attention filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSApplyMultiheadAttention(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSApplyMultiheadAttention(F BNNSFilter, batch_size uintptr, query unsafe.Pointer, query_stride uintptr, key unsafe.Pointer, key_stride uintptr, key_mask *BNNSNDArrayDescriptor, key_mask_stride uintptr, value unsafe.Pointer, value_stride uintptr, output unsafe.Pointer, output_stride uintptr, add_to_attention *BNNSNDArrayDescriptor, backprop_cache_size *uintptr, backprop_cache unsafe.Pointer, workspace_size *uintptr, workspace unsafe.Pointer) int {
	if _bNNSApplyMultiheadAttention == nil {
		panic("Accelerate: symbol BNNSApplyMultiheadAttention not loaded")
	}
	return _bNNSApplyMultiheadAttention(F, batch_size, query, query_stride, key, key_stride, key_mask, key_mask_stride, value, value_stride, output, output_stride, add_to_attention, backprop_cache_size, backprop_cache, workspace_size, workspace)
}


var _bNNSApplyMultiheadAttentionBackward func(F BNNSFilter, batch_size uintptr, query unsafe.Pointer, query_stride uintptr, query_param_delta *BNNSMHAProjectionParameters, key unsafe.Pointer, key_stride uintptr, key_mask *BNNSNDArrayDescriptor, key_mask_stride uintptr, key_param_delta *BNNSMHAProjectionParameters, value unsafe.Pointer, value_stride uintptr, value_param_delta *BNNSMHAProjectionParameters, add_to_attention *BNNSNDArrayDescriptor, key_attn_bias_delta *BNNSNDArrayDescriptor, value_attn_bias_delta *BNNSNDArrayDescriptor, output unsafe.Pointer, output_stride uintptr, output_param_delta *BNNSMHAProjectionParameters, backprop_cache_size uintptr, backprop_cache unsafe.Pointer, workspace_size *uintptr, workspace unsafe.Pointer) int

// BNNSApplyMultiheadAttentionBackward applies a multihead attention filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSApplyMultiheadAttentionBackward(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSApplyMultiheadAttentionBackward(F BNNSFilter, batch_size uintptr, query unsafe.Pointer, query_stride uintptr, query_param_delta *BNNSMHAProjectionParameters, key unsafe.Pointer, key_stride uintptr, key_mask *BNNSNDArrayDescriptor, key_mask_stride uintptr, key_param_delta *BNNSMHAProjectionParameters, value unsafe.Pointer, value_stride uintptr, value_param_delta *BNNSMHAProjectionParameters, add_to_attention *BNNSNDArrayDescriptor, key_attn_bias_delta *BNNSNDArrayDescriptor, value_attn_bias_delta *BNNSNDArrayDescriptor, output unsafe.Pointer, output_stride uintptr, output_param_delta *BNNSMHAProjectionParameters, backprop_cache_size uintptr, backprop_cache unsafe.Pointer, workspace_size *uintptr, workspace unsafe.Pointer) int {
	if _bNNSApplyMultiheadAttentionBackward == nil {
		panic("Accelerate: symbol BNNSApplyMultiheadAttentionBackward not loaded")
	}
	return _bNNSApplyMultiheadAttentionBackward(F, batch_size, query, query_stride, query_param_delta, key, key_stride, key_mask, key_mask_stride, key_param_delta, value, value_stride, value_param_delta, add_to_attention, key_attn_bias_delta, value_attn_bias_delta, output, output_stride, output_param_delta, backprop_cache_size, backprop_cache, workspace_size, workspace)
}


var _bNNSArithmeticFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride *uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int

// BNNSArithmeticFilterApplyBackwardBatch applies an arithmetic filter backward to generate input gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:)
func BNNSArithmeticFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride *uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int {
	if _bNNSArithmeticFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSArithmeticFilterApplyBackwardBatch not loaded")
	}
	return _bNNSArithmeticFilterApplyBackwardBatch(filter, batch_size, number_of_inputs, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride)
}


var _bNNSArithmeticFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, out unsafe.Pointer, out_stride uintptr) int

// BNNSArithmeticFilterApplyBatch applies an arithmetic filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticFilterApplyBatch(_:_:_:_:_:_:_:)
func BNNSArithmeticFilterApplyBatch(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, out unsafe.Pointer, out_stride uintptr) int {
	if _bNNSArithmeticFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSArithmeticFilterApplyBatch not loaded")
	}
	return _bNNSArithmeticFilterApplyBatch(filter, batch_size, number_of_inputs, in, in_stride, out, out_stride)
}


var _bNNSBandPart func(num_lower int, num_upper int, input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSBandPart copies the specified subdiagonals and superdiagonals of a matrix, and sets other elements to zero.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSBandPart(_:_:_:_:_:)
func BNNSBandPart(num_lower int, num_upper int, input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSBandPart == nil {
		panic("Accelerate: symbol BNNSBandPart not loaded")
	}
	return _bNNSBandPart(num_lower, num_upper, input, output, filter_params)
}


var _bNNSClipByGlobalNorm func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, count uintptr, max_norm float32, use_norm float32) int

// BNNSClipByGlobalNorm clips a tensor’s values to a maximum global Euclidean norm.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSClipByGlobalNorm(_:_:_:_:_:)
func BNNSClipByGlobalNorm(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, count uintptr, max_norm float32, use_norm float32) int {
	if _bNNSClipByGlobalNorm == nil {
		panic("Accelerate: symbol BNNSClipByGlobalNorm not loaded")
	}
	return _bNNSClipByGlobalNorm(dest, src, count, max_norm, use_norm)
}


var _bNNSClipByNorm func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, max_norm float32, axis_flags uint32) int

// BNNSClipByNorm clips a tensor’s values to a maximum Euclidean norm.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSClipByNorm(_:_:_:_:)
func BNNSClipByNorm(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, max_norm float32, axis_flags uint32) int {
	if _bNNSClipByNorm == nil {
		panic("Accelerate: symbol BNNSClipByNorm not loaded")
	}
	return _bNNSClipByNorm(dest, src, max_norm, axis_flags)
}


var _bNNSClipByValue func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, min_val float32, max_val float32) int

// BNNSClipByValue clips a tensor’s values to the specified minimum and maximum values.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSClipByValue(_:_:_:_:)
func BNNSClipByValue(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, min_val float32, max_val float32) int {
	if _bNNSClipByValue == nil {
		panic("Accelerate: symbol BNNSClipByValue not loaded")
	}
	return _bNNSClipByValue(dest, src, min_val, max_val)
}


var _bNNSCompareTensor func(in0 *BNNSNDArrayDescriptor, in1 *BNNSNDArrayDescriptor, op unsafe.Pointer, out *BNNSNDArrayDescriptor) int

// BNNSCompareTensor returns a tensor of Boolean type by comparing or performing a logical operation between two inputs.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCompareTensor(_:_:_:_:)
func BNNSCompareTensor(in0 *BNNSNDArrayDescriptor, in1 *BNNSNDArrayDescriptor, op unsafe.Pointer, out *BNNSNDArrayDescriptor) int {
	if _bNNSCompareTensor == nil {
		panic("Accelerate: symbol BNNSCompareTensor not loaded")
	}
	return _bNNSCompareTensor(in0, in1, op, out)
}


var _bNNSComputeLSTMTrainingCacheCapacity func(layer_params *BNNSLayerParametersLSTM) uintptr

// BNNSComputeLSTMTrainingCacheCapacity returns the minimum bytes capacity of the training cache buffer a long short-term memory (LSTM) layer uses when it’s applied.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSComputeLSTMTrainingCacheCapacity(_:)
func BNNSComputeLSTMTrainingCacheCapacity(layer_params *BNNSLayerParametersLSTM) uintptr {
	if _bNNSComputeLSTMTrainingCacheCapacity == nil {
		panic("Accelerate: symbol BNNSComputeLSTMTrainingCacheCapacity not loaded")
	}
	return _bNNSComputeLSTMTrainingCacheCapacity(layer_params)
}


var _bNNSComputeNorm func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, norm_type unsafe.Pointer, axis_flags uint32) int

// BNNSComputeNorm computes the specified norm over an entire tensor or the specified axes.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSComputeNorm(_:_:_:_:)
func BNNSComputeNorm(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, norm_type unsafe.Pointer, axis_flags uint32) int {
	if _bNNSComputeNorm == nil {
		panic("Accelerate: symbol BNNSComputeNorm not loaded")
	}
	return _bNNSComputeNorm(dest, src, norm_type, axis_flags)
}


var _bNNSComputeNormBackward func(in unsafe.Pointer, in_delta *BNNSNDArrayDescriptor, out unsafe.Pointer, out_delta *BNNSNDArrayDescriptor, norm_type unsafe.Pointer, axis_flags uint32) int

// BNNSComputeNormBackward backpropogates gradients for the compute norm function.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSComputeNormBackward(_:_:_:_:_:_:)
func BNNSComputeNormBackward(in unsafe.Pointer, in_delta *BNNSNDArrayDescriptor, out unsafe.Pointer, out_delta *BNNSNDArrayDescriptor, norm_type unsafe.Pointer, axis_flags uint32) int {
	if _bNNSComputeNormBackward == nil {
		panic("Accelerate: symbol BNNSComputeNormBackward not loaded")
	}
	return _bNNSComputeNormBackward(in, in_delta, out, out_delta, norm_type, axis_flags)
}


var _bNNSCopy func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSCopy copies the contents of an n-dimensional array descriptor to another of the same shape.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCopy(_:_:_:)
func BNNSCopy(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSCopy == nil {
		panic("Accelerate: symbol BNNSCopy not loaded")
	}
	return _bNNSCopy(dest, src, filter_params)
}


var _bNNSCreateNearestNeighbors func(max_n_samples uint, n_features uint, n_neighbors uint, data_type unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSNearestNeighbors

// BNNSCreateNearestNeighbors returns a new k-nearest neighbors object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCreateNearestNeighbors(_:_:_:_:_:)
func BNNSCreateNearestNeighbors(max_n_samples uint, n_features uint, n_neighbors uint, data_type unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSNearestNeighbors {
	if _bNNSCreateNearestNeighbors == nil {
		panic("Accelerate: symbol BNNSCreateNearestNeighbors not loaded")
	}
	return _bNNSCreateNearestNeighbors(max_n_samples, n_features, n_neighbors, data_type, filter_params)
}


var _bNNSCreateRandomGenerator func(method unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSRandomGenerator

// BNNSCreateRandomGenerator returns a new random number generator using an internally generated random seed.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCreateRandomGenerator(_:_:)
func BNNSCreateRandomGenerator(method unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSRandomGenerator {
	if _bNNSCreateRandomGenerator == nil {
		panic("Accelerate: symbol BNNSCreateRandomGenerator not loaded")
	}
	return _bNNSCreateRandomGenerator(method, filter_params)
}


var _bNNSCreateRandomGeneratorWithSeed func(method unsafe.Pointer, seed uint64, filter_params *BNNSFilterParameters) BNNSRandomGenerator

// BNNSCreateRandomGeneratorWithSeed returns a new random number generator using the specified seed.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCreateRandomGeneratorWithSeed(_:_:_:)
func BNNSCreateRandomGeneratorWithSeed(method unsafe.Pointer, seed uint64, filter_params *BNNSFilterParameters) BNNSRandomGenerator {
	if _bNNSCreateRandomGeneratorWithSeed == nil {
		panic("Accelerate: symbol BNNSCreateRandomGeneratorWithSeed not loaded")
	}
	return _bNNSCreateRandomGeneratorWithSeed(method, seed, filter_params)
}


var _bNNSCropResize func(layer_params *BNNSLayerParametersCropResize, input *BNNSNDArrayDescriptor, roi *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSCropResize extracts and resizes regions of interest of an input tensor.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCropResize(_:_:_:_:_:)
func BNNSCropResize(layer_params *BNNSLayerParametersCropResize, input *BNNSNDArrayDescriptor, roi *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSCropResize == nil {
		panic("Accelerate: symbol BNNSCropResize not loaded")
	}
	return _bNNSCropResize(layer_params, input, roi, output, filter_params)
}


var _bNNSCropResizeBackward func(layer_params *BNNSLayerParametersCropResize, in_delta *BNNSNDArrayDescriptor, roi *BNNSNDArrayDescriptor, out_delta *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSCropResizeBackward applies a crop-resize filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSCropResizeBackward(_:_:_:_:_:)
func BNNSCropResizeBackward(layer_params *BNNSLayerParametersCropResize, in_delta *BNNSNDArrayDescriptor, roi *BNNSNDArrayDescriptor, out_delta *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSCropResizeBackward == nil {
		panic("Accelerate: symbol BNNSCropResizeBackward not loaded")
	}
	return _bNNSCropResizeBackward(layer_params, in_delta, roi, out_delta, filter_params)
}


var _bNNSDataLayoutGetRank func(layout unsafe.Pointer) uintptr

// BNNSDataLayoutGetRank.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDataLayoutGetRank(_:)
func BNNSDataLayoutGetRank(layout unsafe.Pointer) uintptr {
	if _bNNSDataLayoutGetRank == nil {
		panic("Accelerate: symbol BNNSDataLayoutGetRank not loaded")
	}
	return _bNNSDataLayoutGetRank(layout)
}


var _bNNSDestroyNearestNeighbors func(knn BNNSNearestNeighbors)

// BNNSDestroyNearestNeighbors destroys a k-nearest neighbors object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDestroyNearestNeighbors(_:)
func BNNSDestroyNearestNeighbors(knn BNNSNearestNeighbors) {
	if _bNNSDestroyNearestNeighbors == nil {
		panic("Accelerate: symbol BNNSDestroyNearestNeighbors not loaded")
	}
	_bNNSDestroyNearestNeighbors(knn)
}


var _bNNSDestroyRandomGenerator func(generator BNNSRandomGenerator)

// BNNSDestroyRandomGenerator destroys a random number generator.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDestroyRandomGenerator(_:)
func BNNSDestroyRandomGenerator(generator BNNSRandomGenerator) {
	if _bNNSDestroyRandomGenerator == nil {
		panic("Accelerate: symbol BNNSDestroyRandomGenerator not loaded")
	}
	_bNNSDestroyRandomGenerator(generator)
}


var _bNNSDirectApplyActivationBatch func(layer_params *BNNSLayerParametersActivation, filter_params *BNNSFilterParameters, batch_size uintptr, in_stride uintptr, out_stride uintptr) int

// BNNSDirectApplyActivationBatch applies an activation filter to a set of input objects, writing out the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyActivationBatch(_:_:_:_:_:)
func BNNSDirectApplyActivationBatch(layer_params *BNNSLayerParametersActivation, filter_params *BNNSFilterParameters, batch_size uintptr, in_stride uintptr, out_stride uintptr) int {
	if _bNNSDirectApplyActivationBatch == nil {
		panic("Accelerate: symbol BNNSDirectApplyActivationBatch not loaded")
	}
	return _bNNSDirectApplyActivationBatch(layer_params, filter_params, batch_size, in_stride, out_stride)
}


var _bNNSDirectApplyInTopK func(K uintptr, axis uintptr, batch_size uintptr, input *BNNSNDArrayDescriptor, input_batch_stride uintptr, test_indices *BNNSNDArrayDescriptor, test_indices_batch_stride uintptr, output *BNNSNDArrayDescriptor, output_batch_stride uintptr, filter_params *BNNSFilterParameters) int

// BNNSDirectApplyInTopK applies an in-top-k filter directly to an input.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyInTopK(_:_:_:_:_:_:_:_:_:_:)
func BNNSDirectApplyInTopK(K uintptr, axis uintptr, batch_size uintptr, input *BNNSNDArrayDescriptor, input_batch_stride uintptr, test_indices *BNNSNDArrayDescriptor, test_indices_batch_stride uintptr, output *BNNSNDArrayDescriptor, output_batch_stride uintptr, filter_params *BNNSFilterParameters) int {
	if _bNNSDirectApplyInTopK == nil {
		panic("Accelerate: symbol BNNSDirectApplyInTopK not loaded")
	}
	return _bNNSDirectApplyInTopK(K, axis, batch_size, input, input_batch_stride, test_indices, test_indices_batch_stride, output, output_batch_stride, filter_params)
}


var _bNNSDirectApplyLSTMBatchBackward func(layer_params *BNNSLayerParametersLSTM, layer_delta_params *BNNSLayerParametersLSTM, filter_params *BNNSFilterParameters, training_cache_ptr unsafe.Pointer, training_cache_capacity uintptr) int

// BNNSDirectApplyLSTMBatchBackward applies a long short-term memory (LSTM) filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyLSTMBatchBackward(_:_:_:_:_:)
func BNNSDirectApplyLSTMBatchBackward(layer_params *BNNSLayerParametersLSTM, layer_delta_params *BNNSLayerParametersLSTM, filter_params *BNNSFilterParameters, training_cache_ptr unsafe.Pointer, training_cache_capacity uintptr) int {
	if _bNNSDirectApplyLSTMBatchBackward == nil {
		panic("Accelerate: symbol BNNSDirectApplyLSTMBatchBackward not loaded")
	}
	return _bNNSDirectApplyLSTMBatchBackward(layer_params, layer_delta_params, filter_params, training_cache_ptr, training_cache_capacity)
}


var _bNNSDirectApplyLSTMBatchTrainingCaching func(layer_params *BNNSLayerParametersLSTM, filter_params *BNNSFilterParameters, training_cache_ptr unsafe.Pointer, training_cache_capacity uintptr) int

// BNNSDirectApplyLSTMBatchTrainingCaching applies a long short-term memory (LSTM) layer directly to an input.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyLSTMBatchTrainingCaching(_:_:_:_:)
func BNNSDirectApplyLSTMBatchTrainingCaching(layer_params *BNNSLayerParametersLSTM, filter_params *BNNSFilterParameters, training_cache_ptr unsafe.Pointer, training_cache_capacity uintptr) int {
	if _bNNSDirectApplyLSTMBatchTrainingCaching == nil {
		panic("Accelerate: symbol BNNSDirectApplyLSTMBatchTrainingCaching not loaded")
	}
	return _bNNSDirectApplyLSTMBatchTrainingCaching(layer_params, filter_params, training_cache_ptr, training_cache_capacity)
}


var _bNNSDirectApplyQuantizer func(layer_params *BNNSLayerParametersQuantization, filter_params *BNNSFilterParameters, batch_size uintptr, input_stride uintptr, output_stride uintptr) int

// BNNSDirectApplyQuantizer applies a quantization layer directly to two input matrices.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyQuantizer(_:_:_:_:_:)
func BNNSDirectApplyQuantizer(layer_params *BNNSLayerParametersQuantization, filter_params *BNNSFilterParameters, batch_size uintptr, input_stride uintptr, output_stride uintptr) int {
	if _bNNSDirectApplyQuantizer == nil {
		panic("Accelerate: symbol BNNSDirectApplyQuantizer not loaded")
	}
	return _bNNSDirectApplyQuantizer(layer_params, filter_params, batch_size, input_stride, output_stride)
}


var _bNNSDirectApplyReduction func(layer_params *BNNSLayerParametersReduction, filter_params *BNNSFilterParameters) int

// BNNSDirectApplyReduction applies a reduction operation directly to an input tensor.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyReduction(_:_:)
func BNNSDirectApplyReduction(layer_params *BNNSLayerParametersReduction, filter_params *BNNSFilterParameters) int {
	if _bNNSDirectApplyReduction == nil {
		panic("Accelerate: symbol BNNSDirectApplyReduction not loaded")
	}
	return _bNNSDirectApplyReduction(layer_params, filter_params)
}


var _bNNSDirectApplyTopK func(K uintptr, axis uintptr, batch_size uintptr, input *BNNSNDArrayDescriptor, input_batch_stride uintptr, best_values *BNNSNDArrayDescriptor, best_values_batch_stride uintptr, best_indices *BNNSNDArrayDescriptor, best_indices_batch_stride uintptr, filter_params *BNNSFilterParameters) int

// BNNSDirectApplyTopK applies a top-k filter directly to an input.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyTopK(_:_:_:_:_:_:_:_:_:_:)
func BNNSDirectApplyTopK(K uintptr, axis uintptr, batch_size uintptr, input *BNNSNDArrayDescriptor, input_batch_stride uintptr, best_values *BNNSNDArrayDescriptor, best_values_batch_stride uintptr, best_indices *BNNSNDArrayDescriptor, best_indices_batch_stride uintptr, filter_params *BNNSFilterParameters) int {
	if _bNNSDirectApplyTopK == nil {
		panic("Accelerate: symbol BNNSDirectApplyTopK not loaded")
	}
	return _bNNSDirectApplyTopK(K, axis, batch_size, input, input_batch_stride, best_values, best_values_batch_stride, best_indices, best_indices_batch_stride, filter_params)
}


var _bNNSFilterApply func(filter BNNSFilter, in unsafe.Pointer, out unsafe.Pointer) int

// BNNSFilterApply applies a filter to an input, writing the result to a specified output.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApply(_:_:_:)
func BNNSFilterApply(filter BNNSFilter, in unsafe.Pointer, out unsafe.Pointer) int {
	if _bNNSFilterApply == nil {
		panic("Accelerate: symbol BNNSFilterApply not loaded")
	}
	return _bNNSFilterApply(filter, in, out)
}


var _bNNSFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, weights_delta *BNNSNDArrayDescriptor, bias_delta *BNNSNDArrayDescriptor) int

// BNNSFilterApplyBackwardBatch applies a filter backward to generate input delta, weights delta and bias delta.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, weights_delta *BNNSNDArrayDescriptor, bias_delta *BNNSNDArrayDescriptor) int {
	if _bNNSFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSFilterApplyBackwardBatch not loaded")
	}
	return _bNNSFilterApplyBackwardBatch(filter, batch_size, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, weights_delta, bias_delta)
}


var _bNNSFilterApplyBackwardTwoInputBatch func(filter BNNSFilter, batch_size uintptr, inA unsafe.Pointer, inA_stride uintptr, inA_delta *BNNSNDArrayDescriptor, inA_delta_stride uintptr, inB unsafe.Pointer, inB_stride uintptr, inB_delta *BNNSNDArrayDescriptor, inB_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, weights_delta *BNNSNDArrayDescriptor, bias_delta *BNNSNDArrayDescriptor) int

// BNNSFilterApplyBackwardTwoInputBatch applies a filter backward to generate input deltas, weights delta and bias delta.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyBackwardTwoInputBatch(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSFilterApplyBackwardTwoInputBatch(filter BNNSFilter, batch_size uintptr, inA unsafe.Pointer, inA_stride uintptr, inA_delta *BNNSNDArrayDescriptor, inA_delta_stride uintptr, inB unsafe.Pointer, inB_stride uintptr, inB_delta *BNNSNDArrayDescriptor, inB_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, weights_delta *BNNSNDArrayDescriptor, bias_delta *BNNSNDArrayDescriptor) int {
	if _bNNSFilterApplyBackwardTwoInputBatch == nil {
		panic("Accelerate: symbol BNNSFilterApplyBackwardTwoInputBatch not loaded")
	}
	return _bNNSFilterApplyBackwardTwoInputBatch(filter, batch_size, inA, inA_stride, inA_delta, inA_delta_stride, inB, inB_stride, inB_delta, inB_delta_stride, out, out_stride, out_delta, out_delta_stride, weights_delta, bias_delta)
}


var _bNNSFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr) int

// BNNSFilterApplyBatch applies a filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyBatch(_:_:_:_:_:_:)
func BNNSFilterApplyBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr) int {
	if _bNNSFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSFilterApplyBatch not loaded")
	}
	return _bNNSFilterApplyBatch(filter, batch_size, in, in_stride, out, out_stride)
}


var _bNNSFilterApplyTwoInput func(filter BNNSFilter, inA unsafe.Pointer, inB unsafe.Pointer, out unsafe.Pointer) int

// BNNSFilterApplyTwoInput applies a filter to a pair of inputs, writing the result to a specified output.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyTwoInput(_:_:_:_:)
func BNNSFilterApplyTwoInput(filter BNNSFilter, inA unsafe.Pointer, inB unsafe.Pointer, out unsafe.Pointer) int {
	if _bNNSFilterApplyTwoInput == nil {
		panic("Accelerate: symbol BNNSFilterApplyTwoInput not loaded")
	}
	return _bNNSFilterApplyTwoInput(filter, inA, inB, out)
}


var _bNNSFilterApplyTwoInputBatch func(filter BNNSFilter, batch_size uintptr, inA unsafe.Pointer, inA_stride uintptr, inB unsafe.Pointer, inB_stride uintptr, out unsafe.Pointer, out_stride uintptr) int

// BNNSFilterApplyTwoInputBatch applies a filter to a set of input object pairs, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyTwoInputBatch(_:_:_:_:_:_:_:_:)
func BNNSFilterApplyTwoInputBatch(filter BNNSFilter, batch_size uintptr, inA unsafe.Pointer, inA_stride uintptr, inB unsafe.Pointer, inB_stride uintptr, out unsafe.Pointer, out_stride uintptr) int {
	if _bNNSFilterApplyTwoInputBatch == nil {
		panic("Accelerate: symbol BNNSFilterApplyTwoInputBatch not loaded")
	}
	return _bNNSFilterApplyTwoInputBatch(filter, batch_size, inA, inA_stride, inB, inB_stride, out, out_stride)
}


var _bNNSFilterCreateFusedLayer func(number_of_fused_filters uintptr, filter_type *uintptr, layer_params unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateFusedLayer returns a new fused layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateFusedLayer(_:_:_:_:)
func BNNSFilterCreateFusedLayer(number_of_fused_filters uintptr, filter_type *uintptr, layer_params unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateFusedLayer == nil {
		panic("Accelerate: symbol BNNSFilterCreateFusedLayer not loaded")
	}
	return _bNNSFilterCreateFusedLayer(number_of_fused_filters, filter_type, layer_params, filter_params)
}


var _bNNSFilterCreateLayerActivation func(layer_params *BNNSLayerParametersActivation, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerActivation returns a new activation layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerActivation(_:_:)
func BNNSFilterCreateLayerActivation(layer_params *BNNSLayerParametersActivation, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerActivation == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerActivation not loaded")
	}
	return _bNNSFilterCreateLayerActivation(layer_params, filter_params)
}


var _bNNSFilterCreateLayerArithmetic func(layer_params *BNNSLayerParametersArithmetic, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerArithmetic returns a new arithmetic layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerArithmetic(_:_:)
func BNNSFilterCreateLayerArithmetic(layer_params *BNNSLayerParametersArithmetic, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerArithmetic == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerArithmetic not loaded")
	}
	return _bNNSFilterCreateLayerArithmetic(layer_params, filter_params)
}


var _bNNSFilterCreateLayerBroadcastMatMul func(layer_params *BNNSLayerParametersBroadcastMatMul, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerBroadcastMatMul returns a new broadcast matrix multiply layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerBroadcastMatMul(_:_:)
func BNNSFilterCreateLayerBroadcastMatMul(layer_params *BNNSLayerParametersBroadcastMatMul, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerBroadcastMatMul == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerBroadcastMatMul not loaded")
	}
	return _bNNSFilterCreateLayerBroadcastMatMul(layer_params, filter_params)
}


var _bNNSFilterCreateLayerConvolution func(layer_params *BNNSLayerParametersConvolution, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerConvolution returns a new convolution layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerConvolution(_:_:)
func BNNSFilterCreateLayerConvolution(layer_params *BNNSLayerParametersConvolution, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerConvolution == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerConvolution not loaded")
	}
	return _bNNSFilterCreateLayerConvolution(layer_params, filter_params)
}


var _bNNSFilterCreateLayerDropout func(layer_params *BNNSLayerParametersDropout, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerDropout returns a new dropout layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerDropout(_:_:)
func BNNSFilterCreateLayerDropout(layer_params *BNNSLayerParametersDropout, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerDropout == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerDropout not loaded")
	}
	return _bNNSFilterCreateLayerDropout(layer_params, filter_params)
}


var _bNNSFilterCreateLayerEmbedding func(layer_params *BNNSLayerParametersEmbedding, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerEmbedding returns a new embedding layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerEmbedding(_:_:)
func BNNSFilterCreateLayerEmbedding(layer_params *BNNSLayerParametersEmbedding, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerEmbedding == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerEmbedding not loaded")
	}
	return _bNNSFilterCreateLayerEmbedding(layer_params, filter_params)
}


var _bNNSFilterCreateLayerFullyConnected func(layer_params *BNNSLayerParametersFullyConnected, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerFullyConnected returns a new fully connected layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerFullyConnected(_:_:)
func BNNSFilterCreateLayerFullyConnected(layer_params *BNNSLayerParametersFullyConnected, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerFullyConnected == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerFullyConnected not loaded")
	}
	return _bNNSFilterCreateLayerFullyConnected(layer_params, filter_params)
}


var _bNNSFilterCreateLayerGram func(layer_params *BNNSLayerParametersGram, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerGram returns a new Gram matrix layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerGram(_:_:)
func BNNSFilterCreateLayerGram(layer_params *BNNSLayerParametersGram, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerGram == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerGram not loaded")
	}
	return _bNNSFilterCreateLayerGram(layer_params, filter_params)
}


var _bNNSFilterCreateLayerLoss func(layer_params unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerLoss returns a new loss layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerLoss(_:_:)
func BNNSFilterCreateLayerLoss(layer_params unsafe.Pointer, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerLoss == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerLoss not loaded")
	}
	return _bNNSFilterCreateLayerLoss(layer_params, filter_params)
}


var _bNNSFilterCreateLayerMultiheadAttention func(layer_params *BNNSLayerParametersMultiheadAttention, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerMultiheadAttention returns a new multihead attention layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerMultiheadAttention(_:_:)
func BNNSFilterCreateLayerMultiheadAttention(layer_params *BNNSLayerParametersMultiheadAttention, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerMultiheadAttention == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerMultiheadAttention not loaded")
	}
	return _bNNSFilterCreateLayerMultiheadAttention(layer_params, filter_params)
}


var _bNNSFilterCreateLayerNormalization func(normType unsafe.Pointer, layer_params *BNNSLayerParametersNormalization, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerNormalization returns a new normalization layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerNormalization(_:_:_:)
func BNNSFilterCreateLayerNormalization(normType unsafe.Pointer, layer_params *BNNSLayerParametersNormalization, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerNormalization == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerNormalization not loaded")
	}
	return _bNNSFilterCreateLayerNormalization(normType, layer_params, filter_params)
}


var _bNNSFilterCreateLayerPadding func(layer_params *BNNSLayerParametersPadding, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerPadding returns a new loss layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerPadding(_:_:)
func BNNSFilterCreateLayerPadding(layer_params *BNNSLayerParametersPadding, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerPadding == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerPadding not loaded")
	}
	return _bNNSFilterCreateLayerPadding(layer_params, filter_params)
}


var _bNNSFilterCreateLayerPermute func(layer_params *BNNSLayerParametersPermute, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerPermute returns a new permute layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerPermute(_:_:)
func BNNSFilterCreateLayerPermute(layer_params *BNNSLayerParametersPermute, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerPermute == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerPermute not loaded")
	}
	return _bNNSFilterCreateLayerPermute(layer_params, filter_params)
}


var _bNNSFilterCreateLayerPooling func(layer_params *BNNSLayerParametersPooling, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerPooling returns a new pooling layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerPooling(_:_:)
func BNNSFilterCreateLayerPooling(layer_params *BNNSLayerParametersPooling, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerPooling == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerPooling not loaded")
	}
	return _bNNSFilterCreateLayerPooling(layer_params, filter_params)
}


var _bNNSFilterCreateLayerReduction func(layer_params *BNNSLayerParametersReduction, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerReduction returns a new reduction layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerReduction(_:_:)
func BNNSFilterCreateLayerReduction(layer_params *BNNSLayerParametersReduction, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerReduction == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerReduction not loaded")
	}
	return _bNNSFilterCreateLayerReduction(layer_params, filter_params)
}


var _bNNSFilterCreateLayerResize func(layer_params *BNNSLayerParametersResize, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerResize returns a new resize layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerResize(_:_:)
func BNNSFilterCreateLayerResize(layer_params *BNNSLayerParametersResize, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerResize == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerResize not loaded")
	}
	return _bNNSFilterCreateLayerResize(layer_params, filter_params)
}


var _bNNSFilterCreateLayerTensorContraction func(layer_params *BNNSLayerParametersTensorContraction, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerTensorContraction returns a new tensor-contraction layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerTensorContraction(_:_:)
func BNNSFilterCreateLayerTensorContraction(layer_params *BNNSLayerParametersTensorContraction, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerTensorContraction == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerTensorContraction not loaded")
	}
	return _bNNSFilterCreateLayerTensorContraction(layer_params, filter_params)
}


var _bNNSFilterCreateLayerTransposedConvolution func(layer_params *BNNSLayerParametersConvolution, filter_params *BNNSFilterParameters) BNNSFilter

// BNNSFilterCreateLayerTransposedConvolution returns a new transposed convolution layer.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerTransposedConvolution(_:_:)
func BNNSFilterCreateLayerTransposedConvolution(layer_params *BNNSLayerParametersConvolution, filter_params *BNNSFilterParameters) BNNSFilter {
	if _bNNSFilterCreateLayerTransposedConvolution == nil {
		panic("Accelerate: symbol BNNSFilterCreateLayerTransposedConvolution not loaded")
	}
	return _bNNSFilterCreateLayerTransposedConvolution(layer_params, filter_params)
}


var _bNNSFilterDestroy func(filter BNNSFilter)

// BNNSFilterDestroy destroys the specified filter, releasing all resources allocated for it.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilterDestroy(_:)
func BNNSFilterDestroy(filter BNNSFilter) {
	if _bNNSFilterDestroy == nil {
		panic("Accelerate: symbol BNNSFilterDestroy not loaded")
	}
	_bNNSFilterDestroy(filter)
}


var _bNNSFusedFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, delta_parameters *BNNSNDArrayDescriptor) int

// BNNSFusedFilterApplyBackwardBatch applies a fused filter backward to generate input gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFusedFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:)
func BNNSFusedFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, delta_parameters *BNNSNDArrayDescriptor) int {
	if _bNNSFusedFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSFusedFilterApplyBackwardBatch not loaded")
	}
	return _bNNSFusedFilterApplyBackwardBatch(filter, batch_size, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, delta_parameters)
}


var _bNNSFusedFilterApplyBackwardMultiInputBatch func(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride *uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, delta_parameters *BNNSNDArrayDescriptor) int

// BNNSFusedFilterApplyBackwardMultiInputBatch applies a multiple-input fused filter backward to generate input gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFusedFilterApplyBackwardMultiInputBatch(_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSFusedFilterApplyBackwardMultiInputBatch(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride *uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, delta_parameters *BNNSNDArrayDescriptor) int {
	if _bNNSFusedFilterApplyBackwardMultiInputBatch == nil {
		panic("Accelerate: symbol BNNSFusedFilterApplyBackwardMultiInputBatch not loaded")
	}
	return _bNNSFusedFilterApplyBackwardMultiInputBatch(filter, batch_size, number_of_inputs, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, delta_parameters)
}


var _bNNSFusedFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int

// BNNSFusedFilterApplyBatch applies a fused filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFusedFilterApplyBatch(_:_:_:_:_:_:_:)
func BNNSFusedFilterApplyBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int {
	if _bNNSFusedFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSFusedFilterApplyBatch not loaded")
	}
	return _bNNSFusedFilterApplyBatch(filter, batch_size, in, in_stride, out, out_stride, training)
}


var _bNNSFusedFilterApplyMultiInputBatch func(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int

// BNNSFusedFilterApplyMultiInputBatch applies a multiple-input fused filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFusedFilterApplyMultiInputBatch(_:_:_:_:_:_:_:_:)
func BNNSFusedFilterApplyMultiInputBatch(filter BNNSFilter, batch_size uintptr, number_of_inputs uintptr, in unsafe.Pointer, in_stride *uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int {
	if _bNNSFusedFilterApplyMultiInputBatch == nil {
		panic("Accelerate: symbol BNNSFusedFilterApplyMultiInputBatch not loaded")
	}
	return _bNNSFusedFilterApplyMultiInputBatch(filter, batch_size, number_of_inputs, in, in_stride, out, out_stride, training)
}


var _bNNSGather func(axis uintptr, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSGather gathers the elements of a tensor along a single axis.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGather(_:_:_:_:_:)
func BNNSGather(axis uintptr, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSGather == nil {
		panic("Accelerate: symbol BNNSGather not loaded")
	}
	return _bNNSGather(axis, input, indices, output, filter_params)
}


var _bNNSGatherND func(input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSGatherND gathers the slices of a tensor.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGatherND(_:_:_:_:)
func BNNSGatherND(input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSGatherND == nil {
		panic("Accelerate: symbol BNNSGatherND not loaded")
	}
	return _bNNSGatherND(input, indices, output, filter_params)
}


var _bNNSGetPointer func(filter BNNSFilter, target unsafe.Pointer) BNNSNDArrayDescriptor

// BNNSGetPointer returns an n-dimensional array descriptor that contains a reference to a filter-data member.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGetPointer(_:_:)
func BNNSGetPointer(filter BNNSFilter, target unsafe.Pointer) BNNSNDArrayDescriptor {
	if _bNNSGetPointer == nil {
		panic("Accelerate: symbol BNNSGetPointer not loaded")
	}
	return _bNNSGetPointer(filter, target)
}


var _bNNSGraphCompileFromFile func(filename *byte, function *byte, options Bnns_graph_compile_options_t) Bnns_graph_t

// BNNSGraphCompileFromFile compiles a source mlmodelc file to a graph object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileFromFile(_:_:_:)
func BNNSGraphCompileFromFile(filename *byte, function *byte, options Bnns_graph_compile_options_t) Bnns_graph_t {
	if _bNNSGraphCompileFromFile == nil {
		panic("Accelerate: symbol BNNSGraphCompileFromFile not loaded")
	}
	return _bNNSGraphCompileFromFile(filename, function, options)
}


var _bNNSGraphCompileOptionsDestroy func(options Bnns_graph_compile_options_t)

// BNNSGraphCompileOptionsDestroy destroys the specified compilation options object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsDestroy(_:)
func BNNSGraphCompileOptionsDestroy(options Bnns_graph_compile_options_t) {
	if _bNNSGraphCompileOptionsDestroy == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsDestroy not loaded")
	}
	_bNNSGraphCompileOptionsDestroy(options)
}


var _bNNSGraphCompileOptionsGetGenerateDebugInfo func(options Bnns_graph_compile_options_t) bool

// BNNSGraphCompileOptionsGetGenerateDebugInfo returns the option for the compiled graph to include debugging information.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetGenerateDebugInfo(_:)
func BNNSGraphCompileOptionsGetGenerateDebugInfo(options Bnns_graph_compile_options_t) bool {
	if _bNNSGraphCompileOptionsGetGenerateDebugInfo == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsGetGenerateDebugInfo not loaded")
	}
	return _bNNSGraphCompileOptionsGetGenerateDebugInfo(options)
}


var _bNNSGraphCompileOptionsGetOptimizationPreference func(options Bnns_graph_compile_options_t) unsafe.Pointer

// BNNSGraphCompileOptionsGetOptimizationPreference returns the option for the compiled graph to optimize for either size or performance.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetOptimizationPreference(_:)
func BNNSGraphCompileOptionsGetOptimizationPreference(options Bnns_graph_compile_options_t) unsafe.Pointer {
	if _bNNSGraphCompileOptionsGetOptimizationPreference == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsGetOptimizationPreference not loaded")
	}
	return _bNNSGraphCompileOptionsGetOptimizationPreference(options)
}


var _bNNSGraphCompileOptionsGetOutputFD func(options Bnns_graph_compile_options_t) int

// BNNSGraphCompileOptionsGetOutputFD returns the option for the compiled graph’s output file descriptor.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetOutputFD(_:)
func BNNSGraphCompileOptionsGetOutputFD(options Bnns_graph_compile_options_t) int {
	if _bNNSGraphCompileOptionsGetOutputFD == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsGetOutputFD not loaded")
	}
	return _bNNSGraphCompileOptionsGetOutputFD(options)
}


var _bNNSGraphCompileOptionsGetOutputPath func(options Bnns_graph_compile_options_t) *byte

// BNNSGraphCompileOptionsGetOutputPath returns the option for the compiled graph’s output path.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetOutputPath(_:)
func BNNSGraphCompileOptionsGetOutputPath(options Bnns_graph_compile_options_t) *byte {
	if _bNNSGraphCompileOptionsGetOutputPath == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsGetOutputPath not loaded")
	}
	return _bNNSGraphCompileOptionsGetOutputPath(options)
}


var _bNNSGraphCompileOptionsGetTargetSingleThread func(options Bnns_graph_compile_options_t) bool

// BNNSGraphCompileOptionsGetTargetSingleThread returns the option for the compiled graph to execute on a single thread.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetTargetSingleThread(_:)
func BNNSGraphCompileOptionsGetTargetSingleThread(options Bnns_graph_compile_options_t) bool {
	if _bNNSGraphCompileOptionsGetTargetSingleThread == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsGetTargetSingleThread not loaded")
	}
	return _bNNSGraphCompileOptionsGetTargetSingleThread(options)
}


var _bNNSGraphCompileOptionsMakeDefault func() Bnns_graph_compile_options_t

// BNNSGraphCompileOptionsMakeDefault returns an allocated compilation options object with default values.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsMakeDefault()
func BNNSGraphCompileOptionsMakeDefault() Bnns_graph_compile_options_t {
	if _bNNSGraphCompileOptionsMakeDefault == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsMakeDefault not loaded")
	}
	return _bNNSGraphCompileOptionsMakeDefault()
}


var _bNNSGraphCompileOptionsSetGenerateDebugInfo func(options Bnns_graph_compile_options_t, value bool)

// BNNSGraphCompileOptionsSetGenerateDebugInfo sets the option for the compiled graph to include debugging information.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetGenerateDebugInfo(_:_:)
func BNNSGraphCompileOptionsSetGenerateDebugInfo(options Bnns_graph_compile_options_t, value bool) {
	if _bNNSGraphCompileOptionsSetGenerateDebugInfo == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetGenerateDebugInfo not loaded")
	}
	_bNNSGraphCompileOptionsSetGenerateDebugInfo(options, value)
}



var _bNNSGraphCompileOptionsSetMessageLogMask func(options Bnns_graph_compile_options_t, log_level_mask uint32)

// BNNSGraphCompileOptionsSetMessageLogMask sets the mask for compile-time messages.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetMessageLogMask(_:_:)
func BNNSGraphCompileOptionsSetMessageLogMask(options Bnns_graph_compile_options_t, log_level_mask uint32) {
	if _bNNSGraphCompileOptionsSetMessageLogMask == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetMessageLogMask not loaded")
	}
	_bNNSGraphCompileOptionsSetMessageLogMask(options, log_level_mask)
}


var _bNNSGraphCompileOptionsSetOptimizationPreference func(options Bnns_graph_compile_options_t, preference unsafe.Pointer)

// BNNSGraphCompileOptionsSetOptimizationPreference sets the option for the compiled graph to optimize for either size or performance.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOptimizationPreference(_:_:)
func BNNSGraphCompileOptionsSetOptimizationPreference(options Bnns_graph_compile_options_t, preference unsafe.Pointer) {
	if _bNNSGraphCompileOptionsSetOptimizationPreference == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetOptimizationPreference not loaded")
	}
	_bNNSGraphCompileOptionsSetOptimizationPreference(options, preference)
}


var _bNNSGraphCompileOptionsSetOutputFD func(options Bnns_graph_compile_options_t, fd int)

// BNNSGraphCompileOptionsSetOutputFD sets the option for graph compilation to generate the graph object directly to the specified file descriptor.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOutputFD(_:_:)
func BNNSGraphCompileOptionsSetOutputFD(options Bnns_graph_compile_options_t, fd int) {
	if _bNNSGraphCompileOptionsSetOutputFD == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetOutputFD not loaded")
	}
	_bNNSGraphCompileOptionsSetOutputFD(options, fd)
}


var _bNNSGraphCompileOptionsSetOutputPath func(options Bnns_graph_compile_options_t, path *byte)

// BNNSGraphCompileOptionsSetOutputPath sets the option for graph compilation to generate the graph object directly to the specified file.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOutputPath(_:_:)
func BNNSGraphCompileOptionsSetOutputPath(options Bnns_graph_compile_options_t, path *byte) {
	if _bNNSGraphCompileOptionsSetOutputPath == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetOutputPath not loaded")
	}
	_bNNSGraphCompileOptionsSetOutputPath(options, path)
}


var _bNNSGraphCompileOptionsSetTargetSingleThread func(options Bnns_graph_compile_options_t, value bool)

// BNNSGraphCompileOptionsSetTargetSingleThread sets the option for the compiled graph to execute on a single thread.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetTargetSingleThread(_:_:)
func BNNSGraphCompileOptionsSetTargetSingleThread(options Bnns_graph_compile_options_t, value bool) {
	if _bNNSGraphCompileOptionsSetTargetSingleThread == nil {
		panic("Accelerate: symbol BNNSGraphCompileOptionsSetTargetSingleThread not loaded")
	}
	_bNNSGraphCompileOptionsSetTargetSingleThread(options, value)
}


var _bNNSGraphContextDestroy func(context Bnns_graph_context_t)

// BNNSGraphContextDestroy destroys the specified graph context.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextDestroy(_:)
func BNNSGraphContextDestroy(context Bnns_graph_context_t) {
	if _bNNSGraphContextDestroy == nil {
		panic("Accelerate: symbol BNNSGraphContextDestroy not loaded")
	}
	_bNNSGraphContextDestroy(context)
}


var _bNNSGraphContextEnableNanAndInfChecks func(context Bnns_graph_context_t, enable_check_for_nans_inf bool)

// BNNSGraphContextEnableNanAndInfChecks specifies that the context checks intermediate tensors for NaNs and infinities.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextEnableNanAndInfChecks(_:_:)
func BNNSGraphContextEnableNanAndInfChecks(context Bnns_graph_context_t, enable_check_for_nans_inf bool) {
	if _bNNSGraphContextEnableNanAndInfChecks == nil {
		panic("Accelerate: symbol BNNSGraphContextEnableNanAndInfChecks not loaded")
	}
	_bNNSGraphContextEnableNanAndInfChecks(context, enable_check_for_nans_inf)
}


var _bNNSGraphContextExecute func(context Bnns_graph_context_t, function *byte, argument_count uintptr, arguments *Bnns_graph_argument_t, workspace_size uintptr, workspace *byte) int

// BNNSGraphContextExecute executes the specified function with the given context.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextExecute(_:_:_:_:_:_:)
func BNNSGraphContextExecute(context Bnns_graph_context_t, function *byte, argument_count uintptr, arguments *Bnns_graph_argument_t, workspace_size uintptr, workspace *byte) int {
	if _bNNSGraphContextExecute == nil {
		panic("Accelerate: symbol BNNSGraphContextExecute not loaded")
	}
	return _bNNSGraphContextExecute(context, function, argument_count, arguments, workspace_size, workspace)
}


var _bNNSGraphContextGetTensor func(context Bnns_graph_context_t, function *byte, argument *byte, fill_known_dynamic_shapes bool, tensor *BNNSTensor) int

// BNNSGraphContextGetTensor sets the properties of a tensor for the specified function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextGetTensor(_:_:_:_:_:)
func BNNSGraphContextGetTensor(context Bnns_graph_context_t, function *byte, argument *byte, fill_known_dynamic_shapes bool, tensor *BNNSTensor) int {
	if _bNNSGraphContextGetTensor == nil {
		panic("Accelerate: symbol BNNSGraphContextGetTensor not loaded")
	}
	return _bNNSGraphContextGetTensor(context, function, argument, fill_known_dynamic_shapes, tensor)
}


var _bNNSGraphContextGetWorkspaceSize func(context Bnns_graph_context_t, function *byte) uintptr

// BNNSGraphContextGetWorkspaceSize returns the minimum size, in bytes, of the workspace that graph context execution requires.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextGetWorkspaceSize(_:_:)
func BNNSGraphContextGetWorkspaceSize(context Bnns_graph_context_t, function *byte) uintptr {
	if _bNNSGraphContextGetWorkspaceSize == nil {
		panic("Accelerate: symbol BNNSGraphContextGetWorkspaceSize not loaded")
	}
	return _bNNSGraphContextGetWorkspaceSize(context, function)
}


var _bNNSGraphContextMake func(graph Bnns_graph_t) Bnns_graph_context_t

// BNNSGraphContextMake returns an allocated and initialized graph context from the specified graph.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextMake(_:)
func BNNSGraphContextMake(graph Bnns_graph_t) Bnns_graph_context_t {
	if _bNNSGraphContextMake == nil {
		panic("Accelerate: symbol BNNSGraphContextMake not loaded")
	}
	return _bNNSGraphContextMake(graph)
}


var _bNNSGraphContextMakeStreaming func(graph Bnns_graph_t, function *byte, initial_states_count uintptr, initial_states *BNNSTensor) Bnns_graph_context_t

// BNNSGraphContextMakeStreaming returns an allocated and initialized graph context with streaming support from the specified graph.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextMakeStreaming(_:_:_:_:)
func BNNSGraphContextMakeStreaming(graph Bnns_graph_t, function *byte, initial_states_count uintptr, initial_states *BNNSTensor) Bnns_graph_context_t {
	if _bNNSGraphContextMakeStreaming == nil {
		panic("Accelerate: symbol BNNSGraphContextMakeStreaming not loaded")
	}
	return _bNNSGraphContextMakeStreaming(graph, function, initial_states_count, initial_states)
}


var _bNNSGraphContextSetArgumentType func(context Bnns_graph_context_t, argument_type unsafe.Pointer) int

// BNNSGraphContextSetArgumentType specifies the argument type for a graph context.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetArgumentType(_:_:)
func BNNSGraphContextSetArgumentType(context Bnns_graph_context_t, argument_type unsafe.Pointer) int {
	if _bNNSGraphContextSetArgumentType == nil {
		panic("Accelerate: symbol BNNSGraphContextSetArgumentType not loaded")
	}
	return _bNNSGraphContextSetArgumentType(context, argument_type)
}


var _bNNSGraphContextSetBatchSize func(context Bnns_graph_context_t, function *byte, batch_size uint64) int

// BNNSGraphContextSetBatchSize sets the batch size for a graph.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetBatchSize(_:_:_:)
func BNNSGraphContextSetBatchSize(context Bnns_graph_context_t, function *byte, batch_size uint64) int {
	if _bNNSGraphContextSetBatchSize == nil {
		panic("Accelerate: symbol BNNSGraphContextSetBatchSize not loaded")
	}
	return _bNNSGraphContextSetBatchSize(context, function, batch_size)
}


var _bNNSGraphContextSetDynamicShapes func(context Bnns_graph_context_t, function *byte, shapes_count uintptr, shapes *Bnns_graph_shape_t) int

// BNNSGraphContextSetDynamicShapes specifies the dynamic shapes for a graph and, if possible, infers, the output shapes.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetDynamicShapes(_:_:_:_:)
func BNNSGraphContextSetDynamicShapes(context Bnns_graph_context_t, function *byte, shapes_count uintptr, shapes *Bnns_graph_shape_t) int {
	if _bNNSGraphContextSetDynamicShapes == nil {
		panic("Accelerate: symbol BNNSGraphContextSetDynamicShapes not loaded")
	}
	return _bNNSGraphContextSetDynamicShapes(context, function, shapes_count, shapes)
}



var _bNNSGraphContextSetMessageLogMask func(context Bnns_graph_context_t, log_level_mask uint32) int

// BNNSGraphContextSetMessageLogMask sets mask for log messages that are logged (either via `os_log` or the user specified callback)
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetMessageLogMask(_:_:)
func BNNSGraphContextSetMessageLogMask(context Bnns_graph_context_t, log_level_mask uint32) int {
	if _bNNSGraphContextSetMessageLogMask == nil {
		panic("Accelerate: symbol BNNSGraphContextSetMessageLogMask not loaded")
	}
	return _bNNSGraphContextSetMessageLogMask(context, log_level_mask)
}



var _bNNSGraphContextSetStreamingAdvanceCount func(context Bnns_graph_context_t, advance_count uintptr) int

// BNNSGraphContextSetStreamingAdvanceCount sets the streaming advancement amount for cases with dynamically shaped inputs.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetStreamingAdvanceCount(_:_:)
func BNNSGraphContextSetStreamingAdvanceCount(context Bnns_graph_context_t, advance_count uintptr) int {
	if _bNNSGraphContextSetStreamingAdvanceCount == nil {
		panic("Accelerate: symbol BNNSGraphContextSetStreamingAdvanceCount not loaded")
	}
	return _bNNSGraphContextSetStreamingAdvanceCount(context, advance_count)
}



var _bNNSGraphGetArgumentCount func(graph Bnns_graph_t, function *byte) uintptr

// BNNSGraphGetArgumentCount returns the number of arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentCount(_:_:)
func BNNSGraphGetArgumentCount(graph Bnns_graph_t, function *byte) uintptr {
	if _bNNSGraphGetArgumentCount == nil {
		panic("Accelerate: symbol BNNSGraphGetArgumentCount not loaded")
	}
	return _bNNSGraphGetArgumentCount(graph, function)
}


var _bNNSGraphGetArgumentIntents func(graph Bnns_graph_t, function *byte, argument_intents_count uintptr, argument_intents *uintptr) int

// BNNSGraphGetArgumentIntents extracts the intents of arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentIntents(_:_:_:_:)
func BNNSGraphGetArgumentIntents(graph Bnns_graph_t, function *byte, argument_intents_count uintptr, argument_intents *uintptr) int {
	if _bNNSGraphGetArgumentIntents == nil {
		panic("Accelerate: symbol BNNSGraphGetArgumentIntents not loaded")
	}
	return _bNNSGraphGetArgumentIntents(graph, function, argument_intents_count, argument_intents)
}


var _bNNSGraphGetArgumentInterleaveFactors func(graph Bnns_graph_t, function *byte, argument_count uintptr, argument_interleave *uint16, argument_interleave_counts *uintptr) int

// BNNSGraphGetArgumentInterleaveFactors returns the interleave factors for arguments, if present
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentInterleaveFactors(_:_:_:_:_:)
func BNNSGraphGetArgumentInterleaveFactors(graph Bnns_graph_t, function *byte, argument_count uintptr, argument_interleave *uint16, argument_interleave_counts *uintptr) int {
	if _bNNSGraphGetArgumentInterleaveFactors == nil {
		panic("Accelerate: symbol BNNSGraphGetArgumentInterleaveFactors not loaded")
	}
	return _bNNSGraphGetArgumentInterleaveFactors(graph, function, argument_count, argument_interleave, argument_interleave_counts)
}


var _bNNSGraphGetArgumentNames func(graph Bnns_graph_t, function *byte, argument_names_count uintptr, argument_names *byte) int

// BNNSGraphGetArgumentNames extracts the names of arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentNames(_:_:_:_:)
func BNNSGraphGetArgumentNames(graph Bnns_graph_t, function *byte, argument_names_count uintptr, argument_names *byte) int {
	if _bNNSGraphGetArgumentNames == nil {
		panic("Accelerate: symbol BNNSGraphGetArgumentNames not loaded")
	}
	return _bNNSGraphGetArgumentNames(graph, function, argument_names_count, argument_names)
}


var _bNNSGraphGetArgumentPosition func(graph Bnns_graph_t, function *byte, argument *byte) uintptr

// BNNSGraphGetArgumentPosition returns the index into the arguments array for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentPosition(_:_:_:)
func BNNSGraphGetArgumentPosition(graph Bnns_graph_t, function *byte, argument *byte) uintptr {
	if _bNNSGraphGetArgumentPosition == nil {
		panic("Accelerate: symbol BNNSGraphGetArgumentPosition not loaded")
	}
	return _bNNSGraphGetArgumentPosition(graph, function, argument)
}


var _bNNSGraphGetFunctionCount func(graph Bnns_graph_t) uintptr

// BNNSGraphGetFunctionCount returns the number of callable functions in the specified graph.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetFunctionCount(_:)
func BNNSGraphGetFunctionCount(graph Bnns_graph_t) uintptr {
	if _bNNSGraphGetFunctionCount == nil {
		panic("Accelerate: symbol BNNSGraphGetFunctionCount not loaded")
	}
	return _bNNSGraphGetFunctionCount(graph)
}


var _bNNSGraphGetFunctionNames func(graph Bnns_graph_t, function_name_count uintptr, function_names *byte) int

// BNNSGraphGetFunctionNames extracts the names of callable functions in the graph.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetFunctionNames(_:_:_:)
func BNNSGraphGetFunctionNames(graph Bnns_graph_t, function_name_count uintptr, function_names *byte) int {
	if _bNNSGraphGetFunctionNames == nil {
		panic("Accelerate: symbol BNNSGraphGetFunctionNames not loaded")
	}
	return _bNNSGraphGetFunctionNames(graph, function_name_count, function_names)
}


var _bNNSGraphGetInputCount func(graph Bnns_graph_t, function *byte) uintptr

// BNNSGraphGetInputCount returns the number of input arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetInputCount(_:_:)
func BNNSGraphGetInputCount(graph Bnns_graph_t, function *byte) uintptr {
	if _bNNSGraphGetInputCount == nil {
		panic("Accelerate: symbol BNNSGraphGetInputCount not loaded")
	}
	return _bNNSGraphGetInputCount(graph, function)
}


var _bNNSGraphGetInputNames func(graph Bnns_graph_t, function *byte, input_names_count uintptr, input_names *byte) int

// BNNSGraphGetInputNames extracts the names of input arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetInputNames(_:_:_:_:)
func BNNSGraphGetInputNames(graph Bnns_graph_t, function *byte, input_names_count uintptr, input_names *byte) int {
	if _bNNSGraphGetInputNames == nil {
		panic("Accelerate: symbol BNNSGraphGetInputNames not loaded")
	}
	return _bNNSGraphGetInputNames(graph, function, input_names_count, input_names)
}


var _bNNSGraphGetOutputCount func(graph Bnns_graph_t, function *byte) uintptr

// BNNSGraphGetOutputCount returns the number of output arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetOutputCount(_:_:)
func BNNSGraphGetOutputCount(graph Bnns_graph_t, function *byte) uintptr {
	if _bNNSGraphGetOutputCount == nil {
		panic("Accelerate: symbol BNNSGraphGetOutputCount not loaded")
	}
	return _bNNSGraphGetOutputCount(graph, function)
}


var _bNNSGraphGetOutputNames func(graph Bnns_graph_t, function *byte, output_names_count uintptr, output_names *byte) int

// BNNSGraphGetOutputNames extracts the names of output arguments for the given function argument.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetOutputNames(_:_:_:_:)
func BNNSGraphGetOutputNames(graph Bnns_graph_t, function *byte, output_names_count uintptr, output_names *byte) int {
	if _bNNSGraphGetOutputNames == nil {
		panic("Accelerate: symbol BNNSGraphGetOutputNames not loaded")
	}
	return _bNNSGraphGetOutputNames(graph, function, output_names_count, output_names)
}


var _bNNSGraphTensorFillStrides func(graph Bnns_graph_t, function *byte, argument *byte, tensor *BNNSTensor) int

// BNNSGraphTensorFillStrides sets the stride of the specifed tensor for compatibility with the given model’s input or output argument based on its current shape.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSGraphTensorFillStrides(_:_:_:_:)
func BNNSGraphTensorFillStrides(graph Bnns_graph_t, function *byte, argument *byte, tensor *BNNSTensor) int {
	if _bNNSGraphTensorFillStrides == nil {
		panic("Accelerate: symbol BNNSGraphTensorFillStrides not loaded")
	}
	return _bNNSGraphTensorFillStrides(graph, function, argument, tensor)
}


var _bNNSLossFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, labels unsafe.Pointer, labels_stride uintptr, weights unsafe.Pointer, weights_size uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int

// BNNSLossFilterApplyBackwardBatch applies a loss filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSLossFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSLossFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, labels unsafe.Pointer, labels_stride uintptr, weights unsafe.Pointer, weights_size uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int {
	if _bNNSLossFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSLossFilterApplyBackwardBatch not loaded")
	}
	return _bNNSLossFilterApplyBackwardBatch(filter, batch_size, in, in_stride, in_delta, in_delta_stride, labels, labels_stride, weights, weights_size, out_delta, out_delta_stride)
}


var _bNNSLossFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, labels unsafe.Pointer, labels_stride uintptr, weights unsafe.Pointer, weights_size uintptr, out unsafe.Pointer, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr) int

// BNNSLossFilterApplyBatch applies a loss filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSLossFilterApplyBatch(_:_:_:_:_:_:_:_:_:_:_:)
func BNNSLossFilterApplyBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, labels unsafe.Pointer, labels_stride uintptr, weights unsafe.Pointer, weights_size uintptr, out unsafe.Pointer, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr) int {
	if _bNNSLossFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSLossFilterApplyBatch not loaded")
	}
	return _bNNSLossFilterApplyBatch(filter, batch_size, in, in_stride, labels, labels_stride, weights, weights_size, out, in_delta, in_delta_stride)
}


var _bNNSMatMul func(transA bool, transB bool, alpha float32, inputA *BNNSNDArrayDescriptor, inputB *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, workspace unsafe.Pointer, filter_params *BNNSFilterParameters) int

// BNNSMatMul applies a matrix multiplication operation directly to two input matrices.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSMatMul(_:_:_:_:_:_:_:_:)
func BNNSMatMul(transA bool, transB bool, alpha float32, inputA *BNNSNDArrayDescriptor, inputB *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, workspace unsafe.Pointer, filter_params *BNNSFilterParameters) int {
	if _bNNSMatMul == nil {
		panic("Accelerate: symbol BNNSMatMul not loaded")
	}
	return _bNNSMatMul(transA, transB, alpha, inputA, inputB, output, workspace, filter_params)
}


var _bNNSMatMulWorkspaceSize func(transA bool, transB bool, alpha float32, inputA *BNNSNDArrayDescriptor, inputB *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSMatMulWorkspaceSize returns the workspace size that a matrix multiply operation requires.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSMatMulWorkspaceSize(_:_:_:_:_:_:_:)
func BNNSMatMulWorkspaceSize(transA bool, transB bool, alpha float32, inputA *BNNSNDArrayDescriptor, inputB *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSMatMulWorkspaceSize == nil {
		panic("Accelerate: symbol BNNSMatMulWorkspaceSize not loaded")
	}
	return _bNNSMatMulWorkspaceSize(transA, transB, alpha, inputA, inputB, output, filter_params)
}


var _bNNSNDArrayFullyConnectedSparsifySparseCOO func(in_dense_shape *BNNSNDArrayDescriptor, in_indices *BNNSNDArrayDescriptor, in_values *BNNSNDArrayDescriptor, out *BNNSNDArrayDescriptor, sparse_params *BNNSSparsityParameters, batch_size uintptr, workspace unsafe.Pointer, workspace_size uintptr, filter_params *BNNSFilterParameters) int

// BNNSNDArrayFullyConnectedSparsifySparseCOO converts a sparse tensor from the standardized coordinate list (COO) layout to a device-specific sparse layout that BNNS fully connected layers use.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayFullyConnectedSparsifySparseCOO(_:_:_:_:_:_:_:_:_:)
func BNNSNDArrayFullyConnectedSparsifySparseCOO(in_dense_shape *BNNSNDArrayDescriptor, in_indices *BNNSNDArrayDescriptor, in_values *BNNSNDArrayDescriptor, out *BNNSNDArrayDescriptor, sparse_params *BNNSSparsityParameters, batch_size uintptr, workspace unsafe.Pointer, workspace_size uintptr, filter_params *BNNSFilterParameters) int {
	if _bNNSNDArrayFullyConnectedSparsifySparseCOO == nil {
		panic("Accelerate: symbol BNNSNDArrayFullyConnectedSparsifySparseCOO not loaded")
	}
	return _bNNSNDArrayFullyConnectedSparsifySparseCOO(in_dense_shape, in_indices, in_values, out, sparse_params, batch_size, workspace, workspace_size, filter_params)
}


var _bNNSNDArrayFullyConnectedSparsifySparseCSR func(in_dense_shape *BNNSNDArrayDescriptor, in_column_indices *BNNSNDArrayDescriptor, in_row_starts *BNNSNDArrayDescriptor, in_values *BNNSNDArrayDescriptor, out *BNNSNDArrayDescriptor, sparse_params *BNNSSparsityParameters, batch_size uintptr, workspace unsafe.Pointer, workspace_size uintptr, filter_params *BNNSFilterParameters) int

// BNNSNDArrayFullyConnectedSparsifySparseCSR converts a sparse tensor from the standardized compressed sparse row (CSR) layout to a device-specific sparse layout that BNNS fully connected layers use.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayFullyConnectedSparsifySparseCSR(_:_:_:_:_:_:_:_:_:_:)
func BNNSNDArrayFullyConnectedSparsifySparseCSR(in_dense_shape *BNNSNDArrayDescriptor, in_column_indices *BNNSNDArrayDescriptor, in_row_starts *BNNSNDArrayDescriptor, in_values *BNNSNDArrayDescriptor, out *BNNSNDArrayDescriptor, sparse_params *BNNSSparsityParameters, batch_size uintptr, workspace unsafe.Pointer, workspace_size uintptr, filter_params *BNNSFilterParameters) int {
	if _bNNSNDArrayFullyConnectedSparsifySparseCSR == nil {
		panic("Accelerate: symbol BNNSNDArrayFullyConnectedSparsifySparseCSR not loaded")
	}
	return _bNNSNDArrayFullyConnectedSparsifySparseCSR(in_dense_shape, in_column_indices, in_row_starts, in_values, out, sparse_params, batch_size, workspace, workspace_size, filter_params)
}


var _bNNSNDArrayGetDataSize func(array *BNNSNDArrayDescriptor) uintptr

// BNNSNDArrayGetDataSize returns the size, in bytes, that an array descriptor requires.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayGetDataSize(_:)
func BNNSNDArrayGetDataSize(array *BNNSNDArrayDescriptor) uintptr {
	if _bNNSNDArrayGetDataSize == nil {
		panic("Accelerate: symbol BNNSNDArrayGetDataSize not loaded")
	}
	return _bNNSNDArrayGetDataSize(array)
}


var _bNNSNearestNeighborsGetInfo func(knn BNNSNearestNeighbors, sample_number int, indices *int, distances unsafe.Pointer) int

// BNNSNearestNeighborsGetInfo calculates the sorted indices and Euclidean distances of the k-nearest neighbors to a specified sample data point.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNearestNeighborsGetInfo(_:_:_:_:)
func BNNSNearestNeighborsGetInfo(knn BNNSNearestNeighbors, sample_number int, indices []int, distances unsafe.Pointer) int {
	if _bNNSNearestNeighborsGetInfo == nil {
		panic("Accelerate: symbol BNNSNearestNeighborsGetInfo not loaded")
	}
	return _bNNSNearestNeighborsGetInfo(knn, sample_number, unsafe.SliceData(indices), distances)
}


var _bNNSNearestNeighborsLoad func(knn BNNSNearestNeighbors, n_new_samples uint, data_ptr unsafe.Pointer) int

// BNNSNearestNeighborsLoad adds new sample data to a k-nearest neighbors object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNearestNeighborsLoad(_:_:_:)
func BNNSNearestNeighborsLoad(knn BNNSNearestNeighbors, n_new_samples uint, data_ptr unsafe.Pointer) int {
	if _bNNSNearestNeighborsLoad == nil {
		panic("Accelerate: symbol BNNSNearestNeighborsLoad not loaded")
	}
	return _bNNSNearestNeighborsLoad(knn, n_new_samples, data_ptr)
}


var _bNNSNormalizationFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, beta_delta *BNNSNDArrayDescriptor, gamma_delta *BNNSNDArrayDescriptor) int

// BNNSNormalizationFilterApplyBackwardBatch applies a normalization filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNormalizationFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:)
func BNNSNormalizationFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, beta_delta *BNNSNDArrayDescriptor, gamma_delta *BNNSNDArrayDescriptor) int {
	if _bNNSNormalizationFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSNormalizationFilterApplyBackwardBatch not loaded")
	}
	return _bNNSNormalizationFilterApplyBackwardBatch(filter, batch_size, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, beta_delta, gamma_delta)
}


var _bNNSNormalizationFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int

// BNNSNormalizationFilterApplyBatch applies a normalization filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNormalizationFilterApplyBatch(_:_:_:_:_:_:_:)
func BNNSNormalizationFilterApplyBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, training bool) int {
	if _bNNSNormalizationFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSNormalizationFilterApplyBatch not loaded")
	}
	return _bNNSNormalizationFilterApplyBatch(filter, batch_size, in, in_stride, out, out_stride, training)
}


var _bNNSOptimizerStep func(function unsafe.Pointer, OptimizerAlgFields unsafe.Pointer, number_of_parameters uintptr, parameters *BNNSNDArrayDescriptor, gradients *BNNSNDArrayDescriptor, accumulators *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSOptimizerStep applies a single optimization step to one or more parameters.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerStep(_:_:_:_:_:_:_:)
func BNNSOptimizerStep(function unsafe.Pointer, OptimizerAlgFields unsafe.Pointer, number_of_parameters uintptr, parameters *BNNSNDArrayDescriptor, gradients *BNNSNDArrayDescriptor, accumulators *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSOptimizerStep == nil {
		panic("Accelerate: symbol BNNSOptimizerStep not loaded")
	}
	return _bNNSOptimizerStep(function, OptimizerAlgFields, number_of_parameters, parameters, gradients, accumulators, filter_params)
}


var _bNNSPermuteFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int

// BNNSPermuteFilterApplyBackwardBatch applies a permute filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSPermuteFilterApplyBackwardBatch(_:_:_:_:_:_:)
func BNNSPermuteFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr) int {
	if _bNNSPermuteFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSPermuteFilterApplyBackwardBatch not loaded")
	}
	return _bNNSPermuteFilterApplyBackwardBatch(filter, batch_size, in_delta, in_delta_stride, out_delta, out_delta_stride)
}


var _bNNSPoolingFilterApplyBackwardBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, bias_delta *BNNSNDArrayDescriptor, indices *uintptr, idx_stride uintptr) int

// BNNSPoolingFilterApplyBackwardBatch applies a pooling filter backward to generate gradients.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSPoolingFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSPoolingFilterApplyBackwardBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, bias_delta *BNNSNDArrayDescriptor, indices *uintptr, idx_stride uintptr) int {
	if _bNNSPoolingFilterApplyBackwardBatch == nil {
		panic("Accelerate: symbol BNNSPoolingFilterApplyBackwardBatch not loaded")
	}
	return _bNNSPoolingFilterApplyBackwardBatch(filter, batch_size, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, bias_delta, indices, idx_stride)
}


var _bNNSPoolingFilterApplyBackwardBatchEx func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, bias_delta *BNNSNDArrayDescriptor, indices_data_type unsafe.Pointer, indices unsafe.Pointer, idx_stride uintptr) int

// BNNSPoolingFilterApplyBackwardBatchEx applies a pooling filter backward to generate gradients with support for multiple data types for indices.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSPoolingFilterApplyBackwardBatchEx(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSPoolingFilterApplyBackwardBatchEx(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, in_delta *BNNSNDArrayDescriptor, in_delta_stride uintptr, out unsafe.Pointer, out_stride uintptr, out_delta *BNNSNDArrayDescriptor, out_delta_stride uintptr, bias_delta *BNNSNDArrayDescriptor, indices_data_type unsafe.Pointer, indices unsafe.Pointer, idx_stride uintptr) int {
	if _bNNSPoolingFilterApplyBackwardBatchEx == nil {
		panic("Accelerate: symbol BNNSPoolingFilterApplyBackwardBatchEx not loaded")
	}
	return _bNNSPoolingFilterApplyBackwardBatchEx(filter, batch_size, in, in_stride, in_delta, in_delta_stride, out, out_stride, out_delta, out_delta_stride, bias_delta, indices_data_type, indices, idx_stride)
}


var _bNNSPoolingFilterApplyBatch func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, indices *uintptr, idx_stride uintptr) int

// BNNSPoolingFilterApplyBatch applies a pooling filter to a set of input objects, writing the result to a set of output objects.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSPoolingFilterApplyBatch(_:_:_:_:_:_:_:_:)
func BNNSPoolingFilterApplyBatch(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, indices *uintptr, idx_stride uintptr) int {
	if _bNNSPoolingFilterApplyBatch == nil {
		panic("Accelerate: symbol BNNSPoolingFilterApplyBatch not loaded")
	}
	return _bNNSPoolingFilterApplyBatch(filter, batch_size, in, in_stride, out, out_stride, indices, idx_stride)
}


var _bNNSPoolingFilterApplyBatchEx func(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, indices_data_type unsafe.Pointer, indices unsafe.Pointer, idx_stride uintptr) int

// BNNSPoolingFilterApplyBatchEx applies a pooling filter to a set of input objects with support for multiple data types for indices.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSPoolingFilterApplyBatchEx(_:_:_:_:_:_:_:_:_:)
func BNNSPoolingFilterApplyBatchEx(filter BNNSFilter, batch_size uintptr, in unsafe.Pointer, in_stride uintptr, out unsafe.Pointer, out_stride uintptr, indices_data_type unsafe.Pointer, indices unsafe.Pointer, idx_stride uintptr) int {
	if _bNNSPoolingFilterApplyBatchEx == nil {
		panic("Accelerate: symbol BNNSPoolingFilterApplyBatchEx not loaded")
	}
	return _bNNSPoolingFilterApplyBatchEx(filter, batch_size, in, in_stride, out, out_stride, indices_data_type, indices, idx_stride)
}


var _bNNSRandomFillCategoricalFloat func(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, probabilities *BNNSNDArrayDescriptor, log_probabilities bool) int

// BNNSRandomFillCategoricalFloat fills the specified tensor with random values from the categorical distributions with the given event probabilities.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomFillCategoricalFloat(_:_:_:_:)
func BNNSRandomFillCategoricalFloat(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, probabilities *BNNSNDArrayDescriptor, log_probabilities bool) int {
	if _bNNSRandomFillCategoricalFloat == nil {
		panic("Accelerate: symbol BNNSRandomFillCategoricalFloat not loaded")
	}
	return _bNNSRandomFillCategoricalFloat(generator, desc, probabilities, log_probabilities)
}


var _bNNSRandomFillNormalFloat func(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, mean float32, stddev float32) int

// BNNSRandomFillNormalFloat fills the specified tensor with random floating-point values mapped to a normal distribution.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomFillNormalFloat(_:_:_:_:)
func BNNSRandomFillNormalFloat(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, mean float32, stddev float32) int {
	if _bNNSRandomFillNormalFloat == nil {
		panic("Accelerate: symbol BNNSRandomFillNormalFloat not loaded")
	}
	return _bNNSRandomFillNormalFloat(generator, desc, mean, stddev)
}


var _bNNSRandomFillUniformFloat func(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, a float32, b float32) int

// BNNSRandomFillUniformFloat fills the specified tensor with random floating-point values from the continuous uniform distribution within a range.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomFillUniformFloat(_:_:_:_:)
func BNNSRandomFillUniformFloat(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, a float32, b float32) int {
	if _bNNSRandomFillUniformFloat == nil {
		panic("Accelerate: symbol BNNSRandomFillUniformFloat not loaded")
	}
	return _bNNSRandomFillUniformFloat(generator, desc, a, b)
}


var _bNNSRandomFillUniformInt func(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, a int64, b int64) int

// BNNSRandomFillUniformInt fills the specified tensor with random integer values from the continuous uniform distribution within a range.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomFillUniformInt(_:_:_:_:)
func BNNSRandomFillUniformInt(generator BNNSRandomGenerator, desc *BNNSNDArrayDescriptor, a int64, b int64) int {
	if _bNNSRandomFillUniformInt == nil {
		panic("Accelerate: symbol BNNSRandomFillUniformInt not loaded")
	}
	return _bNNSRandomFillUniformInt(generator, desc, a, b)
}


var _bNNSRandomGeneratorGetState func(generator BNNSRandomGenerator, state_size uintptr, state unsafe.Pointer) int

// BNNSRandomGeneratorGetState returns the state of a random number generator.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomGeneratorGetState(_:_:_:)
func BNNSRandomGeneratorGetState(generator BNNSRandomGenerator, state_size uintptr, state unsafe.Pointer) int {
	if _bNNSRandomGeneratorGetState == nil {
		panic("Accelerate: symbol BNNSRandomGeneratorGetState not loaded")
	}
	return _bNNSRandomGeneratorGetState(generator, state_size, state)
}


var _bNNSRandomGeneratorSetState func(generator BNNSRandomGenerator, state_size uintptr, state unsafe.Pointer) int

// BNNSRandomGeneratorSetState sets the state of a random number generator.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomGeneratorSetState(_:_:_:)
func BNNSRandomGeneratorSetState(generator BNNSRandomGenerator, state_size uintptr, state unsafe.Pointer) int {
	if _bNNSRandomGeneratorSetState == nil {
		panic("Accelerate: symbol BNNSRandomGeneratorSetState not loaded")
	}
	return _bNNSRandomGeneratorSetState(generator, state_size, state)
}


var _bNNSRandomGeneratorStateSize func(generator BNNSRandomGenerator) uintptr

// BNNSRandomGeneratorStateSize returns the state size, in bytes, of a random number generator.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomGeneratorStateSize(_:)
func BNNSRandomGeneratorStateSize(generator BNNSRandomGenerator) uintptr {
	if _bNNSRandomGeneratorStateSize == nil {
		panic("Accelerate: symbol BNNSRandomGeneratorStateSize not loaded")
	}
	return _bNNSRandomGeneratorStateSize(generator)
}


var _bNNSScatter func(axis uintptr, op unsafe.Pointer, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSScatter scatters the elements of a tensor along a single axis.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSScatter(_:_:_:_:_:_:)
func BNNSScatter(axis uintptr, op unsafe.Pointer, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSScatter == nil {
		panic("Accelerate: symbol BNNSScatter not loaded")
	}
	return _bNNSScatter(axis, op, input, indices, output, filter_params)
}


var _bNNSScatterND func(op unsafe.Pointer, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSScatterND scatters the slices of a tensor.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSScatterND(_:_:_:_:_:)
func BNNSScatterND(op unsafe.Pointer, input *BNNSNDArrayDescriptor, indices *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSScatterND == nil {
		panic("Accelerate: symbol BNNSScatterND not loaded")
	}
	return _bNNSScatterND(op, input, indices, output, filter_params)
}


var _bNNSShuffle func(type_ unsafe.Pointer, input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSShuffle rearranges elements in a tensor according to shuffle type.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSShuffle(_:_:_:_:)
func BNNSShuffle(type_ unsafe.Pointer, input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSShuffle == nil {
		panic("Accelerate: symbol BNNSShuffle not loaded")
	}
	return _bNNSShuffle(type_, input, output, filter_params)
}


var _bNNSTensorGetAllocationSize func(tensor *BNNSTensor) uintptr

// BNNSTensorGetAllocationSize returns the minimum allocation size, in bytes, of the specified tensor.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSTensorGetAllocationSize(_:)
func BNNSTensorGetAllocationSize(tensor *BNNSTensor) uintptr {
	if _bNNSTensorGetAllocationSize == nil {
		panic("Accelerate: symbol BNNSTensorGetAllocationSize not loaded")
	}
	return _bNNSTensorGetAllocationSize(tensor)
}


var _bNNSTile func(input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSTile generates an output tensor by tiling an input tensor multiple times.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSTile(_:_:_:)
func BNNSTile(input *BNNSNDArrayDescriptor, output *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSTile == nil {
		panic("Accelerate: symbol BNNSTile not loaded")
	}
	return _bNNSTile(input, output, filter_params)
}


var _bNNSTileBackward func(in_delta *BNNSNDArrayDescriptor, out_delta *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int

// BNNSTileBackward applies a tile filter backward to generate an input gradient.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSTileBackward(_:_:_:)
func BNNSTileBackward(in_delta *BNNSNDArrayDescriptor, out_delta *BNNSNDArrayDescriptor, filter_params *BNNSFilterParameters) int {
	if _bNNSTileBackward == nil {
		panic("Accelerate: symbol BNNSTileBackward not loaded")
	}
	return _bNNSTileBackward(in_delta, out_delta, filter_params)
}


var _bNNSTranspose func(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, axis0 uintptr, axis1 uintptr, filter_params *BNNSFilterParameters) int

// BNNSTranspose transposes a tensor by swapping two of its dimensions.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSTranspose(_:_:_:_:_:)
func BNNSTranspose(dest *BNNSNDArrayDescriptor, src *BNNSNDArrayDescriptor, axis0 uintptr, axis1 uintptr, filter_params *BNNSFilterParameters) int {
	if _bNNSTranspose == nil {
		panic("Accelerate: symbol BNNSTranspose not loaded")
	}
	return _bNNSTranspose(dest, src, axis0, axis1, filter_params)
}


var _setBLASParamErrorProc func(__ErrorProc BLASParamErrorProc)

// SetBLASParamErrorProc sets an error handler function.
//
// See: https://developer.apple.com/documentation/Accelerate/SetBLASParamErrorProc
func SetBLASParamErrorProc(__ErrorProc BLASParamErrorProc) {
	if _setBLASParamErrorProc == nil {
		panic("Accelerate: symbol SetBLASParamErrorProc not loaded")
	}
	_setBLASParamErrorProc(__ErrorProc)
}











var _sparseGetInertia func(Factored SparseOpaqueFactorization_Complex_Double, num_positive *int, num_zero *int, num_negative *int) int

// SparseGetInertia returns the inertia of an LDLT factorization in complex double.
//
// See: https://developer.apple.com/documentation/Accelerate/SparseGetInertia(_:_:_:_:)-2gc7f
func SparseGetInertia(Factored SparseOpaqueFactorization_Complex_Double, num_positive []int, num_zero []int, num_negative []int) int {
	if _sparseGetInertia == nil {
		panic("Accelerate: symbol SparseGetInertia not loaded")
	}
	return _sparseGetInertia(Factored, unsafe.SliceData(num_positive), unsafe.SliceData(num_zero), unsafe.SliceData(num_negative))
}















var _appleblas_dgeadd func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, ALPHA float64, A *float64, LDA int, BETA float64, B *float64, LDB int, C *float64, LDC int)

// Appleblas_dgeadd.
//
// See: https://developer.apple.com/documentation/Accelerate/appleblas_dgeadd(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Appleblas_dgeadd(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, ALPHA float64, A []float64, LDA int, BETA float64, B []float64, LDB int, C []float64, LDC int) {
	if _appleblas_dgeadd == nil {
		panic("Accelerate: symbol appleblas_dgeadd not loaded")
	}
	_appleblas_dgeadd(ORDER, TRANSA, TRANSB, M, N, ALPHA, unsafe.SliceData(A), LDA, BETA, unsafe.SliceData(B), LDB, unsafe.SliceData(C), LDC)
}


var _appleblas_sgeadd func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, ALPHA float32, A *float32, LDA int, BETA float32, B *float32, LDB int, C *float32, LDC int)

// Appleblas_sgeadd.
//
// See: https://developer.apple.com/documentation/Accelerate/appleblas_sgeadd(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Appleblas_sgeadd(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, ALPHA float32, A []float32, LDA int, BETA float32, B []float32, LDB int, C []float32, LDC int) {
	if _appleblas_sgeadd == nil {
		panic("Accelerate: symbol appleblas_sgeadd not loaded")
	}
	_appleblas_sgeadd(ORDER, TRANSA, TRANSB, M, N, ALPHA, unsafe.SliceData(A), LDA, BETA, unsafe.SliceData(B), LDB, unsafe.SliceData(C), LDC)
}


var _catlas_caxpby func(N int, ALPHA uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Catlas_caxpby computes the product of two vectors, scaling each one separately (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_caxpby(_:_:_:_:_:_:_:)
func Catlas_caxpby(N int, ALPHA uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _catlas_caxpby == nil {
		panic("Accelerate: symbol catlas_caxpby not loaded")
	}
	_catlas_caxpby(N, ALPHA, X, INCX, BETA, Y, INCY)
}


var _catlas_cset func(N int, ALPHA uintptr, X uintptr, INCX int)

// Catlas_cset modifies a vector (single-precision complex) in place, setting each element to a given value.
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_cset(_:_:_:_:)
func Catlas_cset(N int, ALPHA uintptr, X uintptr, INCX int) {
	if _catlas_cset == nil {
		panic("Accelerate: symbol catlas_cset not loaded")
	}
	_catlas_cset(N, ALPHA, X, INCX)
}


var _catlas_daxpby func(N int, ALPHA float64, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Catlas_daxpby computes the sum of two vectors, scaling each one separately (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_daxpby(_:_:_:_:_:_:_:)
func Catlas_daxpby(N int, ALPHA float64, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _catlas_daxpby == nil {
		panic("Accelerate: symbol catlas_daxpby not loaded")
	}
	_catlas_daxpby(N, ALPHA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _catlas_dset func(N int, ALPHA float64, X *float64, INCX int)

// Catlas_dset modifies a vector (double-precision) in place, setting each element to a given value.
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_dset(_:_:_:_:)
func Catlas_dset(N int, ALPHA float64, X []float64, INCX int) {
	if _catlas_dset == nil {
		panic("Accelerate: symbol catlas_dset not loaded")
	}
	_catlas_dset(N, ALPHA, unsafe.SliceData(X), INCX)
}


var _catlas_saxpby func(N int, ALPHA float32, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Catlas_saxpby computes the sum of two vectors, scaling each one separately (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_saxpby(_:_:_:_:_:_:_:)
func Catlas_saxpby(N int, ALPHA float32, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _catlas_saxpby == nil {
		panic("Accelerate: symbol catlas_saxpby not loaded")
	}
	_catlas_saxpby(N, ALPHA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _catlas_sset func(N int, ALPHA float32, X *float32, INCX int)

// Catlas_sset modifies a vector (single-precision) in place, setting each element to a given value.
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_sset(_:_:_:_:)
func Catlas_sset(N int, ALPHA float32, X []float32, INCX int) {
	if _catlas_sset == nil {
		panic("Accelerate: symbol catlas_sset not loaded")
	}
	_catlas_sset(N, ALPHA, unsafe.SliceData(X), INCX)
}


var _catlas_zaxpby func(N int, ALPHA uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Catlas_zaxpby computes the sum of two vectors, scaling each one separately (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_zaxpby(_:_:_:_:_:_:_:)
func Catlas_zaxpby(N int, ALPHA uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _catlas_zaxpby == nil {
		panic("Accelerate: symbol catlas_zaxpby not loaded")
	}
	_catlas_zaxpby(N, ALPHA, X, INCX, BETA, Y, INCY)
}


var _catlas_zset func(N int, ALPHA uintptr, X uintptr, INCX int)

// Catlas_zset modifies a vector (double-precision complex) in place, setting each element to a given value.
//
// See: https://developer.apple.com/documentation/Accelerate/catlas_zset(_:_:_:_:)
func Catlas_zset(N int, ALPHA uintptr, X uintptr, INCX int) {
	if _catlas_zset == nil {
		panic("Accelerate: symbol catlas_zset not loaded")
	}
	_catlas_zset(N, ALPHA, X, INCX)
}






var _cblas_caxpy func(N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_caxpy computes a constant times a vector plus a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_caxpy(_:_:_:_:_:_:)
func Cblas_caxpy(N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_caxpy == nil {
		panic("Accelerate: symbol cblas_caxpy not loaded")
	}
	_cblas_caxpy(N, ALPHA, X, INCX, Y, INCY)
}


var _cblas_ccopy func(N int, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_ccopy copies a vector to another vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ccopy(_:_:_:_:_:)
func Cblas_ccopy(N int, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_ccopy == nil {
		panic("Accelerate: symbol cblas_ccopy not loaded")
	}
	_cblas_ccopy(N, X, INCX, Y, INCY)
}


var _cblas_cdotc_sub func(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTC uintptr)

// Cblas_cdotc_sub calculates the dot product of the complex conjugate of a single-precision complex vector with a second single-precision complex vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cdotc_sub(_:_:_:_:_:_:)
func Cblas_cdotc_sub(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTC uintptr) {
	if _cblas_cdotc_sub == nil {
		panic("Accelerate: symbol cblas_cdotc_sub not loaded")
	}
	_cblas_cdotc_sub(N, X, INCX, Y, INCY, DOTC)
}


var _cblas_cdotu_sub func(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTU uintptr)

// Cblas_cdotu_sub computes the dot product of two single-precision complex vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cdotu_sub(_:_:_:_:_:_:)
func Cblas_cdotu_sub(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTU uintptr) {
	if _cblas_cdotu_sub == nil {
		panic("Accelerate: symbol cblas_cdotu_sub not loaded")
	}
	_cblas_cdotu_sub(N, X, INCX, Y, INCY, DOTU)
}


var _cblas_cgbmv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_cgbmv scales a general band matrix, then multiplies by a vector, then adds a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cgbmv(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_cgbmv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_cgbmv == nil {
		panic("Accelerate: symbol cblas_cgbmv not loaded")
	}
	_cblas_cgbmv(ORDER, TRANSA, M, N, KL, KU, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_cgemm func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_cgemm multiplies two matrices (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cgemm(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_cgemm(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_cgemm == nil {
		panic("Accelerate: symbol cblas_cgemm not loaded")
	}
	_cblas_cgemm(ORDER, TRANSA, TRANSB, M, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_cgemv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_cgemv multiplies a matrix by a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cgemv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_cgemv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_cgemv == nil {
		panic("Accelerate: symbol cblas_cgemv not loaded")
	}
	_cblas_cgemv(ORDER, TRANSA, M, N, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_cgerc func(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_cgerc multiplies vector X by the conjugate transpose of vector Y, then adds matrix A (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cgerc(_:_:_:_:_:_:_:_:_:_:)
func Cblas_cgerc(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_cgerc == nil {
		panic("Accelerate: symbol cblas_cgerc not loaded")
	}
	_cblas_cgerc(ORDER, M, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_cgeru func(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_cgeru multiplies vector X by the transpose of vector Y, then adds matrix A (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cgeru(_:_:_:_:_:_:_:_:_:_:)
func Cblas_cgeru(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_cgeru == nil {
		panic("Accelerate: symbol cblas_cgeru not loaded")
	}
	_cblas_cgeru(ORDER, M, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_chbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_chbmv scales a Hermitian band matrix, then multiplies by a vector, then adds a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chbmv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_chbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_chbmv == nil {
		panic("Accelerate: symbol cblas_chbmv not loaded")
	}
	_cblas_chbmv(ORDER, UPLO, N, K, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_chemm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_chemm multiplies two Hermitian matrices (single-precision complex), then adds a third (with scaling).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chemm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_chemm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_chemm == nil {
		panic("Accelerate: symbol cblas_chemm not loaded")
	}
	_cblas_chemm(ORDER, SIDE, UPLO, M, N, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_chemv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_chemv scales and multiplies a Hermitian matrix by a vector, then adds a second (scaled) vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chemv(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_chemv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_chemv == nil {
		panic("Accelerate: symbol cblas_chemv not loaded")
	}
	_cblas_chemv(ORDER, UPLO, N, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_cher func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X uintptr, INCX int, A uintptr, LDA int)

// Cblas_cher hermitian rank 1 update: adds the product of a scaling factor, vector [X], and the conjugate transpose of [X] to matrix [A].
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cher(_:_:_:_:_:_:_:_:)
func Cblas_cher(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X uintptr, INCX int, A uintptr, LDA int) {
	if _cblas_cher == nil {
		panic("Accelerate: symbol cblas_cher not loaded")
	}
	_cblas_cher(ORDER, UPLO, N, ALPHA, X, INCX, A, LDA)
}


var _cblas_cher2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_cher2 hermitian rank 2 update: adds the product of a scaling factor, vector [X], and the conjugate transpose of vector [Y] to the product of the conjugate of the scaling factor, vector [Y], and the conjugate transpose of vector [X], and adds the result to matrix [A].
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cher2(_:_:_:_:_:_:_:_:_:_:)
func Cblas_cher2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_cher2 == nil {
		panic("Accelerate: symbol cblas_cher2 not loaded")
	}
	_cblas_cher2(ORDER, UPLO, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_cher2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA float32, C uintptr, LDC int)

// Cblas_cher2k performs a rank-2k update of a complex Hermitian matrix (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cher2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_cher2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA float32, C uintptr, LDC int) {
	if _cblas_cher2k == nil {
		panic("Accelerate: symbol cblas_cher2k not loaded")
	}
	_cblas_cher2k(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_cherk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A uintptr, LDA int, BETA float32, C uintptr, LDC int)

// Cblas_cherk rank-k update—multiplies a Hermitian matrix by its transpose and adds a second matrix (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cherk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_cherk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A uintptr, LDA int, BETA float32, C uintptr, LDC int) {
	if _cblas_cherk == nil {
		panic("Accelerate: symbol cblas_cherk not loaded")
	}
	_cblas_cherk(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, BETA, C, LDC)
}


var _cblas_chpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, AP uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_chpmv scales a packed hermitian matrix, multiplies it by a vector, and adds a scaled vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chpmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_chpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, AP uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_chpmv == nil {
		panic("Accelerate: symbol cblas_chpmv not loaded")
	}
	_cblas_chpmv(ORDER, UPLO, N, ALPHA, AP, X, INCX, BETA, Y, INCY)
}


var _cblas_chpr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X uintptr, INCX int, A uintptr)

// Cblas_chpr scales and multiplies a vector times its conjugate transpose, then adds a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chpr(_:_:_:_:_:_:_:)
func Cblas_chpr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X uintptr, INCX int, A uintptr) {
	if _cblas_chpr == nil {
		panic("Accelerate: symbol cblas_chpr not loaded")
	}
	_cblas_chpr(ORDER, UPLO, N, ALPHA, X, INCX, A)
}


var _cblas_chpr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, AP uintptr)

// Cblas_chpr2 multiplies a vector times the conjugate transpose of a second vector and vice-versa, sums the results, and adds a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_chpr2(_:_:_:_:_:_:_:_:_:)
func Cblas_chpr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, AP uintptr) {
	if _cblas_chpr2 == nil {
		panic("Accelerate: symbol cblas_chpr2 not loaded")
	}
	_cblas_chpr2(ORDER, UPLO, N, ALPHA, X, INCX, Y, INCY, AP)
}


var _cblas_crotg func(A uintptr, B uintptr, C *float32, S uintptr)

// Cblas_crotg constructs a complex Givens rotation.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_crotg(_:_:_:_:)
func Cblas_crotg(A uintptr, B uintptr, C []float32, S uintptr) {
	if _cblas_crotg == nil {
		panic("Accelerate: symbol cblas_crotg not loaded")
	}
	_cblas_crotg(A, B, unsafe.SliceData(C), S)
}


var _cblas_cscal func(N int, ALPHA uintptr, X uintptr, INCX int)

// Cblas_cscal multiplies each element of a vector by a constant (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cscal(_:_:_:_:)
func Cblas_cscal(N int, ALPHA uintptr, X uintptr, INCX int) {
	if _cblas_cscal == nil {
		panic("Accelerate: symbol cblas_cscal not loaded")
	}
	_cblas_cscal(N, ALPHA, X, INCX)
}


var _cblas_csrot func(N int, X uintptr, INCX int, Y uintptr, INCY int, C float32, S float32)

// Cblas_csrot applies a Givens rotation matrix to a pair of complex vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_csrot(_:_:_:_:_:_:_:)
func Cblas_csrot(N int, X uintptr, INCX int, Y uintptr, INCY int, C float32, S float32) {
	if _cblas_csrot == nil {
		panic("Accelerate: symbol cblas_csrot not loaded")
	}
	_cblas_csrot(N, X, INCX, Y, INCY, C, S)
}


var _cblas_csscal func(N int, ALPHA float32, X uintptr, INCX int)

// Cblas_csscal multiplies each element of a vector by a constant (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_csscal(_:_:_:_:)
func Cblas_csscal(N int, ALPHA float32, X uintptr, INCX int) {
	if _cblas_csscal == nil {
		panic("Accelerate: symbol cblas_csscal not loaded")
	}
	_cblas_csscal(N, ALPHA, X, INCX)
}


var _cblas_cswap func(N int, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_cswap exchanges the elements of two vectors (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_cswap(_:_:_:_:_:)
func Cblas_cswap(N int, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_cswap == nil {
		panic("Accelerate: symbol cblas_cswap not loaded")
	}
	_cblas_cswap(N, X, INCX, Y, INCY)
}


var _cblas_csymm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_csymm multiplies a matrix by a symmetric matrix (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_csymm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_csymm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_csymm == nil {
		panic("Accelerate: symbol cblas_csymm not loaded")
	}
	_cblas_csymm(ORDER, SIDE, UPLO, M, N, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_csyr2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_csyr2k performs a rank-2k update of a symmetric matrix (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_csyr2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_csyr2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_csyr2k == nil {
		panic("Accelerate: symbol cblas_csyr2k not loaded")
	}
	_cblas_csyr2k(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_csyrk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, BETA uintptr, C uintptr, LDC int)

// Cblas_csyrk rank-k update—multiplies a symmetric matrix by its transpose and adds a second matrix (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_csyrk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_csyrk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_csyrk == nil {
		panic("Accelerate: symbol cblas_csyrk not loaded")
	}
	_cblas_csyrk(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, BETA, C, LDC)
}


var _cblas_ctbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ctbmv scales a triangular band matrix, then multiplies by a vector (single-precision compex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctbmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_ctbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ctbmv == nil {
		panic("Accelerate: symbol cblas_ctbmv not loaded")
	}
	_cblas_ctbmv(ORDER, UPLO, TRANSA, DIAG, N, K, A, LDA, X, INCX)
}


var _cblas_ctbsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ctbsv solves a triangular banded system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctbsv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_ctbsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ctbsv == nil {
		panic("Accelerate: symbol cblas_ctbsv not loaded")
	}
	_cblas_ctbsv(ORDER, UPLO, TRANSA, DIAG, N, K, A, LDA, X, INCX)
}


var _cblas_ctpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int)

// Cblas_ctpmv multiplies a triangular matrix by a vector, then adds a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctpmv(_:_:_:_:_:_:_:_:)
func Cblas_ctpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int) {
	if _cblas_ctpmv == nil {
		panic("Accelerate: symbol cblas_ctpmv not loaded")
	}
	_cblas_ctpmv(ORDER, UPLO, TRANSA, DIAG, N, AP, X, INCX)
}


var _cblas_ctpsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int)

// Cblas_ctpsv solves a packed triangular system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctpsv(_:_:_:_:_:_:_:_:)
func Cblas_ctpsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int) {
	if _cblas_ctpsv == nil {
		panic("Accelerate: symbol cblas_ctpsv not loaded")
	}
	_cblas_ctpsv(ORDER, UPLO, TRANSA, DIAG, N, AP, X, INCX)
}


var _cblas_ctrmm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int)

// Cblas_ctrmm scales a triangular matrix and multiplies it by a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctrmm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ctrmm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int) {
	if _cblas_ctrmm == nil {
		panic("Accelerate: symbol cblas_ctrmm not loaded")
	}
	_cblas_ctrmm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, A, LDA, B, LDB)
}


var _cblas_ctrmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ctrmv multiplies a triangular matrix by a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctrmv(_:_:_:_:_:_:_:_:_:)
func Cblas_ctrmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ctrmv == nil {
		panic("Accelerate: symbol cblas_ctrmv not loaded")
	}
	_cblas_ctrmv(ORDER, UPLO, TRANSA, DIAG, N, A, LDA, X, INCX)
}


var _cblas_ctrsm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int)

// Cblas_ctrsm solves a triangular system of equations with multiple values for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctrsm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ctrsm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int) {
	if _cblas_ctrsm == nil {
		panic("Accelerate: symbol cblas_ctrsm not loaded")
	}
	_cblas_ctrsm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, A, LDA, B, LDB)
}


var _cblas_ctrsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ctrsv solves a triangular system of equations with a single value for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ctrsv(_:_:_:_:_:_:_:_:_:)
func Cblas_ctrsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ctrsv == nil {
		panic("Accelerate: symbol cblas_ctrsv not loaded")
	}
	_cblas_ctrsv(ORDER, UPLO, TRANSA, DIAG, N, A, LDA, X, INCX)
}


var _cblas_dasum func(N int, X *float64, INCX int) float64

// Cblas_dasum computes the sum of the absolute values of elements in a vector (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dasum(_:_:_:)
func Cblas_dasum(N int, X []float64, INCX int) float64 {
	if _cblas_dasum == nil {
		panic("Accelerate: symbol cblas_dasum not loaded")
	}
	return _cblas_dasum(N, unsafe.SliceData(X), INCX)
}


var _cblas_daxpy func(N int, ALPHA float64, X *float64, INCX int, Y *float64, INCY int)

// Cblas_daxpy computes a constant times a vector plus a vector (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_daxpy(_:_:_:_:_:_:)
func Cblas_daxpy(N int, ALPHA float64, X []float64, INCX int, Y []float64, INCY int) {
	if _cblas_daxpy == nil {
		panic("Accelerate: symbol cblas_daxpy not loaded")
	}
	_cblas_daxpy(N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_dcopy func(N int, X *float64, INCX int, Y *float64, INCY int)

// Cblas_dcopy copies a vector to another vector (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dcopy(_:_:_:_:_:)
func Cblas_dcopy(N int, X []float64, INCX int, Y []float64, INCY int) {
	if _cblas_dcopy == nil {
		panic("Accelerate: symbol cblas_dcopy not loaded")
	}
	_cblas_dcopy(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_ddot func(N int, X *float64, INCX int, Y *float64, INCY int) float64

// Cblas_ddot computes the dot product of two vectors (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ddot(_:_:_:_:_:)
func Cblas_ddot(N int, X []float64, INCX int, Y []float64, INCY int) float64 {
	if _cblas_ddot == nil {
		panic("Accelerate: symbol cblas_ddot not loaded")
	}
	return _cblas_ddot(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_dgbmv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA float64, A *float64, LDA int, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Cblas_dgbmv scales a general band matrix, then multiplies by a vector, then adds a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dgbmv(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dgbmv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA float64, A []float64, LDA int, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _cblas_dgbmv == nil {
		panic("Accelerate: symbol cblas_dgbmv not loaded")
	}
	_cblas_dgbmv(ORDER, TRANSA, M, N, KL, KU, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_dgemm func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA float64, A *float64, LDA int, B *float64, LDB int, BETA float64, C *float64, LDC int)

// Cblas_dgemm multiplies two matrices (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dgemm(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dgemm(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA float64, A []float64, LDA int, B []float64, LDB int, BETA float64, C []float64, LDC int) {
	if _cblas_dgemm == nil {
		panic("Accelerate: symbol cblas_dgemm not loaded")
	}
	_cblas_dgemm(ORDER, TRANSA, TRANSB, M, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_dgemv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA float64, A *float64, LDA int, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Cblas_dgemv multiplies a matrix by a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dgemv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dgemv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA float64, A []float64, LDA int, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _cblas_dgemv == nil {
		panic("Accelerate: symbol cblas_dgemv not loaded")
	}
	_cblas_dgemv(ORDER, TRANSA, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_dger func(ORDER CBLAS_ORDER, M int, N int, ALPHA float64, X *float64, INCX int, Y *float64, INCY int, A *float64, LDA int)

// Cblas_dger multiplies vector X by the transpose of vector Y, then adds matrix A (double precison).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dger(_:_:_:_:_:_:_:_:_:_:)
func Cblas_dger(ORDER CBLAS_ORDER, M int, N int, ALPHA float64, X []float64, INCX int, Y []float64, INCY int, A []float64, LDA int) {
	if _cblas_dger == nil {
		panic("Accelerate: symbol cblas_dger not loaded")
	}
	_cblas_dger(ORDER, M, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A), LDA)
}


var _cblas_dnrm2 func(N int, X *float64, INCX int) float64

// Cblas_dnrm2 computes the L2 norm (Euclidian length) of a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dnrm2(_:_:_:)
func Cblas_dnrm2(N int, X []float64, INCX int) float64 {
	if _cblas_dnrm2 == nil {
		panic("Accelerate: symbol cblas_dnrm2 not loaded")
	}
	return _cblas_dnrm2(N, unsafe.SliceData(X), INCX)
}


var _cblas_drot func(N int, X *float64, INCX int, Y *float64, INCY int, C float64, S float64)

// Cblas_drot applies a Givens rotation matrix to a pair of vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_drot(_:_:_:_:_:_:_:)
func Cblas_drot(N int, X []float64, INCX int, Y []float64, INCY int, C float64, S float64) {
	if _cblas_drot == nil {
		panic("Accelerate: symbol cblas_drot not loaded")
	}
	_cblas_drot(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, C, S)
}


var _cblas_drotg func(A *float64, B *float64, C *float64, S *float64)

// Cblas_drotg constructs a Givens rotation matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_drotg(_:_:_:_:)
func Cblas_drotg(A []float64, B []float64, C []float64, S []float64) {
	if _cblas_drotg == nil {
		panic("Accelerate: symbol cblas_drotg not loaded")
	}
	_cblas_drotg(unsafe.SliceData(A), unsafe.SliceData(B), unsafe.SliceData(C), unsafe.SliceData(S))
}


var _cblas_drotm func(N int, X *float64, INCX int, Y *float64, INCY int, P *float64)

// Cblas_drotm applies a modified Givens transformation (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_drotm(_:_:_:_:_:_:)
func Cblas_drotm(N int, X []float64, INCX int, Y []float64, INCY int, P []float64) {
	if _cblas_drotm == nil {
		panic("Accelerate: symbol cblas_drotm not loaded")
	}
	_cblas_drotm(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(P))
}


var _cblas_drotmg func(D1 *float64, D2 *float64, B1 *float64, B2 float64, P *float64)

// Cblas_drotmg generates a modified Givens rotation matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_drotmg(_:_:_:_:_:)
func Cblas_drotmg(D1 []float64, D2 []float64, B1 []float64, B2 float64, P []float64) {
	if _cblas_drotmg == nil {
		panic("Accelerate: symbol cblas_drotmg not loaded")
	}
	_cblas_drotmg(unsafe.SliceData(D1), unsafe.SliceData(D2), unsafe.SliceData(B1), B2, unsafe.SliceData(P))
}


var _cblas_dsbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA float64, A *float64, LDA int, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Cblas_dsbmv scales a symmetric band matrix, then multiplies by a vector, then adds a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsbmv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA float64, A []float64, LDA int, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _cblas_dsbmv == nil {
		panic("Accelerate: symbol cblas_dsbmv not loaded")
	}
	_cblas_dsbmv(ORDER, UPLO, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_dscal func(N int, ALPHA float64, X *float64, INCX int)

// Cblas_dscal multiplies each element of a vector by a constant (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dscal(_:_:_:_:)
func Cblas_dscal(N int, ALPHA float64, X []float64, INCX int) {
	if _cblas_dscal == nil {
		panic("Accelerate: symbol cblas_dscal not loaded")
	}
	_cblas_dscal(N, ALPHA, unsafe.SliceData(X), INCX)
}


var _cblas_dsdot func(N int, X *float32, INCX int, Y *float32, INCY int) float64

// Cblas_dsdot computes the double-precision dot product of a pair of single-precision vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsdot(_:_:_:_:_:)
func Cblas_dsdot(N int, X []float32, INCX int, Y []float32, INCY int) float64 {
	if _cblas_dsdot == nil {
		panic("Accelerate: symbol cblas_dsdot not loaded")
	}
	return _cblas_dsdot(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_dspmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, AP *float64, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Cblas_dspmv scales a packed symmetric matrix, then multiplies by a vector, then scales and adds another vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dspmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_dspmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, AP []float64, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _cblas_dspmv == nil {
		panic("Accelerate: symbol cblas_dspmv not loaded")
	}
	_cblas_dspmv(ORDER, UPLO, N, ALPHA, unsafe.SliceData(AP), unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_dspr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X *float64, INCX int, AP *float64)

// Cblas_dspr rank one update: adds a packed symmetric matrix to the product of a scaling factor, a vector, and its transpose (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dspr(_:_:_:_:_:_:_:)
func Cblas_dspr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X []float64, INCX int, AP []float64) {
	if _cblas_dspr == nil {
		panic("Accelerate: symbol cblas_dspr not loaded")
	}
	_cblas_dspr(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(AP))
}


var _cblas_dspr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X *float64, INCX int, Y *float64, INCY int, A *float64)

// Cblas_dspr2 rank two update of a packed symmetric matrix using two vectors (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dspr2(_:_:_:_:_:_:_:_:_:)
func Cblas_dspr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X []float64, INCX int, Y []float64, INCY int, A []float64) {
	if _cblas_dspr2 == nil {
		panic("Accelerate: symbol cblas_dspr2 not loaded")
	}
	_cblas_dspr2(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A))
}


var _cblas_dswap func(N int, X *float64, INCX int, Y *float64, INCY int)

// Cblas_dswap exchanges the elements of two vectors (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dswap(_:_:_:_:_:)
func Cblas_dswap(N int, X []float64, INCX int, Y []float64, INCY int) {
	if _cblas_dswap == nil {
		panic("Accelerate: symbol cblas_dswap not loaded")
	}
	_cblas_dswap(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_dsymm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA float64, A *float64, LDA int, B *float64, LDB int, BETA float64, C *float64, LDC int)

// Cblas_dsymm multiplies a matrix by a symmetric matrix (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsymm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsymm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA float64, A []float64, LDA int, B []float64, LDB int, BETA float64, C []float64, LDC int) {
	if _cblas_dsymm == nil {
		panic("Accelerate: symbol cblas_dsymm not loaded")
	}
	_cblas_dsymm(ORDER, SIDE, UPLO, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_dsymv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, A *float64, LDA int, X *float64, INCX int, BETA float64, Y *float64, INCY int)

// Cblas_dsymv scales a symmetric matrix, multiplies by a vector, then scales and adds another vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsymv(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsymv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, A []float64, LDA int, X []float64, INCX int, BETA float64, Y []float64, INCY int) {
	if _cblas_dsymv == nil {
		panic("Accelerate: symbol cblas_dsymv not loaded")
	}
	_cblas_dsymv(ORDER, UPLO, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_dsyr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X *float64, INCX int, A *float64, LDA int)

// Cblas_dsyr rank one update: adds a symmetric matrix to the product of a scaling factor, a vector, and its transpose (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsyr(_:_:_:_:_:_:_:_:)
func Cblas_dsyr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X []float64, INCX int, A []float64, LDA int) {
	if _cblas_dsyr == nil {
		panic("Accelerate: symbol cblas_dsyr not loaded")
	}
	_cblas_dsyr(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(A), LDA)
}


var _cblas_dsyr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X *float64, INCX int, Y *float64, INCY int, A *float64, LDA int)

// Cblas_dsyr2 rank two update of a symmetric matrix using two vectors (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsyr2(_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsyr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X []float64, INCX int, Y []float64, INCY int, A []float64, LDA int) {
	if _cblas_dsyr2 == nil {
		panic("Accelerate: symbol cblas_dsyr2 not loaded")
	}
	_cblas_dsyr2(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A), LDA)
}


var _cblas_dsyr2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A *float64, LDA int, B *float64, LDB int, BETA float64, C *float64, LDC int)

// Cblas_dsyr2k performs a rank-2k update of a symmetric matrix (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsyr2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsyr2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A []float64, LDA int, B []float64, LDB int, BETA float64, C []float64, LDC int) {
	if _cblas_dsyr2k == nil {
		panic("Accelerate: symbol cblas_dsyr2k not loaded")
	}
	_cblas_dsyr2k(ORDER, UPLO, TRANS, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_dsyrk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A *float64, LDA int, BETA float64, C *float64, LDC int)

// Cblas_dsyrk rank-k update—multiplies a symmetric matrix by its transpose and adds a second matrix (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dsyrk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dsyrk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A []float64, LDA int, BETA float64, C []float64, LDC int) {
	if _cblas_dsyrk == nil {
		panic("Accelerate: symbol cblas_dsyrk not loaded")
	}
	_cblas_dsyrk(ORDER, UPLO, TRANS, N, K, ALPHA, unsafe.SliceData(A), LDA, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_dtbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A *float64, LDA int, X *float64, INCX int)

// Cblas_dtbmv scales a triangular band matrix, then multiplies by a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtbmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_dtbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A []float64, LDA int, X []float64, INCX int) {
	if _cblas_dtbmv == nil {
		panic("Accelerate: symbol cblas_dtbmv not loaded")
	}
	_cblas_dtbmv(ORDER, UPLO, TRANSA, DIAG, N, K, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_dtbsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A *float64, LDA int, X *float64, INCX int)

// Cblas_dtbsv solves a triangular banded system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtbsv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_dtbsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A []float64, LDA int, X []float64, INCX int) {
	if _cblas_dtbsv == nil {
		panic("Accelerate: symbol cblas_dtbsv not loaded")
	}
	_cblas_dtbsv(ORDER, UPLO, TRANSA, DIAG, N, K, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_dtpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP *float64, X *float64, INCX int)

// Cblas_dtpmv multiplies a triangular matrix by a vector, then adds a vector (double precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtpmv(_:_:_:_:_:_:_:_:)
func Cblas_dtpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP []float64, X []float64, INCX int) {
	if _cblas_dtpmv == nil {
		panic("Accelerate: symbol cblas_dtpmv not loaded")
	}
	_cblas_dtpmv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(AP), unsafe.SliceData(X), INCX)
}


var _cblas_dtpsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP *float64, X *float64, INCX int)

// Cblas_dtpsv solves a packed triangular system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtpsv(_:_:_:_:_:_:_:_:)
func Cblas_dtpsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP []float64, X []float64, INCX int) {
	if _cblas_dtpsv == nil {
		panic("Accelerate: symbol cblas_dtpsv not loaded")
	}
	_cblas_dtpsv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(AP), unsafe.SliceData(X), INCX)
}


var _cblas_dtrmm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float64, A *float64, LDA int, B *float64, LDB int)

// Cblas_dtrmm scales a triangular matrix and multiplies it by a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtrmm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dtrmm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float64, A []float64, LDA int, B []float64, LDB int) {
	if _cblas_dtrmm == nil {
		panic("Accelerate: symbol cblas_dtrmm not loaded")
	}
	_cblas_dtrmm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB)
}


var _cblas_dtrmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A *float64, LDA int, X *float64, INCX int)

// Cblas_dtrmv multiplies a triangular matrix by a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtrmv(_:_:_:_:_:_:_:_:_:)
func Cblas_dtrmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A []float64, LDA int, X []float64, INCX int) {
	if _cblas_dtrmv == nil {
		panic("Accelerate: symbol cblas_dtrmv not loaded")
	}
	_cblas_dtrmv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_dtrsm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float64, A *float64, LDA int, B *float64, LDB int)

// Cblas_dtrsm solves a triangular system of equations with multiple values for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtrsm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_dtrsm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float64, A []float64, LDA int, B []float64, LDB int) {
	if _cblas_dtrsm == nil {
		panic("Accelerate: symbol cblas_dtrsm not loaded")
	}
	_cblas_dtrsm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB)
}


var _cblas_dtrsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A *float64, LDA int, X *float64, INCX int)

// Cblas_dtrsv solves a triangular system of equations with a single value for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dtrsv(_:_:_:_:_:_:_:_:_:)
func Cblas_dtrsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A []float64, LDA int, X []float64, INCX int) {
	if _cblas_dtrsv == nil {
		panic("Accelerate: symbol cblas_dtrsv not loaded")
	}
	_cblas_dtrsv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_dzasum func(N int, X uintptr, INCX int) float64

// Cblas_dzasum computes the sum of the absolute values of real and imaginary parts of elements in a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dzasum(_:_:_:)
func Cblas_dzasum(N int, X uintptr, INCX int) float64 {
	if _cblas_dzasum == nil {
		panic("Accelerate: symbol cblas_dzasum not loaded")
	}
	return _cblas_dzasum(N, X, INCX)
}


var _cblas_dznrm2 func(N int, X uintptr, INCX int) float64

// Cblas_dznrm2 computes the unitary norm of a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_dznrm2(_:_:_:)
func Cblas_dznrm2(N int, X uintptr, INCX int) float64 {
	if _cblas_dznrm2 == nil {
		panic("Accelerate: symbol cblas_dznrm2 not loaded")
	}
	return _cblas_dznrm2(N, X, INCX)
}


var _cblas_errprn func(ierr int, info int, form *byte) int

// Cblas_errprn prints an error message.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_errprn
func Cblas_errprn(ierr int, info int, form *byte) int {
	if _cblas_errprn == nil {
		panic("Accelerate: symbol cblas_errprn not loaded")
	}
	return _cblas_errprn(ierr, info, form)
}


var _cblas_icamax func(N int, X uintptr, INCX int) int

// Cblas_icamax returns the index of the element with the largest absolute value in a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_icamax(_:_:_:)
func Cblas_icamax(N int, X uintptr, INCX int) int {
	if _cblas_icamax == nil {
		panic("Accelerate: symbol cblas_icamax not loaded")
	}
	return _cblas_icamax(N, X, INCX)
}


var _cblas_idamax func(N int, X *float64, INCX int) int

// Cblas_idamax returns the index of the element with the largest absolute value in a vector (double-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_idamax(_:_:_:)
func Cblas_idamax(N int, X []float64, INCX int) int {
	if _cblas_idamax == nil {
		panic("Accelerate: symbol cblas_idamax not loaded")
	}
	return _cblas_idamax(N, unsafe.SliceData(X), INCX)
}


var _cblas_isamax func(N int, X *float32, INCX int) int

// Cblas_isamax returns the index of the element with the largest absolute value in a vector (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_isamax(_:_:_:)
func Cblas_isamax(N int, X []float32, INCX int) int {
	if _cblas_isamax == nil {
		panic("Accelerate: symbol cblas_isamax not loaded")
	}
	return _cblas_isamax(N, unsafe.SliceData(X), INCX)
}


var _cblas_izamax func(N int, X uintptr, INCX int) int

// Cblas_izamax returns the index of the element with the largest absolute value in a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_izamax(_:_:_:)
func Cblas_izamax(N int, X uintptr, INCX int) int {
	if _cblas_izamax == nil {
		panic("Accelerate: symbol cblas_izamax not loaded")
	}
	return _cblas_izamax(N, X, INCX)
}


var _cblas_sasum func(N int, X *float32, INCX int) float32

// Cblas_sasum computes the sum of the absolute values of elements in a vector (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sasum(_:_:_:)
func Cblas_sasum(N int, X []float32, INCX int) float32 {
	if _cblas_sasum == nil {
		panic("Accelerate: symbol cblas_sasum not loaded")
	}
	return _cblas_sasum(N, unsafe.SliceData(X), INCX)
}


var _cblas_saxpy func(N int, ALPHA float32, X *float32, INCX int, Y *float32, INCY int)

// Cblas_saxpy computes a constant times a vector plus a vector (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_saxpy(_:_:_:_:_:_:)
func Cblas_saxpy(N int, ALPHA float32, X []float32, INCX int, Y []float32, INCY int) {
	if _cblas_saxpy == nil {
		panic("Accelerate: symbol cblas_saxpy not loaded")
	}
	_cblas_saxpy(N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_scasum func(N int, X uintptr, INCX int) float32

// Cblas_scasum computes the sum of the absolute values of real and imaginary parts of elements in a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_scasum(_:_:_:)
func Cblas_scasum(N int, X uintptr, INCX int) float32 {
	if _cblas_scasum == nil {
		panic("Accelerate: symbol cblas_scasum not loaded")
	}
	return _cblas_scasum(N, X, INCX)
}


var _cblas_scnrm2 func(N int, X uintptr, INCX int) float32

// Cblas_scnrm2 computes the unitary norm of a vector (single-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_scnrm2(_:_:_:)
func Cblas_scnrm2(N int, X uintptr, INCX int) float32 {
	if _cblas_scnrm2 == nil {
		panic("Accelerate: symbol cblas_scnrm2 not loaded")
	}
	return _cblas_scnrm2(N, X, INCX)
}


var _cblas_scopy func(N int, X *float32, INCX int, Y *float32, INCY int)

// Cblas_scopy copies a vector to another vector (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_scopy(_:_:_:_:_:)
func Cblas_scopy(N int, X []float32, INCX int, Y []float32, INCY int) {
	if _cblas_scopy == nil {
		panic("Accelerate: symbol cblas_scopy not loaded")
	}
	_cblas_scopy(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_sdot func(N int, X *float32, INCX int, Y *float32, INCY int) float32

// Cblas_sdot computes the dot product of two vectors (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sdot(_:_:_:_:_:)
func Cblas_sdot(N int, X []float32, INCX int, Y []float32, INCY int) float32 {
	if _cblas_sdot == nil {
		panic("Accelerate: symbol cblas_sdot not loaded")
	}
	return _cblas_sdot(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_sdsdot func(N int, ALPHA float32, X *float32, INCX int, Y *float32, INCY int) float32

// Cblas_sdsdot computes the dot product of two single-precision vectors plus an initial single-precision value.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sdsdot(_:_:_:_:_:_:)
func Cblas_sdsdot(N int, ALPHA float32, X []float32, INCX int, Y []float32, INCY int) float32 {
	if _cblas_sdsdot == nil {
		panic("Accelerate: symbol cblas_sdsdot not loaded")
	}
	return _cblas_sdsdot(N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_sgbmv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA float32, A *float32, LDA int, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Cblas_sgbmv scales a general band matrix, then multiplies by a vector, then adds a vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sgbmv(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_sgbmv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA float32, A []float32, LDA int, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _cblas_sgbmv == nil {
		panic("Accelerate: symbol cblas_sgbmv not loaded")
	}
	_cblas_sgbmv(ORDER, TRANSA, M, N, KL, KU, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_sgemm func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA float32, A *float32, LDA int, B *float32, LDB int, BETA float32, C *float32, LDC int)

// Cblas_sgemm multiplies two matrices (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sgemm(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_sgemm(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA float32, A []float32, LDA int, B []float32, LDB int, BETA float32, C []float32, LDC int) {
	if _cblas_sgemm == nil {
		panic("Accelerate: symbol cblas_sgemm not loaded")
	}
	_cblas_sgemm(ORDER, TRANSA, TRANSB, M, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_sgemv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA float32, A *float32, LDA int, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Cblas_sgemv multiplies a single-precision matrix by a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sgemv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_sgemv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA float32, A []float32, LDA int, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _cblas_sgemv == nil {
		panic("Accelerate: symbol cblas_sgemv not loaded")
	}
	_cblas_sgemv(ORDER, TRANSA, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_sger func(ORDER CBLAS_ORDER, M int, N int, ALPHA float32, X *float32, INCX int, Y *float32, INCY int, A *float32, LDA int)

// Cblas_sger multiplies vector X by the transpose of vector Y, then adds matrix A (single precison).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sger(_:_:_:_:_:_:_:_:_:_:)
func Cblas_sger(ORDER CBLAS_ORDER, M int, N int, ALPHA float32, X []float32, INCX int, Y []float32, INCY int, A []float32, LDA int) {
	if _cblas_sger == nil {
		panic("Accelerate: symbol cblas_sger not loaded")
	}
	_cblas_sger(ORDER, M, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A), LDA)
}


var _cblas_snrm2 func(N int, X *float32, INCX int) float32

// Cblas_snrm2 computes the L2 norm (Euclidian length) of a vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_snrm2(_:_:_:)
func Cblas_snrm2(N int, X []float32, INCX int) float32 {
	if _cblas_snrm2 == nil {
		panic("Accelerate: symbol cblas_snrm2 not loaded")
	}
	return _cblas_snrm2(N, unsafe.SliceData(X), INCX)
}


var _cblas_srot func(N int, X *float32, INCX int, Y *float32, INCY int, C float32, S float32)

// Cblas_srot applies a Givens rotation matrix to a pair of vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_srot(_:_:_:_:_:_:_:)
func Cblas_srot(N int, X []float32, INCX int, Y []float32, INCY int, C float32, S float32) {
	if _cblas_srot == nil {
		panic("Accelerate: symbol cblas_srot not loaded")
	}
	_cblas_srot(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, C, S)
}


var _cblas_srotg func(A *float32, B *float32, C *float32, S *float32)

// Cblas_srotg constructs a Givens rotation matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_srotg(_:_:_:_:)
func Cblas_srotg(A []float32, B []float32, C []float32, S []float32) {
	if _cblas_srotg == nil {
		panic("Accelerate: symbol cblas_srotg not loaded")
	}
	_cblas_srotg(unsafe.SliceData(A), unsafe.SliceData(B), unsafe.SliceData(C), unsafe.SliceData(S))
}


var _cblas_srotm func(N int, X *float32, INCX int, Y *float32, INCY int, P *float32)

// Cblas_srotm applies a modified Givens transformation (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_srotm(_:_:_:_:_:_:)
func Cblas_srotm(N int, X []float32, INCX int, Y []float32, INCY int, P []float32) {
	if _cblas_srotm == nil {
		panic("Accelerate: symbol cblas_srotm not loaded")
	}
	_cblas_srotm(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(P))
}


var _cblas_srotmg func(D1 *float32, D2 *float32, B1 *float32, B2 float32, P *float32)

// Cblas_srotmg generates a modified Givens rotation matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_srotmg(_:_:_:_:_:)
func Cblas_srotmg(D1 []float32, D2 []float32, B1 []float32, B2 float32, P []float32) {
	if _cblas_srotmg == nil {
		panic("Accelerate: symbol cblas_srotmg not loaded")
	}
	_cblas_srotmg(unsafe.SliceData(D1), unsafe.SliceData(D2), unsafe.SliceData(B1), B2, unsafe.SliceData(P))
}


var _cblas_ssbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA float32, A *float32, LDA int, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Cblas_ssbmv scales a symmetric band matrix, then multiplies by a vector, then adds a vector (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssbmv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA float32, A []float32, LDA int, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _cblas_ssbmv == nil {
		panic("Accelerate: symbol cblas_ssbmv not loaded")
	}
	_cblas_ssbmv(ORDER, UPLO, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_sscal func(N int, ALPHA float32, X *float32, INCX int)

// Cblas_sscal multiplies each element of a vector by a constant (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sscal(_:_:_:_:)
func Cblas_sscal(N int, ALPHA float32, X []float32, INCX int) {
	if _cblas_sscal == nil {
		panic("Accelerate: symbol cblas_sscal not loaded")
	}
	_cblas_sscal(N, ALPHA, unsafe.SliceData(X), INCX)
}


var _cblas_sspmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, AP *float32, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Cblas_sspmv scales a packed symmetric matrix, then multiplies by a vector, then scales and adds another vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sspmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_sspmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, AP []float32, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _cblas_sspmv == nil {
		panic("Accelerate: symbol cblas_sspmv not loaded")
	}
	_cblas_sspmv(ORDER, UPLO, N, ALPHA, unsafe.SliceData(AP), unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_sspr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X *float32, INCX int, AP *float32)

// Cblas_sspr rank one update: adds a packed symmetric matrix to the product of a scaling factor, a vector, and its transpose (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sspr(_:_:_:_:_:_:_:)
func Cblas_sspr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X []float32, INCX int, AP []float32) {
	if _cblas_sspr == nil {
		panic("Accelerate: symbol cblas_sspr not loaded")
	}
	_cblas_sspr(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(AP))
}


var _cblas_sspr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X *float32, INCX int, Y *float32, INCY int, A *float32)

// Cblas_sspr2 rank two update of a packed symmetric matrix using two vectors (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sspr2(_:_:_:_:_:_:_:_:_:)
func Cblas_sspr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X []float32, INCX int, Y []float32, INCY int, A []float32) {
	if _cblas_sspr2 == nil {
		panic("Accelerate: symbol cblas_sspr2 not loaded")
	}
	_cblas_sspr2(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A))
}


var _cblas_sswap func(N int, X *float32, INCX int, Y *float32, INCY int)

// Cblas_sswap exchanges the elements of two vectors (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_sswap(_:_:_:_:_:)
func Cblas_sswap(N int, X []float32, INCX int, Y []float32, INCY int) {
	if _cblas_sswap == nil {
		panic("Accelerate: symbol cblas_sswap not loaded")
	}
	_cblas_sswap(N, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY)
}


var _cblas_ssymm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA float32, A *float32, LDA int, B *float32, LDB int, BETA float32, C *float32, LDC int)

// Cblas_ssymm multiplies a matrix by a symmetric matrix (single-precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssymm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssymm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA float32, A []float32, LDA int, B []float32, LDB int, BETA float32, C []float32, LDC int) {
	if _cblas_ssymm == nil {
		panic("Accelerate: symbol cblas_ssymm not loaded")
	}
	_cblas_ssymm(ORDER, SIDE, UPLO, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_ssymv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, A *float32, LDA int, X *float32, INCX int, BETA float32, Y *float32, INCY int)

// Cblas_ssymv scales a symmetric matrix, multiplies by a vector, then scales and adds another vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssymv(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssymv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, A []float32, LDA int, X []float32, INCX int, BETA float32, Y []float32, INCY int) {
	if _cblas_ssymv == nil {
		panic("Accelerate: symbol cblas_ssymv not loaded")
	}
	_cblas_ssymv(ORDER, UPLO, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX, BETA, unsafe.SliceData(Y), INCY)
}


var _cblas_ssyr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X *float32, INCX int, A *float32, LDA int)

// Cblas_ssyr rank one update: adds a symmetric matrix to the product of a scaling factor, a vector, and its transpose (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssyr(_:_:_:_:_:_:_:_:)
func Cblas_ssyr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X []float32, INCX int, A []float32, LDA int) {
	if _cblas_ssyr == nil {
		panic("Accelerate: symbol cblas_ssyr not loaded")
	}
	_cblas_ssyr(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(A), LDA)
}


var _cblas_ssyr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X *float32, INCX int, Y *float32, INCY int, A *float32, LDA int)

// Cblas_ssyr2 rank two update of a symmetric matrix using two vectors (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssyr2(_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssyr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float32, X []float32, INCX int, Y []float32, INCY int, A []float32, LDA int) {
	if _cblas_ssyr2 == nil {
		panic("Accelerate: symbol cblas_ssyr2 not loaded")
	}
	_cblas_ssyr2(ORDER, UPLO, N, ALPHA, unsafe.SliceData(X), INCX, unsafe.SliceData(Y), INCY, unsafe.SliceData(A), LDA)
}


var _cblas_ssyr2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A *float32, LDA int, B *float32, LDB int, BETA float32, C *float32, LDC int)

// Cblas_ssyr2k performs a rank-2k update of a symmetric matrix (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssyr2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssyr2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A []float32, LDA int, B []float32, LDB int, BETA float32, C []float32, LDC int) {
	if _cblas_ssyr2k == nil {
		panic("Accelerate: symbol cblas_ssyr2k not loaded")
	}
	_cblas_ssyr2k(ORDER, UPLO, TRANS, N, K, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_ssyrk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A *float32, LDA int, BETA float32, C *float32, LDC int)

// Cblas_ssyrk rank-k update—multiplies a symmetric matrix by its transpose and adds a second matrix (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ssyrk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ssyrk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float32, A []float32, LDA int, BETA float32, C []float32, LDC int) {
	if _cblas_ssyrk == nil {
		panic("Accelerate: symbol cblas_ssyrk not loaded")
	}
	_cblas_ssyrk(ORDER, UPLO, TRANS, N, K, ALPHA, unsafe.SliceData(A), LDA, BETA, unsafe.SliceData(C), LDC)
}


var _cblas_stbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A *float32, LDA int, X *float32, INCX int)

// Cblas_stbmv scales a triangular band matrix, then multiplies by a vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_stbmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_stbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A []float32, LDA int, X []float32, INCX int) {
	if _cblas_stbmv == nil {
		panic("Accelerate: symbol cblas_stbmv not loaded")
	}
	_cblas_stbmv(ORDER, UPLO, TRANSA, DIAG, N, K, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_stbsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A *float32, LDA int, X *float32, INCX int)

// Cblas_stbsv solves a triangular banded system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_stbsv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_stbsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A []float32, LDA int, X []float32, INCX int) {
	if _cblas_stbsv == nil {
		panic("Accelerate: symbol cblas_stbsv not loaded")
	}
	_cblas_stbsv(ORDER, UPLO, TRANSA, DIAG, N, K, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_stpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP *float32, X *float32, INCX int)

// Cblas_stpmv multiplies a triangular matrix by a vector, then adds a vector (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_stpmv(_:_:_:_:_:_:_:_:)
func Cblas_stpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP []float32, X []float32, INCX int) {
	if _cblas_stpmv == nil {
		panic("Accelerate: symbol cblas_stpmv not loaded")
	}
	_cblas_stpmv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(AP), unsafe.SliceData(X), INCX)
}


var _cblas_stpsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP *float32, X *float32, INCX int)

// Cblas_stpsv solves a packed triangular system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_stpsv(_:_:_:_:_:_:_:_:)
func Cblas_stpsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP []float32, X []float32, INCX int) {
	if _cblas_stpsv == nil {
		panic("Accelerate: symbol cblas_stpsv not loaded")
	}
	_cblas_stpsv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(AP), unsafe.SliceData(X), INCX)
}


var _cblas_strmm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float32, A *float32, LDA int, B *float32, LDB int)

// Cblas_strmm scales a triangular matrix and multiplies it by a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_strmm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_strmm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float32, A []float32, LDA int, B []float32, LDB int) {
	if _cblas_strmm == nil {
		panic("Accelerate: symbol cblas_strmm not loaded")
	}
	_cblas_strmm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB)
}


var _cblas_strmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A *float32, LDA int, X *float32, INCX int)

// Cblas_strmv multiplies a triangular matrix by a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_strmv(_:_:_:_:_:_:_:_:_:)
func Cblas_strmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A []float32, LDA int, X []float32, INCX int) {
	if _cblas_strmv == nil {
		panic("Accelerate: symbol cblas_strmv not loaded")
	}
	_cblas_strmv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_strsm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float32, A *float32, LDA int, B *float32, LDB int)

// Cblas_strsm solves a triangular system of equations with multiple values for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_strsm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_strsm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA float32, A []float32, LDA int, B []float32, LDB int) {
	if _cblas_strsm == nil {
		panic("Accelerate: symbol cblas_strsm not loaded")
	}
	_cblas_strsm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, unsafe.SliceData(A), LDA, unsafe.SliceData(B), LDB)
}


var _cblas_strsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A *float32, LDA int, X *float32, INCX int)

// Cblas_strsv solves a triangular system of equations with a single value for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_strsv(_:_:_:_:_:_:_:_:_:)
func Cblas_strsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A []float32, LDA int, X []float32, INCX int) {
	if _cblas_strsv == nil {
		panic("Accelerate: symbol cblas_strsv not loaded")
	}
	_cblas_strsv(ORDER, UPLO, TRANSA, DIAG, N, unsafe.SliceData(A), LDA, unsafe.SliceData(X), INCX)
}


var _cblas_xerbla func(p int, rout *byte, form *byte)

// Cblas_xerbla the default error handler for BLAS routines.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_xerbla
func Cblas_xerbla(p int, rout *byte, form *byte) {
	if _cblas_xerbla == nil {
		panic("Accelerate: symbol cblas_xerbla not loaded")
	}
	_cblas_xerbla(p, rout, form)
}


var _cblas_zaxpy func(N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_zaxpy computes a constant times a vector plus a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zaxpy(_:_:_:_:_:_:)
func Cblas_zaxpy(N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_zaxpy == nil {
		panic("Accelerate: symbol cblas_zaxpy not loaded")
	}
	_cblas_zaxpy(N, ALPHA, X, INCX, Y, INCY)
}


var _cblas_zcopy func(N int, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_zcopy copies a vector to another vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zcopy(_:_:_:_:_:)
func Cblas_zcopy(N int, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_zcopy == nil {
		panic("Accelerate: symbol cblas_zcopy not loaded")
	}
	_cblas_zcopy(N, X, INCX, Y, INCY)
}


var _cblas_zdotc_sub func(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTC uintptr)

// Cblas_zdotc_sub calculates the dot product of the complex conjugate of a double-precision complex vector with a second double-precision complex vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zdotc_sub(_:_:_:_:_:_:)
func Cblas_zdotc_sub(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTC uintptr) {
	if _cblas_zdotc_sub == nil {
		panic("Accelerate: symbol cblas_zdotc_sub not loaded")
	}
	_cblas_zdotc_sub(N, X, INCX, Y, INCY, DOTC)
}


var _cblas_zdotu_sub func(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTU uintptr)

// Cblas_zdotu_sub computes the dot product of two double-precision complex vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zdotu_sub(_:_:_:_:_:_:)
func Cblas_zdotu_sub(N int, X uintptr, INCX int, Y uintptr, INCY int, DOTU uintptr) {
	if _cblas_zdotu_sub == nil {
		panic("Accelerate: symbol cblas_zdotu_sub not loaded")
	}
	_cblas_zdotu_sub(N, X, INCX, Y, INCY, DOTU)
}


var _cblas_zdrot func(N int, X uintptr, INCX int, Y uintptr, INCY int, C float64, S float64)

// Cblas_zdrot applies a Givens rotation matrix to a pair of complex vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zdrot(_:_:_:_:_:_:_:)
func Cblas_zdrot(N int, X uintptr, INCX int, Y uintptr, INCY int, C float64, S float64) {
	if _cblas_zdrot == nil {
		panic("Accelerate: symbol cblas_zdrot not loaded")
	}
	_cblas_zdrot(N, X, INCX, Y, INCY, C, S)
}


var _cblas_zdscal func(N int, ALPHA float64, X uintptr, INCX int)

// Cblas_zdscal multiplies each element of a vector by a constant (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zdscal(_:_:_:_:)
func Cblas_zdscal(N int, ALPHA float64, X uintptr, INCX int) {
	if _cblas_zdscal == nil {
		panic("Accelerate: symbol cblas_zdscal not loaded")
	}
	_cblas_zdscal(N, ALPHA, X, INCX)
}


var _cblas_zgbmv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_zgbmv scales a general band matrix, then multiplies by a vector, then adds a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zgbmv(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zgbmv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, KL int, KU int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_zgbmv == nil {
		panic("Accelerate: symbol cblas_zgbmv not loaded")
	}
	_cblas_zgbmv(ORDER, TRANSA, M, N, KL, KU, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_zgemm func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_zgemm multiplies two matrices (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zgemm(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zgemm(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, TRANSB CBLAS_TRANSPOSE, M int, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_zgemm == nil {
		panic("Accelerate: symbol cblas_zgemm not loaded")
	}
	_cblas_zgemm(ORDER, TRANSA, TRANSB, M, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_zgemv func(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_zgemv multiplies a matrix by a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zgemv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zgemv(ORDER CBLAS_ORDER, TRANSA CBLAS_TRANSPOSE, M int, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_zgemv == nil {
		panic("Accelerate: symbol cblas_zgemv not loaded")
	}
	_cblas_zgemv(ORDER, TRANSA, M, N, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_zgerc func(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_zgerc multiplies vector X by the conjugate transpose of vector Y, then adds matrix A (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zgerc(_:_:_:_:_:_:_:_:_:_:)
func Cblas_zgerc(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_zgerc == nil {
		panic("Accelerate: symbol cblas_zgerc not loaded")
	}
	_cblas_zgerc(ORDER, M, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_zgeru func(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_zgeru multiplies vector X by the transpose of vector Y, then adds matrix A (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zgeru(_:_:_:_:_:_:_:_:_:_:)
func Cblas_zgeru(ORDER CBLAS_ORDER, M int, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_zgeru == nil {
		panic("Accelerate: symbol cblas_zgeru not loaded")
	}
	_cblas_zgeru(ORDER, M, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_zhbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_zhbmv scales a Hermitian band matrix, then multiplies by a vector, then adds a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhbmv(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zhbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, K int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_zhbmv == nil {
		panic("Accelerate: symbol cblas_zhbmv not loaded")
	}
	_cblas_zhbmv(ORDER, UPLO, N, K, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_zhemm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_zhemm multiplies two Hermitian matrices (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhemm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zhemm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_zhemm == nil {
		panic("Accelerate: symbol cblas_zhemm not loaded")
	}
	_cblas_zhemm(ORDER, SIDE, UPLO, M, N, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_zhemv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_zhemv scales and multiplies a Hermitian matrix by a vector, then adds a second (scaled) vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhemv(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zhemv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, A uintptr, LDA int, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_zhemv == nil {
		panic("Accelerate: symbol cblas_zhemv not loaded")
	}
	_cblas_zhemv(ORDER, UPLO, N, ALPHA, A, LDA, X, INCX, BETA, Y, INCY)
}


var _cblas_zher func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X uintptr, INCX int, A uintptr, LDA int)

// Cblas_zher adds the product of a scaling factor, vector [X], and the conjugate transpose of [X] to matrix [A].
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zher(_:_:_:_:_:_:_:_:)
func Cblas_zher(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X uintptr, INCX int, A uintptr, LDA int) {
	if _cblas_zher == nil {
		panic("Accelerate: symbol cblas_zher not loaded")
	}
	_cblas_zher(ORDER, UPLO, N, ALPHA, X, INCX, A, LDA)
}


var _cblas_zher2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int)

// Cblas_zher2 hermitian rank 2 update: adds the product of a scaling factor, vector [X], and the conjugate transpose of vector [Y] to the product of the conjugate of the scaling factor, vector [Y], and the conjugate transpose of vector [X], and adds the result to matrix [A].
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zher2(_:_:_:_:_:_:_:_:_:_:)
func Cblas_zher2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, A uintptr, LDA int) {
	if _cblas_zher2 == nil {
		panic("Accelerate: symbol cblas_zher2 not loaded")
	}
	_cblas_zher2(ORDER, UPLO, N, ALPHA, X, INCX, Y, INCY, A, LDA)
}


var _cblas_zher2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA float64, C uintptr, LDC int)

// Cblas_zher2k performs a rank-2k update of a complex Hermitian matrix (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zher2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zher2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA float64, C uintptr, LDC int) {
	if _cblas_zher2k == nil {
		panic("Accelerate: symbol cblas_zher2k not loaded")
	}
	_cblas_zher2k(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_zherk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A uintptr, LDA int, BETA float64, C uintptr, LDC int)

// Cblas_zherk rank-k update—multiplies a Hermitian matrix by its transpose and adds a second matrix (single precision).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zherk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zherk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA float64, A uintptr, LDA int, BETA float64, C uintptr, LDC int) {
	if _cblas_zherk == nil {
		panic("Accelerate: symbol cblas_zherk not loaded")
	}
	_cblas_zherk(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, BETA, C, LDC)
}


var _cblas_zhpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, AP uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int)

// Cblas_zhpmv scales a packed hermitian matrix, multiplies it by a vector, and adds a scaled vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhpmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_zhpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, AP uintptr, X uintptr, INCX int, BETA uintptr, Y uintptr, INCY int) {
	if _cblas_zhpmv == nil {
		panic("Accelerate: symbol cblas_zhpmv not loaded")
	}
	_cblas_zhpmv(ORDER, UPLO, N, ALPHA, AP, X, INCX, BETA, Y, INCY)
}


var _cblas_zhpr func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X uintptr, INCX int, A uintptr)

// Cblas_zhpr scales and multiplies a vector times its conjugate transpose, then adds a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhpr(_:_:_:_:_:_:_:)
func Cblas_zhpr(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA float64, X uintptr, INCX int, A uintptr) {
	if _cblas_zhpr == nil {
		panic("Accelerate: symbol cblas_zhpr not loaded")
	}
	_cblas_zhpr(ORDER, UPLO, N, ALPHA, X, INCX, A)
}


var _cblas_zhpr2 func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, AP uintptr)

// Cblas_zhpr2 multiplies a vector times the conjugate transpose of a second vector and vice-versa, sums the results, and adds a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zhpr2(_:_:_:_:_:_:_:_:_:)
func Cblas_zhpr2(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, N int, ALPHA uintptr, X uintptr, INCX int, Y uintptr, INCY int, AP uintptr) {
	if _cblas_zhpr2 == nil {
		panic("Accelerate: symbol cblas_zhpr2 not loaded")
	}
	_cblas_zhpr2(ORDER, UPLO, N, ALPHA, X, INCX, Y, INCY, AP)
}


var _cblas_zrotg func(A uintptr, B uintptr, C *float64, S uintptr)

// Cblas_zrotg constructs a complex Givens rotation.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zrotg(_:_:_:_:)
func Cblas_zrotg(A uintptr, B uintptr, C []float64, S uintptr) {
	if _cblas_zrotg == nil {
		panic("Accelerate: symbol cblas_zrotg not loaded")
	}
	_cblas_zrotg(A, B, unsafe.SliceData(C), S)
}


var _cblas_zscal func(N int, ALPHA uintptr, X uintptr, INCX int)

// Cblas_zscal multiplies each element of a vector by a constant (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zscal(_:_:_:_:)
func Cblas_zscal(N int, ALPHA uintptr, X uintptr, INCX int) {
	if _cblas_zscal == nil {
		panic("Accelerate: symbol cblas_zscal not loaded")
	}
	_cblas_zscal(N, ALPHA, X, INCX)
}


var _cblas_zswap func(N int, X uintptr, INCX int, Y uintptr, INCY int)

// Cblas_zswap exchanges the elements of two vectors (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zswap(_:_:_:_:_:)
func Cblas_zswap(N int, X uintptr, INCX int, Y uintptr, INCY int) {
	if _cblas_zswap == nil {
		panic("Accelerate: symbol cblas_zswap not loaded")
	}
	_cblas_zswap(N, X, INCX, Y, INCY)
}


var _cblas_zsymm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_zsymm multiplies a matrix by a symmetric matrix (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zsymm(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zsymm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_zsymm == nil {
		panic("Accelerate: symbol cblas_zsymm not loaded")
	}
	_cblas_zsymm(ORDER, SIDE, UPLO, M, N, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_zsyr2k func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int)

// Cblas_zsyr2k performs a rank-2k update of a symmetric matrix (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zsyr2k(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zsyr2k(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_zsyr2k == nil {
		panic("Accelerate: symbol cblas_zsyr2k not loaded")
	}
	_cblas_zsyr2k(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}


var _cblas_zsyrk func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, BETA uintptr, C uintptr, LDC int)

// Cblas_zsyrk rank-k update—multiplies a symmetric matrix by its transpose and adds a second matrix (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_zsyrk(_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_zsyrk(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANS CBLAS_TRANSPOSE, N int, K int, ALPHA uintptr, A uintptr, LDA int, BETA uintptr, C uintptr, LDC int) {
	if _cblas_zsyrk == nil {
		panic("Accelerate: symbol cblas_zsyrk not loaded")
	}
	_cblas_zsyrk(ORDER, UPLO, TRANS, N, K, ALPHA, A, LDA, BETA, C, LDC)
}


var _cblas_ztbmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ztbmv scales a triangular band matrix, then multiplies by a vector (double-precision complex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztbmv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_ztbmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ztbmv == nil {
		panic("Accelerate: symbol cblas_ztbmv not loaded")
	}
	_cblas_ztbmv(ORDER, UPLO, TRANSA, DIAG, N, K, A, LDA, X, INCX)
}


var _cblas_ztbsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ztbsv solves a triangular banded system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztbsv(_:_:_:_:_:_:_:_:_:_:)
func Cblas_ztbsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, K int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ztbsv == nil {
		panic("Accelerate: symbol cblas_ztbsv not loaded")
	}
	_cblas_ztbsv(ORDER, UPLO, TRANSA, DIAG, N, K, A, LDA, X, INCX)
}


var _cblas_ztpmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int)

// Cblas_ztpmv multiplies a triangular matrix by a vector, then adds a vector (double-precision compex).
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztpmv(_:_:_:_:_:_:_:_:)
func Cblas_ztpmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int) {
	if _cblas_ztpmv == nil {
		panic("Accelerate: symbol cblas_ztpmv not loaded")
	}
	_cblas_ztpmv(ORDER, UPLO, TRANSA, DIAG, N, AP, X, INCX)
}


var _cblas_ztpsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int)

// Cblas_ztpsv solves a packed triangular system of equations.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztpsv(_:_:_:_:_:_:_:_:)
func Cblas_ztpsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, AP uintptr, X uintptr, INCX int) {
	if _cblas_ztpsv == nil {
		panic("Accelerate: symbol cblas_ztpsv not loaded")
	}
	_cblas_ztpsv(ORDER, UPLO, TRANSA, DIAG, N, AP, X, INCX)
}


var _cblas_ztrmm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int)

// Cblas_ztrmm scales a triangular matrix and multiplies it by a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztrmm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ztrmm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int) {
	if _cblas_ztrmm == nil {
		panic("Accelerate: symbol cblas_ztrmm not loaded")
	}
	_cblas_ztrmm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, A, LDA, B, LDB)
}


var _cblas_ztrmv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ztrmv multiplies a triangular matrix by a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztrmv(_:_:_:_:_:_:_:_:_:)
func Cblas_ztrmv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ztrmv == nil {
		panic("Accelerate: symbol cblas_ztrmv not loaded")
	}
	_cblas_ztrmv(ORDER, UPLO, TRANSA, DIAG, N, A, LDA, X, INCX)
}


var _cblas_ztrsm func(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int)

// Cblas_ztrsm solves a triangular system of equations with multiple values for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztrsm(_:_:_:_:_:_:_:_:_:_:_:_:)
func Cblas_ztrsm(ORDER CBLAS_ORDER, SIDE CBLAS_SIDE, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, M int, N int, ALPHA uintptr, A uintptr, LDA int, B uintptr, LDB int) {
	if _cblas_ztrsm == nil {
		panic("Accelerate: symbol cblas_ztrsm not loaded")
	}
	_cblas_ztrsm(ORDER, SIDE, UPLO, TRANSA, DIAG, M, N, ALPHA, A, LDA, B, LDB)
}


var _cblas_ztrsv func(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int)

// Cblas_ztrsv solves a triangular system of equations with a single value for the right side.
//
// See: https://developer.apple.com/documentation/Accelerate/cblas_ztrsv(_:_:_:_:_:_:_:_:_:)
func Cblas_ztrsv(ORDER CBLAS_ORDER, UPLO CBLAS_UPLO, TRANSA CBLAS_TRANSPOSE, DIAG CBLAS_DIAG, N int, A uintptr, LDA int, X uintptr, INCX int) {
	if _cblas_ztrsv == nil {
		panic("Accelerate: symbol cblas_ztrsv not loaded")
	}
	_cblas_ztrsv(ORDER, UPLO, TRANSA, DIAG, N, A, LDA, X, INCX)
}

























































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































var _quadrature_integrate func(__f *Quadrature_integrate_function, __a float64, __b float64, options *Quadrature_integrate_options, status uintptr, abs_error *float64, workspace_size uintptr, workspace uintptr) float64

// Quadrature_integrate computes an approximation to the definite integral of a function on a specified interval.
//
// See: https://developer.apple.com/documentation/Accelerate/quadrature_integrate
func Quadrature_integrate(__f *Quadrature_integrate_function, __a float64, __b float64, options *Quadrature_integrate_options, status uintptr, abs_error []float64, workspace_size uintptr, workspace uintptr) float64 {
	if _quadrature_integrate == nil {
		panic("Accelerate: symbol quadrature_integrate not loaded")
	}
	return _quadrature_integrate(__f, __a, __b, options, status, unsafe.SliceData(abs_error), workspace_size, workspace)
}




























































































































































































































































































































































var _sparse_commit func(A unsafe.Pointer) unsafe.Pointer

// Sparse_commit puts values that you recently added to the matrix into the internal sparse storage format.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_commit(_:)
func Sparse_commit(A unsafe.Pointer) unsafe.Pointer {
	if _sparse_commit == nil {
		panic("Accelerate: symbol sparse_commit not loaded")
	}
	return _sparse_commit(A)
}




















var _sparse_get_matrix_nonzero_count func(A unsafe.Pointer) int

// Sparse_get_matrix_nonzero_count returns the number of nonzero values of a matrix.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_get_matrix_nonzero_count(_:)
func Sparse_get_matrix_nonzero_count(A unsafe.Pointer) int {
	if _sparse_get_matrix_nonzero_count == nil {
		panic("Accelerate: symbol sparse_get_matrix_nonzero_count not loaded")
	}
	return _sparse_get_matrix_nonzero_count(A)
}






var _sparse_get_matrix_property func(A unsafe.Pointer, pname unsafe.Pointer) int

// Sparse_get_matrix_property returns the value of the given property name.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_get_matrix_property(_:_:)
func Sparse_get_matrix_property(A unsafe.Pointer, pname unsafe.Pointer) int {
	if _sparse_get_matrix_property == nil {
		panic("Accelerate: symbol sparse_get_matrix_property not loaded")
	}
	return _sparse_get_matrix_property(A, pname)
}










































var _sparse_matrix_destroy func(A unsafe.Pointer) unsafe.Pointer

// Sparse_matrix_destroy releases any memory associated with the matrix object.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_destroy(_:)
func Sparse_matrix_destroy(A unsafe.Pointer) unsafe.Pointer {
	if _sparse_matrix_destroy == nil {
		panic("Accelerate: symbol sparse_matrix_destroy not loaded")
	}
	return _sparse_matrix_destroy(A)
}










































var _sparse_set_matrix_property func(A unsafe.Pointer, pname unsafe.Pointer) unsafe.Pointer

// Sparse_set_matrix_property sets the given property for a matrix object.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_set_matrix_property(_:_:)
func Sparse_set_matrix_property(A unsafe.Pointer, pname unsafe.Pointer) unsafe.Pointer {
	if _sparse_set_matrix_property == nil {
		panic("Accelerate: symbol sparse_set_matrix_property not loaded")
	}
	return _sparse_set_matrix_property(A, pname)
}





























































































































































































































var _vA1024Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VA1024Shift 1024-bit arithmetic shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vA1024Shift(_:_:_:)
func VA1024Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vA1024Shift == nil {
		panic("Accelerate: symbol vA1024Shift not loaded")
	}
	_vA1024Shift(a, shiftAmount, result)
}



var _vA256Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VA256Shift 256-bit arithmetic shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vA256Shift(_:_:_:)
func VA256Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vA256Shift == nil {
		panic("Accelerate: symbol vA256Shift not loaded")
	}
	_vA256Shift(a, shiftAmount, result)
}


var _vA512Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VA512Shift 512-bit arithmetic shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vA512Shift(_:_:_:)
func VA512Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vA512Shift == nil {
		panic("Accelerate: symbol vA512Shift not loaded")
	}
	_vA512Shift(a, shiftAmount, result)
}





var _vDSP_DCT_Execute func(__Setup uintptr, __Input *float32, __Output *float32)

// VDSP_DCT_Execute calculates the discrete cosine transform for a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Execute
func VDSP_DCT_Execute(__Setup uintptr, __Input []float32, __Output []float32) {
	if _vDSP_DCT_Execute == nil {
		panic("Accelerate: symbol vDSP_DCT_Execute not loaded")
	}
	_vDSP_DCT_Execute(__Setup, unsafe.SliceData(__Input), unsafe.SliceData(__Output))
}





var _vDSP_DFT_Execute func(__Setup uintptr, __Ir *float32, __Ii *float32, __Or *float32, __Oi *float32)

// VDSP_DFT_Execute calculates the discrete single-precision Fourier transform for a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Execute
func VDSP_DFT_Execute(__Setup uintptr, __Ir []float32, __Ii []float32, __Or []float32, __Oi []float32) {
	if _vDSP_DFT_Execute == nil {
		panic("Accelerate: symbol vDSP_DFT_Execute not loaded")
	}
	_vDSP_DFT_Execute(__Setup, unsafe.SliceData(__Ir), unsafe.SliceData(__Ii), unsafe.SliceData(__Or), unsafe.SliceData(__Oi))
}


var _vDSP_DFT_ExecuteD func(__Setup uintptr, __Ir *float64, __Ii *float64, __Or *float64, __Oi *float64)

// VDSP_DFT_ExecuteD calculates the discrete double-precision Fourier transform for a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_ExecuteD
func VDSP_DFT_ExecuteD(__Setup uintptr, __Ir []float64, __Ii []float64, __Or []float64, __Oi []float64) {
	if _vDSP_DFT_ExecuteD == nil {
		panic("Accelerate: symbol vDSP_DFT_ExecuteD not loaded")
	}
	_vDSP_DFT_ExecuteD(__Setup, unsafe.SliceData(__Ir), unsafe.SliceData(__Ii), unsafe.SliceData(__Or), unsafe.SliceData(__Oi))
}













var _vDSP_FFT16_copv func(__Output *float32, __Input *float32, __Direction FFTDirection)

// VDSP_FFT16_copv performs a 16-element FFT on interleaved-complex data.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_FFT16_copv
func VDSP_FFT16_copv(__Output []float32, __Input []float32, __Direction FFTDirection) {
	if _vDSP_FFT16_copv == nil {
		panic("Accelerate: symbol vDSP_FFT16_copv not loaded")
	}
	_vDSP_FFT16_copv(unsafe.SliceData(__Output), unsafe.SliceData(__Input), __Direction)
}


var _vDSP_FFT16_zopv func(__Or *float32, __Oi *float32, __Ir *float32, __Ii *float32, __Direction FFTDirection)

// VDSP_FFT16_zopv performs a 16-element FFT on split-complex data.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_FFT16_zopv
func VDSP_FFT16_zopv(__Or []float32, __Oi []float32, __Ir []float32, __Ii []float32, __Direction FFTDirection) {
	if _vDSP_FFT16_zopv == nil {
		panic("Accelerate: symbol vDSP_FFT16_zopv not loaded")
	}
	_vDSP_FFT16_zopv(unsafe.SliceData(__Or), unsafe.SliceData(__Oi), unsafe.SliceData(__Ir), unsafe.SliceData(__Ii), __Direction)
}


var _vDSP_FFT32_copv func(__Output *float32, __Input *float32, __Direction FFTDirection)

// VDSP_FFT32_copv performs a 32-element FFT on interleaved-complex data.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_FFT32_copv
func VDSP_FFT32_copv(__Output []float32, __Input []float32, __Direction FFTDirection) {
	if _vDSP_FFT32_copv == nil {
		panic("Accelerate: symbol vDSP_FFT32_copv not loaded")
	}
	_vDSP_FFT32_copv(unsafe.SliceData(__Output), unsafe.SliceData(__Input), __Direction)
}


var _vDSP_FFT32_zopv func(__Or *float32, __Oi *float32, __Ir *float32, __Ii *float32, __Direction FFTDirection)

// VDSP_FFT32_zopv performs a 32-element FFT on split-complex data.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_FFT32_zopv
func VDSP_FFT32_zopv(__Or []float32, __Oi []float32, __Ir []float32, __Ii []float32, __Direction FFTDirection) {
	if _vDSP_FFT32_zopv == nil {
		panic("Accelerate: symbol vDSP_FFT32_zopv not loaded")
	}
	_vDSP_FFT32_zopv(unsafe.SliceData(__Or), unsafe.SliceData(__Oi), unsafe.SliceData(__Ir), unsafe.SliceData(__Ii), __Direction)
}










































var _vDSP_destroy_fftsetup func(__setup FFTSetup)

// VDSP_destroy_fftsetup deallocates an existing single-precision FFT setup structure.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_destroy_fftsetup
func VDSP_destroy_fftsetup(__setup FFTSetup) {
	if _vDSP_destroy_fftsetup == nil {
		panic("Accelerate: symbol vDSP_destroy_fftsetup not loaded")
	}
	_vDSP_destroy_fftsetup(__setup)
}


var _vDSP_destroy_fftsetupD func(__setup FFTSetupD)

// VDSP_destroy_fftsetupD deallocates an existing double-precision FFT setup structure.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_destroy_fftsetupD
func VDSP_destroy_fftsetupD(__setup FFTSetupD) {
	if _vDSP_destroy_fftsetupD == nil {
		panic("Accelerate: symbol vDSP_destroy_fftsetupD not loaded")
	}
	_vDSP_destroy_fftsetupD(__setup)
}





























































































































































































































































































































































































































var _vImageAffineWarpD_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16F, flags uint32) int

// VImageAffineWarpD_ARGB16F applies a double-precision affine transformation to a floating-point 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGB16F(_:_:_:_:_:_:)
func VImageAffineWarpD_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageAffineWarpD_ARGB16F == nil {
		panic("Accelerate: symbol vImageAffineWarpD_ARGB16F not loaded")
	}
	return _vImageAffineWarpD_ARGB16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16S, flags uint32) int

// VImageAffineWarpD_ARGB16S applies a double-precision affine transformation to a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGB16S(_:_:_:_:_:_:)
func VImageAffineWarpD_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageAffineWarpD_ARGB16S == nil {
		panic("Accelerate: symbol vImageAffineWarpD_ARGB16S not loaded")
	}
	return _vImageAffineWarpD_ARGB16S(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16U, flags uint32) int

// VImageAffineWarpD_ARGB16U applies a double-precision affine transformation to an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGB16U(_:_:_:_:_:_:)
func VImageAffineWarpD_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageAffineWarpD_ARGB16U == nil {
		panic("Accelerate: symbol vImageAffineWarpD_ARGB16U not loaded")
	}
	return _vImageAffineWarpD_ARGB16U(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_8888, flags uint32) int

// VImageAffineWarpD_ARGB8888 applies a double-precision affine transformation to an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGB8888(_:_:_:_:_:_:)
func VImageAffineWarpD_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_8888, flags uint32) int {
	if _vImageAffineWarpD_ARGB8888 == nil {
		panic("Accelerate: symbol vImageAffineWarpD_ARGB8888 not loaded")
	}
	return _vImageAffineWarpD_ARGB8888(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_FFFF, flags uint32) int

// VImageAffineWarpD_ARGBFFFF applies a double-precision affine transformation to a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGBFFFF(_:_:_:_:_:_:)
func VImageAffineWarpD_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_FFFF, flags uint32) int {
	if _vImageAffineWarpD_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageAffineWarpD_ARGBFFFF not loaded")
	}
	return _vImageAffineWarpD_ARGBFFFF(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_16F16F, flags uint32) int

// VImageAffineWarpD_CbCr16F applies a double-precision affine transformation to a floating-point 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_CbCr16F(_:_:_:_:_:_:)
func VImageAffineWarpD_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_16F16F, flags uint32) int {
	if _vImageAffineWarpD_CbCr16F == nil {
		panic("Accelerate: symbol vImageAffineWarpD_CbCr16F not loaded")
	}
	return _vImageAffineWarpD_CbCr16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_16F, flags uint32) int

// VImageAffineWarpD_Planar16F applies a double-precision affine transformation to a floating-point 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_Planar16F(_:_:_:_:_:_:)
func VImageAffineWarpD_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_16F, flags uint32) int {
	if _vImageAffineWarpD_Planar16F == nil {
		panic("Accelerate: symbol vImageAffineWarpD_Planar16F not loaded")
	}
	return _vImageAffineWarpD_Planar16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_8, flags uint32) int

// VImageAffineWarpD_Planar8 applies a double-precision affine transformation to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_Planar8(_:_:_:_:_:_:)
func VImageAffineWarpD_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_8, flags uint32) int {
	if _vImageAffineWarpD_Planar8 == nil {
		panic("Accelerate: symbol vImageAffineWarpD_Planar8 not loaded")
	}
	return _vImageAffineWarpD_Planar8(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarpD_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_F, flags uint32) int

// VImageAffineWarpD_PlanarF applies a double-precision affine transformation to a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_PlanarF(_:_:_:_:_:_:)
func VImageAffineWarpD_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform_Double, backColor Pixel_F, flags uint32) int {
	if _vImageAffineWarpD_PlanarF == nil {
		panic("Accelerate: symbol vImageAffineWarpD_PlanarF not loaded")
	}
	return _vImageAffineWarpD_PlanarF(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16F, flags uint32) int

// VImageAffineWarp_ARGB16F.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_ARGB16F(_:_:_:_:_:_:)
func VImageAffineWarp_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageAffineWarp_ARGB16F == nil {
		panic("Accelerate: symbol vImageAffineWarp_ARGB16F not loaded")
	}
	return _vImageAffineWarp_ARGB16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16S, flags uint32) int

// VImageAffineWarp_ARGB16S applies a single-precision affine transformation to a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_ARGB16S(_:_:_:_:_:_:)
func VImageAffineWarp_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageAffineWarp_ARGB16S == nil {
		panic("Accelerate: symbol vImageAffineWarp_ARGB16S not loaded")
	}
	return _vImageAffineWarp_ARGB16S(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16U, flags uint32) int

// VImageAffineWarp_ARGB16U applies a single-precision affine transformation to an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_ARGB16U(_:_:_:_:_:_:)
func VImageAffineWarp_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageAffineWarp_ARGB16U == nil {
		panic("Accelerate: symbol vImageAffineWarp_ARGB16U not loaded")
	}
	return _vImageAffineWarp_ARGB16U(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_8888, flags uint32) int

// VImageAffineWarp_ARGB8888 applies a single-precision affine transformation to an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_ARGB8888(_:_:_:_:_:_:)
func VImageAffineWarp_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_8888, flags uint32) int {
	if _vImageAffineWarp_ARGB8888 == nil {
		panic("Accelerate: symbol vImageAffineWarp_ARGB8888 not loaded")
	}
	return _vImageAffineWarp_ARGB8888(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_FFFF, flags uint32) int

// VImageAffineWarp_ARGBFFFF applies a single-precision affine transformation to a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_ARGBFFFF(_:_:_:_:_:_:)
func VImageAffineWarp_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_FFFF, flags uint32) int {
	if _vImageAffineWarp_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageAffineWarp_ARGBFFFF not loaded")
	}
	return _vImageAffineWarp_ARGBFFFF(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_16F16F, flags uint32) int

// VImageAffineWarp_CbCr16F.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_CbCr16F(_:_:_:_:_:_:)
func VImageAffineWarp_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_16F16F, flags uint32) int {
	if _vImageAffineWarp_CbCr16F == nil {
		panic("Accelerate: symbol vImageAffineWarp_CbCr16F not loaded")
	}
	return _vImageAffineWarp_CbCr16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_16F, flags uint32) int

// VImageAffineWarp_Planar16F.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_Planar16F(_:_:_:_:_:_:)
func VImageAffineWarp_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_16F, flags uint32) int {
	if _vImageAffineWarp_Planar16F == nil {
		panic("Accelerate: symbol vImageAffineWarp_Planar16F not loaded")
	}
	return _vImageAffineWarp_Planar16F(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_8, flags uint32) int

// VImageAffineWarp_Planar8 applies a single-precision affine transformation to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_Planar8(_:_:_:_:_:_:)
func VImageAffineWarp_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_8, flags uint32) int {
	if _vImageAffineWarp_Planar8 == nil {
		panic("Accelerate: symbol vImageAffineWarp_Planar8 not loaded")
	}
	return _vImageAffineWarp_Planar8(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAffineWarp_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_F, flags uint32) int

// VImageAffineWarp_PlanarF applies a single-precision affine transformation to a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAffineWarp_PlanarF(_:_:_:_:_:_:)
func VImageAffineWarp_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform *VImage_AffineTransform, backColor Pixel_F, flags uint32) int {
	if _vImageAffineWarp_PlanarF == nil {
		panic("Accelerate: symbol vImageAffineWarp_PlanarF not loaded")
	}
	return _vImageAffineWarp_PlanarF(src, dest, tempBuffer, transform, backColor, flags)
}


var _vImageAlphaBlend_ARGB8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_ARGB8888 performs nonpremultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_ARGB8888(_:_:_:_:)
func VImageAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_ARGB8888 == nil {
		panic("Accelerate: symbol vImageAlphaBlend_ARGB8888 not loaded")
	}
	return _vImageAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
}


var _vImageAlphaBlend_ARGBFFFF func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_ARGBFFFF performs nonpremultiplied alpha compositing of two 32-bit-per-channel, 4-channel ARGB buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_ARGBFFFF(_:_:_:_:)
func VImageAlphaBlend_ARGBFFFF(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageAlphaBlend_ARGBFFFF not loaded")
	}
	return _vImageAlphaBlend_ARGBFFFF(srcTop, srcBottom, dest, flags)
}


var _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888 composites a nonpremultiplied 8-bit-per-channel, ARGB buffer over a premultiplied ARGB buffer and generates a premultiplied result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888(_:_:_:_:)
func VImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888 == nil {
		panic("Accelerate: symbol vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888 not loaded")
	}
	return _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888(srcTop, srcBottom, dest, flags)
}


var _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF composites a nonpremultiplied 32-bit-per-channel, ARGB buffer over a premultiplied ARGB buffer and generates a premultiplied result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF(_:_:_:_:)
func VImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF not loaded")
	}
	return _vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF(srcTop, srcBottom, dest, flags)
}


var _vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8 func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8 composites a nonpremultiplied 8-bit planar buffer over a premultiplied 8-bit planar buffer and generates a premultiplied result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8(_:_:_:_:_:)
func VImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8 == nil {
		panic("Accelerate: symbol vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8 not loaded")
	}
	return _vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8(srcTop, srcTopAlpha, srcBottom, dest, flags)
}


var _vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF composites a nonpremultiplied 32-bit planar buffer over a premultiplied 32-bit planar buffer and generates a premultiplied result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF(_:_:_:_:_:)
func VImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF == nil {
		panic("Accelerate: symbol vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF not loaded")
	}
	return _vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF(srcTop, srcTopAlpha, srcBottom, dest, flags)
}


var _vImageAlphaBlend_Planar8 func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, srcBottomAlpha unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_Planar8 performs nonpremultiplied alpha compositing of two 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_Planar8(_:_:_:_:_:_:_:)
func VImageAlphaBlend_Planar8(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, srcBottomAlpha unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_Planar8 == nil {
		panic("Accelerate: symbol vImageAlphaBlend_Planar8 not loaded")
	}
	return _vImageAlphaBlend_Planar8(srcTop, srcTopAlpha, srcBottom, srcBottomAlpha, alpha, dest, flags)
}


var _vImageAlphaBlend_PlanarF func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, srcBottomAlpha unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageAlphaBlend_PlanarF performs nonpremultiplied alpha compositing of two 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_PlanarF(_:_:_:_:_:_:_:)
func VImageAlphaBlend_PlanarF(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, srcBottomAlpha unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageAlphaBlend_PlanarF == nil {
		panic("Accelerate: symbol vImageAlphaBlend_PlanarF not loaded")
	}
	return _vImageAlphaBlend_PlanarF(srcTop, srcTopAlpha, srcBottom, srcBottomAlpha, alpha, dest, flags)
}


var _vImageBoxConvolve_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8888, flags uint32) int

// VImageBoxConvolve_ARGB8888 applies a box filter to an 8-bit-per-channel, 4-channel interleaved source image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBoxConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageBoxConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageBoxConvolve_ARGB8888 == nil {
		panic("Accelerate: symbol vImageBoxConvolve_ARGB8888 not loaded")
	}
	return _vImageBoxConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageBoxConvolve_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8, flags uint32) int

// VImageBoxConvolve_Planar8 applies a box filter to an 8-bit planar source image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBoxConvolve_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageBoxConvolve_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8, flags uint32) int {
	if _vImageBoxConvolve_Planar8 == nil {
		panic("Accelerate: symbol vImageBoxConvolve_Planar8 not loaded")
	}
	return _vImageBoxConvolve_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageBufferFill_ARGB16F func(dest unsafe.Pointer, color Pixel_ARGB_16F, flags uint32) int

// VImageBufferFill_ARGB16F fills a floating-point 16-bit-per-channel, 4-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_ARGB16F(_:_:_:)
func VImageBufferFill_ARGB16F(dest unsafe.Pointer, color Pixel_ARGB_16F, flags uint32) int {
	if _vImageBufferFill_ARGB16F == nil {
		panic("Accelerate: symbol vImageBufferFill_ARGB16F not loaded")
	}
	return _vImageBufferFill_ARGB16F(dest, color, flags)
}


var _vImageBufferFill_ARGB16S func(dest unsafe.Pointer, color Pixel_ARGB_16S, flags uint32) int

// VImageBufferFill_ARGB16S fills a signed 16-bit-per-channel, 4-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_ARGB16S(_:_:_:)
func VImageBufferFill_ARGB16S(dest unsafe.Pointer, color Pixel_ARGB_16S, flags uint32) int {
	if _vImageBufferFill_ARGB16S == nil {
		panic("Accelerate: symbol vImageBufferFill_ARGB16S not loaded")
	}
	return _vImageBufferFill_ARGB16S(dest, color, flags)
}


var _vImageBufferFill_ARGB16U func(dest unsafe.Pointer, color Pixel_ARGB_16U, flags uint32) int

// VImageBufferFill_ARGB16U fills an unsigned 16-bit-per-channel, 4-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_ARGB16U(_:_:_:)
func VImageBufferFill_ARGB16U(dest unsafe.Pointer, color Pixel_ARGB_16U, flags uint32) int {
	if _vImageBufferFill_ARGB16U == nil {
		panic("Accelerate: symbol vImageBufferFill_ARGB16U not loaded")
	}
	return _vImageBufferFill_ARGB16U(dest, color, flags)
}


var _vImageBufferFill_ARGB8888 func(dest unsafe.Pointer, color Pixel_8888, flags uint32) int

// VImageBufferFill_ARGB8888 fills an 8-bit-per-channel, 4-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_ARGB8888(_:_:_:)
func VImageBufferFill_ARGB8888(dest unsafe.Pointer, color Pixel_8888, flags uint32) int {
	if _vImageBufferFill_ARGB8888 == nil {
		panic("Accelerate: symbol vImageBufferFill_ARGB8888 not loaded")
	}
	return _vImageBufferFill_ARGB8888(dest, color, flags)
}


var _vImageBufferFill_ARGBFFFF func(dest unsafe.Pointer, color Pixel_FFFF, flags uint32) int

// VImageBufferFill_ARGBFFFF fills a floating-point 32-bit-per-channel, 4-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_ARGBFFFF(_:_:_:)
func VImageBufferFill_ARGBFFFF(dest unsafe.Pointer, color Pixel_FFFF, flags uint32) int {
	if _vImageBufferFill_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageBufferFill_ARGBFFFF not loaded")
	}
	return _vImageBufferFill_ARGBFFFF(dest, color, flags)
}


var _vImageBufferFill_CbCr16S func(dest unsafe.Pointer, color Pixel_16S16S, flags uint32) int

// VImageBufferFill_CbCr16S.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_CbCr16S(_:_:_:)
func VImageBufferFill_CbCr16S(dest unsafe.Pointer, color Pixel_16S16S, flags uint32) int {
	if _vImageBufferFill_CbCr16S == nil {
		panic("Accelerate: symbol vImageBufferFill_CbCr16S not loaded")
	}
	return _vImageBufferFill_CbCr16S(dest, color, flags)
}


var _vImageBufferFill_CbCr16U func(dest unsafe.Pointer, color Pixel_16U16U, flags uint32) int

// VImageBufferFill_CbCr16U fills an unsigned 16-bit-per-channel, 2-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_CbCr16U(_:_:_:)
func VImageBufferFill_CbCr16U(dest unsafe.Pointer, color Pixel_16U16U, flags uint32) int {
	if _vImageBufferFill_CbCr16U == nil {
		panic("Accelerate: symbol vImageBufferFill_CbCr16U not loaded")
	}
	return _vImageBufferFill_CbCr16U(dest, color, flags)
}


var _vImageBufferFill_CbCr8 func(dest unsafe.Pointer, color Pixel_88, flags uint32) int

// VImageBufferFill_CbCr8 fills an 8-bit-per-channel, 2-channel interleaved buffer with a specified color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferFill_CbCr8(_:_:_:)
func VImageBufferFill_CbCr8(dest unsafe.Pointer, color Pixel_88, flags uint32) int {
	if _vImageBufferFill_CbCr8 == nil {
		panic("Accelerate: symbol vImageBufferFill_CbCr8 not loaded")
	}
	return _vImageBufferFill_CbCr8(dest, color, flags)
}


var _vImageBuffer_CopyToCVPixelBuffer func(buffer unsafe.Pointer, bufferFormat *VImage_CGImageFormat, cvPixelBuffer corevideo.CVPixelBufferRef, cvImageFormat uintptr, backgroundColor *float64, flags uint32) int

// VImageBuffer_CopyToCVPixelBuffer copies the contents of a vImage buffer to a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_CopyToCVPixelBuffer(_:_:_:_:_:_:)
func VImageBuffer_CopyToCVPixelBuffer(buffer unsafe.Pointer, bufferFormat *VImage_CGImageFormat, cvPixelBuffer corevideo.CVPixelBufferRef, cvImageFormat uintptr, backgroundColor *float64, flags uint32) int {
	if _vImageBuffer_CopyToCVPixelBuffer == nil {
		panic("Accelerate: symbol vImageBuffer_CopyToCVPixelBuffer not loaded")
	}
	return _vImageBuffer_CopyToCVPixelBuffer(buffer, bufferFormat, cvPixelBuffer, cvImageFormat, backgroundColor, flags)
}


var _vImageBuffer_GetSize func(buf unsafe.Pointer) corefoundation.CGSize

// VImageBuffer_GetSize returns the size, in pixels, of a vImage buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_GetSize(_:)
func VImageBuffer_GetSize(buf unsafe.Pointer) corefoundation.CGSize {
	if _vImageBuffer_GetSize == nil {
		panic("Accelerate: symbol vImageBuffer_GetSize not loaded")
	}
	return _vImageBuffer_GetSize(buf)
}


var _vImageBuffer_Init func(buf unsafe.Pointer, height uint, width uint, pixelBits uint32, flags uint32) int

// VImageBuffer_Init initializes a vImage buffer with a specified width, height, and bits per pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_Init(_:_:_:_:_:)
func VImageBuffer_Init(buf unsafe.Pointer, height uint, width uint, pixelBits uint32, flags uint32) int {
	if _vImageBuffer_Init == nil {
		panic("Accelerate: symbol vImageBuffer_Init not loaded")
	}
	return _vImageBuffer_Init(buf, height, width, pixelBits, flags)
}


var _vImageBuffer_InitForCopyFromCVPixelBuffer func(buffers unsafe.Pointer, converter unsafe.Pointer, pixelBuffer corevideo.CVPixelBufferRef, flags uint32) int

// VImageBuffer_InitForCopyFromCVPixelBuffer initializes an array of vImage buffers in the order necessary to copy from a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitForCopyFromCVPixelBuffer(_:_:_:_:)
func VImageBuffer_InitForCopyFromCVPixelBuffer(buffers unsafe.Pointer, converter unsafe.Pointer, pixelBuffer corevideo.CVPixelBufferRef, flags uint32) int {
	if _vImageBuffer_InitForCopyFromCVPixelBuffer == nil {
		panic("Accelerate: symbol vImageBuffer_InitForCopyFromCVPixelBuffer not loaded")
	}
	return _vImageBuffer_InitForCopyFromCVPixelBuffer(buffers, converter, pixelBuffer, flags)
}


var _vImageBuffer_InitForCopyToCVPixelBuffer func(buffers unsafe.Pointer, converter unsafe.Pointer, pixelBuffer corevideo.CVPixelBufferRef, flags uint32) int

// VImageBuffer_InitForCopyToCVPixelBuffer initializes an array of vImage buffers in the order necessary to copy to a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitForCopyToCVPixelBuffer(_:_:_:_:)
func VImageBuffer_InitForCopyToCVPixelBuffer(buffers unsafe.Pointer, converter unsafe.Pointer, pixelBuffer corevideo.CVPixelBufferRef, flags uint32) int {
	if _vImageBuffer_InitForCopyToCVPixelBuffer == nil {
		panic("Accelerate: symbol vImageBuffer_InitForCopyToCVPixelBuffer not loaded")
	}
	return _vImageBuffer_InitForCopyToCVPixelBuffer(buffers, converter, pixelBuffer, flags)
}


var _vImageBuffer_InitWithCGImage func(buf unsafe.Pointer, format *VImage_CGImageFormat, backgroundColor *float64, image coregraphics.CGImageRef, flags uint32) int

// VImageBuffer_InitWithCGImage initializes a vImage buffer with the contents of a Core Graphics image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitWithCGImage(_:_:_:_:_:)
func VImageBuffer_InitWithCGImage(buf unsafe.Pointer, format *VImage_CGImageFormat, backgroundColor *float64, image coregraphics.CGImageRef, flags uint32) int {
	if _vImageBuffer_InitWithCGImage == nil {
		panic("Accelerate: symbol vImageBuffer_InitWithCGImage not loaded")
	}
	return _vImageBuffer_InitWithCGImage(buf, format, backgroundColor, image, flags)
}


var _vImageBuffer_InitWithCVPixelBuffer func(buffer unsafe.Pointer, desiredFormat *VImage_CGImageFormat, cvPixelBuffer corevideo.CVPixelBufferRef, cvImageFormat uintptr, backgroundColor *float64, flags uint32) int

// VImageBuffer_InitWithCVPixelBuffer initializes a vImage buffer with a copy of the contents of a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitWithCVPixelBuffer(_:_:_:_:_:_:)
func VImageBuffer_InitWithCVPixelBuffer(buffer unsafe.Pointer, desiredFormat *VImage_CGImageFormat, cvPixelBuffer corevideo.CVPixelBufferRef, cvImageFormat uintptr, backgroundColor *float64, flags uint32) int {
	if _vImageBuffer_InitWithCVPixelBuffer == nil {
		panic("Accelerate: symbol vImageBuffer_InitWithCVPixelBuffer not loaded")
	}
	return _vImageBuffer_InitWithCVPixelBuffer(buffer, desiredFormat, cvPixelBuffer, cvImageFormat, backgroundColor, flags)
}


var _vImageByteSwap_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageByteSwap_Planar16U byte-swaps an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageByteSwap_Planar16U(_:_:_:)
func VImageByteSwap_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageByteSwap_Planar16U == nil {
		panic("Accelerate: symbol vImageByteSwap_Planar16U not loaded")
	}
	return _vImageByteSwap_Planar16U(src, dest, flags)
}


var _vImageCGImageFormat_GetComponentCount func(format *VImage_CGImageFormat) uint32

// VImageCGImageFormat_GetComponentCount calculates the number of color and alpha channels for a specified image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCGImageFormat_GetComponentCount(_:)
func VImageCGImageFormat_GetComponentCount(format *VImage_CGImageFormat) uint32 {
	if _vImageCGImageFormat_GetComponentCount == nil {
		panic("Accelerate: symbol vImageCGImageFormat_GetComponentCount not loaded")
	}
	return _vImageCGImageFormat_GetComponentCount(format)
}


var _vImageCGImageFormat_IsEqual func(f1 *VImage_CGImageFormat, f2 *VImage_CGImageFormat) bool

// VImageCGImageFormat_IsEqual returns a Boolean value that indicates whether two vImage Core Graphics image formats are equal.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCGImageFormat_IsEqual(_:_:)
func VImageCGImageFormat_IsEqual(f1 *VImage_CGImageFormat, f2 *VImage_CGImageFormat) bool {
	if _vImageCGImageFormat_IsEqual == nil {
		panic("Accelerate: symbol vImageCGImageFormat_IsEqual not loaded")
	}
	return _vImageCGImageFormat_IsEqual(f1, f2)
}


var _vImageCVImageFormat_Copy func(format unsafe.Pointer) objectivec.IObject

// VImageCVImageFormat_Copy returns a mutable copy of an immutable Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_Copy(_:)
func VImageCVImageFormat_Copy(format unsafe.Pointer) objectivec.IObject {
	if _vImageCVImageFormat_Copy == nil {
		panic("Accelerate: symbol vImageCVImageFormat_Copy not loaded")
	}
	return _vImageCVImageFormat_Copy(format)
}




var _vImageCVImageFormat_Create func(imageFormatType uint32, matrix *VImage_ARGBToYpCbCrMatrix, cvImageBufferChromaLocation corefoundation.CFStringRef, baseColorspace coregraphics.CGColorSpaceRef, alphaIsOneHint int) objectivec.IObject

// VImageCVImageFormat_Create creates the description of image encoding in a Core Video pixel buffer from the specified properties.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_Create(_:_:_:_:_:)
func VImageCVImageFormat_Create(imageFormatType uint32, matrix *VImage_ARGBToYpCbCrMatrix, cvImageBufferChromaLocation corefoundation.CFStringRef, baseColorspace coregraphics.CGColorSpaceRef, alphaIsOneHint int) objectivec.IObject {
	if _vImageCVImageFormat_Create == nil {
		panic("Accelerate: symbol vImageCVImageFormat_Create not loaded")
	}
	return _vImageCVImageFormat_Create(imageFormatType, matrix, cvImageBufferChromaLocation, baseColorspace, alphaIsOneHint)
}


var _vImageCVImageFormat_CreateWithCVPixelBuffer func(buffer corevideo.CVPixelBufferRef) objectivec.IObject

// VImageCVImageFormat_CreateWithCVPixelBuffer creates the description of the image encoding in an existing Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_CreateWithCVPixelBuffer(_:)
func VImageCVImageFormat_CreateWithCVPixelBuffer(buffer corevideo.CVPixelBufferRef) objectivec.IObject {
	if _vImageCVImageFormat_CreateWithCVPixelBuffer == nil {
		panic("Accelerate: symbol vImageCVImageFormat_CreateWithCVPixelBuffer not loaded")
	}
	return _vImageCVImageFormat_CreateWithCVPixelBuffer(buffer)
}


var _vImageCVImageFormat_GetAlphaHint func(format unsafe.Pointer) int

// VImageCVImageFormat_GetAlphaHint returns the alpha hint of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetAlphaHint(_:)
func VImageCVImageFormat_GetAlphaHint(format unsafe.Pointer) int {
	if _vImageCVImageFormat_GetAlphaHint == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetAlphaHint not loaded")
	}
	return _vImageCVImageFormat_GetAlphaHint(format)
}


var _vImageCVImageFormat_GetChannelCount func(format unsafe.Pointer) uint32

// VImageCVImageFormat_GetChannelCount returns the number of channels, including alpha, for the Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetChannelCount(_:)
func VImageCVImageFormat_GetChannelCount(format unsafe.Pointer) uint32 {
	if _vImageCVImageFormat_GetChannelCount == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetChannelCount not loaded")
	}
	return _vImageCVImageFormat_GetChannelCount(format)
}




var _vImageCVImageFormat_GetChromaSiting func(format unsafe.Pointer) corefoundation.CFStringRef

// VImageCVImageFormat_GetChromaSiting returns the chrominance siting of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetChromaSiting(_:)
func VImageCVImageFormat_GetChromaSiting(format unsafe.Pointer) corefoundation.CFStringRef {
	if _vImageCVImageFormat_GetChromaSiting == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetChromaSiting not loaded")
	}
	return _vImageCVImageFormat_GetChromaSiting(format)
}


var _vImageCVImageFormat_GetColorSpace func(format unsafe.Pointer) coregraphics.CGColorSpaceRef

// VImageCVImageFormat_GetColorSpace returns the color space of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetColorSpace(_:)
func VImageCVImageFormat_GetColorSpace(format unsafe.Pointer) coregraphics.CGColorSpaceRef {
	if _vImageCVImageFormat_GetColorSpace == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetColorSpace not loaded")
	}
	return _vImageCVImageFormat_GetColorSpace(format)
}



var _vImageCVImageFormat_GetFormatCode func(format unsafe.Pointer) uint32

// VImageCVImageFormat_GetFormatCode returns the four-character code that encodes the pixel format of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetFormatCode(_:)
func VImageCVImageFormat_GetFormatCode(format unsafe.Pointer) uint32 {
	if _vImageCVImageFormat_GetFormatCode == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetFormatCode not loaded")
	}
	return _vImageCVImageFormat_GetFormatCode(format)
}


var _vImageCVImageFormat_GetUserData func(format unsafe.Pointer) unsafe.Pointer

// VImageCVImageFormat_GetUserData returns the user data of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetUserData(_:)
func VImageCVImageFormat_GetUserData(format unsafe.Pointer) unsafe.Pointer {
	if _vImageCVImageFormat_GetUserData == nil {
		panic("Accelerate: symbol vImageCVImageFormat_GetUserData not loaded")
	}
	return _vImageCVImageFormat_GetUserData(format)
}


var _vImageCVImageFormat_Release func(fmt uintptr)

// VImageCVImageFormat_Release releases a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_Release
func VImageCVImageFormat_Release(fmt uintptr) {
	if _vImageCVImageFormat_Release == nil {
		panic("Accelerate: symbol vImageCVImageFormat_Release not loaded")
	}
	_vImageCVImageFormat_Release(fmt)
}


var _vImageCVImageFormat_Retain func(fmt uintptr)

// VImageCVImageFormat_Retain retains a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_Retain
func VImageCVImageFormat_Retain(fmt uintptr) {
	if _vImageCVImageFormat_Retain == nil {
		panic("Accelerate: symbol vImageCVImageFormat_Retain not loaded")
	}
	_vImageCVImageFormat_Retain(fmt)
}


var _vImageCVImageFormat_SetAlphaHint func(format uintptr, alphaIsOne int) int

// VImageCVImageFormat_SetAlphaHint sets the alpha hint of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_SetAlphaHint(_:_:)
func VImageCVImageFormat_SetAlphaHint(format uintptr, alphaIsOne int) int {
	if _vImageCVImageFormat_SetAlphaHint == nil {
		panic("Accelerate: symbol vImageCVImageFormat_SetAlphaHint not loaded")
	}
	return _vImageCVImageFormat_SetAlphaHint(format, alphaIsOne)
}


var _vImageCVImageFormat_SetChromaSiting func(format uintptr, siting corefoundation.CFStringRef) int

// VImageCVImageFormat_SetChromaSiting sets the chrominance siting of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_SetChromaSiting(_:_:)
func VImageCVImageFormat_SetChromaSiting(format uintptr, siting corefoundation.CFStringRef) int {
	if _vImageCVImageFormat_SetChromaSiting == nil {
		panic("Accelerate: symbol vImageCVImageFormat_SetChromaSiting not loaded")
	}
	return _vImageCVImageFormat_SetChromaSiting(format, siting)
}


var _vImageCVImageFormat_SetColorSpace func(format uintptr, colorspace coregraphics.CGColorSpaceRef) int

// VImageCVImageFormat_SetColorSpace sets the color space of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_SetColorSpace(_:_:)
func VImageCVImageFormat_SetColorSpace(format uintptr, colorspace coregraphics.CGColorSpaceRef) int {
	if _vImageCVImageFormat_SetColorSpace == nil {
		panic("Accelerate: symbol vImageCVImageFormat_SetColorSpace not loaded")
	}
	return _vImageCVImageFormat_SetColorSpace(format, colorspace)
}


var _vImageCVImageFormat_SetUserData func(format uintptr, userData unsafe.Pointer) int

// VImageCVImageFormat_SetUserData sets the user data of a Core Video image format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_SetUserData(_:_:_:)
func VImageCVImageFormat_SetUserData(format uintptr, userData unsafe.Pointer) int {
	if _vImageCVImageFormat_SetUserData == nil {
		panic("Accelerate: symbol vImageCVImageFormat_SetUserData not loaded")
	}
	return _vImageCVImageFormat_SetUserData(format, userData)
}


var _vImageClipToAlpha_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_ARGB8888 clamps the values of an 8-bit-per-channel, 4-channel ARGB buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_ARGB8888(_:_:_:)
func VImageClipToAlpha_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_ARGB8888 == nil {
		panic("Accelerate: symbol vImageClipToAlpha_ARGB8888 not loaded")
	}
	return _vImageClipToAlpha_ARGB8888(src, dest, flags)
}


var _vImageClipToAlpha_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_ARGBFFFF clamps the values of a 32-bit-per-channel, 4-channel ARGB buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_ARGBFFFF(_:_:_:)
func VImageClipToAlpha_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageClipToAlpha_ARGBFFFF not loaded")
	}
	return _vImageClipToAlpha_ARGBFFFF(src, dest, flags)
}


var _vImageClipToAlpha_Planar8 func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_Planar8 clamps the values of an 8-bit planar buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_Planar8(_:_:_:_:)
func VImageClipToAlpha_Planar8(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_Planar8 == nil {
		panic("Accelerate: symbol vImageClipToAlpha_Planar8 not loaded")
	}
	return _vImageClipToAlpha_Planar8(src, alpha, dest, flags)
}


var _vImageClipToAlpha_PlanarF func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_PlanarF clamps the values of a 32-bit planar buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_PlanarF(_:_:_:_:)
func VImageClipToAlpha_PlanarF(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_PlanarF == nil {
		panic("Accelerate: symbol vImageClipToAlpha_PlanarF not loaded")
	}
	return _vImageClipToAlpha_PlanarF(src, alpha, dest, flags)
}


var _vImageClipToAlpha_RGBA8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_RGBA8888 clamps the values of an 8-bit-per-channel, 4-channel RGBA buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_RGBA8888(_:_:_:)
func VImageClipToAlpha_RGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_RGBA8888 == nil {
		panic("Accelerate: symbol vImageClipToAlpha_RGBA8888 not loaded")
	}
	return _vImageClipToAlpha_RGBA8888(src, dest, flags)
}


var _vImageClipToAlpha_RGBAFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageClipToAlpha_RGBAFFFF clamps the values of a 32-bit-per-channel, 4-channel RGBA buffer to the corresponding alpha values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClipToAlpha_RGBAFFFF(_:_:_:)
func VImageClipToAlpha_RGBAFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageClipToAlpha_RGBAFFFF == nil {
		panic("Accelerate: symbol vImageClipToAlpha_RGBAFFFF not loaded")
	}
	return _vImageClipToAlpha_RGBAFFFF(src, dest, flags)
}


var _vImageClip_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int

// VImageClip_PlanarF clips the values in a floating-point 32-bit planar buffer to the specified minimum and maximum values.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageClip_PlanarF(_:_:_:_:_:)
func VImageClip_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int {
	if _vImageClip_PlanarF == nil {
		panic("Accelerate: symbol vImageClip_PlanarF not loaded")
	}
	return _vImageClip_PlanarF(src, dest, maxFloat, minFloat, flags)
}


var _vImageContrastStretch_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageContrastStretch_ARGB8888 performs contrast stretching on an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_ARGB8888(_:_:_:)
func VImageContrastStretch_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageContrastStretch_ARGB8888 == nil {
		panic("Accelerate: symbol vImageContrastStretch_ARGB8888 not loaded")
	}
	return _vImageContrastStretch_ARGB8888(src, dest, flags)
}


var _vImageContrastStretch_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageContrastStretch_ARGBFFFF performs contrast stretching on a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_ARGBFFFF(_:_:_:_:_:_:_:)
func VImageContrastStretch_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageContrastStretch_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageContrastStretch_ARGBFFFF not loaded")
	}
	return _vImageContrastStretch_ARGBFFFF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
}


var _vImageContrastStretch_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageContrastStretch_Planar8 performs contrast stretching on an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_Planar8(_:_:_:)
func VImageContrastStretch_Planar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageContrastStretch_Planar8 == nil {
		panic("Accelerate: symbol vImageContrastStretch_Planar8 not loaded")
	}
	return _vImageContrastStretch_Planar8(src, dest, flags)
}


var _vImageContrastStretch_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageContrastStretch_PlanarF performs contrast stretching on a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_PlanarF(_:_:_:_:_:_:_:)
func VImageContrastStretch_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageContrastStretch_PlanarF == nil {
		panic("Accelerate: symbol vImageContrastStretch_PlanarF not loaded")
	}
	return _vImageContrastStretch_PlanarF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
}


var _vImageConvert_12UTo16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_12UTo16U converts an unsigned 12-bit planar buffer to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_12UTo16U(_:_:_:)
func VImageConvert_12UTo16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_12UTo16U == nil {
		panic("Accelerate: symbol vImageConvert_12UTo16U not loaded")
	}
	return _vImageConvert_12UTo16U(src, dest, flags)
}


var _vImageConvert_16Fto16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Fto16Q12 converts a floating-point 16-bit planar buffer to a fixed-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Fto16Q12(_:_:_:)
func VImageConvert_16Fto16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Fto16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_16Fto16Q12 not loaded")
	}
	return _vImageConvert_16Fto16Q12(src, dest, flags)
}


var _vImageConvert_16Fto16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Fto16U converts a floating-point 16-bit planar buffer to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Fto16U(_:_:_:)
func VImageConvert_16Fto16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Fto16U == nil {
		panic("Accelerate: symbol vImageConvert_16Fto16U not loaded")
	}
	return _vImageConvert_16Fto16U(src, dest, flags)
}


var _vImageConvert_16Q12to16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Q12to16F converts a fixed-point 16-bit planar buffer to a floating-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Q12to16F(_:_:_:)
func VImageConvert_16Q12to16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Q12to16F == nil {
		panic("Accelerate: symbol vImageConvert_16Q12to16F not loaded")
	}
	return _vImageConvert_16Q12to16F(src, dest, flags)
}


var _vImageConvert_16Q12to16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Q12to16U converts a fixed-point 16-bit planar buffer to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Q12to16U(_:_:_:)
func VImageConvert_16Q12to16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Q12to16U == nil {
		panic("Accelerate: symbol vImageConvert_16Q12to16U not loaded")
	}
	return _vImageConvert_16Q12to16U(src, dest, flags)
}


var _vImageConvert_16Q12to8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Q12to8 converts a fixed-point 16-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Q12to8(_:_:_:)
func VImageConvert_16Q12to8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Q12to8 == nil {
		panic("Accelerate: symbol vImageConvert_16Q12to8 not loaded")
	}
	return _vImageConvert_16Q12to8(src, dest, flags)
}


var _vImageConvert_16Q12toF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Q12toF converts a fixed-point 16-bit planar buffer to a floating-point 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Q12toF(_:_:_:)
func VImageConvert_16Q12toF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Q12toF == nil {
		panic("Accelerate: symbol vImageConvert_16Q12toF not loaded")
	}
	return _vImageConvert_16Q12toF(src, dest, flags)
}


var _vImageConvert_16SToF func(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int

// VImageConvert_16SToF converts a signed 16-bit planar buffer to a floating-point 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16SToF(_:_:_:_:_:)
func VImageConvert_16SToF(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int {
	if _vImageConvert_16SToF == nil {
		panic("Accelerate: symbol vImageConvert_16SToF not loaded")
	}
	return _vImageConvert_16SToF(src, dest, offset, scale, flags)
}


var _vImageConvert_16UTo12U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16UTo12U converts an unsigned 16-bit planar buffer to an unsigned 12-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16UTo12U(_:_:_:)
func VImageConvert_16UTo12U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16UTo12U == nil {
		panic("Accelerate: symbol vImageConvert_16UTo12U not loaded")
	}
	return _vImageConvert_16UTo12U(src, dest, flags)
}


var _vImageConvert_16UToF func(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int

// VImageConvert_16UToF converts an unsigned 16-bit planar buffer to a floating-point 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16UToF(_:_:_:_:_:)
func VImageConvert_16UToF(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int {
	if _vImageConvert_16UToF == nil {
		panic("Accelerate: symbol vImageConvert_16UToF not loaded")
	}
	return _vImageConvert_16UToF(src, dest, offset, scale, flags)
}


var _vImageConvert_16UToPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16UToPlanar8 converts an unsigned 16-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16UToPlanar8(_:_:_:)
func VImageConvert_16UToPlanar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16UToPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_16UToPlanar8 not loaded")
	}
	return _vImageConvert_16UToPlanar8(src, dest, flags)
}


var _vImageConvert_16Uto16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Uto16F converts an unsigned 16-bit planar buffer to a floating-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Uto16F(_:_:_:)
func VImageConvert_16Uto16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Uto16F == nil {
		panic("Accelerate: symbol vImageConvert_16Uto16F not loaded")
	}
	return _vImageConvert_16Uto16F(src, dest, flags)
}


var _vImageConvert_16Uto16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_16Uto16Q12 converts an unsigned 16-bit planar buffer to a fixed-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_16Uto16Q12(_:_:_:)
func VImageConvert_16Uto16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_16Uto16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_16Uto16Q12 not loaded")
	}
	return _vImageConvert_16Uto16Q12(src, dest, flags)
}


var _vImageConvert_420Yp8_Cb8_Cr8ToARGB8888 func(srcYp unsafe.Pointer, srcCb unsafe.Pointer, srcCr unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_420Yp8_Cb8_Cr8ToARGB8888 converts planar Yp, Cb, and Cr buffers to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_420Yp8_Cb8_Cr8ToARGB8888(_:_:_:_:_:_:_:_:)
func VImageConvert_420Yp8_Cb8_Cr8ToARGB8888(srcYp unsafe.Pointer, srcCb unsafe.Pointer, srcCr unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_420Yp8_Cb8_Cr8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_420Yp8_Cb8_Cr8ToARGB8888 not loaded")
	}
	return _vImageConvert_420Yp8_Cb8_Cr8ToARGB8888(srcYp, srcCb, srcCr, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_420Yp8_CbCr8ToARGB8888 func(srcYp unsafe.Pointer, srcCbCr unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_420Yp8_CbCr8ToARGB8888 converts a planar Yp buffer and a 2-channel CbCr buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_420Yp8_CbCr8ToARGB8888(_:_:_:_:_:_:_:)
func VImageConvert_420Yp8_CbCr8ToARGB8888(srcYp unsafe.Pointer, srcCbCr unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_420Yp8_CbCr8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_420Yp8_CbCr8ToARGB8888 not loaded")
	}
	return _vImageConvert_420Yp8_CbCr8ToARGB8888(srcYp, srcCbCr, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422CbYpCrYp16ToARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint16, flags uint32) int

// VImageConvert_422CbYpCrYp16ToARGB16U converts a 16-bit-per-channel 4:2:2 CbYpCrYp buffer to an unsigned 16-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CbYpCrYp16ToARGB16U(_:_:_:_:_:_:)
func VImageConvert_422CbYpCrYp16ToARGB16U(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint16, flags uint32) int {
	if _vImageConvert_422CbYpCrYp16ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_422CbYpCrYp16ToARGB16U not loaded")
	}
	return _vImageConvert_422CbYpCrYp16ToARGB16U(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422CbYpCrYp16ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_422CbYpCrYp16ToARGB8888 converts a 16-bit-per-channel 4:2:2 CbYpCrYp buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CbYpCrYp16ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_422CbYpCrYp16ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_422CbYpCrYp16ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_422CbYpCrYp16ToARGB8888 not loaded")
	}
	return _vImageConvert_422CbYpCrYp16ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422CbYpCrYp8ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_422CbYpCrYp8ToARGB8888 converts an 8-bit-per-channel 4:2:2 CbYpCrYp buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CbYpCrYp8ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_422CbYpCrYp8ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_422CbYpCrYp8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_422CbYpCrYp8ToARGB8888 not loaded")
	}
	return _vImageConvert_422CbYpCrYp8ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422CbYpCrYp8_AA8ToARGB8888 func(src unsafe.Pointer, srcA unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int

// VImageConvert_422CbYpCrYp8_AA8ToARGB8888 converts an 8-bit-per-channel 4:2:2 CbYpCrYp buffer and an 8-bit alpha buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CbYpCrYp8_AA8ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_422CbYpCrYp8_AA8ToARGB8888(src unsafe.Pointer, srcA unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int {
	if _vImageConvert_422CbYpCrYp8_AA8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_422CbYpCrYp8_AA8ToARGB8888 not loaded")
	}
	return _vImageConvert_422CbYpCrYp8_AA8ToARGB8888(src, srcA, dest, info, permuteMap, flags)
}


var _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha Pixel_16Q12, flags uint32) int

// VImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12 converts a 10-bit-per-channel 4:2:2 CrYpCbYpCbYpCbYpCrYpCrYp buffer to a fixed-point 16-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12(_:_:_:_:_:_:)
func VImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha Pixel_16Q12, flags uint32) int {
	if _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12 not loaded")
	}
	return _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888 converts a 10-bit-per-channel 4:2:2 CrYpCbYpCbYpCbYpCrYpCrYp buffer to a 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888 not loaded")
	}
	return _vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_422YpCbYpCr8ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_422YpCbYpCr8ToARGB8888 converts an 8-bit-per-channel 4:2:2 YpCbYpCr buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_422YpCbYpCr8ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_422YpCbYpCr8ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_422YpCbYpCr8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_422YpCbYpCr8ToARGB8888 not loaded")
	}
	return _vImageConvert_422YpCbYpCr8ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_444AYpCbCr16ToARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int

// VImageConvert_444AYpCbCr16ToARGB16U converts an 10-bit-per-channel 4:4:4 CrYpCb buffer to an unsigned 16-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444AYpCbCr16ToARGB16U(_:_:_:_:_:)
func VImageConvert_444AYpCbCr16ToARGB16U(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int {
	if _vImageConvert_444AYpCbCr16ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_444AYpCbCr16ToARGB16U not loaded")
	}
	return _vImageConvert_444AYpCbCr16ToARGB16U(src, dest, info, permuteMap, flags)
}


var _vImageConvert_444AYpCbCr16ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int

// VImageConvert_444AYpCbCr16ToARGB8888 converts an 16-bit-per-channel 4:4:4 YpCbCr buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444AYpCbCr16ToARGB8888(_:_:_:_:_:)
func VImageConvert_444AYpCbCr16ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int {
	if _vImageConvert_444AYpCbCr16ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_444AYpCbCr16ToARGB8888 not loaded")
	}
	return _vImageConvert_444AYpCbCr16ToARGB8888(src, dest, info, permuteMap, flags)
}


var _vImageConvert_444AYpCbCr8ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int

// VImageConvert_444AYpCbCr8ToARGB8888 converts an 8-bit-per-channel 4:4:4 YpCbCr buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444AYpCbCr8ToARGB8888(_:_:_:_:_:)
func VImageConvert_444AYpCbCr8ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int {
	if _vImageConvert_444AYpCbCr8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_444AYpCbCr8ToARGB8888 not loaded")
	}
	return _vImageConvert_444AYpCbCr8ToARGB8888(src, dest, info, permuteMap, flags)
}


var _vImageConvert_444CbYpCrA8ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int

// VImageConvert_444CbYpCrA8ToARGB8888 converts an 8-bit-per-channel 4:4:4 CbYpCrA buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444CbYpCrA8ToARGB8888(_:_:_:_:_:)
func VImageConvert_444CbYpCrA8ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, flags uint32) int {
	if _vImageConvert_444CbYpCrA8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_444CbYpCrA8ToARGB8888 not loaded")
	}
	return _vImageConvert_444CbYpCrA8ToARGB8888(src, dest, info, permuteMap, flags)
}


var _vImageConvert_444CrYpCb10ToARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha Pixel_16Q12, flags uint32) int

// VImageConvert_444CrYpCb10ToARGB16Q12 converts an 10-bit-per-channel 4:4:4 CrYpCb buffer to a fixed-point 16-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444CrYpCb10ToARGB16Q12(_:_:_:_:_:_:)
func VImageConvert_444CrYpCb10ToARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha Pixel_16Q12, flags uint32) int {
	if _vImageConvert_444CrYpCb10ToARGB16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_444CrYpCb10ToARGB16Q12 not loaded")
	}
	return _vImageConvert_444CrYpCb10ToARGB16Q12(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_444CrYpCb10ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_444CrYpCb10ToARGB8888 converts an 10-bit-per-channel 4:4:4 CrYpCb buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444CrYpCb10ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_444CrYpCb10ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_444CrYpCb10ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_444CrYpCb10ToARGB8888 not loaded")
	}
	return _vImageConvert_444CrYpCb10ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_444CrYpCb8ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int

// VImageConvert_444CrYpCb8ToARGB8888 converts an 8-bit-per-channel 4:4:4 CrYpCb buffer to an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_444CrYpCb8ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_444CrYpCb8ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_YpCbCrToARGB, permuteMap uint8, alpha uint8, flags uint32) int {
	if _vImageConvert_444CrYpCb8ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_444CrYpCb8ToARGB8888 not loaded")
	}
	return _vImageConvert_444CrYpCb8ToARGB8888(src, dest, info, permuteMap, alpha, flags)
}


var _vImageConvert_8to16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_8to16Q12 converts an 8-bit planar buffer to a fixed-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_8to16Q12(_:_:_:)
func VImageConvert_8to16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_8to16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_8to16Q12 not loaded")
	}
	return _vImageConvert_8to16Q12(src, dest, flags)
}


var _vImageConvert_ARGB1555toARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB1555toARGB8888 converts an ARGB1555 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB1555toARGB8888(_:_:_:)
func VImageConvert_ARGB1555toARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB1555toARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB1555toARGB8888 not loaded")
	}
	return _vImageConvert_ARGB1555toARGB8888(src, dest, flags)
}


var _vImageConvert_ARGB1555toPlanar8 func(src unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB1555toPlanar8 deinterleaves an ARGB1555 4-channel interleaved buffer into four 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB1555toPlanar8(_:_:_:_:_:_:)
func VImageConvert_ARGB1555toPlanar8(src unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB1555toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB1555toPlanar8 not loaded")
	}
	return _vImageConvert_ARGB1555toPlanar8(src, destA, destR, destG, destB, flags)
}


var _vImageConvert_ARGB1555toRGB565 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB1555toRGB565 removes the alpha channel from an ARGB1555 buffer to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB1555toRGB565(_:_:_:)
func VImageConvert_ARGB1555toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB1555toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB1555toRGB565 not loaded")
	}
	return _vImageConvert_ARGB1555toRGB565(src, dest, flags)
}


var _vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10 converts a fixed-point 16-bit-per-channel, 4-channel ARGB buffer to a 10-bit-per-channel 4:2:2 CrYpCbYpCbYpCbYpCrYpCrYp buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10(_:_:_:_:_:)
func VImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10 not loaded")
	}
	return _vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB16Q12To444CrYpCb10 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16Q12To444CrYpCb10 converts a fixed-point 16-bit-per-channel, 4-channel ARGB buffer to an 10-bit-per-channel 4:4:4 CrYpCb buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16Q12To444CrYpCb10(_:_:_:_:_:)
func VImageConvert_ARGB16Q12To444CrYpCb10(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16Q12To444CrYpCb10 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16Q12To444CrYpCb10 not loaded")
	}
	return _vImageConvert_ARGB16Q12To444CrYpCb10(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB16Q12ToARGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16Q12ToARGB2101010 converts a fixed-point 16-bit-per-channel, 4-channel interleaved buffer to an ARGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16Q12ToARGB2101010(_:_:_:_:_:_:_:_:)
func VImageConvert_ARGB16Q12ToARGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16Q12ToARGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16Q12ToARGB2101010 not loaded")
	}
	return _vImageConvert_ARGB16Q12ToARGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, RGB101010Min, RGB101010Max, permuteMap, flags)
}


var _vImageConvert_ARGB16Q12ToRGBA1010102 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16Q12ToRGBA1010102 converts a fixed-point 16-bit-per-channel, 4-channel interleaved buffer to an RGBA1010102 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16Q12ToRGBA1010102(_:_:_:_:_:_:_:_:)
func VImageConvert_ARGB16Q12ToRGBA1010102(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16Q12ToRGBA1010102 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16Q12ToRGBA1010102 not loaded")
	}
	return _vImageConvert_ARGB16Q12ToRGBA1010102(src, dest, RGB101010RangeMin, RGB101010RangeMax, RGB101010Min, RGB101010Max, permuteMap, flags)
}


var _vImageConvert_ARGB16Q12ToXRGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16Q12ToXRGB2101010 converts a fixed-point 16-bit-per-channel, 4-channel interleaved buffer to an XRGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16Q12ToXRGB2101010(_:_:_:_:_:_:_:_:)
func VImageConvert_ARGB16Q12ToXRGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, RGB101010Min int32, RGB101010Max int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16Q12ToXRGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16Q12ToXRGB2101010 not loaded")
	}
	return _vImageConvert_ARGB16Q12ToXRGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, RGB101010Min, RGB101010Max, permuteMap, flags)
}


var _vImageConvert_ARGB16UTo422CbYpCrYp16 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UTo422CbYpCrYp16 converts an unsigned 16-bit-per-channel, 4-channel ARGB buffer to a 16-bit-per-channel 4:2:2 CbYpCrYp buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UTo422CbYpCrYp16(_:_:_:_:_:)
func VImageConvert_ARGB16UTo422CbYpCrYp16(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UTo422CbYpCrYp16 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UTo422CbYpCrYp16 not loaded")
	}
	return _vImageConvert_ARGB16UTo422CbYpCrYp16(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB16UTo444AYpCbCr16 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UTo444AYpCbCr16 converts an unsigned 16-bit-per-channel, 4-channel ARGB buffer to an 16-bit-per-channel 4:4:4 AYpCbCr buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UTo444AYpCbCr16(_:_:_:_:_:)
func VImageConvert_ARGB16UTo444AYpCbCr16(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UTo444AYpCbCr16 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UTo444AYpCbCr16 not loaded")
	}
	return _vImageConvert_ARGB16UTo444AYpCbCr16(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB16UToARGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UToARGB2101010 converts an unsigned 16-bit-per-channel, 4-channel interleaved buffer to an ARGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UToARGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGB16UToARGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UToARGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UToARGB2101010 not loaded")
	}
	return _vImageConvert_ARGB16UToARGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB16UToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int

// VImageConvert_ARGB16UToARGB8888 converts an unsigned 16-bit-per-channel, 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UToARGB8888(_:_:_:_:_:_:)
func VImageConvert_ARGB16UToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvert_ARGB16UToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UToARGB8888 not loaded")
	}
	return _vImageConvert_ARGB16UToARGB8888(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImageConvert_ARGB16UToRGBA1010102 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UToRGBA1010102 converts an unsigned 16-bit-per-channel, 4-channel interleaved buffer to an RGBA1010102 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UToRGBA1010102(_:_:_:_:_:_:)
func VImageConvert_ARGB16UToRGBA1010102(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UToRGBA1010102 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UToRGBA1010102 not loaded")
	}
	return _vImageConvert_ARGB16UToRGBA1010102(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB16UToXRGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UToXRGB2101010 converts an unsigned 16-bit-per-channel, 4-channel interleaved buffer to an XRGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UToXRGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGB16UToXRGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UToXRGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UToXRGB2101010 not loaded")
	}
	return _vImageConvert_ARGB16UToXRGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB16UtoARGB8888_dithered func(src unsafe.Pointer, dest unsafe.Pointer, dither int, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB16UtoARGB8888_dithered converts an unsigned 16-bit-per-channel, 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UtoARGB8888_dithered(_:_:_:_:_:)
func VImageConvert_ARGB16UtoARGB8888_dithered(src unsafe.Pointer, dest unsafe.Pointer, dither int, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB16UtoARGB8888_dithered == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UtoARGB8888_dithered not loaded")
	}
	return _vImageConvert_ARGB16UtoARGB8888_dithered(src, dest, dither, permuteMap, flags)
}


var _vImageConvert_ARGB16UtoPlanar16U func(argbSrc unsafe.Pointer, aDest unsafe.Pointer, rDest unsafe.Pointer, gDest unsafe.Pointer, bDest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB16UtoPlanar16U deinterleaves an unsigned 16-bit-per-channel, 4-channel interleaved buffer into four unsigned 16-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UtoPlanar16U(_:_:_:_:_:_:)
func VImageConvert_ARGB16UtoPlanar16U(argbSrc unsafe.Pointer, aDest unsafe.Pointer, rDest unsafe.Pointer, gDest unsafe.Pointer, bDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB16UtoPlanar16U == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UtoPlanar16U not loaded")
	}
	return _vImageConvert_ARGB16UtoPlanar16U(argbSrc, aDest, rDest, gDest, bDest, flags)
}


var _vImageConvert_ARGB16UtoRGB16U func(argbSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB16UtoRGB16U removes the alpha channel from an unsigned 16-bit-per-channel ARGB buffer to produce an unsigned 16-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB16UtoRGB16U(_:_:_:)
func VImageConvert_ARGB16UtoRGB16U(argbSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB16UtoRGB16U == nil {
		panic("Accelerate: symbol vImageConvert_ARGB16UtoRGB16U not loaded")
	}
	return _vImageConvert_ARGB16UtoRGB16U(argbSrc, rgbDest, flags)
}


var _vImageConvert_ARGB2101010ToARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB2101010ToARGB16F converts an ARGB2101010 32-bit, 4-channel interleaved buffer to a floating-point 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB2101010ToARGB16F(_:_:_:_:_:_:)
func VImageConvert_ARGB2101010ToARGB16F(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB2101010ToARGB16F == nil {
		panic("Accelerate: symbol vImageConvert_ARGB2101010ToARGB16F not loaded")
	}
	return _vImageConvert_ARGB2101010ToARGB16F(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB2101010ToARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB2101010ToARGB16Q12 converts an ARGB2101010 32-bit, 4-channel interleaved buffer to a fixed-point 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB2101010ToARGB16Q12(_:_:_:_:_:_:)
func VImageConvert_ARGB2101010ToARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB2101010ToARGB16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB2101010ToARGB16Q12 not loaded")
	}
	return _vImageConvert_ARGB2101010ToARGB16Q12(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB2101010ToARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB2101010ToARGB16U converts an ARGB2101010 32-bit, 4-channel interleaved buffer to an unsigned 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB2101010ToARGB16U(_:_:_:_:_:_:)
func VImageConvert_ARGB2101010ToARGB16U(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB2101010ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_ARGB2101010ToARGB16U not loaded")
	}
	return _vImageConvert_ARGB2101010ToARGB16U(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB2101010ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB2101010ToARGB8888 converts an ARGB2101010 32-bit, 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB2101010ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_ARGB2101010ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB2101010ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB2101010ToARGB8888 not loaded")
	}
	return _vImageConvert_ARGB2101010ToARGB8888(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB2101010ToARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB2101010ToARGBFFFF converts an ARGB2101010 32-bit, 4-channel interleaved buffer to a floating-point 32-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB2101010ToARGBFFFF(_:_:_:_:_:_:)
func VImageConvert_ARGB2101010ToARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB2101010ToARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_ARGB2101010ToARGBFFFF not loaded")
	}
	return _vImageConvert_ARGB2101010ToARGBFFFF(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB8888To420Yp8_Cb8_Cr8 func(src unsafe.Pointer, destYp unsafe.Pointer, destCb unsafe.Pointer, destCr unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To420Yp8_Cb8_Cr8 converts an 8-bit-per-channel, 4-channel ARGB buffer to planar Yp, Cb, and Cr buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To420Yp8_Cb8_Cr8(_:_:_:_:_:_:_:)
func VImageConvert_ARGB8888To420Yp8_Cb8_Cr8(src unsafe.Pointer, destYp unsafe.Pointer, destCb unsafe.Pointer, destCr unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To420Yp8_Cb8_Cr8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To420Yp8_Cb8_Cr8 not loaded")
	}
	return _vImageConvert_ARGB8888To420Yp8_Cb8_Cr8(src, destYp, destCb, destCr, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To420Yp8_CbCr8 func(src unsafe.Pointer, destYp unsafe.Pointer, destCbCr unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To420Yp8_CbCr8 converts an 8-bit-per-channel, 4-channel ARGB buffer to a planar Yp buffer and a 2-channel CbCr buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To420Yp8_CbCr8(_:_:_:_:_:_:)
func VImageConvert_ARGB8888To420Yp8_CbCr8(src unsafe.Pointer, destYp unsafe.Pointer, destCbCr unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To420Yp8_CbCr8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To420Yp8_CbCr8 not loaded")
	}
	return _vImageConvert_ARGB8888To420Yp8_CbCr8(src, destYp, destCbCr, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To422CbYpCrYp16 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To422CbYpCrYp16 converts an 8-bit-per-channel, 4-channel ARGB buffer to a 16-bit-per-channel 4:2:2 CbYpCrYp buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To422CbYpCrYp16(_:_:_:_:_:)
func VImageConvert_ARGB8888To422CbYpCrYp16(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To422CbYpCrYp16 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To422CbYpCrYp16 not loaded")
	}
	return _vImageConvert_ARGB8888To422CbYpCrYp16(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To422CbYpCrYp8 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To422CbYpCrYp8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:2:2 CbCrYp buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To422CbYpCrYp8(_:_:_:_:_:)
func VImageConvert_ARGB8888To422CbYpCrYp8(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To422CbYpCrYp8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To422CbYpCrYp8 not loaded")
	}
	return _vImageConvert_ARGB8888To422CbYpCrYp8(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To422CbYpCrYp8_AA8 func(src unsafe.Pointer, dest unsafe.Pointer, destA unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To422CbYpCrYp8_AA8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:2:2 CbYpCrYp buffer and an 8-bit alpha buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To422CbYpCrYp8_AA8(_:_:_:_:_:_:)
func VImageConvert_ARGB8888To422CbYpCrYp8_AA8(src unsafe.Pointer, dest unsafe.Pointer, destA unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To422CbYpCrYp8_AA8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To422CbYpCrYp8_AA8 not loaded")
	}
	return _vImageConvert_ARGB8888To422CbYpCrYp8_AA8(src, dest, destA, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10 converts an 8-bit-per-channel, 4-channel ARGB buffer to a 10-bit-per-channel 4:2:2 CrYpCbYpCbYpCbYpCrYpCrYp buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10(_:_:_:_:_:)
func VImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10 not loaded")
	}
	return _vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To422YpCbYpCr8 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To422YpCbYpCr8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:2:2 YpCbYpCr buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To422YpCbYpCr8(_:_:_:_:_:)
func VImageConvert_ARGB8888To422YpCbYpCr8(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To422YpCbYpCr8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To422YpCbYpCr8 not loaded")
	}
	return _vImageConvert_ARGB8888To422YpCbYpCr8(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To444AYpCbCr16 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To444AYpCbCr16 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 16-bit-per-channel 4:4:4 YpCbCr buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To444AYpCbCr16(_:_:_:_:_:)
func VImageConvert_ARGB8888To444AYpCbCr16(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To444AYpCbCr16 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To444AYpCbCr16 not loaded")
	}
	return _vImageConvert_ARGB8888To444AYpCbCr16(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To444AYpCbCr8 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To444AYpCbCr8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:4:4 YpCbCr buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To444AYpCbCr8(_:_:_:_:_:)
func VImageConvert_ARGB8888To444AYpCbCr8(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To444AYpCbCr8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To444AYpCbCr8 not loaded")
	}
	return _vImageConvert_ARGB8888To444AYpCbCr8(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To444CbYpCrA8 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To444CbYpCrA8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:4:4 CrYpCbA buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To444CbYpCrA8(_:_:_:_:_:)
func VImageConvert_ARGB8888To444CbYpCrA8(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To444CbYpCrA8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To444CbYpCrA8 not loaded")
	}
	return _vImageConvert_ARGB8888To444CbYpCrA8(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To444CrYpCb10 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To444CrYpCb10 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 10-bit-per-channel 4:4:4 CrYpCb buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To444CrYpCb10(_:_:_:_:_:)
func VImageConvert_ARGB8888To444CrYpCb10(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To444CrYpCb10 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To444CrYpCb10 not loaded")
	}
	return _vImageConvert_ARGB8888To444CrYpCb10(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888To444CrYpCb8 func(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888To444CrYpCb8 converts an 8-bit-per-channel, 4-channel ARGB buffer to an 8-bit-per-channel 4:4:4 CrYpCb buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To444CrYpCb8(_:_:_:_:_:)
func VImageConvert_ARGB8888To444CrYpCb8(src unsafe.Pointer, dest unsafe.Pointer, info *VImage_ARGBToYpCbCr, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888To444CrYpCb8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888To444CrYpCb8 not loaded")
	}
	return _vImageConvert_ARGB8888To444CrYpCb8(src, dest, info, permuteMap, flags)
}


var _vImageConvert_ARGB8888ToARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_ARGB_16U, flags uint32) int

// VImageConvert_ARGB8888ToARGB16U converts an 8-bit-per-channel, 4-channel interleaved buffer to an unsigned 16-bit-per-channel, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888ToARGB16U(_:_:_:_:_:_:)
func VImageConvert_ARGB8888ToARGB16U(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageConvert_ARGB8888ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888ToARGB16U not loaded")
	}
	return _vImageConvert_ARGB8888ToARGB16U(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImageConvert_ARGB8888ToARGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888ToARGB2101010 converts an 8-bit-per-channel, 4-channel interleaved buffer to an ARGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888ToARGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGB8888ToARGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888ToARGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888ToARGB2101010 not loaded")
	}
	return _vImageConvert_ARGB8888ToARGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB8888ToRGB16U func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_16U, flags uint32) int

// VImageConvert_ARGB8888ToRGB16U removes the alpha channel from an 8-bit-per-channel ARGB buffer to produce an unsigned 16-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888ToRGB16U(_:_:_:_:_:_:)
func VImageConvert_ARGB8888ToRGB16U(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_16U, flags uint32) int {
	if _vImageConvert_ARGB8888ToRGB16U == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888ToRGB16U not loaded")
	}
	return _vImageConvert_ARGB8888ToRGB16U(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImageConvert_ARGB8888ToRGBA1010102 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888ToRGBA1010102 converts an 8-bit-per-channel, 4-channel interleaved buffer to an RGBA1010102 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888ToRGBA1010102(_:_:_:_:_:_:)
func VImageConvert_ARGB8888ToRGBA1010102(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888ToRGBA1010102 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888ToRGBA1010102 not loaded")
	}
	return _vImageConvert_ARGB8888ToRGBA1010102(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB8888ToXRGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGB8888ToXRGB2101010 converts an 8-bit-per-channel, 4-channel interleaved buffer to an ARGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888ToXRGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGB8888ToXRGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGB8888ToXRGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888ToXRGB2101010 not loaded")
	}
	return _vImageConvert_ARGB8888ToXRGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGB8888toARGB1555 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB8888toARGB1555 converts an 8-bit-per-channel, 4-channel interleaved buffer to an ARGB1555 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toARGB1555(_:_:_:)
func VImageConvert_ARGB8888toARGB1555(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB8888toARGB1555 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toARGB1555 not loaded")
	}
	return _vImageConvert_ARGB8888toARGB1555(src, dest, flags)
}


var _vImageConvert_ARGB8888toARGB1555_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_ARGB8888toARGB1555_dithered converts an 8-bit-per-channel, 4-channel interleaved buffer to an ARGB1555 4-channel interleaved buffer usng the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toARGB1555_dithered(_:_:_:_:_:)
func VImageConvert_ARGB8888toARGB1555_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_ARGB8888toARGB1555_dithered == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toARGB1555_dithered not loaded")
	}
	return _vImageConvert_ARGB8888toARGB1555_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_ARGB8888toPlanar16Q12 func(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB8888toPlanar16Q12 deinterleaves an 8-bit-per-channel, 4-channel interleaved buffer into four fixed-point 16-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toPlanar16Q12(_:_:_:_:_:_:)
func VImageConvert_ARGB8888toPlanar16Q12(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB8888toPlanar16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toPlanar16Q12 not loaded")
	}
	return _vImageConvert_ARGB8888toPlanar16Q12(src, alpha, red, green, blue, flags)
}


var _vImageConvert_ARGB8888toPlanar8 func(srcARGB unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB8888toPlanar8 deinterleaves an 8-bit-per-channel, 4-channel interleaved buffer into four 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toPlanar8(_:_:_:_:_:_:)
func VImageConvert_ARGB8888toPlanar8(srcARGB unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB8888toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toPlanar8 not loaded")
	}
	return _vImageConvert_ARGB8888toPlanar8(srcARGB, destA, destR, destG, destB, flags)
}


var _vImageConvert_ARGB8888toPlanarF func(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_ARGB8888toPlanarF deinterleaves an 8-bit-per-channel, 4-channel interleaved buffer into four floating-point 32-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toPlanarF(_:_:_:_:_:_:_:_:)
func VImageConvert_ARGB8888toPlanarF(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_ARGB8888toPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toPlanarF not loaded")
	}
	return _vImageConvert_ARGB8888toPlanarF(src, alpha, red, green, blue, maxFloat, minFloat, flags)
}


var _vImageConvert_ARGB8888toRGB565 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGB8888toRGB565 removes the alpha channel from an 8-bit-per-channel ARGB buffer to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toRGB565(_:_:_:)
func VImageConvert_ARGB8888toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGB8888toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toRGB565 not loaded")
	}
	return _vImageConvert_ARGB8888toRGB565(src, dest, flags)
}


var _vImageConvert_ARGB8888toRGB565_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_ARGB8888toRGB565_dithered removes the alpha channel from an 8-bit-per-channel ARGB buffer using the specified dithering algorithm to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toRGB565_dithered(_:_:_:_:_:)
func VImageConvert_ARGB8888toRGB565_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_ARGB8888toRGB565_dithered == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toRGB565_dithered not loaded")
	}
	return _vImageConvert_ARGB8888toRGB565_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_ARGB8888toRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int

// VImageConvert_ARGB8888toRGB888 removes the alpha channel from an 8-bit-per-channel ARGB buffer to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888toRGB888(_:_:_:)
func VImageConvert_ARGB8888toRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int {
	if _vImageConvert_ARGB8888toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_ARGB8888toRGB888 not loaded")
	}
	return _vImageConvert_ARGB8888toRGB888(arg0, arg1, arg2)
}


var _vImageConvert_ARGBFFFFToARGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGBFFFFToARGB2101010 converts a floating-point 32-bit-per-channel, 4-channel interleaved buffer to an ARGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFToARGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGBFFFFToARGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGBFFFFToARGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFToARGB2101010 not loaded")
	}
	return _vImageConvert_ARGBFFFFToARGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGBFFFFToXRGB2101010 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_ARGBFFFFToXRGB2101010 converts a floating-point 32-bit-per-channel, 4-channel interleaved buffer to an XRGB2101010 32-bit, 4-channel buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFToXRGB2101010(_:_:_:_:_:_:)
func VImageConvert_ARGBFFFFToXRGB2101010(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGBFFFFToXRGB2101010 == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFToXRGB2101010 not loaded")
	}
	return _vImageConvert_ARGBFFFFToXRGB2101010(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_ARGBFFFFtoARGB8888_dithered func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, dither int, permuteMap uint8, flags uint32) int

// VImageConvert_ARGBFFFFtoARGB8888_dithered converts a floating-point 32-bit-per-channel, 4-channel buffer to an 8-bit-per-channel, 4-channel buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFtoARGB8888_dithered(_:_:_:_:_:_:_:)
func VImageConvert_ARGBFFFFtoARGB8888_dithered(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, dither int, permuteMap uint8, flags uint32) int {
	if _vImageConvert_ARGBFFFFtoARGB8888_dithered == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFtoARGB8888_dithered not loaded")
	}
	return _vImageConvert_ARGBFFFFtoARGB8888_dithered(src, dest, maxFloat, minFloat, dither, permuteMap, flags)
}


var _vImageConvert_ARGBFFFFtoPlanar8 func(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_ARGBFFFFtoPlanar8 deinterleaves a floating-point 32-bit-per-channel, 4-channel interleaved buffer into four 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFtoPlanar8(_:_:_:_:_:_:_:_:)
func VImageConvert_ARGBFFFFtoPlanar8(src unsafe.Pointer, alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_ARGBFFFFtoPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFtoPlanar8 not loaded")
	}
	return _vImageConvert_ARGBFFFFtoPlanar8(src, alpha, red, green, blue, maxFloat, minFloat, flags)
}


var _vImageConvert_ARGBFFFFtoPlanarF func(srcARGB unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int

// VImageConvert_ARGBFFFFtoPlanarF deinterleaves a floating-point 32-bit-per-channel, 4-channel interleaved buffer into four floating-point 38-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFtoPlanarF(_:_:_:_:_:_:)
func VImageConvert_ARGBFFFFtoPlanarF(srcARGB unsafe.Pointer, destA unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGBFFFFtoPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFtoPlanarF not loaded")
	}
	return _vImageConvert_ARGBFFFFtoPlanarF(srcARGB, destA, destR, destG, destB, flags)
}


var _vImageConvert_ARGBFFFFtoRGBFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_ARGBFFFFtoRGBFFF removes the alpha channel from a floating-point 32-bit-per-channel ARGB buffer to produce a floating-point 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBFFFFtoRGBFFF(_:_:_:)
func VImageConvert_ARGBFFFFtoRGBFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGBFFFFtoRGBFFF == nil {
		panic("Accelerate: symbol vImageConvert_ARGBFFFFtoRGBFFF not loaded")
	}
	return _vImageConvert_ARGBFFFFtoRGBFFF(src, dest, flags)
}


var _vImageConvert_ARGBToYpCbCr_GenerateConversion func(matrix *VImage_ARGBToYpCbCrMatrix, pixelRange *VImage_YpCbCrPixelRange, outInfo *VImage_ARGBToYpCbCr, inARGBType unsafe.Pointer, outYpCbCrType unsafe.Pointer, flags uint32) int

// VImageConvert_ARGBToYpCbCr_GenerateConversion generates the information that describes the conversion from ARGB to YpCbCr.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBToYpCbCr_GenerateConversion(_:_:_:_:_:_:)
func VImageConvert_ARGBToYpCbCr_GenerateConversion(matrix *VImage_ARGBToYpCbCrMatrix, pixelRange *VImage_YpCbCrPixelRange, outInfo *VImage_ARGBToYpCbCr, inARGBType unsafe.Pointer, outYpCbCrType unsafe.Pointer, flags uint32) int {
	if _vImageConvert_ARGBToYpCbCr_GenerateConversion == nil {
		panic("Accelerate: symbol vImageConvert_ARGBToYpCbCr_GenerateConversion not loaded")
	}
	return _vImageConvert_ARGBToYpCbCr_GenerateConversion(matrix, pixelRange, outInfo, inARGBType, outYpCbCrType, flags)
}


var _vImageConvert_AnyToAny func(converter unsafe.Pointer, srcs unsafe.Pointer, dests unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageConvert_AnyToAny converts the pixels in a vImage buffer to another format, using the specified converter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_AnyToAny(_:_:_:_:_:)
func VImageConvert_AnyToAny(converter unsafe.Pointer, srcs unsafe.Pointer, dests unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageConvert_AnyToAny == nil {
		panic("Accelerate: symbol vImageConvert_AnyToAny not loaded")
	}
	return _vImageConvert_AnyToAny(converter, srcs, dests, tempBuffer, flags)
}


var _vImageConvert_BGRA16UtoRGB16U func(bgraSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_BGRA16UtoRGB16U removes the alpha channel from an unsigned 16-bit-per-channel BGRA buffer to produce an unsigned 16-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRA16UtoRGB16U(_:_:_:)
func VImageConvert_BGRA16UtoRGB16U(bgraSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_BGRA16UtoRGB16U == nil {
		panic("Accelerate: symbol vImageConvert_BGRA16UtoRGB16U not loaded")
	}
	return _vImageConvert_BGRA16UtoRGB16U(bgraSrc, rgbDest, flags)
}


var _vImageConvert_BGRA8888toRGB565 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_BGRA8888toRGB565 removes the alpha channel from an 8-bit-per-channel RGBA buffer to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRA8888toRGB565(_:_:_:)
func VImageConvert_BGRA8888toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_BGRA8888toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_BGRA8888toRGB565 not loaded")
	}
	return _vImageConvert_BGRA8888toRGB565(src, dest, flags)
}


var _vImageConvert_BGRA8888toRGB565_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_BGRA8888toRGB565_dithered removes the alpha channel from an 8-bit-per-channel BGRA buffer using the specified dithering algorithm to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRA8888toRGB565_dithered(_:_:_:_:_:)
func VImageConvert_BGRA8888toRGB565_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_BGRA8888toRGB565_dithered == nil {
		panic("Accelerate: symbol vImageConvert_BGRA8888toRGB565_dithered not loaded")
	}
	return _vImageConvert_BGRA8888toRGB565_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_BGRA8888toRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int

// VImageConvert_BGRA8888toRGB888 removes the alpha channel from an 8-bit-per-channel BGRA buffer to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRA8888toRGB888(_:_:_:)
func VImageConvert_BGRA8888toRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int {
	if _vImageConvert_BGRA8888toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_BGRA8888toRGB888 not loaded")
	}
	return _vImageConvert_BGRA8888toRGB888(arg0, arg1, arg2)
}


var _vImageConvert_BGRAFFFFtoRGBFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_BGRAFFFFtoRGBFFF removes the alpha channel from a floating-point 32-bit-per-channel BGRA buffer to produce a floating-point 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRAFFFFtoRGBFFF(_:_:_:)
func VImageConvert_BGRAFFFFtoRGBFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_BGRAFFFFtoRGBFFF == nil {
		panic("Accelerate: symbol vImageConvert_BGRAFFFFtoRGBFFF not loaded")
	}
	return _vImageConvert_BGRAFFFFtoRGBFFF(src, dest, flags)
}


var _vImageConvert_BGRX8888ToPlanar8 func(src unsafe.Pointer, blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, flags uint32) int

// VImageConvert_BGRX8888ToPlanar8 deinterleaves an 8-bit-per-channel, 4-channel interleaved buffer into three 8-bit planar buffers and discards the last channel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRX8888ToPlanar8(_:_:_:_:_:)
func VImageConvert_BGRX8888ToPlanar8(src unsafe.Pointer, blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, flags uint32) int {
	if _vImageConvert_BGRX8888ToPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_BGRX8888ToPlanar8 not loaded")
	}
	return _vImageConvert_BGRX8888ToPlanar8(src, blue, green, red, flags)
}


var _vImageConvert_BGRXFFFFToPlanarF func(src unsafe.Pointer, blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, flags uint32) int

// VImageConvert_BGRXFFFFToPlanarF deinterleaves a floating-point 32-bit-per-channel, 4-channel interleaved buffer into three floating-point 32-bit planar buffers and discards the last channel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_BGRXFFFFToPlanarF(_:_:_:_:_:)
func VImageConvert_BGRXFFFFToPlanarF(src unsafe.Pointer, blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, flags uint32) int {
	if _vImageConvert_BGRXFFFFToPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_BGRXFFFFToPlanarF not loaded")
	}
	return _vImageConvert_BGRXFFFFToPlanarF(src, blue, green, red, flags)
}


var _vImageConvert_ChunkyToPlanar8 func(srcChannels unsafe.Pointer, destPlanarBuffers unsafe.Pointer, channelCount uint, srcStrideBytes uintptr, srcWidth uint, srcHeight uint, srcRowBytes uintptr, flags uint32) int

// VImageConvert_ChunkyToPlanar8 deinterleaves an 8-bit-per-channel interleaved buffer with an arbitrary number of channels into the corresponding number of 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ChunkyToPlanar8(_:_:_:_:_:_:_:_:)
func VImageConvert_ChunkyToPlanar8(srcChannels unsafe.Pointer, destPlanarBuffers unsafe.Pointer, channelCount uint, srcStrideBytes uintptr, srcWidth uint, srcHeight uint, srcRowBytes uintptr, flags uint32) int {
	if _vImageConvert_ChunkyToPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_ChunkyToPlanar8 not loaded")
	}
	return _vImageConvert_ChunkyToPlanar8(srcChannels, destPlanarBuffers, channelCount, srcStrideBytes, srcWidth, srcHeight, srcRowBytes, flags)
}


var _vImageConvert_ChunkyToPlanarF func(srcChannels unsafe.Pointer, destPlanarBuffers unsafe.Pointer, channelCount uint, srcStrideBytes uintptr, srcWidth uint, srcHeight uint, srcRowBytes uintptr, flags uint32) int

// VImageConvert_ChunkyToPlanarF deinterleaves a floating-point 32-bit-per-channel interleaved buffer with an arbitrary number of channels into the corresponding number of 32-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_ChunkyToPlanarF(_:_:_:_:_:_:_:_:)
func VImageConvert_ChunkyToPlanarF(srcChannels unsafe.Pointer, destPlanarBuffers unsafe.Pointer, channelCount uint, srcStrideBytes uintptr, srcWidth uint, srcHeight uint, srcRowBytes uintptr, flags uint32) int {
	if _vImageConvert_ChunkyToPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_ChunkyToPlanarF not loaded")
	}
	return _vImageConvert_ChunkyToPlanarF(srcChannels, destPlanarBuffers, channelCount, srcStrideBytes, srcWidth, srcHeight, srcRowBytes, flags)
}


var _vImageConvert_FTo16S func(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int

// VImageConvert_FTo16S converts a floating-point 32-bit planar buffer to a signed 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_FTo16S(_:_:_:_:_:)
func VImageConvert_FTo16S(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int {
	if _vImageConvert_FTo16S == nil {
		panic("Accelerate: symbol vImageConvert_FTo16S not loaded")
	}
	return _vImageConvert_FTo16S(src, dest, offset, scale, flags)
}


var _vImageConvert_FTo16U func(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int

// VImageConvert_FTo16U converts a floating-point 32-bit planar buffer to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_FTo16U(_:_:_:_:_:)
func VImageConvert_FTo16U(src unsafe.Pointer, dest unsafe.Pointer, offset float32, scale float32, flags uint32) int {
	if _vImageConvert_FTo16U == nil {
		panic("Accelerate: symbol vImageConvert_FTo16U not loaded")
	}
	return _vImageConvert_FTo16U(src, dest, offset, scale, flags)
}


var _vImageConvert_Fto16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Fto16Q12 converts a floating-point 32-bit planar buffer to a fixed-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Fto16Q12(_:_:_:)
func VImageConvert_Fto16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Fto16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_Fto16Q12 not loaded")
	}
	return _vImageConvert_Fto16Q12(src, dest, flags)
}


var _vImageConvert_Indexed1toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int

// VImageConvert_Indexed1toPlanar8 converts an indexed 1-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Indexed1toPlanar8(_:_:_:_:)
func VImageConvert_Indexed1toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int {
	if _vImageConvert_Indexed1toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Indexed1toPlanar8 not loaded")
	}
	return _vImageConvert_Indexed1toPlanar8(src, dest, colors, flags)
}


var _vImageConvert_Indexed2toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int

// VImageConvert_Indexed2toPlanar8 converts an indexed 2-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Indexed2toPlanar8(_:_:_:_:)
func VImageConvert_Indexed2toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int {
	if _vImageConvert_Indexed2toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Indexed2toPlanar8 not loaded")
	}
	return _vImageConvert_Indexed2toPlanar8(src, dest, colors, flags)
}


var _vImageConvert_Indexed4toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int

// VImageConvert_Indexed4toPlanar8 converts an indexed 4-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Indexed4toPlanar8(_:_:_:_:)
func VImageConvert_Indexed4toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, colors Pixel_8, flags uint32) int {
	if _vImageConvert_Indexed4toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Indexed4toPlanar8 not loaded")
	}
	return _vImageConvert_Indexed4toPlanar8(src, dest, colors, flags)
}


var _vImageConvert_Planar16FtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16FtoPlanar8 converts a floaing-point 16-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16FtoPlanar8(_:_:_:)
func VImageConvert_Planar16FtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16FtoPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Planar16FtoPlanar8 not loaded")
	}
	return _vImageConvert_Planar16FtoPlanar8(src, dest, flags)
}


var _vImageConvert_Planar16FtoPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16FtoPlanarF converts a floating-point 16-bit planar buffer to a floating-point 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16FtoPlanarF(_:_:_:)
func VImageConvert_Planar16FtoPlanarF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16FtoPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_Planar16FtoPlanarF not loaded")
	}
	return _vImageConvert_Planar16FtoPlanarF(src, dest, flags)
}


var _vImageConvert_Planar16Q12toARGB16F func(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16Q12toARGB16F interleaves four fixed-point 16-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16Q12toARGB16F(_:_:_:_:_:_:)
func VImageConvert_Planar16Q12toARGB16F(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16Q12toARGB16F == nil {
		panic("Accelerate: symbol vImageConvert_Planar16Q12toARGB16F not loaded")
	}
	return _vImageConvert_Planar16Q12toARGB16F(alpha, red, green, blue, dest, flags)
}


var _vImageConvert_Planar16Q12toARGB8888 func(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16Q12toARGB8888 interleaves four fixed-point 16-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16Q12toARGB8888(_:_:_:_:_:_:)
func VImageConvert_Planar16Q12toARGB8888(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16Q12toARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar16Q12toARGB8888 not loaded")
	}
	return _vImageConvert_Planar16Q12toARGB8888(alpha, red, green, blue, dest, flags)
}


var _vImageConvert_Planar16Q12toRGB16F func(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16Q12toRGB16F interleaves three fixed-point 16-bit planar buffers into a floating-point 32-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16Q12toRGB16F(_:_:_:_:_:)
func VImageConvert_Planar16Q12toRGB16F(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16Q12toRGB16F == nil {
		panic("Accelerate: symbol vImageConvert_Planar16Q12toRGB16F not loaded")
	}
	return _vImageConvert_Planar16Q12toRGB16F(red, green, blue, dest, flags)
}


var _vImageConvert_Planar16Q12toRGB888 func(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16Q12toRGB888 interleaves three fixed-point 16-bit planar buffers into an 8-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16Q12toRGB888(_:_:_:_:_:)
func VImageConvert_Planar16Q12toRGB888(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16Q12toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar16Q12toRGB888 not loaded")
	}
	return _vImageConvert_Planar16Q12toRGB888(red, green, blue, dest, flags)
}


var _vImageConvert_Planar16UtoARGB16U func(aSrc unsafe.Pointer, rSrc unsafe.Pointer, gSrc unsafe.Pointer, bSrc unsafe.Pointer, argbDest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16UtoARGB16U interleaves four unsigned 16-bit planar buffers into an unsigned 16-bit-per-channel, 4-channel interleaved ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16UtoARGB16U(_:_:_:_:_:_:)
func VImageConvert_Planar16UtoARGB16U(aSrc unsafe.Pointer, rSrc unsafe.Pointer, gSrc unsafe.Pointer, bSrc unsafe.Pointer, argbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16UtoARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_Planar16UtoARGB16U not loaded")
	}
	return _vImageConvert_Planar16UtoARGB16U(aSrc, rSrc, gSrc, bSrc, argbDest, flags)
}


var _vImageConvert_Planar16UtoPlanar8_dithered func(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_Planar16UtoPlanar8_dithered converts an unsigned 16-bit planar buffer to an 8-bit planar buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16UtoPlanar8_dithered(_:_:_:_:)
func VImageConvert_Planar16UtoPlanar8_dithered(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_Planar16UtoPlanar8_dithered == nil {
		panic("Accelerate: symbol vImageConvert_Planar16UtoPlanar8_dithered not loaded")
	}
	return _vImageConvert_Planar16UtoPlanar8_dithered(src, dest, dither, flags)
}


var _vImageConvert_Planar16UtoRGB16U func(rSrc unsafe.Pointer, gSrc unsafe.Pointer, bSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar16UtoRGB16U interleaves three unsigned 16-bit planar buffers into an unsigned 16-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar16UtoRGB16U(_:_:_:_:_:)
func VImageConvert_Planar16UtoRGB16U(rSrc unsafe.Pointer, gSrc unsafe.Pointer, bSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar16UtoRGB16U == nil {
		panic("Accelerate: symbol vImageConvert_Planar16UtoRGB16U not loaded")
	}
	return _vImageConvert_Planar16UtoRGB16U(rSrc, gSrc, bSrc, rgbDest, flags)
}


var _vImageConvert_Planar1toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar1toPlanar8 converts a 1-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar1toPlanar8(_:_:_:)
func VImageConvert_Planar1toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar1toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Planar1toPlanar8 not loaded")
	}
	return _vImageConvert_Planar1toPlanar8(src, dest, flags)
}


var _vImageConvert_Planar2toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar2toPlanar8 converts a 2-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar2toPlanar8(_:_:_:)
func VImageConvert_Planar2toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar2toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Planar2toPlanar8 not loaded")
	}
	return _vImageConvert_Planar2toPlanar8(src, dest, flags)
}


var _vImageConvert_Planar4toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar4toPlanar8 converts a 4-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar4toPlanar8(_:_:_:)
func VImageConvert_Planar4toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar4toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_Planar4toPlanar8 not loaded")
	}
	return _vImageConvert_Planar4toPlanar8(src, dest, flags)
}


var _vImageConvert_Planar8To16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8To16U converts an 8-bit planar buffer to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8To16U(_:_:_:)
func VImageConvert_Planar8To16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8To16U == nil {
		panic("Accelerate: symbol vImageConvert_Planar8To16U not loaded")
	}
	return _vImageConvert_Planar8To16U(src, dest, flags)
}


var _vImageConvert_Planar8ToARGBFFFF func(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_Planar8ToARGBFFFF interleaves four 8-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel interleaved ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8ToARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageConvert_Planar8ToARGBFFFF(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_Planar8ToARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_Planar8ToARGBFFFF not loaded")
	}
	return _vImageConvert_Planar8ToARGBFFFF(alpha, red, green, blue, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_Planar8ToBGRX8888 func(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8ToBGRX8888 interleaves three 8-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved BGRX buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8ToBGRX8888(_:_:_:_:_:_:)
func VImageConvert_Planar8ToBGRX8888(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8ToBGRX8888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8ToBGRX8888 not loaded")
	}
	return _vImageConvert_Planar8ToBGRX8888(blue, green, red, alpha, dest, flags)
}


var _vImageConvert_Planar8ToBGRXFFFF func(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_Planar8ToBGRXFFFF interleaves three 8-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel interleaved BGRX buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8ToBGRXFFFF(_:_:_:_:_:_:_:_:)
func VImageConvert_Planar8ToBGRXFFFF(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_Planar8ToBGRXFFFF == nil {
		panic("Accelerate: symbol vImageConvert_Planar8ToBGRXFFFF not loaded")
	}
	return _vImageConvert_Planar8ToBGRXFFFF(blue, green, red, alpha, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_Planar8ToXRGB8888 func(alpha Pixel_8, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8ToXRGB8888 interleaves three 8-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved XRGB buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8ToXRGB8888(_:_:_:_:_:_:)
func VImageConvert_Planar8ToXRGB8888(alpha Pixel_8, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8ToXRGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8ToXRGB8888 not loaded")
	}
	return _vImageConvert_Planar8ToXRGB8888(alpha, red, green, blue, dest, flags)
}


var _vImageConvert_Planar8ToXRGBFFFF func(alpha Pixel_F, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_Planar8ToXRGBFFFF interleaves three 8-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel interleaved XRGB buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8ToXRGBFFFF(_:_:_:_:_:_:_:_:)
func VImageConvert_Planar8ToXRGBFFFF(alpha Pixel_F, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_Planar8ToXRGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_Planar8ToXRGBFFFF not loaded")
	}
	return _vImageConvert_Planar8ToXRGBFFFF(alpha, red, green, blue, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_Planar8toARGB1555 func(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8toARGB1555 interleaves four 8-bit planar buffers into an ARGB1555 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toARGB1555(_:_:_:_:_:_:)
func VImageConvert_Planar8toARGB1555(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8toARGB1555 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toARGB1555 not loaded")
	}
	return _vImageConvert_Planar8toARGB1555(srcA, srcR, srcG, srcB, dest, flags)
}


var _vImageConvert_Planar8toARGB8888 func(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8toARGB8888 interleaves four 8-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toARGB8888(_:_:_:_:_:_:)
func VImageConvert_Planar8toARGB8888(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8toARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toARGB8888 not loaded")
	}
	return _vImageConvert_Planar8toARGB8888(srcA, srcR, srcG, srcB, dest, flags)
}


var _vImageConvert_Planar8toIndexed1 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int

// VImageConvert_Planar8toIndexed1 converts an 8-bit planar buffer to an indexed 1-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toIndexed1(_:_:_:_:_:_:)
func VImageConvert_Planar8toIndexed1(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int {
	if _vImageConvert_Planar8toIndexed1 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toIndexed1 not loaded")
	}
	return _vImageConvert_Planar8toIndexed1(src, dest, tempBuffer, colors, dither, flags)
}


var _vImageConvert_Planar8toIndexed2 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int

// VImageConvert_Planar8toIndexed2 converts an 8-bit planar buffer to an indexed 2-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toIndexed2(_:_:_:_:_:_:)
func VImageConvert_Planar8toIndexed2(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int {
	if _vImageConvert_Planar8toIndexed2 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toIndexed2 not loaded")
	}
	return _vImageConvert_Planar8toIndexed2(src, dest, tempBuffer, colors, dither, flags)
}


var _vImageConvert_Planar8toIndexed4 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int

// VImageConvert_Planar8toIndexed4 converts an 8-bit planar buffer to an indexed 4-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toIndexed4(_:_:_:_:_:_:)
func VImageConvert_Planar8toIndexed4(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, colors Pixel_8, dither int, flags uint32) int {
	if _vImageConvert_Planar8toIndexed4 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toIndexed4 not loaded")
	}
	return _vImageConvert_Planar8toIndexed4(src, dest, tempBuffer, colors, dither, flags)
}


var _vImageConvert_Planar8toPlanar1 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_Planar8toPlanar1 converts an 8-bit planar buffer to a 1-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toPlanar1(_:_:_:_:_:)
func VImageConvert_Planar8toPlanar1(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_Planar8toPlanar1 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toPlanar1 not loaded")
	}
	return _vImageConvert_Planar8toPlanar1(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_Planar8toPlanar16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8toPlanar16F converts an 8-bit planar buffer to a floating-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toPlanar16F(_:_:_:)
func VImageConvert_Planar8toPlanar16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8toPlanar16F == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toPlanar16F not loaded")
	}
	return _vImageConvert_Planar8toPlanar16F(src, dest, flags)
}


var _vImageConvert_Planar8toPlanar2 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_Planar8toPlanar2 converts an 8-bit planar buffer to a 2-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toPlanar2(_:_:_:_:_:)
func VImageConvert_Planar8toPlanar2(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_Planar8toPlanar2 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toPlanar2 not loaded")
	}
	return _vImageConvert_Planar8toPlanar2(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_Planar8toPlanar4 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_Planar8toPlanar4 converts an 8-bit planar buffer to a 4-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toPlanar4(_:_:_:_:_:)
func VImageConvert_Planar8toPlanar4(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_Planar8toPlanar4 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toPlanar4 not loaded")
	}
	return _vImageConvert_Planar8toPlanar4(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_Planar8toPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int

// VImageConvert_Planar8toPlanarF converts an 8-bit planar buffer to a floating-point 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toPlanarF(_:_:_:_:_:)
func VImageConvert_Planar8toPlanarF(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int {
	if _vImageConvert_Planar8toPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toPlanarF not loaded")
	}
	return _vImageConvert_Planar8toPlanarF(src, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_Planar8toRGB565 func(srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8toRGB565 interleaves three 8-bit planar buffers into an RGB565 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toRGB565(_:_:_:_:_:)
func VImageConvert_Planar8toRGB565(srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toRGB565 not loaded")
	}
	return _vImageConvert_Planar8toRGB565(srcR, srcG, srcB, dest, flags)
}


var _vImageConvert_Planar8toRGB888 func(planarRed unsafe.Pointer, planarGreen unsafe.Pointer, planarBlue unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_Planar8toRGB888 interleaves three 8-bit planar buffers into an 8-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_Planar8toRGB888(_:_:_:_:_:)
func VImageConvert_Planar8toRGB888(planarRed unsafe.Pointer, planarGreen unsafe.Pointer, planarBlue unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_Planar8toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_Planar8toRGB888 not loaded")
	}
	return _vImageConvert_Planar8toRGB888(planarRed, planarGreen, planarBlue, rgbDest, flags)
}


var _vImageConvert_PlanarFToARGB8888 func(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_PlanarFToARGB8888 interleaves four 32-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFToARGB8888(_:_:_:_:_:_:_:_:)
func VImageConvert_PlanarFToARGB8888(alpha unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_PlanarFToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFToARGB8888 not loaded")
	}
	return _vImageConvert_PlanarFToARGB8888(alpha, red, green, blue, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_PlanarFToBGRX8888 func(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_PlanarFToBGRX8888 interleaves three 32-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved BGRX buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFToBGRX8888(_:_:_:_:_:_:_:_:)
func VImageConvert_PlanarFToBGRX8888(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_PlanarFToBGRX8888 == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFToBGRX8888 not loaded")
	}
	return _vImageConvert_PlanarFToBGRX8888(blue, green, red, alpha, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_PlanarFToBGRXFFFF func(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, flags uint32) int

// VImageConvert_PlanarFToBGRXFFFF interleaves four floating-point 32-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel BGRXARGB interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFToBGRXFFFF(_:_:_:_:_:_:)
func VImageConvert_PlanarFToBGRXFFFF(blue unsafe.Pointer, green unsafe.Pointer, red unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_PlanarFToBGRXFFFF == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFToBGRXFFFF not loaded")
	}
	return _vImageConvert_PlanarFToBGRXFFFF(blue, green, red, alpha, dest, flags)
}


var _vImageConvert_PlanarFToXRGB8888 func(alpha Pixel_8, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int

// VImageConvert_PlanarFToXRGB8888 interleaves three 32-bit planar buffers into an 8-bit-per-channel, 4-channel interleaved XRGB buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFToXRGB8888(_:_:_:_:_:_:_:_:)
func VImageConvert_PlanarFToXRGB8888(alpha Pixel_8, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_FFFF, minFloat Pixel_FFFF, flags uint32) int {
	if _vImageConvert_PlanarFToXRGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFToXRGB8888 not loaded")
	}
	return _vImageConvert_PlanarFToXRGB8888(alpha, red, green, blue, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_PlanarFToXRGBFFFF func(alpha Pixel_F, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_PlanarFToXRGBFFFF interleaves three 32-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel interleaved XRGB buffer with the specified constant alpha value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFToXRGBFFFF(_:_:_:_:_:_:)
func VImageConvert_PlanarFToXRGBFFFF(alpha Pixel_F, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_PlanarFToXRGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFToXRGBFFFF not loaded")
	}
	return _vImageConvert_PlanarFToXRGBFFFF(alpha, red, green, blue, dest, flags)
}


var _vImageConvert_PlanarFtoARGBFFFF func(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_PlanarFtoARGBFFFF interleaves four floating-point 32-bit planar buffers into a floating-point 32-bit-per-channel, 4-channel ARGB interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFtoARGBFFFF(_:_:_:_:_:_:)
func VImageConvert_PlanarFtoARGBFFFF(srcA unsafe.Pointer, srcR unsafe.Pointer, srcG unsafe.Pointer, srcB unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_PlanarFtoARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFtoARGBFFFF not loaded")
	}
	return _vImageConvert_PlanarFtoARGBFFFF(srcA, srcR, srcG, srcB, dest, flags)
}


var _vImageConvert_PlanarFtoPlanar16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_PlanarFtoPlanar16F converts a floating-point 32-bit planar buffer to a floating-point 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFtoPlanar16F(_:_:_:)
func VImageConvert_PlanarFtoPlanar16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_PlanarFtoPlanar16F == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFtoPlanar16F not loaded")
	}
	return _vImageConvert_PlanarFtoPlanar16F(src, dest, flags)
}


var _vImageConvert_PlanarFtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int

// VImageConvert_PlanarFtoPlanar8 converts a floating-point 32-bit planar buffer to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFtoPlanar8(_:_:_:_:_:)
func VImageConvert_PlanarFtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, flags uint32) int {
	if _vImageConvert_PlanarFtoPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFtoPlanar8 not loaded")
	}
	return _vImageConvert_PlanarFtoPlanar8(src, dest, maxFloat, minFloat, flags)
}


var _vImageConvert_PlanarFtoPlanar8_dithered func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, dither int, flags uint32) int

// VImageConvert_PlanarFtoPlanar8_dithered converts a floating-point 32-bit planar buffer to an 8-bit planar buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFtoPlanar8_dithered(_:_:_:_:_:_:)
func VImageConvert_PlanarFtoPlanar8_dithered(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, dither int, flags uint32) int {
	if _vImageConvert_PlanarFtoPlanar8_dithered == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFtoPlanar8_dithered not loaded")
	}
	return _vImageConvert_PlanarFtoPlanar8_dithered(src, dest, maxFloat, minFloat, dither, flags)
}


var _vImageConvert_PlanarFtoRGBFFF func(planarRed unsafe.Pointer, planarGreen unsafe.Pointer, planarBlue unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_PlanarFtoRGBFFF interleaves three floating-point 32-bit planar buffers into a floating-point 32-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarFtoRGBFFF(_:_:_:_:_:)
func VImageConvert_PlanarFtoRGBFFF(planarRed unsafe.Pointer, planarGreen unsafe.Pointer, planarBlue unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_PlanarFtoRGBFFF == nil {
		panic("Accelerate: symbol vImageConvert_PlanarFtoRGBFFF not loaded")
	}
	return _vImageConvert_PlanarFtoRGBFFF(planarRed, planarGreen, planarBlue, rgbDest, flags)
}


var _vImageConvert_PlanarToChunky8 func(srcPlanarBuffers unsafe.Pointer, destChannels unsafe.Pointer, channelCount uint, destStrideBytes uintptr, destWidth uint, destHeight uint, destRowBytes uintptr, flags uint32) int

// VImageConvert_PlanarToChunky8 interleaves the specifed number of 8-bit planar buffers into an 8-bit-per-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarToChunky8(_:_:_:_:_:_:_:_:)
func VImageConvert_PlanarToChunky8(srcPlanarBuffers unsafe.Pointer, destChannels unsafe.Pointer, channelCount uint, destStrideBytes uintptr, destWidth uint, destHeight uint, destRowBytes uintptr, flags uint32) int {
	if _vImageConvert_PlanarToChunky8 == nil {
		panic("Accelerate: symbol vImageConvert_PlanarToChunky8 not loaded")
	}
	return _vImageConvert_PlanarToChunky8(srcPlanarBuffers, destChannels, channelCount, destStrideBytes, destWidth, destHeight, destRowBytes, flags)
}


var _vImageConvert_PlanarToChunkyF func(srcPlanarBuffers unsafe.Pointer, destChannels unsafe.Pointer, channelCount uint, destStrideBytes uintptr, destWidth uint, destHeight uint, destRowBytes uintptr, flags uint32) int

// VImageConvert_PlanarToChunkyF interleaves the specifed number of floating-point 32-bit planar buffers into a floating-point 32 -bit-per-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_PlanarToChunkyF(_:_:_:_:_:_:_:_:)
func VImageConvert_PlanarToChunkyF(srcPlanarBuffers unsafe.Pointer, destChannels unsafe.Pointer, channelCount uint, destStrideBytes uintptr, destWidth uint, destHeight uint, destRowBytes uintptr, flags uint32) int {
	if _vImageConvert_PlanarToChunkyF == nil {
		panic("Accelerate: symbol vImageConvert_PlanarToChunkyF not loaded")
	}
	return _vImageConvert_PlanarToChunkyF(srcPlanarBuffers, destChannels, channelCount, destStrideBytes, destWidth, destHeight, destRowBytes, flags)
}


var _vImageConvert_RGB16UToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int

// VImageConvert_RGB16UToARGB8888 converts an unsigned 16-bit-per-channel, 3-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer using permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UToARGB8888(_:_:_:_:_:_:)
func VImageConvert_RGB16UToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvert_RGB16UToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UToARGB8888 not loaded")
	}
	return _vImageConvert_RGB16UToARGB8888(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImageConvert_RGB16UtoARGB16U func(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, argbDest unsafe.Pointer, premultiply bool, flags uint32) int

// VImageConvert_RGB16UtoARGB16U combines an unsigned 16-bit-per-channel, 3-channel RGB buffer and either an unsigned 16-bit alpha buffer or constant alpha value to produce an ARGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UtoARGB16U(_:_:_:_:_:_:)
func VImageConvert_RGB16UtoARGB16U(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, argbDest unsafe.Pointer, premultiply bool, flags uint32) int {
	if _vImageConvert_RGB16UtoARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UtoARGB16U not loaded")
	}
	return _vImageConvert_RGB16UtoARGB16U(rgbSrc, aSrc, alpha, argbDest, premultiply, flags)
}


var _vImageConvert_RGB16UtoBGRA16U func(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, bgraDest unsafe.Pointer, premultiply bool, flags uint32) int

// VImageConvert_RGB16UtoBGRA16U combines an unsigned 16-bit-per-channel, 3-channel RGB buffer and either an unsigned 16-bit alpha buffer or constant alpha value to produce a BGRA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UtoBGRA16U(_:_:_:_:_:_:)
func VImageConvert_RGB16UtoBGRA16U(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, bgraDest unsafe.Pointer, premultiply bool, flags uint32) int {
	if _vImageConvert_RGB16UtoBGRA16U == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UtoBGRA16U not loaded")
	}
	return _vImageConvert_RGB16UtoBGRA16U(rgbSrc, aSrc, alpha, bgraDest, premultiply, flags)
}


var _vImageConvert_RGB16UtoPlanar16U func(rgbSrc unsafe.Pointer, rDest unsafe.Pointer, gDest unsafe.Pointer, bDest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB16UtoPlanar16U deinterleaves an unsigned 16-bit-per-channel, 3-channel interleaved buffer into three unsigned 16-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UtoPlanar16U(_:_:_:_:_:)
func VImageConvert_RGB16UtoPlanar16U(rgbSrc unsafe.Pointer, rDest unsafe.Pointer, gDest unsafe.Pointer, bDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB16UtoPlanar16U == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UtoPlanar16U not loaded")
	}
	return _vImageConvert_RGB16UtoPlanar16U(rgbSrc, rDest, gDest, bDest, flags)
}


var _vImageConvert_RGB16UtoRGB888_dithered func(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGB16UtoRGB888_dithered converts an unsigned 16-bit-per-channel, 3-channel interleaved buffer to an 8-bit-per-channel, 3-channel interleaved buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UtoRGB888_dithered(_:_:_:_:)
func VImageConvert_RGB16UtoRGB888_dithered(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGB16UtoRGB888_dithered == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UtoRGB888_dithered not loaded")
	}
	return _vImageConvert_RGB16UtoRGB888_dithered(src, dest, dither, flags)
}


var _vImageConvert_RGB16UtoRGBA16U func(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, rgbaDest unsafe.Pointer, premultiply bool, flags uint32) int

// VImageConvert_RGB16UtoRGBA16U combines an unsigned 16-bit-per-channel, 3-channel RGB buffer and either an unsigned 16-bit alpha buffer or constant alpha value to produce an RGBA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB16UtoRGBA16U(_:_:_:_:_:_:)
func VImageConvert_RGB16UtoRGBA16U(rgbSrc unsafe.Pointer, aSrc unsafe.Pointer, alpha Pixel_16U, rgbaDest unsafe.Pointer, premultiply bool, flags uint32) int {
	if _vImageConvert_RGB16UtoRGBA16U == nil {
		panic("Accelerate: symbol vImageConvert_RGB16UtoRGBA16U not loaded")
	}
	return _vImageConvert_RGB16UtoRGBA16U(rgbSrc, aSrc, alpha, rgbaDest, premultiply, flags)
}


var _vImageConvert_RGB565toARGB1555 func(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGB565toARGB1555 combines an RGB565 3-channel RGB buffer and a constant alpha value to produce a 4-channel ARGB1555 buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toARGB1555(_:_:_:_:)
func VImageConvert_RGB565toARGB1555(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGB565toARGB1555 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toARGB1555 not loaded")
	}
	return _vImageConvert_RGB565toARGB1555(src, dest, dither, flags)
}


var _vImageConvert_RGB565toARGB8888 func(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB565toARGB8888 combines an RGB565 3-channel RGB buffer and a constant alpha value to produce an 8-bit-per-channel, 4-channel ARGB buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toARGB8888(_:_:_:_:)
func VImageConvert_RGB565toARGB8888(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB565toARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toARGB8888 not loaded")
	}
	return _vImageConvert_RGB565toARGB8888(alpha, src, dest, flags)
}


var _vImageConvert_RGB565toBGRA8888 func(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB565toBGRA8888 combines an RGB565 3-channel RGB buffer and a constant alpha value to produce an 8-bit-per-channel, 4-channel BGRA buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toBGRA8888(_:_:_:_:)
func VImageConvert_RGB565toBGRA8888(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB565toBGRA8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toBGRA8888 not loaded")
	}
	return _vImageConvert_RGB565toBGRA8888(alpha, src, dest, flags)
}


var _vImageConvert_RGB565toPlanar8 func(src unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int

// VImageConvert_RGB565toPlanar8 deinterleaves an RGB565 3-channel interleaved buffer into three 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toPlanar8(_:_:_:_:_:)
func VImageConvert_RGB565toPlanar8(src unsafe.Pointer, destR unsafe.Pointer, destG unsafe.Pointer, destB unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB565toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toPlanar8 not loaded")
	}
	return _vImageConvert_RGB565toPlanar8(src, destR, destG, destB, flags)
}


var _vImageConvert_RGB565toRGB888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB565toRGB888 converts an RGB565 3-channel interleaved buffer to an 8-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toRGB888(_:_:_:)
func VImageConvert_RGB565toRGB888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB565toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toRGB888 not loaded")
	}
	return _vImageConvert_RGB565toRGB888(src, dest, flags)
}


var _vImageConvert_RGB565toRGBA5551 func(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGB565toRGBA5551 combines an RGB565 3-channel RGB buffer and a constant alpha value to produce a 4-channel RGBA5551 buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toRGBA5551(_:_:_:_:)
func VImageConvert_RGB565toRGBA5551(src unsafe.Pointer, dest unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGB565toRGBA5551 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toRGBA5551 not loaded")
	}
	return _vImageConvert_RGB565toRGBA5551(src, dest, dither, flags)
}


var _vImageConvert_RGB565toRGBA8888 func(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB565toRGBA8888 combines an RGB565 3-channel RGB buffer and a constant alpha value to produce an 8-bit-per-channel, 4-channel RGBA buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB565toRGBA8888(_:_:_:_:)
func VImageConvert_RGB565toRGBA8888(alpha Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB565toRGBA8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB565toRGBA8888 not loaded")
	}
	return _vImageConvert_RGB565toRGBA8888(alpha, src, dest, flags)
}


var _vImageConvert_RGB888toARGB8888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int

// VImageConvert_RGB888toARGB8888 combines an 8-bit-per-channel, 3-channel RGB buffer and either an 8-bit alpha buffer or constant alpha value to produce an ARGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toARGB8888(_:_:_:_:_:_:)
func VImageConvert_RGB888toARGB8888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int {
	if _vImageConvert_RGB888toARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toARGB8888 not loaded")
	}
	return _vImageConvert_RGB888toARGB8888(arg0, arg1, arg2, arg3, arg4, arg5)
}


var _vImageConvert_RGB888toBGRA8888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int

// VImageConvert_RGB888toBGRA8888 combines an 8-bit-per-channel, 3-channel RGB buffer and either an 8-bit alpha buffer or constant alpha value to produce a BGRA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toBGRA8888(_:_:_:_:_:_:)
func VImageConvert_RGB888toBGRA8888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int {
	if _vImageConvert_RGB888toBGRA8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toBGRA8888 not loaded")
	}
	return _vImageConvert_RGB888toBGRA8888(arg0, arg1, arg2, arg3, arg4, arg5)
}


var _vImageConvert_RGB888toPlanar16Q12 func(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int

// VImageConvert_RGB888toPlanar16Q12 deinterleaves an 8-bit-per-channel, 3-channel interleaved buffer into three fixed-point 16-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toPlanar16Q12(_:_:_:_:_:)
func VImageConvert_RGB888toPlanar16Q12(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB888toPlanar16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toPlanar16Q12 not loaded")
	}
	return _vImageConvert_RGB888toPlanar16Q12(src, red, green, blue, flags)
}


var _vImageConvert_RGB888toPlanar8 func(rgbSrc unsafe.Pointer, redDest unsafe.Pointer, greenDest unsafe.Pointer, blueDest unsafe.Pointer, flags uint32) int

// VImageConvert_RGB888toPlanar8 deinterleaves an 8-bit-per-channel, 3-channel interleaved buffer into three 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toPlanar8(_:_:_:_:_:)
func VImageConvert_RGB888toPlanar8(rgbSrc unsafe.Pointer, redDest unsafe.Pointer, greenDest unsafe.Pointer, blueDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGB888toPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toPlanar8 not loaded")
	}
	return _vImageConvert_RGB888toPlanar8(rgbSrc, redDest, greenDest, blueDest, flags)
}


var _vImageConvert_RGB888toRGB565_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGB888toRGB565_dithered converts an 8-bit-per-channel, 3-channel interleaved buffer to an RGB565 3-channel interleaved buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toRGB565_dithered(_:_:_:_:_:)
func VImageConvert_RGB888toRGB565_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGB888toRGB565_dithered == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toRGB565_dithered not loaded")
	}
	return _vImageConvert_RGB888toRGB565_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_RGB888toRGBA8888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int

// VImageConvert_RGB888toRGBA8888 combines an 8-bit-per-channel, 3-channel RGB buffer and either an 8-bit alpha buffer or constant alpha value to produce an RGBA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGB888toRGBA8888(_:_:_:_:_:_:)
func VImageConvert_RGB888toRGBA8888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8, arg3 unsafe.Pointer, arg4 bool, arg5 uint32) int {
	if _vImageConvert_RGB888toRGBA8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGB888toRGBA8888 not loaded")
	}
	return _vImageConvert_RGB888toRGBA8888(arg0, arg1, arg2, arg3, arg4, arg5)
}


var _vImageConvert_RGBA1010102ToARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_RGBA1010102ToARGB16Q12 converts an RGBA1010102 32-bit, 4-channel interleaved buffer to a fixed-point 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA1010102ToARGB16Q12(_:_:_:_:_:_:)
func VImageConvert_RGBA1010102ToARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_RGBA1010102ToARGB16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA1010102ToARGB16Q12 not loaded")
	}
	return _vImageConvert_RGBA1010102ToARGB16Q12(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_RGBA1010102ToARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_RGBA1010102ToARGB16U converts an RGBA1010102 32-bit, 4-channel interleaved buffer to an unsigned 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA1010102ToARGB16U(_:_:_:_:_:_:)
func VImageConvert_RGBA1010102ToARGB16U(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_RGBA1010102ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_RGBA1010102ToARGB16U not loaded")
	}
	return _vImageConvert_RGBA1010102ToARGB16U(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_RGBA1010102ToARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_RGBA1010102ToARGB8888 converts an RGBA1010102 32-bit, 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA1010102ToARGB8888(_:_:_:_:_:_:)
func VImageConvert_RGBA1010102ToARGB8888(src unsafe.Pointer, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_RGBA1010102ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA1010102ToARGB8888 not loaded")
	}
	return _vImageConvert_RGBA1010102ToARGB8888(src, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_RGBA16UtoRGB16U func(rgbaSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBA16UtoRGB16U removes the alpha channel from an unsigned 16-bit-per-channel RGBA buffer to produce an unsigned 16-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA16UtoRGB16U(_:_:_:)
func VImageConvert_RGBA16UtoRGB16U(rgbaSrc unsafe.Pointer, rgbDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBA16UtoRGB16U == nil {
		panic("Accelerate: symbol vImageConvert_RGBA16UtoRGB16U not loaded")
	}
	return _vImageConvert_RGBA16UtoRGB16U(rgbaSrc, rgbDest, flags)
}


var _vImageConvert_RGBA5551toRGB565 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBA5551toRGB565 removes the alpha channel from an RGBA5551 buffer to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA5551toRGB565(_:_:_:)
func VImageConvert_RGBA5551toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBA5551toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA5551toRGB565 not loaded")
	}
	return _vImageConvert_RGBA5551toRGB565(src, dest, flags)
}


var _vImageConvert_RGBA5551toRGBA8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBA5551toRGBA8888 converts an RGB5651 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA5551toRGBA8888(_:_:_:)
func VImageConvert_RGBA5551toRGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBA5551toRGBA8888 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA5551toRGBA8888 not loaded")
	}
	return _vImageConvert_RGBA5551toRGBA8888(src, dest, flags)
}


var _vImageConvert_RGBA8888toRGB565 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBA8888toRGB565 removes the alpha channel from an 8-bit-per-channel RGBA buffer to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGB565(_:_:_:)
func VImageConvert_RGBA8888toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBA8888toRGB565 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA8888toRGB565 not loaded")
	}
	return _vImageConvert_RGBA8888toRGB565(src, dest, flags)
}


var _vImageConvert_RGBA8888toRGB565_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGBA8888toRGB565_dithered removes the alpha channel from an 8-bit-per-channel RGBA buffer using the specified dithering algorithm to produce an RGB565 result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGB565_dithered(_:_:_:_:_:)
func VImageConvert_RGBA8888toRGB565_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGBA8888toRGB565_dithered == nil {
		panic("Accelerate: symbol vImageConvert_RGBA8888toRGB565_dithered not loaded")
	}
	return _vImageConvert_RGBA8888toRGB565_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_RGBA8888toRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int

// VImageConvert_RGBA8888toRGB888 removes the alpha channel from an 8-bit-per-channel RGBA buffer to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGB888(_:_:_:)
func VImageConvert_RGBA8888toRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint32) int {
	if _vImageConvert_RGBA8888toRGB888 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA8888toRGB888 not loaded")
	}
	return _vImageConvert_RGBA8888toRGB888(arg0, arg1, arg2)
}


var _vImageConvert_RGBA8888toRGBA5551 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBA8888toRGBA5551 converts an 8-bit-per-channel, 4-channel interleaved buffer to an RGBA5551 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGBA5551(_:_:_:)
func VImageConvert_RGBA8888toRGBA5551(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBA8888toRGBA5551 == nil {
		panic("Accelerate: symbol vImageConvert_RGBA8888toRGBA5551 not loaded")
	}
	return _vImageConvert_RGBA8888toRGBA5551(src, dest, flags)
}


var _vImageConvert_RGBA8888toRGBA5551_dithered func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int

// VImageConvert_RGBA8888toRGBA5551_dithered converts an 8-bit-per-channel, 4-channel interleaved buffer to an RGBA5551 4-channel interleaved buffer usng the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGBA5551_dithered(_:_:_:_:_:)
func VImageConvert_RGBA8888toRGBA5551_dithered(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, dither int, flags uint32) int {
	if _vImageConvert_RGBA8888toRGBA5551_dithered == nil {
		panic("Accelerate: symbol vImageConvert_RGBA8888toRGBA5551_dithered not loaded")
	}
	return _vImageConvert_RGBA8888toRGBA5551_dithered(src, dest, tempBuffer, dither, flags)
}


var _vImageConvert_RGBAFFFFtoRGBFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBAFFFFtoRGBFFF removes the alpha channel from a floating-point 32-bit-per-channel RGBA buffer to produce a floating-point 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBAFFFFtoRGBFFF(_:_:_:)
func VImageConvert_RGBAFFFFtoRGBFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBAFFFFtoRGBFFF == nil {
		panic("Accelerate: symbol vImageConvert_RGBAFFFFtoRGBFFF not loaded")
	}
	return _vImageConvert_RGBAFFFFtoRGBFFF(src, dest, flags)
}


var _vImageConvert_RGBFFFtoARGBFFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int

// VImageConvert_RGBFFFtoARGBFFFF combines a floating-point 32-bit-per-channel, 3-channel RGB buffer and either an 32-bit alpha buffer or constant alpha value to produce an ARGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBFFFtoARGBFFFF(_:_:_:_:_:_:)
func VImageConvert_RGBFFFtoARGBFFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int {
	if _vImageConvert_RGBFFFtoARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_RGBFFFtoARGBFFFF not loaded")
	}
	return _vImageConvert_RGBFFFtoARGBFFFF(arg0, arg1, arg2, arg3, arg4, flags)
}


var _vImageConvert_RGBFFFtoBGRAFFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int

// VImageConvert_RGBFFFtoBGRAFFFF combines a floating-point 32-bit-per-channel, 3-channel RGB buffer and either an 32-bit alpha buffer or constant alpha value to produce a BGRA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBFFFtoBGRAFFFF(_:_:_:_:_:_:)
func VImageConvert_RGBFFFtoBGRAFFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int {
	if _vImageConvert_RGBFFFtoBGRAFFFF == nil {
		panic("Accelerate: symbol vImageConvert_RGBFFFtoBGRAFFFF not loaded")
	}
	return _vImageConvert_RGBFFFtoBGRAFFFF(arg0, arg1, arg2, arg3, arg4, flags)
}


var _vImageConvert_RGBFFFtoPlanarF func(rgbSrc unsafe.Pointer, redDest unsafe.Pointer, greenDest unsafe.Pointer, blueDest unsafe.Pointer, flags uint32) int

// VImageConvert_RGBFFFtoPlanarF deinterleaves a floating-point 32-bit-per-channel, 3-channel interleaved buffer into three floating-point 32-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBFFFtoPlanarF(_:_:_:_:_:)
func VImageConvert_RGBFFFtoPlanarF(rgbSrc unsafe.Pointer, redDest unsafe.Pointer, greenDest unsafe.Pointer, blueDest unsafe.Pointer, flags uint32) int {
	if _vImageConvert_RGBFFFtoPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_RGBFFFtoPlanarF not loaded")
	}
	return _vImageConvert_RGBFFFtoPlanarF(rgbSrc, redDest, greenDest, blueDest, flags)
}


var _vImageConvert_RGBFFFtoRGB888_dithered func(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, dither int, flags uint32) int

// VImageConvert_RGBFFFtoRGB888_dithered converts a floating-point 32-bit-per-channel, 3-channel buffer to an 8-bit-per-channel, 3-channel buffer using the specified dithering algorithm.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBFFFtoRGB888_dithered(_:_:_:_:_:_:)
func VImageConvert_RGBFFFtoRGB888_dithered(src unsafe.Pointer, dest unsafe.Pointer, maxFloat Pixel_F, minFloat Pixel_F, dither int, flags uint32) int {
	if _vImageConvert_RGBFFFtoRGB888_dithered == nil {
		panic("Accelerate: symbol vImageConvert_RGBFFFtoRGB888_dithered not loaded")
	}
	return _vImageConvert_RGBFFFtoRGB888_dithered(src, dest, maxFloat, minFloat, dither, flags)
}


var _vImageConvert_RGBFFFtoRGBAFFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int

// VImageConvert_RGBFFFtoRGBAFFFF combines a floating-point 32-bit-per-channel, 3-channel RGB buffer and either an 32-bit alpha buffer or constant alpha value to produce an RGBA result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBFFFtoRGBAFFFF(_:_:_:_:_:_:)
func VImageConvert_RGBFFFtoRGBAFFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_F, arg3 unsafe.Pointer, arg4 bool, flags uint32) int {
	if _vImageConvert_RGBFFFtoRGBAFFFF == nil {
		panic("Accelerate: symbol vImageConvert_RGBFFFtoRGBAFFFF not loaded")
	}
	return _vImageConvert_RGBFFFtoRGBAFFFF(arg0, arg1, arg2, arg3, arg4, flags)
}


var _vImageConvert_XRGB2101010ToARGB16F func(src unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_XRGB2101010ToARGB16F converts an XRGB2101010 32-bit, 4-channel interleaved buffer to a floating-point 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB2101010ToARGB16F(_:_:_:_:_:_:_:)
func VImageConvert_XRGB2101010ToARGB16F(src unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_XRGB2101010ToARGB16F == nil {
		panic("Accelerate: symbol vImageConvert_XRGB2101010ToARGB16F not loaded")
	}
	return _vImageConvert_XRGB2101010ToARGB16F(src, alpha, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_XRGB2101010ToARGB16Q12 func(src unsafe.Pointer, alpha Pixel_16Q12, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_XRGB2101010ToARGB16Q12 converts an XRGB2101010 32-bit, 4-channel interleaved buffer to a fixed-point 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB2101010ToARGB16Q12(_:_:_:_:_:_:_:)
func VImageConvert_XRGB2101010ToARGB16Q12(src unsafe.Pointer, alpha Pixel_16Q12, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_XRGB2101010ToARGB16Q12 == nil {
		panic("Accelerate: symbol vImageConvert_XRGB2101010ToARGB16Q12 not loaded")
	}
	return _vImageConvert_XRGB2101010ToARGB16Q12(src, alpha, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_XRGB2101010ToARGB16U func(src unsafe.Pointer, alpha uint16, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_XRGB2101010ToARGB16U converts an XRGB2101010 32-bit, 4-channel interleaved buffer to an unsigned 16-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB2101010ToARGB16U(_:_:_:_:_:_:_:)
func VImageConvert_XRGB2101010ToARGB16U(src unsafe.Pointer, alpha uint16, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_XRGB2101010ToARGB16U == nil {
		panic("Accelerate: symbol vImageConvert_XRGB2101010ToARGB16U not loaded")
	}
	return _vImageConvert_XRGB2101010ToARGB16U(src, alpha, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_XRGB2101010ToARGB8888 func(src unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_XRGB2101010ToARGB8888 converts an XRGB2101010 32-bit, 4-channel interleaved buffer to an 8-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB2101010ToARGB8888(_:_:_:_:_:_:_:)
func VImageConvert_XRGB2101010ToARGB8888(src unsafe.Pointer, alpha Pixel_8, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_XRGB2101010ToARGB8888 == nil {
		panic("Accelerate: symbol vImageConvert_XRGB2101010ToARGB8888 not loaded")
	}
	return _vImageConvert_XRGB2101010ToARGB8888(src, alpha, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_XRGB2101010ToARGBFFFF func(src unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int

// VImageConvert_XRGB2101010ToARGBFFFF converts an XRGB2101010 32-bit, 4-channel interleaved buffer to a floating-point 32-bit-per-channel, 4-channel interleaved buffer with permutation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB2101010ToARGBFFFF(_:_:_:_:_:_:_:)
func VImageConvert_XRGB2101010ToARGBFFFF(src unsafe.Pointer, alpha Pixel_F, dest unsafe.Pointer, RGB101010RangeMin int32, RGB101010RangeMax int32, permuteMap uint8, flags uint32) int {
	if _vImageConvert_XRGB2101010ToARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvert_XRGB2101010ToARGBFFFF not loaded")
	}
	return _vImageConvert_XRGB2101010ToARGBFFFF(src, alpha, dest, RGB101010RangeMin, RGB101010RangeMax, permuteMap, flags)
}


var _vImageConvert_XRGB8888ToPlanar8 func(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int

// VImageConvert_XRGB8888ToPlanar8 deinterleaves an 8-bit-per-channel, 4-channel interleaved buffer into three 8-bit planar buffers and discards the first channel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGB8888ToPlanar8(_:_:_:_:_:)
func VImageConvert_XRGB8888ToPlanar8(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int {
	if _vImageConvert_XRGB8888ToPlanar8 == nil {
		panic("Accelerate: symbol vImageConvert_XRGB8888ToPlanar8 not loaded")
	}
	return _vImageConvert_XRGB8888ToPlanar8(src, red, green, blue, flags)
}


var _vImageConvert_XRGBFFFFToPlanarF func(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int

// VImageConvert_XRGBFFFFToPlanarF deinterleaves a floating-point 32-bit-per-channel, 4-channel interleaved buffer into three floating-point 32-bit planar buffers and discards the first channel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_XRGBFFFFToPlanarF(_:_:_:_:_:)
func VImageConvert_XRGBFFFFToPlanarF(src unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, flags uint32) int {
	if _vImageConvert_XRGBFFFFToPlanarF == nil {
		panic("Accelerate: symbol vImageConvert_XRGBFFFFToPlanarF not loaded")
	}
	return _vImageConvert_XRGBFFFFToPlanarF(src, red, green, blue, flags)
}


var _vImageConvert_YpCbCrToARGB_GenerateConversion func(matrix *VImage_YpCbCrToARGBMatrix, pixelRange *VImage_YpCbCrPixelRange, outInfo *VImage_YpCbCrToARGB, inYpCbCrType unsafe.Pointer, outARGBType unsafe.Pointer, flags uint32) int

// VImageConvert_YpCbCrToARGB_GenerateConversion generates the information that describes the conversion from YpCbCr to ARGB.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvert_YpCbCrToARGB_GenerateConversion(_:_:_:_:_:_:)
func VImageConvert_YpCbCrToARGB_GenerateConversion(matrix *VImage_YpCbCrToARGBMatrix, pixelRange *VImage_YpCbCrPixelRange, outInfo *VImage_YpCbCrToARGB, inYpCbCrType unsafe.Pointer, outARGBType unsafe.Pointer, flags uint32) int {
	if _vImageConvert_YpCbCrToARGB_GenerateConversion == nil {
		panic("Accelerate: symbol vImageConvert_YpCbCrToARGB_GenerateConversion not loaded")
	}
	return _vImageConvert_YpCbCrToARGB_GenerateConversion(matrix, pixelRange, outInfo, inYpCbCrType, outARGBType, flags)
}


var _vImageConverter_CreateForCGToCVImageFormat func(srcFormat *VImage_CGImageFormat, destFormat uintptr, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer

// VImageConverter_CreateForCGToCVImageFormat creates a vImage converter that converts a Core Graphics-formatted image to a Core Video-formatted image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateForCGToCVImageFormat(_:_:_:_:_:)
func VImageConverter_CreateForCGToCVImageFormat(srcFormat *VImage_CGImageFormat, destFormat uintptr, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer {
	if _vImageConverter_CreateForCGToCVImageFormat == nil {
		panic("Accelerate: symbol vImageConverter_CreateForCGToCVImageFormat not loaded")
	}
	return _vImageConverter_CreateForCGToCVImageFormat(srcFormat, destFormat, backgroundColor, flags, err)
}


var _vImageConverter_CreateForCVToCGImageFormat func(srcFormat uintptr, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer

// VImageConverter_CreateForCVToCGImageFormat creates a vImage converter that converts a Core Video-formatted image to a Core Graphics-formatted image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateForCVToCGImageFormat(_:_:_:_:_:)
func VImageConverter_CreateForCVToCGImageFormat(srcFormat uintptr, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer {
	if _vImageConverter_CreateForCVToCGImageFormat == nil {
		panic("Accelerate: symbol vImageConverter_CreateForCVToCGImageFormat not loaded")
	}
	return _vImageConverter_CreateForCVToCGImageFormat(srcFormat, destFormat, backgroundColor, flags, err)
}


var _vImageConverter_CreateWithCGColorConversionInfo func(colorConversionInfoRef coregraphics.CGColorConversionInfoRef, sFormat *VImage_CGImageFormat, dFormat *VImage_CGImageFormat, bg *float64, flags uint32, err *int) unsafe.Pointer

// VImageConverter_CreateWithCGColorConversionInfo creates an any-to-any converter that uses a color conversion information object to convert from one image format to another.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithCGColorConversionInfo(_:_:_:_:_:_:)
func VImageConverter_CreateWithCGColorConversionInfo(colorConversionInfoRef coregraphics.CGColorConversionInfoRef, sFormat *VImage_CGImageFormat, dFormat *VImage_CGImageFormat, bg *float64, flags uint32, err *int) unsafe.Pointer {
	if _vImageConverter_CreateWithCGColorConversionInfo == nil {
		panic("Accelerate: symbol vImageConverter_CreateWithCGColorConversionInfo not loaded")
	}
	return _vImageConverter_CreateWithCGColorConversionInfo(colorConversionInfoRef, sFormat, dFormat, bg, flags, err)
}


var _vImageConverter_CreateWithCGImageFormat func(srcFormat *VImage_CGImageFormat, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer

// VImageConverter_CreateWithCGImageFormat creates a vImage converter that converts from one vImage Core Graphics image format to another.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithCGImageFormat(_:_:_:_:_:)
func VImageConverter_CreateWithCGImageFormat(srcFormat *VImage_CGImageFormat, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer {
	if _vImageConverter_CreateWithCGImageFormat == nil {
		panic("Accelerate: symbol vImageConverter_CreateWithCGImageFormat not loaded")
	}
	return _vImageConverter_CreateWithCGImageFormat(srcFormat, destFormat, backgroundColor, flags, err)
}


var _vImageConverter_CreateWithColorSyncCodeFragment func(codeFragment corefoundation.CFTypeRef, srcFormat *VImage_CGImageFormat, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer

// VImageConverter_CreateWithColorSyncCodeFragment creates a vImage converter to convert from one vImage Core Graphics image format to another, using custom ColorSync transform.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithColorSyncCodeFragment(_:_:_:_:_:_:)
func VImageConverter_CreateWithColorSyncCodeFragment(codeFragment corefoundation.CFTypeRef, srcFormat *VImage_CGImageFormat, destFormat *VImage_CGImageFormat, backgroundColor *float64, flags uint32, err *int) unsafe.Pointer {
	if _vImageConverter_CreateWithColorSyncCodeFragment == nil {
		panic("Accelerate: symbol vImageConverter_CreateWithColorSyncCodeFragment not loaded")
	}
	return _vImageConverter_CreateWithColorSyncCodeFragment(codeFragment, srcFormat, destFormat, backgroundColor, flags, err)
}



var _vImageConverter_GetNumberOfDestinationBuffers func(converter unsafe.Pointer) uint

// VImageConverter_GetNumberOfDestinationBuffers returns the number of destination buffers written to by the converter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_GetNumberOfDestinationBuffers(_:)
func VImageConverter_GetNumberOfDestinationBuffers(converter unsafe.Pointer) uint {
	if _vImageConverter_GetNumberOfDestinationBuffers == nil {
		panic("Accelerate: symbol vImageConverter_GetNumberOfDestinationBuffers not loaded")
	}
	return _vImageConverter_GetNumberOfDestinationBuffers(converter)
}


var _vImageConverter_GetNumberOfSourceBuffers func(converter unsafe.Pointer) uint

// VImageConverter_GetNumberOfSourceBuffers returns the number of source buffers consumed by the converter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_GetNumberOfSourceBuffers(_:)
func VImageConverter_GetNumberOfSourceBuffers(converter unsafe.Pointer) uint {
	if _vImageConverter_GetNumberOfSourceBuffers == nil {
		panic("Accelerate: symbol vImageConverter_GetNumberOfSourceBuffers not loaded")
	}
	return _vImageConverter_GetNumberOfSourceBuffers(converter)
}



var _vImageConverter_MustOperateOutOfPlace func(converter unsafe.Pointer, srcs unsafe.Pointer, dests unsafe.Pointer, flags uint32) int

// VImageConverter_MustOperateOutOfPlace determines whether a converter is capable of operating in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_MustOperateOutOfPlace(_:_:_:_:)
func VImageConverter_MustOperateOutOfPlace(converter unsafe.Pointer, srcs unsafe.Pointer, dests unsafe.Pointer, flags uint32) int {
	if _vImageConverter_MustOperateOutOfPlace == nil {
		panic("Accelerate: symbol vImageConverter_MustOperateOutOfPlace not loaded")
	}
	return _vImageConverter_MustOperateOutOfPlace(converter, srcs, dests, flags)
}


var _vImageConverter_Release func(converter unsafe.Pointer)

// VImageConverter_Release releases a vImage converter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_Release
func VImageConverter_Release(converter unsafe.Pointer) {
	if _vImageConverter_Release == nil {
		panic("Accelerate: symbol vImageConverter_Release not loaded")
	}
	_vImageConverter_Release(converter)
}


var _vImageConverter_Retain func(converter unsafe.Pointer)

// VImageConverter_Retain retains a vImage converter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter_Retain
func VImageConverter_Retain(converter unsafe.Pointer) {
	if _vImageConverter_Retain == nil {
		panic("Accelerate: symbol vImageConverter_Retain not loaded")
	}
	_vImageConverter_Retain(converter)
}


var _vImageConvolveFloatKernel_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernelHeight uint32, kernelWidth uint32, bias float32, backgroundColor Pixel_8888, flags uint32) int

// VImageConvolveFloatKernel_ARGB8888 convolves an 8-bit-per-channel, 4-channel interleaved image using 32-bit weights.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveFloatKernel_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveFloatKernel_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernelHeight uint32, kernelWidth uint32, bias float32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvolveFloatKernel_ARGB8888 == nil {
		panic("Accelerate: symbol vImageConvolveFloatKernel_ARGB8888 not loaded")
	}
	return _vImageConvolveFloatKernel_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernelHeight, kernelWidth, bias, backgroundColor, flags)
}


var _vImageConvolveMultiKernel_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernels *int16, kernel_height uint32, kernel_width uint32, divisors int32, biases int32, backgroundColor Pixel_8888, flags uint32) int

// VImageConvolveMultiKernel_ARGB8888 convolves each channel of an 8-bit-per-channel, 4-channel interleaved image by one of the four 2D kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveMultiKernel_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveMultiKernel_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernels *int16, kernel_height uint32, kernel_width uint32, divisors int32, biases int32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvolveMultiKernel_ARGB8888 == nil {
		panic("Accelerate: symbol vImageConvolveMultiKernel_ARGB8888 not loaded")
	}
	return _vImageConvolveMultiKernel_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernels, kernel_height, kernel_width, divisors, biases, backgroundColor, flags)
}


var _vImageConvolveMultiKernel_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernels uintptr, kernel_height uint32, kernel_width uint32, biases unsafe.Pointer, backgroundColor Pixel_FFFF, flags uint32) int

// VImageConvolveMultiKernel_ARGBFFFF convolves each channel of a floating-point 32-bit-per-channel, 4-channel interleaved image by one of the four 2D kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveMultiKernel_ARGBFFFF(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveMultiKernel_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernels uintptr, kernel_height uint32, kernel_width uint32, biases unsafe.Pointer, backgroundColor Pixel_FFFF, flags uint32) int {
	if _vImageConvolveMultiKernel_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvolveMultiKernel_ARGBFFFF not loaded")
	}
	return _vImageConvolveMultiKernel_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernels, kernel_height, kernel_width, biases, backgroundColor, flags)
}


var _vImageConvolveWithBias_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_ARGB_16F, flags uint32) int

// VImageConvolveWithBias_ARGB16F convolves a floating-point 16-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_ARGB16F(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageConvolveWithBias_ARGB16F == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_ARGB16F not loaded")
	}
	return _vImageConvolveWithBias_ARGB16F(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, bias, backgroundColor, flags)
}


var _vImageConvolveWithBias_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, bias int32, backgroundColor Pixel_8888, flags uint32) int

// VImageConvolveWithBias_ARGB8888 convolves an 8-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, bias int32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvolveWithBias_ARGB8888 == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_ARGB8888 not loaded")
	}
	return _vImageConvolveWithBias_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, bias, backgroundColor, flags)
}


var _vImageConvolveWithBias_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_FFFF, flags uint32) int

// VImageConvolveWithBias_ARGBFFFF convolves a floating-point 32-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_ARGBFFFF(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_FFFF, flags uint32) int {
	if _vImageConvolveWithBias_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_ARGBFFFF not loaded")
	}
	return _vImageConvolveWithBias_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, bias, backgroundColor, flags)
}


var _vImageConvolveWithBias_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_16F, flags uint32) int

// VImageConvolveWithBias_Planar16F convolves a floating-point 16-bit planar image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_Planar16F(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_16F, flags uint32) int {
	if _vImageConvolveWithBias_Planar16F == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_Planar16F not loaded")
	}
	return _vImageConvolveWithBias_Planar16F(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, bias, backgroundColor, flags)
}


var _vImageConvolveWithBias_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, bias int32, backgroundColor Pixel_8, flags uint32) int

// VImageConvolveWithBias_Planar8 convolves an 8-bit planar image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_Planar8(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, bias int32, backgroundColor Pixel_8, flags uint32) int {
	if _vImageConvolveWithBias_Planar8 == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_Planar8 not loaded")
	}
	return _vImageConvolveWithBias_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, bias, backgroundColor, flags)
}


var _vImageConvolveWithBias_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_F, flags uint32) int

// VImageConvolveWithBias_PlanarF convolves a floating-point 32-bit planar image by a 2D kernel and adds a bias.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_PlanarF(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolveWithBias_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, bias float32, backgroundColor Pixel_F, flags uint32) int {
	if _vImageConvolveWithBias_PlanarF == nil {
		panic("Accelerate: symbol vImageConvolveWithBias_PlanarF not loaded")
	}
	return _vImageConvolveWithBias_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, bias, backgroundColor, flags)
}


var _vImageConvolve_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_ARGB_16F, flags uint32) int

// VImageConvolve_ARGB16F convolves a floating-point 16-bit-per-channel, 4-channel interleaved image by a 2D kernel, then divides the pixel values by a divisor.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_ARGB16F(_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageConvolve_ARGB16F == nil {
		panic("Accelerate: symbol vImageConvolve_ARGB16F not loaded")
	}
	return _vImageConvolve_ARGB16F(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageConvolve_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, backgroundColor Pixel_8888, flags uint32) int

// VImageConvolve_ARGB8888 convolves an 8-bit-per-channel, 4-channel interleaved image by a 2D kernel and divides the pixel values by a divisor.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageConvolve_ARGB8888 == nil {
		panic("Accelerate: symbol vImageConvolve_ARGB8888 not loaded")
	}
	return _vImageConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, backgroundColor, flags)
}


var _vImageConvolve_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_FFFF, flags uint32) int

// VImageConvolve_ARGBFFFF convolves a floating-point 32-bit-per-channel, 4-channel interleaved image by a 2D kernel, then divides the pixel values by a divisor.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_ARGBFFFF(_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_FFFF, flags uint32) int {
	if _vImageConvolve_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageConvolve_ARGBFFFF not loaded")
	}
	return _vImageConvolve_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageConvolve_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_16F, flags uint32) int

// VImageConvolve_Planar16F convolves a floating-point 16-bit planar image by a 2D kernel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_Planar16F(_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_16F, flags uint32) int {
	if _vImageConvolve_Planar16F == nil {
		panic("Accelerate: symbol vImageConvolve_Planar16F not loaded")
	}
	return _vImageConvolve_Planar16F(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageConvolve_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, backgroundColor Pixel_8, flags uint32) int

// VImageConvolve_Planar8 convolves an 8-bit planar image by a 2D kernel and divides the pixel values by a divisor.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_Planar8(_:_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel_height uint32, kernel_width uint32, divisor int32, backgroundColor Pixel_8, flags uint32) int {
	if _vImageConvolve_Planar8 == nil {
		panic("Accelerate: symbol vImageConvolve_Planar8 not loaded")
	}
	return _vImageConvolve_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, backgroundColor, flags)
}


var _vImageConvolve_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_F, flags uint32) int

// VImageConvolve_PlanarF convolves a floating-point 32-bit planar image by a 2D kernel.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConvolve_PlanarF(_:_:_:_:_:_:_:_:_:_:)
func VImageConvolve_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_F, flags uint32) int {
	if _vImageConvolve_PlanarF == nil {
		panic("Accelerate: symbol vImageConvolve_PlanarF not loaded")
	}
	return _vImageConvolve_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageCopyBuffer func(src unsafe.Pointer, dest unsafe.Pointer, pixelSize uintptr, flags uint32) int

// VImageCopyBuffer copies the contents of a vImage buffer to a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCopyBuffer(_:_:_:_:)
func VImageCopyBuffer(src unsafe.Pointer, dest unsafe.Pointer, pixelSize uintptr, flags uint32) int {
	if _vImageCopyBuffer == nil {
		panic("Accelerate: symbol vImageCopyBuffer not loaded")
	}
	return _vImageCopyBuffer(src, dest, pixelSize, flags)
}


var _vImageCreateCGImageFromBuffer func(buf unsafe.Pointer, format *VImage_CGImageFormat, userData unsafe.Pointer, flags uint32, err *int) coregraphics.CGImageRef

// VImageCreateCGImageFromBuffer creates a Core Graphics image from a vImage buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCreateCGImageFromBuffer(_:_:_:_:_:_:)
func VImageCreateCGImageFromBuffer(buf unsafe.Pointer, format *VImage_CGImageFormat, userData unsafe.Pointer, flags uint32, err *int) coregraphics.CGImageRef {
	if _vImageCreateCGImageFromBuffer == nil {
		panic("Accelerate: symbol vImageCreateCGImageFromBuffer not loaded")
	}
	return _vImageCreateCGImageFromBuffer(buf, format, userData, flags, err)
}


var _vImageCreateGammaFunction func(gamma float32, gamma_type int, flags uint32) GammaFunction

// VImageCreateGammaFunction returns a gamma function object.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCreateGammaFunction(_:_:_:)
func VImageCreateGammaFunction(gamma float32, gamma_type int, flags uint32) GammaFunction {
	if _vImageCreateGammaFunction == nil {
		panic("Accelerate: symbol vImageCreateGammaFunction not loaded")
	}
	return _vImageCreateGammaFunction(gamma, gamma_type, flags)
}


var _vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction func(whitePoint *VImageWhitePoint, tf *VImageTransferFunction, intent coregraphics.CGColorRenderingIntent, flags uint32, err *int) coregraphics.CGColorSpaceRef

// VImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction creates a monochrome color space based on primitives from Y’CbCr specifications.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction(_:_:_:_:_:)
func VImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction(whitePoint *VImageWhitePoint, tf *VImageTransferFunction, intent coregraphics.CGColorRenderingIntent, flags uint32, err *int) coregraphics.CGColorSpaceRef {
	if _vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction == nil {
		panic("Accelerate: symbol vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction not loaded")
	}
	return _vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction(whitePoint, tf, intent, flags, err)
}


var _vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction func(primaries *VImageRGBPrimaries, tf *VImageTransferFunction, intent coregraphics.CGColorRenderingIntent, flags uint32, err *int) coregraphics.CGColorSpaceRef

// VImageCreateRGBColorSpaceWithPrimariesAndTransferFunction creates an RGB color space based on primitives from Y’CbCr specifications.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction(_:_:_:_:_:)
func VImageCreateRGBColorSpaceWithPrimariesAndTransferFunction(primaries *VImageRGBPrimaries, tf *VImageTransferFunction, intent coregraphics.CGColorRenderingIntent, flags uint32, err *int) coregraphics.CGColorSpaceRef {
	if _vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction == nil {
		panic("Accelerate: symbol vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction not loaded")
	}
	return _vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction(primaries, tf, intent, flags, err)
}


var _vImageDestroyGammaFunction func(f GammaFunction)

// VImageDestroyGammaFunction destroys a gamma function object.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDestroyGammaFunction(_:)
func VImageDestroyGammaFunction(f GammaFunction) {
	if _vImageDestroyGammaFunction == nil {
		panic("Accelerate: symbol vImageDestroyGammaFunction not loaded")
	}
	_vImageDestroyGammaFunction(f)
}


var _vImageDestroyResamplingFilter func(filter ResamplingFilter)

// VImageDestroyResamplingFilter disposes of a resampling filter object.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDestroyResamplingFilter(_:)
func VImageDestroyResamplingFilter(filter ResamplingFilter) {
	if _vImageDestroyResamplingFilter == nil {
		panic("Accelerate: symbol vImageDestroyResamplingFilter not loaded")
	}
	_vImageDestroyResamplingFilter(filter)
}


var _vImageDilate_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int

// VImageDilate_ARGB8888 dilates an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDilate_ARGB8888(_:_:_:_:_:_:_:_:)
func VImageDilate_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageDilate_ARGB8888 == nil {
		panic("Accelerate: symbol vImageDilate_ARGB8888 not loaded")
	}
	return _vImageDilate_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
}


var _vImageDilate_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint, kernel_width uint, flags uint32) int

// VImageDilate_ARGBFFFF dilates a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDilate_ARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageDilate_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageDilate_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageDilate_ARGBFFFF not loaded")
	}
	return _vImageDilate_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, flags)
}


var _vImageDilate_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int

// VImageDilate_Planar8 dilates an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDilate_Planar8(_:_:_:_:_:_:_:_:)
func VImageDilate_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageDilate_Planar8 == nil {
		panic("Accelerate: symbol vImageDilate_Planar8 not loaded")
	}
	return _vImageDilate_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
}


var _vImageDilate_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint, kernel_width uint, flags uint32) int

// VImageDilate_PlanarF dilates a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageDilate_PlanarF(_:_:_:_:_:_:_:_:)
func VImageDilate_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageDilate_PlanarF == nil {
		panic("Accelerate: symbol vImageDilate_PlanarF not loaded")
	}
	return _vImageDilate_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, flags)
}


var _vImageEndsInContrastStretch_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, percent_low unsafe.Pointer, percent_high unsafe.Pointer, flags uint32) int

// VImageEndsInContrastStretch_ARGB8888 performs ends-in contrast stretching on an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEndsInContrastStretch_ARGB8888(_:_:_:_:_:)
func VImageEndsInContrastStretch_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, percent_low unsafe.Pointer, percent_high unsafe.Pointer, flags uint32) int {
	if _vImageEndsInContrastStretch_ARGB8888 == nil {
		panic("Accelerate: symbol vImageEndsInContrastStretch_ARGB8888 not loaded")
	}
	return _vImageEndsInContrastStretch_ARGB8888(src, dest, percent_low, percent_high, flags)
}


var _vImageEndsInContrastStretch_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, percent_low unsafe.Pointer, percent_high unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageEndsInContrastStretch_ARGBFFFF performs ends-in contrast stretching on a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEndsInContrastStretch_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func VImageEndsInContrastStretch_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, percent_low unsafe.Pointer, percent_high unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageEndsInContrastStretch_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageEndsInContrastStretch_ARGBFFFF not loaded")
	}
	return _vImageEndsInContrastStretch_ARGBFFFF(src, dest, tempBuffer, percent_low, percent_high, histogram_entries, minVal, maxVal, flags)
}


var _vImageEndsInContrastStretch_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, percent_low uint, percent_high uint, flags uint32) int

// VImageEndsInContrastStretch_Planar8 performs ends-in contrast stretching on an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEndsInContrastStretch_Planar8(_:_:_:_:_:)
func VImageEndsInContrastStretch_Planar8(src unsafe.Pointer, dest unsafe.Pointer, percent_low uint, percent_high uint, flags uint32) int {
	if _vImageEndsInContrastStretch_Planar8 == nil {
		panic("Accelerate: symbol vImageEndsInContrastStretch_Planar8 not loaded")
	}
	return _vImageEndsInContrastStretch_Planar8(src, dest, percent_low, percent_high, flags)
}


var _vImageEndsInContrastStretch_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, percent_low uint, percent_high uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageEndsInContrastStretch_PlanarF performs ends-in contrast stretching on a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEndsInContrastStretch_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImageEndsInContrastStretch_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, percent_low uint, percent_high uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageEndsInContrastStretch_PlanarF == nil {
		panic("Accelerate: symbol vImageEndsInContrastStretch_PlanarF not loaded")
	}
	return _vImageEndsInContrastStretch_PlanarF(src, dest, tempBuffer, percent_low, percent_high, histogram_entries, minVal, maxVal, flags)
}


var _vImageEqualization_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageEqualization_ARGB8888 performs histogram equalization on an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEqualization_ARGB8888(_:_:_:)
func VImageEqualization_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageEqualization_ARGB8888 == nil {
		panic("Accelerate: symbol vImageEqualization_ARGB8888 not loaded")
	}
	return _vImageEqualization_ARGB8888(src, dest, flags)
}


var _vImageEqualization_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageEqualization_ARGBFFFF performs histogram equalization on a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEqualization_ARGBFFFF(_:_:_:_:_:_:_:)
func VImageEqualization_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageEqualization_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageEqualization_ARGBFFFF not loaded")
	}
	return _vImageEqualization_ARGBFFFF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
}


var _vImageEqualization_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageEqualization_Planar8 performs histogram equalization on an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEqualization_Planar8(_:_:_:)
func VImageEqualization_Planar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageEqualization_Planar8 == nil {
		panic("Accelerate: symbol vImageEqualization_Planar8 not loaded")
	}
	return _vImageEqualization_Planar8(src, dest, flags)
}


var _vImageEqualization_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageEqualization_PlanarF performs histogram equalization on a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageEqualization_PlanarF(_:_:_:_:_:_:_:)
func VImageEqualization_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageEqualization_PlanarF == nil {
		panic("Accelerate: symbol vImageEqualization_PlanarF not loaded")
	}
	return _vImageEqualization_PlanarF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
}


var _vImageErode_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int

// VImageErode_ARGB8888 erodes an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageErode_ARGB8888(_:_:_:_:_:_:_:_:)
func VImageErode_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageErode_ARGB8888 == nil {
		panic("Accelerate: symbol vImageErode_ARGB8888 not loaded")
	}
	return _vImageErode_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
}


var _vImageErode_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint, kernel_width uint, flags uint32) int

// VImageErode_ARGBFFFF erodes a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageErode_ARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageErode_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageErode_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageErode_ARGBFFFF not loaded")
	}
	return _vImageErode_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, flags)
}


var _vImageErode_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int

// VImageErode_Planar8 erodes an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageErode_Planar8(_:_:_:_:_:_:_:_:)
func VImageErode_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *byte, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageErode_Planar8 == nil {
		panic("Accelerate: symbol vImageErode_Planar8 not loaded")
	}
	return _vImageErode_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
}


var _vImageErode_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel_height uint, kernel_width uint, flags uint32) int

// VImageErode_PlanarF erodes a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageErode_PlanarF(_:_:_:_:_:_:_:_:)
func VImageErode_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageErode_PlanarF == nil {
		panic("Accelerate: symbol vImageErode_PlanarF not loaded")
	}
	return _vImageErode_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), kernel_height, kernel_width, flags)
}


var _vImageExtractChannel_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int

// VImageExtractChannel_ARGB16U extracts a single channel from an unsigned 16-bit-per-channel, 4-channel interleaved buffer and writes the result to an unsigned 16-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageExtractChannel_ARGB16U(_:_:_:_:)
func VImageExtractChannel_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int {
	if _vImageExtractChannel_ARGB16U == nil {
		panic("Accelerate: symbol vImageExtractChannel_ARGB16U not loaded")
	}
	return _vImageExtractChannel_ARGB16U(src, dest, channelIndex, flags)
}


var _vImageExtractChannel_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int

// VImageExtractChannel_ARGB8888 extracts a single channel from an 8-bit-per-channel, 4-channel interleaved buffer and writes the result to an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageExtractChannel_ARGB8888(_:_:_:_:)
func VImageExtractChannel_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int {
	if _vImageExtractChannel_ARGB8888 == nil {
		panic("Accelerate: symbol vImageExtractChannel_ARGB8888 not loaded")
	}
	return _vImageExtractChannel_ARGB8888(src, dest, channelIndex, flags)
}


var _vImageExtractChannel_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int

// VImageExtractChannel_ARGBFFFF extracts a single channel from a 32-bit-per-channel, 4-channel interleaved buffer and writes the result to a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageExtractChannel_ARGBFFFF(_:_:_:_:)
func VImageExtractChannel_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, channelIndex int, flags uint32) int {
	if _vImageExtractChannel_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageExtractChannel_ARGBFFFF not loaded")
	}
	return _vImageExtractChannel_ARGBFFFF(src, dest, channelIndex, flags)
}


var _vImageFlatten_ARGB16Q12 func(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16S, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_ARGB16Q12 performs an alpha composite of a fixed-point 16-bit-per-channel, 4-channel ARGB buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGB16Q12(_:_:_:_:_:)
func VImageFlatten_ARGB16Q12(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16S, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_ARGB16Q12 == nil {
		panic("Accelerate: symbol vImageFlatten_ARGB16Q12 not loaded")
	}
	return _vImageFlatten_ARGB16Q12(argbSrc, argbDst, argbBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_ARGB16U func(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16U, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_ARGB16U performs an alpha composite of an unsigned 16-bit-per-channel, 4-channel ARGB buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGB16U(_:_:_:_:_:)
func VImageFlatten_ARGB16U(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16U, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_ARGB16U == nil {
		panic("Accelerate: symbol vImageFlatten_ARGB16U not loaded")
	}
	return _vImageFlatten_ARGB16U(argbSrc, argbDst, argbBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_ARGB8888 func(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_8888, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_ARGB8888 performs an alpha composite of an 8-bit-per-channel, 4-channel ARGB buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGB8888(_:_:_:_:_:)
func VImageFlatten_ARGB8888(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_8888, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_ARGB8888 == nil {
		panic("Accelerate: symbol vImageFlatten_ARGB8888 not loaded")
	}
	return _vImageFlatten_ARGB8888(argbSrc, argbDst, argbBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_ARGB8888ToRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int

// VImageFlatten_ARGB8888ToRGB888 flattens an 8-bit-per-channel ARGB buffer against a solid background to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGB8888ToRGB888(_:_:_:_:_:)
func VImageFlatten_ARGB8888ToRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_ARGB8888ToRGB888 == nil {
		panic("Accelerate: symbol vImageFlatten_ARGB8888ToRGB888 not loaded")
	}
	return _vImageFlatten_ARGB8888ToRGB888(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFlatten_ARGBFFFF func(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_FFFF, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_ARGBFFFF performs an alpha composite of a 32-bit-per-channel, 4-channel ARGB buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGBFFFF(_:_:_:_:_:)
func VImageFlatten_ARGBFFFF(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_FFFF, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageFlatten_ARGBFFFF not loaded")
	}
	return _vImageFlatten_ARGBFFFF(argbSrc, argbDst, argbBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_ARGBFFFFToRGBFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int

// VImageFlatten_ARGBFFFFToRGBFFF flattens a 32-bit-per-channel ARGB buffer against a solid background to produce a 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_ARGBFFFFToRGBFFF(_:_:_:_:_:)
func VImageFlatten_ARGBFFFFToRGBFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_ARGBFFFFToRGBFFF == nil {
		panic("Accelerate: symbol vImageFlatten_ARGBFFFFToRGBFFF not loaded")
	}
	return _vImageFlatten_ARGBFFFFToRGBFFF(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFlatten_BGRA8888ToRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int

// VImageFlatten_BGRA8888ToRGB888 flattens an 8-bit-per-channel BGRA buffer against a solid background to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_BGRA8888ToRGB888(_:_:_:_:_:)
func VImageFlatten_BGRA8888ToRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_BGRA8888ToRGB888 == nil {
		panic("Accelerate: symbol vImageFlatten_BGRA8888ToRGB888 not loaded")
	}
	return _vImageFlatten_BGRA8888ToRGB888(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFlatten_BGRAFFFFToRGBFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int

// VImageFlatten_BGRAFFFFToRGBFFF flattens a 32-bit-per-channel BGRA buffer against a solid background to produce a 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_BGRAFFFFToRGBFFF(_:_:_:_:_:)
func VImageFlatten_BGRAFFFFToRGBFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_BGRAFFFFToRGBFFF == nil {
		panic("Accelerate: symbol vImageFlatten_BGRAFFFFToRGBFFF not loaded")
	}
	return _vImageFlatten_BGRAFFFFToRGBFFF(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFlatten_RGBA16Q12 func(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16S, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_RGBA16Q12 performs an alpha composite of a fixed-point 16-bit-per-channel, 4-channel RGBA buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBA16Q12(_:_:_:_:_:)
func VImageFlatten_RGBA16Q12(argbSrc unsafe.Pointer, argbDst unsafe.Pointer, argbBackgroundColorPtr Pixel_ARGB_16S, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_RGBA16Q12 == nil {
		panic("Accelerate: symbol vImageFlatten_RGBA16Q12 not loaded")
	}
	return _vImageFlatten_RGBA16Q12(argbSrc, argbDst, argbBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_RGBA16U func(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_ARGB_16U, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_RGBA16U performs an alpha composite of an unsigned 16-bit-per-channel, 4-channel RGBA buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBA16U(_:_:_:_:_:)
func VImageFlatten_RGBA16U(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_ARGB_16U, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_RGBA16U == nil {
		panic("Accelerate: symbol vImageFlatten_RGBA16U not loaded")
	}
	return _vImageFlatten_RGBA16U(rgbaSrc, rgbaDst, rgbaBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_RGBA8888 func(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_8888, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_RGBA8888 performs an alpha composite of an 8-bit-per-channel, 4-channel RGBA buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBA8888(_:_:_:_:_:)
func VImageFlatten_RGBA8888(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_8888, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_RGBA8888 == nil {
		panic("Accelerate: symbol vImageFlatten_RGBA8888 not loaded")
	}
	return _vImageFlatten_RGBA8888(rgbaSrc, rgbaDst, rgbaBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_RGBA8888ToRGB888 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int

// VImageFlatten_RGBA8888ToRGB888 flattens an 8-bit-per-channel RGBA buffer against a solid background to produce an 8-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBA8888ToRGB888(_:_:_:_:_:)
func VImageFlatten_RGBA8888ToRGB888(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_8888, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_RGBA8888ToRGB888 == nil {
		panic("Accelerate: symbol vImageFlatten_RGBA8888ToRGB888 not loaded")
	}
	return _vImageFlatten_RGBA8888ToRGB888(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFlatten_RGBAFFFF func(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_FFFF, isImagePremultiplied bool, flags uint32) int

// VImageFlatten_RGBAFFFF performs an alpha composite of a 32-bit-per-channel, 4-channel RGBA buffer over a solid background color.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBAFFFF(_:_:_:_:_:)
func VImageFlatten_RGBAFFFF(rgbaSrc unsafe.Pointer, rgbaDst unsafe.Pointer, rgbaBackgroundColorPtr Pixel_FFFF, isImagePremultiplied bool, flags uint32) int {
	if _vImageFlatten_RGBAFFFF == nil {
		panic("Accelerate: symbol vImageFlatten_RGBAFFFF not loaded")
	}
	return _vImageFlatten_RGBAFFFF(rgbaSrc, rgbaDst, rgbaBackgroundColorPtr, isImagePremultiplied, flags)
}


var _vImageFlatten_RGBAFFFFToRGBFFF func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int

// VImageFlatten_RGBAFFFFToRGBFFF flattens a 32-bit-per-channel RGBA buffer against a solid background to produce a 32-bit-per-channel RGB result.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFlatten_RGBAFFFFToRGBFFF(_:_:_:_:_:)
func VImageFlatten_RGBAFFFFToRGBFFF(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 Pixel_FFFF, arg3 bool, arg4 uint32) int {
	if _vImageFlatten_RGBAFFFFToRGBFFF == nil {
		panic("Accelerate: symbol vImageFlatten_RGBAFFFFToRGBFFF not loaded")
	}
	return _vImageFlatten_RGBAFFFFToRGBFFF(arg0, arg1, arg2, arg3, arg4)
}


var _vImageFloodFill_ARGB16U func(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_ARGB_16U, connectivity int, flags uint32) int

// VImageFloodFill_ARGB16U applies a flood-fill operation to an unsigned 16-bit-per-channel, four-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFloodFill_ARGB16U(_:_:_:_:_:_:_:)
func VImageFloodFill_ARGB16U(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_ARGB_16U, connectivity int, flags uint32) int {
	if _vImageFloodFill_ARGB16U == nil {
		panic("Accelerate: symbol vImageFloodFill_ARGB16U not loaded")
	}
	return _vImageFloodFill_ARGB16U(srcDest, tempBuffer, seedX, seedY, newValue, connectivity, flags)
}


var _vImageFloodFill_ARGB8888 func(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_8888, connectivity int, flags uint32) int

// VImageFloodFill_ARGB8888 applies a flood-fill operation to an 8-bit-per-channel, four-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFloodFill_ARGB8888(_:_:_:_:_:_:_:)
func VImageFloodFill_ARGB8888(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_8888, connectivity int, flags uint32) int {
	if _vImageFloodFill_ARGB8888 == nil {
		panic("Accelerate: symbol vImageFloodFill_ARGB8888 not loaded")
	}
	return _vImageFloodFill_ARGB8888(srcDest, tempBuffer, seedX, seedY, newValue, connectivity, flags)
}


var _vImageFloodFill_Planar16U func(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_16U, connectivity int, flags uint32) int

// VImageFloodFill_Planar16U applies a flood fill-operation to an unsigned 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFloodFill_Planar16U(_:_:_:_:_:_:_:)
func VImageFloodFill_Planar16U(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_16U, connectivity int, flags uint32) int {
	if _vImageFloodFill_Planar16U == nil {
		panic("Accelerate: symbol vImageFloodFill_Planar16U not loaded")
	}
	return _vImageFloodFill_Planar16U(srcDest, tempBuffer, seedX, seedY, newValue, connectivity, flags)
}


var _vImageFloodFill_Planar8 func(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_8, connectivity int, flags uint32) int

// VImageFloodFill_Planar8 applies a flood-fill operation to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageFloodFill_Planar8(_:_:_:_:_:_:_:)
func VImageFloodFill_Planar8(srcDest unsafe.Pointer, tempBuffer unsafe.Pointer, seedX uint, seedY uint, newValue Pixel_8, connectivity int, flags uint32) int {
	if _vImageFloodFill_Planar8 == nil {
		panic("Accelerate: symbol vImageFloodFill_Planar8 not loaded")
	}
	return _vImageFloodFill_Planar8(srcDest, tempBuffer, seedX, seedY, newValue, connectivity, flags)
}


var _vImageGamma_Planar8toPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int

// VImageGamma_Planar8toPlanarF applies a gamma function to an 8-bit planar image to produce a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGamma_Planar8toPlanarF(_:_:_:_:)
func VImageGamma_Planar8toPlanarF(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int {
	if _vImageGamma_Planar8toPlanarF == nil {
		panic("Accelerate: symbol vImageGamma_Planar8toPlanarF not loaded")
	}
	return _vImageGamma_Planar8toPlanarF(src, dest, gamma, flags)
}


var _vImageGamma_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int

// VImageGamma_PlanarF applies a gamma function to a PlanarF image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGamma_PlanarF(_:_:_:_:)
func VImageGamma_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int {
	if _vImageGamma_PlanarF == nil {
		panic("Accelerate: symbol vImageGamma_PlanarF not loaded")
	}
	return _vImageGamma_PlanarF(src, dest, gamma, flags)
}


var _vImageGamma_PlanarFtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int

// VImageGamma_PlanarFtoPlanar8 applies a gamma function to a 32-bit planar image to produce an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGamma_PlanarFtoPlanar8(_:_:_:_:)
func VImageGamma_PlanarFtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, gamma GammaFunction, flags uint32) int {
	if _vImageGamma_PlanarFtoPlanar8 == nil {
		panic("Accelerate: symbol vImageGamma_PlanarFtoPlanar8 not loaded")
	}
	return _vImageGamma_PlanarFtoPlanar8(src, dest, gamma, flags)
}


var _vImageGetPerspectiveWarp func(srcPoints unsafe.Pointer, destPoints unsafe.Pointer, transform *VImage_PerpsectiveTransform, flags uint32) int

// VImageGetPerspectiveWarp returns a projective-transformation structure that defines the mapping between a source quadrilateral and a destination quadrilateral.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGetPerspectiveWarp(_:_:_:_:)
func VImageGetPerspectiveWarp(srcPoints unsafe.Pointer, destPoints unsafe.Pointer, transform *VImage_PerpsectiveTransform, flags uint32) int {
	if _vImageGetPerspectiveWarp == nil {
		panic("Accelerate: symbol vImageGetPerspectiveWarp not loaded")
	}
	return _vImageGetPerspectiveWarp(srcPoints, destPoints, transform, flags)
}


var _vImageGetResamplingFilterExtent func(filter ResamplingFilter, flags uint32) uint

// VImageGetResamplingFilterExtent returns the maximum sampling radius for a resampling filter.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGetResamplingFilterExtent(_:_:)
func VImageGetResamplingFilterExtent(filter ResamplingFilter, flags uint32) uint {
	if _vImageGetResamplingFilterExtent == nil {
		panic("Accelerate: symbol vImageGetResamplingFilterExtent not loaded")
	}
	return _vImageGetResamplingFilterExtent(filter, flags)
}


var _vImageGetResamplingFilterSize func(scale float32, kernelWidth float32, flags uint32) uintptr

// VImageGetResamplingFilterSize returns the minimum size, in bytes, for the buffer needed by the new resampling filter function.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageGetResamplingFilterSize(_:_:_:_:)
func VImageGetResamplingFilterSize(scale float32, kernelWidth float32, flags uint32) uintptr {
	if _vImageGetResamplingFilterSize == nil {
		panic("Accelerate: symbol vImageGetResamplingFilterSize not loaded")
	}
	return _vImageGetResamplingFilterSize(scale, kernelWidth, flags)
}


var _vImageHistogramCalculation_ARGB8888 func(src unsafe.Pointer, histogram *uint, flags uint32) int

// VImageHistogramCalculation_ARGB8888 calculates the histogram of an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramCalculation_ARGB8888(_:_:_:)
func VImageHistogramCalculation_ARGB8888(src unsafe.Pointer, histogram *uint, flags uint32) int {
	if _vImageHistogramCalculation_ARGB8888 == nil {
		panic("Accelerate: symbol vImageHistogramCalculation_ARGB8888 not loaded")
	}
	return _vImageHistogramCalculation_ARGB8888(src, histogram, flags)
}


var _vImageHistogramCalculation_ARGBFFFF func(src unsafe.Pointer, histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageHistogramCalculation_ARGBFFFF calculates the histogram of a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramCalculation_ARGBFFFF(_:_:_:_:_:_:)
func VImageHistogramCalculation_ARGBFFFF(src unsafe.Pointer, histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageHistogramCalculation_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageHistogramCalculation_ARGBFFFF not loaded")
	}
	return _vImageHistogramCalculation_ARGBFFFF(src, histogram, histogram_entries, minVal, maxVal, flags)
}


var _vImageHistogramCalculation_Planar8 func(src unsafe.Pointer, histogram *uint, flags uint32) int

// VImageHistogramCalculation_Planar8 calculates the histogram of an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramCalculation_Planar8(_:_:_:)
func VImageHistogramCalculation_Planar8(src unsafe.Pointer, histogram *uint, flags uint32) int {
	if _vImageHistogramCalculation_Planar8 == nil {
		panic("Accelerate: symbol vImageHistogramCalculation_Planar8 not loaded")
	}
	return _vImageHistogramCalculation_Planar8(src, histogram, flags)
}


var _vImageHistogramCalculation_PlanarF func(src unsafe.Pointer, histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageHistogramCalculation_PlanarF calculates the histogram of a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramCalculation_PlanarF(_:_:_:_:_:_:)
func VImageHistogramCalculation_PlanarF(src unsafe.Pointer, histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageHistogramCalculation_PlanarF == nil {
		panic("Accelerate: symbol vImageHistogramCalculation_PlanarF not loaded")
	}
	return _vImageHistogramCalculation_PlanarF(src, histogram, histogram_entries, minVal, maxVal, flags)
}


var _vImageHistogramSpecification_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, desired_histogram *uint, flags uint32) int

// VImageHistogramSpecification_ARGB8888 specifies the histogram of an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramSpecification_ARGB8888(_:_:_:_:)
func VImageHistogramSpecification_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, desired_histogram *uint, flags uint32) int {
	if _vImageHistogramSpecification_ARGB8888 == nil {
		panic("Accelerate: symbol vImageHistogramSpecification_ARGB8888 not loaded")
	}
	return _vImageHistogramSpecification_ARGB8888(src, dest, desired_histogram, flags)
}


var _vImageHistogramSpecification_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, desired_histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageHistogramSpecification_ARGBFFFF specifes the histogram of a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramSpecification_ARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageHistogramSpecification_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, desired_histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageHistogramSpecification_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageHistogramSpecification_ARGBFFFF not loaded")
	}
	return _vImageHistogramSpecification_ARGBFFFF(src, dest, tempBuffer, desired_histogram, histogram_entries, minVal, maxVal, flags)
}


var _vImageHistogramSpecification_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, desired_histogram *uint, flags uint32) int

// VImageHistogramSpecification_Planar8 specifies the histogram of an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramSpecification_Planar8(_:_:_:_:)
func VImageHistogramSpecification_Planar8(src unsafe.Pointer, dest unsafe.Pointer, desired_histogram *uint, flags uint32) int {
	if _vImageHistogramSpecification_Planar8 == nil {
		panic("Accelerate: symbol vImageHistogramSpecification_Planar8 not loaded")
	}
	return _vImageHistogramSpecification_Planar8(src, dest, desired_histogram, flags)
}


var _vImageHistogramSpecification_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, desired_histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int

// VImageHistogramSpecification_PlanarF specifies the histogram of a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHistogramSpecification_PlanarF(_:_:_:_:_:_:_:_:)
func VImageHistogramSpecification_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, desired_histogram *uint, histogram_entries uint, minVal Pixel_F, maxVal Pixel_F, flags uint32) int {
	if _vImageHistogramSpecification_PlanarF == nil {
		panic("Accelerate: symbol vImageHistogramSpecification_PlanarF not loaded")
	}
	return _vImageHistogramSpecification_PlanarF(src, dest, tempBuffer, desired_histogram, histogram_entries, minVal, maxVal, flags)
}


var _vImageHorizontalReflect_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_ARGB16F reflects a floating-point 16-bit-per-channel, 4-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_ARGB16F(_:_:_:)
func VImageHorizontalReflect_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_ARGB16F == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_ARGB16F not loaded")
	}
	return _vImageHorizontalReflect_ARGB16F(src, dest, flags)
}


var _vImageHorizontalReflect_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_ARGB16S reflects a signed 16-bit-per-channel, 4-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_ARGB16S(_:_:_:)
func VImageHorizontalReflect_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_ARGB16S == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_ARGB16S not loaded")
	}
	return _vImageHorizontalReflect_ARGB16S(src, dest, flags)
}


var _vImageHorizontalReflect_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_ARGB16U reflects an unsigned 16-bit-per-channel, 4-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_ARGB16U(_:_:_:)
func VImageHorizontalReflect_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_ARGB16U == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_ARGB16U not loaded")
	}
	return _vImageHorizontalReflect_ARGB16U(src, dest, flags)
}


var _vImageHorizontalReflect_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_ARGB8888 reflects an 8-bit-per-channel, 4-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_ARGB8888(_:_:_:)
func VImageHorizontalReflect_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_ARGB8888 == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_ARGB8888 not loaded")
	}
	return _vImageHorizontalReflect_ARGB8888(src, dest, flags)
}


var _vImageHorizontalReflect_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_ARGBFFFF reflects a 32-bit-per-channel, 4-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_ARGBFFFF(_:_:_:)
func VImageHorizontalReflect_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_ARGBFFFF not loaded")
	}
	return _vImageHorizontalReflect_ARGBFFFF(src, dest, flags)
}


var _vImageHorizontalReflect_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_CbCr16F reflects a floating-point 16-bit-per-channel, 2-channel interleaved image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_CbCr16F(_:_:_:)
func VImageHorizontalReflect_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_CbCr16F == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_CbCr16F not loaded")
	}
	return _vImageHorizontalReflect_CbCr16F(src, dest, flags)
}


var _vImageHorizontalReflect_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_Planar16F reflects a floating-point 16-bit planar image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_Planar16F(_:_:_:)
func VImageHorizontalReflect_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_Planar16F == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_Planar16F not loaded")
	}
	return _vImageHorizontalReflect_Planar16F(src, dest, flags)
}


var _vImageHorizontalReflect_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_Planar16U reflects an unsigned 16-bit planar image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_Planar16U(_:_:_:)
func VImageHorizontalReflect_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_Planar16U == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_Planar16U not loaded")
	}
	return _vImageHorizontalReflect_Planar16U(src, dest, flags)
}


var _vImageHorizontalReflect_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_Planar8 reflects an 8-bit planar image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_Planar8(_:_:_:)
func VImageHorizontalReflect_Planar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_Planar8 == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_Planar8 not loaded")
	}
	return _vImageHorizontalReflect_Planar8(src, dest, flags)
}


var _vImageHorizontalReflect_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageHorizontalReflect_PlanarF reflects a 32-bit planar image horizontally across the center vertical line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalReflect_PlanarF(_:_:_:)
func VImageHorizontalReflect_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageHorizontalReflect_PlanarF == nil {
		panic("Accelerate: symbol vImageHorizontalReflect_PlanarF not loaded")
	}
	return _vImageHorizontalReflect_PlanarF(src, dest, flags)
}


var _vImageHorizontalShearD_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int

// VImageHorizontalShearD_ARGB16F performs a double-precision horizontal shear on a region of interest within a floating-point 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_ARGB16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageHorizontalShearD_ARGB16F == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_ARGB16F not loaded")
	}
	return _vImageHorizontalShearD_ARGB16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int

// VImageHorizontalShearD_ARGB16S performs a double-precision horizontal shear on a region of interest within a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_ARGB16S(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageHorizontalShearD_ARGB16S == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_ARGB16S not loaded")
	}
	return _vImageHorizontalShearD_ARGB16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int

// VImageHorizontalShearD_ARGB16U performs a double-precision horizontal shear on a region of interest within an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_ARGB16U(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageHorizontalShearD_ARGB16U == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_ARGB16U not loaded")
	}
	return _vImageHorizontalShearD_ARGB16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int

// VImageHorizontalShearD_ARGB8888 performs a double-precision horizontal shear on a region of interest within an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int {
	if _vImageHorizontalShearD_ARGB8888 == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_ARGB8888 not loaded")
	}
	return _vImageHorizontalShearD_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int

// VImageHorizontalShearD_ARGBFFFF performs a double-precision horizontal shear on a region of interest within a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int {
	if _vImageHorizontalShearD_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_ARGBFFFF not loaded")
	}
	return _vImageHorizontalShearD_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int

// VImageHorizontalShearD_CbCr16F performs a double-precision horizontal shear on a region of interest within a floating-point 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_CbCr16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int {
	if _vImageHorizontalShearD_CbCr16F == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_CbCr16F not loaded")
	}
	return _vImageHorizontalShearD_CbCr16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_CbCr16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int

// VImageHorizontalShearD_CbCr16S.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_CbCr16S(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_CbCr16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int {
	if _vImageHorizontalShearD_CbCr16S == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_CbCr16S not loaded")
	}
	return _vImageHorizontalShearD_CbCr16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_CbCr16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int

// VImageHorizontalShearD_CbCr16U.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_CbCr16U(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_CbCr16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int {
	if _vImageHorizontalShearD_CbCr16U == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_CbCr16U not loaded")
	}
	return _vImageHorizontalShearD_CbCr16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int

// VImageHorizontalShearD_Planar16F performs a double-precision horizontal shear on a region of interest within a floating-point 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_Planar16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int {
	if _vImageHorizontalShearD_Planar16F == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_Planar16F not loaded")
	}
	return _vImageHorizontalShearD_Planar16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8, flags uint32) int

// VImageHorizontalShearD_Planar8 performs a double-precision horizontal shear on a region of interest within an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8, flags uint32) int {
	if _vImageHorizontalShearD_Planar8 == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_Planar8 not loaded")
	}
	return _vImageHorizontalShearD_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShearD_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_F, flags uint32) int

// VImageHorizontalShearD_PlanarF performs a double-precision horizontal shear on a region of interest within a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShearD_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShearD_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_F, flags uint32) int {
	if _vImageHorizontalShearD_PlanarF == nil {
		panic("Accelerate: symbol vImageHorizontalShearD_PlanarF not loaded")
	}
	return _vImageHorizontalShearD_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int

// VImageHorizontalShear_ARGB16F performs a single-precision horizontal shear on a region of interest within a floating-point 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageHorizontalShear_ARGB16F == nil {
		panic("Accelerate: symbol vImageHorizontalShear_ARGB16F not loaded")
	}
	return _vImageHorizontalShear_ARGB16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int

// VImageHorizontalShear_ARGB16S performs a single-precision horizontal shear on a region of interest within a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB16S(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageHorizontalShear_ARGB16S == nil {
		panic("Accelerate: symbol vImageHorizontalShear_ARGB16S not loaded")
	}
	return _vImageHorizontalShear_ARGB16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int

// VImageHorizontalShear_ARGB16U performs a single-precision horizontal shear on a region of interest within an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB16U(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageHorizontalShear_ARGB16U == nil {
		panic("Accelerate: symbol vImageHorizontalShear_ARGB16U not loaded")
	}
	return _vImageHorizontalShear_ARGB16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int

// VImageHorizontalShear_ARGB8888 performs a single-precision horizontal shear on a region of interest within an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int {
	if _vImageHorizontalShear_ARGB8888 == nil {
		panic("Accelerate: symbol vImageHorizontalShear_ARGB8888 not loaded")
	}
	return _vImageHorizontalShear_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int

// VImageHorizontalShear_ARGBFFFF performs a single-precision horizontal shear on a region of interest within a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int {
	if _vImageHorizontalShear_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageHorizontalShear_ARGBFFFF not loaded")
	}
	return _vImageHorizontalShear_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int

// VImageHorizontalShear_CbCr16F performs a single-precision horizontal shear on a region of interest within a flloating-point 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_CbCr16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int {
	if _vImageHorizontalShear_CbCr16F == nil {
		panic("Accelerate: symbol vImageHorizontalShear_CbCr16F not loaded")
	}
	return _vImageHorizontalShear_CbCr16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_CbCr16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int

// VImageHorizontalShear_CbCr16S.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_CbCr16S(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_CbCr16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int {
	if _vImageHorizontalShear_CbCr16S == nil {
		panic("Accelerate: symbol vImageHorizontalShear_CbCr16S not loaded")
	}
	return _vImageHorizontalShear_CbCr16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_CbCr16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int

// VImageHorizontalShear_CbCr16U performs a single-precision horizontal shear on a region of interest within an unsigned 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_CbCr16U(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_CbCr16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int {
	if _vImageHorizontalShear_CbCr16U == nil {
		panic("Accelerate: symbol vImageHorizontalShear_CbCr16U not loaded")
	}
	return _vImageHorizontalShear_CbCr16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_CbCr8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_88, flags uint32) int

// VImageHorizontalShear_CbCr8 performs a single-precision horizontal shear on a region of interest within an 8-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_CbCr8(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_CbCr8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_88, flags uint32) int {
	if _vImageHorizontalShear_CbCr8 == nil {
		panic("Accelerate: symbol vImageHorizontalShear_CbCr8 not loaded")
	}
	return _vImageHorizontalShear_CbCr8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int

// VImageHorizontalShear_Planar16F performs a single-precision horizontal shear on a region of interest within a floating-point 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_Planar16F(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int {
	if _vImageHorizontalShear_Planar16F == nil {
		panic("Accelerate: symbol vImageHorizontalShear_Planar16F not loaded")
	}
	return _vImageHorizontalShear_Planar16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_Planar16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S, flags uint32) int

// VImageHorizontalShear_Planar16S performs a single-precision horizontal shear on a region of interest within a signed 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_Planar16S(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_Planar16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S, flags uint32) int {
	if _vImageHorizontalShear_Planar16S == nil {
		panic("Accelerate: symbol vImageHorizontalShear_Planar16S not loaded")
	}
	return _vImageHorizontalShear_Planar16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U, flags uint32) int

// VImageHorizontalShear_Planar16U performs a single-precision horizontal shear on a region of interest within an unsigned 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_Planar16U(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U, flags uint32) int {
	if _vImageHorizontalShear_Planar16U == nil {
		panic("Accelerate: symbol vImageHorizontalShear_Planar16U not loaded")
	}
	return _vImageHorizontalShear_Planar16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8, flags uint32) int

// VImageHorizontalShear_Planar8 performs a single-precision horizontal shear on a region of interest within an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8, flags uint32) int {
	if _vImageHorizontalShear_Planar8 == nil {
		panic("Accelerate: symbol vImageHorizontalShear_Planar8 not loaded")
	}
	return _vImageHorizontalShear_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_F, flags uint32) int

// VImageHorizontalShear_PlanarF performs a single-precision horizontal shear on a region of interest within a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_F, flags uint32) int {
	if _vImageHorizontalShear_PlanarF == nil {
		panic("Accelerate: symbol vImageHorizontalShear_PlanarF not loaded")
	}
	return _vImageHorizontalShear_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageHorizontalShear_XRGB2101010W func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_32U, flags uint32) int

// VImageHorizontalShear_XRGB2101010W performs a single-precision horizontal shear on a region of interest within a 2-bit alpha, 10-bit RGB interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_XRGB2101010W(_:_:_:_:_:_:_:_:_:)
func VImageHorizontalShear_XRGB2101010W(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, xTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_32U, flags uint32) int {
	if _vImageHorizontalShear_XRGB2101010W == nil {
		panic("Accelerate: symbol vImageHorizontalShear_XRGB2101010W not loaded")
	}
	return _vImageHorizontalShear_XRGB2101010W(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
}


var _vImageInterpolatedLookupTable_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, table *Pixel_F, tableEntries uint, maxFloat float32, minFloat float32, flags uint32) int

// VImageInterpolatedLookupTable_PlanarF uses an interpolated lookup table to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageInterpolatedLookupTable_PlanarF(_:_:_:_:_:_:_:)
func VImageInterpolatedLookupTable_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, table *Pixel_F, tableEntries uint, maxFloat float32, minFloat float32, flags uint32) int {
	if _vImageInterpolatedLookupTable_PlanarF == nil {
		panic("Accelerate: symbol vImageInterpolatedLookupTable_PlanarF not loaded")
	}
	return _vImageInterpolatedLookupTable_PlanarF(src, dest, table, tableEntries, maxFloat, minFloat, flags)
}


var _vImageLookupTable_8to64U func(src unsafe.Pointer, dest unsafe.Pointer, LUT uint64, flags uint32) int

// VImageLookupTable_8to64U uses a lookup table to transform an 8-bit planar image to a 64-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_8to64U(_:_:_:_:)
func VImageLookupTable_8to64U(src unsafe.Pointer, dest unsafe.Pointer, LUT uint64, flags uint32) int {
	if _vImageLookupTable_8to64U == nil {
		panic("Accelerate: symbol vImageLookupTable_8to64U not loaded")
	}
	return _vImageLookupTable_8to64U(src, dest, LUT, flags)
}


var _vImageLookupTable_Planar16 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_16U, flags uint32) int

// VImageLookupTable_Planar16 uses a lookup table to transform a 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar16(_:_:_:_:)
func VImageLookupTable_Planar16(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_16U, flags uint32) int {
	if _vImageLookupTable_Planar16 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar16 not loaded")
	}
	return _vImageLookupTable_Planar16(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanar128 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_FFFF, flags uint32) int

// VImageLookupTable_Planar8toPlanar128 uses a lookup table to transform an 8-bit planar image to a 32-bit-per-channel, four-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanar128(_:_:_:_:)
func VImageLookupTable_Planar8toPlanar128(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_FFFF, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanar128 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanar128 not loaded")
	}
	return _vImageLookupTable_Planar8toPlanar128(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanar16 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_16U, flags uint32) int

// VImageLookupTable_Planar8toPlanar16 uses a lookup table to transform an 8-bit planar image to an unsigned 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanar16(_:_:_:_:)
func VImageLookupTable_Planar8toPlanar16(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_16U, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanar16 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanar16 not loaded")
	}
	return _vImageLookupTable_Planar8toPlanar16(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanar24 func(src unsafe.Pointer, dest unsafe.Pointer, table uint32, flags uint32) int

// VImageLookupTable_Planar8toPlanar24 uses a lookup table to transform an 8-bit planar image to an 8-bit-per-channel, three-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanar24(_:_:_:_:)
func VImageLookupTable_Planar8toPlanar24(src unsafe.Pointer, dest unsafe.Pointer, table uint32, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanar24 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanar24 not loaded")
	}
	return _vImageLookupTable_Planar8toPlanar24(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanar48 func(src unsafe.Pointer, dest unsafe.Pointer, table uint64, flags uint32) int

// VImageLookupTable_Planar8toPlanar48 uses a lookup table to transform an 8-bit planar image to a 16-bit-per-channel, three-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanar48(_:_:_:_:)
func VImageLookupTable_Planar8toPlanar48(src unsafe.Pointer, dest unsafe.Pointer, table uint64, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanar48 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanar48 not loaded")
	}
	return _vImageLookupTable_Planar8toPlanar48(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanar96 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_FFFF, flags uint32) int

// VImageLookupTable_Planar8toPlanar96 uses a lookup table to transform an 8-bit planar image to a 32-bit-per-channel, three-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanar96(_:_:_:_:)
func VImageLookupTable_Planar8toPlanar96(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_FFFF, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanar96 == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanar96 not loaded")
	}
	return _vImageLookupTable_Planar8toPlanar96(src, dest, table, flags)
}


var _vImageLookupTable_Planar8toPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_F, flags uint32) int

// VImageLookupTable_Planar8toPlanarF uses a lookup table to transform an 8-bit planar image to a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_Planar8toPlanarF(_:_:_:_:)
func VImageLookupTable_Planar8toPlanarF(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_F, flags uint32) int {
	if _vImageLookupTable_Planar8toPlanarF == nil {
		panic("Accelerate: symbol vImageLookupTable_Planar8toPlanarF not loaded")
	}
	return _vImageLookupTable_Planar8toPlanarF(src, dest, table, flags)
}


var _vImageLookupTable_PlanarFtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_8, flags uint32) int

// VImageLookupTable_PlanarFtoPlanar8 uses a lookup table to transform a 32-bit planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageLookupTable_PlanarFtoPlanar8(_:_:_:_:)
func VImageLookupTable_PlanarFtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_8, flags uint32) int {
	if _vImageLookupTable_PlanarFtoPlanar8 == nil {
		panic("Accelerate: symbol vImageLookupTable_PlanarFtoPlanar8 not loaded")
	}
	return _vImageLookupTable_PlanarFtoPlanar8(src, dest, table, flags)
}


var _vImageMatrixMultiply_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int

// VImageMatrixMultiply_ARGB8888 multiplies each pixel in an interleaved four-channel, 8-bit source image by a matrix to produce an interleaved four-channel, 8-bit destination image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGB8888(_:_:_:_:_:_:_:)
func VImageMatrixMultiply_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int {
	if _vImageMatrixMultiply_ARGB8888 == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_ARGB8888 not loaded")
	}
	return _vImageMatrixMultiply_ARGB8888(src, dest, matrix, divisor, pre_bias, post_bias, flags)
}


var _vImageMatrixMultiply_ARGB8888ToPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, matrix int16, divisor int32, pre_bias int16, post_bias int32, flags uint32) int

// VImageMatrixMultiply_ARGB8888ToPlanar8 multiplies each pixel in an interleaved four-channel, 8-bit source image by a matrix to produce a planar 8-bit destination image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGB8888ToPlanar8(_:_:_:_:_:_:_:)
func VImageMatrixMultiply_ARGB8888ToPlanar8(src unsafe.Pointer, dest unsafe.Pointer, matrix int16, divisor int32, pre_bias int16, post_bias int32, flags uint32) int {
	if _vImageMatrixMultiply_ARGB8888ToPlanar8 == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_ARGB8888ToPlanar8 not loaded")
	}
	return _vImageMatrixMultiply_ARGB8888ToPlanar8(src, dest, matrix, divisor, pre_bias, post_bias, flags)
}


var _vImageMatrixMultiply_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, pre_bias *float32, post_bias *float32, flags uint32) int

// VImageMatrixMultiply_ARGBFFFF multiplies each pixel in an interleaved four-channel, 32-bit source image by a matrix to produce an interleaved four-channel, 32-bit destination image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGBFFFF(_:_:_:_:_:_:)
func VImageMatrixMultiply_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, pre_bias []float32, post_bias []float32, flags uint32) int {
	if _vImageMatrixMultiply_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_ARGBFFFF not loaded")
	}
	return _vImageMatrixMultiply_ARGBFFFF(src, dest, matrix, unsafe.SliceData(pre_bias), unsafe.SliceData(post_bias), flags)
}


var _vImageMatrixMultiply_ARGBFFFFToPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, pre_bias unsafe.Pointer, post_bias float32, flags uint32) int

// VImageMatrixMultiply_ARGBFFFFToPlanarF multiplies each pixel in an interleaved four-channel, 32-bit source image by a matrix to produce a planar 32-bit destination image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGBFFFFToPlanarF(_:_:_:_:_:_:)
func VImageMatrixMultiply_ARGBFFFFToPlanarF(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, pre_bias unsafe.Pointer, post_bias float32, flags uint32) int {
	if _vImageMatrixMultiply_ARGBFFFFToPlanarF == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_ARGBFFFFToPlanarF not loaded")
	}
	return _vImageMatrixMultiply_ARGBFFFFToPlanarF(src, dest, matrix, pre_bias, post_bias, flags)
}


var _vImageMatrixMultiply_Planar16S func(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int

// VImageMatrixMultiply_Planar16S multiplies each pixel in a set of 16-bit source image planes by a matrix to produce a set of 8-bit destination image planes.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_Planar16S(_:_:_:_:_:_:_:_:_:)
func VImageMatrixMultiply_Planar16S(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int {
	if _vImageMatrixMultiply_Planar16S == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_Planar16S not loaded")
	}
	return _vImageMatrixMultiply_Planar16S(srcs, dests, src_planes, dest_planes, matrix, divisor, pre_bias, post_bias, flags)
}


var _vImageMatrixMultiply_Planar8 func(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int

// VImageMatrixMultiply_Planar8 multiplies each pixel in a set of 8-bit source image planes by a matrix to produce a set of 8-bit destination image planes.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageMatrixMultiply_Planar8(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix int16, divisor int32, pre_bias *int16, post_bias *int32, flags uint32) int {
	if _vImageMatrixMultiply_Planar8 == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_Planar8 not loaded")
	}
	return _vImageMatrixMultiply_Planar8(srcs, dests, src_planes, dest_planes, matrix, divisor, pre_bias, post_bias, flags)
}


var _vImageMatrixMultiply_PlanarF func(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix float32, pre_bias *float32, post_bias *float32, flags uint32) int

// VImageMatrixMultiply_PlanarF multiplies each pixel in a set of 32-bit source image planes by a matrix to produce a set of 32-bit destination image planes.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_PlanarF(_:_:_:_:_:_:_:_:)
func VImageMatrixMultiply_PlanarF(srcs unsafe.Pointer, dests unsafe.Pointer, src_planes uint32, dest_planes uint32, matrix float32, pre_bias []float32, post_bias []float32, flags uint32) int {
	if _vImageMatrixMultiply_PlanarF == nil {
		panic("Accelerate: symbol vImageMatrixMultiply_PlanarF not loaded")
	}
	return _vImageMatrixMultiply_PlanarF(srcs, dests, src_planes, dest_planes, matrix, unsafe.SliceData(pre_bias), unsafe.SliceData(post_bias), flags)
}


var _vImageMax_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMax_ARGB8888 maximizes an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMax_ARGB8888(_:_:_:_:_:_:_:_:)
func VImageMax_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMax_ARGB8888 == nil {
		panic("Accelerate: symbol vImageMax_ARGB8888 not loaded")
	}
	return _vImageMax_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMax_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMax_ARGBFFFF maximizes a 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMax_ARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageMax_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMax_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageMax_ARGBFFFF not loaded")
	}
	return _vImageMax_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMax_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMax_Planar8 maximizes an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMax_Planar8(_:_:_:_:_:_:_:_:)
func VImageMax_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMax_Planar8 == nil {
		panic("Accelerate: symbol vImageMax_Planar8 not loaded")
	}
	return _vImageMax_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMax_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMax_PlanarF maximizes a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMax_PlanarF(_:_:_:_:_:_:_:_:)
func VImageMax_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMax_PlanarF == nil {
		panic("Accelerate: symbol vImageMax_PlanarF not loaded")
	}
	return _vImageMax_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMin_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMin_ARGB8888 minimizes an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMin_ARGB8888(_:_:_:_:_:_:_:_:)
func VImageMin_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMin_ARGB8888 == nil {
		panic("Accelerate: symbol vImageMin_ARGB8888 not loaded")
	}
	return _vImageMin_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMin_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMin_ARGBFFFF minimizes an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMin_ARGBFFFF(_:_:_:_:_:_:_:_:)
func VImageMin_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMin_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageMin_ARGBFFFF not loaded")
	}
	return _vImageMin_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMin_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMin_Planar8 minimizes an 8-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMin_Planar8(_:_:_:_:_:_:_:_:)
func VImageMin_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMin_Planar8 == nil {
		panic("Accelerate: symbol vImageMin_Planar8 not loaded")
	}
	return _vImageMin_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}


var _vImageMin_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int

// VImageMin_PlanarF minimizes a 32-bit planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMin_PlanarF(_:_:_:_:_:_:_:_:)
func VImageMin_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint, kernel_width uint, flags uint32) int {
	if _vImageMin_PlanarF == nil {
		panic("Accelerate: symbol vImageMin_PlanarF not loaded")
	}
	return _vImageMin_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
}







var _vImageNewResamplingFilter func(scale float32, flags uint32) ResamplingFilter

// VImageNewResamplingFilter creates a resampling filter object that corresponds to the default kernel supplied by the vImage framework.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageNewResamplingFilter(_:_:)
func VImageNewResamplingFilter(scale float32, flags uint32) ResamplingFilter {
	if _vImageNewResamplingFilter == nil {
		panic("Accelerate: symbol vImageNewResamplingFilter not loaded")
	}
	return _vImageNewResamplingFilter(scale, flags)
}


var _vImageNewResamplingFilterForFunctionUsingBuffer func(filter ResamplingFilter, scale float32, kernelWidth float32, userData unsafe.Pointer, flags uint32) int

// VImageNewResamplingFilterForFunctionUsingBuffer creates a resampling filter object that encapsulates a resampling kernel function that you provide.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageNewResamplingFilterForFunctionUsingBuffer(_:_:_:_:_:_:)
func VImageNewResamplingFilterForFunctionUsingBuffer(filter ResamplingFilter, scale float32, kernelWidth float32, userData unsafe.Pointer, flags uint32) int {
	if _vImageNewResamplingFilterForFunctionUsingBuffer == nil {
		panic("Accelerate: symbol vImageNewResamplingFilterForFunctionUsingBuffer not loaded")
	}
	return _vImageNewResamplingFilterForFunctionUsingBuffer(filter, scale, kernelWidth, userData, flags)
}


var _vImageOverwriteChannelsWithPixel_ARGB16U func(the_pixel Pixel_ARGB_16U, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannelsWithPixel_ARGB16U overwrites the channels of an unsigned 16-bit-per-channel, 4-channel interleaved buffer with the specified channels of a pixel value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithPixel_ARGB16U(_:_:_:_:_:)
func VImageOverwriteChannelsWithPixel_ARGB16U(the_pixel Pixel_ARGB_16U, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannelsWithPixel_ARGB16U == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithPixel_ARGB16U not loaded")
	}
	return _vImageOverwriteChannelsWithPixel_ARGB16U(the_pixel, src, dest, copyMask, flags)
}


var _vImageOverwriteChannelsWithPixel_ARGB8888 func(the_pixel Pixel_8888, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannelsWithPixel_ARGB8888 overwrites the channels of an 8-bit-per-channel, 4-channel interleaved buffer with the specified channels of a pixel value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithPixel_ARGB8888(_:_:_:_:_:)
func VImageOverwriteChannelsWithPixel_ARGB8888(the_pixel Pixel_8888, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannelsWithPixel_ARGB8888 == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithPixel_ARGB8888 not loaded")
	}
	return _vImageOverwriteChannelsWithPixel_ARGB8888(the_pixel, src, dest, copyMask, flags)
}


var _vImageOverwriteChannelsWithPixel_ARGBFFFF func(the_pixel Pixel_FFFF, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannelsWithPixel_ARGBFFFF overwrites the channels of a floating-point 32-bit-per-channel, 4-channel interleaved buffer with the specified channels of a pixel value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithPixel_ARGBFFFF(_:_:_:_:_:)
func VImageOverwriteChannelsWithPixel_ARGBFFFF(the_pixel Pixel_FFFF, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannelsWithPixel_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithPixel_ARGBFFFF not loaded")
	}
	return _vImageOverwriteChannelsWithPixel_ARGBFFFF(the_pixel, src, dest, copyMask, flags)
}


var _vImageOverwriteChannelsWithScalar_ARGB8888 func(scalar Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannelsWithScalar_ARGB8888 overwrites the selected channels of an 8-bit-per-channel, 4-channel interleaved buffer with the specified scalar value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_ARGB8888(_:_:_:_:_:)
func VImageOverwriteChannelsWithScalar_ARGB8888(scalar Pixel_8, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_ARGB8888 == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_ARGB8888 not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_ARGB8888(scalar, src, dest, copyMask, flags)
}


var _vImageOverwriteChannelsWithScalar_ARGBFFFF func(scalar Pixel_F, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannelsWithScalar_ARGBFFFF overwrites the selected channels of a 32-bit-per-channel, 4-channel interleaved buffer with the specified scalar value.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_ARGBFFFF(_:_:_:_:_:)
func VImageOverwriteChannelsWithScalar_ARGBFFFF(scalar Pixel_F, src unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_ARGBFFFF not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_ARGBFFFF(scalar, src, dest, copyMask, flags)
}


var _vImageOverwriteChannelsWithScalar_Planar16F func(scalar Pixel_16F, dest unsafe.Pointer, flags uint32) int

// VImageOverwriteChannelsWithScalar_Planar16F overwrites a floating-point 16-bit planar buffer with the specified scalar value in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_Planar16F(_:_:_:)
func VImageOverwriteChannelsWithScalar_Planar16F(scalar Pixel_16F, dest unsafe.Pointer, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_Planar16F == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_Planar16F not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_Planar16F(scalar, dest, flags)
}


var _vImageOverwriteChannelsWithScalar_Planar16S func(scalar Pixel_16S, dest unsafe.Pointer, flags uint32) int

// VImageOverwriteChannelsWithScalar_Planar16S overwrites a signed 16-bit planar buffer with the specified scalar value in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_Planar16S(_:_:_:)
func VImageOverwriteChannelsWithScalar_Planar16S(scalar Pixel_16S, dest unsafe.Pointer, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_Planar16S == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_Planar16S not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_Planar16S(scalar, dest, flags)
}


var _vImageOverwriteChannelsWithScalar_Planar16U func(scalar Pixel_16U, dest unsafe.Pointer, flags uint32) int

// VImageOverwriteChannelsWithScalar_Planar16U overwrites an unsigned 16-bit planar buffer with the specified scalar value in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_Planar16U(_:_:_:)
func VImageOverwriteChannelsWithScalar_Planar16U(scalar Pixel_16U, dest unsafe.Pointer, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_Planar16U == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_Planar16U not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_Planar16U(scalar, dest, flags)
}


var _vImageOverwriteChannelsWithScalar_Planar8 func(scalar Pixel_8, dest unsafe.Pointer, flags uint32) int

// VImageOverwriteChannelsWithScalar_Planar8 overwrites an 8-bit planar buffer with the specified scalar value in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_Planar8(_:_:_:)
func VImageOverwriteChannelsWithScalar_Planar8(scalar Pixel_8, dest unsafe.Pointer, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_Planar8 == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_Planar8 not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_Planar8(scalar, dest, flags)
}


var _vImageOverwriteChannelsWithScalar_PlanarF func(scalar Pixel_F, dest unsafe.Pointer, flags uint32) int

// VImageOverwriteChannelsWithScalar_PlanarF overwrites a floating-point 32-bit planar buffer with the specified scalar value in place.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannelsWithScalar_PlanarF(_:_:_:)
func VImageOverwriteChannelsWithScalar_PlanarF(scalar Pixel_F, dest unsafe.Pointer, flags uint32) int {
	if _vImageOverwriteChannelsWithScalar_PlanarF == nil {
		panic("Accelerate: symbol vImageOverwriteChannelsWithScalar_PlanarF not loaded")
	}
	return _vImageOverwriteChannelsWithScalar_PlanarF(scalar, dest, flags)
}


var _vImageOverwriteChannels_ARGB8888 func(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannels_ARGB8888 overwrites the channels of an 8-bit-per-channel, 4-channel interleaved buffer with the corresponding pixels of a planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannels_ARGB8888(_:_:_:_:_:)
func VImageOverwriteChannels_ARGB8888(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannels_ARGB8888 == nil {
		panic("Accelerate: symbol vImageOverwriteChannels_ARGB8888 not loaded")
	}
	return _vImageOverwriteChannels_ARGB8888(newSrc, origSrc, dest, copyMask, flags)
}


var _vImageOverwriteChannels_ARGBFFFF func(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageOverwriteChannels_ARGBFFFF overwrites the channels of a floating-point 32-bit-per-channel, 4-channel interleaved buffer with the corresponding pixels of a planar buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageOverwriteChannels_ARGBFFFF(_:_:_:_:_:)
func VImageOverwriteChannels_ARGBFFFF(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageOverwriteChannels_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageOverwriteChannels_ARGBFFFF not loaded")
	}
	return _vImageOverwriteChannels_ARGBFFFF(newSrc, origSrc, dest, copyMask, flags)
}


var _vImagePNGDecompressionFilter func(buffer unsafe.Pointer, startScanline uint, scanlineCount uint, bitsPerPixel uint32, filterMethodNumber uint32, filterType uint32, flags uint32) int

// VImagePNGDecompressionFilter performs PNG decompression filtering.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePNGDecompressionFilter(_:_:_:_:_:_:_:)
func VImagePNGDecompressionFilter(buffer unsafe.Pointer, startScanline uint, scanlineCount uint, bitsPerPixel uint32, filterMethodNumber uint32, filterType uint32, flags uint32) int {
	if _vImagePNGDecompressionFilter == nil {
		panic("Accelerate: symbol vImagePNGDecompressionFilter not loaded")
	}
	return _vImagePNGDecompressionFilter(buffer, startScanline, scanlineCount, bitsPerPixel, filterMethodNumber, filterType, flags)
}


var _vImagePermuteChannelsWithMaskedInsert_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_ARGB_16U, flags uint32) int

// VImagePermuteChannelsWithMaskedInsert_ARGB16U permutes and overwrites the channels of an unsigned 16-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannelsWithMaskedInsert_ARGB16U(_:_:_:_:_:_:)
func VImagePermuteChannelsWithMaskedInsert_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_ARGB_16U, flags uint32) int {
	if _vImagePermuteChannelsWithMaskedInsert_ARGB16U == nil {
		panic("Accelerate: symbol vImagePermuteChannelsWithMaskedInsert_ARGB16U not loaded")
	}
	return _vImagePermuteChannelsWithMaskedInsert_ARGB16U(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImagePermuteChannelsWithMaskedInsert_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int

// VImagePermuteChannelsWithMaskedInsert_ARGB8888 permutes and overwrites the channels of an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannelsWithMaskedInsert_ARGB8888(_:_:_:_:_:_:)
func VImagePermuteChannelsWithMaskedInsert_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_8888, flags uint32) int {
	if _vImagePermuteChannelsWithMaskedInsert_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePermuteChannelsWithMaskedInsert_ARGB8888 not loaded")
	}
	return _vImagePermuteChannelsWithMaskedInsert_ARGB8888(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImagePermuteChannelsWithMaskedInsert_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_FFFF, flags uint32) int

// VImagePermuteChannelsWithMaskedInsert_ARGBFFFF permutes and overwrites the channels of a floating-point 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannelsWithMaskedInsert_ARGBFFFF(_:_:_:_:_:_:)
func VImagePermuteChannelsWithMaskedInsert_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, copyMask uint8, backgroundColor Pixel_FFFF, flags uint32) int {
	if _vImagePermuteChannelsWithMaskedInsert_ARGBFFFF == nil {
		panic("Accelerate: symbol vImagePermuteChannelsWithMaskedInsert_ARGBFFFF not loaded")
	}
	return _vImagePermuteChannelsWithMaskedInsert_ARGBFFFF(src, dest, permuteMap, copyMask, backgroundColor, flags)
}


var _vImagePermuteChannels_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int

// VImagePermuteChannels_ARGB16F permutes the channels of a floating-point 16-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannels_ARGB16F(_:_:_:_:)
func VImagePermuteChannels_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int {
	if _vImagePermuteChannels_ARGB16F == nil {
		panic("Accelerate: symbol vImagePermuteChannels_ARGB16F not loaded")
	}
	return _vImagePermuteChannels_ARGB16F(src, dest, permuteMap, flags)
}


var _vImagePermuteChannels_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int

// VImagePermuteChannels_ARGB16U permutes the channels of an unsigned 16-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannels_ARGB16U(_:_:_:_:)
func VImagePermuteChannels_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int {
	if _vImagePermuteChannels_ARGB16U == nil {
		panic("Accelerate: symbol vImagePermuteChannels_ARGB16U not loaded")
	}
	return _vImagePermuteChannels_ARGB16U(src, dest, permuteMap, flags)
}


var _vImagePermuteChannels_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int

// VImagePermuteChannels_ARGB8888 permutes the channels of an 8-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannels_ARGB8888(_:_:_:_:)
func VImagePermuteChannels_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int {
	if _vImagePermuteChannels_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePermuteChannels_ARGB8888 not loaded")
	}
	return _vImagePermuteChannels_ARGB8888(src, dest, permuteMap, flags)
}


var _vImagePermuteChannels_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int

// VImagePermuteChannels_ARGBFFFF permutes the channels of a floating-point 32-bit-per-channel, 4-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannels_ARGBFFFF(_:_:_:_:)
func VImagePermuteChannels_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int {
	if _vImagePermuteChannels_ARGBFFFF == nil {
		panic("Accelerate: symbol vImagePermuteChannels_ARGBFFFF not loaded")
	}
	return _vImagePermuteChannels_ARGBFFFF(src, dest, permuteMap, flags)
}


var _vImagePermuteChannels_RGB888 func(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int

// VImagePermuteChannels_RGB888 permutes the channels of an 8-bit-per-channel, 3-channel interleaved buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePermuteChannels_RGB888(_:_:_:_:)
func VImagePermuteChannels_RGB888(src unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, flags uint32) int {
	if _vImagePermuteChannels_RGB888 == nil {
		panic("Accelerate: symbol vImagePermuteChannels_RGB888 not loaded")
	}
	return _vImagePermuteChannels_RGB888(src, dest, permuteMap, flags)
}








var _vImagePiecewiseGamma_Planar16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int

// VImagePiecewiseGamma_Planar16Q12 applies a piecewise gamma function to transform a 16Q12 planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_Planar16Q12(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_Planar16Q12(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int {
	if _vImagePiecewiseGamma_Planar16Q12 == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_Planar16Q12 not loaded")
	}
	return _vImagePiecewiseGamma_Planar16Q12(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_Planar16Q12toPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int

// VImagePiecewiseGamma_Planar16Q12toPlanar8 applies a piecewise gamma function to transform a 16Q12 planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_Planar16Q12toPlanar8(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_Planar16Q12toPlanar8(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int {
	if _vImagePiecewiseGamma_Planar16Q12toPlanar8 == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_Planar16Q12toPlanar8 not loaded")
	}
	return _vImagePiecewiseGamma_Planar16Q12toPlanar8(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int

// VImagePiecewiseGamma_Planar8 applies a piecewise gamma function to transform an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_Planar8(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_Planar8(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int {
	if _vImagePiecewiseGamma_Planar8 == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_Planar8 not loaded")
	}
	return _vImagePiecewiseGamma_Planar8(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_Planar8toPlanar16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int

// VImagePiecewiseGamma_Planar8toPlanar16Q12 applies a piecewise gamma function to transform an 8-bit planar image to a 16Q12 planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_Planar8toPlanar16Q12(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_Planar8toPlanar16Q12(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int {
	if _vImagePiecewiseGamma_Planar8toPlanar16Q12 == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_Planar8toPlanar16Q12 not loaded")
	}
	return _vImagePiecewiseGamma_Planar8toPlanar16Q12(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_Planar8toPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int

// VImagePiecewiseGamma_Planar8toPlanarF applies a piecewise gamma function to transform an 8-bit planar image to a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_Planar8toPlanarF(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_Planar8toPlanarF(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_8, flags uint32) int {
	if _vImagePiecewiseGamma_Planar8toPlanarF == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_Planar8toPlanarF not loaded")
	}
	return _vImagePiecewiseGamma_Planar8toPlanarF(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int

// VImagePiecewiseGamma_PlanarF applies a piecewise gamma function to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_PlanarF(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int {
	if _vImagePiecewiseGamma_PlanarF == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_PlanarF not loaded")
	}
	return _vImagePiecewiseGamma_PlanarF(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewiseGamma_PlanarFtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int

// VImagePiecewiseGamma_PlanarFtoPlanar8 applies a piecewise gamma function to transform a 32-bit planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseGamma_PlanarFtoPlanar8(_:_:_:_:_:_:_:)
func VImagePiecewiseGamma_PlanarFtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int {
	if _vImagePiecewiseGamma_PlanarFtoPlanar8 == nil {
		panic("Accelerate: symbol vImagePiecewiseGamma_PlanarFtoPlanar8 not loaded")
	}
	return _vImagePiecewiseGamma_PlanarFtoPlanar8(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImagePiecewisePolynomial_Planar8toPlanarF func(src unsafe.Pointer, dest unsafe.Pointer, coefficients *float32, boundaries *float32, order uint32, log2segments uint32, flags uint32) int

// VImagePiecewisePolynomial_Planar8toPlanarF applies a set of piecewise polynomials to transform an 8-bit planar image to a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewisePolynomial_Planar8toPlanarF(_:_:_:_:_:_:_:)
func VImagePiecewisePolynomial_Planar8toPlanarF(src unsafe.Pointer, dest unsafe.Pointer, coefficients []float32, boundaries []float32, order uint32, log2segments uint32, flags uint32) int {
	if _vImagePiecewisePolynomial_Planar8toPlanarF == nil {
		panic("Accelerate: symbol vImagePiecewisePolynomial_Planar8toPlanarF not loaded")
	}
	return _vImagePiecewisePolynomial_Planar8toPlanarF(src, dest, unsafe.SliceData(coefficients), unsafe.SliceData(boundaries), order, log2segments, flags)
}


var _vImagePiecewisePolynomial_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, coefficients *float32, boundaries *float32, order uint32, log2segments uint32, flags uint32) int

// VImagePiecewisePolynomial_PlanarF applies a set of piecewise polynomials to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewisePolynomial_PlanarF(_:_:_:_:_:_:_:)
func VImagePiecewisePolynomial_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, coefficients []float32, boundaries []float32, order uint32, log2segments uint32, flags uint32) int {
	if _vImagePiecewisePolynomial_PlanarF == nil {
		panic("Accelerate: symbol vImagePiecewisePolynomial_PlanarF not loaded")
	}
	return _vImagePiecewisePolynomial_PlanarF(src, dest, unsafe.SliceData(coefficients), unsafe.SliceData(boundaries), order, log2segments, flags)
}


var _vImagePiecewisePolynomial_PlanarFtoPlanar8 func(src unsafe.Pointer, dest unsafe.Pointer, coefficients *float32, boundaries *float32, order uint32, log2segments uint32, flags uint32) int

// VImagePiecewisePolynomial_PlanarFtoPlanar8 applies a set of piecewise polynomials to transform a 32-bit planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewisePolynomial_PlanarFtoPlanar8(_:_:_:_:_:_:_:)
func VImagePiecewisePolynomial_PlanarFtoPlanar8(src unsafe.Pointer, dest unsafe.Pointer, coefficients []float32, boundaries []float32, order uint32, log2segments uint32, flags uint32) int {
	if _vImagePiecewisePolynomial_PlanarFtoPlanar8 == nil {
		panic("Accelerate: symbol vImagePiecewisePolynomial_PlanarFtoPlanar8 not loaded")
	}
	return _vImagePiecewisePolynomial_PlanarFtoPlanar8(src, dest, unsafe.SliceData(coefficients), unsafe.SliceData(boundaries), order, log2segments, flags)
}


var _vImagePiecewiseRational_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, topCoefficients *float32, bottomCoefficients *float32, boundaries *float32, topOrder uint32, bottomOrder uint32, log2segments uint32, flags uint32) int

// VImagePiecewiseRational_PlanarF applies a set of piecewise rational expressions to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePiecewiseRational_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImagePiecewiseRational_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, topCoefficients []float32, bottomCoefficients []float32, boundaries []float32, topOrder uint32, bottomOrder uint32, log2segments uint32, flags uint32) int {
	if _vImagePiecewiseRational_PlanarF == nil {
		panic("Accelerate: symbol vImagePiecewiseRational_PlanarF not loaded")
	}
	return _vImagePiecewiseRational_PlanarF(src, dest, unsafe.SliceData(topCoefficients), unsafe.SliceData(bottomCoefficients), unsafe.SliceData(boundaries), topOrder, bottomOrder, log2segments, flags)
}


var _vImagePremultipliedAlphaBlendDarken_RGBA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlendDarken_RGBA8888 performs alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers using the darken blend mode.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendDarken_RGBA8888(_:_:_:_:)
func VImagePremultipliedAlphaBlendDarken_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlendDarken_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendDarken_RGBA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendDarken_RGBA8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlendLighten_RGBA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlendLighten_RGBA8888 performs alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers using the lighten blend mode.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendLighten_RGBA8888(_:_:_:_:)
func VImagePremultipliedAlphaBlendLighten_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlendLighten_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendLighten_RGBA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendLighten_RGBA8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlendMultiply_RGBA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlendMultiply_RGBA8888 performs alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers using the multiply blend mode.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendMultiply_RGBA8888(_:_:_:_:)
func VImagePremultipliedAlphaBlendMultiply_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlendMultiply_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendMultiply_RGBA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendMultiply_RGBA8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlendScreen_RGBA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlendScreen_RGBA8888 performs alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers using the screen blend mode.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendScreen_RGBA8888(_:_:_:_:)
func VImagePremultipliedAlphaBlendScreen_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlendScreen_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendScreen_RGBA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendScreen_RGBA8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlendWithPermute_ARGB8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, makeDestAlphaOpaque bool, flags uint32) int

// VImagePremultipliedAlphaBlendWithPermute_ARGB8888 permutes the top 8-bit, 4-channel premultiplied buffer, and composites with the bottom buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendWithPermute_ARGB8888(_:_:_:_:_:_:)
func VImagePremultipliedAlphaBlendWithPermute_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, makeDestAlphaOpaque bool, flags uint32) int {
	if _vImagePremultipliedAlphaBlendWithPermute_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendWithPermute_ARGB8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendWithPermute_ARGB8888(srcTop, srcBottom, dest, permuteMap, makeDestAlphaOpaque, flags)
}


var _vImagePremultipliedAlphaBlendWithPermute_RGBA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, makeDestAlphaOpaque bool, flags uint32) int

// VImagePremultipliedAlphaBlendWithPermute_RGBA8888 permutes the top 8-bit, 4-channel premultiplied buffer, and composites with the bottom buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendWithPermute_RGBA8888(_:_:_:_:_:_:)
func VImagePremultipliedAlphaBlendWithPermute_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, permuteMap uint8, makeDestAlphaOpaque bool, flags uint32) int {
	if _vImagePremultipliedAlphaBlendWithPermute_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlendWithPermute_RGBA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlendWithPermute_RGBA8888(srcTop, srcBottom, dest, permuteMap, makeDestAlphaOpaque, flags)
}


var _vImagePremultipliedAlphaBlend_ARGB8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_ARGB8888 performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_ARGB8888(_:_:_:_:)
func VImagePremultipliedAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_ARGB8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlend_ARGBFFFF func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_ARGBFFFF performs premultiplied alpha compositing of two 32-bit-per-channel, 4-channel ARGB buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_ARGBFFFF(_:_:_:_:)
func VImagePremultipliedAlphaBlend_ARGBFFFF(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_ARGBFFFF == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_ARGBFFFF not loaded")
	}
	return _vImagePremultipliedAlphaBlend_ARGBFFFF(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlend_BGRA8888 func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_BGRA8888 performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_BGRA8888(_:_:_:_:)
func VImagePremultipliedAlphaBlend_BGRA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_BGRA8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_BGRA8888 not loaded")
	}
	return _vImagePremultipliedAlphaBlend_BGRA8888(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlend_BGRAFFFF func(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_BGRAFFFF performs premultiplied alpha compositing of two 32-bit-per-channel, 4-channel BGRA buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_BGRAFFFF(_:_:_:_:)
func VImagePremultipliedAlphaBlend_BGRAFFFF(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_BGRAFFFF == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_BGRAFFFF not loaded")
	}
	return _vImagePremultipliedAlphaBlend_BGRAFFFF(srcTop, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlend_Planar8 func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_Planar8 performs premultiplied alpha compositing of two 8-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_Planar8(_:_:_:_:_:)
func VImagePremultipliedAlphaBlend_Planar8(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_Planar8 == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_Planar8 not loaded")
	}
	return _vImagePremultipliedAlphaBlend_Planar8(srcTop, srcTopAlpha, srcBottom, dest, flags)
}


var _vImagePremultipliedAlphaBlend_PlanarF func(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedAlphaBlend_PlanarF performs premultiplied alpha compositing of two 32-bit planar buffers.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_PlanarF(_:_:_:_:_:)
func VImagePremultipliedAlphaBlend_PlanarF(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedAlphaBlend_PlanarF == nil {
		panic("Accelerate: symbol vImagePremultipliedAlphaBlend_PlanarF not loaded")
	}
	return _vImagePremultipliedAlphaBlend_PlanarF(srcTop, srcTopAlpha, srcBottom, dest, flags)
}


var _vImagePremultipliedConstAlphaBlend_ARGB8888 func(srcTop unsafe.Pointer, constAlpha Pixel_8, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedConstAlphaBlend_ARGB8888 performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel interleaved buffers and applies an extra alpha value to the top buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_ARGB8888(_:_:_:_:_:)
func VImagePremultipliedConstAlphaBlend_ARGB8888(srcTop unsafe.Pointer, constAlpha Pixel_8, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedConstAlphaBlend_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePremultipliedConstAlphaBlend_ARGB8888 not loaded")
	}
	return _vImagePremultipliedConstAlphaBlend_ARGB8888(srcTop, constAlpha, srcBottom, dest, flags)
}


var _vImagePremultipliedConstAlphaBlend_ARGBFFFF func(srcTop unsafe.Pointer, constAlpha Pixel_F, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedConstAlphaBlend_ARGBFFFF performs premultiplied alpha compositing of two 32-bit-per-channel, 4-channel interleaved buffers and applies an extra alpha value to the top buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_ARGBFFFF(_:_:_:_:_:)
func VImagePremultipliedConstAlphaBlend_ARGBFFFF(srcTop unsafe.Pointer, constAlpha Pixel_F, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedConstAlphaBlend_ARGBFFFF == nil {
		panic("Accelerate: symbol vImagePremultipliedConstAlphaBlend_ARGBFFFF not loaded")
	}
	return _vImagePremultipliedConstAlphaBlend_ARGBFFFF(srcTop, constAlpha, srcBottom, dest, flags)
}


var _vImagePremultipliedConstAlphaBlend_Planar8 func(srcTop unsafe.Pointer, constAlpha Pixel_8, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedConstAlphaBlend_Planar8 performs premultiplied alpha compositing of two 8-bit planar buffers and applies an extra alpha value to the top buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_Planar8(_:_:_:_:_:_:)
func VImagePremultipliedConstAlphaBlend_Planar8(srcTop unsafe.Pointer, constAlpha Pixel_8, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedConstAlphaBlend_Planar8 == nil {
		panic("Accelerate: symbol vImagePremultipliedConstAlphaBlend_Planar8 not loaded")
	}
	return _vImagePremultipliedConstAlphaBlend_Planar8(srcTop, constAlpha, srcTopAlpha, srcBottom, dest, flags)
}


var _vImagePremultipliedConstAlphaBlend_PlanarF func(srcTop unsafe.Pointer, constAlpha Pixel_F, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultipliedConstAlphaBlend_PlanarF performs premultiplied alpha compositing of two 32-bit planar buffers and applies an extra alpha value to the top buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_PlanarF(_:_:_:_:_:_:)
func VImagePremultipliedConstAlphaBlend_PlanarF(srcTop unsafe.Pointer, constAlpha Pixel_F, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultipliedConstAlphaBlend_PlanarF == nil {
		panic("Accelerate: symbol vImagePremultipliedConstAlphaBlend_PlanarF not loaded")
	}
	return _vImagePremultipliedConstAlphaBlend_PlanarF(srcTop, constAlpha, srcTopAlpha, srcBottom, dest, flags)
}


var _vImagePremultiplyData_ARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_ARGB16Q12 transforms a fixed-point 16-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_ARGB16Q12(_:_:_:)
func VImagePremultiplyData_ARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_ARGB16Q12 == nil {
		panic("Accelerate: symbol vImagePremultiplyData_ARGB16Q12 not loaded")
	}
	return _vImagePremultiplyData_ARGB16Q12(src, dest, flags)
}


var _vImagePremultiplyData_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_ARGB16U transforms an unsigned 16-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_ARGB16U(_:_:_:)
func VImagePremultiplyData_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_ARGB16U == nil {
		panic("Accelerate: symbol vImagePremultiplyData_ARGB16U not loaded")
	}
	return _vImagePremultiplyData_ARGB16U(src, dest, flags)
}


var _vImagePremultiplyData_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_ARGB8888 transforms an 8-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_ARGB8888(_:_:_:)
func VImagePremultiplyData_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_ARGB8888 == nil {
		panic("Accelerate: symbol vImagePremultiplyData_ARGB8888 not loaded")
	}
	return _vImagePremultiplyData_ARGB8888(src, dest, flags)
}


var _vImagePremultiplyData_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_ARGBFFFF transforms a floating-point 32-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_ARGBFFFF(_:_:_:)
func VImagePremultiplyData_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_ARGBFFFF == nil {
		panic("Accelerate: symbol vImagePremultiplyData_ARGBFFFF not loaded")
	}
	return _vImagePremultiplyData_ARGBFFFF(src, dest, flags)
}


var _vImagePremultiplyData_Planar8 func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_Planar8 transforms an 8-bit planar buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_Planar8(_:_:_:_:)
func VImagePremultiplyData_Planar8(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_Planar8 == nil {
		panic("Accelerate: symbol vImagePremultiplyData_Planar8 not loaded")
	}
	return _vImagePremultiplyData_Planar8(src, alpha, dest, flags)
}


var _vImagePremultiplyData_PlanarF func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_PlanarF transforms a 32-bit planar buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_PlanarF(_:_:_:_:)
func VImagePremultiplyData_PlanarF(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_PlanarF == nil {
		panic("Accelerate: symbol vImagePremultiplyData_PlanarF not loaded")
	}
	return _vImagePremultiplyData_PlanarF(src, alpha, dest, flags)
}


var _vImagePremultiplyData_RGBA16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_RGBA16F transforms a floating-point 16-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA16F(_:_:_:)
func VImagePremultiplyData_RGBA16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_RGBA16F == nil {
		panic("Accelerate: symbol vImagePremultiplyData_RGBA16F not loaded")
	}
	return _vImagePremultiplyData_RGBA16F(src, dest, flags)
}


var _vImagePremultiplyData_RGBA16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_RGBA16Q12 transforms a fixed-point 16-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA16Q12(_:_:_:)
func VImagePremultiplyData_RGBA16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_RGBA16Q12 == nil {
		panic("Accelerate: symbol vImagePremultiplyData_RGBA16Q12 not loaded")
	}
	return _vImagePremultiplyData_RGBA16Q12(src, dest, flags)
}


var _vImagePremultiplyData_RGBA16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_RGBA16U transforms an unsigned 16-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA16U(_:_:_:)
func VImagePremultiplyData_RGBA16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_RGBA16U == nil {
		panic("Accelerate: symbol vImagePremultiplyData_RGBA16U not loaded")
	}
	return _vImagePremultiplyData_RGBA16U(src, dest, flags)
}


var _vImagePremultiplyData_RGBA8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_RGBA8888 transforms an 8-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA8888(_:_:_:)
func VImagePremultiplyData_RGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_RGBA8888 == nil {
		panic("Accelerate: symbol vImagePremultiplyData_RGBA8888 not loaded")
	}
	return _vImagePremultiplyData_RGBA8888(src, dest, flags)
}


var _vImagePremultiplyData_RGBAFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImagePremultiplyData_RGBAFFFF transforms a floating-point 32-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBAFFFF(_:_:_:)
func VImagePremultiplyData_RGBAFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImagePremultiplyData_RGBAFFFF == nil {
		panic("Accelerate: symbol vImagePremultiplyData_RGBAFFFF not loaded")
	}
	return _vImagePremultiplyData_RGBAFFFF(src, dest, flags)
}


var _vImageRichardsonLucyDeConvolve_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel2 *int16, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, divisor int32, divisor2 int32, backgroundColor Pixel_8888, iterationCount uint32, flags uint32) int

// VImageRichardsonLucyDeConvolve_ARGB8888 deconvolves an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRichardsonLucyDeConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageRichardsonLucyDeConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel2 *int16, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, divisor int32, divisor2 int32, backgroundColor Pixel_8888, iterationCount uint32, flags uint32) int {
	if _vImageRichardsonLucyDeConvolve_ARGB8888 == nil {
		panic("Accelerate: symbol vImageRichardsonLucyDeConvolve_ARGB8888 not loaded")
	}
	return _vImageRichardsonLucyDeConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel2, kernel_height, kernel_width, kernel_height2, kernel_width2, divisor, divisor2, backgroundColor, iterationCount, flags)
}


var _vImageRichardsonLucyDeConvolve_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel2 *float32, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, backgroundColor Pixel_FFFF, iterationCount uint32, flags uint32) int

// VImageRichardsonLucyDeConvolve_ARGBFFFF deconvolves a floating-point 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRichardsonLucyDeConvolve_ARGBFFFF(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageRichardsonLucyDeConvolve_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel2 []float32, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, backgroundColor Pixel_FFFF, iterationCount uint32, flags uint32) int {
	if _vImageRichardsonLucyDeConvolve_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageRichardsonLucyDeConvolve_ARGBFFFF not loaded")
	}
	return _vImageRichardsonLucyDeConvolve_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), unsafe.SliceData(kernel2), kernel_height, kernel_width, kernel_height2, kernel_width2, backgroundColor, iterationCount, flags)
}


var _vImageRichardsonLucyDeConvolve_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel2 *int16, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, divisor int32, divisor2 int32, backgroundColor Pixel_8, iterationCount uint32, flags uint32) int

// VImageRichardsonLucyDeConvolve_Planar8 deconvolves an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRichardsonLucyDeConvolve_Planar8(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageRichardsonLucyDeConvolve_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *int16, kernel2 *int16, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, divisor int32, divisor2 int32, backgroundColor Pixel_8, iterationCount uint32, flags uint32) int {
	if _vImageRichardsonLucyDeConvolve_Planar8 == nil {
		panic("Accelerate: symbol vImageRichardsonLucyDeConvolve_Planar8 not loaded")
	}
	return _vImageRichardsonLucyDeConvolve_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel2, kernel_height, kernel_width, kernel_height2, kernel_width2, divisor, divisor2, backgroundColor, iterationCount, flags)
}


var _vImageRichardsonLucyDeConvolve_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel *float32, kernel2 *float32, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, backgroundColor Pixel_F, iterationCount uint32, flags uint32) int

// VImageRichardsonLucyDeConvolve_PlanarF deconvolves a floating-point 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRichardsonLucyDeConvolve_PlanarF(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageRichardsonLucyDeConvolve_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel []float32, kernel2 []float32, kernel_height uint32, kernel_width uint32, kernel_height2 uint32, kernel_width2 uint32, backgroundColor Pixel_F, iterationCount uint32, flags uint32) int {
	if _vImageRichardsonLucyDeConvolve_PlanarF == nil {
		panic("Accelerate: symbol vImageRichardsonLucyDeConvolve_PlanarF not loaded")
	}
	return _vImageRichardsonLucyDeConvolve_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernel), unsafe.SliceData(kernel2), kernel_height, kernel_width, kernel_height2, kernel_width2, backgroundColor, iterationCount, flags)
}


var _vImageRotate90_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16F, flags uint32) int

// VImageRotate90_ARGB16F rotates a floating-point 16-bit-per-channel, 4-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_ARGB16F(_:_:_:_:_:)
func VImageRotate90_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageRotate90_ARGB16F == nil {
		panic("Accelerate: symbol vImageRotate90_ARGB16F not loaded")
	}
	return _vImageRotate90_ARGB16F(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16S, flags uint32) int

// VImageRotate90_ARGB16S rotates a signed 16-bit-per-channel, 4-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_ARGB16S(_:_:_:_:_:)
func VImageRotate90_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageRotate90_ARGB16S == nil {
		panic("Accelerate: symbol vImageRotate90_ARGB16S not loaded")
	}
	return _vImageRotate90_ARGB16S(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16U, flags uint32) int

// VImageRotate90_ARGB16U rotates an unsigned 16-bit-per-channel, 4-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_ARGB16U(_:_:_:_:_:)
func VImageRotate90_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageRotate90_ARGB16U == nil {
		panic("Accelerate: symbol vImageRotate90_ARGB16U not loaded")
	}
	return _vImageRotate90_ARGB16U(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_8888, flags uint32) int

// VImageRotate90_ARGB8888 rotates an 8-bit-per-channel, 4-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_ARGB8888(_:_:_:_:_:)
func VImageRotate90_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_8888, flags uint32) int {
	if _vImageRotate90_ARGB8888 == nil {
		panic("Accelerate: symbol vImageRotate90_ARGB8888 not loaded")
	}
	return _vImageRotate90_ARGB8888(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_FFFF, flags uint32) int

// VImageRotate90_ARGBFFFF rotates a 32-bit-per-channel, 4-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_ARGBFFFF(_:_:_:_:_:)
func VImageRotate90_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_FFFF, flags uint32) int {
	if _vImageRotate90_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageRotate90_ARGBFFFF not loaded")
	}
	return _vImageRotate90_ARGBFFFF(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16F16F, flags uint32) int

// VImageRotate90_CbCr16F rotates a floating-point 16-bit-per-channel, 2-channel interleaved image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_CbCr16F(_:_:_:_:_:)
func VImageRotate90_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16F16F, flags uint32) int {
	if _vImageRotate90_CbCr16F == nil {
		panic("Accelerate: symbol vImageRotate90_CbCr16F not loaded")
	}
	return _vImageRotate90_CbCr16F(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16F, flags uint32) int

// VImageRotate90_Planar16F rotates a floating-point 16-bit planar image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_Planar16F(_:_:_:_:_:)
func VImageRotate90_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16F, flags uint32) int {
	if _vImageRotate90_Planar16F == nil {
		panic("Accelerate: symbol vImageRotate90_Planar16F not loaded")
	}
	return _vImageRotate90_Planar16F(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16U, flags uint32) int

// VImageRotate90_Planar16U rotates an unsigned 16-bit planar image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_Planar16U(_:_:_:_:_:)
func VImageRotate90_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_16U, flags uint32) int {
	if _vImageRotate90_Planar16U == nil {
		panic("Accelerate: symbol vImageRotate90_Planar16U not loaded")
	}
	return _vImageRotate90_Planar16U(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_8, flags uint32) int

// VImageRotate90_Planar8 rotates an 8-bit planar image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_Planar8(_:_:_:_:_:)
func VImageRotate90_Planar8(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_8, flags uint32) int {
	if _vImageRotate90_Planar8 == nil {
		panic("Accelerate: symbol vImageRotate90_Planar8 not loaded")
	}
	return _vImageRotate90_Planar8(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate90_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_F, flags uint32) int

// VImageRotate90_PlanarF rotates a 32-bit planar image by a multiple of 90°.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate90_PlanarF(_:_:_:_:_:)
func VImageRotate90_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, rotationConstant uint8, backColor Pixel_F, flags uint32) int {
	if _vImageRotate90_PlanarF == nil {
		panic("Accelerate: symbol vImageRotate90_PlanarF not loaded")
	}
	return _vImageRotate90_PlanarF(src, dest, rotationConstant, backColor, flags)
}


var _vImageRotate_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16F, flags uint32) int

// VImageRotate_ARGB16F rotates a floating-point 16-bit-per-channel, 4-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_ARGB16F(_:_:_:_:_:_:)
func VImageRotate_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageRotate_ARGB16F == nil {
		panic("Accelerate: symbol vImageRotate_ARGB16F not loaded")
	}
	return _vImageRotate_ARGB16F(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16S, flags uint32) int

// VImageRotate_ARGB16S rotates a signed 16-bit-per-channel, 4-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_ARGB16S(_:_:_:_:_:_:)
func VImageRotate_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageRotate_ARGB16S == nil {
		panic("Accelerate: symbol vImageRotate_ARGB16S not loaded")
	}
	return _vImageRotate_ARGB16S(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16U, flags uint32) int

// VImageRotate_ARGB16U rotates an unsigned 16-bit-per-channel, 4-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_ARGB16U(_:_:_:_:_:_:)
func VImageRotate_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageRotate_ARGB16U == nil {
		panic("Accelerate: symbol vImageRotate_ARGB16U not loaded")
	}
	return _vImageRotate_ARGB16U(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_8888, flags uint32) int

// VImageRotate_ARGB8888 rotates an 8-bit-per-channel, 4-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_ARGB8888(_:_:_:_:_:_:)
func VImageRotate_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_8888, flags uint32) int {
	if _vImageRotate_ARGB8888 == nil {
		panic("Accelerate: symbol vImageRotate_ARGB8888 not loaded")
	}
	return _vImageRotate_ARGB8888(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_FFFF, flags uint32) int

// VImageRotate_ARGBFFFF rotates a 32-bit-per-channel, 4-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_ARGBFFFF(_:_:_:_:_:_:)
func VImageRotate_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_FFFF, flags uint32) int {
	if _vImageRotate_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageRotate_ARGBFFFF not loaded")
	}
	return _vImageRotate_ARGBFFFF(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_16F16F, flags uint32) int

// VImageRotate_CbCr16F rotates a floating-point 16-bit-per-channel, 2-channel interleaved image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_CbCr16F(_:_:_:_:_:_:)
func VImageRotate_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_16F16F, flags uint32) int {
	if _vImageRotate_CbCr16F == nil {
		panic("Accelerate: symbol vImageRotate_CbCr16F not loaded")
	}
	return _vImageRotate_CbCr16F(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_16F, flags uint32) int

// VImageRotate_Planar16F rotates a floating-point 16-bit planar image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_Planar16F(_:_:_:_:_:_:)
func VImageRotate_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_16F, flags uint32) int {
	if _vImageRotate_Planar16F == nil {
		panic("Accelerate: symbol vImageRotate_Planar16F not loaded")
	}
	return _vImageRotate_Planar16F(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_8, flags uint32) int

// VImageRotate_Planar8 rotates an 8-bit planar image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_Planar8(_:_:_:_:_:_:)
func VImageRotate_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_8, flags uint32) int {
	if _vImageRotate_Planar8 == nil {
		panic("Accelerate: symbol vImageRotate_Planar8 not loaded")
	}
	return _vImageRotate_Planar8(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageRotate_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_F, flags uint32) int

// VImageRotate_PlanarF rotates a 32-bit planar image by any angle, which you specify in radians.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageRotate_PlanarF(_:_:_:_:_:_:)
func VImageRotate_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, angleInRadians float32, backColor Pixel_F, flags uint32) int {
	if _vImageRotate_PlanarF == nil {
		panic("Accelerate: symbol vImageRotate_PlanarF not loaded")
	}
	return _vImageRotate_PlanarF(src, dest, tempBuffer, angleInRadians, backColor, flags)
}


var _vImageScale_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_ARGB16F scales a floating-point 16-bit-per-channel, 4-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_ARGB16F(_:_:_:_:)
func VImageScale_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_ARGB16F == nil {
		panic("Accelerate: symbol vImageScale_ARGB16F not loaded")
	}
	return _vImageScale_ARGB16F(src, dest, tempBuffer, flags)
}


var _vImageScale_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_ARGB16S scales a signed 16-bit-per-channel, 4-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_ARGB16S(_:_:_:_:)
func VImageScale_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_ARGB16S == nil {
		panic("Accelerate: symbol vImageScale_ARGB16S not loaded")
	}
	return _vImageScale_ARGB16S(src, dest, tempBuffer, flags)
}


var _vImageScale_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_ARGB16U scales an unsigned 16-bit-per-channel, 4-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_ARGB16U(_:_:_:_:)
func VImageScale_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_ARGB16U == nil {
		panic("Accelerate: symbol vImageScale_ARGB16U not loaded")
	}
	return _vImageScale_ARGB16U(src, dest, tempBuffer, flags)
}


var _vImageScale_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_ARGB8888 scales an 8-bit-per-channel, 4-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_ARGB8888(_:_:_:_:)
func VImageScale_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_ARGB8888 == nil {
		panic("Accelerate: symbol vImageScale_ARGB8888 not loaded")
	}
	return _vImageScale_ARGB8888(src, dest, tempBuffer, flags)
}


var _vImageScale_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_ARGBFFFF scales a 32-bit-per-channel, 4-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_ARGBFFFF(_:_:_:_:)
func VImageScale_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageScale_ARGBFFFF not loaded")
	}
	return _vImageScale_ARGBFFFF(src, dest, tempBuffer, flags)
}


var _vImageScale_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_CbCr16F scales a floating-point 16-bit-per-channel, 2-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_CbCr16F(_:_:_:_:)
func VImageScale_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_CbCr16F == nil {
		panic("Accelerate: symbol vImageScale_CbCr16F not loaded")
	}
	return _vImageScale_CbCr16F(src, dest, tempBuffer, flags)
}


var _vImageScale_CbCr16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_CbCr16U scales an unsigned 16-bit-per-channel, 2-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_CbCr16U(_:_:_:_:)
func VImageScale_CbCr16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_CbCr16U == nil {
		panic("Accelerate: symbol vImageScale_CbCr16U not loaded")
	}
	return _vImageScale_CbCr16U(src, dest, tempBuffer, flags)
}


var _vImageScale_CbCr8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_CbCr8 scales an 8-bit-per-channel, 2-channel interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_CbCr8(_:_:_:_:)
func VImageScale_CbCr8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_CbCr8 == nil {
		panic("Accelerate: symbol vImageScale_CbCr8 not loaded")
	}
	return _vImageScale_CbCr8(src, dest, tempBuffer, flags)
}


var _vImageScale_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_Planar16F scales a floating-point 16-bit planar image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar16F(_:_:_:_:)
func VImageScale_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_Planar16F == nil {
		panic("Accelerate: symbol vImageScale_Planar16F not loaded")
	}
	return _vImageScale_Planar16F(src, dest, tempBuffer, flags)
}


var _vImageScale_Planar16S func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_Planar16S scales a signed 16-bit planar image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar16S(_:_:_:_:)
func VImageScale_Planar16S(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_Planar16S == nil {
		panic("Accelerate: symbol vImageScale_Planar16S not loaded")
	}
	return _vImageScale_Planar16S(src, dest, tempBuffer, flags)
}


var _vImageScale_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_Planar16U scales an unsigned 16-bit planar image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar16U(_:_:_:_:)
func VImageScale_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_Planar16U == nil {
		panic("Accelerate: symbol vImageScale_Planar16U not loaded")
	}
	return _vImageScale_Planar16U(src, dest, tempBuffer, flags)
}


var _vImageScale_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_Planar8 scales an 8-bit planar image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar8(_:_:_:_:)
func VImageScale_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_Planar8 == nil {
		panic("Accelerate: symbol vImageScale_Planar8 not loaded")
	}
	return _vImageScale_Planar8(src, dest, tempBuffer, flags)
}


var _vImageScale_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_PlanarF scales a 32-bit planar image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_PlanarF(_:_:_:_:)
func VImageScale_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_PlanarF == nil {
		panic("Accelerate: symbol vImageScale_PlanarF not loaded")
	}
	return _vImageScale_PlanarF(src, dest, tempBuffer, flags)
}


var _vImageScale_XRGB2101010W func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int

// VImageScale_XRGB2101010W scales a 2-bit alpha, 10-bit RGB interleaved image to fit a destination buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageScale_XRGB2101010W(_:_:_:_:)
func VImageScale_XRGB2101010W(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags uint32) int {
	if _vImageScale_XRGB2101010W == nil {
		panic("Accelerate: symbol vImageScale_XRGB2101010W not loaded")
	}
	return _vImageScale_XRGB2101010W(src, dest, tempBuffer, flags)
}


var _vImageSelectChannels_ARGB8888 func(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageSelectChannels_ARGB8888 overwrites the channels of an 8-bit-per-channel, 4-channel interleaved buffer with the specified channels of the corresponding pixels of a second buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSelectChannels_ARGB8888(_:_:_:_:_:)
func VImageSelectChannels_ARGB8888(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageSelectChannels_ARGB8888 == nil {
		panic("Accelerate: symbol vImageSelectChannels_ARGB8888 not loaded")
	}
	return _vImageSelectChannels_ARGB8888(newSrc, origSrc, dest, copyMask, flags)
}


var _vImageSelectChannels_ARGBFFFF func(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int

// VImageSelectChannels_ARGBFFFF overwrites the channels of a floating-point 32-bit-per-channel, 4-channel interleaved buffer with the specified channels of the corresponding pixels of a second buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSelectChannels_ARGBFFFF(_:_:_:_:_:)
func VImageSelectChannels_ARGBFFFF(newSrc unsafe.Pointer, origSrc unsafe.Pointer, dest unsafe.Pointer, copyMask uint8, flags uint32) int {
	if _vImageSelectChannels_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageSelectChannels_ARGBFFFF not loaded")
	}
	return _vImageSelectChannels_ARGBFFFF(newSrc, origSrc, dest, copyMask, flags)
}


var _vImageSepConvolve_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, bias float32, backgroundColor Pixel_8888, flags uint32) int

// VImageSepConvolve_ARGB8888 convolves an 8-bit-per-channel, 4-channel interleaved image by separate horizontal and vertical separable kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, bias float32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageSepConvolve_ARGB8888 == nil {
		panic("Accelerate: symbol vImageSepConvolve_ARGB8888 not loaded")
	}
	return _vImageSepConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, bias, backgroundColor, flags)
}


var _vImageSepConvolve_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16F, flags uint32) int

// VImageSepConvolve_Planar16F convolves a floating-point 16-bit planar image by separate horizontal and vertical separable kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_Planar16F(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16F, flags uint32) int {
	if _vImageSepConvolve_Planar16F == nil {
		panic("Accelerate: symbol vImageSepConvolve_Planar16F not loaded")
	}
	return _vImageSepConvolve_Planar16F(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, bias, backgroundColor, flags)
}


var _vImageSepConvolve_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16U, flags uint32) int

// VImageSepConvolve_Planar16U convolves an unsigned 16-bit planar image by separate horizontal and vertical separable kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_Planar16U(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16U, flags uint32) int {
	if _vImageSepConvolve_Planar16U == nil {
		panic("Accelerate: symbol vImageSepConvolve_Planar16U not loaded")
	}
	return _vImageSepConvolve_Planar16U(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, bias, backgroundColor, flags)
}


var _vImageSepConvolve_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16U, flags uint32) int

// VImageSepConvolve_Planar8 convolves an 8-bit planar image by separate horizontal and vertical separable kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_Planar8(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, bias float32, backgroundColor Pixel_16U, flags uint32) int {
	if _vImageSepConvolve_Planar8 == nil {
		panic("Accelerate: symbol vImageSepConvolve_Planar8 not loaded")
	}
	return _vImageSepConvolve_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, bias, backgroundColor, flags)
}


var _vImageSepConvolve_Planar8to16U func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, scale float32, bias float32, backgroundColor Pixel_8, flags uint32) int

// VImageSepConvolve_Planar8to16U convolves an 8-bit planar image by separate horizontal and vertical separable kernels, and writes the result to an unsigned 16-bit planar destination.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_Planar8to16U(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_Planar8to16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, scale float32, bias float32, backgroundColor Pixel_8, flags uint32) int {
	if _vImageSepConvolve_Planar8to16U == nil {
		panic("Accelerate: symbol vImageSepConvolve_Planar8to16U not loaded")
	}
	return _vImageSepConvolve_Planar8to16U(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, scale, bias, backgroundColor, flags)
}


var _vImageSepConvolve_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX *float32, kernelX_width uint32, kernelY *float32, kernelY_width uint32, bias float32, backgroundColor Pixel_F, flags uint32) int

// VImageSepConvolve_PlanarF convolves a floating-point 32-bit planar image by separate horizontal and vertical separable kernels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_PlanarF(_:_:_:_:_:_:_:_:_:_:_:_:)
func VImageSepConvolve_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernelX []float32, kernelX_width uint32, kernelY []float32, kernelY_width uint32, bias float32, backgroundColor Pixel_F, flags uint32) int {
	if _vImageSepConvolve_PlanarF == nil {
		panic("Accelerate: symbol vImageSepConvolve_PlanarF not loaded")
	}
	return _vImageSepConvolve_PlanarF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, unsafe.SliceData(kernelX), kernelX_width, unsafe.SliceData(kernelY), kernelY_width, bias, backgroundColor, flags)
}


var _vImageSymmetricPiecewiseGamma_Planar16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int

// VImageSymmetricPiecewiseGamma_Planar16Q12 applies a symmetric piecewise gamma function to transform a 16Q12 planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSymmetricPiecewiseGamma_Planar16Q12(_:_:_:_:_:_:_:)
func VImageSymmetricPiecewiseGamma_Planar16Q12(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary Pixel_16S, flags uint32) int {
	if _vImageSymmetricPiecewiseGamma_Planar16Q12 == nil {
		panic("Accelerate: symbol vImageSymmetricPiecewiseGamma_Planar16Q12 not loaded")
	}
	return _vImageSymmetricPiecewiseGamma_Planar16Q12(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImageSymmetricPiecewiseGamma_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int

// VImageSymmetricPiecewiseGamma_PlanarF applies a symmetric piecewise gamma function to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSymmetricPiecewiseGamma_PlanarF(_:_:_:_:_:_:_:)
func VImageSymmetricPiecewiseGamma_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, exponentialCoeffs unsafe.Pointer, gamma float32, linearCoeffs unsafe.Pointer, boundary float32, flags uint32) int {
	if _vImageSymmetricPiecewiseGamma_PlanarF == nil {
		panic("Accelerate: symbol vImageSymmetricPiecewiseGamma_PlanarF not loaded")
	}
	return _vImageSymmetricPiecewiseGamma_PlanarF(src, dest, exponentialCoeffs, gamma, linearCoeffs, boundary, flags)
}


var _vImageSymmetricPiecewisePolynomial_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, coefficients *float32, boundaries *float32, order uint32, log2segments uint32, flags uint32) int

// VImageSymmetricPiecewisePolynomial_PlanarF applies a set of symmetric piecewise polynomials to transform a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageSymmetricPiecewisePolynomial_PlanarF(_:_:_:_:_:_:_:)
func VImageSymmetricPiecewisePolynomial_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, coefficients []float32, boundaries []float32, order uint32, log2segments uint32, flags uint32) int {
	if _vImageSymmetricPiecewisePolynomial_PlanarF == nil {
		panic("Accelerate: symbol vImageSymmetricPiecewisePolynomial_PlanarF not loaded")
	}
	return _vImageSymmetricPiecewisePolynomial_PlanarF(src, dest, unsafe.SliceData(coefficients), unsafe.SliceData(boundaries), order, log2segments, flags)
}


var _vImageTableLookUp_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, alphaTable Pixel_8, redTable Pixel_8, greenTable Pixel_8, blueTable Pixel_8, flags uint32) int

// VImageTableLookUp_ARGB8888 uses a lookup table to transform an interleaved, four-channel 8-bit planar image to an interleaved, four-channel 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageTableLookUp_ARGB8888(_:_:_:_:_:_:_:)
func VImageTableLookUp_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, alphaTable Pixel_8, redTable Pixel_8, greenTable Pixel_8, blueTable Pixel_8, flags uint32) int {
	if _vImageTableLookUp_ARGB8888 == nil {
		panic("Accelerate: symbol vImageTableLookUp_ARGB8888 not loaded")
	}
	return _vImageTableLookUp_ARGB8888(src, dest, alphaTable, redTable, greenTable, blueTable, flags)
}


var _vImageTableLookUp_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_8, flags uint32) int

// VImageTableLookUp_Planar8 uses a lookup table to transform an 8-bit planar image to an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageTableLookUp_Planar8(_:_:_:_:)
func VImageTableLookUp_Planar8(src unsafe.Pointer, dest unsafe.Pointer, table Pixel_8, flags uint32) int {
	if _vImageTableLookUp_Planar8 == nil {
		panic("Accelerate: symbol vImageTableLookUp_Planar8 not loaded")
	}
	return _vImageTableLookUp_Planar8(src, dest, table, flags)
}


var _vImageTentConvolve_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8888, flags uint32) int

// VImageTentConvolve_ARGB8888 applies a tent filter to an 8-bit-per-channel, 4-channel interleaved source image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageTentConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageTentConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8888, flags uint32) int {
	if _vImageTentConvolve_ARGB8888 == nil {
		panic("Accelerate: symbol vImageTentConvolve_ARGB8888 not loaded")
	}
	return _vImageTentConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageTentConvolve_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8, flags uint32) int

// VImageTentConvolve_Planar8 applies a tent filter to an 8-bit planar source image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageTentConvolve_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageTentConvolve_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, kernel_height uint32, kernel_width uint32, backgroundColor Pixel_8, flags uint32) int {
	if _vImageTentConvolve_Planar8 == nil {
		panic("Accelerate: symbol vImageTentConvolve_Planar8 not loaded")
	}
	return _vImageTentConvolve_Planar8(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, backgroundColor, flags)
}


var _vImageUnpremultiplyData_ARGB16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_ARGB16Q12 transforms a fixed-point 16-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGB16Q12(_:_:_:)
func VImageUnpremultiplyData_ARGB16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_ARGB16Q12 == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_ARGB16Q12 not loaded")
	}
	return _vImageUnpremultiplyData_ARGB16Q12(src, dest, flags)
}


var _vImageUnpremultiplyData_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_ARGB16U transforms an unsigned 16-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGB16U(_:_:_:)
func VImageUnpremultiplyData_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_ARGB16U == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_ARGB16U not loaded")
	}
	return _vImageUnpremultiplyData_ARGB16U(src, dest, flags)
}


var _vImageUnpremultiplyData_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_ARGB8888 transforms an 8-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGB8888(_:_:_:)
func VImageUnpremultiplyData_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_ARGB8888 == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_ARGB8888 not loaded")
	}
	return _vImageUnpremultiplyData_ARGB8888(src, dest, flags)
}


var _vImageUnpremultiplyData_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_ARGBFFFF transforms a 32-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGBFFFF(_:_:_:)
func VImageUnpremultiplyData_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_ARGBFFFF not loaded")
	}
	return _vImageUnpremultiplyData_ARGBFFFF(src, dest, flags)
}


var _vImageUnpremultiplyData_Planar8 func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_Planar8 transforms an 8-bit planar buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_Planar8(_:_:_:_:)
func VImageUnpremultiplyData_Planar8(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_Planar8 == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_Planar8 not loaded")
	}
	return _vImageUnpremultiplyData_Planar8(src, alpha, dest, flags)
}


var _vImageUnpremultiplyData_PlanarF func(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_PlanarF transforms a 32-bit planar buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_PlanarF(_:_:_:_:)
func VImageUnpremultiplyData_PlanarF(src unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_PlanarF == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_PlanarF not loaded")
	}
	return _vImageUnpremultiplyData_PlanarF(src, alpha, dest, flags)
}


var _vImageUnpremultiplyData_RGBA16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_RGBA16F transforms a floating-point 16-bit-per-channel, 4-channel RGBA buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_RGBA16F(_:_:_:)
func VImageUnpremultiplyData_RGBA16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_RGBA16F == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_RGBA16F not loaded")
	}
	return _vImageUnpremultiplyData_RGBA16F(src, dest, flags)
}


var _vImageUnpremultiplyData_RGBA16Q12 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_RGBA16Q12 transforms a fixed-point 16-bit-per-channel, 4-channel RGBA buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_RGBA16Q12(_:_:_:)
func VImageUnpremultiplyData_RGBA16Q12(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_RGBA16Q12 == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_RGBA16Q12 not loaded")
	}
	return _vImageUnpremultiplyData_RGBA16Q12(src, dest, flags)
}


var _vImageUnpremultiplyData_RGBA16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_RGBA16U transforms an unsigned 16-bit-per-channel, 4-channel RGBA buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_RGBA16U(_:_:_:)
func VImageUnpremultiplyData_RGBA16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_RGBA16U == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_RGBA16U not loaded")
	}
	return _vImageUnpremultiplyData_RGBA16U(src, dest, flags)
}


var _vImageUnpremultiplyData_RGBA8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_RGBA8888 transforms an 8-bit-per-channel, 4-channel RGBA buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_RGBA8888(_:_:_:)
func VImageUnpremultiplyData_RGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_RGBA8888 == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_RGBA8888 not loaded")
	}
	return _vImageUnpremultiplyData_RGBA8888(src, dest, flags)
}


var _vImageUnpremultiplyData_RGBAFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageUnpremultiplyData_RGBAFFFF transforms a 32-bit-per-channel, 4-channel RGBA buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_RGBAFFFF(_:_:_:)
func VImageUnpremultiplyData_RGBAFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageUnpremultiplyData_RGBAFFFF == nil {
		panic("Accelerate: symbol vImageUnpremultiplyData_RGBAFFFF not loaded")
	}
	return _vImageUnpremultiplyData_RGBAFFFF(src, dest, flags)
}


var _vImageVerticalReflect_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_ARGB16F reflects a floating-point 16-bit-per-channel, 4-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_ARGB16F(_:_:_:)
func VImageVerticalReflect_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_ARGB16F == nil {
		panic("Accelerate: symbol vImageVerticalReflect_ARGB16F not loaded")
	}
	return _vImageVerticalReflect_ARGB16F(src, dest, flags)
}


var _vImageVerticalReflect_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_ARGB16S reflects a signed 16-bit-per-channel, 4-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_ARGB16S(_:_:_:)
func VImageVerticalReflect_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_ARGB16S == nil {
		panic("Accelerate: symbol vImageVerticalReflect_ARGB16S not loaded")
	}
	return _vImageVerticalReflect_ARGB16S(src, dest, flags)
}


var _vImageVerticalReflect_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_ARGB16U reflects an unsigned 16-bit-per-channel, 4-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_ARGB16U(_:_:_:)
func VImageVerticalReflect_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_ARGB16U == nil {
		panic("Accelerate: symbol vImageVerticalReflect_ARGB16U not loaded")
	}
	return _vImageVerticalReflect_ARGB16U(src, dest, flags)
}


var _vImageVerticalReflect_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_ARGB8888 reflects an 8-bit-per-channel, 4-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_ARGB8888(_:_:_:)
func VImageVerticalReflect_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_ARGB8888 == nil {
		panic("Accelerate: symbol vImageVerticalReflect_ARGB8888 not loaded")
	}
	return _vImageVerticalReflect_ARGB8888(src, dest, flags)
}


var _vImageVerticalReflect_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_ARGBFFFF reflects a 32-bit-per-channel, 4-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_ARGBFFFF(_:_:_:)
func VImageVerticalReflect_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageVerticalReflect_ARGBFFFF not loaded")
	}
	return _vImageVerticalReflect_ARGBFFFF(src, dest, flags)
}


var _vImageVerticalReflect_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_CbCr16F reflects a floating-point 16-bit-per-channel, 2-channel interleaved image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_CbCr16F(_:_:_:)
func VImageVerticalReflect_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_CbCr16F == nil {
		panic("Accelerate: symbol vImageVerticalReflect_CbCr16F not loaded")
	}
	return _vImageVerticalReflect_CbCr16F(src, dest, flags)
}


var _vImageVerticalReflect_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_Planar16F reflects a floating-point 16-bit planar image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_Planar16F(_:_:_:)
func VImageVerticalReflect_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_Planar16F == nil {
		panic("Accelerate: symbol vImageVerticalReflect_Planar16F not loaded")
	}
	return _vImageVerticalReflect_Planar16F(src, dest, flags)
}


var _vImageVerticalReflect_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_Planar16U reflects an unsigned 16-bit planar image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_Planar16U(_:_:_:)
func VImageVerticalReflect_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_Planar16U == nil {
		panic("Accelerate: symbol vImageVerticalReflect_Planar16U not loaded")
	}
	return _vImageVerticalReflect_Planar16U(src, dest, flags)
}


var _vImageVerticalReflect_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_Planar8 reflects an 8-bit planar image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_Planar8(_:_:_:)
func VImageVerticalReflect_Planar8(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_Planar8 == nil {
		panic("Accelerate: symbol vImageVerticalReflect_Planar8 not loaded")
	}
	return _vImageVerticalReflect_Planar8(src, dest, flags)
}


var _vImageVerticalReflect_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int

// VImageVerticalReflect_PlanarF reflects a 32-bit planar image vertically across the center horizontal line.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalReflect_PlanarF(_:_:_:)
func VImageVerticalReflect_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, flags uint32) int {
	if _vImageVerticalReflect_PlanarF == nil {
		panic("Accelerate: symbol vImageVerticalReflect_PlanarF not loaded")
	}
	return _vImageVerticalReflect_PlanarF(src, dest, flags)
}


var _vImageVerticalShearD_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int

// VImageVerticalShearD_ARGB16F performs a double-precision vertical shear on a region of interest within a floating-point 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_ARGB16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageVerticalShearD_ARGB16F == nil {
		panic("Accelerate: symbol vImageVerticalShearD_ARGB16F not loaded")
	}
	return _vImageVerticalShearD_ARGB16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int

// VImageVerticalShearD_ARGB16S performs a double-precision vertical shear on a region of interest within a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_ARGB16S(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageVerticalShearD_ARGB16S == nil {
		panic("Accelerate: symbol vImageVerticalShearD_ARGB16S not loaded")
	}
	return _vImageVerticalShearD_ARGB16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int

// VImageVerticalShearD_ARGB16U performs a double-precision vertical shear on a region of interest within an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_ARGB16U(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageVerticalShearD_ARGB16U == nil {
		panic("Accelerate: symbol vImageVerticalShearD_ARGB16U not loaded")
	}
	return _vImageVerticalShearD_ARGB16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int

// VImageVerticalShearD_ARGB8888 performs a double-precision vertical shear on a region of interest within an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int {
	if _vImageVerticalShearD_ARGB8888 == nil {
		panic("Accelerate: symbol vImageVerticalShearD_ARGB8888 not loaded")
	}
	return _vImageVerticalShearD_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int

// VImageVerticalShearD_ARGBFFFF performs a double-precision vertical shear on a region of interest within a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int {
	if _vImageVerticalShearD_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageVerticalShearD_ARGBFFFF not loaded")
	}
	return _vImageVerticalShearD_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int

// VImageVerticalShearD_CbCr16F performs a double-precision vertical shear on a region of interest within a floating-point 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_CbCr16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int {
	if _vImageVerticalShearD_CbCr16F == nil {
		panic("Accelerate: symbol vImageVerticalShearD_CbCr16F not loaded")
	}
	return _vImageVerticalShearD_CbCr16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_CbCr16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int

// VImageVerticalShearD_CbCr16S.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_CbCr16S(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_CbCr16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int {
	if _vImageVerticalShearD_CbCr16S == nil {
		panic("Accelerate: symbol vImageVerticalShearD_CbCr16S not loaded")
	}
	return _vImageVerticalShearD_CbCr16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_CbCr16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int

// VImageVerticalShearD_CbCr16U.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_CbCr16U(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_CbCr16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int {
	if _vImageVerticalShearD_CbCr16U == nil {
		panic("Accelerate: symbol vImageVerticalShearD_CbCr16U not loaded")
	}
	return _vImageVerticalShearD_CbCr16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int

// VImageVerticalShearD_Planar16F performs a double-precision vertical shear on a region of interest within a floating-point 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_Planar16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int {
	if _vImageVerticalShearD_Planar16F == nil {
		panic("Accelerate: symbol vImageVerticalShearD_Planar16F not loaded")
	}
	return _vImageVerticalShearD_Planar16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8, flags uint32) int

// VImageVerticalShearD_Planar8 performs a double-precision vertical shear on a region of interest within an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_8, flags uint32) int {
	if _vImageVerticalShearD_Planar8 == nil {
		panic("Accelerate: symbol vImageVerticalShearD_Planar8 not loaded")
	}
	return _vImageVerticalShearD_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShearD_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_F, flags uint32) int

// VImageVerticalShearD_PlanarF performs a double-precision vertical shear on a region of interest within a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShearD_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShearD_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float64, shearSlope float64, filter ResamplingFilter, backColor Pixel_F, flags uint32) int {
	if _vImageVerticalShearD_PlanarF == nil {
		panic("Accelerate: symbol vImageVerticalShearD_PlanarF not loaded")
	}
	return _vImageVerticalShearD_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_ARGB16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int

// VImageVerticalShear_ARGB16F performs a single-precision vertical shear on a region of interest within a floating-point 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGB16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16F, flags uint32) int {
	if _vImageVerticalShear_ARGB16F == nil {
		panic("Accelerate: symbol vImageVerticalShear_ARGB16F not loaded")
	}
	return _vImageVerticalShear_ARGB16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_ARGB16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int

// VImageVerticalShear_ARGB16S performs a single-precision vertical shear on a region of interest within a signed 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGB16S(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_ARGB16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16S, flags uint32) int {
	if _vImageVerticalShear_ARGB16S == nil {
		panic("Accelerate: symbol vImageVerticalShear_ARGB16S not loaded")
	}
	return _vImageVerticalShear_ARGB16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_ARGB16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int

// VImageVerticalShear_ARGB16U performs a single-precision vertical shear on a region of interest within an unsigned 16-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGB16U(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_ARGB_16U, flags uint32) int {
	if _vImageVerticalShear_ARGB16U == nil {
		panic("Accelerate: symbol vImageVerticalShear_ARGB16U not loaded")
	}
	return _vImageVerticalShear_ARGB16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_ARGB8888 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int

// VImageVerticalShear_ARGB8888 performs a single-precision vertical shear on a region of interest within an 8-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGB8888(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8888, flags uint32) int {
	if _vImageVerticalShear_ARGB8888 == nil {
		panic("Accelerate: symbol vImageVerticalShear_ARGB8888 not loaded")
	}
	return _vImageVerticalShear_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_ARGBFFFF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int

// VImageVerticalShear_ARGBFFFF performs a single-precision vertical shear on a region of interest within a 32-bit-per-channel, 4-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_FFFF, flags uint32) int {
	if _vImageVerticalShear_ARGBFFFF == nil {
		panic("Accelerate: symbol vImageVerticalShear_ARGBFFFF not loaded")
	}
	return _vImageVerticalShear_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_CbCr16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int

// VImageVerticalShear_CbCr16F performs a single-precision vertical shear on a region of interest within a floating-point 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_CbCr16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_CbCr16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F16F, flags uint32) int {
	if _vImageVerticalShear_CbCr16F == nil {
		panic("Accelerate: symbol vImageVerticalShear_CbCr16F not loaded")
	}
	return _vImageVerticalShear_CbCr16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_CbCr16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int

// VImageVerticalShear_CbCr16S.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_CbCr16S(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_CbCr16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S16S, flags uint32) int {
	if _vImageVerticalShear_CbCr16S == nil {
		panic("Accelerate: symbol vImageVerticalShear_CbCr16S not loaded")
	}
	return _vImageVerticalShear_CbCr16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_CbCr16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int

// VImageVerticalShear_CbCr16U performs a single-precision vertical shear on a region of interest within an unsigned 16-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_CbCr16U(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_CbCr16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U16U, flags uint32) int {
	if _vImageVerticalShear_CbCr16U == nil {
		panic("Accelerate: symbol vImageVerticalShear_CbCr16U not loaded")
	}
	return _vImageVerticalShear_CbCr16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_CbCr8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_88, flags uint32) int

// VImageVerticalShear_CbCr8 performs a single-precision vertical shear on a region of interest within an 8-bit-per-channel, 2-channel interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_CbCr8(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_CbCr8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_88, flags uint32) int {
	if _vImageVerticalShear_CbCr8 == nil {
		panic("Accelerate: symbol vImageVerticalShear_CbCr8 not loaded")
	}
	return _vImageVerticalShear_CbCr8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_Planar16F func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int

// VImageVerticalShear_Planar16F performs a single-precision vertical shear on a region of interest within a floating-point 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_Planar16F(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16F, flags uint32) int {
	if _vImageVerticalShear_Planar16F == nil {
		panic("Accelerate: symbol vImageVerticalShear_Planar16F not loaded")
	}
	return _vImageVerticalShear_Planar16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_Planar16S func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S, flags uint32) int

// VImageVerticalShear_Planar16S performs a single-precision vertical shear on a region of interest within a signed 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_Planar16S(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_Planar16S(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16S, flags uint32) int {
	if _vImageVerticalShear_Planar16S == nil {
		panic("Accelerate: symbol vImageVerticalShear_Planar16S not loaded")
	}
	return _vImageVerticalShear_Planar16S(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_Planar16U func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U, flags uint32) int

// VImageVerticalShear_Planar16U performs a single-precision vertical shear on a region of interest within an unsigned 16-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_Planar16U(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_16U, flags uint32) int {
	if _vImageVerticalShear_Planar16U == nil {
		panic("Accelerate: symbol vImageVerticalShear_Planar16U not loaded")
	}
	return _vImageVerticalShear_Planar16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_Planar8 func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8, flags uint32) int

// VImageVerticalShear_Planar8 performs a single-precision vertical shear on a region of interest within an 8-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_Planar8(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_Planar8(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_8, flags uint32) int {
	if _vImageVerticalShear_Planar8 == nil {
		panic("Accelerate: symbol vImageVerticalShear_Planar8 not loaded")
	}
	return _vImageVerticalShear_Planar8(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_PlanarF func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_F, flags uint32) int

// VImageVerticalShear_PlanarF performs a single-precision vertical shear on a region of interest within a 32-bit planar image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_PlanarF(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_F, flags uint32) int {
	if _vImageVerticalShear_PlanarF == nil {
		panic("Accelerate: symbol vImageVerticalShear_PlanarF not loaded")
	}
	return _vImageVerticalShear_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vImageVerticalShear_XRGB2101010W func(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_32U, flags uint32) int

// VImageVerticalShear_XRGB2101010W performs a single-precision vertical shear on a region of interest within a 2-bit alpha, 10-bit RGB interleaved image.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_XRGB2101010W(_:_:_:_:_:_:_:_:_:)
func VImageVerticalShear_XRGB2101010W(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X uint, srcOffsetToROI_Y uint, yTranslate float32, shearSlope float32, filter ResamplingFilter, backColor Pixel_32U, flags uint32) int {
	if _vImageVerticalShear_XRGB2101010W == nil {
		panic("Accelerate: symbol vImageVerticalShear_XRGB2101010W not loaded")
	}
	return _vImageVerticalShear_XRGB2101010W(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
}


var _vL1024Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VL1024Rotate 1024-bit left rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vL1024Rotate(_:_:_:)
func VL1024Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vL1024Rotate == nil {
		panic("Accelerate: symbol vL1024Rotate not loaded")
	}
	_vL1024Rotate(a, rotateAmount, result)
}



var _vL256Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VL256Rotate 256-bit left rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vL256Rotate(_:_:_:)
func VL256Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vL256Rotate == nil {
		panic("Accelerate: symbol vL256Rotate not loaded")
	}
	_vL256Rotate(a, rotateAmount, result)
}


var _vL512Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VL512Rotate 512-bit left rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vL512Rotate(_:_:_:)
func VL512Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vL512Rotate == nil {
		panic("Accelerate: symbol vL512Rotate not loaded")
	}
	_vL512Rotate(a, rotateAmount, result)
}




var _vLL1024Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLL1024Shift 1024-bit logical left shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vLL1024Shift(_:_:_:)
func VLL1024Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLL1024Shift == nil {
		panic("Accelerate: symbol vLL1024Shift not loaded")
	}
	_vLL1024Shift(a, shiftAmount, result)
}



var _vLL256Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLL256Shift 256-bit logical left shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vLL256Shift(_:_:_:)
func VLL256Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLL256Shift == nil {
		panic("Accelerate: symbol vLL256Shift not loaded")
	}
	_vLL256Shift(a, shiftAmount, result)
}


var _vLL512Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLL512Shift 512-bit logical left shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vLL512Shift(_:_:_:)
func VLL512Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLL512Shift == nil {
		panic("Accelerate: symbol vLL512Shift not loaded")
	}
	_vLL512Shift(a, shiftAmount, result)
}




var _vLR1024Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLR1024Shift 1024-bit logical right shift .
//
// See: https://developer.apple.com/documentation/Accelerate/vLR1024Shift(_:_:_:)
func VLR1024Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLR1024Shift == nil {
		panic("Accelerate: symbol vLR1024Shift not loaded")
	}
	_vLR1024Shift(a, shiftAmount, result)
}



var _vLR256Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLR256Shift 256-bit logical right shift.
//
// See: https://developer.apple.com/documentation/Accelerate/vLR256Shift(_:_:_:)
func VLR256Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLR256Shift == nil {
		panic("Accelerate: symbol vLR256Shift not loaded")
	}
	_vLR256Shift(a, shiftAmount, result)
}


var _vLR512Shift func(a uintptr, shiftAmount uint32, result uintptr)

// VLR512Shift 512-bit logical right shift .
//
// See: https://developer.apple.com/documentation/Accelerate/vLR512Shift(_:_:_:)
func VLR512Shift(a uintptr, shiftAmount uint32, result uintptr) {
	if _vLR512Shift == nil {
		panic("Accelerate: symbol vLR512Shift not loaded")
	}
	_vLR512Shift(a, shiftAmount, result)
}




var _vR1024Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VR1024Rotate 1024-bit right rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vR1024Rotate(_:_:_:)
func VR1024Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vR1024Rotate == nil {
		panic("Accelerate: symbol vR1024Rotate not loaded")
	}
	_vR1024Rotate(a, rotateAmount, result)
}



var _vR256Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VR256Rotate 256-bit right rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vR256Rotate(_:_:_:)
func VR256Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vR256Rotate == nil {
		panic("Accelerate: symbol vR256Rotate not loaded")
	}
	_vR256Rotate(a, rotateAmount, result)
}


var _vR512Rotate func(a uintptr, rotateAmount uint32, result uintptr)

// VR512Rotate 512-bit right rotate.
//
// See: https://developer.apple.com/documentation/Accelerate/vR512Rotate(_:_:_:)
func VR512Rotate(a uintptr, rotateAmount uint32, result uintptr) {
	if _vR512Rotate == nil {
		panic("Accelerate: symbol vR512Rotate not loaded")
	}
	_vR512Rotate(a, rotateAmount, result)
}




var _vS1024Add func(a uintptr, b uintptr, result uintptr)

// VS1024Add signed 1024-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024Add(_:_:_:)
func VS1024Add(a uintptr, b uintptr, result uintptr) {
	if _vS1024Add == nil {
		panic("Accelerate: symbol vS1024Add not loaded")
	}
	_vS1024Add(a, b, result)
}


var _vS1024AddS func(a uintptr, b uintptr, result uintptr)

// VS1024AddS signed 1024-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024AddS(_:_:_:)
func VS1024AddS(a uintptr, b uintptr, result uintptr) {
	if _vS1024AddS == nil {
		panic("Accelerate: symbol vS1024AddS not loaded")
	}
	_vS1024AddS(a, b, result)
}


var _vS1024Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VS1024Divide signed 1024-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024Divide(_:_:_:_:)
func VS1024Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vS1024Divide == nil {
		panic("Accelerate: symbol vS1024Divide not loaded")
	}
	_vS1024Divide(numerator, divisor, result, remainder)
}


var _vS1024HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VS1024HalfMultiply signed 1024-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024HalfMultiply(_:_:_:)
func VS1024HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS1024HalfMultiply == nil {
		panic("Accelerate: symbol vS1024HalfMultiply not loaded")
	}
	_vS1024HalfMultiply(a, b, result)
}


var _vS1024Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VS1024Mod signed 256-bit Mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024Mod(_:_:_:)
func VS1024Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vS1024Mod == nil {
		panic("Accelerate: symbol vS1024Mod not loaded")
	}
	_vS1024Mod(numerator, divisor, remainder)
}


var _vS1024Neg func(a uintptr, result uintptr)

// VS1024Neg signed 1024-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024Neg(_:_:)
func VS1024Neg(a uintptr, result uintptr) {
	if _vS1024Neg == nil {
		panic("Accelerate: symbol vS1024Neg not loaded")
	}
	_vS1024Neg(a, result)
}


var _vS1024Sub func(a uintptr, b uintptr, result uintptr)

// VS1024Sub signed 1024-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024Sub(_:_:_:)
func VS1024Sub(a uintptr, b uintptr, result uintptr) {
	if _vS1024Sub == nil {
		panic("Accelerate: symbol vS1024Sub not loaded")
	}
	_vS1024Sub(a, b, result)
}


var _vS1024SubS func(a uintptr, b uintptr, result uintptr)

// VS1024SubS signed 1024-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS1024SubS(_:_:_:)
func VS1024SubS(a uintptr, b uintptr, result uintptr) {
	if _vS1024SubS == nil {
		panic("Accelerate: symbol vS1024SubS not loaded")
	}
	_vS1024SubS(a, b, result)
}





var _vS128FullMultiply func(a uintptr, b uintptr, result uintptr)

// VS128FullMultiply signed 128-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS128FullMultiply(_:_:_:)
func VS128FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS128FullMultiply == nil {
		panic("Accelerate: symbol vS128FullMultiply not loaded")
	}
	_vS128FullMultiply(a, b, result)
}








var _vS256Add func(a uintptr, b uintptr, result uintptr)

// VS256Add signed 256-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS256Add(_:_:_:)
func VS256Add(a uintptr, b uintptr, result uintptr) {
	if _vS256Add == nil {
		panic("Accelerate: symbol vS256Add not loaded")
	}
	_vS256Add(a, b, result)
}


var _vS256AddS func(a uintptr, b uintptr, result uintptr)

// VS256AddS signed 256-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS256AddS(_:_:_:)
func VS256AddS(a uintptr, b uintptr, result uintptr) {
	if _vS256AddS == nil {
		panic("Accelerate: symbol vS256AddS not loaded")
	}
	_vS256AddS(a, b, result)
}


var _vS256Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VS256Divide computes the signed 256-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vS256Divide(_:_:_:_:)
func VS256Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vS256Divide == nil {
		panic("Accelerate: symbol vS256Divide not loaded")
	}
	_vS256Divide(numerator, divisor, result, remainder)
}


var _vS256FullMultiply func(a uintptr, b uintptr, result uintptr)

// VS256FullMultiply signed 256-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS256FullMultiply(_:_:_:)
func VS256FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS256FullMultiply == nil {
		panic("Accelerate: symbol vS256FullMultiply not loaded")
	}
	_vS256FullMultiply(a, b, result)
}


var _vS256HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VS256HalfMultiply signed 256-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS256HalfMultiply(_:_:_:)
func VS256HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS256HalfMultiply == nil {
		panic("Accelerate: symbol vS256HalfMultiply not loaded")
	}
	_vS256HalfMultiply(a, b, result)
}


var _vS256Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VS256Mod signed 256-bit mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vS256Mod(_:_:_:)
func VS256Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vS256Mod == nil {
		panic("Accelerate: symbol vS256Mod not loaded")
	}
	_vS256Mod(numerator, divisor, remainder)
}


var _vS256Neg func(a uintptr, result uintptr)

// VS256Neg signed 256-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vS256Neg(_:_:)
func VS256Neg(a uintptr, result uintptr) {
	if _vS256Neg == nil {
		panic("Accelerate: symbol vS256Neg not loaded")
	}
	_vS256Neg(a, result)
}


var _vS256Sub func(a uintptr, b uintptr, result uintptr)

// VS256Sub signed 256-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS256Sub(_:_:_:)
func VS256Sub(a uintptr, b uintptr, result uintptr) {
	if _vS256Sub == nil {
		panic("Accelerate: symbol vS256Sub not loaded")
	}
	_vS256Sub(a, b, result)
}


var _vS256SubS func(a uintptr, b uintptr, result uintptr)

// VS256SubS signed 256-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS256SubS(_:_:_:)
func VS256SubS(a uintptr, b uintptr, result uintptr) {
	if _vS256SubS == nil {
		panic("Accelerate: symbol vS256SubS not loaded")
	}
	_vS256SubS(a, b, result)
}






var _vS512Add func(a uintptr, b uintptr, result uintptr)

// VS512Add signed 512-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS512Add(_:_:_:)
func VS512Add(a uintptr, b uintptr, result uintptr) {
	if _vS512Add == nil {
		panic("Accelerate: symbol vS512Add not loaded")
	}
	_vS512Add(a, b, result)
}


var _vS512AddS func(a uintptr, b uintptr, result uintptr)

// VS512AddS signed 512-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS512AddS(_:_:_:)
func VS512AddS(a uintptr, b uintptr, result uintptr) {
	if _vS512AddS == nil {
		panic("Accelerate: symbol vS512AddS not loaded")
	}
	_vS512AddS(a, b, result)
}


var _vS512Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VS512Divide signed 512-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vS512Divide(_:_:_:_:)
func VS512Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vS512Divide == nil {
		panic("Accelerate: symbol vS512Divide not loaded")
	}
	_vS512Divide(numerator, divisor, result, remainder)
}


var _vS512FullMultiply func(a uintptr, b uintptr, result uintptr)

// VS512FullMultiply signed 512-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS512FullMultiply(_:_:_:)
func VS512FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS512FullMultiply == nil {
		panic("Accelerate: symbol vS512FullMultiply not loaded")
	}
	_vS512FullMultiply(a, b, result)
}


var _vS512HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VS512HalfMultiply signed 512-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vS512HalfMultiply(_:_:_:)
func VS512HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vS512HalfMultiply == nil {
		panic("Accelerate: symbol vS512HalfMultiply not loaded")
	}
	_vS512HalfMultiply(a, b, result)
}


var _vS512Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VS512Mod signed 512-bit mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vS512Mod(_:_:_:)
func VS512Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vS512Mod == nil {
		panic("Accelerate: symbol vS512Mod not loaded")
	}
	_vS512Mod(numerator, divisor, remainder)
}


var _vS512Neg func(a uintptr, result uintptr)

// VS512Neg signed 512-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vS512Neg(_:_:)
func VS512Neg(a uintptr, result uintptr) {
	if _vS512Neg == nil {
		panic("Accelerate: symbol vS512Neg not loaded")
	}
	_vS512Neg(a, result)
}


var _vS512Sub func(a uintptr, b uintptr, result uintptr)

// VS512Sub signed 512-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vS512Sub(_:_:_:)
func VS512Sub(a uintptr, b uintptr, result uintptr) {
	if _vS512Sub == nil {
		panic("Accelerate: symbol vS512Sub not loaded")
	}
	_vS512Sub(a, b, result)
}


var _vS512SubS func(a uintptr, b uintptr, result uintptr)

// VS512SubS signed 512-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vS512SubS(_:_:_:)
func VS512SubS(a uintptr, b uintptr, result uintptr) {
	if _vS512SubS == nil {
		panic("Accelerate: symbol vS512SubS not loaded")
	}
	_vS512SubS(a, b, result)
}













var _vU1024Add func(a uintptr, b uintptr, result uintptr)

// VU1024Add unsigned 1024-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024Add(_:_:_:)
func VU1024Add(a uintptr, b uintptr, result uintptr) {
	if _vU1024Add == nil {
		panic("Accelerate: symbol vU1024Add not loaded")
	}
	_vU1024Add(a, b, result)
}


var _vU1024AddS func(a uintptr, b uintptr, result uintptr)

// VU1024AddS unsigned 1024-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024AddS(_:_:_:)
func VU1024AddS(a uintptr, b uintptr, result uintptr) {
	if _vU1024AddS == nil {
		panic("Accelerate: symbol vU1024AddS not loaded")
	}
	_vU1024AddS(a, b, result)
}


var _vU1024Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VU1024Divide unsigned 1024-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024Divide(_:_:_:_:)
func VU1024Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vU1024Divide == nil {
		panic("Accelerate: symbol vU1024Divide not loaded")
	}
	_vU1024Divide(numerator, divisor, result, remainder)
}


var _vU1024HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VU1024HalfMultiply unsigned 1024-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024HalfMultiply(_:_:_:)
func VU1024HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU1024HalfMultiply == nil {
		panic("Accelerate: symbol vU1024HalfMultiply not loaded")
	}
	_vU1024HalfMultiply(a, b, result)
}


var _vU1024Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VU1024Mod unsigned 1024-bit mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024Mod(_:_:_:)
func VU1024Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vU1024Mod == nil {
		panic("Accelerate: symbol vU1024Mod not loaded")
	}
	_vU1024Mod(numerator, divisor, remainder)
}


var _vU1024Neg func(a uintptr, result uintptr)

// VU1024Neg unsigned 1024-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024Neg(_:_:)
func VU1024Neg(a uintptr, result uintptr) {
	if _vU1024Neg == nil {
		panic("Accelerate: symbol vU1024Neg not loaded")
	}
	_vU1024Neg(a, result)
}


var _vU1024Sub func(a uintptr, b uintptr, result uintptr)

// VU1024Sub unsigned 1024-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024Sub(_:_:_:)
func VU1024Sub(a uintptr, b uintptr, result uintptr) {
	if _vU1024Sub == nil {
		panic("Accelerate: symbol vU1024Sub not loaded")
	}
	_vU1024Sub(a, b, result)
}


var _vU1024SubS func(a uintptr, b uintptr, result uintptr)

// VU1024SubS unsigned 1024-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU1024SubS(_:_:_:)
func VU1024SubS(a uintptr, b uintptr, result uintptr) {
	if _vU1024SubS == nil {
		panic("Accelerate: symbol vU1024SubS not loaded")
	}
	_vU1024SubS(a, b, result)
}





var _vU128FullMultiply func(a uintptr, b uintptr, result uintptr)

// VU128FullMultiply unsigned 128-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU128FullMultiply(_:_:_:)
func VU128FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU128FullMultiply == nil {
		panic("Accelerate: symbol vU128FullMultiply not loaded")
	}
	_vU128FullMultiply(a, b, result)
}








var _vU256Add func(a uintptr, b uintptr, result uintptr)

// VU256Add unsigned 256-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU256Add(_:_:_:)
func VU256Add(a uintptr, b uintptr, result uintptr) {
	if _vU256Add == nil {
		panic("Accelerate: symbol vU256Add not loaded")
	}
	_vU256Add(a, b, result)
}


var _vU256AddS func(a uintptr, b uintptr, result uintptr)

// VU256AddS unsigned 256-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU256AddS(_:_:_:)
func VU256AddS(a uintptr, b uintptr, result uintptr) {
	if _vU256AddS == nil {
		panic("Accelerate: symbol vU256AddS not loaded")
	}
	_vU256AddS(a, b, result)
}


var _vU256Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VU256Divide unsigned 256-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vU256Divide(_:_:_:_:)
func VU256Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vU256Divide == nil {
		panic("Accelerate: symbol vU256Divide not loaded")
	}
	_vU256Divide(numerator, divisor, result, remainder)
}


var _vU256FullMultiply func(a uintptr, b uintptr, result uintptr)

// VU256FullMultiply unsigned 256-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU256FullMultiply(_:_:_:)
func VU256FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU256FullMultiply == nil {
		panic("Accelerate: symbol vU256FullMultiply not loaded")
	}
	_vU256FullMultiply(a, b, result)
}


var _vU256HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VU256HalfMultiply unsigned 256-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU256HalfMultiply(_:_:_:)
func VU256HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU256HalfMultiply == nil {
		panic("Accelerate: symbol vU256HalfMultiply not loaded")
	}
	_vU256HalfMultiply(a, b, result)
}


var _vU256Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VU256Mod unsigned 256-bit mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vU256Mod(_:_:_:)
func VU256Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vU256Mod == nil {
		panic("Accelerate: symbol vU256Mod not loaded")
	}
	_vU256Mod(numerator, divisor, remainder)
}


var _vU256Neg func(a uintptr, result uintptr)

// VU256Neg unsigned 256-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vU256Neg(_:_:)
func VU256Neg(a uintptr, result uintptr) {
	if _vU256Neg == nil {
		panic("Accelerate: symbol vU256Neg not loaded")
	}
	_vU256Neg(a, result)
}


var _vU256Sub func(a uintptr, b uintptr, result uintptr)

// VU256Sub unsigned 256-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU256Sub(_:_:_:)
func VU256Sub(a uintptr, b uintptr, result uintptr) {
	if _vU256Sub == nil {
		panic("Accelerate: symbol vU256Sub not loaded")
	}
	_vU256Sub(a, b, result)
}


var _vU256SubS func(a uintptr, b uintptr, result uintptr)

// VU256SubS unsigned 256-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU256SubS(_:_:_:)
func VU256SubS(a uintptr, b uintptr, result uintptr) {
	if _vU256SubS == nil {
		panic("Accelerate: symbol vU256SubS not loaded")
	}
	_vU256SubS(a, b, result)
}






var _vU512Add func(a uintptr, b uintptr, result uintptr)

// VU512Add unsigned 512-bit addition (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU512Add(_:_:_:)
func VU512Add(a uintptr, b uintptr, result uintptr) {
	if _vU512Add == nil {
		panic("Accelerate: symbol vU512Add not loaded")
	}
	_vU512Add(a, b, result)
}


var _vU512AddS func(a uintptr, b uintptr, result uintptr)

// VU512AddS unsigned 512-bit addition with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU512AddS(_:_:_:)
func VU512AddS(a uintptr, b uintptr, result uintptr) {
	if _vU512AddS == nil {
		panic("Accelerate: symbol vU512AddS not loaded")
	}
	_vU512AddS(a, b, result)
}


var _vU512Divide func(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr)

// VU512Divide computes the unsigned 512-bit division.
//
// See: https://developer.apple.com/documentation/Accelerate/vU512Divide(_:_:_:_:)
func VU512Divide(numerator uintptr, divisor uintptr, result uintptr, remainder uintptr) {
	if _vU512Divide == nil {
		panic("Accelerate: symbol vU512Divide not loaded")
	}
	_vU512Divide(numerator, divisor, result, remainder)
}


var _vU512FullMultiply func(a uintptr, b uintptr, result uintptr)

// VU512FullMultiply unsigned 512-bit multiplication; result is twice as wide as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU512FullMultiply(_:_:_:)
func VU512FullMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU512FullMultiply == nil {
		panic("Accelerate: symbol vU512FullMultiply not loaded")
	}
	_vU512FullMultiply(a, b, result)
}


var _vU512HalfMultiply func(a uintptr, b uintptr, result uintptr)

// VU512HalfMultiply unsigned 512-bit multiplication; result is the same width as multiplicands.
//
// See: https://developer.apple.com/documentation/Accelerate/vU512HalfMultiply(_:_:_:)
func VU512HalfMultiply(a uintptr, b uintptr, result uintptr) {
	if _vU512HalfMultiply == nil {
		panic("Accelerate: symbol vU512HalfMultiply not loaded")
	}
	_vU512HalfMultiply(a, b, result)
}


var _vU512Mod func(numerator uintptr, divisor uintptr, remainder uintptr)

// VU512Mod unsigned 512-bit mod.
//
// See: https://developer.apple.com/documentation/Accelerate/vU512Mod(_:_:_:)
func VU512Mod(numerator uintptr, divisor uintptr, remainder uintptr) {
	if _vU512Mod == nil {
		panic("Accelerate: symbol vU512Mod not loaded")
	}
	_vU512Mod(numerator, divisor, remainder)
}


var _vU512Neg func(a uintptr, result uintptr)

// VU512Neg unsigned 512-bit negation.
//
// See: https://developer.apple.com/documentation/Accelerate/vU512Neg(_:_:)
func VU512Neg(a uintptr, result uintptr) {
	if _vU512Neg == nil {
		panic("Accelerate: symbol vU512Neg not loaded")
	}
	_vU512Neg(a, result)
}


var _vU512Sub func(a uintptr, b uintptr, result uintptr)

// VU512Sub unsigned 512-bit subtraction (modular arithmetic).
//
// See: https://developer.apple.com/documentation/Accelerate/vU512Sub(_:_:_:)
func VU512Sub(a uintptr, b uintptr, result uintptr) {
	if _vU512Sub == nil {
		panic("Accelerate: symbol vU512Sub not loaded")
	}
	_vU512Sub(a, b, result)
}


var _vU512SubS func(a uintptr, b uintptr, result uintptr)

// VU512SubS unsigned 512-bit subtraction with saturation (clipping).
//
// See: https://developer.apple.com/documentation/Accelerate/vU512SubS(_:_:_:)
func VU512SubS(a uintptr, b uintptr, result uintptr) {
	if _vU512SubS == nil {
		panic("Accelerate: symbol vU512SubS not loaded")
	}
	_vU512SubS(a, b, result)
}




























































var _vvacos func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvacos calculates the arccosine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvacos(_:_:_:)
func Vvacos(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvacos == nil {
		panic("Accelerate: symbol vvacos not loaded")
	}
	_vvacos(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvacosf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvacosf calculates the arccosine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvacosf(_:_:_:)
func Vvacosf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvacosf == nil {
		panic("Accelerate: symbol vvacosf not loaded")
	}
	_vvacosf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvacosh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvacosh calculates the inverse hyperbolic cosine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvacosh(_:_:_:)
func Vvacosh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvacosh == nil {
		panic("Accelerate: symbol vvacosh not loaded")
	}
	_vvacosh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvacoshf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvacoshf calculates the inverse hyperbolic cosine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvacoshf(_:_:_:)
func Vvacoshf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvacoshf == nil {
		panic("Accelerate: symbol vvacoshf not loaded")
	}
	_vvacoshf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvasin func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvasin calculates the arcsine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvasin(_:_:_:)
func Vvasin(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvasin == nil {
		panic("Accelerate: symbol vvasin not loaded")
	}
	_vvasin(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvasinf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvasinf calculates the arcsine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvasinf(_:_:_:)
func Vvasinf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvasinf == nil {
		panic("Accelerate: symbol vvasinf not loaded")
	}
	_vvasinf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvasinh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvasinh calculates the inverse hyperbolic sine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvasinh(_:_:_:)
func Vvasinh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvasinh == nil {
		panic("Accelerate: symbol vvasinh not loaded")
	}
	_vvasinh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvasinhf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvasinhf calculates the inverse hyperbolic sine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvasinhf(_:_:_:)
func Vvasinhf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvasinhf == nil {
		panic("Accelerate: symbol vvasinhf not loaded")
	}
	_vvasinhf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvatan func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvatan calculates the arctangent of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatan(_:_:_:)
func Vvatan(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvatan == nil {
		panic("Accelerate: symbol vvatan not loaded")
	}
	_vvatan(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvatan2 func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvatan2 calculates the arctangent of each pair of elements in two arrays of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatan2(_:_:_:_:)
func Vvatan2(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvatan2 == nil {
		panic("Accelerate: symbol vvatan2 not loaded")
	}
	_vvatan2(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvatan2f func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvatan2f calculates the arctangent of each pair of elements in two arrays of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatan2f(_:_:_:_:)
func Vvatan2f(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvatan2f == nil {
		panic("Accelerate: symbol vvatan2f not loaded")
	}
	_vvatan2f(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvatanf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvatanf calculates the arctangent of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatanf(_:_:_:)
func Vvatanf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvatanf == nil {
		panic("Accelerate: symbol vvatanf not loaded")
	}
	_vvatanf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvatanh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvatanh calculates the inverse hyperbolic tangent of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatanh(_:_:_:)
func Vvatanh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvatanh == nil {
		panic("Accelerate: symbol vvatanh not loaded")
	}
	_vvatanh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvatanhf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvatanhf calculates the inverse hyperbolic tangent of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvatanhf(_:_:_:)
func Vvatanhf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvatanhf == nil {
		panic("Accelerate: symbol vvatanhf not loaded")
	}
	_vvatanhf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcbrt func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvcbrt calculates the cube root for each element of a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcbrt(_:_:_:)
func Vvcbrt(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvcbrt == nil {
		panic("Accelerate: symbol vvcbrt not loaded")
	}
	_vvcbrt(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcbrtf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvcbrtf calculates the cube root for each element of a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcbrtf(_:_:_:)
func Vvcbrtf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvcbrtf == nil {
		panic("Accelerate: symbol vvcbrtf not loaded")
	}
	_vvcbrtf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvceil func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvceil calculates the ceiling of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvceil(_:_:_:)
func Vvceil(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvceil == nil {
		panic("Accelerate: symbol vvceil not loaded")
	}
	_vvceil(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvceilf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvceilf calculates the ceiling of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvceilf(_:_:_:)
func Vvceilf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvceilf == nil {
		panic("Accelerate: symbol vvceilf not loaded")
	}
	_vvceilf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcopysign func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvcopysign copies an array, setting the sign of each element based on a second array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcopysign(_:_:_:_:)
func Vvcopysign(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvcopysign == nil {
		panic("Accelerate: symbol vvcopysign not loaded")
	}
	_vvcopysign(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvcopysignf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvcopysignf copies an array, setting the sign of each element based on a second array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcopysignf(_:_:_:_:)
func Vvcopysignf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvcopysignf == nil {
		panic("Accelerate: symbol vvcopysignf not loaded")
	}
	_vvcopysignf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvcos func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvcos calculates the cosine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcos(_:_:_:)
func Vvcos(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvcos == nil {
		panic("Accelerate: symbol vvcos not loaded")
	}
	_vvcos(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcosf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvcosf calculates the cosine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcosf(_:_:_:)
func Vvcosf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvcosf == nil {
		panic("Accelerate: symbol vvcosf not loaded")
	}
	_vvcosf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcosh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvcosh calculates the hyperbolic cosine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcosh(_:_:_:)
func Vvcosh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvcosh == nil {
		panic("Accelerate: symbol vvcosh not loaded")
	}
	_vvcosh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcoshf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvcoshf calculates the hyperbolic cosine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcoshf(_:_:_:)
func Vvcoshf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvcoshf == nil {
		panic("Accelerate: symbol vvcoshf not loaded")
	}
	_vvcoshf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcosisin func(arg0 uintptr, arg1 *float64, arg2 *int)

// Vvcosisin calculates the cosine and sine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcosisin(_:_:_:)
func Vvcosisin(arg0 uintptr, arg1 []float64, arg2 []int) {
	if _vvcosisin == nil {
		panic("Accelerate: symbol vvcosisin not loaded")
	}
	_vvcosisin(arg0, unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcosisinf func(arg0 uintptr, arg1 *float32, arg2 *int)

// Vvcosisinf calculates the cosine and sine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcosisinf(_:_:_:)
func Vvcosisinf(arg0 uintptr, arg1 []float32, arg2 []int) {
	if _vvcosisinf == nil {
		panic("Accelerate: symbol vvcosisinf not loaded")
	}
	_vvcosisinf(arg0, unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcospi func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvcospi calculates the cosine of pi multiplied by each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcospi(_:_:_:)
func Vvcospi(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvcospi == nil {
		panic("Accelerate: symbol vvcospi not loaded")
	}
	_vvcospi(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvcospif func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvcospif calculates the cosine of pi multiplied by each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvcospif(_:_:_:)
func Vvcospif(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvcospif == nil {
		panic("Accelerate: symbol vvcospif not loaded")
	}
	_vvcospif(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvdiv func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvdiv divides each element in an array by the corresponding value in a second array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvdiv(_:_:_:_:)
func Vvdiv(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvdiv == nil {
		panic("Accelerate: symbol vvdiv not loaded")
	}
	_vvdiv(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvdivf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvdivf divides each element in an array by the corresponding value in a second array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvdivf(_:_:_:_:)
func Vvdivf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvdivf == nil {
		panic("Accelerate: symbol vvdivf not loaded")
	}
	_vvdivf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvexp func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvexp calculates raised to the power of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexp(_:_:_:)
func Vvexp(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvexp == nil {
		panic("Accelerate: symbol vvexp not loaded")
	}
	_vvexp(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvexp2 func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvexp2 calculates 2 raised to the power of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexp2(_:_:_:)
func Vvexp2(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvexp2 == nil {
		panic("Accelerate: symbol vvexp2 not loaded")
	}
	_vvexp2(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvexp2f func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvexp2f calculates 2 raised to the power of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexp2f(_:_:_:)
func Vvexp2f(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvexp2f == nil {
		panic("Accelerate: symbol vvexp2f not loaded")
	}
	_vvexp2f(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvexpf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvexpf calculates raised to the power of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexpf(_:_:_:)
func Vvexpf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvexpf == nil {
		panic("Accelerate: symbol vvexpf not loaded")
	}
	_vvexpf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvexpm1 func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvexpm1 calculates for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexpm1(_:_:_:)
func Vvexpm1(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvexpm1 == nil {
		panic("Accelerate: symbol vvexpm1 not loaded")
	}
	_vvexpm1(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvexpm1f func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvexpm1f calculates for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvexpm1f(_:_:_:)
func Vvexpm1f(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvexpm1f == nil {
		panic("Accelerate: symbol vvexpm1f not loaded")
	}
	_vvexpm1f(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvfabs func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvfabs calculates the absolute value for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfabs(_:_:_:)
func Vvfabs(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvfabs == nil {
		panic("Accelerate: symbol vvfabs not loaded")
	}
	_vvfabs(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvfabsf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvfabsf calculates the absolute value for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfabsf(_:_:_:)
func Vvfabsf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvfabsf == nil {
		panic("Accelerate: symbol vvfabsf not loaded")
	}
	_vvfabsf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvfloor func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvfloor calculates the floor of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfloor(_:_:_:)
func Vvfloor(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvfloor == nil {
		panic("Accelerate: symbol vvfloor not loaded")
	}
	_vvfloor(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvfloorf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvfloorf calculates the floor of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfloorf(_:_:_:)
func Vvfloorf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvfloorf == nil {
		panic("Accelerate: symbol vvfloorf not loaded")
	}
	_vvfloorf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvfmod func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvfmod calculates the modulus after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfmod(_:_:_:_:)
func Vvfmod(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvfmod == nil {
		panic("Accelerate: symbol vvfmod not loaded")
	}
	_vvfmod(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvfmodf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvfmodf calculates the modulus after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvfmodf(_:_:_:_:)
func Vvfmodf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvfmodf == nil {
		panic("Accelerate: symbol vvfmodf not loaded")
	}
	_vvfmodf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvint func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvint calculates the integer truncation for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvint(_:_:_:)
func Vvint(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvint == nil {
		panic("Accelerate: symbol vvint not loaded")
	}
	_vvint(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvintf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvintf calculates the integer truncation for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvintf(_:_:_:)
func Vvintf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvintf == nil {
		panic("Accelerate: symbol vvintf not loaded")
	}
	_vvintf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvlog calculates the natural logarithm for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog(_:_:_:)
func Vvlog(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvlog == nil {
		panic("Accelerate: symbol vvlog not loaded")
	}
	_vvlog(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog10 func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvlog10 calculates the base 10 logarithm of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog10(_:_:_:)
func Vvlog10(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvlog10 == nil {
		panic("Accelerate: symbol vvlog10 not loaded")
	}
	_vvlog10(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog10f func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvlog10f calculates the base 10 logarithm of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog10f(_:_:_:)
func Vvlog10f(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvlog10f == nil {
		panic("Accelerate: symbol vvlog10f not loaded")
	}
	_vvlog10f(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog1p func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvlog1p calculates for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog1p(_:_:_:)
func Vvlog1p(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvlog1p == nil {
		panic("Accelerate: symbol vvlog1p not loaded")
	}
	_vvlog1p(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog1pf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvlog1pf calculates for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog1pf(_:_:_:)
func Vvlog1pf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvlog1pf == nil {
		panic("Accelerate: symbol vvlog1pf not loaded")
	}
	_vvlog1pf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog2 func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvlog2 calculates the base 2 logarithm of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog2(_:_:_:)
func Vvlog2(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvlog2 == nil {
		panic("Accelerate: symbol vvlog2 not loaded")
	}
	_vvlog2(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlog2f func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvlog2f calculates the base 2 logarithm of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlog2f(_:_:_:)
func Vvlog2f(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvlog2f == nil {
		panic("Accelerate: symbol vvlog2f not loaded")
	}
	_vvlog2f(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlogb func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvlogb calculates the unbiased exponent of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlogb(_:_:_:)
func Vvlogb(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvlogb == nil {
		panic("Accelerate: symbol vvlogb not loaded")
	}
	_vvlogb(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlogbf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvlogbf calculates the unbiased exponent of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlogbf(_:_:_:)
func Vvlogbf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvlogbf == nil {
		panic("Accelerate: symbol vvlogbf not loaded")
	}
	_vvlogbf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvlogf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvlogf calculates the natural logarithm for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvlogf(_:_:_:)
func Vvlogf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvlogf == nil {
		panic("Accelerate: symbol vvlogf not loaded")
	}
	_vvlogf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvnextafter func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvnextafter calculates the next machine-representable value for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvnextafter(_:_:_:_:)
func Vvnextafter(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvnextafter == nil {
		panic("Accelerate: symbol vvnextafter not loaded")
	}
	_vvnextafter(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvnextafterf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvnextafterf calculates the next machine-representable value for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvnextafterf(_:_:_:_:)
func Vvnextafterf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvnextafterf == nil {
		panic("Accelerate: symbol vvnextafterf not loaded")
	}
	_vvnextafterf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvnint func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvnint calculates the nearest integer for each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvnint(_:_:_:)
func Vvnint(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvnint == nil {
		panic("Accelerate: symbol vvnint not loaded")
	}
	_vvnint(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvnintf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvnintf calculates the nearest integer for each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvnintf(_:_:_:)
func Vvnintf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvnintf == nil {
		panic("Accelerate: symbol vvnintf not loaded")
	}
	_vvnintf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvpow func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvpow raises each element in an array to the power of the corresponding element in a second array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvpow(_:_:_:_:)
func Vvpow(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvpow == nil {
		panic("Accelerate: symbol vvpow not loaded")
	}
	_vvpow(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvpowf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvpowf raises each element in an array to the power of the corresponding element in a second array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvpowf(_:_:_:_:)
func Vvpowf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvpowf == nil {
		panic("Accelerate: symbol vvpowf not loaded")
	}
	_vvpowf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvpows func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvpows calculates the cube root for each element of a vector.
//
// See: https://developer.apple.com/documentation/Accelerate/vvpows(_:_:_:_:)
func Vvpows(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvpows == nil {
		panic("Accelerate: symbol vvpows not loaded")
	}
	_vvpows(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvpowsf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvpowsf calculates, elementwise, x**y for a vector x and a scalar y.
//
// See: https://developer.apple.com/documentation/Accelerate/vvpowsf(_:_:_:_:)
func Vvpowsf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvpowsf == nil {
		panic("Accelerate: symbol vvpowsf not loaded")
	}
	_vvpowsf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvrec func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvrec calculates the reciprocal of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvrec(_:_:_:)
func Vvrec(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvrec == nil {
		panic("Accelerate: symbol vvrec not loaded")
	}
	_vvrec(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvrecf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvrecf calculates the reciprocal of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvrecf(_:_:_:)
func Vvrecf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvrecf == nil {
		panic("Accelerate: symbol vvrecf not loaded")
	}
	_vvrecf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvremainder func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvremainder calculates the remainder after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvremainder(_:_:_:_:)
func Vvremainder(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvremainder == nil {
		panic("Accelerate: symbol vvremainder not loaded")
	}
	_vvremainder(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvremainderf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvremainderf calculates the remainder after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvremainderf(_:_:_:_:)
func Vvremainderf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvremainderf == nil {
		panic("Accelerate: symbol vvremainderf not loaded")
	}
	_vvremainderf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvrsqrt func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvrsqrt calculates the reciprocal square root of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvrsqrt(_:_:_:)
func Vvrsqrt(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvrsqrt == nil {
		panic("Accelerate: symbol vvrsqrt not loaded")
	}
	_vvrsqrt(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvrsqrtf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvrsqrtf calculates the reciprocal square root of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvrsqrtf(_:_:_:)
func Vvrsqrtf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvrsqrtf == nil {
		panic("Accelerate: symbol vvrsqrtf not loaded")
	}
	_vvrsqrtf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsin func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvsin calculates the sine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsin(_:_:_:)
func Vvsin(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvsin == nil {
		panic("Accelerate: symbol vvsin not loaded")
	}
	_vvsin(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsincos func(arg0 *float64, arg1 *float64, arg2 *float64, arg3 *int)

// Vvsincos calculates the cosine and sine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsincos(_:_:_:_:)
func Vvsincos(arg0 []float64, arg1 []float64, arg2 []float64, arg3 []int) {
	if _vvsincos == nil {
		panic("Accelerate: symbol vvsincos not loaded")
	}
	_vvsincos(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvsincosf func(arg0 *float32, arg1 *float32, arg2 *float32, arg3 *int)

// Vvsincosf calculates the cosine and sine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsincosf(_:_:_:_:)
func Vvsincosf(arg0 []float32, arg1 []float32, arg2 []float32, arg3 []int) {
	if _vvsincosf == nil {
		panic("Accelerate: symbol vvsincosf not loaded")
	}
	_vvsincosf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2), unsafe.SliceData(arg3))
}


var _vvsinf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvsinf calculates the sine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsinf(_:_:_:)
func Vvsinf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvsinf == nil {
		panic("Accelerate: symbol vvsinf not loaded")
	}
	_vvsinf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsinh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvsinh calculates the hyperbolic sine of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsinh(_:_:_:)
func Vvsinh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvsinh == nil {
		panic("Accelerate: symbol vvsinh not loaded")
	}
	_vvsinh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsinhf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvsinhf calculates the hyperbolic sine of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsinhf(_:_:_:)
func Vvsinhf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvsinhf == nil {
		panic("Accelerate: symbol vvsinhf not loaded")
	}
	_vvsinhf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsinpi func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvsinpi calculates the sine of pi multiplied by each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsinpi(_:_:_:)
func Vvsinpi(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvsinpi == nil {
		panic("Accelerate: symbol vvsinpi not loaded")
	}
	_vvsinpi(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsinpif func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvsinpif calculates the sine of pi multiplied by each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsinpif(_:_:_:)
func Vvsinpif(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvsinpif == nil {
		panic("Accelerate: symbol vvsinpif not loaded")
	}
	_vvsinpif(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsqrt func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvsqrt calculates the square root of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsqrt(_:_:_:)
func Vvsqrt(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvsqrt == nil {
		panic("Accelerate: symbol vvsqrt not loaded")
	}
	_vvsqrt(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvsqrtf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvsqrtf calculates the square root of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvsqrtf(_:_:_:)
func Vvsqrtf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvsqrtf == nil {
		panic("Accelerate: symbol vvsqrtf not loaded")
	}
	_vvsqrtf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtan func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvtan calculates the tangent of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtan(_:_:_:)
func Vvtan(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvtan == nil {
		panic("Accelerate: symbol vvtan not loaded")
	}
	_vvtan(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtanf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvtanf calculates the tangent of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtanf(_:_:_:)
func Vvtanf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvtanf == nil {
		panic("Accelerate: symbol vvtanf not loaded")
	}
	_vvtanf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtanh func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvtanh calculates the hyperbolic tangent of each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtanh(_:_:_:)
func Vvtanh(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvtanh == nil {
		panic("Accelerate: symbol vvtanh not loaded")
	}
	_vvtanh(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtanhf func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvtanhf calculates the hyperbolic tangent of each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtanhf(_:_:_:)
func Vvtanhf(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvtanhf == nil {
		panic("Accelerate: symbol vvtanhf not loaded")
	}
	_vvtanhf(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtanpi func(arg0 *float64, arg1 *float64, arg2 *int)

// Vvtanpi calculates the tangent of pi multiplied by each element in an array of double-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtanpi(_:_:_:)
func Vvtanpi(arg0 []float64, arg1 []float64, arg2 []int) {
	if _vvtanpi == nil {
		panic("Accelerate: symbol vvtanpi not loaded")
	}
	_vvtanpi(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}


var _vvtanpif func(arg0 *float32, arg1 *float32, arg2 *int)

// Vvtanpif calculates the tangent of pi multiplied by each element in an array of single-precision values.
//
// See: https://developer.apple.com/documentation/Accelerate/vvtanpif(_:_:_:)
func Vvtanpif(arg0 []float32, arg1 []float32, arg2 []int) {
	if _vvtanpif == nil {
		panic("Accelerate: symbol vvtanpif not loaded")
	}
	_vvtanpif(unsafe.SliceData(arg0), unsafe.SliceData(arg1), unsafe.SliceData(arg2))
}
































































































































































































































































































































































































































































































































































func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_bLASGetThreading, frameworkHandle, "BLASGetThreading")
		registerFunc(&_bLASSetThreading, frameworkHandle, "BLASSetThreading")
		registerFunc(&_bNNSApplyMultiheadAttention, frameworkHandle, "BNNSApplyMultiheadAttention")
		registerFunc(&_bNNSApplyMultiheadAttentionBackward, frameworkHandle, "BNNSApplyMultiheadAttentionBackward")
		registerFunc(&_bNNSArithmeticFilterApplyBackwardBatch, frameworkHandle, "BNNSArithmeticFilterApplyBackwardBatch")
		registerFunc(&_bNNSArithmeticFilterApplyBatch, frameworkHandle, "BNNSArithmeticFilterApplyBatch")
		registerFunc(&_bNNSBandPart, frameworkHandle, "BNNSBandPart")
		registerFunc(&_bNNSClipByGlobalNorm, frameworkHandle, "BNNSClipByGlobalNorm")
		registerFunc(&_bNNSClipByNorm, frameworkHandle, "BNNSClipByNorm")
		registerFunc(&_bNNSClipByValue, frameworkHandle, "BNNSClipByValue")
		registerFunc(&_bNNSCompareTensor, frameworkHandle, "BNNSCompareTensor")
		registerFunc(&_bNNSComputeLSTMTrainingCacheCapacity, frameworkHandle, "BNNSComputeLSTMTrainingCacheCapacity")
		registerFunc(&_bNNSComputeNorm, frameworkHandle, "BNNSComputeNorm")
		registerFunc(&_bNNSComputeNormBackward, frameworkHandle, "BNNSComputeNormBackward")
		registerFunc(&_bNNSCopy, frameworkHandle, "BNNSCopy")
		registerFunc(&_bNNSCreateNearestNeighbors, frameworkHandle, "BNNSCreateNearestNeighbors")
		registerFunc(&_bNNSCreateRandomGenerator, frameworkHandle, "BNNSCreateRandomGenerator")
		registerFunc(&_bNNSCreateRandomGeneratorWithSeed, frameworkHandle, "BNNSCreateRandomGeneratorWithSeed")
		registerFunc(&_bNNSCropResize, frameworkHandle, "BNNSCropResize")
		registerFunc(&_bNNSCropResizeBackward, frameworkHandle, "BNNSCropResizeBackward")
		registerFunc(&_bNNSDataLayoutGetRank, frameworkHandle, "BNNSDataLayoutGetRank")
		registerFunc(&_bNNSDestroyNearestNeighbors, frameworkHandle, "BNNSDestroyNearestNeighbors")
		registerFunc(&_bNNSDestroyRandomGenerator, frameworkHandle, "BNNSDestroyRandomGenerator")
		registerFunc(&_bNNSDirectApplyActivationBatch, frameworkHandle, "BNNSDirectApplyActivationBatch")
		registerFunc(&_bNNSDirectApplyInTopK, frameworkHandle, "BNNSDirectApplyInTopK")
		registerFunc(&_bNNSDirectApplyLSTMBatchBackward, frameworkHandle, "BNNSDirectApplyLSTMBatchBackward")
		registerFunc(&_bNNSDirectApplyLSTMBatchTrainingCaching, frameworkHandle, "BNNSDirectApplyLSTMBatchTrainingCaching")
		registerFunc(&_bNNSDirectApplyQuantizer, frameworkHandle, "BNNSDirectApplyQuantizer")
		registerFunc(&_bNNSDirectApplyReduction, frameworkHandle, "BNNSDirectApplyReduction")
		registerFunc(&_bNNSDirectApplyTopK, frameworkHandle, "BNNSDirectApplyTopK")
		registerFunc(&_bNNSFilterApply, frameworkHandle, "BNNSFilterApply")
		registerFunc(&_bNNSFilterApplyBackwardBatch, frameworkHandle, "BNNSFilterApplyBackwardBatch")
		registerFunc(&_bNNSFilterApplyBackwardTwoInputBatch, frameworkHandle, "BNNSFilterApplyBackwardTwoInputBatch")
		registerFunc(&_bNNSFilterApplyBatch, frameworkHandle, "BNNSFilterApplyBatch")
		registerFunc(&_bNNSFilterApplyTwoInput, frameworkHandle, "BNNSFilterApplyTwoInput")
		registerFunc(&_bNNSFilterApplyTwoInputBatch, frameworkHandle, "BNNSFilterApplyTwoInputBatch")
		registerFunc(&_bNNSFilterCreateFusedLayer, frameworkHandle, "BNNSFilterCreateFusedLayer")
		registerFunc(&_bNNSFilterCreateLayerActivation, frameworkHandle, "BNNSFilterCreateLayerActivation")
		registerFunc(&_bNNSFilterCreateLayerArithmetic, frameworkHandle, "BNNSFilterCreateLayerArithmetic")
		registerFunc(&_bNNSFilterCreateLayerBroadcastMatMul, frameworkHandle, "BNNSFilterCreateLayerBroadcastMatMul")
		registerFunc(&_bNNSFilterCreateLayerConvolution, frameworkHandle, "BNNSFilterCreateLayerConvolution")
		registerFunc(&_bNNSFilterCreateLayerDropout, frameworkHandle, "BNNSFilterCreateLayerDropout")
		registerFunc(&_bNNSFilterCreateLayerEmbedding, frameworkHandle, "BNNSFilterCreateLayerEmbedding")
		registerFunc(&_bNNSFilterCreateLayerFullyConnected, frameworkHandle, "BNNSFilterCreateLayerFullyConnected")
		registerFunc(&_bNNSFilterCreateLayerGram, frameworkHandle, "BNNSFilterCreateLayerGram")
		registerFunc(&_bNNSFilterCreateLayerLoss, frameworkHandle, "BNNSFilterCreateLayerLoss")
		registerFunc(&_bNNSFilterCreateLayerMultiheadAttention, frameworkHandle, "BNNSFilterCreateLayerMultiheadAttention")
		registerFunc(&_bNNSFilterCreateLayerNormalization, frameworkHandle, "BNNSFilterCreateLayerNormalization")
		registerFunc(&_bNNSFilterCreateLayerPadding, frameworkHandle, "BNNSFilterCreateLayerPadding")
		registerFunc(&_bNNSFilterCreateLayerPermute, frameworkHandle, "BNNSFilterCreateLayerPermute")
		registerFunc(&_bNNSFilterCreateLayerPooling, frameworkHandle, "BNNSFilterCreateLayerPooling")
		registerFunc(&_bNNSFilterCreateLayerReduction, frameworkHandle, "BNNSFilterCreateLayerReduction")
		registerFunc(&_bNNSFilterCreateLayerResize, frameworkHandle, "BNNSFilterCreateLayerResize")
		registerFunc(&_bNNSFilterCreateLayerTensorContraction, frameworkHandle, "BNNSFilterCreateLayerTensorContraction")
		registerFunc(&_bNNSFilterCreateLayerTransposedConvolution, frameworkHandle, "BNNSFilterCreateLayerTransposedConvolution")
		registerFunc(&_bNNSFilterDestroy, frameworkHandle, "BNNSFilterDestroy")
		registerFunc(&_bNNSFusedFilterApplyBackwardBatch, frameworkHandle, "BNNSFusedFilterApplyBackwardBatch")
		registerFunc(&_bNNSFusedFilterApplyBackwardMultiInputBatch, frameworkHandle, "BNNSFusedFilterApplyBackwardMultiInputBatch")
		registerFunc(&_bNNSFusedFilterApplyBatch, frameworkHandle, "BNNSFusedFilterApplyBatch")
		registerFunc(&_bNNSFusedFilterApplyMultiInputBatch, frameworkHandle, "BNNSFusedFilterApplyMultiInputBatch")
		registerFunc(&_bNNSGather, frameworkHandle, "BNNSGather")
		registerFunc(&_bNNSGatherND, frameworkHandle, "BNNSGatherND")
		registerFunc(&_bNNSGetPointer, frameworkHandle, "BNNSGetPointer")
		registerFunc(&_bNNSGraphCompileFromFile, frameworkHandle, "BNNSGraphCompileFromFile")
		registerFunc(&_bNNSGraphCompileOptionsDestroy, frameworkHandle, "BNNSGraphCompileOptionsDestroy")
		registerFunc(&_bNNSGraphCompileOptionsGetGenerateDebugInfo, frameworkHandle, "BNNSGraphCompileOptionsGetGenerateDebugInfo")
		registerFunc(&_bNNSGraphCompileOptionsGetOptimizationPreference, frameworkHandle, "BNNSGraphCompileOptionsGetOptimizationPreference")
		registerFunc(&_bNNSGraphCompileOptionsGetOutputFD, frameworkHandle, "BNNSGraphCompileOptionsGetOutputFD")
		registerFunc(&_bNNSGraphCompileOptionsGetOutputPath, frameworkHandle, "BNNSGraphCompileOptionsGetOutputPath")
		registerFunc(&_bNNSGraphCompileOptionsGetTargetSingleThread, frameworkHandle, "BNNSGraphCompileOptionsGetTargetSingleThread")
		registerFunc(&_bNNSGraphCompileOptionsMakeDefault, frameworkHandle, "BNNSGraphCompileOptionsMakeDefault")
		registerFunc(&_bNNSGraphCompileOptionsSetGenerateDebugInfo, frameworkHandle, "BNNSGraphCompileOptionsSetGenerateDebugInfo")
		registerFunc(&_bNNSGraphCompileOptionsSetMessageLogMask, frameworkHandle, "BNNSGraphCompileOptionsSetMessageLogMask")
		registerFunc(&_bNNSGraphCompileOptionsSetOptimizationPreference, frameworkHandle, "BNNSGraphCompileOptionsSetOptimizationPreference")
		registerFunc(&_bNNSGraphCompileOptionsSetOutputFD, frameworkHandle, "BNNSGraphCompileOptionsSetOutputFD")
		registerFunc(&_bNNSGraphCompileOptionsSetOutputPath, frameworkHandle, "BNNSGraphCompileOptionsSetOutputPath")
		registerFunc(&_bNNSGraphCompileOptionsSetTargetSingleThread, frameworkHandle, "BNNSGraphCompileOptionsSetTargetSingleThread")
		registerFunc(&_bNNSGraphContextDestroy, frameworkHandle, "BNNSGraphContextDestroy")
		registerFunc(&_bNNSGraphContextEnableNanAndInfChecks, frameworkHandle, "BNNSGraphContextEnableNanAndInfChecks")
		registerFunc(&_bNNSGraphContextExecute, frameworkHandle, "BNNSGraphContextExecute")
		registerFunc(&_bNNSGraphContextGetTensor, frameworkHandle, "BNNSGraphContextGetTensor")
		registerFunc(&_bNNSGraphContextGetWorkspaceSize, frameworkHandle, "BNNSGraphContextGetWorkspaceSize")
		registerFunc(&_bNNSGraphContextMake, frameworkHandle, "BNNSGraphContextMake")
		registerFunc(&_bNNSGraphContextMakeStreaming, frameworkHandle, "BNNSGraphContextMakeStreaming")
		registerFunc(&_bNNSGraphContextSetArgumentType, frameworkHandle, "BNNSGraphContextSetArgumentType")
		registerFunc(&_bNNSGraphContextSetBatchSize, frameworkHandle, "BNNSGraphContextSetBatchSize")
		registerFunc(&_bNNSGraphContextSetDynamicShapes, frameworkHandle, "BNNSGraphContextSetDynamicShapes")
		registerFunc(&_bNNSGraphContextSetMessageLogMask, frameworkHandle, "BNNSGraphContextSetMessageLogMask")
		registerFunc(&_bNNSGraphContextSetStreamingAdvanceCount, frameworkHandle, "BNNSGraphContextSetStreamingAdvanceCount")
		registerFunc(&_bNNSGraphGetArgumentCount, frameworkHandle, "BNNSGraphGetArgumentCount")
		registerFunc(&_bNNSGraphGetArgumentIntents, frameworkHandle, "BNNSGraphGetArgumentIntents")
		registerFunc(&_bNNSGraphGetArgumentInterleaveFactors, frameworkHandle, "BNNSGraphGetArgumentInterleaveFactors")
		registerFunc(&_bNNSGraphGetArgumentNames, frameworkHandle, "BNNSGraphGetArgumentNames")
		registerFunc(&_bNNSGraphGetArgumentPosition, frameworkHandle, "BNNSGraphGetArgumentPosition")
		registerFunc(&_bNNSGraphGetFunctionCount, frameworkHandle, "BNNSGraphGetFunctionCount")
		registerFunc(&_bNNSGraphGetFunctionNames, frameworkHandle, "BNNSGraphGetFunctionNames")
		registerFunc(&_bNNSGraphGetInputCount, frameworkHandle, "BNNSGraphGetInputCount")
		registerFunc(&_bNNSGraphGetInputNames, frameworkHandle, "BNNSGraphGetInputNames")
		registerFunc(&_bNNSGraphGetOutputCount, frameworkHandle, "BNNSGraphGetOutputCount")
		registerFunc(&_bNNSGraphGetOutputNames, frameworkHandle, "BNNSGraphGetOutputNames")
		registerFunc(&_bNNSGraphTensorFillStrides, frameworkHandle, "BNNSGraphTensorFillStrides")
		registerFunc(&_bNNSLossFilterApplyBackwardBatch, frameworkHandle, "BNNSLossFilterApplyBackwardBatch")
		registerFunc(&_bNNSLossFilterApplyBatch, frameworkHandle, "BNNSLossFilterApplyBatch")
		registerFunc(&_bNNSMatMul, frameworkHandle, "BNNSMatMul")
		registerFunc(&_bNNSMatMulWorkspaceSize, frameworkHandle, "BNNSMatMulWorkspaceSize")
		registerFunc(&_bNNSNDArrayFullyConnectedSparsifySparseCOO, frameworkHandle, "BNNSNDArrayFullyConnectedSparsifySparseCOO")
		registerFunc(&_bNNSNDArrayFullyConnectedSparsifySparseCSR, frameworkHandle, "BNNSNDArrayFullyConnectedSparsifySparseCSR")
		registerFunc(&_bNNSNDArrayGetDataSize, frameworkHandle, "BNNSNDArrayGetDataSize")
		registerFunc(&_bNNSNearestNeighborsGetInfo, frameworkHandle, "BNNSNearestNeighborsGetInfo")
		registerFunc(&_bNNSNearestNeighborsLoad, frameworkHandle, "BNNSNearestNeighborsLoad")
		registerFunc(&_bNNSNormalizationFilterApplyBackwardBatch, frameworkHandle, "BNNSNormalizationFilterApplyBackwardBatch")
		registerFunc(&_bNNSNormalizationFilterApplyBatch, frameworkHandle, "BNNSNormalizationFilterApplyBatch")
		registerFunc(&_bNNSOptimizerStep, frameworkHandle, "BNNSOptimizerStep")
		registerFunc(&_bNNSPermuteFilterApplyBackwardBatch, frameworkHandle, "BNNSPermuteFilterApplyBackwardBatch")
		registerFunc(&_bNNSPoolingFilterApplyBackwardBatch, frameworkHandle, "BNNSPoolingFilterApplyBackwardBatch")
		registerFunc(&_bNNSPoolingFilterApplyBackwardBatchEx, frameworkHandle, "BNNSPoolingFilterApplyBackwardBatchEx")
		registerFunc(&_bNNSPoolingFilterApplyBatch, frameworkHandle, "BNNSPoolingFilterApplyBatch")
		registerFunc(&_bNNSPoolingFilterApplyBatchEx, frameworkHandle, "BNNSPoolingFilterApplyBatchEx")
		registerFunc(&_bNNSRandomFillCategoricalFloat, frameworkHandle, "BNNSRandomFillCategoricalFloat")
		registerFunc(&_bNNSRandomFillNormalFloat, frameworkHandle, "BNNSRandomFillNormalFloat")
		registerFunc(&_bNNSRandomFillUniformFloat, frameworkHandle, "BNNSRandomFillUniformFloat")
		registerFunc(&_bNNSRandomFillUniformInt, frameworkHandle, "BNNSRandomFillUniformInt")
		registerFunc(&_bNNSRandomGeneratorGetState, frameworkHandle, "BNNSRandomGeneratorGetState")
		registerFunc(&_bNNSRandomGeneratorSetState, frameworkHandle, "BNNSRandomGeneratorSetState")
		registerFunc(&_bNNSRandomGeneratorStateSize, frameworkHandle, "BNNSRandomGeneratorStateSize")
		registerFunc(&_bNNSScatter, frameworkHandle, "BNNSScatter")
		registerFunc(&_bNNSScatterND, frameworkHandle, "BNNSScatterND")
		registerFunc(&_bNNSShuffle, frameworkHandle, "BNNSShuffle")
		registerFunc(&_bNNSTensorGetAllocationSize, frameworkHandle, "BNNSTensorGetAllocationSize")
		registerFunc(&_bNNSTile, frameworkHandle, "BNNSTile")
		registerFunc(&_bNNSTileBackward, frameworkHandle, "BNNSTileBackward")
		registerFunc(&_bNNSTranspose, frameworkHandle, "BNNSTranspose")
		registerFunc(&_setBLASParamErrorProc, frameworkHandle, "SetBLASParamErrorProc")
		registerFunc(&_sparseGetInertia, frameworkHandle, "SparseGetInertia")
		registerFunc(&_appleblas_dgeadd, frameworkHandle, "appleblas_dgeadd")
		registerFunc(&_appleblas_sgeadd, frameworkHandle, "appleblas_sgeadd")
		registerFunc(&_catlas_caxpby, frameworkHandle, "catlas_caxpby")
		registerFunc(&_catlas_cset, frameworkHandle, "catlas_cset")
		registerFunc(&_catlas_daxpby, frameworkHandle, "catlas_daxpby")
		registerFunc(&_catlas_dset, frameworkHandle, "catlas_dset")
		registerFunc(&_catlas_saxpby, frameworkHandle, "catlas_saxpby")
		registerFunc(&_catlas_sset, frameworkHandle, "catlas_sset")
		registerFunc(&_catlas_zaxpby, frameworkHandle, "catlas_zaxpby")
		registerFunc(&_catlas_zset, frameworkHandle, "catlas_zset")
		registerFunc(&_cblas_caxpy, frameworkHandle, "cblas_caxpy")
		registerFunc(&_cblas_ccopy, frameworkHandle, "cblas_ccopy")
		registerFunc(&_cblas_cdotc_sub, frameworkHandle, "cblas_cdotc_sub")
		registerFunc(&_cblas_cdotu_sub, frameworkHandle, "cblas_cdotu_sub")
		registerFunc(&_cblas_cgbmv, frameworkHandle, "cblas_cgbmv")
		registerFunc(&_cblas_cgemm, frameworkHandle, "cblas_cgemm")
		registerFunc(&_cblas_cgemv, frameworkHandle, "cblas_cgemv")
		registerFunc(&_cblas_cgerc, frameworkHandle, "cblas_cgerc")
		registerFunc(&_cblas_cgeru, frameworkHandle, "cblas_cgeru")
		registerFunc(&_cblas_chbmv, frameworkHandle, "cblas_chbmv")
		registerFunc(&_cblas_chemm, frameworkHandle, "cblas_chemm")
		registerFunc(&_cblas_chemv, frameworkHandle, "cblas_chemv")
		registerFunc(&_cblas_cher, frameworkHandle, "cblas_cher")
		registerFunc(&_cblas_cher2, frameworkHandle, "cblas_cher2")
		registerFunc(&_cblas_cher2k, frameworkHandle, "cblas_cher2k")
		registerFunc(&_cblas_cherk, frameworkHandle, "cblas_cherk")
		registerFunc(&_cblas_chpmv, frameworkHandle, "cblas_chpmv")
		registerFunc(&_cblas_chpr, frameworkHandle, "cblas_chpr")
		registerFunc(&_cblas_chpr2, frameworkHandle, "cblas_chpr2")
		registerFunc(&_cblas_crotg, frameworkHandle, "cblas_crotg")
		registerFunc(&_cblas_cscal, frameworkHandle, "cblas_cscal")
		registerFunc(&_cblas_csrot, frameworkHandle, "cblas_csrot")
		registerFunc(&_cblas_csscal, frameworkHandle, "cblas_csscal")
		registerFunc(&_cblas_cswap, frameworkHandle, "cblas_cswap")
		registerFunc(&_cblas_csymm, frameworkHandle, "cblas_csymm")
		registerFunc(&_cblas_csyr2k, frameworkHandle, "cblas_csyr2k")
		registerFunc(&_cblas_csyrk, frameworkHandle, "cblas_csyrk")
		registerFunc(&_cblas_ctbmv, frameworkHandle, "cblas_ctbmv")
		registerFunc(&_cblas_ctbsv, frameworkHandle, "cblas_ctbsv")
		registerFunc(&_cblas_ctpmv, frameworkHandle, "cblas_ctpmv")
		registerFunc(&_cblas_ctpsv, frameworkHandle, "cblas_ctpsv")
		registerFunc(&_cblas_ctrmm, frameworkHandle, "cblas_ctrmm")
		registerFunc(&_cblas_ctrmv, frameworkHandle, "cblas_ctrmv")
		registerFunc(&_cblas_ctrsm, frameworkHandle, "cblas_ctrsm")
		registerFunc(&_cblas_ctrsv, frameworkHandle, "cblas_ctrsv")
		registerFunc(&_cblas_dasum, frameworkHandle, "cblas_dasum")
		registerFunc(&_cblas_daxpy, frameworkHandle, "cblas_daxpy")
		registerFunc(&_cblas_dcopy, frameworkHandle, "cblas_dcopy")
		registerFunc(&_cblas_ddot, frameworkHandle, "cblas_ddot")
		registerFunc(&_cblas_dgbmv, frameworkHandle, "cblas_dgbmv")
		registerFunc(&_cblas_dgemm, frameworkHandle, "cblas_dgemm")
		registerFunc(&_cblas_dgemv, frameworkHandle, "cblas_dgemv")
		registerFunc(&_cblas_dger, frameworkHandle, "cblas_dger")
		registerFunc(&_cblas_dnrm2, frameworkHandle, "cblas_dnrm2")
		registerFunc(&_cblas_drot, frameworkHandle, "cblas_drot")
		registerFunc(&_cblas_drotg, frameworkHandle, "cblas_drotg")
		registerFunc(&_cblas_drotm, frameworkHandle, "cblas_drotm")
		registerFunc(&_cblas_drotmg, frameworkHandle, "cblas_drotmg")
		registerFunc(&_cblas_dsbmv, frameworkHandle, "cblas_dsbmv")
		registerFunc(&_cblas_dscal, frameworkHandle, "cblas_dscal")
		registerFunc(&_cblas_dsdot, frameworkHandle, "cblas_dsdot")
		registerFunc(&_cblas_dspmv, frameworkHandle, "cblas_dspmv")
		registerFunc(&_cblas_dspr, frameworkHandle, "cblas_dspr")
		registerFunc(&_cblas_dspr2, frameworkHandle, "cblas_dspr2")
		registerFunc(&_cblas_dswap, frameworkHandle, "cblas_dswap")
		registerFunc(&_cblas_dsymm, frameworkHandle, "cblas_dsymm")
		registerFunc(&_cblas_dsymv, frameworkHandle, "cblas_dsymv")
		registerFunc(&_cblas_dsyr, frameworkHandle, "cblas_dsyr")
		registerFunc(&_cblas_dsyr2, frameworkHandle, "cblas_dsyr2")
		registerFunc(&_cblas_dsyr2k, frameworkHandle, "cblas_dsyr2k")
		registerFunc(&_cblas_dsyrk, frameworkHandle, "cblas_dsyrk")
		registerFunc(&_cblas_dtbmv, frameworkHandle, "cblas_dtbmv")
		registerFunc(&_cblas_dtbsv, frameworkHandle, "cblas_dtbsv")
		registerFunc(&_cblas_dtpmv, frameworkHandle, "cblas_dtpmv")
		registerFunc(&_cblas_dtpsv, frameworkHandle, "cblas_dtpsv")
		registerFunc(&_cblas_dtrmm, frameworkHandle, "cblas_dtrmm")
		registerFunc(&_cblas_dtrmv, frameworkHandle, "cblas_dtrmv")
		registerFunc(&_cblas_dtrsm, frameworkHandle, "cblas_dtrsm")
		registerFunc(&_cblas_dtrsv, frameworkHandle, "cblas_dtrsv")
		registerFunc(&_cblas_dzasum, frameworkHandle, "cblas_dzasum")
		registerFunc(&_cblas_dznrm2, frameworkHandle, "cblas_dznrm2")
		registerFunc(&_cblas_errprn, frameworkHandle, "cblas_errprn")
		registerFunc(&_cblas_icamax, frameworkHandle, "cblas_icamax")
		registerFunc(&_cblas_idamax, frameworkHandle, "cblas_idamax")
		registerFunc(&_cblas_isamax, frameworkHandle, "cblas_isamax")
		registerFunc(&_cblas_izamax, frameworkHandle, "cblas_izamax")
		registerFunc(&_cblas_sasum, frameworkHandle, "cblas_sasum")
		registerFunc(&_cblas_saxpy, frameworkHandle, "cblas_saxpy")
		registerFunc(&_cblas_scasum, frameworkHandle, "cblas_scasum")
		registerFunc(&_cblas_scnrm2, frameworkHandle, "cblas_scnrm2")
		registerFunc(&_cblas_scopy, frameworkHandle, "cblas_scopy")
		registerFunc(&_cblas_sdot, frameworkHandle, "cblas_sdot")
		registerFunc(&_cblas_sdsdot, frameworkHandle, "cblas_sdsdot")
		registerFunc(&_cblas_sgbmv, frameworkHandle, "cblas_sgbmv")
		registerFunc(&_cblas_sgemm, frameworkHandle, "cblas_sgemm")
		registerFunc(&_cblas_sgemv, frameworkHandle, "cblas_sgemv")
		registerFunc(&_cblas_sger, frameworkHandle, "cblas_sger")
		registerFunc(&_cblas_snrm2, frameworkHandle, "cblas_snrm2")
		registerFunc(&_cblas_srot, frameworkHandle, "cblas_srot")
		registerFunc(&_cblas_srotg, frameworkHandle, "cblas_srotg")
		registerFunc(&_cblas_srotm, frameworkHandle, "cblas_srotm")
		registerFunc(&_cblas_srotmg, frameworkHandle, "cblas_srotmg")
		registerFunc(&_cblas_ssbmv, frameworkHandle, "cblas_ssbmv")
		registerFunc(&_cblas_sscal, frameworkHandle, "cblas_sscal")
		registerFunc(&_cblas_sspmv, frameworkHandle, "cblas_sspmv")
		registerFunc(&_cblas_sspr, frameworkHandle, "cblas_sspr")
		registerFunc(&_cblas_sspr2, frameworkHandle, "cblas_sspr2")
		registerFunc(&_cblas_sswap, frameworkHandle, "cblas_sswap")
		registerFunc(&_cblas_ssymm, frameworkHandle, "cblas_ssymm")
		registerFunc(&_cblas_ssymv, frameworkHandle, "cblas_ssymv")
		registerFunc(&_cblas_ssyr, frameworkHandle, "cblas_ssyr")
		registerFunc(&_cblas_ssyr2, frameworkHandle, "cblas_ssyr2")
		registerFunc(&_cblas_ssyr2k, frameworkHandle, "cblas_ssyr2k")
		registerFunc(&_cblas_ssyrk, frameworkHandle, "cblas_ssyrk")
		registerFunc(&_cblas_stbmv, frameworkHandle, "cblas_stbmv")
		registerFunc(&_cblas_stbsv, frameworkHandle, "cblas_stbsv")
		registerFunc(&_cblas_stpmv, frameworkHandle, "cblas_stpmv")
		registerFunc(&_cblas_stpsv, frameworkHandle, "cblas_stpsv")
		registerFunc(&_cblas_strmm, frameworkHandle, "cblas_strmm")
		registerFunc(&_cblas_strmv, frameworkHandle, "cblas_strmv")
		registerFunc(&_cblas_strsm, frameworkHandle, "cblas_strsm")
		registerFunc(&_cblas_strsv, frameworkHandle, "cblas_strsv")
		registerFunc(&_cblas_xerbla, frameworkHandle, "cblas_xerbla")
		registerFunc(&_cblas_zaxpy, frameworkHandle, "cblas_zaxpy")
		registerFunc(&_cblas_zcopy, frameworkHandle, "cblas_zcopy")
		registerFunc(&_cblas_zdotc_sub, frameworkHandle, "cblas_zdotc_sub")
		registerFunc(&_cblas_zdotu_sub, frameworkHandle, "cblas_zdotu_sub")
		registerFunc(&_cblas_zdrot, frameworkHandle, "cblas_zdrot")
		registerFunc(&_cblas_zdscal, frameworkHandle, "cblas_zdscal")
		registerFunc(&_cblas_zgbmv, frameworkHandle, "cblas_zgbmv")
		registerFunc(&_cblas_zgemm, frameworkHandle, "cblas_zgemm")
		registerFunc(&_cblas_zgemv, frameworkHandle, "cblas_zgemv")
		registerFunc(&_cblas_zgerc, frameworkHandle, "cblas_zgerc")
		registerFunc(&_cblas_zgeru, frameworkHandle, "cblas_zgeru")
		registerFunc(&_cblas_zhbmv, frameworkHandle, "cblas_zhbmv")
		registerFunc(&_cblas_zhemm, frameworkHandle, "cblas_zhemm")
		registerFunc(&_cblas_zhemv, frameworkHandle, "cblas_zhemv")
		registerFunc(&_cblas_zher, frameworkHandle, "cblas_zher")
		registerFunc(&_cblas_zher2, frameworkHandle, "cblas_zher2")
		registerFunc(&_cblas_zher2k, frameworkHandle, "cblas_zher2k")
		registerFunc(&_cblas_zherk, frameworkHandle, "cblas_zherk")
		registerFunc(&_cblas_zhpmv, frameworkHandle, "cblas_zhpmv")
		registerFunc(&_cblas_zhpr, frameworkHandle, "cblas_zhpr")
		registerFunc(&_cblas_zhpr2, frameworkHandle, "cblas_zhpr2")
		registerFunc(&_cblas_zrotg, frameworkHandle, "cblas_zrotg")
		registerFunc(&_cblas_zscal, frameworkHandle, "cblas_zscal")
		registerFunc(&_cblas_zswap, frameworkHandle, "cblas_zswap")
		registerFunc(&_cblas_zsymm, frameworkHandle, "cblas_zsymm")
		registerFunc(&_cblas_zsyr2k, frameworkHandle, "cblas_zsyr2k")
		registerFunc(&_cblas_zsyrk, frameworkHandle, "cblas_zsyrk")
		registerFunc(&_cblas_ztbmv, frameworkHandle, "cblas_ztbmv")
		registerFunc(&_cblas_ztbsv, frameworkHandle, "cblas_ztbsv")
		registerFunc(&_cblas_ztpmv, frameworkHandle, "cblas_ztpmv")
		registerFunc(&_cblas_ztpsv, frameworkHandle, "cblas_ztpsv")
		registerFunc(&_cblas_ztrmm, frameworkHandle, "cblas_ztrmm")
		registerFunc(&_cblas_ztrmv, frameworkHandle, "cblas_ztrmv")
		registerFunc(&_cblas_ztrsm, frameworkHandle, "cblas_ztrsm")
		registerFunc(&_cblas_ztrsv, frameworkHandle, "cblas_ztrsv")
		registerFunc(&_quadrature_integrate, frameworkHandle, "quadrature_integrate")
		registerFunc(&_sparse_commit, frameworkHandle, "sparse_commit")
		registerFunc(&_sparse_get_matrix_nonzero_count, frameworkHandle, "sparse_get_matrix_nonzero_count")
		registerFunc(&_sparse_get_matrix_property, frameworkHandle, "sparse_get_matrix_property")
		registerFunc(&_sparse_matrix_destroy, frameworkHandle, "sparse_matrix_destroy")
		registerFunc(&_sparse_set_matrix_property, frameworkHandle, "sparse_set_matrix_property")
		registerFunc(&_vA1024Shift, frameworkHandle, "vA1024Shift")
		registerFunc(&_vA256Shift, frameworkHandle, "vA256Shift")
		registerFunc(&_vA512Shift, frameworkHandle, "vA512Shift")
		registerFunc(&_vDSP_DCT_Execute, frameworkHandle, "vDSP_DCT_Execute")
		registerFunc(&_vDSP_DFT_Execute, frameworkHandle, "vDSP_DFT_Execute")
		registerFunc(&_vDSP_DFT_ExecuteD, frameworkHandle, "vDSP_DFT_ExecuteD")
		registerFunc(&_vDSP_FFT16_copv, frameworkHandle, "vDSP_FFT16_copv")
		registerFunc(&_vDSP_FFT16_zopv, frameworkHandle, "vDSP_FFT16_zopv")
		registerFunc(&_vDSP_FFT32_copv, frameworkHandle, "vDSP_FFT32_copv")
		registerFunc(&_vDSP_FFT32_zopv, frameworkHandle, "vDSP_FFT32_zopv")
		registerFunc(&_vDSP_destroy_fftsetup, frameworkHandle, "vDSP_destroy_fftsetup")
		registerFunc(&_vDSP_destroy_fftsetupD, frameworkHandle, "vDSP_destroy_fftsetupD")
		registerFunc(&_vImageAffineWarpD_ARGB16F, frameworkHandle, "vImageAffineWarpD_ARGB16F")
		registerFunc(&_vImageAffineWarpD_ARGB16S, frameworkHandle, "vImageAffineWarpD_ARGB16S")
		registerFunc(&_vImageAffineWarpD_ARGB16U, frameworkHandle, "vImageAffineWarpD_ARGB16U")
		registerFunc(&_vImageAffineWarpD_ARGB8888, frameworkHandle, "vImageAffineWarpD_ARGB8888")
		registerFunc(&_vImageAffineWarpD_ARGBFFFF, frameworkHandle, "vImageAffineWarpD_ARGBFFFF")
		registerFunc(&_vImageAffineWarpD_CbCr16F, frameworkHandle, "vImageAffineWarpD_CbCr16F")
		registerFunc(&_vImageAffineWarpD_Planar16F, frameworkHandle, "vImageAffineWarpD_Planar16F")
		registerFunc(&_vImageAffineWarpD_Planar8, frameworkHandle, "vImageAffineWarpD_Planar8")
		registerFunc(&_vImageAffineWarpD_PlanarF, frameworkHandle, "vImageAffineWarpD_PlanarF")
		registerFunc(&_vImageAffineWarp_ARGB16F, frameworkHandle, "vImageAffineWarp_ARGB16F")
		registerFunc(&_vImageAffineWarp_ARGB16S, frameworkHandle, "vImageAffineWarp_ARGB16S")
		registerFunc(&_vImageAffineWarp_ARGB16U, frameworkHandle, "vImageAffineWarp_ARGB16U")
		registerFunc(&_vImageAffineWarp_ARGB8888, frameworkHandle, "vImageAffineWarp_ARGB8888")
		registerFunc(&_vImageAffineWarp_ARGBFFFF, frameworkHandle, "vImageAffineWarp_ARGBFFFF")
		registerFunc(&_vImageAffineWarp_CbCr16F, frameworkHandle, "vImageAffineWarp_CbCr16F")
		registerFunc(&_vImageAffineWarp_Planar16F, frameworkHandle, "vImageAffineWarp_Planar16F")
		registerFunc(&_vImageAffineWarp_Planar8, frameworkHandle, "vImageAffineWarp_Planar8")
		registerFunc(&_vImageAffineWarp_PlanarF, frameworkHandle, "vImageAffineWarp_PlanarF")
		registerFunc(&_vImageAlphaBlend_ARGB8888, frameworkHandle, "vImageAlphaBlend_ARGB8888")
		registerFunc(&_vImageAlphaBlend_ARGBFFFF, frameworkHandle, "vImageAlphaBlend_ARGBFFFF")
		registerFunc(&_vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888, frameworkHandle, "vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGB8888")
		registerFunc(&_vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF, frameworkHandle, "vImageAlphaBlend_NonpremultipliedToPremultiplied_ARGBFFFF")
		registerFunc(&_vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8, frameworkHandle, "vImageAlphaBlend_NonpremultipliedToPremultiplied_Planar8")
		registerFunc(&_vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF, frameworkHandle, "vImageAlphaBlend_NonpremultipliedToPremultiplied_PlanarF")
		registerFunc(&_vImageAlphaBlend_Planar8, frameworkHandle, "vImageAlphaBlend_Planar8")
		registerFunc(&_vImageAlphaBlend_PlanarF, frameworkHandle, "vImageAlphaBlend_PlanarF")
		registerFunc(&_vImageBoxConvolve_ARGB8888, frameworkHandle, "vImageBoxConvolve_ARGB8888")
		registerFunc(&_vImageBoxConvolve_Planar8, frameworkHandle, "vImageBoxConvolve_Planar8")
		registerFunc(&_vImageBufferFill_ARGB16F, frameworkHandle, "vImageBufferFill_ARGB16F")
		registerFunc(&_vImageBufferFill_ARGB16S, frameworkHandle, "vImageBufferFill_ARGB16S")
		registerFunc(&_vImageBufferFill_ARGB16U, frameworkHandle, "vImageBufferFill_ARGB16U")
		registerFunc(&_vImageBufferFill_ARGB8888, frameworkHandle, "vImageBufferFill_ARGB8888")
		registerFunc(&_vImageBufferFill_ARGBFFFF, frameworkHandle, "vImageBufferFill_ARGBFFFF")
		registerFunc(&_vImageBufferFill_CbCr16S, frameworkHandle, "vImageBufferFill_CbCr16S")
		registerFunc(&_vImageBufferFill_CbCr16U, frameworkHandle, "vImageBufferFill_CbCr16U")
		registerFunc(&_vImageBufferFill_CbCr8, frameworkHandle, "vImageBufferFill_CbCr8")
		registerFunc(&_vImageBuffer_CopyToCVPixelBuffer, frameworkHandle, "vImageBuffer_CopyToCVPixelBuffer")
		registerFunc(&_vImageBuffer_GetSize, frameworkHandle, "vImageBuffer_GetSize")
		registerFunc(&_vImageBuffer_Init, frameworkHandle, "vImageBuffer_Init")
		registerFunc(&_vImageBuffer_InitForCopyFromCVPixelBuffer, frameworkHandle, "vImageBuffer_InitForCopyFromCVPixelBuffer")
		registerFunc(&_vImageBuffer_InitForCopyToCVPixelBuffer, frameworkHandle, "vImageBuffer_InitForCopyToCVPixelBuffer")
		registerFunc(&_vImageBuffer_InitWithCGImage, frameworkHandle, "vImageBuffer_InitWithCGImage")
		registerFunc(&_vImageBuffer_InitWithCVPixelBuffer, frameworkHandle, "vImageBuffer_InitWithCVPixelBuffer")
		registerFunc(&_vImageByteSwap_Planar16U, frameworkHandle, "vImageByteSwap_Planar16U")
		registerFunc(&_vImageCGImageFormat_GetComponentCount, frameworkHandle, "vImageCGImageFormat_GetComponentCount")
		registerFunc(&_vImageCGImageFormat_IsEqual, frameworkHandle, "vImageCGImageFormat_IsEqual")
		registerFunc(&_vImageCVImageFormat_Copy, frameworkHandle, "vImageCVImageFormat_Copy")
		registerFunc(&_vImageCVImageFormat_Create, frameworkHandle, "vImageCVImageFormat_Create")
		registerFunc(&_vImageCVImageFormat_CreateWithCVPixelBuffer, frameworkHandle, "vImageCVImageFormat_CreateWithCVPixelBuffer")
		registerFunc(&_vImageCVImageFormat_GetAlphaHint, frameworkHandle, "vImageCVImageFormat_GetAlphaHint")
		registerFunc(&_vImageCVImageFormat_GetChannelCount, frameworkHandle, "vImageCVImageFormat_GetChannelCount")
		registerFunc(&_vImageCVImageFormat_GetChromaSiting, frameworkHandle, "vImageCVImageFormat_GetChromaSiting")
		registerFunc(&_vImageCVImageFormat_GetColorSpace, frameworkHandle, "vImageCVImageFormat_GetColorSpace")
		registerFunc(&_vImageCVImageFormat_GetFormatCode, frameworkHandle, "vImageCVImageFormat_GetFormatCode")
		registerFunc(&_vImageCVImageFormat_GetUserData, frameworkHandle, "vImageCVImageFormat_GetUserData")
		registerFunc(&_vImageCVImageFormat_Release, frameworkHandle, "vImageCVImageFormat_Release")
		registerFunc(&_vImageCVImageFormat_Retain, frameworkHandle, "vImageCVImageFormat_Retain")
		registerFunc(&_vImageCVImageFormat_SetAlphaHint, frameworkHandle, "vImageCVImageFormat_SetAlphaHint")
		registerFunc(&_vImageCVImageFormat_SetChromaSiting, frameworkHandle, "vImageCVImageFormat_SetChromaSiting")
		registerFunc(&_vImageCVImageFormat_SetColorSpace, frameworkHandle, "vImageCVImageFormat_SetColorSpace")
		registerFunc(&_vImageCVImageFormat_SetUserData, frameworkHandle, "vImageCVImageFormat_SetUserData")
		registerFunc(&_vImageClipToAlpha_ARGB8888, frameworkHandle, "vImageClipToAlpha_ARGB8888")
		registerFunc(&_vImageClipToAlpha_ARGBFFFF, frameworkHandle, "vImageClipToAlpha_ARGBFFFF")
		registerFunc(&_vImageClipToAlpha_Planar8, frameworkHandle, "vImageClipToAlpha_Planar8")
		registerFunc(&_vImageClipToAlpha_PlanarF, frameworkHandle, "vImageClipToAlpha_PlanarF")
		registerFunc(&_vImageClipToAlpha_RGBA8888, frameworkHandle, "vImageClipToAlpha_RGBA8888")
		registerFunc(&_vImageClipToAlpha_RGBAFFFF, frameworkHandle, "vImageClipToAlpha_RGBAFFFF")
		registerFunc(&_vImageClip_PlanarF, frameworkHandle, "vImageClip_PlanarF")
		registerFunc(&_vImageContrastStretch_ARGB8888, frameworkHandle, "vImageContrastStretch_ARGB8888")
		registerFunc(&_vImageContrastStretch_ARGBFFFF, frameworkHandle, "vImageContrastStretch_ARGBFFFF")
		registerFunc(&_vImageContrastStretch_Planar8, frameworkHandle, "vImageContrastStretch_Planar8")
		registerFunc(&_vImageContrastStretch_PlanarF, frameworkHandle, "vImageContrastStretch_PlanarF")
		registerFunc(&_vImageConvert_12UTo16U, frameworkHandle, "vImageConvert_12UTo16U")
		registerFunc(&_vImageConvert_16Fto16Q12, frameworkHandle, "vImageConvert_16Fto16Q12")
		registerFunc(&_vImageConvert_16Fto16U, frameworkHandle, "vImageConvert_16Fto16U")
		registerFunc(&_vImageConvert_16Q12to16F, frameworkHandle, "vImageConvert_16Q12to16F")
		registerFunc(&_vImageConvert_16Q12to16U, frameworkHandle, "vImageConvert_16Q12to16U")
		registerFunc(&_vImageConvert_16Q12to8, frameworkHandle, "vImageConvert_16Q12to8")
		registerFunc(&_vImageConvert_16Q12toF, frameworkHandle, "vImageConvert_16Q12toF")
		registerFunc(&_vImageConvert_16SToF, frameworkHandle, "vImageConvert_16SToF")
		registerFunc(&_vImageConvert_16UTo12U, frameworkHandle, "vImageConvert_16UTo12U")
		registerFunc(&_vImageConvert_16UToF, frameworkHandle, "vImageConvert_16UToF")
		registerFunc(&_vImageConvert_16UToPlanar8, frameworkHandle, "vImageConvert_16UToPlanar8")
		registerFunc(&_vImageConvert_16Uto16F, frameworkHandle, "vImageConvert_16Uto16F")
		registerFunc(&_vImageConvert_16Uto16Q12, frameworkHandle, "vImageConvert_16Uto16Q12")
		registerFunc(&_vImageConvert_420Yp8_Cb8_Cr8ToARGB8888, frameworkHandle, "vImageConvert_420Yp8_Cb8_Cr8ToARGB8888")
		registerFunc(&_vImageConvert_420Yp8_CbCr8ToARGB8888, frameworkHandle, "vImageConvert_420Yp8_CbCr8ToARGB8888")
		registerFunc(&_vImageConvert_422CbYpCrYp16ToARGB16U, frameworkHandle, "vImageConvert_422CbYpCrYp16ToARGB16U")
		registerFunc(&_vImageConvert_422CbYpCrYp16ToARGB8888, frameworkHandle, "vImageConvert_422CbYpCrYp16ToARGB8888")
		registerFunc(&_vImageConvert_422CbYpCrYp8ToARGB8888, frameworkHandle, "vImageConvert_422CbYpCrYp8ToARGB8888")
		registerFunc(&_vImageConvert_422CbYpCrYp8_AA8ToARGB8888, frameworkHandle, "vImageConvert_422CbYpCrYp8_AA8ToARGB8888")
		registerFunc(&_vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12, frameworkHandle, "vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB16Q12")
		registerFunc(&_vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888, frameworkHandle, "vImageConvert_422CrYpCbYpCbYpCbYpCrYpCrYp10ToARGB8888")
		registerFunc(&_vImageConvert_422YpCbYpCr8ToARGB8888, frameworkHandle, "vImageConvert_422YpCbYpCr8ToARGB8888")
		registerFunc(&_vImageConvert_444AYpCbCr16ToARGB16U, frameworkHandle, "vImageConvert_444AYpCbCr16ToARGB16U")
		registerFunc(&_vImageConvert_444AYpCbCr16ToARGB8888, frameworkHandle, "vImageConvert_444AYpCbCr16ToARGB8888")
		registerFunc(&_vImageConvert_444AYpCbCr8ToARGB8888, frameworkHandle, "vImageConvert_444AYpCbCr8ToARGB8888")
		registerFunc(&_vImageConvert_444CbYpCrA8ToARGB8888, frameworkHandle, "vImageConvert_444CbYpCrA8ToARGB8888")
		registerFunc(&_vImageConvert_444CrYpCb10ToARGB16Q12, frameworkHandle, "vImageConvert_444CrYpCb10ToARGB16Q12")
		registerFunc(&_vImageConvert_444CrYpCb10ToARGB8888, frameworkHandle, "vImageConvert_444CrYpCb10ToARGB8888")
		registerFunc(&_vImageConvert_444CrYpCb8ToARGB8888, frameworkHandle, "vImageConvert_444CrYpCb8ToARGB8888")
		registerFunc(&_vImageConvert_8to16Q12, frameworkHandle, "vImageConvert_8to16Q12")
		registerFunc(&_vImageConvert_ARGB1555toARGB8888, frameworkHandle, "vImageConvert_ARGB1555toARGB8888")
		registerFunc(&_vImageConvert_ARGB1555toPlanar8, frameworkHandle, "vImageConvert_ARGB1555toPlanar8")
		registerFunc(&_vImageConvert_ARGB1555toRGB565, frameworkHandle, "vImageConvert_ARGB1555toRGB565")
		registerFunc(&_vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10, frameworkHandle, "vImageConvert_ARGB16Q12To422CrYpCbYpCbYpCbYpCrYpCrYp10")
		registerFunc(&_vImageConvert_ARGB16Q12To444CrYpCb10, frameworkHandle, "vImageConvert_ARGB16Q12To444CrYpCb10")
		registerFunc(&_vImageConvert_ARGB16Q12ToARGB2101010, frameworkHandle, "vImageConvert_ARGB16Q12ToARGB2101010")
		registerFunc(&_vImageConvert_ARGB16Q12ToRGBA1010102, frameworkHandle, "vImageConvert_ARGB16Q12ToRGBA1010102")
		registerFunc(&_vImageConvert_ARGB16Q12ToXRGB2101010, frameworkHandle, "vImageConvert_ARGB16Q12ToXRGB2101010")
		registerFunc(&_vImageConvert_ARGB16UTo422CbYpCrYp16, frameworkHandle, "vImageConvert_ARGB16UTo422CbYpCrYp16")
		registerFunc(&_vImageConvert_ARGB16UTo444AYpCbCr16, frameworkHandle, "vImageConvert_ARGB16UTo444AYpCbCr16")
		registerFunc(&_vImageConvert_ARGB16UToARGB2101010, frameworkHandle, "vImageConvert_ARGB16UToARGB2101010")
		registerFunc(&_vImageConvert_ARGB16UToARGB8888, frameworkHandle, "vImageConvert_ARGB16UToARGB8888")
		registerFunc(&_vImageConvert_ARGB16UToRGBA1010102, frameworkHandle, "vImageConvert_ARGB16UToRGBA1010102")
		registerFunc(&_vImageConvert_ARGB16UToXRGB2101010, frameworkHandle, "vImageConvert_ARGB16UToXRGB2101010")
		registerFunc(&_vImageConvert_ARGB16UtoARGB8888_dithered, frameworkHandle, "vImageConvert_ARGB16UtoARGB8888_dithered")
		registerFunc(&_vImageConvert_ARGB16UtoPlanar16U, frameworkHandle, "vImageConvert_ARGB16UtoPlanar16U")
		registerFunc(&_vImageConvert_ARGB16UtoRGB16U, frameworkHandle, "vImageConvert_ARGB16UtoRGB16U")
		registerFunc(&_vImageConvert_ARGB2101010ToARGB16F, frameworkHandle, "vImageConvert_ARGB2101010ToARGB16F")
		registerFunc(&_vImageConvert_ARGB2101010ToARGB16Q12, frameworkHandle, "vImageConvert_ARGB2101010ToARGB16Q12")
		registerFunc(&_vImageConvert_ARGB2101010ToARGB16U, frameworkHandle, "vImageConvert_ARGB2101010ToARGB16U")
		registerFunc(&_vImageConvert_ARGB2101010ToARGB8888, frameworkHandle, "vImageConvert_ARGB2101010ToARGB8888")
		registerFunc(&_vImageConvert_ARGB2101010ToARGBFFFF, frameworkHandle, "vImageConvert_ARGB2101010ToARGBFFFF")
		registerFunc(&_vImageConvert_ARGB8888To420Yp8_Cb8_Cr8, frameworkHandle, "vImageConvert_ARGB8888To420Yp8_Cb8_Cr8")
		registerFunc(&_vImageConvert_ARGB8888To420Yp8_CbCr8, frameworkHandle, "vImageConvert_ARGB8888To420Yp8_CbCr8")
		registerFunc(&_vImageConvert_ARGB8888To422CbYpCrYp16, frameworkHandle, "vImageConvert_ARGB8888To422CbYpCrYp16")
		registerFunc(&_vImageConvert_ARGB8888To422CbYpCrYp8, frameworkHandle, "vImageConvert_ARGB8888To422CbYpCrYp8")
		registerFunc(&_vImageConvert_ARGB8888To422CbYpCrYp8_AA8, frameworkHandle, "vImageConvert_ARGB8888To422CbYpCrYp8_AA8")
		registerFunc(&_vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10, frameworkHandle, "vImageConvert_ARGB8888To422CrYpCbYpCbYpCbYpCrYpCrYp10")
		registerFunc(&_vImageConvert_ARGB8888To422YpCbYpCr8, frameworkHandle, "vImageConvert_ARGB8888To422YpCbYpCr8")
		registerFunc(&_vImageConvert_ARGB8888To444AYpCbCr16, frameworkHandle, "vImageConvert_ARGB8888To444AYpCbCr16")
		registerFunc(&_vImageConvert_ARGB8888To444AYpCbCr8, frameworkHandle, "vImageConvert_ARGB8888To444AYpCbCr8")
		registerFunc(&_vImageConvert_ARGB8888To444CbYpCrA8, frameworkHandle, "vImageConvert_ARGB8888To444CbYpCrA8")
		registerFunc(&_vImageConvert_ARGB8888To444CrYpCb10, frameworkHandle, "vImageConvert_ARGB8888To444CrYpCb10")
		registerFunc(&_vImageConvert_ARGB8888To444CrYpCb8, frameworkHandle, "vImageConvert_ARGB8888To444CrYpCb8")
		registerFunc(&_vImageConvert_ARGB8888ToARGB16U, frameworkHandle, "vImageConvert_ARGB8888ToARGB16U")
		registerFunc(&_vImageConvert_ARGB8888ToARGB2101010, frameworkHandle, "vImageConvert_ARGB8888ToARGB2101010")
		registerFunc(&_vImageConvert_ARGB8888ToRGB16U, frameworkHandle, "vImageConvert_ARGB8888ToRGB16U")
		registerFunc(&_vImageConvert_ARGB8888ToRGBA1010102, frameworkHandle, "vImageConvert_ARGB8888ToRGBA1010102")
		registerFunc(&_vImageConvert_ARGB8888ToXRGB2101010, frameworkHandle, "vImageConvert_ARGB8888ToXRGB2101010")
		registerFunc(&_vImageConvert_ARGB8888toARGB1555, frameworkHandle, "vImageConvert_ARGB8888toARGB1555")
		registerFunc(&_vImageConvert_ARGB8888toARGB1555_dithered, frameworkHandle, "vImageConvert_ARGB8888toARGB1555_dithered")
		registerFunc(&_vImageConvert_ARGB8888toPlanar16Q12, frameworkHandle, "vImageConvert_ARGB8888toPlanar16Q12")
		registerFunc(&_vImageConvert_ARGB8888toPlanar8, frameworkHandle, "vImageConvert_ARGB8888toPlanar8")
		registerFunc(&_vImageConvert_ARGB8888toPlanarF, frameworkHandle, "vImageConvert_ARGB8888toPlanarF")
		registerFunc(&_vImageConvert_ARGB8888toRGB565, frameworkHandle, "vImageConvert_ARGB8888toRGB565")
		registerFunc(&_vImageConvert_ARGB8888toRGB565_dithered, frameworkHandle, "vImageConvert_ARGB8888toRGB565_dithered")
		registerFunc(&_vImageConvert_ARGB8888toRGB888, frameworkHandle, "vImageConvert_ARGB8888toRGB888")
		registerFunc(&_vImageConvert_ARGBFFFFToARGB2101010, frameworkHandle, "vImageConvert_ARGBFFFFToARGB2101010")
		registerFunc(&_vImageConvert_ARGBFFFFToXRGB2101010, frameworkHandle, "vImageConvert_ARGBFFFFToXRGB2101010")
		registerFunc(&_vImageConvert_ARGBFFFFtoARGB8888_dithered, frameworkHandle, "vImageConvert_ARGBFFFFtoARGB8888_dithered")
		registerFunc(&_vImageConvert_ARGBFFFFtoPlanar8, frameworkHandle, "vImageConvert_ARGBFFFFtoPlanar8")
		registerFunc(&_vImageConvert_ARGBFFFFtoPlanarF, frameworkHandle, "vImageConvert_ARGBFFFFtoPlanarF")
		registerFunc(&_vImageConvert_ARGBFFFFtoRGBFFF, frameworkHandle, "vImageConvert_ARGBFFFFtoRGBFFF")
		registerFunc(&_vImageConvert_ARGBToYpCbCr_GenerateConversion, frameworkHandle, "vImageConvert_ARGBToYpCbCr_GenerateConversion")
		registerFunc(&_vImageConvert_AnyToAny, frameworkHandle, "vImageConvert_AnyToAny")
		registerFunc(&_vImageConvert_BGRA16UtoRGB16U, frameworkHandle, "vImageConvert_BGRA16UtoRGB16U")
		registerFunc(&_vImageConvert_BGRA8888toRGB565, frameworkHandle, "vImageConvert_BGRA8888toRGB565")
		registerFunc(&_vImageConvert_BGRA8888toRGB565_dithered, frameworkHandle, "vImageConvert_BGRA8888toRGB565_dithered")
		registerFunc(&_vImageConvert_BGRA8888toRGB888, frameworkHandle, "vImageConvert_BGRA8888toRGB888")
		registerFunc(&_vImageConvert_BGRAFFFFtoRGBFFF, frameworkHandle, "vImageConvert_BGRAFFFFtoRGBFFF")
		registerFunc(&_vImageConvert_BGRX8888ToPlanar8, frameworkHandle, "vImageConvert_BGRX8888ToPlanar8")
		registerFunc(&_vImageConvert_BGRXFFFFToPlanarF, frameworkHandle, "vImageConvert_BGRXFFFFToPlanarF")
		registerFunc(&_vImageConvert_ChunkyToPlanar8, frameworkHandle, "vImageConvert_ChunkyToPlanar8")
		registerFunc(&_vImageConvert_ChunkyToPlanarF, frameworkHandle, "vImageConvert_ChunkyToPlanarF")
		registerFunc(&_vImageConvert_FTo16S, frameworkHandle, "vImageConvert_FTo16S")
		registerFunc(&_vImageConvert_FTo16U, frameworkHandle, "vImageConvert_FTo16U")
		registerFunc(&_vImageConvert_Fto16Q12, frameworkHandle, "vImageConvert_Fto16Q12")
		registerFunc(&_vImageConvert_Indexed1toPlanar8, frameworkHandle, "vImageConvert_Indexed1toPlanar8")
		registerFunc(&_vImageConvert_Indexed2toPlanar8, frameworkHandle, "vImageConvert_Indexed2toPlanar8")
		registerFunc(&_vImageConvert_Indexed4toPlanar8, frameworkHandle, "vImageConvert_Indexed4toPlanar8")
		registerFunc(&_vImageConvert_Planar16FtoPlanar8, frameworkHandle, "vImageConvert_Planar16FtoPlanar8")
		registerFunc(&_vImageConvert_Planar16FtoPlanarF, frameworkHandle, "vImageConvert_Planar16FtoPlanarF")
		registerFunc(&_vImageConvert_Planar16Q12toARGB16F, frameworkHandle, "vImageConvert_Planar16Q12toARGB16F")
		registerFunc(&_vImageConvert_Planar16Q12toARGB8888, frameworkHandle, "vImageConvert_Planar16Q12toARGB8888")
		registerFunc(&_vImageConvert_Planar16Q12toRGB16F, frameworkHandle, "vImageConvert_Planar16Q12toRGB16F")
		registerFunc(&_vImageConvert_Planar16Q12toRGB888, frameworkHandle, "vImageConvert_Planar16Q12toRGB888")
		registerFunc(&_vImageConvert_Planar16UtoARGB16U, frameworkHandle, "vImageConvert_Planar16UtoARGB16U")
		registerFunc(&_vImageConvert_Planar16UtoPlanar8_dithered, frameworkHandle, "vImageConvert_Planar16UtoPlanar8_dithered")
		registerFunc(&_vImageConvert_Planar16UtoRGB16U, frameworkHandle, "vImageConvert_Planar16UtoRGB16U")
		registerFunc(&_vImageConvert_Planar1toPlanar8, frameworkHandle, "vImageConvert_Planar1toPlanar8")
		registerFunc(&_vImageConvert_Planar2toPlanar8, frameworkHandle, "vImageConvert_Planar2toPlanar8")
		registerFunc(&_vImageConvert_Planar4toPlanar8, frameworkHandle, "vImageConvert_Planar4toPlanar8")
		registerFunc(&_vImageConvert_Planar8To16U, frameworkHandle, "vImageConvert_Planar8To16U")
		registerFunc(&_vImageConvert_Planar8ToARGBFFFF, frameworkHandle, "vImageConvert_Planar8ToARGBFFFF")
		registerFunc(&_vImageConvert_Planar8ToBGRX8888, frameworkHandle, "vImageConvert_Planar8ToBGRX8888")
		registerFunc(&_vImageConvert_Planar8ToBGRXFFFF, frameworkHandle, "vImageConvert_Planar8ToBGRXFFFF")
		registerFunc(&_vImageConvert_Planar8ToXRGB8888, frameworkHandle, "vImageConvert_Planar8ToXRGB8888")
		registerFunc(&_vImageConvert_Planar8ToXRGBFFFF, frameworkHandle, "vImageConvert_Planar8ToXRGBFFFF")
		registerFunc(&_vImageConvert_Planar8toARGB1555, frameworkHandle, "vImageConvert_Planar8toARGB1555")
		registerFunc(&_vImageConvert_Planar8toARGB8888, frameworkHandle, "vImageConvert_Planar8toARGB8888")
		registerFunc(&_vImageConvert_Planar8toIndexed1, frameworkHandle, "vImageConvert_Planar8toIndexed1")
		registerFunc(&_vImageConvert_Planar8toIndexed2, frameworkHandle, "vImageConvert_Planar8toIndexed2")
		registerFunc(&_vImageConvert_Planar8toIndexed4, frameworkHandle, "vImageConvert_Planar8toIndexed4")
		registerFunc(&_vImageConvert_Planar8toPlanar1, frameworkHandle, "vImageConvert_Planar8toPlanar1")
		registerFunc(&_vImageConvert_Planar8toPlanar16F, frameworkHandle, "vImageConvert_Planar8toPlanar16F")
		registerFunc(&_vImageConvert_Planar8toPlanar2, frameworkHandle, "vImageConvert_Planar8toPlanar2")
		registerFunc(&_vImageConvert_Planar8toPlanar4, frameworkHandle, "vImageConvert_Planar8toPlanar4")
		registerFunc(&_vImageConvert_Planar8toPlanarF, frameworkHandle, "vImageConvert_Planar8toPlanarF")
		registerFunc(&_vImageConvert_Planar8toRGB565, frameworkHandle, "vImageConvert_Planar8toRGB565")
		registerFunc(&_vImageConvert_Planar8toRGB888, frameworkHandle, "vImageConvert_Planar8toRGB888")
		registerFunc(&_vImageConvert_PlanarFToARGB8888, frameworkHandle, "vImageConvert_PlanarFToARGB8888")
		registerFunc(&_vImageConvert_PlanarFToBGRX8888, frameworkHandle, "vImageConvert_PlanarFToBGRX8888")
		registerFunc(&_vImageConvert_PlanarFToBGRXFFFF, frameworkHandle, "vImageConvert_PlanarFToBGRXFFFF")
		registerFunc(&_vImageConvert_PlanarFToXRGB8888, frameworkHandle, "vImageConvert_PlanarFToXRGB8888")
		registerFunc(&_vImageConvert_PlanarFToXRGBFFFF, frameworkHandle, "vImageConvert_PlanarFToXRGBFFFF")
		registerFunc(&_vImageConvert_PlanarFtoARGBFFFF, frameworkHandle, "vImageConvert_PlanarFtoARGBFFFF")
		registerFunc(&_vImageConvert_PlanarFtoPlanar16F, frameworkHandle, "vImageConvert_PlanarFtoPlanar16F")
		registerFunc(&_vImageConvert_PlanarFtoPlanar8, frameworkHandle, "vImageConvert_PlanarFtoPlanar8")
		registerFunc(&_vImageConvert_PlanarFtoPlanar8_dithered, frameworkHandle, "vImageConvert_PlanarFtoPlanar8_dithered")
		registerFunc(&_vImageConvert_PlanarFtoRGBFFF, frameworkHandle, "vImageConvert_PlanarFtoRGBFFF")
		registerFunc(&_vImageConvert_PlanarToChunky8, frameworkHandle, "vImageConvert_PlanarToChunky8")
		registerFunc(&_vImageConvert_PlanarToChunkyF, frameworkHandle, "vImageConvert_PlanarToChunkyF")
		registerFunc(&_vImageConvert_RGB16UToARGB8888, frameworkHandle, "vImageConvert_RGB16UToARGB8888")
		registerFunc(&_vImageConvert_RGB16UtoARGB16U, frameworkHandle, "vImageConvert_RGB16UtoARGB16U")
		registerFunc(&_vImageConvert_RGB16UtoBGRA16U, frameworkHandle, "vImageConvert_RGB16UtoBGRA16U")
		registerFunc(&_vImageConvert_RGB16UtoPlanar16U, frameworkHandle, "vImageConvert_RGB16UtoPlanar16U")
		registerFunc(&_vImageConvert_RGB16UtoRGB888_dithered, frameworkHandle, "vImageConvert_RGB16UtoRGB888_dithered")
		registerFunc(&_vImageConvert_RGB16UtoRGBA16U, frameworkHandle, "vImageConvert_RGB16UtoRGBA16U")
		registerFunc(&_vImageConvert_RGB565toARGB1555, frameworkHandle, "vImageConvert_RGB565toARGB1555")
		registerFunc(&_vImageConvert_RGB565toARGB8888, frameworkHandle, "vImageConvert_RGB565toARGB8888")
		registerFunc(&_vImageConvert_RGB565toBGRA8888, frameworkHandle, "vImageConvert_RGB565toBGRA8888")
		registerFunc(&_vImageConvert_RGB565toPlanar8, frameworkHandle, "vImageConvert_RGB565toPlanar8")
		registerFunc(&_vImageConvert_RGB565toRGB888, frameworkHandle, "vImageConvert_RGB565toRGB888")
		registerFunc(&_vImageConvert_RGB565toRGBA5551, frameworkHandle, "vImageConvert_RGB565toRGBA5551")
		registerFunc(&_vImageConvert_RGB565toRGBA8888, frameworkHandle, "vImageConvert_RGB565toRGBA8888")
		registerFunc(&_vImageConvert_RGB888toARGB8888, frameworkHandle, "vImageConvert_RGB888toARGB8888")
		registerFunc(&_vImageConvert_RGB888toBGRA8888, frameworkHandle, "vImageConvert_RGB888toBGRA8888")
		registerFunc(&_vImageConvert_RGB888toPlanar16Q12, frameworkHandle, "vImageConvert_RGB888toPlanar16Q12")
		registerFunc(&_vImageConvert_RGB888toPlanar8, frameworkHandle, "vImageConvert_RGB888toPlanar8")
		registerFunc(&_vImageConvert_RGB888toRGB565_dithered, frameworkHandle, "vImageConvert_RGB888toRGB565_dithered")
		registerFunc(&_vImageConvert_RGB888toRGBA8888, frameworkHandle, "vImageConvert_RGB888toRGBA8888")
		registerFunc(&_vImageConvert_RGBA1010102ToARGB16Q12, frameworkHandle, "vImageConvert_RGBA1010102ToARGB16Q12")
		registerFunc(&_vImageConvert_RGBA1010102ToARGB16U, frameworkHandle, "vImageConvert_RGBA1010102ToARGB16U")
		registerFunc(&_vImageConvert_RGBA1010102ToARGB8888, frameworkHandle, "vImageConvert_RGBA1010102ToARGB8888")
		registerFunc(&_vImageConvert_RGBA16UtoRGB16U, frameworkHandle, "vImageConvert_RGBA16UtoRGB16U")
		registerFunc(&_vImageConvert_RGBA5551toRGB565, frameworkHandle, "vImageConvert_RGBA5551toRGB565")
		registerFunc(&_vImageConvert_RGBA5551toRGBA8888, frameworkHandle, "vImageConvert_RGBA5551toRGBA8888")
		registerFunc(&_vImageConvert_RGBA8888toRGB565, frameworkHandle, "vImageConvert_RGBA8888toRGB565")
		registerFunc(&_vImageConvert_RGBA8888toRGB565_dithered, frameworkHandle, "vImageConvert_RGBA8888toRGB565_dithered")
		registerFunc(&_vImageConvert_RGBA8888toRGB888, frameworkHandle, "vImageConvert_RGBA8888toRGB888")
		registerFunc(&_vImageConvert_RGBA8888toRGBA5551, frameworkHandle, "vImageConvert_RGBA8888toRGBA5551")
		registerFunc(&_vImageConvert_RGBA8888toRGBA5551_dithered, frameworkHandle, "vImageConvert_RGBA8888toRGBA5551_dithered")
		registerFunc(&_vImageConvert_RGBAFFFFtoRGBFFF, frameworkHandle, "vImageConvert_RGBAFFFFtoRGBFFF")
		registerFunc(&_vImageConvert_RGBFFFtoARGBFFFF, frameworkHandle, "vImageConvert_RGBFFFtoARGBFFFF")
		registerFunc(&_vImageConvert_RGBFFFtoBGRAFFFF, frameworkHandle, "vImageConvert_RGBFFFtoBGRAFFFF")
		registerFunc(&_vImageConvert_RGBFFFtoPlanarF, frameworkHandle, "vImageConvert_RGBFFFtoPlanarF")
		registerFunc(&_vImageConvert_RGBFFFtoRGB888_dithered, frameworkHandle, "vImageConvert_RGBFFFtoRGB888_dithered")
		registerFunc(&_vImageConvert_RGBFFFtoRGBAFFFF, frameworkHandle, "vImageConvert_RGBFFFtoRGBAFFFF")
		registerFunc(&_vImageConvert_XRGB2101010ToARGB16F, frameworkHandle, "vImageConvert_XRGB2101010ToARGB16F")
		registerFunc(&_vImageConvert_XRGB2101010ToARGB16Q12, frameworkHandle, "vImageConvert_XRGB2101010ToARGB16Q12")
		registerFunc(&_vImageConvert_XRGB2101010ToARGB16U, frameworkHandle, "vImageConvert_XRGB2101010ToARGB16U")
		registerFunc(&_vImageConvert_XRGB2101010ToARGB8888, frameworkHandle, "vImageConvert_XRGB2101010ToARGB8888")
		registerFunc(&_vImageConvert_XRGB2101010ToARGBFFFF, frameworkHandle, "vImageConvert_XRGB2101010ToARGBFFFF")
		registerFunc(&_vImageConvert_XRGB8888ToPlanar8, frameworkHandle, "vImageConvert_XRGB8888ToPlanar8")
		registerFunc(&_vImageConvert_XRGBFFFFToPlanarF, frameworkHandle, "vImageConvert_XRGBFFFFToPlanarF")
		registerFunc(&_vImageConvert_YpCbCrToARGB_GenerateConversion, frameworkHandle, "vImageConvert_YpCbCrToARGB_GenerateConversion")
		registerFunc(&_vImageConverter_CreateForCGToCVImageFormat, frameworkHandle, "vImageConverter_CreateForCGToCVImageFormat")
		registerFunc(&_vImageConverter_CreateForCVToCGImageFormat, frameworkHandle, "vImageConverter_CreateForCVToCGImageFormat")
		registerFunc(&_vImageConverter_CreateWithCGColorConversionInfo, frameworkHandle, "vImageConverter_CreateWithCGColorConversionInfo")
		registerFunc(&_vImageConverter_CreateWithCGImageFormat, frameworkHandle, "vImageConverter_CreateWithCGImageFormat")
		registerFunc(&_vImageConverter_CreateWithColorSyncCodeFragment, frameworkHandle, "vImageConverter_CreateWithColorSyncCodeFragment")
		registerFunc(&_vImageConverter_GetNumberOfDestinationBuffers, frameworkHandle, "vImageConverter_GetNumberOfDestinationBuffers")
		registerFunc(&_vImageConverter_GetNumberOfSourceBuffers, frameworkHandle, "vImageConverter_GetNumberOfSourceBuffers")
		registerFunc(&_vImageConverter_MustOperateOutOfPlace, frameworkHandle, "vImageConverter_MustOperateOutOfPlace")
		registerFunc(&_vImageConverter_Release, frameworkHandle, "vImageConverter_Release")
		registerFunc(&_vImageConverter_Retain, frameworkHandle, "vImageConverter_Retain")
		registerFunc(&_vImageConvolveFloatKernel_ARGB8888, frameworkHandle, "vImageConvolveFloatKernel_ARGB8888")
		registerFunc(&_vImageConvolveMultiKernel_ARGB8888, frameworkHandle, "vImageConvolveMultiKernel_ARGB8888")
		registerFunc(&_vImageConvolveMultiKernel_ARGBFFFF, frameworkHandle, "vImageConvolveMultiKernel_ARGBFFFF")
		registerFunc(&_vImageConvolveWithBias_ARGB16F, frameworkHandle, "vImageConvolveWithBias_ARGB16F")
		registerFunc(&_vImageConvolveWithBias_ARGB8888, frameworkHandle, "vImageConvolveWithBias_ARGB8888")
		registerFunc(&_vImageConvolveWithBias_ARGBFFFF, frameworkHandle, "vImageConvolveWithBias_ARGBFFFF")
		registerFunc(&_vImageConvolveWithBias_Planar16F, frameworkHandle, "vImageConvolveWithBias_Planar16F")
		registerFunc(&_vImageConvolveWithBias_Planar8, frameworkHandle, "vImageConvolveWithBias_Planar8")
		registerFunc(&_vImageConvolveWithBias_PlanarF, frameworkHandle, "vImageConvolveWithBias_PlanarF")
		registerFunc(&_vImageConvolve_ARGB16F, frameworkHandle, "vImageConvolve_ARGB16F")
		registerFunc(&_vImageConvolve_ARGB8888, frameworkHandle, "vImageConvolve_ARGB8888")
		registerFunc(&_vImageConvolve_ARGBFFFF, frameworkHandle, "vImageConvolve_ARGBFFFF")
		registerFunc(&_vImageConvolve_Planar16F, frameworkHandle, "vImageConvolve_Planar16F")
		registerFunc(&_vImageConvolve_Planar8, frameworkHandle, "vImageConvolve_Planar8")
		registerFunc(&_vImageConvolve_PlanarF, frameworkHandle, "vImageConvolve_PlanarF")
		registerFunc(&_vImageCopyBuffer, frameworkHandle, "vImageCopyBuffer")
		registerFunc(&_vImageCreateCGImageFromBuffer, frameworkHandle, "vImageCreateCGImageFromBuffer")
		registerFunc(&_vImageCreateGammaFunction, frameworkHandle, "vImageCreateGammaFunction")
		registerFunc(&_vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction, frameworkHandle, "vImageCreateMonochromeColorSpaceWithWhitePointAndTransferFunction")
		registerFunc(&_vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction, frameworkHandle, "vImageCreateRGBColorSpaceWithPrimariesAndTransferFunction")
		registerFunc(&_vImageDestroyGammaFunction, frameworkHandle, "vImageDestroyGammaFunction")
		registerFunc(&_vImageDestroyResamplingFilter, frameworkHandle, "vImageDestroyResamplingFilter")
		registerFunc(&_vImageDilate_ARGB8888, frameworkHandle, "vImageDilate_ARGB8888")
		registerFunc(&_vImageDilate_ARGBFFFF, frameworkHandle, "vImageDilate_ARGBFFFF")
		registerFunc(&_vImageDilate_Planar8, frameworkHandle, "vImageDilate_Planar8")
		registerFunc(&_vImageDilate_PlanarF, frameworkHandle, "vImageDilate_PlanarF")
		registerFunc(&_vImageEndsInContrastStretch_ARGB8888, frameworkHandle, "vImageEndsInContrastStretch_ARGB8888")
		registerFunc(&_vImageEndsInContrastStretch_ARGBFFFF, frameworkHandle, "vImageEndsInContrastStretch_ARGBFFFF")
		registerFunc(&_vImageEndsInContrastStretch_Planar8, frameworkHandle, "vImageEndsInContrastStretch_Planar8")
		registerFunc(&_vImageEndsInContrastStretch_PlanarF, frameworkHandle, "vImageEndsInContrastStretch_PlanarF")
		registerFunc(&_vImageEqualization_ARGB8888, frameworkHandle, "vImageEqualization_ARGB8888")
		registerFunc(&_vImageEqualization_ARGBFFFF, frameworkHandle, "vImageEqualization_ARGBFFFF")
		registerFunc(&_vImageEqualization_Planar8, frameworkHandle, "vImageEqualization_Planar8")
		registerFunc(&_vImageEqualization_PlanarF, frameworkHandle, "vImageEqualization_PlanarF")
		registerFunc(&_vImageErode_ARGB8888, frameworkHandle, "vImageErode_ARGB8888")
		registerFunc(&_vImageErode_ARGBFFFF, frameworkHandle, "vImageErode_ARGBFFFF")
		registerFunc(&_vImageErode_Planar8, frameworkHandle, "vImageErode_Planar8")
		registerFunc(&_vImageErode_PlanarF, frameworkHandle, "vImageErode_PlanarF")
		registerFunc(&_vImageExtractChannel_ARGB16U, frameworkHandle, "vImageExtractChannel_ARGB16U")
		registerFunc(&_vImageExtractChannel_ARGB8888, frameworkHandle, "vImageExtractChannel_ARGB8888")
		registerFunc(&_vImageExtractChannel_ARGBFFFF, frameworkHandle, "vImageExtractChannel_ARGBFFFF")
		registerFunc(&_vImageFlatten_ARGB16Q12, frameworkHandle, "vImageFlatten_ARGB16Q12")
		registerFunc(&_vImageFlatten_ARGB16U, frameworkHandle, "vImageFlatten_ARGB16U")
		registerFunc(&_vImageFlatten_ARGB8888, frameworkHandle, "vImageFlatten_ARGB8888")
		registerFunc(&_vImageFlatten_ARGB8888ToRGB888, frameworkHandle, "vImageFlatten_ARGB8888ToRGB888")
		registerFunc(&_vImageFlatten_ARGBFFFF, frameworkHandle, "vImageFlatten_ARGBFFFF")
		registerFunc(&_vImageFlatten_ARGBFFFFToRGBFFF, frameworkHandle, "vImageFlatten_ARGBFFFFToRGBFFF")
		registerFunc(&_vImageFlatten_BGRA8888ToRGB888, frameworkHandle, "vImageFlatten_BGRA8888ToRGB888")
		registerFunc(&_vImageFlatten_BGRAFFFFToRGBFFF, frameworkHandle, "vImageFlatten_BGRAFFFFToRGBFFF")
		registerFunc(&_vImageFlatten_RGBA16Q12, frameworkHandle, "vImageFlatten_RGBA16Q12")
		registerFunc(&_vImageFlatten_RGBA16U, frameworkHandle, "vImageFlatten_RGBA16U")
		registerFunc(&_vImageFlatten_RGBA8888, frameworkHandle, "vImageFlatten_RGBA8888")
		registerFunc(&_vImageFlatten_RGBA8888ToRGB888, frameworkHandle, "vImageFlatten_RGBA8888ToRGB888")
		registerFunc(&_vImageFlatten_RGBAFFFF, frameworkHandle, "vImageFlatten_RGBAFFFF")
		registerFunc(&_vImageFlatten_RGBAFFFFToRGBFFF, frameworkHandle, "vImageFlatten_RGBAFFFFToRGBFFF")
		registerFunc(&_vImageFloodFill_ARGB16U, frameworkHandle, "vImageFloodFill_ARGB16U")
		registerFunc(&_vImageFloodFill_ARGB8888, frameworkHandle, "vImageFloodFill_ARGB8888")
		registerFunc(&_vImageFloodFill_Planar16U, frameworkHandle, "vImageFloodFill_Planar16U")
		registerFunc(&_vImageFloodFill_Planar8, frameworkHandle, "vImageFloodFill_Planar8")
		registerFunc(&_vImageGamma_Planar8toPlanarF, frameworkHandle, "vImageGamma_Planar8toPlanarF")
		registerFunc(&_vImageGamma_PlanarF, frameworkHandle, "vImageGamma_PlanarF")
		registerFunc(&_vImageGamma_PlanarFtoPlanar8, frameworkHandle, "vImageGamma_PlanarFtoPlanar8")
		registerFunc(&_vImageGetPerspectiveWarp, frameworkHandle, "vImageGetPerspectiveWarp")
		registerFunc(&_vImageGetResamplingFilterExtent, frameworkHandle, "vImageGetResamplingFilterExtent")
		registerFunc(&_vImageGetResamplingFilterSize, frameworkHandle, "vImageGetResamplingFilterSize")
		registerFunc(&_vImageHistogramCalculation_ARGB8888, frameworkHandle, "vImageHistogramCalculation_ARGB8888")
		registerFunc(&_vImageHistogramCalculation_ARGBFFFF, frameworkHandle, "vImageHistogramCalculation_ARGBFFFF")
		registerFunc(&_vImageHistogramCalculation_Planar8, frameworkHandle, "vImageHistogramCalculation_Planar8")
		registerFunc(&_vImageHistogramCalculation_PlanarF, frameworkHandle, "vImageHistogramCalculation_PlanarF")
		registerFunc(&_vImageHistogramSpecification_ARGB8888, frameworkHandle, "vImageHistogramSpecification_ARGB8888")
		registerFunc(&_vImageHistogramSpecification_ARGBFFFF, frameworkHandle, "vImageHistogramSpecification_ARGBFFFF")
		registerFunc(&_vImageHistogramSpecification_Planar8, frameworkHandle, "vImageHistogramSpecification_Planar8")
		registerFunc(&_vImageHistogramSpecification_PlanarF, frameworkHandle, "vImageHistogramSpecification_PlanarF")
		registerFunc(&_vImageHorizontalReflect_ARGB16F, frameworkHandle, "vImageHorizontalReflect_ARGB16F")
		registerFunc(&_vImageHorizontalReflect_ARGB16S, frameworkHandle, "vImageHorizontalReflect_ARGB16S")
		registerFunc(&_vImageHorizontalReflect_ARGB16U, frameworkHandle, "vImageHorizontalReflect_ARGB16U")
		registerFunc(&_vImageHorizontalReflect_ARGB8888, frameworkHandle, "vImageHorizontalReflect_ARGB8888")
		registerFunc(&_vImageHorizontalReflect_ARGBFFFF, frameworkHandle, "vImageHorizontalReflect_ARGBFFFF")
		registerFunc(&_vImageHorizontalReflect_CbCr16F, frameworkHandle, "vImageHorizontalReflect_CbCr16F")
		registerFunc(&_vImageHorizontalReflect_Planar16F, frameworkHandle, "vImageHorizontalReflect_Planar16F")
		registerFunc(&_vImageHorizontalReflect_Planar16U, frameworkHandle, "vImageHorizontalReflect_Planar16U")
		registerFunc(&_vImageHorizontalReflect_Planar8, frameworkHandle, "vImageHorizontalReflect_Planar8")
		registerFunc(&_vImageHorizontalReflect_PlanarF, frameworkHandle, "vImageHorizontalReflect_PlanarF")
		registerFunc(&_vImageHorizontalShearD_ARGB16F, frameworkHandle, "vImageHorizontalShearD_ARGB16F")
		registerFunc(&_vImageHorizontalShearD_ARGB16S, frameworkHandle, "vImageHorizontalShearD_ARGB16S")
		registerFunc(&_vImageHorizontalShearD_ARGB16U, frameworkHandle, "vImageHorizontalShearD_ARGB16U")
		registerFunc(&_vImageHorizontalShearD_ARGB8888, frameworkHandle, "vImageHorizontalShearD_ARGB8888")
		registerFunc(&_vImageHorizontalShearD_ARGBFFFF, frameworkHandle, "vImageHorizontalShearD_ARGBFFFF")
		registerFunc(&_vImageHorizontalShearD_CbCr16F, frameworkHandle, "vImageHorizontalShearD_CbCr16F")
		registerFunc(&_vImageHorizontalShearD_CbCr16S, frameworkHandle, "vImageHorizontalShearD_CbCr16S")
		registerFunc(&_vImageHorizontalShearD_CbCr16U, frameworkHandle, "vImageHorizontalShearD_CbCr16U")
		registerFunc(&_vImageHorizontalShearD_Planar16F, frameworkHandle, "vImageHorizontalShearD_Planar16F")
		registerFunc(&_vImageHorizontalShearD_Planar8, frameworkHandle, "vImageHorizontalShearD_Planar8")
		registerFunc(&_vImageHorizontalShearD_PlanarF, frameworkHandle, "vImageHorizontalShearD_PlanarF")
		registerFunc(&_vImageHorizontalShear_ARGB16F, frameworkHandle, "vImageHorizontalShear_ARGB16F")
		registerFunc(&_vImageHorizontalShear_ARGB16S, frameworkHandle, "vImageHorizontalShear_ARGB16S")
		registerFunc(&_vImageHorizontalShear_ARGB16U, frameworkHandle, "vImageHorizontalShear_ARGB16U")
		registerFunc(&_vImageHorizontalShear_ARGB8888, frameworkHandle, "vImageHorizontalShear_ARGB8888")
		registerFunc(&_vImageHorizontalShear_ARGBFFFF, frameworkHandle, "vImageHorizontalShear_ARGBFFFF")
		registerFunc(&_vImageHorizontalShear_CbCr16F, frameworkHandle, "vImageHorizontalShear_CbCr16F")
		registerFunc(&_vImageHorizontalShear_CbCr16S, frameworkHandle, "vImageHorizontalShear_CbCr16S")
		registerFunc(&_vImageHorizontalShear_CbCr16U, frameworkHandle, "vImageHorizontalShear_CbCr16U")
		registerFunc(&_vImageHorizontalShear_CbCr8, frameworkHandle, "vImageHorizontalShear_CbCr8")
		registerFunc(&_vImageHorizontalShear_Planar16F, frameworkHandle, "vImageHorizontalShear_Planar16F")
		registerFunc(&_vImageHorizontalShear_Planar16S, frameworkHandle, "vImageHorizontalShear_Planar16S")
		registerFunc(&_vImageHorizontalShear_Planar16U, frameworkHandle, "vImageHorizontalShear_Planar16U")
		registerFunc(&_vImageHorizontalShear_Planar8, frameworkHandle, "vImageHorizontalShear_Planar8")
		registerFunc(&_vImageHorizontalShear_PlanarF, frameworkHandle, "vImageHorizontalShear_PlanarF")
		registerFunc(&_vImageHorizontalShear_XRGB2101010W, frameworkHandle, "vImageHorizontalShear_XRGB2101010W")
		registerFunc(&_vImageInterpolatedLookupTable_PlanarF, frameworkHandle, "vImageInterpolatedLookupTable_PlanarF")
		registerFunc(&_vImageLookupTable_8to64U, frameworkHandle, "vImageLookupTable_8to64U")
		registerFunc(&_vImageLookupTable_Planar16, frameworkHandle, "vImageLookupTable_Planar16")
		registerFunc(&_vImageLookupTable_Planar8toPlanar128, frameworkHandle, "vImageLookupTable_Planar8toPlanar128")
		registerFunc(&_vImageLookupTable_Planar8toPlanar16, frameworkHandle, "vImageLookupTable_Planar8toPlanar16")
		registerFunc(&_vImageLookupTable_Planar8toPlanar24, frameworkHandle, "vImageLookupTable_Planar8toPlanar24")
		registerFunc(&_vImageLookupTable_Planar8toPlanar48, frameworkHandle, "vImageLookupTable_Planar8toPlanar48")
		registerFunc(&_vImageLookupTable_Planar8toPlanar96, frameworkHandle, "vImageLookupTable_Planar8toPlanar96")
		registerFunc(&_vImageLookupTable_Planar8toPlanarF, frameworkHandle, "vImageLookupTable_Planar8toPlanarF")
		registerFunc(&_vImageLookupTable_PlanarFtoPlanar8, frameworkHandle, "vImageLookupTable_PlanarFtoPlanar8")
		registerFunc(&_vImageMatrixMultiply_ARGB8888, frameworkHandle, "vImageMatrixMultiply_ARGB8888")
		registerFunc(&_vImageMatrixMultiply_ARGB8888ToPlanar8, frameworkHandle, "vImageMatrixMultiply_ARGB8888ToPlanar8")
		registerFunc(&_vImageMatrixMultiply_ARGBFFFF, frameworkHandle, "vImageMatrixMultiply_ARGBFFFF")
		registerFunc(&_vImageMatrixMultiply_ARGBFFFFToPlanarF, frameworkHandle, "vImageMatrixMultiply_ARGBFFFFToPlanarF")
		registerFunc(&_vImageMatrixMultiply_Planar16S, frameworkHandle, "vImageMatrixMultiply_Planar16S")
		registerFunc(&_vImageMatrixMultiply_Planar8, frameworkHandle, "vImageMatrixMultiply_Planar8")
		registerFunc(&_vImageMatrixMultiply_PlanarF, frameworkHandle, "vImageMatrixMultiply_PlanarF")
		registerFunc(&_vImageMax_ARGB8888, frameworkHandle, "vImageMax_ARGB8888")
		registerFunc(&_vImageMax_ARGBFFFF, frameworkHandle, "vImageMax_ARGBFFFF")
		registerFunc(&_vImageMax_Planar8, frameworkHandle, "vImageMax_Planar8")
		registerFunc(&_vImageMax_PlanarF, frameworkHandle, "vImageMax_PlanarF")
		registerFunc(&_vImageMin_ARGB8888, frameworkHandle, "vImageMin_ARGB8888")
		registerFunc(&_vImageMin_ARGBFFFF, frameworkHandle, "vImageMin_ARGBFFFF")
		registerFunc(&_vImageMin_Planar8, frameworkHandle, "vImageMin_Planar8")
		registerFunc(&_vImageMin_PlanarF, frameworkHandle, "vImageMin_PlanarF")
		registerFunc(&_vImageNewResamplingFilter, frameworkHandle, "vImageNewResamplingFilter")
		registerFunc(&_vImageNewResamplingFilterForFunctionUsingBuffer, frameworkHandle, "vImageNewResamplingFilterForFunctionUsingBuffer")
		registerFunc(&_vImageOverwriteChannelsWithPixel_ARGB16U, frameworkHandle, "vImageOverwriteChannelsWithPixel_ARGB16U")
		registerFunc(&_vImageOverwriteChannelsWithPixel_ARGB8888, frameworkHandle, "vImageOverwriteChannelsWithPixel_ARGB8888")
		registerFunc(&_vImageOverwriteChannelsWithPixel_ARGBFFFF, frameworkHandle, "vImageOverwriteChannelsWithPixel_ARGBFFFF")
		registerFunc(&_vImageOverwriteChannelsWithScalar_ARGB8888, frameworkHandle, "vImageOverwriteChannelsWithScalar_ARGB8888")
		registerFunc(&_vImageOverwriteChannelsWithScalar_ARGBFFFF, frameworkHandle, "vImageOverwriteChannelsWithScalar_ARGBFFFF")
		registerFunc(&_vImageOverwriteChannelsWithScalar_Planar16F, frameworkHandle, "vImageOverwriteChannelsWithScalar_Planar16F")
		registerFunc(&_vImageOverwriteChannelsWithScalar_Planar16S, frameworkHandle, "vImageOverwriteChannelsWithScalar_Planar16S")
		registerFunc(&_vImageOverwriteChannelsWithScalar_Planar16U, frameworkHandle, "vImageOverwriteChannelsWithScalar_Planar16U")
		registerFunc(&_vImageOverwriteChannelsWithScalar_Planar8, frameworkHandle, "vImageOverwriteChannelsWithScalar_Planar8")
		registerFunc(&_vImageOverwriteChannelsWithScalar_PlanarF, frameworkHandle, "vImageOverwriteChannelsWithScalar_PlanarF")
		registerFunc(&_vImageOverwriteChannels_ARGB8888, frameworkHandle, "vImageOverwriteChannels_ARGB8888")
		registerFunc(&_vImageOverwriteChannels_ARGBFFFF, frameworkHandle, "vImageOverwriteChannels_ARGBFFFF")
		registerFunc(&_vImagePNGDecompressionFilter, frameworkHandle, "vImagePNGDecompressionFilter")
		registerFunc(&_vImagePermuteChannelsWithMaskedInsert_ARGB16U, frameworkHandle, "vImagePermuteChannelsWithMaskedInsert_ARGB16U")
		registerFunc(&_vImagePermuteChannelsWithMaskedInsert_ARGB8888, frameworkHandle, "vImagePermuteChannelsWithMaskedInsert_ARGB8888")
		registerFunc(&_vImagePermuteChannelsWithMaskedInsert_ARGBFFFF, frameworkHandle, "vImagePermuteChannelsWithMaskedInsert_ARGBFFFF")
		registerFunc(&_vImagePermuteChannels_ARGB16F, frameworkHandle, "vImagePermuteChannels_ARGB16F")
		registerFunc(&_vImagePermuteChannels_ARGB16U, frameworkHandle, "vImagePermuteChannels_ARGB16U")
		registerFunc(&_vImagePermuteChannels_ARGB8888, frameworkHandle, "vImagePermuteChannels_ARGB8888")
		registerFunc(&_vImagePermuteChannels_ARGBFFFF, frameworkHandle, "vImagePermuteChannels_ARGBFFFF")
		registerFunc(&_vImagePermuteChannels_RGB888, frameworkHandle, "vImagePermuteChannels_RGB888")
		registerFunc(&_vImagePiecewiseGamma_Planar16Q12, frameworkHandle, "vImagePiecewiseGamma_Planar16Q12")
		registerFunc(&_vImagePiecewiseGamma_Planar16Q12toPlanar8, frameworkHandle, "vImagePiecewiseGamma_Planar16Q12toPlanar8")
		registerFunc(&_vImagePiecewiseGamma_Planar8, frameworkHandle, "vImagePiecewiseGamma_Planar8")
		registerFunc(&_vImagePiecewiseGamma_Planar8toPlanar16Q12, frameworkHandle, "vImagePiecewiseGamma_Planar8toPlanar16Q12")
		registerFunc(&_vImagePiecewiseGamma_Planar8toPlanarF, frameworkHandle, "vImagePiecewiseGamma_Planar8toPlanarF")
		registerFunc(&_vImagePiecewiseGamma_PlanarF, frameworkHandle, "vImagePiecewiseGamma_PlanarF")
		registerFunc(&_vImagePiecewiseGamma_PlanarFtoPlanar8, frameworkHandle, "vImagePiecewiseGamma_PlanarFtoPlanar8")
		registerFunc(&_vImagePiecewisePolynomial_Planar8toPlanarF, frameworkHandle, "vImagePiecewisePolynomial_Planar8toPlanarF")
		registerFunc(&_vImagePiecewisePolynomial_PlanarF, frameworkHandle, "vImagePiecewisePolynomial_PlanarF")
		registerFunc(&_vImagePiecewisePolynomial_PlanarFtoPlanar8, frameworkHandle, "vImagePiecewisePolynomial_PlanarFtoPlanar8")
		registerFunc(&_vImagePiecewiseRational_PlanarF, frameworkHandle, "vImagePiecewiseRational_PlanarF")
		registerFunc(&_vImagePremultipliedAlphaBlendDarken_RGBA8888, frameworkHandle, "vImagePremultipliedAlphaBlendDarken_RGBA8888")
		registerFunc(&_vImagePremultipliedAlphaBlendLighten_RGBA8888, frameworkHandle, "vImagePremultipliedAlphaBlendLighten_RGBA8888")
		registerFunc(&_vImagePremultipliedAlphaBlendMultiply_RGBA8888, frameworkHandle, "vImagePremultipliedAlphaBlendMultiply_RGBA8888")
		registerFunc(&_vImagePremultipliedAlphaBlendScreen_RGBA8888, frameworkHandle, "vImagePremultipliedAlphaBlendScreen_RGBA8888")
		registerFunc(&_vImagePremultipliedAlphaBlendWithPermute_ARGB8888, frameworkHandle, "vImagePremultipliedAlphaBlendWithPermute_ARGB8888")
		registerFunc(&_vImagePremultipliedAlphaBlendWithPermute_RGBA8888, frameworkHandle, "vImagePremultipliedAlphaBlendWithPermute_RGBA8888")
		registerFunc(&_vImagePremultipliedAlphaBlend_ARGB8888, frameworkHandle, "vImagePremultipliedAlphaBlend_ARGB8888")
		registerFunc(&_vImagePremultipliedAlphaBlend_ARGBFFFF, frameworkHandle, "vImagePremultipliedAlphaBlend_ARGBFFFF")
		registerFunc(&_vImagePremultipliedAlphaBlend_BGRA8888, frameworkHandle, "vImagePremultipliedAlphaBlend_BGRA8888")
		registerFunc(&_vImagePremultipliedAlphaBlend_BGRAFFFF, frameworkHandle, "vImagePremultipliedAlphaBlend_BGRAFFFF")
		registerFunc(&_vImagePremultipliedAlphaBlend_Planar8, frameworkHandle, "vImagePremultipliedAlphaBlend_Planar8")
		registerFunc(&_vImagePremultipliedAlphaBlend_PlanarF, frameworkHandle, "vImagePremultipliedAlphaBlend_PlanarF")
		registerFunc(&_vImagePremultipliedConstAlphaBlend_ARGB8888, frameworkHandle, "vImagePremultipliedConstAlphaBlend_ARGB8888")
		registerFunc(&_vImagePremultipliedConstAlphaBlend_ARGBFFFF, frameworkHandle, "vImagePremultipliedConstAlphaBlend_ARGBFFFF")
		registerFunc(&_vImagePremultipliedConstAlphaBlend_Planar8, frameworkHandle, "vImagePremultipliedConstAlphaBlend_Planar8")
		registerFunc(&_vImagePremultipliedConstAlphaBlend_PlanarF, frameworkHandle, "vImagePremultipliedConstAlphaBlend_PlanarF")
		registerFunc(&_vImagePremultiplyData_ARGB16Q12, frameworkHandle, "vImagePremultiplyData_ARGB16Q12")
		registerFunc(&_vImagePremultiplyData_ARGB16U, frameworkHandle, "vImagePremultiplyData_ARGB16U")
		registerFunc(&_vImagePremultiplyData_ARGB8888, frameworkHandle, "vImagePremultiplyData_ARGB8888")
		registerFunc(&_vImagePremultiplyData_ARGBFFFF, frameworkHandle, "vImagePremultiplyData_ARGBFFFF")
		registerFunc(&_vImagePremultiplyData_Planar8, frameworkHandle, "vImagePremultiplyData_Planar8")
		registerFunc(&_vImagePremultiplyData_PlanarF, frameworkHandle, "vImagePremultiplyData_PlanarF")
		registerFunc(&_vImagePremultiplyData_RGBA16F, frameworkHandle, "vImagePremultiplyData_RGBA16F")
		registerFunc(&_vImagePremultiplyData_RGBA16Q12, frameworkHandle, "vImagePremultiplyData_RGBA16Q12")
		registerFunc(&_vImagePremultiplyData_RGBA16U, frameworkHandle, "vImagePremultiplyData_RGBA16U")
		registerFunc(&_vImagePremultiplyData_RGBA8888, frameworkHandle, "vImagePremultiplyData_RGBA8888")
		registerFunc(&_vImagePremultiplyData_RGBAFFFF, frameworkHandle, "vImagePremultiplyData_RGBAFFFF")
		registerFunc(&_vImageRichardsonLucyDeConvolve_ARGB8888, frameworkHandle, "vImageRichardsonLucyDeConvolve_ARGB8888")
		registerFunc(&_vImageRichardsonLucyDeConvolve_ARGBFFFF, frameworkHandle, "vImageRichardsonLucyDeConvolve_ARGBFFFF")
		registerFunc(&_vImageRichardsonLucyDeConvolve_Planar8, frameworkHandle, "vImageRichardsonLucyDeConvolve_Planar8")
		registerFunc(&_vImageRichardsonLucyDeConvolve_PlanarF, frameworkHandle, "vImageRichardsonLucyDeConvolve_PlanarF")
		registerFunc(&_vImageRotate90_ARGB16F, frameworkHandle, "vImageRotate90_ARGB16F")
		registerFunc(&_vImageRotate90_ARGB16S, frameworkHandle, "vImageRotate90_ARGB16S")
		registerFunc(&_vImageRotate90_ARGB16U, frameworkHandle, "vImageRotate90_ARGB16U")
		registerFunc(&_vImageRotate90_ARGB8888, frameworkHandle, "vImageRotate90_ARGB8888")
		registerFunc(&_vImageRotate90_ARGBFFFF, frameworkHandle, "vImageRotate90_ARGBFFFF")
		registerFunc(&_vImageRotate90_CbCr16F, frameworkHandle, "vImageRotate90_CbCr16F")
		registerFunc(&_vImageRotate90_Planar16F, frameworkHandle, "vImageRotate90_Planar16F")
		registerFunc(&_vImageRotate90_Planar16U, frameworkHandle, "vImageRotate90_Planar16U")
		registerFunc(&_vImageRotate90_Planar8, frameworkHandle, "vImageRotate90_Planar8")
		registerFunc(&_vImageRotate90_PlanarF, frameworkHandle, "vImageRotate90_PlanarF")
		registerFunc(&_vImageRotate_ARGB16F, frameworkHandle, "vImageRotate_ARGB16F")
		registerFunc(&_vImageRotate_ARGB16S, frameworkHandle, "vImageRotate_ARGB16S")
		registerFunc(&_vImageRotate_ARGB16U, frameworkHandle, "vImageRotate_ARGB16U")
		registerFunc(&_vImageRotate_ARGB8888, frameworkHandle, "vImageRotate_ARGB8888")
		registerFunc(&_vImageRotate_ARGBFFFF, frameworkHandle, "vImageRotate_ARGBFFFF")
		registerFunc(&_vImageRotate_CbCr16F, frameworkHandle, "vImageRotate_CbCr16F")
		registerFunc(&_vImageRotate_Planar16F, frameworkHandle, "vImageRotate_Planar16F")
		registerFunc(&_vImageRotate_Planar8, frameworkHandle, "vImageRotate_Planar8")
		registerFunc(&_vImageRotate_PlanarF, frameworkHandle, "vImageRotate_PlanarF")
		registerFunc(&_vImageScale_ARGB16F, frameworkHandle, "vImageScale_ARGB16F")
		registerFunc(&_vImageScale_ARGB16S, frameworkHandle, "vImageScale_ARGB16S")
		registerFunc(&_vImageScale_ARGB16U, frameworkHandle, "vImageScale_ARGB16U")
		registerFunc(&_vImageScale_ARGB8888, frameworkHandle, "vImageScale_ARGB8888")
		registerFunc(&_vImageScale_ARGBFFFF, frameworkHandle, "vImageScale_ARGBFFFF")
		registerFunc(&_vImageScale_CbCr16F, frameworkHandle, "vImageScale_CbCr16F")
		registerFunc(&_vImageScale_CbCr16U, frameworkHandle, "vImageScale_CbCr16U")
		registerFunc(&_vImageScale_CbCr8, frameworkHandle, "vImageScale_CbCr8")
		registerFunc(&_vImageScale_Planar16F, frameworkHandle, "vImageScale_Planar16F")
		registerFunc(&_vImageScale_Planar16S, frameworkHandle, "vImageScale_Planar16S")
		registerFunc(&_vImageScale_Planar16U, frameworkHandle, "vImageScale_Planar16U")
		registerFunc(&_vImageScale_Planar8, frameworkHandle, "vImageScale_Planar8")
		registerFunc(&_vImageScale_PlanarF, frameworkHandle, "vImageScale_PlanarF")
		registerFunc(&_vImageScale_XRGB2101010W, frameworkHandle, "vImageScale_XRGB2101010W")
		registerFunc(&_vImageSelectChannels_ARGB8888, frameworkHandle, "vImageSelectChannels_ARGB8888")
		registerFunc(&_vImageSelectChannels_ARGBFFFF, frameworkHandle, "vImageSelectChannels_ARGBFFFF")
		registerFunc(&_vImageSepConvolve_ARGB8888, frameworkHandle, "vImageSepConvolve_ARGB8888")
		registerFunc(&_vImageSepConvolve_Planar16F, frameworkHandle, "vImageSepConvolve_Planar16F")
		registerFunc(&_vImageSepConvolve_Planar16U, frameworkHandle, "vImageSepConvolve_Planar16U")
		registerFunc(&_vImageSepConvolve_Planar8, frameworkHandle, "vImageSepConvolve_Planar8")
		registerFunc(&_vImageSepConvolve_Planar8to16U, frameworkHandle, "vImageSepConvolve_Planar8to16U")
		registerFunc(&_vImageSepConvolve_PlanarF, frameworkHandle, "vImageSepConvolve_PlanarF")
		registerFunc(&_vImageSymmetricPiecewiseGamma_Planar16Q12, frameworkHandle, "vImageSymmetricPiecewiseGamma_Planar16Q12")
		registerFunc(&_vImageSymmetricPiecewiseGamma_PlanarF, frameworkHandle, "vImageSymmetricPiecewiseGamma_PlanarF")
		registerFunc(&_vImageSymmetricPiecewisePolynomial_PlanarF, frameworkHandle, "vImageSymmetricPiecewisePolynomial_PlanarF")
		registerFunc(&_vImageTableLookUp_ARGB8888, frameworkHandle, "vImageTableLookUp_ARGB8888")
		registerFunc(&_vImageTableLookUp_Planar8, frameworkHandle, "vImageTableLookUp_Planar8")
		registerFunc(&_vImageTentConvolve_ARGB8888, frameworkHandle, "vImageTentConvolve_ARGB8888")
		registerFunc(&_vImageTentConvolve_Planar8, frameworkHandle, "vImageTentConvolve_Planar8")
		registerFunc(&_vImageUnpremultiplyData_ARGB16Q12, frameworkHandle, "vImageUnpremultiplyData_ARGB16Q12")
		registerFunc(&_vImageUnpremultiplyData_ARGB16U, frameworkHandle, "vImageUnpremultiplyData_ARGB16U")
		registerFunc(&_vImageUnpremultiplyData_ARGB8888, frameworkHandle, "vImageUnpremultiplyData_ARGB8888")
		registerFunc(&_vImageUnpremultiplyData_ARGBFFFF, frameworkHandle, "vImageUnpremultiplyData_ARGBFFFF")
		registerFunc(&_vImageUnpremultiplyData_Planar8, frameworkHandle, "vImageUnpremultiplyData_Planar8")
		registerFunc(&_vImageUnpremultiplyData_PlanarF, frameworkHandle, "vImageUnpremultiplyData_PlanarF")
		registerFunc(&_vImageUnpremultiplyData_RGBA16F, frameworkHandle, "vImageUnpremultiplyData_RGBA16F")
		registerFunc(&_vImageUnpremultiplyData_RGBA16Q12, frameworkHandle, "vImageUnpremultiplyData_RGBA16Q12")
		registerFunc(&_vImageUnpremultiplyData_RGBA16U, frameworkHandle, "vImageUnpremultiplyData_RGBA16U")
		registerFunc(&_vImageUnpremultiplyData_RGBA8888, frameworkHandle, "vImageUnpremultiplyData_RGBA8888")
		registerFunc(&_vImageUnpremultiplyData_RGBAFFFF, frameworkHandle, "vImageUnpremultiplyData_RGBAFFFF")
		registerFunc(&_vImageVerticalReflect_ARGB16F, frameworkHandle, "vImageVerticalReflect_ARGB16F")
		registerFunc(&_vImageVerticalReflect_ARGB16S, frameworkHandle, "vImageVerticalReflect_ARGB16S")
		registerFunc(&_vImageVerticalReflect_ARGB16U, frameworkHandle, "vImageVerticalReflect_ARGB16U")
		registerFunc(&_vImageVerticalReflect_ARGB8888, frameworkHandle, "vImageVerticalReflect_ARGB8888")
		registerFunc(&_vImageVerticalReflect_ARGBFFFF, frameworkHandle, "vImageVerticalReflect_ARGBFFFF")
		registerFunc(&_vImageVerticalReflect_CbCr16F, frameworkHandle, "vImageVerticalReflect_CbCr16F")
		registerFunc(&_vImageVerticalReflect_Planar16F, frameworkHandle, "vImageVerticalReflect_Planar16F")
		registerFunc(&_vImageVerticalReflect_Planar16U, frameworkHandle, "vImageVerticalReflect_Planar16U")
		registerFunc(&_vImageVerticalReflect_Planar8, frameworkHandle, "vImageVerticalReflect_Planar8")
		registerFunc(&_vImageVerticalReflect_PlanarF, frameworkHandle, "vImageVerticalReflect_PlanarF")
		registerFunc(&_vImageVerticalShearD_ARGB16F, frameworkHandle, "vImageVerticalShearD_ARGB16F")
		registerFunc(&_vImageVerticalShearD_ARGB16S, frameworkHandle, "vImageVerticalShearD_ARGB16S")
		registerFunc(&_vImageVerticalShearD_ARGB16U, frameworkHandle, "vImageVerticalShearD_ARGB16U")
		registerFunc(&_vImageVerticalShearD_ARGB8888, frameworkHandle, "vImageVerticalShearD_ARGB8888")
		registerFunc(&_vImageVerticalShearD_ARGBFFFF, frameworkHandle, "vImageVerticalShearD_ARGBFFFF")
		registerFunc(&_vImageVerticalShearD_CbCr16F, frameworkHandle, "vImageVerticalShearD_CbCr16F")
		registerFunc(&_vImageVerticalShearD_CbCr16S, frameworkHandle, "vImageVerticalShearD_CbCr16S")
		registerFunc(&_vImageVerticalShearD_CbCr16U, frameworkHandle, "vImageVerticalShearD_CbCr16U")
		registerFunc(&_vImageVerticalShearD_Planar16F, frameworkHandle, "vImageVerticalShearD_Planar16F")
		registerFunc(&_vImageVerticalShearD_Planar8, frameworkHandle, "vImageVerticalShearD_Planar8")
		registerFunc(&_vImageVerticalShearD_PlanarF, frameworkHandle, "vImageVerticalShearD_PlanarF")
		registerFunc(&_vImageVerticalShear_ARGB16F, frameworkHandle, "vImageVerticalShear_ARGB16F")
		registerFunc(&_vImageVerticalShear_ARGB16S, frameworkHandle, "vImageVerticalShear_ARGB16S")
		registerFunc(&_vImageVerticalShear_ARGB16U, frameworkHandle, "vImageVerticalShear_ARGB16U")
		registerFunc(&_vImageVerticalShear_ARGB8888, frameworkHandle, "vImageVerticalShear_ARGB8888")
		registerFunc(&_vImageVerticalShear_ARGBFFFF, frameworkHandle, "vImageVerticalShear_ARGBFFFF")
		registerFunc(&_vImageVerticalShear_CbCr16F, frameworkHandle, "vImageVerticalShear_CbCr16F")
		registerFunc(&_vImageVerticalShear_CbCr16S, frameworkHandle, "vImageVerticalShear_CbCr16S")
		registerFunc(&_vImageVerticalShear_CbCr16U, frameworkHandle, "vImageVerticalShear_CbCr16U")
		registerFunc(&_vImageVerticalShear_CbCr8, frameworkHandle, "vImageVerticalShear_CbCr8")
		registerFunc(&_vImageVerticalShear_Planar16F, frameworkHandle, "vImageVerticalShear_Planar16F")
		registerFunc(&_vImageVerticalShear_Planar16S, frameworkHandle, "vImageVerticalShear_Planar16S")
		registerFunc(&_vImageVerticalShear_Planar16U, frameworkHandle, "vImageVerticalShear_Planar16U")
		registerFunc(&_vImageVerticalShear_Planar8, frameworkHandle, "vImageVerticalShear_Planar8")
		registerFunc(&_vImageVerticalShear_PlanarF, frameworkHandle, "vImageVerticalShear_PlanarF")
		registerFunc(&_vImageVerticalShear_XRGB2101010W, frameworkHandle, "vImageVerticalShear_XRGB2101010W")
		registerFunc(&_vL1024Rotate, frameworkHandle, "vL1024Rotate")
		registerFunc(&_vL256Rotate, frameworkHandle, "vL256Rotate")
		registerFunc(&_vL512Rotate, frameworkHandle, "vL512Rotate")
		registerFunc(&_vLL1024Shift, frameworkHandle, "vLL1024Shift")
		registerFunc(&_vLL256Shift, frameworkHandle, "vLL256Shift")
		registerFunc(&_vLL512Shift, frameworkHandle, "vLL512Shift")
		registerFunc(&_vLR1024Shift, frameworkHandle, "vLR1024Shift")
		registerFunc(&_vLR256Shift, frameworkHandle, "vLR256Shift")
		registerFunc(&_vLR512Shift, frameworkHandle, "vLR512Shift")
		registerFunc(&_vR1024Rotate, frameworkHandle, "vR1024Rotate")
		registerFunc(&_vR256Rotate, frameworkHandle, "vR256Rotate")
		registerFunc(&_vR512Rotate, frameworkHandle, "vR512Rotate")
		registerFunc(&_vS1024Add, frameworkHandle, "vS1024Add")
		registerFunc(&_vS1024AddS, frameworkHandle, "vS1024AddS")
		registerFunc(&_vS1024Divide, frameworkHandle, "vS1024Divide")
		registerFunc(&_vS1024HalfMultiply, frameworkHandle, "vS1024HalfMultiply")
		registerFunc(&_vS1024Mod, frameworkHandle, "vS1024Mod")
		registerFunc(&_vS1024Neg, frameworkHandle, "vS1024Neg")
		registerFunc(&_vS1024Sub, frameworkHandle, "vS1024Sub")
		registerFunc(&_vS1024SubS, frameworkHandle, "vS1024SubS")
		registerFunc(&_vS128FullMultiply, frameworkHandle, "vS128FullMultiply")
		registerFunc(&_vS256Add, frameworkHandle, "vS256Add")
		registerFunc(&_vS256AddS, frameworkHandle, "vS256AddS")
		registerFunc(&_vS256Divide, frameworkHandle, "vS256Divide")
		registerFunc(&_vS256FullMultiply, frameworkHandle, "vS256FullMultiply")
		registerFunc(&_vS256HalfMultiply, frameworkHandle, "vS256HalfMultiply")
		registerFunc(&_vS256Mod, frameworkHandle, "vS256Mod")
		registerFunc(&_vS256Neg, frameworkHandle, "vS256Neg")
		registerFunc(&_vS256Sub, frameworkHandle, "vS256Sub")
		registerFunc(&_vS256SubS, frameworkHandle, "vS256SubS")
		registerFunc(&_vS512Add, frameworkHandle, "vS512Add")
		registerFunc(&_vS512AddS, frameworkHandle, "vS512AddS")
		registerFunc(&_vS512Divide, frameworkHandle, "vS512Divide")
		registerFunc(&_vS512FullMultiply, frameworkHandle, "vS512FullMultiply")
		registerFunc(&_vS512HalfMultiply, frameworkHandle, "vS512HalfMultiply")
		registerFunc(&_vS512Mod, frameworkHandle, "vS512Mod")
		registerFunc(&_vS512Neg, frameworkHandle, "vS512Neg")
		registerFunc(&_vS512Sub, frameworkHandle, "vS512Sub")
		registerFunc(&_vS512SubS, frameworkHandle, "vS512SubS")
		registerFunc(&_vU1024Add, frameworkHandle, "vU1024Add")
		registerFunc(&_vU1024AddS, frameworkHandle, "vU1024AddS")
		registerFunc(&_vU1024Divide, frameworkHandle, "vU1024Divide")
		registerFunc(&_vU1024HalfMultiply, frameworkHandle, "vU1024HalfMultiply")
		registerFunc(&_vU1024Mod, frameworkHandle, "vU1024Mod")
		registerFunc(&_vU1024Neg, frameworkHandle, "vU1024Neg")
		registerFunc(&_vU1024Sub, frameworkHandle, "vU1024Sub")
		registerFunc(&_vU1024SubS, frameworkHandle, "vU1024SubS")
		registerFunc(&_vU128FullMultiply, frameworkHandle, "vU128FullMultiply")
		registerFunc(&_vU256Add, frameworkHandle, "vU256Add")
		registerFunc(&_vU256AddS, frameworkHandle, "vU256AddS")
		registerFunc(&_vU256Divide, frameworkHandle, "vU256Divide")
		registerFunc(&_vU256FullMultiply, frameworkHandle, "vU256FullMultiply")
		registerFunc(&_vU256HalfMultiply, frameworkHandle, "vU256HalfMultiply")
		registerFunc(&_vU256Mod, frameworkHandle, "vU256Mod")
		registerFunc(&_vU256Neg, frameworkHandle, "vU256Neg")
		registerFunc(&_vU256Sub, frameworkHandle, "vU256Sub")
		registerFunc(&_vU256SubS, frameworkHandle, "vU256SubS")
		registerFunc(&_vU512Add, frameworkHandle, "vU512Add")
		registerFunc(&_vU512AddS, frameworkHandle, "vU512AddS")
		registerFunc(&_vU512Divide, frameworkHandle, "vU512Divide")
		registerFunc(&_vU512FullMultiply, frameworkHandle, "vU512FullMultiply")
		registerFunc(&_vU512HalfMultiply, frameworkHandle, "vU512HalfMultiply")
		registerFunc(&_vU512Mod, frameworkHandle, "vU512Mod")
		registerFunc(&_vU512Neg, frameworkHandle, "vU512Neg")
		registerFunc(&_vU512Sub, frameworkHandle, "vU512Sub")
		registerFunc(&_vU512SubS, frameworkHandle, "vU512SubS")
		registerFunc(&_vvacos, frameworkHandle, "vvacos")
		registerFunc(&_vvacosf, frameworkHandle, "vvacosf")
		registerFunc(&_vvacosh, frameworkHandle, "vvacosh")
		registerFunc(&_vvacoshf, frameworkHandle, "vvacoshf")
		registerFunc(&_vvasin, frameworkHandle, "vvasin")
		registerFunc(&_vvasinf, frameworkHandle, "vvasinf")
		registerFunc(&_vvasinh, frameworkHandle, "vvasinh")
		registerFunc(&_vvasinhf, frameworkHandle, "vvasinhf")
		registerFunc(&_vvatan, frameworkHandle, "vvatan")
		registerFunc(&_vvatan2, frameworkHandle, "vvatan2")
		registerFunc(&_vvatan2f, frameworkHandle, "vvatan2f")
		registerFunc(&_vvatanf, frameworkHandle, "vvatanf")
		registerFunc(&_vvatanh, frameworkHandle, "vvatanh")
		registerFunc(&_vvatanhf, frameworkHandle, "vvatanhf")
		registerFunc(&_vvcbrt, frameworkHandle, "vvcbrt")
		registerFunc(&_vvcbrtf, frameworkHandle, "vvcbrtf")
		registerFunc(&_vvceil, frameworkHandle, "vvceil")
		registerFunc(&_vvceilf, frameworkHandle, "vvceilf")
		registerFunc(&_vvcopysign, frameworkHandle, "vvcopysign")
		registerFunc(&_vvcopysignf, frameworkHandle, "vvcopysignf")
		registerFunc(&_vvcos, frameworkHandle, "vvcos")
		registerFunc(&_vvcosf, frameworkHandle, "vvcosf")
		registerFunc(&_vvcosh, frameworkHandle, "vvcosh")
		registerFunc(&_vvcoshf, frameworkHandle, "vvcoshf")
		registerFunc(&_vvcosisin, frameworkHandle, "vvcosisin")
		registerFunc(&_vvcosisinf, frameworkHandle, "vvcosisinf")
		registerFunc(&_vvcospi, frameworkHandle, "vvcospi")
		registerFunc(&_vvcospif, frameworkHandle, "vvcospif")
		registerFunc(&_vvdiv, frameworkHandle, "vvdiv")
		registerFunc(&_vvdivf, frameworkHandle, "vvdivf")
		registerFunc(&_vvexp, frameworkHandle, "vvexp")
		registerFunc(&_vvexp2, frameworkHandle, "vvexp2")
		registerFunc(&_vvexp2f, frameworkHandle, "vvexp2f")
		registerFunc(&_vvexpf, frameworkHandle, "vvexpf")
		registerFunc(&_vvexpm1, frameworkHandle, "vvexpm1")
		registerFunc(&_vvexpm1f, frameworkHandle, "vvexpm1f")
		registerFunc(&_vvfabs, frameworkHandle, "vvfabs")
		registerFunc(&_vvfabsf, frameworkHandle, "vvfabsf")
		registerFunc(&_vvfloor, frameworkHandle, "vvfloor")
		registerFunc(&_vvfloorf, frameworkHandle, "vvfloorf")
		registerFunc(&_vvfmod, frameworkHandle, "vvfmod")
		registerFunc(&_vvfmodf, frameworkHandle, "vvfmodf")
		registerFunc(&_vvint, frameworkHandle, "vvint")
		registerFunc(&_vvintf, frameworkHandle, "vvintf")
		registerFunc(&_vvlog, frameworkHandle, "vvlog")
		registerFunc(&_vvlog10, frameworkHandle, "vvlog10")
		registerFunc(&_vvlog10f, frameworkHandle, "vvlog10f")
		registerFunc(&_vvlog1p, frameworkHandle, "vvlog1p")
		registerFunc(&_vvlog1pf, frameworkHandle, "vvlog1pf")
		registerFunc(&_vvlog2, frameworkHandle, "vvlog2")
		registerFunc(&_vvlog2f, frameworkHandle, "vvlog2f")
		registerFunc(&_vvlogb, frameworkHandle, "vvlogb")
		registerFunc(&_vvlogbf, frameworkHandle, "vvlogbf")
		registerFunc(&_vvlogf, frameworkHandle, "vvlogf")
		registerFunc(&_vvnextafter, frameworkHandle, "vvnextafter")
		registerFunc(&_vvnextafterf, frameworkHandle, "vvnextafterf")
		registerFunc(&_vvnint, frameworkHandle, "vvnint")
		registerFunc(&_vvnintf, frameworkHandle, "vvnintf")
		registerFunc(&_vvpow, frameworkHandle, "vvpow")
		registerFunc(&_vvpowf, frameworkHandle, "vvpowf")
		registerFunc(&_vvpows, frameworkHandle, "vvpows")
		registerFunc(&_vvpowsf, frameworkHandle, "vvpowsf")
		registerFunc(&_vvrec, frameworkHandle, "vvrec")
		registerFunc(&_vvrecf, frameworkHandle, "vvrecf")
		registerFunc(&_vvremainder, frameworkHandle, "vvremainder")
		registerFunc(&_vvremainderf, frameworkHandle, "vvremainderf")
		registerFunc(&_vvrsqrt, frameworkHandle, "vvrsqrt")
		registerFunc(&_vvrsqrtf, frameworkHandle, "vvrsqrtf")
		registerFunc(&_vvsin, frameworkHandle, "vvsin")
		registerFunc(&_vvsincos, frameworkHandle, "vvsincos")
		registerFunc(&_vvsincosf, frameworkHandle, "vvsincosf")
		registerFunc(&_vvsinf, frameworkHandle, "vvsinf")
		registerFunc(&_vvsinh, frameworkHandle, "vvsinh")
		registerFunc(&_vvsinhf, frameworkHandle, "vvsinhf")
		registerFunc(&_vvsinpi, frameworkHandle, "vvsinpi")
		registerFunc(&_vvsinpif, frameworkHandle, "vvsinpif")
		registerFunc(&_vvsqrt, frameworkHandle, "vvsqrt")
		registerFunc(&_vvsqrtf, frameworkHandle, "vvsqrtf")
		registerFunc(&_vvtan, frameworkHandle, "vvtan")
		registerFunc(&_vvtanf, frameworkHandle, "vvtanf")
		registerFunc(&_vvtanh, frameworkHandle, "vvtanh")
		registerFunc(&_vvtanhf, frameworkHandle, "vvtanhf")
		registerFunc(&_vvtanpi, frameworkHandle, "vvtanpi")
		registerFunc(&_vvtanpif, frameworkHandle, "vvtanpif")
	}





