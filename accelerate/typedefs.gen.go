// Code generated from Apple documentation. DO NOT EDIT.

package accelerate

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// BLASParamErrorProc is a BLAS error handler callback type.
//
// See: https://developer.apple.com/documentation/Accelerate/BLASParamErrorProc
type BLASParamErrorProc = func(*byte, *byte, *int, *int)

// BNNSAlloc is a type-alias for a user-provided memory allocation function.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSAlloc
type BNNSAlloc = func(unsafe.Pointer, uint, uint) int

// BNNSFilter is an opaque type that represents a filter.
//
// Deprecated: Deprecated since macOS 15.0. Use BNNSGraph* APIs
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFilter
type BNNSFilter = unsafe.Pointer

// BNNSFree is a type-alias for a user-provided memory deallocation function.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSFree
type BNNSFree = func(unsafe.Pointer)

// BNNSNearestNeighbors is a k-nearest neighbors object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSNearestNeighbors
type BNNSNearestNeighbors = unsafe.Pointer

// BNNSRandomGenerator is a pointer to a random number generator object.
//
// See: https://developer.apple.com/documentation/Accelerate/BNNSRandomGenerator
type BNNSRandomGenerator = unsafe.Pointer

// See: https://developer.apple.com/documentation/Accelerate/COMPLEX
type COMPLEX = DSPComplex

// See: https://developer.apple.com/documentation/Accelerate/COMPLEX_SPLIT
type COMPLEX_SPLIT = DSPSplitComplex

// See: https://developer.apple.com/documentation/Accelerate/DOUBLE_COMPLEX
type DOUBLE_COMPLEX = DSPDoubleComplex

// See: https://developer.apple.com/documentation/Accelerate/DOUBLE_COMPLEX_SPLIT
type DOUBLE_COMPLEX_SPLIT = DSPDoubleSplitComplex

// FFTDirection is constants that specify whether to perform a forward or inverse FFT.
//
// See: https://developer.apple.com/documentation/Accelerate/FFTDirection
type FFTDirection = int

// FFTRadix is the radix of the FFT decomposition.
//
// See: https://developer.apple.com/documentation/Accelerate/FFTRadix
type FFTRadix = int

// FFTSetup is an opaque type that contains setup information for a single-precision FFT transform.
//
// See: https://developer.apple.com/documentation/Accelerate/FFTSetup
type FFTSetup = uintptr

// FFTSetupD is an opaque type that contains setup information for a double-precision FFT transform.
//
// See: https://developer.apple.com/documentation/Accelerate/FFTSetupD
type FFTSetupD = uintptr

// GammaFunction is a type for a gamma function.
//
// See: https://developer.apple.com/documentation/Accelerate/GammaFunction
type GammaFunction = unsafe.Pointer

// See: https://developer.apple.com/documentation/Accelerate/Pixel_16F
type Pixel_16F = uint16

// See: https://developer.apple.com/documentation/Accelerate/Pixel_16F16F
type Pixel_16F16F = unsafe.Pointer

// Pixel_16Q12 is a type for a signed 16-bit, fixed-point number with 12 bits of fractional precision.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_16Q12
type Pixel_16Q12 = int16

// Pixel_16S is a type for a planar, 16-bits-per-channel, signed pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_16S
type Pixel_16S = int16

// See: https://developer.apple.com/documentation/Accelerate/Pixel_16S16S
type Pixel_16S16S = unsafe.Pointer

// Pixel_16U is a type for a planar, 16-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_16U
type Pixel_16U = uint16

// Pixel_16U16U is a type for a two-channel, 16-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_16U16U
type Pixel_16U16U = unsafe.Pointer

// Pixel_32U is a type you use for the XRGB2101010 format.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_32U
type Pixel_32U = uint32

// Pixel_8 is a type for a planar, 8-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_8
type Pixel_8 = uint8

// Pixel_88 is a type for a two-channel, 8-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_88
type Pixel_88 = unsafe.Pointer

// Pixel_8888 is a type for a four-channel, 8-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_8888
type Pixel_8888 = unsafe.Pointer

// See: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16F
type Pixel_ARGB_16F = unsafe.Pointer

// Pixel_ARGB_16S is a type for a four-channel, 16-bits-per-channel, signed pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16S
type Pixel_ARGB_16S = unsafe.Pointer

// Pixel_ARGB_16U is a type for a four-channel, 16-bits-per-channel, unsigned pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16U
type Pixel_ARGB_16U = unsafe.Pointer

// Pixel_F is a type for a planar, 32-bits-per-channel, floating-point pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_F
type Pixel_F = float32

// See: https://developer.apple.com/documentation/Accelerate/Pixel_FF
type Pixel_FF = unsafe.Pointer

// Pixel_FFFF is a type for a four-channel, 32-bits-per-channel, floating-point pixel.
//
// See: https://developer.apple.com/documentation/Accelerate/Pixel_FFFF
type Pixel_FFFF = unsafe.Pointer

// ResamplingFilter is a pointer to a resampling filter callback function.
//
// See: https://developer.apple.com/documentation/Accelerate/ResamplingFilter
type ResamplingFilter = unsafe.Pointer

// Bnns_graph_compile_message_fn_t is the graph compile-message logging callback function.
//
// See: https://developer.apple.com/documentation/Accelerate/bnns_graph_compile_message_fn_t
type Bnns_graph_compile_message_fn_t = func(BNNSGraphMessageLevel, *byte, *byte, *Bnns_user_message_data_t)

// Bnns_graph_execute_message_fn_t is the graph execute-message logging callback function.
//
// See: https://developer.apple.com/documentation/Accelerate/bnns_graph_execute_message_fn_t
type Bnns_graph_execute_message_fn_t = func(BNNSGraphMessageLevel, *byte, *byte, *Bnns_user_message_data_t)

// Bnns_graph_free_all_fn_t is the workspace and output deallocation function.
//
// See: https://developer.apple.com/documentation/Accelerate/bnns_graph_free_all_fn_t
type Bnns_graph_free_all_fn_t = func(unsafe.Pointer, uint)

// Bnns_graph_realloc_fn_t is the workspace and output allocation function.
//
// See: https://developer.apple.com/documentation/Accelerate/bnns_graph_realloc_fn_t
type Bnns_graph_realloc_fn_t = func(unsafe.Pointer, uint, unsafe.Pointer, uint, uint) int

