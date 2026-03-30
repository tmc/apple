// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioInputNode] class.
var (
	_AVAudioInputNodeClass     AVAudioInputNodeClass
	_AVAudioInputNodeClassOnce sync.Once
)

func getAVAudioInputNodeClass() AVAudioInputNodeClass {
	_AVAudioInputNodeClassOnce.Do(func() {
		_AVAudioInputNodeClass = AVAudioInputNodeClass{class: objc.GetClass("AVAudioInputNode")}
	})
	return _AVAudioInputNodeClass
}

// GetAVAudioInputNodeClass returns the class object for AVAudioInputNode.
func GetAVAudioInputNodeClass() AVAudioInputNodeClass {
	return getAVAudioInputNodeClass()
}

type AVAudioInputNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioInputNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioInputNodeClass) Alloc() AVAudioInputNode {
	rv := objc.Send[AVAudioInputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that connects to the system’s audio input.
//
// # Overview
//
// This node connects to the system’s audio input when rendering to or from
// an audio device. In manual rendering mode, this node supplies input data to
// the engine.
//
// This audio node has one element. The format of the input scope reflects:
//
// - The audio hardware sample rate and channel count when it connects to
// hardware. - The format of the PCM audio data that the node supplies to the
// engine in manual rendering mode. For more information, see
// [AVAudioInputNode.SetManualRenderingInputPCMFormatInputBlock]
//
// When rendering from an audio device, the input node doesn’t support
// format conversion. In this case, the format of the output scope must be the
// same as the input and the formats for all nodes connected to the input
// chain.
//
// In manual rendering mode, the format of the output scope is initially the
// same as the input, but you may set it to a different format, which converts
// the node.
//
// # Manually Giving Data to an Audio Engine
//
//   - [AVAudioInputNode.SetManualRenderingInputPCMFormatInputBlock]: Supplies the data through the input node to the engine while operating in the manual rendering mode.
//
// # Getting and Setting Voice Processing Properties
//
//   - [AVAudioInputNode.VoiceProcessingInputMuted]: A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//   - [AVAudioInputNode.SetVoiceProcessingInputMuted]
//   - [AVAudioInputNode.VoiceProcessingBypassed]: A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//   - [AVAudioInputNode.SetVoiceProcessingBypassed]
//   - [AVAudioInputNode.VoiceProcessingAGCEnabled]: A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//   - [AVAudioInputNode.SetVoiceProcessingAGCEnabled]
//   - [AVAudioInputNode.VoiceProcessingOtherAudioDuckingConfiguration]: The ducking configuration of nonvoice audio.
//   - [AVAudioInputNode.SetVoiceProcessingOtherAudioDuckingConfiguration]
//
// # Handling Muted Speech Events
//
//   - [AVAudioInputNode.SetMutedSpeechActivityEventListener]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode
type AVAudioInputNode struct {
	AVAudioIONode
}

// AVAudioInputNodeFromID constructs a [AVAudioInputNode] from an objc.ID.
//
// An object that connects to the system’s audio input.
func AVAudioInputNodeFromID(id objc.ID) AVAudioInputNode {
	return AVAudioInputNode{AVAudioIONode: AVAudioIONodeFromID(id)}
}

// NOTE: AVAudioInputNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioInputNode] class.
//
// # Manually Giving Data to an Audio Engine
//
//   - [IAVAudioInputNode.SetManualRenderingInputPCMFormatInputBlock]: Supplies the data through the input node to the engine while operating in the manual rendering mode.
//
// # Getting and Setting Voice Processing Properties
//
//   - [IAVAudioInputNode.VoiceProcessingInputMuted]: A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//   - [IAVAudioInputNode.SetVoiceProcessingInputMuted]
//   - [IAVAudioInputNode.VoiceProcessingBypassed]: A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//   - [IAVAudioInputNode.SetVoiceProcessingBypassed]
//   - [IAVAudioInputNode.VoiceProcessingAGCEnabled]: A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//   - [IAVAudioInputNode.SetVoiceProcessingAGCEnabled]
//   - [IAVAudioInputNode.VoiceProcessingOtherAudioDuckingConfiguration]: The ducking configuration of nonvoice audio.
//   - [IAVAudioInputNode.SetVoiceProcessingOtherAudioDuckingConfiguration]
//
// # Handling Muted Speech Events
//
//   - [IAVAudioInputNode.SetMutedSpeechActivityEventListener]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode
type IAVAudioInputNode interface {
	IAVAudioIONode
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Manually Giving Data to an Audio Engine

	// Supplies the data through the input node to the engine while operating in the manual rendering mode.
	SetManualRenderingInputPCMFormatInputBlock(format IAVAudioFormat, block AVAudioIONodeInputBlock) bool

	// Topic: Getting and Setting Voice Processing Properties

	// A Boolean that indicates whether the input of the voice processing unit is in a muted state.
	VoiceProcessingInputMuted() bool
	SetVoiceProcessingInputMuted(value bool)
	// A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
	VoiceProcessingBypassed() bool
	SetVoiceProcessingBypassed(value bool)
	// A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
	VoiceProcessingAGCEnabled() bool
	SetVoiceProcessingAGCEnabled(value bool)
	// The ducking configuration of nonvoice audio.
	VoiceProcessingOtherAudioDuckingConfiguration() AVAudioVoiceProcessingOtherAudioDuckingConfiguration
	SetVoiceProcessingOtherAudioDuckingConfiguration(value AVAudioVoiceProcessingOtherAudioDuckingConfiguration)

	// Topic: Handling Muted Speech Events

	SetMutedSpeechActivityEventListener(listenerBlock AVAudioVoiceProcessingSpeechActivityEventHandler) bool
}

// Init initializes the instance.
func (a AVAudioInputNode) Init() AVAudioInputNode {
	rv := objc.Send[AVAudioInputNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioInputNode) Autorelease() AVAudioInputNode {
	rv := objc.Send[AVAudioInputNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioInputNode creates a new AVAudioInputNode instance.
func NewAVAudioInputNode() AVAudioInputNode {
	class := getAVAudioInputNodeClass()
	rv := objc.Send[AVAudioInputNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Supplies the data through the input node to the engine while operating in
// the manual rendering mode.
//
// format: The format of the PCM audio data the block supplies to the engine.
//
// block: The block the engine calls on the input node to get the audio to send to
// the output when operating in the manual rendering mode. For more
// information, see [AVAudioIONodeInputBlock].
//
// # Discussion
//
// The block must be non-`nil` when using an input node while the engine is
// operating in manual rendering mode. If you switch the engine to render to
// and from an audio device, it invalidates any previous block.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/setManualRenderingInputPCMFormat(_:inputBlock:)
func (a AVAudioInputNode) SetManualRenderingInputPCMFormatInputBlock(format IAVAudioFormat, block AVAudioIONodeInputBlock) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setManualRenderingInputPCMFormat:inputBlock:"), format, block)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/setMutedSpeechActivityEventListener(_:)
func (a AVAudioInputNode) SetMutedSpeechActivityEventListener(listenerBlock AVAudioVoiceProcessingSpeechActivityEventHandler) bool {
	_block0, _ := NewAVAudioVoiceProcessingSpeechActivityEventBlock(listenerBlock)
	rv := objc.Send[bool](a.ID, objc.Sel("setMutedSpeechActivityEventListener:"), _block0)
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
func (a AVAudioInputNode) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// A Boolean that indicates whether the input of the voice processing unit is
// in a muted state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingInputMuted
func (a AVAudioInputNode) VoiceProcessingInputMuted() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isVoiceProcessingInputMuted"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingInputMuted(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setVoiceProcessingInputMuted:"), value)
}

// A Boolean that indicates whether the node bypasses all microphone uplink
// processing of the voice-processing unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingBypassed
func (a AVAudioInputNode) VoiceProcessingBypassed() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isVoiceProcessingBypassed"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingBypassed(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setVoiceProcessingBypassed:"), value)
}

// A Boolean that indicates whether automatic gain control on the processed
// microphone uplink signal is active.
//
// # Discussion
//
// This property is in an enabled state by default.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingAGCEnabled
func (a AVAudioInputNode) VoiceProcessingAGCEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isVoiceProcessingAGCEnabled"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingAGCEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setVoiceProcessingAGCEnabled:"), value)
}

