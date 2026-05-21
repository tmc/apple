// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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

// # Methods
//
//   - [AVAudioOutputNode.ManualRenderingFormat]
//   - [AVAudioOutputNode.ManualRenderingMaximumFrameCount]
//   - [AVAudioOutputNode.SetManualRenderingPCMFormatMaximumFrameCount]
type AVAudioOutputNode struct {
	AVAudioIONode
}

// AVAudioOutputNodeFromID constructs a [AVAudioOutputNode] from an objc.ID.
func AVAudioOutputNodeFromID(id objc.ID) AVAudioOutputNode {
	return AVAudioOutputNode{AVAudioIONode: AVAudioIONodeFromID(id)}
}

// Ensure AVAudioOutputNode implements IAVAudioOutputNode.
var _ IAVAudioOutputNode = AVAudioOutputNode{}

// An interface definition for the [AVAudioOutputNode] class.
//
// # Methods
//
//   - [IAVAudioOutputNode.ManualRenderingFormat]
//   - [IAVAudioOutputNode.ManualRenderingMaximumFrameCount]
//   - [IAVAudioOutputNode.SetManualRenderingPCMFormatMaximumFrameCount]
type IAVAudioOutputNode interface {
	IAVAudioIONode

	// Topic: Methods

	ManualRenderingFormat() objectivec.IObject
	ManualRenderingMaximumFrameCount() uint32
	SetManualRenderingPCMFormatMaximumFrameCount(pCMFormat objectivec.IObject, count uint32) bool
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

func NewAudioOutputNodeWithIOUnitIsInput(iOUnit unsafe.Pointer, input bool) AVAudioOutputNode {
	instance := getAVAudioOutputNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOUnit:isInput:"), iOUnit, input)
	return AVAudioOutputNodeFromID(rv)
}

func NewAudioOutputNodeWithImpl(impl unsafe.Pointer) AVAudioOutputNode {
	instance := getAVAudioOutputNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioOutputNodeFromID(rv)
}

func (a AVAudioOutputNode) ManualRenderingFormat() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manualRenderingFormat"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioOutputNode) ManualRenderingMaximumFrameCount() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("manualRenderingMaximumFrameCount"))
	return rv
}
func (a AVAudioOutputNode) SetManualRenderingPCMFormatMaximumFrameCount(pCMFormat objectivec.IObject, count uint32) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setManualRenderingPCMFormat:maximumFrameCount:"), pCMFormat, count)
	return rv
}
