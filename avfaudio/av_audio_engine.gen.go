// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEngine] class.
var (
	_AVAudioEngineClass     AVAudioEngineClass
	_AVAudioEngineClassOnce sync.Once
)

func getAVAudioEngineClass() AVAudioEngineClass {
	_AVAudioEngineClassOnce.Do(func() {
		_AVAudioEngineClass = AVAudioEngineClass{class: objc.GetClass("AVAudioEngine")}
	})
	return _AVAudioEngineClass
}

// GetAVAudioEngineClass returns the class object for AVAudioEngine.
func GetAVAudioEngineClass() AVAudioEngineClass {
	return getAVAudioEngineClass()
}

type AVAudioEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEngineClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEngineClass) Alloc() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that manages a graph of audio nodes, controls playback, and
// configures real-time rendering constraints.
//
// # Overview
//
// An audio engine object contains a group of [AVAudioNode] instances that you
// attach to form an audio processing chain.
//
// [media-3901205]
//
// You can connect, disconnect, and remove audio nodes during runtime with
// minor limitations. Removing an audio node that has differing channel
// counts, or that’s a mixer, can break the graph. Reconnect audio nodes
// only when they’re upstream of a mixer.
//
// By default, Audio Engine renders to a connected audio device in real time.
// You can configure the engine to operate in manual rendering mode when you
// need to render at, or faster than, real time. In that mode, the engine
// disconnects from audio devices and your app drives the rendering.
//
// # Create an Engine for Audio File Playback
//
// To play an audio file, you create an [AVAudioFile] with a file that’s
// open for reading. Create an audio engine object and an [AVAudioPlayerNode]
// instance, and then attach the player node to the engine. Next, connect the
// player node to the audio engine’s output node. The engine performs audio
// output through an output node, which is a singleton that the engine creates
// the first time you access it.
//
// Then schedule the audio file for full playback. The callback notifies your
// app when playback completes.
//
// Before you play the audio, start the engine.
//
// When you’re done, stop the player and the engine.
//
// # Attaching and Detaching Audio Nodes
//
//   - [AVAudioEngine.AttachNode]: Attaches an audio node to the audio engine.
//   - [AVAudioEngine.DetachNode]: Detaches an audio node from the audio engine.
//   - [AVAudioEngine.AttachedNodes]: A read-only set that contains the nodes you attach to the audio engine.
//
// # Getting the Input, Output, and Main Mixer Nodes
//
//   - [AVAudioEngine.InputNode]: The audio engine’s singleton input audio node.
//   - [AVAudioEngine.OutputNode]: The audio engine’s singleton output audio node.
//   - [AVAudioEngine.MainMixerNode]: The audio engine’s optional singleton main mixer node.
//
// # Connecting and Disconnecting Audio Nodes
//
//   - [AVAudioEngine.ConnectToFormat]: Establishes a connection between two nodes.
//   - [AVAudioEngine.ConnectToFromBusToBusFormat]: Establishes a connection between two nodes, specifying the input and output busses.
//   - [AVAudioEngine.DisconnectNodeInput]: Removes all input connections of the node.
//   - [AVAudioEngine.DisconnectNodeInputBus]: Removes the input connection of a node on the specified bus.
//   - [AVAudioEngine.DisconnectNodeOutput]: Removes all output connections of a node.
//   - [AVAudioEngine.DisconnectNodeOutputBus]: Removes the output connection of a node on the specified bus.
//
// # Managing MIDI Nodes
//
//   - [AVAudioEngine.ConnectMIDIToFormatEventListBlock]: Establishes a MIDI connection between two nodes.
//   - [AVAudioEngine.ConnectMIDIToNodesFormatEventListBlock]: Establishes a MIDI connection between a source node and multiple destination nodes.
//   - [AVAudioEngine.DisconnectMIDIFrom]: Removes a MIDI connection between two nodes.
//   - [AVAudioEngine.DisconnectMIDIFromNodes]: Removes a MIDI connection between one source node and multiple destination nodes.
//   - [AVAudioEngine.DisconnectMIDIInput]: Disconnects all input MIDI connections from a node.
//   - [AVAudioEngine.DisconnectMIDIOutput]: Disconnects all output MIDI connections from a node.
//
// # Playing Audio
//
//   - [AVAudioEngine.Prepare]: Prepares the audio engine for starting.
//   - [AVAudioEngine.StartAndReturnError]: Starts the audio engine.
//   - [AVAudioEngine.Running]: A Boolean value that indicates whether the audio engine is running.
//   - [AVAudioEngine.Pause]: Pauses the audio engine.
//   - [AVAudioEngine.Stop]: Stops the audio engine and releases any previously prepared resources.
//   - [AVAudioEngine.Reset]: Resets all audio nodes in the audio engine.
//   - [AVAudioEngine.MusicSequence]: The music sequence instance that you attach to the audio engine, if any.
//   - [AVAudioEngine.SetMusicSequence]
//
// # Manually Rendering an Audio Engine
//
//   - [AVAudioEngine.EnableManualRenderingModeFormatMaximumFrameCountError]: Sets the engine to operate in manual rendering mode with the render format and maximum frame count you specify.
//   - [AVAudioEngine.DisableManualRenderingMode]: Sets the engine to render to or from an audio device.
//   - [AVAudioEngine.RenderOfflineToBufferError]: Makes a render call to the engine operating in the offline manual rendering mode.
//
// # Getting Manual Rendering Properties
//
//   - [AVAudioEngine.ManualRenderingBlock]: The block that renders the engine when operating in manual rendering mode.
//   - [AVAudioEngine.ManualRenderingFormat]: The render format of the engine in manual rendering mode.
//   - [AVAudioEngine.ManualRenderingMaximumFrameCount]: The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
//   - [AVAudioEngine.ManualRenderingMode]: The manual rendering mode configured on the engine.
//   - [AVAudioEngine.ManualRenderingSampleTime]: An indication of where the engine is on its render timeline in manual rendering mode.
//   - [AVAudioEngine.AutoShutdownEnabled]: A Boolean value that indicates whether autoshutdown is in an enabled state.
//   - [AVAudioEngine.SetAutoShutdownEnabled]
//   - [AVAudioEngine.IsInManualRenderingMode]: A Boolean value that indicates whether the engine is operating in manual rendering mode.
//
// # Using Connection Points
//
//   - [AVAudioEngine.ConnectToConnectionPointsFromBusFormat]: Establishes a connection between a source node and multiple destination nodes.
//   - [AVAudioEngine.InputConnectionPointForNodeInputBus]: Returns connection information about a node’s input bus.
//   - [AVAudioEngine.OutputConnectionPointsForNodeOutputBus]: Returns connection information about a node’s output bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type AVAudioEngine struct {
	objectivec.Object
}

