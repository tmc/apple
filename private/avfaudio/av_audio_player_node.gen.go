// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioPlayerNode] class.
var (
	_AVAudioPlayerNodeClass     AVAudioPlayerNodeClass
	_AVAudioPlayerNodeClassOnce sync.Once
)

func getAVAudioPlayerNodeClass() AVAudioPlayerNodeClass {
	_AVAudioPlayerNodeClassOnce.Do(func() {
		_AVAudioPlayerNodeClass = AVAudioPlayerNodeClass{class: objc.GetClass("AVAudioPlayerNode")}
	})
	return _AVAudioPlayerNodeClass
}

// GetAVAudioPlayerNodeClass returns the class object for AVAudioPlayerNode.
func GetAVAudioPlayerNodeClass() AVAudioPlayerNodeClass {
	return getAVAudioPlayerNodeClass()
}

type AVAudioPlayerNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPlayerNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPlayerNodeClass) Alloc() AVAudioPlayerNode {
	rv := objc.SendIfResponds[AVAudioPlayerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioPlayerNode.CallLegacyCompletionHandlerForTypeLegacyHandler]
//   - [AVAudioPlayerNode.DebugDescription]
//   - [AVAudioPlayerNode.Description]
//   - [AVAudioPlayerNode.Hash]
//   - [AVAudioPlayerNode.Playing]
//   - [AVAudioPlayerNode.Superclass]
type AVAudioPlayerNode struct {
	AVAudioNode
}

// AVAudioPlayerNodeFromID constructs a [AVAudioPlayerNode] from an objc.ID.
func AVAudioPlayerNodeFromID(id objc.ID) AVAudioPlayerNode {
	return AVAudioPlayerNode{AVAudioNode: AVAudioNodeFromID(id)}
}

// Ensure AVAudioPlayerNode implements IAVAudioPlayerNode.
var _ IAVAudioPlayerNode = AVAudioPlayerNode{}

// An interface definition for the [AVAudioPlayerNode] class.
//
// # Methods
//
//   - [IAVAudioPlayerNode.CallLegacyCompletionHandlerForTypeLegacyHandler]
//   - [IAVAudioPlayerNode.DebugDescription]
//   - [IAVAudioPlayerNode.Description]
//   - [IAVAudioPlayerNode.Hash]
//   - [IAVAudioPlayerNode.Playing]
//   - [IAVAudioPlayerNode.Superclass]
type IAVAudioPlayerNode interface {
	IAVAudioNode

	// Topic: Methods

	CallLegacyCompletionHandlerForTypeLegacyHandler(type_ int64, handler VoidHandler)
	DebugDescription() string
	Description() string
	Hash() uint64
	Playing() bool
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (a AVAudioPlayerNode) Init() AVAudioPlayerNode {
	rv := objc.SendIfResponds[AVAudioPlayerNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPlayerNode) Autorelease() AVAudioPlayerNode {
	rv := objc.SendIfResponds[AVAudioPlayerNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPlayerNode creates a new AVAudioPlayerNode instance.
func NewAVAudioPlayerNode() AVAudioPlayerNode {
	class := getAVAudioPlayerNodeClass()
	rv := objc.SendIfResponds[AVAudioPlayerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioPlayerNodeWithImpl(impl unsafe.Pointer) AVAudioPlayerNode {
	instance := getAVAudioPlayerNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioPlayerNodeFromID(rv)
}

func (a AVAudioPlayerNode) CallLegacyCompletionHandlerForTypeLegacyHandler(type_ int64, handler VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("callLegacyCompletionHandlerForType:legacyHandler:"), type_, _block1)
}

func (a AVAudioPlayerNode) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioPlayerNode) Description() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioPlayerNode) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("hash"))
	return rv
}
func (a AVAudioPlayerNode) Playing() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("playing"))
	return rv
}
func (a AVAudioPlayerNode) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](a.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// CallLegacyCompletionHandlerForTypeLegacyHandlerSync is a synchronous wrapper around [AVAudioPlayerNode.CallLegacyCompletionHandlerForTypeLegacyHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioPlayerNode) CallLegacyCompletionHandlerForTypeLegacyHandlerSync(ctx context.Context, type_ int64) error {
	done := make(chan struct{}, 1)
	a.CallLegacyCompletionHandlerForTypeLegacyHandler(type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
