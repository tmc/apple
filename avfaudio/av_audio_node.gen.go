// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

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

// An object you use for audio generation, processing, or an I/O block.
//
// # Overview
//
// An [AVAudioEngine] object contains instances of audio nodes that you
// attach, and this base class provides common functionality. Instances of
// this class don’t provide useful functionality until you attach them to an
// engine.
//
// Nodes have input and output busses that serve as connection points. For
// example, an effect has one input bus and one output bus, and a mixer has
// multiple input busses and one output bus.
//
// A bus contains a format the framework expresses in terms of sample rate and
// channel count. Formats must match exactly when making connections between
// nodes, excluding [AVAudioMixerNode] and [AVAudioOutputNode].
//
// # Configuring an Input Format Bus
//
//   - [AVAudioNode.InputFormatForBus]: Gets the input format for the bus you specify.
//   - [AVAudioNode.NameForInputBus]: Gets the name of the input bus you specify.
//   - [AVAudioNode.NumberOfInputs]: The number of input busses for the node.
//
// # Creating an Output Format Bus
//
//   - [AVAudioNode.OutputFormatForBus]: Retrieves the output format for the bus you specify.
//   - [AVAudioNode.NameForOutputBus]: Retrieves the name of the output bus you specify.
//   - [AVAudioNode.NumberOfOutputs]: The number of output busses for the node.
//
// # Installing and Removing an Audio Tap
//
//   - [AVAudioNode.InstallTapOnBusBufferSizeFormatBlock]: Installs an audio tap on a bus you specify to record, monitor, and observe the output of the node.
//   - [AVAudioNode.RemoveTapOnBus]: Removes an audio tap on a bus you specify.
//
// # Getting the Audio Engine for the Node
//
//   - [AVAudioNode.Engine]: The audio engine that manages the node, if any.
//
// # Getting the Latest Node Render Time
//
//   - [AVAudioNode.LastRenderTime]: The most recent render time.
//
// # Getting Audio Node Properties
//
//   - [AVAudioNode.AUAudioUnit]: An audio unit object that wraps or underlies the implementation’s audio unit.
//   - [AVAudioNode.Latency]: The processing latency of the node, in seconds.
//   - [AVAudioNode.OutputPresentationLatency]: The maximum render pipeline latency downstream of the node, in seconds.
//
// # Resetting the Audio Node
//
//   - [AVAudioNode.Reset]: Clears a unit’s previous processing state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
type AVAudioNode struct {
	objectivec.Object
}

// AVAudioNodeFromID constructs a [AVAudioNode] from an objc.ID.
//
// An object you use for audio generation, processing, or an I/O block.
func AVAudioNodeFromID(id objc.ID) AVAudioNode {
	return AVAudioNode{objectivec.Object{ID: id}}
}

// NOTE: AVAudioNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioNode] class.
//
// # Configuring an Input Format Bus
//
//   - [IAVAudioNode.InputFormatForBus]: Gets the input format for the bus you specify.
//   - [IAVAudioNode.NameForInputBus]: Gets the name of the input bus you specify.
//   - [IAVAudioNode.NumberOfInputs]: The number of input busses for the node.
//
// # Creating an Output Format Bus
//
//   - [IAVAudioNode.OutputFormatForBus]: Retrieves the output format for the bus you specify.
//   - [IAVAudioNode.NameForOutputBus]: Retrieves the name of the output bus you specify.
//   - [IAVAudioNode.NumberOfOutputs]: The number of output busses for the node.
//
// # Installing and Removing an Audio Tap
//
//   - [IAVAudioNode.InstallTapOnBusBufferSizeFormatBlock]: Installs an audio tap on a bus you specify to record, monitor, and observe the output of the node.
//   - [IAVAudioNode.RemoveTapOnBus]: Removes an audio tap on a bus you specify.
//
// # Getting the Audio Engine for the Node
//
//   - [IAVAudioNode.Engine]: The audio engine that manages the node, if any.
//
// # Getting the Latest Node Render Time
//
//   - [IAVAudioNode.LastRenderTime]: The most recent render time.
//
// # Getting Audio Node Properties
//
//   - [IAVAudioNode.AUAudioUnit]: An audio unit object that wraps or underlies the implementation’s audio unit.
//   - [IAVAudioNode.Latency]: The processing latency of the node, in seconds.
//   - [IAVAudioNode.OutputPresentationLatency]: The maximum render pipeline latency downstream of the node, in seconds.
//
// # Resetting the Audio Node
//
//   - [IAVAudioNode.Reset]: Clears a unit’s previous processing state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
type IAVAudioNode interface {
	objectivec.IObject

	// Topic: Configuring an Input Format Bus

	// Gets the input format for the bus you specify.
	InputFormatForBus(bus AVAudioNodeBus) IAVAudioFormat
	// Gets the name of the input bus you specify.
	NameForInputBus(bus AVAudioNodeBus) string
	// The number of input busses for the node.
	NumberOfInputs() uint

	// Topic: Creating an Output Format Bus

	// Retrieves the output format for the bus you specify.
	OutputFormatForBus(bus AVAudioNodeBus) IAVAudioFormat
	// Retrieves the name of the output bus you specify.
	NameForOutputBus(bus AVAudioNodeBus) string
	// The number of output busses for the node.
	NumberOfOutputs() uint

	// Topic: Installing and Removing an Audio Tap

	// Installs an audio tap on a bus you specify to record, monitor, and observe the output of the node.
	InstallTapOnBusBufferSizeFormatBlock(bus AVAudioNodeBus, bufferSize AVAudioFrameCount, format IAVAudioFormat, tapBlock AVAudioNodeTapBlock)
	// Removes an audio tap on a bus you specify.
	RemoveTapOnBus(bus AVAudioNodeBus)

	// Topic: Getting the Audio Engine for the Node

	// The audio engine that manages the node, if any.
	Engine() IAVAudioEngine

	// Topic: Getting the Latest Node Render Time

	// The most recent render time.
	LastRenderTime() IAVAudioTime

	// Topic: Getting Audio Node Properties

	// An audio unit object that wraps or underlies the implementation’s audio unit.
	AUAudioUnit() objectivec.IObject
	// The processing latency of the node, in seconds.
	Latency() float64
	// The maximum render pipeline latency downstream of the node, in seconds.
	OutputPresentationLatency() float64

	// Topic: Resetting the Audio Node

	// Clears a unit’s previous processing state.
	Reset()
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

// Gets the input format for the bus you specify.
//
// bus: An audio node bus.
//
// # Return Value
//
// An [AVAudioFormat] instance that represents the input format of the bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/inputFormat(forBus:)
func (a AVAudioNode) InputFormatForBus(bus AVAudioNodeBus) IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFormatForBus:"), bus)
	return AVAudioFormatFromID(rv)
}