// AVAudioEngineFromID constructs a [AVAudioEngine] from an objc.ID.
//
// An object that manages a graph of audio nodes, controls playback, and
// configures real-time rendering constraints.
func AVAudioEngineFromID(id objc.ID) AVAudioEngine {
	return AVAudioEngine{objectivec.Object{ID: id}}
}

// NOTE: AVAudioEngine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioEngine] class.
//
// # Attaching and Detaching Audio Nodes
//
//   - [IAVAudioEngine.AttachNode]: Attaches an audio node to the audio engine.
//   - [IAVAudioEngine.DetachNode]: Detaches an audio node from the audio engine.
//   - [IAVAudioEngine.AttachedNodes]: A read-only set that contains the nodes you attach to the audio engine.
//
// # Getting the Input, Output, and Main Mixer Nodes
//
//   - [IAVAudioEngine.InputNode]: The audio engine’s singleton input audio node.
//   - [IAVAudioEngine.OutputNode]: The audio engine’s singleton output audio node.
//   - [IAVAudioEngine.MainMixerNode]: The audio engine’s optional singleton main mixer node.
//
// # Connecting and Disconnecting Audio Nodes
//
//   - [IAVAudioEngine.ConnectToFormat]: Establishes a connection between two nodes.
//   - [IAVAudioEngine.ConnectToFromBusToBusFormat]: Establishes a connection between two nodes, specifying the input and output busses.
//   - [IAVAudioEngine.DisconnectNodeInput]: Removes all input connections of the node.
//   - [IAVAudioEngine.DisconnectNodeInputBus]: Removes the input connection of a node on the specified bus.
//   - [IAVAudioEngine.DisconnectNodeOutput]: Removes all output connections of a node.
//   - [IAVAudioEngine.DisconnectNodeOutputBus]: Removes the output connection of a node on the specified bus.
//
// # Managing MIDI Nodes
//
//   - [IAVAudioEngine.ConnectMIDIToFormatEventListBlock]: Establishes a MIDI connection between two nodes.
//   - [IAVAudioEngine.ConnectMIDIToNodesFormatEventListBlock]: Establishes a MIDI connection between a source node and multiple destination nodes.
//   - [IAVAudioEngine.DisconnectMIDIFrom]: Removes a MIDI connection between two nodes.
//   - [IAVAudioEngine.DisconnectMIDIFromNodes]: Removes a MIDI connection between one source node and multiple destination nodes.
//   - [IAVAudioEngine.DisconnectMIDIInput]: Disconnects all input MIDI connections from a node.
//   - [IAVAudioEngine.DisconnectMIDIOutput]: Disconnects all output MIDI connections from a node.
//
// # Playing Audio
//
//   - [IAVAudioEngine.Prepare]: Prepares the audio engine for starting.
//   - [IAVAudioEngine.StartAndReturnError]: Starts the audio engine.
//   - [IAVAudioEngine.Running]: A Boolean value that indicates whether the audio engine is running.
//   - [IAVAudioEngine.Pause]: Pauses the audio engine.
//   - [IAVAudioEngine.Stop]: Stops the audio engine and releases any previously prepared resources.
//   - [IAVAudioEngine.Reset]: Resets all audio nodes in the audio engine.
//   - [IAVAudioEngine.MusicSequence]: The music sequence instance that you attach to the audio engine, if any.
//   - [IAVAudioEngine.SetMusicSequence]
//
// # Manually Rendering an Audio Engine
//
//   - [IAVAudioEngine.EnableManualRenderingModeFormatMaximumFrameCountError]: Sets the engine to operate in manual rendering mode with the render format and maximum frame count you specify.
//   - [IAVAudioEngine.DisableManualRenderingMode]: Sets the engine to render to or from an audio device.
//   - [IAVAudioEngine.RenderOfflineToBufferError]: Makes a render call to the engine operating in the offline manual rendering mode.
//
// # Getting Manual Rendering Properties
//
//   - [IAVAudioEngine.ManualRenderingBlock]: The block that renders the engine when operating in manual rendering mode.
//   - [IAVAudioEngine.ManualRenderingFormat]: The render format of the engine in manual rendering mode.
//   - [IAVAudioEngine.ManualRenderingMaximumFrameCount]: The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
//   - [IAVAudioEngine.ManualRenderingMode]: The manual rendering mode configured on the engine.
//   - [IAVAudioEngine.ManualRenderingSampleTime]: An indication of where the engine is on its render timeline in manual rendering mode.
//   - [IAVAudioEngine.AutoShutdownEnabled]: A Boolean value that indicates whether autoshutdown is in an enabled state.
//   - [IAVAudioEngine.SetAutoShutdownEnabled]
//   - [IAVAudioEngine.IsInManualRenderingMode]: A Boolean value that indicates whether the engine is operating in manual rendering mode.
//
// # Using Connection Points
//
//   - [IAVAudioEngine.ConnectToConnectionPointsFromBusFormat]: Establishes a connection between a source node and multiple destination nodes.
//   - [IAVAudioEngine.InputConnectionPointForNodeInputBus]: Returns connection information about a node’s input bus.
//   - [IAVAudioEngine.OutputConnectionPointsForNodeOutputBus]: Returns connection information about a node’s output bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type IAVAudioEngine interface {
	objectivec.IObject

	// Topic: Attaching and Detaching Audio Nodes

	// Attaches an audio node to the audio engine.
	AttachNode(node IAVAudioNode)
	// Detaches an audio node from the audio engine.
	DetachNode(node IAVAudioNode)
	// A read-only set that contains the nodes you attach to the audio engine.
	AttachedNodes() foundation.INSSet

	// Topic: Getting the Input, Output, and Main Mixer Nodes

	// The audio engine’s singleton input audio node.
	InputNode() IAVAudioInputNode
	// The audio engine’s singleton output audio node.
	OutputNode() IAVAudioOutputNode
	// The audio engine’s optional singleton main mixer node.
	MainMixerNode() IAVAudioMixerNode

	// Topic: Connecting and Disconnecting Audio Nodes

	// Establishes a connection between two nodes.
	ConnectToFormat(node1 IAVAudioNode, node2 IAVAudioNode, format IAVAudioFormat)
	// Establishes a connection between two nodes, specifying the input and output busses.
	ConnectToFromBusToBusFormat(node1 IAVAudioNode, node2 IAVAudioNode, bus1 AVAudioNodeBus, bus2 AVAudioNodeBus, format IAVAudioFormat)
	// Removes all input connections of the node.
	DisconnectNodeInput(node IAVAudioNode)
	// Removes the input connection of a node on the specified bus.
	DisconnectNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus)
	// Removes all output connections of a node.
	DisconnectNodeOutput(node IAVAudioNode)
	// Removes the output connection of a node on the specified bus.
	DisconnectNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus)

	// Topic: Managing MIDI Nodes

	// Establishes a MIDI connection between two nodes.
	ConnectMIDIToFormatEventListBlock(sourceNode IAVAudioNode, destinationNode IAVAudioNode, format IAVAudioFormat, tapBlock objectivec.IObject)
	// Establishes a MIDI connection between a source node and multiple destination nodes.
	ConnectMIDIToNodesFormatEventListBlock(sourceNode IAVAudioNode, destinationNodes []AVAudioNode, format IAVAudioFormat, tapBlock objectivec.IObject)
	// Removes a MIDI connection between two nodes.
	DisconnectMIDIFrom(sourceNode IAVAudioNode, destinationNode IAVAudioNode)
	// Removes a MIDI connection between one source node and multiple destination nodes.
	DisconnectMIDIFromNodes(sourceNode IAVAudioNode, destinationNodes []AVAudioNode)
	// Disconnects all input MIDI connections from a node.
	DisconnectMIDIInput(node IAVAudioNode)
	// Disconnects all output MIDI connections from a node.
	DisconnectMIDIOutput(node IAVAudioNode)

	// Topic: Playing Audio

	// Prepares the audio engine for starting.
	Prepare()
	// Starts the audio engine.
	StartAndReturnError() (bool, error)
	// A Boolean value that indicates whether the audio engine is running.
	Running() bool
	// Pauses the audio engine.
	Pause()
	// Stops the audio engine and releases any previously prepared resources.
	Stop()
	// Resets all audio nodes in the audio engine.
	Reset()
	// The music sequence instance that you attach to the audio engine, if any.
	MusicSequence() objectivec.IObject
	SetMusicSequence(value objectivec.IObject)

	// Topic: Manually Rendering an Audio Engine

	// Sets the engine to operate in manual rendering mode with the render format and maximum frame count you specify.
	EnableManualRenderingModeFormatMaximumFrameCountError(mode AVAudioEngineManualRenderingMode, pcmFormat IAVAudioFormat, maximumFrameCount AVAudioFrameCount) (bool, error)
	// Sets the engine to render to or from an audio device.
	DisableManualRenderingMode()
	// Makes a render call to the engine operating in the offline manual rendering mode.
	RenderOfflineToBufferError(numberOfFrames AVAudioFrameCount, buffer IAVAudioPCMBuffer) (AVAudioEngineManualRenderingStatus, error)

	// Topic: Getting Manual Rendering Properties

	// The block that renders the engine when operating in manual rendering mode.
	ManualRenderingBlock() AVAudioEngineManualRenderingBlock
	// The render format of the engine in manual rendering mode.
	ManualRenderingFormat() IAVAudioFormat
	// The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
	ManualRenderingMaximumFrameCount() AVAudioFrameCount
	// The manual rendering mode configured on the engine.
	ManualRenderingMode() AVAudioEngineManualRenderingMode
	// An indication of where the engine is on its render timeline in manual rendering mode.
	ManualRenderingSampleTime() AVAudioFramePosition
	// A Boolean value that indicates whether autoshutdown is in an enabled state.
	AutoShutdownEnabled() bool
	SetAutoShutdownEnabled(value bool)
	// A Boolean value that indicates whether the engine is operating in manual rendering mode.
	IsInManualRenderingMode() bool

	// Topic: Using Connection Points

	// Establishes a connection between a source node and multiple destination nodes.
	ConnectToConnectionPointsFromBusFormat(sourceNode IAVAudioNode, destNodes []AVAudioConnectionPoint, sourceBus AVAudioNodeBus, format IAVAudioFormat)
	// Returns connection information about a node’s input bus.
	InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus) IAVAudioConnectionPoint
	// Returns connection information about a node’s output bus.
	OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus) []AVAudioConnectionPoint
}