// See: https://developer.apple.com/documentation/Accelerate/la_attribute_t
type La_attribute_t = uint

// See: https://developer.apple.com/documentation/Accelerate/la_count_t
type La_count_t = uint

// See: https://developer.apple.com/documentation/Accelerate/la_deallocator_t
type La_deallocator_t = func(unsafe.Pointer)

// See: https://developer.apple.com/documentation/Accelerate/la_hint_t
type La_hint_t = uint

// See: https://developer.apple.com/documentation/Accelerate/la_index_t
type La_index_t = int

// See: https://developer.apple.com/documentation/Accelerate/la_norm_t
type La_norm_t = uint

// See: https://developer.apple.com/documentation/Accelerate/la_object_t
type La_object_t = objectivec.Object

// See: https://developer.apple.com/documentation/Accelerate/la_scalar_type_t
type La_scalar_type_t = uint

// See: https://developer.apple.com/documentation/Accelerate/la_status_t
type La_status_t = int

// See: https://developer.apple.com/documentation/Accelerate/quadrature_function_array
type Quadrature_function_array = func(unsafe.Pointer, uint, []float64, []float64)

// Simd_bool is a Boolean scalar value.
//
// See: https://developer.apple.com/documentation/simd/simd_bool
type Simd_bool = bool

// Simd_char1 is a vector of one 8-bit signed integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_char1
type Simd_char1 = int8

// Simd_char16 is a vector of sixteen 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char16
type Simd_char16 = int8

// Simd_char2 is a vector of two 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char2
type Simd_char2 = int8

// Simd_char3 is a vector of three 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char3
type Simd_char3 = int8

// Simd_char32 is a vector of thirty-two 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char32
type Simd_char32 = int8

// Simd_char4 is a vector of four 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char4
type Simd_char4 = int8

// Simd_char64 is a vector of sixty-four 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char64
type Simd_char64 = int8

// Simd_char8 is a vector of eight 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_char8
type Simd_char8 = int8

// Simd_double1 is a vector of one 64-bit floating-point element.
//
// See: https://developer.apple.com/documentation/simd/simd_double1
type Simd_double1 = float64

// Simd_double2 is a vector of two 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_double2
type Simd_double2 = float64

// Simd_double3 is a vector of three 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_double3
type Simd_double3 = float64

// Simd_double4 is a vector of four 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_double4
type Simd_double4 = float64

// Simd_double8 is a vector of eight 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_double8
type Simd_double8 = float64

// Simd_float1 is a vector of one 32-bit floating-point element.
//
// See: https://developer.apple.com/documentation/simd/simd_float1
type Simd_float1 = float32

// Simd_float16 is a vector of sixteen 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_float16
type Simd_float16 = float32

// Simd_float2 is a vector of two 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_float2
type Simd_float2 = float32

// Simd_float3 is a vector of three 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_float3
type Simd_float3 = float32

// Simd_float4 is a vector of four 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_float4
type Simd_float4 = float32

// Simd_float8 is a vector of eight 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_float8
type Simd_float8 = float32

// Simd_half1 is a vector of one 16-bit floating-point element.
//
// See: https://developer.apple.com/documentation/simd/simd_half1
type Simd_half1 = unsafe.Pointer

// Simd_half16 is a vector of sixteen 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half16
type Simd_half16 = unsafe.Pointer

// Simd_half2 is a vector of two 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half2
type Simd_half2 = unsafe.Pointer

// Simd_half3 is a vector of three 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half3
type Simd_half3 = unsafe.Pointer

// Simd_half32 is a vector of thirty-two 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half32
type Simd_half32 = unsafe.Pointer

// Simd_half4 is a vector of four 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half4
type Simd_half4 = unsafe.Pointer

// Simd_half8 is a vector of eight 16-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_half8
type Simd_half8 = unsafe.Pointer

// Simd_int1 is a vector of one 32-bit signed integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_int1
type Simd_int1 = int

// Simd_int16 is a vector of sixteen 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_int16
type Simd_int16 = int

// Simd_int2 is a vector of two 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_int2
type Simd_int2 = int

// Simd_int3 is a vector of three 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_int3
type Simd_int3 = int

// Simd_int4 is a vector of four 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_int4
type Simd_int4 = int

// Simd_int8 is a vector of eight 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_int8
type Simd_int8 = int

// Simd_long1 is a vector of one 64-bit signed integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_long1
type Simd_long1 = int

// Simd_long2 is a vector of two 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_long2
type Simd_long2 = int

// Simd_long3 is a vector of three 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_long3
type Simd_long3 = int

// Simd_long4 is a vector of four 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_long4
type Simd_long4 = int

// Simd_long8 is a vector of eight 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_long8
type Simd_long8 = int

// Simd_packed_char16 is a packed vector of sixteen 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char16
type Simd_packed_char16 = int8

// Simd_packed_char2 is a packed vector of two 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char2
type Simd_packed_char2 = int8

// Simd_packed_char32 is a packed vector of thirty-two 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char32
type Simd_packed_char32 = int8

// Simd_packed_char4 is a packed vector of four 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char4
type Simd_packed_char4 = int8

// Simd_packed_char64 is a packed vector of sixty-four 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char64
type Simd_packed_char64 = int8

// Simd_packed_char8 is a packed vector of eight 8-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_char8
type Simd_packed_char8 = int8

// Simd_packed_double2 is a packed vector of two 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_double2
type Simd_packed_double2 = float64

// Simd_packed_double4 is a packed vector of four 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_double4
type Simd_packed_double4 = float64

// Simd_packed_double8 is a packed vector of eight 64-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_double8
type Simd_packed_double8 = float64

// Simd_packed_float16 is a packed vector of sixteen 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_float16
type Simd_packed_float16 = float32

// Simd_packed_float2 is a packed vector of two 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_float2
type Simd_packed_float2 = float32

// Simd_packed_float4 is a packed vector of four 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_float4
type Simd_packed_float4 = float32

// Simd_packed_float8 is a packed vector of eight 32-bit floating-point elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_float8
type Simd_packed_float8 = float32

// Simd_packed_int16 is a packed vector of sixteen 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_int16
type Simd_packed_int16 = int

