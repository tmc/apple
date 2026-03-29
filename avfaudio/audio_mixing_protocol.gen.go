// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A collection of properties that are applicable to the input bus of a mixer node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing
type AVAudioMixing interface {
	objectivec.IObject
	AVAudio3DMixing
	AVAudioStereoMixing

	// Gets the audio mixing destination object that corresponds to the specified mixer node and input bus.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/destination(forMixer:bus:)
	DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination

	// The bus’s input volume.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
	Volume() float32

	// The bus’s input volume.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
	SetVolume(value float32)
}

// AVAudioMixingObject wraps an existing Objective-C object that conforms to the AVAudioMixing protocol.
type AVAudioMixingObject struct {
	objectivec.Object
}
func (o AVAudioMixingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioMixingObjectFromID constructs a [AVAudioMixingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioMixingObjectFromID(id objc.ID) AVAudioMixingObject {
	return AVAudioMixingObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o AVAudioMixingObject) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
	}
// The bus’s input volume.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (o AVAudioMixingObject) Volume() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("volume"))
	return rv
	}
// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (o AVAudioMixingObject) Obstruction() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("obstruction"))
	return rv
	}
// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (o AVAudioMixingObject) Occlusion() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("occlusion"))
	return rv
	}
// The location of the source in the 3D environment.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (o AVAudioMixingObject) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](o.ID, objc.Sel("position"))
	return rv
	}
// A value that changes the playback rate of the input signal.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (o AVAudioMixingObject) Rate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("rate"))
	return rv
	}
// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (o AVAudioMixingObject) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](o.ID, objc.Sel("pointSourceInHeadMode"))
	return rv
	}
// A value that controls the blend of dry and reverb processed audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (o AVAudioMixingObject) ReverbBlend() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("reverbBlend"))
	return rv
	}
// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (o AVAudioMixingObject) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](o.ID, objc.Sel("sourceMode"))
	return rv
	}
// The type of rendering algorithm the mixer uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (o AVAudioMixingObject) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](o.ID, objc.Sel("renderingAlgorithm"))
	return rv
	}
// The bus’s stereo pan.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (o AVAudioMixingObject) Pan() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("pan"))
	return rv
	}

func (o AVAudioMixingObject) SetVolume(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setVolume:"), value)
}

func (o AVAudioMixingObject) SetObstruction(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setObstruction:"), value)
}

func (o AVAudioMixingObject) SetOcclusion(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setOcclusion:"), value)
}

func (o AVAudioMixingObject) SetPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPosition:"), value)
}

func (o AVAudioMixingObject) SetRate(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRate:"), value)
}

func (o AVAudioMixingObject) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setPointSourceInHeadMode:"), value)
}

func (o AVAudioMixingObject) SetReverbBlend(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setReverbBlend:"), value)
}

func (o AVAudioMixingObject) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceMode:"), value)
}

func (o AVAudioMixingObject) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](o.ID, objc.Sel("setRenderingAlgorithm:"), value)
}

func (o AVAudioMixingObject) SetPan(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPan:"), value)
}