// Init initializes the instance.
func (a AVAudioEngine) Init() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEngine) Autorelease() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEngine creates a new AVAudioEngine instance.
func NewAVAudioEngine() AVAudioEngine {
	class := getAVAudioEngineClass()
	rv := objc.Send[AVAudioEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Attaches an audio node to the audio engine.
//
// node: The audio node to attach.
//
// # Discussion
//
// An instance of [AVAudioNode] isn’t usable until you attach it to the
// audio engine using this method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/attach(_:)
func (a AVAudioEngine) AttachNode(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("attachNode:"), node)
}

// Detaches an audio node from the audio engine.
//
// node: The audio node to detach.
//
// # Discussion
//
// If necessary, the audio engine safely disconnects the audio node before
// detaching it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/detach(_:)
func (a AVAudioEngine) DetachNode(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("detachNode:"), node)
}

// Establishes a connection between two nodes.
//
// node1: The source audio node.
//
// node2: The destination audio node.
//
// format: If not [NULL], the engine uses this value for the format of the source
// audio node’s output bus. In all cases, the engine matches the format of
// the destination audio node’s input bus to the source audio node’s
// output bus.
//
// # Discussion
//
// This method calls [ConnectToFromBusToBusFormat] using bus `0` for the
// source audio node, and bus `0` for the destination audio node, except when
// a destination is a mixer, in which case, the destination is the mixer’s
// [NextAvailableInputBus].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:format:)
func (a AVAudioEngine) ConnectToFormat(node1 IAVAudioNode, node2 IAVAudioNode, format IAVAudioFormat) {
	objc.Send[objc.ID](a.ID, objc.Sel("connect:to:format:"), node1, node2, format)
}