// Simd_packed_int2 is a packed vector of two 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_int2
type Simd_packed_int2 = int

// Simd_packed_int4 is a packed vector of four 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_int4
type Simd_packed_int4 = int

// Simd_packed_int8 is a packed vector of eight 32-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_int8
type Simd_packed_int8 = int

// Simd_packed_long2 is a packed vector of two 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_long2
type Simd_packed_long2 = int

// Simd_packed_long4 is a packed vector of four 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_long4
type Simd_packed_long4 = int

// Simd_packed_long8 is a packed vector of eight 64-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_long8
type Simd_packed_long8 = int

// Simd_packed_short16 is a packed vector of sixteen 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_short16
type Simd_packed_short16 = int16

// Simd_packed_short2 is a packed vector of two 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_short2
type Simd_packed_short2 = int16

// Simd_packed_short32 is a packed vector of thirty-two 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_short32
type Simd_packed_short32 = int16

// Simd_packed_short4 is a packed vector of four 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_short4
type Simd_packed_short4 = int16

// Simd_packed_short8 is a packed vector of eight 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_short8
type Simd_packed_short8 = int16

// Simd_packed_uchar16 is a packed vector of sixteen 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar16
type Simd_packed_uchar16 = uint8

// Simd_packed_uchar2 is a packed vector of two 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar2
type Simd_packed_uchar2 = uint8

// Simd_packed_uchar32 is a packed vector of thirty-two 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar32
type Simd_packed_uchar32 = uint8

// Simd_packed_uchar4 is a packed vector of four 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar4
type Simd_packed_uchar4 = uint8

// Simd_packed_uchar64 is a packed vector of sixty-four 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar64
type Simd_packed_uchar64 = uint8

// Simd_packed_uchar8 is a packed vector of eight 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uchar8
type Simd_packed_uchar8 = uint8

// Simd_packed_uint16 is a packed vector of sixteen 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uint16
type Simd_packed_uint16 = uint

// Simd_packed_uint2 is a packed vector of two 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uint2
type Simd_packed_uint2 = uint

// Simd_packed_uint4 is a packed vector of four 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uint4
type Simd_packed_uint4 = uint

// Simd_packed_uint8 is a packed vector of eight 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_uint8
type Simd_packed_uint8 = uint

// Simd_packed_ulong2 is a packed vector of two 64-bit unsigned integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ulong2
type Simd_packed_ulong2 = uint

// Simd_packed_ulong4 is a packed vector of four 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ulong4
type Simd_packed_ulong4 = uint

// Simd_packed_ulong8 is a packed vector of eight 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ulong8
type Simd_packed_ulong8 = uint

// Simd_packed_ushort16 is a packed vector of sixteen 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ushort16
type Simd_packed_ushort16 = uint16

// Simd_packed_ushort2 is a packed vector of two 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ushort2
type Simd_packed_ushort2 = uint16

// Simd_packed_ushort32 is a packed vector of thirty-two 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ushort32
type Simd_packed_ushort32 = uint16

// Simd_packed_ushort4 is a packed vector of four 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ushort4
type Simd_packed_ushort4 = uint16

// Simd_packed_ushort8 is a packed vector of eight 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_packed_ushort8
type Simd_packed_ushort8 = uint16

// Simd_short1 is a vector of one 16-bit signed integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_short1
type Simd_short1 = int16

// Simd_short16 is a vector of sixteen 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short16
type Simd_short16 = int16

// Simd_short2 is a vector of two 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short2
type Simd_short2 = int16

// Simd_short3 is a vector of three 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short3
type Simd_short3 = int16

// Simd_short32 is a vector of thirty-two 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short32
type Simd_short32 = int16

// Simd_short4 is a vector of four 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short4
type Simd_short4 = int16

// Simd_short8 is a vector of eight 16-bit signed integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_short8
type Simd_short8 = int16

// Simd_uchar1 is a vector of one 8-bit unsigned integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar1
type Simd_uchar1 = uint8

// Simd_uchar16 is a vector of sixteen 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar16
type Simd_uchar16 = uint8

// Simd_uchar2 is a vector of two 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar2
type Simd_uchar2 = uint8

// Simd_uchar3 is a vector of three 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar3
type Simd_uchar3 = uint8

// Simd_uchar32 is a vector of thirty-two 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar32
type Simd_uchar32 = uint8

// Simd_uchar4 is a vector of four 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar4
type Simd_uchar4 = uint8

// Simd_uchar64 is a vector of sixty-four 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar64
type Simd_uchar64 = uint8

// Simd_uchar8 is a vector of eight 8-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uchar8
type Simd_uchar8 = uint8

// Simd_uint1 is a vector of one 32-bit unsigned integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_uint1
type Simd_uint1 = uint

// Simd_uint16 is a vector of sixteen 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uint16
type Simd_uint16 = uint

// Simd_uint2 is a vector of two 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uint2
type Simd_uint2 = uint

// Simd_uint3 is a vector of three 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uint3
type Simd_uint3 = uint

// Simd_uint4 is a vector of four 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uint4
type Simd_uint4 = uint

// Simd_uint8 is a vector of eight 32-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_uint8
type Simd_uint8 = uint

// Simd_ulong1 is a vector of one 64-bit unsigned integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_ulong1
type Simd_ulong1 = uint

// Simd_ulong2 is a vector of two 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ulong2
type Simd_ulong2 = uint

// Simd_ulong3 is a vector of three 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ulong3
type Simd_ulong3 = uint

// Simd_ulong4 is a vector of four 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ulong4
type Simd_ulong4 = uint

// Simd_ulong8 is a vector of eight 64-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ulong8
type Simd_ulong8 = uint

// Simd_ushort1 is a vector of one 16-bit unsigned integer element.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort1
type Simd_ushort1 = uint16

// Simd_ushort16 is a vector of sixteen 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort16
type Simd_ushort16 = uint16

// Simd_ushort2 is a vector of two 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort2
type Simd_ushort2 = uint16

// Simd_ushort3 is a vector of three 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort3
type Simd_ushort3 = uint16

// Simd_ushort32 is a vector of thirty-two 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort32
type Simd_ushort32 = uint16

// Simd_ushort4 is a vector of four 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort4
type Simd_ushort4 = uint16

