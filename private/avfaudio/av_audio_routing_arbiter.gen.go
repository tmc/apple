// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioRoutingArbiter] class.
var (
	_AVAudioRoutingArbiterClass     AVAudioRoutingArbiterClass
	_AVAudioRoutingArbiterClassOnce sync.Once
)

func getAVAudioRoutingArbiterClass() AVAudioRoutingArbiterClass {
	_AVAudioRoutingArbiterClassOnce.Do(func() {
		_AVAudioRoutingArbiterClass = AVAudioRoutingArbiterClass{class: objc.GetClass("AVAudioRoutingArbiter")}
	})
	return _AVAudioRoutingArbiterClass
}

// GetAVAudioRoutingArbiterClass returns the class object for AVAudioRoutingArbiter.
func GetAVAudioRoutingArbiterClass() AVAudioRoutingArbiterClass {
	return getAVAudioRoutingArbiterClass()
}

type AVAudioRoutingArbiterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioRoutingArbiterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioRoutingArbiterClass) Alloc() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsError]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler]
//   - [AVAudioRoutingArbiter.CreateBTSessionWithCategoryModeFlags]
//   - [AVAudioRoutingArbiter.DispatchQueue]
//   - [AVAudioRoutingArbiter.SetDispatchQueue]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter
type AVAudioRoutingArbiter struct {
	objectivec.Object
}

// AVAudioRoutingArbiterFromID constructs a [AVAudioRoutingArbiter] from an objc.ID.
func AVAudioRoutingArbiterFromID(id objc.ID) AVAudioRoutingArbiter {
	return AVAudioRoutingArbiter{objectivec.Object{ID: id}}
}
// Ensure AVAudioRoutingArbiter implements IAVAudioRoutingArbiter.
var _ IAVAudioRoutingArbiter = AVAudioRoutingArbiter{}

// An interface definition for the [AVAudioRoutingArbiter] class.
//
// # Methods
//
//   - [IAVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler]
//   - [IAVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsError]
//   - [IAVAudioRoutingArbiter.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler]
//   - [IAVAudioRoutingArbiter.CreateBTSessionWithCategoryModeFlags]
//   - [IAVAudioRoutingArbiter.DispatchQueue]
//   - [IAVAudioRoutingArbiter.SetDispatchQueue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter
type IAVAudioRoutingArbiter interface {
	objectivec.IObject

	// Topic: Methods

	BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler(category objectivec.IObject, mode objectivec.IObject, options uint64, handler ErrorHandler)
	BeginArbitrationWithAudioSessionCategoryModeOptionsError(category objectivec.IObject, mode objectivec.IObject, options uint64) (bool, error)
	BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler(category int, mode int, flags uint32, handler ErrorHandler)
	CreateBTSessionWithCategoryModeFlags(category int, mode int, flags uint32)
	DispatchQueue() objectivec.Object
	SetDispatchQueue(value objectivec.Object)
}

// Init initializes the instance.
func (a AVAudioRoutingArbiter) Init() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioRoutingArbiter) Autorelease() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioRoutingArbiter creates a new AVAudioRoutingArbiter instance.
func NewAVAudioRoutingArbiter() AVAudioRoutingArbiter {
	class := getAVAudioRoutingArbiterClass()
	rv := objc.Send[AVAudioRoutingArbiter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/beginArbitrationWithAudioSessionCategory:mode:options:completionHandler:
func (a AVAudioRoutingArbiter) BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler(category objectivec.IObject, mode objectivec.IObject, options uint64, handler ErrorHandler) {
_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("beginArbitrationWithAudioSessionCategory:mode:options:completionHandler:"), category, mode, options, _block3)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/beginArbitrationWithAudioSessionCategory:mode:options:error:
func (a AVAudioRoutingArbiter) BeginArbitrationWithAudioSessionCategoryModeOptionsError(category objectivec.IObject, mode objectivec.IObject, options uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("beginArbitrationWithAudioSessionCategory:mode:options:error:"), category, mode, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("beginArbitrationWithAudioSessionCategory:mode:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/beginArbitrationWithBTSessionCategory:mode:flags:completionHandler:
func (a AVAudioRoutingArbiter) BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler(category int, mode int, flags uint32, handler ErrorHandler) {
_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("beginArbitrationWithBTSessionCategory:mode:flags:completionHandler:"), category, mode, flags, _block3)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/createBTSessionWithCategory:mode:flags:
func (a AVAudioRoutingArbiter) CreateBTSessionWithCategoryModeFlags(category int, mode int, flags uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("createBTSessionWithCategory:mode:flags:"), category, mode, flags)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/dispatchQueue
func (a AVAudioRoutingArbiter) DispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a AVAudioRoutingArbiter) SetDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](a.ID, objc.Sel("setDispatchQueue:"), value)
}

// BeginArbitrationWithAudioSessionCategoryModeOptions is a synchronous wrapper around [AVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioRoutingArbiter) BeginArbitrationWithAudioSessionCategoryModeOptions(ctx context.Context, category objectivec.IObject, mode objectivec.IObject, options uint64) error {
	done := make(chan error, 1)
	a.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler(category, mode, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// BeginArbitrationWithBTSessionCategoryModeFlags is a synchronous wrapper around [AVAudioRoutingArbiter.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioRoutingArbiter) BeginArbitrationWithBTSessionCategoryModeFlags(ctx context.Context, category int, mode int, flags uint32) error {
	done := make(chan error, 1)
	a.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler(category, mode, flags, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

