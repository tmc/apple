// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVAudioSourceNode] class.
var (
	_AVAudioSourceNodeClass     AVAudioSourceNodeClass
	_AVAudioSourceNodeClassOnce sync.Once
)

func getAVAudioSourceNodeClass() AVAudioSourceNodeClass {
	_AVAudioSourceNodeClassOnce.Do(func() {
		_AVAudioSourceNodeClass = AVAudioSourceNodeClass{class: objc.GetClass("AVAudioSourceNode")}
	})
	return _AVAudioSourceNodeClass
}

// GetAVAudioSourceNodeClass returns the class object for AVAudioSourceNode.
func GetAVAudioSourceNodeClass() AVAudioSourceNodeClass {
	return getAVAudioSourceNodeClass()
}

type AVAudioSourceNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSourceNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSourceNodeClass) Alloc() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioSourceNode.DebugDescription]
//   - [AVAudioSourceNode.Description]
//   - [AVAudioSourceNode.Hash]
//   - [AVAudioSourceNode.Superclass]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode
type AVAudioSourceNode struct {
	AVAudioNode
}

// AVAudioSourceNodeFromID constructs a [AVAudioSourceNode] from an objc.ID.
func AVAudioSourceNodeFromID(id objc.ID) AVAudioSourceNode {
	return AVAudioSourceNode{AVAudioNode: AVAudioNodeFromID(id)}
}
// Ensure AVAudioSourceNode implements IAVAudioSourceNode.
var _ IAVAudioSourceNode = AVAudioSourceNode{}

// An interface definition for the [AVAudioSourceNode] class.
//
// # Methods
//
//   - [IAVAudioSourceNode.DebugDescription]
//   - [IAVAudioSourceNode.Description]
//   - [IAVAudioSourceNode.Hash]
//   - [IAVAudioSourceNode.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode
type IAVAudioSourceNode interface {
	IAVAudioNode

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a AVAudioSourceNode) Init() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSourceNode) Autorelease() AVAudioSourceNode {
	rv := objc.Send[AVAudioSourceNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSourceNode creates a new AVAudioSourceNode instance.
func NewAVAudioSourceNode() AVAudioSourceNode {
	class := getAVAudioSourceNodeClass()
	rv := objc.Send[AVAudioSourceNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioSourceNodeWithImpl(impl unsafe.Pointer) AVAudioSourceNode {
	instance := getAVAudioSourceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioSourceNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/pullInputBlockFromRenderBlock:
func (_AVAudioSourceNodeClass AVAudioSourceNodeClass) PullInputBlockFromRenderBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](objc.ID(_AVAudioSourceNodeClass.class), objc.Sel("pullInputBlockFromRenderBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/debugDescription
func (a AVAudioSourceNode) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/description
func (a AVAudioSourceNode) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/hash
func (a AVAudioSourceNode) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/superclass
func (a AVAudioSourceNode) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}

// PullInputBlockFromRenderBlockSync is a synchronous wrapper around [AVAudioSourceNode.PullInputBlockFromRenderBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioSourceNodeClass) PullInputBlockFromRenderBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	ac.PullInputBlockFromRenderBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