// Simd_ushort8 is a vector of eight 16-bit unsigned integer elements.
//
// See: https://developer.apple.com/documentation/simd/simd_ushort8
type Simd_ushort8 = uint16

// Sparse_dimension is the dimension type.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_dimension
type Sparse_dimension = uint64

// Sparse_index is the index type.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_index
type Sparse_index = int64

// Sparse_matrix_double is sparse matrix opaque type for double.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_double
type Sparse_matrix_double = uintptr

// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_double_complex
type Sparse_matrix_double_complex = uintptr

// Sparse_matrix_float is sparse matrix opaque type for float.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_float
type Sparse_matrix_float = uintptr

// See: https://developer.apple.com/documentation/Accelerate/sparse_matrix_float_complex
type Sparse_matrix_float_complex = uintptr

// Sparse_stride is the stride type.
//
// See: https://developer.apple.com/documentation/Accelerate/sparse_stride
type Sparse_stride = int64

// VBool32 is a 128-bit vector packed with `bool int` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vBool32
type VBool32 = uint

// VDSP_DFT_Interleaved_Setup is an opaque type that contains setup information for an interleaved single-precision discrete Fourier transform (DFT).
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_Setup
type VDSP_DFT_Interleaved_Setup = uintptr

// VDSP_DFT_Interleaved_SetupD is an opaque type that contains setup information for an interleaved double-precision discrete Fourier transform (DFT).
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_SetupD
type VDSP_DFT_Interleaved_SetupD = uintptr

// VDSP_DFT_Setup is an opaque type that contains setup information for a single-precision discrete Fourier transform (DFT).
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Setup
type VDSP_DFT_Setup = uintptr

// VDSP_DFT_SetupD is an opaque type that contains setup information for a double-precision discrete Fourier transform (DFT).
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_SetupD
type VDSP_DFT_SetupD = uintptr

// VDSP_Length is an unsigned-integer value that represents the size of vectors and the indices of elements in vectors.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_Length
type VDSP_Length = uint

// VDSP_Stride is an integer value that represents the differences between indices of elements, including the lengths of strides.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_Stride
type VDSP_Stride = int

// VDSP_biquad_Setup is a data structure that contains precalculated data for use by the single-precision cascaded biquadratic IIR filter function.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_Setup
type VDSP_biquad_Setup = uintptr

// VDSP_biquad_SetupD is a data structure that contains precalculated data for use by the double-precision cascaded biquadratic IIR filter function.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetupD
type VDSP_biquad_SetupD = uintptr

// VDSP_biquadm_Setup is a data structure that contains precalculated data for use by a single-precision, multichannel cascaded biquadratic filter function.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_biquadm_Setup
type VDSP_biquadm_Setup = uintptr

// VDSP_biquadm_SetupD is a data structure that contains precalculated data for use by a double-precision, multichannel cascaded biquadratic filter function.
//
// See: https://developer.apple.com/documentation/Accelerate/vDSP_biquadm_SetupD
type VDSP_biquadm_SetupD = uintptr

// VDouble is a 128-bit vector packed with `double` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vDouble
type VDouble = float64

// VFloat is a 128-bit vector packed with `float` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vFloat
type VFloat = float32

// See: https://developer.apple.com/documentation/Accelerate/vFloatPacked
type VFloatPacked = float32

// VImageBufferTypeCode is type codes, such as chrominance or luminance, for the contents of a vImage buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageBufferTypeCode
type VImageBufferTypeCode = uint32

// VImageCVImageFormatError is additional error codes for functions that use the vImageCVImageFormatRef.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormatError
type VImageCVImageFormatError = int

// VImageCVImageFormatRef is a mutable description of image encoding in a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat
type VImageCVImageFormatRef uintptr

// VImageConstCVImageFormatRef is an immutable description of image encoding in a Core Video pixel buffer.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConstCVImageFormat
type VImageConstCVImageFormatRef = unsafe.Pointer

// VImageConverterRef is a description of a conversion from one image format to another.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageConverter
type VImageConverterRef = unsafe.Pointer

// VImageMatrixType is an enumeration of RGB -> Y’CbCr conversion matrix types.
//
// See: https://developer.apple.com/documentation/Accelerate/vImageMatrixType
type VImageMatrixType = uint32

// VImagePixelCount is a type for the number of pixels.
//
// See: https://developer.apple.com/documentation/Accelerate/vImagePixelCount
type VImagePixelCount = uint

// VImage_CGAffineTransform is a structure for values that represent a Core Graphics–compatible affine transformation.
//
// See: https://developer.apple.com/documentation/Accelerate/vImage_CGAffineTransform
type VImage_CGAffineTransform = VImage_AffineTransform_Double

// VImage_Error is a type for image errors.
//
// See: https://developer.apple.com/documentation/Accelerate/vImage_Error
type VImage_Error = int

// VImage_Flags is a type for processing options.
//
// See: https://developer.apple.com/documentation/Accelerate/vImage_Flags
type VImage_Flags = uint32

// VImage_MultidimensionalTable is an opaque pointer that represents a multidimensional lookup table.
//
// See: https://developer.apple.com/documentation/Accelerate/vImage_MultidimensionalTable
type VImage_MultidimensionalTable = uintptr

// VImage_WarpInterpolation is constants for selecting the interpolation mode.
//
// See: https://developer.apple.com/documentation/Accelerate/vImage_WarpInterpolation
type VImage_WarpInterpolation = int32

// VSInt16 is a 128-bit vector packed with `signed short` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vSInt16
type VSInt16 = int16

// VSInt32 is a 128-bit vector packed with `signed int` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vSInt32
type VSInt32 = int

// VSInt64 is a 128-bit vector packed with `int64_t` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vSInt64
type VSInt64 = int64

// VSInt8 is a 128-bit vector packed with `signed char` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vSInt8
type VSInt8 = int8

// VUInt16 is a 128-bit vector packed with `unsigned short` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vUInt16
type VUInt16 = uint16

// VUInt32 is a 128-bit vector packed with `unsigned int` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vUInt32
type VUInt32 = uint

// VUInt64 is a 128-bit vector packed with `uint64_t` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vUInt64
type VUInt64 = uint64

// VUInt8 is a 128-bit vector packed with `unsigned char` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vUInt8
type VUInt8 = uint8

