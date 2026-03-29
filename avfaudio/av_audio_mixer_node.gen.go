// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioMixerNode] class.
var (
	_AVAudioMixerNodeClass     AVAudioMixerNodeClass
	_AVAudioMixerNodeClassOnce sync.Once
)

func getAVAudioMixerNodeClass() AVAudioMixerNodeClass {
	_AVAudioMixerNodeClassOnce.Do(func() {
		_AVAudioMixerNodeClass = AVAudioMixerNodeClass{class: objc.GetClass("AVAudioMixerNode")}
	})
	return _AVAudioMixerNodeClass
}

// GetAVAudioMixerNodeClass returns the class object for AVAudioMixerNode.
func GetAVAudioMixerNodeClass() AVAudioMixerNodeClass {
	return getAVAudioMixerNodeClass()
}

type AVAudioMixerNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioMixerNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioMixerNodeClass) Alloc() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that takes any number of inputs and converts them into a single
// output.
//
// # Overview
// 
// The mixer accepts input at any sample rate and efficiently combines sample
// rate conversions. It also accepts any channel count and correctly upmixes
// or downmixes to the output channel count.
//
// # Getting and Setting the Mixer Volume
//
//   - [AVAudioMixerNode.OutputVolume]: The mixer’s output volume.
//   - [AVAudioMixerNode.SetOutputVolume]
//
// # Getting an Input Bus
//
//   - [AVAudioMixerNode.NextAvailableInputBus]: An audio bus that isn’t in a connected state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode
type AVAudioMixerNode struct {
	AVAudioNode
}

// AVAudioMixerNodeFromID constructs a [AVAudioMixerNode] from an objc.ID.
//
// An object that takes any number of inputs and converts them into a single
// output.
func AVAudioMixerNodeFromID(id objc.ID) AVAudioMixerNode {
	return AVAudioMixerNode{AVAudioNode: AVAudioNodeFromID(id)}
}
// NOTE: AVAudioMixerNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioMixerNode] class.
//
// # Getting and Setting the Mixer Volume
//
//   - [IAVAudioMixerNode.OutputVolume]: The mixer’s output volume.
//   - [IAVAudioMixerNode.SetOutputVolume]
//
// # Getting an Input Bus
//
//   - [IAVAudioMixerNode.NextAvailableInputBus]: An audio bus that isn’t in a connected state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode
type IAVAudioMixerNode interface {
	IAVAudioNode
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Getting and Setting the Mixer Volume

	// The mixer’s output volume.
	OutputVolume() float32
	SetOutputVolume(value float32)

	// Topic: Getting an Input Bus

	// An audio bus that isn’t in a connected state.
	NextAvailableInputBus() AVAudioNodeBus
}

// Init initializes the instance.
func (a AVAudioMixerNode) Init() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioMixerNode) Autorelease() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMixerNode creates a new AVAudioMixerNode instance.
func NewAVAudioMixerNode() AVAudioMixerNode {
	class := getAVAudioMixerNodeClass()
	rv := objc.Send[AVAudioMixerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets the audio mixing destination object that corresponds to the specified
// mixer node and input bus.
//
// mixer: The mixer to get destination details for.
//
// bus: The input bus.
//
// # Return Value
// 
// Returns `self` if the specified mixer or input bus matches its connection
// point. If the mixer or input bus doesn’t match its connection point, or
// if the source node isn’t in a connected state to the mixer or input bus,
// the method returns `nil.`
//
// # Discussion
// 
// When you connect a source node to multiple mixers downstream, setting
// [AVAudioMixing] properties directly on the source node applies the change
// to all of them. Use this method to get the corresponding
// [AVAudioMixingDestination] for a specific mixer. Properties set on
// individual destination instances don’t reflect at the source node level.
// 
// If there’s any disconnection between the source and mixer nodes, the
// return value can be invalid. Fetch the return value every time you want to
// set or get properties on a specific mixer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/destination(forMixer:bus:)
func (a AVAudioMixerNode) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// The mixer’s output volume.
//
// # Discussion
// 
// The range of valid values is `0.0` to `1.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode/outputVolume
func (a AVAudioMixerNode) OutputVolume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("outputVolume"))
	return rv
}
func (a AVAudioMixerNode) SetOutputVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputVolume:"), value)
}
// An audio bus that isn’t in a connected state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode/nextAvailableInputBus
func (a AVAudioMixerNode) NextAvailableInputBus() AVAudioNodeBus {
	rv := objc.Send[AVAudioNodeBus](a.ID, objc.Sel("nextAvailableInputBus"))
	return AVAudioNodeBus(rv)
}
// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// # Discussion
// 
// The value of `obstruction` is in decibels. The system blocks only the
// direct path of sound between the source and listener.
// 
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (a AVAudioMixerNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioMixerNode) SetObstruction(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setObstruction:"), value)
}
// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// # Discussion
// 
// The value of `obstruction` is in decibels. The system blocks the direct and
// reverb paths of sound between the source and listener.
// 
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (a AVAudioMixerNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioMixerNode) SetOcclusion(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOcclusion:"), value)
}
// The bus’s stereo pan.
//
// # Discussion
// 
// The default value is `0.0`, and the range of valid values is `-1.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (a AVAudioMixerNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioMixerNode) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}
// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioMixerNode) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioMixerNode) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setPointSourceInHeadMode:"), value)
}
// The location of the source in the 3D environment.
//
// # Discussion
// 
// The system specifies the coordinates in meters. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (a AVAudioMixerNode) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioMixerNode) SetPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPosition:"), value)
}
// A value that changes the playback rate of the input signal.
//
// # Discussion
// 
// A value of `2.0` results in the output audio playing one octave higher. A
// value of `0.5` results in the output audio playing one octave lower.
// 
// The default value is `1.0`, and the range of valid values is `0.5` to
// `2.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (a AVAudioMixerNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioMixerNode) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}
// The type of rendering algorithm the mixer uses.
//
// # Discussion
// 
// Depending on the current output format of the [AVAudioEnvironmentNode]
// instance, the system may only support a subset of the rendering algorithms.
// You can retrieve an array of valid rendering algorithms by calling the
// [ApplicableRenderingAlgorithms] function of the [AVAudioEnvironmentNode]
// instance.
// 
// The default rendering algorithm is
// [Audio3DMixingRenderingAlgorithmEqualPowerPanning]. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (a AVAudioMixerNode) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioMixerNode) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenderingAlgorithm:"), value)
}
// A value that controls the blend of dry and reverb processed audio.
//
// # Discussion
// 
// This property controls the amount of the source’s audio that the
// [AVAudioEnvironmentNode] instance processes. A value of `0.5` results in an
// equal blend of dry and processed (wet) audio.
// 
// The default is `0.0`, and the range of valid values is `0.0` (completely
// dry) to `1.0` (completely wet). Only the [AVAudioEnvironmentNode] class
// implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (a AVAudioMixerNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioMixerNode) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}
// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioMixerNode) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioMixerNode) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setSourceMode:"), value)
}
// The bus’s input volume.
//
// # Discussion
// 
// The default value is `1.0`, and the range of valid values is `0.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] and the [AVAudioMixerNode]
// implement this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (a AVAudioMixerNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioMixerNode) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

			// Protocol methods for AVAudioMixing
			

