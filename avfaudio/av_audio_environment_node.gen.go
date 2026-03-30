// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEnvironmentNode] class.
var (
	_AVAudioEnvironmentNodeClass     AVAudioEnvironmentNodeClass
	_AVAudioEnvironmentNodeClassOnce sync.Once
)

func getAVAudioEnvironmentNodeClass() AVAudioEnvironmentNodeClass {
	_AVAudioEnvironmentNodeClassOnce.Do(func() {
		_AVAudioEnvironmentNodeClass = AVAudioEnvironmentNodeClass{class: objc.GetClass("AVAudioEnvironmentNode")}
	})
	return _AVAudioEnvironmentNodeClass
}

// GetAVAudioEnvironmentNodeClass returns the class object for AVAudioEnvironmentNode.
func GetAVAudioEnvironmentNodeClass() AVAudioEnvironmentNodeClass {
	return getAVAudioEnvironmentNodeClass()
}

type AVAudioEnvironmentNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEnvironmentNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEnvironmentNodeClass) Alloc() AVAudioEnvironmentNode {
	rv := objc.Send[AVAudioEnvironmentNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that simulates a 3D audio environment.
//
// # Overview
//
// The [AVAudioEnvironmentNode] class is a mixer node that simulates a 3D
// audio environment. Any node that conforms to [AVAudioMixing] can act as a
// source node, such as [AVAudioPlayerNode].
//
// The environment node has an implicit listener. You set the listener’s
// position and orientation, and the system then controls the way the user
// experiences the virtual world.
//
// To help characterize the environment, this class defines properties for
// distance attenuation and reverberation.
//
// [AVAudio3DMixingSourceMode] affects how inputs with different channel
// configurations render. Spatialization applies only to inputs with a mono
// channel connection format. This class doesn’t spatialize stereo inputs or
// support inputs with connection formats of more than two channels.
//
// To set the node’s output to a multichannel format, use an [AVAudioFormat]
// that has one of the following [Audio Channel Layout Tags]:
//
// - [KAudioChannelLayoutTag_AudioUnit_4] -
// [KAudioChannelLayoutTag_AudioUnit_5_0] -
// [KAudioChannelLayoutTag_AudioUnit_6_0] -
// [KAudioChannelLayoutTag_AudioUnit_7_0] -
// [KAudioChannelLayoutTag_AudioUnit_7_0_Front] -
// [KAudioChannelLayoutTag_AudioUnit_8]
//
// # Getting and Setting Positional Properties
//
//   - [AVAudioEnvironmentNode.ListenerPosition]: The listener’s position in the 3D environment.
//   - [AVAudioEnvironmentNode.SetListenerPosition]
//   - [AVAudioEnvironmentNode.ListenerAngularOrientation]: The listener’s angular orientation in the environment.
//   - [AVAudioEnvironmentNode.SetListenerAngularOrientation]
//   - [AVAudioEnvironmentNode.ListenerVectorOrientation]: The listener’s vector orientation in the environment.
//   - [AVAudioEnvironmentNode.SetListenerVectorOrientation]
//
// # Getting Attenuation and Reverb Properties
//
//   - [AVAudioEnvironmentNode.DistanceAttenuationParameters]: The distance attenuation parameters for the environment.
//   - [AVAudioEnvironmentNode.ReverbParameters]: The reverb parameters for the environment.
//
// # Getting and Setting Environment Properties
//
//   - [AVAudioEnvironmentNode.OutputVolume]: The mixer’s output volume.
//   - [AVAudioEnvironmentNode.SetOutputVolume]
//   - [AVAudioEnvironmentNode.OutputType]: The type of output hardware.
//   - [AVAudioEnvironmentNode.SetOutputType]
//
// # Getting the Available Rendering Algorithms
//
//   - [AVAudioEnvironmentNode.ApplicableRenderingAlgorithms]: An array of rendering algorithms applicable to the environment node.
//
// # Getting the Head Tracking Status
//
//   - [AVAudioEnvironmentNode.ListenerHeadTrackingEnabled]: A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//   - [AVAudioEnvironmentNode.SetListenerHeadTrackingEnabled]
//
// # Getting the Input Bus
//
//   - [AVAudioEnvironmentNode.NextAvailableInputBus]: An unused input bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode
//
// [AVAudio3DMixingSourceMode]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode
// [Audio Channel Layout Tags]: https://developer.apple.com/documentation/CoreAudioTypes/audio-channel-layout-tags
type AVAudioEnvironmentNode struct {
	AVAudioNode
}

// AVAudioEnvironmentNodeFromID constructs a [AVAudioEnvironmentNode] from an objc.ID.
//
// An object that simulates a 3D audio environment.
func AVAudioEnvironmentNodeFromID(id objc.ID) AVAudioEnvironmentNode {
	return AVAudioEnvironmentNode{AVAudioNode: AVAudioNodeFromID(id)}
}

// NOTE: AVAudioEnvironmentNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioEnvironmentNode] class.
//
// # Getting and Setting Positional Properties
//
//   - [IAVAudioEnvironmentNode.ListenerPosition]: The listener’s position in the 3D environment.
//   - [IAVAudioEnvironmentNode.SetListenerPosition]
//   - [IAVAudioEnvironmentNode.ListenerAngularOrientation]: The listener’s angular orientation in the environment.
//   - [IAVAudioEnvironmentNode.SetListenerAngularOrientation]
//   - [IAVAudioEnvironmentNode.ListenerVectorOrientation]: The listener’s vector orientation in the environment.
//   - [IAVAudioEnvironmentNode.SetListenerVectorOrientation]
//
// # Getting Attenuation and Reverb Properties
//
//   - [IAVAudioEnvironmentNode.DistanceAttenuationParameters]: The distance attenuation parameters for the environment.
//   - [IAVAudioEnvironmentNode.ReverbParameters]: The reverb parameters for the environment.
//
// # Getting and Setting Environment Properties
//
//   - [IAVAudioEnvironmentNode.OutputVolume]: The mixer’s output volume.
//   - [IAVAudioEnvironmentNode.SetOutputVolume]
//   - [IAVAudioEnvironmentNode.OutputType]: The type of output hardware.
//   - [IAVAudioEnvironmentNode.SetOutputType]
//
// # Getting the Available Rendering Algorithms
//
//   - [IAVAudioEnvironmentNode.ApplicableRenderingAlgorithms]: An array of rendering algorithms applicable to the environment node.
//
// # Getting the Head Tracking Status
//
//   - [IAVAudioEnvironmentNode.ListenerHeadTrackingEnabled]: A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//   - [IAVAudioEnvironmentNode.SetListenerHeadTrackingEnabled]
//
// # Getting the Input Bus
//
//   - [IAVAudioEnvironmentNode.NextAvailableInputBus]: An unused input bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode
type IAVAudioEnvironmentNode interface {
	IAVAudioNode
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Getting and Setting Positional Properties

	// The listener’s position in the 3D environment.
	ListenerPosition() AVAudio3DPoint
	SetListenerPosition(value AVAudio3DPoint)
	// The listener’s angular orientation in the environment.
	ListenerAngularOrientation() AVAudio3DAngularOrientation
	SetListenerAngularOrientation(value AVAudio3DAngularOrientation)
	// The listener’s vector orientation in the environment.
	ListenerVectorOrientation() AVAudio3DVectorOrientation
	SetListenerVectorOrientation(value AVAudio3DVectorOrientation)

	// Topic: Getting Attenuation and Reverb Properties

	// The distance attenuation parameters for the environment.
	DistanceAttenuationParameters() IAVAudioEnvironmentDistanceAttenuationParameters
	// The reverb parameters for the environment.
	ReverbParameters() IAVAudioEnvironmentReverbParameters

	// Topic: Getting and Setting Environment Properties

	// The mixer’s output volume.
	OutputVolume() float32
	SetOutputVolume(value float32)
	// The type of output hardware.
	OutputType() AVAudioEnvironmentOutputType
	SetOutputType(value AVAudioEnvironmentOutputType)

	// Topic: Getting the Available Rendering Algorithms

	// An array of rendering algorithms applicable to the environment node.
	ApplicableRenderingAlgorithms() []foundation.NSNumber

	// Topic: Getting the Head Tracking Status

	// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
	ListenerHeadTrackingEnabled() bool
	SetListenerHeadTrackingEnabled(value bool)

	// Topic: Getting the Input Bus

	// An unused input bus.
	NextAvailableInputBus() AVAudioNodeBus

	// A quadraphonic symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_4() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_4(value objectivec.IObject)
	// A 5-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_5_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_5_0(value objectivec.IObject)
	// A 6-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_6_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_6_0(value objectivec.IObject)
	// A 7-channel surround-based layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_7_0(value objectivec.IObject)
	// An alternate 7-channel surround-based layout, for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_7_0_Front() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value objectivec.IObject)
	// An octagonal symmetrical layout, recommended for use by audio units.
	KAudioChannelLayoutTag_AudioUnit_8() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_8(value objectivec.IObject)
}

// Init initializes the instance.
func (a AVAudioEnvironmentNode) Init() AVAudioEnvironmentNode {
	rv := objc.Send[AVAudioEnvironmentNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEnvironmentNode) Autorelease() AVAudioEnvironmentNode {
	rv := objc.Send[AVAudioEnvironmentNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEnvironmentNode creates a new AVAudioEnvironmentNode instance.
func NewAVAudioEnvironmentNode() AVAudioEnvironmentNode {
	class := getAVAudioEnvironmentNodeClass()
	rv := objc.Send[AVAudioEnvironmentNode](objc.ID(class.class), objc.Sel("new"))
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
func (a AVAudioEnvironmentNode) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// The listener’s position in the 3D environment.
//
// # Discussion
//
// The system specifies the coordinates in meters.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerPosition
func (a AVAudioEnvironmentNode) ListenerPosition() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("listenerPosition"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioEnvironmentNode) SetListenerPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setListenerPosition:"), value)
}

// The listener’s angular orientation in the environment.
//
// # Discussion
//
// The system specifies all angles in degrees.
//
// The default orientation is with the listener looking directly along the
// negative z-axis (forward). This orientation has a yaw of `0.0` degrees, a
// pitch of `0.0` degrees, and a roll of `0.0` degrees.
//
// Changing this property results in a corresponding change in the
// [ListenerVectorOrientation] property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation
func (a AVAudioEnvironmentNode) ListenerAngularOrientation() AVAudio3DAngularOrientation {
	rv := objc.Send[AVAudio3DAngularOrientation](a.ID, objc.Sel("listenerAngularOrientation"))
	return AVAudio3DAngularOrientation(rv)
}
func (a AVAudioEnvironmentNode) SetListenerAngularOrientation(value AVAudio3DAngularOrientation) {
	objc.Send[struct{}](a.ID, objc.Sel("setListenerAngularOrientation:"), value)
}

// The listener’s vector orientation in the environment.
//
// # Discussion
//
// The default orientation is with the listener looking directly along the
// negative z-axis (forward).
//
// The node expresses a forward vector orientation as `(0, 0, -1)`, and an up
// vector as `(0, 1, 0)`.
//
// Changing this property results in a corresponding change in the
// [ListenerAngularOrientation] property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerVectorOrientation
func (a AVAudioEnvironmentNode) ListenerVectorOrientation() AVAudio3DVectorOrientation {
	rv := objc.Send[AVAudio3DVectorOrientation](a.ID, objc.Sel("listenerVectorOrientation"))
	return AVAudio3DVectorOrientation(rv)
}
func (a AVAudioEnvironmentNode) SetListenerVectorOrientation(value AVAudio3DVectorOrientation) {
	objc.Send[struct{}](a.ID, objc.Sel("setListenerVectorOrientation:"), value)
}

// The distance attenuation parameters for the environment.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/distanceAttenuationParameters
func (a AVAudioEnvironmentNode) DistanceAttenuationParameters() IAVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("distanceAttenuationParameters"))
	return AVAudioEnvironmentDistanceAttenuationParametersFromID(objc.ID(rv))
}

// The reverb parameters for the environment.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/reverbParameters
func (a AVAudioEnvironmentNode) ReverbParameters() IAVAudioEnvironmentReverbParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("reverbParameters"))
	return AVAudioEnvironmentReverbParametersFromID(objc.ID(rv))
}

// The mixer’s output volume.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputVolume
func (a AVAudioEnvironmentNode) OutputVolume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("outputVolume"))
	return rv
}
func (a AVAudioEnvironmentNode) SetOutputVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputVolume:"), value)
}