// ComplexSplit is a Go-name alias for COMPLEX_SPLIT.
type ComplexSplit = COMPLEX_SPLIT

// DoubleComplex is a Go-name alias for DOUBLE_COMPLEX.
type DoubleComplex = DOUBLE_COMPLEX

// DoubleComplexSplit is a Go-name alias for DOUBLE_COMPLEX_SPLIT.
type DoubleComplexSplit = DOUBLE_COMPLEX_SPLIT

// Pixel16f is a Go-name alias for Pixel_16F.
type Pixel16f = Pixel_16F

// Pixel16f16f is a Go-name alias for Pixel_16F16F.
type Pixel16f16f = Pixel_16F16F

// Pixel16q12 is a Go-name alias for Pixel_16Q12.
type Pixel16q12 = Pixel_16Q12

// Pixel16s is a Go-name alias for Pixel_16S.
type Pixel16s = Pixel_16S

// Pixel16s16s is a Go-name alias for Pixel_16S16S.
type Pixel16s16s = Pixel_16S16S

// Pixel16u is a Go-name alias for Pixel_16U.
type Pixel16u = Pixel_16U

// Pixel16u16u is a Go-name alias for Pixel_16U16U.
type Pixel16u16u = Pixel_16U16U

// Pixel32u is a Go-name alias for Pixel_32U.
type Pixel32u = Pixel_32U

// Pixel8 is a Go-name alias for Pixel_8.
type Pixel8 = Pixel_8

// Pixel88 is a Go-name alias for Pixel_88.
type Pixel88 = Pixel_88

// Pixel8888 is a Go-name alias for Pixel_8888.
type Pixel8888 = Pixel_8888

// PixelArgb16f is a Go-name alias for Pixel_ARGB_16F.
type PixelArgb16f = Pixel_ARGB_16F

// PixelArgb16s is a Go-name alias for Pixel_ARGB_16S.
type PixelArgb16s = Pixel_ARGB_16S

// PixelArgb16u is a Go-name alias for Pixel_ARGB_16U.
type PixelArgb16u = Pixel_ARGB_16U

// PixelF is a Go-name alias for Pixel_F.
type PixelF = Pixel_F

// PixelFf is a Go-name alias for Pixel_FF.
type PixelFf = Pixel_FF

// PixelFfff is a Go-name alias for Pixel_FFFF.
type PixelFfff = Pixel_FFFF

// BnnsGraphCompileMessageFn is a Go-name alias for Bnns_graph_compile_message_fn_t.
type BnnsGraphCompileMessageFn = Bnns_graph_compile_message_fn_t

// BnnsGraphExecuteMessageFn is a Go-name alias for Bnns_graph_execute_message_fn_t.
type BnnsGraphExecuteMessageFn = Bnns_graph_execute_message_fn_t

// BnnsGraphFreeAllFn is a Go-name alias for Bnns_graph_free_all_fn_t.
type BnnsGraphFreeAllFn = Bnns_graph_free_all_fn_t

// BnnsGraphReallocFn is a Go-name alias for Bnns_graph_realloc_fn_t.
type BnnsGraphReallocFn = Bnns_graph_realloc_fn_t

// LaAttribute is a Go-name alias for La_attribute_t.
type LaAttribute = La_attribute_t

// LaCount is a Go-name alias for La_count_t.
type LaCount = La_count_t

// LaDeallocator is a Go-name alias for La_deallocator_t.
type LaDeallocator = La_deallocator_t

// LaHint is a Go-name alias for La_hint_t.
type LaHint = La_hint_t

// LaIndex is a Go-name alias for La_index_t.
type LaIndex = La_index_t

// LaNorm is a Go-name alias for La_norm_t.
type LaNorm = La_norm_t

// LaObject is a Go-name alias for La_object_t.
type LaObject = La_object_t

// LaScalarType is a Go-name alias for La_scalar_type_t.
type LaScalarType = La_scalar_type_t

// LaStatus is a Go-name alias for La_status_t.
type LaStatus = La_status_t

// QuadratureFunctionArray is a Go-name alias for Quadrature_function_array.
type QuadratureFunctionArray = Quadrature_function_array

// SIMDBool is a Go-name alias for Simd_bool.
type SIMDBool = Simd_bool

// SIMDChar1 is a Go-name alias for Simd_char1.
type SIMDChar1 = Simd_char1

// SIMDChar16 is a Go-name alias for Simd_char16.
type SIMDChar16 = Simd_char16

// SIMDChar2 is a Go-name alias for Simd_char2.
type SIMDChar2 = Simd_char2

// SIMDChar3 is a Go-name alias for Simd_char3.
type SIMDChar3 = Simd_char3

// SIMDChar32 is a Go-name alias for Simd_char32.
type SIMDChar32 = Simd_char32

// SIMDChar4 is a Go-name alias for Simd_char4.
type SIMDChar4 = Simd_char4

// SIMDChar64 is a Go-name alias for Simd_char64.
type SIMDChar64 = Simd_char64

// SIMDChar8 is a Go-name alias for Simd_char8.
type SIMDChar8 = Simd_char8

// SIMDDouble1 is a Go-name alias for Simd_double1.
type SIMDDouble1 = Simd_double1

// SIMDDouble2 is a Go-name alias for Simd_double2.
type SIMDDouble2 = Simd_double2

// SIMDDouble3 is a Go-name alias for Simd_double3.
type SIMDDouble3 = Simd_double3

// SIMDDouble4 is a Go-name alias for Simd_double4.
type SIMDDouble4 = Simd_double4

// SIMDDouble8 is a Go-name alias for Simd_double8.
type SIMDDouble8 = Simd_double8

// SIMDFloat1 is a Go-name alias for Simd_float1.
type SIMDFloat1 = Simd_float1

// SIMDFloat16 is a Go-name alias for Simd_float16.
type SIMDFloat16 = Simd_float16

// SIMDFloat2 is a Go-name alias for Simd_float2.
type SIMDFloat2 = Simd_float2

// SIMDFloat3 is a Go-name alias for Simd_float3.
type SIMDFloat3 = Simd_float3

// SIMDFloat4 is a Go-name alias for Simd_float4.
type SIMDFloat4 = Simd_float4