// Establishes a connection between two nodes, specifying the input and output
// busses.
//
// node1: The source audio node.
//
// node2: The destination audio node.
//
// bus1: The output bus of the source audio node.
//
// bus2: The input bus of the destination audio node.
//
// format: If not [NULL], the engine uses this value for the format of the source
// audio node’s output bus. In all cases, the engine matches the format of
// the destination audio node’s input bus to the source audio node’s
// output bus.
//
// # Discussion
//
// Audio nodes have input and output busses ([AVAudioNodeBus]). Use this
// method to establish connections between audio nodes. Connections are always
// one-to-one, never one-to-many or many-to-one.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:fromBus:toBus:format:)
func (a AVAudioEngine) ConnectToFromBusToBusFormat(node1 IAVAudioNode, node2 IAVAudioNode, bus1 AVAudioNodeBus, bus2 AVAudioNodeBus, format IAVAudioFormat) {
	objc.Send[objc.ID](a.ID, objc.Sel("connect:to:fromBus:toBus:format:"), node1, node2, bus1, bus2, format)
}

// Removes all input connections of the node.
//
// node: The audio node with the inputs you want to disconnect.
//
// # Discussion
//
// Connections break on each of the audio node’s input buses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeInput(_:)
func (a AVAudioEngine) DisconnectNodeInput(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectNodeInput:"), node)
}

