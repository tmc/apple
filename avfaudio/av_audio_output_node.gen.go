// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioOutputNode] class.
var (
	_AVAudioOutputNodeClass     AVAudioOutputNodeClass
	_AVAudioOutputNodeClassOnce sync.Once
)

func getAVAudioOutputNodeClass() AVAudioOutputNodeClass {
	_AVAudioOutputNodeClassOnce.Do(func() {
		_AVAudioOutputNodeClass = AVAudioOutputNodeClass{class: objc.GetClass("AVAudioOutputNode")}
	})
	return _AVAudioOutputNodeClass
}

// GetAVAudioOutputNodeClass returns the class object for AVAudioOutputNode.
func GetAVAudioOutputNodeClass() AVAudioOutputNodeClass {
	return getAVAudioOutputNodeClass()
}

type AVAudioOutputNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioOutputNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioOutputNodeClass) Alloc() AVAudioOutputNode {
	rv := objc.Send[AVAudioOutputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that connects to the system’s audio output.
//
// # Overview
//
// This node connects to the system’s audio output when rendering to or from
// an audio device. This node performs output in response to client’s
// requests when the engine is in manual rendering mode.
//
// This audio node has one element. The format of the output scope reflects:
//
// - The audio hardware sample rate and channel count when it connects to the
// hardware. - The engine’s manual rendering mode output format (see
// [AVAudioOutputNode.ManualRenderingFormat]).
//
// The format of the input scope is initially the same as that of the output,
// but you may set it to a different format, in which case the audio node
// converts.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode
type AVAudioOutputNode struct {
	AVAudioIONode
}

// AVAudioOutputNodeFromID constructs a [AVAudioOutputNode] from an objc.ID.
//
// An object that connects to the system’s audio output.
func AVAudioOutputNodeFromID(id objc.ID) AVAudioOutputNode {
	return AVAudioOutputNode{AVAudioIONode: AVAudioIONodeFromID(id)}
}

// NOTE: AVAudioOutputNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioOutputNode] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode
type IAVAudioOutputNode interface {
	IAVAudioIONode

	// The render format of the engine in manual rendering mode.
	ManualRenderingFormat() IAVAudioFormat
	SetManualRenderingFormat(value IAVAudioFormat)
}

// Init initializes the instance.
func (a AVAudioOutputNode) Init() AVAudioOutputNode {
	rv := objc.Send[AVAudioOutputNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioOutputNode) Autorelease() AVAudioOutputNode {
	rv := objc.Send[AVAudioOutputNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioOutputNode creates a new AVAudioOutputNode instance.
func NewAVAudioOutputNode() AVAudioOutputNode {
	class := getAVAudioOutputNodeClass()
	rv := objc.Send[AVAudioOutputNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The render format of the engine in manual rendering mode.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingformat
func (a AVAudioOutputNode) ManualRenderingFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manualRenderingFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}
func (a AVAudioOutputNode) SetManualRenderingFormat(value IAVAudioFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setManualRenderingFormat:"), value)
}
