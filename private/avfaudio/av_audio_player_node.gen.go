// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.Send[AVAudioPlayerNode](objc.ID(ac.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode
type IAVAudioPlayerNode interface {
	IAVAudioNode

	// Topic: Methods

	CallLegacyCompletionHandlerForTypeLegacyHandler(type_ int64, handler VoidHandler)
	DebugDescription() string
	Description() string
	Hash() uint64
	Playing() bool
	Superclass() objc.Class
}

// Init initializes the instance.
func (a AVAudioPlayerNode) Init() AVAudioPlayerNode {
	rv := objc.Send[AVAudioPlayerNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPlayerNode) Autorelease() AVAudioPlayerNode {
	rv := objc.Send[AVAudioPlayerNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPlayerNode creates a new AVAudioPlayerNode instance.
func NewAVAudioPlayerNode() AVAudioPlayerNode {
	class := getAVAudioPlayerNodeClass()
	rv := objc.Send[AVAudioPlayerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioPlayerNodeWithImpl(impl unsafe.Pointer) AVAudioPlayerNode {
	instance := getAVAudioPlayerNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioPlayerNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/callLegacyCompletionHandlerForType:legacyHandler:
func (a AVAudioPlayerNode) CallLegacyCompletionHandlerForTypeLegacyHandler(type_ int64, handler VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("callLegacyCompletionHandlerForType:legacyHandler:"), type_, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/debugDescription
func (a AVAudioPlayerNode) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/description
func (a AVAudioPlayerNode) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/hash
func (a AVAudioPlayerNode) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/playing
func (a AVAudioPlayerNode) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("playing"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/superclass
func (a AVAudioPlayerNode) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
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
