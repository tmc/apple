// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// C struct types

// CGColorSpace
type CGColorSpace struct {
}

// CGImage
type CGImage struct {
}

// FaceLandmarkDetectorPoint
type FaceLandmarkDetectorPoint struct {
}

// CVBuffer
type CVBuffer struct {
}

// IOSurface
type IOSurface struct {
}

// Long
type Long struct {
	__data_    *byte
	__size_    uint64
	__cap_     objectivec.Object
	__is_long_ objectivec.Object
}

// Rep
type Rep struct {
	__s int16
	__l int
}

// SharedWeakCount
type SharedWeakCount struct {
}

// Short
type Short struct {
	__data_    [23]int8
	__size_    objectivec.Object
	__is_long_ objectivec.Object
}

// AbstractBlobContainer
type AbstractBlobContainer struct {
}

// Abstract_blob_container is a type alias for AbstractBlobContainer for use in objc.Send[T] calls.
type Abstract_blob_container = AbstractBlobContainer

// AbstractContext
type AbstractContext struct {
}

// Abstract_context is a type alias for AbstractContext for use in objc.Send[T] calls.
type Abstract_context = AbstractContext

// BlobCPU
type BlobCPU struct {
}

// Blob_cpu is a type alias for BlobCPU for use in objc.Send[T] calls.
type Blob_cpu = BlobCPU

// ConvolutionUniforms
type ConvolutionUniforms struct {
	Field1  int32
	Field2  int32
	Field3  int32
	Field4  int32
	Field5  int32
	Field6  int32
	Field7  int32
	Field8  int32
	Field9  int32
	Field10 float32
	Field11 int32
	Field12 int32
	Field13 int32
	Field14 int32
	Field15 int32
	Field16 int32
	Field17 float32
	Field18 float32
	Field19 PaddingParams
	Field20 int16
	Field21 int16
	Field22 int16
	Field23 int16
	Field24 uint16
	Field25 uint16
	Field26 int32
	Field27 int32
	Field28 int32
	Field29 int32
	Field30 int32
	Field31 uint16
	Field32 uint16
	Field33 uint16
	Field34 uint16
	Field35 uint16
	Field36 int16
	Field37 int32
	Field38 int32
	Field39 int32
	Field40 int32
	Field41 int32
	Field42 int16
	Field43 int32
	Field44 bool
}

// Convolution_uniforms is a type alias for ConvolutionUniforms for use in objc.Send[T] calls.
type Convolution_uniforms = ConvolutionUniforms

// FastPyramidResizer
type FastPyramidResizer struct {
}

// Fast_pyramid_resizer is a type alias for FastPyramidResizer for use in objc.Send[T] calls.
type Fast_pyramid_resizer = FastPyramidResizer

// FloatBuffer
type FloatBuffer struct {
	Field1 *float32
	Field2 uint64
	Field3 bool
}

// Float_buffer_t is a type alias for FloatBuffer for use in objc.Send[T] calls.
type Float_buffer_t = FloatBuffer

// GenericLoadConstantKernel
type GenericLoadConstantKernel struct {
}

// Generic_load_constant_kernel is a type alias for GenericLoadConstantKernel for use in objc.Send[T] calls.
type Generic_load_constant_kernel = GenericLoadConstantKernel

// GradientBuilder
type GradientBuilder struct {
}

// Gradient_builder is a type alias for GradientBuilder for use in objc.Send[T] calls.
type Gradient_builder = GradientBuilder

// InnerProductUniforms
type InnerProductUniforms struct {
	Field1  uint32
	Field2  uint32
	Field3  int32
	Field4  int32
	Field5  int32
	Field6  float32
	Field7  float32
	Field8  int32
	Field9  int32
	Field10 int32
	Field11 bool
	Field12 int32
	Field13 int32
	Field14 int32
	Field15 float32
	Field16 float32
	Field17 uint32
	Field18 uint32
	Field19 uint32
	Field20 uint32
	Field21 uint32
	Field22 uint32
	Field23 uint32
	Field24 uint32
	Field25 uint32
	Field26 int32
	Field27 int32
	Field28 int32
	Field29 int32
	Field30 int32
	Field31 int32
	Field32 int32
	Field33 int32
}

// Inner_product_uniforms is a type alias for InnerProductUniforms for use in objc.Send[T] calls.
type Inner_product_uniforms = InnerProductUniforms

// MxnetToolsImageHeaderT
type MxnetToolsImageHeaderT struct {
	Field1 uint32
	Field2 float32
	Field3 MxnetToolsImageIDT
}

