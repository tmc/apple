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

// _CFArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFArray
type _CFArray struct {
}

// _CFDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFDictionary
type _CFDictionary struct {
}

// _CFString
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CFString
type _CFString struct {
}

// _CVBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__CVBuffer
type _CVBuffer struct {
}

// _IOSurface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__IOSurface
type _IOSurface struct {
}

// _long
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__long
type _long struct {
	__data_    *byte
	__size_    uint64
	__cap_     objectivec.Object
	__is_long_ objectivec.Object
}

// _sFILE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sFILE
type _sFILE struct {
}

// _sFILEX
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sFILEX
type _sFILEX struct {
}

// _sbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__sbuf
type _sbuf struct {
}

// _shared_weak_count
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__shared_weak_count
type _shared_weak_count struct {
}

// _short
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/__short
type _short struct {
	__data_    unsafe.Pointer
	__size_    objectivec.Object
	__is_long_ objectivec.Object
}

// Abstract_blob_container
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/abstract_blob_container
type Abstract_blob_container struct {
}

// Abstract_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/abstract_context
type Abstract_context struct {
}

// Blob_cpu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/blob_cpu
type Blob_cpu struct {
}

// Condition_variable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/condition_variable
type Condition_variable struct {
	__cv_ Opaque_pthread_cond_t
}

// Convolution_uniforms
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/convolution_uniforms
type Convolution_uniforms struct {
}

// Fast_pyramid_resizer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/fast_pyramid_resizer
type Fast_pyramid_resizer struct {
}

// Float_buffer_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/float_buffer_t
type Float_buffer_t struct {
}

// Generic_load_constant_kernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/generic_load_constant_kernel
type Generic_load_constant_kernel struct {
}

// Gradient_builder
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/gradient_builder
type Gradient_builder struct {
}

// Inner_product_uniforms
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/inner_product_uniforms
type Inner_product_uniforms struct {
}

// MxnetTools_imageHeader_t_
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_imageHeader_t_
type MxnetTools_imageHeader_t_ struct {
}

// MxnetTools_imageID_t_
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_imageID_t_
type MxnetTools_imageID_t_ struct {
}

// MxnetTools_recordHeader_t_
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_mxnetTools_recordHeader_t_
type MxnetTools_recordHeader_t_ struct {
}

// Net
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/net
type Net struct {
}

// Net_strides_configuration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/net_strides_configuration
type Net_strides_configuration struct {
}

// Opaque_pthread_cond_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_opaque_pthread_cond_t
type Opaque_pthread_cond_t struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// Opaque_pthread_mutex_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/_opaque_pthread_mutex_t
type Opaque_pthread_mutex_t struct {
	__sig    int64
	__opaque unsafe.Pointer
}

// Padding_params_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/padding_params_t
type Padding_params_t struct {
}

// Postprocessing_settings_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/postprocessing_settings_t
type Postprocessing_settings_t struct {
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

// Surface_and_buffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/surface_and_buffer
type Surface_and_buffer struct {
}

// V9_noise_kernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/v9_noise_kernel
type V9_noise_kernel struct {
}

// Vimage2espresso_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleNeuralEngine/vimage2espresso_param
type Vimage2espresso_param struct {
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