// The type of output hardware.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputType
func (a AVAudioEnvironmentNode) OutputType() AVAudioEnvironmentOutputType {
	rv := objc.Send[AVAudioEnvironmentOutputType](a.ID, objc.Sel("outputType"))
	return AVAudioEnvironmentOutputType(rv)
}
func (a AVAudioEnvironmentNode) SetOutputType(value AVAudioEnvironmentOutputType) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputType:"), value)
}

// An array of rendering algorithms applicable to the environment node.
//
// # Discussion
//
// The [AVAudioEnvironmentNode] class supports several rendering algorithms
// for each input bus as [AVAudio3DMixingRenderingAlgorithm] defines.
//
// Depending on the current output format of the environment node, this method
// returns an immutable array of the applicable rendering algorithms. This
// subset of applicable rendering algorithms is important when you configure
// the environment node to a multichannel output format because only a subset
// of the algorithms render to all of the channels.
//
// Retrieve the applicable algorithms after a successful connection to the
// destination node through one of the [AVAudioEngine] connect methods.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/applicableRenderingAlgorithms
//
// [AVAudio3DMixingRenderingAlgorithm]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm
func (a AVAudioEnvironmentNode) ApplicableRenderingAlgorithms() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("applicableRenderingAlgorithms"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A Boolean value that indicates whether the listener orientation is
// automatically rotated based on head orientation.
//
// # Discussion
//
// To enable head tracking, your app must have the
// [com.apple.developer.coremotion.head-pose] entitlement.
//
// Set this value to true to enable head tracking with compatible AirPods.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled
//
// [com.apple.developer.coremotion.head-pose]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.coremotion.head-pose
func (a AVAudioEnvironmentNode) ListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isListenerHeadTrackingEnabled"))
	return rv
}
func (a AVAudioEnvironmentNode) SetListenerHeadTrackingEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setListenerHeadTrackingEnabled:"), value)
}

