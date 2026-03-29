// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioIONode] class.
var (
	_AVAudioIONodeClass     AVAudioIONodeClass
	_AVAudioIONodeClassOnce sync.Once
)

func getAVAudioIONodeClass() AVAudioIONodeClass {
	_AVAudioIONodeClassOnce.Do(func() {
		_AVAudioIONodeClass = AVAudioIONodeClass{class: objc.GetClass("AVAudioIONode")}
	})
	return _AVAudioIONodeClass
}

// GetAVAudioIONodeClass returns the class object for AVAudioIONode.
func GetAVAudioIONodeClass() AVAudioIONodeClass {
	return getAVAudioIONodeClass()
}

type AVAudioIONodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioIONodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioIONodeClass) Alloc() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioIONode.EnableManualRenderingModeIsInput]
//   - [AVAudioIONode.EnableRealtimeRenderingModeWithIOUnitIsInputForceIOUnitReset]
//   - [AVAudioIONode.IsInManualRenderingMode]
//   - [AVAudioIONode.ManualRenderingMode]
//   - [AVAudioIONode.InitWithIOUnitIsInput]
//   - [AVAudioIONode.VoiceProcessingEnabled]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
type AVAudioIONode struct {
	AVAudioNode
}

// AVAudioIONodeFromID constructs a [AVAudioIONode] from an objc.ID.
func AVAudioIONodeFromID(id objc.ID) AVAudioIONode {
	return AVAudioIONode{AVAudioNode: AVAudioNodeFromID(id)}
}
// Ensure AVAudioIONode implements IAVAudioIONode.
var _ IAVAudioIONode = AVAudioIONode{}

// An interface definition for the [AVAudioIONode] class.
//
// # Methods
//
//   - [IAVAudioIONode.EnableManualRenderingModeIsInput]
//   - [IAVAudioIONode.EnableRealtimeRenderingModeWithIOUnitIsInputForceIOUnitReset]
//   - [IAVAudioIONode.IsInManualRenderingMode]
//   - [IAVAudioIONode.ManualRenderingMode]
//   - [IAVAudioIONode.InitWithIOUnitIsInput]
//   - [IAVAudioIONode.VoiceProcessingEnabled]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
type IAVAudioIONode interface {
	IAVAudioNode

	// Topic: Methods

	EnableManualRenderingModeIsInput(mode int64, input bool) bool
	EnableRealtimeRenderingModeWithIOUnitIsInputForceIOUnitReset(iOUnit unsafe.Pointer, input bool, reset bool) bool
	IsInManualRenderingMode() bool
	ManualRenderingMode() int64
	InitWithIOUnitIsInput(iOUnit unsafe.Pointer, input bool) AVAudioIONode
	VoiceProcessingEnabled() bool
}

// Init initializes the instance.
func (a AVAudioIONode) Init() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioIONode) Autorelease() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioIONode creates a new AVAudioIONode instance.
func NewAVAudioIONode() AVAudioIONode {
	class := getAVAudioIONodeClass()
	rv := objc.Send[AVAudioIONode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/initWithIOUnit:isInput:
func NewAudioIONodeWithIOUnitIsInput(iOUnit unsafe.Pointer, input bool) AVAudioIONode {
	instance := getAVAudioIONodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOUnit:isInput:"), iOUnit, input)
	return AVAudioIONodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioIONodeWithImpl(impl unsafe.Pointer) AVAudioIONode {
	instance := getAVAudioIONodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioIONodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/enableManualRenderingMode:isInput:
func (a AVAudioIONode) EnableManualRenderingModeIsInput(mode int64, input bool) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("enableManualRenderingMode:isInput:"), mode, input)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/enableRealtimeRenderingModeWithIOUnit:isInput:forceIOUnitReset:
func (a AVAudioIONode) EnableRealtimeRenderingModeWithIOUnitIsInputForceIOUnitReset(iOUnit unsafe.Pointer, input bool, reset bool) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("enableRealtimeRenderingModeWithIOUnit:isInput:forceIOUnitReset:"), iOUnit, input, reset)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/isInManualRenderingMode
func (a AVAudioIONode) IsInManualRenderingMode() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInManualRenderingMode"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/manualRenderingMode
func (a AVAudioIONode) ManualRenderingMode() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("manualRenderingMode"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/initWithIOUnit:isInput:
func (a AVAudioIONode) InitWithIOUnitIsInput(iOUnit unsafe.Pointer, input bool) AVAudioIONode {
	rv := objc.Send[AVAudioIONode](a.ID, objc.Sel("initWithIOUnit:isInput:"), iOUnit, input)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/voiceProcessingEnabled
func (a AVAudioIONode) VoiceProcessingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("voiceProcessingEnabled"))
	return rv
}

