// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoFDOverfeatNetwork] class.
var (
	_EspressoFDOverfeatNetworkClass     EspressoFDOverfeatNetworkClass
	_EspressoFDOverfeatNetworkClassOnce sync.Once
)

func getEspressoFDOverfeatNetworkClass() EspressoFDOverfeatNetworkClass {
	_EspressoFDOverfeatNetworkClassOnce.Do(func() {
		_EspressoFDOverfeatNetworkClass = EspressoFDOverfeatNetworkClass{class: objc.GetClass("EspressoFDOverfeatNetwork")}
	})
	return _EspressoFDOverfeatNetworkClass
}

// GetEspressoFDOverfeatNetworkClass returns the class object for EspressoFDOverfeatNetwork.
func GetEspressoFDOverfeatNetworkClass() EspressoFDOverfeatNetworkClass {
	return getEspressoFDOverfeatNetworkClass()
}

type EspressoFDOverfeatNetworkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoFDOverfeatNetworkClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoFDOverfeatNetworkClass) Alloc() EspressoFDOverfeatNetwork {
	rv := objc.Send[EspressoFDOverfeatNetwork](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoFDOverfeatNetwork.AutoResizeForAspectRatioUseLowPriorityModeGpuPriority]
//   - [EspressoFDOverfeatNetwork.AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority]
//   - [EspressoFDOverfeatNetwork.Basename]
//   - [EspressoFDOverfeatNetwork.SetBasename]
//   - [EspressoFDOverfeatNetwork.Context_cpu]
//   - [EspressoFDOverfeatNetwork.SetContext_cpu]
//   - [EspressoFDOverfeatNetwork.Context_metal]
//   - [EspressoFDOverfeatNetwork.SetContext_metal]
//   - [EspressoFDOverfeatNetwork.Cpin]
//   - [EspressoFDOverfeatNetwork.SetCpin]
//   - [EspressoFDOverfeatNetwork.Default_retile_outputs]
//   - [EspressoFDOverfeatNetwork.ErrorForLayers]
//   - [EspressoFDOverfeatNetwork.ForceMaxNScales]
//   - [EspressoFDOverfeatNetwork.SetForceMaxNScales]
//   - [EspressoFDOverfeatNetwork.Forward_cpu_network_at_indexPyr]
//   - [EspressoFDOverfeatNetwork.GeneratePyramidTex]
//   - [EspressoFDOverfeatNetwork.GetNumScales]
//   - [EspressoFDOverfeatNetwork.GetScale]
//   - [EspressoFDOverfeatNetwork.MaxScale]
//   - [EspressoFDOverfeatNetwork.SetMaxScale]
//   - [EspressoFDOverfeatNetwork.Mode]
//   - [EspressoFDOverfeatNetwork.SetMode]
//   - [EspressoFDOverfeatNetwork.NeedRetiling]
//   - [EspressoFDOverfeatNetwork.ProcessBlobTex]
//   - [EspressoFDOverfeatNetwork.ProcessBlobNoRotationTexDoBGRA2RGBA]
//   - [EspressoFDOverfeatNetwork.ProcessVimageNoRotationTexDoBGRA2RGBA]
//   - [EspressoFDOverfeatNetwork.Reset]
//   - [EspressoFDOverfeatNetwork.ResizerCount]
//   - [EspressoFDOverfeatNetwork.Retile_and_forward_espresso_gpu_network_at_indexNetPyr]
//   - [EspressoFDOverfeatNetwork.Retile_and_forward_espresso_network_at_indexNetPyr]
//   - [EspressoFDOverfeatNetwork.ScaleConfig]
//   - [EspressoFDOverfeatNetwork.SetScaleConfig]
//   - [EspressoFDOverfeatNetwork.ScalingMode]
//   - [EspressoFDOverfeatNetwork.SetScalingMode]
//   - [EspressoFDOverfeatNetwork.SetContextCpu]
//   - [EspressoFDOverfeatNetwork.SetContextMetal]
//   - [EspressoFDOverfeatNetwork.Setup_retile]
//   - [EspressoFDOverfeatNetwork.StrideConfiguration]
//   - [EspressoFDOverfeatNetwork.UseGPUScaler]
//   - [EspressoFDOverfeatNetwork.SetUseGPUScaler]
//   - [EspressoFDOverfeatNetwork.Weights]
//   - [EspressoFDOverfeatNetwork.SetWeights]
//   - [EspressoFDOverfeatNetwork.WipeLayersMemory]
type EspressoFDOverfeatNetwork struct {
	objectivec.Object
}

// EspressoFDOverfeatNetworkFromID constructs a [EspressoFDOverfeatNetwork] from an objc.ID.
func EspressoFDOverfeatNetworkFromID(id objc.ID) EspressoFDOverfeatNetwork {
	return EspressoFDOverfeatNetwork{objectivec.Object{ID: id}}
}

// Ensure EspressoFDOverfeatNetwork implements IEspressoFDOverfeatNetwork.
var _ IEspressoFDOverfeatNetwork = EspressoFDOverfeatNetwork{}

// An interface definition for the [EspressoFDOverfeatNetwork] class.
//
// # Methods
//
//   - [IEspressoFDOverfeatNetwork.AutoResizeForAspectRatioUseLowPriorityModeGpuPriority]
//   - [IEspressoFDOverfeatNetwork.AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority]
//   - [IEspressoFDOverfeatNetwork.Basename]
//   - [IEspressoFDOverfeatNetwork.SetBasename]
//   - [IEspressoFDOverfeatNetwork.Context_cpu]
//   - [IEspressoFDOverfeatNetwork.SetContext_cpu]
//   - [IEspressoFDOverfeatNetwork.Context_metal]
//   - [IEspressoFDOverfeatNetwork.SetContext_metal]
//   - [IEspressoFDOverfeatNetwork.Cpin]
//   - [IEspressoFDOverfeatNetwork.SetCpin]
//   - [IEspressoFDOverfeatNetwork.Default_retile_outputs]
//   - [IEspressoFDOverfeatNetwork.ErrorForLayers]
//   - [IEspressoFDOverfeatNetwork.ForceMaxNScales]
//   - [IEspressoFDOverfeatNetwork.SetForceMaxNScales]
//   - [IEspressoFDOverfeatNetwork.Forward_cpu_network_at_indexPyr]
//   - [IEspressoFDOverfeatNetwork.GeneratePyramidTex]
//   - [IEspressoFDOverfeatNetwork.GetNumScales]
//   - [IEspressoFDOverfeatNetwork.GetScale]
//   - [IEspressoFDOverfeatNetwork.MaxScale]
//   - [IEspressoFDOverfeatNetwork.SetMaxScale]
//   - [IEspressoFDOverfeatNetwork.Mode]
//   - [IEspressoFDOverfeatNetwork.SetMode]
//   - [IEspressoFDOverfeatNetwork.NeedRetiling]
//   - [IEspressoFDOverfeatNetwork.ProcessBlobTex]
//   - [IEspressoFDOverfeatNetwork.ProcessBlobNoRotationTexDoBGRA2RGBA]
//   - [IEspressoFDOverfeatNetwork.ProcessVimageNoRotationTexDoBGRA2RGBA]
//   - [IEspressoFDOverfeatNetwork.Reset]
//   - [IEspressoFDOverfeatNetwork.ResizerCount]
//   - [IEspressoFDOverfeatNetwork.Retile_and_forward_espresso_gpu_network_at_indexNetPyr]
//   - [IEspressoFDOverfeatNetwork.Retile_and_forward_espresso_network_at_indexNetPyr]
//   - [IEspressoFDOverfeatNetwork.ScaleConfig]
//   - [IEspressoFDOverfeatNetwork.SetScaleConfig]
//   - [IEspressoFDOverfeatNetwork.ScalingMode]
//   - [IEspressoFDOverfeatNetwork.SetScalingMode]
//   - [IEspressoFDOverfeatNetwork.SetContextCpu]
//   - [IEspressoFDOverfeatNetwork.SetContextMetal]
//   - [IEspressoFDOverfeatNetwork.Setup_retile]
//   - [IEspressoFDOverfeatNetwork.StrideConfiguration]
//   - [IEspressoFDOverfeatNetwork.UseGPUScaler]
//   - [IEspressoFDOverfeatNetwork.SetUseGPUScaler]
//   - [IEspressoFDOverfeatNetwork.Weights]
//   - [IEspressoFDOverfeatNetwork.SetWeights]
//   - [IEspressoFDOverfeatNetwork.WipeLayersMemory]
type IEspressoFDOverfeatNetwork interface {
	objectivec.IObject

	// Topic: Methods

	AutoResizeForAspectRatioUseLowPriorityModeGpuPriority(ratio float32, mode bool, priority uint32)
	AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority(name objectivec.IObject, weights objectivec.IObject, config int, mode int, path int, ratio float32, reset bool, mode2 bool, priority uint32)
	Basename() string
	SetBasename(value string)
	Context_cpu() unsafe.Pointer
	SetContext_cpu(value unsafe.Pointer)
	Context_metal() unsafe.Pointer
	SetContext_metal(value unsafe.Pointer)
	Cpin() int
	SetCpin(value int)
	Default_retile_outputs() int
	ErrorForLayers() objectivec.IObject
	ForceMaxNScales() int
	SetForceMaxNScales(value int)
	Forward_cpu_network_at_indexPyr(forward_cpu_network_at_index int, pyr unsafe.Pointer)
	GeneratePyramidTex(pyramid unsafe.Pointer, tex objectivec.IObject)
	GetNumScales() int
	GetScale(scale int) float64
	MaxScale() float32
	SetMaxScale(value float32)
	Mode() int
	SetMode(value int)
	NeedRetiling(retiling int) bool
	ProcessBlobTex(blob unsafe.Pointer, tex objectivec.IObject)
	ProcessBlobNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool)
	ProcessVimageNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool)
	Reset()
	ResizerCount() int
	Retile_and_forward_espresso_gpu_network_at_indexNetPyr(retile_and_forward_espresso_gpu_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer)
	Retile_and_forward_espresso_network_at_indexNetPyr(retile_and_forward_espresso_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer)
	ScaleConfig() int
	SetScaleConfig(value int)
	ScalingMode() int
	SetScalingMode(value int)
	SetContextCpu(cpu objectivec.IObject)
	SetContextMetal(metal objectivec.IObject)
	Setup_retile()
	StrideConfiguration() NetStridesConfiguration
	UseGPUScaler() bool
	SetUseGPUScaler(value bool)
	Weights() string
	SetWeights(value string)
	WipeLayersMemory()
}

