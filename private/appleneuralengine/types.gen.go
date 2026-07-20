// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// ANEBufferStruct
type ANEBufferStruct struct {
}

// ANEDeviceStruct
type ANEDeviceStruct struct {
}

// ANEMemoryMappingParamsStruct
type ANEMemoryMappingParamsStruct struct {
}

// ANENotificationMessageStruct
type ANENotificationMessageStruct struct {
}

// AnalyticsData
type AnalyticsData struct {
}

// AnalyticsGroupInfo
type AnalyticsGroupInfo struct {
}

// AnalyticsLayerInfo
type AnalyticsLayerInfo struct {
}

// AnalyticsProcedureInfo
type AnalyticsProcedureInfo struct {
}

// AnalyticsTaskInfo
type AnalyticsTaskInfo struct {
}

// Attribute
type Attribute struct {
}

// AttributeStorage
type AttributeStorage struct {
}

// Autodiff
type Autodiff struct {
}

// AutodiffOpData
type AutodiffOpData struct {
}

// BaseModuleRef
type BaseModuleRef struct {
}

// BaseRuntime
type BaseRuntime struct {
}

// BuildVersionInfo
type BuildVersionInfo struct {
}

// CGColorSpace
type CGColorSpace struct {
}

// CGImage
type CGImage struct {
}

// DeviceExtendedInfo
type DeviceExtendedInfo struct {
}

// DeviceInfo
type DeviceInfo struct {
}

// FaceLandmarkDetectorPoint
type FaceLandmarkDetectorPoint struct {
}

// FuncOp
type FuncOp struct {
}

// InMemoryModuleRef
type InMemoryModuleRef struct {
}

// Location
type Location struct {
}

// LocationAttr
type LocationAttr struct {
}

// LockFileManager
type LockFileManager struct {
}

// LockGuard
type LockGuard struct {
}

// MLIRContext
type MLIRContext struct {
}

// MPSCommandBufferDescriptor
type MPSCommandBufferDescriptor struct {
}

// MPSGraphExecutableCacheValue
type MPSGraphExecutableCacheValue struct {
}

// MPSGraphExecutableSpecializedModule
type MPSGraphExecutableSpecializedModule struct {
}

// MPSGraphModuleKey
type MPSGraphModuleKey struct {
}

// MPSGraphOperatingSystemVersion
type MPSGraphOperatingSystemVersion struct {
}

// MPSGraphSpecializationCache
type MPSGraphSpecializationCache struct {
	_moduleStorage       unsafe.Pointer
	_currentCache        unsafe.Pointer
	_failedToLoadModules unsafe.Pointer
}

// MPSResourceBlobEntry
type MPSResourceBlobEntry struct {
}

// MemoryBuffer
type MemoryBuffer struct {
}

// ModuleOp
type ModuleOp struct {
	State foundation.NSOperation
}

// ModuleResourcesLoader
type ModuleResourcesLoader struct {
}

// NamedAttribute
type NamedAttribute struct {
}

// OriginalModuleRef
type OriginalModuleRef struct {
	_module unsafe.Pointer
}

// ReadDataFromFileCache
type ReadDataFromFileCache struct {
}

// ReturnOp
type ReturnOp struct {
}

// ShapeFunctionRegistry
type ShapeFunctionRegistry struct {
}

// StringMapEntryBase
type StringMapEntryBase struct {
}

// StringRef
type StringRef struct {
}

// VMData
type VMData struct {
}

// Value
type Value struct {
}

// ValueImpl
type ValueImpl struct {
}

// VirtANEModel
type VirtANEModel struct {
}

// CFArray
type CFArray struct {
}

// CFDictionary
type CFDictionary struct {
}

// CFString
type CFString struct {
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

// SFILE
type SFILE struct {
}

// SFILEX
type SFILEX struct {
}

// Sbuf
type Sbuf struct {
}

// SharedWeakCount
type SharedWeakCount struct {
}

// Short
type Short struct {
	__data_    unsafe.Pointer
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

// ConditionVariable
type ConditionVariable struct {
	__cv_ OpaquePthreadCond
}

// Condition_variable is a type alias for ConditionVariable for use in objc.Send[T] calls.
type Condition_variable = ConditionVariable

// ConvolutionUniforms
type ConvolutionUniforms struct {
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
}

// Inner_product_uniforms is a type alias for InnerProductUniforms for use in objc.Send[T] calls.
type Inner_product_uniforms = InnerProductUniforms

// MxnetToolsImageHeaderT
type MxnetToolsImageHeaderT struct {
}

// MxnetTools_imageHeader_t_ is a type alias for MxnetToolsImageHeaderT for use in objc.Send[T] calls.
type MxnetTools_imageHeader_t_ = MxnetToolsImageHeaderT

// MxnetToolsImageIDT
type MxnetToolsImageIDT struct {
}

// MxnetTools_imageID_t_ is a type alias for MxnetToolsImageIDT for use in objc.Send[T] calls.
type MxnetTools_imageID_t_ = MxnetToolsImageIDT

// MxnetToolsRecordHeaderT
type MxnetToolsRecordHeaderT struct {
}

// MxnetTools_recordHeader_t_ is a type alias for MxnetToolsRecordHeaderT for use in objc.Send[T] calls.
type MxnetTools_recordHeader_t_ = MxnetToolsRecordHeaderT

// Net
type Net struct {
}

// NetStridesConfiguration
type NetStridesConfiguration struct {
}

// Net_strides_configuration is a type alias for NetStridesConfiguration for use in objc.Send[T] calls.
type Net_strides_configuration = NetStridesConfiguration

// OpaquePthreadCond
type OpaquePthreadCond struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// Opaque_pthread_cond_t is a type alias for OpaquePthreadCond for use in objc.Send[T] calls.
type Opaque_pthread_cond_t = OpaquePthreadCond

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

// PaddingParams
type PaddingParams struct {
}

// Padding_params_t is a type alias for PaddingParams for use in objc.Send[T] calls.
type Padding_params_t = PaddingParams

// PostprocessingSettings
type PostprocessingSettings struct {
	Name                               unsafe.Pointer
	Network                            unsafe.Pointer
	Do_blend                           int
	Blend_alpha                        float32
	Grayscale_i0                       int
	Grayscale_i1                       int
	Temporal_alpha                     float32
	Gamma                              float32
	Saturation                         float32
	Contrast                           float32
	Brightness                         float32
	Color_transfer_mode                int
	Width                              int
	Height                             int
	Width_fast                         int
	Height_fast                        int
	Width_capture                      int
	Height_capture                     int
	Width_miniature                    int
	Height_miniature                   int
	Width_fullscreen                   int
	Height_fullscreen                  int
	Width_hd                           int
	Height_hd                          int
	Old_frame_scale                    float32
	Noise_strength                     float32
	Dyn_noise                          int
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

// Vimage2espressoParam
type Vimage2espressoParam struct {
	Scale                    float32
	Center_mean              int
	Is_image_bgr             int
	Is_network_bgr           int
	Bias_r                   float32
	Bias_g                   float32
	Bias_b                   float32
	Bias_a                   float32
	Metal_output_plane       int
	Width                    uint
	Height                   uint
	Rowbytes                 uint
	Rotate_deg               int
	Use_direct_cvpixelbuffer int
	No_alpha_premultiply     int
}

// Vimage2espresso_param is a type alias for Vimage2espressoParam for use in objc.Send[T] calls.
type Vimage2espresso_param = Vimage2espressoParam