// An unused input bus.
//
// # Discussion
//
// This method finds and returns the first input bus that doesn’t have a
// connection with a node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/nextAvailableInputBus
func (a AVAudioEnvironmentNode) NextAvailableInputBus() AVAudioNodeBus {
	rv := objc.Send[AVAudioNodeBus](a.ID, objc.Sel("nextAvailableInputBus"))
	return AVAudioNodeBus(rv)
}

// A quadraphonic symmetrical layout, recommended for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_4
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_4() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_4"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_4(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_4:"), value)
}

// A 5-channel surround-based layout, recommended for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_5_0
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_5_0() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_5_0"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_5_0(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_5_0:"), value)
}

// A 6-channel surround-based layout, recommended for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_6_0
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_6_0() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_6_0"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_6_0(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_6_0:"), value)
}

// A 7-channel surround-based layout, recommended for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0:"), value)
}

// An alternate 7-channel surround-based layout, for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0_Front
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0_Front() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0_Front"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0_Front:"), value)
}

// An octagonal symmetrical layout, recommended for use by audio units.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_8
func (a AVAudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_8() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_8"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_8(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_8:"), value)
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
func (a AVAudioEnvironmentNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioEnvironmentNode) SetObstruction(value float32) {
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
func (a AVAudioEnvironmentNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioEnvironmentNode) SetOcclusion(value float32) {
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
func (a AVAudioEnvironmentNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioEnvironmentNode) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioEnvironmentNode) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioEnvironmentNode) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
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
func (a AVAudioEnvironmentNode) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioEnvironmentNode) SetPosition(value AVAudio3DPoint) {
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
func (a AVAudioEnvironmentNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioEnvironmentNode) SetRate(value float32) {
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
func (a AVAudioEnvironmentNode) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioEnvironmentNode) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
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
func (a AVAudioEnvironmentNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioEnvironmentNode) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioEnvironmentNode) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioEnvironmentNode) SetSourceMode(value AVAudio3DMixingSourceMode) {
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
func (a AVAudioEnvironmentNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioEnvironmentNode) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

// Protocol methods for AVAudioMixing