// Removes the input connection of a node on the specified bus.
//
// node: The audio node with the input to disconnect.
//
// bus: The destination’s input bus to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeInput(_:bus:)
func (a AVAudioEngine) DisconnectNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectNodeInput:bus:"), node, bus)
}

// Removes all output connections of a node.
//
// node: The audio node with the outputs to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeOutput(_:)
func (a AVAudioEngine) DisconnectNodeOutput(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectNodeOutput:"), node)
}

// Removes the output connection of a node on the specified bus.
//
// node: The audio node with the output to disconnect.
//
// bus: The destination’s output bus to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeOutput(_:bus:)
func (a AVAudioEngine) DisconnectNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectNodeOutput:bus:"), node, bus)
}

// Establishes a MIDI connection between two nodes.
//
// sourceNode: The source node.
//
// destinationNode: The destination node.
//
// format: If not [NULL], the engine uses this value for the format of the source
// audio node’s output bus. In all cases, the engine matches the format of
// the destination audio node’s input bus to the source node’s output bus.
//
// tapBlock: If not [NULL], the source node’s event list block calls this on the
// real-time thread. The host can tap the MIDI data of the source node through
// this block.
//
// tapBlock is a [audiotoolbox.AUMIDIEventListBlock].
//
// # Discussion
//
// Use this to establish a MIDI connection between a source node and a
// destination node that has MIDI input capability. This method disconnects
// any existing MIDI connection that involves the destination node. When
// making the MIDI connection, this method overwrites the source node’s
// event list block.
//
// The source node can only be an [AVAudioUnit] node with the type
// [kAudioUnitType_MIDIProcessor]. The destination node types can be
// [kAudioUnitType_MusicDevice], [kAudioUnitType_MusicEffect], or
// [kAudioUnitType_MIDIProcessor].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI(_:to:format:eventListBlock:)-73cd1
// tapBlock is a [audiotoolbox.AUMIDIEventListBlock].
//
// [kAudioUnitType_MIDIProcessor]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MIDIProcessor
// [kAudioUnitType_MusicDevice]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MusicDevice
// [kAudioUnitType_MusicEffect]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MusicEffect
func (a AVAudioEngine) ConnectMIDIToFormatEventListBlock(sourceNode IAVAudioNode, destinationNode IAVAudioNode, format IAVAudioFormat, tapBlock objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("connectMIDI:to:format:eventListBlock:"), sourceNode, destinationNode, format, tapBlock)
}

// Establishes a MIDI connection between a source node and multiple
// destination nodes.
//
// sourceNode: The source node.
//
// destinationNodes: An array of objects that specify the destination nodes.
//
// format: If not [NULL], the engine uses this value for the format of the source
// audio node’s output bus. In all cases, the format of the source node’s
// output bus has to match with the destination node’s output bus format.
//
// tapBlock: If not [NULL], the source node’s event list block calls this on the
// real-time thread. The host can tap the MIDI data of the source node through
// this block.
//
// tapBlock is a [audiotoolbox.AUMIDIEventListBlock].
//
// # Discussion
//
// Use this to establish a MIDI connection between a source node and multiple
// destination nodes that have MIDI input capability. This method disconnects
// any existing MIDI connection that involves the destination node. When
// making the MIDI connection, this method overwrites the source node’s
// event list block.
//
// The source node can only be an [AVAudioUnit] node with the type
// [kAudioUnitType_MIDIProcessor]. The destination node types can be
// [kAudioUnitType_MusicDevice], [kAudioUnitType_MusicEffect], or
// [kAudioUnitType_MIDIProcessor].
//
// MIDI connections made with this method specify a single destination
// connection (one-to-one) or multiple connections (one-to-many), but never
// many-to-one.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI(_:to:format:eventListBlock:)-7qtd5
// tapBlock is a [audiotoolbox.AUMIDIEventListBlock].
//
// [kAudioUnitType_MIDIProcessor]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MIDIProcessor
// [kAudioUnitType_MusicDevice]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MusicDevice
// [kAudioUnitType_MusicEffect]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitType_MusicEffect
func (a AVAudioEngine) ConnectMIDIToNodesFormatEventListBlock(sourceNode IAVAudioNode, destinationNodes []AVAudioNode, format IAVAudioFormat, tapBlock objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("connectMIDI:toNodes:format:eventListBlock:"), sourceNode, objectivec.IObjectSliceToNSArray(destinationNodes), format, tapBlock)
}