// SIMDFloat8 is a Go-name alias for Simd_float8.
type SIMDFloat8 = Simd_float8

// SIMDHalf1 is a Go-name alias for Simd_half1.
type SIMDHalf1 = Simd_half1

// SIMDHalf16 is a Go-name alias for Simd_half16.
type SIMDHalf16 = Simd_half16

// SIMDHalf2 is a Go-name alias for Simd_half2.
type SIMDHalf2 = Simd_half2

// SIMDHalf3 is a Go-name alias for Simd_half3.
type SIMDHalf3 = Simd_half3

// SIMDHalf32 is a Go-name alias for Simd_half32.
type SIMDHalf32 = Simd_half32

// SIMDHalf4 is a Go-name alias for Simd_half4.
type SIMDHalf4 = Simd_half4

// SIMDHalf8 is a Go-name alias for Simd_half8.
type SIMDHalf8 = Simd_half8

// SIMDInt1 is a Go-name alias for Simd_int1.
type SIMDInt1 = Simd_int1

// SIMDInt16 is a Go-name alias for Simd_int16.
type SIMDInt16 = Simd_int16

// SIMDInt2 is a Go-name alias for Simd_int2.
type SIMDInt2 = Simd_int2

// SIMDInt3 is a Go-name alias for Simd_int3.
type SIMDInt3 = Simd_int3

// SIMDInt4 is a Go-name alias for Simd_int4.
type SIMDInt4 = Simd_int4

// SIMDInt8 is a Go-name alias for Simd_int8.
type SIMDInt8 = Simd_int8

// SIMDLong1 is a Go-name alias for Simd_long1.
type SIMDLong1 = Simd_long1

// SIMDLong2 is a Go-name alias for Simd_long2.
type SIMDLong2 = Simd_long2

// SIMDLong3 is a Go-name alias for Simd_long3.
type SIMDLong3 = Simd_long3

// SIMDLong4 is a Go-name alias for Simd_long4.
type SIMDLong4 = Simd_long4

// SIMDLong8 is a Go-name alias for Simd_long8.
type SIMDLong8 = Simd_long8

// SIMDPackedChar16 is a Go-name alias for Simd_packed_char16.
type SIMDPackedChar16 = Simd_packed_char16

// SIMDPackedChar2 is a Go-name alias for Simd_packed_char2.
type SIMDPackedChar2 = Simd_packed_char2

// SIMDPackedChar32 is a Go-name alias for Simd_packed_char32.
type SIMDPackedChar32 = Simd_packed_char32

// SIMDPackedChar4 is a Go-name alias for Simd_packed_char4.
type SIMDPackedChar4 = Simd_packed_char4

// SIMDPackedChar64 is a Go-name alias for Simd_packed_char64.
type SIMDPackedChar64 = Simd_packed_char64

// SIMDPackedChar8 is a Go-name alias for Simd_packed_char8.
type SIMDPackedChar8 = Simd_packed_char8

// SIMDPackedDouble2 is a Go-name alias for Simd_packed_double2.
type SIMDPackedDouble2 = Simd_packed_double2

// SIMDPackedDouble4 is a Go-name alias for Simd_packed_double4.
type SIMDPackedDouble4 = Simd_packed_double4

// SIMDPackedDouble8 is a Go-name alias for Simd_packed_double8.
type SIMDPackedDouble8 = Simd_packed_double8

// SIMDPackedFloat16 is a Go-name alias for Simd_packed_float16.
type SIMDPackedFloat16 = Simd_packed_float16

// SIMDPackedFloat2 is a Go-name alias for Simd_packed_float2.
type SIMDPackedFloat2 = Simd_packed_float2

// SIMDPackedFloat4 is a Go-name alias for Simd_packed_float4.
type SIMDPackedFloat4 = Simd_packed_float4

// SIMDPackedFloat8 is a Go-name alias for Simd_packed_float8.
type SIMDPackedFloat8 = Simd_packed_float8

// SIMDPackedInt16 is a Go-name alias for Simd_packed_int16.
type SIMDPackedInt16 = Simd_packed_int16

// SIMDPackedInt2 is a Go-name alias for Simd_packed_int2.
type SIMDPackedInt2 = Simd_packed_int2

// SIMDPackedInt4 is a Go-name alias for Simd_packed_int4.
type SIMDPackedInt4 = Simd_packed_int4

// SIMDPackedInt8 is a Go-name alias for Simd_packed_int8.
type SIMDPackedInt8 = Simd_packed_int8

// SIMDPackedLong2 is a Go-name alias for Simd_packed_long2.
type SIMDPackedLong2 = Simd_packed_long2

// SIMDPackedLong4 is a Go-name alias for Simd_packed_long4.
type SIMDPackedLong4 = Simd_packed_long4

// SIMDPackedLong8 is a Go-name alias for Simd_packed_long8.
type SIMDPackedLong8 = Simd_packed_long8

// SIMDPackedShort16 is a Go-name alias for Simd_packed_short16.
type SIMDPackedShort16 = Simd_packed_short16

// SIMDPackedShort2 is a Go-name alias for Simd_packed_short2.
type SIMDPackedShort2 = Simd_packed_short2

// SIMDPackedShort32 is a Go-name alias for Simd_packed_short32.
type SIMDPackedShort32 = Simd_packed_short32

// SIMDPackedShort4 is a Go-name alias for Simd_packed_short4.
type SIMDPackedShort4 = Simd_packed_short4

// SIMDPackedShort8 is a Go-name alias for Simd_packed_short8.
type SIMDPackedShort8 = Simd_packed_short8

// SIMDPackedUchar16 is a Go-name alias for Simd_packed_uchar16.
type SIMDPackedUchar16 = Simd_packed_uchar16

// SIMDPackedUchar2 is a Go-name alias for Simd_packed_uchar2.
type SIMDPackedUchar2 = Simd_packed_uchar2

// SIMDPackedUchar32 is a Go-name alias for Simd_packed_uchar32.
type SIMDPackedUchar32 = Simd_packed_uchar32

// SIMDPackedUchar4 is a Go-name alias for Simd_packed_uchar4.
type SIMDPackedUchar4 = Simd_packed_uchar4

// SIMDPackedUchar64 is a Go-name alias for Simd_packed_uchar64.
type SIMDPackedUchar64 = Simd_packed_uchar64