// Init initializes the instance.
func (e EspressoFDOverfeatNetwork) Init() EspressoFDOverfeatNetwork {
	rv := objc.Send[EspressoFDOverfeatNetwork](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoFDOverfeatNetwork) Autorelease() EspressoFDOverfeatNetwork {
	rv := objc.Send[EspressoFDOverfeatNetwork](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoFDOverfeatNetwork creates a new EspressoFDOverfeatNetwork instance.
func NewEspressoFDOverfeatNetwork() EspressoFDOverfeatNetwork {
	class := getEspressoFDOverfeatNetworkClass()
	rv := objc.Send[EspressoFDOverfeatNetwork](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e EspressoFDOverfeatNetwork) AutoResizeForAspectRatioUseLowPriorityModeGpuPriority(ratio float32, mode bool, priority uint32) {
	objc.Send[objc.ID](e.ID, objc.Sel("autoResizeForAspectRatio:useLowPriorityMode:gpuPriority:"), ratio, mode, priority)
}
func (e EspressoFDOverfeatNetwork) AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority(name objectivec.IObject, weights objectivec.IObject, config int, mode int, path int, ratio float32, reset bool, mode2 bool, priority uint32) {
	objc.Send[objc.ID](e.ID, objc.Sel("autoSetupNetBaseName:weights:scaleConfig:setupMode:computePath:autoAspectRatio:forceReset:useLowPriorityMode:gpuPriority:"), name, weights, config, mode, path, ratio, reset, mode2, priority)
}
func (e EspressoFDOverfeatNetwork) Default_retile_outputs() int {
	rv := objc.Send[int](e.ID, objc.Sel("default_retile_outputs"))
	return rv
}
func (e EspressoFDOverfeatNetwork) ErrorForLayers() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("errorForLayers"))
	return objectivec.Object{ID: rv}
}
func (e EspressoFDOverfeatNetwork) Forward_cpu_network_at_indexPyr(forward_cpu_network_at_index int, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("forward_cpu_network_at_index:pyr:"), forward_cpu_network_at_index, pyr)
}
func (e EspressoFDOverfeatNetwork) GeneratePyramidTex(pyramid unsafe.Pointer, tex objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("generatePyramid:tex:"), pyramid, tex)
}
func (e EspressoFDOverfeatNetwork) GetNumScales() int {
	rv := objc.Send[int](e.ID, objc.Sel("getNumScales"))
	return rv
}
func (e EspressoFDOverfeatNetwork) GetScale(scale int) float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("getScale:"), scale)
	return rv
}
func (e EspressoFDOverfeatNetwork) NeedRetiling(retiling int) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("needRetiling:"), retiling)
	return rv
}
func (e EspressoFDOverfeatNetwork) ProcessBlobTex(blob unsafe.Pointer, tex objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("processBlob:tex:"), blob, tex)
}
func (e EspressoFDOverfeatNetwork) ProcessBlobNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("processBlobNoRotation:tex:doBGRA2RGBA:"), rotation, tex, bgra2rgba)
}
func (e EspressoFDOverfeatNetwork) ProcessVimageNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("processVimageNoRotation:tex:doBGRA2RGBA:"), rotation, tex, bgra2rgba)
}
func (e EspressoFDOverfeatNetwork) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}
func (e EspressoFDOverfeatNetwork) ResizerCount() int {
	rv := objc.Send[int](e.ID, objc.Sel("resizerCount"))
	return rv
}
func (e EspressoFDOverfeatNetwork) Retile_and_forward_espresso_gpu_network_at_indexNetPyr(retile_and_forward_espresso_gpu_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("retile_and_forward_espresso_gpu_network_at_index:net:pyr:"), retile_and_forward_espresso_gpu_network_at_index, net, pyr)
}
func (e EspressoFDOverfeatNetwork) Retile_and_forward_espresso_network_at_indexNetPyr(retile_and_forward_espresso_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("retile_and_forward_espresso_network_at_index:net:pyr:"), retile_and_forward_espresso_network_at_index, net, pyr)
}
func (e EspressoFDOverfeatNetwork) SetContextCpu(cpu objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setContextCpu:"), cpu)
}
func (e EspressoFDOverfeatNetwork) SetContextMetal(metal objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setContextMetal:"), metal)
}
func (e EspressoFDOverfeatNetwork) Setup_retile() {
	objc.Send[objc.ID](e.ID, objc.Sel("setup_retile"))
}
func (e EspressoFDOverfeatNetwork) StrideConfiguration() NetStridesConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("strideConfiguration"))
	return NetStridesConfiguration(rv)
}
func (e EspressoFDOverfeatNetwork) WipeLayersMemory() {
	objc.Send[objc.ID](e.ID, objc.Sel("wipeLayersMemory"))
}

