// Code generated from Apple documentation. DO NOT EDIT.

package accelerate

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

type BLASParamErrorProc = func(*byte, *byte, *int, *int)







type BNNSAlloc = func(unsafe.Pointer, uint, uint) int







// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
type BNNSFilter = unsafe.Pointer







type BNNSFree = func(unsafe.Pointer)







type BNNSNearestNeighbors = unsafe.Pointer







type BNNSRandomGenerator = unsafe.Pointer







type COMPLEX = DSPComplex







type COMPLEX_SPLIT = DSPSplitComplex







type DOUBLE_COMPLEX = DSPDoubleComplex







type DOUBLE_COMPLEX_SPLIT = DSPDoubleSplitComplex







type FFTDirection = int







type FFTRadix = int







type FFTSetup = uintptr







type FFTSetupD = uintptr







type GammaFunction = unsafe.Pointer







type Pixel_16F = uint16







type Pixel_16F16F = unsafe.Pointer







type Pixel_16Q12 = int16







type Pixel_16S = int16







type Pixel_16S16S = unsafe.Pointer







type Pixel_16U = uint16







type Pixel_16U16U = unsafe.Pointer







type Pixel_32U = uint32







type Pixel_8 = uint8







type Pixel_88 = unsafe.Pointer







type Pixel_8888 = unsafe.Pointer







type Pixel_ARGB_16F = unsafe.Pointer







type Pixel_ARGB_16S = unsafe.Pointer







type Pixel_ARGB_16U = unsafe.Pointer







type Pixel_F = float32







type Pixel_FF = unsafe.Pointer







type Pixel_FFFF = unsafe.Pointer







type ResamplingFilter = unsafe.Pointer







type Bnns_graph_compile_message_fn_t = func(unsafe.Pointer, *byte, *byte, *Bnns_user_message_data_t)







type Bnns_graph_execute_message_fn_t = func(unsafe.Pointer, *byte, *byte, *Bnns_user_message_data_t)







type Bnns_graph_free_all_fn_t = func(unsafe.Pointer, uint)







type Bnns_graph_realloc_fn_t = func(unsafe.Pointer, uint, unsafe.Pointer, uint, uint) int







type La_attribute_t = uint







type La_count_t = uint







type La_deallocator_t = func(unsafe.Pointer)







type La_hint_t = uint







type La_index_t = int







type La_norm_t = uint







type La_object_t = objectivec.Object







type La_scalar_type_t = uint







type La_status_t = int







type Quadrature_function_array = func(unsafe.Pointer, uint, []float64, []float64)







type Simd_bool = bool







type Simd_char1 = int8







type Simd_char16 = int8







type Simd_char2 = int8







type Simd_char3 = int8







type Simd_char32 = int8







type Simd_char4 = int8







type Simd_char64 = int8







type Simd_char8 = int8







type Simd_double1 = float64







type Simd_double2 = float64







type Simd_double3 = float64







type Simd_double4 = float64







type Simd_double8 = float64







type Simd_float1 = float32







type Simd_float16 = float32







type Simd_float2 = float32







type Simd_float3 = float32







type Simd_float4 = float32







type Simd_float8 = float32







type Simd_half1 = unsafe.Pointer







type Simd_half16 = unsafe.Pointer







type Simd_half2 = unsafe.Pointer







type Simd_half3 = unsafe.Pointer







type Simd_half32 = unsafe.Pointer







type Simd_half4 = unsafe.Pointer







type Simd_half8 = unsafe.Pointer







type Simd_int1 = int







type Simd_int16 = int







type Simd_int2 = int







type Simd_int3 = int







type Simd_int4 = int







type Simd_int8 = int







type Simd_long1 = int







type Simd_long2 = int







type Simd_long3 = int







type Simd_long4 = int







type Simd_long8 = int







type Simd_packed_char16 = int8







type Simd_packed_char2 = int8







type Simd_packed_char32 = int8







type Simd_packed_char4 = int8







type Simd_packed_char64 = int8







type Simd_packed_char8 = int8







type Simd_packed_double2 = float64







type Simd_packed_double4 = float64







type Simd_packed_double8 = float64







type Simd_packed_float16 = float32







type Simd_packed_float2 = float32







type Simd_packed_float4 = float32