// Removes a MIDI connection between two nodes.
//
// sourceNode: The node with the MIDI output to disconnect.
//
// destinationNode: The node with the MIDI input to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDI(_:from:)-1kssy
func (a AVAudioEngine) DisconnectMIDIFrom(sourceNode IAVAudioNode, destinationNode IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectMIDI:from:"), sourceNode, destinationNode)
}

// Removes a MIDI connection between one source node and multiple destination
// nodes.
//
// sourceNode: The node with the MIDI output to disconnect.
//
// destinationNodes: A list of nodes with the MIDI input to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDI(_:from:)-7oaab
func (a AVAudioEngine) DisconnectMIDIFromNodes(sourceNode IAVAudioNode, destinationNodes []AVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectMIDI:fromNodes:"), sourceNode, objectivec.IObjectSliceToNSArray(destinationNodes))
}

// Disconnects all input MIDI connections from a node.
//
// node: The node with the MIDI input to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDIInput(_:)
func (a AVAudioEngine) DisconnectMIDIInput(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectMIDIInput:"), node)
}

// Disconnects all output MIDI connections from a node.
//
// node: The node with the MIDI outputs to disconnect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDIOutput(_:)
func (a AVAudioEngine) DisconnectMIDIOutput(node IAVAudioNode) {
	objc.Send[objc.ID](a.ID, objc.Sel("disconnectMIDIOutput:"), node)
}

// Prepares the audio engine for starting.
//
// # Discussion
//
// This method preallocates many resources the audio engine requires to start.
// Use it to responsively start audio input or output.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/prepare()
func (a AVAudioEngine) Prepare() {
	objc.Send[objc.ID](a.ID, objc.Sel("prepare"))
}

// Starts the audio engine.
//
// # Discussion
//
// This method calls the [Prepare] method if you don’t call it after
// invoking [Stop]. It then starts the audio hardware through the
// [AVAudioInputNode] and [AVAudioOutputNode] instances in the audio engine.
// This method throws an error when:
//
// - There’s a problem in the structure of the graph, such as the input
// can’t route to an output or to a recording tap through converter nodes. -
// An [AVAudioSession] error occurs. - The driver fails to start the hardware.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/start()
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
func (a AVAudioEngine) StartAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("startAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Pauses the audio engine.
//
// # Discussion
//
// This method stops the audio engine and the audio hardware, but doesn’t
// deallocate the resources for the [Prepare] method. When your app doesn’t
// need to play audio, consider pausing or stopping the engine to minimize
// power consumption.
//
// You resume the audio engine by invoking [StartAndReturnError].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/pause()
func (a AVAudioEngine) Pause() {
	objc.Send[objc.ID](a.ID, objc.Sel("pause"))
}

// Stops the audio engine and releases any previously prepared resources.
//
// # Discussion
//
// This method stops the audio engine and the audio hardware, and releases any
// allocated resources for the [Prepare] method. When your app doesn’t need
// to play audio, consider pausing or stopping the engine to minimize power
// consumption.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/stop()
func (a AVAudioEngine) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}

// Resets all audio nodes in the audio engine.
//
// # Discussion
//
// This methods resets all audio nodes in the audio engine. For example, use
// it to silence reverb and delay tails.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/reset()
func (a AVAudioEngine) Reset() {
	objc.Send[objc.ID](a.ID, objc.Sel("reset"))
}

// Sets the engine to operate in manual rendering mode with the render format
// and maximum frame count you specify.
//
// mode: The manual rendering mode to use.
//
// pcmFormat: The format of the output PCM audio data from the engine.
//
// maximumFrameCount: The maximum number of PCM sample frames the engine produces in a single
// render call.
//
// # Discussion
//
// Use this method to configure the engine to render in response to requests
// from the client. You must stop the engine before calling this method. The
// render format must be a PCM format and match the format of the rendering
// buffer.
//
// The source nodes can supply the input data in manual rendering mode. For
// more information, see [AVAudioPlayerNode] and [AVAudioInputNode].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/enableManualRenderingMode(_:format:maximumFrameCount:)
func (a AVAudioEngine) EnableManualRenderingModeFormatMaximumFrameCountError(mode AVAudioEngineManualRenderingMode, pcmFormat IAVAudioFormat, maximumFrameCount AVAudioFrameCount) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("enableManualRenderingMode:format:maximumFrameCount:error:"), mode, pcmFormat, maximumFrameCount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableManualRenderingMode:format:maximumFrameCount:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the engine to render to or from an audio device.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disableManualRenderingMode()
func (a AVAudioEngine) DisableManualRenderingMode() {
	objc.Send[objc.ID](a.ID, objc.Sel("disableManualRenderingMode"))
}

