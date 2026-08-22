// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

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

// Class returns the underlying Objective-C class pointer.
func (ec EspressoImage2ImageClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoImage2ImageClass) Alloc() EspressoImage2Image {
	rv := objc.SendIfResponds[EspressoImage2Image](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

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
//   - [EspressoImage2Image._resetTemporalState]
//   - [EspressoImage2Image._reshapeToResolutionPreset]
//   - [EspressoImage2Image._reshapeToWidthAndHeight]
//   - [EspressoImage2Image._tune]
//   - [EspressoImage2Image.InitWithQueue]
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
//   - [IEspressoImage2Image._resetTemporalState]
//   - [IEspressoImage2Image._reshapeToResolutionPreset]
//   - [IEspressoImage2Image._reshapeToWidthAndHeight]
//   - [IEspressoImage2Image._tune]
//   - [IEspressoImage2Image.InitWithQueue]
type IEspressoImage2Image interface {
	objectivec.IObject

	// Topic: Methods

	AddNoiseLayer()
	AggregateWisdom(wisdom unsafe.Pointer)
	Benchmark() float32
	EncodeToCommandBufferSourceTextureDestinationTextureCropRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer) int
	EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer, rect2 unsafe.Pointer) int
	Flip_y() int
	SetFlip_y(value int)
	GetEspressoNetwork() unsafe.Pointer
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
	ResolutionForPreset(preset int64) unsafe.Pointer
	Rotation_degrees() int
	SetRotation_degrees(value int)
	SetupWithQueue(queue objectivec.IObject) objectivec.IObject
	SimpleLinearResizeSourceTextureDestinationTexture(resize objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject)
	StyleName() objectivec.IObject
	SubmitToQueueWithSourceTextureDestinationTexture(texture objectivec.IObject, texture2 objectivec.IObject) int
	SubmitToQueueWithSourceTextureDestinationTextureCropRect(texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer) int
	Tune()
	TweakValue(tweak objectivec.IObject, value float32)
	WasReshaped() int
	Width() int
	_resetTemporalState()
	_reshapeToResolutionPreset(preset int64) int
	_reshapeToWidthAndHeight(width int, height int) int
	_tune()
	InitWithQueue(queue objectivec.IObject) EspressoImage2Image
}