type Simd_packed_float8 = float32







type Simd_packed_int16 = int







type Simd_packed_int2 = int







type Simd_packed_int4 = int







type Simd_packed_int8 = int







type Simd_packed_long2 = int







type Simd_packed_long4 = int







type Simd_packed_long8 = int







type Simd_packed_short16 = int16







type Simd_packed_short2 = int16







type Simd_packed_short32 = int16







type Simd_packed_short4 = int16







type Simd_packed_short8 = int16







type Simd_packed_uchar16 = uint8







type Simd_packed_uchar2 = uint8







type Simd_packed_uchar32 = uint8







type Simd_packed_uchar4 = uint8







type Simd_packed_uchar64 = uint8







type Simd_packed_uchar8 = uint8







type Simd_packed_uint16 = uint







type Simd_packed_uint2 = uint







type Simd_packed_uint4 = uint







type Simd_packed_uint8 = uint







type Simd_packed_ulong2 = uint







type Simd_packed_ulong4 = uint







type Simd_packed_ulong8 = uint







type Simd_packed_ushort16 = uint16







type Simd_packed_ushort2 = uint16







type Simd_packed_ushort32 = uint16







type Simd_packed_ushort4 = uint16







type Simd_packed_ushort8 = uint16







type Simd_short1 = int16







type Simd_short16 = int16







type Simd_short2 = int16







type Simd_short3 = int16







type Simd_short32 = int16







type Simd_short4 = int16







type Simd_short8 = int16







type Simd_uchar1 = uint8







type Simd_uchar16 = uint8







type Simd_uchar2 = uint8







type Simd_uchar3 = uint8







type Simd_uchar32 = uint8







type Simd_uchar4 = uint8







type Simd_uchar64 = uint8







type Simd_uchar8 = uint8







type Simd_uint1 = uint







type Simd_uint16 = uint







type Simd_uint2 = uint







type Simd_uint3 = uint







type Simd_uint4 = uint







type Simd_uint8 = uint







type Simd_ulong1 = uint







type Simd_ulong2 = uint







type Simd_ulong3 = uint







type Simd_ulong4 = uint







type Simd_ulong8 = uint







type Simd_ushort1 = uint16







type Simd_ushort16 = uint16







type Simd_ushort2 = uint16







type Simd_ushort3 = uint16







type Simd_ushort32 = uint16







type Simd_ushort4 = uint16







type Simd_ushort8 = uint16







type Sparse_dimension = uint64







type Sparse_index = int64







type Sparse_matrix_double = uintptr







type Sparse_matrix_double_complex = uintptr







type Sparse_matrix_float = uintptr







type Sparse_matrix_float_complex = uintptr







type Sparse_stride = int64







type VBool32 = unsafe.Pointer







type VDSP_DFT_Interleaved_Setup = uintptr







type VDSP_DFT_Interleaved_SetupD = uintptr







type VDSP_DFT_Setup = uintptr







type VDSP_DFT_SetupD = uintptr







type VDSP_Length = uint







type VDSP_Stride = int







type VDSP_biquad_Setup = uintptr







type VDSP_biquad_SetupD = uintptr







type VDSP_biquadm_Setup = uintptr







type VDSP_biquadm_SetupD = uintptr







type VDouble = unsafe.Pointer







type VFloat = unsafe.Pointer







type VFloatPacked = unsafe.Pointer







type VImageBufferTypeCode = uint32







type VImageCVImageFormatError = int







type VImageCVImageFormatRef uintptr







type VImageConstCVImageFormatRef = unsafe.Pointer







type VImageConverterRef = unsafe.Pointer







type VImageMatrixType = uint32







type VImagePixelCount = uint







type VImage_CGAffineTransform = VImage_AffineTransform_Double







type VImage_Error = int







type VImage_Flags = uint32







type VImage_MultidimensionalTable = uintptr







type VImage_WarpInterpolation = int32







type VSInt16 = unsafe.Pointer







type VSInt32 = unsafe.Pointer







type VSInt64 = unsafe.Pointer







type VSInt8 = unsafe.Pointer







type VUInt16 = unsafe.Pointer







type VUInt32 = unsafe.Pointer







type VUInt64 = unsafe.Pointer







type VUInt8 = unsafe.Pointer





