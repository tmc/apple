// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[AVAudioSinkNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioSinkNode.InitWithRealtimeSafeReceiverBlock]
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
// # Methods
//
//   - [IAVAudioSinkNode.InitWithRealtimeSafeReceiverBlock]
type IAVAudioSinkNode interface {
	IAVAudioNode

	// Topic: Methods

	InitWithRealtimeSafeReceiverBlock(block VoidHandler) AVAudioSinkNode
}

// Init initializes the instance.
func (a AVAudioSinkNode) Init() AVAudioSinkNode {
	rv := objc.SendIfResponds[AVAudioSinkNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSinkNode) Autorelease() AVAudioSinkNode {
	rv := objc.SendIfResponds[AVAudioSinkNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSinkNode creates a new AVAudioSinkNode instance.
func NewAVAudioSinkNode() AVAudioSinkNode {
	class := getAVAudioSinkNodeClass()
	rv := objc.SendIfResponds[AVAudioSinkNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAVAudioSinkNodeWithImpl(impl unsafe.Pointer) AVAudioSinkNode {
	instance := getAVAudioSinkNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioSinkNodeFromID(rv)
}

var _avaudiosinknode_initwithrealtimesafereceiverblock_p0_key byte

func (a AVAudioSinkNode) InitWithRealtimeSafeReceiverBlock(block VoidHandler) AVAudioSinkNode {
	_block0, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[AVAudioSinkNode](a.ID, objc.Sel("initWithRealtimeSafeReceiverBlock:"), _block0)
	return rv
}

func (_AVAudioSinkNodeClass AVAudioSinkNodeClass) PullInputBlockFromRealtimeSafeReceiverBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](objc.ID(_AVAudioSinkNodeClass.class), objc.Sel("pullInputBlockFromRealtimeSafeReceiverBlock:"), _block0)
}
func (_AVAudioSinkNodeClass AVAudioSinkNodeClass) PullInputBlockFromReceiverBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](objc.ID(_AVAudioSinkNodeClass.class), objc.Sel("pullInputBlockFromReceiverBlock:"), _block0)
}

// InitWithRealtimeSafeReceiverBlockSync is a synchronous wrapper around [AVAudioSinkNode.InitWithRealtimeSafeReceiverBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioSinkNode) InitWithRealtimeSafeReceiverBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.InitWithRealtimeSafeReceiverBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PullInputBlockFromRealtimeSafeReceiverBlockSync is a synchronous wrapper around [AVAudioSinkNode.PullInputBlockFromRealtimeSafeReceiverBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioSinkNodeClass) PullInputBlockFromRealtimeSafeReceiverBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	ac.PullInputBlockFromRealtimeSafeReceiverBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
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