// Makes a render call to the engine operating in the offline manual rendering
// mode.
//
// numberOfFrames: The number of PCM sample frames to render.
//
// buffer: The PCM buffer the engine must render the audio for.
//
// # Return Value
//
// One of the status codes from [AVAudioEngineManualRenderingStatus].
// Irrespective of the returned status code, on exit, the output buffer’s
// [FrameLength] indicates the number of PCM samples the engine renders.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/renderOffline(_:to:)
//
// [AVAudioEngineManualRenderingStatus]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus
func (a AVAudioEngine) RenderOfflineToBufferError(numberOfFrames AVAudioFrameCount, buffer IAVAudioPCMBuffer) (AVAudioEngineManualRenderingStatus, error) {
	var errorPtr objc.ID
	rv := objc.Send[AVAudioEngineManualRenderingStatus](a.ID, objc.Sel("renderOffline:toBuffer:error:"), numberOfFrames, buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Establishes a connection between a source node and multiple destination
// nodes.
//
// sourceNode: The source node.
//
// destNodes: An array of [AVAudioConnectionPoint] objects that specify destination nodes
// and busses.
//
// sourceBus: The output bus on the source node.
//
// format: If not [NULL], the framework uses this value for the format of the source
// audio node’s output bus. In all cases, the framework matches the format
// of the destination audio node’s input bus to the source audio node’s
// output bus.
//
// # Discussion
//
// Connections that use this method are either one-to-one or one-to-many.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:fromBus:format:)
func (a AVAudioEngine) ConnectToConnectionPointsFromBusFormat(sourceNode IAVAudioNode, destNodes []AVAudioConnectionPoint, sourceBus AVAudioNodeBus, format IAVAudioFormat) {
	objc.Send[objc.ID](a.ID, objc.Sel("connect:toConnectionPoints:fromBus:format:"), sourceNode, objectivec.IObjectSliceToNSArray(destNodes), sourceBus, format)
}

// Returns connection information about a node’s input bus.
//
// node: The node with the input connection you’re querying.
//
// bus: The node’s input bus for the connection you’re querying.
//
// # Return Value
//
// An [AVAudioConnectionPoint] object with connection information on the
// node’s input bus.
//
// # Discussion
//
// Connections are always one-to-one or one-to-many. This method returns `nil`
// if there’s no connection on the node’s specified input bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputConnectionPoint(for:inputBus:)
func (a AVAudioEngine) InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus) IAVAudioConnectionPoint {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputConnectionPointForNode:inputBus:"), node, bus)
	return AVAudioConnectionPointFromID(rv)
}

// Returns connection information about a node’s output bus.
//
// node: The node with the output connections you’re querying.
//
// bus: The node’s output bus for connections you’re querying.
//
// # Return Value
//
// An array of [AVAudioConnectionPoint] objects with connection information on
// the node’s output bus.
//
// # Discussion
//
// Connections are always one-to-one or one-to-many. This method returns an
// empty array if there are no connections on the node’s specified output
// bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputConnectionPoints(for:outputBus:)
func (a AVAudioEngine) OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus) []AVAudioConnectionPoint {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("outputConnectionPointsForNode:outputBus:"), node, bus)
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioConnectionPoint {
		return AVAudioConnectionPointFromID(id)
	})
}

// A read-only set that contains the nodes you attach to the audio engine.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/attachedNodes
func (a AVAudioEngine) AttachedNodes() foundation.INSSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attachedNodes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The audio engine’s singleton input audio node.
//
// # Discussion
//
// The framework performs audio input through an input node. The audio engine
// creates a singleton on demand when first accessing this variable. To
// receive input, connect another node from the output of the input node, or
// create a recording tap on it.
//
// When the engine renders to and from an audio device, the [AVAudioSession]
// category and the availability of hardware determines whether an app
// performs input (for example, input hardware isn’t available in tvOS).
// Check the input node’s input format (specifically, the hardware format)
// for a nonzero sample rate and channel count to see if input is in an
// enabled state.
//
// Trying to perform input through the input node when it isn’t available or
// in an enabled state causes the engine to throw an error (when possible) or
// an exception.
//
// In manual rendering mode, the input node can synchronously supply data to
// the engine while it’s rendering. For more information, see
// [SetManualRenderingInputPCMFormatInputBlock].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputNode
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
func (a AVAudioEngine) InputNode() IAVAudioInputNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputNode"))
	return AVAudioInputNodeFromID(objc.ID(rv))
}

