// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioSinkNode] class.
var (
	_AVAudioSinkNodeClass     AVAudioSinkNodeClass
	_AVAudioSinkNodeClassOnce sync.Once
)

func getAVAudioSinkNodeClass() AVAudioSinkNodeClass {
	_AVAudioSinkNodeClassOnce.Do(func() {
		_AVAudioSinkNodeClass = AVAudioSinkNodeClass{class: objc.GetClass("AVAudioSinkNode")}
	})
	return _AVAudioSinkNodeClass
}

// GetAVAudioSinkNodeClass returns the class object for AVAudioSinkNode.
func GetAVAudioSinkNodeClass() AVAudioSinkNodeClass {
	return getAVAudioSinkNodeClass()
}

type AVAudioSinkNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSinkNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSinkNodeClass) Alloc() AVAudioSinkNode {
	rv := objc.Send[AVAudioSinkNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that receives audio data.
//
// # Overview
//
// You use an [AVAudioSinkNode] to receive audio data through
// [AVAudioSinkNodeReceiverBlock]. You only use it in the input chain.
//
// An audio sink node doesn’t support format conversion. When connecting,
// use the output format of the input for the format for the connection. The
// format should match the hardware input sample rate.
//
// The voice processing I/O unit is an exception to the above because it
// supports sample rate conversion. The input scope format (hardware format)
// and output scope format (client format) of the input node can differ in
// that case.
//
// An audio sink node doesn’t support manual rendering mode, and doesn’t
// have an output bus, so you can’t install a tap on it.
//
// # Creating an Audio Sink Node
//
//   - [AVAudioSinkNode.InitWithReceiverBlock]: Creates an audio sink node with a block that receives audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode
type AVAudioSinkNode struct {
	AVAudioNode
}

// AVAudioSinkNodeFromID constructs a [AVAudioSinkNode] from an objc.ID.
//
// An object that receives audio data.
func AVAudioSinkNodeFromID(id objc.ID) AVAudioSinkNode {
	return AVAudioSinkNode{AVAudioNode: AVAudioNodeFromID(id)}
}

// NOTE: AVAudioSinkNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioSinkNode] class.
//
// # Creating an Audio Sink Node
//
//   - [IAVAudioSinkNode.InitWithReceiverBlock]: Creates an audio sink node with a block that receives audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode
type IAVAudioSinkNode interface {
	IAVAudioNode

	// Topic: Creating an Audio Sink Node

	// Creates an audio sink node with a block that receives audio data.
	InitWithReceiverBlock(block AVAudioSinkNodeReceiverBlock) AVAudioSinkNode
}

// Init initializes the instance.
func (a AVAudioSinkNode) Init() AVAudioSinkNode {
	rv := objc.Send[AVAudioSinkNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSinkNode) Autorelease() AVAudioSinkNode {
	rv := objc.Send[AVAudioSinkNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSinkNode creates a new AVAudioSinkNode instance.
func NewAVAudioSinkNode() AVAudioSinkNode {
	class := getAVAudioSinkNodeClass()
	rv := objc.Send[AVAudioSinkNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio sink node with a block that receives audio data.
//
// block: The block that receives audio data from the input.
//
// # Discussion
//
// When connecting the audio sink node to another node, the system uses the
// connection format to set the audio format for the input bus.
//
// The system calls the block on the real-time thread when receiving input
// data. Avoid making blocking calls within the block.
//
// When receiving data, the system sets the audio format using the node’s
// input format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode/init(receiverBlock:)
func NewAudioSinkNodeWithReceiverBlock(block AVAudioSinkNodeReceiverBlock) AVAudioSinkNode {
	instance := getAVAudioSinkNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithReceiverBlock:"), block)
	return AVAudioSinkNodeFromID(rv)
}

// Creates an audio sink node with a block that receives audio data.
//
// block: The block that receives audio data from the input.
//
// # Discussion
//
// When connecting the audio sink node to another node, the system uses the
// connection format to set the audio format for the input bus.
//
// The system calls the block on the real-time thread when receiving input
// data. Avoid making blocking calls within the block.
//
// When receiving data, the system sets the audio format using the node’s
// input format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode/init(receiverBlock:)
func (a AVAudioSinkNode) InitWithReceiverBlock(block AVAudioSinkNodeReceiverBlock) AVAudioSinkNode {
	rv := objc.Send[AVAudioSinkNode](a.ID, objc.Sel("initWithReceiverBlock:"), block)
	return rv
}