func (e EspressoFDOverfeatNetwork) Basename() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("basename"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoFDOverfeatNetwork) SetBasename(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setBasename:"), objc.String(value))
}
func (e EspressoFDOverfeatNetwork) Context_cpu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("context_cpu"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetContext_cpu(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setContext_cpu:"), value)
}
func (e EspressoFDOverfeatNetwork) Context_metal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("context_metal"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetContext_metal(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setContext_metal:"), value)
}
func (e EspressoFDOverfeatNetwork) Cpin() int {
	rv := objc.Send[int](e.ID, objc.Sel("cpin"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetCpin(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setCpin:"), value)
}
func (e EspressoFDOverfeatNetwork) ForceMaxNScales() int {
	rv := objc.Send[int](e.ID, objc.Sel("forceMaxNScales"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetForceMaxNScales(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setForceMaxNScales:"), value)
}
func (e EspressoFDOverfeatNetwork) MaxScale() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("maxScale"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetMaxScale(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxScale:"), value)
}
func (e EspressoFDOverfeatNetwork) Mode() int {
	rv := objc.Send[int](e.ID, objc.Sel("mode"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetMode(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setMode:"), value)
}
func (e EspressoFDOverfeatNetwork) ScaleConfig() int {
	rv := objc.Send[int](e.ID, objc.Sel("scaleConfig"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetScaleConfig(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setScaleConfig:"), value)
}
func (e EspressoFDOverfeatNetwork) ScalingMode() int {
	rv := objc.Send[int](e.ID, objc.Sel("scalingMode"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetScalingMode(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setScalingMode:"), value)
}
func (e EspressoFDOverfeatNetwork) UseGPUScaler() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("useGPUScaler"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetUseGPUScaler(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setUseGPUScaler:"), value)
}
func (e EspressoFDOverfeatNetwork) Weights() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("weights"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoFDOverfeatNetwork) SetWeights(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setWeights:"), objc.String(value))
}
