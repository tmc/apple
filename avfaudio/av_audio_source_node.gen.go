// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioSourceNode] class.
var (
	_AVAudioSourceNodeClass     AVAudioSourceNodeClass
	_AVAudioSourceNodeClassOnce sync.Once
)

func getAVAudioSourceNodeClass() AVAudioSourceNodeClass {
	_AVAudioSourceNodeClassOnce.Do(func() {
		_AVAudioSourceNodeClass = AVAudioSourceNodeClass{class: objc.GetClass("AVAudioSourceNode")}
	})
	return _AVAudioSourceNodeClass
}

// GetAVAudioSourceNodeClass returns the class object for AVAudioSourceNode.
func GetAVAudioSourceNodeClass() AVAudioSourceNodeClass {
	return getAVAudioSourceNodeClass()
}

type AVAudioSourceNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSourceNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSourceNodeClass) Alloc() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that supplies audio data.
//
// # Overview
// 
// The [AVAudioSourceNode] class allows for supplying audio data for rendering
// through [AVAudioSourceNodeRenderBlock]. It’s a convenient method for
// delievering audio data instead of setting the input callback on an audio
// unit with `kAudioUnitProperty_SetRenderCallback`.
//
// # Creating an Audio Source Node
//
//   - [AVAudioSourceNode.InitWithRenderBlock]: Creates an audio source node with a block that supplies audio data.
//   - [AVAudioSourceNode.InitWithFormatRenderBlock]: Creates an audio source node with the audio format and a block that supplies audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode
type AVAudioSourceNode struct {
	AVAudioNode
}

// AVAudioSourceNodeFromID constructs a [AVAudioSourceNode] from an objc.ID.
//
// An object that supplies audio data.
func AVAudioSourceNodeFromID(id objc.ID) AVAudioSourceNode {
	return AVAudioSourceNode{AVAudioNode: AVAudioNodeFromID(id)}
}
// NOTE: AVAudioSourceNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioSourceNode] class.
//
// # Creating an Audio Source Node
//
//   - [IAVAudioSourceNode.InitWithRenderBlock]: Creates an audio source node with a block that supplies audio data.
//   - [IAVAudioSourceNode.InitWithFormatRenderBlock]: Creates an audio source node with the audio format and a block that supplies audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode
type IAVAudioSourceNode interface {
	IAVAudioNode
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Creating an Audio Source Node

	// Creates an audio source node with a block that supplies audio data.
	InitWithRenderBlock(block AVAudioSourceNodeRenderBlock) AVAudioSourceNode
	// Creates an audio source node with the audio format and a block that supplies audio data.
	InitWithFormatRenderBlock(format IAVAudioFormat, block AVAudioSourceNodeRenderBlock) AVAudioSourceNode
}

// Init initializes the instance.
func (a AVAudioSourceNode) Init() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSourceNode) Autorelease() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSourceNode creates a new AVAudioSourceNode instance.
func NewAVAudioSourceNode() AVAudioSourceNode {
	class := getAVAudioSourceNodeClass()
	rv := objc.Send[AVAudioSourceNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio source node with the audio format and a block that
// supplies audio data.
//
// format: The format of the pulse-code modulated (PCM) audio data the block supplies.
//
// block: The block to supply audio data to the output.
//
// # Discussion
// 
// When connecting the audio source node to another node, the system uses the
// connection format to set the audio format for the output bus.
// 
// Depending on the audio engine’s operating mode, call the block on
// real-time or nonreal-time threads. When rendering to a device, avoid making
// blocking calls within the block.
// 
// [AVAudioSourceNode] supports different audio formats for the block and the
// output, but the system only supports linear PCM conversions with sample
// rate, bit depth, and interleaving.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(format:renderBlock:)
func NewAudioSourceNodeWithFormatRenderBlock(format IAVAudioFormat, block AVAudioSourceNodeRenderBlock) AVAudioSourceNode {
	instance := getAVAudioSourceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:renderBlock:"), format, block)
	return AVAudioSourceNodeFromID(rv)
}

// Creates an audio source node with a block that supplies audio data.
//
// block: The block to supply audio data to the output.
//
// # Discussion
// 
// When connecting the audio source node to another node, the system uses the
// connection format to set the audio format for the output bus.
// 
// Depending on the audio engine’s operating mode, call the block on
// real-time or nonreal-time threads. When rendering to a device, avoid making
// blocking calls within the block.
// 
// The system sets the node’s output format using the audio format for the
// render block. When reconnecting the node with a different output format,
// the audio format for the block changes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(renderBlock:)
func NewAudioSourceNodeWithRenderBlock(block AVAudioSourceNodeRenderBlock) AVAudioSourceNode {
	instance := getAVAudioSourceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRenderBlock:"), block)
	return AVAudioSourceNodeFromID(rv)
}

// Creates an audio source node with a block that supplies audio data.
//
// block: The block to supply audio data to the output.
//
// # Discussion
// 
// When connecting the audio source node to another node, the system uses the
// connection format to set the audio format for the output bus.
// 
// Depending on the audio engine’s operating mode, call the block on
// real-time or nonreal-time threads. When rendering to a device, avoid making
// blocking calls within the block.
// 
// The system sets the node’s output format using the audio format for the
// render block. When reconnecting the node with a different output format,
// the audio format for the block changes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(renderBlock:)
func (a AVAudioSourceNode) InitWithRenderBlock(block AVAudioSourceNodeRenderBlock) AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("initWithRenderBlock:"), block)
	return rv
}
// Creates an audio source node with the audio format and a block that
// supplies audio data.
//
// format: The format of the pulse-code modulated (PCM) audio data the block supplies.
//
// block: The block to supply audio data to the output.
//
// # Discussion
// 
// When connecting the audio source node to another node, the system uses the
// connection format to set the audio format for the output bus.
// 
// Depending on the audio engine’s operating mode, call the block on
// real-time or nonreal-time threads. When rendering to a device, avoid making
// blocking calls within the block.
// 
// [AVAudioSourceNode] supports different audio formats for the block and the
// output, but the system only supports linear PCM conversions with sample
// rate, bit depth, and interleaving.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(format:renderBlock:)
func (a AVAudioSourceNode) InitWithFormatRenderBlock(format IAVAudioFormat, block AVAudioSourceNodeRenderBlock) AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("initWithFormat:renderBlock:"), format, block)
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
func (a AVAudioSourceNode) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
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
func (a AVAudioSourceNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioSourceNode) SetObstruction(value float32) {
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
func (a AVAudioSourceNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioSourceNode) SetOcclusion(value float32) {
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
func (a AVAudioSourceNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioSourceNode) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}
// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioSourceNode) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioSourceNode) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
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
func (a AVAudioSourceNode) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioSourceNode) SetPosition(value AVAudio3DPoint) {
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
func (a AVAudioSourceNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioSourceNode) SetRate(value float32) {
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
func (a AVAudioSourceNode) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioSourceNode) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
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
func (a AVAudioSourceNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioSourceNode) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}
// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioSourceNode) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioSourceNode) SetSourceMode(value AVAudio3DMixingSourceMode) {
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
func (a AVAudioSourceNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioSourceNode) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

			// Protocol methods for AVAudioMixing
			