// SIMDPackedUchar8 is a Go-name alias for Simd_packed_uchar8.
type SIMDPackedUchar8 = Simd_packed_uchar8

// SIMDPackedUint16 is a Go-name alias for Simd_packed_uint16.
type SIMDPackedUint16 = Simd_packed_uint16

// SIMDPackedUint2 is a Go-name alias for Simd_packed_uint2.
type SIMDPackedUint2 = Simd_packed_uint2

// SIMDPackedUint4 is a Go-name alias for Simd_packed_uint4.
type SIMDPackedUint4 = Simd_packed_uint4

// SIMDPackedUint8 is a Go-name alias for Simd_packed_uint8.
type SIMDPackedUint8 = Simd_packed_uint8

// SIMDPackedUlong2 is a Go-name alias for Simd_packed_ulong2.
type SIMDPackedUlong2 = Simd_packed_ulong2

// SIMDPackedUlong4 is a Go-name alias for Simd_packed_ulong4.
type SIMDPackedUlong4 = Simd_packed_ulong4

// SIMDPackedUlong8 is a Go-name alias for Simd_packed_ulong8.
type SIMDPackedUlong8 = Simd_packed_ulong8

// SIMDPackedUshort16 is a Go-name alias for Simd_packed_ushort16.
type SIMDPackedUshort16 = Simd_packed_ushort16

// SIMDPackedUshort2 is a Go-name alias for Simd_packed_ushort2.
type SIMDPackedUshort2 = Simd_packed_ushort2

// SIMDPackedUshort32 is a Go-name alias for Simd_packed_ushort32.
type SIMDPackedUshort32 = Simd_packed_ushort32

// SIMDPackedUshort4 is a Go-name alias for Simd_packed_ushort4.
type SIMDPackedUshort4 = Simd_packed_ushort4

// SIMDPackedUshort8 is a Go-name alias for Simd_packed_ushort8.
type SIMDPackedUshort8 = Simd_packed_ushort8

// SIMDShort1 is a Go-name alias for Simd_short1.
type SIMDShort1 = Simd_short1

// SIMDShort16 is a Go-name alias for Simd_short16.
type SIMDShort16 = Simd_short16

// SIMDShort2 is a Go-name alias for Simd_short2.
type SIMDShort2 = Simd_short2

// SIMDShort3 is a Go-name alias for Simd_short3.
type SIMDShort3 = Simd_short3

// SIMDShort32 is a Go-name alias for Simd_short32.
type SIMDShort32 = Simd_short32

// SIMDShort4 is a Go-name alias for Simd_short4.
type SIMDShort4 = Simd_short4

// SIMDShort8 is a Go-name alias for Simd_short8.
type SIMDShort8 = Simd_short8

// SIMDUchar1 is a Go-name alias for Simd_uchar1.
type SIMDUchar1 = Simd_uchar1

// SIMDUchar16 is a Go-name alias for Simd_uchar16.
type SIMDUchar16 = Simd_uchar16

// SIMDUchar2 is a Go-name alias for Simd_uchar2.
type SIMDUchar2 = Simd_uchar2

// SIMDUchar3 is a Go-name alias for Simd_uchar3.
type SIMDUchar3 = Simd_uchar3

// SIMDUchar32 is a Go-name alias for Simd_uchar32.
type SIMDUchar32 = Simd_uchar32

// SIMDUchar4 is a Go-name alias for Simd_uchar4.
type SIMDUchar4 = Simd_uchar4

// SIMDUchar64 is a Go-name alias for Simd_uchar64.
type SIMDUchar64 = Simd_uchar64

// SIMDUchar8 is a Go-name alias for Simd_uchar8.
type SIMDUchar8 = Simd_uchar8

// SIMDUint1 is a Go-name alias for Simd_uint1.
type SIMDUint1 = Simd_uint1

// SIMDUint16 is a Go-name alias for Simd_uint16.
type SIMDUint16 = Simd_uint16

// SIMDUint2 is a Go-name alias for Simd_uint2.
type SIMDUint2 = Simd_uint2

// SIMDUint3 is a Go-name alias for Simd_uint3.
type SIMDUint3 = Simd_uint3

// SIMDUint4 is a Go-name alias for Simd_uint4.
type SIMDUint4 = Simd_uint4

// SIMDUint8 is a Go-name alias for Simd_uint8.
type SIMDUint8 = Simd_uint8

// SIMDUlong1 is a Go-name alias for Simd_ulong1.
type SIMDUlong1 = Simd_ulong1

// SIMDUlong2 is a Go-name alias for Simd_ulong2.
type SIMDUlong2 = Simd_ulong2

// SIMDUlong3 is a Go-name alias for Simd_ulong3.
type SIMDUlong3 = Simd_ulong3

// SIMDUlong4 is a Go-name alias for Simd_ulong4.
type SIMDUlong4 = Simd_ulong4

// SIMDUlong8 is a Go-name alias for Simd_ulong8.
type SIMDUlong8 = Simd_ulong8

// SIMDUshort1 is a Go-name alias for Simd_ushort1.
type SIMDUshort1 = Simd_ushort1

// SIMDUshort16 is a Go-name alias for Simd_ushort16.
type SIMDUshort16 = Simd_ushort16

// SIMDUshort2 is a Go-name alias for Simd_ushort2.
type SIMDUshort2 = Simd_ushort2

// SIMDUshort3 is a Go-name alias for Simd_ushort3.
type SIMDUshort3 = Simd_ushort3

// SIMDUshort32 is a Go-name alias for Simd_ushort32.
type SIMDUshort32 = Simd_ushort32

// SIMDUshort4 is a Go-name alias for Simd_ushort4.
type SIMDUshort4 = Simd_ushort4

// SIMDUshort8 is a Go-name alias for Simd_ushort8.
type SIMDUshort8 = Simd_ushort8

// SparseDimension is a Go-name alias for Sparse_dimension.
type SparseDimension = Sparse_dimension

// SparseIndex is a Go-name alias for Sparse_index.
type SparseIndex = Sparse_index

// SparseMatrixDouble is a Go-name alias for Sparse_matrix_double.
type SparseMatrixDouble = Sparse_matrix_double

// SparseMatrixDoubleComplex is a Go-name alias for Sparse_matrix_double_complex.
type SparseMatrixDoubleComplex = Sparse_matrix_double_complex

