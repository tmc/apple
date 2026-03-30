// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A collection of properties that define 3D mixing properties.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing
type AVAudio3DMixing interface {
	objectivec.IObject

	// A value that simulates filtering of the direct path of sound due to an obstacle.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
	Obstruction() float32

	// A value that simulates filtering of the direct and reverb paths of sound due to an obstacle.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
	Occlusion() float32

	// The location of the source in the 3D environment.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
	Position() AVAudio3DPoint

	// A value that changes the playback rate of the input signal.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
	Rate() float32

	// The in-head mode for a point source.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
	PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode

	// A value that controls the blend of dry and reverb processed audio.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
	ReverbBlend() float32

	// The source mode for the input bus of the audio environment node.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
	SourceMode() AVAudio3DMixingSourceMode

	// The type of rendering algorithm the mixer uses.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
	RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm

	// A value that simulates filtering of the direct path of sound due to an obstacle.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
	SetObstruction(value float32)

	// A value that simulates filtering of the direct and reverb paths of sound due to an obstacle.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
	SetOcclusion(value float32)

	// The location of the source in the 3D environment.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
	SetPosition(value AVAudio3DPoint)

	// A value that changes the playback rate of the input signal.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
	SetRate(value float32)

	// The in-head mode for a point source.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
	SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode)

	// A value that controls the blend of dry and reverb processed audio.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
	SetReverbBlend(value float32)

	// The source mode for the input bus of the audio environment node.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
	SetSourceMode(value AVAudio3DMixingSourceMode)

	// The type of rendering algorithm the mixer uses.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
	SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm)
}

// AVAudio3DMixingObject wraps an existing Objective-C object that conforms to the AVAudio3DMixing protocol.
type AVAudio3DMixingObject struct {
	objectivec.Object
}

func (o AVAudio3DMixingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudio3DMixingObjectFromID constructs a [AVAudio3DMixingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudio3DMixingObjectFromID(id objc.ID) AVAudio3DMixingObject {
	return AVAudio3DMixingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (o AVAudio3DMixingObject) Obstruction() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("obstruction"))
	return rv
}

// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (o AVAudio3DMixingObject) Occlusion() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("occlusion"))
	return rv
}

// The location of the source in the 3D environment.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (o AVAudio3DMixingObject) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](o.ID, objc.Sel("position"))
	return rv
}

// A value that changes the playback rate of the input signal.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (o AVAudio3DMixingObject) Rate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("rate"))
	return rv
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (o AVAudio3DMixingObject) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](o.ID, objc.Sel("pointSourceInHeadMode"))
	return rv
}

// A value that controls the blend of dry and reverb processed audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (o AVAudio3DMixingObject) ReverbBlend() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("reverbBlend"))
	return rv
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (o AVAudio3DMixingObject) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](o.ID, objc.Sel("sourceMode"))
	return rv
}

// The type of rendering algorithm the mixer uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (o AVAudio3DMixingObject) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](o.ID, objc.Sel("renderingAlgorithm"))
	return rv
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
func (o AVAudio3DMixingObject) SetObstruction(value float32) {
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
func (o AVAudio3DMixingObject) SetOcclusion(value float32) {
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
func (o AVAudio3DMixingObject) SetPosition(value AVAudio3DPoint) {
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
func (o AVAudio3DMixingObject) SetRate(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRate:"), value)
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (o AVAudio3DMixingObject) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
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
func (o AVAudio3DMixingObject) SetReverbBlend(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setReverbBlend:"), value)
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (o AVAudio3DMixingObject) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceMode:"), value)
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
func (o AVAudio3DMixingObject) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](o.ID, objc.Sel("setRenderingAlgorithm:"), value)
}
