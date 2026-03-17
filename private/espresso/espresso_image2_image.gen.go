// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoImage2Image] class.
var (
	_EspressoImage2ImageClass     EspressoImage2ImageClass
	_EspressoImage2ImageClassOnce sync.Once
)

func getEspressoImage2ImageClass() EspressoImage2ImageClass {
	_EspressoImage2ImageClassOnce.Do(func() {
		_EspressoImage2ImageClass = EspressoImage2ImageClass{class: objc.GetClass("EspressoImage2Image")}
	})
	return _EspressoImage2ImageClass
}

// GetEspressoImage2ImageClass returns the class object for EspressoImage2Image.
func GetEspressoImage2ImageClass() EspressoImage2ImageClass {
	return getEspressoImage2ImageClass()
}

type EspressoImage2ImageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoImage2ImageClass) Alloc() EspressoImage2Image {
	rv := objc.Send[EspressoImage2Image](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoImage2Image.AddNoiseLayer]
//   - [EspressoImage2Image.AggregateWisdom]
//   - [EspressoImage2Image.Benchmark]
//   - [EspressoImage2Image.EncodeToCommandBufferSourceTextureDestinationTextureCropRect]
//   - [EspressoImage2Image.EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect]
//   - [EspressoImage2Image.Flip_y]
//   - [EspressoImage2Image.SetFlip_y]
//   - [EspressoImage2Image.GetEspressoNetwork]
//   - [EspressoImage2Image.GetInternalDataForKey]
//   - [EspressoImage2Image.Height]
//   - [EspressoImage2Image.Load]
//   - [EspressoImage2Image.LoadResolutionPreset]
//   - [EspressoImage2Image.NewOutputTexture]
//   - [EspressoImage2Image.PostProcessCameraSourceTextureInputTextureDestinationTexture]
//   - [EspressoImage2Image.Reload]
//   - [EspressoImage2Image.ResetTemporalState]
//   - [EspressoImage2Image.ReshapeToResolutionPreset]
//   - [EspressoImage2Image.ReshapeToResolutionPresetAspectRatio]
//   - [EspressoImage2Image.ReshapeToWidthAndHeight]
//   - [EspressoImage2Image.ResolutionForPreset]
//   - [EspressoImage2Image.Rotation_degrees]
//   - [EspressoImage2Image.SetRotation_degrees]
//   - [EspressoImage2Image.SetupWithQueue]
//   - [EspressoImage2Image.SimpleLinearResizeSourceTextureDestinationTexture]
//   - [EspressoImage2Image.StyleName]
//   - [EspressoImage2Image.SubmitToQueueWithSourceTextureDestinationTexture]
//   - [EspressoImage2Image.SubmitToQueueWithSourceTextureDestinationTextureCropRect]
//   - [EspressoImage2Image.Tune]
//   - [EspressoImage2Image.TweakValue]
//   - [EspressoImage2Image.WasReshaped]
//   - [EspressoImage2Image.Width]
//   - [EspressoImage2Image.InitWithQueue]
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image
type EspressoImage2Image struct {
	objectivec.Object
}

// EspressoImage2ImageFromID constructs a [EspressoImage2Image] from an objc.ID.
func EspressoImage2ImageFromID(id objc.ID) EspressoImage2Image {
	return EspressoImage2Image{objectivec.Object{ID: id}}
}
// Ensure EspressoImage2Image implements IEspressoImage2Image.
var _ IEspressoImage2Image = EspressoImage2Image{}

// An interface definition for the [EspressoImage2Image] class.
//
// # Methods
//
//   - [IEspressoImage2Image.AddNoiseLayer]
//   - [IEspressoImage2Image.AggregateWisdom]
//   - [IEspressoImage2Image.Benchmark]
//   - [IEspressoImage2Image.EncodeToCommandBufferSourceTextureDestinationTextureCropRect]
//   - [IEspressoImage2Image.EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect]
//   - [IEspressoImage2Image.Flip_y]
//   - [IEspressoImage2Image.SetFlip_y]
//   - [IEspressoImage2Image.GetEspressoNetwork]
//   - [IEspressoImage2Image.GetInternalDataForKey]
//   - [IEspressoImage2Image.Height]
//   - [IEspressoImage2Image.Load]
//   - [IEspressoImage2Image.LoadResolutionPreset]
//   - [IEspressoImage2Image.NewOutputTexture]
//   - [IEspressoImage2Image.PostProcessCameraSourceTextureInputTextureDestinationTexture]
//   - [IEspressoImage2Image.Reload]
//   - [IEspressoImage2Image.ResetTemporalState]
//   - [IEspressoImage2Image.ReshapeToResolutionPreset]
//   - [IEspressoImage2Image.ReshapeToResolutionPresetAspectRatio]
//   - [IEspressoImage2Image.ReshapeToWidthAndHeight]
//   - [IEspressoImage2Image.ResolutionForPreset]
//   - [IEspressoImage2Image.Rotation_degrees]
//   - [IEspressoImage2Image.SetRotation_degrees]
//   - [IEspressoImage2Image.SetupWithQueue]
//   - [IEspressoImage2Image.SimpleLinearResizeSourceTextureDestinationTexture]
//   - [IEspressoImage2Image.StyleName]
//   - [IEspressoImage2Image.SubmitToQueueWithSourceTextureDestinationTexture]
//   - [IEspressoImage2Image.SubmitToQueueWithSourceTextureDestinationTextureCropRect]
//   - [IEspressoImage2Image.Tune]
//   - [IEspressoImage2Image.TweakValue]
//   - [IEspressoImage2Image.WasReshaped]
//   - [IEspressoImage2Image.Width]
//   - [IEspressoImage2Image.InitWithQueue]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image
type IEspressoImage2Image interface {
	objectivec.IObject

	// Topic: Methods

	AddNoiseLayer()
	AggregateWisdom(wisdom unsafe.Pointer)
	Benchmark() float32
	EncodeToCommandBufferSourceTextureDestinationTextureCropRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject) int
	EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject, rect2 objectivec.IObject) int
	Flip_y() int
	SetFlip_y(value int)
	GetEspressoNetwork() objectivec.IObject
	GetInternalDataForKey(key objectivec.IObject) objectivec.IObject
	Height() int
	Load(load objectivec.IObject) int
	LoadResolutionPreset(load objectivec.IObject, preset int64) int
	NewOutputTexture() objectivec.IObject
	PostProcessCameraSourceTextureInputTextureDestinationTexture(process objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, texture3 objectivec.IObject)
	Reload() int
	ResetTemporalState()
	ReshapeToResolutionPreset(preset int64) int
	ReshapeToResolutionPresetAspectRatio(preset int64, ratio float32) int
	ReshapeToWidthAndHeight(width int, height int) int
	ResolutionForPreset(preset int64) objectivec.IObject
	Rotation_degrees() int
	SetRotation_degrees(value int)
	SetupWithQueue(queue objectivec.IObject) objectivec.IObject
	SimpleLinearResizeSourceTextureDestinationTexture(resize objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject)
	StyleName() objectivec.IObject
	SubmitToQueueWithSourceTextureDestinationTexture(texture objectivec.IObject, texture2 objectivec.IObject) int
	SubmitToQueueWithSourceTextureDestinationTextureCropRect(texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject) int
	Tune()
	TweakValue(tweak objectivec.IObject, value float32)
	WasReshaped() int
	Width() int
	InitWithQueue(queue objectivec.IObject) EspressoImage2Image
}