// SparseMatrixFloat is a Go-name alias for Sparse_matrix_float.
type SparseMatrixFloat = Sparse_matrix_float

// SparseMatrixFloatComplex is a Go-name alias for Sparse_matrix_float_complex.
type SparseMatrixFloatComplex = Sparse_matrix_float_complex

// SparseStride is a Go-name alias for Sparse_stride.
type SparseStride = Sparse_stride

// VdspDftInterleavedSetup is a Go-name alias for VDSP_DFT_Interleaved_Setup.
type VdspDftInterleavedSetup = VDSP_DFT_Interleaved_Setup

// VdspDftInterleavedSetupD is a Go-name alias for VDSP_DFT_Interleaved_SetupD.
type VdspDftInterleavedSetupD = VDSP_DFT_Interleaved_SetupD

// VdspDftSetup is a Go-name alias for VDSP_DFT_Setup.
type VdspDftSetup = VDSP_DFT_Setup

// VdspDftSetupD is a Go-name alias for VDSP_DFT_SetupD.
type VdspDftSetupD = VDSP_DFT_SetupD

// VdspLength is a Go-name alias for VDSP_Length.
type VdspLength = VDSP_Length

// VdspStride is a Go-name alias for VDSP_Stride.
type VdspStride = VDSP_Stride

// VdspBiquadSetup is a Go-name alias for VDSP_biquad_Setup.
type VdspBiquadSetup = VDSP_biquad_Setup

// VdspBiquadSetupD is a Go-name alias for VDSP_biquad_SetupD.
type VdspBiquadSetupD = VDSP_biquad_SetupD

// VdspBiquadmSetup is a Go-name alias for VDSP_biquadm_Setup.
type VdspBiquadmSetup = VDSP_biquadm_Setup

// VdspBiquadmSetupD is a Go-name alias for VDSP_biquadm_SetupD.
type VdspBiquadmSetupD = VDSP_biquadm_SetupD

// VImageCGAffineTransform is a Go-name alias for VImage_CGAffineTransform.
type VImageCGAffineTransform = VImage_CGAffineTransform

// VImageError is a Go-name alias for VImage_Error.
type VImageError = VImage_Error

// VImageFlags is a Go-name alias for VImage_Flags.
type VImageFlags = VImage_Flags

// VImageMultidimensionalTable is a Go-name alias for VImage_MultidimensionalTable.
type VImageMultidimensionalTable = VImage_MultidimensionalTable

// VImageWarpInterpolation is a Go-name alias for VImage_WarpInterpolation.
type VImageWarpInterpolation = VImage_WarpInterpolation

// BNNSArithmeticFunction aliases the generated arithmetic-function enum.
type BNNSArithmeticFunction = BNNSArithmetic

// BNNSBoxCoordinateMode aliases the generated box-coordinate enum family.
type BNNSBoxCoordinateMode = Bnns

// BNNSDescriptorType is the descriptor-kind enum used by BNNS tensor parameter structs.
type BNNSDescriptorType = Bnns

// BNNSDataLayout is the layout enum used by BNNSNDArrayDescriptor.
type BNNSDataLayout = BNNSData

// BNNSEmbeddingFlags aliases the generated embedding-flags enum.
type BNNSEmbeddingFlags = BNNSEmbeddingFlagScaleGradientBy

// BNNSFilterType aliases the generated BNNS filter-type enum family.
type BNNSFilterType = Bnns

// BNNSLinearSamplingMode aliases the generated linear-sampling enum.
type BNNSLinearSamplingMode = BNNSLinearSampling

// BNNSLossReductionFunction aliases the generated loss-reduction enum.
type BNNSLossReductionFunction = BNNSLossReduction

// BNNSNDArrayFlags is the flag enum used by BNNSNDArrayDescriptor.
type BNNSNDArrayFlags = BNNSNDArrayFlagBackprop

// BNNSNormType aliases the generated norm-type enum.
type BNNSNormType = Bnnsl2

// BNNSOptimizerClippingFunction aliases the generated optimizer-clipping enum.
type BNNSOptimizerClippingFunction = BNNSOptimizerClipping

// BNNSOptimizerRegularizationFunction aliases the generated regularization enum.
type BNNSOptimizerRegularizationFunction = BNNSOptimizerRegularization

// BNNSOptimizerSGDMomentumVariant aliases the generated SGD-momentum enum.
type BNNSOptimizerSGDMomentumVariant = BNNSSGDMomentum

// BNNSRandomGeneratorMethod aliases the generated random-generator enum.
type BNNSRandomGeneratorMethod = BNNSRandomGeneratorMethodAES

// Quadrature_integrator aliases the generated quadrature-integrator enum.
type Quadrature_integrator = QuadratureIntegrateQ

// Quadrature_status aliases the generated quadrature status enum.
type Quadrature_status = Quadrature

// SparseControl_t aliases the generated sparse-control enum.
type SparseControl_t = SparseDefault

// SparseGMRESVariant_t aliases the generated sparse-variant enum.
type SparseGMRESVariant_t = SparseVariant

// SparseKind_t aliases the generated sparse matrix-kind enum.
type SparseKind_t = Sparse

// SparseLSMRConvergenceTest_t aliases the generated sparse LSMR convergence enum.
type SparseLSMRConvergenceTest_t = SparseLSMRCT

// SparseStatus_t aliases the generated sparse status enum.
type SparseStatus_t = Sparse

// SparseTriangle_t aliases the generated sparse triangle enum.
type SparseTriangle_t = Sparse

// Sparse_matrix_property aliases the generated sparse matrix-property enum.
type Sparse_matrix_property = Sparse

// Sparse_status aliases the generated sparse status enum.
type Sparse_status = Sparse

// VImageARGBType aliases the generated vImage ARGB pixel-type enum.
type VImageARGBType = KvImageARG

// VImageMDTableUsageHint aliases the generated multidimensional-table hint enum.
type VImageMDTableUsageHint = KvImageMDTableHint

// VImage_InterpolationMethod aliases the generated vImage interpolation enum.
type VImage_InterpolationMethod = KvImage

// VImageYpCbCrType aliases the generated vImage YpCbCr pixel-type enum.
type VImageYpCbCrType = Kv