// Gets the name of the input bus you specify.
//
// bus: The input bus from an audio node.
//
// # Return Value
//
// The name of the input bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forInputBus:)
func (a AVAudioNode) NameForInputBus(bus AVAudioNodeBus) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nameForInputBus:"), bus)
	return foundation.NSStringFromID(rv).String()
}

// Retrieves the output format for the bus you specify.
//
// bus: An audio node bus.
//
// # Return Value
//
// An [AVAudioFormat] instance that represents the output format of the bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputFormat(forBus:)
func (a AVAudioNode) OutputFormatForBus(bus AVAudioNodeBus) IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFormatForBus:"), bus)
	return AVAudioFormatFromID(rv)
}

// Retrieves the name of the output bus you specify.
//
// bus: The output bus from an audio node.
//
// # Return Value
//
// The name of the output bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forOutputBus:)
func (a AVAudioNode) NameForOutputBus(bus AVAudioNodeBus) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nameForOutputBus:"), bus)
	return foundation.NSStringFromID(rv).String()
}

// Installs an audio tap on a bus you specify to record, monitor, and observe
// the output of the node.
//
// bus: The output bus to attach the tap to.
//
// bufferSize: The size of the incoming buffers. The implementation may choose another
// size.
//
// format: If non-`nil`, the framework applies this format to the output bus you
// specify. An error occurs when attaching to an output bus that’s already
// in a connected state. The tap and connection formats (if non-`nil`) on the
// bus need to be identical. Otherwise, the latter operation overrides the
// previous format.
//
// For [AVAudioOutputNode], you must specify the tap format as `nil`.
//
// tapBlock: A block the framework calls with audio buffers.
//
// # Discussion
//
// You can install and remove taps while the engine is in a running state. You
// can install only one tap on any bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/installTap(onBus:bufferSize:format:block:)
func (a AVAudioNode) InstallTapOnBusBufferSizeFormatBlock(bus AVAudioNodeBus, bufferSize AVAudioFrameCount, format IAVAudioFormat, tapBlock AVAudioNodeTapBlock) {
	objc.Send[objc.ID](a.ID, objc.Sel("installTapOnBus:bufferSize:format:block:"), bus, bufferSize, format, tapBlock)
}

// Removes an audio tap on a bus you specify.
//
// bus: The node output bus with the tap to remove.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/removeTap(onBus:)
func (a AVAudioNode) RemoveTapOnBus(bus AVAudioNodeBus) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeTapOnBus:"), bus)
}

// Clears a unit’s previous processing state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/reset()
func (a AVAudioNode) Reset() {
	objc.Send[objc.ID](a.ID, objc.Sel("reset"))
}

// The number of input busses for the node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfInputs
func (a AVAudioNode) NumberOfInputs() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("numberOfInputs"))
	return rv
}

// The number of output busses for the node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfOutputs
func (a AVAudioNode) NumberOfOutputs() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("numberOfOutputs"))
	return rv
}

// The audio engine that manages the node, if any.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/engine
func (a AVAudioNode) Engine() IAVAudioEngine {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("engine"))
	return AVAudioEngineFromID(objc.ID(rv))
}

// The most recent render time.
//
// # Discussion
//
// This value is `nil` if the engine isn’t running or if you don’t connect
// the node to an input or output node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/lastRenderTime
func (a AVAudioNode) LastRenderTime() IAVAudioTime {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("lastRenderTime"))
	return AVAudioTimeFromID(objc.ID(rv))
}

// An audio unit object that wraps or underlies the implementation’s audio
// unit.
//
// # Discussion
//
// This provides an [AUAudioUnit] that either wraps or underlies the
// implementation’s audio unit, depending on how the app packages the audio
// unit. Apps interact with this to control custom properties, select presets,
// and change parameters.
//
// Don’t perform operations directly on the audio unit that may conflict
// with the engine’s state, which includes changing the initialization
// state, stream formats, channel layouts, or connections to other audio
// units.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/auAudioUnit
//
// [AUAudioUnit]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit
func (a AVAudioNode) AUAudioUnit() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AUAudioUnit"))
	return objectivec.Object{ID: rv}
}

// The processing latency of the node, in seconds.
//
// # Discussion
//
// This latency reflects the delay due to signal processing. A value of `0`
// indicates either no latency or an unknown latency.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/latency
func (a AVAudioNode) Latency() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("latency"))
	return rv
}

// The maximum render pipeline latency downstream of the node, in seconds.
//
// # Discussion
//
// This latency describes the maximum time it takes to present the audio at
// the output of a node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputPresentationLatency
func (a AVAudioNode) OutputPresentationLatency() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("outputPresentationLatency"))
	return rv
}
