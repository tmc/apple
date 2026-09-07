// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	rv := objc.SendIfResponds[AVAudioSourceNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioSourceNode.InitWithFormatRealtimeSafeRenderBlock]
//   - [AVAudioSourceNode.InitWithRealtimeSafeRenderBlock]
//   - [AVAudioSourceNode.DebugDescription]
//   - [AVAudioSourceNode.Description]
//   - [AVAudioSourceNode.Hash]
//   - [AVAudioSourceNode.Superclass]
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
//   - [IAVAudioSourceNode.InitWithFormatRealtimeSafeRenderBlock]
//   - [IAVAudioSourceNode.InitWithRealtimeSafeRenderBlock]
//   - [IAVAudioSourceNode.DebugDescription]
//   - [IAVAudioSourceNode.Description]
//   - [IAVAudioSourceNode.Hash]
//   - [IAVAudioSourceNode.Superclass]
type IAVAudioSourceNode interface {
	IAVAudioNode

	// Topic: Methods

	InitWithFormatRealtimeSafeRenderBlock(format objectivec.IObject, block VoidHandler) AVAudioSourceNode
	InitWithRealtimeSafeRenderBlock(block VoidHandler) AVAudioSourceNode
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (a AVAudioSourceNode) Init() AVAudioSourceNode {
	rv := objc.SendIfResponds[AVAudioSourceNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSourceNode) Autorelease() AVAudioSourceNode {
	rv := objc.SendIfResponds[AVAudioSourceNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSourceNode creates a new AVAudioSourceNode instance.
func NewAVAudioSourceNode() AVAudioSourceNode {
	class := getAVAudioSourceNodeClass()
	rv := objc.SendIfResponds[AVAudioSourceNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAVAudioSourceNodeWithImpl(impl unsafe.Pointer) AVAudioSourceNode {
	instance := getAVAudioSourceNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioSourceNodeFromID(rv)
}

var _avaudiosourcenode_initwithformat_realtimesaferenderblock_p1_key byte

func (a AVAudioSourceNode) InitWithFormatRealtimeSafeRenderBlock(format objectivec.IObject, block VoidHandler) AVAudioSourceNode {
	_block1, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[AVAudioSourceNode](a.ID, objc.Sel("initWithFormat:realtimeSafeRenderBlock:"), format, _block1)
	return rv
}

var _avaudiosourcenode_initwithrealtimesaferenderblock_p0_key byte

func (a AVAudioSourceNode) InitWithRealtimeSafeRenderBlock(block VoidHandler) AVAudioSourceNode {
	_block0, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[AVAudioSourceNode](a.ID, objc.Sel("initWithRealtimeSafeRenderBlock:"), _block0)
	return rv
}

func (_AVAudioSourceNodeClass AVAudioSourceNodeClass) PullInputBlockFromRealtimeSafeRenderBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](objc.ID(_AVAudioSourceNodeClass.class), objc.Sel("pullInputBlockFromRealtimeSafeRenderBlock:"), _block0)
}
func (_AVAudioSourceNodeClass AVAudioSourceNodeClass) PullInputBlockFromRenderBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](objc.ID(_AVAudioSourceNodeClass.class), objc.Sel("pullInputBlockFromRenderBlock:"), _block0)
}

func (a AVAudioSourceNode) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioSourceNode) Description() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioSourceNode) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("hash"))
	return rv
}
func (a AVAudioSourceNode) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](a.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// InitWithFormatRealtimeSafeRenderBlockSync is a synchronous wrapper around [AVAudioSourceNode.InitWithFormatRealtimeSafeRenderBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioSourceNode) InitWithFormatRealtimeSafeRenderBlockSync(ctx context.Context, format objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.InitWithFormatRealtimeSafeRenderBlock(format, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithRealtimeSafeRenderBlockSync is a synchronous wrapper around [AVAudioSourceNode.InitWithRealtimeSafeRenderBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioSourceNode) InitWithRealtimeSafeRenderBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.InitWithRealtimeSafeRenderBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PullInputBlockFromRealtimeSafeRenderBlockSync is a synchronous wrapper around [AVAudioSourceNode.PullInputBlockFromRealtimeSafeRenderBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioSourceNodeClass) PullInputBlockFromRealtimeSafeRenderBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	ac.PullInputBlockFromRealtimeSafeRenderBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
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
