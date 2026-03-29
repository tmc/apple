// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
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

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode
type AVAudioSinkNode struct {
	AVAudioNode
}

// AVAudioSinkNodeFromID constructs a [AVAudioSinkNode] from an objc.ID.
func AVAudioSinkNodeFromID(id objc.ID) AVAudioSinkNode {
	return AVAudioSinkNode{AVAudioNode: AVAudioNodeFromID(id)}
}
// Ensure AVAudioSinkNode implements IAVAudioSinkNode.
var _ IAVAudioSinkNode = AVAudioSinkNode{}

// An interface definition for the [AVAudioSinkNode] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode
type IAVAudioSinkNode interface {
	IAVAudioNode
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

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioSinkNodeWithImpl(impl unsafe.Pointer) AVAudioSinkNode {
	instance := getAVAudioSinkNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioSinkNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode/pullInputBlockFromReceiverBlock:
func (_AVAudioSinkNodeClass AVAudioSinkNodeClass) PullInputBlockFromReceiverBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](objc.ID(_AVAudioSinkNodeClass.class), objc.Sel("pullInputBlockFromReceiverBlock:"), _block0)
}

// PullInputBlockFromReceiverBlockSync is a synchronous wrapper around [AVAudioSinkNode.PullInputBlockFromReceiverBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioSinkNodeClass) PullInputBlockFromReceiverBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	ac.PullInputBlockFromReceiverBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

