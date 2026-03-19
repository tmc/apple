// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (ec EspressoFDOverfeatNetworkClass) Alloc() EspressoFDOverfeatNetwork {
	rv := objc.Send[EspressoFDOverfeatNetwork](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoFDOverfeatNetwork.AutoResizeForAspectRatioUseLowPriorityModeGpuPriority]
//   - [EspressoFDOverfeatNetwork.AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority]
//   - [EspressoFDOverfeatNetwork.Basename]
//   - [EspressoFDOverfeatNetwork.SetBasename]
//   - [EspressoFDOverfeatNetwork.BoxBlobForScale]
//   - [EspressoFDOverfeatNetwork.Context_cpu]
//   - [EspressoFDOverfeatNetwork.SetContext_cpu]
//   - [EspressoFDOverfeatNetwork.Context_metal]
//   - [EspressoFDOverfeatNetwork.SetContext_metal]
//   - [EspressoFDOverfeatNetwork.Cpin]
//   - [EspressoFDOverfeatNetwork.SetCpin]
//   - [EspressoFDOverfeatNetwork.Cpu_net]
//   - [EspressoFDOverfeatNetwork.Default_retile_outputs]
//   - [EspressoFDOverfeatNetwork.ErrorForLayers]
//   - [EspressoFDOverfeatNetwork.ForceMaxNScales]
//   - [EspressoFDOverfeatNetwork.SetForceMaxNScales]
//   - [EspressoFDOverfeatNetwork.Forward_cpu_network_at_indexPyr]
//   - [EspressoFDOverfeatNetwork.GeneratePyramidTex]
//   - [EspressoFDOverfeatNetwork.GetNumScales]
//   - [EspressoFDOverfeatNetwork.GetScale]
//   - [EspressoFDOverfeatNetwork.Gpu_net]
//   - [EspressoFDOverfeatNetwork.MaxScale]
//   - [EspressoFDOverfeatNetwork.SetMaxScale]
//   - [EspressoFDOverfeatNetwork.Mode]
//   - [EspressoFDOverfeatNetwork.SetMode]
//   - [EspressoFDOverfeatNetwork.NeedRetiling]
//   - [EspressoFDOverfeatNetwork.ProbBlobForScale]
//   - [EspressoFDOverfeatNetwork.ProcessBlobTex]
//   - [EspressoFDOverfeatNetwork.ProcessBlobNoRotationTexDoBGRA2RGBA]
//   - [EspressoFDOverfeatNetwork.ProcessPyramid]
//   - [EspressoFDOverfeatNetwork.ProcessPyramidGpu_resizer]
//   - [EspressoFDOverfeatNetwork.ProcessVimageNoRotationTexDoBGRA2RGBA]
//   - [EspressoFDOverfeatNetwork.Reset]
//   - [EspressoFDOverfeatNetwork.ResizerAtIndex]
//   - [EspressoFDOverfeatNetwork.ResizerCount]
//   - [EspressoFDOverfeatNetwork.Retile_and_forward_espresso_gpu_network_at_indexNetPyr]
//   - [EspressoFDOverfeatNetwork.Retile_and_forward_espresso_network_at_indexNetPyr]
//   - [EspressoFDOverfeatNetwork.RetryLoadingCaffeNetNameWeightsContextCp]
//   - [EspressoFDOverfeatNetwork.ScaleConfig]
//   - [EspressoFDOverfeatNetwork.SetScaleConfig]
//   - [EspressoFDOverfeatNetwork.ScalingMode]
//   - [EspressoFDOverfeatNetwork.SetScalingMode]
//   - [EspressoFDOverfeatNetwork.SetContextCpu]
//   - [EspressoFDOverfeatNetwork.SetContextMetal]
//   - [EspressoFDOverfeatNetwork.Setup_retile]
//   - [EspressoFDOverfeatNetwork.StoreDataForPruningProb]
//   - [EspressoFDOverfeatNetwork.StrideConfiguration]
//   - [EspressoFDOverfeatNetwork.UseGPUScaler]
//   - [EspressoFDOverfeatNetwork.SetUseGPUScaler]
//   - [EspressoFDOverfeatNetwork.Weights]
//   - [EspressoFDOverfeatNetwork.SetWeights]
//   - [EspressoFDOverfeatNetwork.WipeLayersMemory]
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork
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
//   - [IEspressoFDOverfeatNetwork.BoxBlobForScale]
//   - [IEspressoFDOverfeatNetwork.Context_cpu]
//   - [IEspressoFDOverfeatNetwork.SetContext_cpu]
//   - [IEspressoFDOverfeatNetwork.Context_metal]
//   - [IEspressoFDOverfeatNetwork.SetContext_metal]
//   - [IEspressoFDOverfeatNetwork.Cpin]
//   - [IEspressoFDOverfeatNetwork.SetCpin]
//   - [IEspressoFDOverfeatNetwork.Cpu_net]
//   - [IEspressoFDOverfeatNetwork.Default_retile_outputs]
//   - [IEspressoFDOverfeatNetwork.ErrorForLayers]
//   - [IEspressoFDOverfeatNetwork.ForceMaxNScales]
//   - [IEspressoFDOverfeatNetwork.SetForceMaxNScales]
//   - [IEspressoFDOverfeatNetwork.Forward_cpu_network_at_indexPyr]
//   - [IEspressoFDOverfeatNetwork.GeneratePyramidTex]
//   - [IEspressoFDOverfeatNetwork.GetNumScales]
//   - [IEspressoFDOverfeatNetwork.GetScale]
//   - [IEspressoFDOverfeatNetwork.Gpu_net]
//   - [IEspressoFDOverfeatNetwork.MaxScale]
//   - [IEspressoFDOverfeatNetwork.SetMaxScale]
//   - [IEspressoFDOverfeatNetwork.Mode]
//   - [IEspressoFDOverfeatNetwork.SetMode]
//   - [IEspressoFDOverfeatNetwork.NeedRetiling]
//   - [IEspressoFDOverfeatNetwork.ProbBlobForScale]
//   - [IEspressoFDOverfeatNetwork.ProcessBlobTex]
//   - [IEspressoFDOverfeatNetwork.ProcessBlobNoRotationTexDoBGRA2RGBA]
//   - [IEspressoFDOverfeatNetwork.ProcessPyramid]
//   - [IEspressoFDOverfeatNetwork.ProcessPyramidGpu_resizer]
//   - [IEspressoFDOverfeatNetwork.ProcessVimageNoRotationTexDoBGRA2RGBA]
//   - [IEspressoFDOverfeatNetwork.Reset]
//   - [IEspressoFDOverfeatNetwork.ResizerAtIndex]
//   - [IEspressoFDOverfeatNetwork.ResizerCount]
//   - [IEspressoFDOverfeatNetwork.Retile_and_forward_espresso_gpu_network_at_indexNetPyr]
//   - [IEspressoFDOverfeatNetwork.Retile_and_forward_espresso_network_at_indexNetPyr]
//   - [IEspressoFDOverfeatNetwork.RetryLoadingCaffeNetNameWeightsContextCp]
//   - [IEspressoFDOverfeatNetwork.ScaleConfig]
//   - [IEspressoFDOverfeatNetwork.SetScaleConfig]
//   - [IEspressoFDOverfeatNetwork.ScalingMode]
//   - [IEspressoFDOverfeatNetwork.SetScalingMode]
//   - [IEspressoFDOverfeatNetwork.SetContextCpu]
//   - [IEspressoFDOverfeatNetwork.SetContextMetal]
//   - [IEspressoFDOverfeatNetwork.Setup_retile]
//   - [IEspressoFDOverfeatNetwork.StoreDataForPruningProb]
//   - [IEspressoFDOverfeatNetwork.StrideConfiguration]
//   - [IEspressoFDOverfeatNetwork.UseGPUScaler]
//   - [IEspressoFDOverfeatNetwork.SetUseGPUScaler]
//   - [IEspressoFDOverfeatNetwork.Weights]
//   - [IEspressoFDOverfeatNetwork.SetWeights]
//   - [IEspressoFDOverfeatNetwork.WipeLayersMemory]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork
type IEspressoFDOverfeatNetwork interface {
	objectivec.IObject

	// Topic: Methods

	AutoResizeForAspectRatioUseLowPriorityModeGpuPriority(ratio float32, mode bool, priority uint32)
	AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority(name objectivec.IObject, weights objectivec.IObject, config int, mode int, path int, ratio float32, reset bool, mode2 bool, priority uint32)
	Basename() string
	SetBasename(value string)
	BoxBlobForScale(scale int) objectivec.IObject
	Context_cpu() objectivec.IObject
	SetContext_cpu(value objectivec.IObject)
	Context_metal() objectivec.IObject
	SetContext_metal(value objectivec.IObject)
	Cpin() int
	SetCpin(value int)
	Cpu_net(cpu_net int) objectivec.IObject
	Default_retile_outputs() int
	ErrorForLayers() objectivec.IObject
	ForceMaxNScales() int
	SetForceMaxNScales(value int)
	Forward_cpu_network_at_indexPyr(forward_cpu_network_at_index int, pyr unsafe.Pointer)
	GeneratePyramidTex(pyramid unsafe.Pointer, tex objectivec.IObject)
	GetNumScales() int
	GetScale(scale int) float64
	Gpu_net(gpu_net int) objectivec.IObject
	MaxScale() float32
	SetMaxScale(value float32)
	Mode() int
	SetMode(value int)
	NeedRetiling(retiling int) bool
	ProbBlobForScale(scale int) objectivec.IObject
	ProcessBlobTex(blob unsafe.Pointer, tex objectivec.IObject)
	ProcessBlobNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool)
	ProcessPyramid(pyramid objectivec.IObject)
	ProcessPyramidGpu_resizer(pyramid objectivec.IObject, gpu_resizer objectivec.IObject)
	ProcessVimageNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool)
	Reset()
	ResizerAtIndex(index int) objectivec.IObject
	ResizerCount() int
	Retile_and_forward_espresso_gpu_network_at_indexNetPyr(retile_and_forward_espresso_gpu_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer)
	Retile_and_forward_espresso_network_at_indexNetPyr(retile_and_forward_espresso_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer)
	RetryLoadingCaffeNetNameWeightsContextCp(net unsafe.Pointer, name objectivec.IObject, weights objectivec.IObject, context objectivec.IObject, cp int)
	ScaleConfig() int
	SetScaleConfig(value int)
	ScalingMode() int
	SetScalingMode(value int)
	SetContextCpu(cpu objectivec.IObject)
	SetContextMetal(metal objectivec.IObject)
	Setup_retile()
	StoreDataForPruningProb(pruning objectivec.IObject, prob float32)
	StrideConfiguration() objectivec.IObject
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

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/autoResizeForAspectRatio:useLowPriorityMode:gpuPriority:
func (e EspressoFDOverfeatNetwork) AutoResizeForAspectRatioUseLowPriorityModeGpuPriority(ratio float32, mode bool, priority uint32) {
	objc.Send[objc.ID](e.ID, objc.Sel("autoResizeForAspectRatio:useLowPriorityMode:gpuPriority:"), ratio, mode, priority)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/autoSetupNetBaseName:weights:scaleConfig:setupMode:computePath:autoAspectRatio:forceReset:useLowPriorityMode:gpuPriority:
func (e EspressoFDOverfeatNetwork) AutoSetupNetBaseNameWeightsScaleConfigSetupModeComputePathAutoAspectRatioForceResetUseLowPriorityModeGpuPriority(name objectivec.IObject, weights objectivec.IObject, config int, mode int, path int, ratio float32, reset bool, mode2 bool, priority uint32) {
	objc.Send[objc.ID](e.ID, objc.Sel("autoSetupNetBaseName:weights:scaleConfig:setupMode:computePath:autoAspectRatio:forceReset:useLowPriorityMode:gpuPriority:"), name, weights, config, mode, path, ratio, reset, mode2, priority)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/boxBlobForScale:
func (e EspressoFDOverfeatNetwork) BoxBlobForScale(scale int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("boxBlobForScale:"), scale)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/cpu_net:
func (e EspressoFDOverfeatNetwork) Cpu_net(cpu_net int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("cpu_net:"), cpu_net)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/default_retile_outputs
func (e EspressoFDOverfeatNetwork) Default_retile_outputs() int {
	rv := objc.Send[int](e.ID, objc.Sel("default_retile_outputs"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/errorForLayers
func (e EspressoFDOverfeatNetwork) ErrorForLayers() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("errorForLayers"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/forward_cpu_network_at_index:pyr:
func (e EspressoFDOverfeatNetwork) Forward_cpu_network_at_indexPyr(forward_cpu_network_at_index int, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("forward_cpu_network_at_index:pyr:"), forward_cpu_network_at_index, pyr)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/generatePyramid:tex:
func (e EspressoFDOverfeatNetwork) GeneratePyramidTex(pyramid unsafe.Pointer, tex objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("generatePyramid:tex:"), pyramid, tex)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/getNumScales
func (e EspressoFDOverfeatNetwork) GetNumScales() int {
	rv := objc.Send[int](e.ID, objc.Sel("getNumScales"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/getScale:
func (e EspressoFDOverfeatNetwork) GetScale(scale int) float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("getScale:"), scale)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/gpu_net:
func (e EspressoFDOverfeatNetwork) Gpu_net(gpu_net int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("gpu_net:"), gpu_net)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/needRetiling:
func (e EspressoFDOverfeatNetwork) NeedRetiling(retiling int) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("needRetiling:"), retiling)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/probBlobForScale:
func (e EspressoFDOverfeatNetwork) ProbBlobForScale(scale int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("probBlobForScale:"), scale)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/processBlob:tex:
func (e EspressoFDOverfeatNetwork) ProcessBlobTex(blob unsafe.Pointer, tex objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("processBlob:tex:"), blob, tex)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/processBlobNoRotation:tex:doBGRA2RGBA:
func (e EspressoFDOverfeatNetwork) ProcessBlobNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("processBlobNoRotation:tex:doBGRA2RGBA:"), rotation, tex, bgra2rgba)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/processPyramid:
func (e EspressoFDOverfeatNetwork) ProcessPyramid(pyramid objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("processPyramid:"), pyramid)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/processPyramid:gpu_resizer:
func (e EspressoFDOverfeatNetwork) ProcessPyramidGpu_resizer(pyramid objectivec.IObject, gpu_resizer objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("processPyramid:gpu_resizer:"), pyramid, gpu_resizer)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/processVimageNoRotation:tex:doBGRA2RGBA:
func (e EspressoFDOverfeatNetwork) ProcessVimageNoRotationTexDoBGRA2RGBA(rotation unsafe.Pointer, tex objectivec.IObject, bgra2rgba bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("processVimageNoRotation:tex:doBGRA2RGBA:"), rotation, tex, bgra2rgba)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/reset
func (e EspressoFDOverfeatNetwork) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/resizerAtIndex:
func (e EspressoFDOverfeatNetwork) ResizerAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("resizerAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/resizerCount
func (e EspressoFDOverfeatNetwork) ResizerCount() int {
	rv := objc.Send[int](e.ID, objc.Sel("resizerCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/retile_and_forward_espresso_gpu_network_at_index:net:pyr:
func (e EspressoFDOverfeatNetwork) Retile_and_forward_espresso_gpu_network_at_indexNetPyr(retile_and_forward_espresso_gpu_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("retile_and_forward_espresso_gpu_network_at_index:net:pyr:"), retile_and_forward_espresso_gpu_network_at_index, net, pyr)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/retile_and_forward_espresso_network_at_index:net:pyr:
func (e EspressoFDOverfeatNetwork) Retile_and_forward_espresso_network_at_indexNetPyr(retile_and_forward_espresso_network_at_index int, net unsafe.Pointer, pyr unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("retile_and_forward_espresso_network_at_index:net:pyr:"), retile_and_forward_espresso_network_at_index, net, pyr)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/retryLoadingCaffeNet:name:weights:context:cp:
func (e EspressoFDOverfeatNetwork) RetryLoadingCaffeNetNameWeightsContextCp(net unsafe.Pointer, name objectivec.IObject, weights objectivec.IObject, context objectivec.IObject, cp int) {
	objc.Send[objc.ID](e.ID, objc.Sel("retryLoadingCaffeNet:name:weights:context:cp:"), net, name, weights, context, cp)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/setContextCpu:
func (e EspressoFDOverfeatNetwork) SetContextCpu(cpu objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setContextCpu:"), cpu)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/setContextMetal:
func (e EspressoFDOverfeatNetwork) SetContextMetal(metal objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setContextMetal:"), metal)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/setup_retile
func (e EspressoFDOverfeatNetwork) Setup_retile() {
	objc.Send[objc.ID](e.ID, objc.Sel("setup_retile"))
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/storeDataForPruning:prob:
func (e EspressoFDOverfeatNetwork) StoreDataForPruningProb(pruning objectivec.IObject, prob float32) {
	objc.Send[objc.ID](e.ID, objc.Sel("storeDataForPruning:prob:"), pruning, prob)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/strideConfiguration
func (e EspressoFDOverfeatNetwork) StrideConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("strideConfiguration"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/wipeLayersMemory
func (e EspressoFDOverfeatNetwork) WipeLayersMemory() {
	objc.Send[objc.ID](e.ID, objc.Sel("wipeLayersMemory"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/basename
func (e EspressoFDOverfeatNetwork) Basename() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("basename"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoFDOverfeatNetwork) SetBasename(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setBasename:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/context_cpu
func (e EspressoFDOverfeatNetwork) Context_cpu() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("context_cpu"))
	return objectivec.Object{ID: rv}
}
func (e EspressoFDOverfeatNetwork) SetContext_cpu(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setContext_cpu:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/context_metal
func (e EspressoFDOverfeatNetwork) Context_metal() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("context_metal"))
	return objectivec.Object{ID: rv}
}
func (e EspressoFDOverfeatNetwork) SetContext_metal(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setContext_metal:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/cpin
func (e EspressoFDOverfeatNetwork) Cpin() int {
	rv := objc.Send[int](e.ID, objc.Sel("cpin"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetCpin(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setCpin:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/forceMaxNScales
func (e EspressoFDOverfeatNetwork) ForceMaxNScales() int {
	rv := objc.Send[int](e.ID, objc.Sel("forceMaxNScales"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetForceMaxNScales(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setForceMaxNScales:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/maxScale
func (e EspressoFDOverfeatNetwork) MaxScale() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("maxScale"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetMaxScale(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxScale:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/mode
func (e EspressoFDOverfeatNetwork) Mode() int {
	rv := objc.Send[int](e.ID, objc.Sel("mode"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetMode(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setMode:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/scaleConfig
func (e EspressoFDOverfeatNetwork) ScaleConfig() int {
	rv := objc.Send[int](e.ID, objc.Sel("scaleConfig"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetScaleConfig(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setScaleConfig:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/scalingMode
func (e EspressoFDOverfeatNetwork) ScalingMode() int {
	rv := objc.Send[int](e.ID, objc.Sel("scalingMode"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetScalingMode(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setScalingMode:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/useGPUScaler
func (e EspressoFDOverfeatNetwork) UseGPUScaler() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("useGPUScaler"))
	return rv
}
func (e EspressoFDOverfeatNetwork) SetUseGPUScaler(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setUseGPUScaler:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFDOverfeatNetwork/weights
func (e EspressoFDOverfeatNetwork) Weights() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("weights"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoFDOverfeatNetwork) SetWeights(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setWeights:"), objc.String(value))
}

