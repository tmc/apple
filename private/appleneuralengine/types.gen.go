// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// C struct types

// ANEBufferStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ANEBufferStruct
type ANEBufferStruct struct {
}

// ANEDeviceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ANEDeviceStruct
type ANEDeviceStruct struct {
}

// ANEMemoryMappingParamsStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ANEMemoryMappingParamsStruct
type ANEMemoryMappingParamsStruct struct {
}

// ANENotificationMessageStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ANENotificationMessageStruct
type ANENotificationMessageStruct struct {
}

// AnalyticsData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_AnalyticsData
type AnalyticsData struct {
}

// AnalyticsGroupInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_AnalyticsGroupInfo
type AnalyticsGroupInfo struct {
}

// AnalyticsLayerInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_AnalyticsLayerInfo
type AnalyticsLayerInfo struct {
}

// AnalyticsProcedureInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_AnalyticsProcedureInfo
type AnalyticsProcedureInfo struct {
}

// AnalyticsTaskInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_AnalyticsTaskInfo
type AnalyticsTaskInfo struct {
}

// Attribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/Attribute
type Attribute struct {
}

// AttributeStorage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/AttributeStorage
type AttributeStorage struct {
}

// Autodiff
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/Autodiff
type Autodiff struct {
}

// AutodiffOpData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/AutodiffOpData
type AutodiffOpData struct {
}

// BaseModuleRef
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/BaseModuleRef
type BaseModuleRef struct {
}

// BaseRuntime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/BaseRuntime
type BaseRuntime struct {
}

// BuildVersionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/BuildVersionInfo
type BuildVersionInfo struct {
}

// CGColorSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/CGColorSpace
type CGColorSpace struct {
}

// CGImage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/CGImage
type CGImage struct {
}

// DeviceExtendedInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/DeviceExtendedInfo
type DeviceExtendedInfo struct {
}

// DeviceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/DeviceInfo
type DeviceInfo struct {
}

// FaceLandmarkDetectorPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/FaceLandmarkDetectorPoint
type FaceLandmarkDetectorPoint struct {
}

// FuncOp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/FuncOp
type FuncOp struct {
}

// InMemoryModuleRef
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/InMemoryModuleRef
type InMemoryModuleRef struct {
}

// Location
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/Location
type Location struct {
}

// LocationAttr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/LocationAttr
type LocationAttr struct {
}

// LockFileManager
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/LockFileManager
type LockFileManager struct {
}

// LockGuard
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/LockGuard
type LockGuard struct {
}

// MLIRContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MLIRContext
type MLIRContext struct {
}

// MPSCommandBufferDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCommandBufferDescriptor
type MPSCommandBufferDescriptor struct {
}

// MPSGraphExecutableCacheValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphExecutableCacheValue
type MPSGraphExecutableCacheValue struct {
}

// MPSGraphExecutableSpecializedModule
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphExecutableSpecializedModule
type MPSGraphExecutableSpecializedModule struct {
}

// MPSGraphModuleKey
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphModuleKey
type MPSGraphModuleKey struct {
}

// MPSGraphOperatingSystemVersion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphOperatingSystemVersion
type MPSGraphOperatingSystemVersion struct {
}

// MPSGraphSpecializationCache
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphSpecializationCache
type MPSGraphSpecializationCache struct {
	_moduleStorage       unsafe.Pointer
	_currentCache        unsafe.Pointer
	_failedToLoadModules unsafe.Pointer
}

// MPSResourceBlobEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MPSResourceBlobEntry
type MPSResourceBlobEntry struct {
}

// MemoryBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/MemoryBuffer
type MemoryBuffer struct {
}

// ModuleOp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ModuleOp
type ModuleOp struct {
	State *Operation
}

// ModuleResourcesLoader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ModuleResourcesLoader
type ModuleResourcesLoader struct {
}

// NamedAttribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/NamedAttribute
type NamedAttribute struct {
}

// Operation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/Operation
type Operation struct {
}

// OriginalModuleRef
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/OriginalModuleRef
type OriginalModuleRef struct {
	_module unsafe.Pointer
}

// ReadDataFromFileCache
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ReadDataFromFileCache
type ReadDataFromFileCache struct {
}

// ReturnOp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ReturnOp
type ReturnOp struct {
}

// ShapeFunctionRegistry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ShapeFunctionRegistry
type ShapeFunctionRegistry struct {
}

// StringMapEntryBase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/StringMapEntryBase
type StringMapEntryBase struct {
}

// StringRef
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/StringRef
type StringRef struct {
}

// VMData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/VMData
type VMData struct {
}

// Value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/Value
type Value struct {
}

// ValueImpl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/ValueImpl
type ValueImpl struct {
}

// VirtANEModel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/VirtANEModel
type VirtANEModel struct {
}

// CFArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFArray
type CFArray struct {
}

// CFDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFDictionary
type CFDictionary struct {
}

// CFString
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFString
type CFString struct {
}

// CVBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CVBuffer
type CVBuffer struct {
}

// IOSurface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__IOSurface
type IOSurface struct {
}

// Long
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__long
type Long struct {
	__data_    *byte
	__size_    uint64
	__cap_     objectivec.Object
	__is_long_ objectivec.Object
}

// SFILE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sFILE
type SFILE struct {
}

// SFILEX
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sFILEX
type SFILEX struct {
}

// Sbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sbuf
type Sbuf struct {
}

// SharedWeakCount
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__shared_weak_count
type SharedWeakCount struct {
}

// Short
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__short
type Short struct {
	__data_    unsafe.Pointer
	__size_    objectivec.Object
	__is_long_ objectivec.Object
}

// AbstractBlobContainer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/abstract_blob_container
type AbstractBlobContainer struct {
}

// AbstractContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/abstract_context
type AbstractContext struct {
}

// BlobCPU
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/blob_cpu
type BlobCPU struct {
}

// ConditionVariable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/condition_variable
type ConditionVariable struct {
	__cv_ OpaquePthreadCond
}

// ConvolutionUniforms
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/convolution_uniforms
type ConvolutionUniforms struct {
}

// FastPyramidResizer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/fast_pyramid_resizer
type FastPyramidResizer struct {
}

// FloatBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/float_buffer_t
type FloatBuffer struct {
}

// GenericLoadConstantKernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/generic_load_constant_kernel
type GenericLoadConstantKernel struct {
}

// GradientBuilder
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/gradient_builder
type GradientBuilder struct {
}

// InnerProductUniforms
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/inner_product_uniforms
type InnerProductUniforms struct {
}

// MxnetToolsImageHeaderT
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_imageHeader_t_
type MxnetToolsImageHeaderT struct {
}

// MxnetToolsImageIDT
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_imageID_t_
type MxnetToolsImageIDT struct {
}

// MxnetToolsRecordHeaderT
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_recordHeader_t_
type MxnetToolsRecordHeaderT struct {
}

// Net
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/net
type Net struct {
}

// NetStridesConfiguration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/net_strides_configuration
type NetStridesConfiguration struct {
}

// OpaquePthreadCond
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_opaque_pthread_cond_t
type OpaquePthreadCond struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// OpaquePthreadMutex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_opaque_pthread_mutex_t
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// PaddingParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/padding_params_t
type PaddingParams struct {
}

// PostprocessingSettings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/postprocessing_settings_t
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

// SurfaceAndBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/surface_and_buffer
type SurfaceAndBuffer struct {
}

// V9NoiseKernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/v9_noise_kernel
type V9NoiseKernel struct {
}

// Vimage2espressoParam
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/vimage2espresso_param
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
