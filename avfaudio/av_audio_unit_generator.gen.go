// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/audiotoolbox"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitGenerator] class.
var (
	_AVAudioUnitGeneratorClass     AVAudioUnitGeneratorClass
	_AVAudioUnitGeneratorClassOnce sync.Once
)

func getAVAudioUnitGeneratorClass() AVAudioUnitGeneratorClass {
	_AVAudioUnitGeneratorClassOnce.Do(func() {
		_AVAudioUnitGeneratorClass = AVAudioUnitGeneratorClass{class: objc.GetClass("AVAudioUnitGenerator")}
	})
	return _AVAudioUnitGeneratorClass
}

// GetAVAudioUnitGeneratorClass returns the class object for AVAudioUnitGenerator.
func GetAVAudioUnitGeneratorClass() AVAudioUnitGeneratorClass {
	return getAVAudioUnitGeneratorClass()
}

type AVAudioUnitGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitGeneratorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitGeneratorClass) Alloc() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that generates audio output.
//
// # Overview
//
// A generator represents an [AudioUnit] of type `kAudioUnitType_Generator` or
// `kAudioUnitType_RemoteGenerator`. A generator has no audio input, but
// produces audio output. An example is a tone generator.
//
// # Creating an audio unit generator
//
//   - [AVAudioUnitGenerator.InitWithAudioComponentDescription]: Creates a generator audio unit with the specified description.
//
// # Getting and setting the bypass status
//
//   - [AVAudioUnitGenerator.Bypass]: The bypass state of the audio unit.
//   - [AVAudioUnitGenerator.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
//
// [AudioUnit]: https://developer.apple.com/documentation/AudioToolbox/AudioUnit
type AVAudioUnitGenerator struct {
	AVAudioUnit
}

// AVAudioUnitGeneratorFromID constructs a [AVAudioUnitGenerator] from an objc.ID.
//
// An object that generates audio output.
func AVAudioUnitGeneratorFromID(id objc.ID) AVAudioUnitGenerator {
	return AVAudioUnitGenerator{AVAudioUnit: AVAudioUnitFromID(id)}
}

// NOTE: AVAudioUnitGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitGenerator] class.
//
// # Creating an audio unit generator
//
//   - [IAVAudioUnitGenerator.InitWithAudioComponentDescription]: Creates a generator audio unit with the specified description.
//
// # Getting and setting the bypass status
//
//   - [IAVAudioUnitGenerator.Bypass]: The bypass state of the audio unit.
//   - [IAVAudioUnitGenerator.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
type IAVAudioUnitGenerator interface {
	IAVAudioUnit

	// Topic: Creating an audio unit generator

	// Creates a generator audio unit with the specified description.
	InitWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AVAudioUnitGenerator

	// Topic: Getting and setting the bypass status

	// The bypass state of the audio unit.
	Bypass() bool
	SetBypass(value bool)

	// Gets the audio mixing destination object that corresponds to the specified mixer node and input bus.
	DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination
	// A value that simulates filtering of the direct path of sound due to an obstacle.
	Obstruction() float32
	// A value that simulates filtering of the direct and reverb paths of sound due to an obstacle.
	Occlusion() float32
	// The bus’s stereo pan.
	Pan() float32
	// The in-head mode for a point source.
	PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode
	// The location of the source in the 3D environment.
	Position() AVAudio3DPoint
	// A value that changes the playback rate of the input signal.
	Rate() float32
	// The type of rendering algorithm the mixer uses.
	RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm
	// A value that controls the blend of dry and reverb processed audio.
	ReverbBlend() float32
	// The source mode for the input bus of the audio environment node.
	SourceMode() AVAudio3DMixingSourceMode
	// The bus’s input volume.
	Volume() float32
}

