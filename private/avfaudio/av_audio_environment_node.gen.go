// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

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

// # Methods
//
//   - [AVAudioEnvironmentNode.DebugDescription]
//   - [AVAudioEnvironmentNode.Description]
//   - [AVAudioEnvironmentNode.Hash]
//   - [AVAudioEnvironmentNode.ListenerHeadTrackingEnabled]
//   - [AVAudioEnvironmentNode.SetListenerHeadTrackingEnabled]
//   - [AVAudioEnvironmentNode.Superclass]
type AVAudioEnvironmentNode struct {
	AVAudioNode
}

// AVAudioEnvironmentNodeFromID constructs a [AVAudioEnvironmentNode] from an objc.ID.
func AVAudioEnvironmentNodeFromID(id objc.ID) AVAudioEnvironmentNode {
	return AVAudioEnvironmentNode{AVAudioNode: AVAudioNodeFromID(id)}
}

// Ensure AVAudioEnvironmentNode implements IAVAudioEnvironmentNode.
var _ IAVAudioEnvironmentNode = AVAudioEnvironmentNode{}

// An interface definition for the [AVAudioEnvironmentNode] class.
//
// # Methods
//
//   - [IAVAudioEnvironmentNode.DebugDescription]
//   - [IAVAudioEnvironmentNode.Description]
//   - [IAVAudioEnvironmentNode.Hash]
//   - [IAVAudioEnvironmentNode.ListenerHeadTrackingEnabled]
//   - [IAVAudioEnvironmentNode.SetListenerHeadTrackingEnabled]
//   - [IAVAudioEnvironmentNode.Superclass]
type IAVAudioEnvironmentNode interface {
	IAVAudioNode

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	ListenerHeadTrackingEnabled() bool
	SetListenerHeadTrackingEnabled(value bool)
	Superclass() objectivec.Class
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

func NewAudioEnvironmentNodeWithImpl(impl unsafe.Pointer) AVAudioEnvironmentNode {
	instance := getAVAudioEnvironmentNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioEnvironmentNodeFromID(rv)
}

func (a AVAudioEnvironmentNode) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioEnvironmentNode) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioEnvironmentNode) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}
func (a AVAudioEnvironmentNode) ListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("listenerHeadTrackingEnabled"))
	return rv
}
func (a AVAudioEnvironmentNode) SetListenerHeadTrackingEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setListenerHeadTrackingEnabled:"), value)
}
func (a AVAudioEnvironmentNode) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](a.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