// Init initializes the instance.
func (e EspressoImage2Image) Init() EspressoImage2Image {
	rv := objc.Send[EspressoImage2Image](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoImage2Image) Autorelease() EspressoImage2Image {
	rv := objc.Send[EspressoImage2Image](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoImage2Image creates a new EspressoImage2Image instance.
func NewEspressoImage2Image() EspressoImage2Image {
	class := getEspressoImage2ImageClass()
	rv := objc.Send[EspressoImage2Image](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/initWithQueue:
func NewEspressoImage2ImageWithQueue(queue objectivec.IObject) EspressoImage2Image {
	instance := getEspressoImage2ImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithQueue:"), queue)
	return EspressoImage2ImageFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/addNoiseLayer
func (e EspressoImage2Image) AddNoiseLayer() {
	objc.Send[objc.ID](e.ID, objc.Sel("addNoiseLayer"))
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/aggregateWisdom:
func (e EspressoImage2Image) AggregateWisdom(wisdom unsafe.Pointer) {
	objc.Send[objc.ID](e.ID, objc.Sel("aggregateWisdom:"), wisdom)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/benchmark
func (e EspressoImage2Image) Benchmark() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("benchmark"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:
func (e EspressoImage2Image) EncodeToCommandBufferSourceTextureDestinationTextureCropRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject) int {
	rv := objc.Send[int](e.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:"), buffer, texture, texture2, rect)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:destinationRect:
func (e EspressoImage2Image) EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject, rect2 objectivec.IObject) int {
	rv := objc.Send[int](e.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:destinationRect:"), buffer, texture, texture2, rect, rect2)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/getEspressoNetwork
func (e EspressoImage2Image) GetEspressoNetwork() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getEspressoNetwork"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/getInternalDataForKey:
func (e EspressoImage2Image) GetInternalDataForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getInternalDataForKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/height
func (e EspressoImage2Image) Height() int {
	rv := objc.Send[int](e.ID, objc.Sel("height"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/load:
func (e EspressoImage2Image) Load(load objectivec.IObject) int {
	rv := objc.Send[int](e.ID, objc.Sel("load:"), load)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/load:resolutionPreset:
func (e EspressoImage2Image) LoadResolutionPreset(load objectivec.IObject, preset int64) int {
	rv := objc.Send[int](e.ID, objc.Sel("load:resolutionPreset:"), load, preset)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/newOutputTexture
func (e EspressoImage2Image) NewOutputTexture() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("newOutputTexture"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/postProcess:cameraSourceTexture:inputTexture:destinationTexture:
func (e EspressoImage2Image) PostProcessCameraSourceTextureInputTextureDestinationTexture(process objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, texture3 objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("postProcess:cameraSourceTexture:inputTexture:destinationTexture:"), process, texture, texture2, texture3)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/reload
func (e EspressoImage2Image) Reload() int {
	rv := objc.Send[int](e.ID, objc.Sel("reload"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/resetTemporalState
func (e EspressoImage2Image) ResetTemporalState() {
	objc.Send[objc.ID](e.ID, objc.Sel("resetTemporalState"))
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/reshapeToResolutionPreset:
func (e EspressoImage2Image) ReshapeToResolutionPreset(preset int64) int {
	rv := objc.Send[int](e.ID, objc.Sel("reshapeToResolutionPreset:"), preset)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/reshapeToResolutionPreset:aspectRatio:
func (e EspressoImage2Image) ReshapeToResolutionPresetAspectRatio(preset int64, ratio float32) int {
	rv := objc.Send[int](e.ID, objc.Sel("reshapeToResolutionPreset:aspectRatio:"), preset, ratio)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/reshapeToWidth:andHeight:
func (e EspressoImage2Image) ReshapeToWidthAndHeight(width int, height int) int {
	rv := objc.Send[int](e.ID, objc.Sel("reshapeToWidth:andHeight:"), width, height)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/resolutionForPreset:
func (e EspressoImage2Image) ResolutionForPreset(preset int64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("resolutionForPreset:"), preset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/setupWithQueue:
func (e EspressoImage2Image) SetupWithQueue(queue objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("setupWithQueue:"), queue)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/simpleLinearResize:sourceTexture:destinationTexture:
func (e EspressoImage2Image) SimpleLinearResizeSourceTextureDestinationTexture(resize objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("simpleLinearResize:sourceTexture:destinationTexture:"), resize, texture, texture2)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/styleName
func (e EspressoImage2Image) StyleName() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("styleName"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/submitToQueueWithSourceTexture:destinationTexture:
func (e EspressoImage2Image) SubmitToQueueWithSourceTextureDestinationTexture(texture objectivec.IObject, texture2 objectivec.IObject) int {
	rv := objc.Send[int](e.ID, objc.Sel("submitToQueueWithSourceTexture:destinationTexture:"), texture, texture2)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/submitToQueueWithSourceTexture:destinationTexture:cropRect:
func (e EspressoImage2Image) SubmitToQueueWithSourceTextureDestinationTextureCropRect(texture objectivec.IObject, texture2 objectivec.IObject, rect objectivec.IObject) int {
	rv := objc.Send[int](e.ID, objc.Sel("submitToQueueWithSourceTexture:destinationTexture:cropRect:"), texture, texture2, rect)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/tune
func (e EspressoImage2Image) Tune() {
	objc.Send[objc.ID](e.ID, objc.Sel("tune"))
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/tweak:value:
func (e EspressoImage2Image) TweakValue(tweak objectivec.IObject, value float32) {
	objc.Send[objc.ID](e.ID, objc.Sel("tweak:value:"), tweak, value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/wasReshaped
func (e EspressoImage2Image) WasReshaped() int {
	rv := objc.Send[int](e.ID, objc.Sel("wasReshaped"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/width
func (e EspressoImage2Image) Width() int {
	rv := objc.Send[int](e.ID, objc.Sel("width"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/initWithQueue:
func (e EspressoImage2Image) InitWithQueue(queue objectivec.IObject) EspressoImage2Image {
	rv := objc.Send[EspressoImage2Image](e.ID, objc.Sel("initWithQueue:"), queue)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/featureVersion
func (_EspressoImage2ImageClass EspressoImage2ImageClass) FeatureVersion() int {
	rv := objc.Send[int](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("featureVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/getStylesKeys
func (_EspressoImage2ImageClass EspressoImage2ImageClass) GetStylesKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("getStylesKeys"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/gpuSync:tex:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) GpuSyncTex(sync objectivec.IObject, tex objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("gpuSync:tex:"), sync, tex)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/loadStylesConfigAtDefaultsKey:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) LoadStylesConfigAtDefaultsKey(key objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("loadStylesConfigAtDefaultsKey:"), key)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/loadStylesConfigAtPath:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) LoadStylesConfigAtPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("loadStylesConfigAtPath:"), path)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/setDefaultOption:toValue:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) SetDefaultOptionToValue(option objectivec.IObject, value objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("setDefaultOption:toValue:"), option, value)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/tuneNetworks:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) TuneNetworks(networks objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("tuneNetworks:"), networks)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/tuneNetworksWGWindowSize:
func (_EspressoImage2ImageClass EspressoImage2ImageClass) TuneNetworksWGWindowSize(size objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("tuneNetworksWGWindowSize:"), size)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/flip_y
func (e EspressoImage2Image) Flip_y() int {
	rv := objc.Send[int](e.ID, objc.Sel("flip_y"))
	return rv
}
func (e EspressoImage2Image) SetFlip_y(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setFlip_y:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoImage2Image/rotation_degrees
func (e EspressoImage2Image) Rotation_degrees() int {
	rv := objc.Send[int](e.ID, objc.Sel("rotation_degrees"))
	return rv
}
func (e EspressoImage2Image) SetRotation_degrees(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setRotation_degrees:"), value)
}