// The ducking configuration of nonvoice audio.
//
// # Discussion
//
// Use this property to configures the ducking of nonvoice audio, including
// advanced enablement and ducking level. Typically, when playing other audio
// during voice chat, applying a higher level of ducking can increase the
// intelligibility of the voice chat.
//
// If not set, the default behavior is to disable advanced ducking, with a
// ducking level set to [AVAudioVoiceProcessingOtherAudioDuckingLevelDefault].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/voiceProcessingOtherAudioDuckingConfiguration
func (a AVAudioInputNode) VoiceProcessingOtherAudioDuckingConfiguration() AVAudioVoiceProcessingOtherAudioDuckingConfiguration {
	rv := objc.Send[AVAudioVoiceProcessingOtherAudioDuckingConfiguration](a.ID, objc.Sel("voiceProcessingOtherAudioDuckingConfiguration"))
	return AVAudioVoiceProcessingOtherAudioDuckingConfiguration(rv)
}
func (a AVAudioInputNode) SetVoiceProcessingOtherAudioDuckingConfiguration(value AVAudioVoiceProcessingOtherAudioDuckingConfiguration) {
	objc.Send[struct{}](a.ID, objc.Sel("setVoiceProcessingOtherAudioDuckingConfiguration:"), value)
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
func (a AVAudioInputNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioInputNode) SetObstruction(value float32) {
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
func (a AVAudioInputNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioInputNode) SetOcclusion(value float32) {
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
func (a AVAudioInputNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioInputNode) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioInputNode) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioInputNode) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
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
func (a AVAudioInputNode) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioInputNode) SetPosition(value AVAudio3DPoint) {
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
func (a AVAudioInputNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioInputNode) SetRate(value float32) {
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
// [AVAudio3DMixingRenderingAlgorithmEqualPowerPanning]. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (a AVAudioInputNode) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioInputNode) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
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
func (a AVAudioInputNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioInputNode) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioInputNode) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioInputNode) SetSourceMode(value AVAudio3DMixingSourceMode) {
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
func (a AVAudioInputNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioInputNode) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

// Protocol methods for AVAudioMixing

// SetMutedSpeechActivityEventListenerSync is a synchronous wrapper around [AVAudioInputNode.SetMutedSpeechActivityEventListener].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioInputNode) SetMutedSpeechActivityEventListenerSync(ctx context.Context) (AVAudioVoiceProcessingSpeechActivityEvent, error) {
	done := make(chan AVAudioVoiceProcessingSpeechActivityEvent, 1)
	a.SetMutedSpeechActivityEventListener(func(val AVAudioVoiceProcessingSpeechActivityEvent) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
