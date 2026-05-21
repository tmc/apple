// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioMixerNode] class.
var (
	_AVAudioMixerNodeClass     AVAudioMixerNodeClass
	_AVAudioMixerNodeClassOnce sync.Once
)

func getAVAudioMixerNodeClass() AVAudioMixerNodeClass {
	_AVAudioMixerNodeClassOnce.Do(func() {
		_AVAudioMixerNodeClass = AVAudioMixerNodeClass{class: objc.GetClass("AVAudioMixerNode")}
	})
	return _AVAudioMixerNodeClass
}

// GetAVAudioMixerNodeClass returns the class object for AVAudioMixerNode.
func GetAVAudioMixerNodeClass() AVAudioMixerNodeClass {
	return getAVAudioMixerNodeClass()
}

type AVAudioMixerNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioMixerNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioMixerNodeClass) Alloc() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioMixerNode.InputConnected]
//   - [AVAudioMixerNode.SetInputPanBus]
//   - [AVAudioMixerNode.SetInputVolumeBus]
//   - [AVAudioMixerNode.DebugDescription]
//   - [AVAudioMixerNode.Description]
//   - [AVAudioMixerNode.Hash]
//   - [AVAudioMixerNode.Superclass]
type AVAudioMixerNode struct {
	AVAudioNode
}

// AVAudioMixerNodeFromID constructs a [AVAudioMixerNode] from an objc.ID.
func AVAudioMixerNodeFromID(id objc.ID) AVAudioMixerNode {
	return AVAudioMixerNode{AVAudioNode: AVAudioNodeFromID(id)}
}

// Ensure AVAudioMixerNode implements IAVAudioMixerNode.
var _ IAVAudioMixerNode = AVAudioMixerNode{}

// An interface definition for the [AVAudioMixerNode] class.
//
// # Methods
//
//   - [IAVAudioMixerNode.InputConnected]
//   - [IAVAudioMixerNode.SetInputPanBus]
//   - [IAVAudioMixerNode.SetInputVolumeBus]
//   - [IAVAudioMixerNode.DebugDescription]
//   - [IAVAudioMixerNode.Description]
//   - [IAVAudioMixerNode.Hash]
//   - [IAVAudioMixerNode.Superclass]
type IAVAudioMixerNode interface {
	IAVAudioNode

	// Topic: Methods

	InputConnected(connected uint64)
	SetInputPanBus(pan float32, bus uint64)
	SetInputVolumeBus(volume float32, bus uint64)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (a AVAudioMixerNode) Init() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioMixerNode) Autorelease() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMixerNode creates a new AVAudioMixerNode instance.
func NewAVAudioMixerNode() AVAudioMixerNode {
	class := getAVAudioMixerNodeClass()
	rv := objc.Send[AVAudioMixerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioMixerNodeWithImpl(impl unsafe.Pointer) AVAudioMixerNode {
	instance := getAVAudioMixerNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioMixerNodeFromID(rv)
}

func (a AVAudioMixerNode) InputConnected(connected uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("inputConnected:"), connected)
}
func (a AVAudioMixerNode) SetInputPanBus(pan float32, bus uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setInputPan:bus:"), pan, bus)
}
func (a AVAudioMixerNode) SetInputVolumeBus(volume float32, bus uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setInputVolume:bus:"), volume, bus)
}

func (a AVAudioMixerNode) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioMixerNode) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioMixerNode) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}
func (a AVAudioMixerNode) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](a.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
