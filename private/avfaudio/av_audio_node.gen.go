// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioNode] class.
var (
	_AVAudioNodeClass     AVAudioNodeClass
	_AVAudioNodeClassOnce sync.Once
)

func getAVAudioNodeClass() AVAudioNodeClass {
	_AVAudioNodeClassOnce.Do(func() {
		_AVAudioNodeClass = AVAudioNodeClass{class: objc.GetClass("AVAudioNode")}
	})
	return _AVAudioNodeClass
}

// GetAVAudioNodeClass returns the class object for AVAudioNode.
func GetAVAudioNodeClass() AVAudioNodeClass {
	return getAVAudioNodeClass()
}

type AVAudioNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioNodeClass) Alloc() AVAudioNode {
	rv := objc.Send[AVAudioNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioNode.Clock]
//   - [AVAudioNode.DestinationForMixerBus]
//   - [AVAudioNode.DidAttachToEngine]
//   - [AVAudioNode.DidDetachFromEngineError]
//   - [AVAudioNode.Impl]
//   - [AVAudioNode.Obstruction]
//   - [AVAudioNode.Occlusion]
//   - [AVAudioNode.Pan]
//   - [AVAudioNode.PointSourceInHeadMode]
//   - [AVAudioNode.Rate]
//   - [AVAudioNode.RenderingAlgorithm]
//   - [AVAudioNode.ResetImpl]
//   - [AVAudioNode.ReverbBlend]
//   - [AVAudioNode.SetInputFormatForBus]
//   - [AVAudioNode.SetNumberOfInputs]
//   - [AVAudioNode.SetNumberOfOutputs]
//   - [AVAudioNode.SetObstruction]
//   - [AVAudioNode.SetOcclusion]
//   - [AVAudioNode.SetOutputFormatForBus]
//   - [AVAudioNode.SetPan]
//   - [AVAudioNode.SetPointSourceInHeadMode]
//   - [AVAudioNode.SetRate]
//   - [AVAudioNode.SetRenderingAlgorithm]
//   - [AVAudioNode.SetReverbBlend]
//   - [AVAudioNode.SetSourceMode]
//   - [AVAudioNode.SetVolume]
//   - [AVAudioNode.SourceMode]
//   - [AVAudioNode.Volume]
//   - [AVAudioNode.InitWithImpl]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
type AVAudioNode struct {
	objectivec.Object
}

// AVAudioNodeFromID constructs a [AVAudioNode] from an objc.ID.
func AVAudioNodeFromID(id objc.ID) AVAudioNode {
	return AVAudioNode{objectivec.Object{ID: id}}
}

// Ensure AVAudioNode implements IAVAudioNode.
var _ IAVAudioNode = AVAudioNode{}

// An interface definition for the [AVAudioNode] class.
//
// # Methods
//
//   - [IAVAudioNode.Clock]
//   - [IAVAudioNode.DestinationForMixerBus]
//   - [IAVAudioNode.DidAttachToEngine]
//   - [IAVAudioNode.DidDetachFromEngineError]
//   - [IAVAudioNode.Impl]
//   - [IAVAudioNode.Obstruction]
//   - [IAVAudioNode.Occlusion]
//   - [IAVAudioNode.Pan]
//   - [IAVAudioNode.PointSourceInHeadMode]
//   - [IAVAudioNode.Rate]
//   - [IAVAudioNode.RenderingAlgorithm]
//   - [IAVAudioNode.ResetImpl]
//   - [IAVAudioNode.ReverbBlend]
//   - [IAVAudioNode.SetInputFormatForBus]
//   - [IAVAudioNode.SetNumberOfInputs]
//   - [IAVAudioNode.SetNumberOfOutputs]
//   - [IAVAudioNode.SetObstruction]
//   - [IAVAudioNode.SetOcclusion]
//   - [IAVAudioNode.SetOutputFormatForBus]
//   - [IAVAudioNode.SetPan]
//   - [IAVAudioNode.SetPointSourceInHeadMode]
//   - [IAVAudioNode.SetRate]
//   - [IAVAudioNode.SetRenderingAlgorithm]
//   - [IAVAudioNode.SetReverbBlend]
//   - [IAVAudioNode.SetSourceMode]
//   - [IAVAudioNode.SetVolume]
//   - [IAVAudioNode.SourceMode]
//   - [IAVAudioNode.Volume]
//   - [IAVAudioNode.InitWithImpl]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
type IAVAudioNode interface {
	objectivec.IObject

	// Topic: Methods

	Clock() objectivec.IObject
	DestinationForMixerBus(mixer objectivec.IObject, bus uint64) objectivec.IObject
	DidAttachToEngine(engine objectivec.IObject)
	DidDetachFromEngineError(engine objectivec.IObject) error
	Impl() unsafe.Pointer
	Obstruction() float32
	Occlusion() float32
	Pan() float32
	PointSourceInHeadMode() int64
	Rate() float32
	RenderingAlgorithm() int64
	ResetImpl(impl unsafe.Pointer) bool
	ReverbBlend() float32
	SetInputFormatForBus(format objectivec.IObject, bus uint64) bool
	SetNumberOfInputs(inputs uint32)
	SetNumberOfOutputs(outputs uint32)
	SetObstruction(obstruction float32)
	SetOcclusion(occlusion float32)
	SetOutputFormatForBus(format objectivec.IObject, bus uint64) bool
	SetPan(pan float32)
	SetPointSourceInHeadMode(mode int64)
	SetRate(rate float32)
	SetRenderingAlgorithm(algorithm int64)
	SetReverbBlend(blend float32)
	SetSourceMode(mode int64)
	SetVolume(volume float32)
	SourceMode() int64
	Volume() float32
	InitWithImpl(impl unsafe.Pointer) AVAudioNode
}

// Init initializes the instance.
func (a AVAudioNode) Init() AVAudioNode {
	rv := objc.Send[AVAudioNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioNode) Autorelease() AVAudioNode {
	rv := objc.Send[AVAudioNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioNode creates a new AVAudioNode instance.
func NewAVAudioNode() AVAudioNode {
	class := getAVAudioNodeClass()
	rv := objc.Send[AVAudioNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioNodeWithImpl(impl unsafe.Pointer) AVAudioNode {
	instance := getAVAudioNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/clock
func (a AVAudioNode) Clock() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("clock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/destinationForMixer:bus:
func (a AVAudioNode) DestinationForMixerBus(mixer objectivec.IObject, bus uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/didAttachToEngine:
func (a AVAudioNode) DidAttachToEngine(engine objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("didAttachToEngine:"), engine)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/didDetachFromEngine:error:
func (a AVAudioNode) DidDetachFromEngineError(engine objectivec.IObject) error {
	var errorPtr objc.ID
	objc.Send[struct{}](a.ID, objc.Sel("didDetachFromEngine:error:"), engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/obstruction
func (a AVAudioNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/occlusion
func (a AVAudioNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/pan
func (a AVAudioNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/pointSourceInHeadMode
func (a AVAudioNode) PointSourceInHeadMode() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("pointSourceInHeadMode"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/rate
func (a AVAudioNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/renderingAlgorithm
func (a AVAudioNode) RenderingAlgorithm() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("renderingAlgorithm"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/resetImpl:
func (a AVAudioNode) ResetImpl(impl unsafe.Pointer) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("resetImpl:"), impl)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/reverbBlend
func (a AVAudioNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setInputFormat:forBus:
func (a AVAudioNode) SetInputFormatForBus(format objectivec.IObject, bus uint64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setInputFormat:forBus:"), format, bus)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setNumberOfInputs:
func (a AVAudioNode) SetNumberOfInputs(inputs uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setNumberOfInputs:"), inputs)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setNumberOfOutputs:
func (a AVAudioNode) SetNumberOfOutputs(outputs uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setNumberOfOutputs:"), outputs)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setObstruction:
func (a AVAudioNode) SetObstruction(obstruction float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setObstruction:"), obstruction)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setOcclusion:
func (a AVAudioNode) SetOcclusion(occlusion float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setOcclusion:"), occlusion)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setOutputFormat:forBus:
func (a AVAudioNode) SetOutputFormatForBus(format objectivec.IObject, bus uint64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setOutputFormat:forBus:"), format, bus)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setPan:
func (a AVAudioNode) SetPan(pan float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setPan:"), pan)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setPointSourceInHeadMode:
func (a AVAudioNode) SetPointSourceInHeadMode(mode int64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setPointSourceInHeadMode:"), mode)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setRate:
func (a AVAudioNode) SetRate(rate float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setRate:"), rate)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setRenderingAlgorithm:
func (a AVAudioNode) SetRenderingAlgorithm(algorithm int64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setRenderingAlgorithm:"), algorithm)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setReverbBlend:
func (a AVAudioNode) SetReverbBlend(blend float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setReverbBlend:"), blend)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setSourceMode:
func (a AVAudioNode) SetSourceMode(mode int64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSourceMode:"), mode)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/setVolume:
func (a AVAudioNode) SetVolume(volume float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setVolume:"), volume)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/sourceMode
func (a AVAudioNode) SourceMode() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("sourceMode"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/volume
func (a AVAudioNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func (a AVAudioNode) InitWithImpl(impl unsafe.Pointer) AVAudioNode {
	rv := objc.Send[AVAudioNode](a.ID, objc.Sel("initWithImpl:"), impl)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/impl
func (a AVAudioNode) Impl() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("impl"))
	return rv
}