// Init initializes the instance.
func (a AVAudioUnitGenerator) Init() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitGenerator) Autorelease() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitGenerator creates a new AVAudioUnitGenerator instance.
func NewAVAudioUnitGenerator() AVAudioUnitGenerator {
	class := getAVAudioUnitGeneratorClass()
	rv := objc.Send[AVAudioUnitGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a generator audio unit with the specified description.
//
// audioComponentDescription: The audio component description.
//
// # Return Value
//
// A new [AVAudioUnitGenerator] instance.
//
// # Discussion
//
// The [AudioComponentDescription] structure `componentType` field must be
// `kAudioUnitType_Generator` or [kAudioUnitType_RemoteGenerator].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/init(audioComponentDescription:)
//
// [AudioComponentDescription]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
// [kAudioUnitType_RemoteGenerator]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_RemoteGenerator
func NewAudioUnitGeneratorWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AVAudioUnitGenerator {
	instance := getAVAudioUnitGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitGeneratorFromID(rv)
}

// Creates a generator audio unit with the specified description.
//
// audioComponentDescription: The audio component description.
//
// # Return Value
//
// A new [AVAudioUnitGenerator] instance.
//
// # Discussion
//
// The [AudioComponentDescription] structure `componentType` field must be
// `kAudioUnitType_Generator` or [kAudioUnitType_RemoteGenerator].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/init(audioComponentDescription:)
//
// [AudioComponentDescription]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
// [kAudioUnitType_RemoteGenerator]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_RemoteGenerator
func (a AVAudioUnitGenerator) InitWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](a.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
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
func (a AVAudioUnitGenerator) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (a AVAudioUnitGenerator) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}

// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (a AVAudioUnitGenerator) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}

// The bus’s stereo pan.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (a AVAudioUnitGenerator) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioUnitGenerator) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}

// The location of the source in the 3D environment.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (a AVAudioUnitGenerator) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}

// A value that changes the playback rate of the input signal.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (a AVAudioUnitGenerator) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}

// The type of rendering algorithm the mixer uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (a AVAudioUnitGenerator) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}

// A value that controls the blend of dry and reverb processed audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (a AVAudioUnitGenerator) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioUnitGenerator) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}

// The bus’s input volume.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (a AVAudioUnitGenerator) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}

// The bypass state of the audio unit.
//
// # Discussion
//
// If true, the audio unit bypasses audio processing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/bypass
func (a AVAudioUnitGenerator) Bypass() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("bypass"))
	return rv
}
func (a AVAudioUnitGenerator) SetBypass(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setBypass:"), value)
}

// Protocol methods for AVAudioMixing

// The bus’s input volume.
//
// # Discussion
//
// The default value is `1.0`, and the range of valid values is `0.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] and the [AVAudioMixerNode]
// implement this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (o AVAudioUnitGenerator) SetVolume(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setVolume:"), value)
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
func (o AVAudioUnitGenerator) SetObstruction(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setObstruction:"), value)
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
func (o AVAudioUnitGenerator) SetOcclusion(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setOcclusion:"), value)
}

// The location of the source in the 3D environment.
//
// # Discussion
//
// The system specifies the coordinates in meters. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (o AVAudioUnitGenerator) SetPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPosition:"), value)
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
func (o AVAudioUnitGenerator) SetRate(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRate:"), value)
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (o AVAudioUnitGenerator) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setPointSourceInHeadMode:"), value)
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
func (o AVAudioUnitGenerator) SetReverbBlend(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setReverbBlend:"), value)
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (o AVAudioUnitGenerator) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceMode:"), value)
}

// The type of rendering algorithm the mixer uses.
//
// # Discussion
//
// Depending on the current output format of the [AVAudioEnvironmentNode]
// instance, the system may only support a subset of the rendering algorithms.
// You can retrieve an array of valid rendering algorithms by calling the
// [AVAudioEnvironmentNode.ApplicableRenderingAlgorithms] function of the
// [AVAudioEnvironmentNode] instance.
//
// The default rendering algorithm is
// [AVAudio3DMixingRenderingAlgorithmEqualPowerPanning]. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (o AVAudioUnitGenerator) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](o.ID, objc.Sel("setRenderingAlgorithm:"), value)
}

// The bus’s stereo pan.
//
// # Discussion
//
// The default value is `0.0`, and the range of valid values is `-1.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (o AVAudioUnitGenerator) SetPan(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPan:"), value)
}