// The audio engine’s singleton output audio node.
//
// # Discussion
//
// The framework performs audio output through an output node. The audio
// engine creates a singleton on demand when first accessing this variable.
// Connect another node to the input of the output node, or get a mixer using
// the [MainMixerNode] property.
//
// When the engine renders to and from an audio device, the [AVAudioSession]
// category and the availability of hardware determines whether an app
// performs output. Check the output node’s output format (specifically, the
// hardware format) for a nonzero sample rate and channel count to see if
// output is in an enabled state.
//
// Trying to perform output through the output node when it isn’t available
// or in an enabled state causes the engine to throw an error (when possible)
// or an exception.
//
// In manual rendering mode, the output node’s format determines the render
// format of the engine. For more information about changing it, see
// [EnableManualRenderingModeFormatMaximumFrameCountError].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputNode
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
func (a AVAudioEngine) OutputNode() IAVAudioOutputNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputNode"))
	return AVAudioOutputNodeFromID(objc.ID(rv))
}

// The audio engine’s optional singleton main mixer node.
//
// # Discussion
//
// The audio engine constructs a singleton main mixer and connects it to the
// [OutputNode] when first accessing this property. You can then connect
// additional audio nodes to the mixer.
//
// If the client never sets the connection format between the `mainMixerNode`
// and the `outputNode`, the engine always updates the format to track the
// format of the `outputNode` on startup or restart, even after an
// [AVAudioEngineConfigurationChangeNotification]. Otherwise, it’s the
// client’s responsibility to update the connection format after an
// [AVAudioEngineConfigurationChangeNotification].
//
// By default, the mixer’s output format (sample rate and channel count)
// tracks the format of the output node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/mainMixerNode
//
// [AVAudioEngineConfigurationChangeNotification]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineConfigurationChangeNotification
func (a AVAudioEngine) MainMixerNode() IAVAudioMixerNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mainMixerNode"))
	return AVAudioMixerNodeFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the audio engine is running.
//
// # Discussion
//
// The value is true if the audio engine is in a running state; otherwise,
// false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isRunning
func (a AVAudioEngine) Running() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRunning"))
	return rv
}

// The music sequence instance that you attach to the audio engine, if any.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a AVAudioEngine) MusicSequence() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("musicSequence"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioEngine) SetMusicSequence(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setMusicSequence:"), value)
}

// The block that renders the engine when operating in manual rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingBlock
func (a AVAudioEngine) ManualRenderingBlock() AVAudioEngineManualRenderingBlock {
	rv := objc.Send[AVAudioEngineManualRenderingBlock](a.ID, objc.Sel("manualRenderingBlock"))
	return AVAudioEngineManualRenderingBlock(rv)
}

// The render format of the engine in manual rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingFormat
func (a AVAudioEngine) ManualRenderingFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manualRenderingFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}

// The maximum number of PCM sample frames the engine produces in any single
// render call in manual rendering mode.
//
// # Discussion
//
// If you get this property when the engine isn’t in manual rendering mode,
// it returns zero.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingMaximumFrameCount
func (a AVAudioEngine) ManualRenderingMaximumFrameCount() AVAudioFrameCount {
	rv := objc.Send[AVAudioFrameCount](a.ID, objc.Sel("manualRenderingMaximumFrameCount"))
	return AVAudioFrameCount(rv)
}

// The manual rendering mode configured on the engine.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingMode
func (a AVAudioEngine) ManualRenderingMode() AVAudioEngineManualRenderingMode {
	rv := objc.Send[AVAudioEngineManualRenderingMode](a.ID, objc.Sel("manualRenderingMode"))
	return AVAudioEngineManualRenderingMode(rv)
}

// An indication of where the engine is on its render timeline in manual
// rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingSampleTime
func (a AVAudioEngine) ManualRenderingSampleTime() AVAudioFramePosition {
	rv := objc.Send[AVAudioFramePosition](a.ID, objc.Sel("manualRenderingSampleTime"))
	return AVAudioFramePosition(rv)
}

// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// # Discussion
//
// If autoshutdown is in an enabled state, the engine can start and stop the
// audio hardware dynamically to conserve power. In watchOS, autoshutdown is
// always in an enabled state. For other platforms, it’s in a disabled state
// by default.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isAutoShutdownEnabled
func (a AVAudioEngine) AutoShutdownEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAutoShutdownEnabled"))
	return rv
}
func (a AVAudioEngine) SetAutoShutdownEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutoShutdownEnabled:"), value)
}

// A Boolean value that indicates whether the engine is operating in manual
// rendering mode.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isInManualRenderingMode
func (a AVAudioEngine) IsInManualRenderingMode() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInManualRenderingMode"))
	return rv
}