// MxnetTools_imageHeader_t_ is a type alias for MxnetToolsImageHeaderT for use in objc.Send[T] calls.
type MxnetTools_imageHeader_t_ = MxnetToolsImageHeaderT

// MxnetToolsImageIDT
type MxnetToolsImageIDT struct {
	Field1 [2]uint64
}

// MxnetTools_imageID_t_ is a type alias for MxnetToolsImageIDT for use in objc.Send[T] calls.
type MxnetTools_imageID_t_ = MxnetToolsImageIDT

// MxnetToolsRecordHeaderT
type MxnetToolsRecordHeaderT struct {
	Field1 uint32
	Field2 uint32
}

// MxnetTools_recordHeader_t_ is a type alias for MxnetToolsRecordHeaderT for use in objc.Send[T] calls.
type MxnetTools_recordHeader_t_ = MxnetToolsRecordHeaderT

// Net
type Net struct {
}

// NetStridesConfiguration
type NetStridesConfiguration struct {
	Field1 int32
	Field2 [3]uint64
	Field3 [3]uint64
	Field4 [3]uint64
}

// Net_strides_configuration is a type alias for NetStridesConfiguration for use in objc.Send[T] calls.
type Net_strides_configuration = NetStridesConfiguration

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque [56]int8
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

// PaddingParams
type PaddingParams struct {
	Field1 int16
	Field2 int16
	Field3 float32
	Field4 int16
	Field5 int16
	Field6 int16
	Field7 int16
	Field8 int16
	Field9 int16
}

// Padding_params_t is a type alias for PaddingParams for use in objc.Send[T] calls.
type Padding_params_t = PaddingParams

// PostprocessingSettings
type PostprocessingSettings struct {
	Name                               unsafe.Pointer
	Network                            unsafe.Pointer
	Do_blend                           int32
	Blend_alpha                        float32
	Grayscale_i0                       int32
	Grayscale_i1                       int32
	Temporal_alpha                     float32
	Gamma                              float32
	Saturation                         float32
	Contrast                           float32
	Brightness                         float32
	Color_transfer_mode                int32
	Width                              int32
	Height                             int32
	Width_fast                         int32
	Height_fast                        int32
	Width_capture                      int32
	Height_capture                     int32
	Width_miniature                    int32
	Height_miniature                   int32
	Width_fullscreen                   int32
	Height_fullscreen                  int32
	Width_hd                           int32
	Height_hd                          int32
	Old_frame_scale                    float32
	Noise_strength                     float32
	Dyn_noise                          int32
	Noise_speed                        float32
	Preprocessing_bias_b               float32
	Preprocessing_bias_g               float32
	Preprocessing_bias_r               float32
	Preprocessing_scale                float32
	Preprocessing_old_frame_bias_scale float32
	High_quality_scaling               bool
}

// Postprocessing_settings_t is a type alias for PostprocessingSettings for use in objc.Send[T] calls.
type Postprocessing_settings_t = PostprocessingSettings

// SurfaceAndBuffer
type SurfaceAndBuffer struct {
}

// Surface_and_buffer is a type alias for SurfaceAndBuffer for use in objc.Send[T] calls.
type Surface_and_buffer = SurfaceAndBuffer

// V9NoiseKernel
type V9NoiseKernel struct {
}

// V9_noise_kernel is a type alias for V9NoiseKernel for use in objc.Send[T] calls.
type V9_noise_kernel = V9NoiseKernel

// VImageBuffer
type VImageBuffer struct {
	Field1 unsafe.Pointer
	Field2 uint64
	Field3 uint64
	Field4 uint64
}

// VImage_Buffer is a type alias for VImageBuffer for use in objc.Send[T] calls.
type VImage_Buffer = VImageBuffer

// Vimage2espressoParam
type Vimage2espressoParam struct {
	Scale                    float32
	Center_mean              int32
	Is_image_bgr             int32
	Is_network_bgr           int32
	Bias_r                   float32
	Bias_g                   float32
	Bias_b                   float32
	Bias_a                   float32
	Metal_output_plane       int32
	Width                    uint32
	Height                   uint32
	Rowbytes                 uint32
	Rotate_deg               int32
	Use_direct_cvpixelbuffer int32
	No_alpha_premultiply     int32
}

// Vimage2espresso_param is a type alias for Vimage2espressoParam for use in objc.Send[T] calls.
type Vimage2espresso_param = Vimage2espressoParam
