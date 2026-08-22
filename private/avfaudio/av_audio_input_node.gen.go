// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	rv := objc.SendIfResponds[AVAudioInputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioInputNode.DebugDescription]
//   - [AVAudioInputNode.Description]
//   - [AVAudioInputNode.Hash]
//   - [AVAudioInputNode.Superclass]
//   - [AVAudioInputNode.VoiceProcessingAGCEnabled]
//   - [AVAudioInputNode.SetVoiceProcessingAGCEnabled]
//   - [AVAudioInputNode.VoiceProcessingBypassed]
//   - [AVAudioInputNode.SetVoiceProcessingBypassed]
//   - [AVAudioInputNode.VoiceProcessingInputMuted]
//   - [AVAudioInputNode.SetVoiceProcessingInputMuted]
type AVAudioInputNode struct {
	AVAudioIONode
}

// AVAudioInputNodeFromID constructs a [AVAudioInputNode] from an objc.ID.
func AVAudioInputNodeFromID(id objc.ID) AVAudioInputNode {
	return AVAudioInputNode{AVAudioIONode: AVAudioIONodeFromID(id)}
}

// Ensure AVAudioInputNode implements IAVAudioInputNode.
var _ IAVAudioInputNode = AVAudioInputNode{}

// An interface definition for the [AVAudioInputNode] class.
//
// # Methods
//
//   - [IAVAudioInputNode.DebugDescription]
//   - [IAVAudioInputNode.Description]
//   - [IAVAudioInputNode.Hash]
//   - [IAVAudioInputNode.Superclass]
//   - [IAVAudioInputNode.VoiceProcessingAGCEnabled]
//   - [IAVAudioInputNode.SetVoiceProcessingAGCEnabled]
//   - [IAVAudioInputNode.VoiceProcessingBypassed]
//   - [IAVAudioInputNode.SetVoiceProcessingBypassed]
//   - [IAVAudioInputNode.VoiceProcessingInputMuted]
//   - [IAVAudioInputNode.SetVoiceProcessingInputMuted]
type IAVAudioInputNode interface {
	IAVAudioIONode

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
	VoiceProcessingAGCEnabled() bool
	SetVoiceProcessingAGCEnabled(value bool)
	VoiceProcessingBypassed() bool
	SetVoiceProcessingBypassed(value bool)
	VoiceProcessingInputMuted() bool
	SetVoiceProcessingInputMuted(value bool)
}

// Init initializes the instance.
func (a AVAudioInputNode) Init() AVAudioInputNode {
	rv := objc.SendIfResponds[AVAudioInputNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioInputNode) Autorelease() AVAudioInputNode {
	rv := objc.SendIfResponds[AVAudioInputNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioInputNode creates a new AVAudioInputNode instance.
func NewAVAudioInputNode() AVAudioInputNode {
	class := getAVAudioInputNodeClass()
	rv := objc.SendIfResponds[AVAudioInputNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioInputNodeWithIOUnitIsInput(iOUnit unsafe.Pointer, input bool) AVAudioInputNode {
	instance := getAVAudioInputNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIOUnit:isInput:"), iOUnit, input)
	return AVAudioInputNodeFromID(rv)
}

func NewAudioInputNodeWithImpl(impl unsafe.Pointer) AVAudioInputNode {
	instance := getAVAudioInputNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioInputNodeFromID(rv)
}

func (a AVAudioInputNode) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioInputNode) Description() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioInputNode) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("hash"))
	return rv
}
func (a AVAudioInputNode) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](a.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (a AVAudioInputNode) VoiceProcessingAGCEnabled() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("voiceProcessingAGCEnabled"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingAGCEnabled(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setVoiceProcessingAGCEnabled:"), value)
}
func (a AVAudioInputNode) VoiceProcessingBypassed() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("voiceProcessingBypassed"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingBypassed(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setVoiceProcessingBypassed:"), value)
}
func (a AVAudioInputNode) VoiceProcessingInputMuted() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("voiceProcessingInputMuted"))
	return rv
}
func (a AVAudioInputNode) SetVoiceProcessingInputMuted(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setVoiceProcessingInputMuted:"), value)
}