// Init initializes the instance.
func (e EspressoImage2Image) Init() EspressoImage2Image {
	rv := objc.SendIfResponds[EspressoImage2Image](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoImage2Image) Autorelease() EspressoImage2Image {
	rv := objc.SendIfResponds[EspressoImage2Image](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoImage2Image creates a new EspressoImage2Image instance.
func NewEspressoImage2Image() EspressoImage2Image {
	class := getEspressoImage2ImageClass()
	rv := objc.SendIfResponds[EspressoImage2Image](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoImage2ImageWithQueue(queue objectivec.IObject) EspressoImage2Image {
	instance := getEspressoImage2ImageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithQueue:"), queue)
	return EspressoImage2ImageFromID(rv)
}

func (e EspressoImage2Image) AddNoiseLayer() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("addNoiseLayer"))
}
func (e EspressoImage2Image) AggregateWisdom(wisdom unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("aggregateWisdom:"), wisdom)
}
func (e EspressoImage2Image) Benchmark() float32 {
	rv := objc.SendIfResponds[float32](e.ID, objc.Sel("benchmark"))
	return rv
}
func (e EspressoImage2Image) EncodeToCommandBufferSourceTextureDestinationTextureCropRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:"), buffer, texture, texture2, rect)
	return rv
}
func (e EspressoImage2Image) EncodeToCommandBufferSourceTextureDestinationTextureCropRectDestinationRect(buffer objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer, rect2 unsafe.Pointer) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:cropRect:destinationRect:"), buffer, texture, texture2, rect, rect2)
	return rv
}
func (e EspressoImage2Image) GetEspressoNetwork() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("getEspressoNetwork"))
	return rv
}
func (e EspressoImage2Image) GetInternalDataForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("getInternalDataForKey:"), key)
	return objectivec.Object{ID: rv}
}
func (e EspressoImage2Image) Height() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("height"))
	return rv
}
func (e EspressoImage2Image) Load(load objectivec.IObject) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("load:"), load)
	return rv
}
func (e EspressoImage2Image) LoadResolutionPreset(load objectivec.IObject, preset int64) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("load:resolutionPreset:"), load, preset)
	return rv
}
func (e EspressoImage2Image) NewOutputTexture() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("newOutputTexture"))
	return objectivec.Object{ID: rv}
}
func (e EspressoImage2Image) PostProcessCameraSourceTextureInputTextureDestinationTexture(process objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject, texture3 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("postProcess:cameraSourceTexture:inputTexture:destinationTexture:"), process, texture, texture2, texture3)
}
func (e EspressoImage2Image) Reload() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("reload"))
	return rv
}
func (e EspressoImage2Image) ResetTemporalState() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("resetTemporalState"))
}
func (e EspressoImage2Image) ReshapeToResolutionPreset(preset int64) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("reshapeToResolutionPreset:"), preset)
	return rv
}
func (e EspressoImage2Image) ReshapeToResolutionPresetAspectRatio(preset int64, ratio float32) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("reshapeToResolutionPreset:aspectRatio:"), preset, ratio)
	return rv
}
func (e EspressoImage2Image) ReshapeToWidthAndHeight(width int, height int) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("reshapeToWidth:andHeight:"), width, height)
	return rv
}
func (e EspressoImage2Image) ResolutionForPreset(preset int64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("resolutionForPreset:"), preset)
	return rv
}
func (e EspressoImage2Image) SetupWithQueue(queue objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("setupWithQueue:"), queue)
	return objectivec.Object{ID: rv}
}
func (e EspressoImage2Image) SimpleLinearResizeSourceTextureDestinationTexture(resize objectivec.IObject, texture objectivec.IObject, texture2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("simpleLinearResize:sourceTexture:destinationTexture:"), resize, texture, texture2)
}
func (e EspressoImage2Image) StyleName() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("styleName"))
	return objectivec.Object{ID: rv}
}
func (e EspressoImage2Image) SubmitToQueueWithSourceTextureDestinationTexture(texture objectivec.IObject, texture2 objectivec.IObject) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("submitToQueueWithSourceTexture:destinationTexture:"), texture, texture2)
	return rv
}
func (e EspressoImage2Image) SubmitToQueueWithSourceTextureDestinationTextureCropRect(texture objectivec.IObject, texture2 objectivec.IObject, rect unsafe.Pointer) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("submitToQueueWithSourceTexture:destinationTexture:cropRect:"), texture, texture2, rect)
	return rv
}
func (e EspressoImage2Image) Tune() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("tune"))
}
func (e EspressoImage2Image) TweakValue(tweak objectivec.IObject, value float32) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("tweak:value:"), tweak, value)
}
func (e EspressoImage2Image) WasReshaped() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("wasReshaped"))
	return rv
}
func (e EspressoImage2Image) Width() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("width"))
	return rv
}
func (e EspressoImage2Image) _resetTemporalState() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("_resetTemporalState"))
}
func (e EspressoImage2Image) _reshapeToResolutionPreset(preset int64) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("_reshapeToResolutionPreset:"), preset)
	return rv
}
func (e EspressoImage2Image) _reshapeToWidthAndHeight(width int, height int) int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("_reshapeToWidth:andHeight:"), width, height)
	return rv
}
func (e EspressoImage2Image) _tune() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("_tune"))
}
func (e EspressoImage2Image) InitWithQueue(queue objectivec.IObject) EspressoImage2Image {
	rv := objc.SendIfResponds[EspressoImage2Image](e.ID, objc.Sel("initWithQueue:"), queue)
	return rv
}

func (_EspressoImage2ImageClass EspressoImage2ImageClass) FeatureVersion() int {
	rv := objc.SendIfResponds[int](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("featureVersion"))
	return rv
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) GetStylesKeys() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("getStylesKeys"))
	return objectivec.Object{ID: rv}
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) GpuSyncTex(sync objectivec.IObject, tex objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("gpuSync:tex:"), sync, tex)
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) LoadStylesConfigAtDefaultsKey(key objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("loadStylesConfigAtDefaultsKey:"), key)
	return rv
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) LoadStylesConfigAtPath(path objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("loadStylesConfigAtPath:"), path)
	return rv
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) SetDefaultOptionToValue(option objectivec.IObject, value objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("setDefaultOption:toValue:"), option, value)
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) TuneNetworks(networks objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("tuneNetworks:"), networks)
}
func (_EspressoImage2ImageClass EspressoImage2ImageClass) TuneNetworksWGWindowSize(size objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_EspressoImage2ImageClass.class), objc.Sel("tuneNetworksWGWindowSize:"), size)
}

func (e EspressoImage2Image) Flip_y() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("flip_y"))
	return rv
}
func (e EspressoImage2Image) SetFlip_y(value int) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setFlip_y:"), value)
}
func (e EspressoImage2Image) Rotation_degrees() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("rotation_degrees"))
	return rv
}
func (e EspressoImage2Image) SetRotation_degrees(value int) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setRotation_degrees:"), value)
}
